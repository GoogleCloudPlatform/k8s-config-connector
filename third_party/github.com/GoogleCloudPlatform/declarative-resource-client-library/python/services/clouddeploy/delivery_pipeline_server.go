// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	clouddeploypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/clouddeploy/clouddeploy_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy"
)

// DeliveryPipelineServer implements the gRPC interface for DeliveryPipeline.
type DeliveryPipelineServer struct{}

// ProtoToDeliveryPipelineSerialPipeline converts a DeliveryPipelineSerialPipeline object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipeline(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipeline) *clouddeploy.DeliveryPipelineSerialPipeline {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipeline{}
	for _, r := range p.GetStages() {
		obj.Stages = append(obj.Stages, *ProtoToClouddeployDeliveryPipelineSerialPipelineStages(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStages converts a DeliveryPipelineSerialPipelineStages object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStages(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStages) *clouddeploy.DeliveryPipelineSerialPipelineStages {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStages{
		TargetId: dcl.StringOrNil(p.GetTargetId()),
		Strategy: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategy(p.GetStrategy()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	for _, r := range p.GetDeployParameters() {
		obj.DeployParameters = append(obj.DeployParameters, *ProtoToClouddeployDeliveryPipelineSerialPipelineStagesDeployParameters(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategy converts a DeliveryPipelineSerialPipelineStagesStrategy object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategy(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategy) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategy {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategy{
		Standard: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard(p.GetStandard()),
		Canary:   ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary(p.GetCanary()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandard converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandard{
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(p.GetPostdeploy()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanary converts a DeliveryPipelineSerialPipelineStagesStrategyCanary object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanary {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanary{
		RuntimeConfig:          ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(p.GetRuntimeConfig()),
		CanaryDeployment:       ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(p.GetCanaryDeployment()),
		CustomCanaryDeployment: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(p.GetCustomCanaryDeployment()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{
		Kubernetes: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(p.GetKubernetes()),
		CloudRun:   ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(p.GetCloudRun()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{
		GatewayServiceMesh: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(p.GetGatewayServiceMesh()),
		ServiceNetworking:  ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(p.GetServiceNetworking()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{
		HttpRoute:             dcl.StringOrNil(p.GetHttpRoute()),
		Service:               dcl.StringOrNil(p.GetService()),
		Deployment:            dcl.StringOrNil(p.GetDeployment()),
		RouteUpdateWaitTime:   dcl.StringOrNil(p.GetRouteUpdateWaitTime()),
		StableCutbackDuration: dcl.StringOrNil(p.GetStableCutbackDuration()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{
		Service:                    dcl.StringOrNil(p.GetService()),
		Deployment:                 dcl.StringOrNil(p.GetDeployment()),
		DisablePodOverprovisioning: dcl.Bool(p.GetDisablePodOverprovisioning()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{
		AutomaticTrafficControl: dcl.Bool(p.GetAutomaticTrafficControl()),
	}
	for _, r := range p.GetCanaryRevisionTags() {
		obj.CanaryRevisionTags = append(obj.CanaryRevisionTags, r)
	}
	for _, r := range p.GetPriorRevisionTags() {
		obj.PriorRevisionTags = append(obj.PriorRevisionTags, r)
	}
	for _, r := range p.GetStableRevisionTags() {
		obj.StableRevisionTags = append(obj.StableRevisionTags, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(p.GetPostdeploy()),
	}
	for _, r := range p.GetPercentages() {
		obj.Percentages = append(obj.Percentages, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	for _, r := range p.GetPhaseConfigs() {
		obj.PhaseConfigs = append(obj.PhaseConfigs, *ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{
		PhaseId:    dcl.StringOrNil(p.GetPhaseId()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(p.GetPostdeploy()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesDeployParameters converts a DeliveryPipelineSerialPipelineStagesDeployParameters object from its proto representation.
func ProtoToClouddeployDeliveryPipelineSerialPipelineStagesDeployParameters(p *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesDeployParameters) *clouddeploy.DeliveryPipelineSerialPipelineStagesDeployParameters {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineSerialPipelineStagesDeployParameters{}
	return obj
}

// ProtoToDeliveryPipelineCondition converts a DeliveryPipelineCondition object from its proto representation.
func ProtoToClouddeployDeliveryPipelineCondition(p *clouddeploypb.ClouddeployDeliveryPipelineCondition) *clouddeploy.DeliveryPipelineCondition {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineCondition{
		PipelineReadyCondition:  ProtoToClouddeployDeliveryPipelineConditionPipelineReadyCondition(p.GetPipelineReadyCondition()),
		TargetsPresentCondition: ProtoToClouddeployDeliveryPipelineConditionTargetsPresentCondition(p.GetTargetsPresentCondition()),
		TargetsTypeCondition:    ProtoToClouddeployDeliveryPipelineConditionTargetsTypeCondition(p.GetTargetsTypeCondition()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionPipelineReadyCondition converts a DeliveryPipelineConditionPipelineReadyCondition object from its proto representation.
func ProtoToClouddeployDeliveryPipelineConditionPipelineReadyCondition(p *clouddeploypb.ClouddeployDeliveryPipelineConditionPipelineReadyCondition) *clouddeploy.DeliveryPipelineConditionPipelineReadyCondition {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineConditionPipelineReadyCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsPresentCondition converts a DeliveryPipelineConditionTargetsPresentCondition object from its proto representation.
func ProtoToClouddeployDeliveryPipelineConditionTargetsPresentCondition(p *clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsPresentCondition) *clouddeploy.DeliveryPipelineConditionTargetsPresentCondition {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineConditionTargetsPresentCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetMissingTargets() {
		obj.MissingTargets = append(obj.MissingTargets, r)
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsTypeCondition converts a DeliveryPipelineConditionTargetsTypeCondition object from its proto representation.
func ProtoToClouddeployDeliveryPipelineConditionTargetsTypeCondition(p *clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsTypeCondition) *clouddeploy.DeliveryPipelineConditionTargetsTypeCondition {
	if p == nil {
		return nil
	}
	obj := &clouddeploy.DeliveryPipelineConditionTargetsTypeCondition{
		Status:       dcl.Bool(p.GetStatus()),
		ErrorDetails: dcl.StringOrNil(p.GetErrorDetails()),
	}
	return obj
}

// ProtoToDeliveryPipeline converts a DeliveryPipeline resource from its proto representation.
func ProtoToDeliveryPipeline(p *clouddeploypb.ClouddeployDeliveryPipeline) *clouddeploy.DeliveryPipeline {
	obj := &clouddeploy.DeliveryPipeline{
		Name:           dcl.StringOrNil(p.GetName()),
		Uid:            dcl.StringOrNil(p.GetUid()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		SerialPipeline: ProtoToClouddeployDeliveryPipelineSerialPipeline(p.GetSerialPipeline()),
		Condition:      ProtoToClouddeployDeliveryPipelineCondition(p.GetCondition()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
		Suspended:      dcl.Bool(p.GetSuspended()),
	}
	return obj
}

// DeliveryPipelineSerialPipelineToProto converts a DeliveryPipelineSerialPipeline object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineToProto(o *clouddeploy.DeliveryPipelineSerialPipeline) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipeline {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipeline{}
	sStages := make([]*clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStages, len(o.Stages))
	for i, r := range o.Stages {
		sStages[i] = ClouddeployDeliveryPipelineSerialPipelineStagesToProto(&r)
	}
	p.SetStages(sStages)
	return p
}

// DeliveryPipelineSerialPipelineStagesToProto converts a DeliveryPipelineSerialPipelineStages object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStages) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStages {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStages{}
	p.SetTargetId(dcl.ValueOrEmptyString(o.TargetId))
	p.SetStrategy(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyToProto(o.Strategy))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	sDeployParameters := make([]*clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesDeployParameters, len(o.DeployParameters))
	for i, r := range o.DeployParameters {
		sDeployParameters[i] = ClouddeployDeliveryPipelineSerialPipelineStagesDeployParametersToProto(&r)
	}
	p.SetDeployParameters(sDeployParameters)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyToProto converts a DeliveryPipelineSerialPipelineStagesStrategy object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategy) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategy {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategy{}
	p.SetStandard(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o.Standard))
	p.SetCanary(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryToProto(o.Canary))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandard) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard{}
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto(o.Postdeploy))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanary object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanary) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary{}
	p.SetRuntimeConfig(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto(o.RuntimeConfig))
	p.SetCanaryDeployment(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentToProto(o.CanaryDeployment))
	p.SetCustomCanaryDeployment(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto(o.CustomCanaryDeployment))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	p.SetKubernetes(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto(o.Kubernetes))
	p.SetCloudRun(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto(o.CloudRun))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	p.SetGatewayServiceMesh(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto(o.GatewayServiceMesh))
	p.SetServiceNetworking(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto(o.ServiceNetworking))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	p.SetHttpRoute(dcl.ValueOrEmptyString(o.HttpRoute))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetDeployment(dcl.ValueOrEmptyString(o.Deployment))
	p.SetRouteUpdateWaitTime(dcl.ValueOrEmptyString(o.RouteUpdateWaitTime))
	p.SetStableCutbackDuration(dcl.ValueOrEmptyString(o.StableCutbackDuration))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetDeployment(dcl.ValueOrEmptyString(o.Deployment))
	p.SetDisablePodOverprovisioning(dcl.ValueOrEmptyBool(o.DisablePodOverprovisioning))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
	p.SetAutomaticTrafficControl(dcl.ValueOrEmptyBool(o.AutomaticTrafficControl))
	sCanaryRevisionTags := make([]string, len(o.CanaryRevisionTags))
	for i, r := range o.CanaryRevisionTags {
		sCanaryRevisionTags[i] = r
	}
	p.SetCanaryRevisionTags(sCanaryRevisionTags)
	sPriorRevisionTags := make([]string, len(o.PriorRevisionTags))
	for i, r := range o.PriorRevisionTags {
		sPriorRevisionTags[i] = r
	}
	p.SetPriorRevisionTags(sPriorRevisionTags)
	sStableRevisionTags := make([]string, len(o.StableRevisionTags))
	for i, r := range o.StableRevisionTags {
		sStableRevisionTags[i] = r
	}
	p.SetStableRevisionTags(sStableRevisionTags)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto(o.Postdeploy))
	sPercentages := make([]int64, len(o.Percentages))
	for i, r := range o.Percentages {
		sPercentages[i] = r
	}
	p.SetPercentages(sPercentages)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	sPhaseConfigs := make([]*clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, len(o.PhaseConfigs))
	for i, r := range o.PhaseConfigs {
		sPhaseConfigs[i] = ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto(&r)
	}
	p.SetPhaseConfigs(sPhaseConfigs)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}
	p.SetPhaseId(dcl.ValueOrEmptyString(o.PhaseId))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto(o.Postdeploy))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesDeployParametersToProto converts a DeliveryPipelineSerialPipelineStagesDeployParameters object to its proto representation.
func ClouddeployDeliveryPipelineSerialPipelineStagesDeployParametersToProto(o *clouddeploy.DeliveryPipelineSerialPipelineStagesDeployParameters) *clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesDeployParameters {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineSerialPipelineStagesDeployParameters{}
	mValues := make(map[string]string, len(o.Values))
	for k, r := range o.Values {
		mValues[k] = r
	}
	p.SetValues(mValues)
	mMatchTargetLabels := make(map[string]string, len(o.MatchTargetLabels))
	for k, r := range o.MatchTargetLabels {
		mMatchTargetLabels[k] = r
	}
	p.SetMatchTargetLabels(mMatchTargetLabels)
	return p
}

// DeliveryPipelineConditionToProto converts a DeliveryPipelineCondition object to its proto representation.
func ClouddeployDeliveryPipelineConditionToProto(o *clouddeploy.DeliveryPipelineCondition) *clouddeploypb.ClouddeployDeliveryPipelineCondition {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineCondition{}
	p.SetPipelineReadyCondition(ClouddeployDeliveryPipelineConditionPipelineReadyConditionToProto(o.PipelineReadyCondition))
	p.SetTargetsPresentCondition(ClouddeployDeliveryPipelineConditionTargetsPresentConditionToProto(o.TargetsPresentCondition))
	p.SetTargetsTypeCondition(ClouddeployDeliveryPipelineConditionTargetsTypeConditionToProto(o.TargetsTypeCondition))
	return p
}

// DeliveryPipelineConditionPipelineReadyConditionToProto converts a DeliveryPipelineConditionPipelineReadyCondition object to its proto representation.
func ClouddeployDeliveryPipelineConditionPipelineReadyConditionToProto(o *clouddeploy.DeliveryPipelineConditionPipelineReadyCondition) *clouddeploypb.ClouddeployDeliveryPipelineConditionPipelineReadyCondition {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineConditionPipelineReadyCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// DeliveryPipelineConditionTargetsPresentConditionToProto converts a DeliveryPipelineConditionTargetsPresentCondition object to its proto representation.
func ClouddeployDeliveryPipelineConditionTargetsPresentConditionToProto(o *clouddeploy.DeliveryPipelineConditionTargetsPresentCondition) *clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsPresentCondition {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsPresentCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	sMissingTargets := make([]string, len(o.MissingTargets))
	for i, r := range o.MissingTargets {
		sMissingTargets[i] = r
	}
	p.SetMissingTargets(sMissingTargets)
	return p
}

// DeliveryPipelineConditionTargetsTypeConditionToProto converts a DeliveryPipelineConditionTargetsTypeCondition object to its proto representation.
func ClouddeployDeliveryPipelineConditionTargetsTypeConditionToProto(o *clouddeploy.DeliveryPipelineConditionTargetsTypeCondition) *clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsTypeCondition {
	if o == nil {
		return nil
	}
	p := &clouddeploypb.ClouddeployDeliveryPipelineConditionTargetsTypeCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetErrorDetails(dcl.ValueOrEmptyString(o.ErrorDetails))
	return p
}

// DeliveryPipelineToProto converts a DeliveryPipeline resource to its proto representation.
func DeliveryPipelineToProto(resource *clouddeploy.DeliveryPipeline) *clouddeploypb.ClouddeployDeliveryPipeline {
	p := &clouddeploypb.ClouddeployDeliveryPipeline{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetSerialPipeline(ClouddeployDeliveryPipelineSerialPipelineToProto(resource.SerialPipeline))
	p.SetCondition(ClouddeployDeliveryPipelineConditionToProto(resource.Condition))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSuspended(dcl.ValueOrEmptyBool(resource.Suspended))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) applyDeliveryPipeline(ctx context.Context, c *clouddeploy.Client, request *clouddeploypb.ApplyClouddeployDeliveryPipelineRequest) (*clouddeploypb.ClouddeployDeliveryPipeline, error) {
	p := ProtoToDeliveryPipeline(request.GetResource())
	res, err := c.ApplyDeliveryPipeline(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DeliveryPipelineToProto(res)
	return r, nil
}

// applyClouddeployDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) ApplyClouddeployDeliveryPipeline(ctx context.Context, request *clouddeploypb.ApplyClouddeployDeliveryPipelineRequest) (*clouddeploypb.ClouddeployDeliveryPipeline, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDeliveryPipeline(ctx, cl, request)
}

// DeleteDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Delete() method.
func (s *DeliveryPipelineServer) DeleteClouddeployDeliveryPipeline(ctx context.Context, request *clouddeploypb.DeleteClouddeployDeliveryPipelineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDeliveryPipeline(ctx, ProtoToDeliveryPipeline(request.GetResource()))

}

// ListClouddeployDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipelineList() method.
func (s *DeliveryPipelineServer) ListClouddeployDeliveryPipeline(ctx context.Context, request *clouddeploypb.ListClouddeployDeliveryPipelineRequest) (*clouddeploypb.ListClouddeployDeliveryPipelineResponse, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDeliveryPipeline(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*clouddeploypb.ClouddeployDeliveryPipeline
	for _, r := range resources.Items {
		rp := DeliveryPipelineToProto(r)
		protos = append(protos, rp)
	}
	p := &clouddeploypb.ListClouddeployDeliveryPipelineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDeliveryPipeline(ctx context.Context, service_account_file string) (*clouddeploy.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return clouddeploy.NewClient(conf), nil
}
