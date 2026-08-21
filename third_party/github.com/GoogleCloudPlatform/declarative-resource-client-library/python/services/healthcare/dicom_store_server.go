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
	healthcarepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/healthcare/healthcare_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/healthcare"
)

// DicomStoreServer implements the gRPC interface for DicomStore.
type DicomStoreServer struct{}

// ProtoToDicomStoreNotificationConfig converts a DicomStoreNotificationConfig object from its proto representation.
func ProtoToHealthcareDicomStoreNotificationConfig(p *healthcarepb.HealthcareDicomStoreNotificationConfig) *healthcare.DicomStoreNotificationConfig {
	if p == nil {
		return nil
	}
	obj := &healthcare.DicomStoreNotificationConfig{
		PubsubTopic: dcl.StringOrNil(p.GetPubsubTopic()),
	}
	return obj
}

// ProtoToDicomStore converts a DicomStore resource from its proto representation.
func ProtoToDicomStore(p *healthcarepb.HealthcareDicomStore) *healthcare.DicomStore {
	obj := &healthcare.DicomStore{
		Name:               dcl.StringOrNil(p.GetName()),
		NotificationConfig: ProtoToHealthcareDicomStoreNotificationConfig(p.GetNotificationConfig()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
		Dataset:            dcl.StringOrNil(p.GetDataset()),
	}
	return obj
}

// DicomStoreNotificationConfigToProto converts a DicomStoreNotificationConfig object to its proto representation.
func HealthcareDicomStoreNotificationConfigToProto(o *healthcare.DicomStoreNotificationConfig) *healthcarepb.HealthcareDicomStoreNotificationConfig {
	if o == nil {
		return nil
	}
	p := &healthcarepb.HealthcareDicomStoreNotificationConfig{}
	p.SetPubsubTopic(dcl.ValueOrEmptyString(o.PubsubTopic))
	return p
}

// DicomStoreToProto converts a DicomStore resource to its proto representation.
func DicomStoreToProto(resource *healthcare.DicomStore) *healthcarepb.HealthcareDicomStore {
	p := &healthcarepb.HealthcareDicomStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetNotificationConfig(HealthcareDicomStoreNotificationConfigToProto(resource.NotificationConfig))
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
func (s *DicomStoreServer) applyDicomStore(ctx context.Context, c *healthcare.Client, request *healthcarepb.ApplyHealthcareDicomStoreRequest) (*healthcarepb.HealthcareDicomStore, error) {
	p := ProtoToDicomStore(request.GetResource())
	res, err := c.ApplyDicomStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DicomStoreToProto(res)
	return r, nil
}

// applyHealthcareDicomStore handles the gRPC request by passing it to the underlying DicomStore Apply() method.
func (s *DicomStoreServer) ApplyHealthcareDicomStore(ctx context.Context, request *healthcarepb.ApplyHealthcareDicomStoreRequest) (*healthcarepb.HealthcareDicomStore, error) {
	cl, err := createConfigDicomStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDicomStore(ctx, cl, request)
}

// DeleteDicomStore handles the gRPC request by passing it to the underlying DicomStore Delete() method.
func (s *DicomStoreServer) DeleteHealthcareDicomStore(ctx context.Context, request *healthcarepb.DeleteHealthcareDicomStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDicomStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDicomStore(ctx, ProtoToDicomStore(request.GetResource()))

}

// ListHealthcareDicomStore handles the gRPC request by passing it to the underlying DicomStoreList() method.
func (s *DicomStoreServer) ListHealthcareDicomStore(ctx context.Context, request *healthcarepb.ListHealthcareDicomStoreRequest) (*healthcarepb.ListHealthcareDicomStoreResponse, error) {
	cl, err := createConfigDicomStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDicomStore(ctx, request.GetProject(), request.GetLocation(), request.GetDataset())
	if err != nil {
		return nil, err
	}
	var protos []*healthcarepb.HealthcareDicomStore
	for _, r := range resources.Items {
		rp := DicomStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &healthcarepb.ListHealthcareDicomStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDicomStore(ctx context.Context, service_account_file string) (*healthcare.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return healthcare.NewClient(conf), nil
}
