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

package mockdns

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

var placeholderNameServers = []string{"ns-cloud-a1.googledomains.com.", "ns-cloud-a2.googledomains.com.", "ns-cloud-a3.googledomains.com.", "ns-cloud-a4.googledomains.com."}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !isDNSAPI(url) {
		return
	}

	// DNS ManagedZone
	{
		replacements.ReplacePath(".nameServers", placeholderNameServers)
		replacements.ReplacePath(".managedZones[].nameServers", placeholderNameServers)
		replacements.ReplacePath(".zoneContext.newValue.nameServers", placeholderNameServers)
		replacements.ReplacePath(".zoneContext.oldValue.nameServers", placeholderNameServers)

		replacements.ReplacePath(".creationTime", mockgcpregistry.PlaceholderTimestamp)
		replacements.ReplacePath(".managedZones[].creationTime", mockgcpregistry.PlaceholderTimestamp)
		replacements.ReplacePath(".zoneContext.newValue.creationTime", mockgcpregistry.PlaceholderTimestamp)
		replacements.ReplacePath(".zoneContext.oldValue.creationTime", mockgcpregistry.PlaceholderTimestamp)
	}
}

func isDNSAPI(url string) bool {
	return strings.HasPrefix(url, "https://dns.googleapis.com/")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isDNSAPI(event.URL()) {
		return
	}

	parentToKind := make(map[string]string)

	event.VisitResponseStringValues(func(path string, value string) {
		if strings.HasSuffix(path, ".kind") {
			parent := strings.TrimSuffix(path, ".kind")
			parentToKind[parent] = value
		}
	})

	event.VisitResponseStringValues(func(path string, value string) {
		if strings.HasSuffix(path, ".id") {
			parent := strings.TrimSuffix(path, ".id")
			if kind, ok := parentToKind[parent]; ok {
				if kind == "dns#policy" {
					replacements.ReplaceStringValue(value, "${dnsPolicyId}")
				}
				if kind == "dns#managedZone" {
					replacements.ReplaceStringValue(value, "${managedZoneId}")
				}
				if kind == "dns#responsePolicy" {
					replacements.ReplaceStringValue(value, "${dnsResponsePolicyId}")
				}
			}
		}
	})
}
