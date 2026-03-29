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

package mockresourcemanager

import (
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"k8s.io/klog/v2"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	if !strings.Contains(url, "cloudresourcemanager.googleapis.com") {
		return
	}

	// Normalization for TagKeys and TagValues
	replacements.ReplacePath(".createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".etag", "abcdef0123A=")

	replacements.ReplacePath(".response.createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.etag", "abcdef0123A=")

	replacements.ReplacePath(".tagKeys[].createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".tagKeys[].updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".tagKeys[].etag", "abcdef0123A=")

	replacements.ReplacePath(".tagValues[].createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".tagValues[].updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".tagValues[].etag", "abcdef0123A=")

	replacements.ReplacePath(".response.tagKeys[].createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.tagKeys[].updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.tagKeys[].etag", "abcdef0123A=")

	replacements.ReplacePath(".response.tagValues[].createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.tagValues[].updateTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".response.tagValues[].etag", "abcdef0123A=")
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isCloudResourceManagerAPI(event) {
		return
	}

	name := ""
	event.VisitResponseStringValues(func(path string, value string) {
		if path == ".name" || path == ".response.name" {
			name = value
		}
		if path == ".projectNumber" {
			replacements.ReplaceStringValue(value, "${projectNumber}")
		}
	})

	tokens := strings.Split(name, "/")
	if len(tokens) == 2 && tokens[0] == "tagKeys" {
		if tokens[1] == "namespaced" {
			// This is actually a search operation: https://cloud.google.com/resource-manager/reference/rest/v3/tagKeys/getNamespaced
		} else {
			replacements.ReplaceStringValue(tokens[1], "${tagKeyID}")
		}
	}
	if len(tokens) == 2 && tokens[0] == "tagValues" {
		if tokens[1] == "namespaced" {
			// This is actually a search operation: https://cloud.google.com/resource-manager/reference/rest/v3/tagValues/getNamespaced
		} else {
			replacements.ReplaceStringValue(tokens[1], "${tagValueID}")
		}
	}
	if len(tokens) == 2 && tokens[0] == "tagBindings" {
		replacements.ReplaceStringValue(tokens[1], "${tagBindingID}")
	}
}

// isCloudResourceManagerAPI returns true if this is a cloud resource manager URL
func isCloudResourceManagerAPI(event mockgcpregistry.Event) bool {
	u, err := url.Parse(event.URL())
	if err != nil {
		klog.Fatalf("cannot parse URL %q", event.URL())
	}
	switch u.Host {
	case "cloudresourcemanager.googleapis.com":
		return true
	}
	return false
}
