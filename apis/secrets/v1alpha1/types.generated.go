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


// +kcc:proto=google.cloud.secrets.v1beta1.SecretVersion
type SecretVersion struct {
}

// +kcc:proto=google.cloud.secrets.v1beta1.SecretVersion
type SecretVersionObservedState struct {
	// Output only. The resource name of the [SecretVersion][google.cloud.secrets.v1beta1.SecretVersion] in the
	//  format `projects/*/secrets/*/versions/*`.
	//
	//  [SecretVersion][google.cloud.secrets.v1beta1.SecretVersion] IDs in a [Secret][google.cloud.secrets.v1beta1.Secret] start at 1 and
	//  are incremented for each subsequent version of the secret.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.SecretVersion.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which the [SecretVersion][google.cloud.secrets.v1beta1.SecretVersion] was created.
	// +kcc:proto:field=google.cloud.secrets.v1beta1.SecretVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time this [SecretVersion][google.cloud.secrets.v1beta1.SecretVersion] was destroyed.
	//  Only present if [state][google.cloud.secrets.v1beta1.SecretVersion.state] is
	//  [DESTROYED][google.cloud.secrets.v1beta1.SecretVersion.State.DESTROYED].
	// +kcc:proto:field=google.cloud.secrets.v1beta1.SecretVersion.destroy_time
	DestroyTime *string `json:"destroyTime,omitempty"`

	// Output only. The current state of the [SecretVersion][google.cloud.secrets.v1beta1.SecretVersion].
	// +kcc:proto:field=google.cloud.secrets.v1beta1.SecretVersion.state
	State *string `json:"state,omitempty"`
}
