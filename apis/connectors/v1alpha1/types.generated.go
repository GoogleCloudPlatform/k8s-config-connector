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


// +kcc:proto=google.cloud.connectors.v1.AuthConfigTemplate
type AuthConfigTemplate struct {
	// The type of authentication configured.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfigTemplate.auth_type
	AuthType *string `json:"authType,omitempty"`

	// Config variables to describe an `AuthConfig` for a `Connection`.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfigTemplate.config_variable_templates
	ConfigVariableTemplates []ConfigVariableTemplate `json:"configVariableTemplates,omitempty"`

	// Display name for authentication template.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfigTemplate.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Connector specific description for an authentication template.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthConfigTemplate.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.AuthorizationCodeLink
type AuthorizationCodeLink struct {
	// The base URI the user must click to trigger the authorization code login
	//  flow.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthorizationCodeLink.uri
	URI *string `json:"uri,omitempty"`

	// The scopes for which the user will authorize GCP Connectors on the
	//  connector data source.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthorizationCodeLink.scopes
	Scopes []string `json:"scopes,omitempty"`

	// The client ID assigned to the GCP Connectors OAuth app for the connector
	//  data source.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthorizationCodeLink.client_id
	ClientID *string `json:"clientID,omitempty"`

	// Whether to enable PKCE for the auth code flow.
	// +kcc:proto:field=google.cloud.connectors.v1.AuthorizationCodeLink.enable_pkce
	EnablePkce *bool `json:"enablePkce,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.ConfigVariableTemplate
type ConfigVariableTemplate struct {
	// Key of the config variable.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.key
	Key *string `json:"key,omitempty"`

	// Type of the parameter: string, int, bool etc.
	//  consider custom type for the benefit for the validation.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.value_type
	ValueType *string `json:"valueType,omitempty"`

	// Display name of the parameter.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Description.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.description
	Description *string `json:"description,omitempty"`

	// Regular expression in RE2 syntax used for validating the `value` of a
	//  `ConfigVariable`.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.validation_regex
	ValidationRegex *string `json:"validationRegex,omitempty"`

	// Flag represents that this `ConfigVariable` must be provided for a
	//  connection.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.required
	Required *bool `json:"required,omitempty"`

	// Role grant configuration for the config variable.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.role_grant
	RoleGrant *RoleGrant `json:"roleGrant,omitempty"`

	// Enum options. To be populated if `ValueType` is `ENUM`
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.enum_options
	EnumOptions []EnumOption `json:"enumOptions,omitempty"`

	// Authorization code link options. To be populated if `ValueType` is
	//  `AUTHORIZATION_CODE`
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.authorization_code_link
	AuthorizationCodeLink *AuthorizationCodeLink `json:"authorizationCodeLink,omitempty"`

	// State of the config variable.
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.state
	State *string `json:"state,omitempty"`

	// Indicates if current template is part of advanced settings
	// +kcc:proto:field=google.cloud.connectors.v1.ConfigVariableTemplate.is_advanced
	IsAdvanced *bool `json:"isAdvanced,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.ConnectorVersion
type ConnectorVersion struct {
}

// +kcc:proto=google.cloud.connectors.v1.EgressControlConfig
type EgressControlConfig struct {
	// Static Comma separated backends which are common for all Connection
	//  resources. Supported formats for each backend are host:port or just
	//  host (host can be ip address or domain name).
	// +kcc:proto:field=google.cloud.connectors.v1.EgressControlConfig.backends
	Backends *string `json:"backends,omitempty"`

	// Extractions Rules to extract the backends from customer provided
	//  configuration.
	// +kcc:proto:field=google.cloud.connectors.v1.EgressControlConfig.extraction_rules
	ExtractionRules *ExtractionRules `json:"extractionRules,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.EnumOption
type EnumOption struct {
	// Id of the option.
	// +kcc:proto:field=google.cloud.connectors.v1.EnumOption.id
	ID *string `json:"id,omitempty"`

	// Display name of the option.
	// +kcc:proto:field=google.cloud.connectors.v1.EnumOption.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.ExtractionRule
type ExtractionRule struct {
	// Source on which the rule is applied.
	// +kcc:proto:field=google.cloud.connectors.v1.ExtractionRule.source
	Source *ExtractionRule_Source `json:"source,omitempty"`

	// Regex used to extract backend details from source. If empty, whole source
	//  value will be used.
	// +kcc:proto:field=google.cloud.connectors.v1.ExtractionRule.extraction_regex
	ExtractionRegex *string `json:"extractionRegex,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.ExtractionRule.Source
type ExtractionRule_Source struct {
	// Type of the source.
	// +kcc:proto:field=google.cloud.connectors.v1.ExtractionRule.Source.source_type
	SourceType *string `json:"sourceType,omitempty"`

	// Field identifier. For example config vaiable name.
	// +kcc:proto:field=google.cloud.connectors.v1.ExtractionRule.Source.field_id
	FieldID *string `json:"fieldID,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.ExtractionRules
type ExtractionRules struct {
	// Collection of Extraction Rule.
	// +kcc:proto:field=google.cloud.connectors.v1.ExtractionRules.extraction_rule
	ExtractionRule []ExtractionRule `json:"extractionRule,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.RoleGrant
type RoleGrant struct {
	// Prinicipal/Identity for whom the role need to assigned.
	// +kcc:proto:field=google.cloud.connectors.v1.RoleGrant.principal
	Principal *string `json:"principal,omitempty"`

	// List of roles that need to be granted.
	// +kcc:proto:field=google.cloud.connectors.v1.RoleGrant.roles
	Roles []string `json:"roles,omitempty"`

	// Resource on which the roles needs to be granted for the principal.
	// +kcc:proto:field=google.cloud.connectors.v1.RoleGrant.resource
	Resource *RoleGrant_Resource `json:"resource,omitempty"`

	// Template that UI can use to provide helper text to customers.
	// +kcc:proto:field=google.cloud.connectors.v1.RoleGrant.helper_text_template
	HelperTextTemplate *string `json:"helperTextTemplate,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.RoleGrant.Resource
type RoleGrant_Resource struct {
	// Different types of resource supported.
	// +kcc:proto:field=google.cloud.connectors.v1.RoleGrant.Resource.type
	Type *string `json:"type,omitempty"`

	// Template to uniquely represent a GCP resource in a format IAM expects
	//  This is a template that can have references to other values provided in
	//  the config variable template.
	// +kcc:proto:field=google.cloud.connectors.v1.RoleGrant.Resource.path_template
	PathTemplate *string `json:"pathTemplate,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.SslConfigTemplate
type SslConfigTemplate struct {
	// Controls the ssl type for the given connector version
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfigTemplate.ssl_type
	SslType *string `json:"sslType,omitempty"`

	// Boolean for determining if the connector version mandates TLS.
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfigTemplate.is_tls_mandatory
	IsTlsMandatory *bool `json:"isTlsMandatory,omitempty"`

	// List of supported Server Cert Types
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfigTemplate.server_cert_type
	ServerCertType []string `json:"serverCertType,omitempty"`

	// List of supported Client Cert Types
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfigTemplate.client_cert_type
	ClientCertType []string `json:"clientCertType,omitempty"`

	// Any additional fields that need to be rendered
	// +kcc:proto:field=google.cloud.connectors.v1.SslConfigTemplate.additional_variables
	AdditionalVariables []ConfigVariableTemplate `json:"additionalVariables,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.SupportedRuntimeFeatures
type SupportedRuntimeFeatures struct {
	// Specifies if the connector supports entity apis like 'createEntity'.
	// +kcc:proto:field=google.cloud.connectors.v1.SupportedRuntimeFeatures.entity_apis
	EntityApis *bool `json:"entityApis,omitempty"`

	// Specifies if the connector supports action apis like 'executeAction'.
	// +kcc:proto:field=google.cloud.connectors.v1.SupportedRuntimeFeatures.action_apis
	ActionApis *bool `json:"actionApis,omitempty"`

	// Specifies if the connector supports 'ExecuteSqlQuery' operation.
	// +kcc:proto:field=google.cloud.connectors.v1.SupportedRuntimeFeatures.sql_query
	SqlQuery *bool `json:"sqlQuery,omitempty"`
}

// +kcc:proto=google.cloud.connectors.v1.ConnectorVersion
type ConnectorVersionObservedState struct {
	// Output only. Resource name of the Version.
	//  Format:
	//  projects/{project}/locations/{location}/providers/{provider}/connectors/{connector}/versions/{version}
	//  Only global location is supported for Connector resource.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.name
	Name *string `json:"name,omitempty"`

	// Output only. Created time.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Updated time.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Resource labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Flag to mark the version indicating the launch stage.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`

	// Output only. ReleaseVersion of the connector, for example: "1.0.1-alpha".
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.release_version
	ReleaseVersion *string `json:"releaseVersion,omitempty"`

	// Output only. List of auth configs supported by the Connector Version.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.auth_config_templates
	AuthConfigTemplates []AuthConfigTemplate `json:"authConfigTemplates,omitempty"`

	// Output only. List of config variables needed to create a connection.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.config_variable_templates
	ConfigVariableTemplates []ConfigVariableTemplate `json:"configVariableTemplates,omitempty"`

	// Output only. Information about the runtime features supported by the
	//  Connector.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.supported_runtime_features
	SupportedRuntimeFeatures *SupportedRuntimeFeatures `json:"supportedRuntimeFeatures,omitempty"`

	// Output only. Display name.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Configuration for Egress Control.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.egress_control_config
	EgressControlConfig *EgressControlConfig `json:"egressControlConfig,omitempty"`

	// Output only. Role grant configurations for this connector version.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.role_grants
	RoleGrants []RoleGrant `json:"roleGrants,omitempty"`

	// Output only. Role grant configuration for this config variable. It will be
	//  DEPRECATED soon.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.role_grant
	RoleGrant *RoleGrant `json:"roleGrant,omitempty"`

	// Output only. Ssl configuration supported by the Connector.
	// +kcc:proto:field=google.cloud.connectors.v1.ConnectorVersion.ssl_config_template
	SslConfigTemplate *SslConfigTemplate `json:"sslConfigTemplate,omitempty"`
}
