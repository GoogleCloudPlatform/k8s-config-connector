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
	"strings"
	"time"

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *NetworkSecurityV1Server) CreateUrlList(ctx context.Context, req *pbv1.CreateUrlListRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/urlLists/" + req.UrlListId

	fqn := name

	obj := proto.Clone(req.UrlList).(*pbv1.UrlList)
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
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, req.Parent, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkSecurityV1Server) GetUrlList(ctx context.Context, req *pbv1.GetUrlListRequest) (*pbv1.UrlList, error) {
	fqn := req.Name

	obj := &pbv1.UrlList{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "UrlList %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkSecurityV1Server) UpdateUrlList(ctx context.Context, req *pbv1.UpdateUrlListRequest) (*longrunning.Operation, error) {
	fqn := req.UrlList.Name

	existing := &pbv1.UrlList{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	obj := proto.Clone(existing).(*pbv1.UrlList)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj = proto.Clone(req.UrlList).(*pbv1.UrlList)
	} else {
		if err := fields.UpdateByFieldMask(obj, req.UrlList, paths); err != nil {
			return nil, err
		}
	}

	obj.Name = fqn
	obj.CreateTime = existing.CreateTime
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix, _, _ := strings.Cut(fqn, "/urlLists/")

	metadata := &pbv1.OperationMetadata{
		CreateTime: obj.UpdateTime,
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *NetworkSecurityV1Server) DeleteUrlList(ctx context.Context, req *pbv1.DeleteUrlListRequest) (*longrunning.Operation, error) {
	fqn := req.Name

	existing := &pbv1.UrlList{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	deleted := &pbv1.UrlList{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix, _, _ := strings.Cut(fqn, "/urlLists/")

	metadata := &pbv1.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, metadata, func() (proto.Message, error) {
		return deleted, nil
	})
}
