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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/firebase/beta/firebase_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebase/beta"
)

// AppleAppServer implements the gRPC interface for AppleApp.
type AppleAppServer struct{}

// ProtoToAppleApp converts a AppleApp resource from its proto representation.
func ProtoToAppleApp(p *betapb.FirebaseBetaAppleApp) *beta.AppleApp {
	obj := &beta.AppleApp{
		Name:        dcl.StringOrNil(p.GetName()),
		AppId:       dcl.StringOrNil(p.GetAppId()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		ProjectId:   dcl.StringOrNil(p.GetProjectId()),
		BundleId:    dcl.StringOrNil(p.GetBundleId()),
		AppStoreId:  dcl.StringOrNil(p.GetAppStoreId()),
		TeamId:      dcl.StringOrNil(p.GetTeamId()),
		ApiKeyId:    dcl.StringOrNil(p.GetApiKeyId()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// AppleAppToProto converts a AppleApp resource to its proto representation.
func AppleAppToProto(resource *beta.AppleApp) *betapb.FirebaseBetaAppleApp {
	p := &betapb.FirebaseBetaAppleApp{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAppId(dcl.ValueOrEmptyString(resource.AppId))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetProjectId(dcl.ValueOrEmptyString(resource.ProjectId))
	p.SetBundleId(dcl.ValueOrEmptyString(resource.BundleId))
	p.SetAppStoreId(dcl.ValueOrEmptyString(resource.AppStoreId))
	p.SetTeamId(dcl.ValueOrEmptyString(resource.TeamId))
	p.SetApiKeyId(dcl.ValueOrEmptyString(resource.ApiKeyId))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyAppleApp handles the gRPC request by passing it to the underlying AppleApp Apply() method.
func (s *AppleAppServer) applyAppleApp(ctx context.Context, c *beta.Client, request *betapb.ApplyFirebaseBetaAppleAppRequest) (*betapb.FirebaseBetaAppleApp, error) {
	p := ProtoToAppleApp(request.GetResource())
	res, err := c.ApplyAppleApp(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AppleAppToProto(res)
	return r, nil
}

// applyFirebaseBetaAppleApp handles the gRPC request by passing it to the underlying AppleApp Apply() method.
func (s *AppleAppServer) ApplyFirebaseBetaAppleApp(ctx context.Context, request *betapb.ApplyFirebaseBetaAppleAppRequest) (*betapb.FirebaseBetaAppleApp, error) {
	cl, err := createConfigAppleApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAppleApp(ctx, cl, request)
}

// DeleteAppleApp handles the gRPC request by passing it to the underlying AppleApp Delete() method.
func (s *AppleAppServer) DeleteFirebaseBetaAppleApp(ctx context.Context, request *betapb.DeleteFirebaseBetaAppleAppRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAppleApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAppleApp(ctx, ProtoToAppleApp(request.GetResource()))

}

// ListFirebaseBetaAppleApp handles the gRPC request by passing it to the underlying AppleAppList() method.
func (s *AppleAppServer) ListFirebaseBetaAppleApp(ctx context.Context, request *betapb.ListFirebaseBetaAppleAppRequest) (*betapb.ListFirebaseBetaAppleAppResponse, error) {
	cl, err := createConfigAppleApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAppleApp(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.FirebaseBetaAppleApp
	for _, r := range resources.Items {
		rp := AppleAppToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListFirebaseBetaAppleAppResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAppleApp(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
