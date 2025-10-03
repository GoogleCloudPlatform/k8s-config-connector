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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/osconfig/beta/osconfig_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/beta"
)

// OSPolicyAssignmentServer implements the gRPC interface for OSPolicyAssignment.
type OSPolicyAssignmentServer struct{}

// ProtoToOSPolicyAssignmentOSPoliciesModeEnum converts a OSPolicyAssignmentOSPoliciesModeEnum enum from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum(e betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum) *beta.OSPolicyAssignmentOSPoliciesModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum_name[int32(e)]; ok {
		e := beta.OSPolicyAssignmentOSPoliciesModeEnum(n[len("OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(e betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum_name[int32(e)]; ok {
		e := beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(n[len("OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(e betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_name[int32(e)]; ok {
		e := beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(n[len("OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(e betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum_name[int32(e)]; ok {
		e := beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(n[len("OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(e betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum_name[int32(e)]; ok {
		e := beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(n[len("OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum enum from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(e betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum_name[int32(e)]; ok {
		e := beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(n[len("OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentRolloutStateEnum converts a OSPolicyAssignmentRolloutStateEnum enum from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentRolloutStateEnum(e betapb.OsconfigBetaOSPolicyAssignmentRolloutStateEnum) *beta.OSPolicyAssignmentRolloutStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaOSPolicyAssignmentRolloutStateEnum_name[int32(e)]; ok {
		e := beta.OSPolicyAssignmentRolloutStateEnum(n[len("OsconfigBetaOSPolicyAssignmentRolloutStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPolicies converts a OSPolicyAssignmentOSPolicies object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPolicies(p *betapb.OsconfigBetaOSPolicyAssignmentOSPolicies) *beta.OSPolicyAssignmentOSPolicies {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPolicies{
		Id:                        dcl.StringOrNil(p.GetId()),
		Description:               dcl.StringOrNil(p.GetDescription()),
		Mode:                      ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum(p.GetMode()),
		AllowNoResourceGroupMatch: dcl.Bool(p.GetAllowNoResourceGroupMatch()),
	}
	for _, r := range p.GetResourceGroups() {
		obj.ResourceGroups = append(obj.ResourceGroups, *ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroups(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroups converts a OSPolicyAssignmentOSPoliciesResourceGroups object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroups(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroups) *beta.OSPolicyAssignmentOSPoliciesResourceGroups {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroups{}
	for _, r := range p.GetInventoryFilters() {
		obj.InventoryFilters = append(obj.InventoryFilters, *ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters(r))
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResources(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters converts a OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters{
		OSShortName: dcl.StringOrNil(p.GetOsShortName()),
		OSVersion:   dcl.StringOrNil(p.GetOsVersion()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResources converts a OSPolicyAssignmentOSPoliciesResourceGroupsResources object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResources(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResources) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResources {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResources{
		Id:         dcl.StringOrNil(p.GetId()),
		Pkg:        ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(p.GetPkg()),
		Repository: ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(p.GetRepository()),
		Exec:       ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(p.GetExec()),
		File:       ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(p.GetFile()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{
		DesiredState: ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(p.GetDesiredState()),
		Apt:          ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(p.GetApt()),
		Deb:          ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(p.GetDeb()),
		Yum:          ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(p.GetYum()),
		Zypper:       ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(p.GetZypper()),
		Rpm:          ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(p.GetRpm()),
		Googet:       ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(p.GetGooget()),
		Msi:          ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(p.GetMsi()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{
		Source:   ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(p.GetSource()),
		PullDeps: dcl.Bool(p.GetPullDeps()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource{
		Remote:        ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{
		Source:   ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource(p.GetSource()),
		PullDeps: dcl.Bool(p.GetPullDeps()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource{
		Remote:        ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{
		Source: ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource(p.GetSource()),
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource{
		Remote:        ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{
		Apt:    ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(p.GetApt()),
		Yum:    ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(p.GetYum()),
		Zypper: ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{
		ArchiveType:  ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(p.GetArchiveType()),
		Uri:          dcl.StringOrNil(p.GetUri()),
		Distribution: dcl.StringOrNil(p.GetDistribution()),
		GpgKey:       dcl.StringOrNil(p.GetGpgKey()),
	}
	for _, r := range p.GetComponents() {
		obj.Components = append(obj.Components, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{
		Id:          dcl.StringOrNil(p.GetId()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		BaseUrl:     dcl.StringOrNil(p.GetBaseUrl()),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{
		Id:          dcl.StringOrNil(p.GetId()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		BaseUrl:     dcl.StringOrNil(p.GetBaseUrl()),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{
		Name: dcl.StringOrNil(p.GetName()),
		Url:  dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{
		Validate: ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate(p.GetValidate()),
		Enforce:  ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce(p.GetEnforce()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate{
		File:           ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile(p.GetFile()),
		Script:         dcl.StringOrNil(p.GetScript()),
		Interpreter:    ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(p.GetInterpreter()),
		OutputFilePath: dcl.StringOrNil(p.GetOutputFilePath()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile{
		Remote:        ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce{
		File:           ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile(p.GetFile()),
		Script:         dcl.StringOrNil(p.GetScript()),
		Interpreter:    ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(p.GetInterpreter()),
		OutputFilePath: dcl.StringOrNil(p.GetOutputFilePath()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile{
		Remote:        ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{
		File:        ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile(p.GetFile()),
		Content:     dcl.StringOrNil(p.GetContent()),
		Path:        dcl.StringOrNil(p.GetPath()),
		State:       ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(p.GetState()),
		Permissions: dcl.StringOrNil(p.GetPermissions()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile{
		Remote:        ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs(p *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs) *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilter converts a OSPolicyAssignmentInstanceFilter object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentInstanceFilter(p *betapb.OsconfigBetaOSPolicyAssignmentInstanceFilter) *beta.OSPolicyAssignmentInstanceFilter {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentInstanceFilter{
		All: dcl.Bool(p.GetAll()),
	}
	for _, r := range p.GetInclusionLabels() {
		obj.InclusionLabels = append(obj.InclusionLabels, *ProtoToOsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabels(r))
	}
	for _, r := range p.GetExclusionLabels() {
		obj.ExclusionLabels = append(obj.ExclusionLabels, *ProtoToOsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabels(r))
	}
	for _, r := range p.GetInventories() {
		obj.Inventories = append(obj.Inventories, *ProtoToOsconfigBetaOSPolicyAssignmentInstanceFilterInventories(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterInclusionLabels converts a OSPolicyAssignmentInstanceFilterInclusionLabels object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabels(p *betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabels) *beta.OSPolicyAssignmentInstanceFilterInclusionLabels {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentInstanceFilterInclusionLabels{}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterExclusionLabels converts a OSPolicyAssignmentInstanceFilterExclusionLabels object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabels(p *betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabels) *beta.OSPolicyAssignmentInstanceFilterExclusionLabels {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentInstanceFilterExclusionLabels{}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterInventories converts a OSPolicyAssignmentInstanceFilterInventories object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentInstanceFilterInventories(p *betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterInventories) *beta.OSPolicyAssignmentInstanceFilterInventories {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentInstanceFilterInventories{
		OSShortName: dcl.StringOrNil(p.GetOsShortName()),
		OSVersion:   dcl.StringOrNil(p.GetOsVersion()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentRollout converts a OSPolicyAssignmentRollout object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentRollout(p *betapb.OsconfigBetaOSPolicyAssignmentRollout) *beta.OSPolicyAssignmentRollout {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentRollout{
		DisruptionBudget: ProtoToOsconfigBetaOSPolicyAssignmentRolloutDisruptionBudget(p.GetDisruptionBudget()),
		MinWaitDuration:  dcl.StringOrNil(p.GetMinWaitDuration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentRolloutDisruptionBudget converts a OSPolicyAssignmentRolloutDisruptionBudget object from its proto representation.
func ProtoToOsconfigBetaOSPolicyAssignmentRolloutDisruptionBudget(p *betapb.OsconfigBetaOSPolicyAssignmentRolloutDisruptionBudget) *beta.OSPolicyAssignmentRolloutDisruptionBudget {
	if p == nil {
		return nil
	}
	obj := &beta.OSPolicyAssignmentRolloutDisruptionBudget{
		Fixed:   dcl.Int64OrNil(p.GetFixed()),
		Percent: dcl.Int64OrNil(p.GetPercent()),
	}
	return obj
}

// ProtoToOSPolicyAssignment converts a OSPolicyAssignment resource from its proto representation.
func ProtoToOSPolicyAssignment(p *betapb.OsconfigBetaOSPolicyAssignment) *beta.OSPolicyAssignment {
	obj := &beta.OSPolicyAssignment{
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		InstanceFilter:     ProtoToOsconfigBetaOSPolicyAssignmentInstanceFilter(p.GetInstanceFilter()),
		Rollout:            ProtoToOsconfigBetaOSPolicyAssignmentRollout(p.GetRollout()),
		RevisionId:         dcl.StringOrNil(p.GetRevisionId()),
		RevisionCreateTime: dcl.StringOrNil(p.GetRevisionCreateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		RolloutState:       ProtoToOsconfigBetaOSPolicyAssignmentRolloutStateEnum(p.GetRolloutState()),
		Baseline:           dcl.Bool(p.GetBaseline()),
		Deleted:            dcl.Bool(p.GetDeleted()),
		Reconciling:        dcl.Bool(p.GetReconciling()),
		Uid:                dcl.StringOrNil(p.GetUid()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		SkipAwaitRollout:   dcl.Bool(p.GetSkipAwaitRollout()),
	}
	for _, r := range p.GetOsPolicies() {
		obj.OSPolicies = append(obj.OSPolicies, *ProtoToOsconfigBetaOSPolicyAssignmentOSPolicies(r))
	}
	return obj
}

// OSPolicyAssignmentOSPoliciesModeEnumToProto converts a OSPolicyAssignmentOSPoliciesModeEnum enum to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnumToProto(e *beta.OSPolicyAssignmentOSPoliciesModeEnum) betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum {
	if e == nil {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum_value["OSPolicyAssignmentOSPoliciesModeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum(v)
	}
	return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(e *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum) betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == nil {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
	}
	if v, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(v)
	}
	return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(e *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == nil {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(v)
	}
	return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum enum to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto(e *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum) betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(v)
	}
	return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum enum to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto(e *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum) betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(v)
	}
	return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum enum to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto(e *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum) betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum {
	if e == nil {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(0)
	}
	if v, ok := betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(v)
	}
	return betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(0)
}

// OSPolicyAssignmentRolloutStateEnumToProto converts a OSPolicyAssignmentRolloutStateEnum enum to its proto representation.
func OsconfigBetaOSPolicyAssignmentRolloutStateEnumToProto(e *beta.OSPolicyAssignmentRolloutStateEnum) betapb.OsconfigBetaOSPolicyAssignmentRolloutStateEnum {
	if e == nil {
		return betapb.OsconfigBetaOSPolicyAssignmentRolloutStateEnum(0)
	}
	if v, ok := betapb.OsconfigBetaOSPolicyAssignmentRolloutStateEnum_value["OSPolicyAssignmentRolloutStateEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaOSPolicyAssignmentRolloutStateEnum(v)
	}
	return betapb.OsconfigBetaOSPolicyAssignmentRolloutStateEnum(0)
}

// OSPolicyAssignmentOSPoliciesToProto converts a OSPolicyAssignmentOSPolicies object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesToProto(o *beta.OSPolicyAssignmentOSPolicies) *betapb.OsconfigBetaOSPolicyAssignmentOSPolicies {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPolicies{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetMode(OsconfigBetaOSPolicyAssignmentOSPoliciesModeEnumToProto(o.Mode))
	p.SetAllowNoResourceGroupMatch(dcl.ValueOrEmptyBool(o.AllowNoResourceGroupMatch))
	sResourceGroups := make([]*betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroups, len(o.ResourceGroups))
	for i, r := range o.ResourceGroups {
		sResourceGroups[i] = OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsToProto(&r)
	}
	p.SetResourceGroups(sResourceGroups)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroups object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroups) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroups {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroups{}
	sInventoryFilters := make([]*betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters, len(o.InventoryFilters))
	for i, r := range o.InventoryFilters {
		sInventoryFilters[i] = OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto(&r)
	}
	p.SetInventoryFilters(sInventoryFilters)
	sResources := make([]*betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResources, len(o.Resources))
	for i, r := range o.Resources {
		sResources[i] = OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto(&r)
	}
	p.SetResources(sResources)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters{}
	p.SetOsShortName(dcl.ValueOrEmptyString(o.OSShortName))
	p.SetOsVersion(dcl.ValueOrEmptyString(o.OSVersion))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResources object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResources) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResources {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResources{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetPkg(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto(o.Pkg))
	p.SetRepository(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto(o.Repository))
	p.SetExec(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto(o.Exec))
	p.SetFile(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto(o.File))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{}
	p.SetDesiredState(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(o.DesiredState))
	p.SetApt(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto(o.Apt))
	p.SetDeb(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto(o.Deb))
	p.SetYum(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto(o.Yum))
	p.SetZypper(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto(o.Zypper))
	p.SetRpm(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto(o.Rpm))
	p.SetGooget(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto(o.Googet))
	p.SetMsi(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto(o.Msi))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{}
	p.SetSource(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto(o.Source))
	p.SetPullDeps(dcl.ValueOrEmptyBool(o.PullDeps))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource{}
	p.SetRemote(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{}
	p.SetSource(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto(o.Source))
	p.SetPullDeps(dcl.ValueOrEmptyBool(o.PullDeps))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource{}
	p.SetRemote(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{}
	p.SetSource(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto(o.Source))
	sProperties := make([]string, len(o.Properties))
	for i, r := range o.Properties {
		sProperties[i] = r
	}
	p.SetProperties(sProperties)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource{}
	p.SetRemote(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{}
	p.SetApt(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto(o.Apt))
	p.SetYum(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto(o.Yum))
	p.SetZypper(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto(o.Zypper))
	p.SetGoo(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto(o.Goo))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{}
	p.SetArchiveType(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(o.ArchiveType))
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetDistribution(dcl.ValueOrEmptyString(o.Distribution))
	p.SetGpgKey(dcl.ValueOrEmptyString(o.GpgKey))
	sComponents := make([]string, len(o.Components))
	for i, r := range o.Components {
		sComponents[i] = r
	}
	p.SetComponents(sComponents)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetBaseUrl(dcl.ValueOrEmptyString(o.BaseUrl))
	sGpgKeys := make([]string, len(o.GpgKeys))
	for i, r := range o.GpgKeys {
		sGpgKeys[i] = r
	}
	p.SetGpgKeys(sGpgKeys)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetBaseUrl(dcl.ValueOrEmptyString(o.BaseUrl))
	sGpgKeys := make([]string, len(o.GpgKeys))
	for i, r := range o.GpgKeys {
		sGpgKeys[i] = r
	}
	p.SetGpgKeys(sGpgKeys)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{}
	p.SetValidate(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto(o.Validate))
	p.SetEnforce(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto(o.Enforce))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate{}
	p.SetFile(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto(o.File))
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto(o.Interpreter))
	p.SetOutputFilePath(dcl.ValueOrEmptyString(o.OutputFilePath))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile{}
	p.SetRemote(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce{}
	p.SetFile(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto(o.File))
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto(o.Interpreter))
	p.SetOutputFilePath(dcl.ValueOrEmptyString(o.OutputFilePath))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile{}
	p.SetRemote(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{}
	p.SetFile(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto(o.File))
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetState(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto(o.State))
	p.SetPermissions(dcl.ValueOrEmptyString(o.Permissions))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile{}
	p.SetRemote(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs object to its proto representation.
func OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto(o *beta.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs) *betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentInstanceFilterToProto converts a OSPolicyAssignmentInstanceFilter object to its proto representation.
func OsconfigBetaOSPolicyAssignmentInstanceFilterToProto(o *beta.OSPolicyAssignmentInstanceFilter) *betapb.OsconfigBetaOSPolicyAssignmentInstanceFilter {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentInstanceFilter{}
	p.SetAll(dcl.ValueOrEmptyBool(o.All))
	sInclusionLabels := make([]*betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabels, len(o.InclusionLabels))
	for i, r := range o.InclusionLabels {
		sInclusionLabels[i] = OsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabelsToProto(&r)
	}
	p.SetInclusionLabels(sInclusionLabels)
	sExclusionLabels := make([]*betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabels, len(o.ExclusionLabels))
	for i, r := range o.ExclusionLabels {
		sExclusionLabels[i] = OsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabelsToProto(&r)
	}
	p.SetExclusionLabels(sExclusionLabels)
	sInventories := make([]*betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterInventories, len(o.Inventories))
	for i, r := range o.Inventories {
		sInventories[i] = OsconfigBetaOSPolicyAssignmentInstanceFilterInventoriesToProto(&r)
	}
	p.SetInventories(sInventories)
	return p
}

// OSPolicyAssignmentInstanceFilterInclusionLabelsToProto converts a OSPolicyAssignmentInstanceFilterInclusionLabels object to its proto representation.
func OsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabelsToProto(o *beta.OSPolicyAssignmentInstanceFilterInclusionLabels) *betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabels {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterInclusionLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// OSPolicyAssignmentInstanceFilterExclusionLabelsToProto converts a OSPolicyAssignmentInstanceFilterExclusionLabels object to its proto representation.
func OsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabelsToProto(o *beta.OSPolicyAssignmentInstanceFilterExclusionLabels) *betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabels {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterExclusionLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// OSPolicyAssignmentInstanceFilterInventoriesToProto converts a OSPolicyAssignmentInstanceFilterInventories object to its proto representation.
func OsconfigBetaOSPolicyAssignmentInstanceFilterInventoriesToProto(o *beta.OSPolicyAssignmentInstanceFilterInventories) *betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterInventories {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentInstanceFilterInventories{}
	p.SetOsShortName(dcl.ValueOrEmptyString(o.OSShortName))
	p.SetOsVersion(dcl.ValueOrEmptyString(o.OSVersion))
	return p
}

// OSPolicyAssignmentRolloutToProto converts a OSPolicyAssignmentRollout object to its proto representation.
func OsconfigBetaOSPolicyAssignmentRolloutToProto(o *beta.OSPolicyAssignmentRollout) *betapb.OsconfigBetaOSPolicyAssignmentRollout {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentRollout{}
	p.SetDisruptionBudget(OsconfigBetaOSPolicyAssignmentRolloutDisruptionBudgetToProto(o.DisruptionBudget))
	p.SetMinWaitDuration(dcl.ValueOrEmptyString(o.MinWaitDuration))
	return p
}

// OSPolicyAssignmentRolloutDisruptionBudgetToProto converts a OSPolicyAssignmentRolloutDisruptionBudget object to its proto representation.
func OsconfigBetaOSPolicyAssignmentRolloutDisruptionBudgetToProto(o *beta.OSPolicyAssignmentRolloutDisruptionBudget) *betapb.OsconfigBetaOSPolicyAssignmentRolloutDisruptionBudget {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaOSPolicyAssignmentRolloutDisruptionBudget{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	return p
}

// OSPolicyAssignmentToProto converts a OSPolicyAssignment resource to its proto representation.
func OSPolicyAssignmentToProto(resource *beta.OSPolicyAssignment) *betapb.OsconfigBetaOSPolicyAssignment {
	p := &betapb.OsconfigBetaOSPolicyAssignment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInstanceFilter(OsconfigBetaOSPolicyAssignmentInstanceFilterToProto(resource.InstanceFilter))
	p.SetRollout(OsconfigBetaOSPolicyAssignmentRolloutToProto(resource.Rollout))
	p.SetRevisionId(dcl.ValueOrEmptyString(resource.RevisionId))
	p.SetRevisionCreateTime(dcl.ValueOrEmptyString(resource.RevisionCreateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetRolloutState(OsconfigBetaOSPolicyAssignmentRolloutStateEnumToProto(resource.RolloutState))
	p.SetBaseline(dcl.ValueOrEmptyBool(resource.Baseline))
	p.SetDeleted(dcl.ValueOrEmptyBool(resource.Deleted))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSkipAwaitRollout(dcl.ValueOrEmptyBool(resource.SkipAwaitRollout))
	sOSPolicies := make([]*betapb.OsconfigBetaOSPolicyAssignmentOSPolicies, len(resource.OSPolicies))
	for i, r := range resource.OSPolicies {
		sOSPolicies[i] = OsconfigBetaOSPolicyAssignmentOSPoliciesToProto(&r)
	}
	p.SetOsPolicies(sOSPolicies)

	return p
}

// applyOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Apply() method.
func (s *OSPolicyAssignmentServer) applyOSPolicyAssignment(ctx context.Context, c *beta.Client, request *betapb.ApplyOsconfigBetaOSPolicyAssignmentRequest) (*betapb.OsconfigBetaOSPolicyAssignment, error) {
	p := ProtoToOSPolicyAssignment(request.GetResource())
	res, err := c.ApplyOSPolicyAssignment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OSPolicyAssignmentToProto(res)
	return r, nil
}

// applyOsconfigBetaOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Apply() method.
func (s *OSPolicyAssignmentServer) ApplyOsconfigBetaOSPolicyAssignment(ctx context.Context, request *betapb.ApplyOsconfigBetaOSPolicyAssignmentRequest) (*betapb.OsconfigBetaOSPolicyAssignment, error) {
	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyOSPolicyAssignment(ctx, cl, request)
}

// DeleteOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Delete() method.
func (s *OSPolicyAssignmentServer) DeleteOsconfigBetaOSPolicyAssignment(ctx context.Context, request *betapb.DeleteOsconfigBetaOSPolicyAssignmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOSPolicyAssignment(ctx, ProtoToOSPolicyAssignment(request.GetResource()))

}

// ListOsconfigBetaOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignmentList() method.
func (s *OSPolicyAssignmentServer) ListOsconfigBetaOSPolicyAssignment(ctx context.Context, request *betapb.ListOsconfigBetaOSPolicyAssignmentRequest) (*betapb.ListOsconfigBetaOSPolicyAssignmentResponse, error) {
	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOSPolicyAssignment(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.OsconfigBetaOSPolicyAssignment
	for _, r := range resources.Items {
		rp := OSPolicyAssignmentToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListOsconfigBetaOSPolicyAssignmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigOSPolicyAssignment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
