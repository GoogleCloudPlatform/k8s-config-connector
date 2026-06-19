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

package mockredis

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "redis.googleapis.com") {
		return
	}
	replacements.ReplacePath(".pscConnections[].address", "10.11.12.13")
	replacements.ReplacePath(".response.pscConnections[].address", "10.11.12.13")
	replacements.ReplacePath(".discoveryEndpoints[].address", "10.11.12.13")
	replacements.ReplacePath(".response.discoveryEndpoints[].address", "10.11.12.13")
	replacements.ReplacePath(".encryptionInfo.lastUpdateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.encryptionInfo.lastUpdateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.uid", "0123456789abcdef")
	replacements.ReplacePath(".crossClusterReplicationConfig.updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.crossClusterReplicationConfig.updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".crossClusterReplicationConfig.membership.primaryCluster.uid", "0123456789abcdef")
	replacements.ReplacePath(".crossClusterReplicationConfig.membership.secondaryClusters[].uid", "0123456789abcdef")
	replacements.ReplacePath(".crossClusterReplicationConfig.primaryCluster.uid", "0123456789abcdef")
	replacements.ReplacePath(".crossClusterReplicationConfig.secondaryClusters[].uid", "0123456789abcdef")
	replacements.ReplacePath(".response.crossClusterReplicationConfig.membership.primaryCluster.uid", "0123456789abcdef")
	replacements.ReplacePath(".response.crossClusterReplicationConfig.membership.secondaryClusters[].uid", "0123456789abcdef")
	replacements.ReplacePath(".response.crossClusterReplicationConfig.primaryCluster.uid", "0123456789abcdef")
	replacements.ReplacePath(".response.crossClusterReplicationConfig.secondaryClusters[].uid", "0123456789abcdef")

	replacements.ReplacePath(".status.observedState.pscConnections[].pscConnectionID", "${pscConnectionID}")
	replacements.ReplacePath(".status.observedState.pscConnections[].address", "10.11.12.13")
	replacements.ReplacePath(".status.observedState.discoveryEndpoints[].address", "10.11.12.13")
	replacements.ReplacePath(".status.observedState.crossClusterReplicationConfig.updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".status.observedState.crossClusterReplicationConfig.membership.primaryCluster.uid", "0123456789abcdef")
	replacements.ReplacePath(".status.observedState.crossClusterReplicationConfig.primaryCluster.uid", "0123456789abcdef")
	replacements.ReplacePath(".status.observedState.crossClusterReplicationConfig.secondaryClusters[].uid", "0123456789abcdef")
	replacements.ReplacePath(".status.observedState.crossClusterReplicationConfig.secondaryClusters[].uid", "0123456789abcdef")

	// Standardize zone locations to us-central1-a
	replacements.ReplacePath(".currentLocationId", "us-central1-a")
	replacements.ReplacePath(".response.currentLocationId", "us-central1-a")
	replacements.ReplacePath(".locationId", "us-central1-a")
	replacements.ReplacePath(".response.locationId", "us-central1-a")
	replacements.ReplacePath(".nodes[].zone", "us-central1-a")
	replacements.ReplacePath(".response.nodes[].zone", "us-central1-a")

	// Remove maintenance fields since they are volatile and not supported by the mock
	replacements.RemovePath(".availableMaintenanceVersions")
	replacements.RemovePath(".response.availableMaintenanceVersions")
	replacements.RemovePath(".maintenanceVersion")
	replacements.RemovePath(".response.maintenanceVersion")

	// Remove LRO 'done' field to align mock and real GCP formats
	replacements.RemovePath(".done")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// No-op for now
}
