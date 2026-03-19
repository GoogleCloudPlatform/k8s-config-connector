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

func TestDataprocAutoscalingPolicyRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantErr  bool
	}{
		{
			name:     "valid reference",
			external: "projects/my-project/regions/us-central1/autoscalingPolicies/my-policy",
			wantErr:  false,
		},
		{
			name:     "invalid reference format",
			external: "invalid/format",
			wantErr:  true,
		},
		{
			name:     "missing project",
			external: "regions/us-central1/autoscalingPolicies/my-policy",
			wantErr:  true,
		},
		{
			name:     "missing region",
			external: "projects/my-project/autoscalingPolicies/my-policy",
			wantErr:  true,
		},
		{
			name:     "missing policy",
			external: "projects/my-project/regions/us-central1",
			wantErr:  true,
		},
		{
			name:     "empty reference",
			external: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DataprocAutoscalingPolicyRef{}
			err := r.ValidateExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataprocAutoscalingPolicyIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DataprocAutoscalingPolicyIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/regions/us-central1/autoscalingPolicies/my-policy",
			want: &DataprocAutoscalingPolicyIdentity{
				Project:           "my-project",
				Region:            "us-central1",
				AutoscalingPolicy: "my-policy",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://dataproc.googleapis.com/projects/my-project/regions/us-central1/autoscalingPolicies/my-policy",
			want: &DataprocAutoscalingPolicyIdentity{
				Project:           "my-project",
				Region:            "us-central1",
				AutoscalingPolicy: "my-policy",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DataprocAutoscalingPolicyIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Region != tt.want.Region {
					t.Errorf("Region = %v, want %v", i.Region, tt.want.Region)
				}
				if i.AutoscalingPolicy != tt.want.AutoscalingPolicy {
					t.Errorf("AutoscalingPolicy = %v, want %v", i.AutoscalingPolicy, tt.want.AutoscalingPolicy)
				}
			}
		})
	}
}
