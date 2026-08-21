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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/beta/containerazure_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/beta"
)

// NodePoolServer implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolStateEnum converts a NodePoolStateEnum enum from its proto representation.
func ProtoToContainerazureBetaNodePoolStateEnum(e betapb.ContainerazureBetaNodePoolStateEnum) *beta.NodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ContainerazureBetaNodePoolStateEnum_name[int32(e)]; ok {
		e := beta.NodePoolStateEnum(n[len("ContainerazureBetaNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig object from its proto representation.
func ProtoToContainerazureBetaNodePoolConfig(p *betapb.ContainerazureBetaNodePoolConfig) *beta.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfig{
		VmSize:      dcl.StringOrNil(p.GetVmSize()),
		RootVolume:  ProtoToContainerazureBetaNodePoolConfigRootVolume(p.GetRootVolume()),
		SshConfig:   ProtoToContainerazureBetaNodePoolConfigSshConfig(p.GetSshConfig()),
		ImageType:   dcl.StringOrNil(p.GetImageType()),
		ProxyConfig: ProtoToContainerazureBetaNodePoolConfigProxyConfig(p.GetProxyConfig()),
	}
	return obj
}

// ProtoToNodePoolConfigRootVolume converts a NodePoolConfigRootVolume object from its proto representation.
func ProtoToContainerazureBetaNodePoolConfigRootVolume(p *betapb.ContainerazureBetaNodePoolConfigRootVolume) *beta.NodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigRootVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToNodePoolConfigSshConfig converts a NodePoolConfigSshConfig object from its proto representation.
func ProtoToContainerazureBetaNodePoolConfigSshConfig(p *betapb.ContainerazureBetaNodePoolConfigSshConfig) *beta.NodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.GetAuthorizedKey()),
	}
	return obj
}

// ProtoToNodePoolConfigProxyConfig converts a NodePoolConfigProxyConfig object from its proto representation.
func ProtoToContainerazureBetaNodePoolConfigProxyConfig(p *betapb.ContainerazureBetaNodePoolConfigProxyConfig) *beta.NodePoolConfigProxyConfig {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolConfigProxyConfig{
		ResourceGroupId: dcl.StringOrNil(p.GetResourceGroupId()),
		SecretId:        dcl.StringOrNil(p.GetSecretId()),
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling object from its proto representation.
func ProtoToContainerazureBetaNodePoolAutoscaling(p *betapb.ContainerazureBetaNodePoolAutoscaling) *beta.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.GetMinNodeCount()),
		MaxNodeCount: dcl.Int64OrNil(p.GetMaxNodeCount()),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint object from its proto representation.
func ProtoToContainerazureBetaNodePoolMaxPodsConstraint(p *betapb.ContainerazureBetaNodePoolMaxPodsConstraint) *beta.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.GetMaxPodsPerNode()),
	}
	return obj
}

// ProtoToNodePoolManagement converts a NodePoolManagement object from its proto representation.
func ProtoToContainerazureBetaNodePoolManagement(p *betapb.ContainerazureBetaNodePoolManagement) *beta.NodePoolManagement {
	if p == nil {
		return nil
	}
	obj := &beta.NodePoolManagement{
		AutoRepair: dcl.Bool(p.GetAutoRepair()),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *betapb.ContainerazureBetaNodePool) *beta.NodePool {
	obj := &beta.NodePool{
		Name:                  dcl.StringOrNil(p.GetName()),
		Version:               dcl.StringOrNil(p.GetVersion()),
		Config:                ProtoToContainerazureBetaNodePoolConfig(p.GetConfig()),
		SubnetId:              dcl.StringOrNil(p.GetSubnetId()),
		Autoscaling:           ProtoToContainerazureBetaNodePoolAutoscaling(p.GetAutoscaling()),
		State:                 ProtoToContainerazureBetaNodePoolStateEnum(p.GetState()),
		Uid:                   dcl.StringOrNil(p.GetUid()),
		Reconciling:           dcl.Bool(p.GetReconciling()),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                  dcl.StringOrNil(p.GetEtag()),
		MaxPodsConstraint:     ProtoToContainerazureBetaNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Management:            ProtoToContainerazureBetaNodePoolManagement(p.GetManagement()),
		AzureAvailabilityZone: dcl.StringOrNil(p.GetAzureAvailabilityZone()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
		Cluster:               dcl.StringOrNil(p.GetCluster()),
	}
	return obj
}

// NodePoolStateEnumToProto converts a NodePoolStateEnum enum to its proto representation.
func ContainerazureBetaNodePoolStateEnumToProto(e *beta.NodePoolStateEnum) betapb.ContainerazureBetaNodePoolStateEnum {
	if e == nil {
		return betapb.ContainerazureBetaNodePoolStateEnum(0)
	}
	if v, ok := betapb.ContainerazureBetaNodePoolStateEnum_value["NodePoolStateEnum"+string(*e)]; ok {
		return betapb.ContainerazureBetaNodePoolStateEnum(v)
	}
	return betapb.ContainerazureBetaNodePoolStateEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig object to its proto representation.
func ContainerazureBetaNodePoolConfigToProto(o *beta.NodePoolConfig) *betapb.ContainerazureBetaNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaNodePoolConfig{}
	p.SetVmSize(dcl.ValueOrEmptyString(o.VmSize))
	p.SetRootVolume(ContainerazureBetaNodePoolConfigRootVolumeToProto(o.RootVolume))
	p.SetSshConfig(ContainerazureBetaNodePoolConfigSshConfigToProto(o.SshConfig))
	p.SetImageType(dcl.ValueOrEmptyString(o.ImageType))
	p.SetProxyConfig(ContainerazureBetaNodePoolConfigProxyConfigToProto(o.ProxyConfig))
	mTags := make(map[string]string, len(o.Tags))
	for k, r := range o.Tags {
		mTags[k] = r
	}
	p.SetTags(mTags)
	mLabels := make(map[string]string, len(o.Labels))
	for k, r := range o.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	return p
}

// NodePoolConfigRootVolumeToProto converts a NodePoolConfigRootVolume object to its proto representation.
func ContainerazureBetaNodePoolConfigRootVolumeToProto(o *beta.NodePoolConfigRootVolume) *betapb.ContainerazureBetaNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaNodePoolConfigRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// NodePoolConfigSshConfigToProto converts a NodePoolConfigSshConfig object to its proto representation.
func ContainerazureBetaNodePoolConfigSshConfigToProto(o *beta.NodePoolConfigSshConfig) *betapb.ContainerazureBetaNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaNodePoolConfigSshConfig{}
	p.SetAuthorizedKey(dcl.ValueOrEmptyString(o.AuthorizedKey))
	return p
}

// NodePoolConfigProxyConfigToProto converts a NodePoolConfigProxyConfig object to its proto representation.
func ContainerazureBetaNodePoolConfigProxyConfigToProto(o *beta.NodePoolConfigProxyConfig) *betapb.ContainerazureBetaNodePoolConfigProxyConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaNodePoolConfigProxyConfig{}
	p.SetResourceGroupId(dcl.ValueOrEmptyString(o.ResourceGroupId))
	p.SetSecretId(dcl.ValueOrEmptyString(o.SecretId))
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling object to its proto representation.
func ContainerazureBetaNodePoolAutoscalingToProto(o *beta.NodePoolAutoscaling) *betapb.ContainerazureBetaNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaNodePoolAutoscaling{}
	p.SetMinNodeCount(dcl.ValueOrEmptyInt64(o.MinNodeCount))
	p.SetMaxNodeCount(dcl.ValueOrEmptyInt64(o.MaxNodeCount))
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint object to its proto representation.
func ContainerazureBetaNodePoolMaxPodsConstraintToProto(o *beta.NodePoolMaxPodsConstraint) *betapb.ContainerazureBetaNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaNodePoolMaxPodsConstraint{}
	p.SetMaxPodsPerNode(dcl.ValueOrEmptyInt64(o.MaxPodsPerNode))
	return p
}

// NodePoolManagementToProto converts a NodePoolManagement object to its proto representation.
func ContainerazureBetaNodePoolManagementToProto(o *beta.NodePoolManagement) *betapb.ContainerazureBetaNodePoolManagement {
	if o == nil {
		return nil
	}
	p := &betapb.ContainerazureBetaNodePoolManagement{}
	p.SetAutoRepair(dcl.ValueOrEmptyBool(o.AutoRepair))
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *beta.NodePool) *betapb.ContainerazureBetaNodePool {
	p := &betapb.ContainerazureBetaNodePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyString(resource.Version))
	p.SetConfig(ContainerazureBetaNodePoolConfigToProto(resource.Config))
	p.SetSubnetId(dcl.ValueOrEmptyString(resource.SubnetId))
	p.SetAutoscaling(ContainerazureBetaNodePoolAutoscalingToProto(resource.Autoscaling))
	p.SetState(ContainerazureBetaNodePoolStateEnumToProto(resource.State))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetMaxPodsConstraint(ContainerazureBetaNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint))
	p.SetManagement(ContainerazureBetaNodePoolManagementToProto(resource.Management))
	p.SetAzureAvailabilityZone(dcl.ValueOrEmptyString(resource.AzureAvailabilityZone))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetCluster(dcl.ValueOrEmptyString(resource.Cluster))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *beta.Client, request *betapb.ApplyContainerazureBetaNodePoolRequest) (*betapb.ContainerazureBetaNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// applyContainerazureBetaNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerazureBetaNodePool(ctx context.Context, request *betapb.ApplyContainerazureBetaNodePoolRequest) (*betapb.ContainerazureBetaNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerazureBetaNodePool(ctx context.Context, request *betapb.DeleteContainerazureBetaNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerazureBetaNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerazureBetaNodePool(ctx context.Context, request *betapb.ListContainerazureBetaNodePoolRequest) (*betapb.ListContainerazureBetaNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.GetProject(), request.GetLocation(), request.GetCluster())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ContainerazureBetaNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListContainerazureBetaNodePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
