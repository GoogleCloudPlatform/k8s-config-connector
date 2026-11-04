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

package ha

import (
	"fmt"
	"strings"
)

// gkeClusterInfo holds the GCP project, location, and cluster name parsed from a GKE context.
type gkeClusterInfo struct {
	contextName string
	identity    string
	project     string
	location    string
	clusterName string
}

// parseGKEContext parses a standard gcloud kubeconfig context name of the form
// "gke_<project>_<location>_<cluster-name>" and derives its DNS-1123 cluster identity.
//
// TODO: Support custom kubeconfig context names that do not follow the default
// gcloud "gke_<project>_<location>_<cluster>" naming convention (e.g. by inspecting the
// kubeconfig AuthInfo/cluster metadata or accepting explicit project/location overrides).
func parseGKEContext(contextName string) (gkeClusterInfo, error) {
	trimmed := strings.TrimSpace(contextName)
	parts := strings.Split(trimmed, "_")
	if !isValidGKEContextParts(parts) {
		return gkeClusterInfo{}, fmt.Errorf(
			"context %q does not match expected GKE context format \"gke_<project>_<location>_<cluster>\"; run `gcloud container clusters get-credentials <cluster> --location=<location> --project=<project>`",
			contextName,
		)
	}

	identity, err := sanitizeClusterIdentity(trimmed)
	if err != nil {
		return gkeClusterInfo{}, err
	}

	return gkeClusterInfo{
		contextName: trimmed,
		identity:    identity,
		project:     parts[1],
		location:    parts[2],
		clusterName: parts[3],
	}, nil
}

func isValidGKEContextParts(parts []string) bool {
	if len(parts) != 4 {
		return false
	}

	hasGKEPrefix := parts[0] == "gke"
	hasNonEmptySegments := parts[1] != "" && parts[2] != "" && parts[3] != ""
	return hasGKEPrefix && hasNonEmptySegments
}
