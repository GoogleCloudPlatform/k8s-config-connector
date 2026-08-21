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

package resourceoverrides

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/crdutil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	instanceTypeField    = []string{"instanceType"}
	instanceTypeRefField = []string{"instanceTypeRef"}
)

func GetAlloyDBInstanceResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "AlloyDBInstance",
	}
	// Preserve the legacy non-reference field 'instanceType' after it is changed to
	// a reference field, 'instanceTypeRef'.
	ro.Overrides = append(ro.Overrides, keepInstanceTypeField())
	return ro
}

func keepInstanceTypeField() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return preserveBothInstanceTypeAndInstanceTypeRefInCRD(crd, instanceTypeRefField[0], instanceTypeField[0])
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		if err := PreserveUserSpecifiedLegacyField(original, reconciled, instanceTypeField...); err != nil {
			return fmt.Errorf("error preserving '%v' in post-actuation transformation: %w", strings.Join(instanceTypeField, "."), err)
		}
		if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified(original, reconciled, instanceTypeField, instanceTypeRefField); err != nil {
			return fmt.Errorf("error conditionally pruning '%v' in post-actuation transformation: %w", strings.Join(instanceTypeRefField, "."), err)
		}
		return nil
	}
	return o
}

func preserveBothInstanceTypeAndInstanceTypeRefInCRD(crd *apiextensions.CustomResourceDefinition, referenceFieldName, nonReferenceFieldName string) error {

	var err error

	v1beta1Schema := crd.Spec.Versions[0].Schema.OpenAPIV3Schema
	err = preserveBothInstanceTypeAndInstanceTypeRefInCRDForSchema(crd, referenceFieldName, nonReferenceFieldName, v1beta1Schema)
	if err != nil {
		return err
	}

	v1alpha1Schema := crd.Spec.Versions[1].Schema.OpenAPIV3Schema
	err = preserveBothInstanceTypeAndInstanceTypeRefInCRDForSchema(crd, referenceFieldName, nonReferenceFieldName, v1alpha1Schema)
	if err != nil {
		return err
	}

	return nil
}

// Allows side-by-side mutually exclusive non-ref field instanceType and new ref field instanceTypeRef
// Based on resourceoverrides.PreserveMutuallyExclusiveNonReferenceField
func preserveBothInstanceTypeAndInstanceTypeRefInCRDForSchema(crd *apiextensions.CustomResourceDefinition, referenceFieldName, nonReferenceFieldName string, schema *apiextensions.JSONSchemaProps) error {

	var err error
	var parent *apiextensions.JSONSchemaProps

	parentPath := []string{"spec"}

	parentPathStr := strings.Join(parentPath, ".")
	parent, err = getSchemaForPath(schema, parentPath)
	if err != nil {
		return fmt.Errorf("can't get schema for path '%v' in CRD %s: %w", parentPathStr, crd.Name, err)
	}

	oneOfRule, err := crdutil.GetOneOfRuleForObjectOrArray(parent)
	if err != nil {
		return fmt.Errorf("error getting the oneOf rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
	}

	oneOfRule = []*apiextensions.JSONSchemaProps{
		{
			Required: []string{nonReferenceFieldName},
		},
		{
			Required: []string{referenceFieldName},
		},
	}

	if err := crdutil.SetOneOfRuleForObjectOrArray(parent, oneOfRule); err != nil {
		return fmt.Errorf("error setting the oneOf rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
	}

	requiredFields, err := crdutil.GetRequiredRuleForObjectOrArray(parent)
	if err != nil {
		return fmt.Errorf("error getting the required rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
	}

	var updatedRequiredFields []string
	for _, field := range requiredFields {
		if field != referenceFieldName {
			updatedRequiredFields = append(updatedRequiredFields, field)
		}
	}
	if err := crdutil.SetRequiredRuleForObjectOrArray(parent, updatedRequiredFields); err != nil {
		return fmt.Errorf("error setting the required rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
	}

	referenceFieldSchema, ok, err := crdutil.GetSchemaForFieldUnderObjectOrArray(referenceFieldName, parent)
	if err != nil {
		return fmt.Errorf("error getting schema for reference field %s under path %s for CRD %s: %w", referenceFieldName, parentPath, crd.Name, err)
	}
	if !ok {
		return fmt.Errorf("can't find reference field %s under path %s for CRD %s", referenceFieldName, parentPath, crd.Name)
	}
	fieldType := referenceFieldSchema.Type
	if fieldType != "object" && fieldType != "array" {
		return fmt.Errorf("wrong type for reference field %s under path %s for CRD %s: %s", referenceFieldName, parentPath, crd.Name, fieldType)
	}

	var nonReferenceFieldSchema *apiextensions.JSONSchemaProps
	description := fmt.Sprintf("We recommend that you use `%s` instead.\nThe type of the instance. Possible values: [PRIMARY, READ_POOL, SECONDARY]", referenceFieldName)

	nonReferenceFieldSchema = &apiextensions.JSONSchemaProps{
		Description: description,
		Type:        "string",
	}

	if err := crdutil.SetSchemaForFieldUnderObjectOrArray(nonReferenceFieldName, parent, nonReferenceFieldSchema); err != nil {
		return fmt.Errorf("error setting schema for non-reference field %s under path %s for CRD %s: %w", nonReferenceFieldName, parentPath, crd.Name, err)
	}

	updatedSchema, err := getSchemaForPath(schema, parentPath[:len(parentPath)-1])
	if err != nil {
		return fmt.Errorf("can't get schema for path '%v' in CRD %s: %w", parentPathStr, crd.Name, err)
	}
	if err := crdutil.SetSchemaForFieldUnderObjectOrArray(parentPath[len(parentPath)-1], updatedSchema, parent); err != nil {
		return fmt.Errorf("error setting updated schema for parent path '%v' for CRD %s: %w", parentPathStr, crd.Name, err)
	}

	return nil
}
