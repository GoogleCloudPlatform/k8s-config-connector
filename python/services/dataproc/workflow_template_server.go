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
	dataprocpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dataproc/dataproc_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc"
)

// WorkflowTemplateServer implements the gRPC interface for WorkflowTemplate.
type WorkflowTemplateServer struct{}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(e dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(n[len("DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(e dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(n[len("DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(e dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum) *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(n[len("DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(e dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum) *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(n[len("DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(e dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(n[len("DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum enum from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(e dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_name[int32(e)]; ok {
		e := dataproc.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(n[len("DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkflowTemplatePlacement converts a WorkflowTemplatePlacement object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacement(p *dataprocpb.DataprocWorkflowTemplatePlacement) *dataproc.WorkflowTemplatePlacement {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacement{
		ManagedCluster:  ProtoToDataprocWorkflowTemplatePlacementManagedCluster(p.GetManagedCluster()),
		ClusterSelector: ProtoToDataprocWorkflowTemplatePlacementClusterSelector(p.GetClusterSelector()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedCluster converts a WorkflowTemplatePlacementManagedCluster object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedCluster(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster) *dataproc.WorkflowTemplatePlacementManagedCluster {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.StringOrNil(p.GetClusterName()),
		Config:      ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfig(p.GetConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfig converts a WorkflowTemplatePlacementManagedClusterConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfig{
		StagingBucket:         dcl.StringOrNil(p.GetStagingBucket()),
		TempBucket:            dcl.StringOrNil(p.GetTempBucket()),
		GceClusterConfig:      ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig(p.GetGceClusterConfig()),
		MasterConfig:          ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig(p.GetMasterConfig()),
		WorkerConfig:          ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig(p.GetWorkerConfig()),
		SecondaryWorkerConfig: ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(p.GetSecondaryWorkerConfig()),
		SoftwareConfig:        ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig(p.GetSoftwareConfig()),
		EncryptionConfig:      ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig(p.GetEncryptionConfig()),
		AutoscalingConfig:     ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig(p.GetAutoscalingConfig()),
		SecurityConfig:        ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig(p.GetSecurityConfig()),
		LifecycleConfig:       ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig(p.GetLifecycleConfig()),
		EndpointConfig:        ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig(p.GetEndpointConfig()),
	}
	for _, r := range p.GetInitializationActions() {
		obj.InitializationActions = append(obj.InitializationActions, *ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig{
		Zone:                    dcl.StringOrNil(p.GetZone()),
		Network:                 dcl.StringOrNil(p.GetNetwork()),
		Subnetwork:              dcl.StringOrNil(p.GetSubnetwork()),
		InternalIPOnly:          dcl.Bool(p.GetInternalIpOnly()),
		PrivateIPv6GoogleAccess: ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(p.GetPrivateIpv6GoogleAccess()),
		ServiceAccount:          dcl.StringOrNil(p.GetServiceAccount()),
		ReservationAffinity:     ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity(p.GetReservationAffinity()),
		NodeGroupAffinity:       ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity(p.GetNodeGroupAffinity()),
		ShieldedInstanceConfig:  ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
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
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity) *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.GetKey()),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity) *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.StringOrNil(p.GetNodeGroup()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.GetEnableSecureBoot()),
		EnableVtpm:                dcl.Bool(p.GetEnableVtpm()),
		EnableIntegrityMonitoring: dcl.Bool(p.GetEnableIntegrityMonitoring()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators) *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators) *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig{
		NumInstances:       dcl.Int64OrNil(p.GetNumInstances()),
		Image:              dcl.StringOrNil(p.GetImage()),
		MachineType:        dcl.StringOrNil(p.GetMachineType()),
		DiskConfig:         ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig(p.GetDiskConfig()),
		IsPreemptible:      dcl.Bool(p.GetIsPreemptible()),
		Preemptibility:     ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(p.GetPreemptibility()),
		ManagedGroupConfig: ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig(p.GetManagedGroupConfig()),
		MinCpuPlatform:     dcl.StringOrNil(p.GetMinCpuPlatform()),
	}
	for _, r := range p.GetInstanceNames() {
		obj.InstanceNames = append(obj.InstanceNames, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig{
		BootDiskType:   dcl.StringOrNil(p.GetBootDiskType()),
		BootDiskSizeGb: dcl.Int64OrNil(p.GetBootDiskSizeGb()),
		NumLocalSsds:   dcl.Int64OrNil(p.GetNumLocalSsds()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig{
		InstanceTemplateName:     dcl.StringOrNil(p.GetInstanceTemplateName()),
		InstanceGroupManagerName: dcl.StringOrNil(p.GetInstanceGroupManagerName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators{
		AcceleratorType:  dcl.StringOrNil(p.GetAcceleratorType()),
		AcceleratorCount: dcl.Int64OrNil(p.GetAcceleratorCount()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig{
		ImageVersion: dcl.StringOrNil(p.GetImageVersion()),
	}
	for _, r := range p.GetOptionalComponents() {
		obj.OptionalComponents = append(obj.OptionalComponents, *ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(r))
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigInitializationActions converts a WorkflowTemplatePlacementManagedClusterConfigInitializationActions object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions) *dataproc.WorkflowTemplatePlacementManagedClusterConfigInitializationActions {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigInitializationActions{
		ExecutableFile:   dcl.StringOrNil(p.GetExecutableFile()),
		ExecutionTimeout: dcl.StringOrNil(p.GetExecutionTimeout()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig converts a WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.StringOrNil(p.GetGcePdKmsKeyName()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig converts a WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig{
		Policy: dcl.StringOrNil(p.GetPolicy()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecurityConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig{
		KerberosConfig: ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig(p.GetKerberosConfig()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig{
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
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig{
		IdleDeleteTtl:  dcl.StringOrNil(p.GetIdleDeleteTtl()),
		AutoDeleteTime: dcl.StringOrNil(p.GetAutoDeleteTime()),
		AutoDeleteTtl:  dcl.StringOrNil(p.GetAutoDeleteTtl()),
		IdleStartTime:  dcl.StringOrNil(p.GetIdleStartTime()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementManagedClusterConfigEndpointConfig converts a WorkflowTemplatePlacementManagedClusterConfigEndpointConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig(p *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig) *dataproc.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(p.GetEnableHttpPortAccess()),
	}
	return obj
}

// ProtoToWorkflowTemplatePlacementClusterSelector converts a WorkflowTemplatePlacementClusterSelector object from its proto representation.
func ProtoToDataprocWorkflowTemplatePlacementClusterSelector(p *dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector) *dataproc.WorkflowTemplatePlacementClusterSelector {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplatePlacementClusterSelector{
		Zone: dcl.StringOrNil(p.GetZone()),
	}
	return obj
}

// ProtoToWorkflowTemplateJobs converts a WorkflowTemplateJobs object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobs(p *dataprocpb.DataprocWorkflowTemplateJobs) *dataproc.WorkflowTemplateJobs {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobs{
		StepId:      dcl.StringOrNil(p.GetStepId()),
		HadoopJob:   ProtoToDataprocWorkflowTemplateJobsHadoopJob(p.GetHadoopJob()),
		SparkJob:    ProtoToDataprocWorkflowTemplateJobsSparkJob(p.GetSparkJob()),
		PysparkJob:  ProtoToDataprocWorkflowTemplateJobsPysparkJob(p.GetPysparkJob()),
		HiveJob:     ProtoToDataprocWorkflowTemplateJobsHiveJob(p.GetHiveJob()),
		PigJob:      ProtoToDataprocWorkflowTemplateJobsPigJob(p.GetPigJob()),
		SparkRJob:   ProtoToDataprocWorkflowTemplateJobsSparkRJob(p.GetSparkRJob()),
		SparkSqlJob: ProtoToDataprocWorkflowTemplateJobsSparkSqlJob(p.GetSparkSqlJob()),
		PrestoJob:   ProtoToDataprocWorkflowTemplateJobsPrestoJob(p.GetPrestoJob()),
		Scheduling:  ProtoToDataprocWorkflowTemplateJobsScheduling(p.GetScheduling()),
	}
	for _, r := range p.GetPrerequisiteStepIds() {
		obj.PrerequisiteStepIds = append(obj.PrerequisiteStepIds, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHadoopJob converts a WorkflowTemplateJobsHadoopJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHadoopJob(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJob) *dataproc.WorkflowTemplateJobsHadoopJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
		LoggingConfig:  ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig) *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkJob converts a WorkflowTemplateJobsSparkJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJob) *dataproc.WorkflowTemplateJobsSparkJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJob{
		MainJarFileUri: dcl.StringOrNil(p.GetMainJarFileUri()),
		MainClass:      dcl.StringOrNil(p.GetMainClass()),
		LoggingConfig:  ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocWorkflowTemplateJobsSparkJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPysparkJob converts a WorkflowTemplateJobsPysparkJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPysparkJob(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJob) *dataproc.WorkflowTemplateJobsPysparkJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.StringOrNil(p.GetMainPythonFileUri()),
		LoggingConfig:     ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig) *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJob converts a WorkflowTemplateJobsHiveJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHiveJob(p *dataprocpb.DataprocWorkflowTemplateJobsHiveJob) *dataproc.WorkflowTemplateJobsHiveJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHiveJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsHiveJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsHiveJobQueryList converts a WorkflowTemplateJobsHiveJobQueryList object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsHiveJobQueryList(p *dataprocpb.DataprocWorkflowTemplateJobsHiveJobQueryList) *dataproc.WorkflowTemplateJobsHiveJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsHiveJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJob converts a WorkflowTemplateJobsPigJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJob(p *dataprocpb.DataprocWorkflowTemplateJobsPigJob) *dataproc.WorkflowTemplateJobsPigJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsPigJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		LoggingConfig:     ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobQueryList converts a WorkflowTemplateJobsPigJobQueryList object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobQueryList(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobQueryList) *dataproc.WorkflowTemplateJobsPigJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPigJobLoggingConfig converts a WorkflowTemplateJobsPigJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPigJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig) *dataproc.WorkflowTemplateJobsPigJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPigJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkRJob converts a WorkflowTemplateJobsSparkRJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkRJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJob) *dataproc.WorkflowTemplateJobsSparkRJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.StringOrNil(p.GetMainRFileUri()),
		LoggingConfig: ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(p.GetLoggingConfig()),
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
func ProtoToDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJob converts a WorkflowTemplateJobsSparkSqlJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJob(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob) *dataproc.WorkflowTemplateJobsSparkSqlJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJob{
		QueryFileUri:  dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:     ProtoToDataprocWorkflowTemplateJobsSparkSqlJobQueryList(p.GetQueryList()),
		LoggingConfig: ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetJarFileUris() {
		obj.JarFileUris = append(obj.JarFileUris, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobQueryList converts a WorkflowTemplateJobsSparkSqlJobQueryList object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobQueryList(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobQueryList) *dataproc.WorkflowTemplateJobsSparkSqlJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsSparkSqlJobLoggingConfig converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig) *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJob converts a WorkflowTemplateJobsPrestoJob object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJob(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJob) *dataproc.WorkflowTemplateJobsPrestoJob {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJob{
		QueryFileUri:      dcl.StringOrNil(p.GetQueryFileUri()),
		QueryList:         ProtoToDataprocWorkflowTemplateJobsPrestoJobQueryList(p.GetQueryList()),
		ContinueOnFailure: dcl.Bool(p.GetContinueOnFailure()),
		OutputFormat:      dcl.StringOrNil(p.GetOutputFormat()),
		LoggingConfig:     ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(p.GetLoggingConfig()),
	}
	for _, r := range p.GetClientTags() {
		obj.ClientTags = append(obj.ClientTags, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobQueryList converts a WorkflowTemplateJobsPrestoJobQueryList object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobQueryList(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobQueryList) *dataproc.WorkflowTemplateJobsPrestoJobQueryList {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJobQueryList{}
	for _, r := range p.GetQueries() {
		obj.Queries = append(obj.Queries, r)
	}
	return obj
}

// ProtoToWorkflowTemplateJobsPrestoJobLoggingConfig converts a WorkflowTemplateJobsPrestoJobLoggingConfig object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(p *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig) *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig{}
	return obj
}

// ProtoToWorkflowTemplateJobsScheduling converts a WorkflowTemplateJobsScheduling object from its proto representation.
func ProtoToDataprocWorkflowTemplateJobsScheduling(p *dataprocpb.DataprocWorkflowTemplateJobsScheduling) *dataproc.WorkflowTemplateJobsScheduling {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64OrNil(p.GetMaxFailuresPerHour()),
		MaxFailuresTotal:   dcl.Int64OrNil(p.GetMaxFailuresTotal()),
	}
	return obj
}

// ProtoToWorkflowTemplateParameters converts a WorkflowTemplateParameters object from its proto representation.
func ProtoToDataprocWorkflowTemplateParameters(p *dataprocpb.DataprocWorkflowTemplateParameters) *dataproc.WorkflowTemplateParameters {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParameters{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Validation:  ProtoToDataprocWorkflowTemplateParametersValidation(p.GetValidation()),
	}
	for _, r := range p.GetFields() {
		obj.Fields = append(obj.Fields, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidation converts a WorkflowTemplateParametersValidation object from its proto representation.
func ProtoToDataprocWorkflowTemplateParametersValidation(p *dataprocpb.DataprocWorkflowTemplateParametersValidation) *dataproc.WorkflowTemplateParametersValidation {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParametersValidation{
		Regex:  ProtoToDataprocWorkflowTemplateParametersValidationRegex(p.GetRegex()),
		Values: ProtoToDataprocWorkflowTemplateParametersValidationValues(p.GetValues()),
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationRegex converts a WorkflowTemplateParametersValidationRegex object from its proto representation.
func ProtoToDataprocWorkflowTemplateParametersValidationRegex(p *dataprocpb.DataprocWorkflowTemplateParametersValidationRegex) *dataproc.WorkflowTemplateParametersValidationRegex {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParametersValidationRegex{}
	for _, r := range p.GetRegexes() {
		obj.Regexes = append(obj.Regexes, r)
	}
	return obj
}

// ProtoToWorkflowTemplateParametersValidationValues converts a WorkflowTemplateParametersValidationValues object from its proto representation.
func ProtoToDataprocWorkflowTemplateParametersValidationValues(p *dataprocpb.DataprocWorkflowTemplateParametersValidationValues) *dataproc.WorkflowTemplateParametersValidationValues {
	if p == nil {
		return nil
	}
	obj := &dataproc.WorkflowTemplateParametersValidationValues{}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToWorkflowTemplate converts a WorkflowTemplate resource from its proto representation.
func ProtoToWorkflowTemplate(p *dataprocpb.DataprocWorkflowTemplate) *dataproc.WorkflowTemplate {
	obj := &dataproc.WorkflowTemplate{
		Name:       dcl.StringOrNil(p.GetName()),
		Version:    dcl.Int64OrNil(p.GetVersion()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Placement:  ProtoToDataprocWorkflowTemplatePlacement(p.GetPlacement()),
		DagTimeout: dcl.StringOrNil(p.GetDagTimeout()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetJobs() {
		obj.Jobs = append(obj.Jobs, *ProtoToDataprocWorkflowTemplateJobs(r))
	}
	for _, r := range p.GetParameters() {
		obj.Parameters = append(obj.Parameters, *ProtoToDataprocWorkflowTemplateParameters(r))
	}
	return obj
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum enum to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(e *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum) dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum_value["WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(e *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum) dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum_value["WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum enum to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto(e *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum) dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto(e *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum) dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum enum to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(e *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum) dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum_value["WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnum(0)
}

// WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumToProto converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum enum to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnumToProto(e *dataproc.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum) dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum {
	if e == nil {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(0)
	}
	if v, ok := dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_value["WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum"+string(*e)]; ok {
		return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(v)
	}
	return dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(0)
}

// WorkflowTemplatePlacementToProto converts a WorkflowTemplatePlacement object to its proto representation.
func DataprocWorkflowTemplatePlacementToProto(o *dataproc.WorkflowTemplatePlacement) *dataprocpb.DataprocWorkflowTemplatePlacement {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacement{}
	p.SetManagedCluster(DataprocWorkflowTemplatePlacementManagedClusterToProto(o.ManagedCluster))
	p.SetClusterSelector(DataprocWorkflowTemplatePlacementClusterSelectorToProto(o.ClusterSelector))
	return p
}

// WorkflowTemplatePlacementManagedClusterToProto converts a WorkflowTemplatePlacementManagedCluster object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterToProto(o *dataproc.WorkflowTemplatePlacementManagedCluster) *dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedCluster{}
	p.SetClusterName(dcl.ValueOrEmptyString(o.ClusterName))
	p.SetConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigToProto(o.Config))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfig{}
	p.SetStagingBucket(dcl.ValueOrEmptyString(o.StagingBucket))
	p.SetTempBucket(dcl.ValueOrEmptyString(o.TempBucket))
	p.SetGceClusterConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto(o.GceClusterConfig))
	p.SetMasterConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto(o.MasterConfig))
	p.SetWorkerConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto(o.WorkerConfig))
	p.SetSecondaryWorkerConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto(o.SecondaryWorkerConfig))
	p.SetSoftwareConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto(o.SoftwareConfig))
	p.SetEncryptionConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto(o.EncryptionConfig))
	p.SetAutoscalingConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto(o.AutoscalingConfig))
	p.SetSecurityConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto(o.SecurityConfig))
	p.SetLifecycleConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigToProto(o.LifecycleConfig))
	p.SetEndpointConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto(o.EndpointConfig))
	sInitializationActions := make([]*dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions, len(o.InitializationActions))
	for i, r := range o.InitializationActions {
		sInitializationActions[i] = DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto(&r)
	}
	p.SetInitializationActions(sInitializationActions)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	p.SetNetwork(dcl.ValueOrEmptyString(o.Network))
	p.SetSubnetwork(dcl.ValueOrEmptyString(o.Subnetwork))
	p.SetInternalIpOnly(dcl.ValueOrEmptyBool(o.InternalIPOnly))
	p.SetPrivateIpv6GoogleAccess(DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumToProto(o.PrivateIPv6GoogleAccess))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetReservationAffinity(DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityToProto(o.ReservationAffinity))
	p.SetNodeGroupAffinity(DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto(o.NodeGroupAffinity))
	p.SetShieldedInstanceConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig))
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
func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity{}
	p.SetConsumeReservationType(DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity{}
	p.SetNodeGroup(dcl.ValueOrEmptyString(o.NodeGroup))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig{}
	p.SetEnableSecureBoot(dcl.ValueOrEmptyBool(o.EnableSecureBoot))
	p.SetEnableVtpm(dcl.ValueOrEmptyBool(o.EnableVtpm))
	p.SetEnableIntegrityMonitoring(dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig{}
	p.SetNumInstances(dcl.ValueOrEmptyInt64(o.NumInstances))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	p.SetDiskConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto(o.DiskConfig))
	p.SetIsPreemptible(dcl.ValueOrEmptyBool(o.IsPreemptible))
	p.SetPreemptibility(DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigPreemptibilityEnumToProto(o.Preemptibility))
	p.SetManagedGroupConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o.ManagedGroupConfig))
	p.SetMinCpuPlatform(dcl.ValueOrEmptyString(o.MinCpuPlatform))
	sInstanceNames := make([]string, len(o.InstanceNames))
	for i, r := range o.InstanceNames {
		sInstanceNames[i] = r
	}
	p.SetInstanceNames(sInstanceNames)
	sAccelerators := make([]*dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators, len(o.Accelerators))
	for i, r := range o.Accelerators {
		sAccelerators[i] = DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto(&r)
	}
	p.SetAccelerators(sAccelerators)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig{}
	p.SetBootDiskType(dcl.ValueOrEmptyString(o.BootDiskType))
	p.SetBootDiskSizeGb(dcl.ValueOrEmptyInt64(o.BootDiskSizeGb))
	p.SetNumLocalSsds(dcl.ValueOrEmptyInt64(o.NumLocalSsds))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig{}
	p.SetInstanceTemplateName(dcl.ValueOrEmptyString(o.InstanceTemplateName))
	p.SetInstanceGroupManagerName(dcl.ValueOrEmptyString(o.InstanceGroupManagerName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators{}
	p.SetAcceleratorType(dcl.ValueOrEmptyString(o.AcceleratorType))
	p.SetAcceleratorCount(dcl.ValueOrEmptyInt64(o.AcceleratorCount))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig{}
	p.SetImageVersion(dcl.ValueOrEmptyString(o.ImageVersion))
	mProperties := make(map[string]string, len(o.Properties))
	for k, r := range o.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)
	sOptionalComponents := make([]dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum, len(o.OptionalComponents))
	for i, r := range o.OptionalComponents {
		sOptionalComponents[i] = dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum(dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOptionalComponentsEnum_value[string(r)])
	}
	p.SetOptionalComponents(sOptionalComponents)
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto converts a WorkflowTemplatePlacementManagedClusterConfigInitializationActions object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigInitializationActions) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions{}
	p.SetExecutableFile(dcl.ValueOrEmptyString(o.ExecutableFile))
	p.SetExecutionTimeout(dcl.ValueOrEmptyString(o.ExecutionTimeout))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig{}
	p.SetGcePdKmsKeyName(dcl.ValueOrEmptyString(o.GcePdKmsKeyName))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig{}
	p.SetPolicy(dcl.ValueOrEmptyString(o.Policy))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecurityConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig{}
	p.SetKerberosConfig(DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto(o.KerberosConfig))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig{}
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
func DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig{}
	p.SetIdleDeleteTtl(dcl.ValueOrEmptyString(o.IdleDeleteTtl))
	p.SetAutoDeleteTime(dcl.ValueOrEmptyString(o.AutoDeleteTime))
	p.SetAutoDeleteTtl(dcl.ValueOrEmptyString(o.AutoDeleteTtl))
	p.SetIdleStartTime(dcl.ValueOrEmptyString(o.IdleStartTime))
	return p
}

// WorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto converts a WorkflowTemplatePlacementManagedClusterConfigEndpointConfig object to its proto representation.
func DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigToProto(o *dataproc.WorkflowTemplatePlacementManagedClusterConfigEndpointConfig) *dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig{}
	p.SetEnableHttpPortAccess(dcl.ValueOrEmptyBool(o.EnableHttpPortAccess))
	mHttpPorts := make(map[string]string, len(o.HttpPorts))
	for k, r := range o.HttpPorts {
		mHttpPorts[k] = r
	}
	p.SetHttpPorts(mHttpPorts)
	return p
}

// WorkflowTemplatePlacementClusterSelectorToProto converts a WorkflowTemplatePlacementClusterSelector object to its proto representation.
func DataprocWorkflowTemplatePlacementClusterSelectorToProto(o *dataproc.WorkflowTemplatePlacementClusterSelector) *dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplatePlacementClusterSelector{}
	p.SetZone(dcl.ValueOrEmptyString(o.Zone))
	mClusterLabels := make(map[string]string, len(o.ClusterLabels))
	for k, r := range o.ClusterLabels {
		mClusterLabels[k] = r
	}
	p.SetClusterLabels(mClusterLabels)
	return p
}

// WorkflowTemplateJobsToProto converts a WorkflowTemplateJobs object to its proto representation.
func DataprocWorkflowTemplateJobsToProto(o *dataproc.WorkflowTemplateJobs) *dataprocpb.DataprocWorkflowTemplateJobs {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobs{}
	p.SetStepId(dcl.ValueOrEmptyString(o.StepId))
	p.SetHadoopJob(DataprocWorkflowTemplateJobsHadoopJobToProto(o.HadoopJob))
	p.SetSparkJob(DataprocWorkflowTemplateJobsSparkJobToProto(o.SparkJob))
	p.SetPysparkJob(DataprocWorkflowTemplateJobsPysparkJobToProto(o.PysparkJob))
	p.SetHiveJob(DataprocWorkflowTemplateJobsHiveJobToProto(o.HiveJob))
	p.SetPigJob(DataprocWorkflowTemplateJobsPigJobToProto(o.PigJob))
	p.SetSparkRJob(DataprocWorkflowTemplateJobsSparkRJobToProto(o.SparkRJob))
	p.SetSparkSqlJob(DataprocWorkflowTemplateJobsSparkSqlJobToProto(o.SparkSqlJob))
	p.SetPrestoJob(DataprocWorkflowTemplateJobsPrestoJobToProto(o.PrestoJob))
	p.SetScheduling(DataprocWorkflowTemplateJobsSchedulingToProto(o.Scheduling))
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
func DataprocWorkflowTemplateJobsHadoopJobToProto(o *dataproc.WorkflowTemplateJobsHadoopJob) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsHadoopJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHadoopJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkJobToProto converts a WorkflowTemplateJobsSparkJob object to its proto representation.
func DataprocWorkflowTemplateJobsSparkJobToProto(o *dataproc.WorkflowTemplateJobsSparkJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJob{}
	p.SetMainJarFileUri(dcl.ValueOrEmptyString(o.MainJarFileUri))
	p.SetMainClass(dcl.ValueOrEmptyString(o.MainClass))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsSparkJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsSparkJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPysparkJobToProto converts a WorkflowTemplateJobsPysparkJob object to its proto representation.
func DataprocWorkflowTemplateJobsPysparkJobToProto(o *dataproc.WorkflowTemplateJobsPysparkJob) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJob{}
	p.SetMainPythonFileUri(dcl.ValueOrEmptyString(o.MainPythonFileUri))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsPysparkJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPysparkJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsHiveJobToProto converts a WorkflowTemplateJobsHiveJob object to its proto representation.
func DataprocWorkflowTemplateJobsHiveJobToProto(o *dataproc.WorkflowTemplateJobsHiveJob) *dataprocpb.DataprocWorkflowTemplateJobsHiveJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHiveJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocWorkflowTemplateJobsHiveJobQueryListToProto(o.QueryList))
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
func DataprocWorkflowTemplateJobsHiveJobQueryListToProto(o *dataproc.WorkflowTemplateJobsHiveJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsHiveJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobToProto converts a WorkflowTemplateJobsPigJob object to its proto representation.
func DataprocWorkflowTemplateJobsPigJobToProto(o *dataproc.WorkflowTemplateJobsPigJob) *dataprocpb.DataprocWorkflowTemplateJobsPigJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocWorkflowTemplateJobsPigJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsPigJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsPigJobQueryListToProto(o *dataproc.WorkflowTemplateJobsPigJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPigJobLoggingConfigToProto converts a WorkflowTemplateJobsPigJobLoggingConfig object to its proto representation.
func DataprocWorkflowTemplateJobsPigJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPigJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPigJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkRJobToProto converts a WorkflowTemplateJobsSparkRJob object to its proto representation.
func DataprocWorkflowTemplateJobsSparkRJobToProto(o *dataproc.WorkflowTemplateJobsSparkRJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJob{}
	p.SetMainRFileUri(dcl.ValueOrEmptyString(o.MainRFileUri))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsSparkRJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkRJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSparkSqlJobToProto converts a WorkflowTemplateJobsSparkSqlJob object to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJob) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocWorkflowTemplateJobsSparkSqlJobQueryListToProto(o.QueryList))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsSparkSqlJobQueryListToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsSparkSqlJobLoggingConfigToProto converts a WorkflowTemplateJobsSparkSqlJobLoggingConfig object to its proto representation.
func DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsPrestoJobToProto converts a WorkflowTemplateJobsPrestoJob object to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobToProto(o *dataproc.WorkflowTemplateJobsPrestoJob) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJob {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJob{}
	p.SetQueryFileUri(dcl.ValueOrEmptyString(o.QueryFileUri))
	p.SetQueryList(DataprocWorkflowTemplateJobsPrestoJobQueryListToProto(o.QueryList))
	p.SetContinueOnFailure(dcl.ValueOrEmptyBool(o.ContinueOnFailure))
	p.SetOutputFormat(dcl.ValueOrEmptyString(o.OutputFormat))
	p.SetLoggingConfig(DataprocWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o.LoggingConfig))
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
func DataprocWorkflowTemplateJobsPrestoJobQueryListToProto(o *dataproc.WorkflowTemplateJobsPrestoJobQueryList) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobQueryList{}
	sQueries := make([]string, len(o.Queries))
	for i, r := range o.Queries {
		sQueries[i] = r
	}
	p.SetQueries(sQueries)
	return p
}

// WorkflowTemplateJobsPrestoJobLoggingConfigToProto converts a WorkflowTemplateJobsPrestoJobLoggingConfig object to its proto representation.
func DataprocWorkflowTemplateJobsPrestoJobLoggingConfigToProto(o *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig) *dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsPrestoJobLoggingConfig{}
	mDriverLogLevels := make(map[string]string, len(o.DriverLogLevels))
	for k, r := range o.DriverLogLevels {
		mDriverLogLevels[k] = r
	}
	p.SetDriverLogLevels(mDriverLogLevels)
	return p
}

// WorkflowTemplateJobsSchedulingToProto converts a WorkflowTemplateJobsScheduling object to its proto representation.
func DataprocWorkflowTemplateJobsSchedulingToProto(o *dataproc.WorkflowTemplateJobsScheduling) *dataprocpb.DataprocWorkflowTemplateJobsScheduling {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateJobsScheduling{}
	p.SetMaxFailuresPerHour(dcl.ValueOrEmptyInt64(o.MaxFailuresPerHour))
	p.SetMaxFailuresTotal(dcl.ValueOrEmptyInt64(o.MaxFailuresTotal))
	return p
}

// WorkflowTemplateParametersToProto converts a WorkflowTemplateParameters object to its proto representation.
func DataprocWorkflowTemplateParametersToProto(o *dataproc.WorkflowTemplateParameters) *dataprocpb.DataprocWorkflowTemplateParameters {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParameters{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetValidation(DataprocWorkflowTemplateParametersValidationToProto(o.Validation))
	sFields := make([]string, len(o.Fields))
	for i, r := range o.Fields {
		sFields[i] = r
	}
	p.SetFields(sFields)
	return p
}

// WorkflowTemplateParametersValidationToProto converts a WorkflowTemplateParametersValidation object to its proto representation.
func DataprocWorkflowTemplateParametersValidationToProto(o *dataproc.WorkflowTemplateParametersValidation) *dataprocpb.DataprocWorkflowTemplateParametersValidation {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidation{}
	p.SetRegex(DataprocWorkflowTemplateParametersValidationRegexToProto(o.Regex))
	p.SetValues(DataprocWorkflowTemplateParametersValidationValuesToProto(o.Values))
	return p
}

// WorkflowTemplateParametersValidationRegexToProto converts a WorkflowTemplateParametersValidationRegex object to its proto representation.
func DataprocWorkflowTemplateParametersValidationRegexToProto(o *dataproc.WorkflowTemplateParametersValidationRegex) *dataprocpb.DataprocWorkflowTemplateParametersValidationRegex {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidationRegex{}
	sRegexes := make([]string, len(o.Regexes))
	for i, r := range o.Regexes {
		sRegexes[i] = r
	}
	p.SetRegexes(sRegexes)
	return p
}

// WorkflowTemplateParametersValidationValuesToProto converts a WorkflowTemplateParametersValidationValues object to its proto representation.
func DataprocWorkflowTemplateParametersValidationValuesToProto(o *dataproc.WorkflowTemplateParametersValidationValues) *dataprocpb.DataprocWorkflowTemplateParametersValidationValues {
	if o == nil {
		return nil
	}
	p := &dataprocpb.DataprocWorkflowTemplateParametersValidationValues{}
	sValues := make([]string, len(o.Values))
	for i, r := range o.Values {
		sValues[i] = r
	}
	p.SetValues(sValues)
	return p
}

// WorkflowTemplateToProto converts a WorkflowTemplate resource to its proto representation.
func WorkflowTemplateToProto(resource *dataproc.WorkflowTemplate) *dataprocpb.DataprocWorkflowTemplate {
	p := &dataprocpb.DataprocWorkflowTemplate{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyInt64(resource.Version))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetPlacement(DataprocWorkflowTemplatePlacementToProto(resource.Placement))
	p.SetDagTimeout(dcl.ValueOrEmptyString(resource.DagTimeout))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sJobs := make([]*dataprocpb.DataprocWorkflowTemplateJobs, len(resource.Jobs))
	for i, r := range resource.Jobs {
		sJobs[i] = DataprocWorkflowTemplateJobsToProto(&r)
	}
	p.SetJobs(sJobs)
	sParameters := make([]*dataprocpb.DataprocWorkflowTemplateParameters, len(resource.Parameters))
	for i, r := range resource.Parameters {
		sParameters[i] = DataprocWorkflowTemplateParametersToProto(&r)
	}
	p.SetParameters(sParameters)

	return p
}

// applyWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) applyWorkflowTemplate(ctx context.Context, c *dataproc.Client, request *dataprocpb.ApplyDataprocWorkflowTemplateRequest) (*dataprocpb.DataprocWorkflowTemplate, error) {
	p := ProtoToWorkflowTemplate(request.GetResource())
	res, err := c.ApplyWorkflowTemplate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkflowTemplateToProto(res)
	return r, nil
}

// applyDataprocWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Apply() method.
func (s *WorkflowTemplateServer) ApplyDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.ApplyDataprocWorkflowTemplateRequest) (*dataprocpb.DataprocWorkflowTemplate, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkflowTemplate(ctx, cl, request)
}

// DeleteWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplate Delete() method.
func (s *WorkflowTemplateServer) DeleteDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.DeleteDataprocWorkflowTemplateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkflowTemplate(ctx, ProtoToWorkflowTemplate(request.GetResource()))

}

// ListDataprocWorkflowTemplate handles the gRPC request by passing it to the underlying WorkflowTemplateList() method.
func (s *WorkflowTemplateServer) ListDataprocWorkflowTemplate(ctx context.Context, request *dataprocpb.ListDataprocWorkflowTemplateRequest) (*dataprocpb.ListDataprocWorkflowTemplateResponse, error) {
	cl, err := createConfigWorkflowTemplate(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkflowTemplate(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*dataprocpb.DataprocWorkflowTemplate
	for _, r := range resources.Items {
		rp := WorkflowTemplateToProto(r)
		protos = append(protos, rp)
	}
	p := &dataprocpb.ListDataprocWorkflowTemplateResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkflowTemplate(ctx context.Context, service_account_file string) (*dataproc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dataproc.NewClient(conf), nil
}
