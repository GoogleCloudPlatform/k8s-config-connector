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
	_ identity.Identity   = &GKEHubMembershipBindingIdentity{}
	_ identity.IdentityV2 = &GKEHubMembershipBindingIdentity{}
	_ identity.Resource   = &GKEHubMembershipBinding{}

	membershipbindingURL = gcpurls.Template[GKEHubMembershipBindingIdentity](
		"gkehub.googleapis.com",
		"projects/{projectID}/locations/{location}/memberships/{membershipID}/bindings/{membershipbindingID}",
	)
)

// GKEHubMembershipBindingIdentity defines the resource reference to GKEHubMembershipBinding, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type GKEHubMembershipBindingIdentity struct {
	ProjectID           string
	Location            string
	MembershipID        string
	MembershipbindingID string
}

func (i *GKEHubMembershipBindingIdentity) String() string {
	return membershipbindingURL.ToString(*i)
}

func (i *GKEHubMembershipBindingIdentity) ID() string {
	return i.MembershipbindingID
}

func (i *GKEHubMembershipBindingIdentity) Host() string {
	return membershipbindingURL.Host()
}

func (i *GKEHubMembershipBindingIdentity) Parent() *GKEHubMembershipIdentity {
	return NewGKEHubMembershipIdentity(i.ProjectID, i.Location, i.MembershipID)
}

func (i *GKEHubMembershipBindingIdentity) FromExternal(external string) error {
	out, match, err := membershipbindingURL.Parse(external)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of GKEHubMembershipBinding external=%q was not known (use %s)", external, membershipbindingURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func NewGKEHubMembershipBindingIdentity(project, location, membershipID, membershipbindingID string) *GKEHubMembershipBindingIdentity {
	return &GKEHubMembershipBindingIdentity{
		ProjectID:           project,
		Location:            location,
		MembershipID:        membershipID,
		MembershipbindingID: membershipbindingID,
	}
}

func (obj *GKEHubMembershipBinding) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := project.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	membershipRef, err := ResolveGKEHubMembershipRef(ctx, reader, obj, &obj.Spec.MembershipRef)
	if err != nil {
		return nil, err
	}
	membershipID := membershipRef.ID()

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	location := direct.ValueOf(obj.Spec.Location)
	if location == "" {
		location = "global"
	}

	return NewGKEHubMembershipBindingIdentity(projectID, location, membershipID, resourceID), nil
}
