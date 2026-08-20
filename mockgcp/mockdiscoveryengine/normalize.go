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
	if !strings.Contains(url, "discoveryengine.googleapis.com") {
		return
	}
	replacements.ReplacePath(".startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.RemovePath(".servingConfigDataStore")
	replacements.RemovePath(".response.servingConfigDataStore")
	replacements.RemovePath(".marketplaceAgentVisibility")
	replacements.RemovePath(".response.marketplaceAgentVisibility")
	replacements.RemovePath(".observabilityConfig")
	replacements.RemovePath(".response.observabilityConfig")
	replacements.RemovePath(".naturalLanguageQueryUnderstandingConfig")
	replacements.RemovePath(".response.naturalLanguageQueryUnderstandingConfig")
	replacements.RemovePath(".solutionTypes")
	replacements.RemovePath(".response.solutionTypes")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "discoveryengine.googleapis.com") {
		return
	}

	normalizeID := func(path string, value string) {
		if path == ".name" || strings.HasSuffix(path, ".name") || path == "name" {
			tokens := strings.Split(value, "/")
			if len(tokens) >= 10 && tokens[len(tokens)-2] == "sessions" {
				sessionID := tokens[len(tokens)-1]
				replacements.ReplaceStringValue(sessionID, "session-${uniqueId}")
			}
			if len(tokens) >= 10 && tokens[len(tokens)-2] == "conversations" {
				conversationID := tokens[len(tokens)-1]
				replacements.ReplaceStringValue(conversationID, "conversation-${uniqueId}")
			}
		}
	}

	event.VisitResponseStringValues(normalizeID)
	event.VisitRequestStringValues(normalizeID)

	// Normalize numeric IDs in the URL if we haven't seen them in a "name" field yet
	url := event.URL()
	if idx := strings.Index(url, "?"); idx != -1 {
		url = url[:idx]
	}
	tokens := strings.Split(url, "/")
	for i := range tokens {
		if i > 0 && i < len(tokens)-1 {
			if tokens[i] == "sessions" || tokens[i] == "conversations" {
				id := tokens[i+1]
				// Basic check for numeric-ish ID
				if len(id) > 10 {
					if tokens[i] == "sessions" {
						replacements.ReplaceStringValue(id, "session-${uniqueId}")
					} else {
						replacements.ReplaceStringValue(id, "conversation-${uniqueId}")
					}
				}
			}
		}
	}
}
