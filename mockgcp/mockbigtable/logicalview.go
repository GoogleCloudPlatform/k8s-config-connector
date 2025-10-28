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
// proto.message: google.bigtable.admin.v2.LogicalView

package mockbigtable

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/apimachinery/pkg/util/sets"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

func (s *instanceAdminServer) GetLogicalView(ctx context.Context, req *pb.GetLogicalViewRequest) (*pb.LogicalView, error) {
	name, err := s.parseLogicalViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.LogicalView{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%v not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *instanceAdminServer) CreateLogicalView(ctx context.Context, req *pb.CreateLogicalViewRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/logicalViews/" + req.LogicalViewId
	name, err := s.parseLogicalViewName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := ProtoClone(req.LogicalView)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	isAsync := false

	metadata := &pb.CreateLogicalViewMetadata{}
	zone := "us-central1-a" // TODO
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)

	lroRet := ProtoClone(obj)

	if isAsync {
		return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
			return lroRet, nil
		})
	} else {
		return s.operations.DoneLRO(ctx, prefix, metadata, lroRet)
	}
}

func (s *instanceAdminServer) UpdateLogicalView(ctx context.Context, req *pb.UpdateLogicalViewRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseLogicalViewName(req.GetLogicalView().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.LogicalView{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)

	// Required. The set of fields to update.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	isAsync := false

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "query":
			updated.Query = req.GetLogicalView().GetQuery()
		case "deletion_protection":
			updated.DeletionProtection = req.GetLogicalView().GetDeletionProtection()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.UpdateLogicalViewMetadata{}
	zone := "us-central1-a" // TODO
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)

	lroRet := ProtoClone(updated)
	updatePaths := sets.New(req.GetUpdateMask().GetPaths()...)
	// Only return in LRO whatever has actually been updated/changed.
	if !updatePaths.Has("query") {
		lroRet.Query = ""
	}
	if !updatePaths.Has("deletion_protection") {
		lroRet.DeletionProtection = false
	}

	if isAsync {
		return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
			return lroRet, nil
		})
	} else {
		return s.operations.DoneLRO(ctx, prefix, metadata, lroRet)
	}
}

func (s *instanceAdminServer) ListLogicalViews(ctx context.Context, req *pb.ListLogicalViewsRequest) (*pb.ListLogicalViewsResponse, error) {
	instanceName, err := s.parseInstanceName(req.GetParent())
	if err != nil {
		return nil, err
	}

	logicalView, err := s.listLogicalViewsForInstance(ctx, instanceName)
	if err != nil {
		return nil, err
	}

	response := &pb.ListLogicalViewsResponse{}
	response.LogicalViews = logicalView

	return response, nil
}

func (s *instanceAdminServer) listLogicalViewsForInstance(ctx context.Context, instanceName *instanceName) ([]*pb.LogicalView, error) {
	if instanceName.InstanceName == "-" {
		return nil, fmt.Errorf("mock does not implement ListLogicalViews for wildcard instances")
	}

	var response []*pb.LogicalView

	findKind := (&pb.LogicalView{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: instanceName.String() + "/logicalViews/",
	}, func(obj proto.Message) error {
		logicalView := obj.(*pb.LogicalView)
		response = append(response, logicalView)
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Name < response[j].Name
	})

	return response, nil
}

func (s *instanceAdminServer) DeleteLogicalView(ctx context.Context, req *pb.DeleteLogicalViewRequest) (*emptypb.Empty, error) {
	name, err := s.parseLogicalViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.LogicalView{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type logicalViewName struct {
	instanceName
	LogicalView string
}

func (n *logicalViewName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/logicalViews/%s", n.Project.ID, n.InstanceName, n.LogicalView)
}

// parseLogicalViewName parses a string into a logicalViewName.
// The expected form is `projects/*/instances/*/logicalViews/*`.
func (s *instanceAdminServer) parseLogicalViewName(name string) (*logicalViewName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "logicalViews" {
		instanceName, err := s.parseInstanceName(strings.Join(tokens[0:4], "/"))
		if err != nil {
			return nil, err
		}

		name := &logicalViewName{
			instanceName: *instanceName,
			LogicalView:  tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
