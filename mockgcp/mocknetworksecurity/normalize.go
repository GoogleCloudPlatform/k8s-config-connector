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

package mocknetworksecurity

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.RemovePath(".error.details")
	replacements.ReplacePath(".pairingKey.expireTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.pairingKey.expireTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)

	if strings.Contains(url, "gatewaySecurityPolicies") && strings.Contains(url, "rules") {
		// applicationMatcher and tlsInspectionEnabled are zero-values ("" and false) in real GCP
		// which proto3 JSON marshaling omits.
		cleanRule := func(m map[string]any) {
			if v, ok := m["applicationMatcher"].(string); ok && v == "" {
				delete(m, "applicationMatcher")
			}
			if v, ok := m["tlsInspectionEnabled"].(bool); ok && !v {
				delete(m, "tlsInspectionEnabled")
			}
		}
		replacements.TransformObject("", cleanRule)
		replacements.TransformObject(".response", cleanRule)
	}
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
