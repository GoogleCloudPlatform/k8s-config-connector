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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/beta/dataproc_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
)

// WorkflowTemplateServer implements the gRPC interface for WorkflowTemplate.
type WorkflowTemplateServer struct{}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(e betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum) *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(n[len("DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(e betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum) *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(n[len("DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(e betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum) *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(n[len("DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(e betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum) *beta.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := beta.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacement(p *betapb.DataprocBetaWorkflowTemplatePlacement) *beta.WorkflowTemplatePlacement {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacement{
		ManagedCluster:  ProtoToDataprocBetaWorkflowTemplatePlacementManagedCluster(p.GetManagedCluster()),
		ClusterSelector: ProtoToDataprocBetaWorkflowTemplatePlacementClusterSelector(p.GetClusterSelector()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedCluster(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster) *beta.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.GetClusterName()),
		Config:      ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfig converts a WorkflowTemplatePlacementManagedClusterConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfig) *beta.WorkflowTemplatePlacementManagedClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
		GceClusterConfig:      ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig(p.GetEndpointConfig()),
		GkeClusterConfig:      ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig(p.GetGkeClusterConfig()),
		MetastoreConfig:       ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig(p.GetMetastoreConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.GetZone()),
		Network:                 dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:              dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:          dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess: ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
		ReservationAffinity:     ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:       ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
		ShieldedInstanceConfig:  ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
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
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity) *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity) *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.GetEnableSecureBoot()),
		EnableVtpm:                dcl.Bool(p.GetEnableVtpm()),
		EnableIntegrityMonitoring: dcl.Bool(p.GetEnableIntegrityMonitoring()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators) *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators) *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators) *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigInitializationActions converts a WorkflowTemplatePlacementManagedClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigInitializationActions(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigInitializationActions) *beta.WorkflowTemplatePlacementManagedClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig converts a WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig converts a WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecurityConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig{
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
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigEndpointConfig converts a WorkflowTemplatePlacementManagedClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.GetEnableHttpPortAccess()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig converts a WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p.GetNamespacedGkeDeploymentTarget()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget converts a WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *beta.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		TargetGkeCluster: dcl.StringOrNil(p.GetTargetGkeCluster()),
		ClusterNamespace: dcl.StringOrNil(p.GetClusterNamespace()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig converts a WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig(p *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig) *beta.WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.StringOrNil(p.GetDataprocMetastoreService()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplatePlacementClusterSelector(p *betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector) *beta.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobs(p *betapb.DataprocBetaWorkflowTemplateJobs) *beta.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.GetStepId()),
		HadoopJob:   ProtoToDataprocBetaWorkflowTemplateJobsHadoopJob(p.GetHadoopJob()),
		SparkJob:    ProtoToDataprocBetaWorkflowTemplateJobsSparkJob(p.GetSparkJob()),
		PysparkJob:  ProtoToDataprocBetaWorkflowTemplateJobsPysparkJob(p.GetPysparkJob()),
		HiveJob:     ProtoToDataprocBetaWorkflowTemplateJobsHiveJob(p.GetHiveJob()),
		PigJob:      ProtoToDataprocBetaWorkflowTemplateJobsPigJob(p.GetPigJob()),
		SparkRJob:   ProtoToDataprocBetaWorkflowTemplateJobsSparkRJob(p.GetSparkRJob()),
		SparkSqlJob: ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJob(p.GetSparkSqlJob()),
		PrestoJob:   ProtoToDataprocBetaWorkflowTemplateJobsPrestoJob(p.GetPrestoJob()),
		Scheduling:  ProtoToDataprocBetaWorkflowTemplateJobsScheduling(p.GetScheduling()),
	}
	for _, r := range p.GetPrerequisiteStepIds() {
		obj.PrerequisiteStepIds = append(obj.PrerequisiteStepIds, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHadoopJob(p *betapb.DataprocBetaWorkflowTemplateJobsHadoopJob) *beta.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
		LoggingConfig:  ProtoToDataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig) *beta.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkJob) *beta.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
		LoggingConfig:  ProtoToDataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig) *beta.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPysparkJob(p *betapb.DataprocBetaWorkflowTemplateJobsPysparkJob) *beta.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.GetMainPythonFileUri()),
		LoggingConfig:     ProtoToDataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig) *beta.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHiveJob(p *betapb.DataprocBetaWorkflowTemplateJobsHiveJob) *beta.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsHiveJobQueryList(p *betapb.DataprocBetaWorkflowTemplateJobsHiveJobQueryList) *beta.WorkflowTemplateJobsHiveJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPigJob(p *betapb.DataprocBetaWorkflowTemplateJobsPigJob) *beta.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		LoggingConfig:     ProtoToDataprocBetaWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPigJobQueryList(p *betapb.DataprocBetaWorkflowTemplateJobsPigJobQueryList) *beta.WorkflowTemplateJobsPigJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPigJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPigJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig) *beta.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkRJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkRJob) *beta.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.GetMainRFileUri()),
		LoggingConfig: ProtoToDataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig) *beta.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJob(p *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob) *beta.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:     ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList(p *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList) *beta.WorkflowTemplateJobsSparkSqlJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig) *beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPrestoJob(p *betapb.DataprocBetaWorkflowTemplateJobsPrestoJob) *beta.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		OutputFormat:      dcl.StringOrNil(p.GetOutputFormat()),
		LoggingConfig:     ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobQueryList(p *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobQueryList) *beta.WorkflowTemplateJobsPrestoJobQueryList {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig(p *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig) *beta.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateJobsScheduling(p *betapb.DataprocBetaWorkflowTemplateJobsScheduling) *beta.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.GetMaxFailuresPerHour()),
		MaxFailuresTotal:   dcl.Int64OrNil(p.GetMaxFailuresTotal()),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParameters(p *betapb.DataprocBetaWorkflowTemplateParameters) *beta.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Validation:  ProtoToDataprocBetaWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParametersValidation(p *betapb.DataprocBetaWorkflowTemplateParametersValidation) *beta.WorkflowTemplateParametersValidation {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParametersValidation{
		Regex:  ProtoToDataprocBetaWorkflowTemplateParametersValidationRegex(p.GetRegex()),
		Values: ProtoToDataprocBetaWorkflowTemplateParametersValidationValues(p.GetValues()),
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParametersValidationRegex(p *betapb.DataprocBetaWorkflowTemplateParametersValidationRegex) *beta.WorkflowTemplateParametersValidationRegex {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParametersValidationRegex{}
	for _, r := range p.GetRegexes() {
		obj.Regexes = append(obj.Regexes, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues object from its proto representation.
func ProtoToDataprocBetaWorkflowTemplateParametersValidationValues(p *betapb.DataprocBetaWorkflowTemplateParametersValidationValues) *beta.WorkflowTemplateParametersValidationValues {
	if p == nil {
		return nil
	}
	obj := &beta.WorkflowTemplateParametersValidationValues{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplate converts a WorkflowTemplate resource from its proto representation.
func ProtoToWorkflowTemplate(p *betapb.DataprocBetaWorkflowTemplate) *beta.WorkflowTemplate {
	obj := &beta.WorkflowTemplate{
		Name:       dcl.StringOrNil(p.GetName()),
		Version:    dcl.Int64OrNil(p.GetVersion()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocBetaWorkflowTemplatePlacement(p.GetPlacement()),
		DagTimeout: dcl.StringOrNil(p.GetDagTimeout()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocBetaWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocBetaWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum enum to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto(e *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum) betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum {
	if e == nil {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(0)
	}
	if v, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum"+string(*e)]; ok {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(v)
	}
	return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto(e *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum) betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum {
	if e == nil {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(v)
	}
	return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(e *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum) betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == nil {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(v)
	}
	return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *beta.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum) betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_value["WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(0)
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementToProto(o *beta.WorkflowTemplatePlacement) *betapb.DataprocBetaWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacement{}
	p.SetManagedCluster(DataprocBetaWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster))
	p.SetClusterSelector(DataprocBetaWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector))
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterToProto(o *beta.WorkflowTemplatePlacementManagedCluster) *betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedCluster{}
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigToProto(o.Config))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto(o.EndpointConfig))
	p.SetGkeClusterConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigToProto(o.GkeClusterConfig))
	p.SetMetastoreConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigToProto(o.MetastoreConfig))
	sInitializationActions := make([]*betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocBetaWorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
	p.SetShieldedInstanceConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig))
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
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig{}
	p.SetEnableSecureBoot(dcl.ValueOrEmptyBool(o.EnableSecureBoot))
	p.SetEnableVtpm(dcl.ValueOrEmptyBool(o.EnableVtpm))
	p.SetEnableIntegrityMonitoring(dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto converts a WorkflowTemplatePlacementManagedClusterConfigInitializationActions object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigInitializationActions) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig{}
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
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigEndpointConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigEndpointConfig{}
	p.SetEnableHttpPortAccess(dcl.ValueOrEmptyBool(o.EnableHttpPortAccess))
	mHttpPorts := make(map[string]string, len(o.HttpPorts))
	for k, r := range o.HttpPorts {
		mHttpPorts[k] = r
	}
	p.SetHttpPorts(mHttpPorts)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig{}
	p.SetNamespacedGkeDeploymentTarget(DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o.NamespacedGkeDeploymentTarget))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto converts a WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{}
	p.SetTargetGkeCluster(dcl.ValueOrEmptyString(o.TargetGkeCluster))
	p.SetClusterNamespace(dcl.ValueOrEmptyString(o.ClusterNamespace))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMetastoreConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigToProto(o *beta.WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig) *betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig{}
	p.SetDataprocMetastoreService(dcl.ValueOrEmptyString(o.DataprocMetastoreService))
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector object to its proto representation.
func DataprocBetaWorkflowTemplatePlacementClusterSelectorToProto(o *beta.WorkflowTemplatePlacementClusterSelector) *betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplatePlacementClusterSelector{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	mClusterLabels := make(map[string]string, len(o.ClusterLabels))
	for k, r := range o.ClusterLabels {
		mClusterLabels[k] = r
	}
	p.SetClusterLabels(mClusterLabels)
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs object to its proto representation.
func DataprocBetaWorkflowTemplateJobsToProto(o *beta.WorkflowTemplateJobs) *betapb.DataprocBetaWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobs{}
	p.SetStepId(dcl.ValueOrEmptyString(o.StepId))
	p.SetHadoopJob(DataprocBetaWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob))
	p.SetSparkJob(DataprocBetaWorkflowTemplateJobsSparkJobToProto(o.SparkJob))
	p.SetPysparkJob(DataprocBetaWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob))
	p.SetHiveJob(DataprocBetaWorkflowTemplateJobsHiveJobToProto(o.HiveJob))
	p.SetPigJob(DataprocBetaWorkflowTemplateJobsPigJobToProto(o.PigJob))
	p.SetSparkRJob(DataprocBetaWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob))
	p.SetSparkSqlJob(DataprocBetaWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob))
	p.SetPrestoJob(DataprocBetaWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob))
	p.SetScheduling(DataprocBetaWorkflowTemplateJobsSchedulingToProto(o.Scheduling))
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
func DataprocBetaWorkflowTemplateJobsHadoopJobToProto(o *beta.WorkflowTemplateJobsHadoopJob) *betapb.DataprocBetaWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHadoopJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsHadoopJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHadoopJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkJobToProto(o *beta.WorkflowTemplateJobsSparkJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPysparkJobToProto(o *beta.WorkflowTemplateJobsPysparkJob) *betapb.DataprocBetaWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPysparkJob{}
	p.SetMainPythonFileUri(dcl.ValueOrEmptyString(o.MainPythonFileUri))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPysparkJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPysparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsHiveJobToProto(o *beta.WorkflowTemplateJobsHiveJob) *betapb.DataprocBetaWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHiveJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocBetaWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList))
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
func DataprocBetaWorkflowTemplateJobsHiveJobQueryListToProto(o *beta.WorkflowTemplateJobsHiveJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsHiveJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPigJobToProto(o *beta.WorkflowTemplateJobsPigJob) *betapb.DataprocBetaWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocBetaWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocBetaWorkflowTemplateJobsPigJobQueryListToProto(o *beta.WorkflowTemplateJobsPigJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPigJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPigJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPigJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkRJobToProto(o *beta.WorkflowTemplateJobsSparkRJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkRJob{}
	p.SetMainRFileUri(dcl.ValueOrEmptyString(o.MainRFileUri))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkRJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkRJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkSqlJobToProto(o *beta.WorkflowTemplateJobsSparkSqlJob) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *beta.WorkflowTemplateJobsSparkSqlJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPrestoJobToProto(o *beta.WorkflowTemplateJobsPrestoJob) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocBetaWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetOutputFormat(dcl.ValueOrEmptyString(o.OutputFormat))
	p.SetLoggingConfig(DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocBetaWorkflowTemplateJobsPrestoJobQueryListToProto(o *beta.WorkflowTemplateJobsPrestoJobQueryList) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig object to its proto representation.
func DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *beta.WorkflowTemplateJobsPrestoJobLoggingConfig) *betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsPrestoJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling object to its proto representation.
func DataprocBetaWorkflowTemplateJobsSchedulingToProto(o *beta.WorkflowTemplateJobsScheduling) *betapb.DataprocBetaWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateJobsScheduling{}
	p.SetMaxFailuresPerHour(dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour))
	p.SetMaxFailuresTotal(dcl.ValueOrEmptyInt64(o.MaxFailuresTotal))
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters object to its proto representation.
func DataprocBetaWorkflowTemplateParametersToProto(o *beta.WorkflowTemplateParameters) *betapb.DataprocBetaWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParameters{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetValidation(DataprocBetaWorkflowTemplateParametersValidationToProto(o.Validation))
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation object to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationToProto(o *beta.WorkflowTemplateParametersValidation) *betapb.DataprocBetaWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidation{}
	p.SetRegex(DataprocBetaWorkflowTemplateParametersValidationRegexToProto(o.Regex))
	p.SetValues(DataprocBetaWorkflowTemplateParametersValidationValuesToProto(o.Values))
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex object to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationRegexToProto(o *beta.WorkflowTemplateParametersValidationRegex) *betapb.DataprocBetaWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidationRegex{}
	sRegexes := make([]string, len(o.Regexes))
	for i, r := range o.Regexes {
		sRegexes[i] = r
	}
	p.SetRegexes(sRegexes)
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues object to its proto representation.
func DataprocBetaWorkflowTemplateParametersValidationValuesToProto(o *beta.WorkflowTemplateParametersValidationValues) *betapb.DataprocBetaWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &betapb.DataprocBetaWorkflowTemplateParametersValidationValues{}
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *beta.WorkflowTemplate) *betapb.DataprocBetaWorkflowTemplate {
	p := &betapb.DataprocBetaWorkflowTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyInt64(resource.Version))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetPlacement(DataprocBetaWorkflowTemplatePlacementToProto(resource.Placement))
	p.SetDagTimeout(dcl.ValueOrEmptyString(resource.DagTimeout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sJobs := make([]*betapb.DataprocBetaWorkflowTemplateJobs, len(resource.Jobs))
	for i, r := range resource.Jobs {
		sJobs[i] = DataprocBetaWorkflowTemplateJobsToProto(&r)
	}
	p.SetJobs(sJobs)
	sParameters := make([]*betapb.DataprocBetaWorkflowTemplateParameters, len(resource.Parameters))
	for i, r := range resource.Parameters {
		sParameters[i] = DataprocBetaWorkflowTemplateParametersToProto(&r)
	}
	p.SetParameters(sParameters)

	return p
}

// applyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *beta.Client, request *betapb.ApplyDataprocBetaWorkflowTemplateRequest) (*betapb.DataprocBetaWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// applyDataprocBetaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.ApplyDataprocBetaWorkflowTemplateRequest) (*betapb.DataprocBetaWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.DeleteDataprocBetaWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListDataprocBetaWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocBetaWorkflowTemplate(ctx context.Context, request *betapb.ListDataprocBetaWorkflowTemplateRequest) (*betapb.ListDataprocBetaWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.DataprocBetaWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListDataprocBetaWorkflowTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
