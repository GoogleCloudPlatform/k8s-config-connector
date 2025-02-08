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


// +kcc:proto=google.cloud.orchestration.airflow.service.v1.UserWorkloadsSecret
type UserWorkloadsSecret struct {
	// Identifier. The resource name of the Secret, in the form:
	//  "projects/{projectId}/locations/{locationId}/environments/{environmentId}/userWorkloadsSecrets/{userWorkloadsSecretId}"
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.UserWorkloadsSecret.name
	Name *string `json:"name,omitempty"`

	// Optional. The "data" field of Kubernetes Secret, organized in key-value
	//  pairs, which can contain sensitive values such as a password, a token, or a
	//  key. The values for all keys have to be base64-encoded strings. For details
	//  see: https://kubernetes.io/docs/concepts/configuration/secret/
	//
	//  Example:
	//
	//  {
	//    "example": "ZXhhbXBsZV92YWx1ZQ==",
	//    "another-example": "YW5vdGhlcl9leGFtcGxlX3ZhbHVl"
	//  }
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1.UserWorkloadsSecret.data
	Data map[string]string `json:"data,omitempty"`
}
