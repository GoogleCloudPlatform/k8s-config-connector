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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath("pscAttachmentDetails[].serviceAttachment", "projects/1234567890/regions/us-central1/serviceAttachments/sampleServiceAttachment")
	// 	visitor.ReplacePath(".response.pscConnections[].address", "10.11.12.13")
	// 	visitor.ReplacePath(".discoveryEndpoints[].address", "10.11.12.13")
	// 	visitor.ReplacePath(".response.discoveryEndpoints[].address", "10.11.12.13")
	// }

	// // Specific to Memorystore
	// {
	// 	visitor.ReplacePath(".response.pscAttachmentDetails[].serviceAttachment", "projects/1234567890/regions/us-central1/serviceAttachments/sampleServiceAttachment")
	// 	visitor.ReplacePath(".response.endpoints[].connections[].pscAutoConnection.forwardingRule", "https://www.googleapis.com/compute/v1/projects/project-tp/regions/us-central1/forwardingRules/sca-auto-fr-sample")
	// 	visitor.ReplacePath(".response.endpoints[].connections[].pscAutoConnection.ipAddress", "10.11.12.13")
	// 	visitor.ReplacePath(".response.endpoints[].connections[].pscAutoConnection.serviceAttachment", "projects/1234567890/regions/us-central1/serviceAttachments/sampleServiceAttachment")
	// 	visitor.ReplacePath(".response.satisfiesPzi", "true")
	// }

	// // Specific to ServiceConnectionPolicy
	// {
	// 	visitor.ReplacePath(".response.pscConnections[].consumerAddress", "https://www.googleapis.com/compute/v1/projects/${projectId}/regions/us-central1/addresses/sca-auto-addr-f4949b71-c18b-49b1-ab5c-991b7badec5e")
	// 	visitor.ReplacePath(".response.pscConnections[].consumerForwardingRule", "https://www.googleapis.com/compute/v1/projects/${projectId}/regions/us-central1/forwardingRules/sca-auto-fr-8f3463ea-1df2-4c48-bcfc-9dc8eb6e8b01")
	// }
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// No-op for now
}
