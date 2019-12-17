package v1alpha1

import (
	"github.com/ghodss/yaml"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindServiceQuota        = "ServiceQuota"
	kindServiceQuotaRequest = "ServiceQuotaRequest"
)

const serviceQuotaCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: servicequota.example.giantswarm.io
spec:
  group: example.giantswarm.io
  scope: Namespaced
  version: v1alpha1
  names:
    kind: ServiceQuota
    plural: servicequotas
    singular: servicequota
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          type: object
          properties:
            quotaCode:
              type: string
            serviceCode:
			  type: string
			account:
			  type: string
          required:
          - quotaCode
		  - serviceCode
		  - account
        status:
          type: object
          properties:
		  	currentValue:
				type: resource.Quantity
			serviceName:
				type: string
			unit:
				type: string
`

const serviceQuotaRequestCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: servicequotarequest.example.giantswarm.io
spec:
  group: example.giantswarm.io
  scope: Namespaced
  version: v1alpha1
  names:
    kind: ServiceQuotaRequest
    plural: servicequotarequests
    singular: servicequotarequest
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          type: object
          properties:
            quotaCode:
              type: string
            serviceCode:
			  type: string
			account:
			  type: string
			desiredValue:
			  type: resource.Quantity
          required:
          - quotaCode
		  - serviceCode
		  - account
		  - desiredValue
        status:
          type: object
          properties:
		  	status:
			  enum:
			  - pending
			  - requested
			  - success
			  - error
			statusMessage:
			  type: string
`

var serviceQuotaCRD *apiextensionsv1beta1.CustomResourceDefinition
var serviceQuotaRequestCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(serviceQuotaCRDYAML), &serviceQuotaCRD)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(serviceQuotaRequestCRDYAML), &serviceQuotaRequestCRD)
	if err != nil {
		panic(err)
	}
}

func NewServiceQuotaCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return serviceQuotaCRD.DeepCopy()
}

func NewServiceQuotaRequestCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return serviceQuotaRequestCRD.DeepCopy()
}

func NewServiceQuotaTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindServiceQuota,
	}
}

func NewServiceQuotaRequestTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindServiceQuotaRequest,
	}
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceQuota CRs might look something like the following.
//
//	apiVersion: "example.giantswarm.io/v1alpha1"
//	kind: "ServiceQuota"
//	metadata:
//	  name: "cloudformation-stackcount-orgname"
//	  labels:
//	    giantswarm.io/managed-by: "aws-service-quota-operator"
//	    giantswarm.io/provider: "aws"
//	spec:
//	  quotaCode: "L-0485CB21"
//	  serviceCode: "cloudformation"
//	  account: "bacedfd6e09d98eacf"
//
type ServiceQuota struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              ServiceQuotaSpec   `json:"spec" yaml:"spec"`
	Status            ServiceQuotaStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type ServiceQuotaSpec struct {
	// Specifies the service that you want to use.
	ServiceCode string `json:"serviceCode" yaml:"serviceCode"`
	// The code identifier for the service quota specified.
	QuotaCode string `json:"cuotaCode" yaml:"cuotaCode"`
	// AWS account to check service quota
	Account string `json:"account" yaml:"account"`
}

type ServiceQuotaStatus struct {
	// The name of the AWS service specified in the increase request
	ServiceName string `json:"serviceName" yaml:"serviceName"`
	// Current Value represents the limit amount for this resource
	CurrentValue resource.Quantity `json:"currentValue" yaml:"currentValue"`
	// If the Service Quota has a unit it will be stored here
	Unit string `json:"unit" yaml:"unit"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServiceQuotaList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []ServiceQuota `json:"items" yaml:"items"`
}

// ServiceQuota CRs might look something like the following.
//
//	apiVersion: "example.giantswarm.io/v1alpha1"
//	kind: "ServiceQuotaRequest"
//	metadata:
//	  name: "cloudformation-stackcount-orgname-increase1000"
//	  labels:
//	    giantswarm.io/managed-by: "aws-service-quota-operator"
//	    giantswarm.io/provider: "aws"
//	spec:
//	  quotaCode: "L-0485CB21"
//	  serviceCode: "cloudformation"
//	  account: "bacedfd6e09d98eacf"
//	  desiredValue: 1200
//
type ServiceQuotaRequest struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`
	Spec              ServiceQuotaSpec   `json:"spec" yaml:"spec"`
	Status            ServiceQuotaStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

type ServiceQuotaRequestSpec struct {
	// Specifies the service that you want to use.
	ServiceCode string `json:"serviceCode" yaml:"serviceCode"`
	// The code identifier for the service quota specified.
	QuotaCode string `json:"cuotaCode" yaml:"cuotaCode"`
	// AWS account to check service quota
	Account string `json:"account" yaml:"account"`
	// Specifies the new value for the quota.
	DesiredValue resource.Quantity `json:"desiredValue" yaml:"desiredValue"`
}

type ServiceQuotaRequestStatus struct {
	// Specifies the status value of the quota increase request.
	Status string `json:"status" yaml:"status"`
	// Current Value represents the limit amount for this resource
	StatusMessage string `json:"statusMessage" yaml:"statusMessage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServiceQuotaRequestList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []ServiceQuotaRequest `json:"items" yaml:"items"`
}
