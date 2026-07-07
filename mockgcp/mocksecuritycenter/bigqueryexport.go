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

package mocksecuritycenter

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
)

type securityCenterServer struct {
	*MockService
	pb.UnimplementedSecurityCenterServer
}

func (s *securityCenterServer) GetBigQueryExport(ctx context.Context, req *pb.GetBigQueryExportRequest) (*pb.BigQueryExport, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	name = normalizeBigQueryExportName(name)

	obj := &pb.BigQueryExport{}
	if err := s.storage.Get(ctx, name, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *securityCenterServer) CreateBigQueryExport(ctx context.Context, req *pb.CreateBigQueryExportRequest) (*pb.BigQueryExport, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}
	exportID := req.GetBigQueryExportId()
	if exportID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "big_query_export_id is required")
	}

	parent = normalizeBigQueryExportParent(parent)
	fqn := parent + "/bigQueryExports/" + exportID

	obj := proto.Clone(req.GetBigQueryExport()).(*pb.BigQueryExport)
	obj.Name = fqn

	now := timestamppb.New(time.Now())
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.MostRecentEditor = "user@example.com"
	obj.Principal = "service-org-123456789@gcp-sa-scc.iam.gserviceaccount.com"

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *securityCenterServer) UpdateBigQueryExport(ctx context.Context, req *pb.UpdateBigQueryExportRequest) (*pb.BigQueryExport, error) {
	reqObj := req.GetBigQueryExport()
	if reqObj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "big_query_export is required")
	}
	fqn := reqObj.GetName()
	if fqn == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	fqn = normalizeBigQueryExportName(fqn)

	existing := &pb.BigQueryExport{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	// Just update the storage with the new object
	obj := proto.Clone(reqObj).(*pb.BigQueryExport)
	obj.Name = fqn
	obj.CreateTime = existing.CreateTime
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.MostRecentEditor = "user@example.com"
	obj.Principal = existing.Principal

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *securityCenterServer) DeleteBigQueryExport(ctx context.Context, req *pb.DeleteBigQueryExportRequest) (*emptypb.Empty, error) {
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	name = normalizeBigQueryExportName(name)

	obj := &pb.BigQueryExport{}
	if err := s.storage.Delete(ctx, name, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func normalizeBigQueryExportName(name string) string {
	if !strings.Contains(name, "/locations/") {
		if strings.HasPrefix(name, "organizations/") {
			parts := strings.Split(name, "/")
			if len(parts) >= 4 && parts[2] == "bigQueryExports" {
				return parts[0] + "/" + parts[1] + "/locations/global/bigQueryExports/" + strings.Join(parts[3:], "/")
			}
		}
		if strings.HasPrefix(name, "folders/") {
			parts := strings.Split(name, "/")
			if len(parts) >= 4 && parts[2] == "bigQueryExports" {
				return parts[0] + "/" + parts[1] + "/locations/global/bigQueryExports/" + strings.Join(parts[3:], "/")
			}
		}
		if strings.HasPrefix(name, "projects/") {
			parts := strings.Split(name, "/")
			if len(parts) >= 4 && parts[2] == "bigQueryExports" {
				return parts[0] + "/" + parts[1] + "/locations/global/bigQueryExports/" + strings.Join(parts[3:], "/")
			}
		}
	}
	return name
}

func normalizeBigQueryExportParent(parent string) string {
	if !strings.Contains(parent, "/locations/") {
		return parent + "/locations/global"
	}
	return parent
}
