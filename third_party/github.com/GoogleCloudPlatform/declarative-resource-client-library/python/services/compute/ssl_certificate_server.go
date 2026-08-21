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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for SslCertificate.
type SslCertificateServer struct{}

// ProtoToSslCertificateTypeEnum converts a SslCertificateTypeEnum enum from its proto representation.
func ProtoToComputeSslCertificateTypeEnum(e computepb.ComputeSslCertificateTypeEnum) *compute.SslCertificateTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeSslCertificateTypeEnum_name[int32(e)]; ok {
		e := compute.SslCertificateTypeEnum(n[len("ComputeSslCertificateTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToSslCertificateSelfManaged converts a SslCertificateSelfManaged resource from its proto representation.
func ProtoToComputeSslCertificateSelfManaged(p *computepb.ComputeSslCertificateSelfManaged) *compute.SslCertificateSelfManaged {
	if p == nil {
		return nil
	}
	obj := &compute.SslCertificateSelfManaged{
		Certificate: dcl.StringOrNil(p.Certificate),
		PrivateKey:  dcl.StringOrNil(p.PrivateKey),
	}
	return obj
}

// ProtoToSslCertificate converts a SslCertificate resource from its proto representation.
func ProtoToSslCertificate(p *computepb.ComputeSslCertificate) *compute.SslCertificate {
	obj := &compute.SslCertificate{
		Name:              dcl.StringOrNil(p.Name),
		Id:                dcl.Int64OrNil(p.Id),
		CreationTimestamp: dcl.StringOrNil(p.CreationTimestamp),
		Description:       dcl.StringOrNil(p.Description),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		SelfManaged:       ProtoToComputeSslCertificateSelfManaged(p.GetSelfManaged()),
		Type:              ProtoToComputeSslCertificateTypeEnum(p.GetType()),
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
func ComputeSslCertificateTypeEnumToProto(e *compute.SslCertificateTypeEnum) computepb.ComputeSslCertificateTypeEnum {
	if e == nil {
		return computepb.ComputeSslCertificateTypeEnum(0)
	}
	if v, ok := computepb.ComputeSslCertificateTypeEnum_value["SslCertificateTypeEnum"+string(*e)]; ok {
		return computepb.ComputeSslCertificateTypeEnum(v)
	}
	return computepb.ComputeSslCertificateTypeEnum(0)
}

// SslCertificateSelfManagedToProto converts a SslCertificateSelfManaged resource to its proto representation.
func ComputeSslCertificateSelfManagedToProto(o *compute.SslCertificateSelfManaged) *computepb.ComputeSslCertificateSelfManaged {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSslCertificateSelfManaged{
		Certificate: dcl.ValueOrEmptyString(o.Certificate),
		PrivateKey:  dcl.ValueOrEmptyString(o.PrivateKey),
	}
	return p
}

// SslCertificateToProto converts a SslCertificate resource to its proto representation.
func SslCertificateToProto(resource *compute.SslCertificate) *computepb.ComputeSslCertificate {
	p := &computepb.ComputeSslCertificate{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		SelfManaged:       ComputeSslCertificateSelfManagedToProto(resource.SelfManaged),
		Type:              ComputeSslCertificateTypeEnumToProto(resource.Type),
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
func (s *SslCertificateServer) applySslCertificate(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeSslCertificateRequest) (*computepb.ComputeSslCertificate, error) {
	p := ProtoToSslCertificate(request.GetResource())
	res, err := c.ApplySslCertificate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SslCertificateToProto(res)
	return r, nil
}

// ApplySslCertificate handles the gRPC request by passing it to the underlying SslCertificate Apply() method.
func (s *SslCertificateServer) ApplyComputeSslCertificate(ctx context.Context, request *computepb.ApplyComputeSslCertificateRequest) (*computepb.ComputeSslCertificate, error) {
	cl, err := createConfigSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySslCertificate(ctx, cl, request)
}

// DeleteSslCertificate handles the gRPC request by passing it to the underlying SslCertificate Delete() method.
func (s *SslCertificateServer) DeleteComputeSslCertificate(ctx context.Context, request *computepb.DeleteComputeSslCertificateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSslCertificate(ctx, ProtoToSslCertificate(request.GetResource()))

}

// ListComputeSslCertificate handles the gRPC request by passing it to the underlying SslCertificateList() method.
func (s *SslCertificateServer) ListComputeSslCertificate(ctx context.Context, request *computepb.ListComputeSslCertificateRequest) (*computepb.ListComputeSslCertificateResponse, error) {
	cl, err := createConfigSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSslCertificate(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeSslCertificate
	for _, r := range resources.Items {
		rp := SslCertificateToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeSslCertificateResponse{Items: protos}, nil
}

func createConfigSslCertificate(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
