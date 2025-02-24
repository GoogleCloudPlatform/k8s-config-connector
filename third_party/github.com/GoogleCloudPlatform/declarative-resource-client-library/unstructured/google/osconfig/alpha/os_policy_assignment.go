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

type OSPolicyAssignment struct{}

func OSPolicyAssignmentToUnstructured(r *dclService.OSPolicyAssignment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "osconfig",
			Version: "alpha",
			Type:    "OSPolicyAssignment",
		},
		Object: make(map[string]interface{}),
	}
	if r.Baseline != nil {
		u.Object["baseline"] = *r.Baseline
	}
	if r.Deleted != nil {
		u.Object["deleted"] = *r.Deleted
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.InstanceFilter != nil && r.InstanceFilter != dclService.EmptyOSPolicyAssignmentInstanceFilter {
		rInstanceFilter := make(map[string]interface{})
		if r.InstanceFilter.All != nil {
			rInstanceFilter["all"] = *r.InstanceFilter.All
		}
		var rInstanceFilterExclusionLabels []interface{}
		for _, rInstanceFilterExclusionLabelsVal := range r.InstanceFilter.ExclusionLabels {
			rInstanceFilterExclusionLabelsObject := make(map[string]interface{})
			if rInstanceFilterExclusionLabelsVal.Labels != nil {
				rInstanceFilterExclusionLabelsValLabels := make(map[string]interface{})
				for k, v := range rInstanceFilterExclusionLabelsVal.Labels {
					rInstanceFilterExclusionLabelsValLabels[k] = v
				}
				rInstanceFilterExclusionLabelsObject["labels"] = rInstanceFilterExclusionLabelsValLabels
			}
			rInstanceFilterExclusionLabels = append(rInstanceFilterExclusionLabels, rInstanceFilterExclusionLabelsObject)
		}
		rInstanceFilter["exclusionLabels"] = rInstanceFilterExclusionLabels
		var rInstanceFilterInclusionLabels []interface{}
		for _, rInstanceFilterInclusionLabelsVal := range r.InstanceFilter.InclusionLabels {
			rInstanceFilterInclusionLabelsObject := make(map[string]interface{})
			if rInstanceFilterInclusionLabelsVal.Labels != nil {
				rInstanceFilterInclusionLabelsValLabels := make(map[string]interface{})
				for k, v := range rInstanceFilterInclusionLabelsVal.Labels {
					rInstanceFilterInclusionLabelsValLabels[k] = v
				}
				rInstanceFilterInclusionLabelsObject["labels"] = rInstanceFilterInclusionLabelsValLabels
			}
			rInstanceFilterInclusionLabels = append(rInstanceFilterInclusionLabels, rInstanceFilterInclusionLabelsObject)
		}
		rInstanceFilter["inclusionLabels"] = rInstanceFilterInclusionLabels
		var rInstanceFilterInventories []interface{}
		for _, rInstanceFilterInventoriesVal := range r.InstanceFilter.Inventories {
			rInstanceFilterInventoriesObject := make(map[string]interface{})
			if rInstanceFilterInventoriesVal.OSShortName != nil {
				rInstanceFilterInventoriesObject["osShortName"] = *rInstanceFilterInventoriesVal.OSShortName
			}
			if rInstanceFilterInventoriesVal.OSVersion != nil {
				rInstanceFilterInventoriesObject["osVersion"] = *rInstanceFilterInventoriesVal.OSVersion
			}
			rInstanceFilterInventories = append(rInstanceFilterInventories, rInstanceFilterInventoriesObject)
		}
		rInstanceFilter["inventories"] = rInstanceFilterInventories
		u.Object["instanceFilter"] = rInstanceFilter
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	var rOSPolicies []interface{}
	for _, rOSPoliciesVal := range r.OSPolicies {
		rOSPoliciesObject := make(map[string]interface{})
		if rOSPoliciesVal.AllowNoResourceGroupMatch != nil {
			rOSPoliciesObject["allowNoResourceGroupMatch"] = *rOSPoliciesVal.AllowNoResourceGroupMatch
		}
		if rOSPoliciesVal.Description != nil {
			rOSPoliciesObject["description"] = *rOSPoliciesVal.Description
		}
		if rOSPoliciesVal.Id != nil {
			rOSPoliciesObject["id"] = *rOSPoliciesVal.Id
		}
		if rOSPoliciesVal.Mode != nil {
			rOSPoliciesObject["mode"] = string(*rOSPoliciesVal.Mode)
		}
		var rOSPoliciesValResourceGroups []interface{}
		for _, rOSPoliciesValResourceGroupsVal := range rOSPoliciesVal.ResourceGroups {
			rOSPoliciesValResourceGroupsObject := make(map[string]interface{})
			var rOSPoliciesValResourceGroupsValInventoryFilters []interface{}
			for _, rOSPoliciesValResourceGroupsValInventoryFiltersVal := range rOSPoliciesValResourceGroupsVal.InventoryFilters {
				rOSPoliciesValResourceGroupsValInventoryFiltersObject := make(map[string]interface{})
				if rOSPoliciesValResourceGroupsValInventoryFiltersVal.OSShortName != nil {
					rOSPoliciesValResourceGroupsValInventoryFiltersObject["osShortName"] = *rOSPoliciesValResourceGroupsValInventoryFiltersVal.OSShortName
				}
				if rOSPoliciesValResourceGroupsValInventoryFiltersVal.OSVersion != nil {
					rOSPoliciesValResourceGroupsValInventoryFiltersObject["osVersion"] = *rOSPoliciesValResourceGroupsValInventoryFiltersVal.OSVersion
				}
				rOSPoliciesValResourceGroupsValInventoryFilters = append(rOSPoliciesValResourceGroupsValInventoryFilters, rOSPoliciesValResourceGroupsValInventoryFiltersObject)
			}
			rOSPoliciesValResourceGroupsObject["inventoryFilters"] = rOSPoliciesValResourceGroupsValInventoryFilters
			var rOSPoliciesValResourceGroupsValResources []interface{}
			for _, rOSPoliciesValResourceGroupsValResourcesVal := range rOSPoliciesValResourceGroupsVal.Resources {
				rOSPoliciesValResourceGroupsValResourcesObject := make(map[string]interface{})
				if rOSPoliciesValResourceGroupsValResourcesVal.Exec != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec {
					rOSPoliciesValResourceGroupsValResourcesValExec := make(map[string]interface{})
					if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce {
						rOSPoliciesValResourceGroupsValResourcesValExecEnforce := make(map[string]interface{})
						var rOSPoliciesValResourceGroupsValResourcesValExecEnforceArgs []interface{}
						for _, rOSPoliciesValResourceGroupsValResourcesValExecEnforceArgsVal := range rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.Args {
							rOSPoliciesValResourceGroupsValResourcesValExecEnforceArgs = append(rOSPoliciesValResourceGroupsValResourcesValExecEnforceArgs, rOSPoliciesValResourceGroupsValResourcesValExecEnforceArgsVal)
						}
						rOSPoliciesValResourceGroupsValResourcesValExecEnforce["args"] = rOSPoliciesValResourceGroupsValResourcesValExecEnforceArgs
						if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile {
							rOSPoliciesValResourceGroupsValResourcesValExecEnforceFile := make(map[string]interface{})
							if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.AllowInsecure != nil {
								rOSPoliciesValResourceGroupsValResourcesValExecEnforceFile["allowInsecure"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.AllowInsecure
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Gcs != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Gcs != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs {
								rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileGcs := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Gcs.Bucket != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileGcs["bucket"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Gcs.Bucket
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Gcs.Generation != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileGcs["generation"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Gcs.Generation
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Gcs.Object != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileGcs["object"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Gcs.Object
								}
								rOSPoliciesValResourceGroupsValResourcesValExecEnforceFile["gcs"] = rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileGcs
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.LocalPath != nil {
								rOSPoliciesValResourceGroupsValResourcesValExecEnforceFile["localPath"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.LocalPath
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Remote != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Remote != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote {
								rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileRemote := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Remote.Sha256Checksum != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileRemote["sha256Checksum"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Remote.Sha256Checksum
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Remote.Uri != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileRemote["uri"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.File.Remote.Uri
								}
								rOSPoliciesValResourceGroupsValResourcesValExecEnforceFile["remote"] = rOSPoliciesValResourceGroupsValResourcesValExecEnforceFileRemote
							}
							rOSPoliciesValResourceGroupsValResourcesValExecEnforce["file"] = rOSPoliciesValResourceGroupsValResourcesValExecEnforceFile
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.Interpreter != nil {
							rOSPoliciesValResourceGroupsValResourcesValExecEnforce["interpreter"] = string(*rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.Interpreter)
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.OutputFilePath != nil {
							rOSPoliciesValResourceGroupsValResourcesValExecEnforce["outputFilePath"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.OutputFilePath
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.Script != nil {
							rOSPoliciesValResourceGroupsValResourcesValExecEnforce["script"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Enforce.Script
						}
						rOSPoliciesValResourceGroupsValResourcesValExec["enforce"] = rOSPoliciesValResourceGroupsValResourcesValExecEnforce
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate {
						rOSPoliciesValResourceGroupsValResourcesValExecValidate := make(map[string]interface{})
						var rOSPoliciesValResourceGroupsValResourcesValExecValidateArgs []interface{}
						for _, rOSPoliciesValResourceGroupsValResourcesValExecValidateArgsVal := range rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.Args {
							rOSPoliciesValResourceGroupsValResourcesValExecValidateArgs = append(rOSPoliciesValResourceGroupsValResourcesValExecValidateArgs, rOSPoliciesValResourceGroupsValResourcesValExecValidateArgsVal)
						}
						rOSPoliciesValResourceGroupsValResourcesValExecValidate["args"] = rOSPoliciesValResourceGroupsValResourcesValExecValidateArgs
						if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile {
							rOSPoliciesValResourceGroupsValResourcesValExecValidateFile := make(map[string]interface{})
							if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.AllowInsecure != nil {
								rOSPoliciesValResourceGroupsValResourcesValExecValidateFile["allowInsecure"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.AllowInsecure
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Gcs != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Gcs != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs {
								rOSPoliciesValResourceGroupsValResourcesValExecValidateFileGcs := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Gcs.Bucket != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecValidateFileGcs["bucket"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Gcs.Bucket
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Gcs.Generation != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecValidateFileGcs["generation"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Gcs.Generation
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Gcs.Object != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecValidateFileGcs["object"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Gcs.Object
								}
								rOSPoliciesValResourceGroupsValResourcesValExecValidateFile["gcs"] = rOSPoliciesValResourceGroupsValResourcesValExecValidateFileGcs
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.LocalPath != nil {
								rOSPoliciesValResourceGroupsValResourcesValExecValidateFile["localPath"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.LocalPath
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Remote != nil && rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Remote != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote {
								rOSPoliciesValResourceGroupsValResourcesValExecValidateFileRemote := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Remote.Sha256Checksum != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecValidateFileRemote["sha256Checksum"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Remote.Sha256Checksum
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Remote.Uri != nil {
									rOSPoliciesValResourceGroupsValResourcesValExecValidateFileRemote["uri"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.File.Remote.Uri
								}
								rOSPoliciesValResourceGroupsValResourcesValExecValidateFile["remote"] = rOSPoliciesValResourceGroupsValResourcesValExecValidateFileRemote
							}
							rOSPoliciesValResourceGroupsValResourcesValExecValidate["file"] = rOSPoliciesValResourceGroupsValResourcesValExecValidateFile
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.Interpreter != nil {
							rOSPoliciesValResourceGroupsValResourcesValExecValidate["interpreter"] = string(*rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.Interpreter)
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.OutputFilePath != nil {
							rOSPoliciesValResourceGroupsValResourcesValExecValidate["outputFilePath"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.OutputFilePath
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.Script != nil {
							rOSPoliciesValResourceGroupsValResourcesValExecValidate["script"] = *rOSPoliciesValResourceGroupsValResourcesVal.Exec.Validate.Script
						}
						rOSPoliciesValResourceGroupsValResourcesValExec["validate"] = rOSPoliciesValResourceGroupsValResourcesValExecValidate
					}
					rOSPoliciesValResourceGroupsValResourcesObject["exec"] = rOSPoliciesValResourceGroupsValResourcesValExec
				}
				if rOSPoliciesValResourceGroupsValResourcesVal.File != nil && rOSPoliciesValResourceGroupsValResourcesVal.File != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile {
					rOSPoliciesValResourceGroupsValResourcesValFile := make(map[string]interface{})
					if rOSPoliciesValResourceGroupsValResourcesVal.File.Content != nil {
						rOSPoliciesValResourceGroupsValResourcesValFile["content"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.Content
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.File.File != nil && rOSPoliciesValResourceGroupsValResourcesVal.File.File != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile {
						rOSPoliciesValResourceGroupsValResourcesValFileFile := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.File.File.AllowInsecure != nil {
							rOSPoliciesValResourceGroupsValResourcesValFileFile["allowInsecure"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.File.AllowInsecure
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.File.File.Gcs != nil && rOSPoliciesValResourceGroupsValResourcesVal.File.File.Gcs != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs {
							rOSPoliciesValResourceGroupsValResourcesValFileFileGcs := make(map[string]interface{})
							if rOSPoliciesValResourceGroupsValResourcesVal.File.File.Gcs.Bucket != nil {
								rOSPoliciesValResourceGroupsValResourcesValFileFileGcs["bucket"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.File.Gcs.Bucket
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.File.File.Gcs.Generation != nil {
								rOSPoliciesValResourceGroupsValResourcesValFileFileGcs["generation"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.File.Gcs.Generation
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.File.File.Gcs.Object != nil {
								rOSPoliciesValResourceGroupsValResourcesValFileFileGcs["object"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.File.Gcs.Object
							}
							rOSPoliciesValResourceGroupsValResourcesValFileFile["gcs"] = rOSPoliciesValResourceGroupsValResourcesValFileFileGcs
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.File.File.LocalPath != nil {
							rOSPoliciesValResourceGroupsValResourcesValFileFile["localPath"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.File.LocalPath
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.File.File.Remote != nil && rOSPoliciesValResourceGroupsValResourcesVal.File.File.Remote != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote {
							rOSPoliciesValResourceGroupsValResourcesValFileFileRemote := make(map[string]interface{})
							if rOSPoliciesValResourceGroupsValResourcesVal.File.File.Remote.Sha256Checksum != nil {
								rOSPoliciesValResourceGroupsValResourcesValFileFileRemote["sha256Checksum"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.File.Remote.Sha256Checksum
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.File.File.Remote.Uri != nil {
								rOSPoliciesValResourceGroupsValResourcesValFileFileRemote["uri"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.File.Remote.Uri
							}
							rOSPoliciesValResourceGroupsValResourcesValFileFile["remote"] = rOSPoliciesValResourceGroupsValResourcesValFileFileRemote
						}
						rOSPoliciesValResourceGroupsValResourcesValFile["file"] = rOSPoliciesValResourceGroupsValResourcesValFileFile
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.File.Path != nil {
						rOSPoliciesValResourceGroupsValResourcesValFile["path"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.Path
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.File.Permissions != nil {
						rOSPoliciesValResourceGroupsValResourcesValFile["permissions"] = *rOSPoliciesValResourceGroupsValResourcesVal.File.Permissions
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.File.State != nil {
						rOSPoliciesValResourceGroupsValResourcesValFile["state"] = string(*rOSPoliciesValResourceGroupsValResourcesVal.File.State)
					}
					rOSPoliciesValResourceGroupsValResourcesObject["file"] = rOSPoliciesValResourceGroupsValResourcesValFile
				}
				if rOSPoliciesValResourceGroupsValResourcesVal.Id != nil {
					rOSPoliciesValResourceGroupsValResourcesObject["id"] = *rOSPoliciesValResourceGroupsValResourcesVal.Id
				}
				if rOSPoliciesValResourceGroupsValResourcesVal.Pkg != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg {
					rOSPoliciesValResourceGroupsValResourcesValPkg := make(map[string]interface{})
					if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Apt != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Apt != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt {
						rOSPoliciesValResourceGroupsValResourcesValPkgApt := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Apt.Name != nil {
							rOSPoliciesValResourceGroupsValResourcesValPkgApt["name"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Apt.Name
						}
						rOSPoliciesValResourceGroupsValResourcesValPkg["apt"] = rOSPoliciesValResourceGroupsValResourcesValPkgApt
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb {
						rOSPoliciesValResourceGroupsValResourcesValPkgDeb := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.PullDeps != nil {
							rOSPoliciesValResourceGroupsValResourcesValPkgDeb["pullDeps"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.PullDeps
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource {
							rOSPoliciesValResourceGroupsValResourcesValPkgDebSource := make(map[string]interface{})
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.AllowInsecure != nil {
								rOSPoliciesValResourceGroupsValResourcesValPkgDebSource["allowInsecure"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.AllowInsecure
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Gcs != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Gcs != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs {
								rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceGcs := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Gcs.Bucket != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceGcs["bucket"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Gcs.Bucket
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Gcs.Generation != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceGcs["generation"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Gcs.Generation
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Gcs.Object != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceGcs["object"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Gcs.Object
								}
								rOSPoliciesValResourceGroupsValResourcesValPkgDebSource["gcs"] = rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceGcs
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.LocalPath != nil {
								rOSPoliciesValResourceGroupsValResourcesValPkgDebSource["localPath"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.LocalPath
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Remote != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Remote != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote {
								rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceRemote := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Remote.Sha256Checksum != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceRemote["sha256Checksum"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Remote.Sha256Checksum
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Remote.Uri != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceRemote["uri"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Deb.Source.Remote.Uri
								}
								rOSPoliciesValResourceGroupsValResourcesValPkgDebSource["remote"] = rOSPoliciesValResourceGroupsValResourcesValPkgDebSourceRemote
							}
							rOSPoliciesValResourceGroupsValResourcesValPkgDeb["source"] = rOSPoliciesValResourceGroupsValResourcesValPkgDebSource
						}
						rOSPoliciesValResourceGroupsValResourcesValPkg["deb"] = rOSPoliciesValResourceGroupsValResourcesValPkgDeb
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.DesiredState != nil {
						rOSPoliciesValResourceGroupsValResourcesValPkg["desiredState"] = string(*rOSPoliciesValResourceGroupsValResourcesVal.Pkg.DesiredState)
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Googet != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Googet != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget {
						rOSPoliciesValResourceGroupsValResourcesValPkgGooget := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Googet.Name != nil {
							rOSPoliciesValResourceGroupsValResourcesValPkgGooget["name"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Googet.Name
						}
						rOSPoliciesValResourceGroupsValResourcesValPkg["googet"] = rOSPoliciesValResourceGroupsValResourcesValPkgGooget
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi {
						rOSPoliciesValResourceGroupsValResourcesValPkgMsi := make(map[string]interface{})
						var rOSPoliciesValResourceGroupsValResourcesValPkgMsiProperties []interface{}
						for _, rOSPoliciesValResourceGroupsValResourcesValPkgMsiPropertiesVal := range rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Properties {
							rOSPoliciesValResourceGroupsValResourcesValPkgMsiProperties = append(rOSPoliciesValResourceGroupsValResourcesValPkgMsiProperties, rOSPoliciesValResourceGroupsValResourcesValPkgMsiPropertiesVal)
						}
						rOSPoliciesValResourceGroupsValResourcesValPkgMsi["properties"] = rOSPoliciesValResourceGroupsValResourcesValPkgMsiProperties
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource {
							rOSPoliciesValResourceGroupsValResourcesValPkgMsiSource := make(map[string]interface{})
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.AllowInsecure != nil {
								rOSPoliciesValResourceGroupsValResourcesValPkgMsiSource["allowInsecure"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.AllowInsecure
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Gcs != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Gcs != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs {
								rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceGcs := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Gcs.Bucket != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceGcs["bucket"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Gcs.Bucket
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Gcs.Generation != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceGcs["generation"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Gcs.Generation
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Gcs.Object != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceGcs["object"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Gcs.Object
								}
								rOSPoliciesValResourceGroupsValResourcesValPkgMsiSource["gcs"] = rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceGcs
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.LocalPath != nil {
								rOSPoliciesValResourceGroupsValResourcesValPkgMsiSource["localPath"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.LocalPath
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Remote != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Remote != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote {
								rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceRemote := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Remote.Sha256Checksum != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceRemote["sha256Checksum"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Remote.Sha256Checksum
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Remote.Uri != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceRemote["uri"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Msi.Source.Remote.Uri
								}
								rOSPoliciesValResourceGroupsValResourcesValPkgMsiSource["remote"] = rOSPoliciesValResourceGroupsValResourcesValPkgMsiSourceRemote
							}
							rOSPoliciesValResourceGroupsValResourcesValPkgMsi["source"] = rOSPoliciesValResourceGroupsValResourcesValPkgMsiSource
						}
						rOSPoliciesValResourceGroupsValResourcesValPkg["msi"] = rOSPoliciesValResourceGroupsValResourcesValPkgMsi
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm {
						rOSPoliciesValResourceGroupsValResourcesValPkgRpm := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.PullDeps != nil {
							rOSPoliciesValResourceGroupsValResourcesValPkgRpm["pullDeps"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.PullDeps
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource {
							rOSPoliciesValResourceGroupsValResourcesValPkgRpmSource := make(map[string]interface{})
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.AllowInsecure != nil {
								rOSPoliciesValResourceGroupsValResourcesValPkgRpmSource["allowInsecure"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.AllowInsecure
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Gcs != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Gcs != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs {
								rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceGcs := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Gcs.Bucket != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceGcs["bucket"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Gcs.Bucket
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Gcs.Generation != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceGcs["generation"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Gcs.Generation
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Gcs.Object != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceGcs["object"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Gcs.Object
								}
								rOSPoliciesValResourceGroupsValResourcesValPkgRpmSource["gcs"] = rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceGcs
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.LocalPath != nil {
								rOSPoliciesValResourceGroupsValResourcesValPkgRpmSource["localPath"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.LocalPath
							}
							if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Remote != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Remote != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote {
								rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceRemote := make(map[string]interface{})
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Remote.Sha256Checksum != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceRemote["sha256Checksum"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Remote.Sha256Checksum
								}
								if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Remote.Uri != nil {
									rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceRemote["uri"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Rpm.Source.Remote.Uri
								}
								rOSPoliciesValResourceGroupsValResourcesValPkgRpmSource["remote"] = rOSPoliciesValResourceGroupsValResourcesValPkgRpmSourceRemote
							}
							rOSPoliciesValResourceGroupsValResourcesValPkgRpm["source"] = rOSPoliciesValResourceGroupsValResourcesValPkgRpmSource
						}
						rOSPoliciesValResourceGroupsValResourcesValPkg["rpm"] = rOSPoliciesValResourceGroupsValResourcesValPkgRpm
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Yum != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Yum != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum {
						rOSPoliciesValResourceGroupsValResourcesValPkgYum := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Yum.Name != nil {
							rOSPoliciesValResourceGroupsValResourcesValPkgYum["name"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Yum.Name
						}
						rOSPoliciesValResourceGroupsValResourcesValPkg["yum"] = rOSPoliciesValResourceGroupsValResourcesValPkgYum
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Zypper != nil && rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Zypper != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper {
						rOSPoliciesValResourceGroupsValResourcesValPkgZypper := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Zypper.Name != nil {
							rOSPoliciesValResourceGroupsValResourcesValPkgZypper["name"] = *rOSPoliciesValResourceGroupsValResourcesVal.Pkg.Zypper.Name
						}
						rOSPoliciesValResourceGroupsValResourcesValPkg["zypper"] = rOSPoliciesValResourceGroupsValResourcesValPkgZypper
					}
					rOSPoliciesValResourceGroupsValResourcesObject["pkg"] = rOSPoliciesValResourceGroupsValResourcesValPkg
				}
				if rOSPoliciesValResourceGroupsValResourcesVal.Repository != nil && rOSPoliciesValResourceGroupsValResourcesVal.Repository != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository {
					rOSPoliciesValResourceGroupsValResourcesValRepository := make(map[string]interface{})
					if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt != nil && rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt {
						rOSPoliciesValResourceGroupsValResourcesValRepositoryApt := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.ArchiveType != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryApt["archiveType"] = string(*rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.ArchiveType)
						}
						var rOSPoliciesValResourceGroupsValResourcesValRepositoryAptComponents []interface{}
						for _, rOSPoliciesValResourceGroupsValResourcesValRepositoryAptComponentsVal := range rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.Components {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryAptComponents = append(rOSPoliciesValResourceGroupsValResourcesValRepositoryAptComponents, rOSPoliciesValResourceGroupsValResourcesValRepositoryAptComponentsVal)
						}
						rOSPoliciesValResourceGroupsValResourcesValRepositoryApt["components"] = rOSPoliciesValResourceGroupsValResourcesValRepositoryAptComponents
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.Distribution != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryApt["distribution"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.Distribution
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.GpgKey != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryApt["gpgKey"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.GpgKey
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.Uri != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryApt["uri"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Apt.Uri
						}
						rOSPoliciesValResourceGroupsValResourcesValRepository["apt"] = rOSPoliciesValResourceGroupsValResourcesValRepositoryApt
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Goo != nil && rOSPoliciesValResourceGroupsValResourcesVal.Repository.Goo != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo {
						rOSPoliciesValResourceGroupsValResourcesValRepositoryGoo := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Goo.Name != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryGoo["name"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Goo.Name
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Goo.Url != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryGoo["url"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Goo.Url
						}
						rOSPoliciesValResourceGroupsValResourcesValRepository["goo"] = rOSPoliciesValResourceGroupsValResourcesValRepositoryGoo
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum != nil && rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum {
						rOSPoliciesValResourceGroupsValResourcesValRepositoryYum := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum.BaseUrl != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryYum["baseUrl"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum.BaseUrl
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum.DisplayName != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryYum["displayName"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum.DisplayName
						}
						var rOSPoliciesValResourceGroupsValResourcesValRepositoryYumGpgKeys []interface{}
						for _, rOSPoliciesValResourceGroupsValResourcesValRepositoryYumGpgKeysVal := range rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum.GpgKeys {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryYumGpgKeys = append(rOSPoliciesValResourceGroupsValResourcesValRepositoryYumGpgKeys, rOSPoliciesValResourceGroupsValResourcesValRepositoryYumGpgKeysVal)
						}
						rOSPoliciesValResourceGroupsValResourcesValRepositoryYum["gpgKeys"] = rOSPoliciesValResourceGroupsValResourcesValRepositoryYumGpgKeys
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum.Id != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryYum["id"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Yum.Id
						}
						rOSPoliciesValResourceGroupsValResourcesValRepository["yum"] = rOSPoliciesValResourceGroupsValResourcesValRepositoryYum
					}
					if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper != nil && rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper != dclService.EmptyOSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper {
						rOSPoliciesValResourceGroupsValResourcesValRepositoryZypper := make(map[string]interface{})
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper.BaseUrl != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryZypper["baseUrl"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper.BaseUrl
						}
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper.DisplayName != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryZypper["displayName"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper.DisplayName
						}
						var rOSPoliciesValResourceGroupsValResourcesValRepositoryZypperGpgKeys []interface{}
						for _, rOSPoliciesValResourceGroupsValResourcesValRepositoryZypperGpgKeysVal := range rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper.GpgKeys {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryZypperGpgKeys = append(rOSPoliciesValResourceGroupsValResourcesValRepositoryZypperGpgKeys, rOSPoliciesValResourceGroupsValResourcesValRepositoryZypperGpgKeysVal)
						}
						rOSPoliciesValResourceGroupsValResourcesValRepositoryZypper["gpgKeys"] = rOSPoliciesValResourceGroupsValResourcesValRepositoryZypperGpgKeys
						if rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper.Id != nil {
							rOSPoliciesValResourceGroupsValResourcesValRepositoryZypper["id"] = *rOSPoliciesValResourceGroupsValResourcesVal.Repository.Zypper.Id
						}
						rOSPoliciesValResourceGroupsValResourcesValRepository["zypper"] = rOSPoliciesValResourceGroupsValResourcesValRepositoryZypper
					}
					rOSPoliciesValResourceGroupsValResourcesObject["repository"] = rOSPoliciesValResourceGroupsValResourcesValRepository
				}
				rOSPoliciesValResourceGroupsValResources = append(rOSPoliciesValResourceGroupsValResources, rOSPoliciesValResourceGroupsValResourcesObject)
			}
			rOSPoliciesValResourceGroupsObject["resources"] = rOSPoliciesValResourceGroupsValResources
			rOSPoliciesValResourceGroups = append(rOSPoliciesValResourceGroups, rOSPoliciesValResourceGroupsObject)
		}
		rOSPoliciesObject["resourceGroups"] = rOSPoliciesValResourceGroups
		rOSPolicies = append(rOSPolicies, rOSPoliciesObject)
	}
	u.Object["osPolicies"] = rOSPolicies
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Reconciling != nil {
		u.Object["reconciling"] = *r.Reconciling
	}
	if r.RevisionCreateTime != nil {
		u.Object["revisionCreateTime"] = *r.RevisionCreateTime
	}
	if r.RevisionId != nil {
		u.Object["revisionId"] = *r.RevisionId
	}
	if r.Rollout != nil && r.Rollout != dclService.EmptyOSPolicyAssignmentRollout {
		rRollout := make(map[string]interface{})
		if r.Rollout.DisruptionBudget != nil && r.Rollout.DisruptionBudget != dclService.EmptyOSPolicyAssignmentRolloutDisruptionBudget {
			rRolloutDisruptionBudget := make(map[string]interface{})
			if r.Rollout.DisruptionBudget.Fixed != nil {
				rRolloutDisruptionBudget["fixed"] = *r.Rollout.DisruptionBudget.Fixed
			}
			if r.Rollout.DisruptionBudget.Percent != nil {
				rRolloutDisruptionBudget["percent"] = *r.Rollout.DisruptionBudget.Percent
			}
			rRollout["disruptionBudget"] = rRolloutDisruptionBudget
		}
		if r.Rollout.MinWaitDuration != nil {
			rRollout["minWaitDuration"] = *r.Rollout.MinWaitDuration
		}
		u.Object["rollout"] = rRollout
	}
	if r.RolloutState != nil {
		u.Object["rolloutState"] = string(*r.RolloutState)
	}
	if r.SkipAwaitRollout != nil {
		u.Object["skipAwaitRollout"] = *r.SkipAwaitRollout
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	return u
}

func UnstructuredToOSPolicyAssignment(u *unstructured.Resource) (*dclService.OSPolicyAssignment, error) {
	r := &dclService.OSPolicyAssignment{}
	if _, ok := u.Object["baseline"]; ok {
		if b, ok := u.Object["baseline"].(bool); ok {
			r.Baseline = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Baseline: expected bool")
		}
	}
	if _, ok := u.Object["deleted"]; ok {
		if b, ok := u.Object["deleted"].(bool); ok {
			r.Deleted = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Deleted: expected bool")
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
	if _, ok := u.Object["instanceFilter"]; ok {
		if rInstanceFilter, ok := u.Object["instanceFilter"].(map[string]interface{}); ok {
			r.InstanceFilter = &dclService.OSPolicyAssignmentInstanceFilter{}
			if _, ok := rInstanceFilter["all"]; ok {
				if b, ok := rInstanceFilter["all"].(bool); ok {
					r.InstanceFilter.All = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.All: expected bool")
				}
			}
			if _, ok := rInstanceFilter["exclusionLabels"]; ok {
				if s, ok := rInstanceFilter["exclusionLabels"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rInstanceFilterExclusionLabels dclService.OSPolicyAssignmentInstanceFilterExclusionLabels
							if _, ok := objval["labels"]; ok {
								if rInstanceFilterExclusionLabelsLabels, ok := objval["labels"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rInstanceFilterExclusionLabelsLabels {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rInstanceFilterExclusionLabels.Labels = m
								} else {
									return nil, fmt.Errorf("rInstanceFilterExclusionLabels.Labels: expected map[string]interface{}")
								}
							}
							r.InstanceFilter.ExclusionLabels = append(r.InstanceFilter.ExclusionLabels, rInstanceFilterExclusionLabels)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.ExclusionLabels: expected []interface{}")
				}
			}
			if _, ok := rInstanceFilter["inclusionLabels"]; ok {
				if s, ok := rInstanceFilter["inclusionLabels"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rInstanceFilterInclusionLabels dclService.OSPolicyAssignmentInstanceFilterInclusionLabels
							if _, ok := objval["labels"]; ok {
								if rInstanceFilterInclusionLabelsLabels, ok := objval["labels"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rInstanceFilterInclusionLabelsLabels {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									rInstanceFilterInclusionLabels.Labels = m
								} else {
									return nil, fmt.Errorf("rInstanceFilterInclusionLabels.Labels: expected map[string]interface{}")
								}
							}
							r.InstanceFilter.InclusionLabels = append(r.InstanceFilter.InclusionLabels, rInstanceFilterInclusionLabels)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.InclusionLabels: expected []interface{}")
				}
			}
			if _, ok := rInstanceFilter["inventories"]; ok {
				if s, ok := rInstanceFilter["inventories"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rInstanceFilterInventories dclService.OSPolicyAssignmentInstanceFilterInventories
							if _, ok := objval["osShortName"]; ok {
								if s, ok := objval["osShortName"].(string); ok {
									rInstanceFilterInventories.OSShortName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rInstanceFilterInventories.OSShortName: expected string")
								}
							}
							if _, ok := objval["osVersion"]; ok {
								if s, ok := objval["osVersion"].(string); ok {
									rInstanceFilterInventories.OSVersion = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rInstanceFilterInventories.OSVersion: expected string")
								}
							}
							r.InstanceFilter.Inventories = append(r.InstanceFilter.Inventories, rInstanceFilterInventories)
						}
					}
				} else {
					return nil, fmt.Errorf("r.InstanceFilter.Inventories: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.InstanceFilter: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["osPolicies"]; ok {
		if s, ok := u.Object["osPolicies"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rOSPolicies dclService.OSPolicyAssignmentOSPolicies
					if _, ok := objval["allowNoResourceGroupMatch"]; ok {
						if b, ok := objval["allowNoResourceGroupMatch"].(bool); ok {
							rOSPolicies.AllowNoResourceGroupMatch = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("rOSPolicies.AllowNoResourceGroupMatch: expected bool")
						}
					}
					if _, ok := objval["description"]; ok {
						if s, ok := objval["description"].(string); ok {
							rOSPolicies.Description = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rOSPolicies.Description: expected string")
						}
					}
					if _, ok := objval["id"]; ok {
						if s, ok := objval["id"].(string); ok {
							rOSPolicies.Id = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rOSPolicies.Id: expected string")
						}
					}
					if _, ok := objval["mode"]; ok {
						if s, ok := objval["mode"].(string); ok {
							rOSPolicies.Mode = dclService.OSPolicyAssignmentOSPoliciesModeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rOSPolicies.Mode: expected string")
						}
					}
					if _, ok := objval["resourceGroups"]; ok {
						if s, ok := objval["resourceGroups"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rOSPoliciesResourceGroups dclService.OSPolicyAssignmentOSPoliciesResourceGroups
									if _, ok := objval["inventoryFilters"]; ok {
										if s, ok := objval["inventoryFilters"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rOSPoliciesResourceGroupsInventoryFilters dclService.OSPolicyAssignmentOSPoliciesResourceGroupsInventoryFilters
													if _, ok := objval["osShortName"]; ok {
														if s, ok := objval["osShortName"].(string); ok {
															rOSPoliciesResourceGroupsInventoryFilters.OSShortName = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rOSPoliciesResourceGroupsInventoryFilters.OSShortName: expected string")
														}
													}
													if _, ok := objval["osVersion"]; ok {
														if s, ok := objval["osVersion"].(string); ok {
															rOSPoliciesResourceGroupsInventoryFilters.OSVersion = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rOSPoliciesResourceGroupsInventoryFilters.OSVersion: expected string")
														}
													}
													rOSPoliciesResourceGroups.InventoryFilters = append(rOSPoliciesResourceGroups.InventoryFilters, rOSPoliciesResourceGroupsInventoryFilters)
												}
											}
										} else {
											return nil, fmt.Errorf("rOSPoliciesResourceGroups.InventoryFilters: expected []interface{}")
										}
									}
									if _, ok := objval["resources"]; ok {
										if s, ok := objval["resources"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rOSPoliciesResourceGroupsResources dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResources
													if _, ok := objval["exec"]; ok {
														if rOSPoliciesResourceGroupsResourcesExec, ok := objval["exec"].(map[string]interface{}); ok {
															rOSPoliciesResourceGroupsResources.Exec = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExec{}
															if _, ok := rOSPoliciesResourceGroupsResourcesExec["enforce"]; ok {
																if rOSPoliciesResourceGroupsResourcesExecEnforce, ok := rOSPoliciesResourceGroupsResourcesExec["enforce"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Exec.Enforce = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforce{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["args"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["args"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Enforce.Args = append(rOSPoliciesResourceGroupsResources.Exec.Enforce.Args, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.Args: expected []interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["file"]; ok {
																		if rOSPoliciesResourceGroupsResourcesExecEnforceFile, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["file"].(map[string]interface{}); ok {
																			rOSPoliciesResourceGroupsResources.Exec.Enforce.File = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFile{}
																			if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFile["allowInsecure"]; ok {
																				if b, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFile["allowInsecure"].(bool); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Enforce.File.AllowInsecure = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.AllowInsecure: expected bool")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFile["gcs"]; ok {
																				if rOSPoliciesResourceGroupsResourcesExecEnforceFileGcs, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFile["gcs"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Gcs = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileGcs{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileGcs["bucket"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileGcs["bucket"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Gcs.Bucket = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Gcs.Bucket: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileGcs["generation"]; ok {
																						if i, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileGcs["generation"].(int64); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Gcs.Generation = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Gcs.Generation: expected int64")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileGcs["object"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileGcs["object"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Gcs.Object = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Gcs.Object: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Gcs: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFile["localPath"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFile["localPath"].(string); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Enforce.File.LocalPath = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.LocalPath: expected string")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFile["remote"]; ok {
																				if rOSPoliciesResourceGroupsResourcesExecEnforceFileRemote, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFile["remote"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Remote = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceFileRemote{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileRemote["sha256Checksum"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileRemote["sha256Checksum"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Remote.Sha256Checksum = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Remote.Sha256Checksum: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileRemote["uri"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforceFileRemote["uri"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Remote.Uri = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Remote.Uri: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File.Remote: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.File: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["interpreter"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["interpreter"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Exec.Enforce.Interpreter = dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecEnforceInterpreterEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.Interpreter: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["outputFilePath"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["outputFilePath"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Exec.Enforce.OutputFilePath = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.OutputFilePath: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["script"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesExecEnforce["script"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Exec.Enforce.Script = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce.Script: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Enforce: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesExec["validate"]; ok {
																if rOSPoliciesResourceGroupsResourcesExecValidate, ok := rOSPoliciesResourceGroupsResourcesExec["validate"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Exec.Validate = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidate{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecValidate["args"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesExecValidate["args"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Validate.Args = append(rOSPoliciesResourceGroupsResources.Exec.Validate.Args, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.Args: expected []interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecValidate["file"]; ok {
																		if rOSPoliciesResourceGroupsResourcesExecValidateFile, ok := rOSPoliciesResourceGroupsResourcesExecValidate["file"].(map[string]interface{}); ok {
																			rOSPoliciesResourceGroupsResources.Exec.Validate.File = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFile{}
																			if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFile["allowInsecure"]; ok {
																				if b, ok := rOSPoliciesResourceGroupsResourcesExecValidateFile["allowInsecure"].(bool); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Validate.File.AllowInsecure = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.AllowInsecure: expected bool")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFile["gcs"]; ok {
																				if rOSPoliciesResourceGroupsResourcesExecValidateFileGcs, ok := rOSPoliciesResourceGroupsResourcesExecValidateFile["gcs"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Validate.File.Gcs = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileGcs{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileGcs["bucket"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileGcs["bucket"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Validate.File.Gcs.Bucket = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.Gcs.Bucket: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileGcs["generation"]; ok {
																						if i, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileGcs["generation"].(int64); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Validate.File.Gcs.Generation = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.Gcs.Generation: expected int64")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileGcs["object"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileGcs["object"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Validate.File.Gcs.Object = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.Gcs.Object: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.Gcs: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFile["localPath"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesExecValidateFile["localPath"].(string); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Validate.File.LocalPath = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.LocalPath: expected string")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFile["remote"]; ok {
																				if rOSPoliciesResourceGroupsResourcesExecValidateFileRemote, ok := rOSPoliciesResourceGroupsResourcesExecValidateFile["remote"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Exec.Validate.File.Remote = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateFileRemote{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileRemote["sha256Checksum"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileRemote["sha256Checksum"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Validate.File.Remote.Sha256Checksum = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.Remote.Sha256Checksum: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileRemote["uri"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesExecValidateFileRemote["uri"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Exec.Validate.File.Remote.Uri = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.Remote.Uri: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File.Remote: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.File: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecValidate["interpreter"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesExecValidate["interpreter"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Exec.Validate.Interpreter = dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesExecValidateInterpreterEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.Interpreter: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecValidate["outputFilePath"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesExecValidate["outputFilePath"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Exec.Validate.OutputFilePath = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.OutputFilePath: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesExecValidate["script"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesExecValidate["script"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Exec.Validate.Script = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate.Script: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec.Validate: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Exec: expected map[string]interface{}")
														}
													}
													if _, ok := objval["file"]; ok {
														if rOSPoliciesResourceGroupsResourcesFile, ok := objval["file"].(map[string]interface{}); ok {
															rOSPoliciesResourceGroupsResources.File = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFile{}
															if _, ok := rOSPoliciesResourceGroupsResourcesFile["content"]; ok {
																if s, ok := rOSPoliciesResourceGroupsResourcesFile["content"].(string); ok {
																	rOSPoliciesResourceGroupsResources.File.Content = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.Content: expected string")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesFile["file"]; ok {
																if rOSPoliciesResourceGroupsResourcesFileFile, ok := rOSPoliciesResourceGroupsResourcesFile["file"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.File.File = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFile{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesFileFile["allowInsecure"]; ok {
																		if b, ok := rOSPoliciesResourceGroupsResourcesFileFile["allowInsecure"].(bool); ok {
																			rOSPoliciesResourceGroupsResources.File.File.AllowInsecure = dcl.Bool(b)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.AllowInsecure: expected bool")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesFileFile["gcs"]; ok {
																		if rOSPoliciesResourceGroupsResourcesFileFileGcs, ok := rOSPoliciesResourceGroupsResourcesFileFile["gcs"].(map[string]interface{}); ok {
																			rOSPoliciesResourceGroupsResources.File.File.Gcs = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileGcs{}
																			if _, ok := rOSPoliciesResourceGroupsResourcesFileFileGcs["bucket"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesFileFileGcs["bucket"].(string); ok {
																					rOSPoliciesResourceGroupsResources.File.File.Gcs.Bucket = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.Gcs.Bucket: expected string")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesFileFileGcs["generation"]; ok {
																				if i, ok := rOSPoliciesResourceGroupsResourcesFileFileGcs["generation"].(int64); ok {
																					rOSPoliciesResourceGroupsResources.File.File.Gcs.Generation = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.Gcs.Generation: expected int64")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesFileFileGcs["object"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesFileFileGcs["object"].(string); ok {
																					rOSPoliciesResourceGroupsResources.File.File.Gcs.Object = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.Gcs.Object: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.Gcs: expected map[string]interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesFileFile["localPath"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesFileFile["localPath"].(string); ok {
																			rOSPoliciesResourceGroupsResources.File.File.LocalPath = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.LocalPath: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesFileFile["remote"]; ok {
																		if rOSPoliciesResourceGroupsResourcesFileFileRemote, ok := rOSPoliciesResourceGroupsResourcesFileFile["remote"].(map[string]interface{}); ok {
																			rOSPoliciesResourceGroupsResources.File.File.Remote = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileFileRemote{}
																			if _, ok := rOSPoliciesResourceGroupsResourcesFileFileRemote["sha256Checksum"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesFileFileRemote["sha256Checksum"].(string); ok {
																					rOSPoliciesResourceGroupsResources.File.File.Remote.Sha256Checksum = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.Remote.Sha256Checksum: expected string")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesFileFileRemote["uri"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesFileFileRemote["uri"].(string); ok {
																					rOSPoliciesResourceGroupsResources.File.File.Remote.Uri = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.Remote.Uri: expected string")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File.Remote: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.File: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesFile["path"]; ok {
																if s, ok := rOSPoliciesResourceGroupsResourcesFile["path"].(string); ok {
																	rOSPoliciesResourceGroupsResources.File.Path = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.Path: expected string")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesFile["permissions"]; ok {
																if s, ok := rOSPoliciesResourceGroupsResourcesFile["permissions"].(string); ok {
																	rOSPoliciesResourceGroupsResources.File.Permissions = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.Permissions: expected string")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesFile["state"]; ok {
																if s, ok := rOSPoliciesResourceGroupsResourcesFile["state"].(string); ok {
																	rOSPoliciesResourceGroupsResources.File.State = dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesFileStateEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File.State: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.File: expected map[string]interface{}")
														}
													}
													if _, ok := objval["id"]; ok {
														if s, ok := objval["id"].(string); ok {
															rOSPoliciesResourceGroupsResources.Id = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Id: expected string")
														}
													}
													if _, ok := objval["pkg"]; ok {
														if rOSPoliciesResourceGroupsResourcesPkg, ok := objval["pkg"].(map[string]interface{}); ok {
															rOSPoliciesResourceGroupsResources.Pkg = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkg{}
															if _, ok := rOSPoliciesResourceGroupsResourcesPkg["apt"]; ok {
																if rOSPoliciesResourceGroupsResourcesPkgApt, ok := rOSPoliciesResourceGroupsResourcesPkg["apt"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Pkg.Apt = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgApt{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgApt["name"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesPkgApt["name"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Apt.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Apt.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Apt: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesPkg["deb"]; ok {
																if rOSPoliciesResourceGroupsResourcesPkgDeb, ok := rOSPoliciesResourceGroupsResourcesPkg["deb"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Pkg.Deb = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDeb{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgDeb["pullDeps"]; ok {
																		if b, ok := rOSPoliciesResourceGroupsResourcesPkgDeb["pullDeps"].(bool); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Deb.PullDeps = dcl.Bool(b)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.PullDeps: expected bool")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgDeb["source"]; ok {
																		if rOSPoliciesResourceGroupsResourcesPkgDebSource, ok := rOSPoliciesResourceGroupsResourcesPkgDeb["source"].(map[string]interface{}); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Deb.Source = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSource{}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSource["allowInsecure"]; ok {
																				if b, ok := rOSPoliciesResourceGroupsResourcesPkgDebSource["allowInsecure"].(bool); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.AllowInsecure = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.AllowInsecure: expected bool")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSource["gcs"]; ok {
																				if rOSPoliciesResourceGroupsResourcesPkgDebSourceGcs, ok := rOSPoliciesResourceGroupsResourcesPkgDebSource["gcs"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Gcs = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceGcs{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceGcs["bucket"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceGcs["bucket"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Gcs.Bucket = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Gcs.Bucket: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceGcs["generation"]; ok {
																						if i, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceGcs["generation"].(int64); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Gcs.Generation = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Gcs.Generation: expected int64")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceGcs["object"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceGcs["object"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Gcs.Object = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Gcs.Object: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Gcs: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSource["localPath"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesPkgDebSource["localPath"].(string); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.LocalPath = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.LocalPath: expected string")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSource["remote"]; ok {
																				if rOSPoliciesResourceGroupsResourcesPkgDebSourceRemote, ok := rOSPoliciesResourceGroupsResourcesPkgDebSource["remote"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Remote = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDebSourceRemote{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceRemote["sha256Checksum"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceRemote["sha256Checksum"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Remote.Sha256Checksum = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Remote.Sha256Checksum: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceRemote["uri"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgDebSourceRemote["uri"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Remote.Uri = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Remote.Uri: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source.Remote: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb.Source: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Deb: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesPkg["desiredState"]; ok {
																if s, ok := rOSPoliciesResourceGroupsResourcesPkg["desiredState"].(string); ok {
																	rOSPoliciesResourceGroupsResources.Pkg.DesiredState = dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgDesiredStateEnumRef(s)
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.DesiredState: expected string")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesPkg["googet"]; ok {
																if rOSPoliciesResourceGroupsResourcesPkgGooget, ok := rOSPoliciesResourceGroupsResourcesPkg["googet"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Pkg.Googet = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgGooget{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgGooget["name"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesPkgGooget["name"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Googet.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Googet.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Googet: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesPkg["msi"]; ok {
																if rOSPoliciesResourceGroupsResourcesPkgMsi, ok := rOSPoliciesResourceGroupsResourcesPkg["msi"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Pkg.Msi = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsi{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsi["properties"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesPkgMsi["properties"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Msi.Properties = append(rOSPoliciesResourceGroupsResources.Pkg.Msi.Properties, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Properties: expected []interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsi["source"]; ok {
																		if rOSPoliciesResourceGroupsResourcesPkgMsiSource, ok := rOSPoliciesResourceGroupsResourcesPkgMsi["source"].(map[string]interface{}); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Msi.Source = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSource{}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSource["allowInsecure"]; ok {
																				if b, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSource["allowInsecure"].(bool); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.AllowInsecure = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.AllowInsecure: expected bool")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSource["gcs"]; ok {
																				if rOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSource["gcs"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Gcs = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs["bucket"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs["bucket"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Gcs.Bucket = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Gcs.Bucket: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs["generation"]; ok {
																						if i, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs["generation"].(int64); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Gcs.Generation = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Gcs.Generation: expected int64")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs["object"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceGcs["object"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Gcs.Object = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Gcs.Object: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Gcs: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSource["localPath"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSource["localPath"].(string); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.LocalPath = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.LocalPath: expected string")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSource["remote"]; ok {
																				if rOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSource["remote"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Remote = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote["sha256Checksum"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote["sha256Checksum"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Remote.Sha256Checksum = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Remote.Sha256Checksum: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote["uri"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgMsiSourceRemote["uri"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Remote.Uri = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Remote.Uri: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source.Remote: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi.Source: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Msi: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesPkg["rpm"]; ok {
																if rOSPoliciesResourceGroupsResourcesPkgRpm, ok := rOSPoliciesResourceGroupsResourcesPkg["rpm"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Pkg.Rpm = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpm{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpm["pullDeps"]; ok {
																		if b, ok := rOSPoliciesResourceGroupsResourcesPkgRpm["pullDeps"].(bool); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Rpm.PullDeps = dcl.Bool(b)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.PullDeps: expected bool")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpm["source"]; ok {
																		if rOSPoliciesResourceGroupsResourcesPkgRpmSource, ok := rOSPoliciesResourceGroupsResourcesPkgRpm["source"].(map[string]interface{}); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSource{}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSource["allowInsecure"]; ok {
																				if b, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSource["allowInsecure"].(bool); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.AllowInsecure = dcl.Bool(b)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.AllowInsecure: expected bool")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSource["gcs"]; ok {
																				if rOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSource["gcs"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Gcs = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs["bucket"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs["bucket"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Gcs.Bucket = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Gcs.Bucket: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs["generation"]; ok {
																						if i, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs["generation"].(int64); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Gcs.Generation = dcl.Int64(i)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Gcs.Generation: expected int64")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs["object"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceGcs["object"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Gcs.Object = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Gcs.Object: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Gcs: expected map[string]interface{}")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSource["localPath"]; ok {
																				if s, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSource["localPath"].(string); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.LocalPath = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.LocalPath: expected string")
																				}
																			}
																			if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSource["remote"]; ok {
																				if rOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSource["remote"].(map[string]interface{}); ok {
																					rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Remote = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote{}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote["sha256Checksum"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote["sha256Checksum"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Remote.Sha256Checksum = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Remote.Sha256Checksum: expected string")
																						}
																					}
																					if _, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote["uri"]; ok {
																						if s, ok := rOSPoliciesResourceGroupsResourcesPkgRpmSourceRemote["uri"].(string); ok {
																							rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Remote.Uri = dcl.String(s)
																						} else {
																							return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Remote.Uri: expected string")
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source.Remote: expected map[string]interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm.Source: expected map[string]interface{}")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Rpm: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesPkg["yum"]; ok {
																if rOSPoliciesResourceGroupsResourcesPkgYum, ok := rOSPoliciesResourceGroupsResourcesPkg["yum"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Pkg.Yum = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgYum{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgYum["name"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesPkgYum["name"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Yum.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Yum.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Yum: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesPkg["zypper"]; ok {
																if rOSPoliciesResourceGroupsResourcesPkgZypper, ok := rOSPoliciesResourceGroupsResourcesPkg["zypper"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Pkg.Zypper = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesPkgZypper{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesPkgZypper["name"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesPkgZypper["name"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Pkg.Zypper.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Zypper.Name: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg.Zypper: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Pkg: expected map[string]interface{}")
														}
													}
													if _, ok := objval["repository"]; ok {
														if rOSPoliciesResourceGroupsResourcesRepository, ok := objval["repository"].(map[string]interface{}); ok {
															rOSPoliciesResourceGroupsResources.Repository = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepository{}
															if _, ok := rOSPoliciesResourceGroupsResourcesRepository["apt"]; ok {
																if rOSPoliciesResourceGroupsResourcesRepositoryApt, ok := rOSPoliciesResourceGroupsResourcesRepository["apt"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Repository.Apt = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryApt{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["archiveType"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["archiveType"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Apt.ArchiveType = dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryAptArchiveTypeEnumRef(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Apt.ArchiveType: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["components"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["components"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rOSPoliciesResourceGroupsResources.Repository.Apt.Components = append(rOSPoliciesResourceGroupsResources.Repository.Apt.Components, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Apt.Components: expected []interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["distribution"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["distribution"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Apt.Distribution = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Apt.Distribution: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["gpgKey"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["gpgKey"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Apt.GpgKey = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Apt.GpgKey: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["uri"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryApt["uri"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Apt.Uri = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Apt.Uri: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Apt: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesRepository["goo"]; ok {
																if rOSPoliciesResourceGroupsResourcesRepositoryGoo, ok := rOSPoliciesResourceGroupsResourcesRepository["goo"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Repository.Goo = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryGoo{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryGoo["name"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryGoo["name"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Goo.Name = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Goo.Name: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryGoo["url"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryGoo["url"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Goo.Url = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Goo.Url: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Goo: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesRepository["yum"]; ok {
																if rOSPoliciesResourceGroupsResourcesRepositoryYum, ok := rOSPoliciesResourceGroupsResourcesRepository["yum"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Repository.Yum = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryYum{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryYum["baseUrl"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryYum["baseUrl"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Yum.BaseUrl = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Yum.BaseUrl: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryYum["displayName"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryYum["displayName"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Yum.DisplayName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Yum.DisplayName: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryYum["gpgKeys"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryYum["gpgKeys"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rOSPoliciesResourceGroupsResources.Repository.Yum.GpgKeys = append(rOSPoliciesResourceGroupsResources.Repository.Yum.GpgKeys, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Yum.GpgKeys: expected []interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryYum["id"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryYum["id"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Yum.Id = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Yum.Id: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Yum: expected map[string]interface{}")
																}
															}
															if _, ok := rOSPoliciesResourceGroupsResourcesRepository["zypper"]; ok {
																if rOSPoliciesResourceGroupsResourcesRepositoryZypper, ok := rOSPoliciesResourceGroupsResourcesRepository["zypper"].(map[string]interface{}); ok {
																	rOSPoliciesResourceGroupsResources.Repository.Zypper = &dclService.OSPolicyAssignmentOSPoliciesResourceGroupsResourcesRepositoryZypper{}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryZypper["baseUrl"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryZypper["baseUrl"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Zypper.BaseUrl = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Zypper.BaseUrl: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryZypper["displayName"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryZypper["displayName"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Zypper.DisplayName = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Zypper.DisplayName: expected string")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryZypper["gpgKeys"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryZypper["gpgKeys"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rOSPoliciesResourceGroupsResources.Repository.Zypper.GpgKeys = append(rOSPoliciesResourceGroupsResources.Repository.Zypper.GpgKeys, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Zypper.GpgKeys: expected []interface{}")
																		}
																	}
																	if _, ok := rOSPoliciesResourceGroupsResourcesRepositoryZypper["id"]; ok {
																		if s, ok := rOSPoliciesResourceGroupsResourcesRepositoryZypper["id"].(string); ok {
																			rOSPoliciesResourceGroupsResources.Repository.Zypper.Id = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Zypper.Id: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository.Zypper: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rOSPoliciesResourceGroupsResources.Repository: expected map[string]interface{}")
														}
													}
													rOSPoliciesResourceGroups.Resources = append(rOSPoliciesResourceGroups.Resources, rOSPoliciesResourceGroupsResources)
												}
											}
										} else {
											return nil, fmt.Errorf("rOSPoliciesResourceGroups.Resources: expected []interface{}")
										}
									}
									rOSPolicies.ResourceGroups = append(rOSPolicies.ResourceGroups, rOSPoliciesResourceGroups)
								}
							}
						} else {
							return nil, fmt.Errorf("rOSPolicies.ResourceGroups: expected []interface{}")
						}
					}
					r.OSPolicies = append(r.OSPolicies, rOSPolicies)
				}
			}
		} else {
			return nil, fmt.Errorf("r.OSPolicies: expected []interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["reconciling"]; ok {
		if b, ok := u.Object["reconciling"].(bool); ok {
			r.Reconciling = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Reconciling: expected bool")
		}
	}
	if _, ok := u.Object["revisionCreateTime"]; ok {
		if s, ok := u.Object["revisionCreateTime"].(string); ok {
			r.RevisionCreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.RevisionCreateTime: expected string")
		}
	}
	if _, ok := u.Object["revisionId"]; ok {
		if s, ok := u.Object["revisionId"].(string); ok {
			r.RevisionId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.RevisionId: expected string")
		}
	}
	if _, ok := u.Object["rollout"]; ok {
		if rRollout, ok := u.Object["rollout"].(map[string]interface{}); ok {
			r.Rollout = &dclService.OSPolicyAssignmentRollout{}
			if _, ok := rRollout["disruptionBudget"]; ok {
				if rRolloutDisruptionBudget, ok := rRollout["disruptionBudget"].(map[string]interface{}); ok {
					r.Rollout.DisruptionBudget = &dclService.OSPolicyAssignmentRolloutDisruptionBudget{}
					if _, ok := rRolloutDisruptionBudget["fixed"]; ok {
						if i, ok := rRolloutDisruptionBudget["fixed"].(int64); ok {
							r.Rollout.DisruptionBudget.Fixed = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Rollout.DisruptionBudget.Fixed: expected int64")
						}
					}
					if _, ok := rRolloutDisruptionBudget["percent"]; ok {
						if i, ok := rRolloutDisruptionBudget["percent"].(int64); ok {
							r.Rollout.DisruptionBudget.Percent = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Rollout.DisruptionBudget.Percent: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Rollout.DisruptionBudget: expected map[string]interface{}")
				}
			}
			if _, ok := rRollout["minWaitDuration"]; ok {
				if s, ok := rRollout["minWaitDuration"].(string); ok {
					r.Rollout.MinWaitDuration = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Rollout.MinWaitDuration: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Rollout: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["rolloutState"]; ok {
		if s, ok := u.Object["rolloutState"].(string); ok {
			r.RolloutState = dclService.OSPolicyAssignmentRolloutStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.RolloutState: expected string")
		}
	}
	if _, ok := u.Object["skipAwaitRollout"]; ok {
		if b, ok := u.Object["skipAwaitRollout"].(bool); ok {
			r.SkipAwaitRollout = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.SkipAwaitRollout: expected bool")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
		}
	}
	return r, nil
}

func GetOSPolicyAssignment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOSPolicyAssignment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetOSPolicyAssignment(ctx, r)
	if err != nil {
		return nil, err
	}
	return OSPolicyAssignmentToUnstructured(r), nil
}

func ListOSPolicyAssignment(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListOSPolicyAssignment(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, OSPolicyAssignmentToUnstructured(r))
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

func ApplyOSPolicyAssignment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOSPolicyAssignment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToOSPolicyAssignment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyOSPolicyAssignment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return OSPolicyAssignmentToUnstructured(r), nil
}

func OSPolicyAssignmentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOSPolicyAssignment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToOSPolicyAssignment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyOSPolicyAssignment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteOSPolicyAssignment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToOSPolicyAssignment(u)
	if err != nil {
		return err
	}
	return c.DeleteOSPolicyAssignment(ctx, r)
}

func OSPolicyAssignmentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToOSPolicyAssignment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *OSPolicyAssignment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"osconfig",
		"OSPolicyAssignment",
		"alpha",
	}
}

func (r *OSPolicyAssignment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OSPolicyAssignment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OSPolicyAssignment) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *OSPolicyAssignment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OSPolicyAssignment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OSPolicyAssignment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *OSPolicyAssignment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetOSPolicyAssignment(ctx, config, resource)
}

func (r *OSPolicyAssignment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyOSPolicyAssignment(ctx, config, resource, opts...)
}

func (r *OSPolicyAssignment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return OSPolicyAssignmentHasDiff(ctx, config, resource, opts...)
}

func (r *OSPolicyAssignment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteOSPolicyAssignment(ctx, config, resource)
}

func (r *OSPolicyAssignment) ID(resource *unstructured.Resource) (string, error) {
	return OSPolicyAssignmentID(resource)
}

func init() {
	unstructured.Register(&OSPolicyAssignment{})
}
