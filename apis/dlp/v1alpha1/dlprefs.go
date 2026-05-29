// Copyright 2026 Google LLC
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

type DLPInspectTemplateRef struct {
	// A reference to an externally managed DLP Inspect Template.
	// Should be in the format `projects/[project_id]/locations/[location]/inspectTemplates/[inspect_template]`.
	External string `json:"external,omitempty"`

	// The `name` of a `DLPInspectTemplate` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `DLPInspectTemplate` resource.
	Namespace string `json:"namespace,omitempty"`
}

type DLPDeidentifyTemplateRef struct {
	// A reference to an externally managed DLP Deidentify Template.
	// Should be in the format `projects/[project_id]/locations/[location]/deidentifyTemplates/[deidentify_template]`.
	External string `json:"external,omitempty"`

	// The `name` of a `DLPDeidentifyTemplate` resource.
	Name string `json:"name,omitempty"`

	// The `namespace` of a `DLPDeidentifyTemplate` resource.
	Namespace string `json:"namespace,omitempty"`
}
