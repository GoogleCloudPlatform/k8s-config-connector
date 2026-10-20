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

import (
	"context"
	"strings"
	"testing"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestUniverseDomainPropagation(t *testing.T) {
	ctx := context.Background()

	// Set universe domain via Config
	cfg := UnitTestConfig()
	cfg.UniverseDomain = "custom.universe.goog"

	p, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}
	if p == nil {
		t.Fatalf("New() returned nil provider")
	}

	// Verify that MonitoringBasePath in DefaultBasePaths was updated with the universe domain
	basePath := transport_tpg.DefaultBasePaths[transport_tpg.MonitoringBasePathKey]
	if !strings.Contains(basePath, "custom.universe.goog") {
		t.Errorf("expected MonitoringBasePath to contain custom.universe.goog, got %s", basePath)
	}
}
