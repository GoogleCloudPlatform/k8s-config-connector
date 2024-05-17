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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/gkehub/v1beta1"
	api "google.golang.org/api/gkehub/v1beta"
)

func featureMembershipSpecKRMtoMembershipFeatureSpecAPI(r *krm.GKEHubFeatureMembershipSpec) (*api.MembershipFeatureSpec, error) {
	acm, err := convertKRMtoAPI_ConfigManagement(r.Configmanagement)
	if err != nil {
		return nil, err
	}
	return &api.MembershipFeatureSpec{
		Configmanagement: acm,
		Policycontroller: convertKRMtoAPI_Policycontroller(r.Policycontroller),
		Mesh:             convertKRMtoAPI_ServiceMesh(r.Mesh),
	}, nil
}
