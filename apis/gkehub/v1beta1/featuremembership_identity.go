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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClusterIdentity defines the resource reference to AlloyDBCluster, which "External" field
// holds the GCP identifier for the KRM object.
type FeatureMembershipIdentity struct {
	parent *parent.ProjectAndLocationParent
	//membershipID *MembershipIdentity
	featureID *FeatureIdentity
}

//func (i *FeatureMembershipIdentity) String() string {
//	return i.parent.String() + "/memberships/" + i.membershipID.id + "/features/" + i.featureID.id
//}

//func (i *FeatureMembershipIdentity) MembershipID() *MembershipIdentity {
//	return i.membershipID
//}

func (i *FeatureMembershipIdentity) FeatureID() *FeatureIdentity {
	return i.featureID
}

func (i *FeatureMembershipIdentity) Parent() *parent.ProjectAndLocationParent { return i.parent }

// New builds a ClusterIdentity from the Config Connector Cluster object.
func NewFeatureMembershipIdentity(ctx context.Context, reader client.Reader, obj *GKEHubFeatureMembership) (*FeatureMembershipIdentity, error) {
	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := obj.Spec.Location

	//membershipExternal, err := obj.Spec.MembershipRef.NormalizedExternal(ctx, reader, obj.Namespace)
	//if err != nil {
	//	return nil, err
	//}
	//membershipIdentity, err := ParseMembershipExternal(membershipExternal)
	//if err != nil {
	//	return nil, err
	//}

	featureExternal, err := obj.Spec.FeatureRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, err
	}
	featureIdentity, err := ParseFeatureExternal(featureExternal)
	if err != nil {
		return nil, err
	}

	// todo: support '.status.externalRef'

	return &FeatureMembershipIdentity{
		parent: &parent.ProjectAndLocationParent{
			ProjectID: projectID,
			Location:  location,
		},
		//membershipID: membershipIdentity,
		featureID: featureIdentity,
	}, nil
}
