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

		replacements.RemovePath(".networkConfig.networkTierConfig")
		replacements.RemovePath(".nodePools[].networkConfig.networkTierConfig")

		replacements.RemovePath(".config.nodeImageConfig")
		replacements.RemovePath(".nodePools[].config.nodeImageConfig")

		replacements.RemovePath(".initialNodeCount")
		replacements.RemovePath(".nodePools[].initialNodeCount")

		replacements.RemovePath(".instanceGroupUrls")
		replacements.RemovePath(".nodePools[].instanceGroupUrls")

		replacements.RemovePath(".kubeletCertInfo")
		replacements.RemovePath(".nodePools[].kubeletCertInfo")

		// Clean up etag for GKE clusters / node pools without leakage to other services like IAM
		replacements.TransformObject("", func(m map[string]any) {
			isGKECluster := m["monitoringService"] != nil || m["databaseEncryption"] != nil || m["nodeConfig"] != nil || m["addonsConfig"] != nil
			isGKENodePool := m["config"] != nil || m["upgradeSettings"] != nil
			if isGKECluster || isGKENodePool {
				delete(m, "etag")
			}
		})
		replacements.TransformObject(".nodePools[]", func(m map[string]any) {
			delete(m, "etag")
		})
	}
}

func isContainerAPI(url string) bool {
	return strings.HasPrefix(url, "https://container.googleapis.com/")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isContainerAPI(event.URL()) {
		return
	}

	// Normalize GKE version
	versionRegex := regexp.MustCompile(`^1\.\d+\.\d+-gke\.\d+$`)
	// Normalize GKE IP/CIDR
	cidrRegex := regexp.MustCompile(`^10\.\d+\.\d+\.\d+/\d+$`)
	// Normalize GKE pod range names
	podRangeRegex := regexp.MustCompile(`^gke-[a-zA-Z0-9\-\${}]+-pods-[0-9a-fA-F\d]+$`)

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

		if versionRegex.MatchString(value) {
			replacements.ReplaceStringValue(value, "1.35.6-gke.1049000")
		} else if cidrRegex.MatchString(value) {
			replacements.ReplaceStringValue(value, "10.0.0.0/14")
		} else if podRangeRegex.MatchString(value) {
			replacements.ReplaceStringValue(value, "gke-cluster-sample-${uniqueId}-pods-abcdef01")
		}
	})
}

// Simple check for IPv4 address format.
func isIPv4Address(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}
