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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDialogflowCXEntityTypeIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowCXEntityTypeIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/agents/my-agent/entityTypes/my-entity-type",
			want: &DialogflowCXEntityTypeIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Agent:      "my-agent",
				EntityType: "my-entity-type",
			},
		},
		{
			name: "full url with scheme",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/entityTypes/my-entity-type",
			want: &DialogflowCXEntityTypeIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Agent:      "my-agent",
				EntityType: "my-entity-type",
			},
		},
		{
			name: "full url without scheme",
			ref:  "//dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/entityTypes/my-entity-type",
			want: &DialogflowCXEntityTypeIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Agent:      "my-agent",
				EntityType: "my-entity-type",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "empty reference",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DialogflowCXEntityTypeIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != "projects/my-project/locations/us-central1/agents/my-agent/entityTypes/my-entity-type" {
					t.Errorf("String() = %q, want %q", got, "projects/my-project/locations/us-central1/agents/my-agent/entityTypes/my-entity-type")
				}
			}
		})
	}
}

func TestDialogflowCXEntityTypeIdentity_Methods(t *testing.T) {
	id := &DialogflowCXEntityTypeIdentity{
		Project:    "my-project",
		Location:   "us-central1",
		Agent:      "my-agent",
		EntityType: "my-entity-type",
	}

	if got, want := id.ParentString(), "projects/my-project/locations/us-central1/agents/my-agent"; got != want {
		t.Errorf("ParentString() = %q, want %q", got, want)
	}

	if got, want := id.Host(), "dialogflow.googleapis.com"; got != want {
		t.Errorf("Host() = %q, want %q", got, want)
	}

	if got, want := id.String(), "projects/my-project/locations/us-central1/agents/my-agent/entityTypes/my-entity-type"; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}

	if !id.HasIdentitySpecified() {
		t.Errorf("HasIdentitySpecified() = false, want true")
	}

	emptyID := &DialogflowCXEntityTypeIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Agent:    "my-agent",
	}
	if emptyID.HasIdentitySpecified() {
		t.Errorf("emptyID.HasIdentitySpecified() = true, want false")
	}
}

func TestGetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *DialogflowCXEntityType
		want    *DialogflowCXEntityTypeIdentity
		wantErr bool
	}{
		{
			name: "spec with resourceID",
			obj: &DialogflowCXEntityType{
				Spec: DialogflowCXEntityTypeSpec{
					Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
					ResourceID: common.LazyPtr("custom-id"),
				},
			},
			want: &DialogflowCXEntityTypeIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Agent:      "my-agent",
				EntityType: "custom-id",
			},
		},
		{
			name: "spec without resourceID (server-generated identity pending)",
			obj: &DialogflowCXEntityType{
				Spec: DialogflowCXEntityTypeSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
				},
			},
			want: &DialogflowCXEntityTypeIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Agent:      "my-agent",
				EntityType: "",
			},
		},
		{
			name: "spec without resourceID, resolved from status.externalRef",
			obj: &DialogflowCXEntityType{
				Spec: DialogflowCXEntityTypeSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
				},
				Status: DialogflowCXEntityTypeStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/entityTypes/server-id-1"),
				},
			},
			want: &DialogflowCXEntityTypeIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Agent:      "my-agent",
				EntityType: "server-id-1",
			},
		},
		{
			name: "spec without resourceID, resolved from status.name",
			obj: &DialogflowCXEntityType{
				Spec: DialogflowCXEntityTypeSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
				},
				Status: DialogflowCXEntityTypeStatus{
					Name: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/entityTypes/server-id-2"),
				},
			},
			want: &DialogflowCXEntityTypeIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Agent:      "my-agent",
				EntityType: "server-id-2",
			},
		},
		{
			name: "drift detection error when spec resourceID does not match status",
			obj: &DialogflowCXEntityType{
				Spec: DialogflowCXEntityTypeSpec{
					Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
					ResourceID: common.LazyPtr("spec-id"),
				},
				Status: DialogflowCXEntityTypeStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/entityTypes/different-id"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing parent",
			obj: &DialogflowCXEntityType{
				Spec: DialogflowCXEntityTypeSpec{
					ResourceID: common.LazyPtr("custom-id"),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid parent format",
			obj: &DialogflowCXEntityType{
				Spec: DialogflowCXEntityTypeSpec{
					Parent:     common.LazyPtr("invalid/parent"),
					ResourceID: common.LazyPtr("custom-id"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				gotID, ok := got.(*DialogflowCXEntityTypeIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned %T, want *DialogflowCXEntityTypeIdentity", got)
				}
				if diff := cmp.Diff(tt.want, gotID); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestDialogflowCXEntityTypeRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     *DialogflowCXEntityTypeRef
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref: &DialogflowCXEntityTypeRef{
				External: "projects/my-project/locations/us-central1/agents/my-agent/entityTypes/my-entity-type",
			},
			wantErr: false,
		},
		{
			name: "invalid external reference",
			ref: &DialogflowCXEntityTypeRef{
				External: "invalid/external/ref",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.ValidateExternal(tt.ref.External)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDialogflowCXEntityTypeRef_Normalize(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() error = %v", err)
	}

	readyObjWithExternalRef := &DialogflowCXEntityType{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-extref",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXEntityTypeSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXEntityTypeStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/entityTypes/entity-1"),
		},
	}

	readyObjWithName := &DialogflowCXEntityType{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-name",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXEntityTypeSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXEntityTypeStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			Name: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/entityTypes/entity-2"),
		},
	}

	readyObjWithSpecOnly := &DialogflowCXEntityType{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-spec",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXEntityTypeSpec{
			Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
			ResourceID: common.LazyPtr("entity-3"),
		},
		Status: DialogflowCXEntityTypeStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}

	notReadyObj := &DialogflowCXEntityType{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXEntityTypeSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXEntityTypeStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "False",
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(readyObjWithExternalRef, readyObjWithName, readyObjWithSpecOnly, notReadyObj).
		Build()

	tests := []struct {
		name             string
		ref              *DialogflowCXEntityTypeRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &DialogflowCXEntityTypeRef{
				External: "projects/p1/locations/l1/agents/a1/entityTypes/e1",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/p1/locations/l1/agents/a1/entityTypes/e1",
		},
		{
			name: "external reference invalid",
			ref: &DialogflowCXEntityTypeRef{
				External: "invalid/format",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name and namespace when ready with status.externalRef",
			ref: &DialogflowCXEntityTypeRef{
				Name:      "ready-extref",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/entityTypes/entity-1",
		},
		{
			name: "by name and namespace when ready with status.name",
			ref: &DialogflowCXEntityTypeRef{
				Name:      "ready-name",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/entityTypes/entity-2",
		},
		{
			name: "by name and namespace when ready with spec fallback",
			ref: &DialogflowCXEntityTypeRef{
				Name:      "ready-spec",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/entityTypes/entity-3",
		},
		{
			name: "by name and namespace when not ready",
			ref: &DialogflowCXEntityTypeRef{
				Name:      "not-ready",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "not found",
			ref: &DialogflowCXEntityTypeRef{
				Name:      "non-existent",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(ctx, fakeClient, tt.defaultNamespace)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && tt.ref.External != tt.wantExternal {
				t.Errorf("Normalize() external = %q, want %q", tt.ref.External, tt.wantExternal)
			}
		})
	}
}
