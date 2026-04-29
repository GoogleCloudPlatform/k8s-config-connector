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

package mockmemorystore

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "memorystore.googleapis.com") {
		return
	}
	replacements.ReplacePath(".backupFiles[].createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".crossInstanceReplicationConfig.updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".expireTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".maintenancePolicy.createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".maintenancePolicy.updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.backupFiles[].createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.crossInstanceReplicationConfig.updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.expireTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.maintenancePolicy.createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.maintenancePolicy.updateTime", "2024-04-01T12:34:56.123456Z")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// No-op for now
}
