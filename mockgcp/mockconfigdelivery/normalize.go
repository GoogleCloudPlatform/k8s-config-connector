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

package mockconfigdelivery

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func isConfigDeliveryOperation(m map[string]any) bool {
	if metadata, ok := m["metadata"].(map[string]any); ok {
		if typeURL, ok := metadata["@type"].(string); ok {
			if strings.Contains(typeURL, "google.cloud.configdelivery") {
				return true
			}
		}
	}
	return false
}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "configdelivery.googleapis.com") {
		return
	}

	replacements.ReplacePath("createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath("updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath("endTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.updateTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.TransformObject("", func(m map[string]any) {
		if !isConfigDeliveryOperation(m) {
			return
		}
		// Clean up Operation metadata
		if m["metadata"] != nil {
			if metadata, ok := m["metadata"].(map[string]any); ok {
				delete(metadata, "requestedCancellation")
			}
			if done, ok := m["done"].(bool); ok && !done {
				delete(m, "done")
			}
		}
	})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
