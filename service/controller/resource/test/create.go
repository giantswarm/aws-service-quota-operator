package test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/giantswarm/aws-service-quota-operator/service/credential"
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

	r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("Creating %+v\n", obj))

	arn := ""
	if serviceQuota.Spec.Organization == "" {
		r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("Default ARN"))
		arn, _ = credential.GetDefaultARN(r.k8sClient)
	} else {
		r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("Organization ARN"))
		params := organizations.NewGetCredentialsParams()
		params.OrganizationID = serviceQuota.Spec.Organization

		orgCredentials, err := client.Organizations.GetCredentials(params, auth)
		if err != nil {
			return microerror.Mask(err)
		}
		arn = orgCredentials.Payload[0].Aws.Roles.Admin
	}

	r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("Credentials ARN %s", arn))

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		return microerror.Mask(err)
	}
	creds := stscreds.NewCredentials(sess, arn)

	conn := servicequotas.New(sess, &aws.Config{Credentials: creds})

	r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("Service code %s", serviceQuota.Spec.ServiceCode))
	r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("Quota code %s", serviceQuota.Spec.QuotaCode))

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
