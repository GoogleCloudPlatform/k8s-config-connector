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

const TimePlaceholder = "2024-04-01T12:34:56.123456Z"

const EtagPlaceholder = "abcdef0123A="

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// accessPolicy
	replacements.ReplacePath(".response.createTime", TimePlaceholder)
	replacements.ReplacePath(".createTime", TimePlaceholder)
	replacements.ReplacePath(".response.updateTime", TimePlaceholder)
	replacements.ReplacePath(".updateTime", TimePlaceholder)

	replacements.ReplacePath(".etag", EtagPlaceholder)
	replacements.ReplacePath(".accessPolicies[].etag", EtagPlaceholder)

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// ctx := context.TODO()
	// log := klog.FromContext(ctx)

	// // TODO: If all normalizers should filter by host, we could auto-filter (i.e. only call normalizers with matching hosts)
	if strings.Contains(event.URL(), "https://accesscontextmanager.googleapis.com") {

		event.VisitResponseStringValues(func(path string, value string) {
			switch path {
			case ".response.name", ".name":
				tokens := strings.Split(value, "/")

				if len(tokens) == 2 && tokens[0] == "accessPolicies" {
					replacements.ReplaceStringValue(tokens[1], "${accessPolicyID}")
				}

				// Odd operation ID format
				if len(tokens) == 5 && tokens[0] == "operations" && tokens[1] == "accessPolicies" && (tokens[3] == "create" || tokens[3] == "update" || tokens[3] == "delete") {
					replacements.ReplaceStringValue(tokens[4], "${operationID}")
				}
				if len(tokens) == 7 && tokens[0] == "operations" && tokens[1] == "accessPolicies" && tokens[3] == "accessLevels" && (tokens[5] == "create" || tokens[5] == "update" || tokens[5] == "delete") {
					replacements.ReplaceStringValue(tokens[6], "${operationID}")
				}
			}
		})
	}
}
