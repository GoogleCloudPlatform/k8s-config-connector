// Copyright 2023 Google LLC. All Rights Reserved.
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
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Job struct {
	Name                     *string                      `json:"name"`
	Uid                      *string                      `json:"uid"`
	Generation               *int64                       `json:"generation"`
	Labels                   map[string]string            `json:"labels"`
	Annotations              map[string]string            `json:"annotations"`
	CreateTime               *string                      `json:"createTime"`
	UpdateTime               *string                      `json:"updateTime"`
	DeleteTime               *string                      `json:"deleteTime"`
	ExpireTime               *string                      `json:"expireTime"`
	Creator                  *string                      `json:"creator"`
	LastModifier             *string                      `json:"lastModifier"`
	Client                   *string                      `json:"client"`
	ClientVersion            *string                      `json:"clientVersion"`
	LaunchStage              *JobLaunchStageEnum          `json:"launchStage"`
	BinaryAuthorization      *JobBinaryAuthorization      `json:"binaryAuthorization"`
	Template                 *JobTemplate                 `json:"template"`
	ObservedGeneration       *int64                       `json:"observedGeneration"`
	TerminalCondition        *JobTerminalCondition        `json:"terminalCondition"`
	Conditions               []JobConditions              `json:"conditions"`
	ExecutionCount           *int64                       `json:"executionCount"`
	LatestSucceededExecution *JobLatestSucceededExecution `json:"latestSucceededExecution"`
	LatestCreatedExecution   *JobLatestCreatedExecution   `json:"latestCreatedExecution"`
	Reconciling              *bool                        `json:"reconciling"`
	Etag                     *string                      `json:"etag"`
	Project                  *string                      `json:"project"`
	Location                 *string                      `json:"location"`
}

func (r *Job) String() string {
	return dcl.SprintResource(r)
}

// The enum JobLaunchStageEnum.
type JobLaunchStageEnum string

// JobLaunchStageEnumRef returns a *JobLaunchStageEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobLaunchStageEnumRef(s string) *JobLaunchStageEnum {
	v := JobLaunchStageEnum(s)
	return &v
}

func (v JobLaunchStageEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"LAUNCH_STAGE_UNSPECIFIED", "UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobLaunchStageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTemplateTemplateExecutionEnvironmentEnum.
type JobTemplateTemplateExecutionEnvironmentEnum string

// JobTemplateTemplateExecutionEnvironmentEnumRef returns a *JobTemplateTemplateExecutionEnvironmentEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTemplateTemplateExecutionEnvironmentEnumRef(s string) *JobTemplateTemplateExecutionEnvironmentEnum {
	v := JobTemplateTemplateExecutionEnvironmentEnum(s)
	return &v
}

func (v JobTemplateTemplateExecutionEnvironmentEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EXECUTION_ENVIRONMENT_UNSPECIFIED", "EXECUTION_ENVIRONMENT_DEFAULT", "EXECUTION_ENVIRONMENT_GEN2"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTemplateTemplateExecutionEnvironmentEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTemplateTemplateVPCAccessEgressEnum.
type JobTemplateTemplateVPCAccessEgressEnum string

// JobTemplateTemplateVPCAccessEgressEnumRef returns a *JobTemplateTemplateVPCAccessEgressEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTemplateTemplateVPCAccessEgressEnumRef(s string) *JobTemplateTemplateVPCAccessEgressEnum {
	v := JobTemplateTemplateVPCAccessEgressEnum(s)
	return &v
}

func (v JobTemplateTemplateVPCAccessEgressEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"VPC_EGRESS_UNSPECIFIED", "ALL_TRAFFIC", "PRIVATE_RANGES_ONLY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTemplateTemplateVPCAccessEgressEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTerminalConditionStateEnum.
type JobTerminalConditionStateEnum string

// JobTerminalConditionStateEnumRef returns a *JobTerminalConditionStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTerminalConditionStateEnumRef(s string) *JobTerminalConditionStateEnum {
	v := JobTerminalConditionStateEnum(s)
	return &v
}

func (v JobTerminalConditionStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "CONDITION_PENDING", "CONDITION_RECONCILING", "CONDITION_FAILED", "CONDITION_SUCCEEDED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTerminalConditionStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTerminalConditionSeverityEnum.
type JobTerminalConditionSeverityEnum string

// JobTerminalConditionSeverityEnumRef returns a *JobTerminalConditionSeverityEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTerminalConditionSeverityEnumRef(s string) *JobTerminalConditionSeverityEnum {
	v := JobTerminalConditionSeverityEnum(s)
	return &v
}

func (v JobTerminalConditionSeverityEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SEVERITY_UNSPECIFIED", "ERROR", "WARNING", "INFO"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTerminalConditionSeverityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTerminalConditionReasonEnum.
type JobTerminalConditionReasonEnum string

// JobTerminalConditionReasonEnumRef returns a *JobTerminalConditionReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTerminalConditionReasonEnumRef(s string) *JobTerminalConditionReasonEnum {
	v := JobTerminalConditionReasonEnum(s)
	return &v
}

func (v JobTerminalConditionReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"COMMON_REASON_UNDEFINED", "UNKNOWN", "ROUTE_MISSING", "REVISION_FAILED", "PROGRESS_DEADLINE_EXCEEDED", "CONTAINER_MISSING", "CONTAINER_PERMISSION_DENIED", "CONTAINER_IMAGE_UNAUTHORIZED", "CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED", "ENCRYPTION_KEY_PERMISSION_DENIED", "ENCRYPTION_KEY_CHECK_FAILED", "SECRETS_ACCESS_CHECK_FAILED", "WAITING_FOR_OPERATION", "IMMEDIATE_RETRY", "POSTPONED_RETRY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTerminalConditionReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTerminalConditionInternalReasonEnum.
type JobTerminalConditionInternalReasonEnum string

// JobTerminalConditionInternalReasonEnumRef returns a *JobTerminalConditionInternalReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTerminalConditionInternalReasonEnumRef(s string) *JobTerminalConditionInternalReasonEnum {
	v := JobTerminalConditionInternalReasonEnum(s)
	return &v
}

func (v JobTerminalConditionInternalReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INTERNAL_REASON_UNDEFINED", "CONFLICTING_REVISION_NAME", "REVISION_MISSING", "CONFIGURATION_MISSING", "ASSIGNING_TRAFFIC", "UPDATING_INGRESS_TRAFFIC_ALLOWED", "REVISION_ORG_POLICY_VIOLATION", "ENABLING_GCFV2_URI_SUPPORT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTerminalConditionInternalReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTerminalConditionDomainMappingReasonEnum.
type JobTerminalConditionDomainMappingReasonEnum string

// JobTerminalConditionDomainMappingReasonEnumRef returns a *JobTerminalConditionDomainMappingReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTerminalConditionDomainMappingReasonEnumRef(s string) *JobTerminalConditionDomainMappingReasonEnum {
	v := JobTerminalConditionDomainMappingReasonEnum(s)
	return &v
}

func (v JobTerminalConditionDomainMappingReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"DOMAIN_MAPPING_REASON_UNDEFINED", "ROUTE_NOT_READY", "PERMISSION_DENIED", "CERTIFICATE_ALREADY_EXISTS", "MAPPING_ALREADY_EXISTS", "CERTIFICATE_PENDING", "CERTIFICATE_FAILED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTerminalConditionDomainMappingReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTerminalConditionRevisionReasonEnum.
type JobTerminalConditionRevisionReasonEnum string

// JobTerminalConditionRevisionReasonEnumRef returns a *JobTerminalConditionRevisionReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTerminalConditionRevisionReasonEnumRef(s string) *JobTerminalConditionRevisionReasonEnum {
	v := JobTerminalConditionRevisionReasonEnum(s)
	return &v
}

func (v JobTerminalConditionRevisionReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"REVISION_REASON_UNDEFINED", "PENDING", "RESERVE", "RETIRED", "RETIRING", "RECREATING", "HEALTH_CHECK_CONTAINER_ERROR", "CUSTOMIZED_PATH_RESPONSE_PENDING", "MIN_INSTANCES_NOT_PROVISIONED", "ACTIVE_REVISION_LIMIT_REACHED", "NO_DEPLOYMENT", "HEALTH_CHECK_SKIPPED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTerminalConditionRevisionReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobTerminalConditionExecutionReasonEnum.
type JobTerminalConditionExecutionReasonEnum string

// JobTerminalConditionExecutionReasonEnumRef returns a *JobTerminalConditionExecutionReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobTerminalConditionExecutionReasonEnumRef(s string) *JobTerminalConditionExecutionReasonEnum {
	v := JobTerminalConditionExecutionReasonEnum(s)
	return &v
}

func (v JobTerminalConditionExecutionReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EXECUTION_REASON_UNDEFINED", "JOB_STATUS_SERVICE_POLLING_ERROR", "NON_ZERO_EXIT_CODE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobTerminalConditionExecutionReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobConditionsStateEnum.
type JobConditionsStateEnum string

// JobConditionsStateEnumRef returns a *JobConditionsStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobConditionsStateEnumRef(s string) *JobConditionsStateEnum {
	v := JobConditionsStateEnum(s)
	return &v
}

func (v JobConditionsStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATE_UNSPECIFIED", "CONDITION_PENDING", "CONDITION_RECONCILING", "CONDITION_FAILED", "CONDITION_SUCCEEDED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobConditionsStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobConditionsSeverityEnum.
type JobConditionsSeverityEnum string

// JobConditionsSeverityEnumRef returns a *JobConditionsSeverityEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobConditionsSeverityEnumRef(s string) *JobConditionsSeverityEnum {
	v := JobConditionsSeverityEnum(s)
	return &v
}

func (v JobConditionsSeverityEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"SEVERITY_UNSPECIFIED", "ERROR", "WARNING", "INFO"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobConditionsSeverityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobConditionsReasonEnum.
type JobConditionsReasonEnum string

// JobConditionsReasonEnumRef returns a *JobConditionsReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobConditionsReasonEnumRef(s string) *JobConditionsReasonEnum {
	v := JobConditionsReasonEnum(s)
	return &v
}

func (v JobConditionsReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"COMMON_REASON_UNDEFINED", "UNKNOWN", "REVISION_FAILED", "PROGRESS_DEADLINE_EXCEEDED", "BUILD_STEP_FAILED", "CONTAINER_MISSING", "CONTAINER_PERMISSION_DENIED", "CONTAINER_IMAGE_UNAUTHORIZED", "CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED", "ENCRYPTION_KEY_PERMISSION_DENIED", "ENCRYPTION_KEY_CHECK_FAILED", "SECRETS_ACCESS_CHECK_FAILED", "WAITING_FOR_OPERATION", "IMMEDIATE_RETRY", "POSTPONED_RETRY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobConditionsReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobConditionsRevisionReasonEnum.
type JobConditionsRevisionReasonEnum string

// JobConditionsRevisionReasonEnumRef returns a *JobConditionsRevisionReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobConditionsRevisionReasonEnumRef(s string) *JobConditionsRevisionReasonEnum {
	v := JobConditionsRevisionReasonEnum(s)
	return &v
}

func (v JobConditionsRevisionReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"REVISION_REASON_UNDEFINED", "PENDING", "RESERVE", "RETIRED", "RETIRING", "RECREATING", "HEALTH_CHECK_CONTAINER_ERROR", "CUSTOMIZED_PATH_RESPONSE_PENDING", "MIN_INSTANCES_NOT_PROVISIONED", "ACTIVE_REVISION_LIMIT_REACHED", "NO_DEPLOYMENT", "HEALTH_CHECK_SKIPPED"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobConditionsRevisionReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum JobConditionsExecutionReasonEnum.
type JobConditionsExecutionReasonEnum string

// JobConditionsExecutionReasonEnumRef returns a *JobConditionsExecutionReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func JobConditionsExecutionReasonEnumRef(s string) *JobConditionsExecutionReasonEnum {
	v := JobConditionsExecutionReasonEnum(s)
	return &v
}

func (v JobConditionsExecutionReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EXECUTION_REASON_UNDEFINED", "JOB_STATUS_SERVICE_POLLING_ERROR", "NON_ZERO_EXIT_CODE"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "JobConditionsExecutionReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type JobBinaryAuthorization struct {
	empty                   bool    `json:"-"`
	UseDefault              *bool   `json:"useDefault"`
	BreakglassJustification *string `json:"breakglassJustification"`
}

type jsonJobBinaryAuthorization JobBinaryAuthorization

func (r *JobBinaryAuthorization) UnmarshalJSON(data []byte) error {
	var res jsonJobBinaryAuthorization
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobBinaryAuthorization
	} else {

		r.UseDefault = res.UseDefault

		r.BreakglassJustification = res.BreakglassJustification

	}
	return nil
}

// This object is used to assert a desired state where this JobBinaryAuthorization is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobBinaryAuthorization *JobBinaryAuthorization = &JobBinaryAuthorization{empty: true}

func (r *JobBinaryAuthorization) Empty() bool {
	return r.empty
}

func (r *JobBinaryAuthorization) String() string {
	return dcl.SprintResource(r)
}

func (r *JobBinaryAuthorization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplate struct {
	empty       bool                 `json:"-"`
	Labels      map[string]string    `json:"labels"`
	Annotations map[string]string    `json:"annotations"`
	Parallelism *int64               `json:"parallelism"`
	TaskCount   *int64               `json:"taskCount"`
	Template    *JobTemplateTemplate `json:"template"`
}

type jsonJobTemplate JobTemplate

func (r *JobTemplate) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplate
	} else {

		r.Labels = res.Labels

		r.Annotations = res.Annotations

		r.Parallelism = res.Parallelism

		r.TaskCount = res.TaskCount

		r.Template = res.Template

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplate *JobTemplate = &JobTemplate{empty: true}

func (r *JobTemplate) Empty() bool {
	return r.empty
}

func (r *JobTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplate struct {
	empty                bool                                         `json:"-"`
	Containers           []JobTemplateTemplateContainers              `json:"containers"`
	Volumes              []JobTemplateTemplateVolumes                 `json:"volumes"`
	MaxRetries           *int64                                       `json:"maxRetries"`
	Timeout              *string                                      `json:"timeout"`
	ServiceAccount       *string                                      `json:"serviceAccount"`
	ExecutionEnvironment *JobTemplateTemplateExecutionEnvironmentEnum `json:"executionEnvironment"`
	EncryptionKey        *string                                      `json:"encryptionKey"`
	VPCAccess            *JobTemplateTemplateVPCAccess                `json:"vpcAccess"`
}

type jsonJobTemplateTemplate JobTemplateTemplate

func (r *JobTemplateTemplate) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplate
	} else {

		r.Containers = res.Containers

		r.Volumes = res.Volumes

		r.MaxRetries = res.MaxRetries

		r.Timeout = res.Timeout

		r.ServiceAccount = res.ServiceAccount

		r.ExecutionEnvironment = res.ExecutionEnvironment

		r.EncryptionKey = res.EncryptionKey

		r.VPCAccess = res.VPCAccess

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplate *JobTemplateTemplate = &JobTemplateTemplate{empty: true}

func (r *JobTemplateTemplate) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateContainers struct {
	empty        bool                                        `json:"-"`
	Name         *string                                     `json:"name"`
	Image        *string                                     `json:"image"`
	Command      []string                                    `json:"command"`
	Args         []string                                    `json:"args"`
	Env          []JobTemplateTemplateContainersEnv          `json:"env"`
	Resources    *JobTemplateTemplateContainersResources     `json:"resources"`
	Ports        []JobTemplateTemplateContainersPorts        `json:"ports"`
	VolumeMounts []JobTemplateTemplateContainersVolumeMounts `json:"volumeMounts"`
}

type jsonJobTemplateTemplateContainers JobTemplateTemplateContainers

func (r *JobTemplateTemplateContainers) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateContainers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateContainers
	} else {

		r.Name = res.Name

		r.Image = res.Image

		r.Command = res.Command

		r.Args = res.Args

		r.Env = res.Env

		r.Resources = res.Resources

		r.Ports = res.Ports

		r.VolumeMounts = res.VolumeMounts

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateContainers is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateContainers *JobTemplateTemplateContainers = &JobTemplateTemplateContainers{empty: true}

func (r *JobTemplateTemplateContainers) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateContainers) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateContainers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateContainersEnv struct {
	empty       bool                                         `json:"-"`
	Name        *string                                      `json:"name"`
	Value       *string                                      `json:"value"`
	ValueSource *JobTemplateTemplateContainersEnvValueSource `json:"valueSource"`
}

type jsonJobTemplateTemplateContainersEnv JobTemplateTemplateContainersEnv

func (r *JobTemplateTemplateContainersEnv) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateContainersEnv
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateContainersEnv
	} else {

		r.Name = res.Name

		r.Value = res.Value

		r.ValueSource = res.ValueSource

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateContainersEnv is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateContainersEnv *JobTemplateTemplateContainersEnv = &JobTemplateTemplateContainersEnv{empty: true}

func (r *JobTemplateTemplateContainersEnv) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateContainersEnv) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateContainersEnv) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateContainersEnvValueSource struct {
	empty        bool                                                     `json:"-"`
	SecretKeyRef *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef `json:"secretKeyRef"`
}

type jsonJobTemplateTemplateContainersEnvValueSource JobTemplateTemplateContainersEnvValueSource

func (r *JobTemplateTemplateContainersEnvValueSource) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateContainersEnvValueSource
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateContainersEnvValueSource
	} else {

		r.SecretKeyRef = res.SecretKeyRef

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateContainersEnvValueSource is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateContainersEnvValueSource *JobTemplateTemplateContainersEnvValueSource = &JobTemplateTemplateContainersEnvValueSource{empty: true}

func (r *JobTemplateTemplateContainersEnvValueSource) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateContainersEnvValueSource) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateContainersEnvValueSource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateContainersEnvValueSourceSecretKeyRef struct {
	empty   bool    `json:"-"`
	Secret  *string `json:"secret"`
	Version *string `json:"version"`
}

type jsonJobTemplateTemplateContainersEnvValueSourceSecretKeyRef JobTemplateTemplateContainersEnvValueSourceSecretKeyRef

func (r *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateContainersEnvValueSourceSecretKeyRef
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateContainersEnvValueSourceSecretKeyRef
	} else {

		r.Secret = res.Secret

		r.Version = res.Version

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateContainersEnvValueSourceSecretKeyRef is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateContainersEnvValueSourceSecretKeyRef *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef = &JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{empty: true}

func (r *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateContainersResources struct {
	empty   bool              `json:"-"`
	Limits  map[string]string `json:"limits"`
	CpuIdle *bool             `json:"cpuIdle"`
}

type jsonJobTemplateTemplateContainersResources JobTemplateTemplateContainersResources

func (r *JobTemplateTemplateContainersResources) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateContainersResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateContainersResources
	} else {

		r.Limits = res.Limits

		r.CpuIdle = res.CpuIdle

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateContainersResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateContainersResources *JobTemplateTemplateContainersResources = &JobTemplateTemplateContainersResources{empty: true}

func (r *JobTemplateTemplateContainersResources) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateContainersResources) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateContainersResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateContainersPorts struct {
	empty         bool    `json:"-"`
	Name          *string `json:"name"`
	ContainerPort *int64  `json:"containerPort"`
}

type jsonJobTemplateTemplateContainersPorts JobTemplateTemplateContainersPorts

func (r *JobTemplateTemplateContainersPorts) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateContainersPorts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateContainersPorts
	} else {

		r.Name = res.Name

		r.ContainerPort = res.ContainerPort

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateContainersPorts is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateContainersPorts *JobTemplateTemplateContainersPorts = &JobTemplateTemplateContainersPorts{empty: true}

func (r *JobTemplateTemplateContainersPorts) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateContainersPorts) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateContainersPorts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateContainersVolumeMounts struct {
	empty     bool    `json:"-"`
	Name      *string `json:"name"`
	MountPath *string `json:"mountPath"`
}

type jsonJobTemplateTemplateContainersVolumeMounts JobTemplateTemplateContainersVolumeMounts

func (r *JobTemplateTemplateContainersVolumeMounts) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateContainersVolumeMounts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateContainersVolumeMounts
	} else {

		r.Name = res.Name

		r.MountPath = res.MountPath

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateContainersVolumeMounts is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateContainersVolumeMounts *JobTemplateTemplateContainersVolumeMounts = &JobTemplateTemplateContainersVolumeMounts{empty: true}

func (r *JobTemplateTemplateContainersVolumeMounts) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateContainersVolumeMounts) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateContainersVolumeMounts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateVolumes struct {
	empty            bool                                        `json:"-"`
	Name             *string                                     `json:"name"`
	Secret           *JobTemplateTemplateVolumesSecret           `json:"secret"`
	CloudSqlInstance *JobTemplateTemplateVolumesCloudSqlInstance `json:"cloudSqlInstance"`
}

type jsonJobTemplateTemplateVolumes JobTemplateTemplateVolumes

func (r *JobTemplateTemplateVolumes) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateVolumes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateVolumes
	} else {

		r.Name = res.Name

		r.Secret = res.Secret

		r.CloudSqlInstance = res.CloudSqlInstance

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateVolumes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateVolumes *JobTemplateTemplateVolumes = &JobTemplateTemplateVolumes{empty: true}

func (r *JobTemplateTemplateVolumes) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateVolumes) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateVolumes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateVolumesSecret struct {
	empty       bool                                    `json:"-"`
	Secret      *string                                 `json:"secret"`
	Items       []JobTemplateTemplateVolumesSecretItems `json:"items"`
	DefaultMode *int64                                  `json:"defaultMode"`
}

type jsonJobTemplateTemplateVolumesSecret JobTemplateTemplateVolumesSecret

func (r *JobTemplateTemplateVolumesSecret) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateVolumesSecret
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateVolumesSecret
	} else {

		r.Secret = res.Secret

		r.Items = res.Items

		r.DefaultMode = res.DefaultMode

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateVolumesSecret is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateVolumesSecret *JobTemplateTemplateVolumesSecret = &JobTemplateTemplateVolumesSecret{empty: true}

func (r *JobTemplateTemplateVolumesSecret) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateVolumesSecret) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateVolumesSecret) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateVolumesSecretItems struct {
	empty   bool    `json:"-"`
	Path    *string `json:"path"`
	Version *string `json:"version"`
	Mode    *int64  `json:"mode"`
}

type jsonJobTemplateTemplateVolumesSecretItems JobTemplateTemplateVolumesSecretItems

func (r *JobTemplateTemplateVolumesSecretItems) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateVolumesSecretItems
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateVolumesSecretItems
	} else {

		r.Path = res.Path

		r.Version = res.Version

		r.Mode = res.Mode

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateVolumesSecretItems is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateVolumesSecretItems *JobTemplateTemplateVolumesSecretItems = &JobTemplateTemplateVolumesSecretItems{empty: true}

func (r *JobTemplateTemplateVolumesSecretItems) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateVolumesSecretItems) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateVolumesSecretItems) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateVolumesCloudSqlInstance struct {
	empty     bool     `json:"-"`
	Instances []string `json:"instances"`
}

type jsonJobTemplateTemplateVolumesCloudSqlInstance JobTemplateTemplateVolumesCloudSqlInstance

func (r *JobTemplateTemplateVolumesCloudSqlInstance) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateVolumesCloudSqlInstance
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateVolumesCloudSqlInstance
	} else {

		r.Instances = res.Instances

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateVolumesCloudSqlInstance is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateVolumesCloudSqlInstance *JobTemplateTemplateVolumesCloudSqlInstance = &JobTemplateTemplateVolumesCloudSqlInstance{empty: true}

func (r *JobTemplateTemplateVolumesCloudSqlInstance) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateVolumesCloudSqlInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateVolumesCloudSqlInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTemplateTemplateVPCAccess struct {
	empty     bool                                    `json:"-"`
	Connector *string                                 `json:"connector"`
	Egress    *JobTemplateTemplateVPCAccessEgressEnum `json:"egress"`
}

type jsonJobTemplateTemplateVPCAccess JobTemplateTemplateVPCAccess

func (r *JobTemplateTemplateVPCAccess) UnmarshalJSON(data []byte) error {
	var res jsonJobTemplateTemplateVPCAccess
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTemplateTemplateVPCAccess
	} else {

		r.Connector = res.Connector

		r.Egress = res.Egress

	}
	return nil
}

// This object is used to assert a desired state where this JobTemplateTemplateVPCAccess is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTemplateTemplateVPCAccess *JobTemplateTemplateVPCAccess = &JobTemplateTemplateVPCAccess{empty: true}

func (r *JobTemplateTemplateVPCAccess) Empty() bool {
	return r.empty
}

func (r *JobTemplateTemplateVPCAccess) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTemplateTemplateVPCAccess) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobTerminalCondition struct {
	empty               bool                                         `json:"-"`
	Type                *string                                      `json:"type"`
	State               *JobTerminalConditionStateEnum               `json:"state"`
	Message             *string                                      `json:"message"`
	LastTransitionTime  *string                                      `json:"lastTransitionTime"`
	Severity            *JobTerminalConditionSeverityEnum            `json:"severity"`
	Reason              *JobTerminalConditionReasonEnum              `json:"reason"`
	InternalReason      *JobTerminalConditionInternalReasonEnum      `json:"internalReason"`
	DomainMappingReason *JobTerminalConditionDomainMappingReasonEnum `json:"domainMappingReason"`
	RevisionReason      *JobTerminalConditionRevisionReasonEnum      `json:"revisionReason"`
	ExecutionReason     *JobTerminalConditionExecutionReasonEnum     `json:"executionReason"`
}

type jsonJobTerminalCondition JobTerminalCondition

func (r *JobTerminalCondition) UnmarshalJSON(data []byte) error {
	var res jsonJobTerminalCondition
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobTerminalCondition
	} else {

		r.Type = res.Type

		r.State = res.State

		r.Message = res.Message

		r.LastTransitionTime = res.LastTransitionTime

		r.Severity = res.Severity

		r.Reason = res.Reason

		r.InternalReason = res.InternalReason

		r.DomainMappingReason = res.DomainMappingReason

		r.RevisionReason = res.RevisionReason

		r.ExecutionReason = res.ExecutionReason

	}
	return nil
}

// This object is used to assert a desired state where this JobTerminalCondition is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobTerminalCondition *JobTerminalCondition = &JobTerminalCondition{empty: true}

func (r *JobTerminalCondition) Empty() bool {
	return r.empty
}

func (r *JobTerminalCondition) String() string {
	return dcl.SprintResource(r)
}

func (r *JobTerminalCondition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobConditions struct {
	empty              bool                              `json:"-"`
	Type               *string                           `json:"type"`
	State              *JobConditionsStateEnum           `json:"state"`
	Message            *string                           `json:"message"`
	LastTransitionTime *string                           `json:"lastTransitionTime"`
	Severity           *JobConditionsSeverityEnum        `json:"severity"`
	Reason             *JobConditionsReasonEnum          `json:"reason"`
	RevisionReason     *JobConditionsRevisionReasonEnum  `json:"revisionReason"`
	ExecutionReason    *JobConditionsExecutionReasonEnum `json:"executionReason"`
}

type jsonJobConditions JobConditions

func (r *JobConditions) UnmarshalJSON(data []byte) error {
	var res jsonJobConditions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobConditions
	} else {

		r.Type = res.Type

		r.State = res.State

		r.Message = res.Message

		r.LastTransitionTime = res.LastTransitionTime

		r.Severity = res.Severity

		r.Reason = res.Reason

		r.RevisionReason = res.RevisionReason

		r.ExecutionReason = res.ExecutionReason

	}
	return nil
}

// This object is used to assert a desired state where this JobConditions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobConditions *JobConditions = &JobConditions{empty: true}

func (r *JobConditions) Empty() bool {
	return r.empty
}

func (r *JobConditions) String() string {
	return dcl.SprintResource(r)
}

func (r *JobConditions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobLatestSucceededExecution struct {
	empty      bool    `json:"-"`
	Name       *string `json:"name"`
	CreateTime *string `json:"createTime"`
}

type jsonJobLatestSucceededExecution JobLatestSucceededExecution

func (r *JobLatestSucceededExecution) UnmarshalJSON(data []byte) error {
	var res jsonJobLatestSucceededExecution
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobLatestSucceededExecution
	} else {

		r.Name = res.Name

		r.CreateTime = res.CreateTime

	}
	return nil
}

// This object is used to assert a desired state where this JobLatestSucceededExecution is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobLatestSucceededExecution *JobLatestSucceededExecution = &JobLatestSucceededExecution{empty: true}

func (r *JobLatestSucceededExecution) Empty() bool {
	return r.empty
}

func (r *JobLatestSucceededExecution) String() string {
	return dcl.SprintResource(r)
}

func (r *JobLatestSucceededExecution) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type JobLatestCreatedExecution struct {
	empty      bool    `json:"-"`
	Name       *string `json:"name"`
	CreateTime *string `json:"createTime"`
}

type jsonJobLatestCreatedExecution JobLatestCreatedExecution

func (r *JobLatestCreatedExecution) UnmarshalJSON(data []byte) error {
	var res jsonJobLatestCreatedExecution
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyJobLatestCreatedExecution
	} else {

		r.Name = res.Name

		r.CreateTime = res.CreateTime

	}
	return nil
}

// This object is used to assert a desired state where this JobLatestCreatedExecution is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyJobLatestCreatedExecution *JobLatestCreatedExecution = &JobLatestCreatedExecution{empty: true}

func (r *JobLatestCreatedExecution) Empty() bool {
	return r.empty
}

func (r *JobLatestCreatedExecution) String() string {
	return dcl.SprintResource(r)
}

func (r *JobLatestCreatedExecution) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.New().Sum([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Job) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "run",
		Type:    "Job",
		Version: "alpha",
	}
}

func (r *Job) ID() (string, error) {
	if err := extractJobFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                       dcl.ValueOrEmptyString(nr.Name),
		"uid":                        dcl.ValueOrEmptyString(nr.Uid),
		"generation":                 dcl.ValueOrEmptyString(nr.Generation),
		"labels":                     dcl.ValueOrEmptyString(nr.Labels),
		"annotations":                dcl.ValueOrEmptyString(nr.Annotations),
		"create_time":                dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":                dcl.ValueOrEmptyString(nr.UpdateTime),
		"delete_time":                dcl.ValueOrEmptyString(nr.DeleteTime),
		"expire_time":                dcl.ValueOrEmptyString(nr.ExpireTime),
		"creator":                    dcl.ValueOrEmptyString(nr.Creator),
		"last_modifier":              dcl.ValueOrEmptyString(nr.LastModifier),
		"client":                     dcl.ValueOrEmptyString(nr.Client),
		"client_version":             dcl.ValueOrEmptyString(nr.ClientVersion),
		"launch_stage":               dcl.ValueOrEmptyString(nr.LaunchStage),
		"binary_authorization":       dcl.ValueOrEmptyString(nr.BinaryAuthorization),
		"template":                   dcl.ValueOrEmptyString(nr.Template),
		"observed_generation":        dcl.ValueOrEmptyString(nr.ObservedGeneration),
		"terminal_condition":         dcl.ValueOrEmptyString(nr.TerminalCondition),
		"conditions":                 dcl.ValueOrEmptyString(nr.Conditions),
		"execution_count":            dcl.ValueOrEmptyString(nr.ExecutionCount),
		"latest_succeeded_execution": dcl.ValueOrEmptyString(nr.LatestSucceededExecution),
		"latest_created_execution":   dcl.ValueOrEmptyString(nr.LatestCreatedExecution),
		"reconciling":                dcl.ValueOrEmptyString(nr.Reconciling),
		"etag":                       dcl.ValueOrEmptyString(nr.Etag),
		"project":                    dcl.ValueOrEmptyString(nr.Project),
		"location":                   dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/jobs/{{name}}", params), nil
}

const JobMaxPage = -1

type JobList struct {
	Items []*Job

	nextToken string

	pageSize int32

	resource *Job
}

func (l *JobList) HasNext() bool {
	return l.nextToken != ""
}

func (l *JobList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listJob(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListJob(ctx context.Context, project, location string) (*JobList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListJobWithMaxResults(ctx, project, location, JobMaxPage)

}

func (c *Client) ListJobWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*JobList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Job{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listJob(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &JobList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetJob(ctx context.Context, r *Job) (*Job, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractJobFields(r)

	b, err := c.getJobRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalJob(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeJobNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractJobFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteJob(ctx context.Context, r *Job) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Job resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Job...")
	deleteOp := deleteJobOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllJob deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllJob(ctx context.Context, project, location string, filter func(*Job) bool) error {
	listObj, err := c.ListJob(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllJob(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllJob(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyJob(ctx context.Context, rawDesired *Job, opts ...dcl.ApplyOption) (*Job, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Job
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyJobHelper(c, ctx, rawDesired, opts...)
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

func applyJobHelper(c *Client, ctx context.Context, rawDesired *Job, opts ...dcl.ApplyOption) (*Job, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyJob...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractJobFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.jobDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToJobDiffs(c.Config, fieldDiffs, opts)
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
	var ops []jobApiOperation
	if create {
		ops = append(ops, &createJobOperation{})
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
	return applyJobDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyJobDiff(c *Client, ctx context.Context, desired *Job, rawDesired *Job, ops []jobApiOperation, opts ...dcl.ApplyOption) (*Job, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetJob(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createJobOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapJob(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeJobNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeJobNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeJobDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractJobFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractJobFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffJob(c, newDesired, newState)
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

func (r *Job) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
