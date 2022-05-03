// Copyright 2022 Google LLC
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

package dclschemaloader

import (
	"github.com/nasa9084/go-openapi"
)

var correctedDCLResourceValueTemplateMap = map[string]string{
	// TODO(b/179399619): remove following two entries once DCL fixes 'x-dcl-id' for MonitoringGroup and StorageBucket resources
	"monitoring_ga_group": "projects/{{project}}/groups/{{name}}",
	"storage_ga_bucket":   "{{name}}",
}

var unsupportedDCLResourceValueTemplateMap = map[string]string{
	// referenced by RunService
	"secretmanager_beta_secret": "projects/{{project}}/secrets/{{name}}",
	// TODO(b/200559394): Support the relative resource name.
	// {{name}} will be resolved to the short name, i.e. resourceID. It works
	// for RunService because the API accepts the short name.
	// There might be other cases where resources which reference
	// secretmanager_beta_secretversion require the long name - the relative
	// resource name, and it is not supported yet.
	"secretmanager_beta_secretversion": "{{name}}",
	"workflows_beta_workflow":          "projects/{{project}}/locations/{{location}}/workflows/{{name}}",
}

// This function patches DCL schema for KCC's convenience.
func patchDCLSchemas(schemaMap map[string]*openapi.Schema) {
	// Fix the OpenAPI schema for the resources with incorrect 'x-dcl-id's.
	for resourceName, valueTemplate := range correctedDCLResourceValueTemplateMap {
		s := schemaMap[resourceName]
		s.Extension["x-dcl-id"] = valueTemplate
	}

	// Configure 'x-dcl-id' for the resources that are:
	// (1) referenced by a DCL-based KCC resource, and
	// (2) not yet supported in DCL.
	for resourceName, valueTemplate := range unsupportedDCLResourceValueTemplateMap {
		schemaMap[resourceName] = &openapi.Schema{
			Extension: map[string]interface{}{
				"x-dcl-id": valueTemplate,
			},
		}
	}
}
