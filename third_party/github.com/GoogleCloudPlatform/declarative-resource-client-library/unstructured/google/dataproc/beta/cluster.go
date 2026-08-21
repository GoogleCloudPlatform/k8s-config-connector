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
package dataproc

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Cluster struct{}

func ClusterToUnstructured(r *dclService.Cluster) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dataproc",
			Version: "beta",
			Type:    "Cluster",
		},
		Object: make(map[string]interface{}),
	}
	if r.ClusterUuid != nil {
		u.Object["clusterUuid"] = *r.ClusterUuid
	}
	if r.Config != nil && r.Config != dclService.EmptyClusterConfig {
		rConfig := make(map[string]interface{})
		if r.Config.AutoscalingConfig != nil && r.Config.AutoscalingConfig != dclService.EmptyClusterConfigAutoscalingConfig {
			rConfigAutoscalingConfig := make(map[string]interface{})
			if r.Config.AutoscalingConfig.Policy != nil {
				rConfigAutoscalingConfig["policy"] = *r.Config.AutoscalingConfig.Policy
			}
			rConfig["autoscalingConfig"] = rConfigAutoscalingConfig
		}
		if r.Config.DataprocMetricConfig != nil && r.Config.DataprocMetricConfig != dclService.EmptyClusterConfigDataprocMetricConfig {
			rConfigDataprocMetricConfig := make(map[string]interface{})
			var rConfigDataprocMetricConfigMetrics []interface{}
			for _, rConfigDataprocMetricConfigMetricsVal := range r.Config.DataprocMetricConfig.Metrics {
				rConfigDataprocMetricConfigMetricsObject := make(map[string]interface{})
				var rConfigDataprocMetricConfigMetricsValMetricOverrides []interface{}
				for _, rConfigDataprocMetricConfigMetricsValMetricOverridesVal := range rConfigDataprocMetricConfigMetricsVal.MetricOverrides {
					rConfigDataprocMetricConfigMetricsValMetricOverrides = append(rConfigDataprocMetricConfigMetricsValMetricOverrides, rConfigDataprocMetricConfigMetricsValMetricOverridesVal)
				}
				rConfigDataprocMetricConfigMetricsObject["metricOverrides"] = rConfigDataprocMetricConfigMetricsValMetricOverrides
				if rConfigDataprocMetricConfigMetricsVal.MetricSource != nil {
					rConfigDataprocMetricConfigMetricsObject["metricSource"] = string(*rConfigDataprocMetricConfigMetricsVal.MetricSource)
				}
				rConfigDataprocMetricConfigMetrics = append(rConfigDataprocMetricConfigMetrics, rConfigDataprocMetricConfigMetricsObject)
			}
			rConfigDataprocMetricConfig["metrics"] = rConfigDataprocMetricConfigMetrics
			rConfig["dataprocMetricConfig"] = rConfigDataprocMetricConfig
		}
		if r.Config.EncryptionConfig != nil && r.Config.EncryptionConfig != dclService.EmptyClusterConfigEncryptionConfig {
			rConfigEncryptionConfig := make(map[string]interface{})
			if r.Config.EncryptionConfig.GcePdKmsKeyName != nil {
				rConfigEncryptionConfig["gcePdKmsKeyName"] = *r.Config.EncryptionConfig.GcePdKmsKeyName
			}
			rConfig["encryptionConfig"] = rConfigEncryptionConfig
		}
		if r.Config.EndpointConfig != nil && r.Config.EndpointConfig != dclService.EmptyClusterConfigEndpointConfig {
			rConfigEndpointConfig := make(map[string]interface{})
			if r.Config.EndpointConfig.EnableHttpPortAccess != nil {
				rConfigEndpointConfig["enableHttpPortAccess"] = *r.Config.EndpointConfig.EnableHttpPortAccess
			}
			if r.Config.EndpointConfig.HttpPorts != nil {
				rConfigEndpointConfigHttpPorts := make(map[string]interface{})
				for k, v := range r.Config.EndpointConfig.HttpPorts {
					rConfigEndpointConfigHttpPorts[k] = v
				}
				rConfigEndpointConfig["httpPorts"] = rConfigEndpointConfigHttpPorts
			}
			rConfig["endpointConfig"] = rConfigEndpointConfig
		}
		if r.Config.GceClusterConfig != nil && r.Config.GceClusterConfig != dclService.EmptyClusterConfigGceClusterConfig {
			rConfigGceClusterConfig := make(map[string]interface{})
			if r.Config.GceClusterConfig.ConfidentialInstanceConfig != nil && r.Config.GceClusterConfig.ConfidentialInstanceConfig != dclService.EmptyClusterConfigGceClusterConfigConfidentialInstanceConfig {
				rConfigGceClusterConfigConfidentialInstanceConfig := make(map[string]interface{})
				if r.Config.GceClusterConfig.ConfidentialInstanceConfig.EnableConfidentialCompute != nil {
					rConfigGceClusterConfigConfidentialInstanceConfig["enableConfidentialCompute"] = *r.Config.GceClusterConfig.ConfidentialInstanceConfig.EnableConfidentialCompute
				}
				rConfigGceClusterConfig["confidentialInstanceConfig"] = rConfigGceClusterConfigConfidentialInstanceConfig
			}
			if r.Config.GceClusterConfig.InternalIPOnly != nil {
				rConfigGceClusterConfig["internalIPOnly"] = *r.Config.GceClusterConfig.InternalIPOnly
			}
			if r.Config.GceClusterConfig.Metadata != nil {
				rConfigGceClusterConfigMetadata := make(map[string]interface{})
				for k, v := range r.Config.GceClusterConfig.Metadata {
					rConfigGceClusterConfigMetadata[k] = v
				}
				rConfigGceClusterConfig["metadata"] = rConfigGceClusterConfigMetadata
			}
			if r.Config.GceClusterConfig.Network != nil {
				rConfigGceClusterConfig["network"] = *r.Config.GceClusterConfig.Network
			}
			if r.Config.GceClusterConfig.NodeGroupAffinity != nil && r.Config.GceClusterConfig.NodeGroupAffinity != dclService.EmptyClusterConfigGceClusterConfigNodeGroupAffinity {
				rConfigGceClusterConfigNodeGroupAffinity := make(map[string]interface{})
				if r.Config.GceClusterConfig.NodeGroupAffinity.NodeGroup != nil {
					rConfigGceClusterConfigNodeGroupAffinity["nodeGroup"] = *r.Config.GceClusterConfig.NodeGroupAffinity.NodeGroup
				}
				rConfigGceClusterConfig["nodeGroupAffinity"] = rConfigGceClusterConfigNodeGroupAffinity
			}
			if r.Config.GceClusterConfig.PrivateIPv6GoogleAccess != nil {
				rConfigGceClusterConfig["privateIPv6GoogleAccess"] = string(*r.Config.GceClusterConfig.PrivateIPv6GoogleAccess)
			}
			if r.Config.GceClusterConfig.ReservationAffinity != nil && r.Config.GceClusterConfig.ReservationAffinity != dclService.EmptyClusterConfigGceClusterConfigReservationAffinity {
				rConfigGceClusterConfigReservationAffinity := make(map[string]interface{})
				if r.Config.GceClusterConfig.ReservationAffinity.ConsumeReservationType != nil {
					rConfigGceClusterConfigReservationAffinity["consumeReservationType"] = string(*r.Config.GceClusterConfig.ReservationAffinity.ConsumeReservationType)
				}
				if r.Config.GceClusterConfig.ReservationAffinity.Key != nil {
					rConfigGceClusterConfigReservationAffinity["key"] = *r.Config.GceClusterConfig.ReservationAffinity.Key
				}
				var rConfigGceClusterConfigReservationAffinityValues []interface{}
				for _, rConfigGceClusterConfigReservationAffinityValuesVal := range r.Config.GceClusterConfig.ReservationAffinity.Values {
					rConfigGceClusterConfigReservationAffinityValues = append(rConfigGceClusterConfigReservationAffinityValues, rConfigGceClusterConfigReservationAffinityValuesVal)
				}
				rConfigGceClusterConfigReservationAffinity["values"] = rConfigGceClusterConfigReservationAffinityValues
				rConfigGceClusterConfig["reservationAffinity"] = rConfigGceClusterConfigReservationAffinity
			}
			if r.Config.GceClusterConfig.ServiceAccount != nil {
				rConfigGceClusterConfig["serviceAccount"] = *r.Config.GceClusterConfig.ServiceAccount
			}
			var rConfigGceClusterConfigServiceAccountScopes []interface{}
			for _, rConfigGceClusterConfigServiceAccountScopesVal := range r.Config.GceClusterConfig.ServiceAccountScopes {
				rConfigGceClusterConfigServiceAccountScopes = append(rConfigGceClusterConfigServiceAccountScopes, rConfigGceClusterConfigServiceAccountScopesVal)
			}
			rConfigGceClusterConfig["serviceAccountScopes"] = rConfigGceClusterConfigServiceAccountScopes
			if r.Config.GceClusterConfig.ShieldedInstanceConfig != nil && r.Config.GceClusterConfig.ShieldedInstanceConfig != dclService.EmptyClusterConfigGceClusterConfigShieldedInstanceConfig {
				rConfigGceClusterConfigShieldedInstanceConfig := make(map[string]interface{})
				if r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring != nil {
					rConfigGceClusterConfigShieldedInstanceConfig["enableIntegrityMonitoring"] = *r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring
				}
				if r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot != nil {
					rConfigGceClusterConfigShieldedInstanceConfig["enableSecureBoot"] = *r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot
				}
				if r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm != nil {
					rConfigGceClusterConfigShieldedInstanceConfig["enableVtpm"] = *r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm
				}
				rConfigGceClusterConfig["shieldedInstanceConfig"] = rConfigGceClusterConfigShieldedInstanceConfig
			}
			if r.Config.GceClusterConfig.Subnetwork != nil {
				rConfigGceClusterConfig["subnetwork"] = *r.Config.GceClusterConfig.Subnetwork
			}
			var rConfigGceClusterConfigTags []interface{}
			for _, rConfigGceClusterConfigTagsVal := range r.Config.GceClusterConfig.Tags {
				rConfigGceClusterConfigTags = append(rConfigGceClusterConfigTags, rConfigGceClusterConfigTagsVal)
			}
			rConfigGceClusterConfig["tags"] = rConfigGceClusterConfigTags
			if r.Config.GceClusterConfig.Zone != nil {
				rConfigGceClusterConfig["zone"] = *r.Config.GceClusterConfig.Zone
			}
			rConfig["gceClusterConfig"] = rConfigGceClusterConfig
		}
		if r.Config.GkeClusterConfig != nil && r.Config.GkeClusterConfig != dclService.EmptyClusterConfigGkeClusterConfig {
			rConfigGkeClusterConfig := make(map[string]interface{})
			if r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget != nil && r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget != dclService.EmptyClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
				rConfigGkeClusterConfigNamespacedGkeDeploymentTarget := make(map[string]interface{})
				if r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace != nil {
					rConfigGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"] = *r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace
				}
				if r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster != nil {
					rConfigGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"] = *r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster
				}
				rConfigGkeClusterConfig["namespacedGkeDeploymentTarget"] = rConfigGkeClusterConfigNamespacedGkeDeploymentTarget
			}
			rConfig["gkeClusterConfig"] = rConfigGkeClusterConfig
		}
		var rConfigInitializationActions []interface{}
		for _, rConfigInitializationActionsVal := range r.Config.InitializationActions {
			rConfigInitializationActionsObject := make(map[string]interface{})
			if rConfigInitializationActionsVal.ExecutableFile != nil {
				rConfigInitializationActionsObject["executableFile"] = *rConfigInitializationActionsVal.ExecutableFile
			}
			if rConfigInitializationActionsVal.ExecutionTimeout != nil {
				rConfigInitializationActionsObject["executionTimeout"] = *rConfigInitializationActionsVal.ExecutionTimeout
			}
			rConfigInitializationActions = append(rConfigInitializationActions, rConfigInitializationActionsObject)
		}
		rConfig["initializationActions"] = rConfigInitializationActions
		if r.Config.LifecycleConfig != nil && r.Config.LifecycleConfig != dclService.EmptyClusterConfigLifecycleConfig {
			rConfigLifecycleConfig := make(map[string]interface{})
			if r.Config.LifecycleConfig.AutoDeleteTime != nil {
				rConfigLifecycleConfig["autoDeleteTime"] = *r.Config.LifecycleConfig.AutoDeleteTime
			}
			if r.Config.LifecycleConfig.AutoDeleteTtl != nil {
				rConfigLifecycleConfig["autoDeleteTtl"] = *r.Config.LifecycleConfig.AutoDeleteTtl
			}
			if r.Config.LifecycleConfig.IdleDeleteTtl != nil {
				rConfigLifecycleConfig["idleDeleteTtl"] = *r.Config.LifecycleConfig.IdleDeleteTtl
			}
			if r.Config.LifecycleConfig.IdleStartTime != nil {
				rConfigLifecycleConfig["idleStartTime"] = *r.Config.LifecycleConfig.IdleStartTime
			}
			rConfig["lifecycleConfig"] = rConfigLifecycleConfig
		}
		if r.Config.MasterConfig != nil && r.Config.MasterConfig != dclService.EmptyClusterConfigMasterConfig {
			rConfigMasterConfig := make(map[string]interface{})
			var rConfigMasterConfigAccelerators []interface{}
			for _, rConfigMasterConfigAcceleratorsVal := range r.Config.MasterConfig.Accelerators {
				rConfigMasterConfigAcceleratorsObject := make(map[string]interface{})
				if rConfigMasterConfigAcceleratorsVal.AcceleratorCount != nil {
					rConfigMasterConfigAcceleratorsObject["acceleratorCount"] = *rConfigMasterConfigAcceleratorsVal.AcceleratorCount
				}
				if rConfigMasterConfigAcceleratorsVal.AcceleratorType != nil {
					rConfigMasterConfigAcceleratorsObject["acceleratorType"] = *rConfigMasterConfigAcceleratorsVal.AcceleratorType
				}
				rConfigMasterConfigAccelerators = append(rConfigMasterConfigAccelerators, rConfigMasterConfigAcceleratorsObject)
			}
			rConfigMasterConfig["accelerators"] = rConfigMasterConfigAccelerators
			if r.Config.MasterConfig.DiskConfig != nil && r.Config.MasterConfig.DiskConfig != dclService.EmptyClusterConfigMasterConfigDiskConfig {
				rConfigMasterConfigDiskConfig := make(map[string]interface{})
				if r.Config.MasterConfig.DiskConfig.BootDiskSizeGb != nil {
					rConfigMasterConfigDiskConfig["bootDiskSizeGb"] = *r.Config.MasterConfig.DiskConfig.BootDiskSizeGb
				}
				if r.Config.MasterConfig.DiskConfig.BootDiskType != nil {
					rConfigMasterConfigDiskConfig["bootDiskType"] = *r.Config.MasterConfig.DiskConfig.BootDiskType
				}
				if r.Config.MasterConfig.DiskConfig.LocalSsdInterface != nil {
					rConfigMasterConfigDiskConfig["localSsdInterface"] = *r.Config.MasterConfig.DiskConfig.LocalSsdInterface
				}
				if r.Config.MasterConfig.DiskConfig.NumLocalSsds != nil {
					rConfigMasterConfigDiskConfig["numLocalSsds"] = *r.Config.MasterConfig.DiskConfig.NumLocalSsds
				}
				rConfigMasterConfig["diskConfig"] = rConfigMasterConfigDiskConfig
			}
			if r.Config.MasterConfig.Image != nil {
				rConfigMasterConfig["image"] = *r.Config.MasterConfig.Image
			}
			var rConfigMasterConfigInstanceNames []interface{}
			for _, rConfigMasterConfigInstanceNamesVal := range r.Config.MasterConfig.InstanceNames {
				rConfigMasterConfigInstanceNames = append(rConfigMasterConfigInstanceNames, rConfigMasterConfigInstanceNamesVal)
			}
			rConfigMasterConfig["instanceNames"] = rConfigMasterConfigInstanceNames
			var rConfigMasterConfigInstanceReferences []interface{}
			for _, rConfigMasterConfigInstanceReferencesVal := range r.Config.MasterConfig.InstanceReferences {
				rConfigMasterConfigInstanceReferencesObject := make(map[string]interface{})
				if rConfigMasterConfigInstanceReferencesVal.InstanceId != nil {
					rConfigMasterConfigInstanceReferencesObject["instanceId"] = *rConfigMasterConfigInstanceReferencesVal.InstanceId
				}
				if rConfigMasterConfigInstanceReferencesVal.InstanceName != nil {
					rConfigMasterConfigInstanceReferencesObject["instanceName"] = *rConfigMasterConfigInstanceReferencesVal.InstanceName
				}
				if rConfigMasterConfigInstanceReferencesVal.PublicEciesKey != nil {
					rConfigMasterConfigInstanceReferencesObject["publicEciesKey"] = *rConfigMasterConfigInstanceReferencesVal.PublicEciesKey
				}
				if rConfigMasterConfigInstanceReferencesVal.PublicKey != nil {
					rConfigMasterConfigInstanceReferencesObject["publicKey"] = *rConfigMasterConfigInstanceReferencesVal.PublicKey
				}
				rConfigMasterConfigInstanceReferences = append(rConfigMasterConfigInstanceReferences, rConfigMasterConfigInstanceReferencesObject)
			}
			rConfigMasterConfig["instanceReferences"] = rConfigMasterConfigInstanceReferences
			if r.Config.MasterConfig.IsPreemptible != nil {
				rConfigMasterConfig["isPreemptible"] = *r.Config.MasterConfig.IsPreemptible
			}
			if r.Config.MasterConfig.MachineType != nil {
				rConfigMasterConfig["machineType"] = *r.Config.MasterConfig.MachineType
			}
			if r.Config.MasterConfig.ManagedGroupConfig != nil && r.Config.MasterConfig.ManagedGroupConfig != dclService.EmptyClusterConfigMasterConfigManagedGroupConfig {
				rConfigMasterConfigManagedGroupConfig := make(map[string]interface{})
				if r.Config.MasterConfig.ManagedGroupConfig.InstanceGroupManagerName != nil {
					rConfigMasterConfigManagedGroupConfig["instanceGroupManagerName"] = *r.Config.MasterConfig.ManagedGroupConfig.InstanceGroupManagerName
				}
				if r.Config.MasterConfig.ManagedGroupConfig.InstanceTemplateName != nil {
					rConfigMasterConfigManagedGroupConfig["instanceTemplateName"] = *r.Config.MasterConfig.ManagedGroupConfig.InstanceTemplateName
				}
				rConfigMasterConfig["managedGroupConfig"] = rConfigMasterConfigManagedGroupConfig
			}
			if r.Config.MasterConfig.MinCpuPlatform != nil {
				rConfigMasterConfig["minCpuPlatform"] = *r.Config.MasterConfig.MinCpuPlatform
			}
			if r.Config.MasterConfig.NumInstances != nil {
				rConfigMasterConfig["numInstances"] = *r.Config.MasterConfig.NumInstances
			}
			if r.Config.MasterConfig.Preemptibility != nil {
				rConfigMasterConfig["preemptibility"] = string(*r.Config.MasterConfig.Preemptibility)
			}
			rConfig["masterConfig"] = rConfigMasterConfig
		}
		if r.Config.MetastoreConfig != nil && r.Config.MetastoreConfig != dclService.EmptyClusterConfigMetastoreConfig {
			rConfigMetastoreConfig := make(map[string]interface{})
			if r.Config.MetastoreConfig.DataprocMetastoreService != nil {
				rConfigMetastoreConfig["dataprocMetastoreService"] = *r.Config.MetastoreConfig.DataprocMetastoreService
			}
			rConfig["metastoreConfig"] = rConfigMetastoreConfig
		}
		if r.Config.SecondaryWorkerConfig != nil && r.Config.SecondaryWorkerConfig != dclService.EmptyClusterConfigSecondaryWorkerConfig {
			rConfigSecondaryWorkerConfig := make(map[string]interface{})
			var rConfigSecondaryWorkerConfigAccelerators []interface{}
			for _, rConfigSecondaryWorkerConfigAcceleratorsVal := range r.Config.SecondaryWorkerConfig.Accelerators {
				rConfigSecondaryWorkerConfigAcceleratorsObject := make(map[string]interface{})
				if rConfigSecondaryWorkerConfigAcceleratorsVal.AcceleratorCount != nil {
					rConfigSecondaryWorkerConfigAcceleratorsObject["acceleratorCount"] = *rConfigSecondaryWorkerConfigAcceleratorsVal.AcceleratorCount
				}
				if rConfigSecondaryWorkerConfigAcceleratorsVal.AcceleratorType != nil {
					rConfigSecondaryWorkerConfigAcceleratorsObject["acceleratorType"] = *rConfigSecondaryWorkerConfigAcceleratorsVal.AcceleratorType
				}
				rConfigSecondaryWorkerConfigAccelerators = append(rConfigSecondaryWorkerConfigAccelerators, rConfigSecondaryWorkerConfigAcceleratorsObject)
			}
			rConfigSecondaryWorkerConfig["accelerators"] = rConfigSecondaryWorkerConfigAccelerators
			if r.Config.SecondaryWorkerConfig.DiskConfig != nil && r.Config.SecondaryWorkerConfig.DiskConfig != dclService.EmptyClusterConfigSecondaryWorkerConfigDiskConfig {
				rConfigSecondaryWorkerConfigDiskConfig := make(map[string]interface{})
				if r.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb != nil {
					rConfigSecondaryWorkerConfigDiskConfig["bootDiskSizeGb"] = *r.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb
				}
				if r.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType != nil {
					rConfigSecondaryWorkerConfigDiskConfig["bootDiskType"] = *r.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType
				}
				if r.Config.SecondaryWorkerConfig.DiskConfig.LocalSsdInterface != nil {
					rConfigSecondaryWorkerConfigDiskConfig["localSsdInterface"] = *r.Config.SecondaryWorkerConfig.DiskConfig.LocalSsdInterface
				}
				if r.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds != nil {
					rConfigSecondaryWorkerConfigDiskConfig["numLocalSsds"] = *r.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds
				}
				rConfigSecondaryWorkerConfig["diskConfig"] = rConfigSecondaryWorkerConfigDiskConfig
			}
			if r.Config.SecondaryWorkerConfig.Image != nil {
				rConfigSecondaryWorkerConfig["image"] = *r.Config.SecondaryWorkerConfig.Image
			}
			var rConfigSecondaryWorkerConfigInstanceNames []interface{}
			for _, rConfigSecondaryWorkerConfigInstanceNamesVal := range r.Config.SecondaryWorkerConfig.InstanceNames {
				rConfigSecondaryWorkerConfigInstanceNames = append(rConfigSecondaryWorkerConfigInstanceNames, rConfigSecondaryWorkerConfigInstanceNamesVal)
			}
			rConfigSecondaryWorkerConfig["instanceNames"] = rConfigSecondaryWorkerConfigInstanceNames
			var rConfigSecondaryWorkerConfigInstanceReferences []interface{}
			for _, rConfigSecondaryWorkerConfigInstanceReferencesVal := range r.Config.SecondaryWorkerConfig.InstanceReferences {
				rConfigSecondaryWorkerConfigInstanceReferencesObject := make(map[string]interface{})
				if rConfigSecondaryWorkerConfigInstanceReferencesVal.InstanceId != nil {
					rConfigSecondaryWorkerConfigInstanceReferencesObject["instanceId"] = *rConfigSecondaryWorkerConfigInstanceReferencesVal.InstanceId
				}
				if rConfigSecondaryWorkerConfigInstanceReferencesVal.InstanceName != nil {
					rConfigSecondaryWorkerConfigInstanceReferencesObject["instanceName"] = *rConfigSecondaryWorkerConfigInstanceReferencesVal.InstanceName
				}
				if rConfigSecondaryWorkerConfigInstanceReferencesVal.PublicEciesKey != nil {
					rConfigSecondaryWorkerConfigInstanceReferencesObject["publicEciesKey"] = *rConfigSecondaryWorkerConfigInstanceReferencesVal.PublicEciesKey
				}
				if rConfigSecondaryWorkerConfigInstanceReferencesVal.PublicKey != nil {
					rConfigSecondaryWorkerConfigInstanceReferencesObject["publicKey"] = *rConfigSecondaryWorkerConfigInstanceReferencesVal.PublicKey
				}
				rConfigSecondaryWorkerConfigInstanceReferences = append(rConfigSecondaryWorkerConfigInstanceReferences, rConfigSecondaryWorkerConfigInstanceReferencesObject)
			}
			rConfigSecondaryWorkerConfig["instanceReferences"] = rConfigSecondaryWorkerConfigInstanceReferences
			if r.Config.SecondaryWorkerConfig.IsPreemptible != nil {
				rConfigSecondaryWorkerConfig["isPreemptible"] = *r.Config.SecondaryWorkerConfig.IsPreemptible
			}
			if r.Config.SecondaryWorkerConfig.MachineType != nil {
				rConfigSecondaryWorkerConfig["machineType"] = *r.Config.SecondaryWorkerConfig.MachineType
			}
			if r.Config.SecondaryWorkerConfig.ManagedGroupConfig != nil && r.Config.SecondaryWorkerConfig.ManagedGroupConfig != dclService.EmptyClusterConfigSecondaryWorkerConfigManagedGroupConfig {
				rConfigSecondaryWorkerConfigManagedGroupConfig := make(map[string]interface{})
				if r.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceGroupManagerName != nil {
					rConfigSecondaryWorkerConfigManagedGroupConfig["instanceGroupManagerName"] = *r.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceGroupManagerName
				}
				if r.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceTemplateName != nil {
					rConfigSecondaryWorkerConfigManagedGroupConfig["instanceTemplateName"] = *r.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceTemplateName
				}
				rConfigSecondaryWorkerConfig["managedGroupConfig"] = rConfigSecondaryWorkerConfigManagedGroupConfig
			}
			if r.Config.SecondaryWorkerConfig.MinCpuPlatform != nil {
				rConfigSecondaryWorkerConfig["minCpuPlatform"] = *r.Config.SecondaryWorkerConfig.MinCpuPlatform
			}
			if r.Config.SecondaryWorkerConfig.NumInstances != nil {
				rConfigSecondaryWorkerConfig["numInstances"] = *r.Config.SecondaryWorkerConfig.NumInstances
			}
			if r.Config.SecondaryWorkerConfig.Preemptibility != nil {
				rConfigSecondaryWorkerConfig["preemptibility"] = string(*r.Config.SecondaryWorkerConfig.Preemptibility)
			}
			rConfig["secondaryWorkerConfig"] = rConfigSecondaryWorkerConfig
		}
		if r.Config.SecurityConfig != nil && r.Config.SecurityConfig != dclService.EmptyClusterConfigSecurityConfig {
			rConfigSecurityConfig := make(map[string]interface{})
			if r.Config.SecurityConfig.IdentityConfig != nil && r.Config.SecurityConfig.IdentityConfig != dclService.EmptyClusterConfigSecurityConfigIdentityConfig {
				rConfigSecurityConfigIdentityConfig := make(map[string]interface{})
				if r.Config.SecurityConfig.IdentityConfig.UserServiceAccountMapping != nil {
					rConfigSecurityConfigIdentityConfigUserServiceAccountMapping := make(map[string]interface{})
					for k, v := range r.Config.SecurityConfig.IdentityConfig.UserServiceAccountMapping {
						rConfigSecurityConfigIdentityConfigUserServiceAccountMapping[k] = v
					}
					rConfigSecurityConfigIdentityConfig["userServiceAccountMapping"] = rConfigSecurityConfigIdentityConfigUserServiceAccountMapping
				}
				rConfigSecurityConfig["identityConfig"] = rConfigSecurityConfigIdentityConfig
			}
			if r.Config.SecurityConfig.KerberosConfig != nil && r.Config.SecurityConfig.KerberosConfig != dclService.EmptyClusterConfigSecurityConfigKerberosConfig {
				rConfigSecurityConfigKerberosConfig := make(map[string]interface{})
				if r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer != nil {
					rConfigSecurityConfigKerberosConfig["crossRealmTrustAdminServer"] = *r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer
				}
				if r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustKdc != nil {
					rConfigSecurityConfigKerberosConfig["crossRealmTrustKdc"] = *r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustKdc
				}
				if r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustRealm != nil {
					rConfigSecurityConfigKerberosConfig["crossRealmTrustRealm"] = *r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustRealm
				}
				if r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword != nil {
					rConfigSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"] = *r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword
				}
				if r.Config.SecurityConfig.KerberosConfig.EnableKerberos != nil {
					rConfigSecurityConfigKerberosConfig["enableKerberos"] = *r.Config.SecurityConfig.KerberosConfig.EnableKerberos
				}
				if r.Config.SecurityConfig.KerberosConfig.KdcDbKey != nil {
					rConfigSecurityConfigKerberosConfig["kdcDbKey"] = *r.Config.SecurityConfig.KerberosConfig.KdcDbKey
				}
				if r.Config.SecurityConfig.KerberosConfig.KeyPassword != nil {
					rConfigSecurityConfigKerberosConfig["keyPassword"] = *r.Config.SecurityConfig.KerberosConfig.KeyPassword
				}
				if r.Config.SecurityConfig.KerberosConfig.Keystore != nil {
					rConfigSecurityConfigKerberosConfig["keystore"] = *r.Config.SecurityConfig.KerberosConfig.Keystore
				}
				if r.Config.SecurityConfig.KerberosConfig.KeystorePassword != nil {
					rConfigSecurityConfigKerberosConfig["keystorePassword"] = *r.Config.SecurityConfig.KerberosConfig.KeystorePassword
				}
				if r.Config.SecurityConfig.KerberosConfig.KmsKey != nil {
					rConfigSecurityConfigKerberosConfig["kmsKey"] = *r.Config.SecurityConfig.KerberosConfig.KmsKey
				}
				if r.Config.SecurityConfig.KerberosConfig.Realm != nil {
					rConfigSecurityConfigKerberosConfig["realm"] = *r.Config.SecurityConfig.KerberosConfig.Realm
				}
				if r.Config.SecurityConfig.KerberosConfig.RootPrincipalPassword != nil {
					rConfigSecurityConfigKerberosConfig["rootPrincipalPassword"] = *r.Config.SecurityConfig.KerberosConfig.RootPrincipalPassword
				}
				if r.Config.SecurityConfig.KerberosConfig.TgtLifetimeHours != nil {
					rConfigSecurityConfigKerberosConfig["tgtLifetimeHours"] = *r.Config.SecurityConfig.KerberosConfig.TgtLifetimeHours
				}
				if r.Config.SecurityConfig.KerberosConfig.Truststore != nil {
					rConfigSecurityConfigKerberosConfig["truststore"] = *r.Config.SecurityConfig.KerberosConfig.Truststore
				}
				if r.Config.SecurityConfig.KerberosConfig.TruststorePassword != nil {
					rConfigSecurityConfigKerberosConfig["truststorePassword"] = *r.Config.SecurityConfig.KerberosConfig.TruststorePassword
				}
				rConfigSecurityConfig["kerberosConfig"] = rConfigSecurityConfigKerberosConfig
			}
			rConfig["securityConfig"] = rConfigSecurityConfig
		}
		if r.Config.SoftwareConfig != nil && r.Config.SoftwareConfig != dclService.EmptyClusterConfigSoftwareConfig {
			rConfigSoftwareConfig := make(map[string]interface{})
			if r.Config.SoftwareConfig.ImageVersion != nil {
				rConfigSoftwareConfig["imageVersion"] = *r.Config.SoftwareConfig.ImageVersion
			}
			var rConfigSoftwareConfigOptionalComponents []interface{}
			for _, rConfigSoftwareConfigOptionalComponentsVal := range r.Config.SoftwareConfig.OptionalComponents {
				rConfigSoftwareConfigOptionalComponents = append(rConfigSoftwareConfigOptionalComponents, string(rConfigSoftwareConfigOptionalComponentsVal))
			}
			rConfigSoftwareConfig["optionalComponents"] = rConfigSoftwareConfigOptionalComponents
			if r.Config.SoftwareConfig.Properties != nil {
				rConfigSoftwareConfigProperties := make(map[string]interface{})
				for k, v := range r.Config.SoftwareConfig.Properties {
					rConfigSoftwareConfigProperties[k] = v
				}
				rConfigSoftwareConfig["properties"] = rConfigSoftwareConfigProperties
			}
			rConfig["softwareConfig"] = rConfigSoftwareConfig
		}
		if r.Config.StagingBucket != nil {
			rConfig["stagingBucket"] = *r.Config.StagingBucket
		}
		if r.Config.TempBucket != nil {
			rConfig["tempBucket"] = *r.Config.TempBucket
		}
		if r.Config.WorkerConfig != nil && r.Config.WorkerConfig != dclService.EmptyClusterConfigWorkerConfig {
			rConfigWorkerConfig := make(map[string]interface{})
			var rConfigWorkerConfigAccelerators []interface{}
			for _, rConfigWorkerConfigAcceleratorsVal := range r.Config.WorkerConfig.Accelerators {
				rConfigWorkerConfigAcceleratorsObject := make(map[string]interface{})
				if rConfigWorkerConfigAcceleratorsVal.AcceleratorCount != nil {
					rConfigWorkerConfigAcceleratorsObject["acceleratorCount"] = *rConfigWorkerConfigAcceleratorsVal.AcceleratorCount
				}
				if rConfigWorkerConfigAcceleratorsVal.AcceleratorType != nil {
					rConfigWorkerConfigAcceleratorsObject["acceleratorType"] = *rConfigWorkerConfigAcceleratorsVal.AcceleratorType
				}
				rConfigWorkerConfigAccelerators = append(rConfigWorkerConfigAccelerators, rConfigWorkerConfigAcceleratorsObject)
			}
			rConfigWorkerConfig["accelerators"] = rConfigWorkerConfigAccelerators
			if r.Config.WorkerConfig.DiskConfig != nil && r.Config.WorkerConfig.DiskConfig != dclService.EmptyClusterConfigWorkerConfigDiskConfig {
				rConfigWorkerConfigDiskConfig := make(map[string]interface{})
				if r.Config.WorkerConfig.DiskConfig.BootDiskSizeGb != nil {
					rConfigWorkerConfigDiskConfig["bootDiskSizeGb"] = *r.Config.WorkerConfig.DiskConfig.BootDiskSizeGb
				}
				if r.Config.WorkerConfig.DiskConfig.BootDiskType != nil {
					rConfigWorkerConfigDiskConfig["bootDiskType"] = *r.Config.WorkerConfig.DiskConfig.BootDiskType
				}
				if r.Config.WorkerConfig.DiskConfig.LocalSsdInterface != nil {
					rConfigWorkerConfigDiskConfig["localSsdInterface"] = *r.Config.WorkerConfig.DiskConfig.LocalSsdInterface
				}
				if r.Config.WorkerConfig.DiskConfig.NumLocalSsds != nil {
					rConfigWorkerConfigDiskConfig["numLocalSsds"] = *r.Config.WorkerConfig.DiskConfig.NumLocalSsds
				}
				rConfigWorkerConfig["diskConfig"] = rConfigWorkerConfigDiskConfig
			}
			if r.Config.WorkerConfig.Image != nil {
				rConfigWorkerConfig["image"] = *r.Config.WorkerConfig.Image
			}
			var rConfigWorkerConfigInstanceNames []interface{}
			for _, rConfigWorkerConfigInstanceNamesVal := range r.Config.WorkerConfig.InstanceNames {
				rConfigWorkerConfigInstanceNames = append(rConfigWorkerConfigInstanceNames, rConfigWorkerConfigInstanceNamesVal)
			}
			rConfigWorkerConfig["instanceNames"] = rConfigWorkerConfigInstanceNames
			var rConfigWorkerConfigInstanceReferences []interface{}
			for _, rConfigWorkerConfigInstanceReferencesVal := range r.Config.WorkerConfig.InstanceReferences {
				rConfigWorkerConfigInstanceReferencesObject := make(map[string]interface{})
				if rConfigWorkerConfigInstanceReferencesVal.InstanceId != nil {
					rConfigWorkerConfigInstanceReferencesObject["instanceId"] = *rConfigWorkerConfigInstanceReferencesVal.InstanceId
				}
				if rConfigWorkerConfigInstanceReferencesVal.InstanceName != nil {
					rConfigWorkerConfigInstanceReferencesObject["instanceName"] = *rConfigWorkerConfigInstanceReferencesVal.InstanceName
				}
				if rConfigWorkerConfigInstanceReferencesVal.PublicEciesKey != nil {
					rConfigWorkerConfigInstanceReferencesObject["publicEciesKey"] = *rConfigWorkerConfigInstanceReferencesVal.PublicEciesKey
				}
				if rConfigWorkerConfigInstanceReferencesVal.PublicKey != nil {
					rConfigWorkerConfigInstanceReferencesObject["publicKey"] = *rConfigWorkerConfigInstanceReferencesVal.PublicKey
				}
				rConfigWorkerConfigInstanceReferences = append(rConfigWorkerConfigInstanceReferences, rConfigWorkerConfigInstanceReferencesObject)
			}
			rConfigWorkerConfig["instanceReferences"] = rConfigWorkerConfigInstanceReferences
			if r.Config.WorkerConfig.IsPreemptible != nil {
				rConfigWorkerConfig["isPreemptible"] = *r.Config.WorkerConfig.IsPreemptible
			}
			if r.Config.WorkerConfig.MachineType != nil {
				rConfigWorkerConfig["machineType"] = *r.Config.WorkerConfig.MachineType
			}
			if r.Config.WorkerConfig.ManagedGroupConfig != nil && r.Config.WorkerConfig.ManagedGroupConfig != dclService.EmptyClusterConfigWorkerConfigManagedGroupConfig {
				rConfigWorkerConfigManagedGroupConfig := make(map[string]interface{})
				if r.Config.WorkerConfig.ManagedGroupConfig.InstanceGroupManagerName != nil {
					rConfigWorkerConfigManagedGroupConfig["instanceGroupManagerName"] = *r.Config.WorkerConfig.ManagedGroupConfig.InstanceGroupManagerName
				}
				if r.Config.WorkerConfig.ManagedGroupConfig.InstanceTemplateName != nil {
					rConfigWorkerConfigManagedGroupConfig["instanceTemplateName"] = *r.Config.WorkerConfig.ManagedGroupConfig.InstanceTemplateName
				}
				rConfigWorkerConfig["managedGroupConfig"] = rConfigWorkerConfigManagedGroupConfig
			}
			if r.Config.WorkerConfig.MinCpuPlatform != nil {
				rConfigWorkerConfig["minCpuPlatform"] = *r.Config.WorkerConfig.MinCpuPlatform
			}
			if r.Config.WorkerConfig.NumInstances != nil {
				rConfigWorkerConfig["numInstances"] = *r.Config.WorkerConfig.NumInstances
			}
			if r.Config.WorkerConfig.Preemptibility != nil {
				rConfigWorkerConfig["preemptibility"] = string(*r.Config.WorkerConfig.Preemptibility)
			}
			rConfig["workerConfig"] = rConfigWorkerConfig
		}
		u.Object["config"] = rConfig
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
	if r.Metrics != nil && r.Metrics != dclService.EmptyClusterMetrics {
		rMetrics := make(map[string]interface{})
		if r.Metrics.HdfsMetrics != nil {
			rMetricsHdfsMetrics := make(map[string]interface{})
			for k, v := range r.Metrics.HdfsMetrics {
				rMetricsHdfsMetrics[k] = v
			}
			rMetrics["hdfsMetrics"] = rMetricsHdfsMetrics
		}
		if r.Metrics.YarnMetrics != nil {
			rMetricsYarnMetrics := make(map[string]interface{})
			for k, v := range r.Metrics.YarnMetrics {
				rMetricsYarnMetrics[k] = v
			}
			rMetrics["yarnMetrics"] = rMetricsYarnMetrics
		}
		u.Object["metrics"] = rMetrics
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Status != nil && r.Status != dclService.EmptyClusterStatus {
		rStatus := make(map[string]interface{})
		if r.Status.Detail != nil {
			rStatus["detail"] = *r.Status.Detail
		}
		if r.Status.State != nil {
			rStatus["state"] = string(*r.Status.State)
		}
		if r.Status.StateStartTime != nil {
			rStatus["stateStartTime"] = *r.Status.StateStartTime
		}
		if r.Status.Substate != nil {
			rStatus["substate"] = string(*r.Status.Substate)
		}
		u.Object["status"] = rStatus
	}
	var rStatusHistory []interface{}
	for _, rStatusHistoryVal := range r.StatusHistory {
		rStatusHistoryObject := make(map[string]interface{})
		if rStatusHistoryVal.Detail != nil {
			rStatusHistoryObject["detail"] = *rStatusHistoryVal.Detail
		}
		if rStatusHistoryVal.State != nil {
			rStatusHistoryObject["state"] = string(*rStatusHistoryVal.State)
		}
		if rStatusHistoryVal.StateStartTime != nil {
			rStatusHistoryObject["stateStartTime"] = *rStatusHistoryVal.StateStartTime
		}
		if rStatusHistoryVal.Substate != nil {
			rStatusHistoryObject["substate"] = string(*rStatusHistoryVal.Substate)
		}
		rStatusHistory = append(rStatusHistory, rStatusHistoryObject)
	}
	u.Object["statusHistory"] = rStatusHistory
	if r.VirtualClusterConfig != nil && r.VirtualClusterConfig != dclService.EmptyClusterVirtualClusterConfig {
		rVirtualClusterConfig := make(map[string]interface{})
		if r.VirtualClusterConfig.AuxiliaryServicesConfig != nil && r.VirtualClusterConfig.AuxiliaryServicesConfig != dclService.EmptyClusterVirtualClusterConfigAuxiliaryServicesConfig {
			rVirtualClusterConfigAuxiliaryServicesConfig := make(map[string]interface{})
			if r.VirtualClusterConfig.AuxiliaryServicesConfig.MetastoreConfig != nil && r.VirtualClusterConfig.AuxiliaryServicesConfig.MetastoreConfig != dclService.EmptyClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig {
				rVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig := make(map[string]interface{})
				if r.VirtualClusterConfig.AuxiliaryServicesConfig.MetastoreConfig.DataprocMetastoreService != nil {
					rVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig["dataprocMetastoreService"] = *r.VirtualClusterConfig.AuxiliaryServicesConfig.MetastoreConfig.DataprocMetastoreService
				}
				rVirtualClusterConfigAuxiliaryServicesConfig["metastoreConfig"] = rVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig
			}
			if r.VirtualClusterConfig.AuxiliaryServicesConfig.SparkHistoryServerConfig != nil && r.VirtualClusterConfig.AuxiliaryServicesConfig.SparkHistoryServerConfig != dclService.EmptyClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig {
				rVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig := make(map[string]interface{})
				if r.VirtualClusterConfig.AuxiliaryServicesConfig.SparkHistoryServerConfig.DataprocCluster != nil {
					rVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig["dataprocCluster"] = *r.VirtualClusterConfig.AuxiliaryServicesConfig.SparkHistoryServerConfig.DataprocCluster
				}
				rVirtualClusterConfigAuxiliaryServicesConfig["sparkHistoryServerConfig"] = rVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig
			}
			rVirtualClusterConfig["auxiliaryServicesConfig"] = rVirtualClusterConfigAuxiliaryServicesConfig
		}
		if r.VirtualClusterConfig.KubernetesClusterConfig != nil && r.VirtualClusterConfig.KubernetesClusterConfig != dclService.EmptyClusterVirtualClusterConfigKubernetesClusterConfig {
			rVirtualClusterConfigKubernetesClusterConfig := make(map[string]interface{})
			if r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig != nil && r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig != dclService.EmptyClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig {
				rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig := make(map[string]interface{})
				if r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.GkeClusterTarget != nil {
					rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig["gkeClusterTarget"] = *r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.GkeClusterTarget
				}
				var rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget []interface{}
				for _, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal := range r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.NodePoolTarget {
					rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetObject := make(map[string]interface{})
					if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePool != nil {
						rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetObject["nodePool"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePool
					}
					if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig != nil && rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig != dclService.EmptyClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig {
						rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfig := make(map[string]interface{})
						if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Autoscaling != nil && rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Autoscaling != dclService.EmptyClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling {
							rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigAutoscaling := make(map[string]interface{})
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Autoscaling.MaxNodeCount != nil {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigAutoscaling["maxNodeCount"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Autoscaling.MaxNodeCount
							}
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Autoscaling.MinNodeCount != nil {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigAutoscaling["minNodeCount"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Autoscaling.MinNodeCount
							}
							rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfig["autoscaling"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigAutoscaling
						}
						if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config != nil && rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config != dclService.EmptyClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig {
							rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig := make(map[string]interface{})
							var rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAccelerators []interface{}
							for _, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsVal := range rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.Accelerators {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsObject := make(map[string]interface{})
								if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsVal.AcceleratorCount != nil {
									rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsObject["acceleratorCount"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsVal.AcceleratorCount
								}
								if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsVal.AcceleratorType != nil {
									rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsObject["acceleratorType"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsVal.AcceleratorType
								}
								if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsVal.GpuPartitionSize != nil {
									rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsObject["gpuPartitionSize"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsVal.GpuPartitionSize
								}
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAccelerators = append(rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAccelerators, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAcceleratorsObject)
							}
							rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig["accelerators"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigAccelerators
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.BootDiskKmsKey != nil {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig["bootDiskKmsKey"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.BootDiskKmsKey
							}
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.EphemeralStorageConfig != nil && rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.EphemeralStorageConfig != dclService.EmptyClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigEphemeralStorageConfig := make(map[string]interface{})
								if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.EphemeralStorageConfig.LocalSsdCount != nil {
									rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigEphemeralStorageConfig["localSsdCount"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.EphemeralStorageConfig.LocalSsdCount
								}
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig["ephemeralStorageConfig"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfigEphemeralStorageConfig
							}
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.LocalSsdCount != nil {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig["localSsdCount"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.LocalSsdCount
							}
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.MachineType != nil {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig["machineType"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.MachineType
							}
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.MinCpuPlatform != nil {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig["minCpuPlatform"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.MinCpuPlatform
							}
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.Preemptible != nil {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig["preemptible"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.Preemptible
							}
							if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.Spot != nil {
								rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig["spot"] = *rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Config.Spot
							}
							rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfig["config"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigConfig
						}
						var rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigLocations []interface{}
						for _, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigLocationsVal := range rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.NodePoolConfig.Locations {
							rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigLocations = append(rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigLocations, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigLocationsVal)
						}
						rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfig["locations"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfigLocations
						rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetObject["nodePoolConfig"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValNodePoolConfig
					}
					var rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValRoles []interface{}
					for _, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValRolesVal := range rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetVal.Roles {
						rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValRoles = append(rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValRoles, string(rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValRolesVal))
					}
					rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetObject["roles"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetValRoles
					rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget = append(rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetObject)
				}
				rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig["nodePoolTarget"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget
				rVirtualClusterConfigKubernetesClusterConfig["gkeClusterConfig"] = rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig
			}
			if r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesNamespace != nil {
				rVirtualClusterConfigKubernetesClusterConfig["kubernetesNamespace"] = *r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesNamespace
			}
			if r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig != nil && r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig != dclService.EmptyClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig {
				rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig := make(map[string]interface{})
				if r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig.ComponentVersion != nil {
					rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigComponentVersion := make(map[string]interface{})
					for k, v := range r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig.ComponentVersion {
						rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigComponentVersion[k] = v
					}
					rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig["componentVersion"] = rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigComponentVersion
				}
				if r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig.Properties != nil {
					rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigProperties := make(map[string]interface{})
					for k, v := range r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig.Properties {
						rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigProperties[k] = v
					}
					rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig["properties"] = rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigProperties
				}
				rVirtualClusterConfigKubernetesClusterConfig["kubernetesSoftwareConfig"] = rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig
			}
			rVirtualClusterConfig["kubernetesClusterConfig"] = rVirtualClusterConfigKubernetesClusterConfig
		}
		if r.VirtualClusterConfig.StagingBucket != nil {
			rVirtualClusterConfig["stagingBucket"] = *r.VirtualClusterConfig.StagingBucket
		}
		u.Object["virtualClusterConfig"] = rVirtualClusterConfig
	}
	return u
}

func UnstructuredToCluster(u *unstructured.Resource) (*dclService.Cluster, error) {
	r := &dclService.Cluster{}
	if _, ok := u.Object["clusterUuid"]; ok {
		if s, ok := u.Object["clusterUuid"].(string); ok {
			r.ClusterUuid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ClusterUuid: expected string")
		}
	}
	if _, ok := u.Object["config"]; ok {
		if rConfig, ok := u.Object["config"].(map[string]interface{}); ok {
			r.Config = &dclService.ClusterConfig{}
			if _, ok := rConfig["autoscalingConfig"]; ok {
				if rConfigAutoscalingConfig, ok := rConfig["autoscalingConfig"].(map[string]interface{}); ok {
					r.Config.AutoscalingConfig = &dclService.ClusterConfigAutoscalingConfig{}
					if _, ok := rConfigAutoscalingConfig["policy"]; ok {
						if s, ok := rConfigAutoscalingConfig["policy"].(string); ok {
							r.Config.AutoscalingConfig.Policy = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.AutoscalingConfig.Policy: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.AutoscalingConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["dataprocMetricConfig"]; ok {
				if rConfigDataprocMetricConfig, ok := rConfig["dataprocMetricConfig"].(map[string]interface{}); ok {
					r.Config.DataprocMetricConfig = &dclService.ClusterConfigDataprocMetricConfig{}
					if _, ok := rConfigDataprocMetricConfig["metrics"]; ok {
						if s, ok := rConfigDataprocMetricConfig["metrics"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigDataprocMetricConfigMetrics dclService.ClusterConfigDataprocMetricConfigMetrics
									if _, ok := objval["metricOverrides"]; ok {
										if s, ok := objval["metricOverrides"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rConfigDataprocMetricConfigMetrics.MetricOverrides = append(rConfigDataprocMetricConfigMetrics.MetricOverrides, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rConfigDataprocMetricConfigMetrics.MetricOverrides: expected []interface{}")
										}
									}
									if _, ok := objval["metricSource"]; ok {
										if s, ok := objval["metricSource"].(string); ok {
											rConfigDataprocMetricConfigMetrics.MetricSource = dclService.ClusterConfigDataprocMetricConfigMetricsMetricSourceEnumRef(s)
										} else {
											return nil, fmt.Errorf("rConfigDataprocMetricConfigMetrics.MetricSource: expected string")
										}
									}
									r.Config.DataprocMetricConfig.Metrics = append(r.Config.DataprocMetricConfig.Metrics, rConfigDataprocMetricConfigMetrics)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.DataprocMetricConfig.Metrics: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.DataprocMetricConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["encryptionConfig"]; ok {
				if rConfigEncryptionConfig, ok := rConfig["encryptionConfig"].(map[string]interface{}); ok {
					r.Config.EncryptionConfig = &dclService.ClusterConfigEncryptionConfig{}
					if _, ok := rConfigEncryptionConfig["gcePdKmsKeyName"]; ok {
						if s, ok := rConfigEncryptionConfig["gcePdKmsKeyName"].(string); ok {
							r.Config.EncryptionConfig.GcePdKmsKeyName = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.EncryptionConfig.GcePdKmsKeyName: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.EncryptionConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["endpointConfig"]; ok {
				if rConfigEndpointConfig, ok := rConfig["endpointConfig"].(map[string]interface{}); ok {
					r.Config.EndpointConfig = &dclService.ClusterConfigEndpointConfig{}
					if _, ok := rConfigEndpointConfig["enableHttpPortAccess"]; ok {
						if b, ok := rConfigEndpointConfig["enableHttpPortAccess"].(bool); ok {
							r.Config.EndpointConfig.EnableHttpPortAccess = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Config.EndpointConfig.EnableHttpPortAccess: expected bool")
						}
					}
					if _, ok := rConfigEndpointConfig["httpPorts"]; ok {
						if rConfigEndpointConfigHttpPorts, ok := rConfigEndpointConfig["httpPorts"].(map[string]interface{}); ok {
							m := make(map[string]string)
							for k, v := range rConfigEndpointConfigHttpPorts {
								if s, ok := v.(string); ok {
									m[k] = s
								}
							}
							r.Config.EndpointConfig.HttpPorts = m
						} else {
							return nil, fmt.Errorf("r.Config.EndpointConfig.HttpPorts: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.EndpointConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["gceClusterConfig"]; ok {
				if rConfigGceClusterConfig, ok := rConfig["gceClusterConfig"].(map[string]interface{}); ok {
					r.Config.GceClusterConfig = &dclService.ClusterConfigGceClusterConfig{}
					if _, ok := rConfigGceClusterConfig["confidentialInstanceConfig"]; ok {
						if rConfigGceClusterConfigConfidentialInstanceConfig, ok := rConfigGceClusterConfig["confidentialInstanceConfig"].(map[string]interface{}); ok {
							r.Config.GceClusterConfig.ConfidentialInstanceConfig = &dclService.ClusterConfigGceClusterConfigConfidentialInstanceConfig{}
							if _, ok := rConfigGceClusterConfigConfidentialInstanceConfig["enableConfidentialCompute"]; ok {
								if b, ok := rConfigGceClusterConfigConfidentialInstanceConfig["enableConfidentialCompute"].(bool); ok {
									r.Config.GceClusterConfig.ConfidentialInstanceConfig.EnableConfidentialCompute = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Config.GceClusterConfig.ConfidentialInstanceConfig.EnableConfidentialCompute: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.ConfidentialInstanceConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigGceClusterConfig["internalIPOnly"]; ok {
						if b, ok := rConfigGceClusterConfig["internalIPOnly"].(bool); ok {
							r.Config.GceClusterConfig.InternalIPOnly = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.InternalIPOnly: expected bool")
						}
					}
					if _, ok := rConfigGceClusterConfig["metadata"]; ok {
						if rConfigGceClusterConfigMetadata, ok := rConfigGceClusterConfig["metadata"].(map[string]interface{}); ok {
							m := make(map[string]string)
							for k, v := range rConfigGceClusterConfigMetadata {
								if s, ok := v.(string); ok {
									m[k] = s
								}
							}
							r.Config.GceClusterConfig.Metadata = m
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.Metadata: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigGceClusterConfig["network"]; ok {
						if s, ok := rConfigGceClusterConfig["network"].(string); ok {
							r.Config.GceClusterConfig.Network = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.Network: expected string")
						}
					}
					if _, ok := rConfigGceClusterConfig["nodeGroupAffinity"]; ok {
						if rConfigGceClusterConfigNodeGroupAffinity, ok := rConfigGceClusterConfig["nodeGroupAffinity"].(map[string]interface{}); ok {
							r.Config.GceClusterConfig.NodeGroupAffinity = &dclService.ClusterConfigGceClusterConfigNodeGroupAffinity{}
							if _, ok := rConfigGceClusterConfigNodeGroupAffinity["nodeGroup"]; ok {
								if s, ok := rConfigGceClusterConfigNodeGroupAffinity["nodeGroup"].(string); ok {
									r.Config.GceClusterConfig.NodeGroupAffinity.NodeGroup = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.GceClusterConfig.NodeGroupAffinity.NodeGroup: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.NodeGroupAffinity: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigGceClusterConfig["privateIPv6GoogleAccess"]; ok {
						if s, ok := rConfigGceClusterConfig["privateIPv6GoogleAccess"].(string); ok {
							r.Config.GceClusterConfig.PrivateIPv6GoogleAccess = dclService.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.PrivateIPv6GoogleAccess: expected string")
						}
					}
					if _, ok := rConfigGceClusterConfig["reservationAffinity"]; ok {
						if rConfigGceClusterConfigReservationAffinity, ok := rConfigGceClusterConfig["reservationAffinity"].(map[string]interface{}); ok {
							r.Config.GceClusterConfig.ReservationAffinity = &dclService.ClusterConfigGceClusterConfigReservationAffinity{}
							if _, ok := rConfigGceClusterConfigReservationAffinity["consumeReservationType"]; ok {
								if s, ok := rConfigGceClusterConfigReservationAffinity["consumeReservationType"].(string); ok {
									r.Config.GceClusterConfig.ReservationAffinity.ConsumeReservationType = dclService.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef(s)
								} else {
									return nil, fmt.Errorf("r.Config.GceClusterConfig.ReservationAffinity.ConsumeReservationType: expected string")
								}
							}
							if _, ok := rConfigGceClusterConfigReservationAffinity["key"]; ok {
								if s, ok := rConfigGceClusterConfigReservationAffinity["key"].(string); ok {
									r.Config.GceClusterConfig.ReservationAffinity.Key = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.GceClusterConfig.ReservationAffinity.Key: expected string")
								}
							}
							if _, ok := rConfigGceClusterConfigReservationAffinity["values"]; ok {
								if s, ok := rConfigGceClusterConfigReservationAffinity["values"].([]interface{}); ok {
									for _, ss := range s {
										if strval, ok := ss.(string); ok {
											r.Config.GceClusterConfig.ReservationAffinity.Values = append(r.Config.GceClusterConfig.ReservationAffinity.Values, strval)
										}
									}
								} else {
									return nil, fmt.Errorf("r.Config.GceClusterConfig.ReservationAffinity.Values: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.ReservationAffinity: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigGceClusterConfig["serviceAccount"]; ok {
						if s, ok := rConfigGceClusterConfig["serviceAccount"].(string); ok {
							r.Config.GceClusterConfig.ServiceAccount = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.ServiceAccount: expected string")
						}
					}
					if _, ok := rConfigGceClusterConfig["serviceAccountScopes"]; ok {
						if s, ok := rConfigGceClusterConfig["serviceAccountScopes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.GceClusterConfig.ServiceAccountScopes = append(r.Config.GceClusterConfig.ServiceAccountScopes, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.ServiceAccountScopes: expected []interface{}")
						}
					}
					if _, ok := rConfigGceClusterConfig["shieldedInstanceConfig"]; ok {
						if rConfigGceClusterConfigShieldedInstanceConfig, ok := rConfigGceClusterConfig["shieldedInstanceConfig"].(map[string]interface{}); ok {
							r.Config.GceClusterConfig.ShieldedInstanceConfig = &dclService.ClusterConfigGceClusterConfigShieldedInstanceConfig{}
							if _, ok := rConfigGceClusterConfigShieldedInstanceConfig["enableIntegrityMonitoring"]; ok {
								if b, ok := rConfigGceClusterConfigShieldedInstanceConfig["enableIntegrityMonitoring"].(bool); ok {
									r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring: expected bool")
								}
							}
							if _, ok := rConfigGceClusterConfigShieldedInstanceConfig["enableSecureBoot"]; ok {
								if b, ok := rConfigGceClusterConfigShieldedInstanceConfig["enableSecureBoot"].(bool); ok {
									r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot: expected bool")
								}
							}
							if _, ok := rConfigGceClusterConfigShieldedInstanceConfig["enableVtpm"]; ok {
								if b, ok := rConfigGceClusterConfigShieldedInstanceConfig["enableVtpm"].(bool); ok {
									r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.ShieldedInstanceConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigGceClusterConfig["subnetwork"]; ok {
						if s, ok := rConfigGceClusterConfig["subnetwork"].(string); ok {
							r.Config.GceClusterConfig.Subnetwork = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.Subnetwork: expected string")
						}
					}
					if _, ok := rConfigGceClusterConfig["tags"]; ok {
						if s, ok := rConfigGceClusterConfig["tags"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.GceClusterConfig.Tags = append(r.Config.GceClusterConfig.Tags, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.Tags: expected []interface{}")
						}
					}
					if _, ok := rConfigGceClusterConfig["zone"]; ok {
						if s, ok := rConfigGceClusterConfig["zone"].(string); ok {
							r.Config.GceClusterConfig.Zone = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.GceClusterConfig.Zone: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.GceClusterConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["gkeClusterConfig"]; ok {
				if rConfigGkeClusterConfig, ok := rConfig["gkeClusterConfig"].(map[string]interface{}); ok {
					r.Config.GkeClusterConfig = &dclService.ClusterConfigGkeClusterConfig{}
					if _, ok := rConfigGkeClusterConfig["namespacedGkeDeploymentTarget"]; ok {
						if rConfigGkeClusterConfigNamespacedGkeDeploymentTarget, ok := rConfigGkeClusterConfig["namespacedGkeDeploymentTarget"].(map[string]interface{}); ok {
							r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget = &dclService.ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
							if _, ok := rConfigGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"]; ok {
								if s, ok := rConfigGkeClusterConfigNamespacedGkeDeploymentTarget["clusterNamespace"].(string); ok {
									r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.ClusterNamespace: expected string")
								}
							}
							if _, ok := rConfigGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"]; ok {
								if s, ok := rConfigGkeClusterConfigNamespacedGkeDeploymentTarget["targetGkeCluster"].(string); ok {
									r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget.TargetGkeCluster: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.GkeClusterConfig.NamespacedGkeDeploymentTarget: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.GkeClusterConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["initializationActions"]; ok {
				if s, ok := rConfig["initializationActions"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rConfigInitializationActions dclService.ClusterConfigInitializationActions
							if _, ok := objval["executableFile"]; ok {
								if s, ok := objval["executableFile"].(string); ok {
									rConfigInitializationActions.ExecutableFile = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rConfigInitializationActions.ExecutableFile: expected string")
								}
							}
							if _, ok := objval["executionTimeout"]; ok {
								if s, ok := objval["executionTimeout"].(string); ok {
									rConfigInitializationActions.ExecutionTimeout = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rConfigInitializationActions.ExecutionTimeout: expected string")
								}
							}
							r.Config.InitializationActions = append(r.Config.InitializationActions, rConfigInitializationActions)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.InitializationActions: expected []interface{}")
				}
			}
			if _, ok := rConfig["lifecycleConfig"]; ok {
				if rConfigLifecycleConfig, ok := rConfig["lifecycleConfig"].(map[string]interface{}); ok {
					r.Config.LifecycleConfig = &dclService.ClusterConfigLifecycleConfig{}
					if _, ok := rConfigLifecycleConfig["autoDeleteTime"]; ok {
						if s, ok := rConfigLifecycleConfig["autoDeleteTime"].(string); ok {
							r.Config.LifecycleConfig.AutoDeleteTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.LifecycleConfig.AutoDeleteTime: expected string")
						}
					}
					if _, ok := rConfigLifecycleConfig["autoDeleteTtl"]; ok {
						if s, ok := rConfigLifecycleConfig["autoDeleteTtl"].(string); ok {
							r.Config.LifecycleConfig.AutoDeleteTtl = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.LifecycleConfig.AutoDeleteTtl: expected string")
						}
					}
					if _, ok := rConfigLifecycleConfig["idleDeleteTtl"]; ok {
						if s, ok := rConfigLifecycleConfig["idleDeleteTtl"].(string); ok {
							r.Config.LifecycleConfig.IdleDeleteTtl = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.LifecycleConfig.IdleDeleteTtl: expected string")
						}
					}
					if _, ok := rConfigLifecycleConfig["idleStartTime"]; ok {
						if s, ok := rConfigLifecycleConfig["idleStartTime"].(string); ok {
							r.Config.LifecycleConfig.IdleStartTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.LifecycleConfig.IdleStartTime: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.LifecycleConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["masterConfig"]; ok {
				if rConfigMasterConfig, ok := rConfig["masterConfig"].(map[string]interface{}); ok {
					r.Config.MasterConfig = &dclService.ClusterConfigMasterConfig{}
					if _, ok := rConfigMasterConfig["accelerators"]; ok {
						if s, ok := rConfigMasterConfig["accelerators"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigMasterConfigAccelerators dclService.ClusterConfigMasterConfigAccelerators
									if _, ok := objval["acceleratorCount"]; ok {
										if i, ok := objval["acceleratorCount"].(int64); ok {
											rConfigMasterConfigAccelerators.AcceleratorCount = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rConfigMasterConfigAccelerators.AcceleratorCount: expected int64")
										}
									}
									if _, ok := objval["acceleratorType"]; ok {
										if s, ok := objval["acceleratorType"].(string); ok {
											rConfigMasterConfigAccelerators.AcceleratorType = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigMasterConfigAccelerators.AcceleratorType: expected string")
										}
									}
									r.Config.MasterConfig.Accelerators = append(r.Config.MasterConfig.Accelerators, rConfigMasterConfigAccelerators)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.Accelerators: expected []interface{}")
						}
					}
					if _, ok := rConfigMasterConfig["diskConfig"]; ok {
						if rConfigMasterConfigDiskConfig, ok := rConfigMasterConfig["diskConfig"].(map[string]interface{}); ok {
							r.Config.MasterConfig.DiskConfig = &dclService.ClusterConfigMasterConfigDiskConfig{}
							if _, ok := rConfigMasterConfigDiskConfig["bootDiskSizeGb"]; ok {
								if i, ok := rConfigMasterConfigDiskConfig["bootDiskSizeGb"].(int64); ok {
									r.Config.MasterConfig.DiskConfig.BootDiskSizeGb = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Config.MasterConfig.DiskConfig.BootDiskSizeGb: expected int64")
								}
							}
							if _, ok := rConfigMasterConfigDiskConfig["bootDiskType"]; ok {
								if s, ok := rConfigMasterConfigDiskConfig["bootDiskType"].(string); ok {
									r.Config.MasterConfig.DiskConfig.BootDiskType = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.MasterConfig.DiskConfig.BootDiskType: expected string")
								}
							}
							if _, ok := rConfigMasterConfigDiskConfig["localSsdInterface"]; ok {
								if s, ok := rConfigMasterConfigDiskConfig["localSsdInterface"].(string); ok {
									r.Config.MasterConfig.DiskConfig.LocalSsdInterface = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.MasterConfig.DiskConfig.LocalSsdInterface: expected string")
								}
							}
							if _, ok := rConfigMasterConfigDiskConfig["numLocalSsds"]; ok {
								if i, ok := rConfigMasterConfigDiskConfig["numLocalSsds"].(int64); ok {
									r.Config.MasterConfig.DiskConfig.NumLocalSsds = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Config.MasterConfig.DiskConfig.NumLocalSsds: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.DiskConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigMasterConfig["image"]; ok {
						if s, ok := rConfigMasterConfig["image"].(string); ok {
							r.Config.MasterConfig.Image = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.Image: expected string")
						}
					}
					if _, ok := rConfigMasterConfig["instanceNames"]; ok {
						if s, ok := rConfigMasterConfig["instanceNames"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.MasterConfig.InstanceNames = append(r.Config.MasterConfig.InstanceNames, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.InstanceNames: expected []interface{}")
						}
					}
					if _, ok := rConfigMasterConfig["instanceReferences"]; ok {
						if s, ok := rConfigMasterConfig["instanceReferences"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigMasterConfigInstanceReferences dclService.ClusterConfigMasterConfigInstanceReferences
									if _, ok := objval["instanceId"]; ok {
										if s, ok := objval["instanceId"].(string); ok {
											rConfigMasterConfigInstanceReferences.InstanceId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigMasterConfigInstanceReferences.InstanceId: expected string")
										}
									}
									if _, ok := objval["instanceName"]; ok {
										if s, ok := objval["instanceName"].(string); ok {
											rConfigMasterConfigInstanceReferences.InstanceName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigMasterConfigInstanceReferences.InstanceName: expected string")
										}
									}
									if _, ok := objval["publicEciesKey"]; ok {
										if s, ok := objval["publicEciesKey"].(string); ok {
											rConfigMasterConfigInstanceReferences.PublicEciesKey = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigMasterConfigInstanceReferences.PublicEciesKey: expected string")
										}
									}
									if _, ok := objval["publicKey"]; ok {
										if s, ok := objval["publicKey"].(string); ok {
											rConfigMasterConfigInstanceReferences.PublicKey = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigMasterConfigInstanceReferences.PublicKey: expected string")
										}
									}
									r.Config.MasterConfig.InstanceReferences = append(r.Config.MasterConfig.InstanceReferences, rConfigMasterConfigInstanceReferences)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.InstanceReferences: expected []interface{}")
						}
					}
					if _, ok := rConfigMasterConfig["isPreemptible"]; ok {
						if b, ok := rConfigMasterConfig["isPreemptible"].(bool); ok {
							r.Config.MasterConfig.IsPreemptible = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.IsPreemptible: expected bool")
						}
					}
					if _, ok := rConfigMasterConfig["machineType"]; ok {
						if s, ok := rConfigMasterConfig["machineType"].(string); ok {
							r.Config.MasterConfig.MachineType = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.MachineType: expected string")
						}
					}
					if _, ok := rConfigMasterConfig["managedGroupConfig"]; ok {
						if rConfigMasterConfigManagedGroupConfig, ok := rConfigMasterConfig["managedGroupConfig"].(map[string]interface{}); ok {
							r.Config.MasterConfig.ManagedGroupConfig = &dclService.ClusterConfigMasterConfigManagedGroupConfig{}
							if _, ok := rConfigMasterConfigManagedGroupConfig["instanceGroupManagerName"]; ok {
								if s, ok := rConfigMasterConfigManagedGroupConfig["instanceGroupManagerName"].(string); ok {
									r.Config.MasterConfig.ManagedGroupConfig.InstanceGroupManagerName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.MasterConfig.ManagedGroupConfig.InstanceGroupManagerName: expected string")
								}
							}
							if _, ok := rConfigMasterConfigManagedGroupConfig["instanceTemplateName"]; ok {
								if s, ok := rConfigMasterConfigManagedGroupConfig["instanceTemplateName"].(string); ok {
									r.Config.MasterConfig.ManagedGroupConfig.InstanceTemplateName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.MasterConfig.ManagedGroupConfig.InstanceTemplateName: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.ManagedGroupConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigMasterConfig["minCpuPlatform"]; ok {
						if s, ok := rConfigMasterConfig["minCpuPlatform"].(string); ok {
							r.Config.MasterConfig.MinCpuPlatform = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.MinCpuPlatform: expected string")
						}
					}
					if _, ok := rConfigMasterConfig["numInstances"]; ok {
						if i, ok := rConfigMasterConfig["numInstances"].(int64); ok {
							r.Config.MasterConfig.NumInstances = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.NumInstances: expected int64")
						}
					}
					if _, ok := rConfigMasterConfig["preemptibility"]; ok {
						if s, ok := rConfigMasterConfig["preemptibility"].(string); ok {
							r.Config.MasterConfig.Preemptibility = dclService.ClusterConfigMasterConfigPreemptibilityEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Config.MasterConfig.Preemptibility: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.MasterConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["metastoreConfig"]; ok {
				if rConfigMetastoreConfig, ok := rConfig["metastoreConfig"].(map[string]interface{}); ok {
					r.Config.MetastoreConfig = &dclService.ClusterConfigMetastoreConfig{}
					if _, ok := rConfigMetastoreConfig["dataprocMetastoreService"]; ok {
						if s, ok := rConfigMetastoreConfig["dataprocMetastoreService"].(string); ok {
							r.Config.MetastoreConfig.DataprocMetastoreService = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.MetastoreConfig.DataprocMetastoreService: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.MetastoreConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["secondaryWorkerConfig"]; ok {
				if rConfigSecondaryWorkerConfig, ok := rConfig["secondaryWorkerConfig"].(map[string]interface{}); ok {
					r.Config.SecondaryWorkerConfig = &dclService.ClusterConfigSecondaryWorkerConfig{}
					if _, ok := rConfigSecondaryWorkerConfig["accelerators"]; ok {
						if s, ok := rConfigSecondaryWorkerConfig["accelerators"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigSecondaryWorkerConfigAccelerators dclService.ClusterConfigSecondaryWorkerConfigAccelerators
									if _, ok := objval["acceleratorCount"]; ok {
										if i, ok := objval["acceleratorCount"].(int64); ok {
											rConfigSecondaryWorkerConfigAccelerators.AcceleratorCount = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rConfigSecondaryWorkerConfigAccelerators.AcceleratorCount: expected int64")
										}
									}
									if _, ok := objval["acceleratorType"]; ok {
										if s, ok := objval["acceleratorType"].(string); ok {
											rConfigSecondaryWorkerConfigAccelerators.AcceleratorType = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigSecondaryWorkerConfigAccelerators.AcceleratorType: expected string")
										}
									}
									r.Config.SecondaryWorkerConfig.Accelerators = append(r.Config.SecondaryWorkerConfig.Accelerators, rConfigSecondaryWorkerConfigAccelerators)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.Accelerators: expected []interface{}")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["diskConfig"]; ok {
						if rConfigSecondaryWorkerConfigDiskConfig, ok := rConfigSecondaryWorkerConfig["diskConfig"].(map[string]interface{}); ok {
							r.Config.SecondaryWorkerConfig.DiskConfig = &dclService.ClusterConfigSecondaryWorkerConfigDiskConfig{}
							if _, ok := rConfigSecondaryWorkerConfigDiskConfig["bootDiskSizeGb"]; ok {
								if i, ok := rConfigSecondaryWorkerConfigDiskConfig["bootDiskSizeGb"].(int64); ok {
									r.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.DiskConfig.BootDiskSizeGb: expected int64")
								}
							}
							if _, ok := rConfigSecondaryWorkerConfigDiskConfig["bootDiskType"]; ok {
								if s, ok := rConfigSecondaryWorkerConfigDiskConfig["bootDiskType"].(string); ok {
									r.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.DiskConfig.BootDiskType: expected string")
								}
							}
							if _, ok := rConfigSecondaryWorkerConfigDiskConfig["localSsdInterface"]; ok {
								if s, ok := rConfigSecondaryWorkerConfigDiskConfig["localSsdInterface"].(string); ok {
									r.Config.SecondaryWorkerConfig.DiskConfig.LocalSsdInterface = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.DiskConfig.LocalSsdInterface: expected string")
								}
							}
							if _, ok := rConfigSecondaryWorkerConfigDiskConfig["numLocalSsds"]; ok {
								if i, ok := rConfigSecondaryWorkerConfigDiskConfig["numLocalSsds"].(int64); ok {
									r.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.DiskConfig.NumLocalSsds: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.DiskConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["image"]; ok {
						if s, ok := rConfigSecondaryWorkerConfig["image"].(string); ok {
							r.Config.SecondaryWorkerConfig.Image = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.Image: expected string")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["instanceNames"]; ok {
						if s, ok := rConfigSecondaryWorkerConfig["instanceNames"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.SecondaryWorkerConfig.InstanceNames = append(r.Config.SecondaryWorkerConfig.InstanceNames, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.InstanceNames: expected []interface{}")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["instanceReferences"]; ok {
						if s, ok := rConfigSecondaryWorkerConfig["instanceReferences"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigSecondaryWorkerConfigInstanceReferences dclService.ClusterConfigSecondaryWorkerConfigInstanceReferences
									if _, ok := objval["instanceId"]; ok {
										if s, ok := objval["instanceId"].(string); ok {
											rConfigSecondaryWorkerConfigInstanceReferences.InstanceId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigSecondaryWorkerConfigInstanceReferences.InstanceId: expected string")
										}
									}
									if _, ok := objval["instanceName"]; ok {
										if s, ok := objval["instanceName"].(string); ok {
											rConfigSecondaryWorkerConfigInstanceReferences.InstanceName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigSecondaryWorkerConfigInstanceReferences.InstanceName: expected string")
										}
									}
									if _, ok := objval["publicEciesKey"]; ok {
										if s, ok := objval["publicEciesKey"].(string); ok {
											rConfigSecondaryWorkerConfigInstanceReferences.PublicEciesKey = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigSecondaryWorkerConfigInstanceReferences.PublicEciesKey: expected string")
										}
									}
									if _, ok := objval["publicKey"]; ok {
										if s, ok := objval["publicKey"].(string); ok {
											rConfigSecondaryWorkerConfigInstanceReferences.PublicKey = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigSecondaryWorkerConfigInstanceReferences.PublicKey: expected string")
										}
									}
									r.Config.SecondaryWorkerConfig.InstanceReferences = append(r.Config.SecondaryWorkerConfig.InstanceReferences, rConfigSecondaryWorkerConfigInstanceReferences)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.InstanceReferences: expected []interface{}")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["isPreemptible"]; ok {
						if b, ok := rConfigSecondaryWorkerConfig["isPreemptible"].(bool); ok {
							r.Config.SecondaryWorkerConfig.IsPreemptible = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.IsPreemptible: expected bool")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["machineType"]; ok {
						if s, ok := rConfigSecondaryWorkerConfig["machineType"].(string); ok {
							r.Config.SecondaryWorkerConfig.MachineType = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.MachineType: expected string")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["managedGroupConfig"]; ok {
						if rConfigSecondaryWorkerConfigManagedGroupConfig, ok := rConfigSecondaryWorkerConfig["managedGroupConfig"].(map[string]interface{}); ok {
							r.Config.SecondaryWorkerConfig.ManagedGroupConfig = &dclService.ClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
							if _, ok := rConfigSecondaryWorkerConfigManagedGroupConfig["instanceGroupManagerName"]; ok {
								if s, ok := rConfigSecondaryWorkerConfigManagedGroupConfig["instanceGroupManagerName"].(string); ok {
									r.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceGroupManagerName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceGroupManagerName: expected string")
								}
							}
							if _, ok := rConfigSecondaryWorkerConfigManagedGroupConfig["instanceTemplateName"]; ok {
								if s, ok := rConfigSecondaryWorkerConfigManagedGroupConfig["instanceTemplateName"].(string); ok {
									r.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceTemplateName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.ManagedGroupConfig.InstanceTemplateName: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.ManagedGroupConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["minCpuPlatform"]; ok {
						if s, ok := rConfigSecondaryWorkerConfig["minCpuPlatform"].(string); ok {
							r.Config.SecondaryWorkerConfig.MinCpuPlatform = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.MinCpuPlatform: expected string")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["numInstances"]; ok {
						if i, ok := rConfigSecondaryWorkerConfig["numInstances"].(int64); ok {
							r.Config.SecondaryWorkerConfig.NumInstances = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.NumInstances: expected int64")
						}
					}
					if _, ok := rConfigSecondaryWorkerConfig["preemptibility"]; ok {
						if s, ok := rConfigSecondaryWorkerConfig["preemptibility"].(string); ok {
							r.Config.SecondaryWorkerConfig.Preemptibility = dclService.ClusterConfigSecondaryWorkerConfigPreemptibilityEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig.Preemptibility: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.SecondaryWorkerConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["securityConfig"]; ok {
				if rConfigSecurityConfig, ok := rConfig["securityConfig"].(map[string]interface{}); ok {
					r.Config.SecurityConfig = &dclService.ClusterConfigSecurityConfig{}
					if _, ok := rConfigSecurityConfig["identityConfig"]; ok {
						if rConfigSecurityConfigIdentityConfig, ok := rConfigSecurityConfig["identityConfig"].(map[string]interface{}); ok {
							r.Config.SecurityConfig.IdentityConfig = &dclService.ClusterConfigSecurityConfigIdentityConfig{}
							if _, ok := rConfigSecurityConfigIdentityConfig["userServiceAccountMapping"]; ok {
								if rConfigSecurityConfigIdentityConfigUserServiceAccountMapping, ok := rConfigSecurityConfigIdentityConfig["userServiceAccountMapping"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rConfigSecurityConfigIdentityConfigUserServiceAccountMapping {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									r.Config.SecurityConfig.IdentityConfig.UserServiceAccountMapping = m
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.IdentityConfig.UserServiceAccountMapping: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SecurityConfig.IdentityConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigSecurityConfig["kerberosConfig"]; ok {
						if rConfigSecurityConfigKerberosConfig, ok := rConfigSecurityConfig["kerberosConfig"].(map[string]interface{}); ok {
							r.Config.SecurityConfig.KerberosConfig = &dclService.ClusterConfigSecurityConfigKerberosConfig{}
							if _, ok := rConfigSecurityConfigKerberosConfig["crossRealmTrustAdminServer"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["crossRealmTrustAdminServer"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustAdminServer: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["crossRealmTrustKdc"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["crossRealmTrustKdc"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustKdc = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustKdc: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["crossRealmTrustRealm"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["crossRealmTrustRealm"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustRealm = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustRealm: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["crossRealmTrustSharedPassword"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.CrossRealmTrustSharedPassword: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["enableKerberos"]; ok {
								if b, ok := rConfigSecurityConfigKerberosConfig["enableKerberos"].(bool); ok {
									r.Config.SecurityConfig.KerberosConfig.EnableKerberos = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.EnableKerberos: expected bool")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["kdcDbKey"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["kdcDbKey"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.KdcDbKey = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.KdcDbKey: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["keyPassword"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["keyPassword"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.KeyPassword = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.KeyPassword: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["keystore"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["keystore"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.Keystore = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.Keystore: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["keystorePassword"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["keystorePassword"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.KeystorePassword = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.KeystorePassword: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["kmsKey"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["kmsKey"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.KmsKey = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.KmsKey: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["realm"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["realm"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.Realm = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.Realm: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["rootPrincipalPassword"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["rootPrincipalPassword"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.RootPrincipalPassword = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.RootPrincipalPassword: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["tgtLifetimeHours"]; ok {
								if i, ok := rConfigSecurityConfigKerberosConfig["tgtLifetimeHours"].(int64); ok {
									r.Config.SecurityConfig.KerberosConfig.TgtLifetimeHours = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.TgtLifetimeHours: expected int64")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["truststore"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["truststore"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.Truststore = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.Truststore: expected string")
								}
							}
							if _, ok := rConfigSecurityConfigKerberosConfig["truststorePassword"]; ok {
								if s, ok := rConfigSecurityConfigKerberosConfig["truststorePassword"].(string); ok {
									r.Config.SecurityConfig.KerberosConfig.TruststorePassword = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig.TruststorePassword: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SecurityConfig.KerberosConfig: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.SecurityConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["softwareConfig"]; ok {
				if rConfigSoftwareConfig, ok := rConfig["softwareConfig"].(map[string]interface{}); ok {
					r.Config.SoftwareConfig = &dclService.ClusterConfigSoftwareConfig{}
					if _, ok := rConfigSoftwareConfig["imageVersion"]; ok {
						if s, ok := rConfigSoftwareConfig["imageVersion"].(string); ok {
							r.Config.SoftwareConfig.ImageVersion = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.SoftwareConfig.ImageVersion: expected string")
						}
					}
					if _, ok := rConfigSoftwareConfig["optionalComponents"]; ok {
						if s, ok := rConfigSoftwareConfig["optionalComponents"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.SoftwareConfig.OptionalComponents = append(r.Config.SoftwareConfig.OptionalComponents, dclService.ClusterConfigSoftwareConfigOptionalComponentsEnum(strval))
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SoftwareConfig.OptionalComponents: expected []interface{}")
						}
					}
					if _, ok := rConfigSoftwareConfig["properties"]; ok {
						if rConfigSoftwareConfigProperties, ok := rConfigSoftwareConfig["properties"].(map[string]interface{}); ok {
							m := make(map[string]string)
							for k, v := range rConfigSoftwareConfigProperties {
								if s, ok := v.(string); ok {
									m[k] = s
								}
							}
							r.Config.SoftwareConfig.Properties = m
						} else {
							return nil, fmt.Errorf("r.Config.SoftwareConfig.Properties: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.SoftwareConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["stagingBucket"]; ok {
				if s, ok := rConfig["stagingBucket"].(string); ok {
					r.Config.StagingBucket = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Config.StagingBucket: expected string")
				}
			}
			if _, ok := rConfig["tempBucket"]; ok {
				if s, ok := rConfig["tempBucket"].(string); ok {
					r.Config.TempBucket = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Config.TempBucket: expected string")
				}
			}
			if _, ok := rConfig["workerConfig"]; ok {
				if rConfigWorkerConfig, ok := rConfig["workerConfig"].(map[string]interface{}); ok {
					r.Config.WorkerConfig = &dclService.ClusterConfigWorkerConfig{}
					if _, ok := rConfigWorkerConfig["accelerators"]; ok {
						if s, ok := rConfigWorkerConfig["accelerators"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigWorkerConfigAccelerators dclService.ClusterConfigWorkerConfigAccelerators
									if _, ok := objval["acceleratorCount"]; ok {
										if i, ok := objval["acceleratorCount"].(int64); ok {
											rConfigWorkerConfigAccelerators.AcceleratorCount = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rConfigWorkerConfigAccelerators.AcceleratorCount: expected int64")
										}
									}
									if _, ok := objval["acceleratorType"]; ok {
										if s, ok := objval["acceleratorType"].(string); ok {
											rConfigWorkerConfigAccelerators.AcceleratorType = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigWorkerConfigAccelerators.AcceleratorType: expected string")
										}
									}
									r.Config.WorkerConfig.Accelerators = append(r.Config.WorkerConfig.Accelerators, rConfigWorkerConfigAccelerators)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.Accelerators: expected []interface{}")
						}
					}
					if _, ok := rConfigWorkerConfig["diskConfig"]; ok {
						if rConfigWorkerConfigDiskConfig, ok := rConfigWorkerConfig["diskConfig"].(map[string]interface{}); ok {
							r.Config.WorkerConfig.DiskConfig = &dclService.ClusterConfigWorkerConfigDiskConfig{}
							if _, ok := rConfigWorkerConfigDiskConfig["bootDiskSizeGb"]; ok {
								if i, ok := rConfigWorkerConfigDiskConfig["bootDiskSizeGb"].(int64); ok {
									r.Config.WorkerConfig.DiskConfig.BootDiskSizeGb = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Config.WorkerConfig.DiskConfig.BootDiskSizeGb: expected int64")
								}
							}
							if _, ok := rConfigWorkerConfigDiskConfig["bootDiskType"]; ok {
								if s, ok := rConfigWorkerConfigDiskConfig["bootDiskType"].(string); ok {
									r.Config.WorkerConfig.DiskConfig.BootDiskType = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.WorkerConfig.DiskConfig.BootDiskType: expected string")
								}
							}
							if _, ok := rConfigWorkerConfigDiskConfig["localSsdInterface"]; ok {
								if s, ok := rConfigWorkerConfigDiskConfig["localSsdInterface"].(string); ok {
									r.Config.WorkerConfig.DiskConfig.LocalSsdInterface = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.WorkerConfig.DiskConfig.LocalSsdInterface: expected string")
								}
							}
							if _, ok := rConfigWorkerConfigDiskConfig["numLocalSsds"]; ok {
								if i, ok := rConfigWorkerConfigDiskConfig["numLocalSsds"].(int64); ok {
									r.Config.WorkerConfig.DiskConfig.NumLocalSsds = dcl.Int64(i)
								} else {
									return nil, fmt.Errorf("r.Config.WorkerConfig.DiskConfig.NumLocalSsds: expected int64")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.DiskConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigWorkerConfig["image"]; ok {
						if s, ok := rConfigWorkerConfig["image"].(string); ok {
							r.Config.WorkerConfig.Image = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.Image: expected string")
						}
					}
					if _, ok := rConfigWorkerConfig["instanceNames"]; ok {
						if s, ok := rConfigWorkerConfig["instanceNames"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.WorkerConfig.InstanceNames = append(r.Config.WorkerConfig.InstanceNames, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.InstanceNames: expected []interface{}")
						}
					}
					if _, ok := rConfigWorkerConfig["instanceReferences"]; ok {
						if s, ok := rConfigWorkerConfig["instanceReferences"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rConfigWorkerConfigInstanceReferences dclService.ClusterConfigWorkerConfigInstanceReferences
									if _, ok := objval["instanceId"]; ok {
										if s, ok := objval["instanceId"].(string); ok {
											rConfigWorkerConfigInstanceReferences.InstanceId = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigWorkerConfigInstanceReferences.InstanceId: expected string")
										}
									}
									if _, ok := objval["instanceName"]; ok {
										if s, ok := objval["instanceName"].(string); ok {
											rConfigWorkerConfigInstanceReferences.InstanceName = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigWorkerConfigInstanceReferences.InstanceName: expected string")
										}
									}
									if _, ok := objval["publicEciesKey"]; ok {
										if s, ok := objval["publicEciesKey"].(string); ok {
											rConfigWorkerConfigInstanceReferences.PublicEciesKey = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigWorkerConfigInstanceReferences.PublicEciesKey: expected string")
										}
									}
									if _, ok := objval["publicKey"]; ok {
										if s, ok := objval["publicKey"].(string); ok {
											rConfigWorkerConfigInstanceReferences.PublicKey = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rConfigWorkerConfigInstanceReferences.PublicKey: expected string")
										}
									}
									r.Config.WorkerConfig.InstanceReferences = append(r.Config.WorkerConfig.InstanceReferences, rConfigWorkerConfigInstanceReferences)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.InstanceReferences: expected []interface{}")
						}
					}
					if _, ok := rConfigWorkerConfig["isPreemptible"]; ok {
						if b, ok := rConfigWorkerConfig["isPreemptible"].(bool); ok {
							r.Config.WorkerConfig.IsPreemptible = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.IsPreemptible: expected bool")
						}
					}
					if _, ok := rConfigWorkerConfig["machineType"]; ok {
						if s, ok := rConfigWorkerConfig["machineType"].(string); ok {
							r.Config.WorkerConfig.MachineType = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.MachineType: expected string")
						}
					}
					if _, ok := rConfigWorkerConfig["managedGroupConfig"]; ok {
						if rConfigWorkerConfigManagedGroupConfig, ok := rConfigWorkerConfig["managedGroupConfig"].(map[string]interface{}); ok {
							r.Config.WorkerConfig.ManagedGroupConfig = &dclService.ClusterConfigWorkerConfigManagedGroupConfig{}
							if _, ok := rConfigWorkerConfigManagedGroupConfig["instanceGroupManagerName"]; ok {
								if s, ok := rConfigWorkerConfigManagedGroupConfig["instanceGroupManagerName"].(string); ok {
									r.Config.WorkerConfig.ManagedGroupConfig.InstanceGroupManagerName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.WorkerConfig.ManagedGroupConfig.InstanceGroupManagerName: expected string")
								}
							}
							if _, ok := rConfigWorkerConfigManagedGroupConfig["instanceTemplateName"]; ok {
								if s, ok := rConfigWorkerConfigManagedGroupConfig["instanceTemplateName"].(string); ok {
									r.Config.WorkerConfig.ManagedGroupConfig.InstanceTemplateName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.Config.WorkerConfig.ManagedGroupConfig.InstanceTemplateName: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.ManagedGroupConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rConfigWorkerConfig["minCpuPlatform"]; ok {
						if s, ok := rConfigWorkerConfig["minCpuPlatform"].(string); ok {
							r.Config.WorkerConfig.MinCpuPlatform = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.MinCpuPlatform: expected string")
						}
					}
					if _, ok := rConfigWorkerConfig["numInstances"]; ok {
						if i, ok := rConfigWorkerConfig["numInstances"].(int64); ok {
							r.Config.WorkerConfig.NumInstances = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.NumInstances: expected int64")
						}
					}
					if _, ok := rConfigWorkerConfig["preemptibility"]; ok {
						if s, ok := rConfigWorkerConfig["preemptibility"].(string); ok {
							r.Config.WorkerConfig.Preemptibility = dclService.ClusterConfigWorkerConfigPreemptibilityEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Config.WorkerConfig.Preemptibility: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.WorkerConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Config: expected map[string]interface{}")
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
	if _, ok := u.Object["metrics"]; ok {
		if rMetrics, ok := u.Object["metrics"].(map[string]interface{}); ok {
			r.Metrics = &dclService.ClusterMetrics{}
			if _, ok := rMetrics["hdfsMetrics"]; ok {
				if rMetricsHdfsMetrics, ok := rMetrics["hdfsMetrics"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rMetricsHdfsMetrics {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.Metrics.HdfsMetrics = m
				} else {
					return nil, fmt.Errorf("r.Metrics.HdfsMetrics: expected map[string]interface{}")
				}
			}
			if _, ok := rMetrics["yarnMetrics"]; ok {
				if rMetricsYarnMetrics, ok := rMetrics["yarnMetrics"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rMetricsYarnMetrics {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.Metrics.YarnMetrics = m
				} else {
					return nil, fmt.Errorf("r.Metrics.YarnMetrics: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Metrics: expected map[string]interface{}")
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
	if _, ok := u.Object["status"]; ok {
		if rStatus, ok := u.Object["status"].(map[string]interface{}); ok {
			r.Status = &dclService.ClusterStatus{}
			if _, ok := rStatus["detail"]; ok {
				if s, ok := rStatus["detail"].(string); ok {
					r.Status.Detail = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Status.Detail: expected string")
				}
			}
			if _, ok := rStatus["state"]; ok {
				if s, ok := rStatus["state"].(string); ok {
					r.Status.State = dclService.ClusterStatusStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Status.State: expected string")
				}
			}
			if _, ok := rStatus["stateStartTime"]; ok {
				if s, ok := rStatus["stateStartTime"].(string); ok {
					r.Status.StateStartTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Status.StateStartTime: expected string")
				}
			}
			if _, ok := rStatus["substate"]; ok {
				if s, ok := rStatus["substate"].(string); ok {
					r.Status.Substate = dclService.ClusterStatusSubstateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Status.Substate: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Status: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["statusHistory"]; ok {
		if s, ok := u.Object["statusHistory"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rStatusHistory dclService.ClusterStatusHistory
					if _, ok := objval["detail"]; ok {
						if s, ok := objval["detail"].(string); ok {
							rStatusHistory.Detail = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rStatusHistory.Detail: expected string")
						}
					}
					if _, ok := objval["state"]; ok {
						if s, ok := objval["state"].(string); ok {
							rStatusHistory.State = dclService.ClusterStatusHistoryStateEnumRef(s)
						} else {
							return nil, fmt.Errorf("rStatusHistory.State: expected string")
						}
					}
					if _, ok := objval["stateStartTime"]; ok {
						if s, ok := objval["stateStartTime"].(string); ok {
							rStatusHistory.StateStartTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rStatusHistory.StateStartTime: expected string")
						}
					}
					if _, ok := objval["substate"]; ok {
						if s, ok := objval["substate"].(string); ok {
							rStatusHistory.Substate = dclService.ClusterStatusHistorySubstateEnumRef(s)
						} else {
							return nil, fmt.Errorf("rStatusHistory.Substate: expected string")
						}
					}
					r.StatusHistory = append(r.StatusHistory, rStatusHistory)
				}
			}
		} else {
			return nil, fmt.Errorf("r.StatusHistory: expected []interface{}")
		}
	}
	if _, ok := u.Object["virtualClusterConfig"]; ok {
		if rVirtualClusterConfig, ok := u.Object["virtualClusterConfig"].(map[string]interface{}); ok {
			r.VirtualClusterConfig = &dclService.ClusterVirtualClusterConfig{}
			if _, ok := rVirtualClusterConfig["auxiliaryServicesConfig"]; ok {
				if rVirtualClusterConfigAuxiliaryServicesConfig, ok := rVirtualClusterConfig["auxiliaryServicesConfig"].(map[string]interface{}); ok {
					r.VirtualClusterConfig.AuxiliaryServicesConfig = &dclService.ClusterVirtualClusterConfigAuxiliaryServicesConfig{}
					if _, ok := rVirtualClusterConfigAuxiliaryServicesConfig["metastoreConfig"]; ok {
						if rVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig, ok := rVirtualClusterConfigAuxiliaryServicesConfig["metastoreConfig"].(map[string]interface{}); ok {
							r.VirtualClusterConfig.AuxiliaryServicesConfig.MetastoreConfig = &dclService.ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig{}
							if _, ok := rVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig["dataprocMetastoreService"]; ok {
								if s, ok := rVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig["dataprocMetastoreService"].(string); ok {
									r.VirtualClusterConfig.AuxiliaryServicesConfig.MetastoreConfig.DataprocMetastoreService = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.VirtualClusterConfig.AuxiliaryServicesConfig.MetastoreConfig.DataprocMetastoreService: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.VirtualClusterConfig.AuxiliaryServicesConfig.MetastoreConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rVirtualClusterConfigAuxiliaryServicesConfig["sparkHistoryServerConfig"]; ok {
						if rVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig, ok := rVirtualClusterConfigAuxiliaryServicesConfig["sparkHistoryServerConfig"].(map[string]interface{}); ok {
							r.VirtualClusterConfig.AuxiliaryServicesConfig.SparkHistoryServerConfig = &dclService.ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig{}
							if _, ok := rVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig["dataprocCluster"]; ok {
								if s, ok := rVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig["dataprocCluster"].(string); ok {
									r.VirtualClusterConfig.AuxiliaryServicesConfig.SparkHistoryServerConfig.DataprocCluster = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.VirtualClusterConfig.AuxiliaryServicesConfig.SparkHistoryServerConfig.DataprocCluster: expected string")
								}
							}
						} else {
							return nil, fmt.Errorf("r.VirtualClusterConfig.AuxiliaryServicesConfig.SparkHistoryServerConfig: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.VirtualClusterConfig.AuxiliaryServicesConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rVirtualClusterConfig["kubernetesClusterConfig"]; ok {
				if rVirtualClusterConfigKubernetesClusterConfig, ok := rVirtualClusterConfig["kubernetesClusterConfig"].(map[string]interface{}); ok {
					r.VirtualClusterConfig.KubernetesClusterConfig = &dclService.ClusterVirtualClusterConfigKubernetesClusterConfig{}
					if _, ok := rVirtualClusterConfigKubernetesClusterConfig["gkeClusterConfig"]; ok {
						if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig, ok := rVirtualClusterConfigKubernetesClusterConfig["gkeClusterConfig"].(map[string]interface{}); ok {
							r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig = &dclService.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig{}
							if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig["gkeClusterTarget"]; ok {
								if s, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig["gkeClusterTarget"].(string); ok {
									r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.GkeClusterTarget = dcl.String(s)
								} else {
									return nil, fmt.Errorf("r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.GkeClusterTarget: expected string")
								}
							}
							if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig["nodePoolTarget"]; ok {
								if s, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig["nodePoolTarget"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget dclService.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget
											if _, ok := objval["nodePool"]; ok {
												if s, ok := objval["nodePool"].(string); ok {
													rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePool = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePool: expected string")
												}
											}
											if _, ok := objval["nodePoolConfig"]; ok {
												if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig, ok := objval["nodePoolConfig"].(map[string]interface{}); ok {
													rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig = &dclService.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig{}
													if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig["autoscaling"]; ok {
														if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig["autoscaling"].(map[string]interface{}); ok {
															rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Autoscaling = &dclService.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling{}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling["maxNodeCount"]; ok {
																if i, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling["maxNodeCount"].(int64); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Autoscaling.MaxNodeCount = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Autoscaling.MaxNodeCount: expected int64")
																}
															}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling["minNodeCount"]; ok {
																if i, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling["minNodeCount"].(int64); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Autoscaling.MinNodeCount = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Autoscaling.MinNodeCount: expected int64")
																}
															}
														} else {
															return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Autoscaling: expected map[string]interface{}")
														}
													}
													if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig["config"]; ok {
														if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig["config"].(map[string]interface{}); ok {
															rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config = &dclService.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig{}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["accelerators"]; ok {
																if s, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["accelerators"].([]interface{}); ok {
																	for _, o := range s {
																		if objval, ok := o.(map[string]interface{}); ok {
																			var rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators dclService.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators
																			if _, ok := objval["acceleratorCount"]; ok {
																				if i, ok := objval["acceleratorCount"].(int64); ok {
																					rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators.AcceleratorCount = dcl.Int64(i)
																				} else {
																					return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators.AcceleratorCount: expected int64")
																				}
																			}
																			if _, ok := objval["acceleratorType"]; ok {
																				if s, ok := objval["acceleratorType"].(string); ok {
																					rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators.AcceleratorType = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators.AcceleratorType: expected string")
																				}
																			}
																			if _, ok := objval["gpuPartitionSize"]; ok {
																				if s, ok := objval["gpuPartitionSize"].(string); ok {
																					rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators.GpuPartitionSize = dcl.String(s)
																				} else {
																					return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators.GpuPartitionSize: expected string")
																				}
																			}
																			rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.Accelerators = append(rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.Accelerators, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators)
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.Accelerators: expected []interface{}")
																}
															}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["bootDiskKmsKey"]; ok {
																if s, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["bootDiskKmsKey"].(string); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.BootDiskKmsKey = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.BootDiskKmsKey: expected string")
																}
															}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["ephemeralStorageConfig"]; ok {
																if rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["ephemeralStorageConfig"].(map[string]interface{}); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.EphemeralStorageConfig = &dclService.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig{}
																	if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig["localSsdCount"]; ok {
																		if i, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig["localSsdCount"].(int64); ok {
																			rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.EphemeralStorageConfig.LocalSsdCount = dcl.Int64(i)
																		} else {
																			return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.EphemeralStorageConfig.LocalSsdCount: expected int64")
																		}
																	}
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.EphemeralStorageConfig: expected map[string]interface{}")
																}
															}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["localSsdCount"]; ok {
																if i, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["localSsdCount"].(int64); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.LocalSsdCount = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.LocalSsdCount: expected int64")
																}
															}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["machineType"]; ok {
																if s, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["machineType"].(string); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.MachineType = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.MachineType: expected string")
																}
															}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["minCpuPlatform"]; ok {
																if s, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["minCpuPlatform"].(string); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.MinCpuPlatform = dcl.String(s)
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.MinCpuPlatform: expected string")
																}
															}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["preemptible"]; ok {
																if b, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["preemptible"].(bool); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.Preemptible = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.Preemptible: expected bool")
																}
															}
															if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["spot"]; ok {
																if b, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig["spot"].(bool); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.Spot = dcl.Bool(b)
																} else {
																	return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config.Spot: expected bool")
																}
															}
														} else {
															return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Config: expected map[string]interface{}")
														}
													}
													if _, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig["locations"]; ok {
														if s, ok := rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig["locations"].([]interface{}); ok {
															for _, ss := range s {
																if strval, ok := ss.(string); ok {
																	rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Locations = append(rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Locations, strval)
																}
															}
														} else {
															return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig.Locations: expected []interface{}")
														}
													}
												} else {
													return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.NodePoolConfig: expected map[string]interface{}")
												}
											}
											if _, ok := objval["roles"]; ok {
												if s, ok := objval["roles"].([]interface{}); ok {
													for _, ss := range s {
														if strval, ok := ss.(string); ok {
															rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.Roles = append(rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.Roles, dclService.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(strval))
														}
													}
												} else {
													return nil, fmt.Errorf("rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget.Roles: expected []interface{}")
												}
											}
											r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.NodePoolTarget = append(r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.NodePoolTarget, rVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget)
										}
									}
								} else {
									return nil, fmt.Errorf("r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig.NodePoolTarget: expected []interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.VirtualClusterConfig.KubernetesClusterConfig.GkeClusterConfig: expected map[string]interface{}")
						}
					}
					if _, ok := rVirtualClusterConfigKubernetesClusterConfig["kubernetesNamespace"]; ok {
						if s, ok := rVirtualClusterConfigKubernetesClusterConfig["kubernetesNamespace"].(string); ok {
							r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesNamespace = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesNamespace: expected string")
						}
					}
					if _, ok := rVirtualClusterConfigKubernetesClusterConfig["kubernetesSoftwareConfig"]; ok {
						if rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig, ok := rVirtualClusterConfigKubernetesClusterConfig["kubernetesSoftwareConfig"].(map[string]interface{}); ok {
							r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig = &dclService.ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig{}
							if _, ok := rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig["componentVersion"]; ok {
								if rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigComponentVersion, ok := rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig["componentVersion"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigComponentVersion {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig.ComponentVersion = m
								} else {
									return nil, fmt.Errorf("r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig.ComponentVersion: expected map[string]interface{}")
								}
							}
							if _, ok := rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig["properties"]; ok {
								if rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigProperties, ok := rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig["properties"].(map[string]interface{}); ok {
									m := make(map[string]string)
									for k, v := range rVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigProperties {
										if s, ok := v.(string); ok {
											m[k] = s
										}
									}
									r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig.Properties = m
								} else {
									return nil, fmt.Errorf("r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig.Properties: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.VirtualClusterConfig.KubernetesClusterConfig.KubernetesSoftwareConfig: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.VirtualClusterConfig.KubernetesClusterConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rVirtualClusterConfig["stagingBucket"]; ok {
				if s, ok := rVirtualClusterConfig["stagingBucket"].(string); ok {
					r.VirtualClusterConfig.StagingBucket = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.VirtualClusterConfig.StagingBucket: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.VirtualClusterConfig: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetCluster(ctx, r)
	if err != nil {
		return nil, err
	}
	return ClusterToUnstructured(r), nil
}

func ListCluster(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListCluster(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ClusterToUnstructured(r))
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

func ApplyCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCluster(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyCluster(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ClusterToUnstructured(r), nil
}

func ClusterHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToCluster(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyCluster(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return err
	}
	return c.DeleteCluster(ctx, r)
}

func ClusterID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToCluster(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Cluster) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dataproc",
		"Cluster",
		"beta",
	}
}

func SetPolicyCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCluster(u)
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

func SetPolicyWithEtagCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCluster(u)
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

func GetPolicyCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCluster(u)
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

func SetPolicyMemberCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToCluster(u)
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

func GetPolicyMemberCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToCluster(u)
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

func DeletePolicyMemberCluster(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToCluster(u)
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

func (r *Cluster) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberCluster(ctx, config, resource, member)
}

func (r *Cluster) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberCluster(ctx, config, resource, role, member)
}

func (r *Cluster) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberCluster(ctx, config, resource, member)
}

func (r *Cluster) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyCluster(ctx, config, resource, policy)
}

func (r *Cluster) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagCluster(ctx, config, resource, policy)
}

func (r *Cluster) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyCluster(ctx, config, resource)
}

func (r *Cluster) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetCluster(ctx, config, resource)
}

func (r *Cluster) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyCluster(ctx, config, resource, opts...)
}

func (r *Cluster) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ClusterHasDiff(ctx, config, resource, opts...)
}

func (r *Cluster) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteCluster(ctx, config, resource)
}

func (r *Cluster) ID(resource *unstructured.Resource) (string, error) {
	return ClusterID(resource)
}

func init() {
	unstructured.Register(&Cluster{})
}
