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

package mockdeveloperconnect

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "developerconnect.googleapis.com") {
		return
	}

	replacements.ReplacePath("createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath("updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath("endTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.updateTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.RemovePath(".response.labels")

	replacements.TransformObject(".error", func(m map[string]any) {
		delete(m, "errors")
	})
	replacements.TransformObject("", func(m map[string]any) {
		if val, found := m["done"]; found && val == false {
			delete(m, "done")
		}
	})
	replacements.TransformObject(".metadata", func(m map[string]any) {
		if val, found := m["requestedCancellation"]; found && val == false {
			delete(m, "requestedCancellation")
		}
	})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
