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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/alpha/iam_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/alpha"
)

// WorkforcePoolServer implements the gRPC interface for WorkforcePool.
type WorkforcePoolServer struct{}

// ProtoToWorkforcePoolStateEnum converts a WorkforcePoolStateEnum enum from its proto representation.
func ProtoToIamAlphaWorkforcePoolStateEnum(e alphapb.IamAlphaWorkforcePoolStateEnum) *alpha.WorkforcePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IamAlphaWorkforcePoolStateEnum_name[int32(e)]; ok {
		e := alpha.WorkforcePoolStateEnum(n[len("IamAlphaWorkforcePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkforcePool converts a WorkforcePool resource from its proto representation.
func ProtoToWorkforcePool(p *alphapb.IamAlphaWorkforcePool) *alpha.WorkforcePool {
	obj := &alpha.WorkforcePool{
		Name:            dcl.StringOrNil(p.GetName()),
		SelfLink:        dcl.StringOrNil(p.GetSelfLink()),
		Parent:          dcl.StringOrNil(p.GetParent()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToIamAlphaWorkforcePoolStateEnum(p.GetState()),
		Disabled:        dcl.Bool(p.GetDisabled()),
		SessionDuration: dcl.StringOrNil(p.GetSessionDuration()),
		Location:        dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkforcePoolStateEnumToProto converts a WorkforcePoolStateEnum enum to its proto representation.
func IamAlphaWorkforcePoolStateEnumToProto(e *alpha.WorkforcePoolStateEnum) alphapb.IamAlphaWorkforcePoolStateEnum {
	if e == nil {
		return alphapb.IamAlphaWorkforcePoolStateEnum(0)
	}
	if v, ok := alphapb.IamAlphaWorkforcePoolStateEnum_value["WorkforcePoolStateEnum"+string(*e)]; ok {
		return alphapb.IamAlphaWorkforcePoolStateEnum(v)
	}
	return alphapb.IamAlphaWorkforcePoolStateEnum(0)
}

// WorkforcePoolToProto converts a WorkforcePool resource to its proto representation.
func WorkforcePoolToProto(resource *alpha.WorkforcePool) *alphapb.IamAlphaWorkforcePool {
	p := &alphapb.IamAlphaWorkforcePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamAlphaWorkforcePoolStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetSessionDuration(dcl.ValueOrEmptyString(resource.SessionDuration))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Apply() method.
func (s *WorkforcePoolServer) applyWorkforcePool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIamAlphaWorkforcePoolRequest) (*alphapb.IamAlphaWorkforcePool, error) {
	p := ProtoToWorkforcePool(request.GetResource())
	res, err := c.ApplyWorkforcePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkforcePoolToProto(res)
	return r, nil
}

// applyIamAlphaWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Apply() method.
func (s *WorkforcePoolServer) ApplyIamAlphaWorkforcePool(ctx context.Context, request *alphapb.ApplyIamAlphaWorkforcePoolRequest) (*alphapb.IamAlphaWorkforcePool, error) {
	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkforcePool(ctx, cl, request)
}

// DeleteWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Delete() method.
func (s *WorkforcePoolServer) DeleteIamAlphaWorkforcePool(ctx context.Context, request *alphapb.DeleteIamAlphaWorkforcePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkforcePool(ctx, ProtoToWorkforcePool(request.GetResource()))

}

// ListIamAlphaWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePoolList() method.
func (s *WorkforcePoolServer) ListIamAlphaWorkforcePool(ctx context.Context, request *alphapb.ListIamAlphaWorkforcePoolRequest) (*alphapb.ListIamAlphaWorkforcePoolResponse, error) {
	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkforcePool(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IamAlphaWorkforcePool
	for _, r := range resources.Items {
		rp := WorkforcePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIamAlphaWorkforcePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkforcePool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
