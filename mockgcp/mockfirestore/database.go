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

// +tool:mockgcp-support
// proto.service: google.firestore.admin.v1.FirestoreAdmin
// proto.message: google.firestore.admin.v1.Database

package mockfirestore

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type firestoreAdminServer struct {
	*MockService
	pb.UnimplementedFirestoreAdminServer
}

func (s *firestoreAdminServer) GetDatabase(ctx context.Context, req *pb.GetDatabaseRequest) (*pb.Database, error) {
	name, err := s.parseDatabaseName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Project '%s' or database '%s' does not exist.", name.Project.ID, name.Database)
		}
		return nil, err
	}

	return obj, nil
}

func (s *firestoreAdminServer) CreateDatabase(ctx context.Context, req *pb.CreateDatabaseRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDatabaseName(req.Parent + "/databases/" + req.DatabaseId)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.Database).(*pb.Database)
	populateDefaultsForDatabase(obj)
	t := timestamppb.New(time.Now())
	obj.CreateTime = t
	obj.UpdateTime = t
	obj.EarliestVersionTime = t

	obj.Etag = computeEtag(obj)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreateDatabaseMetadata{}
	lroPrefix := fqn
	op, err := s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Database)
		return result, nil
	})
	if err != nil {
		return op, err
	}

	// Unusually, response is populated right away
	response, err := anypb.New(obj)
	if err != nil {
		return op, err
	}
	op.Result = &longrunningpb.Operation_Response{
		Response: response,
	}

	return op, err
}

func (s *firestoreAdminServer) UpdateDatabase(ctx context.Context, req *pb.UpdateDatabaseRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDatabaseName(req.GetDatabase().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if len(req.GetUpdateMask().GetPaths()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Update mask is required")
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "deleteProtectionState":
			obj.DeleteProtectionState = req.GetDatabase().GetDeleteProtectionState()
		case "concurrencyMode":
			obj.ConcurrencyMode = req.GetDatabase().GetConcurrencyMode()
		case "pointInTimeRecoveryEnablement":
			obj.PointInTimeRecoveryEnablement = req.GetDatabase().GetPointInTimeRecoveryEnablement()
		case "name":
			// Not updatable
			if req.GetDatabase().GetName() != obj.Name {
				return nil, status.Errorf(codes.InvalidArgument, "Field %q is not updatable (in mockgcp)", path)
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "Field %q is not updatable (in mockgcp)", path)
		}
	}

	populateDefaultsForDatabase(obj)

	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = computeEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.UpdateDatabaseMetadata{}
	op, err := s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Database)
		return result, nil
	})
	if err != nil {
		return op, err
	}
	response, err := anypb.New(obj)
	if err != nil {
		return op, err
	}
	op.Result = &longrunningpb.Operation_Response{
		Response: response,
	}
	op.Done = true
	return op, err
}

func (s *firestoreAdminServer) DeleteDatabase(ctx context.Context, req *pb.DeleteDatabaseRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDatabaseName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	// Deletion is soft-delete, but with a different name!

	obj := &pb.Database{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// The name changes on deletion
	newName := *name
	newName.Database = fmt.Sprintf("%x", time.Now().UnixNano())
	obj.Name = newName.String()

	obj.DeleteTime = timestamppb.New(time.Now())
	obj.PreviousId = name.Database

	obj.FreeTier = PtrTo(false)

	obj.Etag = computeEtag(obj)
	if err := s.storage.Create(ctx, newName.String(), obj); err != nil {
		return nil, err
	}

	metadata := &pb.DeleteDatabaseMetadata{}
	op, err := s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return &pb.Database{}, nil
	})
	if err != nil {
		return op, err
	}
	response, err := anypb.New(obj)
	if err != nil {
		return op, err
	}
	op.Result = &longrunningpb.Operation_Response{
		Response: response,
	}
	return op, err
}

func populateDefaultsForDatabase(obj *pb.Database) {
	if obj.Uid == "" {
		obj.Uid = fmt.Sprintf("%x", time.Now().UnixNano())
	}
	if obj.ConcurrencyMode == pb.Database_CONCURRENCY_MODE_UNSPECIFIED {
		obj.ConcurrencyMode = pb.Database_PESSIMISTIC
	}

	if obj.AppEngineIntegrationMode == pb.Database_APP_ENGINE_INTEGRATION_MODE_UNSPECIFIED {
		obj.AppEngineIntegrationMode = pb.Database_DISABLED
	}

	if obj.DatabaseEdition == pb.Database_DATABASE_EDITION_UNSPECIFIED {
		obj.DatabaseEdition = pb.Database_STANDARD
	}

	if obj.FreeTier == nil {
		// The first database in each project is free-tier
		obj.FreeTier = PtrTo(true)
	}

	if obj.PointInTimeRecoveryEnablement == pb.Database_POINT_IN_TIME_RECOVERY_ENABLEMENT_UNSPECIFIED {
		obj.PointInTimeRecoveryEnablement = pb.Database_POINT_IN_TIME_RECOVERY_DISABLED
	}

	switch obj.PointInTimeRecoveryEnablement {
	case pb.Database_POINT_IN_TIME_RECOVERY_DISABLED:
		obj.VersionRetentionPeriod = durationpb.New(time.Hour)
	case pb.Database_POINT_IN_TIME_RECOVERY_ENABLED:
		obj.VersionRetentionPeriod = durationpb.New(7 * 24 * time.Hour)
	}

	if obj.Type == pb.Database_DATABASE_TYPE_UNSPECIFIED {
		obj.Type = pb.Database_FIRESTORE_NATIVE
	}

	// Seems to clear empty cmek config rather than storing an empty object
	if obj.CmekConfig != nil && reflect.DeepEqual(obj.CmekConfig, &pb.Database_CmekConfig{}) {
		obj.CmekConfig = nil
	}
}

type databaseName struct {
	Project  *projects.ProjectData
	Database string
}

func (n *databaseName) String() string {
	return "projects/" + n.Project.ID + "/databases/" + n.Database
}

// parseDatabaseName parses a string into a databaseName.
// The expected form is projects/<projectID>/locations/<region>/zones/<zone>/networks/<networkId>
func (s *MockService) parseDatabaseName(name string) (*databaseName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "databases" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &databaseName{
			Project:  project,
			Database: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
