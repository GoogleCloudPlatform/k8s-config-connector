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
package osconfig

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/osconfig/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type GuestPolicy struct{}

func GuestPolicyToUnstructured(r *dclService.GuestPolicy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "osconfig",
			Version: "alpha",
			Type:    "GuestPolicy",
		},
		Object: make(map[string]interface{}),
	}
	if r.Assignment != nil && r.Assignment != dclService.EmptyGuestPolicyAssignment {
		rAssignment := make(map[string]interface{})
		var rAssignmentGroupLabels []interface{}
		for _, rAssignmentGroupLabelsVal := range r.Assignment.GroupLabels {
			rAssignmentGroupLabelsObject := make(map[string]interface{})
			if rAssignmentGroupLabelsVal.Labels != nil {
				rAssignmentGroupLabelsValLabels := make(map[string]interface{})
				for k, v := range rAssignmentGroupLabelsVal.Labels {
					rAssignmentGroupLabelsValLabels[k] = v
				}
				rAssignmentGroupLabelsObject["labels"] = rAssignmentGroupLabelsValLabels
			}
			rAssignmentGroupLabels = append(rAssignmentGroupLabels, rAssignmentGroupLabelsObject)
		}
		rAssignment["groupLabels"] = rAssignmentGroupLabels
		var rAssignmentInstanceNamePrefixes []interface{}
		for _, rAssignmentInstanceNamePrefixesVal := range r.Assignment.InstanceNamePrefixes {
			rAssignmentInstanceNamePrefixes = append(rAssignmentInstanceNamePrefixes, rAssignmentInstanceNamePrefixesVal)
		}
		rAssignment["instanceNamePrefixes"] = rAssignmentInstanceNamePrefixes
		var rAssignmentInstances []interface{}
		for _, rAssignmentInstancesVal := range r.Assignment.Instances {
			rAssignmentInstances = append(rAssignmentInstances, rAssignmentInstancesVal)
		}
		rAssignment["instances"] = rAssignmentInstances
		var rAssignmentOSTypes []interface{}
		for _, rAssignmentOSTypesVal := range r.Assignment.OSTypes {
			rAssignmentOSTypesObject := make(map[string]interface{})
			if rAssignmentOSTypesVal.OSArchitecture != nil {
				rAssignmentOSTypesObject["osArchitecture"] = *rAssignmentOSTypesVal.OSArchitecture
			}
			if rAssignmentOSTypesVal.OSShortName != nil {
				rAssignmentOSTypesObject["osShortName"] = *rAssignmentOSTypesVal.OSShortName
			}
			if rAssignmentOSTypesVal.OSVersion != nil {
				rAssignmentOSTypesObject["osVersion"] = *rAssignmentOSTypesVal.OSVersion
			}
			rAssignmentOSTypes = append(rAssignmentOSTypes, rAssignmentOSTypesObject)
		}
		rAssignment["osTypes"] = rAssignmentOSTypes
		var rAssignmentZones []interface{}
		for _, rAssignmentZonesVal := range r.Assignment.Zones {
			rAssignmentZones = append(rAssignmentZones, rAssignmentZonesVal)
		}
		rAssignment["zones"] = rAssignmentZones
		u.Object["assignment"] = rAssignment
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	var rPackageRepositories []interface{}
	for _, rPackageRepositoriesVal := range r.PackageRepositories {
		rPackageRepositoriesObject := make(map[string]interface{})
		if rPackageRepositoriesVal.Apt != nil && rPackageRepositoriesVal.Apt != dclService.EmptyGuestPolicyPackageRepositoriesApt {
			rPackageRepositoriesValApt := make(map[string]interface{})
			if rPackageRepositoriesVal.Apt.ArchiveType != nil {
				rPackageRepositoriesValApt["archiveType"] = string(*rPackageRepositoriesVal.Apt.ArchiveType)
			}
			var rPackageRepositoriesValAptComponents []interface{}
			for _, rPackageRepositoriesValAptComponentsVal := range rPackageRepositoriesVal.Apt.Components {
				rPackageRepositoriesValAptComponents = append(rPackageRepositoriesValAptComponents, rPackageRepositoriesValAptComponentsVal)
			}
			rPackageRepositoriesValApt["components"] = rPackageRepositoriesValAptComponents
			if rPackageRepositoriesVal.Apt.Distribution != nil {
				rPackageRepositoriesValApt["distribution"] = *rPackageRepositoriesVal.Apt.Distribution
			}
			if rPackageRepositoriesVal.Apt.GpgKey != nil {
				rPackageRepositoriesValApt["gpgKey"] = *rPackageRepositoriesVal.Apt.GpgKey
			}
			if rPackageRepositoriesVal.Apt.Uri != nil {
				rPackageRepositoriesValApt["uri"] = *rPackageRepositoriesVal.Apt.Uri
			}
			rPackageRepositoriesObject["apt"] = rPackageRepositoriesValApt
		}
		if rPackageRepositoriesVal.Goo != nil && rPackageRepositoriesVal.Goo != dclService.EmptyGuestPolicyPackageRepositoriesGoo {
			rPackageRepositoriesValGoo := make(map[string]interface{})
			if rPackageRepositoriesVal.Goo.Name != nil {
				rPackageRepositoriesValGoo["name"] = *rPackageRepositoriesVal.Goo.Name
			}
			if rPackageRepositoriesVal.Goo.Url != nil {
				rPackageRepositoriesValGoo["url"] = *rPackageRepositoriesVal.Goo.Url
			}
			rPackageRepositoriesObject["goo"] = rPackageRepositoriesValGoo
		}
		if rPackageRepositoriesVal.Yum != nil && rPackageRepositoriesVal.Yum != dclService.EmptyGuestPolicyPackageRepositoriesYum {
			rPackageRepositoriesValYum := make(map[string]interface{})
			if rPackageRepositoriesVal.Yum.BaseUrl != nil {
				rPackageRepositoriesValYum["baseUrl"] = *rPackageRepositoriesVal.Yum.BaseUrl
			}
			if rPackageRepositoriesVal.Yum.DisplayName != nil {
				rPackageRepositoriesValYum["displayName"] = *rPackageRepositoriesVal.Yum.DisplayName
			}
			var rPackageRepositoriesValYumGpgKeys []interface{}
			for _, rPackageRepositoriesValYumGpgKeysVal := range rPackageRepositoriesVal.Yum.GpgKeys {
				rPackageRepositoriesValYumGpgKeys = append(rPackageRepositoriesValYumGpgKeys, rPackageRepositoriesValYumGpgKeysVal)
			}
			rPackageRepositoriesValYum["gpgKeys"] = rPackageRepositoriesValYumGpgKeys
			if rPackageRepositoriesVal.Yum.Id != nil {
				rPackageRepositoriesValYum["id"] = *rPackageRepositoriesVal.Yum.Id
			}
			rPackageRepositoriesObject["yum"] = rPackageRepositoriesValYum
		}
		if rPackageRepositoriesVal.Zypper != nil && rPackageRepositoriesVal.Zypper != dclService.EmptyGuestPolicyPackageRepositoriesZypper {
			rPackageRepositoriesValZypper := make(map[string]interface{})
			if rPackageRepositoriesVal.Zypper.BaseUrl != nil {
				rPackageRepositoriesValZypper["baseUrl"] = *rPackageRepositoriesVal.Zypper.BaseUrl
			}
			if rPackageRepositoriesVal.Zypper.DisplayName != nil {
				rPackageRepositoriesValZypper["displayName"] = *rPackageRepositoriesVal.Zypper.DisplayName
			}
			var rPackageRepositoriesValZypperGpgKeys []interface{}
			for _, rPackageRepositoriesValZypperGpgKeysVal := range rPackageRepositoriesVal.Zypper.GpgKeys {
				rPackageRepositoriesValZypperGpgKeys = append(rPackageRepositoriesValZypperGpgKeys, rPackageRepositoriesValZypperGpgKeysVal)
			}
			rPackageRepositoriesValZypper["gpgKeys"] = rPackageRepositoriesValZypperGpgKeys
			if rPackageRepositoriesVal.Zypper.Id != nil {
				rPackageRepositoriesValZypper["id"] = *rPackageRepositoriesVal.Zypper.Id
			}
			rPackageRepositoriesObject["zypper"] = rPackageRepositoriesValZypper
		}
		rPackageRepositories = append(rPackageRepositories, rPackageRepositoriesObject)
	}
	u.Object["packageRepositories"] = rPackageRepositories
	var rPackages []interface{}
	for _, rPackagesVal := range r.Packages {
		rPackagesObject := make(map[string]interface{})
		if rPackagesVal.DesiredState != nil {
			rPackagesObject["desiredState"] = string(*rPackagesVal.DesiredState)
		}
		if rPackagesVal.Manager != nil {
			rPackagesObject["manager"] = string(*rPackagesVal.Manager)
		}
		if rPackagesVal.Name != nil {
			rPackagesObject["name"] = *rPackagesVal.Name
		}
		rPackages = append(rPackages, rPackagesObject)
	}
	u.Object["packages"] = rPackages
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rRecipes []interface{}
	for _, rRecipesVal := range r.Recipes {
		rRecipesObject := make(map[string]interface{})
		var rRecipesValArtifacts []interface{}
		for _, rRecipesValArtifactsVal := range rRecipesVal.Artifacts {
			rRecipesValArtifactsObject := make(map[string]interface{})
			if rRecipesValArtifactsVal.AllowInsecure != nil {
				rRecipesValArtifactsObject["allowInsecure"] = *rRecipesValArtifactsVal.AllowInsecure
			}
			if rRecipesValArtifactsVal.Gcs != nil && rRecipesValArtifactsVal.Gcs != dclService.EmptyGuestPolicyRecipesArtifactsGcs {
				rRecipesValArtifactsValGcs := make(map[string]interface{})
				if rRecipesValArtifactsVal.Gcs.Bucket != nil {
					rRecipesValArtifactsValGcs["bucket"] = *rRecipesValArtifactsVal.Gcs.Bucket
				}
				if rRecipesValArtifactsVal.Gcs.Generation != nil {
					rRecipesValArtifactsValGcs["generation"] = *rRecipesValArtifactsVal.Gcs.Generation
				}
				if rRecipesValArtifactsVal.Gcs.Object != nil {
					rRecipesValArtifactsValGcs["object"] = *rRecipesValArtifactsVal.Gcs.Object
				}
				rRecipesValArtifactsObject["gcs"] = rRecipesValArtifactsValGcs
			}
			if rRecipesValArtifactsVal.Id != nil {
				rRecipesValArtifactsObject["id"] = *rRecipesValArtifactsVal.Id
			}
			if rRecipesValArtifactsVal.Remote != nil && rRecipesValArtifactsVal.Remote != dclService.EmptyGuestPolicyRecipesArtifactsRemote {
				rRecipesValArtifactsValRemote := make(map[string]interface{})
				if rRecipesValArtifactsVal.Remote.Checksum != nil {
					rRecipesValArtifactsValRemote["checksum"] = *rRecipesValArtifactsVal.Remote.Checksum
				}
				if rRecipesValArtifactsVal.Remote.Uri != nil {
					rRecipesValArtifactsValRemote["uri"] = *rRecipesValArtifactsVal.Remote.Uri
				}
				rRecipesValArtifactsObject["remote"] = rRecipesValArtifactsValRemote
			}
			rRecipesValArtifacts = append(rRecipesValArtifacts, rRecipesValArtifactsObject)
		}
		rRecipesObject["artifacts"] = rRecipesValArtifacts
		if rRecipesVal.DesiredState != nil {
			rRecipesObject["desiredState"] = string(*rRecipesVal.DesiredState)
		}
		var rRecipesValInstallSteps []interface{}
		for _, rRecipesValInstallStepsVal := range rRecipesVal.InstallSteps {
			rRecipesValInstallStepsObject := make(map[string]interface{})
			if rRecipesValInstallStepsVal.ArchiveExtraction != nil && rRecipesValInstallStepsVal.ArchiveExtraction != dclService.EmptyGuestPolicyRecipesInstallStepsArchiveExtraction {
				rRecipesValInstallStepsValArchiveExtraction := make(map[string]interface{})
				if rRecipesValInstallStepsVal.ArchiveExtraction.ArtifactId != nil {
					rRecipesValInstallStepsValArchiveExtraction["artifactId"] = *rRecipesValInstallStepsVal.ArchiveExtraction.ArtifactId
				}
				if rRecipesValInstallStepsVal.ArchiveExtraction.Destination != nil {
					rRecipesValInstallStepsValArchiveExtraction["destination"] = *rRecipesValInstallStepsVal.ArchiveExtraction.Destination
				}
				if rRecipesValInstallStepsVal.ArchiveExtraction.Type != nil {
					rRecipesValInstallStepsValArchiveExtraction["type"] = string(*rRecipesValInstallStepsVal.ArchiveExtraction.Type)
				}
				rRecipesValInstallStepsObject["archiveExtraction"] = rRecipesValInstallStepsValArchiveExtraction
			}
			if rRecipesValInstallStepsVal.DpkgInstallation != nil && rRecipesValInstallStepsVal.DpkgInstallation != dclService.EmptyGuestPolicyRecipesInstallStepsDpkgInstallation {
				rRecipesValInstallStepsValDpkgInstallation := make(map[string]interface{})
				if rRecipesValInstallStepsVal.DpkgInstallation.ArtifactId != nil {
					rRecipesValInstallStepsValDpkgInstallation["artifactId"] = *rRecipesValInstallStepsVal.DpkgInstallation.ArtifactId
				}
				rRecipesValInstallStepsObject["dpkgInstallation"] = rRecipesValInstallStepsValDpkgInstallation
			}
			if rRecipesValInstallStepsVal.FileCopy != nil && rRecipesValInstallStepsVal.FileCopy != dclService.EmptyGuestPolicyRecipesInstallStepsFileCopy {
				rRecipesValInstallStepsValFileCopy := make(map[string]interface{})
				if rRecipesValInstallStepsVal.FileCopy.ArtifactId != nil {
					rRecipesValInstallStepsValFileCopy["artifactId"] = *rRecipesValInstallStepsVal.FileCopy.ArtifactId
				}
				if rRecipesValInstallStepsVal.FileCopy.Destination != nil {
					rRecipesValInstallStepsValFileCopy["destination"] = *rRecipesValInstallStepsVal.FileCopy.Destination
				}
				if rRecipesValInstallStepsVal.FileCopy.Overwrite != nil {
					rRecipesValInstallStepsValFileCopy["overwrite"] = *rRecipesValInstallStepsVal.FileCopy.Overwrite
				}
				if rRecipesValInstallStepsVal.FileCopy.Permissions != nil {
					rRecipesValInstallStepsValFileCopy["permissions"] = *rRecipesValInstallStepsVal.FileCopy.Permissions
				}
				rRecipesValInstallStepsObject["fileCopy"] = rRecipesValInstallStepsValFileCopy
			}
			if rRecipesValInstallStepsVal.FileExec != nil && rRecipesValInstallStepsVal.FileExec != dclService.EmptyGuestPolicyRecipesInstallStepsFileExec {
				rRecipesValInstallStepsValFileExec := make(map[string]interface{})
				var rRecipesValInstallStepsValFileExecAllowedExitCodes []interface{}
				for _, rRecipesValInstallStepsValFileExecAllowedExitCodesVal := range rRecipesValInstallStepsVal.FileExec.AllowedExitCodes {
					rRecipesValInstallStepsValFileExecAllowedExitCodes = append(rRecipesValInstallStepsValFileExecAllowedExitCodes, rRecipesValInstallStepsValFileExecAllowedExitCodesVal)
				}
				rRecipesValInstallStepsValFileExec["allowedExitCodes"] = rRecipesValInstallStepsValFileExecAllowedExitCodes
				var rRecipesValInstallStepsValFileExecArgs []interface{}
				for _, rRecipesValInstallStepsValFileExecArgsVal := range rRecipesValInstallStepsVal.FileExec.Args {
					rRecipesValInstallStepsValFileExecArgs = append(rRecipesValInstallStepsValFileExecArgs, rRecipesValInstallStepsValFileExecArgsVal)
				}
				rRecipesValInstallStepsValFileExec["args"] = rRecipesValInstallStepsValFileExecArgs
				if rRecipesValInstallStepsVal.FileExec.ArtifactId != nil {
					rRecipesValInstallStepsValFileExec["artifactId"] = *rRecipesValInstallStepsVal.FileExec.ArtifactId
				}
				if rRecipesValInstallStepsVal.FileExec.LocalPath != nil {
					rRecipesValInstallStepsValFileExec["localPath"] = *rRecipesValInstallStepsVal.FileExec.LocalPath
				}
				rRecipesValInstallStepsObject["fileExec"] = rRecipesValInstallStepsValFileExec
			}
			if rRecipesValInstallStepsVal.MsiInstallation != nil && rRecipesValInstallStepsVal.MsiInstallation != dclService.EmptyGuestPolicyRecipesInstallStepsMsiInstallation {
				rRecipesValInstallStepsValMsiInstallation := make(map[string]interface{})
				var rRecipesValInstallStepsValMsiInstallationAllowedExitCodes []interface{}
				for _, rRecipesValInstallStepsValMsiInstallationAllowedExitCodesVal := range rRecipesValInstallStepsVal.MsiInstallation.AllowedExitCodes {
					rRecipesValInstallStepsValMsiInstallationAllowedExitCodes = append(rRecipesValInstallStepsValMsiInstallationAllowedExitCodes, rRecipesValInstallStepsValMsiInstallationAllowedExitCodesVal)
				}
				rRecipesValInstallStepsValMsiInstallation["allowedExitCodes"] = rRecipesValInstallStepsValMsiInstallationAllowedExitCodes
				if rRecipesValInstallStepsVal.MsiInstallation.ArtifactId != nil {
					rRecipesValInstallStepsValMsiInstallation["artifactId"] = *rRecipesValInstallStepsVal.MsiInstallation.ArtifactId
				}
				var rRecipesValInstallStepsValMsiInstallationFlags []interface{}
				for _, rRecipesValInstallStepsValMsiInstallationFlagsVal := range rRecipesValInstallStepsVal.MsiInstallation.Flags {
					rRecipesValInstallStepsValMsiInstallationFlags = append(rRecipesValInstallStepsValMsiInstallationFlags, rRecipesValInstallStepsValMsiInstallationFlagsVal)
				}
				rRecipesValInstallStepsValMsiInstallation["flags"] = rRecipesValInstallStepsValMsiInstallationFlags
				rRecipesValInstallStepsObject["msiInstallation"] = rRecipesValInstallStepsValMsiInstallation
			}
			if rRecipesValInstallStepsVal.RpmInstallation != nil && rRecipesValInstallStepsVal.RpmInstallation != dclService.EmptyGuestPolicyRecipesInstallStepsRpmInstallation {
				rRecipesValInstallStepsValRpmInstallation := make(map[string]interface{})
				if rRecipesValInstallStepsVal.RpmInstallation.ArtifactId != nil {
					rRecipesValInstallStepsValRpmInstallation["artifactId"] = *rRecipesValInstallStepsVal.RpmInstallation.ArtifactId
				}
				rRecipesValInstallStepsObject["rpmInstallation"] = rRecipesValInstallStepsValRpmInstallation
			}
			if rRecipesValInstallStepsVal.ScriptRun != nil && rRecipesValInstallStepsVal.ScriptRun != dclService.EmptyGuestPolicyRecipesInstallStepsScriptRun {
				rRecipesValInstallStepsValScriptRun := make(map[string]interface{})
				var rRecipesValInstallStepsValScriptRunAllowedExitCodes []interface{}
				for _, rRecipesValInstallStepsValScriptRunAllowedExitCodesVal := range rRecipesValInstallStepsVal.ScriptRun.AllowedExitCodes {
					rRecipesValInstallStepsValScriptRunAllowedExitCodes = append(rRecipesValInstallStepsValScriptRunAllowedExitCodes, rRecipesValInstallStepsValScriptRunAllowedExitCodesVal)
				}
				rRecipesValInstallStepsValScriptRun["allowedExitCodes"] = rRecipesValInstallStepsValScriptRunAllowedExitCodes
				if rRecipesValInstallStepsVal.ScriptRun.Interpreter != nil {
					rRecipesValInstallStepsValScriptRun["interpreter"] = string(*rRecipesValInstallStepsVal.ScriptRun.Interpreter)
				}
				if rRecipesValInstallStepsVal.ScriptRun.Script != nil {
					rRecipesValInstallStepsValScriptRun["script"] = *rRecipesValInstallStepsVal.ScriptRun.Script
				}
				rRecipesValInstallStepsObject["scriptRun"] = rRecipesValInstallStepsValScriptRun
			}
			rRecipesValInstallSteps = append(rRecipesValInstallSteps, rRecipesValInstallStepsObject)
		}
		rRecipesObject["installSteps"] = rRecipesValInstallSteps
		if rRecipesVal.Name != nil {
			rRecipesObject["name"] = *rRecipesVal.Name
		}
		var rRecipesValUpdateSteps []interface{}
		for _, rRecipesValUpdateStepsVal := range rRecipesVal.UpdateSteps {
			rRecipesValUpdateStepsObject := make(map[string]interface{})
			if rRecipesValUpdateStepsVal.ArchiveExtraction != nil && rRecipesValUpdateStepsVal.ArchiveExtraction != dclService.EmptyGuestPolicyRecipesUpdateStepsArchiveExtraction {
				rRecipesValUpdateStepsValArchiveExtraction := make(map[string]interface{})
				if rRecipesValUpdateStepsVal.ArchiveExtraction.ArtifactId != nil {
					rRecipesValUpdateStepsValArchiveExtraction["artifactId"] = *rRecipesValUpdateStepsVal.ArchiveExtraction.ArtifactId
				}
				if rRecipesValUpdateStepsVal.ArchiveExtraction.Destination != nil {
					rRecipesValUpdateStepsValArchiveExtraction["destination"] = *rRecipesValUpdateStepsVal.ArchiveExtraction.Destination
				}
				if rRecipesValUpdateStepsVal.ArchiveExtraction.Type != nil {
					rRecipesValUpdateStepsValArchiveExtraction["type"] = string(*rRecipesValUpdateStepsVal.ArchiveExtraction.Type)
				}
				rRecipesValUpdateStepsObject["archiveExtraction"] = rRecipesValUpdateStepsValArchiveExtraction
			}
			if rRecipesValUpdateStepsVal.DpkgInstallation != nil && rRecipesValUpdateStepsVal.DpkgInstallation != dclService.EmptyGuestPolicyRecipesUpdateStepsDpkgInstallation {
				rRecipesValUpdateStepsValDpkgInstallation := make(map[string]interface{})
				if rRecipesValUpdateStepsVal.DpkgInstallation.ArtifactId != nil {
					rRecipesValUpdateStepsValDpkgInstallation["artifactId"] = *rRecipesValUpdateStepsVal.DpkgInstallation.ArtifactId
				}
				rRecipesValUpdateStepsObject["dpkgInstallation"] = rRecipesValUpdateStepsValDpkgInstallation
			}
			if rRecipesValUpdateStepsVal.FileCopy != nil && rRecipesValUpdateStepsVal.FileCopy != dclService.EmptyGuestPolicyRecipesUpdateStepsFileCopy {
				rRecipesValUpdateStepsValFileCopy := make(map[string]interface{})
				if rRecipesValUpdateStepsVal.FileCopy.ArtifactId != nil {
					rRecipesValUpdateStepsValFileCopy["artifactId"] = *rRecipesValUpdateStepsVal.FileCopy.ArtifactId
				}
				if rRecipesValUpdateStepsVal.FileCopy.Destination != nil {
					rRecipesValUpdateStepsValFileCopy["destination"] = *rRecipesValUpdateStepsVal.FileCopy.Destination
				}
				if rRecipesValUpdateStepsVal.FileCopy.Overwrite != nil {
					rRecipesValUpdateStepsValFileCopy["overwrite"] = *rRecipesValUpdateStepsVal.FileCopy.Overwrite
				}
				if rRecipesValUpdateStepsVal.FileCopy.Permissions != nil {
					rRecipesValUpdateStepsValFileCopy["permissions"] = *rRecipesValUpdateStepsVal.FileCopy.Permissions
				}
				rRecipesValUpdateStepsObject["fileCopy"] = rRecipesValUpdateStepsValFileCopy
			}
			if rRecipesValUpdateStepsVal.FileExec != nil && rRecipesValUpdateStepsVal.FileExec != dclService.EmptyGuestPolicyRecipesUpdateStepsFileExec {
				rRecipesValUpdateStepsValFileExec := make(map[string]interface{})
				var rRecipesValUpdateStepsValFileExecAllowedExitCodes []interface{}
				for _, rRecipesValUpdateStepsValFileExecAllowedExitCodesVal := range rRecipesValUpdateStepsVal.FileExec.AllowedExitCodes {
					rRecipesValUpdateStepsValFileExecAllowedExitCodes = append(rRecipesValUpdateStepsValFileExecAllowedExitCodes, rRecipesValUpdateStepsValFileExecAllowedExitCodesVal)
				}
				rRecipesValUpdateStepsValFileExec["allowedExitCodes"] = rRecipesValUpdateStepsValFileExecAllowedExitCodes
				var rRecipesValUpdateStepsValFileExecArgs []interface{}
				for _, rRecipesValUpdateStepsValFileExecArgsVal := range rRecipesValUpdateStepsVal.FileExec.Args {
					rRecipesValUpdateStepsValFileExecArgs = append(rRecipesValUpdateStepsValFileExecArgs, rRecipesValUpdateStepsValFileExecArgsVal)
				}
				rRecipesValUpdateStepsValFileExec["args"] = rRecipesValUpdateStepsValFileExecArgs
				if rRecipesValUpdateStepsVal.FileExec.ArtifactId != nil {
					rRecipesValUpdateStepsValFileExec["artifactId"] = *rRecipesValUpdateStepsVal.FileExec.ArtifactId
				}
				if rRecipesValUpdateStepsVal.FileExec.LocalPath != nil {
					rRecipesValUpdateStepsValFileExec["localPath"] = *rRecipesValUpdateStepsVal.FileExec.LocalPath
				}
				rRecipesValUpdateStepsObject["fileExec"] = rRecipesValUpdateStepsValFileExec
			}
			if rRecipesValUpdateStepsVal.MsiInstallation != nil && rRecipesValUpdateStepsVal.MsiInstallation != dclService.EmptyGuestPolicyRecipesUpdateStepsMsiInstallation {
				rRecipesValUpdateStepsValMsiInstallation := make(map[string]interface{})
				var rRecipesValUpdateStepsValMsiInstallationAllowedExitCodes []interface{}
				for _, rRecipesValUpdateStepsValMsiInstallationAllowedExitCodesVal := range rRecipesValUpdateStepsVal.MsiInstallation.AllowedExitCodes {
					rRecipesValUpdateStepsValMsiInstallationAllowedExitCodes = append(rRecipesValUpdateStepsValMsiInstallationAllowedExitCodes, rRecipesValUpdateStepsValMsiInstallationAllowedExitCodesVal)
				}
				rRecipesValUpdateStepsValMsiInstallation["allowedExitCodes"] = rRecipesValUpdateStepsValMsiInstallationAllowedExitCodes
				if rRecipesValUpdateStepsVal.MsiInstallation.ArtifactId != nil {
					rRecipesValUpdateStepsValMsiInstallation["artifactId"] = *rRecipesValUpdateStepsVal.MsiInstallation.ArtifactId
				}
				var rRecipesValUpdateStepsValMsiInstallationFlags []interface{}
				for _, rRecipesValUpdateStepsValMsiInstallationFlagsVal := range rRecipesValUpdateStepsVal.MsiInstallation.Flags {
					rRecipesValUpdateStepsValMsiInstallationFlags = append(rRecipesValUpdateStepsValMsiInstallationFlags, rRecipesValUpdateStepsValMsiInstallationFlagsVal)
				}
				rRecipesValUpdateStepsValMsiInstallation["flags"] = rRecipesValUpdateStepsValMsiInstallationFlags
				rRecipesValUpdateStepsObject["msiInstallation"] = rRecipesValUpdateStepsValMsiInstallation
			}
			if rRecipesValUpdateStepsVal.RpmInstallation != nil && rRecipesValUpdateStepsVal.RpmInstallation != dclService.EmptyGuestPolicyRecipesUpdateStepsRpmInstallation {
				rRecipesValUpdateStepsValRpmInstallation := make(map[string]interface{})
				if rRecipesValUpdateStepsVal.RpmInstallation.ArtifactId != nil {
					rRecipesValUpdateStepsValRpmInstallation["artifactId"] = *rRecipesValUpdateStepsVal.RpmInstallation.ArtifactId
				}
				rRecipesValUpdateStepsObject["rpmInstallation"] = rRecipesValUpdateStepsValRpmInstallation
			}
			if rRecipesValUpdateStepsVal.ScriptRun != nil && rRecipesValUpdateStepsVal.ScriptRun != dclService.EmptyGuestPolicyRecipesUpdateStepsScriptRun {
				rRecipesValUpdateStepsValScriptRun := make(map[string]interface{})
				var rRecipesValUpdateStepsValScriptRunAllowedExitCodes []interface{}
				for _, rRecipesValUpdateStepsValScriptRunAllowedExitCodesVal := range rRecipesValUpdateStepsVal.ScriptRun.AllowedExitCodes {
					rRecipesValUpdateStepsValScriptRunAllowedExitCodes = append(rRecipesValUpdateStepsValScriptRunAllowedExitCodes, rRecipesValUpdateStepsValScriptRunAllowedExitCodesVal)
				}
				rRecipesValUpdateStepsValScriptRun["allowedExitCodes"] = rRecipesValUpdateStepsValScriptRunAllowedExitCodes
				if rRecipesValUpdateStepsVal.ScriptRun.Interpreter != nil {
					rRecipesValUpdateStepsValScriptRun["interpreter"] = string(*rRecipesValUpdateStepsVal.ScriptRun.Interpreter)
				}
				if rRecipesValUpdateStepsVal.ScriptRun.Script != nil {
					rRecipesValUpdateStepsValScriptRun["script"] = *rRecipesValUpdateStepsVal.ScriptRun.Script
				}
				rRecipesValUpdateStepsObject["scriptRun"] = rRecipesValUpdateStepsValScriptRun
			}
			rRecipesValUpdateSteps = append(rRecipesValUpdateSteps, rRecipesValUpdateStepsObject)
		}
		rRecipesObject["updateSteps"] = rRecipesValUpdateSteps
		if rRecipesVal.Version != nil {
			rRecipesObject["version"] = *rRecipesVal.Version
		}
		rRecipes = append(rRecipes, rRecipesObject)
	}
	u.Object["recipes"] = rRecipes
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToGuestPolicy(u *unstructured.Resource) (*dclService.GuestPolicy, error) {
	r := &dclService.GuestPolicy{}
	if _, ok := u.Object["assignment"]; ok {
		if rAssignment, ok := u.Object["assignment"].(map[string]interface{}); ok {
			r.Assignment = &dclService.GuestPolicyAssignment{}
			if _, ok := rAssignment["groupLabels"]; ok {
				if s, ok := rAssignment["groupLabels"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rAssignmentGroupLabels dclService.GuestPolicyAssignmentGroupLabels
							if _, ok := objval["labels"]; ok {
								if rAssignmentGroupLabelsLabels, ok := objval["labels"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rAssignmentGroupLabelsLabels {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rAssignmentGroupLabels.Labels = m
								} else {
									return nil, fmt.Errorf("rAssignmentGroupLabels.Labels: expected map[string]interface{}")
								}
							}
							r.Assignment.GroupLabels = append(r.Assignment.GroupLabels, rAssignmentGroupLabels)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Assignment.GroupLabels: expected []interface{}")
				}
			}
			if _, ok := rAssignment["instanceNamePrefixes"]; ok {
				if s, ok := rAssignment["instanceNamePrefixes"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Assignment.InstanceNamePrefixes = append(r.Assignment.InstanceNamePrefixes, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Assignment.InstanceNamePrefixes: expected []interface{}")
				}
			}
			if _, ok := rAssignment["instances"]; ok {
				if s, ok := rAssignment["instances"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Assignment.Instances = append(r.Assignment.Instances, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Assignment.Instances: expected []interface{}")
				}
			}
			if _, ok := rAssignment["osTypes"]; ok {
				if s, ok := rAssignment["osTypes"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rAssignmentOSTypes dclService.GuestPolicyAssignmentOSTypes
							if _, ok := objval["osArchitecture"]; ok {
								if s, ok := objval["osArchitecture"].(string); ok {
									rAssignmentOSTypes.OSArchitecture = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAssignmentOSTypes.OSArchitecture: expected string")
								}
							}
							if _, ok := objval["osShortName"]; ok {
								if s, ok := objval["osShortName"].(string); ok {
									rAssignmentOSTypes.OSShortName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAssignmentOSTypes.OSShortName: expected string")
								}
							}
							if _, ok := objval["osVersion"]; ok {
								if s, ok := objval["osVersion"].(string); ok {
									rAssignmentOSTypes.OSVersion = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rAssignmentOSTypes.OSVersion: expected string")
								}
							}
							r.Assignment.OSTypes = append(r.Assignment.OSTypes, rAssignmentOSTypes)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Assignment.OSTypes: expected []interface{}")
				}
			}
			if _, ok := rAssignment["zones"]; ok {
				if s, ok := rAssignment["zones"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Assignment.Zones = append(r.Assignment.Zones, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Assignment.Zones: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Assignment: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["packageRepositories"]; ok {
		if s, ok := u.Object["packageRepositories"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rPackageRepositories dclService.GuestPolicyPackageRepositories
					if _, ok := objval["apt"]; ok {
						if rPackageRepositoriesApt, ok := objval["apt"].(map[string]interface{}); ok {
							rPackageRepositories.Apt = &dclService.GuestPolicyPackageRepositoriesApt{}
							if _, ok := rPackageRepositoriesApt["archiveType"]; ok {
								if s, ok := rPackageRepositoriesApt["archiveType"].(string); ok {
									rPackageRepositories.Apt.ArchiveType = dclService.GuestPolicyPackageRepositoriesAptArchiveTypeEnumRef(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Apt.ArchiveType: expected string")
								}
							}
							if _, ok := rPackageRepositoriesApt["components"]; ok {
								if s, ok := rPackageRepositoriesApt["components"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rPackageRepositories.Apt.Components = append(rPackageRepositories.Apt.Components, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Apt.Components: expected []interface{}")
								}
							}
							if _, ok := rPackageRepositoriesApt["distribution"]; ok {
								if s, ok := rPackageRepositoriesApt["distribution"].(string); ok {
									rPackageRepositories.Apt.Distribution = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Apt.Distribution: expected string")
								}
							}
							if _, ok := rPackageRepositoriesApt["gpgKey"]; ok {
								if s, ok := rPackageRepositoriesApt["gpgKey"].(string); ok {
									rPackageRepositories.Apt.GpgKey = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Apt.GpgKey: expected string")
								}
							}
							if _, ok := rPackageRepositoriesApt["uri"]; ok {
								if s, ok := rPackageRepositoriesApt["uri"].(string); ok {
									rPackageRepositories.Apt.Uri = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Apt.Uri: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rPackageRepositories.Apt: expected map[string]interface{}")
						}
					}
					if _, ok := objval["goo"]; ok {
						if rPackageRepositoriesGoo, ok := objval["goo"].(map[string]interface{}); ok {
							rPackageRepositories.Goo = &dclService.GuestPolicyPackageRepositoriesGoo{}
							if _, ok := rPackageRepositoriesGoo["name"]; ok {
								if s, ok := rPackageRepositoriesGoo["name"].(string); ok {
									rPackageRepositories.Goo.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Goo.Name: expected string")
								}
							}
							if _, ok := rPackageRepositoriesGoo["url"]; ok {
								if s, ok := rPackageRepositoriesGoo["url"].(string); ok {
									rPackageRepositories.Goo.Url = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Goo.Url: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rPackageRepositories.Goo: expected map[string]interface{}")
						}
					}
					if _, ok := objval["yum"]; ok {
						if rPackageRepositoriesYum, ok := objval["yum"].(map[string]interface{}); ok {
							rPackageRepositories.Yum = &dclService.GuestPolicyPackageRepositoriesYum{}
							if _, ok := rPackageRepositoriesYum["baseUrl"]; ok {
								if s, ok := rPackageRepositoriesYum["baseUrl"].(string); ok {
									rPackageRepositories.Yum.BaseUrl = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Yum.BaseUrl: expected string")
								}
							}
							if _, ok := rPackageRepositoriesYum["displayName"]; ok {
								if s, ok := rPackageRepositoriesYum["displayName"].(string); ok {
									rPackageRepositories.Yum.DisplayName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Yum.DisplayName: expected string")
								}
							}
							if _, ok := rPackageRepositoriesYum["gpgKeys"]; ok {
								if s, ok := rPackageRepositoriesYum["gpgKeys"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rPackageRepositories.Yum.GpgKeys = append(rPackageRepositories.Yum.GpgKeys, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Yum.GpgKeys: expected []interface{}")
								}
							}
							if _, ok := rPackageRepositoriesYum["id"]; ok {
								if s, ok := rPackageRepositoriesYum["id"].(string); ok {
									rPackageRepositories.Yum.Id = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Yum.Id: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rPackageRepositories.Yum: expected map[string]interface{}")
						}
					}
					if _, ok := objval["zypper"]; ok {
						if rPackageRepositoriesZypper, ok := objval["zypper"].(map[string]interface{}); ok {
							rPackageRepositories.Zypper = &dclService.GuestPolicyPackageRepositoriesZypper{}
							if _, ok := rPackageRepositoriesZypper["baseUrl"]; ok {
								if s, ok := rPackageRepositoriesZypper["baseUrl"].(string); ok {
									rPackageRepositories.Zypper.BaseUrl = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Zypper.BaseUrl: expected string")
								}
							}
							if _, ok := rPackageRepositoriesZypper["displayName"]; ok {
								if s, ok := rPackageRepositoriesZypper["displayName"].(string); ok {
									rPackageRepositories.Zypper.DisplayName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Zypper.DisplayName: expected string")
								}
							}
							if _, ok := rPackageRepositoriesZypper["gpgKeys"]; ok {
								if s, ok := rPackageRepositoriesZypper["gpgKeys"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rPackageRepositories.Zypper.GpgKeys = append(rPackageRepositories.Zypper.GpgKeys, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Zypper.GpgKeys: expected []interface{}")
								}
							}
							if _, ok := rPackageRepositoriesZypper["id"]; ok {
								if s, ok := rPackageRepositoriesZypper["id"].(string); ok {
									rPackageRepositories.Zypper.Id = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rPackageRepositories.Zypper.Id: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("rPackageRepositories.Zypper: expected map[string]interface{}")
						}
					}
					r.PackageRepositories = append(r.PackageRepositories, rPackageRepositories)
				}
			}
		} else {
			return nil, fmt.Errorf("r.PackageRepositories: expected []interface{}")
		}
	}
	if _, ok := u.Object["packages"]; ok {
		if s, ok := u.Object["packages"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rPackages dclService.GuestPolicyPackages
					if _, ok := objval["desiredState"]; ok {
						if s, ok := objval["desiredState"].(string); ok {
							rPackages.DesiredState = dclService.GuestPolicyPackagesDesiredStateEnumRef(s)
						} else {
							return nil, fmt.Errorf("rPackages.DesiredState: expected string")
						}
					}
					if _, ok := objval["manager"]; ok {
						if s, ok := objval["manager"].(string); ok {
							rPackages.Manager = dclService.GuestPolicyPackagesManagerEnumRef(s)
						} else {
							return nil, fmt.Errorf("rPackages.Manager: expected string")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rPackages.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rPackages.Name: expected string")
						}
					}
					r.Packages = append(r.Packages, rPackages)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Packages: expected []interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["recipes"]; ok {
		if s, ok := u.Object["recipes"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rRecipes dclService.GuestPolicyRecipes
					if _, ok := objval["artifacts"]; ok {
						if s, ok := objval["artifacts"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rRecipesArtifacts dclService.GuestPolicyRecipesArtifacts
									if _, ok := objval["allowInsecure"]; ok {
										if b, ok := objval["allowInsecure"].(bool); ok {
											rRecipesArtifacts.AllowInsecure = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rRecipesArtifacts.AllowInsecure: expected bool")
										}
									}
									if _, ok := objval["gcs"]; ok {
										if rRecipesArtifactsGcs, ok := objval["gcs"].(map[string]interface{}); ok {
											rRecipesArtifacts.Gcs = &dclService.GuestPolicyRecipesArtifactsGcs{}
											if _, ok := rRecipesArtifactsGcs["bucket"]; ok {
												if s, ok := rRecipesArtifactsGcs["bucket"].(string); ok {
													rRecipesArtifacts.Gcs.Bucket = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesArtifacts.Gcs.Bucket: expected string")
												}
											}
											if _, ok := rRecipesArtifactsGcs["generation"]; ok {
												if i, ok := rRecipesArtifactsGcs["generation"].(int64); ok {
													rRecipesArtifacts.Gcs.Generation = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rRecipesArtifacts.Gcs.Generation: expected int64")
												}
											}
											if _, ok := rRecipesArtifactsGcs["object"]; ok {
												if s, ok := rRecipesArtifactsGcs["object"].(string); ok {
													rRecipesArtifacts.Gcs.Object = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesArtifacts.Gcs.Object: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesArtifacts.Gcs: expected map[string]interface{}")
										}
									}
									if _, ok := objval["id"]; ok {
										if s, ok := objval["id"].(string); ok {
											rRecipesArtifacts.Id = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRecipesArtifacts.Id: expected string")
										}
									}
									if _, ok := objval["remote"]; ok {
										if rRecipesArtifactsRemote, ok := objval["remote"].(map[string]interface{}); ok {
											rRecipesArtifacts.Remote = &dclService.GuestPolicyRecipesArtifactsRemote{}
											if _, ok := rRecipesArtifactsRemote["checksum"]; ok {
												if s, ok := rRecipesArtifactsRemote["checksum"].(string); ok {
													rRecipesArtifacts.Remote.Checksum = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesArtifacts.Remote.Checksum: expected string")
												}
											}
											if _, ok := rRecipesArtifactsRemote["uri"]; ok {
												if s, ok := rRecipesArtifactsRemote["uri"].(string); ok {
													rRecipesArtifacts.Remote.Uri = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesArtifacts.Remote.Uri: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesArtifacts.Remote: expected map[string]interface{}")
										}
									}
									rRecipes.Artifacts = append(rRecipes.Artifacts, rRecipesArtifacts)
								}
							}
						} else {
							return nil, fmt.Errorf("rRecipes.Artifacts: expected []interface{}")
						}
					}
					if _, ok := objval["desiredState"]; ok {
						if s, ok := objval["desiredState"].(string); ok {
							rRecipes.DesiredState = dclService.GuestPolicyRecipesDesiredStateEnumRef(s)
						} else {
							return nil, fmt.Errorf("rRecipes.DesiredState: expected string")
						}
					}
					if _, ok := objval["installSteps"]; ok {
						if s, ok := objval["installSteps"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rRecipesInstallSteps dclService.GuestPolicyRecipesInstallSteps
									if _, ok := objval["archiveExtraction"]; ok {
										if rRecipesInstallStepsArchiveExtraction, ok := objval["archiveExtraction"].(map[string]interface{}); ok {
											rRecipesInstallSteps.ArchiveExtraction = &dclService.GuestPolicyRecipesInstallStepsArchiveExtraction{}
											if _, ok := rRecipesInstallStepsArchiveExtraction["artifactId"]; ok {
												if s, ok := rRecipesInstallStepsArchiveExtraction["artifactId"].(string); ok {
													rRecipesInstallSteps.ArchiveExtraction.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.ArchiveExtraction.ArtifactId: expected string")
												}
											}
											if _, ok := rRecipesInstallStepsArchiveExtraction["destination"]; ok {
												if s, ok := rRecipesInstallStepsArchiveExtraction["destination"].(string); ok {
													rRecipesInstallSteps.ArchiveExtraction.Destination = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.ArchiveExtraction.Destination: expected string")
												}
											}
											if _, ok := rRecipesInstallStepsArchiveExtraction["type"]; ok {
												if s, ok := rRecipesInstallStepsArchiveExtraction["type"].(string); ok {
													rRecipesInstallSteps.ArchiveExtraction.Type = dclService.GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnumRef(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.ArchiveExtraction.Type: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesInstallSteps.ArchiveExtraction: expected map[string]interface{}")
										}
									}
									if _, ok := objval["dpkgInstallation"]; ok {
										if rRecipesInstallStepsDpkgInstallation, ok := objval["dpkgInstallation"].(map[string]interface{}); ok {
											rRecipesInstallSteps.DpkgInstallation = &dclService.GuestPolicyRecipesInstallStepsDpkgInstallation{}
											if _, ok := rRecipesInstallStepsDpkgInstallation["artifactId"]; ok {
												if s, ok := rRecipesInstallStepsDpkgInstallation["artifactId"].(string); ok {
													rRecipesInstallSteps.DpkgInstallation.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.DpkgInstallation.ArtifactId: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesInstallSteps.DpkgInstallation: expected map[string]interface{}")
										}
									}
									if _, ok := objval["fileCopy"]; ok {
										if rRecipesInstallStepsFileCopy, ok := objval["fileCopy"].(map[string]interface{}); ok {
											rRecipesInstallSteps.FileCopy = &dclService.GuestPolicyRecipesInstallStepsFileCopy{}
											if _, ok := rRecipesInstallStepsFileCopy["artifactId"]; ok {
												if s, ok := rRecipesInstallStepsFileCopy["artifactId"].(string); ok {
													rRecipesInstallSteps.FileCopy.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.FileCopy.ArtifactId: expected string")
												}
											}
											if _, ok := rRecipesInstallStepsFileCopy["destination"]; ok {
												if s, ok := rRecipesInstallStepsFileCopy["destination"].(string); ok {
													rRecipesInstallSteps.FileCopy.Destination = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.FileCopy.Destination: expected string")
												}
											}
											if _, ok := rRecipesInstallStepsFileCopy["overwrite"]; ok {
												if b, ok := rRecipesInstallStepsFileCopy["overwrite"].(bool); ok {
													rRecipesInstallSteps.FileCopy.Overwrite = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.FileCopy.Overwrite: expected bool")
												}
											}
											if _, ok := rRecipesInstallStepsFileCopy["permissions"]; ok {
												if s, ok := rRecipesInstallStepsFileCopy["permissions"].(string); ok {
													rRecipesInstallSteps.FileCopy.Permissions = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.FileCopy.Permissions: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesInstallSteps.FileCopy: expected map[string]interface{}")
										}
									}
									if _, ok := objval["fileExec"]; ok {
										if rRecipesInstallStepsFileExec, ok := objval["fileExec"].(map[string]interface{}); ok {
											rRecipesInstallSteps.FileExec = &dclService.GuestPolicyRecipesInstallStepsFileExec{}
											if _, ok := rRecipesInstallStepsFileExec["allowedExitCodes"]; ok {
												if s, ok := rRecipesInstallStepsFileExec["allowedExitCodes"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rRecipesInstallSteps.FileExec.AllowedExitCodes = append(rRecipesInstallSteps.FileExec.AllowedExitCodes, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.FileExec.AllowedExitCodes: expected []interface{}")
												}
											}
											if _, ok := rRecipesInstallStepsFileExec["args"]; ok {
												if s, ok := rRecipesInstallStepsFileExec["args"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															rRecipesInstallSteps.FileExec.Args = append(rRecipesInstallSteps.FileExec.Args, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.FileExec.Args: expected []interface{}")
												}
											}
											if _, ok := rRecipesInstallStepsFileExec["artifactId"]; ok {
												if s, ok := rRecipesInstallStepsFileExec["artifactId"].(string); ok {
													rRecipesInstallSteps.FileExec.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.FileExec.ArtifactId: expected string")
												}
											}
											if _, ok := rRecipesInstallStepsFileExec["localPath"]; ok {
												if s, ok := rRecipesInstallStepsFileExec["localPath"].(string); ok {
													rRecipesInstallSteps.FileExec.LocalPath = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.FileExec.LocalPath: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesInstallSteps.FileExec: expected map[string]interface{}")
										}
									}
									if _, ok := objval["msiInstallation"]; ok {
										if rRecipesInstallStepsMsiInstallation, ok := objval["msiInstallation"].(map[string]interface{}); ok {
											rRecipesInstallSteps.MsiInstallation = &dclService.GuestPolicyRecipesInstallStepsMsiInstallation{}
											if _, ok := rRecipesInstallStepsMsiInstallation["allowedExitCodes"]; ok {
												if s, ok := rRecipesInstallStepsMsiInstallation["allowedExitCodes"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rRecipesInstallSteps.MsiInstallation.AllowedExitCodes = append(rRecipesInstallSteps.MsiInstallation.AllowedExitCodes, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.MsiInstallation.AllowedExitCodes: expected []interface{}")
												}
											}
											if _, ok := rRecipesInstallStepsMsiInstallation["artifactId"]; ok {
												if s, ok := rRecipesInstallStepsMsiInstallation["artifactId"].(string); ok {
													rRecipesInstallSteps.MsiInstallation.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.MsiInstallation.ArtifactId: expected string")
												}
											}
											if _, ok := rRecipesInstallStepsMsiInstallation["flags"]; ok {
												if s, ok := rRecipesInstallStepsMsiInstallation["flags"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															rRecipesInstallSteps.MsiInstallation.Flags = append(rRecipesInstallSteps.MsiInstallation.Flags, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.MsiInstallation.Flags: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesInstallSteps.MsiInstallation: expected map[string]interface{}")
										}
									}
									if _, ok := objval["rpmInstallation"]; ok {
										if rRecipesInstallStepsRpmInstallation, ok := objval["rpmInstallation"].(map[string]interface{}); ok {
											rRecipesInstallSteps.RpmInstallation = &dclService.GuestPolicyRecipesInstallStepsRpmInstallation{}
											if _, ok := rRecipesInstallStepsRpmInstallation["artifactId"]; ok {
												if s, ok := rRecipesInstallStepsRpmInstallation["artifactId"].(string); ok {
													rRecipesInstallSteps.RpmInstallation.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.RpmInstallation.ArtifactId: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesInstallSteps.RpmInstallation: expected map[string]interface{}")
										}
									}
									if _, ok := objval["scriptRun"]; ok {
										if rRecipesInstallStepsScriptRun, ok := objval["scriptRun"].(map[string]interface{}); ok {
											rRecipesInstallSteps.ScriptRun = &dclService.GuestPolicyRecipesInstallStepsScriptRun{}
											if _, ok := rRecipesInstallStepsScriptRun["allowedExitCodes"]; ok {
												if s, ok := rRecipesInstallStepsScriptRun["allowedExitCodes"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rRecipesInstallSteps.ScriptRun.AllowedExitCodes = append(rRecipesInstallSteps.ScriptRun.AllowedExitCodes, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.ScriptRun.AllowedExitCodes: expected []interface{}")
												}
											}
											if _, ok := rRecipesInstallStepsScriptRun["interpreter"]; ok {
												if s, ok := rRecipesInstallStepsScriptRun["interpreter"].(string); ok {
													rRecipesInstallSteps.ScriptRun.Interpreter = dclService.GuestPolicyRecipesInstallStepsScriptRunInterpreterEnumRef(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.ScriptRun.Interpreter: expected string")
												}
											}
											if _, ok := rRecipesInstallStepsScriptRun["script"]; ok {
												if s, ok := rRecipesInstallStepsScriptRun["script"].(string); ok {
													rRecipesInstallSteps.ScriptRun.Script = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesInstallSteps.ScriptRun.Script: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesInstallSteps.ScriptRun: expected map[string]interface{}")
										}
									}
									rRecipes.InstallSteps = append(rRecipes.InstallSteps, rRecipesInstallSteps)
								}
							}
						} else {
							return nil, fmt.Errorf("rRecipes.InstallSteps: expected []interface{}")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rRecipes.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rRecipes.Name: expected string")
						}
					}
					if _, ok := objval["updateSteps"]; ok {
						if s, ok := objval["updateSteps"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rRecipesUpdateSteps dclService.GuestPolicyRecipesUpdateSteps
									if _, ok := objval["archiveExtraction"]; ok {
										if rRecipesUpdateStepsArchiveExtraction, ok := objval["archiveExtraction"].(map[string]interface{}); ok {
											rRecipesUpdateSteps.ArchiveExtraction = &dclService.GuestPolicyRecipesUpdateStepsArchiveExtraction{}
											if _, ok := rRecipesUpdateStepsArchiveExtraction["artifactId"]; ok {
												if s, ok := rRecipesUpdateStepsArchiveExtraction["artifactId"].(string); ok {
													rRecipesUpdateSteps.ArchiveExtraction.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.ArchiveExtraction.ArtifactId: expected string")
												}
											}
											if _, ok := rRecipesUpdateStepsArchiveExtraction["destination"]; ok {
												if s, ok := rRecipesUpdateStepsArchiveExtraction["destination"].(string); ok {
													rRecipesUpdateSteps.ArchiveExtraction.Destination = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.ArchiveExtraction.Destination: expected string")
												}
											}
											if _, ok := rRecipesUpdateStepsArchiveExtraction["type"]; ok {
												if s, ok := rRecipesUpdateStepsArchiveExtraction["type"].(string); ok {
													rRecipesUpdateSteps.ArchiveExtraction.Type = dclService.GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnumRef(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.ArchiveExtraction.Type: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesUpdateSteps.ArchiveExtraction: expected map[string]interface{}")
										}
									}
									if _, ok := objval["dpkgInstallation"]; ok {
										if rRecipesUpdateStepsDpkgInstallation, ok := objval["dpkgInstallation"].(map[string]interface{}); ok {
											rRecipesUpdateSteps.DpkgInstallation = &dclService.GuestPolicyRecipesUpdateStepsDpkgInstallation{}
											if _, ok := rRecipesUpdateStepsDpkgInstallation["artifactId"]; ok {
												if s, ok := rRecipesUpdateStepsDpkgInstallation["artifactId"].(string); ok {
													rRecipesUpdateSteps.DpkgInstallation.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.DpkgInstallation.ArtifactId: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesUpdateSteps.DpkgInstallation: expected map[string]interface{}")
										}
									}
									if _, ok := objval["fileCopy"]; ok {
										if rRecipesUpdateStepsFileCopy, ok := objval["fileCopy"].(map[string]interface{}); ok {
											rRecipesUpdateSteps.FileCopy = &dclService.GuestPolicyRecipesUpdateStepsFileCopy{}
											if _, ok := rRecipesUpdateStepsFileCopy["artifactId"]; ok {
												if s, ok := rRecipesUpdateStepsFileCopy["artifactId"].(string); ok {
													rRecipesUpdateSteps.FileCopy.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.FileCopy.ArtifactId: expected string")
												}
											}
											if _, ok := rRecipesUpdateStepsFileCopy["destination"]; ok {
												if s, ok := rRecipesUpdateStepsFileCopy["destination"].(string); ok {
													rRecipesUpdateSteps.FileCopy.Destination = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.FileCopy.Destination: expected string")
												}
											}
											if _, ok := rRecipesUpdateStepsFileCopy["overwrite"]; ok {
												if b, ok := rRecipesUpdateStepsFileCopy["overwrite"].(bool); ok {
													rRecipesUpdateSteps.FileCopy.Overwrite = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.FileCopy.Overwrite: expected bool")
												}
											}
											if _, ok := rRecipesUpdateStepsFileCopy["permissions"]; ok {
												if s, ok := rRecipesUpdateStepsFileCopy["permissions"].(string); ok {
													rRecipesUpdateSteps.FileCopy.Permissions = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.FileCopy.Permissions: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesUpdateSteps.FileCopy: expected map[string]interface{}")
										}
									}
									if _, ok := objval["fileExec"]; ok {
										if rRecipesUpdateStepsFileExec, ok := objval["fileExec"].(map[string]interface{}); ok {
											rRecipesUpdateSteps.FileExec = &dclService.GuestPolicyRecipesUpdateStepsFileExec{}
											if _, ok := rRecipesUpdateStepsFileExec["allowedExitCodes"]; ok {
												if s, ok := rRecipesUpdateStepsFileExec["allowedExitCodes"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rRecipesUpdateSteps.FileExec.AllowedExitCodes = append(rRecipesUpdateSteps.FileExec.AllowedExitCodes, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.FileExec.AllowedExitCodes: expected []interface{}")
												}
											}
											if _, ok := rRecipesUpdateStepsFileExec["args"]; ok {
												if s, ok := rRecipesUpdateStepsFileExec["args"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															rRecipesUpdateSteps.FileExec.Args = append(rRecipesUpdateSteps.FileExec.Args, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.FileExec.Args: expected []interface{}")
												}
											}
											if _, ok := rRecipesUpdateStepsFileExec["artifactId"]; ok {
												if s, ok := rRecipesUpdateStepsFileExec["artifactId"].(string); ok {
													rRecipesUpdateSteps.FileExec.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.FileExec.ArtifactId: expected string")
												}
											}
											if _, ok := rRecipesUpdateStepsFileExec["localPath"]; ok {
												if s, ok := rRecipesUpdateStepsFileExec["localPath"].(string); ok {
													rRecipesUpdateSteps.FileExec.LocalPath = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.FileExec.LocalPath: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesUpdateSteps.FileExec: expected map[string]interface{}")
										}
									}
									if _, ok := objval["msiInstallation"]; ok {
										if rRecipesUpdateStepsMsiInstallation, ok := objval["msiInstallation"].(map[string]interface{}); ok {
											rRecipesUpdateSteps.MsiInstallation = &dclService.GuestPolicyRecipesUpdateStepsMsiInstallation{}
											if _, ok := rRecipesUpdateStepsMsiInstallation["allowedExitCodes"]; ok {
												if s, ok := rRecipesUpdateStepsMsiInstallation["allowedExitCodes"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rRecipesUpdateSteps.MsiInstallation.AllowedExitCodes = append(rRecipesUpdateSteps.MsiInstallation.AllowedExitCodes, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.MsiInstallation.AllowedExitCodes: expected []interface{}")
												}
											}
											if _, ok := rRecipesUpdateStepsMsiInstallation["artifactId"]; ok {
												if s, ok := rRecipesUpdateStepsMsiInstallation["artifactId"].(string); ok {
													rRecipesUpdateSteps.MsiInstallation.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.MsiInstallation.ArtifactId: expected string")
												}
											}
											if _, ok := rRecipesUpdateStepsMsiInstallation["flags"]; ok {
												if s, ok := rRecipesUpdateStepsMsiInstallation["flags"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															rRecipesUpdateSteps.MsiInstallation.Flags = append(rRecipesUpdateSteps.MsiInstallation.Flags, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.MsiInstallation.Flags: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesUpdateSteps.MsiInstallation: expected map[string]interface{}")
										}
									}
									if _, ok := objval["rpmInstallation"]; ok {
										if rRecipesUpdateStepsRpmInstallation, ok := objval["rpmInstallation"].(map[string]interface{}); ok {
											rRecipesUpdateSteps.RpmInstallation = &dclService.GuestPolicyRecipesUpdateStepsRpmInstallation{}
											if _, ok := rRecipesUpdateStepsRpmInstallation["artifactId"]; ok {
												if s, ok := rRecipesUpdateStepsRpmInstallation["artifactId"].(string); ok {
													rRecipesUpdateSteps.RpmInstallation.ArtifactId = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.RpmInstallation.ArtifactId: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesUpdateSteps.RpmInstallation: expected map[string]interface{}")
										}
									}
									if _, ok := objval["scriptRun"]; ok {
										if rRecipesUpdateStepsScriptRun, ok := objval["scriptRun"].(map[string]interface{}); ok {
											rRecipesUpdateSteps.ScriptRun = &dclService.GuestPolicyRecipesUpdateStepsScriptRun{}
											if _, ok := rRecipesUpdateStepsScriptRun["allowedExitCodes"]; ok {
												if s, ok := rRecipesUpdateStepsScriptRun["allowedExitCodes"].([]interface{}); ok {
													for _, ss := range s {
														if intval, ok := ss.(int64); ok {
															rRecipesUpdateSteps.ScriptRun.AllowedExitCodes = append(rRecipesUpdateSteps.ScriptRun.AllowedExitCodes, intval)
														}
													}
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.ScriptRun.AllowedExitCodes: expected []interface{}")
												}
											}
											if _, ok := rRecipesUpdateStepsScriptRun["interpreter"]; ok {
												if s, ok := rRecipesUpdateStepsScriptRun["interpreter"].(string); ok {
													rRecipesUpdateSteps.ScriptRun.Interpreter = dclService.GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnumRef(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.ScriptRun.Interpreter: expected string")
												}
											}
											if _, ok := rRecipesUpdateStepsScriptRun["script"]; ok {
												if s, ok := rRecipesUpdateStepsScriptRun["script"].(string); ok {
													rRecipesUpdateSteps.ScriptRun.Script = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRecipesUpdateSteps.ScriptRun.Script: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRecipesUpdateSteps.ScriptRun: expected map[string]interface{}")
										}
									}
									rRecipes.UpdateSteps = append(rRecipes.UpdateSteps, rRecipesUpdateSteps)
								}
							}
						} else {
							return nil, fmt.Errorf("rRecipes.UpdateSteps: expected []interface{}")
						}
					}
					if _, ok := objval["version"]; ok {
						if s, ok := objval["version"].(string); ok {
							rRecipes.Version = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rRecipes.Version: expected string")
						}
					}
					r.Recipes = append(r.Recipes, rRecipes)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Recipes: expected []interface{}")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetGuestPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGuestPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetGuestPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return GuestPolicyToUnstructured(r), nil
}

func ListGuestPolicy(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListGuestPolicy(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, GuestPolicyToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyGuestPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGuestPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGuestPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyGuestPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return GuestPolicyToUnstructured(r), nil
}

func GuestPolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGuestPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGuestPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyGuestPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteGuestPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGuestPolicy(u)
	if err != nil {
		return err
	}
	return c.DeleteGuestPolicy(ctx, r)
}

func GuestPolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToGuestPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *GuestPolicy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"osconfig",
		"GuestPolicy",
		"alpha",
	}
}

func (r *GuestPolicy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GuestPolicy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GuestPolicy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *GuestPolicy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GuestPolicy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GuestPolicy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GuestPolicy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetGuestPolicy(ctx, config, resource)
}

func (r *GuestPolicy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyGuestPolicy(ctx, config, resource, opts...)
}

func (r *GuestPolicy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return GuestPolicyHasDiff(ctx, config, resource, opts...)
}

func (r *GuestPolicy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteGuestPolicy(ctx, config, resource)
}

func (r *GuestPolicy) ID(resource *unstructured.Resource) (string, error) {
	return GuestPolicyID(resource)
}

func init() {
	unstructured.Register(&GuestPolicy{})
}
