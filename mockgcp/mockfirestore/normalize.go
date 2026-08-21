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
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"k8s.io/klog/v2"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Database
	replacements.ReplacePath(".response.earliestVersionTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".earliestVersionTime", mockgcpregistry.PlaceholderTimestamp)

	// freeTier cannot be set (the first database is always freeTier=true)
	replacements.ReplacePath(".freeTier", true)
	replacements.ReplacePath(".response.freeTier", true)

	// realtimeUpdatesMode is not yet in our protos.
	replacements.RemovePath(".realtimeUpdatesMode")
	replacements.RemovePath(".response.realtimeUpdatesMode")

	// Fields
	replacements.ReplacePath(".response.startTime", mockgcpregistry.PlaceholderTimestamp)

	// BackupSchedules
	replacements.ReplacePath(".backupSchedules[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".backupSchedules[].updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	ctx := context.TODO()
	log := klog.FromContext(ctx)

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
				tokens := strings.Split(value, "/")

				if previousId != "" {
					if len(tokens) == 4 && tokens[2] == "databases" {
						log.Info("normalizing previousId in database name", "path", path, "name", value, "previousId", previousId)
						replacements.ReplaceStringValue(tokens[3], "${randomDatabaseID}")
					}
				}
				if len(tokens) == 6 && tokens[2] == "databases" && tokens[4] == "backupSchedules" {
					replacements.ReplaceStringValue(tokens[5], "${backupScheduleID}")
				}
				if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "databases" && tokens[4] == "collectionGroups" && tokens[6] == "indexes" {
					replacements.ReplaceStringValue(tokens[7], "${indexID}")
				}
			}
		})
	}
}
