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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/spanner/admin/instance/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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
	s.populateDefaultsForSpannerInstance(obj, obj)
	obj.State = pb.Instance_READY
	obj.DefaultBackupScheduleType = pb.Instance_AUTOMATIC
	if obj.Edition == pb.Instance_EDITION_UNSPECIFIED {
		obj.Edition = pb.Instance_STANDARD
	}

	// Metadata instance include ReplicaComputeCapacity even if not specify
	cloneObj := proto.Clone(obj).(*pb.Instance)
	s.populateReplicaComputeCapacityForSpannerInstance(cloneObj)

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
		metadata.Instance.UpdateTime = now
		metadata.Instance.ReplicaComputeCapacity = nil
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
	if 1000*update.NodeCount > update.ProcessingUnits {
		obj.ProcessingUnits = 1000 * update.NodeCount
		obj.NodeCount = update.NodeCount
	} else {
		obj.ProcessingUnits = update.ProcessingUnits
		obj.NodeCount = update.ProcessingUnits / 1000
	}
}

func (s *SpannerInstanceV1) populateReplicaComputeCapacityForSpannerInstance(obj *pb.Instance) {
	if len(obj.ReplicaComputeCapacity) == 0 {
		tokens := strings.Split(obj.Config, "/")
		var location string
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instanceConfigs" {
			location = strings.TrimPrefix(tokens[3], "regional-")
		}
		r := &pb.ReplicaComputeCapacity{
			ReplicaSelection: &pb.ReplicaSelection{Location: location},
			ComputeCapacity:  &pb.ReplicaComputeCapacity_NodeCount{NodeCount: obj.NodeCount},
		}
		obj.ReplicaComputeCapacity = append(obj.ReplicaComputeCapacity, r)
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
	for _, path := range req.GetFieldMask().GetPaths() {
		switch path {
		case "display_name":
			obj.DisplayName = updated.DisplayName
		case "edition":
			if updated.Edition == pb.Instance_EDITION_UNSPECIFIED && obj.Edition > updated.Edition {
				return nil, fmt.Errorf("Cannot downgrade edition from %s to %s", obj.Edition, updated.Edition)
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
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	s.populateDefaultsForSpannerInstance(req.Instance, obj)
	// Metadata instance include ReplicaComputeCapacity even if not specify
	cloneObj := proto.Clone(obj).(*pb.Instance)
	s.populateReplicaComputeCapacityForSpannerInstance(cloneObj)
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
		metadata.Instance.ReplicaComputeCapacity = nil
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
