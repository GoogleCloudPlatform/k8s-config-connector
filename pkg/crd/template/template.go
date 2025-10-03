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

package crdtemplate

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"gopkg.in/yaml.v2"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func SpecToYAML(crd *apiextensions.CustomResourceDefinition) ([]byte, error) {
	specPropertyName := "spec"
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec, ok := schema.Properties[specPropertyName]
	if !ok {
		// this occurs when a CRD has an empty spec, such as ComputeSharedVPCHostProject
		return make([]byte, 0), nil
	}
	return propsToYAML(spec)
}

func StatusToYAML(crd *apiextensions.CustomResourceDefinition) ([]byte, error) {
	statusPropertyName := "status"
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	status, ok := schema.Properties[statusPropertyName]
	if !ok {
		return nil, fmt.Errorf("unexpected missing '%v' on crd '%v'", statusPropertyName, crd.Spec.Names.Kind)
	}
	return propsToYAML(status)
}

func propsToYAML(props apiextensions.JSONSchemaProps) ([]byte, error) {
	value := propsToValue(props)
	bytes, err := yaml.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("error marshalling value yaml: %w", err)
	}
	return bytes, nil
}

func propsToValue(props apiextensions.JSONSchemaProps) interface{} {
	switch props.Type {
	case "object":
		return objectToValue(props)
	case "array":
		return []interface{}{propsToValue(*props.Items.Schema)}
	case "boolean", "integer", "string":
		return props.Type
	case "number":
		return "float"
	default:
		if props.XPreserveUnknownFields != nil && *props.XPreserveUnknownFields {
			return "schemaless"
		}
		panic(fmt.Sprintf("unhandled type: %v", props.Type))
	}
}

func objectToValue(props apiextensions.JSONSchemaProps) interface{} {
	if isMapType(props) {
		value := make(map[string]string)
		value["string"] = props.AdditionalProperties.Schema.Type
		return value
	}
	value := make(map[string]interface{}, len(props.Properties))
	for k, v := range props.Properties {
		value[k] = propsToValue(v)
	}
	return value
}

func isMapType(props apiextensions.JSONSchemaProps) bool {
	// this property represents a user defined map
	return props.AdditionalProperties != nil && props.AdditionalProperties.Allows
}
