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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudresourcemanager/beta/cloudresourcemanager_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager/beta"
)

// ProjectServer implements the gRPC interface for Project.
type ProjectServer struct{}

// ProtoToProjectLifecycleStateEnum converts a ProjectLifecycleStateEnum enum from its proto representation.
func ProtoToCloudresourcemanagerBetaProjectLifecycleStateEnum(e betapb.CloudresourcemanagerBetaProjectLifecycleStateEnum) *beta.ProjectLifecycleStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.CloudresourcemanagerBetaProjectLifecycleStateEnum_name[int32(e)]; ok {
		e := beta.ProjectLifecycleStateEnum(n[len("CloudresourcemanagerBetaProjectLifecycleStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToProject converts a Project resource from its proto representation.
func ProtoToProject(p *betapb.CloudresourcemanagerBetaProject) *beta.Project {
	obj := &beta.Project{
		LifecycleState: ProtoToCloudresourcemanagerBetaProjectLifecycleStateEnum(p.GetLifecycleState()),
		DisplayName:    dcl.StringOrNil(p.GetDisplayName()),
		Parent:         dcl.StringOrNil(p.GetParent()),
		Name:           dcl.StringOrNil(p.GetName()),
		ProjectNumber:  dcl.Int64OrNil(p.GetProjectNumber()),
	}
	return obj
}

// ProjectLifecycleStateEnumToProto converts a ProjectLifecycleStateEnum enum to its proto representation.
func CloudresourcemanagerBetaProjectLifecycleStateEnumToProto(e *beta.ProjectLifecycleStateEnum) betapb.CloudresourcemanagerBetaProjectLifecycleStateEnum {
	if e == nil {
		return betapb.CloudresourcemanagerBetaProjectLifecycleStateEnum(0)
	}
	if v, ok := betapb.CloudresourcemanagerBetaProjectLifecycleStateEnum_value["ProjectLifecycleStateEnum"+string(*e)]; ok {
		return betapb.CloudresourcemanagerBetaProjectLifecycleStateEnum(v)
	}
	return betapb.CloudresourcemanagerBetaProjectLifecycleStateEnum(0)
}

// ProjectToProto converts a Project resource to its proto representation.
func ProjectToProto(resource *beta.Project) *betapb.CloudresourcemanagerBetaProject {
	p := &betapb.CloudresourcemanagerBetaProject{}
	p.SetLifecycleState(CloudresourcemanagerBetaProjectLifecycleStateEnumToProto(resource.LifecycleState))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetProjectNumber(dcl.ValueOrEmptyInt64(resource.ProjectNumber))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyProject handles the gRPC request by passing it to the underlying Project Apply() method.
func (s *ProjectServer) applyProject(ctx context.Context, c *beta.Client, request *betapb.ApplyCloudresourcemanagerBetaProjectRequest) (*betapb.CloudresourcemanagerBetaProject, error) {
	p := ProtoToProject(request.GetResource())
	res, err := c.ApplyProject(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ProjectToProto(res)
	return r, nil
}

// applyCloudresourcemanagerBetaProject handles the gRPC request by passing it to the underlying Project Apply() method.
func (s *ProjectServer) ApplyCloudresourcemanagerBetaProject(ctx context.Context, request *betapb.ApplyCloudresourcemanagerBetaProjectRequest) (*betapb.CloudresourcemanagerBetaProject, error) {
	cl, err := createConfigProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyProject(ctx, cl, request)
}

// DeleteProject handles the gRPC request by passing it to the underlying Project Delete() method.
func (s *ProjectServer) DeleteCloudresourcemanagerBetaProject(ctx context.Context, request *betapb.DeleteCloudresourcemanagerBetaProjectRequest) (*emptypb.Empty, error) {

	cl, err := createConfigProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteProject(ctx, ProtoToProject(request.GetResource()))

}

// ListCloudresourcemanagerBetaProject handles the gRPC request by passing it to the underlying ProjectList() method.
func (s *ProjectServer) ListCloudresourcemanagerBetaProject(ctx context.Context, request *betapb.ListCloudresourcemanagerBetaProjectRequest) (*betapb.ListCloudresourcemanagerBetaProjectResponse, error) {
	cl, err := createConfigProject(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListProject(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.CloudresourcemanagerBetaProject
	for _, r := range resources.Items {
		rp := ProjectToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListCloudresourcemanagerBetaProjectResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigProject(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
