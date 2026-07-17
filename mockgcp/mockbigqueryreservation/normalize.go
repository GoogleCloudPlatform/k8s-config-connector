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

package mockbigqueryreservation

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// CapacityCommitment API Normalization
	replacements.ReplacePath(".commitmentStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".commitmentEndTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".capacityCommitments[].commitmentStartTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".capacityCommitments[].commitmentEndTime", mockgcpregistry.PlaceholderTimestamp)

	// CapacityCommitment KRM Normalization
	replacements.ReplacePath(".status.commitmentStartTime", mockgcpregistry.PlaceholderTime)
	replacements.ReplacePath(".status.commitmentEndTime", mockgcpregistry.PlaceholderTime)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
