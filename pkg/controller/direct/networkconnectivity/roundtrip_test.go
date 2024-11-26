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

package networkconnectivity

import (
	"math/rand"
	"testing"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"
)

func FuzzServiceConnectionPolicySpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.ServiceConnectionPolicy{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		outputFields := sets.New(".etag")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".name",
			".labels",
			".psc_config.allowed_google_producers_resource_hierarchy_level",
			".psc_connections[].ip_version",
			".psc_connections[].producer_instance_metadata",
			".psc_connections[].service_class",
		)

		statusFields := sets.New(
			".create_time",
			".update_time",
			".infrastructure",
			".psc_connections",
		)

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(outputFields).Union(statusFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		r := &fuzz.ReplaceFields{}
		r.Func = func(path string, val protoreflect.Value) (protoreflect.Value, bool) {
			// TODO: Any values we need to force
			return protoreflect.Value{}, false
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, r)

		ctx := &direct.MapContext{}
		k := NetworkConnectivityServiceConnectionPolicySpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := NetworkConnectivityServiceConnectionPolicySpec_ToProto(ctx, k)
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

func FuzzServiceConnectionPolicyObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.ServiceConnectionPolicy{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		outputFields := sets.New(".etag")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".name",
			".labels",
			".psc_config.allowed_google_producers_resource_hierarchy_level",
			".psc_connections[].ip_version",
			".psc_connections[].producer_instance_metadata",
			".psc_connections[].service_class",
		)

		specFields := sets.New(
			".description",
			".network",
			".service_class",
			".psc_config",
			".psc_connections[].error.details",
		)

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(outputFields).Union(specFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		r := &fuzz.ReplaceFields{}
		r.Func = func(path string, val protoreflect.Value) (protoreflect.Value, bool) {
			// TODO: Any values we need to force
			return protoreflect.Value{}, false
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, r)

		ctx := &direct.MapContext{}
		k := NetworkConnectivityServiceConnectionPolicyObservedState_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := NetworkConnectivityServiceConnectionPolicyObservedState_ToProto(ctx, k)
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
