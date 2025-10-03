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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networksecurity/alpha/networksecurity_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networksecurity/alpha"
)

// ClientTlsPolicyServer implements the gRPC interface for ClientTlsPolicy.
type ClientTlsPolicyServer struct{}

// ProtoToClientTlsPolicyClientCertificate converts a ClientTlsPolicyClientCertificate object from its proto representation.
func ProtoToNetworksecurityAlphaClientTlsPolicyClientCertificate(p *alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificate) *alpha.ClientTlsPolicyClientCertificate {
	if p == nil {
		return nil
	}
	obj := &alpha.ClientTlsPolicyClientCertificate{
		LocalFilepath:               ProtoToNetworksecurityAlphaClientTlsPolicyClientCertificateLocalFilepath(p.GetLocalFilepath()),
		GrpcEndpoint:                ProtoToNetworksecurityAlphaClientTlsPolicyClientCertificateGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityAlphaClientTlsPolicyClientCertificateCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicyClientCertificateLocalFilepath converts a ClientTlsPolicyClientCertificateLocalFilepath object from its proto representation.
func ProtoToNetworksecurityAlphaClientTlsPolicyClientCertificateLocalFilepath(p *alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateLocalFilepath) *alpha.ClientTlsPolicyClientCertificateLocalFilepath {
	if p == nil {
		return nil
	}
	obj := &alpha.ClientTlsPolicyClientCertificateLocalFilepath{
		CertificatePath: dcl.StringOrNil(p.GetCertificatePath()),
		PrivateKeyPath:  dcl.StringOrNil(p.GetPrivateKeyPath()),
	}
	return obj
}

// ProtoToClientTlsPolicyClientCertificateGrpcEndpoint converts a ClientTlsPolicyClientCertificateGrpcEndpoint object from its proto representation.
func ProtoToNetworksecurityAlphaClientTlsPolicyClientCertificateGrpcEndpoint(p *alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateGrpcEndpoint) *alpha.ClientTlsPolicyClientCertificateGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.ClientTlsPolicyClientCertificateGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.GetTargetUri()),
	}
	return obj
}

// ProtoToClientTlsPolicyClientCertificateCertificateProviderInstance converts a ClientTlsPolicyClientCertificateCertificateProviderInstance object from its proto representation.
func ProtoToNetworksecurityAlphaClientTlsPolicyClientCertificateCertificateProviderInstance(p *alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateCertificateProviderInstance) *alpha.ClientTlsPolicyClientCertificateCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &alpha.ClientTlsPolicyClientCertificateCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.GetPluginInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCa converts a ClientTlsPolicyServerValidationCa object from its proto representation.
func ProtoToNetworksecurityAlphaClientTlsPolicyServerValidationCa(p *alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCa) *alpha.ClientTlsPolicyServerValidationCa {
	if p == nil {
		return nil
	}
	obj := &alpha.ClientTlsPolicyServerValidationCa{
		CaCertPath:                  dcl.StringOrNil(p.GetCaCertPath()),
		GrpcEndpoint:                ProtoToNetworksecurityAlphaClientTlsPolicyServerValidationCaGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityAlphaClientTlsPolicyServerValidationCaCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCaGrpcEndpoint converts a ClientTlsPolicyServerValidationCaGrpcEndpoint object from its proto representation.
func ProtoToNetworksecurityAlphaClientTlsPolicyServerValidationCaGrpcEndpoint(p *alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCaGrpcEndpoint) *alpha.ClientTlsPolicyServerValidationCaGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.ClientTlsPolicyServerValidationCaGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.GetTargetUri()),
	}
	return obj
}

// ProtoToClientTlsPolicyServerValidationCaCertificateProviderInstance converts a ClientTlsPolicyServerValidationCaCertificateProviderInstance object from its proto representation.
func ProtoToNetworksecurityAlphaClientTlsPolicyServerValidationCaCertificateProviderInstance(p *alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCaCertificateProviderInstance) *alpha.ClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &alpha.ClientTlsPolicyServerValidationCaCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.GetPluginInstance()),
	}
	return obj
}

// ProtoToClientTlsPolicy converts a ClientTlsPolicy resource from its proto representation.
func ProtoToClientTlsPolicy(p *alphapb.NetworksecurityAlphaClientTlsPolicy) *alpha.ClientTlsPolicy {
	obj := &alpha.ClientTlsPolicy{
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		Sni:               dcl.StringOrNil(p.GetSni()),
		ClientCertificate: ProtoToNetworksecurityAlphaClientTlsPolicyClientCertificate(p.GetClientCertificate()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetServerValidationCa() {
		obj.ServerValidationCa = append(obj.ServerValidationCa, *ProtoToNetworksecurityAlphaClientTlsPolicyServerValidationCa(r))
	}
	return obj
}

// ClientTlsPolicyClientCertificateToProto converts a ClientTlsPolicyClientCertificate object to its proto representation.
func NetworksecurityAlphaClientTlsPolicyClientCertificateToProto(o *alpha.ClientTlsPolicyClientCertificate) *alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificate {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificate{}
	p.SetLocalFilepath(NetworksecurityAlphaClientTlsPolicyClientCertificateLocalFilepathToProto(o.LocalFilepath))
	p.SetGrpcEndpoint(NetworksecurityAlphaClientTlsPolicyClientCertificateGrpcEndpointToProto(o.GrpcEndpoint))
	p.SetCertificateProviderInstance(NetworksecurityAlphaClientTlsPolicyClientCertificateCertificateProviderInstanceToProto(o.CertificateProviderInstance))
	return p
}

// ClientTlsPolicyClientCertificateLocalFilepathToProto converts a ClientTlsPolicyClientCertificateLocalFilepath object to its proto representation.
func NetworksecurityAlphaClientTlsPolicyClientCertificateLocalFilepathToProto(o *alpha.ClientTlsPolicyClientCertificateLocalFilepath) *alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateLocalFilepath {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateLocalFilepath{}
	p.SetCertificatePath(dcl.ValueOrEmptyString(o.CertificatePath))
	p.SetPrivateKeyPath(dcl.ValueOrEmptyString(o.PrivateKeyPath))
	return p
}

// ClientTlsPolicyClientCertificateGrpcEndpointToProto converts a ClientTlsPolicyClientCertificateGrpcEndpoint object to its proto representation.
func NetworksecurityAlphaClientTlsPolicyClientCertificateGrpcEndpointToProto(o *alpha.ClientTlsPolicyClientCertificateGrpcEndpoint) *alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateGrpcEndpoint{}
	p.SetTargetUri(dcl.ValueOrEmptyString(o.TargetUri))
	return p
}

// ClientTlsPolicyClientCertificateCertificateProviderInstanceToProto converts a ClientTlsPolicyClientCertificateCertificateProviderInstance object to its proto representation.
func NetworksecurityAlphaClientTlsPolicyClientCertificateCertificateProviderInstanceToProto(o *alpha.ClientTlsPolicyClientCertificateCertificateProviderInstance) *alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaClientTlsPolicyClientCertificateCertificateProviderInstance{}
	p.SetPluginInstance(dcl.ValueOrEmptyString(o.PluginInstance))
	return p
}

// ClientTlsPolicyServerValidationCaToProto converts a ClientTlsPolicyServerValidationCa object to its proto representation.
func NetworksecurityAlphaClientTlsPolicyServerValidationCaToProto(o *alpha.ClientTlsPolicyServerValidationCa) *alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCa {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCa{}
	p.SetCaCertPath(dcl.ValueOrEmptyString(o.CaCertPath))
	p.SetGrpcEndpoint(NetworksecurityAlphaClientTlsPolicyServerValidationCaGrpcEndpointToProto(o.GrpcEndpoint))
	p.SetCertificateProviderInstance(NetworksecurityAlphaClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto(o.CertificateProviderInstance))
	return p
}

// ClientTlsPolicyServerValidationCaGrpcEndpointToProto converts a ClientTlsPolicyServerValidationCaGrpcEndpoint object to its proto representation.
func NetworksecurityAlphaClientTlsPolicyServerValidationCaGrpcEndpointToProto(o *alpha.ClientTlsPolicyServerValidationCaGrpcEndpoint) *alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCaGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCaGrpcEndpoint{}
	p.SetTargetUri(dcl.ValueOrEmptyString(o.TargetUri))
	return p
}

// ClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto converts a ClientTlsPolicyServerValidationCaCertificateProviderInstance object to its proto representation.
func NetworksecurityAlphaClientTlsPolicyServerValidationCaCertificateProviderInstanceToProto(o *alpha.ClientTlsPolicyServerValidationCaCertificateProviderInstance) *alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCaCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCaCertificateProviderInstance{}
	p.SetPluginInstance(dcl.ValueOrEmptyString(o.PluginInstance))
	return p
}

// ClientTlsPolicyToProto converts a ClientTlsPolicy resource to its proto representation.
func ClientTlsPolicyToProto(resource *alpha.ClientTlsPolicy) *alphapb.NetworksecurityAlphaClientTlsPolicy {
	p := &alphapb.NetworksecurityAlphaClientTlsPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetSni(dcl.ValueOrEmptyString(resource.Sni))
	p.SetClientCertificate(NetworksecurityAlphaClientTlsPolicyClientCertificateToProto(resource.ClientCertificate))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sServerValidationCa := make([]*alphapb.NetworksecurityAlphaClientTlsPolicyServerValidationCa, len(resource.ServerValidationCa))
	for i, r := range resource.ServerValidationCa {
		sServerValidationCa[i] = NetworksecurityAlphaClientTlsPolicyServerValidationCaToProto(&r)
	}
	p.SetServerValidationCa(sServerValidationCa)

	return p
}

// applyClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Apply() method.
func (s *ClientTlsPolicyServer) applyClientTlsPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworksecurityAlphaClientTlsPolicyRequest) (*alphapb.NetworksecurityAlphaClientTlsPolicy, error) {
	p := ProtoToClientTlsPolicy(request.GetResource())
	res, err := c.ApplyClientTlsPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ClientTlsPolicyToProto(res)
	return r, nil
}

// applyNetworksecurityAlphaClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Apply() method.
func (s *ClientTlsPolicyServer) ApplyNetworksecurityAlphaClientTlsPolicy(ctx context.Context, request *alphapb.ApplyNetworksecurityAlphaClientTlsPolicyRequest) (*alphapb.NetworksecurityAlphaClientTlsPolicy, error) {
	cl, err := createConfigClientTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyClientTlsPolicy(ctx, cl, request)
}

// DeleteClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicy Delete() method.
func (s *ClientTlsPolicyServer) DeleteNetworksecurityAlphaClientTlsPolicy(ctx context.Context, request *alphapb.DeleteNetworksecurityAlphaClientTlsPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigClientTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteClientTlsPolicy(ctx, ProtoToClientTlsPolicy(request.GetResource()))

}

// ListNetworksecurityAlphaClientTlsPolicy handles the gRPC request by passing it to the underlying ClientTlsPolicyList() method.
func (s *ClientTlsPolicyServer) ListNetworksecurityAlphaClientTlsPolicy(ctx context.Context, request *alphapb.ListNetworksecurityAlphaClientTlsPolicyRequest) (*alphapb.ListNetworksecurityAlphaClientTlsPolicyResponse, error) {
	cl, err := createConfigClientTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListClientTlsPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworksecurityAlphaClientTlsPolicy
	for _, r := range resources.Items {
		rp := ClientTlsPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworksecurityAlphaClientTlsPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigClientTlsPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
