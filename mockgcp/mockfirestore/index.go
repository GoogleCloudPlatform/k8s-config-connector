/*
Copyright 2024 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockfirestore

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
)

func (s *firestoreAdminServer) UpdateField(ctx context.Context, req *pb.UpdateFieldRequest) (*longrunningpb.Operation, error) {
	log := klog.FromContext(ctx)
	log.Info("UpdateField", "req", req)

	name, err := s.parseFieldName(req.GetField().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	if req.GetField().GetIndexConfig() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Field.IndexConfig must be specified")
	}

	for _, index := range req.GetField().GetIndexConfig().GetIndexes() {
		if index.GetQueryScope() == pb.Index_QUERY_SCOPE_UNSPECIFIED {
			return nil, status.Errorf(codes.InvalidArgument, "Index.QueryScope must be specified")
		}
		if len(index.GetFields()) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "Index.Fields must be specified")
		}
		for _, indexField := range index.GetFields() {
			if indexField.GetFieldPath() == "" {
				indexField.FieldPath = name.FieldID
				// return nil, status.Errorf(codes.InvalidArgument, "Index.IndexField.FieldPath must be specified")
			}
			if indexField.GetValueMode() == nil {
				return nil, status.Errorf(codes.InvalidArgument, "Index.IndexField.ValueMode must be specified")
			}
		}
	}
	obj := &pb.Field{}

	create := false
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// We auto-create the field here
			create = true

			// Default indexes (?)
			obj.IndexConfig = &pb.Field_IndexConfig{}
			obj.IndexConfig.Indexes = []*pb.Index{
				{
					QueryScope: pb.Index_COLLECTION,
					Fields: []*pb.Index_IndexField{
						{
							FieldPath: name.FieldID,
							ValueMode: &pb.Index_IndexField_Order_{Order: pb.Index_IndexField_ASCENDING},
						},
					},
				},
				{
					QueryScope: pb.Index_COLLECTION,
					Fields: []*pb.Index_IndexField{
						{
							FieldPath: name.FieldID,
							ValueMode: &pb.Index_IndexField_Order_{Order: pb.Index_IndexField_DESCENDING},
						},
					},
				},
				{
					QueryScope: pb.Index_COLLECTION,
					Fields: []*pb.Index_IndexField{
						{
							FieldPath: name.FieldID,
							ValueMode: &pb.Index_IndexField_ArrayConfig_{ArrayConfig: pb.Index_IndexField_CONTAINS},
						},
					},
				},
			}
		} else {
			return nil, err
		}
	}

	oldObj := ProtoClone(obj)

	// TODO: Merge more fields?

	indexConfigDeltas := []*pb.FieldOperationMetadata_IndexConfigDelta{}

	var newIndexes []*pb.Index
	// First go through and remove any indexes we are not keeping
	for _, oldIndex := range oldObj.GetIndexConfig().GetIndexes() {
		keepingIndex := false

		for _, newIndex := range req.GetField().GetIndexConfig().GetIndexes() {
			if indexEquals(oldIndex, newIndex) {
				keepingIndex = true
				break
			}
		}

		if keepingIndex {
			continue
		}
		// Removing an existing index
		indexConfigDelta := &pb.FieldOperationMetadata_IndexConfigDelta{
			ChangeType: pb.FieldOperationMetadata_IndexConfigDelta_REMOVE,
			Index:      ProtoClone(oldIndex),
		}
		indexConfigDelta.Index.State = pb.Index_STATE_UNSPECIFIED // Does not return state
		indexConfigDeltas = append(indexConfigDeltas, indexConfigDelta)
	}

	// Now go through and add any new indexes
	for _, index := range req.GetField().GetIndexConfig().GetIndexes() {
		foundMatch := false

		for _, oldIndex := range oldObj.GetIndexConfig().GetIndexes() {
			if indexEquals(oldIndex, index) {
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			indexConfigDeltas = append(indexConfigDeltas, &pb.FieldOperationMetadata_IndexConfigDelta{
				ChangeType: pb.FieldOperationMetadata_IndexConfigDelta_ADD,
				Index:      ProtoClone(index),
			})
		}

		newIndexes = append(newIndexes, index)
	}

	if obj.IndexConfig == nil {
		obj.IndexConfig = &pb.Field_IndexConfig{}
	}
	if obj.IndexConfig.AncestorField == "" {
		obj.IndexConfig.AncestorField = fmt.Sprintf("projects/%s/databases/%s/collectionGroups/__default__/fields/*", name.ProjectID, name.DatabaseID)
	}
	obj.IndexConfig.Indexes = newIndexes

	for _, index := range obj.GetIndexConfig().GetIndexes() {
		index.State = pb.Index_READY
	}

	if create {
		obj.Name = fqn
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	lroPrefix := fmt.Sprintf("projects/%s/databases/%s", name.ProjectID, name.DatabaseID)
	lroMetadata := &pb.FieldOperationMetadata{
		Field:             name.String(),
		IndexConfigDeltas: indexConfigDeltas,
		StartTime:         timestamppb.Now(),
		State:             pb.OperationState_INITIALIZING,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.State = pb.OperationState_SUCCESSFUL
		lroMetadata.EndTime = timestamppb.Now()
		lroMetadata.ProgressDocuments = &pb.Progress{}
		return obj, nil
	})
}

func fieldNames(a []*pb.Index_IndexField) []string {
	var result []string
	for _, f := range a {
		result = append(result, f.GetFieldPath())
	}
	return result
}

func indexEquals(a, b *pb.Index) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.QueryScope != b.QueryScope {
		return false
	}
	if a.Multikey != b.Multikey {
		return false
	}
	if !slices.Equal(fieldNames(a.Fields), fieldNames(b.Fields)) {
		return false
	}

	aFields := make(map[string]*pb.Index_IndexField)
	for _, f := range a.Fields {
		aFields[f.FieldPath] = f
	}
	for _, bField := range b.Fields {
		if aField, ok := aFields[bField.FieldPath]; !ok || !proto.Equal(bField, aField) {
			return false
		}
	}
	return true
}

func (s *firestoreAdminServer) GetField(ctx context.Context, req *pb.GetFieldRequest) (*pb.Field, error) {
	name, err := s.parseFieldName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Field{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// fieldName is the resource name of a firestore field.
// Format: projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_id}
type fieldName struct {
	ProjectID    string
	DatabaseID   string
	CollectionID string
	FieldID      string
}

func (n *fieldName) String() string {
	return "projects/" + n.ProjectID + "/databases/" + n.DatabaseID + "/collectionGroups/" + n.CollectionID + "/fields/" + n.FieldID
}

// parsefieldName parses the name of a field.
func (s *firestoreAdminServer) parseFieldName(name string) (*fieldName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "databases" && tokens[4] == "collectionGroups" && tokens[6] == "fields" {
		return &fieldName{
			ProjectID:    tokens[1],
			DatabaseID:   tokens[3],
			CollectionID: tokens[5],
			FieldID:      tokens[7],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q does not match expected format", name)
	}
}
