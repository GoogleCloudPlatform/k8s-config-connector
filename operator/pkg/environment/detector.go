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

package environment

import (
	"context"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EnvironmentType classifies the runtime cloud environment.
type EnvironmentType string

const (
	// CommercialGCP is standard public Google Cloud (googleapis.com).
	CommercialGCP EnvironmentType = "CommercialGCP"
	// SovereignCloud is an isolated Google Sovereign Cloud partition.
	SovereignCloud EnvironmentType = "SovereignCloud"
	// AirgappedTPC is an isolated air-gapped Trusted Partner Cloud.
	AirgappedTPC EnvironmentType = "AirgappedTPC"
)

// EnvironmentInfo holds detected environment properties.
type EnvironmentInfo struct {
	Type           EnvironmentType
	UniverseDomain string
	IsSovereign    bool
	PartitionID    string
}

// DetectEnvironment inspects the environment and returns classification details.
func DetectEnvironment(ctx context.Context, kubeClient client.Client) (*EnvironmentInfo, error) {
	universeDomain := gcp.GetUniverseDomain()
	if envVar := os.Getenv(gcp.UniverseDomainEnvVar); envVar != "" {
		universeDomain = strings.TrimSpace(envVar)
	}

	info := &EnvironmentInfo{
		Type:           CommercialGCP,
		UniverseDomain: universeDomain,
		IsSovereign:    false,
	}

	if universeDomain == "" || universeDomain == gcp.DefaultUniverseDomain {
		return info, nil
	}

	// Non-default universe domain indicates sovereign or air-gapped environment
	info.IsSovereign = true
	if strings.Contains(universeDomain, "sovereign") {
		info.Type = SovereignCloud
	} else {
		info.Type = AirgappedTPC
	}

	// Detect partition if project ID contains colon prefix
	if projectID := os.Getenv("PROJECT_ID"); projectID != "" {
		if idx := strings.Index(projectID, ":"); idx != -1 {
			info.PartitionID = projectID[:idx]
		}
	}

	return info, nil
}
