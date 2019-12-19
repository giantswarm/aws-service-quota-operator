package test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/giantswarm/aws-service-quota-operator/service/key"
	gsclient "github.com/giantswarm/gsclientgen/client"
	"github.com/giantswarm/gsclientgen/client/organizations"
	"github.com/giantswarm/microerror"
	httptransport "github.com/go-openapi/runtime/client"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	serviceQuota, err := key.ToServiceQuota(obj)
	if err != nil {
		return err
	}

	var endpointURL *url.URL
	{
		endpointURL, err = url.Parse("http://credentiald:8000")
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var client *gsclient.Gsclientgen
	{
		tp := httptransport.New("credentiald:8000", "", []string{endpointURL.Scheme})
		tp.Transport = &http.Transport{}
		client = gsclient.New(tp, strfmt.Default)
	}

	var auth runtime.ClientAuthInfoWriter
	{
	}

	params := organizations.NewGetCredentialsParams()
	params.OrganizationID = serviceQuota.Spec.Account

	response, err := client.Organizations.GetCredentials(params, auth)
	if err != nil {
		return microerror.Mask(err)
	}

	r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("Creating %+v\n", obj))
	r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("Credentials %+v\n", response))

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		return microerror.Mask(err)
	}

	conn := servicequotas.New(sess)
	input := &servicequotas.GetServiceQuotaInput{
		ServiceCode: aws.String(serviceQuota.Spec.ServiceCode),
		QuotaCode:   aws.String(serviceQuota.Spec.QuotaCode),
	}
	output, err := conn.GetServiceQuota(input)

	if err != nil {
		r.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("error getting Service Quotas Service Quota : %s", err))
		return microerror.Mask(err)
	}
	if output == nil {
		r.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("error getting Service Quotas Service Quota: empty result"))
		return microerror.Mask(err)
	}

	return nil
}
