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
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"cloud.google.com/go/compute/apiv1/computepb"
)

var allGlobalVmExtensionPolicyZones = []string{
	"africa-south1-a",
	"africa-south1-b",
	"africa-south1-c",
	"asia-east1-a",
	"asia-east1-b",
	"asia-east1-c",
	"asia-east2-a",
	"asia-east2-b",
	"asia-east2-c",
	"asia-northeast1-a",
	"asia-northeast1-b",
	"asia-northeast1-c",
	"asia-northeast2-a",
	"asia-northeast2-b",
	"asia-northeast2-c",
	"asia-northeast3-a",
	"asia-northeast3-b",
	"asia-northeast3-c",
	"asia-south1-a",
	"asia-south1-b",
	"asia-south1-c",
	"asia-south2-a",
	"asia-south2-b",
	"asia-south2-c",
	"asia-southeast1-a",
	"asia-southeast1-b",
	"asia-southeast1-c",
	"asia-southeast2-a",
	"asia-southeast2-b",
	"asia-southeast2-c",
	"asia-southeast3-a",
	"asia-southeast3-b",
	"asia-southeast3-c",
	"australia-southeast1-a",
	"australia-southeast1-b",
	"australia-southeast1-c",
	"australia-southeast2-a",
	"australia-southeast2-b",
	"australia-southeast2-c",
	"europe-central2-a",
	"europe-central2-b",
	"europe-central2-c",
	"europe-north1-a",
	"europe-north1-b",
	"europe-north1-c",
	"europe-north2-a",
	"europe-north2-b",
	"europe-north2-c",
	"europe-southwest1-a",
	"europe-southwest1-b",
	"europe-southwest1-c",
	"europe-west1-b",
	"europe-west1-c",
	"europe-west1-d",
	"europe-west10-a",
	"europe-west10-b",
	"europe-west10-c",
	"europe-west12-a",
	"europe-west12-b",
	"europe-west12-c",
	"europe-west2-a",
	"europe-west2-b",
	"europe-west2-c",
	"europe-west3-a",
	"europe-west3-b",
	"europe-west3-c",
	"europe-west4-a",
	"europe-west4-ai1a",
	"europe-west4-b",
	"europe-west4-c",
	"europe-west6-a",
	"europe-west6-b",
	"europe-west6-c",
	"europe-west8-a",
	"europe-west8-b",
	"europe-west8-c",
	"europe-west9-a",
	"europe-west9-b",
	"europe-west9-c",
	"me-central1-a",
	"me-central1-b",
	"me-central1-c",
	"me-west1-a",
	"me-west1-b",
	"me-west1-c",
	"northamerica-northeast1-a",
	"northamerica-northeast1-b",
	"northamerica-northeast1-c",
	"northamerica-northeast2-a",
	"northamerica-northeast2-b",
	"northamerica-northeast2-c",
	"northamerica-south1-a",
	"northamerica-south1-b",
	"northamerica-south1-c",
	"southamerica-east1-a",
	"southamerica-east1-b",
	"southamerica-east1-c",
	"southamerica-west1-a",
	"southamerica-west1-b",
	"southamerica-west1-c",
	"us-central1-a",
	"us-central1-b",
	"us-central1-c",
	"us-central1-f",
	"us-east1-b",
	"us-east1-c",
	"us-east1-d",
	"us-east4-a",
	"us-east4-b",
	"us-east4-c",
	"us-east5-a",
	"us-east5-b",
	"us-east5-c",
	"us-south1-a",
	"us-south1-ai1b",
	"us-south1-b",
	"us-south1-c",
	"us-west1-a",
	"us-west1-b",
	"us-west1-c",
	"us-west2-a",
	"us-west2-b",
	"us-west2-c",
	"us-west3-a",
	"us-west3-b",
	"us-west3-c",
	"us-west4-a",
	"us-west4-b",
	"us-west4-c",
}

type GlobalVmExtensionPoliciesV1 struct {
	*MockService
}

type globalVMExtensionPolicyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *globalVMExtensionPolicyName) String() string {
	return "projects/" + n.Project.ID + "/global/globalVmExtensionPolicies/" + n.Name
}

func (s *MockService) parseGlobalVMExtensionPolicyName(name string) (*globalVMExtensionPolicyName, error) {
	name = strings.TrimPrefix(name, "//compute.googleapis.com/")
	name = strings.TrimPrefix(name, "/")
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && (tokens[3] == "vmExtensionPolicies" || tokens[3] == "globalVmExtensionPolicies") {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &globalVMExtensionPolicyName{
			Project: project,
			Name:    tokens[4],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *GlobalVmExtensionPoliciesV1) buildInitialLocationRolloutStatus() map[string]*computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus {
	lrs := make(map[string]*computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus, len(allGlobalVmExtensionPolicyZones))
	for _, z := range allGlobalVmExtensionPolicyZones {
		lrs[z] = &computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus{
			State: PtrTo("LOCATION_ROLLOUT_STATE_NOT_STARTED"),
		}
	}
	return lrs
}

func (s *GlobalVmExtensionPoliciesV1) Get(ctx context.Context, req *computepb.GetGlobalVmExtensionPolicyRequest) (*computepb.GlobalVmExtensionPolicy, error) {
	reqName := fmt.Sprintf("projects/%s/global/vmExtensionPolicies/%s", req.GetProject(), req.GetGlobalVmExtensionPolicy())
	name, err := s.parseGlobalVMExtensionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &computepb.GlobalVmExtensionPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *GlobalVmExtensionPoliciesV1) Insert(ctx context.Context, req *computepb.InsertGlobalVmExtensionPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/vmExtensionPolicies/%s", req.GetProject(), req.GetGlobalVmExtensionPolicyResource().GetName())
	name, err := s.parseGlobalVMExtensionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetGlobalVmExtensionPolicyResource()).(*computepb.GlobalVmExtensionPolicy)
	selfLink := BuildComputeSelfLink(ctx, fqn)
	obj.SelfLink = PtrTo(selfLink)
	obj.SelfLinkWithId = PtrTo(selfLink)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.UpdateTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#globalVmExtensionPolicy")
	obj.ScopedResourceStatus = PtrTo("SCOPED_RESOURCE_STATUS_UNSPECIFIED")
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.Priority == nil {
		obj.Priority = PtrTo(int32(1000))
	}
	for _, ep := range obj.ExtensionPolicies {
		if ep != nil {
			if ep.PinnedVersion == nil {
				ep.PinnedVersion = PtrTo("")
			}
			if ep.StringConfig == nil {
				ep.StringConfig = PtrTo("")
			}
		}
	}

	plan := "FAST_ROLLOUT"
	if obj.RolloutOperation != nil && obj.RolloutOperation.RolloutInput != nil && obj.RolloutOperation.RolloutInput.PredefinedRolloutPlan != nil {
		plan = *obj.RolloutOperation.RolloutInput.PredefinedRolloutPlan
	}

	rolloutName := fmt.Sprintf("projects/%d/locations/global/rollouts/vm-extension-policy-insert-rollout-20240401-123456", name.Project.Number)
	rolloutPlan := fmt.Sprintf("projects/%d/locations/global/rolloutPlans/vm-policy-fast-rollout-plan", name.Project.Number)

	obj.RolloutOperation = &computepb.GlobalVmExtensionPolicyRolloutOperation{
		RolloutInput: &computepb.GlobalVmExtensionPolicyRolloutOperationRolloutInput{
			ConflictBehavior:      PtrTo(""),
			PredefinedRolloutPlan: PtrTo(plan),
			RetryUuid:             PtrTo(""),
		},
		RolloutStatus: &computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatus{
			CurrentRollouts: []*computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata{
				{
					LocationRolloutStatus: s.buildInitialLocationRolloutStatus(),
					Rollout:               PtrTo(rolloutName),
					RolloutPlan:           PtrTo(rolloutPlan),
					State:                 PtrTo("STATE_PROCESSING"),
				},
			},
		},
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      &id,
		TargetLink:    PtrTo(selfLink),
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalVmExtensionPoliciesV1) Update(ctx context.Context, req *computepb.UpdateGlobalVmExtensionPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/vmExtensionPolicies/%s", req.GetProject(), req.GetGlobalVmExtensionPolicy())
	name, err := s.parseGlobalVMExtensionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &computepb.GlobalVmExtensionPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	newResource := req.GetGlobalVmExtensionPolicyResource()
	if newResource.Description != nil {
		obj.Description = newResource.Description
	}
	if newResource.ExtensionPolicies != nil {
		obj.ExtensionPolicies = newResource.ExtensionPolicies
	}
	if newResource.InstanceSelectors != nil {
		obj.InstanceSelectors = newResource.InstanceSelectors
	}
	if newResource.Priority != nil {
		obj.Priority = newResource.Priority
	}
	obj.UpdateTimestamp = PtrTo(s.nowString())
	for _, ep := range obj.ExtensionPolicies {
		if ep != nil {
			if ep.PinnedVersion == nil {
				ep.PinnedVersion = PtrTo("")
			}
			if ep.StringConfig == nil {
				ep.StringConfig = PtrTo("")
			}
		}
	}

	plan := "FAST_ROLLOUT"
	if newResource.RolloutOperation != nil && newResource.RolloutOperation.RolloutInput != nil && newResource.RolloutOperation.RolloutInput.PredefinedRolloutPlan != nil {
		plan = *newResource.RolloutOperation.RolloutInput.PredefinedRolloutPlan
	}

	var previousRollout *computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata
	if obj.RolloutOperation != nil && obj.RolloutOperation.RolloutStatus != nil && len(obj.RolloutOperation.RolloutStatus.CurrentRollouts) > 0 {
		oldRollout := obj.RolloutOperation.RolloutStatus.CurrentRollouts[0]
		previousRollout = proto.Clone(oldRollout).(*computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata)
		previousRollout.State = PtrTo("STATE_CANCELLED")
	}

	rolloutName := fmt.Sprintf("projects/%d/locations/global/rollouts/vm-extension-policy-update-rollout-20240401-123456", name.Project.Number)
	rolloutPlan := fmt.Sprintf("projects/%d/locations/global/rolloutPlans/vm-policy-fast-rollout-plan", name.Project.Number)

	obj.RolloutOperation = &computepb.GlobalVmExtensionPolicyRolloutOperation{
		RolloutInput: &computepb.GlobalVmExtensionPolicyRolloutOperationRolloutInput{
			ConflictBehavior:      PtrTo(""),
			PredefinedRolloutPlan: PtrTo(plan),
			RetryUuid:             PtrTo(""),
		},
		RolloutStatus: &computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatus{
			CurrentRollouts: []*computepb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata{
				{
					LocationRolloutStatus: s.buildInitialLocationRolloutStatus(),
					Rollout:               PtrTo(rolloutName),
					RolloutPlan:           PtrTo(rolloutPlan),
					State:                 PtrTo("STATE_PROCESSING"),
				},
			},
			PreviousRollout: previousRollout,
		},
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *GlobalVmExtensionPoliciesV1) Delete(ctx context.Context, req *computepb.DeleteGlobalVmExtensionPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/vmExtensionPolicies/%s", req.GetProject(), req.GetGlobalVmExtensionPolicy())
	name, err := s.parseGlobalVMExtensionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &computepb.GlobalVmExtensionPolicy{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func RegisterGlobalVmExtensionPoliciesHandler(ctx context.Context, mux *runtime.ServeMux, s *GlobalVmExtensionPoliciesV1) error {
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// GET /compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}
	getPath := "/compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.GetGlobalVmExtensionPolicyRequest{
			Project:                 pathParams["project"],
			GlobalVmExtensionPolicy: pathParams["global_vm_extension_policy"],
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

	// POST /compute/v1/projects/{project}/global/vmExtensionPolicies
	insertPath := "/compute/v1/projects/{project}/global/vmExtensionPolicies"
	if err := mux.HandlePath("POST", insertPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.InsertGlobalVmExtensionPolicyRequest{
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
		resource := &computepb.GlobalVmExtensionPolicy{}
		if err := protojson.Unmarshal(bodyBytes, resource); err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.GlobalVmExtensionPolicyResource = resource

		op, err := s.Insert(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// PATCH /compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}
	updatePath := "/compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}"
	if err := mux.HandlePath("PATCH", updatePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.UpdateGlobalVmExtensionPolicyRequest{
			Project:                 pathParams["project"],
			GlobalVmExtensionPolicy: pathParams["global_vm_extension_policy"],
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
		resource := &computepb.GlobalVmExtensionPolicy{}
		if err := protojson.Unmarshal(bodyBytes, resource); err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.GlobalVmExtensionPolicyResource = resource

		op, err := s.Update(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// POST /compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}/delete
	deletePath := "/compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}/delete"
	handleDelete := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.DeleteGlobalVmExtensionPolicyRequest{
			Project:                 pathParams["project"],
			GlobalVmExtensionPolicy: pathParams["global_vm_extension_policy"],
		}

		newReader, berr := utilities.IOReaderFactory(r.Body)
		if berr == nil {
			bodyBytes, err := io.ReadAll(newReader())
			if err == nil && len(bodyBytes) > 0 {
				input := &computepb.GlobalVmExtensionPolicyRolloutOperationRolloutInput{}
				if err := protojson.Unmarshal(bodyBytes, input); err == nil {
					protoReq.GlobalVmExtensionPolicyRolloutOperationRolloutInputResource = input
				}
			}
		}

		op, err := s.Delete(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}

	if err := mux.HandlePath("POST", deletePath, handleDelete); err != nil {
		return err
	}
	if err := mux.HandlePath("DELETE", "/compute/v1/projects/{project}/global/vmExtensionPolicies/{global_vm_extension_policy}", handleDelete); err != nil {
		return err
	}

	return nil
}
