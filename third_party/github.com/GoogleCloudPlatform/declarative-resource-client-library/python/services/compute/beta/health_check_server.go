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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for HealthCheck.
type HealthCheckServer struct{}

// ProtoToHealthCheckHttp2HealthCheckPortSpecificationEnum converts a HealthCheckHttp2HealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum(e betapb.ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum) *beta.HealthCheckHttp2HealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := beta.HealthCheckHttp2HealthCheckPortSpecificationEnum(n[len("ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttp2HealthCheckProxyHeaderEnum converts a HealthCheckHttp2HealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum(e betapb.ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum) *beta.HealthCheckHttp2HealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := beta.HealthCheckHttp2HealthCheckProxyHeaderEnum(n[len("ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttpHealthCheckPortSpecificationEnum converts a HealthCheckHttpHealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum(e betapb.ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum) *beta.HealthCheckHttpHealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := beta.HealthCheckHttpHealthCheckPortSpecificationEnum(n[len("ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttpHealthCheckProxyHeaderEnum converts a HealthCheckHttpHealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum(e betapb.ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum) *beta.HealthCheckHttpHealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := beta.HealthCheckHttpHealthCheckProxyHeaderEnum(n[len("ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttpsHealthCheckPortSpecificationEnum converts a HealthCheckHttpsHealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum(e betapb.ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum) *beta.HealthCheckHttpsHealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := beta.HealthCheckHttpsHealthCheckPortSpecificationEnum(n[len("ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttpsHealthCheckProxyHeaderEnum converts a HealthCheckHttpsHealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum(e betapb.ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum) *beta.HealthCheckHttpsHealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := beta.HealthCheckHttpsHealthCheckProxyHeaderEnum(n[len("ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckSslHealthCheckPortSpecificationEnum converts a HealthCheckSslHealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum(e betapb.ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum) *beta.HealthCheckSslHealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := beta.HealthCheckSslHealthCheckPortSpecificationEnum(n[len("ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckSslHealthCheckProxyHeaderEnum converts a HealthCheckSslHealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum(e betapb.ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum) *beta.HealthCheckSslHealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := beta.HealthCheckSslHealthCheckProxyHeaderEnum(n[len("ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckTcpHealthCheckPortSpecificationEnum converts a HealthCheckTcpHealthCheckPortSpecificationEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum(e betapb.ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum) *beta.HealthCheckTcpHealthCheckPortSpecificationEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum_name[int32(e)]; ok {
		e := beta.HealthCheckTcpHealthCheckPortSpecificationEnum(n[len("ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckTcpHealthCheckProxyHeaderEnum converts a HealthCheckTcpHealthCheckProxyHeaderEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum(e betapb.ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum) *beta.HealthCheckTcpHealthCheckProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum_name[int32(e)]; ok {
		e := beta.HealthCheckTcpHealthCheckProxyHeaderEnum(n[len("ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckTypeEnum converts a HealthCheckTypeEnum enum from its proto representation.
func ProtoToComputeBetaHealthCheckTypeEnum(e betapb.ComputeBetaHealthCheckTypeEnum) *beta.HealthCheckTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaHealthCheckTypeEnum_name[int32(e)]; ok {
		e := beta.HealthCheckTypeEnum(n[len("ComputeBetaHealthCheckTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToHealthCheckHttp2HealthCheck converts a HealthCheckHttp2HealthCheck resource from its proto representation.
func ProtoToComputeBetaHealthCheckHttp2HealthCheck(p *betapb.ComputeBetaHealthCheckHttp2HealthCheck) *beta.HealthCheckHttp2HealthCheck {
	if p == nil {
		return nil
	}
	obj := &beta.HealthCheckHttp2HealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Host:              dcl.StringOrNil(p.Host),
		RequestPath:       dcl.StringOrNil(p.RequestPath),
		ProxyHeader:       ProtoToComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum(p.GetProxyHeader()),
		Response:          dcl.StringOrNil(p.Response),
	}
	return obj
}

// ProtoToHealthCheckHttpHealthCheck converts a HealthCheckHttpHealthCheck resource from its proto representation.
func ProtoToComputeBetaHealthCheckHttpHealthCheck(p *betapb.ComputeBetaHealthCheckHttpHealthCheck) *beta.HealthCheckHttpHealthCheck {
	if p == nil {
		return nil
	}
	obj := &beta.HealthCheckHttpHealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Host:              dcl.StringOrNil(p.Host),
		RequestPath:       dcl.StringOrNil(p.RequestPath),
		ProxyHeader:       ProtoToComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum(p.GetProxyHeader()),
		Response:          dcl.StringOrNil(p.Response),
	}
	return obj
}

// ProtoToHealthCheckHttpsHealthCheck converts a HealthCheckHttpsHealthCheck resource from its proto representation.
func ProtoToComputeBetaHealthCheckHttpsHealthCheck(p *betapb.ComputeBetaHealthCheckHttpsHealthCheck) *beta.HealthCheckHttpsHealthCheck {
	if p == nil {
		return nil
	}
	obj := &beta.HealthCheckHttpsHealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Host:              dcl.StringOrNil(p.Host),
		RequestPath:       dcl.StringOrNil(p.RequestPath),
		ProxyHeader:       ProtoToComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum(p.GetProxyHeader()),
		Response:          dcl.StringOrNil(p.Response),
	}
	return obj
}

// ProtoToHealthCheckSslHealthCheck converts a HealthCheckSslHealthCheck resource from its proto representation.
func ProtoToComputeBetaHealthCheckSslHealthCheck(p *betapb.ComputeBetaHealthCheckSslHealthCheck) *beta.HealthCheckSslHealthCheck {
	if p == nil {
		return nil
	}
	obj := &beta.HealthCheckSslHealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Request:           dcl.StringOrNil(p.Request),
		Response:          dcl.StringOrNil(p.Response),
		ProxyHeader:       ProtoToComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum(p.GetProxyHeader()),
	}
	return obj
}

// ProtoToHealthCheckTcpHealthCheck converts a HealthCheckTcpHealthCheck resource from its proto representation.
func ProtoToComputeBetaHealthCheckTcpHealthCheck(p *betapb.ComputeBetaHealthCheckTcpHealthCheck) *beta.HealthCheckTcpHealthCheck {
	if p == nil {
		return nil
	}
	obj := &beta.HealthCheckTcpHealthCheck{
		Port:              dcl.Int64OrNil(p.Port),
		PortName:          dcl.StringOrNil(p.PortName),
		PortSpecification: ProtoToComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum(p.GetPortSpecification()),
		Request:           dcl.StringOrNil(p.Request),
		Response:          dcl.StringOrNil(p.Response),
		ProxyHeader:       ProtoToComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum(p.GetProxyHeader()),
	}
	return obj
}

// ProtoToHealthCheck converts a HealthCheck resource from its proto representation.
func ProtoToHealthCheck(p *betapb.ComputeBetaHealthCheck) *beta.HealthCheck {
	obj := &beta.HealthCheck{
		CheckIntervalSec:   dcl.Int64OrNil(p.CheckIntervalSec),
		Description:        dcl.StringOrNil(p.Description),
		HealthyThreshold:   dcl.Int64OrNil(p.HealthyThreshold),
		Http2HealthCheck:   ProtoToComputeBetaHealthCheckHttp2HealthCheck(p.GetHttp2HealthCheck()),
		HttpHealthCheck:    ProtoToComputeBetaHealthCheckHttpHealthCheck(p.GetHttpHealthCheck()),
		HttpsHealthCheck:   ProtoToComputeBetaHealthCheckHttpsHealthCheck(p.GetHttpsHealthCheck()),
		Name:               dcl.StringOrNil(p.Name),
		SslHealthCheck:     ProtoToComputeBetaHealthCheckSslHealthCheck(p.GetSslHealthCheck()),
		TcpHealthCheck:     ProtoToComputeBetaHealthCheckTcpHealthCheck(p.GetTcpHealthCheck()),
		Type:               ProtoToComputeBetaHealthCheckTypeEnum(p.GetType()),
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
func ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnumToProto(e *beta.HealthCheckHttp2HealthCheckPortSpecificationEnum) betapb.ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum_value["HealthCheckHttp2HealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum(v)
	}
	return betapb.ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnum(0)
}

// HealthCheckHttp2HealthCheckProxyHeaderEnumToProto converts a HealthCheckHttp2HealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnumToProto(e *beta.HealthCheckHttp2HealthCheckProxyHeaderEnum) betapb.ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum_value["HealthCheckHttp2HealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum(v)
	}
	return betapb.ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnum(0)
}

// HealthCheckHttpHealthCheckPortSpecificationEnumToProto converts a HealthCheckHttpHealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnumToProto(e *beta.HealthCheckHttpHealthCheckPortSpecificationEnum) betapb.ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum_value["HealthCheckHttpHealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum(v)
	}
	return betapb.ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnum(0)
}

// HealthCheckHttpHealthCheckProxyHeaderEnumToProto converts a HealthCheckHttpHealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnumToProto(e *beta.HealthCheckHttpHealthCheckProxyHeaderEnum) betapb.ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum_value["HealthCheckHttpHealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum(v)
	}
	return betapb.ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnum(0)
}

// HealthCheckHttpsHealthCheckPortSpecificationEnumToProto converts a HealthCheckHttpsHealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnumToProto(e *beta.HealthCheckHttpsHealthCheckPortSpecificationEnum) betapb.ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum_value["HealthCheckHttpsHealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum(v)
	}
	return betapb.ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnum(0)
}

// HealthCheckHttpsHealthCheckProxyHeaderEnumToProto converts a HealthCheckHttpsHealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnumToProto(e *beta.HealthCheckHttpsHealthCheckProxyHeaderEnum) betapb.ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum_value["HealthCheckHttpsHealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum(v)
	}
	return betapb.ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnum(0)
}

// HealthCheckSslHealthCheckPortSpecificationEnumToProto converts a HealthCheckSslHealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnumToProto(e *beta.HealthCheckSslHealthCheckPortSpecificationEnum) betapb.ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum_value["HealthCheckSslHealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum(v)
	}
	return betapb.ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnum(0)
}

// HealthCheckSslHealthCheckProxyHeaderEnumToProto converts a HealthCheckSslHealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnumToProto(e *beta.HealthCheckSslHealthCheckProxyHeaderEnum) betapb.ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum_value["HealthCheckSslHealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum(v)
	}
	return betapb.ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnum(0)
}

// HealthCheckTcpHealthCheckPortSpecificationEnumToProto converts a HealthCheckTcpHealthCheckPortSpecificationEnum enum to its proto representation.
func ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnumToProto(e *beta.HealthCheckTcpHealthCheckPortSpecificationEnum) betapb.ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum_value["HealthCheckTcpHealthCheckPortSpecificationEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum(v)
	}
	return betapb.ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnum(0)
}

// HealthCheckTcpHealthCheckProxyHeaderEnumToProto converts a HealthCheckTcpHealthCheckProxyHeaderEnum enum to its proto representation.
func ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnumToProto(e *beta.HealthCheckTcpHealthCheckProxyHeaderEnum) betapb.ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum_value["HealthCheckTcpHealthCheckProxyHeaderEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum(v)
	}
	return betapb.ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnum(0)
}

// HealthCheckTypeEnumToProto converts a HealthCheckTypeEnum enum to its proto representation.
func ComputeBetaHealthCheckTypeEnumToProto(e *beta.HealthCheckTypeEnum) betapb.ComputeBetaHealthCheckTypeEnum {
	if e == nil {
		return betapb.ComputeBetaHealthCheckTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaHealthCheckTypeEnum_value["HealthCheckTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaHealthCheckTypeEnum(v)
	}
	return betapb.ComputeBetaHealthCheckTypeEnum(0)
}

// HealthCheckHttp2HealthCheckToProto converts a HealthCheckHttp2HealthCheck resource to its proto representation.
func ComputeBetaHealthCheckHttp2HealthCheckToProto(o *beta.HealthCheckHttp2HealthCheck) *betapb.ComputeBetaHealthCheckHttp2HealthCheck {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaHealthCheckHttp2HealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeBetaHealthCheckHttp2HealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Host:              dcl.ValueOrEmptyString(o.Host),
		RequestPath:       dcl.ValueOrEmptyString(o.RequestPath),
		ProxyHeader:       ComputeBetaHealthCheckHttp2HealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
		Response:          dcl.ValueOrEmptyString(o.Response),
	}
	return p
}

// HealthCheckHttpHealthCheckToProto converts a HealthCheckHttpHealthCheck resource to its proto representation.
func ComputeBetaHealthCheckHttpHealthCheckToProto(o *beta.HealthCheckHttpHealthCheck) *betapb.ComputeBetaHealthCheckHttpHealthCheck {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaHealthCheckHttpHealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeBetaHealthCheckHttpHealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Host:              dcl.ValueOrEmptyString(o.Host),
		RequestPath:       dcl.ValueOrEmptyString(o.RequestPath),
		ProxyHeader:       ComputeBetaHealthCheckHttpHealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
		Response:          dcl.ValueOrEmptyString(o.Response),
	}
	return p
}

// HealthCheckHttpsHealthCheckToProto converts a HealthCheckHttpsHealthCheck resource to its proto representation.
func ComputeBetaHealthCheckHttpsHealthCheckToProto(o *beta.HealthCheckHttpsHealthCheck) *betapb.ComputeBetaHealthCheckHttpsHealthCheck {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaHealthCheckHttpsHealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeBetaHealthCheckHttpsHealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Host:              dcl.ValueOrEmptyString(o.Host),
		RequestPath:       dcl.ValueOrEmptyString(o.RequestPath),
		ProxyHeader:       ComputeBetaHealthCheckHttpsHealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
		Response:          dcl.ValueOrEmptyString(o.Response),
	}
	return p
}

// HealthCheckSslHealthCheckToProto converts a HealthCheckSslHealthCheck resource to its proto representation.
func ComputeBetaHealthCheckSslHealthCheckToProto(o *beta.HealthCheckSslHealthCheck) *betapb.ComputeBetaHealthCheckSslHealthCheck {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaHealthCheckSslHealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeBetaHealthCheckSslHealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Request:           dcl.ValueOrEmptyString(o.Request),
		Response:          dcl.ValueOrEmptyString(o.Response),
		ProxyHeader:       ComputeBetaHealthCheckSslHealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
	}
	return p
}

// HealthCheckTcpHealthCheckToProto converts a HealthCheckTcpHealthCheck resource to its proto representation.
func ComputeBetaHealthCheckTcpHealthCheckToProto(o *beta.HealthCheckTcpHealthCheck) *betapb.ComputeBetaHealthCheckTcpHealthCheck {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaHealthCheckTcpHealthCheck{
		Port:              dcl.ValueOrEmptyInt64(o.Port),
		PortName:          dcl.ValueOrEmptyString(o.PortName),
		PortSpecification: ComputeBetaHealthCheckTcpHealthCheckPortSpecificationEnumToProto(o.PortSpecification),
		Request:           dcl.ValueOrEmptyString(o.Request),
		Response:          dcl.ValueOrEmptyString(o.Response),
		ProxyHeader:       ComputeBetaHealthCheckTcpHealthCheckProxyHeaderEnumToProto(o.ProxyHeader),
	}
	return p
}

// HealthCheckToProto converts a HealthCheck resource to its proto representation.
func HealthCheckToProto(resource *beta.HealthCheck) *betapb.ComputeBetaHealthCheck {
	p := &betapb.ComputeBetaHealthCheck{
		CheckIntervalSec:   dcl.ValueOrEmptyInt64(resource.CheckIntervalSec),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		HealthyThreshold:   dcl.ValueOrEmptyInt64(resource.HealthyThreshold),
		Http2HealthCheck:   ComputeBetaHealthCheckHttp2HealthCheckToProto(resource.Http2HealthCheck),
		HttpHealthCheck:    ComputeBetaHealthCheckHttpHealthCheckToProto(resource.HttpHealthCheck),
		HttpsHealthCheck:   ComputeBetaHealthCheckHttpsHealthCheckToProto(resource.HttpsHealthCheck),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		SslHealthCheck:     ComputeBetaHealthCheckSslHealthCheckToProto(resource.SslHealthCheck),
		TcpHealthCheck:     ComputeBetaHealthCheckTcpHealthCheckToProto(resource.TcpHealthCheck),
		Type:               ComputeBetaHealthCheckTypeEnumToProto(resource.Type),
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
func (s *HealthCheckServer) applyHealthCheck(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaHealthCheckRequest) (*betapb.ComputeBetaHealthCheck, error) {
	p := ProtoToHealthCheck(request.GetResource())
	res, err := c.ApplyHealthCheck(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HealthCheckToProto(res)
	return r, nil
}

// ApplyHealthCheck handles the gRPC request by passing it to the underlying HealthCheck Apply() method.
func (s *HealthCheckServer) ApplyComputeBetaHealthCheck(ctx context.Context, request *betapb.ApplyComputeBetaHealthCheckRequest) (*betapb.ComputeBetaHealthCheck, error) {
	cl, err := createConfigHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyHealthCheck(ctx, cl, request)
}

// DeleteHealthCheck handles the gRPC request by passing it to the underlying HealthCheck Delete() method.
func (s *HealthCheckServer) DeleteComputeBetaHealthCheck(ctx context.Context, request *betapb.DeleteComputeBetaHealthCheckRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHealthCheck(ctx, ProtoToHealthCheck(request.GetResource()))

}

// ListComputeBetaHealthCheck handles the gRPC request by passing it to the underlying HealthCheckList() method.
func (s *HealthCheckServer) ListComputeBetaHealthCheck(ctx context.Context, request *betapb.ListComputeBetaHealthCheckRequest) (*betapb.ListComputeBetaHealthCheckResponse, error) {
	cl, err := createConfigHealthCheck(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHealthCheck(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaHealthCheck
	for _, r := range resources.Items {
		rp := HealthCheckToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaHealthCheckResponse{Items: protos}, nil
}

func createConfigHealthCheck(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
