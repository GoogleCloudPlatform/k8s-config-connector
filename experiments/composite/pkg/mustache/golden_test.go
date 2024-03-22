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

package mustache

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/composite/pkg/testutils"
	"sigs.k8s.io/yaml"
)

func TestWithGoldenTests(t *testing.T) {
	goldenDir := "testdata"
	files, err := os.ReadDir(goldenDir)
	if err != nil {
		t.Fatalf("reading directory %q: %v", goldenDir, err)
	}
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		dir := filepath.Join(goldenDir, f.Name())
		t.Run(f.Name(), func(t *testing.T) {
			ctx := context.TODO()

			b := testutils.MustReadFile(t, filepath.Join(dir, "input.yaml"))

			obj := map[string]interface{}{
				"metadata": map[string]interface{}{
					"name": "foo",
				},
				"spec": map[string]interface{}{
					"machineType": "justright",
				},
			}
			activation := &Activation{
				Object:  obj,
				Version: 1,
				Objects: make(map[string]interface{}),
			}

			activation.Objects["nodepool"] = map[string]interface{}{
				"metadata": map[string]interface{}{
					"name": "nodepool-1",
				},
				"spec": map[string]interface{}{},
				"status": map[string]interface{}{
					"instanceGroupUrls": []string{
						"url1",
					},
				},
			}

			objects, err := Interpret(ctx, b, activation)
			if err != nil {
				t.Fatalf("error from Interpret: %v", err)
			}

			var got bytes.Buffer
			for i, obj := range objects {
				if i != 0 {
					got.WriteString("\n---\n")
				}
				y, err := yaml.Marshal(obj)
				if err != nil {
					t.Fatalf("error from yaml.Marshal: %v", err)
				}
				got.Write(y)
			}

			testutils.CompareGoldenFile(t, filepath.Join(dir, "expected.yaml"), got.String())
		})
	}

}
