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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

const (
	emailPlaceholder = "user@example.com"
	timePlaceholder  = "2024-04-01T12:34:56.123456Z"
)

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	replacements.ReplacePath(".creationRecord.mutateTime", timePlaceholder)
	replacements.ReplacePath(".notificationChannels[].creationRecord.mutateTime", timePlaceholder)

	replacements.ReplacePath(".mutationRecord.mutateTime", timePlaceholder)
	replacements.ReplacePath(".notificationChannels[].mutationRecord.mutateTime", timePlaceholder)

	replacements.ReplacePath(".mutationRecords[].mutateTime", timePlaceholder)
	replacements.ReplacePath(".notificationChannels[].mutationRecords[].mutateTime", timePlaceholder)

	replacements.ReplacePath(".creationRecord.mutatedBy", emailPlaceholder)
	replacements.ReplacePath(".mutationRecord.mutatedBy", emailPlaceholder)
	replacements.ReplacePath(".mutationRecords[].mutatedBy", emailPlaceholder)

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
