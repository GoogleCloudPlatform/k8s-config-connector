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
		replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4Cidr", "10.112.0.0/14")
		replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4CidrBlock", "10.112.0.0/14")
		replacements.ReplacePath(".networkConfig.podIpv4CidrBlock", "10.112.0.0/14")

		replacements.ReplacePath(".maintenancePolicy.resourceVersion", "abcd1234")

		replacements.SortSlice(".monitoringConfig.componentConfig.enableSystemComponents")
	}
}

func isContainerAPI(url string) bool {
	return strings.HasPrefix(url, "https://container.googleapis.com/")
}

func isComputeAPI(url string) bool {
	return strings.HasPrefix(url, "https://compute.googleapis.com/") ||
		strings.HasPrefix(url, "https://www.googleapis.com/compute/")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// =========================================================================
	// 1. Container API (https://container.googleapis.com/...)
	// =========================================================================
	if isContainerAPI(event.URL()) {
		event.VisitResponseStringValues(func(path string, value string) {
			switch {
			// IP Endpoints
			case path == ".controlPlaneEndpointsConfig.ipEndpointsConfig.publicEndpoint" ||
				path == ".privateClusterConfig.publicEndpoint":
				if isIPv4Address(value) {
					replacements.ReplaceStringValue(value, "${publicEndpointIPV4}")
				}
			case path == ".controlPlaneEndpointsConfig.ipEndpointsConfig.privateEndpoint" ||
				path == ".privateClusterConfig.privateEndpoint":
				if isIPv4Address(value) {
					replacements.ReplaceStringValue(value, "${privateEndpointIPV4}")
				}
			// GKE Cluster & NodePool Versions
			case path == ".version" ||
				path == ".currentMasterVersion" ||
				path == ".currentNodeVersion" ||
				path == ".initialClusterVersion" ||
				strings.HasSuffix(path, ".version"):
				if regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+-gke\.[0-9]+$`).MatchString(value) {
					replacements.ReplaceStringValue(value, "1.30.5-gke.1014001")
				}
			// Node Base Image
			case strings.HasSuffix(path, ".nodeImageConfig.image"):
				replacements.ReplaceStringValue(value, "gke-1305-gke1014001-cos-113-17681-1153-62-c-pre")
			// Cluster Master CA Certificate
			case path == ".masterAuth.clusterCaCertificate":
				replacements.ReplaceStringValue(value, "1234567890abcdefghijklmn")
			// TPM Bootstrap Certificate Expiration Timestamp
			case strings.HasSuffix(path, ".kubeletCertInfo.tpmBootstrapCertExpireTime"):
				replacements.ReplaceStringValue(value, mockgcpregistry.PlaceholderTimestamp)
			// GKE DNS Endpoint
			case path == ".controlPlaneEndpointsConfig.dnsEndpointConfig.endpoint":
				dnsEndpointRegex := regexp.MustCompile(`gke-[a-f0-9]{40}`)
				if dnsEndpointRegex.MatchString(value) {
					match := dnsEndpointRegex.FindString(value)
					replacements.ReplaceStringValue(match, "gke-12345trewq")
				}
			// Secondary Pod CIDR Range Name (normalizes the dynamic 8-hex-char hash suffix)
			case strings.HasSuffix(path, ".podRange"):
				podRangeRegex := regexp.MustCompile(`(gke-[a-z0-9-${}]+-pods-)[0-9a-f]{8}`)
				if podRangeRegex.MatchString(value) {
					match := podRangeRegex.FindString(value)
					norm := podRangeRegex.ReplaceAllString(match, "${1}7b2f8c84")
					replacements.ReplaceStringValue(match, norm)
				}
			// Instance Group URLs inside Container responses
			case strings.Contains(path, "instanceGroupUrls"):
				normalizeGKEInstanceGroupNames(value, replacements)
			}
		})
		return
	}
	// =========================================================================
	// 2. Compute API (https://compute.googleapis.com/...)
	// =========================================================================
	if isComputeAPI(event.URL()) {
		event.VisitResponseStringValues(func(path string, value string) {
			switch {
			// Only normalize GKE-generated IGM names in Compute responses
			case path == ".name" || path == ".instanceGroup" || strings.Contains(path, "instanceGroup"):
				normalizeGKEInstanceGroupNames(value, replacements)
			}
		})
		return
	}
}

// Helper to normalize GKE-generated Instance Group Manager names
func normalizeGKEInstanceGroupNames(value string, replacements mockgcpregistry.NormalizingVisitor) {
	igmRegex := regexp.MustCompile(`gke-[a-z0-9-]+-[a-f0-9]+-grp`)
	if igmRegex.MatchString(value) {
		match := igmRegex.FindString(value)
		replacements.ReplaceStringValue(match, "gke-containercluster-abcdef-normalized-grp")
	}
}

// Simple check for IPv4 address format.
func isIPv4Address(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}
