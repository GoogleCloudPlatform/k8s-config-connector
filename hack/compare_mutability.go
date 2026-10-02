package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
)

type KCCMutability map[string][]string

type GCPResourceInfo struct {
	HasUpdateMethod          bool                              `json:"has_update_method"`
	ImmutableFields          []string                          `json:"immutable_fields"`
	MutableFields            []string                          `json:"mutable_fields"`
	PotentiallyMutableFields []string                          `json:"potentially_mutable_fields"`
	ProtoMessage             string                            `json:"proto_message"`
	SetterMethods            map[string]map[string]interface{} `json:"setter_methods"`
	UpdateRequestMessage     string                            `json:"update_request_message"`
}

type GCPMutability map[string]GCPResourceInfo

type SimplifiedReportItem struct {
	ResourceName               string   `json:"resource_name"`
	ShouldBeMutable            []string `json:"should_be_mutable"`
	PotentiallyShouldBeMutable []string `json:"potentially_should_be_mutable"`
}

// Map KCC field path to GCP field path
// Rules: camelCase -> snake_case; [fieldName]Ref -> field_name
func mapKCCToGCP(kccPath string) string {
	segments := strings.Split(kccPath, ".")
	var gcpSegments []string
	for _, s := range segments {
		clean := s
		if strings.HasSuffix(s, "Ref") {
			clean = strings.TrimSuffix(s, "Ref")
		}
		gcpSegments = append(gcpSegments, text.AsSnakeCase(clean))
	}
	return strings.Join(gcpSegments, ".")
}

func isCoreHierarchicalRef(field string) bool {
	lower := strings.ToLower(field)
	coreParents := []string{"projectref", "folderref", "organizationref", "billingaccountref", "locationref", "regionref", "zoneref"}
	for _, p := range coreParents {
		if lower == p {
			return true
		}
	}
	return false
}

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	// 1. Load Data
	kccBytes, err := ioutil.ReadFile(filepath.Join(rootDir, "kcc_mutability.json"))
	if err != nil {
		fmt.Printf("Error reading kcc_mutability.json: %v\n", err)
		os.Exit(1)
	}
	var kccData KCCMutability
	json.Unmarshal(kccBytes, &kccData)

	gcpBytes, err := ioutil.ReadFile(filepath.Join(rootDir, "gcp_mutability.json"))
	if err != nil {
		fmt.Printf("Error reading gcp_mutability.json: %v\n", err)
		os.Exit(1)
	}
	var gcpData GCPMutability
	json.Unmarshal(gcpBytes, &gcpData)

	var simplifiedReports []SimplifiedReportItem

	// 2. Compare
	kinds := []string{}
	for k := range kccData {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)

	for _, kind := range kinds {
		kccImmutables := kccData[kind]
		gcpInfo, ok := gcpData[kind]
		if !ok {
			continue
		}

		item := SimplifiedReportItem{
			ResourceName:               kind,
			ShouldBeMutable:            []string{},
			PotentiallyShouldBeMutable: []string{},
		}
		
		gcpImmutablesSet := make(map[string]bool)
		for _, f := range gcpInfo.ImmutableFields {
			gcpImmutablesSet[f] = true
		}
		gcpMutablesSet := make(map[string]bool)
		for _, f := range gcpInfo.MutableFields {
			gcpMutablesSet[f] = true
		}

		for _, kccField := range kccImmutables {
			// Skip universally immutable identity/location fields
			if kccField == "name" || kccField == "location" || kccField == "region" || kccField == "zone" {
				continue
			}
			
			// Skip core parent references to reduce noise
			if isCoreHierarchicalRef(kccField) {
				continue
			}

			gcpField := mapKCCToGCP(kccField)
			
			// 1. Definite Drift (Verified via Setter)
			isMutable := false
			if gcpMutablesSet[gcpField] {
				isMutable = true
			} else {
				parts := strings.Split(gcpField, ".")
				for i := 1; i <= len(parts); i++ {
					parentPath := strings.Join(parts[:i], ".")
					if gcpMutablesSet[parentPath] {
						isMutable = true
						break
					}
				}
			}

			if isMutable {
				item.ShouldBeMutable = append(item.ShouldBeMutable, kccField)
			} else if len(gcpInfo.ImmutableFields) > 0 && !gcpImmutablesSet[gcpField] {
				// 2. Potential Drift (Missing IMMUTABLE annotation)
				item.PotentiallyShouldBeMutable = append(item.PotentiallyShouldBeMutable, kccField)
			}
		}

		if len(item.ShouldBeMutable) > 0 || len(item.PotentiallyShouldBeMutable) > 0 {
			simplifiedReports = append(simplifiedReports, item)
		}
	}

	// 3. Output
	outData, _ := json.MarshalIndent(simplifiedReports, "", "  ")
	ioutil.WriteFile("mutability_drift_report.json", outData, 0644)
	fmt.Printf("Generated simplified mutability drift report for %d resources.\n", len(simplifiedReports))
}