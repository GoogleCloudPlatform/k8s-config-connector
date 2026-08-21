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

// NOTE: Boilerplate only.  Ignore this file.

// Package v1beta1 contains API Schema definitions for the iam v1beta1 API group
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam
// +k8s:defaulter-gen=TypeMeta
// +groupName=iam.cnrm.cloud.google.com
package v1beta1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	IAMPolicyGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    reflect.TypeOf(IAMPolicy{}).Name(),
	}

	IAMPartialPolicyGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    reflect.TypeOf(IAMPartialPolicy{}).Name(),
	}
	IAMPolicyMemberGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    reflect.TypeOf(IAMPolicyMember{}).Name(),
	}
	IAMAuditConfigGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    reflect.TypeOf(IAMAuditConfig{}).Name(),
	}
	IAMAPIVersion = GroupVersion.String()
)

// IsHandwrittenIAM returns true if the given GVK corresponds to that of a
// handwritten IAM resource.
func IsHandwrittenIAM(gvk schema.GroupVersionKind) bool {
	switch gvk {
	case IAMPolicyGVK, IAMPolicyMemberGVK, IAMAuditConfigGVK, IAMPartialPolicyGVK:
		return true
	default:
		return false
	}
}
