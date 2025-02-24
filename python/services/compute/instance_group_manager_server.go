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

// InstanceGroupManagerServer implements the gRPC interface for InstanceGroupManager.
type InstanceGroupManagerServer struct{}

// ProtoToInstanceGroupManagerDistributionPolicyTargetShapeEnum converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(e computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum) *compute.InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerDistributionPolicyTargetShapeEnum(n[len("ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyTypeEnum converts a InstanceGroupManagerUpdatePolicyTypeEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyTypeEnum(e computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum) *compute.InstanceGroupManagerUpdatePolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerUpdatePolicyTypeEnum(n[len("ComputeInstanceGroupManagerUpdatePolicyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(e computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) *compute.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(n[len("ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMinimalActionEnum converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(e computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum) *compute.InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerUpdatePolicyMinimalActionEnum(n[len("ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyReplacementMethodEnum converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(e computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum) *compute.InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerUpdatePolicyReplacementMethodEnum(n[len("ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum from its proto representation.
func ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(e computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) *compute.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_name[int32(e)]; ok {
		e := compute.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(n[len("ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerDistributionPolicy converts a InstanceGroupManagerDistributionPolicy object from its proto representation.
func ProtoToComputeInstanceGroupManagerDistributionPolicy(p *computepb.ComputeInstanceGroupManagerDistributionPolicy) *compute.InstanceGroupManagerDistributionPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerDistributionPolicy{
		TargetShape: ProtoToComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(p.GetTargetShape()),
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, *ProtoToComputeInstanceGroupManagerDistributionPolicyZones(r))
	}
	return obj
}

// ProtoToInstanceGroupManagerDistributionPolicyZones converts a InstanceGroupManagerDistributionPolicyZones object from its proto representation.
func ProtoToComputeInstanceGroupManagerDistributionPolicyZones(p *computepb.ComputeInstanceGroupManagerDistributionPolicyZones) *compute.InstanceGroupManagerDistributionPolicyZones {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersions converts a InstanceGroupManagerVersions object from its proto representation.
func ProtoToComputeInstanceGroupManagerVersions(p *computepb.ComputeInstanceGroupManagerVersions) *compute.InstanceGroupManagerVersions {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerVersions{
		Name:             dcl.StringOrNil(p.GetName()),
		InstanceTemplate: dcl.StringOrNil(p.GetInstanceTemplate()),
		TargetSize:       ProtoToComputeInstanceGroupManagerVersionsTargetSize(p.GetTargetSize()),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersionsTargetSize converts a InstanceGroupManagerVersionsTargetSize object from its proto representation.
func ProtoToComputeInstanceGroupManagerVersionsTargetSize(p *computepb.ComputeInstanceGroupManagerVersionsTargetSize) *compute.InstanceGroupManagerVersionsTargetSize {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerVersionsTargetSize{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerCurrentActions converts a InstanceGroupManagerCurrentActions object from its proto representation.
func ProtoToComputeInstanceGroupManagerCurrentActions(p *computepb.ComputeInstanceGroupManagerCurrentActions) *compute.InstanceGroupManagerCurrentActions {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerCurrentActions{
		None:                   dcl.Int64OrNil(p.GetNone()),
		Creating:               dcl.Int64OrNil(p.GetCreating()),
		CreatingWithoutRetries: dcl.Int64OrNil(p.GetCreatingWithoutRetries()),
		Verifying:              dcl.Int64OrNil(p.GetVerifying()),
		Recreating:             dcl.Int64OrNil(p.GetRecreating()),
		Deleting:               dcl.Int64OrNil(p.GetDeleting()),
		Abandoning:             dcl.Int64OrNil(p.GetAbandoning()),
		Restarting:             dcl.Int64OrNil(p.GetRestarting()),
		Refreshing:             dcl.Int64OrNil(p.GetRefreshing()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatus converts a InstanceGroupManagerStatus object from its proto representation.
func ProtoToComputeInstanceGroupManagerStatus(p *computepb.ComputeInstanceGroupManagerStatus) *compute.InstanceGroupManagerStatus {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatus{
		IsStable:      dcl.Bool(p.GetIsStable()),
		VersionTarget: ProtoToComputeInstanceGroupManagerStatusVersionTarget(p.GetVersionTarget()),
		Stateful:      ProtoToComputeInstanceGroupManagerStatusStateful(p.GetStateful()),
		Autoscaler:    dcl.StringOrNil(p.GetAutoscaler()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusVersionTarget converts a InstanceGroupManagerStatusVersionTarget object from its proto representation.
func ProtoToComputeInstanceGroupManagerStatusVersionTarget(p *computepb.ComputeInstanceGroupManagerStatusVersionTarget) *compute.InstanceGroupManagerStatusVersionTarget {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.Bool(p.GetIsReached()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStateful converts a InstanceGroupManagerStatusStateful object from its proto representation.
func ProtoToComputeInstanceGroupManagerStatusStateful(p *computepb.ComputeInstanceGroupManagerStatusStateful) *compute.InstanceGroupManagerStatusStateful {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.Bool(p.GetHasStatefulConfig()),
		PerInstanceConfigs: ProtoToComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs(p.GetPerInstanceConfigs()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStatefulPerInstanceConfigs converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs object from its proto representation.
func ProtoToComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs(p *computepb.ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs) *compute.InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.Bool(p.GetAllEffective()),
	}
	return obj
}

// ProtoToInstanceGroupManagerAutoHealingPolicies converts a InstanceGroupManagerAutoHealingPolicies object from its proto representation.
func ProtoToComputeInstanceGroupManagerAutoHealingPolicies(p *computepb.ComputeInstanceGroupManagerAutoHealingPolicies) *compute.InstanceGroupManagerAutoHealingPolicies {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerAutoHealingPolicies{
		HealthCheck:     dcl.StringOrNil(p.GetHealthCheck()),
		InitialDelaySec: dcl.Int64OrNil(p.GetInitialDelaySec()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicy converts a InstanceGroupManagerUpdatePolicy object from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicy(p *computepb.ComputeInstanceGroupManagerUpdatePolicy) *compute.InstanceGroupManagerUpdatePolicy {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerUpdatePolicy{
		Type:                       ProtoToComputeInstanceGroupManagerUpdatePolicyTypeEnum(p.GetType()),
		InstanceRedistributionType: ProtoToComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(p.GetInstanceRedistributionType()),
		MinimalAction:              ProtoToComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(p.GetMinimalAction()),
		MaxSurge:                   ProtoToComputeInstanceGroupManagerUpdatePolicyMaxSurge(p.GetMaxSurge()),
		MaxUnavailable:             ProtoToComputeInstanceGroupManagerUpdatePolicyMaxUnavailable(p.GetMaxUnavailable()),
		ReplacementMethod:          ProtoToComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(p.GetReplacementMethod()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxSurge converts a InstanceGroupManagerUpdatePolicyMaxSurge object from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyMaxSurge(p *computepb.ComputeInstanceGroupManagerUpdatePolicyMaxSurge) *compute.InstanceGroupManagerUpdatePolicyMaxSurge {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerUpdatePolicyMaxSurge{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxUnavailable converts a InstanceGroupManagerUpdatePolicyMaxUnavailable object from its proto representation.
func ProtoToComputeInstanceGroupManagerUpdatePolicyMaxUnavailable(p *computepb.ComputeInstanceGroupManagerUpdatePolicyMaxUnavailable) *compute.InstanceGroupManagerUpdatePolicyMaxUnavailable {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerUpdatePolicyMaxUnavailable{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerNamedPorts converts a InstanceGroupManagerNamedPorts object from its proto representation.
func ProtoToComputeInstanceGroupManagerNamedPorts(p *computepb.ComputeInstanceGroupManagerNamedPorts) *compute.InstanceGroupManagerNamedPorts {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerNamedPorts{
		Name: dcl.StringOrNil(p.GetName()),
		Port: dcl.Int64OrNil(p.GetPort()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicy converts a InstanceGroupManagerStatefulPolicy object from its proto representation.
func ProtoToComputeInstanceGroupManagerStatefulPolicy(p *computepb.ComputeInstanceGroupManagerStatefulPolicy) *compute.InstanceGroupManagerStatefulPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatefulPolicy{
		PreservedState: ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedState(p.GetPreservedState()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedState converts a InstanceGroupManagerStatefulPolicyPreservedState object from its proto representation.
func ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedState(p *computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedState) *compute.InstanceGroupManagerStatefulPolicyPreservedState {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatefulPolicyPreservedState{}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisks converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks object from its proto representation.
func ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks(p *computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks) *compute.InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if p == nil {
		return nil
	}
	obj := &compute.InstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ProtoToComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManager converts a InstanceGroupManager resource from its proto representation.
func ProtoToInstanceGroupManager(p *computepb.ComputeInstanceGroupManager) *compute.InstanceGroupManager {
	obj := &compute.InstanceGroupManager{
		Id:                 dcl.Int64OrNil(p.GetId()),
		CreationTimestamp:  dcl.StringOrNil(p.GetCreationTimestamp()),
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		Zone:               dcl.StringOrNil(p.GetZone()),
		Region:             dcl.StringOrNil(p.GetRegion()),
		DistributionPolicy: ProtoToComputeInstanceGroupManagerDistributionPolicy(p.GetDistributionPolicy()),
		InstanceTemplate:   dcl.StringOrNil(p.GetInstanceTemplate()),
		InstanceGroup:      dcl.StringOrNil(p.GetInstanceGroup()),
		BaseInstanceName:   dcl.StringOrNil(p.GetBaseInstanceName()),
		Fingerprint:        dcl.StringOrNil(p.GetFingerprint()),
		CurrentActions:     ProtoToComputeInstanceGroupManagerCurrentActions(p.GetCurrentActions()),
		Status:             ProtoToComputeInstanceGroupManagerStatus(p.GetStatus()),
		TargetSize:         dcl.Int64OrNil(p.GetTargetSize()),
		SelfLink:           dcl.StringOrNil(p.GetSelfLink()),
		UpdatePolicy:       ProtoToComputeInstanceGroupManagerUpdatePolicy(p.GetUpdatePolicy()),
		StatefulPolicy:     ProtoToComputeInstanceGroupManagerStatefulPolicy(p.GetStatefulPolicy()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetVersions() {
		obj.Versions = append(obj.Versions, *ProtoToComputeInstanceGroupManagerVersions(r))
	}
	for _, r := range p.GetTargetPools() {
		obj.TargetPools = append(obj.TargetPools, r)
	}
	for _, r := range p.GetAutoHealingPolicies() {
		obj.AutoHealingPolicies = append(obj.AutoHealingPolicies, *ProtoToComputeInstanceGroupManagerAutoHealingPolicies(r))
	}
	for _, r := range p.GetNamedPorts() {
		obj.NamedPorts = append(obj.NamedPorts, *ProtoToComputeInstanceGroupManagerNamedPorts(r))
	}
	return obj
}

// InstanceGroupManagerDistributionPolicyTargetShapeEnumToProto converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum to its proto representation.
func ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(e *compute.InstanceGroupManagerDistributionPolicyTargetShapeEnum) computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum_value["InstanceGroupManagerDistributionPolicyTargetShapeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
}

// InstanceGroupManagerUpdatePolicyTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyTypeEnum enum to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyTypeEnumToProto(e *compute.InstanceGroupManagerUpdatePolicyTypeEnum) computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum_value["InstanceGroupManagerUpdatePolicyTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerUpdatePolicyTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(e *compute.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_value["InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyMinimalActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(e *compute.InstanceGroupManagerUpdatePolicyMinimalActionEnum) computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum_value["InstanceGroupManagerUpdatePolicyMinimalActionEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
}

// InstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(e *compute.InstanceGroupManagerUpdatePolicyReplacementMethodEnum) computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum_value["InstanceGroupManagerUpdatePolicyReplacementMethodEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum to its proto representation.
func ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(e *compute.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == nil {
		return computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
	}
	if v, ok := computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"+string(*e)]; ok {
		return computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(v)
	}
	return computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
}

// InstanceGroupManagerDistributionPolicyToProto converts a InstanceGroupManagerDistributionPolicy object to its proto representation.
func ComputeInstanceGroupManagerDistributionPolicyToProto(o *compute.InstanceGroupManagerDistributionPolicy) *computepb.ComputeInstanceGroupManagerDistributionPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerDistributionPolicy{}
	p.SetTargetShape(ComputeInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(o.TargetShape))
	sZones := make([]*computepb.ComputeInstanceGroupManagerDistributionPolicyZones, len(o.Zones))
	for i, r := range o.Zones {
		sZones[i] = ComputeInstanceGroupManagerDistributionPolicyZonesToProto(&r)
	}
	p.SetZones(sZones)
	return p
}

// InstanceGroupManagerDistributionPolicyZonesToProto converts a InstanceGroupManagerDistributionPolicyZones object to its proto representation.
func ComputeInstanceGroupManagerDistributionPolicyZonesToProto(o *compute.InstanceGroupManagerDistributionPolicyZones) *computepb.ComputeInstanceGroupManagerDistributionPolicyZones {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerDistributionPolicyZones{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	return p
}

// InstanceGroupManagerVersionsToProto converts a InstanceGroupManagerVersions object to its proto representation.
func ComputeInstanceGroupManagerVersionsToProto(o *compute.InstanceGroupManagerVersions) *computepb.ComputeInstanceGroupManagerVersions {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerVersions{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetInstanceTemplate(dcl.ValueOrEmptyString(o.InstanceTemplate))
	p.SetTargetSize(ComputeInstanceGroupManagerVersionsTargetSizeToProto(o.TargetSize))
	return p
}

// InstanceGroupManagerVersionsTargetSizeToProto converts a InstanceGroupManagerVersionsTargetSize object to its proto representation.
func ComputeInstanceGroupManagerVersionsTargetSizeToProto(o *compute.InstanceGroupManagerVersionsTargetSize) *computepb.ComputeInstanceGroupManagerVersionsTargetSize {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerVersionsTargetSize{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerCurrentActionsToProto converts a InstanceGroupManagerCurrentActions object to its proto representation.
func ComputeInstanceGroupManagerCurrentActionsToProto(o *compute.InstanceGroupManagerCurrentActions) *computepb.ComputeInstanceGroupManagerCurrentActions {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerCurrentActions{}
	p.SetNone(dcl.ValueOrEmptyInt64(o.None))
	p.SetCreating(dcl.ValueOrEmptyInt64(o.Creating))
	p.SetCreatingWithoutRetries(dcl.ValueOrEmptyInt64(o.CreatingWithoutRetries))
	p.SetVerifying(dcl.ValueOrEmptyInt64(o.Verifying))
	p.SetRecreating(dcl.ValueOrEmptyInt64(o.Recreating))
	p.SetDeleting(dcl.ValueOrEmptyInt64(o.Deleting))
	p.SetAbandoning(dcl.ValueOrEmptyInt64(o.Abandoning))
	p.SetRestarting(dcl.ValueOrEmptyInt64(o.Restarting))
	p.SetRefreshing(dcl.ValueOrEmptyInt64(o.Refreshing))
	return p
}

// InstanceGroupManagerStatusToProto converts a InstanceGroupManagerStatus object to its proto representation.
func ComputeInstanceGroupManagerStatusToProto(o *compute.InstanceGroupManagerStatus) *computepb.ComputeInstanceGroupManagerStatus {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatus{}
	p.SetIsStable(dcl.ValueOrEmptyBool(o.IsStable))
	p.SetVersionTarget(ComputeInstanceGroupManagerStatusVersionTargetToProto(o.VersionTarget))
	p.SetStateful(ComputeInstanceGroupManagerStatusStatefulToProto(o.Stateful))
	p.SetAutoscaler(dcl.ValueOrEmptyString(o.Autoscaler))
	return p
}

// InstanceGroupManagerStatusVersionTargetToProto converts a InstanceGroupManagerStatusVersionTarget object to its proto representation.
func ComputeInstanceGroupManagerStatusVersionTargetToProto(o *compute.InstanceGroupManagerStatusVersionTarget) *computepb.ComputeInstanceGroupManagerStatusVersionTarget {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatusVersionTarget{}
	p.SetIsReached(dcl.ValueOrEmptyBool(o.IsReached))
	return p
}

// InstanceGroupManagerStatusStatefulToProto converts a InstanceGroupManagerStatusStateful object to its proto representation.
func ComputeInstanceGroupManagerStatusStatefulToProto(o *compute.InstanceGroupManagerStatusStateful) *computepb.ComputeInstanceGroupManagerStatusStateful {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatusStateful{}
	p.SetHasStatefulConfig(dcl.ValueOrEmptyBool(o.HasStatefulConfig))
	p.SetPerInstanceConfigs(ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o.PerInstanceConfigs))
	return p
}

// InstanceGroupManagerStatusStatefulPerInstanceConfigsToProto converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs object to its proto representation.
func ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o *compute.InstanceGroupManagerStatusStatefulPerInstanceConfigs) *computepb.ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatusStatefulPerInstanceConfigs{}
	p.SetAllEffective(dcl.ValueOrEmptyBool(o.AllEffective))
	return p
}

// InstanceGroupManagerAutoHealingPoliciesToProto converts a InstanceGroupManagerAutoHealingPolicies object to its proto representation.
func ComputeInstanceGroupManagerAutoHealingPoliciesToProto(o *compute.InstanceGroupManagerAutoHealingPolicies) *computepb.ComputeInstanceGroupManagerAutoHealingPolicies {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerAutoHealingPolicies{}
	p.SetHealthCheck(dcl.ValueOrEmptyString(o.HealthCheck))
	p.SetInitialDelaySec(dcl.ValueOrEmptyInt64(o.InitialDelaySec))
	return p
}

// InstanceGroupManagerUpdatePolicyToProto converts a InstanceGroupManagerUpdatePolicy object to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyToProto(o *compute.InstanceGroupManagerUpdatePolicy) *computepb.ComputeInstanceGroupManagerUpdatePolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerUpdatePolicy{}
	p.SetType(ComputeInstanceGroupManagerUpdatePolicyTypeEnumToProto(o.Type))
	p.SetInstanceRedistributionType(ComputeInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(o.InstanceRedistributionType))
	p.SetMinimalAction(ComputeInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(o.MinimalAction))
	p.SetMaxSurge(ComputeInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o.MaxSurge))
	p.SetMaxUnavailable(ComputeInstanceGroupManagerUpdatePolicyMaxUnavailableToProto(o.MaxUnavailable))
	p.SetReplacementMethod(ComputeInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(o.ReplacementMethod))
	return p
}

// InstanceGroupManagerUpdatePolicyMaxSurgeToProto converts a InstanceGroupManagerUpdatePolicyMaxSurge object to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o *compute.InstanceGroupManagerUpdatePolicyMaxSurge) *computepb.ComputeInstanceGroupManagerUpdatePolicyMaxSurge {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerUpdatePolicyMaxSurge{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerUpdatePolicyMaxUnavailableToProto converts a InstanceGroupManagerUpdatePolicyMaxUnavailable object to its proto representation.
func ComputeInstanceGroupManagerUpdatePolicyMaxUnavailableToProto(o *compute.InstanceGroupManagerUpdatePolicyMaxUnavailable) *computepb.ComputeInstanceGroupManagerUpdatePolicyMaxUnavailable {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerUpdatePolicyMaxUnavailable{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerNamedPortsToProto converts a InstanceGroupManagerNamedPorts object to its proto representation.
func ComputeInstanceGroupManagerNamedPortsToProto(o *compute.InstanceGroupManagerNamedPorts) *computepb.ComputeInstanceGroupManagerNamedPorts {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerNamedPorts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	return p
}

// InstanceGroupManagerStatefulPolicyToProto converts a InstanceGroupManagerStatefulPolicy object to its proto representation.
func ComputeInstanceGroupManagerStatefulPolicyToProto(o *compute.InstanceGroupManagerStatefulPolicy) *computepb.ComputeInstanceGroupManagerStatefulPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatefulPolicy{}
	p.SetPreservedState(ComputeInstanceGroupManagerStatefulPolicyPreservedStateToProto(o.PreservedState))
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateToProto converts a InstanceGroupManagerStatefulPolicyPreservedState object to its proto representation.
func ComputeInstanceGroupManagerStatefulPolicyPreservedStateToProto(o *compute.InstanceGroupManagerStatefulPolicyPreservedState) *computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedState {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedState{}
	mDisks := make(map[string]*computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks, len(o.Disks))
	for k, r := range o.Disks {
		mDisks[k] = ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(&r)
	}
	p.SetDisks(mDisks)
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks object to its proto representation.
func ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(o *compute.InstanceGroupManagerStatefulPolicyPreservedStateDisks) *computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisks{}
	p.SetAutoDelete(ComputeInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(o.AutoDelete))
	return p
}

// InstanceGroupManagerToProto converts a InstanceGroupManager resource to its proto representation.
func InstanceGroupManagerToProto(resource *compute.InstanceGroupManager) *computepb.ComputeInstanceGroupManager {
	p := &computepb.ComputeInstanceGroupManager{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetZone(dcl.ValueOrEmptyString(resource.Zone))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetDistributionPolicy(ComputeInstanceGroupManagerDistributionPolicyToProto(resource.DistributionPolicy))
	p.SetInstanceTemplate(dcl.ValueOrEmptyString(resource.InstanceTemplate))
	p.SetInstanceGroup(dcl.ValueOrEmptyString(resource.InstanceGroup))
	p.SetBaseInstanceName(dcl.ValueOrEmptyString(resource.BaseInstanceName))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetCurrentActions(ComputeInstanceGroupManagerCurrentActionsToProto(resource.CurrentActions))
	p.SetStatus(ComputeInstanceGroupManagerStatusToProto(resource.Status))
	p.SetTargetSize(dcl.ValueOrEmptyInt64(resource.TargetSize))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetUpdatePolicy(ComputeInstanceGroupManagerUpdatePolicyToProto(resource.UpdatePolicy))
	p.SetStatefulPolicy(ComputeInstanceGroupManagerStatefulPolicyToProto(resource.StatefulPolicy))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sVersions := make([]*computepb.ComputeInstanceGroupManagerVersions, len(resource.Versions))
	for i, r := range resource.Versions {
		sVersions[i] = ComputeInstanceGroupManagerVersionsToProto(&r)
	}
	p.SetVersions(sVersions)
	sTargetPools := make([]string, len(resource.TargetPools))
	for i, r := range resource.TargetPools {
		sTargetPools[i] = r
	}
	p.SetTargetPools(sTargetPools)
	sAutoHealingPolicies := make([]*computepb.ComputeInstanceGroupManagerAutoHealingPolicies, len(resource.AutoHealingPolicies))
	for i, r := range resource.AutoHealingPolicies {
		sAutoHealingPolicies[i] = ComputeInstanceGroupManagerAutoHealingPoliciesToProto(&r)
	}
	p.SetAutoHealingPolicies(sAutoHealingPolicies)
	sNamedPorts := make([]*computepb.ComputeInstanceGroupManagerNamedPorts, len(resource.NamedPorts))
	for i, r := range resource.NamedPorts {
		sNamedPorts[i] = ComputeInstanceGroupManagerNamedPortsToProto(&r)
	}
	p.SetNamedPorts(sNamedPorts)

	return p
}

// applyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) applyInstanceGroupManager(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeInstanceGroupManagerRequest) (*computepb.ComputeInstanceGroupManager, error) {
	p := ProtoToInstanceGroupManager(request.GetResource())
	res, err := c.ApplyInstanceGroupManager(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceGroupManagerToProto(res)
	return r, nil
}

// applyComputeInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) ApplyComputeInstanceGroupManager(ctx context.Context, request *computepb.ApplyComputeInstanceGroupManagerRequest) (*computepb.ComputeInstanceGroupManager, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstanceGroupManager(ctx, cl, request)
}

// DeleteInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Delete() method.
func (s *InstanceGroupManagerServer) DeleteComputeInstanceGroupManager(ctx context.Context, request *computepb.DeleteComputeInstanceGroupManagerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstanceGroupManager(ctx, ProtoToInstanceGroupManager(request.GetResource()))

}

// ListComputeInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManagerList() method.
func (s *InstanceGroupManagerServer) ListComputeInstanceGroupManager(ctx context.Context, request *computepb.ListComputeInstanceGroupManagerRequest) (*computepb.ListComputeInstanceGroupManagerResponse, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstanceGroupManager(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeInstanceGroupManager
	for _, r := range resources.Items {
		rp := InstanceGroupManagerToProto(r)
		protos = append(protos, rp)
	}
	p := &computepb.ListComputeInstanceGroupManagerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstanceGroupManager(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
