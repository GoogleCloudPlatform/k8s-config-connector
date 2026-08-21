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

package crddecoration

import (
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var (
	xPreserveUnknownFields = true
)

func DecorateCRD(crd *apiextensions.CustomResourceDefinition) error {
	// A list of decorate functions that either generically apply to all CRDs or only apply to manually crafted IAM CRDs.
	decorateFuncs := []resourceoverrides.CRDDecorate{
		decorateIAMPolicy,
		decorateIAMPolicyMember,
		decorateIAMPartialPolicy,
		// Mark object fields that have no `Properties` or `AdditionalProperties` with xPreserveUnknownFields
		// since such object fields are intended to allow any kind of sub-field defined by the user or by GCP.
		// Examples include user-definable maps (e.g. DataflowJob spec.parameters) and server-definable map fields in the status.
		allowPropertylessObjectFieldsToPreserveUnknownFields,
		addMissingFieldsForCRDFromTypesFile,
		removeExtraneousFieldsFromCRD,
	}
	for _, f := range decorateFuncs {
		if err := f(crd); err != nil {
			return fmt.Errorf("error decorating CRD %v: %w", crd.GetName(), err)
		}
	}

	// Decorate specific resource CRDs.
	if err := resourceoverrides.Handler.CRDDecorate(crd); err != nil {
		return fmt.Errorf("error decorating CRD %v: %w", crd.GetName(), err)
	}
	return nil
}

func decorateIAMPolicy(crd *apiextensions.CustomResourceDefinition) error {
	kind := crd.Spec.Names.Kind
	if kind != "IAMPolicy" {
		return nil
	}
	// Override the resourceRef schema.
	overrideIAMResourceRefSchema(crd)
	return nil
}

func decorateIAMPolicyMember(crd *apiextensions.CustomResourceDefinition) error {
	kind := crd.Spec.Names.Kind
	if kind != "IAMPolicyMember" {
		return nil
	}
	// Override the resourceRef schema.
	overrideIAMResourceRefSchema(crd)
	// Override the spec schema in IAMPolicyMember to require exactly one
	// of member or memberFrom.
	markMemberAndMemberFromMutuallyExclusiveForIAMPolicyMember(crd)
	// Override the memberFrom schema in IAMPolicyMember to require exactly
	// one subfield in memberFrom.
	markSubfieldsInMemberFromMutuallyExclusiveForIAMPolicyMember(crd)
	return nil
}

func decorateIAMPartialPolicy(crd *apiextensions.CustomResourceDefinition) error {
	kind := crd.Spec.Names.Kind
	if kind != "IAMPartialPolicy" {
		return nil
	}
	// Override the resourceRef schema.
	overrideIAMResourceRefSchema(crd)
	// Override the schema for members item in IAMPartialPolicy spec to require exactly
	// one of member or memberFrom.
	markMemberAndMemberFromMutuallyExclusiveForIAMPartialPolicy(crd)
	// Override the memberFrom schema in IAMPartialPolicy to require exactly
	// one subfield in memberFrom.
	markSubfieldsInMemberFromMutuallyExclusiveForIAMPartialPolicy(crd)
	// Remove regex patterns on status fields ('role' and 'members'); this proactively handles the situation where
	// IAM Policy supports some patterns that KCC falls out-of-sync and controllers fail to update the CR's status
	removeRegexPatternsFromIAMPartialPolicyStatusFields(crd)
	return nil
}

func allowPropertylessObjectFieldsToPreserveUnknownFields(crd *apiextensions.CustomResourceDefinition) error {
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	if spec, ok := schema.Properties["spec"]; ok {
		fixPropertylessObjectFields(&spec)
		schema.Properties["spec"] = spec
	}
	if status, ok := schema.Properties["status"]; ok {
		fixPropertylessObjectFields(&status)
		schema.Properties["status"] = status
	}
	return nil
}

func addMissingFieldsForCRDFromTypesFile(crd *apiextensions.CustomResourceDefinition) error {
	k8s.SetAnnotation(k8s.KCCVersionLabel, "0.0.0-dev", crd)
	if crd.Labels == nil {
		crd.Labels = make(map[string]string)
	}
	crd.Labels[k8s.KCCSystemLabel] = "true"
	crd.Labels[crdgeneration.ManagedByKCCLabel] = "true"

	kind := crd.Spec.Names.Kind
	crd.Spec.Names = apiextensions.CustomResourceDefinitionNames{
		Singular:   strings.ToLower(kind),
		Plural:     crd.Spec.Names.Plural,
		Kind:       kind,
		Categories: []string{crdgeneration.GCPCategory},
		ShortNames: crdgeneration.GenerateShortNames(kind),
	}
	return nil
}

func removeExtraneousFieldsFromCRD(crd *apiextensions.CustomResourceDefinition) error {
	k8s.RemoveAnnotation("controller-gen.kubebuilder.io/version", crd)
	return nil
}

func fixPropertylessObjectFields(jsonSchema *apiextensions.JSONSchemaProps) {
	switch jsonSchema.Type {
	case "object":
		if len(jsonSchema.Properties) == 0 && jsonSchema.AdditionalProperties == nil {
			jsonSchema.XPreserveUnknownFields = &xPreserveUnknownFields
			return
		}
		for k, v := range jsonSchema.Properties {
			fixPropertylessObjectFields(&v)
			jsonSchema.Properties[k] = v
		}
		if jsonSchema.AdditionalProperties != nil {
			fixPropertylessObjectFields(jsonSchema.AdditionalProperties.Schema)
		}
	case "array":
		fixPropertylessObjectFields(jsonSchema.Items.Schema)
	}
}

func markSubfieldsInMemberFromMutuallyExclusiveForIAMPolicyMember(crd *apiextensions.CustomResourceDefinition) {
	kind := crd.Spec.Names.Kind
	if kind != "IAMPolicyMember" {
		return
	}
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec := schema.Properties["spec"]
	memberFrom := spec.Properties["memberFrom"]
	for _, f := range sortedFieldsOf(memberFrom) {
		memberFrom.OneOf = append(memberFrom.OneOf,
			apiextensions.JSONSchemaProps{
				Required: []string{f},
			})
	}
	spec.Properties["memberFrom"] = memberFrom
}

func markMemberAndMemberFromMutuallyExclusiveForIAMPolicyMember(crd *apiextensions.CustomResourceDefinition) {
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec := schema.Properties["spec"]
	spec.OneOf = []apiextensions.JSONSchemaProps{
		{
			Required: []string{"member"},
		},
		{
			Required: []string{"memberFrom"},
		},
	}
	schema.Properties["spec"] = spec
}

func markSubfieldsInMemberFromMutuallyExclusiveForIAMPartialPolicy(crd *apiextensions.CustomResourceDefinition) {
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec := schema.Properties["spec"]
	bindings := spec.Properties["bindings"]
	members := bindings.Items.Schema.Properties["members"]
	memberSchema := members.Items.Schema
	memberFrom := memberSchema.Properties["memberFrom"]
	for _, f := range sortedFieldsOf(memberFrom) {
		memberFrom.OneOf = append(memberFrom.OneOf,
			apiextensions.JSONSchemaProps{
				Required: []string{f},
			})
	}
	memberSchema.Properties["memberFrom"] = memberFrom
}

func markMemberAndMemberFromMutuallyExclusiveForIAMPartialPolicy(crd *apiextensions.CustomResourceDefinition) {
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec := schema.Properties["spec"]
	bindings := spec.Properties["bindings"]
	members := bindings.Items.Schema.Properties["members"]
	memberSchema := members.Items.Schema
	memberSchema.OneOf = []apiextensions.JSONSchemaProps{
		{
			Required: []string{"member"},
		},
		{
			Required: []string{"memberFrom"},
		},
	}
	members.Items.Schema = memberSchema
}

func removeRegexPatternsFromIAMPartialPolicyStatusFields(crd *apiextensions.CustomResourceDefinition) {
	kind := crd.Spec.Names.Kind
	if kind != "IAMPartialPolicy" {
		return
	}
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	status := schema.Properties["status"]
	allBindings := status.Properties["allBindings"]
	removeRegexPatternsFromBindings(&allBindings)
	lastAppliedBindings := status.Properties["lastAppliedBindings"]
	removeRegexPatternsFromBindings(&lastAppliedBindings)
	status.Properties["lastAppliedBindings"] = lastAppliedBindings
	return
}

func removeRegexPatternsFromBindings(bindings *apiextensions.JSONSchemaProps) {
	role := bindings.Items.Schema.Properties["role"]
	role.Pattern = ""
	bindings.Items.Schema.Properties["role"] = role
	members := bindings.Items.Schema.Properties["members"]
	members.Items.Schema.Pattern = ""
	bindings.Items.Schema.Properties["members"] = members
}

func overrideIAMResourceRefSchema(crd *apiextensions.CustomResourceDefinition) {
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec := schema.Properties["spec"]
	spec.Properties["resourceRef"] = iamResourceRefSchema()
}

func iamResourceRefSchema() apiextensions.JSONSchemaProps {
	return apiextensions.JSONSchemaProps{
		Description: "Immutable. Required. The GCP resource to set the IAM policy on.",
		Type:        "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"kind":       {Type: "string"},
			"namespace":  {Type: "string"},
			"name":       {Type: "string"},
			"apiVersion": {Type: "string"},
			"external":   {Type: "string"},
		},
		Required: []string{"kind"},

		// Enforces the following rules:
		// * 'name' and 'external' cannot both be specified
		// * 'namespace' can only be specified if 'name' is specified
		// * only 'kind' is specified for headless IAM
		OneOf: []apiextensions.JSONSchemaProps{
			{
				Required: []string{"name"},
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"external"},
				},
			},
			{
				Required: []string{"external"},
				Not: &apiextensions.JSONSchemaProps{
					AnyOf: []apiextensions.JSONSchemaProps{
						{
							Required: []string{"name"},
						},
						{
							Required: []string{"namespace"},
						},
					},
				},
			},
			{
				// Headless IAM; only kind is specified.
				Not: &apiextensions.JSONSchemaProps{
					AnyOf: []apiextensions.JSONSchemaProps{
						{
							Required: []string{"name"},
						},
						{
							Required: []string{"namespace"},
						},
						{
							Required: []string{"apiVersion"},
						},
						{
							Required: []string{"external"},
						},
					},
				},
			},
		},
	}
}

func sortedFieldsOf(schema apiextensions.JSONSchemaProps) []string {
	fields := make([]string, 0)
	for f := range schema.Properties {
		fields = append(fields, f)
	}
	sort.Strings(fields)
	return fields
}
