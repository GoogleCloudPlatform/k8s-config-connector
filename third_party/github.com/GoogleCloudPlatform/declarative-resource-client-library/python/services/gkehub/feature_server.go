// Copyright 2023 Google LLC. All Rights Reserved.
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
	gkehubpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/gkehub_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub"
)

// FeatureServer implements the gRPC interface for Feature.
type FeatureServer struct{}

// ProtoToFeatureResourceStateStateEnum converts a FeatureResourceStateStateEnum enum from its proto representation.
func ProtoToGkehubFeatureResourceStateStateEnum(e gkehubpb.GkehubFeatureResourceStateStateEnum) *gkehub.FeatureResourceStateStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkehubpb.GkehubFeatureResourceStateStateEnum_name[int32(e)]; ok {
		e := gkehub.FeatureResourceStateStateEnum(n[len("GkehubFeatureResourceStateStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureStateStateCodeEnum converts a FeatureStateStateCodeEnum enum from its proto representation.
func ProtoToGkehubFeatureStateStateCodeEnum(e gkehubpb.GkehubFeatureStateStateCodeEnum) *gkehub.FeatureStateStateCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := gkehubpb.GkehubFeatureStateStateCodeEnum_name[int32(e)]; ok {
		e := gkehub.FeatureStateStateCodeEnum(n[len("GkehubFeatureStateStateCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureResourceState converts a FeatureResourceState object from its proto representation.
func ProtoToGkehubFeatureResourceState(p *gkehubpb.GkehubFeatureResourceState) *gkehub.FeatureResourceState {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureResourceState{
		State:        ProtoToGkehubFeatureResourceStateStateEnum(p.GetState()),
		HasResources: dcl.Bool(p.GetHasResources()),
	}
	return obj
}

// ProtoToFeatureSpec converts a FeatureSpec object from its proto representation.
func ProtoToGkehubFeatureSpec(p *gkehubpb.GkehubFeatureSpec) *gkehub.FeatureSpec {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureSpec{
		Multiclusteringress: ProtoToGkehubFeatureSpecMulticlusteringress(p.GetMulticlusteringress()),
	}
	return obj
}

// ProtoToFeatureSpecMulticlusteringress converts a FeatureSpecMulticlusteringress object from its proto representation.
func ProtoToGkehubFeatureSpecMulticlusteringress(p *gkehubpb.GkehubFeatureSpecMulticlusteringress) *gkehub.FeatureSpecMulticlusteringress {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureSpecMulticlusteringress{
		ConfigMembership: dcl.StringOrNil(p.GetConfigMembership()),
	}
	return obj
}

// ProtoToFeatureState converts a FeatureState object from its proto representation.
func ProtoToGkehubFeatureState(p *gkehubpb.GkehubFeatureState) *gkehub.FeatureState {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureState{
		State: ProtoToGkehubFeatureStateState(p.GetState()),
	}
	return obj
}

// ProtoToFeatureStateState converts a FeatureStateState object from its proto representation.
func ProtoToGkehubFeatureStateState(p *gkehubpb.GkehubFeatureStateState) *gkehub.FeatureStateState {
	if p == nil {
		return nil
	}
	obj := &gkehub.FeatureStateState{
		Code:        ProtoToGkehubFeatureStateStateCodeEnum(p.GetCode()),
		Description: dcl.StringOrNil(p.GetDescription()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
	}
	return obj
}

// ProtoToFeature converts a Feature resource from its proto representation.
func ProtoToFeature(p *gkehubpb.GkehubFeature) *gkehub.Feature {
	obj := &gkehub.Feature{
		Name:          dcl.StringOrNil(p.GetName()),
		ResourceState: ProtoToGkehubFeatureResourceState(p.GetResourceState()),
		Spec:          ProtoToGkehubFeatureSpec(p.GetSpec()),
		State:         ProtoToGkehubFeatureState(p.GetState()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:    dcl.StringOrNil(p.GetDeleteTime()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// FeatureResourceStateStateEnumToProto converts a FeatureResourceStateStateEnum enum to its proto representation.
func GkehubFeatureResourceStateStateEnumToProto(e *gkehub.FeatureResourceStateStateEnum) gkehubpb.GkehubFeatureResourceStateStateEnum {
	if e == nil {
		return gkehubpb.GkehubFeatureResourceStateStateEnum(0)
	}
	if v, ok := gkehubpb.GkehubFeatureResourceStateStateEnum_value["FeatureResourceStateStateEnum"+string(*e)]; ok {
		return gkehubpb.GkehubFeatureResourceStateStateEnum(v)
	}
	return gkehubpb.GkehubFeatureResourceStateStateEnum(0)
}

// FeatureStateStateCodeEnumToProto converts a FeatureStateStateCodeEnum enum to its proto representation.
func GkehubFeatureStateStateCodeEnumToProto(e *gkehub.FeatureStateStateCodeEnum) gkehubpb.GkehubFeatureStateStateCodeEnum {
	if e == nil {
		return gkehubpb.GkehubFeatureStateStateCodeEnum(0)
	}
	if v, ok := gkehubpb.GkehubFeatureStateStateCodeEnum_value["FeatureStateStateCodeEnum"+string(*e)]; ok {
		return gkehubpb.GkehubFeatureStateStateCodeEnum(v)
	}
	return gkehubpb.GkehubFeatureStateStateCodeEnum(0)
}

// FeatureResourceStateToProto converts a FeatureResourceState object to its proto representation.
func GkehubFeatureResourceStateToProto(o *gkehub.FeatureResourceState) *gkehubpb.GkehubFeatureResourceState {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureResourceState{}
	p.SetState(GkehubFeatureResourceStateStateEnumToProto(o.State))
	p.SetHasResources(dcl.ValueOrEmptyBool(o.HasResources))
	return p
}

// FeatureSpecToProto converts a FeatureSpec object to its proto representation.
func GkehubFeatureSpecToProto(o *gkehub.FeatureSpec) *gkehubpb.GkehubFeatureSpec {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureSpec{}
	p.SetMulticlusteringress(GkehubFeatureSpecMulticlusteringressToProto(o.Multiclusteringress))
	return p
}

// FeatureSpecMulticlusteringressToProto converts a FeatureSpecMulticlusteringress object to its proto representation.
func GkehubFeatureSpecMulticlusteringressToProto(o *gkehub.FeatureSpecMulticlusteringress) *gkehubpb.GkehubFeatureSpecMulticlusteringress {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureSpecMulticlusteringress{}
	p.SetConfigMembership(dcl.ValueOrEmptyString(o.ConfigMembership))
	return p
}

// FeatureStateToProto converts a FeatureState object to its proto representation.
func GkehubFeatureStateToProto(o *gkehub.FeatureState) *gkehubpb.GkehubFeatureState {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureState{}
	p.SetState(GkehubFeatureStateStateToProto(o.State))
	return p
}

// FeatureStateStateToProto converts a FeatureStateState object to its proto representation.
func GkehubFeatureStateStateToProto(o *gkehub.FeatureStateState) *gkehubpb.GkehubFeatureStateState {
	if o == nil {
		return nil
	}
	p := &gkehubpb.GkehubFeatureStateState{}
	p.SetCode(GkehubFeatureStateStateCodeEnumToProto(o.Code))
	p.SetDescription(dcl.ValueOrEmptyString(o.Description))
	p.SetUpdateTime(dcl.ValueOrEmptyString(o.UpdateTime))
	return p
}

// FeatureToProto converts a Feature resource to its proto representation.
func FeatureToProto(resource *gkehub.Feature) *gkehubpb.GkehubFeature {
	p := &gkehubpb.GkehubFeature{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetResourceState(GkehubFeatureResourceStateToProto(resource.ResourceState))
	p.SetSpec(GkehubFeatureSpecToProto(resource.Spec))
	p.SetState(GkehubFeatureStateToProto(resource.State))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) applyFeature(ctx context.Context, c *gkehub.Client, request *gkehubpb.ApplyGkehubFeatureRequest) (*gkehubpb.GkehubFeature, error) {
	p := ProtoToFeature(request.GetResource())
	res, err := c.ApplyFeature(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureToProto(res)
	return r, nil
}

// applyGkehubFeature handles the gRPC request by passing it to the underlying Feature Apply() method.
func (s *FeatureServer) ApplyGkehubFeature(ctx context.Context, request *gkehubpb.ApplyGkehubFeatureRequest) (*gkehubpb.GkehubFeature, error) {
	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFeature(ctx, cl, request)
}

// DeleteFeature handles the gRPC request by passing it to the underlying Feature Delete() method.
func (s *FeatureServer) DeleteGkehubFeature(ctx context.Context, request *gkehubpb.DeleteGkehubFeatureRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeature(ctx, ProtoToFeature(request.GetResource()))

}

// ListGkehubFeature handles the gRPC request by passing it to the underlying FeatureList() method.
func (s *FeatureServer) ListGkehubFeature(ctx context.Context, request *gkehubpb.ListGkehubFeatureRequest) (*gkehubpb.ListGkehubFeatureResponse, error) {
	cl, err := createConfigFeature(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeature(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*gkehubpb.GkehubFeature
	for _, r := range resources.Items {
		rp := FeatureToProto(r)
		protos = append(protos, rp)
	}
	p := &gkehubpb.ListGkehubFeatureResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFeature(ctx context.Context, service_account_file string) (*gkehub.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return gkehub.NewClient(conf), nil
}
