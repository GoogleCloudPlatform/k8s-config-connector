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

func TestEventarcEnrollmentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name           string
		ref            string
		wantProject    string
		wantLocation   string
		wantEnrollment string
		wantErr        bool
	}{
		{
			name:           "valid full URL",
			ref:            "//eventarc.googleapis.com/projects/my-project/locations/us-central1/enrollments/my-enrollment",
			wantProject:    "my-project",
			wantLocation:   "us-central1",
			wantEnrollment: "my-enrollment",
		},
		{
			name:           "valid relative path",
			ref:            "projects/my-project/locations/us-central1/enrollments/my-enrollment",
			wantProject:    "my-project",
			wantLocation:   "us-central1",
			wantEnrollment: "my-enrollment",
		},
		{
			name:    "invalid format - missing location",
			ref:     "projects/my-project/enrollments/my-enrollment",
			wantErr: true,
		},
		{
			name:    "invalid format - extra segments",
			ref:     "projects/my-project/locations/us-central1/enrollments/my-enrollment/extra",
			wantErr: true,
		},
		{
			name:    "empty enrollment",
			ref:     "projects/my-project/locations/us-central1/enrollments/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &EventarcEnrollmentIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("FromExternal() Project got = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("FromExternal() Location got = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.Enrollment != tt.wantEnrollment {
					t.Errorf("FromExternal() Enrollment got = %v, want %v", i.Enrollment, tt.wantEnrollment)
				}
			}
		})
	}
}

func TestEventarcEnrollmentIdentity_Methods(t *testing.T) {
	i := &EventarcEnrollmentIdentity{
		Project:    "my-project",
		Location:   "us-central1",
		Enrollment: "my-enrollment",
	}

	want := "projects/my-project/locations/us-central1/enrollments/my-enrollment"
	if got := i.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}
