// Copyright 2022 Google LLC. All Rights Reserved.
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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/beta/vertex_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex/beta"
)

// EndpointServer implements the gRPC interface for Endpoint.
type EndpointServer struct{}

// ProtoToEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum converts a EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum enum from its proto representation.
func ProtoToVertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(e betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) *beta.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum_name[int32(e)]; ok {
		e := beta.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(n[len("VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointDeployedModels converts a EndpointDeployedModels object from its proto representation.
func ProtoToVertexBetaEndpointDeployedModels(p *betapb.VertexBetaEndpointDeployedModels) *beta.EndpointDeployedModels {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointDeployedModels{
		DedicatedResources:      ProtoToVertexBetaEndpointDeployedModelsDedicatedResources(p.GetDedicatedResources()),
		AutomaticResources:      ProtoToVertexBetaEndpointDeployedModelsAutomaticResources(p.GetAutomaticResources()),
		Id:                      dcl.StringOrNil(p.GetId()),
		Model:                   dcl.StringOrNil(p.GetModel()),
		ModelVersionId:          dcl.StringOrNil(p.GetModelVersionId()),
		DisplayName:             dcl.StringOrNil(p.GetDisplayName()),
		CreateTime:              dcl.StringOrNil(p.GetCreateTime()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
		DisableContainerLogging: dcl.Bool(p.GetDisableContainerLogging()),
		EnableAccessLogging:     dcl.Bool(p.GetEnableAccessLogging()),
		PrivateEndpoints:        ProtoToVertexBetaEndpointDeployedModelsPrivateEndpoints(p.GetPrivateEndpoints()),
		SharedResources:         dcl.StringOrNil(p.GetSharedResources()),
		EnableContainerLogging:  dcl.Bool(p.GetEnableContainerLogging()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResources converts a EndpointDeployedModelsDedicatedResources object from its proto representation.
func ProtoToVertexBetaEndpointDeployedModelsDedicatedResources(p *betapb.VertexBetaEndpointDeployedModelsDedicatedResources) *beta.EndpointDeployedModelsDedicatedResources {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointDeployedModelsDedicatedResources{
		MachineSpec:     ProtoToVertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpec(p.GetMachineSpec()),
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	for _, r := range p.GetAutoscalingMetricSpecs() {
		obj.AutoscalingMetricSpecs = append(obj.AutoscalingMetricSpecs, *ProtoToVertexBetaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(r))
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResourcesMachineSpec converts a EndpointDeployedModelsDedicatedResourcesMachineSpec object from its proto representation.
func ProtoToVertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpec(p *betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpec) *beta.EndpointDeployedModelsDedicatedResourcesMachineSpec {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointDeployedModelsDedicatedResourcesMachineSpec{
		MachineType:      dcl.StringOrNil(p.GetMachineType()),
		AcceleratorType:  ProtoToVertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs converts a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs object from its proto representation.
func ProtoToVertexBetaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(p *betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *beta.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{
		MetricName: dcl.StringOrNil(p.GetMetricName()),
		Target:     dcl.Int64OrNil(p.GetTarget()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsAutomaticResources converts a EndpointDeployedModelsAutomaticResources object from its proto representation.
func ProtoToVertexBetaEndpointDeployedModelsAutomaticResources(p *betapb.VertexBetaEndpointDeployedModelsAutomaticResources) *beta.EndpointDeployedModelsAutomaticResources {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointDeployedModelsAutomaticResources{
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsPrivateEndpoints converts a EndpointDeployedModelsPrivateEndpoints object from its proto representation.
func ProtoToVertexBetaEndpointDeployedModelsPrivateEndpoints(p *betapb.VertexBetaEndpointDeployedModelsPrivateEndpoints) *beta.EndpointDeployedModelsPrivateEndpoints {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointDeployedModelsPrivateEndpoints{
		PredictHttpUri:    dcl.StringOrNil(p.GetPredictHttpUri()),
		ExplainHttpUri:    dcl.StringOrNil(p.GetExplainHttpUri()),
		HealthHttpUri:     dcl.StringOrNil(p.GetHealthHttpUri()),
		ServiceAttachment: dcl.StringOrNil(p.GetServiceAttachment()),
	}
	return obj
}

// ProtoToEndpointEncryptionSpec converts a EndpointEncryptionSpec object from its proto representation.
func ProtoToVertexBetaEndpointEncryptionSpec(p *betapb.VertexBetaEndpointEncryptionSpec) *beta.EndpointEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToEndpoint converts a Endpoint resource from its proto representation.
func ProtoToEndpoint(p *betapb.VertexBetaEndpoint) *beta.Endpoint {
	obj := &beta.Endpoint{
		Name:                         dcl.StringOrNil(p.GetName()),
		DisplayName:                  dcl.StringOrNil(p.GetDisplayName()),
		Description:                  dcl.StringOrNil(p.GetDescription()),
		Etag:                         dcl.StringOrNil(p.GetEtag()),
		CreateTime:                   dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		EncryptionSpec:               ProtoToVertexBetaEndpointEncryptionSpec(p.GetEncryptionSpec()),
		Network:                      dcl.StringOrNil(p.GetNetwork()),
		ModelDeploymentMonitoringJob: dcl.StringOrNil(p.GetModelDeploymentMonitoringJob()),
		Project:                      dcl.StringOrNil(p.GetProject()),
		Location:                     dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexBetaEndpointDeployedModels(r))
	}
	return obj
}

// EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto converts a EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum enum to its proto representation.
func VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto(e *beta.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	if e == nil {
		return betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(0)
	}
	if v, ok := betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum_value["EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"+string(*e)]; ok {
		return betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(v)
	}
	return betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(0)
}

// EndpointDeployedModelsToProto converts a EndpointDeployedModels object to its proto representation.
func VertexBetaEndpointDeployedModelsToProto(o *beta.EndpointDeployedModels) *betapb.VertexBetaEndpointDeployedModels {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaEndpointDeployedModels{}
	p.SetDedicatedResources(VertexBetaEndpointDeployedModelsDedicatedResourcesToProto(o.DedicatedResources))
	p.SetAutomaticResources(VertexBetaEndpointDeployedModelsAutomaticResourcesToProto(o.AutomaticResources))
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	p.SetModelVersionId(dcl.ValueOrEmptyString(o.ModelVersionId))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetDisableContainerLogging(dcl.ValueOrEmptyBool(o.DisableContainerLogging))
	p.SetEnableAccessLogging(dcl.ValueOrEmptyBool(o.EnableAccessLogging))
	p.SetPrivateEndpoints(VertexBetaEndpointDeployedModelsPrivateEndpointsToProto(o.PrivateEndpoints))
	p.SetSharedResources(dcl.ValueOrEmptyString(o.SharedResources))
	p.SetEnableContainerLogging(dcl.ValueOrEmptyBool(o.EnableContainerLogging))
	return p
}

// EndpointDeployedModelsDedicatedResourcesToProto converts a EndpointDeployedModelsDedicatedResources object to its proto representation.
func VertexBetaEndpointDeployedModelsDedicatedResourcesToProto(o *beta.EndpointDeployedModelsDedicatedResources) *betapb.VertexBetaEndpointDeployedModelsDedicatedResources {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaEndpointDeployedModelsDedicatedResources{}
	p.SetMachineSpec(VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecToProto(o.MachineSpec))
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	sAutoscalingMetricSpecs := make([]*betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, len(o.AutoscalingMetricSpecs))
	for i, r := range o.AutoscalingMetricSpecs {
		sAutoscalingMetricSpecs[i] = VertexBetaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto(&r)
	}
	p.SetAutoscalingMetricSpecs(sAutoscalingMetricSpecs)
	return p
}

// EndpointDeployedModelsDedicatedResourcesMachineSpecToProto converts a EndpointDeployedModelsDedicatedResourcesMachineSpec object to its proto representation.
func VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecToProto(o *beta.EndpointDeployedModelsDedicatedResourcesMachineSpec) *betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpec {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpec{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetAcceleratorType(VertexBetaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto converts a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs object to its proto representation.
func VertexBetaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto(o *beta.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}
	p.SetMetricName(dcl.ValueOrEmptyString(o.MetricName))
	p.SetTarget(dcl.ValueOrEmptyInt64(o.Target))
	return p
}

// EndpointDeployedModelsAutomaticResourcesToProto converts a EndpointDeployedModelsAutomaticResources object to its proto representation.
func VertexBetaEndpointDeployedModelsAutomaticResourcesToProto(o *beta.EndpointDeployedModelsAutomaticResources) *betapb.VertexBetaEndpointDeployedModelsAutomaticResources {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaEndpointDeployedModelsAutomaticResources{}
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	return p
}

// EndpointDeployedModelsPrivateEndpointsToProto converts a EndpointDeployedModelsPrivateEndpoints object to its proto representation.
func VertexBetaEndpointDeployedModelsPrivateEndpointsToProto(o *beta.EndpointDeployedModelsPrivateEndpoints) *betapb.VertexBetaEndpointDeployedModelsPrivateEndpoints {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaEndpointDeployedModelsPrivateEndpoints{}
	p.SetPredictHttpUri(dcl.ValueOrEmptyString(o.PredictHttpUri))
	p.SetExplainHttpUri(dcl.ValueOrEmptyString(o.ExplainHttpUri))
	p.SetHealthHttpUri(dcl.ValueOrEmptyString(o.HealthHttpUri))
	p.SetServiceAttachment(dcl.ValueOrEmptyString(o.ServiceAttachment))
	return p
}

// EndpointEncryptionSpecToProto converts a EndpointEncryptionSpec object to its proto representation.
func VertexBetaEndpointEncryptionSpecToProto(o *beta.EndpointEncryptionSpec) *betapb.VertexBetaEndpointEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &betapb.VertexBetaEndpointEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// EndpointToProto converts a Endpoint resource to its proto representation.
func EndpointToProto(resource *beta.Endpoint) *betapb.VertexBetaEndpoint {
	p := &betapb.VertexBetaEndpoint{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEncryptionSpec(VertexBetaEndpointEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetModelDeploymentMonitoringJob(dcl.ValueOrEmptyString(resource.ModelDeploymentMonitoringJob))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sDeployedModels := make([]*betapb.VertexBetaEndpointDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexBetaEndpointDeployedModelsToProto(&r)
	}
	p.SetDeployedModels(sDeployedModels)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyEndpoint handles the gRPC request by passing it to the underlying Endpoint Apply() method.
func (s *EndpointServer) applyEndpoint(ctx context.Context, c *beta.Client, request *betapb.ApplyVertexBetaEndpointRequest) (*betapb.VertexBetaEndpoint, error) {
	p := ProtoToEndpoint(request.GetResource())
	res, err := c.ApplyEndpoint(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointToProto(res)
	return r, nil
}

// applyVertexBetaEndpoint handles the gRPC request by passing it to the underlying Endpoint Apply() method.
func (s *EndpointServer) ApplyVertexBetaEndpoint(ctx context.Context, request *betapb.ApplyVertexBetaEndpointRequest) (*betapb.VertexBetaEndpoint, error) {
	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEndpoint(ctx, cl, request)
}

// DeleteEndpoint handles the gRPC request by passing it to the underlying Endpoint Delete() method.
func (s *EndpointServer) DeleteVertexBetaEndpoint(ctx context.Context, request *betapb.DeleteVertexBetaEndpointRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpoint(ctx, ProtoToEndpoint(request.GetResource()))

}

// ListVertexBetaEndpoint handles the gRPC request by passing it to the underlying EndpointList() method.
func (s *EndpointServer) ListVertexBetaEndpoint(ctx context.Context, request *betapb.ListVertexBetaEndpointRequest) (*betapb.ListVertexBetaEndpointResponse, error) {
	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpoint(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.VertexBetaEndpoint
	for _, r := range resources.Items {
		rp := EndpointToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListVertexBetaEndpointResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEndpoint(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
