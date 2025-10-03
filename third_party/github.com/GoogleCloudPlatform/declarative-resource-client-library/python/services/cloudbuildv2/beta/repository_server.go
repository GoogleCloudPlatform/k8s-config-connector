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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudbuildv2/beta/cloudbuildv2_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudbuildv2/beta"
)

// RepositoryServer implements the gRPC interface for Repository.
type RepositoryServer struct{}

// ProtoToRepository converts a Repository resource from its proto representation.
func ProtoToRepository(p *betapb.Cloudbuildv2BetaRepository) *beta.Repository {
	obj := &beta.Repository{
		Name:       dcl.StringOrNil(p.GetName()),
		RemoteUri:  dcl.StringOrNil(p.GetRemoteUri()),
		CreateTime: dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime: dcl.StringOrNil(p.GetUpdateTime()),
		Etag:       dcl.StringOrNil(p.GetEtag()),
		Project:    dcl.StringOrNil(p.GetProject()),
		Location:   dcl.StringOrNil(p.GetLocation()),
		Connection: dcl.StringOrNil(p.GetConnection()),
	}
	return obj
}

// RepositoryToProto converts a Repository resource to its proto representation.
func RepositoryToProto(resource *beta.Repository) *betapb.Cloudbuildv2BetaRepository {
	p := &betapb.Cloudbuildv2BetaRepository{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetRemoteUri(dcl.ValueOrEmptyString(resource.RemoteUri))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetConnection(dcl.ValueOrEmptyString(resource.Connection))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyRepository handles the gRPC request by passing it to the underlying Repository Apply() method.
func (s *RepositoryServer) applyRepository(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudbuildv2BetaRepositoryRequest) (*betapb.Cloudbuildv2BetaRepository, error) {
	p := ProtoToRepository(request.GetResource())
	res, err := c.ApplyRepository(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RepositoryToProto(res)
	return r, nil
}

// applyCloudbuildv2BetaRepository handles the gRPC request by passing it to the underlying Repository Apply() method.
func (s *RepositoryServer) ApplyCloudbuildv2BetaRepository(ctx context.Context, request *betapb.ApplyCloudbuildv2BetaRepositoryRequest) (*betapb.Cloudbuildv2BetaRepository, error) {
	cl, err := createConfigRepository(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRepository(ctx, cl, request)
}

// DeleteRepository handles the gRPC request by passing it to the underlying Repository Delete() method.
func (s *RepositoryServer) DeleteCloudbuildv2BetaRepository(ctx context.Context, request *betapb.DeleteCloudbuildv2BetaRepositoryRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRepository(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRepository(ctx, ProtoToRepository(request.GetResource()))

}

// ListCloudbuildv2BetaRepository handles the gRPC request by passing it to the underlying RepositoryList() method.
func (s *RepositoryServer) ListCloudbuildv2BetaRepository(ctx context.Context, request *betapb.ListCloudbuildv2BetaRepositoryRequest) (*betapb.ListCloudbuildv2BetaRepositoryResponse, error) {
	cl, err := createConfigRepository(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRepository(ctx, request.GetProject(), request.GetLocation(), request.GetConnection())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.Cloudbuildv2BetaRepository
	for _, r := range resources.Items {
		rp := RepositoryToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudbuildv2BetaRepositoryResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRepository(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
