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

package mocknetworkservices

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

const TimePlaceholder = "2024-04-01T12:34:56.123456Z"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".createTime", TimePlaceholder)
	replacements.ReplacePath(".updateTime", TimePlaceholder)
	replacements.ReplacePath(".endTime", TimePlaceholder)

	replacements.ReplacePath(".lbRouteExtensions[].createTime", TimePlaceholder)
	replacements.ReplacePath(".lbRouteExtensions[].updateTime", TimePlaceholder)

	replacements.ReplacePath(".metadata.createTime", TimePlaceholder)
	replacements.ReplacePath(".metadata.endTime", TimePlaceholder)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	event.VisitResponseStringValues(func(path string, value string) {
		if strings.HasPrefix(value, "https://compute.googleapis.com/compute/v1/") {
			newValue := strings.Replace(value, "https://compute.googleapis.com/compute/v1/", "https://www.googleapis.com/compute/v1/", 1)
			replacements.ReplaceStringValue(value, newValue)
		}
	})
}
