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

// TriggerServer implements the gRPC interface for Trigger.
type TriggerServer struct{}

// ProtoToTriggerMatchingCriteria converts a TriggerMatchingCriteria object from its proto representation.
func ProtoToEventarcAlphaTriggerMatchingCriteria(p *alphapb.EventarcAlphaTriggerMatchingCriteria) *alpha.TriggerMatchingCriteria {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerMatchingCriteria{
		Attribute: dcl.StringOrNil(p.GetAttribute()),
		Value:     dcl.StringOrNil(p.GetValue()),
		Operator:  dcl.StringOrNil(p.GetOperator()),
	}
	return obj
}

// ProtoToTriggerDestination converts a TriggerDestination object from its proto representation.
func ProtoToEventarcAlphaTriggerDestination(p *alphapb.EventarcAlphaTriggerDestination) *alpha.TriggerDestination {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerDestination{
		CloudRunService: ProtoToEventarcAlphaTriggerDestinationCloudRunService(p.GetCloudRunService()),
		CloudFunction:   dcl.StringOrNil(p.GetCloudFunction()),
		Gke:             ProtoToEventarcAlphaTriggerDestinationGke(p.GetGke()),
		Workflow:        dcl.StringOrNil(p.GetWorkflow()),
		HttpEndpoint:    ProtoToEventarcAlphaTriggerDestinationHttpEndpoint(p.GetHttpEndpoint()),
		NetworkConfig:   ProtoToEventarcAlphaTriggerDestinationNetworkConfig(p.GetNetworkConfig()),
	}
	return obj
}

// ProtoToTriggerDestinationCloudRunService converts a TriggerDestinationCloudRunService object from its proto representation.
func ProtoToEventarcAlphaTriggerDestinationCloudRunService(p *alphapb.EventarcAlphaTriggerDestinationCloudRunService) *alpha.TriggerDestinationCloudRunService {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerDestinationCloudRunService{
		Service: dcl.StringOrNil(p.GetService()),
		Path:    dcl.StringOrNil(p.GetPath()),
		Region:  dcl.StringOrNil(p.GetRegion()),
	}
	return obj
}

// ProtoToTriggerDestinationGke converts a TriggerDestinationGke object from its proto representation.
func ProtoToEventarcAlphaTriggerDestinationGke(p *alphapb.EventarcAlphaTriggerDestinationGke) *alpha.TriggerDestinationGke {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerDestinationGke{
		Cluster:   dcl.StringOrNil(p.GetCluster()),
		Location:  dcl.StringOrNil(p.GetLocation()),
		Namespace: dcl.StringOrNil(p.GetNamespace()),
		Service:   dcl.StringOrNil(p.GetService()),
		Path:      dcl.StringOrNil(p.GetPath()),
	}
	return obj
}

// ProtoToTriggerDestinationHttpEndpoint converts a TriggerDestinationHttpEndpoint object from its proto representation.
func ProtoToEventarcAlphaTriggerDestinationHttpEndpoint(p *alphapb.EventarcAlphaTriggerDestinationHttpEndpoint) *alpha.TriggerDestinationHttpEndpoint {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerDestinationHttpEndpoint{
		Uri: dcl.StringOrNil(p.GetUri()),
	}
	return obj
}

// ProtoToTriggerDestinationNetworkConfig converts a TriggerDestinationNetworkConfig object from its proto representation.
func ProtoToEventarcAlphaTriggerDestinationNetworkConfig(p *alphapb.EventarcAlphaTriggerDestinationNetworkConfig) *alpha.TriggerDestinationNetworkConfig {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerDestinationNetworkConfig{
		NetworkAttachment: dcl.StringOrNil(p.GetNetworkAttachment()),
	}
	return obj
}

// ProtoToTriggerTransport converts a TriggerTransport object from its proto representation.
func ProtoToEventarcAlphaTriggerTransport(p *alphapb.EventarcAlphaTriggerTransport) *alpha.TriggerTransport {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerTransport{
		Pubsub: ProtoToEventarcAlphaTriggerTransportPubsub(p.GetPubsub()),
	}
	return obj
}

// ProtoToTriggerTransportPubsub converts a TriggerTransportPubsub object from its proto representation.
func ProtoToEventarcAlphaTriggerTransportPubsub(p *alphapb.EventarcAlphaTriggerTransportPubsub) *alpha.TriggerTransportPubsub {
	if p == nil {
		return nil
	}
	obj := &alpha.TriggerTransportPubsub{
		Topic:        dcl.StringOrNil(p.GetTopic()),
		Subscription: dcl.StringOrNil(p.GetSubscription()),
	}
	return obj
}

// ProtoToTrigger converts a Trigger resource from its proto representation.
func ProtoToTrigger(p *alphapb.EventarcAlphaTrigger) *alpha.Trigger {
	obj := &alpha.Trigger{
		Name:                 dcl.StringOrNil(p.GetName()),
		Uid:                  dcl.StringOrNil(p.GetUid()),
		CreateTime:           dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:           dcl.StringOrNil(p.GetUpdateTime()),
		ServiceAccount:       dcl.StringOrNil(p.GetServiceAccount()),
		Destination:          ProtoToEventarcAlphaTriggerDestination(p.GetDestination()),
		Transport:            ProtoToEventarcAlphaTriggerTransport(p.GetTransport()),
		Etag:                 dcl.StringOrNil(p.GetEtag()),
		Project:              dcl.StringOrNil(p.GetProject()),
		Location:             dcl.StringOrNil(p.GetLocation()),
		Channel:              dcl.StringOrNil(p.GetChannel()),
		EventDataContentType: dcl.StringOrNil(p.GetEventDataContentType()),
	}
	for _, r := range p.GetMatchingCriteria() {
		obj.MatchingCriteria = append(obj.MatchingCriteria, *ProtoToEventarcAlphaTriggerMatchingCriteria(r))
	}
	return obj
}

// TriggerMatchingCriteriaToProto converts a TriggerMatchingCriteria object to its proto representation.
func EventarcAlphaTriggerMatchingCriteriaToProto(o *alpha.TriggerMatchingCriteria) *alphapb.EventarcAlphaTriggerMatchingCriteria {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerMatchingCriteria{}
	p.SetAttribute(dcl.ValueOrEmptyString(o.Attribute))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	p.SetOperator(dcl.ValueOrEmptyString(o.Operator))
	return p
}

// TriggerDestinationToProto converts a TriggerDestination object to its proto representation.
func EventarcAlphaTriggerDestinationToProto(o *alpha.TriggerDestination) *alphapb.EventarcAlphaTriggerDestination {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerDestination{}
	p.SetCloudRunService(EventarcAlphaTriggerDestinationCloudRunServiceToProto(o.CloudRunService))
	p.SetCloudFunction(dcl.ValueOrEmptyString(o.CloudFunction))
	p.SetGke(EventarcAlphaTriggerDestinationGkeToProto(o.Gke))
	p.SetWorkflow(dcl.ValueOrEmptyString(o.Workflow))
	p.SetHttpEndpoint(EventarcAlphaTriggerDestinationHttpEndpointToProto(o.HttpEndpoint))
	p.SetNetworkConfig(EventarcAlphaTriggerDestinationNetworkConfigToProto(o.NetworkConfig))
	return p
}

// TriggerDestinationCloudRunServiceToProto converts a TriggerDestinationCloudRunService object to its proto representation.
func EventarcAlphaTriggerDestinationCloudRunServiceToProto(o *alpha.TriggerDestinationCloudRunService) *alphapb.EventarcAlphaTriggerDestinationCloudRunService {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerDestinationCloudRunService{}
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	p.SetRegion(dcl.ValueOrEmptyString(o.Region))
	return p
}

// TriggerDestinationGkeToProto converts a TriggerDestinationGke object to its proto representation.
func EventarcAlphaTriggerDestinationGkeToProto(o *alpha.TriggerDestinationGke) *alphapb.EventarcAlphaTriggerDestinationGke {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerDestinationGke{}
	p.SetCluster(dcl.ValueOrEmptyString(o.Cluster))
	p.SetLocation(dcl.ValueOrEmptyString(o.Location))
	p.SetNamespace(dcl.ValueOrEmptyString(o.Namespace))
	p.SetService(dcl.ValueOrEmptyString(o.Service))
	p.SetPath(dcl.ValueOrEmptyString(o.Path))
	return p
}

// TriggerDestinationHttpEndpointToProto converts a TriggerDestinationHttpEndpoint object to its proto representation.
func EventarcAlphaTriggerDestinationHttpEndpointToProto(o *alpha.TriggerDestinationHttpEndpoint) *alphapb.EventarcAlphaTriggerDestinationHttpEndpoint {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerDestinationHttpEndpoint{}
	p.SetUri(dcl.ValueOrEmptyString(o.Uri))
	return p
}

// TriggerDestinationNetworkConfigToProto converts a TriggerDestinationNetworkConfig object to its proto representation.
func EventarcAlphaTriggerDestinationNetworkConfigToProto(o *alpha.TriggerDestinationNetworkConfig) *alphapb.EventarcAlphaTriggerDestinationNetworkConfig {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerDestinationNetworkConfig{}
	p.SetNetworkAttachment(dcl.ValueOrEmptyString(o.NetworkAttachment))
	return p
}

// TriggerTransportToProto converts a TriggerTransport object to its proto representation.
func EventarcAlphaTriggerTransportToProto(o *alpha.TriggerTransport) *alphapb.EventarcAlphaTriggerTransport {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerTransport{}
	p.SetPubsub(EventarcAlphaTriggerTransportPubsubToProto(o.Pubsub))
	return p
}

// TriggerTransportPubsubToProto converts a TriggerTransportPubsub object to its proto representation.
func EventarcAlphaTriggerTransportPubsubToProto(o *alpha.TriggerTransportPubsub) *alphapb.EventarcAlphaTriggerTransportPubsub {
	if o == nil {
		return nil
	}
	p := &alphapb.EventarcAlphaTriggerTransportPubsub{}
	p.SetTopic(dcl.ValueOrEmptyString(o.Topic))
	p.SetSubscription(dcl.ValueOrEmptyString(o.Subscription))
	return p
}

// TriggerToProto converts a Trigger resource to its proto representation.
func TriggerToProto(resource *alpha.Trigger) *alphapb.EventarcAlphaTrigger {
	p := &alphapb.EventarcAlphaTrigger{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetUid(dcl.ValueOrEmptyString(resource.Uid))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetServiceAccount(dcl.ValueOrEmptyString(resource.ServiceAccount))
	p.SetDestination(EventarcAlphaTriggerDestinationToProto(resource.Destination))
	p.SetTransport(EventarcAlphaTriggerTransportToProto(resource.Transport))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetChannel(dcl.ValueOrEmptyString(resource.Channel))
	p.SetEventDataContentType(dcl.ValueOrEmptyString(resource.EventDataContentType))
	sMatchingCriteria := make([]*alphapb.EventarcAlphaTriggerMatchingCriteria, len(resource.MatchingCriteria))
	for i, r := range resource.MatchingCriteria {
		sMatchingCriteria[i] = EventarcAlphaTriggerMatchingCriteriaToProto(&r)
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
func (s *TriggerServer) applyTrigger(ctx context.Context, c *alpha.Client, request *alphapb.ApplyEventarcAlphaTriggerRequest) (*alphapb.EventarcAlphaTrigger, error) {
	p := ProtoToTrigger(request.GetResource())
	res, err := c.ApplyTrigger(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TriggerToProto(res)
	return r, nil
}

// applyEventarcAlphaTrigger handles the gRPC request by passing it to the underlying Trigger Apply() method.
func (s *TriggerServer) ApplyEventarcAlphaTrigger(ctx context.Context, request *alphapb.ApplyEventarcAlphaTriggerRequest) (*alphapb.EventarcAlphaTrigger, error) {
	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyTrigger(ctx, cl, request)
}

// DeleteTrigger handles the gRPC request by passing it to the underlying Trigger Delete() method.
func (s *TriggerServer) DeleteEventarcAlphaTrigger(ctx context.Context, request *alphapb.DeleteEventarcAlphaTriggerRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTrigger(ctx, ProtoToTrigger(request.GetResource()))

}

// ListEventarcAlphaTrigger handles the gRPC request by passing it to the underlying TriggerList() method.
func (s *TriggerServer) ListEventarcAlphaTrigger(ctx context.Context, request *alphapb.ListEventarcAlphaTriggerRequest) (*alphapb.ListEventarcAlphaTriggerResponse, error) {
	cl, err := createConfigTrigger(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTrigger(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.EventarcAlphaTrigger
	for _, r := range resources.Items {
		rp := TriggerToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListEventarcAlphaTriggerResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigTrigger(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
