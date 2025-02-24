// Copyright 2023 Google LLC. All Rights Reserved.
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

type Job struct{}

func JobToUnstructured(r *dclService.Job) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "run",
			Version: "alpha",
			Type:    "Job",
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
	if r.BinaryAuthorization != nil && r.BinaryAuthorization != dclService.EmptyJobBinaryAuthorization {
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
	var rConditions []interface{}
	for _, rConditionsVal := range r.Conditions {
		rConditionsObject := make(map[string]interface{})
		if rConditionsVal.ExecutionReason != nil {
			rConditionsObject["executionReason"] = string(*rConditionsVal.ExecutionReason)
		}
		if rConditionsVal.LastTransitionTime != nil {
			rConditionsObject["lastTransitionTime"] = *rConditionsVal.LastTransitionTime
		}
		if rConditionsVal.Message != nil {
			rConditionsObject["message"] = *rConditionsVal.Message
		}
		if rConditionsVal.Reason != nil {
			rConditionsObject["reason"] = string(*rConditionsVal.Reason)
		}
		if rConditionsVal.RevisionReason != nil {
			rConditionsObject["revisionReason"] = string(*rConditionsVal.RevisionReason)
		}
		if rConditionsVal.Severity != nil {
			rConditionsObject["severity"] = string(*rConditionsVal.Severity)
		}
		if rConditionsVal.State != nil {
			rConditionsObject["state"] = string(*rConditionsVal.State)
		}
		if rConditionsVal.Type != nil {
			rConditionsObject["type"] = *rConditionsVal.Type
		}
		rConditions = append(rConditions, rConditionsObject)
	}
	u.Object["conditions"] = rConditions
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Creator != nil {
		u.Object["creator"] = *r.Creator
	}
	if r.DeleteTime != nil {
		u.Object["deleteTime"] = *r.DeleteTime
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.ExecutionCount != nil {
		u.Object["executionCount"] = *r.ExecutionCount
	}
	if r.ExpireTime != nil {
		u.Object["expireTime"] = *r.ExpireTime
	}
	if r.Generation != nil {
		u.Object["generation"] = *r.Generation
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
	if r.LatestCreatedExecution != nil && r.LatestCreatedExecution != dclService.EmptyJobLatestCreatedExecution {
		rLatestCreatedExecution := make(map[string]interface{})
		if r.LatestCreatedExecution.CreateTime != nil {
			rLatestCreatedExecution["createTime"] = *r.LatestCreatedExecution.CreateTime
		}
		if r.LatestCreatedExecution.Name != nil {
			rLatestCreatedExecution["name"] = *r.LatestCreatedExecution.Name
		}
		u.Object["latestCreatedExecution"] = rLatestCreatedExecution
	}
	if r.LatestSucceededExecution != nil && r.LatestSucceededExecution != dclService.EmptyJobLatestSucceededExecution {
		rLatestSucceededExecution := make(map[string]interface{})
		if r.LatestSucceededExecution.CreateTime != nil {
			rLatestSucceededExecution["createTime"] = *r.LatestSucceededExecution.CreateTime
		}
		if r.LatestSucceededExecution.Name != nil {
			rLatestSucceededExecution["name"] = *r.LatestSucceededExecution.Name
		}
		u.Object["latestSucceededExecution"] = rLatestSucceededExecution
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
	if r.ObservedGeneration != nil {
		u.Object["observedGeneration"] = *r.ObservedGeneration
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Reconciling != nil {
		u.Object["reconciling"] = *r.Reconciling
	}
	if r.Template != nil && r.Template != dclService.EmptyJobTemplate {
		rTemplate := make(map[string]interface{})
		if r.Template.Annotations != nil {
			rTemplateAnnotations := make(map[string]interface{})
			for k, v := range r.Template.Annotations {
				rTemplateAnnotations[k] = v
			}
			rTemplate["annotations"] = rTemplateAnnotations
		}
		if r.Template.Labels != nil {
			rTemplateLabels := make(map[string]interface{})
			for k, v := range r.Template.Labels {
				rTemplateLabels[k] = v
			}
			rTemplate["labels"] = rTemplateLabels
		}
		if r.Template.Parallelism != nil {
			rTemplate["parallelism"] = *r.Template.Parallelism
		}
		if r.Template.TaskCount != nil {
			rTemplate["taskCount"] = *r.Template.TaskCount
		}
		if r.Template.Template != nil && r.Template.Template != dclService.EmptyJobTemplateTemplate {
			rTemplateTemplate := make(map[string]interface{})
			var rTemplateTemplateContainers []interface{}
			for _, rTemplateTemplateContainersVal := range r.Template.Template.Containers {
				rTemplateTemplateContainersObject := make(map[string]interface{})
				var rTemplateTemplateContainersValArgs []interface{}
				for _, rTemplateTemplateContainersValArgsVal := range rTemplateTemplateContainersVal.Args {
					rTemplateTemplateContainersValArgs = append(rTemplateTemplateContainersValArgs, rTemplateTemplateContainersValArgsVal)
				}
				rTemplateTemplateContainersObject["args"] = rTemplateTemplateContainersValArgs
				var rTemplateTemplateContainersValCommand []interface{}
				for _, rTemplateTemplateContainersValCommandVal := range rTemplateTemplateContainersVal.Command {
					rTemplateTemplateContainersValCommand = append(rTemplateTemplateContainersValCommand, rTemplateTemplateContainersValCommandVal)
				}
				rTemplateTemplateContainersObject["command"] = rTemplateTemplateContainersValCommand
				var rTemplateTemplateContainersValEnv []interface{}
				for _, rTemplateTemplateContainersValEnvVal := range rTemplateTemplateContainersVal.Env {
					rTemplateTemplateContainersValEnvObject := make(map[string]interface{})
					if rTemplateTemplateContainersValEnvVal.Name != nil {
						rTemplateTemplateContainersValEnvObject["name"] = *rTemplateTemplateContainersValEnvVal.Name
					}
					if rTemplateTemplateContainersValEnvVal.Value != nil {
						rTemplateTemplateContainersValEnvObject["value"] = *rTemplateTemplateContainersValEnvVal.Value
					}
					if rTemplateTemplateContainersValEnvVal.ValueSource != nil && rTemplateTemplateContainersValEnvVal.ValueSource != dclService.EmptyJobTemplateTemplateContainersEnvValueSource {
						rTemplateTemplateContainersValEnvValValueSource := make(map[string]interface{})
						if rTemplateTemplateContainersValEnvVal.ValueSource.SecretKeyRef != nil && rTemplateTemplateContainersValEnvVal.ValueSource.SecretKeyRef != dclService.EmptyJobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
							rTemplateTemplateContainersValEnvValValueSourceSecretKeyRef := make(map[string]interface{})
							if rTemplateTemplateContainersValEnvVal.ValueSource.SecretKeyRef.Secret != nil {
								rTemplateTemplateContainersValEnvValValueSourceSecretKeyRef["secret"] = *rTemplateTemplateContainersValEnvVal.ValueSource.SecretKeyRef.Secret
							}
							if rTemplateTemplateContainersValEnvVal.ValueSource.SecretKeyRef.Version != nil {
								rTemplateTemplateContainersValEnvValValueSourceSecretKeyRef["version"] = *rTemplateTemplateContainersValEnvVal.ValueSource.SecretKeyRef.Version
							}
							rTemplateTemplateContainersValEnvValValueSource["secretKeyRef"] = rTemplateTemplateContainersValEnvValValueSourceSecretKeyRef
						}
						rTemplateTemplateContainersValEnvObject["valueSource"] = rTemplateTemplateContainersValEnvValValueSource
					}
					rTemplateTemplateContainersValEnv = append(rTemplateTemplateContainersValEnv, rTemplateTemplateContainersValEnvObject)
				}
				rTemplateTemplateContainersObject["env"] = rTemplateTemplateContainersValEnv
				if rTemplateTemplateContainersVal.Image != nil {
					rTemplateTemplateContainersObject["image"] = *rTemplateTemplateContainersVal.Image
				}
				if rTemplateTemplateContainersVal.Name != nil {
					rTemplateTemplateContainersObject["name"] = *rTemplateTemplateContainersVal.Name
				}
				var rTemplateTemplateContainersValPorts []interface{}
				for _, rTemplateTemplateContainersValPortsVal := range rTemplateTemplateContainersVal.Ports {
					rTemplateTemplateContainersValPortsObject := make(map[string]interface{})
					if rTemplateTemplateContainersValPortsVal.ContainerPort != nil {
						rTemplateTemplateContainersValPortsObject["containerPort"] = *rTemplateTemplateContainersValPortsVal.ContainerPort
					}
					if rTemplateTemplateContainersValPortsVal.Name != nil {
						rTemplateTemplateContainersValPortsObject["name"] = *rTemplateTemplateContainersValPortsVal.Name
					}
					rTemplateTemplateContainersValPorts = append(rTemplateTemplateContainersValPorts, rTemplateTemplateContainersValPortsObject)
				}
				rTemplateTemplateContainersObject["ports"] = rTemplateTemplateContainersValPorts
				if rTemplateTemplateContainersVal.Resources != nil && rTemplateTemplateContainersVal.Resources != dclService.EmptyJobTemplateTemplateContainersResources {
					rTemplateTemplateContainersValResources := make(map[string]interface{})
					if rTemplateTemplateContainersVal.Resources.CpuIdle != nil {
						rTemplateTemplateContainersValResources["cpuIdle"] = *rTemplateTemplateContainersVal.Resources.CpuIdle
					}
					if rTemplateTemplateContainersVal.Resources.Limits != nil {
						rTemplateTemplateContainersValResourcesLimits := make(map[string]interface{})
						for k, v := range rTemplateTemplateContainersVal.Resources.Limits {
							rTemplateTemplateContainersValResourcesLimits[k] = v
						}
						rTemplateTemplateContainersValResources["limits"] = rTemplateTemplateContainersValResourcesLimits
					}
					rTemplateTemplateContainersObject["resources"] = rTemplateTemplateContainersValResources
				}
				var rTemplateTemplateContainersValVolumeMounts []interface{}
				for _, rTemplateTemplateContainersValVolumeMountsVal := range rTemplateTemplateContainersVal.VolumeMounts {
					rTemplateTemplateContainersValVolumeMountsObject := make(map[string]interface{})
					if rTemplateTemplateContainersValVolumeMountsVal.MountPath != nil {
						rTemplateTemplateContainersValVolumeMountsObject["mountPath"] = *rTemplateTemplateContainersValVolumeMountsVal.MountPath
					}
					if rTemplateTemplateContainersValVolumeMountsVal.Name != nil {
						rTemplateTemplateContainersValVolumeMountsObject["name"] = *rTemplateTemplateContainersValVolumeMountsVal.Name
					}
					rTemplateTemplateContainersValVolumeMounts = append(rTemplateTemplateContainersValVolumeMounts, rTemplateTemplateContainersValVolumeMountsObject)
				}
				rTemplateTemplateContainersObject["volumeMounts"] = rTemplateTemplateContainersValVolumeMounts
				rTemplateTemplateContainers = append(rTemplateTemplateContainers, rTemplateTemplateContainersObject)
			}
			rTemplateTemplate["containers"] = rTemplateTemplateContainers
			if r.Template.Template.EncryptionKey != nil {
				rTemplateTemplate["encryptionKey"] = *r.Template.Template.EncryptionKey
			}
			if r.Template.Template.ExecutionEnvironment != nil {
				rTemplateTemplate["executionEnvironment"] = string(*r.Template.Template.ExecutionEnvironment)
			}
			if r.Template.Template.MaxRetries != nil {
				rTemplateTemplate["maxRetries"] = *r.Template.Template.MaxRetries
			}
			if r.Template.Template.ServiceAccount != nil {
				rTemplateTemplate["serviceAccount"] = *r.Template.Template.ServiceAccount
			}
			if r.Template.Template.Timeout != nil {
				rTemplateTemplate["timeout"] = *r.Template.Template.Timeout
			}
			var rTemplateTemplateVolumes []interface{}
			for _, rTemplateTemplateVolumesVal := range r.Template.Template.Volumes {
				rTemplateTemplateVolumesObject := make(map[string]interface{})
				if rTemplateTemplateVolumesVal.CloudSqlInstance != nil && rTemplateTemplateVolumesVal.CloudSqlInstance != dclService.EmptyJobTemplateTemplateVolumesCloudSqlInstance {
					rTemplateTemplateVolumesValCloudSqlInstance := make(map[string]interface{})
					var rTemplateTemplateVolumesValCloudSqlInstanceInstances []interface{}
					for _, rTemplateTemplateVolumesValCloudSqlInstanceInstancesVal := range rTemplateTemplateVolumesVal.CloudSqlInstance.Instances {
						rTemplateTemplateVolumesValCloudSqlInstanceInstances = append(rTemplateTemplateVolumesValCloudSqlInstanceInstances, rTemplateTemplateVolumesValCloudSqlInstanceInstancesVal)
					}
					rTemplateTemplateVolumesValCloudSqlInstance["instances"] = rTemplateTemplateVolumesValCloudSqlInstanceInstances
					rTemplateTemplateVolumesObject["cloudSqlInstance"] = rTemplateTemplateVolumesValCloudSqlInstance
				}
				if rTemplateTemplateVolumesVal.Name != nil {
					rTemplateTemplateVolumesObject["name"] = *rTemplateTemplateVolumesVal.Name
				}
				if rTemplateTemplateVolumesVal.Secret != nil && rTemplateTemplateVolumesVal.Secret != dclService.EmptyJobTemplateTemplateVolumesSecret {
					rTemplateTemplateVolumesValSecret := make(map[string]interface{})
					if rTemplateTemplateVolumesVal.Secret.DefaultMode != nil {
						rTemplateTemplateVolumesValSecret["defaultMode"] = *rTemplateTemplateVolumesVal.Secret.DefaultMode
					}
					var rTemplateTemplateVolumesValSecretItems []interface{}
					for _, rTemplateTemplateVolumesValSecretItemsVal := range rTemplateTemplateVolumesVal.Secret.Items {
						rTemplateTemplateVolumesValSecretItemsObject := make(map[string]interface{})
						if rTemplateTemplateVolumesValSecretItemsVal.Mode != nil {
							rTemplateTemplateVolumesValSecretItemsObject["mode"] = *rTemplateTemplateVolumesValSecretItemsVal.Mode
						}
						if rTemplateTemplateVolumesValSecretItemsVal.Path != nil {
							rTemplateTemplateVolumesValSecretItemsObject["path"] = *rTemplateTemplateVolumesValSecretItemsVal.Path
						}
						if rTemplateTemplateVolumesValSecretItemsVal.Version != nil {
							rTemplateTemplateVolumesValSecretItemsObject["version"] = *rTemplateTemplateVolumesValSecretItemsVal.Version
						}
						rTemplateTemplateVolumesValSecretItems = append(rTemplateTemplateVolumesValSecretItems, rTemplateTemplateVolumesValSecretItemsObject)
					}
					rTemplateTemplateVolumesValSecret["items"] = rTemplateTemplateVolumesValSecretItems
					if rTemplateTemplateVolumesVal.Secret.Secret != nil {
						rTemplateTemplateVolumesValSecret["secret"] = *rTemplateTemplateVolumesVal.Secret.Secret
					}
					rTemplateTemplateVolumesObject["secret"] = rTemplateTemplateVolumesValSecret
				}
				rTemplateTemplateVolumes = append(rTemplateTemplateVolumes, rTemplateTemplateVolumesObject)
			}
			rTemplateTemplate["volumes"] = rTemplateTemplateVolumes
			if r.Template.Template.VPCAccess != nil && r.Template.Template.VPCAccess != dclService.EmptyJobTemplateTemplateVPCAccess {
				rTemplateTemplateVPCAccess := make(map[string]interface{})
				if r.Template.Template.VPCAccess.Connector != nil {
					rTemplateTemplateVPCAccess["connector"] = *r.Template.Template.VPCAccess.Connector
				}
				if r.Template.Template.VPCAccess.Egress != nil {
					rTemplateTemplateVPCAccess["egress"] = string(*r.Template.Template.VPCAccess.Egress)
				}
				rTemplateTemplate["vpcAccess"] = rTemplateTemplateVPCAccess
			}
			rTemplate["template"] = rTemplateTemplate
		}
		u.Object["template"] = rTemplate
	}
	if r.TerminalCondition != nil && r.TerminalCondition != dclService.EmptyJobTerminalCondition {
		rTerminalCondition := make(map[string]interface{})
		if r.TerminalCondition.DomainMappingReason != nil {
			rTerminalCondition["domainMappingReason"] = string(*r.TerminalCondition.DomainMappingReason)
		}
		if r.TerminalCondition.ExecutionReason != nil {
			rTerminalCondition["executionReason"] = string(*r.TerminalCondition.ExecutionReason)
		}
		if r.TerminalCondition.InternalReason != nil {
			rTerminalCondition["internalReason"] = string(*r.TerminalCondition.InternalReason)
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
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToJob(u *unstructured.Resource) (*dclService.Job, error) {
	r := &dclService.Job{}
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
			r.BinaryAuthorization = &dclService.JobBinaryAuthorization{}
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
	if _, ok := u.Object["conditions"]; ok {
		if s, ok := u.Object["conditions"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rConditions dclService.JobConditions
					if _, ok := objval["executionReason"]; ok {
						if s, ok := objval["executionReason"].(string); ok {
							rConditions.ExecutionReason = dclService.JobConditionsExecutionReasonEnumRef(s)
						} else {
							return nil, fmt.Errorf("rConditions.ExecutionReason: expected string")
						}
					}
					if _, ok := objval["lastTransitionTime"]; ok {
						if s, ok := objval["lastTransitionTime"].(string); ok {
							rConditions.LastTransitionTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rConditions.LastTransitionTime: expected string")
						}
					}
					if _, ok := objval["message"]; ok {
						if s, ok := objval["message"].(string); ok {
							rConditions.Message = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rConditions.Message: expected string")
						}
					}
					if _, ok := objval["reason"]; ok {
						if s, ok := objval["reason"].(string); ok {
							rConditions.Reason = dclService.JobConditionsReasonEnumRef(s)
						} else {
							return nil, fmt.Errorf("rConditions.Reason: expected string")
						}
					}
					if _, ok := objval["revisionReason"]; ok {
						if s, ok := objval["revisionReason"].(string); ok {
							rConditions.RevisionReason = dclService.JobConditionsRevisionReasonEnumRef(s)
						} else {
							return nil, fmt.Errorf("rConditions.RevisionReason: expected string")
						}
					}
					if _, ok := objval["severity"]; ok {
						if s, ok := objval["severity"].(string); ok {
							rConditions.Severity = dclService.JobConditionsSeverityEnumRef(s)
						} else {
							return nil, fmt.Errorf("rConditions.Severity: expected string")
						}
					}
					if _, ok := objval["state"]; ok {
						if s, ok := objval["state"].(string); ok {
							rConditions.State = dclService.JobConditionsStateEnumRef(s)
						} else {
							return nil, fmt.Errorf("rConditions.State: expected string")
						}
					}
					if _, ok := objval["type"]; ok {
						if s, ok := objval["type"].(string); ok {
							rConditions.Type = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rConditions.Type: expected string")
						}
					}
					r.Conditions = append(r.Conditions, rConditions)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Conditions: expected []interface{}")
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
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["executionCount"]; ok {
		if i, ok := u.Object["executionCount"].(int64); ok {
			r.ExecutionCount = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.ExecutionCount: expected int64")
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
	if _, ok := u.Object["latestCreatedExecution"]; ok {
		if rLatestCreatedExecution, ok := u.Object["latestCreatedExecution"].(map[string]interface{}); ok {
			r.LatestCreatedExecution = &dclService.JobLatestCreatedExecution{}
			if _, ok := rLatestCreatedExecution["createTime"]; ok {
				if s, ok := rLatestCreatedExecution["createTime"].(string); ok {
					r.LatestCreatedExecution.CreateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LatestCreatedExecution.CreateTime: expected string")
				}
			}
			if _, ok := rLatestCreatedExecution["name"]; ok {
				if s, ok := rLatestCreatedExecution["name"].(string); ok {
					r.LatestCreatedExecution.Name = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LatestCreatedExecution.Name: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LatestCreatedExecution: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["latestSucceededExecution"]; ok {
		if rLatestSucceededExecution, ok := u.Object["latestSucceededExecution"].(map[string]interface{}); ok {
			r.LatestSucceededExecution = &dclService.JobLatestSucceededExecution{}
			if _, ok := rLatestSucceededExecution["createTime"]; ok {
				if s, ok := rLatestSucceededExecution["createTime"].(string); ok {
					r.LatestSucceededExecution.CreateTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LatestSucceededExecution.CreateTime: expected string")
				}
			}
			if _, ok := rLatestSucceededExecution["name"]; ok {
				if s, ok := rLatestSucceededExecution["name"].(string); ok {
					r.LatestSucceededExecution.Name = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LatestSucceededExecution.Name: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LatestSucceededExecution: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["launchStage"]; ok {
		if s, ok := u.Object["launchStage"].(string); ok {
			r.LaunchStage = dclService.JobLaunchStageEnumRef(s)
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
	if _, ok := u.Object["observedGeneration"]; ok {
		if i, ok := u.Object["observedGeneration"].(int64); ok {
			r.ObservedGeneration = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.ObservedGeneration: expected int64")
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
			r.Template = &dclService.JobTemplate{}
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
			if _, ok := rTemplate["parallelism"]; ok {
				if i, ok := rTemplate["parallelism"].(int64); ok {
					r.Template.Parallelism = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.Template.Parallelism: expected int64")
				}
			}
			if _, ok := rTemplate["taskCount"]; ok {
				if i, ok := rTemplate["taskCount"].(int64); ok {
					r.Template.TaskCount = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.Template.TaskCount: expected int64")
				}
			}
			if _, ok := rTemplate["template"]; ok {
				if rTemplateTemplate, ok := rTemplate["template"].(map[string]interface{}); ok {
					r.Template.Template = &dclService.JobTemplateTemplate{}
					if _, ok := rTemplateTemplate["containers"]; ok {
						if s, ok := rTemplateTemplate["containers"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rTemplateTemplateContainers dclService.JobTemplateTemplateContainers
									if _, ok := objval["args"]; ok {
										if s, ok := objval["args"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rTemplateTemplateContainers.Args = append(rTemplateTemplateContainers.Args, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateTemplateContainers.Args: expected []interface{}")
										}
									}
									if _, ok := objval["command"]; ok {
										if s, ok := objval["command"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rTemplateTemplateContainers.Command = append(rTemplateTemplateContainers.Command, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateTemplateContainers.Command: expected []interface{}")
										}
									}
									if _, ok := objval["env"]; ok {
										if s, ok := objval["env"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rTemplateTemplateContainersEnv dclService.JobTemplateTemplateContainersEnv
													if _, ok := objval["name"]; ok {
														if s, ok := objval["name"].(string); ok {
															rTemplateTemplateContainersEnv.Name = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rTemplateTemplateContainersEnv.Name: expected string")
														}
													}
													if _, ok := objval["value"]; ok {
														if s, ok := objval["value"].(string); ok {
															rTemplateTemplateContainersEnv.Value = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rTemplateTemplateContainersEnv.Value: expected string")
														}
													}
													if _, ok := objval["valueSource"]; ok {
														if rTemplateTemplateContainersEnvValueSource, ok := objval["valueSource"].(map[string]interface{}); ok {
															rTemplateTemplateContainersEnv.ValueSource = &dclService.JobTemplateTemplateContainersEnvValueSource{}
															if _, ok := rTemplateTemplateContainersEnvValueSource["secretKeyRef"]; ok {
																if rTemplateTemplateContainersEnvValueSourceSecretKeyRef, ok := rTemplateTemplateContainersEnvValueSource["secretKeyRef"].(map[string]interface{}); ok {
																	rTemplateTemplateContainersEnv.ValueSource.SecretKeyRef = &dclService.JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}
																	if _, ok := rTemplateTemplateContainersEnvValueSourceSecretKeyRef["secret"]; ok {
																		if s, ok := rTemplateTemplateContainersEnvValueSourceSecretKeyRef["secret"].(string); ok {
																			rTemplateTemplateContainersEnv.ValueSource.SecretKeyRef.Secret = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rTemplateTemplateContainersEnv.ValueSource.SecretKeyRef.Secret: expected string")
																		}
																	}
																	if _, ok := rTemplateTemplateContainersEnvValueSourceSecretKeyRef["version"]; ok {
																		if s, ok := rTemplateTemplateContainersEnvValueSourceSecretKeyRef["version"].(string); ok {
																			rTemplateTemplateContainersEnv.ValueSource.SecretKeyRef.Version = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rTemplateTemplateContainersEnv.ValueSource.SecretKeyRef.Version: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rTemplateTemplateContainersEnv.ValueSource.SecretKeyRef: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rTemplateTemplateContainersEnv.ValueSource: expected map[string]interface{}")
														}
													}
													rTemplateTemplateContainers.Env = append(rTemplateTemplateContainers.Env, rTemplateTemplateContainersEnv)
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateTemplateContainers.Env: expected []interface{}")
										}
									}
									if _, ok := objval["image"]; ok {
										if s, ok := objval["image"].(string); ok {
											rTemplateTemplateContainers.Image = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rTemplateTemplateContainers.Image: expected string")
										}
									}
									if _, ok := objval["name"]; ok {
										if s, ok := objval["name"].(string); ok {
											rTemplateTemplateContainers.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rTemplateTemplateContainers.Name: expected string")
										}
									}
									if _, ok := objval["ports"]; ok {
										if s, ok := objval["ports"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rTemplateTemplateContainersPorts dclService.JobTemplateTemplateContainersPorts
													if _, ok := objval["containerPort"]; ok {
														if i, ok := objval["containerPort"].(int64); ok {
															rTemplateTemplateContainersPorts.ContainerPort = dcl.Int64(i)
														} else {
															return nil, fmt.Errorf("rTemplateTemplateContainersPorts.ContainerPort: expected int64")
														}
													}
													if _, ok := objval["name"]; ok {
														if s, ok := objval["name"].(string); ok {
															rTemplateTemplateContainersPorts.Name = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rTemplateTemplateContainersPorts.Name: expected string")
														}
													}
													rTemplateTemplateContainers.Ports = append(rTemplateTemplateContainers.Ports, rTemplateTemplateContainersPorts)
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateTemplateContainers.Ports: expected []interface{}")
										}
									}
									if _, ok := objval["resources"]; ok {
										if rTemplateTemplateContainersResources, ok := objval["resources"].(map[string]interface{}); ok {
											rTemplateTemplateContainers.Resources = &dclService.JobTemplateTemplateContainersResources{}
											if _, ok := rTemplateTemplateContainersResources["cpuIdle"]; ok {
												if b, ok := rTemplateTemplateContainersResources["cpuIdle"].(bool); ok {
													rTemplateTemplateContainers.Resources.CpuIdle = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rTemplateTemplateContainers.Resources.CpuIdle: expected bool")
												}
											}
											if _, ok := rTemplateTemplateContainersResources["limits"]; ok {
												if rTemplateTemplateContainersResourcesLimits, ok := rTemplateTemplateContainersResources["limits"].(map[string]interface{}); ok {
													m := make(map[string]string)
													for k, v := range rTemplateTemplateContainersResourcesLimits {
														if s, ok := v.(string); ok {
															m[k] = s
														}
													}
													rTemplateTemplateContainers.Resources.Limits = m
												} else {
													return nil, fmt.Errorf("rTemplateTemplateContainers.Resources.Limits: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateTemplateContainers.Resources: expected map[string]interface{}")
										}
									}
									if _, ok := objval["volumeMounts"]; ok {
										if s, ok := objval["volumeMounts"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rTemplateTemplateContainersVolumeMounts dclService.JobTemplateTemplateContainersVolumeMounts
													if _, ok := objval["mountPath"]; ok {
														if s, ok := objval["mountPath"].(string); ok {
															rTemplateTemplateContainersVolumeMounts.MountPath = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rTemplateTemplateContainersVolumeMounts.MountPath: expected string")
														}
													}
													if _, ok := objval["name"]; ok {
														if s, ok := objval["name"].(string); ok {
															rTemplateTemplateContainersVolumeMounts.Name = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rTemplateTemplateContainersVolumeMounts.Name: expected string")
														}
													}
													rTemplateTemplateContainers.VolumeMounts = append(rTemplateTemplateContainers.VolumeMounts, rTemplateTemplateContainersVolumeMounts)
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateTemplateContainers.VolumeMounts: expected []interface{}")
										}
									}
									r.Template.Template.Containers = append(r.Template.Template.Containers, rTemplateTemplateContainers)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Template.Template.Containers: expected []interface{}")
						}
					}
					if _, ok := rTemplateTemplate["encryptionKey"]; ok {
						if s, ok := rTemplateTemplate["encryptionKey"].(string); ok {
							r.Template.Template.EncryptionKey = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Template.Template.EncryptionKey: expected string")
						}
					}
					if _, ok := rTemplateTemplate["executionEnvironment"]; ok {
						if s, ok := rTemplateTemplate["executionEnvironment"].(string); ok {
							r.Template.Template.ExecutionEnvironment = dclService.JobTemplateTemplateExecutionEnvironmentEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Template.Template.ExecutionEnvironment: expected string")
						}
					}
					if _, ok := rTemplateTemplate["maxRetries"]; ok {
						if i, ok := rTemplateTemplate["maxRetries"].(int64); ok {
							r.Template.Template.MaxRetries = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Template.Template.MaxRetries: expected int64")
						}
					}
					if _, ok := rTemplateTemplate["serviceAccount"]; ok {
						if s, ok := rTemplateTemplate["serviceAccount"].(string); ok {
							r.Template.Template.ServiceAccount = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Template.Template.ServiceAccount: expected string")
						}
					}
					if _, ok := rTemplateTemplate["timeout"]; ok {
						if s, ok := rTemplateTemplate["timeout"].(string); ok {
							r.Template.Template.Timeout = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Template.Template.Timeout: expected string")
						}
					}
					if _, ok := rTemplateTemplate["volumes"]; ok {
						if s, ok := rTemplateTemplate["volumes"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rTemplateTemplateVolumes dclService.JobTemplateTemplateVolumes
									if _, ok := objval["cloudSqlInstance"]; ok {
										if rTemplateTemplateVolumesCloudSqlInstance, ok := objval["cloudSqlInstance"].(map[string]interface{}); ok {
											rTemplateTemplateVolumes.CloudSqlInstance = &dclService.JobTemplateTemplateVolumesCloudSqlInstance{}
											if _, ok := rTemplateTemplateVolumesCloudSqlInstance["instances"]; ok {
												if s, ok := rTemplateTemplateVolumesCloudSqlInstance["instances"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															rTemplateTemplateVolumes.CloudSqlInstance.Instances = append(rTemplateTemplateVolumes.CloudSqlInstance.Instances, strval)
														}
													}
												} else {
													return nil, fmt.Errorf("rTemplateTemplateVolumes.CloudSqlInstance.Instances: expected []interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateTemplateVolumes.CloudSqlInstance: expected map[string]interface{}")
										}
									}
									if _, ok := objval["name"]; ok {
										if s, ok := objval["name"].(string); ok {
											rTemplateTemplateVolumes.Name = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rTemplateTemplateVolumes.Name: expected string")
										}
									}
									if _, ok := objval["secret"]; ok {
										if rTemplateTemplateVolumesSecret, ok := objval["secret"].(map[string]interface{}); ok {
											rTemplateTemplateVolumes.Secret = &dclService.JobTemplateTemplateVolumesSecret{}
											if _, ok := rTemplateTemplateVolumesSecret["defaultMode"]; ok {
												if i, ok := rTemplateTemplateVolumesSecret["defaultMode"].(int64); ok {
													rTemplateTemplateVolumes.Secret.DefaultMode = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rTemplateTemplateVolumes.Secret.DefaultMode: expected int64")
												}
											}
											if _, ok := rTemplateTemplateVolumesSecret["items"]; ok {
												if s, ok := rTemplateTemplateVolumesSecret["items"].([]interface{}); ok {
													for _, o := range s {
														if objval, ok := o.(map[string]interface{}); ok {
															var rTemplateTemplateVolumesSecretItems dclService.JobTemplateTemplateVolumesSecretItems
															if _, ok := objval["mode"]; ok {
																if i, ok := objval["mode"].(int64); ok {
																	rTemplateTemplateVolumesSecretItems.Mode = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rTemplateTemplateVolumesSecretItems.Mode: expected int64")
																}
															}
															if _, ok := objval["path"]; ok {
																if s, ok := objval["path"].(string); ok {
																	rTemplateTemplateVolumesSecretItems.Path = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rTemplateTemplateVolumesSecretItems.Path: expected string")
																}
															}
															if _, ok := objval["version"]; ok {
																if s, ok := objval["version"].(string); ok {
																	rTemplateTemplateVolumesSecretItems.Version = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rTemplateTemplateVolumesSecretItems.Version: expected string")
																}
															}
															rTemplateTemplateVolumes.Secret.Items = append(rTemplateTemplateVolumes.Secret.Items, rTemplateTemplateVolumesSecretItems)
														}
													}
												} else {
													return nil, fmt.Errorf("rTemplateTemplateVolumes.Secret.Items: expected []interface{}")
												}
											}
											if _, ok := rTemplateTemplateVolumesSecret["secret"]; ok {
												if s, ok := rTemplateTemplateVolumesSecret["secret"].(string); ok {
													rTemplateTemplateVolumes.Secret.Secret = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rTemplateTemplateVolumes.Secret.Secret: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rTemplateTemplateVolumes.Secret: expected map[string]interface{}")
										}
									}
									r.Template.Template.Volumes = append(r.Template.Template.Volumes, rTemplateTemplateVolumes)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Template.Template.Volumes: expected []interface{}")
						}
					}
					if _, ok := rTemplateTemplate["vpcAccess"]; ok {
						if rTemplateTemplateVPCAccess, ok := rTemplateTemplate["vpcAccess"].(map[string]interface{}); ok {
							r.Template.Template.VPCAccess = &dclService.JobTemplateTemplateVPCAccess{}
							if _, ok := rTemplateTemplateVPCAccess["connector"]; ok {
								if s, ok := rTemplateTemplateVPCAccess["connector"].(string); ok {
									r.Template.Template.VPCAccess.Connector = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Template.Template.VPCAccess.Connector: expected string")
								}
							}
							if _, ok := rTemplateTemplateVPCAccess["egress"]; ok {
								if s, ok := rTemplateTemplateVPCAccess["egress"].(string); ok {
									r.Template.Template.VPCAccess.Egress = dclService.JobTemplateTemplateVPCAccessEgressEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.Template.Template.VPCAccess.Egress: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Template.Template.VPCAccess: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Template.Template: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Template: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["terminalCondition"]; ok {
		if rTerminalCondition, ok := u.Object["terminalCondition"].(map[string]interface{}); ok {
			r.TerminalCondition = &dclService.JobTerminalCondition{}
			if _, ok := rTerminalCondition["domainMappingReason"]; ok {
				if s, ok := rTerminalCondition["domainMappingReason"].(string); ok {
					r.TerminalCondition.DomainMappingReason = dclService.JobTerminalConditionDomainMappingReasonEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.DomainMappingReason: expected string")
				}
			}
			if _, ok := rTerminalCondition["executionReason"]; ok {
				if s, ok := rTerminalCondition["executionReason"].(string); ok {
					r.TerminalCondition.ExecutionReason = dclService.JobTerminalConditionExecutionReasonEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.ExecutionReason: expected string")
				}
			}
			if _, ok := rTerminalCondition["internalReason"]; ok {
				if s, ok := rTerminalCondition["internalReason"].(string); ok {
					r.TerminalCondition.InternalReason = dclService.JobTerminalConditionInternalReasonEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.InternalReason: expected string")
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
					r.TerminalCondition.Reason = dclService.JobTerminalConditionReasonEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.Reason: expected string")
				}
			}
			if _, ok := rTerminalCondition["revisionReason"]; ok {
				if s, ok := rTerminalCondition["revisionReason"].(string); ok {
					r.TerminalCondition.RevisionReason = dclService.JobTerminalConditionRevisionReasonEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.RevisionReason: expected string")
				}
			}
			if _, ok := rTerminalCondition["severity"]; ok {
				if s, ok := rTerminalCondition["severity"].(string); ok {
					r.TerminalCondition.Severity = dclService.JobTerminalConditionSeverityEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.TerminalCondition.Severity: expected string")
				}
			}
			if _, ok := rTerminalCondition["state"]; ok {
				if s, ok := rTerminalCondition["state"].(string); ok {
					r.TerminalCondition.State = dclService.JobTerminalConditionStateEnumRef(s)
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
	return r, nil
}

func GetJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJob(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetJob(ctx, r)
	if err != nil {
		return nil, err
	}
	return JobToUnstructured(r), nil
}

func ListJob(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListJob(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, JobToUnstructured(r))
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

func ApplyJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJob(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToJob(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyJob(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return JobToUnstructured(r), nil
}

func JobHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJob(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToJob(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyJob(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToJob(u)
	if err != nil {
		return err
	}
	return c.DeleteJob(ctx, r)
}

func JobID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToJob(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Job) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"run",
		"Job",
		"alpha",
	}
}

func SetPolicyJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToJob(u)
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

func SetPolicyWithEtagJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToJob(u)
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

func GetPolicyJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToJob(u)
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

func SetPolicyMemberJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToJob(u)
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

func GetPolicyMemberJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToJob(u)
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

func DeletePolicyMemberJob(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToJob(u)
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

func (r *Job) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberJob(ctx, config, resource, member)
}

func (r *Job) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberJob(ctx, config, resource, role, member)
}

func (r *Job) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberJob(ctx, config, resource, member)
}

func (r *Job) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyJob(ctx, config, resource, policy)
}

func (r *Job) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagJob(ctx, config, resource, policy)
}

func (r *Job) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyJob(ctx, config, resource)
}

func (r *Job) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetJob(ctx, config, resource)
}

func (r *Job) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyJob(ctx, config, resource, opts...)
}

func (r *Job) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return JobHasDiff(ctx, config, resource, opts...)
}

func (r *Job) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteJob(ctx, config, resource)
}

func (r *Job) ID(resource *unstructured.Resource) (string, error) {
	return JobID(resource)
}

func init() {
	unstructured.Register(&Job{})
}
