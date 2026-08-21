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
	_ identity.IdentityV2 = &ComputeInstanceIdentity{}
	_ identity.Resource   = &ComputeInstance{}
)

var ComputeInstanceIdentityFormat = gcpurls.Template[ComputeInstanceIdentity]("compute.googleapis.com", "projects/{project}/zones/{zone}/instances/{instance}")

// ComputeInstanceIdentity is the identity of a GCP ComputeInstance resource.
// +k8s:deepcopy-gen=false
type ComputeInstanceIdentity struct {
	Project  string
	Zone     string
	Instance string
}

func (i *ComputeInstanceIdentity) String() string {
	return ComputeInstanceIdentityFormat.ToString(*i)
}

func (i *ComputeInstanceIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeInstanceIdentityFormat.Parse(trimmedRef)
	if err != nil {
		return fmt.Errorf("format of ComputeInstance external=%q was not known (use %s): %w", ref, ComputeInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of ComputeInstance external=%q was not known (use %s)", ref, ComputeInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *ComputeInstanceIdentity) Host() string {
	return ComputeInstanceIdentityFormat.Host()
}

func (i *ComputeInstanceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/zones/%s", i.Project, i.Zone)
}

func ParseComputeInstanceExternal(external string) (*ComputeInstanceIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeInstance external value")
	}
	id := &ComputeInstanceIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeInstanceSpec(ctx context.Context, reader client.Reader, obj *ComputeInstance) (*ComputeInstanceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	zone := common.ValueOf(obj.Spec.Zone)
	if zone == "" {
		return nil, fmt.Errorf("cannot resolve zone: spec.zone is empty")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeInstanceIdentity{
		Project:  projectID,
		Zone:     zone,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *ComputeInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeInstanceIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeInstance identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
