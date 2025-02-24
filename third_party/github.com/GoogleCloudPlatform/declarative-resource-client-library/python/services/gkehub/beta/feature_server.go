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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/beta/gkehub_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

// FeatureServer implements the gRPC interface for Feature.
type FeatureServer struct{}

// ProtoToFeatureResourceStateStateEnum converts a FeatureResourceStateStateEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureResourceStateStateEnum(e betapb.GkehubBetaFeatureResourceStateStateEnum) *beta.FeatureResourceStateStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureResourceStateStateEnum_name[int32(e)]; ok {
		e := beta.FeatureResourceStateStateEnum(n[len("GkehubBetaFeatureResourceStateStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum converts a FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(e betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum) *beta.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum_name[int32(e)]; ok {
		e := beta.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(n[len("GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum converts a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(e betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum) *beta.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum_name[int32(e)]; ok {
		e := beta.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(n[len("GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureStateStateCodeEnum converts a FeatureStateStateCodeEnum enum from its proto representation.
func ProtoToGkehubBetaFeatureStateStateCodeEnum(e betapb.GkehubBetaFeatureStateStateCodeEnum) *beta.FeatureStateStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.GkehubBetaFeatureStateStateCodeEnum_name[int32(e)]; ok {
		e := beta.FeatureStateStateCodeEnum(n[len("GkehubBetaFeatureStateStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureResourceState converts a FeatureResourceState object from its proto representation.
func ProtoToGkehubBetaFeatureResourceState(p *betapb.GkehubBetaFeatureResourceState) *beta.FeatureResourceState {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureResourceState{
		State:        ProtoToGkehubBetaFeatureResourceStateStateEnum(p.GetState()),
		HasResources: dcl.Bool(p.GetHasResources()),
	}
	return obj
}

// ProtoToFeatureSpec converts a FeatureSpec object from its proto representation.
func ProtoToGkehubBetaFeatureSpec(p *betapb.GkehubBetaFeatureSpec) *beta.FeatureSpec {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpec{
		Multiclusteringress: ProtoToGkehubBetaFeatureSpecMulticlusteringress(p.GetMulticlusteringress()),
		Fleetobservability:  ProtoToGkehubBetaFeatureSpecFleetobservability(p.GetFleetobservability()),
	}
	return obj
}

// ProtoToFeatureSpecMulticlusteringress converts a FeatureSpecMulticlusteringress object from its proto representation.
func ProtoToGkehubBetaFeatureSpecMulticlusteringress(p *betapb.GkehubBetaFeatureSpecMulticlusteringress) *beta.FeatureSpecMulticlusteringress {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecMulticlusteringress{
		ConfigMembership: dcl.StringOrNil(p.GetConfigMembership()),
	}
	return obj
}

// ProtoToFeatureSpecFleetobservability converts a FeatureSpecFleetobservability object from its proto representation.
func ProtoToGkehubBetaFeatureSpecFleetobservability(p *betapb.GkehubBetaFeatureSpecFleetobservability) *beta.FeatureSpecFleetobservability {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecFleetobservability{
		LoggingConfig: ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfig(p.GetLoggingConfig()),
	}
	return obj
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfig converts a FeatureSpecFleetobservabilityLoggingConfig object from its proto representation.
func ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfig(p *betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfig) *beta.FeatureSpecFleetobservabilityLoggingConfig {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecFleetobservabilityLoggingConfig{
		DefaultConfig:        ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(p.GetDefaultConfig()),
		FleetScopeLogsConfig: ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(p.GetFleetScopeLogsConfig()),
	}
	return obj
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfigDefaultConfig converts a FeatureSpecFleetobservabilityLoggingConfigDefaultConfig object from its proto representation.
func ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig(p *betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig) *beta.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{
		Mode: ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig converts a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig object from its proto representation.
func ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig(p *betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) *beta.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{
		Mode: ProtoToGkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(p.GetMode()),
	}
	return obj
}

// ProtoToFeatureState converts a FeatureState object from its proto representation.
func ProtoToGkehubBetaFeatureState(p *betapb.GkehubBetaFeatureState) *beta.FeatureState {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureState{
		State: ProtoToGkehubBetaFeatureStateState(p.GetState()),
	}
	return obj
}

// ProtoToFeatureStateState converts a FeatureStateState object from its proto representation.
func ProtoToGkehubBetaFeatureStateState(p *betapb.GkehubBetaFeatureStateState) *beta.FeatureStateState {
	if p == nil {
		return nil
	}
	obj := &beta.FeatureStateState{
		Code:        ProtoToGkehubBetaFeatureStateStateCodeEnum(p.GetCode()),
		Description: dcl.StringOrNil(p.GetDescription()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToFeature converts a Feature resource from its proto representation.
func ProtoToFeature(p *betapb.GkehubBetaFeature) *beta.Feature {
	obj := &beta.Feature{
		Name:          dcl.StringOrNil(p.GetName()),
		ResourceState: ProtoToGkehubBetaFeatureResourceState(p.GetResourceState()),
		Spec:          ProtoToGkehubBetaFeatureSpec(p.GetSpec()),
		State:         ProtoToGkehubBetaFeatureState(p.GetState()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:    dcl.StringOrNil(p.GetDeleteTime()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// FeatureResourceStateStateEnumToProto converts a FeatureResourceStateStateEnum enum to its proto representation.
func GkehubBetaFeatureResourceStateStateEnumToProto(e *beta.FeatureResourceStateStateEnum) betapb.GkehubBetaFeatureResourceStateStateEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureResourceStateStateEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureResourceStateStateEnum_value["FeatureResourceStateStateEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureResourceStateStateEnum(v)
	}
	return betapb.GkehubBetaFeatureResourceStateStateEnum(0)
}

// FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumToProto converts a FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum enum to its proto representation.
func GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumToProto(e *beta.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum) betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum_value["FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(v)
	}
	return betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnum(0)
}

// FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumToProto converts a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum enum to its proto representation.
func GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumToProto(e *beta.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum) betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum_value["FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(v)
	}
	return betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnum(0)
}

// FeatureStateStateCodeEnumToProto converts a FeatureStateStateCodeEnum enum to its proto representation.
func GkehubBetaFeatureStateStateCodeEnumToProto(e *beta.FeatureStateStateCodeEnum) betapb.GkehubBetaFeatureStateStateCodeEnum {
	if e == nil {
		return betapb.GkehubBetaFeatureStateStateCodeEnum(0)
	}
	if v, ok := betapb.GkehubBetaFeatureStateStateCodeEnum_value["FeatureStateStateCodeEnum"+string(*e)]; ok {
		return betapb.GkehubBetaFeatureStateStateCodeEnum(v)
	}
	return betapb.GkehubBetaFeatureStateStateCodeEnum(0)
}

// FeatureResourceStateToProto converts a FeatureResourceState object to its proto representation.
func GkehubBetaFeatureResourceStateToProto(o *beta.FeatureResourceState) *betapb.GkehubBetaFeatureResourceState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureResourceState{}
	p.SetState(GkehubBetaFeatureResourceStateStateEnumToProto(o.State))
	p.SetHasResources(dcl.ValueOrEmptyBool(o.HasResources))
	return p
}

// FeatureSpecToProto converts a FeatureSpec object to its proto representation.
func GkehubBetaFeatureSpecToProto(o *beta.FeatureSpec) *betapb.GkehubBetaFeatureSpec {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpec{}
	p.SetMulticlusteringress(GkehubBetaFeatureSpecMulticlusteringressToProto(o.Multiclusteringress))
	p.SetFleetobservability(GkehubBetaFeatureSpecFleetobservabilityToProto(o.Fleetobservability))
	return p
}

// FeatureSpecMulticlusteringressToProto converts a FeatureSpecMulticlusteringress object to its proto representation.
func GkehubBetaFeatureSpecMulticlusteringressToProto(o *beta.FeatureSpecMulticlusteringress) *betapb.GkehubBetaFeatureSpecMulticlusteringress {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecMulticlusteringress{}
	p.SetConfigMembership(dcl.ValueOrEmptyString(o.ConfigMembership))
	return p
}

// FeatureSpecFleetobservabilityToProto converts a FeatureSpecFleetobservability object to its proto representation.
func GkehubBetaFeatureSpecFleetobservabilityToProto(o *beta.FeatureSpecFleetobservability) *betapb.GkehubBetaFeatureSpecFleetobservability {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecFleetobservability{}
	p.SetLoggingConfig(GkehubBetaFeatureSpecFleetobservabilityLoggingConfigToProto(o.LoggingConfig))
	return p
}

// FeatureSpecFleetobservabilityLoggingConfigToProto converts a FeatureSpecFleetobservabilityLoggingConfig object to its proto representation.
func GkehubBetaFeatureSpecFleetobservabilityLoggingConfigToProto(o *beta.FeatureSpecFleetobservabilityLoggingConfig) *betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfig{}
	p.SetDefaultConfig(GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigToProto(o.DefaultConfig))
	p.SetFleetScopeLogsConfig(GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigToProto(o.FleetScopeLogsConfig))
	return p
}

// FeatureSpecFleetobservabilityLoggingConfigDefaultConfigToProto converts a FeatureSpecFleetobservabilityLoggingConfigDefaultConfig object to its proto representation.
func GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigToProto(o *beta.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig) *betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
	p.SetMode(GkehubBetaFeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumToProto(o.Mode))
	return p
}

// FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigToProto converts a FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig object to its proto representation.
func GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigToProto(o *beta.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig) *betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
	p.SetMode(GkehubBetaFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumToProto(o.Mode))
	return p
}

// FeatureStateToProto converts a FeatureState object to its proto representation.
func GkehubBetaFeatureStateToProto(o *beta.FeatureState) *betapb.GkehubBetaFeatureState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureState{}
	p.SetState(GkehubBetaFeatureStateStateToProto(o.State))
	return p
}

// FeatureStateStateToProto converts a FeatureStateState object to its proto representation.
func GkehubBetaFeatureStateStateToProto(o *beta.FeatureStateState) *betapb.GkehubBetaFeatureStateState {
	if o == nil {
		return nil
	}
	p := &betapb.GkehubBetaFeatureStateState{}
	p.SetCode(GkehubBetaFeatureStateStateCodeEnumToProto(o.Code))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// FeatureToProto converts a Feature resource to its proto representation.
func FeatureToProto(resource *beta.Feature) *betapb.GkehubBetaFeature {
	p := &betapb.GkehubBetaFeature{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetResourceState(GkehubBetaFeatureResourceStateToProto(resource.ResourceState))
	p.SetSpec(GkehubBetaFeatureSpecToProto(resource.Spec))
	p.SetState(GkehubBetaFeatureStateToProto(resource.State))
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
func (s *FeatureServer) applyFeature(ctx context.Context, c *beta.Client, request *betapb.ApplyGkehubBetaFeatureRequest) (*betapb.GkehubBetaFeature, error) {
	p := ProtoToFeature(request.GetResource())
	res, err := c.ApplyFeature(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureToProto(res)
	return r, nil
}

// applyGkehubBetaFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) ApplyGkehubBetaFeature(ctx context.Context, request *betapb.ApplyGkehubBetaFeatureRequest) (*betapb.GkehubBetaFeature, error) {
	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFeature(ctx, cl, request)
}

// DeleteFeature handles the gRPC request by passing it to the underlying Feature Delete() method.
func (s *FeatureServer) DeleteGkehubBetaFeature(ctx context.Context, request *betapb.DeleteGkehubBetaFeatureRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeature(ctx, ProtoToFeature(request.GetResource()))

}

// ListGkehubBetaFeature handles the gRPC request by passing it to the underlying FeatureList() method.
func (s *FeatureServer) ListGkehubBetaFeature(ctx context.Context, request *betapb.ListGkehubBetaFeatureRequest) (*betapb.ListGkehubBetaFeatureResponse, error) {
	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeature(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.GkehubBetaFeature
	for _, r := range resources.Items {
		rp := FeatureToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListGkehubBetaFeatureResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFeature(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
