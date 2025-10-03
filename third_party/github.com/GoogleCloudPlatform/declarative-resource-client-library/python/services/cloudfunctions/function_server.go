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
	cloudfunctionspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudfunctions/cloudfunctions_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudfunctions"
)

// FunctionServer implements the gRPC interface for Function.
type FunctionServer struct{}

// ProtoToFunctionHttpsTriggerSecurityLevelEnum converts a FunctionHttpsTriggerSecurityLevelEnum enum from its proto representation.
func ProtoToCloudfunctionsFunctionHttpsTriggerSecurityLevelEnum(e cloudfunctionspb.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum) *cloudfunctions.FunctionHttpsTriggerSecurityLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum_name[int32(e)]; ok {
		e := cloudfunctions.FunctionHttpsTriggerSecurityLevelEnum(n[len("CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionStatusEnum converts a FunctionStatusEnum enum from its proto representation.
func ProtoToCloudfunctionsFunctionStatusEnum(e cloudfunctionspb.CloudfunctionsFunctionStatusEnum) *cloudfunctions.FunctionStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsFunctionStatusEnum_name[int32(e)]; ok {
		e := cloudfunctions.FunctionStatusEnum(n[len("CloudfunctionsFunctionStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionVPCConnectorEgressSettingsEnum converts a FunctionVPCConnectorEgressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsFunctionVPCConnectorEgressSettingsEnum(e cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum) *cloudfunctions.FunctionVPCConnectorEgressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum_name[int32(e)]; ok {
		e := cloudfunctions.FunctionVPCConnectorEgressSettingsEnum(n[len("CloudfunctionsFunctionVPCConnectorEgressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionIngressSettingsEnum converts a FunctionIngressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsFunctionIngressSettingsEnum(e cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum) *cloudfunctions.FunctionIngressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum_name[int32(e)]; ok {
		e := cloudfunctions.FunctionIngressSettingsEnum(n[len("CloudfunctionsFunctionIngressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFunctionSourceRepository converts a FunctionSourceRepository object from its proto representation.
func ProtoToCloudfunctionsFunctionSourceRepository(p *cloudfunctionspb.CloudfunctionsFunctionSourceRepository) *cloudfunctions.FunctionSourceRepository {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.FunctionSourceRepository{
		Url:         dcl.StringOrNil(p.GetUrl()),
		DeployedUrl: dcl.StringOrNil(p.GetDeployedUrl()),
	}
	return obj
}

// ProtoToFunctionHttpsTrigger converts a FunctionHttpsTrigger object from its proto representation.
func ProtoToCloudfunctionsFunctionHttpsTrigger(p *cloudfunctionspb.CloudfunctionsFunctionHttpsTrigger) *cloudfunctions.FunctionHttpsTrigger {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.FunctionHttpsTrigger{
		Url:           dcl.StringOrNil(p.GetUrl()),
		SecurityLevel: ProtoToCloudfunctionsFunctionHttpsTriggerSecurityLevelEnum(p.GetSecurityLevel()),
	}
	return obj
}

// ProtoToFunctionEventTrigger converts a FunctionEventTrigger object from its proto representation.
func ProtoToCloudfunctionsFunctionEventTrigger(p *cloudfunctionspb.CloudfunctionsFunctionEventTrigger) *cloudfunctions.FunctionEventTrigger {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.FunctionEventTrigger{
		EventType:     dcl.StringOrNil(p.GetEventType()),
		Resource:      dcl.StringOrNil(p.GetResource()),
		Service:       dcl.StringOrNil(p.GetService()),
		FailurePolicy: dcl.Bool(p.GetFailurePolicy()),
	}
	return obj
}

// ProtoToFunction converts a Function resource from its proto representation.
func ProtoToFunction(p *cloudfunctionspb.CloudfunctionsFunction) *cloudfunctions.Function {
	obj := &cloudfunctions.Function{
		Name:                       dcl.StringOrNil(p.GetName()),
		Description:                dcl.StringOrNil(p.GetDescription()),
		SourceArchiveUrl:           dcl.StringOrNil(p.GetSourceArchiveUrl()),
		SourceRepository:           ProtoToCloudfunctionsFunctionSourceRepository(p.GetSourceRepository()),
		HttpsTrigger:               ProtoToCloudfunctionsFunctionHttpsTrigger(p.GetHttpsTrigger()),
		EventTrigger:               ProtoToCloudfunctionsFunctionEventTrigger(p.GetEventTrigger()),
		Status:                     ProtoToCloudfunctionsFunctionStatusEnum(p.GetStatus()),
		EntryPoint:                 dcl.StringOrNil(p.GetEntryPoint()),
		Runtime:                    dcl.StringOrNil(p.GetRuntime()),
		Timeout:                    dcl.StringOrNil(p.GetTimeout()),
		AvailableMemoryMb:          dcl.Int64OrNil(p.GetAvailableMemoryMb()),
		ServiceAccountEmail:        dcl.StringOrNil(p.GetServiceAccountEmail()),
		UpdateTime:                 dcl.StringOrNil(p.GetUpdateTime()),
		VersionId:                  dcl.Int64OrNil(p.GetVersionId()),
		MaxInstances:               dcl.Int64OrNil(p.GetMaxInstances()),
		VPCConnector:               dcl.StringOrNil(p.GetVpcConnector()),
		VPCConnectorEgressSettings: ProtoToCloudfunctionsFunctionVPCConnectorEgressSettingsEnum(p.GetVpcConnectorEgressSettings()),
		IngressSettings:            ProtoToCloudfunctionsFunctionIngressSettingsEnum(p.GetIngressSettings()),
		Region:                     dcl.StringOrNil(p.GetRegion()),
		Project:                    dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// FunctionHttpsTriggerSecurityLevelEnumToProto converts a FunctionHttpsTriggerSecurityLevelEnum enum to its proto representation.
func CloudfunctionsFunctionHttpsTriggerSecurityLevelEnumToProto(e *cloudfunctions.FunctionHttpsTriggerSecurityLevelEnum) cloudfunctionspb.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum_value["FunctionHttpsTriggerSecurityLevelEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsFunctionHttpsTriggerSecurityLevelEnum(0)
}

// FunctionStatusEnumToProto converts a FunctionStatusEnum enum to its proto representation.
func CloudfunctionsFunctionStatusEnumToProto(e *cloudfunctions.FunctionStatusEnum) cloudfunctionspb.CloudfunctionsFunctionStatusEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsFunctionStatusEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsFunctionStatusEnum_value["FunctionStatusEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsFunctionStatusEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsFunctionStatusEnum(0)
}

// FunctionVPCConnectorEgressSettingsEnumToProto converts a FunctionVPCConnectorEgressSettingsEnum enum to its proto representation.
func CloudfunctionsFunctionVPCConnectorEgressSettingsEnumToProto(e *cloudfunctions.FunctionVPCConnectorEgressSettingsEnum) cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum_value["FunctionVPCConnectorEgressSettingsEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsFunctionVPCConnectorEgressSettingsEnum(0)
}

// FunctionIngressSettingsEnumToProto converts a FunctionIngressSettingsEnum enum to its proto representation.
func CloudfunctionsFunctionIngressSettingsEnumToProto(e *cloudfunctions.FunctionIngressSettingsEnum) cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum_value["FunctionIngressSettingsEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsFunctionIngressSettingsEnum(0)
}

// FunctionSourceRepositoryToProto converts a FunctionSourceRepository object to its proto representation.
func CloudfunctionsFunctionSourceRepositoryToProto(o *cloudfunctions.FunctionSourceRepository) *cloudfunctionspb.CloudfunctionsFunctionSourceRepository {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsFunctionSourceRepository{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetDeployedUrl(dcl.ValueOrEmptyString(o.DeployedUrl))
	return p
}

// FunctionHttpsTriggerToProto converts a FunctionHttpsTrigger object to its proto representation.
func CloudfunctionsFunctionHttpsTriggerToProto(o *cloudfunctions.FunctionHttpsTrigger) *cloudfunctionspb.CloudfunctionsFunctionHttpsTrigger {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsFunctionHttpsTrigger{}
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	p.SetSecurityLevel(CloudfunctionsFunctionHttpsTriggerSecurityLevelEnumToProto(o.SecurityLevel))
	return p
}

// FunctionEventTriggerToProto converts a FunctionEventTrigger object to its proto representation.
func CloudfunctionsFunctionEventTriggerToProto(o *cloudfunctions.FunctionEventTrigger) *cloudfunctionspb.CloudfunctionsFunctionEventTrigger {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsFunctionEventTrigger{}
	p.SetEventType(dcl.ValueOrEmptyString(o.EventType))
	p.SetResource(dcl.ValueOrEmptyString(o.Resource))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetFailurePolicy(dcl.ValueOrEmptyBool(o.FailurePolicy))
	return p
}

// FunctionToProto converts a Function resource to its proto representation.
func FunctionToProto(resource *cloudfunctions.Function) *cloudfunctionspb.CloudfunctionsFunction {
	p := &cloudfunctionspb.CloudfunctionsFunction{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetSourceArchiveUrl(dcl.ValueOrEmptyString(resource.SourceArchiveUrl))
	p.SetSourceRepository(CloudfunctionsFunctionSourceRepositoryToProto(resource.SourceRepository))
	p.SetHttpsTrigger(CloudfunctionsFunctionHttpsTriggerToProto(resource.HttpsTrigger))
	p.SetEventTrigger(CloudfunctionsFunctionEventTriggerToProto(resource.EventTrigger))
	p.SetStatus(CloudfunctionsFunctionStatusEnumToProto(resource.Status))
	p.SetEntryPoint(dcl.ValueOrEmptyString(resource.EntryPoint))
	p.SetRuntime(dcl.ValueOrEmptyString(resource.Runtime))
	p.SetTimeout(dcl.ValueOrEmptyString(resource.Timeout))
	p.SetAvailableMemoryMb(dcl.ValueOrEmptyInt64(resource.AvailableMemoryMb))
	p.SetServiceAccountEmail(dcl.ValueOrEmptyString(resource.ServiceAccountEmail))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetVersionId(dcl.ValueOrEmptyInt64(resource.VersionId))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(resource.MaxInstances))
	p.SetVpcConnector(dcl.ValueOrEmptyString(resource.VPCConnector))
	p.SetVpcConnectorEgressSettings(CloudfunctionsFunctionVPCConnectorEgressSettingsEnumToProto(resource.VPCConnectorEgressSettings))
	p.SetIngressSettings(CloudfunctionsFunctionIngressSettingsEnumToProto(resource.IngressSettings))
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
func (s *FunctionServer) applyFunction(ctx context.Context, c *cloudfunctions.Client, request *cloudfunctionspb.ApplyCloudfunctionsFunctionRequest) (*cloudfunctionspb.CloudfunctionsFunction, error) {
	p := ProtoToFunction(request.GetResource())
	res, err := c.ApplyFunction(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FunctionToProto(res)
	return r, nil
}

// applyCloudfunctionsFunction handles the gRPC request by passing it to the underlying Function Apply() method.
func (s *FunctionServer) ApplyCloudfunctionsFunction(ctx context.Context, request *cloudfunctionspb.ApplyCloudfunctionsFunctionRequest) (*cloudfunctionspb.CloudfunctionsFunction, error) {
	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFunction(ctx, cl, request)
}

// DeleteFunction handles the gRPC request by passing it to the underlying Function Delete() method.
func (s *FunctionServer) DeleteCloudfunctionsFunction(ctx context.Context, request *cloudfunctionspb.DeleteCloudfunctionsFunctionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFunction(ctx, ProtoToFunction(request.GetResource()))

}

// ListCloudfunctionsFunction handles the gRPC request by passing it to the underlying FunctionList() method.
func (s *FunctionServer) ListCloudfunctionsFunction(ctx context.Context, request *cloudfunctionspb.ListCloudfunctionsFunctionRequest) (*cloudfunctionspb.ListCloudfunctionsFunctionResponse, error) {
	cl, err := createConfigFunction(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFunction(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*cloudfunctionspb.CloudfunctionsFunction
	for _, r := range resources.Items {
		rp := FunctionToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudfunctionspb.ListCloudfunctionsFunctionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFunction(ctx context.Context, service_account_file string) (*cloudfunctions.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudfunctions.NewClient(conf), nil
}
