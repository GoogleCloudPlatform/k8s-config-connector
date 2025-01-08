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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	kccyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/yaml"

	"github.com/ghodss/yaml" //nolint:depguard
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
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
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %v: %w", filePath, err)
	}
	yamls, err := kccyaml.SplitYAML(bytes)
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
		bytes, err := yaml.Marshal(crd)
		if err != nil {
			return fmt.Errorf("error marshalling CRD into bytes: %w", err)
		}
		outputName, err := crdgeneration.FileNameForCRD(crd)
		if err != nil {
			return fmt.Errorf("error determining file name for CRD with GVK %v: %w", crd.GroupVersionKind(), err)
		}
		outputPath := filepath.Join(outputDir, outputName)
		if err := ioutil.WriteFile(outputPath, bytes, fileMode); err != nil {
			return fmt.Errorf("error writing file %v: %w", outputName, err)
		}
	}
	return nil
}
