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
	_ identity.IdentityV2 = &ComputeInstantSnapshotIdentity{}
	_ identity.Resource   = &ComputeInstantSnapshot{}
)

var ZonalComputeInstantSnapshotIdentityFormat = gcpurls.Template[ComputeInstantSnapshotIdentity](
	"compute.googleapis.com",
	"projects/{project}/zones/{zone}/instantSnapshots/{instantSnapshot}",
)
var RegionalComputeInstantSnapshotIdentityFormat = gcpurls.Template[ComputeInstantSnapshotIdentity](
	"compute.googleapis.com",
	"projects/{project}/regions/{region}/instantSnapshots/{instantSnapshot}",
)

// ComputeInstantSnapshotIdentity is the identity of a GCP ComputeInstantSnapshot resource.
// +k8s:deepcopy-gen=false
type ComputeInstantSnapshotIdentity struct {
	Project         string
	Zone            string
	Region          string
	InstantSnapshot string
}

func (i *ComputeInstantSnapshotIdentity) IsZonal() bool {
	return i.Zone != ""
}

func (i *ComputeInstantSnapshotIdentity) IsRegional() bool {
	return i.Region != ""
}

func (i *ComputeInstantSnapshotIdentity) Validate() error {
	if i.Zone == "" && i.Region == "" {
		return fmt.Errorf("ComputeInstantSnapshot identity must have either a Zone or a Region")
	}
	if i.Zone != "" && i.Region != "" {
		return fmt.Errorf("ComputeInstantSnapshot identity cannot have both a Zone and a Region")
	}
	return nil
}

func (i *ComputeInstantSnapshotIdentity) String() string {
	if i.Zone != "" {
		return ZonalComputeInstantSnapshotIdentityFormat.ToString(*i)
	}
	return RegionalComputeInstantSnapshotIdentityFormat.ToString(*i)
}

func (i *ComputeInstantSnapshotIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	if parsed, match, err := ZonalComputeInstantSnapshotIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	if parsed, match, err := RegionalComputeInstantSnapshotIdentityFormat.Parse(trimmedRef); err == nil && match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeInstantSnapshot external=%q was not known (use %s or %s)", ref, ZonalComputeInstantSnapshotIdentityFormat.CanonicalForm(), RegionalComputeInstantSnapshotIdentityFormat.CanonicalForm())
}

func (i *ComputeInstantSnapshotIdentity) Host() string {
	return ZonalComputeInstantSnapshotIdentityFormat.Host()
}

func (i *ComputeInstantSnapshotIdentity) ParentString() string {
	if i.Zone != "" {
		return fmt.Sprintf("projects/%s/zones/%s", i.Project, i.Zone)
	}
	return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
}

func ParseComputeInstantSnapshotExternal(external string) (*ComputeInstantSnapshotIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeInstantSnapshot external value")
	}
	id := &ComputeInstantSnapshotIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeInstantSnapshotSpec(ctx context.Context, reader client.Reader, obj *ComputeInstantSnapshot) (*ComputeInstantSnapshotIdentity, error) {
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

	identity := &ComputeInstantSnapshotIdentity{
		Project:         projectID,
		InstantSnapshot: resourceID,
	}

	if len(strings.Split(location, "-")) == 3 {
		identity.Zone = location
	} else {
		identity.Region = location
	}

	return identity, nil
}

func (obj *ComputeInstantSnapshot) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeInstantSnapshotSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &ComputeInstantSnapshotIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeInstantSnapshot identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
