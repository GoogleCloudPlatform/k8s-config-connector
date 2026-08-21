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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gameservices/alpha/gameservices_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gameservices/alpha"
)

// RealmServer implements the gRPC interface for Realm.
type RealmServer struct{}

// ProtoToRealm converts a Realm resource from its proto representation.
func ProtoToRealm(p *alphapb.GameservicesAlphaRealm) *alpha.Realm {
	obj := &alpha.Realm{
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
func RealmToProto(resource *alpha.Realm) *alphapb.GameservicesAlphaRealm {
	p := &alphapb.GameservicesAlphaRealm{}
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
func (s *RealmServer) applyRealm(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGameservicesAlphaRealmRequest) (*alphapb.GameservicesAlphaRealm, error) {
	p := ProtoToRealm(request.GetResource())
	res, err := c.ApplyRealm(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RealmToProto(res)
	return r, nil
}

// applyGameservicesAlphaRealm handles the gRPC request by passing it to the underlying Realm Apply() method.
func (s *RealmServer) ApplyGameservicesAlphaRealm(ctx context.Context, request *alphapb.ApplyGameservicesAlphaRealmRequest) (*alphapb.GameservicesAlphaRealm, error) {
	cl, err := createConfigRealm(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRealm(ctx, cl, request)
}

// DeleteRealm handles the gRPC request by passing it to the underlying Realm Delete() method.
func (s *RealmServer) DeleteGameservicesAlphaRealm(ctx context.Context, request *alphapb.DeleteGameservicesAlphaRealmRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRealm(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRealm(ctx, ProtoToRealm(request.GetResource()))

}

// ListGameservicesAlphaRealm handles the gRPC request by passing it to the underlying RealmList() method.
func (s *RealmServer) ListGameservicesAlphaRealm(ctx context.Context, request *alphapb.ListGameservicesAlphaRealmRequest) (*alphapb.ListGameservicesAlphaRealmResponse, error) {
	cl, err := createConfigRealm(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRealm(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GameservicesAlphaRealm
	for _, r := range resources.Items {
		rp := RealmToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListGameservicesAlphaRealmResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRealm(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
