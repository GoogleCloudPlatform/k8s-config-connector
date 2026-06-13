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
	_ identity.IdentityV2 = &ComputeDiskIdentity{}
	_ identity.Resource   = &ComputeDisk{}
)

var (
	ComputeDiskZonalIdentityFormat    = gcpurls.Template[ComputeDiskIdentity]("compute.googleapis.com", "projects/{project}/zones/{location}/disks/{disk}")
	ComputeDiskRegionalIdentityFormat = gcpurls.Template[ComputeDiskIdentity]("compute.googleapis.com", "projects/{project}/regions/{location}/disks/{disk}")
)

// ComputeDiskIdentity is the identity of a GCP ComputeDisk resource.
// +k8s:deepcopy-gen=false
type ComputeDiskIdentity struct {
	Project  string
	Location string
	Disk     string
}

func (i *ComputeDiskIdentity) IsRegional() bool {
	return isLocationRegional(i.Location)
}

func (i *ComputeDiskIdentity) String() string {
	if i.IsRegional() {
		return ComputeDiskRegionalIdentityFormat.ToString(*i)
	}
	return ComputeDiskZonalIdentityFormat.ToString(*i)
}

func (i *ComputeDiskIdentity) FromExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	parsed, match, err := ComputeDiskZonalIdentityFormat.Parse(trimmedRef)
	if err == nil && match {
		*i = *parsed
		return nil
	}

	parsed, match, err = ComputeDiskRegionalIdentityFormat.Parse(trimmedRef)
	if err == nil && match {
		*i = *parsed
		return nil
	}

	return fmt.Errorf("format of ComputeDisk external=%q was not known (expected zonal %q or regional %q)", ref, ComputeDiskZonalIdentityFormat.CanonicalForm(), ComputeDiskRegionalIdentityFormat.CanonicalForm())
}

func (i *ComputeDiskIdentity) Host() string {
	return ComputeDiskZonalIdentityFormat.Host()
}

func (i *ComputeDiskIdentity) ParentString() string {
	if i.IsRegional() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Location)
	}
	return fmt.Sprintf("projects/%s/zones/%s", i.Project, i.Location)
}

func ParseComputeDiskExternal(external string) (*ComputeDiskIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeDisk external value")
	}
	id := &ComputeDiskIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func isLocationRegional(location string) bool {
	parts := strings.Split(location, "-")
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		if len(lastPart) == 1 {
			return false
		}
	}
	return true
}

func getIdentityFromComputeDiskSpec(ctx context.Context, reader client.Reader, obj *ComputeDisk) (*ComputeDiskIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location: spec.location is empty")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &ComputeDiskIdentity{
		Project:  projectID,
		Location: location,
		Disk:     resourceID,
	}
	return identity, nil
}

func (obj *ComputeDisk) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeDiskSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeDiskIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeDisk identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
