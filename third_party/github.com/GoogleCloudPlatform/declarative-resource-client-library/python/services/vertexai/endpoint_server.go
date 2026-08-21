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
	vertexaipb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/vertexai_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai"
)

// EndpointServer implements the gRPC interface for Endpoint.
type EndpointServer struct{}

// ProtoToEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum converts a EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum enum from its proto representation.
func ProtoToVertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(e vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) *vertexai.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum_name[int32(e)]; ok {
		e := vertexai.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(n[len("VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointDeployedModels converts a EndpointDeployedModels object from its proto representation.
func ProtoToVertexaiEndpointDeployedModels(p *vertexaipb.VertexaiEndpointDeployedModels) *vertexai.EndpointDeployedModels {
	if p == nil {
		return nil
	}
	obj := &vertexai.EndpointDeployedModels{
		DedicatedResources:  ProtoToVertexaiEndpointDeployedModelsDedicatedResources(p.GetDedicatedResources()),
		AutomaticResources:  ProtoToVertexaiEndpointDeployedModelsAutomaticResources(p.GetAutomaticResources()),
		Id:                  dcl.StringOrNil(p.GetId()),
		Model:               dcl.StringOrNil(p.GetModel()),
		ModelVersionId:      dcl.StringOrNil(p.GetModelVersionId()),
		DisplayName:         dcl.StringOrNil(p.GetDisplayName()),
		CreateTime:          dcl.StringOrNil(p.GetCreateTime()),
		ServiceAccount:      dcl.StringOrNil(p.GetServiceAccount()),
		EnableAccessLogging: dcl.Bool(p.GetEnableAccessLogging()),
		PrivateEndpoints:    ProtoToVertexaiEndpointDeployedModelsPrivateEndpoints(p.GetPrivateEndpoints()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResources converts a EndpointDeployedModelsDedicatedResources object from its proto representation.
func ProtoToVertexaiEndpointDeployedModelsDedicatedResources(p *vertexaipb.VertexaiEndpointDeployedModelsDedicatedResources) *vertexai.EndpointDeployedModelsDedicatedResources {
	if p == nil {
		return nil
	}
	obj := &vertexai.EndpointDeployedModelsDedicatedResources{
		MachineSpec:     ProtoToVertexaiEndpointDeployedModelsDedicatedResourcesMachineSpec(p.GetMachineSpec()),
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	for _, r := range p.GetAutoscalingMetricSpecs() {
		obj.AutoscalingMetricSpecs = append(obj.AutoscalingMetricSpecs, *ProtoToVertexaiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(r))
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResourcesMachineSpec converts a EndpointDeployedModelsDedicatedResourcesMachineSpec object from its proto representation.
func ProtoToVertexaiEndpointDeployedModelsDedicatedResourcesMachineSpec(p *vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpec) *vertexai.EndpointDeployedModelsDedicatedResourcesMachineSpec {
	if p == nil {
		return nil
	}
	obj := &vertexai.EndpointDeployedModelsDedicatedResourcesMachineSpec{
		MachineType:      dcl.StringOrNil(p.GetMachineType()),
		AcceleratorType:  ProtoToVertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs converts a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs object from its proto representation.
func ProtoToVertexaiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs(p *vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *vertexai.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if p == nil {
		return nil
	}
	obj := &vertexai.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{
		MetricName: dcl.StringOrNil(p.GetMetricName()),
		Target:     dcl.Int64OrNil(p.GetTarget()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsAutomaticResources converts a EndpointDeployedModelsAutomaticResources object from its proto representation.
func ProtoToVertexaiEndpointDeployedModelsAutomaticResources(p *vertexaipb.VertexaiEndpointDeployedModelsAutomaticResources) *vertexai.EndpointDeployedModelsAutomaticResources {
	if p == nil {
		return nil
	}
	obj := &vertexai.EndpointDeployedModelsAutomaticResources{
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	return obj
}

// ProtoToEndpointDeployedModelsPrivateEndpoints converts a EndpointDeployedModelsPrivateEndpoints object from its proto representation.
func ProtoToVertexaiEndpointDeployedModelsPrivateEndpoints(p *vertexaipb.VertexaiEndpointDeployedModelsPrivateEndpoints) *vertexai.EndpointDeployedModelsPrivateEndpoints {
	if p == nil {
		return nil
	}
	obj := &vertexai.EndpointDeployedModelsPrivateEndpoints{
		PredictHttpUri:    dcl.StringOrNil(p.GetPredictHttpUri()),
		ExplainHttpUri:    dcl.StringOrNil(p.GetExplainHttpUri()),
		HealthHttpUri:     dcl.StringOrNil(p.GetHealthHttpUri()),
		ServiceAttachment: dcl.StringOrNil(p.GetServiceAttachment()),
	}
	return obj
}

// ProtoToEndpointEncryptionSpec converts a EndpointEncryptionSpec object from its proto representation.
func ProtoToVertexaiEndpointEncryptionSpec(p *vertexaipb.VertexaiEndpointEncryptionSpec) *vertexai.EndpointEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &vertexai.EndpointEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToEndpoint converts a Endpoint resource from its proto representation.
func ProtoToEndpoint(p *vertexaipb.VertexaiEndpoint) *vertexai.Endpoint {
	obj := &vertexai.Endpoint{
		Name:                         dcl.StringOrNil(p.GetName()),
		DisplayName:                  dcl.StringOrNil(p.GetDisplayName()),
		Description:                  dcl.StringOrNil(p.GetDescription()),
		Etag:                         dcl.StringOrNil(p.GetEtag()),
		CreateTime:                   dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:                   dcl.StringOrNil(p.GetUpdateTime()),
		EncryptionSpec:               ProtoToVertexaiEndpointEncryptionSpec(p.GetEncryptionSpec()),
		Network:                      dcl.StringOrNil(p.GetNetwork()),
		ModelDeploymentMonitoringJob: dcl.StringOrNil(p.GetModelDeploymentMonitoringJob()),
		Project:                      dcl.StringOrNil(p.GetProject()),
		Location:                     dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetDeployedModels() {
		obj.DeployedModels = append(obj.DeployedModels, *ProtoToVertexaiEndpointDeployedModels(r))
	}
	return obj
}

// EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto converts a EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum enum to its proto representation.
func VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto(e *vertexai.EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum) vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum {
	if e == nil {
		return vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(0)
	}
	if v, ok := vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum_value["EndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum"+string(*e)]; ok {
		return vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(v)
	}
	return vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnum(0)
}

// EndpointDeployedModelsToProto converts a EndpointDeployedModels object to its proto representation.
func VertexaiEndpointDeployedModelsToProto(o *vertexai.EndpointDeployedModels) *vertexaipb.VertexaiEndpointDeployedModels {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiEndpointDeployedModels{}
	p.SetDedicatedResources(VertexaiEndpointDeployedModelsDedicatedResourcesToProto(o.DedicatedResources))
	p.SetAutomaticResources(VertexaiEndpointDeployedModelsAutomaticResourcesToProto(o.AutomaticResources))
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetModel(dcl.ValueOrEmptyString(o.Model))
	p.SetModelVersionId(dcl.ValueOrEmptyString(o.ModelVersionId))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetEnableAccessLogging(dcl.ValueOrEmptyBool(o.EnableAccessLogging))
	p.SetPrivateEndpoints(VertexaiEndpointDeployedModelsPrivateEndpointsToProto(o.PrivateEndpoints))
	return p
}

// EndpointDeployedModelsDedicatedResourcesToProto converts a EndpointDeployedModelsDedicatedResources object to its proto representation.
func VertexaiEndpointDeployedModelsDedicatedResourcesToProto(o *vertexai.EndpointDeployedModelsDedicatedResources) *vertexaipb.VertexaiEndpointDeployedModelsDedicatedResources {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiEndpointDeployedModelsDedicatedResources{}
	p.SetMachineSpec(VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecToProto(o.MachineSpec))
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	sAutoscalingMetricSpecs := make([]*vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs, len(o.AutoscalingMetricSpecs))
	for i, r := range o.AutoscalingMetricSpecs {
		sAutoscalingMetricSpecs[i] = VertexaiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto(&r)
	}
	p.SetAutoscalingMetricSpecs(sAutoscalingMetricSpecs)
	return p
}

// EndpointDeployedModelsDedicatedResourcesMachineSpecToProto converts a EndpointDeployedModelsDedicatedResourcesMachineSpec object to its proto representation.
func VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecToProto(o *vertexai.EndpointDeployedModelsDedicatedResourcesMachineSpec) *vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpec {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpec{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetAcceleratorType(VertexaiEndpointDeployedModelsDedicatedResourcesMachineSpecAcceleratorTypeEnumToProto(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto converts a EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs object to its proto representation.
func VertexaiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecsToProto(o *vertexai.EndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs) *vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiEndpointDeployedModelsDedicatedResourcesAutoscalingMetricSpecs{}
	p.SetMetricName(dcl.ValueOrEmptyString(o.MetricName))
	p.SetTarget(dcl.ValueOrEmptyInt64(o.Target))
	return p
}

// EndpointDeployedModelsAutomaticResourcesToProto converts a EndpointDeployedModelsAutomaticResources object to its proto representation.
func VertexaiEndpointDeployedModelsAutomaticResourcesToProto(o *vertexai.EndpointDeployedModelsAutomaticResources) *vertexaipb.VertexaiEndpointDeployedModelsAutomaticResources {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiEndpointDeployedModelsAutomaticResources{}
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	return p
}

// EndpointDeployedModelsPrivateEndpointsToProto converts a EndpointDeployedModelsPrivateEndpoints object to its proto representation.
func VertexaiEndpointDeployedModelsPrivateEndpointsToProto(o *vertexai.EndpointDeployedModelsPrivateEndpoints) *vertexaipb.VertexaiEndpointDeployedModelsPrivateEndpoints {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiEndpointDeployedModelsPrivateEndpoints{}
	p.SetPredictHttpUri(dcl.ValueOrEmptyString(o.PredictHttpUri))
	p.SetExplainHttpUri(dcl.ValueOrEmptyString(o.ExplainHttpUri))
	p.SetHealthHttpUri(dcl.ValueOrEmptyString(o.HealthHttpUri))
	p.SetServiceAttachment(dcl.ValueOrEmptyString(o.ServiceAttachment))
	return p
}

// EndpointEncryptionSpecToProto converts a EndpointEncryptionSpec object to its proto representation.
func VertexaiEndpointEncryptionSpecToProto(o *vertexai.EndpointEncryptionSpec) *vertexaipb.VertexaiEndpointEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiEndpointEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// EndpointToProto converts a Endpoint resource to its proto representation.
func EndpointToProto(resource *vertexai.Endpoint) *vertexaipb.VertexaiEndpoint {
	p := &vertexaipb.VertexaiEndpoint{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEncryptionSpec(VertexaiEndpointEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetModelDeploymentMonitoringJob(dcl.ValueOrEmptyString(resource.ModelDeploymentMonitoringJob))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sDeployedModels := make([]*vertexaipb.VertexaiEndpointDeployedModels, len(resource.DeployedModels))
	for i, r := range resource.DeployedModels {
		sDeployedModels[i] = VertexaiEndpointDeployedModelsToProto(&r)
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
func (s *EndpointServer) applyEndpoint(ctx context.Context, c *vertexai.Client, request *vertexaipb.ApplyVertexaiEndpointRequest) (*vertexaipb.VertexaiEndpoint, error) {
	p := ProtoToEndpoint(request.GetResource())
	res, err := c.ApplyEndpoint(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointToProto(res)
	return r, nil
}

// applyVertexaiEndpoint handles the gRPC request by passing it to the underlying Endpoint Apply() method.
func (s *EndpointServer) ApplyVertexaiEndpoint(ctx context.Context, request *vertexaipb.ApplyVertexaiEndpointRequest) (*vertexaipb.VertexaiEndpoint, error) {
	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEndpoint(ctx, cl, request)
}

// DeleteEndpoint handles the gRPC request by passing it to the underlying Endpoint Delete() method.
func (s *EndpointServer) DeleteVertexaiEndpoint(ctx context.Context, request *vertexaipb.DeleteVertexaiEndpointRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpoint(ctx, ProtoToEndpoint(request.GetResource()))

}

// ListVertexaiEndpoint handles the gRPC request by passing it to the underlying EndpointList() method.
func (s *EndpointServer) ListVertexaiEndpoint(ctx context.Context, request *vertexaipb.ListVertexaiEndpointRequest) (*vertexaipb.ListVertexaiEndpointResponse, error) {
	cl, err := createConfigEndpoint(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpoint(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*vertexaipb.VertexaiEndpoint
	for _, r := range resources.Items {
		rp := EndpointToProto(r)
		protos = append(protos, rp)
	}
	p := &vertexaipb.ListVertexaiEndpointResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEndpoint(ctx context.Context, service_account_file string) (*vertexai.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return vertexai.NewClient(conf), nil
}
