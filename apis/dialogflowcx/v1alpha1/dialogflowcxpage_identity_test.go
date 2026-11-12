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

func TestDialogflowCXPageIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowCXPageIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/my-page",
			want: &DialogflowCXPageIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Flow:     "my-flow",
				Page:     "my-page",
			},
		},
		{
			name: "full url with scheme",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/my-page",
			want: &DialogflowCXPageIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Flow:     "my-flow",
				Page:     "my-page",
			},
		},
		{
			name: "full url without scheme",
			ref:  "//dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/my-page",
			want: &DialogflowCXPageIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Flow:     "my-flow",
				Page:     "my-page",
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
			i := &DialogflowCXPageIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/my-page" {
					t.Errorf("String() = %q, want %q", got, "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/my-page")
				}
			}
		})
	}
}

func TestDialogflowCXPageIdentity_Methods(t *testing.T) {
	id := &DialogflowCXPageIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Agent:    "my-agent",
		Flow:     "my-flow",
		Page:     "my-page",
	}

	if got, want := id.ParentString(), "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"; got != want {
		t.Errorf("ParentString() = %q, want %q", got, want)
	}

	if got, want := id.Host(), "dialogflow.googleapis.com"; got != want {
		t.Errorf("Host() = %q, want %q", got, want)
	}

	if got, want := id.String(), "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/my-page"; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}

	if !id.HasIdentitySpecified() {
		t.Errorf("HasIdentitySpecified() = false, want true")
	}

	emptyID := &DialogflowCXPageIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Agent:    "my-agent",
		Flow:     "my-flow",
	}
	if emptyID.HasIdentitySpecified() {
		t.Errorf("emptyID.HasIdentitySpecified() = true, want false")
	}
}

func TestDialogflowCXPage_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *DialogflowCXPage
		want    *DialogflowCXPageIdentity
		wantErr bool
	}{
		{
			name: "spec with resourceID",
			obj: &DialogflowCXPage{
				Spec: DialogflowCXPageSpec{
					Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
					ResourceID: common.LazyPtr("custom-id"),
				},
			},
			want: &DialogflowCXPageIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Flow:     "my-flow",
				Page:     "custom-id",
			},
		},
		{
			name: "spec without resourceID (server-generated identity pending)",
			obj: &DialogflowCXPage{
				Spec: DialogflowCXPageSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
				},
			},
			want: &DialogflowCXPageIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Flow:     "my-flow",
				Page:     "",
			},
		},
		{
			name: "spec without resourceID, resolved from status.externalRef",
			obj: &DialogflowCXPage{
				Spec: DialogflowCXPageSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
				},
				Status: DialogflowCXPageStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/server-id-1"),
				},
			},
			want: &DialogflowCXPageIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Flow:     "my-flow",
				Page:     "server-id-1",
			},
		},
		{
			name: "spec without resourceID, resolved from status.name",
			obj: &DialogflowCXPage{
				Spec: DialogflowCXPageSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
				},
				Status: DialogflowCXPageStatus{
					Name: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/server-id-2"),
				},
			},
			want: &DialogflowCXPageIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Flow:     "my-flow",
				Page:     "server-id-2",
			},
		},
		{
			name: "drift detection error when spec resourceID does not match status",
			obj: &DialogflowCXPage{
				Spec: DialogflowCXPageSpec{
					Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
					ResourceID: common.LazyPtr("spec-id"),
				},
				Status: DialogflowCXPageStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/different-id"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing parent",
			obj: &DialogflowCXPage{
				Spec: DialogflowCXPageSpec{
					ResourceID: common.LazyPtr("custom-id"),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid parent format",
			obj: &DialogflowCXPage{
				Spec: DialogflowCXPageSpec{
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
				gotID, ok := got.(*DialogflowCXPageIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned %T, want *DialogflowCXPageIdentity", got)
				}
				if diff := cmp.Diff(tt.want, gotID); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestDialogflowCXPageRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     *DialogflowCXPageRef
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref: &DialogflowCXPageRef{
				External: "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/my-page",
			},
			wantErr: false,
		},
		{
			name: "invalid external reference",
			ref: &DialogflowCXPageRef{
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

func TestDialogflowCXPageRef_Normalize(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() error = %v", err)
	}

	readyObjWithExternalRef := &DialogflowCXPage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-extref",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXPageSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
		},
		Status: DialogflowCXPageStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/page-1"),
		},
	}

	readyObjWithName := &DialogflowCXPage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-name",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXPageSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
		},
		Status: DialogflowCXPageStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			Name: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/page-2"),
		},
	}

	readyObjWithSpecOnly := &DialogflowCXPage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-spec",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXPageSpec{
			Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
			ResourceID: common.LazyPtr("page-3"),
		},
		Status: DialogflowCXPageStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}

	notReadyObj := &DialogflowCXPage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXPageSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow"),
		},
		Status: DialogflowCXPageStatus{
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
		ref              *DialogflowCXPageRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &DialogflowCXPageRef{
				External: "projects/p1/locations/l1/agents/a1/flows/f1/pages/p1",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/p1/locations/l1/agents/a1/flows/f1/pages/p1",
		},
		{
			name: "external reference invalid",
			ref: &DialogflowCXPageRef{
				External: "invalid/format",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name and namespace when ready with status.externalRef",
			ref: &DialogflowCXPageRef{
				Name:      "ready-extref",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/page-1",
		},
		{
			name: "by name and namespace when ready with status.name",
			ref: &DialogflowCXPageRef{
				Name:      "ready-name",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/page-2",
		},
		{
			name: "by name and namespace when ready with spec fallback",
			ref: &DialogflowCXPageRef{
				Name:      "ready-spec",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/flows/my-flow/pages/page-3",
		},
		{
			name: "by name and namespace when not ready",
			ref: &DialogflowCXPageRef{
				Name:      "not-ready",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "not found",
			ref: &DialogflowCXPageRef{
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
