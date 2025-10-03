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

// Server implements the gRPC interface for SslPolicy.
type SslPolicyServer struct{}

// ProtoToSslPolicyProfileEnum converts a SslPolicyProfileEnum enum from its proto representation.
func ProtoToComputeBetaSslPolicyProfileEnum(e betapb.ComputeBetaSslPolicyProfileEnum) *beta.SslPolicyProfileEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaSslPolicyProfileEnum_name[int32(e)]; ok {
		e := beta.SslPolicyProfileEnum(n[len("ComputeBetaSslPolicyProfileEnum"):])
		return &e
	}
	return nil
}

// ProtoToSslPolicyMinTlsVersionEnum converts a SslPolicyMinTlsVersionEnum enum from its proto representation.
func ProtoToComputeBetaSslPolicyMinTlsVersionEnum(e betapb.ComputeBetaSslPolicyMinTlsVersionEnum) *beta.SslPolicyMinTlsVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaSslPolicyMinTlsVersionEnum_name[int32(e)]; ok {
		e := beta.SslPolicyMinTlsVersionEnum(n[len("ComputeBetaSslPolicyMinTlsVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToSslPolicyWarning converts a SslPolicyWarning resource from its proto representation.
func ProtoToComputeBetaSslPolicyWarning(p *betapb.ComputeBetaSslPolicyWarning) *beta.SslPolicyWarning {
	if p == nil {
		return nil
	}
	obj := &beta.SslPolicyWarning{
		Code:    dcl.StringOrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetData() {
		obj.Data = append(obj.Data, *ProtoToComputeBetaSslPolicyWarningData(r))
	}
	return obj
}

// ProtoToSslPolicyWarningData converts a SslPolicyWarningData resource from its proto representation.
func ProtoToComputeBetaSslPolicyWarningData(p *betapb.ComputeBetaSslPolicyWarningData) *beta.SslPolicyWarningData {
	if p == nil {
		return nil
	}
	obj := &beta.SslPolicyWarningData{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToSslPolicy converts a SslPolicy resource from its proto representation.
func ProtoToSslPolicy(p *betapb.ComputeBetaSslPolicy) *beta.SslPolicy {
	obj := &beta.SslPolicy{
		Id:            dcl.Int64OrNil(p.Id),
		SelfLink:      dcl.StringOrNil(p.SelfLink),
		Name:          dcl.StringOrNil(p.Name),
		Description:   dcl.StringOrNil(p.Description),
		Profile:       ProtoToComputeBetaSslPolicyProfileEnum(p.GetProfile()),
		MinTlsVersion: ProtoToComputeBetaSslPolicyMinTlsVersionEnum(p.GetMinTlsVersion()),
		Project:       dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetEnabledFeature() {
		obj.EnabledFeature = append(obj.EnabledFeature, r)
	}
	for _, r := range p.GetCustomFeature() {
		obj.CustomFeature = append(obj.CustomFeature, r)
	}
	for _, r := range p.GetWarning() {
		obj.Warning = append(obj.Warning, *ProtoToComputeBetaSslPolicyWarning(r))
	}
	return obj
}

// SslPolicyProfileEnumToProto converts a SslPolicyProfileEnum enum to its proto representation.
func ComputeBetaSslPolicyProfileEnumToProto(e *beta.SslPolicyProfileEnum) betapb.ComputeBetaSslPolicyProfileEnum {
	if e == nil {
		return betapb.ComputeBetaSslPolicyProfileEnum(0)
	}
	if v, ok := betapb.ComputeBetaSslPolicyProfileEnum_value["SslPolicyProfileEnum"+string(*e)]; ok {
		return betapb.ComputeBetaSslPolicyProfileEnum(v)
	}
	return betapb.ComputeBetaSslPolicyProfileEnum(0)
}

// SslPolicyMinTlsVersionEnumToProto converts a SslPolicyMinTlsVersionEnum enum to its proto representation.
func ComputeBetaSslPolicyMinTlsVersionEnumToProto(e *beta.SslPolicyMinTlsVersionEnum) betapb.ComputeBetaSslPolicyMinTlsVersionEnum {
	if e == nil {
		return betapb.ComputeBetaSslPolicyMinTlsVersionEnum(0)
	}
	if v, ok := betapb.ComputeBetaSslPolicyMinTlsVersionEnum_value["SslPolicyMinTlsVersionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaSslPolicyMinTlsVersionEnum(v)
	}
	return betapb.ComputeBetaSslPolicyMinTlsVersionEnum(0)
}

// SslPolicyWarningToProto converts a SslPolicyWarning resource to its proto representation.
func ComputeBetaSslPolicyWarningToProto(o *beta.SslPolicyWarning) *betapb.ComputeBetaSslPolicyWarning {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaSslPolicyWarning{
		Code:    dcl.ValueOrEmptyString(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Data {
		p.Data = append(p.Data, ComputeBetaSslPolicyWarningDataToProto(&r))
	}
	return p
}

// SslPolicyWarningDataToProto converts a SslPolicyWarningData resource to its proto representation.
func ComputeBetaSslPolicyWarningDataToProto(o *beta.SslPolicyWarningData) *betapb.ComputeBetaSslPolicyWarningData {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaSslPolicyWarningData{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// SslPolicyToProto converts a SslPolicy resource to its proto representation.
func SslPolicyToProto(resource *beta.SslPolicy) *betapb.ComputeBetaSslPolicy {
	p := &betapb.ComputeBetaSslPolicy{
		Id:            dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:      dcl.ValueOrEmptyString(resource.SelfLink),
		Name:          dcl.ValueOrEmptyString(resource.Name),
		Description:   dcl.ValueOrEmptyString(resource.Description),
		Profile:       ComputeBetaSslPolicyProfileEnumToProto(resource.Profile),
		MinTlsVersion: ComputeBetaSslPolicyMinTlsVersionEnumToProto(resource.MinTlsVersion),
		Project:       dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.EnabledFeature {
		p.EnabledFeature = append(p.EnabledFeature, r)
	}
	for _, r := range resource.CustomFeature {
		p.CustomFeature = append(p.CustomFeature, r)
	}
	for _, r := range resource.Warning {
		p.Warning = append(p.Warning, ComputeBetaSslPolicyWarningToProto(&r))
	}

	return p
}

// ApplySslPolicy handles the gRPC request by passing it to the underlying SslPolicy Apply() method.
func (s *SslPolicyServer) applySslPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaSslPolicyRequest) (*betapb.ComputeBetaSslPolicy, error) {
	p := ProtoToSslPolicy(request.GetResource())
	res, err := c.ApplySslPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SslPolicyToProto(res)
	return r, nil
}

// ApplySslPolicy handles the gRPC request by passing it to the underlying SslPolicy Apply() method.
func (s *SslPolicyServer) ApplyComputeBetaSslPolicy(ctx context.Context, request *betapb.ApplyComputeBetaSslPolicyRequest) (*betapb.ComputeBetaSslPolicy, error) {
	cl, err := createConfigSslPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySslPolicy(ctx, cl, request)
}

// DeleteSslPolicy handles the gRPC request by passing it to the underlying SslPolicy Delete() method.
func (s *SslPolicyServer) DeleteComputeBetaSslPolicy(ctx context.Context, request *betapb.DeleteComputeBetaSslPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSslPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSslPolicy(ctx, ProtoToSslPolicy(request.GetResource()))

}

// ListComputeBetaSslPolicy handles the gRPC request by passing it to the underlying SslPolicyList() method.
func (s *SslPolicyServer) ListComputeBetaSslPolicy(ctx context.Context, request *betapb.ListComputeBetaSslPolicyRequest) (*betapb.ListComputeBetaSslPolicyResponse, error) {
	cl, err := createConfigSslPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSslPolicy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaSslPolicy
	for _, r := range resources.Items {
		rp := SslPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaSslPolicyResponse{Items: protos}, nil
}

func createConfigSslPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
