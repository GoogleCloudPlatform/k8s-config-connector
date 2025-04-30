// Copyright 2025 Google LLC
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

package cloudidentity

import (
	"testing"

	"math/rand"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudidentity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func FuzzCloudIdentityGroupSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		// To KRM
		k1 := &krm.CloudIdentityGroupSpec{}
		filler.Fill(t, k1)

		// To API
		ctx := &direct.MapContext{}
		apiObj := CloudIdentityGroupSpec_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API obj: %v \n KRM: %s", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// Back to KRM
		k2 := CloudIdentityGroupSpec_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API obj to KRM: %v \n API: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		opts := cmp.Options{
			cmpopts.IgnoreFields(krm.CloudIdentityGroupSpec{}, "ResourceID"),
			cmpopts.IgnoreFields(krm.CloudIdentityGroupSpec{}, "InitialGroupConfig"), // not an api field
		}

		// Compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}

	})
}

func FuzzCloudIdentityGroupObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		// To KRM
		k1 := &krm.CloudIdentityGroupObservedState{}
		filler.Fill(t, k1)

		// To API
		ctx := &direct.MapContext{}
		apiObj := CloudIdentityGroupObservedState_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting from KRM to API obj: %v \n KRM: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Back to KRM
		k2 := CloudIdentityGroupObservedState_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting from API obj to KRM: %v \n API: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Compare
		if diff := cmp.Diff(k1, k2); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}

	})
}
