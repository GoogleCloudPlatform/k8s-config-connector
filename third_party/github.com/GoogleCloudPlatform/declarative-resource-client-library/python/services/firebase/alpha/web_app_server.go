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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/firebase/alpha/firebase_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebase/alpha"
)

// WebAppServer implements the gRPC interface for WebApp.
type WebAppServer struct{}

// ProtoToWebApp converts a WebApp resource from its proto representation.
func ProtoToWebApp(p *alphapb.FirebaseAlphaWebApp) *alpha.WebApp {
	obj := &alpha.WebApp{
		Name:        dcl.StringOrNil(p.GetName()),
		AppId:       dcl.StringOrNil(p.GetAppId()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		ProjectId:   dcl.StringOrNil(p.GetProjectId()),
		WebId:       dcl.StringOrNil(p.GetWebId()),
		ApiKeyId:    dcl.StringOrNil(p.GetApiKeyId()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	for _, r := range p.GetAppUrls() {
		obj.AppUrls = append(obj.AppUrls, r)
	}
	return obj
}

// WebAppToProto converts a WebApp resource to its proto representation.
func WebAppToProto(resource *alpha.WebApp) *alphapb.FirebaseAlphaWebApp {
	p := &alphapb.FirebaseAlphaWebApp{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAppId(dcl.ValueOrEmptyString(resource.AppId))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetProjectId(dcl.ValueOrEmptyString(resource.ProjectId))
	p.SetWebId(dcl.ValueOrEmptyString(resource.WebId))
	p.SetApiKeyId(dcl.ValueOrEmptyString(resource.ApiKeyId))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	sAppUrls := make([]string, len(resource.AppUrls))
	for i, r := range resource.AppUrls {
		sAppUrls[i] = r
	}
	p.SetAppUrls(sAppUrls)

	return p
}

// applyWebApp handles the gRPC request by passing it to the underlying WebApp Apply() method.
func (s *WebAppServer) applyWebApp(ctx context.Context, c *alpha.Client, request *alphapb.ApplyFirebaseAlphaWebAppRequest) (*alphapb.FirebaseAlphaWebApp, error) {
	p := ProtoToWebApp(request.GetResource())
	res, err := c.ApplyWebApp(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WebAppToProto(res)
	return r, nil
}

// applyFirebaseAlphaWebApp handles the gRPC request by passing it to the underlying WebApp Apply() method.
func (s *WebAppServer) ApplyFirebaseAlphaWebApp(ctx context.Context, request *alphapb.ApplyFirebaseAlphaWebAppRequest) (*alphapb.FirebaseAlphaWebApp, error) {
	cl, err := createConfigWebApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWebApp(ctx, cl, request)
}

// DeleteWebApp handles the gRPC request by passing it to the underlying WebApp Delete() method.
func (s *WebAppServer) DeleteFirebaseAlphaWebApp(ctx context.Context, request *alphapb.DeleteFirebaseAlphaWebAppRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWebApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWebApp(ctx, ProtoToWebApp(request.GetResource()))

}

// ListFirebaseAlphaWebApp handles the gRPC request by passing it to the underlying WebAppList() method.
func (s *WebAppServer) ListFirebaseAlphaWebApp(ctx context.Context, request *alphapb.ListFirebaseAlphaWebAppRequest) (*alphapb.ListFirebaseAlphaWebAppResponse, error) {
	cl, err := createConfigWebApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWebApp(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.FirebaseAlphaWebApp
	for _, r := range resources.Items {
		rp := WebAppToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListFirebaseAlphaWebAppResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWebApp(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
