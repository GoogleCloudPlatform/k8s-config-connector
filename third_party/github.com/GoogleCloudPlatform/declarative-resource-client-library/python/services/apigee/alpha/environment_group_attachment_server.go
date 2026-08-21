// Copyright 2022 Google LLC. All Rights Reserved.
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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/alpha/apigee_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/alpha"
)

// EnvironmentGroupAttachmentServer implements the gRPC interface for EnvironmentGroupAttachment.
type EnvironmentGroupAttachmentServer struct{}

// ProtoToEnvironmentGroupAttachment converts a EnvironmentGroupAttachment resource from its proto representation.
func ProtoToEnvironmentGroupAttachment(p *alphapb.ApigeeAlphaEnvironmentGroupAttachment) *alpha.EnvironmentGroupAttachment {
	obj := &alpha.EnvironmentGroupAttachment{
		Name:        dcl.StringOrNil(p.GetName()),
		Environment: dcl.StringOrNil(p.GetEnvironment()),
		CreatedAt:   dcl.Int64OrNil(p.GetCreatedAt()),
		Envgroup:    dcl.StringOrNil(p.GetEnvgroup()),
	}
	return obj
}

// EnvironmentGroupAttachmentToProto converts a EnvironmentGroupAttachment resource to its proto representation.
func EnvironmentGroupAttachmentToProto(resource *alpha.EnvironmentGroupAttachment) *alphapb.ApigeeAlphaEnvironmentGroupAttachment {
	p := &alphapb.ApigeeAlphaEnvironmentGroupAttachment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetEnvironment(dcl.ValueOrEmptyString(resource.Environment))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetEnvgroup(dcl.ValueOrEmptyString(resource.Envgroup))

	return p
}

// applyEnvironmentGroupAttachment handles the gRPC request by passing it to the underlying EnvironmentGroupAttachment Apply() method.
func (s *EnvironmentGroupAttachmentServer) applyEnvironmentGroupAttachment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyApigeeAlphaEnvironmentGroupAttachmentRequest) (*alphapb.ApigeeAlphaEnvironmentGroupAttachment, error) {
	p := ProtoToEnvironmentGroupAttachment(request.GetResource())
	res, err := c.ApplyEnvironmentGroupAttachment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvironmentGroupAttachmentToProto(res)
	return r, nil
}

// applyApigeeAlphaEnvironmentGroupAttachment handles the gRPC request by passing it to the underlying EnvironmentGroupAttachment Apply() method.
func (s *EnvironmentGroupAttachmentServer) ApplyApigeeAlphaEnvironmentGroupAttachment(ctx context.Context, request *alphapb.ApplyApigeeAlphaEnvironmentGroupAttachmentRequest) (*alphapb.ApigeeAlphaEnvironmentGroupAttachment, error) {
	cl, err := createConfigEnvironmentGroupAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEnvironmentGroupAttachment(ctx, cl, request)
}

// DeleteEnvironmentGroupAttachment handles the gRPC request by passing it to the underlying EnvironmentGroupAttachment Delete() method.
func (s *EnvironmentGroupAttachmentServer) DeleteApigeeAlphaEnvironmentGroupAttachment(ctx context.Context, request *alphapb.DeleteApigeeAlphaEnvironmentGroupAttachmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvironmentGroupAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvironmentGroupAttachment(ctx, ProtoToEnvironmentGroupAttachment(request.GetResource()))

}

// ListApigeeAlphaEnvironmentGroupAttachment handles the gRPC request by passing it to the underlying EnvironmentGroupAttachmentList() method.
func (s *EnvironmentGroupAttachmentServer) ListApigeeAlphaEnvironmentGroupAttachment(ctx context.Context, request *alphapb.ListApigeeAlphaEnvironmentGroupAttachmentRequest) (*alphapb.ListApigeeAlphaEnvironmentGroupAttachmentResponse, error) {
	cl, err := createConfigEnvironmentGroupAttachment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvironmentGroupAttachment(ctx, request.GetEnvgroup())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ApigeeAlphaEnvironmentGroupAttachment
	for _, r := range resources.Items {
		rp := EnvironmentGroupAttachmentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListApigeeAlphaEnvironmentGroupAttachmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEnvironmentGroupAttachment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
