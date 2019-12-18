package test

import (
	"context"
	"fmt"

	"github.com/giantswarm/aws-service-quota-operator/service/key"
)

func (r *Resource) EnsureDeleted(ctx context.Context, obj interface{}) error {
	_, err := key.ToServiceQuota(obj)
	if err != nil {
		return err
	}

	r.logger.LogCtx(ctx, "level", "info", "message", fmt.Sprintf("Deleting %+v\n", obj))

	return nil
}
