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

package discovery

import (
	"strings"
)

// ServiceGroupMapping maps a GCP Service name (from Service Usage) to corresponding KCC API Groups.
var ServiceGroupMapping = map[string][]string{
	"compute.googleapis.com": {
		"compute.cnrm.cloud.google.com",
	},
	"container.googleapis.com": {
		"container.cnrm.cloud.google.com",
		"containerattached.cnrm.cloud.google.com",
	},
	"storage.googleapis.com": {
		"storage.cnrm.cloud.google.com",
	},
	"pubsub.googleapis.com": {
		"pubsub.cnrm.cloud.google.com",
	},
	"iam.googleapis.com": {
		"iam.cnrm.cloud.google.com",
	},
	"bigquery.googleapis.com": {
		"bigquery.cnrm.cloud.google.com",
		"bigquerydatatransfer.cnrm.cloud.google.com",
		"bigqueryreservation.cnrm.cloud.google.com",
	},
	"cloudkms.googleapis.com": {
		"kms.cnrm.cloud.google.com",
	},
	"logging.googleapis.com": {
		"logging.cnrm.cloud.google.com",
	},
	"monitoring.googleapis.com": {
		"monitoring.cnrm.cloud.google.com",
	},
	"dns.googleapis.com": {
		"dns.cnrm.cloud.google.com",
	},
	"sqladmin.googleapis.com": {
		"sql.cnrm.cloud.google.com",
	},
	"dataproc.googleapis.com": {
		"dataproc.cnrm.cloud.google.com",
	},
	"networksecurity.googleapis.com": {
		"networksecurity.cnrm.cloud.google.com",
	},
	"networkconnectivity.googleapis.com": {
		"networkconnectivity.cnrm.cloud.google.com",
	},
	"accesscontextmanager.googleapis.com": {
		"accesscontextmanager.cnrm.cloud.google.com",
	},
	"orgpolicy.googleapis.com": {
		"orgpolicy.cnrm.cloud.google.com",
	},
	"artifactregistry.googleapis.com": {
		"artifactregistry.cnrm.cloud.google.com",
	},
	"servicedirectory.googleapis.com": {
		"servicedirectory.cnrm.cloud.google.com",
	},
	"serviceusage.googleapis.com": {
		"serviceusage.cnrm.cloud.google.com",
	},
	"cloudresourcemanager.googleapis.com": {
		"resourcemanager.cnrm.cloud.google.com",
	},
	"essentialcontacts.googleapis.com": {
		"essentialcontacts.cnrm.cloud.google.com",
	},
}

// CoreAPIGroups are always included regardless of discovery state to ensure operator functionality.
var CoreAPIGroups = map[string]bool{
	"core.cnrm.cloud.google.com":      true,
	"customize.core.cnrm.cloud.google.com": true,
}

// GatingReport details the outcome of CRD discovery filtering.
type GatingReport struct {
	TotalCRDsInCatalog int
	RetainedCRDs       int
	PrunedCRDs         int
	RetainedAPIGroups  []string
	UnsupportedGroups  []string
}

// FilterCRDNames returns the subset of CRDs that correspond to enabled services in the sovereign universe.
func FilterCRDNames(allCRDNames []string, enabledServices []string) ([]string, GatingReport) {
	// Build set of allowed API groups
	allowedGroups := make(map[string]bool)
	for group := range CoreAPIGroups {
		allowedGroups[group] = true
	}

	for _, svc := range enabledServices {
		svcNormalized := strings.ToLower(strings.TrimSpace(svc))
		if !strings.HasSuffix(svcNormalized, ".googleapis.com") {
			svcNormalized += ".googleapis.com"
		}
		if groups, exists := ServiceGroupMapping[svcNormalized]; exists {
			for _, g := range groups {
				allowedGroups[g] = true
			}
		}
	}

	var retained []string
	var prunedCount int
	seenGroups := make(map[string]bool)
	unsupportedGroupsMap := make(map[string]bool)

	for _, crdName := range allCRDNames {
		group := extractGroupFromCRD(crdName)
		if allowedGroups[group] {
			retained = append(retained, crdName)
			seenGroups[group] = true
		} else {
			prunedCount++
			unsupportedGroupsMap[group] = true
		}
	}

	var retainedGroups []string
	for g := range seenGroups {
		retainedGroups = append(retainedGroups, g)
	}

	var unsupportedGroups []string
	for g := range unsupportedGroupsMap {
		unsupportedGroups = append(unsupportedGroups, g)
	}

	report := GatingReport{
		TotalCRDsInCatalog: len(allCRDNames),
		RetainedCRDs:       len(retained),
		PrunedCRDs:         prunedCount,
		RetainedAPIGroups:  retainedGroups,
		UnsupportedGroups:  unsupportedGroups,
	}

	return retained, report
}

// extractGroupFromCRD parses the API group from a standard CRD name (e.g. "computeinstances.compute.cnrm.cloud.google.com").
func extractGroupFromCRD(crdName string) string {
	parts := strings.Split(crdName, ".")
	if len(parts) > 1 {
		return strings.Join(parts[1:], ".")
	}
	return crdName
}
