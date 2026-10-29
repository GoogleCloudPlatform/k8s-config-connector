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

func (g gkeClusterInfo) mclGSAEmail() string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", mclServiceAccountName, g.project)
}

func (g gkeClusterInfo) syncerGSAEmail() string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", syncerServiceAccountName, g.project)
}

// gcsBucketLocation returns a valid GCS bucket region derived from the GKE cluster location.
// If the GKE cluster is zonal (e.g. "us-central1-a"), the zone suffix is stripped ("us-central1").
func (g gkeClusterInfo) gcsBucketLocation() string {
	if isZonalGKELocation(g.location) {
		lastHyphen := strings.LastIndex(g.location, "-")
		return g.location[:lastHyphen]
	}

	return g.location
}

func isZonalGKELocation(location string) bool {
	if strings.Count(location, "-") != 2 {
		return false
	}

	lastHyphen := strings.LastIndex(location, "-")
	suffix := location[lastHyphen+1:]
	return len(suffix) == 1 && suffix[0] >= 'a' && suffix[0] <= 'z'
}

// dedupeProjects returns the first gkeClusterInfo per unique GCP project while preserving order.
func dedupeProjects(clusters []gkeClusterInfo) []gkeClusterInfo {
	seen := make(map[string]bool, len(clusters))
	unique := make([]gkeClusterInfo, 0, len(clusters))
	for _, c := range clusters {
		if !seen[c.project] {
			seen[c.project] = true
			unique = append(unique, c)
		}
	}

	return unique
}
