// Copyright 2022 Google LLC. All Rights Reserved.
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
	runpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/run/run_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run"
)

// Server implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceIngressEnum converts a ServiceIngressEnum enum from its proto representation.
func ProtoToRunServiceIngressEnum(e runpb.RunServiceIngressEnum) *run.ServiceIngressEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceIngressEnum_name[int32(e)]; ok {
		e := run.ServiceIngressEnum(n[len("RunServiceIngressEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLaunchStageEnum converts a ServiceLaunchStageEnum enum from its proto representation.
func ProtoToRunServiceLaunchStageEnum(e runpb.RunServiceLaunchStageEnum) *run.ServiceLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceLaunchStageEnum_name[int32(e)]; ok {
		e := run.ServiceLaunchStageEnum(n[len("RunServiceLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTemplateVPCAccessEgressEnum converts a ServiceTemplateVPCAccessEgressEnum enum from its proto representation.
func ProtoToRunServiceTemplateVPCAccessEgressEnum(e runpb.RunServiceTemplateVPCAccessEgressEnum) *run.ServiceTemplateVPCAccessEgressEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTemplateVPCAccessEgressEnum_name[int32(e)]; ok {
		e := run.ServiceTemplateVPCAccessEgressEnum(n[len("RunServiceTemplateVPCAccessEgressEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTemplateExecutionEnvironmentEnum converts a ServiceTemplateExecutionEnvironmentEnum enum from its proto representation.
func ProtoToRunServiceTemplateExecutionEnvironmentEnum(e runpb.RunServiceTemplateExecutionEnvironmentEnum) *run.ServiceTemplateExecutionEnvironmentEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTemplateExecutionEnvironmentEnum_name[int32(e)]; ok {
		e := run.ServiceTemplateExecutionEnvironmentEnum(n[len("RunServiceTemplateExecutionEnvironmentEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTrafficTypeEnum converts a ServiceTrafficTypeEnum enum from its proto representation.
func ProtoToRunServiceTrafficTypeEnum(e runpb.RunServiceTrafficTypeEnum) *run.ServiceTrafficTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTrafficTypeEnum_name[int32(e)]; ok {
		e := run.ServiceTrafficTypeEnum(n[len("RunServiceTrafficTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionStateEnum converts a ServiceTerminalConditionStateEnum enum from its proto representation.
func ProtoToRunServiceTerminalConditionStateEnum(e runpb.RunServiceTerminalConditionStateEnum) *run.ServiceTerminalConditionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTerminalConditionStateEnum_name[int32(e)]; ok {
		e := run.ServiceTerminalConditionStateEnum(n[len("RunServiceTerminalConditionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionSeverityEnum converts a ServiceTerminalConditionSeverityEnum enum from its proto representation.
func ProtoToRunServiceTerminalConditionSeverityEnum(e runpb.RunServiceTerminalConditionSeverityEnum) *run.ServiceTerminalConditionSeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTerminalConditionSeverityEnum_name[int32(e)]; ok {
		e := run.ServiceTerminalConditionSeverityEnum(n[len("RunServiceTerminalConditionSeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionReasonEnum converts a ServiceTerminalConditionReasonEnum enum from its proto representation.
func ProtoToRunServiceTerminalConditionReasonEnum(e runpb.RunServiceTerminalConditionReasonEnum) *run.ServiceTerminalConditionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTerminalConditionReasonEnum_name[int32(e)]; ok {
		e := run.ServiceTerminalConditionReasonEnum(n[len("RunServiceTerminalConditionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionInternalReasonEnum converts a ServiceTerminalConditionInternalReasonEnum enum from its proto representation.
func ProtoToRunServiceTerminalConditionInternalReasonEnum(e runpb.RunServiceTerminalConditionInternalReasonEnum) *run.ServiceTerminalConditionInternalReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTerminalConditionInternalReasonEnum_name[int32(e)]; ok {
		e := run.ServiceTerminalConditionInternalReasonEnum(n[len("RunServiceTerminalConditionInternalReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionDomainMappingReasonEnum converts a ServiceTerminalConditionDomainMappingReasonEnum enum from its proto representation.
func ProtoToRunServiceTerminalConditionDomainMappingReasonEnum(e runpb.RunServiceTerminalConditionDomainMappingReasonEnum) *run.ServiceTerminalConditionDomainMappingReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTerminalConditionDomainMappingReasonEnum_name[int32(e)]; ok {
		e := run.ServiceTerminalConditionDomainMappingReasonEnum(n[len("RunServiceTerminalConditionDomainMappingReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionRevisionReasonEnum converts a ServiceTerminalConditionRevisionReasonEnum enum from its proto representation.
func ProtoToRunServiceTerminalConditionRevisionReasonEnum(e runpb.RunServiceTerminalConditionRevisionReasonEnum) *run.ServiceTerminalConditionRevisionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTerminalConditionRevisionReasonEnum_name[int32(e)]; ok {
		e := run.ServiceTerminalConditionRevisionReasonEnum(n[len("RunServiceTerminalConditionRevisionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionJobReasonEnum converts a ServiceTerminalConditionJobReasonEnum enum from its proto representation.
func ProtoToRunServiceTerminalConditionJobReasonEnum(e runpb.RunServiceTerminalConditionJobReasonEnum) *run.ServiceTerminalConditionJobReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTerminalConditionJobReasonEnum_name[int32(e)]; ok {
		e := run.ServiceTerminalConditionJobReasonEnum(n[len("RunServiceTerminalConditionJobReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTrafficStatusesTypeEnum converts a ServiceTrafficStatusesTypeEnum enum from its proto representation.
func ProtoToRunServiceTrafficStatusesTypeEnum(e runpb.RunServiceTrafficStatusesTypeEnum) *run.ServiceTrafficStatusesTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := runpb.RunServiceTrafficStatusesTypeEnum_name[int32(e)]; ok {
		e := run.ServiceTrafficStatusesTypeEnum(n[len("RunServiceTrafficStatusesTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceBinaryAuthorization converts a ServiceBinaryAuthorization object from its proto representation.
func ProtoToRunServiceBinaryAuthorization(p *runpb.RunServiceBinaryAuthorization) *run.ServiceBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &run.ServiceBinaryAuthorization{
		UseDefault:              dcl.Bool(p.GetUseDefault()),
		BreakglassJustification: dcl.StringOrNil(p.GetBreakglassJustification()),
	}
	return obj
}

// ProtoToServiceTemplate converts a ServiceTemplate object from its proto representation.
func ProtoToRunServiceTemplate(p *runpb.RunServiceTemplate) *run.ServiceTemplate {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplate{
		Revision:             dcl.StringOrNil(p.GetRevision()),
		Scaling:              ProtoToRunServiceTemplateScaling(p.GetScaling()),
		VPCAccess:            ProtoToRunServiceTemplateVPCAccess(p.GetVpcAccess()),
		ContainerConcurrency: dcl.Int64OrNil(p.GetContainerConcurrency()),
		Timeout:              dcl.StringOrNil(p.GetTimeout()),
		ServiceAccount:       dcl.StringOrNil(p.GetServiceAccount()),
		Confidential:         dcl.Bool(p.GetConfidential()),
		ExecutionEnvironment: ProtoToRunServiceTemplateExecutionEnvironmentEnum(p.GetExecutionEnvironment()),
	}
	for _, r := range p.GetContainers() {
		obj.Containers = append(obj.Containers, *ProtoToRunServiceTemplateContainers(r))
	}
	for _, r := range p.GetVolumes() {
		obj.Volumes = append(obj.Volumes, *ProtoToRunServiceTemplateVolumes(r))
	}
	return obj
}

// ProtoToServiceTemplateScaling converts a ServiceTemplateScaling object from its proto representation.
func ProtoToRunServiceTemplateScaling(p *runpb.RunServiceTemplateScaling) *run.ServiceTemplateScaling {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateScaling{
		MinInstanceCount: dcl.Int64OrNil(p.GetMinInstanceCount()),
		MaxInstanceCount: dcl.Int64OrNil(p.GetMaxInstanceCount()),
	}
	return obj
}

// ProtoToServiceTemplateVPCAccess converts a ServiceTemplateVPCAccess object from its proto representation.
func ProtoToRunServiceTemplateVPCAccess(p *runpb.RunServiceTemplateVPCAccess) *run.ServiceTemplateVPCAccess {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateVPCAccess{
		Connector: dcl.StringOrNil(p.GetConnector()),
		Egress:    ProtoToRunServiceTemplateVPCAccessEgressEnum(p.GetEgress()),
	}
	return obj
}

// ProtoToServiceTemplateContainers converts a ServiceTemplateContainers object from its proto representation.
func ProtoToRunServiceTemplateContainers(p *runpb.RunServiceTemplateContainers) *run.ServiceTemplateContainers {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateContainers{
		Name:      dcl.StringOrNil(p.GetName()),
		Image:     dcl.StringOrNil(p.GetImage()),
		Resources: ProtoToRunServiceTemplateContainersResources(p.GetResources()),
	}
	for _, r := range p.GetCommand() {
		obj.Command = append(obj.Command, r)
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetEnv() {
		obj.Env = append(obj.Env, *ProtoToRunServiceTemplateContainersEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToRunServiceTemplateContainersPorts(r))
	}
	for _, r := range p.GetVolumeMounts() {
		obj.VolumeMounts = append(obj.VolumeMounts, *ProtoToRunServiceTemplateContainersVolumeMounts(r))
	}
	return obj
}

// ProtoToServiceTemplateContainersEnv converts a ServiceTemplateContainersEnv object from its proto representation.
func ProtoToRunServiceTemplateContainersEnv(p *runpb.RunServiceTemplateContainersEnv) *run.ServiceTemplateContainersEnv {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateContainersEnv{
		Name:        dcl.StringOrNil(p.GetName()),
		Value:       dcl.StringOrNil(p.GetValue()),
		ValueSource: ProtoToRunServiceTemplateContainersEnvValueSource(p.GetValueSource()),
	}
	return obj
}

// ProtoToServiceTemplateContainersEnvValueSource converts a ServiceTemplateContainersEnvValueSource object from its proto representation.
func ProtoToRunServiceTemplateContainersEnvValueSource(p *runpb.RunServiceTemplateContainersEnvValueSource) *run.ServiceTemplateContainersEnvValueSource {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateContainersEnvValueSource{
		SecretKeyRef: ProtoToRunServiceTemplateContainersEnvValueSourceSecretKeyRef(p.GetSecretKeyRef()),
	}
	return obj
}

// ProtoToServiceTemplateContainersEnvValueSourceSecretKeyRef converts a ServiceTemplateContainersEnvValueSourceSecretKeyRef object from its proto representation.
func ProtoToRunServiceTemplateContainersEnvValueSourceSecretKeyRef(p *runpb.RunServiceTemplateContainersEnvValueSourceSecretKeyRef) *run.ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateContainersEnvValueSourceSecretKeyRef{
		Secret:  dcl.StringOrNil(p.GetSecret()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToServiceTemplateContainersResources converts a ServiceTemplateContainersResources object from its proto representation.
func ProtoToRunServiceTemplateContainersResources(p *runpb.RunServiceTemplateContainersResources) *run.ServiceTemplateContainersResources {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateContainersResources{
		CpuIdle: dcl.Bool(p.GetCpuIdle()),
	}
	return obj
}

// ProtoToServiceTemplateContainersPorts converts a ServiceTemplateContainersPorts object from its proto representation.
func ProtoToRunServiceTemplateContainersPorts(p *runpb.RunServiceTemplateContainersPorts) *run.ServiceTemplateContainersPorts {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateContainersPorts{
		Name:          dcl.StringOrNil(p.GetName()),
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToServiceTemplateContainersVolumeMounts converts a ServiceTemplateContainersVolumeMounts object from its proto representation.
func ProtoToRunServiceTemplateContainersVolumeMounts(p *runpb.RunServiceTemplateContainersVolumeMounts) *run.ServiceTemplateContainersVolumeMounts {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateContainersVolumeMounts{
		Name:      dcl.StringOrNil(p.GetName()),
		MountPath: dcl.StringOrNil(p.GetMountPath()),
	}
	return obj
}

// ProtoToServiceTemplateVolumes converts a ServiceTemplateVolumes object from its proto representation.
func ProtoToRunServiceTemplateVolumes(p *runpb.RunServiceTemplateVolumes) *run.ServiceTemplateVolumes {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateVolumes{
		Name:             dcl.StringOrNil(p.GetName()),
		Secret:           ProtoToRunServiceTemplateVolumesSecret(p.GetSecret()),
		CloudSqlInstance: ProtoToRunServiceTemplateVolumesCloudSqlInstance(p.GetCloudSqlInstance()),
	}
	return obj
}

// ProtoToServiceTemplateVolumesSecret converts a ServiceTemplateVolumesSecret object from its proto representation.
func ProtoToRunServiceTemplateVolumesSecret(p *runpb.RunServiceTemplateVolumesSecret) *run.ServiceTemplateVolumesSecret {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateVolumesSecret{
		Secret:      dcl.StringOrNil(p.GetSecret()),
		DefaultMode: dcl.Int64OrNil(p.GetDefaultMode()),
	}
	for _, r := range p.GetItems() {
		obj.Items = append(obj.Items, *ProtoToRunServiceTemplateVolumesSecretItems(r))
	}
	return obj
}

// ProtoToServiceTemplateVolumesSecretItems converts a ServiceTemplateVolumesSecretItems object from its proto representation.
func ProtoToRunServiceTemplateVolumesSecretItems(p *runpb.RunServiceTemplateVolumesSecretItems) *run.ServiceTemplateVolumesSecretItems {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateVolumesSecretItems{
		Path:    dcl.StringOrNil(p.GetPath()),
		Version: dcl.StringOrNil(p.GetVersion()),
		Mode:    dcl.Int64OrNil(p.GetMode()),
	}
	return obj
}

// ProtoToServiceTemplateVolumesCloudSqlInstance converts a ServiceTemplateVolumesCloudSqlInstance object from its proto representation.
func ProtoToRunServiceTemplateVolumesCloudSqlInstance(p *runpb.RunServiceTemplateVolumesCloudSqlInstance) *run.ServiceTemplateVolumesCloudSqlInstance {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTemplateVolumesCloudSqlInstance{}
	for _, r := range p.GetConnections() {
		obj.Connections = append(obj.Connections, r)
	}
	return obj
}

// ProtoToServiceTraffic converts a ServiceTraffic object from its proto representation.
func ProtoToRunServiceTraffic(p *runpb.RunServiceTraffic) *run.ServiceTraffic {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTraffic{
		Type:     ProtoToRunServiceTrafficTypeEnum(p.GetType()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Percent:  dcl.Int64OrNil(p.GetPercent()),
		Tag:      dcl.StringOrNil(p.GetTag()),
	}
	return obj
}

// ProtoToServiceTerminalCondition converts a ServiceTerminalCondition object from its proto representation.
func ProtoToRunServiceTerminalCondition(p *runpb.RunServiceTerminalCondition) *run.ServiceTerminalCondition {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTerminalCondition{
		Type:                dcl.StringOrNil(p.GetType()),
		State:               ProtoToRunServiceTerminalConditionStateEnum(p.GetState()),
		Message:             dcl.StringOrNil(p.GetMessage()),
		LastTransitionTime:  dcl.StringOrNil(p.GetLastTransitionTime()),
		Severity:            ProtoToRunServiceTerminalConditionSeverityEnum(p.GetSeverity()),
		Reason:              ProtoToRunServiceTerminalConditionReasonEnum(p.GetReason()),
		InternalReason:      ProtoToRunServiceTerminalConditionInternalReasonEnum(p.GetInternalReason()),
		DomainMappingReason: ProtoToRunServiceTerminalConditionDomainMappingReasonEnum(p.GetDomainMappingReason()),
		RevisionReason:      ProtoToRunServiceTerminalConditionRevisionReasonEnum(p.GetRevisionReason()),
		JobReason:           ProtoToRunServiceTerminalConditionJobReasonEnum(p.GetJobReason()),
	}
	return obj
}

// ProtoToServiceTrafficStatuses converts a ServiceTrafficStatuses object from its proto representation.
func ProtoToRunServiceTrafficStatuses(p *runpb.RunServiceTrafficStatuses) *run.ServiceTrafficStatuses {
	if p == nil {
		return nil
	}
	obj := &run.ServiceTrafficStatuses{
		Type:     ProtoToRunServiceTrafficStatusesTypeEnum(p.GetType()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Percent:  dcl.Int64OrNil(p.GetPercent()),
		Tag:      dcl.StringOrNil(p.GetTag()),
		Uri:      dcl.StringOrNil(p.GetUri()),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *runpb.RunService) *run.Service {
	obj := &run.Service{
		Name:                  dcl.StringOrNil(p.GetName()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		Uid:                   dcl.StringOrNil(p.GetUid()),
		Generation:            dcl.Int64OrNil(p.GetGeneration()),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:            dcl.StringOrNil(p.GetDeleteTime()),
		ExpireTime:            dcl.StringOrNil(p.GetExpireTime()),
		Creator:               dcl.StringOrNil(p.GetCreator()),
		LastModifier:          dcl.StringOrNil(p.GetLastModifier()),
		Client:                dcl.StringOrNil(p.GetClient()),
		ClientVersion:         dcl.StringOrNil(p.GetClientVersion()),
		Ingress:               ProtoToRunServiceIngressEnum(p.GetIngress()),
		LaunchStage:           ProtoToRunServiceLaunchStageEnum(p.GetLaunchStage()),
		BinaryAuthorization:   ProtoToRunServiceBinaryAuthorization(p.GetBinaryAuthorization()),
		Template:              ProtoToRunServiceTemplate(p.GetTemplate()),
		TerminalCondition:     ProtoToRunServiceTerminalCondition(p.GetTerminalCondition()),
		LatestReadyRevision:   dcl.StringOrNil(p.GetLatestReadyRevision()),
		LatestCreatedRevision: dcl.StringOrNil(p.GetLatestCreatedRevision()),
		Uri:                   dcl.StringOrNil(p.GetUri()),
		Reconciling:           dcl.Bool(p.GetReconciling()),
		Etag:                  dcl.StringOrNil(p.GetEtag()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetTraffic() {
		obj.Traffic = append(obj.Traffic, *ProtoToRunServiceTraffic(r))
	}
	for _, r := range p.GetTrafficStatuses() {
		obj.TrafficStatuses = append(obj.TrafficStatuses, *ProtoToRunServiceTrafficStatuses(r))
	}
	return obj
}

// ServiceIngressEnumToProto converts a ServiceIngressEnum enum to its proto representation.
func RunServiceIngressEnumToProto(e *run.ServiceIngressEnum) runpb.RunServiceIngressEnum {
	if e == nil {
		return runpb.RunServiceIngressEnum(0)
	}
	if v, ok := runpb.RunServiceIngressEnum_value["ServiceIngressEnum"+string(*e)]; ok {
		return runpb.RunServiceIngressEnum(v)
	}
	return runpb.RunServiceIngressEnum(0)
}

// ServiceLaunchStageEnumToProto converts a ServiceLaunchStageEnum enum to its proto representation.
func RunServiceLaunchStageEnumToProto(e *run.ServiceLaunchStageEnum) runpb.RunServiceLaunchStageEnum {
	if e == nil {
		return runpb.RunServiceLaunchStageEnum(0)
	}
	if v, ok := runpb.RunServiceLaunchStageEnum_value["ServiceLaunchStageEnum"+string(*e)]; ok {
		return runpb.RunServiceLaunchStageEnum(v)
	}
	return runpb.RunServiceLaunchStageEnum(0)
}

// ServiceTemplateVPCAccessEgressEnumToProto converts a ServiceTemplateVPCAccessEgressEnum enum to its proto representation.
func RunServiceTemplateVPCAccessEgressEnumToProto(e *run.ServiceTemplateVPCAccessEgressEnum) runpb.RunServiceTemplateVPCAccessEgressEnum {
	if e == nil {
		return runpb.RunServiceTemplateVPCAccessEgressEnum(0)
	}
	if v, ok := runpb.RunServiceTemplateVPCAccessEgressEnum_value["ServiceTemplateVPCAccessEgressEnum"+string(*e)]; ok {
		return runpb.RunServiceTemplateVPCAccessEgressEnum(v)
	}
	return runpb.RunServiceTemplateVPCAccessEgressEnum(0)
}

// ServiceTemplateExecutionEnvironmentEnumToProto converts a ServiceTemplateExecutionEnvironmentEnum enum to its proto representation.
func RunServiceTemplateExecutionEnvironmentEnumToProto(e *run.ServiceTemplateExecutionEnvironmentEnum) runpb.RunServiceTemplateExecutionEnvironmentEnum {
	if e == nil {
		return runpb.RunServiceTemplateExecutionEnvironmentEnum(0)
	}
	if v, ok := runpb.RunServiceTemplateExecutionEnvironmentEnum_value["ServiceTemplateExecutionEnvironmentEnum"+string(*e)]; ok {
		return runpb.RunServiceTemplateExecutionEnvironmentEnum(v)
	}
	return runpb.RunServiceTemplateExecutionEnvironmentEnum(0)
}

// ServiceTrafficTypeEnumToProto converts a ServiceTrafficTypeEnum enum to its proto representation.
func RunServiceTrafficTypeEnumToProto(e *run.ServiceTrafficTypeEnum) runpb.RunServiceTrafficTypeEnum {
	if e == nil {
		return runpb.RunServiceTrafficTypeEnum(0)
	}
	if v, ok := runpb.RunServiceTrafficTypeEnum_value["ServiceTrafficTypeEnum"+string(*e)]; ok {
		return runpb.RunServiceTrafficTypeEnum(v)
	}
	return runpb.RunServiceTrafficTypeEnum(0)
}

// ServiceTerminalConditionStateEnumToProto converts a ServiceTerminalConditionStateEnum enum to its proto representation.
func RunServiceTerminalConditionStateEnumToProto(e *run.ServiceTerminalConditionStateEnum) runpb.RunServiceTerminalConditionStateEnum {
	if e == nil {
		return runpb.RunServiceTerminalConditionStateEnum(0)
	}
	if v, ok := runpb.RunServiceTerminalConditionStateEnum_value["ServiceTerminalConditionStateEnum"+string(*e)]; ok {
		return runpb.RunServiceTerminalConditionStateEnum(v)
	}
	return runpb.RunServiceTerminalConditionStateEnum(0)
}

// ServiceTerminalConditionSeverityEnumToProto converts a ServiceTerminalConditionSeverityEnum enum to its proto representation.
func RunServiceTerminalConditionSeverityEnumToProto(e *run.ServiceTerminalConditionSeverityEnum) runpb.RunServiceTerminalConditionSeverityEnum {
	if e == nil {
		return runpb.RunServiceTerminalConditionSeverityEnum(0)
	}
	if v, ok := runpb.RunServiceTerminalConditionSeverityEnum_value["ServiceTerminalConditionSeverityEnum"+string(*e)]; ok {
		return runpb.RunServiceTerminalConditionSeverityEnum(v)
	}
	return runpb.RunServiceTerminalConditionSeverityEnum(0)
}

// ServiceTerminalConditionReasonEnumToProto converts a ServiceTerminalConditionReasonEnum enum to its proto representation.
func RunServiceTerminalConditionReasonEnumToProto(e *run.ServiceTerminalConditionReasonEnum) runpb.RunServiceTerminalConditionReasonEnum {
	if e == nil {
		return runpb.RunServiceTerminalConditionReasonEnum(0)
	}
	if v, ok := runpb.RunServiceTerminalConditionReasonEnum_value["ServiceTerminalConditionReasonEnum"+string(*e)]; ok {
		return runpb.RunServiceTerminalConditionReasonEnum(v)
	}
	return runpb.RunServiceTerminalConditionReasonEnum(0)
}

// ServiceTerminalConditionInternalReasonEnumToProto converts a ServiceTerminalConditionInternalReasonEnum enum to its proto representation.
func RunServiceTerminalConditionInternalReasonEnumToProto(e *run.ServiceTerminalConditionInternalReasonEnum) runpb.RunServiceTerminalConditionInternalReasonEnum {
	if e == nil {
		return runpb.RunServiceTerminalConditionInternalReasonEnum(0)
	}
	if v, ok := runpb.RunServiceTerminalConditionInternalReasonEnum_value["ServiceTerminalConditionInternalReasonEnum"+string(*e)]; ok {
		return runpb.RunServiceTerminalConditionInternalReasonEnum(v)
	}
	return runpb.RunServiceTerminalConditionInternalReasonEnum(0)
}

// ServiceTerminalConditionDomainMappingReasonEnumToProto converts a ServiceTerminalConditionDomainMappingReasonEnum enum to its proto representation.
func RunServiceTerminalConditionDomainMappingReasonEnumToProto(e *run.ServiceTerminalConditionDomainMappingReasonEnum) runpb.RunServiceTerminalConditionDomainMappingReasonEnum {
	if e == nil {
		return runpb.RunServiceTerminalConditionDomainMappingReasonEnum(0)
	}
	if v, ok := runpb.RunServiceTerminalConditionDomainMappingReasonEnum_value["ServiceTerminalConditionDomainMappingReasonEnum"+string(*e)]; ok {
		return runpb.RunServiceTerminalConditionDomainMappingReasonEnum(v)
	}
	return runpb.RunServiceTerminalConditionDomainMappingReasonEnum(0)
}

// ServiceTerminalConditionRevisionReasonEnumToProto converts a ServiceTerminalConditionRevisionReasonEnum enum to its proto representation.
func RunServiceTerminalConditionRevisionReasonEnumToProto(e *run.ServiceTerminalConditionRevisionReasonEnum) runpb.RunServiceTerminalConditionRevisionReasonEnum {
	if e == nil {
		return runpb.RunServiceTerminalConditionRevisionReasonEnum(0)
	}
	if v, ok := runpb.RunServiceTerminalConditionRevisionReasonEnum_value["ServiceTerminalConditionRevisionReasonEnum"+string(*e)]; ok {
		return runpb.RunServiceTerminalConditionRevisionReasonEnum(v)
	}
	return runpb.RunServiceTerminalConditionRevisionReasonEnum(0)
}

// ServiceTerminalConditionJobReasonEnumToProto converts a ServiceTerminalConditionJobReasonEnum enum to its proto representation.
func RunServiceTerminalConditionJobReasonEnumToProto(e *run.ServiceTerminalConditionJobReasonEnum) runpb.RunServiceTerminalConditionJobReasonEnum {
	if e == nil {
		return runpb.RunServiceTerminalConditionJobReasonEnum(0)
	}
	if v, ok := runpb.RunServiceTerminalConditionJobReasonEnum_value["ServiceTerminalConditionJobReasonEnum"+string(*e)]; ok {
		return runpb.RunServiceTerminalConditionJobReasonEnum(v)
	}
	return runpb.RunServiceTerminalConditionJobReasonEnum(0)
}

// ServiceTrafficStatusesTypeEnumToProto converts a ServiceTrafficStatusesTypeEnum enum to its proto representation.
func RunServiceTrafficStatusesTypeEnumToProto(e *run.ServiceTrafficStatusesTypeEnum) runpb.RunServiceTrafficStatusesTypeEnum {
	if e == nil {
		return runpb.RunServiceTrafficStatusesTypeEnum(0)
	}
	if v, ok := runpb.RunServiceTrafficStatusesTypeEnum_value["ServiceTrafficStatusesTypeEnum"+string(*e)]; ok {
		return runpb.RunServiceTrafficStatusesTypeEnum(v)
	}
	return runpb.RunServiceTrafficStatusesTypeEnum(0)
}

// ServiceBinaryAuthorizationToProto converts a ServiceBinaryAuthorization object to its proto representation.
func RunServiceBinaryAuthorizationToProto(o *run.ServiceBinaryAuthorization) *runpb.RunServiceBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceBinaryAuthorization{}
	p.SetUseDefault(dcl.ValueOrEmptyBool(o.UseDefault))
	p.SetBreakglassJustification(dcl.ValueOrEmptyString(o.BreakglassJustification))
	return p
}

// ServiceTemplateToProto converts a ServiceTemplate object to its proto representation.
func RunServiceTemplateToProto(o *run.ServiceTemplate) *runpb.RunServiceTemplate {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplate{}
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetScaling(RunServiceTemplateScalingToProto(o.Scaling))
	p.SetVpcAccess(RunServiceTemplateVPCAccessToProto(o.VPCAccess))
	p.SetContainerConcurrency(dcl.ValueOrEmptyInt64(o.ContainerConcurrency))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetConfidential(dcl.ValueOrEmptyBool(o.Confidential))
	p.SetExecutionEnvironment(RunServiceTemplateExecutionEnvironmentEnumToProto(o.ExecutionEnvironment))
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
	sContainers := make([]*runpb.RunServiceTemplateContainers, len(o.Containers))
	for i, r := range o.Containers {
		sContainers[i] = RunServiceTemplateContainersToProto(&r)
	}
	p.SetContainers(sContainers)
	sVolumes := make([]*runpb.RunServiceTemplateVolumes, len(o.Volumes))
	for i, r := range o.Volumes {
		sVolumes[i] = RunServiceTemplateVolumesToProto(&r)
	}
	p.SetVolumes(sVolumes)
	return p
}

// ServiceTemplateScalingToProto converts a ServiceTemplateScaling object to its proto representation.
func RunServiceTemplateScalingToProto(o *run.ServiceTemplateScaling) *runpb.RunServiceTemplateScaling {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateScaling{}
	p.SetMinInstanceCount(dcl.ValueOrEmptyInt64(o.MinInstanceCount))
	p.SetMaxInstanceCount(dcl.ValueOrEmptyInt64(o.MaxInstanceCount))
	return p
}

// ServiceTemplateVPCAccessToProto converts a ServiceTemplateVPCAccess object to its proto representation.
func RunServiceTemplateVPCAccessToProto(o *run.ServiceTemplateVPCAccess) *runpb.RunServiceTemplateVPCAccess {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateVPCAccess{}
	p.SetConnector(dcl.ValueOrEmptyString(o.Connector))
	p.SetEgress(RunServiceTemplateVPCAccessEgressEnumToProto(o.Egress))
	return p
}

// ServiceTemplateContainersToProto converts a ServiceTemplateContainers object to its proto representation.
func RunServiceTemplateContainersToProto(o *run.ServiceTemplateContainers) *runpb.RunServiceTemplateContainers {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateContainers{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetResources(RunServiceTemplateContainersResourcesToProto(o.Resources))
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
	sEnv := make([]*runpb.RunServiceTemplateContainersEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = RunServiceTemplateContainersEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*runpb.RunServiceTemplateContainersPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = RunServiceTemplateContainersPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	sVolumeMounts := make([]*runpb.RunServiceTemplateContainersVolumeMounts, len(o.VolumeMounts))
	for i, r := range o.VolumeMounts {
		sVolumeMounts[i] = RunServiceTemplateContainersVolumeMountsToProto(&r)
	}
	p.SetVolumeMounts(sVolumeMounts)
	return p
}

// ServiceTemplateContainersEnvToProto converts a ServiceTemplateContainersEnv object to its proto representation.
func RunServiceTemplateContainersEnvToProto(o *run.ServiceTemplateContainersEnv) *runpb.RunServiceTemplateContainersEnv {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateContainersEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetValueSource(RunServiceTemplateContainersEnvValueSourceToProto(o.ValueSource))
	return p
}

// ServiceTemplateContainersEnvValueSourceToProto converts a ServiceTemplateContainersEnvValueSource object to its proto representation.
func RunServiceTemplateContainersEnvValueSourceToProto(o *run.ServiceTemplateContainersEnvValueSource) *runpb.RunServiceTemplateContainersEnvValueSource {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateContainersEnvValueSource{}
	p.SetSecretKeyRef(RunServiceTemplateContainersEnvValueSourceSecretKeyRefToProto(o.SecretKeyRef))
	return p
}

// ServiceTemplateContainersEnvValueSourceSecretKeyRefToProto converts a ServiceTemplateContainersEnvValueSourceSecretKeyRef object to its proto representation.
func RunServiceTemplateContainersEnvValueSourceSecretKeyRefToProto(o *run.ServiceTemplateContainersEnvValueSourceSecretKeyRef) *runpb.RunServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	p.SetSecret(dcl.ValueOrEmptyString(o.Secret))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// ServiceTemplateContainersResourcesToProto converts a ServiceTemplateContainersResources object to its proto representation.
func RunServiceTemplateContainersResourcesToProto(o *run.ServiceTemplateContainersResources) *runpb.RunServiceTemplateContainersResources {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateContainersResources{}
	p.SetCpuIdle(dcl.ValueOrEmptyBool(o.CpuIdle))
	mLimits := make(map[string]string, len(o.Limits))
	for k, r := range o.Limits {
		mLimits[k] = r
	}
	p.SetLimits(mLimits)
	return p
}

// ServiceTemplateContainersPortsToProto converts a ServiceTemplateContainersPorts object to its proto representation.
func RunServiceTemplateContainersPortsToProto(o *run.ServiceTemplateContainersPorts) *runpb.RunServiceTemplateContainersPorts {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateContainersPorts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ServiceTemplateContainersVolumeMountsToProto converts a ServiceTemplateContainersVolumeMounts object to its proto representation.
func RunServiceTemplateContainersVolumeMountsToProto(o *run.ServiceTemplateContainersVolumeMounts) *runpb.RunServiceTemplateContainersVolumeMounts {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateContainersVolumeMounts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetMountPath(dcl.ValueOrEmptyString(o.MountPath))
	return p
}

// ServiceTemplateVolumesToProto converts a ServiceTemplateVolumes object to its proto representation.
func RunServiceTemplateVolumesToProto(o *run.ServiceTemplateVolumes) *runpb.RunServiceTemplateVolumes {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateVolumes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetSecret(RunServiceTemplateVolumesSecretToProto(o.Secret))
	p.SetCloudSqlInstance(RunServiceTemplateVolumesCloudSqlInstanceToProto(o.CloudSqlInstance))
	return p
}

// ServiceTemplateVolumesSecretToProto converts a ServiceTemplateVolumesSecret object to its proto representation.
func RunServiceTemplateVolumesSecretToProto(o *run.ServiceTemplateVolumesSecret) *runpb.RunServiceTemplateVolumesSecret {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateVolumesSecret{}
	p.SetSecret(dcl.ValueOrEmptyString(o.Secret))
	p.SetDefaultMode(dcl.ValueOrEmptyInt64(o.DefaultMode))
	sItems := make([]*runpb.RunServiceTemplateVolumesSecretItems, len(o.Items))
	for i, r := range o.Items {
		sItems[i] = RunServiceTemplateVolumesSecretItemsToProto(&r)
	}
	p.SetItems(sItems)
	return p
}

// ServiceTemplateVolumesSecretItemsToProto converts a ServiceTemplateVolumesSecretItems object to its proto representation.
func RunServiceTemplateVolumesSecretItemsToProto(o *run.ServiceTemplateVolumesSecretItems) *runpb.RunServiceTemplateVolumesSecretItems {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateVolumesSecretItems{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetMode(dcl.ValueOrEmptyInt64(o.Mode))
	return p
}

// ServiceTemplateVolumesCloudSqlInstanceToProto converts a ServiceTemplateVolumesCloudSqlInstance object to its proto representation.
func RunServiceTemplateVolumesCloudSqlInstanceToProto(o *run.ServiceTemplateVolumesCloudSqlInstance) *runpb.RunServiceTemplateVolumesCloudSqlInstance {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTemplateVolumesCloudSqlInstance{}
	sConnections := make([]string, len(o.Connections))
	for i, r := range o.Connections {
		sConnections[i] = r
	}
	p.SetConnections(sConnections)
	return p
}

// ServiceTrafficToProto converts a ServiceTraffic object to its proto representation.
func RunServiceTrafficToProto(o *run.ServiceTraffic) *runpb.RunServiceTraffic {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTraffic{}
	p.SetType(RunServiceTrafficTypeEnumToProto(o.Type))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetTag(dcl.ValueOrEmptyString(o.Tag))
	return p
}

// ServiceTerminalConditionToProto converts a ServiceTerminalCondition object to its proto representation.
func RunServiceTerminalConditionToProto(o *run.ServiceTerminalCondition) *runpb.RunServiceTerminalCondition {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTerminalCondition{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetState(RunServiceTerminalConditionStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetLastTransitionTime(dcl.ValueOrEmptyString(o.LastTransitionTime))
	p.SetSeverity(RunServiceTerminalConditionSeverityEnumToProto(o.Severity))
	p.SetReason(RunServiceTerminalConditionReasonEnumToProto(o.Reason))
	p.SetInternalReason(RunServiceTerminalConditionInternalReasonEnumToProto(o.InternalReason))
	p.SetDomainMappingReason(RunServiceTerminalConditionDomainMappingReasonEnumToProto(o.DomainMappingReason))
	p.SetRevisionReason(RunServiceTerminalConditionRevisionReasonEnumToProto(o.RevisionReason))
	p.SetJobReason(RunServiceTerminalConditionJobReasonEnumToProto(o.JobReason))
	return p
}

// ServiceTrafficStatusesToProto converts a ServiceTrafficStatuses object to its proto representation.
func RunServiceTrafficStatusesToProto(o *run.ServiceTrafficStatuses) *runpb.RunServiceTrafficStatuses {
	if o == nil {
		return nil
	}
	p := &runpb.RunServiceTrafficStatuses{}
	p.SetType(RunServiceTrafficStatusesTypeEnumToProto(o.Type))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetTag(dcl.ValueOrEmptyString(o.Tag))
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *run.Service) *runpb.RunService {
	p := &runpb.RunService{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
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
	p.SetIngress(RunServiceIngressEnumToProto(resource.Ingress))
	p.SetLaunchStage(RunServiceLaunchStageEnumToProto(resource.LaunchStage))
	p.SetBinaryAuthorization(RunServiceBinaryAuthorizationToProto(resource.BinaryAuthorization))
	p.SetTemplate(RunServiceTemplateToProto(resource.Template))
	p.SetTerminalCondition(RunServiceTerminalConditionToProto(resource.TerminalCondition))
	p.SetLatestReadyRevision(dcl.ValueOrEmptyString(resource.LatestReadyRevision))
	p.SetLatestCreatedRevision(dcl.ValueOrEmptyString(resource.LatestCreatedRevision))
	p.SetUri(dcl.ValueOrEmptyString(resource.Uri))
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
	sTraffic := make([]*runpb.RunServiceTraffic, len(resource.Traffic))
	for i, r := range resource.Traffic {
		sTraffic[i] = RunServiceTrafficToProto(&r)
	}
	p.SetTraffic(sTraffic)
	sTrafficStatuses := make([]*runpb.RunServiceTrafficStatuses, len(resource.TrafficStatuses))
	for i, r := range resource.TrafficStatuses {
		sTrafficStatuses[i] = RunServiceTrafficStatusesToProto(&r)
	}
	p.SetTrafficStatuses(sTrafficStatuses)

	return p
}

// applyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *run.Client, request *runpb.ApplyRunServiceRequest) (*runpb.RunService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// applyRunService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyRunService(ctx context.Context, request *runpb.ApplyRunServiceRequest) (*runpb.RunService, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteRunService(ctx context.Context, request *runpb.DeleteRunServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListRunService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListRunService(ctx context.Context, request *runpb.ListRunServiceRequest) (*runpb.ListRunServiceResponse, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*runpb.RunService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &runpb.ListRunServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*run.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return run.NewClient(conf), nil
}
