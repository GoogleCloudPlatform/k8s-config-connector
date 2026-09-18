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
	if !strings.HasPrefix(url, "https://container.googleapis.com/") {
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
	}
}

func isContainerAPI(url string) bool {
	return strings.HasPrefix(url, "https://container.googleapis.com/") ||
		strings.HasPrefix(url, "https://compute.googleapis.com/") ||
		strings.HasPrefix(url, "https://www.googleapis.com/compute/")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isContainerAPI(event.URL()) {
		return
	}

	// Extract unique ID from the URL if present
	var uniqueID string
	if match := regexp.MustCompile(`clusters/cluster-sample-([a-z0-9]+)`).FindStringSubmatch(event.URL()); len(match) > 1 {
		uniqueID = match[1]
	}

	// Replace public IP addresses with placeholders and normalize dynamic strings.
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

		// Normalize GKE versions
		if strings.HasPrefix(event.URL(), "https://container.googleapis.com/") && strings.Contains(event.URL(), "clusters/cluster-sample") {
			if path == ".currentMasterVersion" || path == ".currentNodeVersion" || path == ".initialClusterVersion" || path == ".version" || strings.HasSuffix(path, ".version") {
				if regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+-gke\.[0-9]+$`).MatchString(value) {
					replacements.ReplaceStringValue(value, "1.30.5-gke.1014001")
				}
			}

			// CA Certificate
			if path == ".masterAuth.clusterCaCertificate" {
				replacements.ReplaceStringValue(value, "1234567890abcdefghijklmn")
			}

			// TPM Bootstrap Cert Expire Time
			if strings.HasSuffix(path, ".kubeletCertInfo.tpmBootstrapCertExpireTime") {
				replacements.ReplaceStringValue(value, mockgcpregistry.PlaceholderTimestamp)
			}

			// Image
			if strings.HasSuffix(path, ".nodeImageConfig.image") || path == ".config.nodeImageConfig.image" {
				replacements.ReplaceStringValue(value, "gke-1305-gke1014001-cos-113-17681-1153-62-c-pre")
			}
		}

		// Find and normalize dns endpoint hashes (e.g. gke-7b2f8c845b534c84a78c3f1cb6f6cfdded1ad359)
		dnsEndpointRegex := regexp.MustCompile(`gke-[a-f0-9]{40}`)
		if dnsEndpointRegex.MatchString(value) {
			match := dnsEndpointRegex.FindString(value)
			replacements.ReplaceStringValue(match, "gke-12345trewq")
		}

		// Normalize default-pool instance group name (e.g. gke-cluster-sample-[a-z0-9]+-default-pool-[a-f0-9]+-grp)
		defaultPoolRegex := regexp.MustCompile(`gke-cluster-sample-[a-z0-9]+-default-pool-[a-f0-9]+-grp`)
		if defaultPoolRegex.MatchString(value) {
			match := defaultPoolRegex.FindString(value)
			replacements.ReplaceStringValue(match, "gke-containercluster-abcdef-default-pool-grp")
		}

		// Normalize nodepool-sample instance group name (e.g. gke-cluster-sample-[a-z0-9]+-nodepool-sample--[a-f0-9]+-grp)
		if uniqueID != "" {
			nodepoolRegex := regexp.MustCompile(`gke-cluster-sample-[a-z0-9]+-nodepool-sample--[a-f0-9]+-grp`)
			if nodepoolRegex.MatchString(value) {
				match := nodepoolRegex.FindString(value)
				replacements.ReplaceStringValue(match, "gke-containercluster-abcdef-nodepool-sample-"+uniqueID+"-grp")
			}
		}

		// Normalize default-pool instance template (e.g. gke-cluster-sample-[a-z0-9]+-default-pool-[a-f0-9]{8})
		defaultPoolTemplateRegex := regexp.MustCompile(`gke-cluster-sample-[a-z0-9]+-default-pool-[a-f0-9]{8}`)
		if defaultPoolTemplateRegex.MatchString(value) {
			match := defaultPoolTemplateRegex.FindString(value)
			replacements.ReplaceStringValue(match, "gke-cluster-sample-py3hj-default-pool-da81477a")
		}

		// Normalize nodepool-sample instance template (e.g. gke-cluster-sample-[a-z0-9]+-nodepool-sample--[a-f0-9]{8})
		instanceTemplateRegex := regexp.MustCompile(`gke-cluster-sample-[a-z0-9]+-nodepool-sample--[a-f0-9]{8}`)
		if instanceTemplateRegex.MatchString(value) {
			match := instanceTemplateRegex.FindString(value)
			replacements.ReplaceStringValue(match, "gke-cluster-sample-p-nodepool-sample--08c12cb9")
		}

		// Normalize pod range / secondary range name (e.g. gke-cluster-sample-[a-z0-9]+-pods-[a-f0-9]{8})
		podsRangeRegex := regexp.MustCompile(`gke-cluster-sample-[a-z0-9]+-pods-[a-f0-9]{8}`)
		if podsRangeRegex.MatchString(value) {
			match := podsRangeRegex.FindString(value)
			if uniqueID != "" {
				replacements.ReplaceStringValue(match, "gke-cluster-sample-"+uniqueID+"-pods-7b2f8c84")
			}
		}
	})
}

// Simple check for IPv4 address format.
func isIPv4Address(value string) bool {
	ip := net.ParseIP(value)
	return ip != nil && ip.To4() != nil
}
