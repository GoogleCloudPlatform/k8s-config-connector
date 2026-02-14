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
	// referenced by DLPStoredInfoType
	"bigquery_beta_dataset": "projects/{{project}}/datasets/{{name}}",
	"bigquery_beta_table":   "projects/{{project}}/datasets/{{dataset_id}}/tables/{{name}}",
	// referenced by DLPDeidentifyTemplate
	"cloudkms_ga_cryptokey": "{{selfLink}}",
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
	patchNetworkServicesHTTPRoute(schemaMap)
}

func patchNetworkServicesHTTPRoute(schemaMap map[string]*openapi.Schema) {
	s := schemaMap["networkservices_ga_httproute"]
	if s == nil {
		return
	}

	// Path 1: rules.items.properties.action.properties.destinations.items.properties.serviceName
	rules := s.Properties["rules"]
	if rules != nil && rules.Items != nil {
		action := rules.Items.Properties["action"]
		if action != nil {
			destinations := action.Properties["destinations"]
			if destinations != nil && destinations.Items != nil {
				serviceName := destinations.Items.Properties["serviceName"]
				addStorageBucketRef(serviceName)
			}

			// Path 2: rules.items.properties.action.properties.requestMirrorPolicy.properties.destination.properties.serviceName
			mirror := action.Properties["requestMirrorPolicy"]
			if mirror != nil {
				destination := mirror.Properties["destination"]
				if destination != nil {
					serviceName := destination.Properties["serviceName"]
					addStorageBucketRef(serviceName)
				}
			}
		}
	}
}

func addStorageBucketRef(s *openapi.Schema) {
	if s == nil {
		return
	}
	refs, ok := s.Extension["x-dcl-references"]
	if !ok {
		return
	}
	refList, ok := refs.([]interface{})
	if !ok {
		return
	}

	// Check if already exists (idempotency)
	for _, r := range refList {
		m, ok := r.(map[interface{}]interface{})
		if !ok {
			continue
		}
		if m["resource"] == "Storage/Bucket" {
			return
		}
	}

	// Add Storage/Bucket
	newRef := map[string]interface{}{
		"resource": "Storage/Bucket",
		"field":    "name",
	}
	s.Extension["x-dcl-references"] = append(refList, newRef)
}
