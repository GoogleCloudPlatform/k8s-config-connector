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

package asset

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
)

var locationalPrefixes = []string{"Global", "Region"}

func GetServiceMappingAndResourceConfig(smLoader *servicemappingloader.ServiceMappingLoader, asset *Asset) (*v1alpha1.ServiceMapping, *v1alpha1.ResourceConfig, error) {
	pieces := strings.Split(asset.AssetType, "/")
	if len(pieces) != 2 {
		return nil, nil, fmt.Errorf("unexpected format for asset type value '%v': expected 'service.googleapis.com/ResourceKind'", asset.AssetType)
	}
	serviceHostName := pieces[0]
	sm, err := smLoader.GetServiceMappingForServiceHostName(serviceHostName)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting service mapping: %w", err)
	}
	assetKind := pieces[1]
	resourceConfigs, err := getResourceConfigs(sm, assetKind)
	if err != nil {
		return nil, nil, err
	}
	if len(resourceConfigs) == 1 {
		return sm, &resourceConfigs[0], nil
	}
	rc, err := matchAssetNameToRC(asset.Name, resourceConfigs)
	if err != nil {
		return nil, nil, err
	}
	return sm, rc, err
}

func getResourceConfigs(sm *v1alpha1.ServiceMapping, rawKind string) ([]v1alpha1.ResourceConfig, error) {
	resourceKind := strings.ToLower(getResourceKind(rawKind))
	var results []v1alpha1.ResourceConfig
	for _, rc := range sm.Spec.Resources {
		if !shouldUseResourceConfig(rc) {
			continue
		}
		rcKind := strings.TrimPrefix(rc.Kind, sm.Spec.Name)
		if resourceKind == strings.ToLower(rcKind) {
			results = append(results, rc)
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no resource config found in '%v' for '%v'", sm.Name, rawKind)
	}
	return results, nil
}

func shouldUseResourceConfig(rc v1alpha1.ResourceConfig) bool {
	// google instance from template is only used internally when creating an instance from a template and assets
	// should not be matched to it
	if rc.Name == "google_compute_instance_from_template" {
		return false
	}
	return true
}

func getResourceKind(rawKind string) string {
	for _, prefix := range locationalPrefixes {
		if strings.HasPrefix(rawKind, prefix) {
			return strings.TrimPrefix(rawKind, prefix)
		}
	}
	// TODO: this is hard-coded until we rename IAMCustomRole to IAMRole at which point this can be removed
	if rawKind == "Role" {
		return "CustomRole"
	}
	return rawKind
}

// TODO: do something for compute instance here
func matchAssetNameToRC(assetName string, resourceConfigs []v1alpha1.ResourceConfig) (*v1alpha1.ResourceConfig, error) {
	for _, rc := range resourceConfigs {
		regexIDTemplate := idTemplateToRegex(rc.IDTemplate)
		regex, err := regexp.Compile(regexIDTemplate)
		if err != nil {
			return nil, fmt.Errorf("error compiling '%v' to regex: %w", regexIDTemplate, err)
		}
		if regex.MatchString(assetName) {
			return &rc, nil
		}
	}
	return nil, fmt.Errorf("unable to find a match resource config for asset name '%v'", assetName)
}

var replaceRegex = regexp.MustCompile("/{{.*?}}")

// idTemplateToRegex replaces all fields in the id template that are 'user input' with wildcards, i.e. the id template,
//
//	projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}
//
// becomes
//
//	projects/.*/regions/.*/forwardingRules/.*
func idTemplateToRegex(idTemplate string) string {
	return replaceRegex.ReplaceAllString(idTemplate, ".*")
}
