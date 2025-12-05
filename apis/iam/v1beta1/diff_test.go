// Copyright 2025 Google LLC
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
	"sort"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"github.com/google/go-cmp/cmp"
)

func TestIAMPolicySpecDiffers(t *testing.T) {
	tests := []struct {
		name    string
		desired *IAMPolicySpec
		actual  *IAMPolicySpec
		want    *structuredreporting.Diff
	}{
		{
			name: "no diff",
			desired: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
				},
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType:         "DATA_READ",
								ExemptedMembers: []Member{"user:exempt@example.com"},
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
				},
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType:         "DATA_READ",
								ExemptedMembers: []Member{"user:exempt@example.com"},
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{},
		},
		{
			name: "resource reference kind differs",
			desired: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
			},
			actual: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Folder",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.resourceRef.kind",
						Old: "Project",
						New: "Folder",
					},
				},
			},
		},
		{
			name: "resource reference namespace differs",
			desired: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
			},
			actual: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "other-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.resourceRef.namespace",
						Old: "test-namespace",
						New: "other-namespace",
					},
				},
			},
		},
		{
			name: "resource reference name differs",
			desired: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
			},
			actual: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "other-project",
					External:  "projects/test-project",
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.resourceRef.name",
						Old: "test-project",
						New: "other-project",
					},
				},
			},
		},
		{
			name: "resource reference external differs",
			desired: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
			},
			actual: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/other-project",
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.resourceRef.external",
						Old: "projects/test-project",
						New: "projects/other-project",
					},
				},
			},
		},
		{
			name: "binding added",
			desired: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
					{
						Role:    "roles/editor",
						Members: []Member{"user:editor@example.com"},
					},
				},
			},
			actual: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.bindings[role=roles/editor]",
						Old: "present in desired spec",
						New: "absent in actual spec",
					},
				},
			},
		},
		{
			name: "binding removed",
			desired: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
				},
			},
			actual: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
					{
						Role:    "roles/editor",
						Members: []Member{"user:editor@example.com"},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.bindings[role=roles/editor]",
						Old: "absent in desired spec",
						New: "present in actual spec",
					},
				},
			},
		},
		{
			name: "binding members differ",
			desired: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com", "user:other@example.com"},
					},
				},
			},
			actual: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID: "spec.bindings[role=roles/viewer]",
						Old: IAMPolicyBinding{
							Role:    "roles/viewer",
							Members: []Member{"user:other@example.com", "user:viewer@example.com"},
						},
						New: IAMPolicyBinding{
							Role:    "roles/viewer",
							Members: []Member{"user:viewer@example.com"},
						},
					},
				},
			},
		},
		{
			name: "binding members order differs - no diff",
			desired: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com", "user:other@example.com"},
					},
				},
			},
			actual: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:other@example.com", "user:viewer@example.com"},
					},
				},
			},
			want: &structuredreporting.Diff{},
		},
		{
			name: "binding condition differs",
			desired: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
						Condition: &IAMCondition{
							Title:      "title",
							Expression: "expression",
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
						Condition: &IAMCondition{
							Title:      "other-title",
							Expression: "expression",
						},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID: "spec.bindings[role=roles/viewer]",
						Old: IAMPolicyBinding{
							Role:    "roles/viewer",
							Members: []Member{"user:viewer@example.com"},
							Condition: &IAMCondition{
								Title:      "title",
								Expression: "expression",
							},
						},
						New: IAMPolicyBinding{
							Role:    "roles/viewer",
							Members: []Member{"user:viewer@example.com"},
							Condition: &IAMCondition{
								Title:      "other-title",
								Expression: "expression",
							},
						},
					},
				},
			}},
		{
			name: "empty desired bindings",
			desired: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{},
			},
			actual: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.bindings[role=roles/viewer]",
						Old: "absent in desired spec",
						New: "present in actual spec",
					},
				},
			},
		},
		{
			name: "empty actual bindings",
			desired: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
				},
			},
			actual: &IAMPolicySpec{
				Bindings: []IAMPolicyBinding{},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.bindings[role=roles/viewer]",
						Old: "present in desired spec",
						New: "absent in actual spec",
					},
				},
			},
		},
		{
			name: "audit config added",
			desired: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
						},
					},
					{
						Service: "compute.googleapis.com",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_WRITE",
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.auditConfigs[service=compute.googleapis.com]",
						Old: "present in desired spec",
						New: "absent in actual spec",
					},
				},
			},
		},
		{
			name: "audit config removed",
			desired: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
						},
					},
					{
						Service: "compute.googleapis.com",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_WRITE",
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.auditConfigs[service=compute.googleapis.com]",
						Old: "absent in desired spec",
						New: "present in actual spec",
					},
				},
			}},
		{
			name: "audit config log type differs",
			desired: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_WRITE",
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID: "spec.auditConfigs[service=allServices]",
						Old: IAMPolicyAuditConfig{
							Service: "allServices",
							AuditLogConfigs: []AuditLogConfig{
								{
									LogType: "DATA_READ",
								},
							},
						},
						New: IAMPolicyAuditConfig{
							Service: "allServices",
							AuditLogConfigs: []AuditLogConfig{
								{
									LogType: "DATA_WRITE",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "audit config exempted members differs",
			desired: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType:         "DATA_READ",
								ExemptedMembers: []Member{"user:exempt1@example.com"},
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType:         "DATA_READ",
								ExemptedMembers: []Member{"user:exempt2@example.com"},
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID: "spec.auditConfigs[service=allServices]",
						Old: IAMPolicyAuditConfig{
							Service: "allServices",
							AuditLogConfigs: []AuditLogConfig{
								{
									LogType:         "DATA_READ",
									ExemptedMembers: []Member{"user:exempt1@example.com"},
								},
							},
						},
						New: IAMPolicyAuditConfig{
							Service: "allServices",
							AuditLogConfigs: []AuditLogConfig{
								{
									LogType:         "DATA_READ",
									ExemptedMembers: []Member{"user:exempt2@example.com"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "audit config exempted members order differs - no diff",
			desired: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType:         "DATA_READ",
								ExemptedMembers: []Member{"user:exempt1@example.com", "user:exempt2@example.com"},
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType:         "DATA_READ",
								ExemptedMembers: []Member{"user:exempt2@example.com", "user:exempt1@example.com"},
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{},
		},
		{
			name: "audit config log configs order differs - no diff",
			desired: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
							{
								LogType: "DATA_WRITE",
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_WRITE",
							},
							{
								LogType: "DATA_READ",
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{},
		},
		{
			name: "empty desired audit configs",
			desired: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{},
			},
			actual: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.auditConfigs[service=allServices]",
						Old: "absent in desired spec",
						New: "present in actual spec",
					},
				},
			},
		},
		{
			name: "empty actual audit configs",
			desired: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				AuditConfigs: []IAMPolicyAuditConfig{},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.auditConfigs[service=allServices]",
						Old: "present in desired spec",
						New: "absent in actual spec",
					},
				},
			},
		},
		{
			name: "combined diff",
			desired: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Project",
					Namespace: "test-namespace",
					Name:      "test-project",
					External:  "projects/test-project",
				},
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer@example.com"},
					},
					{
						Role:    "roles/editor",
						Members: []Member{"user:editor@example.com"},
					},
				},
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "allServices",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_READ",
							},
						},
					},
				},
			},
			actual: &IAMPolicySpec{
				ResourceReference: ResourceReference{
					Kind:      "Folder",
					Namespace: "other-namespace",
					Name:      "other-project",
					External:  "projects/other-project",
				},
				Bindings: []IAMPolicyBinding{
					{
						Role:    "roles/viewer",
						Members: []Member{"user:viewer-actual@example.com"},
					},
				},
				AuditConfigs: []IAMPolicyAuditConfig{
					{
						Service: "compute.googleapis.com",
						AuditLogConfigs: []AuditLogConfig{
							{
								LogType: "DATA_WRITE",
							},
						},
					},
				},
			},
			want: &structuredreporting.Diff{
				Fields: []structuredreporting.DiffField{
					{
						ID:  "spec.resourceRef.kind",
						Old: "Project",
						New: "Folder",
					},
					{
						ID:  "spec.resourceRef.namespace",
						Old: "test-namespace",
						New: "other-namespace",
					},
					{
						ID:  "spec.resourceRef.name",
						Old: "test-project",
						New: "other-project",
					},
					{
						ID:  "spec.resourceRef.external",
						Old: "projects/test-project",
						New: "projects/other-project",
					},
					{
						ID:  "spec.bindings[role=roles/editor]",
						Old: "present in desired spec",
						New: "absent in actual spec",
					},
					{
						ID: "spec.bindings[role=roles/viewer]",
						Old: IAMPolicyBinding{
							Role:    "roles/viewer",
							Members: []Member{"user:viewer@example.com"},
						},
						New: IAMPolicyBinding{
							Role:    "roles/viewer",
							Members: []Member{"user:viewer-actual@example.com"},
						},
					},
					{
						ID:  "spec.auditConfigs[service=allServices]",
						Old: "present in desired spec",
						New: "absent in actual spec",
					},
					{
						ID:  "spec.auditConfigs[service=compute.googleapis.com]",
						Old: "absent in desired spec",
						New: "present in actual spec",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a deep copy of the original desired spec for later comparison
			originalDesired := tc.desired.DeepCopy()

			got := IAMPolicySpecDiffers(tc.desired, tc.actual)

			if diff := cmp.Diff(originalDesired, tc.desired); diff != "" {
				t.Errorf("IAMPolicySpecDiffers() mutated the desired spec (-original +mutated):\n%v", diff)
			}

			if diff := cmp.Diff(sortDiffFields(tc.want), sortDiffFields(got)); diff != "" {
				t.Errorf("IAMPolicySpecDiffers() got diff (-want +got):\n%v", diff)
			}
		})
	}
}

// sortDiffFields sorts the fields of a structuredreporting.Diff by their ID for stable comparison.
func sortDiffFields(diff *structuredreporting.Diff) *structuredreporting.Diff {
	if diff == nil || diff.Fields == nil {
		return diff
	}
	sort.Slice(diff.Fields, func(i, j int) bool {
		return diff.Fields[i].ID < diff.Fields[j].ID
	})
	return diff
}
