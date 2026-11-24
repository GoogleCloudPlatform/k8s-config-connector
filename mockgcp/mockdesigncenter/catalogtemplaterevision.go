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
	_ "embed"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

//go:embed template_metadata.json
var defaultTemplateMetadataJSON []byte

var defaultTemplateMetadata *pb.TFBlueprintMetadata

func init() {
	defaultTemplateMetadata = &pb.TFBlueprintMetadata{}
	if err := protojson.Unmarshal(defaultTemplateMetadataJSON, defaultTemplateMetadata); err != nil {
		panic(fmt.Sprintf("failed to unmarshal default template metadata: %v", err))
	}
}

func (s *DesignCenterServer) GetCatalogTemplateRevision(ctx context.Context, req *pb.GetCatalogTemplateRevisionRequest) (*pb.CatalogTemplateRevision, error) {
	name, err := s.parseCatalogTemplateRevisionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CatalogTemplateRevision{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.populateCatalogTemplateRevision(obj), nil
}

func (s *DesignCenterServer) CreateCatalogTemplateRevision(ctx context.Context, req *pb.CreateCatalogTemplateRevisionRequest) (*longrunning.Operation, error) {
	parent, err := s.parseCatalogTemplateName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Reject if parent catalog template does not exist
	parentCatalogTemplate := &pb.CatalogTemplate{}
	if err := s.storage.Get(ctx, parent.String(), parentCatalogTemplate); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	if req.CatalogTemplateRevisionId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "catalogTemplateRevisionId is required")
	}

	if req.CatalogTemplateRevision == nil {
		return nil, status.Errorf(codes.InvalidArgument, "catalogTemplateRevision is required")
	}

	revisionID := req.CatalogTemplateRevisionId
	fqn := fmt.Sprintf("%s/revisions/%s", parent.String(), revisionID)

	name, err := s.parseCatalogTemplateRevisionName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.CatalogTemplateRevision).(*pb.CatalogTemplateRevision)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.CatalogTemplateRevision_ACTIVE
	obj.TemplateCategory = pb.TemplateCategory_APPLICATION_TEMPLATE

	if obj.TemplateMetadata == nil && defaultTemplateMetadata != nil {
		obj.TemplateMetadata = proto.Clone(defaultTemplateMetadata).(*pb.TFBlueprintMetadata)
	}

	if obj.Uuid == "" {
		obj.Uuid = revisionID
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	populated := s.populateCatalogTemplateRevision(obj)

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.Now(),
		Target:                fqn,
		Verb:                  "create",
		RequestedCancellation: false,
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return populated, nil
	})
}

func (s *DesignCenterServer) DeleteCatalogTemplateRevision(ctx context.Context, req *pb.DeleteCatalogTemplateRevisionRequest) (*longrunning.Operation, error) {
	name, err := s.parseCatalogTemplateRevisionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.CatalogTemplateRevision{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

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

func (s *DesignCenterServer) ListCatalogTemplateRevisions(ctx context.Context, req *pb.ListCatalogTemplateRevisionsRequest) (*pb.ListCatalogTemplateRevisionsResponse, error) {
	parent, err := s.parseCatalogTemplateName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/revisions/"
	response := &pb.ListCatalogTemplateRevisionsResponse{}
	findKind := (&pb.CatalogTemplateRevision{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.CatalogTemplateRevision)
		response.CatalogTemplateRevisions = append(response.CatalogTemplateRevisions, s.populateCatalogTemplateRevision(item))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DesignCenterServer) populateCatalogTemplateRevision(revision *pb.CatalogTemplateRevision) *pb.CatalogTemplateRevision {
	result := proto.Clone(revision).(*pb.CatalogTemplateRevision)
	if result.TemplateCategory == pb.TemplateCategory_TEMPLATE_CATEGORY_UNSPECIFIED {
		result.TemplateCategory = pb.TemplateCategory_APPLICATION_TEMPLATE
	}
	if result.State == pb.CatalogTemplateRevision_STATE_UNSPECIFIED {
		result.State = pb.CatalogTemplateRevision_ACTIVE
	}
	if result.TemplateMetadata == nil && defaultTemplateMetadata != nil {
		result.TemplateMetadata = proto.Clone(defaultTemplateMetadata).(*pb.TFBlueprintMetadata)
	}
	if result.Uuid == "" {
		name, _ := s.parseCatalogTemplateRevisionName(result.Name)
		if name != nil {
			result.Uuid = name.CatalogTemplateRevision
		}
	}
	return result
}
