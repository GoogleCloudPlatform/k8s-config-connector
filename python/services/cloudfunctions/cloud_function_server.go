// Copyright 2021 Google LLC. All Rights Reserved.
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

// Server implements the gRPC interface for CloudFunction.
type CloudFunctionServer struct{}

// ProtoToCloudFunctionStatusEnum converts a CloudFunctionStatusEnum enum from its proto representation.
func ProtoToCloudfunctionsCloudFunctionStatusEnum(e cloudfunctionspb.CloudfunctionsCloudFunctionStatusEnum) *cloudfunctions.CloudFunctionStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsCloudFunctionStatusEnum_name[int32(e)]; ok {
		e := cloudfunctions.CloudFunctionStatusEnum(n[len("CloudFunctionStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToCloudFunctionVPCConnectorEgressSettingsEnum converts a CloudFunctionVPCConnectorEgressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum(e cloudfunctionspb.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum) *cloudfunctions.CloudFunctionVPCConnectorEgressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum_name[int32(e)]; ok {
		e := cloudfunctions.CloudFunctionVPCConnectorEgressSettingsEnum(n[len("CloudFunctionVPCConnectorEgressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCloudFunctionIngressSettingsEnum converts a CloudFunctionIngressSettingsEnum enum from its proto representation.
func ProtoToCloudfunctionsCloudFunctionIngressSettingsEnum(e cloudfunctionspb.CloudfunctionsCloudFunctionIngressSettingsEnum) *cloudfunctions.CloudFunctionIngressSettingsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudfunctionspb.CloudfunctionsCloudFunctionIngressSettingsEnum_name[int32(e)]; ok {
		e := cloudfunctions.CloudFunctionIngressSettingsEnum(n[len("CloudFunctionIngressSettingsEnum"):])
		return &e
	}
	return nil
}

// ProtoToCloudFunctionSourceRepository converts a CloudFunctionSourceRepository resource from its proto representation.
func ProtoToCloudfunctionsCloudFunctionSourceRepository(p *cloudfunctionspb.CloudfunctionsCloudFunctionSourceRepository) *cloudfunctions.CloudFunctionSourceRepository {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.CloudFunctionSourceRepository{
		Url:         dcl.StringOrNil(p.Url),
		DeployedUrl: dcl.StringOrNil(p.DeployedUrl),
	}
	return obj
}

// ProtoToCloudFunctionHttpsTrigger converts a CloudFunctionHttpsTrigger resource from its proto representation.
func ProtoToCloudfunctionsCloudFunctionHttpsTrigger(p *cloudfunctionspb.CloudfunctionsCloudFunctionHttpsTrigger) *cloudfunctions.CloudFunctionHttpsTrigger {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.CloudFunctionHttpsTrigger{
		Url: dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToCloudFunctionEventTrigger converts a CloudFunctionEventTrigger resource from its proto representation.
func ProtoToCloudfunctionsCloudFunctionEventTrigger(p *cloudfunctionspb.CloudfunctionsCloudFunctionEventTrigger) *cloudfunctions.CloudFunctionEventTrigger {
	if p == nil {
		return nil
	}
	obj := &cloudfunctions.CloudFunctionEventTrigger{
		EventType:     dcl.StringOrNil(p.EventType),
		Resource:      dcl.StringOrNil(p.Resource),
		Service:       dcl.StringOrNil(p.Service),
		FailurePolicy: dcl.Bool(p.FailurePolicy),
	}
	return obj
}

// ProtoToCloudFunction converts a CloudFunction resource from its proto representation.
func ProtoToCloudFunction(p *cloudfunctionspb.CloudfunctionsCloudFunction) *cloudfunctions.CloudFunction {
	obj := &cloudfunctions.CloudFunction{
		Name:                       dcl.StringOrNil(p.Name),
		Description:                dcl.StringOrNil(p.Description),
		SourceArchiveUrl:           dcl.StringOrNil(p.SourceArchiveUrl),
		SourceRepository:           ProtoToCloudfunctionsCloudFunctionSourceRepository(p.GetSourceRepository()),
		HttpsTrigger:               ProtoToCloudfunctionsCloudFunctionHttpsTrigger(p.GetHttpsTrigger()),
		EventTrigger:               ProtoToCloudfunctionsCloudFunctionEventTrigger(p.GetEventTrigger()),
		Status:                     ProtoToCloudfunctionsCloudFunctionStatusEnum(p.GetStatus()),
		EntryPoint:                 dcl.StringOrNil(p.EntryPoint),
		Runtime:                    dcl.StringOrNil(p.Runtime),
		Timeout:                    dcl.Int64OrNil(p.Timeout),
		AvailableMemoryMb:          dcl.Int64OrNil(p.AvailableMemoryMb),
		ServiceAccountEmail:        dcl.StringOrNil(p.ServiceAccountEmail),
		UpdateTime:                 dcl.StringOrNil(p.UpdateTime),
		VersionId:                  dcl.Int64OrNil(p.VersionId),
		Network:                    dcl.StringOrNil(p.Network),
		MaxInstances:               dcl.Int64OrNil(p.MaxInstances),
		VPCConnector:               dcl.StringOrNil(p.VpcConnector),
		VPCConnectorEgressSettings: ProtoToCloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum(p.GetVpcConnectorEgressSettings()),
		IngressSettings:            ProtoToCloudfunctionsCloudFunctionIngressSettingsEnum(p.GetIngressSettings()),
		Region:                     dcl.StringOrNil(p.Region),
		Project:                    dcl.StringOrNil(p.Project),
	}
	return obj
}

// CloudFunctionStatusEnumToProto converts a CloudFunctionStatusEnum enum to its proto representation.
func CloudfunctionsCloudFunctionStatusEnumToProto(e *cloudfunctions.CloudFunctionStatusEnum) cloudfunctionspb.CloudfunctionsCloudFunctionStatusEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsCloudFunctionStatusEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsCloudFunctionStatusEnum_value["CloudFunctionStatusEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsCloudFunctionStatusEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsCloudFunctionStatusEnum(0)
}

// CloudFunctionVPCConnectorEgressSettingsEnumToProto converts a CloudFunctionVPCConnectorEgressSettingsEnum enum to its proto representation.
func CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnumToProto(e *cloudfunctions.CloudFunctionVPCConnectorEgressSettingsEnum) cloudfunctionspb.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum_value["CloudFunctionVPCConnectorEgressSettingsEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnum(0)
}

// CloudFunctionIngressSettingsEnumToProto converts a CloudFunctionIngressSettingsEnum enum to its proto representation.
func CloudfunctionsCloudFunctionIngressSettingsEnumToProto(e *cloudfunctions.CloudFunctionIngressSettingsEnum) cloudfunctionspb.CloudfunctionsCloudFunctionIngressSettingsEnum {
	if e == nil {
		return cloudfunctionspb.CloudfunctionsCloudFunctionIngressSettingsEnum(0)
	}
	if v, ok := cloudfunctionspb.CloudfunctionsCloudFunctionIngressSettingsEnum_value["CloudFunctionIngressSettingsEnum"+string(*e)]; ok {
		return cloudfunctionspb.CloudfunctionsCloudFunctionIngressSettingsEnum(v)
	}
	return cloudfunctionspb.CloudfunctionsCloudFunctionIngressSettingsEnum(0)
}

// CloudFunctionSourceRepositoryToProto converts a CloudFunctionSourceRepository resource to its proto representation.
func CloudfunctionsCloudFunctionSourceRepositoryToProto(o *cloudfunctions.CloudFunctionSourceRepository) *cloudfunctionspb.CloudfunctionsCloudFunctionSourceRepository {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsCloudFunctionSourceRepository{
		Url:         dcl.ValueOrEmptyString(o.Url),
		DeployedUrl: dcl.ValueOrEmptyString(o.DeployedUrl),
	}
	return p
}

// CloudFunctionHttpsTriggerToProto converts a CloudFunctionHttpsTrigger resource to its proto representation.
func CloudfunctionsCloudFunctionHttpsTriggerToProto(o *cloudfunctions.CloudFunctionHttpsTrigger) *cloudfunctionspb.CloudfunctionsCloudFunctionHttpsTrigger {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsCloudFunctionHttpsTrigger{
		Url: dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// CloudFunctionEventTriggerToProto converts a CloudFunctionEventTrigger resource to its proto representation.
func CloudfunctionsCloudFunctionEventTriggerToProto(o *cloudfunctions.CloudFunctionEventTrigger) *cloudfunctionspb.CloudfunctionsCloudFunctionEventTrigger {
	if o == nil {
		return nil
	}
	p := &cloudfunctionspb.CloudfunctionsCloudFunctionEventTrigger{
		EventType:     dcl.ValueOrEmptyString(o.EventType),
		Resource:      dcl.ValueOrEmptyString(o.Resource),
		Service:       dcl.ValueOrEmptyString(o.Service),
		FailurePolicy: dcl.ValueOrEmptyBool(o.FailurePolicy),
	}
	return p
}

// CloudFunctionToProto converts a CloudFunction resource to its proto representation.
func CloudFunctionToProto(resource *cloudfunctions.CloudFunction) *cloudfunctionspb.CloudfunctionsCloudFunction {
	p := &cloudfunctionspb.CloudfunctionsCloudFunction{
		Name:                       dcl.ValueOrEmptyString(resource.Name),
		Description:                dcl.ValueOrEmptyString(resource.Description),
		SourceArchiveUrl:           dcl.ValueOrEmptyString(resource.SourceArchiveUrl),
		SourceRepository:           CloudfunctionsCloudFunctionSourceRepositoryToProto(resource.SourceRepository),
		HttpsTrigger:               CloudfunctionsCloudFunctionHttpsTriggerToProto(resource.HttpsTrigger),
		EventTrigger:               CloudfunctionsCloudFunctionEventTriggerToProto(resource.EventTrigger),
		Status:                     CloudfunctionsCloudFunctionStatusEnumToProto(resource.Status),
		EntryPoint:                 dcl.ValueOrEmptyString(resource.EntryPoint),
		Runtime:                    dcl.ValueOrEmptyString(resource.Runtime),
		Timeout:                    dcl.ValueOrEmptyInt64(resource.Timeout),
		AvailableMemoryMb:          dcl.ValueOrEmptyInt64(resource.AvailableMemoryMb),
		ServiceAccountEmail:        dcl.ValueOrEmptyString(resource.ServiceAccountEmail),
		UpdateTime:                 dcl.ValueOrEmptyString(resource.UpdateTime),
		VersionId:                  dcl.ValueOrEmptyInt64(resource.VersionId),
		Network:                    dcl.ValueOrEmptyString(resource.Network),
		MaxInstances:               dcl.ValueOrEmptyInt64(resource.MaxInstances),
		VpcConnector:               dcl.ValueOrEmptyString(resource.VPCConnector),
		VpcConnectorEgressSettings: CloudfunctionsCloudFunctionVPCConnectorEgressSettingsEnumToProto(resource.VPCConnectorEgressSettings),
		IngressSettings:            CloudfunctionsCloudFunctionIngressSettingsEnumToProto(resource.IngressSettings),
		Region:                     dcl.ValueOrEmptyString(resource.Region),
		Project:                    dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyCloudFunction handles the gRPC request by passing it to the underlying CloudFunction Apply() method.
func (s *CloudFunctionServer) applyCloudFunction(ctx context.Context, c *cloudfunctions.Client, request *cloudfunctionspb.ApplyCloudfunctionsCloudFunctionRequest) (*cloudfunctionspb.CloudfunctionsCloudFunction, error) {
	p := ProtoToCloudFunction(request.GetResource())
	res, err := c.ApplyCloudFunction(ctx, p)
	if err != nil {
		return nil, err
	}
	r := CloudFunctionToProto(res)
	return r, nil
}

// ApplyCloudFunction handles the gRPC request by passing it to the underlying CloudFunction Apply() method.
func (s *CloudFunctionServer) ApplyCloudfunctionsCloudFunction(ctx context.Context, request *cloudfunctionspb.ApplyCloudfunctionsCloudFunctionRequest) (*cloudfunctionspb.CloudfunctionsCloudFunction, error) {
	cl, err := createConfigCloudFunction(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyCloudFunction(ctx, cl, request)
}

// DeleteCloudFunction handles the gRPC request by passing it to the underlying CloudFunction Delete() method.
func (s *CloudFunctionServer) DeleteCloudfunctionsCloudFunction(ctx context.Context, request *cloudfunctionspb.DeleteCloudfunctionsCloudFunctionRequest) (*emptypb.Empty, error) {
	cl, err := createConfigCloudFunction(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCloudFunction(ctx, ProtoToCloudFunction(request.GetResource()))
}

// ListCloudFunction handles the gRPC request by passing it to the underlying CloudFunctionList() method.
func (s *CloudFunctionServer) ListCloudfunctionsCloudFunction(ctx context.Context, request *cloudfunctionspb.ListCloudfunctionsCloudFunctionRequest) (*cloudfunctionspb.ListCloudfunctionsCloudFunctionResponse, error) {
	cl, err := createConfigCloudFunction(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCloudFunction(ctx, request.Project, request.Region)
	if err != nil {
		return nil, err
	}
	var protos []*cloudfunctionspb.CloudfunctionsCloudFunction
	for _, r := range resources.Items {
		rp := CloudFunctionToProto(r)
		protos = append(protos, rp)
	}
	return &cloudfunctionspb.ListCloudfunctionsCloudFunctionResponse{Items: protos}, nil
}

func createConfigCloudFunction(ctx context.Context, service_account_file string) (*cloudfunctions.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudfunctions.NewClient(conf), nil
}
