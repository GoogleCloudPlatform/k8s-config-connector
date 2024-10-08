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

package mockfirestore

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/firestore/admin/v1"
)

type DatabaseService struct {
	*MockService
	pb.UnimplementedFirestoreAdminServer
}

func (s *DatabaseService) GetDatabase(ctx context.Context, req *pb.GetDatabaseRequest) (*pb.Database, error) {
	fqn := req.GetName()

	obj := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *DatabaseService) CreateDatabase(ctx context.Context, req *pb.CreateDatabaseRequest) (*longrunningpb.Operation, error) {
	fqn := req.GetDatabase().GetName()

	obj := proto.Clone(req.Database).(*pb.Database)
	populateDefaultsForDatabase(obj)
	t := timestamppb.New(time.Now())
	obj.CreateTime = t
	obj.UpdateTime = t
	obj.EarliestVersionTime = t
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreateDatabaseMetadata{}
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
	return op, err
}

func (s *DatabaseService) UpdateDatabase(ctx context.Context, req *pb.UpdateDatabaseRequest) (*longrunningpb.Operation, error) {
	fqn := req.GetDatabase().GetName()

	existing := &pb.Database{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.Database).(*pb.Database)
	populateDefaultsForDatabase(updated)
	updated.UpdateTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.UpdateDatabaseMetadata{}
	op, err := s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		result := proto.Clone(updated).(*pb.Database)
		return result, nil
	})
	if err != nil {
		return op, err
	}
	response, err := anypb.New(updated)
	if err != nil {
		return op, err
	}
	op.Result = &longrunningpb.Operation_Response{
		Response: response,
	}
	return op, err
}

func (s *DatabaseService) DeleteDatabase(ctx context.Context, req *pb.DeleteDatabaseRequest) (*longrunningpb.Operation, error) {
	fqn := req.GetName()

	deleted := &pb.Database{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	metadata := &pb.DeleteDatabaseMetadata{}
	op, err := s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return &pb.Database{}, nil
	})
	if err != nil {
		return op, err
	}
	response, err := anypb.New(deleted)
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
	if obj.PointInTimeRecoveryEnablement == pb.Database_POINT_IN_TIME_RECOVERY_ENABLEMENT_UNSPECIFIED {
		obj.PointInTimeRecoveryEnablement = pb.Database_POINT_IN_TIME_RECOVERY_DISABLED
	}
	if obj.PointInTimeRecoveryEnablement == pb.Database_POINT_IN_TIME_RECOVERY_DISABLED {
		obj.VersionRetentionPeriod = durationpb.New(time.Hour)
	} else if obj.PointInTimeRecoveryEnablement == pb.Database_POINT_IN_TIME_RECOVERY_ENABLED {
		obj.VersionRetentionPeriod = durationpb.New(7 * 24 * time.Hour)
	}
	obj.Etag = computeEtag(obj)
}

func computeEtag(obj proto.Message) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
