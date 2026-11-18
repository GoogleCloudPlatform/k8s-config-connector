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

package mockworkflowexecution

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func isWorkflowExecutionURL(u string) bool {
	return strings.Contains(u, "workflowexecutions.googleapis.com") ||
		strings.Contains(u, "google.cloud.workflows.executions.")
}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !isWorkflowExecutionURL(url) {
		return
	}

	// Execution
	replacements.ReplacePath(".executions[].startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".executions[].endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".status.observedState.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".status.observedState.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".workflowRevisionId", "000001-609")
	replacements.ReplacePath(".status.observedState.workflowRevisionID", "000001-609")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isWorkflowExecutionURL(event.URL()) {
		return
	}
	event.VisitResponseStringValues(func(path string, value string) {
		if path == ".name" || strings.HasSuffix(path, ".name") {
			tokens := strings.Split(value, "/")
			if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workflows" && tokens[6] == "executions" && tokens[7] != "" {
				replacements.ReplaceStringValue("executions/"+tokens[7], "executions/123456789")
			}
		}
	})
}
