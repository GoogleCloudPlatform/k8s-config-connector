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

package dataform

import (
	"math/rand"
	"testing"

	dataformpb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
)

func FuzzDataformRepositorySpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		randStream := rand.New(rand.NewSource(seed))

		p1 := &dataformpb.Repository{}
		fuzz.FillWithRandom(t, randStream, p1)

		// A few fields are not implemented yet in KRM, don't test them
		unimplementedFields := sets.New(
			".name",
			".labels",
			".git_remote_settings.token_status", // not supported by KCC anymore
		)

		// Fields added more recently that can potentially be implemented
		unimplementedFields.Insert(".create_time")
		unimplementedFields.Insert(".data_encryption_state")
		unimplementedFields.Insert(".kms_key_name")
		unimplementedFields.Insert(".internal_metadata")

		// Remove any known-unimplemented fields
		clearFields := &fuzz.ClearFields{
			Paths: unimplementedFields,
		}
		fuzz.Visit("", p1.ProtoReflect(), nil, clearFields)

		ctx := &direct.MapContext{}
		k := DataformRepositorySpec_FromProto(ctx, p1)
		if ctx.Err() != nil {
			t.Fatalf("error mapping from proto to krm: %v", ctx.Err())
		}

		p2 := DataformRepositorySpec_ToProto(ctx, k)
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
