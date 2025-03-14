// Copyright 2024 Google LLC
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
// proto.service: google.cloud.aiplatform.v1beta1.NotebookService
// proto.message: google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type notebookService struct {
	*MockService
	pb.UnimplementedNotebookServiceServer
}

func (s *notebookService) GetNotebookRuntimeTemplate(ctx context.Context, req *pb.GetNotebookRuntimeTemplateRequest) (*pb.NotebookRuntimeTemplate, error) {
	name, err := s.parseNotebookRuntimeTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NotebookRuntimeTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "notebookRuntimeTemplate %q not found", fqn)
		}
		return nil, err
	}
	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))

	return obj, nil
}

func (s *notebookService) CreateNotebookRuntimeTemplate(ctx context.Context, req *pb.CreateNotebookRuntimeTemplateRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/notebookRuntimeTemplates/%s", req.GetParent(), req.GetNotebookRuntimeTemplateId())

	name, err := s.parseNotebookRuntimeTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.NotebookRuntimeTemplate).(*pb.NotebookRuntimeTemplate)

	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = "abcdef0123A="
	idleTimeoutParsed, err := time.ParseDuration("10800s")
	if err != nil {
		return nil, err
	}
	obj.IdleShutdownConfig = &pb.NotebookIdleShutdownConfig{
		IdleTimeout: durationpb.New(idleTimeoutParsed),
	}
	obj.NotebookRuntimeType = pb.NotebookRuntimeType_USER_DEFINED
	obj.NetworkSpec.Network = fmt.Sprintf("projects/%v/global/networks/default", name.Project.Number)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// By default, immediately finish the LRO with success.
	lroPrefix := strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	lroMetadata := &pb.CreateNotebookRuntimeTemplateOperationMetadata{}
	lroMetadata.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	objInOp := &pb.NotebookRuntimeTemplate{
		Name: strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number)),
	}
	return s.operations.DoneLRO(ctx, lroPrefix, lroMetadata, objInOp)
}

func (s *notebookService) DeleteNotebookRuntimeTemplate(ctx context.Context, req *pb.DeleteNotebookRuntimeTemplateRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNotebookRuntimeTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	deleted := &pb.NotebookRuntimeTemplate{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", fmt.Sprintf("%v", name.Project.Number), name.Location)
	return s.operations.DoneLRO(ctx, lroPrefix, op, &emptypb.Empty{})
}

type notebookRuntimeTemplateName struct {
	Project                   *projects.ProjectData
	Location                  string
	NotebookRuntimeTemplateID string
}

func (n *notebookRuntimeTemplateName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/notebookRuntimeTemplates/" + n.NotebookRuntimeTemplateID
}

// parseNotebookRuntimeTemplateName parses a string into a notebookRuntimeTemplateName.
// The expected form is `projects/*/locations/*/notebookRuntimeTemplates/*`.
func (s *MockService) parseNotebookRuntimeTemplateName(name string) (*notebookRuntimeTemplateName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "notebookRuntimeTemplates" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &notebookRuntimeTemplateName{
			Project:                   project,
			Location:                  tokens[3],
			NotebookRuntimeTemplateID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
