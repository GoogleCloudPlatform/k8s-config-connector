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

type hiveMetastoreService struct {
	*MockService
	hivepb.UnimplementedHiveMetastoreServiceServer
}

type hiveCatalogName struct {
	Project   *projects.ProjectData
	CatalogID string
}

func (n *hiveCatalogName) String() string {
	return "projects/" + n.Project.ID + "/catalogs/" + n.CatalogID
}

func (s *MockService) parseHiveCatalogName(name string) (*hiveCatalogName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "catalogs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &hiveCatalogName{
			Project:   project,
			CatalogID: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *hiveMetastoreService) CreateHiveCatalog(ctx context.Context, req *hivepb.CreateHiveCatalogRequest) (*hivepb.HiveCatalog, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	if req.GetHiveCatalogId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "hive-catalog-id is required")
	}

	name := &hiveCatalogName{
		Project:   project,
		CatalogID: req.GetHiveCatalogId(),
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetHiveCatalog()).(*hivepb.HiveCatalog)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	primaryLocation := req.GetPrimaryLocation()
	if primaryLocation != "" {
		obj.Replicas = []*hivepb.HiveCatalog_Replica{
			{
				Region: primaryLocation,
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	ret := proto.Clone(obj).(*hivepb.HiveCatalog)
	ret.Replicas = nil
	return ret, nil
}

func (s *hiveMetastoreService) GetHiveCatalog(ctx context.Context, req *hivepb.GetHiveCatalogRequest) (*hivepb.HiveCatalog, error) {
	name, err := s.parseHiveCatalogName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &hivepb.HiveCatalog{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive Catalog %s not found.", name.CatalogID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *hiveMetastoreService) UpdateHiveCatalog(ctx context.Context, req *hivepb.UpdateHiveCatalogRequest) (*hivepb.HiveCatalog, error) {
	if req.GetHiveCatalog() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "hive_catalog is required")
	}
	name, err := s.parseHiveCatalogName(req.GetHiveCatalog().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &hivepb.HiveCatalog{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive Catalog %s not found.", name.CatalogID)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = []string{"description", "location_uri"}
	}
	for _, path := range paths {
		switch path {
		case "description":
			existing.Description = req.GetHiveCatalog().GetDescription()
		case "location_uri":
			existing.LocationUri = req.GetHiveCatalog().GetLocationUri()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unsupported update mask path %q", path)
		}
	}
	existing.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	ret := proto.Clone(existing).(*hivepb.HiveCatalog)
	ret.Replicas = nil
	return ret, nil
}

func (s *hiveMetastoreService) DeleteHiveCatalog(ctx context.Context, req *hivepb.DeleteHiveCatalogRequest) (*emptypb.Empty, error) {
	name, err := s.parseHiveCatalogName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &hivepb.HiveCatalog{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Hive Catalog %s not found.", name.CatalogID)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *hiveMetastoreService) ListHiveCatalogs(ctx context.Context, req *hivepb.ListHiveCatalogsRequest) (*hivepb.ListHiveCatalogsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/catalogs/", project.ID)
	response := &hivepb.ListHiveCatalogsResponse{}
	catalogKind := (&hivepb.HiveCatalog{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, catalogKind, storage.ListOptions{Prefix: prefix}, func(m proto.Message) error {
		item := m.(*hivepb.HiveCatalog)
		response.Catalogs = append(response.Catalogs, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
