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

// WorkflowTemplateServer implements the gRPC interface for WorkflowTemplate.
type WorkflowTemplateServer struct{}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(e alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum) *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(n[len("DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(e alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum) *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(n[len("DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(e alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum) *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(n[len("DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(e alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum) *alpha.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := alpha.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacement(p *alphapb.DataprocAlphaWorkflowTemplatePlacement) *alpha.WorkflowTemplatePlacement {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacement{
		ManagedCluster:  ProtoToDataprocAlphaWorkflowTemplatePlacementManagedCluster(p.GetManagedCluster()),
		ClusterSelector: ProtoToDataprocAlphaWorkflowTemplatePlacementClusterSelector(p.GetClusterSelector()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedCluster(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster) *alpha.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.GetClusterName()),
		Config:      ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfig converts a WorkflowTemplatePlacementManagedClusterConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
		GceClusterConfig:      ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig(p.GetEndpointConfig()),
		GkeClusterConfig:      ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig(p.GetGkeClusterConfig()),
		MetastoreConfig:       ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig(p.GetMetastoreConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.GetZone()),
		Network:                 dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:              dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:          dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess: ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
		ReservationAffinity:     ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:       ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
		ShieldedInstanceConfig:  ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
	}
	for _, r := range p.GetServiceAccountScopes() {
		obj.ServiceAccountScopes = append(obj.ServiceAccountScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity) *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity) *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.GetEnableSecureBoot()),
		EnableVtpm:                dcl.Bool(p.GetEnableVtpm()),
		EnableIntegrityMonitoring: dcl.Bool(p.GetEnableIntegrityMonitoring()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators) *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators) *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators) *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigInitializationActions converts a WorkflowTemplatePlacementManagedClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActions(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActions) *alpha.WorkflowTemplatePlacementManagedClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig converts a WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig converts a WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecurityConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig{
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

// ProtoToWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig converts a WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigEndpointConfig converts a WorkflowTemplatePlacementManagedClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.GetEnableHttpPortAccess()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig converts a WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p.GetNamespacedGkeDeploymentTarget()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget converts a WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *alpha.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.StringOrNil(p.GetTargetGkeCluster()),
		ClusterNamespace: dcl.StringOrNil(p.GetClusterNamespace()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig converts a WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig(p *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig) *alpha.WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.StringOrNil(p.GetDataprocMetastoreService()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplatePlacementClusterSelector(p *alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector) *alpha.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobs(p *alphapb.DataprocAlphaWorkflowTemplateJobs) *alpha.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.GetStepId()),
		HadoopJob:   ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJob(p.GetHadoopJob()),
		SparkJob:    ProtoToDataprocAlphaWorkflowTemplateJobsSparkJob(p.GetSparkJob()),
		PysparkJob:  ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJob(p.GetPysparkJob()),
		HiveJob:     ProtoToDataprocAlphaWorkflowTemplateJobsHiveJob(p.GetHiveJob()),
		PigJob:      ProtoToDataprocAlphaWorkflowTemplateJobsPigJob(p.GetPigJob()),
		SparkRJob:   ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJob(p.GetSparkRJob()),
		SparkSqlJob: ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJob(p.GetSparkSqlJob()),
		PrestoJob:   ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJob(p.GetPrestoJob()),
		Scheduling:  ProtoToDataprocAlphaWorkflowTemplateJobsScheduling(p.GetScheduling()),
	}
	for _, r := range p.GetPrerequisiteStepIds() {
		obj.PrerequisiteStepIds = append(obj.PrerequisiteStepIds, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob) *alpha.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
		LoggingConfig:  ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJobLoggingConfig converts a WorkflowTemplateJobsHadoopJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig) *alpha.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob) *alpha.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
		LoggingConfig:  ProtoToDataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJobLoggingConfig converts a WorkflowTemplateJobsSparkJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob) *alpha.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.GetMainPythonFileUri()),
		LoggingConfig:     ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetPythonFileUris() {
		obj.PythonFileUris = append(obj.PythonFileUris, r)
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJobLoggingConfig converts a WorkflowTemplateJobsPysparkJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig) *alpha.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHiveJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob) *alpha.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsHiveJobQueryList(p *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList) *alpha.WorkflowTemplateJobsHiveJobQueryList {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPigJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPigJob) *alpha.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		LoggingConfig:     ProtoToDataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPigJobQueryList(p *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobQueryList) *alpha.WorkflowTemplateJobsPigJobQueryList {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPigJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig) *alpha.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob) *alpha.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.GetMainRFileUri()),
		LoggingConfig: ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetFileUris() {
		obj.FileUris = append(obj.FileUris, r)
	}
	for _, r := range p.GetArchiveUris() {
		obj.ArchiveUris = append(obj.ArchiveUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJobLoggingConfig converts a WorkflowTemplateJobsSparkRJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob) *alpha.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:     ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList) *alpha.WorkflowTemplateJobsSparkSqlJobQueryList {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig) *alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJob(p *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob) *alpha.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		OutputFormat:      dcl.StringOrNil(p.GetOutputFormat()),
		LoggingConfig:     ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobQueryList(p *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList) *alpha.WorkflowTemplateJobsPrestoJobQueryList {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig(p *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig) *alpha.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateJobsScheduling(p *alphapb.DataprocAlphaWorkflowTemplateJobsScheduling) *alpha.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.GetMaxFailuresPerHour()),
		MaxFailuresTotal:   dcl.Int64OrNil(p.GetMaxFailuresTotal()),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParameters(p *alphapb.DataprocAlphaWorkflowTemplateParameters) *alpha.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Validation:  ProtoToDataprocAlphaWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParametersValidation(p *alphapb.DataprocAlphaWorkflowTemplateParametersValidation) *alpha.WorkflowTemplateParametersValidation {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParametersValidation{
		Regex:  ProtoToDataprocAlphaWorkflowTemplateParametersValidationRegex(p.GetRegex()),
		Values: ProtoToDataprocAlphaWorkflowTemplateParametersValidationValues(p.GetValues()),
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParametersValidationRegex(p *alphapb.DataprocAlphaWorkflowTemplateParametersValidationRegex) *alpha.WorkflowTemplateParametersValidationRegex {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParametersValidationRegex{}
	for _, r := range p.GetRegexes() {
		obj.Regexes = append(obj.Regexes, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues object from its proto representation.
func ProtoToDataprocAlphaWorkflowTemplateParametersValidationValues(p *alphapb.DataprocAlphaWorkflowTemplateParametersValidationValues) *alpha.WorkflowTemplateParametersValidationValues {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkflowTemplateParametersValidationValues{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplate converts a WorkflowTemplate resource from its proto representation.
func ProtoToWorkflowTemplate(p *alphapb.DataprocAlphaWorkflowTemplate) *alpha.WorkflowTemplate {
	obj := &alpha.WorkflowTemplate{
		Name:       dcl.StringOrNil(p.GetName()),
		Version:    dcl.Int64OrNil(p.GetVersion()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocAlphaWorkflowTemplatePlacement(p.GetPlacement()),
		DagTimeout: dcl.StringOrNil(p.GetDagTimeout()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocAlphaWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocAlphaWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum enum to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto(e *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum) alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum {
	if e == nil {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(v)
	}
	return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto(e *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum) alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum {
	if e == nil {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(v)
	}
	return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(e *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum) alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == nil {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(v)
	}
	return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *alpha.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum) alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_value["WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(0)
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementToProto(o *alpha.WorkflowTemplatePlacement) *alphapb.DataprocAlphaWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacement{}
	p.SetManagedCluster(DataprocAlphaWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster))
	p.SetClusterSelector(DataprocAlphaWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector))
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterToProto(o *alpha.WorkflowTemplatePlacementManagedCluster) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedCluster{}
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigToProto(o.Config))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto(o.EndpointConfig))
	p.SetGkeClusterConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig))
	p.SetMetastoreConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigToProto(o.MetastoreConfig))
	sInitializationActions := make([]*alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
	p.SetShieldedInstanceConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig))
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

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig{}
	p.SetEnableSecureBoot(dcl.ValueOrEmptyBool(o.EnableSecureBoot))
	p.SetEnableVtpm(dcl.ValueOrEmptyBool(o.EnableVtpm))
	p.SetEnableIntegrityMonitoring(dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto converts a WorkflowTemplatePlacementManagedClusterConfigInitializationActions object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigInitializationActions) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig{}
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

// WorkflowTemplatePlacementManagedClusterConfigLifecycleConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigEndpointConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig{}
	p.SetEnableHttpPortAccess(dcl.ValueOrEmptyBool(o.EnableHttpPortAccess))
	mHttpPorts := make(map[string]string, len(o.HttpPorts))
	for k, r := range o.HttpPorts {
		mHttpPorts[k] = r
	}
	p.SetHttpPorts(mHttpPorts)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig{}
	p.SetNamespacedGkeDeploymentTarget(DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o.NamespacedGkeDeploymentTarget))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto converts a WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
	p.SetTargetGkeCluster(dcl.ValueOrEmptyString(o.TargetGkeCluster))
	p.SetClusterNamespace(dcl.ValueOrEmptyString(o.ClusterNamespace))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMetastoreConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigToProto(o *alpha.WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig) *alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig{}
	p.SetDataprocMetastoreService(dcl.ValueOrEmptyString(o.DataprocMetastoreService))
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector object to its proto representation.
func DataprocAlphaWorkflowTemplatePlacementClusterSelectorToProto(o *alpha.WorkflowTemplatePlacementClusterSelector) *alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplatePlacementClusterSelector{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	mClusterLabels := make(map[string]string, len(o.ClusterLabels))
	for k, r := range o.ClusterLabels {
		mClusterLabels[k] = r
	}
	p.SetClusterLabels(mClusterLabels)
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsToProto(o *alpha.WorkflowTemplateJobs) *alphapb.DataprocAlphaWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobs{}
	p.SetStepId(dcl.ValueOrEmptyString(o.StepId))
	p.SetHadoopJob(DataprocAlphaWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob))
	p.SetSparkJob(DataprocAlphaWorkflowTemplateJobsSparkJobToProto(o.SparkJob))
	p.SetPysparkJob(DataprocAlphaWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob))
	p.SetHiveJob(DataprocAlphaWorkflowTemplateJobsHiveJobToProto(o.HiveJob))
	p.SetPigJob(DataprocAlphaWorkflowTemplateJobsPigJobToProto(o.PigJob))
	p.SetSparkRJob(DataprocAlphaWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob))
	p.SetSparkSqlJob(DataprocAlphaWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob))
	p.SetPrestoJob(DataprocAlphaWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob))
	p.SetScheduling(DataprocAlphaWorkflowTemplateJobsSchedulingToProto(o.Scheduling))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sPrerequisiteStepIds := make([]string, len(o.PrerequisiteStepIds))
	for i, r := range o.PrerequisiteStepIds {
		sPrerequisiteStepIds[i] = r
	}
	p.SetPrerequisiteStepIds(sPrerequisiteStepIds)
	return p
}

// WorkflowTemplateJobsHadoopJobToProto converts a WorkflowTemplateJobsHadoopJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsHadoopJobToProto(o *alpha.WorkflowTemplateJobsHadoopJob) *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	sFileUris := make([]string, len(o.FileUris))
	for i, r := range o.FileUris {
		sFileUris[i] = r
	}
	p.SetFileUris(sFileUris)
	sArchiveUris := make([]string, len(o.ArchiveUris))
	for i, r := range o.ArchiveUris {
		sArchiveUris[i] = r
	}
	p.SetArchiveUris(sArchiveUris)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsHadoopJobLoggingConfigToProto converts a WorkflowTemplateJobsHadoopJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsHadoopJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHadoopJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkJobToProto(o *alpha.WorkflowTemplateJobsSparkJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	sFileUris := make([]string, len(o.FileUris))
	for i, r := range o.FileUris {
		sFileUris[i] = r
	}
	p.SetFileUris(sFileUris)
	sArchiveUris := make([]string, len(o.ArchiveUris))
	for i, r := range o.ArchiveUris {
		sArchiveUris[i] = r
	}
	p.SetArchiveUris(sArchiveUris)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsSparkJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPysparkJobToProto(o *alpha.WorkflowTemplateJobsPysparkJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJob{}
	p.SetMainPythonFileUri(dcl.ValueOrEmptyString(o.MainPythonFileUri))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sPythonFileUris := make([]string, len(o.PythonFileUris))
	for i, r := range o.PythonFileUris {
		sPythonFileUris[i] = r
	}
	p.SetPythonFileUris(sPythonFileUris)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	sFileUris := make([]string, len(o.FileUris))
	for i, r := range o.FileUris {
		sFileUris[i] = r
	}
	p.SetFileUris(sFileUris)
	sArchiveUris := make([]string, len(o.ArchiveUris))
	for i, r := range o.ArchiveUris {
		sArchiveUris[i] = r
	}
	p.SetArchiveUris(sArchiveUris)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsPysparkJobLoggingConfigToProto converts a WorkflowTemplateJobsPysparkJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPysparkJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPysparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsHiveJobToProto(o *alpha.WorkflowTemplateJobsHiveJob) *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHiveJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocAlphaWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	mScriptVariables := make(map[string]string, len(o.ScriptVariables))
	for k, r := range o.ScriptVariables {
		mScriptVariables[k] = r
	}
	p.SetScriptVariables(mScriptVariables)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	return p
}

// WorkflowTemplateJobsHiveJobQueryListToProto converts a WorkflowTemplateJobsHiveJobQueryList object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsHiveJobQueryListToProto(o *alpha.WorkflowTemplateJobsHiveJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsHiveJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPigJobToProto(o *alpha.WorkflowTemplateJobsPigJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocAlphaWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig))
	mScriptVariables := make(map[string]string, len(o.ScriptVariables))
	for k, r := range o.ScriptVariables {
		mScriptVariables[k] = r
	}
	p.SetScriptVariables(mScriptVariables)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	return p
}

// WorkflowTemplateJobsPigJobQueryListToProto converts a WorkflowTemplateJobsPigJobQueryList object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPigJobQueryListToProto(o *alpha.WorkflowTemplateJobsPigJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPigJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPigJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkRJobToProto(o *alpha.WorkflowTemplateJobsSparkRJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJob{}
	p.SetMainRFileUri(dcl.ValueOrEmptyString(o.MainRFileUri))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig))
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sFileUris := make([]string, len(o.FileUris))
	for i, r := range o.FileUris {
		sFileUris[i] = r
	}
	p.SetFileUris(sFileUris)
	sArchiveUris := make([]string, len(o.ArchiveUris))
	for i, r := range o.ArchiveUris {
		sArchiveUris[i] = r
	}
	p.SetArchiveUris(sArchiveUris)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsSparkRJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkRJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkRJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkRJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobToProto(o *alpha.WorkflowTemplateJobsSparkSqlJob) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig))
	mScriptVariables := make(map[string]string, len(o.ScriptVariables))
	for k, r := range o.ScriptVariables {
		mScriptVariables[k] = r
	}
	p.SetScriptVariables(mScriptVariables)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sJarFileUris := make([]string, len(o.JarFileUris))
	for i, r := range o.JarFileUris {
		sJarFileUris[i] = r
	}
	p.SetJarFileUris(sJarFileUris)
	return p
}

// WorkflowTemplateJobsSparkSqlJobQueryListToProto converts a WorkflowTemplateJobsSparkSqlJobQueryList object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *alpha.WorkflowTemplateJobsSparkSqlJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPrestoJobToProto(o *alpha.WorkflowTemplateJobsPrestoJob) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocAlphaWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetOutputFormat(dcl.ValueOrEmptyString(o.OutputFormat))
	p.SetLoggingConfig(DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig))
	sClientTags := make([]string, len(o.ClientTags))
	for i, r := range o.ClientTags {
		sClientTags[i] = r
	}
	p.SetClientTags(sClientTags)
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	return p
}

// WorkflowTemplateJobsPrestoJobQueryListToProto converts a WorkflowTemplateJobsPrestoJobQueryList object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPrestoJobQueryListToProto(o *alpha.WorkflowTemplateJobsPrestoJobQueryList) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *alpha.WorkflowTemplateJobsPrestoJobLoggingConfig) *alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsPrestoJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling object to its proto representation.
func DataprocAlphaWorkflowTemplateJobsSchedulingToProto(o *alpha.WorkflowTemplateJobsScheduling) *alphapb.DataprocAlphaWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateJobsScheduling{}
	p.SetMaxFailuresPerHour(dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour))
	p.SetMaxFailuresTotal(dcl.ValueOrEmptyInt64(o.MaxFailuresTotal))
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters object to its proto representation.
func DataprocAlphaWorkflowTemplateParametersToProto(o *alpha.WorkflowTemplateParameters) *alphapb.DataprocAlphaWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParameters{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetValidation(DataprocAlphaWorkflowTemplateParametersValidationToProto(o.Validation))
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation object to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationToProto(o *alpha.WorkflowTemplateParametersValidation) *alphapb.DataprocAlphaWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidation{}
	p.SetRegex(DataprocAlphaWorkflowTemplateParametersValidationRegexToProto(o.Regex))
	p.SetValues(DataprocAlphaWorkflowTemplateParametersValidationValuesToProto(o.Values))
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex object to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationRegexToProto(o *alpha.WorkflowTemplateParametersValidationRegex) *alphapb.DataprocAlphaWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidationRegex{}
	sRegexes := make([]string, len(o.Regexes))
	for i, r := range o.Regexes {
		sRegexes[i] = r
	}
	p.SetRegexes(sRegexes)
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues object to its proto representation.
func DataprocAlphaWorkflowTemplateParametersValidationValuesToProto(o *alpha.WorkflowTemplateParametersValidationValues) *alphapb.DataprocAlphaWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &alphapb.DataprocAlphaWorkflowTemplateParametersValidationValues{}
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *alpha.WorkflowTemplate) *alphapb.DataprocAlphaWorkflowTemplate {
	p := &alphapb.DataprocAlphaWorkflowTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyInt64(resource.Version))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetPlacement(DataprocAlphaWorkflowTemplatePlacementToProto(resource.Placement))
	p.SetDagTimeout(dcl.ValueOrEmptyString(resource.DagTimeout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sJobs := make([]*alphapb.DataprocAlphaWorkflowTemplateJobs, len(resource.Jobs))
	for i, r := range resource.Jobs {
		sJobs[i] = DataprocAlphaWorkflowTemplateJobsToProto(&r)
	}
	p.SetJobs(sJobs)
	sParameters := make([]*alphapb.DataprocAlphaWorkflowTemplateParameters, len(resource.Parameters))
	for i, r := range resource.Parameters {
		sParameters[i] = DataprocAlphaWorkflowTemplateParametersToProto(&r)
	}
	p.SetParameters(sParameters)

	return p
}

// applyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *alpha.Client, request *alphapb.ApplyDataprocAlphaWorkflowTemplateRequest) (*alphapb.DataprocAlphaWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// applyDataprocAlphaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.ApplyDataprocAlphaWorkflowTemplateRequest) (*alphapb.DataprocAlphaWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.DeleteDataprocAlphaWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListDataprocAlphaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocAlphaWorkflowTemplate(ctx context.Context, request *alphapb.ListDataprocAlphaWorkflowTemplateRequest) (*alphapb.ListDataprocAlphaWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.DataprocAlphaWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListDataprocAlphaWorkflowTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
