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

package mockbeyondcorp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
)

type appConnectorServer struct {
	*MockService
	pb.UnimplementedAppConnectorsServiceServer
}

type appConnectorName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *appConnectorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/appConnectors/%s", n.Project.ID, n.Location, n.Name)
}

func (s *appConnectorServer) parseAppConnectorName(name string) (*appConnectorName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "appConnectors" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &appConnectorName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *appConnectorServer) GetAppConnector(ctx context.Context, req *pb.GetAppConnectorRequest) (*pb.AppConnector, error) {
	name, err := s.parseAppConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AppConnector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return proto.CloneOf(obj), nil
}

func (s *appConnectorServer) CreateAppConnector(ctx context.Context, req *pb.CreateAppConnectorRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/appConnectors/%s", req.GetParent(), req.GetAppConnectorId())
	name, err := s.parseAppConnectorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()
	obj := proto.CloneOf(req.GetAppConnector())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.AppConnector_CREATED
	obj.Uid = uuid.New().String()

	if err := s.populateDefaultsForAppConnector(name, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.AppConnectorOperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return proto.CloneOf(obj), nil
	})
}

func (s *appConnectorServer) UpdateAppConnector(ctx context.Context, req *pb.UpdateAppConnectorRequest) (*longrunning.Operation, error) {
	name, err := s.parseAppConnectorName(req.GetAppConnector().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	existing := &pb.AppConnector{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	now := time.Now()
	existing.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = []string{"display_name", "labels", "principal_info"}
	}

	for _, path := range paths {
		switch path {
		case "displayName", "display_name":
			existing.DisplayName = req.GetAppConnector().GetDisplayName()
		case "labels":
			existing.Labels = req.GetAppConnector().GetLabels()
		case "principalInfo", "principal_info":
			existing.PrincipalInfo = req.GetAppConnector().GetPrincipalInfo()
		}
	}

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	metadata := &pb.AppConnectorOperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return proto.CloneOf(existing), nil
	})
}

func (s *appConnectorServer) DeleteAppConnector(ctx context.Context, req *pb.DeleteAppConnectorRequest) (*longrunning.Operation, error) {
	name, err := s.parseAppConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.AppConnector{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	metadata := &pb.AppConnectorOperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *appConnectorServer) ListAppConnectors(ctx context.Context, req *pb.ListAppConnectorsRequest) (*pb.ListAppConnectorsResponse, error) {
	response := &pb.ListAppConnectorsResponse{}

	prefix := req.Parent + "/appConnectors/"
	findKind := (&pb.AppConnector{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.AppConnector)
		response.AppConnectors = append(response.AppConnectors, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *appConnectorServer) ReportStatus(ctx context.Context, req *pb.ReportStatusRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportStatus not implemented")
}

func (s *appConnectorServer) populateDefaultsForAppConnector(name *appConnectorName, obj *pb.AppConnector) error {
	return nil
}
