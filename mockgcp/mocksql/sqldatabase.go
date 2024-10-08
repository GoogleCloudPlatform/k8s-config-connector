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

package mocksql

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Implemented SqlDatabasesServiceServer interface in mockgcp/generated/mockgcp/cloud/sql/v1beta4/cloud_sql_grpc.pb.go
type sqlDatabaseServer struct {
	*MockService
	pb.UnimplementedSqlDatabasesServiceServer
}

func (s *sqlDatabaseServer) Insert(ctx context.Context, req *pb.SqlDatabasesInsertRequest) (*pb.Operation, error) {
	name, err := s.buildDatabaseName(req.GetProject(), req.GetInstance(), req.GetBody().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetBody()).(*pb.Database)
	obj.Name = name.DatabaseName
	obj.Project = name.Project.ID
	obj.Kind = "sql#database"
	obj.Collation = "utf8_general_ci"
	obj.Etag = fields.ComputeWeakEtag(obj)

	obj.SelfLink = fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s/databases/%s",
		name.Project.ID, name.InstanceName, name.DatabaseName)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_CREATE_DATABASE,
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlDatabaseServer) Get(ctx context.Context, req *pb.SqlDatabasesGetRequest) (*pb.Database, error) {
	name, err := s.buildDatabaseName(req.GetProject(), req.GetInstance(), req.GetDatabase())
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

func (s *sqlDatabaseServer) Update(ctx context.Context, req *pb.SqlDatabasesUpdateRequest) (*pb.Operation, error) {
	name, err := s.buildDatabaseName(req.GetProject(), req.GetInstance(), req.GetDatabase())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Collation = req.GetBody().Collation
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_UPDATE_DATABASE,
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlDatabaseServer) Delete(ctx context.Context, req *pb.SqlDatabasesDeleteRequest) (*pb.Operation, error) {
	name, err := s.buildDatabaseName(req.GetProject(), req.GetInstance(), req.GetDatabase())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Database{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_DELETE_DATABASE,
	}

	return s.operations.startLRO(ctx, op, deleted, func() (proto.Message, error) {
		return deleted, nil
	})
}

type DatabaseName struct {
	Project      *projects.ProjectData
	InstanceName string
	DatabaseName string
}

func (n *DatabaseName) String() string {
	return "projects/" + n.Project.ID + "/instances/" + n.InstanceName + "/databases/" + n.DatabaseName
}

// parseDatabaseName parses a string into a DatabaseName.
// The expected form is projects/<projectID>/instances/<SQLInstanceName>/databases/<DatabaseName>
func (s *MockService) parseDatabaseName(name string) (*DatabaseName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "databases" {
		return s.buildDatabaseName(tokens[1], tokens[3], tokens[5])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) buildDatabaseName(projectID, instanceName string, databaseName string) (*DatabaseName, error) {
	project, err := s.projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}
	return &DatabaseName{
		Project:      project,
		InstanceName: instanceName,
		DatabaseName: databaseName,
	}, nil
}
