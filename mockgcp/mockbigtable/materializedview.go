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
// proto.service: google.bigtable.admin.v2.BigtableInstanceAdmin
// proto.message: google.bigtable.admin.v2.MaterializedView

package mockbigtable

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

func (s *instanceAdminServer) GetMaterializedView(ctx context.Context, req *pb.GetMaterializedViewRequest) (*pb.MaterializedView, error) {
	name, err := s.parseMaterializedViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.MaterializedView{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to read: projects/{%v}/instances/%s/materializedViews/%s", name.Project.Number, name.InstanceName, name.MaterializedView)
	}

	return obj, nil
}

func (s *instanceAdminServer) CreateMaterializedView(ctx context.Context, req *pb.CreateMaterializedViewRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/materializedViews/" + req.MaterializedViewId
	name, err := s.parseMaterializedViewName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := ProtoClone(req.MaterializedView)
	obj.Name = fqn
	obj.Etag = "abcdef0123A="

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.CreateMaterializedViewMetadata{
		StartTime:       timestamppb.Now(),
		OriginalRequest: &pb.CreateMaterializedViewRequest{},
	}
	zone := "us-central1-a"
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)

	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now().Add(5 * time.Second))
		return obj, nil
	})

}

func (s *instanceAdminServer) UpdateMaterializedView(ctx context.Context, req *pb.UpdateMaterializedViewRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseMaterializedViewName(req.GetMaterializedView().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.MaterializedView{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)

	// Required. The set of fields to update.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "deletion_protection":
			updated.DeletionProtection = req.MaterializedView.DeletionProtection
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	zone := "us-central1-a"
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)

	metadata := &pb.UpdateMaterializedViewMetadata{
		StartTime: timestamppb.Now(),
		OriginalRequest: &pb.UpdateMaterializedViewRequest{
			MaterializedView: &pb.MaterializedView{},
			UpdateMask:       &fieldmaskpb.FieldMask{},
		},
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now().Add(5 * time.Minute))
		return updated, nil
	})
}

func (s *instanceAdminServer) ListMaterializedViews(ctx context.Context, req *pb.ListMaterializedViewsRequest) (*pb.ListMaterializedViewsResponse, error) {
	instanceName, err := s.parseInstanceName(req.GetParent())
	if err != nil {
		return nil, err
	}

	materializedView, err := s.listMaterializedViewsForInstance(ctx, instanceName)
	if err != nil {
		return nil, err
	}

	response := &pb.ListMaterializedViewsResponse{}
	response.MaterializedViews = materializedView

	return response, nil
}

func (s *instanceAdminServer) listMaterializedViewsForInstance(ctx context.Context, instanceName *instanceName) ([]*pb.MaterializedView, error) {
	if instanceName.InstanceName == "-" {
		return nil, fmt.Errorf("mock does not implement ListMaterializedViews for wildcard instances")
	}

	var response []*pb.MaterializedView

	findKind := (&pb.MaterializedView{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: instanceName.String() + "/materializedViews/",
	}, func(obj proto.Message) error {
		materializedView := obj.(*pb.MaterializedView)
		response = append(response, materializedView)
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Name < response[j].Name
	})

	return response, nil
}

func (s *instanceAdminServer) DeleteMaterializedView(ctx context.Context, req *pb.DeleteMaterializedViewRequest) (*emptypb.Empty, error) {
	name, err := s.parseMaterializedViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.MaterializedView{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type materializedViewName struct {
	instanceName
	MaterializedView string
}

func (n *materializedViewName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/materializedViews/%s", n.Project.ID, n.InstanceName, n.MaterializedView)
}

// parseMaterializedViewName parses a string into a materializedViewName.
// The expected form is `projects/*/instances/*/materializedViews/*`.
func (s *instanceAdminServer) parseMaterializedViewName(name string) (*materializedViewName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "materializedViews" {
		instanceName, err := s.parseInstanceName(strings.Join(tokens[0:4], "/"))
		if err != nil {
			return nil, err
		}

		name := &materializedViewName{
			instanceName:     *instanceName,
			MaterializedView: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
