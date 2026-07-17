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

package mockcertificatemanager

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, visitor mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "certificatemanager.googleapis.com") {
		return
	}
	visitor.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".response.updateTime", mockgcpregistry.PlaceholderTimestamp)

	visitor.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".done", true)

	visitor.ReplacePath(".certificateIssuanceConfigs[].createTime", mockgcpregistry.PlaceholderTimestamp)
	visitor.ReplacePath(".certificateIssuanceConfigs[].updateTime", mockgcpregistry.PlaceholderTimestamp)

	normalizeType := func(m map[string]any) {
		if typeVal, found := m["type"]; found {
			switch v := typeVal.(type) {
			case float64:
				switch int(v) {
				case 1:
					m["type"] = "FIXED_RECORD"
				case 2:
					m["type"] = "PER_PROJECT_RECORD"
				}
			case int:
				switch v {
				case 1:
					m["type"] = "FIXED_RECORD"
				case 2:
					m["type"] = "PER_PROJECT_RECORD"
				}
			}
		}
	}
	visitor.TransformObject("", normalizeType)
	visitor.TransformObject(".response", normalizeType)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, visitor mockgcpregistry.NormalizingVisitor) {
}
