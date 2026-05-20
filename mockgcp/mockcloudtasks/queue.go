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

// +tool:mockgcp-support
// proto.service: google.cloud.tasks.v2.CloudTasks
// proto.message: google.cloud.tasks.v2.Queue

package mockcloudtasks

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/tasks/v2"
)

type cloudTasks struct {
	*MockService
	pb.UnimplementedProjectsLocationsQueuesServerServer
	pb.UnimplementedProjectsLocationsQueuesTasksServerServer
}

func (s *cloudTasks) GetProjectsLocationsQueue(ctx context.Context, req *pb.GetProjectsLocationsQueueRequest) (*pb.Queue, error) {
	name, err := s.parseQueueName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Queue{}
	obj.State = PtrTo("RUNNING")
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "queue %q not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *cloudTasks) PatchProjectsLocationsQueue(ctx context.Context, req *pb.PatchProjectsLocationsQueueRequest) (*pb.Queue, error) {
	reqName := req.GetProjectsLocationsQueue().GetName()

	name, err := s.parseQueueName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Queue{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := strings.Split(req.GetUpdateMask(), ",")
	for _, path := range paths {
		switch path {
		case "rateLimits.maxDispatchesPerSecond":
			if obj.RateLimits == nil {
				obj.RateLimits = &pb.RateLimits{}
			}
			if req.GetProjectsLocationsQueue().GetRateLimits().GetMaxDispatchesPerSecond() == 0 {
				obj.RateLimits.MaxDispatchesPerSecond = PtrTo(500.0)
			} else {
				obj.RateLimits.MaxDispatchesPerSecond = PtrTo(req.GetProjectsLocationsQueue().GetRateLimits().GetMaxDispatchesPerSecond())
			}
		case "rateLimits.maxBurstSize":
			if obj.RateLimits == nil {
				obj.RateLimits = &pb.RateLimits{}
			}
			if req.GetProjectsLocationsQueue().GetRateLimits().GetMaxBurstSize() == 0 {
				obj.RateLimits.MaxBurstSize = PtrTo(int32(100))
			} else {
				obj.RateLimits.MaxBurstSize = PtrTo(req.GetProjectsLocationsQueue().GetRateLimits().GetMaxBurstSize())
			}
		case "rateLimits.maxConcurrentDispatches":
			if obj.RateLimits == nil {
				obj.RateLimits = &pb.RateLimits{}
			}
			if req.GetProjectsLocationsQueue().GetRateLimits().GetMaxConcurrentDispatches() == 0 {
				obj.RateLimits.MaxConcurrentDispatches = PtrTo(int32(1000))
			} else {
				obj.RateLimits.MaxConcurrentDispatches = PtrTo(req.GetProjectsLocationsQueue().GetRateLimits().GetMaxConcurrentDispatches())
			}
		case "retryConfig.maxAttempts":
			if obj.RetryConfig == nil {
				obj.RetryConfig = &pb.RetryConfig{}
			}
			if req.GetProjectsLocationsQueue().GetRetryConfig().GetMaxAttempts() == 0 {
				obj.RetryConfig.MaxAttempts = PtrTo(int32(100))
			} else {
				obj.RetryConfig.MaxAttempts = PtrTo(req.GetProjectsLocationsQueue().GetRetryConfig().GetMaxAttempts())
			}
		case "retryConfig.minBackoff":
			if obj.RetryConfig == nil {
				obj.RetryConfig = &pb.RetryConfig{}
			}
			if req.GetProjectsLocationsQueue().GetRetryConfig().GetMinBackoff() == nil {
				obj.RetryConfig.MinBackoff = durationpb.New(time.Second / 10)
			} else {
				obj.RetryConfig.MinBackoff = req.GetProjectsLocationsQueue().GetRetryConfig().GetMinBackoff()
			}
		case "retryConfig.maxBackoff":
			if obj.RetryConfig == nil {
				obj.RetryConfig = &pb.RetryConfig{}
			}
			if req.GetProjectsLocationsQueue().GetRetryConfig().GetMaxBackoff() == nil {
				obj.RetryConfig.MaxBackoff = durationpb.New(3600 * time.Second)
			} else {
				obj.RetryConfig.MaxBackoff = req.GetProjectsLocationsQueue().GetRetryConfig().GetMaxBackoff()
			}
		case "retryConfig.maxDoublings":
			if obj.RetryConfig == nil {
				obj.RetryConfig = &pb.RetryConfig{}
			}
			if req.GetProjectsLocationsQueue().GetRetryConfig().GetMaxDoublings() == 0 {
				obj.RetryConfig.MaxDoublings = PtrTo(int32(16))
			} else {
				obj.RetryConfig.MaxDoublings = PtrTo(req.GetProjectsLocationsQueue().GetRetryConfig().GetMaxDoublings())
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *cloudTasks) CreateProjectsLocationsQueue(ctx context.Context, req *pb.CreateProjectsLocationsQueueRequest) (*pb.Queue, error) {
	name, err := s.parseQueueName(req.GetProjectsLocationsQueue().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetProjectsLocationsQueue()).(*pb.Queue)
	obj.Name = PtrTo(fqn)
	obj.RateLimits = &pb.RateLimits{
		MaxBurstSize:            PtrTo(int32(100)),
		MaxConcurrentDispatches: PtrTo(int32(1000)),
		MaxDispatchesPerSecond:  PtrTo(500.0),
	}
	obj.RetryConfig = &pb.RetryConfig{
		MaxAttempts:  PtrTo(int32(100)),
		MaxBackoff:   durationpb.New(3600 * time.Second),
		MaxDoublings: PtrTo(int32(16)),
		MinBackoff:   durationpb.New(time.Second / 10),
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *cloudTasks) DeleteProjectsLocationsQueue(ctx context.Context, req *pb.DeleteProjectsLocationsQueueRequest) (*pb.Empty, error) {
	name, err := s.parseQueueName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Queue{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

type queueParent struct {
	Project  *projects.ProjectData
	Location string
}

func (n *queueParent) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

type queueName struct {
	Parent       queueParent
	ResourceName string
}

func (n *queueName) String() string {
	return n.Parent.String() + "/queues/" + n.ResourceName
}

// parseQueueName parses a string into a queueName.
// The expected form is `projects/*/locations/*/queues/*`.
func (s *MockService) parseQueueName(name string) (*queueName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "queues" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &queueName{
			Parent: queueParent{
				Project:  project,
				Location: tokens[3],
			},
			ResourceName: tokens[5],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

// parseQueueParent parses a string into a queueParent.
// The expected form is `projects/*/locations/*`.
func (s *MockService) parseQueueParent(name string) (*queueParent, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		parent := &queueParent{
			Project:  project,
			Location: tokens[3],
		}
		return parent, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", name)
}

func PtrTo[T any](t T) *T {
	return &t
}
