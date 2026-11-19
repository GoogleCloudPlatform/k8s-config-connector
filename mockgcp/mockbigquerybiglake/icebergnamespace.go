// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockbigquerybiglake

import (
	"context"
	"fmt"
	"sort"
	"strings"

	biglakepb "cloud.google.com/go/biglake/apiv1/biglakepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type icebergNamespaceName struct {
	Project     *projects.ProjectData
	CatalogID   string
	NamespaceID string
}

func (n *icebergNamespaceName) String() string {
	return "projects/" + n.Project.ID + "/catalogs/" + n.CatalogID + "/namespaces/" + n.NamespaceID
}

func (s *MockService) parseIcebergNamespaceName(name string) (*icebergNamespaceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "catalogs" && tokens[4] == "namespaces" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &icebergNamespaceName{
			Project:     project,
			CatalogID:   tokens[3],
			NamespaceID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *icebergCatalogService) CreateIcebergNamespace(ctx context.Context, req *biglakepb.CreateIcebergNamespaceRequest) (*biglakepb.IcebergNamespace, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "catalogs" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	catalogFQN := "projects/" + project.ID + "/catalogs/" + tokens[3]
	catalog := &biglakepb.IcebergCatalog{}
	if err := s.storage.Get(ctx, catalogFQN, catalog); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Catalog %s does not exist.", tokens[3])
		}
		return nil, err
	}

	nsReq := req.GetIcebergNamespace()
	if nsReq == nil || len(nsReq.GetNamespace()) == 0 || nsReq.GetNamespace()[0] == "" {
		return nil, status.Errorf(codes.InvalidArgument, "namespace is required")
	}

	namespaceID := strings.Join(nsReq.GetNamespace(), "\x1f")
	name := &icebergNamespaceName{
		Project:     project,
		CatalogID:   tokens[3],
		NamespaceID: namespaceID,
	}

	fqn := name.String()

	storageObj := proto.Clone(nsReq).(*biglakepb.IcebergNamespace)
	if storageObj.Properties == nil {
		storageObj.Properties = make(map[string]string)
	}
	if _, ok := storageObj.Properties["location"]; !ok && catalog.DefaultLocation != "" {
		storageObj.Properties["location"] = strings.TrimSuffix(catalog.DefaultLocation, "/") + "/" + namespaceID
	}

	if err := s.storage.Create(ctx, fqn, storageObj); err != nil {
		return nil, err
	}

	return proto.Clone(nsReq).(*biglakepb.IcebergNamespace), nil
}

func (s *icebergCatalogService) GetIcebergNamespace(ctx context.Context, req *biglakepb.GetIcebergNamespaceRequest) (*biglakepb.IcebergNamespace, error) {
	name, err := s.parseIcebergNamespaceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &biglakepb.IcebergNamespace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Namespace does not exist: %s", name.NamespaceID)
		}
		return nil, err
	}

	return proto.Clone(obj).(*biglakepb.IcebergNamespace), nil
}

func (s *icebergCatalogService) UpdateIcebergNamespace(ctx context.Context, req *biglakepb.UpdateIcebergNamespaceRequest) (*biglakepb.UpdateIcebergNamespaceResponse, error) {
	name, err := s.parseIcebergNamespaceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &biglakepb.IcebergNamespace{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Namespace does not exist: %s", name.NamespaceID)
		}
		return nil, err
	}

	update := req.GetIcebergNamespaceUpdate()
	var removed []string
	var missing []string
	var addedOrUpdated []string

	if obj.Properties == nil {
		obj.Properties = make(map[string]string)
	}

	for _, k := range update.GetRemovals() {
		if _, ok := obj.Properties[k]; ok {
			delete(obj.Properties, k)
			removed = append(removed, k)
		} else {
			missing = append(missing, k)
		}
	}

	var updateKeys []string
	for k := range update.GetUpdates() {
		updateKeys = append(updateKeys, k)
	}
	sort.Strings(updateKeys)

	for _, k := range updateKeys {
		v := update.GetUpdates()[k]
		obj.Properties[k] = v
		addedOrUpdated = append(addedOrUpdated, k)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &biglakepb.UpdateIcebergNamespaceResponse{
		Updated: addedOrUpdated,
		Missing: missing,
		Removed: removed,
	}, nil
}

func (s *icebergCatalogService) DeleteIcebergNamespace(ctx context.Context, req *biglakepb.DeleteIcebergNamespaceRequest) (*emptypb.Empty, error) {
	name, err := s.parseIcebergNamespaceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &biglakepb.IcebergNamespace{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Namespace does not exist: %s", name.NamespaceID)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *icebergCatalogService) ListIcebergNamespaces(ctx context.Context, req *biglakepb.ListIcebergNamespacesRequest) (*biglakepb.ListIcebergNamespacesResponse, error) {
	parent := req.GetApiParent()
	if parent == "" {
		parent = req.GetParent()
	}
	tokens := strings.Split(parent, "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "catalogs" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/catalogs/%s/namespaces/", project.ID, tokens[3])
	response := &biglakepb.ListIcebergNamespacesResponse{}
	namespaceKind := (&biglakepb.IcebergNamespace{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, namespaceKind, storage.ListOptions{Prefix: prefix}, func(m proto.Message) error {
		item := m.(*biglakepb.IcebergNamespace)
		var values []*structpb.Value
		for _, part := range item.GetNamespace() {
			values = append(values, structpb.NewStringValue(part))
		}
		response.Namespaces = append(response.Namespaces, &structpb.ListValue{Values: values})
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
