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
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

func main() {
	ctx := context.Background()

	var opt ConvertOptions

	flag.StringVar(&opt.Version, "version", "", "CRD version to convert")
	flag.StringVar(&opt.CRDFile, "crd-file", "", "CRD file to convert; if empty, reads from stdin")
	flag.StringVar(&opt.DiffCRDFile, "diff-crd-file", "", "CRD file to compare against; if empty, no comparison is done")
	flag.BoolVar(&opt.Flatten, "flatten", false, "Flatten output to path=type lines (easier for diffing)")
	flag.BoolVar(&opt.IgnoreIntegerTypeDifferences, "ignore-integer-type-differences", opt.IgnoreIntegerTypeDifferences, "Treat int32 and int64 as equivalent to integer when diffing.")

	flag.Parse()

	if err := Run(ctx, opt, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type ConvertOptions struct {
	// CRD file to convert; if empty, reads from stdin
	CRDFile string

	// CRD file to compare against; if empty, no comparison is done
	DiffCRDFile string

	// CRD version to convert; if empty, uses the first version found
	Version string

	// Flatten output to path=type lines (easier for diffing)
	Flatten bool

	// IgnoreIntegerTypeDifferences treats int32 and int64 as equivalent to integer when diffing.
	IgnoreIntegerTypeDifferences bool
}

func Run(ctx context.Context, opt ConvertOptions, out io.Writer) error {

	var data []byte
	if opt.CRDFile == "" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("reading stdin: %w", err)
		}
		data = b
	} else {
		b, err := os.ReadFile(opt.CRDFile)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", opt.CRDFile, err)
		}
		data = b
	}

	schema, err := buildSimpleSchema(data, opt)
	if err != nil {
		return fmt.Errorf("building simple schema: %w", err)
	}

	if opt.DiffCRDFile != "" {
		b, err := os.ReadFile(opt.DiffCRDFile)
		if err != nil {
			return fmt.Errorf("reading diff file %q: %w", opt.DiffCRDFile, err)
		}

		schema2, err := buildSimpleSchema(b, opt)
		if err != nil {
			return fmt.Errorf("building simple schema: %w", err)
		}

		if !opt.Flatten {
			return fmt.Errorf("diffing requires --flatten")
		}

		lines1 := make(map[string]string)
		flatten("", schema, lines1)

		lines2 := make(map[string]string)
		flatten("", schema2, lines2)

		var diff []string

		for k, v1 := range lines1 {
			v2, ok := lines2[k]
			if !ok {
				diff = append(diff, fmt.Sprintf("- %s=%s", k, v1))
				continue
			}
			if v1 != v2 {
				// Special case: treat int32 and int64 as equivalent to integer
				if opt.IgnoreIntegerTypeDifferences {
					if (v1 == "int32" || v1 == "int64") && v2 == "integer" {
						continue
					}
					if v1 == "integer" && (v2 == "int32" || v2 == "int64") {
						continue
					}
				}
				diff = append(diff, fmt.Sprintf("- %s=%s", k, v1))
				diff = append(diff, fmt.Sprintf("+ %s=%s", k, v2))
			}
		}
		for k, v2 := range lines2 {
			_, ok := lines1[k]
			if !ok {
				diff = append(diff, fmt.Sprintf("+ %s=%s", k, v2))
			}
		}

		// Funky sort that tries to match up + and - lines
		sort.Slice(diff, func(i, j int) bool {
			s0 := diff[i]
			s1 := diff[j]
			if s0[1:] < s1[1:] {
				return true
			}
			if s0[1:] > s1[1:] {
				return false
			}
			if s0[0] == '-' {
				return false
			}
			return true
		})

		for _, line := range diff {
			fmt.Fprintln(out, line)
		}
		return nil
	}

	if opt.Flatten {
		lines := make(map[string]string)
		flatten("", schema, lines)

		for _, line := range lines {
			fmt.Fprintln(out, line)
		}
		return nil
	}

	res, err := yaml.Marshal(schema)
	if err != nil {
		return fmt.Errorf("marshaling simple schema: %w", err)
	}
	if _, err := out.Write(res); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	return nil
}

func buildSimpleSchema(data []byte, opt ConvertOptions) (any, error) {
	var crd apiextensionsv1.CustomResourceDefinition
	if err := yaml.Unmarshal(data, &crd); err != nil {
		return nil, fmt.Errorf("unmarshaling CRD: %w", err)
	}

	// Find the stored version or the first version
	var schema *apiextensionsv1.JSONSchemaProps
	if opt.Version != "" {
		for _, v := range crd.Spec.Versions {
			if v.Name == opt.Version {
				schema = v.Schema.OpenAPIV3Schema
				break
			}
		}
	} else {
		if len(crd.Spec.Versions) > 0 {
			schema = crd.Spec.Versions[0].Schema.OpenAPIV3Schema
		}
	}

	if schema == nil {
		if opt.Version == "" {
			return nil, fmt.Errorf("no schema found in CRD %q (no schema version specified)", opt.CRDFile)
		}
		return nil, fmt.Errorf("schema version %q not found in CRD %q", opt.Version, opt.CRDFile)
	}

	simple := walk(schema)

	return simple, nil
}

func flatten(path string, schema any, out map[string]string) {
	switch schema := schema.(type) {
	case map[string]any:
		for k, v := range schema {
			flatten(path+"."+k, v, out)
		}

	case []any:
		for _, v := range schema {
			flatten(fmt.Sprintf("%s[]", path), v, out)
		}
	case string:
		out[path] = schema
	default:
		klog.Fatalf("unhandled type %T", schema)
	}

}

func walk(s *apiextensionsv1.JSONSchemaProps) any {
	if s == nil {
		return "unknown"
	}

	// Handle x-kubernetes-preserve-unknown-fields
	if s.XPreserveUnknownFields != nil && *s.XPreserveUnknownFields {
		return "json"
	}

	// Properties takes precedence (it's an object structure)
	if len(s.Properties) > 0 {
		m := make(map[string]any)
		for k, v := range s.Properties {
			val := v
			m[k] = walk(&val)
		}
		return m
	}

	// Array
	if s.Type == "array" {
		if s.Items != nil && s.Items.Schema != nil {
			return []any{walk(s.Items.Schema)}
		}
		// Fallback
		return []any{"any"}
	}

	// Map (AdditionalProperties)
	if s.AdditionalProperties != nil && s.AdditionalProperties.Schema != nil {
		val := walk(s.AdditionalProperties.Schema)
		if str, ok := val.(string); ok {
			return fmt.Sprintf("map[string]%s", str)
		}
		return map[string]any{
			"KEY": val,
		}
	}

	// Primitives
	t := s.Type
	if t == "" {
		// Check if AnyOf/AllOf/OneOf exists
		if len(s.AnyOf) > 0 {
			// Just pick the first one for simplicity, or return a string "AnyOf[...]"
			return "anyOf"
		}
		if len(s.AllOf) > 0 {
			return "allOf"
		}
		if len(s.OneOf) > 0 {
			return "oneOf"
		}
		return "any"
	}

	if s.Format != "" {
		switch s.Format {
		case "int32", "int64":
			return s.Format
		default:
			return fmt.Sprintf("%s (%s)", t, s.Format)
		}
	}

	return t
}
