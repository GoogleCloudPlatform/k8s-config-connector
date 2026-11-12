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

package mockalloydb

import (
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var (
	pscDNSNameRegex            = regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\.([a-z0-9\-]+)\.alloydb-psc\.goog\.`)
	pscAutoDNSNameRegex        = regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\.([a-z0-9\-]+)\.alloydb\.goog\.`)
	serviceAttachmentLinkRegex = regexp.MustCompile(`projects/[^/]+/regions/([^/]+)/serviceAttachments/[^/]+`)
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func normalizePSCDNSName(s string) string {
	return pscDNSNameRegex.ReplaceAllString(s, "00000000-0000-0000-0000-000000000001.00000000-0000-0000-0000-000000000001.$1.alloydb-psc.goog.")
}

func normalizePSCAutoDNSName(s string) string {
	return pscAutoDNSNameRegex.ReplaceAllString(s, "00000000-0000-0000-0000-000000000001.00000000-0000-0000-0000-000000000001.$1.alloydb.goog.")
}

func normalizeServiceAttachmentLink(s string) string {
	return serviceAttachmentLinkRegex.ReplaceAllString(s, "projects/tenant-project/regions/$1/serviceAttachments/alloydb-psc-sa")
}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "alloydb.googleapis.com") {
		return
	}
	replacements.ReplacePath(".continuousBackupInfo.enabledTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.continuousBackupInfo.enabledTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".ipAddress", "10.1.2.3")
	replacements.ReplacePath(".response.ipAddress", "10.1.2.3")
	replacements.ReplacePath(".primary.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".primary.generateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".backupSource.backupUid", "1a1a1a-222b-3cc3-d444-e555ee555555")
	replacements.ReplacePath(".response.backupSource.backupUid", "1a1a1a-222b-3cc3-d444-e555ee555555")
	replacements.ReplacePath(".serviceAccountEmail", "c-${projectNumber}-12345678@gcp-sa-alloydb.iam.gserviceaccount.com")
	replacements.ReplacePath(".response.serviceAccountEmail", "c-${projectNumber}-12345678@gcp-sa-alloydb.iam.gserviceaccount.com")
	replacements.ReplacePath(".pscConfig.serviceOwnedProjectNumber", "123456789")
	replacements.ReplacePath(".response.pscConfig.serviceOwnedProjectNumber", "123456789")
	replacements.TransformString(".pscInstanceConfig.pscDnsName", normalizePSCDNSName)
	replacements.TransformString(".response.pscInstanceConfig.pscDnsName", normalizePSCDNSName)
	replacements.TransformString(".pscInstanceInfo.pscAutoDnsNames[]", normalizePSCAutoDNSName)
	replacements.TransformString(".response.pscInstanceInfo.pscAutoDnsNames[]", normalizePSCAutoDNSName)
	replacements.ReplacePath(".pscInstanceConfig.pscAutoConnections[].ipAddress", "10.0.0.2")
	replacements.ReplacePath(".response.pscInstanceConfig.pscAutoConnections[].ipAddress", "10.0.0.2")
	replacements.TransformString(".pscInstanceConfig.serviceAttachmentLink", normalizeServiceAttachmentLink)
	replacements.TransformString(".response.pscInstanceConfig.serviceAttachmentLink", normalizeServiceAttachmentLink)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
