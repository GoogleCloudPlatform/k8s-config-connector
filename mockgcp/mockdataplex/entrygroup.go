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
// proto.message: google.cloud.dataplex.v1.EntryGroup

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

// CatalogService implements the CatalogService GRPC service.
type CatalogService struct {
	*MockService
	pb.UnimplementedCatalogServiceServer
}

func (s *CatalogService) GetEntryGroup(ctx context.Context, req *pb.GetEntryGroupRequest) (*pb.EntryGroup, error) {
	name, err := s.parseEntryGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.EntryGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *CatalogService) CreateEntryGroup(ctx context.Context, req *pb.CreateEntryGroupRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/entryGroups/%s", req.GetParent(), req.GetEntryGroupId())
	name, err := s.parseEntryGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetEntryGroup()).(*pb.EntryGroup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = uuid.NewString()
	obj.Etag = fields.ComputeWeakEtag(obj)

	s.populateDefaultsForEntryGroup(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CatalogService) populateDefaultsForEntryGroup(obj *pb.EntryGroup) {
	// No specific defaults identified for EntryGroup yet.
}

func (s *CatalogService) UpdateEntryGroup(ctx context.Context, req *pb.UpdateEntryGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEntryGroupName(req.GetEntryGroup().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.EntryGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetEntryGroup().GetDescription()
		case "displayName":
			obj.DisplayName = req.GetEntryGroup().GetDisplayName()
		case "labels":
			obj.Labels = req.GetEntryGroup().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CatalogService) DeleteEntryGroup(ctx context.Context, req *pb.DeleteEntryGroupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEntryGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.EntryGroup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type entryGroupName struct {
	Project      *projects.ProjectData
	Location     string
	EntryGroupID string
}

func (n *entryGroupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/entryGroups/%s", n.Project.ID, n.Location, n.EntryGroupID)
}

// parseEntryGroupName parses a string into an entryGroupName.
// The expected form is `projects/*/locations/*/entryGroups/*`.
func (s *MockService) parseEntryGroupName(name string) (*entryGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entryGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &entryGroupName{
			Project:      project,
			Location:     tokens[3],
			EntryGroupID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
