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

package mockrun

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".metadata.terminalCondition.lastTransitionTime", "2025-01-01T00:00:00Z")
	replacements.ReplacePath(".terminalCondition.lastTransitionTime", "2025-01-01T00:00:00Z")
	replacements.ReplacePath(".response.terminalCondition.lastTransitionTime", "2025-01-01T00:00:00Z")
	replacements.ReplacePath(".metadata.etag", "abcdef0123A=")
	replacements.ReplacePath(".metadata.annotations.run.googleapis.com/operation-id", "00000000-0000-0000-0000-000000000000")
	replacements.ReplacePath(".metadata.annotations.run.googleapis.com/client-version", "0.0.0")
	replacements.ReplacePath(".metadata.annotations.run.googleapis.com/creator", "test@google.com")
	replacements.ReplacePath(".metadata.annotations.run.googleapis.com/lastModifier", "test@google.com")
	replacements.ReplacePath(".response.metadata.annotations.run.googleapis.com/operation-id", "00000000-0000-0000-0000-000000000000")
	replacements.ReplacePath(".response.metadata.annotations.run.googleapis.com/client-version", "0.0.0")
	replacements.ReplacePath(".response.metadata.annotations.run.googleapis.com/client-creator", "test@google.com")
	replacements.ReplacePath(".response.metadata.annotations.run.googleapis.com/client-lastModifier", "test@google.com")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
