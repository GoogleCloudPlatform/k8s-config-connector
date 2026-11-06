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

package mockmemorystore

import (
	"context"

	v1pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	betapb "cloud.google.com/go/memorystore/apiv1beta/memorystorepb"
	"google.golang.org/genproto/googleapis/longrunning"
)

type instanceServerV1 struct {
	v1pb.UnimplementedMemorystoreServer
	beta *instanceServer
}

func (s *instanceServerV1) GetInstance(ctx context.Context, req *v1pb.GetInstanceRequest) (*v1pb.Instance, error) {
	betaReq := &betapb.GetInstanceRequest{Name: req.GetName()}
	betaObj, err := s.beta.GetInstance(ctx, betaReq)
	if err != nil {
		return nil, err
	}
	v1Obj := &v1pb.Instance{}
	if err := roundTripProto(betaObj, v1Obj); err != nil {
		return nil, err
	}
	return v1Obj, nil
}

func (s *instanceServerV1) CreateInstance(ctx context.Context, req *v1pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	betaReq := &betapb.CreateInstanceRequest{}
	if err := roundTripProto(req, betaReq); err != nil {
		return nil, err
	}
	return s.beta.createInstance(ctx, betaReq, true)
}

func (s *instanceServerV1) UpdateInstance(ctx context.Context, req *v1pb.UpdateInstanceRequest) (*longrunning.Operation, error) {
	betaReq := &betapb.UpdateInstanceRequest{}
	if err := roundTripProto(req, betaReq); err != nil {
		return nil, err
	}
	return s.beta.updateInstance(ctx, betaReq, true)
}

func (s *instanceServerV1) DeleteInstance(ctx context.Context, req *v1pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	betaReq := &betapb.DeleteInstanceRequest{Name: req.GetName()}
	return s.beta.deleteInstance(ctx, betaReq, true)
}

func (s *instanceServerV1) GetBackup(ctx context.Context, req *v1pb.GetBackupRequest) (*v1pb.Backup, error) {
	betaReq := &betapb.GetBackupRequest{Name: req.GetName()}
	betaObj, err := s.beta.GetBackup(ctx, betaReq)
	if err != nil {
		return nil, err
	}
	v1Obj := &v1pb.Backup{}
	if err := roundTripProto(betaObj, v1Obj); err != nil {
		return nil, err
	}
	return v1Obj, nil
}

func (s *instanceServerV1) BackupInstance(ctx context.Context, req *v1pb.BackupInstanceRequest) (*longrunning.Operation, error) {
	betaReq := &betapb.BackupInstanceRequest{}
	if err := roundTripProto(req, betaReq); err != nil {
		return nil, err
	}
	return s.beta.backupInstance(ctx, betaReq, true)
}

func (s *instanceServerV1) DeleteBackup(ctx context.Context, req *v1pb.DeleteBackupRequest) (*longrunning.Operation, error) {
	betaReq := &betapb.DeleteBackupRequest{Name: req.GetName()}
	return s.beta.deleteBackup(ctx, betaReq, true)
}
