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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/run/alpha/run_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run/alpha"
)

// JobServer implements the gRPC interface for Job.
type JobServer struct{}

// ProtoToJobLaunchStageEnum converts a JobLaunchStageEnum enum from its proto representation.
func ProtoToRunAlphaJobLaunchStageEnum(e alphapb.RunAlphaJobLaunchStageEnum) *alpha.JobLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobLaunchStageEnum_name[int32(e)]; ok {
		e := alpha.JobLaunchStageEnum(n[len("RunAlphaJobLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTemplateTemplateExecutionEnvironmentEnum converts a JobTemplateTemplateExecutionEnvironmentEnum enum from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateExecutionEnvironmentEnum(e alphapb.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum) *alpha.JobTemplateTemplateExecutionEnvironmentEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum_name[int32(e)]; ok {
		e := alpha.JobTemplateTemplateExecutionEnvironmentEnum(n[len("RunAlphaJobTemplateTemplateExecutionEnvironmentEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTemplateTemplateVPCAccessEgressEnum converts a JobTemplateTemplateVPCAccessEgressEnum enum from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateVPCAccessEgressEnum(e alphapb.RunAlphaJobTemplateTemplateVPCAccessEgressEnum) *alpha.JobTemplateTemplateVPCAccessEgressEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTemplateTemplateVPCAccessEgressEnum_name[int32(e)]; ok {
		e := alpha.JobTemplateTemplateVPCAccessEgressEnum(n[len("RunAlphaJobTemplateTemplateVPCAccessEgressEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTerminalConditionStateEnum converts a JobTerminalConditionStateEnum enum from its proto representation.
func ProtoToRunAlphaJobTerminalConditionStateEnum(e alphapb.RunAlphaJobTerminalConditionStateEnum) *alpha.JobTerminalConditionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTerminalConditionStateEnum_name[int32(e)]; ok {
		e := alpha.JobTerminalConditionStateEnum(n[len("RunAlphaJobTerminalConditionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTerminalConditionSeverityEnum converts a JobTerminalConditionSeverityEnum enum from its proto representation.
func ProtoToRunAlphaJobTerminalConditionSeverityEnum(e alphapb.RunAlphaJobTerminalConditionSeverityEnum) *alpha.JobTerminalConditionSeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTerminalConditionSeverityEnum_name[int32(e)]; ok {
		e := alpha.JobTerminalConditionSeverityEnum(n[len("RunAlphaJobTerminalConditionSeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTerminalConditionReasonEnum converts a JobTerminalConditionReasonEnum enum from its proto representation.
func ProtoToRunAlphaJobTerminalConditionReasonEnum(e alphapb.RunAlphaJobTerminalConditionReasonEnum) *alpha.JobTerminalConditionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTerminalConditionReasonEnum_name[int32(e)]; ok {
		e := alpha.JobTerminalConditionReasonEnum(n[len("RunAlphaJobTerminalConditionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTerminalConditionInternalReasonEnum converts a JobTerminalConditionInternalReasonEnum enum from its proto representation.
func ProtoToRunAlphaJobTerminalConditionInternalReasonEnum(e alphapb.RunAlphaJobTerminalConditionInternalReasonEnum) *alpha.JobTerminalConditionInternalReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTerminalConditionInternalReasonEnum_name[int32(e)]; ok {
		e := alpha.JobTerminalConditionInternalReasonEnum(n[len("RunAlphaJobTerminalConditionInternalReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTerminalConditionDomainMappingReasonEnum converts a JobTerminalConditionDomainMappingReasonEnum enum from its proto representation.
func ProtoToRunAlphaJobTerminalConditionDomainMappingReasonEnum(e alphapb.RunAlphaJobTerminalConditionDomainMappingReasonEnum) *alpha.JobTerminalConditionDomainMappingReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTerminalConditionDomainMappingReasonEnum_name[int32(e)]; ok {
		e := alpha.JobTerminalConditionDomainMappingReasonEnum(n[len("RunAlphaJobTerminalConditionDomainMappingReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTerminalConditionRevisionReasonEnum converts a JobTerminalConditionRevisionReasonEnum enum from its proto representation.
func ProtoToRunAlphaJobTerminalConditionRevisionReasonEnum(e alphapb.RunAlphaJobTerminalConditionRevisionReasonEnum) *alpha.JobTerminalConditionRevisionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTerminalConditionRevisionReasonEnum_name[int32(e)]; ok {
		e := alpha.JobTerminalConditionRevisionReasonEnum(n[len("RunAlphaJobTerminalConditionRevisionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobTerminalConditionExecutionReasonEnum converts a JobTerminalConditionExecutionReasonEnum enum from its proto representation.
func ProtoToRunAlphaJobTerminalConditionExecutionReasonEnum(e alphapb.RunAlphaJobTerminalConditionExecutionReasonEnum) *alpha.JobTerminalConditionExecutionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobTerminalConditionExecutionReasonEnum_name[int32(e)]; ok {
		e := alpha.JobTerminalConditionExecutionReasonEnum(n[len("RunAlphaJobTerminalConditionExecutionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobConditionsStateEnum converts a JobConditionsStateEnum enum from its proto representation.
func ProtoToRunAlphaJobConditionsStateEnum(e alphapb.RunAlphaJobConditionsStateEnum) *alpha.JobConditionsStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobConditionsStateEnum_name[int32(e)]; ok {
		e := alpha.JobConditionsStateEnum(n[len("RunAlphaJobConditionsStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobConditionsSeverityEnum converts a JobConditionsSeverityEnum enum from its proto representation.
func ProtoToRunAlphaJobConditionsSeverityEnum(e alphapb.RunAlphaJobConditionsSeverityEnum) *alpha.JobConditionsSeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobConditionsSeverityEnum_name[int32(e)]; ok {
		e := alpha.JobConditionsSeverityEnum(n[len("RunAlphaJobConditionsSeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobConditionsReasonEnum converts a JobConditionsReasonEnum enum from its proto representation.
func ProtoToRunAlphaJobConditionsReasonEnum(e alphapb.RunAlphaJobConditionsReasonEnum) *alpha.JobConditionsReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobConditionsReasonEnum_name[int32(e)]; ok {
		e := alpha.JobConditionsReasonEnum(n[len("RunAlphaJobConditionsReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobConditionsRevisionReasonEnum converts a JobConditionsRevisionReasonEnum enum from its proto representation.
func ProtoToRunAlphaJobConditionsRevisionReasonEnum(e alphapb.RunAlphaJobConditionsRevisionReasonEnum) *alpha.JobConditionsRevisionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobConditionsRevisionReasonEnum_name[int32(e)]; ok {
		e := alpha.JobConditionsRevisionReasonEnum(n[len("RunAlphaJobConditionsRevisionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobConditionsExecutionReasonEnum converts a JobConditionsExecutionReasonEnum enum from its proto representation.
func ProtoToRunAlphaJobConditionsExecutionReasonEnum(e alphapb.RunAlphaJobConditionsExecutionReasonEnum) *alpha.JobConditionsExecutionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaJobConditionsExecutionReasonEnum_name[int32(e)]; ok {
		e := alpha.JobConditionsExecutionReasonEnum(n[len("RunAlphaJobConditionsExecutionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToJobBinaryAuthorization converts a JobBinaryAuthorization object from its proto representation.
func ProtoToRunAlphaJobBinaryAuthorization(p *alphapb.RunAlphaJobBinaryAuthorization) *alpha.JobBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.JobBinaryAuthorization{
		UseDefault:              dcl.Bool(p.GetUseDefault()),
		BreakglassJustification: dcl.StringOrNil(p.GetBreakglassJustification()),
	}
	return obj
}

// ProtoToJobTemplate converts a JobTemplate object from its proto representation.
func ProtoToRunAlphaJobTemplate(p *alphapb.RunAlphaJobTemplate) *alpha.JobTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplate{
		Parallelism: dcl.Int64OrNil(p.GetParallelism()),
		TaskCount:   dcl.Int64OrNil(p.GetTaskCount()),
		Template:    ProtoToRunAlphaJobTemplateTemplate(p.GetTemplate()),
	}
	return obj
}

// ProtoToJobTemplateTemplate converts a JobTemplateTemplate object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplate(p *alphapb.RunAlphaJobTemplateTemplate) *alpha.JobTemplateTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplate{
		MaxRetries:           dcl.Int64OrNil(p.GetMaxRetries()),
		Timeout:              dcl.StringOrNil(p.GetTimeout()),
		ServiceAccount:       dcl.StringOrNil(p.GetServiceAccount()),
		ExecutionEnvironment: ProtoToRunAlphaJobTemplateTemplateExecutionEnvironmentEnum(p.GetExecutionEnvironment()),
		EncryptionKey:        dcl.StringOrNil(p.GetEncryptionKey()),
		VPCAccess:            ProtoToRunAlphaJobTemplateTemplateVPCAccess(p.GetVpcAccess()),
	}
	for _, r := range p.GetContainers() {
		obj.Containers = append(obj.Containers, *ProtoToRunAlphaJobTemplateTemplateContainers(r))
	}
	for _, r := range p.GetVolumes() {
		obj.Volumes = append(obj.Volumes, *ProtoToRunAlphaJobTemplateTemplateVolumes(r))
	}
	return obj
}

// ProtoToJobTemplateTemplateContainers converts a JobTemplateTemplateContainers object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateContainers(p *alphapb.RunAlphaJobTemplateTemplateContainers) *alpha.JobTemplateTemplateContainers {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateContainers{
		Name:      dcl.StringOrNil(p.GetName()),
		Image:     dcl.StringOrNil(p.GetImage()),
		Resources: ProtoToRunAlphaJobTemplateTemplateContainersResources(p.GetResources()),
	}
	for _, r := range p.GetCommand() {
		obj.Command = append(obj.Command, r)
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetEnv() {
		obj.Env = append(obj.Env, *ProtoToRunAlphaJobTemplateTemplateContainersEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToRunAlphaJobTemplateTemplateContainersPorts(r))
	}
	for _, r := range p.GetVolumeMounts() {
		obj.VolumeMounts = append(obj.VolumeMounts, *ProtoToRunAlphaJobTemplateTemplateContainersVolumeMounts(r))
	}
	return obj
}

// ProtoToJobTemplateTemplateContainersEnv converts a JobTemplateTemplateContainersEnv object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateContainersEnv(p *alphapb.RunAlphaJobTemplateTemplateContainersEnv) *alpha.JobTemplateTemplateContainersEnv {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateContainersEnv{
		Name:        dcl.StringOrNil(p.GetName()),
		Value:       dcl.StringOrNil(p.GetValue()),
		ValueSource: ProtoToRunAlphaJobTemplateTemplateContainersEnvValueSource(p.GetValueSource()),
	}
	return obj
}

// ProtoToJobTemplateTemplateContainersEnvValueSource converts a JobTemplateTemplateContainersEnvValueSource object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateContainersEnvValueSource(p *alphapb.RunAlphaJobTemplateTemplateContainersEnvValueSource) *alpha.JobTemplateTemplateContainersEnvValueSource {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateContainersEnvValueSource{
		SecretKeyRef: ProtoToRunAlphaJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(p.GetSecretKeyRef()),
	}
	return obj
}

// ProtoToJobTemplateTemplateContainersEnvValueSourceSecretKeyRef converts a JobTemplateTemplateContainersEnvValueSourceSecretKeyRef object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateContainersEnvValueSourceSecretKeyRef(p *alphapb.RunAlphaJobTemplateTemplateContainersEnvValueSourceSecretKeyRef) *alpha.JobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateContainersEnvValueSourceSecretKeyRef{
		Secret:  dcl.StringOrNil(p.GetSecret()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToJobTemplateTemplateContainersResources converts a JobTemplateTemplateContainersResources object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateContainersResources(p *alphapb.RunAlphaJobTemplateTemplateContainersResources) *alpha.JobTemplateTemplateContainersResources {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateContainersResources{
		CpuIdle: dcl.Bool(p.GetCpuIdle()),
	}
	return obj
}

// ProtoToJobTemplateTemplateContainersPorts converts a JobTemplateTemplateContainersPorts object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateContainersPorts(p *alphapb.RunAlphaJobTemplateTemplateContainersPorts) *alpha.JobTemplateTemplateContainersPorts {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateContainersPorts{
		Name:          dcl.StringOrNil(p.GetName()),
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToJobTemplateTemplateContainersVolumeMounts converts a JobTemplateTemplateContainersVolumeMounts object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateContainersVolumeMounts(p *alphapb.RunAlphaJobTemplateTemplateContainersVolumeMounts) *alpha.JobTemplateTemplateContainersVolumeMounts {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateContainersVolumeMounts{
		Name:      dcl.StringOrNil(p.GetName()),
		MountPath: dcl.StringOrNil(p.GetMountPath()),
	}
	return obj
}

// ProtoToJobTemplateTemplateVolumes converts a JobTemplateTemplateVolumes object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateVolumes(p *alphapb.RunAlphaJobTemplateTemplateVolumes) *alpha.JobTemplateTemplateVolumes {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateVolumes{
		Name:             dcl.StringOrNil(p.GetName()),
		Secret:           ProtoToRunAlphaJobTemplateTemplateVolumesSecret(p.GetSecret()),
		CloudSqlInstance: ProtoToRunAlphaJobTemplateTemplateVolumesCloudSqlInstance(p.GetCloudSqlInstance()),
	}
	return obj
}

// ProtoToJobTemplateTemplateVolumesSecret converts a JobTemplateTemplateVolumesSecret object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateVolumesSecret(p *alphapb.RunAlphaJobTemplateTemplateVolumesSecret) *alpha.JobTemplateTemplateVolumesSecret {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateVolumesSecret{
		Secret:      dcl.StringOrNil(p.GetSecret()),
		DefaultMode: dcl.Int64OrNil(p.GetDefaultMode()),
	}
	for _, r := range p.GetItems() {
		obj.Items = append(obj.Items, *ProtoToRunAlphaJobTemplateTemplateVolumesSecretItems(r))
	}
	return obj
}

// ProtoToJobTemplateTemplateVolumesSecretItems converts a JobTemplateTemplateVolumesSecretItems object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateVolumesSecretItems(p *alphapb.RunAlphaJobTemplateTemplateVolumesSecretItems) *alpha.JobTemplateTemplateVolumesSecretItems {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateVolumesSecretItems{
		Path:    dcl.StringOrNil(p.GetPath()),
		Version: dcl.StringOrNil(p.GetVersion()),
		Mode:    dcl.Int64OrNil(p.GetMode()),
	}
	return obj
}

// ProtoToJobTemplateTemplateVolumesCloudSqlInstance converts a JobTemplateTemplateVolumesCloudSqlInstance object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateVolumesCloudSqlInstance(p *alphapb.RunAlphaJobTemplateTemplateVolumesCloudSqlInstance) *alpha.JobTemplateTemplateVolumesCloudSqlInstance {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateVolumesCloudSqlInstance{}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, r)
	}
	return obj
}

// ProtoToJobTemplateTemplateVPCAccess converts a JobTemplateTemplateVPCAccess object from its proto representation.
func ProtoToRunAlphaJobTemplateTemplateVPCAccess(p *alphapb.RunAlphaJobTemplateTemplateVPCAccess) *alpha.JobTemplateTemplateVPCAccess {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTemplateTemplateVPCAccess{
		Connector: dcl.StringOrNil(p.GetConnector()),
		Egress:    ProtoToRunAlphaJobTemplateTemplateVPCAccessEgressEnum(p.GetEgress()),
	}
	return obj
}

// ProtoToJobTerminalCondition converts a JobTerminalCondition object from its proto representation.
func ProtoToRunAlphaJobTerminalCondition(p *alphapb.RunAlphaJobTerminalCondition) *alpha.JobTerminalCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.JobTerminalCondition{
		Type:                dcl.StringOrNil(p.GetType()),
		State:               ProtoToRunAlphaJobTerminalConditionStateEnum(p.GetState()),
		Message:             dcl.StringOrNil(p.GetMessage()),
		LastTransitionTime:  dcl.StringOrNil(p.GetLastTransitionTime()),
		Severity:            ProtoToRunAlphaJobTerminalConditionSeverityEnum(p.GetSeverity()),
		Reason:              ProtoToRunAlphaJobTerminalConditionReasonEnum(p.GetReason()),
		InternalReason:      ProtoToRunAlphaJobTerminalConditionInternalReasonEnum(p.GetInternalReason()),
		DomainMappingReason: ProtoToRunAlphaJobTerminalConditionDomainMappingReasonEnum(p.GetDomainMappingReason()),
		RevisionReason:      ProtoToRunAlphaJobTerminalConditionRevisionReasonEnum(p.GetRevisionReason()),
		ExecutionReason:     ProtoToRunAlphaJobTerminalConditionExecutionReasonEnum(p.GetExecutionReason()),
	}
	return obj
}

// ProtoToJobConditions converts a JobConditions object from its proto representation.
func ProtoToRunAlphaJobConditions(p *alphapb.RunAlphaJobConditions) *alpha.JobConditions {
	if p == nil {
		return nil
	}
	obj := &alpha.JobConditions{
		Type:               dcl.StringOrNil(p.GetType()),
		State:              ProtoToRunAlphaJobConditionsStateEnum(p.GetState()),
		Message:            dcl.StringOrNil(p.GetMessage()),
		LastTransitionTime: dcl.StringOrNil(p.GetLastTransitionTime()),
		Severity:           ProtoToRunAlphaJobConditionsSeverityEnum(p.GetSeverity()),
		Reason:             ProtoToRunAlphaJobConditionsReasonEnum(p.GetReason()),
		RevisionReason:     ProtoToRunAlphaJobConditionsRevisionReasonEnum(p.GetRevisionReason()),
		ExecutionReason:    ProtoToRunAlphaJobConditionsExecutionReasonEnum(p.GetExecutionReason()),
	}
	return obj
}

// ProtoToJobLatestSucceededExecution converts a JobLatestSucceededExecution object from its proto representation.
func ProtoToRunAlphaJobLatestSucceededExecution(p *alphapb.RunAlphaJobLatestSucceededExecution) *alpha.JobLatestSucceededExecution {
	if p == nil {
		return nil
	}
	obj := &alpha.JobLatestSucceededExecution{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToJobLatestCreatedExecution converts a JobLatestCreatedExecution object from its proto representation.
func ProtoToRunAlphaJobLatestCreatedExecution(p *alphapb.RunAlphaJobLatestCreatedExecution) *alpha.JobLatestCreatedExecution {
	if p == nil {
		return nil
	}
	obj := &alpha.JobLatestCreatedExecution{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
	}
	return obj
}

// ProtoToJob converts a Job resource from its proto representation.
func ProtoToJob(p *alphapb.RunAlphaJob) *alpha.Job {
	obj := &alpha.Job{
		Name:                     dcl.StringOrNil(p.GetName()),
		Uid:                      dcl.StringOrNil(p.GetUid()),
		Generation:               dcl.Int64OrNil(p.GetGeneration()),
		CreateTime:               dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:               dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:               dcl.StringOrNil(p.GetDeleteTime()),
		ExpireTime:               dcl.StringOrNil(p.GetExpireTime()),
		Creator:                  dcl.StringOrNil(p.GetCreator()),
		LastModifier:             dcl.StringOrNil(p.GetLastModifier()),
		Client:                   dcl.StringOrNil(p.GetClient()),
		ClientVersion:            dcl.StringOrNil(p.GetClientVersion()),
		LaunchStage:              ProtoToRunAlphaJobLaunchStageEnum(p.GetLaunchStage()),
		BinaryAuthorization:      ProtoToRunAlphaJobBinaryAuthorization(p.GetBinaryAuthorization()),
		Template:                 ProtoToRunAlphaJobTemplate(p.GetTemplate()),
		ObservedGeneration:       dcl.Int64OrNil(p.GetObservedGeneration()),
		TerminalCondition:        ProtoToRunAlphaJobTerminalCondition(p.GetTerminalCondition()),
		ExecutionCount:           dcl.Int64OrNil(p.GetExecutionCount()),
		LatestSucceededExecution: ProtoToRunAlphaJobLatestSucceededExecution(p.GetLatestSucceededExecution()),
		LatestCreatedExecution:   ProtoToRunAlphaJobLatestCreatedExecution(p.GetLatestCreatedExecution()),
		Reconciling:              dcl.Bool(p.GetReconciling()),
		Etag:                     dcl.StringOrNil(p.GetEtag()),
		Project:                  dcl.StringOrNil(p.GetProject()),
		Location:                 dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToRunAlphaJobConditions(r))
	}
	return obj
}

// JobLaunchStageEnumToProto converts a JobLaunchStageEnum enum to its proto representation.
func RunAlphaJobLaunchStageEnumToProto(e *alpha.JobLaunchStageEnum) alphapb.RunAlphaJobLaunchStageEnum {
	if e == nil {
		return alphapb.RunAlphaJobLaunchStageEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobLaunchStageEnum_value["JobLaunchStageEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobLaunchStageEnum(v)
	}
	return alphapb.RunAlphaJobLaunchStageEnum(0)
}

// JobTemplateTemplateExecutionEnvironmentEnumToProto converts a JobTemplateTemplateExecutionEnvironmentEnum enum to its proto representation.
func RunAlphaJobTemplateTemplateExecutionEnvironmentEnumToProto(e *alpha.JobTemplateTemplateExecutionEnvironmentEnum) alphapb.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum {
	if e == nil {
		return alphapb.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum_value["JobTemplateTemplateExecutionEnvironmentEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum(v)
	}
	return alphapb.RunAlphaJobTemplateTemplateExecutionEnvironmentEnum(0)
}

// JobTemplateTemplateVPCAccessEgressEnumToProto converts a JobTemplateTemplateVPCAccessEgressEnum enum to its proto representation.
func RunAlphaJobTemplateTemplateVPCAccessEgressEnumToProto(e *alpha.JobTemplateTemplateVPCAccessEgressEnum) alphapb.RunAlphaJobTemplateTemplateVPCAccessEgressEnum {
	if e == nil {
		return alphapb.RunAlphaJobTemplateTemplateVPCAccessEgressEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTemplateTemplateVPCAccessEgressEnum_value["JobTemplateTemplateVPCAccessEgressEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTemplateTemplateVPCAccessEgressEnum(v)
	}
	return alphapb.RunAlphaJobTemplateTemplateVPCAccessEgressEnum(0)
}

// JobTerminalConditionStateEnumToProto converts a JobTerminalConditionStateEnum enum to its proto representation.
func RunAlphaJobTerminalConditionStateEnumToProto(e *alpha.JobTerminalConditionStateEnum) alphapb.RunAlphaJobTerminalConditionStateEnum {
	if e == nil {
		return alphapb.RunAlphaJobTerminalConditionStateEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTerminalConditionStateEnum_value["JobTerminalConditionStateEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTerminalConditionStateEnum(v)
	}
	return alphapb.RunAlphaJobTerminalConditionStateEnum(0)
}

// JobTerminalConditionSeverityEnumToProto converts a JobTerminalConditionSeverityEnum enum to its proto representation.
func RunAlphaJobTerminalConditionSeverityEnumToProto(e *alpha.JobTerminalConditionSeverityEnum) alphapb.RunAlphaJobTerminalConditionSeverityEnum {
	if e == nil {
		return alphapb.RunAlphaJobTerminalConditionSeverityEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTerminalConditionSeverityEnum_value["JobTerminalConditionSeverityEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTerminalConditionSeverityEnum(v)
	}
	return alphapb.RunAlphaJobTerminalConditionSeverityEnum(0)
}

// JobTerminalConditionReasonEnumToProto converts a JobTerminalConditionReasonEnum enum to its proto representation.
func RunAlphaJobTerminalConditionReasonEnumToProto(e *alpha.JobTerminalConditionReasonEnum) alphapb.RunAlphaJobTerminalConditionReasonEnum {
	if e == nil {
		return alphapb.RunAlphaJobTerminalConditionReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTerminalConditionReasonEnum_value["JobTerminalConditionReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTerminalConditionReasonEnum(v)
	}
	return alphapb.RunAlphaJobTerminalConditionReasonEnum(0)
}

// JobTerminalConditionInternalReasonEnumToProto converts a JobTerminalConditionInternalReasonEnum enum to its proto representation.
func RunAlphaJobTerminalConditionInternalReasonEnumToProto(e *alpha.JobTerminalConditionInternalReasonEnum) alphapb.RunAlphaJobTerminalConditionInternalReasonEnum {
	if e == nil {
		return alphapb.RunAlphaJobTerminalConditionInternalReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTerminalConditionInternalReasonEnum_value["JobTerminalConditionInternalReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTerminalConditionInternalReasonEnum(v)
	}
	return alphapb.RunAlphaJobTerminalConditionInternalReasonEnum(0)
}

// JobTerminalConditionDomainMappingReasonEnumToProto converts a JobTerminalConditionDomainMappingReasonEnum enum to its proto representation.
func RunAlphaJobTerminalConditionDomainMappingReasonEnumToProto(e *alpha.JobTerminalConditionDomainMappingReasonEnum) alphapb.RunAlphaJobTerminalConditionDomainMappingReasonEnum {
	if e == nil {
		return alphapb.RunAlphaJobTerminalConditionDomainMappingReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTerminalConditionDomainMappingReasonEnum_value["JobTerminalConditionDomainMappingReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTerminalConditionDomainMappingReasonEnum(v)
	}
	return alphapb.RunAlphaJobTerminalConditionDomainMappingReasonEnum(0)
}

// JobTerminalConditionRevisionReasonEnumToProto converts a JobTerminalConditionRevisionReasonEnum enum to its proto representation.
func RunAlphaJobTerminalConditionRevisionReasonEnumToProto(e *alpha.JobTerminalConditionRevisionReasonEnum) alphapb.RunAlphaJobTerminalConditionRevisionReasonEnum {
	if e == nil {
		return alphapb.RunAlphaJobTerminalConditionRevisionReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTerminalConditionRevisionReasonEnum_value["JobTerminalConditionRevisionReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTerminalConditionRevisionReasonEnum(v)
	}
	return alphapb.RunAlphaJobTerminalConditionRevisionReasonEnum(0)
}

// JobTerminalConditionExecutionReasonEnumToProto converts a JobTerminalConditionExecutionReasonEnum enum to its proto representation.
func RunAlphaJobTerminalConditionExecutionReasonEnumToProto(e *alpha.JobTerminalConditionExecutionReasonEnum) alphapb.RunAlphaJobTerminalConditionExecutionReasonEnum {
	if e == nil {
		return alphapb.RunAlphaJobTerminalConditionExecutionReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobTerminalConditionExecutionReasonEnum_value["JobTerminalConditionExecutionReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobTerminalConditionExecutionReasonEnum(v)
	}
	return alphapb.RunAlphaJobTerminalConditionExecutionReasonEnum(0)
}

// JobConditionsStateEnumToProto converts a JobConditionsStateEnum enum to its proto representation.
func RunAlphaJobConditionsStateEnumToProto(e *alpha.JobConditionsStateEnum) alphapb.RunAlphaJobConditionsStateEnum {
	if e == nil {
		return alphapb.RunAlphaJobConditionsStateEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobConditionsStateEnum_value["JobConditionsStateEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobConditionsStateEnum(v)
	}
	return alphapb.RunAlphaJobConditionsStateEnum(0)
}

// JobConditionsSeverityEnumToProto converts a JobConditionsSeverityEnum enum to its proto representation.
func RunAlphaJobConditionsSeverityEnumToProto(e *alpha.JobConditionsSeverityEnum) alphapb.RunAlphaJobConditionsSeverityEnum {
	if e == nil {
		return alphapb.RunAlphaJobConditionsSeverityEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobConditionsSeverityEnum_value["JobConditionsSeverityEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobConditionsSeverityEnum(v)
	}
	return alphapb.RunAlphaJobConditionsSeverityEnum(0)
}

// JobConditionsReasonEnumToProto converts a JobConditionsReasonEnum enum to its proto representation.
func RunAlphaJobConditionsReasonEnumToProto(e *alpha.JobConditionsReasonEnum) alphapb.RunAlphaJobConditionsReasonEnum {
	if e == nil {
		return alphapb.RunAlphaJobConditionsReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobConditionsReasonEnum_value["JobConditionsReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobConditionsReasonEnum(v)
	}
	return alphapb.RunAlphaJobConditionsReasonEnum(0)
}

// JobConditionsRevisionReasonEnumToProto converts a JobConditionsRevisionReasonEnum enum to its proto representation.
func RunAlphaJobConditionsRevisionReasonEnumToProto(e *alpha.JobConditionsRevisionReasonEnum) alphapb.RunAlphaJobConditionsRevisionReasonEnum {
	if e == nil {
		return alphapb.RunAlphaJobConditionsRevisionReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobConditionsRevisionReasonEnum_value["JobConditionsRevisionReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobConditionsRevisionReasonEnum(v)
	}
	return alphapb.RunAlphaJobConditionsRevisionReasonEnum(0)
}

// JobConditionsExecutionReasonEnumToProto converts a JobConditionsExecutionReasonEnum enum to its proto representation.
func RunAlphaJobConditionsExecutionReasonEnumToProto(e *alpha.JobConditionsExecutionReasonEnum) alphapb.RunAlphaJobConditionsExecutionReasonEnum {
	if e == nil {
		return alphapb.RunAlphaJobConditionsExecutionReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaJobConditionsExecutionReasonEnum_value["JobConditionsExecutionReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaJobConditionsExecutionReasonEnum(v)
	}
	return alphapb.RunAlphaJobConditionsExecutionReasonEnum(0)
}

// JobBinaryAuthorizationToProto converts a JobBinaryAuthorization object to its proto representation.
func RunAlphaJobBinaryAuthorizationToProto(o *alpha.JobBinaryAuthorization) *alphapb.RunAlphaJobBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobBinaryAuthorization{}
	p.SetUseDefault(dcl.ValueOrEmptyBool(o.UseDefault))
	p.SetBreakglassJustification(dcl.ValueOrEmptyString(o.BreakglassJustification))
	return p
}

// JobTemplateToProto converts a JobTemplate object to its proto representation.
func RunAlphaJobTemplateToProto(o *alpha.JobTemplate) *alphapb.RunAlphaJobTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplate{}
	p.SetParallelism(dcl.ValueOrEmptyInt64(o.Parallelism))
	p.SetTaskCount(dcl.ValueOrEmptyInt64(o.TaskCount))
	p.SetTemplate(RunAlphaJobTemplateTemplateToProto(o.Template))
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	mAnnotations := make(map[string]string, len(o.Annotations))
	for k, r := range o.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)
	return p
}

// JobTemplateTemplateToProto converts a JobTemplateTemplate object to its proto representation.
func RunAlphaJobTemplateTemplateToProto(o *alpha.JobTemplateTemplate) *alphapb.RunAlphaJobTemplateTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplate{}
	p.SetMaxRetries(dcl.ValueOrEmptyInt64(o.MaxRetries))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetExecutionEnvironment(RunAlphaJobTemplateTemplateExecutionEnvironmentEnumToProto(o.ExecutionEnvironment))
	p.SetEncryptionKey(dcl.ValueOrEmptyString(o.EncryptionKey))
	p.SetVpcAccess(RunAlphaJobTemplateTemplateVPCAccessToProto(o.VPCAccess))
	sContainers := make([]*alphapb.RunAlphaJobTemplateTemplateContainers, len(o.Containers))
	for i, r := range o.Containers {
		sContainers[i] = RunAlphaJobTemplateTemplateContainersToProto(&r)
	}
	p.SetContainers(sContainers)
	sVolumes := make([]*alphapb.RunAlphaJobTemplateTemplateVolumes, len(o.Volumes))
	for i, r := range o.Volumes {
		sVolumes[i] = RunAlphaJobTemplateTemplateVolumesToProto(&r)
	}
	p.SetVolumes(sVolumes)
	return p
}

// JobTemplateTemplateContainersToProto converts a JobTemplateTemplateContainers object to its proto representation.
func RunAlphaJobTemplateTemplateContainersToProto(o *alpha.JobTemplateTemplateContainers) *alphapb.RunAlphaJobTemplateTemplateContainers {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateContainers{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetResources(RunAlphaJobTemplateTemplateContainersResourcesToProto(o.Resources))
	sCommand := make([]string, len(o.Command))
	for i, r := range o.Command {
		sCommand[i] = r
	}
	p.SetCommand(sCommand)
	sArgs := make([]string, len(o.Args))
	for i, r := range o.Args {
		sArgs[i] = r
	}
	p.SetArgs(sArgs)
	sEnv := make([]*alphapb.RunAlphaJobTemplateTemplateContainersEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = RunAlphaJobTemplateTemplateContainersEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*alphapb.RunAlphaJobTemplateTemplateContainersPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = RunAlphaJobTemplateTemplateContainersPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	sVolumeMounts := make([]*alphapb.RunAlphaJobTemplateTemplateContainersVolumeMounts, len(o.VolumeMounts))
	for i, r := range o.VolumeMounts {
		sVolumeMounts[i] = RunAlphaJobTemplateTemplateContainersVolumeMountsToProto(&r)
	}
	p.SetVolumeMounts(sVolumeMounts)
	return p
}

// JobTemplateTemplateContainersEnvToProto converts a JobTemplateTemplateContainersEnv object to its proto representation.
func RunAlphaJobTemplateTemplateContainersEnvToProto(o *alpha.JobTemplateTemplateContainersEnv) *alphapb.RunAlphaJobTemplateTemplateContainersEnv {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateContainersEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetValueSource(RunAlphaJobTemplateTemplateContainersEnvValueSourceToProto(o.ValueSource))
	return p
}

// JobTemplateTemplateContainersEnvValueSourceToProto converts a JobTemplateTemplateContainersEnvValueSource object to its proto representation.
func RunAlphaJobTemplateTemplateContainersEnvValueSourceToProto(o *alpha.JobTemplateTemplateContainersEnvValueSource) *alphapb.RunAlphaJobTemplateTemplateContainersEnvValueSource {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateContainersEnvValueSource{}
	p.SetSecretKeyRef(RunAlphaJobTemplateTemplateContainersEnvValueSourceSecretKeyRefToProto(o.SecretKeyRef))
	return p
}

// JobTemplateTemplateContainersEnvValueSourceSecretKeyRefToProto converts a JobTemplateTemplateContainersEnvValueSourceSecretKeyRef object to its proto representation.
func RunAlphaJobTemplateTemplateContainersEnvValueSourceSecretKeyRefToProto(o *alpha.JobTemplateTemplateContainersEnvValueSourceSecretKeyRef) *alphapb.RunAlphaJobTemplateTemplateContainersEnvValueSourceSecretKeyRef {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateContainersEnvValueSourceSecretKeyRef{}
	p.SetSecret(dcl.ValueOrEmptyString(o.Secret))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// JobTemplateTemplateContainersResourcesToProto converts a JobTemplateTemplateContainersResources object to its proto representation.
func RunAlphaJobTemplateTemplateContainersResourcesToProto(o *alpha.JobTemplateTemplateContainersResources) *alphapb.RunAlphaJobTemplateTemplateContainersResources {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateContainersResources{}
	p.SetCpuIdle(dcl.ValueOrEmptyBool(o.CpuIdle))
	mLimits := make(map[string]string, len(o.Limits))
	for k, r := range o.Limits {
		mLimits[k] = r
	}
	p.SetLimits(mLimits)
	return p
}

// JobTemplateTemplateContainersPortsToProto converts a JobTemplateTemplateContainersPorts object to its proto representation.
func RunAlphaJobTemplateTemplateContainersPortsToProto(o *alpha.JobTemplateTemplateContainersPorts) *alphapb.RunAlphaJobTemplateTemplateContainersPorts {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateContainersPorts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// JobTemplateTemplateContainersVolumeMountsToProto converts a JobTemplateTemplateContainersVolumeMounts object to its proto representation.
func RunAlphaJobTemplateTemplateContainersVolumeMountsToProto(o *alpha.JobTemplateTemplateContainersVolumeMounts) *alphapb.RunAlphaJobTemplateTemplateContainersVolumeMounts {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateContainersVolumeMounts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetMountPath(dcl.ValueOrEmptyString(o.MountPath))
	return p
}

// JobTemplateTemplateVolumesToProto converts a JobTemplateTemplateVolumes object to its proto representation.
func RunAlphaJobTemplateTemplateVolumesToProto(o *alpha.JobTemplateTemplateVolumes) *alphapb.RunAlphaJobTemplateTemplateVolumes {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateVolumes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetSecret(RunAlphaJobTemplateTemplateVolumesSecretToProto(o.Secret))
	p.SetCloudSqlInstance(RunAlphaJobTemplateTemplateVolumesCloudSqlInstanceToProto(o.CloudSqlInstance))
	return p
}

// JobTemplateTemplateVolumesSecretToProto converts a JobTemplateTemplateVolumesSecret object to its proto representation.
func RunAlphaJobTemplateTemplateVolumesSecretToProto(o *alpha.JobTemplateTemplateVolumesSecret) *alphapb.RunAlphaJobTemplateTemplateVolumesSecret {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateVolumesSecret{}
	p.SetSecret(dcl.ValueOrEmptyString(o.Secret))
	p.SetDefaultMode(dcl.ValueOrEmptyInt64(o.DefaultMode))
	sItems := make([]*alphapb.RunAlphaJobTemplateTemplateVolumesSecretItems, len(o.Items))
	for i, r := range o.Items {
		sItems[i] = RunAlphaJobTemplateTemplateVolumesSecretItemsToProto(&r)
	}
	p.SetItems(sItems)
	return p
}

// JobTemplateTemplateVolumesSecretItemsToProto converts a JobTemplateTemplateVolumesSecretItems object to its proto representation.
func RunAlphaJobTemplateTemplateVolumesSecretItemsToProto(o *alpha.JobTemplateTemplateVolumesSecretItems) *alphapb.RunAlphaJobTemplateTemplateVolumesSecretItems {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateVolumesSecretItems{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetMode(dcl.ValueOrEmptyInt64(o.Mode))
	return p
}

// JobTemplateTemplateVolumesCloudSqlInstanceToProto converts a JobTemplateTemplateVolumesCloudSqlInstance object to its proto representation.
func RunAlphaJobTemplateTemplateVolumesCloudSqlInstanceToProto(o *alpha.JobTemplateTemplateVolumesCloudSqlInstance) *alphapb.RunAlphaJobTemplateTemplateVolumesCloudSqlInstance {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateVolumesCloudSqlInstance{}
	sInstances := make([]string, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = r
	}
	p.SetInstances(sInstances)
	return p
}

// JobTemplateTemplateVPCAccessToProto converts a JobTemplateTemplateVPCAccess object to its proto representation.
func RunAlphaJobTemplateTemplateVPCAccessToProto(o *alpha.JobTemplateTemplateVPCAccess) *alphapb.RunAlphaJobTemplateTemplateVPCAccess {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTemplateTemplateVPCAccess{}
	p.SetConnector(dcl.ValueOrEmptyString(o.Connector))
	p.SetEgress(RunAlphaJobTemplateTemplateVPCAccessEgressEnumToProto(o.Egress))
	return p
}

// JobTerminalConditionToProto converts a JobTerminalCondition object to its proto representation.
func RunAlphaJobTerminalConditionToProto(o *alpha.JobTerminalCondition) *alphapb.RunAlphaJobTerminalCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobTerminalCondition{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetState(RunAlphaJobTerminalConditionStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetLastTransitionTime(dcl.ValueOrEmptyString(o.LastTransitionTime))
	p.SetSeverity(RunAlphaJobTerminalConditionSeverityEnumToProto(o.Severity))
	p.SetReason(RunAlphaJobTerminalConditionReasonEnumToProto(o.Reason))
	p.SetInternalReason(RunAlphaJobTerminalConditionInternalReasonEnumToProto(o.InternalReason))
	p.SetDomainMappingReason(RunAlphaJobTerminalConditionDomainMappingReasonEnumToProto(o.DomainMappingReason))
	p.SetRevisionReason(RunAlphaJobTerminalConditionRevisionReasonEnumToProto(o.RevisionReason))
	p.SetExecutionReason(RunAlphaJobTerminalConditionExecutionReasonEnumToProto(o.ExecutionReason))
	return p
}

// JobConditionsToProto converts a JobConditions object to its proto representation.
func RunAlphaJobConditionsToProto(o *alpha.JobConditions) *alphapb.RunAlphaJobConditions {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobConditions{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetState(RunAlphaJobConditionsStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetLastTransitionTime(dcl.ValueOrEmptyString(o.LastTransitionTime))
	p.SetSeverity(RunAlphaJobConditionsSeverityEnumToProto(o.Severity))
	p.SetReason(RunAlphaJobConditionsReasonEnumToProto(o.Reason))
	p.SetRevisionReason(RunAlphaJobConditionsRevisionReasonEnumToProto(o.RevisionReason))
	p.SetExecutionReason(RunAlphaJobConditionsExecutionReasonEnumToProto(o.ExecutionReason))
	return p
}

// JobLatestSucceededExecutionToProto converts a JobLatestSucceededExecution object to its proto representation.
func RunAlphaJobLatestSucceededExecutionToProto(o *alpha.JobLatestSucceededExecution) *alphapb.RunAlphaJobLatestSucceededExecution {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobLatestSucceededExecution{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// JobLatestCreatedExecutionToProto converts a JobLatestCreatedExecution object to its proto representation.
func RunAlphaJobLatestCreatedExecutionToProto(o *alpha.JobLatestCreatedExecution) *alphapb.RunAlphaJobLatestCreatedExecution {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaJobLatestCreatedExecution{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(o.CreateTime))
	return p
}

// JobToProto converts a Job resource to its proto representation.
func JobToProto(resource *alpha.Job) *alphapb.RunAlphaJob {
	p := &alphapb.RunAlphaJob{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetGeneration(dcl.ValueOrEmptyInt64(resource.Generation))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetExpireTime(dcl.ValueOrEmptyString(resource.ExpireTime))
	p.SetCreator(dcl.ValueOrEmptyString(resource.Creator))
	p.SetLastModifier(dcl.ValueOrEmptyString(resource.LastModifier))
	p.SetClient(dcl.ValueOrEmptyString(resource.Client))
	p.SetClientVersion(dcl.ValueOrEmptyString(resource.ClientVersion))
	p.SetLaunchStage(RunAlphaJobLaunchStageEnumToProto(resource.LaunchStage))
	p.SetBinaryAuthorization(RunAlphaJobBinaryAuthorizationToProto(resource.BinaryAuthorization))
	p.SetTemplate(RunAlphaJobTemplateToProto(resource.Template))
	p.SetObservedGeneration(dcl.ValueOrEmptyInt64(resource.ObservedGeneration))
	p.SetTerminalCondition(RunAlphaJobTerminalConditionToProto(resource.TerminalCondition))
	p.SetExecutionCount(dcl.ValueOrEmptyInt64(resource.ExecutionCount))
	p.SetLatestSucceededExecution(RunAlphaJobLatestSucceededExecutionToProto(resource.LatestSucceededExecution))
	p.SetLatestCreatedExecution(RunAlphaJobLatestCreatedExecutionToProto(resource.LatestCreatedExecution))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)
	sConditions := make([]*alphapb.RunAlphaJobConditions, len(resource.Conditions))
	for i, r := range resource.Conditions {
		sConditions[i] = RunAlphaJobConditionsToProto(&r)
	}
	p.SetConditions(sConditions)

	return p
}

// applyJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) applyJob(ctx context.Context, c *alpha.Client, request *alphapb.ApplyRunAlphaJobRequest) (*alphapb.RunAlphaJob, error) {
	p := ProtoToJob(request.GetResource())
	res, err := c.ApplyJob(ctx, p)
	if err != nil {
		return nil, err
	}
	r := JobToProto(res)
	return r, nil
}

// applyRunAlphaJob handles the gRPC request by passing it to the underlying Job Apply() method.
func (s *JobServer) ApplyRunAlphaJob(ctx context.Context, request *alphapb.ApplyRunAlphaJobRequest) (*alphapb.RunAlphaJob, error) {
	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyJob(ctx, cl, request)
}

// DeleteJob handles the gRPC request by passing it to the underlying Job Delete() method.
func (s *JobServer) DeleteRunAlphaJob(ctx context.Context, request *alphapb.DeleteRunAlphaJobRequest) (*emptypb.Empty, error) {

	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteJob(ctx, ProtoToJob(request.GetResource()))

}

// ListRunAlphaJob handles the gRPC request by passing it to the underlying JobList() method.
func (s *JobServer) ListRunAlphaJob(ctx context.Context, request *alphapb.ListRunAlphaJobRequest) (*alphapb.ListRunAlphaJobResponse, error) {
	cl, err := createConfigJob(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListJob(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.RunAlphaJob
	for _, r := range resources.Items {
		rp := JobToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListRunAlphaJobResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigJob(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
