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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/alpha/dataproc_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
)

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e alphapb.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *alpha.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := alpha.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *alpha.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := alpha.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigMasterConfigPreemptibilityEnum converts a ClusterConfigMasterConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterConfigMasterConfigPreemptibilityEnum(e alphapb.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum) *alpha.ClusterConfigMasterConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := alpha.ClusterConfigMasterConfigPreemptibilityEnum(n[len("DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigWorkerConfigPreemptibilityEnum converts a ClusterConfigWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum(e alphapb.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum) *alpha.ClusterConfigWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := alpha.ClusterConfigWorkerConfigPreemptibilityEnum(n[len("DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigSecondaryWorkerConfigPreemptibilityEnum converts a ClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(e alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum) *alpha.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := alpha.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum(n[len("DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigSoftwareConfigOptionalComponentsEnum converts a ClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum(e alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum) *alpha.ClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := alpha.ClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfigDataprocMetricConfigMetricsMetricSourceEnum converts a ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum(e alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum) *alpha.ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum_name[int32(e)]; ok {
		e := alpha.ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum(n[len("DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusStateEnum converts a ClusterStatusStateEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterStatusStateEnum(e alphapb.DataprocAlphaClusterStatusStateEnum) *alpha.ClusterStatusStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterStatusStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStatusStateEnum(n[len("DataprocAlphaClusterStatusStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusSubstateEnum converts a ClusterStatusSubstateEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterStatusSubstateEnum(e alphapb.DataprocAlphaClusterStatusSubstateEnum) *alpha.ClusterStatusSubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterStatusSubstateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStatusSubstateEnum(n[len("DataprocAlphaClusterStatusSubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusHistoryStateEnum converts a ClusterStatusHistoryStateEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterStatusHistoryStateEnum(e alphapb.DataprocAlphaClusterStatusHistoryStateEnum) *alpha.ClusterStatusHistoryStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterStatusHistoryStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStatusHistoryStateEnum(n[len("DataprocAlphaClusterStatusHistoryStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterStatusHistorySubstateEnum converts a ClusterStatusHistorySubstateEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterStatusHistorySubstateEnum(e alphapb.DataprocAlphaClusterStatusHistorySubstateEnum) *alpha.ClusterStatusHistorySubstateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterStatusHistorySubstateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStatusHistorySubstateEnum(n[len("DataprocAlphaClusterStatusHistorySubstateEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum enum from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(e alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum_name[int32(e)]; ok {
		e := alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(n[len("DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum"):])
		return &e
	}
	return nil
}

// ProtoToClusterConfig converts a ClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfig(p *alphapb.DataprocAlphaClusterConfig) *alpha.ClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
		GceClusterConfig:      ProtoToDataprocAlphaClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocAlphaClusterConfigMasterConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocAlphaClusterConfigWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocAlphaClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocAlphaClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocAlphaClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocAlphaClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocAlphaClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocAlphaClusterConfigEndpointConfig(p.GetEndpointConfig()),
		GkeClusterConfig:      ProtoToDataprocAlphaClusterConfigGkeClusterConfig(p.GetGkeClusterConfig()),
		MetastoreConfig:       ProtoToDataprocAlphaClusterConfigMetastoreConfig(p.GetMetastoreConfig()),
		DataprocMetricConfig:  ProtoToDataprocAlphaClusterConfigDataprocMetricConfig(p.GetDataprocMetricConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocAlphaClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfig converts a ClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigGceClusterConfig(p *alphapb.DataprocAlphaClusterConfigGceClusterConfig) *alpha.ClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigGceClusterConfig{
		Zone:                       dcl.StringOrNil(p.GetZone()),
		Network:                    dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:                 dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:             dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess:    ProtoToDataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:             dcl.StringOrNil(p.GetServiceAccount()),
		ReservationAffinity:        ProtoToDataprocAlphaClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:          ProtoToDataprocAlphaClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
		ShieldedInstanceConfig:     ProtoToDataprocAlphaClusterConfigGceClusterConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
		ConfidentialInstanceConfig: ProtoToDataprocAlphaClusterConfigGceClusterConfigConfidentialInstanceConfig(p.GetConfidentialInstanceConfig()),
	}
	for _, r := range p.GetServiceAccountScopes() {
		obj.ServiceAccountScopes = append(obj.ServiceAccountScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfigReservationAffinity converts a ClusterConfigGceClusterConfigReservationAffinity object from its proto representation.
func ProtoToDataprocAlphaClusterConfigGceClusterConfigReservationAffinity(p *alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinity) *alpha.ClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfigNodeGroupAffinity converts a ClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocAlphaClusterConfigGceClusterConfigNodeGroupAffinity(p *alphapb.DataprocAlphaClusterConfigGceClusterConfigNodeGroupAffinity) *alpha.ClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfigShieldedInstanceConfig converts a ClusterConfigGceClusterConfigShieldedInstanceConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigGceClusterConfigShieldedInstanceConfig(p *alphapb.DataprocAlphaClusterConfigGceClusterConfigShieldedInstanceConfig) *alpha.ClusterConfigGceClusterConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigGceClusterConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.GetEnableSecureBoot()),
		EnableVtpm:                dcl.Bool(p.GetEnableVtpm()),
		EnableIntegrityMonitoring: dcl.Bool(p.GetEnableIntegrityMonitoring()),
	}
	return obj
}

// ProtoToClusterConfigGceClusterConfigConfidentialInstanceConfig converts a ClusterConfigGceClusterConfigConfidentialInstanceConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigGceClusterConfigConfidentialInstanceConfig(p *alphapb.DataprocAlphaClusterConfigGceClusterConfigConfidentialInstanceConfig) *alpha.ClusterConfigGceClusterConfigConfidentialInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigGceClusterConfigConfidentialInstanceConfig{
		EnableConfidentialCompute: dcl.Bool(p.GetEnableConfidentialCompute()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfig converts a ClusterConfigMasterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigMasterConfig(p *alphapb.DataprocAlphaClusterConfigMasterConfig) *alpha.ClusterConfigMasterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigMasterConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocAlphaClusterConfigMasterConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocAlphaClusterConfigMasterConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocAlphaClusterConfigMasterConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaClusterConfigMasterConfigAccelerators(r))
	}
	for _, r := range p.GetInstanceReferences() {
		obj.InstanceReferences = append(obj.InstanceReferences, *ProtoToDataprocAlphaClusterConfigMasterConfigInstanceReferences(r))
	}
	return obj
}

// ProtoToClusterConfigMasterConfigDiskConfig converts a ClusterConfigMasterConfigDiskConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigMasterConfigDiskConfig(p *alphapb.DataprocAlphaClusterConfigMasterConfigDiskConfig) *alpha.ClusterConfigMasterConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigMasterConfigDiskConfig{
		BootDiskType:      dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb:    dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:      dcl.Int64OrNil(p.GetNumLocalSsds()),
		LocalSsdInterface: dcl.StringOrNil(p.GetLocalSsdInterface()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfigManagedGroupConfig converts a ClusterConfigMasterConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigMasterConfigManagedGroupConfig(p *alphapb.DataprocAlphaClusterConfigMasterConfigManagedGroupConfig) *alpha.ClusterConfigMasterConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigMasterConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfigAccelerators converts a ClusterConfigMasterConfigAccelerators object from its proto representation.
func ProtoToDataprocAlphaClusterConfigMasterConfigAccelerators(p *alphapb.DataprocAlphaClusterConfigMasterConfigAccelerators) *alpha.ClusterConfigMasterConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigMasterConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigMasterConfigInstanceReferences converts a ClusterConfigMasterConfigInstanceReferences object from its proto representation.
func ProtoToDataprocAlphaClusterConfigMasterConfigInstanceReferences(p *alphapb.DataprocAlphaClusterConfigMasterConfigInstanceReferences) *alpha.ClusterConfigMasterConfigInstanceReferences {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigMasterConfigInstanceReferences{
		InstanceName:   dcl.StringOrNil(p.GetInstanceName()),
		InstanceId:     dcl.StringOrNil(p.GetInstanceId()),
		PublicKey:      dcl.StringOrNil(p.GetPublicKey()),
		PublicEciesKey: dcl.StringOrNil(p.GetPublicEciesKey()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfig converts a ClusterConfigWorkerConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigWorkerConfig(p *alphapb.DataprocAlphaClusterConfigWorkerConfig) *alpha.ClusterConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocAlphaClusterConfigWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocAlphaClusterConfigWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaClusterConfigWorkerConfigAccelerators(r))
	}
	for _, r := range p.GetInstanceReferences() {
		obj.InstanceReferences = append(obj.InstanceReferences, *ProtoToDataprocAlphaClusterConfigWorkerConfigInstanceReferences(r))
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigDiskConfig converts a ClusterConfigWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigWorkerConfigDiskConfig(p *alphapb.DataprocAlphaClusterConfigWorkerConfigDiskConfig) *alpha.ClusterConfigWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigWorkerConfigDiskConfig{
		BootDiskType:      dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb:    dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:      dcl.Int64OrNil(p.GetNumLocalSsds()),
		LocalSsdInterface: dcl.StringOrNil(p.GetLocalSsdInterface()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigManagedGroupConfig converts a ClusterConfigWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigWorkerConfigManagedGroupConfig(p *alphapb.DataprocAlphaClusterConfigWorkerConfigManagedGroupConfig) *alpha.ClusterConfigWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigAccelerators converts a ClusterConfigWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocAlphaClusterConfigWorkerConfigAccelerators(p *alphapb.DataprocAlphaClusterConfigWorkerConfigAccelerators) *alpha.ClusterConfigWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigWorkerConfigInstanceReferences converts a ClusterConfigWorkerConfigInstanceReferences object from its proto representation.
func ProtoToDataprocAlphaClusterConfigWorkerConfigInstanceReferences(p *alphapb.DataprocAlphaClusterConfigWorkerConfigInstanceReferences) *alpha.ClusterConfigWorkerConfigInstanceReferences {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigWorkerConfigInstanceReferences{
		InstanceName:   dcl.StringOrNil(p.GetInstanceName()),
		InstanceId:     dcl.StringOrNil(p.GetInstanceId()),
		PublicKey:      dcl.StringOrNil(p.GetPublicKey()),
		PublicEciesKey: dcl.StringOrNil(p.GetPublicEciesKey()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfig converts a ClusterConfigSecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfig(p *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfig) *alpha.ClusterConfigSecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSecondaryWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigAccelerators(r))
	}
	for _, r := range p.GetInstanceReferences() {
		obj.InstanceReferences = append(obj.InstanceReferences, *ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferences(r))
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigDiskConfig converts a ClusterConfigSecondaryWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigDiskConfig(p *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigDiskConfig) *alpha.ClusterConfigSecondaryWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSecondaryWorkerConfigDiskConfig{
		BootDiskType:      dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb:    dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:      dcl.Int64OrNil(p.GetNumLocalSsds()),
		LocalSsdInterface: dcl.StringOrNil(p.GetLocalSsdInterface()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigManagedGroupConfig converts a ClusterConfigSecondaryWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigManagedGroupConfig(p *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigManagedGroupConfig) *alpha.ClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSecondaryWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigAccelerators converts a ClusterConfigSecondaryWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigAccelerators(p *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigAccelerators) *alpha.ClusterConfigSecondaryWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSecondaryWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToClusterConfigSecondaryWorkerConfigInstanceReferences converts a ClusterConfigSecondaryWorkerConfigInstanceReferences object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferences(p *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferences) *alpha.ClusterConfigSecondaryWorkerConfigInstanceReferences {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSecondaryWorkerConfigInstanceReferences{
		InstanceName:   dcl.StringOrNil(p.GetInstanceName()),
		InstanceId:     dcl.StringOrNil(p.GetInstanceId()),
		PublicKey:      dcl.StringOrNil(p.GetPublicKey()),
		PublicEciesKey: dcl.StringOrNil(p.GetPublicEciesKey()),
	}
	return obj
}

// ProtoToClusterConfigSoftwareConfig converts a ClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSoftwareConfig(p *alphapb.DataprocAlphaClusterConfigSoftwareConfig) *alpha.ClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToClusterConfigInitializationActions converts a ClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocAlphaClusterConfigInitializationActions(p *alphapb.DataprocAlphaClusterConfigInitializationActions) *alpha.ClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToClusterConfigEncryptionConfig converts a ClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigEncryptionConfig(p *alphapb.DataprocAlphaClusterConfigEncryptionConfig) *alpha.ClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToClusterConfigAutoscalingConfig converts a ClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigAutoscalingConfig(p *alphapb.DataprocAlphaClusterConfigAutoscalingConfig) *alpha.ClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToClusterConfigSecurityConfig converts a ClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecurityConfig(p *alphapb.DataprocAlphaClusterConfigSecurityConfig) *alpha.ClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocAlphaClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
		IdentityConfig: ProtoToDataprocAlphaClusterConfigSecurityConfigIdentityConfig(p.GetIdentityConfig()),
	}
	return obj
}

// ProtoToClusterConfigSecurityConfigKerberosConfig converts a ClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecurityConfigKerberosConfig(p *alphapb.DataprocAlphaClusterConfigSecurityConfigKerberosConfig) *alpha.ClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSecurityConfigKerberosConfig{
		EnableKerberos:                dcl.Bool(p.GetEnableKerberos()),
		RootPrincipalPassword:         dcl.StringOrNil(p.GetRootPrincipalPassword()),
		KmsKey:                        dcl.StringOrNil(p.GetKmsKey()),
		Keystore:                      dcl.StringOrNil(p.GetKeystore()),
		Truststore:                    dcl.StringOrNil(p.GetTruststore()),
		KeystorePassword:              dcl.StringOrNil(p.GetKeystorePassword()),
		KeyPassword:                   dcl.StringOrNil(p.GetKeyPassword()),
		TruststorePassword:            dcl.StringOrNil(p.GetTruststorePassword()),
		CrossRealmTrustRealm:          dcl.StringOrNil(p.GetCrossRealmTrustRealm()),
		CrossRealmTrustKdc:            dcl.StringOrNil(p.GetCrossRealmTrustKdc()),
		CrossRealmTrustAdminServer:    dcl.StringOrNil(p.GetCrossRealmTrustAdminServer()),
		CrossRealmTrustSharedPassword: dcl.StringOrNil(p.GetCrossRealmTrustSharedPassword()),
		KdcDbKey:                      dcl.StringOrNil(p.GetKdcDbKey()),
		TgtLifetimeHours:              dcl.Int64OrNil(p.GetTgtLifetimeHours()),
		Realm:                         dcl.StringOrNil(p.GetRealm()),
	}
	return obj
}

// ProtoToClusterConfigSecurityConfigIdentityConfig converts a ClusterConfigSecurityConfigIdentityConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigSecurityConfigIdentityConfig(p *alphapb.DataprocAlphaClusterConfigSecurityConfigIdentityConfig) *alpha.ClusterConfigSecurityConfigIdentityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigSecurityConfigIdentityConfig{}
	return obj
}

// ProtoToClusterConfigLifecycleConfig converts a ClusterConfigLifecycleConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigLifecycleConfig(p *alphapb.DataprocAlphaClusterConfigLifecycleConfig) *alpha.ClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToClusterConfigEndpointConfig converts a ClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigEndpointConfig(p *alphapb.DataprocAlphaClusterConfigEndpointConfig) *alpha.ClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.GetEnableHttpPortAccess()),
	}
	return obj
}

// ProtoToClusterConfigGkeClusterConfig converts a ClusterConfigGkeClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigGkeClusterConfig(p *alphapb.DataprocAlphaClusterConfigGkeClusterConfig) *alpha.ClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: ProtoToDataprocAlphaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p.GetNamespacedGkeDeploymentTarget()),
	}
	return obj
}

// ProtoToClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget converts a ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object from its proto representation.
func ProtoToDataprocAlphaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p *alphapb.DataprocAlphaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *alpha.ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.StringOrNil(p.GetTargetGkeCluster()),
		ClusterNamespace: dcl.StringOrNil(p.GetClusterNamespace()),
	}
	return obj
}

// ProtoToClusterConfigMetastoreConfig converts a ClusterConfigMetastoreConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigMetastoreConfig(p *alphapb.DataprocAlphaClusterConfigMetastoreConfig) *alpha.ClusterConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.StringOrNil(p.GetDataprocMetastoreService()),
	}
	return obj
}

// ProtoToClusterConfigDataprocMetricConfig converts a ClusterConfigDataprocMetricConfig object from its proto representation.
func ProtoToDataprocAlphaClusterConfigDataprocMetricConfig(p *alphapb.DataprocAlphaClusterConfigDataprocMetricConfig) *alpha.ClusterConfigDataprocMetricConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigDataprocMetricConfig{}
	for _, r := range p.GetMetrics() {
		obj.Metrics = append(obj.Metrics, *ProtoToDataprocAlphaClusterConfigDataprocMetricConfigMetrics(r))
	}
	return obj
}

// ProtoToClusterConfigDataprocMetricConfigMetrics converts a ClusterConfigDataprocMetricConfigMetrics object from its proto representation.
func ProtoToDataprocAlphaClusterConfigDataprocMetricConfigMetrics(p *alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetrics) *alpha.ClusterConfigDataprocMetricConfigMetrics {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterConfigDataprocMetricConfigMetrics{
		MetricSource: ProtoToDataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum(p.GetMetricSource()),
	}
	for _, r := range p.GetMetricOverrides() {
		obj.MetricOverrides = append(obj.MetricOverrides, r)
	}
	return obj
}

// ProtoToClusterStatus converts a ClusterStatus object from its proto representation.
func ProtoToDataprocAlphaClusterStatus(p *alphapb.DataprocAlphaClusterStatus) *alpha.ClusterStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterStatus{
		State:          ProtoToDataprocAlphaClusterStatusStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.GetDetail()),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocAlphaClusterStatusSubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterStatusHistory converts a ClusterStatusHistory object from its proto representation.
func ProtoToDataprocAlphaClusterStatusHistory(p *alphapb.DataprocAlphaClusterStatusHistory) *alpha.ClusterStatusHistory {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterStatusHistory{
		State:          ProtoToDataprocAlphaClusterStatusHistoryStateEnum(p.GetState()),
		Detail:         dcl.StringOrNil(p.GetDetail()),
		StateStartTime: dcl.StringOrNil(p.GetStateStartTime()),
		Substate:       ProtoToDataprocAlphaClusterStatusHistorySubstateEnum(p.GetSubstate()),
	}
	return obj
}

// ProtoToClusterMetrics converts a ClusterMetrics object from its proto representation.
func ProtoToDataprocAlphaClusterMetrics(p *alphapb.DataprocAlphaClusterMetrics) *alpha.ClusterMetrics {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterMetrics{}
	return obj
}

// ProtoToClusterVirtualClusterConfig converts a ClusterVirtualClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfig) *alpha.ClusterVirtualClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfig{
		StagingBucket:           dcl.StringOrNil(p.GetStagingBucket()),
		KubernetesClusterConfig: ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfig(p.GetKubernetesClusterConfig()),
		AuxiliaryServicesConfig: ProtoToDataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfig(p.GetAuxiliaryServicesConfig()),
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfig converts a ClusterVirtualClusterConfigKubernetesClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfig) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfig{
		KubernetesNamespace:      dcl.StringOrNil(p.GetKubernetesNamespace()),
		GkeClusterConfig:         ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig(p.GetGkeClusterConfig()),
		KubernetesSoftwareConfig: ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig(p.GetKubernetesSoftwareConfig()),
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig{
		GkeClusterTarget: dcl.StringOrNil(p.GetGkeClusterTarget()),
	}
	for _, r := range p.GetNodePoolTarget() {
		obj.NodePoolTarget = append(obj.NodePoolTarget, *ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget(r))
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget{
		NodePool:       dcl.StringOrNil(p.GetNodePool()),
		NodePoolConfig: ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig(p.GetNodePoolConfig()),
	}
	for _, r := range p.GetRoles() {
		obj.Roles = append(obj.Roles, *ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(r))
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig{
		Config:      ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig(p.GetConfig()),
		Autoscaling: ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling(p.GetAutoscaling()),
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig{
		MachineType:            dcl.StringOrNil(p.GetMachineType()),
		LocalSsdCount:          dcl.Int64OrNil(p.GetLocalSsdCount()),
		Preemptible:            dcl.Bool(p.GetPreemptible()),
		MinCpuPlatform:         dcl.StringOrNil(p.GetMinCpuPlatform()),
		BootDiskKmsKey:         dcl.StringOrNil(p.GetBootDiskKmsKey()),
		EphemeralStorageConfig: ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig(p.GetEphemeralStorageConfig()),
		Spot:                   dcl.Bool(p.GetSpot()),
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators(r))
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		GpuPartitionSize: dcl.StringOrNil(p.GetGpuPartitionSize()),
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig{
		LocalSsdCount: dcl.Int64OrNil(p.GetLocalSsdCount()),
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.GetMinNodeCount()),
		MaxNodeCount: dcl.Int64OrNil(p.GetMaxNodeCount()),
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig converts a ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig) *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig{}
	return obj
}

// ProtoToClusterVirtualClusterConfigAuxiliaryServicesConfig converts a ClusterVirtualClusterConfigAuxiliaryServicesConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfig) *alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfig{
		MetastoreConfig:          ProtoToDataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig(p.GetMetastoreConfig()),
		SparkHistoryServerConfig: ProtoToDataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig(p.GetSparkHistoryServerConfig()),
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig converts a ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig) *alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig{
		DataprocMetastoreService: dcl.StringOrNil(p.GetDataprocMetastoreService()),
	}
	return obj
}

// ProtoToClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig converts a ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig object from its proto representation.
func ProtoToDataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig(p *alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig) *alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig{
		DataprocCluster: dcl.StringOrNil(p.GetDataprocCluster()),
	}
	return obj
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *alphapb.DataprocAlphaCluster) *alpha.Cluster {
	obj := &alpha.Cluster{
		Project:              dcl.StringOrNil(p.GetProject()),
		Name:                 dcl.StringOrNil(p.GetName()),
		Config:               ProtoToDataprocAlphaClusterConfig(p.GetConfig()),
		Status:               ProtoToDataprocAlphaClusterStatus(p.GetStatus()),
		ClusterUuid:          dcl.StringOrNil(p.GetClusterUuid()),
		Metrics:              ProtoToDataprocAlphaClusterMetrics(p.GetMetrics()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		VirtualClusterConfig: ProtoToDataprocAlphaClusterVirtualClusterConfig(p.GetVirtualClusterConfig()),
	}
	for _, r := range p.GetStatusHistory() {
		obj.StatusHistory = append(obj.StatusHistory, *ProtoToDataprocAlphaClusterStatusHistory(r))
	}
	return obj
}

// ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *alpha.ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) alphapb.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["ClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return alphapb.DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *alpha.ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["ClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// ClusterConfigMasterConfigPreemptibilityEnumToProto converts a ClusterConfigMasterConfigPreemptibilityEnum enum to its proto representation.
func DataprocAlphaClusterConfigMasterConfigPreemptibilityEnumToProto(e *alpha.ClusterConfigMasterConfigPreemptibilityEnum) alphapb.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum_value["ClusterConfigMasterConfigPreemptibilityEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum(v)
	}
	return alphapb.DataprocAlphaClusterConfigMasterConfigPreemptibilityEnum(0)
}

// ClusterConfigWorkerConfigPreemptibilityEnumToProto converts a ClusterConfigWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnumToProto(e *alpha.ClusterConfigWorkerConfigPreemptibilityEnum) alphapb.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum_value["ClusterConfigWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum(v)
	}
	return alphapb.DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnum(0)
}

// ClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto converts a ClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(e *alpha.ClusterConfigSecondaryWorkerConfigPreemptibilityEnum) alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum_value["ClusterConfigSecondaryWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(v)
	}
	return alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
}

// ClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a ClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *alpha.ClusterConfigSoftwareConfigOptionalComponentsEnum) alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum_value["ClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum(0)
}

// ClusterConfigDataprocMetricConfigMetricsMetricSourceEnumToProto converts a ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum enum to its proto representation.
func DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnumToProto(e *alpha.ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum) alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum_value["ClusterConfigDataprocMetricConfigMetricsMetricSourceEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum(v)
	}
	return alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnum(0)
}

// ClusterStatusStateEnumToProto converts a ClusterStatusStateEnum enum to its proto representation.
func DataprocAlphaClusterStatusStateEnumToProto(e *alpha.ClusterStatusStateEnum) alphapb.DataprocAlphaClusterStatusStateEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterStatusStateEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterStatusStateEnum_value["ClusterStatusStateEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterStatusStateEnum(v)
	}
	return alphapb.DataprocAlphaClusterStatusStateEnum(0)
}

// ClusterStatusSubstateEnumToProto converts a ClusterStatusSubstateEnum enum to its proto representation.
func DataprocAlphaClusterStatusSubstateEnumToProto(e *alpha.ClusterStatusSubstateEnum) alphapb.DataprocAlphaClusterStatusSubstateEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterStatusSubstateEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterStatusSubstateEnum_value["ClusterStatusSubstateEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterStatusSubstateEnum(v)
	}
	return alphapb.DataprocAlphaClusterStatusSubstateEnum(0)
}

// ClusterStatusHistoryStateEnumToProto converts a ClusterStatusHistoryStateEnum enum to its proto representation.
func DataprocAlphaClusterStatusHistoryStateEnumToProto(e *alpha.ClusterStatusHistoryStateEnum) alphapb.DataprocAlphaClusterStatusHistoryStateEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterStatusHistoryStateEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterStatusHistoryStateEnum_value["ClusterStatusHistoryStateEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterStatusHistoryStateEnum(v)
	}
	return alphapb.DataprocAlphaClusterStatusHistoryStateEnum(0)
}

// ClusterStatusHistorySubstateEnumToProto converts a ClusterStatusHistorySubstateEnum enum to its proto representation.
func DataprocAlphaClusterStatusHistorySubstateEnumToProto(e *alpha.ClusterStatusHistorySubstateEnum) alphapb.DataprocAlphaClusterStatusHistorySubstateEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterStatusHistorySubstateEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterStatusHistorySubstateEnum_value["ClusterStatusHistorySubstateEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterStatusHistorySubstateEnum(v)
	}
	return alphapb.DataprocAlphaClusterStatusHistorySubstateEnum(0)
}

// ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnumToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum enum to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnumToProto(e *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum) alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum {
	if e == nil {
		return alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum_value["ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(v)
	}
	return alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(0)
}

// ClusterConfigToProto converts a ClusterConfig object to its proto representation.
func DataprocAlphaClusterConfigToProto(o *alpha.ClusterConfig) *alphapb.DataprocAlphaClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocAlphaClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocAlphaClusterConfigMasterConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocAlphaClusterConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocAlphaClusterConfigSecondaryWorkerConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocAlphaClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocAlphaClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocAlphaClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocAlphaClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocAlphaClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocAlphaClusterConfigEndpointConfigToProto(o.EndpointConfig))
	p.SetGkeClusterConfig(DataprocAlphaClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig))
	p.SetMetastoreConfig(DataprocAlphaClusterConfigMetastoreConfigToProto(o.MetastoreConfig))
	p.SetDataprocMetricConfig(DataprocAlphaClusterConfigDataprocMetricConfigToProto(o.DataprocMetricConfig))
	sInitializationActions := make([]*alphapb.DataprocAlphaClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocAlphaClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// ClusterConfigGceClusterConfigToProto converts a ClusterConfigGceClusterConfig object to its proto representation.
func DataprocAlphaClusterConfigGceClusterConfigToProto(o *alpha.ClusterConfigGceClusterConfig) *alphapb.DataprocAlphaClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocAlphaClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocAlphaClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocAlphaClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
	p.SetShieldedInstanceConfig(DataprocAlphaClusterConfigGceClusterConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig))
	p.SetConfidentialInstanceConfig(DataprocAlphaClusterConfigGceClusterConfigConfidentialInstanceConfigToProto(o.ConfidentialInstanceConfig))
	sServiceAccountScopes := make([]string, len(o.ServiceAccountScopes))
	for i, r := range o.ServiceAccountScopes {
		sServiceAccountScopes[i] = r
	}
	p.SetServiceAccountScopes(sServiceAccountScopes)
	sTags := make([]string, len(o.Tags))
	for i, r := range o.Tags {
		sTags[i] = r
	}
	p.SetTags(sTags)
	mMetadata := make(map[string]string, len(o.Metadata))
	for k, r := range o.Metadata {
		mMetadata[k] = r
	}
	p.SetMetadata(mMetadata)
	return p
}

// ClusterConfigGceClusterConfigReservationAffinityToProto converts a ClusterConfigGceClusterConfigReservationAffinity object to its proto representation.
func DataprocAlphaClusterConfigGceClusterConfigReservationAffinityToProto(o *alpha.ClusterConfigGceClusterConfigReservationAffinity) *alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocAlphaClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// ClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a ClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocAlphaClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *alpha.ClusterConfigGceClusterConfigNodeGroupAffinity) *alphapb.DataprocAlphaClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// ClusterConfigGceClusterConfigShieldedInstanceConfigToProto converts a ClusterConfigGceClusterConfigShieldedInstanceConfig object to its proto representation.
func DataprocAlphaClusterConfigGceClusterConfigShieldedInstanceConfigToProto(o *alpha.ClusterConfigGceClusterConfigShieldedInstanceConfig) *alphapb.DataprocAlphaClusterConfigGceClusterConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigGceClusterConfigShieldedInstanceConfig{}
	p.SetEnableSecureBoot(dcl.ValueOrEmptyBool(o.EnableSecureBoot))
	p.SetEnableVtpm(dcl.ValueOrEmptyBool(o.EnableVtpm))
	p.SetEnableIntegrityMonitoring(dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring))
	return p
}

// ClusterConfigGceClusterConfigConfidentialInstanceConfigToProto converts a ClusterConfigGceClusterConfigConfidentialInstanceConfig object to its proto representation.
func DataprocAlphaClusterConfigGceClusterConfigConfidentialInstanceConfigToProto(o *alpha.ClusterConfigGceClusterConfigConfidentialInstanceConfig) *alphapb.DataprocAlphaClusterConfigGceClusterConfigConfidentialInstanceConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigGceClusterConfigConfidentialInstanceConfig{}
	p.SetEnableConfidentialCompute(dcl.ValueOrEmptyBool(o.EnableConfidentialCompute))
	return p
}

// ClusterConfigMasterConfigToProto converts a ClusterConfigMasterConfig object to its proto representation.
func DataprocAlphaClusterConfigMasterConfigToProto(o *alpha.ClusterConfigMasterConfig) *alphapb.DataprocAlphaClusterConfigMasterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigMasterConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocAlphaClusterConfigMasterConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocAlphaClusterConfigMasterConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocAlphaClusterConfigMasterConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*alphapb.DataprocAlphaClusterConfigMasterConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocAlphaClusterConfigMasterConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	sInstanceReferences := make([]*alphapb.DataprocAlphaClusterConfigMasterConfigInstanceReferences, len(o.InstanceReferences))
	for i, r := range o.InstanceReferences {
		sInstanceReferences[i] = DataprocAlphaClusterConfigMasterConfigInstanceReferencesToProto(&r)
	}
	p.SetInstanceReferences(sInstanceReferences)
	return p
}

// ClusterConfigMasterConfigDiskConfigToProto converts a ClusterConfigMasterConfigDiskConfig object to its proto representation.
func DataprocAlphaClusterConfigMasterConfigDiskConfigToProto(o *alpha.ClusterConfigMasterConfigDiskConfig) *alphapb.DataprocAlphaClusterConfigMasterConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigMasterConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	p.SetLocalSsdInterface(dcl.ValueOrEmptyString(o.LocalSsdInterface))
	return p
}

// ClusterConfigMasterConfigManagedGroupConfigToProto converts a ClusterConfigMasterConfigManagedGroupConfig object to its proto representation.
func DataprocAlphaClusterConfigMasterConfigManagedGroupConfigToProto(o *alpha.ClusterConfigMasterConfigManagedGroupConfig) *alphapb.DataprocAlphaClusterConfigMasterConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigMasterConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigMasterConfigAcceleratorsToProto converts a ClusterConfigMasterConfigAccelerators object to its proto representation.
func DataprocAlphaClusterConfigMasterConfigAcceleratorsToProto(o *alpha.ClusterConfigMasterConfigAccelerators) *alphapb.DataprocAlphaClusterConfigMasterConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigMasterConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigMasterConfigInstanceReferencesToProto converts a ClusterConfigMasterConfigInstanceReferences object to its proto representation.
func DataprocAlphaClusterConfigMasterConfigInstanceReferencesToProto(o *alpha.ClusterConfigMasterConfigInstanceReferences) *alphapb.DataprocAlphaClusterConfigMasterConfigInstanceReferences {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigMasterConfigInstanceReferences{}
	p.SetInstanceName(dcl.ValueOrEmptyString(o.InstanceName))
	p.SetInstanceId(dcl.ValueOrEmptyString(o.InstanceId))
	p.SetPublicKey(dcl.ValueOrEmptyString(o.PublicKey))
	p.SetPublicEciesKey(dcl.ValueOrEmptyString(o.PublicEciesKey))
	return p
}

// ClusterConfigWorkerConfigToProto converts a ClusterConfigWorkerConfig object to its proto representation.
func DataprocAlphaClusterConfigWorkerConfigToProto(o *alpha.ClusterConfigWorkerConfig) *alphapb.DataprocAlphaClusterConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocAlphaClusterConfigWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocAlphaClusterConfigWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocAlphaClusterConfigWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*alphapb.DataprocAlphaClusterConfigWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocAlphaClusterConfigWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	sInstanceReferences := make([]*alphapb.DataprocAlphaClusterConfigWorkerConfigInstanceReferences, len(o.InstanceReferences))
	for i, r := range o.InstanceReferences {
		sInstanceReferences[i] = DataprocAlphaClusterConfigWorkerConfigInstanceReferencesToProto(&r)
	}
	p.SetInstanceReferences(sInstanceReferences)
	return p
}

// ClusterConfigWorkerConfigDiskConfigToProto converts a ClusterConfigWorkerConfigDiskConfig object to its proto representation.
func DataprocAlphaClusterConfigWorkerConfigDiskConfigToProto(o *alpha.ClusterConfigWorkerConfigDiskConfig) *alphapb.DataprocAlphaClusterConfigWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	p.SetLocalSsdInterface(dcl.ValueOrEmptyString(o.LocalSsdInterface))
	return p
}

// ClusterConfigWorkerConfigManagedGroupConfigToProto converts a ClusterConfigWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocAlphaClusterConfigWorkerConfigManagedGroupConfigToProto(o *alpha.ClusterConfigWorkerConfigManagedGroupConfig) *alphapb.DataprocAlphaClusterConfigWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigWorkerConfigAcceleratorsToProto converts a ClusterConfigWorkerConfigAccelerators object to its proto representation.
func DataprocAlphaClusterConfigWorkerConfigAcceleratorsToProto(o *alpha.ClusterConfigWorkerConfigAccelerators) *alphapb.DataprocAlphaClusterConfigWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigWorkerConfigInstanceReferencesToProto converts a ClusterConfigWorkerConfigInstanceReferences object to its proto representation.
func DataprocAlphaClusterConfigWorkerConfigInstanceReferencesToProto(o *alpha.ClusterConfigWorkerConfigInstanceReferences) *alphapb.DataprocAlphaClusterConfigWorkerConfigInstanceReferences {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigWorkerConfigInstanceReferences{}
	p.SetInstanceName(dcl.ValueOrEmptyString(o.InstanceName))
	p.SetInstanceId(dcl.ValueOrEmptyString(o.InstanceId))
	p.SetPublicKey(dcl.ValueOrEmptyString(o.PublicKey))
	p.SetPublicEciesKey(dcl.ValueOrEmptyString(o.PublicEciesKey))
	return p
}

// ClusterConfigSecondaryWorkerConfigToProto converts a ClusterConfigSecondaryWorkerConfig object to its proto representation.
func DataprocAlphaClusterConfigSecondaryWorkerConfigToProto(o *alpha.ClusterConfigSecondaryWorkerConfig) *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocAlphaClusterConfigSecondaryWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocAlphaClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocAlphaClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocAlphaClusterConfigSecondaryWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	sInstanceReferences := make([]*alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferences, len(o.InstanceReferences))
	for i, r := range o.InstanceReferences {
		sInstanceReferences[i] = DataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferencesToProto(&r)
	}
	p.SetInstanceReferences(sInstanceReferences)
	return p
}

// ClusterConfigSecondaryWorkerConfigDiskConfigToProto converts a ClusterConfigSecondaryWorkerConfigDiskConfig object to its proto representation.
func DataprocAlphaClusterConfigSecondaryWorkerConfigDiskConfigToProto(o *alpha.ClusterConfigSecondaryWorkerConfigDiskConfig) *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	p.SetLocalSsdInterface(dcl.ValueOrEmptyString(o.LocalSsdInterface))
	return p
}

// ClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto converts a ClusterConfigSecondaryWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocAlphaClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o *alpha.ClusterConfigSecondaryWorkerConfigManagedGroupConfig) *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// ClusterConfigSecondaryWorkerConfigAcceleratorsToProto converts a ClusterConfigSecondaryWorkerConfigAccelerators object to its proto representation.
func DataprocAlphaClusterConfigSecondaryWorkerConfigAcceleratorsToProto(o *alpha.ClusterConfigSecondaryWorkerConfigAccelerators) *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// ClusterConfigSecondaryWorkerConfigInstanceReferencesToProto converts a ClusterConfigSecondaryWorkerConfigInstanceReferences object to its proto representation.
func DataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferencesToProto(o *alpha.ClusterConfigSecondaryWorkerConfigInstanceReferences) *alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferences {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSecondaryWorkerConfigInstanceReferences{}
	p.SetInstanceName(dcl.ValueOrEmptyString(o.InstanceName))
	p.SetInstanceId(dcl.ValueOrEmptyString(o.InstanceId))
	p.SetPublicKey(dcl.ValueOrEmptyString(o.PublicKey))
	p.SetPublicEciesKey(dcl.ValueOrEmptyString(o.PublicEciesKey))
	return p
}

// ClusterConfigSoftwareConfigToProto converts a ClusterConfigSoftwareConfig object to its proto representation.
func DataprocAlphaClusterConfigSoftwareConfigToProto(o *alpha.ClusterConfigSoftwareConfig) *alphapb.DataprocAlphaClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum(alphapb.DataprocAlphaClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// ClusterConfigInitializationActionsToProto converts a ClusterConfigInitializationActions object to its proto representation.
func DataprocAlphaClusterConfigInitializationActionsToProto(o *alpha.ClusterConfigInitializationActions) *alphapb.DataprocAlphaClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// ClusterConfigEncryptionConfigToProto converts a ClusterConfigEncryptionConfig object to its proto representation.
func DataprocAlphaClusterConfigEncryptionConfigToProto(o *alpha.ClusterConfigEncryptionConfig) *alphapb.DataprocAlphaClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// ClusterConfigAutoscalingConfigToProto converts a ClusterConfigAutoscalingConfig object to its proto representation.
func DataprocAlphaClusterConfigAutoscalingConfigToProto(o *alpha.ClusterConfigAutoscalingConfig) *alphapb.DataprocAlphaClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// ClusterConfigSecurityConfigToProto converts a ClusterConfigSecurityConfig object to its proto representation.
func DataprocAlphaClusterConfigSecurityConfigToProto(o *alpha.ClusterConfigSecurityConfig) *alphapb.DataprocAlphaClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocAlphaClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	p.SetIdentityConfig(DataprocAlphaClusterConfigSecurityConfigIdentityConfigToProto(o.IdentityConfig))
	return p
}

// ClusterConfigSecurityConfigKerberosConfigToProto converts a ClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocAlphaClusterConfigSecurityConfigKerberosConfigToProto(o *alpha.ClusterConfigSecurityConfigKerberosConfig) *alphapb.DataprocAlphaClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSecurityConfigKerberosConfig{}
	p.SetEnableKerberos(dcl.ValueOrEmptyBool(o.EnableKerberos))
	p.SetRootPrincipalPassword(dcl.ValueOrEmptyString(o.RootPrincipalPassword))
	p.SetKmsKey(dcl.ValueOrEmptyString(o.KmsKey))
	p.SetKeystore(dcl.ValueOrEmptyString(o.Keystore))
	p.SetTruststore(dcl.ValueOrEmptyString(o.Truststore))
	p.SetKeystorePassword(dcl.ValueOrEmptyString(o.KeystorePassword))
	p.SetKeyPassword(dcl.ValueOrEmptyString(o.KeyPassword))
	p.SetTruststorePassword(dcl.ValueOrEmptyString(o.TruststorePassword))
	p.SetCrossRealmTrustRealm(dcl.ValueOrEmptyString(o.CrossRealmTrustRealm))
	p.SetCrossRealmTrustKdc(dcl.ValueOrEmptyString(o.CrossRealmTrustKdc))
	p.SetCrossRealmTrustAdminServer(dcl.ValueOrEmptyString(o.CrossRealmTrustAdminServer))
	p.SetCrossRealmTrustSharedPassword(dcl.ValueOrEmptyString(o.CrossRealmTrustSharedPassword))
	p.SetKdcDbKey(dcl.ValueOrEmptyString(o.KdcDbKey))
	p.SetTgtLifetimeHours(dcl.ValueOrEmptyInt64(o.TgtLifetimeHours))
	p.SetRealm(dcl.ValueOrEmptyString(o.Realm))
	return p
}

// ClusterConfigSecurityConfigIdentityConfigToProto converts a ClusterConfigSecurityConfigIdentityConfig object to its proto representation.
func DataprocAlphaClusterConfigSecurityConfigIdentityConfigToProto(o *alpha.ClusterConfigSecurityConfigIdentityConfig) *alphapb.DataprocAlphaClusterConfigSecurityConfigIdentityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigSecurityConfigIdentityConfig{}
	mUserServiceAccountMapping := make(map[string]string, len(o.UserServiceAccountMapping))
	for k, r := range o.UserServiceAccountMapping {
		mUserServiceAccountMapping[k] = r
	}
	p.SetUserServiceAccountMapping(mUserServiceAccountMapping)
	return p
}

// ClusterConfigLifecycleConfigToProto converts a ClusterConfigLifecycleConfig object to its proto representation.
func DataprocAlphaClusterConfigLifecycleConfigToProto(o *alpha.ClusterConfigLifecycleConfig) *alphapb.DataprocAlphaClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// ClusterConfigEndpointConfigToProto converts a ClusterConfigEndpointConfig object to its proto representation.
func DataprocAlphaClusterConfigEndpointConfigToProto(o *alpha.ClusterConfigEndpointConfig) *alphapb.DataprocAlphaClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigEndpointConfig{}
	p.SetEnableHttpPortAccess(dcl.ValueOrEmptyBool(o.EnableHttpPortAccess))
	mHttpPorts := make(map[string]string, len(o.HttpPorts))
	for k, r := range o.HttpPorts {
		mHttpPorts[k] = r
	}
	p.SetHttpPorts(mHttpPorts)
	return p
}

// ClusterConfigGkeClusterConfigToProto converts a ClusterConfigGkeClusterConfig object to its proto representation.
func DataprocAlphaClusterConfigGkeClusterConfigToProto(o *alpha.ClusterConfigGkeClusterConfig) *alphapb.DataprocAlphaClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigGkeClusterConfig{}
	p.SetNamespacedGkeDeploymentTarget(DataprocAlphaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o.NamespacedGkeDeploymentTarget))
	return p
}

// ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto converts a ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object to its proto representation.
func DataprocAlphaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o *alpha.ClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *alphapb.DataprocAlphaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
	p.SetTargetGkeCluster(dcl.ValueOrEmptyString(o.TargetGkeCluster))
	p.SetClusterNamespace(dcl.ValueOrEmptyString(o.ClusterNamespace))
	return p
}

// ClusterConfigMetastoreConfigToProto converts a ClusterConfigMetastoreConfig object to its proto representation.
func DataprocAlphaClusterConfigMetastoreConfigToProto(o *alpha.ClusterConfigMetastoreConfig) *alphapb.DataprocAlphaClusterConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigMetastoreConfig{}
	p.SetDataprocMetastoreService(dcl.ValueOrEmptyString(o.DataprocMetastoreService))
	return p
}

// ClusterConfigDataprocMetricConfigToProto converts a ClusterConfigDataprocMetricConfig object to its proto representation.
func DataprocAlphaClusterConfigDataprocMetricConfigToProto(o *alpha.ClusterConfigDataprocMetricConfig) *alphapb.DataprocAlphaClusterConfigDataprocMetricConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigDataprocMetricConfig{}
	sMetrics := make([]*alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetrics, len(o.Metrics))
	for i, r := range o.Metrics {
		sMetrics[i] = DataprocAlphaClusterConfigDataprocMetricConfigMetricsToProto(&r)
	}
	p.SetMetrics(sMetrics)
	return p
}

// ClusterConfigDataprocMetricConfigMetricsToProto converts a ClusterConfigDataprocMetricConfigMetrics object to its proto representation.
func DataprocAlphaClusterConfigDataprocMetricConfigMetricsToProto(o *alpha.ClusterConfigDataprocMetricConfigMetrics) *alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetrics {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterConfigDataprocMetricConfigMetrics{}
	p.SetMetricSource(DataprocAlphaClusterConfigDataprocMetricConfigMetricsMetricSourceEnumToProto(o.MetricSource))
	sMetricOverrides := make([]string, len(o.MetricOverrides))
	for i, r := range o.MetricOverrides {
		sMetricOverrides[i] = r
	}
	p.SetMetricOverrides(sMetricOverrides)
	return p
}

// ClusterStatusToProto converts a ClusterStatus object to its proto representation.
func DataprocAlphaClusterStatusToProto(o *alpha.ClusterStatus) *alphapb.DataprocAlphaClusterStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterStatus{}
	p.SetState(DataprocAlphaClusterStatusStateEnumToProto(o.State))
	p.SetDetail(dcl.ValueOrEmptyString(o.Detail))
	p.SetStateStartTime(dcl.ValueOrEmptyString(o.StateStartTime))
	p.SetSubstate(DataprocAlphaClusterStatusSubstateEnumToProto(o.Substate))
	return p
}

// ClusterStatusHistoryToProto converts a ClusterStatusHistory object to its proto representation.
func DataprocAlphaClusterStatusHistoryToProto(o *alpha.ClusterStatusHistory) *alphapb.DataprocAlphaClusterStatusHistory {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterStatusHistory{}
	p.SetState(DataprocAlphaClusterStatusHistoryStateEnumToProto(o.State))
	p.SetDetail(dcl.ValueOrEmptyString(o.Detail))
	p.SetStateStartTime(dcl.ValueOrEmptyString(o.StateStartTime))
	p.SetSubstate(DataprocAlphaClusterStatusHistorySubstateEnumToProto(o.Substate))
	return p
}

// ClusterMetricsToProto converts a ClusterMetrics object to its proto representation.
func DataprocAlphaClusterMetricsToProto(o *alpha.ClusterMetrics) *alphapb.DataprocAlphaClusterMetrics {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterMetrics{}
	mHdfsMetrics := make(map[string]string, len(o.HdfsMetrics))
	for k, r := range o.HdfsMetrics {
		mHdfsMetrics[k] = r
	}
	p.SetHdfsMetrics(mHdfsMetrics)
	mYarnMetrics := make(map[string]string, len(o.YarnMetrics))
	for k, r := range o.YarnMetrics {
		mYarnMetrics[k] = r
	}
	p.SetYarnMetrics(mYarnMetrics)
	return p
}

// ClusterVirtualClusterConfigToProto converts a ClusterVirtualClusterConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigToProto(o *alpha.ClusterVirtualClusterConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetKubernetesClusterConfig(DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigToProto(o.KubernetesClusterConfig))
	p.SetAuxiliaryServicesConfig(DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigToProto(o.AuxiliaryServicesConfig))
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfig{}
	p.SetKubernetesNamespace(dcl.ValueOrEmptyString(o.KubernetesNamespace))
	p.SetGkeClusterConfig(DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig))
	p.SetKubernetesSoftwareConfig(DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigToProto(o.KubernetesSoftwareConfig))
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig{}
	p.SetGkeClusterTarget(dcl.ValueOrEmptyString(o.GkeClusterTarget))
	sNodePoolTarget := make([]*alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget, len(o.NodePoolTarget))
	for i, r := range o.NodePoolTarget {
		sNodePoolTarget[i] = DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetToProto(&r)
	}
	p.SetNodePoolTarget(sNodePoolTarget)
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget{}
	p.SetNodePool(dcl.ValueOrEmptyString(o.NodePool))
	p.SetNodePoolConfig(DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigToProto(o.NodePoolConfig))
	sRoles := make([]alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum, len(o.Roles))
	for i, r := range o.Roles {
		sRoles[i] = alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum(alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetRolesEnum_value[string(r)])
	}
	p.SetRoles(sRoles)
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig{}
	p.SetConfig(DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigToProto(o.Config))
	p.SetAutoscaling(DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscalingToProto(o.Autoscaling))
	sLocations := make([]string, len(o.Locations))
	for i, r := range o.Locations {
		sLocations[i] = r
	}
	p.SetLocations(sLocations)
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetLocalSsdCount(dcl.ValueOrEmptyInt64(o.LocalSsdCount))
	p.SetPreemptible(dcl.ValueOrEmptyBool(o.Preemptible))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	p.SetBootDiskKmsKey(dcl.ValueOrEmptyString(o.BootDiskKmsKey))
	p.SetEphemeralStorageConfig(DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfigToProto(o.EphemeralStorageConfig))
	p.SetSpot(dcl.ValueOrEmptyBool(o.Spot))
	sAccelerators := make([]*alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAcceleratorsToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAcceleratorsToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigAccelerators{}
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetGpuPartitionSize(dcl.ValueOrEmptyString(o.GpuPartitionSize))
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfigToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfigToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfigEphemeralStorageConfig{}
	p.SetLocalSsdCount(dcl.ValueOrEmptyInt64(o.LocalSsdCount))
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscalingToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscalingToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling{}
	p.SetMinNodeCount(dcl.ValueOrEmptyInt64(o.MinNodeCount))
	p.SetMaxNodeCount(dcl.ValueOrEmptyInt64(o.MaxNodeCount))
	return p
}

// ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigToProto converts a ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigToProto(o *alpha.ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig{}
	mComponentVersion := make(map[string]string, len(o.ComponentVersion))
	for k, r := range o.ComponentVersion {
		mComponentVersion[k] = r
	}
	p.SetComponentVersion(mComponentVersion)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// ClusterVirtualClusterConfigAuxiliaryServicesConfigToProto converts a ClusterVirtualClusterConfigAuxiliaryServicesConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigToProto(o *alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfig{}
	p.SetMetastoreConfig(DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfigToProto(o.MetastoreConfig))
	p.SetSparkHistoryServerConfig(DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfigToProto(o.SparkHistoryServerConfig))
	return p
}

// ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfigToProto converts a ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfigToProto(o *alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig{}
	p.SetDataprocMetastoreService(dcl.ValueOrEmptyString(o.DataprocMetastoreService))
	return p
}

// ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfigToProto converts a ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig object to its proto representation.
func DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfigToProto(o *alpha.ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig) *alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig{}
	p.SetDataprocCluster(dcl.ValueOrEmptyString(o.DataprocCluster))
	return p
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *alpha.Cluster) *alphapb.DataprocAlphaCluster {
	p := &alphapb.DataprocAlphaCluster{}
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetConfig(DataprocAlphaClusterConfigToProto(resource.Config))
	p.SetStatus(DataprocAlphaClusterStatusToProto(resource.Status))
	p.SetClusterUuid(dcl.ValueOrEmptyString(resource.ClusterUuid))
	p.SetMetrics(DataprocAlphaClusterMetricsToProto(resource.Metrics))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetVirtualClusterConfig(DataprocAlphaClusterVirtualClusterConfigToProto(resource.VirtualClusterConfig))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sStatusHistory := make([]*alphapb.DataprocAlphaClusterStatusHistory, len(resource.StatusHistory))
	for i, r := range resource.StatusHistory {
		sStatusHistory[i] = DataprocAlphaClusterStatusHistoryToProto(&r)
	}
	p.SetStatusHistory(sStatusHistory)

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataprocAlphaClusterRequest) (*alphapb.DataprocAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyDataprocAlphaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyDataprocAlphaCluster(ctx context.Context, request *alphapb.ApplyDataprocAlphaClusterRequest) (*alphapb.DataprocAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteDataprocAlphaCluster(ctx context.Context, request *alphapb.DeleteDataprocAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListDataprocAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListDataprocAlphaCluster(ctx context.Context, request *alphapb.ListDataprocAlphaClusterRequest) (*alphapb.ListDataprocAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataprocAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDataprocAlphaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
