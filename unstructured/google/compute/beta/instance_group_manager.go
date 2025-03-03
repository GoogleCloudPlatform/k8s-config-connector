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
package compute

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type InstanceGroupManager struct{}

func InstanceGroupManagerToUnstructured(r *dclService.InstanceGroupManager) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "beta",
			Type:    "InstanceGroupManager",
		},
		Object: make(map[string]interface{}),
	}
	var rAutoHealingPolicies []interface{}
	for _, rAutoHealingPoliciesVal := range r.AutoHealingPolicies {
		rAutoHealingPoliciesObject := make(map[string]interface{})
		if rAutoHealingPoliciesVal.HealthCheck != nil {
			rAutoHealingPoliciesObject["healthCheck"] = *rAutoHealingPoliciesVal.HealthCheck
		}
		if rAutoHealingPoliciesVal.InitialDelaySec != nil {
			rAutoHealingPoliciesObject["initialDelaySec"] = *rAutoHealingPoliciesVal.InitialDelaySec
		}
		rAutoHealingPolicies = append(rAutoHealingPolicies, rAutoHealingPoliciesObject)
	}
	u.Object["autoHealingPolicies"] = rAutoHealingPolicies
	if r.BaseInstanceName != nil {
		u.Object["baseInstanceName"] = *r.BaseInstanceName
	}
	if r.CreationTimestamp != nil {
		u.Object["creationTimestamp"] = *r.CreationTimestamp
	}
	if r.CurrentActions != nil && r.CurrentActions != dclService.EmptyInstanceGroupManagerCurrentActions {
		rCurrentActions := make(map[string]interface{})
		if r.CurrentActions.Abandoning != nil {
			rCurrentActions["abandoning"] = *r.CurrentActions.Abandoning
		}
		if r.CurrentActions.Creating != nil {
			rCurrentActions["creating"] = *r.CurrentActions.Creating
		}
		if r.CurrentActions.CreatingWithoutRetries != nil {
			rCurrentActions["creatingWithoutRetries"] = *r.CurrentActions.CreatingWithoutRetries
		}
		if r.CurrentActions.Deleting != nil {
			rCurrentActions["deleting"] = *r.CurrentActions.Deleting
		}
		if r.CurrentActions.None != nil {
			rCurrentActions["none"] = *r.CurrentActions.None
		}
		if r.CurrentActions.Recreating != nil {
			rCurrentActions["recreating"] = *r.CurrentActions.Recreating
		}
		if r.CurrentActions.Refreshing != nil {
			rCurrentActions["refreshing"] = *r.CurrentActions.Refreshing
		}
		if r.CurrentActions.Restarting != nil {
			rCurrentActions["restarting"] = *r.CurrentActions.Restarting
		}
		if r.CurrentActions.Verifying != nil {
			rCurrentActions["verifying"] = *r.CurrentActions.Verifying
		}
		u.Object["currentActions"] = rCurrentActions
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DistributionPolicy != nil && r.DistributionPolicy != dclService.EmptyInstanceGroupManagerDistributionPolicy {
		rDistributionPolicy := make(map[string]interface{})
		if r.DistributionPolicy.TargetShape != nil {
			rDistributionPolicy["targetShape"] = string(*r.DistributionPolicy.TargetShape)
		}
		var rDistributionPolicyZones []interface{}
		for _, rDistributionPolicyZonesVal := range r.DistributionPolicy.Zones {
			rDistributionPolicyZonesObject := make(map[string]interface{})
			if rDistributionPolicyZonesVal.Zone != nil {
				rDistributionPolicyZonesObject["zone"] = *rDistributionPolicyZonesVal.Zone
			}
			rDistributionPolicyZones = append(rDistributionPolicyZones, rDistributionPolicyZonesObject)
		}
		rDistributionPolicy["zones"] = rDistributionPolicyZones
		u.Object["distributionPolicy"] = rDistributionPolicy
	}
	if r.FailoverAction != nil {
		u.Object["failoverAction"] = string(*r.FailoverAction)
	}
	if r.Fingerprint != nil {
		u.Object["fingerprint"] = *r.Fingerprint
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.InstanceGroup != nil {
		u.Object["instanceGroup"] = *r.InstanceGroup
	}
	if r.InstanceTemplate != nil {
		u.Object["instanceTemplate"] = *r.InstanceTemplate
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	var rNamedPorts []interface{}
	for _, rNamedPortsVal := range r.NamedPorts {
		rNamedPortsObject := make(map[string]interface{})
		if rNamedPortsVal.Name != nil {
			rNamedPortsObject["name"] = *rNamedPortsVal.Name
		}
		if rNamedPortsVal.Port != nil {
			rNamedPortsObject["port"] = *rNamedPortsVal.Port
		}
		rNamedPorts = append(rNamedPorts, rNamedPortsObject)
	}
	u.Object["namedPorts"] = rNamedPorts
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Region != nil {
		u.Object["region"] = *r.Region
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.ServiceAccount != nil {
		u.Object["serviceAccount"] = *r.ServiceAccount
	}
	if r.StatefulPolicy != nil && r.StatefulPolicy != dclService.EmptyInstanceGroupManagerStatefulPolicy {
		rStatefulPolicy := make(map[string]interface{})
		if r.StatefulPolicy.PreservedState != nil && r.StatefulPolicy.PreservedState != dclService.EmptyInstanceGroupManagerStatefulPolicyPreservedState {
			rStatefulPolicyPreservedState := make(map[string]interface{})
			if r.StatefulPolicy.PreservedState.Disks != nil {
				rStatefulPolicyPreservedStateDisks := make(map[string]interface{})
				for k, v := range r.StatefulPolicy.PreservedState.Disks {
					rStatefulPolicyPreservedStateDisksMap := make(map[string]interface{})
					if v.AutoDelete != nil {
						rStatefulPolicyPreservedStateDisksMap["autoDelete"] = string(*v.AutoDelete)
					}
					rStatefulPolicyPreservedStateDisks[k] = rStatefulPolicyPreservedStateDisksMap
				}
				rStatefulPolicyPreservedState["disks"] = rStatefulPolicyPreservedStateDisks
			}
			if r.StatefulPolicy.PreservedState.ExternalIps != nil {
				rStatefulPolicyPreservedStateExternalIps := make(map[string]interface{})
				for k, v := range r.StatefulPolicy.PreservedState.ExternalIps {
					rStatefulPolicyPreservedStateExternalIpsMap := make(map[string]interface{})
					if v.AutoDelete != nil {
						rStatefulPolicyPreservedStateExternalIpsMap["autoDelete"] = string(*v.AutoDelete)
					}
					rStatefulPolicyPreservedStateExternalIps[k] = rStatefulPolicyPreservedStateExternalIpsMap
				}
				rStatefulPolicyPreservedState["externalIps"] = rStatefulPolicyPreservedStateExternalIps
			}
			if r.StatefulPolicy.PreservedState.InternalIps != nil {
				rStatefulPolicyPreservedStateInternalIps := make(map[string]interface{})
				for k, v := range r.StatefulPolicy.PreservedState.InternalIps {
					rStatefulPolicyPreservedStateInternalIpsMap := make(map[string]interface{})
					if v.AutoDelete != nil {
						rStatefulPolicyPreservedStateInternalIpsMap["autoDelete"] = string(*v.AutoDelete)
					}
					rStatefulPolicyPreservedStateInternalIps[k] = rStatefulPolicyPreservedStateInternalIpsMap
				}
				rStatefulPolicyPreservedState["internalIps"] = rStatefulPolicyPreservedStateInternalIps
			}
			rStatefulPolicy["preservedState"] = rStatefulPolicyPreservedState
		}
		u.Object["statefulPolicy"] = rStatefulPolicy
	}
	if r.Status != nil && r.Status != dclService.EmptyInstanceGroupManagerStatus {
		rStatus := make(map[string]interface{})
		if r.Status.Autoscaler != nil {
			rStatus["autoscaler"] = *r.Status.Autoscaler
		}
		if r.Status.IsStable != nil {
			rStatus["isStable"] = *r.Status.IsStable
		}
		if r.Status.Stateful != nil && r.Status.Stateful != dclService.EmptyInstanceGroupManagerStatusStateful {
			rStatusStateful := make(map[string]interface{})
			if r.Status.Stateful.HasStatefulConfig != nil {
				rStatusStateful["hasStatefulConfig"] = *r.Status.Stateful.HasStatefulConfig
			}
			if r.Status.Stateful.IsStateful != nil {
				rStatusStateful["isStateful"] = *r.Status.Stateful.IsStateful
			}
			if r.Status.Stateful.PerInstanceConfigs != nil && r.Status.Stateful.PerInstanceConfigs != dclService.EmptyInstanceGroupManagerStatusStatefulPerInstanceConfigs {
				rStatusStatefulPerInstanceConfigs := make(map[string]interface{})
				if r.Status.Stateful.PerInstanceConfigs.AllEffective != nil {
					rStatusStatefulPerInstanceConfigs["allEffective"] = *r.Status.Stateful.PerInstanceConfigs.AllEffective
				}
				rStatusStateful["perInstanceConfigs"] = rStatusStatefulPerInstanceConfigs
			}
			rStatus["stateful"] = rStatusStateful
		}
		if r.Status.VersionTarget != nil && r.Status.VersionTarget != dclService.EmptyInstanceGroupManagerStatusVersionTarget {
			rStatusVersionTarget := make(map[string]interface{})
			if r.Status.VersionTarget.IsReached != nil {
				rStatusVersionTarget["isReached"] = *r.Status.VersionTarget.IsReached
			}
			rStatus["versionTarget"] = rStatusVersionTarget
		}
		u.Object["status"] = rStatus
	}
	var rTargetPools []interface{}
	for _, rTargetPoolsVal := range r.TargetPools {
		rTargetPools = append(rTargetPools, rTargetPoolsVal)
	}
	u.Object["targetPools"] = rTargetPools
	if r.TargetSize != nil {
		u.Object["targetSize"] = *r.TargetSize
	}
	if r.UpdatePolicy != nil && r.UpdatePolicy != dclService.EmptyInstanceGroupManagerUpdatePolicy {
		rUpdatePolicy := make(map[string]interface{})
		if r.UpdatePolicy.InstanceRedistributionType != nil {
			rUpdatePolicy["instanceRedistributionType"] = string(*r.UpdatePolicy.InstanceRedistributionType)
		}
		if r.UpdatePolicy.MaxSurge != nil && r.UpdatePolicy.MaxSurge != dclService.EmptyInstanceGroupManagerUpdatePolicyMaxSurge {
			rUpdatePolicyMaxSurge := make(map[string]interface{})
			if r.UpdatePolicy.MaxSurge.Calculated != nil {
				rUpdatePolicyMaxSurge["calculated"] = *r.UpdatePolicy.MaxSurge.Calculated
			}
			if r.UpdatePolicy.MaxSurge.Fixed != nil {
				rUpdatePolicyMaxSurge["fixed"] = *r.UpdatePolicy.MaxSurge.Fixed
			}
			if r.UpdatePolicy.MaxSurge.Percent != nil {
				rUpdatePolicyMaxSurge["percent"] = *r.UpdatePolicy.MaxSurge.Percent
			}
			rUpdatePolicy["maxSurge"] = rUpdatePolicyMaxSurge
		}
		if r.UpdatePolicy.MaxUnavailable != nil && r.UpdatePolicy.MaxUnavailable != dclService.EmptyInstanceGroupManagerUpdatePolicyMaxUnavailable {
			rUpdatePolicyMaxUnavailable := make(map[string]interface{})
			if r.UpdatePolicy.MaxUnavailable.Calculated != nil {
				rUpdatePolicyMaxUnavailable["calculated"] = *r.UpdatePolicy.MaxUnavailable.Calculated
			}
			if r.UpdatePolicy.MaxUnavailable.Fixed != nil {
				rUpdatePolicyMaxUnavailable["fixed"] = *r.UpdatePolicy.MaxUnavailable.Fixed
			}
			if r.UpdatePolicy.MaxUnavailable.Percent != nil {
				rUpdatePolicyMaxUnavailable["percent"] = *r.UpdatePolicy.MaxUnavailable.Percent
			}
			rUpdatePolicy["maxUnavailable"] = rUpdatePolicyMaxUnavailable
		}
		if r.UpdatePolicy.MinReadySec != nil {
			rUpdatePolicy["minReadySec"] = *r.UpdatePolicy.MinReadySec
		}
		if r.UpdatePolicy.MinimalAction != nil {
			rUpdatePolicy["minimalAction"] = string(*r.UpdatePolicy.MinimalAction)
		}
		if r.UpdatePolicy.MostDisruptiveAllowedAction != nil {
			rUpdatePolicy["mostDisruptiveAllowedAction"] = string(*r.UpdatePolicy.MostDisruptiveAllowedAction)
		}
		if r.UpdatePolicy.ReplacementMethod != nil {
			rUpdatePolicy["replacementMethod"] = string(*r.UpdatePolicy.ReplacementMethod)
		}
		if r.UpdatePolicy.Type != nil {
			rUpdatePolicy["type"] = string(*r.UpdatePolicy.Type)
		}
		u.Object["updatePolicy"] = rUpdatePolicy
	}
	var rVersions []interface{}
	for _, rVersionsVal := range r.Versions {
		rVersionsObject := make(map[string]interface{})
		if rVersionsVal.InstanceTemplate != nil {
			rVersionsObject["instanceTemplate"] = *rVersionsVal.InstanceTemplate
		}
		if rVersionsVal.Name != nil {
			rVersionsObject["name"] = *rVersionsVal.Name
		}
		if rVersionsVal.TargetSize != nil && rVersionsVal.TargetSize != dclService.EmptyInstanceGroupManagerVersionsTargetSize {
			rVersionsValTargetSize := make(map[string]interface{})
			if rVersionsVal.TargetSize.Calculated != nil {
				rVersionsValTargetSize["calculated"] = *rVersionsVal.TargetSize.Calculated
			}
			if rVersionsVal.TargetSize.Fixed != nil {
				rVersionsValTargetSize["fixed"] = *rVersionsVal.TargetSize.Fixed
			}
			if rVersionsVal.TargetSize.Percent != nil {
				rVersionsValTargetSize["percent"] = *rVersionsVal.TargetSize.Percent
			}
			rVersionsObject["targetSize"] = rVersionsValTargetSize
		}
		rVersions = append(rVersions, rVersionsObject)
	}
	u.Object["versions"] = rVersions
	if r.Zone != nil {
		u.Object["zone"] = *r.Zone
	}
	return u
}

func UnstructuredToInstanceGroupManager(u *unstructured.Resource) (*dclService.InstanceGroupManager, error) {
	r := &dclService.InstanceGroupManager{}
	if _, ok := u.Object["autoHealingPolicies"]; ok {
		if s, ok := u.Object["autoHealingPolicies"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rAutoHealingPolicies dclService.InstanceGroupManagerAutoHealingPolicies
					if _, ok := objval["healthCheck"]; ok {
						if s, ok := objval["healthCheck"].(string); ok {
							rAutoHealingPolicies.HealthCheck = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAutoHealingPolicies.HealthCheck: expected string")
						}
					}
					if _, ok := objval["initialDelaySec"]; ok {
						if i, ok := objval["initialDelaySec"].(int64); ok {
							rAutoHealingPolicies.InitialDelaySec = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rAutoHealingPolicies.InitialDelaySec: expected int64")
						}
					}
					r.AutoHealingPolicies = append(r.AutoHealingPolicies, rAutoHealingPolicies)
				}
			}
		} else {
			return nil, fmt.Errorf("r.AutoHealingPolicies: expected []interface{}")
		}
	}
	if _, ok := u.Object["baseInstanceName"]; ok {
		if s, ok := u.Object["baseInstanceName"].(string); ok {
			r.BaseInstanceName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.BaseInstanceName: expected string")
		}
	}
	if _, ok := u.Object["creationTimestamp"]; ok {
		if s, ok := u.Object["creationTimestamp"].(string); ok {
			r.CreationTimestamp = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreationTimestamp: expected string")
		}
	}
	if _, ok := u.Object["currentActions"]; ok {
		if rCurrentActions, ok := u.Object["currentActions"].(map[string]interface{}); ok {
			r.CurrentActions = &dclService.InstanceGroupManagerCurrentActions{}
			if _, ok := rCurrentActions["abandoning"]; ok {
				if i, ok := rCurrentActions["abandoning"].(int64); ok {
					r.CurrentActions.Abandoning = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.Abandoning: expected int64")
				}
			}
			if _, ok := rCurrentActions["creating"]; ok {
				if i, ok := rCurrentActions["creating"].(int64); ok {
					r.CurrentActions.Creating = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.Creating: expected int64")
				}
			}
			if _, ok := rCurrentActions["creatingWithoutRetries"]; ok {
				if i, ok := rCurrentActions["creatingWithoutRetries"].(int64); ok {
					r.CurrentActions.CreatingWithoutRetries = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.CreatingWithoutRetries: expected int64")
				}
			}
			if _, ok := rCurrentActions["deleting"]; ok {
				if i, ok := rCurrentActions["deleting"].(int64); ok {
					r.CurrentActions.Deleting = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.Deleting: expected int64")
				}
			}
			if _, ok := rCurrentActions["none"]; ok {
				if i, ok := rCurrentActions["none"].(int64); ok {
					r.CurrentActions.None = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.None: expected int64")
				}
			}
			if _, ok := rCurrentActions["recreating"]; ok {
				if i, ok := rCurrentActions["recreating"].(int64); ok {
					r.CurrentActions.Recreating = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.Recreating: expected int64")
				}
			}
			if _, ok := rCurrentActions["refreshing"]; ok {
				if i, ok := rCurrentActions["refreshing"].(int64); ok {
					r.CurrentActions.Refreshing = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.Refreshing: expected int64")
				}
			}
			if _, ok := rCurrentActions["restarting"]; ok {
				if i, ok := rCurrentActions["restarting"].(int64); ok {
					r.CurrentActions.Restarting = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.Restarting: expected int64")
				}
			}
			if _, ok := rCurrentActions["verifying"]; ok {
				if i, ok := rCurrentActions["verifying"].(int64); ok {
					r.CurrentActions.Verifying = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.CurrentActions.Verifying: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.CurrentActions: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["distributionPolicy"]; ok {
		if rDistributionPolicy, ok := u.Object["distributionPolicy"].(map[string]interface{}); ok {
			r.DistributionPolicy = &dclService.InstanceGroupManagerDistributionPolicy{}
			if _, ok := rDistributionPolicy["targetShape"]; ok {
				if s, ok := rDistributionPolicy["targetShape"].(string); ok {
					r.DistributionPolicy.TargetShape = dclService.InstanceGroupManagerDistributionPolicyTargetShapeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.DistributionPolicy.TargetShape: expected string")
				}
			}
			if _, ok := rDistributionPolicy["zones"]; ok {
				if s, ok := rDistributionPolicy["zones"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rDistributionPolicyZones dclService.InstanceGroupManagerDistributionPolicyZones
							if _, ok := objval["zone"]; ok {
								if s, ok := objval["zone"].(string); ok {
									rDistributionPolicyZones.Zone = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rDistributionPolicyZones.Zone: expected string")
								}
							}
							r.DistributionPolicy.Zones = append(r.DistributionPolicy.Zones, rDistributionPolicyZones)
						}
					}
				} else {
					return nil, fmt.Errorf("r.DistributionPolicy.Zones: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.DistributionPolicy: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["failoverAction"]; ok {
		if s, ok := u.Object["failoverAction"].(string); ok {
			r.FailoverAction = dclService.InstanceGroupManagerFailoverActionEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.FailoverAction: expected string")
		}
	}
	if _, ok := u.Object["fingerprint"]; ok {
		if s, ok := u.Object["fingerprint"].(string); ok {
			r.Fingerprint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Fingerprint: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if i, ok := u.Object["id"].(int64); ok {
			r.Id = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Id: expected int64")
		}
	}
	if _, ok := u.Object["instanceGroup"]; ok {
		if s, ok := u.Object["instanceGroup"].(string); ok {
			r.InstanceGroup = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.InstanceGroup: expected string")
		}
	}
	if _, ok := u.Object["instanceTemplate"]; ok {
		if s, ok := u.Object["instanceTemplate"].(string); ok {
			r.InstanceTemplate = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.InstanceTemplate: expected string")
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
	if _, ok := u.Object["namedPorts"]; ok {
		if s, ok := u.Object["namedPorts"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rNamedPorts dclService.InstanceGroupManagerNamedPorts
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rNamedPorts.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rNamedPorts.Name: expected string")
						}
					}
					if _, ok := objval["port"]; ok {
						if i, ok := objval["port"].(int64); ok {
							rNamedPorts.Port = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rNamedPorts.Port: expected int64")
						}
					}
					r.NamedPorts = append(r.NamedPorts, rNamedPorts)
				}
			}
		} else {
			return nil, fmt.Errorf("r.NamedPorts: expected []interface{}")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["region"]; ok {
		if s, ok := u.Object["region"].(string); ok {
			r.Region = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Region: expected string")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["serviceAccount"]; ok {
		if s, ok := u.Object["serviceAccount"].(string); ok {
			r.ServiceAccount = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ServiceAccount: expected string")
		}
	}
	if _, ok := u.Object["statefulPolicy"]; ok {
		if rStatefulPolicy, ok := u.Object["statefulPolicy"].(map[string]interface{}); ok {
			r.StatefulPolicy = &dclService.InstanceGroupManagerStatefulPolicy{}
			if _, ok := rStatefulPolicy["preservedState"]; ok {
				if rStatefulPolicyPreservedState, ok := rStatefulPolicy["preservedState"].(map[string]interface{}); ok {
					r.StatefulPolicy.PreservedState = &dclService.InstanceGroupManagerStatefulPolicyPreservedState{}
					if _, ok := rStatefulPolicyPreservedState["disks"]; ok {
						if rStatefulPolicyPreservedStateDisks, ok := rStatefulPolicyPreservedState["disks"].(map[string]interface{}); ok {
							m := make(map[string]dclService.InstanceGroupManagerStatefulPolicyPreservedStateDisks)
							for k, v := range rStatefulPolicyPreservedStateDisks {
								if objval, ok := v.(map[string]interface{}); ok {
									var rStatefulPolicyPreservedStateDisksObj dclService.InstanceGroupManagerStatefulPolicyPreservedStateDisks
									if _, ok := objval["autoDelete"]; ok {
										if s, ok := objval["autoDelete"].(string); ok {
											rStatefulPolicyPreservedStateDisksObj.AutoDelete = dclService.InstanceGroupManagerStatefulPolicyPreservedStateDisksAutoDeleteEnumRef(s)
										} else {
											return nil, fmt.Errorf("rStatefulPolicyPreservedStateDisksObj.AutoDelete: expected string")
										}
									}
									m[k] = rStatefulPolicyPreservedStateDisksObj
								} else {
									return nil, fmt.Errorf("r.StatefulPolicy.PreservedState.Disks: expected map[string]interface{}")
								}
							}
							r.StatefulPolicy.PreservedState.Disks = m
						} else {
							return nil, fmt.Errorf("r.StatefulPolicy.PreservedState.Disks: expected map[string]interface{}")
						}
					}
					if _, ok := rStatefulPolicyPreservedState["externalIps"]; ok {
						if rStatefulPolicyPreservedStateExternalIps, ok := rStatefulPolicyPreservedState["externalIps"].(map[string]interface{}); ok {
							m := make(map[string]dclService.InstanceGroupManagerStatefulPolicyPreservedStateExternalIps)
							for k, v := range rStatefulPolicyPreservedStateExternalIps {
								if objval, ok := v.(map[string]interface{}); ok {
									var rStatefulPolicyPreservedStateExternalIpsObj dclService.InstanceGroupManagerStatefulPolicyPreservedStateExternalIps
									if _, ok := objval["autoDelete"]; ok {
										if s, ok := objval["autoDelete"].(string); ok {
											rStatefulPolicyPreservedStateExternalIpsObj.AutoDelete = dclService.InstanceGroupManagerStatefulPolicyPreservedStateExternalIpsAutoDeleteEnumRef(s)
										} else {
											return nil, fmt.Errorf("rStatefulPolicyPreservedStateExternalIpsObj.AutoDelete: expected string")
										}
									}
									m[k] = rStatefulPolicyPreservedStateExternalIpsObj
								} else {
									return nil, fmt.Errorf("r.StatefulPolicy.PreservedState.ExternalIps: expected map[string]interface{}")
								}
							}
							r.StatefulPolicy.PreservedState.ExternalIps = m
						} else {
							return nil, fmt.Errorf("r.StatefulPolicy.PreservedState.ExternalIps: expected map[string]interface{}")
						}
					}
					if _, ok := rStatefulPolicyPreservedState["internalIps"]; ok {
						if rStatefulPolicyPreservedStateInternalIps, ok := rStatefulPolicyPreservedState["internalIps"].(map[string]interface{}); ok {
							m := make(map[string]dclService.InstanceGroupManagerStatefulPolicyPreservedStateInternalIps)
							for k, v := range rStatefulPolicyPreservedStateInternalIps {
								if objval, ok := v.(map[string]interface{}); ok {
									var rStatefulPolicyPreservedStateInternalIpsObj dclService.InstanceGroupManagerStatefulPolicyPreservedStateInternalIps
									if _, ok := objval["autoDelete"]; ok {
										if s, ok := objval["autoDelete"].(string); ok {
											rStatefulPolicyPreservedStateInternalIpsObj.AutoDelete = dclService.InstanceGroupManagerStatefulPolicyPreservedStateInternalIpsAutoDeleteEnumRef(s)
										} else {
											return nil, fmt.Errorf("rStatefulPolicyPreservedStateInternalIpsObj.AutoDelete: expected string")
										}
									}
									m[k] = rStatefulPolicyPreservedStateInternalIpsObj
								} else {
									return nil, fmt.Errorf("r.StatefulPolicy.PreservedState.InternalIps: expected map[string]interface{}")
								}
							}
							r.StatefulPolicy.PreservedState.InternalIps = m
						} else {
							return nil, fmt.Errorf("r.StatefulPolicy.PreservedState.InternalIps: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.StatefulPolicy.PreservedState: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.StatefulPolicy: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["status"]; ok {
		if rStatus, ok := u.Object["status"].(map[string]interface{}); ok {
			r.Status = &dclService.InstanceGroupManagerStatus{}
			if _, ok := rStatus["autoscaler"]; ok {
				if s, ok := rStatus["autoscaler"].(string); ok {
					r.Status.Autoscaler = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Status.Autoscaler: expected string")
				}
			}
			if _, ok := rStatus["isStable"]; ok {
				if b, ok := rStatus["isStable"].(bool); ok {
					r.Status.IsStable = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Status.IsStable: expected bool")
				}
			}
			if _, ok := rStatus["stateful"]; ok {
				if rStatusStateful, ok := rStatus["stateful"].(map[string]interface{}); ok {
					r.Status.Stateful = &dclService.InstanceGroupManagerStatusStateful{}
					if _, ok := rStatusStateful["hasStatefulConfig"]; ok {
						if b, ok := rStatusStateful["hasStatefulConfig"].(bool); ok {
							r.Status.Stateful.HasStatefulConfig = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Status.Stateful.HasStatefulConfig: expected bool")
						}
					}
					if _, ok := rStatusStateful["isStateful"]; ok {
						if b, ok := rStatusStateful["isStateful"].(bool); ok {
							r.Status.Stateful.IsStateful = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Status.Stateful.IsStateful: expected bool")
						}
					}
					if _, ok := rStatusStateful["perInstanceConfigs"]; ok {
						if rStatusStatefulPerInstanceConfigs, ok := rStatusStateful["perInstanceConfigs"].(map[string]interface{}); ok {
							r.Status.Stateful.PerInstanceConfigs = &dclService.InstanceGroupManagerStatusStatefulPerInstanceConfigs{}
							if _, ok := rStatusStatefulPerInstanceConfigs["allEffective"]; ok {
								if b, ok := rStatusStatefulPerInstanceConfigs["allEffective"].(bool); ok {
									r.Status.Stateful.PerInstanceConfigs.AllEffective = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Status.Stateful.PerInstanceConfigs.AllEffective: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Status.Stateful.PerInstanceConfigs: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Status.Stateful: expected map[string]interface{}")
				}
			}
			if _, ok := rStatus["versionTarget"]; ok {
				if rStatusVersionTarget, ok := rStatus["versionTarget"].(map[string]interface{}); ok {
					r.Status.VersionTarget = &dclService.InstanceGroupManagerStatusVersionTarget{}
					if _, ok := rStatusVersionTarget["isReached"]; ok {
						if b, ok := rStatusVersionTarget["isReached"].(bool); ok {
							r.Status.VersionTarget.IsReached = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Status.VersionTarget.IsReached: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Status.VersionTarget: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Status: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["targetPools"]; ok {
		if s, ok := u.Object["targetPools"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.TargetPools = append(r.TargetPools, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.TargetPools: expected []interface{}")
		}
	}
	if _, ok := u.Object["targetSize"]; ok {
		if i, ok := u.Object["targetSize"].(int64); ok {
			r.TargetSize = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.TargetSize: expected int64")
		}
	}
	if _, ok := u.Object["updatePolicy"]; ok {
		if rUpdatePolicy, ok := u.Object["updatePolicy"].(map[string]interface{}); ok {
			r.UpdatePolicy = &dclService.InstanceGroupManagerUpdatePolicy{}
			if _, ok := rUpdatePolicy["instanceRedistributionType"]; ok {
				if s, ok := rUpdatePolicy["instanceRedistributionType"].(string); ok {
					r.UpdatePolicy.InstanceRedistributionType = dclService.InstanceGroupManagerUpdatePolicyInstanceRedistributionTypeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.UpdatePolicy.InstanceRedistributionType: expected string")
				}
			}
			if _, ok := rUpdatePolicy["maxSurge"]; ok {
				if rUpdatePolicyMaxSurge, ok := rUpdatePolicy["maxSurge"].(map[string]interface{}); ok {
					r.UpdatePolicy.MaxSurge = &dclService.InstanceGroupManagerUpdatePolicyMaxSurge{}
					if _, ok := rUpdatePolicyMaxSurge["calculated"]; ok {
						if i, ok := rUpdatePolicyMaxSurge["calculated"].(int64); ok {
							r.UpdatePolicy.MaxSurge.Calculated = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.UpdatePolicy.MaxSurge.Calculated: expected int64")
						}
					}
					if _, ok := rUpdatePolicyMaxSurge["fixed"]; ok {
						if i, ok := rUpdatePolicyMaxSurge["fixed"].(int64); ok {
							r.UpdatePolicy.MaxSurge.Fixed = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.UpdatePolicy.MaxSurge.Fixed: expected int64")
						}
					}
					if _, ok := rUpdatePolicyMaxSurge["percent"]; ok {
						if i, ok := rUpdatePolicyMaxSurge["percent"].(int64); ok {
							r.UpdatePolicy.MaxSurge.Percent = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.UpdatePolicy.MaxSurge.Percent: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.UpdatePolicy.MaxSurge: expected map[string]interface{}")
				}
			}
			if _, ok := rUpdatePolicy["maxUnavailable"]; ok {
				if rUpdatePolicyMaxUnavailable, ok := rUpdatePolicy["maxUnavailable"].(map[string]interface{}); ok {
					r.UpdatePolicy.MaxUnavailable = &dclService.InstanceGroupManagerUpdatePolicyMaxUnavailable{}
					if _, ok := rUpdatePolicyMaxUnavailable["calculated"]; ok {
						if i, ok := rUpdatePolicyMaxUnavailable["calculated"].(int64); ok {
							r.UpdatePolicy.MaxUnavailable.Calculated = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.UpdatePolicy.MaxUnavailable.Calculated: expected int64")
						}
					}
					if _, ok := rUpdatePolicyMaxUnavailable["fixed"]; ok {
						if i, ok := rUpdatePolicyMaxUnavailable["fixed"].(int64); ok {
							r.UpdatePolicy.MaxUnavailable.Fixed = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.UpdatePolicy.MaxUnavailable.Fixed: expected int64")
						}
					}
					if _, ok := rUpdatePolicyMaxUnavailable["percent"]; ok {
						if i, ok := rUpdatePolicyMaxUnavailable["percent"].(int64); ok {
							r.UpdatePolicy.MaxUnavailable.Percent = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.UpdatePolicy.MaxUnavailable.Percent: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.UpdatePolicy.MaxUnavailable: expected map[string]interface{}")
				}
			}
			if _, ok := rUpdatePolicy["minReadySec"]; ok {
				if i, ok := rUpdatePolicy["minReadySec"].(int64); ok {
					r.UpdatePolicy.MinReadySec = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.UpdatePolicy.MinReadySec: expected int64")
				}
			}
			if _, ok := rUpdatePolicy["minimalAction"]; ok {
				if s, ok := rUpdatePolicy["minimalAction"].(string); ok {
					r.UpdatePolicy.MinimalAction = dclService.InstanceGroupManagerUpdatePolicyMinimalActionEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.UpdatePolicy.MinimalAction: expected string")
				}
			}
			if _, ok := rUpdatePolicy["mostDisruptiveAllowedAction"]; ok {
				if s, ok := rUpdatePolicy["mostDisruptiveAllowedAction"].(string); ok {
					r.UpdatePolicy.MostDisruptiveAllowedAction = dclService.InstanceGroupManagerUpdatePolicyMostDisruptiveAllowedActionEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.UpdatePolicy.MostDisruptiveAllowedAction: expected string")
				}
			}
			if _, ok := rUpdatePolicy["replacementMethod"]; ok {
				if s, ok := rUpdatePolicy["replacementMethod"].(string); ok {
					r.UpdatePolicy.ReplacementMethod = dclService.InstanceGroupManagerUpdatePolicyReplacementMethodEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.UpdatePolicy.ReplacementMethod: expected string")
				}
			}
			if _, ok := rUpdatePolicy["type"]; ok {
				if s, ok := rUpdatePolicy["type"].(string); ok {
					r.UpdatePolicy.Type = dclService.InstanceGroupManagerUpdatePolicyTypeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.UpdatePolicy.Type: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.UpdatePolicy: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["versions"]; ok {
		if s, ok := u.Object["versions"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rVersions dclService.InstanceGroupManagerVersions
					if _, ok := objval["instanceTemplate"]; ok {
						if s, ok := objval["instanceTemplate"].(string); ok {
							rVersions.InstanceTemplate = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rVersions.InstanceTemplate: expected string")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rVersions.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rVersions.Name: expected string")
						}
					}
					if _, ok := objval["targetSize"]; ok {
						if rVersionsTargetSize, ok := objval["targetSize"].(map[string]interface{}); ok {
							rVersions.TargetSize = &dclService.InstanceGroupManagerVersionsTargetSize{}
							if _, ok := rVersionsTargetSize["calculated"]; ok {
								if i, ok := rVersionsTargetSize["calculated"].(int64); ok {
									rVersions.TargetSize.Calculated = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rVersions.TargetSize.Calculated: expected int64")
								}
							}
							if _, ok := rVersionsTargetSize["fixed"]; ok {
								if i, ok := rVersionsTargetSize["fixed"].(int64); ok {
									rVersions.TargetSize.Fixed = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rVersions.TargetSize.Fixed: expected int64")
								}
							}
							if _, ok := rVersionsTargetSize["percent"]; ok {
								if i, ok := rVersionsTargetSize["percent"].(int64); ok {
									rVersions.TargetSize.Percent = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("rVersions.TargetSize.Percent: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("rVersions.TargetSize: expected map[string]interface{}")
						}
					}
					r.Versions = append(r.Versions, rVersions)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Versions: expected []interface{}")
		}
	}
	if _, ok := u.Object["zone"]; ok {
		if s, ok := u.Object["zone"].(string); ok {
			r.Zone = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Zone: expected string")
		}
	}
	return r, nil
}

func GetInstanceGroupManager(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstanceGroupManager(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetInstanceGroupManager(ctx, r)
	if err != nil {
		return nil, err
	}
	return InstanceGroupManagerToUnstructured(r), nil
}

func ListInstanceGroupManager(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListInstanceGroupManager(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, InstanceGroupManagerToUnstructured(r))
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

func ApplyInstanceGroupManager(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstanceGroupManager(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInstanceGroupManager(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyInstanceGroupManager(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return InstanceGroupManagerToUnstructured(r), nil
}

func InstanceGroupManagerHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstanceGroupManager(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInstanceGroupManager(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyInstanceGroupManager(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteInstanceGroupManager(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstanceGroupManager(u)
	if err != nil {
		return err
	}
	return c.DeleteInstanceGroupManager(ctx, r)
}

func InstanceGroupManagerID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToInstanceGroupManager(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *InstanceGroupManager) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"InstanceGroupManager",
		"beta",
	}
}

func (r *InstanceGroupManager) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InstanceGroupManager) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InstanceGroupManager) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *InstanceGroupManager) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InstanceGroupManager) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InstanceGroupManager) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *InstanceGroupManager) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetInstanceGroupManager(ctx, config, resource)
}

func (r *InstanceGroupManager) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyInstanceGroupManager(ctx, config, resource, opts...)
}

func (r *InstanceGroupManager) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return InstanceGroupManagerHasDiff(ctx, config, resource, opts...)
}

func (r *InstanceGroupManager) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteInstanceGroupManager(ctx, config, resource)
}

func (r *InstanceGroupManager) ID(resource *unstructured.Resource) (string, error) {
	return InstanceGroupManagerID(resource)
}

func init() {
	unstructured.Register(&InstanceGroupManager{})
}
