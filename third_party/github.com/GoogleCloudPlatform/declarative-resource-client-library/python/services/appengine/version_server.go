// Copyright 2021 Google LLC. All Rights Reserved.
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
	appenginepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/appengine/appengine_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/appengine"
)

// Server implements the gRPC interface for Version.
type VersionServer struct{}

// ProtoToVersionInboundServicesEnum converts a VersionInboundServicesEnum enum from its proto representation.
func ProtoToAppengineVersionInboundServicesEnum(e appenginepb.AppengineVersionInboundServicesEnum) *appengine.VersionInboundServicesEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionInboundServicesEnum_name[int32(e)]; ok {
		e := appengine.VersionInboundServicesEnum(n[len("AppengineVersionInboundServicesEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionServingStatusEnum converts a VersionServingStatusEnum enum from its proto representation.
func ProtoToAppengineVersionServingStatusEnum(e appenginepb.AppengineVersionServingStatusEnum) *appengine.VersionServingStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionServingStatusEnum_name[int32(e)]; ok {
		e := appengine.VersionServingStatusEnum(n[len("AppengineVersionServingStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionHandlersSecurityLevelEnum converts a VersionHandlersSecurityLevelEnum enum from its proto representation.
func ProtoToAppengineVersionHandlersSecurityLevelEnum(e appenginepb.AppengineVersionHandlersSecurityLevelEnum) *appengine.VersionHandlersSecurityLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionHandlersSecurityLevelEnum_name[int32(e)]; ok {
		e := appengine.VersionHandlersSecurityLevelEnum(n[len("AppengineVersionHandlersSecurityLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionHandlersLoginEnum converts a VersionHandlersLoginEnum enum from its proto representation.
func ProtoToAppengineVersionHandlersLoginEnum(e appenginepb.AppengineVersionHandlersLoginEnum) *appengine.VersionHandlersLoginEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionHandlersLoginEnum_name[int32(e)]; ok {
		e := appengine.VersionHandlersLoginEnum(n[len("AppengineVersionHandlersLoginEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionHandlersAuthFailActionEnum converts a VersionHandlersAuthFailActionEnum enum from its proto representation.
func ProtoToAppengineVersionHandlersAuthFailActionEnum(e appenginepb.AppengineVersionHandlersAuthFailActionEnum) *appengine.VersionHandlersAuthFailActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionHandlersAuthFailActionEnum_name[int32(e)]; ok {
		e := appengine.VersionHandlersAuthFailActionEnum(n[len("AppengineVersionHandlersAuthFailActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionHandlersRedirectHttpResponseCodeEnum converts a VersionHandlersRedirectHttpResponseCodeEnum enum from its proto representation.
func ProtoToAppengineVersionHandlersRedirectHttpResponseCodeEnum(e appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum) *appengine.VersionHandlersRedirectHttpResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum_name[int32(e)]; ok {
		e := appengine.VersionHandlersRedirectHttpResponseCodeEnum(n[len("AppengineVersionHandlersRedirectHttpResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionErrorHandlersErrorCodeEnum converts a VersionErrorHandlersErrorCodeEnum enum from its proto representation.
func ProtoToAppengineVersionErrorHandlersErrorCodeEnum(e appenginepb.AppengineVersionErrorHandlersErrorCodeEnum) *appengine.VersionErrorHandlersErrorCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionErrorHandlersErrorCodeEnum_name[int32(e)]; ok {
		e := appengine.VersionErrorHandlersErrorCodeEnum(n[len("AppengineVersionErrorHandlersErrorCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionApiConfigAuthFailActionEnum converts a VersionApiConfigAuthFailActionEnum enum from its proto representation.
func ProtoToAppengineVersionApiConfigAuthFailActionEnum(e appenginepb.AppengineVersionApiConfigAuthFailActionEnum) *appengine.VersionApiConfigAuthFailActionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionApiConfigAuthFailActionEnum_name[int32(e)]; ok {
		e := appengine.VersionApiConfigAuthFailActionEnum(n[len("AppengineVersionApiConfigAuthFailActionEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionApiConfigLoginEnum converts a VersionApiConfigLoginEnum enum from its proto representation.
func ProtoToAppengineVersionApiConfigLoginEnum(e appenginepb.AppengineVersionApiConfigLoginEnum) *appengine.VersionApiConfigLoginEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionApiConfigLoginEnum_name[int32(e)]; ok {
		e := appengine.VersionApiConfigLoginEnum(n[len("AppengineVersionApiConfigLoginEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionApiConfigSecurityLevelEnum converts a VersionApiConfigSecurityLevelEnum enum from its proto representation.
func ProtoToAppengineVersionApiConfigSecurityLevelEnum(e appenginepb.AppengineVersionApiConfigSecurityLevelEnum) *appengine.VersionApiConfigSecurityLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineVersionApiConfigSecurityLevelEnum_name[int32(e)]; ok {
		e := appengine.VersionApiConfigSecurityLevelEnum(n[len("AppengineVersionApiConfigSecurityLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToVersionAutomaticScaling converts a VersionAutomaticScaling resource from its proto representation.
func ProtoToAppengineVersionAutomaticScaling(p *appenginepb.AppengineVersionAutomaticScaling) *appengine.VersionAutomaticScaling {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScaling{
		CoolDownPeriod:            dcl.StringOrNil(p.CoolDownPeriod),
		CpuUtilization:            ProtoToAppengineVersionAutomaticScalingCpuUtilization(p.GetCpuUtilization()),
		MaxConcurrentRequests:     dcl.Int64OrNil(p.MaxConcurrentRequests),
		MaxIdleInstances:          dcl.Int64OrNil(p.MaxIdleInstances),
		MaxTotalInstances:         dcl.Int64OrNil(p.MaxTotalInstances),
		MaxPendingLatency:         dcl.StringOrNil(p.MaxPendingLatency),
		MinIdleInstances:          dcl.Int64OrNil(p.MinIdleInstances),
		MinTotalInstances:         dcl.Int64OrNil(p.MinTotalInstances),
		MinPendingLatency:         dcl.StringOrNil(p.MinPendingLatency),
		RequestUtilization:        ProtoToAppengineVersionAutomaticScalingRequestUtilization(p.GetRequestUtilization()),
		DiskUtilization:           ProtoToAppengineVersionAutomaticScalingDiskUtilization(p.GetDiskUtilization()),
		NetworkUtilization:        ProtoToAppengineVersionAutomaticScalingNetworkUtilization(p.GetNetworkUtilization()),
		StandardSchedulerSettings: ProtoToAppengineVersionAutomaticScalingStandardSchedulerSettings(p.GetStandardSchedulerSettings()),
	}
	return obj
}

// ProtoToVersionAutomaticScalingCpuUtilization converts a VersionAutomaticScalingCpuUtilization resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingCpuUtilization(p *appenginepb.AppengineVersionAutomaticScalingCpuUtilization) *appengine.VersionAutomaticScalingCpuUtilization {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingCpuUtilization{
		AggregationWindowLength: dcl.StringOrNil(p.AggregationWindowLength),
		TargetUtilization:       dcl.Float64OrNil(p.TargetUtilization),
	}
	return obj
}

// ProtoToVersionAutomaticScalingRequestUtilization converts a VersionAutomaticScalingRequestUtilization resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingRequestUtilization(p *appenginepb.AppengineVersionAutomaticScalingRequestUtilization) *appengine.VersionAutomaticScalingRequestUtilization {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingRequestUtilization{
		TargetRequestCountPerSecond: dcl.Int64OrNil(p.TargetRequestCountPerSecond),
		TargetConcurrentRequests:    dcl.Int64OrNil(p.TargetConcurrentRequests),
	}
	return obj
}

// ProtoToVersionAutomaticScalingDiskUtilization converts a VersionAutomaticScalingDiskUtilization resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingDiskUtilization(p *appenginepb.AppengineVersionAutomaticScalingDiskUtilization) *appengine.VersionAutomaticScalingDiskUtilization {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingDiskUtilization{
		TargetWriteBytesPerSecond: dcl.Int64OrNil(p.TargetWriteBytesPerSecond),
		TargetWriteOpsPerSecond:   dcl.Int64OrNil(p.TargetWriteOpsPerSecond),
		TargetReadBytesPerSecond:  dcl.Int64OrNil(p.TargetReadBytesPerSecond),
		TargetReadOpsPerSecond:    dcl.Int64OrNil(p.TargetReadOpsPerSecond),
	}
	return obj
}

// ProtoToVersionAutomaticScalingNetworkUtilization converts a VersionAutomaticScalingNetworkUtilization resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingNetworkUtilization(p *appenginepb.AppengineVersionAutomaticScalingNetworkUtilization) *appengine.VersionAutomaticScalingNetworkUtilization {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingNetworkUtilization{
		TargetSentBytesPerSecond:       dcl.Int64OrNil(p.TargetSentBytesPerSecond),
		TargetSentPacketsPerSecond:     dcl.Int64OrNil(p.TargetSentPacketsPerSecond),
		TargetReceivedBytesPerSecond:   dcl.Int64OrNil(p.TargetReceivedBytesPerSecond),
		TargetReceivedPacketsPerSecond: dcl.Int64OrNil(p.TargetReceivedPacketsPerSecond),
	}
	return obj
}

// ProtoToVersionAutomaticScalingStandardSchedulerSettings converts a VersionAutomaticScalingStandardSchedulerSettings resource from its proto representation.
func ProtoToAppengineVersionAutomaticScalingStandardSchedulerSettings(p *appenginepb.AppengineVersionAutomaticScalingStandardSchedulerSettings) *appengine.VersionAutomaticScalingStandardSchedulerSettings {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionAutomaticScalingStandardSchedulerSettings{
		TargetCpuUtilization:        dcl.Float64OrNil(p.TargetCpuUtilization),
		TargetThroughputUtilization: dcl.Float64OrNil(p.TargetThroughputUtilization),
		MinInstances:                dcl.Int64OrNil(p.MinInstances),
		MaxInstances:                dcl.Int64OrNil(p.MaxInstances),
	}
	return obj
}

// ProtoToVersionBasicScaling converts a VersionBasicScaling resource from its proto representation.
func ProtoToAppengineVersionBasicScaling(p *appenginepb.AppengineVersionBasicScaling) *appengine.VersionBasicScaling {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionBasicScaling{
		IdleTimeout:  dcl.StringOrNil(p.IdleTimeout),
		MaxInstances: dcl.Int64OrNil(p.MaxInstances),
	}
	return obj
}

// ProtoToVersionManualScaling converts a VersionManualScaling resource from its proto representation.
func ProtoToAppengineVersionManualScaling(p *appenginepb.AppengineVersionManualScaling) *appengine.VersionManualScaling {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionManualScaling{
		Instances: dcl.Int64OrNil(p.Instances),
	}
	return obj
}

// ProtoToVersionNetwork converts a VersionNetwork resource from its proto representation.
func ProtoToAppengineVersionNetwork(p *appenginepb.AppengineVersionNetwork) *appengine.VersionNetwork {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionNetwork{
		InstanceTag:     dcl.StringOrNil(p.InstanceTag),
		Name:            dcl.StringOrNil(p.Name),
		SubnetworkName:  dcl.StringOrNil(p.SubnetworkName),
		SessionAffinity: dcl.Bool(p.SessionAffinity),
	}
	for _, r := range p.GetForwardedPorts() {
		obj.ForwardedPorts = append(obj.ForwardedPorts, r)
	}
	return obj
}

// ProtoToVersionResources converts a VersionResources resource from its proto representation.
func ProtoToAppengineVersionResources(p *appenginepb.AppengineVersionResources) *appengine.VersionResources {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionResources{
		Cpu:      dcl.Float64OrNil(p.Cpu),
		DiskGb:   dcl.Float64OrNil(p.DiskGb),
		MemoryGb: dcl.Float64OrNil(p.MemoryGb),
	}
	for _, r := range p.GetVolumes() {
		obj.Volumes = append(obj.Volumes, *ProtoToAppengineVersionResourcesVolumes(r))
	}
	return obj
}

// ProtoToVersionResourcesVolumes converts a VersionResourcesVolumes resource from its proto representation.
func ProtoToAppengineVersionResourcesVolumes(p *appenginepb.AppengineVersionResourcesVolumes) *appengine.VersionResourcesVolumes {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionResourcesVolumes{
		Name:       dcl.StringOrNil(p.Name),
		VolumeType: dcl.StringOrNil(p.VolumeType),
		SizeGb:     dcl.Float64OrNil(p.SizeGb),
	}
	return obj
}

// ProtoToVersionHandlers converts a VersionHandlers resource from its proto representation.
func ProtoToAppengineVersionHandlers(p *appenginepb.AppengineVersionHandlers) *appengine.VersionHandlers {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlers{
		UrlRegex:                 dcl.StringOrNil(p.UrlRegex),
		StaticFiles:              ProtoToAppengineVersionHandlersStaticFiles(p.GetStaticFiles()),
		Script:                   ProtoToAppengineVersionHandlersScript(p.GetScript()),
		ApiEndpoint:              ProtoToAppengineVersionHandlersApiEndpoint(p.GetApiEndpoint()),
		SecurityLevel:            ProtoToAppengineVersionHandlersSecurityLevelEnum(p.GetSecurityLevel()),
		Login:                    ProtoToAppengineVersionHandlersLoginEnum(p.GetLogin()),
		AuthFailAction:           ProtoToAppengineVersionHandlersAuthFailActionEnum(p.GetAuthFailAction()),
		RedirectHttpResponseCode: ProtoToAppengineVersionHandlersRedirectHttpResponseCodeEnum(p.GetRedirectHttpResponseCode()),
	}
	return obj
}

// ProtoToVersionHandlersStaticFiles converts a VersionHandlersStaticFiles resource from its proto representation.
func ProtoToAppengineVersionHandlersStaticFiles(p *appenginepb.AppengineVersionHandlersStaticFiles) *appengine.VersionHandlersStaticFiles {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlersStaticFiles{
		Path:                dcl.StringOrNil(p.Path),
		UploadPathRegex:     dcl.StringOrNil(p.UploadPathRegex),
		MimeType:            dcl.StringOrNil(p.MimeType),
		Expiration:          dcl.StringOrNil(p.Expiration),
		RequireMatchingFile: dcl.Bool(p.RequireMatchingFile),
		ApplicationReadable: dcl.Bool(p.ApplicationReadable),
	}
	return obj
}

// ProtoToVersionHandlersScript converts a VersionHandlersScript resource from its proto representation.
func ProtoToAppengineVersionHandlersScript(p *appenginepb.AppengineVersionHandlersScript) *appengine.VersionHandlersScript {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlersScript{
		ScriptPath: dcl.StringOrNil(p.ScriptPath),
	}
	return obj
}

// ProtoToVersionHandlersApiEndpoint converts a VersionHandlersApiEndpoint resource from its proto representation.
func ProtoToAppengineVersionHandlersApiEndpoint(p *appenginepb.AppengineVersionHandlersApiEndpoint) *appengine.VersionHandlersApiEndpoint {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHandlersApiEndpoint{
		ScriptPath: dcl.StringOrNil(p.ScriptPath),
	}
	return obj
}

// ProtoToVersionErrorHandlers converts a VersionErrorHandlers resource from its proto representation.
func ProtoToAppengineVersionErrorHandlers(p *appenginepb.AppengineVersionErrorHandlers) *appengine.VersionErrorHandlers {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionErrorHandlers{
		ErrorCode:  ProtoToAppengineVersionErrorHandlersErrorCodeEnum(p.GetErrorCode()),
		StaticFile: dcl.StringOrNil(p.StaticFile),
		MimeType:   dcl.StringOrNil(p.MimeType),
	}
	return obj
}

// ProtoToVersionLibraries converts a VersionLibraries resource from its proto representation.
func ProtoToAppengineVersionLibraries(p *appenginepb.AppengineVersionLibraries) *appengine.VersionLibraries {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionLibraries{
		Name:    dcl.StringOrNil(p.Name),
		Version: dcl.StringOrNil(p.Version),
	}
	return obj
}

// ProtoToVersionApiConfig converts a VersionApiConfig resource from its proto representation.
func ProtoToAppengineVersionApiConfig(p *appenginepb.AppengineVersionApiConfig) *appengine.VersionApiConfig {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionApiConfig{
		AuthFailAction: ProtoToAppengineVersionApiConfigAuthFailActionEnum(p.GetAuthFailAction()),
		Login:          ProtoToAppengineVersionApiConfigLoginEnum(p.GetLogin()),
		Script:         dcl.StringOrNil(p.Script),
		SecurityLevel:  ProtoToAppengineVersionApiConfigSecurityLevelEnum(p.GetSecurityLevel()),
		Url:            dcl.StringOrNil(p.Url),
	}
	return obj
}

// ProtoToVersionDeployment converts a VersionDeployment resource from its proto representation.
func ProtoToAppengineVersionDeployment(p *appenginepb.AppengineVersionDeployment) *appengine.VersionDeployment {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeployment{
		Container:         ProtoToAppengineVersionDeploymentContainer(p.GetContainer()),
		Zip:               ProtoToAppengineVersionDeploymentZip(p.GetZip()),
		CloudBuildOptions: ProtoToAppengineVersionDeploymentCloudBuildOptions(p.GetCloudBuildOptions()),
	}
	return obj
}

// ProtoToVersionDeploymentFiles converts a VersionDeploymentFiles resource from its proto representation.
func ProtoToAppengineVersionDeploymentFiles(p *appenginepb.AppengineVersionDeploymentFiles) *appengine.VersionDeploymentFiles {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentFiles{
		SourceUrl: dcl.StringOrNil(p.SourceUrl),
		Sha1Sum:   dcl.StringOrNil(p.Sha1Sum),
		MimeType:  dcl.StringOrNil(p.MimeType),
	}
	return obj
}

// ProtoToVersionDeploymentContainer converts a VersionDeploymentContainer resource from its proto representation.
func ProtoToAppengineVersionDeploymentContainer(p *appenginepb.AppengineVersionDeploymentContainer) *appengine.VersionDeploymentContainer {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentContainer{
		Image: dcl.StringOrNil(p.Image),
	}
	return obj
}

// ProtoToVersionDeploymentZip converts a VersionDeploymentZip resource from its proto representation.
func ProtoToAppengineVersionDeploymentZip(p *appenginepb.AppengineVersionDeploymentZip) *appengine.VersionDeploymentZip {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentZip{
		SourceUrl:  dcl.StringOrNil(p.SourceUrl),
		FilesCount: dcl.Int64OrNil(p.FilesCount),
	}
	return obj
}

// ProtoToVersionDeploymentCloudBuildOptions converts a VersionDeploymentCloudBuildOptions resource from its proto representation.
func ProtoToAppengineVersionDeploymentCloudBuildOptions(p *appenginepb.AppengineVersionDeploymentCloudBuildOptions) *appengine.VersionDeploymentCloudBuildOptions {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionDeploymentCloudBuildOptions{
		AppYamlPath:       dcl.StringOrNil(p.AppYamlPath),
		CloudBuildTimeout: dcl.StringOrNil(p.CloudBuildTimeout),
	}
	return obj
}

// ProtoToVersionHealthCheck converts a VersionHealthCheck resource from its proto representation.
func ProtoToAppengineVersionHealthCheck(p *appenginepb.AppengineVersionHealthCheck) *appengine.VersionHealthCheck {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionHealthCheck{
		DisableHealthCheck: dcl.Bool(p.DisableHealthCheck),
		Host:               dcl.StringOrNil(p.Host),
		HealthyThreshold:   dcl.Int64OrNil(p.HealthyThreshold),
		UnhealthyThreshold: dcl.Int64OrNil(p.UnhealthyThreshold),
		RestartThreshold:   dcl.Int64OrNil(p.RestartThreshold),
		CheckInterval:      dcl.StringOrNil(p.CheckInterval),
		Timeout:            dcl.StringOrNil(p.Timeout),
	}
	return obj
}

// ProtoToVersionReadinessCheck converts a VersionReadinessCheck resource from its proto representation.
func ProtoToAppengineVersionReadinessCheck(p *appenginepb.AppengineVersionReadinessCheck) *appengine.VersionReadinessCheck {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionReadinessCheck{
		Path:             dcl.StringOrNil(p.Path),
		Host:             dcl.StringOrNil(p.Host),
		FailureThreshold: dcl.Int64OrNil(p.FailureThreshold),
		SuccessThreshold: dcl.Int64OrNil(p.SuccessThreshold),
		CheckInterval:    dcl.StringOrNil(p.CheckInterval),
		Timeout:          dcl.StringOrNil(p.Timeout),
		AppStartTimeout:  dcl.StringOrNil(p.AppStartTimeout),
	}
	return obj
}

// ProtoToVersionLivenessCheck converts a VersionLivenessCheck resource from its proto representation.
func ProtoToAppengineVersionLivenessCheck(p *appenginepb.AppengineVersionLivenessCheck) *appengine.VersionLivenessCheck {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionLivenessCheck{
		Path:             dcl.StringOrNil(p.Path),
		Host:             dcl.StringOrNil(p.Host),
		FailureThreshold: dcl.Int64OrNil(p.FailureThreshold),
		SuccessThreshold: dcl.Int64OrNil(p.SuccessThreshold),
		CheckInterval:    dcl.StringOrNil(p.CheckInterval),
		Timeout:          dcl.StringOrNil(p.Timeout),
		InitialDelay:     dcl.StringOrNil(p.InitialDelay),
	}
	return obj
}

// ProtoToVersionEntrypoint converts a VersionEntrypoint resource from its proto representation.
func ProtoToAppengineVersionEntrypoint(p *appenginepb.AppengineVersionEntrypoint) *appengine.VersionEntrypoint {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionEntrypoint{
		Shell: dcl.StringOrNil(p.Shell),
	}
	return obj
}

// ProtoToVersionVPCAccessConnector converts a VersionVPCAccessConnector resource from its proto representation.
func ProtoToAppengineVersionVPCAccessConnector(p *appenginepb.AppengineVersionVPCAccessConnector) *appengine.VersionVPCAccessConnector {
	if p == nil {
		return nil
	}
	obj := &appengine.VersionVPCAccessConnector{
		Name: dcl.StringOrNil(p.Name),
	}
	return obj
}

// ProtoToVersion converts a Version resource from its proto representation.
func ProtoToVersion(p *appenginepb.AppengineVersion) *appengine.Version {
	obj := &appengine.Version{
		ConsumerName:              dcl.StringOrNil(p.ConsumerName),
		Name:                      dcl.StringOrNil(p.Name),
		AutomaticScaling:          ProtoToAppengineVersionAutomaticScaling(p.GetAutomaticScaling()),
		BasicScaling:              ProtoToAppengineVersionBasicScaling(p.GetBasicScaling()),
		ManualScaling:             ProtoToAppengineVersionManualScaling(p.GetManualScaling()),
		InstanceClass:             dcl.StringOrNil(p.InstanceClass),
		Network:                   ProtoToAppengineVersionNetwork(p.GetNetwork()),
		Resources:                 ProtoToAppengineVersionResources(p.GetResources()),
		Runtime:                   dcl.StringOrNil(p.Runtime),
		RuntimeChannel:            dcl.StringOrNil(p.RuntimeChannel),
		Threadsafe:                dcl.Bool(p.Threadsafe),
		Vm:                        dcl.Bool(p.Vm),
		Env:                       dcl.StringOrNil(p.Env),
		ServingStatus:             ProtoToAppengineVersionServingStatusEnum(p.GetServingStatus()),
		CreatedBy:                 dcl.StringOrNil(p.CreatedBy),
		CreateTime:                dcl.StringOrNil(p.GetCreateTime()),
		DiskUsageBytes:            dcl.Int64OrNil(p.DiskUsageBytes),
		RuntimeApiVersion:         dcl.StringOrNil(p.RuntimeApiVersion),
		RuntimeMainExecutablePath: dcl.StringOrNil(p.RuntimeMainExecutablePath),
		ApiConfig:                 ProtoToAppengineVersionApiConfig(p.GetApiConfig()),
		DefaultExpiration:         dcl.StringOrNil(p.DefaultExpiration),
		Deployment:                ProtoToAppengineVersionDeployment(p.GetDeployment()),
		HealthCheck:               ProtoToAppengineVersionHealthCheck(p.GetHealthCheck()),
		ReadinessCheck:            ProtoToAppengineVersionReadinessCheck(p.GetReadinessCheck()),
		LivenessCheck:             ProtoToAppengineVersionLivenessCheck(p.GetLivenessCheck()),
		NobuildFilesRegex:         dcl.StringOrNil(p.NobuildFilesRegex),
		VersionUrl:                dcl.StringOrNil(p.VersionUrl),
		Entrypoint:                ProtoToAppengineVersionEntrypoint(p.GetEntrypoint()),
		VPCAccessConnector:        ProtoToAppengineVersionVPCAccessConnector(p.GetVpcAccessConnector()),
		App:                       dcl.StringOrNil(p.App),
		Service:                   dcl.StringOrNil(p.Service),
	}
	for _, r := range p.GetInboundServices() {
		obj.InboundServices = append(obj.InboundServices, *ProtoToAppengineVersionInboundServicesEnum(r))
	}
	for _, r := range p.GetZones() {
		obj.Zones = append(obj.Zones, r)
	}
	for _, r := range p.GetHandlers() {
		obj.Handlers = append(obj.Handlers, *ProtoToAppengineVersionHandlers(r))
	}
	for _, r := range p.GetErrorHandlers() {
		obj.ErrorHandlers = append(obj.ErrorHandlers, *ProtoToAppengineVersionErrorHandlers(r))
	}
	for _, r := range p.GetLibraries() {
		obj.Libraries = append(obj.Libraries, *ProtoToAppengineVersionLibraries(r))
	}
	return obj
}

// VersionInboundServicesEnumToProto converts a VersionInboundServicesEnum enum to its proto representation.
func AppengineVersionInboundServicesEnumToProto(e *appengine.VersionInboundServicesEnum) appenginepb.AppengineVersionInboundServicesEnum {
	if e == nil {
		return appenginepb.AppengineVersionInboundServicesEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionInboundServicesEnum_value["VersionInboundServicesEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionInboundServicesEnum(v)
	}
	return appenginepb.AppengineVersionInboundServicesEnum(0)
}

// VersionServingStatusEnumToProto converts a VersionServingStatusEnum enum to its proto representation.
func AppengineVersionServingStatusEnumToProto(e *appengine.VersionServingStatusEnum) appenginepb.AppengineVersionServingStatusEnum {
	if e == nil {
		return appenginepb.AppengineVersionServingStatusEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionServingStatusEnum_value["VersionServingStatusEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionServingStatusEnum(v)
	}
	return appenginepb.AppengineVersionServingStatusEnum(0)
}

// VersionHandlersSecurityLevelEnumToProto converts a VersionHandlersSecurityLevelEnum enum to its proto representation.
func AppengineVersionHandlersSecurityLevelEnumToProto(e *appengine.VersionHandlersSecurityLevelEnum) appenginepb.AppengineVersionHandlersSecurityLevelEnum {
	if e == nil {
		return appenginepb.AppengineVersionHandlersSecurityLevelEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionHandlersSecurityLevelEnum_value["VersionHandlersSecurityLevelEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionHandlersSecurityLevelEnum(v)
	}
	return appenginepb.AppengineVersionHandlersSecurityLevelEnum(0)
}

// VersionHandlersLoginEnumToProto converts a VersionHandlersLoginEnum enum to its proto representation.
func AppengineVersionHandlersLoginEnumToProto(e *appengine.VersionHandlersLoginEnum) appenginepb.AppengineVersionHandlersLoginEnum {
	if e == nil {
		return appenginepb.AppengineVersionHandlersLoginEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionHandlersLoginEnum_value["VersionHandlersLoginEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionHandlersLoginEnum(v)
	}
	return appenginepb.AppengineVersionHandlersLoginEnum(0)
}

// VersionHandlersAuthFailActionEnumToProto converts a VersionHandlersAuthFailActionEnum enum to its proto representation.
func AppengineVersionHandlersAuthFailActionEnumToProto(e *appengine.VersionHandlersAuthFailActionEnum) appenginepb.AppengineVersionHandlersAuthFailActionEnum {
	if e == nil {
		return appenginepb.AppengineVersionHandlersAuthFailActionEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionHandlersAuthFailActionEnum_value["VersionHandlersAuthFailActionEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionHandlersAuthFailActionEnum(v)
	}
	return appenginepb.AppengineVersionHandlersAuthFailActionEnum(0)
}

// VersionHandlersRedirectHttpResponseCodeEnumToProto converts a VersionHandlersRedirectHttpResponseCodeEnum enum to its proto representation.
func AppengineVersionHandlersRedirectHttpResponseCodeEnumToProto(e *appengine.VersionHandlersRedirectHttpResponseCodeEnum) appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum {
	if e == nil {
		return appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum_value["VersionHandlersRedirectHttpResponseCodeEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum(v)
	}
	return appenginepb.AppengineVersionHandlersRedirectHttpResponseCodeEnum(0)
}

// VersionErrorHandlersErrorCodeEnumToProto converts a VersionErrorHandlersErrorCodeEnum enum to its proto representation.
func AppengineVersionErrorHandlersErrorCodeEnumToProto(e *appengine.VersionErrorHandlersErrorCodeEnum) appenginepb.AppengineVersionErrorHandlersErrorCodeEnum {
	if e == nil {
		return appenginepb.AppengineVersionErrorHandlersErrorCodeEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionErrorHandlersErrorCodeEnum_value["VersionErrorHandlersErrorCodeEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionErrorHandlersErrorCodeEnum(v)
	}
	return appenginepb.AppengineVersionErrorHandlersErrorCodeEnum(0)
}

// VersionApiConfigAuthFailActionEnumToProto converts a VersionApiConfigAuthFailActionEnum enum to its proto representation.
func AppengineVersionApiConfigAuthFailActionEnumToProto(e *appengine.VersionApiConfigAuthFailActionEnum) appenginepb.AppengineVersionApiConfigAuthFailActionEnum {
	if e == nil {
		return appenginepb.AppengineVersionApiConfigAuthFailActionEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionApiConfigAuthFailActionEnum_value["VersionApiConfigAuthFailActionEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionApiConfigAuthFailActionEnum(v)
	}
	return appenginepb.AppengineVersionApiConfigAuthFailActionEnum(0)
}

// VersionApiConfigLoginEnumToProto converts a VersionApiConfigLoginEnum enum to its proto representation.
func AppengineVersionApiConfigLoginEnumToProto(e *appengine.VersionApiConfigLoginEnum) appenginepb.AppengineVersionApiConfigLoginEnum {
	if e == nil {
		return appenginepb.AppengineVersionApiConfigLoginEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionApiConfigLoginEnum_value["VersionApiConfigLoginEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionApiConfigLoginEnum(v)
	}
	return appenginepb.AppengineVersionApiConfigLoginEnum(0)
}

// VersionApiConfigSecurityLevelEnumToProto converts a VersionApiConfigSecurityLevelEnum enum to its proto representation.
func AppengineVersionApiConfigSecurityLevelEnumToProto(e *appengine.VersionApiConfigSecurityLevelEnum) appenginepb.AppengineVersionApiConfigSecurityLevelEnum {
	if e == nil {
		return appenginepb.AppengineVersionApiConfigSecurityLevelEnum(0)
	}
	if v, ok := appenginepb.AppengineVersionApiConfigSecurityLevelEnum_value["VersionApiConfigSecurityLevelEnum"+string(*e)]; ok {
		return appenginepb.AppengineVersionApiConfigSecurityLevelEnum(v)
	}
	return appenginepb.AppengineVersionApiConfigSecurityLevelEnum(0)
}

// VersionAutomaticScalingToProto converts a VersionAutomaticScaling resource to its proto representation.
func AppengineVersionAutomaticScalingToProto(o *appengine.VersionAutomaticScaling) *appenginepb.AppengineVersionAutomaticScaling {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScaling{
		CoolDownPeriod:            dcl.ValueOrEmptyString(o.CoolDownPeriod),
		CpuUtilization:            AppengineVersionAutomaticScalingCpuUtilizationToProto(o.CpuUtilization),
		MaxConcurrentRequests:     dcl.ValueOrEmptyInt64(o.MaxConcurrentRequests),
		MaxIdleInstances:          dcl.ValueOrEmptyInt64(o.MaxIdleInstances),
		MaxTotalInstances:         dcl.ValueOrEmptyInt64(o.MaxTotalInstances),
		MaxPendingLatency:         dcl.ValueOrEmptyString(o.MaxPendingLatency),
		MinIdleInstances:          dcl.ValueOrEmptyInt64(o.MinIdleInstances),
		MinTotalInstances:         dcl.ValueOrEmptyInt64(o.MinTotalInstances),
		MinPendingLatency:         dcl.ValueOrEmptyString(o.MinPendingLatency),
		RequestUtilization:        AppengineVersionAutomaticScalingRequestUtilizationToProto(o.RequestUtilization),
		DiskUtilization:           AppengineVersionAutomaticScalingDiskUtilizationToProto(o.DiskUtilization),
		NetworkUtilization:        AppengineVersionAutomaticScalingNetworkUtilizationToProto(o.NetworkUtilization),
		StandardSchedulerSettings: AppengineVersionAutomaticScalingStandardSchedulerSettingsToProto(o.StandardSchedulerSettings),
	}
	return p
}

// VersionAutomaticScalingCpuUtilizationToProto converts a VersionAutomaticScalingCpuUtilization resource to its proto representation.
func AppengineVersionAutomaticScalingCpuUtilizationToProto(o *appengine.VersionAutomaticScalingCpuUtilization) *appenginepb.AppengineVersionAutomaticScalingCpuUtilization {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingCpuUtilization{
		AggregationWindowLength: dcl.ValueOrEmptyString(o.AggregationWindowLength),
		TargetUtilization:       dcl.ValueOrEmptyDouble(o.TargetUtilization),
	}
	return p
}

// VersionAutomaticScalingRequestUtilizationToProto converts a VersionAutomaticScalingRequestUtilization resource to its proto representation.
func AppengineVersionAutomaticScalingRequestUtilizationToProto(o *appengine.VersionAutomaticScalingRequestUtilization) *appenginepb.AppengineVersionAutomaticScalingRequestUtilization {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingRequestUtilization{
		TargetRequestCountPerSecond: dcl.ValueOrEmptyInt64(o.TargetRequestCountPerSecond),
		TargetConcurrentRequests:    dcl.ValueOrEmptyInt64(o.TargetConcurrentRequests),
	}
	return p
}

// VersionAutomaticScalingDiskUtilizationToProto converts a VersionAutomaticScalingDiskUtilization resource to its proto representation.
func AppengineVersionAutomaticScalingDiskUtilizationToProto(o *appengine.VersionAutomaticScalingDiskUtilization) *appenginepb.AppengineVersionAutomaticScalingDiskUtilization {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingDiskUtilization{
		TargetWriteBytesPerSecond: dcl.ValueOrEmptyInt64(o.TargetWriteBytesPerSecond),
		TargetWriteOpsPerSecond:   dcl.ValueOrEmptyInt64(o.TargetWriteOpsPerSecond),
		TargetReadBytesPerSecond:  dcl.ValueOrEmptyInt64(o.TargetReadBytesPerSecond),
		TargetReadOpsPerSecond:    dcl.ValueOrEmptyInt64(o.TargetReadOpsPerSecond),
	}
	return p
}

// VersionAutomaticScalingNetworkUtilizationToProto converts a VersionAutomaticScalingNetworkUtilization resource to its proto representation.
func AppengineVersionAutomaticScalingNetworkUtilizationToProto(o *appengine.VersionAutomaticScalingNetworkUtilization) *appenginepb.AppengineVersionAutomaticScalingNetworkUtilization {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingNetworkUtilization{
		TargetSentBytesPerSecond:       dcl.ValueOrEmptyInt64(o.TargetSentBytesPerSecond),
		TargetSentPacketsPerSecond:     dcl.ValueOrEmptyInt64(o.TargetSentPacketsPerSecond),
		TargetReceivedBytesPerSecond:   dcl.ValueOrEmptyInt64(o.TargetReceivedBytesPerSecond),
		TargetReceivedPacketsPerSecond: dcl.ValueOrEmptyInt64(o.TargetReceivedPacketsPerSecond),
	}
	return p
}

// VersionAutomaticScalingStandardSchedulerSettingsToProto converts a VersionAutomaticScalingStandardSchedulerSettings resource to its proto representation.
func AppengineVersionAutomaticScalingStandardSchedulerSettingsToProto(o *appengine.VersionAutomaticScalingStandardSchedulerSettings) *appenginepb.AppengineVersionAutomaticScalingStandardSchedulerSettings {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionAutomaticScalingStandardSchedulerSettings{
		TargetCpuUtilization:        dcl.ValueOrEmptyDouble(o.TargetCpuUtilization),
		TargetThroughputUtilization: dcl.ValueOrEmptyDouble(o.TargetThroughputUtilization),
		MinInstances:                dcl.ValueOrEmptyInt64(o.MinInstances),
		MaxInstances:                dcl.ValueOrEmptyInt64(o.MaxInstances),
	}
	return p
}

// VersionBasicScalingToProto converts a VersionBasicScaling resource to its proto representation.
func AppengineVersionBasicScalingToProto(o *appengine.VersionBasicScaling) *appenginepb.AppengineVersionBasicScaling {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionBasicScaling{
		IdleTimeout:  dcl.ValueOrEmptyString(o.IdleTimeout),
		MaxInstances: dcl.ValueOrEmptyInt64(o.MaxInstances),
	}
	return p
}

// VersionManualScalingToProto converts a VersionManualScaling resource to its proto representation.
func AppengineVersionManualScalingToProto(o *appengine.VersionManualScaling) *appenginepb.AppengineVersionManualScaling {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionManualScaling{
		Instances: dcl.ValueOrEmptyInt64(o.Instances),
	}
	return p
}

// VersionNetworkToProto converts a VersionNetwork resource to its proto representation.
func AppengineVersionNetworkToProto(o *appengine.VersionNetwork) *appenginepb.AppengineVersionNetwork {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionNetwork{
		InstanceTag:     dcl.ValueOrEmptyString(o.InstanceTag),
		Name:            dcl.ValueOrEmptyString(o.Name),
		SubnetworkName:  dcl.ValueOrEmptyString(o.SubnetworkName),
		SessionAffinity: dcl.ValueOrEmptyBool(o.SessionAffinity),
	}
	for _, r := range o.ForwardedPorts {
		p.ForwardedPorts = append(p.ForwardedPorts, r)
	}
	return p
}

// VersionResourcesToProto converts a VersionResources resource to its proto representation.
func AppengineVersionResourcesToProto(o *appengine.VersionResources) *appenginepb.AppengineVersionResources {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionResources{
		Cpu:      dcl.ValueOrEmptyDouble(o.Cpu),
		DiskGb:   dcl.ValueOrEmptyDouble(o.DiskGb),
		MemoryGb: dcl.ValueOrEmptyDouble(o.MemoryGb),
	}
	for _, r := range o.Volumes {
		p.Volumes = append(p.Volumes, AppengineVersionResourcesVolumesToProto(&r))
	}
	return p
}

// VersionResourcesVolumesToProto converts a VersionResourcesVolumes resource to its proto representation.
func AppengineVersionResourcesVolumesToProto(o *appengine.VersionResourcesVolumes) *appenginepb.AppengineVersionResourcesVolumes {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionResourcesVolumes{
		Name:       dcl.ValueOrEmptyString(o.Name),
		VolumeType: dcl.ValueOrEmptyString(o.VolumeType),
		SizeGb:     dcl.ValueOrEmptyDouble(o.SizeGb),
	}
	return p
}

// VersionHandlersToProto converts a VersionHandlers resource to its proto representation.
func AppengineVersionHandlersToProto(o *appengine.VersionHandlers) *appenginepb.AppengineVersionHandlers {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlers{
		UrlRegex:                 dcl.ValueOrEmptyString(o.UrlRegex),
		StaticFiles:              AppengineVersionHandlersStaticFilesToProto(o.StaticFiles),
		Script:                   AppengineVersionHandlersScriptToProto(o.Script),
		ApiEndpoint:              AppengineVersionHandlersApiEndpointToProto(o.ApiEndpoint),
		SecurityLevel:            AppengineVersionHandlersSecurityLevelEnumToProto(o.SecurityLevel),
		Login:                    AppengineVersionHandlersLoginEnumToProto(o.Login),
		AuthFailAction:           AppengineVersionHandlersAuthFailActionEnumToProto(o.AuthFailAction),
		RedirectHttpResponseCode: AppengineVersionHandlersRedirectHttpResponseCodeEnumToProto(o.RedirectHttpResponseCode),
	}
	return p
}

// VersionHandlersStaticFilesToProto converts a VersionHandlersStaticFiles resource to its proto representation.
func AppengineVersionHandlersStaticFilesToProto(o *appengine.VersionHandlersStaticFiles) *appenginepb.AppengineVersionHandlersStaticFiles {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlersStaticFiles{
		Path:                dcl.ValueOrEmptyString(o.Path),
		UploadPathRegex:     dcl.ValueOrEmptyString(o.UploadPathRegex),
		MimeType:            dcl.ValueOrEmptyString(o.MimeType),
		Expiration:          dcl.ValueOrEmptyString(o.Expiration),
		RequireMatchingFile: dcl.ValueOrEmptyBool(o.RequireMatchingFile),
		ApplicationReadable: dcl.ValueOrEmptyBool(o.ApplicationReadable),
	}
	p.HttpHeaders = make(map[string]string)
	for k, r := range o.HttpHeaders {
		p.HttpHeaders[k] = r
	}
	return p
}

// VersionHandlersScriptToProto converts a VersionHandlersScript resource to its proto representation.
func AppengineVersionHandlersScriptToProto(o *appengine.VersionHandlersScript) *appenginepb.AppengineVersionHandlersScript {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlersScript{
		ScriptPath: dcl.ValueOrEmptyString(o.ScriptPath),
	}
	return p
}

// VersionHandlersApiEndpointToProto converts a VersionHandlersApiEndpoint resource to its proto representation.
func AppengineVersionHandlersApiEndpointToProto(o *appengine.VersionHandlersApiEndpoint) *appenginepb.AppengineVersionHandlersApiEndpoint {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHandlersApiEndpoint{
		ScriptPath: dcl.ValueOrEmptyString(o.ScriptPath),
	}
	return p
}

// VersionErrorHandlersToProto converts a VersionErrorHandlers resource to its proto representation.
func AppengineVersionErrorHandlersToProto(o *appengine.VersionErrorHandlers) *appenginepb.AppengineVersionErrorHandlers {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionErrorHandlers{
		ErrorCode:  AppengineVersionErrorHandlersErrorCodeEnumToProto(o.ErrorCode),
		StaticFile: dcl.ValueOrEmptyString(o.StaticFile),
		MimeType:   dcl.ValueOrEmptyString(o.MimeType),
	}
	return p
}

// VersionLibrariesToProto converts a VersionLibraries resource to its proto representation.
func AppengineVersionLibrariesToProto(o *appengine.VersionLibraries) *appenginepb.AppengineVersionLibraries {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionLibraries{
		Name:    dcl.ValueOrEmptyString(o.Name),
		Version: dcl.ValueOrEmptyString(o.Version),
	}
	return p
}

// VersionApiConfigToProto converts a VersionApiConfig resource to its proto representation.
func AppengineVersionApiConfigToProto(o *appengine.VersionApiConfig) *appenginepb.AppengineVersionApiConfig {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionApiConfig{
		AuthFailAction: AppengineVersionApiConfigAuthFailActionEnumToProto(o.AuthFailAction),
		Login:          AppengineVersionApiConfigLoginEnumToProto(o.Login),
		Script:         dcl.ValueOrEmptyString(o.Script),
		SecurityLevel:  AppengineVersionApiConfigSecurityLevelEnumToProto(o.SecurityLevel),
		Url:            dcl.ValueOrEmptyString(o.Url),
	}
	return p
}

// VersionDeploymentToProto converts a VersionDeployment resource to its proto representation.
func AppengineVersionDeploymentToProto(o *appengine.VersionDeployment) *appenginepb.AppengineVersionDeployment {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeployment{
		Container:         AppengineVersionDeploymentContainerToProto(o.Container),
		Zip:               AppengineVersionDeploymentZipToProto(o.Zip),
		CloudBuildOptions: AppengineVersionDeploymentCloudBuildOptionsToProto(o.CloudBuildOptions),
	}
	p.Files = make(map[string]*appenginepb.AppengineVersionDeploymentFiles)
	for k, r := range o.Files {
		p.Files[k] = AppengineVersionDeploymentFilesToProto(&r)
	}
	return p
}

// VersionDeploymentFilesToProto converts a VersionDeploymentFiles resource to its proto representation.
func AppengineVersionDeploymentFilesToProto(o *appengine.VersionDeploymentFiles) *appenginepb.AppengineVersionDeploymentFiles {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentFiles{
		SourceUrl: dcl.ValueOrEmptyString(o.SourceUrl),
		Sha1Sum:   dcl.ValueOrEmptyString(o.Sha1Sum),
		MimeType:  dcl.ValueOrEmptyString(o.MimeType),
	}
	return p
}

// VersionDeploymentContainerToProto converts a VersionDeploymentContainer resource to its proto representation.
func AppengineVersionDeploymentContainerToProto(o *appengine.VersionDeploymentContainer) *appenginepb.AppengineVersionDeploymentContainer {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentContainer{
		Image: dcl.ValueOrEmptyString(o.Image),
	}
	return p
}

// VersionDeploymentZipToProto converts a VersionDeploymentZip resource to its proto representation.
func AppengineVersionDeploymentZipToProto(o *appengine.VersionDeploymentZip) *appenginepb.AppengineVersionDeploymentZip {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentZip{
		SourceUrl:  dcl.ValueOrEmptyString(o.SourceUrl),
		FilesCount: dcl.ValueOrEmptyInt64(o.FilesCount),
	}
	return p
}

// VersionDeploymentCloudBuildOptionsToProto converts a VersionDeploymentCloudBuildOptions resource to its proto representation.
func AppengineVersionDeploymentCloudBuildOptionsToProto(o *appengine.VersionDeploymentCloudBuildOptions) *appenginepb.AppengineVersionDeploymentCloudBuildOptions {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionDeploymentCloudBuildOptions{
		AppYamlPath:       dcl.ValueOrEmptyString(o.AppYamlPath),
		CloudBuildTimeout: dcl.ValueOrEmptyString(o.CloudBuildTimeout),
	}
	return p
}

// VersionHealthCheckToProto converts a VersionHealthCheck resource to its proto representation.
func AppengineVersionHealthCheckToProto(o *appengine.VersionHealthCheck) *appenginepb.AppengineVersionHealthCheck {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionHealthCheck{
		DisableHealthCheck: dcl.ValueOrEmptyBool(o.DisableHealthCheck),
		Host:               dcl.ValueOrEmptyString(o.Host),
		HealthyThreshold:   dcl.ValueOrEmptyInt64(o.HealthyThreshold),
		UnhealthyThreshold: dcl.ValueOrEmptyInt64(o.UnhealthyThreshold),
		RestartThreshold:   dcl.ValueOrEmptyInt64(o.RestartThreshold),
		CheckInterval:      dcl.ValueOrEmptyString(o.CheckInterval),
		Timeout:            dcl.ValueOrEmptyString(o.Timeout),
	}
	return p
}

// VersionReadinessCheckToProto converts a VersionReadinessCheck resource to its proto representation.
func AppengineVersionReadinessCheckToProto(o *appengine.VersionReadinessCheck) *appenginepb.AppengineVersionReadinessCheck {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionReadinessCheck{
		Path:             dcl.ValueOrEmptyString(o.Path),
		Host:             dcl.ValueOrEmptyString(o.Host),
		FailureThreshold: dcl.ValueOrEmptyInt64(o.FailureThreshold),
		SuccessThreshold: dcl.ValueOrEmptyInt64(o.SuccessThreshold),
		CheckInterval:    dcl.ValueOrEmptyString(o.CheckInterval),
		Timeout:          dcl.ValueOrEmptyString(o.Timeout),
		AppStartTimeout:  dcl.ValueOrEmptyString(o.AppStartTimeout),
	}
	return p
}

// VersionLivenessCheckToProto converts a VersionLivenessCheck resource to its proto representation.
func AppengineVersionLivenessCheckToProto(o *appengine.VersionLivenessCheck) *appenginepb.AppengineVersionLivenessCheck {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionLivenessCheck{
		Path:             dcl.ValueOrEmptyString(o.Path),
		Host:             dcl.ValueOrEmptyString(o.Host),
		FailureThreshold: dcl.ValueOrEmptyInt64(o.FailureThreshold),
		SuccessThreshold: dcl.ValueOrEmptyInt64(o.SuccessThreshold),
		CheckInterval:    dcl.ValueOrEmptyString(o.CheckInterval),
		Timeout:          dcl.ValueOrEmptyString(o.Timeout),
		InitialDelay:     dcl.ValueOrEmptyString(o.InitialDelay),
	}
	return p
}

// VersionEntrypointToProto converts a VersionEntrypoint resource to its proto representation.
func AppengineVersionEntrypointToProto(o *appengine.VersionEntrypoint) *appenginepb.AppengineVersionEntrypoint {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionEntrypoint{
		Shell: dcl.ValueOrEmptyString(o.Shell),
	}
	return p
}

// VersionVPCAccessConnectorToProto converts a VersionVPCAccessConnector resource to its proto representation.
func AppengineVersionVPCAccessConnectorToProto(o *appengine.VersionVPCAccessConnector) *appenginepb.AppengineVersionVPCAccessConnector {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineVersionVPCAccessConnector{
		Name: dcl.ValueOrEmptyString(o.Name),
	}
	return p
}

// VersionToProto converts a Version resource to its proto representation.
func VersionToProto(resource *appengine.Version) *appenginepb.AppengineVersion {
	p := &appenginepb.AppengineVersion{
		ConsumerName:              dcl.ValueOrEmptyString(resource.ConsumerName),
		Name:                      dcl.ValueOrEmptyString(resource.Name),
		AutomaticScaling:          AppengineVersionAutomaticScalingToProto(resource.AutomaticScaling),
		BasicScaling:              AppengineVersionBasicScalingToProto(resource.BasicScaling),
		ManualScaling:             AppengineVersionManualScalingToProto(resource.ManualScaling),
		InstanceClass:             dcl.ValueOrEmptyString(resource.InstanceClass),
		Network:                   AppengineVersionNetworkToProto(resource.Network),
		Resources:                 AppengineVersionResourcesToProto(resource.Resources),
		Runtime:                   dcl.ValueOrEmptyString(resource.Runtime),
		RuntimeChannel:            dcl.ValueOrEmptyString(resource.RuntimeChannel),
		Threadsafe:                dcl.ValueOrEmptyBool(resource.Threadsafe),
		Vm:                        dcl.ValueOrEmptyBool(resource.Vm),
		Env:                       dcl.ValueOrEmptyString(resource.Env),
		ServingStatus:             AppengineVersionServingStatusEnumToProto(resource.ServingStatus),
		CreatedBy:                 dcl.ValueOrEmptyString(resource.CreatedBy),
		CreateTime:                dcl.ValueOrEmptyString(resource.CreateTime),
		DiskUsageBytes:            dcl.ValueOrEmptyInt64(resource.DiskUsageBytes),
		RuntimeApiVersion:         dcl.ValueOrEmptyString(resource.RuntimeApiVersion),
		RuntimeMainExecutablePath: dcl.ValueOrEmptyString(resource.RuntimeMainExecutablePath),
		ApiConfig:                 AppengineVersionApiConfigToProto(resource.ApiConfig),
		DefaultExpiration:         dcl.ValueOrEmptyString(resource.DefaultExpiration),
		Deployment:                AppengineVersionDeploymentToProto(resource.Deployment),
		HealthCheck:               AppengineVersionHealthCheckToProto(resource.HealthCheck),
		ReadinessCheck:            AppengineVersionReadinessCheckToProto(resource.ReadinessCheck),
		LivenessCheck:             AppengineVersionLivenessCheckToProto(resource.LivenessCheck),
		NobuildFilesRegex:         dcl.ValueOrEmptyString(resource.NobuildFilesRegex),
		VersionUrl:                dcl.ValueOrEmptyString(resource.VersionUrl),
		Entrypoint:                AppengineVersionEntrypointToProto(resource.Entrypoint),
		VpcAccessConnector:        AppengineVersionVPCAccessConnectorToProto(resource.VPCAccessConnector),
		App:                       dcl.ValueOrEmptyString(resource.App),
		Service:                   dcl.ValueOrEmptyString(resource.Service),
	}
	for _, r := range resource.InboundServices {
		p.InboundServices = append(p.InboundServices, appenginepb.AppengineVersionInboundServicesEnum(appenginepb.AppengineVersionInboundServicesEnum_value[string(r)]))
	}
	for _, r := range resource.Zones {
		p.Zones = append(p.Zones, r)
	}
	for _, r := range resource.Handlers {
		p.Handlers = append(p.Handlers, AppengineVersionHandlersToProto(&r))
	}
	for _, r := range resource.ErrorHandlers {
		p.ErrorHandlers = append(p.ErrorHandlers, AppengineVersionErrorHandlersToProto(&r))
	}
	for _, r := range resource.Libraries {
		p.Libraries = append(p.Libraries, AppengineVersionLibrariesToProto(&r))
	}

	return p
}

// ApplyVersion handles the gRPC request by passing it to the underlying Version Apply() method.
func (s *VersionServer) applyVersion(ctx context.Context, c *appengine.Client, request *appenginepb.ApplyAppengineVersionRequest) (*appenginepb.AppengineVersion, error) {
	p := ProtoToVersion(request.GetResource())
	res, err := c.ApplyVersion(ctx, p)
	if err != nil {
		return nil, err
	}
	r := VersionToProto(res)
	return r, nil
}

// ApplyVersion handles the gRPC request by passing it to the underlying Version Apply() method.
func (s *VersionServer) ApplyAppengineVersion(ctx context.Context, request *appenginepb.ApplyAppengineVersionRequest) (*appenginepb.AppengineVersion, error) {
	cl, err := createConfigVersion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyVersion(ctx, cl, request)
}

// DeleteVersion handles the gRPC request by passing it to the underlying Version Delete() method.
func (s *VersionServer) DeleteAppengineVersion(ctx context.Context, request *appenginepb.DeleteAppengineVersionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigVersion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteVersion(ctx, ProtoToVersion(request.GetResource()))

}

// ListAppengineVersion handles the gRPC request by passing it to the underlying VersionList() method.
func (s *VersionServer) ListAppengineVersion(ctx context.Context, request *appenginepb.ListAppengineVersionRequest) (*appenginepb.ListAppengineVersionResponse, error) {
	cl, err := createConfigVersion(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListVersion(ctx, request.App, request.Service)
	if err != nil {
		return nil, err
	}
	var protos []*appenginepb.AppengineVersion
	for _, r := range resources.Items {
		rp := VersionToProto(r)
		protos = append(protos, rp)
	}
	return &appenginepb.ListAppengineVersionResponse{Items: protos}, nil
}

func createConfigVersion(ctx context.Context, service_account_file string) (*appengine.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return appengine.NewClient(conf), nil
}
