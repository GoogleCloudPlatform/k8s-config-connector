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

package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &StorageBucketObjectIdentity{}
)

var (
	StorageBucketObjectIdentityFormat = gcpurls.Template[StorageBucketObjectIdentity]("storage.googleapis.com", "projects/{project}/buckets/{bucket}/objects/{object}")
)

// StorageBucketObjectIdentity is the identity of a GCP StorageBucketObject resource.
// +k8s:deepcopy-gen=false
type StorageBucketObjectIdentity struct {
	Project string
	Bucket  string
	Object  string
}

func (i *StorageBucketObjectIdentity) Host() string {
	return StorageBucketObjectIdentityFormat.Host()
}

func (i *StorageBucketObjectIdentity) String() string {
	return StorageBucketObjectIdentityFormat.ToString(*i)
}

func (i *StorageBucketObjectIdentity) FromExternal(ref string) error {
	// Strip optional scheme
	ref = strings.TrimPrefix(ref, "https:")
	ref = strings.TrimPrefix(ref, "http:")
	ref = strings.TrimPrefix(ref, "//")

	// Check and strip host
	host := "storage.googleapis.com"
	if strings.HasPrefix(ref, host+"/") {
		ref = strings.TrimPrefix(ref, host)
	}

	ref = strings.Trim(ref, "/")
	tokens := strings.Split(ref, "/")

	// Must have at least 6 tokens: projects/{project}/buckets/{bucket}/objects/{object...}
	if len(tokens) < 6 || tokens[0] != "projects" || tokens[2] != "buckets" || tokens[4] != "objects" {
		return fmt.Errorf("format of StorageBucketObject external=%q was not known (use projects/{project}/buckets/{bucket}/objects/{object})", ref)
	}

	i.Project = tokens[1]
	i.Bucket = tokens[3]
	i.Object = strings.Join(tokens[5:], "/")
	return nil
}

func (i *StorageBucketObjectIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/buckets/%s", i.Project, i.Bucket)
}
