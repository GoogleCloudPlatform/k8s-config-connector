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

package v1beta1

import (
	"testing"
)

func TestRecognizerIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *RecognizerIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/recognizers/my-recognizer",
			want: &RecognizerIdentity{
				parent: &RecognizerParent{
					ProjectID: "my-project",
					Location:  "global",
				},
				id: "my-recognizer",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url v2",
			ref:  "https://speech.googleapis.com/v2/projects/my-project/locations/global/recognizers/my-recognizer",
			want: &RecognizerIdentity{
				parent: &RecognizerParent{
					ProjectID: "my-project",
					Location:  "global",
				},
				id: "my-recognizer",
			},
		},
		{
			name: "full url v1beta1",
			ref:  "https://speech.googleapis.com/v1beta1/projects/my-project/locations/global/recognizers/my-recognizer",
			want: &RecognizerIdentity{
				parent: &RecognizerParent{
					ProjectID: "my-project",
					Location:  "global",
				},
				id: "my-recognizer",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &RecognizerIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.parent.ProjectID != tt.want.parent.ProjectID {
					t.Errorf("ProjectID = %v, want %v", i.parent.ProjectID, tt.want.parent.ProjectID)
				}
				if i.parent.Location != tt.want.parent.Location {
					t.Errorf("Location = %v, want %v", i.parent.Location, tt.want.parent.Location)
				}
				if i.id != tt.want.id {
					t.Errorf("ID = %v, want %v", i.id, tt.want.id)
				}
				if i.Host() != "speech.googleapis.com" {
					t.Errorf("Host() = %v, want speech.googleapis.com", i.Host())
				}
			}
		})
	}
}
