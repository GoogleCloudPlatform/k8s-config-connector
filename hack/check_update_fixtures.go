package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type DiscrepancyReportItem struct {
	ResourceName string `json:"resource_name"`
	CheckFailed  bool   `json:"check_failed"`
	Reason       string `json:"reason"`
}

type DiscrepancyReport struct {
	Comment       string                  `json:"_comment"`
	Discrepancies []DiscrepancyReportItem `json:"discrepancies"`
}

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	gcpFile := filepath.Join(rootDir, "gcp_mutability.json")
	content, err := ioutil.ReadFile(gcpFile)
	if err != nil {
		fmt.Printf("Error reading gcp_mutability.json: %v\n", err)
		os.Exit(1)
	}

	var mutabilityData map[string]map[string]interface{}
	if err := json.Unmarshal(content, &mutabilityData); err != nil {
		fmt.Printf("Error unmarshaling gcp_mutability.json: %v\n", err)
		os.Exit(1)
	}

	var report []DiscrepancyReportItem
	testDataDir := filepath.Join(rootDir, "pkg/test/resourcefixture/testdata/basic")

	// Collect keys and sort them for deterministic output
	var keys []string
	for k := range mutabilityData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, kind := range keys {
		data := mutabilityData[kind]
		hasUpdate := data["has_update_method"].(bool)
		setters := data["setter_methods"].(map[string]interface{})

		if !hasUpdate && len(setters) == 0 {
			lowerKind := strings.ToLower(kind)
			foundUpdateYaml := false

			filepath.Walk(testDataDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}
				if !info.IsDir() && info.Name() == "update.yaml" {
					parts := strings.Split(path, string(os.PathSeparator))
					n := len(parts)
					if n >= 2 && strings.ToLower(parts[n-2]) == lowerKind {
						foundUpdateYaml = true
						return filepath.SkipDir
					}
					if n >= 3 && strings.ToLower(parts[n-3]) == lowerKind {
						foundUpdateYaml = true
						return filepath.SkipDir
					}
				}
				return nil
			})
			
			if foundUpdateYaml {
				report = append(report, DiscrepancyReportItem{
					ResourceName: kind,
					CheckFailed:  true,
					Reason:       "no update/patch/setter found but update is supported in fixture test cases; root-cause and add special handling for the resources in the GCP mutability generator",
				})
			}
		}
	}

	finalReport := DiscrepancyReport{
		Comment:       "Resources that successfully passed the check (e.g., resources that correctly lack an update.yaml when no update method was found) are not included in this report to keep it concise.",
		Discrepancies: report,
	}
	
	outFile := filepath.Join(rootDir, "discrepancy_report.json")
	outData, err := json.MarshalIndent(finalReport, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling report: %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(outFile, outData, 0644); err != nil {
		fmt.Printf("Error writing discrepancy_report.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated report with %d discrepancies to discrepancy_report.json\n", len(report))
}