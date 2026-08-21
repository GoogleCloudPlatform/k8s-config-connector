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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/assuredworkloads/beta/assuredworkloads_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta"
)

// WorkloadServer implements the gRPC interface for Workload.
type WorkloadServer struct{}

// ProtoToWorkloadResourcesResourceTypeEnum converts a WorkloadResourcesResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(e betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum) *beta.WorkloadResourcesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum_name[int32(e)]; ok {
		e := beta.WorkloadResourcesResourceTypeEnum(n[len("AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadComplianceRegimeEnum converts a WorkloadComplianceRegimeEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadComplianceRegimeEnum(e betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum) *beta.WorkloadComplianceRegimeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum_name[int32(e)]; ok {
		e := beta.WorkloadComplianceRegimeEnum(n[len("AssuredworkloadsBetaWorkloadComplianceRegimeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResourceSettingsResourceTypeEnum converts a WorkloadResourceSettingsResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(e betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum) *beta.WorkloadResourceSettingsResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum_name[int32(e)]; ok {
		e := beta.WorkloadResourceSettingsResourceTypeEnum(n[len("AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadKajEnrollmentStateEnum converts a WorkloadKajEnrollmentStateEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadKajEnrollmentStateEnum(e betapb.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum) *beta.WorkloadKajEnrollmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum_name[int32(e)]; ok {
		e := beta.WorkloadKajEnrollmentStateEnum(n[len("AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadSaaEnrollmentResponseSetupErrorsEnum converts a WorkloadSaaEnrollmentResponseSetupErrorsEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum(e betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum) *beta.WorkloadSaaEnrollmentResponseSetupErrorsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum_name[int32(e)]; ok {
		e := beta.WorkloadSaaEnrollmentResponseSetupErrorsEnum(n[len("AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadSaaEnrollmentResponseSetupStatusEnum converts a WorkloadSaaEnrollmentResponseSetupStatusEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum(e betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum) *beta.WorkloadSaaEnrollmentResponseSetupStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum_name[int32(e)]; ok {
		e := beta.WorkloadSaaEnrollmentResponseSetupStatusEnum(n[len("AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadPartnerEnum converts a WorkloadPartnerEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadPartnerEnum(e betapb.AssuredworkloadsBetaWorkloadPartnerEnum) *beta.WorkloadPartnerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadPartnerEnum_name[int32(e)]; ok {
		e := beta.WorkloadPartnerEnum(n[len("AssuredworkloadsBetaWorkloadPartnerEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningStateEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningStateEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(e betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum) *beta.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum_name[int32(e)]; ok {
		e := beta.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum(n[len("AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(e betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum) *beta.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum_name[int32(e)]; ok {
		e := beta.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(n[len("AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum enum from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(e betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum) *beta.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum_name[int32(e)]; ok {
		e := beta.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(n[len("AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResources converts a WorkloadResources object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadResources(p *betapb.AssuredworkloadsBetaWorkloadResources) *beta.WorkloadResources {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadResources{
		ResourceId:   dcl.Int64OrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkloadKmsSettings converts a WorkloadKmsSettings object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadKmsSettings(p *betapb.AssuredworkloadsBetaWorkloadKmsSettings) *beta.WorkloadKmsSettings {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadKmsSettings{
		NextRotationTime: dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:   dcl.StringOrNil(p.GetRotationPeriod()),
	}
	return obj
}

// ProtoToWorkloadResourceSettings converts a WorkloadResourceSettings object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadResourceSettings(p *betapb.AssuredworkloadsBetaWorkloadResourceSettings) *beta.WorkloadResourceSettings {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadResourceSettings{
		ResourceId:   dcl.StringOrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(p.GetResourceType()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
	}
	return obj
}

// ProtoToWorkloadSaaEnrollmentResponse converts a WorkloadSaaEnrollmentResponse object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadSaaEnrollmentResponse(p *betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponse) *beta.WorkloadSaaEnrollmentResponse {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadSaaEnrollmentResponse{
		SetupStatus: ProtoToAssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum(p.GetSetupStatus()),
	}
	for _, r := range p.GetSetupErrors() {
		obj.SetupErrors = append(obj.SetupErrors, *ProtoToAssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum(r))
	}
	return obj
}

// ProtoToWorkloadComplianceStatus converts a WorkloadComplianceStatus object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadComplianceStatus(p *betapb.AssuredworkloadsBetaWorkloadComplianceStatus) *beta.WorkloadComplianceStatus {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadComplianceStatus{}
	for _, r := range p.GetActiveViolationCount() {
		obj.ActiveViolationCount = append(obj.ActiveViolationCount, r)
	}
	for _, r := range p.GetAcknowledgedViolationCount() {
		obj.AcknowledgedViolationCount = append(obj.AcknowledgedViolationCount, r)
	}
	return obj
}

// ProtoToWorkloadPartnerPermissions converts a WorkloadPartnerPermissions object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadPartnerPermissions(p *betapb.AssuredworkloadsBetaWorkloadPartnerPermissions) *beta.WorkloadPartnerPermissions {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadPartnerPermissions{
		DataLogsViewer:             dcl.Bool(p.GetDataLogsViewer()),
		ServiceAccessApprover:      dcl.Bool(p.GetServiceAccessApprover()),
		AssuredWorkloadsMonitoring: dcl.Bool(p.GetAssuredWorkloadsMonitoring()),
	}
	return obj
}

// ProtoToWorkloadEkmProvisioningResponse converts a WorkloadEkmProvisioningResponse object from its proto representation.
func ProtoToAssuredworkloadsBetaWorkloadEkmProvisioningResponse(p *betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponse) *beta.WorkloadEkmProvisioningResponse {
	if p == nil {
		return nil
	}
	obj := &beta.WorkloadEkmProvisioningResponse{
		EkmProvisioningState:        ProtoToAssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(p.GetEkmProvisioningState()),
		EkmProvisioningErrorDomain:  ProtoToAssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(p.GetEkm_ProvisioningErrorDomain()),
		EkmProvisioningErrorMapping: ProtoToAssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(p.GetEkmProvisioningErrorMapping()),
	}
	return obj
}

// ProtoToWorkload converts a Workload resource from its proto representation.
func ProtoToWorkload(p *betapb.AssuredworkloadsBetaWorkload) *beta.Workload {
	obj := &beta.Workload{
		Name:                          dcl.StringOrNil(p.GetName()),
		DisplayName:                   dcl.StringOrNil(p.GetDisplayName()),
		ComplianceRegime:              ProtoToAssuredworkloadsBetaWorkloadComplianceRegimeEnum(p.GetComplianceRegime()),
		CreateTime:                    dcl.StringOrNil(p.GetCreateTime()),
		BillingAccount:                dcl.StringOrNil(p.GetBillingAccount()),
		ProvisionedResourcesParent:    dcl.StringOrNil(p.GetProvisionedResourcesParent()),
		KmsSettings:                   ProtoToAssuredworkloadsBetaWorkloadKmsSettings(p.GetKmsSettings()),
		KajEnrollmentState:            ProtoToAssuredworkloadsBetaWorkloadKajEnrollmentStateEnum(p.GetKajEnrollmentState()),
		EnableSovereignControls:       dcl.Bool(p.GetEnableSovereignControls()),
		SaaEnrollmentResponse:         ProtoToAssuredworkloadsBetaWorkloadSaaEnrollmentResponse(p.GetSaaEnrollmentResponse()),
		ComplianceStatus:              ProtoToAssuredworkloadsBetaWorkloadComplianceStatus(p.GetComplianceStatus()),
		Partner:                       ProtoToAssuredworkloadsBetaWorkloadPartnerEnum(p.GetPartner()),
		PartnerPermissions:            ProtoToAssuredworkloadsBetaWorkloadPartnerPermissions(p.GetPartnerPermissions()),
		EkmProvisioningResponse:       ProtoToAssuredworkloadsBetaWorkloadEkmProvisioningResponse(p.GetEkmProvisioningResponse()),
		ViolationNotificationsEnabled: dcl.Bool(p.GetViolationNotificationsEnabled()),
		Organization:                  dcl.StringOrNil(p.GetOrganization()),
		Location:                      dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToAssuredworkloadsBetaWorkloadResources(r))
	}
	for _, r := range p.GetResourceSettings() {
		obj.ResourceSettings = append(obj.ResourceSettings, *ProtoToAssuredworkloadsBetaWorkloadResourceSettings(r))
	}
	for _, r := range p.GetCompliantButDisallowedServices() {
		obj.CompliantButDisallowedServices = append(obj.CompliantButDisallowedServices, r)
	}
	return obj
}

// WorkloadResourcesResourceTypeEnumToProto converts a WorkloadResourcesResourceTypeEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadResourcesResourceTypeEnumToProto(e *beta.WorkloadResourcesResourceTypeEnum) betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum_value["WorkloadResourcesResourceTypeEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadResourcesResourceTypeEnum(0)
}

// WorkloadComplianceRegimeEnumToProto converts a WorkloadComplianceRegimeEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadComplianceRegimeEnumToProto(e *beta.WorkloadComplianceRegimeEnum) betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum_value["WorkloadComplianceRegimeEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadComplianceRegimeEnum(0)
}

// WorkloadResourceSettingsResourceTypeEnumToProto converts a WorkloadResourceSettingsResourceTypeEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnumToProto(e *beta.WorkloadResourceSettingsResourceTypeEnum) betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum_value["WorkloadResourceSettingsResourceTypeEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnum(0)
}

// WorkloadKajEnrollmentStateEnumToProto converts a WorkloadKajEnrollmentStateEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadKajEnrollmentStateEnumToProto(e *beta.WorkloadKajEnrollmentStateEnum) betapb.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum_value["WorkloadKajEnrollmentStateEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadKajEnrollmentStateEnum(0)
}

// WorkloadSaaEnrollmentResponseSetupErrorsEnumToProto converts a WorkloadSaaEnrollmentResponseSetupErrorsEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnumToProto(e *beta.WorkloadSaaEnrollmentResponseSetupErrorsEnum) betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum_value["WorkloadSaaEnrollmentResponseSetupErrorsEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum(0)
}

// WorkloadSaaEnrollmentResponseSetupStatusEnumToProto converts a WorkloadSaaEnrollmentResponseSetupStatusEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnumToProto(e *beta.WorkloadSaaEnrollmentResponseSetupStatusEnum) betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum_value["WorkloadSaaEnrollmentResponseSetupStatusEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnum(0)
}

// WorkloadPartnerEnumToProto converts a WorkloadPartnerEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadPartnerEnumToProto(e *beta.WorkloadPartnerEnum) betapb.AssuredworkloadsBetaWorkloadPartnerEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadPartnerEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadPartnerEnum_value["WorkloadPartnerEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadPartnerEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadPartnerEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningStateEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto(e *beta.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum) betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningStateEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto(e *beta.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum) betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum enum to its proto representation.
func AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto(e *beta.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum) betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	if e == nil {
		return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(0)
	}
	if v, ok := betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum"+string(*e)]; ok {
		return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(v)
	}
	return betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(0)
}

// WorkloadResourcesToProto converts a WorkloadResources object to its proto representation.
func AssuredworkloadsBetaWorkloadResourcesToProto(o *beta.WorkloadResources) *betapb.AssuredworkloadsBetaWorkloadResources {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadResources{}
	p.SetResourceId(dcl.ValueOrEmptyInt64(o.ResourceId))
	p.SetResourceType(AssuredworkloadsBetaWorkloadResourcesResourceTypeEnumToProto(o.ResourceType))
	return p
}

// WorkloadKmsSettingsToProto converts a WorkloadKmsSettings object to its proto representation.
func AssuredworkloadsBetaWorkloadKmsSettingsToProto(o *beta.WorkloadKmsSettings) *betapb.AssuredworkloadsBetaWorkloadKmsSettings {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadKmsSettings{}
	p.SetNextRotationTime(dcl.ValueOrEmptyString(o.NextRotationTime))
	p.SetRotationPeriod(dcl.ValueOrEmptyString(o.RotationPeriod))
	return p
}

// WorkloadResourceSettingsToProto converts a WorkloadResourceSettings object to its proto representation.
func AssuredworkloadsBetaWorkloadResourceSettingsToProto(o *beta.WorkloadResourceSettings) *betapb.AssuredworkloadsBetaWorkloadResourceSettings {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadResourceSettings{}
	p.SetResourceId(dcl.ValueOrEmptyString(o.ResourceId))
	p.SetResourceType(AssuredworkloadsBetaWorkloadResourceSettingsResourceTypeEnumToProto(o.ResourceType))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	return p
}

// WorkloadSaaEnrollmentResponseToProto converts a WorkloadSaaEnrollmentResponse object to its proto representation.
func AssuredworkloadsBetaWorkloadSaaEnrollmentResponseToProto(o *beta.WorkloadSaaEnrollmentResponse) *betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponse {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponse{}
	p.SetSetupStatus(AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupStatusEnumToProto(o.SetupStatus))
	sSetupErrors := make([]betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum, len(o.SetupErrors))
	for i, r := range o.SetupErrors {
		sSetupErrors[i] = betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum(betapb.AssuredworkloadsBetaWorkloadSaaEnrollmentResponseSetupErrorsEnum_value[string(r)])
	}
	p.SetSetupErrors(sSetupErrors)
	return p
}

// WorkloadComplianceStatusToProto converts a WorkloadComplianceStatus object to its proto representation.
func AssuredworkloadsBetaWorkloadComplianceStatusToProto(o *beta.WorkloadComplianceStatus) *betapb.AssuredworkloadsBetaWorkloadComplianceStatus {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadComplianceStatus{}
	sActiveViolationCount := make([]int64, len(o.ActiveViolationCount))
	for i, r := range o.ActiveViolationCount {
		sActiveViolationCount[i] = r
	}
	p.SetActiveViolationCount(sActiveViolationCount)
	sAcknowledgedViolationCount := make([]int64, len(o.AcknowledgedViolationCount))
	for i, r := range o.AcknowledgedViolationCount {
		sAcknowledgedViolationCount[i] = r
	}
	p.SetAcknowledgedViolationCount(sAcknowledgedViolationCount)
	return p
}

// WorkloadPartnerPermissionsToProto converts a WorkloadPartnerPermissions object to its proto representation.
func AssuredworkloadsBetaWorkloadPartnerPermissionsToProto(o *beta.WorkloadPartnerPermissions) *betapb.AssuredworkloadsBetaWorkloadPartnerPermissions {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadPartnerPermissions{}
	p.SetDataLogsViewer(dcl.ValueOrEmptyBool(o.DataLogsViewer))
	p.SetServiceAccessApprover(dcl.ValueOrEmptyBool(o.ServiceAccessApprover))
	p.SetAssuredWorkloadsMonitoring(dcl.ValueOrEmptyBool(o.AssuredWorkloadsMonitoring))
	return p
}

// WorkloadEkmProvisioningResponseToProto converts a WorkloadEkmProvisioningResponse object to its proto representation.
func AssuredworkloadsBetaWorkloadEkmProvisioningResponseToProto(o *beta.WorkloadEkmProvisioningResponse) *betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponse {
	if o == nil {
		return nil
	}
	p := &betapb.AssuredworkloadsBetaWorkloadEkmProvisioningResponse{}
	p.SetEkmProvisioningState(AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto(o.EkmProvisioningState))
	p.SetEkm_ProvisioningErrorDomain(AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto(o.EkmProvisioningErrorDomain))
	p.SetEkmProvisioningErrorMapping(AssuredworkloadsBetaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto(o.EkmProvisioningErrorMapping))
	return p
}

// WorkloadToProto converts a Workload resource to its proto representation.
func WorkloadToProto(resource *beta.Workload) *betapb.AssuredworkloadsBetaWorkload {
	p := &betapb.AssuredworkloadsBetaWorkload{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetComplianceRegime(AssuredworkloadsBetaWorkloadComplianceRegimeEnumToProto(resource.ComplianceRegime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetBillingAccount(dcl.ValueOrEmptyString(resource.BillingAccount))
	p.SetProvisionedResourcesParent(dcl.ValueOrEmptyString(resource.ProvisionedResourcesParent))
	p.SetKmsSettings(AssuredworkloadsBetaWorkloadKmsSettingsToProto(resource.KmsSettings))
	p.SetKajEnrollmentState(AssuredworkloadsBetaWorkloadKajEnrollmentStateEnumToProto(resource.KajEnrollmentState))
	p.SetEnableSovereignControls(dcl.ValueOrEmptyBool(resource.EnableSovereignControls))
	p.SetSaaEnrollmentResponse(AssuredworkloadsBetaWorkloadSaaEnrollmentResponseToProto(resource.SaaEnrollmentResponse))
	p.SetComplianceStatus(AssuredworkloadsBetaWorkloadComplianceStatusToProto(resource.ComplianceStatus))
	p.SetPartner(AssuredworkloadsBetaWorkloadPartnerEnumToProto(resource.Partner))
	p.SetPartnerPermissions(AssuredworkloadsBetaWorkloadPartnerPermissionsToProto(resource.PartnerPermissions))
	p.SetEkmProvisioningResponse(AssuredworkloadsBetaWorkloadEkmProvisioningResponseToProto(resource.EkmProvisioningResponse))
	p.SetViolationNotificationsEnabled(dcl.ValueOrEmptyBool(resource.ViolationNotificationsEnabled))
	p.SetOrganization(dcl.ValueOrEmptyString(resource.Organization))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sResources := make([]*betapb.AssuredworkloadsBetaWorkloadResources, len(resource.Resources))
	for i, r := range resource.Resources {
		sResources[i] = AssuredworkloadsBetaWorkloadResourcesToProto(&r)
	}
	p.SetResources(sResources)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sResourceSettings := make([]*betapb.AssuredworkloadsBetaWorkloadResourceSettings, len(resource.ResourceSettings))
	for i, r := range resource.ResourceSettings {
		sResourceSettings[i] = AssuredworkloadsBetaWorkloadResourceSettingsToProto(&r)
	}
	p.SetResourceSettings(sResourceSettings)
	sCompliantButDisallowedServices := make([]string, len(resource.CompliantButDisallowedServices))
	for i, r := range resource.CompliantButDisallowedServices {
		sCompliantButDisallowedServices[i] = r
	}
	p.SetCompliantButDisallowedServices(sCompliantButDisallowedServices)

	return p
}

// applyWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) applyWorkload(ctx context.Context, c *beta.Client, request *betapb.ApplyAssuredworkloadsBetaWorkloadRequest) (*betapb.AssuredworkloadsBetaWorkload, error) {
	p := ProtoToWorkload(request.GetResource())
	res, err := c.ApplyWorkload(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadToProto(res)
	return r, nil
}

// applyAssuredworkloadsBetaWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) ApplyAssuredworkloadsBetaWorkload(ctx context.Context, request *betapb.ApplyAssuredworkloadsBetaWorkloadRequest) (*betapb.AssuredworkloadsBetaWorkload, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkload(ctx, cl, request)
}

// DeleteWorkload handles the gRPC request by passing it to the underlying Workload Delete() method.
func (s *WorkloadServer) DeleteAssuredworkloadsBetaWorkload(ctx context.Context, request *betapb.DeleteAssuredworkloadsBetaWorkloadRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkload(ctx, ProtoToWorkload(request.GetResource()))

}

// ListAssuredworkloadsBetaWorkload handles the gRPC request by passing it to the underlying WorkloadList() method.
func (s *WorkloadServer) ListAssuredworkloadsBetaWorkload(ctx context.Context, request *betapb.ListAssuredworkloadsBetaWorkloadRequest) (*betapb.ListAssuredworkloadsBetaWorkloadResponse, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkload(ctx, request.GetOrganization(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.AssuredworkloadsBetaWorkload
	for _, r := range resources.Items {
		rp := WorkloadToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListAssuredworkloadsBetaWorkloadResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkload(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
