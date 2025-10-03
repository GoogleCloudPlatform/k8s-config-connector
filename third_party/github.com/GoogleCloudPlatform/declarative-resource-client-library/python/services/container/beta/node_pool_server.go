// Copyright 2021 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/container/beta/container_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container/beta"
)

// Server implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolConfigSandboxConfigTypeEnum converts a NodePoolConfigSandboxConfigTypeEnum enum from its proto representation.
func ProtoToContainerBetaNodePoolConfigSandboxConfigTypeEnum(e betapb.ContainerBetaNodePoolConfigSandboxConfigTypeEnum) *beta.NodePoolConfigSandboxConfigTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaNodePoolConfigSandboxConfigTypeEnum_name[int32(e)]; ok {
		e := beta.NodePoolConfigSandboxConfigTypeEnum(n[len("ContainerBetaNodePoolConfigSandboxConfigTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfigReservationAffinityConsumeReservationTypeEnum converts a NodePoolConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum(e betapb.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum) *beta.NodePoolConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := beta.NodePoolConfigReservationAffinityConsumeReservationTypeEnum(n[len("ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConditionsCodeEnum converts a NodePoolConditionsCodeEnum enum from its proto representation.
func ProtoToContainerBetaNodePoolConditionsCodeEnum(e betapb.ContainerBetaNodePoolConditionsCodeEnum) *beta.NodePoolConditionsCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerBetaNodePoolConditionsCodeEnum_name[int32(e)]; ok {
		e := beta.NodePoolConditionsCodeEnum(n[len("ContainerBetaNodePoolConditionsCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig resource from its proto representation.
func ProtoToContainerBetaNodePoolConfig(p *betapb.ContainerBetaNodePoolConfig) *beta.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfig{
		MachineType:            dcl.StringOrNil(p.MachineType),
		DiskSizeGb:             dcl.Int64OrNil(p.DiskSizeGb),
		ServiceAccount:         dcl.StringOrNil(p.ServiceAccount),
		ImageType:              dcl.StringOrNil(p.ImageType),
		LocalSsdCount:          dcl.Int64OrNil(p.LocalSsdCount),
		Preemptible:            dcl.Bool(p.Preemptible),
		DiskType:               dcl.StringOrNil(p.DiskType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		SandboxConfig:          ProtoToContainerBetaNodePoolConfigSandboxConfig(p.GetSandboxConfig()),
		ReservationAffinity:    ProtoToContainerBetaNodePoolConfigReservationAffinity(p.GetReservationAffinity()),
		ShieldedInstanceConfig: ProtoToContainerBetaNodePoolConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToContainerBetaNodePoolConfigAccelerators(r))
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerBetaNodePoolConfigTaints(r))
	}
	return obj
}

// ProtoToNodePoolConfigAccelerators converts a NodePoolConfigAccelerators resource from its proto representation.
func ProtoToContainerBetaNodePoolConfigAccelerators(p *betapb.ContainerBetaNodePoolConfigAccelerators) *beta.NodePoolConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToNodePoolConfigTaints converts a NodePoolConfigTaints resource from its proto representation.
func ProtoToContainerBetaNodePoolConfigTaints(p *betapb.ContainerBetaNodePoolConfigTaints) *beta.NodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: dcl.StringOrNil(p.Effect),
	}
	return obj
}

// ProtoToNodePoolConfigSandboxConfig converts a NodePoolConfigSandboxConfig resource from its proto representation.
func ProtoToContainerBetaNodePoolConfigSandboxConfig(p *betapb.ContainerBetaNodePoolConfigSandboxConfig) *beta.NodePoolConfigSandboxConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigSandboxConfig{
		Type: ProtoToContainerBetaNodePoolConfigSandboxConfigTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToNodePoolConfigReservationAffinity converts a NodePoolConfigReservationAffinity resource from its proto representation.
func ProtoToContainerBetaNodePoolConfigReservationAffinity(p *betapb.ContainerBetaNodePoolConfigReservationAffinity) *beta.NodePoolConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigReservationAffinity{
		ConsumeReservationType: ProtoToContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToNodePoolConfigShieldedInstanceConfig converts a NodePoolConfigShieldedInstanceConfig resource from its proto representation.
func ProtoToContainerBetaNodePoolConfigShieldedInstanceConfig(p *betapb.ContainerBetaNodePoolConfigShieldedInstanceConfig) *beta.NodePoolConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling resource from its proto representation.
func ProtoToContainerBetaNodePoolAutoscaling(p *betapb.ContainerBetaNodePoolAutoscaling) *beta.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolAutoscaling{
		Enabled:         dcl.Bool(p.Enabled),
		MinNodeCount:    dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount:    dcl.Int64OrNil(p.MaxNodeCount),
		Autoprovisioned: dcl.Bool(p.Autoprovisioned),
	}
	return obj
}

// ProtoToNodePoolManagement converts a NodePoolManagement resource from its proto representation.
func ProtoToContainerBetaNodePoolManagement(p *betapb.ContainerBetaNodePoolManagement) *beta.NodePoolManagement {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolManagement{
		AutoUpgrade:    dcl.Bool(p.AutoUpgrade),
		AutoRepair:     dcl.Bool(p.AutoRepair),
		UpgradeOptions: ProtoToContainerBetaNodePoolManagementUpgradeOptions(p.GetUpgradeOptions()),
	}
	return obj
}

// ProtoToNodePoolManagementUpgradeOptions converts a NodePoolManagementUpgradeOptions resource from its proto representation.
func ProtoToContainerBetaNodePoolManagementUpgradeOptions(p *betapb.ContainerBetaNodePoolManagementUpgradeOptions) *beta.NodePoolManagementUpgradeOptions {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.StringOrNil(p.GetAutoUpgradeStartTime()),
		Description:          dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToContainerBetaNodePoolMaxPodsConstraint(p *betapb.ContainerBetaNodePoolMaxPodsConstraint) *beta.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToNodePoolConditions converts a NodePoolConditions resource from its proto representation.
func ProtoToContainerBetaNodePoolConditions(p *betapb.ContainerBetaNodePoolConditions) *beta.NodePoolConditions {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConditions{
		Code:    ProtoToContainerBetaNodePoolConditionsCodeEnum(p.GetCode()),
		Message: dcl.StringOrNil(p.Message),
	}
	return obj
}

// ProtoToNodePoolUpgradeSettings converts a NodePoolUpgradeSettings resource from its proto representation.
func ProtoToContainerBetaNodePoolUpgradeSettings(p *betapb.ContainerBetaNodePoolUpgradeSettings) *beta.NodePoolUpgradeSettings {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolUpgradeSettings{
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *betapb.ContainerBetaNodePool) *beta.NodePool {
	obj := &beta.NodePool{
		Name:              dcl.StringOrNil(p.Name),
		Config:            ProtoToContainerBetaNodePoolConfig(p.GetConfig()),
		NodeCount:         dcl.Int64OrNil(p.NodeCount),
		Version:           dcl.StringOrNil(p.Version),
		Status:            dcl.StringOrNil(p.Status),
		StatusMessage:     dcl.StringOrNil(p.StatusMessage),
		Autoscaling:       ProtoToContainerBetaNodePoolAutoscaling(p.GetAutoscaling()),
		Management:        ProtoToContainerBetaNodePoolManagement(p.GetManagement()),
		MaxPodsConstraint: ProtoToContainerBetaNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		PodIPv4CidrSize:   dcl.Int64OrNil(p.PodIpv4CidrSize),
		UpgradeSettings:   ProtoToContainerBetaNodePoolUpgradeSettings(p.GetUpgradeSettings()),
		Cluster:           dcl.StringOrNil(p.Cluster),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToContainerBetaNodePoolConditions(r))
	}
	return obj
}

// NodePoolConfigSandboxConfigTypeEnumToProto converts a NodePoolConfigSandboxConfigTypeEnum enum to its proto representation.
func ContainerBetaNodePoolConfigSandboxConfigTypeEnumToProto(e *beta.NodePoolConfigSandboxConfigTypeEnum) betapb.ContainerBetaNodePoolConfigSandboxConfigTypeEnum {
	if e == nil {
		return betapb.ContainerBetaNodePoolConfigSandboxConfigTypeEnum(0)
	}
	if v, ok := betapb.ContainerBetaNodePoolConfigSandboxConfigTypeEnum_value["NodePoolConfigSandboxConfigTypeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaNodePoolConfigSandboxConfigTypeEnum(v)
	}
	return betapb.ContainerBetaNodePoolConfigSandboxConfigTypeEnum(0)
}

// NodePoolConfigReservationAffinityConsumeReservationTypeEnumToProto converts a NodePoolConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnumToProto(e *beta.NodePoolConfigReservationAffinityConsumeReservationTypeEnum) betapb.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return betapb.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := betapb.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum_value["NodePoolConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return betapb.ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// NodePoolConditionsCodeEnumToProto converts a NodePoolConditionsCodeEnum enum to its proto representation.
func ContainerBetaNodePoolConditionsCodeEnumToProto(e *beta.NodePoolConditionsCodeEnum) betapb.ContainerBetaNodePoolConditionsCodeEnum {
	if e == nil {
		return betapb.ContainerBetaNodePoolConditionsCodeEnum(0)
	}
	if v, ok := betapb.ContainerBetaNodePoolConditionsCodeEnum_value["NodePoolConditionsCodeEnum"+string(*e)]; ok {
		return betapb.ContainerBetaNodePoolConditionsCodeEnum(v)
	}
	return betapb.ContainerBetaNodePoolConditionsCodeEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig resource to its proto representation.
func ContainerBetaNodePoolConfigToProto(o *beta.NodePoolConfig) *betapb.ContainerBetaNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolConfig{
		MachineType:            dcl.ValueOrEmptyString(o.MachineType),
		DiskSizeGb:             dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		ServiceAccount:         dcl.ValueOrEmptyString(o.ServiceAccount),
		ImageType:              dcl.ValueOrEmptyString(o.ImageType),
		LocalSsdCount:          dcl.ValueOrEmptyInt64(o.LocalSsdCount),
		Preemptible:            dcl.ValueOrEmptyBool(o.Preemptible),
		DiskType:               dcl.ValueOrEmptyString(o.DiskType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		SandboxConfig:          ContainerBetaNodePoolConfigSandboxConfigToProto(o.SandboxConfig),
		ReservationAffinity:    ContainerBetaNodePoolConfigReservationAffinityToProto(o.ReservationAffinity),
		ShieldedInstanceConfig: ContainerBetaNodePoolConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
	}
	for _, r := range o.OAuthScopes {
		p.OauthScopes = append(p.OauthScopes, r)
	}
	p.Metadata = make(map[string]string)
	for k, r := range o.Metadata {
		p.Metadata[k] = r
	}
	p.Labels = make(map[string]string)
	for k, r := range o.Labels {
		p.Labels[k] = r
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	for _, r := range o.Accelerators {
		p.Accelerators = append(p.Accelerators, ContainerBetaNodePoolConfigAcceleratorsToProto(&r))
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, ContainerBetaNodePoolConfigTaintsToProto(&r))
	}
	return p
}

// NodePoolConfigAcceleratorsToProto converts a NodePoolConfigAccelerators resource to its proto representation.
func ContainerBetaNodePoolConfigAcceleratorsToProto(o *beta.NodePoolConfigAccelerators) *betapb.ContainerBetaNodePoolConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolConfigAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// NodePoolConfigTaintsToProto converts a NodePoolConfigTaints resource to its proto representation.
func ContainerBetaNodePoolConfigTaintsToProto(o *beta.NodePoolConfigTaints) *betapb.ContainerBetaNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: dcl.ValueOrEmptyString(o.Effect),
	}
	return p
}

// NodePoolConfigSandboxConfigToProto converts a NodePoolConfigSandboxConfig resource to its proto representation.
func ContainerBetaNodePoolConfigSandboxConfigToProto(o *beta.NodePoolConfigSandboxConfig) *betapb.ContainerBetaNodePoolConfigSandboxConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolConfigSandboxConfig{
		Type: ContainerBetaNodePoolConfigSandboxConfigTypeEnumToProto(o.Type),
	}
	return p
}

// NodePoolConfigReservationAffinityToProto converts a NodePoolConfigReservationAffinity resource to its proto representation.
func ContainerBetaNodePoolConfigReservationAffinityToProto(o *beta.NodePoolConfigReservationAffinity) *betapb.ContainerBetaNodePoolConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolConfigReservationAffinity{
		ConsumeReservationType: ContainerBetaNodePoolConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// NodePoolConfigShieldedInstanceConfigToProto converts a NodePoolConfigShieldedInstanceConfig resource to its proto representation.
func ContainerBetaNodePoolConfigShieldedInstanceConfigToProto(o *beta.NodePoolConfigShieldedInstanceConfig) *betapb.ContainerBetaNodePoolConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling resource to its proto representation.
func ContainerBetaNodePoolAutoscalingToProto(o *beta.NodePoolAutoscaling) *betapb.ContainerBetaNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolAutoscaling{
		Enabled:         dcl.ValueOrEmptyBool(o.Enabled),
		MinNodeCount:    dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount:    dcl.ValueOrEmptyInt64(o.MaxNodeCount),
		Autoprovisioned: dcl.ValueOrEmptyBool(o.Autoprovisioned),
	}
	return p
}

// NodePoolManagementToProto converts a NodePoolManagement resource to its proto representation.
func ContainerBetaNodePoolManagementToProto(o *beta.NodePoolManagement) *betapb.ContainerBetaNodePoolManagement {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolManagement{
		AutoUpgrade:    dcl.ValueOrEmptyBool(o.AutoUpgrade),
		AutoRepair:     dcl.ValueOrEmptyBool(o.AutoRepair),
		UpgradeOptions: ContainerBetaNodePoolManagementUpgradeOptionsToProto(o.UpgradeOptions),
	}
	return p
}

// NodePoolManagementUpgradeOptionsToProto converts a NodePoolManagementUpgradeOptions resource to its proto representation.
func ContainerBetaNodePoolManagementUpgradeOptionsToProto(o *beta.NodePoolManagementUpgradeOptions) *betapb.ContainerBetaNodePoolManagementUpgradeOptions {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.ValueOrEmptyString(o.AutoUpgradeStartTime),
		Description:          dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint resource to its proto representation.
func ContainerBetaNodePoolMaxPodsConstraintToProto(o *beta.NodePoolMaxPodsConstraint) *betapb.ContainerBetaNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// NodePoolConditionsToProto converts a NodePoolConditions resource to its proto representation.
func ContainerBetaNodePoolConditionsToProto(o *beta.NodePoolConditions) *betapb.ContainerBetaNodePoolConditions {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolConditions{
		Code:    ContainerBetaNodePoolConditionsCodeEnumToProto(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	return p
}

// NodePoolUpgradeSettingsToProto converts a NodePoolUpgradeSettings resource to its proto representation.
func ContainerBetaNodePoolUpgradeSettingsToProto(o *beta.NodePoolUpgradeSettings) *betapb.ContainerBetaNodePoolUpgradeSettings {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerBetaNodePoolUpgradeSettings{
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
	}
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *beta.NodePool) *betapb.ContainerBetaNodePool {
	p := &betapb.ContainerBetaNodePool{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Config:            ContainerBetaNodePoolConfigToProto(resource.Config),
		NodeCount:         dcl.ValueOrEmptyInt64(resource.NodeCount),
		Version:           dcl.ValueOrEmptyString(resource.Version),
		Status:            dcl.ValueOrEmptyString(resource.Status),
		StatusMessage:     dcl.ValueOrEmptyString(resource.StatusMessage),
		Autoscaling:       ContainerBetaNodePoolAutoscalingToProto(resource.Autoscaling),
		Management:        ContainerBetaNodePoolManagementToProto(resource.Management),
		MaxPodsConstraint: ContainerBetaNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		PodIpv4CidrSize:   dcl.ValueOrEmptyInt64(resource.PodIPv4CidrSize),
		UpgradeSettings:   ContainerBetaNodePoolUpgradeSettingsToProto(resource.UpgradeSettings),
		Cluster:           dcl.ValueOrEmptyString(resource.Cluster),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Locations {
		p.Locations = append(p.Locations, r)
	}
	for _, r := range resource.Conditions {
		p.Conditions = append(p.Conditions, ContainerBetaNodePoolConditionsToProto(&r))
	}

	return p
}

// ApplyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *beta.Client, request *betapb.ApplyContainerBetaNodePoolRequest) (*betapb.ContainerBetaNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// ApplyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerBetaNodePool(ctx context.Context, request *betapb.ApplyContainerBetaNodePoolRequest) (*betapb.ContainerBetaNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerBetaNodePool(ctx context.Context, request *betapb.DeleteContainerBetaNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerBetaNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerBetaNodePool(ctx context.Context, request *betapb.ListContainerBetaNodePoolRequest) (*betapb.ListContainerBetaNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.Project, request.Location, request.Cluster)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContainerBetaNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListContainerBetaNodePoolResponse{Items: protos}, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
