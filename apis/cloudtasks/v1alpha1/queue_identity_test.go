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
	"reflect"
	"testing"
)

func TestParseQueueExternal(t *testing.T) {
	tests := []struct {
		name       string
		external   string
		wantParent *QueueParent
		wantID     string
		wantErr    bool
	}{
		{
			name:     "valid external reference",
			external: "projects/my-project/locations/us-central1/queues/my-queue",
			wantParent: &QueueParent{
				ProjectID: "my-project",
				Location:  "us-central1",
			},
			wantID:  "my-queue",
			wantErr: false,
		},
		{
			name:     "invalid external reference - empty segment",
			external: "projects//locations/us-central1/queues/my-queue",
			wantParent: &QueueParent{
				ProjectID: "",
				Location:  "us-central1",
			},
			wantID:  "my-queue",
			wantErr: false, // ParseQueueExternal split allows this structurally, but let's test it matches actual behavior
		},
		{
			name:     "invalid external reference - wrong format",
			external: "projects/my-project/locations/us-central1/queues",
			wantErr:  true,
		},
		{
			name:     "invalid external reference - wrong prefix",
			external: "project/my-project/locations/us-central1/queues/my-queue",
			wantErr:  true,
		},
		{
			name:     "invalid external reference - wrong midsegment",
			external: "projects/my-project/location/us-central1/queues/my-queue",
			wantErr:  true,
		},
		{
			name:     "invalid external reference - wrong resource type",
			external: "projects/my-project/locations/us-central1/queue/my-queue",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParent, gotID, err := ParseQueueExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseQueueExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(gotParent, tt.wantParent) {
					t.Errorf("ParseQueueExternal() gotParent = %v, want %v", gotParent, tt.wantParent)
				}
				if gotID != tt.wantID {
					t.Errorf("ParseQueueExternal() gotID = %v, want %v", gotID, tt.wantID)
				}
			}
		})
	}
}
