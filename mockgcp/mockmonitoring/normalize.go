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

package mockmonitoring

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

const (
	emailPlaceholder = "user@example.com"
	timePlaceholder  = "2024-04-01T12:34:56.123456Z"
)

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".creationRecord.mutateTime", timePlaceholder)
	replacements.ReplacePath(".notificationChannels[].creationRecord.mutateTime", timePlaceholder)
	replacements.ReplacePath(".alertPolicies[].creationRecord.mutateTime", timePlaceholder)

	replacements.ReplacePath(".mutationRecord.mutateTime", timePlaceholder)
	replacements.ReplacePath(".notificationChannels[].mutationRecord.mutateTime", timePlaceholder)
	replacements.ReplacePath(".alertPolicies[].mutationRecord.mutateTime", timePlaceholder)

	replacements.ReplacePath(".mutationRecords[].mutateTime", timePlaceholder)
	replacements.ReplacePath(".notificationChannels[].mutationRecords[].mutateTime", timePlaceholder)
	replacements.ReplacePath(".alertPolicies[].mutationRecords[].mutateTime", timePlaceholder)

	replacements.ReplacePath(".creationRecord.mutatedBy", emailPlaceholder)
	replacements.ReplacePath(".mutationRecord.mutatedBy", emailPlaceholder)
	replacements.ReplacePath(".mutationRecords[].mutatedBy", emailPlaceholder)
	replacements.ReplacePath(".alertPolicies[].creationRecord.mutatedBy", emailPlaceholder)
	replacements.ReplacePath(".alertPolicies[].mutationRecord.mutatedBy", emailPlaceholder)
	replacements.ReplacePath(".alertPolicies[].mutationRecords[].mutatedBy", emailPlaceholder)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	visitLink := func(name string) {
		tokens := strings.Split(name, "/")
		n := len(tokens)
		if n > 2 && tokens[n-2] == "conditions" {
			replacements.ReplaceStringValue(tokens[n-1], "${conditionID}")
		}
	}

	event.VisitResponseStringValues(func(path string, value string) {
		switch path {
		case ".conditions[].name":
			// In alertPolicy
			//  e.g. "name": "projects/${projectId}/alertPolicies/${alertPolicyID}/conditions/5683099026496458999"
			visitLink(value)
		}
	})
}

func (s *MockService) ConfigureKRMObjectVisitor(u *unstructured.Unstructured, replacements mockgcpregistry.NormalizingVisitor) {
}
