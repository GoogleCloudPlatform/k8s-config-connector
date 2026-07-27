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
	"fmt"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "redis.googleapis.com") {
		return
	}
	replacements.ReplacePath(".pscConnections[].address", "10.11.12.13")
	replacements.ReplacePath(".clusterEndpoints[].connections[].pscAutoConnection.address", "10.11.12.13")
	replacements.ReplacePath(".clusterEndpoints[].connections[].pscAutoConnection.pscConnectionId", "${pscConnectionID}")
	replacements.ReplacePath(".clusterEndpoints[].connections[].pscConnection.address", "10.11.12.13")
	replacements.ReplacePath(".clusterEndpoints[].connections[].pscConnection.pscConnectionId", "${pscConnectionID}")
	replacements.ReplacePath(".response.pscConnections[].address", "10.11.12.13")
	replacements.ReplacePath(".response.clusterEndpoints[].connections[].pscAutoConnection.address", "10.11.12.13")
	replacements.ReplacePath(".response.clusterEndpoints[].connections[].pscAutoConnection.pscConnectionId", "${pscConnectionID}")
	replacements.ReplacePath(".response.clusterEndpoints[].connections[].pscConnection.address", "10.11.12.13")
	replacements.ReplacePath(".response.clusterEndpoints[].connections[].pscConnection.pscConnectionId", "${pscConnectionID}")
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
	replacements.ReplacePath(".status.observedState.crossClusterReplicationConfig.membership.secondaryClusters[].uid", "0123456789abcdef")
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
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if strings.Contains(event.URL(), "https://redis.googleapis.com") || strings.Contains(event.URL(), "https://compute.googleapis.com") {
		normalizeSA := func(path string, value string) {
			re := regexp.MustCompile(`projects/[^/]+/regions/([^/]+)/serviceAttachments/gcp-memorystore-auto-([a-f0-9]+)-psc-sa(-2)?`)
			if re.MatchString(value) {
				match := re.FindStringSubmatch(value)
				normalized := fmt.Sprintf("projects/${projectNumber}/regions/%s/serviceAttachments/gcp-memorystore-auto-abcdef0123456789-psc-sa%s", match[1], match[3])
				replacements.ReplaceStringValue(value, normalized)
			}
		}
		event.VisitRequestStringValues(normalizeSA)
		event.VisitResponseStringValues(normalizeSA)
	}

	if strings.Contains(event.URL(), "https://redis.googleapis.com") {
		event.VisitResponseStringValues(func(path string, value string) {
			if strings.HasSuffix(path, ".forwardingRule") {
				tokens := strings.Split(value, "/")
				if len(tokens) == 11 && tokens[9] == "forwardingRules" {
					// e.g. https://www.googleapis.com/compute/v1/projects/${projectId}/regions/us-central1/forwardingRules/sca-auto-fr-bed6...
					if strings.HasSuffix(tokens[10], "-2") {
						replacements.ReplaceStringValue(tokens[10], "${forwardingRuleID}-2")
					} else {
						replacements.ReplaceStringValue(tokens[10], "${forwardingRuleID}")
					}
				}
			}
		})
	}
}
