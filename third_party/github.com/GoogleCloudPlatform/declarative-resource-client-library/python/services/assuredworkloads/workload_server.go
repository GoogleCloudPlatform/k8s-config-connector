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
	assuredworkloadspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/assuredworkloads/assuredworkloads_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads"
)

// WorkloadServer implements the gRPC interface for Workload.
type WorkloadServer struct{}

// ProtoToWorkloadResourcesResourceTypeEnum converts a WorkloadResourcesResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadResourcesResourceTypeEnum(e assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum) *assuredworkloads.WorkloadResourcesResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadResourcesResourceTypeEnum(n[len("AssuredworkloadsWorkloadResourcesResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadComplianceRegimeEnum converts a WorkloadComplianceRegimeEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadComplianceRegimeEnum(e assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum) *assuredworkloads.WorkloadComplianceRegimeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadComplianceRegimeEnum(n[len("AssuredworkloadsWorkloadComplianceRegimeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResourceSettingsResourceTypeEnum converts a WorkloadResourceSettingsResourceTypeEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(e assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum) *assuredworkloads.WorkloadResourceSettingsResourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadResourceSettingsResourceTypeEnum(n[len("AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadKajEnrollmentStateEnum converts a WorkloadKajEnrollmentStateEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadKajEnrollmentStateEnum(e assuredworkloadspb.AssuredworkloadsWorkloadKajEnrollmentStateEnum) *assuredworkloads.WorkloadKajEnrollmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadKajEnrollmentStateEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadKajEnrollmentStateEnum(n[len("AssuredworkloadsWorkloadKajEnrollmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadSaaEnrollmentResponseSetupErrorsEnum converts a WorkloadSaaEnrollmentResponseSetupErrorsEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum(e assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum) *assuredworkloads.WorkloadSaaEnrollmentResponseSetupErrorsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadSaaEnrollmentResponseSetupErrorsEnum(n[len("AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadSaaEnrollmentResponseSetupStatusEnum converts a WorkloadSaaEnrollmentResponseSetupStatusEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum(e assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum) *assuredworkloads.WorkloadSaaEnrollmentResponseSetupStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadSaaEnrollmentResponseSetupStatusEnum(n[len("AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadPartnerEnum converts a WorkloadPartnerEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadPartnerEnum(e assuredworkloadspb.AssuredworkloadsWorkloadPartnerEnum) *assuredworkloads.WorkloadPartnerEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadPartnerEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadPartnerEnum(n[len("AssuredworkloadsWorkloadPartnerEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningStateEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningStateEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(e assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum) *assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum(n[len("AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(e assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum) *assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(n[len("AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum enum from its proto representation.
func ProtoToAssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(e assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum) *assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	if e == 0 {
		return nil
	}
	if n, ok := assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum_name[int32(e)]; ok {
		e := assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(n[len("AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadResources converts a WorkloadResources object from its proto representation.
func ProtoToAssuredworkloadsWorkloadResources(p *assuredworkloadspb.AssuredworkloadsWorkloadResources) *assuredworkloads.WorkloadResources {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadResources{
		ResourceId:   dcl.Int64OrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsWorkloadResourcesResourceTypeEnum(p.GetResourceType()),
	}
	return obj
}

// ProtoToWorkloadKmsSettings converts a WorkloadKmsSettings object from its proto representation.
func ProtoToAssuredworkloadsWorkloadKmsSettings(p *assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings) *assuredworkloads.WorkloadKmsSettings {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadKmsSettings{
		NextRotationTime: dcl.StringOrNil(p.GetNextRotationTime()),
		RotationPeriod:   dcl.StringOrNil(p.GetRotationPeriod()),
	}
	return obj
}

// ProtoToWorkloadResourceSettings converts a WorkloadResourceSettings object from its proto representation.
func ProtoToAssuredworkloadsWorkloadResourceSettings(p *assuredworkloadspb.AssuredworkloadsWorkloadResourceSettings) *assuredworkloads.WorkloadResourceSettings {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadResourceSettings{
		ResourceId:   dcl.StringOrNil(p.GetResourceId()),
		ResourceType: ProtoToAssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(p.GetResourceType()),
		DisplayName:  dcl.StringOrNil(p.GetDisplayName()),
	}
	return obj
}

// ProtoToWorkloadSaaEnrollmentResponse converts a WorkloadSaaEnrollmentResponse object from its proto representation.
func ProtoToAssuredworkloadsWorkloadSaaEnrollmentResponse(p *assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponse) *assuredworkloads.WorkloadSaaEnrollmentResponse {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadSaaEnrollmentResponse{
		SetupStatus: ProtoToAssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum(p.GetSetupStatus()),
	}
	for _, r := range p.GetSetupErrors() {
		obj.SetupErrors = append(obj.SetupErrors, *ProtoToAssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum(r))
	}
	return obj
}

// ProtoToWorkloadComplianceStatus converts a WorkloadComplianceStatus object from its proto representation.
func ProtoToAssuredworkloadsWorkloadComplianceStatus(p *assuredworkloadspb.AssuredworkloadsWorkloadComplianceStatus) *assuredworkloads.WorkloadComplianceStatus {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadComplianceStatus{}
	for _, r := range p.GetActiveViolationCount() {
		obj.ActiveViolationCount = append(obj.ActiveViolationCount, r)
	}
	for _, r := range p.GetAcknowledgedViolationCount() {
		obj.AcknowledgedViolationCount = append(obj.AcknowledgedViolationCount, r)
	}
	return obj
}

// ProtoToWorkloadPartnerPermissions converts a WorkloadPartnerPermissions object from its proto representation.
func ProtoToAssuredworkloadsWorkloadPartnerPermissions(p *assuredworkloadspb.AssuredworkloadsWorkloadPartnerPermissions) *assuredworkloads.WorkloadPartnerPermissions {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadPartnerPermissions{
		DataLogsViewer:             dcl.Bool(p.GetDataLogsViewer()),
		ServiceAccessApprover:      dcl.Bool(p.GetServiceAccessApprover()),
		AssuredWorkloadsMonitoring: dcl.Bool(p.GetAssuredWorkloadsMonitoring()),
	}
	return obj
}

// ProtoToWorkloadEkmProvisioningResponse converts a WorkloadEkmProvisioningResponse object from its proto representation.
func ProtoToAssuredworkloadsWorkloadEkmProvisioningResponse(p *assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponse) *assuredworkloads.WorkloadEkmProvisioningResponse {
	if p == nil {
		return nil
	}
	obj := &assuredworkloads.WorkloadEkmProvisioningResponse{
		EkmProvisioningState:        ProtoToAssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(p.GetEkmProvisioningState()),
		EkmProvisioningErrorDomain:  ProtoToAssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(p.GetEkm_ProvisioningErrorDomain()),
		EkmProvisioningErrorMapping: ProtoToAssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(p.GetEkmProvisioningErrorMapping()),
	}
	return obj
}

// ProtoToWorkload converts a Workload resource from its proto representation.
func ProtoToWorkload(p *assuredworkloadspb.AssuredworkloadsWorkload) *assuredworkloads.Workload {
	obj := &assuredworkloads.Workload{
		Name:                          dcl.StringOrNil(p.GetName()),
		DisplayName:                   dcl.StringOrNil(p.GetDisplayName()),
		ComplianceRegime:              ProtoToAssuredworkloadsWorkloadComplianceRegimeEnum(p.GetComplianceRegime()),
		CreateTime:                    dcl.StringOrNil(p.GetCreateTime()),
		BillingAccount:                dcl.StringOrNil(p.GetBillingAccount()),
		ProvisionedResourcesParent:    dcl.StringOrNil(p.GetProvisionedResourcesParent()),
		KmsSettings:                   ProtoToAssuredworkloadsWorkloadKmsSettings(p.GetKmsSettings()),
		KajEnrollmentState:            ProtoToAssuredworkloadsWorkloadKajEnrollmentStateEnum(p.GetKajEnrollmentState()),
		EnableSovereignControls:       dcl.Bool(p.GetEnableSovereignControls()),
		SaaEnrollmentResponse:         ProtoToAssuredworkloadsWorkloadSaaEnrollmentResponse(p.GetSaaEnrollmentResponse()),
		ComplianceStatus:              ProtoToAssuredworkloadsWorkloadComplianceStatus(p.GetComplianceStatus()),
		Partner:                       ProtoToAssuredworkloadsWorkloadPartnerEnum(p.GetPartner()),
		PartnerPermissions:            ProtoToAssuredworkloadsWorkloadPartnerPermissions(p.GetPartnerPermissions()),
		EkmProvisioningResponse:       ProtoToAssuredworkloadsWorkloadEkmProvisioningResponse(p.GetEkmProvisioningResponse()),
		ViolationNotificationsEnabled: dcl.Bool(p.GetViolationNotificationsEnabled()),
		Organization:                  dcl.StringOrNil(p.GetOrganization()),
		Location:                      dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetResources() {
		obj.Resources = append(obj.Resources, *ProtoToAssuredworkloadsWorkloadResources(r))
	}
	for _, r := range p.GetResourceSettings() {
		obj.ResourceSettings = append(obj.ResourceSettings, *ProtoToAssuredworkloadsWorkloadResourceSettings(r))
	}
	for _, r := range p.GetCompliantButDisallowedServices() {
		obj.CompliantButDisallowedServices = append(obj.CompliantButDisallowedServices, r)
	}
	return obj
}

// WorkloadResourcesResourceTypeEnumToProto converts a WorkloadResourcesResourceTypeEnum enum to its proto representation.
func AssuredworkloadsWorkloadResourcesResourceTypeEnumToProto(e *assuredworkloads.WorkloadResourcesResourceTypeEnum) assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum_value["WorkloadResourcesResourceTypeEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadResourcesResourceTypeEnum(0)
}

// WorkloadComplianceRegimeEnumToProto converts a WorkloadComplianceRegimeEnum enum to its proto representation.
func AssuredworkloadsWorkloadComplianceRegimeEnumToProto(e *assuredworkloads.WorkloadComplianceRegimeEnum) assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum_value["WorkloadComplianceRegimeEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadComplianceRegimeEnum(0)
}

// WorkloadResourceSettingsResourceTypeEnumToProto converts a WorkloadResourceSettingsResourceTypeEnum enum to its proto representation.
func AssuredworkloadsWorkloadResourceSettingsResourceTypeEnumToProto(e *assuredworkloads.WorkloadResourceSettingsResourceTypeEnum) assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum_value["WorkloadResourceSettingsResourceTypeEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadResourceSettingsResourceTypeEnum(0)
}

// WorkloadKajEnrollmentStateEnumToProto converts a WorkloadKajEnrollmentStateEnum enum to its proto representation.
func AssuredworkloadsWorkloadKajEnrollmentStateEnumToProto(e *assuredworkloads.WorkloadKajEnrollmentStateEnum) assuredworkloadspb.AssuredworkloadsWorkloadKajEnrollmentStateEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadKajEnrollmentStateEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadKajEnrollmentStateEnum_value["WorkloadKajEnrollmentStateEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadKajEnrollmentStateEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadKajEnrollmentStateEnum(0)
}

// WorkloadSaaEnrollmentResponseSetupErrorsEnumToProto converts a WorkloadSaaEnrollmentResponseSetupErrorsEnum enum to its proto representation.
func AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnumToProto(e *assuredworkloads.WorkloadSaaEnrollmentResponseSetupErrorsEnum) assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum_value["WorkloadSaaEnrollmentResponseSetupErrorsEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum(0)
}

// WorkloadSaaEnrollmentResponseSetupStatusEnumToProto converts a WorkloadSaaEnrollmentResponseSetupStatusEnum enum to its proto representation.
func AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnumToProto(e *assuredworkloads.WorkloadSaaEnrollmentResponseSetupStatusEnum) assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum_value["WorkloadSaaEnrollmentResponseSetupStatusEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnum(0)
}

// WorkloadPartnerEnumToProto converts a WorkloadPartnerEnum enum to its proto representation.
func AssuredworkloadsWorkloadPartnerEnumToProto(e *assuredworkloads.WorkloadPartnerEnum) assuredworkloadspb.AssuredworkloadsWorkloadPartnerEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadPartnerEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadPartnerEnum_value["WorkloadPartnerEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadPartnerEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadPartnerEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningStateEnum enum to its proto representation.
func AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto(e *assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningStateEnum) assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningStateEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum enum to its proto representation.
func AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto(e *assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum) assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(0)
}

// WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto converts a WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum enum to its proto representation.
func AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto(e *assuredworkloads.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum) assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	if e == nil {
		return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(0)
	}
	if v, ok := assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum_value["WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum"+string(*e)]; ok {
		return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(v)
	}
	return assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(0)
}

// WorkloadResourcesToProto converts a WorkloadResources object to its proto representation.
func AssuredworkloadsWorkloadResourcesToProto(o *assuredworkloads.WorkloadResources) *assuredworkloadspb.AssuredworkloadsWorkloadResources {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadResources{}
	p.SetResourceId(dcl.ValueOrEmptyInt64(o.ResourceId))
	p.SetResourceType(AssuredworkloadsWorkloadResourcesResourceTypeEnumToProto(o.ResourceType))
	return p
}

// WorkloadKmsSettingsToProto converts a WorkloadKmsSettings object to its proto representation.
func AssuredworkloadsWorkloadKmsSettingsToProto(o *assuredworkloads.WorkloadKmsSettings) *assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadKmsSettings{}
	p.SetNextRotationTime(dcl.ValueOrEmptyString(o.NextRotationTime))
	p.SetRotationPeriod(dcl.ValueOrEmptyString(o.RotationPeriod))
	return p
}

// WorkloadResourceSettingsToProto converts a WorkloadResourceSettings object to its proto representation.
func AssuredworkloadsWorkloadResourceSettingsToProto(o *assuredworkloads.WorkloadResourceSettings) *assuredworkloadspb.AssuredworkloadsWorkloadResourceSettings {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadResourceSettings{}
	p.SetResourceId(dcl.ValueOrEmptyString(o.ResourceId))
	p.SetResourceType(AssuredworkloadsWorkloadResourceSettingsResourceTypeEnumToProto(o.ResourceType))
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	return p
}

// WorkloadSaaEnrollmentResponseToProto converts a WorkloadSaaEnrollmentResponse object to its proto representation.
func AssuredworkloadsWorkloadSaaEnrollmentResponseToProto(o *assuredworkloads.WorkloadSaaEnrollmentResponse) *assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponse {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponse{}
	p.SetSetupStatus(AssuredworkloadsWorkloadSaaEnrollmentResponseSetupStatusEnumToProto(o.SetupStatus))
	sSetupErrors := make([]assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum, len(o.SetupErrors))
	for i, r := range o.SetupErrors {
		sSetupErrors[i] = assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum(assuredworkloadspb.AssuredworkloadsWorkloadSaaEnrollmentResponseSetupErrorsEnum_value[string(r)])
	}
	p.SetSetupErrors(sSetupErrors)
	return p
}

// WorkloadComplianceStatusToProto converts a WorkloadComplianceStatus object to its proto representation.
func AssuredworkloadsWorkloadComplianceStatusToProto(o *assuredworkloads.WorkloadComplianceStatus) *assuredworkloadspb.AssuredworkloadsWorkloadComplianceStatus {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadComplianceStatus{}
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
func AssuredworkloadsWorkloadPartnerPermissionsToProto(o *assuredworkloads.WorkloadPartnerPermissions) *assuredworkloadspb.AssuredworkloadsWorkloadPartnerPermissions {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadPartnerPermissions{}
	p.SetDataLogsViewer(dcl.ValueOrEmptyBool(o.DataLogsViewer))
	p.SetServiceAccessApprover(dcl.ValueOrEmptyBool(o.ServiceAccessApprover))
	p.SetAssuredWorkloadsMonitoring(dcl.ValueOrEmptyBool(o.AssuredWorkloadsMonitoring))
	return p
}

// WorkloadEkmProvisioningResponseToProto converts a WorkloadEkmProvisioningResponse object to its proto representation.
func AssuredworkloadsWorkloadEkmProvisioningResponseToProto(o *assuredworkloads.WorkloadEkmProvisioningResponse) *assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponse {
	if o == nil {
		return nil
	}
	p := &assuredworkloadspb.AssuredworkloadsWorkloadEkmProvisioningResponse{}
	p.SetEkmProvisioningState(AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningStateEnumToProto(o.EkmProvisioningState))
	p.SetEkm_ProvisioningErrorDomain(AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumToProto(o.EkmProvisioningErrorDomain))
	p.SetEkmProvisioningErrorMapping(AssuredworkloadsWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumToProto(o.EkmProvisioningErrorMapping))
	return p
}

// WorkloadToProto converts a Workload resource to its proto representation.
func WorkloadToProto(resource *assuredworkloads.Workload) *assuredworkloadspb.AssuredworkloadsWorkload {
	p := &assuredworkloadspb.AssuredworkloadsWorkload{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetComplianceRegime(AssuredworkloadsWorkloadComplianceRegimeEnumToProto(resource.ComplianceRegime))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetBillingAccount(dcl.ValueOrEmptyString(resource.BillingAccount))
	p.SetProvisionedResourcesParent(dcl.ValueOrEmptyString(resource.ProvisionedResourcesParent))
	p.SetKmsSettings(AssuredworkloadsWorkloadKmsSettingsToProto(resource.KmsSettings))
	p.SetKajEnrollmentState(AssuredworkloadsWorkloadKajEnrollmentStateEnumToProto(resource.KajEnrollmentState))
	p.SetEnableSovereignControls(dcl.ValueOrEmptyBool(resource.EnableSovereignControls))
	p.SetSaaEnrollmentResponse(AssuredworkloadsWorkloadSaaEnrollmentResponseToProto(resource.SaaEnrollmentResponse))
	p.SetComplianceStatus(AssuredworkloadsWorkloadComplianceStatusToProto(resource.ComplianceStatus))
	p.SetPartner(AssuredworkloadsWorkloadPartnerEnumToProto(resource.Partner))
	p.SetPartnerPermissions(AssuredworkloadsWorkloadPartnerPermissionsToProto(resource.PartnerPermissions))
	p.SetEkmProvisioningResponse(AssuredworkloadsWorkloadEkmProvisioningResponseToProto(resource.EkmProvisioningResponse))
	p.SetViolationNotificationsEnabled(dcl.ValueOrEmptyBool(resource.ViolationNotificationsEnabled))
	p.SetOrganization(dcl.ValueOrEmptyString(resource.Organization))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sResources := make([]*assuredworkloadspb.AssuredworkloadsWorkloadResources, len(resource.Resources))
	for i, r := range resource.Resources {
		sResources[i] = AssuredworkloadsWorkloadResourcesToProto(&r)
	}
	p.SetResources(sResources)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sResourceSettings := make([]*assuredworkloadspb.AssuredworkloadsWorkloadResourceSettings, len(resource.ResourceSettings))
	for i, r := range resource.ResourceSettings {
		sResourceSettings[i] = AssuredworkloadsWorkloadResourceSettingsToProto(&r)
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
func (s *WorkloadServer) applyWorkload(ctx context.Context, c *assuredworkloads.Client, request *assuredworkloadspb.ApplyAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.AssuredworkloadsWorkload, error) {
	p := ProtoToWorkload(request.GetResource())
	res, err := c.ApplyWorkload(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadToProto(res)
	return r, nil
}

// applyAssuredworkloadsWorkload handles the gRPC request by passing it to the underlying Workload Apply() method.
func (s *WorkloadServer) ApplyAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.ApplyAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.AssuredworkloadsWorkload, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkload(ctx, cl, request)
}

// DeleteWorkload handles the gRPC request by passing it to the underlying Workload Delete() method.
func (s *WorkloadServer) DeleteAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.DeleteAssuredworkloadsWorkloadRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkload(ctx, ProtoToWorkload(request.GetResource()))

}

// ListAssuredworkloadsWorkload handles the gRPC request by passing it to the underlying WorkloadList() method.
func (s *WorkloadServer) ListAssuredworkloadsWorkload(ctx context.Context, request *assuredworkloadspb.ListAssuredworkloadsWorkloadRequest) (*assuredworkloadspb.ListAssuredworkloadsWorkloadResponse, error) {
	cl, err := createConfigWorkload(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkload(ctx, request.GetOrganization(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*assuredworkloadspb.AssuredworkloadsWorkload
	for _, r := range resources.Items {
		rp := WorkloadToProto(r)
		protos = append(protos, rp)
	}
	p := &assuredworkloadspb.ListAssuredworkloadsWorkloadResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkload(ctx context.Context, service_account_file string) (*assuredworkloads.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return assuredworkloads.NewClient(conf), nil
}
