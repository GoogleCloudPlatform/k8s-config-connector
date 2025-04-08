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

// +tool:mockgcp-support
// proto.service: google.cloud.datastream.v1.Datastream
// proto.message: google.cloud.datastream.v1.Route

package mockdatastream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datastream/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *DatastreamV1) GetRoute(ctx context.Context, req *pb.GetRouteRequest) (*pb.Route, error) {
	name, err := s.parseRouteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Route{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DatastreamV1) CreateRoute(ctx context.Context, req *pb.CreateRouteRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/routes/%s", req.GetParent(), req.GetRouteId())
	name, err := s.parseRouteName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetRoute()).(*pb.Route)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", name.Project.ID, name.Location, name.PrivateConnectionID)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DatastreamV1) DeleteRoute(ctx context.Context, req *pb.DeleteRouteRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRouteName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Route{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s/privateConnections/%s", name.Project.ID, name.Location, name.PrivateConnectionID)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type routeName struct {
	Project             *projects.ProjectData
	Location            string
	PrivateConnectionID string
	RouteID             string
}

func (n *routeName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/privateConnections/" + n.PrivateConnectionID + "/routes/" + n.RouteID
}

// parseRouteName parses a string into a routeName.
// The expected form is `projects/*/locations/*/privateConnections/*/routes/*`.
func (s *MockService) parseRouteName(name string) (*routeName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "privateConnections" && tokens[6] == "routes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &routeName{
			Project:             project,
			Location:            tokens[3],
			PrivateConnectionID: tokens[5],
			RouteID:             tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
