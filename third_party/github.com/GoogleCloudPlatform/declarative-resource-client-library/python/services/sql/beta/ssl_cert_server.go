// Copyright 2021 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/sql/beta/sql_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql/beta"
)

// Server implements the gRPC interface for SslCert.
type SslCertServer struct{}

// ProtoToSslCert converts a SslCert resource from its proto representation.
func ProtoToSslCert(p *betapb.SqlBetaSslCert) *beta.SslCert {
	obj := &beta.SslCert{
		CertSerialNumber: dcl.StringOrNil(p.CertSerialNumber),
		Cert:             dcl.StringOrNil(p.Cert),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		CommonName:       dcl.StringOrNil(p.CommonName),
		ExpirationTime:   dcl.StringOrNil(p.GetExpirationTime()),
		Name:             dcl.StringOrNil(p.Name),
		Instance:         dcl.StringOrNil(p.Instance),
		Project:          dcl.StringOrNil(p.Project),
	}
	return obj
}

// SslCertToProto converts a SslCert resource to its proto representation.
func SslCertToProto(resource *beta.SslCert) *betapb.SqlBetaSslCert {
	p := &betapb.SqlBetaSslCert{
		CertSerialNumber: dcl.ValueOrEmptyString(resource.CertSerialNumber),
		Cert:             dcl.ValueOrEmptyString(resource.Cert),
		CreateTime:       dcl.ValueOrEmptyString(resource.CreateTime),
		CommonName:       dcl.ValueOrEmptyString(resource.CommonName),
		ExpirationTime:   dcl.ValueOrEmptyString(resource.ExpirationTime),
		Name:             dcl.ValueOrEmptyString(resource.Name),
		Instance:         dcl.ValueOrEmptyString(resource.Instance),
		Project:          dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplySslCert handles the gRPC request by passing it to the underlying SslCert Apply() method.
func (s *SslCertServer) applySslCert(ctx context.Context, c *beta.Client, request *betapb.ApplySqlBetaSslCertRequest) (*betapb.SqlBetaSslCert, error) {
	p := ProtoToSslCert(request.GetResource())
	res, err := c.ApplySslCert(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SslCertToProto(res)
	return r, nil
}

// ApplySslCert handles the gRPC request by passing it to the underlying SslCert Apply() method.
func (s *SslCertServer) ApplySqlBetaSslCert(ctx context.Context, request *betapb.ApplySqlBetaSslCertRequest) (*betapb.SqlBetaSslCert, error) {
	cl, err := createConfigSslCert(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySslCert(ctx, cl, request)
}

// DeleteSslCert handles the gRPC request by passing it to the underlying SslCert Delete() method.
func (s *SslCertServer) DeleteSqlBetaSslCert(ctx context.Context, request *betapb.DeleteSqlBetaSslCertRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSslCert(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSslCert(ctx, ProtoToSslCert(request.GetResource()))

}

// ListSqlBetaSslCert handles the gRPC request by passing it to the underlying SslCertList() method.
func (s *SslCertServer) ListSqlBetaSslCert(ctx context.Context, request *betapb.ListSqlBetaSslCertRequest) (*betapb.ListSqlBetaSslCertResponse, error) {
	cl, err := createConfigSslCert(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSslCert(ctx, request.Project, request.Instance)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.SqlBetaSslCert
	for _, r := range resources.Items {
		rp := SslCertToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListSqlBetaSslCertResponse{Items: protos}, nil
}

func createConfigSslCert(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
