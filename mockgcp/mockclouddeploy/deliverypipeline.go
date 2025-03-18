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
// proto.service: google.cloud.deploy.v1.CloudDeploy
// proto.message: google.cloud.deploy.v1.DeliveryPipeline

package mockclouddeploy

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/deploy/v1"
	"github.com/google/uuid"
)

type cloudDeploy struct {
	*MockService
	pb.UnimplementedCloudDeployServer
}

func (s *cloudDeploy) GetDeliveryPipeline(ctx context.Context, req *pb.GetDeliveryPipelineRequest) (*pb.DeliveryPipeline, error) {
	name, err := s.parseDeliveryPipelineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DeliveryPipeline{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "deliveryPipeline %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *cloudDeploy) CreateDeliveryPipeline(ctx context.Context, req *pb.CreateDeliveryPipelineRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/deliveryPipelines/%s", req.Parent, req.DeliveryPipelineId)
	name, err := s.parseDeliveryPipelineName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.DeliveryPipeline).(*pb.DeliveryPipeline)
	obj.Name = fqn

	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if obj.Pipeline == nil {
		obj.Pipeline = &pb.DeliveryPipeline_SerialPipeline{
			SerialPipeline: &pb.SerialPipeline{},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *cloudDeploy) UpdateDeliveryPipeline(ctx context.Context, req *pb.UpdateDeliveryPipelineRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDeliveryPipelineName(req.DeliveryPipeline.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DeliveryPipeline{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	req.DeliveryPipeline.Uid = obj.GetUid()
	req.DeliveryPipeline.CreateTime = obj.GetCreateTime()
	req.DeliveryPipeline.UpdateTime = timestamppb.New(time.Now())

	// Apply the update mask to the object.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.DeliveryPipeline.Description
		case "annotations":
			obj.Annotations = req.DeliveryPipeline.Annotations
		case "labels":
			obj.Labels = req.DeliveryPipeline.Labels
		case "serial_pipeline", "serialPipeline":
			obj.Pipeline = &pb.DeliveryPipeline_SerialPipeline{
				SerialPipeline: req.DeliveryPipeline.GetSerialPipeline(),
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *cloudDeploy) DeleteDeliveryPipeline(ctx context.Context, req *pb.DeleteDeliveryPipelineRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDeliveryPipelineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.DeliveryPipeline{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type deliveryPipelineName struct {
	Project          *projects.ProjectData
	Location         string
	DeliveryPipeline string
}

func (n *deliveryPipelineName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/deliveryPipelines/%s", n.Project.ID, n.Location, n.DeliveryPipeline)
}

// parseDeliveryPipelineName parses a string into a deliveryPipelineName.
// The expected form is `projects/*/locations/*/deliveryPipelines/*`.
func (s *MockService) parseDeliveryPipelineName(name string) (*deliveryPipelineName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "deliveryPipelines" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &deliveryPipelineName{
			Project:          project,
			Location:         tokens[3],
			DeliveryPipeline: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
