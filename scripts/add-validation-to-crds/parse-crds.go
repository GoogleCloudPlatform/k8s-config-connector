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

// This program parses CRDs found in a given YAML file and outputs them onto
// individual CRD files.

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	kccyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/yaml"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

func main() {
	err := run(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	var dir string
	flag.StringVar(&dir, "dir", "", "Directory to process")
	flag.Parse()

	if dir == "" {
		return fmt.Errorf("--dir is required")
	}

	if err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(p) != ".yaml" {
			return nil
		}

		if err := processFile(ctx, p); err != nil {
			return fmt.Errorf("processing file %q: %w", p, err)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func processFile(ctx context.Context, p string) error {
	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", p, err)
	}

	var out bytes.Buffer

	// We preserve the yaml header (copyright, typically)
	for _, line := range bytes.Split(b, []byte("\n")) {
		if len(line) == 0 || line[0] == '#' {
			out.Write(line)
			out.WriteString("\n")
		} else {
			break
		}
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
		if err := addRefsToCRD(crd); err != nil {
			return err
		}
	}

	for i, crd := range crds {
		b, err := yaml.Marshal(crd)
		if err != nil {
			return fmt.Errorf("error marshalling CRD into bytes: %w", err)
		}
		out.Write(b)
		if i != 0 {
			out.WriteString("\n---\n")
		}
	}

	if err := os.WriteFile(p, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("error writing file %q: %w", p, err)
	}
	return nil
}

func addRefsToCRD(crd *apiextensions.CustomResourceDefinition) error {
	for _, v := range crd.Spec.Versions {
		if err := addRefsToProps("", v.Schema.OpenAPIV3Schema); err != nil {
			return err
		}
	}
	return nil
}

func addRefsToProps(fieldPath string, props *apiextensions.JSONSchemaProps) error {
	// Descend into arrays
	if props.Items != nil {
		if props.Items.Schema != nil {
			if err := addRefsToProps(fieldPath+"[]", props.Items.Schema); err != nil {
				return err
			}
		}
		for i := range props.Items.JSONSchemas {
			if err := addRefsToProps(fieldPath+"[]", &props.Items.JSONSchemas[i]); err != nil {
				return err
			}
		}
	}

	// Descend into objects
	for k := range props.Properties {
		v := props.Properties[k]
		if err := addRefsToProps(fieldPath+"."+k, &v); err != nil {
			return err
		}
		props.Properties[k] = v
	}

	if err := addValidationToRefs(fieldPath, props); err != nil {
		return err
	}
	return nil
}

const refRuleWithoutKind = `
oneOf:
- not:
    required:
    - external
  required:
  - name
- not:
    anyOf:
    - required:
      - name
    - required:
      - namespace
  required:
  - external
`

const refRuleWithKind = `
oneOf:
- not:
    required:
    - external
  required:
  - name
  - kind
- not:
    anyOf:
    - required:
      - name
    - required:
      - namespace
    - required:
      - kind
  required:
  - external
`

const refRuleWithOptionalKind = `
oneOf:
- not:
    required:
    - external
  required:
  - name
- not:
    anyOf:
    - required:
      - name
    - required:
      - namespace
  required:
  - external
`

func addValidationToRefs(fieldPath string, props *apiextensions.JSONSchemaProps) error {
	// Is this a ref?
	if props.Type != "object" {
		return nil
	}

	fields := sets.New[string]()
	for k := range props.Properties {
		fields.Insert(k)
	}
	signature := strings.Join(sets.List(fields), ",")

	var ruleYAML string
	if signature == "apiVersion,external,kind,name,namespace" {
		// hack for IAMPolicy.spec.resourceRef for backwards compat
		if fieldPath == ".spec.resourceRef" {
			ruleYAML = refRuleWithOptionalKind
		} else {
			ruleYAML = refRuleWithKind
		}
	} else if signature == "external,kind,name,namespace" {
		ruleYAML = refRuleWithKind
		// kind is optional for projectRef (and maybe in future other well-known ref types)
		// fieldPath is the best mechanism we have today (?)
		if fieldPath == ".spec.projectRef" {
			ruleYAML = refRuleWithOptionalKind
		}
	} else if signature == "external,name,namespace" {
		ruleYAML = refRuleWithoutKind
	} else {
		if strings.HasPrefix(signature, "external,") {
			klog.Warningf("unknown signature %q", signature)
		}
		if strings.HasPrefix(signature, "apiVersion,external,") {
			klog.Warningf("unknown signature %q", signature)
		}
		return nil
	}

	rule := &apiextensions.JSONSchemaProps{}
	if err := yaml.Unmarshal([]byte(ruleYAML), &rule); err != nil {
		return err
	}
	props.OneOf = rule.OneOf

	return nil
}
