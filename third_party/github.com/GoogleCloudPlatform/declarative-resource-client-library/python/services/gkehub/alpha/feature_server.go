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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/alpha/gkehub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
)

// FeatureServer implements the gRPC interface for Feature.
type FeatureServer struct{}

// ProtoToFeatureResourceStateStateEnum converts a FeatureResourceStateStateEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureResourceStateStateEnum(e alphapb.GkehubAlphaFeatureResourceStateStateEnum) *alpha.FeatureResourceStateStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureResourceStateStateEnum_name[int32(e)]; ok {
		e := alpha.FeatureResourceStateStateEnum(n[len("GkehubAlphaFeatureResourceStateStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum converts a FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(e alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum) *alpha.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum_name[int32(e)]; ok {
		e := alpha.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(n[len("GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum converts a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(e alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum) *alpha.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum_name[int32(e)]; ok {
		e := alpha.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(n[len("GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureStateStateCodeEnum converts a FeatureStateStateCodeEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureStateStateCodeEnum(e alphapb.GkehubAlphaFeatureStateStateCodeEnum) *alpha.FeatureStateStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureStateStateCodeEnum_name[int32(e)]; ok {
		e := alpha.FeatureStateStateCodeEnum(n[len("GkehubAlphaFeatureStateStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum converts a FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(e alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum) *alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum_name[int32(e)]; ok {
		e := alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(n[len("GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureResourceState converts a FeatureResourceState object from its proto representation.
func ProtoToGkehubAlphaFeatureResourceState(p *alphapb.GkehubAlphaFeatureResourceState) *alpha.FeatureResourceState {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureResourceState{
		State:        ProtoToGkehubAlphaFeatureResourceStateStateEnum(p.GetState()),
		HasResources: dcl.Bool(p.GetHasResources()),
	}
	return obj
}

// ProtoToFeatureSpec converts a FeatureSpec object from its proto representation.
func ProtoToGkehubAlphaFeatureSpec(p *alphapb.GkehubAlphaFeatureSpec) *alpha.FeatureSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpec{
		Multiclusteringress: ProtoToGkehubAlphaFeatureSpecMulticlusteringress(p.GetMulticlusteringress()),
		Cloudauditlogging:   ProtoToGkehubAlphaFeatureSpecCloudauditlogging(p.GetCloudauditlogging()),
		Fleetobservability:  ProtoToGkehubAlphaFeatureSpecFleetobservability(p.GetFleetobservability()),
	}
	return obj
}

// ProtoToFeatureSpecMulticlusteringress converts a FeatureSpecMulticlusteringress object from its proto representation.
func ProtoToGkehubAlphaFeatureSpecMulticlusteringress(p *alphapb.GkehubAlphaFeatureSpecMulticlusteringress) *alpha.FeatureSpecMulticlusteringress {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpecMulticlusteringress{
		ConfigMembership: dcl.StringOrNil(p.GetConfigMembership()),
	}
	return obj
}

// ProtoToFeatureSpecCloudauditlogging converts a FeatureSpecCloudauditlogging object from its proto representation.
func ProtoToGkehubAlphaFeatureSpecCloudauditlogging(p *alphapb.GkehubAlphaFeatureSpecCloudauditlogging) *alpha.FeatureSpecCloudauditlogging {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpecCloudauditlogging{}
	for _, r := range p.GetAllowlistedServiceAccounts() {
		obj.AllowlistedServiceAccounts = append(obj.AllowlistedServiceAccounts, r)
	}
	return obj
}

// ProtoToFeatureSpecFleetobservability converts a FeatureSpecFleetobservability object from its proto representation.
func ProtoToGkehubAlphaFeatureSpecFleetobservability(p *alphapb.GkehubAlphaFeatureSpecFleetobservability) *alpha.FeatureSpecFleetobservability {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpecFleetobservability{
		LoggingConfig: ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfig(p.GetLoggingConfig()),
	}
	return obj
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfig converts a FeatureSpecFleetobservabilityLoggingConfig object from its proto representation.
func ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfig(p *alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfig) *alpha.FeatureSpecFleetobservabilityLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpecFleetobservabilityLoggingConfig{
		DefaultConfig:        ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(p.GetDefaultConfig()),
		FleetScopeLogsConfig: ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(p.GetFleetScopeLogsConfig()),
	}
	return obj
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfigDefaultConfig converts a FeatureSpecFleetobservabilityLoggingConfigDefaultConfig object from its proto representation.
func ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(p *alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig) *alpha.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{
		Mode: ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig converts a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig object from its proto representation.
func ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(p *alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) *alpha.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{
		Mode: ProtoToGkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToFeatureState converts a FeatureState object from its proto representation.
func ProtoToGkehubAlphaFeatureState(p *alphapb.GkehubAlphaFeatureState) *alpha.FeatureState {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureState{
		State:       ProtoToGkehubAlphaFeatureStateState(p.GetState()),
		Servicemesh: ProtoToGkehubAlphaFeatureStateServicemesh(p.GetServicemesh()),
	}
	return obj
}

// ProtoToFeatureStateState converts a FeatureStateState object from its proto representation.
func ProtoToGkehubAlphaFeatureStateState(p *alphapb.GkehubAlphaFeatureStateState) *alpha.FeatureStateState {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateState{
		Code:        ProtoToGkehubAlphaFeatureStateStateCodeEnum(p.GetCode()),
		Description: dcl.StringOrNil(p.GetDescription()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToFeatureStateServicemesh converts a FeatureStateServicemesh object from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemesh(p *alphapb.GkehubAlphaFeatureStateServicemesh) *alpha.FeatureStateServicemesh {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateServicemesh{}
	for _, r := range p.GetAnalysisMessages() {
		obj.AnalysisMessages = append(obj.AnalysisMessages, *ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessages(r))
	}
	return obj
}

// ProtoToFeatureStateServicemeshAnalysisMessages converts a FeatureStateServicemeshAnalysisMessages object from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessages(p *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessages) *alpha.FeatureStateServicemeshAnalysisMessages {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateServicemeshAnalysisMessages{
		MessageBase: ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase(p.GetMessageBase()),
		Description: dcl.StringOrNil(p.GetDescription()),
	}
	for _, r := range p.GetResourcePaths() {
		obj.ResourcePaths = append(obj.ResourcePaths, r)
	}
	return obj
}

// ProtoToFeatureStateServicemeshAnalysisMessagesMessageBase converts a FeatureStateServicemeshAnalysisMessagesMessageBase object from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase(p *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase) *alpha.FeatureStateServicemeshAnalysisMessagesMessageBase {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateServicemeshAnalysisMessagesMessageBase{
		Type:             ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType(p.GetType()),
		Level:            ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(p.GetLevel()),
		DocumentationUrl: dcl.StringOrNil(p.GetDocumentationUrl()),
	}
	return obj
}

// ProtoToFeatureStateServicemeshAnalysisMessagesMessageBaseType converts a FeatureStateServicemeshAnalysisMessagesMessageBaseType object from its proto representation.
func ProtoToGkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType(p *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType) *alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseType {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseType{
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Code:        dcl.StringOrNil(p.GetCode()),
	}
	return obj
}

// ProtoToFeature converts a Feature resource from its proto representation.
func ProtoToFeature(p *alphapb.GkehubAlphaFeature) *alpha.Feature {
	obj := &alpha.Feature{
		Name:          dcl.StringOrNil(p.GetName()),
		ResourceState: ProtoToGkehubAlphaFeatureResourceState(p.GetResourceState()),
		Spec:          ProtoToGkehubAlphaFeatureSpec(p.GetSpec()),
		State:         ProtoToGkehubAlphaFeatureState(p.GetState()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:    dcl.StringOrNil(p.GetDeleteTime()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// FeatureResourceStateStateEnumToProto converts a FeatureResourceStateStateEnum enum to its proto representation.
func GkehubAlphaFeatureResourceStateStateEnumToProto(e *alpha.FeatureResourceStateStateEnum) alphapb.GkehubAlphaFeatureResourceStateStateEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureResourceStateStateEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureResourceStateStateEnum_value["FeatureResourceStateStateEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureResourceStateStateEnum(v)
	}
	return alphapb.GkehubAlphaFeatureResourceStateStateEnum(0)
}

// FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumToProto converts a FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum enum to its proto representation.
func GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumToProto(e *alpha.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum) alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum_value["FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(v)
	}
	return alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(0)
}

// FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumToProto converts a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum enum to its proto representation.
func GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumToProto(e *alpha.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum) alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum_value["FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(v)
	}
	return alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(0)
}

// FeatureStateStateCodeEnumToProto converts a FeatureStateStateCodeEnum enum to its proto representation.
func GkehubAlphaFeatureStateStateCodeEnumToProto(e *alpha.FeatureStateStateCodeEnum) alphapb.GkehubAlphaFeatureStateStateCodeEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureStateStateCodeEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureStateStateCodeEnum_value["FeatureStateStateCodeEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureStateStateCodeEnum(v)
	}
	return alphapb.GkehubAlphaFeatureStateStateCodeEnum(0)
}

// FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumToProto converts a FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum enum to its proto representation.
func GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumToProto(e *alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum) alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum_value["FeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(v)
	}
	return alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnum(0)
}

// FeatureResourceStateToProto converts a FeatureResourceState object to its proto representation.
func GkehubAlphaFeatureResourceStateToProto(o *alpha.FeatureResourceState) *alphapb.GkehubAlphaFeatureResourceState {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureResourceState{}
	p.SetState(GkehubAlphaFeatureResourceStateStateEnumToProto(o.State))
	p.SetHasResources(dcl.ValueOrEmptyBool(o.HasResources))
	return p
}

// FeatureSpecToProto converts a FeatureSpec object to its proto representation.
func GkehubAlphaFeatureSpecToProto(o *alpha.FeatureSpec) *alphapb.GkehubAlphaFeatureSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpec{}
	p.SetMulticlusteringress(GkehubAlphaFeatureSpecMulticlusteringressToProto(o.Multiclusteringress))
	p.SetCloudauditlogging(GkehubAlphaFeatureSpecCloudauditloggingToProto(o.Cloudauditlogging))
	p.SetFleetobservability(GkehubAlphaFeatureSpecFleetobservabilityToProto(o.Fleetobservability))
	return p
}

// FeatureSpecMulticlusteringressToProto converts a FeatureSpecMulticlusteringress object to its proto representation.
func GkehubAlphaFeatureSpecMulticlusteringressToProto(o *alpha.FeatureSpecMulticlusteringress) *alphapb.GkehubAlphaFeatureSpecMulticlusteringress {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpecMulticlusteringress{}
	p.SetConfigMembership(dcl.ValueOrEmptyString(o.ConfigMembership))
	return p
}

// FeatureSpecCloudauditloggingToProto converts a FeatureSpecCloudauditlogging object to its proto representation.
func GkehubAlphaFeatureSpecCloudauditloggingToProto(o *alpha.FeatureSpecCloudauditlogging) *alphapb.GkehubAlphaFeatureSpecCloudauditlogging {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpecCloudauditlogging{}
	sAllowlistedServiceAccounts := make([]string, len(o.AllowlistedServiceAccounts))
	for i, r := range o.AllowlistedServiceAccounts {
		sAllowlistedServiceAccounts[i] = r
	}
	p.SetAllowlistedServiceAccounts(sAllowlistedServiceAccounts)
	return p
}

// FeatureSpecFleetobservabilityToProto converts a FeatureSpecFleetobservability object to its proto representation.
func GkehubAlphaFeatureSpecFleetobservabilityToProto(o *alpha.FeatureSpecFleetobservability) *alphapb.GkehubAlphaFeatureSpecFleetobservability {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpecFleetobservability{}
	p.SetLoggingConfig(GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigToProto(o.LoggingConfig))
	return p
}

// FeatureSpecFleetobservabilityLoggingConfigToProto converts a FeatureSpecFleetobservabilityLoggingConfig object to its proto representation.
func GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigToProto(o *alpha.FeatureSpecFleetobservabilityLoggingConfig) *alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfig{}
	p.SetDefaultConfig(GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigToProto(o.DefaultConfig))
	p.SetFleetScopeLogsConfig(GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigToProto(o.FleetScopeLogsConfig))
	return p
}

// FeatureSpecFleetobservabilityLoggingConfigDefaultConfigToProto converts a FeatureSpecFleetobservabilityLoggingConfigDefaultConfig object to its proto representation.
func GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigToProto(o *alpha.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) *alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
	p.SetMode(GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumToProto(o.Mode))
	return p
}

// FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigToProto converts a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig object to its proto representation.
func GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigToProto(o *alpha.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) *alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
	p.SetMode(GkehubAlphaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumToProto(o.Mode))
	return p
}

// FeatureStateToProto converts a FeatureState object to its proto representation.
func GkehubAlphaFeatureStateToProto(o *alpha.FeatureState) *alphapb.GkehubAlphaFeatureState {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureState{}
	p.SetState(GkehubAlphaFeatureStateStateToProto(o.State))
	p.SetServicemesh(GkehubAlphaFeatureStateServicemeshToProto(o.Servicemesh))
	return p
}

// FeatureStateStateToProto converts a FeatureStateState object to its proto representation.
func GkehubAlphaFeatureStateStateToProto(o *alpha.FeatureStateState) *alphapb.GkehubAlphaFeatureStateState {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateState{}
	p.SetCode(GkehubAlphaFeatureStateStateCodeEnumToProto(o.Code))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// FeatureStateServicemeshToProto converts a FeatureStateServicemesh object to its proto representation.
func GkehubAlphaFeatureStateServicemeshToProto(o *alpha.FeatureStateServicemesh) *alphapb.GkehubAlphaFeatureStateServicemesh {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateServicemesh{}
	sAnalysisMessages := make([]*alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessages, len(o.AnalysisMessages))
	for i, r := range o.AnalysisMessages {
		sAnalysisMessages[i] = GkehubAlphaFeatureStateServicemeshAnalysisMessagesToProto(&r)
	}
	p.SetAnalysisMessages(sAnalysisMessages)
	return p
}

// FeatureStateServicemeshAnalysisMessagesToProto converts a FeatureStateServicemeshAnalysisMessages object to its proto representation.
func GkehubAlphaFeatureStateServicemeshAnalysisMessagesToProto(o *alpha.FeatureStateServicemeshAnalysisMessages) *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessages {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessages{}
	p.SetMessageBase(GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseToProto(o.MessageBase))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	sResourcePaths := make([]string, len(o.ResourcePaths))
	for i, r := range o.ResourcePaths {
		sResourcePaths[i] = r
	}
	p.SetResourcePaths(sResourcePaths)
	mArgs := make(map[string]string, len(o.Args))
	for k, r := range o.Args {
		mArgs[k] = r
	}
	p.SetArgs(mArgs)
	return p
}

// FeatureStateServicemeshAnalysisMessagesMessageBaseToProto converts a FeatureStateServicemeshAnalysisMessagesMessageBase object to its proto representation.
func GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseToProto(o *alpha.FeatureStateServicemeshAnalysisMessagesMessageBase) *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBase{}
	p.SetType(GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseTypeToProto(o.Type))
	p.SetLevel(GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseLevelEnumToProto(o.Level))
	p.SetDocumentationUrl(dcl.ValueOrEmptyString(o.DocumentationUrl))
	return p
}

// FeatureStateServicemeshAnalysisMessagesMessageBaseTypeToProto converts a FeatureStateServicemeshAnalysisMessagesMessageBaseType object to its proto representation.
func GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseTypeToProto(o *alpha.FeatureStateServicemeshAnalysisMessagesMessageBaseType) *alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureStateServicemeshAnalysisMessagesMessageBaseType{}
	p.SetDisplayName(dcl.ValueOrEmptyString(o.DisplayName))
	p.SetCode(dcl.ValueOrEmptyString(o.Code))
	return p
}

// FeatureToProto converts a Feature resource to its proto representation.
func FeatureToProto(resource *alpha.Feature) *alphapb.GkehubAlphaFeature {
	p := &alphapb.GkehubAlphaFeature{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetResourceState(GkehubAlphaFeatureResourceStateToProto(resource.ResourceState))
	p.SetSpec(GkehubAlphaFeatureSpecToProto(resource.Spec))
	p.SetState(GkehubAlphaFeatureStateToProto(resource.State))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) applyFeature(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkehubAlphaFeatureRequest) (*alphapb.GkehubAlphaFeature, error) {
	p := ProtoToFeature(request.GetResource())
	res, err := c.ApplyFeature(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureToProto(res)
	return r, nil
}

// applyGkehubAlphaFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) ApplyGkehubAlphaFeature(ctx context.Context, request *alphapb.ApplyGkehubAlphaFeatureRequest) (*alphapb.GkehubAlphaFeature, error) {
	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFeature(ctx, cl, request)
}

// DeleteFeature handles the gRPC request by passing it to the underlying Feature Delete() method.
func (s *FeatureServer) DeleteGkehubAlphaFeature(ctx context.Context, request *alphapb.DeleteGkehubAlphaFeatureRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeature(ctx, ProtoToFeature(request.GetResource()))

}

// ListGkehubAlphaFeature handles the gRPC request by passing it to the underlying FeatureList() method.
func (s *FeatureServer) ListGkehubAlphaFeature(ctx context.Context, request *alphapb.ListGkehubAlphaFeatureRequest) (*alphapb.ListGkehubAlphaFeatureResponse, error) {
	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeature(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkehubAlphaFeature
	for _, r := range resources.Items {
		rp := FeatureToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListGkehubAlphaFeatureResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFeature(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
