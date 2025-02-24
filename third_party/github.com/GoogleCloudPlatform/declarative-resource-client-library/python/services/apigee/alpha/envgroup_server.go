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

// EnvgroupServer implements the gRPC interface for Envgroup.
type EnvgroupServer struct{}

// ProtoToEnvgroupStateEnum converts a EnvgroupStateEnum enum from its proto representation.
func ProtoToApigeeAlphaEnvgroupStateEnum(e alphapb.ApigeeAlphaEnvgroupStateEnum) *alpha.EnvgroupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaEnvgroupStateEnum_name[int32(e)]; ok {
		e := alpha.EnvgroupStateEnum(n[len("ApigeeAlphaEnvgroupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvgroup converts a Envgroup resource from its proto representation.
func ProtoToEnvgroup(p *alphapb.ApigeeAlphaEnvgroup) *alpha.Envgroup {
	obj := &alpha.Envgroup{
		Name:               dcl.StringOrNil(p.GetName()),
		CreatedAt:          dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:     dcl.Int64OrNil(p.GetLastModifiedAt()),
		State:              ProtoToApigeeAlphaEnvgroupStateEnum(p.GetState()),
		ApigeeOrganization: dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	for _, r := range p.GetHostnames() {
		obj.Hostnames = append(obj.Hostnames, r)
	}
	return obj
}

// EnvgroupStateEnumToProto converts a EnvgroupStateEnum enum to its proto representation.
func ApigeeAlphaEnvgroupStateEnumToProto(e *alpha.EnvgroupStateEnum) alphapb.ApigeeAlphaEnvgroupStateEnum {
	if e == nil {
		return alphapb.ApigeeAlphaEnvgroupStateEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaEnvgroupStateEnum_value["EnvgroupStateEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaEnvgroupStateEnum(v)
	}
	return alphapb.ApigeeAlphaEnvgroupStateEnum(0)
}

// EnvgroupToProto converts a Envgroup resource to its proto representation.
func EnvgroupToProto(resource *alpha.Envgroup) *alphapb.ApigeeAlphaEnvgroup {
	p := &alphapb.ApigeeAlphaEnvgroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetState(ApigeeAlphaEnvgroupStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))
	sHostnames := make([]string, len(resource.Hostnames))
	for i, r := range resource.Hostnames {
		sHostnames[i] = r
	}
	p.SetHostnames(sHostnames)

	return p
}

// applyEnvgroup handles the gRPC request by passing it to the underlying Envgroup Apply() method.
func (s *EnvgroupServer) applyEnvgroup(ctx context.Context, c *alpha.Client, request *alphapb.ApplyApigeeAlphaEnvgroupRequest) (*alphapb.ApigeeAlphaEnvgroup, error) {
	p := ProtoToEnvgroup(request.GetResource())
	res, err := c.ApplyEnvgroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvgroupToProto(res)
	return r, nil
}

// applyApigeeAlphaEnvgroup handles the gRPC request by passing it to the underlying Envgroup Apply() method.
func (s *EnvgroupServer) ApplyApigeeAlphaEnvgroup(ctx context.Context, request *alphapb.ApplyApigeeAlphaEnvgroupRequest) (*alphapb.ApigeeAlphaEnvgroup, error) {
	cl, err := createConfigEnvgroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEnvgroup(ctx, cl, request)
}

// DeleteEnvgroup handles the gRPC request by passing it to the underlying Envgroup Delete() method.
func (s *EnvgroupServer) DeleteApigeeAlphaEnvgroup(ctx context.Context, request *alphapb.DeleteApigeeAlphaEnvgroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvgroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvgroup(ctx, ProtoToEnvgroup(request.GetResource()))

}

// ListApigeeAlphaEnvgroup handles the gRPC request by passing it to the underlying EnvgroupList() method.
func (s *EnvgroupServer) ListApigeeAlphaEnvgroup(ctx context.Context, request *alphapb.ListApigeeAlphaEnvgroupRequest) (*alphapb.ListApigeeAlphaEnvgroupResponse, error) {
	cl, err := createConfigEnvgroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvgroup(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ApigeeAlphaEnvgroup
	for _, r := range resources.Items {
		rp := EnvgroupToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListApigeeAlphaEnvgroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEnvgroup(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
