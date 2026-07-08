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

package mockdataproc

import (
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Cluster
	replacements.ReplacePath(".status.stateStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.status.stateStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".clusters[].status.stateStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.status.stateStartTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".statusHistory[].stateStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.statusHistory[].stateStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".clusters[].statusHistory[].stateStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.statusHistory[].stateStartTime", mockgcpregistry.PlaceholderTimestamp)

	// Job
	replacements.ReplacePath(".statusHistory[].stateStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".status.stateStartTime", mockgcpregistry.PlaceholderTimestamp)

	// metrics are volatile and more "data plane"
	for _, metric := range []string{
		"dfs-capacity-present",
		"dfs-capacity-remaining",
		"dfs-capacity-total",
		"dfs-capacity-used"} {

		replacements.ReplacePath(".metrics.hdfsMetrics."+metric, 12345)
		replacements.ReplacePath(".response.metrics.hdfsMetrics."+metric, 12345)
		replacements.ReplacePath(".clusters[].metrics.hdfsMetrics."+metric, 12345)
	}
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "dataproc.googleapis.com") {
		return
	}

	// We want to normalize the random session UUID to "00000000-0000-0000-0000-000000000001"
	// Let's find any UUID-like string in the response and register a replacement for it!
	uuidRegex := regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)
	event.VisitResponseStringValues(func(path string, value string) {
		// Cluster UUID is already handled and normalized to ${dataStoreClusterUUID}, so don't overwrite it if it's clusterUuid
		if path == ".clusterUuid" || path == ".response.clusterUuid" {
			return
		}
		for _, match := range uuidRegex.FindAllString(value, -1) {
			replacements.ReplaceStringValue(match, "00000000-0000-0000-0000-000000000001")
		}
	})

	event.VisitResponseStringValues(func(path string, value string) {
		switch path {
		case ".config.configBucket":
			replacements.ReplaceStringValue(value, "${dataStoreConfigBucketPath}")
		case ".config.tempBucket":
			replacements.ReplaceStringValue(value, "${dataStoreTempBucketPath}")
		case ".clusterUuid":
			replacements.ReplaceStringValue(value, "${dataStoreClusterUUID}")
		}
	})
}
