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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/alpha/vertexai_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/alpha"
)

// EndpointServer implements the gRPC interface for Endpoint.
type EndpointServer struct{}

// ProtoToEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum converts a EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum enum from its proto representation.
func ProtoToVertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(e alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) *alpha.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum_name[int32(e)]; ok {
		e := alpha.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(n[len("VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointDeployedModels converts a EndpointDeployedModels object from its proto representation.
func ProtoToVertexaiAlphaEndpointDeployedModels(p *alphapb.VertexaiAlphaEndpointDeployedModels) *alpha.EndpointDeployedModels {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModels{
		DedicatedResources:     ProtoToVertexaiAlphaEndpointDeployedModelsDedicatedResources(p.GetDedicatedResources()),
		AutomaticResources:     ProtoToVertexaiAlphaEndpointDeployedModelsAutomaticResources(p.GetAutomaticResources()),
		Id:                     dcl.StringOrNil(p.GetId()),
		Model:                  dcl.StringOrNil(p.GetModel()),
		ModelVersionId:         dcl.StringOrNil(p.GetModelVersionId()),
		DisplayName:            dcl.StringOrNil(p.GetDisplayName()),
		CreateTime:             dcl.StringOrNil(p.GetCreateTime()),
		ServiceAccount:         dcl.StringOrNil(p.GetServiceAccount()),
		EnableAccessLogging:    dcl.Bool(p.GetEnableAccessLogging()),
		PrivateEndpoints:       ProtoToVertexaiAlphaEndpointDeployedModelsPrivateEndpoints(p.GetPrivateEndpoints()),
		SharedResources:        dcl.StringOrNil(p.GetSharedResources()),
		EnableContainerLogging: dcl.Bool(p.GetEnableContainerLogging()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResources converts a EndpointDeployedModelsDedicatedResources object from its proto representation.
func ProtoToVertexaiAlphaEndpointDeployedModelsDedicatedResources(p *alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResources) *alpha.EndpointDeployedModelsDedicatedResources {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModelsDedicatedResources{
		MachineSpec:     ProtoToVertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec(p.GetMachineSpec()),
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	for _, r := range p.GetAutoscalingMetricSpecs() {
		obj.AutoscalingMetricSpecs = append(obj.AutoscalingMetricSpecs, *ProtoToVertexaiAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(r))
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResourcesMachineSpec converts a EndpointDeployedModelsDedicatedResourcesMachineSpec object from its proto representation.
func ProtoToVertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec(p *alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec) *alpha.EndpointDeployedModelsDedicatedResourcesMachineSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModelsDedicatedResourcesMachineSpec{
		MachineType:      dcl.StringOrNil(p.GetMachineType()),
		AcceleratorType:  ProtoToVertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs converts a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs object from its proto representation.
func ProtoToVertexaiAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(p *alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *alpha.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{
		MetricName: dcl.StringOrNil(p.GetMetricName()),
		Target:     dcl.Int64OrNil(p.GetTarget()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsAutomaticResources converts a EndpointDeployedModelsAutomaticResources object from its proto representation.
func ProtoToVertexaiAlphaEndpointDeployedModelsAutomaticResources(p *alphapb.VertexaiAlphaEndpointDeployedModelsAutomaticResources) *alpha.EndpointDeployedModelsAutomaticResources {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModelsAutomaticResources{
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsPrivateEndpoints converts a EndpointDeployedModelsPrivateEndpoints object from its proto representation.
func ProtoToVertexaiAlphaEndpointDeployedModelsPrivateEndpoints(p *alphapb.VertexaiAlphaEndpointDeployedModelsPrivateEndpoints) *alpha.EndpointDeployedModelsPrivateEndpoints {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModelsPrivateEndpoints{
		PredictHttpUri:    dcl.StringOrNil(p.GetPredictHttpUri()),
		ExplainHttpUri:    dcl.StringOrNil(p.GetExplainHttpUri()),
		HealthHttpUri:     dcl.StringOrNil(p.GetHealthHttpUri()),
		ServiceAttachment: dcl.StringOrNil(p.GetServiceAttachment()),
	}
	return obj
}

// ProtoToEndpointEncryptionSpec converts a EndpointEncryptionSpec object from its proto representation.
func ProtoToVertexaiAlphaEndpointEncryptionSpec(p *alphapb.VertexaiAlphaEndpointEncryptionSpec) *alpha.EndpointEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToEndpoint converts a Endpoint resource from its proto representation.
func ProtoToEndpoint(p *alphapb.VertexaiAlphaEndpoint) *alpha.Endpoint {
	obj := &alpha.Endpoint{
		Name:                         dcl.StringOrNil(p.GetName()),
		DisplayName:                  dcl.StringOrNil(p.GetDisplayName()),
		Description:                  dcl.StringOrNil(p.GetDescription()),
		Etag:                         dcl.StringOrNil(p.GetEtag()),
		CreateTime:                   dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		EncryptionSpec:               ProtoToVertexaiAlphaEndpointEncryptionSpec(p.GetEncryptionSpec()),
		Network:                      dcl.StringOrNil(p.GetNetwork()),
		ModelDeploymentMonitoringJob: dcl.StringOrNil(p.GetModelDeploymentMonitoringJob()),
		Project:                      dcl.StringOrNil(p.GetProject()),
		Location:                     dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexaiAlphaEndpointDeployedModels(r))
	}
	return obj
}

// EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto converts a EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum enum to its proto representation.
func VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto(e *alpha.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	if e == nil {
		return alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(0)
	}
	if v, ok := alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum_value["EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"+string(*e)]; ok {
		return alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(v)
	}
	return alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(0)
}

// EndpointDeployedModelsToProto converts a EndpointDeployedModels object to its proto representation.
func VertexaiAlphaEndpointDeployedModelsToProto(o *alpha.EndpointDeployedModels) *alphapb.VertexaiAlphaEndpointDeployedModels {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaEndpointDeployedModels{}
	p.SetDedicatedResources(VertexaiAlphaEndpointDeployedModelsDedicatedResourcesToProto(o.DedicatedResources))
	p.SetAutomaticResources(VertexaiAlphaEndpointDeployedModelsAutomaticResourcesToProto(o.AutomaticResources))
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	p.SetModelVersionId(dcl.ValueOrEmptyString(o.ModelVersionId))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetEnableAccessLogging(dcl.ValueOrEmptyBool(o.EnableAccessLogging))
	p.SetPrivateEndpoints(VertexaiAlphaEndpointDeployedModelsPrivateEndpointsToProto(o.PrivateEndpoints))
	p.SetSharedResources(dcl.ValueOrEmptyString(o.SharedResources))
	p.SetEnableContainerLogging(dcl.ValueOrEmptyBool(o.EnableContainerLogging))
	return p
}

// EndpointDeployedModelsDedicatedResourcesToProto converts a EndpointDeployedModelsDedicatedResources object to its proto representation.
func VertexaiAlphaEndpointDeployedModelsDedicatedResourcesToProto(o *alpha.EndpointDeployedModelsDedicatedResources) *alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResources {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResources{}
	p.SetMachineSpec(VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecToProto(o.MachineSpec))
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	sAutoscalingMetricSpecs := make([]*alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, len(o.AutoscalingMetricSpecs))
	for i, r := range o.AutoscalingMetricSpecs {
		sAutoscalingMetricSpecs[i] = VertexaiAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto(&r)
	}
	p.SetAutoscalingMetricSpecs(sAutoscalingMetricSpecs)
	return p
}

// EndpointDeployedModelsDedicatedResourcesMachineSpecToProto converts a EndpointDeployedModelsDedicatedResourcesMachineSpec object to its proto representation.
func VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecToProto(o *alpha.EndpointDeployedModelsDedicatedResourcesMachineSpec) *alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetAcceleratorType(VertexaiAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto converts a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs object to its proto representation.
func VertexaiAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto(o *alpha.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}
	p.SetMetricName(dcl.ValueOrEmptyString(o.MetricName))
	p.SetTarget(dcl.ValueOrEmptyInt64(o.Target))
	return p
}

// EndpointDeployedModelsAutomaticResourcesToProto converts a EndpointDeployedModelsAutomaticResources object to its proto representation.
func VertexaiAlphaEndpointDeployedModelsAutomaticResourcesToProto(o *alpha.EndpointDeployedModelsAutomaticResources) *alphapb.VertexaiAlphaEndpointDeployedModelsAutomaticResources {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaEndpointDeployedModelsAutomaticResources{}
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	return p
}

// EndpointDeployedModelsPrivateEndpointsToProto converts a EndpointDeployedModelsPrivateEndpoints object to its proto representation.
func VertexaiAlphaEndpointDeployedModelsPrivateEndpointsToProto(o *alpha.EndpointDeployedModelsPrivateEndpoints) *alphapb.VertexaiAlphaEndpointDeployedModelsPrivateEndpoints {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaEndpointDeployedModelsPrivateEndpoints{}
	p.SetPredictHttpUri(dcl.ValueOrEmptyString(o.PredictHttpUri))
	p.SetExplainHttpUri(dcl.ValueOrEmptyString(o.ExplainHttpUri))
	p.SetHealthHttpUri(dcl.ValueOrEmptyString(o.HealthHttpUri))
	p.SetServiceAttachment(dcl.ValueOrEmptyString(o.ServiceAttachment))
	return p
}

// EndpointEncryptionSpecToProto converts a EndpointEncryptionSpec object to its proto representation.
func VertexaiAlphaEndpointEncryptionSpecToProto(o *alpha.EndpointEncryptionSpec) *alphapb.VertexaiAlphaEndpointEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaEndpointEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// EndpointToProto converts a Endpoint resource to its proto representation.
func EndpointToProto(resource *alpha.Endpoint) *alphapb.VertexaiAlphaEndpoint {
	p := &alphapb.VertexaiAlphaEndpoint{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEncryptionSpec(VertexaiAlphaEndpointEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetModelDeploymentMonitoringJob(dcl.ValueOrEmptyString(resource.ModelDeploymentMonitoringJob))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sDeployedModels := make([]*alphapb.VertexaiAlphaEndpointDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexaiAlphaEndpointDeployedModelsToProto(&r)
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
func (s *EndpointServer) applyEndpoint(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVertexaiAlphaEndpointRequest) (*alphapb.VertexaiAlphaEndpoint, error) {
	p := ProtoToEndpoint(request.GetResource())
	res, err := c.ApplyEndpoint(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointToProto(res)
	return r, nil
}

// applyVertexaiAlphaEndpoint handles the gRPC request by passing it to the underlying Endpoint Apply() method.
func (s *EndpointServer) ApplyVertexaiAlphaEndpoint(ctx context.Context, request *alphapb.ApplyVertexaiAlphaEndpointRequest) (*alphapb.VertexaiAlphaEndpoint, error) {
	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEndpoint(ctx, cl, request)
}

// DeleteEndpoint handles the gRPC request by passing it to the underlying Endpoint Delete() method.
func (s *EndpointServer) DeleteVertexaiAlphaEndpoint(ctx context.Context, request *alphapb.DeleteVertexaiAlphaEndpointRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpoint(ctx, ProtoToEndpoint(request.GetResource()))

}

// ListVertexaiAlphaEndpoint handles the gRPC request by passing it to the underlying EndpointList() method.
func (s *EndpointServer) ListVertexaiAlphaEndpoint(ctx context.Context, request *alphapb.ListVertexaiAlphaEndpointRequest) (*alphapb.ListVertexaiAlphaEndpointResponse, error) {
	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpoint(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VertexaiAlphaEndpoint
	for _, r := range resources.Items {
		rp := EndpointToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVertexaiAlphaEndpointResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEndpoint(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
