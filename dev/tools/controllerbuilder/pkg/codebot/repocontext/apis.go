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

func GatherAPIContents(apisPath string, examples int) ([]string, error) {
	var exampleAPIsContents []string
	var exampleAPIsDirPaths []string

	// Walk through the apis directory.
	err := filepath.Walk(apisPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Index on beta resources only for now
		if info.IsDir() && strings.HasSuffix(info.Name(), "v1beta1") && len(exampleAPIsDirPaths) < examples {
			exampleAPIsDirPaths = append(exampleAPIsDirPaths, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("walking apis directory: %w", err)
	}

	// Read up to "examples" files from each directory.
	for _, dirPath := range exampleAPIsDirPaths {
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
				exampleAPIsContents = append(exampleAPIsContents, string(fileContent))
			}
		}
	}

	return exampleAPIsContents, nil
}