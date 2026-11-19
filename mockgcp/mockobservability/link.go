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

package mockobservability

import (
	"context"
	"strings"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/observability/apiv1/observabilitypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type datasetName struct {
	Project  *projects.ProjectData
	Location string
	Bucket   string
	Dataset  string
}

func (n *datasetName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/buckets/" + n.Bucket + "/datasets/" + n.Dataset
}

func (n *datasetName) LROPrefix() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
}

func (s *observabilityServer) parseDatasetName(parent string) (*datasetName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "datasets" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &datasetName{
			Project:  project,
			Location: tokens[3],
			Bucket:   tokens[5],
			Dataset:  tokens[7],
		}, nil
	}
	return nil, status.Errorf(
		codes.InvalidArgument,
		"parent %q is not valid. expected format is projects/<projectID>/locations/<location>/buckets/<bucket>/datasets/<dataset>",
		parent,
	)
}

type linkName struct {
	Project  *projects.ProjectData
	Location string
	Bucket   string
	Dataset  string
	LinkID   string
}

func (n *linkName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/buckets/" + n.Bucket + "/datasets/" + n.Dataset + "/links/" + n.LinkID
}

func (n *linkName) LROPrefix() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
}

func (s *observabilityServer) parseLinkName(name string) (*linkName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "buckets" && tokens[6] == "datasets" && tokens[8] == "links" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &linkName{
			Project:  project,
			Location: tokens[3],
			Bucket:   tokens[5],
			Dataset:  tokens[7],
			LinkID:   tokens[9],
		}, nil
	}
	return nil, status.Errorf(
		codes.InvalidArgument,
		"name %q is not valid. expected format is projects/<projectID>/locations/<location>/buckets/<bucket>/datasets/<dataset>/links/<linkID>",
		name,
	)
}

func (s *observabilityServer) GetLink(ctx context.Context, req *pb.GetLinkRequest) (*pb.Link, error) {
	name, err := s.parseLinkName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Link{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%s does not exist", fqn)
		}
		return nil, err
	}

	res := proto.Clone(obj).(*pb.Link)
	return res, nil
}

func (s *observabilityServer) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*longrunningpb.Operation, error) {
	parent, err := s.parseDatasetName(req.GetParent())
	if err != nil {
		return nil, err
	}

	linkID := req.GetLinkId()
	if linkID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "linkId is required")
	}

	if req.GetLink() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "link is required")
	}

	name := &linkName{
		Project:  parent.Project,
		Location: parent.Location,
		Bucket:   parent.Bucket,
		Dataset:  parent.Dataset,
		LinkID:   linkID,
	}
	fqn := name.String()

	now := timestamppb.Now()

	obj := proto.Clone(req.GetLink()).(*pb.Link)
	obj.Name = fqn
	obj.CreateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		Target:                fqn,
		Verb:                  "create",
		RequestedCancellation: false,
	}

	return s.operations.StartLRO(ctx, parent.LROPrefix(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		result := proto.Clone(obj).(*pb.Link)
		result.CreateTime = nil
		return result, nil
	})
}

func (s *observabilityServer) UpdateLink(ctx context.Context, req *pb.UpdateLinkRequest) (*longrunningpb.Operation, error) {
	if req.GetLink() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "link is required")
	}

	name, err := s.parseLinkName(req.GetLink().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Link{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%s does not exist", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.Link)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated = proto.Clone(req.GetLink()).(*pb.Link)
		updated.Name = fqn
		updated.CreateTime = existing.CreateTime
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				updated.Description = req.GetLink().GetDescription()
			case "displayName", "display_name":
				updated.DisplayName = req.GetLink().GetDisplayName()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q not supported for update", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		Target:                fqn,
		Verb:                  "update",
		RequestedCancellation: false,
	}

	return s.operations.StartLRO(ctx, name.LROPrefix(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return proto.Clone(updated).(*pb.Link), nil
	})
}

func (s *observabilityServer) DeleteLink(ctx context.Context, req *pb.DeleteLinkRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseLinkName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Link{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%s does not exist", fqn)
		}
		return nil, err
	}

	now := timestamppb.Now()
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            now,
		Target:                fqn,
		Verb:                  "delete",
		RequestedCancellation: false,
	}

	return s.operations.StartLRO(ctx, name.LROPrefix(), metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *observabilityServer) ListLinks(ctx context.Context, req *pb.ListLinksRequest) (*pb.ListLinksResponse, error) {
	parent, err := s.parseDatasetName(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/links/"
	var items []*pb.Link

	findKind := (&pb.Link{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := proto.Clone(obj).(*pb.Link)
		items = append(items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListLinksResponse{
		Links: items,
	}, nil
}
