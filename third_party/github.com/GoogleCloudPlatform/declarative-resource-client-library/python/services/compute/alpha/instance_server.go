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

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceDisksInterfaceEnum converts a InstanceDisksInterfaceEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceDisksInterfaceEnum(e alphapb.ComputeAlphaInstanceDisksInterfaceEnum) *alpha.InstanceDisksInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceDisksInterfaceEnum_name[int32(e)]; ok {
		e := alpha.InstanceDisksInterfaceEnum(n[len("ComputeAlphaInstanceDisksInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksModeEnum converts a InstanceDisksModeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceDisksModeEnum(e alphapb.ComputeAlphaInstanceDisksModeEnum) *alpha.InstanceDisksModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceDisksModeEnum_name[int32(e)]; ok {
		e := alpha.InstanceDisksModeEnum(n[len("ComputeAlphaInstanceDisksModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksTypeEnum converts a InstanceDisksTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceDisksTypeEnum(e alphapb.ComputeAlphaInstanceDisksTypeEnum) *alpha.InstanceDisksTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceDisksTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceDisksTypeEnum(n[len("ComputeAlphaInstanceDisksTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesAccessConfigsNetworkTierEnum converts a InstanceNetworkInterfacesAccessConfigsNetworkTierEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(e alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum) *alpha.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum_name[int32(e)]; ok {
		e := alpha.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum(n[len("ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesAccessConfigsTypeEnum converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum(e alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum) *alpha.InstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceNetworkInterfacesAccessConfigsTypeEnum(n[len("ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum converts a InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(e alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum) *alpha.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum_name[int32(e)]; ok {
		e := alpha.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(n[len("ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum converts a InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(e alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum) *alpha.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(n[len("ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStatusEnum converts a InstanceStatusEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceStatusEnum(e alphapb.ComputeAlphaInstanceStatusEnum) *alpha.InstanceStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceStatusEnum_name[int32(e)]; ok {
		e := alpha.InstanceStatusEnum(n[len("ComputeAlphaInstanceStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisks converts a InstanceDisks object from its proto representation.
func ProtoToComputeAlphaInstanceDisks(p *alphapb.ComputeAlphaInstanceDisks) *alpha.InstanceDisks {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDisks{
		AutoDelete:        dcl.Bool(p.GetAutoDelete()),
		Boot:              dcl.Bool(p.GetBoot()),
		DeviceName:        dcl.StringOrNil(p.GetDeviceName()),
		DiskEncryptionKey: ProtoToComputeAlphaInstanceDisksDiskEncryptionKey(p.GetDiskEncryptionKey()),
		Index:             dcl.Int64OrNil(p.GetIndex()),
		InitializeParams:  ProtoToComputeAlphaInstanceDisksInitializeParams(p.GetInitializeParams()),
		Interface:         ProtoToComputeAlphaInstanceDisksInterfaceEnum(p.GetInterface()),
		Mode:              ProtoToComputeAlphaInstanceDisksModeEnum(p.GetMode()),
		Source:            dcl.StringOrNil(p.GetSource()),
		Type:              ProtoToComputeAlphaInstanceDisksTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceDisksDiskEncryptionKey converts a InstanceDisksDiskEncryptionKey object from its proto representation.
func ProtoToComputeAlphaInstanceDisksDiskEncryptionKey(p *alphapb.ComputeAlphaInstanceDisksDiskEncryptionKey) *alpha.InstanceDisksDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDisksDiskEncryptionKey{
		RawKey:          dcl.StringOrNil(p.GetRawKey()),
		RsaEncryptedKey: dcl.StringOrNil(p.GetRsaEncryptedKey()),
		Sha256:          dcl.StringOrNil(p.GetSha256()),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParams converts a InstanceDisksInitializeParams object from its proto representation.
func ProtoToComputeAlphaInstanceDisksInitializeParams(p *alphapb.ComputeAlphaInstanceDisksInitializeParams) *alpha.InstanceDisksInitializeParams {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDisksInitializeParams{
		DiskName:                 dcl.StringOrNil(p.GetDiskName()),
		DiskSizeGb:               dcl.Int64OrNil(p.GetDiskSizeGb()),
		DiskType:                 dcl.StringOrNil(p.GetDiskType()),
		SourceImage:              dcl.StringOrNil(p.GetSourceImage()),
		SourceImageEncryptionKey: ProtoToComputeAlphaInstanceDisksInitializeParamsSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParamsSourceImageEncryptionKey converts a InstanceDisksInitializeParamsSourceImageEncryptionKey object from its proto representation.
func ProtoToComputeAlphaInstanceDisksInitializeParamsSourceImageEncryptionKey(p *alphapb.ComputeAlphaInstanceDisksInitializeParamsSourceImageEncryptionKey) *alpha.InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceDisksInitializeParamsSourceImageEncryptionKey{
		RawKey: dcl.StringOrNil(p.GetRawKey()),
		Sha256: dcl.StringOrNil(p.GetSha256()),
	}
	return obj
}

// ProtoToInstanceGuestAccelerators converts a InstanceGuestAccelerators object from its proto representation.
func ProtoToComputeAlphaInstanceGuestAccelerators(p *alphapb.ComputeAlphaInstanceGuestAccelerators) *alpha.InstanceGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGuestAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfaces converts a InstanceNetworkInterfaces object from its proto representation.
func ProtoToComputeAlphaInstanceNetworkInterfaces(p *alphapb.ComputeAlphaInstanceNetworkInterfaces) *alpha.InstanceNetworkInterfaces {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNetworkInterfaces{
		Name:       dcl.StringOrNil(p.GetName()),
		Network:    dcl.StringOrNil(p.GetNetwork()),
		NetworkIP:  dcl.StringOrNil(p.GetNetworkIp()),
		Subnetwork: dcl.StringOrNil(p.GetSubnetwork()),
	}
	for _, r := range p.GetAccessConfigs() {
		obj.AccessConfigs = append(obj.AccessConfigs, *ProtoToComputeAlphaInstanceNetworkInterfacesAccessConfigs(r))
	}
	for _, r := range p.GetIpv6AccessConfigs() {
		obj.IPv6AccessConfigs = append(obj.IPv6AccessConfigs, *ProtoToComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigs(r))
	}
	for _, r := range p.GetAliasIpRanges() {
		obj.AliasIPRanges = append(obj.AliasIPRanges, *ProtoToComputeAlphaInstanceNetworkInterfacesAliasIPRanges(r))
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAccessConfigs converts a InstanceNetworkInterfacesAccessConfigs object from its proto representation.
func ProtoToComputeAlphaInstanceNetworkInterfacesAccessConfigs(p *alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigs) *alpha.InstanceNetworkInterfacesAccessConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNetworkInterfacesAccessConfigs{
		Name:                     dcl.StringOrNil(p.GetName()),
		NatIP:                    dcl.StringOrNil(p.GetNatIp()),
		ExternalIPv6:             dcl.StringOrNil(p.GetExternalIpv6()),
		ExternalIPv6PrefixLength: dcl.StringOrNil(p.GetExternalIpv6PrefixLength()),
		SetPublicPtr:             dcl.Bool(p.GetSetPublicPtr()),
		PublicPtrDomainName:      dcl.StringOrNil(p.GetPublicPtrDomainName()),
		NetworkTier:              ProtoToComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(p.GetNetworkTier()),
		Type:                     ProtoToComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigs converts a InstanceNetworkInterfacesIPv6AccessConfigs object from its proto representation.
func ProtoToComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigs(p *alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigs) *alpha.InstanceNetworkInterfacesIPv6AccessConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNetworkInterfacesIPv6AccessConfigs{
		Name:                     dcl.StringOrNil(p.GetName()),
		NatIP:                    dcl.StringOrNil(p.GetNatIp()),
		ExternalIPv6:             dcl.StringOrNil(p.GetExternalIpv6()),
		ExternalIPv6PrefixLength: dcl.StringOrNil(p.GetExternalIpv6PrefixLength()),
		SetPublicPtr:             dcl.Bool(p.GetSetPublicPtr()),
		PublicPtrDomainName:      dcl.StringOrNil(p.GetPublicPtrDomainName()),
		NetworkTier:              ProtoToComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(p.GetNetworkTier()),
		Type:                     ProtoToComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAliasIPRanges converts a InstanceNetworkInterfacesAliasIPRanges object from its proto representation.
func ProtoToComputeAlphaInstanceNetworkInterfacesAliasIPRanges(p *alphapb.ComputeAlphaInstanceNetworkInterfacesAliasIPRanges) *alpha.InstanceNetworkInterfacesAliasIPRanges {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceNetworkInterfacesAliasIPRanges{
		IPCidrRange:         dcl.StringOrNil(p.GetIpCidrRange()),
		SubnetworkRangeName: dcl.StringOrNil(p.GetSubnetworkRangeName()),
	}
	return obj
}

// ProtoToInstanceScheduling converts a InstanceScheduling object from its proto representation.
func ProtoToComputeAlphaInstanceScheduling(p *alphapb.ComputeAlphaInstanceScheduling) *alpha.InstanceScheduling {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceScheduling{
		AutomaticRestart:  dcl.Bool(p.GetAutomaticRestart()),
		OnHostMaintenance: dcl.StringOrNil(p.GetOnHostMaintenance()),
		Preemptible:       dcl.Bool(p.GetPreemptible()),
	}
	return obj
}

// ProtoToInstanceServiceAccounts converts a InstanceServiceAccounts object from its proto representation.
func ProtoToComputeAlphaInstanceServiceAccounts(p *alphapb.ComputeAlphaInstanceServiceAccounts) *alpha.InstanceServiceAccounts {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceServiceAccounts{
		Email: dcl.StringOrNil(p.GetEmail()),
	}
	for _, r := range p.GetScopes() {
		obj.Scopes = append(obj.Scopes, r)
	}
	return obj
}

// ProtoToInstanceShieldedInstanceConfig converts a InstanceShieldedInstanceConfig object from its proto representation.
func ProtoToComputeAlphaInstanceShieldedInstanceConfig(p *alphapb.ComputeAlphaInstanceShieldedInstanceConfig) *alpha.InstanceShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.GetEnableSecureBoot()),
		EnableVtpm:                dcl.Bool(p.GetEnableVtpm()),
		EnableIntegrityMonitoring: dcl.Bool(p.GetEnableIntegrityMonitoring()),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.ComputeAlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
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
		Scheduling:             ProtoToComputeAlphaInstanceScheduling(p.GetScheduling()),
		ShieldedInstanceConfig: ProtoToComputeAlphaInstanceShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		Status:                 ProtoToComputeAlphaInstanceStatusEnum(p.GetStatus()),
		StatusMessage:          dcl.StringOrNil(p.GetStatusMessage()),
		Zone:                   dcl.StringOrNil(p.GetZone()),
		Project:                dcl.StringOrNil(p.GetProject()),
		SelfLink:               dcl.StringOrNil(p.GetSelfLink()),
	}
	for _, r := range p.GetDisks() {
		obj.Disks = append(obj.Disks, *ProtoToComputeAlphaInstanceDisks(r))
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeAlphaInstanceGuestAccelerators(r))
	}
	for _, r := range p.GetNetworkInterfaces() {
		obj.NetworkInterfaces = append(obj.NetworkInterfaces, *ProtoToComputeAlphaInstanceNetworkInterfaces(r))
	}
	for _, r := range p.GetServiceAccounts() {
		obj.ServiceAccounts = append(obj.ServiceAccounts, *ProtoToComputeAlphaInstanceServiceAccounts(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// InstanceDisksInterfaceEnumToProto converts a InstanceDisksInterfaceEnum enum to its proto representation.
func ComputeAlphaInstanceDisksInterfaceEnumToProto(e *alpha.InstanceDisksInterfaceEnum) alphapb.ComputeAlphaInstanceDisksInterfaceEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceDisksInterfaceEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceDisksInterfaceEnum_value["InstanceDisksInterfaceEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceDisksInterfaceEnum(v)
	}
	return alphapb.ComputeAlphaInstanceDisksInterfaceEnum(0)
}

// InstanceDisksModeEnumToProto converts a InstanceDisksModeEnum enum to its proto representation.
func ComputeAlphaInstanceDisksModeEnumToProto(e *alpha.InstanceDisksModeEnum) alphapb.ComputeAlphaInstanceDisksModeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceDisksModeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceDisksModeEnum_value["InstanceDisksModeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceDisksModeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceDisksModeEnum(0)
}

// InstanceDisksTypeEnumToProto converts a InstanceDisksTypeEnum enum to its proto representation.
func ComputeAlphaInstanceDisksTypeEnumToProto(e *alpha.InstanceDisksTypeEnum) alphapb.ComputeAlphaInstanceDisksTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceDisksTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceDisksTypeEnum_value["InstanceDisksTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceDisksTypeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceDisksTypeEnum(0)
}

// InstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto converts a InstanceNetworkInterfacesAccessConfigsNetworkTierEnum enum to its proto representation.
func ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto(e *alpha.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum) alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum_value["InstanceNetworkInterfacesAccessConfigsNetworkTierEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(v)
	}
	return alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(0)
}

// InstanceNetworkInterfacesAccessConfigsTypeEnumToProto converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum to its proto representation.
func ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(e *alpha.InstanceNetworkInterfacesAccessConfigsTypeEnum) alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum_value["InstanceNetworkInterfacesAccessConfigsTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
}

// InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto converts a InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum enum to its proto representation.
func ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto(e *alpha.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum) alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum_value["InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(v)
	}
	return alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(0)
}

// InstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto converts a InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum enum to its proto representation.
func ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto(e *alpha.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum) alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum_value["InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(0)
}

// InstanceStatusEnumToProto converts a InstanceStatusEnum enum to its proto representation.
func ComputeAlphaInstanceStatusEnumToProto(e *alpha.InstanceStatusEnum) alphapb.ComputeAlphaInstanceStatusEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceStatusEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceStatusEnum_value["InstanceStatusEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceStatusEnum(v)
	}
	return alphapb.ComputeAlphaInstanceStatusEnum(0)
}

// InstanceDisksToProto converts a InstanceDisks object to its proto representation.
func ComputeAlphaInstanceDisksToProto(o *alpha.InstanceDisks) *alphapb.ComputeAlphaInstanceDisks {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceDisks{}
	p.SetAutoDelete(dcl.ValueOrEmptyBool(o.AutoDelete))
	p.SetBoot(dcl.ValueOrEmptyBool(o.Boot))
	p.SetDeviceName(dcl.ValueOrEmptyString(o.DeviceName))
	p.SetDiskEncryptionKey(ComputeAlphaInstanceDisksDiskEncryptionKeyToProto(o.DiskEncryptionKey))
	p.SetIndex(dcl.ValueOrEmptyInt64(o.Index))
	p.SetInitializeParams(ComputeAlphaInstanceDisksInitializeParamsToProto(o.InitializeParams))
	p.SetInterface(ComputeAlphaInstanceDisksInterfaceEnumToProto(o.Interface))
	p.SetMode(ComputeAlphaInstanceDisksModeEnumToProto(o.Mode))
	p.SetSource(dcl.ValueOrEmptyString(o.Source))
	p.SetType(ComputeAlphaInstanceDisksTypeEnumToProto(o.Type))
	return p
}

// InstanceDisksDiskEncryptionKeyToProto converts a InstanceDisksDiskEncryptionKey object to its proto representation.
func ComputeAlphaInstanceDisksDiskEncryptionKeyToProto(o *alpha.InstanceDisksDiskEncryptionKey) *alphapb.ComputeAlphaInstanceDisksDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceDisksDiskEncryptionKey{}
	p.SetRawKey(dcl.ValueOrEmptyString(o.RawKey))
	p.SetRsaEncryptedKey(dcl.ValueOrEmptyString(o.RsaEncryptedKey))
	p.SetSha256(dcl.ValueOrEmptyString(o.Sha256))
	return p
}

// InstanceDisksInitializeParamsToProto converts a InstanceDisksInitializeParams object to its proto representation.
func ComputeAlphaInstanceDisksInitializeParamsToProto(o *alpha.InstanceDisksInitializeParams) *alphapb.ComputeAlphaInstanceDisksInitializeParams {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceDisksInitializeParams{}
	p.SetDiskName(dcl.ValueOrEmptyString(o.DiskName))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	p.SetDiskType(dcl.ValueOrEmptyString(o.DiskType))
	p.SetSourceImage(dcl.ValueOrEmptyString(o.SourceImage))
	p.SetSourceImageEncryptionKey(ComputeAlphaInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o.SourceImageEncryptionKey))
	return p
}

// InstanceDisksInitializeParamsSourceImageEncryptionKeyToProto converts a InstanceDisksInitializeParamsSourceImageEncryptionKey object to its proto representation.
func ComputeAlphaInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o *alpha.InstanceDisksInitializeParamsSourceImageEncryptionKey) *alphapb.ComputeAlphaInstanceDisksInitializeParamsSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceDisksInitializeParamsSourceImageEncryptionKey{}
	p.SetRawKey(dcl.ValueOrEmptyString(o.RawKey))
	p.SetSha256(dcl.ValueOrEmptyString(o.Sha256))
	return p
}

// InstanceGuestAcceleratorsToProto converts a InstanceGuestAccelerators object to its proto representation.
func ComputeAlphaInstanceGuestAcceleratorsToProto(o *alpha.InstanceGuestAccelerators) *alphapb.ComputeAlphaInstanceGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGuestAccelerators{}
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	return p
}

// InstanceNetworkInterfacesToProto converts a InstanceNetworkInterfaces object to its proto representation.
func ComputeAlphaInstanceNetworkInterfacesToProto(o *alpha.InstanceNetworkInterfaces) *alphapb.ComputeAlphaInstanceNetworkInterfaces {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceNetworkInterfaces{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetNetworkIp(dcl.ValueOrEmptyString(o.NetworkIP))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	sAccessConfigs := make([]*alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigs, len(o.AccessConfigs))
	for i, r := range o.AccessConfigs {
		sAccessConfigs[i] = ComputeAlphaInstanceNetworkInterfacesAccessConfigsToProto(&r)
	}
	p.SetAccessConfigs(sAccessConfigs)
	sIPv6AccessConfigs := make([]*alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigs, len(o.IPv6AccessConfigs))
	for i, r := range o.IPv6AccessConfigs {
		sIPv6AccessConfigs[i] = ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsToProto(&r)
	}
	p.SetIpv6AccessConfigs(sIPv6AccessConfigs)
	sAliasIPRanges := make([]*alphapb.ComputeAlphaInstanceNetworkInterfacesAliasIPRanges, len(o.AliasIPRanges))
	for i, r := range o.AliasIPRanges {
		sAliasIPRanges[i] = ComputeAlphaInstanceNetworkInterfacesAliasIPRangesToProto(&r)
	}
	p.SetAliasIpRanges(sAliasIPRanges)
	return p
}

// InstanceNetworkInterfacesAccessConfigsToProto converts a InstanceNetworkInterfacesAccessConfigs object to its proto representation.
func ComputeAlphaInstanceNetworkInterfacesAccessConfigsToProto(o *alpha.InstanceNetworkInterfacesAccessConfigs) *alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceNetworkInterfacesAccessConfigs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNatIp(dcl.ValueOrEmptyString(o.NatIP))
	p.SetExternalIpv6(dcl.ValueOrEmptyString(o.ExternalIPv6))
	p.SetExternalIpv6PrefixLength(dcl.ValueOrEmptyString(o.ExternalIPv6PrefixLength))
	p.SetSetPublicPtr(dcl.ValueOrEmptyBool(o.SetPublicPtr))
	p.SetPublicPtrDomainName(dcl.ValueOrEmptyString(o.PublicPtrDomainName))
	p.SetNetworkTier(ComputeAlphaInstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto(o.NetworkTier))
	p.SetType(ComputeAlphaInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(o.Type))
	return p
}

// InstanceNetworkInterfacesIPv6AccessConfigsToProto converts a InstanceNetworkInterfacesIPv6AccessConfigs object to its proto representation.
func ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsToProto(o *alpha.InstanceNetworkInterfacesIPv6AccessConfigs) *alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNatIp(dcl.ValueOrEmptyString(o.NatIP))
	p.SetExternalIpv6(dcl.ValueOrEmptyString(o.ExternalIPv6))
	p.SetExternalIpv6PrefixLength(dcl.ValueOrEmptyString(o.ExternalIPv6PrefixLength))
	p.SetSetPublicPtr(dcl.ValueOrEmptyBool(o.SetPublicPtr))
	p.SetPublicPtrDomainName(dcl.ValueOrEmptyString(o.PublicPtrDomainName))
	p.SetNetworkTier(ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto(o.NetworkTier))
	p.SetType(ComputeAlphaInstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto(o.Type))
	return p
}

// InstanceNetworkInterfacesAliasIPRangesToProto converts a InstanceNetworkInterfacesAliasIPRanges object to its proto representation.
func ComputeAlphaInstanceNetworkInterfacesAliasIPRangesToProto(o *alpha.InstanceNetworkInterfacesAliasIPRanges) *alphapb.ComputeAlphaInstanceNetworkInterfacesAliasIPRanges {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceNetworkInterfacesAliasIPRanges{}
	p.SetIpCidrRange(dcl.ValueOrEmptyString(o.IPCidrRange))
	p.SetSubnetworkRangeName(dcl.ValueOrEmptyString(o.SubnetworkRangeName))
	return p
}

// InstanceSchedulingToProto converts a InstanceScheduling object to its proto representation.
func ComputeAlphaInstanceSchedulingToProto(o *alpha.InstanceScheduling) *alphapb.ComputeAlphaInstanceScheduling {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceScheduling{}
	p.SetAutomaticRestart(dcl.ValueOrEmptyBool(o.AutomaticRestart))
	p.SetOnHostMaintenance(dcl.ValueOrEmptyString(o.OnHostMaintenance))
	p.SetPreemptible(dcl.ValueOrEmptyBool(o.Preemptible))
	return p
}

// InstanceServiceAccountsToProto converts a InstanceServiceAccounts object to its proto representation.
func ComputeAlphaInstanceServiceAccountsToProto(o *alpha.InstanceServiceAccounts) *alphapb.ComputeAlphaInstanceServiceAccounts {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceServiceAccounts{}
	p.SetEmail(dcl.ValueOrEmptyString(o.Email))
	sScopes := make([]string, len(o.Scopes))
	for i, r := range o.Scopes {
		sScopes[i] = r
	}
	p.SetScopes(sScopes)
	return p
}

// InstanceShieldedInstanceConfigToProto converts a InstanceShieldedInstanceConfig object to its proto representation.
func ComputeAlphaInstanceShieldedInstanceConfigToProto(o *alpha.InstanceShieldedInstanceConfig) *alphapb.ComputeAlphaInstanceShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceShieldedInstanceConfig{}
	p.SetEnableSecureBoot(dcl.ValueOrEmptyBool(o.EnableSecureBoot))
	p.SetEnableVtpm(dcl.ValueOrEmptyBool(o.EnableVtpm))
	p.SetEnableIntegrityMonitoring(dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring))
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *alpha.Instance) *alphapb.ComputeAlphaInstance {
	p := &alphapb.ComputeAlphaInstance{}
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
	p.SetScheduling(ComputeAlphaInstanceSchedulingToProto(resource.Scheduling))
	p.SetShieldedInstanceConfig(ComputeAlphaInstanceShieldedInstanceConfigToProto(resource.ShieldedInstanceConfig))
	p.SetStatus(ComputeAlphaInstanceStatusEnumToProto(resource.Status))
	p.SetStatusMessage(dcl.ValueOrEmptyString(resource.StatusMessage))
	p.SetZone(dcl.ValueOrEmptyString(resource.Zone))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	sDisks := make([]*alphapb.ComputeAlphaInstanceDisks, len(resource.Disks))
	for i, r := range resource.Disks {
		sDisks[i] = ComputeAlphaInstanceDisksToProto(&r)
	}
	p.SetDisks(sDisks)
	sGuestAccelerators := make([]*alphapb.ComputeAlphaInstanceGuestAccelerators, len(resource.GuestAccelerators))
	for i, r := range resource.GuestAccelerators {
		sGuestAccelerators[i] = ComputeAlphaInstanceGuestAcceleratorsToProto(&r)
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
	sNetworkInterfaces := make([]*alphapb.ComputeAlphaInstanceNetworkInterfaces, len(resource.NetworkInterfaces))
	for i, r := range resource.NetworkInterfaces {
		sNetworkInterfaces[i] = ComputeAlphaInstanceNetworkInterfacesToProto(&r)
	}
	p.SetNetworkInterfaces(sNetworkInterfaces)
	sServiceAccounts := make([]*alphapb.ComputeAlphaInstanceServiceAccounts, len(resource.ServiceAccounts))
	for i, r := range resource.ServiceAccounts {
		sServiceAccounts[i] = ComputeAlphaInstanceServiceAccountsToProto(&r)
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
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaInstanceRequest) (*alphapb.ComputeAlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyComputeAlphaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyComputeAlphaInstance(ctx context.Context, request *alphapb.ApplyComputeAlphaInstanceRequest) (*alphapb.ComputeAlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteComputeAlphaInstance(ctx context.Context, request *alphapb.DeleteComputeAlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListComputeAlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListComputeAlphaInstance(ctx context.Context, request *alphapb.ListComputeAlphaInstanceRequest) (*alphapb.ListComputeAlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetZone())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
