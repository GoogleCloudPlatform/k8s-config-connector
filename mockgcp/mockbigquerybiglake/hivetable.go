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

package mockbigquerybiglake

import (
	"context"
	"fmt"
	"strings"
	"time"

	hivepb "cloud.google.com/go/biglake/hive/apiv1/hivepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type hiveTableName struct {
	Project    *projects.ProjectData
	CatalogID  string
	DatabaseID string
	TableID    string
}

func (n *hiveTableName) String() string {
	return "projects/" + n.Project.ID + "/catalogs/" + n.CatalogID + "/databases/" + n.DatabaseID + "/tables/" + n.TableID
}

func (s *MockService) parseHiveTableName(name string) (*hiveTableName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "catalogs" && tokens[4] == "databases" && tokens[6] == "tables" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &hiveTableName{
			Project:    project,
			CatalogID:  tokens[3],
			DatabaseID: tokens[5],
			TableID:    tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *hiveMetastoreService) CreateHiveTable(ctx context.Context, req *hivepb.CreateHiveTableRequest) (*hivepb.HiveTable, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "catalogs" || tokens[4] != "databases" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	databaseFQN := "projects/" + project.ID + "/catalogs/" + tokens[3] + "/databases/" + tokens[5]
	database := &hivepb.HiveDatabase{}
	if err := s.storage.Get(ctx, databaseFQN, database); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive database %s.%s does not exist", tokens[3], tokens[5])
		}
		return nil, err
	}

	if req.GetHiveTableId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "hive-table-id is required")
	}

	name := &hiveTableName{
		Project:    project,
		CatalogID:  tokens[3],
		DatabaseID: tokens[5],
		TableID:    req.GetHiveTableId(),
	}

	fqn := name.String()

	existing := &hivepb.HiveTable{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Hive table %s already exists", fqn)
	}

	now := time.Now()
	obj := proto.Clone(req.GetHiveTable()).(*hivepb.HiveTable)
	if obj == nil {
		obj = &hivepb.HiveTable{}
	}
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.TableType = "EXTERNAL_TABLE"

	if obj.Parameters == nil {
		obj.Parameters = make(map[string]string)
	}
	if _, ok := obj.Parameters["table_type"]; !ok {
		obj.Parameters["table_type"] = "HIVE"
	}

	if obj.StorageDescriptor == nil {
		obj.StorageDescriptor = &hivepb.StorageDescriptor{}
	}
	if obj.StorageDescriptor.LocationUri == "" && database.LocationUri != "" {
		obj.StorageDescriptor.LocationUri = strings.TrimSuffix(database.LocationUri, "/") + "/" + name.TableID
	}
	if obj.StorageDescriptor.Compressed != nil && !*obj.StorageDescriptor.Compressed {
		obj.StorageDescriptor.Compressed = nil
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return proto.Clone(obj).(*hivepb.HiveTable), nil
}

func (s *hiveMetastoreService) GetHiveTable(ctx context.Context, req *hivepb.GetHiveTableRequest) (*hivepb.HiveTable, error) {
	name, err := s.parseHiveTableName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &hivepb.HiveTable{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive table: %s.%s.%s does not exist", name.CatalogID, name.DatabaseID, name.TableID)
		}
		return nil, err
	}

	return proto.Clone(obj).(*hivepb.HiveTable), nil
}

func (s *hiveMetastoreService) UpdateHiveTable(ctx context.Context, req *hivepb.UpdateHiveTableRequest) (*hivepb.HiveTable, error) {
	if req.GetHiveTable() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "hive_table is required")
	}
	name, err := s.parseHiveTableName(req.GetHiveTable().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &hivepb.HiveTable{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive table: %s.%s.%s does not exist", name.CatalogID, name.DatabaseID, name.TableID)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = []string{"description", "storage_descriptor", "partition_keys", "parameters"}
	}
	for _, path := range paths {
		switch path {
		case "description":
			existing.Description = req.GetHiveTable().GetDescription()
		case "storage_descriptor", "storageDescriptor":
			existing.StorageDescriptor = req.GetHiveTable().GetStorageDescriptor()
			if existing.StorageDescriptor != nil && existing.StorageDescriptor.Compressed != nil && !*existing.StorageDescriptor.Compressed {
				existing.StorageDescriptor.Compressed = nil
			}
		case "partition_keys", "partitionKeys":
			existing.PartitionKeys = req.GetHiveTable().GetPartitionKeys()
		case "parameters":
			existing.Parameters = req.GetHiveTable().GetParameters()
			if existing.Parameters == nil {
				existing.Parameters = make(map[string]string)
			}
			if _, ok := existing.Parameters["table_type"]; !ok {
				existing.Parameters["table_type"] = "HIVE"
			}
		case "view_original_text", "viewOriginalText":
			existing.ViewOriginalText = req.GetHiveTable().GetViewOriginalText()
		case "view_expanded_text", "viewExpandedText":
			existing.ViewExpandedText = req.GetHiveTable().GetViewExpandedText()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unsupported update mask path %q", path)
		}
	}
	existing.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	return proto.Clone(existing).(*hivepb.HiveTable), nil
}

func (s *hiveMetastoreService) DeleteHiveTable(ctx context.Context, req *hivepb.DeleteHiveTableRequest) (*emptypb.Empty, error) {
	name, err := s.parseHiveTableName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &hivepb.HiveTable{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive table: %s.%s.%s does not exist", name.CatalogID, name.DatabaseID, name.TableID)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *hiveMetastoreService) ListHiveTables(ctx context.Context, req *hivepb.ListHiveTablesRequest) (*hivepb.ListHiveTablesResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "catalogs" || tokens[4] != "databases" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/catalogs/%s/databases/%s/tables/", project.ID, tokens[3], tokens[5])
	response := &hivepb.ListHiveTablesResponse{}
	tableKind := (&hivepb.HiveTable{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, tableKind, storage.ListOptions{Prefix: prefix}, func(m proto.Message) error {
		item := m.(*hivepb.HiveTable)
		response.Tables = append(response.Tables, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
