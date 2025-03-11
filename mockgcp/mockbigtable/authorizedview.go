// Copyright 2025 Google LLC
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
// proto.service: google.bigtable.admin.v2.BigtableTableAdmin
// proto.message: google.bigtable.admin.v2.AuthorizedView

package mockbigtable

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

func (s *tableAdminServer) GetAuthorizedView(ctx context.Context, req *pb.GetAuthorizedViewRequest) (*pb.AuthorizedView, error) {
	name, err := s.parseAuthorizedViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AuthorizedView{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "authorizedView %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *tableAdminServer) CreateAuthorizedView(ctx context.Context, req *pb.CreateAuthorizedViewRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/authorizedViews/" + req.GetAuthorizedViewId()
	name, err := s.parseAuthorizedViewName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.AuthorizedView).(*pb.AuthorizedView)
	obj.Name = fqn
	obj.Etag = "abcdef0123A="

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreateAuthorizedViewMetadata{
		RequestTime:     timestamppb.New(time.Now()),
		OriginalRequest: &pb.CreateAuthorizedViewRequest{},
	}
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), "us-east1-c")
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.FinishTime = timestamppb.New(time.Now().Add(5 * time.Minute))
		return obj, nil
	})
}

func (s *tableAdminServer) UpdateAuthorizedView(ctx context.Context, req *pb.UpdateAuthorizedViewRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAuthorizedViewName(req.GetAuthorizedView().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.AuthorizedView{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.AuthorizedView)

	// Required. The set of fields to update.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "subset_view":
			updated.AuthorizedView = &pb.AuthorizedView_SubsetView_{
				SubsetView: &pb.AuthorizedView_SubsetView{
					RowPrefixes:   req.GetAuthorizedView().GetSubsetView().GetRowPrefixes(),
					FamilySubsets: req.GetAuthorizedView().GetSubsetView().GetFamilySubsets(),
				},
			}
		case "deletion_protection":
			updated.DeletionProtection = false
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.UpdateAuthorizedViewMetadata{
		RequestTime:     timestamppb.New(time.Now()),
		OriginalRequest: req,
		FinishTime:      timestamppb.New(time.Now().Add(5 * time.Minute)),
	}
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), "us-east1-c")
	return s.operations.DoneLRO(ctx, prefix, metadata, updated)
}

func (s *tableAdminServer) DeleteAuthorizedView(ctx context.Context, req *pb.DeleteAuthorizedViewRequest) (*emptypb.Empty, error) {
	name, err := s.parseAuthorizedViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.AuthorizedView{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type authorizedViewName struct {
	Project        string
	InstanceID     string
	TableID        string
	AuthorizedView string
}

func (n *authorizedViewName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/tables/%s/authorizedViews/%s", n.Project, n.InstanceID, n.TableID, n.AuthorizedView)
}

// parseAuthorizedViewName parses a string into a authorizedViewName.
// The expected form is `projects/*/instances/*/tables/*/authorizedViews/*`.
func (s *tableAdminServer) parseAuthorizedViewName(name string) (*authorizedViewName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "tables" && tokens[6] == "authorizedViews" {
		name := &authorizedViewName{
			Project:        tokens[1],
			InstanceID:     tokens[3],
			TableID:        tokens[5],
			AuthorizedView: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
