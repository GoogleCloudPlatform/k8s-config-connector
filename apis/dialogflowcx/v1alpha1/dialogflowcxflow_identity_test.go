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

func TestDialogflowCXFlowIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowCXFlowIdentity
	}{
		{
			name: "valid relative reference",
			ref:  "projects/test-project/locations/us-central1/agents/test-agent/flows/test-flow",
			want: &DialogflowCXFlowIdentity{
				Project:  "test-project",
				Location: "us-central1",
				Agent:    "test-agent",
				Flow:     "test-flow",
			},
		},
		{
			name: "valid external reference with host prefix",
			ref:  "//dialogflow.googleapis.com/projects/test-project/locations/us-central1/agents/test-agent/flows/test-flow",
			want: &DialogflowCXFlowIdentity{
				Project:  "test-project",
				Location: "us-central1",
				Agent:    "test-agent",
				Flow:     "test-flow",
			},
		},
		{
			name:    "invalid format",
			ref:     "projects/test-project/locations/us-central1/agents/test-agent",
			wantErr: true,
		},
		{
			name:    "empty reference",
			ref:     "",
			wantErr: true,
		},
		{
			name:    "wrong host",
			ref:     "//other.googleapis.com/projects/test-project/locations/us-central1/agents/test-agent/flows/test-flow",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &DialogflowCXFlowIdentity{}
			err := id.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal(%q) unexpected error: %v, wantErr: %v", tt.ref, err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, id); diff != "" {
					t.Errorf("FromExternal(%q) mismatch (-want +got):\n%s", tt.ref, diff)
				}
			}
		})
	}
}

func TestDialogflowCXFlowIdentity_Methods(t *testing.T) {
	id := &DialogflowCXFlowIdentity{
		Project:  "test-proj",
		Location: "us-central1",
		Agent:    "test-agent",
		Flow:     "test-flow",
	}

	wantStr := "projects/test-proj/locations/us-central1/agents/test-agent/flows/test-flow"
	if diff := cmp.Diff(wantStr, id.String()); diff != "" {
		t.Errorf("String() mismatch (-want +got):\n%s", diff)
	}

	wantParent := "projects/test-proj/locations/us-central1/agents/test-agent"
	if diff := cmp.Diff(wantParent, id.ParentString()); diff != "" {
		t.Errorf("ParentString() mismatch (-want +got):\n%s", diff)
	}

	wantHost := "dialogflow.googleapis.com"
	if diff := cmp.Diff(wantHost, id.Host()); diff != "" {
		t.Errorf("Host() mismatch (-want +got):\n%s", diff)
	}

	if !id.HasIdentitySpecified() {
		t.Errorf("HasIdentitySpecified() = false, want true")
	}

	idEmpty := &DialogflowCXFlowIdentity{
		Project:  "test-proj",
		Location: "us-central1",
		Agent:    "test-agent",
		Flow:     "",
	}
	if idEmpty.HasIdentitySpecified() {
		t.Errorf("HasIdentitySpecified() = true, want false")
	}
}

func TestDialogflowCXFlow_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name                  string
		obj                   *DialogflowCXFlow
		wantErr               bool
		want                  *DialogflowCXFlowIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &DialogflowCXFlow{
				Spec: DialogflowCXFlowSpec{
					Parent:     common.LazyPtr("projects/my-proj/locations/global/agents/my-agent"),
					ResourceID: common.LazyPtr("00000000-0000-0000-0000-000000000001"),
				},
			},
			want: &DialogflowCXFlowIdentity{
				Project:  "my-proj",
				Location: "global",
				Agent:    "my-agent",
				Flow:     "00000000-0000-0000-0000-000000000001",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with server-generated identity defaulted from status.externalRef",
			obj: &DialogflowCXFlow{
				Spec: DialogflowCXFlowSpec{
					Parent: common.LazyPtr("projects/my-proj/locations/global/agents/my-agent"),
				},
				Status: DialogflowCXFlowStatus{
					ExternalRef: common.LazyPtr("projects/my-proj/locations/global/agents/my-agent/flows/server-flow-id"),
				},
			},
			want: &DialogflowCXFlowIdentity{
				Project:  "my-proj",
				Location: "global",
				Agent:    "my-agent",
				Flow:     "server-flow-id",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with server-generated identity defaulted from status.name (legacy TF)",
			obj: &DialogflowCXFlow{
				Spec: DialogflowCXFlowSpec{
					Parent: common.LazyPtr("projects/my-proj/locations/global/agents/my-agent"),
				},
				Status: DialogflowCXFlowStatus{
					Name: common.LazyPtr("projects/my-proj/locations/global/agents/my-agent/flows/tf-flow-id"),
				},
			},
			want: &DialogflowCXFlowIdentity{
				Project:  "my-proj",
				Location: "global",
				Agent:    "my-agent",
				Flow:     "tf-flow-id",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity without resourceID and without status (unreconciled)",
			obj: &DialogflowCXFlow{
				Spec: DialogflowCXFlowSpec{
					Parent: common.LazyPtr("projects/my-proj/locations/global/agents/my-agent"),
				},
			},
			want: &DialogflowCXFlowIdentity{
				Project:  "my-proj",
				Location: "global",
				Agent:    "my-agent",
				Flow:     "",
			},
			wantIdentitySpecified: false,
		},
		{
			name: "GetIdentity error on identity change (spec resourceID does not match status.externalRef)",
			obj: &DialogflowCXFlow{
				Spec: DialogflowCXFlowSpec{
					Parent:     common.LazyPtr("projects/my-proj/locations/global/agents/my-agent"),
					ResourceID: common.LazyPtr("flow-1"),
				},
				Status: DialogflowCXFlowStatus{
					ExternalRef: common.LazyPtr("projects/my-proj/locations/global/agents/my-agent/flows/flow-2"),
				},
			},
			wantErr: true,
		},
		{
			name: "GetIdentity error when spec.parent is missing",
			obj: &DialogflowCXFlow{
				Spec: DialogflowCXFlowSpec{
					ResourceID: common.LazyPtr("flow-1"),
				},
			},
			wantErr: true,
		},
		{
			name: "GetIdentity error when spec.parent is malformed",
			obj: &DialogflowCXFlow{
				Spec: DialogflowCXFlowSpec{
					Parent:     common.LazyPtr("invalid/parent"),
					ResourceID: common.LazyPtr("flow-1"),
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetIdentity() unexpected error: %v, wantErr: %v", err, tc.wantErr)
			}
			if !tc.wantErr {
				flowID, ok := id.(*DialogflowCXFlowIdentity)
				if !ok {
					t.Fatalf("expected *DialogflowCXFlowIdentity, got %T", id)
				}
				if diff := cmp.Diff(tc.want, flowID); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if flowID.HasIdentitySpecified() != tc.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", flowID.HasIdentitySpecified(), tc.wantIdentitySpecified)
				}
			}
		})
	}
}

func TestDialogflowCXFlowRef_Normalize(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() error = %v", err)
	}

	readyObjWithExternalRef := &DialogflowCXFlow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-extref",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXFlowSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXFlowStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/flow-1"),
		},
	}

	readyObjWithName := &DialogflowCXFlow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-name",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXFlowSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXFlowStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			Name: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/flows/flow-2"),
		},
	}

	readyObjWithSpecOnly := &DialogflowCXFlow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-spec",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXFlowSpec{
			Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
			ResourceID: common.LazyPtr("flow-3"),
		},
		Status: DialogflowCXFlowStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}

	notReadyObj := &DialogflowCXFlow{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXFlowSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXFlowStatus{
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
		ref              *DialogflowCXFlowRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &DialogflowCXFlowRef{
				External: "projects/p1/locations/l1/agents/a1/flows/f1",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/p1/locations/l1/agents/a1/flows/f1",
		},
		{
			name: "external reference invalid",
			ref: &DialogflowCXFlowRef{
				External: "invalid/format",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name and namespace when ready with status.externalRef",
			ref: &DialogflowCXFlowRef{
				Name:      "ready-extref",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/flows/flow-1",
		},
		{
			name: "by name and namespace when ready with status.name",
			ref: &DialogflowCXFlowRef{
				Name:      "ready-name",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/flows/flow-2",
		},
		{
			name: "by name and namespace when ready with spec fallback",
			ref: &DialogflowCXFlowRef{
				Name:      "ready-spec",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/flows/flow-3",
		},
		{
			name: "by name and namespace when not ready",
			ref: &DialogflowCXFlowRef{
				Name:      "not-ready",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "not found",
			ref: &DialogflowCXFlowRef{
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
