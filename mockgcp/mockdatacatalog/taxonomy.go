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
// proto.service: google.cloud.datacatalog.v1.PolicyTagManagerServer
// proto.message: google.cloud.datacatalog.v1.Taxonomy

package mockdatacatalog

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datacatalog/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type policyTagManagerServer struct {
	*MockService
	pb.UnimplementedPolicyTagManagerServer
}

func (s *policyTagManagerServer) GetTaxonomy(ctx context.Context, req *pb.GetTaxonomyRequest) (*pb.Taxonomy, error) {
	name, err := s.parseTaxonomyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Taxonomy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Taxonomy %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *policyTagManagerServer) CreateTaxonomy(ctx context.Context, req *pb.CreateTaxonomyRequest) (*pb.Taxonomy, error) {
	// Determine the taxonomy ID from DisplayName if not explicitly set (this is a mock assumption)
	// In a real API, the ID might be generated or passed differently.
	taxonomyID := req.GetTaxonomy().GetDisplayName()
	if taxonomyID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Taxonomy display name is required to derive ID for mock")
	}
	// Replace spaces or invalid characters if needed for ID format
	taxonomyID = strings.ReplaceAll(strings.ToLower(taxonomyID), " ", "-")

	reqName := fmt.Sprintf("%s/taxonomies/%s", req.GetParent(), taxonomyID)
	name, err := s.parseTaxonomyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetTaxonomy()).(*pb.Taxonomy)
	obj.Name = fqn

	now := time.Now()
	obj.TaxonomyTimestamps = &pb.SystemTimestamps{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}

	if obj.Service == nil {
		obj.Service = &pb.Taxonomy_Service{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *policyTagManagerServer) UpdateTaxonomy(ctx context.Context, req *pb.UpdateTaxonomyRequest) (*pb.Taxonomy, error) {
	if req.GetTaxonomy() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "taxonomy is required")
	}

	name, err := s.parseTaxonomyName(req.GetTaxonomy().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Taxonomy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Taxonomy %q not found", fqn)
		}
		return nil, err
	}
	// Apply field mask updates.
	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		for _, path := range req.UpdateMask.Paths {
			switch path {
			case "description":
				obj.Description = req.Taxonomy.Description
				// Add other updatable fields here.
			}
		}
	} else {
		proto.Merge(obj, req.Taxonomy)
	}
	obj.TaxonomyTimestamps.UpdateTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *policyTagManagerServer) DeleteTaxonomy(ctx context.Context, req *pb.DeleteTaxonomyRequest) (*emptypb.Empty, error) {
	name, err := s.parseTaxonomyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Taxonomy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Taxonomy %q not found", fqn)
		}
		return nil, err
	}

	// TODO: Check for child PolicyTags before deletion? The real API might enforce this.

	return &emptypb.Empty{}, nil
}

func (s *policyTagManagerServer) ListTaxonomies(ctx context.Context, req *pb.ListTaxonomiesRequest) (*pb.ListTaxonomiesResponse, error) {
	parent, err := s.parseTaxonomyParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/taxonomies/"
	response := &pb.ListTaxonomiesResponse{}

	taxonomyKind := (&pb.Taxonomy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, taxonomyKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		taxonomy := obj.(*pb.Taxonomy)
		response.Taxonomies = append(response.Taxonomies, taxonomy)
		return nil
	}); err != nil {
		return nil, err
	}

	// TODO: Handle pagination (PageSize, PageToken)

	return response, nil
}

type taxonomyName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *taxonomyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", n.Project.ID, n.Location, n.Name)
}

func (s *MockService) parseTaxonomyName(name string) (*taxonomyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "taxonomies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &taxonomyName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}

type taxonomyParent struct {
	Project  *projects.ProjectData
	Location string
}

func (n *taxonomyParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (s *MockService) parseTaxonomyParent(name string) (*taxonomyParent, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &taxonomyParent{
			Project:  project,
			Location: tokens[3],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid parent name %q", name)
}
