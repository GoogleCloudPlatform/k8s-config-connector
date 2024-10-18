// Copyright 2023 Google LLC
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
// krm.apiVersion: pubsublite.cnrm.cloud.google.com/v1beta1
// krm.kind: PubSubLiteSubscription
// proto.service: google.cloud.pubsublite.v1.AdminService
// proto.resource: Subscription

package mockpubsublite

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/pubsublite/v1"
)

func (s *PubSubLiteV1) GetSubscription(ctx context.Context, req *pb.GetSubscriptionRequest) (*pb.Subscription, error) {
	name, err := s.parseSubscriptionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Subscription{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PubSubLiteV1) CreateSubscription(ctx context.Context, req *pb.CreateSubscriptionRequest) (*pb.Subscription, error) {
	reqName := req.Parent + "/subscriptions/" + req.SubscriptionId
	name, err := s.parseSubscriptionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Subscription).(*pb.Subscription)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PubSubLiteV1) UpdateSubscription(ctx context.Context, req *pb.UpdateSubscriptionRequest) (*pb.Subscription, error) {
	reqName := req.GetSubscription().GetName()

	name, err := s.parseSubscriptionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Subscription{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *PubSubLiteV1) DeleteSubscription(ctx context.Context, req *pb.DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	name, err := s.parseSubscriptionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Subscription{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type subscriptionName struct {
	Project      *projects.ProjectData
	Location     string
	Subscription string
}

func (s *subscriptionName) String() string {
	return "projects/" + s.Project.ID + "/locations/" + s.Location + "/subscriptions/" + s.Subscription
}

func (s *MockService) parseSubscriptionName(name string) (*subscriptionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "subscriptions" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &subscriptionName{
			Project:      project,
			Location:     tokens[3],
			Subscription: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
