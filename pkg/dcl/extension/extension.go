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

// Package extension is used to interpret the dcl extensions.
package extension

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/constants"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"

	"github.com/nasa9084/go-openapi"
)

var (
	trimmableReferenceSuffixes = []string{"Name", "Id", "IdOrNum", "Email", "Link", "Reference"}
)

// IsReferenceField takes the field schema and determines if the field is for resource reference.
func IsReferenceField(schema *openapi.Schema) bool {
	refSchema := schema
	if schema.Type == "array" {
		refSchema = schema.Items
	}
	_, ok := refSchema.Extension["x-dcl-references"]
	return ok
}

// IsSensitiveField takes the field schema and determines if the field is sensitive.
func IsSensitiveField(schema *openapi.Schema) (bool, error) {
	val, ok := schema.Extension["x-dcl-sensitive"]
	if !ok {
		return false, nil
	}
	if schema.Type != "string" {
		return false, fmt.Errorf("only support sensitive fields of `string` type, but got type %v", schema.Type)
	}
	boolVal, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("wrong type for 'x-dcl-sensitive' extension: %T, expect to have bool", val)
	}
	return boolVal, nil
}

func HasSensitiveFields(schema *openapi.Schema) (bool, error) {
	sensitive, err := IsSensitiveField(schema)
	if err != nil {
		return false, err
	}
	if sensitive {
		return true, nil
	}

	switch schema.Type {
	case "array":
		return HasSensitiveFields(schema.Items)
	case "object":
		if schema.AdditionalProperties != nil {
			return HasSensitiveFields(schema.AdditionalProperties)
		}
		for _, fieldSchema := range schema.Properties {
			sensitive, err = HasSensitiveFields(fieldSchema)
			if err != nil {
				return false, err
			}
			if sensitive {
				return true, nil
			}
		}
	}

	return false, nil
}

func HasIam(schema *openapi.Schema) (bool, error) {
	val, ok := schema.Extension["x-dcl-has-iam"]
	if !ok { // extension doesn't exist
		return false, nil
	}
	boolVal, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("wrong type for 'x-dcl-has-iam' extension: %T, expected to have bool", val)
	}
	return boolVal, nil
}

func IsImmutableField(schema *openapi.Schema) (bool, error) {
	val, ok := schema.Extension["x-kubernetes-immutable"]
	if !ok {
		return false, nil
	}
	boolVal, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("wrong type for 'x-kubernetes-immutable' extension: %T, expect to have bool", val)
	}
	return boolVal, nil
}

func GetLabelsFieldSchema(schema *openapi.Schema) (labelsField string, fieldSchema *openapi.Schema, found bool, err error) {
	raw, found := schema.Extension[constants.DCLLabelsField]
	if !found {
		return "", nil, false, nil
	}
	labelsField, ok := raw.(string)
	if !ok {
		return "", nil, false, fmt.Errorf("wrong type for 'x-dcl-labels' extension: %T, expect to have string type", raw)
	}
	if labelsField == "" {
		return "", nil, false, fmt.Errorf("'x-dcl-labels' field exists, but is an empty string")
	}
	path := strings.Split(labelsField, ".")
	var ret *openapi.Schema
	ret = schema
	for _, field := range path {
		if _, ok := ret.Properties[field]; !ok {
			return "", nil, false, fmt.Errorf("couldn't find the schema for field %v", labelsField)
		}
		ret = ret.Properties[field]
	}
	return labelsField, ret, true, nil
}

// GetReferenceFieldName returns the converted field name given the original reference field name.
func GetReferenceFieldName(path []string, schema *openapi.Schema) (string, error) {
	field := pathslice.Base(path)
	if len(path) == 1 && field == "parent" {
		return "", fmt.Errorf("cannot get reference field name for 'parent' " +
			"since 'parent' is typically split into multiple reference fields")
	}
	// If the filed is `array` type, it expects an list of references. We keep the original field name.
	if schema.Type == "array" {
		return field, nil
	}
	raw := schema.Extension["x-dcl-references"]
	_, ok := raw.([]interface{})
	if !ok {
		return "", fmt.Errorf("wrong type for 'x-dcl-references' extension: %T, expect to have []interface{}", raw)
	}
	return formatReferenceFieldName(field), nil
}

func formatReferenceFieldName(fieldName string) string {
	// If the original field name ends with one of the known suffixes, e.g. "xxxName",
	// we want to convert it to "xxxRef" for consistency
	for _, suffix := range trimmableReferenceSuffixes {
		if strings.HasSuffix(fieldName, suffix) {
			return strings.TrimSuffix(fieldName, suffix) + "Ref"
		}
	}
	return fieldName + "Ref"
}

func GetNameFieldSchema(schema *openapi.Schema) (*openapi.Schema, bool) {
	s, ok := schema.Properties["name"]
	if !ok {
		return nil, false
	}
	return s, true
}

func IsResourceIDFieldServerGenerated(nameFieldSchema *openapi.Schema) (bool, error) {
	val, ok := nameFieldSchema.Extension["x-dcl-server-generated-parameter"]
	if !ok {
		return false, nil
	}
	boolVal, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("wrong type for 'x-dcl-server-generated-parameter' extension: %T, expect to have bool", val)
	}
	return boolVal, nil
}

func GetNameValueTemplate(schema *openapi.Schema) (string, error) {
	raw, ok := schema.Extension["x-dcl-id"]
	if !ok {
		return "", fmt.Errorf("'x-dcl-id' is not found")
	}
	template, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("wrong type for 'x-dcl-id' extension: %T, expect to have string type", raw)
	}
	return template, nil
}

func HasStateHint(schema *openapi.Schema) (bool, error) {
	val, ok := schema.Extension["x-dcl-uses-state-hint"]
	if !ok {
		return false, nil
	}

	boolVal, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("wrong type for 'x-dcl-uses-state-hint' extension: %T, expect to have bool type", val)
	}

	return boolVal, nil
}

func IsMutableButUnreadableField(schema *openapi.Schema) (bool, error) {
	// Check if the field is unreadable.
	val, ok := schema.Extension["x-dcl-mutable-unreadable"]
	if !ok {
		return false, nil
	}

	unreadable, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("wrong type for 'x-dcl-mutable-unreadable' extension: %T, expect to have bool type", val)
	}

	if !unreadable {
		return false, nil
	}

	// Check if the field is mutable.
	immutable, err := IsImmutableField(schema)
	if err != nil {
		return false, err
	}
	return !immutable, nil
}
