// Copyright 2024 Google LLC
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

package bigquerydatatransfer

import (
	"math/rand"
	"testing"

	pb "cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
)

func FuzzBigQueryDataTransferConfigSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.TransferConfig{}
		fuzz.FillWithRandom(t, randStream, p1)

		// Status fields
		unimplementedFields := sets.New(
			".name",
			".dataset_region",
			".next_run_time",
			".owner_info",
			".state",
			".update_time",
			".user_id",
		)

		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields,
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}
		k := BigQueryDataTransferConfigSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := BigQueryDataTransferConfigSpec_ToProto(ctx, k)
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
