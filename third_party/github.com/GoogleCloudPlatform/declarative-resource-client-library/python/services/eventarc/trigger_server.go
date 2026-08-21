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

// TriggerServer implements the gRPC interface for Trigger.
type TriggerServer struct{}

// ProtoToTriggerMatchingCriteria converts a TriggerMatchingCriteria object from its proto representation.
func ProtoToEventarcTriggerMatchingCriteria(p *eventarcpb.EventarcTriggerMatchingCriteria) *eventarc.TriggerMatchingCriteria {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerMatchingCriteria{
		Attribute: dcl.StringOrNil(p.GetAttribute()),
		Value:     dcl.StringOrNil(p.GetValue()),
		Operator:  dcl.StringOrNil(p.GetOperator()),
	}
	return obj
}

// ProtoToTriggerDestination converts a TriggerDestination object from its proto representation.
func ProtoToEventarcTriggerDestination(p *eventarcpb.EventarcTriggerDestination) *eventarc.TriggerDestination {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerDestination{
		CloudRunService: ProtoToEventarcTriggerDestinationCloudRunService(p.GetCloudRunService()),
		CloudFunction:   dcl.StringOrNil(p.GetCloudFunction()),
		Gke:             ProtoToEventarcTriggerDestinationGke(p.GetGke()),
		Workflow:        dcl.StringOrNil(p.GetWorkflow()),
		HttpEndpoint:    ProtoToEventarcTriggerDestinationHttpEndpoint(p.GetHttpEndpoint()),
		NetworkConfig:   ProtoToEventarcTriggerDestinationNetworkConfig(p.GetNetworkConfig()),
	}
	return obj
}

// ProtoToTriggerDestinationCloudRunService converts a TriggerDestinationCloudRunService object from its proto representation.
func ProtoToEventarcTriggerDestinationCloudRunService(p *eventarcpb.EventarcTriggerDestinationCloudRunService) *eventarc.TriggerDestinationCloudRunService {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerDestinationCloudRunService{
		Service: dcl.StringOrNil(p.GetService()),
		Path:    dcl.StringOrNil(p.GetPath()),
		Region:  dcl.StringOrNil(p.GetRegion()),
	}
	return obj
}

// ProtoToTriggerDestinationGke converts a TriggerDestinationGke object from its proto representation.
func ProtoToEventarcTriggerDestinationGke(p *eventarcpb.EventarcTriggerDestinationGke) *eventarc.TriggerDestinationGke {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerDestinationGke{
		Cluster:   dcl.StringOrNil(p.GetCluster()),
		Location:  dcl.StringOrNil(p.GetLocation()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
		Service:   dcl.StringOrNil(p.GetService()),
		Path:      dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToTriggerDestinationHttpEndpoint converts a TriggerDestinationHttpEndpoint object from its proto representation.
func ProtoToEventarcTriggerDestinationHttpEndpoint(p *eventarcpb.EventarcTriggerDestinationHttpEndpoint) *eventarc.TriggerDestinationHttpEndpoint {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerDestinationHttpEndpoint{
		Uri: dcl.StringOrNil(p.GetUri()),
	}
	return obj
}

// ProtoToTriggerDestinationNetworkConfig converts a TriggerDestinationNetworkConfig object from its proto representation.
func ProtoToEventarcTriggerDestinationNetworkConfig(p *eventarcpb.EventarcTriggerDestinationNetworkConfig) *eventarc.TriggerDestinationNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerDestinationNetworkConfig{
		NetworkAttachment: dcl.StringOrNil(p.GetNetworkAttachment()),
	}
	return obj
}

// ProtoToTriggerTransport converts a TriggerTransport object from its proto representation.
func ProtoToEventarcTriggerTransport(p *eventarcpb.EventarcTriggerTransport) *eventarc.TriggerTransport {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerTransport{
		Pubsub: ProtoToEventarcTriggerTransportPubsub(p.GetPubsub()),
	}
	return obj
}

// ProtoToTriggerTransportPubsub converts a TriggerTransportPubsub object from its proto representation.
func ProtoToEventarcTriggerTransportPubsub(p *eventarcpb.EventarcTriggerTransportPubsub) *eventarc.TriggerTransportPubsub {
	if p == nil {
		return nil
	}
	obj := &eventarc.TriggerTransportPubsub{
		Topic:        dcl.StringOrNil(p.GetTopic()),
		Subscription: dcl.StringOrNil(p.GetSubscription()),
	}
	return obj
}

// ProtoToTrigger converts a Trigger resource from its proto representation.
func ProtoToTrigger(p *eventarcpb.EventarcTrigger) *eventarc.Trigger {
	obj := &eventarc.Trigger{
		Name:                 dcl.StringOrNil(p.GetName()),
		Uid:                  dcl.StringOrNil(p.GetUid()),
		CreateTime:           dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		ServiceAccount:       dcl.StringOrNil(p.GetServiceAccount()),
		Destination:          ProtoToEventarcTriggerDestination(p.GetDestination()),
		Transport:            ProtoToEventarcTriggerTransport(p.GetTransport()),
		Etag:                 dcl.StringOrNil(p.GetEtag()),
		Project:              dcl.StringOrNil(p.GetProject()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		Channel:              dcl.StringOrNil(p.GetChannel()),
		EventDataContentType: dcl.StringOrNil(p.GetEventDataContentType()),
	}
	for _, r := range p.GetMatchingCriteria() {
		obj.MatchingCriteria = append(obj.MatchingCriteria, *ProtoToEventarcTriggerMatchingCriteria(r))
	}
	return obj
}

// TriggerMatchingCriteriaToProto converts a TriggerMatchingCriteria object to its proto representation.
func EventarcTriggerMatchingCriteriaToProto(o *eventarc.TriggerMatchingCriteria) *eventarcpb.EventarcTriggerMatchingCriteria {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerMatchingCriteria{}
	p.SetAttribute(dcl.ValueOrEmptyString(o.Attribute))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetOperator(dcl.ValueOrEmptyString(o.Operator))
	return p
}

// TriggerDestinationToProto converts a TriggerDestination object to its proto representation.
func EventarcTriggerDestinationToProto(o *eventarc.TriggerDestination) *eventarcpb.EventarcTriggerDestination {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerDestination{}
	p.SetCloudRunService(EventarcTriggerDestinationCloudRunServiceToProto(o.CloudRunService))
	p.SetCloudFunction(dcl.ValueOrEmptyString(o.CloudFunction))
	p.SetGke(EventarcTriggerDestinationGkeToProto(o.Gke))
	p.SetWorkflow(dcl.ValueOrEmptyString(o.Workflow))
	p.SetHttpEndpoint(EventarcTriggerDestinationHttpEndpointToProto(o.HttpEndpoint))
	p.SetNetworkConfig(EventarcTriggerDestinationNetworkConfigToProto(o.NetworkConfig))
	return p
}

// TriggerDestinationCloudRunServiceToProto converts a TriggerDestinationCloudRunService object to its proto representation.
func EventarcTriggerDestinationCloudRunServiceToProto(o *eventarc.TriggerDestinationCloudRunService) *eventarcpb.EventarcTriggerDestinationCloudRunService {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerDestinationCloudRunService{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetRegion(dcl.ValueOrEmptyString(o.Region))
	return p
}

// TriggerDestinationGkeToProto converts a TriggerDestinationGke object to its proto representation.
func EventarcTriggerDestinationGkeToProto(o *eventarc.TriggerDestinationGke) *eventarcpb.EventarcTriggerDestinationGke {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerDestinationGke{}
	p.SetCluster(dcl.ValueOrEmptyString(o.Cluster))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// TriggerDestinationHttpEndpointToProto converts a TriggerDestinationHttpEndpoint object to its proto representation.
func EventarcTriggerDestinationHttpEndpointToProto(o *eventarc.TriggerDestinationHttpEndpoint) *eventarcpb.EventarcTriggerDestinationHttpEndpoint {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerDestinationHttpEndpoint{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	return p
}

// TriggerDestinationNetworkConfigToProto converts a TriggerDestinationNetworkConfig object to its proto representation.
func EventarcTriggerDestinationNetworkConfigToProto(o *eventarc.TriggerDestinationNetworkConfig) *eventarcpb.EventarcTriggerDestinationNetworkConfig {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerDestinationNetworkConfig{}
	p.SetNetworkAttachment(dcl.ValueOrEmptyString(o.NetworkAttachment))
	return p
}

// TriggerTransportToProto converts a TriggerTransport object to its proto representation.
func EventarcTriggerTransportToProto(o *eventarc.TriggerTransport) *eventarcpb.EventarcTriggerTransport {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerTransport{}
	p.SetPubsub(EventarcTriggerTransportPubsubToProto(o.Pubsub))
	return p
}

// TriggerTransportPubsubToProto converts a TriggerTransportPubsub object to its proto representation.
func EventarcTriggerTransportPubsubToProto(o *eventarc.TriggerTransportPubsub) *eventarcpb.EventarcTriggerTransportPubsub {
	if o == nil {
		return nil
	}
	p := &eventarcpb.EventarcTriggerTransportPubsub{}
	p.SetTopic(dcl.ValueOrEmptyString(o.Topic))
	p.SetSubscription(dcl.ValueOrEmptyString(o.Subscription))
	return p
}

// TriggerToProto converts a Trigger resource to its proto representation.
func TriggerToProto(resource *eventarc.Trigger) *eventarcpb.EventarcTrigger {
	p := &eventarcpb.EventarcTrigger{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetServiceAccount(dcl.ValueOrEmptyString(resource.ServiceAccount))
	p.SetDestination(EventarcTriggerDestinationToProto(resource.Destination))
	p.SetTransport(EventarcTriggerTransportToProto(resource.Transport))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetChannel(dcl.ValueOrEmptyString(resource.Channel))
	p.SetEventDataContentType(dcl.ValueOrEmptyString(resource.EventDataContentType))
	sMatchingCriteria := make([]*eventarcpb.EventarcTriggerMatchingCriteria, len(resource.MatchingCriteria))
	for i, r := range resource.MatchingCriteria {
		sMatchingCriteria[i] = EventarcTriggerMatchingCriteriaToProto(&r)
	}
	p.SetMatchingCriteria(sMatchingCriteria)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	mConditions := make(map[string]string, len(resource.Conditions))
	for k, r := range resource.Conditions {
		mConditions[k] = r
	}
	p.SetConditions(mConditions)

	return p
}

// applyTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) applyTrigger(ctx context.Context, c *eventarc.Client, request *eventarcpb.ApplyEventarcTriggerRequest) (*eventarcpb.EventarcTrigger, error) {
	p := ProtoToTrigger(request.GetResource())
	res, err := c.ApplyTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TriggerToProto(res)
	return r, nil
}

// applyEventarcTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) ApplyEventarcTrigger(ctx context.Context, request *eventarcpb.ApplyEventarcTriggerRequest) (*eventarcpb.EventarcTrigger, error) {
	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTrigger(ctx, cl, request)
}

// DeleteTrigger handles the gRPC request by passing it to the underlying Trigger Delete() method.
func (s *TriggerServer) DeleteEventarcTrigger(ctx context.Context, request *eventarcpb.DeleteEventarcTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTrigger(ctx, ProtoToTrigger(request.GetResource()))

}

// ListEventarcTrigger handles the gRPC request by passing it to the underlying TriggerList() method.
func (s *TriggerServer) ListEventarcTrigger(ctx context.Context, request *eventarcpb.ListEventarcTriggerRequest) (*eventarcpb.ListEventarcTriggerResponse, error) {
	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTrigger(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*eventarcpb.EventarcTrigger
	for _, r := range resources.Items {
		rp := TriggerToProto(r)
		protos = append(protos, rp)
	}
	p := &eventarcpb.ListEventarcTriggerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTrigger(ctx context.Context, service_account_file string) (*eventarc.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return eventarc.NewClient(conf), nil
}
