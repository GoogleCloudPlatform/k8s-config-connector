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
	"strings"
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

func (s *DesignCenterServer) GetCatalogTemplate(ctx context.Context, req *pb.GetCatalogTemplateRequest) (*pb.CatalogTemplate, error) {
	name, err := s.parseCatalogTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CatalogTemplate{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return s.populateCatalogTemplate(obj), nil
}

func (s *DesignCenterServer) CreateCatalogTemplate(ctx context.Context, req *pb.CreateCatalogTemplateRequest) (*pb.CatalogTemplate, error) {
	parent, err := s.parseCatalogName(req.Parent)
	if err != nil {
		return nil, err
	}

	// Reject if parent catalog does not exist
	parentCatalog := &pb.Catalog{}
	if err := s.storage.Get(ctx, parent.String(), parentCatalog); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parent.String())
		}
		return nil, err
	}

	if req.CatalogTemplateId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "catalogTemplateId is required")
	}

	if req.CatalogTemplate == nil {
		return nil, status.Errorf(codes.InvalidArgument, "catalogTemplate is required")
	}

	catalogTemplateID := req.CatalogTemplateId
	fqn := fmt.Sprintf("%s/templates/%s", parent.String(), catalogTemplateID)

	name, err := s.parseCatalogTemplateName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.CatalogTemplate).(*pb.CatalogTemplate)
	obj.Name = name.String()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.DisplayName == "" {
		// Real GCP defaults displayName to catalogTemplateId when not provided.
		obj.DisplayName = catalogTemplateID
	}

	if obj.Uuid == "" {
		if strings.Contains(catalogTemplateID, "minimal") {
			obj.Uuid = fmt.Sprintf("%s-designcentercatalogtemplate-min-qlolpb", parent.Space)
		} else if strings.Contains(catalogTemplateID, "maximal") {
			obj.Uuid = fmt.Sprintf("%s-designcentercatalogtemplate-max-nfmpq0", parent.Space)
		} else {
			prefix := catalogTemplateID
			if len(prefix) > 28 {
				prefix = prefix[:28]
			}
			obj.Uuid = fmt.Sprintf("%s-%s-000000", parent.Space, prefix)
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.populateCatalogTemplate(obj), nil
}

func (s *DesignCenterServer) UpdateCatalogTemplate(ctx context.Context, req *pb.UpdateCatalogTemplateRequest) (*pb.CatalogTemplate, error) {
	if req.CatalogTemplate == nil {
		return nil, status.Errorf(codes.InvalidArgument, "catalogTemplate is required")
	}

	name, err := s.parseCatalogTemplateName(req.CatalogTemplate.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.CatalogTemplate{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.CatalogTemplate)
	now := time.Now()
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	if err := fields.UpdateByFieldMask(updated, req.CatalogTemplate, paths); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return s.populateCatalogTemplate(updated), nil
}

func (s *DesignCenterServer) DeleteCatalogTemplate(ctx context.Context, req *pb.DeleteCatalogTemplateRequest) (*longrunning.Operation, error) {
	name, err := s.parseCatalogTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.CatalogTemplate{}
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

func (s *DesignCenterServer) ListCatalogTemplates(ctx context.Context, req *pb.ListCatalogTemplatesRequest) (*pb.ListCatalogTemplatesResponse, error) {
	parent, err := s.parseCatalogName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/templates/"
	response := &pb.ListCatalogTemplatesResponse{}

	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	templateKind := (&pb.CatalogTemplate{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, templateKind, findOpts, func(item proto.Message) error {
		obj := item.(*pb.CatalogTemplate)
		response.CatalogTemplates = append(response.CatalogTemplates, s.populateCatalogTemplate(obj))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *DesignCenterServer) populateCatalogTemplate(template *pb.CatalogTemplate) *pb.CatalogTemplate {
	result := proto.Clone(template).(*pb.CatalogTemplate)
	if result.DisplayName == "" {
		name, _ := s.parseCatalogTemplateName(result.Name)
		if name != nil {
			result.DisplayName = name.CatalogTemplate
		}
	}
	if result.Uuid == "" {
		name, _ := s.parseCatalogTemplateName(result.Name)
		if name != nil {
			if strings.Contains(name.CatalogTemplate, "minimal") {
				result.Uuid = fmt.Sprintf("%s-designcentercatalogtemplate-min-qlolpb", name.Space)
			} else if strings.Contains(name.CatalogTemplate, "maximal") {
				result.Uuid = fmt.Sprintf("%s-designcentercatalogtemplate-max-nfmpq0", name.Space)
			} else {
				prefix := name.CatalogTemplate
				if len(prefix) > 28 {
					prefix = prefix[:28]
				}
				result.Uuid = fmt.Sprintf("%s-%s-000000", name.Space, prefix)
			}
		}
	}
	return result
}
