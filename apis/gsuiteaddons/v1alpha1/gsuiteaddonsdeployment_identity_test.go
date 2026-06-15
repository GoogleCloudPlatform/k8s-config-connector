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

func TestGSuiteAddonsDeploymentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *GSuiteAddonsDeploymentIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/deployments/my-deployment",
			want: &GSuiteAddonsDeploymentIdentity{
				Project:    "my-project",
				Deployment: "my-deployment",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://gsuiteaddons.googleapis.com/projects/my-project/deployments/my-deployment",
			want: &GSuiteAddonsDeploymentIdentity{
				Project:    "my-project",
				Deployment: "my-deployment",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &GSuiteAddonsDeploymentIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Deployment != tt.want.Deployment {
					t.Errorf("Deployment = %v, want %v", i.Deployment, tt.want.Deployment)
				}
			}
		})
	}
}
