// Copyright 2022 Google LLC. All Rights Reserved.
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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/alpha/vertexai_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/alpha"
)

// ModelDeploymentServer implements the gRPC interface for ModelDeployment.
type ModelDeploymentServer struct{}

// ProtoToModelDeploymentDedicatedResources converts a ModelDeploymentDedicatedResources object from its proto representation.
func ProtoToVertexaiAlphaModelDeploymentDedicatedResources(p *alphapb.VertexaiAlphaModelDeploymentDedicatedResources) *alpha.ModelDeploymentDedicatedResources {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelDeploymentDedicatedResources{
		MachineSpec:     ProtoToVertexaiAlphaModelDeploymentDedicatedResourcesMachineSpec(p.GetMachineSpec()),
		MinReplicaCount: dcl.Int64OrNil(p.GetMinReplicaCount()),
		MaxReplicaCount: dcl.Int64OrNil(p.GetMaxReplicaCount()),
	}
	return obj
}

// ProtoToModelDeploymentDedicatedResourcesMachineSpec converts a ModelDeploymentDedicatedResourcesMachineSpec object from its proto representation.
func ProtoToVertexaiAlphaModelDeploymentDedicatedResourcesMachineSpec(p *alphapb.VertexaiAlphaModelDeploymentDedicatedResourcesMachineSpec) *alpha.ModelDeploymentDedicatedResourcesMachineSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.ModelDeploymentDedicatedResourcesMachineSpec{
		MachineType: dcl.StringOrNil(p.GetMachineType()),
	}
	return obj
}

// ProtoToModelDeployment converts a ModelDeployment resource from its proto representation.
func ProtoToModelDeployment(p *alphapb.VertexaiAlphaModelDeployment) *alpha.ModelDeployment {
	obj := &alpha.ModelDeployment{
		Model:              dcl.StringOrNil(p.GetModel()),
		DeployedModelId:    dcl.StringOrNil(p.GetDeployedModelId()),
		DedicatedResources: ProtoToVertexaiAlphaModelDeploymentDedicatedResources(p.GetDedicatedResources()),
		Endpoint:           dcl.StringOrNil(p.GetEndpoint()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		Project:            dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ModelDeploymentDedicatedResourcesToProto converts a ModelDeploymentDedicatedResources object to its proto representation.
func VertexaiAlphaModelDeploymentDedicatedResourcesToProto(o *alpha.ModelDeploymentDedicatedResources) *alphapb.VertexaiAlphaModelDeploymentDedicatedResources {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelDeploymentDedicatedResources{}
	p.SetMachineSpec(VertexaiAlphaModelDeploymentDedicatedResourcesMachineSpecToProto(o.MachineSpec))
	p.SetMinReplicaCount(dcl.ValueOrEmptyInt64(o.MinReplicaCount))
	p.SetMaxReplicaCount(dcl.ValueOrEmptyInt64(o.MaxReplicaCount))
	return p
}

// ModelDeploymentDedicatedResourcesMachineSpecToProto converts a ModelDeploymentDedicatedResourcesMachineSpec object to its proto representation.
func VertexaiAlphaModelDeploymentDedicatedResourcesMachineSpecToProto(o *alpha.ModelDeploymentDedicatedResourcesMachineSpec) *alphapb.VertexaiAlphaModelDeploymentDedicatedResourcesMachineSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaModelDeploymentDedicatedResourcesMachineSpec{}
	p.SetMachineType(dcl.ValueOrEmptyString(o.MachineType))
	return p
}

// ModelDeploymentToProto converts a ModelDeployment resource to its proto representation.
func ModelDeploymentToProto(resource *alpha.ModelDeployment) *alphapb.VertexaiAlphaModelDeployment {
	p := &alphapb.VertexaiAlphaModelDeployment{}
	p.SetModel(dcl.ValueOrEmptyString(resource.Model))
	p.SetDeployedModelId(dcl.ValueOrEmptyString(resource.DeployedModelId))
	p.SetDedicatedResources(VertexaiAlphaModelDeploymentDedicatedResourcesToProto(resource.DedicatedResources))
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyModelDeployment handles the gRPC request by passing it to the underlying ModelDeployment Apply() method.
func (s *ModelDeploymentServer) applyModelDeployment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVertexaiAlphaModelDeploymentRequest) (*alphapb.VertexaiAlphaModelDeployment, error) {
	p := ProtoToModelDeployment(request.GetResource())
	res, err := c.ApplyModelDeployment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ModelDeploymentToProto(res)
	return r, nil
}

// applyVertexaiAlphaModelDeployment handles the gRPC request by passing it to the underlying ModelDeployment Apply() method.
func (s *ModelDeploymentServer) ApplyVertexaiAlphaModelDeployment(ctx context.Context, request *alphapb.ApplyVertexaiAlphaModelDeploymentRequest) (*alphapb.VertexaiAlphaModelDeployment, error) {
	cl, err := createConfigModelDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyModelDeployment(ctx, cl, request)
}

// DeleteModelDeployment handles the gRPC request by passing it to the underlying ModelDeployment Delete() method.
func (s *ModelDeploymentServer) DeleteVertexaiAlphaModelDeployment(ctx context.Context, request *alphapb.DeleteVertexaiAlphaModelDeploymentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigModelDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteModelDeployment(ctx, ProtoToModelDeployment(request.GetResource()))

}

// ListVertexaiAlphaModelDeployment handles the gRPC request by passing it to the underlying ModelDeploymentList() method.
func (s *ModelDeploymentServer) ListVertexaiAlphaModelDeployment(ctx context.Context, request *alphapb.ListVertexaiAlphaModelDeploymentRequest) (*alphapb.ListVertexaiAlphaModelDeploymentResponse, error) {
	cl, err := createConfigModelDeployment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListModelDeployment(ctx, request.GetProject(), request.GetLocation(), request.GetEndpoint())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VertexaiAlphaModelDeployment
	for _, r := range resources.Items {
		rp := ModelDeploymentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVertexaiAlphaModelDeploymentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigModelDeployment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
