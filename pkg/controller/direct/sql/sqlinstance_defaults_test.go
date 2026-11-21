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

package sql

import (
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
)

func TestApplySQLInstanceGCPDefaults_DiskType(t *testing.T) {
	tests := []struct {
		name         string
		in           *krm.SQLInstance
		actual       *api.DatabaseInstance
		wantDiskType string
	}{
		{
			name: "enterprise edition, diskType unset -> defaults to PD_SSD",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Region: direct.LazyPtr("us-central1"),
					Settings: krm.InstanceSettings{
						Edition: direct.LazyPtr("ENTERPRISE"),
					},
				},
			},
			wantDiskType: "PD_SSD",
		},
		{
			name: "edition unset (defaults to Enterprise), diskType unset -> defaults to PD_SSD",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Region: direct.LazyPtr("us-central1"),
				},
			},
			wantDiskType: "PD_SSD",
		},
		{
			name: "enterprise plus edition, diskType unset -> server chooses default (empty string)",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Region: direct.LazyPtr("us-central1"),
					Settings: krm.InstanceSettings{
						Edition: direct.LazyPtr("ENTERPRISE_PLUS"),
					},
				},
			},
			wantDiskType: "",
		},
		{
			name: "enterprise plus edition, explicit HYPERDISK_BALANCED -> preserved",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Region: direct.LazyPtr("us-central1"),
					Settings: krm.InstanceSettings{
						Edition:  direct.LazyPtr("ENTERPRISE_PLUS"),
						DiskType: direct.LazyPtr("HYPERDISK_BALANCED"),
					},
				},
			},
			wantDiskType: "HYPERDISK_BALANCED",
		},
		{
			name: "enterprise plus edition, explicit PD_SSD -> preserved",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Region: direct.LazyPtr("us-central1"),
					Settings: krm.InstanceSettings{
						Edition:  direct.LazyPtr("ENTERPRISE_PLUS"),
						DiskType: direct.LazyPtr("PD_SSD"),
					},
				},
			},
			wantDiskType: "PD_SSD",
		},
		{
			name: "existing enterprise plus instance with HYPERDISK_BALANCED, diskType unset -> preserves actual",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Region: direct.LazyPtr("us-central1"),
					Settings: krm.InstanceSettings{
						Edition: direct.LazyPtr("ENTERPRISE_PLUS"),
					},
				},
			},
			actual: &api.DatabaseInstance{
				Settings: &api.Settings{
					DataDiskType: "HYPERDISK_BALANCED",
				},
			},
			wantDiskType: "HYPERDISK_BALANCED",
		},
		{
			name: "existing standard instance with PD_HDD, diskType unset -> preserves actual",
			in: &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					Region: direct.LazyPtr("us-central1"),
					Settings: krm.InstanceSettings{
						Edition: direct.LazyPtr("ENTERPRISE"),
					},
				},
			},
			actual: &api.DatabaseInstance{
				Settings: &api.Settings{
					DataDiskType: "PD_HDD",
				},
			},
			wantDiskType: "PD_HDD",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := &api.DatabaseInstance{
				Settings: &api.Settings{
					DataDiskType: direct.ValueOf(tc.in.Spec.Settings.DiskType),
				},
			}
			ApplySQLInstanceGCPDefaults(tc.in, out, tc.actual, nil)
			if out.Settings.DataDiskType != tc.wantDiskType {
				t.Errorf("DataDiskType = %q, want %q", out.Settings.DataDiskType, tc.wantDiskType)
			}
		})
	}
}

func TestDiffInstances_HyperdiskBalanced(t *testing.T) {
	desired := &api.DatabaseInstance{
		Settings: &api.Settings{
			DataDiskType: "HYPERDISK_BALANCED",
		},
	}
	actual := &api.DatabaseInstance{
		Settings: &api.Settings{
			DataDiskType: "HYPERDISK_BALANCED",
		},
	}

	diff := DiffInstances(desired, actual)
	if diff.HasDiff() {
		t.Errorf("DiffInstances() unexpected diff for HYPERDISK_BALANCED: %v", diff.Fields)
	}

	actualMismatch := &api.DatabaseInstance{
		Settings: &api.Settings{
			DataDiskType: "PD_SSD",
		},
	}
	diffMismatch := DiffInstances(desired, actualMismatch)
	if !diffMismatch.HasDiff() {
		t.Errorf("DiffInstances() expected diff between HYPERDISK_BALANCED and PD_SSD, got none")
	}
}
