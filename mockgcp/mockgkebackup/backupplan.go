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
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.BackupPlan

package mockgkebackup

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkebackup/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *backupForGKEServer) GetBackupPlan(ctx context.Context, req *pb.GetBackupPlanRequest) (*pb.BackupPlan, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackupPlan %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *backupForGKEServer) CreateBackupPlan(ctx context.Context, req *pb.CreateBackupPlanRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupPlans/%s", req.GetParent(), req.GetBackupPlanId())
	name, err := s.parseBackupPlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetBackupPlan()).(*pb.BackupPlan)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.BackupPlanID

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupForGKEServer) UpdateBackupPlan(ctx context.Context, req *pb.UpdateBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.GetBackupPlan().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetBackupPlan().GetDescription()
		case "labels":
			obj.Labels = req.GetBackupPlan().GetLabels()
		case "backup_config":
			obj.BackupConfig = req.GetBackupPlan().GetBackupConfig()
		case "backup_schedule":
			obj.BackupSchedule = req.GetBackupPlan().GetBackupSchedule()
		case "deactivated":
			obj.Deactivated = req.GetBackupPlan().GetDeactivated()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupForGKEServer) DeleteBackupPlan(ctx context.Context, req *pb.DeleteBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.BackupPlan{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type backupPlanName struct {
	Project      *projects.ProjectData
	Location     string
	BackupPlanID string
}

func (n *backupPlanName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", n.Project.ID, n.Location, n.BackupPlanID)
}

// parseBackupPlanName parses a string into an backupPlanName.
// The expected form is `projects/*/locations/*/backupPlans/*`.
func (s *MockService) parseBackupPlanName(name string) (*backupPlanName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPlans" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupPlanName{
			Project:      project,
			Location:     tokens[3],
			BackupPlanID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

</out>
</example>

```go
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
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.BackupPlan

package mockgkebackup

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkebackup/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *backupForGKEServer) GetBackupPlan(ctx context.Context, req *pb.GetBackupPlanRequest) (*pb.BackupPlan, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackupPlan %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *backupForGKEServer) CreateBackupPlan(ctx context.Context, req *pb.CreateBackupPlanRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupPlans/%s", req.GetParent(), req.GetBackupPlanId())
	name, err := s.parseBackupPlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetBackupPlan()).(*pb.BackupPlan)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.BackupPlanID // TODO: maybe a proper random value?
	obj.State = pb.BackupPlan_READY

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupForGKEServer) UpdateBackupPlan(ctx context.Context, req *pb.UpdateBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.GetBackupPlan().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetBackupPlan().GetDescription()
		case "labels":
			obj.Labels = req.GetBackupPlan().GetLabels()
		case "backup_config":
			obj.BackupConfig = req.GetBackupPlan().GetBackupConfig()
		case "backup_schedule":
			obj.BackupSchedule = req.GetBackupPlan().GetBackupSchedule()
		case "deactivated":
			obj.Deactivated = req.GetBackupPlan().GetDeactivated()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupForGKEServer) DeleteBackupPlan(ctx context.Context, req *pb.DeleteBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.BackupPlan{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type backupPlanName struct {
	Project      *projects.ProjectData
	Location     string
	BackupPlanID string
}

func (n *backupPlanName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", n.Project.ID, n.Location, n.BackupPlanID)
}

// parseBackupPlanName parses a string into an backupPlanName.
// The expected form is `projects/*/locations/*/backupPlans/*`.
func (s *MockService) parseBackupPlanName(name string) (*backupPlanName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPlans" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupPlanName{
			Project:      project,
			Location:     tokens[3],
			BackupPlanID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

</out>
</example>

```go
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
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.BackupPlan

package mockgkebackup

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkebackup/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *backupForGKEServer) GetBackupPlan(ctx context.Context, req *pb.GetBackupPlanRequest) (*pb.BackupPlan, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackupPlan %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *backupForGKEServer) CreateBackupPlan(ctx context.Context, req *pb.CreateBackupPlanRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupPlans/%s", req.GetParent(), req.GetBackupPlanId())
	name, err := s.parseBackupPlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetBackupPlan()).(*pb.BackupPlan)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.BackupPlanID // TODO: maybe a proper random value?
	obj.State = pb.BackupPlan_READY

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupForGKEServer) UpdateBackupPlan(ctx context.Context, req *pb.UpdateBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.GetBackupPlan().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetBackupPlan().GetDescription()
		case "labels":
			obj.Labels = req.GetBackupPlan().GetLabels()
		case "backup_config":
			obj.BackupConfig = req.GetBackupPlan().GetBackupConfig()
		case "backup_schedule":
			obj.BackupSchedule = req.GetBackupPlan().GetBackupSchedule()
		case "deactivated":
			obj.Deactivated = req.GetBackupPlan().GetDeactivated()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupForGKEServer) DeleteBackupPlan(ctx context.Context, req *pb.DeleteBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.BackupPlan{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type backupPlanName struct {
	Project      *projects.ProjectData
	Location     string
	BackupPlanID string
}

func (n *backupPlanName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", n.Project.ID, n.Location, n.BackupPlanID)
}

// parseBackupPlanName parses a string into an backupPlanName.
// The expected form is `projects/*/locations/*/backupPlans/*`.
func (s *MockService) parseBackupPlanName(name string) (*backupPlanName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPlans" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupPlanName{
			Project:      project,
			Location:     tokens[3],
			BackupPlanID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

