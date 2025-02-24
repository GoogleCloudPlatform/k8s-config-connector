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
package gkehub

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type FeatureMembership struct{}

func FeatureMembershipToUnstructured(r *dclService.FeatureMembership) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "gkehub",
			Version: "alpha",
			Type:    "FeatureMembership",
		},
		Object: make(map[string]interface{}),
	}
	if r.Configmanagement != nil && r.Configmanagement != dclService.EmptyFeatureMembershipConfigmanagement {
		rConfigmanagement := make(map[string]interface{})
		if r.Configmanagement.Binauthz != nil && r.Configmanagement.Binauthz != dclService.EmptyFeatureMembershipConfigmanagementBinauthz {
			rConfigmanagementBinauthz := make(map[string]interface{})
			if r.Configmanagement.Binauthz.Enabled != nil {
				rConfigmanagementBinauthz["enabled"] = *r.Configmanagement.Binauthz.Enabled
			}
			rConfigmanagement["binauthz"] = rConfigmanagementBinauthz
		}
		if r.Configmanagement.ConfigSync != nil && r.Configmanagement.ConfigSync != dclService.EmptyFeatureMembershipConfigmanagementConfigSync {
			rConfigmanagementConfigSync := make(map[string]interface{})
			if r.Configmanagement.ConfigSync.Git != nil && r.Configmanagement.ConfigSync.Git != dclService.EmptyFeatureMembershipConfigmanagementConfigSyncGit {
				rConfigmanagementConfigSyncGit := make(map[string]interface{})
				if r.Configmanagement.ConfigSync.Git.GcpServiceAccountEmail != nil {
					rConfigmanagementConfigSyncGit["gcpServiceAccountEmail"] = *r.Configmanagement.ConfigSync.Git.GcpServiceAccountEmail
				}
				if r.Configmanagement.ConfigSync.Git.HttpsProxy != nil {
					rConfigmanagementConfigSyncGit["httpsProxy"] = *r.Configmanagement.ConfigSync.Git.HttpsProxy
				}
				if r.Configmanagement.ConfigSync.Git.PolicyDir != nil {
					rConfigmanagementConfigSyncGit["policyDir"] = *r.Configmanagement.ConfigSync.Git.PolicyDir
				}
				if r.Configmanagement.ConfigSync.Git.SecretType != nil {
					rConfigmanagementConfigSyncGit["secretType"] = *r.Configmanagement.ConfigSync.Git.SecretType
				}
				if r.Configmanagement.ConfigSync.Git.SyncBranch != nil {
					rConfigmanagementConfigSyncGit["syncBranch"] = *r.Configmanagement.ConfigSync.Git.SyncBranch
				}
				if r.Configmanagement.ConfigSync.Git.SyncRepo != nil {
					rConfigmanagementConfigSyncGit["syncRepo"] = *r.Configmanagement.ConfigSync.Git.SyncRepo
				}
				if r.Configmanagement.ConfigSync.Git.SyncRev != nil {
					rConfigmanagementConfigSyncGit["syncRev"] = *r.Configmanagement.ConfigSync.Git.SyncRev
				}
				if r.Configmanagement.ConfigSync.Git.SyncWaitSecs != nil {
					rConfigmanagementConfigSyncGit["syncWaitSecs"] = *r.Configmanagement.ConfigSync.Git.SyncWaitSecs
				}
				rConfigmanagementConfigSync["git"] = rConfigmanagementConfigSyncGit
			}
			if r.Configmanagement.ConfigSync.MetricsGcpServiceAccountEmail != nil {
				rConfigmanagementConfigSync["metricsGcpServiceAccountEmail"] = *r.Configmanagement.ConfigSync.MetricsGcpServiceAccountEmail
			}
			if r.Configmanagement.ConfigSync.Oci != nil && r.Configmanagement.ConfigSync.Oci != dclService.EmptyFeatureMembershipConfigmanagementConfigSyncOci {
				rConfigmanagementConfigSyncOci := make(map[string]interface{})
				if r.Configmanagement.ConfigSync.Oci.GcpServiceAccountEmail != nil {
					rConfigmanagementConfigSyncOci["gcpServiceAccountEmail"] = *r.Configmanagement.ConfigSync.Oci.GcpServiceAccountEmail
				}
				if r.Configmanagement.ConfigSync.Oci.PolicyDir != nil {
					rConfigmanagementConfigSyncOci["policyDir"] = *r.Configmanagement.ConfigSync.Oci.PolicyDir
				}
				if r.Configmanagement.ConfigSync.Oci.SecretType != nil {
					rConfigmanagementConfigSyncOci["secretType"] = *r.Configmanagement.ConfigSync.Oci.SecretType
				}
				if r.Configmanagement.ConfigSync.Oci.SyncRepo != nil {
					rConfigmanagementConfigSyncOci["syncRepo"] = *r.Configmanagement.ConfigSync.Oci.SyncRepo
				}
				if r.Configmanagement.ConfigSync.Oci.SyncWaitSecs != nil {
					rConfigmanagementConfigSyncOci["syncWaitSecs"] = *r.Configmanagement.ConfigSync.Oci.SyncWaitSecs
				}
				rConfigmanagementConfigSync["oci"] = rConfigmanagementConfigSyncOci
			}
			if r.Configmanagement.ConfigSync.PreventDrift != nil {
				rConfigmanagementConfigSync["preventDrift"] = *r.Configmanagement.ConfigSync.PreventDrift
			}
			if r.Configmanagement.ConfigSync.SourceFormat != nil {
				rConfigmanagementConfigSync["sourceFormat"] = *r.Configmanagement.ConfigSync.SourceFormat
			}
			rConfigmanagement["configSync"] = rConfigmanagementConfigSync
		}
		if r.Configmanagement.HierarchyController != nil && r.Configmanagement.HierarchyController != dclService.EmptyFeatureMembershipConfigmanagementHierarchyController {
			rConfigmanagementHierarchyController := make(map[string]interface{})
			if r.Configmanagement.HierarchyController.EnableHierarchicalResourceQuota != nil {
				rConfigmanagementHierarchyController["enableHierarchicalResourceQuota"] = *r.Configmanagement.HierarchyController.EnableHierarchicalResourceQuota
			}
			if r.Configmanagement.HierarchyController.EnablePodTreeLabels != nil {
				rConfigmanagementHierarchyController["enablePodTreeLabels"] = *r.Configmanagement.HierarchyController.EnablePodTreeLabels
			}
			if r.Configmanagement.HierarchyController.Enabled != nil {
				rConfigmanagementHierarchyController["enabled"] = *r.Configmanagement.HierarchyController.Enabled
			}
			rConfigmanagement["hierarchyController"] = rConfigmanagementHierarchyController
		}
		if r.Configmanagement.PolicyController != nil && r.Configmanagement.PolicyController != dclService.EmptyFeatureMembershipConfigmanagementPolicyController {
			rConfigmanagementPolicyController := make(map[string]interface{})
			if r.Configmanagement.PolicyController.AuditIntervalSeconds != nil {
				rConfigmanagementPolicyController["auditIntervalSeconds"] = *r.Configmanagement.PolicyController.AuditIntervalSeconds
			}
			if r.Configmanagement.PolicyController.Enabled != nil {
				rConfigmanagementPolicyController["enabled"] = *r.Configmanagement.PolicyController.Enabled
			}
			var rConfigmanagementPolicyControllerExemptableNamespaces []interface{}
			for _, rConfigmanagementPolicyControllerExemptableNamespacesVal := range r.Configmanagement.PolicyController.ExemptableNamespaces {
				rConfigmanagementPolicyControllerExemptableNamespaces = append(rConfigmanagementPolicyControllerExemptableNamespaces, rConfigmanagementPolicyControllerExemptableNamespacesVal)
			}
			rConfigmanagementPolicyController["exemptableNamespaces"] = rConfigmanagementPolicyControllerExemptableNamespaces
			if r.Configmanagement.PolicyController.LogDeniesEnabled != nil {
				rConfigmanagementPolicyController["logDeniesEnabled"] = *r.Configmanagement.PolicyController.LogDeniesEnabled
			}
			if r.Configmanagement.PolicyController.Monitoring != nil && r.Configmanagement.PolicyController.Monitoring != dclService.EmptyFeatureMembershipConfigmanagementPolicyControllerMonitoring {
				rConfigmanagementPolicyControllerMonitoring := make(map[string]interface{})
				var rConfigmanagementPolicyControllerMonitoringBackends []interface{}
				for _, rConfigmanagementPolicyControllerMonitoringBackendsVal := range r.Configmanagement.PolicyController.Monitoring.Backends {
					rConfigmanagementPolicyControllerMonitoringBackends = append(rConfigmanagementPolicyControllerMonitoringBackends, string(rConfigmanagementPolicyControllerMonitoringBackendsVal))
				}
				rConfigmanagementPolicyControllerMonitoring["backends"] = rConfigmanagementPolicyControllerMonitoringBackends
				rConfigmanagementPolicyController["monitoring"] = rConfigmanagementPolicyControllerMonitoring
			}
			if r.Configmanagement.PolicyController.MutationEnabled != nil {
				rConfigmanagementPolicyController["mutationEnabled"] = *r.Configmanagement.PolicyController.MutationEnabled
			}
			if r.Configmanagement.PolicyController.ReferentialRulesEnabled != nil {
				rConfigmanagementPolicyController["referentialRulesEnabled"] = *r.Configmanagement.PolicyController.ReferentialRulesEnabled
			}
			if r.Configmanagement.PolicyController.TemplateLibraryInstalled != nil {
				rConfigmanagementPolicyController["templateLibraryInstalled"] = *r.Configmanagement.PolicyController.TemplateLibraryInstalled
			}
			rConfigmanagement["policyController"] = rConfigmanagementPolicyController
		}
		if r.Configmanagement.Version != nil {
			rConfigmanagement["version"] = *r.Configmanagement.Version
		}
		u.Object["configmanagement"] = rConfigmanagement
	}
	if r.Feature != nil {
		u.Object["feature"] = *r.Feature
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Membership != nil {
		u.Object["membership"] = *r.Membership
	}
	if r.MembershipLocation != nil {
		u.Object["membershipLocation"] = *r.MembershipLocation
	}
	if r.Mesh != nil && r.Mesh != dclService.EmptyFeatureMembershipMesh {
		rMesh := make(map[string]interface{})
		if r.Mesh.ControlPlane != nil {
			rMesh["controlPlane"] = string(*r.Mesh.ControlPlane)
		}
		if r.Mesh.Management != nil {
			rMesh["management"] = string(*r.Mesh.Management)
		}
		u.Object["mesh"] = rMesh
	}
	if r.Policycontroller != nil && r.Policycontroller != dclService.EmptyFeatureMembershipPolicycontroller {
		rPolicycontroller := make(map[string]interface{})
		if r.Policycontroller.PolicyControllerHubConfig != nil && r.Policycontroller.PolicyControllerHubConfig != dclService.EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfig {
			rPolicycontrollerPolicyControllerHubConfig := make(map[string]interface{})
			if r.Policycontroller.PolicyControllerHubConfig.AuditIntervalSeconds != nil {
				rPolicycontrollerPolicyControllerHubConfig["auditIntervalSeconds"] = *r.Policycontroller.PolicyControllerHubConfig.AuditIntervalSeconds
			}
			if r.Policycontroller.PolicyControllerHubConfig.ConstraintViolationLimit != nil {
				rPolicycontrollerPolicyControllerHubConfig["constraintViolationLimit"] = *r.Policycontroller.PolicyControllerHubConfig.ConstraintViolationLimit
			}
			var rPolicycontrollerPolicyControllerHubConfigExemptableNamespaces []interface{}
			for _, rPolicycontrollerPolicyControllerHubConfigExemptableNamespacesVal := range r.Policycontroller.PolicyControllerHubConfig.ExemptableNamespaces {
				rPolicycontrollerPolicyControllerHubConfigExemptableNamespaces = append(rPolicycontrollerPolicyControllerHubConfigExemptableNamespaces, rPolicycontrollerPolicyControllerHubConfigExemptableNamespacesVal)
			}
			rPolicycontrollerPolicyControllerHubConfig["exemptableNamespaces"] = rPolicycontrollerPolicyControllerHubConfigExemptableNamespaces
			if r.Policycontroller.PolicyControllerHubConfig.InstallSpec != nil {
				rPolicycontrollerPolicyControllerHubConfig["installSpec"] = string(*r.Policycontroller.PolicyControllerHubConfig.InstallSpec)
			}
			if r.Policycontroller.PolicyControllerHubConfig.LogDeniesEnabled != nil {
				rPolicycontrollerPolicyControllerHubConfig["logDeniesEnabled"] = *r.Policycontroller.PolicyControllerHubConfig.LogDeniesEnabled
			}
			if r.Policycontroller.PolicyControllerHubConfig.Monitoring != nil && r.Policycontroller.PolicyControllerHubConfig.Monitoring != dclService.EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
				rPolicycontrollerPolicyControllerHubConfigMonitoring := make(map[string]interface{})
				var rPolicycontrollerPolicyControllerHubConfigMonitoringBackends []interface{}
				for _, rPolicycontrollerPolicyControllerHubConfigMonitoringBackendsVal := range r.Policycontroller.PolicyControllerHubConfig.Monitoring.Backends {
					rPolicycontrollerPolicyControllerHubConfigMonitoringBackends = append(rPolicycontrollerPolicyControllerHubConfigMonitoringBackends, string(rPolicycontrollerPolicyControllerHubConfigMonitoringBackendsVal))
				}
				rPolicycontrollerPolicyControllerHubConfigMonitoring["backends"] = rPolicycontrollerPolicyControllerHubConfigMonitoringBackends
				rPolicycontrollerPolicyControllerHubConfig["monitoring"] = rPolicycontrollerPolicyControllerHubConfigMonitoring
			}
			if r.Policycontroller.PolicyControllerHubConfig.MutationEnabled != nil {
				rPolicycontrollerPolicyControllerHubConfig["mutationEnabled"] = *r.Policycontroller.PolicyControllerHubConfig.MutationEnabled
			}
			if r.Policycontroller.PolicyControllerHubConfig.PolicyContent != nil && r.Policycontroller.PolicyControllerHubConfig.PolicyContent != dclService.EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
				rPolicycontrollerPolicyControllerHubConfigPolicyContent := make(map[string]interface{})
				if r.Policycontroller.PolicyControllerHubConfig.PolicyContent.TemplateLibrary != nil && r.Policycontroller.PolicyControllerHubConfig.PolicyContent.TemplateLibrary != dclService.EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
					rPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary := make(map[string]interface{})
					if r.Policycontroller.PolicyControllerHubConfig.PolicyContent.TemplateLibrary.Installation != nil {
						rPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary["installation"] = string(*r.Policycontroller.PolicyControllerHubConfig.PolicyContent.TemplateLibrary.Installation)
					}
					rPolicycontrollerPolicyControllerHubConfigPolicyContent["templateLibrary"] = rPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary
				}
				rPolicycontrollerPolicyControllerHubConfig["policyContent"] = rPolicycontrollerPolicyControllerHubConfigPolicyContent
			}
			if r.Policycontroller.PolicyControllerHubConfig.ReferentialRulesEnabled != nil {
				rPolicycontrollerPolicyControllerHubConfig["referentialRulesEnabled"] = *r.Policycontroller.PolicyControllerHubConfig.ReferentialRulesEnabled
			}
			rPolicycontroller["policyControllerHubConfig"] = rPolicycontrollerPolicyControllerHubConfig
		}
		if r.Policycontroller.Version != nil {
			rPolicycontroller["version"] = *r.Policycontroller.Version
		}
		u.Object["policycontroller"] = rPolicycontroller
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	return u
}

func UnstructuredToFeatureMembership(u *unstructured.Resource) (*dclService.FeatureMembership, error) {
	r := &dclService.FeatureMembership{}
	if _, ok := u.Object["configmanagement"]; ok {
		if rConfigmanagement, ok := u.Object["configmanagement"].(map[string]interface{}); ok {
			r.Configmanagement = &dclService.FeatureMembershipConfigmanagement{}
			if _, ok := rConfigmanagement["binauthz"]; ok {
				if rConfigmanagementBinauthz, ok := rConfigmanagement["binauthz"].(map[string]interface{}); ok {
					r.Configmanagement.Binauthz = &dclService.FeatureMembershipConfigmanagementBinauthz{}
					if _, ok := rConfigmanagementBinauthz["enabled"]; ok {
						if b, ok := rConfigmanagementBinauthz["enabled"].(bool); ok {
							r.Configmanagement.Binauthz.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.Binauthz.Enabled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Configmanagement.Binauthz: expected map[string]interface{}")
				}
			}
			if _, ok := rConfigmanagement["configSync"]; ok {
				if rConfigmanagementConfigSync, ok := rConfigmanagement["configSync"].(map[string]interface{}); ok {
					r.Configmanagement.ConfigSync = &dclService.FeatureMembershipConfigmanagementConfigSync{}
					if _, ok := rConfigmanagementConfigSync["git"]; ok {
						if rConfigmanagementConfigSyncGit, ok := rConfigmanagementConfigSync["git"].(map[string]interface{}); ok {
							r.Configmanagement.ConfigSync.Git = &dclService.FeatureMembershipConfigmanagementConfigSyncGit{}
							if _, ok := rConfigmanagementConfigSyncGit["gcpServiceAccountEmail"]; ok {
								if s, ok := rConfigmanagementConfigSyncGit["gcpServiceAccountEmail"].(string); ok {
									r.Configmanagement.ConfigSync.Git.GcpServiceAccountEmail = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git.GcpServiceAccountEmail: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncGit["httpsProxy"]; ok {
								if s, ok := rConfigmanagementConfigSyncGit["httpsProxy"].(string); ok {
									r.Configmanagement.ConfigSync.Git.HttpsProxy = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git.HttpsProxy: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncGit["policyDir"]; ok {
								if s, ok := rConfigmanagementConfigSyncGit["policyDir"].(string); ok {
									r.Configmanagement.ConfigSync.Git.PolicyDir = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git.PolicyDir: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncGit["secretType"]; ok {
								if s, ok := rConfigmanagementConfigSyncGit["secretType"].(string); ok {
									r.Configmanagement.ConfigSync.Git.SecretType = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git.SecretType: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncGit["syncBranch"]; ok {
								if s, ok := rConfigmanagementConfigSyncGit["syncBranch"].(string); ok {
									r.Configmanagement.ConfigSync.Git.SyncBranch = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git.SyncBranch: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncGit["syncRepo"]; ok {
								if s, ok := rConfigmanagementConfigSyncGit["syncRepo"].(string); ok {
									r.Configmanagement.ConfigSync.Git.SyncRepo = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git.SyncRepo: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncGit["syncRev"]; ok {
								if s, ok := rConfigmanagementConfigSyncGit["syncRev"].(string); ok {
									r.Configmanagement.ConfigSync.Git.SyncRev = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git.SyncRev: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncGit["syncWaitSecs"]; ok {
								if s, ok := rConfigmanagementConfigSyncGit["syncWaitSecs"].(string); ok {
									r.Configmanagement.ConfigSync.Git.SyncWaitSecs = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git.SyncWaitSecs: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Git: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigmanagementConfigSync["metricsGcpServiceAccountEmail"]; ok {
						if s, ok := rConfigmanagementConfigSync["metricsGcpServiceAccountEmail"].(string); ok {
							r.Configmanagement.ConfigSync.MetricsGcpServiceAccountEmail = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.ConfigSync.MetricsGcpServiceAccountEmail: expected string")
						}
					}
					if _, ok := rConfigmanagementConfigSync["oci"]; ok {
						if rConfigmanagementConfigSyncOci, ok := rConfigmanagementConfigSync["oci"].(map[string]interface{}); ok {
							r.Configmanagement.ConfigSync.Oci = &dclService.FeatureMembershipConfigmanagementConfigSyncOci{}
							if _, ok := rConfigmanagementConfigSyncOci["gcpServiceAccountEmail"]; ok {
								if s, ok := rConfigmanagementConfigSyncOci["gcpServiceAccountEmail"].(string); ok {
									r.Configmanagement.ConfigSync.Oci.GcpServiceAccountEmail = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Oci.GcpServiceAccountEmail: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncOci["policyDir"]; ok {
								if s, ok := rConfigmanagementConfigSyncOci["policyDir"].(string); ok {
									r.Configmanagement.ConfigSync.Oci.PolicyDir = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Oci.PolicyDir: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncOci["secretType"]; ok {
								if s, ok := rConfigmanagementConfigSyncOci["secretType"].(string); ok {
									r.Configmanagement.ConfigSync.Oci.SecretType = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Oci.SecretType: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncOci["syncRepo"]; ok {
								if s, ok := rConfigmanagementConfigSyncOci["syncRepo"].(string); ok {
									r.Configmanagement.ConfigSync.Oci.SyncRepo = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Oci.SyncRepo: expected string")
								}
							}
							if _, ok := rConfigmanagementConfigSyncOci["syncWaitSecs"]; ok {
								if s, ok := rConfigmanagementConfigSyncOci["syncWaitSecs"].(string); ok {
									r.Configmanagement.ConfigSync.Oci.SyncWaitSecs = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Oci.SyncWaitSecs: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Configmanagement.ConfigSync.Oci: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigmanagementConfigSync["preventDrift"]; ok {
						if b, ok := rConfigmanagementConfigSync["preventDrift"].(bool); ok {
							r.Configmanagement.ConfigSync.PreventDrift = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.ConfigSync.PreventDrift: expected bool")
						}
					}
					if _, ok := rConfigmanagementConfigSync["sourceFormat"]; ok {
						if s, ok := rConfigmanagementConfigSync["sourceFormat"].(string); ok {
							r.Configmanagement.ConfigSync.SourceFormat = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.ConfigSync.SourceFormat: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Configmanagement.ConfigSync: expected map[string]interface{}")
				}
			}
			if _, ok := rConfigmanagement["hierarchyController"]; ok {
				if rConfigmanagementHierarchyController, ok := rConfigmanagement["hierarchyController"].(map[string]interface{}); ok {
					r.Configmanagement.HierarchyController = &dclService.FeatureMembershipConfigmanagementHierarchyController{}
					if _, ok := rConfigmanagementHierarchyController["enableHierarchicalResourceQuota"]; ok {
						if b, ok := rConfigmanagementHierarchyController["enableHierarchicalResourceQuota"].(bool); ok {
							r.Configmanagement.HierarchyController.EnableHierarchicalResourceQuota = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.HierarchyController.EnableHierarchicalResourceQuota: expected bool")
						}
					}
					if _, ok := rConfigmanagementHierarchyController["enablePodTreeLabels"]; ok {
						if b, ok := rConfigmanagementHierarchyController["enablePodTreeLabels"].(bool); ok {
							r.Configmanagement.HierarchyController.EnablePodTreeLabels = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.HierarchyController.EnablePodTreeLabels: expected bool")
						}
					}
					if _, ok := rConfigmanagementHierarchyController["enabled"]; ok {
						if b, ok := rConfigmanagementHierarchyController["enabled"].(bool); ok {
							r.Configmanagement.HierarchyController.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.HierarchyController.Enabled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Configmanagement.HierarchyController: expected map[string]interface{}")
				}
			}
			if _, ok := rConfigmanagement["policyController"]; ok {
				if rConfigmanagementPolicyController, ok := rConfigmanagement["policyController"].(map[string]interface{}); ok {
					r.Configmanagement.PolicyController = &dclService.FeatureMembershipConfigmanagementPolicyController{}
					if _, ok := rConfigmanagementPolicyController["auditIntervalSeconds"]; ok {
						if s, ok := rConfigmanagementPolicyController["auditIntervalSeconds"].(string); ok {
							r.Configmanagement.PolicyController.AuditIntervalSeconds = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.PolicyController.AuditIntervalSeconds: expected string")
						}
					}
					if _, ok := rConfigmanagementPolicyController["enabled"]; ok {
						if b, ok := rConfigmanagementPolicyController["enabled"].(bool); ok {
							r.Configmanagement.PolicyController.Enabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.PolicyController.Enabled: expected bool")
						}
					}
					if _, ok := rConfigmanagementPolicyController["exemptableNamespaces"]; ok {
						if s, ok := rConfigmanagementPolicyController["exemptableNamespaces"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Configmanagement.PolicyController.ExemptableNamespaces = append(r.Configmanagement.PolicyController.ExemptableNamespaces, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Configmanagement.PolicyController.ExemptableNamespaces: expected []interface{}")
						}
					}
					if _, ok := rConfigmanagementPolicyController["logDeniesEnabled"]; ok {
						if b, ok := rConfigmanagementPolicyController["logDeniesEnabled"].(bool); ok {
							r.Configmanagement.PolicyController.LogDeniesEnabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.PolicyController.LogDeniesEnabled: expected bool")
						}
					}
					if _, ok := rConfigmanagementPolicyController["monitoring"]; ok {
						if rConfigmanagementPolicyControllerMonitoring, ok := rConfigmanagementPolicyController["monitoring"].(map[string]interface{}); ok {
							r.Configmanagement.PolicyController.Monitoring = &dclService.FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
							if _, ok := rConfigmanagementPolicyControllerMonitoring["backends"]; ok {
								if s, ok := rConfigmanagementPolicyControllerMonitoring["backends"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Configmanagement.PolicyController.Monitoring.Backends = append(r.Configmanagement.PolicyController.Monitoring.Backends, dclService.FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(strval))
										}
									}
								} else {
									return nil, fmt.Errorf("r.Configmanagement.PolicyController.Monitoring.Backends: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Configmanagement.PolicyController.Monitoring: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigmanagementPolicyController["mutationEnabled"]; ok {
						if b, ok := rConfigmanagementPolicyController["mutationEnabled"].(bool); ok {
							r.Configmanagement.PolicyController.MutationEnabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.PolicyController.MutationEnabled: expected bool")
						}
					}
					if _, ok := rConfigmanagementPolicyController["referentialRulesEnabled"]; ok {
						if b, ok := rConfigmanagementPolicyController["referentialRulesEnabled"].(bool); ok {
							r.Configmanagement.PolicyController.ReferentialRulesEnabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.PolicyController.ReferentialRulesEnabled: expected bool")
						}
					}
					if _, ok := rConfigmanagementPolicyController["templateLibraryInstalled"]; ok {
						if b, ok := rConfigmanagementPolicyController["templateLibraryInstalled"].(bool); ok {
							r.Configmanagement.PolicyController.TemplateLibraryInstalled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Configmanagement.PolicyController.TemplateLibraryInstalled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Configmanagement.PolicyController: expected map[string]interface{}")
				}
			}
			if _, ok := rConfigmanagement["version"]; ok {
				if s, ok := rConfigmanagement["version"].(string); ok {
					r.Configmanagement.Version = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Configmanagement.Version: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Configmanagement: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["feature"]; ok {
		if s, ok := u.Object["feature"].(string); ok {
			r.Feature = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Feature: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["membership"]; ok {
		if s, ok := u.Object["membership"].(string); ok {
			r.Membership = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Membership: expected string")
		}
	}
	if _, ok := u.Object["membershipLocation"]; ok {
		if s, ok := u.Object["membershipLocation"].(string); ok {
			r.MembershipLocation = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.MembershipLocation: expected string")
		}
	}
	if _, ok := u.Object["mesh"]; ok {
		if rMesh, ok := u.Object["mesh"].(map[string]interface{}); ok {
			r.Mesh = &dclService.FeatureMembershipMesh{}
			if _, ok := rMesh["controlPlane"]; ok {
				if s, ok := rMesh["controlPlane"].(string); ok {
					r.Mesh.ControlPlane = dclService.FeatureMembershipMeshControlPlaneEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Mesh.ControlPlane: expected string")
				}
			}
			if _, ok := rMesh["management"]; ok {
				if s, ok := rMesh["management"].(string); ok {
					r.Mesh.Management = dclService.FeatureMembershipMeshManagementEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Mesh.Management: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Mesh: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["policycontroller"]; ok {
		if rPolicycontroller, ok := u.Object["policycontroller"].(map[string]interface{}); ok {
			r.Policycontroller = &dclService.FeatureMembershipPolicycontroller{}
			if _, ok := rPolicycontroller["policyControllerHubConfig"]; ok {
				if rPolicycontrollerPolicyControllerHubConfig, ok := rPolicycontroller["policyControllerHubConfig"].(map[string]interface{}); ok {
					r.Policycontroller.PolicyControllerHubConfig = &dclService.FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["auditIntervalSeconds"]; ok {
						if i, ok := rPolicycontrollerPolicyControllerHubConfig["auditIntervalSeconds"].(int64); ok {
							r.Policycontroller.PolicyControllerHubConfig.AuditIntervalSeconds = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.AuditIntervalSeconds: expected int64")
						}
					}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["constraintViolationLimit"]; ok {
						if i, ok := rPolicycontrollerPolicyControllerHubConfig["constraintViolationLimit"].(int64); ok {
							r.Policycontroller.PolicyControllerHubConfig.ConstraintViolationLimit = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.ConstraintViolationLimit: expected int64")
						}
					}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["exemptableNamespaces"]; ok {
						if s, ok := rPolicycontrollerPolicyControllerHubConfig["exemptableNamespaces"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Policycontroller.PolicyControllerHubConfig.ExemptableNamespaces = append(r.Policycontroller.PolicyControllerHubConfig.ExemptableNamespaces, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.ExemptableNamespaces: expected []interface{}")
						}
					}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["installSpec"]; ok {
						if s, ok := rPolicycontrollerPolicyControllerHubConfig["installSpec"].(string); ok {
							r.Policycontroller.PolicyControllerHubConfig.InstallSpec = dclService.FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.InstallSpec: expected string")
						}
					}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["logDeniesEnabled"]; ok {
						if b, ok := rPolicycontrollerPolicyControllerHubConfig["logDeniesEnabled"].(bool); ok {
							r.Policycontroller.PolicyControllerHubConfig.LogDeniesEnabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.LogDeniesEnabled: expected bool")
						}
					}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["monitoring"]; ok {
						if rPolicycontrollerPolicyControllerHubConfigMonitoring, ok := rPolicycontrollerPolicyControllerHubConfig["monitoring"].(map[string]interface{}); ok {
							r.Policycontroller.PolicyControllerHubConfig.Monitoring = &dclService.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
							if _, ok := rPolicycontrollerPolicyControllerHubConfigMonitoring["backends"]; ok {
								if s, ok := rPolicycontrollerPolicyControllerHubConfigMonitoring["backends"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Policycontroller.PolicyControllerHubConfig.Monitoring.Backends = append(r.Policycontroller.PolicyControllerHubConfig.Monitoring.Backends, dclService.FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(strval))
										}
									}
								} else {
									return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.Monitoring.Backends: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.Monitoring: expected map[string]interface{}")
						}
					}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["mutationEnabled"]; ok {
						if b, ok := rPolicycontrollerPolicyControllerHubConfig["mutationEnabled"].(bool); ok {
							r.Policycontroller.PolicyControllerHubConfig.MutationEnabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.MutationEnabled: expected bool")
						}
					}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["policyContent"]; ok {
						if rPolicycontrollerPolicyControllerHubConfigPolicyContent, ok := rPolicycontrollerPolicyControllerHubConfig["policyContent"].(map[string]interface{}); ok {
							r.Policycontroller.PolicyControllerHubConfig.PolicyContent = &dclService.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
							if _, ok := rPolicycontrollerPolicyControllerHubConfigPolicyContent["templateLibrary"]; ok {
								if rPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, ok := rPolicycontrollerPolicyControllerHubConfigPolicyContent["templateLibrary"].(map[string]interface{}); ok {
									r.Policycontroller.PolicyControllerHubConfig.PolicyContent.TemplateLibrary = &dclService.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
									if _, ok := rPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary["installation"]; ok {
										if s, ok := rPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary["installation"].(string); ok {
											r.Policycontroller.PolicyControllerHubConfig.PolicyContent.TemplateLibrary.Installation = dclService.FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.PolicyContent.TemplateLibrary.Installation: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.PolicyContent.TemplateLibrary: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.PolicyContent: expected map[string]interface{}")
						}
					}
					if _, ok := rPolicycontrollerPolicyControllerHubConfig["referentialRulesEnabled"]; ok {
						if b, ok := rPolicycontrollerPolicyControllerHubConfig["referentialRulesEnabled"].(bool); ok {
							r.Policycontroller.PolicyControllerHubConfig.ReferentialRulesEnabled = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig.ReferentialRulesEnabled: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Policycontroller.PolicyControllerHubConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rPolicycontroller["version"]; ok {
				if s, ok := rPolicycontroller["version"].(string); ok {
					r.Policycontroller.Version = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Policycontroller.Version: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Policycontroller: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	return r, nil
}

func GetFeatureMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFeatureMembership(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetFeatureMembership(ctx, r)
	if err != nil {
		return nil, err
	}
	return FeatureMembershipToUnstructured(r), nil
}

func ListFeatureMembership(ctx context.Context, config *dcl.Config, project string, location string, feature string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListFeatureMembership(ctx, project, location, feature)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, FeatureMembershipToUnstructured(r))
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

func ApplyFeatureMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFeatureMembership(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFeatureMembership(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyFeatureMembership(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return FeatureMembershipToUnstructured(r), nil
}

func FeatureMembershipHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFeatureMembership(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFeatureMembership(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyFeatureMembership(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteFeatureMembership(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFeatureMembership(u)
	if err != nil {
		return err
	}
	return c.DeleteFeatureMembership(ctx, r)
}

func FeatureMembershipID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToFeatureMembership(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *FeatureMembership) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"gkehub",
		"FeatureMembership",
		"alpha",
	}
}

func (r *FeatureMembership) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FeatureMembership) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FeatureMembership) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *FeatureMembership) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FeatureMembership) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FeatureMembership) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FeatureMembership) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetFeatureMembership(ctx, config, resource)
}

func (r *FeatureMembership) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyFeatureMembership(ctx, config, resource, opts...)
}

func (r *FeatureMembership) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return FeatureMembershipHasDiff(ctx, config, resource, opts...)
}

func (r *FeatureMembership) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteFeatureMembership(ctx, config, resource)
}

func (r *FeatureMembership) ID(resource *unstructured.Resource) (string, error) {
	return FeatureMembershipID(resource)
}

func init() {
	unstructured.Register(&FeatureMembership{})
}
