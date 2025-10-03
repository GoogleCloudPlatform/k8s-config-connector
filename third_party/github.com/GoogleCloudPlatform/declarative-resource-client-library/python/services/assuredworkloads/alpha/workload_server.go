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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/assuredworkloads/alpha/assuredworkloads_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/alpha"
)

// WorkloadServer implements the gRPC interface for Workload.
type WorkloadServer struct{}

// ProtoToWorkloadResourcesResourceTypeEnum converts a WorkloadResourcesResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(e alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum) *alpha.WorkloadResourcesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum_name[int32(e)]; ok {
		e := alpha.WorkloadResourcesResourceTypeEnum(n[len("AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadComplianceRegimeEnum converts a WorkloadComplianceRegimeEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadComplianceRegimeEnum(e alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum) *alpha.WorkloadComplianceRegimeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum_name[int32(e)]; ok {
		e := alpha.WorkloadComplianceRegimeEnum(n[len("AssuredworkloadsAlphaWorkloadComplianceRegimeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResourceSettingsResourceTypeEnum converts a WorkloadResourceSettingsResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(e alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum) *alpha.WorkloadResourceSettingsResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum_name[int32(e)]; ok {
		e := alpha.WorkloadResourceSettingsResourceTypeEnum(n[len("AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadKajEnrollmentStateEnum converts a WorkloadKajEnrollmentStateEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum(e alphapb.AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum) *alpha.WorkloadKajEnrollmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum_name[int32(e)]; ok {
		e := alpha.WorkloadKajEnrollmentStateEnum(n[len("AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadSaaEnrollmentResponseSetupErrorsEnum converts a WorkloadSaaEnrollmentResponseSetupErrorsEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum(e alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum) *alpha.WorkloadSaaEnrollmentResponseSetupErrorsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum_name[int32(e)]; ok {
		e := alpha.WorkloadSaaEnrollmentResponseSetupErrorsEnum(n[len("AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadSaaEnrollmentResponseSetupStatusEnum converts a WorkloadSaaEnrollmentResponseSetupStatusEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum(e alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum) *alpha.WorkloadSaaEnrollmentResponseSetupStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum_name[int32(e)]; ok {
		e := alpha.WorkloadSaaEnrollmentResponseSetupStatusEnum(n[len("AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadPartnerEnum converts a WorkloadPartnerEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadPartnerEnum(e alphapb.AssuredworkloadsAlphaWorkloadPartnerEnum) *alpha.WorkloadPartnerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadPartnerEnum_name[int32(e)]; ok {
		e := alpha.WorkloadPartnerEnum(n[len("AssuredworkloadsAlphaWorkloadPartnerEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningStateEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningStateEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(e alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum) *alpha.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum_name[int32(e)]; ok {
		e := alpha.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum(n[len("AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(e alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum) *alpha.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum_name[int32(e)]; ok {
		e := alpha.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(n[len("AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum enum from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(e alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum) *alpha.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum_name[int32(e)]; ok {
		e := alpha.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(n[len("AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResources converts a WorkloadResources object from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadResources(p *alphapb.AssuredworkloadsAlphaWorkloadResources) *alpha.WorkloadResources {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadResources{
		ResourceId:   dcl.Int64OrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkloadKmsSettings converts a WorkloadKmsSettings object from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadKmsSettings(p *alphapb.AssuredworkloadsAlphaWorkloadKmsSettings) *alpha.WorkloadKmsSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadKmsSettings{
		NextRotationTime: dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:   dcl.StringOrNil(p.GetRotationPeriod()),
	}
	return obj
}

// ProtoToWorkloadResourceSettings converts a WorkloadResourceSettings object from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadResourceSettings(p *alphapb.AssuredworkloadsAlphaWorkloadResourceSettings) *alpha.WorkloadResourceSettings {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadResourceSettings{
		ResourceId:   dcl.StringOrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(p.GetResourceType()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
	}
	return obj
}

// ProtoToWorkloadSaaEnrollmentResponse converts a WorkloadSaaEnrollmentResponse object from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadSaaEnrollmentResponse(p *alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponse) *alpha.WorkloadSaaEnrollmentResponse {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadSaaEnrollmentResponse{
		SetupStatus: ProtoToAssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum(p.GetSetupStatus()),
	}
	for _, r := range p.GetSetupErrors() {
		obj.SetupErrors = append(obj.SetupErrors, *ProtoToAssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum(r))
	}
	return obj
}

// ProtoToWorkloadComplianceStatus converts a WorkloadComplianceStatus object from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadComplianceStatus(p *alphapb.AssuredworkloadsAlphaWorkloadComplianceStatus) *alpha.WorkloadComplianceStatus {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadComplianceStatus{}
	for _, r := range p.GetActiveViolationCount() {
		obj.ActiveViolationCount = append(obj.ActiveViolationCount, r)
	}
	for _, r := range p.GetAcknowledgedViolationCount() {
		obj.AcknowledgedViolationCount = append(obj.AcknowledgedViolationCount, r)
	}
	return obj
}

// ProtoToWorkloadPartnerPermissions converts a WorkloadPartnerPermissions object from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadPartnerPermissions(p *alphapb.AssuredworkloadsAlphaWorkloadPartnerPermissions) *alpha.WorkloadPartnerPermissions {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadPartnerPermissions{
		DataLogsViewer:             dcl.Bool(p.GetDataLogsViewer()),
		ServiceAccessApprover:      dcl.Bool(p.GetServiceAccessApprover()),
		AssuredWorkloadsMonitoring: dcl.Bool(p.GetAssuredWorkloadsMonitoring()),
	}
	return obj
}

// ProtoToWorkloadEkmProvisioningResponse converts a WorkloadEkmProvisioningResponse object from its proto representation.
func ProtoToAssuredworkloadsAlphaWorkloadEkmProvisioningResponse(p *alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponse) *alpha.WorkloadEkmProvisioningResponse {
	if p == nil {
		return nil
	}
	obj := &alpha.WorkloadEkmProvisioningResponse{
		EkmProvisioningState:        ProtoToAssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(p.GetEkmProvisioningState()),
		EkmProvisioningErrorDomain:  ProtoToAssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(p.GetEkm_ProvisioningErrorDomain()),
		EkmProvisioningErrorMapping: ProtoToAssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(p.GetEkmProvisioningErrorMapping()),
	}
	return obj
}

// ProtoToWorkload converts a Workload resource from its proto representation.
func ProtoToWorkload(p *alphapb.AssuredworkloadsAlphaWorkload) *alpha.Workload {
	obj := &alpha.Workload{
		Name:                          dcl.StringOrNil(p.GetName()),
		DisplayName:                   dcl.StringOrNil(p.GetDisplayName()),
		ComplianceRegime:              ProtoToAssuredworkloadsAlphaWorkloadComplianceRegimeEnum(p.GetComplianceRegime()),
		CreateTime:                    dcl.StringOrNil(p.GetCreateTime()),
		BillingAccount:                dcl.StringOrNil(p.GetBillingAccount()),
		ProvisionedResourcesParent:    dcl.StringOrNil(p.GetProvisionedResourcesParent()),
		KmsSettings:                   ProtoToAssuredworkloadsAlphaWorkloadKmsSettings(p.GetKmsSettings()),
		KajEnrollmentState:            ProtoToAssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum(p.GetKajEnrollmentState()),
		EnableSovereignControls:       dcl.Bool(p.GetEnableSovereignControls()),
		SaaEnrollmentResponse:         ProtoToAssuredworkloadsAlphaWorkloadSaaEnrollmentResponse(p.GetSaaEnrollmentResponse()),
		ComplianceStatus:              ProtoToAssuredworkloadsAlphaWorkloadComplianceStatus(p.GetComplianceStatus()),
		Partner:                       ProtoToAssuredworkloadsAlphaWorkloadPartnerEnum(p.GetPartner()),
		PartnerPermissions:            ProtoToAssuredworkloadsAlphaWorkloadPartnerPermissions(p.GetPartnerPermissions()),
		EkmProvisioningResponse:       ProtoToAssuredworkloadsAlphaWorkloadEkmProvisioningResponse(p.GetEkmProvisioningResponse()),
		ViolationNotificationsEnabled: dcl.Bool(p.GetViolationNotificationsEnabled()),
		Organization:                  dcl.StringOrNil(p.GetOrganization()),
		Location:                      dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToAssuredworkloadsAlphaWorkloadResources(r))
	}
	for _, r := range p.GetResourceSettings() {
		obj.ResourceSettings = append(obj.ResourceSettings, *ProtoToAssuredworkloadsAlphaWorkloadResourceSettings(r))
	}
	for _, r := range p.GetCompliantButDisallowedServices() {
		obj.CompliantButDisallowedServices = append(obj.CompliantButDisallowedServices, r)
	}
	return obj
}

// WorkloadResourcesResourceTypeEnumToProto converts a WorkloadResourcesResourceTypeEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnumToProto(e *alpha.WorkloadResourcesResourceTypeEnum) alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum_value["WorkloadResourcesResourceTypeEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnum(0)
}

// WorkloadComplianceRegimeEnumToProto converts a WorkloadComplianceRegimeEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadComplianceRegimeEnumToProto(e *alpha.WorkloadComplianceRegimeEnum) alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum_value["WorkloadComplianceRegimeEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadComplianceRegimeEnum(0)
}

// WorkloadResourceSettingsResourceTypeEnumToProto converts a WorkloadResourceSettingsResourceTypeEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnumToProto(e *alpha.WorkloadResourceSettingsResourceTypeEnum) alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum_value["WorkloadResourceSettingsResourceTypeEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnum(0)
}

// WorkloadKajEnrollmentStateEnumToProto converts a WorkloadKajEnrollmentStateEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnumToProto(e *alpha.WorkloadKajEnrollmentStateEnum) alphapb.AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum_value["WorkloadKajEnrollmentStateEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnum(0)
}

// WorkloadSaaEnrollmentResponseSetupErrorsEnumToProto converts a WorkloadSaaEnrollmentResponseSetupErrorsEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnumToProto(e *alpha.WorkloadSaaEnrollmentResponseSetupErrorsEnum) alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum_value["WorkloadSaaEnrollmentResponseSetupErrorsEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum(0)
}

// WorkloadSaaEnrollmentResponseSetupStatusEnumToProto converts a WorkloadSaaEnrollmentResponseSetupStatusEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnumToProto(e *alpha.WorkloadSaaEnrollmentResponseSetupStatusEnum) alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum_value["WorkloadSaaEnrollmentResponseSetupStatusEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnum(0)
}

// WorkloadPartnerEnumToProto converts a WorkloadPartnerEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadPartnerEnumToProto(e *alpha.WorkloadPartnerEnum) alphapb.AssuredworkloadsAlphaWorkloadPartnerEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadPartnerEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadPartnerEnum_value["WorkloadPartnerEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadPartnerEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadPartnerEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningStateEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto(e *alpha.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum) alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningStateEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto(e *alpha.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum) alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum enum to its proto representation.
func AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto(e *alpha.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum) alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	if e == nil {
		return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(0)
	}
	if v, ok := alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum"+string(*e)]; ok {
		return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(v)
	}
	return alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(0)
}

// WorkloadResourcesToProto converts a WorkloadResources object to its proto representation.
func AssuredworkloadsAlphaWorkloadResourcesToProto(o *alpha.WorkloadResources) *alphapb.AssuredworkloadsAlphaWorkloadResources {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadResources{}
	p.SetResourceId(dcl.ValueOrEmptyInt64(o.ResourceId))
	p.SetResourceType(AssuredworkloadsAlphaWorkloadResourcesResourceTypeEnumToProto(o.ResourceType))
	return p
}

// WorkloadKmsSettingsToProto converts a WorkloadKmsSettings object to its proto representation.
func AssuredworkloadsAlphaWorkloadKmsSettingsToProto(o *alpha.WorkloadKmsSettings) *alphapb.AssuredworkloadsAlphaWorkloadKmsSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadKmsSettings{}
	p.SetNextRotationTime(dcl.ValueOrEmptyString(o.NextRotationTime))
	p.SetRotationPeriod(dcl.ValueOrEmptyString(o.RotationPeriod))
	return p
}

// WorkloadResourceSettingsToProto converts a WorkloadResourceSettings object to its proto representation.
func AssuredworkloadsAlphaWorkloadResourceSettingsToProto(o *alpha.WorkloadResourceSettings) *alphapb.AssuredworkloadsAlphaWorkloadResourceSettings {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadResourceSettings{}
	p.SetResourceId(dcl.ValueOrEmptyString(o.ResourceId))
	p.SetResourceType(AssuredworkloadsAlphaWorkloadResourceSettingsResourceTypeEnumToProto(o.ResourceType))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	return p
}

// WorkloadSaaEnrollmentResponseToProto converts a WorkloadSaaEnrollmentResponse object to its proto representation.
func AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseToProto(o *alpha.WorkloadSaaEnrollmentResponse) *alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponse {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponse{}
	p.SetSetupStatus(AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupStatusEnumToProto(o.SetupStatus))
	sSetupErrors := make([]alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum, len(o.SetupErrors))
	for i, r := range o.SetupErrors {
		sSetupErrors[i] = alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum(alphapb.AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseSetupErrorsEnum_value[string(r)])
	}
	p.SetSetupErrors(sSetupErrors)
	return p
}

// WorkloadComplianceStatusToProto converts a WorkloadComplianceStatus object to its proto representation.
func AssuredworkloadsAlphaWorkloadComplianceStatusToProto(o *alpha.WorkloadComplianceStatus) *alphapb.AssuredworkloadsAlphaWorkloadComplianceStatus {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadComplianceStatus{}
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
func AssuredworkloadsAlphaWorkloadPartnerPermissionsToProto(o *alpha.WorkloadPartnerPermissions) *alphapb.AssuredworkloadsAlphaWorkloadPartnerPermissions {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadPartnerPermissions{}
	p.SetDataLogsViewer(dcl.ValueOrEmptyBool(o.DataLogsViewer))
	p.SetServiceAccessApprover(dcl.ValueOrEmptyBool(o.ServiceAccessApprover))
	p.SetAssuredWorkloadsMonitoring(dcl.ValueOrEmptyBool(o.AssuredWorkloadsMonitoring))
	return p
}

// WorkloadEkmProvisioningResponseToProto converts a WorkloadEkmProvisioningResponse object to its proto representation.
func AssuredworkloadsAlphaWorkloadEkmProvisioningResponseToProto(o *alpha.WorkloadEkmProvisioningResponse) *alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponse {
	if o == nil {
		return nil
	}
	p := &alphapb.AssuredworkloadsAlphaWorkloadEkmProvisioningResponse{}
	p.SetEkmProvisioningState(AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto(o.EkmProvisioningState))
	p.SetEkm_ProvisioningErrorDomain(AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto(o.EkmProvisioningErrorDomain))
	p.SetEkmProvisioningErrorMapping(AssuredworkloadsAlphaWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto(o.EkmProvisioningErrorMapping))
	return p
}

// WorkloadToProto converts a Workload resource to its proto representation.
func WorkloadToProto(resource *alpha.Workload) *alphapb.AssuredworkloadsAlphaWorkload {
	p := &alphapb.AssuredworkloadsAlphaWorkload{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetComplianceRegime(AssuredworkloadsAlphaWorkloadComplianceRegimeEnumToProto(resource.ComplianceRegime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetBillingAccount(dcl.ValueOrEmptyString(resource.BillingAccount))
	p.SetProvisionedResourcesParent(dcl.ValueOrEmptyString(resource.ProvisionedResourcesParent))
	p.SetKmsSettings(AssuredworkloadsAlphaWorkloadKmsSettingsToProto(resource.KmsSettings))
	p.SetKajEnrollmentState(AssuredworkloadsAlphaWorkloadKajEnrollmentStateEnumToProto(resource.KajEnrollmentState))
	p.SetEnableSovereignControls(dcl.ValueOrEmptyBool(resource.EnableSovereignControls))
	p.SetSaaEnrollmentResponse(AssuredworkloadsAlphaWorkloadSaaEnrollmentResponseToProto(resource.SaaEnrollmentResponse))
	p.SetComplianceStatus(AssuredworkloadsAlphaWorkloadComplianceStatusToProto(resource.ComplianceStatus))
	p.SetPartner(AssuredworkloadsAlphaWorkloadPartnerEnumToProto(resource.Partner))
	p.SetPartnerPermissions(AssuredworkloadsAlphaWorkloadPartnerPermissionsToProto(resource.PartnerPermissions))
	p.SetEkmProvisioningResponse(AssuredworkloadsAlphaWorkloadEkmProvisioningResponseToProto(resource.EkmProvisioningResponse))
	p.SetViolationNotificationsEnabled(dcl.ValueOrEmptyBool(resource.ViolationNotificationsEnabled))
	p.SetOrganization(dcl.ValueOrEmptyString(resource.Organization))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sResources := make([]*alphapb.AssuredworkloadsAlphaWorkloadResources, len(resource.Resources))
	for i, r := range resource.Resources {
		sResources[i] = AssuredworkloadsAlphaWorkloadResourcesToProto(&r)
	}
	p.SetResources(sResources)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sResourceSettings := make([]*alphapb.AssuredworkloadsAlphaWorkloadResourceSettings, len(resource.ResourceSettings))
	for i, r := range resource.ResourceSettings {
		sResourceSettings[i] = AssuredworkloadsAlphaWorkloadResourceSettingsToProto(&r)
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
func (s *WorkloadServer) applyWorkload(ctx context.Context, c *alpha.Client, request *alphapb.ApplyAssuredworkloadsAlphaWorkloadRequest) (*alphapb.AssuredworkloadsAlphaWorkload, error) {
	p := ProtoToWorkload(request.GetResource())
	res, err := c.ApplyWorkload(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadToProto(res)
	return r, nil
}

// applyAssuredworkloadsAlphaWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) ApplyAssuredworkloadsAlphaWorkload(ctx context.Context, request *alphapb.ApplyAssuredworkloadsAlphaWorkloadRequest) (*alphapb.AssuredworkloadsAlphaWorkload, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkload(ctx, cl, request)
}

// DeleteWorkload handles the gRPC request by passing it to the underlying Workload Delete() method.
func (s *WorkloadServer) DeleteAssuredworkloadsAlphaWorkload(ctx context.Context, request *alphapb.DeleteAssuredworkloadsAlphaWorkloadRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkload(ctx, ProtoToWorkload(request.GetResource()))

}

// ListAssuredworkloadsAlphaWorkload handles the gRPC request by passing it to the underlying WorkloadList() method.
func (s *WorkloadServer) ListAssuredworkloadsAlphaWorkload(ctx context.Context, request *alphapb.ListAssuredworkloadsAlphaWorkloadRequest) (*alphapb.ListAssuredworkloadsAlphaWorkloadResponse, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkload(ctx, request.GetOrganization(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.AssuredworkloadsAlphaWorkload
	for _, r := range resources.Items {
		rp := WorkloadToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListAssuredworkloadsAlphaWorkloadResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkload(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
