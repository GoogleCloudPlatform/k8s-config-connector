// Copyright 2022 Google LLC
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

package mockcompute

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type SubnetsV1 struct {
	*MockService
	pb.UnimplementedSubnetworksServer
}

func (s *SubnetsV1) Get(ctx context.Context, req *pb.GetSubnetworkRequest) (*pb.Subnetwork, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Subnetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SubnetsV1) Insert(ctx context.Context, req *pb.InsertSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), req.GetSubnetworkResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetSubnetworkResource()).(*pb.Subnetwork)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#subnetwork")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *SubnetsV1) Delete(ctx context.Context, req *pb.DeleteSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Subnetwork{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *SubnetsV1) SetPrivateIpGoogleAccess(ctx context.Context, req *pb.SetPrivateIpGoogleAccessSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.newSubnetName(req.GetProject(), req.GetRegion(), req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Subnetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.PrivateIpGoogleAccess = PtrTo(req.GetSubnetworksSetPrivateIpGoogleAccessRequestResource().GetPrivateIpGoogleAccess())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type subnetName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *subnetName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/subnetworks/" + n.Name
}

// parseSubnetName parses a string into a subnetName.
// The expected form is `projects/*/regions/*/subnetworks/*`.
func (s *MockService) parseSubnetName(name string) (*subnetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "subnetworks" {
		return s.newSubnetName(tokens[1], tokens[3], tokens[5])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

// newSubnetName builds a normalized subnetName from the constituent parts.
func (s *MockService) newSubnetName(project string, region string, name string) (*subnetName, error) {
	projectObj, err := s.projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}

	return &subnetName{
		Project: projectObj,
		Region:  region,
		Name:    name,
	}, nil
}
