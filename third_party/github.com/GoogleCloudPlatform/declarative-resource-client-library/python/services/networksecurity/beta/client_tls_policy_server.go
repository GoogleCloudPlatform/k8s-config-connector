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

// ClientTlsPolicyServer implements the gRPC interface for ClientTlsPolicy.
type ClientTlsPolicyServer struct{}

// ProtoToClientTlsPolicyClientCertificate converts a ClientTlsPolicyClientCertificate object from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyClientCertificate(p *betapb.NetworksecurityBetaClientTlsPolicyClientCertificate) *beta.ClientTlsPolicyClientCertificate {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyClientCertificate{
		GrpcEndpoint:                ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicyClientCertificateGrpcEndpoint converts a ClientTlsPolicyClientCertificateGrpcEndpoint object from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint(p *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint) *beta.ClientTlsPolicyClientCertificateGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyClientCertificateGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.GetTargetUri()),
	}
	return obj
}

// ProtoToClientTlsPolicyClientCertificateCertificateProviderInstance converts a ClientTlsPolicyClientCertificateCertificateProviderInstance object from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance(p *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance) *beta.ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyClientCertificateCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.GetPluginInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCa converts a ClientTlsPolicyServerValidationCa object from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCa(p *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCa) *beta.ClientTlsPolicyServerValidationCa {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyServerValidationCa{
		GrpcEndpoint:                ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCaGrpcEndpoint converts a ClientTlsPolicyServerValidationCaGrpcEndpoint object from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint(p *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint) *beta.ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyServerValidationCaGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.GetTargetUri()),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCaCertificateProviderInstance converts a ClientTlsPolicyServerValidationCaCertificateProviderInstance object from its proto representation.
func ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance(p *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance) *beta.ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &beta.ClientTlsPolicyServerValidationCaCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.GetPluginInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicy converts a ClientTlsPolicy resource from its proto representation.
func ProtoToClientTlsPolicy(p *betapb.NetworksecurityBetaClientTlsPolicy) *beta.ClientTlsPolicy {
	obj := &beta.ClientTlsPolicy{
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Sni:               dcl.StringOrNil(p.GetSni()),
		ClientCertificate: ProtoToNetworksecurityBetaClientTlsPolicyClientCertificate(p.GetClientCertificate()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetServerValidationCa() {
		obj.ServerValidationCa = append(obj.ServerValidationCa, *ProtoToNetworksecurityBetaClientTlsPolicyServerValidationCa(r))
	}
	return obj
}

// ClientTlsPolicyClientCertificateToProto converts a ClientTlsPolicyClientCertificate object to its proto representation.
func NetworksecurityBetaClientTlsPolicyClientCertificateToProto(o *beta.ClientTlsPolicyClientCertificate) *betapb.NetworksecurityBetaClientTlsPolicyClientCertificate {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyClientCertificate{}
	p.SetGrpcEndpoint(NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpointToProto(o.GrpcEndpoint))
	p.SetCertificateProviderInstance(NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstanceToProto(o.CertificateProviderInstance))
	return p
}

// ClientTlsPolicyClientCertificateGrpcEndpointToProto converts a ClientTlsPolicyClientCertificateGrpcEndpoint object to its proto representation.
func NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpointToProto(o *beta.ClientTlsPolicyClientCertificateGrpcEndpoint) *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyClientCertificateGrpcEndpoint{}
	p.SetTargetUri(dcl.ValueOrEmptyString(o.TargetUri))
	return p
}

// ClientTlsPolicyClientCertificateCertificateProviderInstanceToProto converts a ClientTlsPolicyClientCertificateCertificateProviderInstance object to its proto representation.
func NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstanceToProto(o *beta.ClientTlsPolicyClientCertificateCertificateProviderInstance) *betapb.NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyClientCertificateCertificateProviderInstance{}
	p.SetPluginInstance(dcl.ValueOrEmptyString(o.PluginInstance))
	return p
}

// ClientTlsPolicyServerValidationCaToProto converts a ClientTlsPolicyServerValidationCa object to its proto representation.
func NetworksecurityBetaClientTlsPolicyServerValidationCaToProto(o *beta.ClientTlsPolicyServerValidationCa) *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCa {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyServerValidationCa{}
	p.SetGrpcEndpoint(NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpointToProto(o.GrpcEndpoint))
	p.SetCertificateProviderInstance(NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto(o.CertificateProviderInstance))
	return p
}

// ClientTlsPolicyServerValidationCaGrpcEndpointToProto converts a ClientTlsPolicyServerValidationCaGrpcEndpoint object to its proto representation.
func NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpointToProto(o *beta.ClientTlsPolicyServerValidationCaGrpcEndpoint) *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaGrpcEndpoint{}
	p.SetTargetUri(dcl.ValueOrEmptyString(o.TargetUri))
	return p
}

// ClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto converts a ClientTlsPolicyServerValidationCaCertificateProviderInstance object to its proto representation.
func NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto(o *beta.ClientTlsPolicyServerValidationCaCertificateProviderInstance) *betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &betapb.NetworksecurityBetaClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	p.SetPluginInstance(dcl.ValueOrEmptyString(o.PluginInstance))
	return p
}

// ClientTlsPolicyToProto converts a ClientTlsPolicy resource to its proto representation.
func ClientTlsPolicyToProto(resource *beta.ClientTlsPolicy) *betapb.NetworksecurityBetaClientTlsPolicy {
	p := &betapb.NetworksecurityBetaClientTlsPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetSni(dcl.ValueOrEmptyString(resource.Sni))
	p.SetClientCertificate(NetworksecurityBetaClientTlsPolicyClientCertificateToProto(resource.ClientCertificate))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sServerValidationCa := make([]*betapb.NetworksecurityBetaClientTlsPolicyServerValidationCa, len(resource.ServerValidationCa))
	for i, r := range resource.ServerValidationCa {
		sServerValidationCa[i] = NetworksecurityBetaClientTlsPolicyServerValidationCaToProto(&r)
	}
	p.SetServerValidationCa(sServerValidationCa)

	return p
}

// applyClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Apply() method.
func (s *ClientTlsPolicyServer) applyClientTlsPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworksecurityBetaClientTlsPolicyRequest) (*betapb.NetworksecurityBetaClientTlsPolicy, error) {
	p := ProtoToClientTlsPolicy(request.GetResource())
	res, err := c.ApplyClientTlsPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClientTlsPolicyToProto(res)
	return r, nil
}

// applyNetworksecurityBetaClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Apply() method.
func (s *ClientTlsPolicyServer) ApplyNetworksecurityBetaClientTlsPolicy(ctx context.Context, request *betapb.ApplyNetworksecurityBetaClientTlsPolicyRequest) (*betapb.NetworksecurityBetaClientTlsPolicy, error) {
	cl, err := createConfigClientTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyClientTlsPolicy(ctx, cl, request)
}

// DeleteClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Delete() method.
func (s *ClientTlsPolicyServer) DeleteNetworksecurityBetaClientTlsPolicy(ctx context.Context, request *betapb.DeleteNetworksecurityBetaClientTlsPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigClientTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteClientTlsPolicy(ctx, ProtoToClientTlsPolicy(request.GetResource()))

}

// ListNetworksecurityBetaClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicyList() method.
func (s *ClientTlsPolicyServer) ListNetworksecurityBetaClientTlsPolicy(ctx context.Context, request *betapb.ListNetworksecurityBetaClientTlsPolicyRequest) (*betapb.ListNetworksecurityBetaClientTlsPolicyResponse, error) {
	cl, err := createConfigClientTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListClientTlsPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworksecurityBetaClientTlsPolicy
	for _, r := range resources.Items {
		rp := ClientTlsPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworksecurityBetaClientTlsPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigClientTlsPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
