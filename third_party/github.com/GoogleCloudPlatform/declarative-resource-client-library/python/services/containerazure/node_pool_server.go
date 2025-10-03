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
	containerazurepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/containerazure_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure"
)

// NodePoolServer implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolStateEnum converts a NodePoolStateEnum enum from its proto representation.
func ProtoToContainerazureNodePoolStateEnum(e containerazurepb.ContainerazureNodePoolStateEnum) *containerazure.NodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := containerazurepb.ContainerazureNodePoolStateEnum_name[int32(e)]; ok {
		e := containerazure.NodePoolStateEnum(n[len("ContainerazureNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig object from its proto representation.
func ProtoToContainerazureNodePoolConfig(p *containerazurepb.ContainerazureNodePoolConfig) *containerazure.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &containerazure.NodePoolConfig{
		VmSize:      dcl.StringOrNil(p.GetVmSize()),
		RootVolume:  ProtoToContainerazureNodePoolConfigRootVolume(p.GetRootVolume()),
		SshConfig:   ProtoToContainerazureNodePoolConfigSshConfig(p.GetSshConfig()),
		ProxyConfig: ProtoToContainerazureNodePoolConfigProxyConfig(p.GetProxyConfig()),
	}
	return obj
}

// ProtoToNodePoolConfigRootVolume converts a NodePoolConfigRootVolume object from its proto representation.
func ProtoToContainerazureNodePoolConfigRootVolume(p *containerazurepb.ContainerazureNodePoolConfigRootVolume) *containerazure.NodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &containerazure.NodePoolConfigRootVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToNodePoolConfigSshConfig converts a NodePoolConfigSshConfig object from its proto representation.
func ProtoToContainerazureNodePoolConfigSshConfig(p *containerazurepb.ContainerazureNodePoolConfigSshConfig) *containerazure.NodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &containerazure.NodePoolConfigSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.GetAuthorizedKey()),
	}
	return obj
}

// ProtoToNodePoolConfigProxyConfig converts a NodePoolConfigProxyConfig object from its proto representation.
func ProtoToContainerazureNodePoolConfigProxyConfig(p *containerazurepb.ContainerazureNodePoolConfigProxyConfig) *containerazure.NodePoolConfigProxyConfig {
	if p == nil {
		return nil
	}
	obj := &containerazure.NodePoolConfigProxyConfig{
		ResourceGroupId: dcl.StringOrNil(p.GetResourceGroupId()),
		SecretId:        dcl.StringOrNil(p.GetSecretId()),
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling object from its proto representation.
func ProtoToContainerazureNodePoolAutoscaling(p *containerazurepb.ContainerazureNodePoolAutoscaling) *containerazure.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &containerazure.NodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.GetMinNodeCount()),
		MaxNodeCount: dcl.Int64OrNil(p.GetMaxNodeCount()),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint object from its proto representation.
func ProtoToContainerazureNodePoolMaxPodsConstraint(p *containerazurepb.ContainerazureNodePoolMaxPodsConstraint) *containerazure.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &containerazure.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.GetMaxPodsPerNode()),
	}
	return obj
}

// ProtoToNodePoolManagement converts a NodePoolManagement object from its proto representation.
func ProtoToContainerazureNodePoolManagement(p *containerazurepb.ContainerazureNodePoolManagement) *containerazure.NodePoolManagement {
	if p == nil {
		return nil
	}
	obj := &containerazure.NodePoolManagement{
		AutoRepair: dcl.Bool(p.GetAutoRepair()),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *containerazurepb.ContainerazureNodePool) *containerazure.NodePool {
	obj := &containerazure.NodePool{
		Name:                  dcl.StringOrNil(p.GetName()),
		Version:               dcl.StringOrNil(p.GetVersion()),
		Config:                ProtoToContainerazureNodePoolConfig(p.GetConfig()),
		SubnetId:              dcl.StringOrNil(p.GetSubnetId()),
		Autoscaling:           ProtoToContainerazureNodePoolAutoscaling(p.GetAutoscaling()),
		State:                 ProtoToContainerazureNodePoolStateEnum(p.GetState()),
		Uid:                   dcl.StringOrNil(p.GetUid()),
		Reconciling:           dcl.Bool(p.GetReconciling()),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                  dcl.StringOrNil(p.GetEtag()),
		MaxPodsConstraint:     ProtoToContainerazureNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Management:            ProtoToContainerazureNodePoolManagement(p.GetManagement()),
		AzureAvailabilityZone: dcl.StringOrNil(p.GetAzureAvailabilityZone()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
		Cluster:               dcl.StringOrNil(p.GetCluster()),
	}
	return obj
}

// NodePoolStateEnumToProto converts a NodePoolStateEnum enum to its proto representation.
func ContainerazureNodePoolStateEnumToProto(e *containerazure.NodePoolStateEnum) containerazurepb.ContainerazureNodePoolStateEnum {
	if e == nil {
		return containerazurepb.ContainerazureNodePoolStateEnum(0)
	}
	if v, ok := containerazurepb.ContainerazureNodePoolStateEnum_value["NodePoolStateEnum"+string(*e)]; ok {
		return containerazurepb.ContainerazureNodePoolStateEnum(v)
	}
	return containerazurepb.ContainerazureNodePoolStateEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig object to its proto representation.
func ContainerazureNodePoolConfigToProto(o *containerazure.NodePoolConfig) *containerazurepb.ContainerazureNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureNodePoolConfig{}
	p.SetVmSize(dcl.ValueOrEmptyString(o.VmSize))
	p.SetRootVolume(ContainerazureNodePoolConfigRootVolumeToProto(o.RootVolume))
	p.SetSshConfig(ContainerazureNodePoolConfigSshConfigToProto(o.SshConfig))
	p.SetProxyConfig(ContainerazureNodePoolConfigProxyConfigToProto(o.ProxyConfig))
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
func ContainerazureNodePoolConfigRootVolumeToProto(o *containerazure.NodePoolConfigRootVolume) *containerazurepb.ContainerazureNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureNodePoolConfigRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// NodePoolConfigSshConfigToProto converts a NodePoolConfigSshConfig object to its proto representation.
func ContainerazureNodePoolConfigSshConfigToProto(o *containerazure.NodePoolConfigSshConfig) *containerazurepb.ContainerazureNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureNodePoolConfigSshConfig{}
	p.SetAuthorizedKey(dcl.ValueOrEmptyString(o.AuthorizedKey))
	return p
}

// NodePoolConfigProxyConfigToProto converts a NodePoolConfigProxyConfig object to its proto representation.
func ContainerazureNodePoolConfigProxyConfigToProto(o *containerazure.NodePoolConfigProxyConfig) *containerazurepb.ContainerazureNodePoolConfigProxyConfig {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureNodePoolConfigProxyConfig{}
	p.SetResourceGroupId(dcl.ValueOrEmptyString(o.ResourceGroupId))
	p.SetSecretId(dcl.ValueOrEmptyString(o.SecretId))
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling object to its proto representation.
func ContainerazureNodePoolAutoscalingToProto(o *containerazure.NodePoolAutoscaling) *containerazurepb.ContainerazureNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureNodePoolAutoscaling{}
	p.SetMinNodeCount(dcl.ValueOrEmptyInt64(o.MinNodeCount))
	p.SetMaxNodeCount(dcl.ValueOrEmptyInt64(o.MaxNodeCount))
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint object to its proto representation.
func ContainerazureNodePoolMaxPodsConstraintToProto(o *containerazure.NodePoolMaxPodsConstraint) *containerazurepb.ContainerazureNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureNodePoolMaxPodsConstraint{}
	p.SetMaxPodsPerNode(dcl.ValueOrEmptyInt64(o.MaxPodsPerNode))
	return p
}

// NodePoolManagementToProto converts a NodePoolManagement object to its proto representation.
func ContainerazureNodePoolManagementToProto(o *containerazure.NodePoolManagement) *containerazurepb.ContainerazureNodePoolManagement {
	if o == nil {
		return nil
	}
	p := &containerazurepb.ContainerazureNodePoolManagement{}
	p.SetAutoRepair(dcl.ValueOrEmptyBool(o.AutoRepair))
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *containerazure.NodePool) *containerazurepb.ContainerazureNodePool {
	p := &containerazurepb.ContainerazureNodePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyString(resource.Version))
	p.SetConfig(ContainerazureNodePoolConfigToProto(resource.Config))
	p.SetSubnetId(dcl.ValueOrEmptyString(resource.SubnetId))
	p.SetAutoscaling(ContainerazureNodePoolAutoscalingToProto(resource.Autoscaling))
	p.SetState(ContainerazureNodePoolStateEnumToProto(resource.State))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetMaxPodsConstraint(ContainerazureNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint))
	p.SetManagement(ContainerazureNodePoolManagementToProto(resource.Management))
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
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *containerazure.Client, request *containerazurepb.ApplyContainerazureNodePoolRequest) (*containerazurepb.ContainerazureNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// applyContainerazureNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerazureNodePool(ctx context.Context, request *containerazurepb.ApplyContainerazureNodePoolRequest) (*containerazurepb.ContainerazureNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerazureNodePool(ctx context.Context, request *containerazurepb.DeleteContainerazureNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerazureNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerazureNodePool(ctx context.Context, request *containerazurepb.ListContainerazureNodePoolRequest) (*containerazurepb.ListContainerazureNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.GetProject(), request.GetLocation(), request.GetCluster())
	if err != nil {
		return nil, err
	}
	var protos []*containerazurepb.ContainerazureNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &containerazurepb.ListContainerazureNodePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*containerazure.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return containerazure.NewClient(conf), nil
}
