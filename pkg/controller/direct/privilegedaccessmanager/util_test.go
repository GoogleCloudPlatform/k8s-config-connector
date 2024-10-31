// Copyright 2024 Google LLC
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

package privilegedaccessmanager

import (
	"reflect"
	"testing"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1alpha1"

	"github.com/google/go-cmp/cmp"
)

func TestSortAccessControlEntrySlice(t *testing.T) {
	tests := []struct {
		testName string
		unsorted []krm.AccessControlEntry
		expected []krm.AccessControlEntry
	}{
		{
			testName: "nil input",
			unsorted: nil,
			expected: nil,
		},
		{
			testName: "empty slice",
			unsorted: []krm.AccessControlEntry{},
			expected: []krm.AccessControlEntry{},
		},
		{
			testName: "empty principal slice in single access control entry",
			unsorted: []krm.AccessControlEntry{
				{
					Principals: []string{},
				},
			},
			expected: []krm.AccessControlEntry{
				{
					Principals: []string{},
				},
			},
		},
		{
			testName: "sort single entry",
			unsorted: []krm.AccessControlEntry{
				{
					Principals: []string{
						"user:abc@test.com",
						"serviceAccount:xyz@gservicaccount.com",
					},
				},
			},
			expected: []krm.AccessControlEntry{
				{
					Principals: []string{
						"serviceAccount:xyz@gservicaccount.com",
						"user:abc@test.com",
					},
				},
			},
		},
		{
			testName: "sort single entry containing duplicate principals",
			unsorted: []krm.AccessControlEntry{
				{
					Principals: []string{
						"user:abc@test.com",
						"serviceAccount:xyz@gservicaccount.com",
						"user:abc@test.com",
					},
				},
			},
			expected: []krm.AccessControlEntry{
				{
					Principals: []string{
						"serviceAccount:xyz@gservicaccount.com",
						"user:abc@test.com",
						"user:abc@test.com",
					},
				},
			},
		},
		{
			testName: "sort multiple entries",
			unsorted: []krm.AccessControlEntry{
				{
					Principals: []string{
						"user:abc@test.com",
						"serviceAccount:xyz@gservicaccount.com",
					},
				},
				{
					Principals: []string{
						"user:abc@test.com",
					},
				},
				{
					Principals: []string{
						"user:abc@test.com",
						"serviceAccount:xyz2@gservicaccount.com",
					},
				},
				{
					Principals: []string{
						"user:abc@test.com",
						"serviceAccount:xyz@gservicaccount.com",
						"serviceAccount:xyz2@gservicaccount.com",
					},
				},
			},
			expected: []krm.AccessControlEntry{
				{
					Principals: []string{
						"user:abc@test.com",
					},
				},
				{
					Principals: []string{
						"serviceAccount:xyz2@gservicaccount.com",
						"user:abc@test.com",
					},
				},
				{
					Principals: []string{
						"serviceAccount:xyz@gservicaccount.com",
						"user:abc@test.com",
					},
				},
				{
					Principals: []string{
						"serviceAccount:xyz2@gservicaccount.com",
						"serviceAccount:xyz@gservicaccount.com",
						"user:abc@test.com",
					},
				},
			},
		},
		{
			testName: "sort multiple entries",
			unsorted: []krm.AccessControlEntry{
				{
					Principals: []string{
						"user:abc@test.com",
						"serviceAccount:xyz@gservicaccount.com",
					},
				},
				{
					Principals: []string{
						"user:abc@test.com",
					},
				},
				{
					Principals: []string{
						"user:abc@test.com",
						"serviceAccount:xyz2@gservicaccount.com",
					},
				},
				{
					Principals: []string{
						"user:abc@test.com",
						"serviceAccount:xyz@gservicaccount.com",
						"serviceAccount:xyz2@gservicaccount.com",
					},
				},
			},
			expected: []krm.AccessControlEntry{
				{
					Principals: []string{
						"user:abc@test.com",
					},
				},
				{
					Principals: []string{
						"serviceAccount:xyz2@gservicaccount.com",
						"user:abc@test.com",
					},
				},
				{
					Principals: []string{
						"serviceAccount:xyz@gservicaccount.com",
						"user:abc@test.com",
					},
				},
				{
					Principals: []string{
						"serviceAccount:xyz2@gservicaccount.com",
						"serviceAccount:xyz@gservicaccount.com",
						"user:abc@test.com",
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			sortAccessControlEntrySlice(tc.unsorted)
			if !reflect.DeepEqual(tc.unsorted, tc.expected) {
				t.Fatalf("unexpected diff running sortAccessControlEntrySlice (-want +got): \n%v", cmp.Diff(tc.expected, tc.unsorted))
			}
		})
	}
}

func TestSortArrayFieldsInSpec(t *testing.T) {
	tests := []struct {
		testName string
		unsorted *krm.PrivilegedAccessManagerEntitlementSpec
		expected *krm.PrivilegedAccessManagerEntitlementSpec
	}{
		{
			testName: "nil spec",
			unsorted: nil,
			expected: nil,
		},
		{
			testName: "empty spec with no EligibleUsers nor ApprovalWorkflow",
			unsorted: &krm.PrivilegedAccessManagerEntitlementSpec{},
			expected: &krm.PrivilegedAccessManagerEntitlementSpec{},
		},
		{
			testName: "nil ApprovalWorkflow.ManualApprovals",
			unsorted: &krm.PrivilegedAccessManagerEntitlementSpec{
				EligibleUsers: []krm.AccessControlEntry{
					{
						Principals: []string{
							"user:abc@test.com",
							"serviceAccount:xyz@gservicaccount.com",
						},
					},
				},
				ApprovalWorkflow: &krm.ApprovalWorkflow{},
			},
			expected: &krm.PrivilegedAccessManagerEntitlementSpec{
				EligibleUsers: []krm.AccessControlEntry{
					{
						Principals: []string{
							"serviceAccount:xyz@gservicaccount.com",
							"user:abc@test.com",
						},
					},
				},
				ApprovalWorkflow: &krm.ApprovalWorkflow{},
			},
		},
		{
			testName: "nil ApprovalWorkflow.ManualApprovals.Steps",
			unsorted: &krm.PrivilegedAccessManagerEntitlementSpec{
				EligibleUsers: []krm.AccessControlEntry{
					{
						Principals: []string{
							"user:abc@test.com",
							"serviceAccount:xyz@gservicaccount.com",
						},
					},
				},
				ApprovalWorkflow: &krm.ApprovalWorkflow{
					ManualApprovals: &krm.ManualApprovals{},
				},
			},
			expected: &krm.PrivilegedAccessManagerEntitlementSpec{
				EligibleUsers: []krm.AccessControlEntry{
					{
						Principals: []string{
							"serviceAccount:xyz@gservicaccount.com",
							"user:abc@test.com",
						},
					},
				},
				ApprovalWorkflow: &krm.ApprovalWorkflow{
					ManualApprovals: &krm.ManualApprovals{},
				},
			},
		},
		{
			testName: "EligibleUsers and ApprovalWorkflow.ManualApprovals.Steps.Approvers are sorted",
			unsorted: &krm.PrivilegedAccessManagerEntitlementSpec{
				EligibleUsers: []krm.AccessControlEntry{
					{
						Principals: []string{
							"user:abc@test.com",
							"serviceAccount:xyz@gservicaccount.com",
						},
					},
					{
						Principals: []string{
							"user:abc@test.com",
						},
					},
					{
						Principals: []string{
							"user:abc@test.com",
							"serviceAccount:xyz2@gservicaccount.com",
						},
					},
				},
				ApprovalWorkflow: &krm.ApprovalWorkflow{
					ManualApprovals: &krm.ManualApprovals{
						Steps: []krm.Step{
							{
								Approvers: []krm.AccessControlEntry{
									{
										Principals: []string{
											"user:abc@test.com",
											"serviceAccount:xyz@gservicaccount.com",
										},
									},
									{
										Principals: []string{
											"user:abc@test.com",
										},
									},
									{
										Principals: []string{
											"user:abc@test.com",
											"serviceAccount:xyz2@gservicaccount.com",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: &krm.PrivilegedAccessManagerEntitlementSpec{
				EligibleUsers: []krm.AccessControlEntry{
					{
						Principals: []string{
							"user:abc@test.com",
						},
					},
					{
						Principals: []string{
							"serviceAccount:xyz2@gservicaccount.com",
							"user:abc@test.com",
						},
					},
					{
						Principals: []string{
							"serviceAccount:xyz@gservicaccount.com",
							"user:abc@test.com",
						},
					},
				},
				ApprovalWorkflow: &krm.ApprovalWorkflow{
					ManualApprovals: &krm.ManualApprovals{
						Steps: []krm.Step{
							{
								Approvers: []krm.AccessControlEntry{
									{
										Principals: []string{
											"user:abc@test.com",
										},
									},
									{
										Principals: []string{
											"serviceAccount:xyz2@gservicaccount.com",
											"user:abc@test.com",
										},
									},
									{
										Principals: []string{
											"serviceAccount:xyz@gservicaccount.com",
											"user:abc@test.com",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			sortArrayFieldsInSpec(tc.unsorted)
			if !reflect.DeepEqual(tc.unsorted, tc.expected) {
				t.Fatalf("unexpected diff running sortArrayFieldsInSpec (-want +got): \n%v", cmp.Diff(tc.expected, tc.unsorted))
			}
		})
	}
}
