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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/alpha/apigee_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/alpha"
)

// InstanceServer implements the gRPC interface for Instance.
type InstanceServer struct{}

// ProtoToInstancePeeringCidrRangeEnum converts a InstancePeeringCidrRangeEnum enum from its proto representation.
func ProtoToApigeeAlphaInstancePeeringCidrRangeEnum(e alphapb.ApigeeAlphaInstancePeeringCidrRangeEnum) *alpha.InstancePeeringCidrRangeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaInstancePeeringCidrRangeEnum_name[int32(e)]; ok {
		e := alpha.InstancePeeringCidrRangeEnum(n[len("ApigeeAlphaInstancePeeringCidrRangeEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstanceStateEnum converts a InstanceStateEnum enum from its proto representation.
func ProtoToApigeeAlphaInstanceStateEnum(e alphapb.ApigeeAlphaInstanceStateEnum) *alpha.InstanceStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaInstanceStateEnum_name[int32(e)]; ok {
		e := alpha.InstanceStateEnum(n[len("ApigeeAlphaInstanceStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToInstance converts a Instance resource from its proto representation.
func ProtoToInstance(p *alphapb.ApigeeAlphaInstance) *alpha.Instance {
	obj := &alpha.Instance{
		Name:                  dcl.StringOrNil(p.GetName()),
		Location:              dcl.StringOrNil(p.GetLocation()),
		PeeringCidrRange:      ProtoToApigeeAlphaInstancePeeringCidrRangeEnum(p.GetPeeringCidrRange()),
		Host:                  dcl.StringOrNil(p.GetHost()),
		Port:                  dcl.StringOrNil(p.GetPort()),
		Description:           dcl.StringOrNil(p.GetDescription()),
		DisplayName:           dcl.StringOrNil(p.GetDisplayName()),
		CreatedAt:             dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:        dcl.Int64OrNil(p.GetLastModifiedAt()),
		DiskEncryptionKeyName: dcl.StringOrNil(p.GetDiskEncryptionKeyName()),
		State:                 ProtoToApigeeAlphaInstanceStateEnum(p.GetState()),
		ApigeeOrganization:    dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	return obj
}

// InstancePeeringCidrRangeEnumToProto converts a InstancePeeringCidrRangeEnum enum to its proto representation.
func ApigeeAlphaInstancePeeringCidrRangeEnumToProto(e *alpha.InstancePeeringCidrRangeEnum) alphapb.ApigeeAlphaInstancePeeringCidrRangeEnum {
	if e == nil {
		return alphapb.ApigeeAlphaInstancePeeringCidrRangeEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaInstancePeeringCidrRangeEnum_value["InstancePeeringCidrRangeEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaInstancePeeringCidrRangeEnum(v)
	}
	return alphapb.ApigeeAlphaInstancePeeringCidrRangeEnum(0)
}

// InstanceStateEnumToProto converts a InstanceStateEnum enum to its proto representation.
func ApigeeAlphaInstanceStateEnumToProto(e *alpha.InstanceStateEnum) alphapb.ApigeeAlphaInstanceStateEnum {
	if e == nil {
		return alphapb.ApigeeAlphaInstanceStateEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaInstanceStateEnum_value["InstanceStateEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaInstanceStateEnum(v)
	}
	return alphapb.ApigeeAlphaInstanceStateEnum(0)
}

// InstanceToProto converts a Instance resource to its proto representation.
func InstanceToProto(resource *alpha.Instance) *alphapb.ApigeeAlphaInstance {
	p := &alphapb.ApigeeAlphaInstance{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetPeeringCidrRange(ApigeeAlphaInstancePeeringCidrRangeEnumToProto(resource.PeeringCidrRange))
	p.SetHost(dcl.ValueOrEmptyString(resource.Host))
	p.SetPort(dcl.ValueOrEmptyString(resource.Port))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetDiskEncryptionKeyName(dcl.ValueOrEmptyString(resource.DiskEncryptionKeyName))
	p.SetState(ApigeeAlphaInstanceStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))

	return p
}

// applyInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) applyInstance(ctx context.Context, c *alpha.Client, request *alphapb.ApplyApigeeAlphaInstanceRequest) (*alphapb.ApigeeAlphaInstance, error) {
	p := ProtoToInstance(request.GetResource())
	res, err := c.ApplyInstance(ctx, p)
	if err != nil {
		return nil, err
	}
	r := InstanceToProto(res)
	return r, nil
}

// applyApigeeAlphaInstance handles the gRPC request by passing it to the underlying Instance Apply() method.
func (s *InstanceServer) ApplyApigeeAlphaInstance(ctx context.Context, request *alphapb.ApplyApigeeAlphaInstanceRequest) (*alphapb.ApigeeAlphaInstance, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyInstance(ctx, cl, request)
}

// DeleteInstance handles the gRPC request by passing it to the underlying Instance Delete() method.
func (s *InstanceServer) DeleteApigeeAlphaInstance(ctx context.Context, request *alphapb.DeleteApigeeAlphaInstanceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteInstance(ctx, ProtoToInstance(request.GetResource()))

}

// ListApigeeAlphaInstance handles the gRPC request by passing it to the underlying InstanceList() method.
func (s *InstanceServer) ListApigeeAlphaInstance(ctx context.Context, request *alphapb.ListApigeeAlphaInstanceRequest) (*alphapb.ListApigeeAlphaInstanceResponse, error) {
	cl, err := createConfigInstance(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListInstance(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ApigeeAlphaInstance
	for _, r := range resources.Items {
		rp := InstanceToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListApigeeAlphaInstanceResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigInstance(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
