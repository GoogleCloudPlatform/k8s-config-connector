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
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	cloudkmspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/cloudkms_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms"
)

// EkmConnectionServer implements the gRPC interface for EkmConnection.
type EkmConnectionServer struct{}

// ProtoToEkmConnectionServiceResolvers converts a EkmConnectionServiceResolvers object from its proto representation.
func ProtoToCloudkmsEkmConnectionServiceResolvers(p *cloudkmspb.CloudkmsEkmConnectionServiceResolvers) *cloudkms.EkmConnectionServiceResolvers {
	if p == nil {
		return nil
	}
	obj := &cloudkms.EkmConnectionServiceResolvers{
		ServiceDirectoryService: dcl.StringOrNil(p.GetServiceDirectoryService()),
		EndpointFilter:          dcl.StringOrNil(p.GetEndpointFilter()),
		Hostname:                dcl.StringOrNil(p.GetHostname()),
	}
	for _, r := range p.GetServerCertificates() {
		obj.ServerCertificates = append(obj.ServerCertificates, *ProtoToCloudkmsEkmConnectionServiceResolversServerCertificates(r))
	}
	return obj
}

// ProtoToEkmConnectionServiceResolversServerCertificates converts a EkmConnectionServiceResolversServerCertificates object from its proto representation.
func ProtoToCloudkmsEkmConnectionServiceResolversServerCertificates(p *cloudkmspb.CloudkmsEkmConnectionServiceResolversServerCertificates) *cloudkms.EkmConnectionServiceResolversServerCertificates {
	if p == nil {
		return nil
	}
	obj := &cloudkms.EkmConnectionServiceResolversServerCertificates{
		RawDer:            dcl.StringOrNil(p.GetRawDer()),
		Parsed:            dcl.Bool(p.GetParsed()),
		Issuer:            dcl.StringOrNil(p.GetIssuer()),
		Subject:           dcl.StringOrNil(p.GetSubject()),
		NotBeforeTime:     dcl.StringOrNil(p.GetNotBeforeTime()),
		NotAfterTime:      dcl.StringOrNil(p.GetNotAfterTime()),
		SerialNumber:      dcl.StringOrNil(p.GetSerialNumber()),
		Sha256Fingerprint: dcl.StringOrNil(p.GetSha256Fingerprint()),
	}
	for _, r := range p.GetSubjectAlternativeDnsNames() {
		obj.SubjectAlternativeDnsNames = append(obj.SubjectAlternativeDnsNames, r)
	}
	return obj
}

// ProtoToEkmConnection converts a EkmConnection resource from its proto representation.
func ProtoToEkmConnection(p *cloudkmspb.CloudkmsEkmConnection) *cloudkms.EkmConnection {
	obj := &cloudkms.EkmConnection{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		Etag:       dcl.StringOrNil(p.GetEtag()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	for _, r := range p.GetServiceResolvers() {
		obj.ServiceResolvers = append(obj.ServiceResolvers, *ProtoToCloudkmsEkmConnectionServiceResolvers(r))
	}
	return obj
}

// EkmConnectionServiceResolversToProto converts a EkmConnectionServiceResolvers object to its proto representation.
func CloudkmsEkmConnectionServiceResolversToProto(o *cloudkms.EkmConnectionServiceResolvers) *cloudkmspb.CloudkmsEkmConnectionServiceResolvers {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsEkmConnectionServiceResolvers{}
	p.SetServiceDirectoryService(dcl.ValueOrEmptyString(o.ServiceDirectoryService))
	p.SetEndpointFilter(dcl.ValueOrEmptyString(o.EndpointFilter))
	p.SetHostname(dcl.ValueOrEmptyString(o.Hostname))
	sServerCertificates := make([]*cloudkmspb.CloudkmsEkmConnectionServiceResolversServerCertificates, len(o.ServerCertificates))
	for i, r := range o.ServerCertificates {
		sServerCertificates[i] = CloudkmsEkmConnectionServiceResolversServerCertificatesToProto(&r)
	}
	p.SetServerCertificates(sServerCertificates)
	return p
}

// EkmConnectionServiceResolversServerCertificatesToProto converts a EkmConnectionServiceResolversServerCertificates object to its proto representation.
func CloudkmsEkmConnectionServiceResolversServerCertificatesToProto(o *cloudkms.EkmConnectionServiceResolversServerCertificates) *cloudkmspb.CloudkmsEkmConnectionServiceResolversServerCertificates {
	if o == nil {
		return nil
	}
	p := &cloudkmspb.CloudkmsEkmConnectionServiceResolversServerCertificates{}
	p.SetRawDer(dcl.ValueOrEmptyString(o.RawDer))
	p.SetParsed(dcl.ValueOrEmptyBool(o.Parsed))
	p.SetIssuer(dcl.ValueOrEmptyString(o.Issuer))
	p.SetSubject(dcl.ValueOrEmptyString(o.Subject))
	p.SetNotBeforeTime(dcl.ValueOrEmptyString(o.NotBeforeTime))
	p.SetNotAfterTime(dcl.ValueOrEmptyString(o.NotAfterTime))
	p.SetSerialNumber(dcl.ValueOrEmptyString(o.SerialNumber))
	p.SetSha256Fingerprint(dcl.ValueOrEmptyString(o.Sha256Fingerprint))
	sSubjectAlternativeDnsNames := make([]string, len(o.SubjectAlternativeDnsNames))
	for i, r := range o.SubjectAlternativeDnsNames {
		sSubjectAlternativeDnsNames[i] = r
	}
	p.SetSubjectAlternativeDnsNames(sSubjectAlternativeDnsNames)
	return p
}

// EkmConnectionToProto converts a EkmConnection resource to its proto representation.
func EkmConnectionToProto(resource *cloudkms.EkmConnection) *cloudkmspb.CloudkmsEkmConnection {
	p := &cloudkmspb.CloudkmsEkmConnection{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	sServiceResolvers := make([]*cloudkmspb.CloudkmsEkmConnectionServiceResolvers, len(resource.ServiceResolvers))
	for i, r := range resource.ServiceResolvers {
		sServiceResolvers[i] = CloudkmsEkmConnectionServiceResolversToProto(&r)
	}
	p.SetServiceResolvers(sServiceResolvers)

	return p
}

// applyEkmConnection handles the gRPC request by passing it to the underlying EkmConnection Apply() method.
func (s *EkmConnectionServer) applyEkmConnection(ctx context.Context, c *cloudkms.Client, request *cloudkmspb.ApplyCloudkmsEkmConnectionRequest) (*cloudkmspb.CloudkmsEkmConnection, error) {
	p := ProtoToEkmConnection(request.GetResource())
	res, err := c.ApplyEkmConnection(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EkmConnectionToProto(res)
	return r, nil
}

// applyCloudkmsEkmConnection handles the gRPC request by passing it to the underlying EkmConnection Apply() method.
func (s *EkmConnectionServer) ApplyCloudkmsEkmConnection(ctx context.Context, request *cloudkmspb.ApplyCloudkmsEkmConnectionRequest) (*cloudkmspb.CloudkmsEkmConnection, error) {
	cl, err := createConfigEkmConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEkmConnection(ctx, cl, request)
}

// DeleteEkmConnection handles the gRPC request by passing it to the underlying EkmConnection Delete() method.
func (s *EkmConnectionServer) DeleteCloudkmsEkmConnection(ctx context.Context, request *cloudkmspb.DeleteCloudkmsEkmConnectionRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for EkmConnection")

}

// ListCloudkmsEkmConnection handles the gRPC request by passing it to the underlying EkmConnectionList() method.
func (s *EkmConnectionServer) ListCloudkmsEkmConnection(ctx context.Context, request *cloudkmspb.ListCloudkmsEkmConnectionRequest) (*cloudkmspb.ListCloudkmsEkmConnectionResponse, error) {
	cl, err := createConfigEkmConnection(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEkmConnection(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*cloudkmspb.CloudkmsEkmConnection
	for _, r := range resources.Items {
		rp := EkmConnectionToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudkmspb.ListCloudkmsEkmConnectionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEkmConnection(ctx context.Context, service_account_file string) (*cloudkms.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudkms.NewClient(conf), nil
}
