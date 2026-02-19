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

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	tfmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf/metadata"

	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

type templateData struct {
	Resources []schema.GroupVersionKind
}

func main() {
	if err := clearGeneratedDocDir(); err != nil {
		log.Fatal(fmt.Errorf("error clearing generated doc dir: %w", err))
	}

	smLoader, err := servicemappingloader.New()
	if err != nil {
		log.Fatal(fmt.Errorf("error creating service mapping loader: %w", err))
	}
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		log.Fatal(fmt.Errorf("error creating a DCL schema loader: %w", err))
	}

	if err := generateListOfResourcesWithServiceGeneratedResourceID(smLoader, serviceMetadataLoader, dclSchemaLoader); err != nil {
		log.Fatal(err)
	}
	if err := generateListOfUnacquirableResources(smLoader); err != nil {
		log.Fatal(err)
	}
}

func generateListOfResourcesWithServiceGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) error {
	// Use "service-generated" instead of "server-generated" in the output file
	// name since the former is the terminology used by the public docs.
	const outputFileName = "_resources-with-service-generated-resource-id.html"
	resources, err := resourcesWithServerGeneratedResourceID(smLoader, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return fmt.Errorf("error getting resources with a server-generated resource ID: %w", err)
	}
	return generateListOfResources(resources, outputFileName)
}

func generateListOfUnacquirableResources(smLoader *servicemappingloader.ServiceMappingLoader) error {
	const outputFileName = "_unacquirable-resources.html"
	resources := unacquirableResources(smLoader)
	return generateListOfResources(resources, outputFileName)
}

func generateListOfResources(resources []schema.GroupVersionKind, outputFileName string) error {
	outputFilePath := path.Join(repo.GetG3ResourceListsGeneratedPath(), outputFileName)
	outputFile, err := fileutil.NewEmptyFile(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating empty file for output: %w", err)
	}

	templateFilePath := repo.GetG3ResourceListsTemplatePath()
	template, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("error parsing template file: %w", err)
	}
	templateData := &templateData{
		Resources: resources,
	}
	if err := template.Execute(outputFile, templateData); err != nil {
		return fmt.Errorf("error while executing template: %w", err)
	}
	return nil
}

func resourcesWithServerGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	tfResources := tfBasedResourcesWithServerGeneratedResourceID(smLoader)
	dclResources, err := dclBasedResourcesWithServerGeneratedResourceID(serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting DCL-based resources with a server-generated resource ID: %w", err)
	}
	resourcesFromCRDs, err := resourcesWithServerGeneratedResourceIDFromCRDs()
	if err != nil {
		return nil, fmt.Errorf("error getting resources with a server-generated resource ID from CRDs: %w", err)
	}

	allResources := append(append(tfResources, dclResources...), resourcesFromCRDs...)
	return k8s.SortGVKsByKind(deduplicateGVKs(allResources)), nil
}

func resourcesWithServerGeneratedResourceIDFromCRDs() ([]schema.GroupVersionKind, error) {
	crdsPath := repo.GetCRDsPath()
	files, err := os.ReadDir(crdsPath)
	if err != nil {
		return nil, fmt.Errorf("error reading CRDs directory: %w", err)
	}

	gvkList := make([]schema.GroupVersionKind, 0)
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".yaml" {
			continue
		}
		crdPath := filepath.Join(crdsPath, file.Name())
		crdBytes, err := os.ReadFile(crdPath)
		if err != nil {
			return nil, fmt.Errorf("error reading CRD file %s: %w", crdPath, err)
		}
		var crd apiextensions.CustomResourceDefinition
		if err := yaml.Unmarshal(crdBytes, &crd); err != nil {
			return nil, fmt.Errorf("error unmarshalling CRD file %s: %w", crdPath, err)
		}

		found := false
		for _, version := range crd.Spec.Versions {
			jsonSchema := version.Schema.OpenAPIV3Schema
			if jsonSchema == nil {
				continue
			}

			spec, ok := jsonSchema.Properties["spec"]
			if !ok {
				continue
			}
			resourceID, ok := spec.Properties["resourceID"]
			if !ok {
				continue
			}
			description := strings.ToLower(resourceID.Description)
			if strings.Contains(description, "service-generated") || strings.Contains(description, "server-generated") {
				found = true
				break
			}
		}

		if found {
			gvk := schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: k8s.GetVersionFromCRD(&crd),
				Kind:    crd.Spec.Names.Kind,
			}
			gvkList = append(gvkList, gvk)
		}
	}
	return gvkList, nil
}

func deduplicateGVKs(gvks []schema.GroupVersionKind) []schema.GroupVersionKind {
	uniqueGVKs := make(map[schema.GroupVersionKind]bool)
	result := make([]schema.GroupVersionKind, 0)
	for _, gvk := range gvks {
		if !uniqueGVKs[gvk] {
			uniqueGVKs[gvk] = true
			result = append(result, gvk)
		}
	}
	return result
}

func tfBasedResourcesWithServerGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			if krmtotf.SupportsResourceIDField(&rc) && krmtotf.IsResourceIDFieldServerGenerated(&rc) && !rc.AutoGenerated {
				gvk := tfmetadata.GVKForResource(&sm, &rc)
				gvkList = append(gvkList, gvk)
			}
		}
	}
	return gvkList
}

func dclBasedResourcesWithServerGeneratedResourceID(serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, s := range serviceMetadataLoader.GetAllServiceMetadata() {
		for _, r := range s.Resources {
			if !r.Releasable {
				continue
			}
			gvk := dclmetadata.GVKForResource(s, r)
			schema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
			if err != nil {
				return nil, fmt.Errorf("error getting DCL schema for GroupVersionKind %v: %w", gvk, err)
			}
			nameFieldSchema, ok := dclextension.GetNameFieldSchema(schema)
			if !ok {
				// Resource doesn't support the resourceID field
				continue
			}
			isServerGenerated, err := dclextension.IsResourceIDFieldServerGenerated(nameFieldSchema)
			if err != nil {
				return nil, fmt.Errorf("error determining if resourceID is server-generated for GroupVersionKind %v: %w", gvk, err)
			}
			if !isServerGenerated {
				continue
			}
			gvkList = append(gvkList, gvk)
		}
	}
	return gvkList, nil
}

func unacquirableResources(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	// Only TF resources can be unacquirable because only TF resources can have
	// a server-generated ID that cannot be set via the resourceID field.
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			if !krmtotf.SupportsResourceIDField(&rc) && krmtotf.SupportsServerGeneratedIDField(&rc) {
				gvk := tfmetadata.GVKForResource(&sm, &rc)
				gvkList = append(gvkList, gvk)
			}
		}
	}
	return k8s.SortGVKsByKind(gvkList)
}

func clearGeneratedDocDir() error {
	docDir := repo.GetG3ResourceListsGeneratedPath()
	if err := os.RemoveAll(docDir); err != nil {
		return fmt.Errorf("error deleting generated doc dir at %v: %w", docDir, err)
	}
	if err := os.Mkdir(docDir, 0700); err != nil {
		return fmt.Errorf("error recreating generated doc dir at %v: %w", docDir, err)
	}
	return nil
}
