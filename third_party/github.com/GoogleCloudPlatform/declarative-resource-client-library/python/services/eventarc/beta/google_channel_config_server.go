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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/eventarc/beta/eventarc_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta"
)

// GoogleChannelConfigServer implements the gRPC interface for GoogleChannelConfig.
type GoogleChannelConfigServer struct{}

// ProtoToGoogleChannelConfig converts a GoogleChannelConfig resource from its proto representation.
func ProtoToGoogleChannelConfig(p *betapb.EventarcBetaGoogleChannelConfig) *beta.GoogleChannelConfig {
	obj := &beta.GoogleChannelConfig{
		Name:          dcl.StringOrNil(p.GetName()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		CryptoKeyName: dcl.StringOrNil(p.GetCryptoKeyName()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// GoogleChannelConfigToProto converts a GoogleChannelConfig resource to its proto representation.
func GoogleChannelConfigToProto(resource *beta.GoogleChannelConfig) *betapb.EventarcBetaGoogleChannelConfig {
	p := &betapb.EventarcBetaGoogleChannelConfig{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetCryptoKeyName(dcl.ValueOrEmptyString(resource.CryptoKeyName))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyGoogleChannelConfig handles the gRPC request by passing it to the underlying GoogleChannelConfig Apply() method.
func (s *GoogleChannelConfigServer) applyGoogleChannelConfig(ctx context.Context, c *beta.Client, request *betapb.ApplyEventarcBetaGoogleChannelConfigRequest) (*betapb.EventarcBetaGoogleChannelConfig, error) {
	p := ProtoToGoogleChannelConfig(request.GetResource())
	res, err := c.ApplyGoogleChannelConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GoogleChannelConfigToProto(res)
	return r, nil
}

// applyEventarcBetaGoogleChannelConfig handles the gRPC request by passing it to the underlying GoogleChannelConfig Apply() method.
func (s *GoogleChannelConfigServer) ApplyEventarcBetaGoogleChannelConfig(ctx context.Context, request *betapb.ApplyEventarcBetaGoogleChannelConfigRequest) (*betapb.EventarcBetaGoogleChannelConfig, error) {
	cl, err := createConfigGoogleChannelConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGoogleChannelConfig(ctx, cl, request)
}

// DeleteGoogleChannelConfig handles the gRPC request by passing it to the underlying GoogleChannelConfig Delete() method.
func (s *GoogleChannelConfigServer) DeleteEventarcBetaGoogleChannelConfig(ctx context.Context, request *betapb.DeleteEventarcBetaGoogleChannelConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGoogleChannelConfig(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGoogleChannelConfig(ctx, ProtoToGoogleChannelConfig(request.GetResource()))

}

// ListEventarcBetaGoogleChannelConfig is a no-op method because GoogleChannelConfig has no list method.
func (s *GoogleChannelConfigServer) ListEventarcBetaGoogleChannelConfig(_ context.Context, _ *betapb.ListEventarcBetaGoogleChannelConfigRequest) (*betapb.ListEventarcBetaGoogleChannelConfigResponse, error) {
	return nil, nil
}

func createConfigGoogleChannelConfig(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
