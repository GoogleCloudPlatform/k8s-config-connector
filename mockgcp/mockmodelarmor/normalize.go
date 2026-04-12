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

package mockmodelarmor

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

const PlaceholderTimestamp = "2024-04-01T12:34:56.123456Z"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
        if !strings.Contains(url, "modelarmor.googleapis.com") && !strings.Contains(url, "modelarmor.us-central1.rep.googleapis.com") {
                return
        }
        replacements.ReplacePath(".createTime", PlaceholderTimestamp)
        replacements.ReplacePath(".updateTime", PlaceholderTimestamp)
        replacements.ReplacePath(".templates[].createTime", PlaceholderTimestamp)
        replacements.ReplacePath(".templates[].updateTime", PlaceholderTimestamp)
}
func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// Only apply normalization if the request is for this service
	if !strings.Contains(event.URL(), "modelarmor.googleapis.com") && !strings.Contains(event.URL(), "modelarmor.us-central1.rep.googleapis.com") {
		return
	}
}
