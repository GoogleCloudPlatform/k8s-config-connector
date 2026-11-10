// Copyright 2026 High Five Labs
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

// The
// OrganizationAddressGroupService's create, get, update and delete, served
// by the project service's AddressGroupServer (addressgroup.go), which
// takes any parent.

package mocknetworksecurity

import (
	"context"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/genproto/googleapis/longrunning"
)

// OrganizationAddressGroupServer serves an organization's address groups
// (organizations/{organization}/locations/{location}/addressGroups/{id}),
// stored and changed as a project's are.
type OrganizationAddressGroupServer struct {
	*MockService
	pbv1.UnimplementedOrganizationAddressGroupServiceServer

	projects *AddressGroupServer
}

func (s *OrganizationAddressGroupServer) CreateAddressGroup(ctx context.Context, req *pbv1.CreateAddressGroupRequest) (*longrunning.Operation, error) {
	return s.projects.CreateAddressGroup(ctx, req)
}

func (s *OrganizationAddressGroupServer) GetAddressGroup(ctx context.Context, req *pbv1.GetAddressGroupRequest) (*pbv1.AddressGroup, error) {
	return s.projects.GetAddressGroup(ctx, req)
}

func (s *OrganizationAddressGroupServer) UpdateAddressGroup(ctx context.Context, req *pbv1.UpdateAddressGroupRequest) (*longrunning.Operation, error) {
	return s.projects.UpdateAddressGroup(ctx, req)
}

func (s *OrganizationAddressGroupServer) DeleteAddressGroup(ctx context.Context, req *pbv1.DeleteAddressGroupRequest) (*longrunning.Operation, error) {
	return s.projects.DeleteAddressGroup(ctx, req)
}
