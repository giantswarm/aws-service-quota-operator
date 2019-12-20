package controller

import (
	"github.com/giantswarm/apiextensions/pkg/apis/example/v1alpha1"
	"github.com/giantswarm/k8sclient"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/giantswarm/aws-service-quota-operator/pkg/project"
)

type ServiceQuotaConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

type ServiceQuota struct {
	*controller.Controller
}

func NewServiceQuota(config ServiceQuotaConfig) (*ServiceQuota, error) {
	var err error

	resourceSets, err := newServiceQuotaResourceSets(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var operatorkitController *controller.Controller
	{
		c := controller.Config{
			CRD:          v1alpha1.NewServiceQuotaCRD(),
			Logger:       config.Logger,
			ResourceSets: resourceSets,
			NewRuntimeObjectFunc: func() runtime.Object {
				return new(v1alpha1.ServiceQuota)
			},
			Name: project.Name() + "-controller",
		}

		operatorkitController, err = controller.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c := &ServiceQuota{
		Controller: operatorkitController,
	}

	return c, nil
}

func newServiceQuotaResourceSets(config ServiceQuotaConfig) ([]*controller.ResourceSet, error) {
	var err error

	var resourceSet *controller.ResourceSet
	{
		c := ServiceQuotaResourceSetConfig{
			K8sClient: config.K8sClient.K8sClient(),
			Logger:    config.Logger,
		}

		resourceSet, err = newServiceQuotaResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	resourceSets := []*controller.ResourceSet{
		resourceSet,
	}

	return resourceSets, nil
}
