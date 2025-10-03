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
	iampb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/iam_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
)

// WorkloadIdentityPoolServer implements the gRPC interface for WorkloadIdentityPool.
type WorkloadIdentityPoolServer struct{}

// ProtoToWorkloadIdentityPoolStateEnum converts a WorkloadIdentityPoolStateEnum enum from its proto representation.
func ProtoToIamWorkloadIdentityPoolStateEnum(e iampb.IamWorkloadIdentityPoolStateEnum) *iam.WorkloadIdentityPoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := iampb.IamWorkloadIdentityPoolStateEnum_name[int32(e)]; ok {
		e := iam.WorkloadIdentityPoolStateEnum(n[len("IamWorkloadIdentityPoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkloadIdentityPool converts a WorkloadIdentityPool resource from its proto representation.
func ProtoToWorkloadIdentityPool(p *iampb.IamWorkloadIdentityPool) *iam.WorkloadIdentityPool {
	obj := &iam.WorkloadIdentityPool{
		Name:        dcl.StringOrNil(p.GetName()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		State:       ProtoToIamWorkloadIdentityPoolStateEnum(p.GetState()),
		Disabled:    dcl.Bool(p.GetDisabled()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkloadIdentityPoolStateEnumToProto converts a WorkloadIdentityPoolStateEnum enum to its proto representation.
func IamWorkloadIdentityPoolStateEnumToProto(e *iam.WorkloadIdentityPoolStateEnum) iampb.IamWorkloadIdentityPoolStateEnum {
	if e == nil {
		return iampb.IamWorkloadIdentityPoolStateEnum(0)
	}
	if v, ok := iampb.IamWorkloadIdentityPoolStateEnum_value["WorkloadIdentityPoolStateEnum"+string(*e)]; ok {
		return iampb.IamWorkloadIdentityPoolStateEnum(v)
	}
	return iampb.IamWorkloadIdentityPoolStateEnum(0)
}

// WorkloadIdentityPoolToProto converts a WorkloadIdentityPool resource to its proto representation.
func WorkloadIdentityPoolToProto(resource *iam.WorkloadIdentityPool) *iampb.IamWorkloadIdentityPool {
	p := &iampb.IamWorkloadIdentityPool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamWorkloadIdentityPoolStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Apply() method.
func (s *WorkloadIdentityPoolServer) applyWorkloadIdentityPool(ctx context.Context, c *iam.Client, request *iampb.ApplyIamWorkloadIdentityPoolRequest) (*iampb.IamWorkloadIdentityPool, error) {
	p := ProtoToWorkloadIdentityPool(request.GetResource())
	res, err := c.ApplyWorkloadIdentityPool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkloadIdentityPoolToProto(res)
	return r, nil
}

// applyIamWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Apply() method.
func (s *WorkloadIdentityPoolServer) ApplyIamWorkloadIdentityPool(ctx context.Context, request *iampb.ApplyIamWorkloadIdentityPoolRequest) (*iampb.IamWorkloadIdentityPool, error) {
	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkloadIdentityPool(ctx, cl, request)
}

// DeleteWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPool Delete() method.
func (s *WorkloadIdentityPoolServer) DeleteIamWorkloadIdentityPool(ctx context.Context, request *iampb.DeleteIamWorkloadIdentityPoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkloadIdentityPool(ctx, ProtoToWorkloadIdentityPool(request.GetResource()))

}

// ListIamWorkloadIdentityPool handles the gRPC request by passing it to the underlying WorkloadIdentityPoolList() method.
func (s *WorkloadIdentityPoolServer) ListIamWorkloadIdentityPool(ctx context.Context, request *iampb.ListIamWorkloadIdentityPoolRequest) (*iampb.ListIamWorkloadIdentityPoolResponse, error) {
	cl, err := createConfigWorkloadIdentityPool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkloadIdentityPool(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*iampb.IamWorkloadIdentityPool
	for _, r := range resources.Items {
		rp := WorkloadIdentityPoolToProto(r)
		protos = append(protos, rp)
	}
	p := &iampb.ListIamWorkloadIdentityPoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkloadIdentityPool(ctx context.Context, service_account_file string) (*iam.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return iam.NewClient(conf), nil
}
