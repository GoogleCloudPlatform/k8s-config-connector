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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/beta/networksecurity_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/beta"
)

// ServerTlsPolicyServer implements the gRPC interface for ServerTlsPolicy.
type ServerTlsPolicyServer struct{}

// ProtoToServerTlsPolicyServerCertificate converts a ServerTlsPolicyServerCertificate object from its proto representation.
func ProtoToNetworksecurityBetaServerTlsPolicyServerCertificate(p *betapb.NetworksecurityBetaServerTlsPolicyServerCertificate) *beta.ServerTlsPolicyServerCertificate {
	if p == nil {
		return nil
	}
	obj := &beta.ServerTlsPolicyServerCertificate{
		GrpcEndpoint:                ProtoToNetworksecurityBetaServerTlsPolicyServerCertificateGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityBetaServerTlsPolicyServerCertificateCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicyServerCertificateGrpcEndpoint converts a ServerTlsPolicyServerCertificateGrpcEndpoint object from its proto representation.
func ProtoToNetworksecurityBetaServerTlsPolicyServerCertificateGrpcEndpoint(p *betapb.NetworksecurityBetaServerTlsPolicyServerCertificateGrpcEndpoint) *beta.ServerTlsPolicyServerCertificateGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.ServerTlsPolicyServerCertificateGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.GetTargetUri()),
	}
	return obj
}

// ProtoToServerTlsPolicyServerCertificateCertificateProviderInstance converts a ServerTlsPolicyServerCertificateCertificateProviderInstance object from its proto representation.
func ProtoToNetworksecurityBetaServerTlsPolicyServerCertificateCertificateProviderInstance(p *betapb.NetworksecurityBetaServerTlsPolicyServerCertificateCertificateProviderInstance) *beta.ServerTlsPolicyServerCertificateCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &beta.ServerTlsPolicyServerCertificateCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.GetPluginInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicy converts a ServerTlsPolicyMtlsPolicy object from its proto representation.
func ProtoToNetworksecurityBetaServerTlsPolicyMtlsPolicy(p *betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicy) *beta.ServerTlsPolicyMtlsPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.ServerTlsPolicyMtlsPolicy{}
	for _, r := range p.GetClientValidationCa() {
		obj.ClientValidationCa = append(obj.ClientValidationCa, *ProtoToNetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCa(r))
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCa converts a ServerTlsPolicyMtlsPolicyClientValidationCa object from its proto representation.
func ProtoToNetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCa(p *betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCa) *beta.ServerTlsPolicyMtlsPolicyClientValidationCa {
	if p == nil {
		return nil
	}
	obj := &beta.ServerTlsPolicyMtlsPolicyClientValidationCa{
		GrpcEndpoint:                ProtoToNetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint converts a ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint object from its proto representation.
func ProtoToNetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(p *betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) *beta.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.GetTargetUri()),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance converts a ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance object from its proto representation.
func ProtoToNetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(p *betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) *beta.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &beta.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.GetPluginInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicy converts a ServerTlsPolicy resource from its proto representation.
func ProtoToServerTlsPolicy(p *betapb.NetworksecurityBetaServerTlsPolicy) *beta.ServerTlsPolicy {
	obj := &beta.ServerTlsPolicy{
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		AllowOpen:         dcl.Bool(p.GetAllowOpen()),
		ServerCertificate: ProtoToNetworksecurityBetaServerTlsPolicyServerCertificate(p.GetServerCertificate()),
		MtlsPolicy:        ProtoToNetworksecurityBetaServerTlsPolicyMtlsPolicy(p.GetMtlsPolicy()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ServerTlsPolicyServerCertificateToProto converts a ServerTlsPolicyServerCertificate object to its proto representation.
func NetworksecurityBetaServerTlsPolicyServerCertificateToProto(o *beta.ServerTlsPolicyServerCertificate) *betapb.NetworksecurityBetaServerTlsPolicyServerCertificate {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaServerTlsPolicyServerCertificate{}
	p.SetGrpcEndpoint(NetworksecurityBetaServerTlsPolicyServerCertificateGrpcEndpointToProto(o.GrpcEndpoint))
	p.SetCertificateProviderInstance(NetworksecurityBetaServerTlsPolicyServerCertificateCertificateProviderInstanceToProto(o.CertificateProviderInstance))
	return p
}

// ServerTlsPolicyServerCertificateGrpcEndpointToProto converts a ServerTlsPolicyServerCertificateGrpcEndpoint object to its proto representation.
func NetworksecurityBetaServerTlsPolicyServerCertificateGrpcEndpointToProto(o *beta.ServerTlsPolicyServerCertificateGrpcEndpoint) *betapb.NetworksecurityBetaServerTlsPolicyServerCertificateGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaServerTlsPolicyServerCertificateGrpcEndpoint{}
	p.SetTargetUri(dcl.ValueOrEmptyString(o.TargetUri))
	return p
}

// ServerTlsPolicyServerCertificateCertificateProviderInstanceToProto converts a ServerTlsPolicyServerCertificateCertificateProviderInstance object to its proto representation.
func NetworksecurityBetaServerTlsPolicyServerCertificateCertificateProviderInstanceToProto(o *beta.ServerTlsPolicyServerCertificateCertificateProviderInstance) *betapb.NetworksecurityBetaServerTlsPolicyServerCertificateCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaServerTlsPolicyServerCertificateCertificateProviderInstance{}
	p.SetPluginInstance(dcl.ValueOrEmptyString(o.PluginInstance))
	return p
}

// ServerTlsPolicyMtlsPolicyToProto converts a ServerTlsPolicyMtlsPolicy object to its proto representation.
func NetworksecurityBetaServerTlsPolicyMtlsPolicyToProto(o *beta.ServerTlsPolicyMtlsPolicy) *betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicy{}
	sClientValidationCa := make([]*betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCa, len(o.ClientValidationCa))
	for i, r := range o.ClientValidationCa {
		sClientValidationCa[i] = NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaToProto(&r)
	}
	p.SetClientValidationCa(sClientValidationCa)
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCa object to its proto representation.
func NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaToProto(o *beta.ServerTlsPolicyMtlsPolicyClientValidationCa) *betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCa {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCa{}
	p.SetGrpcEndpoint(NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto(o.GrpcEndpoint))
	p.SetCertificateProviderInstance(NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto(o.CertificateProviderInstance))
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint object to its proto representation.
func NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto(o *beta.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) *betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{}
	p.SetTargetUri(dcl.ValueOrEmptyString(o.TargetUri))
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance object to its proto representation.
func NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto(o *beta.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) *betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{}
	p.SetPluginInstance(dcl.ValueOrEmptyString(o.PluginInstance))
	return p
}

// ServerTlsPolicyToProto converts a ServerTlsPolicy resource to its proto representation.
func ServerTlsPolicyToProto(resource *beta.ServerTlsPolicy) *betapb.NetworksecurityBetaServerTlsPolicy {
	p := &betapb.NetworksecurityBetaServerTlsPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetAllowOpen(dcl.ValueOrEmptyBool(resource.AllowOpen))
	p.SetServerCertificate(NetworksecurityBetaServerTlsPolicyServerCertificateToProto(resource.ServerCertificate))
	p.SetMtlsPolicy(NetworksecurityBetaServerTlsPolicyMtlsPolicyToProto(resource.MtlsPolicy))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicy Apply() method.
func (s *ServerTlsPolicyServer) applyServerTlsPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworksecurityBetaServerTlsPolicyRequest) (*betapb.NetworksecurityBetaServerTlsPolicy, error) {
	p := ProtoToServerTlsPolicy(request.GetResource())
	res, err := c.ApplyServerTlsPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServerTlsPolicyToProto(res)
	return r, nil
}

// applyNetworksecurityBetaServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicy Apply() method.
func (s *ServerTlsPolicyServer) ApplyNetworksecurityBetaServerTlsPolicy(ctx context.Context, request *betapb.ApplyNetworksecurityBetaServerTlsPolicyRequest) (*betapb.NetworksecurityBetaServerTlsPolicy, error) {
	cl, err := createConfigServerTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyServerTlsPolicy(ctx, cl, request)
}

// DeleteServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicy Delete() method.
func (s *ServerTlsPolicyServer) DeleteNetworksecurityBetaServerTlsPolicy(ctx context.Context, request *betapb.DeleteNetworksecurityBetaServerTlsPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServerTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServerTlsPolicy(ctx, ProtoToServerTlsPolicy(request.GetResource()))

}

// ListNetworksecurityBetaServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicyList() method.
func (s *ServerTlsPolicyServer) ListNetworksecurityBetaServerTlsPolicy(ctx context.Context, request *betapb.ListNetworksecurityBetaServerTlsPolicyRequest) (*betapb.ListNetworksecurityBetaServerTlsPolicyResponse, error) {
	cl, err := createConfigServerTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServerTlsPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworksecurityBetaServerTlsPolicy
	for _, r := range resources.Items {
		rp := ServerTlsPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworksecurityBetaServerTlsPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigServerTlsPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
