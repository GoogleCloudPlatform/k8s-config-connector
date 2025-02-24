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
	cloudresourcemanagerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/cloudresourcemanager/cloudresourcemanager_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/cloudresourcemanager"
)

// FolderServer implements the gRPC interface for Folder.
type FolderServer struct{}

// ProtoToFolderStateEnum converts a FolderStateEnum enum from its proto representation.
func ProtoToCloudresourcemanagerFolderStateEnum(e cloudresourcemanagerpb.CloudresourcemanagerFolderStateEnum) *cloudresourcemanager.FolderStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := cloudresourcemanagerpb.CloudresourcemanagerFolderStateEnum_name[int32(e)]; ok {
		e := cloudresourcemanager.FolderStateEnum(n[len("CloudresourcemanagerFolderStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToFolder converts a Folder resource from its proto representation.
func ProtoToFolder(p *cloudresourcemanagerpb.CloudresourcemanagerFolder) *cloudresourcemanager.Folder {
	obj := &cloudresourcemanager.Folder{
		Name:        dcl.StringOrNil(p.GetName()),
		Parent:      dcl.StringOrNil(p.GetParent()),
		DisplayName: dcl.StringOrNil(p.GetDisplayName()),
		State:       ProtoToCloudresourcemanagerFolderStateEnum(p.GetState()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		DeleteTime:  dcl.StringOrNil(p.GetDeleteTime()),
		Etag:        dcl.StringOrNil(p.GetEtag()),
	}
	return obj
}

// FolderStateEnumToProto converts a FolderStateEnum enum to its proto representation.
func CloudresourcemanagerFolderStateEnumToProto(e *cloudresourcemanager.FolderStateEnum) cloudresourcemanagerpb.CloudresourcemanagerFolderStateEnum {
	if e == nil {
		return cloudresourcemanagerpb.CloudresourcemanagerFolderStateEnum(0)
	}
	if v, ok := cloudresourcemanagerpb.CloudresourcemanagerFolderStateEnum_value["FolderStateEnum"+string(*e)]; ok {
		return cloudresourcemanagerpb.CloudresourcemanagerFolderStateEnum(v)
	}
	return cloudresourcemanagerpb.CloudresourcemanagerFolderStateEnum(0)
}

// FolderToProto converts a Folder resource to its proto representation.
func FolderToProto(resource *cloudresourcemanager.Folder) *cloudresourcemanagerpb.CloudresourcemanagerFolder {
	p := &cloudresourcemanagerpb.CloudresourcemanagerFolder{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetState(CloudresourcemanagerFolderStateEnumToProto(resource.State))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDeleteTime(dcl.ValueOrEmptyString(resource.DeleteTime))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))

	return p
}

// applyFolder handles the gRPC request by passing it to the underlying Folder Apply() method.
func (s *FolderServer) applyFolder(ctx context.Context, c *cloudresourcemanager.Client, request *cloudresourcemanagerpb.ApplyCloudresourcemanagerFolderRequest) (*cloudresourcemanagerpb.CloudresourcemanagerFolder, error) {
	p := ProtoToFolder(request.GetResource())
	res, err := c.ApplyFolder(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FolderToProto(res)
	return r, nil
}

// applyCloudresourcemanagerFolder handles the gRPC request by passing it to the underlying Folder Apply() method.
func (s *FolderServer) ApplyCloudresourcemanagerFolder(ctx context.Context, request *cloudresourcemanagerpb.ApplyCloudresourcemanagerFolderRequest) (*cloudresourcemanagerpb.CloudresourcemanagerFolder, error) {
	cl, err := createConfigFolder(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFolder(ctx, cl, request)
}

// DeleteFolder handles the gRPC request by passing it to the underlying Folder Delete() method.
func (s *FolderServer) DeleteCloudresourcemanagerFolder(ctx context.Context, request *cloudresourcemanagerpb.DeleteCloudresourcemanagerFolderRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFolder(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFolder(ctx, ProtoToFolder(request.GetResource()))

}

// ListCloudresourcemanagerFolder handles the gRPC request by passing it to the underlying FolderList() method.
func (s *FolderServer) ListCloudresourcemanagerFolder(ctx context.Context, request *cloudresourcemanagerpb.ListCloudresourcemanagerFolderRequest) (*cloudresourcemanagerpb.ListCloudresourcemanagerFolderResponse, error) {
	cl, err := createConfigFolder(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFolder(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*cloudresourcemanagerpb.CloudresourcemanagerFolder
	for _, r := range resources.Items {
		rp := FolderToProto(r)
		protos = append(protos, rp)
	}
	p := &cloudresourcemanagerpb.ListCloudresourcemanagerFolderResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFolder(ctx context.Context, service_account_file string) (*cloudresourcemanager.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return cloudresourcemanager.NewClient(conf), nil
}
