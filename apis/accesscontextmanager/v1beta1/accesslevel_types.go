// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AccessContextManagerAccessLevelGVK = GroupVersion.WithKind("AccessContextManagerAccessLevel")

type AccessLevelBasic struct {
	// How the conditions list should be combined to determine if a request
	// is granted this AccessLevel. If AND is used, each Condition in
	// conditions must be satisfied for the AccessLevel to be applied. If
	// OR is used, at least one Condition in conditions must be satisfied
	// for the AccessLevel to be applied. Default value: "AND" Possible values: ["AND", "OR"].
	CombiningFunction *string `json:"combiningFunction,omitempty"`

	// A set of requirements for the AccessLevel to be granted.
	// +required
	Conditions []BasicConditions `json:"conditions"`
}

type BasicConditions struct {
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	DevicePolicy *ConditionsDevicePolicy `json:"devicePolicy,omitempty"`

	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	IPSubnetworks []string `json:"ipSubnetworks,omitempty"`

	// An allowed list of members (users, service accounts).
	// Using groups is not supported.
	//
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	Members []ConditionsMembers `json:"members,omitempty"`

	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	Negate *bool `json:"negate,omitempty"`

	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	Regions []string `json:"regions,omitempty"`

	// A list of other access levels defined in the same policy.
	// Referencing an AccessContextManagerAccessLevel which does not exist
	// is an error. All access levels listed must be granted for the
	// condition to be true.
	RequiredAccessLevels []ConditionsRequiredAccessLevels `json:"requiredAccessLevels,omitempty"`
}

type ConditionsDevicePolicy struct {
	// A list of allowed device management levels.
	// An empty list allows all management levels. Possible values: ["MANAGEMENT_UNSPECIFIED", "NONE", "BASIC", "COMPLETE"].
	AllowedDeviceManagementLevels []string `json:"allowedDeviceManagementLevels,omitempty"`

	// A list of allowed encryptions statuses.
	// An empty list allows all statuses. Possible values: ["ENCRYPTION_UNSPECIFIED", "ENCRYPTION_UNSUPPORTED", "UNENCRYPTED", "ENCRYPTED"].
	AllowedEncryptionStatuses []string `json:"allowedEncryptionStatuses,omitempty"`

	// A list of allowed OS versions.
	// An empty list allows all types and all versions.
	OsConstraints []DevicePolicyOsConstraints `json:"osConstraints,omitempty"`

	// Whether the device needs to be approved by the customer admin.
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty"`

	// Whether the device needs to be corp owned.
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty"`

	// Whether or not screenlock is required for the DevicePolicy
	// to be true. Defaults to false.
	RequireScreenLock *bool `json:"requireScreenLock,omitempty"`
}

type DevicePolicyOsConstraints struct {
	// The minimum allowed OS version. If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	MinimumVersion *string `json:"minimumVersion,omitempty"`

	// The operating system type of the device. Possible values: ["OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS", "ANDROID", "IOS"].
	// +required
	OsType *string `json:"osType"`

	// If you specify DESKTOP_CHROME_OS for osType, you can optionally include requireVerifiedChromeOs
	// to require Chrome Verified Access.
	RequireVerifiedChromeOs *bool `json:"requireVerifiedChromeOs,omitempty"`
}

type ConditionsMembers struct {
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
	User              *string               `json:"user,omitempty"`
}

type ConditionsRequiredAccessLevels struct {
	// The 'name' field of an 'AccessContextManagerAccessLevel' resource.
	External *string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace *string `json:"namespace,omitempty"`
}

type AccessLevelCustom struct {
	// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
	// This page details the objects and attributes that are used to the build the CEL expressions for
	// custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.
	// +required
	Expr *CustomExpr `json:"expr"`
}

type CustomExpr struct {
	// Description of the expression.
	Description *string `json:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	// +required
	Expression *string `json:"expression"`

	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `json:"location,omitempty"`

	// Title for the expression, i.e. a short string describing its purpose.
	Title *string `json:"title,omitempty"`
}

// AccessContextManagerAccessLevelSpec defines the desired state of AccessContextManagerAccessLevel
type AccessContextManagerAccessLevelSpec struct {
	// The AccessContextManagerAccessPolicy this
	// AccessContextManagerAccessLevel lives in.
	// +required
	AccessPolicyRef v1alpha1.ResourceRef `json:"accessPolicyRef"`

	// A set of predefined conditions for the access level and a combining function.
	Basic *AccessLevelBasic `json:"basic,omitempty"`

	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	Custom *AccessLevelCustom `json:"custom,omitempty"`

	// Description of the AccessLevel and its use. Does not affect behavior.
	Description *string `json:"description,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Human readable title. Must be unique within the Policy.
	// +required
	Title *string `json:"title"`
}

// AccessContextManagerAccessLevelStatus defines the config connector machine state of AccessContextManagerAccessLevel
type AccessContextManagerAccessLevelStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaccesscontextmanageraccesslevel;gcpaccesscontextmanageraccesslevels
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AccessContextManagerAccessLevel is the Schema for the AccessContextManagerAccessLevel API
// +k8s:openapi-gen=true
type AccessContextManagerAccessLevel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AccessContextManagerAccessLevelSpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessLevelStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AccessContextManagerAccessLevelList contains a list of AccessContextManagerAccessLevel
type AccessContextManagerAccessLevelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerAccessLevel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerAccessLevel{}, &AccessContextManagerAccessLevelList{})
}
