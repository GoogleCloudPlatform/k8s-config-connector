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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type LegacyHttpHealthCheck struct {
	Kind               string  `json:"kind,omitempty"`
	ID                 *uint64 `json:"id,omitempty"`
	CreationTimestamp  *string `json:"creationTimestamp,omitempty"`
	Name               *string `json:"name,omitempty"`
	Description        *string `json:"description,omitempty"`
	Host               *string `json:"host,omitempty"`
	RequestPath        *string `json:"requestPath,omitempty"`
	Port               *int32  `json:"port,omitempty"`
	CheckIntervalSec   *int32  `json:"checkIntervalSec,omitempty"`
	TimeoutSec         *int32  `json:"timeoutSec,omitempty"`
	UnhealthyThreshold *int32  `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold   *int32  `json:"healthyThreshold,omitempty"`
	SelfLink           *string `json:"selfLink,omitempty"`
}

type errorItem struct {
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
}

type errorDetail struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []errorItem `json:"errors"`
}

type errorResponse struct {
	Error errorDetail `json:"error"`
}

func writeError(w http.ResponseWriter, code int, message string, reason string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := errorResponse{
		Error: errorDetail{
			Code:    code,
			Message: message,
			Errors: []errorItem{
				{
					Message: message,
					Domain:  "global",
					Reason:  reason,
				},
			},
		},
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func toLegacy(obj *pb.HealthCheck) *LegacyHttpHealthCheck {
	legacy := &LegacyHttpHealthCheck{
		Kind:               "compute#httpHealthCheck",
		ID:                 obj.Id,
		CreationTimestamp:  obj.CreationTimestamp,
		Name:               obj.Name,
		Description:        obj.Description,
		CheckIntervalSec:   obj.CheckIntervalSec,
		TimeoutSec:         obj.TimeoutSec,
		UnhealthyThreshold: obj.UnhealthyThreshold,
		HealthyThreshold:   obj.HealthyThreshold,
		SelfLink:           obj.SelfLink,
	}
	if obj.HttpHealthCheck != nil {
		legacy.Host = obj.HttpHealthCheck.Host
		legacy.Port = obj.HttpHealthCheck.Port
		legacy.RequestPath = obj.HttpHealthCheck.RequestPath
	}
	return legacy
}

func fromLegacy(legacy *LegacyHttpHealthCheck) *pb.HealthCheck {
	obj := &pb.HealthCheck{
		Kind:               PtrTo("compute#healthCheck"),
		Id:                 legacy.ID,
		CreationTimestamp:  legacy.CreationTimestamp,
		Name:               legacy.Name,
		Description:        legacy.Description,
		CheckIntervalSec:   legacy.CheckIntervalSec,
		TimeoutSec:         legacy.TimeoutSec,
		UnhealthyThreshold: legacy.UnhealthyThreshold,
		HealthyThreshold:   legacy.HealthyThreshold,
		SelfLink:           legacy.SelfLink,
		Type:               PtrTo("HTTP"),
	}
	if legacy.Host != nil || legacy.Port != nil || legacy.RequestPath != nil {
		obj.HttpHealthCheck = &pb.HTTPHealthCheck{
			Host:        legacy.Host,
			Port:        legacy.Port,
			RequestPath: legacy.RequestPath,
		}
	}
	return obj
}

func (s *MockService) handleHttpHealthChecks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	path := strings.TrimPrefix(r.URL.Path, "/compute/v1/")
	parts := strings.Split(path, "/")
	if len(parts) < 4 || parts[0] != "projects" || parts[2] != "global" || parts[3] != "httpHealthChecks" {
		writeError(w, http.StatusBadRequest, "invalid request path", "invalid")
		return
	}

	projectID := parts[1]
	var name string
	if len(parts) > 4 {
		name = parts[4]
	}

	switch r.Method {
	case http.MethodGet:
		if name == "" {
			s.listHttpHealthChecks(ctx, projectID, w, r)
		} else {
			s.getHttpHealthCheck(ctx, projectID, name, w, r)
		}
	case http.MethodPost:
		if name == "" {
			s.insertHttpHealthCheck(ctx, projectID, w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed", "methodNotAllowed")
		}
	case http.MethodPut:
		if name != "" {
			s.updateHttpHealthCheck(ctx, projectID, name, w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed", "methodNotAllowed")
		}
	case http.MethodPatch:
		if name != "" {
			s.patchHttpHealthCheck(ctx, projectID, name, w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed", "methodNotAllowed")
		}
	case http.MethodDelete:
		if name != "" {
			s.deleteHttpHealthCheck(ctx, projectID, name, w, r)
		} else {
			writeError(w, http.StatusMethodNotAllowed, "Method not allowed", "methodNotAllowed")
		}
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed", "methodNotAllowed")
	}
}

func (s *MockService) listHttpHealthChecks(ctx context.Context, projectID string, w http.ResponseWriter, r *http.Request) {
	findPrefix := "projects/" + projectID + "/global/httpHealthChecks/"

	type LegacyList struct {
		Kind     string                   `json:"kind"`
		ID       string                   `json:"id"`
		SelfLink string                   `json:"selfLink"`
		Items    []*LegacyHttpHealthCheck `json:"items"`
	}

	response := &LegacyList{
		Kind:     "compute#httpHealthCheckList",
		ID:       "projects/" + projectID + "/global/httpHealthChecks",
		SelfLink: BuildComputeSelfLink(ctx, "projects/"+projectID+"/global/httpHealthChecks"),
		Items:    []*LegacyHttpHealthCheck{},
	}

	kind := (&pb.HealthCheck{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		hc := obj.(*pb.HealthCheck)
		response.Items = append(response.Items, toLegacy(hc))
		return nil
	}); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error(), "internalError")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func (s *MockService) insertHttpHealthCheck(ctx context.Context, projectID string, w http.ResponseWriter, r *http.Request) {
	var legacy LegacyHttpHealthCheck
	if err := json.NewDecoder(r.Body).Decode(&legacy); err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "badRequest")
		return
	}

	if legacy.Name == nil || *legacy.Name == "" {
		writeError(w, http.StatusBadRequest, "field 'name' is required", "invalid")
		return
	}

	name := *legacy.Name
	fqn := "projects/" + projectID + "/global/httpHealthChecks/" + name

	// Check if already exists
	existing := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, existing); err == nil {
		writeError(w, http.StatusConflict, "The resource '"+fqn+"' already exists", "alreadyExists")
		return
	}

	id := s.generateID()
	legacy.ID = &id
	legacy.CreationTimestamp = PtrTo(s.nowString())
	legacy.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))

	obj := fromLegacy(&legacy)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error(), "internalError")
		return
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("insert"),
		User:          PtrTo("user@example.com"),
	}
	op, err := s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return obj, nil
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error(), "internalError")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	m := &runtime.JSONPb{}
	opBytes, _ := m.Marshal(op)
	_, _ = w.Write(opBytes)
}

func (s *MockService) getHttpHealthCheck(ctx context.Context, projectID string, name string, w http.ResponseWriter, r *http.Request) {
	fqn := "projects/" + projectID + "/global/httpHealthChecks/" + name

	obj := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		writeError(w, http.StatusNotFound, "The resource '"+fqn+"' was not found", "notFound")
		return
	}

	legacy := toLegacy(obj)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(legacy)
}

func (s *MockService) updateHttpHealthCheck(ctx context.Context, projectID string, name string, w http.ResponseWriter, r *http.Request) {
	fqn := "projects/" + projectID + "/global/httpHealthChecks/" + name

	existing := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		writeError(w, http.StatusNotFound, "The resource '"+fqn+"' was not found", "notFound")
		return
	}

	var legacy LegacyHttpHealthCheck
	if err := json.NewDecoder(r.Body).Decode(&legacy); err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "badRequest")
		return
	}

	legacy.ID = existing.Id
	legacy.CreationTimestamp = existing.CreationTimestamp
	legacy.SelfLink = existing.SelfLink
	legacy.Name = PtrTo(name)

	obj := fromLegacy(&legacy)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error(), "internalError")
		return
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	op, err := s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return obj, nil
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error(), "internalError")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	m := &runtime.JSONPb{}
	opBytes, _ := m.Marshal(op)
	_, _ = w.Write(opBytes)
}

func (s *MockService) patchHttpHealthCheck(ctx context.Context, projectID string, name string, w http.ResponseWriter, r *http.Request) {
	fqn := "projects/" + projectID + "/global/httpHealthChecks/" + name

	existing := &pb.HealthCheck{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		writeError(w, http.StatusNotFound, "The resource '"+fqn+"' was not found", "notFound")
		return
	}

	legacyExisting := toLegacy(existing)

	if err := json.NewDecoder(r.Body).Decode(&legacyExisting); err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "badRequest")
		return
	}

	legacyExisting.ID = existing.Id
	legacyExisting.CreationTimestamp = existing.CreationTimestamp
	legacyExisting.SelfLink = existing.SelfLink
	legacyExisting.Name = PtrTo(name)

	obj := fromLegacy(legacyExisting)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error(), "internalError")
		return
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("patch"),
		User:          PtrTo("user@example.com"),
	}
	op, err := s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return obj, nil
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error(), "internalError")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	m := &runtime.JSONPb{}
	opBytes, _ := m.Marshal(op)
	_, _ = w.Write(opBytes)
}

func (s *MockService) deleteHttpHealthCheck(ctx context.Context, projectID string, name string, w http.ResponseWriter, r *http.Request) {
	fqn := "projects/" + projectID + "/global/httpHealthChecks/" + name

	deleted := &pb.HealthCheck{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		writeError(w, http.StatusNotFound, "The resource '"+fqn+"' was not found", "notFound")
		return
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	op, err := s.startGlobalLRO(ctx, projectID, op, func() (proto.Message, error) {
		return deleted, nil
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error(), "internalError")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	m := &runtime.JSONPb{}
	opBytes, _ := m.Marshal(op)
	_, _ = w.Write(opBytes)
}
