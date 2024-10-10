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

package uri

import (
	"fmt"
	"regexp"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
)

func GetServiceMappingAndResourceConfig(smLoader *servicemappingloader.ServiceMappingLoader, urlHost, urlPath string) (*v1alpha1.ServiceMapping, *v1alpha1.ResourceConfig, error) {
	sm, err := smLoader.GetServiceMappingForServiceHostName(urlHost)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting service mapping: %w", err)
	}
	rc, err := matchResourceNameToRC(urlPath, sm)
	if err != nil {
		return nil, nil, err
	}
	return sm, rc, err
}

func matchResourceNameToRC(uriPath string, sm *v1alpha1.ServiceMapping) (*v1alpha1.ResourceConfig, error) {
	if isHierarchicalResource(sm, uriPath) {
		return matchResourceNameToHierarchalRC(uriPath, sm)
	}
	return matchResourceNameToRCGeneral(uriPath, sm)
}

func matchResourceNameToRCGeneral(uriPath string, sm *v1alpha1.ServiceMapping) (*v1alpha1.ResourceConfig, error) {
	for _, rc := range sm.Spec.Resources {
		if rc.Direct {
			continue
		}
		if !*rc.IDTemplateCanBeUsedToMatchResourceName {
			continue
		}
		// resources that skip import often have id templates that are broad and match too many resources, examples:
		// * google_compute_network_peering
		if rc.SkipImport {
			continue
		}
		regexIDTemplate := idTemplateToRegex(rc.IDTemplate)
		regex, err := regexp.Compile(regexIDTemplate)
		if err != nil {
			return nil, fmt.Errorf("error compiling '%v' to regex: %w", regexIDTemplate, err)
		}
		if regex.MatchString(uriPath) {
			return &rc, nil
		}
	}
	return nil, fmt.Errorf("unable to find a matching resource config for uri path '%v'", uriPath)
}

func matchResourceNameToHierarchalRC(urlPath string, sm *v1alpha1.ServiceMapping) (*v1alpha1.ResourceConfig, error) {
	resourceName := getHierarchicalResourceName(urlPath)
	for _, rc := range sm.Spec.Resources {
		if rc.Kind == resourceName {
			return &rc, nil
		}
	}
	return nil, fmt.Errorf("unable to find a hierarchal resource with name '%v' in service mapping '%v'", resourceName, sm.Spec.Name)
}

var replaceRegex = regexp.MustCompile("(^|/){{.*?}}")

// idTemplateToRegex replaces all fields in the id template that are 'user input' with a wildcard that matches everything except
// the '/' character, i.e. the id template,
//
//	projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}
//
// becomes
//
//	projects/[^/]*/regions/[^/]*/forwardingRules/[^/]*
//
// the expression [^/] is a negative match, meaning, match everything that is not a '/'
func idTemplateToRegex(idTemplate string) string {
	// add ($|\?.*) to the end of the id template to ensure that the string we are matching 'ends' with either the
	// last portion of the id template *or* has a query param (the ?). This is to prevent id templates which are a
	// substring of another id template from matching each other.
	//
	// Example, BigQueryDataset's id template is a substring of BigQueryTable's id template.
	// * BigQueryDataset: projects/{{project}}/datasets/{{dataset_id}}
	// * BigQueryTable:   projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
	idTemplate = fmt.Sprintf("%v($|\\?.*)", idTemplate)
	return replaceRegex.ReplaceAllString(idTemplate, "$1[^/]*")
}

var (
	projectsURLRegex = regexp.MustCompile("(/.*)?/projects/.*")
	foldersURLRegex  = regexp.MustCompile("(/.*)?/folders/.*")
)

func getHierarchicalResourceName(urlPath string) string {
	if foldersURLRegex.MatchString(urlPath) {
		return "Folder"
	}
	return "Project"
}

func isHierarchicalResource(sm *v1alpha1.ServiceMapping, urlPath string) bool {
	if sm.Spec.Name != "ResourceManager" {
		return false
	}
	if projectsURLRegex.MatchString(urlPath) {
		return true
	}
	if foldersURLRegex.MatchString(urlPath) {
		return true
	}
	return false
}
