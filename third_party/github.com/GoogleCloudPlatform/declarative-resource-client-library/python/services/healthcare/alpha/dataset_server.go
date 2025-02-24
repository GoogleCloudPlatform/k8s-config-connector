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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/healthcare/alpha/healthcare_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/healthcare/alpha"
)

// DatasetServer implements the gRPC interface for Dataset.
type DatasetServer struct{}

// ProtoToDataset converts a Dataset resource from its proto representation.
func ProtoToDataset(p *alphapb.HealthcareAlphaDataset) *alpha.Dataset {
	obj := &alpha.Dataset{
		Name:     dcl.StringOrNil(p.GetName()),
		TimeZone: dcl.StringOrNil(p.GetTimeZone()),
		Project:  dcl.StringOrNil(p.GetProject()),
		Location: dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// DatasetToProto converts a Dataset resource to its proto representation.
func DatasetToProto(resource *alpha.Dataset) *alphapb.HealthcareAlphaDataset {
	p := &alphapb.HealthcareAlphaDataset{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTimeZone(dcl.ValueOrEmptyString(resource.TimeZone))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) applyDataset(ctx context.Context, c *alpha.Client, request *alphapb.ApplyHealthcareAlphaDatasetRequest) (*alphapb.HealthcareAlphaDataset, error) {
	p := ProtoToDataset(request.GetResource())
	res, err := c.ApplyDataset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DatasetToProto(res)
	return r, nil
}

// applyHealthcareAlphaDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) ApplyHealthcareAlphaDataset(ctx context.Context, request *alphapb.ApplyHealthcareAlphaDatasetRequest) (*alphapb.HealthcareAlphaDataset, error) {
	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDataset(ctx, cl, request)
}

// DeleteDataset handles the gRPC request by passing it to the underlying Dataset Delete() method.
func (s *DatasetServer) DeleteHealthcareAlphaDataset(ctx context.Context, request *alphapb.DeleteHealthcareAlphaDatasetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDataset(ctx, ProtoToDataset(request.GetResource()))

}

// ListHealthcareAlphaDataset handles the gRPC request by passing it to the underlying DatasetList() method.
func (s *DatasetServer) ListHealthcareAlphaDataset(ctx context.Context, request *alphapb.ListHealthcareAlphaDatasetRequest) (*alphapb.ListHealthcareAlphaDatasetResponse, error) {
	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDataset(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.HealthcareAlphaDataset
	for _, r := range resources.Items {
		rp := DatasetToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListHealthcareAlphaDatasetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDataset(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
