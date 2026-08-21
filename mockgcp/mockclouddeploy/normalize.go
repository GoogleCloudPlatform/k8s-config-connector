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

package mockclouddeploy

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "clouddeploy.googleapis.com") {
		return
	}
	// Use standard placeholders to avoid conflicts with global normalizers
	const (
		PlaceholderUID  = "111111111111111111111"
		PlaceholderEtag = "abcdef0123A="
	)

	// Note: .uid and .etag are already handled globally at the root.
	// We only need to specify them for nested paths or if we use non-standard values.

	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)

	// Array normalization for ListTargets
	replacements.ReplacePath(".targets[].uid", PlaceholderUID)
	replacements.ReplacePath(".targets[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".targets[].updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".targets[].etag", PlaceholderEtag)

	// Array normalization for ListAutomations
	replacements.ReplacePath(".automations[].uid", PlaceholderUID)
	replacements.ReplacePath(".automations[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".automations[].updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".automations[].etag", PlaceholderEtag)

	replacements.ReplacePath(".rules[].promoteReleaseRule.condition", map[string]interface{}{
		"targetsPresentCondition": make(map[string]interface{}),
	})

	// Some responses wrap the object in a "response" field (e.g. LROs or some List responses in the harness)
	replacements.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.updateTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".response.targets[].uid", PlaceholderUID)
	replacements.ReplacePath(".response.targets[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.targets[].updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.targets[].etag", PlaceholderEtag)

	replacements.ReplacePath(".response.automations[].uid", PlaceholderUID)
	replacements.ReplacePath(".response.automations[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.automations[].updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.automations[].etag", PlaceholderEtag)

	replacements.ReplacePath(".response.rules[].promoteReleaseRule.condition", map[string]interface{}{
		"targetsPresentCondition": make(map[string]interface{}),
	})

	// LRO metadata
	replacements.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.requestedCancellation", false)

	// LRO root
	replacements.ReplacePath(".done", true)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
