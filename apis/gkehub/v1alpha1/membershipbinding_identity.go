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
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
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
		"projects/{projectID}/locations/{location}/memberships/{membershipID}/bindings/{membershipBindingID}",
	)
)

// GKEHubMembershipBindingIdentity defines the resource reference to GKEHubMembershipBinding, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type GKEHubMembershipBindingIdentity struct {
	ProjectID           string
	Location            string
	MembershipID        string
	MembershipBindingID string
}

func (i *GKEHubMembershipBindingIdentity) String() string {
	return membershipbindingURL.ToString(*i)
}

func (i *GKEHubMembershipBindingIdentity) ID() string {
	return i.MembershipBindingID
}

func (i *GKEHubMembershipBindingIdentity) Host() string {
	return membershipbindingURL.Host()
}

func (i *GKEHubMembershipBindingIdentity) Parent() *krmv1beta1.GKEHubMembershipIdentity {
	return krmv1beta1.NewGKEHubMembershipIdentity(i.ProjectID, i.Location, i.MembershipID)
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

func NewGKEHubMembershipBindingIdentity(project, location, membershipID, membershipBindingID string) *GKEHubMembershipBindingIdentity {
	return &GKEHubMembershipBindingIdentity{
		ProjectID:           project,
		Location:            location,
		MembershipID:        membershipID,
		MembershipBindingID: membershipBindingID,
	}
}

func (obj *GKEHubMembershipBinding) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	membershipRef, err := krmv1beta1.ResolveGKEHubMembershipRef(ctx, reader, obj, &obj.Spec.MembershipRef)
	if err != nil {
		return nil, err
	}
	if membershipRef == nil {
		return nil, fmt.Errorf("membershipRef is required")
	}
	projectID := membershipRef.ProjectID
	location := membershipRef.Location
	membershipID := membershipRef.MembershipID

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	return NewGKEHubMembershipBindingIdentity(projectID, location, membershipID, resourceID), nil
}
