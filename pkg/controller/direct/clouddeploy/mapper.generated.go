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
	cloudbuildv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
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
func DeliveryPipelineSpec_FromProto(mapCtx *direct.MapContext, in *pb.DeliveryPipeline) *krm.DeliveryPipelineSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeliveryPipelineSpec{}
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
func DeliveryPipelineSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeliveryPipelineSpec) *pb.DeliveryPipeline {
	if in == nil {
		return nil
	}
	out := &pb.DeliveryPipeline{}
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
func DeliveryPipelineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeliveryPipeline) *krm.DeliveryPipelineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeliveryPipelineObservedState{}
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
func DeliveryPipelineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeliveryPipelineObservedState) *pb.DeliveryPipeline {
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

func AnthosCluster_FromProto(mapCtx *direct.MapContext, in *pb.AnthosCluster) *krm.AnthosCluster {
	if in == nil {
		return nil
	}
	out := &krm.AnthosCluster{}
	out.Membership = direct.LazyPtr(in.GetMembership())
	return out
}
func AnthosCluster_ToProto(mapCtx *direct.MapContext, in *krm.AnthosCluster) *pb.AnthosCluster {
	if in == nil {
		return nil
	}
	out := &pb.AnthosCluster{}
	out.Membership = direct.ValueOf(in.Membership)
	return out
}
func AssociatedEntities_FromProto(mapCtx *direct.MapContext, in *pb.AssociatedEntities) *krm.AssociatedEntities {
	if in == nil {
		return nil
	}
	out := &krm.AssociatedEntities{}
	out.GkeClusters = direct.Slice_FromProto(mapCtx, in.GkeClusters, GkeCluster_FromProto)
	out.AnthosClusters = direct.Slice_FromProto(mapCtx, in.AnthosClusters, AnthosCluster_FromProto)
	return out
}
func AssociatedEntities_ToProto(mapCtx *direct.MapContext, in *krm.AssociatedEntities) *pb.AssociatedEntities {
	if in == nil {
		return nil
	}
	out := &pb.AssociatedEntities{}
	out.GkeClusters = direct.Slice_ToProto(mapCtx, in.GkeClusters, GkeCluster_ToProto)
	out.AnthosClusters = direct.Slice_ToProto(mapCtx, in.AnthosClusters, AnthosCluster_ToProto)
	return out
}
func CloudRunLocation_FromProto(mapCtx *direct.MapContext, in *pb.CloudRunLocation) *krm.CloudRunLocation {
	if in == nil {
		return nil
	}
	out := &krm.CloudRunLocation{}
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func CloudRunLocation_ToProto(mapCtx *direct.MapContext, in *krm.CloudRunLocation) *pb.CloudRunLocation {
	if in == nil {
		return nil
	}
	out := &pb.CloudRunLocation{}
	out.Location = direct.ValueOf(in.Location)
	return out
}
func CustomTarget_FromProto(mapCtx *direct.MapContext, in *pb.CustomTarget) *krm.CustomTarget {
	if in == nil {
		return nil
	}
	out := &krm.CustomTarget{}
	out.CustomTargetTypeRef.External = in.GetCustomTargetType()
	return out
}
func CustomTarget_ToProto(mapCtx *direct.MapContext, in *krm.CustomTarget) *pb.CustomTarget {
	if in == nil {
		return nil
	}
	out := &pb.CustomTarget{}
	if in.CustomTargetTypeRef != nil {
		out.CustomTargetType = in.CustomTargetTypeRef.External
	}

	return out
}
func DefaultPool_FromProto(mapCtx *direct.MapContext, in *pb.DefaultPool) *krm.DefaultPool {
	if in == nil {
		return nil
	}
	out := &krm.DefaultPool{}
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ArtifactStorage = direct.LazyPtr(in.GetArtifactStorage())
	return out
}
func DefaultPool_ToProto(mapCtx *direct.MapContext, in *krm.DefaultPool) *pb.DefaultPool {
	if in == nil {
		return nil
	}
	out := &pb.DefaultPool{}
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ArtifactStorage = direct.ValueOf(in.ArtifactStorage)
	return out
}

func ExecutionConfig_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionConfig) *krm.ExecutionConfig {
	if in == nil {
		return nil
	}
	out := &krm.ExecutionConfig{}
	out.Usages = direct.EnumSlice_FromProto(mapCtx, in.Usages)
	out.DefaultPool = DefaultPool_FromProto(mapCtx, in.GetDefaultPool())
	out.PrivatePool = PrivatePool_FromProto(mapCtx, in.GetPrivatePool())
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ArtifactStorage = direct.LazyPtr(in.GetArtifactStorage())
	out.ExecutionTimeout = direct.StringDuration_FromProto(mapCtx, in.GetExecutionTimeout())
	out.Verbose = direct.LazyPtr(in.GetVerbose())
	return out
}
func ExecutionConfig_ToProto(mapCtx *direct.MapContext, in *krm.ExecutionConfig) *pb.ExecutionConfig {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionConfig{}
	out.Usages = direct.EnumSlice_ToProto[pb.ExecutionConfig_ExecutionEnvironmentUsage](mapCtx, in.Usages)
	if oneof := DefaultPool_ToProto(mapCtx, in.DefaultPool); oneof != nil {
		out.ExecutionEnvironment = &pb.ExecutionConfig_DefaultPool{DefaultPool: oneof}
	}
	if oneof := PrivatePool_ToProto(mapCtx, in.PrivatePool); oneof != nil {
		out.ExecutionEnvironment = &pb.ExecutionConfig_PrivatePool{PrivatePool: oneof}
	}
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ArtifactStorage = direct.ValueOf(in.ArtifactStorage)
	out.ExecutionTimeout = direct.StringDuration_ToProto(mapCtx, in.ExecutionTimeout)
	out.Verbose = direct.ValueOf(in.Verbose)
	return out
}
func GkeCluster_FromProto(mapCtx *direct.MapContext, in *pb.GkeCluster) *krm.GkeCluster {
	if in == nil {
		return nil
	}
	out := &krm.GkeCluster{}
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ProxyURL = direct.LazyPtr(in.GetProxyUrl())
	out.DNSEndpoint = direct.LazyPtr(in.GetDnsEndpoint())
	return out
}
func GkeCluster_ToProto(mapCtx *direct.MapContext, in *krm.GkeCluster) *pb.GkeCluster {
	if in == nil {
		return nil
	}
	out := &pb.GkeCluster{}
	out.Cluster = direct.ValueOf(in.Cluster)
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.ProxyUrl = direct.ValueOf(in.ProxyURL)
	out.DnsEndpoint = direct.ValueOf(in.DNSEndpoint)
	return out
}
func MultiTarget_FromProto(mapCtx *direct.MapContext, in *pb.MultiTarget) *krm.MultiTarget {
	if in == nil {
		return nil
	}
	out := &krm.MultiTarget{}
	out.TargetIds = in.TargetIds
	return out
}
func MultiTarget_ToProto(mapCtx *direct.MapContext, in *krm.MultiTarget) *pb.MultiTarget {
	if in == nil {
		return nil
	}
	out := &pb.MultiTarget{}
	out.TargetIds = in.TargetIds
	return out
}
func PrivatePool_FromProto(mapCtx *direct.MapContext, in *pb.PrivatePool) *krm.PrivatePool {
	if in == nil {
		return nil
	}
	out := &krm.PrivatePool{}
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.ArtifactStorage = direct.LazyPtr(in.GetArtifactStorage())
	return out
}
func PrivatePool_ToProto(mapCtx *direct.MapContext, in *krm.PrivatePool) *pb.PrivatePool {
	if in == nil {
		return nil
	}
	out := &pb.PrivatePool{}
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.ArtifactStorage = direct.ValueOf(in.ArtifactStorage)
	return out
}
func DeployTargetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.DeployTargetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployTargetSpec{}
	// out.Name = direct.LazyPtr(in.GetName())
	// MISSING: TargetID
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	//out.Annotations = in.Annotations
	//out.Labels = in.Labels
	out.RequireApproval = direct.LazyPtr(in.GetRequireApproval())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Gke = GkeCluster_FromProto(mapCtx, in.GetGke())
	out.AnthosCluster = AnthosCluster_FromProto(mapCtx, in.GetAnthosCluster())
	out.Run = CloudRunLocation_FromProto(mapCtx, in.GetRun())
	out.MultiTarget = MultiTarget_FromProto(mapCtx, in.GetMultiTarget())
	out.CustomTarget = CustomTarget_FromProto(mapCtx, in.GetCustomTarget())
	// MISSING: AssociatedEntities
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ExecutionConfigs = direct.Slice_FromProto(mapCtx, in.ExecutionConfigs, ExecutionConfig_FromProto)
	out.DeployParameters = in.DeployParameters
	return out
}
func DeployTargetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployTargetSpec) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	// out.Name = direct.ValueOf(in.Name)
	// MISSING: TargetID
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	// out.Annotations = in.Annotations
	// out.Labels = in.Labels
	out.RequireApproval = direct.ValueOf(in.RequireApproval)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	if oneof := GkeCluster_ToProto(mapCtx, in.Gke); oneof != nil {
		out.DeploymentTarget = &pb.Target_Gke{Gke: oneof}
	}
	if oneof := AnthosCluster_ToProto(mapCtx, in.AnthosCluster); oneof != nil {
		out.DeploymentTarget = &pb.Target_AnthosCluster{AnthosCluster: oneof}
	}
	if oneof := CloudRunLocation_ToProto(mapCtx, in.Run); oneof != nil {
		out.DeploymentTarget = &pb.Target_Run{Run: oneof}
	}
	if oneof := MultiTarget_ToProto(mapCtx, in.MultiTarget); oneof != nil {
		out.DeploymentTarget = &pb.Target_MultiTarget{MultiTarget: oneof}
	}
	if oneof := CustomTarget_ToProto(mapCtx, in.CustomTarget); oneof != nil {
		out.DeploymentTarget = &pb.Target_CustomTarget{CustomTarget: oneof}
	}
	// MISSING: AssociatedEntities
	out.Etag = direct.ValueOf(in.Etag)
	out.ExecutionConfigs = direct.Slice_ToProto(mapCtx, in.ExecutionConfigs, ExecutionConfig_ToProto)
	out.DeployParameters = in.DeployParameters
	return out
}
func DeployTargetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Target) *krm.DeployTargetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployTargetObservedState{}
	// MISSING: Name
	out.TargetID = direct.LazyPtr(in.GetTargetId())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: RequireApproval
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Gke
	// MISSING: AnthosCluster
	// MISSING: Run
	// MISSING: MultiTarget
	// MISSING: CustomTarget
	// MISSING: AssociatedEntities
	// MISSING: Etag
	// MISSING: ExecutionConfigs
	// MISSING: DeployParameters
	return out
}
func DeployTargetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployTargetObservedState) *pb.Target {
	if in == nil {
		return nil
	}
	out := &pb.Target{}
	// MISSING: Name
	out.TargetId = direct.ValueOf(in.TargetID)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	// MISSING: RequireApproval
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Gke
	// MISSING: AnthosCluster
	// MISSING: Run
	// MISSING: MultiTarget
	// MISSING: CustomTarget
	// MISSING: AssociatedEntities
	// MISSING: Etag
	// MISSING: ExecutionConfigs
	// MISSING: DeployParameters
	return out
}

func CustomTargetSkaffoldActions_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetSkaffoldActions) *krm.CustomTargetSkaffoldActions {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetSkaffoldActions{}
	out.RenderAction = direct.LazyPtr(in.GetRenderAction())
	out.DeployAction = direct.LazyPtr(in.GetDeployAction())
	out.IncludeSkaffoldModules = direct.Slice_FromProto(mapCtx, in.IncludeSkaffoldModules, SkaffoldModules_FromProto)
	return out
}
func CustomTargetSkaffoldActions_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetSkaffoldActions) *pb.CustomTargetSkaffoldActions {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetSkaffoldActions{}
	out.RenderAction = direct.ValueOf(in.RenderAction)
	out.DeployAction = direct.ValueOf(in.DeployAction)
	out.IncludeSkaffoldModules = direct.Slice_ToProto(mapCtx, in.IncludeSkaffoldModules, SkaffoldModules_ToProto)
	return out
}
func CustomTargetType_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetType) *krm.CustomTargetTypeSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetTypeSpec{}
	//out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	out.Description = direct.LazyPtr(in.GetDescription())
	//out.Annotations = in.Annotations
	//out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	//out.Etag = direct.LazyPtr(in.GetEtag())
	out.CustomActions = CustomTargetSkaffoldActions_FromProto(mapCtx, in.GetCustomActions())
	return out
}
func CustomTargetType_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetTypeSpec) *pb.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetType{}
	// out.Name = direct.ValueOf(in.Name)
	// MISSING: CustomTargetTypeID
	// MISSING: Uid
	out.Description = direct.ValueOf(in.Description)
	// out.Annotations = in.Annotations
	// out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// out.Etag = direct.ValueOf(in.Etag)
	if oneof := CustomTargetSkaffoldActions_ToProto(mapCtx, in.CustomActions); oneof != nil {
		out.Definition = &pb.CustomTargetType_CustomActions{CustomActions: oneof}
	}
	return out
}
func CustomTargetTypeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CustomTargetType) *krm.CustomTargetTypeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CustomTargetTypeObservedState{}
	// MISSING: Name
	out.CustomTargetTypeID = direct.LazyPtr(in.GetCustomTargetTypeId())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}
func CustomTargetTypeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CustomTargetTypeObservedState) *pb.CustomTargetType {
	if in == nil {
		return nil
	}
	out := &pb.CustomTargetType{}
	// MISSING: Name
	out.CustomTargetTypeId = direct.ValueOf(in.CustomTargetTypeID)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Description
	// MISSING: Annotations
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	// MISSING: CustomActions
	return out
}

func SkaffoldModules_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules) *krm.SkaffoldModules {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules{}
	out.Configs = in.Configs
	out.Git = SkaffoldModules_SkaffoldGitSource_FromProto(mapCtx, in.GetGit())
	out.GoogleCloudStorage = SkaffoldModules_SkaffoldGCSSource_FromProto(mapCtx, in.GetGoogleCloudStorage())
	out.GoogleCloudBuildRepo = SkaffoldModules_SkaffoldGcbRepoSource_FromProto(mapCtx, in.GetGoogleCloudBuildRepo())
	return out
}
func SkaffoldModules_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules) *pb.SkaffoldModules {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules{}
	out.Configs = in.Configs
	if oneof := SkaffoldModules_SkaffoldGitSource_ToProto(mapCtx, in.Git); oneof != nil {
		out.Source = &pb.SkaffoldModules_Git{Git: oneof}
	}
	if oneof := SkaffoldModules_SkaffoldGCSSource_ToProto(mapCtx, in.GoogleCloudStorage); oneof != nil {
		out.Source = &pb.SkaffoldModules_GoogleCloudStorage{GoogleCloudStorage: oneof}
	}
	if oneof := SkaffoldModules_SkaffoldGcbRepoSource_ToProto(mapCtx, in.GoogleCloudBuildRepo); oneof != nil {
		out.Source = &pb.SkaffoldModules_GoogleCloudBuildRepo{GoogleCloudBuildRepo: oneof}
	}
	return out
}
func SkaffoldModules_SkaffoldGCSSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGCSSource) *krm.SkaffoldModules_SkaffoldGCSSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGCSSource{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func SkaffoldModules_SkaffoldGCSSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGCSSource) *pb.SkaffoldModules_SkaffoldGCSSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGCSSource{}
	out.Source = direct.ValueOf(in.Source)
	out.Path = direct.ValueOf(in.Path)
	return out
}
func SkaffoldModules_SkaffoldGcbRepoSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGCBRepoSource) *krm.SkaffoldModules_SkaffoldGcbRepoSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGcbRepoSource{}
	if in.GetRepository() != "" {
		out.RepositoryRef = &cloudbuildv1alpha1.RepositoryRef{
			External: in.GetRepository(),
		}
	}

	out.Path = direct.LazyPtr(in.GetPath())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}
func SkaffoldModules_SkaffoldGcbRepoSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGcbRepoSource) *pb.SkaffoldModules_SkaffoldGCBRepoSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGCBRepoSource{}
	if in.RepositoryRef != nil {
		out.Repository = in.RepositoryRef.External
	}
	out.Path = direct.ValueOf(in.Path)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}
func SkaffoldModules_SkaffoldGitSource_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldModules_SkaffoldGitSource) *krm.SkaffoldModules_SkaffoldGitSource {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldModules_SkaffoldGitSource{}
	out.Repo = direct.LazyPtr(in.GetRepo())
	out.Path = direct.LazyPtr(in.GetPath())
	out.Ref = direct.LazyPtr(in.GetRef())
	return out
}
func SkaffoldModules_SkaffoldGitSource_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldModules_SkaffoldGitSource) *pb.SkaffoldModules_SkaffoldGitSource {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldModules_SkaffoldGitSource{}
	out.Repo = direct.ValueOf(in.Repo)
	out.Path = direct.ValueOf(in.Path)
	out.Ref = direct.ValueOf(in.Ref)
	return out
}
