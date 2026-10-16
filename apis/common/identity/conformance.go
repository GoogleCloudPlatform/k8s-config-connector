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

package identity

import (
	"reflect"
	"strings"
	"testing"
)

// ConformanceRef is the subset of refs.ExternalRef exercised by AssertConformance.
// Defining it here avoids importing apis/refs/v1beta1 from apis/common/identity.
type ConformanceRef interface {
	SetExternal(ref string)
	ValidateExternal(ref string) error
	ParseExternalToIdentity() (Identity, error)
}

// ConformanceCase configures shared conformance checks for a resource's
// IdentityV2 and Ref implementations.
type ConformanceCase[I any] struct {
	// SampleExternal is a valid canonical external reference (without //host/ prefix),
	// e.g. "projects/my-project/locations/us-central1/widgets/my-widget".
	SampleExternal string

	// WantHost is the expected service host returned by Host(),
	// e.g. "widget.googleapis.com".
	WantHost string

	// WantParent is the expected ParentString() when the identity implements
	// ParentString() string. Leave empty if the resource has no parent prefix.
	WantParent string

	// WantIdentity is the expected struct value after FromExternal(SampleExternal).
	WantIdentity *I

	// ExpectServerGenerated asserts that the identity implements ServerGeneratedIdentity.
	ExpectServerGenerated bool

	// Ref is an optional zero-value pointer to the resource's <Kind>Ref struct
	// (e.g. &WidgetRef{}) to test ValidateExternal and ParseExternalToIdentity.
	Ref ConformanceRef
}

// AssertConformance runs the standard conformance suite against an IdentityV2
// (and optional Ref) implementation.
func AssertConformance[I any](t *testing.T, tc ConformanceCase[I]) {
	t.Helper()

	if tc.SampleExternal == "" {
		t.Fatal("ConformanceCase.SampleExternal must be set")
	}

	t.Run("valid_external_ref", func(t *testing.T) {
		assertValidExternalRef(t, tc)
	})
	if tc.WantHost != "" {
		t.Run("valid_external_ref_with_host_prefix", func(t *testing.T) {
			assertHostPrefixExternalRef(t, tc)
		})
	}
	assertInvalidExternalRefs(t, tc)
	if tc.ExpectServerGenerated {
		t.Run("server_generated_identity", func(t *testing.T) {
			assertServerGeneratedIdentity(t, tc)
		})
	}
	if tc.Ref != nil {
		assertReferenceConformance(t, tc)
	}
}

type parentStringer interface {
	ParentString() string
}

func newConformanceIdentity[I any](t *testing.T) (IdentityV2, *I) {
	t.Helper()
	raw := new(I)
	id, ok := any(raw).(IdentityV2)
	if !ok {
		t.Fatalf("%T does not implement identity.IdentityV2", raw)
	}
	return id, raw
}

func assertValidExternalRef[I any](t *testing.T, tc ConformanceCase[I]) {
	t.Helper()
	id, raw := newConformanceIdentity[I](t)
	if err := id.FromExternal(tc.SampleExternal); err != nil {
		t.Fatalf("FromExternal(%q) unexpected error: %v", tc.SampleExternal, err)
	}
	if got := id.String(); got != tc.SampleExternal {
		t.Errorf("String() = %q, want %q", got, tc.SampleExternal)
	}
	if tc.WantHost != "" {
		if got := id.Host(); got != tc.WantHost {
			t.Errorf("Host() = %q, want %q", got, tc.WantHost)
		}
	}
	if tc.WantIdentity != nil && !reflect.DeepEqual(raw, tc.WantIdentity) {
		t.Errorf("parsed identity = %+v, want %+v", raw, tc.WantIdentity)
	}
	if ps, ok := any(raw).(parentStringer); ok {
		if got := ps.ParentString(); got != tc.WantParent {
			t.Errorf("ParentString() = %q, want %q", got, tc.WantParent)
		}
	} else if tc.WantParent != "" {
		t.Errorf("%T does not implement ParentString() string, but WantParent=%q was set", raw, tc.WantParent)
	}
}

func assertHostPrefixExternalRef[I any](t *testing.T, tc ConformanceCase[I]) {
	t.Helper()
	id, raw := newConformanceIdentity[I](t)
	fullURI := "//" + tc.WantHost + "/" + tc.SampleExternal
	if err := id.FromExternal(fullURI); err != nil {
		t.Fatalf("FromExternal(%q) unexpected error: %v", fullURI, err)
	}
	if got := id.String(); got != tc.SampleExternal {
		t.Errorf("String() after full URI parse = %q, want %q", got, tc.SampleExternal)
	}
	if tc.WantIdentity != nil && !reflect.DeepEqual(raw, tc.WantIdentity) {
		t.Errorf("parsed identity from full URI = %+v, want %+v", raw, tc.WantIdentity)
	}
}

func assertInvalidExternalRefs[I any](t *testing.T, tc ConformanceCase[I]) {
	t.Helper()
	type invalidCase struct {
		name string
		ref  string
	}
	invalids := []invalidCase{
		{name: "empty", ref: ""},
		{name: "malformed", ref: "invalid"},
		{name: "extra_segment", ref: tc.SampleExternal + "/extra/segment"},
	}
	if idx := strings.IndexByte(tc.SampleExternal, '/'); idx > 0 {
		invalids = append(invalids, invalidCase{
			name: "wrong_collection",
			ref:  "wrongCollection" + tc.SampleExternal[idx:],
		})
	}
	if tc.WantHost != "" && tc.WantHost != "wrong.googleapis.com" {
		invalids = append(invalids, invalidCase{
			name: "wrong_host",
			ref:  "//wrong.googleapis.com/" + tc.SampleExternal,
		})
	}
	for _, inv := range invalids {
		inv := inv
		t.Run("invalid_"+inv.name, func(t *testing.T) {
			id, _ := newConformanceIdentity[I](t)
			if err := id.FromExternal(inv.ref); err == nil {
				t.Errorf("FromExternal(%q) succeeded, expected error", inv.ref)
			}
		})
	}
}

func assertServerGeneratedIdentity[I any](t *testing.T, tc ConformanceCase[I]) {
	t.Helper()
	id, raw := newConformanceIdentity[I](t)
	sg, ok := any(raw).(ServerGeneratedIdentity)
	if !ok {
		t.Fatalf("%T does not implement identity.ServerGeneratedIdentity", raw)
	}
	if sg.HasIdentitySpecified() {
		t.Errorf("zero-value %T reported HasIdentitySpecified() = true, want false", raw)
	}
	if err := id.FromExternal(tc.SampleExternal); err != nil {
		t.Fatalf("FromExternal(%q) unexpected error: %v", tc.SampleExternal, err)
	}
	if !sg.HasIdentitySpecified() {
		t.Errorf("populated %T reported HasIdentitySpecified() = false, want true", raw)
	}
}

func assertReferenceConformance[I any](t *testing.T, tc ConformanceCase[I]) {
	t.Helper()
	t.Run("reference_ValidateExternal", func(t *testing.T) {
		if err := tc.Ref.ValidateExternal(tc.SampleExternal); err != nil {
			t.Errorf("ValidateExternal(%q) unexpected error: %v", tc.SampleExternal, err)
		}
		if err := tc.Ref.ValidateExternal("invalid/external/ref"); err == nil {
			t.Errorf("ValidateExternal(invalid) succeeded, expected error")
		}
	})
	t.Run("reference_ParseExternalToIdentity", func(t *testing.T) {
		tc.Ref.SetExternal(tc.SampleExternal)
		parsed, err := tc.Ref.ParseExternalToIdentity()
		if err != nil {
			t.Fatalf("ParseExternalToIdentity(%q) unexpected error: %v", tc.SampleExternal, err)
		}
		if tc.WantIdentity != nil && !reflect.DeepEqual(parsed, tc.WantIdentity) {
			t.Errorf("ParseExternalToIdentity(%q) = %+v, want %+v", tc.SampleExternal, parsed, tc.WantIdentity)
		}
		tc.Ref.SetExternal("invalid/external/ref")
		if _, err := tc.Ref.ParseExternalToIdentity(); err == nil {
			t.Errorf("ParseExternalToIdentity(invalid) succeeded, expected error")
		}
		tc.Ref.SetExternal("")
	})
}
