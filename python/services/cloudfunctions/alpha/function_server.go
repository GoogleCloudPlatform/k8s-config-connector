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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudfunctions/alpha/cloudfunctions_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/alpha"
)

// FunctionServer implements the gRPC interface for Function.
type FunctionServer struct{}

// ProtoToFunctionHttpsTriggerSecurityLevelEnum converts a FunctionHttpsTriggerSecurityLevelEnum enum from its proto representation.
func ProtoToCloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum(e alphapb.CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum) *alpha.FunctionHttpsTriggerSecurityLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum_name[int32(e)]; ok {
		e := alpha.FunctionHttpsTriggerSecurityLevelEnum(n[len("CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionStatusEnum converts a FunctionStatusEnum enum from its proto representation.
func ProtoToCloudfunctionsAlphaFunctionStatusEnum(e alphapb.CloudfunctionsAlphaFunctionStatusEnum) *alpha.FunctionStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudfunctionsAlphaFunctionStatusEnum_name[int32(e)]; ok {
		e := alpha.FunctionStatusEnum(n[len("CloudfunctionsAlphaFunctionStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionVPCConnectorEgressSettingsEnum converts a FunctionVPCConnectorEgressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum(e alphapb.CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum) *alpha.FunctionVPCConnectorEgressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum_name[int32(e)]; ok {
		e := alpha.FunctionVPCConnectorEgressSettingsEnum(n[len("CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionIngressSettingsEnum converts a FunctionIngressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsAlphaFunctionIngressSettingsEnum(e alphapb.CloudfunctionsAlphaFunctionIngressSettingsEnum) *alpha.FunctionIngressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.CloudfunctionsAlphaFunctionIngressSettingsEnum_name[int32(e)]; ok {
		e := alpha.FunctionIngressSettingsEnum(n[len("CloudfunctionsAlphaFunctionIngressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionSourceRepository converts a FunctionSourceRepository object from its proto representation.
func ProtoToCloudfunctionsAlphaFunctionSourceRepository(p *alphapb.CloudfunctionsAlphaFunctionSourceRepository) *alpha.FunctionSourceRepository {
	if p == nil {
		return nil
	}
	obj := &alpha.FunctionSourceRepository{
		Url:         dcl.StringOrNil(p.GetUrl()),
		DeployedUrl: dcl.StringOrNil(p.GetDeployedUrl()),
	}
	return obj
}

// ProtoToFunctionHttpsTrigger converts a FunctionHttpsTrigger object from its proto representation.
func ProtoToCloudfunctionsAlphaFunctionHttpsTrigger(p *alphapb.CloudfunctionsAlphaFunctionHttpsTrigger) *alpha.FunctionHttpsTrigger {
	if p == nil {
		return nil
	}
	obj := &alpha.FunctionHttpsTrigger{
		Url:           dcl.StringOrNil(p.GetUrl()),
		SecurityLevel: ProtoToCloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum(p.GetSecurityLevel()),
	}
	return obj
}

// ProtoToFunctionEventTrigger converts a FunctionEventTrigger object from its proto representation.
func ProtoToCloudfunctionsAlphaFunctionEventTrigger(p *alphapb.CloudfunctionsAlphaFunctionEventTrigger) *alpha.FunctionEventTrigger {
	if p == nil {
		return nil
	}
	obj := &alpha.FunctionEventTrigger{
		EventType:     dcl.StringOrNil(p.GetEventType()),
		Resource:      dcl.StringOrNil(p.GetResource()),
		Service:       dcl.StringOrNil(p.GetService()),
		FailurePolicy: dcl.Bool(p.GetFailurePolicy()),
	}
	return obj
}

// ProtoToFunction converts a Function resource from its proto representation.
func ProtoToFunction(p *alphapb.CloudfunctionsAlphaFunction) *alpha.Function {
	obj := &alpha.Function{
		Name:                       dcl.StringOrNil(p.GetName()),
		Description:                dcl.StringOrNil(p.GetDescription()),
		SourceArchiveUrl:           dcl.StringOrNil(p.GetSourceArchiveUrl()),
		SourceRepository:           ProtoToCloudfunctionsAlphaFunctionSourceRepository(p.GetSourceRepository()),
		HttpsTrigger:               ProtoToCloudfunctionsAlphaFunctionHttpsTrigger(p.GetHttpsTrigger()),
		EventTrigger:               ProtoToCloudfunctionsAlphaFunctionEventTrigger(p.GetEventTrigger()),
		Status:                     ProtoToCloudfunctionsAlphaFunctionStatusEnum(p.GetStatus()),
		EntryPoint:                 dcl.StringOrNil(p.GetEntryPoint()),
		Runtime:                    dcl.StringOrNil(p.GetRuntime()),
		Timeout:                    dcl.StringOrNil(p.GetTimeout()),
		AvailableMemoryMb:          dcl.Int64OrNil(p.GetAvailableMemoryMb()),
		ServiceAccountEmail:        dcl.StringOrNil(p.GetServiceAccountEmail()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
		VersionId:                  dcl.Int64OrNil(p.GetVersionId()),
		MaxInstances:               dcl.Int64OrNil(p.GetMaxInstances()),
		VPCConnector:               dcl.StringOrNil(p.GetVpcConnector()),
		VPCConnectorEgressSettings: ProtoToCloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum(p.GetVpcConnectorEgressSettings()),
		IngressSettings:            ProtoToCloudfunctionsAlphaFunctionIngressSettingsEnum(p.GetIngressSettings()),
		Region:                     dcl.StringOrNil(p.GetRegion()),
		Project:                    dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// FunctionHttpsTriggerSecurityLevelEnumToProto converts a FunctionHttpsTriggerSecurityLevelEnum enum to its proto representation.
func CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnumToProto(e *alpha.FunctionHttpsTriggerSecurityLevelEnum) alphapb.CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum {
	if e == nil {
		return alphapb.CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum(0)
	}
	if v, ok := alphapb.CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum_value["FunctionHttpsTriggerSecurityLevelEnum"+string(*e)]; ok {
		return alphapb.CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum(v)
	}
	return alphapb.CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnum(0)
}

// FunctionStatusEnumToProto converts a FunctionStatusEnum enum to its proto representation.
func CloudfunctionsAlphaFunctionStatusEnumToProto(e *alpha.FunctionStatusEnum) alphapb.CloudfunctionsAlphaFunctionStatusEnum {
	if e == nil {
		return alphapb.CloudfunctionsAlphaFunctionStatusEnum(0)
	}
	if v, ok := alphapb.CloudfunctionsAlphaFunctionStatusEnum_value["FunctionStatusEnum"+string(*e)]; ok {
		return alphapb.CloudfunctionsAlphaFunctionStatusEnum(v)
	}
	return alphapb.CloudfunctionsAlphaFunctionStatusEnum(0)
}

// FunctionVPCConnectorEgressSettingsEnumToProto converts a FunctionVPCConnectorEgressSettingsEnum enum to its proto representation.
func CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnumToProto(e *alpha.FunctionVPCConnectorEgressSettingsEnum) alphapb.CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum {
	if e == nil {
		return alphapb.CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum(0)
	}
	if v, ok := alphapb.CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum_value["FunctionVPCConnectorEgressSettingsEnum"+string(*e)]; ok {
		return alphapb.CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum(v)
	}
	return alphapb.CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnum(0)
}

// FunctionIngressSettingsEnumToProto converts a FunctionIngressSettingsEnum enum to its proto representation.
func CloudfunctionsAlphaFunctionIngressSettingsEnumToProto(e *alpha.FunctionIngressSettingsEnum) alphapb.CloudfunctionsAlphaFunctionIngressSettingsEnum {
	if e == nil {
		return alphapb.CloudfunctionsAlphaFunctionIngressSettingsEnum(0)
	}
	if v, ok := alphapb.CloudfunctionsAlphaFunctionIngressSettingsEnum_value["FunctionIngressSettingsEnum"+string(*e)]; ok {
		return alphapb.CloudfunctionsAlphaFunctionIngressSettingsEnum(v)
	}
	return alphapb.CloudfunctionsAlphaFunctionIngressSettingsEnum(0)
}

// FunctionSourceRepositoryToProto converts a FunctionSourceRepository object to its proto representation.
func CloudfunctionsAlphaFunctionSourceRepositoryToProto(o *alpha.FunctionSourceRepository) *alphapb.CloudfunctionsAlphaFunctionSourceRepository {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudfunctionsAlphaFunctionSourceRepository{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetDeployedUrl(dcl.ValueOrEmptyString(o.DeployedUrl))
	return p
}

// FunctionHttpsTriggerToProto converts a FunctionHttpsTrigger object to its proto representation.
func CloudfunctionsAlphaFunctionHttpsTriggerToProto(o *alpha.FunctionHttpsTrigger) *alphapb.CloudfunctionsAlphaFunctionHttpsTrigger {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudfunctionsAlphaFunctionHttpsTrigger{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetSecurityLevel(CloudfunctionsAlphaFunctionHttpsTriggerSecurityLevelEnumToProto(o.SecurityLevel))
	return p
}

// FunctionEventTriggerToProto converts a FunctionEventTrigger object to its proto representation.
func CloudfunctionsAlphaFunctionEventTriggerToProto(o *alpha.FunctionEventTrigger) *alphapb.CloudfunctionsAlphaFunctionEventTrigger {
	if o == nil {
		return nil
	}
	p := &alphapb.CloudfunctionsAlphaFunctionEventTrigger{}
	p.SetEventType(dcl.ValueOrEmptyString(o.EventType))
	p.SetResource(dcl.ValueOrEmptyString(o.Resource))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetFailurePolicy(dcl.ValueOrEmptyBool(o.FailurePolicy))
	return p
}

// FunctionToProto converts a Function resource to its proto representation.
func FunctionToProto(resource *alpha.Function) *alphapb.CloudfunctionsAlphaFunction {
	p := &alphapb.CloudfunctionsAlphaFunction{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSourceArchiveUrl(dcl.ValueOrEmptyString(resource.SourceArchiveUrl))
	p.SetSourceRepository(CloudfunctionsAlphaFunctionSourceRepositoryToProto(resource.SourceRepository))
	p.SetHttpsTrigger(CloudfunctionsAlphaFunctionHttpsTriggerToProto(resource.HttpsTrigger))
	p.SetEventTrigger(CloudfunctionsAlphaFunctionEventTriggerToProto(resource.EventTrigger))
	p.SetStatus(CloudfunctionsAlphaFunctionStatusEnumToProto(resource.Status))
	p.SetEntryPoint(dcl.ValueOrEmptyString(resource.EntryPoint))
	p.SetRuntime(dcl.ValueOrEmptyString(resource.Runtime))
	p.SetTimeout(dcl.ValueOrEmptyString(resource.Timeout))
	p.SetAvailableMemoryMb(dcl.ValueOrEmptyInt64(resource.AvailableMemoryMb))
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(resource.ServiceAccountEmail))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetVersionId(dcl.ValueOrEmptyInt64(resource.VersionId))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(resource.MaxInstances))
	p.SetVpcConnector(dcl.ValueOrEmptyString(resource.VPCConnector))
	p.SetVpcConnectorEgressSettings(CloudfunctionsAlphaFunctionVPCConnectorEgressSettingsEnumToProto(resource.VPCConnectorEgressSettings))
	p.SetIngressSettings(CloudfunctionsAlphaFunctionIngressSettingsEnumToProto(resource.IngressSettings))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	mEnvironmentVariables := make(map[string]string, len(resource.EnvironmentVariables))
	for k, r := range resource.EnvironmentVariables {
		mEnvironmentVariables[k] = r
	}
	p.SetEnvironmentVariables(mEnvironmentVariables)

	return p
}

// applyFunction handles the gRPC request by passing it to the underlying Function Apply() method.
func (s *FunctionServer) applyFunction(ctx context.Context, c *alpha.Client, request *alphapb.ApplyCloudfunctionsAlphaFunctionRequest) (*alphapb.CloudfunctionsAlphaFunction, error) {
	p := ProtoToFunction(request.GetResource())
	res, err := c.ApplyFunction(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FunctionToProto(res)
	return r, nil
}

// applyCloudfunctionsAlphaFunction handles the gRPC request by passing it to the underlying Function Apply() method.
func (s *FunctionServer) ApplyCloudfunctionsAlphaFunction(ctx context.Context, request *alphapb.ApplyCloudfunctionsAlphaFunctionRequest) (*alphapb.CloudfunctionsAlphaFunction, error) {
	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFunction(ctx, cl, request)
}

// DeleteFunction handles the gRPC request by passing it to the underlying Function Delete() method.
func (s *FunctionServer) DeleteCloudfunctionsAlphaFunction(ctx context.Context, request *alphapb.DeleteCloudfunctionsAlphaFunctionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFunction(ctx, ProtoToFunction(request.GetResource()))

}

// ListCloudfunctionsAlphaFunction handles the gRPC request by passing it to the underlying FunctionList() method.
func (s *FunctionServer) ListCloudfunctionsAlphaFunction(ctx context.Context, request *alphapb.ListCloudfunctionsAlphaFunctionRequest) (*alphapb.ListCloudfunctionsAlphaFunctionResponse, error) {
	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFunction(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.CloudfunctionsAlphaFunction
	for _, r := range resources.Items {
		rp := FunctionToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListCloudfunctionsAlphaFunctionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFunction(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
