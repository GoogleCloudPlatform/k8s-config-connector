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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeTargetPoolIdentity{}
	_ identity.Resource   = &ComputeTargetPool{}
)

var ComputeTargetPoolIdentityFormat = gcpurls.Template[ComputeTargetPoolIdentity](
	"compute.googleapis.com",
	"projects/{project}/regions/{region}/targetPools/{targetpool}",
)

// ComputeTargetPoolIdentity is the identity of a GCP ComputeTargetPool resource.
// +k8s:deepcopy-gen=false
type ComputeTargetPoolIdentity struct {
	Project    string
	Region     string
	TargetPool string
}

func (i *ComputeTargetPoolIdentity) String() string {
	return ComputeTargetPoolIdentityFormat.ToString(*i)
}

func (i *ComputeTargetPoolIdentity) FromExternal(ref string) error {
	trimmedRef := refs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeTargetPoolIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeTargetPool external=%q was not known (use %s): %w", ref, ComputeTargetPoolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeTargetPool external=%q was not known (use %s)", ref, ComputeTargetPoolIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeTargetPoolIdentity) Host() string {
	return ComputeTargetPoolIdentityFormat.Host()
}

func (i *ComputeTargetPoolIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
}

func ParseComputeTargetPoolExternal(external string) (*ComputeTargetPoolIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeTargetPool external value")
	}
	id := &ComputeTargetPoolIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeTargetPoolSpec(ctx context.Context, reader client.Reader, obj *ComputeTargetPool) (*ComputeTargetPoolIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeTargetPoolIdentity{
		Project:    projectID,
		Region:     location,
		TargetPool: resourceID,
	}
	return identity, nil
}

func (obj *ComputeTargetPool) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeTargetPoolSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against the status value, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &ComputeTargetPoolIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeTargetPool identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
