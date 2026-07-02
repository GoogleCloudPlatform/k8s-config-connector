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

func TestCloudDeployAutomationIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *CloudDeployAutomationIdentity
	}{
		{
			name: "valid external reference",
			ref:  "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline/automations/my-automation",
			want: &CloudDeployAutomationIdentity{
				parent: &AutomationParent{
					ProjectID:          "my-project",
					Location:           "us-central1",
					DeliveryPipelineID: "my-pipeline",
				},
				id: "my-automation",
			},
		},
		{
			name:    "invalid format prefix",
			ref:     "project/my-project/locations/us-central1/deliveryPipelines/my-pipeline/automations/my-automation",
			wantErr: true,
		},
		{
			name:    "invalid format empty",
			ref:     "",
			wantErr: true,
		},
		{
			name:    "invalid short format",
			ref:     "projects/my-project/locations/us-central1/deliveryPipelines/my-pipeline",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &CloudDeployAutomationIdentity{}
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
				if i.parent.DeliveryPipelineID != tt.want.parent.DeliveryPipelineID {
					t.Errorf("DeliveryPipelineID = %v, want %v", i.parent.DeliveryPipelineID, tt.want.parent.DeliveryPipelineID)
				}
				if i.id != tt.want.id {
					t.Errorf("ID = %v, want %v", i.id, tt.want.id)
				}
			}
		})
	}
}
