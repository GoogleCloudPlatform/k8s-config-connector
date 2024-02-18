// Copyright 2023 Google LLC
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

package k8s

import (
	"fmt"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type SchemaReference struct {
	CRD        *apiextensions.CustomResourceDefinition
	JSONSchema *apiextensions.JSONSchemaProps
	GVK        schema.GroupVersionKind
}

type SchemaReferenceUpdater interface {
	// UpdateSchema updates the schema reference of the controller using
	// its corresponding CRD.
	UpdateSchema(crd *apiextensions.CustomResourceDefinition) error
}

// UpdateSchema is a helper function to update the input schema reference using the input crd.
func UpdateSchema(schemaRef *SchemaReference, crd *apiextensions.CustomResourceDefinition) error {
	gvk := schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: GetVersionFromCRD(crd),
		Kind:    crd.Spec.Names.Kind,
	}
	if schemaRef.GVK.String() != gvk.String() {
		return fmt.Errorf("unexpected mismatch of GVK when updating schema reference for controller, old GVK = %s, new GVK = %s", schemaRef.GVK.String(), gvk.String())
	}
	schemaRef.CRD = crd
	schemaRef.JSONSchema = GetOpenAPIV3SchemaFromCRD(crd)
	return nil
}
