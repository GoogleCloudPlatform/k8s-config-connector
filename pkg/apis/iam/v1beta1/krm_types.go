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

package v1beta1

import "k8s.io/apimachinery/pkg/runtime/schema"

type ResourceReference struct {
	Kind       string `json:"kind"`
	Namespace  string `json:"namespace,omitempty"`
	Name       string `json:"name,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
	External   string `json:"external,omitempty"`
}

func (ref *ResourceReference) GroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(ref.APIVersion, ref.Kind)
}

func (ref *ResourceReference) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	ref.APIVersion, ref.Kind = gvk.ToAPIVersionAndKind()
}

// MemberSource represents a source for an IAM identity
type MemberSource struct {
	// The IAMServiceAccount to be bound to the role.
	ServiceAccountRef *MemberReference `json:"serviceAccountRef,omitempty"`

	// The LoggingLogSink whose writer identity (i.e. its
	// 'status.writerIdentity') is to be bound to the role.
	LogSinkRef *MemberReference `json:"logSinkRef,omitempty"`

	// The SQLInstance whose service account (i.e. its
	// 'status.serviceAccountEmailAddress') is to be bound to the role.
	SQLInstanceRef *MemberReference `json:"sqlInstanceRef,omitempty"`

	// The ServiceIdentity whose service account (i.e., its
	// 'status.email') is to be bound to the role.
	ServiceIdentityRef *MemberReference `json:"serviceIdentityRef,omitempty"`
}

// MemberReference represents a resource with an IAM identity
type MemberReference struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name"`
}

// IAMCondition defines the IAM condition under which an IAM binding applies
type IAMCondition struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Expression  string `json:"expression"`
}

type AuditLogConfig struct {
	// Permission type for which logging is to be configured. Must be one of
	// 'DATA_READ', 'DATA_WRITE', or 'ADMIN_READ'.
	// +kubebuilder:validation:Pattern=^(DATA_READ|DATA_WRITE|ADMIN_READ)$
	LogType string `json:"logType"`

	// Identities that do not cause logging for this type of permission. The
	// format is the same as that for 'members' in IAMPolicy/IAMPolicyMember.
	ExemptedMembers []Member `json:"exemptedMembers,omitempty"`
}
