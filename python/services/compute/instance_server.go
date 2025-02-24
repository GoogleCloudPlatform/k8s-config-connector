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

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceDisksInterfaceEnum converts a InstanceDisksInterfaceEnum enum from its proto representation.
func ProtoToComputeInstanceDisksInterfaceEnum(e computepb.ComputeInstanceDisksInterfaceEnum) *compute.InstanceDisksInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceDisksInterfaceEnum_name[int32(e)]; ok {
		e := compute.InstanceDisksInterfaceEnum(n[len("ComputeInstanceDisksInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksModeEnum converts a InstanceDisksModeEnum enum from its proto representation.
func ProtoToComputeInstanceDisksModeEnum(e computepb.ComputeInstanceDisksModeEnum) *compute.InstanceDisksModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceDisksModeEnum_name[int32(e)]; ok {
		e := compute.InstanceDisksModeEnum(n[len("ComputeInstanceDisksModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisksTypeEnum converts a InstanceDisksTypeEnum enum from its proto representation.
func ProtoToComputeInstanceDisksTypeEnum(e computepb.ComputeInstanceDisksTypeEnum) *compute.InstanceDisksTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceDisksTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceDisksTypeEnum(n[len("ComputeInstanceDisksTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesAccessConfigsNetworkTierEnum converts a InstanceNetworkInterfacesAccessConfigsNetworkTierEnum enum from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(e computepb.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum) *compute.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum_name[int32(e)]; ok {
		e := compute.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum(n[len("ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesAccessConfigsTypeEnum converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(e computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum) *compute.InstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceNetworkInterfacesAccessConfigsTypeEnum(n[len("ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum converts a InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum enum from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(e computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum) *compute.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum_name[int32(e)]; ok {
		e := compute.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(n[len("ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum converts a InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(e computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum) *compute.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(n[len("ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStatusEnum converts a InstanceStatusEnum enum from its proto representation.
func ProtoToComputeInstanceStatusEnum(e computepb.ComputeInstanceStatusEnum) *compute.InstanceStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceStatusEnum_name[int32(e)]; ok {
		e := compute.InstanceStatusEnum(n[len("ComputeInstanceStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceDisks converts a InstanceDisks object from its proto representation.
func ProtoToComputeInstanceDisks(p *computepb.ComputeInstanceDisks) *compute.InstanceDisks {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceDisks{
		AutoDelete:        dcl.Bool(p.GetAutoDelete()),
		Boot:              dcl.Bool(p.GetBoot()),
		DeviceName:        dcl.StringOrNil(p.GetDeviceName()),
		DiskEncryptionKey: ProtoToComputeInstanceDisksDiskEncryptionKey(p.GetDiskEncryptionKey()),
		Index:             dcl.Int64OrNil(p.GetIndex()),
		InitializeParams:  ProtoToComputeInstanceDisksInitializeParams(p.GetInitializeParams()),
		Interface:         ProtoToComputeInstanceDisksInterfaceEnum(p.GetInterface()),
		Mode:              ProtoToComputeInstanceDisksModeEnum(p.GetMode()),
		Source:            dcl.StringOrNil(p.GetSource()),
		Type:              ProtoToComputeInstanceDisksTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceDisksDiskEncryptionKey converts a InstanceDisksDiskEncryptionKey object from its proto representation.
func ProtoToComputeInstanceDisksDiskEncryptionKey(p *computepb.ComputeInstanceDisksDiskEncryptionKey) *compute.InstanceDisksDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceDisksDiskEncryptionKey{
		RawKey:          dcl.StringOrNil(p.GetRawKey()),
		RsaEncryptedKey: dcl.StringOrNil(p.GetRsaEncryptedKey()),
		Sha256:          dcl.StringOrNil(p.GetSha256()),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParams converts a InstanceDisksInitializeParams object from its proto representation.
func ProtoToComputeInstanceDisksInitializeParams(p *computepb.ComputeInstanceDisksInitializeParams) *compute.InstanceDisksInitializeParams {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceDisksInitializeParams{
		DiskName:                 dcl.StringOrNil(p.GetDiskName()),
		DiskSizeGb:               dcl.Int64OrNil(p.GetDiskSizeGb()),
		DiskType:                 dcl.StringOrNil(p.GetDiskType()),
		SourceImage:              dcl.StringOrNil(p.GetSourceImage()),
		SourceImageEncryptionKey: ProtoToComputeInstanceDisksInitializeParamsSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
	}
	return obj
}

// ProtoToInstanceDisksInitializeParamsSourceImageEncryptionKey converts a InstanceDisksInitializeParamsSourceImageEncryptionKey object from its proto representation.
func ProtoToComputeInstanceDisksInitializeParamsSourceImageEncryptionKey(p *computepb.ComputeInstanceDisksInitializeParamsSourceImageEncryptionKey) *compute.InstanceDisksInitializeParamsSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceDisksInitializeParamsSourceImageEncryptionKey{
		RawKey: dcl.StringOrNil(p.GetRawKey()),
		Sha256: dcl.StringOrNil(p.GetSha256()),
	}
	return obj
}

// ProtoToInstanceGuestAccelerators converts a InstanceGuestAccelerators object from its proto representation.
func ProtoToComputeInstanceGuestAccelerators(p *computepb.ComputeInstanceGuestAccelerators) *compute.InstanceGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGuestAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfaces converts a InstanceNetworkInterfaces object from its proto representation.
func ProtoToComputeInstanceNetworkInterfaces(p *computepb.ComputeInstanceNetworkInterfaces) *compute.InstanceNetworkInterfaces {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceNetworkInterfaces{
		Name:       dcl.StringOrNil(p.GetName()),
		Network:    dcl.StringOrNil(p.GetNetwork()),
		NetworkIP:  dcl.StringOrNil(p.GetNetworkIp()),
		Subnetwork: dcl.StringOrNil(p.GetSubnetwork()),
	}
	for _, r := range p.GetAccessConfigs() {
		obj.AccessConfigs = append(obj.AccessConfigs, *ProtoToComputeInstanceNetworkInterfacesAccessConfigs(r))
	}
	for _, r := range p.GetIpv6AccessConfigs() {
		obj.IPv6AccessConfigs = append(obj.IPv6AccessConfigs, *ProtoToComputeInstanceNetworkInterfacesIPv6AccessConfigs(r))
	}
	for _, r := range p.GetAliasIpRanges() {
		obj.AliasIPRanges = append(obj.AliasIPRanges, *ProtoToComputeInstanceNetworkInterfacesAliasIPRanges(r))
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAccessConfigs converts a InstanceNetworkInterfacesAccessConfigs object from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesAccessConfigs(p *computepb.ComputeInstanceNetworkInterfacesAccessConfigs) *compute.InstanceNetworkInterfacesAccessConfigs {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceNetworkInterfacesAccessConfigs{
		Name:                     dcl.StringOrNil(p.GetName()),
		NatIP:                    dcl.StringOrNil(p.GetNatIp()),
		ExternalIPv6:             dcl.StringOrNil(p.GetExternalIpv6()),
		ExternalIPv6PrefixLength: dcl.StringOrNil(p.GetExternalIpv6PrefixLength()),
		SetPublicPtr:             dcl.Bool(p.GetSetPublicPtr()),
		PublicPtrDomainName:      dcl.StringOrNil(p.GetPublicPtrDomainName()),
		NetworkTier:              ProtoToComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(p.GetNetworkTier()),
		Type:                     ProtoToComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesIPv6AccessConfigs converts a InstanceNetworkInterfacesIPv6AccessConfigs object from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesIPv6AccessConfigs(p *computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigs) *compute.InstanceNetworkInterfacesIPv6AccessConfigs {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceNetworkInterfacesIPv6AccessConfigs{
		Name:                     dcl.StringOrNil(p.GetName()),
		NatIP:                    dcl.StringOrNil(p.GetNatIp()),
		ExternalIPv6:             dcl.StringOrNil(p.GetExternalIpv6()),
		ExternalIPv6PrefixLength: dcl.StringOrNil(p.GetExternalIpv6PrefixLength()),
		SetPublicPtr:             dcl.Bool(p.GetSetPublicPtr()),
		PublicPtrDomainName:      dcl.StringOrNil(p.GetPublicPtrDomainName()),
		NetworkTier:              ProtoToComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(p.GetNetworkTier()),
		Type:                     ProtoToComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToInstanceNetworkInterfacesAliasIPRanges converts a InstanceNetworkInterfacesAliasIPRanges object from its proto representation.
func ProtoToComputeInstanceNetworkInterfacesAliasIPRanges(p *computepb.ComputeInstanceNetworkInterfacesAliasIPRanges) *compute.InstanceNetworkInterfacesAliasIPRanges {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceNetworkInterfacesAliasIPRanges{
		IPCidrRange:         dcl.StringOrNil(p.GetIpCidrRange()),
		SubnetworkRangeName: dcl.StringOrNil(p.GetSubnetworkRangeName()),
	}
	return obj
}

// ProtoToInstanceScheduling converts a InstanceScheduling object from its proto representation.
func ProtoToComputeInstanceScheduling(p *computepb.ComputeInstanceScheduling) *compute.InstanceScheduling {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceScheduling{
		AutomaticRestart:  dcl.Bool(p.GetAutomaticRestart()),
		OnHostMaintenance: dcl.StringOrNil(p.GetOnHostMaintenance()),
		Preemptible:       dcl.Bool(p.GetPreemptible()),
	}
	return obj
}

// ProtoToInstanceServiceAccounts converts a InstanceServiceAccounts object from its proto representation.
func ProtoToComputeInstanceServiceAccounts(p *computepb.ComputeInstanceServiceAccounts) *compute.InstanceServiceAccounts {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceServiceAccounts{
		Email: dcl.StringOrNil(p.GetEmail()),
	}
	for _, r := range p.GetScopes() {
		obj.Scopes = append(obj.Scopes, r)
	}
	return obj
}

// ProtoToInstanceShieldedInstanceConfig converts a InstanceShieldedInstanceConfig object from its proto representation.
func ProtoToComputeInstanceShieldedInstanceConfig(p *computepb.ComputeInstanceShieldedInstanceConfig) *compute.InstanceShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.GetEnableSecureBoot()),
		EnableVtpm:                dcl.Bool(p.GetEnableVtpm()),
		EnableIntegrityMonitoring: dcl.Bool(p.GetEnableIntegrityMonitoring()),
	}
	return obj
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *computepb.ComputeInstance) *compute.Instance {
	obj := &compute.Instance{
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
		Scheduling:             ProtoToComputeInstanceScheduling(p.GetScheduling()),
		ShieldedInstanceConfig: ProtoToComputeInstanceShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		Status:                 ProtoToComputeInstanceStatusEnum(p.GetStatus()),
		StatusMessage:          dcl.StringOrNil(p.GetStatusMessage()),
		Zone:                   dcl.StringOrNil(p.GetZone()),
		Project:                dcl.StringOrNil(p.GetProject()),
		SelfLink:               dcl.StringOrNil(p.GetSelfLink()),
	}
	for _, r := range p.GetDisks() {
		obj.Disks = append(obj.Disks, *ProtoToComputeInstanceDisks(r))
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeInstanceGuestAccelerators(r))
	}
	for _, r := range p.GetNetworkInterfaces() {
		obj.NetworkInterfaces = append(obj.NetworkInterfaces, *ProtoToComputeInstanceNetworkInterfaces(r))
	}
	for _, r := range p.GetServiceAccounts() {
		obj.ServiceAccounts = append(obj.ServiceAccounts, *ProtoToComputeInstanceServiceAccounts(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// InstanceDisksInterfaceEnumToProto converts a InstanceDisksInterfaceEnum enum to its proto representation.
func ComputeInstanceDisksInterfaceEnumToProto(e *compute.InstanceDisksInterfaceEnum) computepb.ComputeInstanceDisksInterfaceEnum {
	if e == nil {
		return computepb.ComputeInstanceDisksInterfaceEnum(0)
	}
	if v, ok := computepb.ComputeInstanceDisksInterfaceEnum_value["InstanceDisksInterfaceEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceDisksInterfaceEnum(v)
	}
	return computepb.ComputeInstanceDisksInterfaceEnum(0)
}

// InstanceDisksModeEnumToProto converts a InstanceDisksModeEnum enum to its proto representation.
func ComputeInstanceDisksModeEnumToProto(e *compute.InstanceDisksModeEnum) computepb.ComputeInstanceDisksModeEnum {
	if e == nil {
		return computepb.ComputeInstanceDisksModeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceDisksModeEnum_value["InstanceDisksModeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceDisksModeEnum(v)
	}
	return computepb.ComputeInstanceDisksModeEnum(0)
}

// InstanceDisksTypeEnumToProto converts a InstanceDisksTypeEnum enum to its proto representation.
func ComputeInstanceDisksTypeEnumToProto(e *compute.InstanceDisksTypeEnum) computepb.ComputeInstanceDisksTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceDisksTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceDisksTypeEnum_value["InstanceDisksTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceDisksTypeEnum(v)
	}
	return computepb.ComputeInstanceDisksTypeEnum(0)
}

// InstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto converts a InstanceNetworkInterfacesAccessConfigsNetworkTierEnum enum to its proto representation.
func ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto(e *compute.InstanceNetworkInterfacesAccessConfigsNetworkTierEnum) computepb.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == nil {
		return computepb.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(0)
	}
	if v, ok := computepb.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum_value["InstanceNetworkInterfacesAccessConfigsNetworkTierEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(v)
	}
	return computepb.ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnum(0)
}

// InstanceNetworkInterfacesAccessConfigsTypeEnumToProto converts a InstanceNetworkInterfacesAccessConfigsTypeEnum enum to its proto representation.
func ComputeInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(e *compute.InstanceNetworkInterfacesAccessConfigsTypeEnum) computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum_value["InstanceNetworkInterfacesAccessConfigsTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(v)
	}
	return computepb.ComputeInstanceNetworkInterfacesAccessConfigsTypeEnum(0)
}

// InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto converts a InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum enum to its proto representation.
func ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto(e *compute.InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum) computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum {
	if e == nil {
		return computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(0)
	}
	if v, ok := computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum_value["InstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(v)
	}
	return computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnum(0)
}

// InstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto converts a InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum enum to its proto representation.
func ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto(e *compute.InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum) computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum_value["InstanceNetworkInterfacesIPv6AccessConfigsTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(v)
	}
	return computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnum(0)
}

// InstanceStatusEnumToProto converts a InstanceStatusEnum enum to its proto representation.
func ComputeInstanceStatusEnumToProto(e *compute.InstanceStatusEnum) computepb.ComputeInstanceStatusEnum {
	if e == nil {
		return computepb.ComputeInstanceStatusEnum(0)
	}
	if v, ok := computepb.ComputeInstanceStatusEnum_value["InstanceStatusEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceStatusEnum(v)
	}
	return computepb.ComputeInstanceStatusEnum(0)
}

// InstanceDisksToProto converts a InstanceDisks object to its proto representation.
func ComputeInstanceDisksToProto(o *compute.InstanceDisks) *computepb.ComputeInstanceDisks {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceDisks{}
	p.SetAutoDelete(dcl.ValueOrEmptyBool(o.AutoDelete))
	p.SetBoot(dcl.ValueOrEmptyBool(o.Boot))
	p.SetDeviceName(dcl.ValueOrEmptyString(o.DeviceName))
	p.SetDiskEncryptionKey(ComputeInstanceDisksDiskEncryptionKeyToProto(o.DiskEncryptionKey))
	p.SetIndex(dcl.ValueOrEmptyInt64(o.Index))
	p.SetInitializeParams(ComputeInstanceDisksInitializeParamsToProto(o.InitializeParams))
	p.SetInterface(ComputeInstanceDisksInterfaceEnumToProto(o.Interface))
	p.SetMode(ComputeInstanceDisksModeEnumToProto(o.Mode))
	p.SetSource(dcl.ValueOrEmptyString(o.Source))
	p.SetType(ComputeInstanceDisksTypeEnumToProto(o.Type))
	return p
}

// InstanceDisksDiskEncryptionKeyToProto converts a InstanceDisksDiskEncryptionKey object to its proto representation.
func ComputeInstanceDisksDiskEncryptionKeyToProto(o *compute.InstanceDisksDiskEncryptionKey) *computepb.ComputeInstanceDisksDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceDisksDiskEncryptionKey{}
	p.SetRawKey(dcl.ValueOrEmptyString(o.RawKey))
	p.SetRsaEncryptedKey(dcl.ValueOrEmptyString(o.RsaEncryptedKey))
	p.SetSha256(dcl.ValueOrEmptyString(o.Sha256))
	return p
}

// InstanceDisksInitializeParamsToProto converts a InstanceDisksInitializeParams object to its proto representation.
func ComputeInstanceDisksInitializeParamsToProto(o *compute.InstanceDisksInitializeParams) *computepb.ComputeInstanceDisksInitializeParams {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceDisksInitializeParams{}
	p.SetDiskName(dcl.ValueOrEmptyString(o.DiskName))
	p.SetDiskSizeGb(dcl.ValueOrEmptyInt64(o.DiskSizeGb))
	p.SetDiskType(dcl.ValueOrEmptyString(o.DiskType))
	p.SetSourceImage(dcl.ValueOrEmptyString(o.SourceImage))
	p.SetSourceImageEncryptionKey(ComputeInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o.SourceImageEncryptionKey))
	return p
}

// InstanceDisksInitializeParamsSourceImageEncryptionKeyToProto converts a InstanceDisksInitializeParamsSourceImageEncryptionKey object to its proto representation.
func ComputeInstanceDisksInitializeParamsSourceImageEncryptionKeyToProto(o *compute.InstanceDisksInitializeParamsSourceImageEncryptionKey) *computepb.ComputeInstanceDisksInitializeParamsSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceDisksInitializeParamsSourceImageEncryptionKey{}
	p.SetRawKey(dcl.ValueOrEmptyString(o.RawKey))
	p.SetSha256(dcl.ValueOrEmptyString(o.Sha256))
	return p
}

// InstanceGuestAcceleratorsToProto converts a InstanceGuestAccelerators object to its proto representation.
func ComputeInstanceGuestAcceleratorsToProto(o *compute.InstanceGuestAccelerators) *computepb.ComputeInstanceGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGuestAccelerators{}
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	return p
}

// InstanceNetworkInterfacesToProto converts a InstanceNetworkInterfaces object to its proto representation.
func ComputeInstanceNetworkInterfacesToProto(o *compute.InstanceNetworkInterfaces) *computepb.ComputeInstanceNetworkInterfaces {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceNetworkInterfaces{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetNetworkIp(dcl.ValueOrEmptyString(o.NetworkIP))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	sAccessConfigs := make([]*computepb.ComputeInstanceNetworkInterfacesAccessConfigs, len(o.AccessConfigs))
	for i, r := range o.AccessConfigs {
		sAccessConfigs[i] = ComputeInstanceNetworkInterfacesAccessConfigsToProto(&r)
	}
	p.SetAccessConfigs(sAccessConfigs)
	sIPv6AccessConfigs := make([]*computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigs, len(o.IPv6AccessConfigs))
	for i, r := range o.IPv6AccessConfigs {
		sIPv6AccessConfigs[i] = ComputeInstanceNetworkInterfacesIPv6AccessConfigsToProto(&r)
	}
	p.SetIpv6AccessConfigs(sIPv6AccessConfigs)
	sAliasIPRanges := make([]*computepb.ComputeInstanceNetworkInterfacesAliasIPRanges, len(o.AliasIPRanges))
	for i, r := range o.AliasIPRanges {
		sAliasIPRanges[i] = ComputeInstanceNetworkInterfacesAliasIPRangesToProto(&r)
	}
	p.SetAliasIpRanges(sAliasIPRanges)
	return p
}

// InstanceNetworkInterfacesAccessConfigsToProto converts a InstanceNetworkInterfacesAccessConfigs object to its proto representation.
func ComputeInstanceNetworkInterfacesAccessConfigsToProto(o *compute.InstanceNetworkInterfacesAccessConfigs) *computepb.ComputeInstanceNetworkInterfacesAccessConfigs {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceNetworkInterfacesAccessConfigs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNatIp(dcl.ValueOrEmptyString(o.NatIP))
	p.SetExternalIpv6(dcl.ValueOrEmptyString(o.ExternalIPv6))
	p.SetExternalIpv6PrefixLength(dcl.ValueOrEmptyString(o.ExternalIPv6PrefixLength))
	p.SetSetPublicPtr(dcl.ValueOrEmptyBool(o.SetPublicPtr))
	p.SetPublicPtrDomainName(dcl.ValueOrEmptyString(o.PublicPtrDomainName))
	p.SetNetworkTier(ComputeInstanceNetworkInterfacesAccessConfigsNetworkTierEnumToProto(o.NetworkTier))
	p.SetType(ComputeInstanceNetworkInterfacesAccessConfigsTypeEnumToProto(o.Type))
	return p
}

// InstanceNetworkInterfacesIPv6AccessConfigsToProto converts a InstanceNetworkInterfacesIPv6AccessConfigs object to its proto representation.
func ComputeInstanceNetworkInterfacesIPv6AccessConfigsToProto(o *compute.InstanceNetworkInterfacesIPv6AccessConfigs) *computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigs {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceNetworkInterfacesIPv6AccessConfigs{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetNatIp(dcl.ValueOrEmptyString(o.NatIP))
	p.SetExternalIpv6(dcl.ValueOrEmptyString(o.ExternalIPv6))
	p.SetExternalIpv6PrefixLength(dcl.ValueOrEmptyString(o.ExternalIPv6PrefixLength))
	p.SetSetPublicPtr(dcl.ValueOrEmptyBool(o.SetPublicPtr))
	p.SetPublicPtrDomainName(dcl.ValueOrEmptyString(o.PublicPtrDomainName))
	p.SetNetworkTier(ComputeInstanceNetworkInterfacesIPv6AccessConfigsNetworkTierEnumToProto(o.NetworkTier))
	p.SetType(ComputeInstanceNetworkInterfacesIPv6AccessConfigsTypeEnumToProto(o.Type))
	return p
}

// InstanceNetworkInterfacesAliasIPRangesToProto converts a InstanceNetworkInterfacesAliasIPRanges object to its proto representation.
func ComputeInstanceNetworkInterfacesAliasIPRangesToProto(o *compute.InstanceNetworkInterfacesAliasIPRanges) *computepb.ComputeInstanceNetworkInterfacesAliasIPRanges {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceNetworkInterfacesAliasIPRanges{}
	p.SetIpCidrRange(dcl.ValueOrEmptyString(o.IPCidrRange))
	p.SetSubnetworkRangeName(dcl.ValueOrEmptyString(o.SubnetworkRangeName))
	return p
}

// InstanceSchedulingToProto converts a InstanceScheduling object to its proto representation.
func ComputeInstanceSchedulingToProto(o *compute.InstanceScheduling) *computepb.ComputeInstanceScheduling {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceScheduling{}
	p.SetAutomaticRestart(dcl.ValueOrEmptyBool(o.AutomaticRestart))
	p.SetOnHostMaintenance(dcl.ValueOrEmptyString(o.OnHostMaintenance))
	p.SetPreemptible(dcl.ValueOrEmptyBool(o.Preemptible))
	return p
}

// InstanceServiceAccountsToProto converts a InstanceServiceAccounts object to its proto representation.
func ComputeInstanceServiceAccountsToProto(o *compute.InstanceServiceAccounts) *computepb.ComputeInstanceServiceAccounts {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceServiceAccounts{}
	p.SetEmail(dcl.ValueOrEmptyString(o.Email))
	sScopes := make([]string, len(o.Scopes))
	for i, r := range o.Scopes {
		sScopes[i] = r
	}
	p.SetScopes(sScopes)
	return p
}

// InstanceShieldedInstanceConfigToProto converts a InstanceShieldedInstanceConfig object to its proto representation.
func ComputeInstanceShieldedInstanceConfigToProto(o *compute.InstanceShieldedInstanceConfig) *computepb.ComputeInstanceShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceShieldedInstanceConfig{}
	p.SetEnableSecureBoot(dcl.ValueOrEmptyBool(o.EnableSecureBoot))
	p.SetEnableVtpm(dcl.ValueOrEmptyBool(o.EnableVtpm))
	p.SetEnableIntegrityMonitoring(dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring))
	return p
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *compute.Instance) *computepb.ComputeInstance {
	p := &computepb.ComputeInstance{}
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
	p.SetScheduling(ComputeInstanceSchedulingToProto(resource.Scheduling))
	p.SetShieldedInstanceConfig(ComputeInstanceShieldedInstanceConfigToProto(resource.ShieldedInstanceConfig))
	p.SetStatus(ComputeInstanceStatusEnumToProto(resource.Status))
	p.SetStatusMessage(dcl.ValueOrEmptyString(resource.StatusMessage))
	p.SetZone(dcl.ValueOrEmptyString(resource.Zone))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	sDisks := make([]*computepb.ComputeInstanceDisks, len(resource.Disks))
	for i, r := range resource.Disks {
		sDisks[i] = ComputeInstanceDisksToProto(&r)
	}
	p.SetDisks(sDisks)
	sGuestAccelerators := make([]*computepb.ComputeInstanceGuestAccelerators, len(resource.GuestAccelerators))
	for i, r := range resource.GuestAccelerators {
		sGuestAccelerators[i] = ComputeInstanceGuestAcceleratorsToProto(&r)
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
	sNetworkInterfaces := make([]*computepb.ComputeInstanceNetworkInterfaces, len(resource.NetworkInterfaces))
	for i, r := range resource.NetworkInterfaces {
		sNetworkInterfaces[i] = ComputeInstanceNetworkInterfacesToProto(&r)
	}
	p.SetNetworkInterfaces(sNetworkInterfaces)
	sServiceAccounts := make([]*computepb.ComputeInstanceServiceAccounts, len(resource.ServiceAccounts))
	for i, r := range resource.ServiceAccounts {
		sServiceAccounts[i] = ComputeInstanceServiceAccountsToProto(&r)
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
func (s *InstanceServer) applyInstance(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeInstanceRequest) (*computepb.ComputeInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyComputeInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyComputeInstance(ctx context.Context, request *computepb.ApplyComputeInstanceRequest) (*computepb.ComputeInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteComputeInstance(ctx context.Context, request *computepb.DeleteComputeInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListComputeInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListComputeInstance(ctx context.Context, request *computepb.ListComputeInstanceRequest) (*computepb.ListComputeInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetProject(), request.GetZone())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
