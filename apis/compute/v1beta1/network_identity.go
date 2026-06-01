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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refcommon "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NetworkIdentity{}
	_ identity.Resource   = &ComputeNetwork{}
)

var ComputeNetworkIdentityFormat = gcpurls.Template[NetworkIdentity]("compute.googleapis.com", "projects/{project}/global/networks/{network}")

// +k8s:deepcopy-gen=false
type NetworkIdentity struct {
	Project string
	Network string
}

func (i *NetworkIdentity) Parent() parent.ProjectParent {
	return parent.ProjectParent{ProjectID: i.Project}
}

func (i *NetworkIdentity) ID() string {
	return i.Network
}

func (i *NetworkIdentity) String() string {
	return ComputeNetworkIdentityFormat.ToString(*i)
}

func (i *NetworkIdentity) FromExternal(ref string) error {
	parsed, match, err := ComputeNetworkIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeNetwork external=%q was not known (use %s): %w", ref, ComputeNetworkIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeNetwork external=%q was not known (use %s)", ref, ComputeNetworkIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkIdentity) Host() string {
	return ComputeNetworkIdentityFormat.Host()
}

func ParseComputeNetworkExternal(external string) (*NetworkIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeNetwork external value")
	}
	trimmedExternal := refcommon.FixStaleComputeExternalFormat(external)
	id := &NetworkIdentity{}
	if err := id.FromExternal(trimmedExternal); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeNetworkSpec(ctx context.Context, reader client.Reader, obj client.Object) (*NetworkIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NetworkIdentity{
		Project: projectID,
		Network: resourceID,
	}
	return identity, nil
}

func (obj *ComputeNetwork) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeNetworkSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &NetworkIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeNetwork identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
