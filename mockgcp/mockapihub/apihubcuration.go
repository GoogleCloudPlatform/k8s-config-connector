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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type curationName struct {
	Project      *projects.ProjectData
	Location     string
	CurationName string
}

func (n *curationName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/curations/" + n.CurationName
}

func (s *MockService) parseCurationName(name string) (*curationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "curations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &curationName{
			Project:      project,
			Location:     tokens[3],
			CurationName: tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *ApiHubServer) CreateCuration(ctx context.Context, req *pb.CreateCurationRequest) (*pb.Curation, error) {
	reqName := req.Parent + "/curations/" + req.CurationId
	name, err := s.parseCurationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Curation).(*pb.Curation)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) GetCuration(ctx context.Context, req *pb.GetCurationRequest) (*pb.Curation, error) {
	name, err := s.parseCurationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Curation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) UpdateCuration(ctx context.Context, req *pb.UpdateCurationRequest) (*pb.Curation, error) {
	curationName := req.GetCuration().GetName()

	name, err := s.parseCurationName(curationName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Curation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Update display_name & description if specified in update_mask
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// By default, update display_name and description if they are set on desired
		obj.DisplayName = req.GetCuration().GetDisplayName()
		obj.Description = req.GetCuration().GetDescription()
	} else {
		for _, path := range paths {
			switch path {
			case "display_name", "displayName":
				obj.DisplayName = req.GetCuration().GetDisplayName()
			case "description":
				obj.Description = req.GetCuration().GetDescription()
			case "endpoint":
				// immutable but we can accept it if it is the same or just copy it
				obj.Endpoint = req.GetCuration().GetEndpoint()
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) DeleteCuration(ctx context.Context, req *pb.DeleteCurationRequest) (*emptypb.Empty, error) {
	name, err := s.parseCurationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Curation{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiHubServer) ListCurations(ctx context.Context, req *pb.ListCurationsRequest) (*pb.ListCurationsResponse, error) {
	name, err := s.parseCurationName(req.Parent + "/curations/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListCurationsResponse{}

	findPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location + "/curations/"

	curationKind := (&pb.Curation{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, curationKind, storage.ListOptions{}, func(obj proto.Message) error {
		curation := obj.(*pb.Curation)
		if strings.HasPrefix(curation.GetName(), findPrefix) {
			response.Curations = append(response.Curations, curation)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
