package key

import (
	"github.com/giantswarm/apiextensions/pkg/apis/example/v1alpha1"
	"github.com/giantswarm/microerror"
)

func ToServiceQuota(v interface{}) (v1alpha1.ServiceQuota, error) {
	if v == nil {
		return v1alpha1.ServiceQuota{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.ServiceQuota{}, v)
	}

	p, ok := v.(*v1alpha1.ServiceQuota)
	if !ok {
		return v1alpha1.ServiceQuota{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &v1alpha1.ServiceQuota{}, v)
	}

	c := p.DeepCopy()

	return *c, nil
}
