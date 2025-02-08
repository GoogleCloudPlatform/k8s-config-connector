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

package v1alpha1


// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeter
type ServicePerimeter struct {
	// Required. Resource name for the ServicePerimeter.  The `short_name`
	//  component must begin with a letter and only include alphanumeric and '_'.
	//  Format:
	//  `accessPolicies/{access_policy}/servicePerimeters/{service_perimeter}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.name
	Name *string `json:"name,omitempty"`

	// Human readable title. Must be unique within the Policy.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.title
	Title *string `json:"title,omitempty"`

	// Description of the `ServicePerimeter` and its use. Does not affect
	//  behavior.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.description
	Description *string `json:"description,omitempty"`

	// Output only. Time the `ServicePerimeter` was created in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time the `ServicePerimeter` was updated in UTC.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Perimeter type indicator. A single project is
	//  allowed to be a member of single regular perimeter, but multiple service
	//  perimeter bridges. A project cannot be a included in a perimeter bridge
	//  without being included in regular perimeter. For perimeter bridges,
	//  the restricted service list as well as access level lists must be
	//  empty.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.perimeter_type
	PerimeterType *string `json:"perimeterType,omitempty"`

	// Current ServicePerimeter configuration. Specifies sets of resources,
	//  restricted services and access levels that determine perimeter
	//  content and boundaries.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.status
	Status *ServicePerimeterConfig `json:"status,omitempty"`

	// Proposed (or dry run) ServicePerimeter configuration. This configuration
	//  allows to specify and test ServicePerimeter configuration without enforcing
	//  actual access restrictions. Only allowed to be set when the
	//  "use_explicit_dry_run_spec" flag is set.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.spec
	Spec *ServicePerimeterConfig `json:"spec,omitempty"`

	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly
	//  exists  for all Service Perimeters, and that spec is identical to the
	//  status for those Service Perimeters. When this flag is set, it inhibits the
	//  generation of the implicit spec, thereby allowing the user to explicitly
	//  provide a configuration ("spec") to use in a dry-run version of the Service
	//  Perimeter. This allows the user to test changes to the enforced config
	//  ("status") without actually enforcing them. This testing is done through
	//  analyzing the differences between currently enforced and suggested
	//  restrictions. use_explicit_dry_run_spec must bet set to True if any of the
	//  fields in the spec are set to non-default values.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeter.use_explicit_dry_run_spec
	UseExplicitDryRunSpec *bool `json:"useExplicitDryRunSpec,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig
type ServicePerimeterConfig struct {
	// A list of Google Cloud resources that are inside of the service perimeter.
	//  Currently only projects are allowed. Format: `projects/{project_number}`
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.resources
	Resources []string `json:"resources,omitempty"`

	// A list of `AccessLevel` resource names that allow resources within the
	//  `ServicePerimeter` to be accessed from the internet. `AccessLevels` listed
	//  must be in the same policy as this `ServicePerimeter`. Referencing a
	//  nonexistent `AccessLevel` is a syntax error. If no `AccessLevel` names are
	//  listed, resources within the perimeter can only be accessed via Google
	//  Cloud calls with request origins within the perimeter. Example:
	//  `"accessPolicies/MY_POLICY/accessLevels/MY_LEVEL"`.
	//  For Service Perimeter Bridge, must be empty.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.access_levels
	AccessLevels []string `json:"accessLevels,omitempty"`

	// Google Cloud services that are subject to the Service Perimeter
	//  restrictions. For example, if `storage.googleapis.com` is specified, access
	//  to the storage buckets inside the perimeter must meet the perimeter's
	//  access restrictions.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.restricted_services
	RestrictedServices []string `json:"restrictedServices,omitempty"`

	// Configuration for APIs allowed within Perimeter.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.vpc_accessible_services
	VpcAccessibleServices *ServicePerimeterConfig_VpcAccessibleServices `json:"vpcAccessibleServices,omitempty"`

	// List of [IngressPolicies]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  to apply to the perimeter. A perimeter may have multiple [IngressPolicies]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy],
	//  each of which is evaluated separately. Access is granted if any [Ingress
	//  Policy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  grants it. Must be empty for a perimeter bridge.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ingress_policies
	IngressPolicies []ServicePerimeterConfig_IngressPolicy `json:"ingressPolicies,omitempty"`

	// List of [EgressPolicies]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  to apply to the perimeter. A perimeter may have multiple [EgressPolicies]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy],
	//  each of which is evaluated separately. Access is granted if any
	//  [EgressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  grants it. Must be empty for a perimeter bridge.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.egress_policies
	EgressPolicies []ServicePerimeterConfig_EgressPolicy `json:"egressPolicies,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation
type ServicePerimeterConfig_ApiOperation struct {
	// The name of the API whose methods or permissions the [IngressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  or [EgressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  want to allow. A single [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  with `service_name` field set to `*` will allow all methods AND
	//  permissions for all services.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation.service_name
	ServiceName *string `json:"serviceName,omitempty"`

	// API methods or permissions to allow. Method or permission must belong to
	//  the service specified by `service_name` field. A single [MethodSelector]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.MethodSelector]
	//  entry with `*` specified for the `method` field will allow all methods
	//  AND permissions for the service specified in `service_name`.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation.method_selectors
	MethodSelectors []ServicePerimeterConfig_MethodSelector `json:"methodSelectors,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom
type ServicePerimeterConfig_EgressFrom struct {
	// A list of identities that are allowed access through this [EgressPolicy].
	//  Should be in the format of email address. The email address should
	//  represent individual user or service account only.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom.identities
	Identities []string `json:"identities,omitempty"`

	// Specifies the type of identities that are allowed access to outside the
	//  perimeter. If left unspecified, then members of `identities` field will
	//  be allowed access.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom.identity_type
	IdentityType *string `json:"identityType,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy
type ServicePerimeterConfig_EgressPolicy struct {
	// Defines conditions on the source of a request causing this [EgressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  to apply.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy.egress_from
	EgressFrom *ServicePerimeterConfig_EgressFrom `json:"egressFrom,omitempty"`

	// Defines the conditions on the [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  and destination resources that cause this [EgressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy]
	//  to apply.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressPolicy.egress_to
	EgressTo *ServicePerimeterConfig_EgressTo `json:"egressTo,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo
type ServicePerimeterConfig_EgressTo struct {
	// A list of resources, currently only projects in the form
	//  `projects/<projectnumber>`, that are allowed to be accessed by sources
	//  defined in the corresponding [EgressFrom]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom].
	//  A request matches if it contains a resource in this list.  If `*` is
	//  specified for `resources`, then this [EgressTo]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo]
	//  rule will authorize access to all resources outside the perimeter.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo.resources
	Resources []string `json:"resources,omitempty"`

	// A list of [ApiOperations]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  allowed to be performed by the sources specified in the corresponding
	//  [EgressFrom]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressFrom].
	//  A request matches if it uses an operation/service in this list.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo.operations
	Operations []ServicePerimeterConfig_ApiOperation `json:"operations,omitempty"`

	// A list of external resources that are allowed to be accessed. Only AWS
	//  and Azure resources are supported. For Amazon S3, the supported format is
	//  s3://BUCKET_NAME. For Azure Storage, the supported format is
	//  azure://myaccount.blob.core.windows.net/CONTAINER_NAME. A request matches
	//  if it contains an external resource in this list (Example:
	//  s3://bucket/path). Currently '*' is not allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.EgressTo.external_resources
	ExternalResources []string `json:"externalResources,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom
type ServicePerimeterConfig_IngressFrom struct {
	// Sources that this [IngressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  authorizes access from.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom.sources
	Sources []ServicePerimeterConfig_IngressSource `json:"sources,omitempty"`

	// A list of identities that are allowed access through this ingress
	//  policy. Should be in the format of email address. The email address
	//  should represent individual user or service account only.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom.identities
	Identities []string `json:"identities,omitempty"`

	// Specifies the type of identities that are allowed access from outside the
	//  perimeter. If left unspecified, then members of `identities` field will
	//  be allowed access.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom.identity_type
	IdentityType *string `json:"identityType,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy
type ServicePerimeterConfig_IngressPolicy struct {
	// Defines the conditions on the source of a request causing this
	//  [IngressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  to apply.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy.ingress_from
	IngressFrom *ServicePerimeterConfig_IngressFrom `json:"ingressFrom,omitempty"`

	// Defines the conditions on the [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  and request destination that cause this [IngressPolicy]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy]
	//  to apply.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressPolicy.ingress_to
	IngressTo *ServicePerimeterConfig_IngressTo `json:"ingressTo,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressSource
type ServicePerimeterConfig_IngressSource struct {
	// An [AccessLevel]
	//  [google.identity.accesscontextmanager.v1.AccessLevel] resource
	//  name that allow resources within the [ServicePerimeters]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeter] to be
	//  accessed from the internet. [AccessLevels]
	//  [google.identity.accesscontextmanager.v1.AccessLevel] listed must
	//  be in the same policy as this [ServicePerimeter]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeter].
	//  Referencing a nonexistent [AccessLevel]
	//  [google.identity.accesscontextmanager.v1.AccessLevel] will cause
	//  an error. If no [AccessLevel]
	//  [google.identity.accesscontextmanager.v1.AccessLevel] names are
	//  listed, resources within the perimeter can only be accessed via Google
	//  Cloud calls with request origins within the perimeter. Example:
	//  `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL`. If a single `*` is
	//  specified for `access_level`, then all [IngressSources]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressSource]
	//  will be allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressSource.access_level
	AccessLevel *string `json:"accessLevel,omitempty"`

	// A Google Cloud resource that is allowed to ingress the perimeter.
	//  Requests from these resources will be allowed to access perimeter data.
	//  Currently only projects are allowed.
	//  Format: `projects/{project_number}`
	//  The project may be in any Google Cloud organization, not just the
	//  organization that the perimeter is defined in. `*` is not allowed, the
	//  case of allowing all Google Cloud resources only is not supported.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressSource.resource
	Resource *string `json:"resource,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressTo
type ServicePerimeterConfig_IngressTo struct {
	// A list of [ApiOperations]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation]
	//  allowed to be performed by the sources specified in corresponding
	//  [IngressFrom]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom]
	//  in this [ServicePerimeter]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeter].
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressTo.operations
	Operations []ServicePerimeterConfig_ApiOperation `json:"operations,omitempty"`

	// A list of resources, currently only projects in the form
	//  `projects/<projectnumber>`, protected by this [ServicePerimeter]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeter] that are
	//  allowed to be accessed by sources defined in the corresponding
	//  [IngressFrom]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressFrom].
	//  If a single `*` is specified, then access to all resources inside the
	//  perimeter are allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.IngressTo.resources
	Resources []string `json:"resources,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.MethodSelector
type ServicePerimeterConfig_MethodSelector struct {
	// Value for `method` should be a valid method name for the corresponding
	//  `service_name` in [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation].
	//  If `*` used as value for `method`, then ALL methods and permissions are
	//  allowed.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.MethodSelector.method
	Method *string `json:"method,omitempty"`

	// Value for `permission` should be a valid Cloud IAM permission for the
	//  corresponding `service_name` in [ApiOperation]
	//  [google.identity.accesscontextmanager.v1.ServicePerimeterConfig.ApiOperation].
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.MethodSelector.permission
	Permission *string `json:"permission,omitempty"`
}

// +kcc:proto=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.VpcAccessibleServices
type ServicePerimeterConfig_VpcAccessibleServices struct {
	// Whether to restrict API calls within the Service Perimeter to the list of
	//  APIs specified in 'allowed_services'.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.VpcAccessibleServices.enable_restriction
	EnableRestriction *bool `json:"enableRestriction,omitempty"`

	// The list of APIs usable within the Service Perimeter. Must be empty
	//  unless 'enable_restriction' is True. You can specify a list of individual
	//  services, as well as include the 'RESTRICTED-SERVICES' value, which
	//  automatically includes all of the services protected by the perimeter.
	// +kcc:proto:field=google.identity.accesscontextmanager.v1.ServicePerimeterConfig.VpcAccessibleServices.allowed_services
	AllowedServices []string `json:"allowedServices,omitempty"`
}
