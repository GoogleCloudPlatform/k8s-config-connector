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

package crdutil

import (
	"fmt"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var (
	errObjectSchema = fmt.Errorf("the object schema has neither properties nor additionalProperties in the schema")
	errSchemaType   = fmt.Errorf("type of schema must be 'object' or 'array'")
)

func getConfigurableSchemaForObjectOrArray(schema *apiextensions.JSONSchemaProps) (configurableSchema *apiextensions.JSONSchemaProps, err error) {
	switch schema.Type {
	case "object":
		if schema.Properties != nil {
			configurableSchema = schema
		} else if schema.AdditionalProperties != nil {
			configurableSchema = schema.AdditionalProperties.Schema
		} else {
			return nil, errObjectSchema
		}
	case "array":
		return getConfigurableSchemaForObjectOrArray(schema.Items.Schema)
	default:
		return nil, errSchemaType
	}
	return configurableSchema, nil
}

func GetSchemaForFieldUnderObjectOrArray(fieldName string, parent *apiextensions.JSONSchemaProps) (schema *apiextensions.JSONSchemaProps, ok bool, err error) {
	var schemaValue apiextensions.JSONSchemaProps
	configurableSchema, err := getConfigurableSchemaForObjectOrArray(parent)
	if err != nil {
		return nil, false, err
	}
	schemaValue, ok = configurableSchema.Properties[fieldName]
	if ok {
		schema = &schemaValue
	}
	return schema, ok, nil
}

func SetSchemaForFieldUnderObjectOrArray(fieldName string, parent *apiextensions.JSONSchemaProps, schema *apiextensions.JSONSchemaProps) error {
	configurableSchema, err := getConfigurableSchemaForObjectOrArray(parent)
	if err != nil {
		return err
	}
	configurableSchema.Properties[fieldName] = *schema
	return nil
}

func GetRequiredRuleForObjectOrArray(schema *apiextensions.JSONSchemaProps) ([]string, error) {
	configurableSchema, err := getConfigurableSchemaForObjectOrArray(schema)
	if err != nil {
		return nil, err
	}
	return configurableSchema.Required, nil
}

func SetRequiredRuleForObjectOrArray(schema *apiextensions.JSONSchemaProps, requiredRule []string) error {
	configurableSchema, err := getConfigurableSchemaForObjectOrArray(schema)
	if err != nil {
		return err
	}
	configurableSchema.Required = requiredRule
	return nil
}

func GetOneOfRuleForObjectOrArray(schema *apiextensions.JSONSchemaProps) ([]*apiextensions.JSONSchemaProps, error) {
	var oneOfRule []*apiextensions.JSONSchemaProps
	var oneOfRuleValue []apiextensions.JSONSchemaProps
	configurableSchema, err := getConfigurableSchemaForObjectOrArray(schema)
	if err != nil {
		return nil, err
	}
	oneOfRuleValue = configurableSchema.OneOf
	if oneOfRuleValue != nil {
		for i := range oneOfRuleValue {
			oneOfRule = append(oneOfRule, &oneOfRuleValue[i])
		}
	}
	return oneOfRule, nil
}

func SetOneOfRuleForObjectOrArray(schema *apiextensions.JSONSchemaProps, oneOfRule []*apiextensions.JSONSchemaProps) error {
	var oneOfRuleValue []apiextensions.JSONSchemaProps
	for _, rule := range oneOfRule {
		oneOfRuleValue = append(oneOfRuleValue, *rule)
	}
	configurableSchema, err := getConfigurableSchemaForObjectOrArray(schema)
	if err != nil {
		return err
	}
	configurableSchema.OneOf = oneOfRuleValue
	return nil
}

func GetNotRuleForObjectOrArray(schema *apiextensions.JSONSchemaProps) (*apiextensions.JSONSchemaProps, error) {
	configurableSchema, err := getConfigurableSchemaForObjectOrArray(schema)
	if err != nil {
		return nil, err
	}
	return configurableSchema.Not, nil
}

func SetNotRuleForObjectOrArray(schema *apiextensions.JSONSchemaProps, notRule *apiextensions.JSONSchemaProps) error {
	configurableSchema, err := getConfigurableSchemaForObjectOrArray(schema)
	if err != nil {
		return err
	}
	configurableSchema.Not = notRule
	return nil
}
