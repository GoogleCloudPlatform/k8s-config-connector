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

// +tool:mockgcp-support
// proto.service: google.cloud.securitycenter.v1.SecurityCenter
// proto.message: google.cloud.securitycenter.v1.BigQueryExport

package mocksecuritycenter

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type bigQueryExportName struct {
	Organization string
	Folder       string
	Project      string
	Location     string
	ID           string
}

func (s *SecurityCenterServer) parseBigQueryExportName(name string) (*bigQueryExportName, error) {
	name = strings.TrimPrefix(name, "/")
	name = strings.TrimPrefix(name, "v1/")
	tokens := strings.Split(name, "/")
	// Expected formats:
	// organizations/{organization}/bigQueryExports/{export}
	// organizations/{organization}/locations/{location}/bigQueryExports/{export}
	// folders/{folder}/bigQueryExports/{export}
	// folders/{folder}/locations/{location}/bigQueryExports/{export}
	// projects/{project}/bigQueryExports/{export}
	// projects/{project}/locations/{location}/bigQueryExports/{export}
	if len(tokens) == 4 {
		switch tokens[0] {
		case "organizations":
			if tokens[2] == "bigQueryExports" {
				return &bigQueryExportName{
					Organization: tokens[1],
					ID:           tokens[3],
				}, nil
			}
		case "folders":
			if tokens[2] == "bigQueryExports" {
				return &bigQueryExportName{
					Folder: tokens[1],
					ID:     tokens[3],
				}, nil
			}
		case "projects":
			if tokens[2] == "bigQueryExports" {
				return &bigQueryExportName{
					Project: tokens[1],
					ID:      tokens[3],
				}, nil
			}
		}
	} else if len(tokens) == 6 {
		switch tokens[0] {
		case "organizations":
			if tokens[2] == "locations" && tokens[4] == "bigQueryExports" {
				return &bigQueryExportName{
					Organization: tokens[1],
					Location:     tokens[3],
					ID:           tokens[5],
				}, nil
			}
		case "folders":
			if tokens[2] == "locations" && tokens[4] == "bigQueryExports" {
				return &bigQueryExportName{
					Folder:   tokens[1],
					Location: tokens[3],
					ID:       tokens[5],
				}, nil
			}
		case "projects":
			if tokens[2] == "locations" && tokens[4] == "bigQueryExports" {
				return &bigQueryExportName{
					Project:  tokens[1],
					Location: tokens[3],
					ID:       tokens[5],
				}, nil
			}
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not a valid BigQuery export name", name)
}

func (n *bigQueryExportName) String() string {
	if n.Location != "" {
		if n.Organization != "" {
			return fmt.Sprintf("organizations/%s/locations/%s/bigQueryExports/%s", n.Organization, n.Location, n.ID)
		}
		if n.Folder != "" {
			return fmt.Sprintf("folders/%s/locations/%s/bigQueryExports/%s", n.Folder, n.Location, n.ID)
		}
		if n.Project != "" {
			return fmt.Sprintf("projects/%s/locations/%s/bigQueryExports/%s", n.Project, n.Location, n.ID)
		}
	}
	if n.Organization != "" {
		return fmt.Sprintf("organizations/%s/bigQueryExports/%s", n.Organization, n.ID)
	}
	if n.Folder != "" {
		return fmt.Sprintf("folders/%s/bigQueryExports/%s", n.Folder, n.ID)
	}
	if n.Project != "" {
		return fmt.Sprintf("projects/%s/bigQueryExports/%s", n.Project, n.ID)
	}
	return ""
}

func (s *SecurityCenterServer) GetBigQueryExport(ctx context.Context, req *pb.GetBigQueryExportRequest) (*pb.BigQueryExport, error) {
	name, err := s.parseBigQueryExportName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BigQueryExport{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterServer) CreateBigQueryExport(ctx context.Context, req *pb.CreateBigQueryExportRequest) (*pb.BigQueryExport, error) {
	parent := req.GetParent()
	var exportID string
	if req.GetBigQueryExportId() != "" {
		exportID = req.GetBigQueryExportId()
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "big_query_export_id is required")
	}

	reqName := fmt.Sprintf("%s/bigQueryExports/%s", parent, exportID)
	name, err := s.parseBigQueryExportName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetBigQueryExport()).(*pb.BigQueryExport)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()
	obj.MostRecentEditor = "mock-editor@gcp-mock.iam.gserviceaccount.com"
	if name.Organization != "" {
		obj.Principal = fmt.Sprintf("serviceAccount:service-org-%s@gcp-sa-scc-notification.iam.gserviceaccount.com", name.Organization)
	} else if name.Folder != "" {
		obj.Principal = fmt.Sprintf("serviceAccount:service-folder-%s@gcp-sa-scc-notification.iam.gserviceaccount.com", name.Folder)
	} else if name.Project != "" {
		obj.Principal = fmt.Sprintf("serviceAccount:service-project-%s@gcp-sa-scc-notification.iam.gserviceaccount.com", name.Project)
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterServer) UpdateBigQueryExport(ctx context.Context, req *pb.UpdateBigQueryExportRequest) (*pb.BigQueryExport, error) {
	reqName := req.GetBigQueryExport().GetName()

	name, err := s.parseBigQueryExportName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BigQueryExport{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply update mask
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.Description = req.GetBigQueryExport().GetDescription()
		obj.Filter = req.GetBigQueryExport().GetFilter()
		obj.Dataset = req.GetBigQueryExport().GetDataset()
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				obj.Description = req.GetBigQueryExport().GetDescription()
			case "filter":
				obj.Filter = req.GetBigQueryExport().GetFilter()
			case "dataset":
				obj.Dataset = req.GetBigQueryExport().GetDataset()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecurityCenterServer) DeleteBigQueryExport(ctx context.Context, req *pb.DeleteBigQueryExportRequest) (*emptypb.Empty, error) {
	name, err := s.parseBigQueryExportName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.BigQueryExport{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SecurityCenterServer) ListBigQueryExports(ctx context.Context, req *pb.ListBigQueryExportsRequest) (*pb.ListBigQueryExportsResponse, error) {
	response := &pb.ListBigQueryExportsResponse{}

	var prefix string
	if req.GetParent() != "" {
		prefix = req.GetParent() + "/bigQueryExports/"
	}

	findBigQueryExportKind := (&pb.BigQueryExport{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findBigQueryExportKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		bigQueryExport := obj.(*pb.BigQueryExport)
		response.BigQueryExports = append(response.BigQueryExports, bigQueryExport)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
