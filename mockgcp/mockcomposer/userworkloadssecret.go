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
// proto.service: google.cloud.orchestration.airflow.service.v1.Environments
// proto.message: google.cloud.orchestration.airflow.service.v1.UserWorkloadsSecret

package mockcomposer

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type userWorkloadsSecretName struct {
	Project             *projects.ProjectData
	Location            string
	Environment         string
	UserWorkloadsSecret string
}

func (n *userWorkloadsSecretName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/environments/%s/userWorkloadsSecrets/%s", n.Project.ID, n.Location, n.Environment, n.UserWorkloadsSecret)
}

func (s *MockService) parseUserWorkloadsSecretName(name string) (*userWorkloadsSecretName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "environments" && tokens[6] == "userWorkloadsSecrets" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &userWorkloadsSecretName{
			Project:             project,
			Location:            tokens[3],
			Environment:         tokens[5],
			UserWorkloadsSecret: tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ComposerV1) GetUserWorkloadsSecret(ctx context.Context, req *pb.GetUserWorkloadsSecretRequest) (*pb.UserWorkloadsSecret, error) {
	name, err := s.parseUserWorkloadsSecretName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UserWorkloadsSecret{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "secrets %q not found", name.UserWorkloadsSecret)
		}
		return nil, err
	}

	res := proto.Clone(obj).(*pb.UserWorkloadsSecret)
	s.maskSecretData(res)
	return res, nil
}

func (s *ComposerV1) CreateUserWorkloadsSecret(ctx context.Context, req *pb.CreateUserWorkloadsSecretRequest) (*pb.UserWorkloadsSecret, error) {
	secretName := req.GetUserWorkloadsSecret().GetName()
	if secretName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "secret name is required")
	}
	if !strings.HasPrefix(secretName, "projects/") {
		secretName = fmt.Sprintf("%s/userWorkloadsSecrets/%s", req.GetParent(), secretName)
	}

	name, err := s.parseUserWorkloadsSecretName(secretName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetUserWorkloadsSecret()).(*pb.UserWorkloadsSecret)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	res := proto.Clone(obj).(*pb.UserWorkloadsSecret)
	s.maskSecretData(res)
	return res, nil
}

func (s *ComposerV1) UpdateUserWorkloadsSecret(ctx context.Context, req *pb.UpdateUserWorkloadsSecretRequest) (*pb.UserWorkloadsSecret, error) {
	secretName := req.GetUserWorkloadsSecret().GetName()
	name, err := s.parseUserWorkloadsSecretName(secretName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UserWorkloadsSecret{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "secrets %q not found", name.UserWorkloadsSecret)
		}
		return nil, err
	}

	obj.Data = req.GetUserWorkloadsSecret().GetData()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	res := proto.Clone(obj).(*pb.UserWorkloadsSecret)
	s.maskSecretData(res)
	return res, nil
}

func (s *ComposerV1) DeleteUserWorkloadsSecret(ctx context.Context, req *pb.DeleteUserWorkloadsSecretRequest) (*emptypb.Empty, error) {
	name, err := s.parseUserWorkloadsSecretName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.UserWorkloadsSecret{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "secrets %q not found", name.UserWorkloadsSecret)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ComposerV1) ListUserWorkloadsSecrets(ctx context.Context, req *pb.ListUserWorkloadsSecretsRequest) (*pb.ListUserWorkloadsSecretsResponse, error) {
	findPrefix := req.GetParent() + "/userWorkloadsSecrets/"

	response := &pb.ListUserWorkloadsSecretsResponse{}
	findKind := (&pb.UserWorkloadsSecret{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		secret := proto.Clone(obj).(*pb.UserWorkloadsSecret)
		s.maskSecretData(secret)
		response.UserWorkloadsSecrets = append(response.UserWorkloadsSecrets, secret)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ComposerV1) maskSecretData(obj *pb.UserWorkloadsSecret) {
	if obj.Data != nil {
		for k := range obj.Data {
			obj.Data[k] = "**********"
		}
	}
}
