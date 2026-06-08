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

package mockredis

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "redis.googleapis.com") {
		return
	}
	replacements.ReplacePath(".pscConnections[].address", "10.11.12.13")
	replacements.ReplacePath(".response.pscConnections[].address", "10.11.12.13")
	replacements.ReplacePath(".discoveryEndpoints[].address", "10.11.12.13")
	replacements.ReplacePath(".response.discoveryEndpoints[].address", "10.11.12.13")
	replacements.ReplacePath(".clusterEndpoints[].connections[].pscConnection.address", "10.11.12.13")
	replacements.ReplacePath(".response.clusterEndpoints[].connections[].pscConnection.address", "10.11.12.13")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "redis.googleapis.com") {
		return
	}
	event.VisitResponseStringValues(func(path string, value string) {
		if strings.HasSuffix(path, "pscConnectionId") || strings.HasSuffix(path, "pscConnectionID") {
			replacements.ReplaceStringValue(value, "${pscConnectionID}")
		}
		if strings.Contains(value, "serviceAttachments/ssc-auto-sa-") {
			tokens := strings.Split(value, "/")
			for _, token := range tokens {
				if strings.HasPrefix(token, "ssc-auto-sa-") {
					replacements.ReplaceStringValue(token, "ssc-auto-sa-abcde")
				}
			}
		}
		if strings.Contains(value, "forwardingRules/ssc-auto-fr-") {
			tokens := strings.Split(value, "/")
			for _, token := range tokens {
				if strings.HasPrefix(token, "ssc-auto-fr-") {
					replacements.ReplaceStringValue(token, "ssc-auto-fr-abcde")
				}
			}
		}
	})
}
