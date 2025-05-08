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
// proto.service: google.cloud.bigquery.biglake.v1.MetastoreService
// proto.message: google.cloud.bigquery.biglake.v1.Catalog

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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *bigLakeService) CreateCatalog(ctx context.Context, req *pb.CreateCatalogRequest) (*pb.Catalog, error) {
	reqName := fmt.Sprintf("%s/catalogs/%s", req.GetParent(), req.GetCatalogId())
	name, err := s.parseCatalogName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetCatalog()).(*pb.Catalog)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *bigLakeService) DeleteCatalog(ctx context.Context, req *pb.DeleteCatalogRequest) (*pb.Catalog, error) {
	name, err := s.parseCatalogName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Catalog{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	// The real API returns the deleted object, although the RPC docs say Empty.
	// Set delete/expire times like the real API.
	now := time.Now()
	deletedObj.DeleteTime = timestamppb.New(now)
	deletedObj.ExpireTime = timestamppb.New(now.Add(time.Hour * 24 * 30)) // 30 days expiration?
	return deletedObj, nil
}

func (s *bigLakeService) GetCatalog(ctx context.Context, req *pb.GetCatalogRequest) (*pb.Catalog, error) {
	name, err := s.parseCatalogName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Catalog{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *bigLakeService) ListCatalogs(ctx context.Context, req *pb.ListCatalogsRequest) (*pb.ListCatalogsResponse, error) {
	response := &pb.ListCatalogsResponse{}

	findPrefix := req.GetParent()
	catalogKind := (&pb.Catalog{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, catalogKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		catalog := obj.(*pb.Catalog)
		response.Catalogs = append(response.Catalogs, catalog)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

type catalogName struct {
	Project  *projects.ProjectData
	Location string
	Catalog  string
}

func (n *catalogName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/catalogs/" + n.Catalog
}

// parseCatalogName parses a string into a catalogName.
// The expected form is `projects/*/locations/*/catalogs/*`.
func (s *bigLakeService) parseCatalogName(name string) (*catalogName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 &&
		tokens[0] == "projects" &&
		tokens[2] == "locations" &&
		tokens[4] == "catalogs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &catalogName{
			Project:  project,
			Location: tokens[3],
			Catalog:  tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
