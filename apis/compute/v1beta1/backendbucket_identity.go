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
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeBackendBucketIdentity{}
	_ identity.Resource   = &ComputeBackendBucket{}
)

var ComputeBackendBucketIdentityFormat = gcpurls.Template[ComputeBackendBucketIdentity]("compute.googleapis.com", "projects/{project}/global/backendBuckets/{backendBucket}")

// +k8s:deepcopy-gen=false
type ComputeBackendBucketIdentity struct {
	Project       string
	BackendBucket string
}

func (i *ComputeBackendBucketIdentity) String() string {
	return ComputeBackendBucketIdentityFormat.ToString(*i)
}

func (i *ComputeBackendBucketIdentity) FromExternal(ref string) error {
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

func getIdentityFromComputeBackendBucketSpec(ctx context.Context, reader client.Reader, obj client.Object) (*ComputeBackendBucketIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeBackendBucketIdentity{
		Project:       projectID,
		BackendBucket: resourceID,
	}
	return identity, nil
}

func (obj *ComputeBackendBucket) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeBackendBucketSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.SelfLink)
	if externalRef != "" {
		// Remove the service prefix if present
		if strings.HasPrefix(externalRef, "https://") {
			parts := strings.SplitN(externalRef, "/projects/", 2)
			if len(parts) == 2 {
				externalRef = "projects/" + parts[1]
			}
		}

		// Validate desired with actual
		statusIdentity := &ComputeBackendBucketIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeBackendBucket identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
