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
	"strings"
	"time"

	biglakepb "cloud.google.com/go/biglake/apiv1/biglakepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type icebergCatalogService struct {
	*MockService
	biglakepb.UnimplementedIcebergCatalogServiceServer
}

type icebergCatalogName struct {
	Project   *projects.ProjectData
	CatalogID string
}

func (n *icebergCatalogName) String() string {
	return "projects/" + n.Project.ID + "/catalogs/" + n.CatalogID
}

func (s *MockService) parseIcebergCatalogName(name string) (*icebergCatalogName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "catalogs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &icebergCatalogName{
			Project:   project,
			CatalogID: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *icebergCatalogService) CreateIcebergCatalog(ctx context.Context, req *biglakepb.CreateIcebergCatalogRequest) (*biglakepb.IcebergCatalog, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}
	if req.GetIcebergCatalogId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "iceberg-catalog-id is required")
	}

	name := &icebergCatalogName{
		Project:   project,
		CatalogID: req.GetIcebergCatalogId(),
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.Clone(req.GetIcebergCatalog()).(*biglakepb.IcebergCatalog)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if obj.CatalogType == biglakepb.IcebergCatalog_CATALOG_TYPE_UNSPECIFIED {
		obj.CatalogType = biglakepb.IcebergCatalog_CATALOG_TYPE_BIGLAKE
	}
	if obj.CredentialMode == biglakepb.IcebergCatalog_CREDENTIAL_MODE_UNSPECIFIED {
		obj.CredentialMode = biglakepb.IcebergCatalog_CREDENTIAL_MODE_END_USER
	}
	if obj.RestrictedLocationsConfig == nil || len(obj.RestrictedLocationsConfig.RestrictedLocations) == 0 {
		if obj.DefaultLocation != "" {
			obj.RestrictedLocationsConfig = &biglakepb.IcebergCatalog_RestrictedLocationsConfig{
				RestrictedLocations: []string{obj.DefaultLocation},
			}
		}
	}
	if len(obj.StorageRegions) == 0 {
		obj.StorageRegions = []string{"us-central1"}
	}
	if obj.CredentialMode == biglakepb.IcebergCatalog_CREDENTIAL_MODE_VENDED_CREDENTIALS {
		obj.BiglakeServiceAccount = fmt.Sprintf("blirc-%d-g8bvnavk@gcp-sa-biglakerestcatalog.iam.gserviceaccount.com", name.Project.Number)
		obj.BiglakeServiceAccountUniqueId = "102362709148677022604"
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *icebergCatalogService) GetIcebergCatalog(ctx context.Context, req *biglakepb.GetIcebergCatalogRequest) (*biglakepb.IcebergCatalog, error) {
	name, err := s.parseIcebergCatalogName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &biglakepb.IcebergCatalog{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Catalog %s does not exist.", name.CatalogID)
		}
		return nil, err
	}

	response := proto.Clone(obj).(*biglakepb.IcebergCatalog)
	if response.FederatedCatalogOptions == nil {
		response.FederatedCatalogOptions = &biglakepb.IcebergCatalog_FederatedCatalogOptions{}
	}
	if len(response.Replicas) == 0 {
		response.Replicas = []*biglakepb.IcebergCatalog_Replica{
			{Region: "us-central1"},
		}
	}

	return response, nil
}

func (s *icebergCatalogService) UpdateIcebergCatalog(ctx context.Context, req *biglakepb.UpdateIcebergCatalogRequest) (*biglakepb.IcebergCatalog, error) {
	if req.GetIcebergCatalog() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "icebergCatalog is required")
	}

	name, err := s.parseIcebergCatalogName(req.GetIcebergCatalog().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &biglakepb.IcebergCatalog{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Catalog %s does not exist.", name.CatalogID)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		paths = []string{"description", "credentialMode", "defaultLocation", "restrictedLocationsConfig", "federatedCatalogOptions"}
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetIcebergCatalog().GetDescription()
		case "credential_mode", "credentialMode":
			obj.CredentialMode = req.GetIcebergCatalog().GetCredentialMode()
			if obj.CredentialMode == biglakepb.IcebergCatalog_CREDENTIAL_MODE_VENDED_CREDENTIALS {
				obj.BiglakeServiceAccount = fmt.Sprintf("blirc-%d-g8bvnavk@gcp-sa-biglakerestcatalog.iam.gserviceaccount.com", name.Project.Number)
				obj.BiglakeServiceAccountUniqueId = "102362709148677022604"
			}
		case "default_location", "defaultLocation":
			obj.DefaultLocation = req.GetIcebergCatalog().GetDefaultLocation()
		case "restricted_locations_config", "restrictedLocationsConfig":
			obj.RestrictedLocationsConfig = req.GetIcebergCatalog().GetRestrictedLocationsConfig()
		case "federated_catalog_options", "federatedCatalogOptions":
			obj.FederatedCatalogOptions = req.GetIcebergCatalog().GetFederatedCatalogOptions()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q cannot be updated", path)
		}
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *icebergCatalogService) DeleteIcebergCatalog(ctx context.Context, req *biglakepb.DeleteIcebergCatalogRequest) (*emptypb.Empty, error) {
	name, err := s.parseIcebergCatalogName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &biglakepb.IcebergCatalog{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Catalog %s does not exist.", name.CatalogID)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *icebergCatalogService) ListIcebergCatalogs(ctx context.Context, req *biglakepb.ListIcebergCatalogsRequest) (*biglakepb.ListIcebergCatalogsResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 2 || tokens[0] != "projects" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := s.Projects.GetProjectByID(tokens[1])
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/catalogs/", project.ID)
	response := &biglakepb.ListIcebergCatalogsResponse{}
	catalogKind := (&biglakepb.IcebergCatalog{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, catalogKind, storage.ListOptions{Prefix: prefix}, func(m proto.Message) error {
		item := m.(*biglakepb.IcebergCatalog)
		response.IcebergCatalogs = append(response.IcebergCatalogs, item)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}
