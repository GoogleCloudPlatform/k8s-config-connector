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

package mockfirestore

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/uuid"

	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *firestoreAdminServer) CreateIndex(ctx context.Context, req *pb.CreateIndexRequest) (*longrunningpb.Operation, error) {
	// Server-generated ID
	indexID := string(uuid.NewUUID())

	name, err := s.parseIndexName(req.Parent + "/indexes/" + indexID)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.Index).(*pb.Index)
	obj.Name = fqn
	obj.State = pb.Index_READY

	s.populateDefaultsForIndex(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.IndexOperationMetadata{
		Index:     fqn,
		StartTime: timestamppb.Now(),
		State:     pb.OperationState_INITIALIZING,
	}
	lroPrefix := fmt.Sprintf("projects/%s/databases/%s", name.Project.ID, name.Database)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		result := ProtoClone(obj)
		result.State = pb.Index_CREATING

		lroMetadata.EndTime = timestamppb.Now()
		lroMetadata.State = pb.OperationState_SUCCESSFUL
		lroMetadata.ProgressDocuments = &pb.Progress{}
		return result, nil
	})
}

func (s *firestoreAdminServer) populateDefaultsForIndex(obj *pb.Index) {
	if obj.Density == pb.Index_DENSITY_UNSPECIFIED {
		obj.Density = pb.Index_SPARSE_ALL
	}

	// From https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.collectionGroups.indexes
	// The last field entry is always for the field path __name__.
	// If, on creation, __name__ was not specified as the last field, it will be added automatically with the same direction as that of the last field defined.
	// If the final field in a composite index is not directional, the __name__ will be ordered ASCENDING (unless explicitly specified).
	hasName := false
	for _, field := range obj.Fields {
		if field.FieldPath == "__name__" {
			hasName = true
		}
	}
	if !hasName {
		lastField := obj.Fields[len(obj.Fields)-1]
		direction := lastField.GetOrder()
		if direction == pb.Index_IndexField_ORDER_UNSPECIFIED {
			direction = pb.Index_IndexField_ASCENDING
		}
		obj.Fields = append(obj.Fields, &pb.Index_IndexField{
			FieldPath: "__name__",
			ValueMode: &pb.Index_IndexField_Order_{Order: direction},
		})
	}
}

func (s *firestoreAdminServer) GetIndex(ctx context.Context, req *pb.GetIndexRequest) (*pb.Index, error) {
	name, err := s.parseIndexName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Index{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *firestoreAdminServer) ListIndexes(ctx context.Context, req *pb.ListIndexesRequest) (*pb.ListIndexesResponse, error) {
	// This is a simplified implementation. A real implementation would handle pagination.
	var indexes []*pb.Index
	parentPath := req.Parent + "/indexes/"
	// Special-case: the collectionGroup "-" means "all collection groups"
	if strings.HasSuffix(parentPath, "/collectionGroups/-/indexes/") {
		parentPath = strings.TrimSuffix(parentPath, "/-/indexes/")
	}
	if err := s.storage.List(ctx, (&pb.Index{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: parentPath,
	}, func(obj proto.Message) error {
		index := obj.(*pb.Index)
		indexes = append(indexes, index)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListIndexesResponse{
		Indexes: indexes,
	}, nil
}

func (s *firestoreAdminServer) DeleteIndex(ctx context.Context, req *pb.DeleteIndexRequest) (*emptypb.Empty, error) {
	name, err := s.parseIndexName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.Index{}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type indexName struct {
	Project         *projects.ProjectData
	Database        string
	CollectionGroup string
	Index           string
}

func (n *indexName) String() string {
	return "projects/" + n.Project.ID + "/databases/" + n.Database + "/collectionGroups/" + n.CollectionGroup + "/indexes/" + n.Index
}

// parseIndexName parses a string into a indexName.
// The expected form is projects/{project}/databases/{database}/collectionGroups/{collectionGroup}/indexes/{index}
func (s *MockService) parseIndexName(name string) (*indexName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "databases" && tokens[4] == "collectionGroups" && tokens[6] == "indexes" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &indexName{
			Project:         project,
			Database:        tokens[3],
			CollectionGroup: tokens[5],
			Index:           tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
