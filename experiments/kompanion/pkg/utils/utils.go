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

package utils

import (
	"fmt"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	discovery "k8s.io/client-go/discovery"
	"k8s.io/klog/v2"
)

func GetResources(discoveryClient discovery.DiscoveryInterface, resources []schema.GroupVersionResource) ([]schema.GroupVersionResource, error) {
	apiResourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return nil, fmt.Errorf("failed to get preferred resources: %w", err)
	}

	for _, apiResourceList := range apiResourceLists {
		if !strings.Contains(apiResourceList.GroupVersion, ".cnrm.cloud.google.com/") {

			continue
		}

		apiResourceListGroupVersion, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			klog.Warningf("skipping unparseable groupVersion %q", apiResourceList.GroupVersion)
			continue
		}

		for _, apiResource := range apiResourceList.APIResources {
			if !apiResource.Namespaced {

				continue
			}
			if !contains(apiResource.Verbs, "list") {

				continue
			}

			gvr := schema.GroupVersionResource{
				Group:    apiResource.Group,
				Version:  apiResource.Version,
				Resource: apiResource.Name,
			}

			if gvr.Group == "" {

				gvr.Group = apiResourceListGroupVersion.Group
			}

			if gvr.Version == "" {

				gvr.Version = apiResourceListGroupVersion.Version
			}

			resources = append(resources, gvr)
		}
	}
	// Improve determinism for debuggability and idempotency
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].String() < resources[j].String()
	})
	return resources, nil
}

// contains checks if a slice contains a specific string.
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if strings.ToLower(s) == strings.ToLower(str) {
			return true
		}
	}
	return false
}
