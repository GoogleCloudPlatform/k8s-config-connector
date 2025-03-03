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
	firebaserulespb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/firebaserules/firebaserules_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules"
)

// ReleaseServer implements the gRPC interface for Release.
type ReleaseServer struct{}

// ProtoToRelease converts a Release resource from its proto representation.
func ProtoToRelease(p *firebaserulespb.FirebaserulesRelease) *firebaserules.Release {
	obj := &firebaserules.Release{
		Name:        dcl.StringOrNil(p.GetName()),
		RulesetName: dcl.StringOrNil(p.GetRulesetName()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Disabled:    dcl.Bool(p.GetDisabled()),
		Project:     dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// ReleaseToProto converts a Release resource to its proto representation.
func ReleaseToProto(resource *firebaserules.Release) *firebaserulespb.FirebaserulesRelease {
	p := &firebaserulespb.FirebaserulesRelease{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetRulesetName(dcl.ValueOrEmptyString(resource.RulesetName))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDisabled(dcl.ValueOrEmptyBool(resource.Disabled))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))

	return p
}

// applyRelease handles the gRPC request by passing it to the underlying Release Apply() method.
func (s *ReleaseServer) applyRelease(ctx context.Context, c *firebaserules.Client, request *firebaserulespb.ApplyFirebaserulesReleaseRequest) (*firebaserulespb.FirebaserulesRelease, error) {
	p := ProtoToRelease(request.GetResource())
	res, err := c.ApplyRelease(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ReleaseToProto(res)
	return r, nil
}

// applyFirebaserulesRelease handles the gRPC request by passing it to the underlying Release Apply() method.
func (s *ReleaseServer) ApplyFirebaserulesRelease(ctx context.Context, request *firebaserulespb.ApplyFirebaserulesReleaseRequest) (*firebaserulespb.FirebaserulesRelease, error) {
	cl, err := createConfigRelease(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRelease(ctx, cl, request)
}

// DeleteRelease handles the gRPC request by passing it to the underlying Release Delete() method.
func (s *ReleaseServer) DeleteFirebaserulesRelease(ctx context.Context, request *firebaserulespb.DeleteFirebaserulesReleaseRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRelease(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRelease(ctx, ProtoToRelease(request.GetResource()))

}

// ListFirebaserulesRelease handles the gRPC request by passing it to the underlying ReleaseList() method.
func (s *ReleaseServer) ListFirebaserulesRelease(ctx context.Context, request *firebaserulespb.ListFirebaserulesReleaseRequest) (*firebaserulespb.ListFirebaserulesReleaseResponse, error) {
	cl, err := createConfigRelease(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRelease(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*firebaserulespb.FirebaserulesRelease
	for _, r := range resources.Items {
		rp := ReleaseToProto(r)
		protos = append(protos, rp)
	}
	p := &firebaserulespb.ListFirebaserulesReleaseResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRelease(ctx context.Context, service_account_file string) (*firebaserules.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return firebaserules.NewClient(conf), nil
}
