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
	composerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/composer/composer_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/composer"
)

// Server implements the gRPC interface for Environment.
type EnvironmentServer struct{}

// ProtoToEnvironmentStateEnum converts a EnvironmentStateEnum enum from its proto representation.
func ProtoToComposerEnvironmentStateEnum(e composerpb.ComposerEnvironmentStateEnum) *composer.EnvironmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := composerpb.ComposerEnvironmentStateEnum_name[int32(e)]; ok {
		e := composer.EnvironmentStateEnum(n[len("ComposerEnvironmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvironmentConfig converts a EnvironmentConfig resource from its proto representation.
func ProtoToComposerEnvironmentConfig(p *composerpb.ComposerEnvironmentConfig) *composer.EnvironmentConfig {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfig{
		GkeCluster:                    dcl.StringOrNil(p.GkeCluster),
		DagGcsPrefix:                  dcl.StringOrNil(p.DagGcsPrefix),
		NodeCount:                     dcl.Int64OrNil(p.NodeCount),
		SoftwareConfig:                ProtoToComposerEnvironmentConfigSoftwareConfig(p.GetSoftwareConfig()),
		NodeConfig:                    ProtoToComposerEnvironmentConfigNodeConfig(p.GetNodeConfig()),
		PrivateEnvironmentConfig:      ProtoToComposerEnvironmentConfigPrivateEnvironmentConfig(p.GetPrivateEnvironmentConfig()),
		WebServerNetworkAccessControl: ProtoToComposerEnvironmentConfigWebServerNetworkAccessControl(p.GetWebServerNetworkAccessControl()),
		DatabaseConfig:                ProtoToComposerEnvironmentConfigDatabaseConfig(p.GetDatabaseConfig()),
		WebServerConfig:               ProtoToComposerEnvironmentConfigWebServerConfig(p.GetWebServerConfig()),
		EncryptionConfig:              ProtoToComposerEnvironmentConfigEncryptionConfig(p.GetEncryptionConfig()),
		AirflowUri:                    dcl.StringOrNil(p.AirflowUri),
	}
	return obj
}

// ProtoToEnvironmentConfigSoftwareConfig converts a EnvironmentConfigSoftwareConfig resource from its proto representation.
func ProtoToComposerEnvironmentConfigSoftwareConfig(p *composerpb.ComposerEnvironmentConfigSoftwareConfig) *composer.EnvironmentConfigSoftwareConfig {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigSoftwareConfig{
		ImageVersion:  dcl.StringOrNil(p.ImageVersion),
		PythonVersion: dcl.StringOrNil(p.PythonVersion),
	}
	return obj
}

// ProtoToEnvironmentConfigNodeConfig converts a EnvironmentConfigNodeConfig resource from its proto representation.
func ProtoToComposerEnvironmentConfigNodeConfig(p *composerpb.ComposerEnvironmentConfigNodeConfig) *composer.EnvironmentConfigNodeConfig {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigNodeConfig{
		Location:           dcl.StringOrNil(p.Location),
		MachineType:        dcl.StringOrNil(p.MachineType),
		Network:            dcl.StringOrNil(p.Network),
		Subnetwork:         dcl.StringOrNil(p.Subnetwork),
		DiskSizeGb:         dcl.Int64OrNil(p.DiskSizeGb),
		ServiceAccount:     dcl.StringOrNil(p.ServiceAccount),
		IPAllocationPolicy: ProtoToComposerEnvironmentConfigNodeConfigIPAllocationPolicy(p.GetIpAllocationPolicy()),
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
func ProtoToComposerEnvironmentConfigNodeConfigIPAllocationPolicy(p *composerpb.ComposerEnvironmentConfigNodeConfigIPAllocationPolicy) *composer.EnvironmentConfigNodeConfigIPAllocationPolicy {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigNodeConfigIPAllocationPolicy{
		UseIPAliases:               dcl.Bool(p.UseIpAliases),
		ClusterSecondaryRangeName:  dcl.StringOrNil(p.ClusterSecondaryRangeName),
		ClusterIPv4CidrBlock:       dcl.StringOrNil(p.ClusterIpv4CidrBlock),
		ServicesSecondaryRangeName: dcl.StringOrNil(p.ServicesSecondaryRangeName),
		ServicesIPv4CidrBlock:      dcl.StringOrNil(p.ServicesIpv4CidrBlock),
	}
	return obj
}

// ProtoToEnvironmentConfigPrivateEnvironmentConfig converts a EnvironmentConfigPrivateEnvironmentConfig resource from its proto representation.
func ProtoToComposerEnvironmentConfigPrivateEnvironmentConfig(p *composerpb.ComposerEnvironmentConfigPrivateEnvironmentConfig) *composer.EnvironmentConfigPrivateEnvironmentConfig {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigPrivateEnvironmentConfig{
		EnablePrivateEnvironment:   dcl.Bool(p.EnablePrivateEnvironment),
		PrivateClusterConfig:       ProtoToComposerEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(p.GetPrivateClusterConfig()),
		WebServerIPv4CidrBlock:     dcl.StringOrNil(p.WebServerIpv4CidrBlock),
		CloudSqlIPv4CidrBlock:      dcl.StringOrNil(p.CloudSqlIpv4CidrBlock),
		WebServerIPv4ReservedRange: dcl.StringOrNil(p.WebServerIpv4ReservedRange),
	}
	return obj
}

// ProtoToEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig converts a EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig resource from its proto representation.
func ProtoToComposerEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig(p *composerpb.ComposerEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) *composer.EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{
		EnablePrivateEndpoint:   dcl.Bool(p.EnablePrivateEndpoint),
		MasterIPv4CidrBlock:     dcl.StringOrNil(p.MasterIpv4CidrBlock),
		MasterIPv4ReservedRange: dcl.StringOrNil(p.MasterIpv4ReservedRange),
	}
	return obj
}

// ProtoToEnvironmentConfigWebServerNetworkAccessControl converts a EnvironmentConfigWebServerNetworkAccessControl resource from its proto representation.
func ProtoToComposerEnvironmentConfigWebServerNetworkAccessControl(p *composerpb.ComposerEnvironmentConfigWebServerNetworkAccessControl) *composer.EnvironmentConfigWebServerNetworkAccessControl {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigWebServerNetworkAccessControl{}
	for _, r := range p.GetAllowedIpRanges() {
		obj.AllowedIPRanges = append(obj.AllowedIPRanges, *ProtoToComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(r))
	}
	return obj
}

// ProtoToEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges converts a EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges resource from its proto representation.
func ProtoToComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges(p *composerpb.ComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) *composer.EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{
		Value:       dcl.StringOrNil(p.Value),
		Description: dcl.StringOrNil(p.Description),
	}
	return obj
}

// ProtoToEnvironmentConfigDatabaseConfig converts a EnvironmentConfigDatabaseConfig resource from its proto representation.
func ProtoToComposerEnvironmentConfigDatabaseConfig(p *composerpb.ComposerEnvironmentConfigDatabaseConfig) *composer.EnvironmentConfigDatabaseConfig {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigDatabaseConfig{
		MachineType: dcl.StringOrNil(p.MachineType),
	}
	return obj
}

// ProtoToEnvironmentConfigWebServerConfig converts a EnvironmentConfigWebServerConfig resource from its proto representation.
func ProtoToComposerEnvironmentConfigWebServerConfig(p *composerpb.ComposerEnvironmentConfigWebServerConfig) *composer.EnvironmentConfigWebServerConfig {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigWebServerConfig{
		MachineType: dcl.StringOrNil(p.MachineType),
	}
	return obj
}

// ProtoToEnvironmentConfigEncryptionConfig converts a EnvironmentConfigEncryptionConfig resource from its proto representation.
func ProtoToComposerEnvironmentConfigEncryptionConfig(p *composerpb.ComposerEnvironmentConfigEncryptionConfig) *composer.EnvironmentConfigEncryptionConfig {
	if p == nil {
		return nil
	}
	obj := &composer.EnvironmentConfigEncryptionConfig{
		KmsKeyName: dcl.StringOrNil(p.KmsKeyName),
	}
	return obj
}

// ProtoToEnvironment converts a Environment resource from its proto representation.
func ProtoToEnvironment(p *composerpb.ComposerEnvironment) *composer.Environment {
	obj := &composer.Environment{
		Name:       dcl.StringOrNil(p.Name),
		Config:     ProtoToComposerEnvironmentConfig(p.GetConfig()),
		Uuid:       dcl.StringOrNil(p.Uuid),
		State:      ProtoToComposerEnvironmentStateEnum(p.GetState()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Project:    dcl.StringOrNil(p.Project),
		Location:   dcl.StringOrNil(p.Location),
	}
	return obj
}

// EnvironmentStateEnumToProto converts a EnvironmentStateEnum enum to its proto representation.
func ComposerEnvironmentStateEnumToProto(e *composer.EnvironmentStateEnum) composerpb.ComposerEnvironmentStateEnum {
	if e == nil {
		return composerpb.ComposerEnvironmentStateEnum(0)
	}
	if v, ok := composerpb.ComposerEnvironmentStateEnum_value["EnvironmentStateEnum"+string(*e)]; ok {
		return composerpb.ComposerEnvironmentStateEnum(v)
	}
	return composerpb.ComposerEnvironmentStateEnum(0)
}

// EnvironmentConfigToProto converts a EnvironmentConfig resource to its proto representation.
func ComposerEnvironmentConfigToProto(o *composer.EnvironmentConfig) *composerpb.ComposerEnvironmentConfig {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfig{
		GkeCluster:                    dcl.ValueOrEmptyString(o.GkeCluster),
		DagGcsPrefix:                  dcl.ValueOrEmptyString(o.DagGcsPrefix),
		NodeCount:                     dcl.ValueOrEmptyInt64(o.NodeCount),
		SoftwareConfig:                ComposerEnvironmentConfigSoftwareConfigToProto(o.SoftwareConfig),
		NodeConfig:                    ComposerEnvironmentConfigNodeConfigToProto(o.NodeConfig),
		PrivateEnvironmentConfig:      ComposerEnvironmentConfigPrivateEnvironmentConfigToProto(o.PrivateEnvironmentConfig),
		WebServerNetworkAccessControl: ComposerEnvironmentConfigWebServerNetworkAccessControlToProto(o.WebServerNetworkAccessControl),
		DatabaseConfig:                ComposerEnvironmentConfigDatabaseConfigToProto(o.DatabaseConfig),
		WebServerConfig:               ComposerEnvironmentConfigWebServerConfigToProto(o.WebServerConfig),
		EncryptionConfig:              ComposerEnvironmentConfigEncryptionConfigToProto(o.EncryptionConfig),
		AirflowUri:                    dcl.ValueOrEmptyString(o.AirflowUri),
	}
	return p
}

// EnvironmentConfigSoftwareConfigToProto converts a EnvironmentConfigSoftwareConfig resource to its proto representation.
func ComposerEnvironmentConfigSoftwareConfigToProto(o *composer.EnvironmentConfigSoftwareConfig) *composerpb.ComposerEnvironmentConfigSoftwareConfig {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigSoftwareConfig{
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
func ComposerEnvironmentConfigNodeConfigToProto(o *composer.EnvironmentConfigNodeConfig) *composerpb.ComposerEnvironmentConfigNodeConfig {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigNodeConfig{
		Location:           dcl.ValueOrEmptyString(o.Location),
		MachineType:        dcl.ValueOrEmptyString(o.MachineType),
		Network:            dcl.ValueOrEmptyString(o.Network),
		Subnetwork:         dcl.ValueOrEmptyString(o.Subnetwork),
		DiskSizeGb:         dcl.ValueOrEmptyInt64(o.DiskSizeGb),
		ServiceAccount:     dcl.ValueOrEmptyString(o.ServiceAccount),
		IpAllocationPolicy: ComposerEnvironmentConfigNodeConfigIPAllocationPolicyToProto(o.IPAllocationPolicy),
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
func ComposerEnvironmentConfigNodeConfigIPAllocationPolicyToProto(o *composer.EnvironmentConfigNodeConfigIPAllocationPolicy) *composerpb.ComposerEnvironmentConfigNodeConfigIPAllocationPolicy {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigNodeConfigIPAllocationPolicy{
		UseIpAliases:               dcl.ValueOrEmptyBool(o.UseIPAliases),
		ClusterSecondaryRangeName:  dcl.ValueOrEmptyString(o.ClusterSecondaryRangeName),
		ClusterIpv4CidrBlock:       dcl.ValueOrEmptyString(o.ClusterIPv4CidrBlock),
		ServicesSecondaryRangeName: dcl.ValueOrEmptyString(o.ServicesSecondaryRangeName),
		ServicesIpv4CidrBlock:      dcl.ValueOrEmptyString(o.ServicesIPv4CidrBlock),
	}
	return p
}

// EnvironmentConfigPrivateEnvironmentConfigToProto converts a EnvironmentConfigPrivateEnvironmentConfig resource to its proto representation.
func ComposerEnvironmentConfigPrivateEnvironmentConfigToProto(o *composer.EnvironmentConfigPrivateEnvironmentConfig) *composerpb.ComposerEnvironmentConfigPrivateEnvironmentConfig {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigPrivateEnvironmentConfig{
		EnablePrivateEnvironment:   dcl.ValueOrEmptyBool(o.EnablePrivateEnvironment),
		PrivateClusterConfig:       ComposerEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigToProto(o.PrivateClusterConfig),
		WebServerIpv4CidrBlock:     dcl.ValueOrEmptyString(o.WebServerIPv4CidrBlock),
		CloudSqlIpv4CidrBlock:      dcl.ValueOrEmptyString(o.CloudSqlIPv4CidrBlock),
		WebServerIpv4ReservedRange: dcl.ValueOrEmptyString(o.WebServerIPv4ReservedRange),
	}
	return p
}

// EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigToProto converts a EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig resource to its proto representation.
func ComposerEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfigToProto(o *composer.EnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig) *composerpb.ComposerEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigPrivateEnvironmentConfigPrivateClusterConfig{
		EnablePrivateEndpoint:   dcl.ValueOrEmptyBool(o.EnablePrivateEndpoint),
		MasterIpv4CidrBlock:     dcl.ValueOrEmptyString(o.MasterIPv4CidrBlock),
		MasterIpv4ReservedRange: dcl.ValueOrEmptyString(o.MasterIPv4ReservedRange),
	}
	return p
}

// EnvironmentConfigWebServerNetworkAccessControlToProto converts a EnvironmentConfigWebServerNetworkAccessControl resource to its proto representation.
func ComposerEnvironmentConfigWebServerNetworkAccessControlToProto(o *composer.EnvironmentConfigWebServerNetworkAccessControl) *composerpb.ComposerEnvironmentConfigWebServerNetworkAccessControl {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigWebServerNetworkAccessControl{}
	for _, r := range o.AllowedIPRanges {
		p.AllowedIpRanges = append(p.AllowedIpRanges, ComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesToProto(&r))
	}
	return p
}

// EnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesToProto converts a EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges resource to its proto representation.
func ComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIPRangesToProto(o *composer.EnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges) *composerpb.ComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigWebServerNetworkAccessControlAllowedIPRanges{
		Value:       dcl.ValueOrEmptyString(o.Value),
		Description: dcl.ValueOrEmptyString(o.Description),
	}
	return p
}

// EnvironmentConfigDatabaseConfigToProto converts a EnvironmentConfigDatabaseConfig resource to its proto representation.
func ComposerEnvironmentConfigDatabaseConfigToProto(o *composer.EnvironmentConfigDatabaseConfig) *composerpb.ComposerEnvironmentConfigDatabaseConfig {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigDatabaseConfig{
		MachineType: dcl.ValueOrEmptyString(o.MachineType),
	}
	return p
}

// EnvironmentConfigWebServerConfigToProto converts a EnvironmentConfigWebServerConfig resource to its proto representation.
func ComposerEnvironmentConfigWebServerConfigToProto(o *composer.EnvironmentConfigWebServerConfig) *composerpb.ComposerEnvironmentConfigWebServerConfig {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigWebServerConfig{
		MachineType: dcl.ValueOrEmptyString(o.MachineType),
	}
	return p
}

// EnvironmentConfigEncryptionConfigToProto converts a EnvironmentConfigEncryptionConfig resource to its proto representation.
func ComposerEnvironmentConfigEncryptionConfigToProto(o *composer.EnvironmentConfigEncryptionConfig) *composerpb.ComposerEnvironmentConfigEncryptionConfig {
	if o == nil {
		return nil
	}
	p := &composerpb.ComposerEnvironmentConfigEncryptionConfig{
		KmsKeyName: dcl.ValueOrEmptyString(o.KmsKeyName),
	}
	return p
}

// EnvironmentToProto converts a Environment resource to its proto representation.
func EnvironmentToProto(resource *composer.Environment) *composerpb.ComposerEnvironment {
	p := &composerpb.ComposerEnvironment{
		Name:       dcl.ValueOrEmptyString(resource.Name),
		Config:     ComposerEnvironmentConfigToProto(resource.Config),
		Uuid:       dcl.ValueOrEmptyString(resource.Uuid),
		State:      ComposerEnvironmentStateEnumToProto(resource.State),
		CreateTime: dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime: dcl.ValueOrEmptyString(resource.UpdateTime),
		Project:    dcl.ValueOrEmptyString(resource.Project),
		Location:   dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) applyEnvironment(ctx context.Context, c *composer.Client, request *composerpb.ApplyComposerEnvironmentRequest) (*composerpb.ComposerEnvironment, error) {
	p := ProtoToEnvironment(request.GetResource())
	res, err := c.ApplyEnvironment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvironmentToProto(res)
	return r, nil
}

// ApplyEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) ApplyComposerEnvironment(ctx context.Context, request *composerpb.ApplyComposerEnvironmentRequest) (*composerpb.ComposerEnvironment, error) {
	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyEnvironment(ctx, cl, request)
}

// DeleteEnvironment handles the gRPC request by passing it to the underlying Environment Delete() method.
func (s *EnvironmentServer) DeleteComposerEnvironment(ctx context.Context, request *composerpb.DeleteComposerEnvironmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvironment(ctx, ProtoToEnvironment(request.GetResource()))

}

// ListComposerEnvironment handles the gRPC request by passing it to the underlying EnvironmentList() method.
func (s *EnvironmentServer) ListComposerEnvironment(ctx context.Context, request *composerpb.ListComposerEnvironmentRequest) (*composerpb.ListComposerEnvironmentResponse, error) {
	cl, err := createConfigEnvironment(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvironment(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*composerpb.ComposerEnvironment
	for _, r := range resources.Items {
		rp := EnvironmentToProto(r)
		protos = append(protos, rp)
	}
	return &composerpb.ListComposerEnvironmentResponse{Items: protos}, nil
}

func createConfigEnvironment(ctx context.Context, service_account_file string) (*composer.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return composer.NewClient(conf), nil
}
