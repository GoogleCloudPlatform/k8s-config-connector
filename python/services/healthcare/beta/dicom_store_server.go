// Copyright 2022 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/healthcare/beta/healthcare_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/healthcare/beta"
)

// DicomStoreServer implements the gRPC interface for DicomStore.
type DicomStoreServer struct{}

// ProtoToDicomStoreNotificationConfig converts a DicomStoreNotificationConfig object from its proto representation.
func ProtoToHealthcareBetaDicomStoreNotificationConfig(p *betapb.HealthcareBetaDicomStoreNotificationConfig) *beta.DicomStoreNotificationConfig {
	if p == nil {
		return nil
	}
	obj := &beta.DicomStoreNotificationConfig{
		PubsubTopic: dcl.StringOrNil(p.GetPubsubTopic()),
	}
	return obj
}

// ProtoToDicomStore converts a DicomStore resource from its proto representation.
func ProtoToDicomStore(p *betapb.HealthcareBetaDicomStore) *beta.DicomStore {
	obj := &beta.DicomStore{
		Name:               dcl.StringOrNil(p.GetName()),
		NotificationConfig: ProtoToHealthcareBetaDicomStoreNotificationConfig(p.GetNotificationConfig()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		Dataset:            dcl.StringOrNil(p.GetDataset()),
	}
	return obj
}

// DicomStoreNotificationConfigToProto converts a DicomStoreNotificationConfig object to its proto representation.
func HealthcareBetaDicomStoreNotificationConfigToProto(o *beta.DicomStoreNotificationConfig) *betapb.HealthcareBetaDicomStoreNotificationConfig {
	if o == nil {
		return nil
	}
	p := &betapb.HealthcareBetaDicomStoreNotificationConfig{}
	p.SetPubsubTopic(dcl.ValueOrEmptyString(o.PubsubTopic))
	return p
}

// DicomStoreToProto converts a DicomStore resource to its proto representation.
func DicomStoreToProto(resource *beta.DicomStore) *betapb.HealthcareBetaDicomStore {
	p := &betapb.HealthcareBetaDicomStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNotificationConfig(HealthcareBetaDicomStoreNotificationConfigToProto(resource.NotificationConfig))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetDataset(dcl.ValueOrEmptyString(resource.Dataset))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyDicomStore handles the gRPC request by passing it to the underlying DicomStore Apply() method.
func (s *DicomStoreServer) applyDicomStore(ctx context.Context, c *beta.Client, request *betapb.ApplyHealthcareBetaDicomStoreRequest) (*betapb.HealthcareBetaDicomStore, error) {
	p := ProtoToDicomStore(request.GetResource())
	res, err := c.ApplyDicomStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DicomStoreToProto(res)
	return r, nil
}

// applyHealthcareBetaDicomStore handles the gRPC request by passing it to the underlying DicomStore Apply() method.
func (s *DicomStoreServer) ApplyHealthcareBetaDicomStore(ctx context.Context, request *betapb.ApplyHealthcareBetaDicomStoreRequest) (*betapb.HealthcareBetaDicomStore, error) {
	cl, err := createConfigDicomStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDicomStore(ctx, cl, request)
}

// DeleteDicomStore handles the gRPC request by passing it to the underlying DicomStore Delete() method.
func (s *DicomStoreServer) DeleteHealthcareBetaDicomStore(ctx context.Context, request *betapb.DeleteHealthcareBetaDicomStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDicomStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDicomStore(ctx, ProtoToDicomStore(request.GetResource()))

}

// ListHealthcareBetaDicomStore handles the gRPC request by passing it to the underlying DicomStoreList() method.
func (s *DicomStoreServer) ListHealthcareBetaDicomStore(ctx context.Context, request *betapb.ListHealthcareBetaDicomStoreRequest) (*betapb.ListHealthcareBetaDicomStoreResponse, error) {
	cl, err := createConfigDicomStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDicomStore(ctx, request.GetProject(), request.GetLocation(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.HealthcareBetaDicomStore
	for _, r := range resources.Items {
		rp := DicomStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListHealthcareBetaDicomStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDicomStore(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
