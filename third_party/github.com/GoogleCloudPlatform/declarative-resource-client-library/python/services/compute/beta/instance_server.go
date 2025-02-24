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

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceDisksInterfaceEnum converts a InstanceDisksInterfaceEnum enum from its proto representation.
func ProtoToComputeBetaInstanceDisksInterfaceEnum(e betapb.ComputeBetaInstanceDisksInterfaceEnum) *beta.InstanceDisksInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceDisksInterfaceEnum_name[int32(e)]; ok {
		e := beta.InstanceDisksInterfaceEnum(n[len("ComputeBetaInstanceDisksInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksModeEnum converts a InstanceDisksModeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceDisksModeEnum(e betapb.ComputeBetaInstanceDisksModeEnum) *beta.InstanceDisksModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceDisksModeEnum_name[int32(e)]; ok {
		e := beta.InstanceDisksModeEnum(n[len("ComputeBetaInstanceDisksModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksTypeEnum converts a InstanceDisksTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceDisksTypeEnum(e betapb.ComputeBetaInstanceDisksTypeEnum) *beta.InstanceDisksTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceDisksTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceDisksTypeEnum(n[len("ComputeBetaInstanceDisksTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesAccessConfigsNetworkTierEnum converts a InstanceNetworkInterfacesAccessConfigsNetworkTierEnum enum from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(e betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum) *beta.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum_name[int32(e)]; ok {
		e := beta.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum(n[len("ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesAccessConfigsTypeEnum converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(e betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum) *beta.InstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceNetworkInterfacesAccessConfigsTypeEnum(n[len("ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum converts a InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum enum from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(e betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum) *beta.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum_name[int32(e)]; ok {
		e := beta.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(n[len("ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum converts a InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(e betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum) *beta.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(n[len("ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStatusEnum converts a InstanceStatusEnum enum from its proto representation.
func ProtoToComputeBetaInstanceStatusEnum(e betapb.ComputeBetaInstanceStatusEnum) *beta.InstanceStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceStatusEnum_name[int32(e)]; ok {
		e := beta.InstanceStatusEnum(n[len("ComputeBetaInstanceStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisks converts a InstanceDisks object from its proto representation.
func ProtoToComputeBetaInstanceDisks(p *betapb.ComputeBetaInstanceDisks) *beta.InstanceDisks {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDisks{
		AutoDelete:        dcl.Bool(p.GetAutoDelete()),
		Boot:              dcl.Bool(p.GetBoot()),
		DeviceName:        dcl.StringOrNil(p.GetDeviceName()),
		DiskEncryptionKey: ProtoToComputeBetaInstanceDisksDiskEncryptionKey(p.GetDiskEncryptionKey()),
		Index:             dcl.Int64OrNil(p.GetIndex()),
		InitializeParams:  ProtoToComputeBetaInstanceDisksInitializeParams(p.GetInitializeParams()),
		Interface:         ProtoToComputeBetaInstanceDisksInterfaceEnum(p.GetInterface()),
		Mode:              ProtoToComputeBetaInstanceDisksModeEnum(p.GetMode()),
		Source:            dcl.StringOrNil(p.GetSource()),
		Type:              ProtoToComputeBetaInstanceDisksTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceDisksDiskEncryptionKey converts a InstanceDisksDiskEncryptionKey object from its proto representation.
func ProtoToComputeBetaInstanceDisksDiskEncryptionKey(p *betapb.ComputeBetaInstanceDisksDiskEncryptionKey) *beta.InstanceDisksDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDisksDiskEncryptionKey{
		RawKey:          dcl.StringOrNil(p.GetRawKey()),
		RsaEncryptedKey: dcl.StringOrNil(p.GetRsaEncryptedKey()),
		Sha256:          dcl.StringOrNil(p.GetSha256()),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParams converts a InstanceDisksInitializeParams object from its proto representation.
func ProtoToComputeBetaInstanceDisksInitializeParams(p *betapb.ComputeBetaInstanceDisksInitializeParams) *beta.InstanceDisksInitializeParams {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDisksInitializeParams{
		DiskName:                 dcl.StringOrNil(p.GetDiskName()),
		DiskSizeGb:               dcl.Int64OrNil(p.GetDiskSizeGb()),
		DiskType:                 dcl.StringOrNil(p.GetDiskType()),
		SourceImage:              dcl.StringOrNil(p.GetSourceImage()),
		SourceImageEncryptionKey: ProtoToComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParamsSourceImageEncryptionKey converts a InstanceDisksInitializeParamsSourceImageEncryptionKey object from its proto representation.
func ProtoToComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey(p *betapb.ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey) *beta.InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceDisksInitializeParamsSourceImageEncryptionKey{
		RawKey: dcl.StringOrNil(p.GetRawKey()),
		Sha256: dcl.StringOrNil(p.GetSha256()),
	}
	return obj
}

// ProtoToInstanceGuestAccelerators converts a InstanceGuestAccelerators object from its proto representation.
func ProtoToComputeBetaInstanceGuestAccelerators(p *betapb.ComputeBetaInstanceGuestAccelerators) *beta.InstanceGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGuestAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfaces converts a InstanceNetworkInterfaces object from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfaces(p *betapb.ComputeBetaInstanceNetworkInterfaces) *beta.InstanceNetworkInterfaces {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkInterfaces{
		Name:       dcl.StringOrNil(p.GetName()),
		Network:    dcl.StringOrNil(p.GetNetwork()),
		NetworkIP:  dcl.StringOrNil(p.GetNetworkIp()),
		Subnetwork: dcl.StringOrNil(p.GetSubnetwork()),
	}
	for _, r := range p.GetAccessConfigs() {
		obj.AccessConfigs = append(obj.AccessConfigs, *ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigs(r))
	}
	for _, r := range p.GetIpv6AccessConfigs() {
		obj.IPv6AccessConfigs = append(obj.IPv6AccessConfigs, *ProtoToComputeBetaInstanceNetworkInterfacesIPv6AccessConfigs(r))
	}
	for _, r := range p.GetAliasIpRanges() {
		obj.AliasIPRanges = append(obj.AliasIPRanges, *ProtoToComputeBetaInstanceNetworkInterfacesAliasIPRanges(r))
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAccessConfigs converts a InstanceNetworkInterfacesAccessConfigs object from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigs(p *betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigs) *beta.InstanceNetworkInterfacesAccessConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkInterfacesAccessConfigs{
		Name:                     dcl.StringOrNil(p.GetName()),
		NatIP:                    dcl.StringOrNil(p.GetNatIp()),
		ExternalIPv6:             dcl.StringOrNil(p.GetExternalIpv6()),
		ExternalIPv6PrefixLength: dcl.StringOrNil(p.GetExternalIpv6PrefixLength()),
		SetPublicPtr:             dcl.Bool(p.GetSetPublicPtr()),
		PublicPtrDomainName:      dcl.StringOrNil(p.GetPublicPtrDomainName()),
		NetworkTier:              ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(p.GetNetworkTier()),
		Type:                     ProtoToComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigs converts a InstanceNetworkInterfacesIPv6AccessConfigs object from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesIPv6AccessConfigs(p *betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigs) *beta.InstanceNetworkInterfacesIPv6AccessConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkInterfacesIPv6AccessConfigs{
		Name:                     dcl.StringOrNil(p.GetName()),
		NatIP:                    dcl.StringOrNil(p.GetNatIp()),
		ExternalIPv6:             dcl.StringOrNil(p.GetExternalIpv6()),
		ExternalIPv6PrefixLength: dcl.StringOrNil(p.GetExternalIpv6PrefixLength()),
		SetPublicPtr:             dcl.Bool(p.GetSetPublicPtr()),
		PublicPtrDomainName:      dcl.StringOrNil(p.GetPublicPtrDomainName()),
		NetworkTier:              ProtoToComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(p.GetNetworkTier()),
		Type:                     ProtoToComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAliasIPRanges converts a InstanceNetworkInterfacesAliasIPRanges object from its proto representation.
func ProtoToComputeBetaInstanceNetworkInterfacesAliasIPRanges(p *betapb.ComputeBetaInstanceNetworkInterfacesAliasIPRanges) *beta.InstanceNetworkInterfacesAliasIPRanges {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceNetworkInterfacesAliasIPRanges{
		IPCidrRange:         dcl.StringOrNil(p.GetIpCidrRange()),
		SubnetworkRangeName: dcl.StringOrNil(p.GetSubnetworkRangeName()),
	}
	return obj
}

// ProtoToInstanceScheduling converts a InstanceScheduling object from its proto representation.
func ProtoToComputeBetaInstanceScheduling(p *betapb.ComputeBetaInstanceScheduling) *beta.InstanceScheduling {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceScheduling{
		AutomaticRestart:  dcl.Bool(p.GetAutomaticRestart()),
		OnHostMaintenance: dcl.StringOrNil(p.GetOnHostMaintenance()),
		Preemptible:       dcl.Bool(p.GetPreemptible()),
	}
	return obj
}

// ProtoToInstanceServiceAccounts converts a InstanceServiceAccounts object from its proto representation.
func ProtoToComputeBetaInstanceServiceAccounts(p *betapb.ComputeBetaInstanceServiceAccounts) *beta.InstanceServiceAccounts {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceServiceAccounts{
		Email: dcl.StringOrNil(p.GetEmail()),
	}
	for _, r := range p.GetScopes() {
		obj.Scopes = append(obj.Scopes, r)
	}
	return obj
}

// ProtoToInstanceShieldedInstanceConfig converts a InstanceShieldedInstanceConfig object from its proto representation.
func ProtoToComputeBetaInstanceShieldedInstanceConfig(p *betapb.ComputeBetaInstanceShieldedInstanceConfig) *beta.InstanceShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.GetEnableSecureBoot()),
		EnableVtpm:                dcl.Bool(p.GetEnableVtpm()),
		EnableIntegrityMonitoring: dcl.Bool(p.GetEnableIntegrityMonitoring()),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.ComputeBetaInstance) *beta.Instance {
	obj := &beta.Instance{
		CanIPForward:           dcl.Bool(p.GetCanIpForward()),
		CpuPlatform:            dcl.StringOrNil(p.GetCpuPlatform()),
		CreationTimestamp:      dcl.StringOrNil(p.GetCreationTimestamp()),
		DeletionProtection:     dcl.Bool(p.GetDeletionProtection()),
		Description:            dcl.StringOrNil(p.GetDescription()),
		Hostname:               dcl.StringOrNil(p.GetHostname()),
		Id:                     dcl.StringOrNil(p.GetId()),
		MachineType:            dcl.StringOrNil(p.GetMachineType()),
		MinCpuPlatform:         dcl.StringOrNil(p.GetMinCpuPlatform()),
		Name:                   dcl.StringOrNil(p.GetName()),
		Scheduling:             ProtoToComputeBetaInstanceScheduling(p.GetScheduling()),
		ShieldedInstanceConfig: ProtoToComputeBetaInstanceShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		Status:                 ProtoToComputeBetaInstanceStatusEnum(p.GetStatus()),
		StatusMessage:          dcl.StringOrNil(p.GetStatusMessage()),
		Zone:                   dcl.StringOrNil(p.GetZone()),
		Project:                dcl.StringOrNil(p.GetProject()),
		SelfLink:               dcl.StringOrNil(p.GetSelfLink()),
	}
	for _, r := range p.GetDisks() {
		obj.Disks = append(obj.Disks, *ProtoToComputeBetaInstanceDisks(r))
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeBetaInstanceGuestAccelerators(r))
	}
	for _, r := range p.GetNetworkInterfaces() {
		obj.NetworkInterfaces = append(obj.NetworkInterfaces, *ProtoToComputeBetaInstanceNetworkInterfaces(r))
	}
	for _, r := range p.GetServiceAccounts() {
		obj.ServiceAccounts = append(obj.ServiceAccounts, *ProtoToComputeBetaInstanceServiceAccounts(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// InstanceDisksInterfaceEnumToProto converts a InstanceDisksInterfaceEnum enum to its proto representation.
func ComputeBetaInstanceDisksInterfaceEnumToProto(e *beta.InstanceDisksInterfaceEnum) betapb.ComputeBetaInstanceDisksInterfaceEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceDisksInterfaceEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceDisksInterfaceEnum_value["InstanceDisksInterfaceEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceDisksInterfaceEnum(v)
	}
	return betapb.ComputeBetaInstanceDisksInterfaceEnum(0)
}

// InstanceDisksModeEnumToProto converts a InstanceDisksModeEnum enum to its proto representation.
func ComputeBetaInstanceDisksModeEnumToProto(e *beta.InstanceDisksModeEnum) betapb.ComputeBetaInstanceDisksModeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceDisksModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceDisksModeEnum_value["InstanceDisksModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceDisksModeEnum(v)
	}
	return betapb.ComputeBetaInstanceDisksModeEnum(0)
}

// InstanceDisksTypeEnumToProto converts a InstanceDisksTypeEnum enum to its proto representation.
func ComputeBetaInstanceDisksTypeEnumToProto(e *beta.InstanceDisksTypeEnum) betapb.ComputeBetaInstanceDisksTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceDisksTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceDisksTypeEnum_value["InstanceDisksTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceDisksTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceDisksTypeEnum(0)
}

// InstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto converts a InstanceNetworkInterfacesAccessConfigsNetworkTierEnum enum to its proto representation.
func ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto(e *beta.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum) betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum_value["InstanceNetworkInterfacesAccessConfigsNetworkTierEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(v)
	}
	return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(0)
}

// InstanceNetworkInterfacesAccessConfigsTypeEnumToProto converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum to its proto representation.
func ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(e *beta.InstanceNetworkInterfacesAccessConfigsTypeEnum) betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum_value["InstanceNetworkInterfacesAccessConfigsTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
}

// InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto converts a InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum enum to its proto representation.
func ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto(e *beta.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum) betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum_value["InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(v)
	}
	return betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(0)
}

// InstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto converts a InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum enum to its proto representation.
func ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto(e *beta.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum) betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum_value["InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(0)
}

// InstanceStatusEnumToProto converts a InstanceStatusEnum enum to its proto representation.
func ComputeBetaInstanceStatusEnumToProto(e *beta.InstanceStatusEnum) betapb.ComputeBetaInstanceStatusEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceStatusEnum_value["InstanceStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceStatusEnum(v)
	}
	return betapb.ComputeBetaInstanceStatusEnum(0)
}

// InstanceDisksToProto converts a InstanceDisks object to its proto representation.
func ComputeBetaInstanceDisksToProto(o *beta.InstanceDisks) *betapb.ComputeBetaInstanceDisks {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceDisks{}
	p.SetAutoDelete(dcl.ValueOrEmptyBool(o.AutoDelete))
	p.SetBoot(dcl.ValueOrEmptyBool(o.Boot))
	p.SetDeviceName(dcl.ValueOrEmptyString(o.DeviceName))
	p.SetDiskEncryptionKey(ComputeBetaInstanceDisksDiskEncryptionKeyToProto(o.DiskEncryptionKey))
	p.SetIndex(dcl.ValueOrEmptyInt64(o.Index))
	p.SetInitializeParams(ComputeBetaInstanceDisksInitializeParamsToProto(o.InitializeParams))
	p.SetInterface(ComputeBetaInstanceDisksInterfaceEnumToProto(o.Interface))
	p.SetMode(ComputeBetaInstanceDisksModeEnumToProto(o.Mode))
	p.SetSource(dcl.ValueOrEmptyString(o.Source))
	p.SetType(ComputeBetaInstanceDisksTypeEnumToProto(o.Type))
	return p
}

// InstanceDisksDiskEncryptionKeyToProto converts a InstanceDisksDiskEncryptionKey object to its proto representation.
func ComputeBetaInstanceDisksDiskEncryptionKeyToProto(o *beta.InstanceDisksDiskEncryptionKey) *betapb.ComputeBetaInstanceDisksDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceDisksDiskEncryptionKey{}
	p.SetRawKey(dcl.ValueOrEmptyString(o.RawKey))
	p.SetRsaEncryptedKey(dcl.ValueOrEmptyString(o.RsaEncryptedKey))
	p.SetSha256(dcl.ValueOrEmptyString(o.Sha256))
	return p
}

// InstanceDisksInitializeParamsToProto converts a InstanceDisksInitializeParams object to its proto representation.
func ComputeBetaInstanceDisksInitializeParamsToProto(o *beta.InstanceDisksInitializeParams) *betapb.ComputeBetaInstanceDisksInitializeParams {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceDisksInitializeParams{}
	p.SetDiskName(dcl.ValueOrEmptyString(o.DiskName))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	p.SetDiskType(dcl.ValueOrEmptyString(o.DiskType))
	p.SetSourceImage(dcl.ValueOrEmptyString(o.SourceImage))
	p.SetSourceImageEncryptionKey(ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o.SourceImageEncryptionKey))
	return p
}

// InstanceDisksInitializeParamsSourceImageEncryptionKeyToProto converts a InstanceDisksInitializeParamsSourceImageEncryptionKey object to its proto representation.
func ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o *beta.InstanceDisksInitializeParamsSourceImageEncryptionKey) *betapb.ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceDisksInitializeParamsSourceImageEncryptionKey{}
	p.SetRawKey(dcl.ValueOrEmptyString(o.RawKey))
	p.SetSha256(dcl.ValueOrEmptyString(o.Sha256))
	return p
}

// InstanceGuestAcceleratorsToProto converts a InstanceGuestAccelerators object to its proto representation.
func ComputeBetaInstanceGuestAcceleratorsToProto(o *beta.InstanceGuestAccelerators) *betapb.ComputeBetaInstanceGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGuestAccelerators{}
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	return p
}

// InstanceNetworkInterfacesToProto converts a InstanceNetworkInterfaces object to its proto representation.
func ComputeBetaInstanceNetworkInterfacesToProto(o *beta.InstanceNetworkInterfaces) *betapb.ComputeBetaInstanceNetworkInterfaces {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceNetworkInterfaces{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetNetworkIp(dcl.ValueOrEmptyString(o.NetworkIP))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	sAccessConfigs := make([]*betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigs, len(o.AccessConfigs))
	for i, r := range o.AccessConfigs {
		sAccessConfigs[i] = ComputeBetaInstanceNetworkInterfacesAccessConfigsToProto(&r)
	}
	p.SetAccessConfigs(sAccessConfigs)
	sIPv6AccessConfigs := make([]*betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigs, len(o.IPv6AccessConfigs))
	for i, r := range o.IPv6AccessConfigs {
		sIPv6AccessConfigs[i] = ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsToProto(&r)
	}
	p.SetIpv6AccessConfigs(sIPv6AccessConfigs)
	sAliasIPRanges := make([]*betapb.ComputeBetaInstanceNetworkInterfacesAliasIPRanges, len(o.AliasIPRanges))
	for i, r := range o.AliasIPRanges {
		sAliasIPRanges[i] = ComputeBetaInstanceNetworkInterfacesAliasIPRangesToProto(&r)
	}
	p.SetAliasIpRanges(sAliasIPRanges)
	return p
}

// InstanceNetworkInterfacesAccessConfigsToProto converts a InstanceNetworkInterfacesAccessConfigs object to its proto representation.
func ComputeBetaInstanceNetworkInterfacesAccessConfigsToProto(o *beta.InstanceNetworkInterfacesAccessConfigs) *betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceNetworkInterfacesAccessConfigs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNatIp(dcl.ValueOrEmptyString(o.NatIP))
	p.SetExternalIpv6(dcl.ValueOrEmptyString(o.ExternalIPv6))
	p.SetExternalIpv6PrefixLength(dcl.ValueOrEmptyString(o.ExternalIPv6PrefixLength))
	p.SetSetPublicPtr(dcl.ValueOrEmptyBool(o.SetPublicPtr))
	p.SetPublicPtrDomainName(dcl.ValueOrEmptyString(o.PublicPtrDomainName))
	p.SetNetworkTier(ComputeBetaInstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto(o.NetworkTier))
	p.SetType(ComputeBetaInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(o.Type))
	return p
}

// InstanceNetworkInterfacesIPv6AccessConfigsToProto converts a InstanceNetworkInterfacesIPv6AccessConfigs object to its proto representation.
func ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsToProto(o *beta.InstanceNetworkInterfacesIPv6AccessConfigs) *betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNatIp(dcl.ValueOrEmptyString(o.NatIP))
	p.SetExternalIpv6(dcl.ValueOrEmptyString(o.ExternalIPv6))
	p.SetExternalIpv6PrefixLength(dcl.ValueOrEmptyString(o.ExternalIPv6PrefixLength))
	p.SetSetPublicPtr(dcl.ValueOrEmptyBool(o.SetPublicPtr))
	p.SetPublicPtrDomainName(dcl.ValueOrEmptyString(o.PublicPtrDomainName))
	p.SetNetworkTier(ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto(o.NetworkTier))
	p.SetType(ComputeBetaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto(o.Type))
	return p
}

// InstanceNetworkInterfacesAliasIPRangesToProto converts a InstanceNetworkInterfacesAliasIPRanges object to its proto representation.
func ComputeBetaInstanceNetworkInterfacesAliasIPRangesToProto(o *beta.InstanceNetworkInterfacesAliasIPRanges) *betapb.ComputeBetaInstanceNetworkInterfacesAliasIPRanges {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceNetworkInterfacesAliasIPRanges{}
	p.SetIpCidrRange(dcl.ValueOrEmptyString(o.IPCidrRange))
	p.SetSubnetworkRangeName(dcl.ValueOrEmptyString(o.SubnetworkRangeName))
	return p
}

// InstanceSchedulingToProto converts a InstanceScheduling object to its proto representation.
func ComputeBetaInstanceSchedulingToProto(o *beta.InstanceScheduling) *betapb.ComputeBetaInstanceScheduling {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceScheduling{}
	p.SetAutomaticRestart(dcl.ValueOrEmptyBool(o.AutomaticRestart))
	p.SetOnHostMaintenance(dcl.ValueOrEmptyString(o.OnHostMaintenance))
	p.SetPreemptible(dcl.ValueOrEmptyBool(o.Preemptible))
	return p
}

// InstanceServiceAccountsToProto converts a InstanceServiceAccounts object to its proto representation.
func ComputeBetaInstanceServiceAccountsToProto(o *beta.InstanceServiceAccounts) *betapb.ComputeBetaInstanceServiceAccounts {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceServiceAccounts{}
	p.SetEmail(dcl.ValueOrEmptyString(o.Email))
	sScopes := make([]string, len(o.Scopes))
	for i, r := range o.Scopes {
		sScopes[i] = r
	}
	p.SetScopes(sScopes)
	return p
}

// InstanceShieldedInstanceConfigToProto converts a InstanceShieldedInstanceConfig object to its proto representation.
func ComputeBetaInstanceShieldedInstanceConfigToProto(o *beta.InstanceShieldedInstanceConfig) *betapb.ComputeBetaInstanceShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceShieldedInstanceConfig{}
	p.SetEnableSecureBoot(dcl.ValueOrEmptyBool(o.EnableSecureBoot))
	p.SetEnableVtpm(dcl.ValueOrEmptyBool(o.EnableVtpm))
	p.SetEnableIntegrityMonitoring(dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring))
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *beta.Instance) *betapb.ComputeBetaInstance {
	p := &betapb.ComputeBetaInstance{}
	p.SetCanIpForward(dcl.ValueOrEmptyBool(resource.CanIPForward))
	p.SetCpuPlatform(dcl.ValueOrEmptyString(resource.CpuPlatform))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetDeletionProtection(dcl.ValueOrEmptyBool(resource.DeletionProtection))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetHostname(dcl.ValueOrEmptyString(resource.Hostname))
	p.SetId(dcl.ValueOrEmptyString(resource.Id))
	p.SetMachineType(dcl.ValueOrEmptyString(resource.MachineType))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(resource.MinCpuPlatform))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetScheduling(ComputeBetaInstanceSchedulingToProto(resource.Scheduling))
	p.SetShieldedInstanceConfig(ComputeBetaInstanceShieldedInstanceConfigToProto(resource.ShieldedInstanceConfig))
	p.SetStatus(ComputeBetaInstanceStatusEnumToProto(resource.Status))
	p.SetStatusMessage(dcl.ValueOrEmptyString(resource.StatusMessage))
	p.SetZone(dcl.ValueOrEmptyString(resource.Zone))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	sDisks := make([]*betapb.ComputeBetaInstanceDisks, len(resource.Disks))
	for i, r := range resource.Disks {
		sDisks[i] = ComputeBetaInstanceDisksToProto(&r)
	}
	p.SetDisks(sDisks)
	sGuestAccelerators := make([]*betapb.ComputeBetaInstanceGuestAccelerators, len(resource.GuestAccelerators))
	for i, r := range resource.GuestAccelerators {
		sGuestAccelerators[i] = ComputeBetaInstanceGuestAcceleratorsToProto(&r)
	}
	p.SetGuestAccelerators(sGuestAccelerators)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	mMetadata := make(map[string]string, len(resource.Metadata))
	for k, r := range resource.Metadata {
		mMetadata[k] = r
	}
	p.SetMetadata(mMetadata)
	sNetworkInterfaces := make([]*betapb.ComputeBetaInstanceNetworkInterfaces, len(resource.NetworkInterfaces))
	for i, r := range resource.NetworkInterfaces {
		sNetworkInterfaces[i] = ComputeBetaInstanceNetworkInterfacesToProto(&r)
	}
	p.SetNetworkInterfaces(sNetworkInterfaces)
	sServiceAccounts := make([]*betapb.ComputeBetaInstanceServiceAccounts, len(resource.ServiceAccounts))
	for i, r := range resource.ServiceAccounts {
		sServiceAccounts[i] = ComputeBetaInstanceServiceAccountsToProto(&r)
	}
	p.SetServiceAccounts(sServiceAccounts)
	sTags := make([]string, len(resource.Tags))
	for i, r := range resource.Tags {
		sTags[i] = r
	}
	p.SetTags(sTags)

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaInstanceRequest) (*betapb.ComputeBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyComputeBetaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyComputeBetaInstance(ctx context.Context, request *betapb.ApplyComputeBetaInstanceRequest) (*betapb.ComputeBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteComputeBetaInstance(ctx context.Context, request *betapb.DeleteComputeBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListComputeBetaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListComputeBetaInstance(ctx context.Context, request *betapb.ListComputeBetaInstanceRequest) (*betapb.ListComputeBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetZone())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
