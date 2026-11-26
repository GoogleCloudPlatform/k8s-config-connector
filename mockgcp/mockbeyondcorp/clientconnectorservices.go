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

	pb "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type clientConnectorServicesServer struct {
	*MockService
	pb.UnimplementedClientConnectorServicesServiceServer
}

func (s *clientConnectorServicesServer) GetClientConnectorService(ctx context.Context, req *pb.GetClientConnectorServiceRequest) (*pb.ClientConnectorService, error) {
	name, err := s.parseClientConnectorServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ClientConnectorService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *clientConnectorServicesServer) CreateClientConnectorService(ctx context.Context, req *pb.CreateClientConnectorServiceRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/clientConnectorServices/%s", req.GetParent(), req.GetClientConnectorServiceId())
	name, err := s.parseClientConnectorServiceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetClientConnectorService())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.ClientConnectorService_CREATING

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.ClientConnectorServiceOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		obj.State = pb.ClientConnectorService_RUNNING
		metadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *clientConnectorServicesServer) UpdateClientConnectorService(ctx context.Context, req *pb.UpdateClientConnectorServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseClientConnectorServiceName(req.GetClientConnectorService().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ClientConnectorService{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetClientConnectorService().GetDisplayName()
		case "ingress":
			obj.Ingress = req.GetClientConnectorService().GetIngress()
		case "egress":
			obj.Egress = req.GetClientConnectorService().GetEgress()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.ClientConnectorServiceOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "update",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *clientConnectorServicesServer) DeleteClientConnectorService(ctx context.Context, req *pb.DeleteClientConnectorServiceRequest) (*longrunning.Operation, error) {
	name, err := s.parseClientConnectorServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.ClientConnectorService{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.ClientConnectorServiceOperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "delete",
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *clientConnectorServicesServer) ListClientConnectorServices(ctx context.Context, req *pb.ListClientConnectorServicesRequest) (*pb.ListClientConnectorServicesResponse, error) {
	name, err := s.parseClientConnectorServiceName(req.GetParent() + "/clientConnectorServices/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListClientConnectorServicesResponse{}
	findPrefix := strings.TrimSuffix(name.String(), "dummy")
	listKind := (&pb.ClientConnectorService{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, listKind, storage.ListOptions{}, func(obj proto.Message) error {
		svc := obj.(*pb.ClientConnectorService)
		if strings.HasPrefix(svc.GetName(), findPrefix) {
			response.ClientConnectorServices = append(response.ClientConnectorServices, svc)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
