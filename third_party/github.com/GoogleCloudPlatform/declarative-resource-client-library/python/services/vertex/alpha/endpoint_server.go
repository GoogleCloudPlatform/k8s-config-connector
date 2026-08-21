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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/alpha/vertex_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex/alpha"
)

// EndpointServer implements the gRPC interface for Endpoint.
type EndpointServer struct{}

// ProtoToEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum converts a EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum enum from its proto representation.
func ProtoToVertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(e alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) *alpha.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum_name[int32(e)]; ok {
		e := alpha.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(n[len("VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointDeployedModels converts a EndpointDeployedModels object from its proto representation.
func ProtoToVertexAlphaEndpointDeployedModels(p *alphapb.VertexAlphaEndpointDeployedModels) *alpha.EndpointDeployedModels {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModels{
		DedicatedResources:      ProtoToVertexAlphaEndpointDeployedModelsDedicatedResources(p.GetDedicatedResources()),
		AutomaticResources:      ProtoToVertexAlphaEndpointDeployedModelsAutomaticResources(p.GetAutomaticResources()),
		Id:                      dcl.StringOrNil(p.GetId()),
		Model:                   dcl.StringOrNil(p.GetModel()),
		ModelVersionId:          dcl.StringOrNil(p.GetModelVersionId()),
		DisplayName:             dcl.StringOrNil(p.GetDisplayName()),
		CreateTime:              dcl.StringOrNil(p.GetCreateTime()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
		DisableContainerLogging: dcl.Bool(p.GetDisableContainerLogging()),
		EnableAccessLogging:     dcl.Bool(p.GetEnableAccessLogging()),
		PrivateEndpoints:        ProtoToVertexAlphaEndpointDeployedModelsPrivateEndpoints(p.GetPrivateEndpoints()),
		SharedResources:         dcl.StringOrNil(p.GetSharedResources()),
		EnableContainerLogging:  dcl.Bool(p.GetEnableContainerLogging()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResources converts a EndpointDeployedModelsDedicatedResources object from its proto representation.
func ProtoToVertexAlphaEndpointDeployedModelsDedicatedResources(p *alphapb.VertexAlphaEndpointDeployedModelsDedicatedResources) *alpha.EndpointDeployedModelsDedicatedResources {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModelsDedicatedResources{
		MachineSpec:     ProtoToVertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec(p.GetMachineSpec()),
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	for _, r := range p.GetAutoscalingMetricSpecs() {
		obj.AutoscalingMetricSpecs = append(obj.AutoscalingMetricSpecs, *ProtoToVertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(r))
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResourcesMachineSpec converts a EndpointDeployedModelsDedicatedResourcesMachineSpec object from its proto representation.
func ProtoToVertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec(p *alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec) *alpha.EndpointDeployedModelsDedicatedResourcesMachineSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointDeployedModelsDedicatedResourcesMachineSpec{
		MachineType:      dcl.StringOrNil(p.GetMachineType()),
		AcceleratorType:  ProtoToVertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs converts a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs object from its proto representation.
func ProtoToVertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(p *alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *alpha.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
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
func ProtoToVertexAlphaEndpointDeployedModelsAutomaticResources(p *alphapb.VertexAlphaEndpointDeployedModelsAutomaticResources) *alpha.EndpointDeployedModelsAutomaticResources {
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
func ProtoToVertexAlphaEndpointDeployedModelsPrivateEndpoints(p *alphapb.VertexAlphaEndpointDeployedModelsPrivateEndpoints) *alpha.EndpointDeployedModelsPrivateEndpoints {
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
func ProtoToVertexAlphaEndpointEncryptionSpec(p *alphapb.VertexAlphaEndpointEncryptionSpec) *alpha.EndpointEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToEndpoint converts a Endpoint resource from its proto representation.
func ProtoToEndpoint(p *alphapb.VertexAlphaEndpoint) *alpha.Endpoint {
	obj := &alpha.Endpoint{
		Name:                         dcl.StringOrNil(p.GetName()),
		DisplayName:                  dcl.StringOrNil(p.GetDisplayName()),
		Description:                  dcl.StringOrNil(p.GetDescription()),
		Etag:                         dcl.StringOrNil(p.GetEtag()),
		CreateTime:                   dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		EncryptionSpec:               ProtoToVertexAlphaEndpointEncryptionSpec(p.GetEncryptionSpec()),
		Network:                      dcl.StringOrNil(p.GetNetwork()),
		ModelDeploymentMonitoringJob: dcl.StringOrNil(p.GetModelDeploymentMonitoringJob()),
		Project:                      dcl.StringOrNil(p.GetProject()),
		Location:                     dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexAlphaEndpointDeployedModels(r))
	}
	return obj
}

// EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto converts a EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum enum to its proto representation.
func VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto(e *alpha.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	if e == nil {
		return alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(0)
	}
	if v, ok := alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum_value["EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"+string(*e)]; ok {
		return alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(v)
	}
	return alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(0)
}

// EndpointDeployedModelsToProto converts a EndpointDeployedModels object to its proto representation.
func VertexAlphaEndpointDeployedModelsToProto(o *alpha.EndpointDeployedModels) *alphapb.VertexAlphaEndpointDeployedModels {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaEndpointDeployedModels{}
	p.SetDedicatedResources(VertexAlphaEndpointDeployedModelsDedicatedResourcesToProto(o.DedicatedResources))
	p.SetAutomaticResources(VertexAlphaEndpointDeployedModelsAutomaticResourcesToProto(o.AutomaticResources))
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	p.SetModelVersionId(dcl.ValueOrEmptyString(o.ModelVersionId))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetDisableContainerLogging(dcl.ValueOrEmptyBool(o.DisableContainerLogging))
	p.SetEnableAccessLogging(dcl.ValueOrEmptyBool(o.EnableAccessLogging))
	p.SetPrivateEndpoints(VertexAlphaEndpointDeployedModelsPrivateEndpointsToProto(o.PrivateEndpoints))
	p.SetSharedResources(dcl.ValueOrEmptyString(o.SharedResources))
	p.SetEnableContainerLogging(dcl.ValueOrEmptyBool(o.EnableContainerLogging))
	return p
}

// EndpointDeployedModelsDedicatedResourcesToProto converts a EndpointDeployedModelsDedicatedResources object to its proto representation.
func VertexAlphaEndpointDeployedModelsDedicatedResourcesToProto(o *alpha.EndpointDeployedModelsDedicatedResources) *alphapb.VertexAlphaEndpointDeployedModelsDedicatedResources {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaEndpointDeployedModelsDedicatedResources{}
	p.SetMachineSpec(VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecToProto(o.MachineSpec))
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	sAutoscalingMetricSpecs := make([]*alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, len(o.AutoscalingMetricSpecs))
	for i, r := range o.AutoscalingMetricSpecs {
		sAutoscalingMetricSpecs[i] = VertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto(&r)
	}
	p.SetAutoscalingMetricSpecs(sAutoscalingMetricSpecs)
	return p
}

// EndpointDeployedModelsDedicatedResourcesMachineSpecToProto converts a EndpointDeployedModelsDedicatedResourcesMachineSpec object to its proto representation.
func VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecToProto(o *alpha.EndpointDeployedModelsDedicatedResourcesMachineSpec) *alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpec{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetAcceleratorType(VertexAlphaEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto converts a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs object to its proto representation.
func VertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto(o *alpha.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}
	p.SetMetricName(dcl.ValueOrEmptyString(o.MetricName))
	p.SetTarget(dcl.ValueOrEmptyInt64(o.Target))
	return p
}

// EndpointDeployedModelsAutomaticResourcesToProto converts a EndpointDeployedModelsAutomaticResources object to its proto representation.
func VertexAlphaEndpointDeployedModelsAutomaticResourcesToProto(o *alpha.EndpointDeployedModelsAutomaticResources) *alphapb.VertexAlphaEndpointDeployedModelsAutomaticResources {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaEndpointDeployedModelsAutomaticResources{}
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	return p
}

// EndpointDeployedModelsPrivateEndpointsToProto converts a EndpointDeployedModelsPrivateEndpoints object to its proto representation.
func VertexAlphaEndpointDeployedModelsPrivateEndpointsToProto(o *alpha.EndpointDeployedModelsPrivateEndpoints) *alphapb.VertexAlphaEndpointDeployedModelsPrivateEndpoints {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaEndpointDeployedModelsPrivateEndpoints{}
	p.SetPredictHttpUri(dcl.ValueOrEmptyString(o.PredictHttpUri))
	p.SetExplainHttpUri(dcl.ValueOrEmptyString(o.ExplainHttpUri))
	p.SetHealthHttpUri(dcl.ValueOrEmptyString(o.HealthHttpUri))
	p.SetServiceAttachment(dcl.ValueOrEmptyString(o.ServiceAttachment))
	return p
}

// EndpointEncryptionSpecToProto converts a EndpointEncryptionSpec object to its proto representation.
func VertexAlphaEndpointEncryptionSpecToProto(o *alpha.EndpointEncryptionSpec) *alphapb.VertexAlphaEndpointEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaEndpointEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// EndpointToProto converts a Endpoint resource to its proto representation.
func EndpointToProto(resource *alpha.Endpoint) *alphapb.VertexAlphaEndpoint {
	p := &alphapb.VertexAlphaEndpoint{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEncryptionSpec(VertexAlphaEndpointEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetModelDeploymentMonitoringJob(dcl.ValueOrEmptyString(resource.ModelDeploymentMonitoringJob))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sDeployedModels := make([]*alphapb.VertexAlphaEndpointDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexAlphaEndpointDeployedModelsToProto(&r)
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
func (s *EndpointServer) applyEndpoint(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVertexAlphaEndpointRequest) (*alphapb.VertexAlphaEndpoint, error) {
	p := ProtoToEndpoint(request.GetResource())
	res, err := c.ApplyEndpoint(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointToProto(res)
	return r, nil
}

// applyVertexAlphaEndpoint handles the gRPC request by passing it to the underlying Endpoint Apply() method.
func (s *EndpointServer) ApplyVertexAlphaEndpoint(ctx context.Context, request *alphapb.ApplyVertexAlphaEndpointRequest) (*alphapb.VertexAlphaEndpoint, error) {
	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEndpoint(ctx, cl, request)
}

// DeleteEndpoint handles the gRPC request by passing it to the underlying Endpoint Delete() method.
func (s *EndpointServer) DeleteVertexAlphaEndpoint(ctx context.Context, request *alphapb.DeleteVertexAlphaEndpointRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpoint(ctx, ProtoToEndpoint(request.GetResource()))

}

// ListVertexAlphaEndpoint handles the gRPC request by passing it to the underlying EndpointList() method.
func (s *EndpointServer) ListVertexAlphaEndpoint(ctx context.Context, request *alphapb.ListVertexAlphaEndpointRequest) (*alphapb.ListVertexAlphaEndpointResponse, error) {
	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpoint(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VertexAlphaEndpoint
	for _, r := range resources.Items {
		rp := EndpointToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVertexAlphaEndpointResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEndpoint(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
