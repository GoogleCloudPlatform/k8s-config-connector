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

package diffs

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

type Differ struct {
	skipFields []string
}

var GoogleAPI = Differ{
	skipFields: []string{"ForceSendFields", "NullFields", "ServerResponse"},
}

// Diff compares top-level fields of actual and desired structs (must be pointers to structs)
// using reflection and reflect.DeepEqual, skipping configured skipFields.
func (d Differ) Diff(ctx context.Context, actual, desired any) (*structuredreporting.Diff, []string, error) {
	valActual := reflect.ValueOf(actual)
	valDesired := reflect.ValueOf(desired)

	if valActual.Kind() == reflect.Ptr {
		valActual = valActual.Elem()
	}
	if valDesired.Kind() == reflect.Ptr {
		valDesired = valDesired.Elem()
	}

	if valActual.Kind() != reflect.Struct || valDesired.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("actual and desired must be pointers to structs")
	}

	if valActual.Type() != valDesired.Type() {
		return nil, nil, fmt.Errorf("actual and desired must be of the same struct type")
	}

	diffs := &structuredreporting.Diff{}
	var updatePaths []string

	typ := valActual.Type()
	for i := 0; i < valActual.NumField(); i++ {
		field := typ.Field(i)
		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		// Skip configured skipFields
		if slices.Contains(d.skipFields, field.Name) {
			continue
		}

		// Field name can be determined by JSON/API tags if any, or simply the Go field name.
		fieldName := field.Name
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			if len(parts) > 0 && parts[0] != "" {
				fieldName = parts[0]
			}
		}

		actualFieldVal := valActual.Field(i).Interface()
		desiredFieldVal := valDesired.Field(i).Interface()

		if !reflect.DeepEqual(actualFieldVal, desiredFieldVal) {
			diffs.AddField(fieldName, actualFieldVal, desiredFieldVal)
			updatePaths = append(updatePaths, fieldName)
		}
	}

	slices.Sort(updatePaths)

	return diffs, updatePaths, nil
}
