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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.EntryType

package mockdataplex

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/dataplex/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *CatalogService) GetEntryType(ctx context.Context, req *pb.GetEntryTypeRequest) (*pb.EntryType, error) {
	name, err := s.parseEntryTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EntryType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "EntryType %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *CatalogService) ListEntryTypes(ctx context.Context, req *pb.ListEntryTypesRequest) (*pb.ListEntryTypesResponse, error) {
	name, err := s.parseEntryTypeProjectLocationName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := name.String() + "/entryTypes/"

	response := &pb.ListEntryTypesResponse{}

	entryTypeKind := (&pb.EntryType{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, entryTypeKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		entryType := obj.(*pb.EntryType)
		response.EntryTypes = append(response.EntryTypes, entryType)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *CatalogService) CreateEntryType(ctx context.Context, req *pb.CreateEntryTypeRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/entryTypes/%s", req.GetParent(), req.GetEntryTypeId())
	name, err := s.parseEntryTypeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetEntryType()).(*pb.EntryType)
	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = uuid.NewString()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CatalogService) UpdateEntryType(ctx context.Context, req *pb.UpdateEntryTypeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEntryTypeName(req.GetEntryType().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.EntryType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// If no update mask is provided, update the whole object.
		proto.Merge(obj, req.GetEntryType())
	} else {
		// Otherwise, apply fields from the update mask.
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetEntryType().GetDescription()
			case "display_name":
				obj.DisplayName = req.GetEntryType().GetDisplayName()
			case "labels":
				obj.Labels = req.GetEntryType().GetLabels()
			case "type_aliases":
				obj.TypeAliases = req.GetEntryType().GetTypeAliases()
			case "platform":
				obj.Platform = req.GetEntryType().GetPlatform()
			case "system":
				obj.System = req.GetEntryType().GetSystem()
			case "required_aspects":
				obj.RequiredAspects = req.GetEntryType().GetRequiredAspects()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = uuid.NewString()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CatalogService) DeleteEntryType(ctx context.Context, req *pb.DeleteEntryTypeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEntryTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.EntryType{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type entryTypeName struct {
	Project     *projects.ProjectData
	Location    string
	EntryTypeID string
}

func (n *entryTypeName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/entryTypes/%s", n.Project.ID, n.Location, n.EntryTypeID)
}

func (s *MockService) parseEntryTypeName(name string) (*entryTypeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entryTypes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &entryTypeName{
			Project:     project,
			Location:    tokens[3],
			EntryTypeID: tokens[5],
		}
		return n, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type entryTypeProjectLocationName struct {
	Project  *projects.ProjectData
	Location string
}

func (n *entryTypeProjectLocationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (s *MockService) parseEntryTypeProjectLocationName(name string) (*entryTypeProjectLocationName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &entryTypeProjectLocationName{
			Project:  project,
			Location: tokens[3],
		}
		return n, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
