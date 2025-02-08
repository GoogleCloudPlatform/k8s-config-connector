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


// +kcc:proto=google.cloud.shell.v1.Environment
type Environment struct {
	// Immutable. Full name of this resource, in the format
	//  `users/{owner_email}/environments/{environment_id}`. `{owner_email}` is the
	//  email address of the user to whom this environment belongs, and
	//  `{environment_id}` is the identifier of this environment. For example,
	//  `users/someone@example.com/environments/default`.
	// +kcc:proto:field=google.cloud.shell.v1.Environment.name
	Name *string `json:"name,omitempty"`

	// Required. Immutable. Full path to the Docker image used to run this environment, e.g.
	//  "gcr.io/dev-con/cloud-devshell:latest".
	// +kcc:proto:field=google.cloud.shell.v1.Environment.docker_image
	DockerImage *string `json:"dockerImage,omitempty"`
}

// +kcc:proto=google.cloud.shell.v1.Environment
type EnvironmentObservedState struct {
	// Output only. The environment's identifier, unique among the user's
	//  environments.
	// +kcc:proto:field=google.cloud.shell.v1.Environment.id
	ID *string `json:"id,omitempty"`

	// Output only. Current execution state of this environment.
	// +kcc:proto:field=google.cloud.shell.v1.Environment.state
	State *string `json:"state,omitempty"`

	// Output only. Host to which clients can connect to initiate HTTPS or WSS
	//  connections with the environment.
	// +kcc:proto:field=google.cloud.shell.v1.Environment.web_host
	WebHost *string `json:"webHost,omitempty"`

	// Output only. Username that clients should use when initiating SSH sessions
	//  with the environment.
	// +kcc:proto:field=google.cloud.shell.v1.Environment.ssh_username
	SSHUsername *string `json:"sshUsername,omitempty"`

	// Output only. Host to which clients can connect to initiate SSH sessions
	//  with the environment.
	// +kcc:proto:field=google.cloud.shell.v1.Environment.ssh_host
	SSHHost *string `json:"sshHost,omitempty"`

	// Output only. Port to which clients can connect to initiate SSH sessions
	//  with the environment.
	// +kcc:proto:field=google.cloud.shell.v1.Environment.ssh_port
	SSHPort *int32 `json:"sshPort,omitempty"`

	// Output only. Public keys associated with the environment. Clients can
	//  connect to this environment via SSH only if they possess a private key
	//  corresponding to at least one of these public keys. Keys can be added to or
	//  removed from the environment using the AddPublicKey and RemovePublicKey
	//  methods.
	// +kcc:proto:field=google.cloud.shell.v1.Environment.public_keys
	PublicKeys []string `json:"publicKeys,omitempty"`
}
