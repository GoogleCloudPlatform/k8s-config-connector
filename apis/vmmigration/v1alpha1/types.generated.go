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


// +kcc:proto=google.cloud.vmmigration.v1.AwsSourceDetails
type AwsSourceDetails struct {
	// AWS Credentials using access key id and secret.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.access_key_creds
	AccessKeyCreds *AwsSourceDetails_AccessKeyCredentials `json:"accessKeyCreds,omitempty"`

	// Immutable. The AWS region that the source VMs will be migrated from.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.aws_region
	AwsRegion *string `json:"awsRegion,omitempty"`

	// AWS resource tags to limit the scope of the source inventory.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.inventory_tag_list
	InventoryTagList []AwsSourceDetails_Tag `json:"inventoryTagList,omitempty"`

	// AWS security group names to limit the scope of the source
	//  inventory.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.inventory_security_group_names
	InventorySecurityGroupNames []string `json:"inventorySecurityGroupNames,omitempty"`

	// User specified tags to add to every M2VM generated resource in AWS.
	//  These tags will be set in addition to the default tags that are set as part
	//  of the migration process. The tags must not begin with the reserved prefix
	//  `m2vm`.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.migration_resources_user_tags
	MigrationResourcesUserTags map[string]string `json:"migrationResourcesUserTags,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.AwsSourceDetails.AccessKeyCredentials
type AwsSourceDetails_AccessKeyCredentials struct {
	// AWS access key ID.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.AccessKeyCredentials.access_key_id
	AccessKeyID *string `json:"accessKeyID,omitempty"`

	// Input only. AWS secret access key.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.AccessKeyCredentials.secret_access_key
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.AwsSourceDetails.Tag
type AwsSourceDetails_Tag struct {
	// Key of tag.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.Tag.key
	Key *string `json:"key,omitempty"`

	// Value of tag.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.Tag.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.Source
type Source struct {
	// Vmware type source details.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.vmware
	Vmware *VmwareSourceDetails `json:"vmware,omitempty"`

	// AWS type source details.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.aws
	Aws *AwsSourceDetails `json:"aws,omitempty"`

	// The labels of the source.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.labels
	Labels map[string]string `json:"labels,omitempty"`

	// User-provided description of the source.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.VmwareSourceDetails
type VmwareSourceDetails struct {
	// The credentials username.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareSourceDetails.username
	Username *string `json:"username,omitempty"`

	// Input only. The credentials password. This is write only and can not be
	//  read in a GET operation.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareSourceDetails.password
	Password *string `json:"password,omitempty"`

	// The ip address of the vcenter this Source represents.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareSourceDetails.vcenter_ip
	VcenterIP *string `json:"vcenterIP,omitempty"`

	// The thumbprint representing the certificate for the vcenter.
	// +kcc:proto:field=google.cloud.vmmigration.v1.VmwareSourceDetails.thumbprint
	Thumbprint *string `json:"thumbprint,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.AwsSourceDetails
type AwsSourceDetailsObservedState struct {
	// Output only. State of the source as determined by the health check.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.state
	State *string `json:"state,omitempty"`

	// Output only. Provides details on the state of the Source in case of an
	//  error.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.error
	Error *Status `json:"error,omitempty"`

	// Output only. The source's public IP. All communication initiated by this
	//  source will originate from this IP.
	// +kcc:proto:field=google.cloud.vmmigration.v1.AwsSourceDetails.public_ip
	PublicIP *string `json:"publicIP,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.Source
type SourceObservedState struct {
	// AWS type source details.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.aws
	Aws *AwsSourceDetailsObservedState `json:"aws,omitempty"`

	// Output only. The Source name.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.name
	Name *string `json:"name,omitempty"`

	// Output only. The create time timestamp.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time timestamp.
	// +kcc:proto:field=google.cloud.vmmigration.v1.Source.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
