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
package alpha

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Workload struct {
	Name                           *string                          `json:"name"`
	DisplayName                    *string                          `json:"displayName"`
	Resources                      []WorkloadResources              `json:"resources"`
	ComplianceRegime               *WorkloadComplianceRegimeEnum    `json:"complianceRegime"`
	CreateTime                     *string                          `json:"createTime"`
	BillingAccount                 *string                          `json:"billingAccount"`
	Labels                         map[string]string                `json:"labels"`
	ProvisionedResourcesParent     *string                          `json:"provisionedResourcesParent"`
	KmsSettings                    *WorkloadKmsSettings             `json:"kmsSettings"`
	ResourceSettings               []WorkloadResourceSettings       `json:"resourceSettings"`
	KajEnrollmentState             *WorkloadKajEnrollmentStateEnum  `json:"kajEnrollmentState"`
	EnableSovereignControls        *bool                            `json:"enableSovereignControls"`
	SaaEnrollmentResponse          *WorkloadSaaEnrollmentResponse   `json:"saaEnrollmentResponse"`
	ComplianceStatus               *WorkloadComplianceStatus        `json:"complianceStatus"`
	CompliantButDisallowedServices []string                         `json:"compliantButDisallowedServices"`
	Partner                        *WorkloadPartnerEnum             `json:"partner"`
	PartnerPermissions             *WorkloadPartnerPermissions      `json:"partnerPermissions"`
	EkmProvisioningResponse        *WorkloadEkmProvisioningResponse `json:"ekmProvisioningResponse"`
	ViolationNotificationsEnabled  *bool                            `json:"violationNotificationsEnabled"`
	Organization                   *string                          `json:"organization"`
	Location                       *string                          `json:"location"`
}

func (r *Workload) String() string {
	return dcl.SprintResource(r)
}

// The enum WorkloadResourcesResourceTypeEnum.
type WorkloadResourcesResourceTypeEnum string

// WorkloadResourcesResourceTypeEnumRef returns a *WorkloadResourcesResourceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadResourcesResourceTypeEnumRef(s string) *WorkloadResourcesResourceTypeEnum {
	v := WorkloadResourcesResourceTypeEnum(s)
	return &v
}

func (v WorkloadResourcesResourceTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"RESOURCE_TYPE_UNSPECIFIED", "CONSUMER_PROJECT", "ENCRYPTION_KEYS_PROJECT", "KEYRING", "CONSUMER_FOLDER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadResourcesResourceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadComplianceRegimeEnum.
type WorkloadComplianceRegimeEnum string

// WorkloadComplianceRegimeEnumRef returns a *WorkloadComplianceRegimeEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadComplianceRegimeEnumRef(s string) *WorkloadComplianceRegimeEnum {
	v := WorkloadComplianceRegimeEnum(s)
	return &v
}

func (v WorkloadComplianceRegimeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"COMPLIANCE_REGIME_UNSPECIFIED", "IL4", "CJIS", "FEDRAMP_HIGH", "FEDRAMP_MODERATE", "US_REGIONAL_ACCESS", "HIPAA", "HITRUST", "EU_REGIONS_AND_SUPPORT", "CA_REGIONS_AND_SUPPORT", "ITAR", "AU_REGIONS_AND_US_SUPPORT", "ASSURED_WORKLOADS_FOR_PARTNERS", "ISR_REGIONS", "ISR_REGIONS_AND_SUPPORT", "CA_PROTECTED_B", "IL5", "IL2", "JP_REGIONS_AND_SUPPORT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadComplianceRegimeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadResourceSettingsResourceTypeEnum.
type WorkloadResourceSettingsResourceTypeEnum string

// WorkloadResourceSettingsResourceTypeEnumRef returns a *WorkloadResourceSettingsResourceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadResourceSettingsResourceTypeEnumRef(s string) *WorkloadResourceSettingsResourceTypeEnum {
	v := WorkloadResourceSettingsResourceTypeEnum(s)
	return &v
}

func (v WorkloadResourceSettingsResourceTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"RESOURCE_TYPE_UNSPECIFIED", "CONSUMER_PROJECT", "ENCRYPTION_KEYS_PROJECT", "KEYRING", "CONSUMER_FOLDER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadResourceSettingsResourceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadKajEnrollmentStateEnum.
type WorkloadKajEnrollmentStateEnum string

// WorkloadKajEnrollmentStateEnumRef returns a *WorkloadKajEnrollmentStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadKajEnrollmentStateEnumRef(s string) *WorkloadKajEnrollmentStateEnum {
	v := WorkloadKajEnrollmentStateEnum(s)
	return &v
}

func (v WorkloadKajEnrollmentStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"KAJ_ENROLLMENT_STATE_UNSPECIFIED", "KAJ_ENROLLMENT_STATE_PENDING", "KAJ_ENROLLMENT_STATE_COMPLETE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadKajEnrollmentStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadSaaEnrollmentResponseSetupErrorsEnum.
type WorkloadSaaEnrollmentResponseSetupErrorsEnum string

// WorkloadSaaEnrollmentResponseSetupErrorsEnumRef returns a *WorkloadSaaEnrollmentResponseSetupErrorsEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadSaaEnrollmentResponseSetupErrorsEnumRef(s string) *WorkloadSaaEnrollmentResponseSetupErrorsEnum {
	v := WorkloadSaaEnrollmentResponseSetupErrorsEnum(s)
	return &v
}

func (v WorkloadSaaEnrollmentResponseSetupErrorsEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SETUP_ERROR_UNSPECIFIED", "ERROR_INVALID_BASE_SETUP", "ERROR_MISSING_EXTERNAL_SIGNING_KEY", "ERROR_NOT_ALL_SERVICES_ENROLLED", "ERROR_SETUP_CHECK_FAILED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadSaaEnrollmentResponseSetupErrorsEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadSaaEnrollmentResponseSetupStatusEnum.
type WorkloadSaaEnrollmentResponseSetupStatusEnum string

// WorkloadSaaEnrollmentResponseSetupStatusEnumRef returns a *WorkloadSaaEnrollmentResponseSetupStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadSaaEnrollmentResponseSetupStatusEnumRef(s string) *WorkloadSaaEnrollmentResponseSetupStatusEnum {
	v := WorkloadSaaEnrollmentResponseSetupStatusEnum(s)
	return &v
}

func (v WorkloadSaaEnrollmentResponseSetupStatusEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SETUP_STATE_UNSPECIFIED", "STATUS_PENDING", "STATUS_COMPLETE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadSaaEnrollmentResponseSetupStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadPartnerEnum.
type WorkloadPartnerEnum string

// WorkloadPartnerEnumRef returns a *WorkloadPartnerEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadPartnerEnumRef(s string) *WorkloadPartnerEnum {
	v := WorkloadPartnerEnum(s)
	return &v
}

func (v WorkloadPartnerEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"PARTNER_UNSPECIFIED", "LOCAL_CONTROLS_BY_S3NS", "SOVEREIGN_CONTROLS_BY_T_SYSTEMS", "SOVEREIGN_CONTROLS_BY_SIA_MINSAIT", "SOVEREIGN_CONTROLS_BY_PSN"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadPartnerEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadEkmProvisioningResponseEkmProvisioningStateEnum.
type WorkloadEkmProvisioningResponseEkmProvisioningStateEnum string

// WorkloadEkmProvisioningResponseEkmProvisioningStateEnumRef returns a *WorkloadEkmProvisioningResponseEkmProvisioningStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadEkmProvisioningResponseEkmProvisioningStateEnumRef(s string) *WorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	v := WorkloadEkmProvisioningResponseEkmProvisioningStateEnum(s)
	return &v
}

func (v WorkloadEkmProvisioningResponseEkmProvisioningStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EKM_PROVISIONING_STATE_UNSPECIFIED", "EKM_PROVISIONING_STATE_PENDING", "EKM_PROVISIONING_STATE_FAILED", "EKM_PROVISIONING_STATE_COMPLETED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadEkmProvisioningResponseEkmProvisioningStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum.
type WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum string

// WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumRef returns a *WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumRef(s string) *WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	v := WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(s)
	return &v
}

func (v WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EKM_PROVISIONING_ERROR_DOMAIN_UNSPECIFIED", "UNSPECIFIED_ERROR", "GOOGLE_SERVER_ERROR", "EXTERNAL_USER_ERROR", "EXTERNAL_PARTNER_ERROR", "TIMEOUT_ERROR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum.
type WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum string

// WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumRef returns a *WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum with the value of string s
// If the empty string is provided, nil is returned.
func WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumRef(s string) *WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	v := WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(s)
	return &v
}

func (v WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EKM_PROVISIONING_ERROR_MAPPING_UNSPECIFIED", "INVALID_SERVICE_ACCOUNT", "MISSING_METRICS_SCOPE_ADMIN_PERMISSION", "MISSING_EKM_CONNECTION_ADMIN_PERMISSION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type WorkloadResources struct {
	empty        bool                               `json:"-"`
	ResourceId   *int64                             `json:"resourceId"`
	ResourceType *WorkloadResourcesResourceTypeEnum `json:"resourceType"`
}

type jsonWorkloadResources WorkloadResources

func (r *WorkloadResources) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadResources
	} else {

		r.ResourceId = res.ResourceId

		r.ResourceType = res.ResourceType

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadResources *WorkloadResources = &WorkloadResources{empty: true}

func (r *WorkloadResources) Empty() bool {
	return r.empty
}

func (r *WorkloadResources) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadKmsSettings struct {
	empty            bool    `json:"-"`
	NextRotationTime *string `json:"nextRotationTime"`
	RotationPeriod   *string `json:"rotationPeriod"`
}

type jsonWorkloadKmsSettings WorkloadKmsSettings

func (r *WorkloadKmsSettings) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadKmsSettings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadKmsSettings
	} else {

		r.NextRotationTime = res.NextRotationTime

		r.RotationPeriod = res.RotationPeriod

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadKmsSettings is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadKmsSettings *WorkloadKmsSettings = &WorkloadKmsSettings{empty: true}

func (r *WorkloadKmsSettings) Empty() bool {
	return r.empty
}

func (r *WorkloadKmsSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadKmsSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadResourceSettings struct {
	empty        bool                                      `json:"-"`
	ResourceId   *string                                   `json:"resourceId"`
	ResourceType *WorkloadResourceSettingsResourceTypeEnum `json:"resourceType"`
	DisplayName  *string                                   `json:"displayName"`
}

type jsonWorkloadResourceSettings WorkloadResourceSettings

func (r *WorkloadResourceSettings) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadResourceSettings
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadResourceSettings
	} else {

		r.ResourceId = res.ResourceId

		r.ResourceType = res.ResourceType

		r.DisplayName = res.DisplayName

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadResourceSettings is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadResourceSettings *WorkloadResourceSettings = &WorkloadResourceSettings{empty: true}

func (r *WorkloadResourceSettings) Empty() bool {
	return r.empty
}

func (r *WorkloadResourceSettings) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadResourceSettings) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadSaaEnrollmentResponse struct {
	empty       bool                                           `json:"-"`
	SetupErrors []WorkloadSaaEnrollmentResponseSetupErrorsEnum `json:"setupErrors"`
	SetupStatus *WorkloadSaaEnrollmentResponseSetupStatusEnum  `json:"setupStatus"`
}

type jsonWorkloadSaaEnrollmentResponse WorkloadSaaEnrollmentResponse

func (r *WorkloadSaaEnrollmentResponse) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadSaaEnrollmentResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadSaaEnrollmentResponse
	} else {

		r.SetupErrors = res.SetupErrors

		r.SetupStatus = res.SetupStatus

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadSaaEnrollmentResponse is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadSaaEnrollmentResponse *WorkloadSaaEnrollmentResponse = &WorkloadSaaEnrollmentResponse{empty: true}

func (r *WorkloadSaaEnrollmentResponse) Empty() bool {
	return r.empty
}

func (r *WorkloadSaaEnrollmentResponse) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadSaaEnrollmentResponse) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadComplianceStatus struct {
	empty                      bool    `json:"-"`
	ActiveViolationCount       []int64 `json:"activeViolationCount"`
	AcknowledgedViolationCount []int64 `json:"acknowledgedViolationCount"`
}

type jsonWorkloadComplianceStatus WorkloadComplianceStatus

func (r *WorkloadComplianceStatus) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadComplianceStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadComplianceStatus
	} else {

		r.ActiveViolationCount = res.ActiveViolationCount

		r.AcknowledgedViolationCount = res.AcknowledgedViolationCount

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadComplianceStatus is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadComplianceStatus *WorkloadComplianceStatus = &WorkloadComplianceStatus{empty: true}

func (r *WorkloadComplianceStatus) Empty() bool {
	return r.empty
}

func (r *WorkloadComplianceStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadComplianceStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadPartnerPermissions struct {
	empty                      bool  `json:"-"`
	DataLogsViewer             *bool `json:"dataLogsViewer"`
	ServiceAccessApprover      *bool `json:"serviceAccessApprover"`
	AssuredWorkloadsMonitoring *bool `json:"assuredWorkloadsMonitoring"`
}

type jsonWorkloadPartnerPermissions WorkloadPartnerPermissions

func (r *WorkloadPartnerPermissions) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadPartnerPermissions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadPartnerPermissions
	} else {

		r.DataLogsViewer = res.DataLogsViewer

		r.ServiceAccessApprover = res.ServiceAccessApprover

		r.AssuredWorkloadsMonitoring = res.AssuredWorkloadsMonitoring

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadPartnerPermissions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadPartnerPermissions *WorkloadPartnerPermissions = &WorkloadPartnerPermissions{empty: true}

func (r *WorkloadPartnerPermissions) Empty() bool {
	return r.empty
}

func (r *WorkloadPartnerPermissions) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadPartnerPermissions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type WorkloadEkmProvisioningResponse struct {
	empty                       bool                                                            `json:"-"`
	EkmProvisioningState        *WorkloadEkmProvisioningResponseEkmProvisioningStateEnum        `json:"ekmProvisioningState"`
	EkmProvisioningErrorDomain  *WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum  `json:"ekmProvisioningErrorDomain"`
	EkmProvisioningErrorMapping *WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum `json:"ekmProvisioningErrorMapping"`
}

type jsonWorkloadEkmProvisioningResponse WorkloadEkmProvisioningResponse

func (r *WorkloadEkmProvisioningResponse) UnmarshalJSON(data []byte) error {
	var res jsonWorkloadEkmProvisioningResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyWorkloadEkmProvisioningResponse
	} else {

		r.EkmProvisioningState = res.EkmProvisioningState

		r.EkmProvisioningErrorDomain = res.EkmProvisioningErrorDomain

		r.EkmProvisioningErrorMapping = res.EkmProvisioningErrorMapping

	}
	return nil
}

// This object is used to assert a desired state where this WorkloadEkmProvisioningResponse is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyWorkloadEkmProvisioningResponse *WorkloadEkmProvisioningResponse = &WorkloadEkmProvisioningResponse{empty: true}

func (r *WorkloadEkmProvisioningResponse) Empty() bool {
	return r.empty
}

func (r *WorkloadEkmProvisioningResponse) String() string {
	return dcl.SprintResource(r)
}

func (r *WorkloadEkmProvisioningResponse) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Workload) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "assured_workloads",
		Type:    "Workload",
		Version: "alpha",
	}
}

func (r *Workload) ID() (string, error) {
	if err := extractWorkloadFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                              dcl.ValueOrEmptyString(nr.Name),
		"display_name":                      dcl.ValueOrEmptyString(nr.DisplayName),
		"resources":                         dcl.ValueOrEmptyString(nr.Resources),
		"compliance_regime":                 dcl.ValueOrEmptyString(nr.ComplianceRegime),
		"create_time":                       dcl.ValueOrEmptyString(nr.CreateTime),
		"billing_account":                   dcl.ValueOrEmptyString(nr.BillingAccount),
		"labels":                            dcl.ValueOrEmptyString(nr.Labels),
		"provisioned_resources_parent":      dcl.ValueOrEmptyString(nr.ProvisionedResourcesParent),
		"kms_settings":                      dcl.ValueOrEmptyString(nr.KmsSettings),
		"resource_settings":                 dcl.ValueOrEmptyString(nr.ResourceSettings),
		"kaj_enrollment_state":              dcl.ValueOrEmptyString(nr.KajEnrollmentState),
		"enable_sovereign_controls":         dcl.ValueOrEmptyString(nr.EnableSovereignControls),
		"saa_enrollment_response":           dcl.ValueOrEmptyString(nr.SaaEnrollmentResponse),
		"compliance_status":                 dcl.ValueOrEmptyString(nr.ComplianceStatus),
		"compliant_but_disallowed_services": dcl.ValueOrEmptyString(nr.CompliantButDisallowedServices),
		"partner":                           dcl.ValueOrEmptyString(nr.Partner),
		"partner_permissions":               dcl.ValueOrEmptyString(nr.PartnerPermissions),
		"ekm_provisioning_response":         dcl.ValueOrEmptyString(nr.EkmProvisioningResponse),
		"violation_notifications_enabled":   dcl.ValueOrEmptyString(nr.ViolationNotificationsEnabled),
		"organization":                      dcl.ValueOrEmptyString(nr.Organization),
		"location":                          dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("organizations/{{organization}}/locations/{{location}}/workloads/{{name}}", params), nil
}

const WorkloadMaxPage = -1

type WorkloadList struct {
	Items []*Workload

	nextToken string

	pageSize int32

	resource *Workload
}

func (l *WorkloadList) HasNext() bool {
	return l.nextToken != ""
}

func (l *WorkloadList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listWorkload(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListWorkload(ctx context.Context, organization, location string) (*WorkloadList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	c = NewClient(c.Config.Clone(dcl.WithCodeRetryability(map[int]dcl.Retryability{
		400: dcl.Retryability{
			Retryable: true,
			Pattern:   "contains projects or other resources that are not deleted",
			Timeout:   300000000000,
		},
	})))
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListWorkloadWithMaxResults(ctx, organization, location, WorkloadMaxPage)

}

func (c *Client) ListWorkloadWithMaxResults(ctx context.Context, organization, location string, pageSize int32) (*WorkloadList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Workload{
		Organization: &organization,
		Location:     &location,
	}
	items, token, err := c.listWorkload(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &WorkloadList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetWorkload(ctx context.Context, r *Workload) (*Workload, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	c = NewClient(c.Config.Clone(dcl.WithCodeRetryability(map[int]dcl.Retryability{
		400: dcl.Retryability{
			Retryable: true,
			Pattern:   "contains projects or other resources that are not deleted",
			Timeout:   300000000000,
		},
	})))
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractWorkloadFields(r)

	b, err := c.getWorkloadRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalWorkload(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Organization = r.Organization
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeWorkloadNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractWorkloadFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteWorkload(ctx context.Context, r *Workload) error {
	ctx = dcl.ContextWithRequestID(ctx)
	c = NewClient(c.Config.Clone(dcl.WithCodeRetryability(map[int]dcl.Retryability{
		400: dcl.Retryability{
			Retryable: true,
			Pattern:   "contains projects or other resources that are not deleted",
			Timeout:   300000000000,
		},
	})))
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Workload resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Workload...")
	deleteOp := deleteWorkloadOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllWorkload deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllWorkload(ctx context.Context, organization, location string, filter func(*Workload) bool) error {
	listObj, err := c.ListWorkload(ctx, organization, location)
	if err != nil {
		return err
	}

	err = c.deleteAllWorkload(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllWorkload(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyWorkload(ctx context.Context, rawDesired *Workload, opts ...dcl.ApplyOption) (*Workload, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	c = NewClient(c.Config.Clone(dcl.WithCodeRetryability(map[int]dcl.Retryability{
		400: dcl.Retryability{
			Retryable: true,
			Pattern:   "contains projects or other resources that are not deleted",
			Timeout:   300000000000,
		},
	})))
	var resultNewState *Workload
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyWorkloadHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyWorkloadHelper(c *Client, ctx context.Context, rawDesired *Workload, opts ...dcl.ApplyOption) (*Workload, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyWorkload...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractWorkloadFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.workloadDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToWorkloadDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []workloadApiOperation
	if create {
		ops = append(ops, &createWorkloadOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyWorkloadDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyWorkloadDiff(c *Client, ctx context.Context, desired *Workload, rawDesired *Workload, ops []workloadApiOperation, opts ...dcl.ApplyOption) (*Workload, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetWorkload(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createWorkloadOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapWorkload(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeWorkloadNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeWorkloadNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeWorkloadDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractWorkloadFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractWorkloadFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffWorkload(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
