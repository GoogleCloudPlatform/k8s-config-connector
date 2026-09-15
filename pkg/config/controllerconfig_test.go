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

package config

import (
	"reflect"
	"testing"

	"google.golang.org/api/option"
)

// universeDomainOptionType is the concrete type returned by
// option.WithUniverseDomain. The type is unexported and the DialSettings it
// applies to lives in an internal package, so the option cannot be inspected by
// applying it. Comparing against the type of a known-good option is the
// supported way to recognise it from outside the module.
var universeDomainOptionType = reflect.TypeOf(option.WithUniverseDomain(""))

// universeDomainFrom returns the universe domain carried by opts, or "" if no
// universe domain option is present.
func universeDomainFrom(t *testing.T, opts []option.ClientOption) string {
	t.Helper()

	found := ""
	for _, o := range opts {
		if reflect.TypeOf(o) != universeDomainOptionType {
			continue
		}
		if found != "" {
			t.Fatalf("options carry more than one universe domain")
		}
		// withUniverseDomain is a defined string type.
		found = reflect.ValueOf(o).String()
	}
	return found
}

func TestRESTClientOptions_universeDomain(t *testing.T) {
	tests := []struct {
		name     string
		config   ControllerConfig
		expected string
	}{
		{
			name:     "unset_noUniverseDomainOption",
			config:   ControllerConfig{},
			expected: "",
		},
		{
			name:     "explicitPublicUniverse_noUniverseDomainOption",
			config:   ControllerConfig{UniverseDomain: DefaultUniverseDomain},
			expected: "",
		},
		{
			name:     "otherUniverse_setsUniverseDomainOption",
			config:   ControllerConfig{UniverseDomain: "example-apis.test"},
			expected: "example-apis.test",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts, err := tc.config.RESTClientOptions()
			if err != nil {
				t.Fatalf("RESTClientOptions() returned error: %v", err)
			}
			if got := universeDomainFrom(t, opts); got != tc.expected {
				t.Errorf("universe domain = %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestGRPCClientOptions_universeDomain(t *testing.T) {
	tests := []struct {
		name     string
		config   ControllerConfig
		expected string
	}{
		{
			name:     "unset_noUniverseDomainOption",
			config:   ControllerConfig{},
			expected: "",
		},
		{
			name:     "explicitPublicUniverse_noUniverseDomainOption",
			config:   ControllerConfig{UniverseDomain: DefaultUniverseDomain},
			expected: "",
		},
		{
			name:     "otherUniverse_setsUniverseDomainOption",
			config:   ControllerConfig{UniverseDomain: "example-apis.test"},
			expected: "example-apis.test",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts, err := tc.config.GRPCClientOptions()
			if err != nil {
				t.Fatalf("GRPCClientOptions() returned error: %v", err)
			}
			if got := universeDomainFrom(t, opts); got != tc.expected {
				t.Errorf("universe domain = %q, want %q", got, tc.expected)
			}
		})
	}
}

// TestClientOptions_publicUniverseIsUnchanged is the no-op guarantee: with no
// universe configured, and with the public universe configured explicitly, the
// option list must be exactly what it was before universe support existed.
//
// Asserting on the option count rather than only on the universe domain catches
// an implementation that appends a no-op option in the public universe, which
// would be harmless today but is the kind of thing that drifts.
func TestClientOptions_publicUniverseIsUnchanged(t *testing.T) {
	base := ControllerConfig{UserAgent: "kcc/test"}
	explicitPublic := ControllerConfig{UserAgent: "kcc/test", UniverseDomain: DefaultUniverseDomain}

	t.Run("REST", func(t *testing.T) {
		baseOpts, err := base.RESTClientOptions()
		if err != nil {
			t.Fatalf("RESTClientOptions() returned error: %v", err)
		}
		publicOpts, err := explicitPublic.RESTClientOptions()
		if err != nil {
			t.Fatalf("RESTClientOptions() returned error: %v", err)
		}
		if len(baseOpts) != len(publicOpts) {
			t.Errorf("option count differs: unset=%d, explicit googleapis.com=%d", len(baseOpts), len(publicOpts))
		}
		if got := universeDomainFrom(t, baseOpts); got != "" {
			t.Errorf("unset config produced universe domain %q, want none", got)
		}
	})

	t.Run("GRPC", func(t *testing.T) {
		baseOpts, err := base.GRPCClientOptions()
		if err != nil {
			t.Fatalf("GRPCClientOptions() returned error: %v", err)
		}
		publicOpts, err := explicitPublic.GRPCClientOptions()
		if err != nil {
			t.Fatalf("GRPCClientOptions() returned error: %v", err)
		}
		if len(baseOpts) != len(publicOpts) {
			t.Errorf("option count differs: unset=%d, explicit googleapis.com=%d", len(baseOpts), len(publicOpts))
		}
		if got := universeDomainFrom(t, baseOpts); got != "" {
			t.Errorf("unset config produced universe domain %q, want none", got)
		}
	})
}
