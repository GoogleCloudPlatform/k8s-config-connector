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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/clouddeploy/beta/clouddeploy_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/beta"
)

// DeliveryPipelineServer implements the gRPC interface for DeliveryPipeline.
type DeliveryPipelineServer struct{}

// ProtoToDeliveryPipelineSerialPipeline converts a DeliveryPipelineSerialPipeline object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipeline(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipeline) *beta.DeliveryPipelineSerialPipeline {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipeline{}
	for _, r := range p.GetStages() {
		obj.Stages = append(obj.Stages, *ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStages(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStages converts a DeliveryPipelineSerialPipelineStages object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStages(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStages) *beta.DeliveryPipelineSerialPipelineStages {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStages{
		TargetId: dcl.StringOrNil(p.GetTargetId()),
		Strategy: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy(p.GetStrategy()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	for _, r := range p.GetDeployParameters() {
		obj.DeployParameters = append(obj.DeployParameters, *ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesDeployParameters(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategy converts a DeliveryPipelineSerialPipelineStagesStrategy object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy) *beta.DeliveryPipelineSerialPipelineStagesStrategy {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategy{
		Standard: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard(p.GetStandard()),
		Canary:   ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanary(p.GetCanary()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandard converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard) *beta.DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyStandard{
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(p.GetPostdeploy()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) *beta.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) *beta.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanary converts a DeliveryPipelineSerialPipelineStagesStrategyCanary object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanary(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanary) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanary {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanary{
		RuntimeConfig:          ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(p.GetRuntimeConfig()),
		CanaryDeployment:       ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(p.GetCanaryDeployment()),
		CustomCanaryDeployment: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(p.GetCustomCanaryDeployment()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{
		Kubernetes: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(p.GetKubernetes()),
		CloudRun:   ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(p.GetCloudRun()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{
		GatewayServiceMesh: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(p.GetGatewayServiceMesh()),
		ServiceNetworking:  ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(p.GetServiceNetworking()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{
		HttpRoute:             dcl.StringOrNil(p.GetHttpRoute()),
		Service:               dcl.StringOrNil(p.GetService()),
		Deployment:            dcl.StringOrNil(p.GetDeployment()),
		RouteUpdateWaitTime:   dcl.StringOrNil(p.GetRouteUpdateWaitTime()),
		StableCutbackDuration: dcl.StringOrNil(p.GetStableCutbackDuration()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{
		Service:                    dcl.StringOrNil(p.GetService()),
		Deployment:                 dcl.StringOrNil(p.GetDeployment()),
		DisablePodOverprovisioning: dcl.Bool(p.GetDisablePodOverprovisioning()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{
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
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(p.GetPostdeploy()),
	}
	for _, r := range p.GetPercentages() {
		obj.Percentages = append(obj.Percentages, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	for _, r := range p.GetPhaseConfigs() {
		obj.PhaseConfigs = append(obj.PhaseConfigs, *ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{
		PhaseId:    dcl.StringOrNil(p.GetPhaseId()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(p.GetPostdeploy()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesDeployParameters converts a DeliveryPipelineSerialPipelineStagesDeployParameters object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineSerialPipelineStagesDeployParameters(p *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesDeployParameters) *beta.DeliveryPipelineSerialPipelineStagesDeployParameters {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineSerialPipelineStagesDeployParameters{}
	return obj
}

// ProtoToDeliveryPipelineCondition converts a DeliveryPipelineCondition object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineCondition(p *betapb.ClouddeployBetaDeliveryPipelineCondition) *beta.DeliveryPipelineCondition {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineCondition{
		PipelineReadyCondition:  ProtoToClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition(p.GetPipelineReadyCondition()),
		TargetsPresentCondition: ProtoToClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition(p.GetTargetsPresentCondition()),
		TargetsTypeCondition:    ProtoToClouddeployBetaDeliveryPipelineConditionTargetsTypeCondition(p.GetTargetsTypeCondition()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionPipelineReadyCondition converts a DeliveryPipelineConditionPipelineReadyCondition object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition(p *betapb.ClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition) *beta.DeliveryPipelineConditionPipelineReadyCondition {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineConditionPipelineReadyCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsPresentCondition converts a DeliveryPipelineConditionTargetsPresentCondition object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition(p *betapb.ClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition) *beta.DeliveryPipelineConditionTargetsPresentCondition {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineConditionTargetsPresentCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetMissingTargets() {
		obj.MissingTargets = append(obj.MissingTargets, r)
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsTypeCondition converts a DeliveryPipelineConditionTargetsTypeCondition object from its proto representation.
func ProtoToClouddeployBetaDeliveryPipelineConditionTargetsTypeCondition(p *betapb.ClouddeployBetaDeliveryPipelineConditionTargetsTypeCondition) *beta.DeliveryPipelineConditionTargetsTypeCondition {
	if p == nil {
		return nil
	}
	obj := &beta.DeliveryPipelineConditionTargetsTypeCondition{
		Status:       dcl.Bool(p.GetStatus()),
		ErrorDetails: dcl.StringOrNil(p.GetErrorDetails()),
	}
	return obj
}

// ProtoToDeliveryPipeline converts a DeliveryPipeline resource from its proto representation.
func ProtoToDeliveryPipeline(p *betapb.ClouddeployBetaDeliveryPipeline) *beta.DeliveryPipeline {
	obj := &beta.DeliveryPipeline{
		Name:           dcl.StringOrNil(p.GetName()),
		Uid:            dcl.StringOrNil(p.GetUid()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		SerialPipeline: ProtoToClouddeployBetaDeliveryPipelineSerialPipeline(p.GetSerialPipeline()),
		Condition:      ProtoToClouddeployBetaDeliveryPipelineCondition(p.GetCondition()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
		Suspended:      dcl.Bool(p.GetSuspended()),
	}
	return obj
}

// DeliveryPipelineSerialPipelineToProto converts a DeliveryPipelineSerialPipeline object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineToProto(o *beta.DeliveryPipelineSerialPipeline) *betapb.ClouddeployBetaDeliveryPipelineSerialPipeline {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipeline{}
	sStages := make([]*betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStages, len(o.Stages))
	for i, r := range o.Stages {
		sStages[i] = ClouddeployBetaDeliveryPipelineSerialPipelineStagesToProto(&r)
	}
	p.SetStages(sStages)
	return p
}

// DeliveryPipelineSerialPipelineStagesToProto converts a DeliveryPipelineSerialPipelineStages object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesToProto(o *beta.DeliveryPipelineSerialPipelineStages) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStages {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStages{}
	p.SetTargetId(dcl.ValueOrEmptyString(o.TargetId))
	p.SetStrategy(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyToProto(o.Strategy))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	sDeployParameters := make([]*betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesDeployParameters, len(o.DeployParameters))
	for i, r := range o.DeployParameters {
		sDeployParameters[i] = ClouddeployBetaDeliveryPipelineSerialPipelineStagesDeployParametersToProto(&r)
	}
	p.SetDeployParameters(sDeployParameters)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyToProto converts a DeliveryPipelineSerialPipelineStagesStrategy object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategy) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategy{}
	p.SetStandard(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o.Standard))
	p.SetCanary(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryToProto(o.Canary))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyStandard) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandard{}
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto(o.Postdeploy))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanary object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanary) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanary {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanary{}
	p.SetRuntimeConfig(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto(o.RuntimeConfig))
	p.SetCanaryDeployment(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentToProto(o.CanaryDeployment))
	p.SetCustomCanaryDeployment(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto(o.CustomCanaryDeployment))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	p.SetKubernetes(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto(o.Kubernetes))
	p.SetCloudRun(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto(o.CloudRun))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	p.SetGatewayServiceMesh(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto(o.GatewayServiceMesh))
	p.SetServiceNetworking(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto(o.ServiceNetworking))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	p.SetHttpRoute(dcl.ValueOrEmptyString(o.HttpRoute))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetDeployment(dcl.ValueOrEmptyString(o.Deployment))
	p.SetRouteUpdateWaitTime(dcl.ValueOrEmptyString(o.RouteUpdateWaitTime))
	p.SetStableCutbackDuration(dcl.ValueOrEmptyString(o.StableCutbackDuration))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetDeployment(dcl.ValueOrEmptyString(o.Deployment))
	p.SetDisablePodOverprovisioning(dcl.ValueOrEmptyBool(o.DisablePodOverprovisioning))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
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
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto(o.Postdeploy))
	sPercentages := make([]int64, len(o.Percentages))
	for i, r := range o.Percentages {
		sPercentages[i] = r
	}
	p.SetPercentages(sPercentages)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	sPhaseConfigs := make([]*betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, len(o.PhaseConfigs))
	for i, r := range o.PhaseConfigs {
		sPhaseConfigs[i] = ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto(&r)
	}
	p.SetPhaseConfigs(sPhaseConfigs)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}
	p.SetPhaseId(dcl.ValueOrEmptyString(o.PhaseId))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto(o.Postdeploy))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto(o *beta.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesDeployParametersToProto converts a DeliveryPipelineSerialPipelineStagesDeployParameters object to its proto representation.
func ClouddeployBetaDeliveryPipelineSerialPipelineStagesDeployParametersToProto(o *beta.DeliveryPipelineSerialPipelineStagesDeployParameters) *betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesDeployParameters {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineSerialPipelineStagesDeployParameters{}
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
func ClouddeployBetaDeliveryPipelineConditionToProto(o *beta.DeliveryPipelineCondition) *betapb.ClouddeployBetaDeliveryPipelineCondition {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineCondition{}
	p.SetPipelineReadyCondition(ClouddeployBetaDeliveryPipelineConditionPipelineReadyConditionToProto(o.PipelineReadyCondition))
	p.SetTargetsPresentCondition(ClouddeployBetaDeliveryPipelineConditionTargetsPresentConditionToProto(o.TargetsPresentCondition))
	p.SetTargetsTypeCondition(ClouddeployBetaDeliveryPipelineConditionTargetsTypeConditionToProto(o.TargetsTypeCondition))
	return p
}

// DeliveryPipelineConditionPipelineReadyConditionToProto converts a DeliveryPipelineConditionPipelineReadyCondition object to its proto representation.
func ClouddeployBetaDeliveryPipelineConditionPipelineReadyConditionToProto(o *beta.DeliveryPipelineConditionPipelineReadyCondition) *betapb.ClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineConditionPipelineReadyCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// DeliveryPipelineConditionTargetsPresentConditionToProto converts a DeliveryPipelineConditionTargetsPresentCondition object to its proto representation.
func ClouddeployBetaDeliveryPipelineConditionTargetsPresentConditionToProto(o *beta.DeliveryPipelineConditionTargetsPresentCondition) *betapb.ClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineConditionTargetsPresentCondition{}
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
func ClouddeployBetaDeliveryPipelineConditionTargetsTypeConditionToProto(o *beta.DeliveryPipelineConditionTargetsTypeCondition) *betapb.ClouddeployBetaDeliveryPipelineConditionTargetsTypeCondition {
	if o == nil {
		return nil
	}
	p := &betapb.ClouddeployBetaDeliveryPipelineConditionTargetsTypeCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetErrorDetails(dcl.ValueOrEmptyString(o.ErrorDetails))
	return p
}

// DeliveryPipelineToProto converts a DeliveryPipeline resource to its proto representation.
func DeliveryPipelineToProto(resource *beta.DeliveryPipeline) *betapb.ClouddeployBetaDeliveryPipeline {
	p := &betapb.ClouddeployBetaDeliveryPipeline{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetSerialPipeline(ClouddeployBetaDeliveryPipelineSerialPipelineToProto(resource.SerialPipeline))
	p.SetCondition(ClouddeployBetaDeliveryPipelineConditionToProto(resource.Condition))
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
func (s *DeliveryPipelineServer) applyDeliveryPipeline(ctx context.Context, c *beta.Client, request *betapb.ApplyClouddeployBetaDeliveryPipelineRequest) (*betapb.ClouddeployBetaDeliveryPipeline, error) {
	p := ProtoToDeliveryPipeline(request.GetResource())
	res, err := c.ApplyDeliveryPipeline(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DeliveryPipelineToProto(res)
	return r, nil
}

// applyClouddeployBetaDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) ApplyClouddeployBetaDeliveryPipeline(ctx context.Context, request *betapb.ApplyClouddeployBetaDeliveryPipelineRequest) (*betapb.ClouddeployBetaDeliveryPipeline, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDeliveryPipeline(ctx, cl, request)
}

// DeleteDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Delete() method.
func (s *DeliveryPipelineServer) DeleteClouddeployBetaDeliveryPipeline(ctx context.Context, request *betapb.DeleteClouddeployBetaDeliveryPipelineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDeliveryPipeline(ctx, ProtoToDeliveryPipeline(request.GetResource()))

}

// ListClouddeployBetaDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipelineList() method.
func (s *DeliveryPipelineServer) ListClouddeployBetaDeliveryPipeline(ctx context.Context, request *betapb.ListClouddeployBetaDeliveryPipelineRequest) (*betapb.ListClouddeployBetaDeliveryPipelineResponse, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDeliveryPipeline(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ClouddeployBetaDeliveryPipeline
	for _, r := range resources.Items {
		rp := DeliveryPipelineToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListClouddeployBetaDeliveryPipelineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDeliveryPipeline(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
