// Copyright 2026 High Five Labs
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

// Objects with content.
// Uploads (POST /upload/storage/v1/b/{bucket}/o, uploadType media or
// multipart) and content reads (alt=media) are not gRPC-gateway routes, so
// they are served here, beside the generated mux, as handleBucketIAM is.
// Each upload makes a new generation, as GCS does. Resumable uploads are not
// served.

package mockstorage

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
)

func objectKey(bucket, name string) string {
	return "buckets/" + bucket + "/objects/" + name
}

func objectMediaKey(bucket, name string) string {
	return objectKey(bucket, name) + "#media"
}

func (s *MockService) getStoredObject(ctx context.Context, bucket, name string) (*pb.Object, error) {
	obj := &pb.Object{}
	if err := s.storage.Get(ctx, objectKey(bucket, name), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No such object: %s/%s", bucket, name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *MockService) deleteStoredObject(ctx context.Context, bucket, name string) error {
	if err := s.storage.Delete(ctx, objectKey(bucket, name), &pb.Object{}); err != nil {
		if status.Code(err) == codes.NotFound {
			return status.Errorf(codes.NotFound, "No such object: %s/%s", bucket, name)
		}
		return err
	}
	_ = s.storage.Delete(ctx, objectMediaKey(bucket, name), &wrapperspb.BytesValue{})
	return nil
}

// handleObjectMedia serves uploads and content reads; it reports whether it
// handled the request.
func (s *MockService) handleObjectMedia(w http.ResponseWriter, r *http.Request) bool {
	path := r.URL.EscapedPath()
	switch {
	case r.Method == http.MethodPost && strings.HasPrefix(path, "/upload/storage/v1/b/") && strings.HasSuffix(path, "/o"):
		bucket := strings.TrimSuffix(strings.TrimPrefix(path, "/upload/storage/v1/b/"), "/o")
		s.upload(w, r, bucket)
		return true
	case r.Method == http.MethodGet && r.URL.Query().Get("alt") == "media" &&
		(strings.HasPrefix(path, "/storage/v1/b/") || strings.HasPrefix(path, "/download/storage/v1/b/")):
		rest := strings.TrimPrefix(strings.TrimPrefix(path, "/download"), "/storage/v1/b/")
		bucket, object, ok := strings.Cut(rest, "/o/")
		if !ok {
			return false
		}
		name, err := url.PathUnescape(object)
		if err != nil {
			writeStorageError(w, http.StatusBadRequest, err.Error())
			return true
		}
		media := &wrapperspb.BytesValue{}
		if err := s.storage.Get(r.Context(), objectMediaKey(bucket, name), media); err != nil {
			writeStorageError(w, http.StatusNotFound, fmt.Sprintf("No such object: %s/%s", bucket, name))
			return true
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write(media.GetValue())
		return true
	}
	return false
}

func writeStorageError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]any{"code": code, "message": message}})
}

func (s *MockService) upload(w http.ResponseWriter, r *http.Request, bucket string) {
	ctx := r.Context()
	if err := s.storage.Get(ctx, "buckets/"+bucket, &pb.Bucket{}); err != nil {
		writeStorageError(w, http.StatusNotFound, "The specified bucket does not exist.")
		return
	}
	obj := &pb.Object{}
	var content []byte
	switch r.URL.Query().Get("uploadType") {
	case "media":
		obj.Name = PtrTo(r.URL.Query().Get("name"))
		obj.ContentType = PtrTo(r.Header.Get("Content-Type"))
		b, err := io.ReadAll(r.Body)
		if err != nil {
			writeStorageError(w, http.StatusBadRequest, err.Error())
			return
		}
		content = b
	case "multipart":
		_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			writeStorageError(w, http.StatusBadRequest, err.Error())
			return
		}
		mr := multipart.NewReader(r.Body, params["boundary"])
		meta, err := mr.NextPart()
		if err != nil {
			writeStorageError(w, http.StatusBadRequest, "missing metadata part: "+err.Error())
			return
		}
		mb, _ := io.ReadAll(meta)
		if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(mb, obj); err != nil {
			writeStorageError(w, http.StatusBadRequest, "bad metadata: "+err.Error())
			return
		}
		data, err := mr.NextPart()
		if err != nil {
			writeStorageError(w, http.StatusBadRequest, "missing media part: "+err.Error())
			return
		}
		if obj.GetContentType() == "" {
			obj.ContentType = PtrTo(data.Header.Get("Content-Type"))
		}
		content, _ = io.ReadAll(data)
	default:
		writeStorageError(w, http.StatusBadRequest, "mockgcp serves uploadType media and multipart only")
		return
	}
	if obj.GetName() == "" {
		writeStorageError(w, http.StatusBadRequest, "object name is required")
		return
	}

	sum := md5.Sum(content)
	crc := make([]byte, 4)
	binary.BigEndian.PutUint32(crc, crc32.Checksum(content, crc32.MakeTable(crc32.Castagnoli)))
	now := time.Now()
	obj.Bucket = PtrTo(bucket)
	obj.Kind = PtrTo("storage#object")
	obj.Size = PtrTo(uint64(len(content)))
	obj.Md5Hash = PtrTo(base64.StdEncoding.EncodeToString(sum[:]))
	obj.Crc32C = PtrTo(base64.StdEncoding.EncodeToString(crc))
	obj.Generation = PtrTo(now.UnixMicro())
	obj.Metageneration = PtrTo(int64(1))
	obj.Id = PtrTo(fmt.Sprintf("%s/%s/%d", bucket, obj.GetName(), obj.GetGeneration()))
	obj.SelfLink = PtrTo("https://www.googleapis.com/storage/v1/b/" + bucket + "/o/" + url.PathEscape(obj.GetName()))
	obj.MediaLink = PtrTo(obj.GetSelfLink() + "?alt=media")
	obj.TimeCreated = timestamppb.New(now)
	obj.Updated = timestamppb.New(now)

	key := objectKey(bucket, obj.GetName())
	if err := s.storage.Update(ctx, key, obj); err != nil {
		if status.Code(err) != codes.NotFound {
			writeStorageError(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err := s.storage.Create(ctx, key, obj); err != nil {
			writeStorageError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	media := &wrapperspb.BytesValue{Value: content}
	mkey := objectMediaKey(bucket, obj.GetName())
	if err := s.storage.Update(ctx, mkey, media); err != nil {
		if err := s.storage.Create(ctx, mkey, media); err != nil {
			writeStorageError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	b, err := protojson.Marshal(obj)
	if err != nil {
		writeStorageError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}
