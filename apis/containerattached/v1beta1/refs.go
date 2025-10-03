// Copyright 2024 Google LLC
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

package v1beta1

type FleetProjectRef struct {
	/* The project of the fleet. Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).*/
	External string `json:"external,omitempty"`
	/* Name of the project resource. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the project resource. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}
