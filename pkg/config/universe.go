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
	"os"
	"strings"
)

// DefaultUniverseDomain is the API host suffix of the public Google Cloud
// universe. An empty universe domain is equivalent to this value.
const DefaultUniverseDomain = "googleapis.com"

const (
	// UniverseDomainEnvVar is the cross-SDK Google environment variable naming
	// the universe domain. It is also read directly by google.golang.org/api,
	// at lower precedence than option.WithUniverseDomain.
	UniverseDomainEnvVar = "GOOGLE_CLOUD_UNIVERSE_DOMAIN"

	// UniversePrefixEnvVar names the universe prefix. Unlike
	// UniverseDomainEnvVar this is defined by Config Connector: Google
	// publishes no standard environment variable for the project prefix. If one
	// is standardized later, prefer it and keep this as an alias.
	UniversePrefixEnvVar = "GOOGLE_CLOUD_UNIVERSE_PREFIX"
)

// GetUniverseDomain returns the universe domain to target, resolving
// UniverseDomain, then UniverseDomainEnvVar, then DefaultUniverseDomain.
func (c *ControllerConfig) GetUniverseDomain() string {
	if c.UniverseDomain != "" {
		return c.UniverseDomain
	}
	if fromEnv := os.Getenv(UniverseDomainEnvVar); fromEnv != "" {
		return fromEnv
	}
	return DefaultUniverseDomain
}

// GetUniversePrefix returns the universe prefix applied to project IDs and
// service-agent emails, resolving UniversePrefix then UniversePrefixEnvVar.
// The public universe has no prefix, so the zero value is the empty string.
//
// The prefix is never derived from the universe domain: the two are
// independent values and a universe may use any combination of them.
func (c *ControllerConfig) GetUniversePrefix() string {
	if c.UniversePrefix != "" {
		return c.UniversePrefix
	}
	return os.Getenv(UniversePrefixEnvVar)
}

// IsDefaultUniverse reports whether we are targeting the public Google Cloud
// universe. Callers should branch on this rather than on UniverseDomain being
// empty, so that an explicitly configured "googleapis.com" behaves identically
// to no configuration at all.
func (c *ControllerConfig) IsDefaultUniverse() bool {
	return c.GetUniverseDomain() == DefaultUniverseDomain
}

// Endpoint rewrites an endpoint expressed against the public universe into the
// configured universe, and returns it unchanged when targeting the public
// universe.
//
//	Endpoint("networksecurity.googleapis.com:443")
//	  -> "networksecurity.example.test:443"
//
// It substitutes the host suffix rather than matching a service name, so the
// regional ("us-central1-aiplatform.googleapis.com:443"), scheme-prefixed
// ("https://compute.googleapis.com") and preview-prefixed forms all work
// without special cases.
//
// Use this for the endpoint overrides that some direct controllers pass to
// option.WithEndpoint. Controllers that do not override their endpoint need no
// change: RESTClientOptions and GRPCClientOptions pass the universe domain to
// the client library, which derives the endpoint itself.
func (c *ControllerConfig) Endpoint(defaultEndpoint string) string {
	if c.IsDefaultUniverse() {
		return defaultEndpoint
	}
	return strings.ReplaceAll(defaultEndpoint, DefaultUniverseDomain, c.GetUniverseDomain())
}
