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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
)

const PlaceholderTimestamp = "2024-04-01T12:34:56.123456Z"

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Bucket
	replacements.ReplacePath(".softDeletePolicy.effectiveTime", PlaceholderTimestamp)
	replacements.ReplacePath(".timeCreated", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].timeCreated", PlaceholderTimestamp)

	replacements.ReplacePath(".updated", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].updated", PlaceholderTimestamp)

	replacements.ReplacePath(".softDeletePolicy.effectiveTime", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].softDeletePolicy.effectiveTime", PlaceholderTimestamp)

	replacements.ReplacePath(".acl[].etag", "abcdef0123A")
	replacements.ReplacePath(".items[].acl[].etag", "abcdef0123A")
	replacements.ReplacePath(".defaultObjectAcl[].etag", "abcdef0123A=")
	replacements.ReplacePath(".items[].defaultObjectAcl[].etag", "abcdef0123A=")

	replacements.ReplacePath(".generation", "12345678901234")
	replacements.ReplacePath(".items[].generation", "12345678901234")

	replacements.ReplacePath(".iamConfiguration.bucketPolicyOnly.lockedTime", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].iamConfiguration.bucketPolicyOnly.lockedTime", PlaceholderTimestamp)

	replacements.ReplacePath(".iamConfiguration.uniformBucketLevelAccess.lockedTime", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].iamConfiguration.uniformBucketLevelAccess.lockedTime", PlaceholderTimestamp)

	// Managed Folder
	replacements.ReplacePath(".items[].createTime", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].updateTime", PlaceholderTimestamp)

	// Anywhere Cache
	replacements.ReplacePath(".requestId", "1234")

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
