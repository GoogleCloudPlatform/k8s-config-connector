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
	apigeepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/apigee_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee"
)

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstancePeeringCidrRangeEnum converts a InstancePeeringCidrRangeEnum enum from its proto representation.
func ProtoToApigeeInstancePeeringCidrRangeEnum(e apigeepb.ApigeeInstancePeeringCidrRangeEnum) *apigee.InstancePeeringCidrRangeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeInstancePeeringCidrRangeEnum_name[int32(e)]; ok {
		e := apigee.InstancePeeringCidrRangeEnum(n[len("ApigeeInstancePeeringCidrRangeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToApigeeInstanceStateEnum(e apigeepb.ApigeeInstanceStateEnum) *apigee.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeInstanceStateEnum_name[int32(e)]; ok {
		e := apigee.InstanceStateEnum(n[len("ApigeeInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *apigeepb.ApigeeInstance) *apigee.Instance {
	obj := &apigee.Instance{
		Name:                  dcl.StringOrNil(p.GetName()),
		Location:              dcl.StringOrNil(p.GetLocation()),
		PeeringCidrRange:      ProtoToApigeeInstancePeeringCidrRangeEnum(p.GetPeeringCidrRange()),
		Host:                  dcl.StringOrNil(p.GetHost()),
		Port:                  dcl.StringOrNil(p.GetPort()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		CreatedAt:             dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:        dcl.Int64OrNil(p.GetLastModifiedAt()),
		DiskEncryptionKeyName: dcl.StringOrNil(p.GetDiskEncryptionKeyName()),
		State:                 ProtoToApigeeInstanceStateEnum(p.GetState()),
		ApigeeOrganization:    dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	return obj
}

// InstancePeeringCidrRangeEnumToProto converts a InstancePeeringCidrRangeEnum enum to its proto representation.
func ApigeeInstancePeeringCidrRangeEnumToProto(e *apigee.InstancePeeringCidrRangeEnum) apigeepb.ApigeeInstancePeeringCidrRangeEnum {
	if e == nil {
		return apigeepb.ApigeeInstancePeeringCidrRangeEnum(0)
	}
	if v, ok := apigeepb.ApigeeInstancePeeringCidrRangeEnum_value["InstancePeeringCidrRangeEnum"+string(*e)]; ok {
		return apigeepb.ApigeeInstancePeeringCidrRangeEnum(v)
	}
	return apigeepb.ApigeeInstancePeeringCidrRangeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func ApigeeInstanceStateEnumToProto(e *apigee.InstanceStateEnum) apigeepb.ApigeeInstanceStateEnum {
	if e == nil {
		return apigeepb.ApigeeInstanceStateEnum(0)
	}
	if v, ok := apigeepb.ApigeeInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return apigeepb.ApigeeInstanceStateEnum(v)
	}
	return apigeepb.ApigeeInstanceStateEnum(0)
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *apigee.Instance) *apigeepb.ApigeeInstance {
	p := &apigeepb.ApigeeInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetPeeringCidrRange(ApigeeInstancePeeringCidrRangeEnumToProto(resource.PeeringCidrRange))
	p.SetHost(dcl.ValueOrEmptyString(resource.Host))
	p.SetPort(dcl.ValueOrEmptyString(resource.Port))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetDiskEncryptionKeyName(dcl.ValueOrEmptyString(resource.DiskEncryptionKeyName))
	p.SetState(ApigeeInstanceStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *apigee.Client, request *apigeepb.ApplyApigeeInstanceRequest) (*apigeepb.ApigeeInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyApigeeInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyApigeeInstance(ctx context.Context, request *apigeepb.ApplyApigeeInstanceRequest) (*apigeepb.ApigeeInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteApigeeInstance(ctx context.Context, request *apigeepb.DeleteApigeeInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListApigeeInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListApigeeInstance(ctx context.Context, request *apigeepb.ListApigeeInstanceRequest) (*apigeepb.ListApigeeInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*apigeepb.ApigeeInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &apigeepb.ListApigeeInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*apigee.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apigee.NewClient(conf), nil
}
