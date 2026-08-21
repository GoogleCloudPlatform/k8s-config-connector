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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/osconfig/alpha/osconfig_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha"
)

// OSPolicyAssignmentServer implements the gRPC interface for OSPolicyAssignment.
type OSPolicyAssignmentServer struct{}

// ProtoToOSPolicyAssignmentOSPoliciesModeEnum converts a OSPolicyAssignmentOSPoliciesModeEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum) *alpha.OSPolicyAssignmentOSPoliciesModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesModeEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(n[len("OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentRolloutStateEnum converts a OSPolicyAssignmentRolloutStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentRolloutStateEnum(e alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum) *alpha.OSPolicyAssignmentRolloutStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum_name[int32(e)]; ok {
		e := alpha.OSPolicyAssignmentRolloutStateEnum(n[len("OsconfigAlphaOSPolicyAssignmentRolloutStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPolicies converts a OSPolicyAssignmentOSPolicies object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPolicies(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPolicies) *alpha.OSPolicyAssignmentOSPolicies {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPolicies{
		Id:                        dcl.StringOrNil(p.GetId()),
		Description:               dcl.StringOrNil(p.GetDescription()),
		Mode:                      ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(p.GetMode()),
		AllowNoResourceGroupMatch: dcl.Bool(p.GetAllowNoResourceGroupMatch()),
	}
	for _, r := range p.GetResourceGroups() {
		obj.ResourceGroups = append(obj.ResourceGroups, *ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroups converts a OSPolicyAssignmentOSPoliciesResourceGroups object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups) *alpha.OSPolicyAssignmentOSPoliciesResourceGroups {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroups{}
	for _, r := range p.GetInventoryFilters() {
		obj.InventoryFilters = append(obj.InventoryFilters, *ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters(r))
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters converts a OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters{
		OSShortName: dcl.StringOrNil(p.GetOsShortName()),
		OSVersion:   dcl.StringOrNil(p.GetOsVersion()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResources converts a OSPolicyAssignmentOSPoliciesResourceGroupsResources object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResources {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResources{
		Id:         dcl.StringOrNil(p.GetId()),
		Pkg:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(p.GetPkg()),
		Repository: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(p.GetRepository()),
		Exec:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(p.GetExec()),
		File:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(p.GetFile()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{
		DesiredState: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(p.GetDesiredState()),
		Apt:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(p.GetApt()),
		Deb:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(p.GetDeb()),
		Yum:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(p.GetYum()),
		Zypper:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(p.GetZypper()),
		Rpm:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(p.GetRpm()),
		Googet:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(p.GetGooget()),
		Msi:          ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(p.GetMsi()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{
		Source:   ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(p.GetSource()),
		PullDeps: dcl.Bool(p.GetPullDeps()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource{
		Remote:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{
		Source:   ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource(p.GetSource()),
		PullDeps: dcl.Bool(p.GetPullDeps()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource{
		Remote:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{
		Source: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource(p.GetSource()),
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource{
		Remote:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{
		Apt:    ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(p.GetApt()),
		Yum:    ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(p.GetYum()),
		Zypper: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{
		ArchiveType:  ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(p.GetArchiveType()),
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
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{
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
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{
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
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{
		Name: dcl.StringOrNil(p.GetName()),
		Url:  dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{
		Validate: ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate(p.GetValidate()),
		Enforce:  ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce(p.GetEnforce()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate{
		File:           ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile(p.GetFile()),
		Script:         dcl.StringOrNil(p.GetScript()),
		Interpreter:    ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(p.GetInterpreter()),
		OutputFilePath: dcl.StringOrNil(p.GetOutputFilePath()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile{
		Remote:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce{
		File:           ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile(p.GetFile()),
		Script:         dcl.StringOrNil(p.GetScript()),
		Interpreter:    ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(p.GetInterpreter()),
		OutputFilePath: dcl.StringOrNil(p.GetOutputFilePath()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile{
		Remote:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{
		File:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile(p.GetFile()),
		Content:     dcl.StringOrNil(p.GetContent()),
		Path:        dcl.StringOrNil(p.GetPath()),
		State:       ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(p.GetState()),
		Permissions: dcl.StringOrNil(p.GetPermissions()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile{
		Remote:        ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs(p *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs) *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilter converts a OSPolicyAssignmentInstanceFilter object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilter(p *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilter) *alpha.OSPolicyAssignmentInstanceFilter {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentInstanceFilter{
		All: dcl.Bool(p.GetAll()),
	}
	for _, r := range p.GetInclusionLabels() {
		obj.InclusionLabels = append(obj.InclusionLabels, *ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels(r))
	}
	for _, r := range p.GetExclusionLabels() {
		obj.ExclusionLabels = append(obj.ExclusionLabels, *ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels(r))
	}
	for _, r := range p.GetInventories() {
		obj.Inventories = append(obj.Inventories, *ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterInventories(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterInclusionLabels converts a OSPolicyAssignmentInstanceFilterInclusionLabels object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels(p *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels) *alpha.OSPolicyAssignmentInstanceFilterInclusionLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentInstanceFilterInclusionLabels{}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterExclusionLabels converts a OSPolicyAssignmentInstanceFilterExclusionLabels object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels(p *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels) *alpha.OSPolicyAssignmentInstanceFilterExclusionLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentInstanceFilterExclusionLabels{}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterInventories converts a OSPolicyAssignmentInstanceFilterInventories object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilterInventories(p *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInventories) *alpha.OSPolicyAssignmentInstanceFilterInventories {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentInstanceFilterInventories{
		OSShortName: dcl.StringOrNil(p.GetOsShortName()),
		OSVersion:   dcl.StringOrNil(p.GetOsVersion()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentRollout converts a OSPolicyAssignmentRollout object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentRollout(p *alphapb.OsconfigAlphaOSPolicyAssignmentRollout) *alpha.OSPolicyAssignmentRollout {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentRollout{
		DisruptionBudget: ProtoToOsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget(p.GetDisruptionBudget()),
		MinWaitDuration:  dcl.StringOrNil(p.GetMinWaitDuration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentRolloutDisruptionBudget converts a OSPolicyAssignmentRolloutDisruptionBudget object from its proto representation.
func ProtoToOsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget(p *alphapb.OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget) *alpha.OSPolicyAssignmentRolloutDisruptionBudget {
	if p == nil {
		return nil
	}
	obj := &alpha.OSPolicyAssignmentRolloutDisruptionBudget{
		Fixed:   dcl.Int64OrNil(p.GetFixed()),
		Percent: dcl.Int64OrNil(p.GetPercent()),
	}
	return obj
}

// ProtoToOSPolicyAssignment converts a OSPolicyAssignment resource from its proto representation.
func ProtoToOSPolicyAssignment(p *alphapb.OsconfigAlphaOSPolicyAssignment) *alpha.OSPolicyAssignment {
	obj := &alpha.OSPolicyAssignment{
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		InstanceFilter:     ProtoToOsconfigAlphaOSPolicyAssignmentInstanceFilter(p.GetInstanceFilter()),
		Rollout:            ProtoToOsconfigAlphaOSPolicyAssignmentRollout(p.GetRollout()),
		RevisionId:         dcl.StringOrNil(p.GetRevisionId()),
		RevisionCreateTime: dcl.StringOrNil(p.GetRevisionCreateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		RolloutState:       ProtoToOsconfigAlphaOSPolicyAssignmentRolloutStateEnum(p.GetRolloutState()),
		Baseline:           dcl.Bool(p.GetBaseline()),
		Deleted:            dcl.Bool(p.GetDeleted()),
		Reconciling:        dcl.Bool(p.GetReconciling()),
		Uid:                dcl.StringOrNil(p.GetUid()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		SkipAwaitRollout:   dcl.Bool(p.GetSkipAwaitRollout()),
	}
	for _, r := range p.GetOsPolicies() {
		obj.OSPolicies = append(obj.OSPolicies, *ProtoToOsconfigAlphaOSPolicyAssignmentOSPolicies(r))
	}
	return obj
}

// OSPolicyAssignmentOSPoliciesModeEnumToProto converts a OSPolicyAssignmentOSPoliciesModeEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesModeEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum_value["OSPolicyAssignmentOSPoliciesModeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto(e *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum) alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(0)
}

// OSPolicyAssignmentRolloutStateEnumToProto converts a OSPolicyAssignmentRolloutStateEnum enum to its proto representation.
func OsconfigAlphaOSPolicyAssignmentRolloutStateEnumToProto(e *alpha.OSPolicyAssignmentRolloutStateEnum) alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum_value["OSPolicyAssignmentRolloutStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum(v)
	}
	return alphapb.OsconfigAlphaOSPolicyAssignmentRolloutStateEnum(0)
}

// OSPolicyAssignmentOSPoliciesToProto converts a OSPolicyAssignmentOSPolicies object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesToProto(o *alpha.OSPolicyAssignmentOSPolicies) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPolicies {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPolicies{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetMode(OsconfigAlphaOSPolicyAssignmentOSPoliciesModeEnumToProto(o.Mode))
	p.SetAllowNoResourceGroupMatch(dcl.ValueOrEmptyBool(o.AllowNoResourceGroupMatch))
	sResourceGroups := make([]*alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups, len(o.ResourceGroups))
	for i, r := range o.ResourceGroups {
		sResourceGroups[i] = OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsToProto(&r)
	}
	p.SetResourceGroups(sResourceGroups)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroups object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroups) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroups{}
	sInventoryFilters := make([]*alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters, len(o.InventoryFilters))
	for i, r := range o.InventoryFilters {
		sInventoryFilters[i] = OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto(&r)
	}
	p.SetInventoryFilters(sInventoryFilters)
	sResources := make([]*alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources, len(o.Resources))
	for i, r := range o.Resources {
		sResources[i] = OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto(&r)
	}
	p.SetResources(sResources)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters{}
	p.SetOsShortName(dcl.ValueOrEmptyString(o.OSShortName))
	p.SetOsVersion(dcl.ValueOrEmptyString(o.OSVersion))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResources object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResources) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResources{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetPkg(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto(o.Pkg))
	p.SetRepository(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto(o.Repository))
	p.SetExec(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto(o.Exec))
	p.SetFile(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto(o.File))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{}
	p.SetDesiredState(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(o.DesiredState))
	p.SetApt(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto(o.Apt))
	p.SetDeb(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto(o.Deb))
	p.SetYum(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto(o.Yum))
	p.SetZypper(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto(o.Zypper))
	p.SetRpm(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto(o.Rpm))
	p.SetGooget(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto(o.Googet))
	p.SetMsi(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto(o.Msi))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{}
	p.SetSource(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto(o.Source))
	p.SetPullDeps(dcl.ValueOrEmptyBool(o.PullDeps))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource{}
	p.SetRemote(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{}
	p.SetSource(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto(o.Source))
	p.SetPullDeps(dcl.ValueOrEmptyBool(o.PullDeps))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource{}
	p.SetRemote(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{}
	p.SetSource(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto(o.Source))
	sProperties := make([]string, len(o.Properties))
	for i, r := range o.Properties {
		sProperties[i] = r
	}
	p.SetProperties(sProperties)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource{}
	p.SetRemote(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{}
	p.SetApt(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto(o.Apt))
	p.SetYum(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto(o.Yum))
	p.SetZypper(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto(o.Zypper))
	p.SetGoo(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto(o.Goo))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{}
	p.SetArchiveType(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(o.ArchiveType))
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
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{}
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
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{}
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
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{}
	p.SetValidate(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto(o.Validate))
	p.SetEnforce(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto(o.Enforce))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate{}
	p.SetFile(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto(o.File))
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto(o.Interpreter))
	p.SetOutputFilePath(dcl.ValueOrEmptyString(o.OutputFilePath))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile{}
	p.SetRemote(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce{}
	p.SetFile(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto(o.File))
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto(o.Interpreter))
	p.SetOutputFilePath(dcl.ValueOrEmptyString(o.OutputFilePath))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile{}
	p.SetRemote(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{}
	p.SetFile(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto(o.File))
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetState(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto(o.State))
	p.SetPermissions(dcl.ValueOrEmptyString(o.Permissions))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile{}
	p.SetRemote(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto(o *alpha.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs) *alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentInstanceFilterToProto converts a OSPolicyAssignmentInstanceFilter object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentInstanceFilterToProto(o *alpha.OSPolicyAssignmentInstanceFilter) *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilter {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilter{}
	p.SetAll(dcl.ValueOrEmptyBool(o.All))
	sInclusionLabels := make([]*alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels, len(o.InclusionLabels))
	for i, r := range o.InclusionLabels {
		sInclusionLabels[i] = OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabelsToProto(&r)
	}
	p.SetInclusionLabels(sInclusionLabels)
	sExclusionLabels := make([]*alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels, len(o.ExclusionLabels))
	for i, r := range o.ExclusionLabels {
		sExclusionLabels[i] = OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabelsToProto(&r)
	}
	p.SetExclusionLabels(sExclusionLabels)
	sInventories := make([]*alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInventories, len(o.Inventories))
	for i, r := range o.Inventories {
		sInventories[i] = OsconfigAlphaOSPolicyAssignmentInstanceFilterInventoriesToProto(&r)
	}
	p.SetInventories(sInventories)
	return p
}

// OSPolicyAssignmentInstanceFilterInclusionLabelsToProto converts a OSPolicyAssignmentInstanceFilterInclusionLabels object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabelsToProto(o *alpha.OSPolicyAssignmentInstanceFilterInclusionLabels) *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInclusionLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// OSPolicyAssignmentInstanceFilterExclusionLabelsToProto converts a OSPolicyAssignmentInstanceFilterExclusionLabels object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabelsToProto(o *alpha.OSPolicyAssignmentInstanceFilterExclusionLabels) *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterExclusionLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// OSPolicyAssignmentInstanceFilterInventoriesToProto converts a OSPolicyAssignmentInstanceFilterInventories object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentInstanceFilterInventoriesToProto(o *alpha.OSPolicyAssignmentInstanceFilterInventories) *alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInventories {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentInstanceFilterInventories{}
	p.SetOsShortName(dcl.ValueOrEmptyString(o.OSShortName))
	p.SetOsVersion(dcl.ValueOrEmptyString(o.OSVersion))
	return p
}

// OSPolicyAssignmentRolloutToProto converts a OSPolicyAssignmentRollout object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentRolloutToProto(o *alpha.OSPolicyAssignmentRollout) *alphapb.OsconfigAlphaOSPolicyAssignmentRollout {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentRollout{}
	p.SetDisruptionBudget(OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudgetToProto(o.DisruptionBudget))
	p.SetMinWaitDuration(dcl.ValueOrEmptyString(o.MinWaitDuration))
	return p
}

// OSPolicyAssignmentRolloutDisruptionBudgetToProto converts a OSPolicyAssignmentRolloutDisruptionBudget object to its proto representation.
func OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudgetToProto(o *alpha.OSPolicyAssignmentRolloutDisruptionBudget) *alphapb.OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaOSPolicyAssignmentRolloutDisruptionBudget{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	return p
}

// OSPolicyAssignmentToProto converts a OSPolicyAssignment resource to its proto representation.
func OSPolicyAssignmentToProto(resource *alpha.OSPolicyAssignment) *alphapb.OsconfigAlphaOSPolicyAssignment {
	p := &alphapb.OsconfigAlphaOSPolicyAssignment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInstanceFilter(OsconfigAlphaOSPolicyAssignmentInstanceFilterToProto(resource.InstanceFilter))
	p.SetRollout(OsconfigAlphaOSPolicyAssignmentRolloutToProto(resource.Rollout))
	p.SetRevisionId(dcl.ValueOrEmptyString(resource.RevisionId))
	p.SetRevisionCreateTime(dcl.ValueOrEmptyString(resource.RevisionCreateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetRolloutState(OsconfigAlphaOSPolicyAssignmentRolloutStateEnumToProto(resource.RolloutState))
	p.SetBaseline(dcl.ValueOrEmptyBool(resource.Baseline))
	p.SetDeleted(dcl.ValueOrEmptyBool(resource.Deleted))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSkipAwaitRollout(dcl.ValueOrEmptyBool(resource.SkipAwaitRollout))
	sOSPolicies := make([]*alphapb.OsconfigAlphaOSPolicyAssignmentOSPolicies, len(resource.OSPolicies))
	for i, r := range resource.OSPolicies {
		sOSPolicies[i] = OsconfigAlphaOSPolicyAssignmentOSPoliciesToProto(&r)
	}
	p.SetOsPolicies(sOSPolicies)

	return p
}

// applyOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Apply() method.
func (s *OSPolicyAssignmentServer) applyOSPolicyAssignment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyOsconfigAlphaOSPolicyAssignmentRequest) (*alphapb.OsconfigAlphaOSPolicyAssignment, error) {
	p := ProtoToOSPolicyAssignment(request.GetResource())
	res, err := c.ApplyOSPolicyAssignment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OSPolicyAssignmentToProto(res)
	return r, nil
}

// applyOsconfigAlphaOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Apply() method.
func (s *OSPolicyAssignmentServer) ApplyOsconfigAlphaOSPolicyAssignment(ctx context.Context, request *alphapb.ApplyOsconfigAlphaOSPolicyAssignmentRequest) (*alphapb.OsconfigAlphaOSPolicyAssignment, error) {
	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyOSPolicyAssignment(ctx, cl, request)
}

// DeleteOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Delete() method.
func (s *OSPolicyAssignmentServer) DeleteOsconfigAlphaOSPolicyAssignment(ctx context.Context, request *alphapb.DeleteOsconfigAlphaOSPolicyAssignmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOSPolicyAssignment(ctx, ProtoToOSPolicyAssignment(request.GetResource()))

}

// ListOsconfigAlphaOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignmentList() method.
func (s *OSPolicyAssignmentServer) ListOsconfigAlphaOSPolicyAssignment(ctx context.Context, request *alphapb.ListOsconfigAlphaOSPolicyAssignmentRequest) (*alphapb.ListOsconfigAlphaOSPolicyAssignmentResponse, error) {
	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOSPolicyAssignment(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.OsconfigAlphaOSPolicyAssignment
	for _, r := range resources.Items {
		rp := OSPolicyAssignmentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListOsconfigAlphaOSPolicyAssignmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigOSPolicyAssignment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
