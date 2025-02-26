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

package mocktpu

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Node
	replacements.ReplacePath(".createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".id", "000000000000000000000")
	replacements.ReplacePath(".response.id", "000000000000000000000")

	replacements.ReplacePath(".networkEndpoints[].accessConfig.externalIp", "8.8.8.8")
	replacements.ReplacePath(".response.networkEndpoints[].accessConfig.externalIp", "8.8.8.8")
	replacements.ReplacePath(".networkEndpoints[].ipAddress", "10.20.30.40")
	replacements.ReplacePath(".response.networkEndpoints[].ipAddress", "10.20.30.40")
	replacements.ReplacePath(".cidrBlock", "10.20.30.0/24")
	replacements.ReplacePath(".response.cidrBlock", "10.20.30.0/24")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
