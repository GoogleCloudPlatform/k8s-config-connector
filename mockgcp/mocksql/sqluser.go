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

package mocksql

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type sqlUsersService struct {
	*MockService
	pb.UnimplementedSqlUsersServiceServer
}

func (s *sqlUsersService) Get(ctx context.Context, req *pb.SqlUsersGetRequest) (*pb.User, error) {
	name, err := s.buildUserName(req.GetProject(), req.GetInstance(), req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.User{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.Etag = fields.ComputeWeakEtag(obj)
	if obj.Password != "" {
		obj.Password = ""
	}
	return obj, nil
}

func (s *sqlUsersService) List(ctx context.Context, req *pb.SqlUsersListRequest) (*pb.UsersListResponse, error) {
	name, err := s.buildInstanceName(req.GetProject(), req.GetInstance())
	if err != nil {
		return nil, err
	}

	ret := &pb.UsersListResponse{}
	ret.Kind = "sql#usersList"

	userKind := (&pb.User{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, userKind, storage.ListOptions{}, func(obj proto.Message) error {
		user := obj.(*pb.User)
		if user.GetProject() == name.Project.ID && user.GetInstance() == name.InstanceName {
			ret.Items = append(ret.Items, user)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	for _, item := range ret.Items {
		if item.Password != "" {
			item.Password = ""
		}
		item.Etag = fields.ComputeWeakEtag(item)
	}

	sort.Slice(ret.Items, func(i, j int) bool {
		return ret.Items[i].GetName() < ret.Items[j].GetName()
	})

	return ret, nil
}

func (s *sqlUsersService) Insert(ctx context.Context, req *pb.SqlUsersInsertRequest) (*pb.Operation, error) {
	name, err := s.buildUserName(req.GetProject(), req.GetInstance(), req.GetBody().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetBody()).(*pb.User)
	obj.Name = name.UserName
	obj.Project = name.Project.ID
	obj.Instance = name.Instance
	obj.Kind = "sql#user"
	if obj.PasswordPolicy == nil {
		obj.PasswordPolicy = &pb.UserPasswordValidationPolicy{Status: &pb.PasswordStatus{}}
	} else {
		if obj.PasswordPolicy.PasswordExpirationDuration != nil {
			if obj.PasswordPolicy.Status == nil {
				obj.PasswordPolicy.Status = &pb.PasswordStatus{}
			}
			obj.PasswordPolicy.Status.PasswordExpirationTime = timestamppb.New(time.Now().Add(obj.PasswordPolicy.PasswordExpirationDuration.AsDuration()))
		}
	}

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_CREATE_USER,
		Status:        pb.Operation_DONE, // Operation returns LRO, but it is (always?) done
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlUsersService) Update(ctx context.Context, req *pb.SqlUsersUpdateRequest) (*pb.Operation, error) {
	name, err := s.buildUserName(req.GetProject(), req.GetInstance(), req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.User{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj = proto.Clone(req.GetBody()).(*pb.User)
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_UPDATE_USER,
		Status:        pb.Operation_DONE, // Operation returns LRO, but it is (always?) done
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlUsersService) Delete(ctx context.Context, req *pb.SqlUsersDeleteRequest) (*pb.Operation, error) {
	name, err := s.buildUserName(req.GetProject(), req.GetInstance(), req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.User{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_DELETE_USER,
		Status:        pb.Operation_DONE, // Operation returns LRO, but it is (always?) done
	}
	return s.operations.startLRO(ctx, op, deleted, func() (proto.Message, error) {
		return deleted, nil
	})
}

type UserName struct {
	Project  *projects.ProjectData
	Instance string
	UserName string
}

func (n *UserName) String() string {
	return "projects/" + n.Project.ID + "/instances/" + n.Instance + "/users/" + n.UserName
}

// parseSQLUserName parses a string into a UserName.
// The expected form is projects/<projectID>/users/<SQLUserName>
func (s *MockService) parseUserName(name string) (*UserName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "users" {
		return s.buildUserName(tokens[1], tokens[3], tokens[5])
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *MockService) buildUserName(projectID, instanceName, userName string) (*UserName, error) {
	project, err := s.projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &UserName{
		Project:  project,
		Instance: instanceName,
		UserName: userName,
	}, nil
}
