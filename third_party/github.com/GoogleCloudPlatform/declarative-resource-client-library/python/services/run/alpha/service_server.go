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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/run/alpha/run_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/run/alpha"
)

// ServiceServer implements the gRPC interface for Service.
type ServiceServer struct{}

// ProtoToServiceIngressEnum converts a ServiceIngressEnum enum from its proto representation.
func ProtoToRunAlphaServiceIngressEnum(e alphapb.RunAlphaServiceIngressEnum) *alpha.ServiceIngressEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceIngressEnum_name[int32(e)]; ok {
		e := alpha.ServiceIngressEnum(n[len("RunAlphaServiceIngressEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceLaunchStageEnum converts a ServiceLaunchStageEnum enum from its proto representation.
func ProtoToRunAlphaServiceLaunchStageEnum(e alphapb.RunAlphaServiceLaunchStageEnum) *alpha.ServiceLaunchStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceLaunchStageEnum_name[int32(e)]; ok {
		e := alpha.ServiceLaunchStageEnum(n[len("RunAlphaServiceLaunchStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTemplateVPCAccessEgressEnum converts a ServiceTemplateVPCAccessEgressEnum enum from its proto representation.
func ProtoToRunAlphaServiceTemplateVPCAccessEgressEnum(e alphapb.RunAlphaServiceTemplateVPCAccessEgressEnum) *alpha.ServiceTemplateVPCAccessEgressEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTemplateVPCAccessEgressEnum_name[int32(e)]; ok {
		e := alpha.ServiceTemplateVPCAccessEgressEnum(n[len("RunAlphaServiceTemplateVPCAccessEgressEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTemplateExecutionEnvironmentEnum converts a ServiceTemplateExecutionEnvironmentEnum enum from its proto representation.
func ProtoToRunAlphaServiceTemplateExecutionEnvironmentEnum(e alphapb.RunAlphaServiceTemplateExecutionEnvironmentEnum) *alpha.ServiceTemplateExecutionEnvironmentEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTemplateExecutionEnvironmentEnum_name[int32(e)]; ok {
		e := alpha.ServiceTemplateExecutionEnvironmentEnum(n[len("RunAlphaServiceTemplateExecutionEnvironmentEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTrafficTypeEnum converts a ServiceTrafficTypeEnum enum from its proto representation.
func ProtoToRunAlphaServiceTrafficTypeEnum(e alphapb.RunAlphaServiceTrafficTypeEnum) *alpha.ServiceTrafficTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTrafficTypeEnum_name[int32(e)]; ok {
		e := alpha.ServiceTrafficTypeEnum(n[len("RunAlphaServiceTrafficTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionStateEnum converts a ServiceTerminalConditionStateEnum enum from its proto representation.
func ProtoToRunAlphaServiceTerminalConditionStateEnum(e alphapb.RunAlphaServiceTerminalConditionStateEnum) *alpha.ServiceTerminalConditionStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTerminalConditionStateEnum_name[int32(e)]; ok {
		e := alpha.ServiceTerminalConditionStateEnum(n[len("RunAlphaServiceTerminalConditionStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionSeverityEnum converts a ServiceTerminalConditionSeverityEnum enum from its proto representation.
func ProtoToRunAlphaServiceTerminalConditionSeverityEnum(e alphapb.RunAlphaServiceTerminalConditionSeverityEnum) *alpha.ServiceTerminalConditionSeverityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTerminalConditionSeverityEnum_name[int32(e)]; ok {
		e := alpha.ServiceTerminalConditionSeverityEnum(n[len("RunAlphaServiceTerminalConditionSeverityEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionReasonEnum converts a ServiceTerminalConditionReasonEnum enum from its proto representation.
func ProtoToRunAlphaServiceTerminalConditionReasonEnum(e alphapb.RunAlphaServiceTerminalConditionReasonEnum) *alpha.ServiceTerminalConditionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTerminalConditionReasonEnum_name[int32(e)]; ok {
		e := alpha.ServiceTerminalConditionReasonEnum(n[len("RunAlphaServiceTerminalConditionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionRevisionReasonEnum converts a ServiceTerminalConditionRevisionReasonEnum enum from its proto representation.
func ProtoToRunAlphaServiceTerminalConditionRevisionReasonEnum(e alphapb.RunAlphaServiceTerminalConditionRevisionReasonEnum) *alpha.ServiceTerminalConditionRevisionReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTerminalConditionRevisionReasonEnum_name[int32(e)]; ok {
		e := alpha.ServiceTerminalConditionRevisionReasonEnum(n[len("RunAlphaServiceTerminalConditionRevisionReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTerminalConditionJobReasonEnum converts a ServiceTerminalConditionJobReasonEnum enum from its proto representation.
func ProtoToRunAlphaServiceTerminalConditionJobReasonEnum(e alphapb.RunAlphaServiceTerminalConditionJobReasonEnum) *alpha.ServiceTerminalConditionJobReasonEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTerminalConditionJobReasonEnum_name[int32(e)]; ok {
		e := alpha.ServiceTerminalConditionJobReasonEnum(n[len("RunAlphaServiceTerminalConditionJobReasonEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceTrafficStatusesTypeEnum converts a ServiceTrafficStatusesTypeEnum enum from its proto representation.
func ProtoToRunAlphaServiceTrafficStatusesTypeEnum(e alphapb.RunAlphaServiceTrafficStatusesTypeEnum) *alpha.ServiceTrafficStatusesTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.RunAlphaServiceTrafficStatusesTypeEnum_name[int32(e)]; ok {
		e := alpha.ServiceTrafficStatusesTypeEnum(n[len("RunAlphaServiceTrafficStatusesTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToServiceBinaryAuthorization converts a ServiceBinaryAuthorization object from its proto representation.
func ProtoToRunAlphaServiceBinaryAuthorization(p *alphapb.RunAlphaServiceBinaryAuthorization) *alpha.ServiceBinaryAuthorization {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceBinaryAuthorization{
		UseDefault:              dcl.Bool(p.GetUseDefault()),
		BreakglassJustification: dcl.StringOrNil(p.GetBreakglassJustification()),
	}
	return obj
}

// ProtoToServiceTemplate converts a ServiceTemplate object from its proto representation.
func ProtoToRunAlphaServiceTemplate(p *alphapb.RunAlphaServiceTemplate) *alpha.ServiceTemplate {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplate{
		Revision:             dcl.StringOrNil(p.GetRevision()),
		Scaling:              ProtoToRunAlphaServiceTemplateScaling(p.GetScaling()),
		VPCAccess:            ProtoToRunAlphaServiceTemplateVPCAccess(p.GetVpcAccess()),
		ContainerConcurrency: dcl.Int64OrNil(p.GetContainerConcurrency()),
		Timeout:              dcl.StringOrNil(p.GetTimeout()),
		ServiceAccount:       dcl.StringOrNil(p.GetServiceAccount()),
		ExecutionEnvironment: ProtoToRunAlphaServiceTemplateExecutionEnvironmentEnum(p.GetExecutionEnvironment()),
	}
	for _, r := range p.GetContainers() {
		obj.Containers = append(obj.Containers, *ProtoToRunAlphaServiceTemplateContainers(r))
	}
	for _, r := range p.GetVolumes() {
		obj.Volumes = append(obj.Volumes, *ProtoToRunAlphaServiceTemplateVolumes(r))
	}
	return obj
}

// ProtoToServiceTemplateScaling converts a ServiceTemplateScaling object from its proto representation.
func ProtoToRunAlphaServiceTemplateScaling(p *alphapb.RunAlphaServiceTemplateScaling) *alpha.ServiceTemplateScaling {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateScaling{
		MinInstanceCount: dcl.Int64OrNil(p.GetMinInstanceCount()),
		MaxInstanceCount: dcl.Int64OrNil(p.GetMaxInstanceCount()),
	}
	return obj
}

// ProtoToServiceTemplateVPCAccess converts a ServiceTemplateVPCAccess object from its proto representation.
func ProtoToRunAlphaServiceTemplateVPCAccess(p *alphapb.RunAlphaServiceTemplateVPCAccess) *alpha.ServiceTemplateVPCAccess {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateVPCAccess{
		Connector: dcl.StringOrNil(p.GetConnector()),
		Egress:    ProtoToRunAlphaServiceTemplateVPCAccessEgressEnum(p.GetEgress()),
	}
	return obj
}

// ProtoToServiceTemplateContainers converts a ServiceTemplateContainers object from its proto representation.
func ProtoToRunAlphaServiceTemplateContainers(p *alphapb.RunAlphaServiceTemplateContainers) *alpha.ServiceTemplateContainers {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateContainers{
		Name:      dcl.StringOrNil(p.GetName()),
		Image:     dcl.StringOrNil(p.GetImage()),
		Resources: ProtoToRunAlphaServiceTemplateContainersResources(p.GetResources()),
	}
	for _, r := range p.GetCommand() {
		obj.Command = append(obj.Command, r)
	}
	for _, r := range p.GetArgs() {
		obj.Args = append(obj.Args, r)
	}
	for _, r := range p.GetEnv() {
		obj.Env = append(obj.Env, *ProtoToRunAlphaServiceTemplateContainersEnv(r))
	}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, *ProtoToRunAlphaServiceTemplateContainersPorts(r))
	}
	for _, r := range p.GetVolumeMounts() {
		obj.VolumeMounts = append(obj.VolumeMounts, *ProtoToRunAlphaServiceTemplateContainersVolumeMounts(r))
	}
	return obj
}

// ProtoToServiceTemplateContainersEnv converts a ServiceTemplateContainersEnv object from its proto representation.
func ProtoToRunAlphaServiceTemplateContainersEnv(p *alphapb.RunAlphaServiceTemplateContainersEnv) *alpha.ServiceTemplateContainersEnv {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateContainersEnv{
		Name:        dcl.StringOrNil(p.GetName()),
		Value:       dcl.StringOrNil(p.GetValue()),
		ValueSource: ProtoToRunAlphaServiceTemplateContainersEnvValueSource(p.GetValueSource()),
	}
	return obj
}

// ProtoToServiceTemplateContainersEnvValueSource converts a ServiceTemplateContainersEnvValueSource object from its proto representation.
func ProtoToRunAlphaServiceTemplateContainersEnvValueSource(p *alphapb.RunAlphaServiceTemplateContainersEnvValueSource) *alpha.ServiceTemplateContainersEnvValueSource {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateContainersEnvValueSource{
		SecretKeyRef: ProtoToRunAlphaServiceTemplateContainersEnvValueSourceSecretKeyRef(p.GetSecretKeyRef()),
	}
	return obj
}

// ProtoToServiceTemplateContainersEnvValueSourceSecretKeyRef converts a ServiceTemplateContainersEnvValueSourceSecretKeyRef object from its proto representation.
func ProtoToRunAlphaServiceTemplateContainersEnvValueSourceSecretKeyRef(p *alphapb.RunAlphaServiceTemplateContainersEnvValueSourceSecretKeyRef) *alpha.ServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateContainersEnvValueSourceSecretKeyRef{
		Secret:  dcl.StringOrNil(p.GetSecret()),
		Version: dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToServiceTemplateContainersResources converts a ServiceTemplateContainersResources object from its proto representation.
func ProtoToRunAlphaServiceTemplateContainersResources(p *alphapb.RunAlphaServiceTemplateContainersResources) *alpha.ServiceTemplateContainersResources {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateContainersResources{
		CpuIdle: dcl.Bool(p.GetCpuIdle()),
	}
	return obj
}

// ProtoToServiceTemplateContainersPorts converts a ServiceTemplateContainersPorts object from its proto representation.
func ProtoToRunAlphaServiceTemplateContainersPorts(p *alphapb.RunAlphaServiceTemplateContainersPorts) *alpha.ServiceTemplateContainersPorts {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateContainersPorts{
		Name:          dcl.StringOrNil(p.GetName()),
		ContainerPort: dcl.Int64OrNil(p.GetContainerPort()),
	}
	return obj
}

// ProtoToServiceTemplateContainersVolumeMounts converts a ServiceTemplateContainersVolumeMounts object from its proto representation.
func ProtoToRunAlphaServiceTemplateContainersVolumeMounts(p *alphapb.RunAlphaServiceTemplateContainersVolumeMounts) *alpha.ServiceTemplateContainersVolumeMounts {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateContainersVolumeMounts{
		Name:      dcl.StringOrNil(p.GetName()),
		MountPath: dcl.StringOrNil(p.GetMountPath()),
	}
	return obj
}

// ProtoToServiceTemplateVolumes converts a ServiceTemplateVolumes object from its proto representation.
func ProtoToRunAlphaServiceTemplateVolumes(p *alphapb.RunAlphaServiceTemplateVolumes) *alpha.ServiceTemplateVolumes {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateVolumes{
		Name:             dcl.StringOrNil(p.GetName()),
		Secret:           ProtoToRunAlphaServiceTemplateVolumesSecret(p.GetSecret()),
		CloudSqlInstance: ProtoToRunAlphaServiceTemplateVolumesCloudSqlInstance(p.GetCloudSqlInstance()),
	}
	return obj
}

// ProtoToServiceTemplateVolumesSecret converts a ServiceTemplateVolumesSecret object from its proto representation.
func ProtoToRunAlphaServiceTemplateVolumesSecret(p *alphapb.RunAlphaServiceTemplateVolumesSecret) *alpha.ServiceTemplateVolumesSecret {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateVolumesSecret{
		Secret:      dcl.StringOrNil(p.GetSecret()),
		DefaultMode: dcl.Int64OrNil(p.GetDefaultMode()),
	}
	for _, r := range p.GetItems() {
		obj.Items = append(obj.Items, *ProtoToRunAlphaServiceTemplateVolumesSecretItems(r))
	}
	return obj
}

// ProtoToServiceTemplateVolumesSecretItems converts a ServiceTemplateVolumesSecretItems object from its proto representation.
func ProtoToRunAlphaServiceTemplateVolumesSecretItems(p *alphapb.RunAlphaServiceTemplateVolumesSecretItems) *alpha.ServiceTemplateVolumesSecretItems {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateVolumesSecretItems{
		Path:    dcl.StringOrNil(p.GetPath()),
		Version: dcl.StringOrNil(p.GetVersion()),
		Mode:    dcl.Int64OrNil(p.GetMode()),
	}
	return obj
}

// ProtoToServiceTemplateVolumesCloudSqlInstance converts a ServiceTemplateVolumesCloudSqlInstance object from its proto representation.
func ProtoToRunAlphaServiceTemplateVolumesCloudSqlInstance(p *alphapb.RunAlphaServiceTemplateVolumesCloudSqlInstance) *alpha.ServiceTemplateVolumesCloudSqlInstance {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTemplateVolumesCloudSqlInstance{}
	for _, r := range p.GetInstances() {
		obj.Instances = append(obj.Instances, r)
	}
	return obj
}

// ProtoToServiceTraffic converts a ServiceTraffic object from its proto representation.
func ProtoToRunAlphaServiceTraffic(p *alphapb.RunAlphaServiceTraffic) *alpha.ServiceTraffic {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTraffic{
		Type:     ProtoToRunAlphaServiceTrafficTypeEnum(p.GetType()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Percent:  dcl.Int64OrNil(p.GetPercent()),
		Tag:      dcl.StringOrNil(p.GetTag()),
	}
	return obj
}

// ProtoToServiceTerminalCondition converts a ServiceTerminalCondition object from its proto representation.
func ProtoToRunAlphaServiceTerminalCondition(p *alphapb.RunAlphaServiceTerminalCondition) *alpha.ServiceTerminalCondition {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTerminalCondition{
		Type:               dcl.StringOrNil(p.GetType()),
		State:              ProtoToRunAlphaServiceTerminalConditionStateEnum(p.GetState()),
		Message:            dcl.StringOrNil(p.GetMessage()),
		LastTransitionTime: dcl.StringOrNil(p.GetLastTransitionTime()),
		Severity:           ProtoToRunAlphaServiceTerminalConditionSeverityEnum(p.GetSeverity()),
		Reason:             ProtoToRunAlphaServiceTerminalConditionReasonEnum(p.GetReason()),
		RevisionReason:     ProtoToRunAlphaServiceTerminalConditionRevisionReasonEnum(p.GetRevisionReason()),
		JobReason:          ProtoToRunAlphaServiceTerminalConditionJobReasonEnum(p.GetJobReason()),
	}
	return obj
}

// ProtoToServiceTrafficStatuses converts a ServiceTrafficStatuses object from its proto representation.
func ProtoToRunAlphaServiceTrafficStatuses(p *alphapb.RunAlphaServiceTrafficStatuses) *alpha.ServiceTrafficStatuses {
	if p == nil {
		return nil
	}
	obj := &alpha.ServiceTrafficStatuses{
		Type:     ProtoToRunAlphaServiceTrafficStatusesTypeEnum(p.GetType()),
		Revision: dcl.StringOrNil(p.GetRevision()),
		Percent:  dcl.Int64OrNil(p.GetPercent()),
		Tag:      dcl.StringOrNil(p.GetTag()),
		Uri:      dcl.StringOrNil(p.GetUri()),
	}
	return obj
}

// ProtoToService converts a Service resource from its proto representation.
func ProtoToService(p *alphapb.RunAlphaService) *alpha.Service {
	obj := &alpha.Service{
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
		Ingress:               ProtoToRunAlphaServiceIngressEnum(p.GetIngress()),
		LaunchStage:           ProtoToRunAlphaServiceLaunchStageEnum(p.GetLaunchStage()),
		BinaryAuthorization:   ProtoToRunAlphaServiceBinaryAuthorization(p.GetBinaryAuthorization()),
		Template:              ProtoToRunAlphaServiceTemplate(p.GetTemplate()),
		TerminalCondition:     ProtoToRunAlphaServiceTerminalCondition(p.GetTerminalCondition()),
		LatestReadyRevision:   dcl.StringOrNil(p.GetLatestReadyRevision()),
		LatestCreatedRevision: dcl.StringOrNil(p.GetLatestCreatedRevision()),
		Uri:                   dcl.StringOrNil(p.GetUri()),
		Reconciling:           dcl.Bool(p.GetReconciling()),
		Etag:                  dcl.StringOrNil(p.GetEtag()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetTraffic() {
		obj.Traffic = append(obj.Traffic, *ProtoToRunAlphaServiceTraffic(r))
	}
	for _, r := range p.GetTrafficStatuses() {
		obj.TrafficStatuses = append(obj.TrafficStatuses, *ProtoToRunAlphaServiceTrafficStatuses(r))
	}
	return obj
}

// ServiceIngressEnumToProto converts a ServiceIngressEnum enum to its proto representation.
func RunAlphaServiceIngressEnumToProto(e *alpha.ServiceIngressEnum) alphapb.RunAlphaServiceIngressEnum {
	if e == nil {
		return alphapb.RunAlphaServiceIngressEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceIngressEnum_value["ServiceIngressEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceIngressEnum(v)
	}
	return alphapb.RunAlphaServiceIngressEnum(0)
}

// ServiceLaunchStageEnumToProto converts a ServiceLaunchStageEnum enum to its proto representation.
func RunAlphaServiceLaunchStageEnumToProto(e *alpha.ServiceLaunchStageEnum) alphapb.RunAlphaServiceLaunchStageEnum {
	if e == nil {
		return alphapb.RunAlphaServiceLaunchStageEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceLaunchStageEnum_value["ServiceLaunchStageEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceLaunchStageEnum(v)
	}
	return alphapb.RunAlphaServiceLaunchStageEnum(0)
}

// ServiceTemplateVPCAccessEgressEnumToProto converts a ServiceTemplateVPCAccessEgressEnum enum to its proto representation.
func RunAlphaServiceTemplateVPCAccessEgressEnumToProto(e *alpha.ServiceTemplateVPCAccessEgressEnum) alphapb.RunAlphaServiceTemplateVPCAccessEgressEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTemplateVPCAccessEgressEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTemplateVPCAccessEgressEnum_value["ServiceTemplateVPCAccessEgressEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTemplateVPCAccessEgressEnum(v)
	}
	return alphapb.RunAlphaServiceTemplateVPCAccessEgressEnum(0)
}

// ServiceTemplateExecutionEnvironmentEnumToProto converts a ServiceTemplateExecutionEnvironmentEnum enum to its proto representation.
func RunAlphaServiceTemplateExecutionEnvironmentEnumToProto(e *alpha.ServiceTemplateExecutionEnvironmentEnum) alphapb.RunAlphaServiceTemplateExecutionEnvironmentEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTemplateExecutionEnvironmentEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTemplateExecutionEnvironmentEnum_value["ServiceTemplateExecutionEnvironmentEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTemplateExecutionEnvironmentEnum(v)
	}
	return alphapb.RunAlphaServiceTemplateExecutionEnvironmentEnum(0)
}

// ServiceTrafficTypeEnumToProto converts a ServiceTrafficTypeEnum enum to its proto representation.
func RunAlphaServiceTrafficTypeEnumToProto(e *alpha.ServiceTrafficTypeEnum) alphapb.RunAlphaServiceTrafficTypeEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTrafficTypeEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTrafficTypeEnum_value["ServiceTrafficTypeEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTrafficTypeEnum(v)
	}
	return alphapb.RunAlphaServiceTrafficTypeEnum(0)
}

// ServiceTerminalConditionStateEnumToProto converts a ServiceTerminalConditionStateEnum enum to its proto representation.
func RunAlphaServiceTerminalConditionStateEnumToProto(e *alpha.ServiceTerminalConditionStateEnum) alphapb.RunAlphaServiceTerminalConditionStateEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTerminalConditionStateEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTerminalConditionStateEnum_value["ServiceTerminalConditionStateEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTerminalConditionStateEnum(v)
	}
	return alphapb.RunAlphaServiceTerminalConditionStateEnum(0)
}

// ServiceTerminalConditionSeverityEnumToProto converts a ServiceTerminalConditionSeverityEnum enum to its proto representation.
func RunAlphaServiceTerminalConditionSeverityEnumToProto(e *alpha.ServiceTerminalConditionSeverityEnum) alphapb.RunAlphaServiceTerminalConditionSeverityEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTerminalConditionSeverityEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTerminalConditionSeverityEnum_value["ServiceTerminalConditionSeverityEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTerminalConditionSeverityEnum(v)
	}
	return alphapb.RunAlphaServiceTerminalConditionSeverityEnum(0)
}

// ServiceTerminalConditionReasonEnumToProto converts a ServiceTerminalConditionReasonEnum enum to its proto representation.
func RunAlphaServiceTerminalConditionReasonEnumToProto(e *alpha.ServiceTerminalConditionReasonEnum) alphapb.RunAlphaServiceTerminalConditionReasonEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTerminalConditionReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTerminalConditionReasonEnum_value["ServiceTerminalConditionReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTerminalConditionReasonEnum(v)
	}
	return alphapb.RunAlphaServiceTerminalConditionReasonEnum(0)
}

// ServiceTerminalConditionRevisionReasonEnumToProto converts a ServiceTerminalConditionRevisionReasonEnum enum to its proto representation.
func RunAlphaServiceTerminalConditionRevisionReasonEnumToProto(e *alpha.ServiceTerminalConditionRevisionReasonEnum) alphapb.RunAlphaServiceTerminalConditionRevisionReasonEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTerminalConditionRevisionReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTerminalConditionRevisionReasonEnum_value["ServiceTerminalConditionRevisionReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTerminalConditionRevisionReasonEnum(v)
	}
	return alphapb.RunAlphaServiceTerminalConditionRevisionReasonEnum(0)
}

// ServiceTerminalConditionJobReasonEnumToProto converts a ServiceTerminalConditionJobReasonEnum enum to its proto representation.
func RunAlphaServiceTerminalConditionJobReasonEnumToProto(e *alpha.ServiceTerminalConditionJobReasonEnum) alphapb.RunAlphaServiceTerminalConditionJobReasonEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTerminalConditionJobReasonEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTerminalConditionJobReasonEnum_value["ServiceTerminalConditionJobReasonEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTerminalConditionJobReasonEnum(v)
	}
	return alphapb.RunAlphaServiceTerminalConditionJobReasonEnum(0)
}

// ServiceTrafficStatusesTypeEnumToProto converts a ServiceTrafficStatusesTypeEnum enum to its proto representation.
func RunAlphaServiceTrafficStatusesTypeEnumToProto(e *alpha.ServiceTrafficStatusesTypeEnum) alphapb.RunAlphaServiceTrafficStatusesTypeEnum {
	if e == nil {
		return alphapb.RunAlphaServiceTrafficStatusesTypeEnum(0)
	}
	if v, ok := alphapb.RunAlphaServiceTrafficStatusesTypeEnum_value["ServiceTrafficStatusesTypeEnum"+string(*e)]; ok {
		return alphapb.RunAlphaServiceTrafficStatusesTypeEnum(v)
	}
	return alphapb.RunAlphaServiceTrafficStatusesTypeEnum(0)
}

// ServiceBinaryAuthorizationToProto converts a ServiceBinaryAuthorization object to its proto representation.
func RunAlphaServiceBinaryAuthorizationToProto(o *alpha.ServiceBinaryAuthorization) *alphapb.RunAlphaServiceBinaryAuthorization {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceBinaryAuthorization{}
	p.SetUseDefault(dcl.ValueOrEmptyBool(o.UseDefault))
	p.SetBreakglassJustification(dcl.ValueOrEmptyString(o.BreakglassJustification))
	return p
}

// ServiceTemplateToProto converts a ServiceTemplate object to its proto representation.
func RunAlphaServiceTemplateToProto(o *alpha.ServiceTemplate) *alphapb.RunAlphaServiceTemplate {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplate{}
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetScaling(RunAlphaServiceTemplateScalingToProto(o.Scaling))
	p.SetVpcAccess(RunAlphaServiceTemplateVPCAccessToProto(o.VPCAccess))
	p.SetContainerConcurrency(dcl.ValueOrEmptyInt64(o.ContainerConcurrency))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetServiceAccount(dcl.ValueOrEmptyString(o.ServiceAccount))
	p.SetExecutionEnvironment(RunAlphaServiceTemplateExecutionEnvironmentEnumToProto(o.ExecutionEnvironment))
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
	sContainers := make([]*alphapb.RunAlphaServiceTemplateContainers, len(o.Containers))
	for i, r := range o.Containers {
		sContainers[i] = RunAlphaServiceTemplateContainersToProto(&r)
	}
	p.SetContainers(sContainers)
	sVolumes := make([]*alphapb.RunAlphaServiceTemplateVolumes, len(o.Volumes))
	for i, r := range o.Volumes {
		sVolumes[i] = RunAlphaServiceTemplateVolumesToProto(&r)
	}
	p.SetVolumes(sVolumes)
	return p
}

// ServiceTemplateScalingToProto converts a ServiceTemplateScaling object to its proto representation.
func RunAlphaServiceTemplateScalingToProto(o *alpha.ServiceTemplateScaling) *alphapb.RunAlphaServiceTemplateScaling {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateScaling{}
	p.SetMinInstanceCount(dcl.ValueOrEmptyInt64(o.MinInstanceCount))
	p.SetMaxInstanceCount(dcl.ValueOrEmptyInt64(o.MaxInstanceCount))
	return p
}

// ServiceTemplateVPCAccessToProto converts a ServiceTemplateVPCAccess object to its proto representation.
func RunAlphaServiceTemplateVPCAccessToProto(o *alpha.ServiceTemplateVPCAccess) *alphapb.RunAlphaServiceTemplateVPCAccess {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateVPCAccess{}
	p.SetConnector(dcl.ValueOrEmptyString(o.Connector))
	p.SetEgress(RunAlphaServiceTemplateVPCAccessEgressEnumToProto(o.Egress))
	return p
}

// ServiceTemplateContainersToProto converts a ServiceTemplateContainers object to its proto representation.
func RunAlphaServiceTemplateContainersToProto(o *alpha.ServiceTemplateContainers) *alphapb.RunAlphaServiceTemplateContainers {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateContainers{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetImage(dcl.ValueOrEmptyString(o.Image))
	p.SetResources(RunAlphaServiceTemplateContainersResourcesToProto(o.Resources))
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
	sEnv := make([]*alphapb.RunAlphaServiceTemplateContainersEnv, len(o.Env))
	for i, r := range o.Env {
		sEnv[i] = RunAlphaServiceTemplateContainersEnvToProto(&r)
	}
	p.SetEnv(sEnv)
	sPorts := make([]*alphapb.RunAlphaServiceTemplateContainersPorts, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = RunAlphaServiceTemplateContainersPortsToProto(&r)
	}
	p.SetPorts(sPorts)
	sVolumeMounts := make([]*alphapb.RunAlphaServiceTemplateContainersVolumeMounts, len(o.VolumeMounts))
	for i, r := range o.VolumeMounts {
		sVolumeMounts[i] = RunAlphaServiceTemplateContainersVolumeMountsToProto(&r)
	}
	p.SetVolumeMounts(sVolumeMounts)
	return p
}

// ServiceTemplateContainersEnvToProto converts a ServiceTemplateContainersEnv object to its proto representation.
func RunAlphaServiceTemplateContainersEnvToProto(o *alpha.ServiceTemplateContainersEnv) *alphapb.RunAlphaServiceTemplateContainersEnv {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateContainersEnv{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetValueSource(RunAlphaServiceTemplateContainersEnvValueSourceToProto(o.ValueSource))
	return p
}

// ServiceTemplateContainersEnvValueSourceToProto converts a ServiceTemplateContainersEnvValueSource object to its proto representation.
func RunAlphaServiceTemplateContainersEnvValueSourceToProto(o *alpha.ServiceTemplateContainersEnvValueSource) *alphapb.RunAlphaServiceTemplateContainersEnvValueSource {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateContainersEnvValueSource{}
	p.SetSecretKeyRef(RunAlphaServiceTemplateContainersEnvValueSourceSecretKeyRefToProto(o.SecretKeyRef))
	return p
}

// ServiceTemplateContainersEnvValueSourceSecretKeyRefToProto converts a ServiceTemplateContainersEnvValueSourceSecretKeyRef object to its proto representation.
func RunAlphaServiceTemplateContainersEnvValueSourceSecretKeyRefToProto(o *alpha.ServiceTemplateContainersEnvValueSourceSecretKeyRef) *alphapb.RunAlphaServiceTemplateContainersEnvValueSourceSecretKeyRef {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateContainersEnvValueSourceSecretKeyRef{}
	p.SetSecret(dcl.ValueOrEmptyString(o.Secret))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// ServiceTemplateContainersResourcesToProto converts a ServiceTemplateContainersResources object to its proto representation.
func RunAlphaServiceTemplateContainersResourcesToProto(o *alpha.ServiceTemplateContainersResources) *alphapb.RunAlphaServiceTemplateContainersResources {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateContainersResources{}
	p.SetCpuIdle(dcl.ValueOrEmptyBool(o.CpuIdle))
	mLimits := make(map[string]string, len(o.Limits))
	for k, r := range o.Limits {
		mLimits[k] = r
	}
	p.SetLimits(mLimits)
	return p
}

// ServiceTemplateContainersPortsToProto converts a ServiceTemplateContainersPorts object to its proto representation.
func RunAlphaServiceTemplateContainersPortsToProto(o *alpha.ServiceTemplateContainersPorts) *alphapb.RunAlphaServiceTemplateContainersPorts {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateContainersPorts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetContainerPort(dcl.ValueOrEmptyInt64(o.ContainerPort))
	return p
}

// ServiceTemplateContainersVolumeMountsToProto converts a ServiceTemplateContainersVolumeMounts object to its proto representation.
func RunAlphaServiceTemplateContainersVolumeMountsToProto(o *alpha.ServiceTemplateContainersVolumeMounts) *alphapb.RunAlphaServiceTemplateContainersVolumeMounts {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateContainersVolumeMounts{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetMountPath(dcl.ValueOrEmptyString(o.MountPath))
	return p
}

// ServiceTemplateVolumesToProto converts a ServiceTemplateVolumes object to its proto representation.
func RunAlphaServiceTemplateVolumesToProto(o *alpha.ServiceTemplateVolumes) *alphapb.RunAlphaServiceTemplateVolumes {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateVolumes{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetSecret(RunAlphaServiceTemplateVolumesSecretToProto(o.Secret))
	p.SetCloudSqlInstance(RunAlphaServiceTemplateVolumesCloudSqlInstanceToProto(o.CloudSqlInstance))
	return p
}

// ServiceTemplateVolumesSecretToProto converts a ServiceTemplateVolumesSecret object to its proto representation.
func RunAlphaServiceTemplateVolumesSecretToProto(o *alpha.ServiceTemplateVolumesSecret) *alphapb.RunAlphaServiceTemplateVolumesSecret {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateVolumesSecret{}
	p.SetSecret(dcl.ValueOrEmptyString(o.Secret))
	p.SetDefaultMode(dcl.ValueOrEmptyInt64(o.DefaultMode))
	sItems := make([]*alphapb.RunAlphaServiceTemplateVolumesSecretItems, len(o.Items))
	for i, r := range o.Items {
		sItems[i] = RunAlphaServiceTemplateVolumesSecretItemsToProto(&r)
	}
	p.SetItems(sItems)
	return p
}

// ServiceTemplateVolumesSecretItemsToProto converts a ServiceTemplateVolumesSecretItems object to its proto representation.
func RunAlphaServiceTemplateVolumesSecretItemsToProto(o *alpha.ServiceTemplateVolumesSecretItems) *alphapb.RunAlphaServiceTemplateVolumesSecretItems {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateVolumesSecretItems{}
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	p.SetMode(dcl.ValueOrEmptyInt64(o.Mode))
	return p
}

// ServiceTemplateVolumesCloudSqlInstanceToProto converts a ServiceTemplateVolumesCloudSqlInstance object to its proto representation.
func RunAlphaServiceTemplateVolumesCloudSqlInstanceToProto(o *alpha.ServiceTemplateVolumesCloudSqlInstance) *alphapb.RunAlphaServiceTemplateVolumesCloudSqlInstance {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTemplateVolumesCloudSqlInstance{}
	sInstances := make([]string, len(o.Instances))
	for i, r := range o.Instances {
		sInstances[i] = r
	}
	p.SetInstances(sInstances)
	return p
}

// ServiceTrafficToProto converts a ServiceTraffic object to its proto representation.
func RunAlphaServiceTrafficToProto(o *alpha.ServiceTraffic) *alphapb.RunAlphaServiceTraffic {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTraffic{}
	p.SetType(RunAlphaServiceTrafficTypeEnumToProto(o.Type))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetTag(dcl.ValueOrEmptyString(o.Tag))
	return p
}

// ServiceTerminalConditionToProto converts a ServiceTerminalCondition object to its proto representation.
func RunAlphaServiceTerminalConditionToProto(o *alpha.ServiceTerminalCondition) *alphapb.RunAlphaServiceTerminalCondition {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTerminalCondition{}
	p.SetType(dcl.ValueOrEmptyString(o.Type))
	p.SetState(RunAlphaServiceTerminalConditionStateEnumToProto(o.State))
	p.SetMessage(dcl.ValueOrEmptyString(o.Message))
	p.SetLastTransitionTime(dcl.ValueOrEmptyString(o.LastTransitionTime))
	p.SetSeverity(RunAlphaServiceTerminalConditionSeverityEnumToProto(o.Severity))
	p.SetReason(RunAlphaServiceTerminalConditionReasonEnumToProto(o.Reason))
	p.SetRevisionReason(RunAlphaServiceTerminalConditionRevisionReasonEnumToProto(o.RevisionReason))
	p.SetJobReason(RunAlphaServiceTerminalConditionJobReasonEnumToProto(o.JobReason))
	return p
}

// ServiceTrafficStatusesToProto converts a ServiceTrafficStatuses object to its proto representation.
func RunAlphaServiceTrafficStatusesToProto(o *alpha.ServiceTrafficStatuses) *alphapb.RunAlphaServiceTrafficStatuses {
	if o == nil {
		return nil
	}
	p := &alphapb.RunAlphaServiceTrafficStatuses{}
	p.SetType(RunAlphaServiceTrafficStatusesTypeEnumToProto(o.Type))
	p.SetRevision(dcl.ValueOrEmptyString(o.Revision))
	p.SetPercent(dcl.ValueOrEmptyInt64(o.Percent))
	p.SetTag(dcl.ValueOrEmptyString(o.Tag))
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	return p
}

// ServiceToProto converts a Service resource to its proto representation.
func ServiceToProto(resource *alpha.Service) *alphapb.RunAlphaService {
	p := &alphapb.RunAlphaService{}
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
	p.SetIngress(RunAlphaServiceIngressEnumToProto(resource.Ingress))
	p.SetLaunchStage(RunAlphaServiceLaunchStageEnumToProto(resource.LaunchStage))
	p.SetBinaryAuthorization(RunAlphaServiceBinaryAuthorizationToProto(resource.BinaryAuthorization))
	p.SetTemplate(RunAlphaServiceTemplateToProto(resource.Template))
	p.SetTerminalCondition(RunAlphaServiceTerminalConditionToProto(resource.TerminalCondition))
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
	sTraffic := make([]*alphapb.RunAlphaServiceTraffic, len(resource.Traffic))
	for i, r := range resource.Traffic {
		sTraffic[i] = RunAlphaServiceTrafficToProto(&r)
	}
	p.SetTraffic(sTraffic)
	sTrafficStatuses := make([]*alphapb.RunAlphaServiceTrafficStatuses, len(resource.TrafficStatuses))
	for i, r := range resource.TrafficStatuses {
		sTrafficStatuses[i] = RunAlphaServiceTrafficStatusesToProto(&r)
	}
	p.SetTrafficStatuses(sTrafficStatuses)

	return p
}

// applyService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) applyService(ctx context.Context, c *alpha.Client, request *alphapb.ApplyRunAlphaServiceRequest) (*alphapb.RunAlphaService, error) {
	p := ProtoToService(request.GetResource())
	res, err := c.ApplyService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServiceToProto(res)
	return r, nil
}

// applyRunAlphaService handles the gRPC request by passing it to the underlying Service Apply() method.
func (s *ServiceServer) ApplyRunAlphaService(ctx context.Context, request *alphapb.ApplyRunAlphaServiceRequest) (*alphapb.RunAlphaService, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyService(ctx, cl, request)
}

// DeleteService handles the gRPC request by passing it to the underlying Service Delete() method.
func (s *ServiceServer) DeleteRunAlphaService(ctx context.Context, request *alphapb.DeleteRunAlphaServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteService(ctx, ProtoToService(request.GetResource()))

}

// ListRunAlphaService handles the gRPC request by passing it to the underlying ServiceList() method.
func (s *ServiceServer) ListRunAlphaService(ctx context.Context, request *alphapb.ListRunAlphaServiceRequest) (*alphapb.ListRunAlphaServiceResponse, error) {
	cl, err := createConfigService(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListService(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.RunAlphaService
	for _, r := range resources.Items {
		rp := ServiceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListRunAlphaServiceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigService(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
