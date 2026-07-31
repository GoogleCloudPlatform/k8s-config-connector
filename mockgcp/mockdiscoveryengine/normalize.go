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

package mockdiscoveryengine

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.RemovePath(".servingConfigDataStore")
	replacements.RemovePath(".response.servingConfigDataStore")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "discoveryengine.googleapis.com") {
		return
	}

	event.VisitResponseStringValues(func(path string, value string) {
		if path == ".name" || strings.HasSuffix(path, ".name") {
			tokens := strings.Split(value, "/")
			if len(tokens) >= 10 && tokens[len(tokens)-2] == "sessions" {
				sessionID := tokens[len(tokens)-1]
				// Normalize any generated numeric session ID to the expected test format
				replacements.ReplaceStringValue(sessionID, "session-${uniqueId}")
			}
		}
	})
}
