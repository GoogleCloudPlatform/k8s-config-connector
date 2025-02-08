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


// +kcc:proto=google.cloud.orchestration.airflow.service.v1beta1.UserWorkloadsConfigMap
type UserWorkloadsConfigMap struct {
	// Identifier. The resource name of the ConfigMap, in the form:
	//  "projects/{projectId}/locations/{locationId}/environments/{environmentId}/userWorkloadsConfigMaps/{userWorkloadsConfigMapId}"
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1beta1.UserWorkloadsConfigMap.name
	Name *string `json:"name,omitempty"`

	// Optional. The "data" field of Kubernetes ConfigMap, organized in key-value
	//  pairs. For details see:
	//  https://kubernetes.io/docs/concepts/configuration/configmap/
	//
	//  Example:
	//
	//  {
	//    "example_key": "example_value",
	//    "another_key": "another_value"
	//  }
	// +kcc:proto:field=google.cloud.orchestration.airflow.service.v1beta1.UserWorkloadsConfigMap.data
	Data map[string]string `json:"data,omitempty"`
}
