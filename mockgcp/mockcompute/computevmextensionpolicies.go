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

type ZoneVmExtensionPoliciesV1 struct {
	*MockService
}

type zonalVMExtensionPolicyName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

func (n *zonalVMExtensionPolicyName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/vmExtensionPolicies/" + n.Name
}

func (s *MockService) parseZonalVMExtensionPolicyName(name string) (*zonalVMExtensionPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "zones" && tokens[4] == "vmExtensionPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &zonalVMExtensionPolicyName{
			Project: project,
			Zone:    tokens[3],
			Name:    tokens[5],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ZoneVmExtensionPoliciesV1) Get(ctx context.Context, req *computepb.GetZoneVmExtensionPolicyRequest) (*computepb.VmExtensionPolicy, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/vmExtensionPolicies/%s", req.GetProject(), req.GetZone(), req.GetVmExtensionPolicy())
	name, err := s.parseZonalVMExtensionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &computepb.VmExtensionPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ZoneVmExtensionPoliciesV1) Insert(ctx context.Context, req *computepb.InsertZoneVmExtensionPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/vmExtensionPolicies/%s", req.GetProject(), req.GetZone(), req.GetVmExtensionPolicyResource().GetName())
	name, err := s.parseZonalVMExtensionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetVmExtensionPolicyResource()).(*computepb.VmExtensionPolicy)
	selfLink := BuildComputeSelfLink(ctx, fqn)
	obj.SelfLink = PtrTo(selfLink)
	obj.SelfLinkWithId = PtrTo(selfLink)
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.UpdateTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#vmExtensionPolicy")
	obj.State = PtrTo("ACTIVE")
	if obj.Description == nil {
		obj.Description = PtrTo("")
	}
	if obj.Priority == nil {
		obj.Priority = PtrTo(int32(1000))
	}
	for _, ep := range obj.ExtensionPolicies {
		if ep != nil && ep.PinnedVersion == nil {
			ep.PinnedVersion = PtrTo("")
		}
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
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ZoneVmExtensionPoliciesV1) Update(ctx context.Context, req *computepb.UpdateZoneVmExtensionPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/vmExtensionPolicies/%s", req.GetProject(), req.GetZone(), req.GetVmExtensionPolicy())
	name, err := s.parseZonalVMExtensionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &computepb.VmExtensionPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// In Compute Engine API, PUT/PATCH replaces fields provided in the request
	newResource := req.GetVmExtensionPolicyResource()
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
		if ep != nil && ep.PinnedVersion == nil {
			ep.PinnedVersion = PtrTo("")
		}
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
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *ZoneVmExtensionPoliciesV1) Delete(ctx context.Context, req *computepb.DeleteZoneVmExtensionPolicyRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/zones/%s/vmExtensionPolicies/%s", req.GetProject(), req.GetZone(), req.GetVmExtensionPolicy())
	name, err := s.parseZonalVMExtensionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deleted := &computepb.VmExtensionPolicy{}
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
	return s.startZonalLRO(ctx, name.Project.ID, name.Zone, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func RegisterZoneVmExtensionPoliciesHandler(ctx context.Context, mux *runtime.ServeMux, s *ZoneVmExtensionPoliciesV1) error {
	forwardResponseOptions := mux.GetForwardResponseOptions()

	// GET /compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}
	getPath := "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}"
	if err := mux.HandlePath("GET", getPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.GetZoneVmExtensionPolicyRequest{
			Project:           pathParams["project"],
			Zone:              pathParams["zone"],
			VmExtensionPolicy: pathParams["vm_extension_policy"],
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

	// POST /compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies
	insertPath := "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies"
	if err := mux.HandlePath("POST", insertPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.InsertZoneVmExtensionPolicyRequest{
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
		resource := &computepb.VmExtensionPolicy{}
		if err := protojson.Unmarshal(bodyBytes, resource); err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.VmExtensionPolicyResource = resource

		op, err := s.Insert(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// PATCH /compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}
	updatePath := "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}"
	if err := mux.HandlePath("PATCH", updatePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.UpdateZoneVmExtensionPolicyRequest{
			Project:           pathParams["project"],
			Zone:              pathParams["zone"],
			VmExtensionPolicy: pathParams["vm_extension_policy"],
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
		resource := &computepb.VmExtensionPolicy{}
		if err := protojson.Unmarshal(bodyBytes, resource); err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, status.Errorf(codes.InvalidArgument, "%v", err))
			return
		}
		protoReq.VmExtensionPolicyResource = resource

		op, err := s.Update(reqCtx, protoReq)
		if err != nil {
			runtime.HTTPError(reqCtx, mux, outboundMarshaler, w, r, err)
			return
		}

		runtime.ForwardResponseMessage(reqCtx, mux, outboundMarshaler, w, r, op, forwardResponseOptions...)
	}); err != nil {
		return err
	}

	// DELETE /compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}
	deletePath := "/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies/{vm_extension_policy}"
	if err := mux.HandlePath("DELETE", deletePath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		reqCtx := r.Context()
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, r)
		protoReq := &computepb.DeleteZoneVmExtensionPolicyRequest{
			Project:           pathParams["project"],
			Zone:              pathParams["zone"],
			VmExtensionPolicy: pathParams["vm_extension_policy"],
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
