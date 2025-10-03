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

// +tool:fuzz-gen
// proto.message: google.cloud.dataproc.v1.Job
// api.group: dataproc.cnrm.cloud.google.com

package dataproc

import (
	"math/rand"
	"testing"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"
)

func FuzzDataprocJobSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Job{}
		fuzz.FillWithRandom(t, randStream, p1)

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".status",
			".status_history",
			".hadoop_job.logging_config",
			".spark_job.logging_config",
			".pyspark_job.logging_config",
			".hive_job.logging_config",
			".pig_job.logging_config",
			".spark_r_job.logging_config",
			".spark_sql_job.logging_config",
			".presto_job.logging_config",
			".trino_job.logging_config",
			".flink_job.logging_config",
		)

		statusField := sets.New(
			".placement",
			".yarn_applications",
			".status",
			".driver_output_resource_uri",
			".driver_control_files_uri",
			".job_uuid",
			".done",
		)

		// Remove any known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(statusField),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}
		k := DataprocJobSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := DataprocJobSpec_ToProto(ctx, k)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from krm to proto: %v", ctx.Err())
		}

		if diff := cmp.Diff(p1, p2, protocmp.Transform()); diff != "" {
			t.Logf("p1 = %v", prototext.Format(p1))
			t.Logf("p2 = %v", prototext.Format(p2))
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}
	})
}
