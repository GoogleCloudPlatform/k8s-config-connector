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

	replacements.TransformObject("", func(m map[string]any) {
		isFilestore := false
		if m["tier"] != nil || m["fileShares"] != nil {
			isFilestore = true
		}
		if name, ok := m["name"].(string); ok && strings.Contains(name, "/locations/") && (strings.Contains(name, "/instances/") || strings.Contains(name, "/operations/")) {
			isFilestore = true
		}

		if isFilestore {
			isResponse := m["state"] != nil || m["createTime"] != nil || (m["name"] != nil && strings.Contains(m["name"].(string), "/operations/"))
			if isResponse {
				addSecurityFlavors(m)
			}
			if m["state"] != nil || m["createTime"] != nil {
				if fileShares, ok := m["fileShares"].([]any); ok && len(fileShares) > 0 {
					if firstShare, ok := fileShares[0].(map[string]any); ok {
						if capGb, ok := firstShare["capacityGb"]; ok {
							m["capacityGb"] = capGb
						}
					}
				}
				if networks, ok := m["networks"].([]any); ok {
					for _, net := range networks {
						if netMap, ok := net.(map[string]any); ok {
							netMap["network"] = "${networkID}"
						}
					}
				}
			}
			if name, ok := m["name"].(string); ok && strings.Contains(name, "/operations/") {
				if m["done"] == nil {
					m["done"] = false
				}
			}
		}
	})

	replacements.TransformObject(".response", func(m map[string]any) {
		addSecurityFlavors(m)
		isFilestore := false
		if m["tier"] != nil || m["fileShares"] != nil {
			isFilestore = true
		}

		if isFilestore {
			if fileShares, ok := m["fileShares"].([]any); ok && len(fileShares) > 0 {
				if firstShare, ok := fileShares[0].(map[string]any); ok {
					if capGb, ok := firstShare["capacityGb"]; ok {
						m["capacityGb"] = capGb
					}
				}
			}
			if networks, ok := m["networks"].([]any); ok {
				for _, net := range networks {
					if netMap, ok := net.(map[string]any); ok {
						netMap["network"] = "${networkID}"
					}
				}
			}
		}
	})

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

func addSecurityFlavors(val any) {
	switch v := val.(type) {
	case map[string]any:
		if options, ok := v["nfsExportOptions"].([]any); ok {
			for _, opt := range options {
				if optMap, ok := opt.(map[string]any); ok {
					if optMap["securityFlavors"] == nil {
						optMap["securityFlavors"] = []any{"AUTH_SYS"}
					}
				}
			}
		}
		for _, child := range v {
			addSecurityFlavors(child)
		}
	case []any:
		for _, child := range v {
			addSecurityFlavors(child)
		}
	}
}
