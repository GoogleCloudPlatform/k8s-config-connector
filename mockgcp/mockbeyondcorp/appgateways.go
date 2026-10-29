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

// +tool:mockgcp-support
// proto.service: google.cloud.beyondcorp.appgateways.v1.AppGatewaysService
// proto.message: google.cloud.beyondcorp.appgateways.v1.AppGateway

package mockbeyondcorp

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

	pb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type AppGatewaysV1 struct {
	*MockService
	pb.UnimplementedAppGatewaysServiceServer
}

type appGatewayName struct {
	Project      *projects.ProjectData
	Location     string
	AppGatewayID string
}

func (n *appGatewayName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/appGateways/%s", n.Project.ID, n.Location, n.AppGatewayID)
}

func (s *AppGatewaysV1) parseAppGatewayName(name string) (*appGatewayName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "appGateways" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &appGatewayName{
			Project:      project,
			Location:     tokens[3],
			AppGatewayID: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *AppGatewaysV1) GetAppGateway(ctx context.Context, req *pb.GetAppGatewayRequest) (*pb.AppGateway, error) {
	name, err := s.parseAppGatewayName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AppGateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AppGatewaysV1) CreateAppGateway(ctx context.Context, req *pb.CreateAppGatewayRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/appGateways/" + req.GetAppGatewayId()
	name, err := s.parseAppGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()
	obj := proto.Clone(req.GetAppGateway()).(*pb.AppGateway)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = "00000000-0000-0000-0000-000000000001"
	obj.State = pb.AppGateway_CREATING
	obj.Uri = fmt.Sprintf("//beyondcorp.googleapis.com/%s", fqn)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.AppGatewayOperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		obj.State = pb.AppGateway_CREATED
		obj.UpdateTime = timestamppb.Now()
		lroMetadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *AppGatewaysV1) DeleteAppGateway(ctx context.Context, req *pb.DeleteAppGatewayRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAppGatewayName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AppGateway{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.AppGatewayOperationMetadata{
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *AppGatewaysV1) ListAppGateways(ctx context.Context, req *pb.ListAppGatewaysRequest) (*pb.ListAppGatewaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAppGateways not implemented")
}
