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

package mockcompute

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type InstantSnapshotGroupsV1 struct {
	*MockService
}

type RegionInstantSnapshotGroupsV1 struct {
	*MockService
}

type zonalInstantSnapshotGroupName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalInstantSnapshotGroupName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/instantSnapshotGroups/" + n.Name
}

func (s *MockService) parseZonalInstantSnapshotGroupName(name string) (*zonalInstantSnapshotGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "instantSnapshotGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &zonalInstantSnapshotGroupName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

type regionalInstantSnapshotGroupName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalInstantSnapshotGroupName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/instantSnapshotGroups/" + n.Name
}

func (s *MockService) parseRegionalInstantSnapshotGroupName(name string) (*regionalInstantSnapshotGroupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "instantSnapshotGroups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &regionalInstantSnapshotGroupName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *InstantSnapshotGroupsV1) Get(ctx context.Context, req *computepb.GetInstantSnapshotGroupRequest) (*computepb.InstantSnapshotGroup, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/instantSnapshotGroups/%s", req.GetProject(), req.GetZone(), req.GetInstantSnapshotGroup())
	name, err := s.parseZonalInstantSnapshotGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &computepb.InstantSnapshotGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *InstantSnapshotGroupsV1) Insert(ctx context.Context, req *computepb.InsertInstantSnapshotGroupRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/instantSnapshotGroups/%s", req.GetProject(), req.GetZone(), req.GetInstantSnapshotGroupResource().GetName())
	name, err := s.parseZonalInstantSnapshotGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetInstantSnapshotGroupResource()).(*computepb.InstantSnapshotGroup)
	selfLink := BuildComputeSelfLink(ctx, fqn)
	obj.SelfLink = PtrTo(selfLink)
	obj.SelfLinkWithId = PtrTo(selfLink)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instantSnapshotGroup")
	obj.Status = PtrTo("READY")
	obj.Zone = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/zones/%s", name.Project.ID, name.Zone)))

	if obj.SourceConsistencyGroup != nil && *obj.SourceConsistencyGroup != "" {
		scg := *obj.SourceConsistencyGroup
		if !strings.HasPrefix(scg, "https://") {
			scg = BuildComputeSelfLink(ctx, strings.TrimPrefix(scg, "/"))
		}
		obj.SourceConsistencyGroup = PtrTo(scg)

		tokens := strings.Split(scg, "/")
		policyID := tokens[len(tokens)-1]

		obj.ResourceStatus = &computepb.InstantSnapshotGroupResourceStatus{
			ConsistencyMembershipResolutionTime: PtrTo(s.nowString()),
			SourceInfo: &computepb.InstantSnapshotGroupSourceInfo{
				ConsistencyGroup:   PtrTo(scg),
				ConsistencyGroupId: PtrTo(policyID),
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *InstantSnapshotGroupsV1) Delete(ctx context.Context, req *computepb.DeleteInstantSnapshotGroupRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/instantSnapshotGroups/%s", req.GetProject(), req.GetZone(), req.GetInstantSnapshotGroup())
	name, err := s.parseZonalInstantSnapshotGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &computepb.InstantSnapshotGroup{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *RegionInstantSnapshotGroupsV1) Get(ctx context.Context, req *computepb.GetRegionInstantSnapshotGroupRequest) (*computepb.InstantSnapshotGroup, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/instantSnapshotGroups/%s", req.GetProject(), req.GetRegion(), req.GetInstantSnapshotGroup())
	name, err := s.parseRegionalInstantSnapshotGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &computepb.InstantSnapshotGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *RegionInstantSnapshotGroupsV1) Insert(ctx context.Context, req *computepb.InsertRegionInstantSnapshotGroupRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/instantSnapshotGroups/%s", req.GetProject(), req.GetRegion(), req.GetInstantSnapshotGroupResource().GetName())
	name, err := s.parseRegionalInstantSnapshotGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetInstantSnapshotGroupResource()).(*computepb.InstantSnapshotGroup)
	selfLink := BuildComputeSelfLink(ctx, fqn)
	obj.SelfLink = PtrTo(selfLink)
	obj.SelfLinkWithId = PtrTo(selfLink)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#instantSnapshotGroup")
	obj.Status = PtrTo("READY")
	obj.Region = PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)))

	if obj.SourceConsistencyGroup != nil && *obj.SourceConsistencyGroup != "" {
		scg := *obj.SourceConsistencyGroup
		if !strings.HasPrefix(scg, "https://") {
			scg = BuildComputeSelfLink(ctx, strings.TrimPrefix(scg, "/"))
		}
		obj.SourceConsistencyGroup = PtrTo(scg)

		tokens := strings.Split(scg, "/")
		policyID := tokens[len(tokens)-1]

		obj.ResourceStatus = &computepb.InstantSnapshotGroupResourceStatus{
			ConsistencyMembershipResolutionTime: PtrTo(s.nowString()),
			SourceInfo: &computepb.InstantSnapshotGroupSourceInfo{
				ConsistencyGroup:   PtrTo(scg),
				ConsistencyGroupId: PtrTo(policyID),
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RegionInstantSnapshotGroupsV1) Delete(ctx context.Context, req *computepb.DeleteRegionInstantSnapshotGroupRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/regions/%s/instantSnapshotGroups/%s", req.GetProject(), req.GetRegion(), req.GetInstantSnapshotGroup())
	name, err := s.parseRegionalInstantSnapshotGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &computepb.InstantSnapshotGroup{}
	if err := s.storage.Get(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startRegionalLRO(ctx, name.Project.ID, name.Region, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func RegisterInstantSnapshotGroupsHandler(ctx context.Context, mux *runtime.ServeMux, s *InstantSnapshotGroupsV1) error {
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// GET /compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups/{instant_snapshot_group}
	getPath := "/compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups/{instant_snapshot_group}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.GetInstantSnapshotGroupRequest{
			Project:              pathParams["project"],
			Zone:                 pathParams["zone"],
			InstantSnapshotGroup: pathParams["instant_snapshot_group"],
		}

		obj, err := s.Get(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, obj, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// POST /compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups
	insertPath := "/compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups"
	if err := mux.HandlePath("POST", insertPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.InsertInstantSnapshotGroupRequest{
			Project: pathParams["project"],
			Zone:    pathParams["zone"],
		}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		bodyBytes, err := io.ReadAll(newReader())
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		resource := &computepb.InstantSnapshotGroup{}
		if err := protojson.Unmarshal(bodyBytes, resource); err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.InstantSnapshotGroupResource = resource

		op, err := s.Insert(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// DELETE /compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups/{instant_snapshot_group}
	deletePath := "/compute/v1/projects/{project}/zones/{zone}/instantSnapshotGroups/{instant_snapshot_group}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.DeleteInstantSnapshotGroupRequest{
			Project:              pathParams["project"],
			Zone:                 pathParams["zone"],
			InstantSnapshotGroup: pathParams["instant_snapshot_group"],
		}

		op, err := s.Delete(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	return nil
}

func RegisterRegionInstantSnapshotGroupsHandler(ctx context.Context, mux *runtime.ServeMux, s *RegionInstantSnapshotGroupsV1) error {
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// GET /compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups/{instant_snapshot_group}
	getPath := "/compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups/{instant_snapshot_group}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.GetRegionInstantSnapshotGroupRequest{
			Project:              pathParams["project"],
			Region:               pathParams["region"],
			InstantSnapshotGroup: pathParams["instant_snapshot_group"],
		}

		obj, err := s.Get(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, obj, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// POST /compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups
	insertPath := "/compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups"
	if err := mux.HandlePath("POST", insertPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.InsertRegionInstantSnapshotGroupRequest{
			Project: pathParams["project"],
			Region:  pathParams["region"],
		}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", berr))
			return
		}
		bodyBytes, err := io.ReadAll(newReader())
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		resource := &computepb.InstantSnapshotGroup{}
		if err := protojson.Unmarshal(bodyBytes, resource); err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.InstantSnapshotGroupResource = resource

		op, err := s.Insert(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// DELETE /compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups/{instant_snapshot_group}
	deletePath := "/compute/v1/projects/{project}/regions/{region}/instantSnapshotGroups/{instant_snapshot_group}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.DeleteRegionInstantSnapshotGroupRequest{
			Project:              pathParams["project"],
			Region:               pathParams["region"],
			InstantSnapshotGroup: pathParams["instant_snapshot_group"],
		}

		op, err := s.Delete(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	return nil
}
