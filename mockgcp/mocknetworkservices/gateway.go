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
// proto.service: google.cloud.networkservices.v1.NetworkServices
// proto.message: google.cloud.networkservices.v1.Gateway

package mocknetworkservices

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

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NetworkServicesServer) GetGateway(ctx context.Context, req *pb.GetGatewayRequest) (*pb.Gateway, error) {
	name, err := s.parseGatewayName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Gateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) ListGateways(ctx context.Context, req *pb.ListGatewaysRequest) (*pb.ListGatewaysResponse, error) {
	response := &pb.ListGatewaysResponse{}

	prefixName, err := s.parseGatewayName(req.Parent + "/gateways/placeholder-name")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(prefixName.String(), "placeholder-name")

	findKind := (&pb.Gateway{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		gateway := obj.(*pb.Gateway)
		response.Gateways = append(response.Gateways, gateway)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *NetworkServicesServer) CreateGateway(ctx context.Context, req *pb.CreateGatewayRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/gateways/" + req.GatewayId
	name, err := s.parseGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Gateway).(*pb.Gateway)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.SelfLink = buildSelfLink(ctx, fqn)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj).(*pb.Gateway)
		result.SelfLink = "" // Not populated here
		return result, nil
	})
}

func (s *NetworkServicesServer) UpdateGateway(ctx context.Context, req *pb.UpdateGatewayRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetGateway().GetName()

	name, err := s.parseGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Gateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Field mask is used to specify the fields to be overwritten in the
	// Gateway resource by the update.
	// The fields specified in the update_mask are relative to the resource, not
	// the full request. A field will be overwritten if it is in the mask. If the
	// user does not provide a mask then all fields will be overwritten.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// From API documentation:
		// If the update mask is not provided, all fields will be overwritten.
		// We can't know which fields are updatable, so we'll merge.
		// This is not a perfect implementation.
		// We should only copy fields that are not output-only.
		now := time.Now()
		req.Gateway.CreateTime = obj.CreateTime
		req.Gateway.UpdateTime = timestamppb.New(now)
		req.Gateway.Name = obj.Name
		obj = req.Gateway
	} else {
		// TODO: Some sort of helper for fieldmask?
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetGateway().GetLabels()
			case "description":
				obj.Description = req.GetGateway().GetDescription()
			// Add other updatable fields here
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
		obj.UpdateTime = timestamppb.New(time.Now())
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		result := proto.Clone(obj).(*pb.Gateway)
		return result, nil
	})
}

func (s *NetworkServicesServer) DeleteGateway(ctx context.Context, req *pb.DeleteGatewayRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseGatewayName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Gateway{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())

		result := &emptypb.Empty{}
		return result, nil
	})
}

type gatewayName struct {
	Project     *projects.ProjectData
	Location    string
	GatewayName string
}

func (n *gatewayName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/gateways/" + n.GatewayName
}

// parseGatewayName parses a string into an gatewayName.
// The expected form is `projects/*/locations/*/gateways/*`.
func (s *NetworkServicesServer) parseGatewayName(name string) (*gatewayName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "gateways" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &gatewayName{
			Project:     project,
			Location:    tokens[3],
			GatewayName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
