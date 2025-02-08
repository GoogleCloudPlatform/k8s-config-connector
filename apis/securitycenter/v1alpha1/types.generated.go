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


// +kcc:proto=google.cloud.securitycenter.v2.ResourceValueConfig
type ResourceValueConfig struct {
	// Identifier. Name for the resource value configuration
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.name
	Name *string `json:"name,omitempty"`

	// Resource value level this expression represents
	//  Only required when there is no Sensitive Data Protection mapping in the
	//  request
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.resource_value
	ResourceValue *string `json:"resourceValue,omitempty"`

	// Tag values combined with `AND` to check against.
	//  Values in the form "tagValues/123"
	//  Example: `[ "tagValues/123", "tagValues/456", "tagValues/789" ]`
	//  https://cloud.google.com/resource-manager/docs/tags/tags-creating-and-managing
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.tag_values
	TagValues []string `json:"tagValues,omitempty"`

	// Apply resource_value only to resources that match resource_type.
	//  resource_type will be checked with `AND` of other resources.
	//  For example, "storage.googleapis.com/Bucket" with resource_value "HIGH"
	//  will apply "HIGH" value only to "storage.googleapis.com/Bucket" resources.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// Project or folder to scope this configuration to.
	//  For example, "project/456" would apply this configuration only to resources
	//  in "project/456" scope and will be checked with `AND` of other resources.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.scope
	Scope *string `json:"scope,omitempty"`

	// List of resource labels to search for, evaluated with `AND`.
	//  For example, "resource_labels_selector": {"key": "value", "env": "prod"}
	//  will match resources with labels "key": "value" `AND` "env":
	//  "prod"
	//  https://cloud.google.com/resource-manager/docs/creating-managing-labels
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.resource_labels_selector
	ResourceLabelsSelector map[string]string `json:"resourceLabelsSelector,omitempty"`

	// Description of the resource value configuration.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.description
	Description *string `json:"description,omitempty"`

	// Cloud provider this configuration applies to
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.cloud_provider
	CloudProvider *string `json:"cloudProvider,omitempty"`

	// A mapping of the sensitivity on Sensitive Data Protection finding to
	//  resource values. This mapping can only be used in combination with a
	//  resource_type that is related to BigQuery, e.g.
	//  "bigquery.googleapis.com/Dataset".
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.sensitive_data_protection_mapping
	SensitiveDataProtectionMapping *ResourceValueConfig_SensitiveDataProtectionMapping `json:"sensitiveDataProtectionMapping,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.ResourceValueConfig.SensitiveDataProtectionMapping
type ResourceValueConfig_SensitiveDataProtectionMapping struct {
	// Resource value mapping for high-sensitivity Sensitive Data Protection
	//  findings
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.SensitiveDataProtectionMapping.high_sensitivity_mapping
	HighSensitivityMapping *string `json:"highSensitivityMapping,omitempty"`

	// Resource value mapping for medium-sensitivity Sensitive Data Protection
	//  findings
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.SensitiveDataProtectionMapping.medium_sensitivity_mapping
	MediumSensitivityMapping *string `json:"mediumSensitivityMapping,omitempty"`
}

// +kcc:proto=google.cloud.securitycenter.v2.ResourceValueConfig
type ResourceValueConfigObservedState struct {
	// Output only. Timestamp this resource value configuration was created.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp this resource value configuration was last updated.
	// +kcc:proto:field=google.cloud.securitycenter.v2.ResourceValueConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
