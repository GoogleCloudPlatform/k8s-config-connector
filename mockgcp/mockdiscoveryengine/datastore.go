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
// proto.service: google.cloud.discoveryengine.v1.DataStoreService
// proto.message: google.cloud.discoveryengine.v1.DataStore

package mockdiscoveryengine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *dataStoreService) CreateDataStore(ctx context.Context, req *pb.CreateDataStoreRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/dataStores/%s", req.GetParent(), req.GetDataStoreId())
	name, err := s.parseDataStoreName(reqName)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	fqn := name.String()
	obj := proto.Clone(req.GetDataStore()).(*pb.DataStore)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.DefaultSchemaId = "default_schema"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s", name.Project.Number, name.Location, name.Collection)
	// Returns with no createTime
	lroRet := proto.Clone(obj).(*pb.DataStore)
	lroRet.CreateTime = nil
	// output-only
	lroRet.LanguageInfo.NormalizedLanguageCode = obj.LanguageInfo.LanguageCode
	lroRet.LanguageInfo.Language = obj.LanguageInfo.LanguageCode
	return s.operations.DoneLRO(ctx, prefix, nil, lroRet)
}

func (s *dataStoreService) DeleteDataStore(ctx context.Context, req *pb.DeleteDataStoreRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDataStoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.DataStore{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%d/locations/%s/collections/%s", name.Project.Number, name.Location, name.Collection)
	return s.operations.DoneLRO(ctx, prefix, nil, nil)
}

func (s *dataStoreService) UpdateDataStore(ctx context.Context, req *pb.UpdateDataStoreRequest) (*pb.DataStore, error) {
	name, err := s.parseDataStoreName(req.DataStore.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.DataStore{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "dataStore %q not found", name)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: support update mask

	proto.Merge(obj, req.GetDataStore())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *dataStoreService) GetDataStore(ctx context.Context, req *pb.GetDataStoreRequest) (*pb.DataStore, error) {
	name, err := s.parseDataStoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.DataStore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "DataStore %v not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

type dataStoreName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
}

func (n *dataStoreName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s", n.Project.Number, n.Location, n.Collection, n.DataStore)
}

func (s *MockService) parseDataStoreName(name string) (*dataStoreName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" {

		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &dataStoreName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			DataStore:  tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}
