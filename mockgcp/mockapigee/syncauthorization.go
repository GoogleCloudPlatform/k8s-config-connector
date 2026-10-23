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

package mockapigee

import (
	"context"
	"crypto/md5"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
)

func (s *organizationsServer) GetSyncAuthorizationOrganization(ctx context.Context, req *pb.GetSyncAuthorizationOrganizationRequest) (*pb.GoogleCloudApigeeV1SyncAuthorization, error) {
	name, err := s.parseOrganizationName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String() + "/syncAuthorization"
	obj := &pb.GoogleCloudApigeeV1SyncAuthorization{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// If not set yet, return empty identities with computed etag
			obj.Identities = nil
			obj.Etag = computeEtag(obj)
			return obj, nil
		}
		return nil, err
	}
	return obj, nil
}

func (s *organizationsServer) SetSyncAuthorizationOrganization(ctx context.Context, req *pb.SetSyncAuthorizationOrganizationRequest) (*pb.GoogleCloudApigeeV1SyncAuthorization, error) {
	name, err := s.parseOrganizationName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String() + "/syncAuthorization"

	syncAuth := req.GetOrganization()
	if syncAuth == nil {
		return nil, status.Errorf(codes.InvalidArgument, "organization is required")
	}

	obj := proto.CloneOf(syncAuth)
	obj.Etag = computeEtag(obj)

	existing := &pb.GoogleCloudApigeeV1SyncAuthorization{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			if err := s.storage.Create(ctx, fqn, obj); err != nil {
				return nil, err
			}
			return obj, nil
		}
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func computeEtag(obj proto.Message) []byte {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return hash[:]
}
