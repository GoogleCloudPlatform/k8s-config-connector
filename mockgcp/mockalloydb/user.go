// Copyright 2025 Google LLC
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
	"sort"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/alloydb/v1beta"
)

func (s *AlloyDBAdminV1) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	name, err := s.parseUserName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.User{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "user %s does not exist", name.UserName)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AlloyDBAdminV1) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	reqName := req.Parent + "/users/" + req.GetUserId()
	name, err := s.parseUserName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.User).(*pb.User)
	obj.Name = fqn

	if obj.DatabaseRoles == nil && obj.UserType == pb.User_ALLOYDB_IAM_USER {
		obj.DatabaseRoles = []string{"alloydbiamuser"}
	} else if obj.DatabaseRoles != nil {
		sort.Strings(obj.DatabaseRoles)
	}

	if obj.Password != "" {
		obj.Password = ""
	}

	// AlloyDB user does POST for update in the TF library.
	oldObj := &pb.User{}
	err = s.storage.Get(ctx, fqn, oldObj)
	if err == nil {
		s.storage.Delete(ctx, fqn, oldObj)
	} else if status.Code(err) != codes.NotFound {
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AlloyDBAdminV1) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	reqName := req.GetUser().GetName()

	name, err := s.parseUserName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.User{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		topLevelField := strings.Split(path, ".")[0]
		switch topLevelField {
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AlloyDBAdminV1) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {
	name, err := s.parseUserName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.User{}
	if err := s.storage.Get(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
