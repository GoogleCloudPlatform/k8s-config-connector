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

	"github.com/google/go-cmp/cmp"
)

func TestDialogflowToolIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowToolIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/agents/my-agent/tools/my-tool",
			want: &DialogflowToolIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Tool:     "my-tool",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/tools/my-tool",
			want: &DialogflowToolIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Tool:     "my-tool",
			},
		},
		{
			name: "cai format",
			ref:  "//dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent/tools/my-tool",
			want: &DialogflowToolIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
				Tool:     "my-tool",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &DialogflowToolIdentity{}
			err := got.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got.String() != "projects/my-project/locations/us-central1/agents/my-agent/tools/my-tool" {
					t.Errorf("String() = %v, want %v", got.String(), "projects/my-project/locations/us-central1/agents/my-agent/tools/my-tool")
				}
				if got.ParentString() != "projects/my-project/locations/us-central1/agents/my-agent" {
					t.Errorf("ParentString() = %v, want %v", got.ParentString(), "projects/my-project/locations/us-central1/agents/my-agent")
				}
				if got.Host() != "dialogflow.googleapis.com" {
					t.Errorf("Host() = %v, want %v", got.Host(), "dialogflow.googleapis.com")
				}
			}
		})
	}
}

func TestDialogflowCXAgentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowCXAgentIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/agents/my-agent",
			want: &DialogflowCXAgentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent",
			want: &DialogflowCXAgentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
			},
		},
		{
			name: "cai format",
			ref:  "//dialogflow.googleapis.com/projects/my-project/locations/us-central1/agents/my-agent",
			want: &DialogflowCXAgentIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Agent:    "my-agent",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &DialogflowCXAgentIdentity{}
			err := got.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got.String() != "projects/my-project/locations/us-central1/agents/my-agent" {
					t.Errorf("String() = %v, want %v", got.String(), "projects/my-project/locations/us-central1/agents/my-agent")
				}
				if got.Host() != "dialogflow.googleapis.com" {
					t.Errorf("Host() = %v, want %v", got.Host(), "dialogflow.googleapis.com")
				}
			}
		})
	}
}

func TestDialogflowToolRef_Normalize(t *testing.T) {
	tests := []struct {
		name    string
		ref     *DialogflowToolRef
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref: &DialogflowToolRef{
				External: "projects/my-project/locations/us-central1/agents/my-agent/tools/my-tool",
			},
			wantErr: false,
		},
		{
			name: "empty external reference",
			ref: &DialogflowToolRef{
				External: "",
			},
			wantErr: true,
		},
		{
			name: "invalid external reference format",
			ref: &DialogflowToolRef{
				External: "invalid/format",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(context.Background(), nil, "default")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDialogflowCXAgentRef_Normalize(t *testing.T) {
	tests := []struct {
		name    string
		ref     *DialogflowCXAgentRef
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref: &DialogflowCXAgentRef{
				External: "projects/my-project/locations/us-central1/agents/my-agent",
			},
			wantErr: false,
		},
		{
			name: "empty external reference",
			ref: &DialogflowCXAgentRef{
				External: "",
			},
			wantErr: true,
		},
		{
			name: "invalid external reference format",
			ref: &DialogflowCXAgentRef{
				External: "invalid/format",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ref.Normalize(context.Background(), nil, "default")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
