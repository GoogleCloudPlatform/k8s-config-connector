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

package mockspannerinstance

import (
	"context"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
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

func (s *SpannerInstanceV1) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(404, "Instance not found: "+name.String())
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
	now := time.Now()

	obj := proto.Clone(req.GetInstance()).(*pb.Instance)
	obj.Name = fqn
	obj.Config = req.Instance.Config
	obj.NodeCount = req.Instance.NodeCount
	obj.AutoscalingConfig = req.Instance.AutoscalingConfig
	obj.DisplayName = req.Instance.DisplayName
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *SpannerInstanceV1) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseInstanceName(req.Instance.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}
	now := time.Now()

	obj := proto.Clone(req.GetInstance()).(*pb.Instance)
	obj.Name = fqn
	obj.DisplayName = existing.GetDisplayName()
	obj.NodeCount = existing.GetNodeCount()
	obj.ProcessingUnits = existing.GetProcessingUnits()
	obj.Config = existing.GetConfig()
	// obj.State = existing.State
	obj.CreateTime = existing.CreateTime
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		panic("YUWEN!!!!!!!!! " + err.Error())
		return nil, err
	}
	return s.operations.NewLRO(ctx)
}

func (s *SpannerInstanceV1) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*emptypb.Empty, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &instancepb.Instance{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
