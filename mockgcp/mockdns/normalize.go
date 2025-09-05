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
	"regexp"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var (
	// Example: ns-cloud-e1.googledomains.com.
	dnsNameServerRegex = regexp.MustCompile(`"ns-cloud-[a-z]([1-4]).googledomains.com."`)
)

const PlaceholderTimestamp = "2024-04-01T12:34:56.123456Z"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// DNS ManagedZone
	replacements.ReplacePath(".nameServers", []string{"ns-cloud-a1.googledomains.com.", "ns-cloud-a2.googledomains.com.", "ns-cloud-a3.googledomains.com.", "ns-cloud-a4.googledomains.com."})
	replacements.ReplacePath(".managedZones[].creationTime", PlaceholderTimestamp)
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {

	kind := ""

	event.VisitResponseStringValues(func(path string, value string) {
		switch path {
		case ".kind":
			kind = value
		}
	})

	if kind == "dns#managedZone" {
		event.VisitResponseStringValues(func(path string, value string) {
			switch path {
			case ".id":
				replacements.ReplaceStringValue(value, "${managedZoneId}")
			}
		})
	}
}
