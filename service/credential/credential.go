package credential

import (
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	// AWSOperatorArnKey is the key in the Secret under which the ARN for the
	// aws-operator role is held.
	AWSOperatorArnKey = "aws.awsoperator.arn"
)

const (
	DefaultName      = "credential-default"
	DefaultNamespace = "giantswarm"
)

// GetDefaultARN is used only by the bridgezone resource. It should be removed
// when the resource is removed.
func GetDefaultARN(k8sClient kubernetes.Interface) (string, error) {
	credential, err := k8sClient.CoreV1().Secrets(DefaultNamespace).Get(DefaultName, metav1.GetOptions{})
	if err != nil {
		return "", microerror.Mask(err)
	}

	arn, err := getARN(credential)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return arn, nil
}

func getARN(credential *corev1.Secret) (string, error) {
	arn, ok := credential.Data[AWSOperatorArnKey]
	if !ok {
		return "", nil
	}

	return string(arn), nil
}
