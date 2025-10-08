// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package mocklogging

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func TestCreateLink(t *testing.T) {
	s := &MockService{
		storage: storage.NewEmptyStorage(),
	}
	s.Projects.StoreProject(&projects.ProjectData{
		ID:     "my-project",
		Number: 123456,
	})

	ctx := context.Background()
	req := &pb.CreateLinkRequest{
		Parent: "projects/my-project/locations/global/buckets/my-bucket",
		LinkId: "my-link",
		Link: &pb.Link{
			Description: "my link",
		},
	}
	op, err := s.CreateLink(ctx, req)
	if err != nil {
		t.Fatalf("CreateLink failed: %v", err)
	}

	if op.Done {
		t.Errorf("expected LRO to not be done")
	}

	_, err = s.GetLink(ctx, &pb.GetLinkRequest{Name: "projects/my-project/locations/global/buckets/my-bucket/links/my-link"})
	if err != nil {
		t.Fatalf("GetLink failed: %v", err)
	}
}

func TestGetLinkNotFound(t *testing.T) {
	s := &MockService{
		storage: storage.NewEmptyStorage(),
	}
	s.Projects.StoreProject(&projects.ProjectData{
		ID:     "my-project",
		Number: 123456,
	})

	ctx := context.Background()
	_, err := s.GetLink(ctx, &pb.GetLinkRequest{Name: "projects/my-project/locations/global/buckets/my-bucket/links/my-link"})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("GetLink request should have failed with NotFound, but got %v", err)
	}
}
