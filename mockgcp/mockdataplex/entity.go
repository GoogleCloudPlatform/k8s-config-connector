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
// proto.service: google.cloud.dataplex.v1.MetadataService
// proto.message: google.cloud.dataplex.v1.Entity

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *MetadataService) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.Entity, error) {
	parent, err := s.parseZoneName(req.Parent)
	if err != nil {
		return nil, err
	}

	if req.Entity == nil || req.Entity.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Entity.id is required")
	}

	name := &entityName{
		Project:  parent.Project,
		Location: parent.Location,
		LakeID:   parent.LakeID,
		ZoneID:   parent.ZoneID,
		EntityID: req.Entity.Id,
	}

	fqn := name.String()

	existing := &pb.Entity{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Entity already exists: %s", name.ProjectNumberString())
	}

	obj := proto.CloneOf(req.Entity)
	obj.Name = name.ProjectNumberString()
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = "111111111111111111111"
	obj.Etag = "abcdef0123A="
	obj.Access = &pb.StorageAccess{
		Read: pb.StorageAccess_DIRECT,
	}
	if obj.Compatibility == nil {
		obj.Compatibility = &pb.Entity_CompatibilityStatus{
			Bigquery: &pb.Entity_CompatibilityStatus_Compatibility{},
			HiveMetastore: &pb.Entity_CompatibilityStatus_Compatibility{
				Compatible: true,
			},
		}
	}
	s.populateFormatDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *MetadataService) GetEntity(ctx context.Context, req *pb.GetEntityRequest) (*pb.Entity, error) {
	name, err := s.parseEntityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Entity{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Name)
		}
		return nil, err
	}

	if req.View == pb.GetEntityRequest_BASIC {
		res := proto.CloneOf(obj)
		res.Schema = nil
		return res, nil
	}

	return obj, nil
}

func (s *MetadataService) UpdateEntity(ctx context.Context, req *pb.UpdateEntityRequest) (*pb.Entity, error) {
	if req.Entity == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Entity is required")
	}

	name, err := s.parseEntityName(req.Entity.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Entity{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Entity.Name)
		}
		return nil, err
	}

	if req.Entity.Etag != "" && req.Entity.Etag != existing.Etag {
		return nil, status.Errorf(codes.Aborted, "Resource has been modified; please retrieve the latest version and retry")
	}

	updated := proto.CloneOf(req.Entity)
	updated.Name = existing.Name
	updated.CreateTime = existing.CreateTime
	updated.UpdateTime = timestamppb.New(time.Now())
	updated.Uid = existing.Uid
	updated.Etag = "abcdef0123A="
	updated.Access = existing.Access
	updated.Compatibility = existing.Compatibility
	s.populateFormatDefaults(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *MetadataService) DeleteEntity(ctx context.Context, req *pb.DeleteEntityRequest) (*emptypb.Empty, error) {
	name, err := s.parseEntityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Entity{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", req.Name)
		}
		return nil, err
	}

	if req.Etag != "" && req.Etag != existing.Etag {
		return nil, status.Errorf(codes.Aborted, "Resource has been modified; please retrieve the latest version and retry")
	}

	deleted := &pb.Entity{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *MetadataService) populateFormatDefaults(obj *pb.Entity) {
	if obj.Format == nil {
		return
	}
	if strings.EqualFold(obj.Format.MimeType, "text/csv") {
		obj.Format.Format = pb.StorageFormat_CSV
		csv := obj.Format.GetCsv()
		if csv == nil {
			csv = &pb.StorageFormat_CsvOptions{
				Delimiter: ",",
				Encoding:  "UTF-8",
				Quote:     "\"",
			}
			obj.Format.Options = &pb.StorageFormat_Csv{Csv: csv}
		} else {
			if csv.Delimiter == "" {
				csv.Delimiter = ","
			}
			if csv.Encoding == "" {
				csv.Encoding = "UTF-8"
			}
			if csv.Quote == "" {
				csv.Quote = "\""
			}
		}
	}
}

type entityName struct {
	Project  *projects.ProjectData
	Location string
	LakeID   string
	ZoneID   string
	EntityID string
}

func (n *entityName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/entities/%s", n.Project.ID, n.Location, n.LakeID, n.ZoneID, n.EntityID)
}

func (n *entityName) ProjectNumberString() string {
	return fmt.Sprintf("projects/%d/locations/%s/lakes/%s/zones/%s/entities/%s", n.Project.Number, n.Location, n.LakeID, n.ZoneID, n.EntityID)
}

func (s *MockService) parseEntityName(name string) (*entityName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" && tokens[6] == "zones" && tokens[8] == "entities" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &entityName{
			Project:  project,
			Location: tokens[3],
			LakeID:   tokens[5],
			ZoneID:   tokens[7],
			EntityID: tokens[9],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (expected projects/*/locations/*/lakes/*/zones/*/entities/*)", name)
}
