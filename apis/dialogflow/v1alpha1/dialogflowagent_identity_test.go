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

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDialogflowAgentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name       string
		ref        string
		wantErr    bool
		want       *DialogflowAgentIdentity
		wantString string
		wantParent string
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/agent",
			want: &DialogflowAgentIdentity{
				Project: "my-project",
			},
			wantString: "projects/my-project/agent",
			wantParent: "projects/my-project",
		},
		{
			name: "valid full url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/agent",
			want: &DialogflowAgentIdentity{
				Project: "my-project",
			},
			wantString: "projects/my-project/agent",
			wantParent: "projects/my-project",
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "invalid collection",
			ref:     "projects/my-project/locations/global/agents/agent-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DialogflowAgentIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != tt.wantString {
					t.Errorf("String() = %v, want %v", got, tt.wantString)
				}
				if got := i.ParentString(); got != tt.wantParent {
					t.Errorf("ParentString() = %v, want %v", got, tt.wantParent)
				}
				if got, want := i.Host(), "dialogflow.googleapis.com"; got != want {
					t.Errorf("Host() = %v, want %v", got, want)
				}
			}
		})
	}
}

func TestDialogflowAgent_GetIdentity(t *testing.T) {
	tests := []struct {
		name        string
		obj         *DialogflowAgent
		want        *DialogflowAgentIdentity
		wantErr     bool
		errContains string
	}{
		{
			name: "resolved via project-id annotation",
			obj: &DialogflowAgent{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-agent",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DialogflowAgentSpec{
					DisplayName: "example-agent",
				},
			},
			want: &DialogflowAgentIdentity{
				Project: "my-project",
			},
		},
		{
			name: "matches status.externalRef",
			obj: &DialogflowAgent{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-agent",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DialogflowAgentSpec{
					DisplayName: "example-agent",
				},
				Status: DialogflowAgentStatus{
					ExternalRef: func() *string {
						s := "projects/my-project/agent"
						return &s
					}(),
				},
			},
			want: &DialogflowAgentIdentity{
				Project: "my-project",
			},
		},
		{
			name: "conflicts with status.externalRef",
			obj: &DialogflowAgent{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-agent",
					Namespace: "default",
					Annotations: map[string]string{
						"cnrm.cloud.google.com/project-id": "my-project",
					},
				},
				Spec: DialogflowAgentSpec{
					DisplayName: "example-agent",
				},
				Status: DialogflowAgentStatus{
					ExternalRef: func() *string {
						s := "projects/other-project/agent"
						return &s
					}(),
				},
			},
			wantErr:     true,
			errContains: "cannot change DialogflowAgent identity",
		},
		{
			name: "missing project resolution",
			obj: &DialogflowAgent{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-agent",
					Namespace: "default",
				},
				Spec: DialogflowAgentSpec{
					DisplayName: "example-agent",
				},
			},
			wantErr:     true,
			errContains: "cannot resolve project",
		},
	}

	fakeClient := fake.NewFakeClient()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(context.Background(), fakeClient)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("GetIdentity() error = %v, want substring %q", err, tc.errContains)
				}
				return
			}
			agentID, ok := id.(*DialogflowAgentIdentity)
			if !ok {
				t.Fatalf("unexpected identity type: %T", id)
			}
			if diff := cmp.Diff(tc.want, agentID); diff != "" {
				t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
