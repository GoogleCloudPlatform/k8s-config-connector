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

package mockvmwareengine

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".uid", "111111111111111111111")
	replacements.ReplacePath(".response.uid", "111111111111111111111")

	replacements.ReplacePath(".etag", "abcdef0123A=")
	replacements.ReplacePath(".response.etag", "abcdef0123A=")

	// PrivateConnection peeringId
	replacements.ReplacePath(".peeringId", "peering-55c50401-d392-4c50-94af-d4a1242f70ac")
	replacements.ReplacePath(".response.peeringId", "peering-55c50401-d392-4c50-94af-d4a1242f70ac")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
