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

package dcl

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/typeutil"

	"github.com/nasa9084/go-openapi"
)

func MutableButUnreadableFieldsAnnotationFor(r *Resource) (string, error) {
	paths, err := getMutableButUnreadablePaths(r)
	if err != nil {
		return "", fmt.Errorf("error getting mutable-but-unreadable fields for resource: %w", err)
	}
	return k8s.GenerateMutableButUnreadableFieldsAnnotation(&r.Resource, paths)
}

// getMutableButUnreadablePaths returns the list of fields supported by the
// resource that are mutable but unreadable. Each field is broken down into its
// path elements (i.e. ["spec", "fooBar"] instead of "spec.fooBar").
func getMutableButUnreadablePaths(r *Resource) ([][]string, error) {
	paths, err := getMutableButUnreadableDCLPathsInObject([]string{}, r.Schema)
	if err != nil {
		return nil, fmt.Errorf("error getting mutable-but-unreadable fields from the DCL resource schema: %w", err)
	}
	return paths, nil
}

func GetMutableButUnreadableFieldsFromAnnotations(r *Resource) (map[string]interface{}, error) {
	paths, err := getMutableButUnreadablePaths(r)
	if err != nil {
		return nil, err
	}
	return k8s.GetMutableButUnreadableFieldsFromAnnotations(&r.Resource, paths)
}

func getMutableButUnreadableDCLPathsInObject(path []string, schema *openapi.Schema) ([][]string, error) {
	if schema.Type != "object" {
		return nil, fmt.Errorf("expect the schema type to be 'object', but got %v", schema.Type)
	}

	results := make([][]string, 0)
	for subField, subSchema := range schema.Properties {
		subPath := append(path, subField)

		// ReadOnly fields should be skipped because they and their subfields
		// can't be mutable but unreadable.
		if subSchema.ReadOnly {
			continue
		}

		isMutableButUnreadable, err := extension.IsMutableButUnreadableField(subSchema)
		if err != nil {
			return nil, err
		}
		if isMutableButUnreadable {
			results = append(results, subPath)
			continue
		}

		switch subSchema.Type {
		// We haven't determined whether 'x-dcl-mutable-unreadable' extension should
		// be set at the field-level, or within the item schema for an array. More
		// details here: http://b/186078207.
		// Currently KCC is also checking the item schema within an array.
		case "array":
			// Mutable-but-unreadable feature is not supported for array fields
			// with non-primitive types.
			if !typeutil.IsPrimitiveTypeArray(subSchema.Items) {
				continue
			}

			isMutableButUnreadable, err := extension.IsMutableButUnreadableField(subSchema.Items)
			if err != nil {
				return nil, err
			}
			if isMutableButUnreadable {
				results = append(results, subPath)
			}
		// Object field is not a mutable-but-unreadable field, but might contain fields that are.
		case "object":
			subResults, err := getMutableButUnreadableDCLPathsInObject(subPath, subSchema)
			if err != nil {
				return nil, fmt.Errorf("error getting mutable-but-unreadable fields under sub field %s: %w", subField, err)
			}
			results = append(results, subResults...)
		}
	}
	return results, nil
}
