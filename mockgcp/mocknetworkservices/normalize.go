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

package mocknetworkservices

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "networkservices.googleapis.com") {
		return
	}

	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".endTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".lbRouteExtensions[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".lbRouteExtensions[].updateTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".wasmPlugins[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".wasmPlugins[].updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".versions.*.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".versions.*.updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".wasmPluginVersions[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".wasmPluginVersions[].updateTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.requestedCancellation", nil)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// Only apply this logic if the request is for the networkservices API.
	if !strings.Contains(event.URL(), "networkservices.googleapis.com") {
		return
	}

	// Normalize compute reference URLs in responses to match real GCP behavior.
	event.VisitResponseStringValues(func(path string, value string) {
		if strings.HasPrefix(value, "https://compute.googleapis.com/compute/v1/") {
			newValue := strings.Replace(value, "https://compute.googleapis.com/compute/v1/", "https://www.googleapis.com/compute/v1/", 1)
			replacements.ReplaceStringValue(value, newValue)
		}
	})
}
