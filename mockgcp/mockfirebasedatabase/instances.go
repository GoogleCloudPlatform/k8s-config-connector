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

package mockfirebasedatabase

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/firebasedatabase/v1beta"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type ProjectsLocationsInstancesServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsInstancesServerServer
}

func (s *ProjectsLocationsInstancesServer) GetProjectsLocationsInstance(ctx context.Context, req *pb.GetProjectsLocationsInstanceRequest) (*pb.DatabaseInstance, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Namespace does not exist.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *ProjectsLocationsInstancesServer) CreateProjectsLocationsInstance(ctx context.Context, req *pb.CreateProjectsLocationsInstanceRequest) (*pb.DatabaseInstance, error) {
	databaseID := req.GetDatabaseId()
	if databaseID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "database_id is required")
	}

	reqName := fmt.Sprintf("%s/instances/%s", req.GetParent(), databaseID)
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetProjectsLocationsInstance()).(*pb.DatabaseInstance)
	if obj == nil {
		obj = &pb.DatabaseInstance{}
	}

	s.populateDefaultsForInstance(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ProjectsLocationsInstancesServer) DisableProjectsLocationsInstance(ctx context.Context, req *pb.DisableProjectsLocationsInstanceRequest) (*pb.DatabaseInstance, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Namespace does not exist.")
		}
		return nil, err
	}

	disabled := "DISABLED"
	obj.State = &disabled

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ProjectsLocationsInstancesServer) ReenableProjectsLocationsInstance(ctx context.Context, req *pb.ReenableProjectsLocationsInstanceRequest) (*pb.DatabaseInstance, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Namespace does not exist.")
		}
		return nil, err
	}

	active := "ACTIVE"
	obj.State = &active

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ProjectsLocationsInstancesServer) DeleteProjectsLocationsInstance(ctx context.Context, req *pb.DeleteProjectsLocationsInstanceRequest) (*pb.DatabaseInstance, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	oldObj := &pb.DatabaseInstance{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Namespace does not exist.")
		}
		return nil, err
	}

	deleted := "DELETED"
	oldObj.State = &deleted

	return oldObj, nil
}

func (s *ProjectsLocationsInstancesServer) UndeleteProjectsLocationsInstance(ctx context.Context, req *pb.UndeleteProjectsLocationsInstanceRequest) (*pb.DatabaseInstance, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Namespace does not exist.")
		}
		return nil, err
	}

	active := "ACTIVE"
	obj.State = &active

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ProjectsLocationsInstancesServer) ListProjectsLocationsInstances(ctx context.Context, req *pb.ListProjectsLocationsInstancesRequest) (*pb.ListDatabaseInstancesResponse, error) {
	parentTokens := strings.Split(req.GetParent(), "/")
	if len(parentTokens) != 4 || parentTokens[0] != "projects" || parentTokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is malformed", req.GetParent())
	}
	project, err := s.Projects.GetProjectByIDOrNumber(parentTokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/instances/", project.Number, parentTokens[3])
	var instances []*pb.DatabaseInstance
	if err := s.storage.List(ctx, (&pb.DatabaseInstance{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		if inst, ok := obj.(*pb.DatabaseInstance); ok {
			instances = append(instances, inst)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListDatabaseInstancesResponse{
		Instances: instances,
	}, nil
}

func (s *MockService) populateDefaultsForInstance(obj *pb.DatabaseInstance, name *instanceName) {
	fullName := name.String()
	obj.Name = &fullName

	if obj.Project == nil {
		projectStr := fmt.Sprintf("projects/%d", name.Project.Number)
		obj.Project = &projectStr
	}

	if obj.DatabaseUrl == nil {
		dbURL := fmt.Sprintf("https://%s.firebaseio.com", name.InstanceID)
		obj.DatabaseUrl = &dbURL
	}

	if obj.State == nil {
		active := "ACTIVE"
		obj.State = &active
	}

	if obj.Type == nil {
		dbType := "USER_DATABASE"
		obj.Type = &dbType
	}
}
