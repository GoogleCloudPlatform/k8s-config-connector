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
	"reflect"
	"testing"
	"time"

	"math/rand"

	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/fuzz"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func FuzzApigeeEndpointAttachmentSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		// To KRM
		k1 := &krmv1alpha1.ApigeeEndpointAttachmentSpec{}
		filler.Fill(t, k1)

		// To API
		ctx := &direct.MapContext{}
		apiObj := ApigeeEndpointAttachmentSpec_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API obj: %v \n KRM: %s", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// Back to KRM
		k2 := ApigeeEndpointAttachmentSpec_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API obj to KRM: %v \n API: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		opts := cmp.Options{
			cmpopts.IgnoreFields(krmv1alpha1.ApigeeEndpointAttachmentSpec{}, "OrganizationRef"),
			cmpopts.IgnoreFields(krmv1alpha1.ApigeeEndpointAttachmentSpec{}, "ResourceID"),
			cmpopts.IgnoreFields(refs.ComputeServiceAttachmentRef{}, "Name"),
			cmpopts.IgnoreFields(refs.ComputeServiceAttachmentRef{}, "Namespace"),
		}

		// Compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}

	})
}

func FuzzApigeeEndpointAttachmentObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		overrides := map[string]fuzz.OverrideFiller{
			".CreatedAt": func(t *testing.T, fieldName string, field reflect.Value) {
				// Generate a valid timestamp within 10 years.
				validTime := time.Now().Add(time.Duration(stream.Intn(365*10)) * 24 * time.Hour)
				field.SetString(validTime.Format(time.RFC3339))
			},
		}

		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream, FieldOverrides: overrides})

		k1 := &krmv1alpha1.ApigeeEndpointAttachmentObservedState{}
		filler.Fill(t, k1)

		// KRM -> API
		ctx := &direct.MapContext{}
		apiObj := ApigeeEndpointAttachmentObservedState_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API: %v, krm = %v", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// API -> KRM
		k2 := ApigeeEndpointAttachmentObservedState_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API to KRM: %v, api = %v", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Compare
		if diff := cmp.Diff(k1, k2); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed: diff = %s", diff)
		}
	})
}

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
			t.Fatalf("error converting KRM to API obj: %v \n KRM: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Back to KRM
		k2 := ApigeeEnvgroupSpec_FromApi(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API obj to KRM: %v \n API: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
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
			t.Fatalf("error converting from KRM to API obj: %v \n KRM: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Back to KRM
		k2 := ApigeeEnvgroupObservedState_FromApi(ctx, apiObj)
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

func FuzzApigeeEnvgroupAttachmentSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		// To KRM
		k1 := &krmv1alpha1.ApigeeEnvgroupAttachmentSpec{}
		filler.Fill(t, k1)

		// To API
		ctx := &direct.MapContext{}
		apiObj := ApigeeEnvgroupAttachmentSpec_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API obj: %v \n KRM: %s", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// Back to KRM
		k2 := ApigeeEnvgroupAttachmentSpec_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API obj to KRM: %v \n API: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		opts := cmp.Options{
			cmpopts.IgnoreFields(krmv1alpha1.ApigeeEnvgroupAttachmentSpec{}, "EnvgroupRef"),
			cmpopts.IgnoreFields(krmv1alpha1.ApigeeEnvgroupAttachmentSpec{}, "ResourceID"),
			cmpopts.IgnoreFields(krm.ApigeeEnvironmentRef{}, "Name"),
			cmpopts.IgnoreFields(krm.ApigeeEnvironmentRef{}, "Namespace"),
		}

		// Compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}

	})
}

func FuzzApigeeEnvgroupAttachmentObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		overrides := map[string]fuzz.OverrideFiller{
			".CreatedAt": func(t *testing.T, fieldName string, field reflect.Value) {
				// Generate a valid timestamp within 10 years.
				validTime := time.Now().Add(time.Duration(stream.Intn(365*10)) * 24 * time.Hour)
				field.SetString(validTime.Format(time.RFC3339))
			},
		}

		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream, FieldOverrides: overrides})

		k1 := &krmv1alpha1.ApigeeEnvgroupAttachmentObservedState{}
		filler.Fill(t, k1)

		// KRM -> API
		ctx := &direct.MapContext{}
		apiObj := ApigeeEnvgroupAttachmentObservedState_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API: %v, krm = %v", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// API -> KRM
		k2 := ApigeeEnvgroupAttachmentObservedState_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API to KRM: %v, api = %v", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Compare
		if diff := cmp.Diff(k1, k2); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed: diff = %s", diff)
		}
	})
}

func FuzzApigeeInstanceSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		k1 := &krm.ApigeeInstanceSpec{}
		filler.Fill(t, k1)

		// KRM -> API
		ctx := &direct.MapContext{}
		apiObj := ApigeeInstanceSpec_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API: %v, krm = %v", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// API -> KRM
		k2 := ApigeeInstanceSpec_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API to KRM: %v, api = %v", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Ignore Parent, ResourceID, and ref Name+Namespace fields during comparison
		opts := cmp.Options{
			cmpopts.IgnoreFields(krm.ApigeeInstanceSpec{}, "OrganizationRef"),
			cmpopts.IgnoreFields(krm.ApigeeInstanceSpec{}, "ResourceID"),
			cmpopts.IgnoreFields(refs.KMSCryptoKeyRef{}, "Name"),
			cmpopts.IgnoreFields(refs.KMSCryptoKeyRef{}, "Namespace"),
		}
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed: diff = %s", diff)
		}
	})
}

func FuzzApigeeInstanceObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		k1 := &krm.ApigeeInstanceObservedState{}
		filler.Fill(t, k1)

		// KRM -> API
		ctx := &direct.MapContext{}
		apiObj := ApigeeInstanceObservedState_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API: %v, krm = %v", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// API -> KRM
		k2 := ApigeeInstanceObservedState_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API to KRM: %v, api = %v", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Compare
		if diff := cmp.Diff(k1, k2); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed: diff = %s", diff)
		}
	})
}

func FuzzApigeeInstanceAttachmentSpec(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream})

		// To KRM
		k1 := &krmv1alpha1.ApigeeInstanceAttachmentSpec{}
		filler.Fill(t, k1)

		// To API
		ctx := &direct.MapContext{}
		apiObj := ApigeeInstanceAttachmentSpec_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API obj: %v \n KRM: %s", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// Back to KRM
		k2 := ApigeeInstanceAttachmentSpec_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API obj to KRM: %v \n API: %s", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		opts := cmp.Options{
			cmpopts.IgnoreFields(krmv1alpha1.ApigeeInstanceAttachmentSpec{}, "InstanceRef"),
			cmpopts.IgnoreFields(krmv1alpha1.ApigeeInstanceAttachmentSpec{}, "ResourceID"),
			cmpopts.IgnoreFields(krm.ApigeeEnvironmentRef{}, "Name"),
			cmpopts.IgnoreFields(krm.ApigeeEnvironmentRef{}, "Namespace"),
		}

		// Compare
		if diff := cmp.Diff(k1, k2, opts...); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed; diff:\n%s", diff)
		}

	})
}

func FuzzApigeeInstanceAttachmentObservedState(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		stream := rand.New(rand.NewSource(seed))
		overrides := map[string]fuzz.OverrideFiller{
			".CreatedAt": func(t *testing.T, fieldName string, field reflect.Value) {
				// Generate a valid timestamp within 10 years.
				validTime := time.Now().Add(time.Duration(stream.Intn(365*10)) * 24 * time.Hour)
				field.SetString(validTime.Format(time.RFC3339))
			},
		}

		filler := fuzz.NewRandomFiller(&fuzz.FillerConfig{Stream: stream, FieldOverrides: overrides})

		k1 := &krmv1alpha1.ApigeeInstanceAttachmentObservedState{}
		filler.Fill(t, k1)

		// KRM -> API
		ctx := &direct.MapContext{}
		apiObj := ApigeeInstanceAttachmentObservedState_ToAPI(ctx, k1)
		if ctx.Err() != nil {
			t.Fatalf("error converting KRM to API: %v, krm = %v", ctx.Err(), test.PrettyPrintJSON(t, k1))
		}

		// API -> KRM
		k2 := ApigeeInstanceAttachmentObservedState_FromAPI(ctx, apiObj)
		if ctx.Err() != nil {
			t.Fatalf("error converting API to KRM: %v, api = %v", ctx.Err(), test.PrettyPrintJSON(t, apiObj))
		}

		// Compare
		if diff := cmp.Diff(k1, k2); diff != "" {
			t.Logf("k1 = %v", k1)
			t.Logf("k2 = %v", k2)
			t.Errorf("roundtrip failed: diff = %s", diff)
		}
	})
}
