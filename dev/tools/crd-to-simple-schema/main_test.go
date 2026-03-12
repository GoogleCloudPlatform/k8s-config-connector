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
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

func TestRun(t *testing.T) {
	dirs, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatalf("reading testdata: %v", err)
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		t.Run(d.Name(), func(t *testing.T) {
			dir := filepath.Join("testdata", d.Name())

			crd1 := filepath.Join(dir, "crd1.yaml")
			crd2 := filepath.Join(dir, "crd2.yaml")
			outputFile := filepath.Join(dir, "_output.txt")

			var opt ConvertOptions
			opt.CRDFile = crd1
			opt.DiffCRDFile = crd2
			opt.Flatten = true

			if _, err := os.Stat(filepath.Join(dir, "ignore_int_diff")); err == nil {
				opt.IgnoreIntegerTypeDifferences = true
			}

			var buf bytes.Buffer
			if err := Run(t.Context(), opt, &buf); err != nil {
				t.Fatalf("Run failed: %v", err)
			}

			test.CompareGoldenFile(t, outputFile, buf.String())
		})
	}
}
