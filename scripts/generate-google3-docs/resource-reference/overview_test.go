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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestOverviewIsAlphabetical(t *testing.T) {
	file, err := os.Open("overview.md")
	if err != nil {
		t.Error("Error opening file:", err)
	}
	defer file.Close()

	var resources []string
	scanner := bufio.NewReader(file)
	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("ReadString('\n') failed: %s", err)
		}

		if strings.Contains(line, "td><a href=\"/config-connector/docs/reference/resource-doc") {
			// Process the line to extract the resource name
			parts := strings.Split(line, "\">")
			if len(parts) == 2 {
				nameParts := strings.Split(parts[1], "<")
				name := nameParts[0]
				resources = append(resources, strings.TrimSpace(name))
			}
		}
	}

	// Check for sorting violations
	var violations []string
	for i := 0; i < len(resources)-1; i++ {
		if resources[i] > resources[i+1] {
			violations = append(violations, fmt.Sprintf("'%s' should come before '%s';", resources[i+1], resources[i]))
		}
	}

	if len(violations) == 0 {
		t.Log("The resource names are sorted alphabetically.")
	} else {
		t.Fatal("Found sorting violations:", violations)
	}
}
