// Copyright 2022 Google LLC
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

package mockprivateca

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

const NormalizedTimestamp = "2024-04-01T12:34:56.123456Z"

func (s *MockService) ConfigureVisitor(url string, visitor mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "privateca.googleapis.com") {
		return
	}

	visitor.ReplacePath(".createTime", NormalizedTimestamp)
	visitor.ReplacePath(".updateTime", NormalizedTimestamp)
	visitor.ReplacePath(".deleteTime", NormalizedTimestamp)
	visitor.ReplacePath(".expireTime", NormalizedTimestamp)
	visitor.ReplacePath(".endTime", NormalizedTimestamp)

	visitor.ReplacePath(".response.createTime", NormalizedTimestamp)
	visitor.ReplacePath(".response.updateTime", NormalizedTimestamp)
	visitor.ReplacePath(".response.deleteTime", NormalizedTimestamp)
	visitor.ReplacePath(".response.expireTime", NormalizedTimestamp)

	visitor.ReplacePath(".metadata.createTime", NormalizedTimestamp)
	visitor.ReplacePath(".metadata.endTime", NormalizedTimestamp)
	visitor.ReplacePath(".metadata.requestedCancellation", false)
	visitor.ReplacePath(".done", true)

	visitor.ReplacePath(".caCertificateDescriptions[].subjectDescription.notBeforeTime", NormalizedTimestamp)
	visitor.ReplacePath(".caCertificateDescriptions[].subjectDescription.notAfterTime", NormalizedTimestamp)

	visitor.ReplacePath(".response.caCertificateDescriptions[].subjectDescription.notBeforeTime", NormalizedTimestamp)
	visitor.ReplacePath(".response.caCertificateDescriptions[].subjectDescription.notAfterTime", NormalizedTimestamp)

	visitor.ReplacePath(".status.caCertificateDescriptions[].subjectDescription.notBeforeTime", NormalizedTimestamp)
	visitor.ReplacePath(".status.caCertificateDescriptions[].subjectDescription.notAfterTime", NormalizedTimestamp)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
