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

package iamclient

import (
	"context"
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"

	dcliam "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These variables need to be initialized here because the DCL structs expect
// string pointers.
var (
	emptyEtag        = ""
	testEtag         = "ACAB"
	testRole1        = "roles/computeAdmin"
	testRole2        = "roles/viewer"
	testTitle1       = "FooBar"
	testTitle2       = "BarBaz"
	testDescription1 = "test description"
	testDescription2 = "second description"
	testExpression1  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
	testExpression2  = "true"
	testMember1      = "foo@bar.iam.gserviceaccount.com"
)

func TestDCLPolicyFromKRMPolicy(t *testing.T) {
	tests := []struct {
		name           string
		dclResource    *dclunstruct.Resource
		krmPolicy      *v1beta1.IAMPolicy
		expectedResult *dcliam.Policy
		hasError       bool
	}{
		{
			name: "iam policy no bindings, empty etag",
			krmPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "DataprocCluster",
					},
					Bindings: []v1beta1.IAMPolicyBinding{},
					Etag:     "",
				},
			},
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			expectedResult: &dcliam.Policy{
				Bindings: []dcliam.Binding{},
				Etag:     &emptyEtag,
				Version:  nil,
				Resource: &dclunstruct.Resource{
					STV: dclunstruct.ServiceTypeVersion{
						Service: "dataproc",
						Type:    "Cluster",
						Version: "ga",
					},
				},
			},
		},
		{
			name: "iam policy with bindings no conditions",
			krmPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{},
					Bindings: []v1beta1.IAMPolicyBinding{
						{
							Members: []v1beta1.Member{
								"foo@bar.iam.gserviceaccount.com",
							},
							Role: testRole1,
						},
					},
					Etag: "",
				},
			},
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			expectedResult: &dcliam.Policy{
				Bindings: []dcliam.Binding{
					{
						Role: &testRole1,
						Members: []string{
							"foo@bar.iam.gserviceaccount.com",
						},
					},
				},
				Etag:    &emptyEtag,
				Version: nil,
				Resource: &dclunstruct.Resource{
					STV: dclunstruct.ServiceTypeVersion{
						Service: "dataproc",
						Type:    "Cluster",
						Version: "ga",
					},
				},
			},
		},
		{
			name: "iam policy with single binding with conditions",
			krmPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{},
					Bindings: []v1beta1.IAMPolicyBinding{
						{
							Members: []v1beta1.Member{
								"foo@bar.iam.gserviceaccount.com",
							},
							Role: testRole1,
							Condition: &v1beta1.IAMCondition{
								Title:       testTitle1,
								Description: testDescription1,
								Expression:  testExpression1,
							},
						},
					},
				},
			},
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			expectedResult: &dcliam.Policy{
				Bindings: []dcliam.Binding{
					{
						Role: &testRole1,
						Members: []string{
							"foo@bar.iam.gserviceaccount.com",
						},
						Condition: &dcliam.Condition{
							Title:       &testTitle1,
							Description: &testDescription1,
							Expression:  &testExpression1,
						},
					},
				},
				Etag: &emptyEtag,
				Resource: &dclunstruct.Resource{
					STV: dclunstruct.ServiceTypeVersion{
						Service: "dataproc",
						Type:    "Cluster",
						Version: "ga",
					},
				},
			},
		},
		{
			name: "iam policy with multiple bindings/member/conditions",
			krmPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{},
					Bindings: []v1beta1.IAMPolicyBinding{
						{
							Members: []v1beta1.Member{
								"foo@bar.iam.gserviceaccount.com",
								"user@test.iam.gserviceaccount.com",
							},
							Role: testRole1,
							Condition: &v1beta1.IAMCondition{
								Title:       testTitle1,
								Description: testDescription1,
								Expression:  testExpression1,
							},
						},
						{
							Role: "roles/viewer",
							Condition: &v1beta1.IAMCondition{
								Title:       testTitle2,
								Description: testDescription2,
								Expression:  testExpression2,
							},
						},
					},
				},
			},
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			expectedResult: &dcliam.Policy{
				Bindings: []dcliam.Binding{
					{
						Role: &testRole1,
						Members: []string{
							"foo@bar.iam.gserviceaccount.com",
							"user@test.iam.gserviceaccount.com",
						},
						Condition: &dcliam.Condition{
							Title:       &testTitle1,
							Description: &testDescription1,
							Expression:  &testExpression1,
						},
					},
					{
						Role: &testRole2,
						Condition: &dcliam.Condition{
							Title:       &testTitle2,
							Description: &testDescription2,
							Expression:  &testExpression2,
						},
					},
				},
				Etag: &emptyEtag,
				Resource: &dclunstruct.Resource{
					STV: dclunstruct.ServiceTypeVersion{
						Service: "dataproc",
						Type:    "Cluster",
						Version: "ga",
					},
				},
			},
		},
		{
			name: "iam policy with audit config",
			krmPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{},
					AuditConfigs: []v1beta1.IAMPolicyAuditConfig{
						{
							Service: "allServices",
							AuditLogConfigs: []v1beta1.AuditLogConfig{
								{
									LogType: "DATA_READ",
									ExemptedMembers: []v1beta1.Member{
										"foo@bar.iam.gserviceaccount.com",
									},
								},
							},
						},
					},
				},
			},
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			expectedResult: &dcliam.Policy{},
			hasError:       true,
		},
		{
			name: "iam policy with etag set by partial policy controller",
			krmPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "DataprocCluster",
					},
					Bindings: []v1beta1.IAMPolicyBinding{},
					Etag:     testEtag,
				},
			},
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			expectedResult: &dcliam.Policy{
				Bindings: []dcliam.Binding{},
				Etag:     &testEtag,
				Resource: &dclunstruct.Resource{
					STV: dclunstruct.ServiceTypeVersion{
						Service: "dataproc",
						Type:    "Cluster",
						Version: "ga",
					},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := newDCLPolicyFromKRMPolicy(tc.krmPolicy, tc.dclResource)
			if tc.hasError {
				if err == nil {
					t.Fatalf("expected to get error when converting IAMPolicy from DCLPolicyResource")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tc.expectedResult) {
				t.Fatalf("got %v, want %v", actual, tc.expectedResult)
			}
		})
	}
}

func TestIAMPolicyFromDCLResource(t *testing.T) {
	tests := []struct {
		name           string
		dclResource    *dclunstruct.Resource
		origPolicy     *v1beta1.IAMPolicy
		expectedResult *v1beta1.IAMPolicy
		hasError       bool
	}{
		{
			name: "iam policy no bindings",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "iam",
					Type:    "Policy",
					Version: "ga",
				},
				Object: map[string]interface{}{
					"etag": testEtag,
				},
			},
			origPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
				},
			},
			expectedResult: &v1beta1.IAMPolicy{
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
					Bindings:     nil,
					AuditConfigs: nil,
					Etag:         testEtag,
				},
			},
		},
		{
			name: "iam policy with binding no conditions",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "iam",
					Type:    "Policy",
					Version: "ga",
				},
				Object: map[string]interface{}{
					"etag": testEtag,
					"bindings": []interface{}{
						map[string]interface{}{
							"members": []interface{}{
								"foo@bar.iam.gserviceaccount.com",
							},
							"role": testRole1,
						},
					},
				},
			},
			origPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
				},
			},
			expectedResult: &v1beta1.IAMPolicy{
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
					Bindings: []v1beta1.IAMPolicyBinding{
						{
							Members: []v1beta1.Member{
								"foo@bar.iam.gserviceaccount.com",
							},
							Role: testRole1,
						},
					},
					Etag: testEtag,
				},
			},
		},
		{
			name: "iam policy with binding with condition",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "iam",
					Type:    "Policy",
					Version: "ga",
				},
				Object: map[string]interface{}{
					"etag": testEtag,
					"bindings": []interface{}{
						map[string]interface{}{
							"members": []interface{}{
								"foo@bar.iam.gserviceaccount.com",
							},
							"role": testRole1,
							"condition": map[string]interface{}{
								"title":       testTitle1,
								"description": testDescription1,
								"expression":  testExpression1,
							},
						},
					},
				},
			},
			origPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
				},
			},
			expectedResult: &v1beta1.IAMPolicy{
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
					Bindings: []v1beta1.IAMPolicyBinding{
						{
							Members: []v1beta1.Member{
								"foo@bar.iam.gserviceaccount.com",
							},
							Role: testRole1,
							Condition: &v1beta1.IAMCondition{
								Title:       testTitle1,
								Description: testDescription1,
								Expression:  testExpression1,
							},
						},
					},
					Etag: testEtag,
				},
			},
		},
		{
			name: "iam policy with multiple bindings/members/conditions",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "iam",
					Type:    "Policy",
					Version: "ga",
				},
				Object: map[string]interface{}{
					"etag": testEtag,
					"bindings": []interface{}{
						map[string]interface{}{
							"members": []interface{}{
								"foo@bar.iam.gserviceaccount.com",
								"user@test.iam.gserviceaccount.com",
							},
							"role": testRole1,
							"condition": map[string]interface{}{
								"title":       testTitle1,
								"description": testDescription1,
								"expression":  testExpression1,
							},
						},
						map[string]interface{}{
							"role": testRole2,
							"condition": map[string]interface{}{
								"title":       testTitle2,
								"description": testDescription2,
								"expression":  testExpression2,
							},
						},
					},
				},
			},
			origPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
				},
			},
			expectedResult: &v1beta1.IAMPolicy{
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
					Bindings: []v1beta1.IAMPolicyBinding{
						{
							Members: []v1beta1.Member{
								"foo@bar.iam.gserviceaccount.com",
								"user@test.iam.gserviceaccount.com",
							},
							Role: testRole1,
							Condition: &v1beta1.IAMCondition{
								Title:       testTitle1,
								Description: testDescription1,
								Expression:  testExpression1,
							},
						},
						{
							Role: testRole2,
							Condition: &v1beta1.IAMCondition{
								Title:       testTitle2,
								Description: testDescription2,
								Expression:  testExpression2,
							},
						},
					},
					Etag: testEtag,
				},
			},
		},
		{
			name: "error from wrong type",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "binaryauthorization",
					Type:    "Policy",
					Version: "ga",
				},
				Object: map[string]interface{}{
					"etag": testEtag,
				},
			},
			origPolicy:     &v1beta1.IAMPolicy{},
			expectedResult: &v1beta1.IAMPolicy{},
			hasError:       true,
		},
		{
			name: "error from empty DCLObject etag",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "iam",
					Type:    "Policy",
					Version: "ga",
				},
				Object: map[string]interface{}{
					"etag": emptyEtag,
				},
			},
			origPolicy: &v1beta1.IAMPolicy{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicy",
					APIVersion: "v1beta1",
				},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
				},
			},
			expectedResult: &v1beta1.IAMPolicy{
				Spec: v1beta1.IAMPolicySpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "KindBar",
					},
					Bindings:     nil,
					AuditConfigs: nil,
					Etag:         testEtag,
				},
			},
			hasError: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := newIAMPolicyFromDCLResource(tc.dclResource, tc.origPolicy)
			if tc.hasError {
				if err == nil {
					t.Fatalf("expected to get error when converting IAMPolicy from DCLPolicyResource")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tc.expectedResult) {
				t.Fatalf("got %v, want %v", actual, tc.expectedResult)
			}
		})
	}
}

func TestDCLMemberFromKRMMember(t *testing.T) {
	tests := []struct {
		name           string
		dclResource    *dclunstruct.Resource
		krmMember      *v1beta1.IAMPolicyMember
		expectedResult *dcliam.Member
		hasError       bool
	}{
		{
			name: "simple iam policy member",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			krmMember: &v1beta1.IAMPolicyMember{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicyMember",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicyMemberSpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "DataprocCluster",
					},
					Member: v1beta1.Member(testMember1),
					Role:   testRole1,
				},
			},
			expectedResult: &dcliam.Member{
				Role:   &testRole1,
				Member: &testMember1,
				Resource: &dclunstruct.Resource{
					STV: dclunstruct.ServiceTypeVersion{
						Service: "dataproc",
						Type:    "Cluster",
						Version: "ga",
					},
				},
			},
		},
		{
			name: "error from both member and memberFrom used",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			krmMember: &v1beta1.IAMPolicyMember{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicyMember",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicyMemberSpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "DataprocCluster",
					},
					Member: v1beta1.Member(testMember1),
					MemberFrom: &v1beta1.MemberSource{
						ServiceAccountRef: &v1beta1.MemberReference{
							Namespace: "ns-a",
							Name:      "foobar",
						},
					},
					Role: testRole1,
				},
			},
			hasError: true,
		},
		{
			name: "error from neither member and memberFrom used",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			krmMember: &v1beta1.IAMPolicyMember{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicyMember",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicyMemberSpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "DataprocCluster",
					},
					Role: testRole1,
				},
			},
			hasError: true,
		},
		{
			name: "error from member having conditions",
			dclResource: &dclunstruct.Resource{
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Cluster",
					Version: "ga",
				},
			},
			krmMember: &v1beta1.IAMPolicyMember{
				TypeMeta: metav1.TypeMeta{
					Kind:       "IAMPolicyMember",
					APIVersion: "v1beta1",
				},
				Spec: v1beta1.IAMPolicyMemberSpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind: "DataprocCluster",
					},
					Role:   testRole1,
					Member: v1beta1.Member(testMember1),
					Condition: &v1beta1.IAMCondition{
						Title:       testTitle1,
						Description: testDescription1,
						Expression:  testExpression1,
					},
				},
			},
			hasError: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Unit tests will not include testing the `memberFrom` field. This will
			// instead be tested in the integration tests because `memberFrom` will require
			// calling on an actual TF-based resource to resolve (which should not occur
			// in unit tests).
			actual, err := newDCLPolicyMemberFromKRMPolicyMember(context.TODO(), tc.krmMember, tc.dclResource, "", nil)
			if tc.hasError {
				if err == nil {
					t.Fatalf("expected to get error when converting IAMPolicy from DCLPolicyResource")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tc.expectedResult) {
				t.Fatalf("got %v, want %v", actual, tc.expectedResult)
			}
		})
	}
}

func TestKRMMemberFromDCLMember(t *testing.T) {
	tests := []struct {
		name           string
		dclResource    *dclunstruct.Resource
		krmMember      *v1beta1.IAMPolicyMember
		expectedResult *v1beta1.IAMPolicyMember
		hasError       bool
	}{
		{
			name: "simple dcl policy member",
			dclResource: &dclunstruct.Resource{
				Object: map[string]interface{}{
					"role":   testRole1,
					"member": testMember1,
				},
				STV: dclunstruct.ServiceTypeVersion{
					Service: "iam",
					Type:    "PolicyMember",
					Version: "ga",
				},
			},
			krmMember: &v1beta1.IAMPolicyMember{
				ObjectMeta: metav1.ObjectMeta{
					Name: "FooBar",
				},
				Spec: v1beta1.IAMPolicyMemberSpec{
					ResourceReference: v1beta1.ResourceReference{
						Kind:       "DataprocCluster",
						Name:       "Dep",
						APIVersion: "v1beta1",
					},
					Member: v1beta1.Member(testMember1),
				},
			},
			expectedResult: &v1beta1.IAMPolicyMember{
				ObjectMeta: metav1.ObjectMeta{
					Name: "FooBar",
				},
				Spec: v1beta1.IAMPolicyMemberSpec{
					Role:   testRole1,
					Member: v1beta1.Member(testMember1),
					ResourceReference: v1beta1.ResourceReference{
						Kind:       "DataprocCluster",
						Name:       "Dep",
						APIVersion: "v1beta1",
					},
				},
			},
		},
		{
			name: "error from wrong type",
			dclResource: &dclunstruct.Resource{
				Object: map[string]interface{}{
					"role":   testRole1,
					"member": testMember1,
				},
				STV: dclunstruct.ServiceTypeVersion{
					Service: "dataproc",
					Type:    "Member",
					Version: "ga",
				},
			},
			hasError: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := newIAMPolicyMemberFromDCLResource(tc.dclResource, tc.krmMember)
			if tc.hasError {
				if err == nil {
					t.Fatalf("expected to get error when converting IAMPolicy from DCLPolicyResource")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tc.expectedResult) {
				t.Fatalf("got %v, want %v", actual, tc.expectedResult)
			}
		})
	}
}
