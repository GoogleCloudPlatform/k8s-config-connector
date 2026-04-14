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

package mockmemorystore

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

const (
	PlaceholderTimestamp = "2024-04-01T12:34:56.123456Z"
	PlaceholderUID       = "11111111-1111-1111-1111-111111111111"
)

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.HasPrefix(url, "https://memorystore.googleapis.com") && !strings.HasPrefix(url, "https://memorystore.mtls.googleapis.com") {
		return
	}

	replacements.ReplacePath(".crossInstanceReplicationConfig.updateTime", PlaceholderTimestamp)
	replacements.ReplacePath(".expireTime", PlaceholderTimestamp)
	replacements.ReplacePath(".maintenancePolicy.createTime", PlaceholderTimestamp)
	replacements.ReplacePath(".maintenancePolicy.updateTime", PlaceholderTimestamp)
	replacements.ReplacePath(".response.crossInstanceReplicationConfig.updateTime", PlaceholderTimestamp)
	replacements.ReplacePath(".response.expireTime", PlaceholderTimestamp)
	replacements.ReplacePath(".response.maintenancePolicy.createTime", PlaceholderTimestamp)
	replacements.ReplacePath(".response.maintenancePolicy.updateTime", PlaceholderTimestamp)

	// Backups
	replacements.ReplacePath(".backupFiles[].createTime", PlaceholderTimestamp)
	replacements.ReplacePath(".backupFiles[].fileName", "fakefilename.rdb")
	replacements.ReplacePath(".response.backupFiles[].createTime", PlaceholderTimestamp)
	replacements.ReplacePath(".instanceUid", PlaceholderUID)
	replacements.ReplacePath(".encryptionInfo.lastUpdateTime", PlaceholderTimestamp)

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if strings.Contains(event.URL(), "https://memorystore.googleapis.com") || strings.HasPrefix(event.URL(), "https://memorystore.mtls.googleapis.com") {

		event.VisitResponseStringValues(func(path string, value string) {
			switch path {
			case ".name":
				tokens := strings.Split(value, "/")

				if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupCollections" && tokens[6] == "backups" {
					replacements.ReplaceStringValue(tokens[5], "${backupCollectionID}")
					replacements.ReplaceStringValue(tokens[7], "${backup}")
				}

			case ".pscAttachmentDetails[].serviceAttachment":
				tokens := strings.Split(value, "/")

				if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "serviceAttachments" {
					replacements.ReplaceStringValue(tokens[1], "${pscProjectNumber}")
					//replacements.ReplaceStringValue(tokens[3], "${region}")
					if strings.HasSuffix(tokens[5], "-2") {
						replacements.ReplaceStringValue(tokens[5], "${pscServiceAttachment}-2")
					} else {
						replacements.ReplaceStringValue(tokens[5], "${pscServiceAttachment}")
					}
				}
			}
		})
	}
}
