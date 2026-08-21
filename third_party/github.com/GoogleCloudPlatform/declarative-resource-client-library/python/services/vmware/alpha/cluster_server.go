// Copyright 2023 Google LLC. All Rights Reserved.
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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vmware/alpha/vmware_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmware/alpha"
)

// ClusterServer implements the gRPC interface for Cluster.
type ClusterServer struct{}

// ProtoToClusterStateEnum converts a ClusterStateEnum enum from its proto representation.
func ProtoToVmwareAlphaClusterStateEnum(e alphapb.VmwareAlphaClusterStateEnum) *alpha.ClusterStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VmwareAlphaClusterStateEnum_name[int32(e)]; ok {
		e := alpha.ClusterStateEnum(n[len("VmwareAlphaClusterStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToCluster converts a Cluster resource from its proto representation.
func ProtoToCluster(p *alphapb.VmwareAlphaCluster) *alpha.Cluster {
	obj := &alpha.Cluster{
		Name:         dcl.StringOrNil(p.GetName()),
		CreateTime:   dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:   dcl.StringOrNil(p.GetUpdateTime()),
		State:        ProtoToVmwareAlphaClusterStateEnum(p.GetState()),
		Management:   dcl.Bool(p.GetManagement()),
		Uid:          dcl.StringOrNil(p.GetUid()),
		Project:      dcl.StringOrNil(p.GetProject()),
		Location:     dcl.StringOrNil(p.GetLocation()),
		PrivateCloud: dcl.StringOrNil(p.GetPrivateCloud()),
	}
	return obj
}

// ClusterStateEnumToProto converts a ClusterStateEnum enum to its proto representation.
func VmwareAlphaClusterStateEnumToProto(e *alpha.ClusterStateEnum) alphapb.VmwareAlphaClusterStateEnum {
	if e == nil {
		return alphapb.VmwareAlphaClusterStateEnum(0)
	}
	if v, ok := alphapb.VmwareAlphaClusterStateEnum_value["ClusterStateEnum"+string(*e)]; ok {
		return alphapb.VmwareAlphaClusterStateEnum(v)
	}
	return alphapb.VmwareAlphaClusterStateEnum(0)
}

// ClusterToProto converts a Cluster resource to its proto representation.
func ClusterToProto(resource *alpha.Cluster) *alphapb.VmwareAlphaCluster {
	p := &alphapb.VmwareAlphaCluster{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetState(VmwareAlphaClusterStateEnumToProto(resource.State))
	p.SetManagement(dcl.ValueOrEmptyBool(resource.Management))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetPrivateCloud(dcl.ValueOrEmptyString(resource.PrivateCloud))

	return p
}

// applyCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) applyCluster(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVmwareAlphaClusterRequest) (*alphapb.VmwareAlphaCluster, error) {
	p := ProtoToCluster(request.GetResource())
	res, err := c.ApplyCluster(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClusterToProto(res)
	return r, nil
}

// applyVmwareAlphaCluster handles the gRPC request by passing it to the underlying Cluster Apply() method.
func (s *ClusterServer) ApplyVmwareAlphaCluster(ctx context.Context, request *alphapb.ApplyVmwareAlphaClusterRequest) (*alphapb.VmwareAlphaCluster, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyCluster(ctx, cl, request)
}

// DeleteCluster handles the gRPC request by passing it to the underlying Cluster Delete() method.
func (s *ClusterServer) DeleteVmwareAlphaCluster(ctx context.Context, request *alphapb.DeleteVmwareAlphaClusterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteCluster(ctx, ProtoToCluster(request.GetResource()))

}

// ListVmwareAlphaCluster handles the gRPC request by passing it to the underlying ClusterList() method.
func (s *ClusterServer) ListVmwareAlphaCluster(ctx context.Context, request *alphapb.ListVmwareAlphaClusterRequest) (*alphapb.ListVmwareAlphaClusterResponse, error) {
	cl, err := createConfigCluster(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListCluster(ctx, request.GetProject(), request.GetLocation(), request.GetPrivateCloud())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VmwareAlphaCluster
	for _, r := range resources.Items {
		rp := ClusterToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVmwareAlphaClusterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigCluster(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
