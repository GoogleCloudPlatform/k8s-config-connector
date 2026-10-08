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
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDialogflowKnowledgeBaseIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DialogflowKnowledgeBaseIdentity
	}{
		{
			name: "valid regional reference",
			ref:  "projects/my-project/locations/us-central1/knowledgeBases/my-kb",
			want: &DialogflowKnowledgeBaseIdentity{
				Project:        "my-project",
				Location:       "us-central1",
				Knowledge_base: "my-kb",
			},
		},
		{
			name: "valid global reference",
			ref:  "projects/my-project/knowledgeBases/my-kb",
			want: &DialogflowKnowledgeBaseIdentity{
				Project:        "my-project",
				Location:       "",
				Knowledge_base: "my-kb",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full regional url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/locations/us-central1/knowledgeBases/my-kb",
			want: &DialogflowKnowledgeBaseIdentity{
				Project:        "my-project",
				Location:       "us-central1",
				Knowledge_base: "my-kb",
			},
		},
		{
			name: "full global url",
			ref:  "https://dialogflow.googleapis.com/projects/my-project/knowledgeBases/my-kb",
			want: &DialogflowKnowledgeBaseIdentity{
				Project:        "my-project",
				Location:       "",
				Knowledge_base: "my-kb",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DialogflowKnowledgeBaseIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
