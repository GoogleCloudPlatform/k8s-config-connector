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

package mockcompute

import (
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"k8s.io/klog/v2"
)

const PlaceholderTimestamp = "2024-04-01T12:34:56.123456Z"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// General
	replacements.ReplacePath(".creationTimestamp", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].creationTimestamp", PlaceholderTimestamp)

	// Addresses
	replacements.ReplacePath(".labelFingerprint", "abcdef0123A=")
	replacements.ReplacePath(".items[].labelFingerprint", "abcdef0123A=")

	replacements.ReplacePath(".address", "8.8.8.8")
	replacements.ReplacePath(".items[].address", "8.8.8.8")

	replacements.SortSlice(".subnetworks")

	// BackendBuckets

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	addReplacement := func(kind, value string) {
		// Never replace empty values
		if value == "" {
			klog.Fatalf("ignoring empty replacement kind=%q value=%q", kind, value)
			return
		}

		// TODO: unify with placeholderForGCPResource

		placeholder := "${" + strings.TrimSuffix(kind, "s") + "ID}"
		if strings.HasSuffix(kind, "ies") {
			placeholder = "${" + strings.TrimSuffix(kind, "ies") + "yID}"
		}
		switch kind {
		case "addresses":
			placeholder = "${addressID}"
		case "projects":
			// Specially handled
			return
		case "regions":
			// Specially handled (not replaced)
			return
		}

		if value == "default" {
			// Don't replace, "default" is a well-known value used for both subnetwork and network
			// We could instead do something like this:  replacements.ReplaceStringValue(kind + "/" + v, kind + "/" + placeholder)
		} else {
			replacements.ReplaceStringValue(value, placeholder)
		}
	}

	visitLink := func(link string) {
		s, _, _ := strings.Cut(link, "?")

		s = strings.TrimPrefix(s, "https://compute.googleapis.com/")
		s = strings.TrimPrefix(s, "https://www.googleapis.com/")
		s = strings.TrimPrefix(s, "/")
		s = strings.TrimPrefix(s, "compute/v1/")

		s = strings.TrimSuffix(s, "/")

		tokens := strings.Split(s, "/")
		for i := 0; i+1 < len(tokens); i += 2 {
			kind := tokens[i]
			if kind == "global" {
				// Special singleton token
				i--
				continue
			}
			value := tokens[i+1]

			addReplacement(kind, value)
		}
	}

	// Extract IDs from normal resource URLs
	if isComputeAPI(event) && !isGetOperation(event) {
		visitLink(event.URL())
	}

	// Extract IDs from operations (via the targetId and targetLink fields)
	if isComputeAPI(event) && isGetOperation(event) {
		targetLink := ""
		targetId := ""

		event.VisitResponseStringValues(func(path string, value string) {
			switch path {
			case ".targetLink":
				targetLink = value
			case ".targetId":
				targetId = value
			}
		})

		if targetLink != "" && targetId != "" {
			tokens := strings.Split(targetLink, "/")
			n := len(tokens)
			if n >= 2 {
				kind := tokens[n-2]

				// We _should_ differentiate between ID and number.
				// But this causes too many diffs right now.

				addReplacement(kind, targetId)
				addReplacement(kind, tokens[n-1])
			}

			visitLink(targetLink)
		}
	}
}

// isGetOperation returns true if this is an operation poll request
func isGetOperation(event mockgcpregistry.Event) bool {
	u := event.URL()
	// A normal GET poll
	if event.Method() == "GET" && strings.Contains(u, "/operations/") {
		return true
	}
	// A call to the /wait endpoint
	if event.Method() == "POST" && strings.Contains(u, "/operations/") && strings.Contains(u, "/wait") {
		return true
	}
	// A GRPC call
	if u == "/google.longrunning.Operations/GetOperation" {
		return true
	}
	return false
}

// isComputeAPI returns true if this is a compute URL
func isComputeAPI(event mockgcpregistry.Event) bool {
	u, err := url.Parse(event.URL())
	if err != nil {
		klog.Fatalf("cannot parse URL %q", event.URL())
	}
	switch u.Host {
	case "compute.googleapis.com":
		return true
	}
	return false
}
