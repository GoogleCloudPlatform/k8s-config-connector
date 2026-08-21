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

// Server implements the gRPC interface for ManagedSslCertificate.
type ManagedSslCertificateServer struct{}

// ProtoToManagedSslCertificateManagedStatusEnum converts a ManagedSslCertificateManagedStatusEnum enum from its proto representation.
func ProtoToComputeManagedSslCertificateManagedStatusEnum(e computepb.ComputeManagedSslCertificateManagedStatusEnum) *compute.ManagedSslCertificateManagedStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeManagedSslCertificateManagedStatusEnum_name[int32(e)]; ok {
		e := compute.ManagedSslCertificateManagedStatusEnum(n[len("ComputeManagedSslCertificateManagedStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedSslCertificateManagedDomainStatusEnum converts a ManagedSslCertificateManagedDomainStatusEnum enum from its proto representation.
func ProtoToComputeManagedSslCertificateManagedDomainStatusEnum(e computepb.ComputeManagedSslCertificateManagedDomainStatusEnum) *compute.ManagedSslCertificateManagedDomainStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeManagedSslCertificateManagedDomainStatusEnum_name[int32(e)]; ok {
		e := compute.ManagedSslCertificateManagedDomainStatusEnum(n[len("ComputeManagedSslCertificateManagedDomainStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedSslCertificateTypeEnum converts a ManagedSslCertificateTypeEnum enum from its proto representation.
func ProtoToComputeManagedSslCertificateTypeEnum(e computepb.ComputeManagedSslCertificateTypeEnum) *compute.ManagedSslCertificateTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeManagedSslCertificateTypeEnum_name[int32(e)]; ok {
		e := compute.ManagedSslCertificateTypeEnum(n[len("ComputeManagedSslCertificateTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToManagedSslCertificateManaged converts a ManagedSslCertificateManaged resource from its proto representation.
func ProtoToComputeManagedSslCertificateManaged(p *computepb.ComputeManagedSslCertificateManaged) *compute.ManagedSslCertificateManaged {
	if p == nil {
		return nil
	}
	obj := &compute.ManagedSslCertificateManaged{
		Status: ProtoToComputeManagedSslCertificateManagedStatusEnum(p.GetStatus()),
	}
	for _, r := range p.GetDomains() {
		obj.Domains = append(obj.Domains, r)
	}
	return obj
}

// ProtoToManagedSslCertificate converts a ManagedSslCertificate resource from its proto representation.
func ProtoToManagedSslCertificate(p *computepb.ComputeManagedSslCertificate) *compute.ManagedSslCertificate {
	obj := &compute.ManagedSslCertificate{
		Name:              dcl.StringOrNil(p.Name),
		Id:                dcl.Int64OrNil(p.Id),
		CreationTimestamp: dcl.StringOrNil(p.CreationTimestamp),
		Description:       dcl.StringOrNil(p.Description),
		SelfLink:          dcl.StringOrNil(p.SelfLink),
		Managed:           ProtoToComputeManagedSslCertificateManaged(p.GetManaged()),
		Type:              ProtoToComputeManagedSslCertificateTypeEnum(p.GetType()),
		ExpireTime:        dcl.StringOrNil(p.ExpireTime),
		Project:           dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetSubjectAlternativeNames() {
		obj.SubjectAlternativeNames = append(obj.SubjectAlternativeNames, r)
	}
	return obj
}

// ManagedSslCertificateManagedStatusEnumToProto converts a ManagedSslCertificateManagedStatusEnum enum to its proto representation.
func ComputeManagedSslCertificateManagedStatusEnumToProto(e *compute.ManagedSslCertificateManagedStatusEnum) computepb.ComputeManagedSslCertificateManagedStatusEnum {
	if e == nil {
		return computepb.ComputeManagedSslCertificateManagedStatusEnum(0)
	}
	if v, ok := computepb.ComputeManagedSslCertificateManagedStatusEnum_value["ManagedSslCertificateManagedStatusEnum"+string(*e)]; ok {
		return computepb.ComputeManagedSslCertificateManagedStatusEnum(v)
	}
	return computepb.ComputeManagedSslCertificateManagedStatusEnum(0)
}

// ManagedSslCertificateManagedDomainStatusEnumToProto converts a ManagedSslCertificateManagedDomainStatusEnum enum to its proto representation.
func ComputeManagedSslCertificateManagedDomainStatusEnumToProto(e *compute.ManagedSslCertificateManagedDomainStatusEnum) computepb.ComputeManagedSslCertificateManagedDomainStatusEnum {
	if e == nil {
		return computepb.ComputeManagedSslCertificateManagedDomainStatusEnum(0)
	}
	if v, ok := computepb.ComputeManagedSslCertificateManagedDomainStatusEnum_value["ManagedSslCertificateManagedDomainStatusEnum"+string(*e)]; ok {
		return computepb.ComputeManagedSslCertificateManagedDomainStatusEnum(v)
	}
	return computepb.ComputeManagedSslCertificateManagedDomainStatusEnum(0)
}

// ManagedSslCertificateTypeEnumToProto converts a ManagedSslCertificateTypeEnum enum to its proto representation.
func ComputeManagedSslCertificateTypeEnumToProto(e *compute.ManagedSslCertificateTypeEnum) computepb.ComputeManagedSslCertificateTypeEnum {
	if e == nil {
		return computepb.ComputeManagedSslCertificateTypeEnum(0)
	}
	if v, ok := computepb.ComputeManagedSslCertificateTypeEnum_value["ManagedSslCertificateTypeEnum"+string(*e)]; ok {
		return computepb.ComputeManagedSslCertificateTypeEnum(v)
	}
	return computepb.ComputeManagedSslCertificateTypeEnum(0)
}

// ManagedSslCertificateManagedToProto converts a ManagedSslCertificateManaged resource to its proto representation.
func ComputeManagedSslCertificateManagedToProto(o *compute.ManagedSslCertificateManaged) *computepb.ComputeManagedSslCertificateManaged {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeManagedSslCertificateManaged{
		Status: ComputeManagedSslCertificateManagedStatusEnumToProto(o.Status),
	}
	for _, r := range o.Domains {
		p.Domains = append(p.Domains, r)
	}
	p.DomainStatus = make(map[string]computepb.ComputeManagedSslCertificateManagedDomainStatusEnum)
	for k, r := range o.DomainStatus {
		p.DomainStatus[k] = computepb.ComputeManagedSslCertificateManagedDomainStatusEnum(computepb.ComputeManagedSslCertificateManagedDomainStatusEnum_value[string(r)])
	}
	return p
}

// ManagedSslCertificateToProto converts a ManagedSslCertificate resource to its proto representation.
func ManagedSslCertificateToProto(resource *compute.ManagedSslCertificate) *computepb.ComputeManagedSslCertificate {
	p := &computepb.ComputeManagedSslCertificate{
		Name:              dcl.ValueOrEmptyString(resource.Name),
		Id:                dcl.ValueOrEmptyInt64(resource.Id),
		CreationTimestamp: dcl.ValueOrEmptyString(resource.CreationTimestamp),
		Description:       dcl.ValueOrEmptyString(resource.Description),
		SelfLink:          dcl.ValueOrEmptyString(resource.SelfLink),
		Managed:           ComputeManagedSslCertificateManagedToProto(resource.Managed),
		Type:              ComputeManagedSslCertificateTypeEnumToProto(resource.Type),
		ExpireTime:        dcl.ValueOrEmptyString(resource.ExpireTime),
		Project:           dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.SubjectAlternativeNames {
		p.SubjectAlternativeNames = append(p.SubjectAlternativeNames, r)
	}

	return p
}

// ApplyManagedSslCertificate handles the gRPC request by passing it to the underlying ManagedSslCertificate Apply() method.
func (s *ManagedSslCertificateServer) applyManagedSslCertificate(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeManagedSslCertificateRequest) (*computepb.ComputeManagedSslCertificate, error) {
	p := ProtoToManagedSslCertificate(request.GetResource())
	res, err := c.ApplyManagedSslCertificate(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ManagedSslCertificateToProto(res)
	return r, nil
}

// ApplyManagedSslCertificate handles the gRPC request by passing it to the underlying ManagedSslCertificate Apply() method.
func (s *ManagedSslCertificateServer) ApplyComputeManagedSslCertificate(ctx context.Context, request *computepb.ApplyComputeManagedSslCertificateRequest) (*computepb.ComputeManagedSslCertificate, error) {
	cl, err := createConfigManagedSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyManagedSslCertificate(ctx, cl, request)
}

// DeleteManagedSslCertificate handles the gRPC request by passing it to the underlying ManagedSslCertificate Delete() method.
func (s *ManagedSslCertificateServer) DeleteComputeManagedSslCertificate(ctx context.Context, request *computepb.DeleteComputeManagedSslCertificateRequest) (*emptypb.Empty, error) {

	cl, err := createConfigManagedSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteManagedSslCertificate(ctx, ProtoToManagedSslCertificate(request.GetResource()))

}

// ListComputeManagedSslCertificate handles the gRPC request by passing it to the underlying ManagedSslCertificateList() method.
func (s *ManagedSslCertificateServer) ListComputeManagedSslCertificate(ctx context.Context, request *computepb.ListComputeManagedSslCertificateRequest) (*computepb.ListComputeManagedSslCertificateResponse, error) {
	cl, err := createConfigManagedSslCertificate(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListManagedSslCertificate(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeManagedSslCertificate
	for _, r := range resources.Items {
		rp := ManagedSslCertificateToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeManagedSslCertificateResponse{Items: protos}, nil
}

func createConfigManagedSslCertificate(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
