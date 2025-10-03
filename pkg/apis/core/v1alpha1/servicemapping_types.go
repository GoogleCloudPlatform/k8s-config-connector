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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type IAMReferenceType string

const (
	IAMReferenceTypeName IAMReferenceType = "name"
	IAMReferenceTypeId   IAMReferenceType = "id" //nolint:revive
)

type ServiceMappingSpec struct {
	// Name is the name of the service being mapped (e.g. Spanner, PubSub). This is
	// used for the construction of the generated CRDs' API group and kind.
	Name string `json:"name"`

	// Version is the default API version for all the resource CRDs being
	// generated.
	Version string `json:"version"`

	// ServiceHostName is the host portion of the URL for the associated service. IE, for Spanner, it is 'spanner.googleapis.com'
	ServiceHostName string `json:"serviceHostName"`

	// Resources is a list of configurations specifying how to map a specific resource
	// from the Terraform provider to KRM.
	Resources []ResourceConfig `json:"resources"`
}

type ResourceConfig struct {
	// Name is the Terraform name of the resource (e.g. google_spanner_instance)
	Name string `json:"name"`

	// Kind is the Kubernetes kind you wish the resource to have.
	Kind string `json:"kind"`

	// Version is the API version of the resource CRD.
	// If unset, the default API version of the service mapping will be used.
	Version *string `json:"version"`

	// Direct tells if the ResourceConfig is for ConfigConnector directly managed resources.
	// Directly managed resource does not use Terraform or DCL controller, and do not rely on any TF specified fields like `SkipImport`
	// A direct ResourceConfig is used to generate the reference doc.
	Direct bool `json:"direct"`

	// SkipImport skips the import step when fetching the live state of the underlying
	// resource. If specified, IDTemplate must also be specified, and its expanded
	// form will be used as the TF resource's `id` field.
	SkipImport bool `json:"skipImport,omitempty"`

	// IAMConfig contains the mappings from a given resource onto its associated terraform IAM resources (policies, bindings, and members)
	IAMConfig IAMConfig `json:"iamConfig,omitempty"`

	// IAMMemberReferenceConfig configures the resource as a resource that can
	// be referenced as an IAM member.
	IAMMemberReferenceConfig IAMMemberReferenceConfig `json:"iamMemberReferenceConfig,omitempty"`

	// IDTemplate defines the format in which the ID fed into the TF resource's importer
	// should look. Fields may be sourced from the TF resource by using the `{{foo}}`
	// syntax. (e.g. {{project}}/{{location}}/{{name}}.
	//
	// All fields are required. A field can be marked as optional with the ? suffix, e.g. with {{project}}/{{host?}},
	// the host field is optional
	//
	// An OR condition can be defined on a portion of the template by enclosing the portion with brackets `[...]` and using
	// a bar character, `|`, to deliminate the OR. Example, `my-template/[{{field1}}|text_{{field2}]`.
	//
	// If SkipImport is true, this must be specified, and its expanded form will be directly
	// used as the TF resource's `id` field.
	IDTemplate string `json:"idTemplate,omitempty"`

	// The resource name is the One Platform / GCP resource name, when this value is true, it means the IDTemplate
	// can be converted to a regex and used to match against a given URL to determine if it is a name for the given
	// ResourceConfig. If this flag is true then the ID Template is used by "config-connector export" to match against
	// URLs.
	//
	// see: https://cloud.google.com/apis/design/resource_names
	IDTemplateCanBeUsedToMatchResourceName *bool `json:"idTemplateCanBeUsedToMatchResourceName,omitempty"`

	// ServerGeneratedIDField is the field in the resource's status that corresponds to
	// the server-generated resource ID. If unset, it's assumed the resource ID is specified
	// by the user. Resources with this set do not support acquisition.
	ServerGeneratedIDField string `json:"serverGeneratedIDField,omitempty"`

	// Locationality categorizes the GCP resources as global, regional, or zonal. It's only applicable to the effort of
	// unifying multiple locational TF resources into one, e.g. KCC could have a single ComputeAddress CRD to represent
	// two TF/GCE resources - compute address and global compute address. The location field in ComputeAddress CRD is used to specify
	// whether it is a global address or regional address. If unset, it's assumed that there is no multiple TF locational resources
	// mapping to the same compute resource schema. Currently, this supports the following values: global, regional, zonal.
	Locationality string `json:"locationality,omitempty"`

	// MetadataMapping determines how to map Kubernetes metadata fields to the Terraform
	// resource's configuration.
	MetadataMapping MetadataMapping `json:"metadataMapping,omitempty"`

	// ResourceID determines how to map the `spec.resourceID` field to the
	// Terraform resource's configuration.
	// For multiple ResourceConfigs that map to the same Kind, their ResourceID
	// definition must be the same.
	ResourceID ResourceID `json:"resourceID,omitempty"`

	// ResourceReferences configures the mapping of fields in the Terraform resource that
	// implicitly define references to other GCP resources into explicit Kubernetes-style
	// references.
	ResourceReferences []ReferenceConfig `json:"resourceReferences,omitempty"`

	// Directives is a list of Terraform fields that perform unique behaviors on
	// top of the resource which are not part of a GET response. If the KCC annotation's
	// key contains a directive from this list (e.g. `cnrm.cloud.google.com/force-destroy`),
	// the value from the annotation is stored/overwritten in the TF config (e.g. force_destroy -> true)
	Directives []string `json:"directives,omitempty"`

	// MutableButUnreadableFields is a list of Terraform fields that are
	// mutable but not returned by the Terraform read. KCC tracks the values of
	// such fields to be able to determine if the user changed their values on
	// the spec.
	MutableButUnreadableFields []string `json:"mutableButUnreadableFields,omitempty"`

	// IgnoredFields is a list of fields that should be dropped from the underlying
	// Terraform resource.
	IgnoredFields []string `json:"ignoredFields,omitempty"`

	// IgnoredOutputOnlySpecFields is a list of fields that should not be added
	// to spec because they are output-only.
	// We have a legacy bug that adds all the fields under a writable top-level
	// field into spec during CRD generation even if the subfield itself is
	// output-only. We should stop the bleeding by manually adding any new
	// output-only subfields under a writable top-level field into this list.
	IgnoredOutputOnlySpecFields *[]string `json:"ignoredOutputOnlySpecFields,omitempty"`

	// Deprecated: use HierarchicalReferences instead. Only resources that
	// already specify Containers should continue to specify Containers so that
	// these resources can continue to support resource-level container
	// annotations. Since new resources should not support resource-level
	// container annotations, they should not specify Containers.
	//
	// Containers describes all the container mappings this resource understands. Config Connector maps Kubernetes
	// namespaces to the abstract GCP container objects they are scoped by via namespaces. For most resource types,
	// this is a project, but certain resources live outside the scope of a project, like folders or projects
	// themselves. Containers are expressed as annotations on a given Namespace, though users may provide
	// resource-level overrides.
	Containers []Container `json:"containers,omitempty"`

	// HierarchicalReferences lists the resource references that represent
	// references to hierarchical resources (e.g. project, folder,
	// organization).
	HierarchicalReferences []HierarchicalReference `json:"hierarchicalReferences,omitempty"`

	// ResourceAvailableInAssetInventory specifies if the resource exists in asset inventory
	// visible here: https://cloud.google.com/asset-inventory/docs/supported-asset-types#supported_resource_types
	ResourceAvailableInAssetInventory bool `json:"resourceAvailableInAssetInventory,omitempty"`

	// AutoGenerated indicates whether the resource is generated automatically
	// via the KCC resource auto-generation process.
	// A ResourceConfig should not set this field to 'true' unless it is
	// generated via auto generation.
	AutoGenerated bool `json:"autoGenerated,omitempty"`

	// ReconciliationIntervalInSeconds specifies the default mean reconciliation interval for this resource.
	// Providing the value in servicemapping config is optional. If not explicitly configured a global
	// default value of 600 will be used
	ReconciliationIntervalInSeconds *uint32 `json:"reconciliationIntervalInSeconds,omitempty"`

	// Unreadable indicates whether the resource can be read from the underlying API.
	// This field should be set to true if the underlying API does not support read,
	// or if the read is no-op.
	// This field is optional. If unset, the default value of nil indicates the
	// underlying API does support read.
	// If set to true, ReconciliationIntervalInSeconds should also be set to 0 to avoid
	// repeated creation of the resource.
	Unreadable *bool `json:"unreadable,omitempty"`

	// V1alpha1ToV1beta1 indicates whether the resource is during v1alpha1 to
	// v1beta1 conversion.
	V1alpha1ToV1beta1 *bool `json:"v1alpha1ToV1beta1,omitempty"`

	// StorageVersion indicates which storage version the resource uses.
	// It must be set when V1alpha1ToV1beta1 is set to `true`. It must be
	// `v1alpha1` when first set.
	StorageVersion *string `json:"storageVersion,omitempty"`

	// ObservedFields specifies which `spec` fields should be exposed under
	// `status.observedState` in the CRD, and in the CR after a successful
	// reconciliation.
	// The fields should be snake case paths in TF. For example,
	// `master_auth.client_certificate`.
	ObservedFields *[]string `json:"observedFields,omitempty"`
}

type IAMConfig struct {
	// PolicyName is the terraform name of the associated IAM Policy resource (e.g. google_spanner_instance_iam_policy)
	PolicyName string `json:"policyName"`

	// PolicyMemberName is the terraform name of the associated IAM Policy Member resource (e.g. google_spanner_instance_iam_member)
	PolicyMemberName string `json:"policyMemberName"`

	// AuditConfigName is the terraform name of the associated IAM Audit Config resource, if there is any (e.g. google_project_iam_audit_config for the Project resource)
	AuditConfigName string `json:"auditConfigName,omitempty"`

	// A description of the manner in which the IAM Policy references its resource.
	ReferenceField IAMReferenceField `json:"referenceField,omitempty"`

	// SupportsConditions indicates whether or not the resource supports IAM Conditions.
	SupportsConditions bool `json:"supportsConditions"`
}

// A reference from an IAM policy or binding to a resource.
type IAMReferenceField struct {
	// The name of the field in the policy or binding which references the resource. For
	// 'google_spanner_instance_iam_policy' this value is 'instance'.
	Name string `json:"name"`

	// The type of value that should be used in this field. It can be one of { name, id }. For
	// 'google_spanner_instance_iam_policy' it would be 'name' for 'google_kms_key_ring_iam_policy'
	// it would be 'id'.
	Type IAMReferenceType `json:"type"`
}

type MetadataMapping struct {
	// Name is a JSONPath to the field in the TF resource where the KRM "metadata.name" field will be mapped to. By
	// default, this is mapped to the "name" field, if that field is found in the TF resource schema.
	Name string `json:"name,omitempty"`

	// NameValueTemplate is a template by which the value of the metadata.name field
	// should be interpreted before being passed to the Terraform provider. {{value}}
	// is used in place of this sourced value.
	//
	// e.g. If the value sourced from metadata.name is "foo_bar", a nameValueTemplate of
	// "resource/{{value}}" would mean the final value passed to the provider is
	// "resource/foo_bar"
	NameValueTemplate string `json:"nameValueTemplate,omitempty"`

	// Labels is a JSONPath to the field in the TF resource where the KRM "metadata.labels" field will be mapped to. By
	// default, this is mapped to the "labels" field, if that field is found in the TF resource schema.
	Labels string `json:"labels,omitempty"`
}

type ResourceID struct {
	// TargetField is the name of the field in the TF resource where the KRM
	// `spec.resourceID` field will be mapped to.
	TargetField string `json:"targetField,omitempty"`

	// ValueTemplate is a template by which the value of the `spec.resourceID`
	// field should be interpreted before being passed to the Terraform
	// provider.
	// {{value}} is used in place of the source value, i.e. the value of
	// `spec.resourceID`.
	//
	// E.g. If `spec.resourceID` is "foo", a ValueTemplate of
	// "resources/{{value}}" means the final value passed to the Terraform
	// provider is "resources/foo".
	ValueTemplate string `json:"valueTemplate,omitempty"`
}

type ReferenceConfig struct {
	// The inlined type configuration for this reference. Must not be filled
	// out if Types is set.
	TypeConfig `json:",inline"`

	// TFField is the path to the field in the underlying Terraform provider that is
	// the implicit reference. Use periods to delimit the fields in the path. For
	// example, if the reference field is "bar" nested inside "foo" ("foo" being
	// either an object or a list of objects), the associated TFField should be
	// "foo.bar")
	TFField string `json:"tfField"`

	// Description is the description for the resource reference that will be
	// exposed in the CRD.
	Description string `json:"description,omitempty"`

	// Types is the supported types this resource reference supports. Must not
	// be specified if the inlined TypeConfig is filled out.
	//
	// If the value for the reference is not specified in the KRM spec, it is
	// possible that a default value may be set by GCP. This default reference
	// value will be populated in the KRM resource's spec. In cases where a
	// resource reference has multiple types, the first type in this list will
	// become the default TypeConfig for that value.
	Types []TypeConfig `json:"types,omitempty"`
}

type TypeConfig struct {
	// Key is the field name that will be exposed through the KRM resource's spec. It
	// should follow the Kubernetes reference naming semantics:
	//   `fooRef`, where foo is some describer of what is being referenced (e.g.
	//   instanceRef, healthCheckRef)
	// Complex references (those with a "Types" list defined) or lists of references
	// should not specify a key.
	Key string `json:"key,omitempty"`

	// TargetField is the referenced resource's Terraform field that will
	// be extracted and set as the value of the TFField. For example, a
	// ComputeSubnetwork can reference a ComputeNetwork's self link by
	// setting TargetField to "self_link", a field defined on the
	// google_compute_network resource.
	TargetField string `json:"targetField,omitempty"`

	// GVK is the Group,Version,Kind of the resource being referenced.
	//
	// This field is mutually exclusive with JSONSchemaType.
	GVK schema.GroupVersionKind `json:"gvk,omitempty"`

	// Parent specifies whether the referenced resource is a parent. If the parent
	// is successfully deleted, this resource may be deleted without any call to the
	// underlying API. Only one parent may be present. A parent reference's TFField
	// must not be a nested path.
	Parent bool `json:"parent,omitempty"`

	// JSONSchemaType specifies the type as understood by JSON schema validation of this
	// reference field. Should never be specified for a TypeConfig inlined in the
	// ReferenceConfig.
	//
	// This field is mutually exclusive with Kind and TargetField.
	JSONSchemaType string `json:"jsonSchemaType,omitempty"`

	// ValueTemplate is a template by which the value sourced from the reference should
	// be interpreted before being passed to the Terraform provider. {{value}} is used
	// in place of this sourced value. The template can contain other value placeholders
	// that need to be sourced from the reference resource.
	//
	// e.g. If the value sourced from the reference is "foo@domain.com", a valueTemplate
	// of "serviceAccount:{{value}}" would mean the final value passed to the provider
	// is "serviceAccount:foo@domain.com"
	// e.g. If the template is "projects/{{project}}/topics/{{value}}", the project value
	// will be sourced from the referenced resource.
	ValueTemplate string `json:"valueTemplate,omitempty"`

	// DCLBasedResource specifies whether or not the referenced resource is a DCL-based
	// resource. If this value is omitted, it is assumed to be false, and will only be
	// included with a value of `true` in the case that the referenced resource is a
	// DCL-based resource.
	DCLBasedResource bool `json:"dclBasedResource,omitempty"`
}

type IAMMemberReferenceConfig struct {
	// TargetField is the referenced resource's Terraform field that will be
	// extracted and used as a member identity. For example, a LoggingLogSink's
	// writer identity can be referenced by setting TargetField to
	// "writer_identity", a field defined on the google_logging_log_sink
	// resource.
	TargetField string `json:"targetField"`

	// ValueTemplate is a template by which the value sourced from the reference should
	// be interpreted before being used as a member identity. {{value}} is used
	// in place of this sourced value. The template can contain other value placeholders
	// that need to be sourced from the reference resource.
	//
	// e.g. If the value sourced from the reference is "foo@domain.com", a
	// valueTemplate of "serviceAccount:{{value}}" would mean the final value
	// used as a member identity is "serviceAccount:foo@domain.com"
	// e.g. If the template is "projects/{{project}}/topics/{{value}}", the
	// project value will be sourced from the referenced resource.
	ValueTemplate string `json:"valueTemplate,omitempty"`
}

type ContainerType string

// The following constants are the valid container types.
const (
	ContainerTypeProject      = "project"
	ContainerTypeFolder       = "folder"
	ContainerTypeOrganization = "organization"
)

type Container struct {
	// Type is the type of container this represents.
	Type ContainerType `json:"type"`

	// TFField is the path to the field in the underlying Terraform provider that
	// represents the implicit reference to the container object. Use periods to delimit
	// the fields in the path. For example, if the field is "bar" nested inside "foo" ("foo"
	// being either an object or a list of objects), the associated TFField should be
	// "foo.bar")
	TFField string `json:"tfField"`

	// ValueTemplate is a template by which the value of the container annotation
	// should be interpreted before being passed to the Terraform provider. {{value}}
	// is used in place of this sourced value.
	//
	// e.g. If the value sourced from the container annotation is "123456789", a
	// valueTemplate of "folders/{{value}}" would mean the final value passed to the
	// provider is "folders/123456789"
	ValueTemplate string `json:"valueTemplate,omitempty"`
}

type HierarchicalReferenceType string

// The following constants are the valid hierarchical reference types.
const (
	HierarchicalReferenceTypeProject        = HierarchicalReferenceType("project")
	HierarchicalReferenceTypeFolder         = HierarchicalReferenceType("folder")
	HierarchicalReferenceTypeOrganization   = HierarchicalReferenceType("organization")
	HierarchicalReferenceTypeBillingAccount = HierarchicalReferenceType("billingAccount")
)

type HierarchicalReference struct {
	// Type is the type of hierarchical reference that this hierarchical
	// reference configuration represents.
	Type HierarchicalReferenceType `json:"type"`

	// Key is the field name of the resource reference that this hierarchical
	// reference configuration corresponds to (e.g. "projectRef"). It is
	// assumed that all resource references marked as hierarchical references
	// are located at the root-level of the spec.
	Key string `json:"key"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceMapping is configuration mapping Terraform resources to Kubernetes Resource Model
// (KRM). It is used both during initial CRD generation as well as calling the Terraform
// provider at runtime.
// +k8s:openapi-gen=true
type ServiceMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// ServiceMappingSpec defines the aspects common to all resources of a particular
	// service being mapped from the Terraform provider to Kubernetes Resource Model (KRM).
	Spec ServiceMappingSpec `json:"spec,omitempty"`
}

func (sm *ServiceMapping) GetVersionFor(rc *ResourceConfig) string {
	version := sm.Spec.Version
	if rc.Version != nil {
		version = *rc.Version
	}
	return version
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceMappingList contains a list of ServiceMapping
type ServiceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceMapping `json:"items"`
}
