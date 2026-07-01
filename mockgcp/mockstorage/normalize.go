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

var _ mockgcpregistry.SupportsNormalization = &MockService{}

func (s *MockService) ConfigureVisitor(url string, replacements mockgcpregistry.NormalizingVisitor) {
	// Bucket
	replacements.ReplacePath(".softDeletePolicy.effectiveTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".timeCreated", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].timeCreated", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".updated", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].updated", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".softDeletePolicy.effectiveTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].softDeletePolicy.effectiveTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".autoclass.toggleTime", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].autoclass.toggleTime", PlaceholderTimestamp)
	replacements.ReplacePath(".autoclass.terminalStorageClassUpdateTime", PlaceholderTimestamp)
	replacements.ReplacePath(".items[].autoclass.terminalStorageClassUpdateTime", PlaceholderTimestamp)

	replacements.ReplacePath(".acl[].etag", "abcdef0123A")
	replacements.ReplacePath(".items[].acl[].etag", "abcdef0123A")
	replacements.ReplacePath(".defaultObjectAcl[].etag", "abcdef0123A=")
	replacements.ReplacePath(".items[].defaultObjectAcl[].etag", "abcdef0123A=")

	replacements.ReplacePath(".generation", "12345678901234")
	replacements.ReplacePath(".items[].generation", "12345678901234")

	replacements.ReplacePath(".iamConfiguration.bucketPolicyOnly.lockedTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].iamConfiguration.bucketPolicyOnly.lockedTime", mockgcpregistry.PlaceholderTimestamp)

	replacements.ReplacePath(".iamConfiguration.uniformBucketLevelAccess.lockedTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].iamConfiguration.uniformBucketLevelAccess.lockedTime", mockgcpregistry.PlaceholderTimestamp)

	// Managed Folder
	replacements.ReplacePath(".items[].createTime", mockgcpregistry.PlaceholderTimestamp)
	replacements.ReplacePath(".items[].updateTime", mockgcpregistry.PlaceholderTimestamp)

	// Anywhere Cache
	replacements.ReplacePath(".requestId", "1234")

}

func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
}
