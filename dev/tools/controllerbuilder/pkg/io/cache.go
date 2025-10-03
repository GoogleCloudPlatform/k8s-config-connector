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
	"io"
	"net/http"
	"os"
)

// WriteToCache write `out` to a temporary file under `dir`.
func WriteToCache(ctx context.Context, dest string, out string, namePattern string) (string, error) {
	if dest == "" {
		return "", fmt.Errorf("dir is required")
	}

	// if dest is File, then write to file
	fi, err := os.Stat(dest)
	if err != nil {
		return "", err
	}
	if fi.Mode().IsRegular() {
		return dest, os.WriteFile(dest, []byte(out), 0644)
	}

	// If dest is dir, then write to dir
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.Mkdir(dest, 0755); err != nil {
			return "", fmt.Errorf("create dir %s: %w", dest, err)
		}
	}
	if namePattern == "" {
		namePattern = "unnamed-*"
	}
	tmpFile, err := os.CreateTemp(dest, namePattern)
	if err != nil {
		return "", fmt.Errorf("create file under %s: %w", dest, err)
	}
	defer tmpFile.Close()

	fmt.Printf("out %s\n", out)
	if _, err := tmpFile.WriteString(out); err != nil {
		return "", fmt.Errorf("write to file %s: %w", tmpFile.Name(), err)
	}

	return tmpFile.Name(), nil
}

func GetFromRemote(ctx context.Context, rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	defer resp.Body.Close()
	return string(body), nil
}
