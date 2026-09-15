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

package mockassuredworkloads

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "assuredworkloads.googleapis.com") {
		return
	}

	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)

	// Since complianceStatus, partnerPermissions, resourceMonitoringEnabled, violationNotificationsEnabled
	// are not present in the v1 protobuf, they are not returned by MockGCP. We strip them from the real
	// GCP logs to align with MockGCP.
	replacements.RemovePath(".complianceStatus")
	replacements.RemovePath(".response.complianceStatus")

	replacements.RemovePath(".partnerPermissions")
	replacements.RemovePath(".response.partnerPermissions")

	replacements.RemovePath(".resourceMonitoringEnabled")
	replacements.RemovePath(".response.resourceMonitoringEnabled")

	replacements.RemovePath(".violationNotificationsEnabled")
	replacements.RemovePath(".response.violationNotificationsEnabled")

	// billingAccount is not returned by Real GCP GET/PATCH responses, so we strip it to align.
	replacements.RemovePath(".billingAccount")
	replacements.RemovePath(".response.billingAccount")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
