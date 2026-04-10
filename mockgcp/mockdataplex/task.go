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
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Task

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/google/uuid"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *DataplexService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	name, err := s.parseTaskName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Task{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataplexService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/tasks/%s", req.Parent, req.TaskId)
	name, err := s.parseTaskName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := timestamppb.New(time.Now())
	uid := uuid.NewString()
	obj := proto.Clone(req.Task).(*pb.Task)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Uid = uid
	obj.State = pb.State_ACTIVE
	obj.ExecutionStatus = &pb.Task_ExecutionStatus{
		LatestJob:  &pb.Job{Name: fmt.Sprintf("projects/%d/locations/%s/lakes/%s/tasks/%s/jobs/%s", name.Project.Number, name.Location, name.LakeID, name.TaskID, uid), Uid: uid, Trigger: pb.Job_TASK_CONFIG},
		UpdateTime: now,
	}

	// Ensure nested required fields have defaults if not provided, matching potential real API behavior
	if obj.TriggerSpec == nil {
		obj.TriggerSpec = &pb.Task_TriggerSpec{} // Might need more defaults depending on TriggerSpec.Type
	}
	if obj.ExecutionSpec == nil {
		// ExecutionSpec.ServiceAccount is required, but we let validation handle it for now
		obj.ExecutionSpec = &pb.Task_ExecutionSpec{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, name.lroPrefix(), lroMetadata, func() (proto.Message, error) {
		// Simulate task becoming fully ready/active after creation
		obj.State = pb.State_ACTIVE
		obj.ExecutionStatus = &pb.Task_ExecutionStatus{}
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *DataplexService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*longrunning.Operation, error) {
	name, err := s.parseTaskName(req.GetTask().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Task{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	if updateMask == nil {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required for update")
	}

	now := timestamppb.New(time.Now())
	obj.UpdateTime = now

	// Apply updates based on the update mask
	for _, path := range updateMask.GetPaths() {
		switch path {
		case "description":
			obj.Description = req.GetTask().GetDescription()
		case "display_name":
			obj.DisplayName = req.GetTask().GetDisplayName()
		case "labels":
			obj.Labels = req.GetTask().GetLabels()
		case "trigger_spec":
			// Be careful: trigger_spec.type is immutable
			if obj.TriggerSpec.Type != req.GetTask().GetTriggerSpec().GetType() && req.GetTask().GetTriggerSpec().GetType() != pb.Task_TriggerSpec_TYPE_UNSPECIFIED {
				return nil, status.Errorf(codes.InvalidArgument, "field `trigger_spec.type` is immutable")
			}
			obj.TriggerSpec = req.GetTask().GetTriggerSpec()
		case "execution_spec":
			obj.ExecutionSpec = req.GetTask().GetExecutionSpec()
		case "spark":
			obj.Config = &pb.Task_Spark{Spark: req.GetTask().GetSpark()}
		case "notebook":
			obj.Config = &pb.Task_Notebook{Notebook: req.GetTask().GetNotebook()}
		// Add other updatable fields here based on the proto definition
		default:
			// Handle nested fields like execution_spec.args, trigger_spec.schedule etc. if needed.
			// For simplicity, we are only handling top-level fields or full replacement of message fields.
			// A more sophisticated mock might handle field paths like "execution_spec.args".
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q is not supported by the mock", path)
		}
	}

	// Re-apply potentially modified nested structures
	if obj.TriggerSpec == nil {
		obj.TriggerSpec = &pb.Task_TriggerSpec{}
	}
	if obj.ExecutionSpec == nil {
		obj.ExecutionSpec = &pb.Task_ExecutionSpec{}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "update",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, name.lroPrefix(), lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *DataplexService) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	// TODO: Handle parent parsing more robustly if needed (e.g., validating lake existence)
	parentName, err := s.parseLakeName(req.Parent)
	if err != nil {
		// List operations usually return InvalidArgument for bad parent format
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent format: %v", err)
	}
	parentFqnPrefix := parentName.String() + "/tasks/"

	response := &pb.ListTasksResponse{}

	taskKind := (&pb.Task{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, taskKind, storage.ListOptions{}, func(obj proto.Message) error {
		task := obj.(*pb.Task)
		// Ensure the task belongs to the requested parent lake
		if strings.HasPrefix(task.GetName(), parentFqnPrefix) {
			response.Tasks = append(response.Tasks, task)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// TODO: Implement filtering and pagination if necessary

	return response, nil
}

func (s *DataplexService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*longrunning.Operation, error) {
	name, err := s.parseTaskName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Task{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "delete",
		CreateTime: timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, name.lroPrefix(), lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type taskName struct {
	Project  *projects.ProjectData
	Location string
	LakeID   string
	TaskID   string
}

func (n *taskName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", n.Project.ID, n.Location, n.LakeID, n.TaskID)
}

func (n *taskName) lroPrefix() string {
	// LROs for tasks are typically scoped to the location
	return fmt.Sprintf("projects/%s/locations/%s", n.Project.ID, n.Location)
}

// parseTaskName parses a string into a taskName.
// The expected form is `projects/*/locations/*/lakes/*/tasks/*`.
func (s *MockService) parseTaskName(name string) (*taskName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" && tokens[6] == "tasks" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "project %q not found", tokens[1])
		}

		n := &taskName{
			Project:  project,
			Location: tokens[3],
			LakeID:   tokens[5],
			TaskID:   tokens[7],
		}
		return n, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid task resource name", name)
}
