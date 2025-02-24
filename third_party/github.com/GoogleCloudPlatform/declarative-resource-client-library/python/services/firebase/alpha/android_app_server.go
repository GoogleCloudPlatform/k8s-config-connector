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

// AndroidAppServer implements the gRPC interface for AndroidApp.
type AndroidAppServer struct{}

// ProtoToAndroidApp converts a AndroidApp resource from its proto representation.
func ProtoToAndroidApp(p *alphapb.FirebaseAlphaAndroidApp) *alpha.AndroidApp {
	obj := &alpha.AndroidApp{
		Name:        dcl.StringOrNil(p.GetName()),
		AppId:       dcl.StringOrNil(p.GetAppId()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		ProjectId:   dcl.StringOrNil(p.GetProjectId()),
		PackageName: dcl.StringOrNil(p.GetPackageName()),
		ApiKeyId:    dcl.StringOrNil(p.GetApiKeyId()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// AndroidAppToProto converts a AndroidApp resource to its proto representation.
func AndroidAppToProto(resource *alpha.AndroidApp) *alphapb.FirebaseAlphaAndroidApp {
	p := &alphapb.FirebaseAlphaAndroidApp{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetAppId(dcl.ValueOrEmptyString(resource.AppId))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetProjectId(dcl.ValueOrEmptyString(resource.ProjectId))
	p.SetPackageName(dcl.ValueOrEmptyString(resource.PackageName))
	p.SetApiKeyId(dcl.ValueOrEmptyString(resource.ApiKeyId))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyAndroidApp handles the gRPC request by passing it to the underlying AndroidApp Apply() method.
func (s *AndroidAppServer) applyAndroidApp(ctx context.Context, c *alpha.Client, request *alphapb.ApplyFirebaseAlphaAndroidAppRequest) (*alphapb.FirebaseAlphaAndroidApp, error) {
	p := ProtoToAndroidApp(request.GetResource())
	res, err := c.ApplyAndroidApp(ctx, p)
	if err != nil {
		return nil, err
	}
	r := AndroidAppToProto(res)
	return r, nil
}

// applyFirebaseAlphaAndroidApp handles the gRPC request by passing it to the underlying AndroidApp Apply() method.
func (s *AndroidAppServer) ApplyFirebaseAlphaAndroidApp(ctx context.Context, request *alphapb.ApplyFirebaseAlphaAndroidAppRequest) (*alphapb.FirebaseAlphaAndroidApp, error) {
	cl, err := createConfigAndroidApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyAndroidApp(ctx, cl, request)
}

// DeleteAndroidApp handles the gRPC request by passing it to the underlying AndroidApp Delete() method.
func (s *AndroidAppServer) DeleteFirebaseAlphaAndroidApp(ctx context.Context, request *alphapb.DeleteFirebaseAlphaAndroidAppRequest) (*emptypb.Empty, error) {

	cl, err := createConfigAndroidApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteAndroidApp(ctx, ProtoToAndroidApp(request.GetResource()))

}

// ListFirebaseAlphaAndroidApp handles the gRPC request by passing it to the underlying AndroidAppList() method.
func (s *AndroidAppServer) ListFirebaseAlphaAndroidApp(ctx context.Context, request *alphapb.ListFirebaseAlphaAndroidAppRequest) (*alphapb.ListFirebaseAlphaAndroidAppResponse, error) {
	cl, err := createConfigAndroidApp(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListAndroidApp(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.FirebaseAlphaAndroidApp
	for _, r := range resources.Items {
		rp := AndroidAppToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListFirebaseAlphaAndroidAppResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigAndroidApp(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
