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

package tfprovider_test

import (
	"os"
	"testing"

	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
)

func TestUniverseDomainConfiguration(t *testing.T) {
	orig := os.Getenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN")
	defer os.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", orig)

	// Test default empty universe domain
	os.Unsetenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN")
	cfg := tfprovider.NewConfig()
	if cfg.UniverseDomain != "" {
		t.Fatalf("expected empty UniverseDomain by default, got %q", cfg.UniverseDomain)
	}

	// Test universe domain populated from environment
	expectedDomain := "custom.universe.goog"
	os.Setenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN", expectedDomain)
	cfgWithEnv := tfprovider.NewConfig()
	if cfgWithEnv.UniverseDomain != expectedDomain {
		t.Fatalf("expected UniverseDomain %q, got %q", expectedDomain, cfgWithEnv.UniverseDomain)
	}
}
