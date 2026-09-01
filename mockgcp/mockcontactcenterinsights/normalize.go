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

package mockcontactcenterinsights

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "contactcenterinsights") {
		return
	}

	replacements.ReplacePath("createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath("updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath("activationUpdateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath("revisionCreateTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".activationUpdateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".revisionCreateTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.RemovePath("source")
	replacements.RemovePath(".source")
	replacements.RemovePath(".response.source")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "contactcenterinsights") {
		return
	}

	previsitValue := func(val string) {
		if strings.Contains(val, "/phraseMatchers/") {
			tokens := strings.Split(val, "/")
			for i := 0; i < len(tokens)-1; i++ {
				if tokens[i] == "phraseMatchers" {
					id := tokens[i+1]
					if idx := strings.Index(id, "?"); idx != -1 {
						id = id[:idx]
					}
					if isNumeric(id) {
						replacements.ReplaceStringValue(id, "${phraseMatcherID}")
					}
				}
			}
		}
	}

	previsitValue(event.URL())

	event.VisitRequestStringValues(func(path string, value string) {
		previsitValue(value)
	})

	event.VisitResponseStringValues(func(path string, value string) {
		previsitValue(value)
	})
}
