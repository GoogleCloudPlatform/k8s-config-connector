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

package v1alpha1

type RBACRoleBindingRole struct {
	/* Optional. custom_role is the name of a custom KubernetesClusterRole to use. */
	// +optional
	CustomRole *string `json:"customRole,omitempty"`

	/* predefined_role is the Kubernetes default role to use. Possible values: UNKNOWN, ADMIN, EDIT, VIEW, ANTHOS_SUPPORT */
	// +optional
	// +kubebuilder:validation:Enum=UNKNOWN;ADMIN;EDIT;VIEW;ANTHOS_SUPPORT
	PredefinedRole *string `json:"predefinedRole,omitempty"`
}

type RBACRoleBindingStateStatus struct {
	/* Output only. Code describes the state of a RBACRoleBinding resource. Possible values: CODE_UNSPECIFIED, CREATING, READY, DELETING, UPDATING */
	// +kubebuilder:validation:Enum=CODE_UNSPECIFIED;CREATING;READY;DELETING;UPDATING
	Code *string `json:"code,omitempty"`
}
