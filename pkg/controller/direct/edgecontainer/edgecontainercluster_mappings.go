// Copyright 2026 Google LLC
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

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	status "google.golang.org/genproto/googleapis/rpc/status"
)

func EdgeContainerClusterFleet_FromProto(mapCtx *direct.MapContext, in *pb.Fleet) krm.EdgeContainerClusterFleet {
	if in == nil {
		return krm.EdgeContainerClusterFleet{}
	}
	out := krm.EdgeContainerClusterFleet{}
	out.Membership = direct.LazyPtr(in.GetMembership())
	if in.GetProject() != "" {
		out.ProjectRef.External = in.GetProject()
	}
	return out
}

func EdgeContainerClusterFleet_ToProto(mapCtx *direct.MapContext, in krm.EdgeContainerClusterFleet) *pb.Fleet {
	out := &pb.Fleet{}
	out.Membership = direct.ValueOf(in.Membership)
	if in.ProjectRef.External != "" {
		out.Project = in.ProjectRef.External
	}
	return out
}

func EdgeContainerClusterNetworking_FromProto(mapCtx *direct.MapContext, in *pb.ClusterNetworking) krm.EdgeContainerClusterNetworking {
	if in == nil {
		return krm.EdgeContainerClusterNetworking{}
	}
	out := krm.EdgeContainerClusterNetworking{}
	out.ClusterIpv4CidrBlocks = in.GetClusterIpv4CidrBlocks()
	out.ServicesIpv4CidrBlocks = in.GetServicesIpv4CidrBlocks()
	return out
}

func EdgeContainerClusterNetworking_ToProto(mapCtx *direct.MapContext, in krm.EdgeContainerClusterNetworking) *pb.ClusterNetworking {
	out := &pb.ClusterNetworking{}
	out.ClusterIpv4CidrBlocks = in.ClusterIpv4CidrBlocks
	out.ServicesIpv4CidrBlocks = in.ServicesIpv4CidrBlocks
	return out
}

func EdgeContainerClusterAuthorization_FromProto(mapCtx *direct.MapContext, in *pb.Authorization) krm.EdgeContainerClusterAuthorization {
	if in == nil {
		return krm.EdgeContainerClusterAuthorization{}
	}
	out := krm.EdgeContainerClusterAuthorization{}
	if in.GetAdminUsers() != nil {
		out.AdminUsers.UsernameRef.External = in.GetAdminUsers().GetUsername()
	}
	return out
}

func EdgeContainerClusterAuthorization_ToProto(mapCtx *direct.MapContext, in krm.EdgeContainerClusterAuthorization) *pb.Authorization {
	out := &pb.Authorization{}
	if in.AdminUsers.UsernameRef.External != "" {
		out.AdminUsers = &pb.ClusterUser{
			Username: in.AdminUsers.UsernameRef.External,
		}
	}
	return out
}

func EdgeContainerClusterMaintenancePolicy_FromProto(mapCtx *direct.MapContext, in *pb.MaintenancePolicy) *krm.EdgeContainerClusterMaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterMaintenancePolicy{}
	if in.GetWindow() != nil {
		out.Window.RecurringWindow.Recurrence = direct.LazyPtr(in.GetWindow().GetRecurringWindow().GetRecurrence())
		if in.GetWindow().GetRecurringWindow().GetWindow() != nil {
			out.Window.RecurringWindow.Window = &krm.EdgeContainerClusterTimeWindow{
				StartTime: direct.StringTimestamp_FromProto(mapCtx, in.GetWindow().GetRecurringWindow().GetWindow().GetStartTime()),
				EndTime:   direct.StringTimestamp_FromProto(mapCtx, in.GetWindow().GetRecurringWindow().GetWindow().GetEndTime()),
			}
		}
	}
	return out
}

func EdgeContainerClusterMaintenancePolicy_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterMaintenancePolicy) *pb.MaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.MaintenancePolicy{}
	out.Window = &pb.MaintenanceWindow{}
	out.Window.RecurringWindow = &pb.RecurringTimeWindow{}
	out.Window.RecurringWindow.Recurrence = direct.ValueOf(in.Window.RecurringWindow.Recurrence)
	if in.Window.RecurringWindow.Window != nil {
		out.Window.RecurringWindow.Window = &pb.TimeWindow{
			StartTime: direct.StringTimestamp_ToProto(mapCtx, in.Window.RecurringWindow.Window.StartTime),
			EndTime:   direct.StringTimestamp_ToProto(mapCtx, in.Window.RecurringWindow.Window.EndTime),
		}
	}
	return out
}

func EdgeContainerClusterControlPlane_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlane) *krm.EdgeContainerClusterControlPlane {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterControlPlane{}
	if in.GetLocal() != nil {
		out.Local = &krm.EdgeContainerClusterControlPlaneLocal{
			MachineFilter:          direct.LazyPtr(in.GetLocal().GetMachineFilter()),
			NodeCount:              direct.LazyPtr(in.GetLocal().GetNodeCount()),
			NodeLocation:           direct.LazyPtr(in.GetLocal().GetNodeLocation()),
			SharedDeploymentPolicy: direct.Enum_FromProto(mapCtx, in.GetLocal().GetSharedDeploymentPolicy()),
		}
	}
	if in.GetRemote() != nil {
		out.Remote = &krm.EdgeContainerClusterControlPlaneRemote{}
	}
	return out
}

func EdgeContainerClusterControlPlane_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterControlPlane) *pb.Cluster_ControlPlane {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlane{}
	if in.Local != nil {
		out.Config = &pb.Cluster_ControlPlane_Local_{
			Local: &pb.Cluster_ControlPlane_Local{
				MachineFilter:          direct.ValueOf(in.Local.MachineFilter),
				NodeCount:              direct.ValueOf(in.Local.NodeCount),
				NodeLocation:           direct.ValueOf(in.Local.NodeLocation),
				SharedDeploymentPolicy: direct.Enum_ToProto[pb.Cluster_ControlPlane_SharedDeploymentPolicy](mapCtx, in.Local.SharedDeploymentPolicy),
			},
		}
	} else if in.Remote != nil {
		out.Config = &pb.Cluster_ControlPlane_Remote_{
			Remote: &pb.Cluster_ControlPlane_Remote{},
		}
	}
	return out
}

func EdgeContainerClusterSystemAddonsConfig_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_SystemAddonsConfig) *krm.EdgeContainerClusterSystemAddonsConfig {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterSystemAddonsConfig{}
	if in.GetIngress() != nil {
		out.Ingress = &krm.EdgeContainerClusterSystemAddonsConfigIngress{
			Disabled: direct.LazyPtr(in.GetIngress().GetDisabled()),
			Ipv4Vip:  direct.LazyPtr(in.GetIngress().GetIpv4Vip()),
		}
	}
	return out
}

func EdgeContainerClusterSystemAddonsConfig_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterSystemAddonsConfig) *pb.Cluster_SystemAddonsConfig {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_SystemAddonsConfig{}
	if in.Ingress != nil {
		out.Ingress = &pb.Cluster_SystemAddonsConfig_Ingress{
			Disabled: direct.ValueOf(in.Ingress.Disabled),
			Ipv4Vip:  direct.ValueOf(in.Ingress.Ipv4Vip),
		}
	}
	return out
}

func EdgeContainerClusterControlPlaneEncryption_FromProto(mapCtx *direct.MapContext, in *pb.Cluster_ControlPlaneEncryption) *krm.EdgeContainerClusterControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterControlPlaneEncryption{}
	out.KmsKeyActiveVersion = direct.LazyPtr(in.GetKmsKeyActiveVersion())
	if in.GetKmsKey() != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{
			External: in.GetKmsKey(),
		}
	}
	out.KmsKeyState = direct.Enum_FromProto(mapCtx, in.GetKmsKeyState())
	if in.GetKmsStatus() != nil {
		out.KmsStatus = []krm.EdgeContainerClusterKmsStatus{*EdgeContainerClusterKmsStatus_FromProto(mapCtx, in.GetKmsStatus())}
	}
	return out
}

func EdgeContainerClusterControlPlaneEncryption_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterControlPlaneEncryption) *pb.Cluster_ControlPlaneEncryption {
	if in == nil {
		return nil
	}
	out := &pb.Cluster_ControlPlaneEncryption{}
	out.KmsKeyActiveVersion = direct.ValueOf(in.KmsKeyActiveVersion)
	if in.KmsKeyRef != nil {
		out.KmsKey = in.KmsKeyRef.External
	}
	out.KmsKeyState = direct.Enum_ToProto[pb.KmsKeyState](mapCtx, in.KmsKeyState)
	if len(in.KmsStatus) > 0 {
		out.KmsStatus = EdgeContainerClusterKmsStatus_ToProto(mapCtx, &in.KmsStatus[0])
	}
	return out
}

func EdgeContainerClusterKmsStatus_FromProto(mapCtx *direct.MapContext, in *status.Status) *krm.EdgeContainerClusterKmsStatus {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterKmsStatus{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}

func EdgeContainerClusterKmsStatus_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterKmsStatus) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}

func EdgeContainerClusterAdminUsers_FromProto(mapCtx *direct.MapContext, in *pb.ClusterUser) *krm.EdgeContainerClusterAdminUsers {
	if in == nil {
		return nil
	}
	out := &krm.EdgeContainerClusterAdminUsers{}
	if in.GetUsername() != "" {
		out.UsernameRef = refsv1beta1.IAMServiceAccountRef{External: in.GetUsername()}
	}
	return out
}

func EdgeContainerClusterAdminUsers_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterAdminUsers) *pb.ClusterUser {
	if in == nil {
		return nil
	}
	out := &pb.ClusterUser{}
	if in.UsernameRef.External != "" {
		out.Username = in.UsernameRef.External
	}
	return out
}

func EdgeContainerClusterMaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *pb.MaintenanceWindow) krm.EdgeContainerClusterMaintenanceWindow {
	if in == nil {
		return krm.EdgeContainerClusterMaintenanceWindow{}
	}
	out := krm.EdgeContainerClusterMaintenanceWindow{}
	out.RecurringWindow = EdgeContainerClusterRecurringTimeWindow_FromProto(mapCtx, in.GetRecurringWindow())
	return out
}

func EdgeContainerClusterMaintenanceWindow_ToProto(mapCtx *direct.MapContext, in krm.EdgeContainerClusterMaintenanceWindow) *pb.MaintenanceWindow {
	out := &pb.MaintenanceWindow{}
	out.RecurringWindow = EdgeContainerClusterRecurringTimeWindow_ToProto(mapCtx, &in.RecurringWindow)
	return out
}

func EdgeContainerClusterRecurringTimeWindow_FromProto(mapCtx *direct.MapContext, in *pb.RecurringTimeWindow) krm.EdgeContainerClusterRecurringTimeWindow {
	if in == nil {
		return krm.EdgeContainerClusterRecurringTimeWindow{}
	}
	out := krm.EdgeContainerClusterRecurringTimeWindow{}
	out.Recurrence = direct.LazyPtr(in.GetRecurrence())
	out.Window = EdgeContainerClusterTimeWindow_FromProto(mapCtx, in.GetWindow())
	return out
}

func EdgeContainerClusterRecurringTimeWindow_ToProto(mapCtx *direct.MapContext, in *krm.EdgeContainerClusterRecurringTimeWindow) *pb.RecurringTimeWindow {
	if in == nil {
		return nil
	}
	out := &pb.RecurringTimeWindow{}
	out.Recurrence = direct.ValueOf(in.Recurrence)
	out.Window = EdgeContainerClusterTimeWindow_ToProto(mapCtx, in.Window)
	return out
}
