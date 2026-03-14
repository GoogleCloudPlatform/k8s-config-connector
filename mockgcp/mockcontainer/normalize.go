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

package mockcontainer

import (
	"net"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

const PlaceholderTimestamp = "2024-04-01T12:34:56.123456Z"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !isContainerAPI(url) {
		return
	}

	// Cluster
	{
		replacements.ReplacePath(".clusterIpv4Cidr", "${clusterIpv4Cidr}")
		replacements.ReplacePath(".servicesIpv4Cidr", "${servicesIpv4Cidr}")

		replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4Cidr", "${clusterIpv4Cidr}")
		replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4CidrBlock", "${clusterIpv4Cidr}")
		replacements.ReplacePath(".ipAllocationPolicy.servicesIpv4Cidr", "${servicesIpv4Cidr}")
		replacements.ReplacePath(".ipAllocationPolicy.servicesIpv4CidrBlock", "${servicesIpv4Cidr}")

		replacements.ReplacePath(".maintenancePolicy.resourceVersion", "abcd1234")

		replacements.SortSlice(".monitoringConfig.componentConfig.enableSystemComponents")
	}

	// NodePool
	{
		replacements.ReplacePath(".podIpv4CidrSize", "24")
		replacements.ReplacePath(".networkConfig.podIpv4CidrBlock", "${podIpv4CidrBlock}")
	}
}

func isContainerAPI(url string) bool {
	return strings.HasPrefix(url, "https://container.googleapis.com/")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isContainerAPI(event.URL()) {
		return
	}

	// Replace public IP addresses with placeholders.
	event.VisitResponseStringValues(func(path string, value string) {
		switch path {
		case ".controlPlaneEndpointsConfig.ipEndpointsConfig.publicEndpoint",
			".privateClusterConfig.publicEndpoint":
			if isIPv4Address(value) {
				replacements.ReplaceStringValue(value, "${publicEndpointIPV4}")
			}

		case ".controlPlaneEndpointsConfig.ipEndpointsConfig.privateEndpoint",
			".privateClusterConfig.privateEndpoint":
			if isIPv4Address(value) {
				replacements.ReplaceStringValue(value, "${privateEndpointIPV4}")
			}
		}
	})
}

// Simple check for IPv4 address format.
func isIPv4Address(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}
