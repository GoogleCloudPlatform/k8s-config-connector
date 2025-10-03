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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for HealthCheck.
type HealthCheckServer struct{}

// ProtoToHealthCheckHttp2HealthCheckPortSpecificationEnum converts a HealthCheckHttp2HealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeHealthCheckHttp2HealthCheckPortSpecificationEnum(e computepb.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum) *compute.HealthCheckHttp2HealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := compute.HealthCheckHttp2HealthCheckPortSpecificationEnum(n[len("ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttp2HealthCheckProxyHeaderEnum converts a HealthCheckHttp2HealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeHealthCheckHttp2HealthCheckProxyHeaderEnum(e computepb.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum) *compute.HealthCheckHttp2HealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := compute.HealthCheckHttp2HealthCheckProxyHeaderEnum(n[len("ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttpHealthCheckPortSpecificationEnum converts a HealthCheckHttpHealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeHealthCheckHttpHealthCheckPortSpecificationEnum(e computepb.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum) *compute.HealthCheckHttpHealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := compute.HealthCheckHttpHealthCheckPortSpecificationEnum(n[len("ComputeHealthCheckHttpHealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttpHealthCheckProxyHeaderEnum converts a HealthCheckHttpHealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeHealthCheckHttpHealthCheckProxyHeaderEnum(e computepb.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum) *compute.HealthCheckHttpHealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := compute.HealthCheckHttpHealthCheckProxyHeaderEnum(n[len("ComputeHealthCheckHttpHealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttpsHealthCheckPortSpecificationEnum converts a HealthCheckHttpsHealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeHealthCheckHttpsHealthCheckPortSpecificationEnum(e computepb.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum) *compute.HealthCheckHttpsHealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := compute.HealthCheckHttpsHealthCheckPortSpecificationEnum(n[len("ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttpsHealthCheckProxyHeaderEnum converts a HealthCheckHttpsHealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeHealthCheckHttpsHealthCheckProxyHeaderEnum(e computepb.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum) *compute.HealthCheckHttpsHealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := compute.HealthCheckHttpsHealthCheckProxyHeaderEnum(n[len("ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckSslHealthCheckPortSpecificationEnum converts a HealthCheckSslHealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeHealthCheckSslHealthCheckPortSpecificationEnum(e computepb.ComputeHealthCheckSslHealthCheckPortSpecificationEnum) *compute.HealthCheckSslHealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckSslHealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := compute.HealthCheckSslHealthCheckPortSpecificationEnum(n[len("ComputeHealthCheckSslHealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckSslHealthCheckProxyHeaderEnum converts a HealthCheckSslHealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeHealthCheckSslHealthCheckProxyHeaderEnum(e computepb.ComputeHealthCheckSslHealthCheckProxyHeaderEnum) *compute.HealthCheckSslHealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckSslHealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := compute.HealthCheckSslHealthCheckProxyHeaderEnum(n[len("ComputeHealthCheckSslHealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckTcpHealthCheckPortSpecificationEnum converts a HealthCheckTcpHealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeHealthCheckTcpHealthCheckPortSpecificationEnum(e computepb.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum) *compute.HealthCheckTcpHealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := compute.HealthCheckTcpHealthCheckPortSpecificationEnum(n[len("ComputeHealthCheckTcpHealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckTcpHealthCheckProxyHeaderEnum converts a HealthCheckTcpHealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeHealthCheckTcpHealthCheckProxyHeaderEnum(e computepb.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum) *compute.HealthCheckTcpHealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := compute.HealthCheckTcpHealthCheckProxyHeaderEnum(n[len("ComputeHealthCheckTcpHealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckTypeEnum converts a HealthCheckTypeEnum enum from its proto representation.
func ProtoToComputeHealthCheckTypeEnum(e computepb.ComputeHealthCheckTypeEnum) *compute.HealthCheckTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeHealthCheckTypeEnum_name[int32(e)]; ok {
		e := compute.HealthCheckTypeEnum(n[len("ComputeHealthCheckTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttp2HealthCheck converts a HealthCheckHttp2HealthCheck resource from its proto representation.
func ProtoToComputeHealthCheckHttp2HealthCheck(p *computepb.ComputeHealthCheckHttp2HealthCheck) *compute.HealthCheckHttp2HealthCheck {
	if p == nil {
		return nil
	}
	obj := &compute.HealthCheckHttp2HealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeHealthCheckHttp2HealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Host:              dcl.StringOrNil(p.Host),
		RequestPath:       dcl.StringOrNil(p.RequestPath),
		ProxyHeader:       ProtoToComputeHealthCheckHttp2HealthCheckProxyHeaderEnum(p.GetProxyHeader()),
		Response:          dcl.StringOrNil(p.Response),
	}
	return obj
}

// ProtoToHealthCheckHttpHealthCheck converts a HealthCheckHttpHealthCheck resource from its proto representation.
func ProtoToComputeHealthCheckHttpHealthCheck(p *computepb.ComputeHealthCheckHttpHealthCheck) *compute.HealthCheckHttpHealthCheck {
	if p == nil {
		return nil
	}
	obj := &compute.HealthCheckHttpHealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeHealthCheckHttpHealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Host:              dcl.StringOrNil(p.Host),
		RequestPath:       dcl.StringOrNil(p.RequestPath),
		ProxyHeader:       ProtoToComputeHealthCheckHttpHealthCheckProxyHeaderEnum(p.GetProxyHeader()),
		Response:          dcl.StringOrNil(p.Response),
	}
	return obj
}

// ProtoToHealthCheckHttpsHealthCheck converts a HealthCheckHttpsHealthCheck resource from its proto representation.
func ProtoToComputeHealthCheckHttpsHealthCheck(p *computepb.ComputeHealthCheckHttpsHealthCheck) *compute.HealthCheckHttpsHealthCheck {
	if p == nil {
		return nil
	}
	obj := &compute.HealthCheckHttpsHealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeHealthCheckHttpsHealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Host:              dcl.StringOrNil(p.Host),
		RequestPath:       dcl.StringOrNil(p.RequestPath),
		ProxyHeader:       ProtoToComputeHealthCheckHttpsHealthCheckProxyHeaderEnum(p.GetProxyHeader()),
		Response:          dcl.StringOrNil(p.Response),
	}
	return obj
}

// ProtoToHealthCheckSslHealthCheck converts a HealthCheckSslHealthCheck resource from its proto representation.
func ProtoToComputeHealthCheckSslHealthCheck(p *computepb.ComputeHealthCheckSslHealthCheck) *compute.HealthCheckSslHealthCheck {
	if p == nil {
		return nil
	}
	obj := &compute.HealthCheckSslHealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeHealthCheckSslHealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Request:           dcl.StringOrNil(p.Request),
		Response:          dcl.StringOrNil(p.Response),
		ProxyHeader:       ProtoToComputeHealthCheckSslHealthCheckProxyHeaderEnum(p.GetProxyHeader()),
	}
	return obj
}

// ProtoToHealthCheckTcpHealthCheck converts a HealthCheckTcpHealthCheck resource from its proto representation.
func ProtoToComputeHealthCheckTcpHealthCheck(p *computepb.ComputeHealthCheckTcpHealthCheck) *compute.HealthCheckTcpHealthCheck {
	if p == nil {
		return nil
	}
	obj := &compute.HealthCheckTcpHealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeHealthCheckTcpHealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Request:           dcl.StringOrNil(p.Request),
		Response:          dcl.StringOrNil(p.Response),
		ProxyHeader:       ProtoToComputeHealthCheckTcpHealthCheckProxyHeaderEnum(p.GetProxyHeader()),
	}
	return obj
}

// ProtoToHealthCheck converts a HealthCheck resource from its proto representation.
func ProtoToHealthCheck(p *computepb.ComputeHealthCheck) *compute.HealthCheck {
	obj := &compute.HealthCheck{
		CheckIntervalSec:   dcl.Int64OrNil(p.CheckIntervalSec),
		Description:        dcl.StringOrNil(p.Description),
		HealthyThreshold:   dcl.Int64OrNil(p.HealthyThreshold),
		Http2HealthCheck:   ProtoToComputeHealthCheckHttp2HealthCheck(p.GetHttp2HealthCheck()),
		HttpHealthCheck:    ProtoToComputeHealthCheckHttpHealthCheck(p.GetHttpHealthCheck()),
		HttpsHealthCheck:   ProtoToComputeHealthCheckHttpsHealthCheck(p.GetHttpsHealthCheck()),
		Name:               dcl.StringOrNil(p.Name),
		SslHealthCheck:     ProtoToComputeHealthCheckSslHealthCheck(p.GetSslHealthCheck()),
		TcpHealthCheck:     ProtoToComputeHealthCheckTcpHealthCheck(p.GetTcpHealthCheck()),
		Type:               ProtoToComputeHealthCheckTypeEnum(p.GetType()),
		UnhealthyThreshold: dcl.Int64OrNil(p.UnhealthyThreshold),
		TimeoutSec:         dcl.Int64OrNil(p.TimeoutSec),
		Region:             dcl.StringOrNil(p.Region),
		Project:            dcl.StringOrNil(p.Project),
		SelfLink:           dcl.StringOrNil(p.SelfLink),
		Location:           dcl.StringOrNil(p.Location),
	}
	return obj
}

// HealthCheckHttp2HealthCheckPortSpecificationEnumToProto converts a HealthCheckHttp2HealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeHealthCheckHttp2HealthCheckPortSpecificationEnumToProto(e *compute.HealthCheckHttp2HealthCheckPortSpecificationEnum) computepb.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum {
	if e == nil {
		return computepb.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum_value["HealthCheckHttp2HealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum(v)
	}
	return computepb.ComputeHealthCheckHttp2HealthCheckPortSpecificationEnum(0)
}

// HealthCheckHttp2HealthCheckProxyHeaderEnumToProto converts a HealthCheckHttp2HealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeHealthCheckHttp2HealthCheckProxyHeaderEnumToProto(e *compute.HealthCheckHttp2HealthCheckProxyHeaderEnum) computepb.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum {
	if e == nil {
		return computepb.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum_value["HealthCheckHttp2HealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum(v)
	}
	return computepb.ComputeHealthCheckHttp2HealthCheckProxyHeaderEnum(0)
}

// HealthCheckHttpHealthCheckPortSpecificationEnumToProto converts a HealthCheckHttpHealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeHealthCheckHttpHealthCheckPortSpecificationEnumToProto(e *compute.HealthCheckHttpHealthCheckPortSpecificationEnum) computepb.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum {
	if e == nil {
		return computepb.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum_value["HealthCheckHttpHealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum(v)
	}
	return computepb.ComputeHealthCheckHttpHealthCheckPortSpecificationEnum(0)
}

// HealthCheckHttpHealthCheckProxyHeaderEnumToProto converts a HealthCheckHttpHealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeHealthCheckHttpHealthCheckProxyHeaderEnumToProto(e *compute.HealthCheckHttpHealthCheckProxyHeaderEnum) computepb.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum {
	if e == nil {
		return computepb.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum_value["HealthCheckHttpHealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum(v)
	}
	return computepb.ComputeHealthCheckHttpHealthCheckProxyHeaderEnum(0)
}

// HealthCheckHttpsHealthCheckPortSpecificationEnumToProto converts a HealthCheckHttpsHealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeHealthCheckHttpsHealthCheckPortSpecificationEnumToProto(e *compute.HealthCheckHttpsHealthCheckPortSpecificationEnum) computepb.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum {
	if e == nil {
		return computepb.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum_value["HealthCheckHttpsHealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum(v)
	}
	return computepb.ComputeHealthCheckHttpsHealthCheckPortSpecificationEnum(0)
}

// HealthCheckHttpsHealthCheckProxyHeaderEnumToProto converts a HealthCheckHttpsHealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeHealthCheckHttpsHealthCheckProxyHeaderEnumToProto(e *compute.HealthCheckHttpsHealthCheckProxyHeaderEnum) computepb.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum {
	if e == nil {
		return computepb.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum_value["HealthCheckHttpsHealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum(v)
	}
	return computepb.ComputeHealthCheckHttpsHealthCheckProxyHeaderEnum(0)
}

// HealthCheckSslHealthCheckPortSpecificationEnumToProto converts a HealthCheckSslHealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeHealthCheckSslHealthCheckPortSpecificationEnumToProto(e *compute.HealthCheckSslHealthCheckPortSpecificationEnum) computepb.ComputeHealthCheckSslHealthCheckPortSpecificationEnum {
	if e == nil {
		return computepb.ComputeHealthCheckSslHealthCheckPortSpecificationEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckSslHealthCheckPortSpecificationEnum_value["HealthCheckSslHealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckSslHealthCheckPortSpecificationEnum(v)
	}
	return computepb.ComputeHealthCheckSslHealthCheckPortSpecificationEnum(0)
}

// HealthCheckSslHealthCheckProxyHeaderEnumToProto converts a HealthCheckSslHealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeHealthCheckSslHealthCheckProxyHeaderEnumToProto(e *compute.HealthCheckSslHealthCheckProxyHeaderEnum) computepb.ComputeHealthCheckSslHealthCheckProxyHeaderEnum {
	if e == nil {
		return computepb.ComputeHealthCheckSslHealthCheckProxyHeaderEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckSslHealthCheckProxyHeaderEnum_value["HealthCheckSslHealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckSslHealthCheckProxyHeaderEnum(v)
	}
	return computepb.ComputeHealthCheckSslHealthCheckProxyHeaderEnum(0)
}

// HealthCheckTcpHealthCheckPortSpecificationEnumToProto converts a HealthCheckTcpHealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeHealthCheckTcpHealthCheckPortSpecificationEnumToProto(e *compute.HealthCheckTcpHealthCheckPortSpecificationEnum) computepb.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum {
	if e == nil {
		return computepb.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum_value["HealthCheckTcpHealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum(v)
	}
	return computepb.ComputeHealthCheckTcpHealthCheckPortSpecificationEnum(0)
}

// HealthCheckTcpHealthCheckProxyHeaderEnumToProto converts a HealthCheckTcpHealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeHealthCheckTcpHealthCheckProxyHeaderEnumToProto(e *compute.HealthCheckTcpHealthCheckProxyHeaderEnum) computepb.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum {
	if e == nil {
		return computepb.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum_value["HealthCheckTcpHealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum(v)
	}
	return computepb.ComputeHealthCheckTcpHealthCheckProxyHeaderEnum(0)
}

// HealthCheckTypeEnumToProto converts a HealthCheckTypeEnum enum to its proto representation.
func ComputeHealthCheckTypeEnumToProto(e *compute.HealthCheckTypeEnum) computepb.ComputeHealthCheckTypeEnum {
	if e == nil {
		return computepb.ComputeHealthCheckTypeEnum(0)
	}
	if v, ok := computepb.ComputeHealthCheckTypeEnum_value["HealthCheckTypeEnum"+string(*e)]; ok {
		return computepb.ComputeHealthCheckTypeEnum(v)
	}
	return computepb.ComputeHealthCheckTypeEnum(0)
}

// HealthCheckHttp2HealthCheckToProto converts a HealthCheckHttp2HealthCheck resource to its proto representation.
func ComputeHealthCheckHttp2HealthCheckToProto(o *compute.HealthCheckHttp2HealthCheck) *computepb.ComputeHealthCheckHttp2HealthCheck {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeHealthCheckHttp2HealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeHealthCheckHttp2HealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Host:              dcl.ValueOrEmptyString(o.Host),
		RequestPath:       dcl.ValueOrEmptyString(o.RequestPath),
		ProxyHeader:       ComputeHealthCheckHttp2HealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
		Response:          dcl.ValueOrEmptyString(o.Response),
	}
	return p
}

// HealthCheckHttpHealthCheckToProto converts a HealthCheckHttpHealthCheck resource to its proto representation.
func ComputeHealthCheckHttpHealthCheckToProto(o *compute.HealthCheckHttpHealthCheck) *computepb.ComputeHealthCheckHttpHealthCheck {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeHealthCheckHttpHealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeHealthCheckHttpHealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Host:              dcl.ValueOrEmptyString(o.Host),
		RequestPath:       dcl.ValueOrEmptyString(o.RequestPath),
		ProxyHeader:       ComputeHealthCheckHttpHealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
		Response:          dcl.ValueOrEmptyString(o.Response),
	}
	return p
}

// HealthCheckHttpsHealthCheckToProto converts a HealthCheckHttpsHealthCheck resource to its proto representation.
func ComputeHealthCheckHttpsHealthCheckToProto(o *compute.HealthCheckHttpsHealthCheck) *computepb.ComputeHealthCheckHttpsHealthCheck {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeHealthCheckHttpsHealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeHealthCheckHttpsHealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Host:              dcl.ValueOrEmptyString(o.Host),
		RequestPath:       dcl.ValueOrEmptyString(o.RequestPath),
		ProxyHeader:       ComputeHealthCheckHttpsHealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
		Response:          dcl.ValueOrEmptyString(o.Response),
	}
	return p
}

// HealthCheckSslHealthCheckToProto converts a HealthCheckSslHealthCheck resource to its proto representation.
func ComputeHealthCheckSslHealthCheckToProto(o *compute.HealthCheckSslHealthCheck) *computepb.ComputeHealthCheckSslHealthCheck {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeHealthCheckSslHealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeHealthCheckSslHealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Request:           dcl.ValueOrEmptyString(o.Request),
		Response:          dcl.ValueOrEmptyString(o.Response),
		ProxyHeader:       ComputeHealthCheckSslHealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
	}
	return p
}

// HealthCheckTcpHealthCheckToProto converts a HealthCheckTcpHealthCheck resource to its proto representation.
func ComputeHealthCheckTcpHealthCheckToProto(o *compute.HealthCheckTcpHealthCheck) *computepb.ComputeHealthCheckTcpHealthCheck {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeHealthCheckTcpHealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeHealthCheckTcpHealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Request:           dcl.ValueOrEmptyString(o.Request),
		Response:          dcl.ValueOrEmptyString(o.Response),
		ProxyHeader:       ComputeHealthCheckTcpHealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
	}
	return p
}

// HealthCheckToProto converts a HealthCheck resource to its proto representation.
func HealthCheckToProto(resource *compute.HealthCheck) *computepb.ComputeHealthCheck {
	p := &computepb.ComputeHealthCheck{
		CheckIntervalSec:   dcl.ValueOrEmptyInt64(resource.CheckIntervalSec),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		HealthyThreshold:   dcl.ValueOrEmptyInt64(resource.HealthyThreshold),
		Http2HealthCheck:   ComputeHealthCheckHttp2HealthCheckToProto(resource.Http2HealthCheck),
		HttpHealthCheck:    ComputeHealthCheckHttpHealthCheckToProto(resource.HttpHealthCheck),
		HttpsHealthCheck:   ComputeHealthCheckHttpsHealthCheckToProto(resource.HttpsHealthCheck),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		SslHealthCheck:     ComputeHealthCheckSslHealthCheckToProto(resource.SslHealthCheck),
		TcpHealthCheck:     ComputeHealthCheckTcpHealthCheckToProto(resource.TcpHealthCheck),
		Type:               ComputeHealthCheckTypeEnumToProto(resource.Type),
		UnhealthyThreshold: dcl.ValueOrEmptyInt64(resource.UnhealthyThreshold),
		TimeoutSec:         dcl.ValueOrEmptyInt64(resource.TimeoutSec),
		Region:             dcl.ValueOrEmptyString(resource.Region),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		SelfLink:           dcl.ValueOrEmptyString(resource.SelfLink),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyHealthCheck handles the gRPC request by passing it to the underlying HealthCheck Apply() method.
func (s *HealthCheckServer) applyHealthCheck(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeHealthCheckRequest) (*computepb.ComputeHealthCheck, error) {
	p := ProtoToHealthCheck(request.GetResource())
	res, err := c.ApplyHealthCheck(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HealthCheckToProto(res)
	return r, nil
}

// ApplyHealthCheck handles the gRPC request by passing it to the underlying HealthCheck Apply() method.
func (s *HealthCheckServer) ApplyComputeHealthCheck(ctx context.Context, request *computepb.ApplyComputeHealthCheckRequest) (*computepb.ComputeHealthCheck, error) {
	cl, err := createConfigHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyHealthCheck(ctx, cl, request)
}

// DeleteHealthCheck handles the gRPC request by passing it to the underlying HealthCheck Delete() method.
func (s *HealthCheckServer) DeleteComputeHealthCheck(ctx context.Context, request *computepb.DeleteComputeHealthCheckRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHealthCheck(ctx, ProtoToHealthCheck(request.GetResource()))

}

// ListComputeHealthCheck handles the gRPC request by passing it to the underlying HealthCheckList() method.
func (s *HealthCheckServer) ListComputeHealthCheck(ctx context.Context, request *computepb.ListComputeHealthCheckRequest) (*computepb.ListComputeHealthCheckResponse, error) {
	cl, err := createConfigHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHealthCheck(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeHealthCheck
	for _, r := range resources.Items {
		rp := HealthCheckToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeHealthCheckResponse{Items: protos}, nil
}

func createConfigHealthCheck(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
