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

package v1alpha1_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
)

func TestContentWarehouseRuleSetIdentity_FromExternal(t *testing.T) {
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
			ref:          "projects/my-project/locations/us-central1/ruleSets/my-ruleset",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-ruleset",
			wantErr:      false,
		},
		{
			name:         "full url",
			ref:          "https://contentwarehouse.googleapis.com/projects/my-project/locations/us-central1/ruleSets/my-ruleset",
			wantProject:  "my-project",
			wantLocation: "us-central1",
			wantID:       "my-ruleset",
			wantErr:      false,
		},
		{
			name:    "invalid external ref",
			ref:     "projects/my-project/ruleSets/my-ruleset",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &v1alpha1.ContentWarehouseRuleSetIdentity{}
			if err := i.FromExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("ContentWarehouseRuleSetIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Project != tt.wantProject {
					t.Errorf("ContentWarehouseRuleSetIdentity.FromExternal() Project = %v, want %v", i.Project, tt.wantProject)
				}
				if i.Location != tt.wantLocation {
					t.Errorf("ContentWarehouseRuleSetIdentity.FromExternal() Location = %v, want %v", i.Location, tt.wantLocation)
				}
				if i.RuleSet != tt.wantID {
					t.Errorf("ContentWarehouseRuleSetIdentity.FromExternal() RuleSet = %v, want %v", i.RuleSet, tt.wantID)
				}
				if got := i.String(); got != "projects/"+tt.wantProject+"/locations/"+tt.wantLocation+"/ruleSets/"+tt.wantID {
					t.Errorf("ContentWarehouseRuleSetIdentity.String() = %v", got)
				}
			}
		})
	}
}
