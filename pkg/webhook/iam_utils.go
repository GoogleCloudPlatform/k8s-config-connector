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
	"reflect"

	iamapi "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func isIAMResource(obj *unstructured.Unstructured) bool {
	return isIAMPolicy(obj) || isIAMPartialPolicy(obj) || isIAMPolicyMember(obj) || isIAMAuditConfig(obj)
}

func isIAMPolicy(obj *unstructured.Unstructured) bool {
	return obj.GroupVersionKind() == iamapi.IAMPolicyGVK
}

func isIAMPartialPolicy(obj *unstructured.Unstructured) bool {
	return obj.GroupVersionKind() == iamapi.IAMPartialPolicyGVK
}

func isIAMPolicyMember(obj *unstructured.Unstructured) bool {
	return obj.GroupVersionKind() == iamapi.IAMPolicyMemberGVK
}

func isIAMAuditConfig(obj *unstructured.Unstructured) bool {
	return obj.GroupVersionKind() == iamapi.IAMAuditConfigGVK
}

func isIAMSpecModified(oldSpec, newSpec map[string]interface{}) bool {
	return !reflect.DeepEqual(oldSpec, newSpec)
}

func isIAMResourceReferenceModified(oldSpec, newSpec map[string]interface{}) bool {
	return isRequiredFieldModified(oldSpec, newSpec, "resourceRef")
}

func isIAMAuditConfigServiceModified(oldSpec, newSpec map[string]interface{}) bool {
	return isRequiredFieldModified(oldSpec, newSpec, "service")
}

// isRequiredFieldModified returns true if the given field has been modified.
// It is assumed that the field is present in the spec (hence "required"). If
// the field cannot be found, then the function defaults to true.
func isRequiredFieldModified(oldSpec, newSpec map[string]interface{}, field ...string) bool {
	oldVal, ok, err := unstructured.NestedFieldCopy(oldSpec, field...)
	if !ok || err != nil {
		return true
	}
	newVal, ok, err := unstructured.NestedFieldCopy(newSpec, field...)
	if !ok || err != nil {
		return true
	}
	return !reflect.DeepEqual(oldVal, newVal)
}
