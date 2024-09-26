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

package certificatemanager

import (
	"math/rand"
	"testing"

	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/testing/protocmp"
	"k8s.io/apimachinery/pkg/util/sets"
)

func FuzzDNSAuthorizationSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &pb.DnsAuthorization{}
		fuzz.FillWithRandom(t, randStream, p1)

		// We don't expect output fields to round-trip
		outputFields := sets.New(".create_time", ".update_time")

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New[string]()

		// Status fields
		unimplementedFields.Insert(".dns_resource_record", ".name")
		unimplementedFields.Insert(".dns_resource_record", ".type")
		unimplementedFields.Insert(".dns_resource_record", ".data")

		// Labels comes from KRM metadata instead of spec
		unimplementedFields.Insert(".labels")

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
		k := CertificateManagerDNSAuthorizationSpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := CertificateManagerDNSAuthorizationSpec_ToProto(ctx, k)
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
