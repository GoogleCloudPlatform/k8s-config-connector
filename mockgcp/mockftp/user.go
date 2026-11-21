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

package mockftp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/ftp/apiv1/ftppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func toBasicUserView(user *pb.User) *pb.User {
	return &pb.User{
		Name:                   user.Name,
		Username:               user.Username,
		CustomerServiceAccount: user.CustomerServiceAccount,
		State:                  user.State,
		CreateTime:             user.CreateTime,
		UpdateTime:             user.UpdateTime,
	}
}

func (s *MockService) populateUserDefaults(name *userName, obj *pb.User) {
	if obj.State == pb.User_STATE_UNSPECIFIED {
		obj.State = pb.User_ACTIVE
	}
	if obj.Username == "" && name != nil {
		obj.Username = name.User
	}
}

func (s *CloudFtpServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	name, err := s.parseUserName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.User{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	s.populateUserDefaults(name, obj)

	if req.View == pb.UserView_USER_VIEW_FULL {
		return obj, nil
	}

	return toBasicUserView(obj), nil
}

func (s *CloudFtpServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*longrunning.Operation, error) {
	serverName, err := s.parseServerName(req.Parent)
	if err != nil {
		return nil, err
	}

	if req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "userId is required")
	}

	// Verify parent server exists
	server := &pb.Server{}
	if err := s.storage.Get(ctx, serverName.String(), server); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", serverName.String())
		}
		return nil, err
	}

	fqn := fmt.Sprintf("%s/users/%s", serverName.String(), req.UserId)
	name, err := s.parseUserName(fqn)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	obj := proto.Clone(req.User).(*pb.User)
	obj.Name = fqn
	obj.Username = req.UserId
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	s.populateUserDefaults(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	prefix := serverName.LocationString()
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return toBasicUserView(obj), nil
	})
}

func (s *CloudFtpServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*longrunning.Operation, error) {
	if req.GetUser() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "user is required")
	}
	name, err := s.parseUserName(req.GetUser().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.User{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	updated := proto.Clone(existing).(*pb.User)
	updated.UpdateTime = timestamppb.New(now)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated.Labels = req.GetUser().GetLabels()
		updated.StorageDirectoryMappings = req.GetUser().GetStorageDirectoryMappings()
		updated.UserCredentials = req.GetUser().GetUserCredentials()
		if req.GetUser().GetCustomerServiceAccount() != "" {
			updated.CustomerServiceAccount = req.GetUser().GetCustomerServiceAccount()
		}
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				updated.Labels = req.GetUser().GetLabels()
			case "storage_directory_mappings", "storageDirectoryMappings":
				updated.StorageDirectoryMappings = req.GetUser().GetStorageDirectoryMappings()
			case "user_credentials", "userCredentials":
				updated.UserCredentials = req.GetUser().GetUserCredentials()
			case "customer_service_account", "customerServiceAccount":
				updated.CustomerServiceAccount = req.GetUser().GetCustomerServiceAccount()
			}
		}
	}

	s.populateUserDefaults(name, updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	prefix := name.LocationString()
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return toBasicUserView(updated), nil
	})
}

func (s *CloudFtpServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*longrunning.Operation, error) {
	name, err := s.parseUserName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.User{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	prefix := name.LocationString()
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *CloudFtpServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	serverName, err := s.parseServerName(req.Parent)
	if err != nil {
		return nil, err
	}

	response := &pb.ListUsersResponse{}
	userPrefix := serverName.String() + "/users/"

	userKind := (&pb.User{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, userKind, storage.ListOptions{
		Prefix: userPrefix,
	}, func(obj proto.Message) error {
		user := obj.(*pb.User)
		if strings.HasPrefix(user.Name, userPrefix) {
			uName, _ := s.parseUserName(user.Name)
			s.populateUserDefaults(uName, user)
			if req.View == pb.UserView_USER_VIEW_FULL {
				response.Users = append(response.Users, user)
			} else {
				response.Users = append(response.Users, toBasicUserView(user))
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
