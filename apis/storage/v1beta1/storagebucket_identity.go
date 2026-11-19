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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &StorageBucketIdentity{}
	_ identity.Resource   = &StorageBucket{}
)

var (
	StorageBucketIdentityFormat = gcpurls.Template[StorageBucketIdentity]("storage.googleapis.com", "projects/{project}/buckets/{bucket}")
)

// StorageBucketURLFormat is the format for the externalRef of a StorageBucket.
const StorageBucketURLFormat = "projects/{{project}}/buckets/{{bucket}}"

// StorageBucketIdentity is the identity of a GCP StorageBucket resource.
// +k8s:deepcopy-gen=false
type StorageBucketIdentity struct {
	Project string
	Bucket  string
}

func (i *StorageBucketIdentity) Host() string {
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
	if strings.HasPrefix(ref, "gs://") {
		bucket := strings.TrimPrefix(ref, "gs://")
		bucket = strings.TrimSuffix(bucket, "/")
		if !isValidGCSBucketName(bucket) {
			return fmt.Errorf("format of StorageBucket external=%q was not known (use %s)", ref, StorageBucketURLFormat)
		}
		i.Project = "_"
		i.Bucket = bucket
		return nil
	}

	for _, prefix := range []string{
		"https://www.googleapis.com/storage/v1/b/",
		"http://www.googleapis.com/storage/v1/b/",
		"//www.googleapis.com/storage/v1/b/",
		"www.googleapis.com/storage/v1/b/",
		"https://storage.googleapis.com/storage/v1/b/",
		"http://storage.googleapis.com/storage/v1/b/",
		"//storage.googleapis.com/storage/v1/b/",
		"storage.googleapis.com/storage/v1/b/",
	} {
		if strings.HasPrefix(ref, prefix) {
			bucket := strings.TrimPrefix(ref, prefix)
			bucket = strings.TrimSuffix(bucket, "/")
			if !isValidGCSBucketName(bucket) {
				return fmt.Errorf("format of StorageBucket external=%q was not known (use %s)", ref, StorageBucketURLFormat)
			}
			i.Project = "_"
			i.Bucket = bucket
			return nil
		}
	}

	if !strings.Contains(ref, "/") {
		if !isValidGCSBucketName(ref) {
			return fmt.Errorf("format of StorageBucket external=%q was not known (use %s)", ref, StorageBucketURLFormat)
		}
		i.Project = "_"
		i.Bucket = ref
		return nil
	}

	parsed, match, err := StorageBucketIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of StorageBucket external=%q was not known (use %s): %w", ref, StorageBucketURLFormat, err)
	}
	if !match || !isValidGCSBucketName(parsed.Bucket) {
		return fmt.Errorf("format of StorageBucket external=%q was not known (use %s)", ref, StorageBucketURLFormat)
	}

	*i = *parsed
	return nil
}

const maxGCSBucketNameLength = 222

func isValidGCSBucketName(bucket string) bool {
	if len(bucket) == 0 || len(bucket) > maxGCSBucketNameLength {
		return false
	}
	if !gcpurls.IsValidSegment(bucket) || strings.Contains(bucket, "..") {
		return false
	}
	for idx := 0; idx < len(bucket); idx++ {
		c := bucket[idx]
		isLowerAlphaNum := (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
		isAllowedPunct := c == '-' || c == '_' || c == '.'
		if !isLowerAlphaNum && !isAllowedPunct {
			return false
		}
		if (idx == 0 || idx == len(bucket)-1) && !isLowerAlphaNum {
			return false
		}
	}
	return true
}

func getIdentityFromStorageBucketSpec(ctx context.Context, reader client.Reader, obj *StorageBucket) (*StorageBucketIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &StorageBucketIdentity{
		Project: projectID,
		Bucket:  resourceID,
	}
	if err := StorageBucketIdentityFormat.Validate(*identity); err != nil {
		return nil, err
	}
	return identity, nil
}

func (obj *StorageBucket) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return getIdentityFromStorageBucketSpec(ctx, reader, obj)
}
