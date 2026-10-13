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

package mockoslogin

import (
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "oslogin.googleapis.com") {
		return
	}
	replacements.ReplacePath(".fingerprint", "abcdef0123A=")
	replacements.ReplacePath(".response.fingerprint", "abcdef0123A=")
	replacements.TransformObject(".loginProfile.sshPublicKeys", func(m map[string]any) {
		for k, v := range m {
			if k != "abcdef0123A=" {
				delete(m, k)
				m["abcdef0123A="] = v
			}
		}
	})
	replacements.TransformObject(".response.loginProfile.sshPublicKeys", func(m map[string]any) {
		for k, v := range m {
			if k != "abcdef0123A=" {
				delete(m, k)
				m["abcdef0123A="] = v
			}
		}
	})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(event.URL(), "oslogin.googleapis.com") {
		return
	}
	event.VisitResponseStringValues(func(path string, value string) {
		if strings.HasSuffix(path, "fingerprint") {
			replacements.ReplaceStringValue(value, "abcdef0123A=")
		}
		if strings.HasPrefix(value, "users/") && strings.Contains(value, "/sshPublicKeys/") {
			tokens := strings.Split(value, "/")
			if len(tokens) == 4 && tokens[0] == "users" && tokens[2] == "sshPublicKeys" {
				replacements.ReplaceStringValue(tokens[3], "abcdef0123A=")
			}
		}
	})
}
