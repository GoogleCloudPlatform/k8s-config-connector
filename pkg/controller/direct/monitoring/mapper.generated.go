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

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MonitoringServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.MonitoringServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringServiceObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Custom
	// MISSING: AppEngine
	// MISSING: CloudEndpoints
	// MISSING: ClusterIstio
	// MISSING: MeshIstio
	// MISSING: IstioCanonicalService
	// MISSING: CloudRun
	// MISSING: GkeNamespace
	// MISSING: GkeWorkload
	// MISSING: GkeService
	// MISSING: BasicService
	// MISSING: Telemetry
	// MISSING: UserLabels
	return out
}
func MonitoringServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringServiceObservedState) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Custom
	// MISSING: AppEngine
	// MISSING: CloudEndpoints
	// MISSING: ClusterIstio
	// MISSING: MeshIstio
	// MISSING: IstioCanonicalService
	// MISSING: CloudRun
	// MISSING: GkeNamespace
	// MISSING: GkeWorkload
	// MISSING: GkeService
	// MISSING: BasicService
	// MISSING: Telemetry
	// MISSING: UserLabels
	return out
}
func MonitoringServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.MonitoringServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringServiceSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Custom
	// MISSING: AppEngine
	// MISSING: CloudEndpoints
	// MISSING: ClusterIstio
	// MISSING: MeshIstio
	// MISSING: IstioCanonicalService
	// MISSING: CloudRun
	// MISSING: GkeNamespace
	// MISSING: GkeWorkload
	// MISSING: GkeService
	// MISSING: BasicService
	// MISSING: Telemetry
	// MISSING: UserLabels
	return out
}
func MonitoringServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringServiceSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Custom
	// MISSING: AppEngine
	// MISSING: CloudEndpoints
	// MISSING: ClusterIstio
	// MISSING: MeshIstio
	// MISSING: IstioCanonicalService
	// MISSING: CloudRun
	// MISSING: GkeNamespace
	// MISSING: GkeWorkload
	// MISSING: GkeService
	// MISSING: BasicService
	// MISSING: Telemetry
	// MISSING: UserLabels
	return out
}
func Service_AppEngine_FromProto(mapCtx *direct.MapContext, in *pb.Service_AppEngine) *krm.Service_AppEngine {
	if in == nil {
		return nil
	}
	out := &krm.Service_AppEngine{}
	out.ModuleID = direct.LazyPtr(in.GetModuleId())
	return out
}
func Service_AppEngine_ToProto(mapCtx *direct.MapContext, in *krm.Service_AppEngine) *pb.Service_AppEngine {
	if in == nil {
		return nil
	}
	out := &pb.Service_AppEngine{}
	out.ModuleId = direct.ValueOf(in.ModuleID)
	return out
}
func Service_BasicService_FromProto(mapCtx *direct.MapContext, in *pb.Service_BasicService) *krm.Service_BasicService {
	if in == nil {
		return nil
	}
	out := &krm.Service_BasicService{}
	out.ServiceType = direct.LazyPtr(in.GetServiceType())
	out.ServiceLabels = in.ServiceLabels
	return out
}
func Service_BasicService_ToProto(mapCtx *direct.MapContext, in *krm.Service_BasicService) *pb.Service_BasicService {
	if in == nil {
		return nil
	}
	out := &pb.Service_BasicService{}
	out.ServiceType = direct.ValueOf(in.ServiceType)
	out.ServiceLabels = in.ServiceLabels
	return out
}
func Service_CloudEndpoints_FromProto(mapCtx *direct.MapContext, in *pb.Service_CloudEndpoints) *krm.Service_CloudEndpoints {
	if in == nil {
		return nil
	}
	out := &krm.Service_CloudEndpoints{}
	out.Service = direct.LazyPtr(in.GetService())
	return out
}
func Service_CloudEndpoints_ToProto(mapCtx *direct.MapContext, in *krm.Service_CloudEndpoints) *pb.Service_CloudEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.Service_CloudEndpoints{}
	out.Service = direct.ValueOf(in.Service)
	return out
}
func Service_CloudRun_FromProto(mapCtx *direct.MapContext, in *pb.Service_CloudRun) *krm.Service_CloudRun {
	if in == nil {
		return nil
	}
	out := &krm.Service_CloudRun{}
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func Service_CloudRun_ToProto(mapCtx *direct.MapContext, in *krm.Service_CloudRun) *pb.Service_CloudRun {
	if in == nil {
		return nil
	}
	out := &pb.Service_CloudRun{}
	out.ServiceName = direct.ValueOf(in.ServiceName)
	out.Location = direct.ValueOf(in.Location)
	return out
}
func Service_ClusterIstio_FromProto(mapCtx *direct.MapContext, in *pb.Service_ClusterIstio) *krm.Service_ClusterIstio {
	if in == nil {
		return nil
	}
	out := &krm.Service_ClusterIstio{}
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	out.ServiceNamespace = direct.LazyPtr(in.GetServiceNamespace())
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	return out
}
func Service_ClusterIstio_ToProto(mapCtx *direct.MapContext, in *krm.Service_ClusterIstio) *pb.Service_ClusterIstio {
	if in == nil {
		return nil
	}
	out := &pb.Service_ClusterIstio{}
	out.Location = direct.ValueOf(in.Location)
	out.ClusterName = direct.ValueOf(in.ClusterName)
	out.ServiceNamespace = direct.ValueOf(in.ServiceNamespace)
	out.ServiceName = direct.ValueOf(in.ServiceName)
	return out
}
func Service_Custom_FromProto(mapCtx *direct.MapContext, in *pb.Service_Custom) *krm.Service_Custom {
	if in == nil {
		return nil
	}
	out := &krm.Service_Custom{}
	return out
}
func Service_Custom_ToProto(mapCtx *direct.MapContext, in *krm.Service_Custom) *pb.Service_Custom {
	if in == nil {
		return nil
	}
	out := &pb.Service_Custom{}
	return out
}
func Service_GkeNamespace_FromProto(mapCtx *direct.MapContext, in *pb.Service_GkeNamespace) *krm.Service_GkeNamespace {
	if in == nil {
		return nil
	}
	out := &krm.Service_GkeNamespace{}
	// MISSING: ProjectID
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	out.NamespaceName = direct.LazyPtr(in.GetNamespaceName())
	return out
}
func Service_GkeNamespace_ToProto(mapCtx *direct.MapContext, in *krm.Service_GkeNamespace) *pb.Service_GkeNamespace {
	if in == nil {
		return nil
	}
	out := &pb.Service_GkeNamespace{}
	// MISSING: ProjectID
	out.Location = direct.ValueOf(in.Location)
	out.ClusterName = direct.ValueOf(in.ClusterName)
	out.NamespaceName = direct.ValueOf(in.NamespaceName)
	return out
}
func Service_GkeNamespaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service_GkeNamespace) *krm.Service_GkeNamespaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Service_GkeNamespaceObservedState{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	// MISSING: Location
	// MISSING: ClusterName
	// MISSING: NamespaceName
	return out
}
func Service_GkeNamespaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Service_GkeNamespaceObservedState) *pb.Service_GkeNamespace {
	if in == nil {
		return nil
	}
	out := &pb.Service_GkeNamespace{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	// MISSING: Location
	// MISSING: ClusterName
	// MISSING: NamespaceName
	return out
}
func Service_GkeService_FromProto(mapCtx *direct.MapContext, in *pb.Service_GkeService) *krm.Service_GkeService {
	if in == nil {
		return nil
	}
	out := &krm.Service_GkeService{}
	// MISSING: ProjectID
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	out.NamespaceName = direct.LazyPtr(in.GetNamespaceName())
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	return out
}
func Service_GkeService_ToProto(mapCtx *direct.MapContext, in *krm.Service_GkeService) *pb.Service_GkeService {
	if in == nil {
		return nil
	}
	out := &pb.Service_GkeService{}
	// MISSING: ProjectID
	out.Location = direct.ValueOf(in.Location)
	out.ClusterName = direct.ValueOf(in.ClusterName)
	out.NamespaceName = direct.ValueOf(in.NamespaceName)
	out.ServiceName = direct.ValueOf(in.ServiceName)
	return out
}
func Service_GkeServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service_GkeService) *krm.Service_GkeServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Service_GkeServiceObservedState{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	// MISSING: Location
	// MISSING: ClusterName
	// MISSING: NamespaceName
	// MISSING: ServiceName
	return out
}
func Service_GkeServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Service_GkeServiceObservedState) *pb.Service_GkeService {
	if in == nil {
		return nil
	}
	out := &pb.Service_GkeService{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	// MISSING: Location
	// MISSING: ClusterName
	// MISSING: NamespaceName
	// MISSING: ServiceName
	return out
}
func Service_GkeWorkload_FromProto(mapCtx *direct.MapContext, in *pb.Service_GkeWorkload) *krm.Service_GkeWorkload {
	if in == nil {
		return nil
	}
	out := &krm.Service_GkeWorkload{}
	// MISSING: ProjectID
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ClusterName = direct.LazyPtr(in.GetClusterName())
	out.NamespaceName = direct.LazyPtr(in.GetNamespaceName())
	out.TopLevelControllerType = direct.LazyPtr(in.GetTopLevelControllerType())
	out.TopLevelControllerName = direct.LazyPtr(in.GetTopLevelControllerName())
	return out
}
func Service_GkeWorkload_ToProto(mapCtx *direct.MapContext, in *krm.Service_GkeWorkload) *pb.Service_GkeWorkload {
	if in == nil {
		return nil
	}
	out := &pb.Service_GkeWorkload{}
	// MISSING: ProjectID
	out.Location = direct.ValueOf(in.Location)
	out.ClusterName = direct.ValueOf(in.ClusterName)
	out.NamespaceName = direct.ValueOf(in.NamespaceName)
	out.TopLevelControllerType = direct.ValueOf(in.TopLevelControllerType)
	out.TopLevelControllerName = direct.ValueOf(in.TopLevelControllerName)
	return out
}
func Service_GkeWorkloadObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service_GkeWorkload) *krm.Service_GkeWorkloadObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Service_GkeWorkloadObservedState{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	// MISSING: Location
	// MISSING: ClusterName
	// MISSING: NamespaceName
	// MISSING: TopLevelControllerType
	// MISSING: TopLevelControllerName
	return out
}
func Service_GkeWorkloadObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Service_GkeWorkloadObservedState) *pb.Service_GkeWorkload {
	if in == nil {
		return nil
	}
	out := &pb.Service_GkeWorkload{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	// MISSING: Location
	// MISSING: ClusterName
	// MISSING: NamespaceName
	// MISSING: TopLevelControllerType
	// MISSING: TopLevelControllerName
	return out
}
func Service_IstioCanonicalService_FromProto(mapCtx *direct.MapContext, in *pb.Service_IstioCanonicalService) *krm.Service_IstioCanonicalService {
	if in == nil {
		return nil
	}
	out := &krm.Service_IstioCanonicalService{}
	out.MeshUid = direct.LazyPtr(in.GetMeshUid())
	out.CanonicalServiceNamespace = direct.LazyPtr(in.GetCanonicalServiceNamespace())
	out.CanonicalService = direct.LazyPtr(in.GetCanonicalService())
	return out
}
func Service_IstioCanonicalService_ToProto(mapCtx *direct.MapContext, in *krm.Service_IstioCanonicalService) *pb.Service_IstioCanonicalService {
	if in == nil {
		return nil
	}
	out := &pb.Service_IstioCanonicalService{}
	out.MeshUid = direct.ValueOf(in.MeshUid)
	out.CanonicalServiceNamespace = direct.ValueOf(in.CanonicalServiceNamespace)
	out.CanonicalService = direct.ValueOf(in.CanonicalService)
	return out
}
func Service_MeshIstio_FromProto(mapCtx *direct.MapContext, in *pb.Service_MeshIstio) *krm.Service_MeshIstio {
	if in == nil {
		return nil
	}
	out := &krm.Service_MeshIstio{}
	out.MeshUid = direct.LazyPtr(in.GetMeshUid())
	out.ServiceNamespace = direct.LazyPtr(in.GetServiceNamespace())
	out.ServiceName = direct.LazyPtr(in.GetServiceName())
	return out
}
func Service_MeshIstio_ToProto(mapCtx *direct.MapContext, in *krm.Service_MeshIstio) *pb.Service_MeshIstio {
	if in == nil {
		return nil
	}
	out := &pb.Service_MeshIstio{}
	out.MeshUid = direct.ValueOf(in.MeshUid)
	out.ServiceNamespace = direct.ValueOf(in.ServiceNamespace)
	out.ServiceName = direct.ValueOf(in.ServiceName)
	return out
}
func Service_Telemetry_FromProto(mapCtx *direct.MapContext, in *pb.Service_Telemetry) *krm.Service_Telemetry {
	if in == nil {
		return nil
	}
	out := &krm.Service_Telemetry{}
	out.ResourceName = direct.LazyPtr(in.GetResourceName())
	return out
}
func Service_Telemetry_ToProto(mapCtx *direct.MapContext, in *krm.Service_Telemetry) *pb.Service_Telemetry {
	if in == nil {
		return nil
	}
	out := &pb.Service_Telemetry{}
	out.ResourceName = direct.ValueOf(in.ResourceName)
	return out
}
