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
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Service struct {
	Name                  *string                     `json:"name"`
	Description           *string                     `json:"description"`
	Uid                   *string                     `json:"uid"`
	Generation            *int64                      `json:"generation"`
	Labels                map[string]string           `json:"labels"`
	Annotations           map[string]string           `json:"annotations"`
	CreateTime            *string                     `json:"createTime"`
	UpdateTime            *string                     `json:"updateTime"`
	DeleteTime            *string                     `json:"deleteTime"`
	ExpireTime            *string                     `json:"expireTime"`
	Creator               *string                     `json:"creator"`
	LastModifier          *string                     `json:"lastModifier"`
	Client                *string                     `json:"client"`
	ClientVersion         *string                     `json:"clientVersion"`
	Ingress               *ServiceIngressEnum         `json:"ingress"`
	LaunchStage           *ServiceLaunchStageEnum     `json:"launchStage"`
	BinaryAuthorization   *ServiceBinaryAuthorization `json:"binaryAuthorization"`
	Template              *ServiceTemplate            `json:"template"`
	Traffic               []ServiceTraffic            `json:"traffic"`
	TerminalCondition     *ServiceTerminalCondition   `json:"terminalCondition"`
	LatestReadyRevision   *string                     `json:"latestReadyRevision"`
	LatestCreatedRevision *string                     `json:"latestCreatedRevision"`
	TrafficStatuses       []ServiceTrafficStatuses    `json:"trafficStatuses"`
	Uri                   *string                     `json:"uri"`
	Reconciling           *bool                       `json:"reconciling"`
	Etag                  *string                     `json:"etag"`
	Project               *string                     `json:"project"`
	Location              *string                     `json:"location"`
}

func (r *Service) String() string {
	return dcl.SprintResource(r)
}

// The enum ServiceIngressEnum.
type ServiceIngressEnum string

// ServiceIngressEnumRef returns a *ServiceIngressEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceIngressEnumRef(s string) *ServiceIngressEnum {
	v := ServiceIngressEnum(s)
	return &v
}

func (v ServiceIngressEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INGRESS_TRAFFIC_UNSPECIFIED", "INGRESS_TRAFFIC_ALL", "INGRESS_TRAFFIC_INTERNAL_ONLY", "INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ServiceIngressEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceLaunchStageEnum.
type ServiceLaunchStageEnum string

// ServiceLaunchStageEnumRef returns a *ServiceLaunchStageEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceLaunchStageEnumRef(s string) *ServiceLaunchStageEnum {
	v := ServiceLaunchStageEnum(s)
	return &v
}

func (v ServiceLaunchStageEnum) Validate() error {
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
		Enum:  "ServiceLaunchStageEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTemplateVPCAccessEgressEnum.
type ServiceTemplateVPCAccessEgressEnum string

// ServiceTemplateVPCAccessEgressEnumRef returns a *ServiceTemplateVPCAccessEgressEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTemplateVPCAccessEgressEnumRef(s string) *ServiceTemplateVPCAccessEgressEnum {
	v := ServiceTemplateVPCAccessEgressEnum(s)
	return &v
}

func (v ServiceTemplateVPCAccessEgressEnum) Validate() error {
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
		Enum:  "ServiceTemplateVPCAccessEgressEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTemplateExecutionEnvironmentEnum.
type ServiceTemplateExecutionEnvironmentEnum string

// ServiceTemplateExecutionEnvironmentEnumRef returns a *ServiceTemplateExecutionEnvironmentEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTemplateExecutionEnvironmentEnumRef(s string) *ServiceTemplateExecutionEnvironmentEnum {
	v := ServiceTemplateExecutionEnvironmentEnum(s)
	return &v
}

func (v ServiceTemplateExecutionEnvironmentEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"EXECUTION_ENVIRONMENT_UNSPECIFIED", "EXECUTION_ENVIRONMENT_GEN1", "EXECUTION_ENVIRONMENT_GEN2"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ServiceTemplateExecutionEnvironmentEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTrafficTypeEnum.
type ServiceTrafficTypeEnum string

// ServiceTrafficTypeEnumRef returns a *ServiceTrafficTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTrafficTypeEnumRef(s string) *ServiceTrafficTypeEnum {
	v := ServiceTrafficTypeEnum(s)
	return &v
}

func (v ServiceTrafficTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED", "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST", "TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ServiceTrafficTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTerminalConditionStateEnum.
type ServiceTerminalConditionStateEnum string

// ServiceTerminalConditionStateEnumRef returns a *ServiceTerminalConditionStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTerminalConditionStateEnumRef(s string) *ServiceTerminalConditionStateEnum {
	v := ServiceTerminalConditionStateEnum(s)
	return &v
}

func (v ServiceTerminalConditionStateEnum) Validate() error {
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
		Enum:  "ServiceTerminalConditionStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTerminalConditionSeverityEnum.
type ServiceTerminalConditionSeverityEnum string

// ServiceTerminalConditionSeverityEnumRef returns a *ServiceTerminalConditionSeverityEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTerminalConditionSeverityEnumRef(s string) *ServiceTerminalConditionSeverityEnum {
	v := ServiceTerminalConditionSeverityEnum(s)
	return &v
}

func (v ServiceTerminalConditionSeverityEnum) Validate() error {
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
		Enum:  "ServiceTerminalConditionSeverityEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTerminalConditionReasonEnum.
type ServiceTerminalConditionReasonEnum string

// ServiceTerminalConditionReasonEnumRef returns a *ServiceTerminalConditionReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTerminalConditionReasonEnumRef(s string) *ServiceTerminalConditionReasonEnum {
	v := ServiceTerminalConditionReasonEnum(s)
	return &v
}

func (v ServiceTerminalConditionReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"COMMON_REASON_UNDEFINED", "UNKNOWN", "REVISION_FAILED", "PROGRESS_DEADLINE_EXCEEDED", "CONTAINER_MISSING", "CONTAINER_PERMISSION_DENIED", "CONTAINER_IMAGE_UNAUTHORIZED", "CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED", "ENCRYPTION_KEY_PERMISSION_DENIED", "ENCRYPTION_KEY_CHECK_FAILED", "SECRETS_ACCESS_CHECK_FAILED", "WAITING_FOR_OPERATION", "IMMEDIATE_RETRY", "POSTPONED_RETRY", "INTERNAL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ServiceTerminalConditionReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTerminalConditionRevisionReasonEnum.
type ServiceTerminalConditionRevisionReasonEnum string

// ServiceTerminalConditionRevisionReasonEnumRef returns a *ServiceTerminalConditionRevisionReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTerminalConditionRevisionReasonEnumRef(s string) *ServiceTerminalConditionRevisionReasonEnum {
	v := ServiceTerminalConditionRevisionReasonEnum(s)
	return &v
}

func (v ServiceTerminalConditionRevisionReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"REVISION_REASON_UNDEFINED", "PENDING", "RESERVE", "RETIRED", "RETIRING", "RECREATING", "HEALTH_CHECK_CONTAINER_ERROR", "CUSTOMIZED_PATH_RESPONSE_PENDING", "MIN_INSTANCES_NOT_PROVISIONED", "ACTIVE_REVISION_LIMIT_REACHED", "NO_DEPLOYMENT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ServiceTerminalConditionRevisionReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTerminalConditionJobReasonEnum.
type ServiceTerminalConditionJobReasonEnum string

// ServiceTerminalConditionJobReasonEnumRef returns a *ServiceTerminalConditionJobReasonEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTerminalConditionJobReasonEnumRef(s string) *ServiceTerminalConditionJobReasonEnum {
	v := ServiceTerminalConditionJobReasonEnum(s)
	return &v
}

func (v ServiceTerminalConditionJobReasonEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"JOB_REASON_UNDEFINED", "JOB_STATUS_SERVICE_POLLING_ERROR"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ServiceTerminalConditionJobReasonEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum ServiceTrafficStatusesTypeEnum.
type ServiceTrafficStatusesTypeEnum string

// ServiceTrafficStatusesTypeEnumRef returns a *ServiceTrafficStatusesTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func ServiceTrafficStatusesTypeEnumRef(s string) *ServiceTrafficStatusesTypeEnum {
	v := ServiceTrafficStatusesTypeEnum(s)
	return &v
}

func (v ServiceTrafficStatusesTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED", "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST", "TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "ServiceTrafficStatusesTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type ServiceBinaryAuthorization struct {
	empty                   bool    `json:"-"`
	UseDefault              *bool   `json:"useDefault"`
	BreakglassJustification *string `json:"breakglassJustification"`
}

type jsonServiceBinaryAuthorization ServiceBinaryAuthorization

func (r *ServiceBinaryAuthorization) UnmarshalJSON(data []byte) error {
	var res jsonServiceBinaryAuthorization
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceBinaryAuthorization
	} else {

		r.UseDefault = res.UseDefault

		r.BreakglassJustification = res.BreakglassJustification

	}
	return nil
}

// This object is used to assert a desired state where this ServiceBinaryAuthorization is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceBinaryAuthorization *ServiceBinaryAuthorization = &ServiceBinaryAuthorization{empty: true}

func (r *ServiceBinaryAuthorization) Empty() bool {
	return r.empty
}

func (r *ServiceBinaryAuthorization) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceBinaryAuthorization) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplate struct {
	empty                bool                                     `json:"-"`
	Revision             *string                                  `json:"revision"`
	Labels               map[string]string                        `json:"labels"`
	Annotations          map[string]string                        `json:"annotations"`
	Scaling              *ServiceTemplateScaling                  `json:"scaling"`
	VPCAccess            *ServiceTemplateVPCAccess                `json:"vpcAccess"`
	ContainerConcurrency *int64                                   `json:"containerConcurrency"`
	Timeout              *string                                  `json:"timeout"`
	ServiceAccount       *string                                  `json:"serviceAccount"`
	Containers           []ServiceTemplateContainers              `json:"containers"`
	Volumes              []ServiceTemplateVolumes                 `json:"volumes"`
	ExecutionEnvironment *ServiceTemplateExecutionEnvironmentEnum `json:"executionEnvironment"`
}

type jsonServiceTemplate ServiceTemplate

func (r *ServiceTemplate) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplate
	} else {

		r.Revision = res.Revision

		r.Labels = res.Labels

		r.Annotations = res.Annotations

		r.Scaling = res.Scaling

		r.VPCAccess = res.VPCAccess

		r.ContainerConcurrency = res.ContainerConcurrency

		r.Timeout = res.Timeout

		r.ServiceAccount = res.ServiceAccount

		r.Containers = res.Containers

		r.Volumes = res.Volumes

		r.ExecutionEnvironment = res.ExecutionEnvironment

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplate *ServiceTemplate = &ServiceTemplate{empty: true}

func (r *ServiceTemplate) Empty() bool {
	return r.empty
}

func (r *ServiceTemplate) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateScaling struct {
	empty            bool   `json:"-"`
	MinInstanceCount *int64 `json:"minInstanceCount"`
	MaxInstanceCount *int64 `json:"maxInstanceCount"`
}

type jsonServiceTemplateScaling ServiceTemplateScaling

func (r *ServiceTemplateScaling) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateScaling
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateScaling
	} else {

		r.MinInstanceCount = res.MinInstanceCount

		r.MaxInstanceCount = res.MaxInstanceCount

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateScaling is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateScaling *ServiceTemplateScaling = &ServiceTemplateScaling{empty: true}

func (r *ServiceTemplateScaling) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateScaling) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateScaling) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateVPCAccess struct {
	empty     bool                                `json:"-"`
	Connector *string                             `json:"connector"`
	Egress    *ServiceTemplateVPCAccessEgressEnum `json:"egress"`
}

type jsonServiceTemplateVPCAccess ServiceTemplateVPCAccess

func (r *ServiceTemplateVPCAccess) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateVPCAccess
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateVPCAccess
	} else {

		r.Connector = res.Connector

		r.Egress = res.Egress

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateVPCAccess is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateVPCAccess *ServiceTemplateVPCAccess = &ServiceTemplateVPCAccess{empty: true}

func (r *ServiceTemplateVPCAccess) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateVPCAccess) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateVPCAccess) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateContainers struct {
	empty        bool                                    `json:"-"`
	Name         *string                                 `json:"name"`
	Image        *string                                 `json:"image"`
	Command      []string                                `json:"command"`
	Args         []string                                `json:"args"`
	Env          []ServiceTemplateContainersEnv          `json:"env"`
	Resources    *ServiceTemplateContainersResources     `json:"resources"`
	Ports        []ServiceTemplateContainersPorts        `json:"ports"`
	VolumeMounts []ServiceTemplateContainersVolumeMounts `json:"volumeMounts"`
}

type jsonServiceTemplateContainers ServiceTemplateContainers

func (r *ServiceTemplateContainers) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateContainers
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateContainers
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

// This object is used to assert a desired state where this ServiceTemplateContainers is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateContainers *ServiceTemplateContainers = &ServiceTemplateContainers{empty: true}

func (r *ServiceTemplateContainers) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateContainers) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateContainers) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateContainersEnv struct {
	empty       bool                                     `json:"-"`
	Name        *string                                  `json:"name"`
	Value       *string                                  `json:"value"`
	ValueSource *ServiceTemplateContainersEnvValueSource `json:"valueSource"`
}

type jsonServiceTemplateContainersEnv ServiceTemplateContainersEnv

func (r *ServiceTemplateContainersEnv) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateContainersEnv
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateContainersEnv
	} else {

		r.Name = res.Name

		r.Value = res.Value

		r.ValueSource = res.ValueSource

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateContainersEnv is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateContainersEnv *ServiceTemplateContainersEnv = &ServiceTemplateContainersEnv{empty: true}

func (r *ServiceTemplateContainersEnv) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateContainersEnv) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateContainersEnv) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateContainersEnvValueSource struct {
	empty        bool                                                 `json:"-"`
	SecretKeyRef *ServiceTemplateContainersEnvValueSourceSecretKeyRef `json:"secretKeyRef"`
}

type jsonServiceTemplateContainersEnvValueSource ServiceTemplateContainersEnvValueSource

func (r *ServiceTemplateContainersEnvValueSource) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateContainersEnvValueSource
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateContainersEnvValueSource
	} else {

		r.SecretKeyRef = res.SecretKeyRef

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateContainersEnvValueSource is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateContainersEnvValueSource *ServiceTemplateContainersEnvValueSource = &ServiceTemplateContainersEnvValueSource{empty: true}

func (r *ServiceTemplateContainersEnvValueSource) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateContainersEnvValueSource) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateContainersEnvValueSource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateContainersEnvValueSourceSecretKeyRef struct {
	empty   bool    `json:"-"`
	Secret  *string `json:"secret"`
	Version *string `json:"version"`
}

type jsonServiceTemplateContainersEnvValueSourceSecretKeyRef ServiceTemplateContainersEnvValueSourceSecretKeyRef

func (r *ServiceTemplateContainersEnvValueSourceSecretKeyRef) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateContainersEnvValueSourceSecretKeyRef
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateContainersEnvValueSourceSecretKeyRef
	} else {

		r.Secret = res.Secret

		r.Version = res.Version

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateContainersEnvValueSourceSecretKeyRef is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateContainersEnvValueSourceSecretKeyRef *ServiceTemplateContainersEnvValueSourceSecretKeyRef = &ServiceTemplateContainersEnvValueSourceSecretKeyRef{empty: true}

func (r *ServiceTemplateContainersEnvValueSourceSecretKeyRef) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateContainersEnvValueSourceSecretKeyRef) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateContainersEnvValueSourceSecretKeyRef) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateContainersResources struct {
	empty   bool              `json:"-"`
	Limits  map[string]string `json:"limits"`
	CpuIdle *bool             `json:"cpuIdle"`
}

type jsonServiceTemplateContainersResources ServiceTemplateContainersResources

func (r *ServiceTemplateContainersResources) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateContainersResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateContainersResources
	} else {

		r.Limits = res.Limits

		r.CpuIdle = res.CpuIdle

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateContainersResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateContainersResources *ServiceTemplateContainersResources = &ServiceTemplateContainersResources{empty: true}

func (r *ServiceTemplateContainersResources) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateContainersResources) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateContainersResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateContainersPorts struct {
	empty         bool    `json:"-"`
	Name          *string `json:"name"`
	ContainerPort *int64  `json:"containerPort"`
}

type jsonServiceTemplateContainersPorts ServiceTemplateContainersPorts

func (r *ServiceTemplateContainersPorts) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateContainersPorts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateContainersPorts
	} else {

		r.Name = res.Name

		r.ContainerPort = res.ContainerPort

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateContainersPorts is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateContainersPorts *ServiceTemplateContainersPorts = &ServiceTemplateContainersPorts{empty: true}

func (r *ServiceTemplateContainersPorts) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateContainersPorts) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateContainersPorts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateContainersVolumeMounts struct {
	empty     bool    `json:"-"`
	Name      *string `json:"name"`
	MountPath *string `json:"mountPath"`
}

type jsonServiceTemplateContainersVolumeMounts ServiceTemplateContainersVolumeMounts

func (r *ServiceTemplateContainersVolumeMounts) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateContainersVolumeMounts
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateContainersVolumeMounts
	} else {

		r.Name = res.Name

		r.MountPath = res.MountPath

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateContainersVolumeMounts is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateContainersVolumeMounts *ServiceTemplateContainersVolumeMounts = &ServiceTemplateContainersVolumeMounts{empty: true}

func (r *ServiceTemplateContainersVolumeMounts) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateContainersVolumeMounts) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateContainersVolumeMounts) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateVolumes struct {
	empty            bool                                    `json:"-"`
	Name             *string                                 `json:"name"`
	Secret           *ServiceTemplateVolumesSecret           `json:"secret"`
	CloudSqlInstance *ServiceTemplateVolumesCloudSqlInstance `json:"cloudSqlInstance"`
}

type jsonServiceTemplateVolumes ServiceTemplateVolumes

func (r *ServiceTemplateVolumes) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateVolumes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateVolumes
	} else {

		r.Name = res.Name

		r.Secret = res.Secret

		r.CloudSqlInstance = res.CloudSqlInstance

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateVolumes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateVolumes *ServiceTemplateVolumes = &ServiceTemplateVolumes{empty: true}

func (r *ServiceTemplateVolumes) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateVolumes) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateVolumes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateVolumesSecret struct {
	empty       bool                                `json:"-"`
	Secret      *string                             `json:"secret"`
	Items       []ServiceTemplateVolumesSecretItems `json:"items"`
	DefaultMode *int64                              `json:"defaultMode"`
}

type jsonServiceTemplateVolumesSecret ServiceTemplateVolumesSecret

func (r *ServiceTemplateVolumesSecret) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateVolumesSecret
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateVolumesSecret
	} else {

		r.Secret = res.Secret

		r.Items = res.Items

		r.DefaultMode = res.DefaultMode

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateVolumesSecret is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateVolumesSecret *ServiceTemplateVolumesSecret = &ServiceTemplateVolumesSecret{empty: true}

func (r *ServiceTemplateVolumesSecret) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateVolumesSecret) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateVolumesSecret) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateVolumesSecretItems struct {
	empty   bool    `json:"-"`
	Path    *string `json:"path"`
	Version *string `json:"version"`
	Mode    *int64  `json:"mode"`
}

type jsonServiceTemplateVolumesSecretItems ServiceTemplateVolumesSecretItems

func (r *ServiceTemplateVolumesSecretItems) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateVolumesSecretItems
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateVolumesSecretItems
	} else {

		r.Path = res.Path

		r.Version = res.Version

		r.Mode = res.Mode

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateVolumesSecretItems is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateVolumesSecretItems *ServiceTemplateVolumesSecretItems = &ServiceTemplateVolumesSecretItems{empty: true}

func (r *ServiceTemplateVolumesSecretItems) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateVolumesSecretItems) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateVolumesSecretItems) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTemplateVolumesCloudSqlInstance struct {
	empty     bool     `json:"-"`
	Instances []string `json:"instances"`
}

type jsonServiceTemplateVolumesCloudSqlInstance ServiceTemplateVolumesCloudSqlInstance

func (r *ServiceTemplateVolumesCloudSqlInstance) UnmarshalJSON(data []byte) error {
	var res jsonServiceTemplateVolumesCloudSqlInstance
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTemplateVolumesCloudSqlInstance
	} else {

		r.Instances = res.Instances

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTemplateVolumesCloudSqlInstance is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTemplateVolumesCloudSqlInstance *ServiceTemplateVolumesCloudSqlInstance = &ServiceTemplateVolumesCloudSqlInstance{empty: true}

func (r *ServiceTemplateVolumesCloudSqlInstance) Empty() bool {
	return r.empty
}

func (r *ServiceTemplateVolumesCloudSqlInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTemplateVolumesCloudSqlInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTraffic struct {
	empty    bool                    `json:"-"`
	Type     *ServiceTrafficTypeEnum `json:"type"`
	Revision *string                 `json:"revision"`
	Percent  *int64                  `json:"percent"`
	Tag      *string                 `json:"tag"`
}

type jsonServiceTraffic ServiceTraffic

func (r *ServiceTraffic) UnmarshalJSON(data []byte) error {
	var res jsonServiceTraffic
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTraffic
	} else {

		r.Type = res.Type

		r.Revision = res.Revision

		r.Percent = res.Percent

		r.Tag = res.Tag

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTraffic is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTraffic *ServiceTraffic = &ServiceTraffic{empty: true}

func (r *ServiceTraffic) Empty() bool {
	return r.empty
}

func (r *ServiceTraffic) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTraffic) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTerminalCondition struct {
	empty              bool                                        `json:"-"`
	Type               *string                                     `json:"type"`
	State              *ServiceTerminalConditionStateEnum          `json:"state"`
	Message            *string                                     `json:"message"`
	LastTransitionTime *string                                     `json:"lastTransitionTime"`
	Severity           *ServiceTerminalConditionSeverityEnum       `json:"severity"`
	Reason             *ServiceTerminalConditionReasonEnum         `json:"reason"`
	RevisionReason     *ServiceTerminalConditionRevisionReasonEnum `json:"revisionReason"`
	JobReason          *ServiceTerminalConditionJobReasonEnum      `json:"jobReason"`
}

type jsonServiceTerminalCondition ServiceTerminalCondition

func (r *ServiceTerminalCondition) UnmarshalJSON(data []byte) error {
	var res jsonServiceTerminalCondition
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTerminalCondition
	} else {

		r.Type = res.Type

		r.State = res.State

		r.Message = res.Message

		r.LastTransitionTime = res.LastTransitionTime

		r.Severity = res.Severity

		r.Reason = res.Reason

		r.RevisionReason = res.RevisionReason

		r.JobReason = res.JobReason

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTerminalCondition is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTerminalCondition *ServiceTerminalCondition = &ServiceTerminalCondition{empty: true}

func (r *ServiceTerminalCondition) Empty() bool {
	return r.empty
}

func (r *ServiceTerminalCondition) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTerminalCondition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ServiceTrafficStatuses struct {
	empty    bool                            `json:"-"`
	Type     *ServiceTrafficStatusesTypeEnum `json:"type"`
	Revision *string                         `json:"revision"`
	Percent  *int64                          `json:"percent"`
	Tag      *string                         `json:"tag"`
	Uri      *string                         `json:"uri"`
}

type jsonServiceTrafficStatuses ServiceTrafficStatuses

func (r *ServiceTrafficStatuses) UnmarshalJSON(data []byte) error {
	var res jsonServiceTrafficStatuses
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyServiceTrafficStatuses
	} else {

		r.Type = res.Type

		r.Revision = res.Revision

		r.Percent = res.Percent

		r.Tag = res.Tag

		r.Uri = res.Uri

	}
	return nil
}

// This object is used to assert a desired state where this ServiceTrafficStatuses is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyServiceTrafficStatuses *ServiceTrafficStatuses = &ServiceTrafficStatuses{empty: true}

func (r *ServiceTrafficStatuses) Empty() bool {
	return r.empty
}

func (r *ServiceTrafficStatuses) String() string {
	return dcl.SprintResource(r)
}

func (r *ServiceTrafficStatuses) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Service) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "run",
		Type:    "Service",
		Version: "alpha",
	}
}

func (r *Service) ID() (string, error) {
	if err := extractServiceFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                    dcl.ValueOrEmptyString(nr.Name),
		"description":             dcl.ValueOrEmptyString(nr.Description),
		"uid":                     dcl.ValueOrEmptyString(nr.Uid),
		"generation":              dcl.ValueOrEmptyString(nr.Generation),
		"labels":                  dcl.ValueOrEmptyString(nr.Labels),
		"annotations":             dcl.ValueOrEmptyString(nr.Annotations),
		"create_time":             dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":             dcl.ValueOrEmptyString(nr.UpdateTime),
		"delete_time":             dcl.ValueOrEmptyString(nr.DeleteTime),
		"expire_time":             dcl.ValueOrEmptyString(nr.ExpireTime),
		"creator":                 dcl.ValueOrEmptyString(nr.Creator),
		"last_modifier":           dcl.ValueOrEmptyString(nr.LastModifier),
		"client":                  dcl.ValueOrEmptyString(nr.Client),
		"client_version":          dcl.ValueOrEmptyString(nr.ClientVersion),
		"ingress":                 dcl.ValueOrEmptyString(nr.Ingress),
		"launch_stage":            dcl.ValueOrEmptyString(nr.LaunchStage),
		"binary_authorization":    dcl.ValueOrEmptyString(nr.BinaryAuthorization),
		"template":                dcl.ValueOrEmptyString(nr.Template),
		"traffic":                 dcl.ValueOrEmptyString(nr.Traffic),
		"terminal_condition":      dcl.ValueOrEmptyString(nr.TerminalCondition),
		"latest_ready_revision":   dcl.ValueOrEmptyString(nr.LatestReadyRevision),
		"latest_created_revision": dcl.ValueOrEmptyString(nr.LatestCreatedRevision),
		"traffic_statuses":        dcl.ValueOrEmptyString(nr.TrafficStatuses),
		"uri":                     dcl.ValueOrEmptyString(nr.Uri),
		"reconciling":             dcl.ValueOrEmptyString(nr.Reconciling),
		"etag":                    dcl.ValueOrEmptyString(nr.Etag),
		"project":                 dcl.ValueOrEmptyString(nr.Project),
		"location":                dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/services/{{name}}", params), nil
}

const ServiceMaxPage = -1

type ServiceList struct {
	Items []*Service

	nextToken string

	pageSize int32

	resource *Service
}

func (l *ServiceList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ServiceList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listService(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListService(ctx context.Context, project, location string) (*ServiceList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListServiceWithMaxResults(ctx, project, location, ServiceMaxPage)

}

func (c *Client) ListServiceWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ServiceList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Service{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listService(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ServiceList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetService(ctx context.Context, r *Service) (*Service, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractServiceFields(r)

	b, err := c.getServiceRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalService(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeServiceNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractServiceFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteService(ctx context.Context, r *Service) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Service resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Service...")
	deleteOp := deleteServiceOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllService deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllService(ctx context.Context, project, location string, filter func(*Service) bool) error {
	listObj, err := c.ListService(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllService(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllService(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyService(ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (*Service, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Service
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyServiceHelper(c, ctx, rawDesired, opts...)
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

func applyServiceHelper(c *Client, ctx context.Context, rawDesired *Service, opts ...dcl.ApplyOption) (*Service, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyService...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractServiceFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.serviceDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToServiceDiffs(c.Config, fieldDiffs, opts)
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
	var ops []serviceApiOperation
	if create {
		ops = append(ops, &createServiceOperation{})
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
	return applyServiceDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyServiceDiff(c *Client, ctx context.Context, desired *Service, rawDesired *Service, ops []serviceApiOperation, opts ...dcl.ApplyOption) (*Service, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetService(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createServiceOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapService(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeServiceNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeServiceNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeServiceDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractServiceFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractServiceFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffService(c, newDesired, newState)
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

func (r *Service) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
