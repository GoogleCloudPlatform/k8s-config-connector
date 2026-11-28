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

package mockvmwareengine

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "vmwareengine.googleapis.com") {
		return
	}
	replacements.ReplacePath(".createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".response.updateTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".metadata.endTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".uid", "111111111111111111111")
	replacements.ReplacePath(".response.uid", "111111111111111111111")

	replacements.ReplacePath(".etag", "abcdef0123A=")
	replacements.ReplacePath(".response.etag", "abcdef0123A=")

	// PrivateConnection peeringId
	replacements.ReplacePath(".peeringId", "peering-55c50401-d392-4c50-94af-d4a1242f70ac")
	replacements.ReplacePath(".response.peeringId", "peering-55c50401-d392-4c50-94af-d4a1242f70ac")

	// PrivateCloud generated fields
	components := []struct {
		field   string
		prefix  string
		version string
	}{
		{field: "hcx", prefix: "hcx-414861.c8819727", version: "4.10.3.24447633"},
		{field: "nsx", prefix: "nsx-414860.c8819727", version: "3.2.3.1"},
		{field: "vcenter", prefix: "vcsa-359395.c8819727", version: "7.0.3.23085514"},
	}
	for _, c := range components {
		c := c
		for _, root := range []string{"", ".response", ".status.observedState"} {
			replacements.TransformString(root+"."+c.field+".fqdn", func(s string) string {
				return normalizeVMwareEngineFQDN(s, c.prefix)
			})
			replacements.ReplacePath(root+"."+c.field+".version", c.version)
		}
	}
}

func normalizeVMwareEngineFQDN(fqdn string, hostClusterPrefix string) string {
	if fqdn == "" {
		return ""
	}
	region := "us-west2"
	if parts := strings.Split(fqdn, "."); len(parts) == 5 && parts[3] == "gve" && parts[4] == "goog" && parts[2] != "" {
		region = parts[2]
	}
	return hostClusterPrefix + "." + region + ".gve.goog"
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
