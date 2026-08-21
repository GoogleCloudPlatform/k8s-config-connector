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

package policy

import (
	"testing"

	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	condition "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsAPIServerUpdateRequired(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		policy         *iamv1beta1.IAMPolicy
		expectedResult bool
	}{
		{
			name: "no previous conditions and observed generation",
			policy: &iamv1beta1.IAMPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 1,
				},
				Status: iamv1beta1.IAMPolicyStatus{},
			},
			expectedResult: true,
		},
		{
			name: "conditions are update to date, no observed generation",
			policy: &iamv1beta1.IAMPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 1,
				},
				Status: iamv1beta1.IAMPolicyStatus{
					Conditions: []condition.Condition{
						k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage),
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "conditions are update to date, observed generation is stale",
			policy: &iamv1beta1.IAMPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 2,
				},
				Status: iamv1beta1.IAMPolicyStatus{
					Conditions: []condition.Condition{
						k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage),
					},
					ObservedGeneration: 1,
				},
			},
			expectedResult: true,
		},
		{
			name: "conditions are update to date, observed generation matches with the generation",
			policy: &iamv1beta1.IAMPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Generation: 2,
				},
				Status: iamv1beta1.IAMPolicyStatus{
					Conditions: []condition.Condition{
						k8s.NewCustomReadyCondition(corev1.ConditionTrue, k8s.UpToDate, k8s.UpToDateMessage),
					},
					ObservedGeneration: 2,
				},
			},
			expectedResult: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := isAPIServerUpdateRequired(tc.policy)
			if actual != tc.expectedResult {
				t.Fatalf("got %v, want %v", actual, tc.expectedResult)
			}
		})
	}
}
