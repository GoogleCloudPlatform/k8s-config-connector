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

package cloudbuild

import (
	"testing"
	"time"

	pb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCloudBuildTriggerFuzz(t *testing.T) {
	fuzzer := cloudbuildTriggerFuzzer()
	for i := 0; i < 100; i++ {
		seed := time.Now().UnixNano()
		fuzzer.FuzzSpec(t, seed)
		fuzzer.FuzzStatus(t, seed)
	}
}

func TestCloudBuildTriggerSpec_ToProto(t *testing.T) {
	tests := []struct {
		name string
		in   *krm.CloudBuildTriggerSpec
		want *pb.BuildTrigger
	}{
		{
			name: "basic",
			in: &krm.CloudBuildTriggerSpec{
				Description: direct.LazyPtr("test description"),
				Disabled:    direct.LazyPtr(true),
			},
			want: &pb.BuildTrigger{
				Description: "test description",
				Disabled:    true,
			},
		},
		{
			name: "filename",
			in: &krm.CloudBuildTriggerSpec{
				Filename: direct.LazyPtr("cloudbuild.yaml"),
			},
			want: &pb.BuildTrigger{
				BuildTemplate: &pb.BuildTrigger_Filename{
					Filename: "cloudbuild.yaml",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			got := CloudBuildTriggerSpec_ToProto(mapCtx, tt.in)
			if mapCtx.Err() != nil {
				t.Fatalf("unexpected error: %v", mapCtx.Err())
			}
			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("CloudBuildTriggerSpec_ToProto() diff (-want +got):\n%s", diff)
			}
		})
	}
}
