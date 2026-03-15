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

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Cluster
	{
		replacements.ReplacePath(".maintenancePolicy.resourceVersion", "abcd1234")

		replacements.SortSlice(".monitoringConfig.componentConfig.enableSystemComponents")
		replacements.SortSlice(".loggingConfig.componentConfig.enableComponents")
	}

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isContainerAPI(event.URL()) {
		return
	}

	// Capture IP ranges and addresses for normalization.
	event.VisitResponseStringValues(func(path string, value string) {
		if value == "" {
			return
		}

		// Normalize Cluster and Service CIDR ranges.
		// These values are often reused in different parts of the response (e.g. IpAllocationPolicy).
		// We use ReplacePath to handle masks (like "/20") which might be the same for both cluster and services.
		// We also use ReplaceStringValue for full ranges (containing ".") to ensure they are replaced globally.
		if strings.HasSuffix(path, "clusterIpv4Cidr") || strings.HasSuffix(path, "clusterIpv4CidrBlock") || strings.HasSuffix(path, "podIpv4CidrBlock") {
			replacements.ReplacePath(path, "${clusterIpv4Cidr}")
			if strings.Contains(value, ".") {
				replacements.ReplaceStringValue(value, "${clusterIpv4Cidr}")
			}
		}
		if strings.HasSuffix(path, "servicesIpv4Cidr") || strings.HasSuffix(path, "servicesIpv4CidrBlock") {
			replacements.ReplacePath(path, "${servicesIpv4Cidr}")
			if strings.Contains(value, ".") {
				replacements.ReplaceStringValue(value, "${servicesIpv4Cidr}")
			}
		}

		// Replace public/private endpoint IP addresses with placeholders.
		switch path {
		case ".controlPlaneEndpointsConfig.ipEndpointsConfig.publicEndpoint",
			".privateClusterConfig.publicEndpoint",
			".endpoint":
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

func isContainerAPI(url string) bool {
	return strings.Contains(url, "container.googleapis.com")
}

// Simple check for IPv4 address format.
func isIPv4Address(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}
