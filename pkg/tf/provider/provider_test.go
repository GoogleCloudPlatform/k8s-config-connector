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

package tfprovider

import "testing"

func TestBuildProviderConfig_universeDomain(t *testing.T) {
	tests := []struct {
		name           string
		universeDomain string
		wantPresent    bool
		wantValue      string
	}{
		{
			name:        "unset_attributeIsAbsent",
			wantPresent: false,
		},
		{
			name:           "otherUniverse_attributeIsSet",
			universeDomain: "example-apis.test",
			wantPresent:    true,
			wantValue:      "example-apis.test",
		},
		{
			// The provider rewrites transport_tpg.DefaultBasePaths in place
			// when universe_domain is set. Passing "googleapis.com" would be a
			// no-op substitution, but it would still put the public universe
			// through a code path it does not take today, so we do not set it.
			name:           "explicitPublicUniverse_attributeIsSet",
			universeDomain: "googleapis.com",
			wantPresent:    true,
			wantValue:      "googleapis.com",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfgMap := buildProviderConfig(Config{UniverseDomain: tc.universeDomain})

			got, present := cfgMap["universe_domain"]
			if present != tc.wantPresent {
				t.Fatalf("universe_domain present = %v, want %v (cfgMap=%v)", present, tc.wantPresent, cfgMap)
			}
			if !tc.wantPresent {
				return
			}
			if got != tc.wantValue {
				t.Errorf("universe_domain = %v, want %v", got, tc.wantValue)
			}
		})
	}
}

// TestBuildProviderConfig_unrelatedAttributesAreUnchanged guards the no-op
// guarantee for the Terraform path: adding universe support must not perturb
// the attributes KCC already sets.
func TestBuildProviderConfig_unrelatedAttributesAreUnchanged(t *testing.T) {
	config := Config{
		GCPAccessToken:      "token",
		Scopes:              []string{"scope-a", "scope-b"},
		UserProjectOverride: true,
		BillingProject:      "billing-project",
	}

	withoutUniverse := buildProviderConfig(config)

	config.UniverseDomain = "example-apis.test"
	withUniverse := buildProviderConfig(config)

	for _, key := range []string{"access_token", "user_project_override", "billing_project"} {
		if !equal(withoutUniverse[key], withUniverse[key]) {
			t.Errorf("%s changed when universe was set: %v -> %v", key, withoutUniverse[key], withUniverse[key])
		}
	}
	if len(withUniverse) != len(withoutUniverse)+1 {
		t.Errorf("setting a universe changed %d attributes, want exactly 1 added", len(withUniverse)-len(withoutUniverse))
	}
}

func equal(a, b interface{}) bool {
	return a == b
}
