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

package mockstorage

import (
	// "fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Bucket
	// fmt.Printf("VRX CONFIG PATH... VRO %w", url)
	replacements.ReplacePath(".softDeletePolicy.effectiveTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".timeCreated", "2025-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".generation", "786786")
	replacements.ReplacePath(".updated", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".items[].timeCreated", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".items[].updated", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".acl[].etag", "abcdef0123A")
	replacements.ReplacePath(".defaultObjectAcl[].etag", "abcdef0123A=")

	// Managed Folder
	replacements.ReplacePath(".items[].createTime", "2024-04-01T12:34:56.123456Z")
	replacements.ReplacePath(".items[].updateTime", "2024-04-01T12:34:56.123456Z")

	// Anywhere Cache
	replacements.ReplacePath(".requestId", "1234")

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
	// No-op for now
	// fmt.Printf("\n\nIn BUCKET Previsit %w\n\n", event)
	// fmt.Printf("\n\nevent URL %s\n\n", event.URL())
	// fmt.Printf("\n\nevent method %s\n\n", event.Method())
	// event.VisitResponseStringValues(func(path string, value string) {
	// 	fmt.Printf("event path: %s\n", path)
	// 	fmt.Printf("event value: %s\n\n", value)
	// })
	// fmt.Printf("\n\n----------------------------------\n\n")
	// event.VisitRequestStringValues(func(path string, value string) {
	// 	fmt.Printf("event path: %s\n", path)
	// 	fmt.Printf("event value: %s\n\n", value)
	// })
	// fmt.Printf("\n\n----------------------------------\n\n")
}
