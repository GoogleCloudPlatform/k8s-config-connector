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

package mockdatacatalog

import (
	"context"
	"strings"

	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datacatalog/v1"
)

func (s *DataCatalogV1) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}

func (s *DataCatalogV1) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}

func (s *DataCatalogV1) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}

func (s *DataCatalogV1) CreateTaxonomy(ctx context.Context, req *pb.CreateTaxonomyRequest) (*pb.Taxonomy, error) {
	if req.Parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}
	if req.Taxonomy == nil {
		return nil, status.Errorf(codes.InvalidArgument, "taxonomy is required")
	}

	taxonomyID := "mock-taxonomy-id"
	if req.Taxonomy.Name != "" {
		tokens := strings.Split(req.Taxonomy.Name, "/")
		taxonomyID = tokens[len(tokens)-1]
	}

	fqn := req.Parent + "/taxonomies/" + taxonomyID

	obj := proto.Clone(req.Taxonomy).(*pb.Taxonomy)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) GetTaxonomy(ctx context.Context, req *pb.GetTaxonomyRequest) (*pb.Taxonomy, error) {
	fqn := req.Name
	obj := &pb.Taxonomy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Taxonomy %s", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *DataCatalogV1) DeleteTaxonomy(ctx context.Context, req *pb.DeleteTaxonomyRequest) (*emptypb.Empty, error) {
	fqn := req.Name
	obj := &pb.Taxonomy{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Taxonomy %s", fqn)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *DataCatalogV1) UpdateTaxonomy(ctx context.Context, req *pb.UpdateTaxonomyRequest) (*pb.Taxonomy, error) {
	if req.Taxonomy == nil {
		return nil, status.Errorf(codes.InvalidArgument, "taxonomy is required")
	}
	fqn := req.Taxonomy.Name
	existing := &pb.Taxonomy{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	updated := proto.Clone(existing).(*pb.Taxonomy)
	if req.Taxonomy.DisplayName != "" {
		updated.DisplayName = req.Taxonomy.DisplayName
	}
	if req.Taxonomy.Description != "" {
		updated.Description = req.Taxonomy.Description
	}
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *DataCatalogV1) CreatePolicyTag(ctx context.Context, req *pb.CreatePolicyTagRequest) (*pb.PolicyTag, error) {
	if req.Parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}
	if req.PolicyTag == nil {
		return nil, status.Errorf(codes.InvalidArgument, "policytag is required")
	}

	policyTagID := "mock-policytag-id"
	if req.PolicyTag.Name != "" {
		tokens := strings.Split(req.PolicyTag.Name, "/")
		policyTagID = tokens[len(tokens)-1]
	}

	fqn := req.Parent + "/policyTags/" + policyTagID

	obj := proto.Clone(req.PolicyTag).(*pb.PolicyTag)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) GetPolicyTag(ctx context.Context, req *pb.GetPolicyTagRequest) (*pb.PolicyTag, error) {
	fqn := req.Name
	obj := &pb.PolicyTag{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: PolicyTag %s", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *DataCatalogV1) DeletePolicyTag(ctx context.Context, req *pb.DeletePolicyTagRequest) (*emptypb.Empty, error) {
	fqn := req.Name
	obj := &pb.PolicyTag{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: PolicyTag %s", fqn)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *DataCatalogV1) UpdatePolicyTag(ctx context.Context, req *pb.UpdatePolicyTagRequest) (*pb.PolicyTag, error) {
	if req.PolicyTag == nil {
		return nil, status.Errorf(codes.InvalidArgument, "policytag is required")
	}
	fqn := req.PolicyTag.Name
	existing := &pb.PolicyTag{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	updated := proto.Clone(existing).(*pb.PolicyTag)
	if req.PolicyTag.DisplayName != "" {
		updated.DisplayName = req.PolicyTag.DisplayName
	}
	if req.PolicyTag.Description != "" {
		updated.Description = req.PolicyTag.Description
	}
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}
