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
	appenginepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/appengine/appengine_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/appengine"
)

// Server implements the gRPC interface for DomainMapping.
type DomainMappingServer struct{}

// ProtoToDomainMappingSslSettingsSslManagementTypeEnum converts a DomainMappingSslSettingsSslManagementTypeEnum enum from its proto representation.
func ProtoToAppengineDomainMappingSslSettingsSslManagementTypeEnum(e appenginepb.AppengineDomainMappingSslSettingsSslManagementTypeEnum) *appengine.DomainMappingSslSettingsSslManagementTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineDomainMappingSslSettingsSslManagementTypeEnum_name[int32(e)]; ok {
		e := appengine.DomainMappingSslSettingsSslManagementTypeEnum(n[len("AppengineDomainMappingSslSettingsSslManagementTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDomainMappingResourceRecordsTypeEnum converts a DomainMappingResourceRecordsTypeEnum enum from its proto representation.
func ProtoToAppengineDomainMappingResourceRecordsTypeEnum(e appenginepb.AppengineDomainMappingResourceRecordsTypeEnum) *appengine.DomainMappingResourceRecordsTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := appenginepb.AppengineDomainMappingResourceRecordsTypeEnum_name[int32(e)]; ok {
		e := appengine.DomainMappingResourceRecordsTypeEnum(n[len("AppengineDomainMappingResourceRecordsTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToDomainMappingSslSettings converts a DomainMappingSslSettings resource from its proto representation.
func ProtoToAppengineDomainMappingSslSettings(p *appenginepb.AppengineDomainMappingSslSettings) *appengine.DomainMappingSslSettings {
	if p == nil {
		return nil
	}
	obj := &appengine.DomainMappingSslSettings{
		CertificateId:               dcl.StringOrNil(p.CertificateId),
		SslManagementType:           ProtoToAppengineDomainMappingSslSettingsSslManagementTypeEnum(p.GetSslManagementType()),
		PendingManagedCertificateId: dcl.StringOrNil(p.PendingManagedCertificateId),
	}
	return obj
}

// ProtoToDomainMappingResourceRecords converts a DomainMappingResourceRecords resource from its proto representation.
func ProtoToAppengineDomainMappingResourceRecords(p *appenginepb.AppengineDomainMappingResourceRecords) *appengine.DomainMappingResourceRecords {
	if p == nil {
		return nil
	}
	obj := &appengine.DomainMappingResourceRecords{
		Name:   dcl.StringOrNil(p.Name),
		Rrdata: dcl.StringOrNil(p.Rrdata),
		Type:   ProtoToAppengineDomainMappingResourceRecordsTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToDomainMapping converts a DomainMapping resource from its proto representation.
func ProtoToDomainMapping(p *appenginepb.AppengineDomainMapping) *appengine.DomainMapping {
	obj := &appengine.DomainMapping{
		SelfLink:    dcl.StringOrNil(p.SelfLink),
		Name:        dcl.StringOrNil(p.Name),
		SslSettings: ProtoToAppengineDomainMappingSslSettings(p.GetSslSettings()),
		App:         dcl.StringOrNil(p.App),
	}
	for _, r := range p.GetResourceRecords() {
		obj.ResourceRecords = append(obj.ResourceRecords, *ProtoToAppengineDomainMappingResourceRecords(r))
	}
	return obj
}

// DomainMappingSslSettingsSslManagementTypeEnumToProto converts a DomainMappingSslSettingsSslManagementTypeEnum enum to its proto representation.
func AppengineDomainMappingSslSettingsSslManagementTypeEnumToProto(e *appengine.DomainMappingSslSettingsSslManagementTypeEnum) appenginepb.AppengineDomainMappingSslSettingsSslManagementTypeEnum {
	if e == nil {
		return appenginepb.AppengineDomainMappingSslSettingsSslManagementTypeEnum(0)
	}
	if v, ok := appenginepb.AppengineDomainMappingSslSettingsSslManagementTypeEnum_value["DomainMappingSslSettingsSslManagementTypeEnum"+string(*e)]; ok {
		return appenginepb.AppengineDomainMappingSslSettingsSslManagementTypeEnum(v)
	}
	return appenginepb.AppengineDomainMappingSslSettingsSslManagementTypeEnum(0)
}

// DomainMappingResourceRecordsTypeEnumToProto converts a DomainMappingResourceRecordsTypeEnum enum to its proto representation.
func AppengineDomainMappingResourceRecordsTypeEnumToProto(e *appengine.DomainMappingResourceRecordsTypeEnum) appenginepb.AppengineDomainMappingResourceRecordsTypeEnum {
	if e == nil {
		return appenginepb.AppengineDomainMappingResourceRecordsTypeEnum(0)
	}
	if v, ok := appenginepb.AppengineDomainMappingResourceRecordsTypeEnum_value["DomainMappingResourceRecordsTypeEnum"+string(*e)]; ok {
		return appenginepb.AppengineDomainMappingResourceRecordsTypeEnum(v)
	}
	return appenginepb.AppengineDomainMappingResourceRecordsTypeEnum(0)
}

// DomainMappingSslSettingsToProto converts a DomainMappingSslSettings resource to its proto representation.
func AppengineDomainMappingSslSettingsToProto(o *appengine.DomainMappingSslSettings) *appenginepb.AppengineDomainMappingSslSettings {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineDomainMappingSslSettings{
		CertificateId:               dcl.ValueOrEmptyString(o.CertificateId),
		SslManagementType:           AppengineDomainMappingSslSettingsSslManagementTypeEnumToProto(o.SslManagementType),
		PendingManagedCertificateId: dcl.ValueOrEmptyString(o.PendingManagedCertificateId),
	}
	return p
}

// DomainMappingResourceRecordsToProto converts a DomainMappingResourceRecords resource to its proto representation.
func AppengineDomainMappingResourceRecordsToProto(o *appengine.DomainMappingResourceRecords) *appenginepb.AppengineDomainMappingResourceRecords {
	if o == nil {
		return nil
	}
	p := &appenginepb.AppengineDomainMappingResourceRecords{
		Name:   dcl.ValueOrEmptyString(o.Name),
		Rrdata: dcl.ValueOrEmptyString(o.Rrdata),
		Type:   AppengineDomainMappingResourceRecordsTypeEnumToProto(o.Type),
	}
	return p
}

// DomainMappingToProto converts a DomainMapping resource to its proto representation.
func DomainMappingToProto(resource *appengine.DomainMapping) *appenginepb.AppengineDomainMapping {
	p := &appenginepb.AppengineDomainMapping{
		SelfLink:    dcl.ValueOrEmptyString(resource.SelfLink),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		SslSettings: AppengineDomainMappingSslSettingsToProto(resource.SslSettings),
		App:         dcl.ValueOrEmptyString(resource.App),
	}
	for _, r := range resource.ResourceRecords {
		p.ResourceRecords = append(p.ResourceRecords, AppengineDomainMappingResourceRecordsToProto(&r))
	}

	return p
}

// ApplyDomainMapping handles the gRPC request by passing it to the underlying DomainMapping Apply() method.
func (s *DomainMappingServer) applyDomainMapping(ctx context.Context, c *appengine.Client, request *appenginepb.ApplyAppengineDomainMappingRequest) (*appenginepb.AppengineDomainMapping, error) {
	p := ProtoToDomainMapping(request.GetResource())
	res, err := c.ApplyDomainMapping(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DomainMappingToProto(res)
	return r, nil
}

// ApplyDomainMapping handles the gRPC request by passing it to the underlying DomainMapping Apply() method.
func (s *DomainMappingServer) ApplyAppengineDomainMapping(ctx context.Context, request *appenginepb.ApplyAppengineDomainMappingRequest) (*appenginepb.AppengineDomainMapping, error) {
	cl, err := createConfigDomainMapping(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDomainMapping(ctx, cl, request)
}

// DeleteDomainMapping handles the gRPC request by passing it to the underlying DomainMapping Delete() method.
func (s *DomainMappingServer) DeleteAppengineDomainMapping(ctx context.Context, request *appenginepb.DeleteAppengineDomainMappingRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDomainMapping(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDomainMapping(ctx, ProtoToDomainMapping(request.GetResource()))

}

// ListAppengineDomainMapping handles the gRPC request by passing it to the underlying DomainMappingList() method.
func (s *DomainMappingServer) ListAppengineDomainMapping(ctx context.Context, request *appenginepb.ListAppengineDomainMappingRequest) (*appenginepb.ListAppengineDomainMappingResponse, error) {
	cl, err := createConfigDomainMapping(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDomainMapping(ctx, request.App)
	if err != nil {
		return nil, err
	}
	var protos []*appenginepb.AppengineDomainMapping
	for _, r := range resources.Items {
		rp := DomainMappingToProto(r)
		protos = append(protos, rp)
	}
	return &appenginepb.ListAppengineDomainMappingResponse{Items: protos}, nil
}

func createConfigDomainMapping(ctx context.Context, service_account_file string) (*appengine.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return appengine.NewClient(conf), nil
}
