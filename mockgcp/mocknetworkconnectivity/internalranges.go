// Copyright 2024 Google LLC
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

package mocknetworkconnectivity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
)

type internalRanges struct {
	*MockService
	pb.UnimplementedProjectsLocationsInternalRangesServerServer
}

func (r *internalRanges) GetProjectsLocationsInternalRange(ctx context.Context, req *pb.GetProjectsLocationsInternalRangeRequest) (*pb.InternalRange, error) {
	name, err := r.parseInternalRangeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.InternalRange{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *internalRanges) CreateProjectsLocationsInternalRange(ctx context.Context, req *pb.CreateProjectsLocationsInternalRangeRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/internalRanges/%s", req.GetParent(), req.GetInternalRangeId())
	name, err := r.parseInternalRangeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetProjectsLocationsInternalRange()).(*pb.InternalRange)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		if err := r.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (r *internalRanges) PatchProjectsLocationsInternalRange(ctx context.Context, req *pb.PatchProjectsLocationsInternalRangeRequest) (*longrunning.Operation, error) {
	log := klog.FromContext(ctx)

	reqName := req.GetName()

	name, err := r.parseInternalRangeName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.InternalRange{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.UpdateTime = timestamppb.New(now)

	if req.GetUpdateMask() != "" {
		paths := strings.Split(req.GetUpdateMask(), ",")

		patch := req.GetProjectsLocationsInternalRange()
		// TODO: Some sort of helper for fieldmask?
		for _, path := range paths {
			switch path {
			case "prefixLength":
				obj.PrefixLength = patch.PrefixLength

			default:
				log.Info("unsupported update_mask", "req", req)
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock", path)
			}
		}
	}

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (r *internalRanges) DeleteProjectsLocationsInternalRange(ctx context.Context, req *pb.DeleteProjectsLocationsInternalRangeRequest) (*longrunning.Operation, error) {
	name, err := r.parseInternalRangeName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	oldObj := &pb.InternalRange{}
	if err := r.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type internalRangeName struct {
	Project           *projects.ProjectData
	Location          string
	InternalRangeName string
}

func (n *internalRangeName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/internalRanges/" + n.InternalRangeName
}

// parseInternalRangeName parses a string into an internalRangeName.
// The expected form is `projects/*/locations/*/internalRanges/*`.
func (r *internalRanges) parseInternalRangeName(name string) (*internalRangeName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "internalRanges" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &internalRangeName{
			Project:           project,
			Location:          tokens[3],
			InternalRangeName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
