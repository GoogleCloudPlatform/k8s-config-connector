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
	accesscontextmanagerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/accesscontextmanager/accesscontextmanager_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/accesscontextmanager"
)

// Server implements the gRPC interface for AccessPolicy.
type AccessPolicyServer struct{}

// ProtoToAccessPolicy converts a AccessPolicy resource from its proto representation.
func ProtoToAccessPolicy(p *accesscontextmanagerpb.AccesscontextmanagerAccessPolicy) *accesscontextmanager.AccessPolicy {
	obj := &accesscontextmanager.AccessPolicy{
		Name:       dcl.StringOrNil(p.Name),
		Parent:     dcl.StringOrNil(p.Parent),
		Title:      dcl.StringOrNil(p.Title),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// AccessPolicyToProto converts a AccessPolicy resource to its proto representation.
func AccessPolicyToProto(resource *accesscontextmanager.AccessPolicy) *accesscontextmanagerpb.AccesscontextmanagerAccessPolicy {
	p := &accesscontextmanagerpb.AccesscontextmanagerAccessPolicy{
		Name:       dcl.ValueOrEmptyString(resource.Name),
		Parent:     dcl.ValueOrEmptyString(resource.Parent),
		Title:      dcl.ValueOrEmptyString(resource.Title),
		CreateTime: dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime: dcl.ValueOrEmptyString(resource.UpdateTime),
	}

	return p
}

// ApplyAccessPolicy handles the gRPC request by passing it to the underlying AccessPolicy Apply() method.
func (s *AccessPolicyServer) applyAccessPolicy(ctx context.Context, c *accesscontextmanager.Client, request *accesscontextmanagerpb.ApplyAccesscontextmanagerAccessPolicyRequest) (*accesscontextmanagerpb.AccesscontextmanagerAccessPolicy, error) {
	p := ProtoToAccessPolicy(request.GetResource())
	res, err := c.ApplyAccessPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AccessPolicyToProto(res)
	return r, nil
}

// ApplyAccessPolicy handles the gRPC request by passing it to the underlying AccessPolicy Apply() method.
func (s *AccessPolicyServer) ApplyAccesscontextmanagerAccessPolicy(ctx context.Context, request *accesscontextmanagerpb.ApplyAccesscontextmanagerAccessPolicyRequest) (*accesscontextmanagerpb.AccesscontextmanagerAccessPolicy, error) {
	cl, err := createConfigAccessPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyAccessPolicy(ctx, cl, request)
}

// DeleteAccessPolicy handles the gRPC request by passing it to the underlying AccessPolicy Delete() method.
func (s *AccessPolicyServer) DeleteAccesscontextmanagerAccessPolicy(ctx context.Context, request *accesscontextmanagerpb.DeleteAccesscontextmanagerAccessPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAccessPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAccessPolicy(ctx, ProtoToAccessPolicy(request.GetResource()))

}

// ListAccesscontextmanagerAccessPolicy handles the gRPC request by passing it to the underlying AccessPolicyList() method.
func (s *AccessPolicyServer) ListAccesscontextmanagerAccessPolicy(ctx context.Context, request *accesscontextmanagerpb.ListAccesscontextmanagerAccessPolicyRequest) (*accesscontextmanagerpb.ListAccesscontextmanagerAccessPolicyResponse, error) {
	cl, err := createConfigAccessPolicy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAccessPolicy(ctx, request.Parent)
	if err != nil {
		return nil, err
	}
	var protos []*accesscontextmanagerpb.AccesscontextmanagerAccessPolicy
	for _, r := range resources.Items {
		rp := AccessPolicyToProto(r)
		protos = append(protos, rp)
	}
	return &accesscontextmanagerpb.ListAccesscontextmanagerAccessPolicyResponse{Items: protos}, nil
}

func createConfigAccessPolicy(ctx context.Context, service_account_file string) (*accesscontextmanager.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return accesscontextmanager.NewClient(conf), nil
}
