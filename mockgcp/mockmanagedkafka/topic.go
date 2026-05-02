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

// +tool:mockgcp-support
// proto.service: google.cloud.managedkafka.v1.ManagedKafka
// proto.message: google.cloud.managedkafka.v1.Topic

package mockmanagedkafka

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"cloud.google.com/go/iam/apiv1/iampb"
	pb "cloud.google.com/go/managedkafka/apiv1/managedkafkapb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func (s *managedKafka) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	// For now, we only support Topics
	if _, err := s.parseTopicName(req.GetResource()); err != nil {
		return nil, err
	}

	fqn := req.GetResource() + "/iamPolicy"

	policy := &iampb.Policy{}
	if err := s.storage.Get(ctx, fqn, policy); err != nil {
		if status.Code(err) == codes.NotFound {
			return &iampb.Policy{
				Version: 3,
				Etag:    []byte("default"),
			}, nil
		}
		return nil, err
	}

	return policy, nil
}

func (s *managedKafka) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	if _, err := s.parseTopicName(req.GetResource()); err != nil {
		return nil, err
	}

	fqn := req.GetResource() + "/iamPolicy"

	policy := req.GetPolicy()
	policy.Etag = []byte(fmt.Sprintf("etag-%d", time.Now().UnixNano()))

	oldPolicy := &iampb.Policy{}
	if err := s.storage.Get(ctx, fqn, oldPolicy); err != nil {
		if status.Code(err) == codes.NotFound {
			if err := s.storage.Create(ctx, fqn, policy); err != nil {
				return nil, err
			}
			return policy, nil
		}
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, policy); err != nil {
		return nil, err
	}

	return policy, nil
}

func (s *managedKafka) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return &iampb.TestIamPermissionsResponse{
		Permissions: req.Permissions,
	}, nil
}

func (s *managedKafka) GetTopic(ctx context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {
	name, err := s.parseTopicName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Topic{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "org.apache.kafka.common.errors.UnknownTopicOrPartitionException: This server does not host this topic-partition.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) CreateTopic(ctx context.Context, req *pb.CreateTopicRequest) (*pb.Topic, error) {
	reqName := fmt.Sprintf("%s/topics/%s", req.GetParent(), req.GetTopicId())
	name, err := s.parseTopicName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetTopic()).(*pb.Topic)
	obj.Name = fqn
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) UpdateTopic(ctx context.Context, req *pb.UpdateTopicRequest) (*pb.Topic, error) {
	name, err := s.parseTopicName(req.GetTopic().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Topic{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}
	// updateMask=configs%2CpartitionCount
	for _, path := range paths {
		switch path {
		case "configs":
			obj.Configs = req.GetTopic().GetConfigs()
		case "partitionCount":
			obj.PartitionCount = req.GetTopic().GetPartitionCount()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) DeleteTopic(ctx context.Context, req *pb.DeleteTopicRequest) (*emptypb.Empty, error) {
	name, err := s.parseTopicName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Topic{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type topicName struct {
	Project  *projects.ProjectData
	Location string
	Cluster  string
	Topic    string
}

func (n *topicName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/topics/%s", n.Project.ID, n.Location, n.Cluster, n.Topic)
}

func (s *MockService) parseTopicName(name string) (*topicName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" && tokens[6] == "topics" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &topicName{
			Project:  project,
			Location: tokens[3],
			Cluster:  tokens[5],
			Topic:    tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid topic name %q", name)
}
