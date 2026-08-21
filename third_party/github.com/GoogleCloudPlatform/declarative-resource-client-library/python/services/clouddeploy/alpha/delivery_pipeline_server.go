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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/clouddeploy/alpha/clouddeploy_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/alpha"
)

// DeliveryPipelineServer implements the gRPC interface for DeliveryPipeline.
type DeliveryPipelineServer struct{}

// ProtoToDeliveryPipelineSerialPipeline converts a DeliveryPipelineSerialPipeline object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipeline(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipeline) *alpha.DeliveryPipelineSerialPipeline {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipeline{}
	for _, r := range p.GetStages() {
		obj.Stages = append(obj.Stages, *ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStages(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStages converts a DeliveryPipelineSerialPipelineStages object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStages(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStages) *alpha.DeliveryPipelineSerialPipelineStages {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStages{
		TargetId: dcl.StringOrNil(p.GetTargetId()),
		Strategy: ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategy(p.GetStrategy()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	for _, r := range p.GetDeployParameters() {
		obj.DeployParameters = append(obj.DeployParameters, *ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesDeployParameters(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategy converts a DeliveryPipelineSerialPipelineStagesStrategy object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategy(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategy) *alpha.DeliveryPipelineSerialPipelineStagesStrategy {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategy{
		Standard: ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandard(p.GetStandard()),
		Canary:   ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanary(p.GetCanary()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandard converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandard(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandard) *alpha.DeliveryPipelineSerialPipelineStagesStrategyStandard {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyStandard{
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(p.GetPostdeploy()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) *alpha.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) *alpha.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanary converts a DeliveryPipelineSerialPipelineStagesStrategyCanary object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanary(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanary) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanary {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanary{
		RuntimeConfig:          ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(p.GetRuntimeConfig()),
		CanaryDeployment:       ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(p.GetCanaryDeployment()),
		CustomCanaryDeployment: ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(p.GetCustomCanaryDeployment()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{
		Kubernetes: ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(p.GetKubernetes()),
		CloudRun:   ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(p.GetCloudRun()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{
		GatewayServiceMesh: ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(p.GetGatewayServiceMesh()),
		ServiceNetworking:  ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(p.GetServiceNetworking()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{
		HttpRoute:             dcl.StringOrNil(p.GetHttpRoute()),
		Service:               dcl.StringOrNil(p.GetService()),
		Deployment:            dcl.StringOrNil(p.GetDeployment()),
		RouteUpdateWaitTime:   dcl.StringOrNil(p.GetRouteUpdateWaitTime()),
		StableCutbackDuration: dcl.StringOrNil(p.GetStableCutbackDuration()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{
		Service:                    dcl.StringOrNil(p.GetService()),
		Deployment:                 dcl.StringOrNil(p.GetDeployment()),
		DisablePodOverprovisioning: dcl.Bool(p.GetDisablePodOverprovisioning()),
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{
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
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(p.GetPostdeploy()),
	}
	for _, r := range p.GetPercentages() {
		obj.Percentages = append(obj.Percentages, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	for _, r := range p.GetPhaseConfigs() {
		obj.PhaseConfigs = append(obj.PhaseConfigs, *ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(r))
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{
		PhaseId:    dcl.StringOrNil(p.GetPhaseId()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
		Verify:     dcl.Bool(p.GetVerify()),
		Predeploy:  ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(p.GetPredeploy()),
		Postdeploy: ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(p.GetPostdeploy()),
	}
	for _, r := range p.GetProfiles() {
		obj.Profiles = append(obj.Profiles, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	for _, r := range p.GetActions() {
		obj.Actions = append(obj.Actions, r)
	}
	return obj
}

// ProtoToDeliveryPipelineSerialPipelineStagesDeployParameters converts a DeliveryPipelineSerialPipelineStagesDeployParameters object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineSerialPipelineStagesDeployParameters(p *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesDeployParameters) *alpha.DeliveryPipelineSerialPipelineStagesDeployParameters {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineSerialPipelineStagesDeployParameters{}
	return obj
}

// ProtoToDeliveryPipelineCondition converts a DeliveryPipelineCondition object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineCondition(p *alphapb.ClouddeployAlphaDeliveryPipelineCondition) *alpha.DeliveryPipelineCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineCondition{
		PipelineReadyCondition:  ProtoToClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition(p.GetPipelineReadyCondition()),
		TargetsPresentCondition: ProtoToClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition(p.GetTargetsPresentCondition()),
		TargetsTypeCondition:    ProtoToClouddeployAlphaDeliveryPipelineConditionTargetsTypeCondition(p.GetTargetsTypeCondition()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionPipelineReadyCondition converts a DeliveryPipelineConditionPipelineReadyCondition object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition(p *alphapb.ClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition) *alpha.DeliveryPipelineConditionPipelineReadyCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineConditionPipelineReadyCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsPresentCondition converts a DeliveryPipelineConditionTargetsPresentCondition object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition(p *alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition) *alpha.DeliveryPipelineConditionTargetsPresentCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineConditionTargetsPresentCondition{
		Status:     dcl.Bool(p.GetStatus()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	for _, r := range p.GetMissingTargets() {
		obj.MissingTargets = append(obj.MissingTargets, r)
	}
	return obj
}

// ProtoToDeliveryPipelineConditionTargetsTypeCondition converts a DeliveryPipelineConditionTargetsTypeCondition object from its proto representation.
func ProtoToClouddeployAlphaDeliveryPipelineConditionTargetsTypeCondition(p *alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsTypeCondition) *alpha.DeliveryPipelineConditionTargetsTypeCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.DeliveryPipelineConditionTargetsTypeCondition{
		Status:       dcl.Bool(p.GetStatus()),
		ErrorDetails: dcl.StringOrNil(p.GetErrorDetails()),
	}
	return obj
}

// ProtoToDeliveryPipeline converts a DeliveryPipeline resource from its proto representation.
func ProtoToDeliveryPipeline(p *alphapb.ClouddeployAlphaDeliveryPipeline) *alpha.DeliveryPipeline {
	obj := &alpha.DeliveryPipeline{
		Name:           dcl.StringOrNil(p.GetName()),
		Uid:            dcl.StringOrNil(p.GetUid()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		SerialPipeline: ProtoToClouddeployAlphaDeliveryPipelineSerialPipeline(p.GetSerialPipeline()),
		Condition:      ProtoToClouddeployAlphaDeliveryPipelineCondition(p.GetCondition()),
		Etag:           dcl.StringOrNil(p.GetEtag()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
		Suspended:      dcl.Bool(p.GetSuspended()),
	}
	return obj
}

// DeliveryPipelineSerialPipelineToProto converts a DeliveryPipelineSerialPipeline object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineToProto(o *alpha.DeliveryPipelineSerialPipeline) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipeline {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipeline{}
	sStages := make([]*alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStages, len(o.Stages))
	for i, r := range o.Stages {
		sStages[i] = ClouddeployAlphaDeliveryPipelineSerialPipelineStagesToProto(&r)
	}
	p.SetStages(sStages)
	return p
}

// DeliveryPipelineSerialPipelineStagesToProto converts a DeliveryPipelineSerialPipelineStages object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesToProto(o *alpha.DeliveryPipelineSerialPipelineStages) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStages {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStages{}
	p.SetTargetId(dcl.ValueOrEmptyString(o.TargetId))
	p.SetStrategy(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyToProto(o.Strategy))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	sDeployParameters := make([]*alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesDeployParameters, len(o.DeployParameters))
	for i, r := range o.DeployParameters {
		sDeployParameters[i] = ClouddeployAlphaDeliveryPipelineSerialPipelineStagesDeployParametersToProto(&r)
	}
	p.SetDeployParameters(sDeployParameters)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyToProto converts a DeliveryPipelineSerialPipelineStagesStrategy object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategy) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategy {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategy{}
	p.SetStandard(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o.Standard))
	p.SetCanary(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryToProto(o.Canary))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandard object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyStandard) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandard {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandard{}
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto(o.Postdeploy))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanary object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanary) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanary {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanary{}
	p.SetRuntimeConfig(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto(o.RuntimeConfig))
	p.SetCanaryDeployment(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentToProto(o.CanaryDeployment))
	p.SetCustomCanaryDeployment(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto(o.CustomCanaryDeployment))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
	p.SetKubernetes(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto(o.Kubernetes))
	p.SetCloudRun(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto(o.CloudRun))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
	p.SetGatewayServiceMesh(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto(o.GatewayServiceMesh))
	p.SetServiceNetworking(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto(o.ServiceNetworking))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMeshToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
	p.SetHttpRoute(dcl.ValueOrEmptyString(o.HttpRoute))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetDeployment(dcl.ValueOrEmptyString(o.Deployment))
	p.SetRouteUpdateWaitTime(dcl.ValueOrEmptyString(o.RouteUpdateWaitTime))
	p.SetStableCutbackDuration(dcl.ValueOrEmptyString(o.StableCutbackDuration))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworkingToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetDeployment(dcl.ValueOrEmptyString(o.Deployment))
	p.SetDisablePodOverprovisioning(dcl.ValueOrEmptyBool(o.DisablePodOverprovisioning))
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRunToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
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
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto(o.Postdeploy))
	sPercentages := make([]int64, len(o.Percentages))
	for i, r := range o.Percentages {
		sPercentages[i] = r
	}
	p.SetPercentages(sPercentages)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
	sPhaseConfigs := make([]*alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs, len(o.PhaseConfigs))
	for i, r := range o.PhaseConfigs {
		sPhaseConfigs[i] = ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto(&r)
	}
	p.SetPhaseConfigs(sPhaseConfigs)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{}
	p.SetPhaseId(dcl.ValueOrEmptyString(o.PhaseId))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	p.SetVerify(dcl.ValueOrEmptyBool(o.Verify))
	p.SetPredeploy(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto(o.Predeploy))
	p.SetPostdeploy(ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto(o.Postdeploy))
	sProfiles := make([]string, len(o.Profiles))
	for i, r := range o.Profiles {
		sProfiles[i] = r
	}
	p.SetProfiles(sProfiles)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeployToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto converts a DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeployToProto(o *alpha.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
	sActions := make([]string, len(o.Actions))
	for i, r := range o.Actions {
		sActions[i] = r
	}
	p.SetActions(sActions)
	return p
}

// DeliveryPipelineSerialPipelineStagesDeployParametersToProto converts a DeliveryPipelineSerialPipelineStagesDeployParameters object to its proto representation.
func ClouddeployAlphaDeliveryPipelineSerialPipelineStagesDeployParametersToProto(o *alpha.DeliveryPipelineSerialPipelineStagesDeployParameters) *alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesDeployParameters {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineSerialPipelineStagesDeployParameters{}
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
func ClouddeployAlphaDeliveryPipelineConditionToProto(o *alpha.DeliveryPipelineCondition) *alphapb.ClouddeployAlphaDeliveryPipelineCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineCondition{}
	p.SetPipelineReadyCondition(ClouddeployAlphaDeliveryPipelineConditionPipelineReadyConditionToProto(o.PipelineReadyCondition))
	p.SetTargetsPresentCondition(ClouddeployAlphaDeliveryPipelineConditionTargetsPresentConditionToProto(o.TargetsPresentCondition))
	p.SetTargetsTypeCondition(ClouddeployAlphaDeliveryPipelineConditionTargetsTypeConditionToProto(o.TargetsTypeCondition))
	return p
}

// DeliveryPipelineConditionPipelineReadyConditionToProto converts a DeliveryPipelineConditionPipelineReadyCondition object to its proto representation.
func ClouddeployAlphaDeliveryPipelineConditionPipelineReadyConditionToProto(o *alpha.DeliveryPipelineConditionPipelineReadyCondition) *alphapb.ClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineConditionPipelineReadyCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// DeliveryPipelineConditionTargetsPresentConditionToProto converts a DeliveryPipelineConditionTargetsPresentCondition object to its proto representation.
func ClouddeployAlphaDeliveryPipelineConditionTargetsPresentConditionToProto(o *alpha.DeliveryPipelineConditionTargetsPresentCondition) *alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsPresentCondition{}
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
func ClouddeployAlphaDeliveryPipelineConditionTargetsTypeConditionToProto(o *alpha.DeliveryPipelineConditionTargetsTypeCondition) *alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsTypeCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.ClouddeployAlphaDeliveryPipelineConditionTargetsTypeCondition{}
	p.SetStatus(dcl.ValueOrEmptyBool(o.Status))
	p.SetErrorDetails(dcl.ValueOrEmptyString(o.ErrorDetails))
	return p
}

// DeliveryPipelineToProto converts a DeliveryPipeline resource to its proto representation.
func DeliveryPipelineToProto(resource *alpha.DeliveryPipeline) *alphapb.ClouddeployAlphaDeliveryPipeline {
	p := &alphapb.ClouddeployAlphaDeliveryPipeline{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetSerialPipeline(ClouddeployAlphaDeliveryPipelineSerialPipelineToProto(resource.SerialPipeline))
	p.SetCondition(ClouddeployAlphaDeliveryPipelineConditionToProto(resource.Condition))
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
func (s *DeliveryPipelineServer) applyDeliveryPipeline(ctx context.Context, c *alpha.Client, request *alphapb.ApplyClouddeployAlphaDeliveryPipelineRequest) (*alphapb.ClouddeployAlphaDeliveryPipeline, error) {
	p := ProtoToDeliveryPipeline(request.GetResource())
	res, err := c.ApplyDeliveryPipeline(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DeliveryPipelineToProto(res)
	return r, nil
}

// applyClouddeployAlphaDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Apply() method.
func (s *DeliveryPipelineServer) ApplyClouddeployAlphaDeliveryPipeline(ctx context.Context, request *alphapb.ApplyClouddeployAlphaDeliveryPipelineRequest) (*alphapb.ClouddeployAlphaDeliveryPipeline, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDeliveryPipeline(ctx, cl, request)
}

// DeleteDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipeline Delete() method.
func (s *DeliveryPipelineServer) DeleteClouddeployAlphaDeliveryPipeline(ctx context.Context, request *alphapb.DeleteClouddeployAlphaDeliveryPipelineRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDeliveryPipeline(ctx, ProtoToDeliveryPipeline(request.GetResource()))

}

// ListClouddeployAlphaDeliveryPipeline handles the gRPC request by passing it to the underlying DeliveryPipelineList() method.
func (s *DeliveryPipelineServer) ListClouddeployAlphaDeliveryPipeline(ctx context.Context, request *alphapb.ListClouddeployAlphaDeliveryPipelineRequest) (*alphapb.ListClouddeployAlphaDeliveryPipelineResponse, error) {
	cl, err := createConfigDeliveryPipeline(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDeliveryPipeline(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ClouddeployAlphaDeliveryPipeline
	for _, r := range resources.Items {
		rp := DeliveryPipelineToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListClouddeployAlphaDeliveryPipelineResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDeliveryPipeline(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
