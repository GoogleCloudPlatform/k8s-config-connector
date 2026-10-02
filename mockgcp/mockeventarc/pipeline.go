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

// +tool:mockgcp-support
// proto.service: google.cloud.eventarc.v1.Eventarc
// proto.message: google.cloud.eventarc.v1.Pipeline

package mockeventarc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *EventarcV1) GetPipeline(ctx context.Context, req *pb.GetPipelineRequest) (*pb.Pipeline, error) {
	name, err := s.parsePipelineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Pipeline{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *EventarcV1) CreatePipeline(ctx context.Context, req *pb.CreatePipelineRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/pipelines/%s", req.GetParent(), req.GetPipelineId())
	name, err := s.parsePipelineName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetPipeline()).(*pb.Pipeline)
	obj.Name = fqn
	obj.Uid = "111111111111111111111"
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.RetryPolicy == nil {
		obj.RetryPolicy = &pb.Pipeline_RetryPolicy{
			MaxAttempts:   5,
			MinRetryDelay: &durationpb.Duration{Seconds: 1},
			MaxRetryDelay: &durationpb.Duration{Seconds: 60},
		}
	} else {
		if obj.RetryPolicy.MaxAttempts == 0 {
			obj.RetryPolicy.MaxAttempts = 5
		}
		if obj.RetryPolicy.MinRetryDelay == nil {
			obj.RetryPolicy.MinRetryDelay = &durationpb.Duration{Seconds: 1}
		}
		if obj.RetryPolicy.MaxRetryDelay == nil {
			obj.RetryPolicy.MaxRetryDelay = &durationpb.Duration{Seconds: 60}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

func (s *EventarcV1) UpdatePipeline(ctx context.Context, req *pb.UpdatePipelineRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetPipeline().GetName()
	name, err := s.parsePipelineName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Pipeline{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetPipeline().GetLabels()
		case "annotations":
			obj.Annotations = req.GetPipeline().GetAnnotations()
		case "display_name", "displayName":
			obj.DisplayName = req.GetPipeline().GetDisplayName()
		case "destinations":
			obj.Destinations = req.GetPipeline().GetDestinations()
		case "mediations":
			obj.Mediations = req.GetPipeline().GetMediations()
		case "crypto_key_name", "cryptoKeyName":
			obj.CryptoKeyName = req.GetPipeline().GetCryptoKeyName()
		case "logging_config", "loggingConfig":
			obj.LoggingConfig = req.GetPipeline().GetLoggingConfig()
		case "retry_policy", "retryPolicy":
			obj.RetryPolicy = req.GetPipeline().GetRetryPolicy()
		case "input_payload_format", "inputPayloadFormat":
			obj.InputPayloadFormat = req.GetPipeline().GetInputPayloadFormat()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not supported for update", path)
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroRet := proto.Clone(obj).(*pb.Pipeline)
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return lroRet, nil
	})
}

func (s *EventarcV1) DeletePipeline(ctx context.Context, req *pb.DeletePipelineRequest) (*longrunningpb.Operation, error) {
	name, err := s.parsePipelineName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.Pipeline{}

	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	deletedObj.Name = fmt.Sprintf("projects/%d/locations/%s/pipelines/%s", name.Project.Number, name.Location, name.Pipeline)
	deletedObj.Uid = ""
	deletedObj.CreateTime = nil
	deletedObj.UpdateTime = nil
	deletedObj.Labels = nil

	now := time.Now()
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return deletedObj, nil
	})
}

func (s *EventarcV1) ListPipelines(ctx context.Context, req *pb.ListPipelinesRequest) (*pb.ListPipelinesResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	location := tokens[3]

	prefix := fmt.Sprintf("projects/%s/locations/%s/pipelines/", project.ID, location)

	response := &pb.ListPipelinesResponse{}
	kind := (&pb.Pipeline{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		pipeline := obj.(*pb.Pipeline)
		response.Pipelines = append(response.Pipelines, pipeline)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type pipelineName struct {
	Project  *projects.ProjectData
	Location string
	Pipeline string
}

func (n *pipelineName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/pipelines/%s", n.Project.ID, n.Location, n.Pipeline)
}

// parsePipelineName parses a string into a pipelineName.
// The expected form is `projects/*/locations/*/pipelines/*`.
func (s *MockService) parsePipelineName(name string) (*pipelineName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "pipelines" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &pipelineName{
			Project:  project,
			Location: tokens[3],
			Pipeline: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
