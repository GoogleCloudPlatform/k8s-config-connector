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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/beta/iam_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/beta"
)

// WorkforcePoolServer implements the gRPC interface for WorkforcePool.
type WorkforcePoolServer struct{}

// ProtoToWorkforcePoolStateEnum converts a WorkforcePoolStateEnum enum from its proto representation.
func ProtoToIamBetaWorkforcePoolStateEnum(e betapb.IamBetaWorkforcePoolStateEnum) *beta.WorkforcePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IamBetaWorkforcePoolStateEnum_name[int32(e)]; ok {
		e := beta.WorkforcePoolStateEnum(n[len("IamBetaWorkforcePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkforcePool converts a WorkforcePool resource from its proto representation.
func ProtoToWorkforcePool(p *betapb.IamBetaWorkforcePool) *beta.WorkforcePool {
	obj := &beta.WorkforcePool{
		Name:            dcl.StringOrNil(p.GetName()),
		SelfLink:        dcl.StringOrNil(p.GetSelfLink()),
		Parent:          dcl.StringOrNil(p.GetParent()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToIamBetaWorkforcePoolStateEnum(p.GetState()),
		Disabled:        dcl.Bool(p.GetDisabled()),
		SessionDuration: dcl.StringOrNil(p.GetSessionDuration()),
		Location:        dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkforcePoolStateEnumToProto converts a WorkforcePoolStateEnum enum to its proto representation.
func IamBetaWorkforcePoolStateEnumToProto(e *beta.WorkforcePoolStateEnum) betapb.IamBetaWorkforcePoolStateEnum {
	if e == nil {
		return betapb.IamBetaWorkforcePoolStateEnum(0)
	}
	if v, ok := betapb.IamBetaWorkforcePoolStateEnum_value["WorkforcePoolStateEnum"+string(*e)]; ok {
		return betapb.IamBetaWorkforcePoolStateEnum(v)
	}
	return betapb.IamBetaWorkforcePoolStateEnum(0)
}

// WorkforcePoolToProto converts a WorkforcePool resource to its proto representation.
func WorkforcePoolToProto(resource *beta.WorkforcePool) *betapb.IamBetaWorkforcePool {
	p := &betapb.IamBetaWorkforcePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamBetaWorkforcePoolStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetSessionDuration(dcl.ValueOrEmptyString(resource.SessionDuration))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Apply() method.
func (s *WorkforcePoolServer) applyWorkforcePool(ctx context.Context, c *beta.Client, request *betapb.ApplyIamBetaWorkforcePoolRequest) (*betapb.IamBetaWorkforcePool, error) {
	p := ProtoToWorkforcePool(request.GetResource())
	res, err := c.ApplyWorkforcePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkforcePoolToProto(res)
	return r, nil
}

// applyIamBetaWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Apply() method.
func (s *WorkforcePoolServer) ApplyIamBetaWorkforcePool(ctx context.Context, request *betapb.ApplyIamBetaWorkforcePoolRequest) (*betapb.IamBetaWorkforcePool, error) {
	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkforcePool(ctx, cl, request)
}

// DeleteWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Delete() method.
func (s *WorkforcePoolServer) DeleteIamBetaWorkforcePool(ctx context.Context, request *betapb.DeleteIamBetaWorkforcePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkforcePool(ctx, ProtoToWorkforcePool(request.GetResource()))

}

// ListIamBetaWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePoolList() method.
func (s *WorkforcePoolServer) ListIamBetaWorkforcePool(ctx context.Context, request *betapb.ListIamBetaWorkforcePoolRequest) (*betapb.ListIamBetaWorkforcePoolResponse, error) {
	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkforcePool(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IamBetaWorkforcePool
	for _, r := range resources.Items {
		rp := WorkforcePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListIamBetaWorkforcePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkforcePool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
