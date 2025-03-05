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
// proto.service: google.cloud.batch.v1.BatchService
// proto.message: google.cloud.batch.v1.Task

package mockbatch

import (
	"context"
	"fmt"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/batch/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/protobuf/proto"
)

func (s *BatchV1) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	name, err := s.parseTaskName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Task{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Task %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BatchV1) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	name, err := s.parseTaskName(req.Parent)
	if err != nil {
		return nil, err
	}

	fqn := name.String() + "/tasks"
	response := &pb.ListTasksResponse{}

	if err := s.storage.List(ctx, (&pb.Task{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: fqn}, func(obj proto.Message) error {
		task, ok := obj.(*pb.Task)
		if !ok {
			return status.Errorf(codes.Internal, "unexpected resource type: %T", obj)
		}
		response.Tasks = append(response.Tasks, task)
		return nil
	}); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Tasks in TaskGroup %s not found.", fqn)
		}
		return nil, err
	}

	return response, nil
}

type taskName struct {
	Project    string
	Location   string
	JobName    string
	TaskGroup  string
	TaskNumber string
}

func (n *taskName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/jobs/%s/taskGroups/%s/tasks/%s", n.Project, n.Location, n.JobName, n.TaskGroup, n.TaskNumber)
}

// parseTaskName parses a string into a taskName.
// The expected form is `projects/*/locations/*/jobs/*/taskGroups/*/tasks/*`.
func (s *BatchV1) parseTaskName(name string) (*taskName, error) {
	r := regexp.MustCompile(`^projects/([^/]+)/locations/([^/]+)/jobs/([^/]+)/taskGroups/([^/]+)/tasks/([^/]+)$`)
	tokens := r.FindStringSubmatch(name)
	if tokens == nil {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
	return &taskName{
		Project:    tokens[1],
		Location:   tokens[2],
		JobName:    tokens[3],
		TaskGroup:  tokens[4],
		TaskNumber: tokens[5],
	}, nil
}
