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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/beta/apigee_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/beta"
)

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstancePeeringCidrRangeEnum converts a InstancePeeringCidrRangeEnum enum from its proto representation.
func ProtoToApigeeBetaInstancePeeringCidrRangeEnum(e betapb.ApigeeBetaInstancePeeringCidrRangeEnum) *beta.InstancePeeringCidrRangeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ApigeeBetaInstancePeeringCidrRangeEnum_name[int32(e)]; ok {
		e := beta.InstancePeeringCidrRangeEnum(n[len("ApigeeBetaInstancePeeringCidrRangeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToApigeeBetaInstanceStateEnum(e betapb.ApigeeBetaInstanceStateEnum) *beta.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ApigeeBetaInstanceStateEnum_name[int32(e)]; ok {
		e := beta.InstanceStateEnum(n[len("ApigeeBetaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *betapb.ApigeeBetaInstance) *beta.Instance {
	obj := &beta.Instance{
		Name:                  dcl.StringOrNil(p.GetName()),
		Location:              dcl.StringOrNil(p.GetLocation()),
		PeeringCidrRange:      ProtoToApigeeBetaInstancePeeringCidrRangeEnum(p.GetPeeringCidrRange()),
		Host:                  dcl.StringOrNil(p.GetHost()),
		Port:                  dcl.StringOrNil(p.GetPort()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		CreatedAt:             dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:        dcl.Int64OrNil(p.GetLastModifiedAt()),
		DiskEncryptionKeyName: dcl.StringOrNil(p.GetDiskEncryptionKeyName()),
		State:                 ProtoToApigeeBetaInstanceStateEnum(p.GetState()),
		ApigeeOrganization:    dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	return obj
}

// InstancePeeringCidrRangeEnumToProto converts a InstancePeeringCidrRangeEnum enum to its proto representation.
func ApigeeBetaInstancePeeringCidrRangeEnumToProto(e *beta.InstancePeeringCidrRangeEnum) betapb.ApigeeBetaInstancePeeringCidrRangeEnum {
	if e == nil {
		return betapb.ApigeeBetaInstancePeeringCidrRangeEnum(0)
	}
	if v, ok := betapb.ApigeeBetaInstancePeeringCidrRangeEnum_value["InstancePeeringCidrRangeEnum"+string(*e)]; ok {
		return betapb.ApigeeBetaInstancePeeringCidrRangeEnum(v)
	}
	return betapb.ApigeeBetaInstancePeeringCidrRangeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func ApigeeBetaInstanceStateEnumToProto(e *beta.InstanceStateEnum) betapb.ApigeeBetaInstanceStateEnum {
	if e == nil {
		return betapb.ApigeeBetaInstanceStateEnum(0)
	}
	if v, ok := betapb.ApigeeBetaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return betapb.ApigeeBetaInstanceStateEnum(v)
	}
	return betapb.ApigeeBetaInstanceStateEnum(0)
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *beta.Instance) *betapb.ApigeeBetaInstance {
	p := &betapb.ApigeeBetaInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetPeeringCidrRange(ApigeeBetaInstancePeeringCidrRangeEnumToProto(resource.PeeringCidrRange))
	p.SetHost(dcl.ValueOrEmptyString(resource.Host))
	p.SetPort(dcl.ValueOrEmptyString(resource.Port))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetDiskEncryptionKeyName(dcl.ValueOrEmptyString(resource.DiskEncryptionKeyName))
	p.SetState(ApigeeBetaInstanceStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *beta.Client, request *betapb.ApplyApigeeBetaInstanceRequest) (*betapb.ApigeeBetaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyApigeeBetaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyApigeeBetaInstance(ctx context.Context, request *betapb.ApplyApigeeBetaInstanceRequest) (*betapb.ApigeeBetaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteApigeeBetaInstance(ctx context.Context, request *betapb.DeleteApigeeBetaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListApigeeBetaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListApigeeBetaInstance(ctx context.Context, request *betapb.ListApigeeBetaInstanceRequest) (*betapb.ListApigeeBetaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ApigeeBetaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListApigeeBetaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
