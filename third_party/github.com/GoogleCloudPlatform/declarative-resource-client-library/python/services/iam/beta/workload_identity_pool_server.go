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

// WorkloadIdentityPoolServer implements the gRPC interface for WorkloadIdentityPool.
type WorkloadIdentityPoolServer struct{}

// ProtoToWorkloadIdentityPoolStateEnum converts a WorkloadIdentityPoolStateEnum enum from its proto representation.
func ProtoToIamBetaWorkloadIdentityPoolStateEnum(e betapb.IamBetaWorkloadIdentityPoolStateEnum) *beta.WorkloadIdentityPoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IamBetaWorkloadIdentityPoolStateEnum_name[int32(e)]; ok {
		e := beta.WorkloadIdentityPoolStateEnum(n[len("IamBetaWorkloadIdentityPoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadIdentityPool converts a WorkloadIdentityPool resource from its proto representation.
func ProtoToWorkloadIdentityPool(p *betapb.IamBetaWorkloadIdentityPool) *beta.WorkloadIdentityPool {
	obj := &beta.WorkloadIdentityPool{
		Name:        dcl.StringOrNil(p.GetName()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		State:       ProtoToIamBetaWorkloadIdentityPoolStateEnum(p.GetState()),
		Disabled:    dcl.Bool(p.GetDisabled()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkloadIdentityPoolStateEnumToProto converts a WorkloadIdentityPoolStateEnum enum to its proto representation.
func IamBetaWorkloadIdentityPoolStateEnumToProto(e *beta.WorkloadIdentityPoolStateEnum) betapb.IamBetaWorkloadIdentityPoolStateEnum {
	if e == nil {
		return betapb.IamBetaWorkloadIdentityPoolStateEnum(0)
	}
	if v, ok := betapb.IamBetaWorkloadIdentityPoolStateEnum_value["WorkloadIdentityPoolStateEnum"+string(*e)]; ok {
		return betapb.IamBetaWorkloadIdentityPoolStateEnum(v)
	}
	return betapb.IamBetaWorkloadIdentityPoolStateEnum(0)
}

// WorkloadIdentityPoolToProto converts a WorkloadIdentityPool resource to its proto representation.
func WorkloadIdentityPoolToProto(resource *beta.WorkloadIdentityPool) *betapb.IamBetaWorkloadIdentityPool {
	p := &betapb.IamBetaWorkloadIdentityPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamBetaWorkloadIdentityPoolStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Apply() method.
func (s *WorkloadIdentityPoolServer) applyWorkloadIdentityPool(ctx context.Context, c *beta.Client, request *betapb.ApplyIamBetaWorkloadIdentityPoolRequest) (*betapb.IamBetaWorkloadIdentityPool, error) {
	p := ProtoToWorkloadIdentityPool(request.GetResource())
	res, err := c.ApplyWorkloadIdentityPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadIdentityPoolToProto(res)
	return r, nil
}

// applyIamBetaWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Apply() method.
func (s *WorkloadIdentityPoolServer) ApplyIamBetaWorkloadIdentityPool(ctx context.Context, request *betapb.ApplyIamBetaWorkloadIdentityPoolRequest) (*betapb.IamBetaWorkloadIdentityPool, error) {
	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkloadIdentityPool(ctx, cl, request)
}

// DeleteWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Delete() method.
func (s *WorkloadIdentityPoolServer) DeleteIamBetaWorkloadIdentityPool(ctx context.Context, request *betapb.DeleteIamBetaWorkloadIdentityPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkloadIdentityPool(ctx, ProtoToWorkloadIdentityPool(request.GetResource()))

}

// ListIamBetaWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPoolList() method.
func (s *WorkloadIdentityPoolServer) ListIamBetaWorkloadIdentityPool(ctx context.Context, request *betapb.ListIamBetaWorkloadIdentityPoolRequest) (*betapb.ListIamBetaWorkloadIdentityPoolResponse, error) {
	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkloadIdentityPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IamBetaWorkloadIdentityPool
	for _, r := range resources.Items {
		rp := WorkloadIdentityPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListIamBetaWorkloadIdentityPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkloadIdentityPool(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
