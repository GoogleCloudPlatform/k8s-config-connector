// Copyright 2024 Google LLC
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

package gkehub

import (
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	api "google.golang.org/api/gkehub/v1beta"
)

func featureMembershipSpecKRMtoMembershipFeatureSpecAPI(r *krm.GKEHubFeatureMembershipSpec) (*api.MembershipFeatureSpec, error) {
	var acm *api.ConfigManagementMembershipSpec
	var err error
	if r.Configmanagement != nil {
		acm, err = convertKRMtoAPI_ConfigManagement(r.Configmanagement)
		if err != nil {
			return nil, err
		}
	}
	var poco *api.PolicyControllerMembershipSpec
	if r.Policycontroller != nil {
		poco = convertKRMtoAPI_Policycontroller(r.Policycontroller)
	}

	var mesh *api.ServiceMeshMembershipSpec
	if r.Mesh != nil {
		mesh = convertKRMtoAPI_ServiceMesh(r.Mesh)
	}

	return &api.MembershipFeatureSpec{
		Configmanagement: acm,
		Policycontroller: poco,
		Mesh:             mesh,
	}, nil
}

func adapterToFeatureMembershipKRM(a *gkeHubAdapter) (*krm.GKEHubFeatureMembership, error) {
	mId := a.membershipID
	if mId == "" {
		return nil, fmt.Errorf("membershipId is empty, expected not empty")
	}
	membership, err := membershipFromFullyQualifiedName(mId)
	if err != nil {
		return nil, err
	}
	fId := a.featureID
	if fId == "" {
		return nil, fmt.Errorf("featureId is empty, expected not empty")
	}
	feature, err := featureFromFullyQualifiedName(fId)
	if err != nil {
		return nil, err
	}
	if a.actual == nil {
		return nil, fmt.Errorf("actual api feature is nil, expected not nil")
	}
	featureMembership := &krm.GKEHubFeatureMembership{}
	featureMembership.SetGroupVersionKind(krm.GKEHubFeatureMembershipGVK)
	featureMembership.SetName(a.resourceName)
	featureMembership.SetNamespace(a.resourceNamespace)
	spec := featureMembership.Spec
	membershipSpec := a.actual.MembershipSpecs[a.membershipID]
	if membershipSpec.Configmanagement != nil {
		spec.Configmanagement = convertAPItoKRM_ConfigManagement(membershipSpec.Configmanagement)
	}
	if membershipSpec.Mesh != nil {
		spec.Mesh = convertAPItoKRM_ServiceMesh(membershipSpec.Mesh)
	}
	if membershipSpec.Policycontroller != nil {
		spec.Policycontroller = convertAPItoKRM_Policycontroller(membershipSpec.Policycontroller)
	}

	// set references
	spec.Location = feature.location
	if membership.location != "" {
		spec.MembershipLocation = &membership.location
	}
	spec.FeatureRef = refs.FeatureRef{
		External: fId,
	}
	spec.MembershipRef = refs.MembershipRef{
		External: mId,
	}
	spec.ProjectRef = refs.FeatureProjectRef{
		External: a.projectID,
	}
	return featureMembership, nil
}
