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

func TestBigQueryReservationReservationGroupIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name         string
		ref          string
		wantProject  string
		wantLocation string
		wantID       string
		wantErr      bool
	}{
		{
			name:         "valid external ref",
			ref:          "projects/my-project/locations/us-west2/reservationGroups/my-res-group",
			wantProject:  "my-project",
			wantLocation: "us-west2",
			wantID:       "my-res-group",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/reservationGroups/my-res-group",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigQueryReservationReservationGroupIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("BigQueryReservationReservationGroupIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("BigQueryReservationReservationGroupIdentity.FromExternal() Project = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("BigQueryReservationReservationGroupIdentity.FromExternal() Location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.Reservation_group != tt.wantID {
					t.Errorf("BigQueryReservationReservationGroupIdentity.FromExternal() Reservation_group = %v, want %v", i.Reservation_group, tt.wantID)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("BigQueryReservationReservationGroupIdentity.String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}
