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

// ServerTlsPolicyServer implements the gRPC interface for ServerTlsPolicy.
type ServerTlsPolicyServer struct{}

// ProtoToServerTlsPolicyServerCertificate converts a ServerTlsPolicyServerCertificate object from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificate(p *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificate) *alpha.ServerTlsPolicyServerCertificate {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyServerCertificate{
		LocalFilepath:               ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath(p.GetLocalFilepath()),
		GrpcEndpoint:                ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicyServerCertificateLocalFilepath converts a ServerTlsPolicyServerCertificateLocalFilepath object from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath(p *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath) *alpha.ServerTlsPolicyServerCertificateLocalFilepath {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyServerCertificateLocalFilepath{
		CertificatePath: dcl.StringOrNil(p.GetCertificatePath()),
		PrivateKeyPath:  dcl.StringOrNil(p.GetPrivateKeyPath()),
	}
	return obj
}

// ProtoToServerTlsPolicyServerCertificateGrpcEndpoint converts a ServerTlsPolicyServerCertificateGrpcEndpoint object from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint(p *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint) *alpha.ServerTlsPolicyServerCertificateGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyServerCertificateGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.GetTargetUri()),
	}
	return obj
}

// ProtoToServerTlsPolicyServerCertificateCertificateProviderInstance converts a ServerTlsPolicyServerCertificateCertificateProviderInstance object from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance(p *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance) *alpha.ServerTlsPolicyServerCertificateCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyServerCertificateCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.GetPluginInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicy converts a ServerTlsPolicyMtlsPolicy object from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicy(p *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicy) *alpha.ServerTlsPolicyMtlsPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyMtlsPolicy{}
	for _, r := range p.GetClientValidationCa() {
		obj.ClientValidationCa = append(obj.ClientValidationCa, *ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa(r))
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCa converts a ServerTlsPolicyMtlsPolicyClientValidationCa object from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa(p *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa) *alpha.ServerTlsPolicyMtlsPolicyClientValidationCa {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyMtlsPolicyClientValidationCa{
		CaCertPath:                  dcl.StringOrNil(p.GetCaCertPath()),
		GrpcEndpoint:                ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(p.GetGrpcEndpoint()),
		CertificateProviderInstance: ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(p.GetCertificateProviderInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint converts a ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint object from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint(p *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) *alpha.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{
		TargetUri: dcl.StringOrNil(p.GetTargetUri()),
	}
	return obj
}

// ProtoToServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance converts a ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance object from its proto representation.
func ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance(p *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) *alpha.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if p == nil {
		return nil
	}
	obj := &alpha.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{
		PluginInstance: dcl.StringOrNil(p.GetPluginInstance()),
	}
	return obj
}

// ProtoToServerTlsPolicy converts a ServerTlsPolicy resource from its proto representation.
func ProtoToServerTlsPolicy(p *alphapb.NetworksecurityAlphaServerTlsPolicy) *alpha.ServerTlsPolicy {
	obj := &alpha.ServerTlsPolicy{
		Name:              dcl.StringOrNil(p.GetName()),
		Description:       dcl.StringOrNil(p.GetDescription()),
		CreateTime:        dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:        dcl.StringOrNil(p.GetUpdateTime()),
		AllowOpen:         dcl.Bool(p.GetAllowOpen()),
		ServerCertificate: ProtoToNetworksecurityAlphaServerTlsPolicyServerCertificate(p.GetServerCertificate()),
		MtlsPolicy:        ProtoToNetworksecurityAlphaServerTlsPolicyMtlsPolicy(p.GetMtlsPolicy()),
		Project:           dcl.StringOrNil(p.GetProject()),
		Location:          dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ServerTlsPolicyServerCertificateToProto converts a ServerTlsPolicyServerCertificate object to its proto representation.
func NetworksecurityAlphaServerTlsPolicyServerCertificateToProto(o *alpha.ServerTlsPolicyServerCertificate) *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificate {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificate{}
	p.SetLocalFilepath(NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepathToProto(o.LocalFilepath))
	p.SetGrpcEndpoint(NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpointToProto(o.GrpcEndpoint))
	p.SetCertificateProviderInstance(NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstanceToProto(o.CertificateProviderInstance))
	return p
}

// ServerTlsPolicyServerCertificateLocalFilepathToProto converts a ServerTlsPolicyServerCertificateLocalFilepath object to its proto representation.
func NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepathToProto(o *alpha.ServerTlsPolicyServerCertificateLocalFilepath) *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateLocalFilepath{}
	p.SetCertificatePath(dcl.ValueOrEmptyString(o.CertificatePath))
	p.SetPrivateKeyPath(dcl.ValueOrEmptyString(o.PrivateKeyPath))
	return p
}

// ServerTlsPolicyServerCertificateGrpcEndpointToProto converts a ServerTlsPolicyServerCertificateGrpcEndpoint object to its proto representation.
func NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpointToProto(o *alpha.ServerTlsPolicyServerCertificateGrpcEndpoint) *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateGrpcEndpoint{}
	p.SetTargetUri(dcl.ValueOrEmptyString(o.TargetUri))
	return p
}

// ServerTlsPolicyServerCertificateCertificateProviderInstanceToProto converts a ServerTlsPolicyServerCertificateCertificateProviderInstance object to its proto representation.
func NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstanceToProto(o *alpha.ServerTlsPolicyServerCertificateCertificateProviderInstance) *alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyServerCertificateCertificateProviderInstance{}
	p.SetPluginInstance(dcl.ValueOrEmptyString(o.PluginInstance))
	return p
}

// ServerTlsPolicyMtlsPolicyToProto converts a ServerTlsPolicyMtlsPolicy object to its proto representation.
func NetworksecurityAlphaServerTlsPolicyMtlsPolicyToProto(o *alpha.ServerTlsPolicyMtlsPolicy) *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicy{}
	sClientValidationCa := make([]*alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa, len(o.ClientValidationCa))
	for i, r := range o.ClientValidationCa {
		sClientValidationCa[i] = NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaToProto(&r)
	}
	p.SetClientValidationCa(sClientValidationCa)
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCa object to its proto representation.
func NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaToProto(o *alpha.ServerTlsPolicyMtlsPolicyClientValidationCa) *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCa{}
	p.SetCaCertPath(dcl.ValueOrEmptyString(o.CaCertPath))
	p.SetGrpcEndpoint(NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto(o.GrpcEndpoint))
	p.SetCertificateProviderInstance(NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto(o.CertificateProviderInstance))
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint object to its proto representation.
func NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpointToProto(o *alpha.ServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint) *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint{}
	p.SetTargetUri(dcl.ValueOrEmptyString(o.TargetUri))
	return p
}

// ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto converts a ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance object to its proto representation.
func NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstanceToProto(o *alpha.ServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance) *alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworksecurityAlphaServerTlsPolicyMtlsPolicyClientValidationCaCertificateProviderInstance{}
	p.SetPluginInstance(dcl.ValueOrEmptyString(o.PluginInstance))
	return p
}

// ServerTlsPolicyToProto converts a ServerTlsPolicy resource to its proto representation.
func ServerTlsPolicyToProto(resource *alpha.ServerTlsPolicy) *alphapb.NetworksecurityAlphaServerTlsPolicy {
	p := &alphapb.NetworksecurityAlphaServerTlsPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetAllowOpen(dcl.ValueOrEmptyBool(resource.AllowOpen))
	p.SetServerCertificate(NetworksecurityAlphaServerTlsPolicyServerCertificateToProto(resource.ServerCertificate))
	p.SetMtlsPolicy(NetworksecurityAlphaServerTlsPolicyMtlsPolicyToProto(resource.MtlsPolicy))
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
func (s *ServerTlsPolicyServer) applyServerTlsPolicy(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworksecurityAlphaServerTlsPolicyRequest) (*alphapb.NetworksecurityAlphaServerTlsPolicy, error) {
	p := ProtoToServerTlsPolicy(request.GetResource())
	res, err := c.ApplyServerTlsPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ServerTlsPolicyToProto(res)
	return r, nil
}

// applyNetworksecurityAlphaServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicy Apply() method.
func (s *ServerTlsPolicyServer) ApplyNetworksecurityAlphaServerTlsPolicy(ctx context.Context, request *alphapb.ApplyNetworksecurityAlphaServerTlsPolicyRequest) (*alphapb.NetworksecurityAlphaServerTlsPolicy, error) {
	cl, err := createConfigServerTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyServerTlsPolicy(ctx, cl, request)
}

// DeleteServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicy Delete() method.
func (s *ServerTlsPolicyServer) DeleteNetworksecurityAlphaServerTlsPolicy(ctx context.Context, request *alphapb.DeleteNetworksecurityAlphaServerTlsPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigServerTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteServerTlsPolicy(ctx, ProtoToServerTlsPolicy(request.GetResource()))

}

// ListNetworksecurityAlphaServerTlsPolicy handles the gRPC request by passing it to the underlying ServerTlsPolicyList() method.
func (s *ServerTlsPolicyServer) ListNetworksecurityAlphaServerTlsPolicy(ctx context.Context, request *alphapb.ListNetworksecurityAlphaServerTlsPolicyRequest) (*alphapb.ListNetworksecurityAlphaServerTlsPolicyResponse, error) {
	cl, err := createConfigServerTlsPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListServerTlsPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworksecurityAlphaServerTlsPolicy
	for _, r := range resources.Items {
		rp := ServerTlsPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworksecurityAlphaServerTlsPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigServerTlsPolicy(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
