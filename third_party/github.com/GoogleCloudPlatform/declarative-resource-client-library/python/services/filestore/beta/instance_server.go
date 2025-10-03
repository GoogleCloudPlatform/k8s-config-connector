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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/filestore/beta/filestore_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/beta"
)

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToFilestoreBetaInstanceStateEnum(e betapb.FilestoreBetaInstanceStateEnum) *beta.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FilestoreBetaInstanceStateEnum_name[int32(e)]; ok {
		e := beta.InstanceStateEnum(n[len("FilestoreBetaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTierEnum converts a InstanceTierEnum enum from its proto representation.
func ProtoToFilestoreBetaInstanceTierEnum(e betapb.FilestoreBetaInstanceTierEnum) *beta.InstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FilestoreBetaInstanceTierEnum_name[int32(e)]; ok {
		e := beta.InstanceTierEnum(n[len("FilestoreBetaInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsAccessModeEnum converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum from its proto representation.
func ProtoToFilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(e betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum) *beta.InstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum_name[int32(e)]; ok {
		e := beta.InstanceFileSharesNfsExportOptionsAccessModeEnum(n[len("FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsSquashModeEnum converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum from its proto representation.
func ProtoToFilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(e betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum) *beta.InstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum_name[int32(e)]; ok {
		e := beta.InstanceFileSharesNfsExportOptionsSquashModeEnum(n[len("FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworksModesEnum converts a InstanceNetworksModesEnum enum from its proto representation.
func ProtoToFilestoreBetaInstanceNetworksModesEnum(e betapb.FilestoreBetaInstanceNetworksModesEnum) *beta.InstanceNetworksModesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FilestoreBetaInstanceNetworksModesEnum_name[int32(e)]; ok {
		e := beta.InstanceNetworksModesEnum(n[len("FilestoreBetaInstanceNetworksModesEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileShares converts a InstanceFileShares object from its proto representation.
func ProtoToFilestoreBetaInstanceFileShares(p *betapb.FilestoreBetaInstanceFileShares) *beta.InstanceFileShares {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceFileShares{
		Name:         dcl.StringOrNil(p.GetName()),
		CapacityGb:   dcl.Int64OrNil(p.GetCapacityGb()),
		SourceBackup: dcl.StringOrNil(p.GetSourceBackup()),
	}
	for _, r := range p.GetNfsExportOptions() {
		obj.NfsExportOptions = append(obj.NfsExportOptions, *ProtoToFilestoreBetaInstanceFileSharesNfsExportOptions(r))
	}
	return obj
}

// ProtoToInstanceFileSharesNfsExportOptions converts a InstanceFileSharesNfsExportOptions object from its proto representation.
func ProtoToFilestoreBetaInstanceFileSharesNfsExportOptions(p *betapb.FilestoreBetaInstanceFileSharesNfsExportOptions) *beta.InstanceFileSharesNfsExportOptions {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceFileSharesNfsExportOptions{
		AccessMode: ProtoToFilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(p.GetAccessMode()),
		SquashMode: ProtoToFilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(p.GetSquashMode()),
		AnonUid:    dcl.Int64OrNil(p.GetAnonUid()),
		AnonGid:    dcl.Int64OrNil(p.GetAnonGid()),
	}
	for _, r := range p.GetIpRanges() {
		obj.IPRanges = append(obj.IPRanges, r)
	}
	return obj
}

// ProtoToInstanceNetworks converts a InstanceNetworks object from its proto representation.
func ProtoToFilestoreBetaInstanceNetworks(p *betapb.FilestoreBetaInstanceNetworks) *beta.InstanceNetworks {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworks{
		Network:         dcl.StringOrNil(p.GetNetwork()),
		ReservedIPRange: dcl.StringOrNil(p.GetReservedIpRange()),
	}
	for _, r := range p.GetModes() {
		obj.Modes = append(obj.Modes, *ProtoToFilestoreBetaInstanceNetworksModesEnum(r))
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.FilestoreBetaInstance) *beta.Instance {
	obj := &beta.Instance{
		Name:          dcl.StringOrNil(p.GetName()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		State:         ProtoToFilestoreBetaInstanceStateEnum(p.GetState()),
		StatusMessage: dcl.StringOrNil(p.GetStatusMessage()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Tier:          ProtoToFilestoreBetaInstanceTierEnum(p.GetTier()),
		Etag:          dcl.StringOrNil(p.GetEtag()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetFileShares() {
		obj.FileShares = append(obj.FileShares, *ProtoToFilestoreBetaInstanceFileShares(r))
	}
	for _, r := range p.GetNetworks() {
		obj.Networks = append(obj.Networks, *ProtoToFilestoreBetaInstanceNetworks(r))
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func FilestoreBetaInstanceStateEnumToProto(e *beta.InstanceStateEnum) betapb.FilestoreBetaInstanceStateEnum {
	if e == nil {
		return betapb.FilestoreBetaInstanceStateEnum(0)
	}
	if v, ok := betapb.FilestoreBetaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return betapb.FilestoreBetaInstanceStateEnum(v)
	}
	return betapb.FilestoreBetaInstanceStateEnum(0)
}

// InstanceTierEnumToProto converts a InstanceTierEnum enum to its proto representation.
func FilestoreBetaInstanceTierEnumToProto(e *beta.InstanceTierEnum) betapb.FilestoreBetaInstanceTierEnum {
	if e == nil {
		return betapb.FilestoreBetaInstanceTierEnum(0)
	}
	if v, ok := betapb.FilestoreBetaInstanceTierEnum_value["InstanceTierEnum"+string(*e)]; ok {
		return betapb.FilestoreBetaInstanceTierEnum(v)
	}
	return betapb.FilestoreBetaInstanceTierEnum(0)
}

// InstanceFileSharesNfsExportOptionsAccessModeEnumToProto converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum to its proto representation.
func FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(e *beta.InstanceFileSharesNfsExportOptionsAccessModeEnum) betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == nil {
		return betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
	}
	if v, ok := betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum_value["InstanceFileSharesNfsExportOptionsAccessModeEnum"+string(*e)]; ok {
		return betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(v)
	}
	return betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
}

// InstanceFileSharesNfsExportOptionsSquashModeEnumToProto converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum to its proto representation.
func FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(e *beta.InstanceFileSharesNfsExportOptionsSquashModeEnum) betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == nil {
		return betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
	}
	if v, ok := betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum_value["InstanceFileSharesNfsExportOptionsSquashModeEnum"+string(*e)]; ok {
		return betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(v)
	}
	return betapb.FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
}

// InstanceNetworksModesEnumToProto converts a InstanceNetworksModesEnum enum to its proto representation.
func FilestoreBetaInstanceNetworksModesEnumToProto(e *beta.InstanceNetworksModesEnum) betapb.FilestoreBetaInstanceNetworksModesEnum {
	if e == nil {
		return betapb.FilestoreBetaInstanceNetworksModesEnum(0)
	}
	if v, ok := betapb.FilestoreBetaInstanceNetworksModesEnum_value["InstanceNetworksModesEnum"+string(*e)]; ok {
		return betapb.FilestoreBetaInstanceNetworksModesEnum(v)
	}
	return betapb.FilestoreBetaInstanceNetworksModesEnum(0)
}

// InstanceFileSharesToProto converts a InstanceFileShares object to its proto representation.
func FilestoreBetaInstanceFileSharesToProto(o *beta.InstanceFileShares) *betapb.FilestoreBetaInstanceFileShares {
	if o == nil {
		return nil
	}
	p := &betapb.FilestoreBetaInstanceFileShares{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCapacityGb(dcl.ValueOrEmptyInt64(o.CapacityGb))
	p.SetSourceBackup(dcl.ValueOrEmptyString(o.SourceBackup))
	sNfsExportOptions := make([]*betapb.FilestoreBetaInstanceFileSharesNfsExportOptions, len(o.NfsExportOptions))
	for i, r := range o.NfsExportOptions {
		sNfsExportOptions[i] = FilestoreBetaInstanceFileSharesNfsExportOptionsToProto(&r)
	}
	p.SetNfsExportOptions(sNfsExportOptions)
	return p
}

// InstanceFileSharesNfsExportOptionsToProto converts a InstanceFileSharesNfsExportOptions object to its proto representation.
func FilestoreBetaInstanceFileSharesNfsExportOptionsToProto(o *beta.InstanceFileSharesNfsExportOptions) *betapb.FilestoreBetaInstanceFileSharesNfsExportOptions {
	if o == nil {
		return nil
	}
	p := &betapb.FilestoreBetaInstanceFileSharesNfsExportOptions{}
	p.SetAccessMode(FilestoreBetaInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(o.AccessMode))
	p.SetSquashMode(FilestoreBetaInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(o.SquashMode))
	p.SetAnonUid(dcl.ValueOrEmptyInt64(o.AnonUid))
	p.SetAnonGid(dcl.ValueOrEmptyInt64(o.AnonGid))
	sIPRanges := make([]string, len(o.IPRanges))
	for i, r := range o.IPRanges {
		sIPRanges[i] = r
	}
	p.SetIpRanges(sIPRanges)
	return p
}

// InstanceNetworksToProto converts a InstanceNetworks object to its proto representation.
func FilestoreBetaInstanceNetworksToProto(o *beta.InstanceNetworks) *betapb.FilestoreBetaInstanceNetworks {
	if o == nil {
		return nil
	}
	p := &betapb.FilestoreBetaInstanceNetworks{}
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetReservedIpRange(dcl.ValueOrEmptyString(o.ReservedIPRange))
	sModes := make([]betapb.FilestoreBetaInstanceNetworksModesEnum, len(o.Modes))
	for i, r := range o.Modes {
		sModes[i] = betapb.FilestoreBetaInstanceNetworksModesEnum(betapb.FilestoreBetaInstanceNetworksModesEnum_value[string(r)])
	}
	p.SetModes(sModes)
	sIPAddresses := make([]string, len(o.IPAddresses))
	for i, r := range o.IPAddresses {
		sIPAddresses[i] = r
	}
	p.SetIpAddresses(sIPAddresses)
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *beta.Instance) *betapb.FilestoreBetaInstance {
	p := &betapb.FilestoreBetaInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(FilestoreBetaInstanceStateEnumToProto(resource.State))
	p.SetStatusMessage(dcl.ValueOrEmptyString(resource.StatusMessage))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetTier(FilestoreBetaInstanceTierEnumToProto(resource.Tier))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sFileShares := make([]*betapb.FilestoreBetaInstanceFileShares, len(resource.FileShares))
	for i, r := range resource.FileShares {
		sFileShares[i] = FilestoreBetaInstanceFileSharesToProto(&r)
	}
	p.SetFileShares(sFileShares)
	sNetworks := make([]*betapb.FilestoreBetaInstanceNetworks, len(resource.Networks))
	for i, r := range resource.Networks {
		sNetworks[i] = FilestoreBetaInstanceNetworksToProto(&r)
	}
	p.SetNetworks(sNetworks)

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplyFilestoreBetaInstanceRequest) (*betapb.FilestoreBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyFilestoreBetaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyFilestoreBetaInstance(ctx context.Context, request *betapb.ApplyFilestoreBetaInstanceRequest) (*betapb.FilestoreBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteFilestoreBetaInstance(ctx context.Context, request *betapb.DeleteFilestoreBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListFilestoreBetaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListFilestoreBetaInstance(ctx context.Context, request *betapb.ListFilestoreBetaInstanceRequest) (*betapb.ListFilestoreBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.FilestoreBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListFilestoreBetaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
