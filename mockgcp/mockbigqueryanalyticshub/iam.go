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

package mockbigqueryanalyticshub

import (
	"context"

	"cloud.google.com/go/iam/apiv1/iampb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *analyticsHubServer) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	policy := &iampb.Policy{}
	if err := s.storage.Get(ctx, req.Resource, policy); err != nil {
		if status.Code(err) == codes.NotFound {
			return &iampb.Policy{}, nil
		}
		return nil, err
	}
	return policy, nil
}

func (s *analyticsHubServer) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	policy := req.Policy
	if err := s.storage.Create(ctx, req.Resource, policy); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			if err := s.storage.Update(ctx, req.Resource, policy); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return policy, nil
}

func (s *analyticsHubServer) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return &iampb.TestIamPermissionsResponse{
		Permissions: req.Permissions,
	}, nil
}
