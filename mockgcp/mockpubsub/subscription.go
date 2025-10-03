// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockpubsub

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/pubsub/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type subscriberService struct {
	*MockService
	pb.UnimplementedSubscriberServer
}

func (s *subscriberService) CreateSubscription(ctx context.Context, req *pb.Subscription) (*pb.Subscription, error) {
	name, err := s.parseSubscriptionName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := ProtoClone(req)
	obj.Name = name.String()

	obj.State = pb.Subscription_ACTIVE

	if obj.AckDeadlineSeconds == 0 {
		obj.AckDeadlineSeconds = 10
	}

	s.populateDefaultsForSubscription(obj)

	// Unlike many other APIs, creation is eventually consistent (not immediately visible in GET)
	go func() {
		ctx := context.Background()
		time.Sleep(2 * time.Second)
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			klog.Errorf("error creaing pubsub Subscription: %v", err)
		}
	}()

	ret := ProtoClone(obj)
	// If the original subscription contains PushConfig then add the version attribute.
	if req.PushConfig != nil {
		if ret.PushConfig.Attributes == nil {
			ret.PushConfig.Attributes = make(map[string]string)
		}
		if ret.PushConfig.Attributes["x-goog-version"] == "" {
			ret.PushConfig.Attributes["x-goog-version"] = "v1"
		}
	} else {
		ret.PushConfig = &pb.PushConfig{}
	}

	return ret, nil
}

func (s *subscriberService) populateDefaultsForSubscription(obj *pb.Subscription) {
	if obj.BigqueryConfig != nil {
		obj.BigqueryConfig.State = pb.BigQueryConfig_ACTIVE
	}
	if obj.CloudStorageConfig != nil {
		obj.CloudStorageConfig.State = pb.CloudStorageConfig_ACTIVE
	}

	if obj.ExpirationPolicy == nil {
		obj.ExpirationPolicy = &pb.ExpirationPolicy{
			Ttl: &durationpb.Duration{
				Seconds: 3600 * 24 * 31,
			},
		}
	}

	if obj.PushConfig == nil {
		obj.PushConfig = &pb.PushConfig{}
	}
}

func (s *subscriberService) UpdateSubscription(ctx context.Context, req *pb.UpdateSubscriptionRequest) (*pb.Subscription, error) {
	reqName := req.Subscription.Name
	name, err := s.parseSubscriptionName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.Subscription{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)
	updated.Name = name.String()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// Documented as required, but not passed by terraform...
		// return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
		updated.AckDeadlineSeconds = req.GetSubscription().AckDeadlineSeconds
		updated.EnableExactlyOnceDelivery = req.GetSubscription().EnableExactlyOnceDelivery
		updated.Labels = req.GetSubscription().Labels
		updated.PushConfig = req.GetSubscription().PushConfig
		updated.BigqueryConfig = req.GetSubscription().BigqueryConfig
		updated.CloudStorageConfig = req.GetSubscription().CloudStorageConfig
		if updated.CloudStorageConfig != nil {
			updated.CloudStorageConfig.OutputFormat = &pb.CloudStorageConfig_TextConfig_{}
			if updated.CloudStorageConfig.MaxDuration != nil && int64(updated.AckDeadlineSeconds) < updated.CloudStorageConfig.MaxDuration.Seconds {
				updated.AckDeadlineSeconds = int32(updated.CloudStorageConfig.MaxDuration.Seconds)
			}
		}
		updated.DeadLetterPolicy = req.GetSubscription().DeadLetterPolicy
		updated.ExpirationPolicy = req.GetSubscription().ExpirationPolicy
		updated.RetryPolicy = req.GetSubscription().RetryPolicy
		updated.RetainAckedMessages = req.GetSubscription().RetainAckedMessages
		updated.MessageRetentionDuration = req.GetSubscription().MessageRetentionDuration
	}
	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "ackDeadlineSeconds":
			updated.AckDeadlineSeconds = req.GetSubscription().AckDeadlineSeconds
		case "enableExactlyOnceDelivery":
			updated.EnableExactlyOnceDelivery = req.GetSubscription().EnableExactlyOnceDelivery
		case "labels":
			updated.Labels = req.GetSubscription().Labels
		case "pushConfig":
			updated.PushConfig = req.GetSubscription().PushConfig
		// case "expirationPolicy":
		// 	updated.ExpirationPolicy = req.GetSubscription().ExpirationPolicy
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock", path)
		}
	}

	s.populateDefaultsForSubscription(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	// Very unusual behaviour on the return value pushConfig ... maybe this is actually async also?
	ret := ProtoClone(updated)
	if ret.PushConfig == nil {
		ret.PushConfig = &pb.PushConfig{}
	}
	if ret.PushConfig.Attributes == nil {
		ret.PushConfig.Attributes = make(map[string]string)
	}
	if ret.PushConfig.Attributes["x-goog-version"] == "" {
		ret.PushConfig.Attributes["x-goog-version"] = "v1"
	}
	return ret, nil
}

func (s *subscriberService) GetSubscription(ctx context.Context, req *pb.GetSubscriptionRequest) (*pb.Subscription, error) {
	name, err := s.parseSubscriptionName(req.Subscription)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Subscription{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource not found (resource=%s).", name.ID)
		}
		return nil, err
	}
	// API doesn't return attributes even if it is set.
	if obj.PushConfig != nil && obj.PushConfig.Attributes != nil {
		obj.PushConfig.Attributes = nil
	}

	return obj, nil
}

func (s *subscriberService) ListSubscriptions(ctx context.Context, req *pb.ListSubscriptionsRequest) (*pb.ListSubscriptionsResponse, error) {
	project, err := s.Projects.GetProjectByID(req.Project)
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%v/", project.ID)

	var subscriptions []*pb.Subscription

	subscriptionKind := (&pb.Subscription{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, subscriptionKind, storage.ListOptions{}, func(obj proto.Message) error {
		subscription := obj.(*pb.Subscription)
		if strings.HasPrefix(subscription.Name, findPrefix) {
			subscriptions = append(subscriptions, subscription)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListSubscriptionsResponse{
		Subscriptions: subscriptions,
		NextPageToken: "",
	}, nil
}

func (s *subscriberService) DeleteSubscription(ctx context.Context, req *pb.DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	name, err := s.parseSubscriptionName(req.Subscription)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.Subscription{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type subscriptionName struct {
	Project *projects.ProjectData
	ID      string
}

func (n *subscriptionName) String() string {
	return fmt.Sprintf("projects/%s/subscriptions/%s", n.Project.ID, n.ID)
}

// parseSubscriptionName parses a string into a subscriptionName.
// The expected form is `projects/*/subscriptions/*`.
func (s *MockService) parseSubscriptionName(name string) (*subscriptionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "subscriptions" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &subscriptionName{
			Project: project,
			ID:      tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
