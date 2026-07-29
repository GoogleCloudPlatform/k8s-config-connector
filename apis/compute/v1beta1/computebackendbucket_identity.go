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

package v1beta1

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &ComputeBackendBucketIdentity{}
)

var ComputeBackendBucketIdentityFormat = gcpurls.Template[ComputeBackendBucketIdentity](
	"compute.googleapis.com",
	"projects/{project}/regions/{region}/backendBuckets/{backendBucket}",
)

// ComputeBackendBucketIdentity is the identity of a GCP ComputeBackendBucket resource.
// +k8s:deepcopy-gen=false
type ComputeBackendBucketIdentity struct {
	Project       string
	region        string
	BackendBucket string
}

func (i *ComputeBackendBucketIdentity) String() string {
	return ComputeBackendBucketIdentityFormat.ToString(*i)
}

func (i *ComputeBackendBucketIdentity) FromExternal(ref string) error {
	ref = refs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeBackendBucketIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeBackendBucket external=%q was not known (use %s): %w", ref, ComputeBackendBucketIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeBackendBucket external=%q was not known (use %s)", ref, ComputeBackendBucketIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeBackendBucketIdentity) Host() string {
	return ComputeBackendBucketIdentityFormat.Host()
}
