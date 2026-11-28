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

package environment

import "testing"

func TestDetectFromEnv(t *testing.T) {
	tests := []struct {
		name           string
		env            map[string]string
		wantType       CloudType
		wantDomain     string
		wantCustom     bool
		checkCRD       string
		wantCRDAllowed bool
	}{
		{
			name:           "default commercial environment",
			env:            map[string]string{},
			wantType:       CommercialGCP,
			wantDomain:     "googleapis.com",
			wantCustom:     false,
			checkCRD:       "spannerinstances.spanner.cnrm.cloud.google.com",
			wantCRDAllowed: true,
		},
		{
			name: "sovereign cloud with supported services filter",
			env: map[string]string{
				"GOOGLE_CLOUD_UNIVERSE_DOMAIN": "apis-eu0.sovereign.goog",
				"KCC_SUPPORTED_SERVICES":       "compute,storage,iam",
			},
			wantType:       SovereignCloud,
			wantDomain:     "apis-eu0.sovereign.goog",
			wantCustom:     true,
			checkCRD:       "spannerinstances.spanner.cnrm.cloud.google.com",
			wantCRDAllowed: false,
		},
		{
			name: "sovereign cloud allows listed service",
			env: map[string]string{
				"GOOGLE_CLOUD_UNIVERSE_DOMAIN": "apis-fr0.sovereign.goog",
				"KCC_SUPPORTED_SERVICES":       "compute.apis-fr0.sovereign.goog,storage",
			},
			wantType:       SovereignCloud,
			wantDomain:     "apis-fr0.sovereign.goog",
			wantCustom:     true,
			checkCRD:       "computeinstances.compute.cnrm.cloud.google.com",
			wantCRDAllowed: true,
		},
		{
			name: "airgapped TPC environment always allows core CRDs",
			env: map[string]string{
				"UNIVERSE_DOMAIN":        "cloud.airgap.internal",
				"KCC_SUPPORTED_SERVICES": "compute",
			},
			wantType:       AirgappedTPC,
			wantDomain:     "cloud.airgap.internal",
			wantCustom:     true,
			checkCRD:       "configconnectors.core.cnrm.cloud.google.com",
			wantCRDAllowed: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := DetectFromEnv(func(k string) string {
				return tc.env[k]
			})
			if cfg.Type != tc.wantType {
				t.Errorf("Type = %q, want %q", cfg.Type, tc.wantType)
			}
			if cfg.UniverseDomain != tc.wantDomain {
				t.Errorf("UniverseDomain = %q, want %q", cfg.UniverseDomain, tc.wantDomain)
			}
			if cfg.IsCustomUniverse() != tc.wantCustom {
				t.Errorf("IsCustomUniverse() = %v, want %v", cfg.IsCustomUniverse(), tc.wantCustom)
			}
			if got := cfg.IsCRDGroupSupported(tc.checkCRD); got != tc.wantCRDAllowed {
				t.Errorf("IsCRDGroupSupported(%q) = %v, want %v", tc.checkCRD, got, tc.wantCRDAllowed)
			}
		})
	}
}
