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
	_ identity.IdentityV2 = &ComputeInstanceGroupIdentity{}
	_ identity.Resource   = &ComputeInstanceGroup{}
)

var ComputeInstanceGroupIdentityFormat = gcpurls.Template[ComputeInstanceGroupIdentity]("compute.googleapis.com", "projects/{project}/zones/{zone}/instanceGroups/{instanceGroup}")

// ComputeInstanceGroupIdentity is the identity of a GCP ComputeInstanceGroup resource.
// +k8s:deepcopy-gen=false
type ComputeInstanceGroupIdentity struct {
	Project       string
	Zone          string
	InstanceGroup string
}

func (i *ComputeInstanceGroupIdentity) String() string {
	return ComputeInstanceGroupIdentityFormat.ToString(*i)
}

func (i *ComputeInstanceGroupIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeInstanceGroupIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeInstanceGroup external=%q was not known (use %s): %w", ref, ComputeInstanceGroupIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeInstanceGroup external=%q was not known (use %s)", ref, ComputeInstanceGroupIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeInstanceGroupIdentity) Host() string {
	return ComputeInstanceGroupIdentityFormat.Host()
}

func (i *ComputeInstanceGroupIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/zones/%s", i.Project, i.Zone)
}

func ParseComputeInstanceGroupExternal(external string) (*ComputeInstanceGroupIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeInstanceGroup external value")
	}
	id := &ComputeInstanceGroupIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeInstanceGroupSpec(ctx context.Context, reader client.Reader, obj *ComputeInstanceGroup) (*ComputeInstanceGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	zone := obj.Spec.Zone
	if zone == "" {
		return nil, fmt.Errorf("cannot resolve zone: spec.zone is empty")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeInstanceGroupIdentity{
		Project:       projectID,
		Zone:          zone,
		InstanceGroup: resourceID,
	}
	return identity, nil
}

func (obj *ComputeInstanceGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeInstanceGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeInstanceGroupIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeInstanceGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
