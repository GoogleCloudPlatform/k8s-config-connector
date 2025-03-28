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
// proto.message: google.cloud.aiplatform.v1beta1.NotebookRuntime

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

func (s *notebookService) GetNotebookRuntime(ctx context.Context, req *pb.GetNotebookRuntimeRequest) (*pb.NotebookRuntime, error) {
	name, err := s.parseNotebookRuntimeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NotebookRuntime{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "notebookRuntime %q not found", fqn)
		}
		return nil, err
	}
	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))

	return obj, nil
}

func (s *notebookService) AssignNotebookRuntime(ctx context.Context, req *pb.AssignNotebookRuntimeRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/notebookRuntimes/%s", req.GetParent(), req.GetNotebookRuntimeId())

	name, err := s.parseNotebookRuntimeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.NotebookRuntime).(*pb.NotebookRuntime)

	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.RuntimeState = pb.NotebookRuntime_RUNNING
	obj.HealthState = pb.NotebookRuntime_HEALTHY
	obj.Version = "colab-on-gcp-20250226-1800-rc0-cos"
	idleTimeoutParsed, err := time.ParseDuration("10800s")
	if err != nil {
		return nil, err
	}
	obj.IdleShutdownConfig = &pb.NotebookIdleShutdownConfig{
		IdleTimeout: durationpb.New(idleTimeoutParsed),
	}
	obj.NotebookRuntimeType = pb.NotebookRuntimeType_USER_DEFINED
	obj.ProxyUri = fmt.Sprintf("test-%d-dot-notebooks.googleusercontent.com", name.Project.Number)
	obj.RuntimeUser = "${user}"
	obj.ExpirationTime = timestamppb.New(now)
	obj.Labels = map[string]string{
		"aiplatform.googleapis.com/colab_enterprise_pool":            "false",
		"aiplatform.googleapis.com/notebook_runtime_gce_instance_id": "1234567890",
	}
	obj.NotebookRuntimeTemplateRef = &pb.NotebookRuntimeTemplateRef{
		NotebookRuntimeTemplate: "projects/${projectNumber}/locations/us-central1/notebookRuntimeTemplates/test-${uniqueId}",
	}
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	op := &pb.AssignNotebookRuntimeOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := obj.Name
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		op.ProgressMessage = "NotebookRuntime is ready to use."
		// LRO response only contains name field
		result := &pb.NotebookRuntime{}
		result.Name = obj.Name
		return result, nil
	})
}

func (s *notebookService) DeleteNotebookRuntime(ctx context.Context, req *pb.DeleteNotebookRuntimeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseNotebookRuntimeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	deleted := &pb.NotebookRuntime{}
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

type notebookRuntimeName struct {
	Project           *projects.ProjectData
	Location          string
	NotebookRuntimeID string
}

func (n *notebookRuntimeName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/notebookRuntimes/" + n.NotebookRuntimeID
}

// parseNotebookRuntimeName parses a string into a notebookRuntimeName.
// The expected form is `projects/*/locations/*/notebookRuntimes/*`.
func (s *MockService) parseNotebookRuntimeName(name string) (*notebookRuntimeName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "notebookRuntimes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &notebookRuntimeName{
			Project:           project,
			Location:          tokens[3],
			NotebookRuntimeID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
