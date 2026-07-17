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
	"time"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NetworkSecurityV1Server struct {
	*MockService
	pbv1.UnimplementedNetworkSecurityServer
}

func (s *NetworkSecurityV1Server) CreateTlsInspectionPolicy(ctx context.Context, req *pbv1.CreateTlsInspectionPolicyRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/tlsInspectionPolicies/" + req.TlsInspectionPolicyId

	fqn := name

	obj := proto.Clone(req.TlsInspectionPolicy).(*pbv1.TlsInspectionPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

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

func (s *NetworkSecurityV1Server) GetTlsInspectionPolicy(ctx context.Context, req *pbv1.GetTlsInspectionPolicyRequest) (*pbv1.TlsInspectionPolicy, error) {
	fqn := req.Name

	obj := &pbv1.TlsInspectionPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "TlsInspectionPolicy %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkSecurityV1Server) UpdateTlsInspectionPolicy(ctx context.Context, req *pbv1.UpdateTlsInspectionPolicyRequest) (*longrunning.Operation, error) {
	fqn := req.TlsInspectionPolicy.GetName()

	actual := &pbv1.TlsInspectionPolicy{}
	if err := s.storage.Get(ctx, fqn, actual); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.TlsInspectionPolicy).(*pbv1.TlsInspectionPolicy)
	updated.CreateTime = actual.CreateTime
	updated.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pbv1.OperationMetadata{
		CreateTime: updated.CreateTime,
		Target:     fqn,
		Verb:       "update",
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return updated, nil
	})
}

func (s *NetworkSecurityV1Server) DeleteTlsInspectionPolicy(ctx context.Context, req *pbv1.DeleteTlsInspectionPolicyRequest) (*longrunning.Operation, error) {
	fqn := req.Name

	actual := &pbv1.TlsInspectionPolicy{}
	if err := s.storage.Get(ctx, fqn, actual); err != nil {
		return nil, err
	}

	deleted := &pbv1.TlsInspectionPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	metadata := &pbv1.OperationMetadata{
		CreateTime: actual.CreateTime,
		Target:     fqn,
		Verb:       "delete",
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return deleted, nil
	})
}
