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

package sqladmin

import (
	"math/rand"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/sql/v1beta4"
)

func FuzzSQLAdminBackupSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.BackupRun{}
		fuzz.FillWithRandom(t, randStream, p1)

		unimplementedFields := sets.New(
			".kind",
			".self_link",
			".max_chargeable_bytes",
		)

		// Status fields
		statusFields := sets.New(
			".status",
			".enqueued_time",
			".id",
			".start_time",
			".end_time",
			".error",
			".type",
			".window_start_time",
			".disk_encryption_status",
			".time_zone",
		)

		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(statusFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}

		k := SQLAdminBackupSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := SQLAdminBackupSpec_ToProto(ctx, k)
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

func FuzzSQLAdminBackupObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.BackupRun{}
		fuzz.FillWithRandom(t, randStream, p1)

		unimplementedFields := sets.New(
			".kind",
			".self_link",
			".max_chargeable_bytes",
		)

		// Spec fields
		specFields := sets.New(
			".description",
			".instance",
			".location",
			".disk_encryption_configuration",
			".backup_kind",
		)

		clearFields := &fuzz.ClearFields{
			Paths: specFields.Union(unimplementedFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}

		k := SQLAdminBackupObservedState_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := SQLAdminBackupObservedState_ToProto(ctx, k)
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
