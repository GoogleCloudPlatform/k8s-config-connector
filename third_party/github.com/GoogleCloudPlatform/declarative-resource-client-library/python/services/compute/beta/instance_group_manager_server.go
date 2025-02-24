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

// InstanceGroupManagerServer implements the gRPC interface for InstanceGroupManager.
type InstanceGroupManagerServer struct{}

// ProtoToInstanceGroupManagerDistributionPolicyTargetShapeEnum converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(e betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum) *beta.InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerDistributionPolicyTargetShapeEnum(n[len("ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyTypeEnum converts a InstanceGroupManagerUpdatePolicyTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum) *beta.InstanceGroupManagerUpdatePolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyTypeEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) *beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMinimalActionEnum converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum) *beta.InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyMinimalActionEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyReplacementMethodEnum converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum) *beta.InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyReplacementMethodEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum converts a InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(e betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum) *beta.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(n[len("ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(e betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) *beta.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(n[len("ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(e betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum) *beta.InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(n[len("ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(e betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum) *beta.InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(n[len("ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerFailoverActionEnum converts a InstanceGroupManagerFailoverActionEnum enum from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerFailoverActionEnum(e betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum) *beta.InstanceGroupManagerFailoverActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum_name[int32(e)]; ok {
		e := beta.InstanceGroupManagerFailoverActionEnum(n[len("ComputeBetaInstanceGroupManagerFailoverActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerDistributionPolicy converts a InstanceGroupManagerDistributionPolicy object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerDistributionPolicy(p *betapb.ComputeBetaInstanceGroupManagerDistributionPolicy) *beta.InstanceGroupManagerDistributionPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerDistributionPolicy{
		TargetShape: ProtoToComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(p.GetTargetShape()),
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, *ProtoToComputeBetaInstanceGroupManagerDistributionPolicyZones(r))
	}
	return obj
}

// ProtoToInstanceGroupManagerDistributionPolicyZones converts a InstanceGroupManagerDistributionPolicyZones object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerDistributionPolicyZones(p *betapb.ComputeBetaInstanceGroupManagerDistributionPolicyZones) *beta.InstanceGroupManagerDistributionPolicyZones {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersions converts a InstanceGroupManagerVersions object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerVersions(p *betapb.ComputeBetaInstanceGroupManagerVersions) *beta.InstanceGroupManagerVersions {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerVersions{
		Name:             dcl.StringOrNil(p.GetName()),
		InstanceTemplate: dcl.StringOrNil(p.GetInstanceTemplate()),
		TargetSize:       ProtoToComputeBetaInstanceGroupManagerVersionsTargetSize(p.GetTargetSize()),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersionsTargetSize converts a InstanceGroupManagerVersionsTargetSize object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerVersionsTargetSize(p *betapb.ComputeBetaInstanceGroupManagerVersionsTargetSize) *beta.InstanceGroupManagerVersionsTargetSize {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerVersionsTargetSize{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerCurrentActions converts a InstanceGroupManagerCurrentActions object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerCurrentActions(p *betapb.ComputeBetaInstanceGroupManagerCurrentActions) *beta.InstanceGroupManagerCurrentActions {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerCurrentActions{
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
func ProtoToComputeBetaInstanceGroupManagerStatus(p *betapb.ComputeBetaInstanceGroupManagerStatus) *beta.InstanceGroupManagerStatus {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatus{
		IsStable:      dcl.Bool(p.GetIsStable()),
		VersionTarget: ProtoToComputeBetaInstanceGroupManagerStatusVersionTarget(p.GetVersionTarget()),
		Stateful:      ProtoToComputeBetaInstanceGroupManagerStatusStateful(p.GetStateful()),
		Autoscaler:    dcl.StringOrNil(p.GetAutoscaler()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusVersionTarget converts a InstanceGroupManagerStatusVersionTarget object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatusVersionTarget(p *betapb.ComputeBetaInstanceGroupManagerStatusVersionTarget) *beta.InstanceGroupManagerStatusVersionTarget {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.Bool(p.GetIsReached()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStateful converts a InstanceGroupManagerStatusStateful object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatusStateful(p *betapb.ComputeBetaInstanceGroupManagerStatusStateful) *beta.InstanceGroupManagerStatusStateful {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.Bool(p.GetHasStatefulConfig()),
		PerInstanceConfigs: ProtoToComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs(p.GetPerInstanceConfigs()),
		IsStateful:         dcl.Bool(p.GetIsStateful()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStatefulPerInstanceConfigs converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs(p *betapb.ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs) *beta.InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.Bool(p.GetAllEffective()),
	}
	return obj
}

// ProtoToInstanceGroupManagerAutoHealingPolicies converts a InstanceGroupManagerAutoHealingPolicies object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerAutoHealingPolicies(p *betapb.ComputeBetaInstanceGroupManagerAutoHealingPolicies) *beta.InstanceGroupManagerAutoHealingPolicies {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerAutoHealingPolicies{
		HealthCheck:     dcl.StringOrNil(p.GetHealthCheck()),
		InitialDelaySec: dcl.Int64OrNil(p.GetInitialDelaySec()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicy converts a InstanceGroupManagerUpdatePolicy object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicy(p *betapb.ComputeBetaInstanceGroupManagerUpdatePolicy) *beta.InstanceGroupManagerUpdatePolicy {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerUpdatePolicy{
		Type:                        ProtoToComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(p.GetType()),
		InstanceRedistributionType:  ProtoToComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(p.GetInstanceRedistributionType()),
		MinimalAction:               ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(p.GetMinimalAction()),
		MaxSurge:                    ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge(p.GetMaxSurge()),
		MaxUnavailable:              ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMaxUnavailable(p.GetMaxUnavailable()),
		ReplacementMethod:           ProtoToComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(p.GetReplacementMethod()),
		MostDisruptiveAllowedAction: ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(p.GetMostDisruptiveAllowedAction()),
		MinReadySec:                 dcl.Int64OrNil(p.GetMinReadySec()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxSurge converts a InstanceGroupManagerUpdatePolicyMaxSurge object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge(p *betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge) *beta.InstanceGroupManagerUpdatePolicyMaxSurge {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerUpdatePolicyMaxSurge{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxUnavailable converts a InstanceGroupManagerUpdatePolicyMaxUnavailable object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerUpdatePolicyMaxUnavailable(p *betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxUnavailable) *beta.InstanceGroupManagerUpdatePolicyMaxUnavailable {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerUpdatePolicyMaxUnavailable{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerNamedPorts converts a InstanceGroupManagerNamedPorts object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerNamedPorts(p *betapb.ComputeBetaInstanceGroupManagerNamedPorts) *beta.InstanceGroupManagerNamedPorts {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerNamedPorts{
		Name: dcl.StringOrNil(p.GetName()),
		Port: dcl.Int64OrNil(p.GetPort()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicy converts a InstanceGroupManagerStatefulPolicy object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicy(p *betapb.ComputeBetaInstanceGroupManagerStatefulPolicy) *beta.InstanceGroupManagerStatefulPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatefulPolicy{
		PreservedState: ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedState(p.GetPreservedState()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedState converts a InstanceGroupManagerStatefulPolicyPreservedState object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedState(p *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedState) *beta.InstanceGroupManagerStatefulPolicyPreservedState {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatefulPolicyPreservedState{}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisks converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks(p *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks) *beta.InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateInternalIps converts a InstanceGroupManagerStatefulPolicyPreservedStateInternalIps object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps(p *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps) *beta.InstanceGroupManagerStatefulPolicyPreservedStateInternalIps {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatefulPolicyPreservedStateInternalIps{
		AutoDelete: ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateExternalIps converts a InstanceGroupManagerStatefulPolicyPreservedStateExternalIps object from its proto representation.
func ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps(p *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps) *beta.InstanceGroupManagerStatefulPolicyPreservedStateExternalIps {
	if p == nil {
		return nil
	}
	obj := &beta.InstanceGroupManagerStatefulPolicyPreservedStateExternalIps{
		AutoDelete: ProtoToComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManager converts a InstanceGroupManager resource from its proto representation.
func ProtoToInstanceGroupManager(p *betapb.ComputeBetaInstanceGroupManager) *beta.InstanceGroupManager {
	obj := &beta.InstanceGroupManager{
		Id:                 dcl.Int64OrNil(p.GetId()),
		CreationTimestamp:  dcl.StringOrNil(p.GetCreationTimestamp()),
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		Zone:               dcl.StringOrNil(p.GetZone()),
		Region:             dcl.StringOrNil(p.GetRegion()),
		DistributionPolicy: ProtoToComputeBetaInstanceGroupManagerDistributionPolicy(p.GetDistributionPolicy()),
		InstanceTemplate:   dcl.StringOrNil(p.GetInstanceTemplate()),
		InstanceGroup:      dcl.StringOrNil(p.GetInstanceGroup()),
		BaseInstanceName:   dcl.StringOrNil(p.GetBaseInstanceName()),
		Fingerprint:        dcl.StringOrNil(p.GetFingerprint()),
		CurrentActions:     ProtoToComputeBetaInstanceGroupManagerCurrentActions(p.GetCurrentActions()),
		Status:             ProtoToComputeBetaInstanceGroupManagerStatus(p.GetStatus()),
		TargetSize:         dcl.Int64OrNil(p.GetTargetSize()),
		SelfLink:           dcl.StringOrNil(p.GetSelfLink()),
		UpdatePolicy:       ProtoToComputeBetaInstanceGroupManagerUpdatePolicy(p.GetUpdatePolicy()),
		StatefulPolicy:     ProtoToComputeBetaInstanceGroupManagerStatefulPolicy(p.GetStatefulPolicy()),
		ServiceAccount:     dcl.StringOrNil(p.GetServiceAccount()),
		FailoverAction:     ProtoToComputeBetaInstanceGroupManagerFailoverActionEnum(p.GetFailoverAction()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetVersions() {
		obj.Versions = append(obj.Versions, *ProtoToComputeBetaInstanceGroupManagerVersions(r))
	}
	for _, r := range p.GetTargetPools() {
		obj.TargetPools = append(obj.TargetPools, r)
	}
	for _, r := range p.GetAutoHealingPolicies() {
		obj.AutoHealingPolicies = append(obj.AutoHealingPolicies, *ProtoToComputeBetaInstanceGroupManagerAutoHealingPolicies(r))
	}
	for _, r := range p.GetNamedPorts() {
		obj.NamedPorts = append(obj.NamedPorts, *ProtoToComputeBetaInstanceGroupManagerNamedPorts(r))
	}
	return obj
}

// InstanceGroupManagerDistributionPolicyTargetShapeEnumToProto converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(e *beta.InstanceGroupManagerDistributionPolicyTargetShapeEnum) betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum_value["InstanceGroupManagerDistributionPolicyTargetShapeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
}

// InstanceGroupManagerUpdatePolicyTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyTypeEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyTypeEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum_value["InstanceGroupManagerUpdatePolicyTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_value["InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyMinimalActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyMinimalActionEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum_value["InstanceGroupManagerUpdatePolicyMinimalActionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
}

// InstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyReplacementMethodEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum_value["InstanceGroupManagerUpdatePolicyReplacementMethodEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
}

// InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto(e *beta.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum) betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum_value["InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(e *beta.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnumToProto(e *beta.InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum) betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnumToProto(e *beta.InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum) betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(0)
}

// InstanceGroupManagerFailoverActionEnumToProto converts a InstanceGroupManagerFailoverActionEnum enum to its proto representation.
func ComputeBetaInstanceGroupManagerFailoverActionEnumToProto(e *beta.InstanceGroupManagerFailoverActionEnum) betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum {
	if e == nil {
		return betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum(0)
	}
	if v, ok := betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum_value["InstanceGroupManagerFailoverActionEnum"+string(*e)]; ok {
		return betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum(v)
	}
	return betapb.ComputeBetaInstanceGroupManagerFailoverActionEnum(0)
}

// InstanceGroupManagerDistributionPolicyToProto converts a InstanceGroupManagerDistributionPolicy object to its proto representation.
func ComputeBetaInstanceGroupManagerDistributionPolicyToProto(o *beta.InstanceGroupManagerDistributionPolicy) *betapb.ComputeBetaInstanceGroupManagerDistributionPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerDistributionPolicy{}
	p.SetTargetShape(ComputeBetaInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(o.TargetShape))
	sZones := make([]*betapb.ComputeBetaInstanceGroupManagerDistributionPolicyZones, len(o.Zones))
	for i, r := range o.Zones {
		sZones[i] = ComputeBetaInstanceGroupManagerDistributionPolicyZonesToProto(&r)
	}
	p.SetZones(sZones)
	return p
}

// InstanceGroupManagerDistributionPolicyZonesToProto converts a InstanceGroupManagerDistributionPolicyZones object to its proto representation.
func ComputeBetaInstanceGroupManagerDistributionPolicyZonesToProto(o *beta.InstanceGroupManagerDistributionPolicyZones) *betapb.ComputeBetaInstanceGroupManagerDistributionPolicyZones {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerDistributionPolicyZones{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	return p
}

// InstanceGroupManagerVersionsToProto converts a InstanceGroupManagerVersions object to its proto representation.
func ComputeBetaInstanceGroupManagerVersionsToProto(o *beta.InstanceGroupManagerVersions) *betapb.ComputeBetaInstanceGroupManagerVersions {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerVersions{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetInstanceTemplate(dcl.ValueOrEmptyString(o.InstanceTemplate))
	p.SetTargetSize(ComputeBetaInstanceGroupManagerVersionsTargetSizeToProto(o.TargetSize))
	return p
}

// InstanceGroupManagerVersionsTargetSizeToProto converts a InstanceGroupManagerVersionsTargetSize object to its proto representation.
func ComputeBetaInstanceGroupManagerVersionsTargetSizeToProto(o *beta.InstanceGroupManagerVersionsTargetSize) *betapb.ComputeBetaInstanceGroupManagerVersionsTargetSize {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerVersionsTargetSize{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerCurrentActionsToProto converts a InstanceGroupManagerCurrentActions object to its proto representation.
func ComputeBetaInstanceGroupManagerCurrentActionsToProto(o *beta.InstanceGroupManagerCurrentActions) *betapb.ComputeBetaInstanceGroupManagerCurrentActions {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerCurrentActions{}
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
func ComputeBetaInstanceGroupManagerStatusToProto(o *beta.InstanceGroupManagerStatus) *betapb.ComputeBetaInstanceGroupManagerStatus {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatus{}
	p.SetIsStable(dcl.ValueOrEmptyBool(o.IsStable))
	p.SetVersionTarget(ComputeBetaInstanceGroupManagerStatusVersionTargetToProto(o.VersionTarget))
	p.SetStateful(ComputeBetaInstanceGroupManagerStatusStatefulToProto(o.Stateful))
	p.SetAutoscaler(dcl.ValueOrEmptyString(o.Autoscaler))
	return p
}

// InstanceGroupManagerStatusVersionTargetToProto converts a InstanceGroupManagerStatusVersionTarget object to its proto representation.
func ComputeBetaInstanceGroupManagerStatusVersionTargetToProto(o *beta.InstanceGroupManagerStatusVersionTarget) *betapb.ComputeBetaInstanceGroupManagerStatusVersionTarget {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatusVersionTarget{}
	p.SetIsReached(dcl.ValueOrEmptyBool(o.IsReached))
	return p
}

// InstanceGroupManagerStatusStatefulToProto converts a InstanceGroupManagerStatusStateful object to its proto representation.
func ComputeBetaInstanceGroupManagerStatusStatefulToProto(o *beta.InstanceGroupManagerStatusStateful) *betapb.ComputeBetaInstanceGroupManagerStatusStateful {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatusStateful{}
	p.SetHasStatefulConfig(dcl.ValueOrEmptyBool(o.HasStatefulConfig))
	p.SetPerInstanceConfigs(ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o.PerInstanceConfigs))
	p.SetIsStateful(dcl.ValueOrEmptyBool(o.IsStateful))
	return p
}

// InstanceGroupManagerStatusStatefulPerInstanceConfigsToProto converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs object to its proto representation.
func ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o *beta.InstanceGroupManagerStatusStatefulPerInstanceConfigs) *betapb.ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatusStatefulPerInstanceConfigs{}
	p.SetAllEffective(dcl.ValueOrEmptyBool(o.AllEffective))
	return p
}

// InstanceGroupManagerAutoHealingPoliciesToProto converts a InstanceGroupManagerAutoHealingPolicies object to its proto representation.
func ComputeBetaInstanceGroupManagerAutoHealingPoliciesToProto(o *beta.InstanceGroupManagerAutoHealingPolicies) *betapb.ComputeBetaInstanceGroupManagerAutoHealingPolicies {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerAutoHealingPolicies{}
	p.SetHealthCheck(dcl.ValueOrEmptyString(o.HealthCheck))
	p.SetInitialDelaySec(dcl.ValueOrEmptyInt64(o.InitialDelaySec))
	return p
}

// InstanceGroupManagerUpdatePolicyToProto converts a InstanceGroupManagerUpdatePolicy object to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyToProto(o *beta.InstanceGroupManagerUpdatePolicy) *betapb.ComputeBetaInstanceGroupManagerUpdatePolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerUpdatePolicy{}
	p.SetType(ComputeBetaInstanceGroupManagerUpdatePolicyTypeEnumToProto(o.Type))
	p.SetInstanceRedistributionType(ComputeBetaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(o.InstanceRedistributionType))
	p.SetMinimalAction(ComputeBetaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(o.MinimalAction))
	p.SetMaxSurge(ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o.MaxSurge))
	p.SetMaxUnavailable(ComputeBetaInstanceGroupManagerUpdatePolicyMaxUnavailableToProto(o.MaxUnavailable))
	p.SetReplacementMethod(ComputeBetaInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(o.ReplacementMethod))
	p.SetMostDisruptiveAllowedAction(ComputeBetaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto(o.MostDisruptiveAllowedAction))
	p.SetMinReadySec(dcl.ValueOrEmptyInt64(o.MinReadySec))
	return p
}

// InstanceGroupManagerUpdatePolicyMaxSurgeToProto converts a InstanceGroupManagerUpdatePolicyMaxSurge object to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o *beta.InstanceGroupManagerUpdatePolicyMaxSurge) *betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxSurge{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerUpdatePolicyMaxUnavailableToProto converts a InstanceGroupManagerUpdatePolicyMaxUnavailable object to its proto representation.
func ComputeBetaInstanceGroupManagerUpdatePolicyMaxUnavailableToProto(o *beta.InstanceGroupManagerUpdatePolicyMaxUnavailable) *betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxUnavailable {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerUpdatePolicyMaxUnavailable{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerNamedPortsToProto converts a InstanceGroupManagerNamedPorts object to its proto representation.
func ComputeBetaInstanceGroupManagerNamedPortsToProto(o *beta.InstanceGroupManagerNamedPorts) *betapb.ComputeBetaInstanceGroupManagerNamedPorts {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerNamedPorts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	return p
}

// InstanceGroupManagerStatefulPolicyToProto converts a InstanceGroupManagerStatefulPolicy object to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyToProto(o *beta.InstanceGroupManagerStatefulPolicy) *betapb.ComputeBetaInstanceGroupManagerStatefulPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatefulPolicy{}
	p.SetPreservedState(ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateToProto(o.PreservedState))
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateToProto converts a InstanceGroupManagerStatefulPolicyPreservedState object to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateToProto(o *beta.InstanceGroupManagerStatefulPolicyPreservedState) *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedState {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedState{}
	mDisks := make(map[string]*betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks, len(o.Disks))
	for k, r := range o.Disks {
		mDisks[k] = ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(&r)
	}
	p.SetDisks(mDisks)
	mInternalIps := make(map[string]*betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps, len(o.InternalIps))
	for k, r := range o.InternalIps {
		mInternalIps[k] = ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsToProto(&r)
	}
	p.SetInternalIps(mInternalIps)
	mExternalIps := make(map[string]*betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps, len(o.ExternalIps))
	for k, r := range o.ExternalIps {
		mExternalIps[k] = ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsToProto(&r)
	}
	p.SetExternalIps(mExternalIps)
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks object to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(o *beta.InstanceGroupManagerStatefulPolicyPreservedStateDisks) *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisks{}
	p.SetAutoDelete(ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(o.AutoDelete))
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateInternalIps object to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsToProto(o *beta.InstanceGroupManagerStatefulPolicyPreservedStateInternalIps) *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps{}
	p.SetAutoDelete(ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnumToProto(o.AutoDelete))
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateExternalIps object to its proto representation.
func ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsToProto(o *beta.InstanceGroupManagerStatefulPolicyPreservedStateExternalIps) *betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps{}
	p.SetAutoDelete(ComputeBetaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnumToProto(o.AutoDelete))
	return p
}

// InstanceGroupManagerToProto converts a InstanceGroupManager resource to its proto representation.
func InstanceGroupManagerToProto(resource *beta.InstanceGroupManager) *betapb.ComputeBetaInstanceGroupManager {
	p := &betapb.ComputeBetaInstanceGroupManager{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetZone(dcl.ValueOrEmptyString(resource.Zone))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetDistributionPolicy(ComputeBetaInstanceGroupManagerDistributionPolicyToProto(resource.DistributionPolicy))
	p.SetInstanceTemplate(dcl.ValueOrEmptyString(resource.InstanceTemplate))
	p.SetInstanceGroup(dcl.ValueOrEmptyString(resource.InstanceGroup))
	p.SetBaseInstanceName(dcl.ValueOrEmptyString(resource.BaseInstanceName))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetCurrentActions(ComputeBetaInstanceGroupManagerCurrentActionsToProto(resource.CurrentActions))
	p.SetStatus(ComputeBetaInstanceGroupManagerStatusToProto(resource.Status))
	p.SetTargetSize(dcl.ValueOrEmptyInt64(resource.TargetSize))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetUpdatePolicy(ComputeBetaInstanceGroupManagerUpdatePolicyToProto(resource.UpdatePolicy))
	p.SetStatefulPolicy(ComputeBetaInstanceGroupManagerStatefulPolicyToProto(resource.StatefulPolicy))
	p.SetServiceAccount(dcl.ValueOrEmptyString(resource.ServiceAccount))
	p.SetFailoverAction(ComputeBetaInstanceGroupManagerFailoverActionEnumToProto(resource.FailoverAction))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sVersions := make([]*betapb.ComputeBetaInstanceGroupManagerVersions, len(resource.Versions))
	for i, r := range resource.Versions {
		sVersions[i] = ComputeBetaInstanceGroupManagerVersionsToProto(&r)
	}
	p.SetVersions(sVersions)
	sTargetPools := make([]string, len(resource.TargetPools))
	for i, r := range resource.TargetPools {
		sTargetPools[i] = r
	}
	p.SetTargetPools(sTargetPools)
	sAutoHealingPolicies := make([]*betapb.ComputeBetaInstanceGroupManagerAutoHealingPolicies, len(resource.AutoHealingPolicies))
	for i, r := range resource.AutoHealingPolicies {
		sAutoHealingPolicies[i] = ComputeBetaInstanceGroupManagerAutoHealingPoliciesToProto(&r)
	}
	p.SetAutoHealingPolicies(sAutoHealingPolicies)
	sNamedPorts := make([]*betapb.ComputeBetaInstanceGroupManagerNamedPorts, len(resource.NamedPorts))
	for i, r := range resource.NamedPorts {
		sNamedPorts[i] = ComputeBetaInstanceGroupManagerNamedPortsToProto(&r)
	}
	p.SetNamedPorts(sNamedPorts)

	return p
}

// applyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) applyInstanceGroupManager(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaInstanceGroupManagerRequest) (*betapb.ComputeBetaInstanceGroupManager, error) {
	p := ProtoToInstanceGroupManager(request.GetResource())
	res, err := c.ApplyInstanceGroupManager(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceGroupManagerToProto(res)
	return r, nil
}

// applyComputeBetaInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) ApplyComputeBetaInstanceGroupManager(ctx context.Context, request *betapb.ApplyComputeBetaInstanceGroupManagerRequest) (*betapb.ComputeBetaInstanceGroupManager, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstanceGroupManager(ctx, cl, request)
}

// DeleteInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Delete() method.
func (s *InstanceGroupManagerServer) DeleteComputeBetaInstanceGroupManager(ctx context.Context, request *betapb.DeleteComputeBetaInstanceGroupManagerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstanceGroupManager(ctx, ProtoToInstanceGroupManager(request.GetResource()))

}

// ListComputeBetaInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManagerList() method.
func (s *InstanceGroupManagerServer) ListComputeBetaInstanceGroupManager(ctx context.Context, request *betapb.ListComputeBetaInstanceGroupManagerRequest) (*betapb.ListComputeBetaInstanceGroupManagerResponse, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstanceGroupManager(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaInstanceGroupManager
	for _, r := range resources.Items {
		rp := InstanceGroupManagerToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListComputeBetaInstanceGroupManagerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstanceGroupManager(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
