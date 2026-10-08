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

package mocknetworkconnectivity

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
	"k8s.io/apimachinery/pkg/util/uuid"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
)

type multicloudDataTransferConfigsServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsMulticloudDataTransferConfigsServerServer
}

type multicloudDataTransferConfigName struct {
	Project                        *projects.ProjectData
	Location                       string
	MulticloudDataTransferConfigID string
}

func (n *multicloudDataTransferConfigName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/multicloudDataTransferConfigs/%s", n.Project.ID, n.Location, n.MulticloudDataTransferConfigID)
}

func (s *MockService) parseMulticloudDataTransferConfigName(name string) (*multicloudDataTransferConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "multicloudDataTransferConfigs" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &multicloudDataTransferConfigName{
			Project:                        project,
			Location:                       tokens[3],
			MulticloudDataTransferConfigID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *multicloudDataTransferConfigsServer) GetProjectsLocationsMulticloudDataTransferConfig(ctx context.Context, req *pb.GetProjectsLocationsMulticloudDataTransferConfigRequest) (*pb.MulticloudDataTransferConfig, error) {
	name, err := s.parseMulticloudDataTransferConfigName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MulticloudDataTransferConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *multicloudDataTransferConfigsServer) CreateProjectsLocationsMulticloudDataTransferConfig(ctx context.Context, req *pb.CreateProjectsLocationsMulticloudDataTransferConfigRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/multicloudDataTransferConfigs/%s", req.GetParent(), req.GetMulticloudDataTransferConfigId())
	name, err := s.parseMulticloudDataTransferConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetProjectsLocationsMulticloudDataTransferConfig()).(*pb.MulticloudDataTransferConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = string(uuid.NewUUID())
	obj.Etag = computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (s *multicloudDataTransferConfigsServer) DeleteProjectsLocationsMulticloudDataTransferConfig(ctx context.Context, req *pb.DeleteProjectsLocationsMulticloudDataTransferConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseMulticloudDataTransferConfigName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	oldObj := &pb.MulticloudDataTransferConfig{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *multicloudDataTransferConfigsServer) PatchProjectsLocationsMulticloudDataTransferConfig(ctx context.Context, req *pb.PatchProjectsLocationsMulticloudDataTransferConfigRequest) (*longrunning.Operation, error) {
	name, err := s.parseMulticloudDataTransferConfigName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.MulticloudDataTransferConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	reqResource := req.GetProjectsLocationsMulticloudDataTransferConfig()

	// Apply field mask updates
	paths := strings.Split(req.GetUpdateMask(), ",")
	for _, path := range paths {
		if path == "" {
			continue
		}
		switch path {
		case "description":
			obj.Description = reqResource.Description
		case "labels":
			obj.Labels = reqResource.Labels
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}
