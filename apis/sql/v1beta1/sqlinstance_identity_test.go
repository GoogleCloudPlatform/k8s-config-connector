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

package v1beta1

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSQLInstanceIdentityFromExternal(t *testing.T) {
	tests := []struct {
		name         string
		external     string
		wantProject  string
		wantInstance string
		wantErr      bool
	}{
		{
			name:         "full URI",
			external:     "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
			wantProject:  "my-project",
			wantInstance: "my-instance",
		},
		{
			name:         "relative resource name",
			external:     "projects/my-project/instances/my-instance",
			wantProject:  "my-project",
			wantInstance: "my-instance",
		},
		{
			name:     "empty string",
			external: "",
			wantErr:  true,
		},
		{
			name:     "invalid string without format",
			external: "invalid",
			wantErr:  true,
		},
		{
			name:     "missing instance name",
			external: "projects/my-project/instances/",
			wantErr:  true,
		},
		{
			name:     "wrong resource type",
			external: "projects/my-project/databases/my-database",
			wantErr:  true,
		},
		{
			name:     "invalid location-qualified resource name",
			external: "projects/my-project/locations/us-central1/instances/my-instance",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := &SQLInstanceIdentity{}
			err := identity.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if identity.Project != tt.wantProject {
				t.Errorf("FromExternal() project = %q, want %q", identity.Project, tt.wantProject)
			}
			if identity.Instance != tt.wantInstance {
				t.Errorf("FromExternal() instance = %q, want %q", identity.Instance, tt.wantInstance)
			}
		})
	}
}

func TestSQLInstanceIdentityString(t *testing.T) {
	identity := &SQLInstanceIdentity{
		Project:  "my-project",
		Instance: "my-instance",
	}

	want := "projects/my-project/instances/my-instance"
	if got := identity.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}

	wantParent := "projects/my-project"
	if got := identity.ParentString(); got != wantParent {
		t.Errorf("ParentString() = %q, want %q", got, wantParent)
	}

	// Test nil safety
	var nilIdentity *SQLInstanceIdentity
	if got := nilIdentity.String(); got != "" {
		t.Errorf("nil.String() = %q, want empty string", got)
	}
	if got := nilIdentity.ParentString(); got != "" {
		t.Errorf("nil.ParentString() = %q, want empty string", got)
	}
}

func TestSQLInstanceRef_Normalize(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() error = %v", err)
	}

	sqlInstance := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": SQLInstanceGVK.GroupVersion().String(),
			"kind":       SQLInstanceGVK.Kind,
			"metadata": map[string]interface{}{
				"name":      "my-instance",
				"namespace": "test-ns",
				"annotations": map[string]interface{}{
					"cnrm.cloud.google.com/project-id": "my-project",
				},
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "True",
					},
				},
			},
		},
	}

	sqlInstanceWithResourceID := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": SQLInstanceGVK.GroupVersion().String(),
			"kind":       SQLInstanceGVK.Kind,
			"metadata": map[string]interface{}{
				"name":      "k8s-name",
				"namespace": "test-ns",
				"annotations": map[string]interface{}{
					"cnrm.cloud.google.com/project-id": "my-project",
				},
			},
			"spec": map[string]interface{}{
				"resourceID": "custom-instance-id",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "True",
					},
				},
			},
		},
	}

	sqlInstanceNotReady := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": SQLInstanceGVK.GroupVersion().String(),
			"kind":       SQLInstanceGVK.Kind,
			"metadata": map[string]interface{}{
				"name":      "not-ready-instance",
				"namespace": "test-ns",
				"annotations": map[string]interface{}{
					"cnrm.cloud.google.com/project-id": "my-project",
				},
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "False",
					},
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(sqlInstance, sqlInstanceWithResourceID, sqlInstanceNotReady).Build()

	tests := []struct {
		name         string
		ref          *SQLInstanceRef
		defaultNS    string
		wantExternal string
		wantErr      bool
	}{
		{
			name: "full external URI",
			ref: &SQLInstanceRef{
				External: "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
			},
			wantExternal: "//sqladmin.googleapis.com/projects/my-project/instances/my-instance",
		},
		{
			name: "relative external URI",
			ref: &SQLInstanceRef{
				External: "projects/my-project/instances/my-instance",
			},
			wantExternal: "projects/my-project/instances/my-instance",
		},
		{
			name: "resolve by name with default namespace",
			ref: &SQLInstanceRef{
				Name: "my-instance",
			},
			defaultNS:    "test-ns",
			wantExternal: "projects/my-project/instances/my-instance",
		},
		{
			name: "resolve by name with explicit namespace",
			ref: &SQLInstanceRef{
				Name:      "my-instance",
				Namespace: "test-ns",
			},
			wantExternal: "projects/my-project/instances/my-instance",
		},
		{
			name: "resolve by name with spec.resourceID taking precedence",
			ref: &SQLInstanceRef{
				Name: "k8s-name",
			},
			defaultNS:    "test-ns",
			wantExternal: "projects/my-project/instances/custom-instance-id",
		},
		{
			name: "resource not ready",
			ref: &SQLInstanceRef{
				Name: "not-ready-instance",
			},
			defaultNS: "test-ns",
			wantErr:   true,
		},
		{
			name: "resource not found",
			ref: &SQLInstanceRef{
				Name: "non-existent",
			},
			defaultNS: "test-ns",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(ctx, fakeClient, tt.defaultNS)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if tt.ref.External != tt.wantExternal {
				t.Errorf("Normalize() external = %q, want %q", tt.ref.External, tt.wantExternal)
			}
		})
	}
}

func TestSQLInstance_GetIdentity(t *testing.T) {
	ctx := context.Background()

	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() error = %v", err)
	}

	fakeClient := fake.NewClientBuilder().
		WithScheme(scheme).
		Build()

	var _ identity.Resource = &SQLInstance{}

	tests := []struct {
		name         string
		obj          *SQLInstance
		wantIdentity string
		wantErr      bool
	}{
		{
			name: "from metadata.name and annotation",
			obj: &SQLInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-instance",
					Namespace: "test-ns",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
			},
			wantIdentity: "projects/my-project/instances/test-instance",
		},
		{
			name: "from spec.resourceID and annotation",
			obj: &SQLInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-name",
					Namespace: "test-ns",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "custom-project",
					},
				},
				Spec: SQLInstanceSpec{
					ResourceID: ptr.To("custom-instance-id"),
				},
			},
			wantIdentity: "projects/custom-project/instances/custom-instance-id",
		},
		{
			name: "missing project",
			obj: &SQLInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-instance",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if id.String() != tt.wantIdentity {
				t.Errorf("GetIdentity() = %q, want %q", id.String(), tt.wantIdentity)
			}
			if idV2, ok := id.(identity.IdentityV2); !ok {
				t.Errorf("expected IdentityV2")
			} else if idV2.Host() != "sqladmin.googleapis.com" {
				t.Errorf("Host() = %q, want sqladmin.googleapis.com", idV2.Host())
			}
		})
	}
}
