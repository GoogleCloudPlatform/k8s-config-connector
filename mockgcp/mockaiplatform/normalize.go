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

package mockaiplatform

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/google/uuid"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	visitor := func(path string, value string) {
		if strings.Contains(value, "/dataLabelingJobs/") {
			tokens := strings.Split(value, "/")
			if len(tokens) == 6 && tokens[4] == "dataLabelingJobs" {
				// Only replace actual server-generated UUIDs, not user-specified GVK names
				if _, err := uuid.Parse(tokens[5]); err == nil {
					replacements.ReplaceStringValue(tokens[5], "${dataLabelingJobID}")
				}
			}
		}
	}
	event.VisitRequestStringValues(visitor)
	event.VisitResponseStringValues(visitor)
}
