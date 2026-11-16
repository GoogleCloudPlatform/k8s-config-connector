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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeInstantSnapshotGroupIdentity{}
	_ identity.Resource   = &ComputeInstantSnapshotGroup{}
)

var ZonalComputeInstantSnapshotGroupIdentityFormat = gcpurls.Template[ComputeInstantSnapshotGroupIdentity](
	"compute.googleapis.com",
	"projects/{project}/zones/{zone}/instantSnapshotGroups/{instantSnapshotGroup}",
)
var RegionalComputeInstantSnapshotGroupIdentityFormat = gcpurls.Template[ComputeInstantSnapshotGroupIdentity](
	"compute.googleapis.com",
	"projects/{project}/regions/{region}/instantSnapshotGroups/{instantSnapshotGroup}",
)

// ComputeInstantSnapshotGroupIdentity is the identity of a GCP ComputeInstantSnapshotGroup resource.
// +k8s:deepcopy-gen=false
type ComputeInstantSnapshotGroupIdentity struct {
	Project              string
	Zone                 string
	Region               string
	InstantSnapshotGroup string
}

func (i *ComputeInstantSnapshotGroupIdentity) IsZonal() bool {
	return i.Zone != ""
}

func (i *ComputeInstantSnapshotGroupIdentity) IsRegional() bool {
	return i.Region != ""
}

func (i *ComputeInstantSnapshotGroupIdentity) Validate() error {
	if i.Zone == "" && i.Region == "" {
		return fmt.Errorf("ComputeInstantSnapshotGroup identity must have either a Zone or a Region")
	}
	if i.Zone != "" && i.Region != "" {
		return fmt.Errorf("ComputeInstantSnapshotGroup identity cannot have both a Zone and a Region")
	}
	return nil
}

func (i *ComputeInstantSnapshotGroupIdentity) String() string {
	if i.Zone != "" {
		return ZonalComputeInstantSnapshotGroupIdentityFormat.ToString(*i)
	}
	return RegionalComputeInstantSnapshotGroupIdentityFormat.ToString(*i)
}

func (i *ComputeInstantSnapshotGroupIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	if parsed, match, err := ZonalComputeInstantSnapshotGroupIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := RegionalComputeInstantSnapshotGroupIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeInstantSnapshotGroup external=%q was not known (use %s or %s)", ref, ZonalComputeInstantSnapshotGroupIdentityFormat.CanonicalForm(), RegionalComputeInstantSnapshotGroupIdentityFormat.CanonicalForm())
}

func (i *ComputeInstantSnapshotGroupIdentity) Host() string {
	return ZonalComputeInstantSnapshotGroupIdentityFormat.Host()
}

func (i *ComputeInstantSnapshotGroupIdentity) ParentString() string {
	if i.Zone != "" {
		return fmt.Sprintf("projects/%s/zones/%s", i.Project, i.Zone)
	}
	return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
}

func ParseComputeInstantSnapshotGroupExternal(external string) (*ComputeInstantSnapshotGroupIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeInstantSnapshotGroup external value")
	}
	id := &ComputeInstantSnapshotGroupIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeInstantSnapshotGroupSpec(ctx context.Context, reader client.Reader, obj *ComputeInstantSnapshotGroup) (*ComputeInstantSnapshotGroupIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID: %w", err)
	}

	location, err := refs.GetLocation(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve location: %w", err)
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project: %w", err)
	}

	identity := &ComputeInstantSnapshotGroupIdentity{
		Project:              projectID,
		InstantSnapshotGroup: resourceID,
	}

	if len(strings.Split(location, "-")) == 3 {
		identity.Zone = location
	} else {
		identity.Region = location
	}

	return identity, nil
}

func (obj *ComputeInstantSnapshotGroup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeInstantSnapshotGroupSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ComputeInstantSnapshotGroupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeInstantSnapshotGroup identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
