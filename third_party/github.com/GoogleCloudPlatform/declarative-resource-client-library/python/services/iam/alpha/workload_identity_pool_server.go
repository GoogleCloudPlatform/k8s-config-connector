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

// WorkloadIdentityPoolServer implements the gRPC interface for WorkloadIdentityPool.
type WorkloadIdentityPoolServer struct{}

// ProtoToWorkloadIdentityPoolStateEnum converts a WorkloadIdentityPoolStateEnum enum from its proto representation.
func ProtoToIamAlphaWorkloadIdentityPoolStateEnum(e alphapb.IamAlphaWorkloadIdentityPoolStateEnum) *alpha.WorkloadIdentityPoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.IamAlphaWorkloadIdentityPoolStateEnum_name[int32(e)]; ok {
		e := alpha.WorkloadIdentityPoolStateEnum(n[len("IamAlphaWorkloadIdentityPoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadIdentityPool converts a WorkloadIdentityPool resource from its proto representation.
func ProtoToWorkloadIdentityPool(p *alphapb.IamAlphaWorkloadIdentityPool) *alpha.WorkloadIdentityPool {
	obj := &alpha.WorkloadIdentityPool{
		Name:        dcl.StringOrNil(p.GetName()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		State:       ProtoToIamAlphaWorkloadIdentityPoolStateEnum(p.GetState()),
		Disabled:    dcl.Bool(p.GetDisabled()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkloadIdentityPoolStateEnumToProto converts a WorkloadIdentityPoolStateEnum enum to its proto representation.
func IamAlphaWorkloadIdentityPoolStateEnumToProto(e *alpha.WorkloadIdentityPoolStateEnum) alphapb.IamAlphaWorkloadIdentityPoolStateEnum {
	if e == nil {
		return alphapb.IamAlphaWorkloadIdentityPoolStateEnum(0)
	}
	if v, ok := alphapb.IamAlphaWorkloadIdentityPoolStateEnum_value["WorkloadIdentityPoolStateEnum"+string(*e)]; ok {
		return alphapb.IamAlphaWorkloadIdentityPoolStateEnum(v)
	}
	return alphapb.IamAlphaWorkloadIdentityPoolStateEnum(0)
}

// WorkloadIdentityPoolToProto converts a WorkloadIdentityPool resource to its proto representation.
func WorkloadIdentityPoolToProto(resource *alpha.WorkloadIdentityPool) *alphapb.IamAlphaWorkloadIdentityPool {
	p := &alphapb.IamAlphaWorkloadIdentityPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamAlphaWorkloadIdentityPoolStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Apply() method.
func (s *WorkloadIdentityPoolServer) applyWorkloadIdentityPool(ctx context.Context, c *alpha.Client, request *alphapb.ApplyIamAlphaWorkloadIdentityPoolRequest) (*alphapb.IamAlphaWorkloadIdentityPool, error) {
	p := ProtoToWorkloadIdentityPool(request.GetResource())
	res, err := c.ApplyWorkloadIdentityPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadIdentityPoolToProto(res)
	return r, nil
}

// applyIamAlphaWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Apply() method.
func (s *WorkloadIdentityPoolServer) ApplyIamAlphaWorkloadIdentityPool(ctx context.Context, request *alphapb.ApplyIamAlphaWorkloadIdentityPoolRequest) (*alphapb.IamAlphaWorkloadIdentityPool, error) {
	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkloadIdentityPool(ctx, cl, request)
}

// DeleteWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Delete() method.
func (s *WorkloadIdentityPoolServer) DeleteIamAlphaWorkloadIdentityPool(ctx context.Context, request *alphapb.DeleteIamAlphaWorkloadIdentityPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkloadIdentityPool(ctx, ProtoToWorkloadIdentityPool(request.GetResource()))

}

// ListIamAlphaWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPoolList() method.
func (s *WorkloadIdentityPoolServer) ListIamAlphaWorkloadIdentityPool(ctx context.Context, request *alphapb.ListIamAlphaWorkloadIdentityPoolRequest) (*alphapb.ListIamAlphaWorkloadIdentityPoolResponse, error) {
	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkloadIdentityPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.IamAlphaWorkloadIdentityPool
	for _, r := range resources.Items {
		rp := WorkloadIdentityPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListIamAlphaWorkloadIdentityPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkloadIdentityPool(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
