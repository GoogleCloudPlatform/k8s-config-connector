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
package assuredworkloads

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Workload struct{}

func WorkloadToUnstructured(r *dclService.Workload) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "assuredworkloads",
			Version: "alpha",
			Type:    "Workload",
		},
		Object: make(map[string]interface{}),
	}
	if r.BillingAccount != nil {
		u.Object["billingAccount"] = *r.BillingAccount
	}
	if r.ComplianceRegime != nil {
		u.Object["complianceRegime"] = string(*r.ComplianceRegime)
	}
	if r.ComplianceStatus != nil && r.ComplianceStatus != dclService.EmptyWorkloadComplianceStatus {
		rComplianceStatus := make(map[string]interface{})
		var rComplianceStatusAcknowledgedViolationCount []interface{}
		for _, rComplianceStatusAcknowledgedViolationCountVal := range r.ComplianceStatus.AcknowledgedViolationCount {
			rComplianceStatusAcknowledgedViolationCount = append(rComplianceStatusAcknowledgedViolationCount, rComplianceStatusAcknowledgedViolationCountVal)
		}
		rComplianceStatus["acknowledgedViolationCount"] = rComplianceStatusAcknowledgedViolationCount
		var rComplianceStatusActiveViolationCount []interface{}
		for _, rComplianceStatusActiveViolationCountVal := range r.ComplianceStatus.ActiveViolationCount {
			rComplianceStatusActiveViolationCount = append(rComplianceStatusActiveViolationCount, rComplianceStatusActiveViolationCountVal)
		}
		rComplianceStatus["activeViolationCount"] = rComplianceStatusActiveViolationCount
		u.Object["complianceStatus"] = rComplianceStatus
	}
	var rCompliantButDisallowedServices []interface{}
	for _, rCompliantButDisallowedServicesVal := range r.CompliantButDisallowedServices {
		rCompliantButDisallowedServices = append(rCompliantButDisallowedServices, rCompliantButDisallowedServicesVal)
	}
	u.Object["compliantButDisallowedServices"] = rCompliantButDisallowedServices
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.EkmProvisioningResponse != nil && r.EkmProvisioningResponse != dclService.EmptyWorkloadEkmProvisioningResponse {
		rEkmProvisioningResponse := make(map[string]interface{})
		if r.EkmProvisioningResponse.EkmProvisioningErrorDomain != nil {
			rEkmProvisioningResponse["ekmProvisioningErrorDomain"] = string(*r.EkmProvisioningResponse.EkmProvisioningErrorDomain)
		}
		if r.EkmProvisioningResponse.EkmProvisioningErrorMapping != nil {
			rEkmProvisioningResponse["ekmProvisioningErrorMapping"] = string(*r.EkmProvisioningResponse.EkmProvisioningErrorMapping)
		}
		if r.EkmProvisioningResponse.EkmProvisioningState != nil {
			rEkmProvisioningResponse["ekmProvisioningState"] = string(*r.EkmProvisioningResponse.EkmProvisioningState)
		}
		u.Object["ekmProvisioningResponse"] = rEkmProvisioningResponse
	}
	if r.EnableSovereignControls != nil {
		u.Object["enableSovereignControls"] = *r.EnableSovereignControls
	}
	if r.KajEnrollmentState != nil {
		u.Object["kajEnrollmentState"] = string(*r.KajEnrollmentState)
	}
	if r.KmsSettings != nil && r.KmsSettings != dclService.EmptyWorkloadKmsSettings {
		rKmsSettings := make(map[string]interface{})
		if r.KmsSettings.NextRotationTime != nil {
			rKmsSettings["nextRotationTime"] = *r.KmsSettings.NextRotationTime
		}
		if r.KmsSettings.RotationPeriod != nil {
			rKmsSettings["rotationPeriod"] = *r.KmsSettings.RotationPeriod
		}
		u.Object["kmsSettings"] = rKmsSettings
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
	if r.Organization != nil {
		u.Object["organization"] = *r.Organization
	}
	if r.Partner != nil {
		u.Object["partner"] = string(*r.Partner)
	}
	if r.PartnerPermissions != nil && r.PartnerPermissions != dclService.EmptyWorkloadPartnerPermissions {
		rPartnerPermissions := make(map[string]interface{})
		if r.PartnerPermissions.AssuredWorkloadsMonitoring != nil {
			rPartnerPermissions["assuredWorkloadsMonitoring"] = *r.PartnerPermissions.AssuredWorkloadsMonitoring
		}
		if r.PartnerPermissions.DataLogsViewer != nil {
			rPartnerPermissions["dataLogsViewer"] = *r.PartnerPermissions.DataLogsViewer
		}
		if r.PartnerPermissions.ServiceAccessApprover != nil {
			rPartnerPermissions["serviceAccessApprover"] = *r.PartnerPermissions.ServiceAccessApprover
		}
		u.Object["partnerPermissions"] = rPartnerPermissions
	}
	if r.ProvisionedResourcesParent != nil {
		u.Object["provisionedResourcesParent"] = *r.ProvisionedResourcesParent
	}
	var rResourceSettings []interface{}
	for _, rResourceSettingsVal := range r.ResourceSettings {
		rResourceSettingsObject := make(map[string]interface{})
		if rResourceSettingsVal.DisplayName != nil {
			rResourceSettingsObject["displayName"] = *rResourceSettingsVal.DisplayName
		}
		if rResourceSettingsVal.ResourceId != nil {
			rResourceSettingsObject["resourceId"] = *rResourceSettingsVal.ResourceId
		}
		if rResourceSettingsVal.ResourceType != nil {
			rResourceSettingsObject["resourceType"] = string(*rResourceSettingsVal.ResourceType)
		}
		rResourceSettings = append(rResourceSettings, rResourceSettingsObject)
	}
	u.Object["resourceSettings"] = rResourceSettings
	var rResources []interface{}
	for _, rResourcesVal := range r.Resources {
		rResourcesObject := make(map[string]interface{})
		if rResourcesVal.ResourceId != nil {
			rResourcesObject["resourceId"] = *rResourcesVal.ResourceId
		}
		if rResourcesVal.ResourceType != nil {
			rResourcesObject["resourceType"] = string(*rResourcesVal.ResourceType)
		}
		rResources = append(rResources, rResourcesObject)
	}
	u.Object["resources"] = rResources
	if r.SaaEnrollmentResponse != nil && r.SaaEnrollmentResponse != dclService.EmptyWorkloadSaaEnrollmentResponse {
		rSaaEnrollmentResponse := make(map[string]interface{})
		var rSaaEnrollmentResponseSetupErrors []interface{}
		for _, rSaaEnrollmentResponseSetupErrorsVal := range r.SaaEnrollmentResponse.SetupErrors {
			rSaaEnrollmentResponseSetupErrors = append(rSaaEnrollmentResponseSetupErrors, string(rSaaEnrollmentResponseSetupErrorsVal))
		}
		rSaaEnrollmentResponse["setupErrors"] = rSaaEnrollmentResponseSetupErrors
		if r.SaaEnrollmentResponse.SetupStatus != nil {
			rSaaEnrollmentResponse["setupStatus"] = string(*r.SaaEnrollmentResponse.SetupStatus)
		}
		u.Object["saaEnrollmentResponse"] = rSaaEnrollmentResponse
	}
	if r.ViolationNotificationsEnabled != nil {
		u.Object["violationNotificationsEnabled"] = *r.ViolationNotificationsEnabled
	}
	return u
}

func UnstructuredToWorkload(u *unstructured.Resource) (*dclService.Workload, error) {
	r := &dclService.Workload{}
	if _, ok := u.Object["billingAccount"]; ok {
		if s, ok := u.Object["billingAccount"].(string); ok {
			r.BillingAccount = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.BillingAccount: expected string")
		}
	}
	if _, ok := u.Object["complianceRegime"]; ok {
		if s, ok := u.Object["complianceRegime"].(string); ok {
			r.ComplianceRegime = dclService.WorkloadComplianceRegimeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.ComplianceRegime: expected string")
		}
	}
	if _, ok := u.Object["complianceStatus"]; ok {
		if rComplianceStatus, ok := u.Object["complianceStatus"].(map[string]interface{}); ok {
			r.ComplianceStatus = &dclService.WorkloadComplianceStatus{}
			if _, ok := rComplianceStatus["acknowledgedViolationCount"]; ok {
				if s, ok := rComplianceStatus["acknowledgedViolationCount"].([]interface{}); ok {
					for _, ss := range s {
						if intval, ok := ss.(int64); ok {
							r.ComplianceStatus.AcknowledgedViolationCount = append(r.ComplianceStatus.AcknowledgedViolationCount, intval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ComplianceStatus.AcknowledgedViolationCount: expected []interface{}")
				}
			}
			if _, ok := rComplianceStatus["activeViolationCount"]; ok {
				if s, ok := rComplianceStatus["activeViolationCount"].([]interface{}); ok {
					for _, ss := range s {
						if intval, ok := ss.(int64); ok {
							r.ComplianceStatus.ActiveViolationCount = append(r.ComplianceStatus.ActiveViolationCount, intval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ComplianceStatus.ActiveViolationCount: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ComplianceStatus: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["compliantButDisallowedServices"]; ok {
		if s, ok := u.Object["compliantButDisallowedServices"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.CompliantButDisallowedServices = append(r.CompliantButDisallowedServices, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.CompliantButDisallowedServices: expected []interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["ekmProvisioningResponse"]; ok {
		if rEkmProvisioningResponse, ok := u.Object["ekmProvisioningResponse"].(map[string]interface{}); ok {
			r.EkmProvisioningResponse = &dclService.WorkloadEkmProvisioningResponse{}
			if _, ok := rEkmProvisioningResponse["ekmProvisioningErrorDomain"]; ok {
				if s, ok := rEkmProvisioningResponse["ekmProvisioningErrorDomain"].(string); ok {
					r.EkmProvisioningResponse.EkmProvisioningErrorDomain = dclService.WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.EkmProvisioningResponse.EkmProvisioningErrorDomain: expected string")
				}
			}
			if _, ok := rEkmProvisioningResponse["ekmProvisioningErrorMapping"]; ok {
				if s, ok := rEkmProvisioningResponse["ekmProvisioningErrorMapping"].(string); ok {
					r.EkmProvisioningResponse.EkmProvisioningErrorMapping = dclService.WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.EkmProvisioningResponse.EkmProvisioningErrorMapping: expected string")
				}
			}
			if _, ok := rEkmProvisioningResponse["ekmProvisioningState"]; ok {
				if s, ok := rEkmProvisioningResponse["ekmProvisioningState"].(string); ok {
					r.EkmProvisioningResponse.EkmProvisioningState = dclService.WorkloadEkmProvisioningResponseEkmProvisioningStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.EkmProvisioningResponse.EkmProvisioningState: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.EkmProvisioningResponse: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["enableSovereignControls"]; ok {
		if b, ok := u.Object["enableSovereignControls"].(bool); ok {
			r.EnableSovereignControls = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableSovereignControls: expected bool")
		}
	}
	if _, ok := u.Object["kajEnrollmentState"]; ok {
		if s, ok := u.Object["kajEnrollmentState"].(string); ok {
			r.KajEnrollmentState = dclService.WorkloadKajEnrollmentStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.KajEnrollmentState: expected string")
		}
	}
	if _, ok := u.Object["kmsSettings"]; ok {
		if rKmsSettings, ok := u.Object["kmsSettings"].(map[string]interface{}); ok {
			r.KmsSettings = &dclService.WorkloadKmsSettings{}
			if _, ok := rKmsSettings["nextRotationTime"]; ok {
				if s, ok := rKmsSettings["nextRotationTime"].(string); ok {
					r.KmsSettings.NextRotationTime = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.KmsSettings.NextRotationTime: expected string")
				}
			}
			if _, ok := rKmsSettings["rotationPeriod"]; ok {
				if s, ok := rKmsSettings["rotationPeriod"].(string); ok {
					r.KmsSettings.RotationPeriod = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.KmsSettings.RotationPeriod: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.KmsSettings: expected map[string]interface{}")
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
	if _, ok := u.Object["organization"]; ok {
		if s, ok := u.Object["organization"].(string); ok {
			r.Organization = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Organization: expected string")
		}
	}
	if _, ok := u.Object["partner"]; ok {
		if s, ok := u.Object["partner"].(string); ok {
			r.Partner = dclService.WorkloadPartnerEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Partner: expected string")
		}
	}
	if _, ok := u.Object["partnerPermissions"]; ok {
		if rPartnerPermissions, ok := u.Object["partnerPermissions"].(map[string]interface{}); ok {
			r.PartnerPermissions = &dclService.WorkloadPartnerPermissions{}
			if _, ok := rPartnerPermissions["assuredWorkloadsMonitoring"]; ok {
				if b, ok := rPartnerPermissions["assuredWorkloadsMonitoring"].(bool); ok {
					r.PartnerPermissions.AssuredWorkloadsMonitoring = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.PartnerPermissions.AssuredWorkloadsMonitoring: expected bool")
				}
			}
			if _, ok := rPartnerPermissions["dataLogsViewer"]; ok {
				if b, ok := rPartnerPermissions["dataLogsViewer"].(bool); ok {
					r.PartnerPermissions.DataLogsViewer = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.PartnerPermissions.DataLogsViewer: expected bool")
				}
			}
			if _, ok := rPartnerPermissions["serviceAccessApprover"]; ok {
				if b, ok := rPartnerPermissions["serviceAccessApprover"].(bool); ok {
					r.PartnerPermissions.ServiceAccessApprover = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.PartnerPermissions.ServiceAccessApprover: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.PartnerPermissions: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["provisionedResourcesParent"]; ok {
		if s, ok := u.Object["provisionedResourcesParent"].(string); ok {
			r.ProvisionedResourcesParent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ProvisionedResourcesParent: expected string")
		}
	}
	if _, ok := u.Object["resourceSettings"]; ok {
		if s, ok := u.Object["resourceSettings"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rResourceSettings dclService.WorkloadResourceSettings
					if _, ok := objval["displayName"]; ok {
						if s, ok := objval["displayName"].(string); ok {
							rResourceSettings.DisplayName = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rResourceSettings.DisplayName: expected string")
						}
					}
					if _, ok := objval["resourceId"]; ok {
						if s, ok := objval["resourceId"].(string); ok {
							rResourceSettings.ResourceId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rResourceSettings.ResourceId: expected string")
						}
					}
					if _, ok := objval["resourceType"]; ok {
						if s, ok := objval["resourceType"].(string); ok {
							rResourceSettings.ResourceType = dclService.WorkloadResourceSettingsResourceTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rResourceSettings.ResourceType: expected string")
						}
					}
					r.ResourceSettings = append(r.ResourceSettings, rResourceSettings)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResourceSettings: expected []interface{}")
		}
	}
	if _, ok := u.Object["resources"]; ok {
		if s, ok := u.Object["resources"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rResources dclService.WorkloadResources
					if _, ok := objval["resourceId"]; ok {
						if i, ok := objval["resourceId"].(int64); ok {
							rResources.ResourceId = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rResources.ResourceId: expected int64")
						}
					}
					if _, ok := objval["resourceType"]; ok {
						if s, ok := objval["resourceType"].(string); ok {
							rResources.ResourceType = dclService.WorkloadResourcesResourceTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rResources.ResourceType: expected string")
						}
					}
					r.Resources = append(r.Resources, rResources)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Resources: expected []interface{}")
		}
	}
	if _, ok := u.Object["saaEnrollmentResponse"]; ok {
		if rSaaEnrollmentResponse, ok := u.Object["saaEnrollmentResponse"].(map[string]interface{}); ok {
			r.SaaEnrollmentResponse = &dclService.WorkloadSaaEnrollmentResponse{}
			if _, ok := rSaaEnrollmentResponse["setupErrors"]; ok {
				if s, ok := rSaaEnrollmentResponse["setupErrors"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.SaaEnrollmentResponse.SetupErrors = append(r.SaaEnrollmentResponse.SetupErrors, dclService.WorkloadSaaEnrollmentResponseSetupErrorsEnum(strval))
						}
					}
				} else {
					return nil, fmt.Errorf("r.SaaEnrollmentResponse.SetupErrors: expected []interface{}")
				}
			}
			if _, ok := rSaaEnrollmentResponse["setupStatus"]; ok {
				if s, ok := rSaaEnrollmentResponse["setupStatus"].(string); ok {
					r.SaaEnrollmentResponse.SetupStatus = dclService.WorkloadSaaEnrollmentResponseSetupStatusEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.SaaEnrollmentResponse.SetupStatus: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SaaEnrollmentResponse: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["violationNotificationsEnabled"]; ok {
		if b, ok := u.Object["violationNotificationsEnabled"].(bool); ok {
			r.ViolationNotificationsEnabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.ViolationNotificationsEnabled: expected bool")
		}
	}
	return r, nil
}

func GetWorkload(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetWorkload(ctx, r)
	if err != nil {
		return nil, err
	}
	return WorkloadToUnstructured(r), nil
}

func ListWorkload(ctx context.Context, config *dcl.Config, organization string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListWorkload(ctx, organization, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, WorkloadToUnstructured(r))
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

func ApplyWorkload(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkload(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyWorkload(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return WorkloadToUnstructured(r), nil
}

func WorkloadHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToWorkload(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyWorkload(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteWorkload(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return err
	}
	return c.DeleteWorkload(ctx, r)
}

func WorkloadID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToWorkload(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Workload) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"assuredworkloads",
		"Workload",
		"alpha",
	}
}

func (r *Workload) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Workload) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Workload) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetWorkload(ctx, config, resource)
}

func (r *Workload) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyWorkload(ctx, config, resource, opts...)
}

func (r *Workload) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return WorkloadHasDiff(ctx, config, resource, opts...)
}

func (r *Workload) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteWorkload(ctx, config, resource)
}

func (r *Workload) ID(resource *unstructured.Resource) (string, error) {
	return WorkloadID(resource)
}

func init() {
	unstructured.Register(&Workload{})
}
