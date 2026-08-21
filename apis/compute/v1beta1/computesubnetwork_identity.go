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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeSubnetworkIdentity{}
	_ identity.Resource   = &ComputeSubnetwork{}
)

var ComputeSubnetworkIdentityFormat = gcpurls.Template[ComputeSubnetworkIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/subnetworks/{subnetwork}")

// ComputeSubnetworkIdentity is the identity of a GCP ComputeSubnetwork resource.
// +k8s:deepcopy-gen=false
type ComputeSubnetworkIdentity struct {
	Project    string
	Region     string
	Subnetwork string
}

func (i *ComputeSubnetworkIdentity) String() string {
	return ComputeSubnetworkIdentityFormat.ToString(*i)
}

func (i *ComputeSubnetworkIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeSubnetworkIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeSubnetwork external=%q was not known (use %s): %w", ref, ComputeSubnetworkIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeSubnetwork external=%q was not known (use %s)", ref, ComputeSubnetworkIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeSubnetworkIdentity) Host() string {
	return ComputeSubnetworkIdentityFormat.Host()
}

func (i *ComputeSubnetworkIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
}

func ParseComputeSubnetworkExternal(external string) (*ComputeSubnetworkIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeSubnetwork external value")
	}
	id := &ComputeSubnetworkIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeSubnetworkSpec(ctx context.Context, reader client.Reader, obj *ComputeSubnetwork) (*ComputeSubnetworkIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	region := common.ValueOf(obj.Spec.Region)
	if region == "" {
		return nil, fmt.Errorf("cannot resolve region: spec.region is empty")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeSubnetworkIdentity{
		Project:    projectID,
		Region:     region,
		Subnetwork: resourceID,
	}
	return identity, nil
}

func (obj *ComputeSubnetwork) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeSubnetworkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeSubnetworkIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeSubnetwork identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
