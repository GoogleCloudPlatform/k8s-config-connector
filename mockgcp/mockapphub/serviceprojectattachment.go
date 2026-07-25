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
// proto.service: google.cloud.apphub.v1.AppHub
// proto.message: google.cloud.apphub.v1.ServiceProjectAttachment

package mockapphub

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

	pb "cloud.google.com/go/apphub/apiv1/apphubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *AppHubV1Service) GetServiceProjectAttachment(ctx context.Context, req *pb.GetServiceProjectAttachmentRequest) (*pb.ServiceProjectAttachment, error) {
	name, err := s.parseServiceProjectAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ServiceProjectAttachment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AppHubV1Service) CreateServiceProjectAttachment(ctx context.Context, req *pb.CreateServiceProjectAttachmentRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/serviceProjectAttachments/" + req.ServiceProjectAttachmentId
	name, err := s.parseServiceProjectAttachmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ServiceProjectAttachment).(*pb.ServiceProjectAttachment)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.Uid = "7683c772-1f55-4270-a0ea-8b49c9f42d00" // TODO: generate a unique UUID
	obj.State = pb.ServiceProjectAttachment_ACTIVE

	// Resolve the service project to its project number format as per real GCP
	serviceProjectName, err := projects.ParseProjectName(obj.ServiceProject)
	if err != nil {
		return nil, err
	}
	serviceProjectData, err := s.Projects.GetProjectByIDOrNumber(serviceProjectName.OriginalValue)
	if err != nil {
		return nil, err
	}
	obj.ServiceProject = fmt.Sprintf("projects/%d", serviceProjectData.Number)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *AppHubV1Service) DeleteServiceProjectAttachment(ctx context.Context, req *pb.DeleteServiceProjectAttachmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServiceProjectAttachmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ServiceProjectAttachment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type serviceProjectAttachmentName struct {
	Project                    *projects.ProjectData
	Location                   string
	ServiceProjectAttachmentId string
}

func (n *serviceProjectAttachmentName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/serviceProjectAttachments/%s", n.Project.ID, n.Location, n.ServiceProjectAttachmentId)
}

func (s *AppHubV1Service) parseServiceProjectAttachmentName(name string) (*serviceProjectAttachmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "serviceProjectAttachments" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &serviceProjectAttachmentName{
			Project:                    project,
			Location:                   tokens[3],
			ServiceProjectAttachmentId: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
