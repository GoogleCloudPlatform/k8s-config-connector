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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"cloud.google.com/go/compute/apiv1/computepb"
)

type GlobalRolloutPlansV1 struct {
	*MockService
}

type globalRolloutPlanName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalRolloutPlanName) String() string {
	return "projects/" + n.Project.ID + "/global/rolloutPlans/" + n.Name
}

func (s *MockService) parseGlobalRolloutPlanName(name string) (*globalRolloutPlanName, error) {
	name = strings.TrimPrefix(name, "//compute.googleapis.com/")
	name = strings.TrimPrefix(name, "/")
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "rolloutPlans" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &globalRolloutPlanName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *GlobalRolloutPlansV1) populateRolloutPlanDefaults(obj *computepb.RolloutPlan) {
	if obj.Kind == nil {
		obj.Kind = PtrTo("compute#rolloutPlan")
	}
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	for i, wave := range obj.Waves {
		if wave != nil {
			if wave.DisplayName == nil {
				wave.DisplayName = PtrTo("")
			}
			if wave.Number == nil {
				wave.Number = PtrTo(int64(i + 1))
			}
		}
	}
}

func (s *GlobalRolloutPlansV1) Get(ctx context.Context, req *computepb.GetRolloutPlanRequest) (*computepb.RolloutPlan, error) {
	reqName := fmt.Sprintf("projects/%s/global/rolloutPlans/%s", req.GetProject(), req.GetRolloutPlan())
	name, err := s.parseGlobalRolloutPlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &computepb.RolloutPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}
	s.populateRolloutPlanDefaults(obj)

	return obj, nil
}

func (s *GlobalRolloutPlansV1) Insert(ctx context.Context, req *computepb.InsertRolloutPlanRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/rolloutPlans/%s", req.GetProject(), req.GetRolloutPlanResource().GetName())
	name, err := s.parseGlobalRolloutPlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetRolloutPlanResource()).(*computepb.RolloutPlan)
	selfLink := BuildComputeSelfLink(ctx, fqn)
	obj.SelfLink = PtrTo(selfLink)
	obj.SelfLinkWithId = PtrTo(selfLink)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	s.populateRolloutPlanDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalRolloutPlansV1) Delete(ctx context.Context, req *computepb.DeleteRolloutPlanRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/rolloutPlans/%s", req.GetProject(), req.GetRolloutPlan())
	name, err := s.parseGlobalRolloutPlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &computepb.RolloutPlan{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *GlobalRolloutPlansV1) List(ctx context.Context, req *computepb.ListRolloutPlansRequest) (*computepb.RolloutPlansListResponse, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	response := &computepb.RolloutPlansListResponse{
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/rolloutPlans", project.ID))),
	}

	rolloutPlanKind := (&computepb.RolloutPlan{}).ProtoReflect().Descriptor()
	prefix := fmt.Sprintf("projects/%s/global/rolloutPlans/", project.ID)
	if err := s.storage.List(ctx, rolloutPlanKind, storage.ListOptions{Prefix: prefix}, func(msg proto.Message) error {
		obj := msg.(*computepb.RolloutPlan)
		s.populateRolloutPlanDefaults(obj)
		response.Items = append(response.Items, obj)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func RegisterGlobalRolloutPlansHandler(ctx context.Context, mux *runtime.ServeMux, s *GlobalRolloutPlansV1) error {
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// GET /compute/v1/projects/{project}/global/rolloutPlans/{rollout_plan}
	getPath := "/compute/v1/projects/{project}/global/rolloutPlans/{rollout_plan}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.GetRolloutPlanRequest{
			Project:     pathParams["project"],
			RolloutPlan: pathParams["rollout_plan"],
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

	// POST /compute/v1/projects/{project}/global/rolloutPlans
	insertPath := "/compute/v1/projects/{project}/global/rolloutPlans"
	if err := mux.HandlePath("POST", insertPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.InsertRolloutPlanRequest{
			Project: pathParams["project"],
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
		resource := &computepb.RolloutPlan{}
		if err := protojson.Unmarshal(bodyBytes, resource); err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.RolloutPlanResource = resource

		op, err := s.Insert(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// DELETE /compute/v1/projects/{project}/global/rolloutPlans/{rollout_plan}
	deletePath := "/compute/v1/projects/{project}/global/rolloutPlans/{rollout_plan}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.DeleteRolloutPlanRequest{
			Project:     pathParams["project"],
			RolloutPlan: pathParams["rollout_plan"],
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

	// GET /compute/v1/projects/{project}/global/rolloutPlans
	listPath := "/compute/v1/projects/{project}/global/rolloutPlans"
	if err := mux.HandlePath("GET", listPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.ListRolloutPlansRequest{
			Project: pathParams["project"],
		}

		obj, err := s.List(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, obj, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	return nil
}
