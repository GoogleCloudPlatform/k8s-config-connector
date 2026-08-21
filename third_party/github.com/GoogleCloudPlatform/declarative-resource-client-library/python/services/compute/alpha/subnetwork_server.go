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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/alpha/compute_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
)

// SubnetworkServer implements the gRPC interface for Subnetwork.
type SubnetworkServer struct{}

// ProtoToSubnetworkPurposeEnum converts a SubnetworkPurposeEnum enum from its proto representation.
func ProtoToComputeAlphaSubnetworkPurposeEnum(e alphapb.ComputeAlphaSubnetworkPurposeEnum) *alpha.SubnetworkPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaSubnetworkPurposeEnum_name[int32(e)]; ok {
		e := alpha.SubnetworkPurposeEnum(n[len("ComputeAlphaSubnetworkPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkRoleEnum converts a SubnetworkRoleEnum enum from its proto representation.
func ProtoToComputeAlphaSubnetworkRoleEnum(e alphapb.ComputeAlphaSubnetworkRoleEnum) *alpha.SubnetworkRoleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaSubnetworkRoleEnum_name[int32(e)]; ok {
		e := alpha.SubnetworkRoleEnum(n[len("ComputeAlphaSubnetworkRoleEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkLogConfigAggregationIntervalEnum converts a SubnetworkLogConfigAggregationIntervalEnum enum from its proto representation.
func ProtoToComputeAlphaSubnetworkLogConfigAggregationIntervalEnum(e alphapb.ComputeAlphaSubnetworkLogConfigAggregationIntervalEnum) *alpha.SubnetworkLogConfigAggregationIntervalEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaSubnetworkLogConfigAggregationIntervalEnum_name[int32(e)]; ok {
		e := alpha.SubnetworkLogConfigAggregationIntervalEnum(n[len("ComputeAlphaSubnetworkLogConfigAggregationIntervalEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkLogConfigMetadataEnum converts a SubnetworkLogConfigMetadataEnum enum from its proto representation.
func ProtoToComputeAlphaSubnetworkLogConfigMetadataEnum(e alphapb.ComputeAlphaSubnetworkLogConfigMetadataEnum) *alpha.SubnetworkLogConfigMetadataEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaSubnetworkLogConfigMetadataEnum_name[int32(e)]; ok {
		e := alpha.SubnetworkLogConfigMetadataEnum(n[len("ComputeAlphaSubnetworkLogConfigMetadataEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkSecondaryIPRanges converts a SubnetworkSecondaryIPRanges object from its proto representation.
func ProtoToComputeAlphaSubnetworkSecondaryIPRanges(p *alphapb.ComputeAlphaSubnetworkSecondaryIPRanges) *alpha.SubnetworkSecondaryIPRanges {
	if p == nil {
		return nil
	}
	obj := &alpha.SubnetworkSecondaryIPRanges{
		RangeName:   dcl.StringOrNil(p.GetRangeName()),
		IPCidrRange: dcl.StringOrNil(p.GetIpCidrRange()),
	}
	return obj
}

// ProtoToSubnetworkLogConfig converts a SubnetworkLogConfig object from its proto representation.
func ProtoToComputeAlphaSubnetworkLogConfig(p *alphapb.ComputeAlphaSubnetworkLogConfig) *alpha.SubnetworkLogConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.SubnetworkLogConfig{
		AggregationInterval: ProtoToComputeAlphaSubnetworkLogConfigAggregationIntervalEnum(p.GetAggregationInterval()),
		FlowSampling:        dcl.Float64OrNil(p.GetFlowSampling()),
		Metadata:            ProtoToComputeAlphaSubnetworkLogConfigMetadataEnum(p.GetMetadata()),
	}
	return obj
}

// ProtoToSubnetwork converts a Subnetwork resource from its proto representation.
func ProtoToSubnetwork(p *alphapb.ComputeAlphaSubnetwork) *alpha.Subnetwork {
	obj := &alpha.Subnetwork{
		CreationTimestamp:     dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		GatewayAddress:        dcl.StringOrNil(p.GetGatewayAddress()),
		IPCidrRange:           dcl.StringOrNil(p.GetIpCidrRange()),
		Name:                  dcl.StringOrNil(p.GetName()),
		Network:               dcl.StringOrNil(p.GetNetwork()),
		Fingerprint:           dcl.StringOrNil(p.GetFingerprint()),
		Purpose:               ProtoToComputeAlphaSubnetworkPurposeEnum(p.GetPurpose()),
		Role:                  ProtoToComputeAlphaSubnetworkRoleEnum(p.GetRole()),
		PrivateIPGoogleAccess: dcl.Bool(p.GetPrivateIpGoogleAccess()),
		Region:                dcl.StringOrNil(p.GetRegion()),
		LogConfig:             ProtoToComputeAlphaSubnetworkLogConfig(p.GetLogConfig()),
		Project:               dcl.StringOrNil(p.GetProject()),
		SelfLink:              dcl.StringOrNil(p.GetSelfLink()),
		EnableFlowLogs:        dcl.Bool(p.GetEnableFlowLogs()),
	}
	for _, r := range p.GetSecondaryIpRanges() {
		obj.SecondaryIPRanges = append(obj.SecondaryIPRanges, *ProtoToComputeAlphaSubnetworkSecondaryIPRanges(r))
	}
	return obj
}

// SubnetworkPurposeEnumToProto converts a SubnetworkPurposeEnum enum to its proto representation.
func ComputeAlphaSubnetworkPurposeEnumToProto(e *alpha.SubnetworkPurposeEnum) alphapb.ComputeAlphaSubnetworkPurposeEnum {
	if e == nil {
		return alphapb.ComputeAlphaSubnetworkPurposeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaSubnetworkPurposeEnum_value["SubnetworkPurposeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaSubnetworkPurposeEnum(v)
	}
	return alphapb.ComputeAlphaSubnetworkPurposeEnum(0)
}

// SubnetworkRoleEnumToProto converts a SubnetworkRoleEnum enum to its proto representation.
func ComputeAlphaSubnetworkRoleEnumToProto(e *alpha.SubnetworkRoleEnum) alphapb.ComputeAlphaSubnetworkRoleEnum {
	if e == nil {
		return alphapb.ComputeAlphaSubnetworkRoleEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaSubnetworkRoleEnum_value["SubnetworkRoleEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaSubnetworkRoleEnum(v)
	}
	return alphapb.ComputeAlphaSubnetworkRoleEnum(0)
}

// SubnetworkLogConfigAggregationIntervalEnumToProto converts a SubnetworkLogConfigAggregationIntervalEnum enum to its proto representation.
func ComputeAlphaSubnetworkLogConfigAggregationIntervalEnumToProto(e *alpha.SubnetworkLogConfigAggregationIntervalEnum) alphapb.ComputeAlphaSubnetworkLogConfigAggregationIntervalEnum {
	if e == nil {
		return alphapb.ComputeAlphaSubnetworkLogConfigAggregationIntervalEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaSubnetworkLogConfigAggregationIntervalEnum_value["SubnetworkLogConfigAggregationIntervalEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaSubnetworkLogConfigAggregationIntervalEnum(v)
	}
	return alphapb.ComputeAlphaSubnetworkLogConfigAggregationIntervalEnum(0)
}

// SubnetworkLogConfigMetadataEnumToProto converts a SubnetworkLogConfigMetadataEnum enum to its proto representation.
func ComputeAlphaSubnetworkLogConfigMetadataEnumToProto(e *alpha.SubnetworkLogConfigMetadataEnum) alphapb.ComputeAlphaSubnetworkLogConfigMetadataEnum {
	if e == nil {
		return alphapb.ComputeAlphaSubnetworkLogConfigMetadataEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaSubnetworkLogConfigMetadataEnum_value["SubnetworkLogConfigMetadataEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaSubnetworkLogConfigMetadataEnum(v)
	}
	return alphapb.ComputeAlphaSubnetworkLogConfigMetadataEnum(0)
}

// SubnetworkSecondaryIPRangesToProto converts a SubnetworkSecondaryIPRanges object to its proto representation.
func ComputeAlphaSubnetworkSecondaryIPRangesToProto(o *alpha.SubnetworkSecondaryIPRanges) *alphapb.ComputeAlphaSubnetworkSecondaryIPRanges {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaSubnetworkSecondaryIPRanges{}
	p.SetRangeName(dcl.ValueOrEmptyString(o.RangeName))
	p.SetIpCidrRange(dcl.ValueOrEmptyString(o.IPCidrRange))
	return p
}

// SubnetworkLogConfigToProto converts a SubnetworkLogConfig object to its proto representation.
func ComputeAlphaSubnetworkLogConfigToProto(o *alpha.SubnetworkLogConfig) *alphapb.ComputeAlphaSubnetworkLogConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaSubnetworkLogConfig{}
	p.SetAggregationInterval(ComputeAlphaSubnetworkLogConfigAggregationIntervalEnumToProto(o.AggregationInterval))
	p.SetFlowSampling(dcl.ValueOrEmptyDouble(o.FlowSampling))
	p.SetMetadata(ComputeAlphaSubnetworkLogConfigMetadataEnumToProto(o.Metadata))
	return p
}

// SubnetworkToProto converts a Subnetwork resource to its proto representation.
func SubnetworkToProto(resource *alpha.Subnetwork) *alphapb.ComputeAlphaSubnetwork {
	p := &alphapb.ComputeAlphaSubnetwork{}
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGatewayAddress(dcl.ValueOrEmptyString(resource.GatewayAddress))
	p.SetIpCidrRange(dcl.ValueOrEmptyString(resource.IPCidrRange))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetPurpose(ComputeAlphaSubnetworkPurposeEnumToProto(resource.Purpose))
	p.SetRole(ComputeAlphaSubnetworkRoleEnumToProto(resource.Role))
	p.SetPrivateIpGoogleAccess(dcl.ValueOrEmptyBool(resource.PrivateIPGoogleAccess))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetLogConfig(ComputeAlphaSubnetworkLogConfigToProto(resource.LogConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetEnableFlowLogs(dcl.ValueOrEmptyBool(resource.EnableFlowLogs))
	sSecondaryIPRanges := make([]*alphapb.ComputeAlphaSubnetworkSecondaryIPRanges, len(resource.SecondaryIPRanges))
	for i, r := range resource.SecondaryIPRanges {
		sSecondaryIPRanges[i] = ComputeAlphaSubnetworkSecondaryIPRangesToProto(&r)
	}
	p.SetSecondaryIpRanges(sSecondaryIPRanges)

	return p
}

// applySubnetwork handles the gRPC request by passing it to the underlying Subnetwork Apply() method.
func (s *SubnetworkServer) applySubnetwork(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaSubnetworkRequest) (*alphapb.ComputeAlphaSubnetwork, error) {
	p := ProtoToSubnetwork(request.GetResource())
	res, err := c.ApplySubnetwork(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SubnetworkToProto(res)
	return r, nil
}

// applyComputeAlphaSubnetwork handles the gRPC request by passing it to the underlying Subnetwork Apply() method.
func (s *SubnetworkServer) ApplyComputeAlphaSubnetwork(ctx context.Context, request *alphapb.ApplyComputeAlphaSubnetworkRequest) (*alphapb.ComputeAlphaSubnetwork, error) {
	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applySubnetwork(ctx, cl, request)
}

// DeleteSubnetwork handles the gRPC request by passing it to the underlying Subnetwork Delete() method.
func (s *SubnetworkServer) DeleteComputeAlphaSubnetwork(ctx context.Context, request *alphapb.DeleteComputeAlphaSubnetworkRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSubnetwork(ctx, ProtoToSubnetwork(request.GetResource()))

}

// ListComputeAlphaSubnetwork handles the gRPC request by passing it to the underlying SubnetworkList() method.
func (s *SubnetworkServer) ListComputeAlphaSubnetwork(ctx context.Context, request *alphapb.ListComputeAlphaSubnetworkRequest) (*alphapb.ListComputeAlphaSubnetworkResponse, error) {
	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSubnetwork(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaSubnetwork
	for _, r := range resources.Items {
		rp := SubnetworkToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaSubnetworkResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigSubnetwork(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
