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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/pubsub/alpha/pubsub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/pubsub/alpha"
)

// TopicServer implements the gRPC interface for Topic.
type TopicServer struct{}

// ProtoToTopicMessageStoragePolicy converts a TopicMessageStoragePolicy object from its proto representation.
func ProtoToPubsubAlphaTopicMessageStoragePolicy(p *alphapb.PubsubAlphaTopicMessageStoragePolicy) *alpha.TopicMessageStoragePolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.TopicMessageStoragePolicy{}
	for _, r := range p.GetAllowedPersistenceRegions() {
		obj.AllowedPersistenceRegions = append(obj.AllowedPersistenceRegions, r)
	}
	return obj
}

// ProtoToTopic converts a Topic resource from its proto representation.
func ProtoToTopic(p *alphapb.PubsubAlphaTopic) *alpha.Topic {
	obj := &alpha.Topic{
		Name:                 dcl.StringOrNil(p.GetName()),
		KmsKeyName:           dcl.StringOrNil(p.GetKmsKeyName()),
		MessageStoragePolicy: ProtoToPubsubAlphaTopicMessageStoragePolicy(p.GetMessageStoragePolicy()),
		Project:              dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// TopicMessageStoragePolicyToProto converts a TopicMessageStoragePolicy object to its proto representation.
func PubsubAlphaTopicMessageStoragePolicyToProto(o *alpha.TopicMessageStoragePolicy) *alphapb.PubsubAlphaTopicMessageStoragePolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.PubsubAlphaTopicMessageStoragePolicy{}
	sAllowedPersistenceRegions := make([]string, len(o.AllowedPersistenceRegions))
	for i, r := range o.AllowedPersistenceRegions {
		sAllowedPersistenceRegions[i] = r
	}
	p.SetAllowedPersistenceRegions(sAllowedPersistenceRegions)
	return p
}

// TopicToProto converts a Topic resource to its proto representation.
func TopicToProto(resource *alpha.Topic) *alphapb.PubsubAlphaTopic {
	p := &alphapb.PubsubAlphaTopic{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetKmsKeyName(dcl.ValueOrEmptyString(resource.KmsKeyName))
	p.SetMessageStoragePolicy(PubsubAlphaTopicMessageStoragePolicyToProto(resource.MessageStoragePolicy))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) applyTopic(ctx context.Context, c *alpha.Client, request *alphapb.ApplyPubsubAlphaTopicRequest) (*alphapb.PubsubAlphaTopic, error) {
	p := ProtoToTopic(request.GetResource())
	res, err := c.ApplyTopic(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TopicToProto(res)
	return r, nil
}

// applyPubsubAlphaTopic handles the gRPC request by passing it to the underlying Topic Apply() method.
func (s *TopicServer) ApplyPubsubAlphaTopic(ctx context.Context, request *alphapb.ApplyPubsubAlphaTopicRequest) (*alphapb.PubsubAlphaTopic, error) {
	cl, err := createConfigTopic(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTopic(ctx, cl, request)
}

// DeleteTopic handles the gRPC request by passing it to the underlying Topic Delete() method.
func (s *TopicServer) DeletePubsubAlphaTopic(ctx context.Context, request *alphapb.DeletePubsubAlphaTopicRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTopic(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTopic(ctx, ProtoToTopic(request.GetResource()))

}

// ListPubsubAlphaTopic handles the gRPC request by passing it to the underlying TopicList() method.
func (s *TopicServer) ListPubsubAlphaTopic(ctx context.Context, request *alphapb.ListPubsubAlphaTopicRequest) (*alphapb.ListPubsubAlphaTopicResponse, error) {
	cl, err := createConfigTopic(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTopic(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.PubsubAlphaTopic
	for _, r := range resources.Items {
		rp := TopicToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListPubsubAlphaTopicResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTopic(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
