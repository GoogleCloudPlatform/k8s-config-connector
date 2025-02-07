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


// +kcc:proto=google.cloud.baremetalsolution.v2.SSHKey
type SSHKey struct {

	// The public SSH key. This must be in OpenSSH .authorized_keys format.
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.SSHKey.public_key
	PublicKey *string `json:"publicKey,omitempty"`
}

// +kcc:proto=google.cloud.baremetalsolution.v2.SSHKey
type SSHKeyObservedState struct {
	// Output only. The name of this SSH key.
	//  Currently, the only valid value for the location is "global".
	// +kcc:proto:field=google.cloud.baremetalsolution.v2.SSHKey.name
	Name *string `json:"name,omitempty"`
}
