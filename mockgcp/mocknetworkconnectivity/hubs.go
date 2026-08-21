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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
)

type hubsServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsGlobalHubsServerServer
}

func (s *hubsServer) GetProjectsLocationsGlobalHub(ctx context.Context, req *pb.GetProjectsLocationsGlobalHubRequest) (*pb.Hub, error) {
	name, err := s.parseHubName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Hub{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *hubsServer) CreateProjectsLocationsGlobalHub(ctx context.Context, req *pb.CreateProjectsLocationsGlobalHubRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/hubs/%s", req.GetParent(), req.GetHubId())
	name, err := s.parseHubName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetProjectsLocationsGlobalHub()).(*pb.Hub)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = "ACTIVE"
	obj.UniqueId = string(uuid.NewUUID())

	// Set resolved defaults
	obj.PolicyMode = "PRESET"
	obj.PresetTopology = "MESH"
	obj.RouteTables = []string{
		fmt.Sprintf("projects/%s/locations/global/hubs/%s/routeTables/default", name.Project.ID, name.HubID),
	}
	obj.ExportPsc = false

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
	prefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		lroObj := proto.Clone(obj).(*pb.Hub)
		lroObj.RouteTables = []string{
			fmt.Sprintf("projects/%d/locations/global/hubs/%s/routeTables/default", name.Project.Number, name.HubID),
		}
		return lroObj, nil
	})
}

func (s *hubsServer) DeleteProjectsLocationsGlobalHub(ctx context.Context, req *pb.DeleteProjectsLocationsGlobalHubRequest) (*longrunning.Operation, error) {
	name, err := s.parseHubName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	oldObj := &pb.Hub{}
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
	prefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *hubsServer) PatchProjectsLocationsGlobalHub(ctx context.Context, req *pb.PatchProjectsLocationsGlobalHubRequest) (*longrunning.Operation, error) {
	name, err := s.parseHubName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Hub{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	reqResource := req.GetProjectsLocationsGlobalHub()

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
	prefix := fmt.Sprintf("projects/%s/locations/global", name.Project.ID)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}
