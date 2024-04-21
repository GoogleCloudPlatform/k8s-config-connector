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

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
)

type endpointService struct {
	*MockService
	pb.UnimplementedEndpointServiceServer
}

func (s *endpointService) GetEndpoint(ctx context.Context, req *pb.GetEndpointRequest) (*pb.Endpoint, error) {
	name, err := s.parseEndpointName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Endpoint{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *endpointService) CreateEndpoint(ctx context.Context, req *pb.CreateEndpointRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/endpoints/" + req.EndpointId
	name, err := s.parseEndpointName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Endpoint).(*pb.Endpoint)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateEndpointOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.Clone(obj).(*pb.Endpoint)
		result.CreateTime = nil
		result.UpdateTime = nil
		result.Etag = ""

		return result, nil
	})
}

func (s *endpointService) UpdateEndpoint(ctx context.Context, req *pb.UpdateEndpointRequest) (*pb.Endpoint, error) {
	name, err := s.parseEndpointName(req.GetEndpoint().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Endpoint{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// See docs for UpdateMask
	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetEndpoint().GetDisplayName()

		case "description":
			obj.Description = req.GetEndpoint().GetDescription()

		case "labels":
			obj.Labels = req.GetEndpoint().GetLabels()

		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *endpointService) DeleteEndpoint(ctx context.Context, req *pb.DeleteEndpointRequest) (*longrunning.Operation, error) {
	name, err := s.parseEndpointName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Endpoint{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type EndpointName struct {
	Project    *projects.ProjectData
	Location   string
	EndpointID string
}

func (n *EndpointName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/endpoints/%s", n.Project.Number, n.Location, n.EndpointID)
}

// parseEndpointName parses a string into a EndpointName.
// The expected form of input string is projects/<projectID>/locations/<location>/endpoints/<endpointID>
func (s *MockService) parseEndpointName(name string) (*EndpointName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "endpoints" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &EndpointName{
			Project:    project,
			Location:   tokens[3],
			EndpointID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
