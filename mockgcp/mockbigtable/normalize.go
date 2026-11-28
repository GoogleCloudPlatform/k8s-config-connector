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

package mockbigtable

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".metadata.requestTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.finishTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)

	// Instances
	replacements.ReplacePath(".instances[].createTime", mockgcpregistry.PlaceholderTimestamp)

	// Backups
	replacements.ReplacePath(".startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".expireTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".status.observedState.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".status.observedState.endTime", mockgcpregistry.PlaceholderTimestamp)
	if strings.Contains(url, "google.bigtable.admin.") {
		// Backup expiry/tiering times are relative to "now" when running against real GCP
		// (Bigtable only accepts an expireTime up to 90 days out), so normalize them everywhere they appear.
		replacements.ReplacePath(".backup.expireTime", mockgcpregistry.PlaceholderTimestamp)
		replacements.ReplacePath(".backup.hotToStandardTime", mockgcpregistry.PlaceholderTimestamp)
		replacements.ReplacePath(".hotToStandardTime", mockgcpregistry.PlaceholderTimestamp)
		replacements.ReplacePath(".response.hotToStandardTime", mockgcpregistry.PlaceholderTimestamp)
		replacements.ReplacePath(".spec.hotToStandardTime", mockgcpregistry.PlaceholderTimestamp)
	}

	// Materialized Views
	replacements.ReplacePath(".materializedViews[].etag", "abcdef0123A=")
	if strings.Contains(url, "/materializedViews") {
		replacements.RemovePath(".metadata.requestTime")
		replacements.RemovePath(".metadata.finishTime")
	}

	// Schema Bundles
	replacements.ReplacePath(".schemaBundles[].etag", "abcdef0123A=")
	replacements.ReplacePath(".etag", "abcdef0123A=")

	// Memory Layers
	replacements.ReplacePath(".memoryLayer.etag", "abcdef0123A=")
	replacements.ReplacePath(".metadata.originalRequest.memoryLayer.etag", "abcdef0123A=")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
