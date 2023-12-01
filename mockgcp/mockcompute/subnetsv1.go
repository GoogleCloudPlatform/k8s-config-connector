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
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type SubnetsV1 struct {
	*MockService
	pb.UnimplementedSubnetworksServer
}

func (s *SubnetsV1) Get(ctx context.Context, req *pb.GetSubnetworkRequest) (*pb.Subnetwork, error) {
	name, err := s.parseSubnetName("projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/subnetworks/" + req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Subnetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "subnet %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading subnet: %v", err)
		}
	}

	return obj, nil
}

func (s *SubnetsV1) Insert(ctx context.Context, req *pb.InsertSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.parseSubnetName("projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/subnetworks/" + req.GetSubnetworkResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetSubnetworkResource()).(*pb.Subnetwork)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#subsubnet")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating subnet: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *SubnetsV1) Delete(ctx context.Context, req *pb.DeleteSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.parseSubnetName("projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/subnetworks/" + req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Subnetwork{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "subnet %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting subnet: %v", err)
		}
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *SubnetsV1) SetPrivateIpGoogleAccess(ctx context.Context, req *pb.SetPrivateIpGoogleAccessSubnetworkRequest) (*pb.Operation, error) {
	name, err := s.parseSubnetName("projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/subnetworks/" + req.GetSubnetwork())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Subnetwork{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "subnet %q not found", fqn)
		}
		return nil, status.Errorf(codes.Internal, "error reading subnet: %v", err)
	}

	obj.PrivateIpGoogleAccess = PtrTo(req.GetSubnetworksSetPrivateIpGoogleAccessRequestResource().GetPrivateIpGoogleAccess())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating subnet: %v", err)
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
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &subnetName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
