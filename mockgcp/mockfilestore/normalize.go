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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Instance
	replacements.ReplacePath(".networks[].reservedIpRange", "10.20.30.0/24")
	replacements.ReplacePath(".networks[].ipAddresses", []string{"10.20.30.1"})
	replacements.ReplacePath(".response.networks[].reservedIpRange", "10.20.30.0/24")
	replacements.ReplacePath(".response.networks[].ipAddresses", []string{"10.20.30.1"})
}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}

func (s *MockService) ConfigureKRMObjectVisitor(u *unstructured.Unstructured, replacements mockgcpregistry.NormalizingVisitor) {
}
