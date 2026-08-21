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
	osconfigpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/osconfig/osconfig_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig"
)

// OSPolicyAssignmentServer implements the gRPC interface for OSPolicyAssignment.
type OSPolicyAssignmentServer struct{}

// ProtoToOSPolicyAssignmentOSPoliciesModeEnum converts a OSPolicyAssignmentOSPoliciesModeEnum enum from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesModeEnum(e osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesModeEnum) *osconfig.OSPolicyAssignmentOSPoliciesModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesModeEnum_name[int32(e)]; ok {
		e := osconfig.OSPolicyAssignmentOSPoliciesModeEnum(n[len("OsconfigOSPolicyAssignmentOSPoliciesModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(e osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum_name[int32(e)]; ok {
		e := osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(n[len("OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(e osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_name[int32(e)]; ok {
		e := osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(n[len("OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum enum from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(e osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum_name[int32(e)]; ok {
		e := osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(n[len("OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum enum from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(e osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum_name[int32(e)]; ok {
		e := osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(n[len("OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum enum from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(e osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum_name[int32(e)]; ok {
		e := osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(n[len("OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentRolloutStateEnum converts a OSPolicyAssignmentRolloutStateEnum enum from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentRolloutStateEnum(e osconfigpb.OsconfigOSPolicyAssignmentRolloutStateEnum) *osconfig.OSPolicyAssignmentRolloutStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := osconfigpb.OsconfigOSPolicyAssignmentRolloutStateEnum_name[int32(e)]; ok {
		e := osconfig.OSPolicyAssignmentRolloutStateEnum(n[len("OsconfigOSPolicyAssignmentRolloutStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToOSPolicyAssignmentOSPolicies converts a OSPolicyAssignmentOSPolicies object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPolicies(p *osconfigpb.OsconfigOSPolicyAssignmentOSPolicies) *osconfig.OSPolicyAssignmentOSPolicies {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPolicies{
		Id:                        dcl.StringOrNil(p.GetId()),
		Description:               dcl.StringOrNil(p.GetDescription()),
		Mode:                      ProtoToOsconfigOSPolicyAssignmentOSPoliciesModeEnum(p.GetMode()),
		AllowNoResourceGroupMatch: dcl.Bool(p.GetAllowNoResourceGroupMatch()),
	}
	for _, r := range p.GetResourceGroups() {
		obj.ResourceGroups = append(obj.ResourceGroups, *ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroups(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroups converts a OSPolicyAssignmentOSPoliciesResourceGroups object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroups(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroups) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroups {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroups{}
	for _, r := range p.GetInventoryFilters() {
		obj.InventoryFilters = append(obj.InventoryFilters, *ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters(r))
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResources(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters converts a OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters{
		OSShortName: dcl.StringOrNil(p.GetOsShortName()),
		OSVersion:   dcl.StringOrNil(p.GetOsVersion()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResources converts a OSPolicyAssignmentOSPoliciesResourceGroupsResources object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResources(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResources) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResources {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResources{
		Id:         dcl.StringOrNil(p.GetId()),
		Pkg:        ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(p.GetPkg()),
		Repository: ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(p.GetRepository()),
		Exec:       ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(p.GetExec()),
		File:       ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(p.GetFile()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{
		DesiredState: ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(p.GetDesiredState()),
		Apt:          ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(p.GetApt()),
		Deb:          ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(p.GetDeb()),
		Yum:          ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(p.GetYum()),
		Zypper:       ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(p.GetZypper()),
		Rpm:          ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(p.GetRpm()),
		Googet:       ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(p.GetGooget()),
		Msi:          ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(p.GetMsi()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{
		Source:   ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(p.GetSource()),
		PullDeps: dcl.Bool(p.GetPullDeps()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource{
		Remote:        ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{
		Source:   ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource(p.GetSource()),
		PullDeps: dcl.Bool(p.GetPullDeps()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource{
		Remote:        ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{
		Name: dcl.StringOrNil(p.GetName()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{
		Source: ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource(p.GetSource()),
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource{
		Remote:        ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{
		Apt:    ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(p.GetApt()),
		Yum:    ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(p.GetYum()),
		Zypper: ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{
		ArchiveType:  ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(p.GetArchiveType()),
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
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{
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
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{
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
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{
		Name: dcl.StringOrNil(p.GetName()),
		Url:  dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{
		Validate: ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate(p.GetValidate()),
		Enforce:  ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce(p.GetEnforce()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate{
		File:           ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile(p.GetFile()),
		Script:         dcl.StringOrNil(p.GetScript()),
		Interpreter:    ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(p.GetInterpreter()),
		OutputFilePath: dcl.StringOrNil(p.GetOutputFilePath()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile{
		Remote:        ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce{
		File:           ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile(p.GetFile()),
		Script:         dcl.StringOrNil(p.GetScript()),
		Interpreter:    ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(p.GetInterpreter()),
		OutputFilePath: dcl.StringOrNil(p.GetOutputFilePath()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile{
		Remote:        ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{
		File:        ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile(p.GetFile()),
		Content:     dcl.StringOrNil(p.GetContent()),
		Path:        dcl.StringOrNil(p.GetPath()),
		State:       ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(p.GetState()),
		Permissions: dcl.StringOrNil(p.GetPermissions()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile{
		Remote:        ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs(p.GetGcs()),
		LocalPath:     dcl.StringOrNil(p.GetLocalPath()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote{
		Uri:            dcl.StringOrNil(p.GetUri()),
		Sha256Checksum: dcl.StringOrNil(p.GetSha256Checksum()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs(p *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs) *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilter converts a OSPolicyAssignmentInstanceFilter object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentInstanceFilter(p *osconfigpb.OsconfigOSPolicyAssignmentInstanceFilter) *osconfig.OSPolicyAssignmentInstanceFilter {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentInstanceFilter{
		All: dcl.Bool(p.GetAll()),
	}
	for _, r := range p.GetInclusionLabels() {
		obj.InclusionLabels = append(obj.InclusionLabels, *ProtoToOsconfigOSPolicyAssignmentInstanceFilterInclusionLabels(r))
	}
	for _, r := range p.GetExclusionLabels() {
		obj.ExclusionLabels = append(obj.ExclusionLabels, *ProtoToOsconfigOSPolicyAssignmentInstanceFilterExclusionLabels(r))
	}
	for _, r := range p.GetInventories() {
		obj.Inventories = append(obj.Inventories, *ProtoToOsconfigOSPolicyAssignmentInstanceFilterInventories(r))
	}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterInclusionLabels converts a OSPolicyAssignmentInstanceFilterInclusionLabels object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentInstanceFilterInclusionLabels(p *osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterInclusionLabels) *osconfig.OSPolicyAssignmentInstanceFilterInclusionLabels {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentInstanceFilterInclusionLabels{}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterExclusionLabels converts a OSPolicyAssignmentInstanceFilterExclusionLabels object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentInstanceFilterExclusionLabels(p *osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterExclusionLabels) *osconfig.OSPolicyAssignmentInstanceFilterExclusionLabels {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentInstanceFilterExclusionLabels{}
	return obj
}

// ProtoToOSPolicyAssignmentInstanceFilterInventories converts a OSPolicyAssignmentInstanceFilterInventories object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentInstanceFilterInventories(p *osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterInventories) *osconfig.OSPolicyAssignmentInstanceFilterInventories {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentInstanceFilterInventories{
		OSShortName: dcl.StringOrNil(p.GetOsShortName()),
		OSVersion:   dcl.StringOrNil(p.GetOsVersion()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentRollout converts a OSPolicyAssignmentRollout object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentRollout(p *osconfigpb.OsconfigOSPolicyAssignmentRollout) *osconfig.OSPolicyAssignmentRollout {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentRollout{
		DisruptionBudget: ProtoToOsconfigOSPolicyAssignmentRolloutDisruptionBudget(p.GetDisruptionBudget()),
		MinWaitDuration:  dcl.StringOrNil(p.GetMinWaitDuration()),
	}
	return obj
}

// ProtoToOSPolicyAssignmentRolloutDisruptionBudget converts a OSPolicyAssignmentRolloutDisruptionBudget object from its proto representation.
func ProtoToOsconfigOSPolicyAssignmentRolloutDisruptionBudget(p *osconfigpb.OsconfigOSPolicyAssignmentRolloutDisruptionBudget) *osconfig.OSPolicyAssignmentRolloutDisruptionBudget {
	if p == nil {
		return nil
	}
	obj := &osconfig.OSPolicyAssignmentRolloutDisruptionBudget{
		Fixed:   dcl.Int64OrNil(p.GetFixed()),
		Percent: dcl.Int64OrNil(p.GetPercent()),
	}
	return obj
}

// ProtoToOSPolicyAssignment converts a OSPolicyAssignment resource from its proto representation.
func ProtoToOSPolicyAssignment(p *osconfigpb.OsconfigOSPolicyAssignment) *osconfig.OSPolicyAssignment {
	obj := &osconfig.OSPolicyAssignment{
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		InstanceFilter:     ProtoToOsconfigOSPolicyAssignmentInstanceFilter(p.GetInstanceFilter()),
		Rollout:            ProtoToOsconfigOSPolicyAssignmentRollout(p.GetRollout()),
		RevisionId:         dcl.StringOrNil(p.GetRevisionId()),
		RevisionCreateTime: dcl.StringOrNil(p.GetRevisionCreateTime()),
		Etag:               dcl.StringOrNil(p.GetEtag()),
		RolloutState:       ProtoToOsconfigOSPolicyAssignmentRolloutStateEnum(p.GetRolloutState()),
		Baseline:           dcl.Bool(p.GetBaseline()),
		Deleted:            dcl.Bool(p.GetDeleted()),
		Reconciling:        dcl.Bool(p.GetReconciling()),
		Uid:                dcl.StringOrNil(p.GetUid()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		SkipAwaitRollout:   dcl.Bool(p.GetSkipAwaitRollout()),
	}
	for _, r := range p.GetOsPolicies() {
		obj.OSPolicies = append(obj.OSPolicies, *ProtoToOsconfigOSPolicyAssignmentOSPolicies(r))
	}
	return obj
}

// OSPolicyAssignmentOSPoliciesModeEnumToProto converts a OSPolicyAssignmentOSPoliciesModeEnum enum to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesModeEnumToProto(e *osconfig.OSPolicyAssignmentOSPoliciesModeEnum) osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesModeEnum {
	if e == nil {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesModeEnum(0)
	}
	if v, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesModeEnum_value["OSPolicyAssignmentOSPoliciesModeEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesModeEnum(v)
	}
	return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesModeEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum enum to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(e *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum) osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum {
	if e == nil {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
	}
	if v, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(v)
	}
	return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum enum to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(e *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum) osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum {
	if e == nil {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
	}
	if v, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(v)
	}
	return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum enum to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto(e *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum) osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum {
	if e == nil {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(0)
	}
	if v, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(v)
	}
	return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum enum to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto(e *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum) osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum {
	if e == nil {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(0)
	}
	if v, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(v)
	}
	return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnum(0)
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum enum to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto(e *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum) osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum {
	if e == nil {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(0)
	}
	if v, ok := osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum_value["OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(v)
	}
	return osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnum(0)
}

// OSPolicyAssignmentRolloutStateEnumToProto converts a OSPolicyAssignmentRolloutStateEnum enum to its proto representation.
func OsconfigOSPolicyAssignmentRolloutStateEnumToProto(e *osconfig.OSPolicyAssignmentRolloutStateEnum) osconfigpb.OsconfigOSPolicyAssignmentRolloutStateEnum {
	if e == nil {
		return osconfigpb.OsconfigOSPolicyAssignmentRolloutStateEnum(0)
	}
	if v, ok := osconfigpb.OsconfigOSPolicyAssignmentRolloutStateEnum_value["OSPolicyAssignmentRolloutStateEnum"+string(*e)]; ok {
		return osconfigpb.OsconfigOSPolicyAssignmentRolloutStateEnum(v)
	}
	return osconfigpb.OsconfigOSPolicyAssignmentRolloutStateEnum(0)
}

// OSPolicyAssignmentOSPoliciesToProto converts a OSPolicyAssignmentOSPolicies object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesToProto(o *osconfig.OSPolicyAssignmentOSPolicies) *osconfigpb.OsconfigOSPolicyAssignmentOSPolicies {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPolicies{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetMode(OsconfigOSPolicyAssignmentOSPoliciesModeEnumToProto(o.Mode))
	p.SetAllowNoResourceGroupMatch(dcl.ValueOrEmptyBool(o.AllowNoResourceGroupMatch))
	sResourceGroups := make([]*osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroups, len(o.ResourceGroups))
	for i, r := range o.ResourceGroups {
		sResourceGroups[i] = OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsToProto(&r)
	}
	p.SetResourceGroups(sResourceGroups)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroups object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroups) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroups {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroups{}
	sInventoryFilters := make([]*osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters, len(o.InventoryFilters))
	for i, r := range o.InventoryFilters {
		sInventoryFilters[i] = OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto(&r)
	}
	p.SetInventoryFilters(sInventoryFilters)
	sResources := make([]*osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResources, len(o.Resources))
	for i, r := range o.Resources {
		sResources[i] = OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto(&r)
	}
	p.SetResources(sResources)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFiltersToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters{}
	p.SetOsShortName(dcl.ValueOrEmptyString(o.OSShortName))
	p.SetOsVersion(dcl.ValueOrEmptyString(o.OSVersion))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResources object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResources) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResources {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResources{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetPkg(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto(o.Pkg))
	p.SetRepository(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto(o.Repository))
	p.SetExec(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto(o.Exec))
	p.SetFile(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto(o.File))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{}
	p.SetDesiredState(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumToProto(o.DesiredState))
	p.SetApt(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto(o.Apt))
	p.SetDeb(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto(o.Deb))
	p.SetYum(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto(o.Yum))
	p.SetZypper(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto(o.Zypper))
	p.SetRpm(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto(o.Rpm))
	p.SetGooget(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto(o.Googet))
	p.SetMsi(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto(o.Msi))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgAptToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{}
	p.SetSource(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto(o.Source))
	p.SetPullDeps(dcl.ValueOrEmptyBool(o.PullDeps))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource{}
	p.SetRemote(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemoteToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcsToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYumToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypperToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{}
	p.SetSource(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto(o.Source))
	p.SetPullDeps(dcl.ValueOrEmptyBool(o.PullDeps))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource{}
	p.SetRemote(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemoteToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcsToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGoogetToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{}
	p.SetSource(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto(o.Source))
	sProperties := make([]string, len(o.Properties))
	for i, r := range o.Properties {
		sProperties[i] = r
	}
	p.SetProperties(sProperties)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource{}
	p.SetRemote(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto(o.Remote))
	p.SetGcs(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemoteToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcsToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{}
	p.SetApt(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto(o.Apt))
	p.SetYum(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto(o.Yum))
	p.SetZypper(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto(o.Zypper))
	p.SetGoo(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto(o.Goo))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{}
	p.SetArchiveType(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumToProto(o.ArchiveType))
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
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYumToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{}
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
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypperToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{}
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
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGooToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{}
	p.SetValidate(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto(o.Validate))
	p.SetEnforce(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto(o.Enforce))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate{}
	p.SetFile(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto(o.File))
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumToProto(o.Interpreter))
	p.SetOutputFilePath(dcl.ValueOrEmptyString(o.OutputFilePath))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile{}
	p.SetRemote(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemoteToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcsToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce{}
	p.SetFile(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto(o.File))
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumToProto(o.Interpreter))
	p.SetOutputFilePath(dcl.ValueOrEmptyString(o.OutputFilePath))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile{}
	p.SetRemote(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemoteToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcsToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{}
	p.SetFile(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto(o.File))
	p.SetContent(dcl.ValueOrEmptyString(o.Content))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetState(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumToProto(o.State))
	p.SetPermissions(dcl.ValueOrEmptyString(o.Permissions))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile{}
	p.SetRemote(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto(o.Remote))
	p.SetGcs(OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto(o.Gcs))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemoteToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetSha256Checksum(dcl.ValueOrEmptyString(o.Sha256Checksum))
	return p
}

// OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto converts a OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs object to its proto representation.
func OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcsToProto(o *osconfig.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs) *osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// OSPolicyAssignmentInstanceFilterToProto converts a OSPolicyAssignmentInstanceFilter object to its proto representation.
func OsconfigOSPolicyAssignmentInstanceFilterToProto(o *osconfig.OSPolicyAssignmentInstanceFilter) *osconfigpb.OsconfigOSPolicyAssignmentInstanceFilter {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentInstanceFilter{}
	p.SetAll(dcl.ValueOrEmptyBool(o.All))
	sInclusionLabels := make([]*osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterInclusionLabels, len(o.InclusionLabels))
	for i, r := range o.InclusionLabels {
		sInclusionLabels[i] = OsconfigOSPolicyAssignmentInstanceFilterInclusionLabelsToProto(&r)
	}
	p.SetInclusionLabels(sInclusionLabels)
	sExclusionLabels := make([]*osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterExclusionLabels, len(o.ExclusionLabels))
	for i, r := range o.ExclusionLabels {
		sExclusionLabels[i] = OsconfigOSPolicyAssignmentInstanceFilterExclusionLabelsToProto(&r)
	}
	p.SetExclusionLabels(sExclusionLabels)
	sInventories := make([]*osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterInventories, len(o.Inventories))
	for i, r := range o.Inventories {
		sInventories[i] = OsconfigOSPolicyAssignmentInstanceFilterInventoriesToProto(&r)
	}
	p.SetInventories(sInventories)
	return p
}

// OSPolicyAssignmentInstanceFilterInclusionLabelsToProto converts a OSPolicyAssignmentInstanceFilterInclusionLabels object to its proto representation.
func OsconfigOSPolicyAssignmentInstanceFilterInclusionLabelsToProto(o *osconfig.OSPolicyAssignmentInstanceFilterInclusionLabels) *osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterInclusionLabels {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterInclusionLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// OSPolicyAssignmentInstanceFilterExclusionLabelsToProto converts a OSPolicyAssignmentInstanceFilterExclusionLabels object to its proto representation.
func OsconfigOSPolicyAssignmentInstanceFilterExclusionLabelsToProto(o *osconfig.OSPolicyAssignmentInstanceFilterExclusionLabels) *osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterExclusionLabels {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterExclusionLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// OSPolicyAssignmentInstanceFilterInventoriesToProto converts a OSPolicyAssignmentInstanceFilterInventories object to its proto representation.
func OsconfigOSPolicyAssignmentInstanceFilterInventoriesToProto(o *osconfig.OSPolicyAssignmentInstanceFilterInventories) *osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterInventories {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentInstanceFilterInventories{}
	p.SetOsShortName(dcl.ValueOrEmptyString(o.OSShortName))
	p.SetOsVersion(dcl.ValueOrEmptyString(o.OSVersion))
	return p
}

// OSPolicyAssignmentRolloutToProto converts a OSPolicyAssignmentRollout object to its proto representation.
func OsconfigOSPolicyAssignmentRolloutToProto(o *osconfig.OSPolicyAssignmentRollout) *osconfigpb.OsconfigOSPolicyAssignmentRollout {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentRollout{}
	p.SetDisruptionBudget(OsconfigOSPolicyAssignmentRolloutDisruptionBudgetToProto(o.DisruptionBudget))
	p.SetMinWaitDuration(dcl.ValueOrEmptyString(o.MinWaitDuration))
	return p
}

// OSPolicyAssignmentRolloutDisruptionBudgetToProto converts a OSPolicyAssignmentRolloutDisruptionBudget object to its proto representation.
func OsconfigOSPolicyAssignmentRolloutDisruptionBudgetToProto(o *osconfig.OSPolicyAssignmentRolloutDisruptionBudget) *osconfigpb.OsconfigOSPolicyAssignmentRolloutDisruptionBudget {
	if o == nil {
		return nil
	}
	p := &osconfigpb.OsconfigOSPolicyAssignmentRolloutDisruptionBudget{}
	p.SetFixed(dcl.ValueOrEmptyInt64(o.Fixed))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	return p
}

// OSPolicyAssignmentToProto converts a OSPolicyAssignment resource to its proto representation.
func OSPolicyAssignmentToProto(resource *osconfig.OSPolicyAssignment) *osconfigpb.OsconfigOSPolicyAssignment {
	p := &osconfigpb.OsconfigOSPolicyAssignment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetInstanceFilter(OsconfigOSPolicyAssignmentInstanceFilterToProto(resource.InstanceFilter))
	p.SetRollout(OsconfigOSPolicyAssignmentRolloutToProto(resource.Rollout))
	p.SetRevisionId(dcl.ValueOrEmptyString(resource.RevisionId))
	p.SetRevisionCreateTime(dcl.ValueOrEmptyString(resource.RevisionCreateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetRolloutState(OsconfigOSPolicyAssignmentRolloutStateEnumToProto(resource.RolloutState))
	p.SetBaseline(dcl.ValueOrEmptyBool(resource.Baseline))
	p.SetDeleted(dcl.ValueOrEmptyBool(resource.Deleted))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSkipAwaitRollout(dcl.ValueOrEmptyBool(resource.SkipAwaitRollout))
	sOSPolicies := make([]*osconfigpb.OsconfigOSPolicyAssignmentOSPolicies, len(resource.OSPolicies))
	for i, r := range resource.OSPolicies {
		sOSPolicies[i] = OsconfigOSPolicyAssignmentOSPoliciesToProto(&r)
	}
	p.SetOsPolicies(sOSPolicies)

	return p
}

// applyOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Apply() method.
func (s *OSPolicyAssignmentServer) applyOSPolicyAssignment(ctx context.Context, c *osconfig.Client, request *osconfigpb.ApplyOsconfigOSPolicyAssignmentRequest) (*osconfigpb.OsconfigOSPolicyAssignment, error) {
	p := ProtoToOSPolicyAssignment(request.GetResource())
	res, err := c.ApplyOSPolicyAssignment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := OSPolicyAssignmentToProto(res)
	return r, nil
}

// applyOsconfigOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Apply() method.
func (s *OSPolicyAssignmentServer) ApplyOsconfigOSPolicyAssignment(ctx context.Context, request *osconfigpb.ApplyOsconfigOSPolicyAssignmentRequest) (*osconfigpb.OsconfigOSPolicyAssignment, error) {
	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyOSPolicyAssignment(ctx, cl, request)
}

// DeleteOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignment Delete() method.
func (s *OSPolicyAssignmentServer) DeleteOsconfigOSPolicyAssignment(ctx context.Context, request *osconfigpb.DeleteOsconfigOSPolicyAssignmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteOSPolicyAssignment(ctx, ProtoToOSPolicyAssignment(request.GetResource()))

}

// ListOsconfigOSPolicyAssignment handles the gRPC request by passing it to the underlying OSPolicyAssignmentList() method.
func (s *OSPolicyAssignmentServer) ListOsconfigOSPolicyAssignment(ctx context.Context, request *osconfigpb.ListOsconfigOSPolicyAssignmentRequest) (*osconfigpb.ListOsconfigOSPolicyAssignmentResponse, error) {
	cl, err := createConfigOSPolicyAssignment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListOSPolicyAssignment(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*osconfigpb.OsconfigOSPolicyAssignment
	for _, r := range resources.Items {
		rp := OSPolicyAssignmentToProto(r)
		protos = append(protos, rp)
	}
	p := &osconfigpb.ListOsconfigOSPolicyAssignmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigOSPolicyAssignment(ctx context.Context, service_account_file string) (*osconfig.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return osconfig.NewClient(conf), nil
}
