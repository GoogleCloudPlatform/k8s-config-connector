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

package clouddeploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Canary_FromProto(mapCtx *direct.MapContext, in *pb.Canary) *krm.Canary {
	if in == nil {
		return nil
	}
	out := &krm.Canary{}
	out.RuntimeConfig = RuntimeConfig_FromProto(mapCtx, in.GetRuntimeConfig())
	out.CanaryDeployment = CanaryDeployment_FromProto(mapCtx, in.GetCanaryDeployment())
	out.CustomCanaryDeployment = CustomCanaryDeployment_FromProto(mapCtx, in.GetCustomCanaryDeployment())
	return out
}
func Canary_ToProto(mapCtx *direct.MapContext, in *krm.Canary) *pb.Canary {
	if in == nil {
		return nil
	}
	out := &pb.Canary{}
	out.RuntimeConfig = RuntimeConfig_ToProto(mapCtx, in.RuntimeConfig)
	if oneof := CanaryDeployment_ToProto(mapCtx, in.CanaryDeployment); oneof != nil {
		out.Mode = &pb.Canary_CanaryDeployment{CanaryDeployment: oneof}
	}
	if oneof := CustomCanaryDeployment_ToProto(mapCtx, in.CustomCanaryDeployment); oneof != nil {
		out.Mode = &pb.Canary_CustomCanaryDeployment{CustomCanaryDeployment: oneof}
	}
	return out
}
func CanaryDeployment_FromProto(mapCtx *direct.MapContext, in *pb.CanaryDeployment) *krm.CanaryDeployment {
	if in == nil {
		return nil
	}
	out := &krm.CanaryDeployment{}
	out.Percentages = in.Percentages
	out.Verify = direct.LazyPtr(in.GetVerify())
	out.Predeploy = Predeploy_FromProto(mapCtx, in.GetPredeploy())
	out.Postdeploy = Postdeploy_FromProto(mapCtx, in.GetPostdeploy())
	return out
}
func CanaryDeployment_ToProto(mapCtx *direct.MapContext, in *krm.CanaryDeployment) *pb.CanaryDeployment {
	if in == nil {
		return nil
	}
	out := &pb.CanaryDeployment{}
	out.Percentages = in.Percentages
	out.Verify = direct.ValueOf(in.Verify)
	out.Predeploy = Predeploy_ToProto(mapCtx, in.Predeploy)
	out.Postdeploy = Postdeploy_ToProto(mapCtx, in.Postdeploy)
	return out
}
func CloudRunConfig_FromProto(mapCtx *direct.MapContext, in *pb.CloudRunConfig) *krm.CloudRunConfig {
	if in == nil {
		return nil
	}
	out := &krm.CloudRunConfig{}
	out.AutomaticTrafficControl = direct.LazyPtr(in.GetAutomaticTrafficControl())
	out.CanaryRevisionTags = in.CanaryRevisionTags
	out.PriorRevisionTags = in.PriorRevisionTags
	out.StableRevisionTags = in.StableRevisionTags
	return out
}
func CloudRunConfig_ToProto(mapCtx *direct.MapContext, in *krm.CloudRunConfig) *pb.CloudRunConfig {
	if in == nil {
		return nil
	}
	out := &pb.CloudRunConfig{}
	out.AutomaticTrafficControl = direct.ValueOf(in.AutomaticTrafficControl)
	out.CanaryRevisionTags = in.CanaryRevisionTags
	out.PriorRevisionTags = in.PriorRevisionTags
	out.StableRevisionTags = in.StableRevisionTags
	return out
}
func CustomCanaryDeployment_FromProto(mapCtx *direct.MapContext, in *pb.CustomCanaryDeployment) *krm.CustomCanaryDeployment {
	if in == nil {
		return nil
	}
	out := &krm.CustomCanaryDeployment{}
	out.PhaseConfigs = direct.Slice_FromProto(mapCtx, in.PhaseConfigs, CustomCanaryDeployment_PhaseConfig_FromProto)
	return out
}
func CustomCanaryDeployment_ToProto(mapCtx *direct.MapContext, in *krm.CustomCanaryDeployment) *pb.CustomCanaryDeployment {
	if in == nil {
		return nil
	}
	out := &pb.CustomCanaryDeployment{}
	out.PhaseConfigs = direct.Slice_ToProto(mapCtx, in.PhaseConfigs, CustomCanaryDeployment_PhaseConfig_ToProto)
	return out
}
func CustomCanaryDeployment_PhaseConfig_FromProto(mapCtx *direct.MapContext, in *pb.CustomCanaryDeployment_PhaseConfig) *krm.CustomCanaryDeployment_PhaseConfig {
	if in == nil {
		return nil
	}
	out := &krm.CustomCanaryDeployment_PhaseConfig{}
	out.PhaseID = direct.LazyPtr(in.GetPhaseId())
	out.Percentage = direct.LazyPtr(in.GetPercentage())
	out.Profiles = in.Profiles
	out.Verify = direct.LazyPtr(in.GetVerify())
	out.Predeploy = Predeploy_FromProto(mapCtx, in.GetPredeploy())
	out.Postdeploy = Postdeploy_FromProto(mapCtx, in.GetPostdeploy())
	return out
}
func CustomCanaryDeployment_PhaseConfig_ToProto(mapCtx *direct.MapContext, in *krm.CustomCanaryDeployment_PhaseConfig) *pb.CustomCanaryDeployment_PhaseConfig {
	if in == nil {
		return nil
	}
	out := &pb.CustomCanaryDeployment_PhaseConfig{}
	out.PhaseId = direct.ValueOf(in.PhaseID)
	out.Percentage = direct.ValueOf(in.Percentage)
	out.Profiles = in.Profiles
	out.Verify = direct.ValueOf(in.Verify)
	out.Predeploy = Predeploy_ToProto(mapCtx, in.Predeploy)
	out.Postdeploy = Postdeploy_ToProto(mapCtx, in.Postdeploy)
	return out
}
func DeployDeliveryPipelineSpec_FromProto(mapCtx *direct.MapContext, in *pb.DeliveryPipeline) *krm.DeployDeliveryPipelineSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployDeliveryPipelineSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	// ObservedState: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Annotations = in.Annotations
	//out.Labels = in.Labels // Not yet
	// ObservedState: CreateTime
	// ObservedState: UpdateTime
	out.SerialPipeline = SerialPipeline_FromProto(mapCtx, in.GetSerialPipeline())
	// ObservedState: Condition
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Suspended = direct.LazyPtr(in.GetSuspended())
	return out
}
func DeployDeliveryPipelineSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployDeliveryPipelineSpec) *pb.DeliveryPipeline {
	if in == nil {
		return nil
	}
	out := &pb.DeliveryPipeline{}
	out.Name = direct.ValueOf(in.Name)
	// ObservedState: Uid
	out.Description = direct.ValueOf(in.Description)
	out.Annotations = in.Annotations
	//out.Labels = in.Labels // Not yet
	// ObservedState: CreateTime
	// ObservedState: UpdateTime
	if oneof := SerialPipeline_ToProto(mapCtx, in.SerialPipeline); oneof != nil {
		out.Pipeline = &pb.DeliveryPipeline_SerialPipeline{SerialPipeline: oneof}
	}
	// ObservedState: Condition
	out.Etag = direct.ValueOf(in.Etag)
	out.Suspended = direct.ValueOf(in.Suspended)
	return out
}
func DeployDeliveryPipelineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeliveryPipeline) *krm.DeployDeliveryPipelineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployDeliveryPipelineObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: SerialPipeline
	out.Condition = PipelineCondition_FromProto(mapCtx, in.GetCondition())
	// MISSING: Etag
	// MISSING: Suspended
	return out
}
func DeployDeliveryPipelineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployDeliveryPipelineObservedState) *pb.DeliveryPipeline {
	if in == nil {
		return nil
	}
	out := &pb.DeliveryPipeline{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: SerialPipeline
	out.Condition = PipelineCondition_ToProto(mapCtx, in.Condition)
	// MISSING: Etag
	// MISSING: Suspended
	return out
}

func DeployParameters_FromProto(mapCtx *direct.MapContext, in *pb.DeployParameters) *krm.DeployParameters {
	if in == nil {
		return nil
	}
	out := &krm.DeployParameters{}
	out.Values = in.Values
	out.MatchTargetLabels = in.MatchTargetLabels
	return out
}
func DeployParameters_ToProto(mapCtx *direct.MapContext, in *krm.DeployParameters) *pb.DeployParameters {
	if in == nil {
		return nil
	}
	out := &pb.DeployParameters{}
	out.Values = in.Values
	out.MatchTargetLabels = in.MatchTargetLabels
	return out
}
func KubernetesConfig_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesConfig) *krm.KubernetesConfig {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesConfig{}
	out.GatewayServiceMesh = KubernetesConfig_GatewayServiceMesh_FromProto(mapCtx, in.GetGatewayServiceMesh())
	out.ServiceNetworking = KubernetesConfig_ServiceNetworking_FromProto(mapCtx, in.GetServiceNetworking())
	return out
}
func KubernetesConfig_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesConfig) *pb.KubernetesConfig {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesConfig{}
	if oneof := KubernetesConfig_GatewayServiceMesh_ToProto(mapCtx, in.GatewayServiceMesh); oneof != nil {
		out.ServiceDefinition = &pb.KubernetesConfig_GatewayServiceMesh_{GatewayServiceMesh: oneof}
	}
	if oneof := KubernetesConfig_ServiceNetworking_ToProto(mapCtx, in.ServiceNetworking); oneof != nil {
		out.ServiceDefinition = &pb.KubernetesConfig_ServiceNetworking_{ServiceNetworking: oneof}
	}
	return out
}
func KubernetesConfig_GatewayServiceMesh_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesConfig_GatewayServiceMesh) *krm.KubernetesConfig_GatewayServiceMesh {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesConfig_GatewayServiceMesh{}
	out.HTTPRoute = direct.LazyPtr(in.GetHttpRoute())
	out.Service = direct.LazyPtr(in.GetService())
	out.Deployment = direct.LazyPtr(in.GetDeployment())
	out.RouteUpdateWaitTime = direct.StringDuration_FromProto(mapCtx, in.GetRouteUpdateWaitTime())
	out.StableCutbackDuration = direct.StringDuration_FromProto(mapCtx, in.GetStableCutbackDuration())
	out.PodSelectorLabel = direct.LazyPtr(in.GetPodSelectorLabel())
	out.RouteDestinations = KubernetesConfig_GatewayServiceMesh_RouteDestinations_FromProto(mapCtx, in.GetRouteDestinations())
	return out
}
func KubernetesConfig_GatewayServiceMesh_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesConfig_GatewayServiceMesh) *pb.KubernetesConfig_GatewayServiceMesh {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesConfig_GatewayServiceMesh{}
	out.HttpRoute = direct.ValueOf(in.HTTPRoute)
	out.Service = direct.ValueOf(in.Service)
	out.Deployment = direct.ValueOf(in.Deployment)
	out.RouteUpdateWaitTime = direct.StringDuration_ToProto(mapCtx, in.RouteUpdateWaitTime)
	out.StableCutbackDuration = direct.StringDuration_ToProto(mapCtx, in.StableCutbackDuration)
	out.PodSelectorLabel = direct.ValueOf(in.PodSelectorLabel)
	out.RouteDestinations = KubernetesConfig_GatewayServiceMesh_RouteDestinations_ToProto(mapCtx, in.RouteDestinations)
	return out
}
func KubernetesConfig_GatewayServiceMesh_RouteDestinations_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesConfig_GatewayServiceMesh_RouteDestinations) *krm.KubernetesConfig_GatewayServiceMesh_RouteDestinations {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesConfig_GatewayServiceMesh_RouteDestinations{}
	out.DestinationIDs = in.DestinationIds
	out.PropagateService = direct.LazyPtr(in.GetPropagateService())
	return out
}
func KubernetesConfig_GatewayServiceMesh_RouteDestinations_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesConfig_GatewayServiceMesh_RouteDestinations) *pb.KubernetesConfig_GatewayServiceMesh_RouteDestinations {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesConfig_GatewayServiceMesh_RouteDestinations{}
	out.DestinationIds = in.DestinationIDs
	out.PropagateService = direct.ValueOf(in.PropagateService)
	return out
}
func KubernetesConfig_ServiceNetworking_FromProto(mapCtx *direct.MapContext, in *pb.KubernetesConfig_ServiceNetworking) *krm.KubernetesConfig_ServiceNetworking {
	if in == nil {
		return nil
	}
	out := &krm.KubernetesConfig_ServiceNetworking{}
	out.Service = direct.LazyPtr(in.GetService())
	out.Deployment = direct.LazyPtr(in.GetDeployment())
	out.DisablePodOverprovisioning = direct.LazyPtr(in.GetDisablePodOverprovisioning())
	out.PodSelectorLabel = direct.LazyPtr(in.GetPodSelectorLabel())
	return out
}
func KubernetesConfig_ServiceNetworking_ToProto(mapCtx *direct.MapContext, in *krm.KubernetesConfig_ServiceNetworking) *pb.KubernetesConfig_ServiceNetworking {
	if in == nil {
		return nil
	}
	out := &pb.KubernetesConfig_ServiceNetworking{}
	out.Service = direct.ValueOf(in.Service)
	out.Deployment = direct.ValueOf(in.Deployment)
	out.DisablePodOverprovisioning = direct.ValueOf(in.DisablePodOverprovisioning)
	out.PodSelectorLabel = direct.ValueOf(in.PodSelectorLabel)
	return out
}
func PipelineCondition_FromProto(mapCtx *direct.MapContext, in *pb.PipelineCondition) *krm.PipelineCondition {
	if in == nil {
		return nil
	}
	out := &krm.PipelineCondition{}
	out.PipelineReadyCondition = PipelineReadyCondition_FromProto(mapCtx, in.GetPipelineReadyCondition())
	out.TargetsPresentCondition = TargetsPresentCondition_FromProto(mapCtx, in.GetTargetsPresentCondition())
	out.TargetsTypeCondition = TargetsTypeCondition_FromProto(mapCtx, in.GetTargetsTypeCondition())
	return out
}
func PipelineCondition_ToProto(mapCtx *direct.MapContext, in *krm.PipelineCondition) *pb.PipelineCondition {
	if in == nil {
		return nil
	}
	out := &pb.PipelineCondition{}
	out.PipelineReadyCondition = PipelineReadyCondition_ToProto(mapCtx, in.PipelineReadyCondition)
	out.TargetsPresentCondition = TargetsPresentCondition_ToProto(mapCtx, in.TargetsPresentCondition)
	out.TargetsTypeCondition = TargetsTypeCondition_ToProto(mapCtx, in.TargetsTypeCondition)
	return out
}
func PipelineReadyCondition_FromProto(mapCtx *direct.MapContext, in *pb.PipelineReadyCondition) *krm.PipelineReadyCondition {
	if in == nil {
		return nil
	}
	out := &krm.PipelineReadyCondition{}
	out.Status = direct.LazyPtr(in.GetStatus())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func PipelineReadyCondition_ToProto(mapCtx *direct.MapContext, in *krm.PipelineReadyCondition) *pb.PipelineReadyCondition {
	if in == nil {
		return nil
	}
	out := &pb.PipelineReadyCondition{}
	out.Status = direct.ValueOf(in.Status)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Postdeploy_FromProto(mapCtx *direct.MapContext, in *pb.Postdeploy) *krm.Postdeploy {
	if in == nil {
		return nil
	}
	out := &krm.Postdeploy{}
	out.Actions = in.Actions
	return out
}
func Postdeploy_ToProto(mapCtx *direct.MapContext, in *krm.Postdeploy) *pb.Postdeploy {
	if in == nil {
		return nil
	}
	out := &pb.Postdeploy{}
	out.Actions = in.Actions
	return out
}
func Predeploy_FromProto(mapCtx *direct.MapContext, in *pb.Predeploy) *krm.Predeploy {
	if in == nil {
		return nil
	}
	out := &krm.Predeploy{}
	out.Actions = in.Actions
	return out
}
func Predeploy_ToProto(mapCtx *direct.MapContext, in *krm.Predeploy) *pb.Predeploy {
	if in == nil {
		return nil
	}
	out := &pb.Predeploy{}
	out.Actions = in.Actions
	return out
}
func RuntimeConfig_FromProto(mapCtx *direct.MapContext, in *pb.RuntimeConfig) *krm.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.RuntimeConfig{}
	out.Kubernetes = KubernetesConfig_FromProto(mapCtx, in.GetKubernetes())
	out.CloudRun = CloudRunConfig_FromProto(mapCtx, in.GetCloudRun())
	return out
}
func RuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.RuntimeConfig) *pb.RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.RuntimeConfig{}
	if oneof := KubernetesConfig_ToProto(mapCtx, in.Kubernetes); oneof != nil {
		out.RuntimeConfig = &pb.RuntimeConfig_Kubernetes{Kubernetes: oneof}
	}
	if oneof := CloudRunConfig_ToProto(mapCtx, in.CloudRun); oneof != nil {
		out.RuntimeConfig = &pb.RuntimeConfig_CloudRun{CloudRun: oneof}
	}
	return out
}
func SerialPipeline_FromProto(mapCtx *direct.MapContext, in *pb.SerialPipeline) *krm.SerialPipeline {
	if in == nil {
		return nil
	}
	out := &krm.SerialPipeline{}
	out.Stages = direct.Slice_FromProto(mapCtx, in.Stages, Stage_FromProto)
	return out
}
func SerialPipeline_ToProto(mapCtx *direct.MapContext, in *krm.SerialPipeline) *pb.SerialPipeline {
	if in == nil {
		return nil
	}
	out := &pb.SerialPipeline{}
	out.Stages = direct.Slice_ToProto(mapCtx, in.Stages, Stage_ToProto)
	return out
}
func Stage_FromProto(mapCtx *direct.MapContext, in *pb.Stage) *krm.Stage {
	if in == nil {
		return nil
	}
	out := &krm.Stage{}
	out.TargetID = direct.LazyPtr(in.GetTargetId())
	out.Profiles = in.Profiles
	out.Strategy = Strategy_FromProto(mapCtx, in.GetStrategy())
	out.DeployParameters = direct.Slice_FromProto(mapCtx, in.DeployParameters, DeployParameters_FromProto)
	return out
}
func Stage_ToProto(mapCtx *direct.MapContext, in *krm.Stage) *pb.Stage {
	if in == nil {
		return nil
	}
	out := &pb.Stage{}
	out.TargetId = direct.ValueOf(in.TargetID)
	out.Profiles = in.Profiles
	out.Strategy = Strategy_ToProto(mapCtx, in.Strategy)
	out.DeployParameters = direct.Slice_ToProto(mapCtx, in.DeployParameters, DeployParameters_ToProto)
	return out
}
func Standard_FromProto(mapCtx *direct.MapContext, in *pb.Standard) *krm.Standard {
	if in == nil {
		return nil
	}
	out := &krm.Standard{}
	out.Verify = direct.LazyPtr(in.GetVerify())
	out.Predeploy = Predeploy_FromProto(mapCtx, in.GetPredeploy())
	out.Postdeploy = Postdeploy_FromProto(mapCtx, in.GetPostdeploy())
	return out
}
func Standard_ToProto(mapCtx *direct.MapContext, in *krm.Standard) *pb.Standard {
	if in == nil {
		return nil
	}
	out := &pb.Standard{}
	out.Verify = direct.ValueOf(in.Verify)
	out.Predeploy = Predeploy_ToProto(mapCtx, in.Predeploy)
	out.Postdeploy = Postdeploy_ToProto(mapCtx, in.Postdeploy)
	return out
}
func Strategy_FromProto(mapCtx *direct.MapContext, in *pb.Strategy) *krm.Strategy {
	if in == nil {
		return nil
	}
	out := &krm.Strategy{}
	out.Standard = Standard_FromProto(mapCtx, in.GetStandard())
	out.Canary = Canary_FromProto(mapCtx, in.GetCanary())
	return out
}
func Strategy_ToProto(mapCtx *direct.MapContext, in *krm.Strategy) *pb.Strategy {
	if in == nil {
		return nil
	}
	out := &pb.Strategy{}
	if oneof := Standard_ToProto(mapCtx, in.Standard); oneof != nil {
		out.DeploymentStrategy = &pb.Strategy_Standard{Standard: oneof}
	}
	if oneof := Canary_ToProto(mapCtx, in.Canary); oneof != nil {
		out.DeploymentStrategy = &pb.Strategy_Canary{Canary: oneof}
	}
	return out
}
func TargetsPresentCondition_FromProto(mapCtx *direct.MapContext, in *pb.TargetsPresentCondition) *krm.TargetsPresentCondition {
	if in == nil {
		return nil
	}
	out := &krm.TargetsPresentCondition{}
	out.Status = direct.LazyPtr(in.GetStatus())
	out.MissingTargets = in.MissingTargets
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func TargetsPresentCondition_ToProto(mapCtx *direct.MapContext, in *krm.TargetsPresentCondition) *pb.TargetsPresentCondition {
	if in == nil {
		return nil
	}
	out := &pb.TargetsPresentCondition{}
	out.Status = direct.ValueOf(in.Status)
	out.MissingTargets = in.MissingTargets
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func TargetsTypeCondition_FromProto(mapCtx *direct.MapContext, in *pb.TargetsTypeCondition) *krm.TargetsTypeCondition {
	if in == nil {
		return nil
	}
	out := &krm.TargetsTypeCondition{}
	out.Status = direct.LazyPtr(in.GetStatus())
	out.ErrorDetails = direct.LazyPtr(in.GetErrorDetails())
	return out
}
func TargetsTypeCondition_ToProto(mapCtx *direct.MapContext, in *krm.TargetsTypeCondition) *pb.TargetsTypeCondition {
	if in == nil {
		return nil
	}
	out := &pb.TargetsTypeCondition{}
	out.Status = direct.ValueOf(in.Status)
	out.ErrorDetails = direct.ValueOf(in.ErrorDetails)
	return out
}
