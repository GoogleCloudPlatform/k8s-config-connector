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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func TestApplyImagePatch(t *testing.T) {
	stored := &pb.Image{
		Name:        PtrTo("img"),
		Description: PtrTo("old description"),
		SourceDisk:  PtrTo("https://www.googleapis.com/compute/v1/projects/p/zones/us-central1-a/disks/d"),
		Labels:      map[string]string{"managed-by-cnrm": "true"},
	}
	tests := []struct {
		name     string
		patch    *pb.Image
		wantCode codes.Code
		wantMsg  string
		want     *pb.Image
	}{
		{
			name:  "a changed description and family are applied",
			patch: &pb.Image{Description: PtrTo("new description"), Family: PtrTo("f")},
			want: &pb.Image{
				Name:        PtrTo("img"),
				Description: PtrTo("new description"),
				Family:      PtrTo("f"),
				SourceDisk:  PtrTo("https://www.googleapis.com/compute/v1/projects/p/zones/us-central1-a/disks/d"),
				Labels:      map[string]string{"managed-by-cnrm": "true"},
			},
		},
		{
			name:  "immutable fields sent with their stored value are accepted",
			patch: &pb.Image{Name: PtrTo("img"), SourceDisk: PtrTo("https://www.googleapis.com/compute/v1/projects/p/zones/us-central1-a/disks/d")},
			want:  stored,
		},
		{
			name:     "a changed immutable field is rejected",
			patch:    &pb.Image{SourceDisk: PtrTo("projects/p/zones/us-central1-a/disks/d")},
			wantCode: codes.InvalidArgument,
			wantMsg:  "Invalid value for field 'resource.sourceDisk': 'projects/p/zones/us-central1-a/disks/d'. Updating 'sourceDisk' is not supported.",
		},
		{
			name:     "unchanged labels are rejected",
			patch:    &pb.Image{Labels: map[string]string{"managed-by-cnrm": "true"}},
			wantCode: codes.InvalidArgument,
			wantMsg:  "Invalid value for field 'resource.labels': ''. Updating 'labels' is not supported.",
		},
		{
			name:     "a changed immutable field is reported before labels",
			patch:    &pb.Image{SourceDisk: PtrTo("projects/p/zones/us-central1-a/disks/d"), Labels: map[string]string{"managed-by-cnrm": "true"}},
			wantCode: codes.InvalidArgument,
			wantMsg:  "Invalid value for field 'resource.sourceDisk': 'projects/p/zones/us-central1-a/disks/d'. Updating 'sourceDisk' is not supported.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			storedBefore := proto.CloneOf(stored)

			got, err := applyImagePatch(stored, tc.patch)

			if gotCode := status.Code(err); gotCode != tc.wantCode {
				t.Errorf("applyImagePatch() code = %v, want %v", gotCode, tc.wantCode)
			}
			if gotMsg := status.Convert(err).Message(); gotMsg != tc.wantMsg {
				t.Errorf("applyImagePatch() message = %q, want %q", gotMsg, tc.wantMsg)
			}
			if !proto.Equal(got, tc.want) {
				t.Errorf("applyImagePatch() = %v, want %v", got, tc.want)
			}
			if !proto.Equal(stored, storedBefore) {
				t.Errorf("applyImagePatch() modified stored: got %v, want %v", stored, storedBefore)
			}
		})
	}
}
