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

type hiveDatabaseName struct {
	Project    *projects.ProjectData
	CatalogID  string
	DatabaseID string
}

func (n *hiveDatabaseName) String() string {
	return "projects/" + n.Project.ID + "/catalogs/" + n.CatalogID + "/databases/" + n.DatabaseID
}

func (s *MockService) parseHiveDatabaseName(name string) (*hiveDatabaseName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "catalogs" && tokens[4] == "databases" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &hiveDatabaseName{
			Project:    project,
			CatalogID:  tokens[3],
			DatabaseID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *hiveMetastoreService) CreateHiveDatabase(ctx context.Context, req *hivepb.CreateHiveDatabaseRequest) (*hivepb.HiveDatabase, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "catalogs" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	catalogFQN := "projects/" + project.ID + "/catalogs/" + tokens[3]
	catalog := &hivepb.HiveCatalog{}
	if err := s.storage.Get(ctx, catalogFQN, catalog); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive Catalog %s not found.", tokens[3])
		}
		return nil, err
	}

	if req.GetHiveDatabaseId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "hive-database-id is required")
	}

	name := &hiveDatabaseName{
		Project:    project,
		CatalogID:  tokens[3],
		DatabaseID: req.GetHiveDatabaseId(),
	}

	fqn := name.String()

	existing := &hivepb.HiveDatabase{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Hive database %s already exists", fqn)
	}

	now := time.Now()
	obj := proto.Clone(req.GetHiveDatabase()).(*hivepb.HiveDatabase)
	if obj == nil {
		obj = &hivepb.HiveDatabase{}
	}
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.LocationUri == "" && catalog.LocationUri != "" {
		obj.LocationUri = strings.TrimSuffix(catalog.LocationUri, "/") + "/" + name.DatabaseID + ".db"
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return proto.Clone(obj).(*hivepb.HiveDatabase), nil
}

func (s *hiveMetastoreService) GetHiveDatabase(ctx context.Context, req *hivepb.GetHiveDatabaseRequest) (*hivepb.HiveDatabase, error) {
	name, err := s.parseHiveDatabaseName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &hivepb.HiveDatabase{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive database %s.%s does not exist", name.CatalogID, name.DatabaseID)
		}
		return nil, err
	}

	return proto.Clone(obj).(*hivepb.HiveDatabase), nil
}

func (s *hiveMetastoreService) UpdateHiveDatabase(ctx context.Context, req *hivepb.UpdateHiveDatabaseRequest) (*hivepb.HiveDatabase, error) {
	if req.GetHiveDatabase() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "hive_database is required")
	}
	name, err := s.parseHiveDatabaseName(req.GetHiveDatabase().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &hivepb.HiveDatabase{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive database %s.%s does not exist", name.CatalogID, name.DatabaseID)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = []string{"description", "location_uri", "parameters"}
	}
	for _, path := range paths {
		switch path {
		case "description":
			existing.Description = req.GetHiveDatabase().GetDescription()
		case "location_uri":
			existing.LocationUri = req.GetHiveDatabase().GetLocationUri()
		case "parameters":
			existing.Parameters = req.GetHiveDatabase().GetParameters()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unsupported update mask path %q", path)
		}
	}
	existing.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	return proto.Clone(existing).(*hivepb.HiveDatabase), nil
}

func (s *hiveMetastoreService) DeleteHiveDatabase(ctx context.Context, req *hivepb.DeleteHiveDatabaseRequest) (*emptypb.Empty, error) {
	name, err := s.parseHiveDatabaseName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &hivepb.HiveDatabase{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive database %s.%s does not exist", name.CatalogID, name.DatabaseID)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *hiveMetastoreService) ListHiveDatabases(ctx context.Context, req *hivepb.ListHiveDatabasesRequest) (*hivepb.ListHiveDatabasesResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "catalogs" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/catalogs/%s/databases/", project.ID, tokens[3])
	response := &hivepb.ListHiveDatabasesResponse{}
	databaseKind := (&hivepb.HiveDatabase{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, databaseKind, storage.ListOptions{Prefix: prefix}, func(m proto.Message) error {
		item := m.(*hivepb.HiveDatabase)
		response.Databases = append(response.Databases, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
