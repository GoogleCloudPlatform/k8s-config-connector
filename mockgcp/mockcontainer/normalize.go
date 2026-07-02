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

package mockcontainer

import (
	"net"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

var gkeVersionRegex = regexp.MustCompile(`\d+\.\d+\.\d+-gke\.\d+`)

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !isContainerAPI(url) {
		return
	}

	// Cluster
	{
		replacements.ReplacePath(".clusterIpv4Cidr", "10.112.0.0/14")
		replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4Cidr", "10.112.0.0/14")
		replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4CidrBlock", "10.112.0.0/14")

		replacements.ReplacePath(".maintenancePolicy.resourceVersion", "abcd1234")

		replacements.SortSlice(".monitoringConfig.componentConfig.enableSystemComponents")

		// Remove unsupported newer fields
		replacements.RemovePath(".controlPlaneEgress")
		replacements.RemovePath(".response.controlPlaneEgress")
		replacements.RemovePath(".nodeCreationConfig")
		replacements.RemovePath(".response.nodeCreationConfig")
		replacements.RemovePath(".nodeImageConfig")
		replacements.RemovePath(".response.nodeImageConfig")
		replacements.RemovePath(".nodeConfig.nodeImageConfig")
		replacements.RemovePath(".response.nodeConfig.nodeImageConfig")
		replacements.RemovePath(".nodePools[].config.nodeImageConfig")
		replacements.RemovePath(".response.nodePools[].config.nodeImageConfig")
		replacements.RemovePath(".config.nodeImageConfig")
		replacements.RemovePath(".response.config.nodeImageConfig")
		replacements.RemovePath(".addonsConfig.nodeReadinessConfig")
		replacements.RemovePath(".response.addonsConfig.nodeReadinessConfig")

		// Remove volatile or mismatched network tier/utilization configs
		replacements.RemovePath(".networkConfig.networkTierConfig")
		replacements.RemovePath(".response.networkConfig.networkTierConfig")
		replacements.RemovePath(".networkConfig.defaultSnatStatus")
		replacements.RemovePath(".response.networkConfig.defaultSnatStatus")
		replacements.RemovePath(".nodePools[].networkConfig.networkTierConfig")
		replacements.RemovePath(".response.nodePools[].networkConfig.networkTierConfig")
		replacements.RemovePath(".ipAllocationPolicy.networkTierConfig")
		replacements.RemovePath(".response.ipAllocationPolicy.networkTierConfig")
		replacements.RemovePath(".ipAllocationPolicy.defaultPodIpv4RangeUtilization")
		replacements.RemovePath(".response.ipAllocationPolicy.defaultPodIpv4RangeUtilization")
		replacements.RemovePath(".nodePools[].networkConfig.podIpv4RangeUtilization")
		replacements.RemovePath(".response.nodePools[].networkConfig.podIpv4RangeUtilization")

		// Normalize or remove GKE modern defaulted fields to maintain backward compatibility with existing cluster tests
		replacements.ReplacePath(".anonymousAuthenticationConfig.mode", "ENABLED")
		replacements.ReplacePath(".response.anonymousAuthenticationConfig.mode", "ENABLED")
		replacements.RemovePath(".addonsConfig.dnsCacheConfig")
		replacements.RemovePath(".response.addonsConfig.dnsCacheConfig")
	}
}

func isContainerAPI(url string) bool {
	return strings.HasPrefix(url, "https://container.googleapis.com/")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "container.googleapis.com") && !strings.Contains(event.URL(), "compute.googleapis.com") {
		return
	}

	event.VisitResponseStringValues(func(path string, value string) {
		// 1. Replace GKE versions
		if gkeVersionRegex.MatchString(value) {
			newValue := gkeVersionRegex.ReplaceAllString(value, "1.30.5-gke.1014001")
			replacements.ReplaceStringValue(value, newValue)
			value = newValue
		}

		// 2. Normalize compute instance templates URLs (global vs regional)
		if strings.Contains(value, "global/instanceTemplates/gke-") {
			newValue := strings.ReplaceAll(value, "global/instanceTemplates/gke-", "regions/us-central1/instanceTemplates/gke-")
			replacements.ReplaceStringValue(value, newValue)
			value = newValue
		}

		// 3. Normalize custom node pool instance template names
		if strings.Contains(value, "gke-containercluster-abcdef-nodepool-sample-") {
			re := regexp.MustCompile(`gke-containercluster-abcdef-nodepool-sample-[a-zA-Z0-9]+`)
			newValue := re.ReplaceAllString(value, "gke-cluster-sample-k-nodepool-sample--73473afd")
			replacements.ReplaceStringValue(value, newValue)
			value = newValue
		}

		// 4. Normalize custom node pool IGM names
		if strings.Contains(value, "-nodepool-sample-") {
			re := regexp.MustCompile(`gke-(containercluster-abcdef|cluster-sample-[a-zA-Z0-9]+)-nodepool-sample-[a-zA-Z0-9-]*(-grp)?`)
			newValue := re.ReplaceAllStringFunc(value, func(m string) string {
				if strings.HasSuffix(m, "-grp") {
					return "gke-cluster-sample-k-nodepool-sample--15cfd728-grp"
				}
				return "gke-cluster-sample-k-nodepool-sample--15cfd728"
			})
			replacements.ReplaceStringValue(value, newValue)
			value = newValue
		}

		// 5. Normalize default pool names
		if strings.Contains(value, "-default-pool") {
			re := regexp.MustCompile(`gke-cluster-sample-[a-zA-Z0-9]+-default-pool(-[a-zA-Z0-9-]*)?(-grp)?`)
			newValue := re.ReplaceAllStringFunc(value, func(m string) string {
				if strings.HasSuffix(m, "-grp") {
					return "gke-cluster-sample-ke4qs-default-pool-170ea918-grp"
				}
				return "gke-cluster-sample-ke4qs-default-pool-170ea918"
			})
			replacements.ReplaceStringValue(value, newValue)
			value = newValue
		}

		// 6. Replace public/private IP addresses
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

	event.VisitRequestStringValues(func(path string, value string) {
		if gkeVersionRegex.MatchString(value) {
			newValue := gkeVersionRegex.ReplaceAllString(value, "1.30.5-gke.1014001")
			replacements.ReplaceStringValue(value, newValue)
		}
		if strings.Contains(value, "global/instanceTemplates/gke-") {
			newValue := strings.ReplaceAll(value, "global/instanceTemplates/gke-", "regions/us-central1/instanceTemplates/gke-")
			replacements.ReplaceStringValue(value, newValue)
		}
	})
}

// Simple check for IPv4 address format.
func isIPv4Address(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}
