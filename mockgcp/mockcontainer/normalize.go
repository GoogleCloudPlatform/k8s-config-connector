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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

const PlaceholderEtag = "abcdef0123A="

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Cluster
	replacements.ReplacePath(".nodePools[].etag", PlaceholderEtag)
	replacements.ReplacePath(".clusterIpv4Cidr", "10.48.0.0/14")
	replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4Cidr", "10.48.0.0/14")
	replacements.ReplacePath(".ipAllocationPolicy.clusterIpv4CidrBlock", "10.48.0.0/14")
	replacements.ReplacePath(".masterAuth.clusterCaCertificate", "removed-for-normalization")
	replacements.ReplacePath(".maintenancePolicy.resourceVersion", PlaceholderEtag)
	replacements.ReplacePath(".ipAllocationPolicy.defaultPodIpv4RangeUtilization", 0.1)
	replacements.ReplacePath(".nodePools[].networkConfig.podIpv4RangeUtilization", 0.1)

	replacements.ReplacePath(".controlPlaneEndpointsConfig.ipEndpointsConfig.privateEndpoint", "10.128.1.1")
	replacements.ReplacePath(".controlPlaneEndpointsConfig.ipEndpointsConfig.publicEndpoint", "35.1.1.1")
	replacements.ReplacePath(".controlPlaneEndpointsConfig.dnsEndpointConfig.endpoint", "gke-hash-projectnumber.location.gke.goog")
	replacements.ReplacePath(".privateClusterConfig.privateEndpoint", "10.128.1.1")
	replacements.ReplacePath(".privateClusterConfig.publicEndpoint", "35.1.1.1")
	replacements.ReplacePath(".endpoint", "35.1.1.1")

	replacements.SortSlice(".monitoringConfig.componentConfig.enableComponents")

	// NodePool
	replacements.ReplacePath(".etag", PlaceholderEtag)
	replacements.ReplacePath(".networkConfig.podIpv4CidrBlock", "10.4.0.0/14")
	replacements.ReplacePath(".nodePools[].networkConfig.podIpv4CidrBlock", "10.4.0.0/14")

	replacements.TransformString(func(path string, s string) string {
		if path == ".networkConfig.podRange" || path == ".nodePools[].networkConfig.podRange" {
			tokens := strings.Split(s, "-")
			n := len(tokens)
			if len(tokens) > 3 && tokens[0] == "gke" && tokens[n-2] == "pods" {
				tokens[n-1] = "{hash}"
			}
			s = strings.Join(tokens, "-")
		}
		if path == ".ipAllocationPolicy.clusterSecondaryRangeName" || path == ".nodePools[].ipAllocationPolicy.clusterSecondaryRangeName" {
			tokens := strings.Split(s, "-")
			n := len(tokens)
			if len(tokens) > 3 && tokens[0] == "gke" && tokens[n-2] == "pods" {
				tokens[n-1] = "{hash}"
			}
			s = strings.Join(tokens, "-")
		}
		return s
	})

	replacements.TransformString(func(path string, s string) string {
		if path == ".instanceGroupUrls[]" || path == ".nodePools[].instanceGroupUrls[]" {
			tokens := strings.Split(s, "/")
			if len(tokens) == 11 {
				s := tokens[10]
				if strings.HasPrefix(s, "gke-") && strings.HasSuffix(s, "-grp") {
					tokens[10] = "gke-{normalized}-grp"
				}
			}
			s = strings.Join(tokens, "/")
		}
		return s
	})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// No-op for now
}
