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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/filestore/alpha/filestore_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/alpha"
)

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceStateEnum(e alphapb.FilestoreAlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("FilestoreAlphaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTierEnum converts a InstanceTierEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceTierEnum(e alphapb.FilestoreAlphaInstanceTierEnum) *alpha.InstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceTierEnum_name[int32(e)]; ok {
		e := alpha.InstanceTierEnum(n[len("FilestoreAlphaInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsAccessModeEnum converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(e alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum) *alpha.InstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceFileSharesNfsExportOptionsAccessModeEnum(n[len("FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsSquashModeEnum converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(e alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum) *alpha.InstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceFileSharesNfsExportOptionsSquashModeEnum(n[len("FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworksModesEnum converts a InstanceNetworksModesEnum enum from its proto representation.
func ProtoToFilestoreAlphaInstanceNetworksModesEnum(e alphapb.FilestoreAlphaInstanceNetworksModesEnum) *alpha.InstanceNetworksModesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaInstanceNetworksModesEnum_name[int32(e)]; ok {
		e := alpha.InstanceNetworksModesEnum(n[len("FilestoreAlphaInstanceNetworksModesEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileShares converts a InstanceFileShares object from its proto representation.
func ProtoToFilestoreAlphaInstanceFileShares(p *alphapb.FilestoreAlphaInstanceFileShares) *alpha.InstanceFileShares {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFileShares{
		Name:         dcl.StringOrNil(p.GetName()),
		CapacityGb:   dcl.Int64OrNil(p.GetCapacityGb()),
		SourceBackup: dcl.StringOrNil(p.GetSourceBackup()),
	}
	for _, r := range p.GetNfsExportOptions() {
		obj.NfsExportOptions = append(obj.NfsExportOptions, *ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptions(r))
	}
	return obj
}

// ProtoToInstanceFileSharesNfsExportOptions converts a InstanceFileSharesNfsExportOptions object from its proto representation.
func ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptions(p *alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptions) *alpha.InstanceFileSharesNfsExportOptions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceFileSharesNfsExportOptions{
		AccessMode: ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(p.GetAccessMode()),
		SquashMode: ProtoToFilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(p.GetSquashMode()),
		AnonUid:    dcl.Int64OrNil(p.GetAnonUid()),
		AnonGid:    dcl.Int64OrNil(p.GetAnonGid()),
	}
	for _, r := range p.GetIpRanges() {
		obj.IPRanges = append(obj.IPRanges, r)
	}
	return obj
}

// ProtoToInstanceNetworks converts a InstanceNetworks object from its proto representation.
func ProtoToFilestoreAlphaInstanceNetworks(p *alphapb.FilestoreAlphaInstanceNetworks) *alpha.InstanceNetworks {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNetworks{
		Network:         dcl.StringOrNil(p.GetNetwork()),
		ReservedIPRange: dcl.StringOrNil(p.GetReservedIpRange()),
	}
	for _, r := range p.GetModes() {
		obj.Modes = append(obj.Modes, *ProtoToFilestoreAlphaInstanceNetworksModesEnum(r))
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.FilestoreAlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:          dcl.StringOrNil(p.GetName()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		State:         ProtoToFilestoreAlphaInstanceStateEnum(p.GetState()),
		StatusMessage: dcl.StringOrNil(p.GetStatusMessage()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Tier:          ProtoToFilestoreAlphaInstanceTierEnum(p.GetTier()),
		Etag:          dcl.StringOrNil(p.GetEtag()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetFileShares() {
		obj.FileShares = append(obj.FileShares, *ProtoToFilestoreAlphaInstanceFileShares(r))
	}
	for _, r := range p.GetNetworks() {
		obj.Networks = append(obj.Networks, *ProtoToFilestoreAlphaInstanceNetworks(r))
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func FilestoreAlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.FilestoreAlphaInstanceStateEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceStateEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceStateEnum(0)
}

// InstanceTierEnumToProto converts a InstanceTierEnum enum to its proto representation.
func FilestoreAlphaInstanceTierEnumToProto(e *alpha.InstanceTierEnum) alphapb.FilestoreAlphaInstanceTierEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceTierEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceTierEnum_value["InstanceTierEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceTierEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceTierEnum(0)
}

// InstanceFileSharesNfsExportOptionsAccessModeEnumToProto converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum to its proto representation.
func FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(e *alpha.InstanceFileSharesNfsExportOptionsAccessModeEnum) alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum_value["InstanceFileSharesNfsExportOptionsAccessModeEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
}

// InstanceFileSharesNfsExportOptionsSquashModeEnumToProto converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum to its proto representation.
func FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(e *alpha.InstanceFileSharesNfsExportOptionsSquashModeEnum) alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum_value["InstanceFileSharesNfsExportOptionsSquashModeEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
}

// InstanceNetworksModesEnumToProto converts a InstanceNetworksModesEnum enum to its proto representation.
func FilestoreAlphaInstanceNetworksModesEnumToProto(e *alpha.InstanceNetworksModesEnum) alphapb.FilestoreAlphaInstanceNetworksModesEnum {
	if e == nil {
		return alphapb.FilestoreAlphaInstanceNetworksModesEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaInstanceNetworksModesEnum_value["InstanceNetworksModesEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaInstanceNetworksModesEnum(v)
	}
	return alphapb.FilestoreAlphaInstanceNetworksModesEnum(0)
}

// InstanceFileSharesToProto converts a InstanceFileShares object to its proto representation.
func FilestoreAlphaInstanceFileSharesToProto(o *alpha.InstanceFileShares) *alphapb.FilestoreAlphaInstanceFileShares {
	if o == nil {
		return nil
	}
	p := &alphapb.FilestoreAlphaInstanceFileShares{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCapacityGb(dcl.ValueOrEmptyInt64(o.CapacityGb))
	p.SetSourceBackup(dcl.ValueOrEmptyString(o.SourceBackup))
	sNfsExportOptions := make([]*alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptions, len(o.NfsExportOptions))
	for i, r := range o.NfsExportOptions {
		sNfsExportOptions[i] = FilestoreAlphaInstanceFileSharesNfsExportOptionsToProto(&r)
	}
	p.SetNfsExportOptions(sNfsExportOptions)
	return p
}

// InstanceFileSharesNfsExportOptionsToProto converts a InstanceFileSharesNfsExportOptions object to its proto representation.
func FilestoreAlphaInstanceFileSharesNfsExportOptionsToProto(o *alpha.InstanceFileSharesNfsExportOptions) *alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptions {
	if o == nil {
		return nil
	}
	p := &alphapb.FilestoreAlphaInstanceFileSharesNfsExportOptions{}
	p.SetAccessMode(FilestoreAlphaInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(o.AccessMode))
	p.SetSquashMode(FilestoreAlphaInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(o.SquashMode))
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
func FilestoreAlphaInstanceNetworksToProto(o *alpha.InstanceNetworks) *alphapb.FilestoreAlphaInstanceNetworks {
	if o == nil {
		return nil
	}
	p := &alphapb.FilestoreAlphaInstanceNetworks{}
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetReservedIpRange(dcl.ValueOrEmptyString(o.ReservedIPRange))
	sModes := make([]alphapb.FilestoreAlphaInstanceNetworksModesEnum, len(o.Modes))
	for i, r := range o.Modes {
		sModes[i] = alphapb.FilestoreAlphaInstanceNetworksModesEnum(alphapb.FilestoreAlphaInstanceNetworksModesEnum_value[string(r)])
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
func InstanceToProto(resource *alpha.Instance) *alphapb.FilestoreAlphaInstance {
	p := &alphapb.FilestoreAlphaInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(FilestoreAlphaInstanceStateEnumToProto(resource.State))
	p.SetStatusMessage(dcl.ValueOrEmptyString(resource.StatusMessage))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetTier(FilestoreAlphaInstanceTierEnumToProto(resource.Tier))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sFileShares := make([]*alphapb.FilestoreAlphaInstanceFileShares, len(resource.FileShares))
	for i, r := range resource.FileShares {
		sFileShares[i] = FilestoreAlphaInstanceFileSharesToProto(&r)
	}
	p.SetFileShares(sFileShares)
	sNetworks := make([]*alphapb.FilestoreAlphaInstanceNetworks, len(resource.Networks))
	for i, r := range resource.Networks {
		sNetworks[i] = FilestoreAlphaInstanceNetworksToProto(&r)
	}
	p.SetNetworks(sNetworks)

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyFilestoreAlphaInstanceRequest) (*alphapb.FilestoreAlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyFilestoreAlphaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyFilestoreAlphaInstance(ctx context.Context, request *alphapb.ApplyFilestoreAlphaInstanceRequest) (*alphapb.FilestoreAlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteFilestoreAlphaInstance(ctx context.Context, request *alphapb.DeleteFilestoreAlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListFilestoreAlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListFilestoreAlphaInstance(ctx context.Context, request *alphapb.ListFilestoreAlphaInstanceRequest) (*alphapb.ListFilestoreAlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.FilestoreAlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListFilestoreAlphaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
