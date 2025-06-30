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

	// // Node
	// replacements.ReplacePath(".createTime", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".response.createTime", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".updateTime", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".response.updateTime", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".id", "000000000000000000000")
	// replacements.ReplacePath(".response.id", "000000000000000000000")

	// replacements.ReplacePath(".networkEndpoints[].accessConfig.externalIp", "8.8.8.8")
	// replacements.ReplacePath(".response.networkEndpoints[].accessConfig.externalIp", "8.8.8.8")
	// replacements.ReplacePath(".networkEndpoints[].ipAddress", "10.20.30.40")
	// replacements.ReplacePath(".response.networkEndpoints[].ipAddress", "10.20.30.40")
	// replacements.ReplacePath(".cidrBlock", "10.20.30.0/24")
	// replacements.ReplacePath(".response.cidrBlock", "10.20.30.0/24")
	// // Bucket
	// replacements.ReplacePath(".softDeletePolicy.effectiveTime", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".timeCreated", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".updated", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".items[].timeCreated", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".items[].updated", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".acl[].etag", "abcdef0123A")
	// replacements.ReplacePath(".defaultObjectAcl[].etag", "abcdef0123A=")

	// // Managed Folder
	// replacements.ReplacePath(".items[].createTime", "2024-04-01T12:34:56.123456Z")
	// replacements.ReplacePath(".items[].updateTime", "2024-04-01T12:34:56.123456Z")

	// // Anywhere Cache
	// replacements.ReplacePath(".requestId", "1234")
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
