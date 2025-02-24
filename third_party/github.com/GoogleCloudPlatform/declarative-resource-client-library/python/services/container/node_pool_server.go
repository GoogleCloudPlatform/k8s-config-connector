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
	containerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/container/container_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/container"
)

// Server implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolConfigSandboxConfigTypeEnum converts a NodePoolConfigSandboxConfigTypeEnum enum from its proto representation.
func ProtoToContainerNodePoolConfigSandboxConfigTypeEnum(e containerpb.ContainerNodePoolConfigSandboxConfigTypeEnum) *container.NodePoolConfigSandboxConfigTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerNodePoolConfigSandboxConfigTypeEnum_name[int32(e)]; ok {
		e := container.NodePoolConfigSandboxConfigTypeEnum(n[len("ContainerNodePoolConfigSandboxConfigTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfigReservationAffinityConsumeReservationTypeEnum converts a NodePoolConfigReservationAffinityConsumeReservationTypeEnum enum from its proto representation.
func ProtoToContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum(e containerpb.ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum) *container.NodePoolConfigReservationAffinityConsumeReservationTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum_name[int32(e)]; ok {
		e := container.NodePoolConfigReservationAffinityConsumeReservationTypeEnum(n[len("ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConditionsCodeEnum converts a NodePoolConditionsCodeEnum enum from its proto representation.
func ProtoToContainerNodePoolConditionsCodeEnum(e containerpb.ContainerNodePoolConditionsCodeEnum) *container.NodePoolConditionsCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerpb.ContainerNodePoolConditionsCodeEnum_name[int32(e)]; ok {
		e := container.NodePoolConditionsCodeEnum(n[len("ContainerNodePoolConditionsCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig resource from its proto representation.
func ProtoToContainerNodePoolConfig(p *containerpb.ContainerNodePoolConfig) *container.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolConfig{
		MachineType:            dcl.StringOrNil(p.MachineType),
		DiskSizeGb:             dcl.Int64OrNil(p.DiskSizeGb),
		ServiceAccount:         dcl.StringOrNil(p.ServiceAccount),
		ImageType:              dcl.StringOrNil(p.ImageType),
		LocalSsdCount:          dcl.Int64OrNil(p.LocalSsdCount),
		Preemptible:            dcl.Bool(p.Preemptible),
		DiskType:               dcl.StringOrNil(p.DiskType),
		MinCpuPlatform:         dcl.StringOrNil(p.MinCpuPlatform),
		SandboxConfig:          ProtoToContainerNodePoolConfigSandboxConfig(p.GetSandboxConfig()),
		ReservationAffinity:    ProtoToContainerNodePoolConfigReservationAffinity(p.GetReservationAffinity()),
		ShieldedInstanceConfig: ProtoToContainerNodePoolConfigShieldedInstanceConfig(p.GetShieldedInstanceConfig()),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	for _, r := range p.GetAccelerators() {
		obj.Accelerators = append(obj.Accelerators, *ProtoToContainerNodePoolConfigAccelerators(r))
	}
	for _, r := range p.GetTaints() {
		obj.Taints = append(obj.Taints, *ProtoToContainerNodePoolConfigTaints(r))
	}
	return obj
}

// ProtoToNodePoolConfigAccelerators converts a NodePoolConfigAccelerators resource from its proto representation.
func ProtoToContainerNodePoolConfigAccelerators(p *containerpb.ContainerNodePoolConfigAccelerators) *container.NodePoolConfigAccelerators {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolConfigAccelerators{
		AcceleratorCount: dcl.Int64OrNil(p.AcceleratorCount),
		AcceleratorType:  dcl.StringOrNil(p.AcceleratorType),
	}
	return obj
}

// ProtoToNodePoolConfigTaints converts a NodePoolConfigTaints resource from its proto representation.
func ProtoToContainerNodePoolConfigTaints(p *containerpb.ContainerNodePoolConfigTaints) *container.NodePoolConfigTaints {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolConfigTaints{
		Key:    dcl.StringOrNil(p.Key),
		Value:  dcl.StringOrNil(p.Value),
		Effect: dcl.StringOrNil(p.Effect),
	}
	return obj
}

// ProtoToNodePoolConfigSandboxConfig converts a NodePoolConfigSandboxConfig resource from its proto representation.
func ProtoToContainerNodePoolConfigSandboxConfig(p *containerpb.ContainerNodePoolConfigSandboxConfig) *container.NodePoolConfigSandboxConfig {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolConfigSandboxConfig{
		Type: ProtoToContainerNodePoolConfigSandboxConfigTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToNodePoolConfigReservationAffinity converts a NodePoolConfigReservationAffinity resource from its proto representation.
func ProtoToContainerNodePoolConfigReservationAffinity(p *containerpb.ContainerNodePoolConfigReservationAffinity) *container.NodePoolConfigReservationAffinity {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolConfigReservationAffinity{
		ConsumeReservationType: ProtoToContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum(p.GetConsumeReservationType()),
		Key:                    dcl.StringOrNil(p.Key),
	}
	for _, r := range p.GetValues() {
		obj.Values = append(obj.Values, r)
	}
	return obj
}

// ProtoToNodePoolConfigShieldedInstanceConfig converts a NodePoolConfigShieldedInstanceConfig resource from its proto representation.
func ProtoToContainerNodePoolConfigShieldedInstanceConfig(p *containerpb.ContainerNodePoolConfigShieldedInstanceConfig) *container.NodePoolConfigShieldedInstanceConfig {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.Bool(p.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.Bool(p.EnableIntegrityMonitoring),
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling resource from its proto representation.
func ProtoToContainerNodePoolAutoscaling(p *containerpb.ContainerNodePoolAutoscaling) *container.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolAutoscaling{
		Enabled:         dcl.Bool(p.Enabled),
		MinNodeCount:    dcl.Int64OrNil(p.MinNodeCount),
		MaxNodeCount:    dcl.Int64OrNil(p.MaxNodeCount),
		Autoprovisioned: dcl.Bool(p.Autoprovisioned),
	}
	return obj
}

// ProtoToNodePoolManagement converts a NodePoolManagement resource from its proto representation.
func ProtoToContainerNodePoolManagement(p *containerpb.ContainerNodePoolManagement) *container.NodePoolManagement {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolManagement{
		AutoUpgrade:    dcl.Bool(p.AutoUpgrade),
		AutoRepair:     dcl.Bool(p.AutoRepair),
		UpgradeOptions: ProtoToContainerNodePoolManagementUpgradeOptions(p.GetUpgradeOptions()),
	}
	return obj
}

// ProtoToNodePoolManagementUpgradeOptions converts a NodePoolManagementUpgradeOptions resource from its proto representation.
func ProtoToContainerNodePoolManagementUpgradeOptions(p *containerpb.ContainerNodePoolManagementUpgradeOptions) *container.NodePoolManagementUpgradeOptions {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.StringOrNil(p.GetAutoUpgradeStartTime()),
		Description:          dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint resource from its proto representation.
func ProtoToContainerNodePoolMaxPodsConstraint(p *containerpb.ContainerNodePoolMaxPodsConstraint) *container.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	return obj
}

// ProtoToNodePoolConditions converts a NodePoolConditions resource from its proto representation.
func ProtoToContainerNodePoolConditions(p *containerpb.ContainerNodePoolConditions) *container.NodePoolConditions {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolConditions{
		Code:    ProtoToContainerNodePoolConditionsCodeEnum(p.GetCode()),
		Message: dcl.StringOrNil(p.Message),
	}
	return obj
}

// ProtoToNodePoolUpgradeSettings converts a NodePoolUpgradeSettings resource from its proto representation.
func ProtoToContainerNodePoolUpgradeSettings(p *containerpb.ContainerNodePoolUpgradeSettings) *container.NodePoolUpgradeSettings {
	if p == nil {
		return nil
	}
	obj := &container.NodePoolUpgradeSettings{
		MaxSurge:       dcl.Int64OrNil(p.MaxSurge),
		MaxUnavailable: dcl.Int64OrNil(p.MaxUnavailable),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *containerpb.ContainerNodePool) *container.NodePool {
	obj := &container.NodePool{
		Name:              dcl.StringOrNil(p.Name),
		Config:            ProtoToContainerNodePoolConfig(p.GetConfig()),
		NodeCount:         dcl.Int64OrNil(p.NodeCount),
		Version:           dcl.StringOrNil(p.Version),
		Status:            dcl.StringOrNil(p.Status),
		StatusMessage:     dcl.StringOrNil(p.StatusMessage),
		Autoscaling:       ProtoToContainerNodePoolAutoscaling(p.GetAutoscaling()),
		Management:        ProtoToContainerNodePoolManagement(p.GetManagement()),
		MaxPodsConstraint: ProtoToContainerNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		PodIPv4CidrSize:   dcl.Int64OrNil(p.PodIpv4CidrSize),
		UpgradeSettings:   ProtoToContainerNodePoolUpgradeSettings(p.GetUpgradeSettings()),
		Cluster:           dcl.StringOrNil(p.Cluster),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetLocations() {
		obj.Locations = append(obj.Locations, r)
	}
	for _, r := range p.GetConditions() {
		obj.Conditions = append(obj.Conditions, *ProtoToContainerNodePoolConditions(r))
	}
	return obj
}

// NodePoolConfigSandboxConfigTypeEnumToProto converts a NodePoolConfigSandboxConfigTypeEnum enum to its proto representation.
func ContainerNodePoolConfigSandboxConfigTypeEnumToProto(e *container.NodePoolConfigSandboxConfigTypeEnum) containerpb.ContainerNodePoolConfigSandboxConfigTypeEnum {
	if e == nil {
		return containerpb.ContainerNodePoolConfigSandboxConfigTypeEnum(0)
	}
	if v, ok := containerpb.ContainerNodePoolConfigSandboxConfigTypeEnum_value["NodePoolConfigSandboxConfigTypeEnum"+string(*e)]; ok {
		return containerpb.ContainerNodePoolConfigSandboxConfigTypeEnum(v)
	}
	return containerpb.ContainerNodePoolConfigSandboxConfigTypeEnum(0)
}

// NodePoolConfigReservationAffinityConsumeReservationTypeEnumToProto converts a NodePoolConfigReservationAffinityConsumeReservationTypeEnum enum to its proto representation.
func ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnumToProto(e *container.NodePoolConfigReservationAffinityConsumeReservationTypeEnum) containerpb.ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum {
	if e == nil {
		return containerpb.ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum(0)
	}
	if v, ok := containerpb.ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum_value["NodePoolConfigReservationAffinityConsumeReservationTypeEnum"+string(*e)]; ok {
		return containerpb.ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum(v)
	}
	return containerpb.ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnum(0)
}

// NodePoolConditionsCodeEnumToProto converts a NodePoolConditionsCodeEnum enum to its proto representation.
func ContainerNodePoolConditionsCodeEnumToProto(e *container.NodePoolConditionsCodeEnum) containerpb.ContainerNodePoolConditionsCodeEnum {
	if e == nil {
		return containerpb.ContainerNodePoolConditionsCodeEnum(0)
	}
	if v, ok := containerpb.ContainerNodePoolConditionsCodeEnum_value["NodePoolConditionsCodeEnum"+string(*e)]; ok {
		return containerpb.ContainerNodePoolConditionsCodeEnum(v)
	}
	return containerpb.ContainerNodePoolConditionsCodeEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig resource to its proto representation.
func ContainerNodePoolConfigToProto(o *container.NodePoolConfig) *containerpb.ContainerNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolConfig{
		MachineType:            dcl.ValueOrEmptyString(o.MachineType),
		DiskSizeGb:             dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		ServiceAccount:         dcl.ValueOrEmptyString(o.ServiceAccount),
		ImageType:              dcl.ValueOrEmptyString(o.ImageType),
		LocalSsdCount:          dcl.ValueOrEmptyInt64(o.LocalSsdCount),
		Preemptible:            dcl.ValueOrEmptyBool(o.Preemptible),
		DiskType:               dcl.ValueOrEmptyString(o.DiskType),
		MinCpuPlatform:         dcl.ValueOrEmptyString(o.MinCpuPlatform),
		SandboxConfig:          ContainerNodePoolConfigSandboxConfigToProto(o.SandboxConfig),
		ReservationAffinity:    ContainerNodePoolConfigReservationAffinityToProto(o.ReservationAffinity),
		ShieldedInstanceConfig: ContainerNodePoolConfigShieldedInstanceConfigToProto(o.ShieldedInstanceConfig),
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
		p.Accelerators = append(p.Accelerators, ContainerNodePoolConfigAcceleratorsToProto(&r))
	}
	for _, r := range o.Taints {
		p.Taints = append(p.Taints, ContainerNodePoolConfigTaintsToProto(&r))
	}
	return p
}

// NodePoolConfigAcceleratorsToProto converts a NodePoolConfigAccelerators resource to its proto representation.
func ContainerNodePoolConfigAcceleratorsToProto(o *container.NodePoolConfigAccelerators) *containerpb.ContainerNodePoolConfigAccelerators {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolConfigAccelerators{
		AcceleratorCount: dcl.ValueOrEmptyInt64(o.AcceleratorCount),
		AcceleratorType:  dcl.ValueOrEmptyString(o.AcceleratorType),
	}
	return p
}

// NodePoolConfigTaintsToProto converts a NodePoolConfigTaints resource to its proto representation.
func ContainerNodePoolConfigTaintsToProto(o *container.NodePoolConfigTaints) *containerpb.ContainerNodePoolConfigTaints {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolConfigTaints{
		Key:    dcl.ValueOrEmptyString(o.Key),
		Value:  dcl.ValueOrEmptyString(o.Value),
		Effect: dcl.ValueOrEmptyString(o.Effect),
	}
	return p
}

// NodePoolConfigSandboxConfigToProto converts a NodePoolConfigSandboxConfig resource to its proto representation.
func ContainerNodePoolConfigSandboxConfigToProto(o *container.NodePoolConfigSandboxConfig) *containerpb.ContainerNodePoolConfigSandboxConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolConfigSandboxConfig{
		Type: ContainerNodePoolConfigSandboxConfigTypeEnumToProto(o.Type),
	}
	return p
}

// NodePoolConfigReservationAffinityToProto converts a NodePoolConfigReservationAffinity resource to its proto representation.
func ContainerNodePoolConfigReservationAffinityToProto(o *container.NodePoolConfigReservationAffinity) *containerpb.ContainerNodePoolConfigReservationAffinity {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolConfigReservationAffinity{
		ConsumeReservationType: ContainerNodePoolConfigReservationAffinityConsumeReservationTypeEnumToProto(o.ConsumeReservationType),
		Key:                    dcl.ValueOrEmptyString(o.Key),
	}
	for _, r := range o.Values {
		p.Values = append(p.Values, r)
	}
	return p
}

// NodePoolConfigShieldedInstanceConfigToProto converts a NodePoolConfigShieldedInstanceConfig resource to its proto representation.
func ContainerNodePoolConfigShieldedInstanceConfigToProto(o *container.NodePoolConfigShieldedInstanceConfig) *containerpb.ContainerNodePoolConfigShieldedInstanceConfig {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolConfigShieldedInstanceConfig{
		EnableSecureBoot:          dcl.ValueOrEmptyBool(o.EnableSecureBoot),
		EnableIntegrityMonitoring: dcl.ValueOrEmptyBool(o.EnableIntegrityMonitoring),
	}
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling resource to its proto representation.
func ContainerNodePoolAutoscalingToProto(o *container.NodePoolAutoscaling) *containerpb.ContainerNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolAutoscaling{
		Enabled:         dcl.ValueOrEmptyBool(o.Enabled),
		MinNodeCount:    dcl.ValueOrEmptyInt64(o.MinNodeCount),
		MaxNodeCount:    dcl.ValueOrEmptyInt64(o.MaxNodeCount),
		Autoprovisioned: dcl.ValueOrEmptyBool(o.Autoprovisioned),
	}
	return p
}

// NodePoolManagementToProto converts a NodePoolManagement resource to its proto representation.
func ContainerNodePoolManagementToProto(o *container.NodePoolManagement) *containerpb.ContainerNodePoolManagement {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolManagement{
		AutoUpgrade:    dcl.ValueOrEmptyBool(o.AutoUpgrade),
		AutoRepair:     dcl.ValueOrEmptyBool(o.AutoRepair),
		UpgradeOptions: ContainerNodePoolManagementUpgradeOptionsToProto(o.UpgradeOptions),
	}
	return p
}

// NodePoolManagementUpgradeOptionsToProto converts a NodePoolManagementUpgradeOptions resource to its proto representation.
func ContainerNodePoolManagementUpgradeOptionsToProto(o *container.NodePoolManagementUpgradeOptions) *containerpb.ContainerNodePoolManagementUpgradeOptions {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolManagementUpgradeOptions{
		AutoUpgradeStartTime: dcl.ValueOrEmptyString(o.AutoUpgradeStartTime),
		Description:          dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint resource to its proto representation.
func ContainerNodePoolMaxPodsConstraintToProto(o *container.NodePoolMaxPodsConstraint) *containerpb.ContainerNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	return p
}

// NodePoolConditionsToProto converts a NodePoolConditions resource to its proto representation.
func ContainerNodePoolConditionsToProto(o *container.NodePoolConditions) *containerpb.ContainerNodePoolConditions {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolConditions{
		Code:    ContainerNodePoolConditionsCodeEnumToProto(o.Code),
		Message: dcl.ValueOrEmptyString(o.Message),
	}
	return p
}

// NodePoolUpgradeSettingsToProto converts a NodePoolUpgradeSettings resource to its proto representation.
func ContainerNodePoolUpgradeSettingsToProto(o *container.NodePoolUpgradeSettings) *containerpb.ContainerNodePoolUpgradeSettings {
	if o == nil {
		return nil
	}
	p := &containerpb.ContainerNodePoolUpgradeSettings{
		MaxSurge:       dcl.ValueOrEmptyInt64(o.MaxSurge),
		MaxUnavailable: dcl.ValueOrEmptyInt64(o.MaxUnavailable),
	}
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *container.NodePool) *containerpb.ContainerNodePool {
	p := &containerpb.ContainerNodePool{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Config:            ContainerNodePoolConfigToProto(resource.Config),
		NodeCount:         dcl.ValueOrEmptyInt64(resource.NodeCount),
		Version:           dcl.ValueOrEmptyString(resource.Version),
		Status:            dcl.ValueOrEmptyString(resource.Status),
		StatusMessage:     dcl.ValueOrEmptyString(resource.StatusMessage),
		Autoscaling:       ContainerNodePoolAutoscalingToProto(resource.Autoscaling),
		Management:        ContainerNodePoolManagementToProto(resource.Management),
		MaxPodsConstraint: ContainerNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint),
		PodIpv4CidrSize:   dcl.ValueOrEmptyInt64(resource.PodIPv4CidrSize),
		UpgradeSettings:   ContainerNodePoolUpgradeSettingsToProto(resource.UpgradeSettings),
		Cluster:           dcl.ValueOrEmptyString(resource.Cluster),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.Locations {
		p.Locations = append(p.Locations, r)
	}
	for _, r := range resource.Conditions {
		p.Conditions = append(p.Conditions, ContainerNodePoolConditionsToProto(&r))
	}

	return p
}

// ApplyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *container.Client, request *containerpb.ApplyContainerNodePoolRequest) (*containerpb.ContainerNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// ApplyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerNodePool(ctx context.Context, request *containerpb.ApplyContainerNodePoolRequest) (*containerpb.ContainerNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerNodePool(ctx context.Context, request *containerpb.DeleteContainerNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerNodePool(ctx context.Context, request *containerpb.ListContainerNodePoolRequest) (*containerpb.ListContainerNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.Project, request.Location, request.Cluster)
	if err != nil {
		return nil, err
	}
	var protos []*containerpb.ContainerNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	return &containerpb.ListContainerNodePoolResponse{Items: protos}, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*container.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return container.NewClient(conf), nil
}
