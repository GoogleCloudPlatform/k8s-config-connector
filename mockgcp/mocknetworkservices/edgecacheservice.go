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

package mocknetworkservices

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	networkservicespb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/networkservices/v1"
)

func (s *NetworkServicesServer) GetEdgeCacheService(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	project := pathParams["project"]
	name := pathParams["name"]
	fqn := fmt.Sprintf("projects/%s/locations/global/edgeCacheServices/%s", project, name)

	obj := &pb.EdgeCacheService{}
	if err := s.storage.Get(r.Context(), fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{}"))
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := protojson.Marshal(obj)
	w.Write(b)
}

func (s *NetworkServicesServer) CreateEdgeCacheService(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	project := pathParams["project"]
	parent := fmt.Sprintf("projects/%s/locations/global", project)

	id := r.URL.Query().Get("edgeCacheServiceId")
	if id == "" {
		http.Error(w, "edgeCacheServiceId is required", http.StatusBadRequest)
		return
	}

	fqn := parent + "/edgeCacheServices/" + id

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	obj := &pb.EdgeCacheService{}
	if err := protojson.Unmarshal(body, obj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(r.Context(), fqn, obj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// We need to return an Operation
	opMetadata := &networkservicespb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/global", project)
	lro, err := s.operations.StartLRO(r.Context(), lroPrefix, opMetadata, func() (proto.Message, error) {
		return obj, nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := protojson.Marshal(lro)
	w.Write(b)
}

func (s *NetworkServicesServer) PatchEdgeCacheService(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	project := pathParams["project"]
	name := pathParams["name"]
	fqn := fmt.Sprintf("projects/%s/locations/global/edgeCacheServices/%s", project, name)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	patchObj := &pb.EdgeCacheService{}
	if err := protojson.Unmarshal(body, patchObj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existing := &pb.EdgeCacheService{}
	if err := s.storage.Get(r.Context(), fqn, existing); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Simple patch implementation
	proto.Merge(existing, patchObj)

	now := time.Now()
	existing.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(r.Context(), fqn, existing); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	opMetadata := &networkservicespb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/global", project)
	lro, err := s.operations.StartLRO(r.Context(), lroPrefix, opMetadata, func() (proto.Message, error) {
		return existing, nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := protojson.Marshal(lro)
	w.Write(b)
}

func (s *NetworkServicesServer) DeleteEdgeCacheService(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	project := pathParams["project"]
	name := pathParams["name"]
	fqn := fmt.Sprintf("projects/%s/locations/global/edgeCacheServices/%s", project, name)

	deleted := &pb.EdgeCacheService{}
	if err := s.storage.Delete(r.Context(), fqn, deleted); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now()
	opMetadata := &networkservicespb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/global", project)
	lro, err := s.operations.StartLRO(r.Context(), lroPrefix, opMetadata, func() (proto.Message, error) {
		return &pb.EdgeCacheService{}, nil // Should be Empty but we'll use empty service for now
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	b, _ := protojson.Marshal(lro)
	w.Write(b)
}
