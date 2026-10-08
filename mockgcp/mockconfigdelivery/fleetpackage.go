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

package mockconfigdelivery

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

	pb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type ConfigDeliveryServer struct {
	*MockService
	pb.UnimplementedConfigDeliveryServer
}

type fleetPackageName struct {
	Project      *projects.ProjectData
	Location     string
	FleetPackage string
}

func (n *fleetPackageName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/fleetPackages/" + n.FleetPackage
}

func (s *MockService) parseFleetPackageName(name string) (*fleetPackageName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "fleetPackages" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &fleetPackageName{
			Project:      project,
			Location:     tokens[3],
			FleetPackage: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ConfigDeliveryServer) GetFleetPackage(ctx context.Context, req *pb.GetFleetPackageRequest) (*pb.FleetPackage, error) {
	name, err := s.parseFleetPackageName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.FleetPackage{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ConfigDeliveryServer) CreateFleetPackage(ctx context.Context, req *pb.CreateFleetPackageRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/fleetPackages/" + req.FleetPackageId
	name, err := s.parseFleetPackageName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.Clone(req.FleetPackage).(*pb.FleetPackage)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.Target != nil && obj.Target.GetFleet() != nil {
		if obj.Target.GetFleet().Selector == nil {
			obj.Target.GetFleet().Selector = &pb.Fleet_LabelSelector{}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "create",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		lroResponse := proto.Clone(obj).(*pb.FleetPackage)
		return lroResponse, nil
	})
}

func (s *ConfigDeliveryServer) UpdateFleetPackage(ctx context.Context, req *pb.UpdateFleetPackageRequest) (*longrunning.Operation, error) {
	reqName := req.GetFleetPackage().GetName()

	name, err := s.parseFleetPackageName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := &pb.FleetPackage{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "resourceBundleSelector":
			obj.ResourceBundleSelector = req.GetFleetPackage().GetResourceBundleSelector()
		case "target":
			obj.Target = req.GetFleetPackage().GetTarget()
		case "rolloutStrategy":
			obj.RolloutStrategy = req.GetFleetPackage().GetRolloutStrategy()
		case "variantSelector":
			obj.VariantSelector = req.GetFleetPackage().GetVariantSelector()
		case "deletionPropagationPolicy":
			obj.DeletionPropagationPolicy = req.GetFleetPackage().GetDeletionPropagationPolicy()
		case "state":
			obj.State = req.GetFleetPackage().GetState()
		case "labels":
			obj.Labels = req.GetFleetPackage().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if obj.Target != nil && obj.Target.GetFleet() != nil {
		if obj.Target.GetFleet().Selector == nil {
			obj.Target.GetFleet().Selector = &pb.Fleet_LabelSelector{}
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "update",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		lroResponse := proto.Clone(obj).(*pb.FleetPackage)
		return lroResponse, nil
	})
}

func (s *ConfigDeliveryServer) DeleteFleetPackage(ctx context.Context, req *pb.DeleteFleetPackageRequest) (*longrunning.Operation, error) {
	name, err := s.parseFleetPackageName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	oldObj := &pb.FleetPackage{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound && req.AllowMissing {
			// If allowed, return a completed successful LRO representing already deleted or not found
			opMetadata := &pb.OperationMetadata{
				ApiVersion:            "v1",
				CreateTime:            timestamppb.New(now),
				EndTime:               timestamppb.New(now),
				Verb:                  "delete",
				RequestedCancellation: false,
				Target:                fqn,
			}
			opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
			return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
				return &emptypb.Empty{}, nil
			})
		}
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "delete",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
