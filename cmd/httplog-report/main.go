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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport"
)

func main() {
	var (
		fixtureFlag   = flag.String("fixture", "", "Path to a fixture directory or parent directory to scan recursively for _http.log files")
		writeFlag     = flag.Bool("write", false, "Write _http_report.md in each matched fixture directory")
		validateFlag  = flag.Bool("validate", false, "Exit with non-zero status if any validation check (V1-V7) fails")
		verbosityFlag = flag.String("verbosity", "summary", "Terminal output verbosity: summary | mutations | full")
	)
	flag.Parse()

	var roots []string
	if *fixtureFlag != "" {
		roots = append(roots, *fixtureFlag)
	}
	roots = append(roots, flag.Args()...)
	if len(roots) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: httplog-report --fixture <dir> [--write] [--validate] [--verbosity=summary|mutations|full]")
		os.Exit(2)
	}

	var fixtureDirs []string
	for _, r := range roots {
		dirs, err := httplogreport.DiscoverFixtureDirs(r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error discovering fixtures in %s: %v\n", r, err)
			os.Exit(1)
		}
		fixtureDirs = append(fixtureDirs, dirs...)
	}

	if len(fixtureDirs) == 0 {
		fmt.Fprintf(os.Stderr, "No fixture directories containing create.yaml and _http.log found in %v\n", roots)
		os.Exit(1)
	}

	failedValidation := 0
	runtimeErrors := 0
	for _, dir := range fixtureDirs {
		report, err := httplogreport.AnalyzeFixture(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error analyzing %s: %v\n", dir, err)
			runtimeErrors++
			continue
		}

		if *writeFlag {
			outPath, err := httplogreport.WriteReport(dir, report)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing report in %s: %v\n", dir, err)
				runtimeErrors++
				continue
			}
			if len(fixtureDirs) == 1 {
				fmt.Printf("Wrote %s\n\n", outPath)
			}
		}

		if len(fixtureDirs) == 1 || (!report.Passed && *validateFlag) {
			fmt.Println(httplogreport.RenderTerminal(report, *verbosityFlag))
		} else {
			status := "PASS"
			if !report.Passed {
				status = "FAIL"
			}
			fmt.Printf("[%s] %s (%d entries, %d primary RPCs, %d spec diffs)\n",
				status, dir, report.TotalEntries, len(report.PrimaryEntries), len(report.SpecDiff))
		}

		if !report.Passed {
			failedValidation++
		}
	}

	if len(fixtureDirs) > 1 {
		fmt.Printf("\nProcessed %d fixture(s): %d passed, %d failed validation, %d runtime error(s).\n",
			len(fixtureDirs), len(fixtureDirs)-failedValidation-runtimeErrors, failedValidation, runtimeErrors)
	}

	if runtimeErrors > 0 || (*validateFlag && failedValidation > 0) {
		os.Exit(1)
	}
}
