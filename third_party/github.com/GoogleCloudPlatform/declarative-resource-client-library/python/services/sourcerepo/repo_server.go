// Copyright 2021 Google LLC. All Rights Reserved.
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
	sourcerepopb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/sourcerepo/sourcerepo_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sourcerepo"
)

// Server implements the gRPC interface for Repo.
type RepoServer struct{}

// ProtoToRepoPubsubConfigs converts a RepoPubsubConfigs resource from its proto representation.
func ProtoToSourcerepoRepoPubsubConfigs(p *sourcerepopb.SourcerepoRepoPubsubConfigs) *sourcerepo.RepoPubsubConfigs {
	if p == nil {
		return nil
	}
	obj := &sourcerepo.RepoPubsubConfigs{
		Topic:               dcl.StringOrNil(p.Topic),
		MessageFormat:       dcl.StringOrNil(p.MessageFormat),
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
	}
	return obj
}

// ProtoToRepo converts a Repo resource from its proto representation.
func ProtoToRepo(p *sourcerepopb.SourcerepoRepo) *sourcerepo.Repo {
	obj := &sourcerepo.Repo{
		Name:    dcl.StringOrNil(p.Name),
		Size:    dcl.Int64OrNil(p.Size),
		Url:     dcl.StringOrNil(p.Url),
		Project: dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetPubsubConfigs() {
		obj.PubsubConfigs = append(obj.PubsubConfigs, *ProtoToSourcerepoRepoPubsubConfigs(r))
	}
	return obj
}

// RepoPubsubConfigsToProto converts a RepoPubsubConfigs resource to its proto representation.
func SourcerepoRepoPubsubConfigsToProto(o *sourcerepo.RepoPubsubConfigs) *sourcerepopb.SourcerepoRepoPubsubConfigs {
	if o == nil {
		return nil
	}
	p := &sourcerepopb.SourcerepoRepoPubsubConfigs{
		Topic:               dcl.ValueOrEmptyString(o.Topic),
		MessageFormat:       dcl.ValueOrEmptyString(o.MessageFormat),
		ServiceAccountEmail: dcl.ValueOrEmptyString(o.ServiceAccountEmail),
	}
	return p
}

// RepoToProto converts a Repo resource to its proto representation.
func RepoToProto(resource *sourcerepo.Repo) *sourcerepopb.SourcerepoRepo {
	p := &sourcerepopb.SourcerepoRepo{
		Name:    dcl.ValueOrEmptyString(resource.Name),
		Size:    dcl.ValueOrEmptyInt64(resource.Size),
		Url:     dcl.ValueOrEmptyString(resource.Url),
		Project: dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.PubsubConfigs {
		p.PubsubConfigs = append(p.PubsubConfigs, SourcerepoRepoPubsubConfigsToProto(&r))
	}

	return p
}

// ApplyRepo handles the gRPC request by passing it to the underlying Repo Apply() method.
func (s *RepoServer) applyRepo(ctx context.Context, c *sourcerepo.Client, request *sourcerepopb.ApplySourcerepoRepoRequest) (*sourcerepopb.SourcerepoRepo, error) {
	p := ProtoToRepo(request.GetResource())
	res, err := c.ApplyRepo(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RepoToProto(res)
	return r, nil
}

// ApplyRepo handles the gRPC request by passing it to the underlying Repo Apply() method.
func (s *RepoServer) ApplySourcerepoRepo(ctx context.Context, request *sourcerepopb.ApplySourcerepoRepoRequest) (*sourcerepopb.SourcerepoRepo, error) {
	cl, err := createConfigRepo(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyRepo(ctx, cl, request)
}

// DeleteRepo handles the gRPC request by passing it to the underlying Repo Delete() method.
func (s *RepoServer) DeleteSourcerepoRepo(ctx context.Context, request *sourcerepopb.DeleteSourcerepoRepoRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRepo(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRepo(ctx, ProtoToRepo(request.GetResource()))

}

// ListSourcerepoRepo handles the gRPC request by passing it to the underlying RepoList() method.
func (s *RepoServer) ListSourcerepoRepo(ctx context.Context, request *sourcerepopb.ListSourcerepoRepoRequest) (*sourcerepopb.ListSourcerepoRepoResponse, error) {
	cl, err := createConfigRepo(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRepo(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*sourcerepopb.SourcerepoRepo
	for _, r := range resources.Items {
		rp := RepoToProto(r)
		protos = append(protos, rp)
	}
	return &sourcerepopb.ListSourcerepoRepoResponse{Items: protos}, nil
}

func createConfigRepo(ctx context.Context, service_account_file string) (*sourcerepo.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return sourcerepo.NewClient(conf), nil
}
