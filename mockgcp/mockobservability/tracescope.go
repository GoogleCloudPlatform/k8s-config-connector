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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/observability/apiv1/observabilitypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type observabilityServer struct {
	*MockService
	pb.UnimplementedObservabilityServer
}

type parentName struct {
	Project  *projects.ProjectData
	Location string
}

func (n *parentName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location
}

func (s *observabilityServer) parseParent(parent string) (*parentName, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &parentName{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(
		codes.InvalidArgument,
		"parent %q is not valid. expected format is projects/<projectID>/locations/<location>",
		parent,
	)
}

type traceScopeName struct {
	Project      *projects.ProjectData
	Location     string
	TraceScopeID string
}

func (n *traceScopeName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/traceScopes/" + n.TraceScopeID
}

func (s *observabilityServer) parseTraceScopeName(name string) (*traceScopeName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "traceScopes" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &traceScopeName{
			Project:      project,
			Location:     tokens[3],
			TraceScopeID: tokens[5],
		}, nil
	}
	return nil, status.Errorf(
		codes.InvalidArgument,
		"name %q is not valid. expected format is projects/<projectID>/locations/<location>/traceScopes/<traceScopeID>",
		name,
	)
}

func populateDefaultsForTraceScope(obj *pb.TraceScope) {
	// TraceScope currently has no optional proto3 scalar fields requiring default initialization.
	// Any future proto3 zero-value gap should be documented here.
}

func (s *observabilityServer) GetTraceScope(ctx context.Context, req *pb.GetTraceScopeRequest) (*pb.TraceScope, error) {
	name, err := s.parseTraceScopeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TraceScope{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%s does not exist", fqn)
		}
		return nil, err
	}

	res := proto.Clone(obj).(*pb.TraceScope)
	populateDefaultsForTraceScope(res)
	return res, nil
}

func (s *observabilityServer) CreateTraceScope(ctx context.Context, req *pb.CreateTraceScopeRequest) (*pb.TraceScope, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	traceScopeID := req.GetTraceScopeId()
	if traceScopeID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "traceScopeId is required")
	}

	if req.GetTraceScope() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "traceScope is required")
	}

	name := &traceScopeName{
		Project:      parent.Project,
		Location:     parent.Location,
		TraceScopeID: traceScopeID,
	}
	fqn := name.String()

	now := timestamppb.Now()

	obj := proto.Clone(req.GetTraceScope()).(*pb.TraceScope)
	obj.Name = fqn
	obj.CreateTime = now
	obj.UpdateTime = now

	populateDefaultsForTraceScope(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return proto.Clone(obj).(*pb.TraceScope), nil
}

func (s *observabilityServer) UpdateTraceScope(ctx context.Context, req *pb.UpdateTraceScopeRequest) (*pb.TraceScope, error) {
	if req.GetTraceScope() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "traceScope is required")
	}

	name, err := s.parseTraceScopeName(req.GetTraceScope().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.TraceScope{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%s does not exist", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(existing).(*pb.TraceScope)

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated = proto.Clone(req.GetTraceScope()).(*pb.TraceScope)
		updated.Name = fqn
		updated.CreateTime = existing.CreateTime
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				updated.Description = req.GetTraceScope().GetDescription()
			case "resource_names", "resourceNames":
				updated.ResourceNames = req.GetTraceScope().GetResourceNames()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q not supported for update", path)
			}
		}
	}

	updated.UpdateTime = timestamppb.Now()
	populateDefaultsForTraceScope(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return proto.Clone(updated).(*pb.TraceScope), nil
}

func (s *observabilityServer) DeleteTraceScope(ctx context.Context, req *pb.DeleteTraceScopeRequest) (*emptypb.Empty, error) {
	name, err := s.parseTraceScopeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TraceScope{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%s does not exist", fqn)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *observabilityServer) ListTraceScopes(ctx context.Context, req *pb.ListTraceScopesRequest) (*pb.ListTraceScopesResponse, error) {
	parent, err := s.parseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/traceScopes/"
	var items []*pb.TraceScope

	findKind := (&pb.TraceScope{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := proto.Clone(obj).(*pb.TraceScope)
		populateDefaultsForTraceScope(item)
		items = append(items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListTraceScopesResponse{
		TraceScopes: items,
	}, nil
}
