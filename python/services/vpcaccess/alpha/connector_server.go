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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vpcaccess/alpha/vpcaccess_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess/alpha"
)

// ConnectorServer implements the gRPC interface for Connector.
type ConnectorServer struct{}

// ProtoToConnectorStateEnum converts a ConnectorStateEnum enum from its proto representation.
func ProtoToVpcaccessAlphaConnectorStateEnum(e alphapb.VpcaccessAlphaConnectorStateEnum) *alpha.ConnectorStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VpcaccessAlphaConnectorStateEnum_name[int32(e)]; ok {
		e := alpha.ConnectorStateEnum(n[len("VpcaccessAlphaConnectorStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToConnectorSubnet converts a ConnectorSubnet object from its proto representation.
func ProtoToVpcaccessAlphaConnectorSubnet(p *alphapb.VpcaccessAlphaConnectorSubnet) *alpha.ConnectorSubnet {
	if p == nil {
		return nil
	}
	obj := &alpha.ConnectorSubnet{
		Name:      dcl.StringOrNil(p.GetName()),
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
	}
	return obj
}

// ProtoToConnector converts a Connector resource from its proto representation.
func ProtoToConnector(p *alphapb.VpcaccessAlphaConnector) *alpha.Connector {
	obj := &alpha.Connector{
		Name:          dcl.StringOrNil(p.GetName()),
		Network:       dcl.StringOrNil(p.GetNetwork()),
		IPCidrRange:   dcl.StringOrNil(p.GetIpCidrRange()),
		State:         ProtoToVpcaccessAlphaConnectorStateEnum(p.GetState()),
		MinThroughput: dcl.Int64OrNil(p.GetMinThroughput()),
		MaxThroughput: dcl.Int64OrNil(p.GetMaxThroughput()),
		Subnet:        ProtoToVpcaccessAlphaConnectorSubnet(p.GetSubnet()),
		MachineType:   dcl.StringOrNil(p.GetMachineType()),
		MinInstances:  dcl.Int64OrNil(p.GetMinInstances()),
		MaxInstances:  dcl.Int64OrNil(p.GetMaxInstances()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetConnectedProjects() {
		obj.ConnectedProjects = append(obj.ConnectedProjects, r)
	}
	return obj
}

// ConnectorStateEnumToProto converts a ConnectorStateEnum enum to its proto representation.
func VpcaccessAlphaConnectorStateEnumToProto(e *alpha.ConnectorStateEnum) alphapb.VpcaccessAlphaConnectorStateEnum {
	if e == nil {
		return alphapb.VpcaccessAlphaConnectorStateEnum(0)
	}
	if v, ok := alphapb.VpcaccessAlphaConnectorStateEnum_value["ConnectorStateEnum"+string(*e)]; ok {
		return alphapb.VpcaccessAlphaConnectorStateEnum(v)
	}
	return alphapb.VpcaccessAlphaConnectorStateEnum(0)
}

// ConnectorSubnetToProto converts a ConnectorSubnet object to its proto representation.
func VpcaccessAlphaConnectorSubnetToProto(o *alpha.ConnectorSubnet) *alphapb.VpcaccessAlphaConnectorSubnet {
	if o == nil {
		return nil
	}
	p := &alphapb.VpcaccessAlphaConnectorSubnet{}
	p.SetName(dcl.ValueOrEmptyString(o.Name))
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	return p
}

// ConnectorToProto converts a Connector resource to its proto representation.
func ConnectorToProto(resource *alpha.Connector) *alphapb.VpcaccessAlphaConnector {
	p := &alphapb.VpcaccessAlphaConnector{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNetwork(dcl.ValueOrEmptyString(resource.Network))
	p.SetIpCidrRange(dcl.ValueOrEmptyString(resource.IPCidrRange))
	p.SetState(VpcaccessAlphaConnectorStateEnumToProto(resource.State))
	p.SetMinThroughput(dcl.ValueOrEmptyInt64(resource.MinThroughput))
	p.SetMaxThroughput(dcl.ValueOrEmptyInt64(resource.MaxThroughput))
	p.SetSubnet(VpcaccessAlphaConnectorSubnetToProto(resource.Subnet))
	p.SetMachineType(dcl.ValueOrEmptyString(resource.MachineType))
	p.SetMinInstances(dcl.ValueOrEmptyInt64(resource.MinInstances))
	p.SetMaxInstances(dcl.ValueOrEmptyInt64(resource.MaxInstances))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sConnectedProjects := make([]string, len(resource.ConnectedProjects))
	for i, r := range resource.ConnectedProjects {
		sConnectedProjects[i] = r
	}
	p.SetConnectedProjects(sConnectedProjects)

	return p
}

// applyConnector handles the gRPC request by passing it to the underlying Connector Apply() method.
func (s *ConnectorServer) applyConnector(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVpcaccessAlphaConnectorRequest) (*alphapb.VpcaccessAlphaConnector, error) {
	p := ProtoToConnector(request.GetResource())
	res, err := c.ApplyConnector(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConnectorToProto(res)
	return r, nil
}

// applyVpcaccessAlphaConnector handles the gRPC request by passing it to the underlying Connector Apply() method.
func (s *ConnectorServer) ApplyVpcaccessAlphaConnector(ctx context.Context, request *alphapb.ApplyVpcaccessAlphaConnectorRequest) (*alphapb.VpcaccessAlphaConnector, error) {
	cl, err := createConfigConnector(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyConnector(ctx, cl, request)
}

// DeleteConnector handles the gRPC request by passing it to the underlying Connector Delete() method.
func (s *ConnectorServer) DeleteVpcaccessAlphaConnector(ctx context.Context, request *alphapb.DeleteVpcaccessAlphaConnectorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConnector(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConnector(ctx, ProtoToConnector(request.GetResource()))

}

// ListVpcaccessAlphaConnector handles the gRPC request by passing it to the underlying ConnectorList() method.
func (s *ConnectorServer) ListVpcaccessAlphaConnector(ctx context.Context, request *alphapb.ListVpcaccessAlphaConnectorRequest) (*alphapb.ListVpcaccessAlphaConnectorResponse, error) {
	cl, err := createConfigConnector(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConnector(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VpcaccessAlphaConnector
	for _, r := range resources.Items {
		rp := ConnectorToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVpcaccessAlphaConnectorResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigConnector(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
