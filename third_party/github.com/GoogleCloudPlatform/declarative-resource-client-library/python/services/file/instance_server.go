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
	filepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/file/file_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToFileInstanceStateEnum(e filepb.FileInstanceStateEnum) *file.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filepb.FileInstanceStateEnum_name[int32(e)]; ok {
		e := file.InstanceStateEnum(n[len("FileInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTierEnum converts a InstanceTierEnum enum from its proto representation.
func ProtoToFileInstanceTierEnum(e filepb.FileInstanceTierEnum) *file.InstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filepb.FileInstanceTierEnum_name[int32(e)]; ok {
		e := file.InstanceTierEnum(n[len("FileInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsAccessModeEnum converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum from its proto representation.
func ProtoToFileInstanceFileSharesNfsExportOptionsAccessModeEnum(e filepb.FileInstanceFileSharesNfsExportOptionsAccessModeEnum) *file.InstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filepb.FileInstanceFileSharesNfsExportOptionsAccessModeEnum_name[int32(e)]; ok {
		e := file.InstanceFileSharesNfsExportOptionsAccessModeEnum(n[len("FileInstanceFileSharesNfsExportOptionsAccessModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileSharesNfsExportOptionsSquashModeEnum converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum from its proto representation.
func ProtoToFileInstanceFileSharesNfsExportOptionsSquashModeEnum(e filepb.FileInstanceFileSharesNfsExportOptionsSquashModeEnum) *file.InstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filepb.FileInstanceFileSharesNfsExportOptionsSquashModeEnum_name[int32(e)]; ok {
		e := file.InstanceFileSharesNfsExportOptionsSquashModeEnum(n[len("FileInstanceFileSharesNfsExportOptionsSquashModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworksModesEnum converts a InstanceNetworksModesEnum enum from its proto representation.
func ProtoToFileInstanceNetworksModesEnum(e filepb.FileInstanceNetworksModesEnum) *file.InstanceNetworksModesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := filepb.FileInstanceNetworksModesEnum_name[int32(e)]; ok {
		e := file.InstanceNetworksModesEnum(n[len("FileInstanceNetworksModesEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceFileShares converts a InstanceFileShares resource from its proto representation.
func ProtoToFileInstanceFileShares(p *filepb.FileInstanceFileShares) *file.InstanceFileShares {
	if p == nil {
		return nil
	}
	obj := &file.InstanceFileShares{
		Name:         dcl.StringOrNil(p.Name),
		CapacityGb:   dcl.Int64OrNil(p.CapacityGb),
		SourceBackup: dcl.StringOrNil(p.SourceBackup),
	}
	for _, r := range p.GetNfsExportOptions() {
		obj.NfsExportOptions = append(obj.NfsExportOptions, *ProtoToFileInstanceFileSharesNfsExportOptions(r))
	}
	return obj
}

// ProtoToInstanceFileSharesNfsExportOptions converts a InstanceFileSharesNfsExportOptions resource from its proto representation.
func ProtoToFileInstanceFileSharesNfsExportOptions(p *filepb.FileInstanceFileSharesNfsExportOptions) *file.InstanceFileSharesNfsExportOptions {
	if p == nil {
		return nil
	}
	obj := &file.InstanceFileSharesNfsExportOptions{
		AccessMode: ProtoToFileInstanceFileSharesNfsExportOptionsAccessModeEnum(p.GetAccessMode()),
		SquashMode: ProtoToFileInstanceFileSharesNfsExportOptionsSquashModeEnum(p.GetSquashMode()),
		AnonUid:    dcl.Int64OrNil(p.AnonUid),
		AnonGid:    dcl.Int64OrNil(p.AnonGid),
	}
	for _, r := range p.GetIpRanges() {
		obj.IPRanges = append(obj.IPRanges, r)
	}
	return obj
}

// ProtoToInstanceNetworks converts a InstanceNetworks resource from its proto representation.
func ProtoToFileInstanceNetworks(p *filepb.FileInstanceNetworks) *file.InstanceNetworks {
	if p == nil {
		return nil
	}
	obj := &file.InstanceNetworks{
		Network:         dcl.StringOrNil(p.Network),
		ReservedIPRange: dcl.StringOrNil(p.ReservedIpRange),
	}
	for _, r := range p.GetModes() {
		obj.Modes = append(obj.Modes, *ProtoToFileInstanceNetworksModesEnum(r))
	}
	for _, r := range p.GetIpAddresses() {
		obj.IPAddresses = append(obj.IPAddresses, r)
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *filepb.FileInstance) *file.Instance {
	obj := &file.Instance{
		Name:          dcl.StringOrNil(p.Name),
		Description:   dcl.StringOrNil(p.Description),
		State:         ProtoToFileInstanceStateEnum(p.GetState()),
		StatusMessage: dcl.StringOrNil(p.StatusMessage),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Tier:          ProtoToFileInstanceTierEnum(p.GetTier()),
		Etag:          dcl.StringOrNil(p.Etag),
		Project:       dcl.StringOrNil(p.Project),
		Location:      dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetFileShares() {
		obj.FileShares = append(obj.FileShares, *ProtoToFileInstanceFileShares(r))
	}
	for _, r := range p.GetNetworks() {
		obj.Networks = append(obj.Networks, *ProtoToFileInstanceNetworks(r))
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func FileInstanceStateEnumToProto(e *file.InstanceStateEnum) filepb.FileInstanceStateEnum {
	if e == nil {
		return filepb.FileInstanceStateEnum(0)
	}
	if v, ok := filepb.FileInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return filepb.FileInstanceStateEnum(v)
	}
	return filepb.FileInstanceStateEnum(0)
}

// InstanceTierEnumToProto converts a InstanceTierEnum enum to its proto representation.
func FileInstanceTierEnumToProto(e *file.InstanceTierEnum) filepb.FileInstanceTierEnum {
	if e == nil {
		return filepb.FileInstanceTierEnum(0)
	}
	if v, ok := filepb.FileInstanceTierEnum_value["InstanceTierEnum"+string(*e)]; ok {
		return filepb.FileInstanceTierEnum(v)
	}
	return filepb.FileInstanceTierEnum(0)
}

// InstanceFileSharesNfsExportOptionsAccessModeEnumToProto converts a InstanceFileSharesNfsExportOptionsAccessModeEnum enum to its proto representation.
func FileInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(e *file.InstanceFileSharesNfsExportOptionsAccessModeEnum) filepb.FileInstanceFileSharesNfsExportOptionsAccessModeEnum {
	if e == nil {
		return filepb.FileInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
	}
	if v, ok := filepb.FileInstanceFileSharesNfsExportOptionsAccessModeEnum_value["InstanceFileSharesNfsExportOptionsAccessModeEnum"+string(*e)]; ok {
		return filepb.FileInstanceFileSharesNfsExportOptionsAccessModeEnum(v)
	}
	return filepb.FileInstanceFileSharesNfsExportOptionsAccessModeEnum(0)
}

// InstanceFileSharesNfsExportOptionsSquashModeEnumToProto converts a InstanceFileSharesNfsExportOptionsSquashModeEnum enum to its proto representation.
func FileInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(e *file.InstanceFileSharesNfsExportOptionsSquashModeEnum) filepb.FileInstanceFileSharesNfsExportOptionsSquashModeEnum {
	if e == nil {
		return filepb.FileInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
	}
	if v, ok := filepb.FileInstanceFileSharesNfsExportOptionsSquashModeEnum_value["InstanceFileSharesNfsExportOptionsSquashModeEnum"+string(*e)]; ok {
		return filepb.FileInstanceFileSharesNfsExportOptionsSquashModeEnum(v)
	}
	return filepb.FileInstanceFileSharesNfsExportOptionsSquashModeEnum(0)
}

// InstanceNetworksModesEnumToProto converts a InstanceNetworksModesEnum enum to its proto representation.
func FileInstanceNetworksModesEnumToProto(e *file.InstanceNetworksModesEnum) filepb.FileInstanceNetworksModesEnum {
	if e == nil {
		return filepb.FileInstanceNetworksModesEnum(0)
	}
	if v, ok := filepb.FileInstanceNetworksModesEnum_value["InstanceNetworksModesEnum"+string(*e)]; ok {
		return filepb.FileInstanceNetworksModesEnum(v)
	}
	return filepb.FileInstanceNetworksModesEnum(0)
}

// InstanceFileSharesToProto converts a InstanceFileShares resource to its proto representation.
func FileInstanceFileSharesToProto(o *file.InstanceFileShares) *filepb.FileInstanceFileShares {
	if o == nil {
		return nil
	}
	p := &filepb.FileInstanceFileShares{
		Name:         dcl.ValueOrEmptyString(o.Name),
		CapacityGb:   dcl.ValueOrEmptyInt64(o.CapacityGb),
		SourceBackup: dcl.ValueOrEmptyString(o.SourceBackup),
	}
	for _, r := range o.NfsExportOptions {
		p.NfsExportOptions = append(p.NfsExportOptions, FileInstanceFileSharesNfsExportOptionsToProto(&r))
	}
	return p
}

// InstanceFileSharesNfsExportOptionsToProto converts a InstanceFileSharesNfsExportOptions resource to its proto representation.
func FileInstanceFileSharesNfsExportOptionsToProto(o *file.InstanceFileSharesNfsExportOptions) *filepb.FileInstanceFileSharesNfsExportOptions {
	if o == nil {
		return nil
	}
	p := &filepb.FileInstanceFileSharesNfsExportOptions{
		AccessMode: FileInstanceFileSharesNfsExportOptionsAccessModeEnumToProto(o.AccessMode),
		SquashMode: FileInstanceFileSharesNfsExportOptionsSquashModeEnumToProto(o.SquashMode),
		AnonUid:    dcl.ValueOrEmptyInt64(o.AnonUid),
		AnonGid:    dcl.ValueOrEmptyInt64(o.AnonGid),
	}
	for _, r := range o.IPRanges {
		p.IpRanges = append(p.IpRanges, r)
	}
	return p
}

// InstanceNetworksToProto converts a InstanceNetworks resource to its proto representation.
func FileInstanceNetworksToProto(o *file.InstanceNetworks) *filepb.FileInstanceNetworks {
	if o == nil {
		return nil
	}
	p := &filepb.FileInstanceNetworks{
		Network:         dcl.ValueOrEmptyString(o.Network),
		ReservedIpRange: dcl.ValueOrEmptyString(o.ReservedIPRange),
	}
	for _, r := range o.Modes {
		p.Modes = append(p.Modes, filepb.FileInstanceNetworksModesEnum(filepb.FileInstanceNetworksModesEnum_value[string(r)]))
	}
	for _, r := range o.IPAddresses {
		p.IpAddresses = append(p.IpAddresses, r)
	}
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *file.Instance) *filepb.FileInstance {
	p := &filepb.FileInstance{
		Name:          dcl.ValueOrEmptyString(resource.Name),
		Description:   dcl.ValueOrEmptyString(resource.Description),
		State:         FileInstanceStateEnumToProto(resource.State),
		StatusMessage: dcl.ValueOrEmptyString(resource.StatusMessage),
		CreateTime:    dcl.ValueOrEmptyString(resource.CreateTime),
		Tier:          FileInstanceTierEnumToProto(resource.Tier),
		Etag:          dcl.ValueOrEmptyString(resource.Etag),
		Project:       dcl.ValueOrEmptyString(resource.Project),
		Location:      dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.FileShares {
		p.FileShares = append(p.FileShares, FileInstanceFileSharesToProto(&r))
	}
	for _, r := range resource.Networks {
		p.Networks = append(p.Networks, FileInstanceNetworksToProto(&r))
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *file.Client, request *filepb.ApplyFileInstanceRequest) (*filepb.FileInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyFileInstance(ctx context.Context, request *filepb.ApplyFileInstanceRequest) (*filepb.FileInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteFileInstance(ctx context.Context, request *filepb.DeleteFileInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListFileInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListFileInstance(ctx context.Context, request *filepb.ListFileInstanceRequest) (*filepb.ListFileInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, ProtoToInstance(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*filepb.FileInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &filepb.ListFileInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*file.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return file.NewClient(conf), nil
}
