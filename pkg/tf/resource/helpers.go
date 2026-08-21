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

package resource

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func GetTFSchemaForField(tfResource *tfschema.Resource, field string) (*tfschema.Schema, error) {
	return getTFSchemaForNestedField(tfResource, strings.Split(field, ".")...)
}

func getTFSchemaForNestedField(tfResource *tfschema.Resource, fields ...string) (*tfschema.Schema, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no field specified")
	}
	schema, ok := tfResource.Schema[fields[0]]
	if !ok || schema == nil {
		return nil, fmt.Errorf("no schema found for field '%v'", fields[0])
	}
	if len(fields) == 1 {
		return schema, nil
	}
	resource, ok := schema.Elem.(*tfschema.Resource)
	if !ok {
		return nil, fmt.Errorf("expected field '%v' to be a resource, but it was not", fields[0])
	}
	return getTFSchemaForNestedField(resource, fields[1:]...)
}

func TFResourceHasField(tfResource *tfschema.Resource, field string) bool {
	return tfResourceHasNestedField(tfResource, strings.Split(field, ".")...)
}

func tfResourceHasNestedField(tfResource *tfschema.Resource, fields ...string) bool {
	if len(fields) == 0 {
		return false
	}
	schema, ok := tfResource.Schema[fields[0]]
	if !ok || schema == nil {
		return false
	}
	if len(fields) == 1 {
		return true
	}
	resource, ok := schema.Elem.(*tfschema.Resource)
	if !ok {
		return false
	}
	return tfResourceHasNestedField(resource, fields[1:]...)
}

func IsFieldNestedInList(tfResource *tfschema.Resource, field string) bool {
	return isNestedFieldNestedInList(tfResource, strings.Split(field, ".")...)
}

func isNestedFieldNestedInList(tfResource *tfschema.Resource, fields ...string) bool {
	if len(fields) <= 1 {
		return false
	}
	schema, ok := tfResource.Schema[fields[0]]
	if !ok || schema == nil {
		return false
	}
	resource, ok := schema.Elem.(*tfschema.Resource)
	if !ok {
		return false
	}
	if schema.MaxItems != 1 {
		// Field can take more than 1 object => field is a list/set of objects
		return true
	}
	return isNestedFieldNestedInList(resource, fields[1:]...)
}

func TFResourceHasSensitiveFields(tfResource *tfschema.Resource) bool {
	for _, schema := range tfResource.Schema {
		if IsSensitiveConfigurableField(schema) {
			return true
		}
		resource, ok := schema.Elem.(*tfschema.Resource)
		if ok && TFResourceHasSensitiveFields(resource) {
			return true
		}
	}
	return false
}

func IsSensitiveConfigurableField(tfSchema *tfschema.Schema) bool {
	return tfSchema.Sensitive && IsConfigurableField(tfSchema)
}

func IsConfigurableField(tfSchema *tfschema.Schema) bool {
	return tfSchema.Required || tfSchema.Optional
}

func IsObjectField(tfSchema *tfschema.Schema) bool {
	_, ok := tfSchema.Elem.(*tfschema.Resource)
	return tfSchema.Type == tfschema.TypeList && tfSchema.MaxItems == 1 && ok
}

func ImmutableFieldsFromDiff(diff *terraform.InstanceDiff) []string {
	fields := make([]string, 0)
	for field, rd := range diff.Attributes {
		if rd != nil && rd.RequiresNew {
			// TODO(kcc-eng): more deeply KRM-ify the fields coming back (use reference keys, remove ".0." from
			//  MaxItems==1 lists)
			fieldDiff := fmt.Sprintf("Field Name: %s, Got: [redacted], Wanted: [redacted]", text.SnakeCaseToLowerCamelCase(field))
			if !rd.Sensitive {
				fieldDiff = fmt.Sprintf("Field Name: %s, Got: %s, Wanted: %s", text.SnakeCaseToLowerCamelCase(field), rd.New, rd.Old)
			}
			fields = append(fields, fieldDiff)
		}
	}
	return fields
}

func IsSensitiveField(tfSchema *tfschema.Schema) bool {
	return tfSchema.Sensitive
}
