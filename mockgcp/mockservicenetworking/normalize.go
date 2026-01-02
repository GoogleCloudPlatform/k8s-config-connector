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

package mockservicenetworking

import (
	"net/url"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"k8s.io/klog/v2"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	if !isServiceNetworkingAPI(event) {
		return
	}

	// Recognize unusual operation ID format
	{
		name, _ := event.GetResponseStringValue(".name")
		tokens := strings.Split(name, "/")

		if len(tokens) == 2 && tokens[0] == "operations" {
			replacements.ReplaceStringValue(tokens[1], "${operationID}")
		}
	}
}

// isServiceNetworkingAPI returns true if this is a service networking URL
func isServiceNetworkingAPI(event mockgcpregistry.Event) bool {
	u, err := url.Parse(event.URL())
	if err != nil {
		klog.Fatalf("cannot parse URL %q", event.URL())
	}
	switch u.Host {
	case "servicenetworking.googleapis.com":
		return true
	}
	return false
}
