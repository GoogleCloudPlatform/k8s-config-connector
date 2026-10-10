// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

import (
	"context"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/iamrefs"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestContainerNodePool_NormalizeReferences_SecretRef(t *testing.T) {
	ctx := context.Background()

	readySecretVersion := &unstructured.Unstructured{}
	readySecretVersion.SetGroupVersionKind(secretmanagerv1beta1.SecretManagerSecretVersionGVK)
	readySecretVersion.SetName("test-secret-version")
	readySecretVersion.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(readySecretVersion.Object, "projects/test-project/secrets/test-secret/versions/1", "status", "externalRef"); err != nil {
		t.Fatalf("failed to set status.externalRef: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readySecretVersion).Build()

	clusterRef := krm.ContainerClusterRef{
		External: "projects/test-project/locations/us-central1-a/clusters/test-cluster",
	}

	t.Run("resolves SecretVersionRef in ContainerdConfig by name via status.externalRef", func(t *testing.T) {
		obj := &krm.ContainerNodePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pool",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerNodePoolSpec{
				ClusterRef: clusterRef,
				NodeConfig: &krm.NodePoolNodeConfig{
					ContainerdConfig: &krm.ContainerdConfig{
						PrivateRegistryAccessConfig: &krm.PrivateRegistryAccessConfig{
							CertificateAuthorityDomainConfig: []krm.CertificateAuthorityDomainConfig{
								{
									Fqdns: []string{"example.com"},
									GCPSecretManagerCertificateConfig: &krm.GCPSecretManagerCertificateConfig{
										SecretRef: &secretmanagerv1beta1.SecretVersionRef{
											Name: "test-secret-version",
										},
									},
								},
							},
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		want := "projects/test-project/secrets/test-secret/versions/1"
		got := obj.Spec.NodeConfig.ContainerdConfig.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig[0].GCPSecretManagerCertificateConfig.SecretRef.External
		if got != want {
			t.Errorf("NormalizeReferences() External = %q, want %q", got, want)
		}
	})

	t.Run("resolves external SecretVersionRef directly", func(t *testing.T) {
		externalURI := "projects/test-project/secrets/test-secret/versions/2"
		obj := &krm.ContainerNodePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pool",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerNodePoolSpec{
				ClusterRef: clusterRef,
				NodeConfig: &krm.NodePoolNodeConfig{
					ContainerdConfig: &krm.ContainerdConfig{
						PrivateRegistryAccessConfig: &krm.PrivateRegistryAccessConfig{
							CertificateAuthorityDomainConfig: []krm.CertificateAuthorityDomainConfig{
								{
									Fqdns: []string{"example.com"},
									GCPSecretManagerCertificateConfig: &krm.GCPSecretManagerCertificateConfig{
										SecretRef: &secretmanagerv1beta1.SecretVersionRef{
											External: externalURI,
										},
									},
								},
							},
						},
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		got := obj.Spec.NodeConfig.ContainerdConfig.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig[0].GCPSecretManagerCertificateConfig.SecretRef.External
		if got != externalURI {
			t.Errorf("NormalizeReferences() External = %q, want %q", got, externalURI)
		}
	})
}

func TestContainerNodePool_NormalizeReferences_NodeConfigRefs(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-key")
	readyKey.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(readyKey.Object, "projects/test-project/locations/us/keyRings/test-keyring/cryptoKeys/test-key", "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	readySA := &unstructured.Unstructured{}
	readySA.SetGroupVersionKind(iamrefs.IAMServiceAccountGVK)
	readySA.SetName("test-sa")
	readySA.SetNamespace("test-ns")
	if err := unstructured.SetNestedField(readySA.Object, "test-sa@test-project.iam.gserviceaccount.com", "status", "email"); err != nil {
		t.Fatalf("failed to set status.email: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey, readySA).Build()

	clusterRef := krm.ContainerClusterRef{
		External: "projects/test-project/locations/us-central1-a/clusters/test-cluster",
	}

	t.Run("resolves BootDiskKMSCryptoKeyRef and ServiceAccountRef by name", func(t *testing.T) {
		obj := &krm.ContainerNodePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pool",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerNodePoolSpec{
				ClusterRef: clusterRef,
				NodeConfig: &krm.NodePoolNodeConfig{
					BootDiskKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						Name: "test-key",
					},
					ServiceAccountRef: &iamrefs.IAMServiceAccountRef{
						Name: "test-sa",
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		wantKey := "projects/test-project/locations/us/keyRings/test-keyring/cryptoKeys/test-key"
		if obj.Spec.NodeConfig.BootDiskKMSCryptoKeyRef.External != wantKey {
			t.Errorf("NormalizeReferences() BootDiskKMSCryptoKeyRef.External = %q, want %q", obj.Spec.NodeConfig.BootDiskKMSCryptoKeyRef.External, wantKey)
		}

		wantSA := "test-sa@test-project.iam.gserviceaccount.com"
		if obj.Spec.NodeConfig.ServiceAccountRef.External != wantSA {
			t.Errorf("NormalizeReferences() ServiceAccountRef.External = %q, want %q", obj.Spec.NodeConfig.ServiceAccountRef.External, wantSA)
		}

		mapCtx := &direct.MapContext{}
		proto := ContainerNodePoolSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("ContainerNodePoolSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if proto.Config.BootDiskKmsKey != wantKey {
			t.Errorf("proto BootDiskKmsKey = %q, want %q", proto.Config.BootDiskKmsKey, wantKey)
		}
		if proto.Config.ServiceAccount != wantSA {
			t.Errorf("proto ServiceAccount = %q, want %q", proto.Config.ServiceAccount, wantSA)
		}
	})

	t.Run("handles external BootDiskKMSCryptoKeyRef and ServiceAccountRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us/keyRings/test-keyring/cryptoKeys/external-key"
		externalSA := "external-sa@test-project.iam.gserviceaccount.com"
		obj := &krm.ContainerNodePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pool",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerNodePoolSpec{
				ClusterRef: clusterRef,
				NodeConfig: &krm.NodePoolNodeConfig{
					BootDiskKMSCryptoKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
						External: externalKey,
					},
					ServiceAccountRef: &iamrefs.IAMServiceAccountRef{
						External: externalSA,
					},
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.NodeConfig.BootDiskKMSCryptoKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() BootDiskKMSCryptoKeyRef.External = %q, want %q", obj.Spec.NodeConfig.BootDiskKMSCryptoKeyRef.External, externalKey)
		}
		if obj.Spec.NodeConfig.ServiceAccountRef.External != externalSA {
			t.Errorf("NormalizeReferences() ServiceAccountRef.External = %q, want %q", obj.Spec.NodeConfig.ServiceAccountRef.External, externalSA)
		}
	})
}
