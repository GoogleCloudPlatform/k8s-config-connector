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

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type extensionRegistryService struct {
	*MockService
	pb.UnimplementedExtensionRegistryServiceServer
}

func (s *extensionRegistryService) GetExtension(ctx context.Context, req *pb.GetExtensionRequest) (*pb.Extension, error) {
	name, err := s.parseExtensionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Extension{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *extensionRegistryService) ImportExtension(ctx context.Context, req *pb.ImportExtensionRequest) (*longrunning.Operation, error) {
	var name *ExtensionName
	var err error
	reqName := req.GetExtension().GetName()
	if reqName != "" {
		name, err = s.parseExtensionName(reqName)
		if err != nil {
			return nil, err
		}
	} else {
		parent := req.GetParent()
		tokens := strings.Split(parent, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
			projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
			if err != nil {
				return nil, err
			}
			project, err := s.Projects.GetProject(projectName)
			if err != nil {
				return nil, err
			}
			name = &ExtensionName{
				Project:     project,
				Location:    tokens[3],
				ExtensionID: "1234567890123456789",
			}
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
		}
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Extension).(*pb.Extension)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Etag = computeEtag(obj)

	// Set extension operations and other output-only fields if needed
	// In the real GCP, extension operations are populated based on OpenAPI spec.
	// For simplicity, we can populate a mock extension operation so that it returns something in E2E tests.
	if len(obj.ExtensionOperations) == 0 {
		obj.ExtensionOperations = []*pb.ExtensionOperation{
			{
				OperationId: "mock-operation",
				FunctionDeclaration: &pb.FunctionDeclaration{
					Name:        "mockFunction",
					Description: "A mock function parsed from OpenAPI spec.",
				},
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.ImportExtensionOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *extensionRegistryService) UpdateExtension(ctx context.Context, req *pb.UpdateExtensionRequest) (*pb.Extension, error) {
	name, err := s.parseExtensionName(req.GetExtension().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Extension{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "display_name":
			obj.DisplayName = req.GetExtension().GetDisplayName()
		case "description":
			obj.Description = req.GetExtension().GetDescription()
		case "runtime_config":
			obj.RuntimeConfig = req.GetExtension().GetRuntimeConfig()
		case "tool_use_examples":
			obj.ToolUseExamples = req.GetExtension().GetToolUseExamples()
		case "manifest.description":
			if obj.Manifest == nil {
				obj.Manifest = &pb.ExtensionManifest{}
			}
			obj.Manifest.Description = req.GetExtension().GetManifest().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *extensionRegistryService) DeleteExtension(ctx context.Context, req *pb.DeleteExtensionRequest) (*longrunning.Operation, error) {
	name, err := s.parseExtensionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Extension{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type ExtensionName struct {
	Project     *projects.ProjectData
	Location    string
	ExtensionID string
}

func (n *ExtensionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/extensions/%s", n.Project.ID, n.Location, n.ExtensionID)
}

func (s *MockService) parseExtensionName(name string) (*ExtensionName, error) {
	name = strings.TrimPrefix(name, "//aiplatform.googleapis.com/")
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "extensions" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &ExtensionName{
			Project:     project,
			Location:    tokens[3],
			ExtensionID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
