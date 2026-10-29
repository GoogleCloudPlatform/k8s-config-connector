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

package mockstorage

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"k8s.io/klog/v2"
)

func (s *MockService) handleBucketIAM(w http.ResponseWriter, r *http.Request) bool {
	path := r.URL.Path
	if !strings.HasPrefix(path, "/storage/v1/b/") || !strings.HasSuffix(path, "/iam") {
		return false
	}
	remainder := strings.TrimPrefix(path, "/storage/v1/b/")
	bucket := strings.TrimSuffix(remainder, "/iam")
	if strings.Contains(bucket, "/") {
		return false
	}

	ctx := r.Context()
	switch r.Method {
	case http.MethodGet:
		s.getBucketIAM(ctx, w, r, bucket)
		return true
	case http.MethodPut:
		s.putBucketIAM(ctx, w, r, bucket)
		return true
	default:
		return false
	}
}

func (s *MockService) getBucketIAM(ctx context.Context, w http.ResponseWriter, r *http.Request, bucket string) {
	bucketObj := &pb.Bucket{}
	if err := s.storage.Get(ctx, "buckets/"+bucket, bucketObj); err != nil {
		s.writeStorageError(w, http.StatusNotFound, "Not Found", "notFound")
		return
	}

	policyKey := fmt.Sprintf("projects/_/buckets/%s/iam", bucket)
	policyObj := &pb.Policy{}
	if err := s.storage.Get(ctx, policyKey, policyObj); err != nil {
		projectId := "mock-project"
		project, err := s.Projects.GetProjectByNumber(fmt.Sprintf("%d", bucketObj.GetProjectNumber()))
		if err == nil && project != nil {
			projectId = project.ID
		}

		policyObj = &pb.Policy{
			Kind:       PtrTo("storage#policy"),
			ResourceId: PtrTo(fmt.Sprintf("projects/_/buckets/%s", bucket)),
			Version:    PtrTo(int32(1)),
			Etag:       s.defaultPolicyEtag(),
			Bindings: []*pb.PolicyBindings{
				{
					Role: PtrTo("roles/storage.legacyBucketOwner"),
					Members: []string{
						fmt.Sprintf("projectEditor:%s", projectId),
						fmt.Sprintf("projectOwner:%s", projectId),
					},
				},
				{
					Role: PtrTo("roles/storage.legacyBucketReader"),
					Members: []string{
						fmt.Sprintf("projectViewer:%s", projectId),
					},
				},
			},
		}
	}

	s.writePolicyResponse(w, policyObj)
}

func (s *MockService) putBucketIAM(ctx context.Context, w http.ResponseWriter, r *http.Request, bucket string) {
	bucketObj := &pb.Bucket{}
	if err := s.storage.Get(ctx, "buckets/"+bucket, bucketObj); err != nil {
		s.writeStorageError(w, http.StatusNotFound, "Not Found", "notFound")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.writeStorageError(w, http.StatusBadRequest, "Failed to read body", "badRequest")
		return
	}

	policyObj := &pb.Policy{}
	unmarshalOpts := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err := unmarshalOpts.Unmarshal(body, policyObj); err != nil {
		s.writeStorageError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse policy: %v", err), "badRequest")
		return
	}

	policyObj.Kind = PtrTo("storage#policy")
	policyObj.ResourceId = PtrTo(fmt.Sprintf("projects/_/buckets/%s", bucket))

	hasConditions := false
	for _, b := range policyObj.GetBindings() {
		if b.GetCondition() != nil {
			hasConditions = true
			break
		}
	}
	if !hasConditions {
		policyObj.Version = PtrTo(int32(1))
	}

	if len(policyObj.Etag) == 0 {
		policyObj.Etag = s.defaultPolicyEtag()
	}

	policyKey := fmt.Sprintf("projects/_/buckets/%s/iam", bucket)
	if err := s.storage.Update(ctx, policyKey, policyObj); err != nil {
		if err := s.storage.Create(ctx, policyKey, policyObj); err != nil {
			klog.Errorf("error storing policy %s: %v", policyKey, err)
			s.writeStorageError(w, http.StatusInternalServerError, "Internal server error", "internalError")
			return
		}
	}

	s.writePolicyResponse(w, policyObj)
}

func (s *MockService) defaultPolicyEtag() []byte {
	etag, _ := base64.StdEncoding.DecodeString("abcdef0123A=")
	return etag
}

func (s *MockService) writePolicyResponse(w http.ResponseWriter, policy *pb.Policy) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Server", "UploadServer")
	w.Header().Set("Vary", "Origin")
	w.Header().Add("Vary", "X-Origin")

	marshalOpts := protojson.MarshalOptions{EmitUnpopulated: false}
	data, err := marshalOpts.Marshal(policy)
	if err != nil {
		klog.Errorf("error marshaling policy: %v", err)
		s.writeStorageError(w, http.StatusInternalServerError, "Internal server error", "internalError")
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		klog.Warningf("error writing response: %v", err)
	}
}

func (s *MockService) writeStorageError(w http.ResponseWriter, code int, message string, reason string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Server", "UploadServer")
	w.Header().Set("Vary", "Origin")
	w.Header().Add("Vary", "X-Origin")

	resp := map[string]any{
		"error": map[string]any{
			"code":    code,
			"message": message,
			"errors": []map[string]any{
				{
					"domain":  "global",
					"message": message,
					"reason":  reason,
				},
			},
		},
	}
	data, _ := json.Marshal(resp)
	w.WriteHeader(code)
	_, _ = w.Write(data)
}
