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
	"errors"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	s := runtime.NewScheme()

	uFlowWithExternalRef := &unstructured.Unstructured{}
	uFlowWithExternalRef.SetGroupVersionKind(DialogflowCXFlowGVK)
	uFlowWithExternalRef.SetName("flow-with-externalref")
	uFlowWithExternalRef.SetNamespace("test-ns")
	uFlowWithExternalRef.Object["status"] = map[string]interface{}{
		"externalRef": "projects/test-project/locations/us-central1/agents/test-agent/flows/flow-ext",
	}

	uFlowWithStatusName := &unstructured.Unstructured{}
	uFlowWithStatusName.SetGroupVersionKind(DialogflowCXFlowGVK)
	uFlowWithStatusName.SetName("flow-with-statusname")
	uFlowWithStatusName.SetNamespace("test-ns")
	uFlowWithStatusName.Object["status"] = map[string]interface{}{
		"name": "projects/test-project/locations/us-central1/agents/test-agent/flows/flow-tf",
	}

	uFlowUnready := &unstructured.Unstructured{}
	uFlowUnready.SetGroupVersionKind(DialogflowCXFlowGVK)
	uFlowUnready.SetName("flow-unready")
	uFlowUnready.SetNamespace("test-ns")
	uFlowUnready.Object["spec"] = map[string]interface{}{
		"parent": "projects/test-project/locations/us-central1/agents/test-agent",
	}

	fakeClient := fake.NewClientBuilder().WithScheme(s).WithObjects(
		uFlowWithExternalRef,
		uFlowWithStatusName,
		uFlowUnready,
	).Build()

	tests := []struct {
		name         string
		ref          *DialogflowCXFlowRef
		wantExternal string
		wantErr      bool
		wantNotReady bool
	}{
		{
			name: "external reference already set",
			ref: &DialogflowCXFlowRef{
				External: "projects/test-project/locations/us-central1/agents/test-agent/flows/my-flow",
			},
			wantExternal: "projects/test-project/locations/us-central1/agents/test-agent/flows/my-flow",
		},
		{
			name: "internal reference resolved from status.externalRef",
			ref: &DialogflowCXFlowRef{
				Name:      "flow-with-externalref",
				Namespace: "test-ns",
			},
			wantExternal: "projects/test-project/locations/us-central1/agents/test-agent/flows/flow-ext",
		},
		{
			name: "internal reference resolved from status.name (legacy fallback)",
			ref: &DialogflowCXFlowRef{
				Name:      "flow-with-statusname",
				Namespace: "test-ns",
			},
			wantExternal: "projects/test-project/locations/us-central1/agents/test-agent/flows/flow-tf",
		},
		{
			name: "internal reference not ready",
			ref: &DialogflowCXFlowRef{
				Name:      "flow-unready",
				Namespace: "test-ns",
			},
			wantErr:      true,
			wantNotReady: true,
		},
		{
			name: "internal reference not found",
			ref: &DialogflowCXFlowRef{
				Name:      "flow-not-found",
				Namespace: "test-ns",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.ref.Normalize(ctx, fakeClient, "test-ns")
			if (err != nil) != tc.wantErr {
				t.Fatalf("Normalize() unexpected error: %v, wantErr: %v", err, tc.wantErr)
			}
			if tc.wantNotReady {
				var notReadyErr *k8s.ReferenceNotReadyError
				if !errors.As(err, &notReadyErr) {
					t.Fatalf("expected ReferenceNotReadyError, got %T: %v", err, err)
				}
			}
			if !tc.wantErr {
				if diff := cmp.Diff(tc.wantExternal, tc.ref.External); diff != "" {
					t.Errorf("Normalize() External mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
