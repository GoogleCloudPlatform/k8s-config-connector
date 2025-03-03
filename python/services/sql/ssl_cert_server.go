// Copyright 2020 Google LLC. All Rights Reserved.
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
	sqlpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/sql/sql_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql"
)

// Server implements the gRPC interface for SslCert.
type SslCertServer struct{}

// ProtoToSslCert converts a SslCert resource from its proto representation.
func ProtoToSslCert(p *sqlpb.SqlSslCert) *sql.SslCert {
	obj := &sql.SslCert{
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
func SslCertToProto(resource *sql.SslCert) *sqlpb.SqlSslCert {
	p := &sqlpb.SqlSslCert{
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
func (s *SslCertServer) applySslCert(ctx context.Context, c *sql.Client, request *sqlpb.ApplySqlSslCertRequest) (*sqlpb.SqlSslCert, error) {
	p := ProtoToSslCert(request.GetResource())
	res, err := c.ApplySslCert(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SslCertToProto(res)
	return r, nil
}

// ApplySslCert handles the gRPC request by passing it to the underlying SslCert Apply() method.
func (s *SslCertServer) ApplySqlSslCert(ctx context.Context, request *sqlpb.ApplySqlSslCertRequest) (*sqlpb.SqlSslCert, error) {
	cl, err := createConfigSslCert(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySslCert(ctx, cl, request)
}

// DeleteSslCert handles the gRPC request by passing it to the underlying SslCert Delete() method.
func (s *SslCertServer) DeleteSqlSslCert(ctx context.Context, request *sqlpb.DeleteSqlSslCertRequest) (*emptypb.Empty, error) {
	cl, err := createConfigSslCert(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSslCert(ctx, ProtoToSslCert(request.GetResource()))
}

// ListSslCert handles the gRPC request by passing it to the underlying SslCertList() method.
func (s *SslCertServer) ListSqlSslCert(ctx context.Context, request *sqlpb.ListSqlSslCertRequest) (*sqlpb.ListSqlSslCertResponse, error) {
	cl, err := createConfigSslCert(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSslCert(ctx, request.Project, request.Instance)
	if err != nil {
		return nil, err
	}
	var protos []*sqlpb.SqlSslCert
	for _, r := range resources.Items {
		rp := SslCertToProto(r)
		protos = append(protos, rp)
	}
	return &sqlpb.ListSqlSslCertResponse{Items: protos}, nil
}

func createConfigSslCert(ctx context.Context, service_account_file string) (*sql.Client, error) {

	client, err := dcl.FromCredentialsFile(ctx, service_account_file)
	if err != nil {
		return nil, err
	}

	conf := dcl.NewConfig(client, dcl.WithUserAgent("dcl-test"))
	if err != nil {
		return nil, err
	}
	return sql.NewClient(conf), nil
}
