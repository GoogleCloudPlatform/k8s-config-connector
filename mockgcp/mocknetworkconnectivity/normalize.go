// Copyright 2026 Google LLC
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

package mocknetworkconnectivity

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func isNetworkConnectivityOperation(m map[string]any) bool {
	if metadata, ok := m["metadata"].(map[string]any); ok {
		if typeURL, ok := metadata["@type"].(string); ok {
			if strings.Contains(typeURL, "google.cloud.networkconnectivity") {
				return true
			}
		}
	}
	return false
}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "networkconnectivity.googleapis.com") {
		return
	}

	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".uniqueId", "111111111111111111111")

	replacements.TransformObject("", func(m map[string]any) {
		if linked, ok := m["linkedRouterApplianceInstances"].(map[string]any); ok {
			if instances, ok := linked["instances"].([]any); ok {
				for _, instAny := range instances {
					if inst, ok := instAny.(map[string]any); ok {
						if vm, ok := inst["virtualMachine"].(string); ok {
							inst["virtualMachine"] = strings.TrimPrefix(vm, "https://www.googleapis.com/compute/v1/")
						}
					}
				}
			}
		}
	})

	transformInternalRange := func(m map[string]any) {
		if m["prefixLength"] != nil && m["ipCidrRange"] != nil {
			if cidr, ok := m["ipCidrRange"].(string); ok {
				parts := strings.Split(cidr, "/")
				if len(parts) == 2 && strings.HasPrefix(parts[0], "10.0.") {
					m["ipCidrRange"] = "10.0.1.0/" + parts[1]
				}
			}
		}
	}
	replacements.TransformObject("", transformInternalRange)
	replacements.TransformObject(".response", transformInternalRange)

	replacements.TransformObject("", func(m map[string]any) {
		if !isNetworkConnectivityOperation(m) {
			return
		}
		// Clean up Operation metadata
		if m["metadata"] != nil {
			if metadata, ok := m["metadata"].(map[string]any); ok {
				delete(metadata, "requestedCancellation")
			}
			// Real GCP LRO response does not return labels (unlike GET / Create response)
			if response, ok := m["response"].(map[string]any); ok {
				delete(response, "labels")
			}
			if done, ok := m["done"].(bool); ok && !done {
				delete(m, "done")
			}
		}
	})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
