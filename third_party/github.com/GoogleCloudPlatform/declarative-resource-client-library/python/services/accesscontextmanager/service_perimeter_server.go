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
	accesscontextmanagerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/accesscontextmanager/accesscontextmanager_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/accesscontextmanager"
)

// Server implements the gRPC interface for ServicePerimeter.
type ServicePerimeterServer struct{}

// ProtoToServicePerimeterPerimeterTypeEnum converts a ServicePerimeterPerimeterTypeEnum enum from its proto representation.
func ProtoToAccesscontextmanagerServicePerimeterPerimeterTypeEnum(e accesscontextmanagerpb.AccesscontextmanagerServicePerimeterPerimeterTypeEnum) *accesscontextmanager.ServicePerimeterPerimeterTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := accesscontextmanagerpb.AccesscontextmanagerServicePerimeterPerimeterTypeEnum_name[int32(e)]; ok {
		e := accesscontextmanager.ServicePerimeterPerimeterTypeEnum(n[len("AccesscontextmanagerServicePerimeterPerimeterTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToServicePerimeterStatus converts a ServicePerimeterStatus resource from its proto representation.
func ProtoToAccesscontextmanagerServicePerimeterStatus(p *accesscontextmanagerpb.AccesscontextmanagerServicePerimeterStatus) *accesscontextmanager.ServicePerimeterStatus {
	if p == nil {
		return nil
	}
	obj := &accesscontextmanager.ServicePerimeterStatus{
		VPCAccessibleServices: ProtoToAccesscontextmanagerServicePerimeterStatusVPCAccessibleServices(p.GetVpcAccessibleServices()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, r)
	}
	for _, r := range p.GetAccessLevels() {
		obj.AccessLevels = append(obj.AccessLevels, r)
	}
	for _, r := range p.GetRestrictedServices() {
		obj.RestrictedServices = append(obj.RestrictedServices, r)
	}
	return obj
}

// ProtoToServicePerimeterStatusVPCAccessibleServices converts a ServicePerimeterStatusVPCAccessibleServices resource from its proto representation.
func ProtoToAccesscontextmanagerServicePerimeterStatusVPCAccessibleServices(p *accesscontextmanagerpb.AccesscontextmanagerServicePerimeterStatusVPCAccessibleServices) *accesscontextmanager.ServicePerimeterStatusVPCAccessibleServices {
	if p == nil {
		return nil
	}
	obj := &accesscontextmanager.ServicePerimeterStatusVPCAccessibleServices{
		EnableRestriction: dcl.Bool(p.EnableRestriction),
	}
	for _, r := range p.GetAllowedServices() {
		obj.AllowedServices = append(obj.AllowedServices, r)
	}
	return obj
}

// ProtoToServicePerimeterSpec converts a ServicePerimeterSpec resource from its proto representation.
func ProtoToAccesscontextmanagerServicePerimeterSpec(p *accesscontextmanagerpb.AccesscontextmanagerServicePerimeterSpec) *accesscontextmanager.ServicePerimeterSpec {
	if p == nil {
		return nil
	}
	obj := &accesscontextmanager.ServicePerimeterSpec{
		VPCAccessibleServices: ProtoToAccesscontextmanagerServicePerimeterSpecVPCAccessibleServices(p.GetVpcAccessibleServices()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, r)
	}
	for _, r := range p.GetAccessLevels() {
		obj.AccessLevels = append(obj.AccessLevels, r)
	}
	for _, r := range p.GetRestrictedServices() {
		obj.RestrictedServices = append(obj.RestrictedServices, r)
	}
	return obj
}

// ProtoToServicePerimeterSpecVPCAccessibleServices converts a ServicePerimeterSpecVPCAccessibleServices resource from its proto representation.
func ProtoToAccesscontextmanagerServicePerimeterSpecVPCAccessibleServices(p *accesscontextmanagerpb.AccesscontextmanagerServicePerimeterSpecVPCAccessibleServices) *accesscontextmanager.ServicePerimeterSpecVPCAccessibleServices {
	if p == nil {
		return nil
	}
	obj := &accesscontextmanager.ServicePerimeterSpecVPCAccessibleServices{
		EnableRestriction: dcl.Bool(p.EnableRestriction),
	}
	for _, r := range p.GetAllowedServices() {
		obj.AllowedServices = append(obj.AllowedServices, r)
	}
	return obj
}

// ProtoToServicePerimeter converts a ServicePerimeter resource from its proto representation.
func ProtoToServicePerimeter(p *accesscontextmanagerpb.AccesscontextmanagerServicePerimeter) *accesscontextmanager.ServicePerimeter {
	obj := &accesscontextmanager.ServicePerimeter{
		Title:                 dcl.StringOrNil(p.Title),
		Description:           dcl.StringOrNil(p.Description),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		PerimeterType:         ProtoToAccesscontextmanagerServicePerimeterPerimeterTypeEnum(p.GetPerimeterType()),
		Status:                ProtoToAccesscontextmanagerServicePerimeterStatus(p.GetStatus()),
		Policy:                dcl.StringOrNil(p.Policy),
		Name:                  dcl.StringOrNil(p.Name),
		UseExplicitDryRunSpec: dcl.Bool(p.UseExplicitDryRunSpec),
		Spec:                  ProtoToAccesscontextmanagerServicePerimeterSpec(p.GetSpec()),
	}
	return obj
}

// ServicePerimeterPerimeterTypeEnumToProto converts a ServicePerimeterPerimeterTypeEnum enum to its proto representation.
func AccesscontextmanagerServicePerimeterPerimeterTypeEnumToProto(e *accesscontextmanager.ServicePerimeterPerimeterTypeEnum) accesscontextmanagerpb.AccesscontextmanagerServicePerimeterPerimeterTypeEnum {
	if e == nil {
		return accesscontextmanagerpb.AccesscontextmanagerServicePerimeterPerimeterTypeEnum(0)
	}
	if v, ok := accesscontextmanagerpb.AccesscontextmanagerServicePerimeterPerimeterTypeEnum_value["ServicePerimeterPerimeterTypeEnum"+string(*e)]; ok {
		return accesscontextmanagerpb.AccesscontextmanagerServicePerimeterPerimeterTypeEnum(v)
	}
	return accesscontextmanagerpb.AccesscontextmanagerServicePerimeterPerimeterTypeEnum(0)
}

// ServicePerimeterStatusToProto converts a ServicePerimeterStatus resource to its proto representation.
func AccesscontextmanagerServicePerimeterStatusToProto(o *accesscontextmanager.ServicePerimeterStatus) *accesscontextmanagerpb.AccesscontextmanagerServicePerimeterStatus {
	if o == nil {
		return nil
	}
	p := &accesscontextmanagerpb.AccesscontextmanagerServicePerimeterStatus{
		VpcAccessibleServices: AccesscontextmanagerServicePerimeterStatusVPCAccessibleServicesToProto(o.VPCAccessibleServices),
	}
	for _, r := range o.Resources {
		p.Resources = append(p.Resources, r)
	}
	for _, r := range o.AccessLevels {
		p.AccessLevels = append(p.AccessLevels, r)
	}
	for _, r := range o.RestrictedServices {
		p.RestrictedServices = append(p.RestrictedServices, r)
	}
	return p
}

// ServicePerimeterStatusVPCAccessibleServicesToProto converts a ServicePerimeterStatusVPCAccessibleServices resource to its proto representation.
func AccesscontextmanagerServicePerimeterStatusVPCAccessibleServicesToProto(o *accesscontextmanager.ServicePerimeterStatusVPCAccessibleServices) *accesscontextmanagerpb.AccesscontextmanagerServicePerimeterStatusVPCAccessibleServices {
	if o == nil {
		return nil
	}
	p := &accesscontextmanagerpb.AccesscontextmanagerServicePerimeterStatusVPCAccessibleServices{
		EnableRestriction: dcl.ValueOrEmptyBool(o.EnableRestriction),
	}
	for _, r := range o.AllowedServices {
		p.AllowedServices = append(p.AllowedServices, r)
	}
	return p
}

// ServicePerimeterSpecToProto converts a ServicePerimeterSpec resource to its proto representation.
func AccesscontextmanagerServicePerimeterSpecToProto(o *accesscontextmanager.ServicePerimeterSpec) *accesscontextmanagerpb.AccesscontextmanagerServicePerimeterSpec {
	if o == nil {
		return nil
	}
	p := &accesscontextmanagerpb.AccesscontextmanagerServicePerimeterSpec{
		VpcAccessibleServices: AccesscontextmanagerServicePerimeterSpecVPCAccessibleServicesToProto(o.VPCAccessibleServices),
	}
	for _, r := range o.Resources {
		p.Resources = append(p.Resources, r)
	}
	for _, r := range o.AccessLevels {
		p.AccessLevels = append(p.AccessLevels, r)
	}
	for _, r := range o.RestrictedServices {
		p.RestrictedServices = append(p.RestrictedServices, r)
	}
	return p
}

// ServicePerimeterSpecVPCAccessibleServicesToProto converts a ServicePerimeterSpecVPCAccessibleServices resource to its proto representation.
func AccesscontextmanagerServicePerimeterSpecVPCAccessibleServicesToProto(o *accesscontextmanager.ServicePerimeterSpecVPCAccessibleServices) *accesscontextmanagerpb.AccesscontextmanagerServicePerimeterSpecVPCAccessibleServices {
	if o == nil {
		return nil
	}
	p := &accesscontextmanagerpb.AccesscontextmanagerServicePerimeterSpecVPCAccessibleServices{
		EnableRestriction: dcl.ValueOrEmptyBool(o.EnableRestriction),
	}
	for _, r := range o.AllowedServices {
		p.AllowedServices = append(p.AllowedServices, r)
	}
	return p
}

// ServicePerimeterToProto converts a ServicePerimeter resource to its proto representation.
func ServicePerimeterToProto(resource *accesscontextmanager.ServicePerimeter) *accesscontextmanagerpb.AccesscontextmanagerServicePerimeter {
	p := &accesscontextmanagerpb.AccesscontextmanagerServicePerimeter{
		Title:                 dcl.ValueOrEmptyString(resource.Title),
		Description:           dcl.ValueOrEmptyString(resource.Description),
		CreateTime:            dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:            dcl.ValueOrEmptyString(resource.UpdateTime),
		PerimeterType:         AccesscontextmanagerServicePerimeterPerimeterTypeEnumToProto(resource.PerimeterType),
		Status:                AccesscontextmanagerServicePerimeterStatusToProto(resource.Status),
		Policy:                dcl.ValueOrEmptyString(resource.Policy),
		Name:                  dcl.ValueOrEmptyString(resource.Name),
		UseExplicitDryRunSpec: dcl.ValueOrEmptyBool(resource.UseExplicitDryRunSpec),
		Spec:                  AccesscontextmanagerServicePerimeterSpecToProto(resource.Spec),
	}

	return p
}

// ApplyServicePerimeter handles the gRPC request by passing it to the underlying ServicePerimeter Apply() method.
func (s *ServicePerimeterServer) applyServicePerimeter(ctx context.Context, c *accesscontextmanager.Client, request *accesscontextmanagerpb.ApplyAccesscontextmanagerServicePerimeterRequest) (*accesscontextmanagerpb.AccesscontextmanagerServicePerimeter, error) {
	p := ProtoToServicePerimeter(request.GetResource())
	res, err := c.ApplyServicePerimeter(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServicePerimeterToProto(res)
	return r, nil
}

// ApplyServicePerimeter handles the gRPC request by passing it to the underlying ServicePerimeter Apply() method.
func (s *ServicePerimeterServer) ApplyAccesscontextmanagerServicePerimeter(ctx context.Context, request *accesscontextmanagerpb.ApplyAccesscontextmanagerServicePerimeterRequest) (*accesscontextmanagerpb.AccesscontextmanagerServicePerimeter, error) {
	cl, err := createConfigServicePerimeter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyServicePerimeter(ctx, cl, request)
}

// DeleteServicePerimeter handles the gRPC request by passing it to the underlying ServicePerimeter Delete() method.
func (s *ServicePerimeterServer) DeleteAccesscontextmanagerServicePerimeter(ctx context.Context, request *accesscontextmanagerpb.DeleteAccesscontextmanagerServicePerimeterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServicePerimeter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServicePerimeter(ctx, ProtoToServicePerimeter(request.GetResource()))

}

// ListAccesscontextmanagerServicePerimeter handles the gRPC request by passing it to the underlying ServicePerimeterList() method.
func (s *ServicePerimeterServer) ListAccesscontextmanagerServicePerimeter(ctx context.Context, request *accesscontextmanagerpb.ListAccesscontextmanagerServicePerimeterRequest) (*accesscontextmanagerpb.ListAccesscontextmanagerServicePerimeterResponse, error) {
	cl, err := createConfigServicePerimeter(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServicePerimeter(ctx, request.Policy)
	if err != nil {
		return nil, err
	}
	var protos []*accesscontextmanagerpb.AccesscontextmanagerServicePerimeter
	for _, r := range resources.Items {
		rp := ServicePerimeterToProto(r)
		protos = append(protos, rp)
	}
	return &accesscontextmanagerpb.ListAccesscontextmanagerServicePerimeterResponse{Items: protos}, nil
}

func createConfigServicePerimeter(ctx context.Context, service_account_file string) (*accesscontextmanager.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return accesscontextmanager.NewClient(conf), nil
}
