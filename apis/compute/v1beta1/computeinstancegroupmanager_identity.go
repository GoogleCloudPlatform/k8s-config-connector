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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeInstanceGroupManagerIdentity{}
	_ identity.Resource   = &ComputeInstanceGroupManager{}
)

var ZonalComputeInstanceGroupManagerIdentityFormat = gcpurls.Template[ComputeInstanceGroupManagerIdentity]("compute.googleapis.com", "projects/{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}")
var RegionalComputeInstanceGroupManagerIdentityFormat = gcpurls.Template[ComputeInstanceGroupManagerIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/instanceGroupManagers/{instanceGroupManager}")

// ComputeInstanceGroupManagerIdentity is the identity of a GCP ComputeInstanceGroupManager resource.
// +k8s:deepcopy-gen=false
type ComputeInstanceGroupManagerIdentity struct {
	Project              string
	Zone                 string
	Region               string
	InstanceGroupManager string
}

func (i *ComputeInstanceGroupManagerIdentity) IsZonal() bool {
	return i.Zone != ""
}

func (i *ComputeInstanceGroupManagerIdentity) IsRegional() bool {
	return i.Region != ""
}

func (i *ComputeInstanceGroupManagerIdentity) Validate() error {
	if i.Zone == "" && i.Region == "" {
		return fmt.Errorf("ComputeInstanceGroupManager identity must have either a Zone or a Region")
	}
	if i.Zone != "" && i.Region != "" {
		return fmt.Errorf("ComputeInstanceGroupManager identity cannot have both a Zone and a Region")
	}
	return nil
}

func (i *ComputeInstanceGroupManagerIdentity) String() string {
	if i.Zone != "" {
		return ZonalComputeInstanceGroupManagerIdentityFormat.ToString(*i)
	}
	return RegionalComputeInstanceGroupManagerIdentityFormat.ToString(*i)
}

func (i *ComputeInstanceGroupManagerIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	if parsed, match, err := ZonalComputeInstanceGroupManagerIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := RegionalComputeInstanceGroupManagerIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeInstanceGroupManager external=%q was not known (use %s or %s)", ref, ZonalComputeInstanceGroupManagerIdentityFormat.CanonicalForm(), RegionalComputeInstanceGroupManagerIdentityFormat.CanonicalForm())
}

func (i *ComputeInstanceGroupManagerIdentity) Host() string {
	return ZonalComputeInstanceGroupManagerIdentityFormat.Host()
}

func (i *ComputeInstanceGroupManagerIdentity) ParentString() string {
	if i.Zone != "" {
		return fmt.Sprintf("projects/%s/zones/%s", i.Project, i.Zone)
	}
	return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
}

func ParseComputeInstanceGroupManagerExternal(external string) (*ComputeInstanceGroupManagerIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeInstanceGroupManager external value")
	}
	id := &ComputeInstanceGroupManagerIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeInstanceGroupManagerSpec(ctx context.Context, reader client.Reader, obj *ComputeInstanceGroupManager) (*ComputeInstanceGroupManagerIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := ""
	if obj.Spec.Location != nil {
		location = *obj.Spec.Location
	}
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location: spec.location is empty or nil")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeInstanceGroupManagerIdentity{
		Project:              projectID,
		InstanceGroupManager: resourceID,
	}

	if len(strings.Split(location, "-")) == 3 {
		identity.Zone = location
	} else {
		identity.Region = location
	}

	return identity, nil
}

func (obj *ComputeInstanceGroupManager) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeInstanceGroupManagerSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeInstanceGroupManagerIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeInstanceGroupManager identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
