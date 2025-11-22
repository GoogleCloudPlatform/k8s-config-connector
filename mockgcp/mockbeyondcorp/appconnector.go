// Copyright 2025 Google LLC
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
// proto.service: google.cloud.beyondcorp.appconnectors.v1.AppConnectorsService
// proto.message: google.cloud.beyondcorp.appconnectors.v1.AppConnector

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

	pb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

// GetAppConnector gets details of a single AppConnector.
func (s *AppConnectorsService) GetAppConnector(ctx context.Context, req *pb.GetAppConnectorRequest) (*pb.AppConnector, error) {
	name, err := s.parseAppConnectorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AppConnector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "AppConnector %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

// ListAppConnectors lists AppConnectors in a given project and location.
func (s *AppConnectorsService) ListAppConnectors(ctx context.Context, req *pb.ListAppConnectorsRequest) (*pb.ListAppConnectorsResponse, error) {
	dummyName, err := s.parseAppConnectorName(req.GetParent() + "/appConnectors/dummy")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(dummyName.String(), "dummy")

	response := &pb.ListAppConnectorsResponse{}
	s.storage.List(ctx, (&pb.AppConnector{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		response.AppConnectors = append(response.AppConnectors, obj.(*pb.AppConnector))
		return nil
	})

	return response, nil
}

// CreateAppConnector creates a new AppConnector.
func (s *AppConnectorsService) CreateAppConnector(ctx context.Context, req *pb.CreateAppConnectorRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/appConnectors/" + req.GetAppConnectorId()
	name, err := s.parseAppConnectorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetAppConnector()).(*pb.AppConnector)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = uuid.New().String()
	obj.State = pb.AppConnector_CREATING

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.AppConnectorOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
	}

	return s.operations.StartLRO(ctx, req.GetParent(), lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		obj.State = pb.AppConnector_CREATED
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

// UpdateAppConnector updates an existing AppConnector.
func (s *AppConnectorsService) UpdateAppConnector(ctx context.Context, req *pb.UpdateAppConnectorRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAppConnectorName(req.GetAppConnector().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.AppConnector{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update fields based on update_mask
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "labels":
			obj.Labels = req.GetAppConnector().GetLabels()
		case "display_name", "displayName":
			obj.DisplayName = req.GetAppConnector().GetDisplayName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.AppConnectorOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
	}

	return s.operations.StartLRO(ctx, name.Parent(), lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		return obj, nil
	})
}

// DeleteAppConnector deletes an AppConnector.
func (s *AppConnectorsService) DeleteAppConnector(ctx context.Context, req *pb.DeleteAppConnectorRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAppConnectorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.AppConnector{}
	if err := s.storage.Get(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Deleting a non-existent resource is a no-op LRO for some services.
			// Let's assume it completes successfully.
			return s.operations.DoneLRO(ctx, name.Parent(), &pb.AppConnectorOperationMetadata{}, &emptypb.Empty{})
		}
		return nil, err
	}

	now := time.Now()
	oldObj.State = pb.AppConnector_DELETING
	if err := s.storage.Update(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.AppConnectorOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
	}

	return s.operations.StartLRO(ctx, name.Parent(), lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		if err := s.storage.Delete(ctx, fqn, &pb.AppConnector{}); err != nil {
			return nil, err
		}
		return &emptypb.Empty{}, nil
	})
}

// appConnectorName defines the structure of a parsed AppConnector name.
type appConnectorName struct {
	Project        *projects.ProjectData
	Location       string
	AppConnectorID string
}

func (n *appConnectorName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/appConnectors/%s", n.Project.ID, n.Location, n.AppConnectorID)
}

func (n *appConnectorName) Parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

// parseAppConnectorName parses a string into an appConnectorName.
// The expected form is `projects/{project}/locations/{location}/appConnectors/{app_connector}`.
func (s *MockService) parseAppConnectorName(name string) (*appConnectorName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "appConnectors" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &appConnectorName{
			Project:        project,
			Location:       tokens[3],
			AppConnectorID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid appConnector name", name)
}
