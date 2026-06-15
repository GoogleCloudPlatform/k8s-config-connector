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

func TestAutoMLDatasetIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *AutoMLDatasetIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/datasets/my-dataset",
			want: &AutoMLDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://automl.googleapis.com/projects/my-project/locations/us-central1/datasets/my-dataset",
			want: &AutoMLDatasetIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Dataset:  "my-dataset",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AutoMLDatasetIdentity{}
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
				if i.Dataset != tt.want.Dataset {
					t.Errorf("Dataset = %v, want %v", i.Dataset, tt.want.Dataset)
				}
			}
		})
	}
}
