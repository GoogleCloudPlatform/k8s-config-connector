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

package mockorgpolicy

import (
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

// projectNumericPolicyPrefix matches Org Policy policy resource names that use a
// numeric project id segment. Real and mock responses may differ (project id vs
// project number); normalize to ${projectId} for golden HTTP comparisons.
var projectNumericPolicyPrefix = regexp.MustCompile(`projects/[0-9]+/policies/`)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// policy
	replacements.ReplacePath(".spec.updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".spec.etag", "abcdef0123A=")
	replacements.ReplacePath(".dryRunSpec.updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".dryRunSpec.etag", "abcdef0123A=")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "orgpolicy.googleapis.com") {
		return
	}
	if entry, ok := event.(*test.LogEntry); ok {
		entry.Request.RemoveHeader("X-Goog-User-Project")
	}
	registerPolicyNameProjectCanonicalization(event.URL(), replacements)
	event.VisitRequestStringValues(func(_ string, value string) {
		registerPolicyNameProjectCanonicalization(value, replacements)
	})
	event.VisitResponseStringValues(func(_ string, value string) {
		registerPolicyNameProjectCanonicalization(value, replacements)
	})
}

func registerPolicyNameProjectCanonicalization(s string, replacements mockgcpregistry.NormalizingVisitor) {
	for _, m := range projectNumericPolicyPrefix.FindAllString(s, -1) {
		replacements.ReplaceStringValue(m, "projects/${projectId}/policies/")
	}
}
