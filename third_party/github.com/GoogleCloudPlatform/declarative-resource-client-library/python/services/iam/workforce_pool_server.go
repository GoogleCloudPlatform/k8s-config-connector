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

// WorkforcePoolServer implements the gRPC interface for WorkforcePool.
type WorkforcePoolServer struct{}

// ProtoToWorkforcePoolStateEnum converts a WorkforcePoolStateEnum enum from its proto representation.
func ProtoToIamWorkforcePoolStateEnum(e iampb.IamWorkforcePoolStateEnum) *iam.WorkforcePoolStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := iampb.IamWorkforcePoolStateEnum_name[int32(e)]; ok {
		e := iam.WorkforcePoolStateEnum(n[len("IamWorkforcePoolStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToWorkforcePool converts a WorkforcePool resource from its proto representation.
func ProtoToWorkforcePool(p *iampb.IamWorkforcePool) *iam.WorkforcePool {
	obj := &iam.WorkforcePool{
		Name:            dcl.StringOrNil(p.GetName()),
		SelfLink:        dcl.StringOrNil(p.GetSelfLink()),
		Parent:          dcl.StringOrNil(p.GetParent()),
		DisplayName:     dcl.StringOrNil(p.GetDisplayName()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		State:           ProtoToIamWorkforcePoolStateEnum(p.GetState()),
		Disabled:        dcl.Bool(p.GetDisabled()),
		SessionDuration: dcl.StringOrNil(p.GetSessionDuration()),
		Location:        dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// WorkforcePoolStateEnumToProto converts a WorkforcePoolStateEnum enum to its proto representation.
func IamWorkforcePoolStateEnumToProto(e *iam.WorkforcePoolStateEnum) iampb.IamWorkforcePoolStateEnum {
	if e == nil {
		return iampb.IamWorkforcePoolStateEnum(0)
	}
	if v, ok := iampb.IamWorkforcePoolStateEnum_value["WorkforcePoolStateEnum"+string(*e)]; ok {
		return iampb.IamWorkforcePoolStateEnum(v)
	}
	return iampb.IamWorkforcePoolStateEnum(0)
}

// WorkforcePoolToProto converts a WorkforcePool resource to its proto representation.
func WorkforcePoolToProto(resource *iam.WorkforcePool) *iampb.IamWorkforcePool {
	p := &iampb.IamWorkforcePool{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(IamWorkforcePoolStateEnumToProto(resource.State))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetSessionDuration(dcl.ValueOrEmptyString(resource.SessionDuration))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Apply() method.
func (s *WorkforcePoolServer) applyWorkforcePool(ctx context.Context, c *iam.Client, request *iampb.ApplyIamWorkforcePoolRequest) (*iampb.IamWorkforcePool, error) {
	p := ProtoToWorkforcePool(request.GetResource())
	res, err := c.ApplyWorkforcePool(ctx, p)
	if err != nil {
		return nil, err
	}
	r := WorkforcePoolToProto(res)
	return r, nil
}

// applyIamWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Apply() method.
func (s *WorkforcePoolServer) ApplyIamWorkforcePool(ctx context.Context, request *iampb.ApplyIamWorkforcePoolRequest) (*iampb.IamWorkforcePool, error) {
	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyWorkforcePool(ctx, cl, request)
}

// DeleteWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePool Delete() method.
func (s *WorkforcePoolServer) DeleteIamWorkforcePool(ctx context.Context, request *iampb.DeleteIamWorkforcePoolRequest) (*emptypb.Empty, error) {

	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteWorkforcePool(ctx, ProtoToWorkforcePool(request.GetResource()))

}

// ListIamWorkforcePool handles the gRPC request by passing it to the underlying WorkforcePoolList() method.
func (s *WorkforcePoolServer) ListIamWorkforcePool(ctx context.Context, request *iampb.ListIamWorkforcePoolRequest) (*iampb.ListIamWorkforcePoolResponse, error) {
	cl, err := createConfigWorkforcePool(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListWorkforcePool(ctx, request.GetLocation(), request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*iampb.IamWorkforcePool
	for _, r := range resources.Items {
		rp := WorkforcePoolToProto(r)
		protos = append(protos, rp)
	}
	p := &iampb.ListIamWorkforcePoolResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigWorkforcePool(ctx context.Context, service_account_file string) (*iam.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return iam.NewClient(conf), nil
}
