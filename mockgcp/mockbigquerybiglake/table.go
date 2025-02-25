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
// proto.service: mockgcp.cloud.bigquery.biglake.v1.Table
// proto.message: google.cloud.bigquery.biglake.v1.Table

package mockbigquerybiglake

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/biglake/v1"
)

type bigLakeService struct {
	*MockService
	pb.UnimplementedMetastoreServiceServer
}

func (s *bigLakeService) CreateTable(ctx context.Context, req *pb.CreateTableRequest) (*pb.Table, error) {
	reqName := fmt.Sprintf("%s/tables/%s", req.GetParent(), req.GetTableId())
	name, err := s.parseTableName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetTable()).(*pb.Table)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *bigLakeService) UpdateTable(ctx context.Context, req *pb.UpdateTableRequest) (*pb.Table, error) {
	reqName := req.GetTable().GetName()
	name, err := s.parseTableName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	// TODO: Handle field masks.
	proto.Merge(obj, req.GetTable())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *bigLakeService) DeleteTable(ctx context.Context, req *pb.DeleteTableRequest) (*pb.Table, error) {
	name, err := s.parseTableName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Table{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return deletedObj, nil
}

func (s *bigLakeService) GetTable(ctx context.Context, req *pb.GetTableRequest) (*pb.Table, error) {
	name, err := s.parseTableName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type tableName struct {
	Project    *projects.ProjectData
	Location   string
	CatalogID  string
	DatabaseID string
	Table      string
}

func (n *tableName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/catalogs/" + n.CatalogID + "/databases/" + n.DatabaseID + "/tables/" + n.Table
}

// parseTableName parses a string into a tableName.
// The expected form is `projects/*/locations/*/catalogs/*/databases/*/tables/*`.
func (s *bigLakeService) parseTableName(name string) (*tableName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 &&
		tokens[0] == "projects" &&
		tokens[2] == "locations" &&
		tokens[4] == "catalogs" &&
		tokens[6] == "databases" &&
		tokens[8] == "tables" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &tableName{
			Project:    project,
			Location:   tokens[3],
			CatalogID:  tokens[5],
			DatabaseID: tokens[7],
			Table:      tokens[9],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
