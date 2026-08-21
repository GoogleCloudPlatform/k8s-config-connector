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

// TriggerServer implements the gRPC interface for Trigger.
type TriggerServer struct{}

// ProtoToTriggerMatchingCriteria converts a TriggerMatchingCriteria object from its proto representation.
func ProtoToEventarcBetaTriggerMatchingCriteria(p *betapb.EventarcBetaTriggerMatchingCriteria) *beta.TriggerMatchingCriteria {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerMatchingCriteria{
		Attribute: dcl.StringOrNil(p.GetAttribute()),
		Value:     dcl.StringOrNil(p.GetValue()),
		Operator:  dcl.StringOrNil(p.GetOperator()),
	}
	return obj
}

// ProtoToTriggerDestination converts a TriggerDestination object from its proto representation.
func ProtoToEventarcBetaTriggerDestination(p *betapb.EventarcBetaTriggerDestination) *beta.TriggerDestination {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerDestination{
		CloudRunService: ProtoToEventarcBetaTriggerDestinationCloudRunService(p.GetCloudRunService()),
		CloudFunction:   dcl.StringOrNil(p.GetCloudFunction()),
		Gke:             ProtoToEventarcBetaTriggerDestinationGke(p.GetGke()),
		Workflow:        dcl.StringOrNil(p.GetWorkflow()),
		HttpEndpoint:    ProtoToEventarcBetaTriggerDestinationHttpEndpoint(p.GetHttpEndpoint()),
		NetworkConfig:   ProtoToEventarcBetaTriggerDestinationNetworkConfig(p.GetNetworkConfig()),
	}
	return obj
}

// ProtoToTriggerDestinationCloudRunService converts a TriggerDestinationCloudRunService object from its proto representation.
func ProtoToEventarcBetaTriggerDestinationCloudRunService(p *betapb.EventarcBetaTriggerDestinationCloudRunService) *beta.TriggerDestinationCloudRunService {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerDestinationCloudRunService{
		Service: dcl.StringOrNil(p.GetService()),
		Path:    dcl.StringOrNil(p.GetPath()),
		Region:  dcl.StringOrNil(p.GetRegion()),
	}
	return obj
}

// ProtoToTriggerDestinationGke converts a TriggerDestinationGke object from its proto representation.
func ProtoToEventarcBetaTriggerDestinationGke(p *betapb.EventarcBetaTriggerDestinationGke) *beta.TriggerDestinationGke {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerDestinationGke{
		Cluster:   dcl.StringOrNil(p.GetCluster()),
		Location:  dcl.StringOrNil(p.GetLocation()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
		Service:   dcl.StringOrNil(p.GetService()),
		Path:      dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToTriggerDestinationHttpEndpoint converts a TriggerDestinationHttpEndpoint object from its proto representation.
func ProtoToEventarcBetaTriggerDestinationHttpEndpoint(p *betapb.EventarcBetaTriggerDestinationHttpEndpoint) *beta.TriggerDestinationHttpEndpoint {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerDestinationHttpEndpoint{
		Uri: dcl.StringOrNil(p.GetUri()),
	}
	return obj
}

// ProtoToTriggerDestinationNetworkConfig converts a TriggerDestinationNetworkConfig object from its proto representation.
func ProtoToEventarcBetaTriggerDestinationNetworkConfig(p *betapb.EventarcBetaTriggerDestinationNetworkConfig) *beta.TriggerDestinationNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerDestinationNetworkConfig{
		NetworkAttachment: dcl.StringOrNil(p.GetNetworkAttachment()),
	}
	return obj
}

// ProtoToTriggerTransport converts a TriggerTransport object from its proto representation.
func ProtoToEventarcBetaTriggerTransport(p *betapb.EventarcBetaTriggerTransport) *beta.TriggerTransport {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerTransport{
		Pubsub: ProtoToEventarcBetaTriggerTransportPubsub(p.GetPubsub()),
	}
	return obj
}

// ProtoToTriggerTransportPubsub converts a TriggerTransportPubsub object from its proto representation.
func ProtoToEventarcBetaTriggerTransportPubsub(p *betapb.EventarcBetaTriggerTransportPubsub) *beta.TriggerTransportPubsub {
	if p == nil {
		return nil
	}
	obj := &beta.TriggerTransportPubsub{
		Topic:        dcl.StringOrNil(p.GetTopic()),
		Subscription: dcl.StringOrNil(p.GetSubscription()),
	}
	return obj
}

// ProtoToTrigger converts a Trigger resource from its proto representation.
func ProtoToTrigger(p *betapb.EventarcBetaTrigger) *beta.Trigger {
	obj := &beta.Trigger{
		Name:                 dcl.StringOrNil(p.GetName()),
		Uid:                  dcl.StringOrNil(p.GetUid()),
		CreateTime:           dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		ServiceAccount:       dcl.StringOrNil(p.GetServiceAccount()),
		Destination:          ProtoToEventarcBetaTriggerDestination(p.GetDestination()),
		Transport:            ProtoToEventarcBetaTriggerTransport(p.GetTransport()),
		Etag:                 dcl.StringOrNil(p.GetEtag()),
		Project:              dcl.StringOrNil(p.GetProject()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		Channel:              dcl.StringOrNil(p.GetChannel()),
		EventDataContentType: dcl.StringOrNil(p.GetEventDataContentType()),
	}
	for _, r := range p.GetMatchingCriteria() {
		obj.MatchingCriteria = append(obj.MatchingCriteria, *ProtoToEventarcBetaTriggerMatchingCriteria(r))
	}
	return obj
}

// TriggerMatchingCriteriaToProto converts a TriggerMatchingCriteria object to its proto representation.
func EventarcBetaTriggerMatchingCriteriaToProto(o *beta.TriggerMatchingCriteria) *betapb.EventarcBetaTriggerMatchingCriteria {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerMatchingCriteria{}
	p.SetAttribute(dcl.ValueOrEmptyString(o.Attribute))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetOperator(dcl.ValueOrEmptyString(o.Operator))
	return p
}

// TriggerDestinationToProto converts a TriggerDestination object to its proto representation.
func EventarcBetaTriggerDestinationToProto(o *beta.TriggerDestination) *betapb.EventarcBetaTriggerDestination {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerDestination{}
	p.SetCloudRunService(EventarcBetaTriggerDestinationCloudRunServiceToProto(o.CloudRunService))
	p.SetCloudFunction(dcl.ValueOrEmptyString(o.CloudFunction))
	p.SetGke(EventarcBetaTriggerDestinationGkeToProto(o.Gke))
	p.SetWorkflow(dcl.ValueOrEmptyString(o.Workflow))
	p.SetHttpEndpoint(EventarcBetaTriggerDestinationHttpEndpointToProto(o.HttpEndpoint))
	p.SetNetworkConfig(EventarcBetaTriggerDestinationNetworkConfigToProto(o.NetworkConfig))
	return p
}

// TriggerDestinationCloudRunServiceToProto converts a TriggerDestinationCloudRunService object to its proto representation.
func EventarcBetaTriggerDestinationCloudRunServiceToProto(o *beta.TriggerDestinationCloudRunService) *betapb.EventarcBetaTriggerDestinationCloudRunService {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerDestinationCloudRunService{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetRegion(dcl.ValueOrEmptyString(o.Region))
	return p
}

// TriggerDestinationGkeToProto converts a TriggerDestinationGke object to its proto representation.
func EventarcBetaTriggerDestinationGkeToProto(o *beta.TriggerDestinationGke) *betapb.EventarcBetaTriggerDestinationGke {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerDestinationGke{}
	p.SetCluster(dcl.ValueOrEmptyString(o.Cluster))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// TriggerDestinationHttpEndpointToProto converts a TriggerDestinationHttpEndpoint object to its proto representation.
func EventarcBetaTriggerDestinationHttpEndpointToProto(o *beta.TriggerDestinationHttpEndpoint) *betapb.EventarcBetaTriggerDestinationHttpEndpoint {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerDestinationHttpEndpoint{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	return p
}

// TriggerDestinationNetworkConfigToProto converts a TriggerDestinationNetworkConfig object to its proto representation.
func EventarcBetaTriggerDestinationNetworkConfigToProto(o *beta.TriggerDestinationNetworkConfig) *betapb.EventarcBetaTriggerDestinationNetworkConfig {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerDestinationNetworkConfig{}
	p.SetNetworkAttachment(dcl.ValueOrEmptyString(o.NetworkAttachment))
	return p
}

// TriggerTransportToProto converts a TriggerTransport object to its proto representation.
func EventarcBetaTriggerTransportToProto(o *beta.TriggerTransport) *betapb.EventarcBetaTriggerTransport {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerTransport{}
	p.SetPubsub(EventarcBetaTriggerTransportPubsubToProto(o.Pubsub))
	return p
}

// TriggerTransportPubsubToProto converts a TriggerTransportPubsub object to its proto representation.
func EventarcBetaTriggerTransportPubsubToProto(o *beta.TriggerTransportPubsub) *betapb.EventarcBetaTriggerTransportPubsub {
	if o == nil {
		return nil
	}
	p := &betapb.EventarcBetaTriggerTransportPubsub{}
	p.SetTopic(dcl.ValueOrEmptyString(o.Topic))
	p.SetSubscription(dcl.ValueOrEmptyString(o.Subscription))
	return p
}

// TriggerToProto converts a Trigger resource to its proto representation.
func TriggerToProto(resource *beta.Trigger) *betapb.EventarcBetaTrigger {
	p := &betapb.EventarcBetaTrigger{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetServiceAccount(dcl.ValueOrEmptyString(resource.ServiceAccount))
	p.SetDestination(EventarcBetaTriggerDestinationToProto(resource.Destination))
	p.SetTransport(EventarcBetaTriggerTransportToProto(resource.Transport))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetChannel(dcl.ValueOrEmptyString(resource.Channel))
	p.SetEventDataContentType(dcl.ValueOrEmptyString(resource.EventDataContentType))
	sMatchingCriteria := make([]*betapb.EventarcBetaTriggerMatchingCriteria, len(resource.MatchingCriteria))
	for i, r := range resource.MatchingCriteria {
		sMatchingCriteria[i] = EventarcBetaTriggerMatchingCriteriaToProto(&r)
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
func (s *TriggerServer) applyTrigger(ctx context.Context, c *beta.Client, request *betapb.ApplyEventarcBetaTriggerRequest) (*betapb.EventarcBetaTrigger, error) {
	p := ProtoToTrigger(request.GetResource())
	res, err := c.ApplyTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TriggerToProto(res)
	return r, nil
}

// applyEventarcBetaTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) ApplyEventarcBetaTrigger(ctx context.Context, request *betapb.ApplyEventarcBetaTriggerRequest) (*betapb.EventarcBetaTrigger, error) {
	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTrigger(ctx, cl, request)
}

// DeleteTrigger handles the gRPC request by passing it to the underlying Trigger Delete() method.
func (s *TriggerServer) DeleteEventarcBetaTrigger(ctx context.Context, request *betapb.DeleteEventarcBetaTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTrigger(ctx, ProtoToTrigger(request.GetResource()))

}

// ListEventarcBetaTrigger handles the gRPC request by passing it to the underlying TriggerList() method.
func (s *TriggerServer) ListEventarcBetaTrigger(ctx context.Context, request *betapb.ListEventarcBetaTriggerRequest) (*betapb.ListEventarcBetaTriggerResponse, error) {
	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTrigger(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.EventarcBetaTrigger
	for _, r := range resources.Items {
		rp := TriggerToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListEventarcBetaTriggerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTrigger(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
