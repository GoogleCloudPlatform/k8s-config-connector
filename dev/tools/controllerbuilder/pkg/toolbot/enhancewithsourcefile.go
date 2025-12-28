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
	"fmt"
	"os"
	"path/filepath"
)

// EnhanceWithSource is an enhancer that adds a source file from our repo to the data point.
type EnhanceWithSource struct {
	repoRoot string
}

// NewEnhanceWithSource creates a new enhancer that will attach a source file from our repo.
func NewEnhanceWithSource(repoRoot string) (*EnhanceWithSource, error) {
	x := &EnhanceWithSource{
		repoRoot: repoRoot,
	}
	return x, nil
}

var _ Enhancer = &EnhanceWithSource{}

// EnhanceDataPoint enhances the data point by adding the referenced source code.
func (x *EnhanceWithSource) EnhanceDataPoint(ctx context.Context, data *DataPoint) error {
	if src := data.Input["terraform.src"]; src != "" {
		p := filepath.Join(x.repoRoot, "third_party", src)
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", p, err)
		}
		data.SetInput("terraform.src.content", string(b))
	}

	return nil
}
