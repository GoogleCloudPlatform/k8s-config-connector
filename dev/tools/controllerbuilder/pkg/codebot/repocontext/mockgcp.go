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

package repocontext

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// gathers example mock content from multiple directories.
func GatherMockContents(mockgcpPath string, examples int) ([]string, error) {
	var exampleMockContents []string
	var exampleMockDirPaths []string

	// Walk through the mockgcp directory.
	err := filepath.Walk(mockgcpPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Look for directories that might contain mocks
		if info.IsDir() && strings.HasPrefix(info.Name(), "mock") && info.Name() != "mockgcp" && len(exampleMockDirPaths) < examples {
			exampleMockDirPaths = append(exampleMockDirPaths, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walking mockgcp directory: %w", err)
	}

	// Read up to "examples" files from each directory.
	for _, dirPath := range exampleMockDirPaths {
		files, err := os.ReadDir(dirPath)
		if err != nil {
			return nil, fmt.Errorf("reading directory %s: %w", dirPath, err)
		}

		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
				filePath := filepath.Join(dirPath, file.Name())
				fileContent, err := os.ReadFile(filePath)
				if err != nil {
					return nil, fmt.Errorf("reading file %s: %w", filePath, err)
				}
				exampleMockContents = append(exampleMockContents, string(fileContent))
			}
		}
	}

	return exampleMockContents, nil
}