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


// +kcc:proto=google.cloud.speech.v2.CustomClass
type CustomClass struct {

	// Optional. User-settable, human-readable name for the CustomClass. Must be
	//  63 characters or less.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A collection of class items.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.items
	Items []CustomClass_ClassItem `json:"items,omitempty"`

	// Optional. Allows users to store small amounts of arbitrary data.
	//  Both the key and the value must be 63 characters or less each.
	//  At most 100 annotations.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.CustomClass.ClassItem
type CustomClass_ClassItem struct {
	// The class item's value.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.ClassItem.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.speech.v2.CustomClass
type CustomClassObservedState struct {
	// Output only. Identifier. The resource name of the CustomClass.
	//  Format:
	//  `projects/{project}/locations/{location}/customClasses/{custom_class}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.name
	Name *string `json:"name,omitempty"`

	// Output only. System-assigned unique identifier for the CustomClass.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The CustomClass lifecycle state.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.state
	State *string `json:"state,omitempty"`

	// Output only. Creation time.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The most recent time this resource was modified.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time at which this resource was requested for deletion.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. The time at which this resource will be purged.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields. This may be sent on update, undelete, and delete requests to
	//  ensure the client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Whether or not this CustomClass is in the process of being
	//  updated.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The [KMS key
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which
	//  the CustomClass is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// Output only. The [KMS key version
	//  name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions)
	//  with which the CustomClass is encrypted. The expected format is
	//  `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
	// +kcc:proto:field=google.cloud.speech.v2.CustomClass.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`
}
