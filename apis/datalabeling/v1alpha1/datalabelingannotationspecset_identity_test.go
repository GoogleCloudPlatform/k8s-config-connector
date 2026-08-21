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
)

func TestDataLabelingAnnotationSpecSetIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name                  string
		ref                   string
		wantProject           string
		wantAnnotationSpecSet string
		wantErr               bool
	}{
		{
			name:                  "valid external ref",
			ref:                   "projects/my-project/annotationSpecSets/my-annotation-spec-set",
			wantProject:           "my-project",
			wantAnnotationSpecSet: "my-annotation-spec-set",
			wantErr:               false,
		},
		{
			name:                  "full url",
			ref:                   "https://datalabeling.googleapis.com/projects/my-project/annotationSpecSets/my-annotation-spec-set",
			wantProject:           "my-project",
			wantAnnotationSpecSet: "my-annotation-spec-set",
			wantErr:               false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/annotationSpecSets",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DataLabelingAnnotationSpecSetIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("DataLabelingAnnotationSpecSetIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("DataLabelingAnnotationSpecSetIdentity.FromExternal() Project = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Annotation_spec_set != tt.wantAnnotationSpecSet {
					t.Errorf("DataLabelingAnnotationSpecSetIdentity.FromExternal() Annotation_spec_set = %v, want %v", i.Annotation_spec_set, tt.wantAnnotationSpecSet)
				}
				if got := i.String(); got != "projects/"+tt.wantProject+"/annotationSpecSets/"+tt.wantAnnotationSpecSet {
					t.Errorf("DataLabelingAnnotationSpecSetIdentity.String() = %v", got)
				}
			}
		})
	}
}
