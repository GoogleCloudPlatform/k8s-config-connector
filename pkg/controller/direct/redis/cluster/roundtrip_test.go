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

package cluster

import (
	"math/rand"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"

	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
)

func FuzzRedisClusterSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Cluster{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		outputFields := sets.New(".etag")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".name",
			".labels",
		)

		// Status fields
		unimplementedFields.Insert(".discovery_endpoints")
		unimplementedFields.Insert(".uid")
		unimplementedFields.Insert(".precise_size_gb")
		unimplementedFields.Insert(".size_gb")
		unimplementedFields.Insert(".state_info")
		unimplementedFields.Insert(".create_time")
		unimplementedFields.Insert(".state")
		unimplementedFields.Insert(".psc_connections")

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(outputFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		r := &fuzz.ReplaceFields{}
		r.Func = func(path string, val protoreflect.Value) (protoreflect.Value, bool) {
			// TODO: Any values that must follow a pattern
			return protoreflect.Value{}, false
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, r)

		ctx := &direct.MapContext{}
		k := RedisClusterSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := RedisClusterSpec_ToProto(ctx, k)
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

func FuzzRedisClusterObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Cluster{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		outputFields := sets.New(".etag")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".name",
			".labels",
		)

		// Spec fields
		unimplementedFields.Insert(".persistence_config")
		unimplementedFields.Insert(".psc_configs")
		unimplementedFields.Insert(".zone_distribution_config")
		unimplementedFields.Insert(".redis_configs")
		unimplementedFields.Insert(".shard_count")
		unimplementedFields.Insert(".transit_encryption_mode")
		unimplementedFields.Insert(".node_type")
		unimplementedFields.Insert(".authorization_mode")
		unimplementedFields.Insert(".replica_count")
		unimplementedFields.Insert(".deletion_protection_enabled")

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(outputFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		r := &fuzz.ReplaceFields{}
		r.Func = func(path string, val protoreflect.Value) (protoreflect.Value, bool) {
			// TODO: Any values that must follow a pattern
			return protoreflect.Value{}, false
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, r)

		ctx := &direct.MapContext{}
		k := RedisClusterObservedState_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := RedisClusterObservedState_ToProto(ctx, k)
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
