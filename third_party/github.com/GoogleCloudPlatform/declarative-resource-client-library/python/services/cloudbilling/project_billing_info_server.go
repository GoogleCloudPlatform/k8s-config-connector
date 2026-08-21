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
	cloudbillingpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbilling/cloudbilling_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbilling"
)

// Server implements the gRPC interface for ProjectBillingInfo.
type ProjectBillingInfoServer struct{}

// ProtoToProjectBillingInfo converts a ProjectBillingInfo resource from its proto representation.
func ProtoToProjectBillingInfo(p *cloudbillingpb.CloudbillingProjectBillingInfo) *cloudbilling.ProjectBillingInfo {
	obj := &cloudbilling.ProjectBillingInfo{
		Name:               dcl.StringOrNil(p.Name),
		BillingAccountName: dcl.StringOrNil(p.BillingAccountName),
		BillingEnabled:     dcl.StringOrNil(p.BillingEnabled),
	}
	return obj
}

// ProjectBillingInfoToProto converts a ProjectBillingInfo resource to its proto representation.
func ProjectBillingInfoToProto(resource *cloudbilling.ProjectBillingInfo) *cloudbillingpb.CloudbillingProjectBillingInfo {
	p := &cloudbillingpb.CloudbillingProjectBillingInfo{
		Name:               dcl.ValueOrEmptyString(resource.Name),
		BillingAccountName: dcl.ValueOrEmptyString(resource.BillingAccountName),
		BillingEnabled:     dcl.ValueOrEmptyString(resource.BillingEnabled),
	}

	return p
}

// ApplyProjectBillingInfo handles the gRPC request by passing it to the underlying ProjectBillingInfo Apply() method.
func (s *ProjectBillingInfoServer) applyProjectBillingInfo(ctx context.Context, c *cloudbilling.Client, request *cloudbillingpb.ApplyCloudbillingProjectBillingInfoRequest) (*cloudbillingpb.CloudbillingProjectBillingInfo, error) {
	p := ProtoToProjectBillingInfo(request.GetResource())
	res, err := c.ApplyProjectBillingInfo(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ProjectBillingInfoToProto(res)
	return r, nil
}

// ApplyProjectBillingInfo handles the gRPC request by passing it to the underlying ProjectBillingInfo Apply() method.
func (s *ProjectBillingInfoServer) ApplyCloudbillingProjectBillingInfo(ctx context.Context, request *cloudbillingpb.ApplyCloudbillingProjectBillingInfoRequest) (*cloudbillingpb.CloudbillingProjectBillingInfo, error) {
	cl, err := createConfigProjectBillingInfo(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyProjectBillingInfo(ctx, cl, request)
}

// DeleteProjectBillingInfo handles the gRPC request by passing it to the underlying ProjectBillingInfo Delete() method.
func (s *ProjectBillingInfoServer) DeleteCloudbillingProjectBillingInfo(ctx context.Context, request *cloudbillingpb.DeleteCloudbillingProjectBillingInfoRequest) (*emptypb.Empty, error) {

	cl, err := createConfigProjectBillingInfo(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteProjectBillingInfo(ctx, ProtoToProjectBillingInfo(request.GetResource()))

}

// ListCloudbillingProjectBillingInfo is a no-op method because ProjectBillingInfo has no list method.
func (s *ProjectBillingInfoServer) ListCloudbillingProjectBillingInfo(_ context.Context, _ *cloudbillingpb.ListCloudbillingProjectBillingInfoRequest) (*cloudbillingpb.ListCloudbillingProjectBillingInfoResponse, error) {
	return nil, nil
}

func createConfigProjectBillingInfo(ctx context.Context, service_account_file string) (*cloudbilling.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudbilling.NewClient(conf), nil
}
