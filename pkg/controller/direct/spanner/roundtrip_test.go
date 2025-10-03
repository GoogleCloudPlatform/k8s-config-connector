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

package spanner

import (
	"math/rand"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"

	pb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
)

func FuzzSpannerBackupScheduleSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.BackupSchedule{}
		fuzz.FillWithRandom(t, randStream, p1)

		unimplementedFields := sets.New(
			".name",
		)

		// Status fields
		statusFields := sets.New(
			".update_time",
			".cron_spec",
			".spec.cron_spec.time_zone",
			".spec.cron_spec.creation_window",
		)

		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(statusFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}

		k := SpannerBackupScheduleSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := SpannerBackupScheduleSpec_ToProto(ctx, k)
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

func FuzzSpannerBackupScheduleObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.BackupSchedule{}
		fuzz.FillWithRandom(t, randStream, p1)

		unimplementedFields := sets.New(
			".name",
		)

		// Spec fields
		stpecFields := sets.New(
			".spec",
			".retention_duration",
			".encryption_config",
			".full_backup_spec",
			".incremental_backup_spec",
		)

		clearFields := &fuzz.ClearFields{
			Paths: stpecFields.Union(unimplementedFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}

		k := SpannerBackupScheduleObservedState_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := SpannerBackupScheduleObservedState_ToProto(ctx, k)
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
