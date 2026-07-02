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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.Identity   = &GKEHubFleetIdentity{}
	_ identity.IdentityV2 = &GKEHubFleetIdentity{}
	_ identity.Resource   = &GKEHubFleet{}

	fleetURL = gcpurls.Template[GKEHubFleetIdentity](
		"gkehub.googleapis.com",
		"projects/{projectID}/locations/{location}/fleets/{fleetID}",
	)
)

// GKEHubFleetIdentity is the identity of a GKEHubFleet.
// +k8s:deepcopy-gen=false
type GKEHubFleetIdentity struct {
	ProjectID string
	Location  string
	FleetID   string
}

func (i *GKEHubFleetIdentity) String() string {
	return fleetURL.ToString(*i)
}

func (i *GKEHubFleetIdentity) ID() string {
	return i.FleetID
}

func (i *GKEHubFleetIdentity) Host() string {
	return fleetURL.Host()
}

func (i *GKEHubFleetIdentity) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.ProjectID, i.Location)
}

func (i *GKEHubFleetIdentity) ParentString() string {
	return i.Parent()
}

func (i *GKEHubFleetIdentity) FromExternal(external string) error {
	out, match, err := fleetURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of GKEHubFleet external=%q was not known (use %s)", external, fleetURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func NewGKEHubFleetIdentity(project, location, fleetID string) *GKEHubFleetIdentity {
	return &GKEHubFleetIdentity{
		ProjectID: project,
		Location:  location,
		FleetID:   fleetID,
	}
}

func (i *GKEHubFleetIdentity) DefaultProjectState(project string) {
	if i.ProjectID == "" {
		i.ProjectID = project
	}
}

func (i *GKEHubFleetIdentity) DefaultLocationState(location string) {
	if i.Location == "" {
		i.Location = location
	}
}

func (obj *GKEHubFleet) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := project.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	location := obj.Spec.Location
	if location == "" {
		location = "global"
	}

	return NewGKEHubFleetIdentity(projectID, location, resourceID), nil
}
