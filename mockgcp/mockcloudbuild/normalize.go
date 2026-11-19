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

package mockcloudbuild

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "cloudbuild.googleapis.com") && !strings.Contains(url, "cloudbuild") {
		return
	}

	replacements.ReplacePath(".startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".finishTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.finishTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".steps[].pullTiming.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".steps[].pullTiming.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".steps[].timing.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".steps[].timing.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.steps[].pullTiming.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.steps[].pullTiming.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.steps[].timing.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.steps[].timing.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".timing.BUILD.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".timing.BUILD.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".timing.FETCHSOURCE.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".timing.FETCHSOURCE.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".timing.PUSH.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".timing.PUSH.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.timing.BUILD.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.timing.BUILD.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.timing.FETCHSOURCE.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.timing.FETCHSOURCE.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.timing.PUSH.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.build.timing.PUSH.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.finishTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.steps[].pullTiming.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.steps[].pullTiming.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.steps[].timing.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.steps[].timing.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.timing.BUILD.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.timing.BUILD.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.timing.FETCHSOURCE.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.timing.FETCHSOURCE.endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.timing.PUSH.startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.timing.PUSH.endTime", mockgcpregistry.PlaceholderTimestamp)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "cloudbuild.googleapis.com") {
		return
	}
}
