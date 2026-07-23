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
// proto.service: google.cloud.discoveryengine.v1.IdentityMappingStoreService
// proto.message: google.cloud.discoveryengine.v1.IdentityMappingStore

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *identityMappingStoreService) CreateIdentityMappingStore(ctx context.Context, req *pb.CreateIdentityMappingStoreRequest) (*pb.IdentityMappingStore, error) {
	reqName := fmt.Sprintf("%s/identityMappingStores/%s", req.GetParent(), req.GetIdentityMappingStoreId())
	name, err := s.parseIdentityMappingStoreName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetIdentityMappingStore()).(*pb.IdentityMappingStore)
	obj.Name = fqn

	// Set cmek config if kms key is configured
	if obj.GetKmsKeyName() != "" {
		obj.CmekConfig = &pb.CmekConfig{
			KmsKey: obj.GetKmsKeyName(),
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *identityMappingStoreService) GetIdentityMappingStore(ctx context.Context, req *pb.GetIdentityMappingStoreRequest) (*pb.IdentityMappingStore, error) {
	name, err := s.parseIdentityMappingStoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.IdentityMappingStore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "IdentityMappingStore %v not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *identityMappingStoreService) DeleteIdentityMappingStore(ctx context.Context, req *pb.DeleteIdentityMappingStoreRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseIdentityMappingStoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.IdentityMappingStore{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, prefix, nil, nil)
}

type identityMappingStoreName struct {
	Project              *projects.ProjectData
	Location             string
	IdentityMappingStore string
}

func (n *identityMappingStoreName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/identityMappingStores/%s", n.Project.Number, n.Location, n.IdentityMappingStore)
}

func (s *MockService) parseIdentityMappingStoreName(name string) (*identityMappingStoreName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "identityMappingStores" {

		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &identityMappingStoreName{
			Project:              project,
			Location:             tokens[3],
			IdentityMappingStore: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
