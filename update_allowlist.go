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
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	promotableContent, err := ioutil.ReadFile("promotable_resources.txt")
	if err != nil {
		log.Fatal(err)
	}
	promotable := make(map[string]bool)
	for _, line := range strings.Split(string(promotableContent), "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			promotable[trimmed] = true
		}
	}

	allowlistContent, err := ioutil.ReadFile("scripts/resource-autogen/allowlist/allowlist.go")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(allowlistContent), "\n")
	newAlpha := []string{}
	newBeta := []string{}
	
	// Pre-populate newBeta with existing ones
	inBeta := false
	for _, line := range lines {
		if strings.Contains(line, "betaAllowlist = []string{") {
			inBeta = true
			continue
		}
		if inBeta && strings.Contains(line, "}") {
			inBeta = false
			continue
		}
		if inBeta {
			trimmed := strings.Trim(line, "\t \",")
			if trimmed != "" {
				newBeta = append(newBeta, trimmed)
			}
		}
	}

	inAlpha := false
	for _, line := range lines {
		if strings.Contains(line, "alphaAllowlist = []string{") {
			inAlpha = true
			continue
		}
		if inAlpha && strings.Contains(line, "}") {
			inAlpha = false
			continue
		}
		if inAlpha {
			trimmed := strings.Trim(line, "\t \",")
			if trimmed == "" {
				continue
			}
			if promotable[trimmed] {
				newBeta = append(newBeta, trimmed)
			} else {
				newAlpha = append(newAlpha, trimmed)
			}
		}
	}

	// Now construct the new file content
	var output strings.Builder
	skip := false
	for _, line := range lines {
		if strings.Contains(line, "alphaAllowlist = []string{") {
			output.WriteString(line + "\n")
			for _, item := range newAlpha {
				output.WriteString(fmt.Sprintf("\t\t\"%s\",\n", item))
			}
			skip = true
			continue
		}
		if skip && strings.Contains(line, "}") {
			output.WriteString(line + "\n")
			skip = false
			continue
		}
		if strings.Contains(line, "betaAllowlist = []string{") {
			output.WriteString(line + "\n")
			// Sort and dedup newBeta
			uniqueBeta := make(map[string]bool)
			sortedBeta := []string{}
			for _, item := range newBeta {
				if !uniqueBeta[item] {
					uniqueBeta[item] = true
					sortedBeta = append(sortedBeta, item)
				}
			}
			for _, item := range sortedBeta {
				output.WriteString(fmt.Sprintf("\t\t\"%s\",\n", item))
			}
			skip = true
			continue
		}
		if !skip {
			output.WriteString(line + "\n")
		}
	}

	err = ioutil.WriteFile("scripts/resource-autogen/allowlist/allowlist.go", []byte(output.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully updated allowlist.go")
}
