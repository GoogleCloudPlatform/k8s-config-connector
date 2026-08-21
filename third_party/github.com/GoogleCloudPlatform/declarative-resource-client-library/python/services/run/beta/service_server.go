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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/run/beta/run_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run/beta"
)

// Server implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceIngressEnum converts a ServiceIngressEnum enum from its proto representation.
func ProtoToRunBetaServiceIngressEnum(e betapb.RunBetaServiceIngressEnum) *beta.ServiceIngressEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceIngressEnum_name[int32(e)]; ok {
		e := beta.ServiceIngressEnum(n[len("RunBetaServiceIngressEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLaunchStageEnum converts a ServiceLaunchStageEnum enum from its proto representation.
func ProtoToRunBetaServiceLaunchStageEnum(e betapb.RunBetaServiceLaunchStageEnum) *beta.ServiceLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceLaunchStageEnum_name[int32(e)]; ok {
		e := beta.ServiceLaunchStageEnum(n[len("RunBetaServiceLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTemplateVPCAccessEgressEnum converts a ServiceTemplateVPCAccessEgressEnum enum from its proto representation.
func ProtoToRunBetaServiceTemplateVPCAccessEgressEnum(e betapb.RunBetaServiceTemplateVPCAccessEgressEnum) *beta.ServiceTemplateVPCAccessEgressEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTemplateVPCAccessEgressEnum_name[int32(e)]; ok {
		e := beta.ServiceTemplateVPCAccessEgressEnum(n[len("RunBetaServiceTemplateVPCAccessEgressEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTemplateExecutionEnvironmentEnum converts a ServiceTemplateExecutionEnvironmentEnum enum from its proto representation.
func ProtoToRunBetaServiceTemplateExecutionEnvironmentEnum(e betapb.RunBetaServiceTemplateExecutionEnvironmentEnum) *beta.ServiceTemplateExecutionEnvironmentEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTemplateExecutionEnvironmentEnum_name[int32(e)]; ok {
		e := beta.ServiceTemplateExecutionEnvironmentEnum(n[len("RunBetaServiceTemplateExecutionEnvironmentEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTrafficTypeEnum converts a ServiceTrafficTypeEnum enum from its proto representation.
func ProtoToRunBetaServiceTrafficTypeEnum(e betapb.RunBetaServiceTrafficTypeEnum) *beta.ServiceTrafficTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTrafficTypeEnum_name[int32(e)]; ok {
		e := beta.ServiceTrafficTypeEnum(n[len("RunBetaServiceTrafficTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionStateEnum converts a ServiceTerminalConditionStateEnum enum from its proto representation.
func ProtoToRunBetaServiceTerminalConditionStateEnum(e betapb.RunBetaServiceTerminalConditionStateEnum) *beta.ServiceTerminalConditionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTerminalConditionStateEnum_name[int32(e)]; ok {
		e := beta.ServiceTerminalConditionStateEnum(n[len("RunBetaServiceTerminalConditionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionSeverityEnum converts a ServiceTerminalConditionSeverityEnum enum from its proto representation.
func ProtoToRunBetaServiceTerminalConditionSeverityEnum(e betapb.RunBetaServiceTerminalConditionSeverityEnum) *beta.ServiceTerminalConditionSeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTerminalConditionSeverityEnum_name[int32(e)]; ok {
		e := beta.ServiceTerminalConditionSeverityEnum(n[len("RunBetaServiceTerminalConditionSeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionReasonEnum converts a ServiceTerminalConditionReasonEnum enum from its proto representation.
func ProtoToRunBetaServiceTerminalConditionReasonEnum(e betapb.RunBetaServiceTerminalConditionReasonEnum) *beta.ServiceTerminalConditionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTerminalConditionReasonEnum_name[int32(e)]; ok {
		e := beta.ServiceTerminalConditionReasonEnum(n[len("RunBetaServiceTerminalConditionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionInternalReasonEnum converts a ServiceTerminalConditionInternalReasonEnum enum from its proto representation.
func ProtoToRunBetaServiceTerminalConditionInternalReasonEnum(e betapb.RunBetaServiceTerminalConditionInternalReasonEnum) *beta.ServiceTerminalConditionInternalReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTerminalConditionInternalReasonEnum_name[int32(e)]; ok {
		e := beta.ServiceTerminalConditionInternalReasonEnum(n[len("RunBetaServiceTerminalConditionInternalReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionDomainMappingReasonEnum converts a ServiceTerminalConditionDomainMappingReasonEnum enum from its proto representation.
func ProtoToRunBetaServiceTerminalConditionDomainMappingReasonEnum(e betapb.RunBetaServiceTerminalConditionDomainMappingReasonEnum) *beta.ServiceTerminalConditionDomainMappingReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTerminalConditionDomainMappingReasonEnum_name[int32(e)]; ok {
		e := beta.ServiceTerminalConditionDomainMappingReasonEnum(n[len("RunBetaServiceTerminalConditionDomainMappingReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionRevisionReasonEnum converts a ServiceTerminalConditionRevisionReasonEnum enum from its proto representation.
func ProtoToRunBetaServiceTerminalConditionRevisionReasonEnum(e betapb.RunBetaServiceTerminalConditionRevisionReasonEnum) *beta.ServiceTerminalConditionRevisionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTerminalConditionRevisionReasonEnum_name[int32(e)]; ok {
		e := beta.ServiceTerminalConditionRevisionReasonEnum(n[len("RunBetaServiceTerminalConditionRevisionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionJobReasonEnum converts a ServiceTerminalConditionJobReasonEnum enum from its proto representation.
func ProtoToRunBetaServiceTerminalConditionJobReasonEnum(e betapb.RunBetaServiceTerminalConditionJobReasonEnum) *beta.ServiceTerminalConditionJobReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTerminalConditionJobReasonEnum_name[int32(e)]; ok {
		e := beta.ServiceTerminalConditionJobReasonEnum(n[len("RunBetaServiceTerminalConditionJobReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTrafficStatusesTypeEnum converts a ServiceTrafficStatusesTypeEnum enum from its proto representation.
func ProtoToRunBetaServiceTrafficStatusesTypeEnum(e betapb.RunBetaServiceTrafficStatusesTypeEnum) *beta.ServiceTrafficStatusesTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.RunBetaServiceTrafficStatusesTypeEnum_name[int32(e)]; ok {
		e := beta.ServiceTrafficStatusesTypeEnum(n[len("RunBetaServiceTrafficStatusesTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceBinaryAuthorization converts a ServiceBinaryAuthorization object from its proto representation.
func ProtoToRunBetaServiceBinaryAuthorization(p *betapb.RunBetaServiceBinaryAuthorization) *beta.ServiceBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceBinaryAuthorization{
		UseDefault:              dcl.Bool(p.GetUseDefault()),
		BreakglassJustification: dcl.StringOrNil(p.GetBreakglassJustification()),
	}
	return obj
}

// ProtoToServiceTemplate converts a ServiceTemplate object from its proto representation.
func ProtoToRunBetaServiceTemplate(p *betapb.RunBetaServiceTemplate) *beta.ServiceTemplate {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplate{
		Revision:             dcl.StringOrNil(p.GetRevision()),
		Scaling:              ProtoToRunBetaServiceTemplateScaling(p.GetScaling()),
		VPCAccess:            ProtoToRunBetaServiceTemplateVPCAccess(p.GetVpcAccess()),
		ContainerConcurrency: dcl.Int64OrNil(p.GetContainerConcurrency()),
		Timeout:              dcl.StringOrNil(p.GetTimeout()),
		ServiceAccount:       dcl.StringOrNil(p.GetServiceAccount()),
		Confidential:         dcl.Bool(p.GetConfidential()),
		ExecutionEnvironment: ProtoToRunBetaServiceTemplateExecutionEnvironmentEnum(p.GetExecutionEnvironment()),
	}
	for _, r := range p.GetContainers() {
		obj.Containers = append(obj.Containers, *ProtoToRunBetaServiceTemplateContainers(r))
	}
	for _, r := range p.GetVolumes() {
		obj.Volumes = append(obj.Volumes, *ProtoToRunBetaServiceTemplateVolumes(r))
	}
	return obj
}

// ProtoToServiceTemplateScaling converts a ServiceTemplateScaling object from its proto representation.
func ProtoToRunBetaServiceTemplateScaling(p *betapb.RunBetaServiceTemplateScaling) *beta.ServiceTemplateScaling {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateScaling{
		MinInstanceCount: dcl.Int64OrNil(p.GetMinInstanceCount()),
		MaxInstanceCount: dcl.Int64OrNil(p.GetMaxInstanceCount()),
	}
	return obj
}

// ProtoToServiceTemplateVPCAccess converts a ServiceTemplateVPCAccess object from its proto representation.
func ProtoToRunBetaServiceTemplateVPCAccess(p *betapb.RunBetaServiceTemplateVPCAccess) *beta.ServiceTemplateVPCAccess {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateVPCAccess{
		Connector: dcl.StringOrNil(p.GetConnector()),
		Egress:    ProtoToRunBetaServiceTemplateVPCAccessEgressEnum(p.GetEgress()),
	}
	return obj
}

// ProtoToServiceTemplateContainers converts a ServiceTemplateContainers object from its proto representation.
func ProtoToRunBetaServiceTemplateContainers(p *betapb.RunBetaServiceTemplateContainers) *beta.ServiceTemplateContainers {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateContainers{
		Name:      dcl.StringOrNil(p.GetName()),
		Image:     dcl.StringOrNil(p.GetImage()),
		Resources: ProtoToRunBetaServiceTemplateContainersResources(p.GetResources()),
	}
	for _, r := range p.GetCommand() {
		obj.Command = append(obj.Command, r)
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetEnv() {
		obj.Env = append(obj.Env, *ProtoToRunBetaServiceTemplateContainersEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToRunBetaServiceTemplateContainersPorts(r))
	}
	for _, r := range p.GetVolumeMounts() {
		obj.VolumeMounts = append(obj.VolumeMounts, *ProtoToRunBetaServiceTemplateContainersVolumeMounts(r))
	}
	return obj
}

// ProtoToServiceTemplateContainersEnv converts a ServiceTemplateContainersEnv object from its proto representation.
func ProtoToRunBetaServiceTemplateContainersEnv(p *betapb.RunBetaServiceTemplateContainersEnv) *beta.ServiceTemplateContainersEnv {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateContainersEnv{
		Name:        dcl.StringOrNil(p.GetName()),
		Value:       dcl.StringOrNil(p.GetValue()),
		ValueSource: ProtoToRunBetaServiceTemplateContainersEnvValueSource(p.GetValueSource()),
	}
	return obj
}

// ProtoToServiceTemplateContainersEnvValueSource converts a ServiceTemplateContainersEnvValueSource object from its proto representation.
func ProtoToRunBetaServiceTemplateContainersEnvValueSource(p *betapb.RunBetaServiceTemplateContainersEnvValueSource) *beta.ServiceTemplateContainersEnvValueSource {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateContainersEnvValueSource{
		SecretKeyRef: ProtoToRunBetaServiceTemplateContainersEnvValueSourceSecretKeyRef(p.GetSecretKeyRef()),
	}
	return obj
}

// ProtoToServiceTemplateContainersEnvValueSourceSecretKeyRef converts a ServiceTemplateContainersEnvValueSourceSecretKeyRef object from its proto representation.
func ProtoToRunBetaServiceTemplateContainersEnvValueSourceSecretKeyRef(p *betapb.RunBetaServiceTemplateContainersEnvValueSourceSecretKeyRef) *beta.ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateContainersEnvValueSourceSecretKeyRef{
		Secret:  dcl.StringOrNil(p.GetSecret()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToServiceTemplateContainersResources converts a ServiceTemplateContainersResources object from its proto representation.
func ProtoToRunBetaServiceTemplateContainersResources(p *betapb.RunBetaServiceTemplateContainersResources) *beta.ServiceTemplateContainersResources {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateContainersResources{
		CpuIdle: dcl.Bool(p.GetCpuIdle()),
	}
	return obj
}

// ProtoToServiceTemplateContainersPorts converts a ServiceTemplateContainersPorts object from its proto representation.
func ProtoToRunBetaServiceTemplateContainersPorts(p *betapb.RunBetaServiceTemplateContainersPorts) *beta.ServiceTemplateContainersPorts {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateContainersPorts{
		Name:          dcl.StringOrNil(p.GetName()),
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToServiceTemplateContainersVolumeMounts converts a ServiceTemplateContainersVolumeMounts object from its proto representation.
func ProtoToRunBetaServiceTemplateContainersVolumeMounts(p *betapb.RunBetaServiceTemplateContainersVolumeMounts) *beta.ServiceTemplateContainersVolumeMounts {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateContainersVolumeMounts{
		Name:      dcl.StringOrNil(p.GetName()),
		MountPath: dcl.StringOrNil(p.GetMountPath()),
	}
	return obj
}

// ProtoToServiceTemplateVolumes converts a ServiceTemplateVolumes object from its proto representation.
func ProtoToRunBetaServiceTemplateVolumes(p *betapb.RunBetaServiceTemplateVolumes) *beta.ServiceTemplateVolumes {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateVolumes{
		Name:             dcl.StringOrNil(p.GetName()),
		Secret:           ProtoToRunBetaServiceTemplateVolumesSecret(p.GetSecret()),
		CloudSqlInstance: ProtoToRunBetaServiceTemplateVolumesCloudSqlInstance(p.GetCloudSqlInstance()),
	}
	return obj
}

// ProtoToServiceTemplateVolumesSecret converts a ServiceTemplateVolumesSecret object from its proto representation.
func ProtoToRunBetaServiceTemplateVolumesSecret(p *betapb.RunBetaServiceTemplateVolumesSecret) *beta.ServiceTemplateVolumesSecret {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateVolumesSecret{
		Secret:      dcl.StringOrNil(p.GetSecret()),
		DefaultMode: dcl.Int64OrNil(p.GetDefaultMode()),
	}
	for _, r := range p.GetItems() {
		obj.Items = append(obj.Items, *ProtoToRunBetaServiceTemplateVolumesSecretItems(r))
	}
	return obj
}

// ProtoToServiceTemplateVolumesSecretItems converts a ServiceTemplateVolumesSecretItems object from its proto representation.
func ProtoToRunBetaServiceTemplateVolumesSecretItems(p *betapb.RunBetaServiceTemplateVolumesSecretItems) *beta.ServiceTemplateVolumesSecretItems {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateVolumesSecretItems{
		Path:    dcl.StringOrNil(p.GetPath()),
		Version: dcl.StringOrNil(p.GetVersion()),
		Mode:    dcl.Int64OrNil(p.GetMode()),
	}
	return obj
}

// ProtoToServiceTemplateVolumesCloudSqlInstance converts a ServiceTemplateVolumesCloudSqlInstance object from its proto representation.
func ProtoToRunBetaServiceTemplateVolumesCloudSqlInstance(p *betapb.RunBetaServiceTemplateVolumesCloudSqlInstance) *beta.ServiceTemplateVolumesCloudSqlInstance {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTemplateVolumesCloudSqlInstance{}
	for _, r := range p.GetConnections() {
		obj.Connections = append(obj.Connections, r)
	}
	return obj
}

// ProtoToServiceTraffic converts a ServiceTraffic object from its proto representation.
func ProtoToRunBetaServiceTraffic(p *betapb.RunBetaServiceTraffic) *beta.ServiceTraffic {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTraffic{
		Type:     ProtoToRunBetaServiceTrafficTypeEnum(p.GetType()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Percent:  dcl.Int64OrNil(p.GetPercent()),
		Tag:      dcl.StringOrNil(p.GetTag()),
	}
	return obj
}

// ProtoToServiceTerminalCondition converts a ServiceTerminalCondition object from its proto representation.
func ProtoToRunBetaServiceTerminalCondition(p *betapb.RunBetaServiceTerminalCondition) *beta.ServiceTerminalCondition {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTerminalCondition{
		Type:                dcl.StringOrNil(p.GetType()),
		State:               ProtoToRunBetaServiceTerminalConditionStateEnum(p.GetState()),
		Message:             dcl.StringOrNil(p.GetMessage()),
		LastTransitionTime:  dcl.StringOrNil(p.GetLastTransitionTime()),
		Severity:            ProtoToRunBetaServiceTerminalConditionSeverityEnum(p.GetSeverity()),
		Reason:              ProtoToRunBetaServiceTerminalConditionReasonEnum(p.GetReason()),
		InternalReason:      ProtoToRunBetaServiceTerminalConditionInternalReasonEnum(p.GetInternalReason()),
		DomainMappingReason: ProtoToRunBetaServiceTerminalConditionDomainMappingReasonEnum(p.GetDomainMappingReason()),
		RevisionReason:      ProtoToRunBetaServiceTerminalConditionRevisionReasonEnum(p.GetRevisionReason()),
		JobReason:           ProtoToRunBetaServiceTerminalConditionJobReasonEnum(p.GetJobReason()),
	}
	return obj
}

// ProtoToServiceTrafficStatuses converts a ServiceTrafficStatuses object from its proto representation.
func ProtoToRunBetaServiceTrafficStatuses(p *betapb.RunBetaServiceTrafficStatuses) *beta.ServiceTrafficStatuses {
	if p == nil {
		return nil
	}
	obj := &beta.ServiceTrafficStatuses{
		Type:     ProtoToRunBetaServiceTrafficStatusesTypeEnum(p.GetType()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Percent:  dcl.Int64OrNil(p.GetPercent()),
		Tag:      dcl.StringOrNil(p.GetTag()),
		Uri:      dcl.StringOrNil(p.GetUri()),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *betapb.RunBetaService) *beta.Service {
	obj := &beta.Service{
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
		Ingress:               ProtoToRunBetaServiceIngressEnum(p.GetIngress()),
		LaunchStage:           ProtoToRunBetaServiceLaunchStageEnum(p.GetLaunchStage()),
		BinaryAuthorization:   ProtoToRunBetaServiceBinaryAuthorization(p.GetBinaryAuthorization()),
		Template:              ProtoToRunBetaServiceTemplate(p.GetTemplate()),
		TerminalCondition:     ProtoToRunBetaServiceTerminalCondition(p.GetTerminalCondition()),
		LatestReadyRevision:   dcl.StringOrNil(p.GetLatestReadyRevision()),
		LatestCreatedRevision: dcl.StringOrNil(p.GetLatestCreatedRevision()),
		Uri:                   dcl.StringOrNil(p.GetUri()),
		Reconciling:           dcl.Bool(p.GetReconciling()),
		Etag:                  dcl.StringOrNil(p.GetEtag()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetTraffic() {
		obj.Traffic = append(obj.Traffic, *ProtoToRunBetaServiceTraffic(r))
	}
	for _, r := range p.GetTrafficStatuses() {
		obj.TrafficStatuses = append(obj.TrafficStatuses, *ProtoToRunBetaServiceTrafficStatuses(r))
	}
	return obj
}

// ServiceIngressEnumToProto converts a ServiceIngressEnum enum to its proto representation.
func RunBetaServiceIngressEnumToProto(e *beta.ServiceIngressEnum) betapb.RunBetaServiceIngressEnum {
	if e == nil {
		return betapb.RunBetaServiceIngressEnum(0)
	}
	if v, ok := betapb.RunBetaServiceIngressEnum_value["ServiceIngressEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceIngressEnum(v)
	}
	return betapb.RunBetaServiceIngressEnum(0)
}

// ServiceLaunchStageEnumToProto converts a ServiceLaunchStageEnum enum to its proto representation.
func RunBetaServiceLaunchStageEnumToProto(e *beta.ServiceLaunchStageEnum) betapb.RunBetaServiceLaunchStageEnum {
	if e == nil {
		return betapb.RunBetaServiceLaunchStageEnum(0)
	}
	if v, ok := betapb.RunBetaServiceLaunchStageEnum_value["ServiceLaunchStageEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceLaunchStageEnum(v)
	}
	return betapb.RunBetaServiceLaunchStageEnum(0)
}

// ServiceTemplateVPCAccessEgressEnumToProto converts a ServiceTemplateVPCAccessEgressEnum enum to its proto representation.
func RunBetaServiceTemplateVPCAccessEgressEnumToProto(e *beta.ServiceTemplateVPCAccessEgressEnum) betapb.RunBetaServiceTemplateVPCAccessEgressEnum {
	if e == nil {
		return betapb.RunBetaServiceTemplateVPCAccessEgressEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTemplateVPCAccessEgressEnum_value["ServiceTemplateVPCAccessEgressEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTemplateVPCAccessEgressEnum(v)
	}
	return betapb.RunBetaServiceTemplateVPCAccessEgressEnum(0)
}

// ServiceTemplateExecutionEnvironmentEnumToProto converts a ServiceTemplateExecutionEnvironmentEnum enum to its proto representation.
func RunBetaServiceTemplateExecutionEnvironmentEnumToProto(e *beta.ServiceTemplateExecutionEnvironmentEnum) betapb.RunBetaServiceTemplateExecutionEnvironmentEnum {
	if e == nil {
		return betapb.RunBetaServiceTemplateExecutionEnvironmentEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTemplateExecutionEnvironmentEnum_value["ServiceTemplateExecutionEnvironmentEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTemplateExecutionEnvironmentEnum(v)
	}
	return betapb.RunBetaServiceTemplateExecutionEnvironmentEnum(0)
}

// ServiceTrafficTypeEnumToProto converts a ServiceTrafficTypeEnum enum to its proto representation.
func RunBetaServiceTrafficTypeEnumToProto(e *beta.ServiceTrafficTypeEnum) betapb.RunBetaServiceTrafficTypeEnum {
	if e == nil {
		return betapb.RunBetaServiceTrafficTypeEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTrafficTypeEnum_value["ServiceTrafficTypeEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTrafficTypeEnum(v)
	}
	return betapb.RunBetaServiceTrafficTypeEnum(0)
}

// ServiceTerminalConditionStateEnumToProto converts a ServiceTerminalConditionStateEnum enum to its proto representation.
func RunBetaServiceTerminalConditionStateEnumToProto(e *beta.ServiceTerminalConditionStateEnum) betapb.RunBetaServiceTerminalConditionStateEnum {
	if e == nil {
		return betapb.RunBetaServiceTerminalConditionStateEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTerminalConditionStateEnum_value["ServiceTerminalConditionStateEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTerminalConditionStateEnum(v)
	}
	return betapb.RunBetaServiceTerminalConditionStateEnum(0)
}

// ServiceTerminalConditionSeverityEnumToProto converts a ServiceTerminalConditionSeverityEnum enum to its proto representation.
func RunBetaServiceTerminalConditionSeverityEnumToProto(e *beta.ServiceTerminalConditionSeverityEnum) betapb.RunBetaServiceTerminalConditionSeverityEnum {
	if e == nil {
		return betapb.RunBetaServiceTerminalConditionSeverityEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTerminalConditionSeverityEnum_value["ServiceTerminalConditionSeverityEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTerminalConditionSeverityEnum(v)
	}
	return betapb.RunBetaServiceTerminalConditionSeverityEnum(0)
}

// ServiceTerminalConditionReasonEnumToProto converts a ServiceTerminalConditionReasonEnum enum to its proto representation.
func RunBetaServiceTerminalConditionReasonEnumToProto(e *beta.ServiceTerminalConditionReasonEnum) betapb.RunBetaServiceTerminalConditionReasonEnum {
	if e == nil {
		return betapb.RunBetaServiceTerminalConditionReasonEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTerminalConditionReasonEnum_value["ServiceTerminalConditionReasonEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTerminalConditionReasonEnum(v)
	}
	return betapb.RunBetaServiceTerminalConditionReasonEnum(0)
}

// ServiceTerminalConditionInternalReasonEnumToProto converts a ServiceTerminalConditionInternalReasonEnum enum to its proto representation.
func RunBetaServiceTerminalConditionInternalReasonEnumToProto(e *beta.ServiceTerminalConditionInternalReasonEnum) betapb.RunBetaServiceTerminalConditionInternalReasonEnum {
	if e == nil {
		return betapb.RunBetaServiceTerminalConditionInternalReasonEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTerminalConditionInternalReasonEnum_value["ServiceTerminalConditionInternalReasonEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTerminalConditionInternalReasonEnum(v)
	}
	return betapb.RunBetaServiceTerminalConditionInternalReasonEnum(0)
}

// ServiceTerminalConditionDomainMappingReasonEnumToProto converts a ServiceTerminalConditionDomainMappingReasonEnum enum to its proto representation.
func RunBetaServiceTerminalConditionDomainMappingReasonEnumToProto(e *beta.ServiceTerminalConditionDomainMappingReasonEnum) betapb.RunBetaServiceTerminalConditionDomainMappingReasonEnum {
	if e == nil {
		return betapb.RunBetaServiceTerminalConditionDomainMappingReasonEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTerminalConditionDomainMappingReasonEnum_value["ServiceTerminalConditionDomainMappingReasonEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTerminalConditionDomainMappingReasonEnum(v)
	}
	return betapb.RunBetaServiceTerminalConditionDomainMappingReasonEnum(0)
}

// ServiceTerminalConditionRevisionReasonEnumToProto converts a ServiceTerminalConditionRevisionReasonEnum enum to its proto representation.
func RunBetaServiceTerminalConditionRevisionReasonEnumToProto(e *beta.ServiceTerminalConditionRevisionReasonEnum) betapb.RunBetaServiceTerminalConditionRevisionReasonEnum {
	if e == nil {
		return betapb.RunBetaServiceTerminalConditionRevisionReasonEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTerminalConditionRevisionReasonEnum_value["ServiceTerminalConditionRevisionReasonEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTerminalConditionRevisionReasonEnum(v)
	}
	return betapb.RunBetaServiceTerminalConditionRevisionReasonEnum(0)
}

// ServiceTerminalConditionJobReasonEnumToProto converts a ServiceTerminalConditionJobReasonEnum enum to its proto representation.
func RunBetaServiceTerminalConditionJobReasonEnumToProto(e *beta.ServiceTerminalConditionJobReasonEnum) betapb.RunBetaServiceTerminalConditionJobReasonEnum {
	if e == nil {
		return betapb.RunBetaServiceTerminalConditionJobReasonEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTerminalConditionJobReasonEnum_value["ServiceTerminalConditionJobReasonEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTerminalConditionJobReasonEnum(v)
	}
	return betapb.RunBetaServiceTerminalConditionJobReasonEnum(0)
}

// ServiceTrafficStatusesTypeEnumToProto converts a ServiceTrafficStatusesTypeEnum enum to its proto representation.
func RunBetaServiceTrafficStatusesTypeEnumToProto(e *beta.ServiceTrafficStatusesTypeEnum) betapb.RunBetaServiceTrafficStatusesTypeEnum {
	if e == nil {
		return betapb.RunBetaServiceTrafficStatusesTypeEnum(0)
	}
	if v, ok := betapb.RunBetaServiceTrafficStatusesTypeEnum_value["ServiceTrafficStatusesTypeEnum"+string(*e)]; ok {
		return betapb.RunBetaServiceTrafficStatusesTypeEnum(v)
	}
	return betapb.RunBetaServiceTrafficStatusesTypeEnum(0)
}

// ServiceBinaryAuthorizationToProto converts a ServiceBinaryAuthorization object to its proto representation.
func RunBetaServiceBinaryAuthorizationToProto(o *beta.ServiceBinaryAuthorization) *betapb.RunBetaServiceBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceBinaryAuthorization{}
	p.SetUseDefault(dcl.ValueOrEmptyBool(o.UseDefault))
	p.SetBreakglassJustification(dcl.ValueOrEmptyString(o.BreakglassJustification))
	return p
}

// ServiceTemplateToProto converts a ServiceTemplate object to its proto representation.
func RunBetaServiceTemplateToProto(o *beta.ServiceTemplate) *betapb.RunBetaServiceTemplate {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplate{}
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetScaling(RunBetaServiceTemplateScalingToProto(o.Scaling))
	p.SetVpcAccess(RunBetaServiceTemplateVPCAccessToProto(o.VPCAccess))
	p.SetContainerConcurrency(dcl.ValueOrEmptyInt64(o.ContainerConcurrency))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetConfidential(dcl.ValueOrEmptyBool(o.Confidential))
	p.SetExecutionEnvironment(RunBetaServiceTemplateExecutionEnvironmentEnumToProto(o.ExecutionEnvironment))
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
	sContainers := make([]*betapb.RunBetaServiceTemplateContainers, len(o.Containers))
	for i, r := range o.Containers {
		sContainers[i] = RunBetaServiceTemplateContainersToProto(&r)
	}
	p.SetContainers(sContainers)
	sVolumes := make([]*betapb.RunBetaServiceTemplateVolumes, len(o.Volumes))
	for i, r := range o.Volumes {
		sVolumes[i] = RunBetaServiceTemplateVolumesToProto(&r)
	}
	p.SetVolumes(sVolumes)
	return p
}

// ServiceTemplateScalingToProto converts a ServiceTemplateScaling object to its proto representation.
func RunBetaServiceTemplateScalingToProto(o *beta.ServiceTemplateScaling) *betapb.RunBetaServiceTemplateScaling {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateScaling{}
	p.SetMinInstanceCount(dcl.ValueOrEmptyInt64(o.MinInstanceCount))
	p.SetMaxInstanceCount(dcl.ValueOrEmptyInt64(o.MaxInstanceCount))
	return p
}

// ServiceTemplateVPCAccessToProto converts a ServiceTemplateVPCAccess object to its proto representation.
func RunBetaServiceTemplateVPCAccessToProto(o *beta.ServiceTemplateVPCAccess) *betapb.RunBetaServiceTemplateVPCAccess {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateVPCAccess{}
	p.SetConnector(dcl.ValueOrEmptyString(o.Connector))
	p.SetEgress(RunBetaServiceTemplateVPCAccessEgressEnumToProto(o.Egress))
	return p
}

// ServiceTemplateContainersToProto converts a ServiceTemplateContainers object to its proto representation.
func RunBetaServiceTemplateContainersToProto(o *beta.ServiceTemplateContainers) *betapb.RunBetaServiceTemplateContainers {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateContainers{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetResources(RunBetaServiceTemplateContainersResourcesToProto(o.Resources))
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
	sEnv := make([]*betapb.RunBetaServiceTemplateContainersEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = RunBetaServiceTemplateContainersEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*betapb.RunBetaServiceTemplateContainersPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = RunBetaServiceTemplateContainersPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	sVolumeMounts := make([]*betapb.RunBetaServiceTemplateContainersVolumeMounts, len(o.VolumeMounts))
	for i, r := range o.VolumeMounts {
		sVolumeMounts[i] = RunBetaServiceTemplateContainersVolumeMountsToProto(&r)
	}
	p.SetVolumeMounts(sVolumeMounts)
	return p
}

// ServiceTemplateContainersEnvToProto converts a ServiceTemplateContainersEnv object to its proto representation.
func RunBetaServiceTemplateContainersEnvToProto(o *beta.ServiceTemplateContainersEnv) *betapb.RunBetaServiceTemplateContainersEnv {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateContainersEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetValueSource(RunBetaServiceTemplateContainersEnvValueSourceToProto(o.ValueSource))
	return p
}

// ServiceTemplateContainersEnvValueSourceToProto converts a ServiceTemplateContainersEnvValueSource object to its proto representation.
func RunBetaServiceTemplateContainersEnvValueSourceToProto(o *beta.ServiceTemplateContainersEnvValueSource) *betapb.RunBetaServiceTemplateContainersEnvValueSource {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateContainersEnvValueSource{}
	p.SetSecretKeyRef(RunBetaServiceTemplateContainersEnvValueSourceSecretKeyRefToProto(o.SecretKeyRef))
	return p
}

// ServiceTemplateContainersEnvValueSourceSecretKeyRefToProto converts a ServiceTemplateContainersEnvValueSourceSecretKeyRef object to its proto representation.
func RunBetaServiceTemplateContainersEnvValueSourceSecretKeyRefToProto(o *beta.ServiceTemplateContainersEnvValueSourceSecretKeyRef) *betapb.RunBetaServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	p.SetSecret(dcl.ValueOrEmptyString(o.Secret))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// ServiceTemplateContainersResourcesToProto converts a ServiceTemplateContainersResources object to its proto representation.
func RunBetaServiceTemplateContainersResourcesToProto(o *beta.ServiceTemplateContainersResources) *betapb.RunBetaServiceTemplateContainersResources {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateContainersResources{}
	p.SetCpuIdle(dcl.ValueOrEmptyBool(o.CpuIdle))
	mLimits := make(map[string]string, len(o.Limits))
	for k, r := range o.Limits {
		mLimits[k] = r
	}
	p.SetLimits(mLimits)
	return p
}

// ServiceTemplateContainersPortsToProto converts a ServiceTemplateContainersPorts object to its proto representation.
func RunBetaServiceTemplateContainersPortsToProto(o *beta.ServiceTemplateContainersPorts) *betapb.RunBetaServiceTemplateContainersPorts {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateContainersPorts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ServiceTemplateContainersVolumeMountsToProto converts a ServiceTemplateContainersVolumeMounts object to its proto representation.
func RunBetaServiceTemplateContainersVolumeMountsToProto(o *beta.ServiceTemplateContainersVolumeMounts) *betapb.RunBetaServiceTemplateContainersVolumeMounts {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateContainersVolumeMounts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetMountPath(dcl.ValueOrEmptyString(o.MountPath))
	return p
}

// ServiceTemplateVolumesToProto converts a ServiceTemplateVolumes object to its proto representation.
func RunBetaServiceTemplateVolumesToProto(o *beta.ServiceTemplateVolumes) *betapb.RunBetaServiceTemplateVolumes {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateVolumes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetSecret(RunBetaServiceTemplateVolumesSecretToProto(o.Secret))
	p.SetCloudSqlInstance(RunBetaServiceTemplateVolumesCloudSqlInstanceToProto(o.CloudSqlInstance))
	return p
}

// ServiceTemplateVolumesSecretToProto converts a ServiceTemplateVolumesSecret object to its proto representation.
func RunBetaServiceTemplateVolumesSecretToProto(o *beta.ServiceTemplateVolumesSecret) *betapb.RunBetaServiceTemplateVolumesSecret {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateVolumesSecret{}
	p.SetSecret(dcl.ValueOrEmptyString(o.Secret))
	p.SetDefaultMode(dcl.ValueOrEmptyInt64(o.DefaultMode))
	sItems := make([]*betapb.RunBetaServiceTemplateVolumesSecretItems, len(o.Items))
	for i, r := range o.Items {
		sItems[i] = RunBetaServiceTemplateVolumesSecretItemsToProto(&r)
	}
	p.SetItems(sItems)
	return p
}

// ServiceTemplateVolumesSecretItemsToProto converts a ServiceTemplateVolumesSecretItems object to its proto representation.
func RunBetaServiceTemplateVolumesSecretItemsToProto(o *beta.ServiceTemplateVolumesSecretItems) *betapb.RunBetaServiceTemplateVolumesSecretItems {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateVolumesSecretItems{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetMode(dcl.ValueOrEmptyInt64(o.Mode))
	return p
}

// ServiceTemplateVolumesCloudSqlInstanceToProto converts a ServiceTemplateVolumesCloudSqlInstance object to its proto representation.
func RunBetaServiceTemplateVolumesCloudSqlInstanceToProto(o *beta.ServiceTemplateVolumesCloudSqlInstance) *betapb.RunBetaServiceTemplateVolumesCloudSqlInstance {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTemplateVolumesCloudSqlInstance{}
	sConnections := make([]string, len(o.Connections))
	for i, r := range o.Connections {
		sConnections[i] = r
	}
	p.SetConnections(sConnections)
	return p
}

// ServiceTrafficToProto converts a ServiceTraffic object to its proto representation.
func RunBetaServiceTrafficToProto(o *beta.ServiceTraffic) *betapb.RunBetaServiceTraffic {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTraffic{}
	p.SetType(RunBetaServiceTrafficTypeEnumToProto(o.Type))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetTag(dcl.ValueOrEmptyString(o.Tag))
	return p
}

// ServiceTerminalConditionToProto converts a ServiceTerminalCondition object to its proto representation.
func RunBetaServiceTerminalConditionToProto(o *beta.ServiceTerminalCondition) *betapb.RunBetaServiceTerminalCondition {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTerminalCondition{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetState(RunBetaServiceTerminalConditionStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetLastTransitionTime(dcl.ValueOrEmptyString(o.LastTransitionTime))
	p.SetSeverity(RunBetaServiceTerminalConditionSeverityEnumToProto(o.Severity))
	p.SetReason(RunBetaServiceTerminalConditionReasonEnumToProto(o.Reason))
	p.SetInternalReason(RunBetaServiceTerminalConditionInternalReasonEnumToProto(o.InternalReason))
	p.SetDomainMappingReason(RunBetaServiceTerminalConditionDomainMappingReasonEnumToProto(o.DomainMappingReason))
	p.SetRevisionReason(RunBetaServiceTerminalConditionRevisionReasonEnumToProto(o.RevisionReason))
	p.SetJobReason(RunBetaServiceTerminalConditionJobReasonEnumToProto(o.JobReason))
	return p
}

// ServiceTrafficStatusesToProto converts a ServiceTrafficStatuses object to its proto representation.
func RunBetaServiceTrafficStatusesToProto(o *beta.ServiceTrafficStatuses) *betapb.RunBetaServiceTrafficStatuses {
	if o == nil {
		return nil
	}
	p := &betapb.RunBetaServiceTrafficStatuses{}
	p.SetType(RunBetaServiceTrafficStatusesTypeEnumToProto(o.Type))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetTag(dcl.ValueOrEmptyString(o.Tag))
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *beta.Service) *betapb.RunBetaService {
	p := &betapb.RunBetaService{}
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
	p.SetIngress(RunBetaServiceIngressEnumToProto(resource.Ingress))
	p.SetLaunchStage(RunBetaServiceLaunchStageEnumToProto(resource.LaunchStage))
	p.SetBinaryAuthorization(RunBetaServiceBinaryAuthorizationToProto(resource.BinaryAuthorization))
	p.SetTemplate(RunBetaServiceTemplateToProto(resource.Template))
	p.SetTerminalCondition(RunBetaServiceTerminalConditionToProto(resource.TerminalCondition))
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
	sTraffic := make([]*betapb.RunBetaServiceTraffic, len(resource.Traffic))
	for i, r := range resource.Traffic {
		sTraffic[i] = RunBetaServiceTrafficToProto(&r)
	}
	p.SetTraffic(sTraffic)
	sTrafficStatuses := make([]*betapb.RunBetaServiceTrafficStatuses, len(resource.TrafficStatuses))
	for i, r := range resource.TrafficStatuses {
		sTrafficStatuses[i] = RunBetaServiceTrafficStatusesToProto(&r)
	}
	p.SetTrafficStatuses(sTrafficStatuses)

	return p
}

// applyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *beta.Client, request *betapb.ApplyRunBetaServiceRequest) (*betapb.RunBetaService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// applyRunBetaService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyRunBetaService(ctx context.Context, request *betapb.ApplyRunBetaServiceRequest) (*betapb.RunBetaService, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteRunBetaService(ctx context.Context, request *betapb.DeleteRunBetaServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListRunBetaService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListRunBetaService(ctx context.Context, request *betapb.ListRunBetaServiceRequest) (*betapb.ListRunBetaServiceResponse, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.RunBetaService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListRunBetaServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
