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
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/firebase/alpha/firebase_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebase/alpha"
)

// FirebaseProjectServer implements the gRPC interface for FirebaseProject.
type FirebaseProjectServer struct{}

// ProtoToFirebaseProjectStateEnum converts a FirebaseProjectStateEnum enum from its proto representation.
func ProtoToFirebaseAlphaFirebaseProjectStateEnum(e alphapb.FirebaseAlphaFirebaseProjectStateEnum) *alpha.FirebaseProjectStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FirebaseAlphaFirebaseProjectStateEnum_name[int32(e)]; ok {
		e := alpha.FirebaseProjectStateEnum(n[len("FirebaseAlphaFirebaseProjectStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFirebaseProjectResources converts a FirebaseProjectResources object from its proto representation.
func ProtoToFirebaseAlphaFirebaseProjectResources(p *alphapb.FirebaseAlphaFirebaseProjectResources) *alpha.FirebaseProjectResources {
	if p == nil {
		return nil
	}
	obj := &alpha.FirebaseProjectResources{
		HostingSite:              dcl.StringOrNil(p.GetHostingSite()),
		RealtimeDatabaseInstance: dcl.StringOrNil(p.GetRealtimeDatabaseInstance()),
		StorageBucket:            dcl.StringOrNil(p.GetStorageBucket()),
		LocationId:               dcl.StringOrNil(p.GetLocationId()),
	}
	return obj
}

// ProtoToFirebaseProject converts a FirebaseProject resource from its proto representation.
func ProtoToFirebaseProject(p *alphapb.FirebaseAlphaFirebaseProject) *alpha.FirebaseProject {
	obj := &alpha.FirebaseProject{
		ProjectId:     dcl.StringOrNil(p.GetProjectId()),
		ProjectNumber: dcl.Int64OrNil(p.GetProjectNumber()),
		DisplayName:   dcl.StringOrNil(p.GetDisplayName()),
		Resources:     ProtoToFirebaseAlphaFirebaseProjectResources(p.GetResources()),
		State:         ProtoToFirebaseAlphaFirebaseProjectStateEnum(p.GetState()),
		Project:       dcl.StringOrNil(p.GetProject()),
	}
	return obj
}

// FirebaseProjectStateEnumToProto converts a FirebaseProjectStateEnum enum to its proto representation.
func FirebaseAlphaFirebaseProjectStateEnumToProto(e *alpha.FirebaseProjectStateEnum) alphapb.FirebaseAlphaFirebaseProjectStateEnum {
	if e == nil {
		return alphapb.FirebaseAlphaFirebaseProjectStateEnum(0)
	}
	if v, ok := alphapb.FirebaseAlphaFirebaseProjectStateEnum_value["FirebaseProjectStateEnum"+string(*e)]; ok {
		return alphapb.FirebaseAlphaFirebaseProjectStateEnum(v)
	}
	return alphapb.FirebaseAlphaFirebaseProjectStateEnum(0)
}

// FirebaseProjectResourcesToProto converts a FirebaseProjectResources object to its proto representation.
func FirebaseAlphaFirebaseProjectResourcesToProto(o *alpha.FirebaseProjectResources) *alphapb.FirebaseAlphaFirebaseProjectResources {
	if o == nil {
		return nil
	}
	p := &alphapb.FirebaseAlphaFirebaseProjectResources{}
	p.SetHostingSite(dcl.ValueOrEmptyString(o.HostingSite))
	p.SetRealtimeDatabaseInstance(dcl.ValueOrEmptyString(o.RealtimeDatabaseInstance))
	p.SetStorageBucket(dcl.ValueOrEmptyString(o.StorageBucket))
	p.SetLocationId(dcl.ValueOrEmptyString(o.LocationId))
	return p
}

// FirebaseProjectToProto converts a FirebaseProject resource to its proto representation.
func FirebaseProjectToProto(resource *alpha.FirebaseProject) *alphapb.FirebaseAlphaFirebaseProject {
	p := &alphapb.FirebaseAlphaFirebaseProject{}
	p.SetProjectId(dcl.ValueOrEmptyString(resource.ProjectId))
	p.SetProjectNumber(dcl.ValueOrEmptyInt64(resource.ProjectNumber))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetResources(FirebaseAlphaFirebaseProjectResourcesToProto(resource.Resources))
	p.SetState(FirebaseAlphaFirebaseProjectStateEnumToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	mAnnotations := make(map[string]string, len(resource.Annotations))
	for k, r := range resource.Annotations {
		mAnnotations[k] = r
	}
	p.SetAnnotations(mAnnotations)

	return p
}

// applyFirebaseProject handles the gRPC request by passing it to the underlying FirebaseProject Apply() method.
func (s *FirebaseProjectServer) applyFirebaseProject(ctx context.Context, c *alpha.Client, request *alphapb.ApplyFirebaseAlphaFirebaseProjectRequest) (*alphapb.FirebaseAlphaFirebaseProject, error) {
	p := ProtoToFirebaseProject(request.GetResource())
	res, err := c.ApplyFirebaseProject(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FirebaseProjectToProto(res)
	return r, nil
}

// applyFirebaseAlphaFirebaseProject handles the gRPC request by passing it to the underlying FirebaseProject Apply() method.
func (s *FirebaseProjectServer) ApplyFirebaseAlphaFirebaseProject(ctx context.Context, request *alphapb.ApplyFirebaseAlphaFirebaseProjectRequest) (*alphapb.FirebaseAlphaFirebaseProject, error) {
	cl, err := createConfigFirebaseProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFirebaseProject(ctx, cl, request)
}

// DeleteFirebaseProject handles the gRPC request by passing it to the underlying FirebaseProject Delete() method.
func (s *FirebaseProjectServer) DeleteFirebaseAlphaFirebaseProject(ctx context.Context, request *alphapb.DeleteFirebaseAlphaFirebaseProjectRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for FirebaseProject")

}

// ListFirebaseAlphaFirebaseProject handles the gRPC request by passing it to the underlying FirebaseProjectList() method.
func (s *FirebaseProjectServer) ListFirebaseAlphaFirebaseProject(ctx context.Context, request *alphapb.ListFirebaseAlphaFirebaseProjectRequest) (*alphapb.ListFirebaseAlphaFirebaseProjectResponse, error) {
	cl, err := createConfigFirebaseProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFirebaseProject(ctx)
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.FirebaseAlphaFirebaseProject
	for _, r := range resources.Items {
		rp := FirebaseProjectToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListFirebaseAlphaFirebaseProjectResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFirebaseProject(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
