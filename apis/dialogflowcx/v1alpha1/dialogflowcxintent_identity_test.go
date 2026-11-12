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

func TestDialogflowCXIntentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowCXIntentIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/agents/my-agent/intents/my-intent",
			want: &DialogflowCXIntentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Intent:   "my-intent",
			},
		},
		{
			name: "full url with scheme",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/intents/my-intent",
			want: &DialogflowCXIntentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Intent:   "my-intent",
			},
		},
		{
			name: "full url without scheme",
			ref:  "//dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/intents/my-intent",
			want: &DialogflowCXIntentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Intent:   "my-intent",
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
			i := &DialogflowCXIntentIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != "projects/my-project/locations/us-central1/agents/my-agent/intents/my-intent" {
					t.Errorf("String() = %q, want %q", got, "projects/my-project/locations/us-central1/agents/my-agent/intents/my-intent")
				}
			}
		})
	}
}

func TestDialogflowCXIntentIdentity_Methods(t *testing.T) {
	id := &DialogflowCXIntentIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Agent:    "my-agent",
		Intent:   "my-intent",
	}

	if got, want := id.ParentString(), "projects/my-project/locations/us-central1/agents/my-agent"; got != want {
		t.Errorf("ParentString() = %q, want %q", got, want)
	}

	if got, want := id.Host(), "dialogflow.googleapis.com"; got != want {
		t.Errorf("Host() = %q, want %q", got, want)
	}

	if got, want := id.String(), "projects/my-project/locations/us-central1/agents/my-agent/intents/my-intent"; got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}

	if !id.HasIdentitySpecified() {
		t.Errorf("HasIdentitySpecified() = false, want true")
	}

	emptyID := &DialogflowCXIntentIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Agent:    "my-agent",
	}
	if emptyID.HasIdentitySpecified() {
		t.Errorf("emptyID.HasIdentitySpecified() = true, want false")
	}
}

func TestDialogflowCXIntent_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *DialogflowCXIntent
		want    *DialogflowCXIntentIdentity
		wantErr bool
	}{
		{
			name: "spec with resourceID",
			obj: &DialogflowCXIntent{
				Spec: DialogflowCXIntentSpec{
					Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
					ResourceID: common.LazyPtr("custom-id"),
				},
			},
			want: &DialogflowCXIntentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Intent:   "custom-id",
			},
		},
		{
			name: "spec without resourceID (server-generated identity pending)",
			obj: &DialogflowCXIntent{
				Spec: DialogflowCXIntentSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
				},
			},
			want: &DialogflowCXIntentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Intent:   "",
			},
		},
		{
			name: "spec without resourceID, resolved from status.externalRef",
			obj: &DialogflowCXIntent{
				Spec: DialogflowCXIntentSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
				},
				Status: DialogflowCXIntentStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/intents/server-id-1"),
				},
			},
			want: &DialogflowCXIntentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Intent:   "server-id-1",
			},
		},
		{
			name: "spec without resourceID, resolved from status.name",
			obj: &DialogflowCXIntent{
				Spec: DialogflowCXIntentSpec{
					Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
				},
				Status: DialogflowCXIntentStatus{
					Name: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/intents/server-id-2"),
				},
			},
			want: &DialogflowCXIntentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Intent:   "server-id-2",
			},
		},
		{
			name: "drift detection error when spec resourceID does not match status",
			obj: &DialogflowCXIntent{
				Spec: DialogflowCXIntentSpec{
					Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
					ResourceID: common.LazyPtr("spec-id"),
				},
				Status: DialogflowCXIntentStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/intents/different-id"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing parent",
			obj: &DialogflowCXIntent{
				Spec: DialogflowCXIntentSpec{
					ResourceID: common.LazyPtr("custom-id"),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid parent format",
			obj: &DialogflowCXIntent{
				Spec: DialogflowCXIntentSpec{
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
				gotID, ok := got.(*DialogflowCXIntentIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned %T, want *DialogflowCXIntentIdentity", got)
				}
				if diff := cmp.Diff(tt.want, gotID); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestDialogflowCXIntentRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     *DialogflowCXIntentRef
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref: &DialogflowCXIntentRef{
				External: "projects/my-project/locations/us-central1/agents/my-agent/intents/my-intent",
			},
			wantErr: false,
		},
		{
			name: "invalid external reference",
			ref: &DialogflowCXIntentRef{
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

func TestDialogflowCXIntentRef_Normalize(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() error = %v", err)
	}

	readyObjWithExternalRef := &DialogflowCXIntent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-extref",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXIntentSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXIntentStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/intents/intent-1"),
		},
	}

	readyObjWithName := &DialogflowCXIntent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-name",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXIntentSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXIntentStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
			Name: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent/intents/intent-2"),
		},
	}

	readyObjWithSpecOnly := &DialogflowCXIntent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ready-spec",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXIntentSpec{
			Parent:     common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
			ResourceID: common.LazyPtr("intent-3"),
		},
		Status: DialogflowCXIntentStatus{
			Conditions: []k8sv1alpha1.Condition{
				{
					Type:   "Ready",
					Status: "True",
				},
			},
		},
	}

	notReadyObj := &DialogflowCXIntent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "not-ready",
			Namespace: "test-ns",
		},
		Spec: DialogflowCXIntentSpec{
			Parent: common.LazyPtr("projects/my-project/locations/us-central1/agents/my-agent"),
		},
		Status: DialogflowCXIntentStatus{
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
		ref              *DialogflowCXIntentRef
		defaultNamespace string
		wantErr          bool
		wantExternal     string
	}{
		{
			name: "external reference valid",
			ref: &DialogflowCXIntentRef{
				External: "projects/p1/locations/l1/agents/a1/intents/i1",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/p1/locations/l1/agents/a1/intents/i1",
		},
		{
			name: "external reference invalid",
			ref: &DialogflowCXIntentRef{
				External: "invalid/format",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "by name and namespace when ready with status.externalRef",
			ref: &DialogflowCXIntentRef{
				Name:      "ready-extref",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/intents/intent-1",
		},
		{
			name: "by name and namespace when ready with status.name",
			ref: &DialogflowCXIntentRef{
				Name:      "ready-name",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/intents/intent-2",
		},
		{
			name: "by name and namespace when ready with spec fallback",
			ref: &DialogflowCXIntentRef{
				Name:      "ready-spec",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          false,
			wantExternal:     "projects/my-project/locations/us-central1/agents/my-agent/intents/intent-3",
		},
		{
			name: "by name and namespace when not ready",
			ref: &DialogflowCXIntentRef{
				Name:      "not-ready",
				Namespace: "test-ns",
			},
			defaultNamespace: "test-ns",
			wantErr:          true,
		},
		{
			name: "not found",
			ref: &DialogflowCXIntentRef{
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
