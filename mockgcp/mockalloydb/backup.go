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

package mockalloydb

import (
	"context"
	"strings"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
)

func (s *AlloyDBAdminV1) GetBackup(ctx context.Context, req *pb.GetBackupRequest) (*pb.Backup, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func setBackupFields(name *backupName, obj *pb.Backup) {
	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.State = pb.Backup_READY
	obj.Uid = "11111111-1111-1111-1111-111111111111"
	obj.Reconciling = false
	if obj.EncryptionConfig != nil && obj.EncryptionConfig.KmsKeyName != "" {
		obj.EncryptionInfo = &pb.EncryptionInfo{
			EncryptionType: pb.EncryptionInfo_CUSTOMER_MANAGED_ENCRYPTION,
		}
	} else {
		obj.EncryptionInfo = &pb.EncryptionInfo{
			EncryptionType: pb.EncryptionInfo_GOOGLE_DEFAULT_ENCRYPTION,
		}
	}
}

func (s *AlloyDBAdminV1) CreateBackup(ctx context.Context, req *pb.CreateBackupRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/backups/" + req.BackupId
	name, err := s.parseBackupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.Backup)
	obj.Name = fqn
	setBackupFields(name, obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "create")
	return s.operations.StartLRO(ctx, req.Parent, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *AlloyDBAdminV1) UpdateBackup(ctx context.Context, req *pb.UpdateBackupRequest) (*longrunning.Operation, error) {
	reqName := req.GetBackup().GetName()
	name, err := s.parseBackupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		tokens := strings.Split(path, ".")
		switch tokens[0] {
		case "labels":
			obj.Labels = req.Backup.GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	obj.UpdateTime = timestamppb.Now()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "update")
	return s.operations.StartLRO(ctx, name.ProjectAndLocation(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *AlloyDBAdminV1) DeleteBackup(ctx context.Context, req *pb.DeleteBackupRequest) (*longrunning.Operation, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Backup{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "delete")
	return s.operations.StartLRO(ctx, name.ProjectAndLocation(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := &emptypb.Empty{}
		return result, nil
	})
}
