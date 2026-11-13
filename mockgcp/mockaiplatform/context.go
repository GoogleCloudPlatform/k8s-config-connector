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
// proto.service: google.cloud.aiplatform.v1beta1.MetadataService
// proto.message: google.cloud.aiplatform.v1beta1.Context

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
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type ContextName struct {
	Project         *projects.ProjectData
	Location        string
	MetadataStoreId string
	ContextId       string
}

func (n *ContextName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/metadataStores/%s/contexts/%s", n.Project.Number, n.Location, n.MetadataStoreId, n.ContextId)
}

func (s *MockService) parseContextName(name string) (*ContextName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "metadataStores" && tokens[6] == "contexts" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &ContextName{
			Project:         project,
			Location:        tokens[3],
			MetadataStoreId: tokens[5],
			ContextId:       tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *metadataStoreService) GetContext(ctx context.Context, req *pb.GetContextRequest) (*pb.Context, error) {
	name, err := s.parseContextName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Context{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *metadataStoreService) CreateContext(ctx context.Context, req *pb.CreateContextRequest) (*pb.Context, error) {
	contextID := req.GetContextId()
	if contextID == "" {
		contextID = fmt.Sprintf("context-%d", time.Now().UnixNano())
	}
	reqName := req.GetParent() + "/contexts/" + contextID
	name, err := s.parseContextName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetContext())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if obj.Etag == "" {
		obj.Etag = "123456"
	}
	if obj.Metadata == nil {
		obj.Metadata = &structpb.Struct{}
	}
	if obj.SchemaTitle != "" && obj.SchemaVersion == "" {
		obj.SchemaVersion = "0.0.1"
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *metadataStoreService) UpdateContext(ctx context.Context, req *pb.UpdateContextRequest) (*pb.Context, error) {
	reqObj := req.GetContext()
	if reqObj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "context is required")
	}

	name, err := s.parseContextName(reqObj.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.Context{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply updates based on mask or field values
	if reqObj.GetDisplayName() != "" {
		obj.DisplayName = reqObj.GetDisplayName()
	}
	if reqObj.GetDescription() != "" {
		obj.Description = reqObj.GetDescription()
	}
	if reqObj.GetLabels() != nil {
		obj.Labels = reqObj.GetLabels()
	}
	if reqObj.GetMetadata() != nil {
		obj.Metadata = reqObj.GetMetadata()
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *metadataStoreService) DeleteContext(ctx context.Context, req *pb.DeleteContextRequest) (*longrunning.Operation, error) {
	name, err := s.parseContextName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Context{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fqn
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}
