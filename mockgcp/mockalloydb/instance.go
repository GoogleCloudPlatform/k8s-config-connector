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

package mockalloydb

import (
	"context"
	"fmt"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/alloydb/v1beta"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func (s *AlloyDBAdminV1) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func setInstanceFields(name *instanceName, obj *pb.Instance) {
	// Set default values to optional fields when unset.
	if obj.ClientConnectionConfig == nil {
		obj.ClientConnectionConfig = &pb.Instance_ClientConnectionConfig{
			SslConfig: &pb.SslConfig{
				SslMode: pb.SslConfig_ENCRYPTED_ONLY,
			},
		}
	}
	if obj.GeminiConfig == nil {
		obj.GeminiConfig = &pb.GeminiInstanceConfig{}
	}
	if obj.ObservabilityConfig == nil {
		obj.ObservabilityConfig = &pb.Instance_ObservabilityInstanceConfig{
			Enabled:               direct.PtrTo(false),
			MaxQueryStringLength:  direct.PtrTo(int32(10240)),
			PreserveComments:      direct.PtrTo(false),
			QueryPlansPerMinute:   direct.PtrTo(int32(20)),
			RecordApplicationTags: direct.PtrTo(false),
			TrackActiveQueries:    direct.PtrTo(false),
			TrackClientAddress:    direct.PtrTo(false),
			TrackWaitEventTypes:   direct.PtrTo(true),
			TrackWaitEvents:       direct.PtrTo(true),
		}
	}
	if obj.QueryInsightsConfig == nil {
		obj.QueryInsightsConfig = &pb.Instance_QueryInsightsInstanceConfig{
			QueryPlansPerMinute:   direct.PtrTo(uint32(5)),
			QueryStringLength:     uint32(1024),
			RecordApplicationTags: direct.PtrTo(false),
			RecordClientAddress:   direct.PtrTo(false),
		}
		if obj.InstanceType == pb.Instance_SECONDARY {
			obj.QueryInsightsConfig.RecordApplicationTags = direct.PtrTo(true)
			obj.QueryInsightsConfig.RecordClientAddress = direct.PtrTo(true)
		}
	}
	if obj.InstanceType != pb.Instance_READ_POOL &&
		obj.AvailabilityType == pb.Instance_AVAILABILITY_TYPE_UNSPECIFIED {
		obj.AvailabilityType = pb.Instance_REGIONAL
	}

	// Set output-only fields.
	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.IpAddress = "10.1.2.3"
	obj.State = pb.Instance_READY
	obj.Uid = "12345678"
	if obj.InstanceType != pb.Instance_READ_POOL {
		obj.WritableNode = &pb.Instance_Node{
			ZoneId: fmt.Sprintf("%v-b", name.Location),
		}
		if obj.AvailabilityType != pb.Instance_ZONAL {
			obj.Nodes = []*pb.Instance_Node{
				{
					ZoneId: fmt.Sprintf("%v-c", name.Location),
				},
			}
		}
	}
	if obj.NetworkConfig != nil && obj.NetworkConfig.EnableOutboundPublicIp {
		obj.OutboundPublicIpAddresses = []string{
			"35.228.195.235",
			"34.88.204.106",
		}
	}
	if obj.NetworkConfig != nil && obj.NetworkConfig.EnablePublicIp {
		obj.PublicIpAddress = "34.88.151.45"
	}
}

func (s *AlloyDBAdminV1) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/instances/" + req.GetInstanceId()
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = fqn
	setInstanceFields(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "create")
	return s.operations.StartLRO(ctx, name.ProjectAndLocation(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Instance)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *AlloyDBAdminV1) CreateSecondaryInstance(ctx context.Context, req *pb.CreateSecondaryInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/instances/" + req.GetInstanceId()
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = fqn
	setInstanceFields(name, obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "createsecondary")
	return s.operations.StartLRO(ctx, name.ProjectAndLocation(), metadata, func() (proto.Message, error) {
		result := proto.Clone(obj).(*pb.Instance)
		metadata.EndTime = timestamppb.Now()
		return result, nil
	})
}

func (s *AlloyDBAdminV1) UpdateInstance(ctx context.Context, req *pb.UpdateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.GetInstance().GetName()

	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.Instance.GetLabels()
		case "annotations":
			obj.Annotations = req.Instance.GetAnnotations()
		case "displayName":
			obj.DisplayName = req.Instance.GetDisplayName()
		case "gceZone":
			obj.GceZone = req.Instance.GetGceZone()
		case "databaseFlags":
			obj.DatabaseFlags = req.Instance.GetDatabaseFlags()
		case "availabilityType":
			obj.AvailabilityType = req.Instance.GetAvailabilityType()
		case "readPoolConfig":
			obj.ReadPoolConfig = req.Instance.GetReadPoolConfig()
		case "machineConfig":
			obj.MachineConfig = req.Instance.GetMachineConfig()
		case "pscInstanceConfig":
			obj.PscInstanceConfig = req.Instance.GetPscInstanceConfig()
		case "networkConfig":
			obj.NetworkConfig = req.Instance.GetNetworkConfig()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "update")
	return s.operations.StartLRO(ctx, name.ProjectAndLocation(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		result := proto.Clone(obj).(*pb.Instance)
		return result, nil
	})
}

func (s *AlloyDBAdminV1) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "delete")
	return s.operations.StartLRO(ctx, name.ProjectAndLocation(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := &emptypb.Empty{}
		return result, nil
	})
}
