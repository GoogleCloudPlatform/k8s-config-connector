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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &GKEHubMembershipBindingIdentity{}
	_ identity.Resource   = &GKEHubMembershipBinding{}
)

var GKEHubMembershipBindingIdentityFormat = gcpurls.Template[GKEHubMembershipBindingIdentity](
	"gkehub.googleapis.com",
	"projects/{project}/locations/{location}/memberships/{membership}/bindings/{membershipBinding}",
)

// GKEHubMembershipBindingIdentity is the identity of a GKEHubMembershipBinding.
// +k8s:deepcopy-gen=false
type GKEHubMembershipBindingIdentity struct {
	Project           string
	Location          string
	Membership        string
	MembershipBinding string
}

func (i *GKEHubMembershipBindingIdentity) String() string {
	return GKEHubMembershipBindingIdentityFormat.ToString(*i)
}

func (i *GKEHubMembershipBindingIdentity) Parent() string {
	return krmv1beta1.NewGKEHubMembershipIdentity(i.Project, i.Location, i.Membership).String()
}

func (i *GKEHubMembershipBindingIdentity) FromExternal(ref string) error {
	parsed, match, err := GKEHubMembershipBindingIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of GKEHubMembershipBinding external=%q was not known (use %s): %w", ref, GKEHubMembershipBindingIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of GKEHubMembershipBinding external=%q was not known (use %s)", ref, GKEHubMembershipBindingIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *GKEHubMembershipBindingIdentity) Host() string {
	return GKEHubMembershipBindingIdentityFormat.Host()
}

func getIdentityFromGKEHubMembershipBindingSpec(ctx context.Context, reader client.Reader, obj client.Object) (*GKEHubMembershipBindingIdentity, error) {
	membershipBinding := &GKEHubMembershipBinding{}
	switch t := obj.(type) {
	case *GKEHubMembershipBinding:
		membershipBinding = t
	case *unstructured.Unstructured:
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(t.Object, membershipBinding); err != nil {
			return nil, fmt.Errorf("failed to convert unstructured to GKEHubMembershipBinding: %w", err)
		}
	default:
		return nil, fmt.Errorf("expected *GKEHubMembershipBinding or *unstructured.Unstructured, got %T", obj)
	}

	membershipRef, err := krmv1beta1.ResolveGKEHubMembershipRef(ctx, reader, membershipBinding, &membershipBinding.Spec.MembershipRef)
	if err != nil {
		return nil, err
	}
	if membershipRef == nil {
		return nil, fmt.Errorf("membershipRef is required")
	}
	projectID := membershipRef.ProjectID
	location := membershipRef.Location
	membershipID := membershipRef.MembershipID

	resourceID := direct.ValueOf(membershipBinding.Spec.ResourceID)
	if resourceID == "" {
		resourceID = membershipBinding.GetName()
	}

	return &GKEHubMembershipBindingIdentity{
		Project:           projectID,
		Location:          location,
		Membership:        membershipID,
		MembershipBinding: resourceID,
	}, nil
}

func (obj *GKEHubMembershipBinding) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromGKEHubMembershipBindingSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &GKEHubMembershipBindingIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change GKEHubMembershipBinding identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
