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

// +tool:mockgcp-support-spanner
// proto.service: google.spanner.admin.instance.v1.InstanceAdmin
// proto.message: google.spanner.admin.instance.v1.Instance

package mockspanner

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/spanner/admin/instance/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ pb.InstanceAdminServer = &SpannerInstanceV1{}

type SpannerInstanceV1 struct {
	*MockService
	pb.UnimplementedInstanceAdminServer
}

type spannerInstanceName struct {
	Project      *projects.ProjectData
	InstanceName string
}

func (s *SpannerInstanceV1) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Instance not found: %s", name.String())
		}
		return nil, err
	}
	return obj, nil
}

func (s *SpannerInstanceV1) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunningpb.Operation, error) {
	instanceName := req.GetParent() + "/instances/" + req.GetInstanceId()
	name, err := s.parseInstanceName(instanceName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	now := timestamppb.Now()

	obj := proto.Clone(req.GetInstance()).(*pb.Instance)
	// Value of ReplicaComputeCapacity during creation is determined by raw
	// input so it needs to be processed first.
	s.populateReplicaComputeCapacityForSpannerInstance(obj)
	s.populateDefaultsForSpannerInstance(obj, obj)
	obj.State = pb.Instance_READY
	if len(obj.DisplayName) < 4 || len(obj.DisplayName) > 30 {
		return nil, fmt.Errorf("Display name must be between 4-30 characters long")
	}
	if obj.Edition == pb.Instance_EDITION_UNSPECIFIED {
		obj.Edition = pb.Instance_STANDARD
	}
	if obj.DefaultBackupScheduleType == pb.Instance_DEFAULT_BACKUP_SCHEDULE_TYPE_UNSPECIFIED &&
		obj.InstanceType != pb.Instance_FREE_INSTANCE {
		obj.DefaultBackupScheduleType = pb.Instance_AUTOMATIC
	}
	obj.Name = fqn

	cloneObj := proto.Clone(obj).(*pb.Instance)

	obj.CreateTime = now
	obj.UpdateTime = now
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.CreateInstanceMetadata{
		Instance:  cloneObj,
		StartTime: now,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		metadata.ExpectedFulfillmentPeriod = pb.FulfillmentPeriod_FULFILLMENT_PERIOD_NORMAL
		metadata.EndTime = now
		metadata.Instance.CreateTime = now
		metadata.Instance.UpdateTime = now
		metadata.Instance.Name = fqn
		retObj := proto.Clone(obj).(*pb.Instance)
		retObj.Name = fqn
		return retObj, nil
	})
}

func (s *SpannerInstanceV1) populateDefaultsForSpannerInstance(update, obj *pb.Instance) {
	// At most one of either node_count or processing_units should be present.
	// https://cloud.google.com/spanner/docs/compute-capacity
	// 1 nodeCount equals 1000 processingUnits
	if update.AutoscalingConfig != nil {
		update.ProcessingUnits = update.GetAutoscalingConfig().AutoscalingLimits.GetMinProcessingUnits()
		update.NodeCount = update.GetAutoscalingConfig().AutoscalingLimits.GetMinNodes()
	}
	// If either nodeCount or processingUnits fields are unset, it means some
	// changes need to happen. Otherwise, keep obj as is.
	if update.NodeCount != 0 || update.ProcessingUnits != 0 {
		if 1000*update.NodeCount > update.ProcessingUnits {
			obj.ProcessingUnits = 1000 * update.NodeCount
			obj.NodeCount = update.NodeCount
		} else {
			obj.ProcessingUnits = update.ProcessingUnits
			obj.NodeCount = update.ProcessingUnits / 1000
		}
	}

	if update.InstanceType == pb.Instance_INSTANCE_TYPE_UNSPECIFIED {
		obj.InstanceType = pb.Instance_PROVISIONED
	}
}

func (s *SpannerInstanceV1) populateReplicaComputeCapacityForSpannerInstance(obj *pb.Instance) {
	if len(obj.ReplicaComputeCapacity) == 0 {
		obj.ReplicaComputeCapacity = append(obj.ReplicaComputeCapacity, &pb.ReplicaComputeCapacity{})
	}

	tokens := strings.Split(obj.Config, "/")
	var location string
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instanceConfigs" {
		location = strings.TrimPrefix(tokens[3], "regional-")
	}
	obj.ReplicaComputeCapacity[0].ReplicaSelection = &pb.ReplicaSelection{Location: location}

	// When creating an instance, ReplicaComputeCapacity is initialized:
	// - if autoscalingLimits are configured,
	//   - if minNodes is set, set ReplicaComputeCapacity_NodeCount to minNodes
	//   - else set ReplicaComputeCapacity_ProcessingUnits to minProcessingUnits
	// - else
	//   - if nodeCount is set, set ReplicaComputeCapacity_NodeCount to nodeCount
	//   - else set ReplicaComputeCapacity_ProcessingUnits to processingUnits
	// When updating an instance, ReplicaComputeCapacity is updated:
	//   type of ComputeCapacity is not changed, but the value of it is updated accordingly.

	// Check if the computeCapacity is already set.
	var useNodeCount bool
	var useProcessingUnits bool
	if obj.ReplicaComputeCapacity[0].ComputeCapacity != nil {
		if _, ok := obj.ReplicaComputeCapacity[0].ComputeCapacity.(*pb.ReplicaComputeCapacity_NodeCount); ok {
			useNodeCount = true
		} else {
			useProcessingUnits = true
		}
	}

	// Instance creation
	if !(useNodeCount || useProcessingUnits) {
		if obj.AutoscalingConfig != nil && obj.AutoscalingConfig.AutoscalingLimits != nil {
			limits := obj.AutoscalingConfig.AutoscalingLimits
			if _, ok := limits.GetMinLimit().(*pb.AutoscalingConfig_AutoscalingLimits_MinNodes); ok {
				obj.ReplicaComputeCapacity[0].ComputeCapacity = &pb.ReplicaComputeCapacity_NodeCount{NodeCount: limits.GetMinNodes()}
			} else {
				obj.ReplicaComputeCapacity[0].ComputeCapacity = &pb.ReplicaComputeCapacity_ProcessingUnits{ProcessingUnits: limits.GetMinProcessingUnits()}
			}
		} else {
			if obj.NodeCount != 0 {
				obj.ReplicaComputeCapacity[0].ComputeCapacity = &pb.ReplicaComputeCapacity_NodeCount{NodeCount: obj.GetNodeCount()}
			} else {
				obj.ReplicaComputeCapacity[0].ComputeCapacity = &pb.ReplicaComputeCapacity_ProcessingUnits{ProcessingUnits: obj.GetProcessingUnits()}
			}
		}
	} else { // Instance Update
		if obj.AutoscalingConfig != nil && obj.AutoscalingConfig.AutoscalingLimits != nil {
			limits := obj.AutoscalingConfig.AutoscalingLimits
			if useNodeCount {
				obj.ReplicaComputeCapacity[0].ComputeCapacity = &pb.ReplicaComputeCapacity_NodeCount{NodeCount: limits.GetMinNodes()}
			} else {
				obj.ReplicaComputeCapacity[0].ComputeCapacity = &pb.ReplicaComputeCapacity_ProcessingUnits{ProcessingUnits: limits.GetMinProcessingUnits()}
			}
		} else {
			if useNodeCount {
				obj.ReplicaComputeCapacity[0].ComputeCapacity = &pb.ReplicaComputeCapacity_NodeCount{NodeCount: obj.GetNodeCount()}
			} else {
				obj.ReplicaComputeCapacity[0].ComputeCapacity = &pb.ReplicaComputeCapacity_ProcessingUnits{ProcessingUnits: obj.GetProcessingUnits()}
			}
		}
	}
}

func (s *SpannerInstanceV1) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Instance.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	obj.UpdateTime = now
	updated := req.Instance
	paths := req.GetFieldMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "display_name":
			if len(updated.DisplayName) < 4 || len(updated.DisplayName) > 30 {
				return nil, fmt.Errorf("Display name must be between 4-30 characters long")
			}
			obj.DisplayName = updated.DisplayName
		case "edition":
			if obj.Edition > updated.Edition && len(paths) > 1 {
				return nil, fmt.Errorf(
					"Cannot downgrade %s from %s to %s in the same request as other updates. The field mask contains the following paths: %s. Please send a separate request for Edition downgrade",
					fqn,
					obj.Edition.String(),
					updated.Edition.String(),
					strings.Join(req.GetFieldMask().GetPaths(), ","),
				)
			}
			obj.Edition = updated.Edition
		case "labels":
			obj.Labels = updated.Labels
		case "node_count":
			obj.NodeCount = updated.NodeCount
		case "processing_units":
			obj.ProcessingUnits = updated.ProcessingUnits
		case "autoscaling_config":
			obj.AutoscalingConfig = updated.AutoscalingConfig
		case "default_backup_schedule_type":
			obj.DefaultBackupScheduleType = updated.DefaultBackupScheduleType
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	s.populateDefaultsForSpannerInstance(updated, obj)
	// Value of ReplicaComputeCapacity during update is determined by processed
	// input so it needs to be handled after defaults are populated.
	s.populateReplicaComputeCapacityForSpannerInstance(obj)
	cloneObj := proto.Clone(obj).(*pb.Instance)
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	metadata := &pb.UpdateInstanceMetadata{
		Instance:  cloneObj,
		StartTime: now,
	}
	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		metadata.ExpectedFulfillmentPeriod = pb.FulfillmentPeriod_FULFILLMENT_PERIOD_NORMAL
		metadata.EndTime = now
		metadata.Instance.UpdateTime = now
		retObj := proto.Clone(obj).(*pb.Instance)
		return retObj, nil
	})
}

func (s *SpannerInstanceV1) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*emptypb.Empty, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, existing); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SpannerInstanceV1) GetInstanceConfig(ctx context.Context, req *pb.GetInstanceConfigRequest) (*pb.InstanceConfig, error) {
	name, err := s.parseInstanceConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Instance config not found: %s", name.String())
		}
		return nil, err
	}

	return obj, nil
}

func (s *SpannerInstanceV1) CreateInstanceConfig(ctx context.Context, req *pb.CreateInstanceConfigRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceConfigName(req.GetInstanceConfig().Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.InstanceConfig{
		Name:        fqn,
		BaseConfig:  req.InstanceConfig.BaseConfig,
		DisplayName: req.InstanceConfig.DisplayName,
		Replicas:    req.InstanceConfig.Replicas,
	}
	if obj.Etag == "" {
		obj.Etag = fields.ComputeWeakEtag(obj)
	}
	obj.ConfigType = pb.InstanceConfig_USER_MANAGED
	obj.State = pb.InstanceConfig_READY

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := name.String()
	lroMetadata := &pb.CreateInstanceConfigMetadata{
		InstanceConfig: obj,
		Progress:       &pb.OperationProgress{StartTime: timestamppb.Now()},
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.Progress.EndTime = timestamppb.Now()
		lroMetadata.Progress.ProgressPercent = 100
		return obj, nil
	})
}

func (s *SpannerInstanceV1) ListInstanceConfigs(ctx context.Context, req *pb.ListInstanceConfigsRequest) (*pb.ListInstanceConfigsResponse, error) {
	prefix := req.GetParent() + "/instanceConfigs/"
	list := make([]*pb.InstanceConfig, 0)
	snapshotKind := (&pb.InstanceConfig{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, snapshotKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj protoreflect.ProtoMessage) error {
		snapshot, ok := obj.(*pb.InstanceConfig)
		if !ok {
			return status.Errorf(codes.Internal, "unexpected type %T in ListInstanceConfigs", obj)
		}
		if strings.HasPrefix(snapshot.Name, prefix) {
			list = append(list, snapshot)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.ListInstanceConfigsResponse{InstanceConfigs: list}, nil
}

func (s *SpannerInstanceV1) UpdateInstanceConfig(ctx context.Context, req *pb.UpdateInstanceConfigRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceConfigName(req.GetInstanceConfig().Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InstanceConfig{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = fqn

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch camelToUnderscore(path) {
		case "labels":
			obj.Labels = req.GetInstanceConfig().GetLabels()
		case "display_name":
			obj.DisplayName = req.GetInstanceConfig().GetDisplayName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "UpdateInstanceConfig does not support field mask path: %q", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := name.String()
	lroMetadata := &pb.UpdateInstanceConfigMetadata{
		InstanceConfig: obj,
		Progress:       &pb.OperationProgress{StartTime: timestamppb.Now()},
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *SpannerInstanceV1) DeleteInstanceConfig(ctx context.Context, req *pb.DeleteInstanceConfigRequest) (*emptypb.Empty, error) {
	name, err := s.parseInstanceConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.InstanceConfig{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type InstanceConfigName struct {
	Project *projects.ProjectData
	name    string
}

func (n *InstanceConfigName) String() string {
	return fmt.Sprintf("projects/%s/instanceConfigs/%s", n.Project.ID, n.name)
}

// parseInstanceConfigName parses a string into a instanceConfigName.
// The expected form is `projects/*/instanceConfigs/*`.
func (s *MockService) parseInstanceConfigName(name string) (*InstanceConfigName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instanceConfigs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &InstanceConfigName{
			Project: project,
			name:    tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
