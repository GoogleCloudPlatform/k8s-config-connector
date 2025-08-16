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

package mockaccesscontextmanager

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

const placeholderEtag = "abcdef0123A"

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// accessPolicies
	replacements.ReplacePath(".accessPolicies[].etag", placeholderEtag)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	event.VisitResponseStringValues(func(path string, value string) {
		switch path {
		case ".response.name":
			tokens := strings.Split(value, "/")
			if len(tokens) == 2 && tokens[0] == "accessPolicies" {
				replacements.ReplaceStringValue(tokens[1], "${accessPolicyID}")
			}
		case ".name":
			tokens := strings.Split(value, "/")
			// accessPolicy operations look like operations/accessPolicies/${accessPolicyID}/create/${operationID}
			if len(tokens) == 5 && tokens[0] == "operations" && tokens[1] == "accessPolicies" {
				replacements.ReplaceStringValue(tokens[4], "${operationID}")
			}
			// servicePerimeters operations look like operations/accessPolicies/${accessPolicyID}/create/${operationID}
			if len(tokens) == 7 && tokens[0] == "operations" && tokens[1] == "accessPolicies" && tokens[3] == "servicePerimeters" {
				replacements.ReplaceStringValue(tokens[6], "${operationID}")
			}
		}
	})
}
