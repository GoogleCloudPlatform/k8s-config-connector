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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type RegionalBackendServicesV1 struct {
	*MockService
	pb.UnimplementedRegionBackendServicesServer
}

func (s *RegionalBackendServicesV1) Get(ctx context.Context, req *pb.GetRegionBackendServiceRequest) (*pb.BackendService, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/backendServices/" + req.GetBackendService()
	name, err := s.parseRegionalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackendService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
	}

	return obj, nil
}

func (s *RegionalBackendServicesV1) Insert(ctx context.Context, req *pb.InsertRegionBackendServiceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/backendServices/" + req.GetBackendServiceResource().GetName()
	name, err := s.parseRegionalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetBackendServiceResource()).(*pb.BackendService)
	obj.SelfLink = PtrTo("https://www.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#backendService")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalBackendServicesV1) Update(ctx context.Context, req *pb.UpdateRegionBackendServiceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/backendServices/" + req.GetBackendServiceResource().GetName()
	name, err := s.parseRegionalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.BackendService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Implement helper to implement the full rules here
	proto.Merge(obj, req.GetBackendServiceResource())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *RegionalBackendServicesV1) Delete(ctx context.Context, req *pb.DeleteRegionBackendServiceRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/backendServices/" + req.GetBackendService()
	name, err := s.parseRegionalBackendServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.BackendService{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalBackendServiceName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalBackendServiceName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/backendServices/" + n.Name
}

// parseBackendServiceName parses a string into a backendserviceName.
// The expected form is `projects/*/regions/*/backendservice/*`.
func (s *MockService) parseRegionalBackendServiceName(name string) (*regionalBackendServiceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "backendServices" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalBackendServiceName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
