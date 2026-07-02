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
// proto.service: google.cloud.dataplex.v1.CatalogService
// proto.message: google.cloud.dataplex.v1.MetadataJob

package mockdataplex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *CatalogService) GetMetadataJob(ctx context.Context, req *pb.GetMetadataJobRequest) (*pb.MetadataJob, error) {
	name, err := s.parseMetadataJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MetadataJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *CatalogService) CreateMetadataJob(ctx context.Context, req *pb.CreateMetadataJobRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/metadataJobs/%s", req.GetParent(), req.GetMetadataJobId())
	name, err := s.parseMetadataJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetMetadataJob())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = uuid.NewString()

	obj.Status = &pb.MetadataJob_Status{
		State:             pb.MetadataJob_Status_QUEUED,
		CompletionPercent: 0,
		UpdateTime:        timestamppb.New(now),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		nowFinished := time.Now()
		obj.UpdateTime = timestamppb.New(nowFinished)
		obj.Status = &pb.MetadataJob_Status{
			State:             pb.MetadataJob_Status_SUCCEEDED,
			CompletionPercent: 100,
			UpdateTime:        timestamppb.New(nowFinished),
		}

		if obj.GetImportSpec() != nil {
			obj.Result = &pb.MetadataJob_ImportResult{
				ImportResult: &pb.MetadataJob_ImportJobResult{
					CreatedEntries:   10,
					UpdatedEntries:   5,
					DeletedEntries:   0,
					UnchangedEntries: 100,
					RecreatedEntries: 0,
					UpdateTime:       timestamppb.New(nowFinished),
				},
			}
		} else if obj.GetExportSpec() != nil {
			obj.Result = &pb.MetadataJob_ExportResult{
				ExportResult: &pb.MetadataJob_ExportJobResult{
					ExportedEntries: 15,
				},
			}
		}

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		lroMetadata.EndTime = timestamppb.New(nowFinished)
		return obj, nil
	})
}

func (s *CatalogService) CancelMetadataJob(ctx context.Context, req *pb.CancelMetadataJobRequest) (*emptypb.Empty, error) {
	name, err := s.parseMetadataJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MetadataJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	obj.UpdateTime = timestamppb.New(now)
	obj.Status = &pb.MetadataJob_Status{
		State:             pb.MetadataJob_Status_CANCELED,
		CompletionPercent: 100,
		UpdateTime:        timestamppb.New(now),
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type metadataJobName struct {
	Project       *projects.ProjectData
	Location      string
	MetadataJobID string
}

func (n *metadataJobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/metadataJobs/%s", n.Project.ID, n.Location, n.MetadataJobID)
}

// parseMetadataJobName parses a string into an metadataJobName.
// The expected form is `projects/*/locations/*/metadataJobs/*`.
func (s *MockService) parseMetadataJobName(name string) (*metadataJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "metadataJobs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &metadataJobName{
			Project:       project,
			Location:      tokens[3],
			MetadataJobID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
