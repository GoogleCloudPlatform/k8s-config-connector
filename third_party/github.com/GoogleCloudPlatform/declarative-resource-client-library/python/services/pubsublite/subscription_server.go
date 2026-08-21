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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	pubsublitepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/pubsublite/pubsublite_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsublite"
)

// Server implements the gRPC interface for Subscription.
type SubscriptionServer struct{}

// ProtoToSubscriptionDeliveryConfigDeliveryRequirementEnum converts a SubscriptionDeliveryConfigDeliveryRequirementEnum enum from its proto representation.
func ProtoToPubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum(e pubsublitepb.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum) *pubsublite.SubscriptionDeliveryConfigDeliveryRequirementEnum {
	if e == 0 {
		return nil
	}
	if n, ok := pubsublitepb.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum_name[int32(e)]; ok {
		e := pubsublite.SubscriptionDeliveryConfigDeliveryRequirementEnum(n[len("PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum"):])
		return &e
	}
	return nil
}

// ProtoToSubscriptionDeliveryConfig converts a SubscriptionDeliveryConfig resource from its proto representation.
func ProtoToPubsubliteSubscriptionDeliveryConfig(p *pubsublitepb.PubsubliteSubscriptionDeliveryConfig) *pubsublite.SubscriptionDeliveryConfig {
	if p == nil {
		return nil
	}
	obj := &pubsublite.SubscriptionDeliveryConfig{
		DeliveryRequirement: ProtoToPubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum(p.GetDeliveryRequirement()),
	}
	return obj
}

// ProtoToSubscription converts a Subscription resource from its proto representation.
func ProtoToSubscription(p *pubsublitepb.PubsubliteSubscription) *pubsublite.Subscription {
	obj := &pubsublite.Subscription{
		Name:           dcl.StringOrNil(p.Name),
		Topic:          dcl.StringOrNil(p.Topic),
		DeliveryConfig: ProtoToPubsubliteSubscriptionDeliveryConfig(p.GetDeliveryConfig()),
		Project:        dcl.StringOrNil(p.Project),
		Location:       dcl.StringOrNil(p.Location),
	}
	return obj
}

// SubscriptionDeliveryConfigDeliveryRequirementEnumToProto converts a SubscriptionDeliveryConfigDeliveryRequirementEnum enum to its proto representation.
func PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnumToProto(e *pubsublite.SubscriptionDeliveryConfigDeliveryRequirementEnum) pubsublitepb.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum {
	if e == nil {
		return pubsublitepb.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum(0)
	}
	if v, ok := pubsublitepb.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum_value["SubscriptionDeliveryConfigDeliveryRequirementEnum"+string(*e)]; ok {
		return pubsublitepb.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum(v)
	}
	return pubsublitepb.PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnum(0)
}

// SubscriptionDeliveryConfigToProto converts a SubscriptionDeliveryConfig resource to its proto representation.
func PubsubliteSubscriptionDeliveryConfigToProto(o *pubsublite.SubscriptionDeliveryConfig) *pubsublitepb.PubsubliteSubscriptionDeliveryConfig {
	if o == nil {
		return nil
	}
	p := &pubsublitepb.PubsubliteSubscriptionDeliveryConfig{
		DeliveryRequirement: PubsubliteSubscriptionDeliveryConfigDeliveryRequirementEnumToProto(o.DeliveryRequirement),
	}
	return p
}

// SubscriptionToProto converts a Subscription resource to its proto representation.
func SubscriptionToProto(resource *pubsublite.Subscription) *pubsublitepb.PubsubliteSubscription {
	p := &pubsublitepb.PubsubliteSubscription{
		Name:           dcl.ValueOrEmptyString(resource.Name),
		Topic:          dcl.ValueOrEmptyString(resource.Topic),
		DeliveryConfig: PubsubliteSubscriptionDeliveryConfigToProto(resource.DeliveryConfig),
		Project:        dcl.ValueOrEmptyString(resource.Project),
		Location:       dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplySubscription handles the gRPC request by passing it to the underlying Subscription Apply() method.
func (s *SubscriptionServer) applySubscription(ctx context.Context, c *pubsublite.Client, request *pubsublitepb.ApplyPubsubliteSubscriptionRequest) (*pubsublitepb.PubsubliteSubscription, error) {
	p := ProtoToSubscription(request.GetResource())
	res, err := c.ApplySubscription(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SubscriptionToProto(res)
	return r, nil
}

// ApplySubscription handles the gRPC request by passing it to the underlying Subscription Apply() method.
func (s *SubscriptionServer) ApplyPubsubliteSubscription(ctx context.Context, request *pubsublitepb.ApplyPubsubliteSubscriptionRequest) (*pubsublitepb.PubsubliteSubscription, error) {
	cl, err := createConfigSubscription(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySubscription(ctx, cl, request)
}

// DeleteSubscription handles the gRPC request by passing it to the underlying Subscription Delete() method.
func (s *SubscriptionServer) DeletePubsubliteSubscription(ctx context.Context, request *pubsublitepb.DeletePubsubliteSubscriptionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSubscription(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSubscription(ctx, ProtoToSubscription(request.GetResource()))

}

// ListPubsubliteSubscription handles the gRPC request by passing it to the underlying SubscriptionList() method.
func (s *SubscriptionServer) ListPubsubliteSubscription(ctx context.Context, request *pubsublitepb.ListPubsubliteSubscriptionRequest) (*pubsublitepb.ListPubsubliteSubscriptionResponse, error) {
	cl, err := createConfigSubscription(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSubscription(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*pubsublitepb.PubsubliteSubscription
	for _, r := range resources.Items {
		rp := SubscriptionToProto(r)
		protos = append(protos, rp)
	}
	return &pubsublitepb.ListPubsubliteSubscriptionResponse{Items: protos}, nil
}

func createConfigSubscription(ctx context.Context, service_account_file string) (*pubsublite.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return pubsublite.NewClient(conf), nil
}
