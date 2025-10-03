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
	eventarcpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/eventarc/eventarc_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc"
)

// ChannelServer implements the gRPC interface for Channel.
type ChannelServer struct{}

// ProtoToChannelStateEnum converts a ChannelStateEnum enum from its proto representation.
func ProtoToEventarcChannelStateEnum(e eventarcpb.EventarcChannelStateEnum) *eventarc.ChannelStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := eventarcpb.EventarcChannelStateEnum_name[int32(e)]; ok {
		e := eventarc.ChannelStateEnum(n[len("EventarcChannelStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToChannel converts a Channel resource from its proto representation.
func ProtoToChannel(p *eventarcpb.EventarcChannel) *eventarc.Channel {
	obj := &eventarc.Channel{
		Name:               dcl.StringOrNil(p.GetName()),
		Uid:                dcl.StringOrNil(p.GetUid()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:         dcl.StringOrNil(p.GetUpdateTime()),
		ThirdPartyProvider: dcl.StringOrNil(p.GetThirdPartyProvider()),
		PubsubTopic:        dcl.StringOrNil(p.GetPubsubTopic()),
		State:              ProtoToEventarcChannelStateEnum(p.GetState()),
		ActivationToken:    dcl.StringOrNil(p.GetActivationToken()),
		CryptoKeyName:      dcl.StringOrNil(p.GetCryptoKeyName()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// ChannelStateEnumToProto converts a ChannelStateEnum enum to its proto representation.
func EventarcChannelStateEnumToProto(e *eventarc.ChannelStateEnum) eventarcpb.EventarcChannelStateEnum {
	if e == nil {
		return eventarcpb.EventarcChannelStateEnum(0)
	}
	if v, ok := eventarcpb.EventarcChannelStateEnum_value["ChannelStateEnum"+string(*e)]; ok {
		return eventarcpb.EventarcChannelStateEnum(v)
	}
	return eventarcpb.EventarcChannelStateEnum(0)
}

// ChannelToProto converts a Channel resource to its proto representation.
func ChannelToProto(resource *eventarc.Channel) *eventarcpb.EventarcChannel {
	p := &eventarcpb.EventarcChannel{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetThirdPartyProvider(dcl.ValueOrEmptyString(resource.ThirdPartyProvider))
	p.SetPubsubTopic(dcl.ValueOrEmptyString(resource.PubsubTopic))
	p.SetState(EventarcChannelStateEnumToProto(resource.State))
	p.SetActivationToken(dcl.ValueOrEmptyString(resource.ActivationToken))
	p.SetCryptoKeyName(dcl.ValueOrEmptyString(resource.CryptoKeyName))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyChannel handles the gRPC request by passing it to the underlying Channel Apply() method.
func (s *ChannelServer) applyChannel(ctx context.Context, c *eventarc.Client, request *eventarcpb.ApplyEventarcChannelRequest) (*eventarcpb.EventarcChannel, error) {
	p := ProtoToChannel(request.GetResource())
	res, err := c.ApplyChannel(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ChannelToProto(res)
	return r, nil
}

// applyEventarcChannel handles the gRPC request by passing it to the underlying Channel Apply() method.
func (s *ChannelServer) ApplyEventarcChannel(ctx context.Context, request *eventarcpb.ApplyEventarcChannelRequest) (*eventarcpb.EventarcChannel, error) {
	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyChannel(ctx, cl, request)
}

// DeleteChannel handles the gRPC request by passing it to the underlying Channel Delete() method.
func (s *ChannelServer) DeleteEventarcChannel(ctx context.Context, request *eventarcpb.DeleteEventarcChannelRequest) (*emptypb.Empty, error) {

	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteChannel(ctx, ProtoToChannel(request.GetResource()))

}

// ListEventarcChannel handles the gRPC request by passing it to the underlying ChannelList() method.
func (s *ChannelServer) ListEventarcChannel(ctx context.Context, request *eventarcpb.ListEventarcChannelRequest) (*eventarcpb.ListEventarcChannelResponse, error) {
	cl, err := createConfigChannel(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListChannel(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*eventarcpb.EventarcChannel
	for _, r := range resources.Items {
		rp := ChannelToProto(r)
		protos = append(protos, rp)
	}
	p := &eventarcpb.ListEventarcChannelResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigChannel(ctx context.Context, service_account_file string) (*eventarc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return eventarc.NewClient(conf), nil
}
