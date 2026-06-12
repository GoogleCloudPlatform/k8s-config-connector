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

package mockbigtable

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
)

func (s *tableAdminServer) CreateSchemaBundle(ctx context.Context, req *pb.CreateSchemaBundleRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/schemaBundles/" + req.GetSchemaBundleId()
	name, err := s.parseSchemaBundleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.SchemaBundle).(*pb.SchemaBundle)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, "", nil, obj)
}

func (s *tableAdminServer) GetSchemaBundle(ctx context.Context, req *pb.GetSchemaBundleRequest) (*pb.SchemaBundle, error) {
	name, err := s.parseSchemaBundleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SchemaBundle{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *tableAdminServer) UpdateSchemaBundle(ctx context.Context, req *pb.UpdateSchemaBundleRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseSchemaBundleName(req.GetSchemaBundle().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SchemaBundle{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply update mask
	// SchemaBundle mainly updates the Type field. We can replace the fields specified in the mask.
	// For simplicity, we just copy the Type field as it's the only one currently updated.
	obj.Type = req.GetSchemaBundle().GetType()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.DoneLRO(ctx, "", nil, obj)
}

func (s *tableAdminServer) DeleteSchemaBundle(ctx context.Context, req *pb.DeleteSchemaBundleRequest) (*emptypb.Empty, error) {
	name, err := s.parseSchemaBundleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.SchemaBundle{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type schemaBundleName struct {
	tableName
	SchemaBundleID string
}

func (n *schemaBundleName) String() string {
	return n.tableName.String() + "/schemaBundles/" + n.SchemaBundleID
}

func (s *MockService) parseSchemaBundleName(name string) (*schemaBundleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[6] == "schemaBundles" {
		tableName, err := s.parseTableName(strings.Join(tokens[0:6], "/"))
		if err != nil {
			return nil, err
		}

		name := &schemaBundleName{
			tableName:      *tableName,
			SchemaBundleID: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
