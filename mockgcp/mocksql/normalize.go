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

package mocksql

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// SQLUser
	replacements.ReplacePath(".items[].passwordPolicy.status.passwordExpirationTime", "2025-06-19T01:02:03Z")

	// SQLInstance
	replacements.ReplacePath(".etag", "abcdef0123A=")
	replacements.ReplacePath(".createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".ipAddresses[].ipAddress", "10.1.2.3")
	replacements.ReplacePath(".serverCaCert.createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".serverCaCert.expirationTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".serverCaCert.sha1Fingerprint", "12345678")
	replacements.ReplacePath(".serverCaCert.commonName", "common-name")
	replacements.ReplacePath(".settings.settingsVersion", "123")

	// SQL Operation
	replacements.ReplacePath(".insertTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".startTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".endTime", "2024-04-01T12:34:56.123456Z")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	event.VisitResponseStringValues(func(path string, value string) {
		if path == ".serviceAccountEmailAddress" {
			if strings.HasSuffix(value, "@gcp-sa-cloud-sql.iam.gserviceaccount.com") {
				replacements.ReplaceStringValue(value, "p${projectNumber}-abcdef@gcp-sa-cloud-sql.iam.gserviceaccount.com")
			}
		}
	})
}
