// Copyright 2022 Google LLC
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

package iam_test

import (
	"fmt"
	"reflect"
	"testing"

	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/iam"

	"github.com/google/go-cmp/cmp"
)

// mockIdentityResolver helps to resolve referenced member identity
type mockIdentityResolver struct{}

func (t *mockIdentityResolver) Resolve(member iamv1beta1.Member, memberFrom *iamv1beta1.MemberSource, _ string) (string, error) {
	if member != "" {
		return string(member), nil
	}

	if memberFrom.ServiceAccountRef != nil {
		if memberFrom.ServiceAccountRef.Namespace == "cnrm-foo" && memberFrom.ServiceAccountRef.Name == "cnrm-sa" {
			return "serviceAccount:foo@domain.com", nil
		}
	}
	panic(fmt.Errorf("memberFrom is not mocked"))
}

func TestComputePartialPolicyWithMergedBindings(t *testing.T) {
	condition1 := newIAMCondition("test-iam-condition1")
	condition2 := newIAMCondition("test-iam-condition2")
	tests := []struct {
		name          string
		partialPolicy *iamv1beta1.IAMPartialPolicy
		livePolicy    *iamv1beta1.IAMPolicy
		mergedPolicy  *iamv1beta1.IAMPartialPolicy
	}{
		{
			name: "empty partial policy with empty live policy",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{},
					AllBindings:         []iamv1beta1.IAMPolicyBinding{},
				},
			},
		},
		{
			name: "empty partial policy with non-empty live policy",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
		},
		{
			name: "no existing bindings from live policy",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:foo@example.com",
								},
								{
									Member: "serviceAccount:foo@domain.com",
								},
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:foo@example.com",
								},
								{
									Member: "serviceAccount:foo@domain.com",
								},
							},
						},
					},
				},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
		},
		{
			name: "merge bindings with different {role, condition} tuples",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
								{
									// Resolves to "serviceAccount:foo@domain.com",
									MemberFrom: &iamv1beta1.MemberSource{
										ServiceAccountRef: &iamv1beta1.MemberReference{
											Namespace: "cnrm-foo",
											Name:      "cnrm-sa",
										},
									},
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
					},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
								{
									// Resolves to "serviceAccount:foo@domain.com",
									MemberFrom: &iamv1beta1.MemberSource{
										ServiceAccountRef: &iamv1beta1.MemberReference{
											Namespace: "cnrm-foo",
											Name:      "cnrm-sa",
										},
									},
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
							},
						},
					},
				},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
		},
		{
			name: "merge members with same {role, condition} tuples",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
								{
									Member: "serviceAccount:foo@domain.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user2@example.com",
								},
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
							},
						},
					},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
								{
									Member: "serviceAccount:foo@domain.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user2@example.com",
								},
							},
						},
					},
				},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
					},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"serviceAccount:foo@domain.com",
								"user:foo@example.com",
								"user:user1@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
								"user:user2@example.com",
							},
						},
					},
				},
			},
		},
		{
			name: "merge members from multiple entries with same {role, condition} tuples in the same binding slice",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
								{
									Member: "serviceAccount:foo@domain.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user3@example.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user2@example.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition2,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition2,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user2@example.com",
								},
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition2,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
							},
						},
					},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
								{
									Member: "serviceAccount:foo@domain.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user3@example.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user2@example.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition2,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user1@example.com",
								},
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition2,
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:user2@example.com",
								},
							},
						},
					},
				},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"serviceAccount:foo@domain.com",
								"user:user1@example.com",
								"user:user2@example.com",
								"user:user3@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition2,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"user:user2@example.com",
							},
						},
					},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"serviceAccount:foo@domain.com",
								"user:foo@example.com",
								"user:user1@example.com",
								"user:user2@example.com",
								"user:user3@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition2,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
								"user:user1@example.com",
								"user:user2@example.com",
							},
						},
					},
				},
			},
		},
		{
			name: "empty member arrays",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:foo@example.com",
								},
								{
									Member: "serviceAccount:foo@domain.com",
								},
							},
						},
						{
							Role:    "roles/viewer",
							Members: []iamv1beta1.IAMPartialPolicyMember{},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:foo@example.com",
								},
								{
									Member: "serviceAccount:foo@domain.com",
								},
							},
						},
						{
							Role:    "roles/viewer",
							Members: []iamv1beta1.IAMPartialPolicyMember{},
						},
					},
				},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
		},
		{
			name: "remove partial managed bindings from spec compared to lastAppliedBindings",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:foo@example.com",
								},
							},
						},
					},
				},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{
					Bindings: []iamv1beta1.IAMPartialPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.IAMPartialPolicyMember{
								{
									Member: "user:foo@example.com",
								},
							},
						},
					},
				},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
							},
						},
					},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
					},
				},
			},
		},
		{
			name: "remove all managed bindings from spec compared to lastAppliedBindings",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
							},
						},
						{
							Role: "roles/viewer",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
							},
						},
					},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
							},
						},
					},
				},
			},
		},
		{
			name: "remove all bindings",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:foo@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:bar@example.com",
							},
						},
					},
				},
			},
			mergedPolicy: &iamv1beta1.IAMPartialPolicy{
				Spec: iamv1beta1.IAMPartialPolicySpec{},
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{},
					AllBindings:         []iamv1beta1.IAMPolicyBinding{},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			resolver := mockIdentityResolver{}
			res, err := iam.ComputePartialPolicyWithMergedBindings(tc.partialPolicy, tc.livePolicy, &resolver)
			if err != nil {
				t.Fatalf("error when computing partial policy with merged bindings: %v", err)
			}
			if !reflect.DeepEqual(res, tc.mergedPolicy) {
				t.Fatalf("unexpected merged policy diff (-want +got): \n%v", cmp.Diff(tc.mergedPolicy, res))
			}
		})
	}
}

func TestComputePartialPolicyWithRemainingBindings(t *testing.T) {
	condition1 := newIAMCondition("test-iam-condition1")
	tests := []struct {
		name          string
		partialPolicy *iamv1beta1.IAMPartialPolicy
		livePolicy    *iamv1beta1.IAMPolicy
		remaining     *iamv1beta1.IAMPartialPolicy
	}{
		{
			name: "remove previously managed bindings",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
			remaining: &iamv1beta1.IAMPartialPolicy{
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
		},
		{
			name: "no previously managed bindings to remove",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
			remaining: &iamv1beta1.IAMPartialPolicy{
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{},
					AllBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
								"serviceAccount:foo@domain.com",
							},
						},
					},
				},
			},
		},
		{
			name: "all bindings are removed",
			partialPolicy: &iamv1beta1.IAMPartialPolicy{
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						// roles that are removed out of band
						// and KCC hasn't drift-corrected yet before deletion.
						{
							Role: "roles/viewer",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
					},
				},
			},
			livePolicy: &iamv1beta1.IAMPolicy{
				Spec: iamv1beta1.IAMPolicySpec{
					Bindings: []iamv1beta1.IAMPolicyBinding{
						{
							Role: "roles/editor",
							Members: []iamv1beta1.Member{
								"user:user2@example.com",
							},
						},
						{
							Role:      "roles/editor",
							Condition: condition1,
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
						{
							Role: "roles/owner",
							Members: []iamv1beta1.Member{
								"user:user1@example.com",
							},
						},
					},
				},
			},
			remaining: &iamv1beta1.IAMPartialPolicy{
				Status: iamv1beta1.IAMPartialPolicyStatus{
					LastAppliedBindings: []iamv1beta1.IAMPolicyBinding{},
					AllBindings:         []iamv1beta1.IAMPolicyBinding{},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := iam.ComputePartialPolicyWithRemainingBindings(tc.partialPolicy, tc.livePolicy)
			if !reflect.DeepEqual(res, tc.remaining) {
				t.Fatalf("unexpected merged policy diff (-want +got): \n%v", cmp.Diff(tc.remaining, res))
			}
		})
	}
}

func newIAMCondition(title string) *iamv1beta1.IAMCondition {
	return &iamv1beta1.IAMCondition{
		Title:       title,
		Description: "Test IAM Condition",
		Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
	}
}
