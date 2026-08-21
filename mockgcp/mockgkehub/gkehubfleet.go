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

package mockgkehub

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkehub/v1"
)

type GKEHubFleet struct {
	*MockService
	pb.UnimplementedGkeHubServer
}

func (s *GKEHubFleet) GetFleet(ctx context.Context, req *pb.GetFleetRequest) (*pb.Fleet, error) {
	name, err := s.parseFleetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Fleet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *GKEHubFleet) CreateFleet(ctx context.Context, req *pb.CreateFleetRequest) (*longrunning.Operation, error) {
	// GKEHub Fleet name is always the parent + "/fleets/default"
	reqName := req.Parent + "/fleets/default"
	name, err := s.parseFleetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	obj := proto.Clone(req.Resource).(*pb.Fleet)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Uid = "0123456789abcdef"
	obj.State = &pb.FleetLifecycleState{Code: "READY"}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opName := "projects/" + name.Project.ID + "/locations/" + name.Location + "/operations/{{operationID}}"
	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
		ApiVersion: "v1",
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, opName, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Fleet)
		result.CreateTime = now
		result.UpdateTime = now
		result.State = &pb.FleetLifecycleState{Code: "READY"}
		return result, nil
	})
}

func (s *GKEHubFleet) UpdateFleet(ctx context.Context, req *pb.UpdateFleetRequest) (*longrunning.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseFleetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Fleet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := timestamppb.Now()
	paths := req.GetUpdateMask().GetPaths()

	for _, path := range paths {
		switch path {
		case "displayName", "display_name":
			obj.DisplayName = req.Resource.GetDisplayName()
		case "labels":
			obj.Labels = req.Resource.GetLabels()
		case "defaultClusterConfig", "default_cluster_config":
			obj.DefaultClusterConfig = req.Resource.GetDefaultClusterConfig()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opName := "projects/" + name.Project.ID + "/locations/" + name.Location + "/operations/{{operationID}}"
	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
		ApiVersion: "v1",
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, opName, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Fleet)
		result.UpdateTime = now
		result.State = &pb.FleetLifecycleState{Code: "READY"}
		return result, nil
	})
}

func (s *GKEHubFleet) DeleteFleet(ctx context.Context, req *pb.DeleteFleetRequest) (*longrunning.Operation, error) {
	name, err := s.parseFleetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := timestamppb.Now()
	oldObj := &pb.Fleet{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	opName := "projects/" + name.Project.ID + "/locations/" + name.Location + "/operations/{{operationID}}"
	metadata := &pb.OperationMetadata{
		Target:     fqn,
		CreateTime: now,
		EndTime:    now,
		ApiVersion: "v1",
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, opName, metadata, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}
