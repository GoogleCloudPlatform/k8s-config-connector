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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/bigquery/biglake/v1"
)

func (s *bigLakeService) CreateDatabase(ctx context.Context, req *pb.CreateDatabaseRequest) (*pb.Database, error) {
	name, err := s.parseDatabaseName(req.GetParent() + "/databases/" + req.GetDatabaseId())
	if err != nil {
		return nil, err
	}

	obj := req.GetDatabase()
	obj.Name = name.String()

	now := time.Now().UnixNano()
	obj.CreateTime = &timestamppb.Timestamp{Seconds: now / 1e9, Nanos: int32(now % 1e9)}
	obj.UpdateTime = &timestamppb.Timestamp{Seconds: now / 1e9, Nanos: int32(now % 1e9)}

	if err := s.storage.Create(ctx, obj.Name, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *bigLakeService) GetDatabase(ctx context.Context, req *pb.GetDatabaseRequest) (*pb.Database, error) {
	name, err := s.parseDatabaseName(req.GetName())
	if err != nil {
		return nil, err
	}

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *bigLakeService) UpdateDatabase(ctx context.Context, req *pb.UpdateDatabaseRequest) (*pb.Database, error) {
	name, err := s.parseDatabaseName(req.GetDatabase().GetName())
	if err != nil {
		return nil, err
	}

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	// TODO: Handle field masks.
	proto.Merge(obj, req.GetDatabase())

	obj.Name = name.String()

	now := time.Now().UnixNano()
	obj.UpdateTime = &timestamppb.Timestamp{Seconds: now / 1e9, Nanos: int32(now % 1e9)}

	if err := s.storage.Update(ctx, obj.Name, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *bigLakeService) DeleteDatabase(ctx context.Context, req *pb.DeleteDatabaseRequest) (*pb.Database, error) {
	name, err := s.parseDatabaseName(req.GetName())
	if err != nil {
		return nil, err
	}

	obj := &pb.Database{}
	if err := s.storage.Delete(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type databaseName struct {
	Project  string
	Location string
	Catalog  string
	Database string
}

func (n *databaseName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/catalogs/%s/databases/%s", n.Project, n.Location, n.Catalog, n.Database)
}

func (s *bigLakeService) parseDatabaseName(name string) (*databaseName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "catalogs" && tokens[6] == "databases" {
		return &databaseName{
			Project:  tokens[1],
			Location: tokens[3],
			Catalog:  tokens[5],
			Database: tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
