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


// +kcc:proto=google.cloud.secrets.v1beta1.Replication
type Replication struct {
	// The [Secret][google.cloud.secrets.v1beta1.Secret] will automatically be replicated without any restrictions.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.Replication.automatic
	Automatic *Replication_Automatic `json:"automatic,omitempty"`

	// The [Secret][google.cloud.secrets.v1beta1.Secret] will only be replicated into the locations specified.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.Replication.user_managed
	UserManaged *Replication_UserManaged `json:"userManaged,omitempty"`
}

// +kcc:proto=google.cloud.secrets.v1beta1.Replication.Automatic
type Replication_Automatic struct {
}

// +kcc:proto=google.cloud.secrets.v1beta1.Replication.UserManaged
type Replication_UserManaged struct {
	// Required. The list of Replicas for this [Secret][google.cloud.secrets.v1beta1.Secret].
	//
	//  Cannot be empty.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.Replication.UserManaged.replicas
	Replicas []Replication_UserManaged_Replica `json:"replicas,omitempty"`
}

// +kcc:proto=google.cloud.secrets.v1beta1.Replication.UserManaged.Replica
type Replication_UserManaged_Replica struct {
	// The canonical IDs of the location to replicate data.
	//  For example: `"us-east1"`.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.Replication.UserManaged.Replica.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.secrets.v1beta1.Secret
type Secret struct {

	// Required. Immutable. The replication policy of the secret data attached to the [Secret][google.cloud.secrets.v1beta1.Secret].
	//
	//  The replication policy cannot be changed after the Secret has been created.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.Secret.replication
	Replication *Replication `json:"replication,omitempty"`

	// The labels assigned to this Secret.
	//
	//  Label keys must be between 1 and 63 characters long, have a UTF-8 encoding
	//  of maximum 128 bytes, and must conform to the following PCRE regular
	//  expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`
	//
	//  Label values must be between 0 and 63 characters long, have a UTF-8
	//  encoding of maximum 128 bytes, and must conform to the following PCRE
	//  regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`
	//
	//  No more than 64 labels can be assigned to a given resource.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.Secret.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.secrets.v1beta1.Secret
type SecretObservedState struct {
	// Output only. The resource name of the [Secret][google.cloud.secrets.v1beta1.Secret] in the format `projects/*/secrets/*`.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.Secret.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which the [Secret][google.cloud.secrets.v1beta1.Secret] was created.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.Secret.create_time
	CreateTime *string `json:"createTime,omitempty"`
}
