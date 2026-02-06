// Copyright 2022 Google LLC
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

// This program parses CRDs found in a given YAML file and outputs them onto
// individual CRD files.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/yaml"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	kccyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/yaml"
)

const fileMode = 0600

func main() {
	var file string
	var outputDir string
	flag.StringVar(&file, "file", "", "YAML file to parse")
	flag.StringVar(&outputDir, "output-dir", "", "Directory where CRD files are to be written to")
	flag.Parse()
	if file == "" || outputDir == "" {
		fmt.Println("error: incorrect usage. Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if err := run(file, outputDir); err != nil {
		log.Fatal(err)
	}
}

func run(file, outputDir string) error {
	filePath, err := filepath.Abs(file)
	if err != nil {
		return fmt.Errorf("error converting file %v to an absolute path: %w", file, err)
	}
	b, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %v: %w", filePath, err)
	}
	yamls, err := kccyaml.SplitYAML(b)
	if err != nil {
		return fmt.Errorf("error splitting bytes into YAMLs: %w", err)
	}
	crds := make([]*apiextensions.CustomResourceDefinition, 0)
	for _, y := range yamls {
		var crd apiextensions.CustomResourceDefinition
		if err := yaml.Unmarshal(y, &crd); err != nil {
			return fmt.Errorf("error unmarshalling bytes into CRD: %w", err)
		}
		crds = append(crds, &crd)
	}
	for _, crd := range crds {
		b, err := yaml.Marshal(crd)
		if err != nil {
			return fmt.Errorf("error marshalling CRD into bytes: %w", err)
		}
		// sigs.k8s.io/yaml (and the underlying go-yaml) omits preserveUnknownFields: false
		// because of omitempty and it being the default value.
		// For v1 CRDs, we want to match kustomize output which includes it.
		if crd.APIVersion == "apiextensions.k8s.io/v1" {
			// A simple but effective way to ensure it's there if it was in the original struct
			// (which it is for our KCC CRDs).
			if !crd.Spec.PreserveUnknownFields {
				// Insert it into the marshaled YAML if it's missing.
				// We look for 'spec:' and insert it after.
				// This is a bit hacky but safer than using map[string]interface{} which might reorder everything.
				// Actually, since we use sigs.k8s.io/yaml which is based on json, we can just use a map
				// for marshaling if we want full control, but that's overkill.
				
				// Re-marshal into map to preserve fields that the struct might have missed
				// but that's also tricky.
				
				// Let's use a simpler approach: if it's missing in the string, add it.
				yamlStr := string(b)
				if !containsPreserveUnknownFields(yamlStr) {
					// Add it after 'spec:'
					// We use a regex or just simple replacement
					b = []byte(addPreserveUnknownFields(yamlStr))
				}
			}
		}
		outputName, err := crdgeneration.FileNameForCRD(crd)
		if err != nil {
			return fmt.Errorf("error determining file name for CRD with GVK %v: %w", crd.GroupVersionKind(), err)
		}
		outputPath := filepath.Join(outputDir, outputName)
		if err := os.WriteFile(outputPath, b, fileMode); err != nil {
			return fmt.Errorf("error writing file %v: %w", outputPath, err)
		}
	}
	return nil
}

func containsPreserveUnknownFields(yamlStr string) bool {
	return strings.Contains(yamlStr, "preserveUnknownFields:")
}

func addPreserveUnknownFields(yamlStr string) string {
	lines := strings.Split(yamlStr, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "spec:" {
			// Find the next line's indentation
			indent := ""
			if i+1 < len(lines) {
				for _, char := range lines[i+1] {
					if char == ' ' {
						indent += " "
					} else {
						break
					}
				}
			}
			if indent == "" {
				indent = "  "
			}
			// Insert preserveUnknownFields: false
			lines = append(lines[:i+1], append([]string{indent + "preserveUnknownFields: false"}, lines[i+1:]...)...)
			break
		}
	}
	return strings.Join(lines, "\n")
}
