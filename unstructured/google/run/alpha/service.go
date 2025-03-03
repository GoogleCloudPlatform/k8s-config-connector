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
package run

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Service struct{}

func ServiceToUnstructured(r *dclService.Service) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "run",
			Version: "alpha",
			Type:    "Service",
		},
		Object: make(map[string]interface{}),
	}
	if r.Annotations != nil {
		rAnnotations := make(map[string]interface{})
		for k, v := range r.Annotations {
			rAnnotations[k] = v
		}
		u.Object["annotations"] = rAnnotations
	}
	if r.BinaryAuthorization != nil && r.BinaryAuthorization != dclService.EmptyServiceBinaryAuthorization {
		rBinaryAuthorization := make(map[string]interface{})
		if r.BinaryAuthorization.BreakglassJustification != nil {
			rBinaryAuthorization["breakglassJustification"] = *r.BinaryAuthorization.BreakglassJustification
		}
		if r.BinaryAuthorization.UseDefault != nil {
			rBinaryAuthorization["useDefault"] = *r.BinaryAuthorization.UseDefault
		}
		u.Object["binaryAuthorization"] = rBinaryAuthorization
	}
	if r.Client != nil {
		u.Object["client"] = *r.Client
	}
	if r.ClientVersion != nil {
		u.Object["clientVersion"] = *r.ClientVersion
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Creator != nil {
		u.Object["creator"] = *r.Creator
	}
	if r.DeleteTime != nil {
		u.Object["deleteTime"] = *r.DeleteTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.ExpireTime != nil {
		u.Object["expireTime"] = *r.ExpireTime
	}
	if r.Generation != nil {
		u.Object["generation"] = *r.Generation
	}
	if r.Ingress != nil {
		u.Object["ingress"] = string(*r.Ingress)
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.LastModifier != nil {
		u.Object["lastModifier"] = *r.LastModifier
	}
	if r.LatestCreatedRevision != nil {
		u.Object["latestCreatedRevision"] = *r.LatestCreatedRevision
	}
	if r.LatestReadyRevision != nil {
		u.Object["latestReadyRevision"] = *r.LatestReadyRevision
	}
	if r.LaunchStage != nil {
		u.Object["launchStage"] = string(*r.LaunchStage)
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Reconciling != nil {
		u.Object["reconciling"] = *r.Reconciling
	}
	if r.Template != nil && r.Template != dclService.EmptyServiceTemplate {
		rTemplate := make(map[string]interface{})
		if r.Template.Annotations != nil {
			rTemplateAnnotations := make(map[string]interface{})
			for k, v := range r.Template.Annotations {
				rTemplateAnnotations[k] = v
			}
			rTemplate["annotations"] = rTemplateAnnotations
		}
		if r.Template.ContainerConcurrency != nil {
			rTemplate["containerConcurrency"] = *r.Template.ContainerConcurrency
		}
		var rTemplateContainers []interface{}
		for _, rTemplateContainersVal := range r.Template.Containers {
			rTemplateContainersObject := make(map[string]interface{})
			var rTemplateContainersValArgs []interface{}
			for _, rTemplateContainersValArgsVal := range rTemplateContainersVal.Args {
				rTemplateContainersValArgs = append(rTemplateContainersValArgs, rTemplateContainersValArgsVal)
			}
			rTemplateContainersObject["args"] = rTemplateContainersValArgs
			var rTemplateContainersValCommand []interface{}
			for _, rTemplateContainersValCommandVal := range rTemplateContainersVal.Command {
				rTemplateContainersValCommand = append(rTemplateContainersValCommand, rTemplateContainersValCommandVal)
			}
			rTemplateContainersObject["command"] = rTemplateContainersValCommand
			var rTemplateContainersValEnv []interface{}
			for _, rTemplateContainersValEnvVal := range rTemplateContainersVal.Env {
				rTemplateContainersValEnvObject := make(map[string]interface{})
				if rTemplateContainersValEnvVal.Name != nil {
					rTemplateContainersValEnvObject["name"] = *rTemplateContainersValEnvVal.Name
				}
				if rTemplateContainersValEnvVal.Value != nil {
					rTemplateContainersValEnvObject["value"] = *rTemplateContainersValEnvVal.Value
				}
				if rTemplateContainersValEnvVal.ValueSource != nil && rTemplateContainersValEnvVal.ValueSource != dclService.EmptyServiceTemplateContainersEnvValueSource {
					rTemplateContainersValEnvValValueSource := make(map[string]interface{})
					if rTemplateContainersValEnvVal.ValueSource.SecretKeyRef != nil && rTemplateContainersValEnvVal.ValueSource.SecretKeyRef != dclService.EmptyServiceTemplateContainersEnvValueSourceSecretKeyRef {
						rTemplateContainersValEnvValValueSourceSecretKeyRef := make(map[string]interface{})
						if rTemplateContainersValEnvVal.ValueSource.SecretKeyRef.Secret != nil {
							rTemplateContainersValEnvValValueSourceSecretKeyRef["secret"] = *rTemplateContainersValEnvVal.ValueSource.SecretKeyRef.Secret
						}
						if rTemplateContainersValEnvVal.ValueSource.SecretKeyRef.Version != nil {
							rTemplateContainersValEnvValValueSourceSecretKeyRef["version"] = *rTemplateContainersValEnvVal.ValueSource.SecretKeyRef.Version
						}
						rTemplateContainersValEnvValValueSource["secretKeyRef"] = rTemplateContainersValEnvValValueSourceSecretKeyRef
					}
					rTemplateContainersValEnvObject["valueSource"] = rTemplateContainersValEnvValValueSource
				}
				rTemplateContainersValEnv = append(rTemplateContainersValEnv, rTemplateContainersValEnvObject)
			}
			rTemplateContainersObject["env"] = rTemplateContainersValEnv
			if rTemplateContainersVal.Image != nil {
				rTemplateContainersObject["image"] = *rTemplateContainersVal.Image
			}
			if rTemplateContainersVal.Name != nil {
				rTemplateContainersObject["name"] = *rTemplateContainersVal.Name
			}
			var rTemplateContainersValPorts []interface{}
			for _, rTemplateContainersValPortsVal := range rTemplateContainersVal.Ports {
				rTemplateContainersValPortsObject := make(map[string]interface{})
				if rTemplateContainersValPortsVal.ContainerPort != nil {
					rTemplateContainersValPortsObject["containerPort"] = *rTemplateContainersValPortsVal.ContainerPort
				}
				if rTemplateContainersValPortsVal.Name != nil {
					rTemplateContainersValPortsObject["name"] = *rTemplateContainersValPortsVal.Name
				}
				rTemplateContainersValPorts = append(rTemplateContainersValPorts, rTemplateContainersValPortsObject)
			}
			rTemplateContainersObject["ports"] = rTemplateContainersValPorts
			if rTemplateContainersVal.Resources != nil && rTemplateContainersVal.Resources != dclService.EmptyServiceTemplateContainersResources {
				rTemplateContainersValResources := make(map[string]interface{})
				if rTemplateContainersVal.Resources.CpuIdle != nil {
					rTemplateContainersValResources["cpuIdle"] = *rTemplateContainersVal.Resources.CpuIdle
				}
				if rTemplateContainersVal.Resources.Limits != nil {
					rTemplateContainersValResourcesLimits := make(map[string]interface{})
					for k, v := range rTemplateContainersVal.Resources.Limits {
						rTemplateContainersValResourcesLimits[k] = v
					}
					rTemplateContainersValResources["limits"] = rTemplateContainersValResourcesLimits
				}
				rTemplateContainersObject["resources"] = rTemplateContainersValResources
			}
			var rTemplateContainersValVolumeMounts []interface{}
			for _, rTemplateContainersValVolumeMountsVal := range rTemplateContainersVal.VolumeMounts {
				rTemplateContainersValVolumeMountsObject := make(map[string]interface{})
				if rTemplateContainersValVolumeMountsVal.MountPath != nil {
					rTemplateContainersValVolumeMountsObject["mountPath"] = *rTemplateContainersValVolumeMountsVal.MountPath
				}
				if rTemplateContainersValVolumeMountsVal.Name != nil {
					rTemplateContainersValVolumeMountsObject["name"] = *rTemplateContainersValVolumeMountsVal.Name
				}
				rTemplateContainersValVolumeMounts = append(rTemplateContainersValVolumeMounts, rTemplateContainersValVolumeMountsObject)
			}
			rTemplateContainersObject["volumeMounts"] = rTemplateContainersValVolumeMounts
			rTemplateContainers = append(rTemplateContainers, rTemplateContainersObject)
		}
		rTemplate["containers"] = rTemplateContainers
		if r.Template.ExecutionEnvironment != nil {
			rTemplate["executionEnvironment"] = string(*r.Template.ExecutionEnvironment)
		}
		if r.Template.Labels != nil {
			rTemplateLabels := make(map[string]interface{})
			for k, v := range r.Template.Labels {
				rTemplateLabels[k] = v
			}
			rTemplate["labels"] = rTemplateLabels
		}
		if r.Template.Revision != nil {
			rTemplate["revision"] = *r.Template.Revision
		}
		if r.Template.Scaling != nil && r.Template.Scaling != dclService.EmptyServiceTemplateScaling {
			rTemplateScaling := make(map[string]interface{})
			if r.Template.Scaling.MaxInstanceCount != nil {
				rTemplateScaling["maxInstanceCount"] = *r.Template.Scaling.MaxInstanceCount
			}
			if r.Template.Scaling.MinInstanceCount != nil {
				rTemplateScaling["minInstanceCount"] = *r.Template.Scaling.MinInstanceCount
			}
			rTemplate["scaling"] = rTemplateScaling
		}
		if r.Template.ServiceAccount != nil {
			rTemplate["serviceAccount"] = *r.Template.ServiceAccount
		}
		if r.Template.Timeout != nil {
			rTemplate["timeout"] = *r.Template.Timeout
		}
		var rTemplateVolumes []interface{}
		for _, rTemplateVolumesVal := range r.Template.Volumes {
			rTemplateVolumesObject := make(map[string]interface{})
			if rTemplateVolumesVal.CloudSqlInstance != nil && rTemplateVolumesVal.CloudSqlInstance != dclService.EmptyServiceTemplateVolumesCloudSqlInstance {
				rTemplateVolumesValCloudSqlInstance := make(map[string]interface{})
				var rTemplateVolumesValCloudSqlInstanceInstances []interface{}
				for _, rTemplateVolumesValCloudSqlInstanceInstancesVal := range rTemplateVolumesVal.CloudSqlInstance.Instances {
					rTemplateVolumesValCloudSqlInstanceInstances = append(rTemplateVolumesValCloudSqlInstanceInstances, rTemplateVolumesValCloudSqlInstanceInstancesVal)
				}
				rTemplateVolumesValCloudSqlInstance["instances"] = rTemplateVolumesValCloudSqlInstanceInstances
				rTemplateVolumesObject["cloudSqlInstance"] = rTemplateVolumesValCloudSqlInstance
			}
			if rTemplateVolumesVal.Name != nil {
				rTemplateVolumesObject["name"] = *rTemplateVolumesVal.Name
			}
			if rTemplateVolumesVal.Secret != nil && rTemplateVolumesVal.Secret != dclService.EmptyServiceTemplateVolumesSecret {
				rTemplateVolumesValSecret := make(map[string]interface{})
				if rTemplateVolumesVal.Secret.DefaultMode != nil {
					rTemplateVolumesValSecret["defaultMode"] = *rTemplateVolumesVal.Secret.DefaultMode
				}
				var rTemplateVolumesValSecretItems []interface{}
				for _, rTemplateVolumesValSecretItemsVal := range rTemplateVolumesVal.Secret.Items {
					rTemplateVolumesValSecretItemsObject := make(map[string]interface{})
					if rTemplateVolumesValSecretItemsVal.Mode != nil {
						rTemplateVolumesValSecretItemsObject["mode"] = *rTemplateVolumesValSecretItemsVal.Mode
					}
					if rTemplateVolumesValSecretItemsVal.Path != nil {
						rTemplateVolumesValSecretItemsObject["path"] = *rTemplateVolumesValSecretItemsVal.Path
					}
					if rTemplateVolumesValSecretItemsVal.Version != nil {
						rTemplateVolumesValSecretItemsObject["version"] = *rTemplateVolumesValSecretItemsVal.Version
					}
					rTemplateVolumesValSecretItems = append(rTemplateVolumesValSecretItems, rTemplateVolumesValSecretItemsObject)
				}
				rTemplateVolumesValSecret["items"] = rTemplateVolumesValSecretItems
				if rTemplateVolumesVal.Secret.Secret != nil {
					rTemplateVolumesValSecret["secret"] = *rTemplateVolumesVal.Secret.Secret
				}
				rTemplateVolumesObject["secret"] = rTemplateVolumesValSecret
			}
			rTemplateVolumes = append(rTemplateVolumes, rTemplateVolumesObject)
		}
		rTemplate["volumes"] = rTemplateVolumes
		if r.Template.VPCAccess != nil && r.Template.VPCAccess != dclService.EmptyServiceTemplateVPCAccess {
			rTemplateVPCAccess := make(map[string]interface{})
			if r.Template.VPCAccess.Connector != nil {
				rTemplateVPCAccess["connector"] = *r.Template.VPCAccess.Connector
			}
			if r.Template.VPCAccess.Egress != nil {
				rTemplateVPCAccess["egress"] = string(*r.Template.VPCAccess.Egress)
			}
			rTemplate["vpcAccess"] = rTemplateVPCAccess
		}
		u.Object["template"] = rTemplate
	}
	if r.TerminalCondition != nil && r.TerminalCondition != dclService.EmptyServiceTerminalCondition {
		rTerminalCondition := make(map[string]interface{})
		if r.TerminalCondition.JobReason != nil {
			rTerminalCondition["jobReason"] = string(*r.TerminalCondition.JobReason)
		}
		if r.TerminalCondition.LastTransitionTime != nil {
			rTerminalCondition["lastTransitionTime"] = *r.TerminalCondition.LastTransitionTime
		}
		if r.TerminalCondition.Message != nil {
			rTerminalCondition["message"] = *r.TerminalCondition.Message
		}
		if r.TerminalCondition.Reason != nil {
			rTerminalCondition["reason"] = string(*r.TerminalCondition.Reason)
		}
		if r.TerminalCondition.RevisionReason != nil {
			rTerminalCondition["revisionReason"] = string(*r.TerminalCondition.RevisionReason)
		}
		if r.TerminalCondition.Severity != nil {
			rTerminalCondition["severity"] = string(*r.TerminalCondition.Severity)
		}
		if r.TerminalCondition.State != nil {
			rTerminalCondition["state"] = string(*r.TerminalCondition.State)
		}
		if r.TerminalCondition.Type != nil {
			rTerminalCondition["type"] = *r.TerminalCondition.Type
		}
		u.Object["terminalCondition"] = rTerminalCondition
	}
	var rTraffic []interface{}
	for _, rTrafficVal := range r.Traffic {
		rTrafficObject := make(map[string]interface{})
		if rTrafficVal.Percent != nil {
			rTrafficObject["percent"] = *rTrafficVal.Percent
		}
		if rTrafficVal.Revision != nil {
			rTrafficObject["revision"] = *rTrafficVal.Revision
		}
		if rTrafficVal.Tag != nil {
			rTrafficObject["tag"] = *rTrafficVal.Tag
		}
		if rTrafficVal.Type != nil {
			rTrafficObject["type"] = string(*rTrafficVal.Type)
		}
		rTraffic = append(rTraffic, rTrafficObject)
	}
	u.Object["traffic"] = rTraffic
	var rTrafficStatuses []interface{}
	for _, rTrafficStatusesVal := range r.TrafficStatuses {
		rTrafficStatusesObject := make(map[string]interface{})
		if rTrafficStatusesVal.Percent != nil {
			rTrafficStatusesObject["percent"] = *rTrafficStatusesVal.Percent
		}
		if rTrafficStatusesVal.Revision != nil {
			rTrafficStatusesObject["revision"] = *rTrafficStatusesVal.Revision
		}
		if rTrafficStatusesVal.Tag != nil {
			rTrafficStatusesObject["tag"] = *rTrafficStatusesVal.Tag
		}
		if rTrafficStatusesVal.Type != nil {
			rTrafficStatusesObject["type"] = string(*rTrafficStatusesVal.Type)
		}
		if rTrafficStatusesVal.Uri != nil {
			rTrafficStatusesObject["uri"] = *rTrafficStatusesVal.Uri
		}
		rTrafficStatuses = append(rTrafficStatuses, rTrafficStatusesObject)
	}
	u.Object["trafficStatuses"] = rTrafficStatuses
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.Uri != nil {
		u.Object["uri"] = *r.Uri
	}
	return u
}

func UnstructuredToService(u *unstructured.Resource) (*dclService.Service, error) {
	r := &dclService.Service{}
	if _, ok := u.Object["annotations"]; ok {
		if rAnnotations, ok := u.Object["annotations"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rAnnotations {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Annotations = m
		} else {
			return nil, fmt.Errorf("r.Annotations: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["binaryAuthorization"]; ok {
		if rBinaryAuthorization, ok := u.Object["binaryAuthorization"].(map[string]interface{}); ok {
			r.BinaryAuthorization = &dclService.ServiceBinaryAuthorization{}
			if _, ok := rBinaryAuthorization["breakglassJustification"]; ok {
				if s, ok := rBinaryAuthorization["breakglassJustification"].(string); ok {
					r.BinaryAuthorization.BreakglassJustification = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.BinaryAuthorization.BreakglassJustification: expected string")
				}
			}
			if _, ok := rBinaryAuthorization["useDefault"]; ok {
				if b, ok := rBinaryAuthorization["useDefault"].(bool); ok {
					r.BinaryAuthorization.UseDefault = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.BinaryAuthorization.UseDefault: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.BinaryAuthorization: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["client"]; ok {
		if s, ok := u.Object["client"].(string); ok {
			r.Client = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Client: expected string")
		}
	}
	if _, ok := u.Object["clientVersion"]; ok {
		if s, ok := u.Object["clientVersion"].(string); ok {
			r.ClientVersion = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ClientVersion: expected string")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["creator"]; ok {
		if s, ok := u.Object["creator"].(string); ok {
			r.Creator = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Creator: expected string")
		}
	}
	if _, ok := u.Object["deleteTime"]; ok {
		if s, ok := u.Object["deleteTime"].(string); ok {
			r.DeleteTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DeleteTime: expected string")
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
	if _, ok := u.Object["expireTime"]; ok {
		if s, ok := u.Object["expireTime"].(string); ok {
			r.ExpireTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ExpireTime: expected string")
		}
	}
	if _, ok := u.Object["generation"]; ok {
		if i, ok := u.Object["generation"].(int64); ok {
			r.Generation = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Generation: expected int64")
		}
	}
	if _, ok := u.Object["ingress"]; ok {
		if s, ok := u.Object["ingress"].(string); ok {
			r.Ingress = dclService.ServiceIngressEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Ingress: expected string")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["lastModifier"]; ok {
		if s, ok := u.Object["lastModifier"].(string); ok {
			r.LastModifier = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LastModifier: expected string")
		}
	}
	if _, ok := u.Object["latestCreatedRevision"]; ok {
		if s, ok := u.Object["latestCreatedRevision"].(string); ok {
			r.LatestCreatedRevision = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LatestCreatedRevision: expected string")
		}
	}
	if _, ok := u.Object["latestReadyRevision"]; ok {
		if s, ok := u.Object["latestReadyRevision"].(string); ok {
			r.LatestReadyRevision = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.LatestReadyRevision: expected string")
		}
	}
	if _, ok := u.Object["launchStage"]; ok {
		if s, ok := u.Object["launchStage"].(string); ok {
			r.LaunchStage = dclService.ServiceLaunchStageEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.LaunchStage: expected string")
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
	if _, ok := u.Object["template"]; ok {
		if rTemplate, ok := u.Object["template"].(map[string]interface{}); ok {
			r.Template = &dclService.ServiceTemplate{}
			if _, ok := rTemplate["annotations"]; ok {
				if rTemplateAnnotations, ok := rTemplate["annotations"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rTemplateAnnotations {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.Template.Annotations = m
				} else {
					return nil, fmt.Errorf("r.Template.Annotations: expected map[string]interface{}")
				}
			}
			if _, ok := rTemplate["containerConcurrency"]; ok {
				if i, ok := rTemplate["containerConcurrency"].(int64); ok {
					r.Template.ContainerConcurrency = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.Template.ContainerConcurrency: expected int64")
				}
			}
			if _, ok := rTemplate["containers"]; ok {
				if s, ok := rTemplate["containers"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rTemplateContainers dclService.ServiceTemplateContainers
							if _, ok := objval["args"]; ok {
								if s, ok := objval["args"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rTemplateContainers.Args = append(rTemplateContainers.Args, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rTemplateContainers.Args: expected []interface{}")
								}
							}
							if _, ok := objval["command"]; ok {
								if s, ok := objval["command"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rTemplateContainers.Command = append(rTemplateContainers.Command, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rTemplateContainers.Command: expected []interface{}")
								}
							}
							if _, ok := objval["env"]; ok {
								if s, ok := objval["env"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rTemplateContainersEnv dclService.ServiceTemplateContainersEnv
											if _, ok := objval["name"]; ok {
												if s, ok := objval["name"].(string); ok {
													rTemplateContainersEnv.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rTemplateContainersEnv.Name: expected string")
												}
											}
											if _, ok := objval["value"]; ok {
												if s, ok := objval["value"].(string); ok {
													rTemplateContainersEnv.Value = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rTemplateContainersEnv.Value: expected string")
												}
											}
											if _, ok := objval["valueSource"]; ok {
												if rTemplateContainersEnvValueSource, ok := objval["valueSource"].(map[string]interface{}); ok {
													rTemplateContainersEnv.ValueSource = &dclService.ServiceTemplateContainersEnvValueSource{}
													if _, ok := rTemplateContainersEnvValueSource["secretKeyRef"]; ok {
														if rTemplateContainersEnvValueSourceSecretKeyRef, ok := rTemplateContainersEnvValueSource["secretKeyRef"].(map[string]interface{}); ok {
															rTemplateContainersEnv.ValueSource.SecretKeyRef = &dclService.ServiceTemplateContainersEnvValueSourceSecretKeyRef{}
															if _, ok := rTemplateContainersEnvValueSourceSecretKeyRef["secret"]; ok {
																if s, ok := rTemplateContainersEnvValueSourceSecretKeyRef["secret"].(string); ok {
																	rTemplateContainersEnv.ValueSource.SecretKeyRef.Secret = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rTemplateContainersEnv.ValueSource.SecretKeyRef.Secret: expected string")
																}
															}
															if _, ok := rTemplateContainersEnvValueSourceSecretKeyRef["version"]; ok {
																if s, ok := rTemplateContainersEnvValueSourceSecretKeyRef["version"].(string); ok {
																	rTemplateContainersEnv.ValueSource.SecretKeyRef.Version = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rTemplateContainersEnv.ValueSource.SecretKeyRef.Version: expected string")
																}
															}
														} else {
															return nil, fmt.Errorf("rTemplateContainersEnv.ValueSource.SecretKeyRef: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rTemplateContainersEnv.ValueSource: expected map[string]interface{}")
												}
											}
											rTemplateContainers.Env = append(rTemplateContainers.Env, rTemplateContainersEnv)
										}
									}
								} else {
									return nil, fmt.Errorf("rTemplateContainers.Env: expected []interface{}")
								}
							}
							if _, ok := objval["image"]; ok {
								if s, ok := objval["image"].(string); ok {
									rTemplateContainers.Image = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rTemplateContainers.Image: expected string")
								}
							}
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rTemplateContainers.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rTemplateContainers.Name: expected string")
								}
							}
							if _, ok := objval["ports"]; ok {
								if s, ok := objval["ports"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rTemplateContainersPorts dclService.ServiceTemplateContainersPorts
											if _, ok := objval["containerPort"]; ok {
												if i, ok := objval["containerPort"].(int64); ok {
													rTemplateContainersPorts.ContainerPort = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rTemplateContainersPorts.ContainerPort: expected int64")
												}
											}
											if _, ok := objval["name"]; ok {
												if s, ok := objval["name"].(string); ok {
													rTemplateContainersPorts.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rTemplateContainersPorts.Name: expected string")
												}
											}
											rTemplateContainers.Ports = append(rTemplateContainers.Ports, rTemplateContainersPorts)
										}
									}
								} else {
									return nil, fmt.Errorf("rTemplateContainers.Ports: expected []interface{}")
								}
							}
							if _, ok := objval["resources"]; ok {
								if rTemplateContainersResources, ok := objval["resources"].(map[string]interface{}); ok {
									rTemplateContainers.Resources = &dclService.ServiceTemplateContainersResources{}
									if _, ok := rTemplateContainersResources["cpuIdle"]; ok {
										if b, ok := rTemplateContainersResources["cpuIdle"].(bool); ok {
											rTemplateContainers.Resources.CpuIdle = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rTemplateContainers.Resources.CpuIdle: expected bool")
										}
									}
									if _, ok := rTemplateContainersResources["limits"]; ok {
										if rTemplateContainersResourcesLimits, ok := rTemplateContainersResources["limits"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rTemplateContainersResourcesLimits {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rTemplateContainers.Resources.Limits = m
										} else {
											return nil, fmt.Errorf("rTemplateContainers.Resources.Limits: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rTemplateContainers.Resources: expected map[string]interface{}")
								}
							}
							if _, ok := objval["volumeMounts"]; ok {
								if s, ok := objval["volumeMounts"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rTemplateContainersVolumeMounts dclService.ServiceTemplateContainersVolumeMounts
											if _, ok := objval["mountPath"]; ok {
												if s, ok := objval["mountPath"].(string); ok {
													rTemplateContainersVolumeMounts.MountPath = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rTemplateContainersVolumeMounts.MountPath: expected string")
												}
											}
											if _, ok := objval["name"]; ok {
												if s, ok := objval["name"].(string); ok {
													rTemplateContainersVolumeMounts.Name = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rTemplateContainersVolumeMounts.Name: expected string")
												}
											}
											rTemplateContainers.VolumeMounts = append(rTemplateContainers.VolumeMounts, rTemplateContainersVolumeMounts)
										}
									}
								} else {
									return nil, fmt.Errorf("rTemplateContainers.VolumeMounts: expected []interface{}")
								}
							}
							r.Template.Containers = append(r.Template.Containers, rTemplateContainers)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Template.Containers: expected []interface{}")
				}
			}
			if _, ok := rTemplate["executionEnvironment"]; ok {
				if s, ok := rTemplate["executionEnvironment"].(string); ok {
					r.Template.ExecutionEnvironment = dclService.ServiceTemplateExecutionEnvironmentEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Template.ExecutionEnvironment: expected string")
				}
			}
			if _, ok := rTemplate["labels"]; ok {
				if rTemplateLabels, ok := rTemplate["labels"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rTemplateLabels {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.Template.Labels = m
				} else {
					return nil, fmt.Errorf("r.Template.Labels: expected map[string]interface{}")
				}
			}
			if _, ok := rTemplate["revision"]; ok {
				if s, ok := rTemplate["revision"].(string); ok {
					r.Template.Revision = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Template.Revision: expected string")
				}
			}
			if _, ok := rTemplate["scaling"]; ok {
				if rTemplateScaling, ok := rTemplate["scaling"].(map[string]interface{}); ok {
					r.Template.Scaling = &dclService.ServiceTemplateScaling{}
					if _, ok := rTemplateScaling["maxInstanceCount"]; ok {
						if i, ok := rTemplateScaling["maxInstanceCount"].(int64); ok {
							r.Template.Scaling.MaxInstanceCount = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Template.Scaling.MaxInstanceCount: expected int64")
						}
					}
					if _, ok := rTemplateScaling["minInstanceCount"]; ok {
						if i, ok := rTemplateScaling["minInstanceCount"].(int64); ok {
							r.Template.Scaling.MinInstanceCount = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Template.Scaling.MinInstanceCount: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Template.Scaling: expected map[string]interface{}")
				}
			}
			if _, ok := rTemplate["serviceAccount"]; ok {
				if s, ok := rTemplate["serviceAccount"].(string); ok {
					r.Template.ServiceAccount = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Template.ServiceAccount: expected string")
				}
			}
			if _, ok := rTemplate["timeout"]; ok {
				if s, ok := rTemplate["timeout"].(string); ok {
					r.Template.Timeout = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Template.Timeout: expected string")
				}
			}
			if _, ok := rTemplate["volumes"]; ok {
				if s, ok := rTemplate["volumes"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rTemplateVolumes dclService.ServiceTemplateVolumes
							if _, ok := objval["cloudSqlInstance"]; ok {
								if rTemplateVolumesCloudSqlInstance, ok := objval["cloudSqlInstance"].(map[string]interface{}); ok {
									rTemplateVolumes.CloudSqlInstance = &dclService.ServiceTemplateVolumesCloudSqlInstance{}
									if _, ok := rTemplateVolumesCloudSqlInstance["instances"]; ok {
										if s, ok := rTemplateVolumesCloudSqlInstance["instances"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rTemplateVolumes.CloudSqlInstance.Instances = append(rTemplateVolumes.CloudSqlInstance.Instances, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateVolumes.CloudSqlInstance.Instances: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rTemplateVolumes.CloudSqlInstance: expected map[string]interface{}")
								}
							}
							if _, ok := objval["name"]; ok {
								if s, ok := objval["name"].(string); ok {
									rTemplateVolumes.Name = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rTemplateVolumes.Name: expected string")
								}
							}
							if _, ok := objval["secret"]; ok {
								if rTemplateVolumesSecret, ok := objval["secret"].(map[string]interface{}); ok {
									rTemplateVolumes.Secret = &dclService.ServiceTemplateVolumesSecret{}
									if _, ok := rTemplateVolumesSecret["defaultMode"]; ok {
										if i, ok := rTemplateVolumesSecret["defaultMode"].(int64); ok {
											rTemplateVolumes.Secret.DefaultMode = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rTemplateVolumes.Secret.DefaultMode: expected int64")
										}
									}
									if _, ok := rTemplateVolumesSecret["items"]; ok {
										if s, ok := rTemplateVolumesSecret["items"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rTemplateVolumesSecretItems dclService.ServiceTemplateVolumesSecretItems
													if _, ok := objval["mode"]; ok {
														if i, ok := objval["mode"].(int64); ok {
															rTemplateVolumesSecretItems.Mode = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rTemplateVolumesSecretItems.Mode: expected int64")
														}
													}
													if _, ok := objval["path"]; ok {
														if s, ok := objval["path"].(string); ok {
															rTemplateVolumesSecretItems.Path = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rTemplateVolumesSecretItems.Path: expected string")
														}
													}
													if _, ok := objval["version"]; ok {
														if s, ok := objval["version"].(string); ok {
															rTemplateVolumesSecretItems.Version = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rTemplateVolumesSecretItems.Version: expected string")
														}
													}
													rTemplateVolumes.Secret.Items = append(rTemplateVolumes.Secret.Items, rTemplateVolumesSecretItems)
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateVolumes.Secret.Items: expected []interface{}")
										}
									}
									if _, ok := rTemplateVolumesSecret["secret"]; ok {
										if s, ok := rTemplateVolumesSecret["secret"].(string); ok {
											rTemplateVolumes.Secret.Secret = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rTemplateVolumes.Secret.Secret: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rTemplateVolumes.Secret: expected map[string]interface{}")
								}
							}
							r.Template.Volumes = append(r.Template.Volumes, rTemplateVolumes)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Template.Volumes: expected []interface{}")
				}
			}
			if _, ok := rTemplate["vpcAccess"]; ok {
				if rTemplateVPCAccess, ok := rTemplate["vpcAccess"].(map[string]interface{}); ok {
					r.Template.VPCAccess = &dclService.ServiceTemplateVPCAccess{}
					if _, ok := rTemplateVPCAccess["connector"]; ok {
						if s, ok := rTemplateVPCAccess["connector"].(string); ok {
							r.Template.VPCAccess.Connector = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Template.VPCAccess.Connector: expected string")
						}
					}
					if _, ok := rTemplateVPCAccess["egress"]; ok {
						if s, ok := rTemplateVPCAccess["egress"].(string); ok {
							r.Template.VPCAccess.Egress = dclService.ServiceTemplateVPCAccessEgressEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Template.VPCAccess.Egress: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Template.VPCAccess: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Template: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["terminalCondition"]; ok {
		if rTerminalCondition, ok := u.Object["terminalCondition"].(map[string]interface{}); ok {
			r.TerminalCondition = &dclService.ServiceTerminalCondition{}
			if _, ok := rTerminalCondition["jobReason"]; ok {
				if s, ok := rTerminalCondition["jobReason"].(string); ok {
					r.TerminalCondition.JobReason = dclService.ServiceTerminalConditionJobReasonEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.JobReason: expected string")
				}
			}
			if _, ok := rTerminalCondition["lastTransitionTime"]; ok {
				if s, ok := rTerminalCondition["lastTransitionTime"].(string); ok {
					r.TerminalCondition.LastTransitionTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.LastTransitionTime: expected string")
				}
			}
			if _, ok := rTerminalCondition["message"]; ok {
				if s, ok := rTerminalCondition["message"].(string); ok {
					r.TerminalCondition.Message = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.Message: expected string")
				}
			}
			if _, ok := rTerminalCondition["reason"]; ok {
				if s, ok := rTerminalCondition["reason"].(string); ok {
					r.TerminalCondition.Reason = dclService.ServiceTerminalConditionReasonEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.Reason: expected string")
				}
			}
			if _, ok := rTerminalCondition["revisionReason"]; ok {
				if s, ok := rTerminalCondition["revisionReason"].(string); ok {
					r.TerminalCondition.RevisionReason = dclService.ServiceTerminalConditionRevisionReasonEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.RevisionReason: expected string")
				}
			}
			if _, ok := rTerminalCondition["severity"]; ok {
				if s, ok := rTerminalCondition["severity"].(string); ok {
					r.TerminalCondition.Severity = dclService.ServiceTerminalConditionSeverityEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.Severity: expected string")
				}
			}
			if _, ok := rTerminalCondition["state"]; ok {
				if s, ok := rTerminalCondition["state"].(string); ok {
					r.TerminalCondition.State = dclService.ServiceTerminalConditionStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.State: expected string")
				}
			}
			if _, ok := rTerminalCondition["type"]; ok {
				if s, ok := rTerminalCondition["type"].(string); ok {
					r.TerminalCondition.Type = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.Type: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.TerminalCondition: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["traffic"]; ok {
		if s, ok := u.Object["traffic"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rTraffic dclService.ServiceTraffic
					if _, ok := objval["percent"]; ok {
						if i, ok := objval["percent"].(int64); ok {
							rTraffic.Percent = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rTraffic.Percent: expected int64")
						}
					}
					if _, ok := objval["revision"]; ok {
						if s, ok := objval["revision"].(string); ok {
							rTraffic.Revision = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rTraffic.Revision: expected string")
						}
					}
					if _, ok := objval["tag"]; ok {
						if s, ok := objval["tag"].(string); ok {
							rTraffic.Tag = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rTraffic.Tag: expected string")
						}
					}
					if _, ok := objval["type"]; ok {
						if s, ok := objval["type"].(string); ok {
							rTraffic.Type = dclService.ServiceTrafficTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rTraffic.Type: expected string")
						}
					}
					r.Traffic = append(r.Traffic, rTraffic)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Traffic: expected []interface{}")
		}
	}
	if _, ok := u.Object["trafficStatuses"]; ok {
		if s, ok := u.Object["trafficStatuses"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rTrafficStatuses dclService.ServiceTrafficStatuses
					if _, ok := objval["percent"]; ok {
						if i, ok := objval["percent"].(int64); ok {
							rTrafficStatuses.Percent = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rTrafficStatuses.Percent: expected int64")
						}
					}
					if _, ok := objval["revision"]; ok {
						if s, ok := objval["revision"].(string); ok {
							rTrafficStatuses.Revision = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rTrafficStatuses.Revision: expected string")
						}
					}
					if _, ok := objval["tag"]; ok {
						if s, ok := objval["tag"].(string); ok {
							rTrafficStatuses.Tag = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rTrafficStatuses.Tag: expected string")
						}
					}
					if _, ok := objval["type"]; ok {
						if s, ok := objval["type"].(string); ok {
							rTrafficStatuses.Type = dclService.ServiceTrafficStatusesTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rTrafficStatuses.Type: expected string")
						}
					}
					if _, ok := objval["uri"]; ok {
						if s, ok := objval["uri"].(string); ok {
							rTrafficStatuses.Uri = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rTrafficStatuses.Uri: expected string")
						}
					}
					r.TrafficStatuses = append(r.TrafficStatuses, rTrafficStatuses)
				}
			}
		} else {
			return nil, fmt.Errorf("r.TrafficStatuses: expected []interface{}")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["uri"]; ok {
		if s, ok := u.Object["uri"].(string); ok {
			r.Uri = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uri: expected string")
		}
	}
	return r, nil
}

func GetService(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetService(ctx, r)
	if err != nil {
		return nil, err
	}
	return ServiceToUnstructured(r), nil
}

func ListService(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListService(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ServiceToUnstructured(r))
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

func ApplyService(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToService(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyService(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ServiceToUnstructured(r), nil
}

func ServiceHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToService(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToService(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyService(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteService(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToService(u)
	if err != nil {
		return err
	}
	return c.DeleteService(ctx, r)
}

func ServiceID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToService(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Service) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"run",
		"Service",
		"alpha",
	}
}

func SetPolicyService(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func SetPolicyWithEtagService(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicyWithEtag(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyService(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func SetPolicyMemberService(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return nil, err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	policy, err := iamClient.SetMember(ctx, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func GetPolicyMemberService(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToService(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policyMember, err := iamClient.GetMember(ctx, r, role, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.MemberToUnstructured(policyMember), nil
}

func DeletePolicyMemberService(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToService(u)
	if err != nil {
		return err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	if err := iamClient.DeleteMember(ctx, member); err != nil {
		return err
	}
	return nil
}

func (r *Service) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberService(ctx, config, resource, member)
}

func (r *Service) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberService(ctx, config, resource, role, member)
}

func (r *Service) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberService(ctx, config, resource, member)
}

func (r *Service) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyService(ctx, config, resource, policy)
}

func (r *Service) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagService(ctx, config, resource, policy)
}

func (r *Service) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyService(ctx, config, resource)
}

func (r *Service) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetService(ctx, config, resource)
}

func (r *Service) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyService(ctx, config, resource, opts...)
}

func (r *Service) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ServiceHasDiff(ctx, config, resource, opts...)
}

func (r *Service) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteService(ctx, config, resource)
}

func (r *Service) ID(resource *unstructured.Resource) (string, error) {
	return ServiceID(resource)
}

func init() {
	unstructured.Register(&Service{})
}
