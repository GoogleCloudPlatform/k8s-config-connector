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

// Server implements the gRPC interface for SslPolicy.
type SslPolicyServer struct{}

// ProtoToSslPolicyProfileEnum converts a SslPolicyProfileEnum enum from its proto representation.
func ProtoToComputeSslPolicyProfileEnum(e computepb.ComputeSslPolicyProfileEnum) *compute.SslPolicyProfileEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSslPolicyProfileEnum_name[int32(e)]; ok {
		e := compute.SslPolicyProfileEnum(n[len("ComputeSslPolicyProfileEnum"):])
		return &e
	}
	return nil
}

// ProtoToSslPolicyMinTlsVersionEnum converts a SslPolicyMinTlsVersionEnum enum from its proto representation.
func ProtoToComputeSslPolicyMinTlsVersionEnum(e computepb.ComputeSslPolicyMinTlsVersionEnum) *compute.SslPolicyMinTlsVersionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSslPolicyMinTlsVersionEnum_name[int32(e)]; ok {
		e := compute.SslPolicyMinTlsVersionEnum(n[len("ComputeSslPolicyMinTlsVersionEnum"):])
		return &e
	}
	return nil
}

// ProtoToSslPolicyWarning converts a SslPolicyWarning resource from its proto representation.
func ProtoToComputeSslPolicyWarning(p *computepb.ComputeSslPolicyWarning) *compute.SslPolicyWarning {
	if p == nil {
		return nil
	}
	obj := &compute.SslPolicyWarning{
		Code:    dcl.StringOrNil(p.Code),
		Message: dcl.StringOrNil(p.Message),
	}
	for _, r := range p.GetData() {
		obj.Data = append(obj.Data, *ProtoToComputeSslPolicyWarningData(r))
	}
	return obj
}

// ProtoToSslPolicyWarningData converts a SslPolicyWarningData resource from its proto representation.
func ProtoToComputeSslPolicyWarningData(p *computepb.ComputeSslPolicyWarningData) *compute.SslPolicyWarningData {
	if p == nil {
		return nil
	}
	obj := &compute.SslPolicyWarningData{
		Key:   dcl.StringOrNil(p.Key),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToSslPolicy converts a SslPolicy resource from its proto representation.
func ProtoToSslPolicy(p *computepb.ComputeSslPolicy) *compute.SslPolicy {
	obj := &compute.SslPolicy{
		Id:            dcl.Int64OrNil(p.Id),
		SelfLink:      dcl.StringOrNil(p.SelfLink),
		Name:          dcl.StringOrNil(p.Name),
		Description:   dcl.StringOrNil(p.Description),
		Profile:       ProtoToComputeSslPolicyProfileEnum(p.GetProfile()),
		MinTlsVersion: ProtoToComputeSslPolicyMinTlsVersionEnum(p.GetMinTlsVersion()),
		Project:       dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetEnabledFeature() {
		obj.EnabledFeature = append(obj.EnabledFeature, r)
	}
	for _, r := range p.GetCustomFeature() {
		obj.CustomFeature = append(obj.CustomFeature, r)
	}
	for _, r := range p.GetWarning() {
		obj.Warning = append(obj.Warning, *ProtoToComputeSslPolicyWarning(r))
	}
	return obj
}

// SslPolicyProfileEnumToProto converts a SslPolicyProfileEnum enum to its proto representation.
func ComputeSslPolicyProfileEnumToProto(e *compute.SslPolicyProfileEnum) computepb.ComputeSslPolicyProfileEnum {
	if e == nil {
		return computepb.ComputeSslPolicyProfileEnum(0)
	}
	if v, ok := computepb.ComputeSslPolicyProfileEnum_value["SslPolicyProfileEnum"+string(*e)]; ok {
		return computepb.ComputeSslPolicyProfileEnum(v)
	}
	return computepb.ComputeSslPolicyProfileEnum(0)
}

// SslPolicyMinTlsVersionEnumToProto converts a SslPolicyMinTlsVersionEnum enum to its proto representation.
func ComputeSslPolicyMinTlsVersionEnumToProto(e *compute.SslPolicyMinTlsVersionEnum) computepb.ComputeSslPolicyMinTlsVersionEnum {
	if e == nil {
		return computepb.ComputeSslPolicyMinTlsVersionEnum(0)
	}
	if v, ok := computepb.ComputeSslPolicyMinTlsVersionEnum_value["SslPolicyMinTlsVersionEnum"+string(*e)]; ok {
		return computepb.ComputeSslPolicyMinTlsVersionEnum(v)
	}
	return computepb.ComputeSslPolicyMinTlsVersionEnum(0)
}

// SslPolicyWarningToProto converts a SslPolicyWarning resource to its proto representation.
func ComputeSslPolicyWarningToProto(o *compute.SslPolicyWarning) *computepb.ComputeSslPolicyWarning {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSslPolicyWarning{
		Code:    dcl.ValueOrEmptyString(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	for _, r := range o.Data {
		p.Data = append(p.Data, ComputeSslPolicyWarningDataToProto(&r))
	}
	return p
}

// SslPolicyWarningDataToProto converts a SslPolicyWarningData resource to its proto representation.
func ComputeSslPolicyWarningDataToProto(o *compute.SslPolicyWarningData) *computepb.ComputeSslPolicyWarningData {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSslPolicyWarningData{
		Key:   dcl.ValueOrEmptyString(o.Key),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// SslPolicyToProto converts a SslPolicy resource to its proto representation.
func SslPolicyToProto(resource *compute.SslPolicy) *computepb.ComputeSslPolicy {
	p := &computepb.ComputeSslPolicy{
		Id:            dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:      dcl.ValueOrEmptyString(resource.SelfLink),
		Name:          dcl.ValueOrEmptyString(resource.Name),
		Description:   dcl.ValueOrEmptyString(resource.Description),
		Profile:       ComputeSslPolicyProfileEnumToProto(resource.Profile),
		MinTlsVersion: ComputeSslPolicyMinTlsVersionEnumToProto(resource.MinTlsVersion),
		Project:       dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.EnabledFeature {
		p.EnabledFeature = append(p.EnabledFeature, r)
	}
	for _, r := range resource.CustomFeature {
		p.CustomFeature = append(p.CustomFeature, r)
	}
	for _, r := range resource.Warning {
		p.Warning = append(p.Warning, ComputeSslPolicyWarningToProto(&r))
	}

	return p
}

// ApplySslPolicy handles the gRPC request by passing it to the underlying SslPolicy Apply() method.
func (s *SslPolicyServer) applySslPolicy(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeSslPolicyRequest) (*computepb.ComputeSslPolicy, error) {
	p := ProtoToSslPolicy(request.GetResource())
	res, err := c.ApplySslPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SslPolicyToProto(res)
	return r, nil
}

// ApplySslPolicy handles the gRPC request by passing it to the underlying SslPolicy Apply() method.
func (s *SslPolicyServer) ApplyComputeSslPolicy(ctx context.Context, request *computepb.ApplyComputeSslPolicyRequest) (*computepb.ComputeSslPolicy, error) {
	cl, err := createConfigSslPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySslPolicy(ctx, cl, request)
}

// DeleteSslPolicy handles the gRPC request by passing it to the underlying SslPolicy Delete() method.
func (s *SslPolicyServer) DeleteComputeSslPolicy(ctx context.Context, request *computepb.DeleteComputeSslPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSslPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSslPolicy(ctx, ProtoToSslPolicy(request.GetResource()))

}

// ListComputeSslPolicy handles the gRPC request by passing it to the underlying SslPolicyList() method.
func (s *SslPolicyServer) ListComputeSslPolicy(ctx context.Context, request *computepb.ListComputeSslPolicyRequest) (*computepb.ListComputeSslPolicyResponse, error) {
	cl, err := createConfigSslPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSslPolicy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeSslPolicy
	for _, r := range resources.Items {
		rp := SslPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeSslPolicyResponse{Items: protos}, nil
}

func createConfigSslPolicy(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
