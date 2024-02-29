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

package mappings

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"

	"k8s.io/klog/v2"
)

// ValidationError is the type for one of our validation findings.
type ValidationError struct {
	Message  string
	Proposal string
}

// Validate will perform some checks on the mapping.
// We always look for invalid mappings, where (for example) the field does not exist on one or both sides.
// If CHECK_COVERAGE is passed, it will look for missing mappings.
func (m *Mapping) Validate() []ValidationError {
	var errors []ValidationError
	for _, typeMapping := range m.Mappings {
		checkMissing := (os.Getenv("CHECK_COVERAGE") != "")
		switch typeMapping := typeMapping.(type) {
		case *structTypeMapping:
			errors = append(errors, typeMapping.validate(checkMissing)...)
		default:
			klog.Fatalf("unhandled type mapping %T", typeMapping)
		}
	}
	return errors
}

// validate performs validation on a structTypeMapping
func (m *structTypeMapping) validate(checkMissing bool) []ValidationError {
	var errors []ValidationError

	for _, mapping := range m.fields {
		inField := m.inType.lookupField(mapping.InPath)
		outField := m.outType.lookupField(mapping.OutPath)

		if inField == nil {
			err := ValidationError{Message: fmt.Sprintf("field %s not found in input type %v", mapping.InPath, m.inType)}
			if outField != nil {
				err.Proposal = buildGoField(outField)
			}

			errors = append(errors, err)
			continue
		}
		if outField == nil {
			proposal := buildGoField(inField)

			klog.Infof("outType.Fields = %v", m.outType.Fields())

			errors = append(errors, ValidationError{
				Message:  fmt.Sprintf("field %s not found in output type %v", mapping.OutPath, m.outType),
				Proposal: proposal,
			})
			continue
		}

	}

	if checkMissing {

		if m.hasSpecStatus {
			specField := m.inType.lookupField(parseFieldPath("spec"))
			statusField := m.inType.lookupField(parseFieldPath("status"))
			if specField != nil && statusField != nil {
				specType := specField.Type()
				statusType := statusField.Type()

				for _, cloudField := range m.outType.Fields() {
					id := cloudField.ID()

					ignore := false
					for _, ignoreField := range m.ignore {
						if ignoreField.ID == id {
							ignore = true
						}
					}
					if ignore {
						continue
					}
					specField := specType.lookupField(newFieldPath(id))
					statusField := statusType.lookupField(newFieldPath(id))

					if specField == nil && statusField == nil {
						proposal := buildGoField(cloudField)

						errors = append(errors, ValidationError{
							Message:  fmt.Sprintf("field %s not found in KRM spec nor status %v", id, m.inType),
							Proposal: proposal,
						})
						continue
					}
				}
			}
		} else {
			for _, outField := range m.outType.Fields() {
				id := outField.ID()

				found := false
				for _, mapping := range m.fields {
					if len(mapping.OutPath.fields) == 1 && mapping.OutPath.fields[0] == id {
						found = true
					}
				}
				if found {
					continue
				}

				ignore := false
				for _, ignoreField := range m.ignore {
					if ignoreField.ID == id {
						ignore = true
					}
				}
				if ignore {
					continue
				}
				errors = append(errors, ValidationError{
					Message: fmt.Sprintf("field %s is not mapped in %v", id, m.outType),
					// Proposal: proposal,
				})
			}

			for _, inField := range m.inType.Fields() {
				id := inField.ID()

				found := false
				for _, mapping := range m.fields {
					if len(mapping.InPath.fields) == 1 && mapping.InPath.fields[0] == id {
						found = true
					}
				}
				if found {
					continue
				}

				ignore := false
				for _, ignoreField := range m.ignore {
					if ignoreField.ID == id {
						ignore = true
					}
				}
				if ignore {
					continue
				}
				errors = append(errors, ValidationError{
					Message: fmt.Sprintf("field %s is not mapped in %v", id, m.inType),
				})
			}
		}

	}

	return errors
}

// buildGoField builds go code that could be used to add the relevant field.
func buildGoField(f Field) string {
	jsonName := f.getJSONKey()

	fieldName := jsonToGoFieldName(jsonName)
	fieldType := convertToGoType(f.Type().rt)
	jsonTag := jsonName
	jsonTag += ",omitempty"

	requiredTag := ""
	if f.isRequired() {
		requiredTag = "true"
	}

	tags := []string{}
	if jsonTag != "" {
		tags = append(tags, fmt.Sprintf("json:%q", jsonTag))
	}
	if requiredTag != "" {
		tags = append(tags, fmt.Sprintf("required:%q", requiredTag))
	}

	fieldTags := ""
	if len(tags) != 0 {
		fieldTags = " `" + strings.Join(tags, " ") + "`"
	}

	proposal := fmt.Sprintf("%s %s%s", fieldName, fieldType, fieldTags)
	return proposal
}

// jsonToGoFieldName converts a json name into its go field name equivalent,
// used to suggest go code.
func jsonToGoFieldName(jsonName string) string {
	var out []rune
	for i, r := range jsonName {
		if i == 0 {
			r = unicode.ToUpper(r)
		}
		out = append(out, r)
	}
	return string(out)
}

// convertToGoType builds the go type name for the specified type,
// used to suggest go code.
func convertToGoType(t reflect.Type) string {
	fieldGoType := t
	switch fieldGoType.Kind() {
	case reflect.Slice:
		return "[]" + convertToGoType(t.Elem())
	case reflect.Ptr:
		return "*" + convertToGoType(t.Elem())
	case reflect.Struct:
		return t.Name()
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Uint8:
		return "uint8"
	case reflect.Int32:
		return "int32"
	case reflect.Int64:
		return "int64"
	case reflect.Map:
		return "map[todo]todo"
	case reflect.Interface:
		return "interface{}"
	default:
		klog.Fatalf("unsupported kind in convertToGoType %v", fieldGoType.Kind())
		return ""
	}
}
