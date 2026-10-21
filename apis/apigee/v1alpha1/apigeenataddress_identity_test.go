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

package v1alpha1

import (
	"context"
	"fmt"
	"testing"

	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type mockReader struct {
	client.Reader
	objects map[client.ObjectKey]client.Object
}

func (m *mockReader) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	found, ok := m.objects[key]
	if !ok {
		return fmt.Errorf("not found: %v", key)
	}
	if u, ok := obj.(*unstructured.Unstructured); ok {
		if fu, ok := found.(*unstructured.Unstructured); ok {
			u.Object = fu.Object
			return nil
		}
	}
	return nil
}

func TestApigeeNATAddressIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *ApigeeNATAddressIdentity
		wantParent string
		wantID     string
		wantHost   string
		wantString string
	}{
		{
			name: "valid standard external reference",
			ref:  "organizations/my-org/instances/my-instance/natAddresses/my-nat",
			want: &ApigeeNATAddressIdentity{
				Organization: "my-org",
				Instance:     "my-instance",
				NatAddress:   "my-nat",
			},
			wantParent: "organizations/my-org/instances/my-instance",
			wantID:     "my-nat",
			wantHost:   "apigee.googleapis.com",
			wantString: "organizations/my-org/instances/my-instance/natAddresses/my-nat",
		},
		{
			name: "valid reference with double slash host prefix",
			ref:  "//apigee.googleapis.com/organizations/my-org/instances/my-instance/natAddresses/my-nat",
			want: &ApigeeNATAddressIdentity{
				Organization: "my-org",
				Instance:     "my-instance",
				NatAddress:   "my-nat",
			},
			wantParent: "organizations/my-org/instances/my-instance",
			wantID:     "my-nat",
			wantHost:   "apigee.googleapis.com",
			wantString: "organizations/my-org/instances/my-instance/natAddresses/my-nat",
		},
		{
			name: "valid reference with https scheme and v1 api version",
			ref:  "https://apigee.googleapis.com/v1/organizations/my-org/instances/my-instance/natAddresses/my-nat",
			want: &ApigeeNATAddressIdentity{
				Organization: "my-org",
				Instance:     "my-instance",
				NatAddress:   "my-nat",
			},
			wantParent: "organizations/my-org/instances/my-instance",
			wantID:     "my-nat",
			wantHost:   "apigee.googleapis.com",
			wantString: "organizations/my-org/instances/my-instance/natAddresses/my-nat",
		},
		{
			name: "valid reference with https scheme without v1 api version",
			ref:  "https://apigee.googleapis.com/organizations/my-org/instances/my-instance/natAddresses/my-nat",
			want: &ApigeeNATAddressIdentity{
				Organization: "my-org",
				Instance:     "my-instance",
				NatAddress:   "my-nat",
			},
			wantParent: "organizations/my-org/instances/my-instance",
			wantID:     "my-nat",
			wantHost:   "apigee.googleapis.com",
			wantString: "organizations/my-org/instances/my-instance/natAddresses/my-nat",
		},
		{
			name:    "invalid format empty",
			ref:     "",
			wantErr: true,
		},
		{
			name:    "invalid format arbitrary path",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "invalid format missing natAddress",
			ref:     "organizations/my-org/instances/my-instance",
			wantErr: true,
		},
		{
			name:    "invalid format wrong collection segment",
			ref:     "organizations/my-org/instances/my-instance/wrongSegments/my-nat",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &ApigeeNATAddressIdentity{}
			err := id.FromExternal(tc.ref)
			if (err != nil) != tc.wantErr {
				t.Fatalf("FromExternal(%q) error = %v, wantErr %v", tc.ref, err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if diff := cmp.Diff(tc.want, id); diff != "" {
				t.Errorf("FromExternal(%q) mismatch (-want +got):\n%s", tc.ref, diff)
			}
			if id.ParentString() != tc.wantParent {
				t.Errorf("ParentString() = %q, want %q", id.ParentString(), tc.wantParent)
			}
			if id.ID() != tc.wantID {
				t.Errorf("ID() = %q, want %q", id.ID(), tc.wantID)
			}
			if id.Host() != tc.wantHost {
				t.Errorf("Host() = %q, want %q", id.Host(), tc.wantHost)
			}
			if id.String() != tc.wantString {
				t.Errorf("String() = %q, want %q", id.String(), tc.wantString)
			}
		})
	}
}

func TestApigeeNATAddress_GetIdentity(t *testing.T) {
	ctx := context.Background()

	instanceUnstructured := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apigee.cnrm.cloud.google.com/v1beta1",
			"kind":       "ApigeeInstance",
			"metadata": map[string]interface{}{
				"name":      "krm-instance",
				"namespace": "test-ns",
			},
			"status": map[string]interface{}{
				"externalRef": "organizations/org-from-krm/instances/inst-from-krm",
			},
		},
	}
	reader := &mockReader{
		objects: map[client.ObjectKey]client.Object{
			{Namespace: "test-ns", Name: "krm-instance"}: instanceUnstructured,
		},
	}

	resourceID := "custom-nat"
	tests := []struct {
		name       string
		obj        *ApigeeNATAddress
		reader     client.Reader
		wantErr    bool
		wantString string
	}{
		{
			name: "valid spec with explicit resourceID and instanceId path",
			obj: &ApigeeNATAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "krm-name",
					Namespace: "test-ns",
				},
				Spec: ApigeeNATAddressSpec{
					InstanceId: "organizations/my-org/instances/my-inst",
					ResourceID: &resourceID,
				},
			},
			wantString: "organizations/my-org/instances/my-inst/natAddresses/custom-nat",
		},
		{
			name: "valid spec defaulting to metadata.name",
			obj: &ApigeeNATAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "default-nat",
					Namespace: "test-ns",
				},
				Spec: ApigeeNATAddressSpec{
					InstanceId: "organizations/my-org/instances/my-inst",
				},
			},
			wantString: "organizations/my-org/instances/my-inst/natAddresses/default-nat",
		},
		{
			name: "valid spec with instanceId having full URL prefix",
			obj: &ApigeeNATAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "default-nat",
					Namespace: "test-ns",
				},
				Spec: ApigeeNATAddressSpec{
					InstanceId: "//apigee.googleapis.com/organizations/my-org/instances/my-inst",
				},
			},
			wantString: "organizations/my-org/instances/my-inst/natAddresses/default-nat",
		},
		{
			name: "valid spec with instanceId without organizations/ prefix",
			obj: &ApigeeNATAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "default-nat",
					Namespace: "test-ns",
				},
				Spec: ApigeeNATAddressSpec{
					InstanceId: "my-org/instances/my-inst",
				},
			},
			wantString: "organizations/my-org/instances/my-inst/natAddresses/default-nat",
		},
		{
			name: "valid spec resolving instance from local KRM ApigeeInstance",
			obj: &ApigeeNATAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "default-nat",
					Namespace: "test-ns",
				},
				Spec: ApigeeNATAddressSpec{
					InstanceId: "krm-instance",
				},
			},
			reader:     reader,
			wantString: "organizations/org-from-krm/instances/inst-from-krm/natAddresses/default-nat",
		},
		{
			name: "error resolving instance from local KRM ApigeeInstance without reader",
			obj: &ApigeeNATAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "default-nat",
					Namespace: "test-ns",
				},
				Spec: ApigeeNATAddressSpec{
					InstanceId: "krm-instance",
				},
			},
			reader:  nil,
			wantErr: true,
		},
		{
			name: "missing instanceId",
			obj: &ApigeeNATAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "default-nat",
					Namespace: "test-ns",
				},
				Spec: ApigeeNATAddressSpec{},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.reader
			id, err := tc.obj.GetIdentity(ctx, r)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if id.String() != tc.wantString {
				t.Errorf("GetIdentity().String() = %q, want %q", id.String(), tc.wantString)
			}
		})
	}
}

func TestApigeeNATAddressRef_Normalize(t *testing.T) {
	ctx := context.Background()

	readyObj := &ApigeeNATAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-nat",
			Namespace: "test-ns",
		},
		Spec: ApigeeNATAddressSpec{
			InstanceId: "organizations/my-org/instances/my-inst",
		},
		Status: ApigeeNATAddressStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}

	notReadyObj := &ApigeeNATAddress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready-nat",
			Namespace: "test-ns",
		},
		Spec: ApigeeNATAddressSpec{
			InstanceId: "organizations/my-org/instances/my-inst",
		},
		Status: ApigeeNATAddressStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "False",
				},
			},
		},
	}

	uReady := &unstructured.Unstructured{}
	uReady.SetGroupVersionKind(ApigeeNATAddressGVK)
	uReady.SetName("my-nat")
	uReady.SetNamespace("test-ns")
	uReady.Object["spec"] = map[string]interface{}{
		"instanceId": "organizations/my-org/instances/my-inst",
	}
	uReady.Object["status"] = map[string]interface{}{
		"conditions": []interface{}{
			map[string]interface{}{
				"type":   "Ready",
				"status": "True",
			},
		},
	}

	uNotReady := &unstructured.Unstructured{}
	uNotReady.SetGroupVersionKind(ApigeeNATAddressGVK)
	uNotReady.SetName("not-ready-nat")
	uNotReady.SetNamespace("test-ns")
	uNotReady.Object["spec"] = map[string]interface{}{
		"instanceId": "organizations/my-org/instances/my-inst",
	}
	uNotReady.Object["status"] = map[string]interface{}{
		"conditions": []interface{}{
			map[string]interface{}{
				"type":   "Ready",
				"status": "False",
			},
		},
	}

	_ = readyObj
	_ = notReadyObj

	reader := &mockReader{
		objects: map[client.ObjectKey]client.Object{
			{Namespace: "test-ns", Name: "my-nat"}:        uReady,
			{Namespace: "test-ns", Name: "not-ready-nat"}: uNotReady,
		},
	}

	t.Run("external already set", func(t *testing.T) {
		ref := &ApigeeNATAddressRef{
			External: "organizations/my-org/instances/my-inst/natAddresses/my-nat",
		}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() error = %v", err)
		}
		if ref.External != "organizations/my-org/instances/my-inst/natAddresses/my-nat" {
			t.Errorf("Normalize() ref.External = %q, want %q", ref.External, "organizations/my-org/instances/my-inst/natAddresses/my-nat")
		}
	})

	t.Run("normalize ready resource via fallback", func(t *testing.T) {
		ref := &ApigeeNATAddressRef{
			Name: "my-nat",
		}
		if err := ref.Normalize(ctx, reader, "test-ns"); err != nil {
			t.Fatalf("Normalize() error = %v", err)
		}
		want := "organizations/my-org/instances/my-inst/natAddresses/my-nat"
		if ref.External != want {
			t.Errorf("Normalize() ref.External = %q, want %q", ref.External, want)
		}
	})

	t.Run("normalize not ready resource via fallback", func(t *testing.T) {
		ref := &ApigeeNATAddressRef{
			Name: "not-ready-nat",
		}
		if err := ref.Normalize(ctx, reader, "test-ns"); err == nil {
			t.Fatalf("Normalize() expected error for not ready resource, got nil")
		}
	})
}
