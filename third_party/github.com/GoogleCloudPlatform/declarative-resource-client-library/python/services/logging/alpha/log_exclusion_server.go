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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/logging/alpha/logging_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/logging/alpha"
)

// LogExclusionServer implements the gRPC interface for LogExclusion.
type LogExclusionServer struct{}

// ProtoToLogExclusion converts a LogExclusion resource from its proto representation.
func ProtoToLogExclusion(p *alphapb.LoggingAlphaLogExclusion) *alpha.LogExclusion {
	obj := &alpha.LogExclusion{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		Filter:      dcl.StringOrNil(p.GetFilter()),
		Disabled:    dcl.Bool(p.GetDisabled()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Parent:      dcl.StringOrNil(p.GetParent()),
	}
	return obj
}

// LogExclusionToProto converts a LogExclusion resource to its proto representation.
func LogExclusionToProto(resource *alpha.LogExclusion) *alphapb.LoggingAlphaLogExclusion {
	p := &alphapb.LoggingAlphaLogExclusion{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetFilter(dcl.ValueOrEmptyString(resource.Filter))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))

	return p
}

// applyLogExclusion handles the gRPC request by passing it to the underlying LogExclusion Apply() method.
func (s *LogExclusionServer) applyLogExclusion(ctx context.Context, c *alpha.Client, request *alphapb.ApplyLoggingAlphaLogExclusionRequest) (*alphapb.LoggingAlphaLogExclusion, error) {
	p := ProtoToLogExclusion(request.GetResource())
	res, err := c.ApplyLogExclusion(ctx, p)
	if err != nil {
		return nil, err
	}
	r := LogExclusionToProto(res)
	return r, nil
}

// applyLoggingAlphaLogExclusion handles the gRPC request by passing it to the underlying LogExclusion Apply() method.
func (s *LogExclusionServer) ApplyLoggingAlphaLogExclusion(ctx context.Context, request *alphapb.ApplyLoggingAlphaLogExclusionRequest) (*alphapb.LoggingAlphaLogExclusion, error) {
	cl, err := createConfigLogExclusion(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyLogExclusion(ctx, cl, request)
}

// DeleteLogExclusion handles the gRPC request by passing it to the underlying LogExclusion Delete() method.
func (s *LogExclusionServer) DeleteLoggingAlphaLogExclusion(ctx context.Context, request *alphapb.DeleteLoggingAlphaLogExclusionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigLogExclusion(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteLogExclusion(ctx, ProtoToLogExclusion(request.GetResource()))

}

// ListLoggingAlphaLogExclusion handles the gRPC request by passing it to the underlying LogExclusionList() method.
func (s *LogExclusionServer) ListLoggingAlphaLogExclusion(ctx context.Context, request *alphapb.ListLoggingAlphaLogExclusionRequest) (*alphapb.ListLoggingAlphaLogExclusionResponse, error) {
	cl, err := createConfigLogExclusion(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListLogExclusion(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.LoggingAlphaLogExclusion
	for _, r := range resources.Items {
		rp := LogExclusionToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListLoggingAlphaLogExclusionResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigLogExclusion(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
