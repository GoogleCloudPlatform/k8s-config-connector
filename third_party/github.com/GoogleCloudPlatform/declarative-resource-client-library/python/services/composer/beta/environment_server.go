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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/composer/beta/composer_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/composer/beta"
)

// Server implements the gRPC interface for Environment.
type EnvironmentServer struct{}

// ProtoToEnvironmentStateEnum converts a EnvironmentStateEnum enum from its proto representation.
func ProtoToComposerBetaEnvironmentStateEnum(e betapb.ComposerBetaEnvironmentStateEnum) *beta.EnvironmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComposerBetaEnvironmentStateEnum_name[int32(e)]; ok {
		e := beta.EnvironmentStateEnum(n[len("ComposerBetaEnvironmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvironmentConfig converts a EnvironmentConfig resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfig(p *betapb.ComposerBetaEnvironmentConfig) *beta.EnvironmentConfig {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfig{
		GkeCluster:                    dcl.StringOrNil(p.GkeCluster),
		DagGcsPrefix:                  dcl.StringOrNil(p.DagGcsPrefix),
		NodeCount:                     dcl.Int64OrNil(p.NodeCount),
		SoftwareConfig:                ProtoToComposerBetaEnvironmentConfigSoftwareConfig(p.GetSoftwareConfig()),
		NodeConfig:                    ProtoToComposerBetaEnvironmentConfigNodeConfig(p.GetNodeConfig()),
		PrivateEnvironmentConfig:      ProtoToComposerBetaEnvironmentConfigPrivateEnvironmentConfig(p.GetPrivateEnvironmentConfig()),
		WebServerNetworkAccessControl: ProtoToComposerBetaEnvironmentConfigWebServerNetworkAccessControl(p.GetWebServerNetworkAccessControl()),
		DatabaseConfig:                ProtoToComposerBetaEnvironmentConfigDatabaseConfig(p.GetDatabaseConfig()),
		WebServerConfig:               ProtoToComposerBetaEnvironmentConfigWebServerConfig(p.GetWebServerConfig()),
		EncryptionConfig:              ProtoToComposerBetaEnvironmentConfigEncryptionConfig(p.GetEncryptionConfig()),
		AirflowUri:                    dcl.StringOrNil(p.AirflowUri),
		MaintenanceWindow:             ProtoToComposerBetaEnvironmentConfigMaintenanceWindow(p.GetMaintenanceWindow()),
	}
	return obj
}

// ProtoToEnvironmentConfigSoftwareConfig converts a EnvironmentConfigSoftwareConfig resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigSoftwareConfig(p *betapb.ComposerBetaEnvironmentConfigSoftwareConfig) *beta.EnvironmentConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigSoftwareConfig{
		ImageVersion:  dcl.StringOrNil(p.ImageVersion),
		PythonVersion: dcl.StringOrNil(p.PythonVersion),
	}
	return obj
}

// ProtoToEnvironmentConfigNodeConfig converts a EnvironmentConfigNodeConfig resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigNodeConfig(p *betapb.ComposerBetaEnvironmentConfigNodeConfig) *beta.EnvironmentConfigNodeConfig {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigNodeConfig{
		Location:           dcl.StringOrNil(p.Location),
		MachineType:        dcl.StringOrNil(p.MachineType),
		Network:            dcl.StringOrNil(p.Network),
		Subnetwork:         dcl.StringOrNil(p.Subnetwork),
		DiskSizeGb:         dcl.Int64OrNil(p.DiskSizeGb),
		ServiceAccount:     dcl.StringOrNil(p.ServiceAccount),
		IPAllocationPolicy: ProtoToComposerBetaEnvironmentConfigNodeConfigIPAllocationPolicy(p.GetIpAllocationPolicy()),
		MaxPodsPerNode:     dcl.Int64OrNil(p.MaxPodsPerNode),
	}
	for _, r := range p.GetOauthScopes() {
		obj.OAuthScopes = append(obj.OAuthScopes, r)
	}
	for _, r := range p.GetTags() {
		obj.Tags = append(obj.Tags, r)
	}
	return obj
}

// ProtoToEnvironmentConfigNodeConfigIPAllocationPolicy converts a EnvironmentConfigNodeConfigIPAllocationPolicy resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigNodeConfigIPAllocationPolicy(p *betapb.ComposerBetaEnvironmentConfigNodeConfigIPAllocationPolicy) *beta.EnvironmentConfigNodeConfigIPAllocationPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigNodeConfigIPAllocationPolicy{
		UseIPAliases:               dcl.Bool(p.UseIpAliases),
		ClusterSecondaryRangeName:  dcl.StringOrNil(p.ClusterSecondaryRangeName),
		ClusterIPv4CidrBlock:       dcl.StringOrNil(p.ClusterIpv4CidrBlock),
		ServicesSecondaryRangeName: dcl.StringOrNil(p.ServicesSecondaryRangeName),
		ServicesIPv4CidrBlock:      dcl.StringOrNil(p.ServicesIpv4CidrBlock),
	}
	return obj
}

// ProtoToEnvironmentConfigPrivateEnvironmentConfig converts a EnvironmentConfigPrivateEnvironmentConfig resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigPrivateEnvironmentConfig(p *betapb.ComposerBetaEnvironmentConfigPrivateEnvironmentConfig) *beta.EnvironmentConfigPrivateEnvironmentConfig {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigPrivateEnvironmentConfig{
		EnablePrivateEnvironment:   dcl.Bool(p.EnablePrivateEnvironment),
		PrivateClusterConfig:       ProtoToComposerBetaEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(p.GetPrivateClusterConfig()),
		WebServerIPv4CidrBlock:     dcl.StringOrNil(p.WebServerIpv4CidrBlock),
		CloudSqlIPv4CidrBlock:      dcl.StringOrNil(p.CloudSqlIpv4CidrBlock),
		WebServerIPv4ReservedRange: dcl.StringOrNil(p.WebServerIpv4ReservedRange),
	}
	return obj
}

// ProtoToEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig converts a EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(p *betapb.ComposerBetaEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) *beta.EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{
		EnablePrivateEndpoint:   dcl.Bool(p.EnablePrivateEndpoint),
		MasterIPv4CidrBlock:     dcl.StringOrNil(p.MasterIpv4CidrBlock),
		MasterIPv4ReservedRange: dcl.StringOrNil(p.MasterIpv4ReservedRange),
	}
	return obj
}

// ProtoToEnvironmentConfigWebServerNetworkAccessControl converts a EnvironmentConfigWebServerNetworkAccessControl resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigWebServerNetworkAccessControl(p *betapb.ComposerBetaEnvironmentConfigWebServerNetworkAccessControl) *beta.EnvironmentConfigWebServerNetworkAccessControl {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigWebServerNetworkAccessControl{}
	for _, r := range p.GetAllowedIpRanges() {
		obj.AllowedIPRanges = append(obj.AllowedIPRanges, *ProtoToComposerBetaEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(r))
	}
	return obj
}

// ProtoToEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges converts a EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(p *betapb.ComposerBetaEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) *beta.EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{
		Value:       dcl.StringOrNil(p.Value),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToEnvironmentConfigDatabaseConfig converts a EnvironmentConfigDatabaseConfig resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigDatabaseConfig(p *betapb.ComposerBetaEnvironmentConfigDatabaseConfig) *beta.EnvironmentConfigDatabaseConfig {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigDatabaseConfig{
		MachineType: dcl.StringOrNil(p.MachineType),
	}
	return obj
}

// ProtoToEnvironmentConfigWebServerConfig converts a EnvironmentConfigWebServerConfig resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigWebServerConfig(p *betapb.ComposerBetaEnvironmentConfigWebServerConfig) *beta.EnvironmentConfigWebServerConfig {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigWebServerConfig{
		MachineType: dcl.StringOrNil(p.MachineType),
	}
	return obj
}

// ProtoToEnvironmentConfigEncryptionConfig converts a EnvironmentConfigEncryptionConfig resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigEncryptionConfig(p *betapb.ComposerBetaEnvironmentConfigEncryptionConfig) *beta.EnvironmentConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigEncryptionConfig{
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToEnvironmentConfigMaintenanceWindow converts a EnvironmentConfigMaintenanceWindow resource from its proto representation.
func ProtoToComposerBetaEnvironmentConfigMaintenanceWindow(p *betapb.ComposerBetaEnvironmentConfigMaintenanceWindow) *beta.EnvironmentConfigMaintenanceWindow {
	if p == nil {
		return nil
	}
	obj := &beta.EnvironmentConfigMaintenanceWindow{
		StartTime:  dcl.StringOrNil(p.GetStartTime()),
		EndTime:    dcl.StringOrNil(p.GetEndTime()),
		Recurrence: dcl.StringOrNil(p.Recurrence),
	}
	return obj
}

// ProtoToEnvironment converts a Environment resource from its proto representation.
func ProtoToEnvironment(p *betapb.ComposerBetaEnvironment) *beta.Environment {
	obj := &beta.Environment{
		Name:       dcl.StringOrNil(p.Name),
		Config:     ProtoToComposerBetaEnvironmentConfig(p.GetConfig()),
		Uuid:       dcl.StringOrNil(p.Uuid),
		State:      ProtoToComposerBetaEnvironmentStateEnum(p.GetState()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Project:    dcl.StringOrNil(p.Project),
		Location:   dcl.StringOrNil(p.Location),
	}
	return obj
}

// EnvironmentStateEnumToProto converts a EnvironmentStateEnum enum to its proto representation.
func ComposerBetaEnvironmentStateEnumToProto(e *beta.EnvironmentStateEnum) betapb.ComposerBetaEnvironmentStateEnum {
	if e == nil {
		return betapb.ComposerBetaEnvironmentStateEnum(0)
	}
	if v, ok := betapb.ComposerBetaEnvironmentStateEnum_value["EnvironmentStateEnum"+string(*e)]; ok {
		return betapb.ComposerBetaEnvironmentStateEnum(v)
	}
	return betapb.ComposerBetaEnvironmentStateEnum(0)
}

// EnvironmentConfigToProto converts a EnvironmentConfig resource to its proto representation.
func ComposerBetaEnvironmentConfigToProto(o *beta.EnvironmentConfig) *betapb.ComposerBetaEnvironmentConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfig{
		GkeCluster:                    dcl.ValueOrEmptyString(o.GkeCluster),
		DagGcsPrefix:                  dcl.ValueOrEmptyString(o.DagGcsPrefix),
		NodeCount:                     dcl.ValueOrEmptyInt64(o.NodeCount),
		SoftwareConfig:                ComposerBetaEnvironmentConfigSoftwareConfigToProto(o.SoftwareConfig),
		NodeConfig:                    ComposerBetaEnvironmentConfigNodeConfigToProto(o.NodeConfig),
		PrivateEnvironmentConfig:      ComposerBetaEnvironmentConfigPrivateEnvironmentConfigToProto(o.PrivateEnvironmentConfig),
		WebServerNetworkAccessControl: ComposerBetaEnvironmentConfigWebServerNetworkAccessControlToProto(o.WebServerNetworkAccessControl),
		DatabaseConfig:                ComposerBetaEnvironmentConfigDatabaseConfigToProto(o.DatabaseConfig),
		WebServerConfig:               ComposerBetaEnvironmentConfigWebServerConfigToProto(o.WebServerConfig),
		EncryptionConfig:              ComposerBetaEnvironmentConfigEncryptionConfigToProto(o.EncryptionConfig),
		AirflowUri:                    dcl.ValueOrEmptyString(o.AirflowUri),
		MaintenanceWindow:             ComposerBetaEnvironmentConfigMaintenanceWindowToProto(o.MaintenanceWindow),
	}
	return p
}

// EnvironmentConfigSoftwareConfigToProto converts a EnvironmentConfigSoftwareConfig resource to its proto representation.
func ComposerBetaEnvironmentConfigSoftwareConfigToProto(o *beta.EnvironmentConfigSoftwareConfig) *betapb.ComposerBetaEnvironmentConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigSoftwareConfig{
		ImageVersion:  dcl.ValueOrEmptyString(o.ImageVersion),
		PythonVersion: dcl.ValueOrEmptyString(o.PythonVersion),
	}
	p.AirflowConfigOverrides = make(map[string]string)
	for k, r := range o.AirflowConfigOverrides {
		p.AirflowConfigOverrides[k] = r
	}
	p.PypiPackages = make(map[string]string)
	for k, r := range o.PypiPackages {
		p.PypiPackages[k] = r
	}
	p.EnvVariables = make(map[string]string)
	for k, r := range o.EnvVariables {
		p.EnvVariables[k] = r
	}
	return p
}

// EnvironmentConfigNodeConfigToProto converts a EnvironmentConfigNodeConfig resource to its proto representation.
func ComposerBetaEnvironmentConfigNodeConfigToProto(o *beta.EnvironmentConfigNodeConfig) *betapb.ComposerBetaEnvironmentConfigNodeConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigNodeConfig{
		Location:           dcl.ValueOrEmptyString(o.Location),
		MachineType:        dcl.ValueOrEmptyString(o.MachineType),
		Network:            dcl.ValueOrEmptyString(o.Network),
		Subnetwork:         dcl.ValueOrEmptyString(o.Subnetwork),
		DiskSizeGb:         dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		ServiceAccount:     dcl.ValueOrEmptyString(o.ServiceAccount),
		IpAllocationPolicy: ComposerBetaEnvironmentConfigNodeConfigIPAllocationPolicyToProto(o.IPAllocationPolicy),
		MaxPodsPerNode:     dcl.ValueOrEmptyInt64(o.MaxPodsPerNode),
	}
	for _, r := range o.OAuthScopes {
		p.OauthScopes = append(p.OauthScopes, r)
	}
	for _, r := range o.Tags {
		p.Tags = append(p.Tags, r)
	}
	return p
}

// EnvironmentConfigNodeConfigIPAllocationPolicyToProto converts a EnvironmentConfigNodeConfigIPAllocationPolicy resource to its proto representation.
func ComposerBetaEnvironmentConfigNodeConfigIPAllocationPolicyToProto(o *beta.EnvironmentConfigNodeConfigIPAllocationPolicy) *betapb.ComposerBetaEnvironmentConfigNodeConfigIPAllocationPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigNodeConfigIPAllocationPolicy{
		UseIpAliases:               dcl.ValueOrEmptyBool(o.UseIPAliases),
		ClusterSecondaryRangeName:  dcl.ValueOrEmptyString(o.ClusterSecondaryRangeName),
		ClusterIpv4CidrBlock:       dcl.ValueOrEmptyString(o.ClusterIPv4CidrBlock),
		ServicesSecondaryRangeName: dcl.ValueOrEmptyString(o.ServicesSecondaryRangeName),
		ServicesIpv4CidrBlock:      dcl.ValueOrEmptyString(o.ServicesIPv4CidrBlock),
	}
	return p
}

// EnvironmentConfigPrivateEnvironmentConfigToProto converts a EnvironmentConfigPrivateEnvironmentConfig resource to its proto representation.
func ComposerBetaEnvironmentConfigPrivateEnvironmentConfigToProto(o *beta.EnvironmentConfigPrivateEnvironmentConfig) *betapb.ComposerBetaEnvironmentConfigPrivateEnvironmentConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigPrivateEnvironmentConfig{
		EnablePrivateEnvironment:   dcl.ValueOrEmptyBool(o.EnablePrivateEnvironment),
		PrivateClusterConfig:       ComposerBetaEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigToProto(o.PrivateClusterConfig),
		WebServerIpv4CidrBlock:     dcl.ValueOrEmptyString(o.WebServerIPv4CidrBlock),
		CloudSqlIpv4CidrBlock:      dcl.ValueOrEmptyString(o.CloudSqlIPv4CidrBlock),
		WebServerIpv4ReservedRange: dcl.ValueOrEmptyString(o.WebServerIPv4ReservedRange),
	}
	return p
}

// EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigToProto converts a EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig resource to its proto representation.
func ComposerBetaEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigToProto(o *beta.EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) *betapb.ComposerBetaEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{
		EnablePrivateEndpoint:   dcl.ValueOrEmptyBool(o.EnablePrivateEndpoint),
		MasterIpv4CidrBlock:     dcl.ValueOrEmptyString(o.MasterIPv4CidrBlock),
		MasterIpv4ReservedRange: dcl.ValueOrEmptyString(o.MasterIPv4ReservedRange),
	}
	return p
}

// EnvironmentConfigWebServerNetworkAccessControlToProto converts a EnvironmentConfigWebServerNetworkAccessControl resource to its proto representation.
func ComposerBetaEnvironmentConfigWebServerNetworkAccessControlToProto(o *beta.EnvironmentConfigWebServerNetworkAccessControl) *betapb.ComposerBetaEnvironmentConfigWebServerNetworkAccessControl {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigWebServerNetworkAccessControl{}
	for _, r := range o.AllowedIPRanges {
		p.AllowedIpRanges = append(p.AllowedIpRanges, ComposerBetaEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesToProto(&r))
	}
	return p
}

// EnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesToProto converts a EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges resource to its proto representation.
func ComposerBetaEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesToProto(o *beta.EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) *betapb.ComposerBetaEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{
		Value:       dcl.ValueOrEmptyString(o.Value),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// EnvironmentConfigDatabaseConfigToProto converts a EnvironmentConfigDatabaseConfig resource to its proto representation.
func ComposerBetaEnvironmentConfigDatabaseConfigToProto(o *beta.EnvironmentConfigDatabaseConfig) *betapb.ComposerBetaEnvironmentConfigDatabaseConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigDatabaseConfig{
		MachineType: dcl.ValueOrEmptyString(o.MachineType),
	}
	return p
}

// EnvironmentConfigWebServerConfigToProto converts a EnvironmentConfigWebServerConfig resource to its proto representation.
func ComposerBetaEnvironmentConfigWebServerConfigToProto(o *beta.EnvironmentConfigWebServerConfig) *betapb.ComposerBetaEnvironmentConfigWebServerConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigWebServerConfig{
		MachineType: dcl.ValueOrEmptyString(o.MachineType),
	}
	return p
}

// EnvironmentConfigEncryptionConfigToProto converts a EnvironmentConfigEncryptionConfig resource to its proto representation.
func ComposerBetaEnvironmentConfigEncryptionConfigToProto(o *beta.EnvironmentConfigEncryptionConfig) *betapb.ComposerBetaEnvironmentConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigEncryptionConfig{
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	return p
}

// EnvironmentConfigMaintenanceWindowToProto converts a EnvironmentConfigMaintenanceWindow resource to its proto representation.
func ComposerBetaEnvironmentConfigMaintenanceWindowToProto(o *beta.EnvironmentConfigMaintenanceWindow) *betapb.ComposerBetaEnvironmentConfigMaintenanceWindow {
	if o == nil {
		return nil
	}
	p := &betapb.ComposerBetaEnvironmentConfigMaintenanceWindow{
		StartTime:  dcl.ValueOrEmptyString(o.StartTime),
		EndTime:    dcl.ValueOrEmptyString(o.EndTime),
		Recurrence: dcl.ValueOrEmptyString(o.Recurrence),
	}
	return p
}

// EnvironmentToProto converts a Environment resource to its proto representation.
func EnvironmentToProto(resource *beta.Environment) *betapb.ComposerBetaEnvironment {
	p := &betapb.ComposerBetaEnvironment{
		Name:       dcl.ValueOrEmptyString(resource.Name),
		Config:     ComposerBetaEnvironmentConfigToProto(resource.Config),
		Uuid:       dcl.ValueOrEmptyString(resource.Uuid),
		State:      ComposerBetaEnvironmentStateEnumToProto(resource.State),
		CreateTime: dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime: dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:    dcl.ValueOrEmptyString(resource.Project),
		Location:   dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) applyEnvironment(ctx context.Context, c *beta.Client, request *betapb.ApplyComposerBetaEnvironmentRequest) (*betapb.ComposerBetaEnvironment, error) {
	p := ProtoToEnvironment(request.GetResource())
	res, err := c.ApplyEnvironment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvironmentToProto(res)
	return r, nil
}

// ApplyEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) ApplyComposerBetaEnvironment(ctx context.Context, request *betapb.ApplyComposerBetaEnvironmentRequest) (*betapb.ComposerBetaEnvironment, error) {
	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyEnvironment(ctx, cl, request)
}

// DeleteEnvironment handles the gRPC request by passing it to the underlying Environment Delete() method.
func (s *EnvironmentServer) DeleteComposerBetaEnvironment(ctx context.Context, request *betapb.DeleteComposerBetaEnvironmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvironment(ctx, ProtoToEnvironment(request.GetResource()))

}

// ListComposerBetaEnvironment handles the gRPC request by passing it to the underlying EnvironmentList() method.
func (s *EnvironmentServer) ListComposerBetaEnvironment(ctx context.Context, request *betapb.ListComposerBetaEnvironmentRequest) (*betapb.ListComposerBetaEnvironmentResponse, error) {
	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvironment(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComposerBetaEnvironment
	for _, r := range resources.Items {
		rp := EnvironmentToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComposerBetaEnvironmentResponse{Items: protos}, nil
}

func createConfigEnvironment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
