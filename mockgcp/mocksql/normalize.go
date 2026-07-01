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

package mocksql

import (
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"k8s.io/klog/v2"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// SQLUser
	replacements.ReplacePath(".items[].passwordPolicy.status.passwordExpirationTime", "2025-06-19T01:02:03Z")

	// BackupRun
	replacements.ReplacePath(".startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".enqueuedTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".windowStartTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".items[].startTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].endTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].enqueuedTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].windowStartTime", mockgcpregistry.PlaceholderTimestamp)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isSQLAPI(event) {
		return
	}

	u, err := url.Parse(event.URL())
	if err == nil {
		path := u.Path
		idx := strings.LastIndex(path, "/backupRuns/")
		if idx != -1 {
			idStr := path[idx+len("/backupRuns/"):]
			if qIdx := strings.Index(idStr, "?"); qIdx != -1 {
				idStr = idStr[:qIdx]
			}
			if sIdx := strings.Index(idStr, "/"); sIdx != -1 {
				idStr = idStr[:sIdx]
			}
			if idStr != "" {
				replacements.ReplaceStringValue(idStr, "000000000000000000000")
			}
		}
	}

	var sqlInstance struct {
		IPAddresses []struct {
			IPAddress string `json:"ipAddress"`
			Type      string `json:"type"`
		} `json:"ipAddresses"`
	}
	if ok := event.ParseResponseInto(&sqlInstance); ok {
		for _, ipAddress := range sqlInstance.IPAddresses {
			switch ipAddress.Type {
			case "PRIVATE":
				replacements.ReplaceStringValue(ipAddress.IPAddress, "10.1.2.3")
			case "PRIMARY":
				replacements.ReplaceStringValue(ipAddress.IPAddress, "10.10.10.10")
			case "OUTGOING":
				replacements.ReplaceStringValue(ipAddress.IPAddress, "10.10.10.11")
			}
			if ipAddress.Type == "PRIMARY" {
				replacements.ReplaceStringValue(ipAddress.IPAddress, "10.10.10.10")
			}
		}
	}
}

// isSQLAPI returns true if this is a sql URL
func isSQLAPI(event mockgcpregistry.Event) bool {
	u, err := url.Parse(event.URL())
	if err != nil {
		klog.Fatalf("cannot parse URL %q", event.URL())
	}
	switch u.Host {
	case "sqladmin.googleapis.com":
		return true
	}
	return false
}
