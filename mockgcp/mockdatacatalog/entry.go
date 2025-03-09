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
// proto.service: google.cloud.datacatalog.v1.DataCatalog
// proto.message: google.cloud.datacatalog.v1.Entry

package mockdatacatalog

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datacatalog/v1"
)

type DataCatalogV1 struct {
	*MockService
	pb.UnimplementedDataCatalogServer
}

func (s *DataCatalogV1) GetEntry(ctx context.Context, req *pb.GetEntryRequest) (*pb.Entry, error) {
	name, err := s.parseEntryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Entry{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Entry %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) CreateEntry(ctx context.Context, req *pb.CreateEntryRequest) (*pb.Entry, error) {
	reqName := fmt.Sprintf("%s/entries/%s", req.GetParent(), req.GetEntryId())
	name, err := s.parseEntryName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetEntry()).(*pb.Entry)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *DataCatalogV1) UpdateEntry(ctx context.Context, req *pb.UpdateEntryRequest) (*pb.Entry, error) {
	name, err := s.parseEntryName(req.GetEntry().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Entry{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: support update mask

	proto.Merge(obj, req.GetEntry())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *DataCatalogV1) DeleteEntry(ctx context.Context, req *pb.DeleteEntryRequest) (*emptypb.Empty, error) {
	name, err := s.parseEntryName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Entry{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type entryName struct {
	Project    string
	Location   string
	EntryGroup string
	EntryName  string
}

func (n *entryName) String() string {
	return "projects/" + n.Project + "/locations/" + n.Location + "/entryGroups/" + n.EntryGroup + "/entries/" + n.EntryName
}

func (s *MockService) parseEntryName(name string) (*entryName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "entryGroups" && tokens[6] == "entries" {

		name := &entryName{
			Project:    tokens[1],
			Location:   tokens[3],
			EntryGroup: tokens[5],
			EntryName:  tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
