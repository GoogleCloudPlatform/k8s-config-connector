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

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/resource-autogen/allowlist"
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/resource-autogen/sampleconversion"
	autogenloader "github.com/GoogleCloudPlatform/k8s-config-connector/scripts/resource-autogen/servicemapping/servicemappingloader"

	"github.com/hashicorp/go-multierror"
	"github.com/tmccombs/hcl2json/convert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

var (
	randomSuffixKeyword = "%{random_suffix}"
	uniqueIDHolder      = "${uniqueID}"
	// tfReservedFields are reserved field names for TF related metadata.
	// related to TF GCP types.
	tfReservedFields = []string{
		"provider",
		"lifecycle",
		"depends_on",
	}
	// tfOnlyTypes are non GCP TF types.
	tfOnlyTypes = map[string]bool{
		"time_sleep": true,
	}
	// nonIDKRMFieldsRequiringUniqueValuesMap is the map of KRM kinds and the
	// map of their non-ID string fields that require unique values.
	nonIDKRMFieldsRequiringUniqueValuesMap = map[string]map[string]bool{
		"TagsTagKey": {
			"shortName": true,
		},
		"AccessContextManagerServicePerimeter": {
			"title": true,
		},
	}
	// krmFieldsNotAllowingSpecialCharsMap is the map of KRM kinds and the map
	// of their string fields that don't allow special characters in the value.
	// This is needed for the sample converter to explicitly clean up any
	// special characters in the value.
	krmFieldsNotAllowingSpecialCharsMap = map[string]map[string]bool{
		"Project": {
			"name": true,
		},
	}
	// additionalRequiredFieldsMap is the map of KRM kinds and the maps of their
	// required string fields and the default values that are not specified in
	// the TF sample.
	// This is needed so that the sample converter can add required fields and
	// values.
	additionalRequiredFieldsMap = map[string]map[string]string{
		"DataCatalogTaxonomy": {
			"region": "us",
		},
	}
	// defaultOrganizationalResourcesMap is the map of organizational KRM kinds
	// and the relative resource names of the default test instances.
	// There are some organizational resources that KCC shouldn't touch in the
	// integration test. When those resources are needed as the dependency,
	// instead of creating new ones, we should use the default/pre-created ones
	// instead.
	defaultOrganizationalResourcesMap = map[string]string{
		"AccessContextManagerAccessPolicy": "accessPolicies/578359180191",
	}
	// ListFieldsWithAtMostOneItemMap is the map of KRM kinds and the maps of
	// their list fields that support at most one item.
	// TF resources have many list fields that support at most one item, and KCC
	// turns those list fields into object fields. Sample converter should be
	// able to identify those fields and turn them from lists to objects in the
	// converted YAML.
	ListFieldsWithAtMostOneItemMap = map[string]map[string]bool{
		"AccessContextManagerServicePerimeter": {
			"status": true,
		},
	}
	// ResourceIDLengthMap is the map of KRM kinds and the length limitation of
	// their resource IDs.
	// Some resources have more strict ID length limitations than others. KCC's
	// test resourceID may not always fit. This map is used to track the lengths
	// of resource IDs that may cause issues if not properly handled.
	ResourceIDLengthMap = map[string]int{
		"AccessContextManagerServicePerimeter": 50,
	}
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		return fmt.Errorf("error getting new service mapping loader: %w", err)
	}
	generatedSMMap, err := autogenloader.GetGeneratedSMMap()
	if err != nil {
		return fmt.Errorf("error getting the generated ServiceMapping map: %w", err)
	}
	autoGenAllowlist, err := allowlist.LoadAutoGenAllowList(generatedSMMap)
	if err != nil {
		return fmt.Errorf("error loading allowlist for autogen resources: %w", err)
	}
	tfToGVK, err := GetTFTypeToGVKMap(smLoader)
	if err != nil {
		return fmt.Errorf("error getting TF type mapping: %w", err)
	}
	err = convertTFSamplesToKRMTestdata(tfToGVK, smLoader, autoGenAllowlist)
	if err != nil {
		return fmt.Errorf("error converting TF samples:\n%w", err)
	}
	return nil
}

func convertTFSamplesToKRMTestdata(tfToGVK map[string]schema.GroupVersionKind, smLoader *servicemappingloader.ServiceMappingLoader, autoGenAllowlist *allowlist.AutoGenAllowlist) error {
	var errs *multierror.Error
	samplesPath := repo.GetAutoGeneratedTFSamplesPathOrFatal()
	sampleFolders, err := fileutil.SubdirsIn(samplesPath)
	if err != nil {
		return fmt.Errorf("error reading directory %v: %w", samplesPath, err)
	}
	generatedSamples := make(map[string]bool)
	for _, sf := range sampleFolders {
		originalSF := sf
		sf = strings.Replace(sf, "dry-run", "dry_run", -1)
		klog.Infof("Converting TF sample %v (original name: %v)...", sf, originalSF)
		sampleNameInfo := strings.Split(sf, "-")
		if len(sampleNameInfo) < 3 || len(sampleNameInfo) > 4 {
			errs = multierror.Append(errs,
				fmt.Errorf("sample folder name should be in the format of '[Service]-[Kind]-[sample_name]' or '[Service]-[Kind]-[sample_name]-skipped', but it's %v", sf))
			return errs
		}
		service := sampleNameInfo[0]
		kind := sampleNameInfo[1]
		group := fmt.Sprintf("%s.cnrm.cloud.google.com", strings.ToLower(service))

		autoGenType, ok := autoGenAllowlist.GetKRMKind(kind)
		if !ok {
			klog.Infof("Skipping the parse of sample %v. Kind %v not allowlisted.", sf, kind)
			continue
		}
		sm, err := smLoader.GetServiceMapping(group)
		if err != nil {
			// TODO(b/265225406): Check error type.
			klog.Infof("Skipping the parse of sample %v. Group %v not found.", sf, group)
			continue
		}
		rc := servicemappingloader.GetResourceConfigsForKind(sm, kind)
		if rc == nil || len(rc) == 0 {
			klog.Infof("Skipping the parse of sample %v. Kind %v not found in service mapping.", sf, kind)
			continue
		}
		// Auto-generated resources should have one-on-one mapping between kind
		// and resource configs.
		if len(rc) > 1 {
			errs = multierror.Append(errs,
				fmt.Errorf("error retrieving resource configs for "+
					"kind %v, there should only be one matching resource config", kind))
			return errs
		}
		// If the TF type for the sample is not an auto-generated kind, no need
		// to parse the sample.
		if !rc[0].AutoGenerated {
			klog.Infof("Skipping the parse of sample %v. Kind %v is not auto-generated.", sf, kind)
			continue
		}
		// If it is an organizational resource that shouldn't be created, no
		// need to parse the sample.
		if _, ok := defaultOrganizationalResourcesMap[kind]; ok {
			klog.Infof("Skipping the parse of sample %v. Creating a resource of kind %v will impact the use of the test organization.", sf, kind)
			continue
		}

		sampleName := text.SnakeCaseToLowerCase(sampleNameInfo[2])
		// Focus on basic samples for now.
		if !strings.HasSuffix(sampleName, "basic") {
			klog.Infof("Skipping the parse of sample %v. This is not a basic sample.", sf)
			continue
		}

		path := filepath.Join(samplesPath, sf, "main.tf")
		b, err := os.ReadFile(path)
		if err != nil {
			errToReturn := fmt.Errorf("error reading file %v for TF sample %s: %w", path, sf, err)
			klog.Warningf("Failed sample conversion: %v", errToReturn)
			errs = multierror.Append(errs, errToReturn)
			continue
		}

		jsonStruct, err := convertHCLBytesToJSON(b)
		if err != nil {
			errToReturn := fmt.Errorf("error converting HCL to JSON for TF sample %s: %w", sf, err)
			klog.Warningf("Failed sample conversion: %v", errToReturn)
			errs = multierror.Append(errs, errToReturn)
			continue
		}

		create, dependencies, err := tfSampleToKRMTestData(kind, jsonStruct, tfToGVK, smLoader)
		if err != nil {
			errToReturn := fmt.Errorf("error converting TF samples to KRM test data for TF sample %s: %w", sf, err)
			klog.Warningf("Failed sample conversion: %v", errToReturn)
			errs = multierror.Append(errs, errToReturn)
			continue
		}

		// Change 'basic' suffix in sampleName to 'autogen', to avoid name conflict with Direct Controller basic test.
		// Note: We have not used this script for a while, so I manually updated the names of the existing autogen tests.
		// I made the change just in case we want to reuse it in the future.
		sampleName = strings.TrimSuffix(sampleName, "basic") + "autogen"

		if err := insertTestData(create, dependencies, autoGenType, sampleName, generatedSamples); err != nil {
			errToReturn := fmt.Errorf("error unmarshaling json for TF sample %s: %w", sf, err)
			klog.Warningf("Failed sample conversion: %v", errToReturn)
			errs = multierror.Append(errs, errToReturn)
			continue
		}

		klog.Infof("Sample %v converted successfully!", sf)
	}

	return errs.ErrorOrNil()
}

func convertHCLBytesToJSON(raw []byte) (map[string]interface{}, error) {
	lines := strings.Split(string(raw), "\n")
	hcl := ""
	for _, s := range lines {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) == 0 || trimmed == "```hcl" || trimmed == "```" {
			continue
		}
		hcl += s + "\n"
	}
	hcl = strings.TrimSuffix(hcl, "\n")

	// To bypass the "Invalid Template Control Keyword" error.
	if strings.Contains(hcl, randomSuffixKeyword) {
		hcl = strings.ReplaceAll(hcl, randomSuffixKeyword, uniqueIDHolder)
	}

	input := []byte(hcl)
	convertedBytes, err := convert.Bytes(input, "", convert.Options{})
	if err != nil {
		return nil, fmt.Errorf("error parsing bytes: %w", err)
	}

	jsonStruct := make(map[string]interface{})
	err = json.Unmarshal(convertedBytes, &jsonStruct)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}

	return jsonStruct, nil
}

func tfSampleToKRMTestData(testKind string, tf map[string]interface{}, tfToGVK map[string]schema.GroupVersionKind, smLoader *servicemappingloader.ServiceMappingLoader) (create map[string]interface{}, dependencies []map[string]interface{}, err error) {
	resourcesRaw, ok := tf["resource"]
	if !ok {
		return nil, nil, fmt.Errorf("tf struct should contain a 'resource' field: %+v", tf)
	}
	resources, ok := resourcesRaw.(map[string]interface{})
	if !ok {
		return nil, nil, fmt.Errorf("value of 'resource' should be in the format of 'map[string]interface{}' but not %T", resourcesRaw)
	}

	create = make(map[string]interface{})
	dependencyMap := make(map[string]map[string]interface{})
	dependencyGraph := sampleconversion.NewDependencyGraph()
	for tfType, resource := range resources {
		gvk, ok := tfToGVK[tfType]
		if !ok {
			if _, ok := tfOnlyTypes[tfType]; ok {
				continue
			}
			return nil, nil, fmt.Errorf("TF type %v doesn't exist in the service mappings", tfType)
		}
		// No need to parse the config for organizational resources that
		// shouldn't be created.
		if _, ok := defaultOrganizationalResourcesMap[gvk.Kind]; ok {
			continue
		}

		sm, err := smLoader.GetServiceMapping(gvk.Group)
		if err != nil {
			return nil, nil, err
		}
		rc, err := servicemappingloader.GetResourceConfigsForTFType(sm, tfType)
		if err != nil {
			return nil, nil, err
		}

		krmConfig, err := tfConfigToKRMConfig(resource, tfType, dependencyGraph, *rc, tfToGVK)
		if err != nil {
			return nil, nil, fmt.Errorf("error converting TF config to KRM for resource\n%+v\n:\n%w", resource, err)
		}
		if gvk.Kind == testKind {
			if len(create) != 0 {
				return nil, nil, fmt.Errorf("more than one resource of type %s exists, but there should be only one", tfType)
			}
			create = krmConfig
		} else {
			dependencyMap[tfType] = krmConfig
		}
	}

	sortedRefDependencyTypes := dependencyGraph.TopologicalSort()
	dependencies = make([]map[string]interface{}, 0)
	for _, tfType := range sortedRefDependencyTypes {
		gvk, ok := tfToGVK[tfType]
		if !ok {
			return nil, nil, fmt.Errorf("TF type %v doesn't exist in the service mappings", tfType)
		}
		// The "create" struct is covered in the dependencyGraph when there is
		// any dependency, but it shouldn't be added to "dependencies" struct
		// list.
		if gvk.Kind == testKind {
			continue
		}
		dependency, ok := dependencyMap[tfType]
		if !ok {
			return nil, nil, fmt.Errorf("TF type %v doesn't exist in the dependencyMap", tfType)
		}
		dependencies = append(dependencies, dependency)
		delete(dependencyMap, tfType)
	}
	// dependencyGraph only covers TF types that are involved in references. We
	// still need to go through other configs whose TF types are not involved in
	// any reference.
	sortedNonRefDependencyTypes := make([]string, 0)
	for tfType, _ := range dependencyMap {
		sortedNonRefDependencyTypes = append(sortedNonRefDependencyTypes, tfType)
	}
	sort.Strings(sortedNonRefDependencyTypes)
	for _, tfType := range sortedNonRefDependencyTypes {
		dependency, ok := dependencyMap[tfType]
		if !ok {
			return nil, nil, fmt.Errorf("TF type %v doesn't exist in the dependencyMap", tfType)
		}
		dependencies = append(dependencies, dependency)
	}
	return create, dependencies, nil
}

func tfConfigToKRMConfig(tfConfig interface{}, tfType string, dependencyGraph *sampleconversion.DependencyGraph, rc v1alpha1.ResourceConfig, tfToGVK map[string]schema.GroupVersionKind) (spec map[string]interface{}, err error) {
	klog.V(2).Infof("tfConfig: %+v\n", tfConfig)

	name, specs, containerAnnotation, err := cleanupTFFields(tfConfig, tfType, dependencyGraph, rc, tfToGVK)
	if err != nil {
		return nil, fmt.Errorf("error cleaning up the TF config: %w", err)
	}
	// TODO(b/265367038): Handle the samples with multiple resources of the same type.\
	gvk, ok := tfToGVK[tfType]
	if !ok {
		return nil, fmt.Errorf("TF type %v doesn't exist in the service mappings", tfType)
	}
	return handleKRMFields(specs[name], containerAnnotation, gvk, rc)
}

func cleanupTFFields(configRaw interface{}, tfType string, dependencyGraph *sampleconversion.DependencyGraph, rc v1alpha1.ResourceConfig, tfToGVK map[string]schema.GroupVersionKind) (name string, specs map[string]map[string]interface{}, containerAnnotation map[string]string, err error) {
	config, ok := configRaw.(map[string]interface{})
	if !ok {
		return "", nil, nil, fmt.Errorf("TF config should be in the format of 'map[string]interface{}' but not %T", configRaw)
	}
	if len(config) != 1 {
		return "", nil, nil, fmt.Errorf("there should be only 1 element, but got %v", len(config))
	}

	name = reflect.ValueOf(config).MapKeys()[0].String()
	specRaw := config[name]
	specArray, ok := specRaw.([]interface{})
	if !ok {
		return "", nil, nil, fmt.Errorf("value of '%s' should be in the format of '[]interface{}' but not %T", name, specRaw)
	}
	if len(specArray) != 1 {
		return "", nil, nil, fmt.Errorf("there should be only 1 element, but got %v", len(specArray))
	}
	spec, ok := specArray[0].(map[string]interface{})
	if !ok {
		return "", nil, nil, fmt.Errorf("configuration value should be in the format of 'map[string]interface{}' but not %T", specArray[0])
	}

	for _, field := range tfReservedFields {
		_, ok := spec[field]
		if ok {
			delete(spec, field)
		}
	}

	additionalRequiredFields, ok := additionalRequiredFieldsMap[rc.Kind]
	if ok {
		for f, v := range additionalRequiredFields {
			if _, ok := spec[f]; !ok {
				spec[f] = v
			}
		}
	}

	krmSpec, containerAnnotation, err := krmifySpec(spec, tfType, dependencyGraph, rc, tfToGVK)
	if err != nil {
		return "", nil, nil, fmt.Errorf("error krmifying the spec %+v: %w", spec, err)
	}
	specs = make(map[string]map[string]interface{})
	specs[name] = krmSpec
	return name, specs, containerAnnotation, nil
}

func krmifySpec(tfSpec map[string]interface{}, tfType string, dependencyGraph *sampleconversion.DependencyGraph, rc v1alpha1.ResourceConfig, tfToGVK map[string]schema.GroupVersionKind) (krmSpec map[string]interface{}, containerAnnotation map[string]string, err error) {
	krmSpec = make(map[string]interface{})
	containerAnnotation = make(map[string]string)
	refConfigMap := getReferenceConfigMap(rc)
	containerMap := getContainerMap(rc)
	for tfFieldName, value := range tfSpec {
		krmFieldName := text.SnakeCaseToLowerCamelCase(tfFieldName)
		tfRefVal, valueTemplate, containsTFRef, err := sampleconversion.GetTFReferenceValue(value)
		if err != nil {
			return nil, nil, fmt.Errorf("error getting TF reference value for field %v: %w", tfFieldName, err)
		}
		if containsTFRef {
			dependencyGraph.AddDependencyWithTFRefVal(tfRefVal, tfType)
			// The use case that the field has a reference in the Tf sample, but
			// is not a reference field in the KRM resource can't be handled.
			refConfig, ok := refConfigMap[tfFieldName]
			if !ok {
				krmRefVal, err := sampleconversion.ConstructKRMExternalRefValFromTFRefVal(tfRefVal, valueTemplate, tfToGVK)
				if err != nil {
					return nil, nil, fmt.Errorf("cannot construct KRM value for a TF reference field %v: %w", krmFieldName, err)
				}
				krmSpec[krmFieldName] = krmRefVal
				continue
			}

			krmFieldName = refConfig.Key
			// For organizational resources that shouldn't be created, but used
			// as a reference, use the default one instead.
			defaultVal, ok := defaultOrganizationalResourcesMap[refConfig.GVK.Kind]
			if ok {
				krmRefVal := sampleconversion.ConstructKRMExternalReferenceObject(defaultVal)
				krmSpec[krmFieldName] = krmRefVal
				continue
			}

			krmRefVal, err := sampleconversion.ConstructKRMNameReferenceObject(tfRefVal, tfToGVK)
			if err != nil {
				return nil, nil, fmt.Errorf("error constructing KRM reference value for field %v: %w", krmFieldName, err)
			}
			krmSpec[krmFieldName] = krmRefVal
			continue
		}
		if isProjectNameWithNumber(value) {
			testProjectNameWithNumber := "projects/${projectNumber}"
			testProjectID := "${projectId}"
			// It's possible that the field with a value of the relative
			// resource name of a GCP project is a reference field in KRM.
			refConfig, ok := refConfigMap[tfFieldName]
			if !ok {
				container, ok := containerMap[tfFieldName]
				if !ok {
					krmSpec[krmFieldName] = testProjectNameWithNumber
					continue
				}
				if !isProjectContainer(container) {
					return nil, nil, fmt.Errorf("expected container type for field %v to be project but is %+v", tfFieldName, container.Type)
				}
				if len(containerAnnotation) > 0 {
					return nil, nil, fmt.Errorf("more than one container annotation found: '%+v' and '%v: %v'", containerAnnotation, tfFieldName, testProjectNameWithNumber)
				}
				containerAnnotation[k8s.GetAnnotationForContainerType(container.Type)] = testProjectID
				continue
			}
			krmFieldName = refConfig.Key
			krmRefVal := sampleconversion.ConstructKRMExternalReferenceObject(testProjectNameWithNumber)
			krmSpec[krmFieldName] = krmRefVal
			continue
		}
		if isOrganizationName(value) {
			testOrgID := "${TEST_ORG_ID}"
			testOrgName := fmt.Sprintf("organizations/%v", testOrgID)
			// It's possible that the field with a value of the relative
			// resource name of a GCP organization is a reference field in KRM.
			refConfig, ok := refConfigMap[tfFieldName]
			if !ok {
				container, ok := containerMap[tfFieldName]
				if !ok {
					krmSpec[krmFieldName] = testOrgName
					continue
				}
				if !isOrganizationContainer(container) {
					return nil, nil, fmt.Errorf("expected container type for field %v to be organization but is %+v", tfFieldName, container.Type)
				}
				if len(containerAnnotation) > 0 {
					return nil, nil, fmt.Errorf("more than one container annotation found: '%+v' and '%v: %v'", containerAnnotation, tfFieldName, testOrgName)
				}
				containerAnnotation[k8s.GetAnnotationForContainerType(container.Type)] = testOrgID
				continue
			}
			krmFieldName = refConfig.Key
			krmRefVal := sampleconversion.ConstructKRMExternalReferenceObject(testOrgName)
			krmSpec[krmFieldName] = krmRefVal
			continue
		}
		result, err := krmifyNestedField(value)
		if err != nil {
			return nil, nil, fmt.Errorf("error krmifying the nested field %s: %w", krmFieldName, err)
		}
		krmSpec[krmFieldName] = result
	}
	return krmSpec, containerAnnotation, nil
}

func krmifyNestedField(value interface{}) (interface{}, error) {
	switch value.(type) {
	case []interface{}:
		arrayValue := value.([]interface{})
		krmArray := make([]interface{}, 0)
		for _, v := range arrayValue {
			result, err := krmifyNestedField(v)
			if err != nil {
				return nil, fmt.Errorf("error krmifying the array field: %w", err)
			}
			krmArray = append(krmArray, result)
		}
		return krmArray, nil
	case map[string]interface{}:
		mapValue := value.(map[string]interface{})
		krmMap := make(map[string]interface{})
		for k, v := range mapValue {
			krmFieldName := text.SnakeCaseToLowerCamelCase(k)
			result, err := krmifyNestedField(v)
			if err != nil {
				return nil, fmt.Errorf("error krmifying the object field: %w", err)
			}
			krmMap[krmFieldName] = result
		}
		return krmMap, nil
	default:
		// TODO(b/265367198): Handle nested reference fields.
		return value, nil
	}
}

func handleKRMFields(spec map[string]interface{}, containerAnnotation map[string]string, gvk schema.GroupVersionKind, rc v1alpha1.ResourceConfig) (map[string]interface{}, error) {
	krmStruct := make(map[string]interface{})
	krmStruct["apiVersion"] = gvk.GroupVersion().String()
	krmStruct["kind"] = gvk.Kind

	metadata := make(map[string]interface{})
	metadata["name"] = fmt.Sprintf("%s-${uniqueID}", strings.ToLower(gvk.Kind))

	// Fields mapping to `metadata.name` and `metadata.labels` should be removed
	// from `spec`.
	nameField := text.SnakeCaseToLowerCamelCase(rc.MetadataMapping.Name)
	labelsField := text.SnakeCaseToLowerCamelCase(rc.MetadataMapping.Labels)
	if _, ok := spec[nameField]; ok {
		delete(spec, nameField)
	}
	labels, ok := spec[labelsField]
	if ok {
		metadata["labels"] = labels
		delete(spec, labelsField)
	}

	if len(containerAnnotation) > 1 {
		return nil, fmt.Errorf("more than one container annotation provided: %+v", containerAnnotation)
	}
	for key, value := range containerAnnotation {
		annotations := make(map[string]interface{})
		annotations[key] = value
		metadata["annotations"] = annotations
	}
	krmStruct["metadata"] = metadata

	// Setting alphanumeric resourceID for all resources that support the
	// user-specified resourceID field to avoid the edge cases when a resource
	// has a different naming convention from K8s objects'.
	if rc.ResourceID.TargetField != "" && rc.ServerGeneratedIDField == "" {
		resourceID, err := generateResourceID(gvk.Kind)
		if err != nil {
			return nil, fmt.Errorf("error generating resource ID for kind %v: %w", gvk.Kind, err)
		}
		spec["resourceID"] = resourceID
	}

	// Hierarchical references should be represented as reference fields instead.
	var hierarchicalReferenceConfigured bool
	supportedHierarchicalReferenceTypes := make(map[v1alpha1.HierarchicalReferenceType]bool)
	for _, hr := range rc.HierarchicalReferences {
		refConfig, err := krmtotf.GetReferenceConfigForHierarchicalReference(hr, &rc)
		if err != nil {
			return nil, fmt.Errorf("error retrieving reference config: %w", err)
		}
		supportedHierarchicalReferenceTypes[hr.Type] = true
		tfField := text.SnakeCaseToLowerCamelCase(refConfig.TFField)
		_, ok := spec[tfField]
		if !ok {
			continue
		}
		switch hr.Type {
		case v1alpha1.HierarchicalReferenceTypeProject:
			refVal := make(map[string]interface{})
			refVal["name"] = "project-${uniqueID}"
			spec["projectRef"] = refVal
		case v1alpha1.HierarchicalReferenceTypeFolder:
			spec["folderRef"] = map[string]string{"external": "${TEST_FOLDER_ID}"}
		case v1alpha1.HierarchicalReferenceTypeOrganization:
			spec["organizationRef"] = map[string]string{"external": "${TEST_ORG_ID}"}
		default:
			return nil, fmt.Errorf("unsupported hierarchical reference type: %v", hr.Type)
		}
		delete(spec, tfField)
		hierarchicalReferenceConfigured = true
	}
	// If a resource has hierarchical reference field(s), but the field is not
	// explicitly configured, it means that the TF sample uses the default
	// project configured by the TF provider.
	// We need to add the `projectRef` field explicitly.
	if _, ok := supportedHierarchicalReferenceTypes[v1alpha1.HierarchicalReferenceTypeProject]; ok && !hierarchicalReferenceConfigured {
		spec["projectRef"] = map[string]string{"external": "${projectId}"}
	}

	// TODO(b/265367279): Handle nested special field.
	spec = handleSpecialTopLevelFields(spec, gvk.Kind)
	krmStruct["spec"] = spec

	return krmStruct, nil
}

func handleSpecialTopLevelFields(spec map[string]interface{}, kind string) map[string]interface{} {
	nonIDFieldsRequiringUniqueValues, _ := nonIDKRMFieldsRequiringUniqueValuesMap[kind]
	fieldsNotAllowingSpecialChars, _ := krmFieldsNotAllowingSpecialCharsMap[kind]
	listFieldsWithAtMostOneItemMap, _ := ListFieldsWithAtMostOneItemMap[kind]
	if len(nonIDFieldsRequiringUniqueValues) == 0 &&
		len(fieldsNotAllowingSpecialChars) == 0 &&
		len(listFieldsWithAtMostOneItemMap) == 0 {
		return spec
	}

	updatedSpec := make(map[string]interface{})
	for fieldName, value := range spec {
		switch value.(type) {
		case string:
			strVal := value.(string)

			if len(nonIDFieldsRequiringUniqueValues) > 0 {
				if _, ok := nonIDFieldsRequiringUniqueValues[fieldName]; ok {
					strVal += uniqueIDHolder
				}
			}

			if len(fieldsNotAllowingSpecialChars) > 0 {
				if _, ok := fieldsNotAllowingSpecialChars[fieldName]; ok {
					strVal = text.RemoveSpecialCharacters(strVal)
				}
			}
			updatedSpec[fieldName] = strVal
		case []interface{}:
			listVal := value.([]interface{})
			if _, ok := listFieldsWithAtMostOneItemMap[fieldName]; ok {
				updatedSpec[fieldName] = listVal[0]
			}
		default:
			updatedSpec[fieldName] = value
		}
	}
	return updatedSpec
}

func GetTFTypeToGVKMap(smLoader *servicemappingloader.ServiceMappingLoader) (map[string]schema.GroupVersionKind, error) {
	tfTypeToGVK := make(map[string]schema.GroupVersionKind)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			tfType := rc.Name
			gvk := schema.GroupVersionKind{
				Group:   sm.Name,
				Version: sm.GetVersionFor(&rc),
				Kind:    rc.Kind,
			}
			tfTypeToGVK[tfType] = gvk
		}
	}
	return tfTypeToGVK, nil
}

func insertTestData(createConfig map[string]interface{}, dependenciesConfig []map[string]interface{}, autoGenType *allowlist.AutoGenType, sampleName string, generatedSamples map[string]bool) error {
	folderPath := getTestDataFolderPath(autoGenType)

	createFilePath := filepath.Join(folderPath, sampleName, "create.yaml")
	if err := os.MkdirAll(filepath.Dir(createFilePath), 0770); err != nil {
		return fmt.Errorf("error creating folder for path %v: %w", createFilePath, err)
	}
	createConfigInBytes, err := yaml.Marshal(createConfig)
	if err != nil {
		return fmt.Errorf("err marshaling createConfig to yaml: %w", err)
	}
	if err := os.WriteFile(createFilePath, createConfigInBytes, 0644); err != nil {
		return fmt.Errorf("error writing to file %v: %w", createFilePath, err)
	}

	if len(dependenciesConfig) > 0 {
		dependenciesFilePath := filepath.Join(folderPath, sampleName, "dependencies.yaml")
		var dependenciesConfigInBytes []byte

		for i, r := range dependenciesConfig {
			resourceConfigInBytes, err := yaml.Marshal(r)
			if err != nil {
				return fmt.Errorf("err marshaling resource config in dependencies to yaml: %w", err)
			}
			if i != 0 {
				yamlSeparator := []byte("---\n")
				dependenciesConfigInBytes = append(dependenciesConfigInBytes, yamlSeparator...)
			}
			dependenciesConfigInBytes = append(dependenciesConfigInBytes, resourceConfigInBytes...)
		}
		if err := os.WriteFile(dependenciesFilePath, dependenciesConfigInBytes, 0644); err != nil {
			return fmt.Errorf("error writing to file %v: %w", dependenciesFilePath, err)
		}
	}
	return nil
}

func getTestDataFolderPath(autoGenType *allowlist.AutoGenType) string {
	serviceFolderName := autoGenType.ServiceNameInLC
	kindFolderName := strings.ToLower(autoGenType.KRMKindName)
	return filepath.Join(repo.GetBasicIntegrationTestDataPath(), serviceFolderName, autoGenType.Version, kindFolderName)
}

func getReferenceConfigMap(rc v1alpha1.ResourceConfig) map[string]v1alpha1.ReferenceConfig {
	refConfigMap := make(map[string]v1alpha1.ReferenceConfig)
	for _, refConfig := range rc.ResourceReferences {
		refConfigMap[refConfig.TFField] = refConfig
	}
	return refConfigMap
}

func getContainerMap(rc v1alpha1.ResourceConfig) map[string]v1alpha1.Container {
	containerMap := make(map[string]v1alpha1.Container)
	for _, container := range rc.Containers {
		containerMap[container.TFField] = container
	}
	return containerMap
}

func isOrganizationName(value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		return false
	}
	orgNameRegex := regexp.MustCompile(`^organizations/[0-9]{5,15}$`)
	matchResult := orgNameRegex.FindStringSubmatch(str)
	if len(matchResult) == 1 {
		return true
	}
	return false
}

func isProjectNameWithNumber(value interface{}) bool {
	str, ok := value.(string)
	if !ok {
		return false
	}
	projectNameRegex := regexp.MustCompile(`^projects/[0-9]{5,15}$`)
	matchResult := projectNameRegex.FindStringSubmatch(str)
	if len(matchResult) == 1 {
		return true
	}
	return false
}

func isProjectContainer(container v1alpha1.Container) bool {
	switch container.Type {
	case v1alpha1.ContainerTypeProject:
		return true
	default:
		return false
	}
}

func isOrganizationContainer(container v1alpha1.Container) bool {
	switch container.Type {
	case v1alpha1.ContainerTypeOrganization:
		return true
	default:
		return false
	}
}

func generateResourceID(kind string) (string, error) {
	supportedLength, ok := ResourceIDLengthMap[kind]
	resourceID := fmt.Sprintf("%s${uniqueID}", strings.ToLower(kind))
	if !ok {
		return resourceID, nil
	}

	// The generated unique ID has 20 characters.
	// If the supported length of the ID for the resource is shorter than 20
	// characters, we can't convert the resource sample.
	// Otherwise, remove the first X letters of the resourceID to fit.
	if supportedLength <= 20 {
		return "", fmt.Errorf("supported resource ID supportedLength should > 20")
	}
	resourceIDLength := len(kind) + 20
	var numberOfLettersToRemove int
	if resourceIDLength > supportedLength {
		numberOfLettersToRemove = resourceIDLength - supportedLength
	}
	return resourceID[numberOfLettersToRemove:], nil
}
