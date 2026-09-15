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

package environment_test

import (
	"context"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/environment"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
)

func TestDetectEnvironment_Commercial(t *testing.T) {
	_ = os.Unsetenv(gcp.UniverseDomainEnvVar)
	_ = os.Unsetenv("PROJECT_ID")

	info, err := environment.DetectEnvironment(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.Type != environment.CommercialGCP {
		t.Errorf("expected CommercialGCP, got %v", info.Type)
	}
	if info.IsSovereign {
		t.Errorf("expected IsSovereign = false, got true")
	}
}

func TestDetectEnvironment_SovereignPartition(t *testing.T) {
	_ = os.Setenv(gcp.UniverseDomainEnvVar, "sovereign.universe.goog")
	_ = os.Setenv("PROJECT_ID", "eu0:sample-project")
	defer func() {
		_ = os.Unsetenv(gcp.UniverseDomainEnvVar)
		_ = os.Unsetenv("PROJECT_ID")
	}()

	info, err := environment.DetectEnvironment(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.Type != environment.SovereignCloud {
		t.Errorf("expected SovereignCloud, got %v", info.Type)
	}
	if !info.IsSovereign {
		t.Errorf("expected IsSovereign = true, got false")
	}
	if info.UniverseDomain != "sovereign.universe.goog" {
		t.Errorf("expected sovereign.universe.goog, got %s", info.UniverseDomain)
	}
	if info.PartitionID != "eu0" {
		t.Errorf("expected partition eu0, got %s", info.PartitionID)
	}
}

func TestDetectEnvironment_AirgappedTPC(t *testing.T) {
	_ = os.Setenv(gcp.UniverseDomainEnvVar, "custom.airgap.internal")
	_ = os.Setenv("PROJECT_ID", "sec1:secure-corp")
	defer func() {
		_ = os.Unsetenv(gcp.UniverseDomainEnvVar)
		_ = os.Unsetenv("PROJECT_ID")
	}()

	info, err := environment.DetectEnvironment(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.Type != environment.AirgappedTPC {
		t.Errorf("expected AirgappedTPC, got %v", info.Type)
	}
	if !info.IsSovereign {
		t.Errorf("expected IsSovereign = true, got false")
	}
	if info.PartitionID != "sec1" {
		t.Errorf("expected partition sec1, got %s", info.PartitionID)
	}
}
