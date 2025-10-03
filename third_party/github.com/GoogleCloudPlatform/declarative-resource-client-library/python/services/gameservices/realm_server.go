// Copyright 2023 Google LLC. All Rights Reserved.
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
	gameservicespb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gameservices/gameservices_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices"
)

// RealmServer implements the gRPC interface for Realm.
type RealmServer struct{}

// ProtoToRealm converts a Realm resource from its proto representation.
func ProtoToRealm(p *gameservicespb.GameservicesRealm) *gameservices.Realm {
	obj := &gameservices.Realm{
		Name:        dcl.StringOrNil(p.GetName()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		TimeZone:    dcl.StringOrNil(p.GetTimeZone()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Location:    dcl.StringOrNil(p.GetLocation()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// RealmToProto converts a Realm resource to its proto representation.
func RealmToProto(resource *gameservices.Realm) *gameservicespb.GameservicesRealm {
	p := &gameservicespb.GameservicesRealm{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetTimeZone(dcl.ValueOrEmptyString(resource.TimeZone))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyRealm handles the gRPC request by passing it to the underlying Realm Apply() method.
func (s *RealmServer) applyRealm(ctx context.Context, c *gameservices.Client, request *gameservicespb.ApplyGameservicesRealmRequest) (*gameservicespb.GameservicesRealm, error) {
	p := ProtoToRealm(request.GetResource())
	res, err := c.ApplyRealm(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RealmToProto(res)
	return r, nil
}

// applyGameservicesRealm handles the gRPC request by passing it to the underlying Realm Apply() method.
func (s *RealmServer) ApplyGameservicesRealm(ctx context.Context, request *gameservicespb.ApplyGameservicesRealmRequest) (*gameservicespb.GameservicesRealm, error) {
	cl, err := createConfigRealm(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRealm(ctx, cl, request)
}

// DeleteRealm handles the gRPC request by passing it to the underlying Realm Delete() method.
func (s *RealmServer) DeleteGameservicesRealm(ctx context.Context, request *gameservicespb.DeleteGameservicesRealmRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRealm(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRealm(ctx, ProtoToRealm(request.GetResource()))

}

// ListGameservicesRealm handles the gRPC request by passing it to the underlying RealmList() method.
func (s *RealmServer) ListGameservicesRealm(ctx context.Context, request *gameservicespb.ListGameservicesRealmRequest) (*gameservicespb.ListGameservicesRealmResponse, error) {
	cl, err := createConfigRealm(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRealm(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*gameservicespb.GameservicesRealm
	for _, r := range resources.Items {
		rp := RealmToProto(r)
		protos = append(protos, rp)
	}
	p := &gameservicespb.ListGameservicesRealmResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRealm(ctx context.Context, service_account_file string) (*gameservices.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return gameservices.NewClient(conf), nil
}
