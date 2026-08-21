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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// SubnetworkServer implements the gRPC interface for Subnetwork.
type SubnetworkServer struct{}

// ProtoToSubnetworkPurposeEnum converts a SubnetworkPurposeEnum enum from its proto representation.
func ProtoToComputeBetaSubnetworkPurposeEnum(e betapb.ComputeBetaSubnetworkPurposeEnum) *beta.SubnetworkPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaSubnetworkPurposeEnum_name[int32(e)]; ok {
		e := beta.SubnetworkPurposeEnum(n[len("ComputeBetaSubnetworkPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkRoleEnum converts a SubnetworkRoleEnum enum from its proto representation.
func ProtoToComputeBetaSubnetworkRoleEnum(e betapb.ComputeBetaSubnetworkRoleEnum) *beta.SubnetworkRoleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaSubnetworkRoleEnum_name[int32(e)]; ok {
		e := beta.SubnetworkRoleEnum(n[len("ComputeBetaSubnetworkRoleEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkLogConfigAggregationIntervalEnum converts a SubnetworkLogConfigAggregationIntervalEnum enum from its proto representation.
func ProtoToComputeBetaSubnetworkLogConfigAggregationIntervalEnum(e betapb.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum) *beta.SubnetworkLogConfigAggregationIntervalEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum_name[int32(e)]; ok {
		e := beta.SubnetworkLogConfigAggregationIntervalEnum(n[len("ComputeBetaSubnetworkLogConfigAggregationIntervalEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkLogConfigMetadataEnum converts a SubnetworkLogConfigMetadataEnum enum from its proto representation.
func ProtoToComputeBetaSubnetworkLogConfigMetadataEnum(e betapb.ComputeBetaSubnetworkLogConfigMetadataEnum) *beta.SubnetworkLogConfigMetadataEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaSubnetworkLogConfigMetadataEnum_name[int32(e)]; ok {
		e := beta.SubnetworkLogConfigMetadataEnum(n[len("ComputeBetaSubnetworkLogConfigMetadataEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkSecondaryIPRanges converts a SubnetworkSecondaryIPRanges object from its proto representation.
func ProtoToComputeBetaSubnetworkSecondaryIPRanges(p *betapb.ComputeBetaSubnetworkSecondaryIPRanges) *beta.SubnetworkSecondaryIPRanges {
	if p == nil {
		return nil
	}
	obj := &beta.SubnetworkSecondaryIPRanges{
		RangeName:   dcl.StringOrNil(p.GetRangeName()),
		IPCidrRange: dcl.StringOrNil(p.GetIpCidrRange()),
	}
	return obj
}

// ProtoToSubnetworkLogConfig converts a SubnetworkLogConfig object from its proto representation.
func ProtoToComputeBetaSubnetworkLogConfig(p *betapb.ComputeBetaSubnetworkLogConfig) *beta.SubnetworkLogConfig {
	if p == nil {
		return nil
	}
	obj := &beta.SubnetworkLogConfig{
		AggregationInterval: ProtoToComputeBetaSubnetworkLogConfigAggregationIntervalEnum(p.GetAggregationInterval()),
		FlowSampling:        dcl.Float64OrNil(p.GetFlowSampling()),
		Metadata:            ProtoToComputeBetaSubnetworkLogConfigMetadataEnum(p.GetMetadata()),
	}
	return obj
}

// ProtoToSubnetwork converts a Subnetwork resource from its proto representation.
func ProtoToSubnetwork(p *betapb.ComputeBetaSubnetwork) *beta.Subnetwork {
	obj := &beta.Subnetwork{
		CreationTimestamp:     dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		GatewayAddress:        dcl.StringOrNil(p.GetGatewayAddress()),
		IPCidrRange:           dcl.StringOrNil(p.GetIpCidrRange()),
		Name:                  dcl.StringOrNil(p.GetName()),
		Network:               dcl.StringOrNil(p.GetNetwork()),
		Fingerprint:           dcl.StringOrNil(p.GetFingerprint()),
		Purpose:               ProtoToComputeBetaSubnetworkPurposeEnum(p.GetPurpose()),
		Role:                  ProtoToComputeBetaSubnetworkRoleEnum(p.GetRole()),
		PrivateIPGoogleAccess: dcl.Bool(p.GetPrivateIpGoogleAccess()),
		Region:                dcl.StringOrNil(p.GetRegion()),
		LogConfig:             ProtoToComputeBetaSubnetworkLogConfig(p.GetLogConfig()),
		Project:               dcl.StringOrNil(p.GetProject()),
		SelfLink:              dcl.StringOrNil(p.GetSelfLink()),
		EnableFlowLogs:        dcl.Bool(p.GetEnableFlowLogs()),
	}
	for _, r := range p.GetSecondaryIpRanges() {
		obj.SecondaryIPRanges = append(obj.SecondaryIPRanges, *ProtoToComputeBetaSubnetworkSecondaryIPRanges(r))
	}
	return obj
}

// SubnetworkPurposeEnumToProto converts a SubnetworkPurposeEnum enum to its proto representation.
func ComputeBetaSubnetworkPurposeEnumToProto(e *beta.SubnetworkPurposeEnum) betapb.ComputeBetaSubnetworkPurposeEnum {
	if e == nil {
		return betapb.ComputeBetaSubnetworkPurposeEnum(0)
	}
	if v, ok := betapb.ComputeBetaSubnetworkPurposeEnum_value["SubnetworkPurposeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaSubnetworkPurposeEnum(v)
	}
	return betapb.ComputeBetaSubnetworkPurposeEnum(0)
}

// SubnetworkRoleEnumToProto converts a SubnetworkRoleEnum enum to its proto representation.
func ComputeBetaSubnetworkRoleEnumToProto(e *beta.SubnetworkRoleEnum) betapb.ComputeBetaSubnetworkRoleEnum {
	if e == nil {
		return betapb.ComputeBetaSubnetworkRoleEnum(0)
	}
	if v, ok := betapb.ComputeBetaSubnetworkRoleEnum_value["SubnetworkRoleEnum"+string(*e)]; ok {
		return betapb.ComputeBetaSubnetworkRoleEnum(v)
	}
	return betapb.ComputeBetaSubnetworkRoleEnum(0)
}

// SubnetworkLogConfigAggregationIntervalEnumToProto converts a SubnetworkLogConfigAggregationIntervalEnum enum to its proto representation.
func ComputeBetaSubnetworkLogConfigAggregationIntervalEnumToProto(e *beta.SubnetworkLogConfigAggregationIntervalEnum) betapb.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum {
	if e == nil {
		return betapb.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum(0)
	}
	if v, ok := betapb.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum_value["SubnetworkLogConfigAggregationIntervalEnum"+string(*e)]; ok {
		return betapb.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum(v)
	}
	return betapb.ComputeBetaSubnetworkLogConfigAggregationIntervalEnum(0)
}

// SubnetworkLogConfigMetadataEnumToProto converts a SubnetworkLogConfigMetadataEnum enum to its proto representation.
func ComputeBetaSubnetworkLogConfigMetadataEnumToProto(e *beta.SubnetworkLogConfigMetadataEnum) betapb.ComputeBetaSubnetworkLogConfigMetadataEnum {
	if e == nil {
		return betapb.ComputeBetaSubnetworkLogConfigMetadataEnum(0)
	}
	if v, ok := betapb.ComputeBetaSubnetworkLogConfigMetadataEnum_value["SubnetworkLogConfigMetadataEnum"+string(*e)]; ok {
		return betapb.ComputeBetaSubnetworkLogConfigMetadataEnum(v)
	}
	return betapb.ComputeBetaSubnetworkLogConfigMetadataEnum(0)
}

// SubnetworkSecondaryIPRangesToProto converts a SubnetworkSecondaryIPRanges object to its proto representation.
func ComputeBetaSubnetworkSecondaryIPRangesToProto(o *beta.SubnetworkSecondaryIPRanges) *betapb.ComputeBetaSubnetworkSecondaryIPRanges {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaSubnetworkSecondaryIPRanges{}
	p.SetRangeName(dcl.ValueOrEmptyString(o.RangeName))
	p.SetIpCidrRange(dcl.ValueOrEmptyString(o.IPCidrRange))
	return p
}

// SubnetworkLogConfigToProto converts a SubnetworkLogConfig object to its proto representation.
func ComputeBetaSubnetworkLogConfigToProto(o *beta.SubnetworkLogConfig) *betapb.ComputeBetaSubnetworkLogConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaSubnetworkLogConfig{}
	p.SetAggregationInterval(ComputeBetaSubnetworkLogConfigAggregationIntervalEnumToProto(o.AggregationInterval))
	p.SetFlowSampling(dcl.ValueOrEmptyDouble(o.FlowSampling))
	p.SetMetadata(ComputeBetaSubnetworkLogConfigMetadataEnumToProto(o.Metadata))
	return p
}

// SubnetworkToProto converts a Subnetwork resource to its proto representation.
func SubnetworkToProto(resource *beta.Subnetwork) *betapb.ComputeBetaSubnetwork {
	p := &betapb.ComputeBetaSubnetwork{}
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGatewayAddress(dcl.ValueOrEmptyString(resource.GatewayAddress))
	p.SetIpCidrRange(dcl.ValueOrEmptyString(resource.IPCidrRange))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetPurpose(ComputeBetaSubnetworkPurposeEnumToProto(resource.Purpose))
	p.SetRole(ComputeBetaSubnetworkRoleEnumToProto(resource.Role))
	p.SetPrivateIpGoogleAccess(dcl.ValueOrEmptyBool(resource.PrivateIPGoogleAccess))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetLogConfig(ComputeBetaSubnetworkLogConfigToProto(resource.LogConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetEnableFlowLogs(dcl.ValueOrEmptyBool(resource.EnableFlowLogs))
	sSecondaryIPRanges := make([]*betapb.ComputeBetaSubnetworkSecondaryIPRanges, len(resource.SecondaryIPRanges))
	for i, r := range resource.SecondaryIPRanges {
		sSecondaryIPRanges[i] = ComputeBetaSubnetworkSecondaryIPRangesToProto(&r)
	}
	p.SetSecondaryIpRanges(sSecondaryIPRanges)

	return p
}

// applySubnetwork handles the gRPC request by passing it to the underlying Subnetwork Apply() method.
func (s *SubnetworkServer) applySubnetwork(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaSubnetworkRequest) (*betapb.ComputeBetaSubnetwork, error) {
	p := ProtoToSubnetwork(request.GetResource())
	res, err := c.ApplySubnetwork(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SubnetworkToProto(res)
	return r, nil
}

// applyComputeBetaSubnetwork handles the gRPC request by passing it to the underlying Subnetwork Apply() method.
func (s *SubnetworkServer) ApplyComputeBetaSubnetwork(ctx context.Context, request *betapb.ApplyComputeBetaSubnetworkRequest) (*betapb.ComputeBetaSubnetwork, error) {
	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applySubnetwork(ctx, cl, request)
}

// DeleteSubnetwork handles the gRPC request by passing it to the underlying Subnetwork Delete() method.
func (s *SubnetworkServer) DeleteComputeBetaSubnetwork(ctx context.Context, request *betapb.DeleteComputeBetaSubnetworkRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSubnetwork(ctx, ProtoToSubnetwork(request.GetResource()))

}

// ListComputeBetaSubnetwork handles the gRPC request by passing it to the underlying SubnetworkList() method.
func (s *SubnetworkServer) ListComputeBetaSubnetwork(ctx context.Context, request *betapb.ListComputeBetaSubnetworkRequest) (*betapb.ListComputeBetaSubnetworkResponse, error) {
	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSubnetwork(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaSubnetwork
	for _, r := range resources.Items {
		rp := SubnetworkToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaSubnetworkResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigSubnetwork(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
