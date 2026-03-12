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

package fielddesc

import (
	"fmt"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type RequirementLevel string

const (
	OptionalRequirementLevel                  = "Optional"
	RequiredWhenParentPresentRequirementLevel = "RequiredWhenParentPresent"
	RequiredRequirementLevel                  = "Required"
)

type FieldDescription struct {
	FullName             []string
	ShortName            string
	Description          string
	Type                 string
	Format               string
	RequirementLevel     RequirementLevel
	Children             []FieldDescription
	AdditionalProperties []FieldDescription
}

func GetSpecDescription(crd *apiextensions.CustomResourceDefinition, version string) FieldDescription {
	crdDesc := getCRDFieldDescription(crd, version)
	spec, ok := getChildFieldDesc(crdDesc, "spec")
	if !ok {
		// this occurs when a CRD has an empty spec, such as ComputeSharedVPCHostProject
		return FieldDescription{
			Type:             "object",
			RequirementLevel: OptionalRequirementLevel,
			Children:         make([]FieldDescription, 0),
		}
	}
	return *spec
}

func GetStatusDescription(crd *apiextensions.CustomResourceDefinition, version string) (FieldDescription, error) {
	statusPropertyName := "status"
	crdDesc := getCRDFieldDescription(crd, version)
	status, ok := getChildFieldDesc(crdDesc, statusPropertyName)
	if !ok {
		return FieldDescription{}, fmt.Errorf("unexpected missing '%v' on crd '%v'", statusPropertyName, crd.Spec.Names.Kind)
	}
	return *status, nil
}

func getChildFieldDesc(description FieldDescription, childName string) (*FieldDescription, bool) {
	for _, c := range description.Children {
		if c.ShortName == childName {
			return &c, true
		}
	}
	return nil, false
}

func getCRDFieldDescription(crd *apiextensions.CustomResourceDefinition, version string) FieldDescription {
	customResourceDesc := FieldDescription{
		Type:             "object",
		RequirementLevel: RequiredRequirementLevel,
	}
	crdVersionDefinition := k8s.GetCRDVersionDefinition(crd, version)
	return propsToDescription(*crdVersionDefinition.Schema.OpenAPIV3Schema, customResourceDesc, "", true)
}

func propsToDescription(props apiextensions.JSONSchemaProps, parent FieldDescription, name string, required bool) FieldDescription {
	switch props.Type {
	case "object":
		return objectToDescription(props, parent, name, required)
	case "array":
		return sliceToDescriptions(props, parent, name, required)
	case "boolean", "integer", "string", "number":
		return newFieldDescription(props, parent, name, required)
	default:
		if props.XPreserveUnknownFields != nil && *props.XPreserveUnknownFields {
			return schemalessToDescription(props, parent, name, required)
		}
		panic(fmt.Sprintf("unhandled type: %v", props.Type))
	}
}

func schemalessToDescription(props apiextensions.JSONSchemaProps, parent FieldDescription, name string, required bool) FieldDescription {
	fd := newFieldDescription(props, parent, name, required)
	fd.Type = "schemaless"
	return fd
}

func sliceToDescriptions(props apiextensions.JSONSchemaProps, parent FieldDescription, name string, required bool) FieldDescription {
	propsItemSchema := *props.Items.Schema
	fd := newFieldDescription(props, parent, name, required)
	fd.Type = fmt.Sprintf("list (%v)", propsItemSchema.Type)
	fd.Children = []FieldDescription{propsToDescription(propsItemSchema, fd, "[]", required)}
	return fd
}

func objectToDescription(props apiextensions.JSONSchemaProps, parent FieldDescription, name string, required bool) FieldDescription {
	fd := newFieldDescription(props, parent, name, required)
	isMap := isMapType(props)
	if isMap {
		supportedTypes := map[string]bool{
			"boolean": true,
			"integer": true,
			"string":  true,
			"number":  true,
			"object":  true,
		}
		valueType := props.AdditionalProperties.Schema.Type
		if _, ok := supportedTypes[valueType]; ok {
			fd.Type = fmt.Sprintf("map (key: string, value: %v)", valueType)
		} else if props.AdditionalProperties.Schema.XPreserveUnknownFields != nil && *props.AdditionalProperties.Schema.XPreserveUnknownFields {
			fd.Type = fmt.Sprintf("map (key: string, value: %v)", "any")
		} else {
			panic(fmt.Sprintf("objectToDescription: unsupported map type in %+v, only support maps of boolean, integer, string, number, and object types", props.AdditionalProperties.Schema))
		}

		fd.Type = fmt.Sprintf("map (key: string, value: %v)", valueType)
		if valueType != "object" {
			return fd
		}

		props = *props.AdditionalProperties.Schema
	}
	requiredFields := make(map[string]bool)
	for _, s := range props.Required {
		requiredFields[s] = true
	}
	keys := make([]string, 0, len(props.Properties))
	for k := range props.Properties {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := props.Properties[k]
		if isMap {
			fd.AdditionalProperties = append(fd.AdditionalProperties, propsToDescription(v, fd, k, requiredFields[k]))
		} else {
			fd.Children = append(fd.Children, propsToDescription(v, fd, k, requiredFields[k]))
		}
	}
	return fd
}

func newFieldDescription(props apiextensions.JSONSchemaProps, parent FieldDescription, name string, required bool) FieldDescription {
	fullName := make([]string, len(parent.FullName), len(parent.FullName)+1)
	copy(fullName, parent.FullName)
	if name != "" {
		fullName = append(fullName, name)
	}
	fd := FieldDescription{
		Type:        props.Type,
		Format:      props.Format,
		Description: props.Description,
		FullName:    fullName,
		ShortName:   name,
	}
	if fd.Type == "number" {
		fd.Type = "float"
	}
	if required {
		switch parent.RequirementLevel {
		case RequiredRequirementLevel:
			fd.RequirementLevel = RequiredRequirementLevel
		case RequiredWhenParentPresentRequirementLevel, OptionalRequirementLevel:
			fd.RequirementLevel = RequiredWhenParentPresentRequirementLevel
		default:
			panic(fmt.Errorf("unhandled requirement level: %v", parent.RequirementLevel))
		}
	} else {
		fd.RequirementLevel = OptionalRequirementLevel
	}
	return fd
}

func isMapType(props apiextensions.JSONSchemaProps) bool {
	// this property represents a user defined map
	return props.AdditionalProperties != nil && props.AdditionalProperties.Allows
}
