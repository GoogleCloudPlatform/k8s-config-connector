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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// SubnetworkServer implements the gRPC interface for Subnetwork.
type SubnetworkServer struct{}

// ProtoToSubnetworkPurposeEnum converts a SubnetworkPurposeEnum enum from its proto representation.
func ProtoToComputeSubnetworkPurposeEnum(e computepb.ComputeSubnetworkPurposeEnum) *compute.SubnetworkPurposeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSubnetworkPurposeEnum_name[int32(e)]; ok {
		e := compute.SubnetworkPurposeEnum(n[len("ComputeSubnetworkPurposeEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkRoleEnum converts a SubnetworkRoleEnum enum from its proto representation.
func ProtoToComputeSubnetworkRoleEnum(e computepb.ComputeSubnetworkRoleEnum) *compute.SubnetworkRoleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSubnetworkRoleEnum_name[int32(e)]; ok {
		e := compute.SubnetworkRoleEnum(n[len("ComputeSubnetworkRoleEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkLogConfigAggregationIntervalEnum converts a SubnetworkLogConfigAggregationIntervalEnum enum from its proto representation.
func ProtoToComputeSubnetworkLogConfigAggregationIntervalEnum(e computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum) *compute.SubnetworkLogConfigAggregationIntervalEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum_name[int32(e)]; ok {
		e := compute.SubnetworkLogConfigAggregationIntervalEnum(n[len("ComputeSubnetworkLogConfigAggregationIntervalEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkLogConfigMetadataEnum converts a SubnetworkLogConfigMetadataEnum enum from its proto representation.
func ProtoToComputeSubnetworkLogConfigMetadataEnum(e computepb.ComputeSubnetworkLogConfigMetadataEnum) *compute.SubnetworkLogConfigMetadataEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSubnetworkLogConfigMetadataEnum_name[int32(e)]; ok {
		e := compute.SubnetworkLogConfigMetadataEnum(n[len("ComputeSubnetworkLogConfigMetadataEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubnetworkSecondaryIPRanges converts a SubnetworkSecondaryIPRanges object from its proto representation.
func ProtoToComputeSubnetworkSecondaryIPRanges(p *computepb.ComputeSubnetworkSecondaryIPRanges) *compute.SubnetworkSecondaryIPRanges {
	if p == nil {
		return nil
	}
	obj := &compute.SubnetworkSecondaryIPRanges{
		RangeName:   dcl.StringOrNil(p.GetRangeName()),
		IPCidrRange: dcl.StringOrNil(p.GetIpCidrRange()),
	}
	return obj
}

// ProtoToSubnetworkLogConfig converts a SubnetworkLogConfig object from its proto representation.
func ProtoToComputeSubnetworkLogConfig(p *computepb.ComputeSubnetworkLogConfig) *compute.SubnetworkLogConfig {
	if p == nil {
		return nil
	}
	obj := &compute.SubnetworkLogConfig{
		AggregationInterval: ProtoToComputeSubnetworkLogConfigAggregationIntervalEnum(p.GetAggregationInterval()),
		FlowSampling:        dcl.Float64OrNil(p.GetFlowSampling()),
		Metadata:            ProtoToComputeSubnetworkLogConfigMetadataEnum(p.GetMetadata()),
	}
	return obj
}

// ProtoToSubnetwork converts a Subnetwork resource from its proto representation.
func ProtoToSubnetwork(p *computepb.ComputeSubnetwork) *compute.Subnetwork {
	obj := &compute.Subnetwork{
		CreationTimestamp:     dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		GatewayAddress:        dcl.StringOrNil(p.GetGatewayAddress()),
		IPCidrRange:           dcl.StringOrNil(p.GetIpCidrRange()),
		Name:                  dcl.StringOrNil(p.GetName()),
		Network:               dcl.StringOrNil(p.GetNetwork()),
		Fingerprint:           dcl.StringOrNil(p.GetFingerprint()),
		Purpose:               ProtoToComputeSubnetworkPurposeEnum(p.GetPurpose()),
		Role:                  ProtoToComputeSubnetworkRoleEnum(p.GetRole()),
		PrivateIPGoogleAccess: dcl.Bool(p.GetPrivateIpGoogleAccess()),
		Region:                dcl.StringOrNil(p.GetRegion()),
		LogConfig:             ProtoToComputeSubnetworkLogConfig(p.GetLogConfig()),
		Project:               dcl.StringOrNil(p.GetProject()),
		SelfLink:              dcl.StringOrNil(p.GetSelfLink()),
		EnableFlowLogs:        dcl.Bool(p.GetEnableFlowLogs()),
	}
	for _, r := range p.GetSecondaryIpRanges() {
		obj.SecondaryIPRanges = append(obj.SecondaryIPRanges, *ProtoToComputeSubnetworkSecondaryIPRanges(r))
	}
	return obj
}

// SubnetworkPurposeEnumToProto converts a SubnetworkPurposeEnum enum to its proto representation.
func ComputeSubnetworkPurposeEnumToProto(e *compute.SubnetworkPurposeEnum) computepb.ComputeSubnetworkPurposeEnum {
	if e == nil {
		return computepb.ComputeSubnetworkPurposeEnum(0)
	}
	if v, ok := computepb.ComputeSubnetworkPurposeEnum_value["SubnetworkPurposeEnum"+string(*e)]; ok {
		return computepb.ComputeSubnetworkPurposeEnum(v)
	}
	return computepb.ComputeSubnetworkPurposeEnum(0)
}

// SubnetworkRoleEnumToProto converts a SubnetworkRoleEnum enum to its proto representation.
func ComputeSubnetworkRoleEnumToProto(e *compute.SubnetworkRoleEnum) computepb.ComputeSubnetworkRoleEnum {
	if e == nil {
		return computepb.ComputeSubnetworkRoleEnum(0)
	}
	if v, ok := computepb.ComputeSubnetworkRoleEnum_value["SubnetworkRoleEnum"+string(*e)]; ok {
		return computepb.ComputeSubnetworkRoleEnum(v)
	}
	return computepb.ComputeSubnetworkRoleEnum(0)
}

// SubnetworkLogConfigAggregationIntervalEnumToProto converts a SubnetworkLogConfigAggregationIntervalEnum enum to its proto representation.
func ComputeSubnetworkLogConfigAggregationIntervalEnumToProto(e *compute.SubnetworkLogConfigAggregationIntervalEnum) computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum {
	if e == nil {
		return computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum(0)
	}
	if v, ok := computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum_value["SubnetworkLogConfigAggregationIntervalEnum"+string(*e)]; ok {
		return computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum(v)
	}
	return computepb.ComputeSubnetworkLogConfigAggregationIntervalEnum(0)
}

// SubnetworkLogConfigMetadataEnumToProto converts a SubnetworkLogConfigMetadataEnum enum to its proto representation.
func ComputeSubnetworkLogConfigMetadataEnumToProto(e *compute.SubnetworkLogConfigMetadataEnum) computepb.ComputeSubnetworkLogConfigMetadataEnum {
	if e == nil {
		return computepb.ComputeSubnetworkLogConfigMetadataEnum(0)
	}
	if v, ok := computepb.ComputeSubnetworkLogConfigMetadataEnum_value["SubnetworkLogConfigMetadataEnum"+string(*e)]; ok {
		return computepb.ComputeSubnetworkLogConfigMetadataEnum(v)
	}
	return computepb.ComputeSubnetworkLogConfigMetadataEnum(0)
}

// SubnetworkSecondaryIPRangesToProto converts a SubnetworkSecondaryIPRanges object to its proto representation.
func ComputeSubnetworkSecondaryIPRangesToProto(o *compute.SubnetworkSecondaryIPRanges) *computepb.ComputeSubnetworkSecondaryIPRanges {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSubnetworkSecondaryIPRanges{}
	p.SetRangeName(dcl.ValueOrEmptyString(o.RangeName))
	p.SetIpCidrRange(dcl.ValueOrEmptyString(o.IPCidrRange))
	return p
}

// SubnetworkLogConfigToProto converts a SubnetworkLogConfig object to its proto representation.
func ComputeSubnetworkLogConfigToProto(o *compute.SubnetworkLogConfig) *computepb.ComputeSubnetworkLogConfig {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSubnetworkLogConfig{}
	p.SetAggregationInterval(ComputeSubnetworkLogConfigAggregationIntervalEnumToProto(o.AggregationInterval))
	p.SetFlowSampling(dcl.ValueOrEmptyDouble(o.FlowSampling))
	p.SetMetadata(ComputeSubnetworkLogConfigMetadataEnumToProto(o.Metadata))
	return p
}

// SubnetworkToProto converts a Subnetwork resource to its proto representation.
func SubnetworkToProto(resource *compute.Subnetwork) *computepb.ComputeSubnetwork {
	p := &computepb.ComputeSubnetwork{}
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetGatewayAddress(dcl.ValueOrEmptyString(resource.GatewayAddress))
	p.SetIpCidrRange(dcl.ValueOrEmptyString(resource.IPCidrRange))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetPurpose(ComputeSubnetworkPurposeEnumToProto(resource.Purpose))
	p.SetRole(ComputeSubnetworkRoleEnumToProto(resource.Role))
	p.SetPrivateIpGoogleAccess(dcl.ValueOrEmptyBool(resource.PrivateIPGoogleAccess))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetLogConfig(ComputeSubnetworkLogConfigToProto(resource.LogConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetEnableFlowLogs(dcl.ValueOrEmptyBool(resource.EnableFlowLogs))
	sSecondaryIPRanges := make([]*computepb.ComputeSubnetworkSecondaryIPRanges, len(resource.SecondaryIPRanges))
	for i, r := range resource.SecondaryIPRanges {
		sSecondaryIPRanges[i] = ComputeSubnetworkSecondaryIPRangesToProto(&r)
	}
	p.SetSecondaryIpRanges(sSecondaryIPRanges)

	return p
}

// applySubnetwork handles the gRPC request by passing it to the underlying Subnetwork Apply() method.
func (s *SubnetworkServer) applySubnetwork(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeSubnetworkRequest) (*computepb.ComputeSubnetwork, error) {
	p := ProtoToSubnetwork(request.GetResource())
	res, err := c.ApplySubnetwork(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SubnetworkToProto(res)
	return r, nil
}

// applyComputeSubnetwork handles the gRPC request by passing it to the underlying Subnetwork Apply() method.
func (s *SubnetworkServer) ApplyComputeSubnetwork(ctx context.Context, request *computepb.ApplyComputeSubnetworkRequest) (*computepb.ComputeSubnetwork, error) {
	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applySubnetwork(ctx, cl, request)
}

// DeleteSubnetwork handles the gRPC request by passing it to the underlying Subnetwork Delete() method.
func (s *SubnetworkServer) DeleteComputeSubnetwork(ctx context.Context, request *computepb.DeleteComputeSubnetworkRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSubnetwork(ctx, ProtoToSubnetwork(request.GetResource()))

}

// ListComputeSubnetwork handles the gRPC request by passing it to the underlying SubnetworkList() method.
func (s *SubnetworkServer) ListComputeSubnetwork(ctx context.Context, request *computepb.ListComputeSubnetworkRequest) (*computepb.ListComputeSubnetworkResponse, error) {
	cl, err := createConfigSubnetwork(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSubnetwork(ctx, request.GetProject(), request.GetRegion())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeSubnetwork
	for _, r := range resources.Items {
		rp := SubnetworkToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeSubnetworkResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigSubnetwork(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
