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
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func TestImagesV1_GetPublicImage(t *testing.T) {
	tests := []struct {
		name     string
		project  string
		image    string
		wantCode codes.Code
		wantName string
	}{
		{
			name:     "family name is not an image",
			project:  "cos-cloud",
			image:    "cos-81-lts",
			wantCode: codes.NotFound,
		},
		{
			name:     "debian family name is not an image",
			project:  "debian-cloud",
			image:    "debian-11",
			wantCode: codes.NotFound,
		},
		{
			name:     "concrete image name",
			project:  "cos-cloud",
			image:    "cos-81-12871-1196-0",
			wantCode: codes.OK,
			wantName: "cos-81-12871-1196-0",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			got, err := s.images.Get(s.ctx, &pb.GetImageRequest{Project: tc.project, Image: tc.image})
			if status.Code(err) != tc.wantCode {
				t.Fatalf("Get(%s, %s) code = %v (err %v), want %v", tc.project, tc.image, status.Code(err), err, tc.wantCode)
			}
			if tc.wantCode == codes.OK && got.GetName() != tc.wantName {
				t.Errorf("Get() name = %q, want %q", got.GetName(), tc.wantName)
			}
		})
	}
}

func TestImagesV1_GetFromFamily_PublicImages(t *testing.T) {
	tests := []struct {
		name            string
		project         string
		family          string
		wantName        string
		wantSelfLink    string
		wantStatus      string
		wantDescription string
		wantDiskSizeGb  int64
	}{
		{
			name:           "cos LTS family",
			project:        "cos-cloud",
			family:         "cos-81-lts",
			wantName:       "cos-81-12871-1196-0",
			wantSelfLink:   "https://www.googleapis.com/compute/v1/projects/cos-cloud/global/images/cos-81-12871-1196-0",
			wantStatus:     "READY",
			wantDiskSizeGb: 10,
		},
		{
			name:           "current cos LTS family",
			project:        "cos-cloud",
			family:         "cos-121-lts",
			wantName:       "cos-121-18867-90-42",
			wantSelfLink:   "https://www.googleapis.com/compute/v1/projects/cos-cloud/global/images/cos-121-18867-90-42",
			wantStatus:     "READY",
			wantDiskSizeGb: 10,
		},
		{
			name:            "debian-11",
			project:         "debian-cloud",
			family:          "debian-11",
			wantName:        "debian-11-bullseye-v20231010",
			wantSelfLink:    "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-11-bullseye-v20231010",
			wantStatus:      "UP",
			wantDescription: "Debian, Debian GNU/Linux, debian-11, amd64",
		},
		{
			name:            "debian-12",
			project:         "debian-cloud",
			family:          "debian-12",
			wantName:        "debian-12-bookworm-v20261023",
			wantSelfLink:    "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-12-bookworm-v20261023",
			wantStatus:      "UP",
			wantDescription: "Debian, Debian GNU/Linux, debian-12, amd64",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			got, err := s.images.GetFromFamily(s.ctx, &pb.GetFromFamilyImageRequest{Project: tc.project, Family: tc.family})
			if err != nil {
				t.Fatalf("GetFromFamily() unexpected error: %v", err)
			}
			if got.GetName() != tc.wantName {
				t.Errorf("name = %q, want %q", got.GetName(), tc.wantName)
			}
			if got.GetSelfLink() != tc.wantSelfLink {
				t.Errorf("selfLink = %q, want %q", got.GetSelfLink(), tc.wantSelfLink)
			}
			if got.GetFamily() != tc.family {
				t.Errorf("family = %q, want %q", got.GetFamily(), tc.family)
			}
			if got.GetStatus() != tc.wantStatus {
				t.Errorf("status = %q, want %q", got.GetStatus(), tc.wantStatus)
			}
			if got.GetDescription() != tc.wantDescription {
				t.Errorf("description = %q, want %q", got.GetDescription(), tc.wantDescription)
			}
			if got.GetDiskSizeGb() != tc.wantDiskSizeGb {
				t.Errorf("diskSizeGb = %d, want %d", got.GetDiskSizeGb(), tc.wantDiskSizeGb)
			}
		})
	}
}

func TestImagesV1_GetFromFamily_NotFound(t *testing.T) {
	tests := []struct {
		name    string
		project string
		family  string
	}{
		{name: "unknown public family", project: "cos-cloud", family: "nope"},
		{name: "unknown family in project without images", project: "test-project", family: "nope"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := newTestComputeServices()
			_, err := s.images.GetFromFamily(s.ctx, &pb.GetFromFamilyImageRequest{Project: tc.project, Family: tc.family})
			if status.Code(err) != codes.NotFound {
				t.Errorf("GetFromFamily(%s, %s) code = %v (err %v), want %v", tc.project, tc.family, status.Code(err), err, codes.NotFound)
			}
		})
	}
}

func TestImagesV1_GetFromFamily_NewestStoredImage(t *testing.T) {
	s := newTestComputeServices()
	insert := func(name string, deprecated bool) {
		t.Helper()
		obj := &pb.Image{
			Name:   PtrTo(name),
			Family: PtrTo("my-family"),
		}
		if deprecated {
			obj.Deprecated = &pb.DeprecationStatus{State: PtrTo("DEPRECATED")}
		}
		if _, err := s.images.Insert(s.ctx, &pb.InsertImageRequest{Project: "test-project", ImageResource: obj}); err != nil {
			t.Fatalf("Insert image %s: %v", name, err)
		}
	}
	insert("my-image-old", false)
	insert("my-image-new", false)
	// A newer image that is deprecated must not be served by the family.
	insert("my-image-deprecated", true)

	got, err := s.images.GetFromFamily(s.ctx, &pb.GetFromFamilyImageRequest{Project: "test-project", Family: "my-family"})
	if err != nil {
		t.Fatalf("GetFromFamily() unexpected error: %v", err)
	}
	if got.GetName() != "my-image-new" {
		t.Errorf("GetFromFamily() name = %q, want %q", got.GetName(), "my-image-new")
	}
}
