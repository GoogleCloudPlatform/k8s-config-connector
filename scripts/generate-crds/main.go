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

// This file generates CRDs for the type providers defined in the input type providers directory.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crddecoration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	"github.com/ghodss/yaml" //nolint:depguard
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crdgen "sigs.k8s.io/controller-tools/pkg/crd"
	"sigs.k8s.io/controller-tools/pkg/genall"
)

const (
	dirMode  = 0700
	fileMode = 0600
)

var (
	outputDir = ""
)

func main() {
	flag.StringVar(&outputDir, "output-dir", "", "Directory where CRD files are to be written to")
	flag.Parse()
	if outputDir == "" {
		fmt.Printf("error: invalid value for output directory: '%s'\n", "empty string")
		flag.PrintDefaults()
		os.Exit(1)
	}
	// create the output dir if it doesn't exist
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, dirMode); err != nil {
			fmt.Printf("error creating output directory '%v': %v\n", outputDir, err)
			os.Exit(1)
		}
	}

	crds := make([]*apiextensions.CustomResourceDefinition, 0)
	// generate TF2CRD-based resources
	crds = append(crds, generateTFBasedCRDs()...)
	// generate DCL2CRD-based resources
	crds = append(crds, generateDCLBasedCRDs()...)
	// generate CRDs for the strongly-typed Go structs in pkg/apis
	// (e.g. IAMPolicy, IAMPolicyMember, IAMAuditConfig, ServiceMapping)
	crds = append(crds, generateCRDsForTypesFiles()...)

	for _, crd := range crds {
		if err := crddecoration.DecorateCRD(crd); err != nil {
			log.Fatalf("error decorating CRD %v: %v", crd.GetName(), err)
		}
		if err := outputCRDToFile(crd); err != nil {
			log.Fatalf("error outputting CRD %v to file: %v", crd.GetName(), err)
		}
	}
	outputPath, _ := filepath.Abs(outputDir)
	fmt.Printf("CRD manifests generated under '%v'\n", outputPath)
}

func generateTFBasedCRDs() []*apiextensions.CustomResourceDefinition {
	outputCRDs := make([]*apiextensions.CustomResourceDefinition, 0)
	serviceMappings, err := servicemappingloader.GetServiceMappings()
	if err != nil {
		log.Fatalf("could not load service mappings: %v", err)
	}
	for _, sm := range serviceMappings {
		kindResourceConfigsMap := generateKindResourceConfigsMap(&sm)
		for kind, rcs := range kindResourceConfigsMap {
			switch kind {
			case "ComputeInstance":
				// We merge the two TF resources: compute_instance, and compute_instance_from_template
				// into one KCC resource, ComputeInstance, and convert its metadata field from a map
				// to a structured list.
				outputCRDs = append(outputCRDs, generateComputeInstanceCRD(sm, rcs))
				continue
			case "ComputeInstanceTemplate":
				// We convert ComputeInstanceTemplate's metadata field from a map to a structured list.
				outputCRDs = append(outputCRDs, generateComputeInstanceTemplateCRD(sm, rcs))
				continue
			}

			// Another scenario of having multiple resource config mapping to the same kind is to consolidate
			// locational resources into a single CRD, e.g. compute_address and compute_global_address will be
			// merged to a single kind ComputeAddress
			isLocationalResource, err := isLocationalResource(rcs)
			if err != nil {
				log.Fatal(err)
			}
			crds := make([]*apiextensions.CustomResourceDefinition, 0)
			directCount := 0
			for _, rc := range rcs {
				// TODO: remove 'Direct' field from ResourceConfig and remove the if statement.
				// The 'Direct' indicator won't be needed after we finish all the migrations.
				// The 'Direct' indicator is necessary during the migration so
				// that Config Connector uses direct approach to generate CRDs
				// but still allow TF-based controller to reconcile the resource.
				if rc.Direct {
					fmt.Printf("skip generate TF-based CRD for direct resource %s\n", rc.Kind)
					directCount += 1
					continue
				}
				crd, err := crdgeneration.GenerateTF2CRD(&sm, rc)
				if err != nil {
					log.Fatalf("error generating CRD for %v: %v", rc.Name, err)
				}
				crds = append(crds, crd)
			}
			if directCount == len(rcs) {
				continue
			}
			crd, err := mergeCRDs(crds)
			if err != nil {
				log.Fatalf("error merging CRDs for kind %v: %v", kind, err)
			}
			if isLocationalResource {
				addLocationField(crd, rcs)
			}
			crd = addOneOfRulesForMultiTypeResourceReferences(crd, rcs)
			outputCRDs = append(outputCRDs, crd)
		}
	}
	return outputCRDs
}

func generateDCLBasedCRDs() []*apiextensions.CustomResourceDefinition {
	outputCRDs := make([]*apiextensions.CustomResourceDefinition, 0)
	serviceMetadataLoader := dclmetadata.New()
	schemaLoader, err := dclschemaloader.New()
	if err != nil {
		log.Fatalf("could not get a new DCL schema loader: %v", err)
	}
	smLoader, err := servicemappingloader.New()
	if err != nil {
		log.Fatalf("could not create service mapping loader: %v", err)
	}
	generator := crdgeneration.New(serviceMetadataLoader, schemaLoader, supportedgvks.AllWithoutDirect(smLoader, serviceMetadataLoader))
	gvks := supportedgvks.BasedOnDCL(serviceMetadataLoader)
	for _, gvk := range gvks {
		s, err := dclschemaloader.GetDCLSchemaForGVK(gvk, serviceMetadataLoader, schemaLoader)
		if err != nil {
			log.Fatalf("error getting the DCL schema for GVK %v: %v", gvk, err)
		}
		crd, err := generator.GenerateCRDFromOpenAPISchema(s, gvk)
		if err != nil {
			log.Fatalf("error generating CRD for %v: %v", gvk, err)
		}
		outputCRDs = append(outputCRDs, crd)
	}
	return outputCRDs
}

func generateCRDsForTypesFiles() []*apiextensions.CustomResourceDefinition {
	crds := make([]*apiextensions.CustomResourceDefinition, 0)
	tempDir, err := ioutil.TempDir("", "crds_for_types")
	if err != nil {
		log.Fatalf("error creating temporary directory: %v", err)
	}
	defer removeDir(tempDir)
	gens := genall.Generators{}
	crdGen := genall.Generator(crdgen.Generator{})
	gens = append(gens, &crdGen)
	rootPath := repo.GetRootOrLogFatal()
	apisPath := path.Join(rootPath, "pkg", "apis", "iam", "v1beta1")
	roots, err := gens.ForRoots(apisPath)
	if err != nil {
		log.Fatalf("error producing a Runtime to run the generators: %v", err)
	}
	roots.OutputRules = genall.OutputRules{Default: genall.OutputToDirectory(tempDir)}
	if failed := roots.Run(); failed {
		log.Fatalf("failed generating CRDs from Go structs")
	}
	files, err := ioutil.ReadDir(tempDir)
	if err != nil {
		log.Fatalf("error reading temporary directory: %v", err)
	}
	for _, f := range files {
		p := path.Join(tempDir, f.Name())
		b, err := ioutil.ReadFile(p)
		if err != nil {
			log.Fatalf("error reading file %v: %v", p, err)
		}
		crd := &apiextensions.CustomResourceDefinition{}
		if err = yaml.Unmarshal(b, crd); err != nil {
			log.Fatalf("error marshalling file %v into CRD: %v", p, err)
		}
		crds = append(crds, crd)
	}
	return crds
}

func removeDir(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		log.Fatalf("error removing directory '%v': %v", dir, err)
	}
}

// addOneOfRulesForMultiTypeResourceReferences returns a copy of the given CRD
// but with a modified JSON schema which makes it so that keys for resource
// references are mutually exclusive if a resource reference can have one of
// multiple keys (i.e. if the resource reference has multiple TypeConfigs)
func addOneOfRulesForMultiTypeResourceReferences(crd *apiextensions.CustomResourceDefinition, rcs []*corekccv1alpha1.ResourceConfig) *apiextensions.CustomResourceDefinition {
	outCRD := crd.DeepCopy()
	jsonSchema := k8s.GetOpenAPIV3SchemaFromCRD(outCRD)
	tfFieldsToReferenceKeys := getTFFieldsThatMapToMultipleReferenceKeys(rcs)
	for tfField, keys := range tfFieldsToReferenceKeys {
		field := append([]string{"spec"}, text.SnakeCaseStrsToLowerCamelCaseStrs(strings.Split(tfField, "."))...)
		// Enforces that one and only one key can be specified for the resource reference field
		oneOfRule := make([]apiextensions.JSONSchemaProps, 0)
		for _, key := range keys {
			oneOfRule = append(oneOfRule,
				apiextensions.JSONSchemaProps{
					Required: []string{key},
				},
			)
		}
		jsonSchema = setOneOfRuleForField(jsonSchema, field, oneOfRule)
	}
	k8s.PreferredVersion(outCRD).Schema.OpenAPIV3Schema = jsonSchema
	return outCRD
}

func getTFFieldsThatMapToMultipleReferenceKeys(rcs []*corekccv1alpha1.ResourceConfig) map[string][]string {
	tfFieldsToRefKeys := make(map[string][]string)
	for _, rc := range rcs {
		for _, refConfig := range rc.ResourceReferences {
			tfField := refConfig.TFField
			if _, ok := tfFieldsToRefKeys[tfField]; !ok {
				tfFieldsToRefKeys[tfField] = make([]string, 0)
			}
			for _, t := range refConfig.Types {
				// Keep the list of keys in sorted order to keep the output deterministic
				// (i.e. always the same output for the same input)
				tfFieldsToRefKeys[tfField] = slice.IncludeString(tfFieldsToRefKeys[tfField], t.Key)
			}
		}
	}
	for tfField, keys := range tfFieldsToRefKeys {
		if len(keys) <= 1 {
			delete(tfFieldsToRefKeys, tfField)
		}
	}
	return tfFieldsToRefKeys
}

func setOneOfRuleForField(s *apiextensions.JSONSchemaProps, field []string, oneOfRule []apiextensions.JSONSchemaProps) *apiextensions.JSONSchemaProps {
	outSchema := s.DeepCopy()
	subSchema := outSchema.Properties[field[0]]
	if len(field) > 1 {
		switch subSchema.Type {
		case "array":
			subSchema.Items.Schema = setOneOfRuleForField(subSchema.Items.Schema, field[1:], oneOfRule)
		case "object":
			subSchema = *setOneOfRuleForField(&subSchema, field[1:], oneOfRule)
		default:
			panic(fmt.Errorf("error parsing field %v: cannot iterate into type that is not object or array of objects", field))
		}
	} else {
		switch subSchema.Type {
		case "array":
			subSchema.Items.Schema.OneOf = oneOfRule
		default:
			subSchema.OneOf = oneOfRule
		}
	}
	outSchema.Properties[field[0]] = subSchema
	return outSchema
}

// addLocationField removes existing locational fields (region, zone) and replaces them with a 'location' field.
func addLocationField(crd *apiextensions.CustomResourceDefinition, rcs []*corekccv1alpha1.ResourceConfig) {
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec := schema.Properties["spec"]

	locationalFields := map[string]bool{
		"region": true,
		"zone":   true,
	}
	// It is assumed that locational fields (region, zone) would always be
	// at the base level.
	for field := range locationalFields {
		delete(spec.Properties, field)
	}
	requiredFields := make([]string, 0)
	for _, field := range spec.Required {
		if _, ok := locationalFields[field]; !ok {
			requiredFields = append(requiredFields, field)
		}
	}
	spec.Required = requiredFields

	locations := make(map[string]bool)
	for _, rc := range rcs {
		locations[rc.Locationality] = true
	}
	spec.Properties["location"] = apiextensions.JSONSchemaProps{
		Type:        "string",
		Description: generateLocationFieldDescription(locations, crd.Spec.Names.Kind),
	}
	spec.Required = slice.IncludeString(spec.Required, "location")

	schema.Properties["spec"] = spec
	schema.Required = slice.IncludeString(schema.Required, "spec")
}

func generateLocationFieldDescription(locations map[string]bool, resource string) string {
	description1 := fmt.Sprintf("Location represents the geographical location of the %v.", resource)
	description2 := ""
	if locations[gcp.Regional] {
		description2 = "Specify a region name"
	}
	if locations[gcp.Zonal] {
		if description2 == "" {
			description2 = "Specify a zone name"
		} else {
			description2 = description2 + " or a zone name"
		}
	}
	if locations[gcp.Global] {
		if description2 == "" {
			description2 = `Specify "global" for global resources`
		} else {
			description2 = description2 + ` or "global" for global resources`
		}
	}
	if locations[gcp.Regional] || locations[gcp.Zonal] {
		return fmt.Sprintf("%v %v. %v", description1, description2, "Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)")
	} else {
		return fmt.Sprintf("%v %v.", description1, description2)
	}
}

func isLocationalResource(rcs []*corekccv1alpha1.ResourceConfig) (bool, error) {
	m := map[string]bool{
		gcp.Regional: false,
		gcp.Global:   false,
		gcp.Zonal:    false,
	}
	if len(rcs) > 1 {
		for _, rc := range rcs {
			if _, ok := m[rc.Locationality]; !ok {
				return false, fmt.Errorf("unsupported locationality %v for kind %v is set", rc.Locationality, rc.Kind)
			}
			if m[rc.Locationality] {
				return false, fmt.Errorf("there is more than one resource config defined for the same locationality %v for kind %v", rc.Locationality, rc.Kind)
			}
		}
		return true, nil
	}
	if rcs[0].Locationality == "" {
		return false, nil
	}
	if _, ok := m[rcs[0].Locationality]; !ok {
		return false, fmt.Errorf("unsupported locationality %v for kind %v is set", rcs[0].Locationality, rcs[0].Kind)
	}
	return true, nil
}

func generateKindResourceConfigsMap(sm *corekccv1alpha1.ServiceMapping) map[string][]*corekccv1alpha1.ResourceConfig {
	res := make(map[string][]*corekccv1alpha1.ResourceConfig)
	for i, rc := range sm.Spec.Resources {
		if _, ok := res[rc.Kind]; !ok {
			res[rc.Kind] = make([]*corekccv1alpha1.ResourceConfig, 0)
		}
		res[rc.Kind] = append(res[rc.Kind], &sm.Spec.Resources[i])
	}
	return res
}

func mergeCRDs(crds []*apiextensions.CustomResourceDefinition) (*apiextensions.CustomResourceDefinition, error) {
	if len(crds) == 0 {
		return nil, fmt.Errorf("there is no CRD to merge")
	}
	if len(crds) == 1 {
		return crds[0], nil
	}

	var mergedCrd *apiextensions.CustomResourceDefinition
	for _, crd := range crds {
		if mergedCrd == nil {
			mergedCrd = crd.DeepCopy()
		} else {
			if !reflect.DeepEqual(mergedCrd.Name, crd.Name) {
				return nil, fmt.Errorf("couldn't merge crds with different names: %v, %v", mergedCrd.Name, crd.Name)
			}
			if err := mergeJSONSchemaProps(k8s.GetOpenAPIV3SchemaFromCRD(mergedCrd), k8s.GetOpenAPIV3SchemaFromCRD(crd)); err != nil {
				return nil, fmt.Errorf("couldn't merge crds for %v: %w", crd.Name, err)
			}
		}
	}
	return mergedCrd, nil
}

func mergeJSONSchemaProps(s1 *apiextensions.JSONSchemaProps, s2 *apiextensions.JSONSchemaProps) error {
	if s1.Type != s2.Type {
		return fmt.Errorf("types for the field value are different: one is %v and the other is %v", s1.Type, s2.Type)
	}

	if s1.Type == "array" {
		return mergeJSONSchemaPropsForArrayType(s1, s2)
	}

	return mergeJSONSchemaPropsForNonArrayType(s1, s2)
}

func mergeJSONSchemaPropsForArrayType(s1 *apiextensions.JSONSchemaProps, s2 *apiextensions.JSONSchemaProps) error {
	return mergeJSONSchemaPropsForNonArrayType(s1.Items.Schema, s2.Items.Schema)
}

func mergeJSONSchemaPropsForNonArrayType(s1 *apiextensions.JSONSchemaProps, s2 *apiextensions.JSONSchemaProps) error {
	commonRequired := intersection(s1.Required, s2.Required)
	s1.Required = commonRequired

	for k, v2 := range s2.Properties {
		if v1, ok := s1.Properties[k]; !ok {
			s1.Properties[k] = v2
		} else {
			if err := mergeJSONSchemaProps(&v1, &v2); err != nil {
				return fmt.Errorf("error merging JSON schema field %v: %w", k, err)
			}
			s1.Properties[k] = v1
		}
	}
	return nil
}

func intersection(a, b []string) []string {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}
	c := make([]string, 0)
	for _, item := range b {
		if _, ok := m[item]; ok {
			c = slice.IncludeString(c, item)
		}
	}
	return c
}

func generateComputeInstanceCRD(sm corekccv1alpha1.ServiceMapping, rcs []*corekccv1alpha1.ResourceConfig) *apiextensions.CustomResourceDefinition {
	crds := make([]*apiextensions.CustomResourceDefinition, 0)
	for _, rc := range rcs {
		crd, err := crdgeneration.GenerateTF2CRD(&sm, rc)
		if err != nil {
			log.Fatalf("error generating CRD for %v: %v", rc.Name, err)
		}
		crds = append(crds, crd)
	}
	crd, err := mergeComputeInstanceCRDs(crds)
	if err != nil {
		log.Fatalf("error merging CRDs for kind ComputeInstance: %v", err)
	}
	setCustomMetadataSchemaforComputeInstanceAndTemplate(crd)
	return crd
}

func mergeComputeInstanceCRDs(crds []*apiextensions.CustomResourceDefinition) (*apiextensions.CustomResourceDefinition, error) {
	if len(crds) != 2 {
		return nil, fmt.Errorf("there should be 2 ComputeInstance CRDs to merge")
	}

	var mergedCrd *apiextensions.CustomResourceDefinition
	for _, crd := range crds {
		if mergedCrd == nil {
			mergedCrd = crd.DeepCopy()
		} else {
			if mergedCrd.Name != crd.Name {
				return nil, fmt.Errorf("couldn't merge crds with different names: %v, %v", mergedCrd.Name, crd.Name)
			}
			mergeComputeInstanceJSONSchemaProps(k8s.GetOpenAPIV3SchemaFromCRD(mergedCrd), k8s.GetOpenAPIV3SchemaFromCRD(crd))
		}
	}
	return mergedCrd, nil
}

func mergeComputeInstanceJSONSchemaProps(s1 *apiextensions.JSONSchemaProps, s2 *apiextensions.JSONSchemaProps) {
	if !reflect.DeepEqual(s1.Required, s2.Required) {
		s1.AnyOf = []apiextensions.JSONSchemaProps{
			{
				Required: s1.Required,
			},
			{
				Required: s2.Required,
			},
		}
		s1.Required = nil
	}

	for k, v2 := range s2.Properties {
		if v1, ok := s1.Properties[k]; !ok {
			s1.Properties[k] = v2
		} else {
			mergeComputeInstanceJSONSchemaProps(&v1, &v2)
			s1.Properties[k] = v1
		}
	}
}

func generateComputeInstanceTemplateCRD(sm corekccv1alpha1.ServiceMapping, rcs []*corekccv1alpha1.ResourceConfig) *apiextensions.CustomResourceDefinition {
	if len(rcs) != 1 {
		log.Fatalf("expected only one resource config for compute instance template, got %v", len(rcs))
	}
	rc := rcs[0]
	crd, err := crdgeneration.GenerateTF2CRD(&sm, rc)
	if err != nil {
		log.Fatalf("error generating CRD for %v: %v", rc.Name, err)
	}
	setCustomMetadataSchemaforComputeInstanceAndTemplate(crd)
	return crd
}

func setCustomMetadataSchemaforComputeInstanceAndTemplate(crd *apiextensions.CustomResourceDefinition) {
	// In order to make setting custom metadata more in-line with environment-variable setting in
	// k8s pods, as well as leaving it extensible to add "valueFrom" in the future, we convert
	// the TF schema's map into a structured object array.
	metadataSchema := apiextensions.JSONSchemaProps{
		Type: "array",
		Items: &apiextensions.JSONSchemaPropsOrArray{
			Schema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"key":   {Type: "string"},
					"value": {Type: "string"},
				},
				Required: []string{"key", "value"},
			},
		},
	}
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	specSchema := schema.Properties["spec"]
	specSchema.Properties["metadata"] = metadataSchema
	schema.Properties["spec"] = specSchema
}

func outputCRDToFile(crd *apiextensions.CustomResourceDefinition) error {
	crdBytes, err := yaml.Marshal(crd)
	if err != nil {
		return err
	}
	outputFilename, err := crdgeneration.FileNameForCRD(crd)
	if err != nil {
		return err
	}
	outputFilepath := outputDir + "/" + outputFilename
	if err := ioutil.WriteFile(outputFilepath, crdBytes, fileMode); err != nil {
		return err
	}
	return nil
}
