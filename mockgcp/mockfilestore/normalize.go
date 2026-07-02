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

package mockfilestore

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "file.googleapis.com") {
		return
	}

	// Instance
	replacements.ReplacePath(".networks[].reservedIpRange", "10.20.30.0/24")
	replacements.ReplacePath(".networks[].ipAddresses", []string{"10.20.30.1"})
	replacements.ReplacePath(".response.networks[].reservedIpRange", "10.20.30.0/24")
	replacements.ReplacePath(".response.networks[].ipAddresses", []string{"10.20.30.1"})

	replacements.ReplacePath(".etag", "abcdef0123A=")
	replacements.ReplacePath(".response.etag", "abcdef0123A=")

	replacements.ReplaceStringValue("type.googleapis.com/google.cloud.filestore.v1beta1.Instance", "type.googleapis.com/google.cloud.filestore.v1.Instance")

	replacements.RemovePath(".maxCapacityGb")
	replacements.RemovePath(".maxShareCount")
	replacements.RemovePath(".performanceLimits")
	replacements.RemovePath(".satisfiesPzi")
	replacements.RemovePath(".satisfiesPzs")
	replacements.RemovePath(".capacityStepSizeGb")

	replacements.RemovePath(".response.maxCapacityGb")
	replacements.RemovePath(".response.maxShareCount")
	replacements.RemovePath(".response.performanceLimits")
	replacements.RemovePath(".response.satisfiesPzi")
	replacements.RemovePath(".response.satisfiesPzs")
	replacements.RemovePath(".response.capacityStepSizeGb")

	replacements.TransformObject(".error", func(m map[string]any) {
		delete(m, "errors")
		if message, ok := m["message"].(string); ok {
			if strings.Contains(message, "not found") {
				m["message"] = "Resource 'projects/${projectId}/locations/us-central1-c/instances/filestoreinstance-${uniqueId}' was not found"
			}
		}
	})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
