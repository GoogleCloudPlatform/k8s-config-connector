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

func (s *MockService) ConfigureVisitor(url string, visitor mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "discoveryengine.googleapis.com") {
		return
	}

	visitor.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.updateTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".metadata.updateTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, visitor mockgcpregistry.NormalizingVisitor) {
	if name, ok := event.GetResponseStringValue(".name"); ok {
		tokens := strings.Split(name, "/")
		if len(tokens) == 11 && tokens[8] == "siteSearchEngine" && tokens[9] == "targetSites" {
			targetSiteID := tokens[10]
			visitor.ReplaceStringValue(targetSiteID, "1780450824484668193")
		}
	}
	if name, ok := event.GetResponseStringValue(".response.name"); ok {
		tokens := strings.Split(name, "/")
		if len(tokens) == 11 && tokens[8] == "siteSearchEngine" && tokens[9] == "targetSites" {
			targetSiteID := tokens[10]
			visitor.ReplaceStringValue(targetSiteID, "1780450824484668193")
		}
	}
	url := event.URL()
	tokens := strings.Split(url, "/")
	for i, t := range tokens {
		if t == "targetSites" && i+1 < len(tokens) {
			targetSiteID := strings.Split(tokens[i+1], "?")[0]
			targetSiteID = strings.Split(targetSiteID, "%2F")[0]
			if targetSiteID != "" {
				visitor.ReplaceStringValue(targetSiteID, "1780450824484668193")
			}
		}
	}
}
