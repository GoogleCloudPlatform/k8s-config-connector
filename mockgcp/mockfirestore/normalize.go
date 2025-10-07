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

package mockfirestore

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

const TimePlaceholder = "2024-04-01T12:34:56.123456Z"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Database
	replacements.ReplacePath(".response.earliestVersionTime", TimePlaceholder)
	replacements.ReplacePath(".earliestVersionTime", TimePlaceholder)

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// // TODO: If all normalizers should filter by host, we could auto-filter (i.e. only call normalizers with matching hosts)
	if strings.Contains(event.URL(), "https://firestore.googleapis.com") {
		previousId := ""
		event.VisitResponseStringValues(func(path string, value string) {
			switch path {
			case ".response.previousId", ".previousId":
				previousId = value
			}
		})

		event.VisitResponseStringValues(func(path string, value string) {
			switch path {
			case ".response.name", ".name":
				if previousId != "" {
					tokens := strings.Split(value, "/")
					if len(tokens) == 4 && tokens[2] == "databases" {
						replacements.ReplaceStringValue(tokens[3], "${randomDatabaseID}")
					}
				}
			}
		})

	}
}
