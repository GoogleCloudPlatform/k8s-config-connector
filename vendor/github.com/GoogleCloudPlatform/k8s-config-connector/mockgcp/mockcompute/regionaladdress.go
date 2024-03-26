// Copyright 2023 Google LLC
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

type RegionalAddressesV1 struct {
	*MockService
	pb.UnimplementedAddressesServer
}

func (s *RegionalAddressesV1) Get(ctx context.Context, req *pb.GetAddressRequest) (*pb.Address, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + req.GetAddress()
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Address{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RegionalAddressesV1) Insert(ctx context.Context, req *pb.InsertAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + req.GetAddressResource().GetName()
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetAddressResource()).(*pb.Address)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#address")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating address: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalAddressesV1) Delete(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + req.GetAddress()
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Address{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalAddressesV1) SetLabels(ctx context.Context, req *pb.SetLabelsAddressRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/addresses/" + req.GetResource()
	name, err := s.parseRegionalAddressName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Address{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetRegionSetLabelsRequestResource().GetLabels()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalAddressName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalAddressName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/networks/" + n.Name
}

// parseAddressName parses a string into a addressName.
// The expected form is `projects/*/regions/*/address/*`.
func (s *MockService) parseRegionalAddressName(name string) (*regionalAddressName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "addresses" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalAddressName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
