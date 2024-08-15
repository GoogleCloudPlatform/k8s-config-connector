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

package crds

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsvalidation "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/ptr"
)

const (
	FacadeGroup = "facade.compositions.google.com"
)

var scheme = runtime.NewScheme()

func init() {
	if err := apiextensionsv1.AddToScheme(scheme); err != nil {
		panic(err)
	}
	if err := apiextensions.AddToScheme(scheme); err != nil {
		panic(err)
	}
}

type CRDInfo struct {
	Group          string // Should the group be hardcoded to something like facade.compositions.google.com ?
	Kind           string
	Plural         string
	ShortNames     []string
	Categories     []string
	Version        string
	PrinterColumns []apiextensions.CustomResourceColumnDefinition
	Labels         map[string]string
	schema         *apiextensions.JSONSchemaProps
}

func NewFacadeCRDInfo(kind string, plural string,
	shortNames []string, version string,
	printerCols []apiextensions.CustomResourceColumnDefinition,
	labels map[string]string) *CRDInfo {

	crd := CRDInfo{
		Group:          FacadeGroup,
		Kind:           kind,
		Plural:         plural,
		ShortNames:     shortNames,
		Categories:     []string{"facade", "facades"},
		Version:        version,
		PrinterColumns: printerCols,
		schema:         nil,
	}

	crd.Labels = map[string]string{
		"compositions.google.com/facade": "yes",
	}
	for k, v := range labels {
		crd.Labels[k] = v
	}
	return &crd
}

func (c *CRDInfo) SetCRDSchema(schema *apiextensions.JSONSchemaProps) {
	c.schema = schema
}

func (c *CRDInfo) SetSpec(specProperties *apiextensionsv1.JSONSchemaProps) error {
	specUnversionedProperties := &apiextensions.JSONSchemaProps{}
	// Risk ? nil conversion.Scope passed
	if err := apiextensionsv1.Convert_v1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(
		specProperties, specUnversionedProperties, nil); err != nil {
		return err
	}
	statusProperties := map[string]apiextensions.JSONSchemaProps{
		"conditions": {
			Type: "array",
			Items: &apiextensions.JSONSchemaPropsOrArray{
				Schema: &apiextensions.JSONSchemaProps{
					Description: "",
					Required:    []string{"lastTransitionTime", "message", "reason", "status", "type"},
					Type:        "object",
					Properties: map[string]apiextensions.JSONSchemaProps{
						"lastTransitionTime": {
							Description: "",
							Format:      "date-time",
							Type:        "string",
						},
						"message": {
							Description: "human readable message",
							MaxLength:   ptr.To[int64](1024),
							Type:        "string",
						},
						"observedGeneration": {
							Description: "",
							Format:      "int64",
							Minimum:     ptr.To[float64](0),
							Type:        "integer",
						},
						"reason": {
							Description: "",
							MaxLength:   ptr.To[int64](256),
							MinLength:   ptr.To[int64](1),
							Pattern:     "^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$",
							Type:        "string",
						},
						"status": {
							Description: "status of the condition, one of True, False, Unknown.",
							Enum:        []apiextensions.JSON{"True", "False", "Unknown"},
							Type:        "string",
						},
						"type": {
							Description: "type of condition in CamelCase or in foo.example.com/CamelCase." +
								" The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
							MaxLength: ptr.To[int64](316),
							Pattern: "^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)" +
								"?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$",
							Type: "string",
						},
					},
				},
			},
		},
	}

	crdSchema := &apiextensions.JSONSchemaProps{
		Type:        "object",
		Description: "TODO",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"apiVersion": {Type: "string"},
			"kind":       {Type: "string"},
			"metadata":   {Type: "object"},
			"spec":       *specUnversionedProperties.DeepCopy(),
			"status": {
				Type:                   "object",
				Properties:             statusProperties,
				XPreserveUnknownFields: ptr.To[bool](true),
			},
		},
	}

	c.schema = crdSchema
	return nil
}

// CRD takes a schema and converts it to a CRD.
func (c *CRDInfo) CRD() (*apiextensions.CustomResourceDefinition, error) {
	if c.schema == nil {
		return nil, fmt.Errorf("Schema is nil. Use SetSpec or SetCRDSchema first.")
	}
	crd := &apiextensions.CustomResourceDefinition{
		Spec: apiextensions.CustomResourceDefinitionSpec{
			PreserveUnknownFields: ptr.To[bool](false),
			Group:                 c.Group,
			Names: apiextensions.CustomResourceDefinitionNames{
				Kind:       c.Kind,
				ListKind:   c.Kind + "List",
				Plural:     strings.ToLower(c.Plural),
				Singular:   strings.ToLower(c.Kind),
				ShortNames: c.ShortNames,
				Categories: c.Categories,
			},
			Validation: &apiextensions.CustomResourceValidation{
				OpenAPIV3Schema: c.schema,
			},
			Scope:   apiextensions.NamespaceScoped,
			Version: c.Version,
			Subresources: &apiextensions.CustomResourceSubresources{
				Status: &apiextensions.CustomResourceSubresourceStatus{},
				Scale:  nil,
			},
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name:    c.Version,
					Storage: true,
					Served:  true,
				},
			},
			AdditionalPrinterColumns: c.PrinterColumns,
		},
	}

	// Defaulting functions are not found in versionless CRD package
	crdv1 := &apiextensionsv1.CustomResourceDefinition{}
	if err := scheme.Convert(crd, crdv1, nil); err != nil {
		return nil, err
	}
	scheme.Default(crdv1)

	crd2 := &apiextensions.CustomResourceDefinition{}
	if err := scheme.Convert(crdv1, crd2, nil); err != nil {
		return nil, err
	}
	crd2.ObjectMeta.Name = fmt.Sprintf("%s.%s", crd.Spec.Names.Plural, c.Group)

	labels := c.Labels
	if labels == nil {
		labels = make(map[string]string)
	}
	crd2.ObjectMeta.Labels = labels

	return crd2, nil
}

// ValidateCRD calls the CRD package's validation on an internal representation of the CRD.
func ValidateCRD(ctx context.Context, crd *apiextensions.CustomResourceDefinition) error {
	errs := apiextensionsvalidation.ValidateCustomResourceDefinition(ctx, crd)
	if len(errs) > 0 {
		return errs.ToAggregate()
	}
	return nil
}
