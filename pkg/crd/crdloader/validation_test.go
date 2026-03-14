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

package crdloader_test

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestCRDObjectTypes(t *testing.T) {
	crds, err := crdloader.LoadCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	for _, crd := range crds {
		t.Run(crd.Name, func(t *testing.T) {
			for _, version := range crd.Spec.Versions {
				if version.Schema == nil || version.Schema.OpenAPIV3Schema == nil {
					continue
				}
				schema := version.Schema.OpenAPIV3Schema
				for name, subProps := range schema.Properties {
					if name == "metadata" {
						continue
					}
					if err := validateProps(&subProps, fmt.Sprintf("%s.%s", version.Name, name)); err != nil {
						t.Errorf("version %s is invalid: %v", version.Name, err)
					}
				}
			}
		})
	}
}

func validateProps(props *apiextensions.JSONSchemaProps, path string) error {
	if props.Type == "object" {
		if len(props.Properties) == 0 && props.AdditionalProperties == nil && (props.XPreserveUnknownFields == nil || !*props.XPreserveUnknownFields) {
			return fmt.Errorf("object at %s is missing properties, additionalProperties, or x-kubernetes-preserve-unknown-fields", path)
		}
	}
	for name, subProps := range props.Properties {
		if err := validateProps(&subProps, path+"."+name); err != nil {
			return err
		}
	}
	if props.Items != nil {
		if props.Items.Schema != nil {
			if err := validateProps(props.Items.Schema, path+"[]"); err != nil {
				return err
			}
		}
		for i := range props.Items.JSONSchemas {
			if err := validateProps(&props.Items.JSONSchemas[i], fmt.Sprintf("%s[%d]", path, i)); err != nil {
				return err
			}
		}
	}
	if props.AdditionalProperties != nil && props.AdditionalProperties.Schema != nil {
		if err := validateProps(props.AdditionalProperties.Schema, path+"[*]"); err != nil {
			return err
		}
	}
	for i := range props.AllOf {
		if err := validateProps(&props.AllOf[i], fmt.Sprintf("%s.allOf[%d]", path, i)); err != nil {
			return err
		}
	}
	for i := range props.AnyOf {
		if err := validateProps(&props.AnyOf[i], fmt.Sprintf("%s.anyOf[%d]", path, i)); err != nil {
			return err
		}
	}
	for i := range props.OneOf {
		if err := validateProps(&props.OneOf[i], fmt.Sprintf("%s.oneOf[%d]", path, i)); err != nil {
			return err
		}
	}
	if props.Not != nil {
		if err := validateProps(props.Not, path+".not"); err != nil {
			return err
		}
	}
	return nil
}
