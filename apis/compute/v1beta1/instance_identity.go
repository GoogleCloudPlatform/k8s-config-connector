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
	_ identity.IdentityV2 = &InstanceIdentity{}
	_ identity.Resource   = &ComputeInstance{}
)

var InstanceIdentityFormat = gcpurls.Template[InstanceIdentity]("compute.googleapis.com", "projects/{project}/zones/{zone}/instances/{instance}")

// InstanceIdentity is the identity of a GCP ComputeInstance resource.
// +k8s:deepcopy-gen=false
type InstanceIdentity struct {
	Project  string
	Zone     string
	Instance string
}

func (i *InstanceIdentity) String() string {
	return InstanceIdentityFormat.ToString(*i)
}

func (i *InstanceIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := InstanceIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeInstance external=%q was not known (use %s): %w", ref, InstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeInstance external=%q was not known (use %s)", ref, InstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *InstanceIdentity) Host() string {
	return InstanceIdentityFormat.Host()
}

func (i *InstanceIdentity) ParentString() string {
	return "projects/" + i.Project + "/zones/" + i.Zone
}

func ParseInstanceExternal(external string) (*InstanceIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeInstance external value")
	}
	id := &InstanceIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromInstanceSpec(ctx context.Context, reader client.Reader, obj *ComputeInstance) (*InstanceIdentity, error) {
	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	zone := common.ValueOf(obj.Spec.Zone)
	if zone == "" {
		return nil, fmt.Errorf("cannot resolve zone")
	}

	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	identity := &InstanceIdentity{
		Project:  projectID,
		Zone:     zone,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *ComputeInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &InstanceIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeInstance identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
