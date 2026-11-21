// Copyright 2023 Google LLC
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
	"net/http"
	"net/url"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/httpmux"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/storage/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type objects struct {
	*MockService
	pb.UnimplementedObjectsServerServer
}

func (s *objects) ListObjects(ctx context.Context, req *pb.ListObjectsRequest) (*pb.Objects, error) {
	// A stub implementation, just to support deletion (for now)

	httpmux.SetExpiresHeader(ctx, time.Now())

	ret := &pb.Objects{}
	ret.Kind = PtrTo("storage#objects")
	ret.Prefixes = append(ret.Prefixes, "testfolder")
	ret.Prefixes = append(ret.Prefixes, "testmanagedfolder")
	return ret, nil
}

func (s *objects) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.Object, error) {
	// Objects are uploaded through handleObjectMedia (objectmedia.go).
	name, err := url.PathUnescape(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "object name %q: %v", req.GetName(), err)
	}
	return s.getStoredObject(ctx, req.GetBucket(), name)
}

// DeleteObject deletes an object uploaded through handleObjectMedia.
func (s *objects) DeleteObject(ctx context.Context, req *pb.DeleteObjectRequest) (*empty.Empty, error) {
	name, err := url.PathUnescape(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "object name %q: %v", req.GetName(), err)
	}
	if err := s.deleteStoredObject(ctx, req.GetBucket(), name); err != nil {
		return nil, err
	}
	httpmux.SetStatusCode(ctx, http.StatusNoContent)
	return &empty.Empty{}, nil
}
