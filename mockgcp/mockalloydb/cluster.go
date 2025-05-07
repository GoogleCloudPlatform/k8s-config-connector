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
	"strings"

	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/genproto/googleapis/type/dayofweek"
	"google.golang.org/genproto/googleapis/type/timeofday"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/alloydb/v1beta"
)

type AlloyDBAdminV1 struct {
	*MockService
	pb.UnimplementedAlloyDBAdminServer
}

func (s *AlloyDBAdminV1) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	updateNetworkInResponse(obj)
	return obj, nil
}

func setClusterFields(name *clusterName, obj *pb.Cluster) {
	// Remove unreadable field.
	obj.InitialUser = nil
	obj.DisplayName = ""
	// Set default values to optional fields when unset.
	if obj.AutomatedBackupPolicy == nil {
		obj.AutomatedBackupPolicy = &pb.AutomatedBackupPolicy{}
	}
	if obj.AutomatedBackupPolicy.BackupWindow == nil {
		obj.AutomatedBackupPolicy.BackupWindow = PtrTo(duration.Duration{Seconds: 3600})
	}
	if obj.AutomatedBackupPolicy.Enabled == nil {
		obj.AutomatedBackupPolicy.Enabled = PtrTo(false)
	}
	if obj.AutomatedBackupPolicy.Location == "" {
		obj.AutomatedBackupPolicy.Location = name.Location
	}
	if obj.AutomatedBackupPolicy.Retention == nil {
		obj.AutomatedBackupPolicy.Retention = &pb.AutomatedBackupPolicy_TimeBasedRetention_{
			TimeBasedRetention: &pb.AutomatedBackupPolicy_TimeBasedRetention{
				RetentionPeriod: PtrTo((duration.Duration{Seconds: 1209600})),
			},
		}
	}
	if obj.AutomatedBackupPolicy.Schedule == nil {
		obj.AutomatedBackupPolicy.Schedule = &pb.AutomatedBackupPolicy_WeeklySchedule_{
			WeeklySchedule: &pb.AutomatedBackupPolicy_WeeklySchedule{
				DaysOfWeek: []dayofweek.DayOfWeek{
					dayofweek.DayOfWeek_MONDAY,
					dayofweek.DayOfWeek_TUESDAY,
					dayofweek.DayOfWeek_WEDNESDAY,
					dayofweek.DayOfWeek_THURSDAY,
					dayofweek.DayOfWeek_FRIDAY,
					dayofweek.DayOfWeek_SATURDAY,
					dayofweek.DayOfWeek_SUNDAY,
				},
				StartTimes: []*timeofday.TimeOfDay{
					{Hours: 23},
				},
			},
		}
	}

	if obj.ContinuousBackupConfig == nil {
		obj.ContinuousBackupConfig = &pb.ContinuousBackupConfig{
			Enabled:            PtrTo(true),
			RecoveryWindowDays: 14,
		}
	} else {
		if obj.ContinuousBackupConfig.RecoveryWindowDays == 0 {
			obj.ContinuousBackupConfig.RecoveryWindowDays = 14
		}
	}
	if obj.GeminiConfig == nil {
		obj.GeminiConfig = &pb.GeminiClusterConfig{}
	}
	if obj.SubscriptionType == pb.SubscriptionType_SUBSCRIPTION_TYPE_UNSPECIFIED {
		obj.SubscriptionType = pb.SubscriptionType_STANDARD
	}
	// Set output-only fields.
	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.ContinuousBackupInfo = &pb.ContinuousBackupInfo{
		EncryptionInfo: &pb.EncryptionInfo{
			EncryptionType: pb.EncryptionInfo_GOOGLE_DEFAULT_ENCRYPTION,
		},
		Schedule: []dayofweek.DayOfWeek{
			dayofweek.DayOfWeek_MONDAY,
			dayofweek.DayOfWeek_TUESDAY,
			dayofweek.DayOfWeek_WEDNESDAY,
			dayofweek.DayOfWeek_THURSDAY,
			dayofweek.DayOfWeek_FRIDAY,
			dayofweek.DayOfWeek_SATURDAY,
			dayofweek.DayOfWeek_SUNDAY,
		},
	}
	if *obj.ContinuousBackupConfig.Enabled {
		obj.ContinuousBackupInfo.EnabledTime = now
	}
	// DatabaseVersion field was output only a couple of years ago but became
	// configurable early 2024.
	// Context: https://github.com/hashicorp/terraform-provider-google/issues/16960
	// This field needs to be handled differently in mockgcp after we fix the
	// behavior in the controller.
	obj.DatabaseVersion = pb.DatabaseVersion_POSTGRES_15
	if obj.EncryptionConfig != nil && obj.EncryptionConfig.KmsKeyName != "" {
		obj.EncryptionInfo = &pb.EncryptionInfo{
			EncryptionType: pb.EncryptionInfo_CUSTOMER_MANAGED_ENCRYPTION,
		}
	} else {
		obj.EncryptionInfo = &pb.EncryptionInfo{
			EncryptionType: pb.EncryptionInfo_GOOGLE_DEFAULT_ENCRYPTION,
		}
	}
	obj.Reconciling = false
	obj.State = pb.Cluster_READY
	obj.Uid = "111111111111111111111"
	// TODO: Validate the logic for PrimaryConfig.
	// PrimaryConfig is probably set when the primary cluster has a secondary
	// cluster associated with it.
	// It is then set to be an empty struct (non-nil) when the associated
	// instance is deleted (?)
	//if obj.ClusterType == pb.Cluster_PRIMARY {
	//	obj.PrimaryConfig = &pb.Cluster_PrimaryConfig{}
	//}
}

func (s *AlloyDBAdminV1) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/clusters/" + req.ClusterId
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Cluster).(*pb.Cluster)
	obj.Name = fqn
	setClusterFields(name, obj)
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "create")
	return s.operations.StartLRO(ctx, req.Parent, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := proto.Clone(obj).(*pb.Cluster)
		updateNetworkInResponse(result)
		return result, nil
	})
}

func (s *AlloyDBAdminV1) CreateSecondaryCluster(ctx context.Context, req *pb.CreateSecondaryClusterRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/clusters/" + req.ClusterId
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Cluster).(*pb.Cluster)
	obj.Name = fqn
	setClusterFields(name, obj)

	// Configure primaryConfig for the target cluster's primary cluster.
	primaryCluster := &pb.Cluster{}
	primaryClusterName, err := s.parseClusterName(obj.SecondaryConfig.PrimaryClusterName)
	if err != nil {
		return nil, err
	}
	if err := s.storage.Get(ctx, primaryClusterName.String(), primaryCluster); err != nil {
		return nil, err
	}
	if primaryCluster.PrimaryConfig == nil {
		primaryCluster.PrimaryConfig = &pb.Cluster_PrimaryConfig{
			SecondaryClusterNames: make([]string, 0),
		}
	}
	secondaryClusterNames := append(primaryCluster.PrimaryConfig.SecondaryClusterNames, fqn)
	primaryCluster.PrimaryConfig.SecondaryClusterNames = secondaryClusterNames
	if err := s.storage.Update(ctx, primaryClusterName.String(), primaryCluster); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "createsecondary")
	return s.operations.StartLRO(ctx, req.Parent, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := proto.Clone(obj).(*pb.Cluster)
		updateNetworkInResponse(result)
		updateSecondaryConfigInResponse(result)
		return result, nil
	})
}

func (s *AlloyDBAdminV1) RestoreCluster(ctx context.Context, req *pb.RestoreClusterRequest) (*longrunning.Operation, error) {
	// TODO: Implement it once req contains cluster ID.
	return nil, status.Errorf(codes.Unimplemented, "RestoreCluster not implemented yet")
}

func (s *AlloyDBAdminV1) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunning.Operation, error) {
	reqName := req.GetCluster().GetName()
	name, err := s.parseClusterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	for _, path := range paths {
		tokens := strings.Split(path, ".")
		switch tokens[0] {
		case "automatedBackupPolicy":
			obj.AutomatedBackupPolicy = req.Cluster.GetAutomatedBackupPolicy()
		case "continuousBackupConfig":
			obj.ContinuousBackupConfig = req.Cluster.GetContinuousBackupConfig()
		case "displayName":
			obj.DisplayName = req.Cluster.DisplayName
		case "initialUser":
			obj.InitialUser = req.Cluster.InitialUser
		case "labels":
			obj.Labels = req.Cluster.GetLabels()
		case "maintenanceUpdatePolicy":
			obj.MaintenanceUpdatePolicy = req.Cluster.GetMaintenanceUpdatePolicy()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	setClusterFields(name, obj)
	if *obj.ContinuousBackupConfig.Enabled {
		obj.ContinuousBackupInfo.EnabledTime = timestamppb.Now()
	} else {
		obj.ContinuousBackupInfo.EnabledTime = nil
	}
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := constructOperationMetadata(fqn, "update")
	return s.operations.StartLRO(ctx, name.ProjectAndLocation(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		result := proto.Clone(obj).(*pb.Cluster)
		updateNetworkInResponse(result)
		updateSecondaryConfigInResponse(result)
		return result, nil
	})
}

func (s *AlloyDBAdminV1) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunning.Operation, error) {
	name, err := s.parseClusterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Cluster{}
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

func updateNetworkInResponse(obj *pb.Cluster) {
	if (obj.NetworkConfig == nil || obj.NetworkConfig.Network == "") && obj.Network == "" {
		return
	}

	networkVal := ""
	if obj.NetworkConfig != nil && obj.NetworkConfig.Network != "" {
		networkVal = obj.NetworkConfig.Network
	} else if obj.Network != "" {
		networkVal = obj.Network
	}
	// Replace projectID with projectNumber for project "mock-project".
	networkVal = strings.ReplaceAll(networkVal, "mock-project", "518915279")
	obj.Network = networkVal
	networkConfig := &pb.Cluster_NetworkConfig{
		Network: networkVal,
	}
	if obj.NetworkConfig != nil && obj.NetworkConfig.AllocatedIpRange != "" {
		networkConfig.AllocatedIpRange = obj.NetworkConfig.AllocatedIpRange
	}
	obj.NetworkConfig = networkConfig
}

func updateSecondaryConfigInResponse(obj *pb.Cluster) {
	if obj.SecondaryConfig == nil {
		return
	}
	// Replace projectID with projectNumber for project "mock-project".
	primaryClusterName := strings.ReplaceAll(obj.SecondaryConfig.PrimaryClusterName, "mock-project", "518915279")
	obj.SecondaryConfig.PrimaryClusterName = primaryClusterName
}
