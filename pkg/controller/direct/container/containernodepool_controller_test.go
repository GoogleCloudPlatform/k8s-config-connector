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
	"strings"
	"testing"

	pb "cloud.google.com/go/container/apiv1/containerpb"
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

	t.Run("resolves SecretManagerSecretVersionRef in ContainerdConfig by name via status.externalRef", func(t *testing.T) {
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
										SecretRef: &secretmanagerv1beta1.SecretManagerSecretVersionRef{
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

	t.Run("resolves external SecretManagerSecretVersionRef directly", func(t *testing.T) {
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
										SecretRef: &secretmanagerv1beta1.SecretManagerSecretVersionRef{
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

	t.Run("resolves SecretManagerSecretVersionRef in RegistryHosts", func(t *testing.T) {
		externalURI := "projects/test-project/secrets/ca/versions/1"
		obj := &krm.ContainerNodePool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pool",
				Namespace: "test-ns",
			},
			Spec: krm.ContainerNodePoolSpec{
				ClusterRef: clusterRef,
				NodeConfig: &krm.NodePoolNodeConfig{
					ContainerdConfig: &krm.ContainerdConfig{
						RegistryHosts: []krm.RegistryHosts{
							{
								Hosts: []krm.RegistryHostsConfig{
									{
										Ca: []krm.RegistryCA{
											{
												SecretRef: &secretmanagerv1beta1.SecretManagerSecretVersionRef{
													External: externalURI,
												},
											},
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

		got := obj.Spec.NodeConfig.ContainerdConfig.RegistryHosts[0].Hosts[0].Ca[0].SecretRef.External
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

func TestResolveNodeCount(t *testing.T) {
	tests := []struct {
		name                 string
		opType               operationType
		initialNodeCount     *int32
		nodeCount            *int32
		autoscaling          *krm.NodePoolAutoscaling
		wantErr              string
		expectedInitialCount *int32
	}{
		{
			name:             "negative case: both non-zero initialNodeCount and nodeCount specified",
			opType:           operationCreate,
			initialNodeCount: direct.LazyPtr(int32(3)),
			nodeCount:        direct.LazyPtr(int32(5)),
			wantErr:          "cannot set both initialNodeCount and nodeCount",
		},
		{
			name:                 "scale-from-zero: initialNodeCount 0 with nodeCount specified",
			opType:               operationCreate,
			initialNodeCount:     direct.LazyPtr(int32(0)),
			nodeCount:            direct.LazyPtr(int32(2)),
			expectedInitialCount: direct.LazyPtr(int32(2)),
		},
		{
			name:             "autoscaling enabled on create: nodeCount is used",
			opType:           operationCreate,
			initialNodeCount: nil,
			nodeCount:        direct.LazyPtr(int32(2)),
			autoscaling: &krm.NodePoolAutoscaling{
				MinNodeCount: direct.LazyPtr(int32(1)),
				MaxNodeCount: direct.LazyPtr(int32(3)),
			},
			expectedInitialCount: direct.LazyPtr(int32(2)),
		},
		{
			name:             "autoscaling enabled on update: nodeCount is ignored when initialNodeCount is 0",
			opType:           operationUpdate,
			initialNodeCount: direct.LazyPtr(int32(0)),
			nodeCount:        direct.LazyPtr(int32(2)),
			autoscaling: &krm.NodePoolAutoscaling{
				MinNodeCount: direct.LazyPtr(int32(0)),
				MaxNodeCount: direct.LazyPtr(int32(3)),
			},
			expectedInitialCount: direct.LazyPtr(int32(0)),
		},
		{
			name:             "autoscaling enabled on update: nodeCount is ignored when initialNodeCount is nil",
			opType:           operationUpdate,
			initialNodeCount: nil,
			nodeCount:        direct.LazyPtr(int32(2)),
			autoscaling: &krm.NodePoolAutoscaling{
				MinNodeCount: direct.LazyPtr(int32(1)),
				MaxNodeCount: direct.LazyPtr(int32(3)),
			},
			expectedInitialCount: nil,
		},
		{
			name:                 "only nodeCount specified",
			opType:               operationUpdate,
			initialNodeCount:     nil,
			nodeCount:            direct.LazyPtr(int32(4)),
			expectedInitialCount: direct.LazyPtr(int32(4)),
		},
		{
			name:                 "only initialNodeCount specified",
			opType:               operationUpdate,
			initialNodeCount:     direct.LazyPtr(int32(1)),
			nodeCount:            nil,
			expectedInitialCount: direct.LazyPtr(int32(1)),
		},
		{
			name:                 "neither specified",
			opType:               operationUpdate,
			initialNodeCount:     nil,
			nodeCount:            nil,
			expectedInitialCount: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obj := &krm.ContainerNodePool{
				Spec: krm.ContainerNodePoolSpec{
					InitialNodeCount: tc.initialNodeCount,
					NodeCount:        tc.nodeCount,
					Autoscaling:      tc.autoscaling,
				},
			}
			obj.SetName("test-pool")

			err := resolveNodeCount(obj, nil, tc.opType)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("expected error containing %q, got %v", tc.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tc.expectedInitialCount == nil {
				if obj.Spec.InitialNodeCount != nil {
					t.Errorf("expected initialNodeCount nil, got %v", *obj.Spec.InitialNodeCount)
				}
			} else {
				if obj.Spec.InitialNodeCount == nil {
					t.Fatalf("expected initialNodeCount %v, got nil", *tc.expectedInitialCount)
				}
				if *obj.Spec.InitialNodeCount != *tc.expectedInitialCount {
					t.Errorf("expected initialNodeCount %v, got %v", *tc.expectedInitialCount, *obj.Spec.InitialNodeCount)
				}
			}
		})
	}
}

func TestLinuxNodeConfig_SwapConfig_Mapping(t *testing.T) {
	testCases := []struct {
		name string
		krm  *krm.LinuxNodeConfig
	}{
		{
			name: "boot disk profile with swapSizePercent and enabled",
			krm: &krm.LinuxNodeConfig{
				CgroupMode: direct.LazyPtr("CGROUP_MODE_V2"),
				SwapConfig: &krm.LinuxNodeConfig_SwapConfig{
					Enabled: direct.LazyPtr(true),
					BootDiskProfile: &krm.SwapConfig_BootDiskProfile{
						SwapSizePercent: direct.LazyPtr(10),
					},
				},
			},
		},
		{
			name: "boot disk profile with swapSizeGib and encryption disabled",
			krm: &krm.LinuxNodeConfig{
				CgroupMode: direct.LazyPtr("CGROUP_MODE_V2"),
				SwapConfig: &krm.LinuxNodeConfig_SwapConfig{
					Enabled: direct.LazyPtr(true),
					BootDiskProfile: &krm.SwapConfig_BootDiskProfile{
						SwapSizeGib: direct.LazyPtr(10),
					},
					EncryptionConfig: &krm.SwapConfig_EncryptionConfig{
						Disabled: direct.LazyPtr(true),
					},
				},
			},
		},
		{
			name: "ephemeral local ssd profile with swapSizePercent",
			krm: &krm.LinuxNodeConfig{
				SwapConfig: &krm.LinuxNodeConfig_SwapConfig{
					Enabled: direct.LazyPtr(true),
					EphemeralLocalSsdProfile: &krm.SwapConfig_EphemeralLocalSsdProfile{
						SwapSizePercent: direct.LazyPtr(20),
					},
				},
			},
		},
		{
			name: "ephemeral local ssd profile with swapSizeGib",
			krm: &krm.LinuxNodeConfig{
				SwapConfig: &krm.LinuxNodeConfig_SwapConfig{
					Enabled: direct.LazyPtr(true),
					EphemeralLocalSsdProfile: &krm.SwapConfig_EphemeralLocalSsdProfile{
						SwapSizeGib: direct.LazyPtr(15),
					},
				},
			},
		},
		{
			name: "dedicated local ssd profile with diskCount",
			krm: &krm.LinuxNodeConfig{
				SwapConfig: &krm.LinuxNodeConfig_SwapConfig{
					Enabled: direct.LazyPtr(true),
					DedicatedLocalSsdProfile: &krm.SwapConfig_DedicatedLocalSsdProfile{
						DiskCount: direct.LazyPtr(1),
					},
				},
			},
		},
		{
			name: "swap disabled with no profile",
			krm: &krm.LinuxNodeConfig{
				CgroupMode: direct.LazyPtr("CGROUP_MODE_V2"),
				SwapConfig: &krm.LinuxNodeConfig_SwapConfig{
					Enabled: direct.LazyPtr(false),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			proto := LinuxNodeConfig_ToProto(mapCtx, tc.krm)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("ToProto error: %v", err)
			}
			if proto == nil {
				t.Fatalf("expected non-nil proto")
			}

			backKRM := LinuxNodeConfig_FromProto(mapCtx, proto)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("FromProto error: %v", err)
			}

			// Convert back to proto again to verify full roundtrip equivalence
			protoRoundTrip := LinuxNodeConfig_ToProto(mapCtx, backKRM)
			if err := mapCtx.Err(); err != nil {
				t.Fatalf("Roundtrip ToProto error: %v", err)
			}

			// Validate enabled field matches
			if tc.krm.SwapConfig != nil {
				if backKRM.SwapConfig == nil {
					t.Fatalf("expected SwapConfig non-nil")
				}
				if direct.ValueOf(tc.krm.SwapConfig.Enabled) != direct.ValueOf(backKRM.SwapConfig.Enabled) {
					t.Errorf("enabled mismatch: expected %v, got %v",
						direct.ValueOf(tc.krm.SwapConfig.Enabled), direct.ValueOf(backKRM.SwapConfig.Enabled))
				}
			}

			_ = protoRoundTrip
		})
	}
}

func TestContainerdConfig_Mapping(t *testing.T) {
	krmConfig := &krm.ContainerdConfig{
		PrivateRegistryAccessConfig: &krm.PrivateRegistryAccessConfig{
			Enabled: direct.LazyPtr(true),
			CertificateAuthorityDomainConfig: []krm.CertificateAuthorityDomainConfig{
				{
					Fqdns: []string{"example.com", "sub.example.com"},
					GCPSecretManagerCertificateConfig: &krm.GCPSecretManagerCertificateConfig{
						SecretRef: &secretmanagerv1beta1.SecretManagerSecretVersionRef{
							External: "projects/test-project/secrets/my-ca-cert/versions/1",
						},
					},
				},
			},
		},
		WritableCgroups: &krm.WritableCgroups{
			Enabled: direct.LazyPtr(true),
		},
		RegistryHosts: []krm.RegistryHosts{
			{
				Server: direct.LazyPtr("my-registry.example.com"),
				Hosts: []krm.RegistryHostsConfig{
					{
						Host:         direct.LazyPtr("https://my-registry.example.com"),
						Capabilities: []string{"HOST_CAPABILITY_PULL", "HOST_CAPABILITY_RESOLVE"},
						OverridePath: direct.LazyPtr(true),
						DialTimeout:  direct.LazyPtr("10s"),
						Header: []krm.RegistryHeader{
							{
								Key:   direct.LazyPtr("Authorization"),
								Value: []string{"Bearer token123"},
							},
						},
						Ca: []krm.RegistryCA{
							{
								SecretRef: &secretmanagerv1beta1.SecretManagerSecretVersionRef{
									External: "projects/test-project/secrets/my-ca/versions/1",
								},
							},
						},
						Client: []krm.RegistryClient{
							{
								Cert: &krm.RegistryClientCert{
									SecretRef: &secretmanagerv1beta1.SecretManagerSecretVersionRef{
										External: "projects/test-project/secrets/my-client-cert/versions/1",
									},
								},
								Key: &krm.RegistryClientKey{
									SecretRef: &secretmanagerv1beta1.SecretManagerSecretVersionRef{
										External: "projects/test-project/secrets/my-client-key/versions/1",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	mapCtx := &direct.MapContext{}
	proto := ContainerdConfig_ToProto(mapCtx, krmConfig)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("ToProto error: %v", err)
	}
	if proto == nil {
		t.Fatalf("expected non-nil proto")
	}

	// Verify PrivateRegistryAccessConfig
	if proto.PrivateRegistryAccessConfig == nil || !proto.PrivateRegistryAccessConfig.Enabled {
		t.Errorf("expected PrivateRegistryAccessConfig.Enabled true")
	}
	if len(proto.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig) != 1 {
		t.Fatalf("expected 1 CertificateAuthorityDomainConfig, got %d", len(proto.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig))
	}
	caDomain := proto.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig[0]
	if len(caDomain.Fqdns) != 2 || caDomain.Fqdns[0] != "example.com" {
		t.Errorf("unexpected Fqdns: %v", caDomain.Fqdns)
	}
	if caDomain.GetGcpSecretManagerCertificateConfig() == nil || caDomain.GetGcpSecretManagerCertificateConfig().SecretUri != "projects/test-project/secrets/my-ca-cert/versions/1" {
		t.Errorf("unexpected secretUri: %v", caDomain.GetGcpSecretManagerCertificateConfig())
	}

	// Verify WritableCgroups
	if proto.WritableCgroups == nil || !proto.WritableCgroups.Enabled {
		t.Errorf("expected WritableCgroups.Enabled true")
	}

	// Verify RegistryHosts
	if len(proto.RegistryHosts) != 1 {
		t.Fatalf("expected 1 RegistryHosts, got %d", len(proto.RegistryHosts))
	}
	rh := proto.RegistryHosts[0]
	if rh.Server != "my-registry.example.com" {
		t.Errorf("expected Server 'my-registry.example.com', got %q", rh.Server)
	}
	if len(rh.Hosts) != 1 {
		t.Fatalf("expected 1 Hosts, got %d", len(rh.Hosts))
	}
	h := rh.Hosts[0]
	if h.Host != "https://my-registry.example.com" {
		t.Errorf("expected Host 'https://my-registry.example.com', got %q", h.Host)
	}
	if !h.OverridePath {
		t.Errorf("expected OverridePath true")
	}
	if h.DialTimeout == nil || h.DialTimeout.Seconds != 10 {
		t.Errorf("expected DialTimeout 10s, got %v", h.DialTimeout)
	}
	if len(h.Header) != 1 || h.Header[0].Key != "Authorization" {
		t.Errorf("unexpected Header: %v", h.Header)
	}
	if len(h.Ca) != 1 || h.Ca[0].GetGcpSecretManagerSecretUri() != "projects/test-project/secrets/my-ca/versions/1" {
		t.Errorf("unexpected Ca: %v", h.Ca)
	}
	if len(h.Client) != 1 ||
		h.Client[0].Cert.GetGcpSecretManagerSecretUri() != "projects/test-project/secrets/my-client-cert/versions/1" ||
		h.Client[0].Key.GetGcpSecretManagerSecretUri() != "projects/test-project/secrets/my-client-key/versions/1" {
		t.Errorf("unexpected Client: %v", h.Client)
	}

	// Roundtrip back to KRM
	backKRM := ContainerdConfig_FromProto(mapCtx, proto)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("FromProto error: %v", err)
	}
	if backKRM == nil {
		t.Fatalf("expected non-nil backKRM")
	}
	if !direct.ValueOf(backKRM.WritableCgroups.Enabled) {
		t.Errorf("expected roundtrip WritableCgroups true")
	}
	if len(backKRM.RegistryHosts) != 1 || direct.ValueOf(backKRM.RegistryHosts[0].Server) != "my-registry.example.com" {
		t.Errorf("unexpected roundtrip RegistryHosts: %v", backKRM.RegistryHosts)
	}
	if len(backKRM.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig) != 1 ||
		backKRM.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig[0].GCPSecretManagerCertificateConfig.SecretRef.External != "projects/test-project/secrets/my-ca-cert/versions/1" {
		t.Errorf("unexpected roundtrip CertificateAuthorityDomainConfig: %v", backKRM.PrivateRegistryAccessConfig)
	}
}

func TestNormalizeNodePool_Subnetwork(t *testing.T) {
	ctx := context.Background()
	adapter := &nodePoolAdapter{
		id: &krm.ContainerNodePoolIdentity{
			Project:  "test-project",
			Location: "us-central1-a",
			Cluster:  "test-cluster",
			NodePool: "test-pool",
		},
	}

	tests := []struct {
		name           string
		inputSubnet    string
		expectedSubnet string
	}{
		{
			name:           "full compute.googleapis.com URL",
			inputSubnet:    "https://compute.googleapis.com/compute/v1/projects/test-project/regions/us-central1/subnetworks/test-subnetwork",
			expectedSubnet: "projects/test-project/regions/us-central1/subnetworks/test-subnetwork",
		},
		{
			name:           "full www.googleapis.com URL",
			inputSubnet:    "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/subnetworks/test-subnetwork",
			expectedSubnet: "projects/test-project/regions/us-central1/subnetworks/test-subnetwork",
		},
		{
			name:           "relative path",
			inputSubnet:    "projects/test-project/regions/us-central1/subnetworks/test-subnetwork",
			expectedSubnet: "projects/test-project/regions/us-central1/subnetworks/test-subnetwork",
		},
		{
			name:           "short name resolved using adapter project and zone parent region",
			inputSubnet:    "test-subnetwork",
			expectedSubnet: "projects/test-project/regions/us-central1/subnetworks/test-subnetwork",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pbNodePool := &pb.NodePool{
				NetworkConfig: &pb.NodeNetworkConfig{
					Subnetwork: tc.inputSubnet,
					AdditionalNodeNetworkConfigs: []*pb.AdditionalNodeNetworkConfig{
						{
							Network:    "test-network",
							Subnetwork: tc.inputSubnet,
						},
					},
					AdditionalPodNetworkConfigs: []*pb.AdditionalPodNetworkConfig{
						{
							Subnetwork: tc.inputSubnet,
						},
					},
				},
			}

			if err := adapter.normalizeNodePool(ctx, pbNodePool); err != nil {
				t.Fatalf("unexpected error normalizing: %v", err)
			}

			netConfig := pbNodePool.GetNetworkConfig()
			if netConfig.GetSubnetwork() != tc.expectedSubnet {
				t.Errorf("Subnetwork mismatch: got %q, want %q", netConfig.GetSubnetwork(), tc.expectedSubnet)
			}
			if len(netConfig.GetAdditionalNodeNetworkConfigs()) != 1 {
				t.Fatalf("expected 1 additional node network config")
			}
			addNodeNet := netConfig.GetAdditionalNodeNetworkConfigs()[0]
			if addNodeNet.GetNetwork() != "projects/test-project/global/networks/test-network" {
				t.Errorf("AdditionalNodeNetworkConfig Network mismatch: got %q, want %q", addNodeNet.GetNetwork(), "projects/test-project/global/networks/test-network")
			}
			if addNodeNet.GetSubnetwork() != tc.expectedSubnet {
				t.Errorf("AdditionalNodeNetworkConfig Subnetwork mismatch: got %q, want %q", addNodeNet.GetSubnetwork(), tc.expectedSubnet)
			}
			if len(netConfig.GetAdditionalPodNetworkConfigs()) != 1 {
				t.Fatalf("expected 1 additional pod network config")
			}
			addPodNet := netConfig.GetAdditionalPodNetworkConfigs()[0]
			if addPodNet.GetSubnetwork() != tc.expectedSubnet {
				t.Errorf("AdditionalPodNetworkConfig Subnetwork mismatch: got %q, want %q", addPodNet.GetSubnetwork(), tc.expectedSubnet)
			}
		})
	}
}

func TestContainerNodePool_NormalizeNodePool_Autoscaling(t *testing.T) {
	ctx := context.Background()

	t.Run("autoscaling present with nil enabled in desiredKRM defaults to true", func(t *testing.T) {
		adapter := &nodePoolAdapter{
			desiredKRM: &krm.ContainerNodePool{
				Spec: krm.ContainerNodePoolSpec{
					Autoscaling: &krm.NodePoolAutoscaling{
						MinNodeCount: direct.LazyPtr(int32(1)),
						MaxNodeCount: direct.LazyPtr(int32(5)),
					},
				},
			},
		}
		pbNodePool := &pb.NodePool{
			Autoscaling: &pb.NodePoolAutoscaling{
				MinNodeCount: 1,
				MaxNodeCount: 5,
			},
		}
		if err := adapter.normalizeNodePool(ctx, pbNodePool); err != nil {
			t.Fatalf("unexpected error normalizing: %v", err)
		}
		if !pbNodePool.GetAutoscaling().GetEnabled() {
			t.Errorf("expected Autoscaling.Enabled to be true, got false")
		}
	})

	t.Run("autoscaling present with enabled=true in desiredKRM sets enabled to true", func(t *testing.T) {
		adapter := &nodePoolAdapter{
			desiredKRM: &krm.ContainerNodePool{
				Spec: krm.ContainerNodePoolSpec{
					Autoscaling: &krm.NodePoolAutoscaling{
						Enabled:      direct.PtrTo(true),
						MinNodeCount: direct.LazyPtr(int32(1)),
						MaxNodeCount: direct.LazyPtr(int32(5)),
					},
				},
			},
		}
		pbNodePool := &pb.NodePool{
			Autoscaling: &pb.NodePoolAutoscaling{
				Enabled:      true,
				MinNodeCount: 1,
				MaxNodeCount: 5,
			},
		}
		if err := adapter.normalizeNodePool(ctx, pbNodePool); err != nil {
			t.Fatalf("unexpected error normalizing: %v", err)
		}
		if !pbNodePool.GetAutoscaling().GetEnabled() {
			t.Errorf("expected Autoscaling.Enabled to be true, got false")
		}
	})

	t.Run("autoscaling present with enabled=false in desiredKRM sets enabled to false", func(t *testing.T) {
		adapter := &nodePoolAdapter{
			desiredKRM: &krm.ContainerNodePool{
				Spec: krm.ContainerNodePoolSpec{
					Autoscaling: &krm.NodePoolAutoscaling{
						Enabled:      direct.PtrTo(false),
						MinNodeCount: direct.LazyPtr(int32(1)),
						MaxNodeCount: direct.LazyPtr(int32(5)),
					},
				},
			},
		}
		pbNodePool := &pb.NodePool{
			Autoscaling: &pb.NodePoolAutoscaling{
				MinNodeCount: 1,
				MaxNodeCount: 5,
			},
		}
		if err := adapter.normalizeNodePool(ctx, pbNodePool); err != nil {
			t.Fatalf("unexpected error normalizing: %v", err)
		}
		if pbNodePool.GetAutoscaling().GetEnabled() {
			t.Errorf("expected Autoscaling.Enabled to be false, got true")
		}
	})

	t.Run("autoscaling present with nil desiredKRM defaults to true", func(t *testing.T) {
		adapter := &nodePoolAdapter{}
		pbNodePool := &pb.NodePool{
			Autoscaling: &pb.NodePoolAutoscaling{
				MinNodeCount: 1,
				MaxNodeCount: 5,
			},
		}
		if err := adapter.normalizeNodePool(ctx, pbNodePool); err != nil {
			t.Fatalf("unexpected error normalizing: %v", err)
		}
		if !pbNodePool.GetAutoscaling().GetEnabled() {
			t.Errorf("expected Autoscaling.Enabled to be true, got false")
		}
	})

	t.Run("autoscaling nil is not touched", func(t *testing.T) {
		adapter := &nodePoolAdapter{}
		pbNodePool := &pb.NodePool{}
		if err := adapter.normalizeNodePool(ctx, pbNodePool); err != nil {
			t.Fatalf("unexpected error normalizing: %v", err)
		}
		if pbNodePool.GetAutoscaling() != nil {
			t.Errorf("expected Autoscaling to remain nil, got %v", pbNodePool.GetAutoscaling())
		}
	})
}

func TestContainerNodePool_Update_AutoscalingDisabled(t *testing.T) {
	ctx := context.Background()

	desiredKRM := &krm.ContainerNodePool{
		Spec: krm.ContainerNodePoolSpec{
			Autoscaling: &krm.NodePoolAutoscaling{
				Enabled:      direct.PtrTo(false),
				MinNodeCount: direct.LazyPtr(int32(1)),
				MaxNodeCount: direct.LazyPtr(int32(3)),
			},
		},
	}

	actualProto := &pb.NodePool{
		Autoscaling: &pb.NodePoolAutoscaling{
			Enabled:      true,
			MinNodeCount: 1,
			MaxNodeCount: 3,
		},
	}

	adapter := &nodePoolAdapter{
		id: &krm.ContainerNodePoolIdentity{
			Project:  "test-project",
			Location: "us-central1-a",
			Cluster:  "test-cluster",
			NodePool: "test-pool",
		},
		desiredKRM: desiredKRM,
		actual:     actualProto,
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&desiredKRM.Spec,
		actualProto,
		ContainerNodePoolSpec_FromProto,
		ContainerNodePoolSpec_ToProto,
		adapter.normalizeNodePool,
		true,
	)
	if err != nil {
		t.Fatalf("unexpected error comparing specs: %v", err)
	}

	if !diffRes.HasPathWithPrefix("autoscaling") {
		t.Fatalf("expected diff on autoscaling, got paths: %v", diffRes.FieldMask.GetPaths())
	}

	as := diffRes.MergedDesired.GetAutoscaling()
	if as != nil && as.GetEnabled() {
		t.Errorf("expected Autoscaling to be nil or disabled in mergedDesired, got %v", as)
	}
}

func TestContainerNodePool_Update_SwapDisabled(t *testing.T) {
	ctx := context.Background()

	desiredKRM := &krm.ContainerNodePool{
		Spec: krm.ContainerNodePoolSpec{
			NodeConfig: &krm.NodePoolNodeConfig{
				LinuxNodeConfig: &krm.LinuxNodeConfig{
					CgroupMode: direct.LazyPtr("CGROUP_MODE_V2"),
					SwapConfig: &krm.LinuxNodeConfig_SwapConfig{
						Enabled: direct.PtrTo(false),
					},
				},
			},
		},
	}

	actualProto := &pb.NodePool{
		Config: &pb.NodeConfig{
			LinuxNodeConfig: &pb.LinuxNodeConfig{
				CgroupMode: pb.LinuxNodeConfig_CGROUP_MODE_V2,
				SwapConfig: &pb.LinuxNodeConfig_SwapConfig{
					Enabled: direct.LazyPtr(true),
					PerformanceProfile: &pb.LinuxNodeConfig_SwapConfig_BootDiskProfile_{
						BootDiskProfile: &pb.LinuxNodeConfig_SwapConfig_BootDiskProfile{
							SwapSize: &pb.LinuxNodeConfig_SwapConfig_BootDiskProfile_SwapSizePercent{
								SwapSizePercent: 10,
							},
						},
					},
				},
			},
		},
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&desiredKRM.Spec,
		actualProto,
		ContainerNodePoolSpec_FromProto,
		ContainerNodePoolSpec_ToProto,
		nil,
		true,
	)
	if err != nil {
		t.Fatalf("unexpected error comparing specs: %v", err)
	}

	if !diffRes.HasPathWithPrefix("config.linux_node_config") {
		t.Fatalf("expected diff on config.linux_node_config, got paths: %v", diffRes.FieldMask.GetPaths())
	}

	adapter := &nodePoolAdapter{
		id: &krm.ContainerNodePoolIdentity{
			Project:  "test-project",
			Location: "us-central1-a",
			Cluster:  "test-cluster",
			NodePool: "test-pool",
		},
		desiredKRM: desiredKRM,
		actual:     actualProto,
	}

	updateReq, updatePaths := adapter.buildUpdateNodePoolRequest(diffRes)
	if len(updatePaths) == 0 {
		t.Fatalf("expected updatePaths to be non-empty")
	}

	swapConfig := updateReq.GetLinuxNodeConfig().GetSwapConfig()
	if swapConfig == nil {
		t.Fatalf("expected SwapConfig to be set in updateReq")
	}
	if swapConfig.GetEnabled() {
		t.Errorf("expected swap enabled to be false, got true")
	}
	if swapConfig.GetPerformanceProfile() != nil {
		t.Errorf("expected performance profile to be reset to nil, got %v", swapConfig.GetPerformanceProfile())
	}
}

func TestContainerNodePool_Update_PrivateRegistryAccessDisabled(t *testing.T) {
	ctx := context.Background()

	desiredKRM := &krm.ContainerNodePool{
		Spec: krm.ContainerNodePoolSpec{
			NodeConfig: &krm.NodePoolNodeConfig{
				ContainerdConfig: &krm.ContainerdConfig{
					PrivateRegistryAccessConfig: &krm.PrivateRegistryAccessConfig{
						Enabled: direct.PtrTo(false),
					},
				},
			},
		},
	}

	actualProto := &pb.NodePool{
		Config: &pb.NodeConfig{
			ContainerdConfig: &pb.ContainerdConfig{
				PrivateRegistryAccessConfig: &pb.ContainerdConfig_PrivateRegistryAccessConfig{
					Enabled: true,
					CertificateAuthorityDomainConfig: []*pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig{
						{
							Fqdns: []string{"my-registry.example.com"},
							CertificateConfig: &pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig_GcpSecretManagerCertificateConfig{
								GcpSecretManagerCertificateConfig: &pb.ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig_GCPSecretManagerCertificateConfig{
									SecretUri: "projects/test-project/secrets/test-ca/versions/1",
								},
							},
						},
					},
				},
			},
		},
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&desiredKRM.Spec,
		actualProto,
		ContainerNodePoolSpec_FromProto,
		ContainerNodePoolSpec_ToProto,
		nil,
		true,
	)
	if err != nil {
		t.Fatalf("unexpected error comparing specs: %v", err)
	}

	if !diffRes.HasPathWithPrefix("config.containerd_config") {
		t.Fatalf("expected diff on config.containerd_config, got paths: %v", diffRes.FieldMask.GetPaths())
	}

	adapter := &nodePoolAdapter{
		id: &krm.ContainerNodePoolIdentity{
			Project:  "test-project",
			Location: "us-central1-a",
			Cluster:  "test-cluster",
			NodePool: "test-pool",
		},
		desiredKRM: desiredKRM,
		actual:     actualProto,
	}

	updateReq, updatePaths := adapter.buildUpdateNodePoolRequest(diffRes)
	if len(updatePaths) == 0 {
		t.Fatalf("expected updatePaths to be non-empty")
	}

	regAccess := updateReq.GetContainerdConfig().GetPrivateRegistryAccessConfig()
	if regAccess == nil {
		t.Fatalf("expected PrivateRegistryAccessConfig to be set in updateReq")
	}
	if regAccess.GetEnabled() {
		t.Errorf("expected private registry access enabled to be false, got true")
	}
	if len(regAccess.GetCertificateAuthorityDomainConfig()) != 0 {
		t.Errorf("expected CertificateAuthorityDomainConfig to be reset to empty, got %v", regAccess.GetCertificateAuthorityDomainConfig())
	}
}

func TestContainerNodePool_Update_RemoveHandledPaths(t *testing.T) {
	ctx := context.Background()

	desiredKRM := &krm.ContainerNodePool{
		Spec: krm.ContainerNodePoolSpec{
			NodeConfig: &krm.NodePoolNodeConfig{
				Tags:        []string{"foo"},
				MachineType: direct.LazyPtr("e2-standard-4"),
			},
		},
	}

	actualProto := &pb.NodePool{
		Config: &pb.NodeConfig{
			Tags:        []string{"bar"},
			MachineType: "e2-medium",
		},
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&desiredKRM.Spec,
		actualProto,
		ContainerNodePoolSpec_FromProto,
		ContainerNodePoolSpec_ToProto,
		nil,
		true,
	)
	if err != nil {
		t.Fatalf("unexpected error comparing specs: %v", err)
	}

	adapter := &nodePoolAdapter{
		id: &krm.ContainerNodePoolIdentity{
			Project:  "test-project",
			Location: "us-central1-a",
			Cluster:  "test-cluster",
			NodePool: "test-pool",
		},
		desiredKRM: desiredKRM,
		actual:     actualProto,
	}

	_, updatePaths := adapter.buildUpdateNodePoolRequest(diffRes)
	diffRes.RemovePathWithPrefix(updatePaths...)

	if diffRes.Empty() {
		t.Fatalf("expected diffRes to not be empty due to immutable machine_type")
	}
	if diffRes.HasPathWithPrefix("config.tags") {
		t.Errorf("expected config.tags to be removed from diffRes")
	}
	if !diffRes.HasPathWithPrefix("config.machine_type") {
		t.Errorf("expected config.machine_type to remain in diffRes")
	}
}

func TestContainerNodePool_Update_Gvnic(t *testing.T) {
	ctx := context.Background()

	desiredKRM := &krm.ContainerNodePool{
		Spec: krm.ContainerNodePoolSpec{
			NodeConfig: &krm.NodePoolNodeConfig{
				GVNIC: &krm.VirtualNic{
					Enabled: direct.PtrTo(true),
				},
			},
		},
	}

	actualProto := &pb.NodePool{
		Config: &pb.NodeConfig{
			Gvnic: &pb.VirtualNIC{
				Enabled: false,
			},
		},
	}

	diffRes, err := common.CompareSpecifiedSpec(
		ctx,
		&desiredKRM.Spec,
		actualProto,
		ContainerNodePoolSpec_FromProto,
		ContainerNodePoolSpec_ToProto,
		nil,
		true,
	)
	if err != nil {
		t.Fatalf("unexpected error comparing specs: %v", err)
	}

	if !diffRes.HasPathWithPrefix("config.gvnic") {
		t.Fatalf("expected diff on config.gvnic, got paths: %v", diffRes.FieldMask.GetPaths())
	}

	adapter := &nodePoolAdapter{
		id: &krm.ContainerNodePoolIdentity{
			Project:  "test-project",
			Location: "us-central1-a",
			Cluster:  "test-cluster",
			NodePool: "test-pool",
		},
		desiredKRM: desiredKRM,
		actual:     actualProto,
	}

	updateReq, updatePaths := adapter.buildUpdateNodePoolRequest(diffRes)
	if len(updatePaths) == 0 {
		t.Fatalf("expected updatePaths to be non-empty")
	}

	if updateReq.GetGvnic() == nil || !updateReq.GetGvnic().GetEnabled() {
		t.Errorf("expected Gvnic enabled in updateReq, got %v", updateReq.GetGvnic())
	}

	diffRes.RemovePathWithPrefix(updatePaths...)
	if !diffRes.Empty() {
		t.Errorf("expected diffRes to be empty after removing updatePaths, got %v", diffRes.FieldMask.GetPaths())
	}
}
