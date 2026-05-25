// Copyright 2025 Google LLC
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

package mockdiscoveryengine

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "discoveryengine.googleapis.com/") {
		return
	}

	// TargetSite IDs are generated as fixed value in mock.
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "discoveryengine.googleapis.com/") {
		return
	}

	event.VisitResponseStringValues(func(path string, value string) {
		if strings.Contains(value, "/siteSearchEngine/targetSites/") {
			tokens := strings.Split(value, "/")
			if len(tokens) == 11 && tokens[9] == "targetSites" {
				id := tokens[10]
				if strings.HasPrefix(id, "targetsite-") {
					replacements.ReplaceStringValue(id, "targetsite-1234567890")
				}
			}
		}
	})
}
