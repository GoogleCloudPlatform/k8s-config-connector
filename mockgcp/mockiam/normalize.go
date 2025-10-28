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

package mockiam

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

const EtagPlaceholder = "abcdef0123A="
const TimePlaceholder = "2024-04-01T12:34:56.123456Z"
const UIDPlaceholder = "111111111111111111111"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// ServiceAccountKeys
	replacements.ReplacePath(".validAfterTime", "2024-04-01T12:34:56Z")
	replacements.ReplacePath(".keys[].validAfterTime", "2024-04-01T12:34:56Z")

	// Deny Policies
	replacements.ReplacePath(".policy.etag", EtagPlaceholder)
	replacements.ReplacePath(".policies[].etag", EtagPlaceholder)
	replacements.ReplacePath(".policies[].createTime", TimePlaceholder)
	replacements.ReplacePath(".policies[].updateTime", TimePlaceholder)
	replacements.ReplacePath(".policies[].deleteTime", TimePlaceholder)
	replacements.ReplacePath(".response.createTime", TimePlaceholder)
	replacements.ReplacePath(".response.updateTime", TimePlaceholder)
	replacements.ReplacePath(".response.deleteTime", TimePlaceholder)
	replacements.ReplacePath(".policies[].uid", UIDPlaceholder)
	replacements.ReplacePath(".uid", UIDPlaceholder)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// TODO: If all normalizers should filter by host, we could auto-filter (i.e. only call normalizers with matching hosts)
	if strings.Contains(event.URL(), "https://iam.googleapis.com") {
		event.VisitResponseStringValues(func(path string, value string) {
			switch path {
			case ".name":
				tokens := strings.Split(value, "/")
				if len(tokens) == 6 && tokens[4] == "keys" {
					replacements.ReplaceStringValue(tokens[5], "${serviceAccountKeyID}")
				}
			}
		})
	}
}
