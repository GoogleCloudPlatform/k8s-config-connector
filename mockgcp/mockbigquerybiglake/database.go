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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/biglake/v1"
)

func (s *bigLakeService) CreateDatabase(ctx context.Context, req *pb.CreateDatabaseRequest) (*pb.Database, error) {
	reqName := fmt.Sprintf("%s/databases/%s", req.GetParent(), req.GetDatabaseId())
	name, err := s.parseDatabaseName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.CloneOf(req.GetDatabase())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *bigLakeService) UpdateDatabase(ctx context.Context, req *pb.UpdateDatabaseRequest) (*pb.Database, error) {
	reqName := req.GetDatabase().GetName()
	name, err := s.parseDatabaseName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	// TODO: Handle field masks.
	proto.Merge(obj, req.GetDatabase())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *bigLakeService) DeleteDatabase(ctx context.Context, req *pb.DeleteDatabaseRequest) (*pb.Database, error) {
	name, err := s.parseDatabaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.Database{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return deletedObj, nil
}

func (s *bigLakeService) GetDatabase(ctx context.Context, req *pb.GetDatabaseRequest) (*pb.Database, error) {
	name, err := s.parseDatabaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type databaseName struct {
	Project    *projects.ProjectData
	Location   string
	CatalogID  string
	DatabaseID string
}

func (n *databaseName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/catalogs/" + n.CatalogID + "/databases/" + n.DatabaseID
}

// parseDatabaseName parses a string into a databaseName.
// The expected form is `projects/*/locations/*/catalogs/*/databases/*`.
func (s *bigLakeService) parseDatabaseName(name string) (*databaseName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 &&
		tokens[0] == "projects" &&
		tokens[2] == "locations" &&
		tokens[4] == "catalogs" &&
		tokens[6] == "databases" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &databaseName{
			Project:    project,
			Location:   tokens[3],
			CatalogID:  tokens[5],
			DatabaseID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
