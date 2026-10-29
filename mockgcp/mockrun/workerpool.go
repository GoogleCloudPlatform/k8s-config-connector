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

package mockrun

import (
	"context"
	"fmt"
	"strings"

	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/run/v2"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type workerPools struct {
	*MockService
	pb.UnimplementedWorkerPoolsServer
}

func (s *workerPools) GetWorkerPool(ctx context.Context, req *pb.GetWorkerPoolRequest) (*pb.WorkerPool, error) {
	name, err := s.parseWorkerPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' of kind 'WORKER_POOL' in region '%s' in project '%s' does not exist.", name.WorkerPool, name.Location, name.Project.ID)
		}
		return nil, err
	}

	ret := proto.CloneOf(obj)
	// CustomAudiences is mutable-but-unreadable in real GCP Cloud Run v2 API
	ret.CustomAudiences = nil
	return ret, nil
}

func (s *workerPools) populateWorkerPoolDefaults(name *workerPoolName, obj *pb.WorkerPool, isUpdate bool, templateChanged bool) {
	if obj.LaunchStage == 0 {
		obj.LaunchStage = 4 // LaunchStage_GA
	}

	if obj.Scaling == nil {
		defaultCount := int32(1)
		obj.Scaling = &pb.WorkerPoolScaling{
			ManualInstanceCount: &defaultCount,
		}
	} else if obj.Scaling.ManualInstanceCount == nil {
		defaultCount := int32(1)
		obj.Scaling.ManualInstanceCount = &defaultCount
	}

	if len(obj.InstanceSplits) == 0 {
		obj.InstanceSplits = []*pb.InstanceSplit{
			{
				Type:    pb.InstanceSplitAllocationType_INSTANCE_SPLIT_ALLOCATION_TYPE_LATEST,
				Percent: 100,
			},
		}
	}

	obj.LastModifier = fmt.Sprintf("overseer-kcc-tester@%s.iam.gserviceaccount.com", name.Project.ID)
	obj.ObservedGeneration = obj.Generation

	revID := fmt.Sprintf("%s-00001-abc", name.WorkerPool)
	if isUpdate && templateChanged {
		revID = fmt.Sprintf("%s-%05d-abc", name.WorkerPool, obj.Generation)
	}
	revName := fmt.Sprintf("%s/revisions/%s", obj.Name, revID)

	if !isUpdate || obj.LatestCreatedRevision == "" {
		obj.LatestCreatedRevision = revName
		obj.LatestReadyRevision = revName
	} else if templateChanged {
		obj.LatestCreatedRevision = revName
		obj.LatestReadyRevision = revName
	}

	if len(obj.InstanceSplitStatuses) == 0 {
		obj.InstanceSplitStatuses = []*pb.InstanceSplitStatus{
			{
				Type:     pb.InstanceSplitAllocationType_INSTANCE_SPLIT_ALLOCATION_TYPE_LATEST,
				Percent:  100,
				Revision: revID,
			},
		}
	} else {
		for _, status := range obj.InstanceSplitStatuses {
			if status.Type == pb.InstanceSplitAllocationType_INSTANCE_SPLIT_ALLOCATION_TYPE_LATEST {
				status.Revision = revID
			}
		}
	}

	if obj.Template != nil {
		if obj.Template.ServiceAccount == "" {
			obj.Template.ServiceAccount = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)
		}
		for _, container := range obj.Template.Containers {
			if container.Resources == nil {
				container.Resources = &pb.ResourceRequirements{}
			}
			if container.Resources.Limits == nil {
				container.Resources.Limits = make(map[string]string)
			}
			if _, ok := container.Resources.Limits["cpu"]; !ok {
				container.Resources.Limits["cpu"] = "1000m"
			}
			if _, ok := container.Resources.Limits["memory"]; !ok {
				container.Resources.Limits["memory"] = "512Mi"
			}
		}
	}
}

func (s *workerPools) CreateWorkerPool(ctx context.Context, req *pb.CreateWorkerPoolRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/workerPools/" + req.WorkerPoolId
	name, err := s.parseWorkerPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.WorkerPool)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = fields.ComputeWeakEtag(obj)
	obj.Creator = fmt.Sprintf("overseer-kcc-tester@%s.iam.gserviceaccount.com", name.Project.ID)
	obj.LastModifier = fmt.Sprintf("overseer-kcc-tester@%s.iam.gserviceaccount.com", name.Project.ID)

	obj.Uid = uuid.NewString()
	obj.Generation = 1

	if obj.Template == nil {
		obj.Template = &pb.WorkerPoolRevisionTemplate{}
	}

	// Set TerminalCondition to Ready if not set (Mocking immediate success)
	if obj.TerminalCondition == nil {
		obj.TerminalCondition = &pb.Condition{
			LastTransitionTime: timestamppb.Now(),
			State:              pb.Condition_CONDITION_SUCCEEDED,
			Type:               "Ready",
		}
	}

	s.populateWorkerPoolDefaults(name, obj, false, false)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	initialMeta := proto.CloneOf(obj)
	initialMeta.CustomAudiences = nil
	initialMeta.Reconciling = true
	initialMeta.InstanceSplitStatuses = nil
	initialMeta.LatestCreatedRevision = ""
	initialMeta.LatestReadyRevision = ""
	initialMeta.ObservedGeneration = 0
	initialMeta.TerminalCondition = &pb.Condition{
		State: pb.Condition_CONDITION_PENDING,
	}

	return s.operations.StartLRO(ctx, req.Parent, initialMeta, func() (proto.Message, error) {
		initialMeta.Reconciling = false
		initialMeta.InstanceSplitStatuses = obj.InstanceSplitStatuses
		initialMeta.LatestCreatedRevision = obj.LatestCreatedRevision
		initialMeta.LatestReadyRevision = obj.LatestReadyRevision
		initialMeta.ObservedGeneration = obj.ObservedGeneration
		initialMeta.TerminalCondition = obj.TerminalCondition

		ret := proto.CloneOf(obj)
		ret.CustomAudiences = nil
		return ret, nil
	})
}

func (s *workerPools) DeleteWorkerPool(ctx context.Context, req *pb.DeleteWorkerPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkerPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' of kind 'WORKER_POOL' in region '%s' in project '%s' does not exist.", name.WorkerPool, name.Location, name.Project.ID)
		}
		return nil, err
	}
	if err := s.storage.Delete(ctx, fqn, &pb.WorkerPool{}); err != nil {
		return nil, err
	}

	oldObservedGeneration := obj.ObservedGeneration
	obj.DeleteTime = timestamppb.Now()
	obj.ExpireTime = timestamppb.Now()
	obj.Generation++
	obj.ObservedGeneration = obj.Generation

	initialMeta := proto.CloneOf(obj)
	initialMeta.CustomAudiences = nil
	initialMeta.Reconciling = true
	initialMeta.ObservedGeneration = oldObservedGeneration
	initialMeta.TerminalCondition = &pb.Condition{
		LastTransitionTime: timestamppb.Now(),
		State:              pb.Condition_CONDITION_PENDING,
		Type:               "Ready",
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, initialMeta, func() (protoreflect.ProtoMessage, error) {
		initialMeta.Reconciling = false
		initialMeta.ObservedGeneration = obj.Generation
		initialMeta.TerminalCondition = &pb.Condition{
			LastTransitionTime: timestamppb.Now(),
			State:              pb.Condition_CONDITION_SUCCEEDED,
			Type:               "Ready",
		}

		ret := proto.CloneOf(obj)
		ret.CustomAudiences = nil
		return ret, nil
	})
}

func (s *workerPools) UpdateWorkerPool(ctx context.Context, req *pb.UpdateWorkerPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkerPoolName(req.GetWorkerPool().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkerPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' of kind 'WORKER_POOL' in region '%s' in project '%s' does not exist.", name.WorkerPool, name.Location, name.Project.ID)
		}
		return nil, err
	}

	oldTemplate := proto.CloneOf(obj.Template)
	oldLatestCreatedRevision := obj.LatestCreatedRevision
	oldLatestReadyRevision := obj.LatestReadyRevision
	oldObservedGeneration := obj.ObservedGeneration
	var oldInstanceSplitStatuses []*pb.InstanceSplitStatus
	for _, status := range obj.InstanceSplitStatuses {
		oldInstanceSplitStatuses = append(oldInstanceSplitStatuses, proto.Clone(status).(*pb.InstanceSplitStatus))
	}
	updated := req.GetWorkerPool()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// Basic update logic fallback
		if updated.Labels != nil {
			obj.Labels = updated.Labels
		}
		if updated.Annotations != nil {
			obj.Annotations = updated.Annotations
		}
		if updated.Template != nil {
			obj.Template = updated.Template
		}
		if updated.Description != "" {
			obj.Description = updated.Description
		}
		if updated.Scaling != nil {
			obj.Scaling = updated.Scaling
		}
		if updated.LaunchStage != 0 {
			obj.LaunchStage = updated.LaunchStage
		}
		if updated.CustomAudiences != nil {
			obj.CustomAudiences = updated.CustomAudiences
		}
	} else {
		if err := fields.UpdateByFieldMask(obj, updated, paths); err != nil {
			return nil, status.Errorf(codes.Internal, "updating by field mask: %v", err)
		}
	}

	obj.UpdateTime = timestamppb.Now()
	templateChanged := !proto.Equal(oldTemplate, obj.Template)
	if templateChanged {
		obj.Generation++
	}

	s.populateWorkerPoolDefaults(name, obj, true, templateChanged)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	if !templateChanged {
		meta := proto.CloneOf(obj)
		meta.CustomAudiences = nil
		ret := proto.CloneOf(obj)
		ret.CustomAudiences = nil
		return s.operations.DoneLRO(ctx, lroPrefix, meta, ret)
	}

	initialMeta := proto.CloneOf(obj)
	initialMeta.CustomAudiences = nil
	initialMeta.Reconciling = true
	initialMeta.ObservedGeneration = oldObservedGeneration
	initialMeta.LatestCreatedRevision = oldLatestCreatedRevision
	initialMeta.LatestReadyRevision = oldLatestReadyRevision
	initialMeta.InstanceSplitStatuses = oldInstanceSplitStatuses
	initialMeta.TerminalCondition = &pb.Condition{
		LastTransitionTime: timestamppb.Now(),
		State:              pb.Condition_CONDITION_PENDING,
		Type:               "Ready",
	}

	return s.operations.StartLRO(ctx, lroPrefix, initialMeta, func() (protoreflect.ProtoMessage, error) {
		initialMeta.Reconciling = false
		initialMeta.ObservedGeneration = obj.ObservedGeneration
		initialMeta.LatestCreatedRevision = obj.LatestCreatedRevision
		initialMeta.LatestReadyRevision = obj.LatestReadyRevision
		initialMeta.InstanceSplitStatuses = obj.InstanceSplitStatuses
		initialMeta.TerminalCondition = &pb.Condition{
			LastTransitionTime: timestamppb.Now(),
			State:              pb.Condition_CONDITION_SUCCEEDED,
			Type:               "Ready",
		}

		ret := proto.CloneOf(obj)
		ret.CustomAudiences = nil
		return ret, nil
	})
}

func (s *workerPools) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}

func (s *workerPools) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}

func (s *workerPools) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}

type workerPoolName struct {
	Project    *projects.ProjectData
	Location   string
	WorkerPool string
}

func (n *workerPoolName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/workerPools/%s", n.Project.ID, n.Location, n.WorkerPool)
}

func (s *MockService) parseWorkerPoolName(name string) (*workerPoolName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workerPools" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &workerPoolName{
			Project:    project,
			Location:   tokens[3],
			WorkerPool: tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
