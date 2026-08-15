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
	replacements.TransformLRO(func(m map[string]any) {
		if resp, ok := m["response"].(map[string]any); ok {
			if len(resp) == 0 || (len(resp) == 1 && resp["@type"] == "type.googleapis.com/google.protobuf.Empty") {
				delete(m, "response")
			}
		}
	})

	transformFunc := func(m map[string]any) {
		name, _ := m["name"].(string)
		if strings.Contains(name, "/engines/") {
			// For Engines, the real log does not have createTime, updateTime, marketplaceAgentVisibility, observabilityConfig
			delete(m, "createTime")
			delete(m, "updateTime")
			delete(m, "marketplaceAgentVisibility")
			delete(m, "observabilityConfig")
		} else if strings.Contains(name, "/dataStores/") {
			// For DataStores, the real log has createTime, naturalLanguageQueryUnderstandingConfig, solutionTypes
			if m["createTime"] != nil {
				m["createTime"] = mockgcpregistry.PlaceholderTimestamp
			}
		}
	}

	replacements.TransformObject("", transformFunc)
	replacements.TransformObject(".response", transformFunc)

	// Since we got review feedback not to change common HTTP error mapping files,
	// we normalize the error response for discoveryengine here!
	replacements.TransformObject(".error", func(m map[string]any) {
		if _, exists := m["status"]; !exists {
			if code, ok := m["code"].(float64); ok && code == 500 {
				m["status"] = "INTERNAL"
			}
		}
	})
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
