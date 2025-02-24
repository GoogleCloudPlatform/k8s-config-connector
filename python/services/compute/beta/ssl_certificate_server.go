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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for SslCertificate.
type SslCertificateServer struct{}

// ProtoToSslCertificateTypeEnum converts a SslCertificateTypeEnum enum from its proto representation.
func ProtoToComputeBetaSslCertificateTypeEnum(e betapb.ComputeBetaSslCertificateTypeEnum) *beta.SslCertificateTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaSslCertificateTypeEnum_name[int32(e)]; ok {
		e := beta.SslCertificateTypeEnum(n[len("ComputeBetaSslCertificateTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToSslCertificateSelfManaged converts a SslCertificateSelfManaged resource from its proto representation.
func ProtoToComputeBetaSslCertificateSelfManaged(p *betapb.ComputeBetaSslCertificateSelfManaged) *beta.SslCertificateSelfManaged {
	if p == nil {
		return nil
	}
	obj := &beta.SslCertificateSelfManaged{
		Certificate: dcl.StringOrNil(p.Certificate),
		PrivateKey:  dcl.StringOrNil(p.PrivateKey),
	}
	return obj
}

// ProtoToSslCertificate converts a SslCertificate resource from its proto representation.
func ProtoToSslCertificate(p *betapb.ComputeBetaSslCertificate) *beta.SslCertificate {
	obj := &beta.SslCertificate{
		Name:              dcl.StringOrNil(p.Name),
		Id:                dcl.Int64OrNil(p.Id),
		CreationTimestamp: dcl.StringOrNil(p.CreationTimestamp),
		Description:       dcl.StringOrNil(p.Description),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		SelfManaged:       ProtoToComputeBetaSslCertificateSelfManaged(p.GetSelfManaged()),
		Type:              ProtoToComputeBetaSslCertificateTypeEnum(p.GetType()),
		ExpireTime:        dcl.StringOrNil(p.ExpireTime),
		Region:            dcl.StringOrNil(p.Region),
		Project:           dcl.StringOrNil(p.Project),
		Location:          dcl.StringOrNil(p.Location),
	}
	for _, r := range p.GetSubjectAlternativeNames() {
		obj.SubjectAlternativeNames = append(obj.SubjectAlternativeNames, r)
	}
	return obj
}

// SslCertificateTypeEnumToProto converts a SslCertificateTypeEnum enum to its proto representation.
func ComputeBetaSslCertificateTypeEnumToProto(e *beta.SslCertificateTypeEnum) betapb.ComputeBetaSslCertificateTypeEnum {
	if e == nil {
		return betapb.ComputeBetaSslCertificateTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaSslCertificateTypeEnum_value["SslCertificateTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaSslCertificateTypeEnum(v)
	}
	return betapb.ComputeBetaSslCertificateTypeEnum(0)
}

// SslCertificateSelfManagedToProto converts a SslCertificateSelfManaged resource to its proto representation.
func ComputeBetaSslCertificateSelfManagedToProto(o *beta.SslCertificateSelfManaged) *betapb.ComputeBetaSslCertificateSelfManaged {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaSslCertificateSelfManaged{
		Certificate: dcl.ValueOrEmptyString(o.Certificate),
		PrivateKey:  dcl.ValueOrEmptyString(o.PrivateKey),
	}
	return p
}

// SslCertificateToProto converts a SslCertificate resource to its proto representation.
func SslCertificateToProto(resource *beta.SslCertificate) *betapb.ComputeBetaSslCertificate {
	p := &betapb.ComputeBetaSslCertificate{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		SelfManaged:       ComputeBetaSslCertificateSelfManagedToProto(resource.SelfManaged),
		Type:              ComputeBetaSslCertificateTypeEnumToProto(resource.Type),
		ExpireTime:        dcl.ValueOrEmptyString(resource.ExpireTime),
		Region:            dcl.ValueOrEmptyString(resource.Region),
		Project:           dcl.ValueOrEmptyString(resource.Project),
		Location:          dcl.ValueOrEmptyString(resource.Location),
	}
	for _, r := range resource.SubjectAlternativeNames {
		p.SubjectAlternativeNames = append(p.SubjectAlternativeNames, r)
	}

	return p
}

// ApplySslCertificate handles the gRPC request by passing it to the underlying SslCertificate Apply() method.
func (s *SslCertificateServer) applySslCertificate(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaSslCertificateRequest) (*betapb.ComputeBetaSslCertificate, error) {
	p := ProtoToSslCertificate(request.GetResource())
	res, err := c.ApplySslCertificate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SslCertificateToProto(res)
	return r, nil
}

// ApplySslCertificate handles the gRPC request by passing it to the underlying SslCertificate Apply() method.
func (s *SslCertificateServer) ApplyComputeBetaSslCertificate(ctx context.Context, request *betapb.ApplyComputeBetaSslCertificateRequest) (*betapb.ComputeBetaSslCertificate, error) {
	cl, err := createConfigSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySslCertificate(ctx, cl, request)
}

// DeleteSslCertificate handles the gRPC request by passing it to the underlying SslCertificate Delete() method.
func (s *SslCertificateServer) DeleteComputeBetaSslCertificate(ctx context.Context, request *betapb.DeleteComputeBetaSslCertificateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSslCertificate(ctx, ProtoToSslCertificate(request.GetResource()))

}

// ListComputeBetaSslCertificate handles the gRPC request by passing it to the underlying SslCertificateList() method.
func (s *SslCertificateServer) ListComputeBetaSslCertificate(ctx context.Context, request *betapb.ListComputeBetaSslCertificateRequest) (*betapb.ListComputeBetaSslCertificateResponse, error) {
	cl, err := createConfigSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSslCertificate(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaSslCertificate
	for _, r := range resources.Items {
		rp := SslCertificateToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaSslCertificateResponse{Items: protos}, nil
}

func createConfigSslCertificate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
