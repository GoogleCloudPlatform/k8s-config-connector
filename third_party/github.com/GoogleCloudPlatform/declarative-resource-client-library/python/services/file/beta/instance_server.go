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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/file/beta/file_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file/beta"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToFileBetaInstanceStateEnum(e betapb.FileBetaInstanceStateEnum) *beta.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FileBetaInstanceStateEnum_name[int32(e)]; ok {
		e := beta.InstanceStateEnum(n[len("FileBetaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTierEnum converts a InstanceTierEnum enum from its proto representation.
func ProtoToFileBetaInstanceTierEnum(e betapb.FileBetaInstanceTierEnum) *beta.InstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FileBetaInstanceTierEnum_name[int32(e)]; ok {
		e := beta.InstanceTierEnum(n[len("FileBetaInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsAccessModeEnum converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum from its proto representation.
func ProtoToFileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(e betapb.FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum) *beta.InstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum_name[int32(e)]; ok {
		e := beta.InstanceFileSharesNfsExportOptionsAccessModeEnum(n[len("FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsSquashModeEnum converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum from its proto representation.
func ProtoToFileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(e betapb.FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum) *beta.InstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum_name[int32(e)]; ok {
		e := beta.InstanceFileSharesNfsExportOptionsSquashModeEnum(n[len("FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworksModesEnum converts a InstanceNetworksModesEnum enum from its proto representation.
func ProtoToFileBetaInstanceNetworksModesEnum(e betapb.FileBetaInstanceNetworksModesEnum) *beta.InstanceNetworksModesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FileBetaInstanceNetworksModesEnum_name[int32(e)]; ok {
		e := beta.InstanceNetworksModesEnum(n[len("FileBetaInstanceNetworksModesEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileShares converts a InstanceFileShares resource from its proto representation.
func ProtoToFileBetaInstanceFileShares(p *betapb.FileBetaInstanceFileShares) *beta.InstanceFileShares {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceFileShares{
		Name:         dcl.StringOrNil(p.Name),
		CapacityGb:   dcl.Int64OrNil(p.CapacityGb),
		SourceBackup: dcl.StringOrNil(p.SourceBackup),
	}
	for _, r := range p.GetNfsExportOptions() {
		obj.NfsExportOptions = append(obj.NfsExportOptions, *ProtoToFileBetaInstanceFileSharesNfsExportOptions(r))
	}
	return obj
}

// ProtoToInstanceFileSharesNfsExportOptions converts a InstanceFileSharesNfsExportOptions resource from its proto representation.
func ProtoToFileBetaInstanceFileSharesNfsExportOptions(p *betapb.FileBetaInstanceFileSharesNfsExportOptions) *beta.InstanceFileSharesNfsExportOptions {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceFileSharesNfsExportOptions{
		AccessMode: ProtoToFileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(p.GetAccessMode()),
		SquashMode: ProtoToFileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(p.GetSquashMode()),
		AnonUid:    dcl.Int64OrNil(p.AnonUid),
		AnonGid:    dcl.Int64OrNil(p.AnonGid),
	}
	for _, r := range p.GetIpRanges() {
		obj.IPRanges = append(obj.IPRanges, r)
	}
	return obj
}

// ProtoToInstanceNetworks converts a InstanceNetworks resource from its proto representation.
func ProtoToFileBetaInstanceNetworks(p *betapb.FileBetaInstanceNetworks) *beta.InstanceNetworks {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworks{
		Network:         dcl.StringOrNil(p.Network),
		ReservedIPRange: dcl.StringOrNil(p.ReservedIpRange),
	}
	for _, r := range p.GetModes() {
		obj.Modes = append(obj.Modes, *ProtoToFileBetaInstanceNetworksModesEnum(r))
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.FileBetaInstance) *beta.Instance {
	obj := &beta.Instance{
		Name:          dcl.StringOrNil(p.Name),
		Description:   dcl.StringOrNil(p.Description),
		State:         ProtoToFileBetaInstanceStateEnum(p.GetState()),
		StatusMessage: dcl.StringOrNil(p.StatusMessage),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Tier:          ProtoToFileBetaInstanceTierEnum(p.GetTier()),
		Etag:          dcl.StringOrNil(p.Etag),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetFileShares() {
		obj.FileShares = append(obj.FileShares, *ProtoToFileBetaInstanceFileShares(r))
	}
	for _, r := range p.GetNetworks() {
		obj.Networks = append(obj.Networks, *ProtoToFileBetaInstanceNetworks(r))
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func FileBetaInstanceStateEnumToProto(e *beta.InstanceStateEnum) betapb.FileBetaInstanceStateEnum {
	if e == nil {
		return betapb.FileBetaInstanceStateEnum(0)
	}
	if v, ok := betapb.FileBetaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return betapb.FileBetaInstanceStateEnum(v)
	}
	return betapb.FileBetaInstanceStateEnum(0)
}

// InstanceTierEnumToProto converts a InstanceTierEnum enum to its proto representation.
func FileBetaInstanceTierEnumToProto(e *beta.InstanceTierEnum) betapb.FileBetaInstanceTierEnum {
	if e == nil {
		return betapb.FileBetaInstanceTierEnum(0)
	}
	if v, ok := betapb.FileBetaInstanceTierEnum_value["InstanceTierEnum"+string(*e)]; ok {
		return betapb.FileBetaInstanceTierEnum(v)
	}
	return betapb.FileBetaInstanceTierEnum(0)
}

// InstanceFileSharesNfsExportOptionsAccessModeEnumToProto converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum to its proto representation.
func FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(e *beta.InstanceFileSharesNfsExportOptionsAccessModeEnum) betapb.FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == nil {
		return betapb.FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
	}
	if v, ok := betapb.FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum_value["InstanceFileSharesNfsExportOptionsAccessModeEnum"+string(*e)]; ok {
		return betapb.FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(v)
	}
	return betapb.FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
}

// InstanceFileSharesNfsExportOptionsSquashModeEnumToProto converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum to its proto representation.
func FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(e *beta.InstanceFileSharesNfsExportOptionsSquashModeEnum) betapb.FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == nil {
		return betapb.FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
	}
	if v, ok := betapb.FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum_value["InstanceFileSharesNfsExportOptionsSquashModeEnum"+string(*e)]; ok {
		return betapb.FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(v)
	}
	return betapb.FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
}

// InstanceNetworksModesEnumToProto converts a InstanceNetworksModesEnum enum to its proto representation.
func FileBetaInstanceNetworksModesEnumToProto(e *beta.InstanceNetworksModesEnum) betapb.FileBetaInstanceNetworksModesEnum {
	if e == nil {
		return betapb.FileBetaInstanceNetworksModesEnum(0)
	}
	if v, ok := betapb.FileBetaInstanceNetworksModesEnum_value["InstanceNetworksModesEnum"+string(*e)]; ok {
		return betapb.FileBetaInstanceNetworksModesEnum(v)
	}
	return betapb.FileBetaInstanceNetworksModesEnum(0)
}

// InstanceFileSharesToProto converts a InstanceFileShares resource to its proto representation.
func FileBetaInstanceFileSharesToProto(o *beta.InstanceFileShares) *betapb.FileBetaInstanceFileShares {
	if o == nil {
		return nil
	}
	p := &betapb.FileBetaInstanceFileShares{
		Name:         dcl.ValueOrEmptyString(o.Name),
		CapacityGb:   dcl.ValueOrEmptyInt64(o.CapacityGb),
		SourceBackup: dcl.ValueOrEmptyString(o.SourceBackup),
	}
	for _, r := range o.NfsExportOptions {
		p.NfsExportOptions = append(p.NfsExportOptions, FileBetaInstanceFileSharesNfsExportOptionsToProto(&r))
	}
	return p
}

// InstanceFileSharesNfsExportOptionsToProto converts a InstanceFileSharesNfsExportOptions resource to its proto representation.
func FileBetaInstanceFileSharesNfsExportOptionsToProto(o *beta.InstanceFileSharesNfsExportOptions) *betapb.FileBetaInstanceFileSharesNfsExportOptions {
	if o == nil {
		return nil
	}
	p := &betapb.FileBetaInstanceFileSharesNfsExportOptions{
		AccessMode: FileBetaInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(o.AccessMode),
		SquashMode: FileBetaInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(o.SquashMode),
		AnonUid:    dcl.ValueOrEmptyInt64(o.AnonUid),
		AnonGid:    dcl.ValueOrEmptyInt64(o.AnonGid),
	}
	for _, r := range o.IPRanges {
		p.IpRanges = append(p.IpRanges, r)
	}
	return p
}

// InstanceNetworksToProto converts a InstanceNetworks resource to its proto representation.
func FileBetaInstanceNetworksToProto(o *beta.InstanceNetworks) *betapb.FileBetaInstanceNetworks {
	if o == nil {
		return nil
	}
	p := &betapb.FileBetaInstanceNetworks{
		Network:         dcl.ValueOrEmptyString(o.Network),
		ReservedIpRange: dcl.ValueOrEmptyString(o.ReservedIPRange),
	}
	for _, r := range o.Modes {
		p.Modes = append(p.Modes, betapb.FileBetaInstanceNetworksModesEnum(betapb.FileBetaInstanceNetworksModesEnum_value[string(r)]))
	}
	for _, r := range o.IPAddresses {
		p.IpAddresses = append(p.IpAddresses, r)
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *beta.Instance) *betapb.FileBetaInstance {
	p := &betapb.FileBetaInstance{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		Description:   dcl.ValueOrEmptyString(resource.Description),
		State:         FileBetaInstanceStateEnumToProto(resource.State),
		StatusMessage: dcl.ValueOrEmptyString(resource.StatusMessage),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		Tier:          FileBetaInstanceTierEnumToProto(resource.Tier),
		Etag:          dcl.ValueOrEmptyString(resource.Etag),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.FileShares {
		p.FileShares = append(p.FileShares, FileBetaInstanceFileSharesToProto(&r))
	}
	for _, r := range resource.Networks {
		p.Networks = append(p.Networks, FileBetaInstanceNetworksToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplyFileBetaInstanceRequest) (*betapb.FileBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyFileBetaInstance(ctx context.Context, request *betapb.ApplyFileBetaInstanceRequest) (*betapb.FileBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteFileBetaInstance(ctx context.Context, request *betapb.DeleteFileBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListFileBetaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListFileBetaInstance(ctx context.Context, request *betapb.ListFileBetaInstanceRequest) (*betapb.ListFileBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, ProtoToInstance(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.FileBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListFileBetaInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
