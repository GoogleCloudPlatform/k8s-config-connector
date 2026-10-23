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

package mockkms

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "cloudkms.googleapis.com") {
		return
	}

	replacements.ReplacePath(".cryptoKeyVersions[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".cryptoKeyVersions[].generateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".cryptoKeyVersions[].destroyTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".destroyTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".expireTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".generateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".importJobs[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".importJobs[].expireTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".importJobs[].generateTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.TransformObject(".publicKey", func(m map[string]any) {
		delete(m, "data")
	})
	replacements.TransformObject(".importJobs[].publicKey", func(m map[string]any) {
		delete(m, "data")
	})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "cloudkms.googleapis.com") {
		return
	}
	if entry, ok := event.(*test.LogEntry); ok {
		entry.Request.RemoveHeader("X-Allowed-Locations")
	}
}
