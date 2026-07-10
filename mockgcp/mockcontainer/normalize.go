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
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !isContainerAPI(url) {
		return
	}

	// Cluster
	{
		replacements.ReplacePath(".clusterIpv4Cidr", "10.112.0.0/14")

		replacements.ReplacePath(".clusterIpv4Cidr", "10.112.0.0/14")
		replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4Cidr", "10.112.0.0/14")
		replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4CidrBlock", "10.112.0.0/14")

		replacements.ReplacePath(".maintenancePolicy.resourceVersion", "abcd1234")

		replacements.SortSlice(".monitoringConfig.componentConfig.enableSystemComponents")

		replacements.RemovePath(".addonsConfig.dnsCacheConfig")
		replacements.RemovePath(".addonsConfig.nodeReadinessConfig")
		replacements.RemovePath(".controlPlaneEgress")
		replacements.RemovePath(".nodeCreationConfig")
		replacements.RemovePath(".nodePoolAutoConfig")
		replacements.RemovePath(".networkConfig.defaultSnatStatus")
		replacements.RemovePath(".privateCluster")
		replacements.RemovePath(".nodeConfig.nodeImageConfig")
		replacements.RemovePath(".nodePools[].config.nodeImageConfig")
		replacements.RemovePath(".nodePools[].networkConfig.networkTierConfig")
		replacements.RemovePath(".nodePools[].locations")
		replacements.RemovePath(".nodePools[].etag")
		replacements.RemovePath(".nodePools[].selfLink")
		replacements.RemovePath(".anonymousAuthenticationConfig")
		replacements.RemovePath(".ipAllocationPolicy.clusterSecondaryRangeName")
		replacements.RemovePath(".ipAllocationPolicy.defaultPodIpv4RangeUtilization")
		replacements.RemovePath(".ipAllocationPolicy.networkTierConfig")
		replacements.RemovePath(".master")
		replacements.RemovePath(".etag")
		replacements.RemovePath(".locations")
		replacements.RemovePath(".networkConfig.networkTierConfig")
	}
}

func isContainerAPI(url string) bool {
	return strings.HasPrefix(url, "https://container.googleapis.com/")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isContainerAPI(event.URL()) {
		return
	}

	reGKEVersion := regexp.MustCompile(`\b\d+\.\d+\.\d+-gke\.\d+\b`)
	reGKEIGM := regexp.MustCompile(`instanceGroupManagers/gke-([a-z0-9\-]+)-grp`)
	reGKEIG := regexp.MustCompile(`instanceGroups/gke-([a-z0-9\-]+)-grp`)

	// Replace public IP addresses with placeholders.
	event.VisitResponseStringValues(func(path string, value string) {
		if strings.Contains(value, "${uniqueId}") {
			return
		}

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

		if reGKEVersion.MatchString(value) {
			replacements.ReplaceStringValue(value, "1.30.5-gke.1014001")
		}

		if reGKEIGM.MatchString(value) || reGKEIG.MatchString(value) {
			newValue := value
			newValue = reGKEIGM.ReplaceAllStringFunc(newValue, func(match string) string {
				if strings.Contains(match, "default-pool") {
					return "instanceGroupManagers/gke-containercluster-abcdef-default-pool-grp"
				}
				return "instanceGroupManagers/gke-containercluster-abcdef-nodepool-sample-${uniqueId}-grp"
			})
			newValue = reGKEIG.ReplaceAllStringFunc(newValue, func(match string) string {
				if strings.Contains(match, "default-pool") {
					return "instanceGroups/gke-containercluster-abcdef-default-pool-grp"
				}
				return "instanceGroups/gke-containercluster-abcdef-nodepool-sample-${uniqueId}-grp"
			})
			replacements.ReplaceStringValue(value, newValue)
		}

		if value == "PENDING" && (path == ".status" || strings.HasSuffix(path, ".status")) {
			replacements.ReplaceStringValue(value, "RUNNING")
		}
	})
}

// Simple check for IPv4 address format.
func isIPv4Address(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}
