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
package clouddeploy

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/clouddeploy/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type DeliveryPipeline struct{}

func DeliveryPipelineToUnstructured(r *dclService.DeliveryPipeline) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "clouddeploy",
			Version: "beta",
			Type:    "DeliveryPipeline",
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
	if r.Condition != nil && r.Condition != dclService.EmptyDeliveryPipelineCondition {
		rCondition := make(map[string]interface{})
		if r.Condition.PipelineReadyCondition != nil && r.Condition.PipelineReadyCondition != dclService.EmptyDeliveryPipelineConditionPipelineReadyCondition {
			rConditionPipelineReadyCondition := make(map[string]interface{})
			if r.Condition.PipelineReadyCondition.Status != nil {
				rConditionPipelineReadyCondition["status"] = *r.Condition.PipelineReadyCondition.Status
			}
			if r.Condition.PipelineReadyCondition.UpdateTime != nil {
				rConditionPipelineReadyCondition["updateTime"] = *r.Condition.PipelineReadyCondition.UpdateTime
			}
			rCondition["pipelineReadyCondition"] = rConditionPipelineReadyCondition
		}
		if r.Condition.TargetsPresentCondition != nil && r.Condition.TargetsPresentCondition != dclService.EmptyDeliveryPipelineConditionTargetsPresentCondition {
			rConditionTargetsPresentCondition := make(map[string]interface{})
			var rConditionTargetsPresentConditionMissingTargets []interface{}
			for _, rConditionTargetsPresentConditionMissingTargetsVal := range r.Condition.TargetsPresentCondition.MissingTargets {
				rConditionTargetsPresentConditionMissingTargets = append(rConditionTargetsPresentConditionMissingTargets, rConditionTargetsPresentConditionMissingTargetsVal)
			}
			rConditionTargetsPresentCondition["missingTargets"] = rConditionTargetsPresentConditionMissingTargets
			if r.Condition.TargetsPresentCondition.Status != nil {
				rConditionTargetsPresentCondition["status"] = *r.Condition.TargetsPresentCondition.Status
			}
			if r.Condition.TargetsPresentCondition.UpdateTime != nil {
				rConditionTargetsPresentCondition["updateTime"] = *r.Condition.TargetsPresentCondition.UpdateTime
			}
			rCondition["targetsPresentCondition"] = rConditionTargetsPresentCondition
		}
		if r.Condition.TargetsTypeCondition != nil && r.Condition.TargetsTypeCondition != dclService.EmptyDeliveryPipelineConditionTargetsTypeCondition {
			rConditionTargetsTypeCondition := make(map[string]interface{})
			if r.Condition.TargetsTypeCondition.ErrorDetails != nil {
				rConditionTargetsTypeCondition["errorDetails"] = *r.Condition.TargetsTypeCondition.ErrorDetails
			}
			if r.Condition.TargetsTypeCondition.Status != nil {
				rConditionTargetsTypeCondition["status"] = *r.Condition.TargetsTypeCondition.Status
			}
			rCondition["targetsTypeCondition"] = rConditionTargetsTypeCondition
		}
		u.Object["condition"] = rCondition
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
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
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
	if r.SerialPipeline != nil && r.SerialPipeline != dclService.EmptyDeliveryPipelineSerialPipeline {
		rSerialPipeline := make(map[string]interface{})
		var rSerialPipelineStages []interface{}
		for _, rSerialPipelineStagesVal := range r.SerialPipeline.Stages {
			rSerialPipelineStagesObject := make(map[string]interface{})
			var rSerialPipelineStagesValDeployParameters []interface{}
			for _, rSerialPipelineStagesValDeployParametersVal := range rSerialPipelineStagesVal.DeployParameters {
				rSerialPipelineStagesValDeployParametersObject := make(map[string]interface{})
				if rSerialPipelineStagesValDeployParametersVal.MatchTargetLabels != nil {
					rSerialPipelineStagesValDeployParametersValMatchTargetLabels := make(map[string]interface{})
					for k, v := range rSerialPipelineStagesValDeployParametersVal.MatchTargetLabels {
						rSerialPipelineStagesValDeployParametersValMatchTargetLabels[k] = v
					}
					rSerialPipelineStagesValDeployParametersObject["matchTargetLabels"] = rSerialPipelineStagesValDeployParametersValMatchTargetLabels
				}
				if rSerialPipelineStagesValDeployParametersVal.Values != nil {
					rSerialPipelineStagesValDeployParametersValValues := make(map[string]interface{})
					for k, v := range rSerialPipelineStagesValDeployParametersVal.Values {
						rSerialPipelineStagesValDeployParametersValValues[k] = v
					}
					rSerialPipelineStagesValDeployParametersObject["values"] = rSerialPipelineStagesValDeployParametersValValues
				}
				rSerialPipelineStagesValDeployParameters = append(rSerialPipelineStagesValDeployParameters, rSerialPipelineStagesValDeployParametersObject)
			}
			rSerialPipelineStagesObject["deployParameters"] = rSerialPipelineStagesValDeployParameters
			var rSerialPipelineStagesValProfiles []interface{}
			for _, rSerialPipelineStagesValProfilesVal := range rSerialPipelineStagesVal.Profiles {
				rSerialPipelineStagesValProfiles = append(rSerialPipelineStagesValProfiles, rSerialPipelineStagesValProfilesVal)
			}
			rSerialPipelineStagesObject["profiles"] = rSerialPipelineStagesValProfiles
			if rSerialPipelineStagesVal.Strategy != nil && rSerialPipelineStagesVal.Strategy != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategy {
				rSerialPipelineStagesValStrategy := make(map[string]interface{})
				if rSerialPipelineStagesVal.Strategy.Canary != nil && rSerialPipelineStagesVal.Strategy.Canary != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanary {
					rSerialPipelineStagesValStrategyCanary := make(map[string]interface{})
					if rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment != nil && rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
						rSerialPipelineStagesValStrategyCanaryCanaryDeployment := make(map[string]interface{})
						var rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPercentages []interface{}
						for _, rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPercentagesVal := range rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Percentages {
							rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPercentages = append(rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPercentages, rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPercentagesVal)
						}
						rSerialPipelineStagesValStrategyCanaryCanaryDeployment["percentages"] = rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPercentages
						if rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Postdeploy != nil && rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Postdeploy != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
							rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeploy := make(map[string]interface{})
							var rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeployActions []interface{}
							for _, rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeployActionsVal := range rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Postdeploy.Actions {
								rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeployActions = append(rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeployActions, rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeployActionsVal)
							}
							rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeploy["actions"] = rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeployActions
							rSerialPipelineStagesValStrategyCanaryCanaryDeployment["postdeploy"] = rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPostdeploy
						}
						if rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Predeploy != nil && rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Predeploy != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
							rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeploy := make(map[string]interface{})
							var rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeployActions []interface{}
							for _, rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeployActionsVal := range rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Predeploy.Actions {
								rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeployActions = append(rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeployActions, rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeployActionsVal)
							}
							rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeploy["actions"] = rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeployActions
							rSerialPipelineStagesValStrategyCanaryCanaryDeployment["predeploy"] = rSerialPipelineStagesValStrategyCanaryCanaryDeploymentPredeploy
						}
						if rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Verify != nil {
							rSerialPipelineStagesValStrategyCanaryCanaryDeployment["verify"] = *rSerialPipelineStagesVal.Strategy.Canary.CanaryDeployment.Verify
						}
						rSerialPipelineStagesValStrategyCanary["canaryDeployment"] = rSerialPipelineStagesValStrategyCanaryCanaryDeployment
					}
					if rSerialPipelineStagesVal.Strategy.Canary.CustomCanaryDeployment != nil && rSerialPipelineStagesVal.Strategy.Canary.CustomCanaryDeployment != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment {
						rSerialPipelineStagesValStrategyCanaryCustomCanaryDeployment := make(map[string]interface{})
						var rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigs []interface{}
						for _, rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal := range rSerialPipelineStagesVal.Strategy.Canary.CustomCanaryDeployment.PhaseConfigs {
							rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsObject := make(map[string]interface{})
							if rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Percentage != nil {
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsObject["percentage"] = *rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Percentage
							}
							if rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.PhaseId != nil {
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsObject["phaseId"] = *rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.PhaseId
							}
							if rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Postdeploy != nil && rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Postdeploy != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy {
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeploy := make(map[string]interface{})
								var rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeployActions []interface{}
								for _, rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeployActionsVal := range rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Postdeploy.Actions {
									rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeployActions = append(rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeployActions, rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeployActionsVal)
								}
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeploy["actions"] = rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeployActions
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsObject["postdeploy"] = rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPostdeploy
							}
							if rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Predeploy != nil && rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Predeploy != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy {
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeploy := make(map[string]interface{})
								var rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeployActions []interface{}
								for _, rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeployActionsVal := range rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Predeploy.Actions {
									rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeployActions = append(rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeployActions, rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeployActionsVal)
								}
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeploy["actions"] = rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeployActions
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsObject["predeploy"] = rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValPredeploy
							}
							var rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValProfiles []interface{}
							for _, rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValProfilesVal := range rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Profiles {
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValProfiles = append(rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValProfiles, rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValProfilesVal)
							}
							rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsObject["profiles"] = rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsValProfiles
							if rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Verify != nil {
								rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsObject["verify"] = *rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsVal.Verify
							}
							rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigs = append(rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigs, rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigsObject)
						}
						rSerialPipelineStagesValStrategyCanaryCustomCanaryDeployment["phaseConfigs"] = rSerialPipelineStagesValStrategyCanaryCustomCanaryDeploymentPhaseConfigs
						rSerialPipelineStagesValStrategyCanary["customCanaryDeployment"] = rSerialPipelineStagesValStrategyCanaryCustomCanaryDeployment
					}
					if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig != nil && rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig {
						rSerialPipelineStagesValStrategyCanaryRuntimeConfig := make(map[string]interface{})
						if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.CloudRun != nil && rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.CloudRun != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun {
							rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRun := make(map[string]interface{})
							if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.CloudRun.AutomaticTrafficControl != nil {
								rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRun["automaticTrafficControl"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.CloudRun.AutomaticTrafficControl
							}
							var rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunCanaryRevisionTags []interface{}
							for _, rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunCanaryRevisionTagsVal := range rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.CloudRun.CanaryRevisionTags {
								rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunCanaryRevisionTags = append(rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunCanaryRevisionTags, rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunCanaryRevisionTagsVal)
							}
							rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRun["canaryRevisionTags"] = rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunCanaryRevisionTags
							var rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunPriorRevisionTags []interface{}
							for _, rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunPriorRevisionTagsVal := range rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.CloudRun.PriorRevisionTags {
								rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunPriorRevisionTags = append(rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunPriorRevisionTags, rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunPriorRevisionTagsVal)
							}
							rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRun["priorRevisionTags"] = rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunPriorRevisionTags
							var rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunStableRevisionTags []interface{}
							for _, rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunStableRevisionTagsVal := range rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.CloudRun.StableRevisionTags {
								rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunStableRevisionTags = append(rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunStableRevisionTags, rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunStableRevisionTagsVal)
							}
							rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRun["stableRevisionTags"] = rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRunStableRevisionTags
							rSerialPipelineStagesValStrategyCanaryRuntimeConfig["cloudRun"] = rSerialPipelineStagesValStrategyCanaryRuntimeConfigCloudRun
						}
						if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes != nil && rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes {
							rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetes := make(map[string]interface{})
							if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh != nil && rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh {
								rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh := make(map[string]interface{})
								if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.Deployment != nil {
									rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["deployment"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.Deployment
								}
								if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.HttpRoute != nil {
									rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["httpRoute"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.HttpRoute
								}
								if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.RouteUpdateWaitTime != nil {
									rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["routeUpdateWaitTime"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.RouteUpdateWaitTime
								}
								if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.Service != nil {
									rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["service"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.Service
								}
								if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.StableCutbackDuration != nil {
									rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["stableCutbackDuration"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.StableCutbackDuration
								}
								rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetes["gatewayServiceMesh"] = rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh
							}
							if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking != nil && rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking {
								rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesServiceNetworking := make(map[string]interface{})
								if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.Deployment != nil {
									rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesServiceNetworking["deployment"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.Deployment
								}
								if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.DisablePodOverprovisioning != nil {
									rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesServiceNetworking["disablePodOverprovisioning"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.DisablePodOverprovisioning
								}
								if rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.Service != nil {
									rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesServiceNetworking["service"] = *rSerialPipelineStagesVal.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.Service
								}
								rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetes["serviceNetworking"] = rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetesServiceNetworking
							}
							rSerialPipelineStagesValStrategyCanaryRuntimeConfig["kubernetes"] = rSerialPipelineStagesValStrategyCanaryRuntimeConfigKubernetes
						}
						rSerialPipelineStagesValStrategyCanary["runtimeConfig"] = rSerialPipelineStagesValStrategyCanaryRuntimeConfig
					}
					rSerialPipelineStagesValStrategy["canary"] = rSerialPipelineStagesValStrategyCanary
				}
				if rSerialPipelineStagesVal.Strategy.Standard != nil && rSerialPipelineStagesVal.Strategy.Standard != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyStandard {
					rSerialPipelineStagesValStrategyStandard := make(map[string]interface{})
					if rSerialPipelineStagesVal.Strategy.Standard.Postdeploy != nil && rSerialPipelineStagesVal.Strategy.Standard.Postdeploy != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
						rSerialPipelineStagesValStrategyStandardPostdeploy := make(map[string]interface{})
						var rSerialPipelineStagesValStrategyStandardPostdeployActions []interface{}
						for _, rSerialPipelineStagesValStrategyStandardPostdeployActionsVal := range rSerialPipelineStagesVal.Strategy.Standard.Postdeploy.Actions {
							rSerialPipelineStagesValStrategyStandardPostdeployActions = append(rSerialPipelineStagesValStrategyStandardPostdeployActions, rSerialPipelineStagesValStrategyStandardPostdeployActionsVal)
						}
						rSerialPipelineStagesValStrategyStandardPostdeploy["actions"] = rSerialPipelineStagesValStrategyStandardPostdeployActions
						rSerialPipelineStagesValStrategyStandard["postdeploy"] = rSerialPipelineStagesValStrategyStandardPostdeploy
					}
					if rSerialPipelineStagesVal.Strategy.Standard.Predeploy != nil && rSerialPipelineStagesVal.Strategy.Standard.Predeploy != dclService.EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
						rSerialPipelineStagesValStrategyStandardPredeploy := make(map[string]interface{})
						var rSerialPipelineStagesValStrategyStandardPredeployActions []interface{}
						for _, rSerialPipelineStagesValStrategyStandardPredeployActionsVal := range rSerialPipelineStagesVal.Strategy.Standard.Predeploy.Actions {
							rSerialPipelineStagesValStrategyStandardPredeployActions = append(rSerialPipelineStagesValStrategyStandardPredeployActions, rSerialPipelineStagesValStrategyStandardPredeployActionsVal)
						}
						rSerialPipelineStagesValStrategyStandardPredeploy["actions"] = rSerialPipelineStagesValStrategyStandardPredeployActions
						rSerialPipelineStagesValStrategyStandard["predeploy"] = rSerialPipelineStagesValStrategyStandardPredeploy
					}
					if rSerialPipelineStagesVal.Strategy.Standard.Verify != nil {
						rSerialPipelineStagesValStrategyStandard["verify"] = *rSerialPipelineStagesVal.Strategy.Standard.Verify
					}
					rSerialPipelineStagesValStrategy["standard"] = rSerialPipelineStagesValStrategyStandard
				}
				rSerialPipelineStagesObject["strategy"] = rSerialPipelineStagesValStrategy
			}
			if rSerialPipelineStagesVal.TargetId != nil {
				rSerialPipelineStagesObject["targetId"] = *rSerialPipelineStagesVal.TargetId
			}
			rSerialPipelineStages = append(rSerialPipelineStages, rSerialPipelineStagesObject)
		}
		rSerialPipeline["stages"] = rSerialPipelineStages
		u.Object["serialPipeline"] = rSerialPipeline
	}
	if r.Suspended != nil {
		u.Object["suspended"] = *r.Suspended
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToDeliveryPipeline(u *unstructured.Resource) (*dclService.DeliveryPipeline, error) {
	r := &dclService.DeliveryPipeline{}
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
	if _, ok := u.Object["condition"]; ok {
		if rCondition, ok := u.Object["condition"].(map[string]interface{}); ok {
			r.Condition = &dclService.DeliveryPipelineCondition{}
			if _, ok := rCondition["pipelineReadyCondition"]; ok {
				if rConditionPipelineReadyCondition, ok := rCondition["pipelineReadyCondition"].(map[string]interface{}); ok {
					r.Condition.PipelineReadyCondition = &dclService.DeliveryPipelineConditionPipelineReadyCondition{}
					if _, ok := rConditionPipelineReadyCondition["status"]; ok {
						if b, ok := rConditionPipelineReadyCondition["status"].(bool); ok {
							r.Condition.PipelineReadyCondition.Status = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Condition.PipelineReadyCondition.Status: expected bool")
						}
					}
					if _, ok := rConditionPipelineReadyCondition["updateTime"]; ok {
						if s, ok := rConditionPipelineReadyCondition["updateTime"].(string); ok {
							r.Condition.PipelineReadyCondition.UpdateTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Condition.PipelineReadyCondition.UpdateTime: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Condition.PipelineReadyCondition: expected map[string]interface{}")
				}
			}
			if _, ok := rCondition["targetsPresentCondition"]; ok {
				if rConditionTargetsPresentCondition, ok := rCondition["targetsPresentCondition"].(map[string]interface{}); ok {
					r.Condition.TargetsPresentCondition = &dclService.DeliveryPipelineConditionTargetsPresentCondition{}
					if _, ok := rConditionTargetsPresentCondition["missingTargets"]; ok {
						if s, ok := rConditionTargetsPresentCondition["missingTargets"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Condition.TargetsPresentCondition.MissingTargets = append(r.Condition.TargetsPresentCondition.MissingTargets, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Condition.TargetsPresentCondition.MissingTargets: expected []interface{}")
						}
					}
					if _, ok := rConditionTargetsPresentCondition["status"]; ok {
						if b, ok := rConditionTargetsPresentCondition["status"].(bool); ok {
							r.Condition.TargetsPresentCondition.Status = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Condition.TargetsPresentCondition.Status: expected bool")
						}
					}
					if _, ok := rConditionTargetsPresentCondition["updateTime"]; ok {
						if s, ok := rConditionTargetsPresentCondition["updateTime"].(string); ok {
							r.Condition.TargetsPresentCondition.UpdateTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Condition.TargetsPresentCondition.UpdateTime: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Condition.TargetsPresentCondition: expected map[string]interface{}")
				}
			}
			if _, ok := rCondition["targetsTypeCondition"]; ok {
				if rConditionTargetsTypeCondition, ok := rCondition["targetsTypeCondition"].(map[string]interface{}); ok {
					r.Condition.TargetsTypeCondition = &dclService.DeliveryPipelineConditionTargetsTypeCondition{}
					if _, ok := rConditionTargetsTypeCondition["errorDetails"]; ok {
						if s, ok := rConditionTargetsTypeCondition["errorDetails"].(string); ok {
							r.Condition.TargetsTypeCondition.ErrorDetails = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Condition.TargetsTypeCondition.ErrorDetails: expected string")
						}
					}
					if _, ok := rConditionTargetsTypeCondition["status"]; ok {
						if b, ok := rConditionTargetsTypeCondition["status"].(bool); ok {
							r.Condition.TargetsTypeCondition.Status = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Condition.TargetsTypeCondition.Status: expected bool")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Condition.TargetsTypeCondition: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Condition: expected map[string]interface{}")
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
	if _, ok := u.Object["serialPipeline"]; ok {
		if rSerialPipeline, ok := u.Object["serialPipeline"].(map[string]interface{}); ok {
			r.SerialPipeline = &dclService.DeliveryPipelineSerialPipeline{}
			if _, ok := rSerialPipeline["stages"]; ok {
				if s, ok := rSerialPipeline["stages"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rSerialPipelineStages dclService.DeliveryPipelineSerialPipelineStages
							if _, ok := objval["deployParameters"]; ok {
								if s, ok := objval["deployParameters"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rSerialPipelineStagesDeployParameters dclService.DeliveryPipelineSerialPipelineStagesDeployParameters
											if _, ok := objval["matchTargetLabels"]; ok {
												if rSerialPipelineStagesDeployParametersMatchTargetLabels, ok := objval["matchTargetLabels"].(map[string]interface{}); ok {
													m := make(map[string]string)
													for k, v := range rSerialPipelineStagesDeployParametersMatchTargetLabels {
														if s, ok := v.(string); ok {
															m[k] = s
														}
													}
													rSerialPipelineStagesDeployParameters.MatchTargetLabels = m
												} else {
													return nil, fmt.Errorf("rSerialPipelineStagesDeployParameters.MatchTargetLabels: expected map[string]interface{}")
												}
											}
											if _, ok := objval["values"]; ok {
												if rSerialPipelineStagesDeployParametersValues, ok := objval["values"].(map[string]interface{}); ok {
													m := make(map[string]string)
													for k, v := range rSerialPipelineStagesDeployParametersValues {
														if s, ok := v.(string); ok {
															m[k] = s
														}
													}
													rSerialPipelineStagesDeployParameters.Values = m
												} else {
													return nil, fmt.Errorf("rSerialPipelineStagesDeployParameters.Values: expected map[string]interface{}")
												}
											}
											rSerialPipelineStages.DeployParameters = append(rSerialPipelineStages.DeployParameters, rSerialPipelineStagesDeployParameters)
										}
									}
								} else {
									return nil, fmt.Errorf("rSerialPipelineStages.DeployParameters: expected []interface{}")
								}
							}
							if _, ok := objval["profiles"]; ok {
								if s, ok := objval["profiles"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											rSerialPipelineStages.Profiles = append(rSerialPipelineStages.Profiles, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("rSerialPipelineStages.Profiles: expected []interface{}")
								}
							}
							if _, ok := objval["strategy"]; ok {
								if rSerialPipelineStagesStrategy, ok := objval["strategy"].(map[string]interface{}); ok {
									rSerialPipelineStages.Strategy = &dclService.DeliveryPipelineSerialPipelineStagesStrategy{}
									if _, ok := rSerialPipelineStagesStrategy["canary"]; ok {
										if rSerialPipelineStagesStrategyCanary, ok := rSerialPipelineStagesStrategy["canary"].(map[string]interface{}); ok {
											rSerialPipelineStages.Strategy.Canary = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanary{}
											if _, ok := rSerialPipelineStagesStrategyCanary["canaryDeployment"]; ok {
												if rSerialPipelineStagesStrategyCanaryCanaryDeployment, ok := rSerialPipelineStagesStrategyCanary["canaryDeployment"].(map[string]interface{}); ok {
													rSerialPipelineStages.Strategy.Canary.CanaryDeployment = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{}
													if _, ok := rSerialPipelineStagesStrategyCanaryCanaryDeployment["percentages"]; ok {
														if s, ok := rSerialPipelineStagesStrategyCanaryCanaryDeployment["percentages"].([]interface{}); ok {
															for _, ss := range s {
																if intval, ok := ss.(int64); ok {
																	rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Percentages = append(rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Percentages, intval)
																}
															}
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Percentages: expected []interface{}")
														}
													}
													if _, ok := rSerialPipelineStagesStrategyCanaryCanaryDeployment["postdeploy"]; ok {
														if rSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy, ok := rSerialPipelineStagesStrategyCanaryCanaryDeployment["postdeploy"].(map[string]interface{}); ok {
															rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Postdeploy = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{}
															if _, ok := rSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy["actions"]; ok {
																if s, ok := rSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy["actions"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Postdeploy.Actions = append(rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Postdeploy.Actions, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Postdeploy.Actions: expected []interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Postdeploy: expected map[string]interface{}")
														}
													}
													if _, ok := rSerialPipelineStagesStrategyCanaryCanaryDeployment["predeploy"]; ok {
														if rSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy, ok := rSerialPipelineStagesStrategyCanaryCanaryDeployment["predeploy"].(map[string]interface{}); ok {
															rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Predeploy = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{}
															if _, ok := rSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy["actions"]; ok {
																if s, ok := rSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy["actions"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Predeploy.Actions = append(rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Predeploy.Actions, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Predeploy.Actions: expected []interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Predeploy: expected map[string]interface{}")
														}
													}
													if _, ok := rSerialPipelineStagesStrategyCanaryCanaryDeployment["verify"]; ok {
														if b, ok := rSerialPipelineStagesStrategyCanaryCanaryDeployment["verify"].(bool); ok {
															rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Verify = dcl.Bool(b)
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CanaryDeployment.Verify: expected bool")
														}
													}
												} else {
													return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CanaryDeployment: expected map[string]interface{}")
												}
											}
											if _, ok := rSerialPipelineStagesStrategyCanary["customCanaryDeployment"]; ok {
												if rSerialPipelineStagesStrategyCanaryCustomCanaryDeployment, ok := rSerialPipelineStagesStrategyCanary["customCanaryDeployment"].(map[string]interface{}); ok {
													rSerialPipelineStages.Strategy.Canary.CustomCanaryDeployment = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{}
													if _, ok := rSerialPipelineStagesStrategyCanaryCustomCanaryDeployment["phaseConfigs"]; ok {
														if s, ok := rSerialPipelineStagesStrategyCanaryCustomCanaryDeployment["phaseConfigs"].([]interface{}); ok {
															for _, o := range s {
																if objval, ok := o.(map[string]interface{}); ok {
																	var rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs
																	if _, ok := objval["percentage"]; ok {
																		if i, ok := objval["percentage"].(int64); ok {
																			rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Percentage = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Percentage: expected int64")
																		}
																	}
																	if _, ok := objval["phaseId"]; ok {
																		if s, ok := objval["phaseId"].(string); ok {
																			rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.PhaseId = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.PhaseId: expected string")
																		}
																	}
																	if _, ok := objval["postdeploy"]; ok {
																		if rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy, ok := objval["postdeploy"].(map[string]interface{}); ok {
																			rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Postdeploy = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{}
																			if _, ok := rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy["actions"]; ok {
																				if s, ok := rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy["actions"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Postdeploy.Actions = append(rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Postdeploy.Actions, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Postdeploy.Actions: expected []interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Postdeploy: expected map[string]interface{}")
																		}
																	}
																	if _, ok := objval["predeploy"]; ok {
																		if rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy, ok := objval["predeploy"].(map[string]interface{}); ok {
																			rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Predeploy = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{}
																			if _, ok := rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy["actions"]; ok {
																				if s, ok := rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy["actions"].([]interface{}); ok {
																					for _, ss := range s {
																						if strval, ok := ss.(string); ok {
																							rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Predeploy.Actions = append(rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Predeploy.Actions, strval)
																						}
																					}
																				} else {
																					return nil, fmt.Errorf("rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Predeploy.Actions: expected []interface{}")
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Predeploy: expected map[string]interface{}")
																		}
																	}
																	if _, ok := objval["profiles"]; ok {
																		if s, ok := objval["profiles"].([]interface{}); ok {
																			for _, ss := range s {
																				if strval, ok := ss.(string); ok {
																					rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Profiles = append(rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Profiles, strval)
																				}
																			}
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Profiles: expected []interface{}")
																		}
																	}
																	if _, ok := objval["verify"]; ok {
																		if b, ok := objval["verify"].(bool); ok {
																			rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Verify = dcl.Bool(b)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs.Verify: expected bool")
																		}
																	}
																	rSerialPipelineStages.Strategy.Canary.CustomCanaryDeployment.PhaseConfigs = append(rSerialPipelineStages.Strategy.Canary.CustomCanaryDeployment.PhaseConfigs, rSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs)
																}
															}
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CustomCanaryDeployment.PhaseConfigs: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.CustomCanaryDeployment: expected map[string]interface{}")
												}
											}
											if _, ok := rSerialPipelineStagesStrategyCanary["runtimeConfig"]; ok {
												if rSerialPipelineStagesStrategyCanaryRuntimeConfig, ok := rSerialPipelineStagesStrategyCanary["runtimeConfig"].(map[string]interface{}); ok {
													rSerialPipelineStages.Strategy.Canary.RuntimeConfig = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{}
													if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfig["cloudRun"]; ok {
														if rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfig["cloudRun"].(map[string]interface{}); ok {
															rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{}
															if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun["automaticTrafficControl"]; ok {
																if b, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun["automaticTrafficControl"].(bool); ok {
																	rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.AutomaticTrafficControl = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.AutomaticTrafficControl: expected bool")
																}
															}
															if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun["canaryRevisionTags"]; ok {
																if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun["canaryRevisionTags"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.CanaryRevisionTags = append(rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.CanaryRevisionTags, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.CanaryRevisionTags: expected []interface{}")
																}
															}
															if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun["priorRevisionTags"]; ok {
																if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun["priorRevisionTags"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.PriorRevisionTags = append(rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.PriorRevisionTags, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.PriorRevisionTags: expected []interface{}")
																}
															}
															if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun["stableRevisionTags"]; ok {
																if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun["stableRevisionTags"].([]interface{}); ok {
																	for _, ss := range s {
																		if strval, ok := ss.(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.StableRevisionTags = append(rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.StableRevisionTags, strval)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun.StableRevisionTags: expected []interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.CloudRun: expected map[string]interface{}")
														}
													}
													if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfig["kubernetes"]; ok {
														if rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfig["kubernetes"].(map[string]interface{}); ok {
															rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{}
															if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes["gatewayServiceMesh"]; ok {
																if rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes["gatewayServiceMesh"].(map[string]interface{}); ok {
																	rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{}
																	if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["deployment"]; ok {
																		if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["deployment"].(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.Deployment = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.Deployment: expected string")
																		}
																	}
																	if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["httpRoute"]; ok {
																		if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["httpRoute"].(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.HttpRoute = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.HttpRoute: expected string")
																		}
																	}
																	if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["routeUpdateWaitTime"]; ok {
																		if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["routeUpdateWaitTime"].(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.RouteUpdateWaitTime = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.RouteUpdateWaitTime: expected string")
																		}
																	}
																	if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["service"]; ok {
																		if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["service"].(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.Service = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.Service: expected string")
																		}
																	}
																	if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["stableCutbackDuration"]; ok {
																		if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh["stableCutbackDuration"].(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.StableCutbackDuration = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh.StableCutbackDuration: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.GatewayServiceMesh: expected map[string]interface{}")
																}
															}
															if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes["serviceNetworking"]; ok {
																if rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes["serviceNetworking"].(map[string]interface{}); ok {
																	rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking = &dclService.DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{}
																	if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking["deployment"]; ok {
																		if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking["deployment"].(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.Deployment = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.Deployment: expected string")
																		}
																	}
																	if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking["disablePodOverprovisioning"]; ok {
																		if b, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking["disablePodOverprovisioning"].(bool); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.DisablePodOverprovisioning = dcl.Bool(b)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.DisablePodOverprovisioning: expected bool")
																		}
																	}
																	if _, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking["service"]; ok {
																		if s, ok := rSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking["service"].(string); ok {
																			rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.Service = dcl.String(s)
																		} else {
																			return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking.Service: expected string")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes.ServiceNetworking: expected map[string]interface{}")
																}
															}
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig.Kubernetes: expected map[string]interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary.RuntimeConfig: expected map[string]interface{}")
												}
											}
										} else {
											return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Canary: expected map[string]interface{}")
										}
									}
									if _, ok := rSerialPipelineStagesStrategy["standard"]; ok {
										if rSerialPipelineStagesStrategyStandard, ok := rSerialPipelineStagesStrategy["standard"].(map[string]interface{}); ok {
											rSerialPipelineStages.Strategy.Standard = &dclService.DeliveryPipelineSerialPipelineStagesStrategyStandard{}
											if _, ok := rSerialPipelineStagesStrategyStandard["postdeploy"]; ok {
												if rSerialPipelineStagesStrategyStandardPostdeploy, ok := rSerialPipelineStagesStrategyStandard["postdeploy"].(map[string]interface{}); ok {
													rSerialPipelineStages.Strategy.Standard.Postdeploy = &dclService.DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{}
													if _, ok := rSerialPipelineStagesStrategyStandardPostdeploy["actions"]; ok {
														if s, ok := rSerialPipelineStagesStrategyStandardPostdeploy["actions"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rSerialPipelineStages.Strategy.Standard.Postdeploy.Actions = append(rSerialPipelineStages.Strategy.Standard.Postdeploy.Actions, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Standard.Postdeploy.Actions: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Standard.Postdeploy: expected map[string]interface{}")
												}
											}
											if _, ok := rSerialPipelineStagesStrategyStandard["predeploy"]; ok {
												if rSerialPipelineStagesStrategyStandardPredeploy, ok := rSerialPipelineStagesStrategyStandard["predeploy"].(map[string]interface{}); ok {
													rSerialPipelineStages.Strategy.Standard.Predeploy = &dclService.DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{}
													if _, ok := rSerialPipelineStagesStrategyStandardPredeploy["actions"]; ok {
														if s, ok := rSerialPipelineStagesStrategyStandardPredeploy["actions"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rSerialPipelineStages.Strategy.Standard.Predeploy.Actions = append(rSerialPipelineStages.Strategy.Standard.Predeploy.Actions, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Standard.Predeploy.Actions: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Standard.Predeploy: expected map[string]interface{}")
												}
											}
											if _, ok := rSerialPipelineStagesStrategyStandard["verify"]; ok {
												if b, ok := rSerialPipelineStagesStrategyStandard["verify"].(bool); ok {
													rSerialPipelineStages.Strategy.Standard.Verify = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Standard.Verify: expected bool")
												}
											}
										} else {
											return nil, fmt.Errorf("rSerialPipelineStages.Strategy.Standard: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rSerialPipelineStages.Strategy: expected map[string]interface{}")
								}
							}
							if _, ok := objval["targetId"]; ok {
								if s, ok := objval["targetId"].(string); ok {
									rSerialPipelineStages.TargetId = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rSerialPipelineStages.TargetId: expected string")
								}
							}
							r.SerialPipeline.Stages = append(r.SerialPipeline.Stages, rSerialPipelineStages)
						}
					}
				} else {
					return nil, fmt.Errorf("r.SerialPipeline.Stages: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SerialPipeline: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["suspended"]; ok {
		if b, ok := u.Object["suspended"].(bool); ok {
			r.Suspended = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Suspended: expected bool")
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

func GetDeliveryPipeline(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetDeliveryPipeline(ctx, r)
	if err != nil {
		return nil, err
	}
	return DeliveryPipelineToUnstructured(r), nil
}

func ListDeliveryPipeline(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListDeliveryPipeline(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, DeliveryPipelineToUnstructured(r))
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

func ApplyDeliveryPipeline(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDeliveryPipeline(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyDeliveryPipeline(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return DeliveryPipelineToUnstructured(r), nil
}

func DeliveryPipelineHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToDeliveryPipeline(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyDeliveryPipeline(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteDeliveryPipeline(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return err
	}
	return c.DeleteDeliveryPipeline(ctx, r)
}

func DeliveryPipelineID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToDeliveryPipeline(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *DeliveryPipeline) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"clouddeploy",
		"DeliveryPipeline",
		"beta",
	}
}

func (r *DeliveryPipeline) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *DeliveryPipeline) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetDeliveryPipeline(ctx, config, resource)
}

func (r *DeliveryPipeline) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyDeliveryPipeline(ctx, config, resource, opts...)
}

func (r *DeliveryPipeline) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return DeliveryPipelineHasDiff(ctx, config, resource, opts...)
}

func (r *DeliveryPipeline) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteDeliveryPipeline(ctx, config, resource)
}

func (r *DeliveryPipeline) ID(resource *unstructured.Resource) (string, error) {
	return DeliveryPipelineID(resource)
}

func init() {
	unstructured.Register(&DeliveryPipeline{})
}
