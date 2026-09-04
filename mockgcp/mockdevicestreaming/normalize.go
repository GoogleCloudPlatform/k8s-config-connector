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

package mockdevicestreaming

import (
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".expireTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".activeStartTime", mockgcpregistry.PlaceholderTimestamp)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// Only apply normalization if the request is for this service
	if !strings.Contains(event.URL(), "devicestreaming.googleapis.com") {
		return
	}

	// Capture any session ID pattern "session-[a-z0-9]{13}" and replace it with "session-${uniqueId}"
	sessionIDRegex := regexp.MustCompile(`session-[a-z0-9]{13}`)
	event.VisitResponseStringValues(func(path string, value string) {
		for _, match := range sessionIDRegex.FindAllString(value, -1) {
			replacements.ReplaceStringValue(match, "session-${uniqueId}")
		}
	})
}
