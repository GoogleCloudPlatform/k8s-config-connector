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

package mockconnectors

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "connectors.googleapis.com") {
		return
	}

	replacements.ReplacePath("createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath("updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)

	// Since billingConfig, connectionRevision, connectorVersionInfraConfig, connectorVersionLaunchStage,
	// host, subscriptionType, and tlsServiceDirectory are not present in the google.cloud.connectors.v1 protobuf,
	// they are not modeled in MockGCP. We strip them from the HTTP logs to align with MockGCP.
	// TODO: Add these fields to mockgcp proto once supported in googleapis.
	replacements.RemovePath(".billingConfig")
	replacements.RemovePath(".response.billingConfig")

	replacements.RemovePath(".connectionRevision")
	replacements.RemovePath(".response.connectionRevision")

	replacements.RemovePath(".connectorVersionInfraConfig")
	replacements.RemovePath(".response.connectorVersionInfraConfig")

	replacements.RemovePath(".connectorVersionLaunchStage")
	replacements.RemovePath(".response.connectorVersionLaunchStage")

	replacements.RemovePath(".host")
	replacements.RemovePath(".response.host")

	replacements.RemovePath(".subscriptionType")
	replacements.RemovePath(".response.subscriptionType")

	replacements.RemovePath(".tlsServiceDirectory")
	replacements.RemovePath(".response.tlsServiceDirectory")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
