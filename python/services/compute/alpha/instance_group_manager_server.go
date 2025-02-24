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

// InstanceGroupManagerServer implements the gRPC interface for InstanceGroupManager.
type InstanceGroupManagerServer struct{}

// ProtoToInstanceGroupManagerDistributionPolicyTargetShapeEnum converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(e alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum) *alpha.InstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerDistributionPolicyTargetShapeEnum(n[len("ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyTypeEnum converts a InstanceGroupManagerUpdatePolicyTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum) *alpha.InstanceGroupManagerUpdatePolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyTypeEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) *alpha.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMinimalActionEnum converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum) *alpha.InstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyMinimalActionEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyReplacementMethodEnum converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum) *alpha.InstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyReplacementMethodEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum converts a InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(e alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum) *alpha.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(n[len("ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(e alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) *alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(n[len("ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(e alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum) *alpha.InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(n[len("ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum converts a InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(e alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum) *alpha.InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(n[len("ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerFailoverActionEnum converts a InstanceGroupManagerFailoverActionEnum enum from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerFailoverActionEnum(e alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum) *alpha.InstanceGroupManagerFailoverActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum_name[int32(e)]; ok {
		e := alpha.InstanceGroupManagerFailoverActionEnum(n[len("ComputeAlphaInstanceGroupManagerFailoverActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceGroupManagerDistributionPolicy converts a InstanceGroupManagerDistributionPolicy object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerDistributionPolicy(p *alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicy) *alpha.InstanceGroupManagerDistributionPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerDistributionPolicy{
		TargetShape: ProtoToComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(p.GetTargetShape()),
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, *ProtoToComputeAlphaInstanceGroupManagerDistributionPolicyZones(r))
	}
	return obj
}

// ProtoToInstanceGroupManagerDistributionPolicyZones converts a InstanceGroupManagerDistributionPolicyZones object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerDistributionPolicyZones(p *alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyZones) *alpha.InstanceGroupManagerDistributionPolicyZones {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerDistributionPolicyZones{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersions converts a InstanceGroupManagerVersions object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerVersions(p *alphapb.ComputeAlphaInstanceGroupManagerVersions) *alpha.InstanceGroupManagerVersions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerVersions{
		Name:             dcl.StringOrNil(p.GetName()),
		InstanceTemplate: dcl.StringOrNil(p.GetInstanceTemplate()),
		TargetSize:       ProtoToComputeAlphaInstanceGroupManagerVersionsTargetSize(p.GetTargetSize()),
	}
	return obj
}

// ProtoToInstanceGroupManagerVersionsTargetSize converts a InstanceGroupManagerVersionsTargetSize object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerVersionsTargetSize(p *alphapb.ComputeAlphaInstanceGroupManagerVersionsTargetSize) *alpha.InstanceGroupManagerVersionsTargetSize {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerVersionsTargetSize{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerCurrentActions converts a InstanceGroupManagerCurrentActions object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerCurrentActions(p *alphapb.ComputeAlphaInstanceGroupManagerCurrentActions) *alpha.InstanceGroupManagerCurrentActions {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerCurrentActions{
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
func ProtoToComputeAlphaInstanceGroupManagerStatus(p *alphapb.ComputeAlphaInstanceGroupManagerStatus) *alpha.InstanceGroupManagerStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatus{
		IsStable:      dcl.Bool(p.GetIsStable()),
		VersionTarget: ProtoToComputeAlphaInstanceGroupManagerStatusVersionTarget(p.GetVersionTarget()),
		Stateful:      ProtoToComputeAlphaInstanceGroupManagerStatusStateful(p.GetStateful()),
		Autoscaler:    dcl.StringOrNil(p.GetAutoscaler()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusVersionTarget converts a InstanceGroupManagerStatusVersionTarget object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatusVersionTarget(p *alphapb.ComputeAlphaInstanceGroupManagerStatusVersionTarget) *alpha.InstanceGroupManagerStatusVersionTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatusVersionTarget{
		IsReached: dcl.Bool(p.GetIsReached()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStateful converts a InstanceGroupManagerStatusStateful object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatusStateful(p *alphapb.ComputeAlphaInstanceGroupManagerStatusStateful) *alpha.InstanceGroupManagerStatusStateful {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatusStateful{
		HasStatefulConfig:  dcl.Bool(p.GetHasStatefulConfig()),
		PerInstanceConfigs: ProtoToComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs(p.GetPerInstanceConfigs()),
		IsStateful:         dcl.Bool(p.GetIsStateful()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatusStatefulPerInstanceConfigs converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs(p *alphapb.ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs) *alpha.InstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatusStatefulPerInstanceConfigs{
		AllEffective: dcl.Bool(p.GetAllEffective()),
	}
	return obj
}

// ProtoToInstanceGroupManagerAutoHealingPolicies converts a InstanceGroupManagerAutoHealingPolicies object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerAutoHealingPolicies(p *alphapb.ComputeAlphaInstanceGroupManagerAutoHealingPolicies) *alpha.InstanceGroupManagerAutoHealingPolicies {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerAutoHealingPolicies{
		HealthCheck:     dcl.StringOrNil(p.GetHealthCheck()),
		InitialDelaySec: dcl.Int64OrNil(p.GetInitialDelaySec()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicy converts a InstanceGroupManagerUpdatePolicy object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicy(p *alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicy) *alpha.InstanceGroupManagerUpdatePolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerUpdatePolicy{
		Type:                        ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(p.GetType()),
		InstanceRedistributionType:  ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(p.GetInstanceRedistributionType()),
		MinimalAction:               ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(p.GetMinimalAction()),
		MaxSurge:                    ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMaxSurge(p.GetMaxSurge()),
		MaxUnavailable:              ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMaxUnavailable(p.GetMaxUnavailable()),
		ReplacementMethod:           ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(p.GetReplacementMethod()),
		MostDisruptiveAllowedAction: ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(p.GetMostDisruptiveAllowedAction()),
		MinReadySec:                 dcl.Int64OrNil(p.GetMinReadySec()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxSurge converts a InstanceGroupManagerUpdatePolicyMaxSurge object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMaxSurge(p *alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMaxSurge) *alpha.InstanceGroupManagerUpdatePolicyMaxSurge {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerUpdatePolicyMaxSurge{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerUpdatePolicyMaxUnavailable converts a InstanceGroupManagerUpdatePolicyMaxUnavailable object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerUpdatePolicyMaxUnavailable(p *alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMaxUnavailable) *alpha.InstanceGroupManagerUpdatePolicyMaxUnavailable {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerUpdatePolicyMaxUnavailable{
		Fixed:      dcl.Int64OrNil(p.GetFixed()),
		Percent:    dcl.Int64OrNil(p.GetPercent()),
		Calculated: dcl.Int64OrNil(p.GetCalculated()),
	}
	return obj
}

// ProtoToInstanceGroupManagerNamedPorts converts a InstanceGroupManagerNamedPorts object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerNamedPorts(p *alphapb.ComputeAlphaInstanceGroupManagerNamedPorts) *alpha.InstanceGroupManagerNamedPorts {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerNamedPorts{
		Name: dcl.StringOrNil(p.GetName()),
		Port: dcl.Int64OrNil(p.GetPort()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicy converts a InstanceGroupManagerStatefulPolicy object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicy(p *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicy) *alpha.InstanceGroupManagerStatefulPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatefulPolicy{
		PreservedState: ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState(p.GetPreservedState()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedState converts a InstanceGroupManagerStatefulPolicyPreservedState object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState(p *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState) *alpha.InstanceGroupManagerStatefulPolicyPreservedState {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatefulPolicyPreservedState{}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateDisks converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks(p *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks) *alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisks{
		AutoDelete: ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateInternalIps converts a InstanceGroupManagerStatefulPolicyPreservedStateInternalIps object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps(p *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps) *alpha.InstanceGroupManagerStatefulPolicyPreservedStateInternalIps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatefulPolicyPreservedStateInternalIps{
		AutoDelete: ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManagerStatefulPolicyPreservedStateExternalIps converts a InstanceGroupManagerStatefulPolicyPreservedStateExternalIps object from its proto representation.
func ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps(p *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps) *alpha.InstanceGroupManagerStatefulPolicyPreservedStateExternalIps {
	if p == nil {
		return nil
	}
	obj := &alpha.InstanceGroupManagerStatefulPolicyPreservedStateExternalIps{
		AutoDelete: ProtoToComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(p.GetAutoDelete()),
	}
	return obj
}

// ProtoToInstanceGroupManager converts a InstanceGroupManager resource from its proto representation.
func ProtoToInstanceGroupManager(p *alphapb.ComputeAlphaInstanceGroupManager) *alpha.InstanceGroupManager {
	obj := &alpha.InstanceGroupManager{
		Id:                 dcl.Int64OrNil(p.GetId()),
		CreationTimestamp:  dcl.StringOrNil(p.GetCreationTimestamp()),
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		Zone:               dcl.StringOrNil(p.GetZone()),
		Region:             dcl.StringOrNil(p.GetRegion()),
		DistributionPolicy: ProtoToComputeAlphaInstanceGroupManagerDistributionPolicy(p.GetDistributionPolicy()),
		InstanceTemplate:   dcl.StringOrNil(p.GetInstanceTemplate()),
		InstanceGroup:      dcl.StringOrNil(p.GetInstanceGroup()),
		BaseInstanceName:   dcl.StringOrNil(p.GetBaseInstanceName()),
		Fingerprint:        dcl.StringOrNil(p.GetFingerprint()),
		CurrentActions:     ProtoToComputeAlphaInstanceGroupManagerCurrentActions(p.GetCurrentActions()),
		Status:             ProtoToComputeAlphaInstanceGroupManagerStatus(p.GetStatus()),
		TargetSize:         dcl.Int64OrNil(p.GetTargetSize()),
		SelfLink:           dcl.StringOrNil(p.GetSelfLink()),
		UpdatePolicy:       ProtoToComputeAlphaInstanceGroupManagerUpdatePolicy(p.GetUpdatePolicy()),
		StatefulPolicy:     ProtoToComputeAlphaInstanceGroupManagerStatefulPolicy(p.GetStatefulPolicy()),
		ServiceAccount:     dcl.StringOrNil(p.GetServiceAccount()),
		FailoverAction:     ProtoToComputeAlphaInstanceGroupManagerFailoverActionEnum(p.GetFailoverAction()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetVersions() {
		obj.Versions = append(obj.Versions, *ProtoToComputeAlphaInstanceGroupManagerVersions(r))
	}
	for _, r := range p.GetTargetPools() {
		obj.TargetPools = append(obj.TargetPools, r)
	}
	for _, r := range p.GetAutoHealingPolicies() {
		obj.AutoHealingPolicies = append(obj.AutoHealingPolicies, *ProtoToComputeAlphaInstanceGroupManagerAutoHealingPolicies(r))
	}
	for _, r := range p.GetNamedPorts() {
		obj.NamedPorts = append(obj.NamedPorts, *ProtoToComputeAlphaInstanceGroupManagerNamedPorts(r))
	}
	return obj
}

// InstanceGroupManagerDistributionPolicyTargetShapeEnumToProto converts a InstanceGroupManagerDistributionPolicyTargetShapeEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(e *alpha.InstanceGroupManagerDistributionPolicyTargetShapeEnum) alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum_value["InstanceGroupManagerDistributionPolicyTargetShapeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnum(0)
}

// InstanceGroupManagerUpdatePolicyTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyTypeEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyTypeEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum_value["InstanceGroupManagerUpdatePolicyTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto converts a InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum_value["InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnum(0)
}

// InstanceGroupManagerUpdatePolicyMinimalActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMinimalActionEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyMinimalActionEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum_value["InstanceGroupManagerUpdatePolicyMinimalActionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnum(0)
}

// InstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto converts a InstanceGroupManagerUpdatePolicyReplacementMethodEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyReplacementMethodEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum_value["InstanceGroupManagerUpdatePolicyReplacementMethodEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnum(0)
}

// InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto converts a InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto(e *alpha.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum) alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum_value["InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(e *alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum) alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnumToProto(e *alpha.InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum) alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnum(0)
}

// InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnumToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnumToProto(e *alpha.InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum) alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum_value["InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnum(0)
}

// InstanceGroupManagerFailoverActionEnumToProto converts a InstanceGroupManagerFailoverActionEnum enum to its proto representation.
func ComputeAlphaInstanceGroupManagerFailoverActionEnumToProto(e *alpha.InstanceGroupManagerFailoverActionEnum) alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum {
	if e == nil {
		return alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum(0)
	}
	if v, ok := alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum_value["InstanceGroupManagerFailoverActionEnum"+string(*e)]; ok {
		return alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum(v)
	}
	return alphapb.ComputeAlphaInstanceGroupManagerFailoverActionEnum(0)
}

// InstanceGroupManagerDistributionPolicyToProto converts a InstanceGroupManagerDistributionPolicy object to its proto representation.
func ComputeAlphaInstanceGroupManagerDistributionPolicyToProto(o *alpha.InstanceGroupManagerDistributionPolicy) *alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicy{}
	p.SetTargetShape(ComputeAlphaInstanceGroupManagerDistributionPolicyTargetShapeEnumToProto(o.TargetShape))
	sZones := make([]*alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyZones, len(o.Zones))
	for i, r := range o.Zones {
		sZones[i] = ComputeAlphaInstanceGroupManagerDistributionPolicyZonesToProto(&r)
	}
	p.SetZones(sZones)
	return p
}

// InstanceGroupManagerDistributionPolicyZonesToProto converts a InstanceGroupManagerDistributionPolicyZones object to its proto representation.
func ComputeAlphaInstanceGroupManagerDistributionPolicyZonesToProto(o *alpha.InstanceGroupManagerDistributionPolicyZones) *alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyZones {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerDistributionPolicyZones{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	return p
}

// InstanceGroupManagerVersionsToProto converts a InstanceGroupManagerVersions object to its proto representation.
func ComputeAlphaInstanceGroupManagerVersionsToProto(o *alpha.InstanceGroupManagerVersions) *alphapb.ComputeAlphaInstanceGroupManagerVersions {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerVersions{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetInstanceTemplate(dcl.ValueOrEmptyString(o.InstanceTemplate))
	p.SetTargetSize(ComputeAlphaInstanceGroupManagerVersionsTargetSizeToProto(o.TargetSize))
	return p
}

// InstanceGroupManagerVersionsTargetSizeToProto converts a InstanceGroupManagerVersionsTargetSize object to its proto representation.
func ComputeAlphaInstanceGroupManagerVersionsTargetSizeToProto(o *alpha.InstanceGroupManagerVersionsTargetSize) *alphapb.ComputeAlphaInstanceGroupManagerVersionsTargetSize {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerVersionsTargetSize{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerCurrentActionsToProto converts a InstanceGroupManagerCurrentActions object to its proto representation.
func ComputeAlphaInstanceGroupManagerCurrentActionsToProto(o *alpha.InstanceGroupManagerCurrentActions) *alphapb.ComputeAlphaInstanceGroupManagerCurrentActions {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerCurrentActions{}
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
func ComputeAlphaInstanceGroupManagerStatusToProto(o *alpha.InstanceGroupManagerStatus) *alphapb.ComputeAlphaInstanceGroupManagerStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatus{}
	p.SetIsStable(dcl.ValueOrEmptyBool(o.IsStable))
	p.SetVersionTarget(ComputeAlphaInstanceGroupManagerStatusVersionTargetToProto(o.VersionTarget))
	p.SetStateful(ComputeAlphaInstanceGroupManagerStatusStatefulToProto(o.Stateful))
	p.SetAutoscaler(dcl.ValueOrEmptyString(o.Autoscaler))
	return p
}

// InstanceGroupManagerStatusVersionTargetToProto converts a InstanceGroupManagerStatusVersionTarget object to its proto representation.
func ComputeAlphaInstanceGroupManagerStatusVersionTargetToProto(o *alpha.InstanceGroupManagerStatusVersionTarget) *alphapb.ComputeAlphaInstanceGroupManagerStatusVersionTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatusVersionTarget{}
	p.SetIsReached(dcl.ValueOrEmptyBool(o.IsReached))
	return p
}

// InstanceGroupManagerStatusStatefulToProto converts a InstanceGroupManagerStatusStateful object to its proto representation.
func ComputeAlphaInstanceGroupManagerStatusStatefulToProto(o *alpha.InstanceGroupManagerStatusStateful) *alphapb.ComputeAlphaInstanceGroupManagerStatusStateful {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatusStateful{}
	p.SetHasStatefulConfig(dcl.ValueOrEmptyBool(o.HasStatefulConfig))
	p.SetPerInstanceConfigs(ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o.PerInstanceConfigs))
	p.SetIsStateful(dcl.ValueOrEmptyBool(o.IsStateful))
	return p
}

// InstanceGroupManagerStatusStatefulPerInstanceConfigsToProto converts a InstanceGroupManagerStatusStatefulPerInstanceConfigs object to its proto representation.
func ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigsToProto(o *alpha.InstanceGroupManagerStatusStatefulPerInstanceConfigs) *alphapb.ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatusStatefulPerInstanceConfigs{}
	p.SetAllEffective(dcl.ValueOrEmptyBool(o.AllEffective))
	return p
}

// InstanceGroupManagerAutoHealingPoliciesToProto converts a InstanceGroupManagerAutoHealingPolicies object to its proto representation.
func ComputeAlphaInstanceGroupManagerAutoHealingPoliciesToProto(o *alpha.InstanceGroupManagerAutoHealingPolicies) *alphapb.ComputeAlphaInstanceGroupManagerAutoHealingPolicies {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerAutoHealingPolicies{}
	p.SetHealthCheck(dcl.ValueOrEmptyString(o.HealthCheck))
	p.SetInitialDelaySec(dcl.ValueOrEmptyInt64(o.InitialDelaySec))
	return p
}

// InstanceGroupManagerUpdatePolicyToProto converts a InstanceGroupManagerUpdatePolicy object to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyToProto(o *alpha.InstanceGroupManagerUpdatePolicy) *alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicy{}
	p.SetType(ComputeAlphaInstanceGroupManagerUpdatePolicyTypeEnumToProto(o.Type))
	p.SetInstanceRedistributionType(ComputeAlphaInstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumToProto(o.InstanceRedistributionType))
	p.SetMinimalAction(ComputeAlphaInstanceGroupManagerUpdatePolicyMinimalActionEnumToProto(o.MinimalAction))
	p.SetMaxSurge(ComputeAlphaInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o.MaxSurge))
	p.SetMaxUnavailable(ComputeAlphaInstanceGroupManagerUpdatePolicyMaxUnavailableToProto(o.MaxUnavailable))
	p.SetReplacementMethod(ComputeAlphaInstanceGroupManagerUpdatePolicyReplacementMethodEnumToProto(o.ReplacementMethod))
	p.SetMostDisruptiveAllowedAction(ComputeAlphaInstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumToProto(o.MostDisruptiveAllowedAction))
	p.SetMinReadySec(dcl.ValueOrEmptyInt64(o.MinReadySec))
	return p
}

// InstanceGroupManagerUpdatePolicyMaxSurgeToProto converts a InstanceGroupManagerUpdatePolicyMaxSurge object to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyMaxSurgeToProto(o *alpha.InstanceGroupManagerUpdatePolicyMaxSurge) *alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMaxSurge {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMaxSurge{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerUpdatePolicyMaxUnavailableToProto converts a InstanceGroupManagerUpdatePolicyMaxUnavailable object to its proto representation.
func ComputeAlphaInstanceGroupManagerUpdatePolicyMaxUnavailableToProto(o *alpha.InstanceGroupManagerUpdatePolicyMaxUnavailable) *alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMaxUnavailable {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerUpdatePolicyMaxUnavailable{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetCalculated(dcl.ValueOrEmptyInt64(o.Calculated))
	return p
}

// InstanceGroupManagerNamedPortsToProto converts a InstanceGroupManagerNamedPorts object to its proto representation.
func ComputeAlphaInstanceGroupManagerNamedPortsToProto(o *alpha.InstanceGroupManagerNamedPorts) *alphapb.ComputeAlphaInstanceGroupManagerNamedPorts {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerNamedPorts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetPort(dcl.ValueOrEmptyInt64(o.Port))
	return p
}

// InstanceGroupManagerStatefulPolicyToProto converts a InstanceGroupManagerStatefulPolicy object to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyToProto(o *alpha.InstanceGroupManagerStatefulPolicy) *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicy{}
	p.SetPreservedState(ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateToProto(o.PreservedState))
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateToProto converts a InstanceGroupManagerStatefulPolicyPreservedState object to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateToProto(o *alpha.InstanceGroupManagerStatefulPolicyPreservedState) *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedState{}
	mDisks := make(map[string]*alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks, len(o.Disks))
	for k, r := range o.Disks {
		mDisks[k] = ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(&r)
	}
	p.SetDisks(mDisks)
	mInternalIps := make(map[string]*alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps, len(o.InternalIps))
	for k, r := range o.InternalIps {
		mInternalIps[k] = ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsToProto(&r)
	}
	p.SetInternalIps(mInternalIps)
	mExternalIps := make(map[string]*alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps, len(o.ExternalIps))
	for k, r := range o.ExternalIps {
		mExternalIps[k] = ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsToProto(&r)
	}
	p.SetExternalIps(mExternalIps)
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateDisksToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateDisks object to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksToProto(o *alpha.InstanceGroupManagerStatefulPolicyPreservedStateDisks) *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisks{}
	p.SetAutoDelete(ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumToProto(o.AutoDelete))
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateInternalIps object to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsToProto(o *alpha.InstanceGroupManagerStatefulPolicyPreservedStateInternalIps) *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIps{}
	p.SetAutoDelete(ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnumToProto(o.AutoDelete))
	return p
}

// InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsToProto converts a InstanceGroupManagerStatefulPolicyPreservedStateExternalIps object to its proto representation.
func ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsToProto(o *alpha.InstanceGroupManagerStatefulPolicyPreservedStateExternalIps) *alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps {
	if o == nil {
		return nil
	}
	p := &alphapb.ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIps{}
	p.SetAutoDelete(ComputeAlphaInstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnumToProto(o.AutoDelete))
	return p
}

// InstanceGroupManagerToProto converts a InstanceGroupManager resource to its proto representation.
func InstanceGroupManagerToProto(resource *alpha.InstanceGroupManager) *alphapb.ComputeAlphaInstanceGroupManager {
	p := &alphapb.ComputeAlphaInstanceGroupManager{}
	p.SetId(dcl.ValueOrEmptyInt64(resource.Id))
	p.SetCreationTimestamp(dcl.ValueOrEmptyString(resource.CreationTimestamp))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetZone(dcl.ValueOrEmptyString(resource.Zone))
	p.SetRegion(dcl.ValueOrEmptyString(resource.Region))
	p.SetDistributionPolicy(ComputeAlphaInstanceGroupManagerDistributionPolicyToProto(resource.DistributionPolicy))
	p.SetInstanceTemplate(dcl.ValueOrEmptyString(resource.InstanceTemplate))
	p.SetInstanceGroup(dcl.ValueOrEmptyString(resource.InstanceGroup))
	p.SetBaseInstanceName(dcl.ValueOrEmptyString(resource.BaseInstanceName))
	p.SetFingerprint(dcl.ValueOrEmptyString(resource.Fingerprint))
	p.SetCurrentActions(ComputeAlphaInstanceGroupManagerCurrentActionsToProto(resource.CurrentActions))
	p.SetStatus(ComputeAlphaInstanceGroupManagerStatusToProto(resource.Status))
	p.SetTargetSize(dcl.ValueOrEmptyInt64(resource.TargetSize))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetUpdatePolicy(ComputeAlphaInstanceGroupManagerUpdatePolicyToProto(resource.UpdatePolicy))
	p.SetStatefulPolicy(ComputeAlphaInstanceGroupManagerStatefulPolicyToProto(resource.StatefulPolicy))
	p.SetServiceAccount(dcl.ValueOrEmptyString(resource.ServiceAccount))
	p.SetFailoverAction(ComputeAlphaInstanceGroupManagerFailoverActionEnumToProto(resource.FailoverAction))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sVersions := make([]*alphapb.ComputeAlphaInstanceGroupManagerVersions, len(resource.Versions))
	for i, r := range resource.Versions {
		sVersions[i] = ComputeAlphaInstanceGroupManagerVersionsToProto(&r)
	}
	p.SetVersions(sVersions)
	sTargetPools := make([]string, len(resource.TargetPools))
	for i, r := range resource.TargetPools {
		sTargetPools[i] = r
	}
	p.SetTargetPools(sTargetPools)
	sAutoHealingPolicies := make([]*alphapb.ComputeAlphaInstanceGroupManagerAutoHealingPolicies, len(resource.AutoHealingPolicies))
	for i, r := range resource.AutoHealingPolicies {
		sAutoHealingPolicies[i] = ComputeAlphaInstanceGroupManagerAutoHealingPoliciesToProto(&r)
	}
	p.SetAutoHealingPolicies(sAutoHealingPolicies)
	sNamedPorts := make([]*alphapb.ComputeAlphaInstanceGroupManagerNamedPorts, len(resource.NamedPorts))
	for i, r := range resource.NamedPorts {
		sNamedPorts[i] = ComputeAlphaInstanceGroupManagerNamedPortsToProto(&r)
	}
	p.SetNamedPorts(sNamedPorts)

	return p
}

// applyInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) applyInstanceGroupManager(ctx context.Context, c *alpha.Client, request *alphapb.ApplyComputeAlphaInstanceGroupManagerRequest) (*alphapb.ComputeAlphaInstanceGroupManager, error) {
	p := ProtoToInstanceGroupManager(request.GetResource())
	res, err := c.ApplyInstanceGroupManager(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceGroupManagerToProto(res)
	return r, nil
}

// applyComputeAlphaInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Apply() method.
func (s *InstanceGroupManagerServer) ApplyComputeAlphaInstanceGroupManager(ctx context.Context, request *alphapb.ApplyComputeAlphaInstanceGroupManagerRequest) (*alphapb.ComputeAlphaInstanceGroupManager, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstanceGroupManager(ctx, cl, request)
}

// DeleteInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManager Delete() method.
func (s *InstanceGroupManagerServer) DeleteComputeAlphaInstanceGroupManager(ctx context.Context, request *alphapb.DeleteComputeAlphaInstanceGroupManagerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstanceGroupManager(ctx, ProtoToInstanceGroupManager(request.GetResource()))

}

// ListComputeAlphaInstanceGroupManager handles the gRPC request by passing it to the underlying InstanceGroupManagerList() method.
func (s *InstanceGroupManagerServer) ListComputeAlphaInstanceGroupManager(ctx context.Context, request *alphapb.ListComputeAlphaInstanceGroupManagerRequest) (*alphapb.ListComputeAlphaInstanceGroupManagerResponse, error) {
	cl, err := createConfigInstanceGroupManager(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstanceGroupManager(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ComputeAlphaInstanceGroupManager
	for _, r := range resources.Items {
		rp := InstanceGroupManagerToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListComputeAlphaInstanceGroupManagerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstanceGroupManager(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
