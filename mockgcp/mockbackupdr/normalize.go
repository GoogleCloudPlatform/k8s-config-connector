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

package mockbackupdr

import (
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
}

var reBackupVaultSA = regexp.MustCompile(`vault-\d+-\d+@gcp-sa-backupdr-pr\.iam\.gserviceaccount\.com`)
var reRevisionName = regexp.MustCompile(`/revisions/([a-f0-9]{8})`)

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "backupdr.googleapis.com") {
		return
	}
	if !strings.Contains(event.URL(), "/projects/") {
		return
	}

	for _, match := range reRevisionName.FindAllString(event.URL(), -1) {
		replacements.ReplaceStringValue(match, "/revisions/00000000")
	}

	replaceFunc := func(path string, value string) {
		for _, match := range reBackupVaultSA.FindAllString(value, -1) {
			replacements.ReplaceStringValue(match, "vault-${projectNumber}-12345@gcp-sa-backupdr-pr.iam.gserviceaccount.com")
		}
		for _, match := range reRevisionName.FindAllString(value, -1) {
			replacements.ReplaceStringValue(match, "/revisions/00000000")
		}
		if strings.Contains(strings.ToLower(path), "createtime") || strings.Contains(strings.ToLower(path), "updatetime") {
			// standard normalization for timestamps in paths like instanceCreateTime
			replacements.ReplaceStringValue(value, "2024-04-01T12:34:56.123456Z")
		}
	}

	event.VisitRequestStringValues(replaceFunc)
	event.VisitResponseStringValues(replaceFunc)
}
