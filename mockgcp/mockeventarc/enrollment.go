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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type enrollmentName struct {
	Project    *projects.ProjectData
	Location   string
	Enrollment string
}

func (n *enrollmentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/enrollments/%s", n.Project.ID, n.Location, n.Enrollment)
}

func (s *EventarcV1) parseEnrollmentName(name string) (*enrollmentName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}
	parts := strings.Split(name, "/")
	if len(parts) != 6 || parts[0] != "projects" || parts[2] != "locations" || parts[4] != "enrollments" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid enrollment name %q", name)
	}

	project, err := s.Projects.GetProjectByID(parts[1])
	if err != nil {
		return nil, err
	}

	return &enrollmentName{
		Project:    project,
		Location:   parts[3],
		Enrollment: parts[5],
	}, nil
}

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
	obj.Uid = fmt.Sprintf("enrollment-%s-uid", name.Enrollment)
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "etag-v1"

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

func (s *EventarcV1) updateEnrollment(ctx context.Context, fqn string, update func(obj *pb.Enrollment)) (*pb.Enrollment, error) {
	obj := &pb.Enrollment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	update(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *EventarcV1) UpdateEnrollment(ctx context.Context, req *pb.UpdateEnrollmentRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetEnrollment().GetName()
	name, err := s.parseEnrollmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Enrollment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update mask processing
	if req.GetUpdateMask() != nil && len(req.GetUpdateMask().GetPaths()) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.GetEnrollment(), req.GetUpdateMask().GetPaths()); err != nil {
			return nil, err
		}
	} else {
		obj = req.GetEnrollment()
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
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

func (s *EventarcV1) DeleteEnrollment(ctx context.Context, req *pb.DeleteEnrollmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnrollmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Enrollment{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
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
