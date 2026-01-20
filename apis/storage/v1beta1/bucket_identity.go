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

package v1beta1

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	StorageBucketIdentityFormat = gcpurls.Template[StorageBucketIdentity]("storage.googleapis.com", "projects/{project}/buckets/{bucket}")
)

// StorageBucketURLFormat is the format for the externalRef of a StorageBucket.
const StorageBucketURLFormat = "projects/{{project}}/buckets/{{bucket}}"

var _ identity.Identity = &StorageBucketIdentity{}

// +k8s:deepcopy-gen=false
type StorageBucketIdentity struct {
	Project string
	Bucket  string
}

func (i *StorageBucketIdentity) Service() string {
	return StorageBucketIdentityFormat.Host()
}

func (i *StorageBucketIdentity) String() string {
	return StorageBucketIdentityFormat.ToString(*i)
}

func (i *StorageBucketIdentity) BucketName() string {
	return i.Bucket
}

// Deprecated: prefer FromExternal
func ParseStorageBucketExternal(external string) (*StorageBucketIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("missing external value")
	}
	id := &StorageBucketIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func (i *StorageBucketIdentity) FromExternal(ref string) error {
	parsed, match, err := StorageBucketIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of StorageBucket external=%q was not known (use %s): %w", ref, StorageBucketURLFormat, err)
	}
	if !match {
		return fmt.Errorf("format of StorageBucket external=%q was not known (use %s)", ref, StorageBucketURLFormat)
	}

	*i = *parsed
	return nil
}
