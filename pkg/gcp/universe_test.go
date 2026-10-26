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

package gcp

import (
	"testing"
)

func TestResolveUniverseDomain(t *testing.T) {
	t.Run("defaults to googleapis.com", func(t *testing.T) {
		t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "")
		t.Setenv("UNIVERSE_DOMAIN", "")
		if got := ResolveUniverseDomain(""); got != DefaultUniverseDomain {
			t.Fatalf("ResolveUniverseDomain(\"\") = %q, want %q", got, DefaultUniverseDomain)
		}
	})

	t.Run("prefers explicit argument over env", func(t *testing.T) {
		t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "env.universe.goog")
		if got := ResolveUniverseDomain("explicit.universe.goog"); got != "explicit.universe.goog" {
			t.Fatalf("ResolveUniverseDomain() = %q, want %q", got, "explicit.universe.goog")
		}
	})

	t.Run("reads GOOGLE_CLOUD_UNIVERSE_DOMAIN", func(t *testing.T) {
		t.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", "custom.universe.goog")
		if got := ResolveUniverseDomain(""); got != "custom.universe.goog" {
			t.Fatalf("ResolveUniverseDomain(\"\") = %q, want %q", got, "custom.universe.goog")
		}
	})
}

func TestIsCustomUniverse(t *testing.T) {
	if IsCustomUniverse("") {
		t.Errorf("IsCustomUniverse(\"\") = true, want false")
	}
	if IsCustomUniverse("googleapis.com") {
		t.Errorf("IsCustomUniverse(\"googleapis.com\") = true, want false")
	}
	if !IsCustomUniverse("custom.universe.goog") {
		t.Errorf("IsCustomUniverse(\"custom.universe.goog\") = false, want true")
	}
}

func TestServiceHostAndReplace(t *testing.T) {
	if got := ServiceHost("compute", "custom.universe.goog"); got != "compute.custom.universe.goog" {
		t.Errorf("ServiceHost() = %q, want %q", got, "compute.custom.universe.goog")
	}
	if got := ReplaceUniverseDomain("https://storage.googleapis.com/storage/v1/", "custom.universe.goog"); got != "https://storage.custom.universe.goog/storage/v1/" {
		t.Errorf("ReplaceUniverseDomain() = %q, want %q", got, "https://storage.custom.universe.goog/storage/v1/")
	}
}
