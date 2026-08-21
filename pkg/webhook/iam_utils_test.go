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

package webhook

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsIAMPolicyMember(t *testing.T) {
	testCases := []struct {
		Name                            string
		Meta                            metav1.TypeMeta
		IsIAMPolicyMemberExpectedResult bool
		IsIAMPolicyExpectedResult       bool
	}{
		{
			Name: "IAMPolicyMemberVersionAndKind",
			Meta: metav1.TypeMeta{
				Kind:       v1beta1.IAMPolicyMemberGVK.Kind,
				APIVersion: v1beta1.IAMPolicyMemberGVK.GroupVersion().String(),
			},
			IsIAMPolicyMemberExpectedResult: true,
			IsIAMPolicyExpectedResult:       false,
		},
		{
			Name: "IAMPolicyVersionAndKind",
			Meta: metav1.TypeMeta{
				Kind:       v1beta1.IAMPolicyGVK.Kind,
				APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
			},
			IsIAMPolicyMemberExpectedResult: false,
			IsIAMPolicyExpectedResult:       true,
		},
		{
			Name: "PubSubVersionAndKind",
			Meta: metav1.TypeMeta{
				Kind:       "PubSubTopic",
				APIVersion: "pubsub.cnrm.cloud.google.com/v1alpha1",
			},
			IsIAMPolicyMemberExpectedResult: false,
			IsIAMPolicyExpectedResult:       false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			u := newUnstructuredFromObject(t, tc.Meta)
			assertIsIAMPolicyMember(t, u, tc.IsIAMPolicyMemberExpectedResult)
			assertIsIAMPolicy(t, u, tc.IsIAMPolicyExpectedResult)
			assertIsIAMResource(t, u, tc.IsIAMPolicyExpectedResult || tc.IsIAMPolicyMemberExpectedResult)
		})
	}
}

func assertIsIAMPolicyMember(t *testing.T, u *unstructured.Unstructured, expectedResult bool) {
	t.Helper()
	result := isIAMPolicyMember(u)
	if result != expectedResult {
		t.Fatalf("unexpected result: got '%v', want '%v'", result, expectedResult)
	}
}

func assertIsIAMPolicy(t *testing.T, u *unstructured.Unstructured, expectedResult bool) {
	t.Helper()
	result := isIAMPolicy(u)
	if result != expectedResult {
		t.Fatalf("unexpected result: got '%v', want '%v'", result, expectedResult)
	}
}

func assertIsIAMResource(t *testing.T, u *unstructured.Unstructured, expectedResult bool) {
	t.Helper()
	result := isIAMResource(u)
	if result != expectedResult {
		t.Fatalf("unexpected result: got '%v', want '%v'", result, expectedResult)
	}
}
