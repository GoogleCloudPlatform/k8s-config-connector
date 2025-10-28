// Copyright 2025 Google LLC
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
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s PATH_TO_COVERAGE.OUT [PATH_TO_FILE_TO_COVER], "+
			"e.g. go run ./dev/tools/coverageanalyzer topic.coverage.out "+
			"third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/pubsub/resource_pubsub_topic.go\n",
			os.Args[0])
	}
	coverageFile := os.Args[1]

	if _, err := os.Stat(coverageFile); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Coverage file %s doesn't exist", coverageFile)
	}

	uncoveredBlocks, err := findUncoveredBlocks(coverageFile)
	if err != nil {
		log.Fatalf("Error finding uncovered blocks: %v", err)
	}
	if len(uncoveredBlocks) == 0 {
		fmt.Println("No uncovered blocks found. Great job!")
		return
	}

	branches := findBranchDetails(uncoveredBlocks)

	fmt.Println("\nDiscovered uncovered branches (error returns filtered out):")
	if len(branches) == 0 {
		fmt.Println("All uncovered branches were simple error returns. Nothing to report.")
		return
	}
	for _, branch := range branches {
		fmt.Printf("- Type: %s\n  File: %s:%d-%d\n  Desc: %s\n  Content:\n%s\n",
			branch.BranchType,
			branch.FileName,
			branch.StartLine,
			branch.EndLine,
			branch.Description,
			indent(branch.Content, "    "), // Indent content for readability
		)
	}
}
