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
	"time"

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
	// A stub implementation, just to support deletion (for now)

	httpmux.SetExpiresHeader(ctx, time.Now())

	return nil, status.Errorf(codes.NotFound, "object not found")
}
