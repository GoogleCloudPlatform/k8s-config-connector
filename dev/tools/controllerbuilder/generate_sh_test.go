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
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateShFormat(t *testing.T) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatalf("failed to find git repo root: %v", err)
	}
	repoRoot := strings.TrimSpace(string(out))
	apisDir := filepath.Join(repoRoot, "apis")

	var generateScripts []string
	err = filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "generate.sh" {
			generateScripts = append(generateScripts, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to walk apis directory: %v", err)
	}

	if len(generateScripts) == 0 {
		t.Fatalf("no generate.sh scripts found under %s", apisDir)
	}

	subcommandRegex := regexp.MustCompile(`(go\s+run\s+\S*|\./controllerbuilder)\s+(generate-types|generate-mapper|generate-fuzzer|generate-direct|prompt|scaffold)`)

	for _, scriptPath := range generateScripts {
		relPath, _ := filepath.Rel(repoRoot, scriptPath)
		contentBytes, err := os.ReadFile(scriptPath)
		if err != nil {
			t.Errorf("%s: unable to read file: %v", relPath, err)
			continue
		}
		content := string(contentBytes)

		// Skip empty scripts or scripts with no active code
		hasActiveCode := false
		for _, line := range strings.Split(content, "\n") {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" && !strings.HasPrefix(trimmed, "#") && !strings.HasPrefix(trimmed, "set ") {
				hasActiveCode = true
				break
			}
		}
		if !hasActiveCode {
			continue
		}

		// Check if controllerbuilder subcommands are used
		hasControllerBuilderCalls := stringsContainsSubcommand(content)

		if hasControllerBuilderCalls {
			if !strings.Contains(content, "CONTROLLERBUILDER=") {
				t.Errorf("%s: missing CONTROLLERBUILDER header definition. Must set CONTROLLERBUILDER variable with fallback logic.", relPath)
			}
		}

		matches := subcommandRegex.FindAllStringSubmatch(content, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				t.Errorf("%s: direct execution '%s' detected. Must use ${CONTROLLERBUILDER} instead of 'go run' or literal binary path.", relPath, match[0])
			}
		}
	}
}

func stringsContainsSubcommand(content string) bool {
	subcommands := []string{"generate-types", "generate-mapper", "generate-fuzzer", "generate-direct", "prompt", "scaffold"}
	for _, sub := range subcommands {
		if strings.Contains(content, sub) {
			return true
		}
	}
	return false
}
