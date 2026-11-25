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
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDialogflowSipTrunkIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowSipTrunkIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/sipTrunks/my-siptrunk",
			want: &DialogflowSipTrunkIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Siptrunk: "my-siptrunk",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/sipTrunks/my-siptrunk",
			want: &DialogflowSipTrunkIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Siptrunk: "my-siptrunk",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DialogflowSipTrunkIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Location != tt.want.Location {
					t.Errorf("Location = %v, want %v", i.Location, tt.want.Location)
				}
				if i.Siptrunk != tt.want.Siptrunk {
					t.Errorf("Siptrunk = %v, want %v", i.Siptrunk, tt.want.Siptrunk)
				}
			}
		})
	}
}

func TestDialogflowSipTrunk_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		obj         *DialogflowSipTrunk
		wantID      *DialogflowSipTrunkIdentity
		wantErr     bool
		errContains string
	}{
		{
			name: "spec identity without resourceID (server-generated ID before create)",
			obj: &DialogflowSipTrunk{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-trunk",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DialogflowSipTrunkSpec{
					Location: direct.PtrTo("global"),
				},
			},
			wantID: &DialogflowSipTrunkIdentity{
				Project:  "my-project",
				Location: "global",
				Siptrunk: "",
			},
		},
		{
			name: "spec identity without resourceID, status has server-generated ID",
			obj: &DialogflowSipTrunk{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-trunk",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DialogflowSipTrunkSpec{
					Location: direct.PtrTo("global"),
				},
				Status: DialogflowSipTrunkStatus{
					ExternalRef: direct.PtrTo("projects/my-project/locations/global/sipTrunks/server-gen-123"),
				},
			},
			wantID: &DialogflowSipTrunkIdentity{
				Project:  "my-project",
				Location: "global",
				Siptrunk: "server-gen-123",
			},
		},
		{
			name: "spec identity with resourceID, matches status",
			obj: &DialogflowSipTrunk{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-trunk",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DialogflowSipTrunkSpec{
					Location:   direct.PtrTo("global"),
					ResourceID: direct.PtrTo("my-id"),
				},
				Status: DialogflowSipTrunkStatus{
					ExternalRef: direct.PtrTo("projects/my-project/locations/global/sipTrunks/my-id"),
				},
			},
			wantID: &DialogflowSipTrunkIdentity{
				Project:  "my-project",
				Location: "global",
				Siptrunk: "my-id",
			},
		},
		{
			name: "spec identity with resourceID, mismatches status",
			obj: &DialogflowSipTrunk{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-trunk",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DialogflowSipTrunkSpec{
					Location:   direct.PtrTo("global"),
					ResourceID: direct.PtrTo("new-id"),
				},
				Status: DialogflowSipTrunkStatus{
					ExternalRef: direct.PtrTo("projects/my-project/locations/global/sipTrunks/old-id"),
				},
			},
			wantErr:     true,
			errContains: "cannot change DialogflowSipTrunk identity",
		},
		{
			name: "status parent project mismatches spec parent project",
			obj: &DialogflowSipTrunk{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-trunk",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project-2",
					},
				},
				Spec: DialogflowSipTrunkSpec{
					Location: direct.PtrTo("global"),
				},
				Status: DialogflowSipTrunkStatus{
					ExternalRef: direct.PtrTo("projects/my-project-1/locations/global/sipTrunks/server-gen-123"),
				},
			},
			wantErr:     true,
			errContains: "cannot change DialogflowSipTrunk parent project",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, nil)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Fatalf("expected error to contain %q, got %q", tc.errContains, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			gotID := id.(*DialogflowSipTrunkIdentity)
			if *gotID != *tc.wantID {
				t.Fatalf("expected identity %+v, got %+v", tc.wantID, gotID)
			}
		})
	}
}
