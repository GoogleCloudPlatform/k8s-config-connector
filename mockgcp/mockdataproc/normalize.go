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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Cluster
	replacements.ReplacePath(".status.stateStartTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.status.stateStartTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".clusters[].status.stateStartTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".metadata.status.stateStartTime", "2024-04-01T12:34:56.123456Z")

	replacements.ReplacePath(".statusHistory[].stateStartTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.statusHistory[].stateStartTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".clusters[].statusHistory[].stateStartTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".metadata.statusHistory[].stateStartTime", "2024-04-01T12:34:56.123456Z")

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
	// for _, metric := range []string{
	// 	"yarn-memory-mb-total",
	// 	"yarn-memory-mb-available",
	// 	"dfs-capacity-total",
	// 	"dfs-capacity-used"} {

	// 	replacements.RemovePath(".metrics.yarnMetrics." + metric)
	// 	replacements.RemovePath(".response.metrics.yarnMetrics." + metric)
	// 	replacements.RemovePath(".clusters[].metrics.yarnMetrics." + metric)
	// }
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
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
