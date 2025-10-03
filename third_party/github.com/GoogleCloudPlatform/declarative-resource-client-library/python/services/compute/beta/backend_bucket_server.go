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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for BackendBucket.
type BackendBucketServer struct{}

// ProtoToBackendBucketCdnPolicy converts a BackendBucketCdnPolicy resource from its proto representation.
func ProtoToComputeBetaBackendBucketCdnPolicy(p *betapb.ComputeBetaBackendBucketCdnPolicy) *beta.BackendBucketCdnPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.BackendBucketCdnPolicy{
		SignedUrlCacheMaxAgeSec: dcl.Int64OrNil(p.SignedUrlCacheMaxAgeSec),
	}
	for _, r := range p.GetSignedUrlKeyNames() {
		obj.SignedUrlKeyNames = append(obj.SignedUrlKeyNames, r)
	}
	return obj
}

// ProtoToBackendBucket converts a BackendBucket resource from its proto representation.
func ProtoToBackendBucket(p *betapb.ComputeBetaBackendBucket) *beta.BackendBucket {
	obj := &beta.BackendBucket{
		BucketName:  dcl.StringOrNil(p.BucketName),
		CdnPolicy:   ProtoToComputeBetaBackendBucketCdnPolicy(p.GetCdnPolicy()),
		Description: dcl.StringOrNil(p.Description),
		EnableCdn:   dcl.Bool(p.EnableCdn),
		Name:        dcl.StringOrNil(p.Name),
		Project:     dcl.StringOrNil(p.Project),
		SelfLink:    dcl.StringOrNil(p.SelfLink),
	}
	return obj
}

// BackendBucketCdnPolicyToProto converts a BackendBucketCdnPolicy resource to its proto representation.
func ComputeBetaBackendBucketCdnPolicyToProto(o *beta.BackendBucketCdnPolicy) *betapb.ComputeBetaBackendBucketCdnPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendBucketCdnPolicy{
		SignedUrlCacheMaxAgeSec: dcl.ValueOrEmptyInt64(o.SignedUrlCacheMaxAgeSec),
	}
	for _, r := range o.SignedUrlKeyNames {
		p.SignedUrlKeyNames = append(p.SignedUrlKeyNames, r)
	}
	return p
}

// BackendBucketToProto converts a BackendBucket resource to its proto representation.
func BackendBucketToProto(resource *beta.BackendBucket) *betapb.ComputeBetaBackendBucket {
	p := &betapb.ComputeBetaBackendBucket{
		BucketName:  dcl.ValueOrEmptyString(resource.BucketName),
		CdnPolicy:   ComputeBetaBackendBucketCdnPolicyToProto(resource.CdnPolicy),
		Description: dcl.ValueOrEmptyString(resource.Description),
		EnableCdn:   dcl.ValueOrEmptyBool(resource.EnableCdn),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Project:     dcl.ValueOrEmptyString(resource.Project),
		SelfLink:    dcl.ValueOrEmptyString(resource.SelfLink),
	}

	return p
}

// ApplyBackendBucket handles the gRPC request by passing it to the underlying BackendBucket Apply() method.
func (s *BackendBucketServer) applyBackendBucket(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaBackendBucketRequest) (*betapb.ComputeBetaBackendBucket, error) {
	p := ProtoToBackendBucket(request.GetResource())
	res, err := c.ApplyBackendBucket(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BackendBucketToProto(res)
	return r, nil
}

// ApplyBackendBucket handles the gRPC request by passing it to the underlying BackendBucket Apply() method.
func (s *BackendBucketServer) ApplyComputeBetaBackendBucket(ctx context.Context, request *betapb.ApplyComputeBetaBackendBucketRequest) (*betapb.ComputeBetaBackendBucket, error) {
	cl, err := createConfigBackendBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBackendBucket(ctx, cl, request)
}

// DeleteBackendBucket handles the gRPC request by passing it to the underlying BackendBucket Delete() method.
func (s *BackendBucketServer) DeleteComputeBetaBackendBucket(ctx context.Context, request *betapb.DeleteComputeBetaBackendBucketRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBackendBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBackendBucket(ctx, ProtoToBackendBucket(request.GetResource()))

}

// ListComputeBetaBackendBucket handles the gRPC request by passing it to the underlying BackendBucketList() method.
func (s *BackendBucketServer) ListComputeBetaBackendBucket(ctx context.Context, request *betapb.ListComputeBetaBackendBucketRequest) (*betapb.ListComputeBetaBackendBucketResponse, error) {
	cl, err := createConfigBackendBucket(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBackendBucket(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaBackendBucket
	for _, r := range resources.Items {
		rp := BackendBucketToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaBackendBucketResponse{Items: protos}, nil
}

func createConfigBackendBucket(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
