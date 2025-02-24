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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudkms/beta/cloudkms_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudkms/beta"
)

// KeyRingServer implements the gRPC interface for KeyRing.
type KeyRingServer struct{}

// ProtoToKeyRing converts a KeyRing resource from its proto representation.
func ProtoToKeyRing(p *betapb.CloudkmsBetaKeyRing) *beta.KeyRing {
	obj := &beta.KeyRing{
		Name:       dcl.StringOrNil(p.GetName()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// KeyRingToProto converts a KeyRing resource to its proto representation.
func KeyRingToProto(resource *beta.KeyRing) *betapb.CloudkmsBetaKeyRing {
	p := &betapb.CloudkmsBetaKeyRing{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyKeyRing handles the gRPC request by passing it to the underlying KeyRing Apply() method.
func (s *KeyRingServer) applyKeyRing(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudkmsBetaKeyRingRequest) (*betapb.CloudkmsBetaKeyRing, error) {
	p := ProtoToKeyRing(request.GetResource())
	res, err := c.ApplyKeyRing(ctx, p)
	if err != nil {
		return nil, err
	}
	r := KeyRingToProto(res)
	return r, nil
}

// applyCloudkmsBetaKeyRing handles the gRPC request by passing it to the underlying KeyRing Apply() method.
func (s *KeyRingServer) ApplyCloudkmsBetaKeyRing(ctx context.Context, request *betapb.ApplyCloudkmsBetaKeyRingRequest) (*betapb.CloudkmsBetaKeyRing, error) {
	cl, err := createConfigKeyRing(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyKeyRing(ctx, cl, request)
}

// DeleteKeyRing handles the gRPC request by passing it to the underlying KeyRing Delete() method.
func (s *KeyRingServer) DeleteCloudkmsBetaKeyRing(ctx context.Context, request *betapb.DeleteCloudkmsBetaKeyRingRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for KeyRing")

}

// ListCloudkmsBetaKeyRing handles the gRPC request by passing it to the underlying KeyRingList() method.
func (s *KeyRingServer) ListCloudkmsBetaKeyRing(ctx context.Context, request *betapb.ListCloudkmsBetaKeyRingRequest) (*betapb.ListCloudkmsBetaKeyRingResponse, error) {
	cl, err := createConfigKeyRing(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListKeyRing(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudkmsBetaKeyRing
	for _, r := range resources.Items {
		rp := KeyRingToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudkmsBetaKeyRingResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigKeyRing(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
