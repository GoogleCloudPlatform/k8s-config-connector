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

// GuestPolicyServer implements the gRPC interface for GuestPolicy.
type GuestPolicyServer struct{}

// ProtoToGuestPolicyPackagesDesiredStateEnum converts a GuestPolicyPackagesDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackagesDesiredStateEnum(e alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum) *alpha.GuestPolicyPackagesDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyPackagesDesiredStateEnum(n[len("OsconfigAlphaGuestPolicyPackagesDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyPackagesManagerEnum converts a GuestPolicyPackagesManagerEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackagesManagerEnum(e alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum) *alpha.GuestPolicyPackagesManagerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyPackagesManagerEnum(n[len("OsconfigAlphaGuestPolicyPackagesManagerEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyPackageRepositoriesAptArchiveTypeEnum converts a GuestPolicyPackageRepositoriesAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(e alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum) *alpha.GuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyPackageRepositoriesAptArchiveTypeEnum(n[len("OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum converts a GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) *alpha.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(n[len("OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum converts a GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) *alpha.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(n[len("OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum converts a GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) *alpha.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(n[len("OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum converts a GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) *alpha.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(n[len("OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesDesiredStateEnum converts a GuestPolicyRecipesDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesDesiredStateEnum(e alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum) *alpha.GuestPolicyRecipesDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum_name[int32(e)]; ok {
		e := alpha.GuestPolicyRecipesDesiredStateEnum(n[len("OsconfigAlphaGuestPolicyRecipesDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyAssignment converts a GuestPolicyAssignment object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyAssignment(p *alphapb.OsconfigAlphaGuestPolicyAssignment) *alpha.GuestPolicyAssignment {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyAssignment{}
	for _, r := range p.GetGroupLabels() {
		obj.GroupLabels = append(obj.GroupLabels, *ProtoToOsconfigAlphaGuestPolicyAssignmentGroupLabels(r))
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, r)
	}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, r)
	}
	for _, r := range p.GetInstanceNamePrefixes() {
		obj.InstanceNamePrefixes = append(obj.InstanceNamePrefixes, r)
	}
	for _, r := range p.GetOsTypes() {
		obj.OSTypes = append(obj.OSTypes, *ProtoToOsconfigAlphaGuestPolicyAssignmentOSTypes(r))
	}
	return obj
}

// ProtoToGuestPolicyAssignmentGroupLabels converts a GuestPolicyAssignmentGroupLabels object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyAssignmentGroupLabels(p *alphapb.OsconfigAlphaGuestPolicyAssignmentGroupLabels) *alpha.GuestPolicyAssignmentGroupLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyAssignmentGroupLabels{}
	return obj
}

// ProtoToGuestPolicyAssignmentOSTypes converts a GuestPolicyAssignmentOSTypes object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyAssignmentOSTypes(p *alphapb.OsconfigAlphaGuestPolicyAssignmentOSTypes) *alpha.GuestPolicyAssignmentOSTypes {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyAssignmentOSTypes{
		OSShortName:    dcl.StringOrNil(p.GetOsShortName()),
		OSVersion:      dcl.StringOrNil(p.GetOsVersion()),
		OSArchitecture: dcl.StringOrNil(p.GetOsArchitecture()),
	}
	return obj
}

// ProtoToGuestPolicyPackages converts a GuestPolicyPackages object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackages(p *alphapb.OsconfigAlphaGuestPolicyPackages) *alpha.GuestPolicyPackages {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackages{
		Name:         dcl.StringOrNil(p.GetName()),
		DesiredState: ProtoToOsconfigAlphaGuestPolicyPackagesDesiredStateEnum(p.GetDesiredState()),
		Manager:      ProtoToOsconfigAlphaGuestPolicyPackagesManagerEnum(p.GetManager()),
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositories converts a GuestPolicyPackageRepositories object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositories(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositories) *alpha.GuestPolicyPackageRepositories {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositories{
		Apt:    ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesApt(p.GetApt()),
		Yum:    ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesYum(p.GetYum()),
		Zypper: ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesApt converts a GuestPolicyPackageRepositoriesApt object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesApt(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesApt) *alpha.GuestPolicyPackageRepositoriesApt {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositoriesApt{
		ArchiveType:  ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(p.GetArchiveType()),
		Uri:          dcl.StringOrNil(p.GetUri()),
		Distribution: dcl.StringOrNil(p.GetDistribution()),
		GpgKey:       dcl.StringOrNil(p.GetGpgKey()),
	}
	for _, r := range p.GetComponents() {
		obj.Components = append(obj.Components, r)
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesYum converts a GuestPolicyPackageRepositoriesYum object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesYum(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesYum) *alpha.GuestPolicyPackageRepositoriesYum {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositoriesYum{
		Id:          dcl.StringOrNil(p.GetId()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		BaseUrl:     dcl.StringOrNil(p.GetBaseUrl()),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesZypper converts a GuestPolicyPackageRepositoriesZypper object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesZypper(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesZypper) *alpha.GuestPolicyPackageRepositoriesZypper {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositoriesZypper{
		Id:          dcl.StringOrNil(p.GetId()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		BaseUrl:     dcl.StringOrNil(p.GetBaseUrl()),
	}
	for _, r := range p.GetGpgKeys() {
		obj.GpgKeys = append(obj.GpgKeys, r)
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesGoo converts a GuestPolicyPackageRepositoriesGoo object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyPackageRepositoriesGoo(p *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesGoo) *alpha.GuestPolicyPackageRepositoriesGoo {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyPackageRepositoriesGoo{
		Name: dcl.StringOrNil(p.GetName()),
		Url:  dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToGuestPolicyRecipes converts a GuestPolicyRecipes object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipes(p *alphapb.OsconfigAlphaGuestPolicyRecipes) *alpha.GuestPolicyRecipes {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipes{
		Name:         dcl.StringOrNil(p.GetName()),
		Version:      dcl.StringOrNil(p.GetVersion()),
		DesiredState: ProtoToOsconfigAlphaGuestPolicyRecipesDesiredStateEnum(p.GetDesiredState()),
	}
	for _, r := range p.GetArtifacts() {
		obj.Artifacts = append(obj.Artifacts, *ProtoToOsconfigAlphaGuestPolicyRecipesArtifacts(r))
	}
	for _, r := range p.GetInstallSteps() {
		obj.InstallSteps = append(obj.InstallSteps, *ProtoToOsconfigAlphaGuestPolicyRecipesInstallSteps(r))
	}
	for _, r := range p.GetUpdateSteps() {
		obj.UpdateSteps = append(obj.UpdateSteps, *ProtoToOsconfigAlphaGuestPolicyRecipesUpdateSteps(r))
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifacts converts a GuestPolicyRecipesArtifacts object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesArtifacts(p *alphapb.OsconfigAlphaGuestPolicyRecipesArtifacts) *alpha.GuestPolicyRecipesArtifacts {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesArtifacts{
		Id:            dcl.StringOrNil(p.GetId()),
		Remote:        ProtoToOsconfigAlphaGuestPolicyRecipesArtifactsRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigAlphaGuestPolicyRecipesArtifactsGcs(p.GetGcs()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifactsRemote converts a GuestPolicyRecipesArtifactsRemote object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesArtifactsRemote(p *alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsRemote) *alpha.GuestPolicyRecipesArtifactsRemote {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesArtifactsRemote{
		Uri:      dcl.StringOrNil(p.GetUri()),
		Checksum: dcl.StringOrNil(p.GetChecksum()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifactsGcs converts a GuestPolicyRecipesArtifactsGcs object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesArtifactsGcs(p *alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsGcs) *alpha.GuestPolicyRecipesArtifactsGcs {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesArtifactsGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallSteps converts a GuestPolicyRecipesInstallSteps object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallSteps(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallSteps) *alpha.GuestPolicyRecipesInstallSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallSteps{
		FileCopy:          ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy(p.GetFileCopy()),
		ArchiveExtraction: ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction(p.GetArchiveExtraction()),
		MsiInstallation:   ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation(p.GetMsiInstallation()),
		DpkgInstallation:  ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation(p.GetDpkgInstallation()),
		RpmInstallation:   ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation(p.GetRpmInstallation()),
		FileExec:          ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsFileExec(p.GetFileExec()),
		ScriptRun:         ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun(p.GetScriptRun()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsFileCopy converts a GuestPolicyRecipesInstallStepsFileCopy object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy) *alpha.GuestPolicyRecipesInstallStepsFileCopy {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsFileCopy{
		ArtifactId:  dcl.StringOrNil(p.GetArtifactId()),
		Destination: dcl.StringOrNil(p.GetDestination()),
		Overwrite:   dcl.Bool(p.GetOverwrite()),
		Permissions: dcl.StringOrNil(p.GetPermissions()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsArchiveExtraction converts a GuestPolicyRecipesInstallStepsArchiveExtraction object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction) *alpha.GuestPolicyRecipesInstallStepsArchiveExtraction {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsArchiveExtraction{
		ArtifactId:  dcl.StringOrNil(p.GetArtifactId()),
		Destination: dcl.StringOrNil(p.GetDestination()),
		Type:        ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsMsiInstallation converts a GuestPolicyRecipesInstallStepsMsiInstallation object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation) *alpha.GuestPolicyRecipesInstallStepsMsiInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsMsiInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	for _, r := range p.GetFlags() {
		obj.Flags = append(obj.Flags, r)
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsDpkgInstallation converts a GuestPolicyRecipesInstallStepsDpkgInstallation object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation) *alpha.GuestPolicyRecipesInstallStepsDpkgInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsDpkgInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsRpmInstallation converts a GuestPolicyRecipesInstallStepsRpmInstallation object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation) *alpha.GuestPolicyRecipesInstallStepsRpmInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsRpmInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsFileExec converts a GuestPolicyRecipesInstallStepsFileExec object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsFileExec(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileExec) *alpha.GuestPolicyRecipesInstallStepsFileExec {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsFileExec{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
		LocalPath:  dcl.StringOrNil(p.GetLocalPath()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsScriptRun converts a GuestPolicyRecipesInstallStepsScriptRun object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun(p *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun) *alpha.GuestPolicyRecipesInstallStepsScriptRun {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesInstallStepsScriptRun{
		Script:      dcl.StringOrNil(p.GetScript()),
		Interpreter: ProtoToOsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateSteps converts a GuestPolicyRecipesUpdateSteps object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateSteps(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateSteps) *alpha.GuestPolicyRecipesUpdateSteps {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateSteps{
		FileCopy:          ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy(p.GetFileCopy()),
		ArchiveExtraction: ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction(p.GetArchiveExtraction()),
		MsiInstallation:   ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation(p.GetMsiInstallation()),
		DpkgInstallation:  ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation(p.GetDpkgInstallation()),
		RpmInstallation:   ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation(p.GetRpmInstallation()),
		FileExec:          ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec(p.GetFileExec()),
		ScriptRun:         ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun(p.GetScriptRun()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsFileCopy converts a GuestPolicyRecipesUpdateStepsFileCopy object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy) *alpha.GuestPolicyRecipesUpdateStepsFileCopy {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsFileCopy{
		ArtifactId:  dcl.StringOrNil(p.GetArtifactId()),
		Destination: dcl.StringOrNil(p.GetDestination()),
		Overwrite:   dcl.Bool(p.GetOverwrite()),
		Permissions: dcl.StringOrNil(p.GetPermissions()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsArchiveExtraction converts a GuestPolicyRecipesUpdateStepsArchiveExtraction object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction) *alpha.GuestPolicyRecipesUpdateStepsArchiveExtraction {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsArchiveExtraction{
		ArtifactId:  dcl.StringOrNil(p.GetArtifactId()),
		Destination: dcl.StringOrNil(p.GetDestination()),
		Type:        ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsMsiInstallation converts a GuestPolicyRecipesUpdateStepsMsiInstallation object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation) *alpha.GuestPolicyRecipesUpdateStepsMsiInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsMsiInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	for _, r := range p.GetFlags() {
		obj.Flags = append(obj.Flags, r)
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsDpkgInstallation converts a GuestPolicyRecipesUpdateStepsDpkgInstallation object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation) *alpha.GuestPolicyRecipesUpdateStepsDpkgInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsDpkgInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsRpmInstallation converts a GuestPolicyRecipesUpdateStepsRpmInstallation object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation) *alpha.GuestPolicyRecipesUpdateStepsRpmInstallation {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsRpmInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsFileExec converts a GuestPolicyRecipesUpdateStepsFileExec object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec) *alpha.GuestPolicyRecipesUpdateStepsFileExec {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsFileExec{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
		LocalPath:  dcl.StringOrNil(p.GetLocalPath()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsScriptRun converts a GuestPolicyRecipesUpdateStepsScriptRun object from its proto representation.
func ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun(p *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun) *alpha.GuestPolicyRecipesUpdateStepsScriptRun {
	if p == nil {
		return nil
	}
	obj := &alpha.GuestPolicyRecipesUpdateStepsScriptRun{
		Script:      dcl.StringOrNil(p.GetScript()),
		Interpreter: ProtoToOsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicy converts a GuestPolicy resource from its proto representation.
func ProtoToGuestPolicy(p *alphapb.OsconfigAlphaGuestPolicy) *alpha.GuestPolicy {
	obj := &alpha.GuestPolicy{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Assignment:  ProtoToOsconfigAlphaGuestPolicyAssignment(p.GetAssignment()),
		Etag:        dcl.StringOrNil(p.GetEtag()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetPackages() {
		obj.Packages = append(obj.Packages, *ProtoToOsconfigAlphaGuestPolicyPackages(r))
	}
	for _, r := range p.GetPackageRepositories() {
		obj.PackageRepositories = append(obj.PackageRepositories, *ProtoToOsconfigAlphaGuestPolicyPackageRepositories(r))
	}
	for _, r := range p.GetRecipes() {
		obj.Recipes = append(obj.Recipes, *ProtoToOsconfigAlphaGuestPolicyRecipes(r))
	}
	return obj
}

// GuestPolicyPackagesDesiredStateEnumToProto converts a GuestPolicyPackagesDesiredStateEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyPackagesDesiredStateEnumToProto(e *alpha.GuestPolicyPackagesDesiredStateEnum) alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum_value["GuestPolicyPackagesDesiredStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyPackagesDesiredStateEnum(0)
}

// GuestPolicyPackagesManagerEnumToProto converts a GuestPolicyPackagesManagerEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyPackagesManagerEnumToProto(e *alpha.GuestPolicyPackagesManagerEnum) alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum_value["GuestPolicyPackagesManagerEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyPackagesManagerEnum(0)
}

// GuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto converts a GuestPolicyPackageRepositoriesAptArchiveTypeEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto(e *alpha.GuestPolicyPackageRepositoriesAptArchiveTypeEnum) alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum_value["GuestPolicyPackageRepositoriesAptArchiveTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(0)
}

// GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto converts a GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto(e *alpha.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum_value["GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(0)
}

// GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto converts a GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto(e *alpha.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum_value["GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(0)
}

// GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto converts a GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto(e *alpha.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum_value["GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(0)
}

// GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto converts a GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto(e *alpha.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum_value["GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(0)
}

// GuestPolicyRecipesDesiredStateEnumToProto converts a GuestPolicyRecipesDesiredStateEnum enum to its proto representation.
func OsconfigAlphaGuestPolicyRecipesDesiredStateEnumToProto(e *alpha.GuestPolicyRecipesDesiredStateEnum) alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum {
	if e == nil {
		return alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum(0)
	}
	if v, ok := alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum_value["GuestPolicyRecipesDesiredStateEnum"+string(*e)]; ok {
		return alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum(v)
	}
	return alphapb.OsconfigAlphaGuestPolicyRecipesDesiredStateEnum(0)
}

// GuestPolicyAssignmentToProto converts a GuestPolicyAssignment object to its proto representation.
func OsconfigAlphaGuestPolicyAssignmentToProto(o *alpha.GuestPolicyAssignment) *alphapb.OsconfigAlphaGuestPolicyAssignment {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyAssignment{}
	sGroupLabels := make([]*alphapb.OsconfigAlphaGuestPolicyAssignmentGroupLabels, len(o.GroupLabels))
	for i, r := range o.GroupLabels {
		sGroupLabels[i] = OsconfigAlphaGuestPolicyAssignmentGroupLabelsToProto(&r)
	}
	p.SetGroupLabels(sGroupLabels)
	sZones := make([]string, len(o.Zones))
	for i, r := range o.Zones {
		sZones[i] = r
	}
	p.SetZones(sZones)
	sInstances := make([]string, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = r
	}
	p.SetInstances(sInstances)
	sInstanceNamePrefixes := make([]string, len(o.InstanceNamePrefixes))
	for i, r := range o.InstanceNamePrefixes {
		sInstanceNamePrefixes[i] = r
	}
	p.SetInstanceNamePrefixes(sInstanceNamePrefixes)
	sOSTypes := make([]*alphapb.OsconfigAlphaGuestPolicyAssignmentOSTypes, len(o.OSTypes))
	for i, r := range o.OSTypes {
		sOSTypes[i] = OsconfigAlphaGuestPolicyAssignmentOSTypesToProto(&r)
	}
	p.SetOsTypes(sOSTypes)
	return p
}

// GuestPolicyAssignmentGroupLabelsToProto converts a GuestPolicyAssignmentGroupLabels object to its proto representation.
func OsconfigAlphaGuestPolicyAssignmentGroupLabelsToProto(o *alpha.GuestPolicyAssignmentGroupLabels) *alphapb.OsconfigAlphaGuestPolicyAssignmentGroupLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyAssignmentGroupLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// GuestPolicyAssignmentOSTypesToProto converts a GuestPolicyAssignmentOSTypes object to its proto representation.
func OsconfigAlphaGuestPolicyAssignmentOSTypesToProto(o *alpha.GuestPolicyAssignmentOSTypes) *alphapb.OsconfigAlphaGuestPolicyAssignmentOSTypes {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyAssignmentOSTypes{}
	p.SetOsShortName(dcl.ValueOrEmptyString(o.OSShortName))
	p.SetOsVersion(dcl.ValueOrEmptyString(o.OSVersion))
	p.SetOsArchitecture(dcl.ValueOrEmptyString(o.OSArchitecture))
	return p
}

// GuestPolicyPackagesToProto converts a GuestPolicyPackages object to its proto representation.
func OsconfigAlphaGuestPolicyPackagesToProto(o *alpha.GuestPolicyPackages) *alphapb.OsconfigAlphaGuestPolicyPackages {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackages{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDesiredState(OsconfigAlphaGuestPolicyPackagesDesiredStateEnumToProto(o.DesiredState))
	p.SetManager(OsconfigAlphaGuestPolicyPackagesManagerEnumToProto(o.Manager))
	return p
}

// GuestPolicyPackageRepositoriesToProto converts a GuestPolicyPackageRepositories object to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesToProto(o *alpha.GuestPolicyPackageRepositories) *alphapb.OsconfigAlphaGuestPolicyPackageRepositories {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositories{}
	p.SetApt(OsconfigAlphaGuestPolicyPackageRepositoriesAptToProto(o.Apt))
	p.SetYum(OsconfigAlphaGuestPolicyPackageRepositoriesYumToProto(o.Yum))
	p.SetZypper(OsconfigAlphaGuestPolicyPackageRepositoriesZypperToProto(o.Zypper))
	p.SetGoo(OsconfigAlphaGuestPolicyPackageRepositoriesGooToProto(o.Goo))
	return p
}

// GuestPolicyPackageRepositoriesAptToProto converts a GuestPolicyPackageRepositoriesApt object to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesAptToProto(o *alpha.GuestPolicyPackageRepositoriesApt) *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesApt {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesApt{}
	p.SetArchiveType(OsconfigAlphaGuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto(o.ArchiveType))
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

// GuestPolicyPackageRepositoriesYumToProto converts a GuestPolicyPackageRepositoriesYum object to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesYumToProto(o *alpha.GuestPolicyPackageRepositoriesYum) *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesYum {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesYum{}
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

// GuestPolicyPackageRepositoriesZypperToProto converts a GuestPolicyPackageRepositoriesZypper object to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesZypperToProto(o *alpha.GuestPolicyPackageRepositoriesZypper) *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesZypper {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesZypper{}
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

// GuestPolicyPackageRepositoriesGooToProto converts a GuestPolicyPackageRepositoriesGoo object to its proto representation.
func OsconfigAlphaGuestPolicyPackageRepositoriesGooToProto(o *alpha.GuestPolicyPackageRepositoriesGoo) *alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesGoo {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyPackageRepositoriesGoo{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// GuestPolicyRecipesToProto converts a GuestPolicyRecipes object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesToProto(o *alpha.GuestPolicyRecipes) *alphapb.OsconfigAlphaGuestPolicyRecipes {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetDesiredState(OsconfigAlphaGuestPolicyRecipesDesiredStateEnumToProto(o.DesiredState))
	sArtifacts := make([]*alphapb.OsconfigAlphaGuestPolicyRecipesArtifacts, len(o.Artifacts))
	for i, r := range o.Artifacts {
		sArtifacts[i] = OsconfigAlphaGuestPolicyRecipesArtifactsToProto(&r)
	}
	p.SetArtifacts(sArtifacts)
	sInstallSteps := make([]*alphapb.OsconfigAlphaGuestPolicyRecipesInstallSteps, len(o.InstallSteps))
	for i, r := range o.InstallSteps {
		sInstallSteps[i] = OsconfigAlphaGuestPolicyRecipesInstallStepsToProto(&r)
	}
	p.SetInstallSteps(sInstallSteps)
	sUpdateSteps := make([]*alphapb.OsconfigAlphaGuestPolicyRecipesUpdateSteps, len(o.UpdateSteps))
	for i, r := range o.UpdateSteps {
		sUpdateSteps[i] = OsconfigAlphaGuestPolicyRecipesUpdateStepsToProto(&r)
	}
	p.SetUpdateSteps(sUpdateSteps)
	return p
}

// GuestPolicyRecipesArtifactsToProto converts a GuestPolicyRecipesArtifacts object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesArtifactsToProto(o *alpha.GuestPolicyRecipesArtifacts) *alphapb.OsconfigAlphaGuestPolicyRecipesArtifacts {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesArtifacts{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetRemote(OsconfigAlphaGuestPolicyRecipesArtifactsRemoteToProto(o.Remote))
	p.SetGcs(OsconfigAlphaGuestPolicyRecipesArtifactsGcsToProto(o.Gcs))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// GuestPolicyRecipesArtifactsRemoteToProto converts a GuestPolicyRecipesArtifactsRemote object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesArtifactsRemoteToProto(o *alpha.GuestPolicyRecipesArtifactsRemote) *alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsRemote {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetChecksum(dcl.ValueOrEmptyString(o.Checksum))
	return p
}

// GuestPolicyRecipesArtifactsGcsToProto converts a GuestPolicyRecipesArtifactsGcs object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesArtifactsGcsToProto(o *alpha.GuestPolicyRecipesArtifactsGcs) *alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsGcs {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesArtifactsGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// GuestPolicyRecipesInstallStepsToProto converts a GuestPolicyRecipesInstallSteps object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsToProto(o *alpha.GuestPolicyRecipesInstallSteps) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallSteps{}
	p.SetFileCopy(OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopyToProto(o.FileCopy))
	p.SetArchiveExtraction(OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionToProto(o.ArchiveExtraction))
	p.SetMsiInstallation(OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallationToProto(o.MsiInstallation))
	p.SetDpkgInstallation(OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallationToProto(o.DpkgInstallation))
	p.SetRpmInstallation(OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallationToProto(o.RpmInstallation))
	p.SetFileExec(OsconfigAlphaGuestPolicyRecipesInstallStepsFileExecToProto(o.FileExec))
	p.SetScriptRun(OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunToProto(o.ScriptRun))
	return p
}

// GuestPolicyRecipesInstallStepsFileCopyToProto converts a GuestPolicyRecipesInstallStepsFileCopy object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopyToProto(o *alpha.GuestPolicyRecipesInstallStepsFileCopy) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileCopy{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetDestination(dcl.ValueOrEmptyString(o.Destination))
	p.SetOverwrite(dcl.ValueOrEmptyBool(o.Overwrite))
	p.SetPermissions(dcl.ValueOrEmptyString(o.Permissions))
	return p
}

// GuestPolicyRecipesInstallStepsArchiveExtractionToProto converts a GuestPolicyRecipesInstallStepsArchiveExtraction object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionToProto(o *alpha.GuestPolicyRecipesInstallStepsArchiveExtraction) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtraction{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetDestination(dcl.ValueOrEmptyString(o.Destination))
	p.SetType(OsconfigAlphaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto(o.Type))
	return p
}

// GuestPolicyRecipesInstallStepsMsiInstallationToProto converts a GuestPolicyRecipesInstallStepsMsiInstallation object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallationToProto(o *alpha.GuestPolicyRecipesInstallStepsMsiInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsMsiInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	sFlags := make([]string, len(o.Flags))
	for i, r := range o.Flags {
		sFlags[i] = r
	}
	p.SetFlags(sFlags)
	sAllowedExitCodes := make([]int64, len(o.AllowedExitCodes))
	for i, r := range o.AllowedExitCodes {
		sAllowedExitCodes[i] = r
	}
	p.SetAllowedExitCodes(sAllowedExitCodes)
	return p
}

// GuestPolicyRecipesInstallStepsDpkgInstallationToProto converts a GuestPolicyRecipesInstallStepsDpkgInstallation object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallationToProto(o *alpha.GuestPolicyRecipesInstallStepsDpkgInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsDpkgInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	return p
}

// GuestPolicyRecipesInstallStepsRpmInstallationToProto converts a GuestPolicyRecipesInstallStepsRpmInstallation object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallationToProto(o *alpha.GuestPolicyRecipesInstallStepsRpmInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsRpmInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	return p
}

// GuestPolicyRecipesInstallStepsFileExecToProto converts a GuestPolicyRecipesInstallStepsFileExec object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsFileExecToProto(o *alpha.GuestPolicyRecipesInstallStepsFileExec) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsFileExec{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sAllowedExitCodes := make([]int64, len(o.AllowedExitCodes))
	for i, r := range o.AllowedExitCodes {
		sAllowedExitCodes[i] = r
	}
	p.SetAllowedExitCodes(sAllowedExitCodes)
	return p
}

// GuestPolicyRecipesInstallStepsScriptRunToProto converts a GuestPolicyRecipesInstallStepsScriptRun object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunToProto(o *alpha.GuestPolicyRecipesInstallStepsScriptRun) *alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRun{}
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigAlphaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto(o.Interpreter))
	sAllowedExitCodes := make([]int64, len(o.AllowedExitCodes))
	for i, r := range o.AllowedExitCodes {
		sAllowedExitCodes[i] = r
	}
	p.SetAllowedExitCodes(sAllowedExitCodes)
	return p
}

// GuestPolicyRecipesUpdateStepsToProto converts a GuestPolicyRecipesUpdateSteps object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsToProto(o *alpha.GuestPolicyRecipesUpdateSteps) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateSteps {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateSteps{}
	p.SetFileCopy(OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopyToProto(o.FileCopy))
	p.SetArchiveExtraction(OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionToProto(o.ArchiveExtraction))
	p.SetMsiInstallation(OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallationToProto(o.MsiInstallation))
	p.SetDpkgInstallation(OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallationToProto(o.DpkgInstallation))
	p.SetRpmInstallation(OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallationToProto(o.RpmInstallation))
	p.SetFileExec(OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExecToProto(o.FileExec))
	p.SetScriptRun(OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunToProto(o.ScriptRun))
	return p
}

// GuestPolicyRecipesUpdateStepsFileCopyToProto converts a GuestPolicyRecipesUpdateStepsFileCopy object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopyToProto(o *alpha.GuestPolicyRecipesUpdateStepsFileCopy) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileCopy{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetDestination(dcl.ValueOrEmptyString(o.Destination))
	p.SetOverwrite(dcl.ValueOrEmptyBool(o.Overwrite))
	p.SetPermissions(dcl.ValueOrEmptyString(o.Permissions))
	return p
}

// GuestPolicyRecipesUpdateStepsArchiveExtractionToProto converts a GuestPolicyRecipesUpdateStepsArchiveExtraction object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionToProto(o *alpha.GuestPolicyRecipesUpdateStepsArchiveExtraction) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtraction{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetDestination(dcl.ValueOrEmptyString(o.Destination))
	p.SetType(OsconfigAlphaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto(o.Type))
	return p
}

// GuestPolicyRecipesUpdateStepsMsiInstallationToProto converts a GuestPolicyRecipesUpdateStepsMsiInstallation object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallationToProto(o *alpha.GuestPolicyRecipesUpdateStepsMsiInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsMsiInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	sFlags := make([]string, len(o.Flags))
	for i, r := range o.Flags {
		sFlags[i] = r
	}
	p.SetFlags(sFlags)
	sAllowedExitCodes := make([]int64, len(o.AllowedExitCodes))
	for i, r := range o.AllowedExitCodes {
		sAllowedExitCodes[i] = r
	}
	p.SetAllowedExitCodes(sAllowedExitCodes)
	return p
}

// GuestPolicyRecipesUpdateStepsDpkgInstallationToProto converts a GuestPolicyRecipesUpdateStepsDpkgInstallation object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallationToProto(o *alpha.GuestPolicyRecipesUpdateStepsDpkgInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsDpkgInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	return p
}

// GuestPolicyRecipesUpdateStepsRpmInstallationToProto converts a GuestPolicyRecipesUpdateStepsRpmInstallation object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallationToProto(o *alpha.GuestPolicyRecipesUpdateStepsRpmInstallation) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsRpmInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	return p
}

// GuestPolicyRecipesUpdateStepsFileExecToProto converts a GuestPolicyRecipesUpdateStepsFileExec object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExecToProto(o *alpha.GuestPolicyRecipesUpdateStepsFileExec) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsFileExec{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetLocalPath(dcl.ValueOrEmptyString(o.LocalPath))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sAllowedExitCodes := make([]int64, len(o.AllowedExitCodes))
	for i, r := range o.AllowedExitCodes {
		sAllowedExitCodes[i] = r
	}
	p.SetAllowedExitCodes(sAllowedExitCodes)
	return p
}

// GuestPolicyRecipesUpdateStepsScriptRunToProto converts a GuestPolicyRecipesUpdateStepsScriptRun object to its proto representation.
func OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunToProto(o *alpha.GuestPolicyRecipesUpdateStepsScriptRun) *alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun {
	if o == nil {
		return nil
	}
	p := &alphapb.OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRun{}
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigAlphaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto(o.Interpreter))
	sAllowedExitCodes := make([]int64, len(o.AllowedExitCodes))
	for i, r := range o.AllowedExitCodes {
		sAllowedExitCodes[i] = r
	}
	p.SetAllowedExitCodes(sAllowedExitCodes)
	return p
}

// GuestPolicyToProto converts a GuestPolicy resource to its proto representation.
func GuestPolicyToProto(resource *alpha.GuestPolicy) *alphapb.OsconfigAlphaGuestPolicy {
	p := &alphapb.OsconfigAlphaGuestPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetAssignment(OsconfigAlphaGuestPolicyAssignmentToProto(resource.Assignment))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sPackages := make([]*alphapb.OsconfigAlphaGuestPolicyPackages, len(resource.Packages))
	for i, r := range resource.Packages {
		sPackages[i] = OsconfigAlphaGuestPolicyPackagesToProto(&r)
	}
	p.SetPackages(sPackages)
	sPackageRepositories := make([]*alphapb.OsconfigAlphaGuestPolicyPackageRepositories, len(resource.PackageRepositories))
	for i, r := range resource.PackageRepositories {
		sPackageRepositories[i] = OsconfigAlphaGuestPolicyPackageRepositoriesToProto(&r)
	}
	p.SetPackageRepositories(sPackageRepositories)
	sRecipes := make([]*alphapb.OsconfigAlphaGuestPolicyRecipes, len(resource.Recipes))
	for i, r := range resource.Recipes {
		sRecipes[i] = OsconfigAlphaGuestPolicyRecipesToProto(&r)
	}
	p.SetRecipes(sRecipes)

	return p
}

// applyGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Apply() method.
func (s *GuestPolicyServer) applyGuestPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyOsconfigAlphaGuestPolicyRequest) (*alphapb.OsconfigAlphaGuestPolicy, error) {
	p := ProtoToGuestPolicy(request.GetResource())
	res, err := c.ApplyGuestPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GuestPolicyToProto(res)
	return r, nil
}

// applyOsconfigAlphaGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Apply() method.
func (s *GuestPolicyServer) ApplyOsconfigAlphaGuestPolicy(ctx context.Context, request *alphapb.ApplyOsconfigAlphaGuestPolicyRequest) (*alphapb.OsconfigAlphaGuestPolicy, error) {
	cl, err := createConfigGuestPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGuestPolicy(ctx, cl, request)
}

// DeleteGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Delete() method.
func (s *GuestPolicyServer) DeleteOsconfigAlphaGuestPolicy(ctx context.Context, request *alphapb.DeleteOsconfigAlphaGuestPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGuestPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGuestPolicy(ctx, ProtoToGuestPolicy(request.GetResource()))

}

// ListOsconfigAlphaGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicyList() method.
func (s *GuestPolicyServer) ListOsconfigAlphaGuestPolicy(ctx context.Context, request *alphapb.ListOsconfigAlphaGuestPolicyRequest) (*alphapb.ListOsconfigAlphaGuestPolicyResponse, error) {
	cl, err := createConfigGuestPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGuestPolicy(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.OsconfigAlphaGuestPolicy
	for _, r := range resources.Items {
		rp := GuestPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListOsconfigAlphaGuestPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGuestPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
