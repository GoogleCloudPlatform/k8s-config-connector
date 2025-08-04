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
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	testutil "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	// Ensure built-in types are registered.
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

func TestDefaultAPIVersionForIAMResourceRef(t *testing.T) {
	tests := []struct {
		name           string
		resourceRef    map[string]interface{}
		resourceRefNew map[string]interface{}
		denied         bool
	}{
		{
			name:   "spec.resourceRef is not set",
			denied: true,
		},
		{
			name: "spec.resourceRef.kind is not set",
			resourceRef: map[string]interface{}{
				"name": "resource-name",
			},
			denied: true,
		},
		{
			name: "spec.resourceRef.kind is set to unsupported kind",
			resourceRef: map[string]interface{}{
				"kind": "UnsupportedKind",
				"name": "resource-name",
			},
			denied: true,
		},
		{
			name: "spec.resourceRef.apiVersion is already set",
			resourceRef: map[string]interface{}{
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"kind":       "PubSubTopic",
				"name":       "resource-name",
			},
			resourceRefNew: map[string]interface{}{
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"kind":       "PubSubTopic",
				"name":       "resource-name",
			},
		},
		{
			name: "spec.resourceRef.apiVersion is not set",
			resourceRef: map[string]interface{}{
				"kind": "PubSubTopic",
				"name": "resource-name",
			},
			resourceRefNew: map[string]interface{}{
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"kind":       "PubSubTopic",
				"name":       "resource-name",
			},
		},
		{
			name: "spec.resourceRef.apiVersion is not set, and spec.resourceRef.kind is set to an external-only kind Organization",
			resourceRef: map[string]interface{}{
				"kind": "Organization",
				"name": "resource-name",
			},
			resourceRefNew: map[string]interface{}{
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"kind":       "Organization",
				"name":       "resource-name",
			},
		},
		{
			name: "spec.resourceRef.apiVersion is not set, and spec.resourceRef.kind is set to an external-only kind BillingAccount",
			resourceRef: map[string]interface{}{
				"kind": "BillingAccount",
				"name": "resource-name",
			},
			resourceRefNew: map[string]interface{}{
				"apiVersion": "billing.cnrm.cloud.google.com/v1beta1",
				"kind":       "BillingAccount",
				"name":       "resource-name",
			},
		},
	}

	iamTemplates := map[string]map[string]interface{}{
		"IAMPolicy": {
			"apiVersion": v1beta1.IAMPolicyGVK.GroupVersion().String(),
			"kind":       v1beta1.IAMPolicyGVK.Kind,
			"metadata": map[string]interface{}{
				"name": "policy",
			},
			"spec": map[string]interface{}{
				"bindings": []interface{}{
					map[string]interface{}{
						"role": "role",
						"members": []interface{}{
							"member",
						},
					},
				},
			},
		},
		"IAMPartialPolicy": {
			"apiVersion": v1beta1.IAMPartialPolicyGVK.GroupVersion().String(),
			"kind":       v1beta1.IAMPartialPolicyGVK.Kind,
			"metadata": map[string]interface{}{
				"name": "policy",
			},
			"spec": map[string]interface{}{
				"bindings": []interface{}{
					map[string]interface{}{
						"role": "role",
						"members": []interface{}{
							map[string]interface{}{
								"member": "member",
							},
						},
					},
				},
			},
		},
		"IAMPolicyMember": {
			"apiVersion": v1beta1.IAMPolicyMemberGVK.GroupVersion().String(),
			"kind":       v1beta1.IAMPolicyMemberGVK.Kind,
			"metadata": map[string]interface{}{
				"name": "policy",
			},
			"spec": map[string]interface{}{
				"member": "member",
				"role":   "role",
			},
		},
		"IAMAuditConfig": {
			"apiVersion": v1beta1.IAMAuditConfigGVK.GroupVersion().String(),
			"kind":       v1beta1.IAMAuditConfigGVK.Kind,
			"metadata": map[string]interface{}{
				"name": "policy",
			},
			"spec": map[string]interface{}{
				"service": "service",
				"auditLogConfigs": []interface{}{
					map[string]interface{}{
						"logType": "log-type",
						"exemptedMembers": []interface{}{
							"member",
						},
					},
				},
			},
		},
	}

	for iamKind, iamTemplate := range iamTemplates {
		for _, tc := range tests {
			iamKind := iamKind
			iamTemplate := iamTemplate
			tc := tc
			t.Run(fmt.Sprintf("%v (%v)", tc.name, iamKind), func(t *testing.T) {
				t.Parallel()

				// Create IAM resource to pass into webhook
				iamObj := &unstructured.Unstructured{
					Object: deepcopy.MapStringInterface(iamTemplate),
				}
				if tc.resourceRef != nil {
					if err := unstructured.SetNestedField(iamObj.Object, tc.resourceRef, iamResourceRefPath...); err != nil {
						t.Fatalf("error setting '%v' on IAM resource meant to be passed into webhook: %v", iamResourceRefField, err)
					}
				}

				// Create IAM resource that we expect the webhook to return
				iamObjExpected := &unstructured.Unstructured{
					Object: deepcopy.MapStringInterface(iamTemplate),
				}
				if tc.resourceRefNew != nil {
					if err := unstructured.SetNestedField(iamObjExpected.Object, tc.resourceRefNew, iamResourceRefPath...); err != nil {
						t.Fatalf("error setting '%v' on IAM resource meant for testing webhook response: %v", iamResourceRefField, err)
					}
				}

				response := defaultAPIVersionForIAMResourceRef(iamObj, testservicemappingloader.New(t), dclmetadata.New())
				if tc.denied {
					if response.Allowed {
						t.Fatalf("expected request to be denied, but was allowed. Response:\n%v", response)
					}
					return
				}
				if !response.Allowed {
					t.Fatalf("request was unexpectedly denied. Response:\n%v", response)
				}
				expectedResponse := constructPatchResponse(iamObj, iamObjExpected)
				if !testutil.Equals(t, response, expectedResponse) {
					diff := cmp.Diff(expectedResponse, response)
					t.Fatalf("unexpected diff in the response (-want +got):\n%v", diff)
				}
			})
		}
	}
}
