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
// proto.message: google.cloud.eventarc.v1.MessageBus

package mockeventarc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *EventarcV1) GetMessageBus(ctx context.Context, req *pb.GetMessageBusRequest) (*pb.MessageBus, error) {
	name, err := s.parseMessageBusName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MessageBus{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *EventarcV1) CreateMessageBus(ctx context.Context, req *pb.CreateMessageBusRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/messageBuses/%s", req.GetParent(), req.GetMessageBusId())
	name, err := s.parseMessageBusName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetMessageBus()).(*pb.MessageBus)
	obj.Name = fqn
	obj.Uid = "111111111111111111111"
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if obj.LoggingConfig == nil || obj.LoggingConfig.LogSeverity == pb.LoggingConfig_LOG_SEVERITY_UNSPECIFIED {
		obj.LoggingConfig = &pb.LoggingConfig{
			LogSeverity: pb.LoggingConfig_NONE,
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

func (s *EventarcV1) UpdateMessageBus(ctx context.Context, req *pb.UpdateMessageBusRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetMessageBus().GetName()
	name, err := s.parseMessageBusName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.MessageBus{}

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
			obj.Labels = req.GetMessageBus().GetLabels()
		case "annotations":
			obj.Annotations = req.GetMessageBus().GetAnnotations()
		case "display_name", "displayName":
			obj.DisplayName = req.GetMessageBus().GetDisplayName()
		case "crypto_key_name", "cryptoKeyName":
			obj.CryptoKeyName = req.GetMessageBus().GetCryptoKeyName()
		case "logging_config", "loggingConfig":
			obj.LoggingConfig = req.GetMessageBus().GetLoggingConfig()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not supported for update", path)
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroRet := proto.Clone(obj).(*pb.MessageBus)
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return lroRet, nil
	})
}

func (s *EventarcV1) DeleteMessageBus(ctx context.Context, req *pb.DeleteMessageBusRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseMessageBusName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.MessageBus{}

	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	deletedObj.Name = ""
	deletedObj.Uid = ""
	deletedObj.CreateTime = nil
	deletedObj.UpdateTime = nil
	deletedObj.Labels = nil

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return deletedObj, nil
	})
}

func (s *EventarcV1) ListMessageBuses(ctx context.Context, req *pb.ListMessageBusesRequest) (*pb.ListMessageBusesResponse, error) {
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.Parent)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	location := tokens[3]

	prefix := fmt.Sprintf("projects/%s/locations/%s/messageBuses/", project.ID, location)

	response := &pb.ListMessageBusesResponse{}
	kind := (&pb.MessageBus{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		msgBus := obj.(*pb.MessageBus)
		response.MessageBuses = append(response.MessageBuses, msgBus)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *EventarcV1) ListMessageBusEnrollments(ctx context.Context, req *pb.ListMessageBusEnrollmentsRequest) (*pb.ListMessageBusEnrollmentsResponse, error) {
	name, err := s.parseMessageBusName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListMessageBusEnrollmentsResponse{}
	enrollmentKind := (&pb.Enrollment{}).ProtoReflect().Descriptor()
	prefix := fmt.Sprintf("projects/%s/locations/%s/enrollments/", name.Project.ID, name.Location)
	if err := s.storage.List(ctx, enrollmentKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		enrollment := obj.(*pb.Enrollment)
		if enrollment.MessageBus == name.String() {
			response.Enrollments = append(response.Enrollments, enrollment.Name)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type messageBusName struct {
	Project    *projects.ProjectData
	Location   string
	MessageBus string
}

func (n *messageBusName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/messageBuses/%s", n.Project.ID, n.Location, n.MessageBus)
}

// parseMessageBusName parses a string into a messageBusName.
// The expected form is `projects/*/locations/*/messageBuses/*`.
func (s *MockService) parseMessageBusName(name string) (*messageBusName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "messageBuses" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &messageBusName{
			Project:    project,
			Location:   tokens[3],
			MessageBus: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
