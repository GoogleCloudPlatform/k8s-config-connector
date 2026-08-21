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
	dnspb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/dns/dns_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dns"
)

// Server implements the gRPC interface for ResourceRecordSet.
type ResourceRecordSetServer struct{}

// ProtoToResourceRecordSet converts a ResourceRecordSet resource from its proto representation.
func ProtoToResourceRecordSet(p *dnspb.DnsResourceRecordSet) *dns.ResourceRecordSet {
	obj := &dns.ResourceRecordSet{
		DnsName:     dcl.StringOrNil(p.DnsName),
		DnsType:     dcl.StringOrNil(p.DnsType),
		Ttl:         dcl.Int64OrNil(p.Ttl),
		ManagedZone: dcl.StringOrNil(p.ManagedZone),
		Project:     dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetTarget() {
		obj.Target = append(obj.Target, r)
	}
	return obj
}

// ResourceRecordSetToProto converts a ResourceRecordSet resource to its proto representation.
func ResourceRecordSetToProto(resource *dns.ResourceRecordSet) *dnspb.DnsResourceRecordSet {
	p := &dnspb.DnsResourceRecordSet{
		DnsName:     dcl.ValueOrEmptyString(resource.DnsName),
		DnsType:     dcl.ValueOrEmptyString(resource.DnsType),
		Ttl:         dcl.ValueOrEmptyInt64(resource.Ttl),
		ManagedZone: dcl.ValueOrEmptyString(resource.ManagedZone),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Target {
		p.Target = append(p.Target, r)
	}

	return p
}

// ApplyResourceRecordSet handles the gRPC request by passing it to the underlying ResourceRecordSet Apply() method.
func (s *ResourceRecordSetServer) applyResourceRecordSet(ctx context.Context, c *dns.Client, request *dnspb.ApplyDnsResourceRecordSetRequest) (*dnspb.DnsResourceRecordSet, error) {
	p := ProtoToResourceRecordSet(request.GetResource())
	res, err := c.ApplyResourceRecordSet(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ResourceRecordSetToProto(res)
	return r, nil
}

// ApplyResourceRecordSet handles the gRPC request by passing it to the underlying ResourceRecordSet Apply() method.
func (s *ResourceRecordSetServer) ApplyDnsResourceRecordSet(ctx context.Context, request *dnspb.ApplyDnsResourceRecordSetRequest) (*dnspb.DnsResourceRecordSet, error) {
	cl, err := createConfigResourceRecordSet(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyResourceRecordSet(ctx, cl, request)
}

// DeleteResourceRecordSet handles the gRPC request by passing it to the underlying ResourceRecordSet Delete() method.
func (s *ResourceRecordSetServer) DeleteDnsResourceRecordSet(ctx context.Context, request *dnspb.DeleteDnsResourceRecordSetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigResourceRecordSet(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteResourceRecordSet(ctx, ProtoToResourceRecordSet(request.GetResource()))

}

// ListDnsResourceRecordSet handles the gRPC request by passing it to the underlying ResourceRecordSetList() method.
func (s *ResourceRecordSetServer) ListDnsResourceRecordSet(ctx context.Context, request *dnspb.ListDnsResourceRecordSetRequest) (*dnspb.ListDnsResourceRecordSetResponse, error) {
	cl, err := createConfigResourceRecordSet(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListResourceRecordSet(ctx, request.Project, request.ManagedZone)
	if err != nil {
		return nil, err
	}
	var protos []*dnspb.DnsResourceRecordSet
	for _, r := range resources.Items {
		rp := ResourceRecordSetToProto(r)
		protos = append(protos, rp)
	}
	return &dnspb.ListDnsResourceRecordSetResponse{Items: protos}, nil
}

func createConfigResourceRecordSet(ctx context.Context, service_account_file string) (*dns.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return dns.NewClient(conf), nil
}
