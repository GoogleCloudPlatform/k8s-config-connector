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

package io

import (
	"context"
	"fmt"
	"os"
)

// WriteToCache write `out` to a temporary file under `dir`.
func WriteToCache(ctx context.Context, dir string, out string, namePattern string) (string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			return "", fmt.Errorf("create dir %s: %w", dir, err)
		}
	}
	if namePattern == "" {
		namePattern = "unnamed-*"
	}
	tmpFile, err := os.CreateTemp(dir, namePattern)
	if err != nil {
		return "", fmt.Errorf("create file under %s: %w", dir, err)
	}
	defer tmpFile.Close()

	fmt.Printf("out %s\n", out)
	if _, err := tmpFile.WriteString(out); err != nil {
		return "", fmt.Errorf("write to file %s: %w", tmpFile.Name(), err)
	}

	return tmpFile.Name(), nil
}
