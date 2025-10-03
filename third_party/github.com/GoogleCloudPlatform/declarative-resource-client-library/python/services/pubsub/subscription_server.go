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
	pubsubpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/pubsub/pubsub_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub"
)

// Server implements the gRPC interface for Subscription.
type SubscriptionServer struct{}

// ProtoToSubscriptionExpirationPolicy converts a SubscriptionExpirationPolicy resource from its proto representation.
func ProtoToPubsubSubscriptionExpirationPolicy(p *pubsubpb.PubsubSubscriptionExpirationPolicy) *pubsub.SubscriptionExpirationPolicy {
	if p == nil {
		return nil
	}
	obj := &pubsub.SubscriptionExpirationPolicy{
		Ttl: dcl.StringOrNil(p.Ttl),
	}
	return obj
}

// ProtoToSubscriptionDeadLetterPolicy converts a SubscriptionDeadLetterPolicy resource from its proto representation.
func ProtoToPubsubSubscriptionDeadLetterPolicy(p *pubsubpb.PubsubSubscriptionDeadLetterPolicy) *pubsub.SubscriptionDeadLetterPolicy {
	if p == nil {
		return nil
	}
	obj := &pubsub.SubscriptionDeadLetterPolicy{
		DeadLetterTopic:     dcl.StringOrNil(p.DeadLetterTopic),
		MaxDeliveryAttempts: dcl.Int64OrNil(p.MaxDeliveryAttempts),
	}
	return obj
}

// ProtoToSubscriptionPushConfig converts a SubscriptionPushConfig resource from its proto representation.
func ProtoToPubsubSubscriptionPushConfig(p *pubsubpb.PubsubSubscriptionPushConfig) *pubsub.SubscriptionPushConfig {
	if p == nil {
		return nil
	}
	obj := &pubsub.SubscriptionPushConfig{
		PushEndpoint: dcl.StringOrNil(p.PushEndpoint),
		OidcToken:    ProtoToPubsubSubscriptionPushConfigOidcToken(p.GetOidcToken()),
	}
	return obj
}

// ProtoToSubscriptionPushConfigOidcToken converts a SubscriptionPushConfigOidcToken resource from its proto representation.
func ProtoToPubsubSubscriptionPushConfigOidcToken(p *pubsubpb.PubsubSubscriptionPushConfigOidcToken) *pubsub.SubscriptionPushConfigOidcToken {
	if p == nil {
		return nil
	}
	obj := &pubsub.SubscriptionPushConfigOidcToken{
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
		Audience:            dcl.StringOrNil(p.Audience),
	}
	return obj
}

// ProtoToSubscription converts a Subscription resource from its proto representation.
func ProtoToSubscription(p *pubsubpb.PubsubSubscription) *pubsub.Subscription {
	obj := &pubsub.Subscription{
		Name:                     dcl.StringOrNil(p.Name),
		Topic:                    dcl.StringOrNil(p.Topic),
		MessageRetentionDuration: dcl.StringOrNil(p.MessageRetentionDuration),
		RetainAckedMessages:      dcl.Bool(p.RetainAckedMessages),
		ExpirationPolicy:         ProtoToPubsubSubscriptionExpirationPolicy(p.GetExpirationPolicy()),
		Project:                  dcl.StringOrNil(p.Project),
		DeadLetterPolicy:         ProtoToPubsubSubscriptionDeadLetterPolicy(p.GetDeadLetterPolicy()),
		PushConfig:               ProtoToPubsubSubscriptionPushConfig(p.GetPushConfig()),
		AckDeadlineSeconds:       dcl.Int64OrNil(p.AckDeadlineSeconds),
	}
	return obj
}

// SubscriptionExpirationPolicyToProto converts a SubscriptionExpirationPolicy resource to its proto representation.
func PubsubSubscriptionExpirationPolicyToProto(o *pubsub.SubscriptionExpirationPolicy) *pubsubpb.PubsubSubscriptionExpirationPolicy {
	if o == nil {
		return nil
	}
	p := &pubsubpb.PubsubSubscriptionExpirationPolicy{
		Ttl: dcl.ValueOrEmptyString(o.Ttl),
	}
	return p
}

// SubscriptionDeadLetterPolicyToProto converts a SubscriptionDeadLetterPolicy resource to its proto representation.
func PubsubSubscriptionDeadLetterPolicyToProto(o *pubsub.SubscriptionDeadLetterPolicy) *pubsubpb.PubsubSubscriptionDeadLetterPolicy {
	if o == nil {
		return nil
	}
	p := &pubsubpb.PubsubSubscriptionDeadLetterPolicy{
		DeadLetterTopic:     dcl.ValueOrEmptyString(o.DeadLetterTopic),
		MaxDeliveryAttempts: dcl.ValueOrEmptyInt64(o.MaxDeliveryAttempts),
	}
	return p
}

// SubscriptionPushConfigToProto converts a SubscriptionPushConfig resource to its proto representation.
func PubsubSubscriptionPushConfigToProto(o *pubsub.SubscriptionPushConfig) *pubsubpb.PubsubSubscriptionPushConfig {
	if o == nil {
		return nil
	}
	p := &pubsubpb.PubsubSubscriptionPushConfig{
		PushEndpoint: dcl.ValueOrEmptyString(o.PushEndpoint),
		OidcToken:    PubsubSubscriptionPushConfigOidcTokenToProto(o.OidcToken),
	}
	p.Attributes = make(map[string]string)
	for k, r := range o.Attributes {
		p.Attributes[k] = r
	}
	return p
}

// SubscriptionPushConfigOidcTokenToProto converts a SubscriptionPushConfigOidcToken resource to its proto representation.
func PubsubSubscriptionPushConfigOidcTokenToProto(o *pubsub.SubscriptionPushConfigOidcToken) *pubsubpb.PubsubSubscriptionPushConfigOidcToken {
	if o == nil {
		return nil
	}
	p := &pubsubpb.PubsubSubscriptionPushConfigOidcToken{
		ServiceAccountEmail: dcl.ValueOrEmptyString(o.ServiceAccountEmail),
		Audience:            dcl.ValueOrEmptyString(o.Audience),
	}
	return p
}

// SubscriptionToProto converts a Subscription resource to its proto representation.
func SubscriptionToProto(resource *pubsub.Subscription) *pubsubpb.PubsubSubscription {
	p := &pubsubpb.PubsubSubscription{
		Name:                     dcl.ValueOrEmptyString(resource.Name),
		Topic:                    dcl.ValueOrEmptyString(resource.Topic),
		MessageRetentionDuration: dcl.ValueOrEmptyString(resource.MessageRetentionDuration),
		RetainAckedMessages:      dcl.ValueOrEmptyBool(resource.RetainAckedMessages),
		ExpirationPolicy:         PubsubSubscriptionExpirationPolicyToProto(resource.ExpirationPolicy),
		Project:                  dcl.ValueOrEmptyString(resource.Project),
		DeadLetterPolicy:         PubsubSubscriptionDeadLetterPolicyToProto(resource.DeadLetterPolicy),
		PushConfig:               PubsubSubscriptionPushConfigToProto(resource.PushConfig),
		AckDeadlineSeconds:       dcl.ValueOrEmptyInt64(resource.AckDeadlineSeconds),
	}

	return p
}

// ApplySubscription handles the gRPC request by passing it to the underlying Subscription Apply() method.
func (s *SubscriptionServer) applySubscription(ctx context.Context, c *pubsub.Client, request *pubsubpb.ApplyPubsubSubscriptionRequest) (*pubsubpb.PubsubSubscription, error) {
	p := ProtoToSubscription(request.GetResource())
	res, err := c.ApplySubscription(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SubscriptionToProto(res)
	return r, nil
}

// ApplySubscription handles the gRPC request by passing it to the underlying Subscription Apply() method.
func (s *SubscriptionServer) ApplyPubsubSubscription(ctx context.Context, request *pubsubpb.ApplyPubsubSubscriptionRequest) (*pubsubpb.PubsubSubscription, error) {
	cl, err := createConfigSubscription(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySubscription(ctx, cl, request)
}

// DeleteSubscription handles the gRPC request by passing it to the underlying Subscription Delete() method.
func (s *SubscriptionServer) DeletePubsubSubscription(ctx context.Context, request *pubsubpb.DeletePubsubSubscriptionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSubscription(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSubscription(ctx, ProtoToSubscription(request.GetResource()))

}

// ListPubsubSubscription handles the gRPC request by passing it to the underlying SubscriptionList() method.
func (s *SubscriptionServer) ListPubsubSubscription(ctx context.Context, request *pubsubpb.ListPubsubSubscriptionRequest) (*pubsubpb.ListPubsubSubscriptionResponse, error) {
	cl, err := createConfigSubscription(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSubscription(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*pubsubpb.PubsubSubscription
	for _, r := range resources.Items {
		rp := SubscriptionToProto(r)
		protos = append(protos, rp)
	}
	return &pubsubpb.ListPubsubSubscriptionResponse{Items: protos}, nil
}

func createConfigSubscription(ctx context.Context, service_account_file string) (*pubsub.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return pubsub.NewClient(conf), nil
}
