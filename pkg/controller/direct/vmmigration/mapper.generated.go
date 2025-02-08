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

package vmmigration

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmmigration/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ApplianceVersion_FromProto(mapCtx *direct.MapContext, in *pb.ApplianceVersion) *krm.ApplianceVersion {
	if in == nil {
		return nil
	}
	out := &krm.ApplianceVersion{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Critical = direct.LazyPtr(in.GetCritical())
	out.ReleaseNotesURI = direct.LazyPtr(in.GetReleaseNotesUri())
	return out
}
func ApplianceVersion_ToProto(mapCtx *direct.MapContext, in *krm.ApplianceVersion) *pb.ApplianceVersion {
	if in == nil {
		return nil
	}
	out := &pb.ApplianceVersion{}
	out.Version = direct.ValueOf(in.Version)
	out.Uri = direct.ValueOf(in.URI)
	out.Critical = direct.ValueOf(in.Critical)
	out.ReleaseNotesUri = direct.ValueOf(in.ReleaseNotesURI)
	return out
}
func AvailableUpdates_FromProto(mapCtx *direct.MapContext, in *pb.AvailableUpdates) *krm.AvailableUpdates {
	if in == nil {
		return nil
	}
	out := &krm.AvailableUpdates{}
	out.NewDeployableAppliance = ApplianceVersion_FromProto(mapCtx, in.GetNewDeployableAppliance())
	out.InPlaceUpdate = ApplianceVersion_FromProto(mapCtx, in.GetInPlaceUpdate())
	return out
}
func AvailableUpdates_ToProto(mapCtx *direct.MapContext, in *krm.AvailableUpdates) *pb.AvailableUpdates {
	if in == nil {
		return nil
	}
	out := &pb.AvailableUpdates{}
	out.NewDeployableAppliance = ApplianceVersion_ToProto(mapCtx, in.NewDeployableAppliance)
	out.InPlaceUpdate = ApplianceVersion_ToProto(mapCtx, in.InPlaceUpdate)
	return out
}
func DatacenterConnector_FromProto(mapCtx *direct.MapContext, in *pb.DatacenterConnector) *krm.DatacenterConnector {
	if in == nil {
		return nil
	}
	out := &krm.DatacenterConnector{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Name
	out.RegistrationID = direct.LazyPtr(in.GetRegistrationId())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.Version = direct.LazyPtr(in.GetVersion())
	// MISSING: Bucket
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: ApplianceInfrastructureVersion
	// MISSING: ApplianceSoftwareVersion
	// MISSING: AvailableVersions
	// MISSING: UpgradeStatus
	return out
}
func DatacenterConnector_ToProto(mapCtx *direct.MapContext, in *krm.DatacenterConnector) *pb.DatacenterConnector {
	if in == nil {
		return nil
	}
	out := &pb.DatacenterConnector{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Name
	out.RegistrationId = direct.ValueOf(in.RegistrationID)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.Version = direct.ValueOf(in.Version)
	// MISSING: Bucket
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: ApplianceInfrastructureVersion
	// MISSING: ApplianceSoftwareVersion
	// MISSING: AvailableVersions
	// MISSING: UpgradeStatus
	return out
}
func DatacenterConnectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DatacenterConnector) *krm.DatacenterConnectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatacenterConnectorObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: RegistrationID
	// MISSING: ServiceAccount
	// MISSING: Version
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStateTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.ApplianceInfrastructureVersion = direct.LazyPtr(in.GetApplianceInfrastructureVersion())
	out.ApplianceSoftwareVersion = direct.LazyPtr(in.GetApplianceSoftwareVersion())
	out.AvailableVersions = AvailableUpdates_FromProto(mapCtx, in.GetAvailableVersions())
	out.UpgradeStatus = UpgradeStatus_FromProto(mapCtx, in.GetUpgradeStatus())
	return out
}
func DatacenterConnectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatacenterConnectorObservedState) *pb.DatacenterConnector {
	if in == nil {
		return nil
	}
	out := &pb.DatacenterConnector{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Name = direct.ValueOf(in.Name)
	// MISSING: RegistrationID
	// MISSING: ServiceAccount
	// MISSING: Version
	out.Bucket = direct.ValueOf(in.Bucket)
	out.State = direct.Enum_ToProto[pb.DatacenterConnector_State](mapCtx, in.State)
	out.StateTime = direct.StringTimestamp_ToProto(mapCtx, in.StateTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.ApplianceInfrastructureVersion = direct.ValueOf(in.ApplianceInfrastructureVersion)
	out.ApplianceSoftwareVersion = direct.ValueOf(in.ApplianceSoftwareVersion)
	out.AvailableVersions = AvailableUpdates_ToProto(mapCtx, in.AvailableVersions)
	out.UpgradeStatus = UpgradeStatus_ToProto(mapCtx, in.UpgradeStatus)
	return out
}
func UpgradeStatus_FromProto(mapCtx *direct.MapContext, in *pb.UpgradeStatus) *krm.UpgradeStatus {
	if in == nil {
		return nil
	}
	out := &krm.UpgradeStatus{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.PreviousVersion = direct.LazyPtr(in.GetPreviousVersion())
	return out
}
func UpgradeStatus_ToProto(mapCtx *direct.MapContext, in *krm.UpgradeStatus) *pb.UpgradeStatus {
	if in == nil {
		return nil
	}
	out := &pb.UpgradeStatus{}
	out.Version = direct.ValueOf(in.Version)
	out.State = direct.Enum_ToProto[pb.UpgradeStatus_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.PreviousVersion = direct.ValueOf(in.PreviousVersion)
	return out
}
func VmmigrationDatacenterConnectorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DatacenterConnector) *krm.VmmigrationDatacenterConnectorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationDatacenterConnectorObservedState{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Name
	// MISSING: RegistrationID
	// MISSING: ServiceAccount
	// MISSING: Version
	// MISSING: Bucket
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: ApplianceInfrastructureVersion
	// MISSING: ApplianceSoftwareVersion
	// MISSING: AvailableVersions
	// MISSING: UpgradeStatus
	return out
}
func VmmigrationDatacenterConnectorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationDatacenterConnectorObservedState) *pb.DatacenterConnector {
	if in == nil {
		return nil
	}
	out := &pb.DatacenterConnector{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Name
	// MISSING: RegistrationID
	// MISSING: ServiceAccount
	// MISSING: Version
	// MISSING: Bucket
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: ApplianceInfrastructureVersion
	// MISSING: ApplianceSoftwareVersion
	// MISSING: AvailableVersions
	// MISSING: UpgradeStatus
	return out
}
func VmmigrationDatacenterConnectorSpec_FromProto(mapCtx *direct.MapContext, in *pb.DatacenterConnector) *krm.VmmigrationDatacenterConnectorSpec {
	if in == nil {
		return nil
	}
	out := &krm.VmmigrationDatacenterConnectorSpec{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Name
	// MISSING: RegistrationID
	// MISSING: ServiceAccount
	// MISSING: Version
	// MISSING: Bucket
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: ApplianceInfrastructureVersion
	// MISSING: ApplianceSoftwareVersion
	// MISSING: AvailableVersions
	// MISSING: UpgradeStatus
	return out
}
func VmmigrationDatacenterConnectorSpec_ToProto(mapCtx *direct.MapContext, in *krm.VmmigrationDatacenterConnectorSpec) *pb.DatacenterConnector {
	if in == nil {
		return nil
	}
	out := &pb.DatacenterConnector{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Name
	// MISSING: RegistrationID
	// MISSING: ServiceAccount
	// MISSING: Version
	// MISSING: Bucket
	// MISSING: State
	// MISSING: StateTime
	// MISSING: Error
	// MISSING: ApplianceInfrastructureVersion
	// MISSING: ApplianceSoftwareVersion
	// MISSING: AvailableVersions
	// MISSING: UpgradeStatus
	return out
}
