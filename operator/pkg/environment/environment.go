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

import (
	"os"
	"strings"
)

// CloudType classifies the runtime Google Cloud environment.
type CloudType string

const (
	// CommercialGCP represents standard commercial Google Cloud (googleapis.com).
	CommercialGCP CloudType = "CommercialGCP"
	// SovereignCloud represents a regional Google Sovereign Cloud or partner-operated universe.
	SovereignCloud CloudType = "SovereignCloud"
	// AirgappedTPC represents an air-gapped or disconnected Trusted Partner Cloud environment.
	AirgappedTPC CloudType = "AirgappedTPC"

	// DefaultUniverseDomain is the standard commercial GCP universe domain.
	DefaultUniverseDomain = "googleapis.com"
)

// Config holds the detected environment classification, active universe domain,
// and optional service allowlist for CRD filtering.
type Config struct {
	Type              CloudType
	UniverseDomain    string
	SupportedServices map[string]bool
}

// Detect inspects standard operator environment variables to classify the cloud environment.
func Detect() Config {
	return DetectFromEnv(os.Getenv)
}

// DetectFromEnv classifies the cloud environment using the provided getenv lookup function.
func DetectFromEnv(getenv func(string) string) Config {
	domain := strings.TrimSpace(getenv("GOOGLE_CLOUD_UNIVERSE_DOMAIN"))
	if domain == "" {
		domain = strings.TrimSpace(getenv("UNIVERSE_DOMAIN"))
	}
	if domain == "" {
		domain = DefaultUniverseDomain
	}

	cloudType := classifyDomain(domain)

	var supported map[string]bool
	if rawServices := strings.TrimSpace(getenv("KCC_SUPPORTED_SERVICES")); rawServices != "" {
		supported = make(map[string]bool)
		for _, tok := range strings.Split(rawServices, ",") {
			svc := strings.ToLower(strings.TrimSpace(tok))
			svc = strings.TrimSuffix(svc, ".googleapis.com")
			svc = strings.TrimSuffix(svc, "."+domain)
			svc = strings.TrimSuffix(svc, ".cnrm.cloud.google.com")
			if svc != "" {
				supported[svc] = true
			}
		}
	}

	return Config{
		Type:              cloudType,
		UniverseDomain:    domain,
		SupportedServices: supported,
	}
}

func classifyDomain(domain string) CloudType {
	d := strings.ToLower(strings.TrimSpace(domain))
	if d == "" || d == DefaultUniverseDomain {
		return CommercialGCP
	}
	if strings.HasSuffix(d, ".local") ||
		strings.HasSuffix(d, ".internal") ||
		strings.HasSuffix(d, ".airgap") ||
		strings.Contains(d, "airgap") ||
		strings.Contains(d, "disconnected") {
		return AirgappedTPC
	}
	return SovereignCloud
}

// IsCustomUniverse returns true when the environment uses a non-commercial universe domain.
func (c Config) IsCustomUniverse() bool {
	return c.UniverseDomain != "" && c.UniverseDomain != DefaultUniverseDomain
}

// IsCRDGroupSupported reports whether a KCC CRD name (e.g. " computeinstances.compute.cnrm.cloud.google.com")
// belongs to an enabled service in the current environment. Core operator CRDs are always supported.
func (c Config) IsCRDGroupSupported(crdName string) bool {
	if len(c.SupportedServices) == 0 {
		return true
	}
	name := strings.ToLower(strings.TrimSpace(crdName))
	const suffix = ".cnrm.cloud.google.com"
	if !strings.HasSuffix(name, suffix) {
		return true
	}
	if strings.HasSuffix(name, ".core.cnrm.cloud.google.com") {
		return true
	}
	trimmed := strings.TrimSuffix(name, suffix)
	parts := strings.Split(trimmed, ".")
	if len(parts) < 2 {
		return true
	}
	service := parts[len(parts)-1]
	return c.SupportedServices[service]
}
