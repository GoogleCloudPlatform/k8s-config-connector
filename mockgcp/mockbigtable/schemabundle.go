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

// +tool:mockgcp-support
// proto.service: google.bigtable.admin.v2.BigtableTableAdmin
// proto.message: google.bigtable.admin.v2.SchemaBundle

package mockbigtable

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

func (s *tableAdminServer) GetSchemaBundle(ctx context.Context, req *pb.GetSchemaBundleRequest) (*pb.SchemaBundle, error) {
	name, err := s.parseSchemaBundleName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SchemaBundle{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Schema bundle %s not found.", name.String())
		}
		return nil, err
	}

	return obj, nil
}

func (s *tableAdminServer) CreateSchemaBundle(ctx context.Context, req *pb.CreateSchemaBundleRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/schemaBundles/" + req.GetSchemaBundleId()
	name, err := s.parseSchemaBundleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.GetSchemaBundle())
	obj.Name = fqn
	obj.Etag = "abcdef0123A="

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	clusters, err := s.listClustersForInstance(ctx, &name.instanceName)
	if err != nil {
		return nil, err
	}
	zone := pickZoneForInstanceOperation(clusters)
	if zone == "" {
		zone = "us-central1-a"
	}
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)

	metadata := &pb.CreateSchemaBundleMetadata{
		Name:      name.String(),
		StartTime: timestamppb.Now(),
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now().Add(5 * time.Minute))
		return obj, nil
	})
}

func (s *tableAdminServer) UpdateSchemaBundle(ctx context.Context, req *pb.UpdateSchemaBundleRequest) (*longrunningpb.Operation, error) {
	reqBundle := req.GetSchemaBundle()
	if reqBundle == nil {
		return nil, status.Errorf(codes.InvalidArgument, "schema_bundle must be specified")
	}

	name, err := s.parseSchemaBundleName(reqBundle.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.SchemaBundle{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Schema bundle %s not found.", name.String())
		}
		return nil, err
	}

	updated := proto.CloneOf(existing)

	mask := req.GetUpdateMask()
	if mask == nil || len(mask.GetPaths()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be specified")
	}

	for _, path := range mask.GetPaths() {
		switch path {
		case "proto_schema", "protoSchema":
			if protoSchema := reqBundle.GetProtoSchema(); protoSchema != nil {
				updated.Type = &pb.SchemaBundle_ProtoSchema{ProtoSchema: protoSchema}
			} else {
				updated.Type = nil
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q not supported for update", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	clusters, err := s.listClustersForInstance(ctx, &name.instanceName)
	if err != nil {
		return nil, err
	}
	zone := pickZoneForInstanceOperation(clusters)
	if zone == "" {
		zone = "us-central1-a"
	}
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)

	metadata := &pb.UpdateSchemaBundleMetadata{
		Name:      name.String(),
		StartTime: timestamppb.Now(),
	}

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now().Add(5 * time.Minute))
		return updated, nil
	})
}

func (s *tableAdminServer) DeleteSchemaBundle(ctx context.Context, req *pb.DeleteSchemaBundleRequest) (*emptypb.Empty, error) {
	name, err := s.parseSchemaBundleName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.SchemaBundle{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Schema bundle %s not found.", name.String())
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *tableAdminServer) ListSchemaBundles(ctx context.Context, req *pb.ListSchemaBundlesRequest) (*pb.ListSchemaBundlesResponse, error) {
	tableName, err := s.parseTableName(req.GetParent())
	if err != nil {
		return nil, err
	}

	response := &pb.ListSchemaBundlesResponse{}
	findKind := (&pb.SchemaBundle{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: tableName.String() + "/schemaBundles/",
	}, func(obj proto.Message) error {
		schemaBundle := obj.(*pb.SchemaBundle)
		response.SchemaBundles = append(response.SchemaBundles, schemaBundle)
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(response.SchemaBundles, func(i, j int) bool {
		return response.SchemaBundles[i].Name < response.SchemaBundles[j].Name
	})

	return response, nil
}

type schemaBundleName struct {
	tableName
	SchemaBundleId string
}

func (n *schemaBundleName) String() string {
	return n.tableName.String() + "/schemaBundles/" + n.SchemaBundleId
}

// parseSchemaBundleName parses a string into a schemaBundleName.
// The expected form is projects/*/instances/*/tables/*/schemaBundles/*
func (s *MockService) parseSchemaBundleName(name string) (*schemaBundleName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "tables" && tokens[6] == "schemaBundles" {
		tableName, err := s.parseTableName(strings.Join(tokens[0:6], "/"))
		if err != nil {
			return nil, err
		}

		return &schemaBundleName{
			tableName:      *tableName,
			SchemaBundleId: tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
