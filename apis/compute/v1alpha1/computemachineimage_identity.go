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
	_ identity.IdentityV2 = &ComputeMachineImageIdentity{}
	_ identity.Resource   = &ComputeMachineImage{}
)

var ComputeMachineImageIdentityFormat = gcpurls.Template[ComputeMachineImageIdentity](
	"compute.googleapis.com",
	"projects/{project}/global/machineImages/{machineImage}",
)

// ComputeMachineImageIdentity is the identity of a GCP ComputeMachineImage resource.
// +k8s:deepcopy-gen=false
type ComputeMachineImageIdentity struct {
	Project      string
	MachineImage string
}

func (i *ComputeMachineImageIdentity) String() string {
	return ComputeMachineImageIdentityFormat.ToString(*i)
}

func (i *ComputeMachineImageIdentity) FromExternal(ref string) error {
	ref = refs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeMachineImageIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of ComputeMachineImage external=%q was not known (use %s): %w", ref, ComputeMachineImageIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeMachineImage external=%q was not known (use %s)", ref, ComputeMachineImageIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeMachineImageIdentity) Host() string {
	return ComputeMachineImageIdentityFormat.Host()
}

func (i *ComputeMachineImageIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeMachineImageExternal(external string) (*ComputeMachineImageIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeMachineImage external value")
	}
	id := &ComputeMachineImageIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeMachineImageSpec(ctx context.Context, reader client.Reader, obj *ComputeMachineImage) (*ComputeMachineImageIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeMachineImageIdentity{
		Project:      projectID,
		MachineImage: resourceID,
	}
	return identity, nil
}

func (obj *ComputeMachineImage) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeMachineImageSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeMachineImageIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeMachineImage identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
