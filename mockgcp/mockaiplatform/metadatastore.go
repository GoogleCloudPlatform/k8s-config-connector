// Copyright 2025 Google LLC
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
// proto.service: google.cloud.aiplatform.v1beta1.MetadataService
// proto.message: google.cloud.aiplatform.v1beta1.MetadataStore

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type metadataStoreService struct {
	*MockService
	pb.UnimplementedMetadataServiceServer
}

func (s *metadataStoreService) GetMetadataStore(ctx context.Context, req *pb.GetMetadataStoreRequest) (*pb.MetadataStore, error) {
	name, err := s.parseMetadataStoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MetadataStore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	if obj.State == nil {
		obj.State = &pb.MetadataStore_MetadataStoreState{
			DiskUtilizationBytes: 1,
		}
	}

	return obj, nil
}

func (s *metadataStoreService) CreateMetadataStore(ctx context.Context, req *pb.CreateMetadataStoreRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/metadataStores/" + req.MetadataStoreId
	name, err := s.parseMetadataStoreName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.MetadataStore).(*pb.MetadataStore)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.CreateMetadataStoreOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := name.String()
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		// Many fields are not populated in the LRO result
		result := proto.Clone(obj).(*pb.MetadataStore)
		result.CreateTime = nil
		result.UpdateTime = nil
		return result, nil
	})
}

func (s *metadataStoreService) ListSchemas(ctx context.Context, req *pb.ListMetadataStoresRequest) (*pb.ListMetadataStoresResponse, error) {
	project, err := s.Projects.GetProjectByID(req.Parent)
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%v/", project.ID)

	var metadataStores []*pb.MetadataStore

	metadataStoreKind := (&pb.MetadataStore{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, metadataStoreKind, storage.ListOptions{}, func(obj proto.Message) error {
		metadataStore := obj.(*pb.MetadataStore)
		if strings.HasPrefix(metadataStore.GetName(), findPrefix) {
			metadataStores = append(metadataStores, metadataStore)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListMetadataStoresResponse{
		MetadataStores: metadataStores,
	}, nil
}

func (s *metadataStoreService) DeleteMetadataStore(ctx context.Context, req *pb.DeleteMetadataStoreRequest) (*longrunning.Operation, error) {
	name, err := s.parseMetadataStoreName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.MetadataStore{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s/metadataStores/%s", name.Project.Number, name.Location, name.MetadataStoreId)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type MetadataStoreName struct {
	Project         *projects.ProjectData
	Location        string
	MetadataStoreId string
}

func (n *MetadataStoreName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/metadataStores/%s", n.Project.Number, n.Location, n.MetadataStoreId)
}

// parseMetadataStoreName parses a string into a MetadataStoreName.
// The expected form of input string is projects/<projectID>/locations/<location>/metadataStores/<metadataStoreID>
func (s *MockService) parseMetadataStoreName(name string) (*MetadataStoreName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "metadataStores" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &MetadataStoreName{
			Project:         project,
			Location:        tokens[3],
			MetadataStoreId: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
