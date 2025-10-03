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
	spannerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/spanner/spanner_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/spanner"
)

// Server implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToSpannerInstanceStateEnum(e spannerpb.SpannerInstanceStateEnum) *spanner.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := spannerpb.SpannerInstanceStateEnum_name[int32(e)]; ok {
		e := spanner.InstanceStateEnum(n[len("SpannerInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *spannerpb.SpannerInstance) *spanner.Instance {
	obj := &spanner.Instance{
		Name:        dcl.StringOrNil(p.Name),
		Project:     dcl.StringOrNil(p.Project),
		Config:      dcl.StringOrNil(p.Config),
		DisplayName: dcl.StringOrNil(p.DisplayName),
		NodeCount:   dcl.Int64OrNil(p.NodeCount),
		State:       ProtoToSpannerInstanceStateEnum(p.GetState()),
	}
	return obj
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func SpannerInstanceStateEnumToProto(e *spanner.InstanceStateEnum) spannerpb.SpannerInstanceStateEnum {
	if e == nil {
		return spannerpb.SpannerInstanceStateEnum(0)
	}
	if v, ok := spannerpb.SpannerInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return spannerpb.SpannerInstanceStateEnum(v)
	}
	return spannerpb.SpannerInstanceStateEnum(0)
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *spanner.Instance) *spannerpb.SpannerInstance {
	p := &spannerpb.SpannerInstance{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Config:      dcl.ValueOrEmptyString(resource.Config),
		DisplayName: dcl.ValueOrEmptyString(resource.DisplayName),
		NodeCount:   dcl.ValueOrEmptyInt64(resource.NodeCount),
		State:       SpannerInstanceStateEnumToProto(resource.State),
	}

	return p
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *spanner.Client, request *spannerpb.ApplySpannerInstanceRequest) (*spannerpb.SpannerInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// ApplyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplySpannerInstance(ctx context.Context, request *spannerpb.ApplySpannerInstanceRequest) (*spannerpb.SpannerInstance, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteSpannerInstance(ctx context.Context, request *spannerpb.DeleteSpannerInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListSpannerInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListSpannerInstance(ctx context.Context, request *spannerpb.ListSpannerInstanceRequest) (*spannerpb.ListSpannerInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*spannerpb.SpannerInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	return &spannerpb.ListSpannerInstanceResponse{Items: protos}, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*spanner.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return spanner.NewClient(conf), nil
}
