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

// +tool:mockgcp-support
// proto.service: google.cloud.dataplex.v1.ContentService

package mockdataplex

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
)

type ContentService struct {
	*MockService
	pb.UnimplementedContentServiceServer
}

func (s *ContentService) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	name, err := s.parseEntryGroupName(req.Resource)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	policy := &iampb.Policy{}
	if err := s.storage.Get(ctx, fqn, policy); err != nil {
		if status.Code(err) == codes.NotFound {
			policy = &iampb.Policy{
				Etag: []byte("ACAB"),
			}
			return policy, nil
		}
		return nil, err
	}

	return policy, nil
}

func (s *ContentService) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	name, err := s.parseEntryGroupName(req.Resource)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	policy := req.Policy
	policy.Etag = []byte(uuid.NewString())

	if err := s.storage.Update(ctx, fqn, policy); err != nil {
		if status.Code(err) == codes.NotFound {
			if err := s.storage.Create(ctx, fqn, policy); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return policy, nil
}

func (s *ContentService) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return &iampb.TestIamPermissionsResponse{
		Permissions: req.Permissions,
	}, nil
}
