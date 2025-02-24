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

// ChannelServer implements the gRPC interface for Channel.
type ChannelServer struct{}

// ProtoToChannelStateEnum converts a ChannelStateEnum enum from its proto representation.
func ProtoToEventarcBetaChannelStateEnum(e betapb.EventarcBetaChannelStateEnum) *beta.ChannelStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.EventarcBetaChannelStateEnum_name[int32(e)]; ok {
		e := beta.ChannelStateEnum(n[len("EventarcBetaChannelStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToChannel converts a Channel resource from its proto representation.
func ProtoToChannel(p *betapb.EventarcBetaChannel) *beta.Channel {
	obj := &beta.Channel{
		Name:               dcl.StringOrNil(p.GetName()),
		Uid:                dcl.StringOrNil(p.GetUid()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		ThirdPartyProvider: dcl.StringOrNil(p.GetThirdPartyProvider()),
		PubsubTopic:        dcl.StringOrNil(p.GetPubsubTopic()),
		State:              ProtoToEventarcBetaChannelStateEnum(p.GetState()),
		ActivationToken:    dcl.StringOrNil(p.GetActivationToken()),
		CryptoKeyName:      dcl.StringOrNil(p.GetCryptoKeyName()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ChannelStateEnumToProto converts a ChannelStateEnum enum to its proto representation.
func EventarcBetaChannelStateEnumToProto(e *beta.ChannelStateEnum) betapb.EventarcBetaChannelStateEnum {
	if e == nil {
		return betapb.EventarcBetaChannelStateEnum(0)
	}
	if v, ok := betapb.EventarcBetaChannelStateEnum_value["ChannelStateEnum"+string(*e)]; ok {
		return betapb.EventarcBetaChannelStateEnum(v)
	}
	return betapb.EventarcBetaChannelStateEnum(0)
}

// ChannelToProto converts a Channel resource to its proto representation.
func ChannelToProto(resource *beta.Channel) *betapb.EventarcBetaChannel {
	p := &betapb.EventarcBetaChannel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetThirdPartyProvider(dcl.ValueOrEmptyString(resource.ThirdPartyProvider))
	p.SetPubsubTopic(dcl.ValueOrEmptyString(resource.PubsubTopic))
	p.SetState(EventarcBetaChannelStateEnumToProto(resource.State))
	p.SetActivationToken(dcl.ValueOrEmptyString(resource.ActivationToken))
	p.SetCryptoKeyName(dcl.ValueOrEmptyString(resource.CryptoKeyName))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyChannel handles the gRPC request by passing it to the underlying Channel Apply() method.
func (s *ChannelServer) applyChannel(ctx context.Context, c *beta.Client, request *betapb.ApplyEventarcBetaChannelRequest) (*betapb.EventarcBetaChannel, error) {
	p := ProtoToChannel(request.GetResource())
	res, err := c.ApplyChannel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ChannelToProto(res)
	return r, nil
}

// applyEventarcBetaChannel handles the gRPC request by passing it to the underlying Channel Apply() method.
func (s *ChannelServer) ApplyEventarcBetaChannel(ctx context.Context, request *betapb.ApplyEventarcBetaChannelRequest) (*betapb.EventarcBetaChannel, error) {
	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyChannel(ctx, cl, request)
}

// DeleteChannel handles the gRPC request by passing it to the underlying Channel Delete() method.
func (s *ChannelServer) DeleteEventarcBetaChannel(ctx context.Context, request *betapb.DeleteEventarcBetaChannelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteChannel(ctx, ProtoToChannel(request.GetResource()))

}

// ListEventarcBetaChannel handles the gRPC request by passing it to the underlying ChannelList() method.
func (s *ChannelServer) ListEventarcBetaChannel(ctx context.Context, request *betapb.ListEventarcBetaChannelRequest) (*betapb.ListEventarcBetaChannelResponse, error) {
	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListChannel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.EventarcBetaChannel
	for _, r := range resources.Items {
		rp := ChannelToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListEventarcBetaChannelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigChannel(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
