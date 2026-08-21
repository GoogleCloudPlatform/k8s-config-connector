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

// Server implements the gRPC interface for Topic.
type TopicServer struct{}

// ProtoToTopicPartitionConfig converts a TopicPartitionConfig resource from its proto representation.
func ProtoToPubsubliteTopicPartitionConfig(p *pubsublitepb.PubsubliteTopicPartitionConfig) *pubsublite.TopicPartitionConfig {
	if p == nil {
		return nil
	}
	obj := &pubsublite.TopicPartitionConfig{
		Count:    dcl.Int64OrNil(p.Count),
		Capacity: ProtoToPubsubliteTopicPartitionConfigCapacity(p.GetCapacity()),
	}
	return obj
}

// ProtoToTopicPartitionConfigCapacity converts a TopicPartitionConfigCapacity resource from its proto representation.
func ProtoToPubsubliteTopicPartitionConfigCapacity(p *pubsublitepb.PubsubliteTopicPartitionConfigCapacity) *pubsublite.TopicPartitionConfigCapacity {
	if p == nil {
		return nil
	}
	obj := &pubsublite.TopicPartitionConfigCapacity{
		PublishMibPerSec:   dcl.Int64OrNil(p.PublishMibPerSec),
		SubscribeMibPerSec: dcl.Int64OrNil(p.SubscribeMibPerSec),
	}
	return obj
}

// ProtoToTopicRetentionConfig converts a TopicRetentionConfig resource from its proto representation.
func ProtoToPubsubliteTopicRetentionConfig(p *pubsublitepb.PubsubliteTopicRetentionConfig) *pubsublite.TopicRetentionConfig {
	if p == nil {
		return nil
	}
	obj := &pubsublite.TopicRetentionConfig{
		PerPartitionBytes: dcl.Int64OrNil(p.PerPartitionBytes),
		Period:            dcl.StringOrNil(p.Period),
	}
	return obj
}

// ProtoToTopic converts a Topic resource from its proto representation.
func ProtoToTopic(p *pubsublitepb.PubsubliteTopic) *pubsublite.Topic {
	obj := &pubsublite.Topic{
		Name:            dcl.StringOrNil(p.Name),
		PartitionConfig: ProtoToPubsubliteTopicPartitionConfig(p.GetPartitionConfig()),
		RetentionConfig: ProtoToPubsubliteTopicRetentionConfig(p.GetRetentionConfig()),
		Project:         dcl.StringOrNil(p.Project),
		Location:        dcl.StringOrNil(p.Location),
	}
	return obj
}

// TopicPartitionConfigToProto converts a TopicPartitionConfig resource to its proto representation.
func PubsubliteTopicPartitionConfigToProto(o *pubsublite.TopicPartitionConfig) *pubsublitepb.PubsubliteTopicPartitionConfig {
	if o == nil {
		return nil
	}
	p := &pubsublitepb.PubsubliteTopicPartitionConfig{
		Count:    dcl.ValueOrEmptyInt64(o.Count),
		Capacity: PubsubliteTopicPartitionConfigCapacityToProto(o.Capacity),
	}
	return p
}

// TopicPartitionConfigCapacityToProto converts a TopicPartitionConfigCapacity resource to its proto representation.
func PubsubliteTopicPartitionConfigCapacityToProto(o *pubsublite.TopicPartitionConfigCapacity) *pubsublitepb.PubsubliteTopicPartitionConfigCapacity {
	if o == nil {
		return nil
	}
	p := &pubsublitepb.PubsubliteTopicPartitionConfigCapacity{
		PublishMibPerSec:   dcl.ValueOrEmptyInt64(o.PublishMibPerSec),
		SubscribeMibPerSec: dcl.ValueOrEmptyInt64(o.SubscribeMibPerSec),
	}
	return p
}

// TopicRetentionConfigToProto converts a TopicRetentionConfig resource to its proto representation.
func PubsubliteTopicRetentionConfigToProto(o *pubsublite.TopicRetentionConfig) *pubsublitepb.PubsubliteTopicRetentionConfig {
	if o == nil {
		return nil
	}
	p := &pubsublitepb.PubsubliteTopicRetentionConfig{
		PerPartitionBytes: dcl.ValueOrEmptyInt64(o.PerPartitionBytes),
		Period:            dcl.ValueOrEmptyString(o.Period),
	}
	return p
}

// TopicToProto converts a Topic resource to its proto representation.
func TopicToProto(resource *pubsublite.Topic) *pubsublitepb.PubsubliteTopic {
	p := &pubsublitepb.PubsubliteTopic{
		Name:            dcl.ValueOrEmptyString(resource.Name),
		PartitionConfig: PubsubliteTopicPartitionConfigToProto(resource.PartitionConfig),
		RetentionConfig: PubsubliteTopicRetentionConfigToProto(resource.RetentionConfig),
		Project:         dcl.ValueOrEmptyString(resource.Project),
		Location:        dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) applyTopic(ctx context.Context, c *pubsublite.Client, request *pubsublitepb.ApplyPubsubliteTopicRequest) (*pubsublitepb.PubsubliteTopic, error) {
	p := ProtoToTopic(request.GetResource())
	res, err := c.ApplyTopic(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TopicToProto(res)
	return r, nil
}

// ApplyTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) ApplyPubsubliteTopic(ctx context.Context, request *pubsublitepb.ApplyPubsubliteTopicRequest) (*pubsublitepb.PubsubliteTopic, error) {
	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTopic(ctx, cl, request)
}

// DeleteTopic handles the gRPC request by passing it to the underlying Topic Delete() method.
func (s *TopicServer) DeletePubsubliteTopic(ctx context.Context, request *pubsublitepb.DeletePubsubliteTopicRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTopic(ctx, ProtoToTopic(request.GetResource()))

}

// ListPubsubliteTopic handles the gRPC request by passing it to the underlying TopicList() method.
func (s *TopicServer) ListPubsubliteTopic(ctx context.Context, request *pubsublitepb.ListPubsubliteTopicRequest) (*pubsublitepb.ListPubsubliteTopicResponse, error) {
	cl, err := createConfigTopic(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTopic(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*pubsublitepb.PubsubliteTopic
	for _, r := range resources.Items {
		rp := TopicToProto(r)
		protos = append(protos, rp)
	}
	return &pubsublitepb.ListPubsubliteTopicResponse{Items: protos}, nil
}

func createConfigTopic(ctx context.Context, service_account_file string) (*pubsublite.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return pubsublite.NewClient(conf), nil
}
