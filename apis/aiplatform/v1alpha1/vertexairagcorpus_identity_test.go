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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/google/go-cmp/cmp"
)

func TestVertexAIRagCorpusIdentity_Conformance(t *testing.T) {
	identity.AssertConformance(t, identity.ConformanceCase[VertexAIRagCorpusIdentity]{
		SampleExternal: "projects/my-project/locations/us-central1/ragCorpora/my-rag-corpus",
		WantHost:       "aiplatform.googleapis.com",
		WantParent:     "projects/my-project/locations/us-central1",
		WantIdentity: &VertexAIRagCorpusIdentity{
			Project:   "my-project",
			Location:  "us-central1",
			RagCorpus: "my-rag-corpus",
		},
		Ref: &VertexAIRagCorpusRef{},
	})
}

func TestVertexAIRagCorpusIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAIRagCorpusIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/ragCorpora/my-rag-corpus",
			want: &VertexAIRagCorpusIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				RagCorpus: "my-rag-corpus",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/ragCorpora/my-rag-corpus",
			want: &VertexAIRagCorpusIdentity{
				Project:   "my-project",
				Location:  "us-central1",
				RagCorpus: "my-rag-corpus",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &VertexAIRagCorpusIdentity{}
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

func TestExternalReferences_ValidateExternal(t *testing.T) {
	t.Run("VertexAIIndexEndpointRef", func(t *testing.T) {
		r := &VertexAIIndexEndpointRef{}
		if err := r.ValidateExternal("projects/p/locations/l/indexEndpoints/ie"); err != nil {
			t.Errorf("ValidateExternal() failed: %v", err)
		}
		if err := r.ValidateExternal("invalid/ie"); err == nil {
			t.Errorf("ValidateExternal() expected error for invalid external ref, got nil")
		}
	})

	t.Run("VertexAIIndexRef", func(t *testing.T) {
		r := &VertexAIIndexRef{}
		if err := r.ValidateExternal("projects/p/locations/l/indexes/i"); err != nil {
			t.Errorf("ValidateExternal() failed: %v", err)
		}
		if err := r.ValidateExternal("invalid/i"); err == nil {
			t.Errorf("ValidateExternal() expected error for invalid external ref, got nil")
		}
	})

	t.Run("VertexAIEndpointRef", func(t *testing.T) {
		r := &VertexAIEndpointRef{}
		if err := r.ValidateExternal("projects/p/locations/l/endpoints/e"); err != nil {
			t.Errorf("ValidateExternal() failed: %v", err)
		}
		if err := r.ValidateExternal("invalid/e"); err == nil {
			t.Errorf("ValidateExternal() expected error for invalid external ref, got nil")
		}
	})
}
