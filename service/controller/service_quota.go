package controller

import (
	"fmt"

	"github.com/giantswarm/apiextensions/pkg/apis/example/v1alpha1"

	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/aws/aws-sdk-go/aws"

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
	var test = aws.String("bucket")
	fmt.Printf("%+v\n", test)

	resourceSets, err := newServiceQuotaResourceSets(config)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var operatorkitController *controller.Controller
	{
		c := controller.Config{
			CRD:          v1alpha1.NewServiceQuotaCRD(),
			K8sClient:    config.K8sClient,
			Logger:       config.Logger,
			ResourceSets: resourceSets,
			NewRuntimeObjectFunc: func() runtime.Object {
				return new(corev1.Pod)
			},

			// Name is used to compute finalizer names. This here results in something
			// like operatorkit.giantswarm.io/github-search-index-operator-ServiceQuota-controller.
			Name: project.Name() + "-ServiceQuota-controller",
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
			K8sClient: config.K8sClient,
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
