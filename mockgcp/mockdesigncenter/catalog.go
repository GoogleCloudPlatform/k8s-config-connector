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

package mockdesigncenter

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *DesignCenterServer) GetCatalog(ctx context.Context, req *pb.GetCatalogRequest) (*pb.Catalog, error) {
	name, err := s.parseCatalogName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Catalog{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.populateCatalog(obj), nil
}

func (s *DesignCenterServer) CreateCatalog(ctx context.Context, req *pb.CreateCatalogRequest) (*pb.Catalog, error) {
	parent, err := s.parseSpaceName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Reject if parent space does not exist
	parentSpace := &pb.Space{}
	if err := s.storage.Get(ctx, parent.String(), parentSpace); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	if req.CatalogId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "catalogId is required")
	}

	if req.Catalog == nil {
		return nil, status.Errorf(codes.InvalidArgument, "catalog is required")
	}

	catalogID := req.CatalogId
	fqn := fmt.Sprintf("%s/catalogs/%s", parent.String(), catalogID)

	name, err := s.parseCatalogName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.Catalog).(*pb.Catalog)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.DisplayName == "" {
		// Real GCP defaults displayName to catalogId when not provided.
		obj.DisplayName = catalogID
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.populateCatalog(obj), nil
}

func (s *DesignCenterServer) UpdateCatalog(ctx context.Context, req *pb.UpdateCatalogRequest) (*pb.Catalog, error) {
	if req.Catalog == nil {
		return nil, status.Errorf(codes.InvalidArgument, "catalog is required")
	}

	name, err := s.parseCatalogName(req.Catalog.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Catalog{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Catalog)
	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(updated, req.Catalog, paths); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return s.populateCatalog(updated), nil
}

func (s *DesignCenterServer) DeleteCatalog(ctx context.Context, req *pb.DeleteCatalogRequest) (*longrunning.Operation, error) {
	name, err := s.parseCatalogName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Catalog{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Note: requestedCancellation and done are false by default.
	// TODO: Proto3 JSON marshaling omits false boolean values (requestedCancellation=false, done=false).
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		Target:                fqn,
		Verb:                  "delete",
		RequestedCancellation: false,
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *DesignCenterServer) ListCatalogs(ctx context.Context, req *pb.ListCatalogsRequest) (*pb.ListCatalogsResponse, error) {
	parent, err := s.parseSpaceName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/catalogs/"
	response := &pb.ListCatalogsResponse{}

	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	catalogKind := (&pb.Catalog{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, catalogKind, findOpts, func(item proto.Message) error {
		obj := item.(*pb.Catalog)
		response.Catalogs = append(response.Catalogs, s.populateCatalog(obj))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DesignCenterServer) populateCatalog(catalog *pb.Catalog) *pb.Catalog {
	result := proto.Clone(catalog).(*pb.Catalog)
	return result
}
