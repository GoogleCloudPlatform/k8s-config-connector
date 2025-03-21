package direct

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

func KMSCryptoKeyRef(key string) *refsv1beta1.KMSCryptoKeyRef {
	return &refsv1beta1.KMSCryptoKeyRef{
		External: key,
	}
}
