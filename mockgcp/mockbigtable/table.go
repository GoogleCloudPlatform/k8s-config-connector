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

package mockbigtable

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

type tableAdminServer struct {
	*MockService
	pb.UnimplementedBigtableTableAdminServer
}

func (s *tableAdminServer) GetTable(ctx context.Context, req *pb.GetTableRequest) (*pb.Table, error) {
	name, err := s.parseTableName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Table not found: %v", name.String())
		}
		return nil, err
	}

	return returnView(obj, req.GetView()), nil
}

func returnView(obj *pb.Table, view pb.Table_View) *pb.Table {
	if view == pb.Table_VIEW_UNSPECIFIED {
		view = pb.Table_SCHEMA_VIEW
	}

	ret := proto.Clone(obj).(*pb.Table)

	for _, columnFamily := range ret.GetColumnFamilies() {
		if proto.Equal(columnFamily.GcRule, &pb.GcRule{}) {
			columnFamily.GcRule = nil
		}

		if proto.Equal(columnFamily.ValueType, &pb.Type{}) {
			columnFamily.ValueType = nil
		}
	}

	return ret
}

func (s *tableAdminServer) ListTables(ctx context.Context, req *pb.ListTablesRequest) (*pb.ListTablesResponse, error) {
	instanceName, err := s.parseInstanceName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListTablesResponse{}
	findKind := (&pb.Table{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: instanceName.String() + "/tables/",
	}, func(obj proto.Message) error {
		table := obj.(*pb.Table)
		response.Tables = append(response.Tables, table)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *tableAdminServer) CreateTable(ctx context.Context, req *pb.CreateTableRequest) (*pb.Table, error) {
	reqName := req.GetParent() + "/tables/" + req.GetTableId()
	tableName, err := s.parseTableName(reqName)
	if err != nil {
		return nil, err
	}

	tableFQN := tableName.String()

	obj := proto.Clone(req.Table).(*pb.Table)
	obj.Name = tableFQN

	if obj.Granularity == pb.Table_TIMESTAMP_GRANULARITY_UNSPECIFIED {
		obj.Granularity = pb.Table_MILLIS
	}

	for _, columnFamily := range obj.GetColumnFamilies() {
		if columnFamily.ValueType == nil {
			columnFamily.ValueType = &pb.Type{}
		}
	}

	if err := s.storage.Create(ctx, tableFQN, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *tableAdminServer) ModifyColumnFamilies(ctx context.Context, req *pb.ModifyColumnFamiliesRequest) (*pb.Table, error) {
	tableName := req.GetName()

	name, err := s.parseTableName(tableName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Table{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, modReq := range req.GetModifications() {
		id := modReq.GetId()

		switch mod := modReq.Mod.(type) {
		case *pb.ModifyColumnFamiliesRequest_Modification_Create:
			// Fail if already exists
			_, exists := obj.ColumnFamilies[id]
			if exists {
				return nil, status.Errorf(codes.AlreadyExists, "column family %q already exists", id)
			}
			obj.ColumnFamilies[id] = mod.Create
		case *pb.ModifyColumnFamiliesRequest_Modification_Drop:
			// Fail if already exists
			_, exists := obj.ColumnFamilies[id]
			if !exists {
				return nil, status.Errorf(codes.NotFound, "column family %q not found", id)
			}
			delete(obj.ColumnFamilies, id)
		default:
			return nil, fmt.Errorf("modified type %T not implemented by mock", mod)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return returnView(obj, pb.Table_VIEW_UNSPECIFIED), nil
}

func (s *tableAdminServer) DeleteTable(ctx context.Context, req *pb.DeleteTableRequest) (*emptypb.Empty, error) {
	name, err := s.parseTableName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Table{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type tableName struct {
	instanceName
	TableName string
}

func (n *tableName) String() string {
	return n.instanceName.String() + "/tables/" + n.TableName
}

// parseTableName parses a string into a tableName.
// The expected form is projects/<projectID>/locations/global/tables/<tableName>
func (s *MockService) parseTableName(name string) (*tableName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[4] == "tables" {
		instanceName, err := s.parseInstanceName(strings.Join(tokens[0:4], "/"))
		if err != nil {
			return nil, err
		}

		name := &tableName{
			instanceName: *instanceName,
			TableName:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
