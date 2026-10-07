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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type transportsServer struct {
	*MockService
	pb.UnimplementedProjectsLocationsTransportsServerServer
}

func (s *transportsServer) GetProjectsLocationsTransport(ctx context.Context, req *pb.GetProjectsLocationsTransportRequest) (*pb.Transport, error) {
	name, err := s.parseTransportName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Transport{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *transportsServer) CreateProjectsLocationsTransport(ctx context.Context, req *pb.CreateProjectsLocationsTransportRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/transports/%s", req.GetParent(), req.GetTransportId())
	name, err := s.parseTransportName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetProjectsLocationsTransport()).(*pb.Transport)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if obj.StackType == "" {
		obj.StackType = "IPV4_ONLY"
	}
	obj.State = "PENDING_KEY"
	obj.GeneratedActivationKey = "dummy-activation-key"
	obj.PeeringNetwork = fmt.Sprintf("projects/%s/global/networks/transport-%s-vpc", name.Project.ID, name.TransportID)

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

func (s *transportsServer) DeleteProjectsLocationsTransport(ctx context.Context, req *pb.DeleteProjectsLocationsTransportRequest) (*longrunning.Operation, error) {
	name, err := s.parseTransportName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	oldObj := &pb.Transport{}
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

func (s *transportsServer) PatchProjectsLocationsTransport(ctx context.Context, req *pb.PatchProjectsLocationsTransportRequest) (*longrunning.Operation, error) {
	name, err := s.parseTransportName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.Transport{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	reqResource := req.GetProjectsLocationsTransport()

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
		case "advertised_routes", "advertisedRoutes":
			obj.AdvertisedRoutes = reqResource.AdvertisedRoutes
		case "bandwidth":
			obj.Bandwidth = reqResource.Bandwidth
		case "stack_type", "stackType":
			obj.StackType = reqResource.StackType
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
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *transportsServer) ListProjectsLocationsTransports(ctx context.Context, req *pb.ListProjectsLocationsTransportsRequest) (*pb.ListTransportsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/transports/", project.ID, tokens[3])
	var transports []*pb.Transport
	if err := s.storage.List(ctx, (&pb.Transport{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		transports = append(transports, msg.(*pb.Transport))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListTransportsResponse{
		Transports: transports,
	}, nil
}
