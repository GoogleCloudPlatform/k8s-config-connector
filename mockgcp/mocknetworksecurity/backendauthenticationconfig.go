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

package mocknetworksecurity

import (
	"context"
	"fmt"
	"time"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *NetworkSecurityV1Server) CreateBackendAuthenticationConfig(ctx context.Context, req *pbv1.CreateBackendAuthenticationConfigRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/backendAuthenticationConfigs/" + req.BackendAuthenticationConfigId

	fqn := name

	obj := proto.Clone(req.BackendAuthenticationConfig).(*pbv1.BackendAuthenticationConfig)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = fmt.Sprintf("%d", time.Now().UnixNano())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pbv1.OperationMetadata{
		CreateTime: obj.CreateTime,
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkSecurityV1Server) GetBackendAuthenticationConfig(ctx context.Context, req *pbv1.GetBackendAuthenticationConfigRequest) (*pbv1.BackendAuthenticationConfig, error) {
	fqn := req.Name

	obj := &pbv1.BackendAuthenticationConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackendAuthenticationConfig %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkSecurityV1Server) UpdateBackendAuthenticationConfig(ctx context.Context, req *pbv1.UpdateBackendAuthenticationConfigRequest) (*longrunning.Operation, error) {
	fqn := req.BackendAuthenticationConfig.Name

	existing := &pbv1.BackendAuthenticationConfig{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	obj := proto.Clone(req.BackendAuthenticationConfig).(*pbv1.BackendAuthenticationConfig)
	obj.Name = fqn
	obj.CreateTime = existing.CreateTime
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Etag = fmt.Sprintf("%d", time.Now().UnixNano())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pbv1.OperationMetadata{
		CreateTime: obj.UpdateTime,
		Target:     fqn,
		Verb:       "update",
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkSecurityV1Server) DeleteBackendAuthenticationConfig(ctx context.Context, req *pbv1.DeleteBackendAuthenticationConfigRequest) (*longrunning.Operation, error) {
	fqn := req.Name

	existing := &pbv1.BackendAuthenticationConfig{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	deleted := &pbv1.BackendAuthenticationConfig{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	metadata := &pbv1.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "delete",
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return deleted, nil
	})
}
