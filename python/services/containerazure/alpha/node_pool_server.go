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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/containerazure/alpha/containerazure_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure/alpha"
)

// NodePoolServer implements the gRPC interface for NodePool.
type NodePoolServer struct{}

// ProtoToNodePoolStateEnum converts a NodePoolStateEnum enum from its proto representation.
func ProtoToContainerazureAlphaNodePoolStateEnum(e alphapb.ContainerazureAlphaNodePoolStateEnum) *alpha.NodePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ContainerazureAlphaNodePoolStateEnum_name[int32(e)]; ok {
		e := alpha.NodePoolStateEnum(n[len("ContainerazureAlphaNodePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToNodePoolConfig converts a NodePoolConfig object from its proto representation.
func ProtoToContainerazureAlphaNodePoolConfig(p *alphapb.ContainerazureAlphaNodePoolConfig) *alpha.NodePoolConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfig{
		VmSize:      dcl.StringOrNil(p.GetVmSize()),
		RootVolume:  ProtoToContainerazureAlphaNodePoolConfigRootVolume(p.GetRootVolume()),
		SshConfig:   ProtoToContainerazureAlphaNodePoolConfigSshConfig(p.GetSshConfig()),
		ImageType:   dcl.StringOrNil(p.GetImageType()),
		ProxyConfig: ProtoToContainerazureAlphaNodePoolConfigProxyConfig(p.GetProxyConfig()),
	}
	return obj
}

// ProtoToNodePoolConfigRootVolume converts a NodePoolConfigRootVolume object from its proto representation.
func ProtoToContainerazureAlphaNodePoolConfigRootVolume(p *alphapb.ContainerazureAlphaNodePoolConfigRootVolume) *alpha.NodePoolConfigRootVolume {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigRootVolume{
		SizeGib: dcl.Int64OrNil(p.GetSizeGib()),
	}
	return obj
}

// ProtoToNodePoolConfigSshConfig converts a NodePoolConfigSshConfig object from its proto representation.
func ProtoToContainerazureAlphaNodePoolConfigSshConfig(p *alphapb.ContainerazureAlphaNodePoolConfigSshConfig) *alpha.NodePoolConfigSshConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigSshConfig{
		AuthorizedKey: dcl.StringOrNil(p.GetAuthorizedKey()),
	}
	return obj
}

// ProtoToNodePoolConfigProxyConfig converts a NodePoolConfigProxyConfig object from its proto representation.
func ProtoToContainerazureAlphaNodePoolConfigProxyConfig(p *alphapb.ContainerazureAlphaNodePoolConfigProxyConfig) *alpha.NodePoolConfigProxyConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolConfigProxyConfig{
		ResourceGroupId: dcl.StringOrNil(p.GetResourceGroupId()),
		SecretId:        dcl.StringOrNil(p.GetSecretId()),
	}
	return obj
}

// ProtoToNodePoolAutoscaling converts a NodePoolAutoscaling object from its proto representation.
func ProtoToContainerazureAlphaNodePoolAutoscaling(p *alphapb.ContainerazureAlphaNodePoolAutoscaling) *alpha.NodePoolAutoscaling {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolAutoscaling{
		MinNodeCount: dcl.Int64OrNil(p.GetMinNodeCount()),
		MaxNodeCount: dcl.Int64OrNil(p.GetMaxNodeCount()),
	}
	return obj
}

// ProtoToNodePoolMaxPodsConstraint converts a NodePoolMaxPodsConstraint object from its proto representation.
func ProtoToContainerazureAlphaNodePoolMaxPodsConstraint(p *alphapb.ContainerazureAlphaNodePoolMaxPodsConstraint) *alpha.NodePoolMaxPodsConstraint {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64OrNil(p.GetMaxPodsPerNode()),
	}
	return obj
}

// ProtoToNodePoolManagement converts a NodePoolManagement object from its proto representation.
func ProtoToContainerazureAlphaNodePoolManagement(p *alphapb.ContainerazureAlphaNodePoolManagement) *alpha.NodePoolManagement {
	if p == nil {
		return nil
	}
	obj := &alpha.NodePoolManagement{
		AutoRepair: dcl.Bool(p.GetAutoRepair()),
	}
	return obj
}

// ProtoToNodePool converts a NodePool resource from its proto representation.
func ProtoToNodePool(p *alphapb.ContainerazureAlphaNodePool) *alpha.NodePool {
	obj := &alpha.NodePool{
		Name:                  dcl.StringOrNil(p.GetName()),
		Version:               dcl.StringOrNil(p.GetVersion()),
		Config:                ProtoToContainerazureAlphaNodePoolConfig(p.GetConfig()),
		SubnetId:              dcl.StringOrNil(p.GetSubnetId()),
		Autoscaling:           ProtoToContainerazureAlphaNodePoolAutoscaling(p.GetAutoscaling()),
		State:                 ProtoToContainerazureAlphaNodePoolStateEnum(p.GetState()),
		Uid:                   dcl.StringOrNil(p.GetUid()),
		Reconciling:           dcl.Bool(p.GetReconciling()),
		CreateTime:            dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:            dcl.StringOrNil(p.GetUpdateTime()),
		Etag:                  dcl.StringOrNil(p.GetEtag()),
		MaxPodsConstraint:     ProtoToContainerazureAlphaNodePoolMaxPodsConstraint(p.GetMaxPodsConstraint()),
		Management:            ProtoToContainerazureAlphaNodePoolManagement(p.GetManagement()),
		AzureAvailabilityZone: dcl.StringOrNil(p.GetAzureAvailabilityZone()),
		Project:               dcl.StringOrNil(p.GetProject()),
		Location:              dcl.StringOrNil(p.GetLocation()),
		Cluster:               dcl.StringOrNil(p.GetCluster()),
	}
	return obj
}

// NodePoolStateEnumToProto converts a NodePoolStateEnum enum to its proto representation.
func ContainerazureAlphaNodePoolStateEnumToProto(e *alpha.NodePoolStateEnum) alphapb.ContainerazureAlphaNodePoolStateEnum {
	if e == nil {
		return alphapb.ContainerazureAlphaNodePoolStateEnum(0)
	}
	if v, ok := alphapb.ContainerazureAlphaNodePoolStateEnum_value["NodePoolStateEnum"+string(*e)]; ok {
		return alphapb.ContainerazureAlphaNodePoolStateEnum(v)
	}
	return alphapb.ContainerazureAlphaNodePoolStateEnum(0)
}

// NodePoolConfigToProto converts a NodePoolConfig object to its proto representation.
func ContainerazureAlphaNodePoolConfigToProto(o *alpha.NodePoolConfig) *alphapb.ContainerazureAlphaNodePoolConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolConfig{}
	p.SetVmSize(dcl.ValueOrEmptyString(o.VmSize))
	p.SetRootVolume(ContainerazureAlphaNodePoolConfigRootVolumeToProto(o.RootVolume))
	p.SetSshConfig(ContainerazureAlphaNodePoolConfigSshConfigToProto(o.SshConfig))
	p.SetImageType(dcl.ValueOrEmptyString(o.ImageType))
	p.SetProxyConfig(ContainerazureAlphaNodePoolConfigProxyConfigToProto(o.ProxyConfig))
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
func ContainerazureAlphaNodePoolConfigRootVolumeToProto(o *alpha.NodePoolConfigRootVolume) *alphapb.ContainerazureAlphaNodePoolConfigRootVolume {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolConfigRootVolume{}
	p.SetSizeGib(dcl.ValueOrEmptyInt64(o.SizeGib))
	return p
}

// NodePoolConfigSshConfigToProto converts a NodePoolConfigSshConfig object to its proto representation.
func ContainerazureAlphaNodePoolConfigSshConfigToProto(o *alpha.NodePoolConfigSshConfig) *alphapb.ContainerazureAlphaNodePoolConfigSshConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolConfigSshConfig{}
	p.SetAuthorizedKey(dcl.ValueOrEmptyString(o.AuthorizedKey))
	return p
}

// NodePoolConfigProxyConfigToProto converts a NodePoolConfigProxyConfig object to its proto representation.
func ContainerazureAlphaNodePoolConfigProxyConfigToProto(o *alpha.NodePoolConfigProxyConfig) *alphapb.ContainerazureAlphaNodePoolConfigProxyConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolConfigProxyConfig{}
	p.SetResourceGroupId(dcl.ValueOrEmptyString(o.ResourceGroupId))
	p.SetSecretId(dcl.ValueOrEmptyString(o.SecretId))
	return p
}

// NodePoolAutoscalingToProto converts a NodePoolAutoscaling object to its proto representation.
func ContainerazureAlphaNodePoolAutoscalingToProto(o *alpha.NodePoolAutoscaling) *alphapb.ContainerazureAlphaNodePoolAutoscaling {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolAutoscaling{}
	p.SetMinNodeCount(dcl.ValueOrEmptyInt64(o.MinNodeCount))
	p.SetMaxNodeCount(dcl.ValueOrEmptyInt64(o.MaxNodeCount))
	return p
}

// NodePoolMaxPodsConstraintToProto converts a NodePoolMaxPodsConstraint object to its proto representation.
func ContainerazureAlphaNodePoolMaxPodsConstraintToProto(o *alpha.NodePoolMaxPodsConstraint) *alphapb.ContainerazureAlphaNodePoolMaxPodsConstraint {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolMaxPodsConstraint{}
	p.SetMaxPodsPerNode(dcl.ValueOrEmptyInt64(o.MaxPodsPerNode))
	return p
}

// NodePoolManagementToProto converts a NodePoolManagement object to its proto representation.
func ContainerazureAlphaNodePoolManagementToProto(o *alpha.NodePoolManagement) *alphapb.ContainerazureAlphaNodePoolManagement {
	if o == nil {
		return nil
	}
	p := &alphapb.ContainerazureAlphaNodePoolManagement{}
	p.SetAutoRepair(dcl.ValueOrEmptyBool(o.AutoRepair))
	return p
}

// NodePoolToProto converts a NodePool resource to its proto representation.
func NodePoolToProto(resource *alpha.NodePool) *alphapb.ContainerazureAlphaNodePool {
	p := &alphapb.ContainerazureAlphaNodePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetVersion(dcl.ValueOrEmptyString(resource.Version))
	p.SetConfig(ContainerazureAlphaNodePoolConfigToProto(resource.Config))
	p.SetSubnetId(dcl.ValueOrEmptyString(resource.SubnetId))
	p.SetAutoscaling(ContainerazureAlphaNodePoolAutoscalingToProto(resource.Autoscaling))
	p.SetState(ContainerazureAlphaNodePoolStateEnumToProto(resource.State))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetReconciling(dcl.ValueOrEmptyBool(resource.Reconciling))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetMaxPodsConstraint(ContainerazureAlphaNodePoolMaxPodsConstraintToProto(resource.MaxPodsConstraint))
	p.SetManagement(ContainerazureAlphaNodePoolManagementToProto(resource.Management))
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
func (s *NodePoolServer) applyNodePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyContainerazureAlphaNodePoolRequest) (*alphapb.ContainerazureAlphaNodePool, error) {
	p := ProtoToNodePool(request.GetResource())
	res, err := c.ApplyNodePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := NodePoolToProto(res)
	return r, nil
}

// applyContainerazureAlphaNodePool handles the gRPC request by passing it to the underlying NodePool Apply() method.
func (s *NodePoolServer) ApplyContainerazureAlphaNodePool(ctx context.Context, request *alphapb.ApplyContainerazureAlphaNodePoolRequest) (*alphapb.ContainerazureAlphaNodePool, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyNodePool(ctx, cl, request)
}

// DeleteNodePool handles the gRPC request by passing it to the underlying NodePool Delete() method.
func (s *NodePoolServer) DeleteContainerazureAlphaNodePool(ctx context.Context, request *alphapb.DeleteContainerazureAlphaNodePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteNodePool(ctx, ProtoToNodePool(request.GetResource()))

}

// ListContainerazureAlphaNodePool handles the gRPC request by passing it to the underlying NodePoolList() method.
func (s *NodePoolServer) ListContainerazureAlphaNodePool(ctx context.Context, request *alphapb.ListContainerazureAlphaNodePoolRequest) (*alphapb.ListContainerazureAlphaNodePoolResponse, error) {
	cl, err := createConfigNodePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListNodePool(ctx, request.GetProject(), request.GetLocation(), request.GetCluster())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ContainerazureAlphaNodePool
	for _, r := range resources.Items {
		rp := NodePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListContainerazureAlphaNodePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigNodePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
