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

// Server implements the gRPC interface for ManagedSslCertificate.
type ManagedSslCertificateServer struct{}

// ProtoToManagedSslCertificateManagedStatusEnum converts a ManagedSslCertificateManagedStatusEnum enum from its proto representation.
func ProtoToComputeBetaManagedSslCertificateManagedStatusEnum(e betapb.ComputeBetaManagedSslCertificateManagedStatusEnum) *beta.ManagedSslCertificateManagedStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaManagedSslCertificateManagedStatusEnum_name[int32(e)]; ok {
		e := beta.ManagedSslCertificateManagedStatusEnum(n[len("ComputeBetaManagedSslCertificateManagedStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedSslCertificateManagedDomainStatusEnum converts a ManagedSslCertificateManagedDomainStatusEnum enum from its proto representation.
func ProtoToComputeBetaManagedSslCertificateManagedDomainStatusEnum(e betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum) *beta.ManagedSslCertificateManagedDomainStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum_name[int32(e)]; ok {
		e := beta.ManagedSslCertificateManagedDomainStatusEnum(n[len("ComputeBetaManagedSslCertificateManagedDomainStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedSslCertificateTypeEnum converts a ManagedSslCertificateTypeEnum enum from its proto representation.
func ProtoToComputeBetaManagedSslCertificateTypeEnum(e betapb.ComputeBetaManagedSslCertificateTypeEnum) *beta.ManagedSslCertificateTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaManagedSslCertificateTypeEnum_name[int32(e)]; ok {
		e := beta.ManagedSslCertificateTypeEnum(n[len("ComputeBetaManagedSslCertificateTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedSslCertificateManaged converts a ManagedSslCertificateManaged resource from its proto representation.
func ProtoToComputeBetaManagedSslCertificateManaged(p *betapb.ComputeBetaManagedSslCertificateManaged) *beta.ManagedSslCertificateManaged {
	if p == nil {
		return nil
	}
	obj := &beta.ManagedSslCertificateManaged{
		Status: ProtoToComputeBetaManagedSslCertificateManagedStatusEnum(p.GetStatus()),
	}
	for _, r := range p.GetDomains() {
		obj.Domains = append(obj.Domains, r)
	}
	return obj
}

// ProtoToManagedSslCertificate converts a ManagedSslCertificate resource from its proto representation.
func ProtoToManagedSslCertificate(p *betapb.ComputeBetaManagedSslCertificate) *beta.ManagedSslCertificate {
	obj := &beta.ManagedSslCertificate{
		Name:              dcl.StringOrNil(p.Name),
		Id:                dcl.Int64OrNil(p.Id),
		CreationTimestamp: dcl.StringOrNil(p.CreationTimestamp),
		Description:       dcl.StringOrNil(p.Description),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Managed:           ProtoToComputeBetaManagedSslCertificateManaged(p.GetManaged()),
		Type:              ProtoToComputeBetaManagedSslCertificateTypeEnum(p.GetType()),
		ExpireTime:        dcl.StringOrNil(p.ExpireTime),
		Project:           dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetSubjectAlternativeNames() {
		obj.SubjectAlternativeNames = append(obj.SubjectAlternativeNames, r)
	}
	return obj
}

// ManagedSslCertificateManagedStatusEnumToProto converts a ManagedSslCertificateManagedStatusEnum enum to its proto representation.
func ComputeBetaManagedSslCertificateManagedStatusEnumToProto(e *beta.ManagedSslCertificateManagedStatusEnum) betapb.ComputeBetaManagedSslCertificateManagedStatusEnum {
	if e == nil {
		return betapb.ComputeBetaManagedSslCertificateManagedStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaManagedSslCertificateManagedStatusEnum_value["ManagedSslCertificateManagedStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaManagedSslCertificateManagedStatusEnum(v)
	}
	return betapb.ComputeBetaManagedSslCertificateManagedStatusEnum(0)
}

// ManagedSslCertificateManagedDomainStatusEnumToProto converts a ManagedSslCertificateManagedDomainStatusEnum enum to its proto representation.
func ComputeBetaManagedSslCertificateManagedDomainStatusEnumToProto(e *beta.ManagedSslCertificateManagedDomainStatusEnum) betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum {
	if e == nil {
		return betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum_value["ManagedSslCertificateManagedDomainStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum(v)
	}
	return betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum(0)
}

// ManagedSslCertificateTypeEnumToProto converts a ManagedSslCertificateTypeEnum enum to its proto representation.
func ComputeBetaManagedSslCertificateTypeEnumToProto(e *beta.ManagedSslCertificateTypeEnum) betapb.ComputeBetaManagedSslCertificateTypeEnum {
	if e == nil {
		return betapb.ComputeBetaManagedSslCertificateTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaManagedSslCertificateTypeEnum_value["ManagedSslCertificateTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaManagedSslCertificateTypeEnum(v)
	}
	return betapb.ComputeBetaManagedSslCertificateTypeEnum(0)
}

// ManagedSslCertificateManagedToProto converts a ManagedSslCertificateManaged resource to its proto representation.
func ComputeBetaManagedSslCertificateManagedToProto(o *beta.ManagedSslCertificateManaged) *betapb.ComputeBetaManagedSslCertificateManaged {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaManagedSslCertificateManaged{
		Status: ComputeBetaManagedSslCertificateManagedStatusEnumToProto(o.Status),
	}
	for _, r := range o.Domains {
		p.Domains = append(p.Domains, r)
	}
	p.DomainStatus = make(map[string]betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum)
	for k, r := range o.DomainStatus {
		p.DomainStatus[k] = betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum(betapb.ComputeBetaManagedSslCertificateManagedDomainStatusEnum_value[string(r)])
	}
	return p
}

// ManagedSslCertificateToProto converts a ManagedSslCertificate resource to its proto representation.
func ManagedSslCertificateToProto(resource *beta.ManagedSslCertificate) *betapb.ComputeBetaManagedSslCertificate {
	p := &betapb.ComputeBetaManagedSslCertificate{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Managed:           ComputeBetaManagedSslCertificateManagedToProto(resource.Managed),
		Type:              ComputeBetaManagedSslCertificateTypeEnumToProto(resource.Type),
		ExpireTime:        dcl.ValueOrEmptyString(resource.ExpireTime),
		Project:           dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.SubjectAlternativeNames {
		p.SubjectAlternativeNames = append(p.SubjectAlternativeNames, r)
	}

	return p
}

// ApplyManagedSslCertificate handles the gRPC request by passing it to the underlying ManagedSslCertificate Apply() method.
func (s *ManagedSslCertificateServer) applyManagedSslCertificate(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaManagedSslCertificateRequest) (*betapb.ComputeBetaManagedSslCertificate, error) {
	p := ProtoToManagedSslCertificate(request.GetResource())
	res, err := c.ApplyManagedSslCertificate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ManagedSslCertificateToProto(res)
	return r, nil
}

// ApplyManagedSslCertificate handles the gRPC request by passing it to the underlying ManagedSslCertificate Apply() method.
func (s *ManagedSslCertificateServer) ApplyComputeBetaManagedSslCertificate(ctx context.Context, request *betapb.ApplyComputeBetaManagedSslCertificateRequest) (*betapb.ComputeBetaManagedSslCertificate, error) {
	cl, err := createConfigManagedSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyManagedSslCertificate(ctx, cl, request)
}

// DeleteManagedSslCertificate handles the gRPC request by passing it to the underlying ManagedSslCertificate Delete() method.
func (s *ManagedSslCertificateServer) DeleteComputeBetaManagedSslCertificate(ctx context.Context, request *betapb.DeleteComputeBetaManagedSslCertificateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigManagedSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteManagedSslCertificate(ctx, ProtoToManagedSslCertificate(request.GetResource()))

}

// ListComputeBetaManagedSslCertificate handles the gRPC request by passing it to the underlying ManagedSslCertificateList() method.
func (s *ManagedSslCertificateServer) ListComputeBetaManagedSslCertificate(ctx context.Context, request *betapb.ListComputeBetaManagedSslCertificateRequest) (*betapb.ListComputeBetaManagedSslCertificateResponse, error) {
	cl, err := createConfigManagedSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListManagedSslCertificate(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaManagedSslCertificate
	for _, r := range resources.Items {
		rp := ManagedSslCertificateToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaManagedSslCertificateResponse{Items: protos}, nil
}

func createConfigManagedSslCertificate(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
