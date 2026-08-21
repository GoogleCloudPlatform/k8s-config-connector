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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for InstanceTemplate.
type InstanceTemplateServer struct{}

// ProtoToInstanceTemplatePropertiesDisksInterfaceEnum converts a InstanceTemplatePropertiesDisksInterfaceEnum enum from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisksInterfaceEnum(e computepb.ComputeInstanceTemplatePropertiesDisksInterfaceEnum) *compute.InstanceTemplatePropertiesDisksInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceTemplatePropertiesDisksInterfaceEnum_name[int32(e)]; ok {
		e := compute.InstanceTemplatePropertiesDisksInterfaceEnum(n[len("ComputeInstanceTemplatePropertiesDisksInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesDisksModeEnum converts a InstanceTemplatePropertiesDisksModeEnum enum from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisksModeEnum(e computepb.ComputeInstanceTemplatePropertiesDisksModeEnum) *compute.InstanceTemplatePropertiesDisksModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceTemplatePropertiesDisksModeEnum_name[int32(e)]; ok {
		e := compute.InstanceTemplatePropertiesDisksModeEnum(n[len("ComputeInstanceTemplatePropertiesDisksModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesDisksTypeEnum converts a InstanceTemplatePropertiesDisksTypeEnum enum from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisksTypeEnum(e computepb.ComputeInstanceTemplatePropertiesDisksTypeEnum) *compute.InstanceTemplatePropertiesDisksTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceTemplatePropertiesDisksTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceTemplatePropertiesDisksTypeEnum(n[len("ComputeInstanceTemplatePropertiesDisksTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(e computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum) *compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(n[len("ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum enum from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(e computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum) *compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum_name[int32(e)]; ok {
		e := compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(n[len("ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum converts a InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum enum from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(e computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum) *compute.InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum_name[int32(e)]; ok {
		e := compute.InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(n[len("ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplateProperties converts a InstanceTemplateProperties resource from its proto representation.
func ProtoToComputeInstanceTemplateProperties(p *computepb.ComputeInstanceTemplateProperties) *compute.InstanceTemplateProperties {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplateProperties{
		CanIPForward:           dcl.Bool(p.CanIpForward),
		Description:            dcl.StringOrNil(p.Description),
		MachineType:            dcl.StringOrNil(p.MachineType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		ReservationAffinity:    ProtoToComputeInstanceTemplatePropertiesReservationAffinity(p.GetReservationAffinity()),
		ShieldedInstanceConfig: ProtoToComputeInstanceTemplatePropertiesShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		Scheduling:             ProtoToComputeInstanceTemplatePropertiesScheduling(p.GetScheduling()),
	}
	for _, r := range p.GetDisks() {
		obj.Disks = append(obj.Disks, *ProtoToComputeInstanceTemplatePropertiesDisks(r))
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeInstanceTemplatePropertiesGuestAccelerators(r))
	}
	for _, r := range p.GetNetworkInterfaces() {
		obj.NetworkInterfaces = append(obj.NetworkInterfaces, *ProtoToComputeInstanceTemplatePropertiesNetworkInterfaces(r))
	}
	for _, r := range p.GetServiceAccounts() {
		obj.ServiceAccounts = append(obj.ServiceAccounts, *ProtoToComputeInstanceTemplatePropertiesServiceAccounts(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisks converts a InstanceTemplatePropertiesDisks resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisks(p *computepb.ComputeInstanceTemplatePropertiesDisks) *compute.InstanceTemplatePropertiesDisks {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesDisks{
		AutoDelete:        dcl.Bool(p.AutoDelete),
		Boot:              dcl.Bool(p.Boot),
		DeviceName:        dcl.StringOrNil(p.DeviceName),
		DiskEncryptionKey: ProtoToComputeInstanceTemplatePropertiesDisksDiskEncryptionKey(p.GetDiskEncryptionKey()),
		Index:             dcl.Int64OrNil(p.Index),
		InitializeParams:  ProtoToComputeInstanceTemplatePropertiesDisksInitializeParams(p.GetInitializeParams()),
		Interface:         ProtoToComputeInstanceTemplatePropertiesDisksInterfaceEnum(p.GetInterface()),
		Mode:              ProtoToComputeInstanceTemplatePropertiesDisksModeEnum(p.GetMode()),
		Source:            dcl.StringOrNil(p.Source),
		Type:              ProtoToComputeInstanceTemplatePropertiesDisksTypeEnum(p.GetType()),
	}
	for _, r := range p.GetGuestOsFeatures() {
		obj.GuestOSFeatures = append(obj.GuestOSFeatures, *ProtoToComputeInstanceTemplatePropertiesDisksGuestOSFeatures(r))
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksDiskEncryptionKey converts a InstanceTemplatePropertiesDisksDiskEncryptionKey resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisksDiskEncryptionKey(p *computepb.ComputeInstanceTemplatePropertiesDisksDiskEncryptionKey) *compute.InstanceTemplatePropertiesDisksDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesDisksDiskEncryptionKey{
		RawKey:          dcl.StringOrNil(p.RawKey),
		RsaEncryptedKey: dcl.StringOrNil(p.RsaEncryptedKey),
		Sha256:          dcl.StringOrNil(p.Sha256),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksInitializeParams converts a InstanceTemplatePropertiesDisksInitializeParams resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisksInitializeParams(p *computepb.ComputeInstanceTemplatePropertiesDisksInitializeParams) *compute.InstanceTemplatePropertiesDisksInitializeParams {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesDisksInitializeParams{
		DiskName:                    dcl.StringOrNil(p.DiskName),
		DiskSizeGb:                  dcl.Int64OrNil(p.DiskSizeGb),
		DiskType:                    dcl.StringOrNil(p.DiskType),
		SourceImage:                 dcl.StringOrNil(p.SourceImage),
		SourceSnapshot:              dcl.StringOrNil(p.SourceSnapshot),
		SourceSnapshotEncryptionKey: ProtoToComputeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(p.GetSourceSnapshotEncryptionKey()),
		Description:                 dcl.StringOrNil(p.Description),
		OnUpdateAction:              dcl.StringOrNil(p.OnUpdateAction),
		SourceImageEncryptionKey:    ProtoToComputeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
	}
	for _, r := range p.GetResourcePolicies() {
		obj.ResourcePolicies = append(obj.ResourcePolicies, r)
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey converts a InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(p *computepb.ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) *compute.InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{
		RawKey:     dcl.StringOrNil(p.RawKey),
		Sha256:     dcl.StringOrNil(p.Sha256),
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey converts a InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(p *computepb.ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) *compute.InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{
		RawKey:     dcl.StringOrNil(p.RawKey),
		Sha256:     dcl.StringOrNil(p.Sha256),
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksGuestOSFeatures converts a InstanceTemplatePropertiesDisksGuestOSFeatures resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesDisksGuestOSFeatures(p *computepb.ComputeInstanceTemplatePropertiesDisksGuestOSFeatures) *compute.InstanceTemplatePropertiesDisksGuestOSFeatures {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesDisksGuestOSFeatures{
		Type: dcl.StringOrNil(p.Type),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesReservationAffinity converts a InstanceTemplatePropertiesReservationAffinity resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesReservationAffinity(p *computepb.ComputeInstanceTemplatePropertiesReservationAffinity) *compute.InstanceTemplatePropertiesReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesReservationAffinity{
		Key: dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValue() {
		obj.Value = append(obj.Value, r)
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesGuestAccelerators converts a InstanceTemplatePropertiesGuestAccelerators resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesGuestAccelerators(p *computepb.ComputeInstanceTemplatePropertiesGuestAccelerators) *compute.InstanceTemplatePropertiesGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesGuestAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesNetworkInterfaces converts a InstanceTemplatePropertiesNetworkInterfaces resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesNetworkInterfaces(p *computepb.ComputeInstanceTemplatePropertiesNetworkInterfaces) *compute.InstanceTemplatePropertiesNetworkInterfaces {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesNetworkInterfaces{
		Name:       dcl.StringOrNil(p.Name),
		Network:    dcl.StringOrNil(p.Network),
		NetworkIP:  dcl.StringOrNil(p.NetworkIp),
		Subnetwork: dcl.StringOrNil(p.Subnetwork),
	}
	for _, r := range p.GetAccessConfigs() {
		obj.AccessConfigs = append(obj.AccessConfigs, *ProtoToComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(r))
	}
	for _, r := range p.GetAliasIpRanges() {
		obj.AliasIPRanges = append(obj.AliasIPRanges, *ProtoToComputeInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(r))
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesNetworkInterfacesAccessConfigs converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigs resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(p *computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigs) *compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigs{
		Name:                dcl.StringOrNil(p.Name),
		NatIP:               dcl.StringOrNil(p.NatIp),
		Type:                ProtoToComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(p.GetType()),
		SetPublicPtr:        dcl.Bool(p.SetPublicPtr),
		PublicPtrDomainName: dcl.StringOrNil(p.PublicPtrDomainName),
		NetworkTier:         ProtoToComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(p.GetNetworkTier()),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges converts a InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(p *computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) *compute.InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{
		IPCidrRange:         dcl.StringOrNil(p.IpCidrRange),
		SubnetworkRangeName: dcl.StringOrNil(p.SubnetworkRangeName),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesShieldedInstanceConfig converts a InstanceTemplatePropertiesShieldedInstanceConfig resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesShieldedInstanceConfig(p *computepb.ComputeInstanceTemplatePropertiesShieldedInstanceConfig) *compute.InstanceTemplatePropertiesShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableVtpm:                dcl.Bool(p.EnableVtpm),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesScheduling converts a InstanceTemplatePropertiesScheduling resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesScheduling(p *computepb.ComputeInstanceTemplatePropertiesScheduling) *compute.InstanceTemplatePropertiesScheduling {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesScheduling{
		AutomaticRestart:  dcl.Bool(p.AutomaticRestart),
		OnHostMaintenance: dcl.StringOrNil(p.OnHostMaintenance),
		Preemptible:       dcl.Bool(p.Preemptible),
	}
	for _, r := range p.GetNodeAffinities() {
		obj.NodeAffinities = append(obj.NodeAffinities, *ProtoToComputeInstanceTemplatePropertiesSchedulingNodeAffinities(r))
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesSchedulingNodeAffinities converts a InstanceTemplatePropertiesSchedulingNodeAffinities resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesSchedulingNodeAffinities(p *computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinities) *compute.InstanceTemplatePropertiesSchedulingNodeAffinities {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesSchedulingNodeAffinities{
		Key:      dcl.StringOrNil(p.Key),
		Operator: ProtoToComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(p.GetOperator()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesServiceAccounts converts a InstanceTemplatePropertiesServiceAccounts resource from its proto representation.
func ProtoToComputeInstanceTemplatePropertiesServiceAccounts(p *computepb.ComputeInstanceTemplatePropertiesServiceAccounts) *compute.InstanceTemplatePropertiesServiceAccounts {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceTemplatePropertiesServiceAccounts{
		Email: dcl.StringOrNil(p.Email),
	}
	for _, r := range p.GetScopes() {
		obj.Scopes = append(obj.Scopes, r)
	}
	return obj
}

// ProtoToInstanceTemplate converts a InstanceTemplate resource from its proto representation.
func ProtoToInstanceTemplate(p *computepb.ComputeInstanceTemplate) *compute.InstanceTemplate {
	obj := &compute.InstanceTemplate{
		CreationTimestamp: dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:       dcl.StringOrNil(p.Description),
		Id:                dcl.Int64OrNil(p.Id),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Name:              dcl.StringOrNil(p.Name),
		Properties:        ProtoToComputeInstanceTemplateProperties(p.GetProperties()),
		Project:           dcl.StringOrNil(p.Project),
	}
	return obj
}

// InstanceTemplatePropertiesDisksInterfaceEnumToProto converts a InstanceTemplatePropertiesDisksInterfaceEnum enum to its proto representation.
func ComputeInstanceTemplatePropertiesDisksInterfaceEnumToProto(e *compute.InstanceTemplatePropertiesDisksInterfaceEnum) computepb.ComputeInstanceTemplatePropertiesDisksInterfaceEnum {
	if e == nil {
		return computepb.ComputeInstanceTemplatePropertiesDisksInterfaceEnum(0)
	}
	if v, ok := computepb.ComputeInstanceTemplatePropertiesDisksInterfaceEnum_value["InstanceTemplatePropertiesDisksInterfaceEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceTemplatePropertiesDisksInterfaceEnum(v)
	}
	return computepb.ComputeInstanceTemplatePropertiesDisksInterfaceEnum(0)
}

// InstanceTemplatePropertiesDisksModeEnumToProto converts a InstanceTemplatePropertiesDisksModeEnum enum to its proto representation.
func ComputeInstanceTemplatePropertiesDisksModeEnumToProto(e *compute.InstanceTemplatePropertiesDisksModeEnum) computepb.ComputeInstanceTemplatePropertiesDisksModeEnum {
	if e == nil {
		return computepb.ComputeInstanceTemplatePropertiesDisksModeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceTemplatePropertiesDisksModeEnum_value["InstanceTemplatePropertiesDisksModeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceTemplatePropertiesDisksModeEnum(v)
	}
	return computepb.ComputeInstanceTemplatePropertiesDisksModeEnum(0)
}

// InstanceTemplatePropertiesDisksTypeEnumToProto converts a InstanceTemplatePropertiesDisksTypeEnum enum to its proto representation.
func ComputeInstanceTemplatePropertiesDisksTypeEnumToProto(e *compute.InstanceTemplatePropertiesDisksTypeEnum) computepb.ComputeInstanceTemplatePropertiesDisksTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceTemplatePropertiesDisksTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceTemplatePropertiesDisksTypeEnum_value["InstanceTemplatePropertiesDisksTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceTemplatePropertiesDisksTypeEnum(v)
	}
	return computepb.ComputeInstanceTemplatePropertiesDisksTypeEnum(0)
}

// InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumToProto converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum enum to its proto representation.
func ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumToProto(e *compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum) computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum_value["InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(v)
	}
	return computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(0)
}

// InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumToProto converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum enum to its proto representation.
func ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumToProto(e *compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum) computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == nil {
		return computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(0)
	}
	if v, ok := computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum_value["InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(v)
	}
	return computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(0)
}

// InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumToProto converts a InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum enum to its proto representation.
func ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumToProto(e *compute.InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum) computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum {
	if e == nil {
		return computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(0)
	}
	if v, ok := computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum_value["InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(v)
	}
	return computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(0)
}

// InstanceTemplatePropertiesToProto converts a InstanceTemplateProperties resource to its proto representation.
func ComputeInstanceTemplatePropertiesToProto(o *compute.InstanceTemplateProperties) *computepb.ComputeInstanceTemplateProperties {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplateProperties{
		CanIpForward:           dcl.ValueOrEmptyBool(o.CanIPForward),
		Description:            dcl.ValueOrEmptyString(o.Description),
		MachineType:            dcl.ValueOrEmptyString(o.MachineType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		ReservationAffinity:    ComputeInstanceTemplatePropertiesReservationAffinityToProto(o.ReservationAffinity),
		ShieldedInstanceConfig: ComputeInstanceTemplatePropertiesShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
		Scheduling:             ComputeInstanceTemplatePropertiesSchedulingToProto(o.Scheduling),
	}
	for _, r := range o.Disks {
		p.Disks = append(p.Disks, ComputeInstanceTemplatePropertiesDisksToProto(&r))
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	p.Metadata = make(map[string]string)
	for k, r := range o.Metadata {
		p.Metadata[k] = r
	}
	for _, r := range o.GuestAccelerators {
		p.GuestAccelerators = append(p.GuestAccelerators, ComputeInstanceTemplatePropertiesGuestAcceleratorsToProto(&r))
	}
	for _, r := range o.NetworkInterfaces {
		p.NetworkInterfaces = append(p.NetworkInterfaces, ComputeInstanceTemplatePropertiesNetworkInterfacesToProto(&r))
	}
	for _, r := range o.ServiceAccounts {
		p.ServiceAccounts = append(p.ServiceAccounts, ComputeInstanceTemplatePropertiesServiceAccountsToProto(&r))
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	return p
}

// InstanceTemplatePropertiesDisksToProto converts a InstanceTemplatePropertiesDisks resource to its proto representation.
func ComputeInstanceTemplatePropertiesDisksToProto(o *compute.InstanceTemplatePropertiesDisks) *computepb.ComputeInstanceTemplatePropertiesDisks {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesDisks{
		AutoDelete:        dcl.ValueOrEmptyBool(o.AutoDelete),
		Boot:              dcl.ValueOrEmptyBool(o.Boot),
		DeviceName:        dcl.ValueOrEmptyString(o.DeviceName),
		DiskEncryptionKey: ComputeInstanceTemplatePropertiesDisksDiskEncryptionKeyToProto(o.DiskEncryptionKey),
		Index:             dcl.ValueOrEmptyInt64(o.Index),
		InitializeParams:  ComputeInstanceTemplatePropertiesDisksInitializeParamsToProto(o.InitializeParams),
		Interface:         ComputeInstanceTemplatePropertiesDisksInterfaceEnumToProto(o.Interface),
		Mode:              ComputeInstanceTemplatePropertiesDisksModeEnumToProto(o.Mode),
		Source:            dcl.ValueOrEmptyString(o.Source),
		Type:              ComputeInstanceTemplatePropertiesDisksTypeEnumToProto(o.Type),
	}
	for _, r := range o.GuestOSFeatures {
		p.GuestOsFeatures = append(p.GuestOsFeatures, ComputeInstanceTemplatePropertiesDisksGuestOSFeaturesToProto(&r))
	}
	return p
}

// InstanceTemplatePropertiesDisksDiskEncryptionKeyToProto converts a InstanceTemplatePropertiesDisksDiskEncryptionKey resource to its proto representation.
func ComputeInstanceTemplatePropertiesDisksDiskEncryptionKeyToProto(o *compute.InstanceTemplatePropertiesDisksDiskEncryptionKey) *computepb.ComputeInstanceTemplatePropertiesDisksDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesDisksDiskEncryptionKey{
		RawKey:          dcl.ValueOrEmptyString(o.RawKey),
		RsaEncryptedKey: dcl.ValueOrEmptyString(o.RsaEncryptedKey),
		Sha256:          dcl.ValueOrEmptyString(o.Sha256),
	}
	return p
}

// InstanceTemplatePropertiesDisksInitializeParamsToProto converts a InstanceTemplatePropertiesDisksInitializeParams resource to its proto representation.
func ComputeInstanceTemplatePropertiesDisksInitializeParamsToProto(o *compute.InstanceTemplatePropertiesDisksInitializeParams) *computepb.ComputeInstanceTemplatePropertiesDisksInitializeParams {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesDisksInitializeParams{
		DiskName:                    dcl.ValueOrEmptyString(o.DiskName),
		DiskSizeGb:                  dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		DiskType:                    dcl.ValueOrEmptyString(o.DiskType),
		SourceImage:                 dcl.ValueOrEmptyString(o.SourceImage),
		SourceSnapshot:              dcl.ValueOrEmptyString(o.SourceSnapshot),
		SourceSnapshotEncryptionKey: ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyToProto(o.SourceSnapshotEncryptionKey),
		Description:                 dcl.ValueOrEmptyString(o.Description),
		OnUpdateAction:              dcl.ValueOrEmptyString(o.OnUpdateAction),
		SourceImageEncryptionKey:    ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyToProto(o.SourceImageEncryptionKey),
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	for _, r := range o.ResourcePolicies {
		p.ResourcePolicies = append(p.ResourcePolicies, r)
	}
	return p
}

// InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyToProto converts a InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey resource to its proto representation.
func ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyToProto(o *compute.InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) *computepb.ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{
		RawKey:     dcl.ValueOrEmptyString(o.RawKey),
		Sha256:     dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	return p
}

// InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyToProto converts a InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey resource to its proto representation.
func ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyToProto(o *compute.InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) *computepb.ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{
		RawKey:     dcl.ValueOrEmptyString(o.RawKey),
		Sha256:     dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	return p
}

// InstanceTemplatePropertiesDisksGuestOSFeaturesToProto converts a InstanceTemplatePropertiesDisksGuestOSFeatures resource to its proto representation.
func ComputeInstanceTemplatePropertiesDisksGuestOSFeaturesToProto(o *compute.InstanceTemplatePropertiesDisksGuestOSFeatures) *computepb.ComputeInstanceTemplatePropertiesDisksGuestOSFeatures {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesDisksGuestOSFeatures{
		Type: dcl.ValueOrEmptyString(o.Type),
	}
	return p
}

// InstanceTemplatePropertiesReservationAffinityToProto converts a InstanceTemplatePropertiesReservationAffinity resource to its proto representation.
func ComputeInstanceTemplatePropertiesReservationAffinityToProto(o *compute.InstanceTemplatePropertiesReservationAffinity) *computepb.ComputeInstanceTemplatePropertiesReservationAffinity {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesReservationAffinity{
		Key: dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Value {
		p.Value = append(p.Value, r)
	}
	return p
}

// InstanceTemplatePropertiesGuestAcceleratorsToProto converts a InstanceTemplatePropertiesGuestAccelerators resource to its proto representation.
func ComputeInstanceTemplatePropertiesGuestAcceleratorsToProto(o *compute.InstanceTemplatePropertiesGuestAccelerators) *computepb.ComputeInstanceTemplatePropertiesGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesGuestAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// InstanceTemplatePropertiesNetworkInterfacesToProto converts a InstanceTemplatePropertiesNetworkInterfaces resource to its proto representation.
func ComputeInstanceTemplatePropertiesNetworkInterfacesToProto(o *compute.InstanceTemplatePropertiesNetworkInterfaces) *computepb.ComputeInstanceTemplatePropertiesNetworkInterfaces {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesNetworkInterfaces{
		Name:       dcl.ValueOrEmptyString(o.Name),
		Network:    dcl.ValueOrEmptyString(o.Network),
		NetworkIp:  dcl.ValueOrEmptyString(o.NetworkIP),
		Subnetwork: dcl.ValueOrEmptyString(o.Subnetwork),
	}
	for _, r := range o.AccessConfigs {
		p.AccessConfigs = append(p.AccessConfigs, ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsToProto(&r))
	}
	for _, r := range o.AliasIPRanges {
		p.AliasIpRanges = append(p.AliasIpRanges, ComputeInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesToProto(&r))
	}
	return p
}

// InstanceTemplatePropertiesNetworkInterfacesAccessConfigsToProto converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigs resource to its proto representation.
func ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsToProto(o *compute.InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) *computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigs{
		Name:                dcl.ValueOrEmptyString(o.Name),
		NatIp:               dcl.ValueOrEmptyString(o.NatIP),
		Type:                ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumToProto(o.Type),
		SetPublicPtr:        dcl.ValueOrEmptyBool(o.SetPublicPtr),
		PublicPtrDomainName: dcl.ValueOrEmptyString(o.PublicPtrDomainName),
		NetworkTier:         ComputeInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumToProto(o.NetworkTier),
	}
	return p
}

// InstanceTemplatePropertiesNetworkInterfacesAliasIPRangesToProto converts a InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges resource to its proto representation.
func ComputeInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesToProto(o *compute.InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) *computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{
		IpCidrRange:         dcl.ValueOrEmptyString(o.IPCidrRange),
		SubnetworkRangeName: dcl.ValueOrEmptyString(o.SubnetworkRangeName),
	}
	return p
}

// InstanceTemplatePropertiesShieldedInstanceConfigToProto converts a InstanceTemplatePropertiesShieldedInstanceConfig resource to its proto representation.
func ComputeInstanceTemplatePropertiesShieldedInstanceConfigToProto(o *compute.InstanceTemplatePropertiesShieldedInstanceConfig) *computepb.ComputeInstanceTemplatePropertiesShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableVtpm:                dcl.ValueOrEmptyBool(o.EnableVtpm),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// InstanceTemplatePropertiesSchedulingToProto converts a InstanceTemplatePropertiesScheduling resource to its proto representation.
func ComputeInstanceTemplatePropertiesSchedulingToProto(o *compute.InstanceTemplatePropertiesScheduling) *computepb.ComputeInstanceTemplatePropertiesScheduling {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesScheduling{
		AutomaticRestart:  dcl.ValueOrEmptyBool(o.AutomaticRestart),
		OnHostMaintenance: dcl.ValueOrEmptyString(o.OnHostMaintenance),
		Preemptible:       dcl.ValueOrEmptyBool(o.Preemptible),
	}
	for _, r := range o.NodeAffinities {
		p.NodeAffinities = append(p.NodeAffinities, ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesToProto(&r))
	}
	return p
}

// InstanceTemplatePropertiesSchedulingNodeAffinitiesToProto converts a InstanceTemplatePropertiesSchedulingNodeAffinities resource to its proto representation.
func ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesToProto(o *compute.InstanceTemplatePropertiesSchedulingNodeAffinities) *computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinities {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesSchedulingNodeAffinities{
		Key:      dcl.ValueOrEmptyString(o.Key),
		Operator: ComputeInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumToProto(o.Operator),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// InstanceTemplatePropertiesServiceAccountsToProto converts a InstanceTemplatePropertiesServiceAccounts resource to its proto representation.
func ComputeInstanceTemplatePropertiesServiceAccountsToProto(o *compute.InstanceTemplatePropertiesServiceAccounts) *computepb.ComputeInstanceTemplatePropertiesServiceAccounts {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceTemplatePropertiesServiceAccounts{
		Email: dcl.ValueOrEmptyString(o.Email),
	}
	for _, r := range o.Scopes {
		p.Scopes = append(p.Scopes, r)
	}
	return p
}

// InstanceTemplateToProto converts a InstanceTemplate resource to its proto representation.
func InstanceTemplateToProto(resource *compute.InstanceTemplate) *computepb.ComputeInstanceTemplate {
	p := &computepb.ComputeInstanceTemplate{
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Properties:        ComputeInstanceTemplatePropertiesToProto(resource.Properties),
		Project:           dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyInstanceTemplate handles the gRPC request by passing it to the underlying InstanceTemplate Apply() method.
func (s *InstanceTemplateServer) applyInstanceTemplate(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeInstanceTemplateRequest) (*computepb.ComputeInstanceTemplate, error) {
	p := ProtoToInstanceTemplate(request.GetResource())
	res, err := c.ApplyInstanceTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceTemplateToProto(res)
	return r, nil
}

// ApplyInstanceTemplate handles the gRPC request by passing it to the underlying InstanceTemplate Apply() method.
func (s *InstanceTemplateServer) ApplyComputeInstanceTemplate(ctx context.Context, request *computepb.ApplyComputeInstanceTemplateRequest) (*computepb.ComputeInstanceTemplate, error) {
	cl, err := createConfigInstanceTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstanceTemplate(ctx, cl, request)
}

// DeleteInstanceTemplate handles the gRPC request by passing it to the underlying InstanceTemplate Delete() method.
func (s *InstanceTemplateServer) DeleteComputeInstanceTemplate(ctx context.Context, request *computepb.DeleteComputeInstanceTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstanceTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstanceTemplate(ctx, ProtoToInstanceTemplate(request.GetResource()))

}

// ListComputeInstanceTemplate handles the gRPC request by passing it to the underlying InstanceTemplateList() method.
func (s *InstanceTemplateServer) ListComputeInstanceTemplate(ctx context.Context, request *computepb.ListComputeInstanceTemplateRequest) (*computepb.ListComputeInstanceTemplateResponse, error) {
	cl, err := createConfigInstanceTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstanceTemplate(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeInstanceTemplate
	for _, r := range resources.Items {
		rp := InstanceTemplateToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeInstanceTemplateResponse{Items: protos}, nil
}

func createConfigInstanceTemplate(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
