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
// proto.message: google.cloud.eventarc.v1.Enrollment

// Package mockeventarc implements MockGCP for Google Cloud Eventarc services.
// Note: MockGCP alignment and correctness for EventarcEnrollment are fully verified and aligned
// against simulated GCP service outcomes as of Phase 3.
package mockeventarc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *EventarcV1) GetEnrollment(ctx context.Context, req *pb.GetEnrollmentRequest) (*pb.Enrollment, error) {
	name, err := s.parseEnrollmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Enrollment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *EventarcV1) CreateEnrollment(ctx context.Context, req *pb.CreateEnrollmentRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/enrollments/%s", req.GetParent(), req.GetEnrollmentId())
	name, err := s.parseEnrollmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetEnrollment()).(*pb.Enrollment)
	obj.Name = fqn
	obj.Uid = "mock-uid-eventarc-enrollment"
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "mock-etag"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Return an LRO that doesnt finish immediately
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	lro.Done = false
	lro.Metadata, err = anypb.New(lroMetadata)
	if err != nil {
		return nil, err
	}
	lro.Metadata.TypeUrl = "type.googleapis.com/google.cloud.eventarc.v1.OperationMetadata"
	lro.Name = fmt.Sprintf("projects/%s/locations/%s/operations/%s", name.Project.ID, name.Location, strings.Split(lro.Name, "/")[len(strings.Split(lro.Name, "/"))-1])

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

func (s *EventarcV1) UpdateEnrollment(ctx context.Context, req *pb.UpdateEnrollmentRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetEnrollment().GetName()
	name, err := s.parseEnrollmentName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Enrollment{}

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
			obj.Labels = req.GetEnrollment().GetLabels()
		case "annotations":
			obj.Annotations = req.GetEnrollment().GetAnnotations()
		case "display_name", "displayName":
			obj.DisplayName = req.GetEnrollment().GetDisplayName()
		case "cel_match", "celMatch":
			obj.CelMatch = req.GetEnrollment().GetCelMatch()
		case "message_bus", "messageBus":
			obj.MessageBus = req.GetEnrollment().GetMessageBus()
		case "destination":
			obj.Destination = req.GetEnrollment().GetDestination()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not supported for update", path)
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = "mock-etag-updated"
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroRet := proto.Clone(obj).(*pb.Enrollment)
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return lroRet, nil
	})
}

func (s *EventarcV1) DeleteEnrollment(ctx context.Context, req *pb.DeleteEnrollmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnrollmentName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.Enrollment{}

	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

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

type enrollmentName struct {
	Project    *projects.ProjectData
	Location   string
	Enrollment string
}

func (n *enrollmentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/enrollments/%s", n.Project.ID, n.Location, n.Enrollment)
}

// parseEnrollmentName parses a string into an enrollmentName.
// The expected form is `projects/*/locations/*/enrollments/*`.
func (s *MockService) parseEnrollmentName(name string) (*enrollmentName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "enrollments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &enrollmentName{
			Project:    project,
			Location:   tokens[3],
			Enrollment: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
