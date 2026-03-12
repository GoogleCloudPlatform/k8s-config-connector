// Copyright 2026 Google LLC
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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AccessContextManagerServicePerimeterGVK = GroupVersion.WithKind("AccessContextManagerServicePerimeter")

// AccessContextManagerServicePerimeterSpec defines the desired state of AccessContextManagerServicePerimeter
type AccessContextManagerServicePerimeterSpec struct {
	/* The AccessContextManagerAccessPolicy this
	AccessContextManagerServicePerimeter lives in. */
	// +required
	AccessPolicyRef *AccessPolicyRef `json:"accessPolicyRef"`

	/* Description of the ServicePerimeter and its use. Does not affect
	behavior. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Specifies the type of the Perimeter. There are two types: regular and
	bridge. Regular Service Perimeter contains resources, access levels,
	and restricted services. Every resource can be in at most
	ONE regular Service Perimeter.

	In addition to being in a regular service perimeter, a resource can also
	be in zero or more perimeter bridges. A perimeter bridge only contains
	resources. Cross project operations are permitted if all effected
	resources share some perimeter (whether bridge or regular). Perimeter
	Bridge does not contain access levels or services: those are governed
	entirely by the regular perimeter that resource is in.

	Perimeter Bridges are typically useful when building more complex
	topologies with many independent perimeters that need to share some data
	with a common perimeter, but should not be able to share data among
	themselves. Default value: "PERIMETER_TYPE_REGULAR" Possible values: ["PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE"]. */
	// +optional
	PerimeterType *string `json:"perimeterType,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for
	creation and acquisition. When unset, the value of `metadata.name`
	is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Proposed (or dry run) ServicePerimeter configuration.
	This configuration allows to specify and test ServicePerimeter configuration
	without enforcing actual access restrictions. Only allowed to be set when
	the 'useExplicitDryRunSpec' flag is set. */
	// +optional
	Spec *AccessContextManagerServicePerimeterConfig `json:"spec,omitempty"`

	/* ServicePerimeter configuration. Specifies sets of resources,
	restricted services and access levels that determine
	perimeter content and boundaries. */
	// +optional
	Status *AccessContextManagerServicePerimeterConfig `json:"status,omitempty"`

	/* Human readable title. Must be unique within the Policy. */
	// +required
	Title *string `json:"title"`

	/* Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
	for all Service Perimeters, and that spec is identical to the status for those
	Service Perimeters. When this flag is set, it inhibits the generation of the
	implicit spec, thereby allowing the user to explicitly provide a
	configuration ("spec") to use in a dry-run version of the Service Perimeter.
	This allows the user to test changes to the enforced config ("status") without
	actually enforcing them. This testing is done through analyzing the differences
	between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	bet set to True if any of the fields in the spec are set to non-default values. */
	// +optional
	UseExplicitDryRunSpec *bool `json:"useExplicitDryRunSpec,omitempty"`
}

type AccessContextManagerServicePerimeterConfig struct {
	/* (Optional) A list of AccessLevel resource names that allow resources within
	the ServicePerimeter to be accessed from the internet. AccessLevels listed
	must be in the same policy as this ServicePerimeter.
	Referencing a nonexistent AccessLevel is a syntax error. If no
	AccessLevel names are listed, resources within the perimeter can
	only be accessed via GCP calls with request origins within the
	perimeter. For Service Perimeter Bridge, must be empty. */
	// +optional
	AccessLevels []AccessLevelRef `json:"accessLevels,omitempty"`

	/* List of EgressPolicies to apply to the perimeter. A perimeter may
	have multiple EgressPolicies, each of which is evaluated separately.
	Access is granted if any EgressPolicy grants it. Must be empty for
	a perimeter bridge. */
	// +optional
	EgressPolicies []AccessContextManagerServicePerimeterEgressPolicy `json:"egressPolicies,omitempty"`

	/* List of 'IngressPolicies' to apply to the perimeter. A perimeter may
	have multiple 'IngressPolicies', each of which is evaluated
	separately. Access is granted if any 'Ingress Policy' grants it.
	Must be empty for a perimeter bridge. */
	// +optional
	IngressPolicies []AccessContextManagerServicePerimeterIngressPolicy `json:"ingressPolicies,omitempty"`

	/* (Optional) A list of GCP resources that are inside of the service perimeter.
	Currently only projects are allowed. */
	// +optional
	Resources []AccessContextManagerServicePerimeterResource `json:"resources,omitempty"`

	/* GCP services that are subject to the Service Perimeter
	restrictions. Must contain a list of services. For example, if
	'storage.googleapis.com' is specified, access to the storage
	buckets inside the perimeter must meet the perimeter's access
	restrictions. */
	// +optional
	RestrictedServices []string `json:"restrictedServices,omitempty"`

	/* Specifies how APIs are allowed to communicate within the Service
	Perimeter. */
	// +optional
	VPCAccessibleServices *AccessContextManagerServicePerimeterVPCAccessibleServices `json:"vpcAccessibleServices,omitempty"`
}

type AccessContextManagerServicePerimeterEgressPolicy struct {
	/* Defines conditions on the source of a request causing this 'EgressPolicy' to apply. */
	// +optional
	EgressFrom *AccessContextManagerServicePerimeterEgressFrom `json:"egressFrom,omitempty"`

	/* Defines the conditions on the 'ApiOperation' and destination resources that
	cause this 'EgressPolicy' to apply. */
	// +optional
	EgressTo *AccessContextManagerServicePerimeterEgressTo `json:"egressTo,omitempty"`
}

type AccessContextManagerServicePerimeterEgressFrom struct {
	/* (Optional) A list of identities that are allowed access through this
	EgressPolicy. Should be in the format of email address. The email
	address should represent individual user or service account only. */
	// +optional
	Identities []AccessContextManagerServicePerimeterIdentity `json:"identities,omitempty"`

	/* Specifies the type of identities that are allowed access to outside the
	perimeter. If left unspecified, then members of 'identities' field will
	be allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]. */
	// +optional
	IdentityType *string `json:"identityType,omitempty"`
}

type AccessContextManagerServicePerimeterIdentity struct {
	/* A reference to an IAMServiceAccount resource. */
	// +optional
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	/* A user identity, should represent individual user or service account only. */
	// +optional
	User *string `json:"user,omitempty"`
}

type AccessContextManagerServicePerimeterEgressTo struct {
	/* A list of external resources that are allowed to be accessed. A request
	matches if it contains an external resource in this list (Example:
	s3://bucket/path). Currently '*' is not allowed. */
	// +optional
	ExternalResources []string `json:"externalResources,omitempty"`

	/* A list of 'ApiOperations' that this egress rule applies to. A request matches
	if it contains an operation/service in this list. */
	// +optional
	Operations []AccessContextManagerServicePerimeterApiOperation `json:"operations,omitempty"`

	/* (Optional) A list of resources, currently only projects in the form
	"projects/{project_number}". A request
	matches if it contains a resource in this list. */
	// +optional
	Resources []AccessContextManagerServicePerimeterEgressResource `json:"resources,omitempty"`
}

type AccessContextManagerServicePerimeterEgressResource struct {
	/* (Optional) A list of resources, currently only projects in the form
	"projects/{project_number}". A request
	matches if it contains a resource in this list. */
	// +optional
	ProjectRef *ServicePerimeterProjectRef `json:"projectRef,omitempty"`
}

type AccessContextManagerServicePerimeterApiOperation struct {
	/* API methods or permissions to allow. Method or permission must belong
	to the service specified by 'serviceName' field. A single MethodSelector
	entry with '*' specified for the 'method' field will allow all methods
	AND permissions for the service specified in 'serviceName'. */
	// +optional
	MethodSelectors []AccessContextManagerServicePerimeterMethodSelector `json:"methodSelectors,omitempty"`

	/* The name of the API whose methods or permissions the 'IngressPolicy' or
	'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName
	field set to '*' will allow all methods AND permissions for all services. */
	// +optional
	ServiceName *string `json:"serviceName,omitempty"`
}

type AccessContextManagerServicePerimeterMethodSelector struct {
	/* Value for 'method' should be a valid method name for the corresponding
	'serviceName' in 'ApiOperation'. If '*' used as value for method,
	then ALL methods and permissions are allowed. */
	// +optional
	Method *string `json:"method,omitempty"`

	/* Value for permission should be a valid Cloud IAM permission for the
	corresponding 'serviceName' in 'ApiOperation'. */
	// +optional
	Permission *string `json:"permission,omitempty"`
}

type AccessContextManagerServicePerimeterIngressPolicy struct {
	/* Defines the conditions on the source of a request causing this 'IngressPolicy'
	to apply. */
	// +optional
	IngressFrom *AccessContextManagerServicePerimeterIngressFrom `json:"ingressFrom,omitempty"`

	/* Defines the conditions on the 'ApiOperation' and request destination that cause
	this 'IngressPolicy' to apply. */
	// +optional
	IngressTo *AccessContextManagerServicePerimeterIngressTo `json:"ingressTo,omitempty"`
}

type AccessContextManagerServicePerimeterIngressFrom struct {
	/* (Optional) A list of identities that are allowed access through this
	ingress policy. Should be in the format of email address. The email
	address should represent individual user or service account only. */
	// +optional
	Identities []AccessContextManagerServicePerimeterIdentity `json:"identities,omitempty"`

	/* Specifies the type of identities that are allowed access from outside the
	perimeter. If left unspecified, then members of 'identities' field will be
	allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]. */
	// +optional
	IdentityType *string `json:"identityType,omitempty"`

	/* Sources that this 'IngressPolicy' authorizes access from. */
	// +optional
	Sources []AccessContextManagerServicePerimeterIngressSource `json:"sources,omitempty"`
}

type AccessContextManagerServicePerimeterIngressSource struct {
	/* (Optional) A reference to an AccessLevel resource that is allowed to ingress the perimeter. */
	// +optional
	AccessLevelRef *AccessLevelRef `json:"accessLevelRef,omitempty"`

	/* (Optional) A Google Cloud resource that is allowed to ingress the
	perimeter. Requests from these resources will be allowed to access
	perimeter data. Currently only projects are allowed. Format
	"projects/{project_number}" The project may be in any Google Cloud
	organization, not just the organization that the perimeter is defined in. */
	// +optional
	ProjectRef *ServicePerimeterProjectRef `json:"projectRef,omitempty"`
}

type AccessContextManagerServicePerimeterIngressTo struct {
	/* A list of 'ApiOperations' the sources specified in corresponding 'IngressFrom'
	are allowed to perform in this 'ServicePerimeter'. */
	// +optional
	Operations []AccessContextManagerServicePerimeterApiOperation `json:"operations,omitempty"`

	/* A list of resources, currently only projects in the form
	"projects/{project_number}", protected by this ServicePerimeter
	that are allowed to be accessed by sources defined in the
	corresponding IngressFrom. A request matches if it contains a
	resource in this list. */
	// +optional
	Resources []AccessContextManagerServicePerimeterIngressResource `json:"resources,omitempty"`
}

type AccessContextManagerServicePerimeterIngressResource struct {
	/* A list of resources, currently only projects in the form
	"projects/{project_number}", protected by this ServicePerimeter
	that are allowed to be accessed by sources defined in the
	corresponding IngressFrom. A request matches if it contains a
	resource in this list. */
	// +optional
	ProjectRef *ServicePerimeterProjectRef `json:"projectRef,omitempty"`
}

type AccessContextManagerServicePerimeterResource struct {
	/* (Optional) A list of GCP resources that are inside of the service perimeter.
	Currently only projects are allowed. */
	// +optional
	ProjectRef *ServicePerimeterProjectRef `json:"projectRef,omitempty"`
}

type AccessContextManagerServicePerimeterVPCAccessibleServices struct {
	/* The list of APIs usable within the Service Perimeter.
	Must be empty unless 'enableRestriction' is True. */
	// +optional
	AllowedServices []string `json:"allowedServices,omitempty"`

	/* Whether to restrict API calls within the Service Perimeter to the
	list of APIs specified in 'allowedServices'. */
	// +optional
	EnableRestriction *bool `json:"enableRestriction,omitempty"`
}

type ServicePerimeterProjectRef struct {
	/* Allowed value: string of the format `projects/{{value}}`, where {{value}} is the `number` field of a `Project` resource. */
	// +optional
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	// +optional
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// AccessContextManagerServicePerimeterStatus defines the config connector machine state of AccessContextManagerServicePerimeter
type AccessContextManagerServicePerimeterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	// +kubebuilder:validation:Format=""
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AccessContextManagerServicePerimeter resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	/* Time the AccessPolicy was created in UTC. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Time the AccessPolicy was updated in UTC. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaccesscontextmanagerserviceperimeter;gcpaccesscontextmanagerserviceperimeters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AccessContextManagerServicePerimeter is the Schema for the AccessContextManagerServicePerimeter API
// +k8s:openapi-gen=true
type AccessContextManagerServicePerimeter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AccessContextManagerServicePerimeterSpec   `json:"spec,omitempty"`
	Status AccessContextManagerServicePerimeterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AccessContextManagerServicePerimeterList contains a list of AccessContextManagerServicePerimeter
type AccessContextManagerServicePerimeterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerServicePerimeter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerServicePerimeter{}, &AccessContextManagerServicePerimeterList{})
}
