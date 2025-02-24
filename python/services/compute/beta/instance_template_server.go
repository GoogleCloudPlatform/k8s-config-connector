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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for InstanceTemplate.
type InstanceTemplateServer struct{}

// ProtoToInstanceTemplatePropertiesDisksInterfaceEnum converts a InstanceTemplatePropertiesDisksInterfaceEnum enum from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum(e betapb.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum) *beta.InstanceTemplatePropertiesDisksInterfaceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum_name[int32(e)]; ok {
		e := beta.InstanceTemplatePropertiesDisksInterfaceEnum(n[len("ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesDisksModeEnum converts a InstanceTemplatePropertiesDisksModeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisksModeEnum(e betapb.ComputeBetaInstanceTemplatePropertiesDisksModeEnum) *beta.InstanceTemplatePropertiesDisksModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceTemplatePropertiesDisksModeEnum_name[int32(e)]; ok {
		e := beta.InstanceTemplatePropertiesDisksModeEnum(n[len("ComputeBetaInstanceTemplatePropertiesDisksModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesDisksTypeEnum converts a InstanceTemplatePropertiesDisksTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisksTypeEnum(e betapb.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum) *beta.InstanceTemplatePropertiesDisksTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceTemplatePropertiesDisksTypeEnum(n[len("ComputeBetaInstanceTemplatePropertiesDisksTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(e betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum) *beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(n[len("ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum enum from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(e betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum) *beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum_name[int32(e)]; ok {
		e := beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(n[len("ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum converts a InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum enum from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(e betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum) *beta.InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum_name[int32(e)]; ok {
		e := beta.InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(n[len("ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceTemplateProperties converts a InstanceTemplateProperties resource from its proto representation.
func ProtoToComputeBetaInstanceTemplateProperties(p *betapb.ComputeBetaInstanceTemplateProperties) *beta.InstanceTemplateProperties {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplateProperties{
		CanIPForward:           dcl.Bool(p.CanIpForward),
		Description:            dcl.StringOrNil(p.Description),
		MachineType:            dcl.StringOrNil(p.MachineType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		ReservationAffinity:    ProtoToComputeBetaInstanceTemplatePropertiesReservationAffinity(p.GetReservationAffinity()),
		ShieldedInstanceConfig: ProtoToComputeBetaInstanceTemplatePropertiesShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		Scheduling:             ProtoToComputeBetaInstanceTemplatePropertiesScheduling(p.GetScheduling()),
	}
	for _, r := range p.GetDisks() {
		obj.Disks = append(obj.Disks, *ProtoToComputeBetaInstanceTemplatePropertiesDisks(r))
	}
	for _, r := range p.GetGuestAccelerators() {
		obj.GuestAccelerators = append(obj.GuestAccelerators, *ProtoToComputeBetaInstanceTemplatePropertiesGuestAccelerators(r))
	}
	for _, r := range p.GetNetworkInterfaces() {
		obj.NetworkInterfaces = append(obj.NetworkInterfaces, *ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfaces(r))
	}
	for _, r := range p.GetServiceAccounts() {
		obj.ServiceAccounts = append(obj.ServiceAccounts, *ProtoToComputeBetaInstanceTemplatePropertiesServiceAccounts(r))
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisks converts a InstanceTemplatePropertiesDisks resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisks(p *betapb.ComputeBetaInstanceTemplatePropertiesDisks) *beta.InstanceTemplatePropertiesDisks {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesDisks{
		AutoDelete:        dcl.Bool(p.AutoDelete),
		Boot:              dcl.Bool(p.Boot),
		DeviceName:        dcl.StringOrNil(p.DeviceName),
		DiskEncryptionKey: ProtoToComputeBetaInstanceTemplatePropertiesDisksDiskEncryptionKey(p.GetDiskEncryptionKey()),
		Index:             dcl.Int64OrNil(p.Index),
		InitializeParams:  ProtoToComputeBetaInstanceTemplatePropertiesDisksInitializeParams(p.GetInitializeParams()),
		Interface:         ProtoToComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum(p.GetInterface()),
		Mode:              ProtoToComputeBetaInstanceTemplatePropertiesDisksModeEnum(p.GetMode()),
		Source:            dcl.StringOrNil(p.Source),
		Type:              ProtoToComputeBetaInstanceTemplatePropertiesDisksTypeEnum(p.GetType()),
	}
	for _, r := range p.GetGuestOsFeatures() {
		obj.GuestOSFeatures = append(obj.GuestOSFeatures, *ProtoToComputeBetaInstanceTemplatePropertiesDisksGuestOSFeatures(r))
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksDiskEncryptionKey converts a InstanceTemplatePropertiesDisksDiskEncryptionKey resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisksDiskEncryptionKey(p *betapb.ComputeBetaInstanceTemplatePropertiesDisksDiskEncryptionKey) *beta.InstanceTemplatePropertiesDisksDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesDisksDiskEncryptionKey{
		RawKey:          dcl.StringOrNil(p.RawKey),
		RsaEncryptedKey: dcl.StringOrNil(p.RsaEncryptedKey),
		Sha256:          dcl.StringOrNil(p.Sha256),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksInitializeParams converts a InstanceTemplatePropertiesDisksInitializeParams resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisksInitializeParams(p *betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParams) *beta.InstanceTemplatePropertiesDisksInitializeParams {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesDisksInitializeParams{
		DiskName:                    dcl.StringOrNil(p.DiskName),
		DiskSizeGb:                  dcl.Int64OrNil(p.DiskSizeGb),
		DiskType:                    dcl.StringOrNil(p.DiskType),
		SourceImage:                 dcl.StringOrNil(p.SourceImage),
		SourceSnapshot:              dcl.StringOrNil(p.SourceSnapshot),
		SourceSnapshotEncryptionKey: ProtoToComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(p.GetSourceSnapshotEncryptionKey()),
		Description:                 dcl.StringOrNil(p.Description),
		OnUpdateAction:              dcl.StringOrNil(p.OnUpdateAction),
		SourceImageEncryptionKey:    ProtoToComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
	}
	for _, r := range p.GetResourcePolicies() {
		obj.ResourcePolicies = append(obj.ResourcePolicies, r)
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey converts a InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey(p *betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) *beta.InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{
		RawKey:     dcl.StringOrNil(p.RawKey),
		Sha256:     dcl.StringOrNil(p.Sha256),
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey converts a InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey(p *betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) *beta.InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{
		RawKey:     dcl.StringOrNil(p.RawKey),
		Sha256:     dcl.StringOrNil(p.Sha256),
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesDisksGuestOSFeatures converts a InstanceTemplatePropertiesDisksGuestOSFeatures resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesDisksGuestOSFeatures(p *betapb.ComputeBetaInstanceTemplatePropertiesDisksGuestOSFeatures) *beta.InstanceTemplatePropertiesDisksGuestOSFeatures {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesDisksGuestOSFeatures{
		Type: dcl.StringOrNil(p.Type),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesReservationAffinity converts a InstanceTemplatePropertiesReservationAffinity resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesReservationAffinity(p *betapb.ComputeBetaInstanceTemplatePropertiesReservationAffinity) *beta.InstanceTemplatePropertiesReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesReservationAffinity{
		Key: dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValue() {
		obj.Value = append(obj.Value, r)
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesGuestAccelerators converts a InstanceTemplatePropertiesGuestAccelerators resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesGuestAccelerators(p *betapb.ComputeBetaInstanceTemplatePropertiesGuestAccelerators) *beta.InstanceTemplatePropertiesGuestAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesGuestAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesNetworkInterfaces converts a InstanceTemplatePropertiesNetworkInterfaces resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfaces(p *betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfaces) *beta.InstanceTemplatePropertiesNetworkInterfaces {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesNetworkInterfaces{
		Name:       dcl.StringOrNil(p.Name),
		Network:    dcl.StringOrNil(p.Network),
		NetworkIP:  dcl.StringOrNil(p.NetworkIp),
		Subnetwork: dcl.StringOrNil(p.Subnetwork),
	}
	for _, r := range p.GetAccessConfigs() {
		obj.AccessConfigs = append(obj.AccessConfigs, *ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(r))
	}
	for _, r := range p.GetAliasIpRanges() {
		obj.AliasIPRanges = append(obj.AliasIPRanges, *ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(r))
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesNetworkInterfacesAccessConfigs converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigs resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigs(p *betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigs) *beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigs{
		Name:                dcl.StringOrNil(p.Name),
		NatIP:               dcl.StringOrNil(p.NatIp),
		Type:                ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(p.GetType()),
		SetPublicPtr:        dcl.Bool(p.SetPublicPtr),
		PublicPtrDomainName: dcl.StringOrNil(p.PublicPtrDomainName),
		NetworkTier:         ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(p.GetNetworkTier()),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges converts a InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges(p *betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) *beta.InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{
		IPCidrRange:         dcl.StringOrNil(p.IpCidrRange),
		SubnetworkRangeName: dcl.StringOrNil(p.SubnetworkRangeName),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesShieldedInstanceConfig converts a InstanceTemplatePropertiesShieldedInstanceConfig resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesShieldedInstanceConfig(p *betapb.ComputeBetaInstanceTemplatePropertiesShieldedInstanceConfig) *beta.InstanceTemplatePropertiesShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableVtpm:                dcl.Bool(p.EnableVtpm),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesScheduling converts a InstanceTemplatePropertiesScheduling resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesScheduling(p *betapb.ComputeBetaInstanceTemplatePropertiesScheduling) *beta.InstanceTemplatePropertiesScheduling {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesScheduling{
		AutomaticRestart:  dcl.Bool(p.AutomaticRestart),
		OnHostMaintenance: dcl.StringOrNil(p.OnHostMaintenance),
		Preemptible:       dcl.Bool(p.Preemptible),
	}
	for _, r := range p.GetNodeAffinities() {
		obj.NodeAffinities = append(obj.NodeAffinities, *ProtoToComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinities(r))
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesSchedulingNodeAffinities converts a InstanceTemplatePropertiesSchedulingNodeAffinities resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinities(p *betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinities) *beta.InstanceTemplatePropertiesSchedulingNodeAffinities {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesSchedulingNodeAffinities{
		Key:      dcl.StringOrNil(p.Key),
		Operator: ProtoToComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(p.GetOperator()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToInstanceTemplatePropertiesServiceAccounts converts a InstanceTemplatePropertiesServiceAccounts resource from its proto representation.
func ProtoToComputeBetaInstanceTemplatePropertiesServiceAccounts(p *betapb.ComputeBetaInstanceTemplatePropertiesServiceAccounts) *beta.InstanceTemplatePropertiesServiceAccounts {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceTemplatePropertiesServiceAccounts{
		Email: dcl.StringOrNil(p.Email),
	}
	for _, r := range p.GetScopes() {
		obj.Scopes = append(obj.Scopes, r)
	}
	return obj
}

// ProtoToInstanceTemplate converts a InstanceTemplate resource from its proto representation.
func ProtoToInstanceTemplate(p *betapb.ComputeBetaInstanceTemplate) *beta.InstanceTemplate {
	obj := &beta.InstanceTemplate{
		CreationTimestamp: dcl.StringOrNil(p.GetCreationTimestamp()),
		Description:       dcl.StringOrNil(p.Description),
		Id:                dcl.Int64OrNil(p.Id),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Name:              dcl.StringOrNil(p.Name),
		Properties:        ProtoToComputeBetaInstanceTemplateProperties(p.GetProperties()),
		Project:           dcl.StringOrNil(p.Project),
	}
	return obj
}

// InstanceTemplatePropertiesDisksInterfaceEnumToProto converts a InstanceTemplatePropertiesDisksInterfaceEnum enum to its proto representation.
func ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnumToProto(e *beta.InstanceTemplatePropertiesDisksInterfaceEnum) betapb.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum_value["InstanceTemplatePropertiesDisksInterfaceEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum(v)
	}
	return betapb.ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnum(0)
}

// InstanceTemplatePropertiesDisksModeEnumToProto converts a InstanceTemplatePropertiesDisksModeEnum enum to its proto representation.
func ComputeBetaInstanceTemplatePropertiesDisksModeEnumToProto(e *beta.InstanceTemplatePropertiesDisksModeEnum) betapb.ComputeBetaInstanceTemplatePropertiesDisksModeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceTemplatePropertiesDisksModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceTemplatePropertiesDisksModeEnum_value["InstanceTemplatePropertiesDisksModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceTemplatePropertiesDisksModeEnum(v)
	}
	return betapb.ComputeBetaInstanceTemplatePropertiesDisksModeEnum(0)
}

// InstanceTemplatePropertiesDisksTypeEnumToProto converts a InstanceTemplatePropertiesDisksTypeEnum enum to its proto representation.
func ComputeBetaInstanceTemplatePropertiesDisksTypeEnumToProto(e *beta.InstanceTemplatePropertiesDisksTypeEnum) betapb.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum_value["InstanceTemplatePropertiesDisksTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceTemplatePropertiesDisksTypeEnum(0)
}

// InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumToProto converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum enum to its proto representation.
func ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumToProto(e *beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum) betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum_value["InstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnum(0)
}

// InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumToProto converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum enum to its proto representation.
func ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumToProto(e *beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum) betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum_value["InstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(v)
	}
	return betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnum(0)
}

// InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumToProto converts a InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum enum to its proto representation.
func ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumToProto(e *beta.InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum) betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum_value["InstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(v)
	}
	return betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnum(0)
}

// InstanceTemplatePropertiesToProto converts a InstanceTemplateProperties resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesToProto(o *beta.InstanceTemplateProperties) *betapb.ComputeBetaInstanceTemplateProperties {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplateProperties{
		CanIpForward:           dcl.ValueOrEmptyBool(o.CanIPForward),
		Description:            dcl.ValueOrEmptyString(o.Description),
		MachineType:            dcl.ValueOrEmptyString(o.MachineType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		ReservationAffinity:    ComputeBetaInstanceTemplatePropertiesReservationAffinityToProto(o.ReservationAffinity),
		ShieldedInstanceConfig: ComputeBetaInstanceTemplatePropertiesShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
		Scheduling:             ComputeBetaInstanceTemplatePropertiesSchedulingToProto(o.Scheduling),
	}
	for _, r := range o.Disks {
		p.Disks = append(p.Disks, ComputeBetaInstanceTemplatePropertiesDisksToProto(&r))
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
		p.GuestAccelerators = append(p.GuestAccelerators, ComputeBetaInstanceTemplatePropertiesGuestAcceleratorsToProto(&r))
	}
	for _, r := range o.NetworkInterfaces {
		p.NetworkInterfaces = append(p.NetworkInterfaces, ComputeBetaInstanceTemplatePropertiesNetworkInterfacesToProto(&r))
	}
	for _, r := range o.ServiceAccounts {
		p.ServiceAccounts = append(p.ServiceAccounts, ComputeBetaInstanceTemplatePropertiesServiceAccountsToProto(&r))
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	return p
}

// InstanceTemplatePropertiesDisksToProto converts a InstanceTemplatePropertiesDisks resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesDisksToProto(o *beta.InstanceTemplatePropertiesDisks) *betapb.ComputeBetaInstanceTemplatePropertiesDisks {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesDisks{
		AutoDelete:        dcl.ValueOrEmptyBool(o.AutoDelete),
		Boot:              dcl.ValueOrEmptyBool(o.Boot),
		DeviceName:        dcl.ValueOrEmptyString(o.DeviceName),
		DiskEncryptionKey: ComputeBetaInstanceTemplatePropertiesDisksDiskEncryptionKeyToProto(o.DiskEncryptionKey),
		Index:             dcl.ValueOrEmptyInt64(o.Index),
		InitializeParams:  ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsToProto(o.InitializeParams),
		Interface:         ComputeBetaInstanceTemplatePropertiesDisksInterfaceEnumToProto(o.Interface),
		Mode:              ComputeBetaInstanceTemplatePropertiesDisksModeEnumToProto(o.Mode),
		Source:            dcl.ValueOrEmptyString(o.Source),
		Type:              ComputeBetaInstanceTemplatePropertiesDisksTypeEnumToProto(o.Type),
	}
	for _, r := range o.GuestOSFeatures {
		p.GuestOsFeatures = append(p.GuestOsFeatures, ComputeBetaInstanceTemplatePropertiesDisksGuestOSFeaturesToProto(&r))
	}
	return p
}

// InstanceTemplatePropertiesDisksDiskEncryptionKeyToProto converts a InstanceTemplatePropertiesDisksDiskEncryptionKey resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesDisksDiskEncryptionKeyToProto(o *beta.InstanceTemplatePropertiesDisksDiskEncryptionKey) *betapb.ComputeBetaInstanceTemplatePropertiesDisksDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesDisksDiskEncryptionKey{
		RawKey:          dcl.ValueOrEmptyString(o.RawKey),
		RsaEncryptedKey: dcl.ValueOrEmptyString(o.RsaEncryptedKey),
		Sha256:          dcl.ValueOrEmptyString(o.Sha256),
	}
	return p
}

// InstanceTemplatePropertiesDisksInitializeParamsToProto converts a InstanceTemplatePropertiesDisksInitializeParams resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsToProto(o *beta.InstanceTemplatePropertiesDisksInitializeParams) *betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParams {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParams{
		DiskName:                    dcl.ValueOrEmptyString(o.DiskName),
		DiskSizeGb:                  dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		DiskType:                    dcl.ValueOrEmptyString(o.DiskType),
		SourceImage:                 dcl.ValueOrEmptyString(o.SourceImage),
		SourceSnapshot:              dcl.ValueOrEmptyString(o.SourceSnapshot),
		SourceSnapshotEncryptionKey: ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyToProto(o.SourceSnapshotEncryptionKey),
		Description:                 dcl.ValueOrEmptyString(o.Description),
		OnUpdateAction:              dcl.ValueOrEmptyString(o.OnUpdateAction),
		SourceImageEncryptionKey:    ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyToProto(o.SourceImageEncryptionKey),
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
func ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKeyToProto(o *beta.InstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey) *betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceSnapshotEncryptionKey{
		RawKey:     dcl.ValueOrEmptyString(o.RawKey),
		Sha256:     dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	return p
}

// InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyToProto converts a InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKeyToProto(o *beta.InstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey) *betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesDisksInitializeParamsSourceImageEncryptionKey{
		RawKey:     dcl.ValueOrEmptyString(o.RawKey),
		Sha256:     dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	return p
}

// InstanceTemplatePropertiesDisksGuestOSFeaturesToProto converts a InstanceTemplatePropertiesDisksGuestOSFeatures resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesDisksGuestOSFeaturesToProto(o *beta.InstanceTemplatePropertiesDisksGuestOSFeatures) *betapb.ComputeBetaInstanceTemplatePropertiesDisksGuestOSFeatures {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesDisksGuestOSFeatures{
		Type: dcl.ValueOrEmptyString(o.Type),
	}
	return p
}

// InstanceTemplatePropertiesReservationAffinityToProto converts a InstanceTemplatePropertiesReservationAffinity resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesReservationAffinityToProto(o *beta.InstanceTemplatePropertiesReservationAffinity) *betapb.ComputeBetaInstanceTemplatePropertiesReservationAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesReservationAffinity{
		Key: dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Value {
		p.Value = append(p.Value, r)
	}
	return p
}

// InstanceTemplatePropertiesGuestAcceleratorsToProto converts a InstanceTemplatePropertiesGuestAccelerators resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesGuestAcceleratorsToProto(o *beta.InstanceTemplatePropertiesGuestAccelerators) *betapb.ComputeBetaInstanceTemplatePropertiesGuestAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesGuestAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// InstanceTemplatePropertiesNetworkInterfacesToProto converts a InstanceTemplatePropertiesNetworkInterfaces resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesNetworkInterfacesToProto(o *beta.InstanceTemplatePropertiesNetworkInterfaces) *betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfaces {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfaces{
		Name:       dcl.ValueOrEmptyString(o.Name),
		Network:    dcl.ValueOrEmptyString(o.Network),
		NetworkIp:  dcl.ValueOrEmptyString(o.NetworkIP),
		Subnetwork: dcl.ValueOrEmptyString(o.Subnetwork),
	}
	for _, r := range o.AccessConfigs {
		p.AccessConfigs = append(p.AccessConfigs, ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsToProto(&r))
	}
	for _, r := range o.AliasIPRanges {
		p.AliasIpRanges = append(p.AliasIpRanges, ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesToProto(&r))
	}
	return p
}

// InstanceTemplatePropertiesNetworkInterfacesAccessConfigsToProto converts a InstanceTemplatePropertiesNetworkInterfacesAccessConfigs resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsToProto(o *beta.InstanceTemplatePropertiesNetworkInterfacesAccessConfigs) *betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigs{
		Name:                dcl.ValueOrEmptyString(o.Name),
		NatIp:               dcl.ValueOrEmptyString(o.NatIP),
		Type:                ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsTypeEnumToProto(o.Type),
		SetPublicPtr:        dcl.ValueOrEmptyBool(o.SetPublicPtr),
		PublicPtrDomainName: dcl.ValueOrEmptyString(o.PublicPtrDomainName),
		NetworkTier:         ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAccessConfigsNetworkTierEnumToProto(o.NetworkTier),
	}
	return p
}

// InstanceTemplatePropertiesNetworkInterfacesAliasIPRangesToProto converts a InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAliasIPRangesToProto(o *beta.InstanceTemplatePropertiesNetworkInterfacesAliasIPRanges) *betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesNetworkInterfacesAliasIPRanges{
		IpCidrRange:         dcl.ValueOrEmptyString(o.IPCidrRange),
		SubnetworkRangeName: dcl.ValueOrEmptyString(o.SubnetworkRangeName),
	}
	return p
}

// InstanceTemplatePropertiesShieldedInstanceConfigToProto converts a InstanceTemplatePropertiesShieldedInstanceConfig resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesShieldedInstanceConfigToProto(o *beta.InstanceTemplatePropertiesShieldedInstanceConfig) *betapb.ComputeBetaInstanceTemplatePropertiesShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableVtpm:                dcl.ValueOrEmptyBool(o.EnableVtpm),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// InstanceTemplatePropertiesSchedulingToProto converts a InstanceTemplatePropertiesScheduling resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesSchedulingToProto(o *beta.InstanceTemplatePropertiesScheduling) *betapb.ComputeBetaInstanceTemplatePropertiesScheduling {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesScheduling{
		AutomaticRestart:  dcl.ValueOrEmptyBool(o.AutomaticRestart),
		OnHostMaintenance: dcl.ValueOrEmptyString(o.OnHostMaintenance),
		Preemptible:       dcl.ValueOrEmptyBool(o.Preemptible),
	}
	for _, r := range o.NodeAffinities {
		p.NodeAffinities = append(p.NodeAffinities, ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesToProto(&r))
	}
	return p
}

// InstanceTemplatePropertiesSchedulingNodeAffinitiesToProto converts a InstanceTemplatePropertiesSchedulingNodeAffinities resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesToProto(o *beta.InstanceTemplatePropertiesSchedulingNodeAffinities) *betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinities {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinities{
		Key:      dcl.ValueOrEmptyString(o.Key),
		Operator: ComputeBetaInstanceTemplatePropertiesSchedulingNodeAffinitiesOperatorEnumToProto(o.Operator),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// InstanceTemplatePropertiesServiceAccountsToProto converts a InstanceTemplatePropertiesServiceAccounts resource to its proto representation.
func ComputeBetaInstanceTemplatePropertiesServiceAccountsToProto(o *beta.InstanceTemplatePropertiesServiceAccounts) *betapb.ComputeBetaInstanceTemplatePropertiesServiceAccounts {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceTemplatePropertiesServiceAccounts{
		Email: dcl.ValueOrEmptyString(o.Email),
	}
	for _, r := range o.Scopes {
		p.Scopes = append(p.Scopes, r)
	}
	return p
}

// InstanceTemplateToProto converts a InstanceTemplate resource to its proto representation.
func InstanceTemplateToProto(resource *beta.InstanceTemplate) *betapb.ComputeBetaInstanceTemplate {
	p := &betapb.ComputeBetaInstanceTemplate{
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Properties:        ComputeBetaInstanceTemplatePropertiesToProto(resource.Properties),
		Project:           dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyInstanceTemplate handles the gRPC request by passing it to the underlying InstanceTemplate Apply() method.
func (s *InstanceTemplateServer) applyInstanceTemplate(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaInstanceTemplateRequest) (*betapb.ComputeBetaInstanceTemplate, error) {
	p := ProtoToInstanceTemplate(request.GetResource())
	res, err := c.ApplyInstanceTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceTemplateToProto(res)
	return r, nil
}

// ApplyInstanceTemplate handles the gRPC request by passing it to the underlying InstanceTemplate Apply() method.
func (s *InstanceTemplateServer) ApplyComputeBetaInstanceTemplate(ctx context.Context, request *betapb.ApplyComputeBetaInstanceTemplateRequest) (*betapb.ComputeBetaInstanceTemplate, error) {
	cl, err := createConfigInstanceTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstanceTemplate(ctx, cl, request)
}

// DeleteInstanceTemplate handles the gRPC request by passing it to the underlying InstanceTemplate Delete() method.
func (s *InstanceTemplateServer) DeleteComputeBetaInstanceTemplate(ctx context.Context, request *betapb.DeleteComputeBetaInstanceTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstanceTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstanceTemplate(ctx, ProtoToInstanceTemplate(request.GetResource()))

}

// ListComputeBetaInstanceTemplate handles the gRPC request by passing it to the underlying InstanceTemplateList() method.
func (s *InstanceTemplateServer) ListComputeBetaInstanceTemplate(ctx context.Context, request *betapb.ListComputeBetaInstanceTemplateRequest) (*betapb.ListComputeBetaInstanceTemplateResponse, error) {
	cl, err := createConfigInstanceTemplate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstanceTemplate(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaInstanceTemplate
	for _, r := range resources.Items {
		rp := InstanceTemplateToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaInstanceTemplateResponse{Items: protos}, nil
}

func createConfigInstanceTemplate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
