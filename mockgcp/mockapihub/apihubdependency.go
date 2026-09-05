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

package mockapihub

import (
	"context"
	"strings"

	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ApiHubServer) GetDependency(ctx context.Context, req *pb.GetDependencyRequest) (*pb.Dependency, error) {
	name, err := s.parseDependencyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dependency{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	s.enrichDependency(ctx, obj)
	return obj, nil
}

func (s *ApiHubServer) CreateDependency(ctx context.Context, req *pb.CreateDependencyRequest) (*pb.Dependency, error) {
	reqName := req.Parent + "/dependencies/" + req.DependencyId
	name, err := s.parseDependencyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Dependency).(*pb.Dependency)
	obj.Name = fqn

	// Set default enums: discoveryMode: 1 (DISCOVERY_MODE_MANUAL), state: 2 (STATE_ACTIVE)
	// Let's set them directly. Go protobuf enums can be set with typed constants.
	obj.DiscoveryMode = pb.Dependency_DiscoveryMode(1)
	obj.State = pb.Dependency_State(2)

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	populateDependencyAttributes(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	s.enrichDependency(ctx, obj)
	return obj, nil
}

func (s *ApiHubServer) UpdateDependency(ctx context.Context, req *pb.UpdateDependencyRequest) (*pb.Dependency, error) {
	dependencyName := req.GetDependency().GetName()

	name, err := s.parseDependencyName(dependencyName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dependency{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	paths := updateMask.GetPaths()
	if len(paths) == 0 {
		// Treat as all paths if empty
		paths = []string{
			"description", "attributes",
		}
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetDependency().GetDescription()
		case "attributes":
			obj.Attributes = req.GetDependency().GetAttributes()
		}
	}

	obj.UpdateTime = timestamppb.Now()

	populateDependencyAttributes(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	s.enrichDependency(ctx, obj)
	return obj, nil
}

func (s *ApiHubServer) DeleteDependency(ctx context.Context, req *pb.DeleteDependencyRequest) (*emptypb.Empty, error) {
	name, err := s.parseDependencyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Dependency{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiHubServer) ListDependencies(ctx context.Context, req *pb.ListDependenciesRequest) (*pb.ListDependenciesResponse, error) {
	name, err := s.parseDependencyName(req.Parent + "/dependencies/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListDependenciesResponse{}

	findPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location + "/dependencies/"

	dependencyKind := (&pb.Dependency{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, dependencyKind, storage.ListOptions{}, func(obj proto.Message) error {
		dependency := obj.(*pb.Dependency)
		if strings.HasPrefix(dependency.GetName(), findPrefix) {
			s.enrichDependency(ctx, dependency)
			response.Dependencies = append(response.Dependencies, dependency)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ApiHubServer) enrichDependency(ctx context.Context, obj *pb.Dependency) {
	if obj == nil {
		return
	}
	if obj.Consumer != nil {
		extApiName := obj.Consumer.GetExternalApiResourceName()
		if extApiName != "" {
			name, err := s.parseExternalApiName(extApiName)
			if err == nil {
				extApi := &pb.ExternalApi{}
				if err := s.storage.Get(ctx, name.String(), extApi); err == nil {
					obj.Consumer.DisplayName = extApi.DisplayName
				}
			}
		}
	}
	if obj.Supplier != nil {
		extApiName := obj.Supplier.GetExternalApiResourceName()
		if extApiName != "" {
			name, err := s.parseExternalApiName(extApiName)
			if err == nil {
				extApi := &pb.ExternalApi{}
				if err := s.storage.Get(ctx, name.String(), extApi); err == nil {
					obj.Supplier.DisplayName = extApi.DisplayName
				}
			}
		}
	}
}

func populateDependencyAttributes(obj *pb.Dependency) {
	if obj == nil {
		return
	}
	for k, v := range obj.Attributes {
		if v != nil {
			v.Attribute = k
		}
	}
}
