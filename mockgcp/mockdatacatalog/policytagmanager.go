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

package mockdatacatalog

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datacatalog/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type PolicyTagManagerV1 struct {
	*MockService
	pb.UnimplementedPolicyTagManagerServer
}

func (s *PolicyTagManagerV1) CreateTaxonomy(ctx context.Context, req *pb.CreateTaxonomyRequest) (*pb.Taxonomy, error) {
	parent, err := s.parseLocationName(req.Parent)
	if err != nil {
		return nil, err
	}

	taxonomyID := "mocktaxonomyid"
	fqn := fmt.Sprintf("%s/taxonomies/%s", parent.String(), taxonomyID)

	obj := proto.Clone(req.Taxonomy).(*pb.Taxonomy)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PolicyTagManagerV1) GetTaxonomy(ctx context.Context, req *pb.GetTaxonomyRequest) (*pb.Taxonomy, error) {
	name, err := s.parseTaxonomyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Taxonomy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PolicyTagManagerV1) UpdateTaxonomy(ctx context.Context, req *pb.UpdateTaxonomyRequest) (*pb.Taxonomy, error) {
	name, err := s.parseTaxonomyName(req.Taxonomy.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Taxonomy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Simple overwrite for now
	newObj := proto.Clone(req.Taxonomy).(*pb.Taxonomy)
	if err := s.storage.Update(ctx, fqn, newObj); err != nil {
		return nil, err
	}

	return newObj, nil
}

func (s *PolicyTagManagerV1) DeleteTaxonomy(ctx context.Context, req *pb.DeleteTaxonomyRequest) (*emptypb.Empty, error) {
	name, err := s.parseTaxonomyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.Taxonomy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PolicyTagManagerV1) ListTaxonomies(ctx context.Context, req *pb.ListTaxonomiesRequest) (*pb.ListTaxonomiesResponse, error) {
	parent, err := s.parseLocationName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/taxonomies/"
	var taxonomies []*pb.Taxonomy
	err = s.storage.List(ctx, (&pb.Taxonomy{}).ProtoReflect().Descriptor(), storage.ListOptions{}, func(obj proto.Message) error {
		taxonomy := obj.(*pb.Taxonomy)
		if strings.HasPrefix(taxonomy.Name, prefix) {
			taxonomies = append(taxonomies, taxonomy)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListTaxonomiesResponse{
		Taxonomies: taxonomies,
	}, nil
}

func (s *PolicyTagManagerV1) CreatePolicyTag(ctx context.Context, req *pb.CreatePolicyTagRequest) (*pb.PolicyTag, error) {
	parent, err := s.parseTaxonomyName(req.Parent)
	if err != nil {
		return nil, err
	}

	policyTagID := "mockpolicytagid"
	fqn := fmt.Sprintf("%s/policyTags/%s", parent.String(), policyTagID)

	obj := proto.Clone(req.PolicyTag).(*pb.PolicyTag)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PolicyTagManagerV1) GetPolicyTag(ctx context.Context, req *pb.GetPolicyTagRequest) (*pb.PolicyTag, error) {
	name, err := s.parsePolicyTagName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.PolicyTag{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PolicyTagManagerV1) UpdatePolicyTag(ctx context.Context, req *pb.UpdatePolicyTagRequest) (*pb.PolicyTag, error) {
	name, err := s.parsePolicyTagName(req.PolicyTag.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.PolicyTag{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Simple overwrite for now
	newObj := proto.Clone(req.PolicyTag).(*pb.PolicyTag)
	if err := s.storage.Update(ctx, fqn, newObj); err != nil {
		return nil, err
	}

	return newObj, nil
}

func (s *PolicyTagManagerV1) DeletePolicyTag(ctx context.Context, req *pb.DeletePolicyTagRequest) (*emptypb.Empty, error) {
	name, err := s.parsePolicyTagName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &pb.PolicyTag{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PolicyTagManagerV1) ListPolicyTags(ctx context.Context, req *pb.ListPolicyTagsRequest) (*pb.ListPolicyTagsResponse, error) {
	parent, err := s.parseTaxonomyName(req.Parent)
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/policyTags/"
	var policyTags []*pb.PolicyTag
	err = s.storage.List(ctx, (&pb.PolicyTag{}).ProtoReflect().Descriptor(), storage.ListOptions{}, func(obj proto.Message) error {
		pt := obj.(*pb.PolicyTag)
		if strings.HasPrefix(pt.Name, prefix) {
			policyTags = append(policyTags, pt)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListPolicyTagsResponse{
		PolicyTags: policyTags,
	}, nil
}

type locationName struct {
	Project  *projects.ProjectData
	Location string
}

func (n *locationName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

func (s *MockService) parseLocationName(name string) (*locationName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	return &locationName{
		Project:  project,
		Location: tokens[3],
	}, nil
}

type taxonomyName struct {
	Project  *projects.ProjectData
	Location string
	Taxonomy string
}

func (n *taxonomyName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s", n.Project.ID, n.Location, n.Taxonomy)
}

func (s *MockService) parseTaxonomyName(name string) (*taxonomyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "taxonomies" {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	return &taxonomyName{
		Project:  project,
		Location: tokens[3],
		Taxonomy: tokens[5],
	}, nil
}

type policyTagName struct {
	Project   *projects.ProjectData
	Location  string
	Taxonomy  string
	PolicyTag string
}

func (n *policyTagName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/taxonomies/%s/policyTags/%s", n.Project.ID, n.Location, n.Taxonomy, n.PolicyTag)
}

func (s *MockService) parsePolicyTagName(name string) (*policyTagName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 8 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "taxonomies" || tokens[6] != "policyTags" {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	return &policyTagName{
		Project:   project,
		Location:  tokens[3],
		Taxonomy:  tokens[5],
		PolicyTag: tokens[7],
	}, nil
}

type PolicyTagManagerSerializationV1 struct {
	*MockService
	pb.UnimplementedPolicyTagManagerSerializationServer
}
