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

package toolbot

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/klog"
)

var _ Enhancer = &EnhanceWithMappers{}

// EnhanceWithMappers is an enhancer that finds Go mapper functions in a directory
type EnhanceWithMappers struct {
	srcDirectory string
}

// NewEnhanceWithMappers creates a new EnhanceWithMappers
func NewEnhanceWithMappers(srcDirectory string) (*EnhanceWithMappers, error) {
	return &EnhanceWithMappers{
		srcDirectory: srcDirectory,
	}, nil
}

// EnhanceDataPoint enhances the data point by adding mapper function definitions
func (x *EnhanceWithMappers) EnhanceDataPoint(ctx context.Context, p *DataPoint) error {
	if p.Type != "fuzz-gen" { // Only enhance if this is a fuzz-gen tool
		return nil
	}

	apiGroup := p.Input["api.group"]
	if apiGroup == "" {
		return nil
	}

	// Extract resource type from API group (e.g., "bigquerydatatransfer" from "bigquerydatatransfer.cnrm.cloud.google.com")
	resourceType := strings.Split(apiGroup, ".")[0]

	// Find mapper files in the resource-specific directory
	mapperDir := filepath.Join(x.srcDirectory, resourceType)
	files, err := os.ReadDir(mapperDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("api.group %q might be incorrect", apiGroup)
		} else {
			return fmt.Errorf("reading mapper directory: %w", err)
		}
	}

	var mappers []string
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") {
			continue
		}

		// Look for mapper files. .e.g xx_mappings.go, xx_mapper.go, mapper.go, mappers.go, etc.
		if !strings.Contains(file.Name(), "mapp") && !strings.Contains(file.Name(), "mapper.generated.go") {
			continue
		}

		content, err := os.ReadFile(filepath.Join(mapperDir, file.Name()))
		if err != nil {
			return fmt.Errorf("reading mapper file: %w", err)
		}

		// Skip license header, package declaration, and imports
		lines := strings.Split(string(content), "\n")
		start := 0
		for i, line := range lines {
			if strings.HasPrefix(strings.TrimSpace(line), "func ") {
				start = i
				break
			}
		}

		mappers = append(mappers, strings.Join(lines[start:], "\n"))
	}

	if len(mappers) > 0 {
		p.SetInput("go.mappers", strings.Join(mappers, "\n\n"))
	} else {
		klog.Infof("no mapper functions found for resource type %q", resourceType)
	}

	return nil
}
