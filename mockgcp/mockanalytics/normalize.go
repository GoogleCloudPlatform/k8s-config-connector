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

package mockanalytics

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath("accountTicketId", "ASDFGHJKL123456")
	replacements.ReplacePath("createTime", "2025-10-13T10:54:55.445887715Z")
	replacements.ReplacePath("updateTime", "2025-10-13T10:58:19.612207875Z")
	replacements.ReplacePath(".accounts[].createTime", "2025-10-13T10:54:55.445887715Z")
	replacements.ReplacePath(".accounts[].updateTime", "2025-10-13T10:58:19.612207875Z")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
