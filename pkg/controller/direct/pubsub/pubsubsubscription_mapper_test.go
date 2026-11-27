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

package pubsub

import (
	"testing"

	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	krmbigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestBigQueryConfig_FromProto(t *testing.T) {
	tests := []struct {
		name string
		in   *pb.BigQueryConfig
		want *krm.BigQueryConfig
	}{
		{
			name: "table in the Pub/Sub format becomes a canonical table reference",
			in:   &pb.BigQueryConfig{Table: "my-project.my_dataset.my_table", DropUnknownFields: true},
			want: &krm.BigQueryConfig{
				TableRef:          &krmbigqueryv1beta1.BigQueryTableRef{External: "projects/my-project/datasets/my_dataset/tables/my_table"},
				DropUnknownFields: direct.PtrTo(true),
			},
		},
		{
			name: "table with a colon after the project becomes a canonical table reference",
			in:   &pb.BigQueryConfig{Table: "my-project:my_dataset.my_table"},
			want: &krm.BigQueryConfig{
				TableRef: &krmbigqueryv1beta1.BigQueryTableRef{External: "projects/my-project/datasets/my_dataset/tables/my_table"},
			},
		},
		{
			name: "domain-scoped project is kept intact",
			in:   &pb.BigQueryConfig{Table: "example.com:my-project.my_dataset.my_table"},
			want: &krm.BigQueryConfig{
				TableRef: &krmbigqueryv1beta1.BigQueryTableRef{External: "projects/example.com:my-project/datasets/my_dataset/tables/my_table"},
			},
		},
		{
			name: "value that can't be parsed is kept unchanged",
			in:   &pb.BigQueryConfig{Table: "not-a-table"},
			want: &krm.BigQueryConfig{
				TableRef: &krmbigqueryv1beta1.BigQueryTableRef{External: "not-a-table"},
			},
		},
		{
			name: "boolean options are mapped",
			in:   &pb.BigQueryConfig{Table: "my-project.my_dataset.my_table", UseTopicSchema: true, WriteMetadata: true},
			want: &krm.BigQueryConfig{
				TableRef:       &krmbigqueryv1beta1.BigQueryTableRef{External: "projects/my-project/datasets/my_dataset/tables/my_table"},
				UseTopicSchema: direct.PtrTo(true),
				WriteMetadata:  direct.PtrTo(true),
			},
		},
		{
			name: "empty table leaves the table reference unset",
			in:   &pb.BigQueryConfig{},
			want: &krm.BigQueryConfig{},
		},
		{
			name: "nil config",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}

			got := BigQueryConfig_FromProto(mapCtx, tc.in)

			if err := mapCtx.Err(); err != nil {
				t.Fatalf("BigQueryConfig_FromProto() error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("BigQueryConfig_FromProto() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestBigQueryConfig_ToProto(t *testing.T) {
	tests := []struct {
		name string
		in   *krm.BigQueryConfig
		want *pb.BigQueryConfig
	}{
		{
			name: "canonical table reference is sent in the Pub/Sub format",
			in: &krm.BigQueryConfig{
				TableRef:          &krmbigqueryv1beta1.BigQueryTableRef{External: "projects/my-project/datasets/my_dataset/tables/my_table"},
				DropUnknownFields: direct.PtrTo(true),
			},
			want: &pb.BigQueryConfig{Table: "my-project.my_dataset.my_table", DropUnknownFields: true},
		},
		{
			name: "table reference with the bigquery.googleapis.com host is sent in the Pub/Sub format",
			in: &krm.BigQueryConfig{
				TableRef: &krmbigqueryv1beta1.BigQueryTableRef{External: "//bigquery.googleapis.com/projects/my-project/datasets/my_dataset/tables/my_table"},
			},
			want: &pb.BigQueryConfig{Table: "my-project.my_dataset.my_table"},
		},
		{
			name: "domain-scoped project is kept intact",
			in: &krm.BigQueryConfig{
				TableRef: &krmbigqueryv1beta1.BigQueryTableRef{External: "projects/example.com:my-project/datasets/my_dataset/tables/my_table"},
			},
			want: &pb.BigQueryConfig{Table: "example.com:my-project.my_dataset.my_table"},
		},
		{
			name: "value that is not a canonical table reference is kept unchanged",
			in: &krm.BigQueryConfig{
				TableRef: &krmbigqueryv1beta1.BigQueryTableRef{External: "not-a-table"},
			},
			want: &pb.BigQueryConfig{Table: "not-a-table"},
		},
		{
			name: "boolean options are mapped",
			in: &krm.BigQueryConfig{
				TableRef:       &krmbigqueryv1beta1.BigQueryTableRef{External: "projects/my-project/datasets/my_dataset/tables/my_table"},
				UseTopicSchema: direct.PtrTo(true),
				WriteMetadata:  direct.PtrTo(true),
			},
			want: &pb.BigQueryConfig{Table: "my-project.my_dataset.my_table", UseTopicSchema: true, WriteMetadata: true},
		},
		{
			name: "missing table reference leaves the table empty",
			in:   &krm.BigQueryConfig{},
			want: &pb.BigQueryConfig{},
		},
		{
			name: "nil config",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}

			got := BigQueryConfig_ToProto(mapCtx, tc.in)

			if err := mapCtx.Err(); err != nil {
				t.Fatalf("BigQueryConfig_ToProto() error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("BigQueryConfig_ToProto() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
