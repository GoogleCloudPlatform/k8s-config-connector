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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/eventarc/alpha/eventarc_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/alpha"
)

// ChannelServer implements the gRPC interface for Channel.
type ChannelServer struct{}

// ProtoToChannelStateEnum converts a ChannelStateEnum enum from its proto representation.
func ProtoToEventarcAlphaChannelStateEnum(e alphapb.EventarcAlphaChannelStateEnum) *alpha.ChannelStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.EventarcAlphaChannelStateEnum_name[int32(e)]; ok {
		e := alpha.ChannelStateEnum(n[len("EventarcAlphaChannelStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToChannel converts a Channel resource from its proto representation.
func ProtoToChannel(p *alphapb.EventarcAlphaChannel) *alpha.Channel {
	obj := &alpha.Channel{
		Name:               dcl.StringOrNil(p.GetName()),
		Uid:                dcl.StringOrNil(p.GetUid()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		ThirdPartyProvider: dcl.StringOrNil(p.GetThirdPartyProvider()),
		PubsubTopic:        dcl.StringOrNil(p.GetPubsubTopic()),
		State:              ProtoToEventarcAlphaChannelStateEnum(p.GetState()),
		ActivationToken:    dcl.StringOrNil(p.GetActivationToken()),
		CryptoKeyName:      dcl.StringOrNil(p.GetCryptoKeyName()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ChannelStateEnumToProto converts a ChannelStateEnum enum to its proto representation.
func EventarcAlphaChannelStateEnumToProto(e *alpha.ChannelStateEnum) alphapb.EventarcAlphaChannelStateEnum {
	if e == nil {
		return alphapb.EventarcAlphaChannelStateEnum(0)
	}
	if v, ok := alphapb.EventarcAlphaChannelStateEnum_value["ChannelStateEnum"+string(*e)]; ok {
		return alphapb.EventarcAlphaChannelStateEnum(v)
	}
	return alphapb.EventarcAlphaChannelStateEnum(0)
}

// ChannelToProto converts a Channel resource to its proto representation.
func ChannelToProto(resource *alpha.Channel) *alphapb.EventarcAlphaChannel {
	p := &alphapb.EventarcAlphaChannel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetThirdPartyProvider(dcl.ValueOrEmptyString(resource.ThirdPartyProvider))
	p.SetPubsubTopic(dcl.ValueOrEmptyString(resource.PubsubTopic))
	p.SetState(EventarcAlphaChannelStateEnumToProto(resource.State))
	p.SetActivationToken(dcl.ValueOrEmptyString(resource.ActivationToken))
	p.SetCryptoKeyName(dcl.ValueOrEmptyString(resource.CryptoKeyName))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyChannel handles the gRPC request by passing it to the underlying Channel Apply() method.
func (s *ChannelServer) applyChannel(ctx context.Context, c *alpha.Client, request *alphapb.ApplyEventarcAlphaChannelRequest) (*alphapb.EventarcAlphaChannel, error) {
	p := ProtoToChannel(request.GetResource())
	res, err := c.ApplyChannel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ChannelToProto(res)
	return r, nil
}

// applyEventarcAlphaChannel handles the gRPC request by passing it to the underlying Channel Apply() method.
func (s *ChannelServer) ApplyEventarcAlphaChannel(ctx context.Context, request *alphapb.ApplyEventarcAlphaChannelRequest) (*alphapb.EventarcAlphaChannel, error) {
	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyChannel(ctx, cl, request)
}

// DeleteChannel handles the gRPC request by passing it to the underlying Channel Delete() method.
func (s *ChannelServer) DeleteEventarcAlphaChannel(ctx context.Context, request *alphapb.DeleteEventarcAlphaChannelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteChannel(ctx, ProtoToChannel(request.GetResource()))

}

// ListEventarcAlphaChannel handles the gRPC request by passing it to the underlying ChannelList() method.
func (s *ChannelServer) ListEventarcAlphaChannel(ctx context.Context, request *alphapb.ListEventarcAlphaChannelRequest) (*alphapb.ListEventarcAlphaChannelResponse, error) {
	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListChannel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.EventarcAlphaChannel
	for _, r := range resources.Items {
		rp := ChannelToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListEventarcAlphaChannelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigChannel(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
