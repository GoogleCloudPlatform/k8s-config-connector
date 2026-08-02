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

package lint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateScriptsLocation(t *testing.T) {
	apisDir := "../../apis"

	err := filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if info.Name() == "generate.sh" {
			relPath, err := filepath.Rel(apisDir, path)
			if err != nil {
				return err
			}

			// Normalize path separators to forward slash
			relSlash := filepath.ToSlash(relPath)
			parts := strings.Split(relSlash, "/")

			// Valid path must be exactly "apis/<service>/generate.sh" (relative parts: ["<service>", "generate.sh"])
			if len(parts) != 2 {
				t.Errorf("generate.sh script at %s is in an invalid location. generate.sh must live directly under apis/<service>/ (e.g., apis/<service>/generate.sh) to support concurrent generation.", relPath)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking apis directory: %v", err)
	}
}

func TestGenerateScripts(t *testing.T) {
	apisDir := "../../apis"

	err := filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || info.Name() != "generate.sh" {
			return nil
		}

		relPath, err := filepath.Rel(apisDir, path)
		if err != nil {
			return err
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read %s: %v", relPath, err)
		}
		content := string(contentBytes)

		t.Run(relPath, func(t *testing.T) {
			// 1) Must call ./generate-proto.sh
			if !strings.Contains(content, "generate-proto.sh") {
				t.Errorf("generate.sh script at %s is missing call to ./generate-proto.sh", relPath)
			}

			// 2) Must call generate-types (or openapi-to-krm)
			if !strings.Contains(content, "generate-types") && !strings.Contains(content, "openapi-to-krm") {
				t.Errorf("generate.sh script at %s is missing call to generate-types", relPath)
			}

			// 3) Must call dev/tasks/generate-crds
			if !strings.Contains(content, "generate-crds") {
				t.Errorf("generate.sh script at %s is missing call to dev/tasks/generate-crds", relPath)
			}

			// 4) If pkg/controller/direct/<service>/mapper.generated.go exists, generate.sh MUST call generate-mapper (or openapi-to-krm)
			relSlash := filepath.ToSlash(relPath)
			service := strings.Split(relSlash, "/")[0]
			mapperPath := filepath.Join("../../pkg/controller/direct", service, "mapper.generated.go")
			if _, err := os.Stat(mapperPath); err == nil {
				if !strings.Contains(content, "generate-mapper") && !strings.Contains(content, "openapi-to-krm") {
					t.Errorf("generate.sh script at %s is missing call to generate-mapper (mapper.generated.go exists at %s)", relPath, mapperPath)
				}
			}

			// 5) Check CONTROLLERBUILDER variable header definition
			if strings.Contains(content, "generate-types") || strings.Contains(content, "generate-mapper") {
				if !strings.Contains(content, "CONTROLLERBUILDER=") {
					t.Errorf("generate.sh script at %s is missing CONTROLLERBUILDER header definition.", relPath)
				}
			}

			// 6) Disallow 'go run' calls targeting controllerbuilder except the CONTROLLERBUILDER= fallback line
			for _, line := range strings.Split(content, "\n") {
				trimmed := strings.TrimSpace(line)
				if strings.HasPrefix(trimmed, "#") {
					continue
				}
				if strings.Contains(trimmed, "go run") && strings.Contains(trimmed, "controllerbuilder") {
					if !strings.Contains(trimmed, "CONTROLLERBUILDER=") {
						t.Errorf("generate.sh script at %s: forbidden 'go run' call to controllerbuilder detected on line: '%s'. Must use ${CONTROLLERBUILDER} instead.", relPath, trimmed)
					}
				}
			}
		})

		return nil
	})
	if err != nil {
		t.Fatalf("error walking apis directory: %v", err)
	}
}
