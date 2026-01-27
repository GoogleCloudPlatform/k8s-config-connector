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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	tfmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
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

	if err := generateListOfResourcesWithServiceGeneratedResourceID(); err != nil {
		log.Fatal(err)
	}
	if err := generateListOfUnacquirableResources(smLoader); err != nil {
		log.Fatal(err)
	}
}

func generateListOfResourcesWithServiceGeneratedResourceID() error {
	// Use "service-generated" instead of "server-generated" in the output file
	// name since the former is the terminology used by the public docs.
	const outputFileName = "_resources-with-service-generated-resource-id.html"
	resources, err := resourcesWithServerGeneratedResourceID()
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

func resourcesWithServerGeneratedResourceID() ([]schema.GroupVersionKind, error) {
	crdPath := repo.GetCRDsPath()
	entries, err := os.ReadDir(crdPath)
	if err != nil {
		return nil, fmt.Errorf("error reading CRD directory %v: %w", crdPath, err)
	}

	gvkList := make([]schema.GroupVersionKind, 0)
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}

		bytes, err := os.ReadFile(filepath.Join(crdPath, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("error reading CRD file %v: %w", entry.Name(), err)
		}

		var crd apiextensionsv1.CustomResourceDefinition
		if err := yaml.Unmarshal(bytes, &crd); err != nil {
			return nil, fmt.Errorf("error unmarshalling CRD %v: %w", entry.Name(), err)
		}

		if hasServiceGeneratedResourceID(&crd) {
			gvk := schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: "v1beta1",
				Kind:    crd.Spec.Names.Kind,
			}
			gvkList = append(gvkList, gvk)
		}
	}

	return k8s.SortGVKsByKind(gvkList), nil
}

func hasServiceGeneratedResourceID(crd *apiextensionsv1.CustomResourceDefinition) bool {
	for _, v := range crd.Spec.Versions {
		if v.Name != "v1beta1" {
			continue
		}
		if v.Schema == nil || v.Schema.OpenAPIV3Schema == nil {
			continue
		}
		props := v.Schema.OpenAPIV3Schema.Properties
		if spec, ok := props["spec"]; ok {
			if resourceID, ok := spec.Properties["resourceID"]; ok {
				// Search for keyword "service-generated" in resourceID description in CRDs.
				// todo: This requires accurate field description. Do we have a better way to determine resources with service-generated ID?
				if strings.Contains(resourceID.Description, "service-generated") {
					return true
				}
			}
		}
	}
	return false
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
