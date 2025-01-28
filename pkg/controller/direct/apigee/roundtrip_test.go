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

package apigee

import (
	"encoding/json"
	"testing"

	"math/rand"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func FuzzApigeeEnvgroupSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		// To KRM
		k1 := &krm.ApigeeEnvgroupSpec{}
		filler.Fill(t, k1)

		// To API
		ctx := &direct.MapContext{}
		apiObj := ApigeeEnvgroupSpec_ToApi(ctx, k1, *k1.ResourceID)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API obj: %v \n KRM: %s", ctx.Err(), prettyPrint(t, k1))
		}

		// Back to KRM
		k2 := ApigeeEnvgroupSpec_FromApi(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API obj to KRM: %v \n API: %s", ctx.Err(), prettyPrint(t, apiObj))
		}

		// Using cmpopts.IgnoreFields to ignore OrganizationRef
		opts := cmp.Options{
			cmpopts.IgnoreFields(krm.Parent{}, "OrganizationRef"),
		}

		// compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}

	})
}

func FuzzApigeeEnvgroupObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		// To KRM
		k1 := &krm.ApigeeEnvgroupObservedState{}
		filler.Fill(t, k1)

		// To API
		ctx := &direct.MapContext{}
		apiObj := ApigeeEnvgroupObservedState_ToApi(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting from KRM to API obj: %v \n KRM: %s", ctx.Err(), prettyPrint(t, k1))
		}

		// Back to KRM
		k2 := ApigeeEnvgroupObservedState_FromApi(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting from API obj to KRM: %v \n API: %s", ctx.Err(), prettyPrint(t, apiObj))
		}

		// Compare
		if diff := cmp.Diff(k1, k2); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}

	})
}

func prettyPrint(t *testing.T, k any) string {
	encoded, err := json.MarshalIndent(k, "", " ")
	if err != nil {
		t.Fatalf("error encoding to json: %v", err)
	}

	return string(encoded)
}
