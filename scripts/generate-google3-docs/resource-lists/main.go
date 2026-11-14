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

	"k8s.io/apimachinery/pkg/runtime/schema"
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
	if err := generateListOfResourcesWithServiceGeneratedResourceIDv1alpha1(smLoader, serviceMetadataLoader, dclSchemaLoader); err != nil {
		log.Fatal(err)
	}
	if err := generateListOfUnacquirableResources(smLoader); err != nil {
		log.Fatal(err)
	}
}

// Use "service-generated" instead of "server-generated" in the output file
// names since the former is the terminology used by the public docs.

func generateListOfResourcesWithServiceGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) error {
	const outputFileName = "_resources-with-service-generated-resource-id.html"
	resources, err := resourcesWithServiceGeneratedResourceID(smLoader, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return fmt.Errorf("error getting resources with a service-generated resource ID: %w", err)
	}
	return generateListOfResources(resources, outputFileName)
}

func generateListOfResourcesWithServiceGeneratedResourceIDv1alpha1(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) error {
	const outputFileName = "_v1alpha1-resources-with-service-generated-resource-id.html"
	resources, err := resourcesWithServiceGeneratedResourceIDv1alpha1(smLoader, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return fmt.Errorf("error getting v1alpha1 resources with a service-generated resource ID: %w", err)
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

// resourcesWithServiceGeneratedResourceID returns all the v1beta1 resources
// (TF/DCL-based and direct) with a service-generated resource ID, deduplicated
// by GroupKind.
func resourcesWithServiceGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	tfDCLResources, err := tfDCLResourcesWithServiceGeneratedResourceID(smLoader, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return nil, err
	}
	directResources, err := directResourcesWithServiceGeneratedResourceID()
	if err != nil {
		return nil, err
	}
	return k8s.SortGVKsByKind(dedupeGroupKinds(append(tfDCLResources, directResources...))), nil
}

// resourcesWithServiceGeneratedResourceIDv1alpha1 returns all the v1alpha1
// resources (TF/DCL-based and direct) with a service-generated resource ID,
// deduplicated by GroupKind.
func resourcesWithServiceGeneratedResourceIDv1alpha1(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	tfDCLResources, err := tfDCLResourcesWithServiceGeneratedResourceIDv1alpha1(smLoader, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return nil, err
	}
	directResources, err := directResourcesWithServiceGeneratedResourceIDv1alpha1()
	if err != nil {
		return nil, err
	}
	resources := dedupeGroupKinds(append(tfDCLResources, directResources...))

	// Some kinds are still v1alpha1 in the TF/DCL metadata but have since been
	// promoted (e.g. to v1beta1 via a direct controller). Skip the kinds that
	// are already listed in the v1beta1 list.
	nonAlphaResources, err := resourcesWithServiceGeneratedResourceID(smLoader, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return nil, err
	}
	return k8s.SortGVKsByKind(excludeGroupKinds(resources, nonAlphaResources)), nil
}

// tfDCLResourcesWithServiceGeneratedResourceIDv1alpha1 returns the v1alpha1
// TF/DCL-based resources with a service-generated resource ID.
func tfDCLResourcesWithServiceGeneratedResourceIDv1alpha1(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	tfResources := tfBasedResourcesWithServiceGeneratedResourceIDv1alpha1(smLoader)
	dclResources, err := dclBasedResourcesWithServiceGeneratedResourceIDv1alpha1(serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting v1alpha1 DCL-based resources with a service-generated resource ID: %w", err)
	}
	return append(tfResources, dclResources...), nil
}

// tfDCLResourcesWithServiceGeneratedResourceID returns the v1beta1 TF/DCL-based
// resources with a service-generated resource ID.
func tfDCLResourcesWithServiceGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	tfResources := tfBasedResourcesWithServiceGeneratedResourceID(smLoader)
	dclResources, err := dclBasedResourcesWithServiceGeneratedResourceID(serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting DCL-based resources with a service-generated resource ID: %w", err)
	}
	return append(tfResources, dclResources...), nil
}

// directResourcesWithServiceGeneratedResourceID returns the v1beta1 direct
// resources with a service-generated resource ID.
func directResourcesWithServiceGeneratedResourceID() ([]schema.GroupVersionKind, error) {
	return directResourcesWithServiceGeneratedResourceIDForVersion(false)
}

// directResourcesWithServiceGeneratedResourceIDv1alpha1 returns the v1alpha1
// direct resources with a service-generated resource ID.
func directResourcesWithServiceGeneratedResourceIDv1alpha1() ([]schema.GroupVersionKind, error) {
	return directResourcesWithServiceGeneratedResourceIDForVersion(true)
}

func directResourcesWithServiceGeneratedResourceIDForVersion(v1alpha1 bool) ([]schema.GroupVersionKind, error) {
	all, err := allDirectResourcesWithServiceGeneratedResourceID()
	if err != nil {
		return nil, err
	}
	gvks := make([]schema.GroupVersionKind, 0)
	for _, gvk := range all {
		if isV1alpha1(gvk) != v1alpha1 {
			continue
		}
		gvks = append(gvks, gvk)
	}
	return gvks, nil
}

// dedupeGroupKinds returns resources with duplicate GroupKinds removed,
// keeping the first occurrence.
func dedupeGroupKinds(resources []schema.GroupVersionKind) []schema.GroupVersionKind {
	seen := make(map[schema.GroupKind]bool)
	deduped := make([]schema.GroupVersionKind, 0, len(resources))
	for _, gvk := range resources {
		if seen[gvk.GroupKind()] {
			continue
		}
		seen[gvk.GroupKind()] = true
		deduped = append(deduped, gvk)
	}
	return deduped
}

// excludeGroupKinds returns the GVKs in resources whose GroupKind does not
// appear in excluded.
func excludeGroupKinds(resources, excluded []schema.GroupVersionKind) []schema.GroupVersionKind {
	excludedGKs := make(map[schema.GroupKind]bool)
	for _, gvk := range excluded {
		excludedGKs[gvk.GroupKind()] = true
	}
	filtered := make([]schema.GroupVersionKind, 0, len(resources))
	for _, gvk := range resources {
		if excludedGKs[gvk.GroupKind()] {
			continue
		}
		filtered = append(filtered, gvk)
	}
	return filtered
}

func isV1alpha1(gvk schema.GroupVersionKind) bool {
	return gvk.Version == k8s.KCCAPIVersionV1Alpha1
}

// tfBasedResourcesWithServiceGeneratedResourceID returns the v1beta1
// TF-based resources with a service-generated resource ID.
func tfBasedResourcesWithServiceGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	return tfResourcesWithServiceGeneratedResourceID(smLoader, false)
}

// tfBasedResourcesWithServiceGeneratedResourceIDv1alpha1 returns the v1alpha1
// TF-based resources (both hand-written and auto-generated) with a
// service-generated resource ID.
func tfBasedResourcesWithServiceGeneratedResourceIDv1alpha1(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	return tfResourcesWithServiceGeneratedResourceID(smLoader, true)
}

// tfResourcesWithServiceGeneratedResourceID returns the TF-based resources with
// a service-generated resource ID.
func tfResourcesWithServiceGeneratedResourceID(smLoader *servicemappingloader.ServiceMappingLoader, v1alpha1 bool) []schema.GroupVersionKind {
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			if !krmtotf.SupportsResourceIDField(&rc) || !krmtotf.IsResourceIDFieldServerGenerated(&rc) {
				continue
			}
			gvk := tfmetadata.GVKForResource(&sm, &rc)
			if isV1alpha1(gvk) != v1alpha1 {
				continue
			}
			gvkList = append(gvkList, gvk)
		}
	}
	return gvkList
}

// dclBasedResourcesWithServiceGeneratedResourceID returns the v1beta1
// DCL-based resources with a service-generated resource ID.
func dclBasedResourcesWithServiceGeneratedResourceID(serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	return dclResourcesWithServiceGeneratedResourceID(serviceMetadataLoader, dclSchemaLoader, false)
}

// dclBasedResourcesWithServiceGeneratedResourceIDv1alpha1 returns the v1alpha1
// DCL-based resources with a service-generated resource ID.
func dclBasedResourcesWithServiceGeneratedResourceIDv1alpha1(serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ([]schema.GroupVersionKind, error) {
	return dclResourcesWithServiceGeneratedResourceID(serviceMetadataLoader, dclSchemaLoader, true)
}

// dclResourcesWithServiceGeneratedResourceID returns the DCL-based resources
// with a service-generated resource ID.
func dclResourcesWithServiceGeneratedResourceID(serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader, v1alpha1 bool) ([]schema.GroupVersionKind, error) {
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, s := range serviceMetadataLoader.GetAllServiceMetadata() {
		for _, r := range s.Resources {
			if !r.Releasable {
				continue
			}
			gvk := dclmetadata.GVKForResource(s, r)
			if isV1alpha1(gvk) != v1alpha1 {
				continue
			}
			schema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
			if err != nil {
				return nil, fmt.Errorf("error getting DCL schema for GroupVersionKind %v: %w", gvk, err)
			}
			nameFieldSchema, ok := dclextension.GetNameFieldSchema(schema)
			if !ok {
				// Resource doesn't support the resourceID field
				continue
			}
			isServiceGenerated, err := dclextension.IsResourceIDFieldServerGenerated(nameFieldSchema)
			if err != nil {
				return nil, fmt.Errorf("error determining if resourceID is service-generated for GroupVersionKind %v: %w", gvk, err)
			}
			if !isServiceGenerated {
				continue
			}
			gvkList = append(gvkList, gvk)
		}
	}
	return gvkList, nil
}

func unacquirableResources(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	// Only TF resources can be unacquirable because only TF resources can have
	// a service-generated ID that cannot be set via the resourceID field.
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
