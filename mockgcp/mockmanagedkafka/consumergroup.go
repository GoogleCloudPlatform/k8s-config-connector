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
// proto.message: google.cloud.managedkafka.v1.Cluster

package mockmanagedkafka

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/managedkafka/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *managedKafka) GetConsumerGroup(ctx context.Context, req *pb.GetConsumerGroupRequest) (*pb.ConsumerGroup, error) {
	name, err := s.parseConsumerGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConsumerGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%v' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) UpdateConsumerGroup(ctx context.Context, req *pb.UpdateConsumerGroupRequest) (*pb.ConsumerGroup, error) {
	name, err := s.parseConsumerGroupName(req.GetConsumerGroup().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.ConsumerGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedKafka) DeleteConsumerGroup(ctx context.Context, req *pb.DeleteConsumerGroupRequest) (*emptypb.Empty, error) {
	name, err := s.parseConsumerGroupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ConsumerGroup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type ConsumerGroupName struct {
	Project       *projects.ProjectData
	Location      string
	Cluster       string
	ConsumerGroup string
}

func (n *ConsumerGroupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s/consumerGroups/%s", n.Project.ID, n.Location, n.Cluster, n.ConsumerGroup)
}

func (s *MockService) parseConsumerGroupName(name string) (*ConsumerGroupName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clusters" && tokens[6] == "consumerGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &ConsumerGroupName{
			Project:       project,
			Location:      tokens[3],
			Cluster:       tokens[5],
			ConsumerGroup: tokens[7],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid consumer group name %q", name)
}
