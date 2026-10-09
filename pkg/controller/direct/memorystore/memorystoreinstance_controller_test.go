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

package memorystore

import (
	"context"
	"testing"

	memorystorepb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestIsConnectionSubset(t *testing.T) {
	tests := []struct {
		name       string
		desired    []*memorystorepb.Instance_ConnectionDetail
		actual     []*memorystorepb.Instance_ConnectionDetail
		wantSubset bool
	}{
		{
			name:       "empty desired is subset of empty actual",
			desired:    nil,
			actual:     nil,
			wantSubset: true,
		},
		{
			name:    "empty desired is subset of non-empty actual",
			desired: nil,
			actual: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
			},
			wantSubset: true,
		},
		{
			name: "desired output-only fields omitted matches actual with output-only fields populated",
			desired: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
			},
			actual: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId:         "proj-1",
							Network:           "net-1",
							PscConnectionId:   "conn-123",
							IpAddress:         "10.0.0.1",
							ForwardingRule:    "fwd-1",
							ServiceAttachment: "svc-1",
							Ports:             &memorystorepb.PscAutoConnection_Port{Port: 6379},
						},
					},
				},
			},
			wantSubset: true,
		},
		{
			name: "desired subset of actual with extra server-generated connections",
			desired: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
			},
			actual: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "server-internal",
							Network:   "server-net",
						},
					},
				},
			},
			wantSubset: true,
		},
		{
			name: "desired has connection missing in actual",
			desired: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-2",
							Network:   "net-2",
						},
					},
				},
			},
			actual: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
			},
			wantSubset: false,
		},
		{
			name: "duplicate desired connections matching single actual connection succeeds",
			desired: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
			},
			actual: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
			},
			wantSubset: true,
		},
		{
			name: "duplicate desired connections match multiple actual connections",
			desired: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId: "proj-1",
							Network:   "net-1",
						},
					},
				},
			},
			actual: []*memorystorepb.Instance_ConnectionDetail{
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId:       "proj-1",
							Network:         "net-1",
							PscConnectionId: "conn-1",
						},
					},
				},
				{
					Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
						PscAutoConnection: &memorystorepb.PscAutoConnection{
							ProjectId:       "proj-1",
							Network:         "net-1",
							PscConnectionId: "conn-2",
						},
					},
				},
			},
			wantSubset: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isConnectionSubset(tc.desired, tc.actual)
			if got != tc.wantSubset {
				t.Errorf("isConnectionSubset() = %v, want %v", got, tc.wantSubset)
			}
		})
	}
}

func TestAlignEndpointsWithDesired(t *testing.T) {
	desiredEndpoint := &memorystorepb.Instance_InstanceEndpoint{
		Connections: []*memorystorepb.Instance_ConnectionDetail{
			{
				Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
					PscAutoConnection: &memorystorepb.PscAutoConnection{
						ProjectId: "proj-1",
						Network:   "net-1",
					},
				},
			},
		},
	}

	actualEndpointMatched := &memorystorepb.Instance_InstanceEndpoint{
		Connections: []*memorystorepb.Instance_ConnectionDetail{
			{
				Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
					PscAutoConnection: &memorystorepb.PscAutoConnection{
						ProjectId:         "proj-1",
						Network:           "net-1",
						PscConnectionId:   "conn-123",
						IpAddress:         "10.0.0.1",
						ForwardingRule:    "fwd-1",
						ServiceAttachment: "svc-1",
						Ports:             &memorystorepb.PscAutoConnection_Port{Port: 6379},
					},
				},
			},
			{
				Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
					PscAutoConnection: &memorystorepb.PscAutoConnection{
						ProjectId: "internal-tenant",
						Network:   "internal-net",
					},
				},
			},
		},
	}

	actualEndpointExtra := &memorystorepb.Instance_InstanceEndpoint{
		Connections: []*memorystorepb.Instance_ConnectionDetail{
			{
				Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
					PscAutoConnection: &memorystorepb.PscAutoConnection{
						ProjectId: "replica-proj",
						Network:   "replica-net",
					},
				},
			},
		},
	}

	tests := []struct {
		name          string
		actual        *memorystorepb.Instance
		desired       *memorystorepb.Instance
		wantEndpoints []*memorystorepb.Instance_InstanceEndpoint
	}{
		{
			name: "empty desired endpoints sets actual endpoints to nil",
			actual: &memorystorepb.Instance{
				Endpoints: []*memorystorepb.Instance_InstanceEndpoint{actualEndpointExtra},
			},
			desired: &memorystorepb.Instance{
				Endpoints: nil,
			},
			wantEndpoints: nil,
		},
		{
			name: "desired subset of actual aligns matched endpoint and filters unmanaged endpoints",
			actual: &memorystorepb.Instance{
				Endpoints: []*memorystorepb.Instance_InstanceEndpoint{
					actualEndpointMatched,
					actualEndpointExtra,
				},
			},
			desired: &memorystorepb.Instance{
				Endpoints: []*memorystorepb.Instance_InstanceEndpoint{desiredEndpoint},
			},
			wantEndpoints: []*memorystorepb.Instance_InstanceEndpoint{
				{
					Connections: desiredEndpoint.Connections,
				},
			},
		},
		{
			name: "desired not a subset leaves actual endpoints untouched",
			actual: &memorystorepb.Instance{
				Endpoints: []*memorystorepb.Instance_InstanceEndpoint{actualEndpointExtra},
			},
			desired: &memorystorepb.Instance{
				Endpoints: []*memorystorepb.Instance_InstanceEndpoint{desiredEndpoint},
			},
			wantEndpoints: []*memorystorepb.Instance_InstanceEndpoint{actualEndpointExtra},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			alignEndpointsWithDesired(tc.actual, tc.desired)
			if diff := cmp.Diff(tc.wantEndpoints, tc.actual.Endpoints, protocmp.Transform()); diff != "" {
				t.Errorf("alignEndpointsWithDesired() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCompareInstanceDuplicateConnectionsNoDrift(t *testing.T) {
	ctx := context.Background()
	desired := &memorystorepb.Instance{
		Endpoints: []*memorystorepb.Instance_InstanceEndpoint{
			{
				Connections: []*memorystorepb.Instance_ConnectionDetail{
					{
						Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
							PscAutoConnection: &memorystorepb.PscAutoConnection{
								ProjectId: "proj-1",
								Network:   "net-1",
							},
						},
					},
					{
						Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
							PscAutoConnection: &memorystorepb.PscAutoConnection{
								ProjectId: "proj-1",
								Network:   "net-1",
							},
						},
					},
				},
			},
		},
	}

	actual := &memorystorepb.Instance{
		Endpoints: []*memorystorepb.Instance_InstanceEndpoint{
			{
				Connections: []*memorystorepb.Instance_ConnectionDetail{
					{
						Connection: &memorystorepb.Instance_ConnectionDetail_PscAutoConnection{
							PscAutoConnection: &memorystorepb.PscAutoConnection{
								ProjectId:         "proj-1",
								Network:           "net-1",
								PscConnectionId:   "server-conn-1",
								IpAddress:         "10.0.0.1",
								ForwardingRule:    "fwd-1",
								ServiceAttachment: "svc-1",
								Ports:             &memorystorepb.PscAutoConnection_Port{Port: 6379},
							},
						},
					},
				},
			},
		},
	}

	diffs, err := compareInstance(ctx, actual, desired)
	if err != nil {
		t.Fatalf("compareInstance() unexpected error: %v", err)
	}
	if diffs.HasDiff() {
		t.Errorf("compareInstance() reported diffs %v, expected 0 diffs when GCP returns 1 connection for duplicate desired connections", diffs.Fields)
	}
}

func TestMemorystoreInstance_NormalizeReferences(t *testing.T) {
	ctx := context.Background()

	readyKey := &unstructured.Unstructured{}
	readyKey.SetGroupVersionKind(kmsv1beta1.KMSCryptoKeyGVK)
	readyKey.SetName("test-key")
	readyKey.SetNamespace("test-ns")
	selfLink := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/test-key"
	if err := unstructured.SetNestedField(readyKey.Object, selfLink, "status", "selfLink"); err != nil {
		t.Fatalf("failed to set status.selfLink: %v", err)
	}

	scheme := runtime.NewScheme()
	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(readyKey).Build()

	t.Run("resolves KmsKeyRef by name", func(t *testing.T) {
		obj := &krm.MemorystoreInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.MemorystoreInstanceSpec{
				KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					Name: "test-key",
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.KmsKeyRef.External != selfLink {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.KmsKeyRef.External, selfLink)
		}

		mapCtx := &direct.MapContext{}
		protoObj := MemorystoreInstanceSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("MemorystoreInstanceSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetKmsKey() != selfLink {
			t.Errorf("protoObj KmsKey = %q, want %q", protoObj.GetKmsKey(), selfLink)
		}

		fromProto := MemorystoreInstanceSpec_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("MemorystoreInstanceSpec_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KmsKeyRef == nil || fromProto.KmsKeyRef.External != selfLink {
			t.Errorf("fromProto KmsKeyRef.External = %v, want %q", fromProto.KmsKeyRef, selfLink)
		}
	})

	t.Run("handles external KmsKeyRef directly", func(t *testing.T) {
		externalKey := "projects/test-project/locations/us-central1/keyRings/test-keyring/cryptoKeys/external-key"
		obj := &krm.MemorystoreInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.MemorystoreInstanceSpec{
				KmsKeyRef: &kmsv1beta1.KMSCryptoKeyRef{
					External: externalKey,
				},
			},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		if obj.Spec.KmsKeyRef.External != externalKey {
			t.Errorf("NormalizeReferences() External = %q, want %q", obj.Spec.KmsKeyRef.External, externalKey)
		}

		mapCtx := &direct.MapContext{}
		protoObj := MemorystoreInstanceSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("MemorystoreInstanceSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetKmsKey() != externalKey {
			t.Errorf("protoObj KmsKey = %q, want %q", protoObj.GetKmsKey(), externalKey)
		}

		fromProto := MemorystoreInstanceSpec_FromProto(mapCtx, protoObj)
		if mapCtx.Err() != nil {
			t.Fatalf("MemorystoreInstanceSpec_FromProto unexpected error: %v", mapCtx.Err())
		}
		if fromProto.KmsKeyRef == nil || fromProto.KmsKeyRef.External != externalKey {
			t.Errorf("fromProto KmsKeyRef.External = %v, want %q", fromProto.KmsKeyRef, externalKey)
		}
	})

	t.Run("handles nil KmsKeyRef", func(t *testing.T) {
		obj := &krm.MemorystoreInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-instance",
				Namespace: "test-ns",
			},
			Spec: krm.MemorystoreInstanceSpec{},
		}

		if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
			t.Fatalf("NormalizeReferences() unexpected error: %v", err)
		}

		mapCtx := &direct.MapContext{}
		protoObj := MemorystoreInstanceSpec_ToProto(mapCtx, &obj.Spec)
		if mapCtx.Err() != nil {
			t.Fatalf("MemorystoreInstanceSpec_ToProto unexpected error: %v", mapCtx.Err())
		}
		if protoObj.GetKmsKey() != "" {
			t.Errorf("protoObj KmsKey = %q, want empty", protoObj.GetKmsKey())
		}
	})
}
