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

package securesourcemanager

import (
	"math/rand"
	"testing"

	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"
)

func FuzzSecureSourceManagerInstanceSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Instance{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		ignoreFields := sets.New(
			// Handled specially
			".name",

			// ObservedState (output) fields
			".create_time",
			".update_time",
			".host_config",
			".private_config.http_service_attachment",
			".private_config.ssh_service_attachment",
			".state",
			".state_note",
		)

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New[string]()

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(ignoreFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}
		k := SecureSourceManagerInstanceSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := SecureSourceManagerInstanceSpec_ToProto(ctx, k)
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

func FuzzSecureSourceManagerInstanceObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Instance{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect spec fields to round-trip
		ignoreFields := sets.New(
			// Handled specially
			".name",

			// Spec fields
			".kms_key",
			".private_config.ca_pool",
			".private_config.is_private",
			".labels",
		)

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: ignoreFields,
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}
		k := SecureSourceManagerInstanceObservedState_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := SecureSourceManagerInstanceObservedState_ToProto(ctx, k)
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

func FuzzSecureSourceManagerRepositorySpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Repository{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		outputFields := sets.New[string]()
		outputFields.Insert(".uid")
		outputFields.Insert(".uris")
		outputFields.Insert(".etag")
		outputFields.Insert(".create_time")
		outputFields.Insert(".update_time")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New[string]()
		unimplementedFields.Insert(".name")

		// Temporarily not supported due to lack of Update API.
		unimplementedFields.Insert(".description")

		// Remove any output only or known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(outputFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		r := &fuzz.ReplaceFields{}
		r.Func = func(path string, val protoreflect.Value) (protoreflect.Value, bool) {
			return protoreflect.Value{}, false
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, r)

		ctx := &direct.MapContext{}
		k := SecureSourceManagerRepositorySpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := SecureSourceManagerRepositorySpec_ToProto(ctx, k)
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

func FuzzSecureSourceManagerRepositoryObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.Repository{}
		fuzz.FillWithRandom(t, randStream, p1)

		// Ignore spec fields
		specFields := sets.New[string]()
		specFields.Insert(".instance")
		specFields.Insert(".initial_config")
		// Temporarily not supported due to lack of Update API.
		specFields.Insert(".description")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New[string]()
		unimplementedFields.Insert(".name")

		// Remove any output only or spec fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields.Union(specFields),
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		r := &fuzz.ReplaceFields{}
		r.Func = func(path string, val protoreflect.Value) (protoreflect.Value, bool) {
			return protoreflect.Value{}, false
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, r)

		ctx := &direct.MapContext{}
		k := SecureSourceManagerRepositoryObservedState_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := SecureSourceManagerRepositoryObservedState_ToProto(ctx, k)
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
