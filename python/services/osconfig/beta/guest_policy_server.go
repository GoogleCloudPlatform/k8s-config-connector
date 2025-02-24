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

// GuestPolicyServer implements the gRPC interface for GuestPolicy.
type GuestPolicyServer struct{}

// ProtoToGuestPolicyPackagesDesiredStateEnum converts a GuestPolicyPackagesDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackagesDesiredStateEnum(e betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum) *beta.GuestPolicyPackagesDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyPackagesDesiredStateEnum(n[len("OsconfigBetaGuestPolicyPackagesDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyPackagesManagerEnum converts a GuestPolicyPackagesManagerEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackagesManagerEnum(e betapb.OsconfigBetaGuestPolicyPackagesManagerEnum) *beta.GuestPolicyPackagesManagerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyPackagesManagerEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyPackagesManagerEnum(n[len("OsconfigBetaGuestPolicyPackagesManagerEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyPackageRepositoriesAptArchiveTypeEnum converts a GuestPolicyPackageRepositoriesAptArchiveTypeEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(e betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum) *beta.GuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyPackageRepositoriesAptArchiveTypeEnum(n[len("OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum converts a GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(e betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) *beta.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(n[len("OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum converts a GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(e betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) *beta.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(n[len("OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum converts a GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(e betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) *beta.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(n[len("OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum converts a GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(e betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) *beta.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(n[len("OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyRecipesDesiredStateEnum converts a GuestPolicyRecipesDesiredStateEnum enum from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesDesiredStateEnum(e betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum) *beta.GuestPolicyRecipesDesiredStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum_name[int32(e)]; ok {
		e := beta.GuestPolicyRecipesDesiredStateEnum(n[len("OsconfigBetaGuestPolicyRecipesDesiredStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToGuestPolicyAssignment converts a GuestPolicyAssignment object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyAssignment(p *betapb.OsconfigBetaGuestPolicyAssignment) *beta.GuestPolicyAssignment {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyAssignment{}
	for _, r := range p.GetGroupLabels() {
		obj.GroupLabels = append(obj.GroupLabels, *ProtoToOsconfigBetaGuestPolicyAssignmentGroupLabels(r))
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
		obj.OSTypes = append(obj.OSTypes, *ProtoToOsconfigBetaGuestPolicyAssignmentOSTypes(r))
	}
	return obj
}

// ProtoToGuestPolicyAssignmentGroupLabels converts a GuestPolicyAssignmentGroupLabels object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyAssignmentGroupLabels(p *betapb.OsconfigBetaGuestPolicyAssignmentGroupLabels) *beta.GuestPolicyAssignmentGroupLabels {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyAssignmentGroupLabels{}
	return obj
}

// ProtoToGuestPolicyAssignmentOSTypes converts a GuestPolicyAssignmentOSTypes object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyAssignmentOSTypes(p *betapb.OsconfigBetaGuestPolicyAssignmentOSTypes) *beta.GuestPolicyAssignmentOSTypes {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyAssignmentOSTypes{
		OSShortName:    dcl.StringOrNil(p.GetOsShortName()),
		OSVersion:      dcl.StringOrNil(p.GetOsVersion()),
		OSArchitecture: dcl.StringOrNil(p.GetOsArchitecture()),
	}
	return obj
}

// ProtoToGuestPolicyPackages converts a GuestPolicyPackages object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackages(p *betapb.OsconfigBetaGuestPolicyPackages) *beta.GuestPolicyPackages {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackages{
		Name:         dcl.StringOrNil(p.GetName()),
		DesiredState: ProtoToOsconfigBetaGuestPolicyPackagesDesiredStateEnum(p.GetDesiredState()),
		Manager:      ProtoToOsconfigBetaGuestPolicyPackagesManagerEnum(p.GetManager()),
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositories converts a GuestPolicyPackageRepositories object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositories(p *betapb.OsconfigBetaGuestPolicyPackageRepositories) *beta.GuestPolicyPackageRepositories {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositories{
		Apt:    ProtoToOsconfigBetaGuestPolicyPackageRepositoriesApt(p.GetApt()),
		Yum:    ProtoToOsconfigBetaGuestPolicyPackageRepositoriesYum(p.GetYum()),
		Zypper: ProtoToOsconfigBetaGuestPolicyPackageRepositoriesZypper(p.GetZypper()),
		Goo:    ProtoToOsconfigBetaGuestPolicyPackageRepositoriesGoo(p.GetGoo()),
	}
	return obj
}

// ProtoToGuestPolicyPackageRepositoriesApt converts a GuestPolicyPackageRepositoriesApt object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesApt(p *betapb.OsconfigBetaGuestPolicyPackageRepositoriesApt) *beta.GuestPolicyPackageRepositoriesApt {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositoriesApt{
		ArchiveType:  ProtoToOsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(p.GetArchiveType()),
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
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesYum(p *betapb.OsconfigBetaGuestPolicyPackageRepositoriesYum) *beta.GuestPolicyPackageRepositoriesYum {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositoriesYum{
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
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesZypper(p *betapb.OsconfigBetaGuestPolicyPackageRepositoriesZypper) *beta.GuestPolicyPackageRepositoriesZypper {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositoriesZypper{
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
func ProtoToOsconfigBetaGuestPolicyPackageRepositoriesGoo(p *betapb.OsconfigBetaGuestPolicyPackageRepositoriesGoo) *beta.GuestPolicyPackageRepositoriesGoo {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyPackageRepositoriesGoo{
		Name: dcl.StringOrNil(p.GetName()),
		Url:  dcl.StringOrNil(p.GetUrl()),
	}
	return obj
}

// ProtoToGuestPolicyRecipes converts a GuestPolicyRecipes object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipes(p *betapb.OsconfigBetaGuestPolicyRecipes) *beta.GuestPolicyRecipes {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipes{
		Name:         dcl.StringOrNil(p.GetName()),
		Version:      dcl.StringOrNil(p.GetVersion()),
		DesiredState: ProtoToOsconfigBetaGuestPolicyRecipesDesiredStateEnum(p.GetDesiredState()),
	}
	for _, r := range p.GetArtifacts() {
		obj.Artifacts = append(obj.Artifacts, *ProtoToOsconfigBetaGuestPolicyRecipesArtifacts(r))
	}
	for _, r := range p.GetInstallSteps() {
		obj.InstallSteps = append(obj.InstallSteps, *ProtoToOsconfigBetaGuestPolicyRecipesInstallSteps(r))
	}
	for _, r := range p.GetUpdateSteps() {
		obj.UpdateSteps = append(obj.UpdateSteps, *ProtoToOsconfigBetaGuestPolicyRecipesUpdateSteps(r))
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifacts converts a GuestPolicyRecipesArtifacts object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesArtifacts(p *betapb.OsconfigBetaGuestPolicyRecipesArtifacts) *beta.GuestPolicyRecipesArtifacts {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesArtifacts{
		Id:            dcl.StringOrNil(p.GetId()),
		Remote:        ProtoToOsconfigBetaGuestPolicyRecipesArtifactsRemote(p.GetRemote()),
		Gcs:           ProtoToOsconfigBetaGuestPolicyRecipesArtifactsGcs(p.GetGcs()),
		AllowInsecure: dcl.Bool(p.GetAllowInsecure()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifactsRemote converts a GuestPolicyRecipesArtifactsRemote object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesArtifactsRemote(p *betapb.OsconfigBetaGuestPolicyRecipesArtifactsRemote) *beta.GuestPolicyRecipesArtifactsRemote {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesArtifactsRemote{
		Uri:      dcl.StringOrNil(p.GetUri()),
		Checksum: dcl.StringOrNil(p.GetChecksum()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesArtifactsGcs converts a GuestPolicyRecipesArtifactsGcs object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesArtifactsGcs(p *betapb.OsconfigBetaGuestPolicyRecipesArtifactsGcs) *beta.GuestPolicyRecipesArtifactsGcs {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesArtifactsGcs{
		Bucket:     dcl.StringOrNil(p.GetBucket()),
		Object:     dcl.StringOrNil(p.GetObject()),
		Generation: dcl.Int64OrNil(p.GetGeneration()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallSteps converts a GuestPolicyRecipesInstallSteps object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallSteps(p *betapb.OsconfigBetaGuestPolicyRecipesInstallSteps) *beta.GuestPolicyRecipesInstallSteps {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallSteps{
		FileCopy:          ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsFileCopy(p.GetFileCopy()),
		ArchiveExtraction: ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction(p.GetArchiveExtraction()),
		MsiInstallation:   ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation(p.GetMsiInstallation()),
		DpkgInstallation:  ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation(p.GetDpkgInstallation()),
		RpmInstallation:   ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation(p.GetRpmInstallation()),
		FileExec:          ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsFileExec(p.GetFileExec()),
		ScriptRun:         ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsScriptRun(p.GetScriptRun()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsFileCopy converts a GuestPolicyRecipesInstallStepsFileCopy object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsFileCopy(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileCopy) *beta.GuestPolicyRecipesInstallStepsFileCopy {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsFileCopy{
		ArtifactId:  dcl.StringOrNil(p.GetArtifactId()),
		Destination: dcl.StringOrNil(p.GetDestination()),
		Overwrite:   dcl.Bool(p.GetOverwrite()),
		Permissions: dcl.StringOrNil(p.GetPermissions()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsArchiveExtraction converts a GuestPolicyRecipesInstallStepsArchiveExtraction object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction) *beta.GuestPolicyRecipesInstallStepsArchiveExtraction {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsArchiveExtraction{
		ArtifactId:  dcl.StringOrNil(p.GetArtifactId()),
		Destination: dcl.StringOrNil(p.GetDestination()),
		Type:        ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsMsiInstallation converts a GuestPolicyRecipesInstallStepsMsiInstallation object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation) *beta.GuestPolicyRecipesInstallStepsMsiInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsMsiInstallation{
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
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation) *beta.GuestPolicyRecipesInstallStepsDpkgInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsDpkgInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsRpmInstallation converts a GuestPolicyRecipesInstallStepsRpmInstallation object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation) *beta.GuestPolicyRecipesInstallStepsRpmInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsRpmInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesInstallStepsFileExec converts a GuestPolicyRecipesInstallStepsFileExec object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsFileExec(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileExec) *beta.GuestPolicyRecipesInstallStepsFileExec {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsFileExec{
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
func ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsScriptRun(p *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRun) *beta.GuestPolicyRecipesInstallStepsScriptRun {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesInstallStepsScriptRun{
		Script:      dcl.StringOrNil(p.GetScript()),
		Interpreter: ProtoToOsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateSteps converts a GuestPolicyRecipesUpdateSteps object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateSteps(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateSteps) *beta.GuestPolicyRecipesUpdateSteps {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateSteps{
		FileCopy:          ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy(p.GetFileCopy()),
		ArchiveExtraction: ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction(p.GetArchiveExtraction()),
		MsiInstallation:   ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation(p.GetMsiInstallation()),
		DpkgInstallation:  ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation(p.GetDpkgInstallation()),
		RpmInstallation:   ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation(p.GetRpmInstallation()),
		FileExec:          ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsFileExec(p.GetFileExec()),
		ScriptRun:         ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun(p.GetScriptRun()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsFileCopy converts a GuestPolicyRecipesUpdateStepsFileCopy object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy) *beta.GuestPolicyRecipesUpdateStepsFileCopy {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsFileCopy{
		ArtifactId:  dcl.StringOrNil(p.GetArtifactId()),
		Destination: dcl.StringOrNil(p.GetDestination()),
		Overwrite:   dcl.Bool(p.GetOverwrite()),
		Permissions: dcl.StringOrNil(p.GetPermissions()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsArchiveExtraction converts a GuestPolicyRecipesUpdateStepsArchiveExtraction object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction) *beta.GuestPolicyRecipesUpdateStepsArchiveExtraction {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsArchiveExtraction{
		ArtifactId:  dcl.StringOrNil(p.GetArtifactId()),
		Destination: dcl.StringOrNil(p.GetDestination()),
		Type:        ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsMsiInstallation converts a GuestPolicyRecipesUpdateStepsMsiInstallation object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation) *beta.GuestPolicyRecipesUpdateStepsMsiInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsMsiInstallation{
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
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation) *beta.GuestPolicyRecipesUpdateStepsDpkgInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsDpkgInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsRpmInstallation converts a GuestPolicyRecipesUpdateStepsRpmInstallation object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation) *beta.GuestPolicyRecipesUpdateStepsRpmInstallation {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsRpmInstallation{
		ArtifactId: dcl.StringOrNil(p.GetArtifactId()),
	}
	return obj
}

// ProtoToGuestPolicyRecipesUpdateStepsFileExec converts a GuestPolicyRecipesUpdateStepsFileExec object from its proto representation.
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsFileExec(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileExec) *beta.GuestPolicyRecipesUpdateStepsFileExec {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsFileExec{
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
func ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun(p *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun) *beta.GuestPolicyRecipesUpdateStepsScriptRun {
	if p == nil {
		return nil
	}
	obj := &beta.GuestPolicyRecipesUpdateStepsScriptRun{
		Script:      dcl.StringOrNil(p.GetScript()),
		Interpreter: ProtoToOsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(p.GetInterpreter()),
	}
	for _, r := range p.GetAllowedExitCodes() {
		obj.AllowedExitCodes = append(obj.AllowedExitCodes, r)
	}
	return obj
}

// ProtoToGuestPolicy converts a GuestPolicy resource from its proto representation.
func ProtoToGuestPolicy(p *betapb.OsconfigBetaGuestPolicy) *beta.GuestPolicy {
	obj := &beta.GuestPolicy{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Assignment:  ProtoToOsconfigBetaGuestPolicyAssignment(p.GetAssignment()),
		Etag:        dcl.StringOrNil(p.GetEtag()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetPackages() {
		obj.Packages = append(obj.Packages, *ProtoToOsconfigBetaGuestPolicyPackages(r))
	}
	for _, r := range p.GetPackageRepositories() {
		obj.PackageRepositories = append(obj.PackageRepositories, *ProtoToOsconfigBetaGuestPolicyPackageRepositories(r))
	}
	for _, r := range p.GetRecipes() {
		obj.Recipes = append(obj.Recipes, *ProtoToOsconfigBetaGuestPolicyRecipes(r))
	}
	return obj
}

// GuestPolicyPackagesDesiredStateEnumToProto converts a GuestPolicyPackagesDesiredStateEnum enum to its proto representation.
func OsconfigBetaGuestPolicyPackagesDesiredStateEnumToProto(e *beta.GuestPolicyPackagesDesiredStateEnum) betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum_value["GuestPolicyPackagesDesiredStateEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyPackagesDesiredStateEnum(0)
}

// GuestPolicyPackagesManagerEnumToProto converts a GuestPolicyPackagesManagerEnum enum to its proto representation.
func OsconfigBetaGuestPolicyPackagesManagerEnumToProto(e *beta.GuestPolicyPackagesManagerEnum) betapb.OsconfigBetaGuestPolicyPackagesManagerEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyPackagesManagerEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyPackagesManagerEnum_value["GuestPolicyPackagesManagerEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyPackagesManagerEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyPackagesManagerEnum(0)
}

// GuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto converts a GuestPolicyPackageRepositoriesAptArchiveTypeEnum enum to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto(e *beta.GuestPolicyPackageRepositoriesAptArchiveTypeEnum) betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum_value["GuestPolicyPackageRepositoriesAptArchiveTypeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnum(0)
}

// GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto converts a GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto(e *beta.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum) betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum_value["GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum(0)
}

// GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto converts a GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto(e *beta.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum) betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum_value["GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnum(0)
}

// GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto converts a GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto(e *beta.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum) betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum_value["GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum(0)
}

// GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto converts a GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto(e *beta.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum) betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum_value["GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum(0)
}

// GuestPolicyRecipesDesiredStateEnumToProto converts a GuestPolicyRecipesDesiredStateEnum enum to its proto representation.
func OsconfigBetaGuestPolicyRecipesDesiredStateEnumToProto(e *beta.GuestPolicyRecipesDesiredStateEnum) betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum {
	if e == nil {
		return betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum(0)
	}
	if v, ok := betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum_value["GuestPolicyRecipesDesiredStateEnum"+string(*e)]; ok {
		return betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum(v)
	}
	return betapb.OsconfigBetaGuestPolicyRecipesDesiredStateEnum(0)
}

// GuestPolicyAssignmentToProto converts a GuestPolicyAssignment object to its proto representation.
func OsconfigBetaGuestPolicyAssignmentToProto(o *beta.GuestPolicyAssignment) *betapb.OsconfigBetaGuestPolicyAssignment {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyAssignment{}
	sGroupLabels := make([]*betapb.OsconfigBetaGuestPolicyAssignmentGroupLabels, len(o.GroupLabels))
	for i, r := range o.GroupLabels {
		sGroupLabels[i] = OsconfigBetaGuestPolicyAssignmentGroupLabelsToProto(&r)
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
	sOSTypes := make([]*betapb.OsconfigBetaGuestPolicyAssignmentOSTypes, len(o.OSTypes))
	for i, r := range o.OSTypes {
		sOSTypes[i] = OsconfigBetaGuestPolicyAssignmentOSTypesToProto(&r)
	}
	p.SetOsTypes(sOSTypes)
	return p
}

// GuestPolicyAssignmentGroupLabelsToProto converts a GuestPolicyAssignmentGroupLabels object to its proto representation.
func OsconfigBetaGuestPolicyAssignmentGroupLabelsToProto(o *beta.GuestPolicyAssignmentGroupLabels) *betapb.OsconfigBetaGuestPolicyAssignmentGroupLabels {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyAssignmentGroupLabels{}
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// GuestPolicyAssignmentOSTypesToProto converts a GuestPolicyAssignmentOSTypes object to its proto representation.
func OsconfigBetaGuestPolicyAssignmentOSTypesToProto(o *beta.GuestPolicyAssignmentOSTypes) *betapb.OsconfigBetaGuestPolicyAssignmentOSTypes {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyAssignmentOSTypes{}
	p.SetOsShortName(dcl.ValueOrEmptyString(o.OSShortName))
	p.SetOsVersion(dcl.ValueOrEmptyString(o.OSVersion))
	p.SetOsArchitecture(dcl.ValueOrEmptyString(o.OSArchitecture))
	return p
}

// GuestPolicyPackagesToProto converts a GuestPolicyPackages object to its proto representation.
func OsconfigBetaGuestPolicyPackagesToProto(o *beta.GuestPolicyPackages) *betapb.OsconfigBetaGuestPolicyPackages {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackages{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDesiredState(OsconfigBetaGuestPolicyPackagesDesiredStateEnumToProto(o.DesiredState))
	p.SetManager(OsconfigBetaGuestPolicyPackagesManagerEnumToProto(o.Manager))
	return p
}

// GuestPolicyPackageRepositoriesToProto converts a GuestPolicyPackageRepositories object to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesToProto(o *beta.GuestPolicyPackageRepositories) *betapb.OsconfigBetaGuestPolicyPackageRepositories {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositories{}
	p.SetApt(OsconfigBetaGuestPolicyPackageRepositoriesAptToProto(o.Apt))
	p.SetYum(OsconfigBetaGuestPolicyPackageRepositoriesYumToProto(o.Yum))
	p.SetZypper(OsconfigBetaGuestPolicyPackageRepositoriesZypperToProto(o.Zypper))
	p.SetGoo(OsconfigBetaGuestPolicyPackageRepositoriesGooToProto(o.Goo))
	return p
}

// GuestPolicyPackageRepositoriesAptToProto converts a GuestPolicyPackageRepositoriesApt object to its proto representation.
func OsconfigBetaGuestPolicyPackageRepositoriesAptToProto(o *beta.GuestPolicyPackageRepositoriesApt) *betapb.OsconfigBetaGuestPolicyPackageRepositoriesApt {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositoriesApt{}
	p.SetArchiveType(OsconfigBetaGuestPolicyPackageRepositoriesAptArchiveTypeEnumToProto(o.ArchiveType))
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
func OsconfigBetaGuestPolicyPackageRepositoriesYumToProto(o *beta.GuestPolicyPackageRepositoriesYum) *betapb.OsconfigBetaGuestPolicyPackageRepositoriesYum {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositoriesYum{}
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
func OsconfigBetaGuestPolicyPackageRepositoriesZypperToProto(o *beta.GuestPolicyPackageRepositoriesZypper) *betapb.OsconfigBetaGuestPolicyPackageRepositoriesZypper {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositoriesZypper{}
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
func OsconfigBetaGuestPolicyPackageRepositoriesGooToProto(o *beta.GuestPolicyPackageRepositoriesGoo) *betapb.OsconfigBetaGuestPolicyPackageRepositoriesGoo {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyPackageRepositoriesGoo{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetUrl(dcl.ValueOrEmptyString(o.Url))
	return p
}

// GuestPolicyRecipesToProto converts a GuestPolicyRecipes object to its proto representation.
func OsconfigBetaGuestPolicyRecipesToProto(o *beta.GuestPolicyRecipes) *betapb.OsconfigBetaGuestPolicyRecipes {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetDesiredState(OsconfigBetaGuestPolicyRecipesDesiredStateEnumToProto(o.DesiredState))
	sArtifacts := make([]*betapb.OsconfigBetaGuestPolicyRecipesArtifacts, len(o.Artifacts))
	for i, r := range o.Artifacts {
		sArtifacts[i] = OsconfigBetaGuestPolicyRecipesArtifactsToProto(&r)
	}
	p.SetArtifacts(sArtifacts)
	sInstallSteps := make([]*betapb.OsconfigBetaGuestPolicyRecipesInstallSteps, len(o.InstallSteps))
	for i, r := range o.InstallSteps {
		sInstallSteps[i] = OsconfigBetaGuestPolicyRecipesInstallStepsToProto(&r)
	}
	p.SetInstallSteps(sInstallSteps)
	sUpdateSteps := make([]*betapb.OsconfigBetaGuestPolicyRecipesUpdateSteps, len(o.UpdateSteps))
	for i, r := range o.UpdateSteps {
		sUpdateSteps[i] = OsconfigBetaGuestPolicyRecipesUpdateStepsToProto(&r)
	}
	p.SetUpdateSteps(sUpdateSteps)
	return p
}

// GuestPolicyRecipesArtifactsToProto converts a GuestPolicyRecipesArtifacts object to its proto representation.
func OsconfigBetaGuestPolicyRecipesArtifactsToProto(o *beta.GuestPolicyRecipesArtifacts) *betapb.OsconfigBetaGuestPolicyRecipesArtifacts {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesArtifacts{}
	p.SetId(dcl.ValueOrEmptyString(o.Id))
	p.SetRemote(OsconfigBetaGuestPolicyRecipesArtifactsRemoteToProto(o.Remote))
	p.SetGcs(OsconfigBetaGuestPolicyRecipesArtifactsGcsToProto(o.Gcs))
	p.SetAllowInsecure(dcl.ValueOrEmptyBool(o.AllowInsecure))
	return p
}

// GuestPolicyRecipesArtifactsRemoteToProto converts a GuestPolicyRecipesArtifactsRemote object to its proto representation.
func OsconfigBetaGuestPolicyRecipesArtifactsRemoteToProto(o *beta.GuestPolicyRecipesArtifactsRemote) *betapb.OsconfigBetaGuestPolicyRecipesArtifactsRemote {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesArtifactsRemote{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	p.SetChecksum(dcl.ValueOrEmptyString(o.Checksum))
	return p
}

// GuestPolicyRecipesArtifactsGcsToProto converts a GuestPolicyRecipesArtifactsGcs object to its proto representation.
func OsconfigBetaGuestPolicyRecipesArtifactsGcsToProto(o *beta.GuestPolicyRecipesArtifactsGcs) *betapb.OsconfigBetaGuestPolicyRecipesArtifactsGcs {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesArtifactsGcs{}
	p.SetBucket(dcl.ValueOrEmptyString(o.Bucket))
	p.SetObject(dcl.ValueOrEmptyString(o.Object))
	p.SetGeneration(dcl.ValueOrEmptyInt64(o.Generation))
	return p
}

// GuestPolicyRecipesInstallStepsToProto converts a GuestPolicyRecipesInstallSteps object to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsToProto(o *beta.GuestPolicyRecipesInstallSteps) *betapb.OsconfigBetaGuestPolicyRecipesInstallSteps {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallSteps{}
	p.SetFileCopy(OsconfigBetaGuestPolicyRecipesInstallStepsFileCopyToProto(o.FileCopy))
	p.SetArchiveExtraction(OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionToProto(o.ArchiveExtraction))
	p.SetMsiInstallation(OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallationToProto(o.MsiInstallation))
	p.SetDpkgInstallation(OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallationToProto(o.DpkgInstallation))
	p.SetRpmInstallation(OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallationToProto(o.RpmInstallation))
	p.SetFileExec(OsconfigBetaGuestPolicyRecipesInstallStepsFileExecToProto(o.FileExec))
	p.SetScriptRun(OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunToProto(o.ScriptRun))
	return p
}

// GuestPolicyRecipesInstallStepsFileCopyToProto converts a GuestPolicyRecipesInstallStepsFileCopy object to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsFileCopyToProto(o *beta.GuestPolicyRecipesInstallStepsFileCopy) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileCopy {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileCopy{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetDestination(dcl.ValueOrEmptyString(o.Destination))
	p.SetOverwrite(dcl.ValueOrEmptyBool(o.Overwrite))
	p.SetPermissions(dcl.ValueOrEmptyString(o.Permissions))
	return p
}

// GuestPolicyRecipesInstallStepsArchiveExtractionToProto converts a GuestPolicyRecipesInstallStepsArchiveExtraction object to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionToProto(o *beta.GuestPolicyRecipesInstallStepsArchiveExtraction) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtraction{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetDestination(dcl.ValueOrEmptyString(o.Destination))
	p.SetType(OsconfigBetaGuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumToProto(o.Type))
	return p
}

// GuestPolicyRecipesInstallStepsMsiInstallationToProto converts a GuestPolicyRecipesInstallStepsMsiInstallation object to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallationToProto(o *beta.GuestPolicyRecipesInstallStepsMsiInstallation) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsMsiInstallation{}
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
func OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallationToProto(o *beta.GuestPolicyRecipesInstallStepsDpkgInstallation) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsDpkgInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	return p
}

// GuestPolicyRecipesInstallStepsRpmInstallationToProto converts a GuestPolicyRecipesInstallStepsRpmInstallation object to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallationToProto(o *beta.GuestPolicyRecipesInstallStepsRpmInstallation) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsRpmInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	return p
}

// GuestPolicyRecipesInstallStepsFileExecToProto converts a GuestPolicyRecipesInstallStepsFileExec object to its proto representation.
func OsconfigBetaGuestPolicyRecipesInstallStepsFileExecToProto(o *beta.GuestPolicyRecipesInstallStepsFileExec) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileExec {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsFileExec{}
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
func OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunToProto(o *beta.GuestPolicyRecipesInstallStepsScriptRun) *betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRun {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesInstallStepsScriptRun{}
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigBetaGuestPolicyRecipesInstallStepsScriptRunInterpreterEnumToProto(o.Interpreter))
	sAllowedExitCodes := make([]int64, len(o.AllowedExitCodes))
	for i, r := range o.AllowedExitCodes {
		sAllowedExitCodes[i] = r
	}
	p.SetAllowedExitCodes(sAllowedExitCodes)
	return p
}

// GuestPolicyRecipesUpdateStepsToProto converts a GuestPolicyRecipesUpdateSteps object to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsToProto(o *beta.GuestPolicyRecipesUpdateSteps) *betapb.OsconfigBetaGuestPolicyRecipesUpdateSteps {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateSteps{}
	p.SetFileCopy(OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopyToProto(o.FileCopy))
	p.SetArchiveExtraction(OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionToProto(o.ArchiveExtraction))
	p.SetMsiInstallation(OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallationToProto(o.MsiInstallation))
	p.SetDpkgInstallation(OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallationToProto(o.DpkgInstallation))
	p.SetRpmInstallation(OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallationToProto(o.RpmInstallation))
	p.SetFileExec(OsconfigBetaGuestPolicyRecipesUpdateStepsFileExecToProto(o.FileExec))
	p.SetScriptRun(OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunToProto(o.ScriptRun))
	return p
}

// GuestPolicyRecipesUpdateStepsFileCopyToProto converts a GuestPolicyRecipesUpdateStepsFileCopy object to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopyToProto(o *beta.GuestPolicyRecipesUpdateStepsFileCopy) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileCopy{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetDestination(dcl.ValueOrEmptyString(o.Destination))
	p.SetOverwrite(dcl.ValueOrEmptyBool(o.Overwrite))
	p.SetPermissions(dcl.ValueOrEmptyString(o.Permissions))
	return p
}

// GuestPolicyRecipesUpdateStepsArchiveExtractionToProto converts a GuestPolicyRecipesUpdateStepsArchiveExtraction object to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionToProto(o *beta.GuestPolicyRecipesUpdateStepsArchiveExtraction) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtraction{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	p.SetDestination(dcl.ValueOrEmptyString(o.Destination))
	p.SetType(OsconfigBetaGuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumToProto(o.Type))
	return p
}

// GuestPolicyRecipesUpdateStepsMsiInstallationToProto converts a GuestPolicyRecipesUpdateStepsMsiInstallation object to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallationToProto(o *beta.GuestPolicyRecipesUpdateStepsMsiInstallation) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsMsiInstallation{}
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
func OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallationToProto(o *beta.GuestPolicyRecipesUpdateStepsDpkgInstallation) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsDpkgInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	return p
}

// GuestPolicyRecipesUpdateStepsRpmInstallationToProto converts a GuestPolicyRecipesUpdateStepsRpmInstallation object to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallationToProto(o *beta.GuestPolicyRecipesUpdateStepsRpmInstallation) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsRpmInstallation{}
	p.SetArtifactId(dcl.ValueOrEmptyString(o.ArtifactId))
	return p
}

// GuestPolicyRecipesUpdateStepsFileExecToProto converts a GuestPolicyRecipesUpdateStepsFileExec object to its proto representation.
func OsconfigBetaGuestPolicyRecipesUpdateStepsFileExecToProto(o *beta.GuestPolicyRecipesUpdateStepsFileExec) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileExec {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsFileExec{}
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
func OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunToProto(o *beta.GuestPolicyRecipesUpdateStepsScriptRun) *betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun {
	if o == nil {
		return nil
	}
	p := &betapb.OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRun{}
	p.SetScript(dcl.ValueOrEmptyString(o.Script))
	p.SetInterpreter(OsconfigBetaGuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumToProto(o.Interpreter))
	sAllowedExitCodes := make([]int64, len(o.AllowedExitCodes))
	for i, r := range o.AllowedExitCodes {
		sAllowedExitCodes[i] = r
	}
	p.SetAllowedExitCodes(sAllowedExitCodes)
	return p
}

// GuestPolicyToProto converts a GuestPolicy resource to its proto representation.
func GuestPolicyToProto(resource *beta.GuestPolicy) *betapb.OsconfigBetaGuestPolicy {
	p := &betapb.OsconfigBetaGuestPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetAssignment(OsconfigBetaGuestPolicyAssignmentToProto(resource.Assignment))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sPackages := make([]*betapb.OsconfigBetaGuestPolicyPackages, len(resource.Packages))
	for i, r := range resource.Packages {
		sPackages[i] = OsconfigBetaGuestPolicyPackagesToProto(&r)
	}
	p.SetPackages(sPackages)
	sPackageRepositories := make([]*betapb.OsconfigBetaGuestPolicyPackageRepositories, len(resource.PackageRepositories))
	for i, r := range resource.PackageRepositories {
		sPackageRepositories[i] = OsconfigBetaGuestPolicyPackageRepositoriesToProto(&r)
	}
	p.SetPackageRepositories(sPackageRepositories)
	sRecipes := make([]*betapb.OsconfigBetaGuestPolicyRecipes, len(resource.Recipes))
	for i, r := range resource.Recipes {
		sRecipes[i] = OsconfigBetaGuestPolicyRecipesToProto(&r)
	}
	p.SetRecipes(sRecipes)

	return p
}

// applyGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Apply() method.
func (s *GuestPolicyServer) applyGuestPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyOsconfigBetaGuestPolicyRequest) (*betapb.OsconfigBetaGuestPolicy, error) {
	p := ProtoToGuestPolicy(request.GetResource())
	res, err := c.ApplyGuestPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GuestPolicyToProto(res)
	return r, nil
}

// applyOsconfigBetaGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Apply() method.
func (s *GuestPolicyServer) ApplyOsconfigBetaGuestPolicy(ctx context.Context, request *betapb.ApplyOsconfigBetaGuestPolicyRequest) (*betapb.OsconfigBetaGuestPolicy, error) {
	cl, err := createConfigGuestPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGuestPolicy(ctx, cl, request)
}

// DeleteGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicy Delete() method.
func (s *GuestPolicyServer) DeleteOsconfigBetaGuestPolicy(ctx context.Context, request *betapb.DeleteOsconfigBetaGuestPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGuestPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGuestPolicy(ctx, ProtoToGuestPolicy(request.GetResource()))

}

// ListOsconfigBetaGuestPolicy handles the gRPC request by passing it to the underlying GuestPolicyList() method.
func (s *GuestPolicyServer) ListOsconfigBetaGuestPolicy(ctx context.Context, request *betapb.ListOsconfigBetaGuestPolicyRequest) (*betapb.ListOsconfigBetaGuestPolicyResponse, error) {
	cl, err := createConfigGuestPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGuestPolicy(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.OsconfigBetaGuestPolicy
	for _, r := range resources.Items {
		rp := GuestPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListOsconfigBetaGuestPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGuestPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
