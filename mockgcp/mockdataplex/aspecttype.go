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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.AspectType

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *CatalogService) GetAspectType(ctx context.Context, req *pb.GetAspectTypeRequest) (*pb.AspectType, error) {
	name, err := s.parseAspectTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AspectType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *CatalogService) ListAspectTypes(ctx context.Context, req *pb.ListAspectTypesRequest) (*pb.ListAspectTypesResponse, error) {
	response := &pb.ListAspectTypesResponse{}

	aspectTypeKind := (&pb.AspectType{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, aspectTypeKind, storage.ListOptions{}, func(obj proto.Message) error {
		aspectType := obj.(*pb.AspectType)
		if strings.HasPrefix(aspectType.GetName(), req.Parent) {
			response.AspectTypes = append(response.AspectTypes, aspectType)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *CatalogService) CreateAspectType(ctx context.Context, req *pb.CreateAspectTypeRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/aspectTypes/%s", req.GetParent(), req.GetAspectTypeId())
	name, err := s.parseAspectTypeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetAspectType())
	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = uuid.NewString()
	if obj.Authorization == nil {
		obj.Authorization = &pb.AspectType_Authorization{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CatalogService) UpdateAspectType(ctx context.Context, req *pb.UpdateAspectTypeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAspectTypeName(req.GetAspectType().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := time.Now()

	obj := &pb.AspectType{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// If no update mask is provided, update the whole object.
		proto.Merge(obj, req.GetAspectType())
	} else {
		// Otherwise, apply fields from the update mask.
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetAspectType().GetDescription()
			case "display_name":
				obj.DisplayName = req.GetAspectType().GetDisplayName()
			case "labels":
				obj.Labels = req.GetAspectType().GetLabels()
			case "authorization":
				obj.Authorization = req.GetAspectType().GetAuthorization()
			case "metadata_template":
				obj.MetadataTemplate = req.GetAspectType().GetMetadataTemplate()
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
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *CatalogService) DeleteAspectType(ctx context.Context, req *pb.DeleteAspectTypeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAspectTypeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.AspectType{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type aspectTypeName struct {
	Project      *projects.ProjectData
	Location     string
	AspectTypeID string
}

func (n *aspectTypeName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/aspectTypes/%s", n.Project.ID, n.Location, n.AspectTypeID)
}

func (s *MockService) parseAspectTypeName(name string) (*aspectTypeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "aspectTypes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &aspectTypeName{
			Project:      project,
			Location:     tokens[3],
			AspectTypeID: tokens[5],
		}
		return n, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
