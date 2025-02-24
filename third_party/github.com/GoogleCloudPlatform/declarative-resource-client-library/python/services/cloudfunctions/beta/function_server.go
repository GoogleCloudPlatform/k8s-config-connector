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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudfunctions/beta/cloudfunctions_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions/beta"
)

// FunctionServer implements the gRPC interface for Function.
type FunctionServer struct{}

// ProtoToFunctionHttpsTriggerSecurityLevelEnum converts a FunctionHttpsTriggerSecurityLevelEnum enum from its proto representation.
func ProtoToCloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum(e betapb.CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum) *beta.FunctionHttpsTriggerSecurityLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum_name[int32(e)]; ok {
		e := beta.FunctionHttpsTriggerSecurityLevelEnum(n[len("CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionStatusEnum converts a FunctionStatusEnum enum from its proto representation.
func ProtoToCloudfunctionsBetaFunctionStatusEnum(e betapb.CloudfunctionsBetaFunctionStatusEnum) *beta.FunctionStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudfunctionsBetaFunctionStatusEnum_name[int32(e)]; ok {
		e := beta.FunctionStatusEnum(n[len("CloudfunctionsBetaFunctionStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionVPCConnectorEgressSettingsEnum converts a FunctionVPCConnectorEgressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum(e betapb.CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum) *beta.FunctionVPCConnectorEgressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum_name[int32(e)]; ok {
		e := beta.FunctionVPCConnectorEgressSettingsEnum(n[len("CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionIngressSettingsEnum converts a FunctionIngressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsBetaFunctionIngressSettingsEnum(e betapb.CloudfunctionsBetaFunctionIngressSettingsEnum) *beta.FunctionIngressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudfunctionsBetaFunctionIngressSettingsEnum_name[int32(e)]; ok {
		e := beta.FunctionIngressSettingsEnum(n[len("CloudfunctionsBetaFunctionIngressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionSourceRepository converts a FunctionSourceRepository object from its proto representation.
func ProtoToCloudfunctionsBetaFunctionSourceRepository(p *betapb.CloudfunctionsBetaFunctionSourceRepository) *beta.FunctionSourceRepository {
	if p == nil {
		return nil
	}
	obj := &beta.FunctionSourceRepository{
		Url:         dcl.StringOrNil(p.GetUrl()),
		DeployedUrl: dcl.StringOrNil(p.GetDeployedUrl()),
	}
	return obj
}

// ProtoToFunctionHttpsTrigger converts a FunctionHttpsTrigger object from its proto representation.
func ProtoToCloudfunctionsBetaFunctionHttpsTrigger(p *betapb.CloudfunctionsBetaFunctionHttpsTrigger) *beta.FunctionHttpsTrigger {
	if p == nil {
		return nil
	}
	obj := &beta.FunctionHttpsTrigger{
		Url:           dcl.StringOrNil(p.GetUrl()),
		SecurityLevel: ProtoToCloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum(p.GetSecurityLevel()),
	}
	return obj
}

// ProtoToFunctionEventTrigger converts a FunctionEventTrigger object from its proto representation.
func ProtoToCloudfunctionsBetaFunctionEventTrigger(p *betapb.CloudfunctionsBetaFunctionEventTrigger) *beta.FunctionEventTrigger {
	if p == nil {
		return nil
	}
	obj := &beta.FunctionEventTrigger{
		EventType:     dcl.StringOrNil(p.GetEventType()),
		Resource:      dcl.StringOrNil(p.GetResource()),
		Service:       dcl.StringOrNil(p.GetService()),
		FailurePolicy: dcl.Bool(p.GetFailurePolicy()),
	}
	return obj
}

// ProtoToFunction converts a Function resource from its proto representation.
func ProtoToFunction(p *betapb.CloudfunctionsBetaFunction) *beta.Function {
	obj := &beta.Function{
		Name:                       dcl.StringOrNil(p.GetName()),
		Description:                dcl.StringOrNil(p.GetDescription()),
		SourceArchiveUrl:           dcl.StringOrNil(p.GetSourceArchiveUrl()),
		SourceRepository:           ProtoToCloudfunctionsBetaFunctionSourceRepository(p.GetSourceRepository()),
		HttpsTrigger:               ProtoToCloudfunctionsBetaFunctionHttpsTrigger(p.GetHttpsTrigger()),
		EventTrigger:               ProtoToCloudfunctionsBetaFunctionEventTrigger(p.GetEventTrigger()),
		Status:                     ProtoToCloudfunctionsBetaFunctionStatusEnum(p.GetStatus()),
		EntryPoint:                 dcl.StringOrNil(p.GetEntryPoint()),
		Runtime:                    dcl.StringOrNil(p.GetRuntime()),
		Timeout:                    dcl.StringOrNil(p.GetTimeout()),
		AvailableMemoryMb:          dcl.Int64OrNil(p.GetAvailableMemoryMb()),
		ServiceAccountEmail:        dcl.StringOrNil(p.GetServiceAccountEmail()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
		VersionId:                  dcl.Int64OrNil(p.GetVersionId()),
		MaxInstances:               dcl.Int64OrNil(p.GetMaxInstances()),
		VPCConnector:               dcl.StringOrNil(p.GetVpcConnector()),
		VPCConnectorEgressSettings: ProtoToCloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum(p.GetVpcConnectorEgressSettings()),
		IngressSettings:            ProtoToCloudfunctionsBetaFunctionIngressSettingsEnum(p.GetIngressSettings()),
		Region:                     dcl.StringOrNil(p.GetRegion()),
		Project:                    dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// FunctionHttpsTriggerSecurityLevelEnumToProto converts a FunctionHttpsTriggerSecurityLevelEnum enum to its proto representation.
func CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnumToProto(e *beta.FunctionHttpsTriggerSecurityLevelEnum) betapb.CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum {
	if e == nil {
		return betapb.CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum(0)
	}
	if v, ok := betapb.CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum_value["FunctionHttpsTriggerSecurityLevelEnum"+string(*e)]; ok {
		return betapb.CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum(v)
	}
	return betapb.CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnum(0)
}

// FunctionStatusEnumToProto converts a FunctionStatusEnum enum to its proto representation.
func CloudfunctionsBetaFunctionStatusEnumToProto(e *beta.FunctionStatusEnum) betapb.CloudfunctionsBetaFunctionStatusEnum {
	if e == nil {
		return betapb.CloudfunctionsBetaFunctionStatusEnum(0)
	}
	if v, ok := betapb.CloudfunctionsBetaFunctionStatusEnum_value["FunctionStatusEnum"+string(*e)]; ok {
		return betapb.CloudfunctionsBetaFunctionStatusEnum(v)
	}
	return betapb.CloudfunctionsBetaFunctionStatusEnum(0)
}

// FunctionVPCConnectorEgressSettingsEnumToProto converts a FunctionVPCConnectorEgressSettingsEnum enum to its proto representation.
func CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnumToProto(e *beta.FunctionVPCConnectorEgressSettingsEnum) betapb.CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum {
	if e == nil {
		return betapb.CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum(0)
	}
	if v, ok := betapb.CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum_value["FunctionVPCConnectorEgressSettingsEnum"+string(*e)]; ok {
		return betapb.CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum(v)
	}
	return betapb.CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnum(0)
}

// FunctionIngressSettingsEnumToProto converts a FunctionIngressSettingsEnum enum to its proto representation.
func CloudfunctionsBetaFunctionIngressSettingsEnumToProto(e *beta.FunctionIngressSettingsEnum) betapb.CloudfunctionsBetaFunctionIngressSettingsEnum {
	if e == nil {
		return betapb.CloudfunctionsBetaFunctionIngressSettingsEnum(0)
	}
	if v, ok := betapb.CloudfunctionsBetaFunctionIngressSettingsEnum_value["FunctionIngressSettingsEnum"+string(*e)]; ok {
		return betapb.CloudfunctionsBetaFunctionIngressSettingsEnum(v)
	}
	return betapb.CloudfunctionsBetaFunctionIngressSettingsEnum(0)
}

// FunctionSourceRepositoryToProto converts a FunctionSourceRepository object to its proto representation.
func CloudfunctionsBetaFunctionSourceRepositoryToProto(o *beta.FunctionSourceRepository) *betapb.CloudfunctionsBetaFunctionSourceRepository {
	if o == nil {
		return nil
	}
	p := &betapb.CloudfunctionsBetaFunctionSourceRepository{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetDeployedUrl(dcl.ValueOrEmptyString(o.DeployedUrl))
	return p
}

// FunctionHttpsTriggerToProto converts a FunctionHttpsTrigger object to its proto representation.
func CloudfunctionsBetaFunctionHttpsTriggerToProto(o *beta.FunctionHttpsTrigger) *betapb.CloudfunctionsBetaFunctionHttpsTrigger {
	if o == nil {
		return nil
	}
	p := &betapb.CloudfunctionsBetaFunctionHttpsTrigger{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetSecurityLevel(CloudfunctionsBetaFunctionHttpsTriggerSecurityLevelEnumToProto(o.SecurityLevel))
	return p
}

// FunctionEventTriggerToProto converts a FunctionEventTrigger object to its proto representation.
func CloudfunctionsBetaFunctionEventTriggerToProto(o *beta.FunctionEventTrigger) *betapb.CloudfunctionsBetaFunctionEventTrigger {
	if o == nil {
		return nil
	}
	p := &betapb.CloudfunctionsBetaFunctionEventTrigger{}
	p.SetEventType(dcl.ValueOrEmptyString(o.EventType))
	p.SetResource(dcl.ValueOrEmptyString(o.Resource))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetFailurePolicy(dcl.ValueOrEmptyBool(o.FailurePolicy))
	return p
}

// FunctionToProto converts a Function resource to its proto representation.
func FunctionToProto(resource *beta.Function) *betapb.CloudfunctionsBetaFunction {
	p := &betapb.CloudfunctionsBetaFunction{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSourceArchiveUrl(dcl.ValueOrEmptyString(resource.SourceArchiveUrl))
	p.SetSourceRepository(CloudfunctionsBetaFunctionSourceRepositoryToProto(resource.SourceRepository))
	p.SetHttpsTrigger(CloudfunctionsBetaFunctionHttpsTriggerToProto(resource.HttpsTrigger))
	p.SetEventTrigger(CloudfunctionsBetaFunctionEventTriggerToProto(resource.EventTrigger))
	p.SetStatus(CloudfunctionsBetaFunctionStatusEnumToProto(resource.Status))
	p.SetEntryPoint(dcl.ValueOrEmptyString(resource.EntryPoint))
	p.SetRuntime(dcl.ValueOrEmptyString(resource.Runtime))
	p.SetTimeout(dcl.ValueOrEmptyString(resource.Timeout))
	p.SetAvailableMemoryMb(dcl.ValueOrEmptyInt64(resource.AvailableMemoryMb))
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(resource.ServiceAccountEmail))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetVersionId(dcl.ValueOrEmptyInt64(resource.VersionId))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(resource.MaxInstances))
	p.SetVpcConnector(dcl.ValueOrEmptyString(resource.VPCConnector))
	p.SetVpcConnectorEgressSettings(CloudfunctionsBetaFunctionVPCConnectorEgressSettingsEnumToProto(resource.VPCConnectorEgressSettings))
	p.SetIngressSettings(CloudfunctionsBetaFunctionIngressSettingsEnumToProto(resource.IngressSettings))
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
func (s *FunctionServer) applyFunction(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudfunctionsBetaFunctionRequest) (*betapb.CloudfunctionsBetaFunction, error) {
	p := ProtoToFunction(request.GetResource())
	res, err := c.ApplyFunction(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FunctionToProto(res)
	return r, nil
}

// applyCloudfunctionsBetaFunction handles the gRPC request by passing it to the underlying Function Apply() method.
func (s *FunctionServer) ApplyCloudfunctionsBetaFunction(ctx context.Context, request *betapb.ApplyCloudfunctionsBetaFunctionRequest) (*betapb.CloudfunctionsBetaFunction, error) {
	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFunction(ctx, cl, request)
}

// DeleteFunction handles the gRPC request by passing it to the underlying Function Delete() method.
func (s *FunctionServer) DeleteCloudfunctionsBetaFunction(ctx context.Context, request *betapb.DeleteCloudfunctionsBetaFunctionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFunction(ctx, ProtoToFunction(request.GetResource()))

}

// ListCloudfunctionsBetaFunction handles the gRPC request by passing it to the underlying FunctionList() method.
func (s *FunctionServer) ListCloudfunctionsBetaFunction(ctx context.Context, request *betapb.ListCloudfunctionsBetaFunctionRequest) (*betapb.ListCloudfunctionsBetaFunctionResponse, error) {
	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFunction(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudfunctionsBetaFunction
	for _, r := range resources.Items {
		rp := FunctionToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudfunctionsBetaFunctionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFunction(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
