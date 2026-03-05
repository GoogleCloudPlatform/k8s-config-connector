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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

const NormalizedTimestamp = "2024-04-01T12:34:56.123456Z"

func (s *MockService) ConfigureVisitor(url string, visitor mockgcpregistry.NormalizingVisitor) {
	visitor.ReplacePath(".createTime", NormalizedTimestamp)
	visitor.ReplacePath(".updateTime", NormalizedTimestamp)
	visitor.ReplacePath(".response.createTime", NormalizedTimestamp)
	visitor.ReplacePath(".response.updateTime", NormalizedTimestamp)
	visitor.ReplacePath(".metadata.createTime", NormalizedTimestamp)
	visitor.ReplacePath(".certificateIssuanceConfigs[].createTime", NormalizedTimestamp)
	visitor.ReplacePath(".certificateIssuanceConfigs[].updateTime", NormalizedTimestamp)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, visitor mockgcpregistry.NormalizingVisitor) {
}
