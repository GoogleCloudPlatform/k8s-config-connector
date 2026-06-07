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
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/openapi-to-krm/pkg/shared"
)

type DiscoveryDoc struct {
	Schemas map[string]*Schema `json:"schemas"`
}

type Schema struct {
	ID                   string             `json:"id"`
	Description          string             `json:"description"`
	Type                 string             `json:"type"`
	Format               string             `json:"format"`
	Properties           map[string]*Schema `json:"properties"`
	Items                *Schema            `json:"items"`
	Ref                  string             `json:"$ref"`
	AdditionalProperties *Schema            `json:"additionalProperties"`
}

type resourceList []string

func (r *resourceList) String() string {
	return strings.Join(*r, ",")
}

func (r *resourceList) Set(value string) error {
	*r = append(*r, value)
	return nil
}

type Options struct {
	SchemaFile    string
	APIVersion    string
	Resource      []string
	OutputFile    string
	IgnoreFields  []string
	RequireFields []string
}

func main() {
	var opt Options
	var ignoreFieldsStr string
	var requireFieldsStr string
	var rList resourceList

	flag.StringVar(&opt.SchemaFile, "schema-file", "", "Path to the Discovery API / OpenAPI JSON file")
	flag.StringVar(&opt.APIVersion, "api-version", "", "KRM API Version (e.g., dns.cnrm.cloud.google.com/v1beta1)")
	flag.Var(&rList, "resource", "Resource mapping (e.g., DNSManagedZone:ManagedZone)")
	flag.StringVar(&opt.OutputFile, "output-file", "", "Output file path")
	flag.StringVar(&ignoreFieldsStr, "ignore-field", "", "Comma-separated or repeated field to ignore (e.g., *:kind,ManagedZone:labels)")
	flag.StringVar(&requireFieldsStr, "require-field", "", "Comma-separated or repeated field to mark as required (e.g., *:dnsName,ManagedZone:dnsName)")
	flag.Parse()

	opt.Resource = rList

	if opt.SchemaFile == "" || opt.APIVersion == "" || len(opt.Resource) == 0 || opt.OutputFile == "" {
		log.Fatalf("Flags (-schema-file, -api-version, -resource, -output-file) are required")
	}

	if ignoreFieldsStr != "" {
		opt.IgnoreFields = strings.Split(ignoreFieldsStr, ",")
	}

	if requireFieldsStr != "" {
		opt.RequireFields = strings.Split(requireFieldsStr, ",")
	}

	ctx := context.Background()
	if err := Run(ctx, opt); err != nil {
		log.Fatalf("Execution failed: %v", err)
	}
}

func Run(ctx context.Context, opt Options) error {
	// Read and parse schema file
	content, err := os.ReadFile(opt.SchemaFile)
	if err != nil {
		return fmt.Errorf("reading schema file: %w", err)
	}

	var doc DiscoveryDoc
	if err := json.Unmarshal(content, &doc); err != nil {
		return fmt.Errorf("parsing JSON schema: %w", err)
	}

	visited := make(map[string]bool)
	var resourceHeaders []string

	for _, res := range opt.Resource {
		parts := strings.Split(res, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid resource format, expected Kind:OpenAPIType, got %s", res)
		}
		krmKind := parts[0]
		rootType := parts[1]
		resourceHeaders = append(resourceHeaders, fmt.Sprintf("%s:%s", krmKind, rootType))
		findReachable(doc.Schemas, rootType, visited)
	}

	outDir := filepath.Dir(opt.OutputFile)
	handCodedIDs, err := findHandCodedOpenAPIIDs(outDir)
	if err != nil {
		return fmt.Errorf("scanning handcoded openapi tags: %w", err)
	}

	var sortedSchemas []string
	for id := range visited {
		sortedSchemas = append(sortedSchemas, id)
	}
	sort.Strings(sortedSchemas)

	var buf strings.Builder
	buf.WriteString("// Copyright 2026 Google LLC\n")
	buf.WriteString("//\n")
	buf.WriteString("// Licensed under the Apache License, Version 2.0 (the \"License\");\n")
	buf.WriteString("// you may not use this file except in compliance with the License.\n")
	buf.WriteString("// You may obtain a copy of the License at\n")
	buf.WriteString("//\n")
	buf.WriteString("//      http://www.apache.org/licenses/LICENSE-2.0\n")
	buf.WriteString("//\n")
	buf.WriteString("// Unless required by applicable law or agreed to in writing, software\n")
	buf.WriteString("// distributed under the License is distributed on an \"AS IS\" BASIS,\n")
	buf.WriteString("// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n")
	buf.WriteString("// See the License for the specific language governing permissions and\n")
	buf.WriteString("// limitations under the License.\n\n")
	buf.WriteString("// Code generated by dev/tools/openapi-to-krm. DO NOT EDIT.\n")
	buf.WriteString("// +generated:types\n")
	fmt.Fprintf(&buf, "// krm.group: %s\n", strings.Split(opt.APIVersion, "/")[0])
	fmt.Fprintf(&buf, "// krm.version: %s\n", strings.Split(opt.APIVersion, "/")[1])
	for _, rh := range resourceHeaders {
		fmt.Fprintf(&buf, "// resource: %s\n", rh)
	}
	buf.WriteString("\n")
	fmt.Fprintf(&buf, "package %s\n\n", strings.Split(opt.APIVersion, "/")[1])

	for _, id := range sortedSchemas {
		s, ok := doc.Schemas[id]
		if !ok {
			continue
		}
		isHandCoded := handCodedIDs[id]
		if isHandCoded {
			fmt.Fprintf(&buf, "/* found existing non-generated go type with openapi tag %q, skipping\n\n", id)
		}

		fmt.Fprintf(&buf, "// +openapi:%s\n", id)
		fmt.Fprintf(&buf, "type %s struct {\n", id)

		// Print properties sorted
		var props []string
		for p := range s.Properties {
			props = append(props, p)
		}
		sort.Strings(props)

		for _, p := range props {
			if shouldIgnoreField(id, p, opt.IgnoreFields) {
				continue
			}
			prop := s.Properties[p]
			goName := shared.GoFieldName(p)
			gType := goType(prop)
			comment := commentBlock(prop.Description, "\t")

			jsonTag := fmt.Sprintf("`json:\"%s,omitempty\"`", p)
			fieldRequired := isFieldRequired(id, p, opt.RequireFields)
			if fieldRequired {
				if comment != "" {
					comment += "\t// +required\n"
				} else {
					comment = "\t// +required\n"
				}
				jsonTag = fmt.Sprintf("`json:\"%s\"`", p)
			}

			fmt.Fprintf(&buf, "%s\t%s %s %s\n\n", comment, goName, gType, jsonTag)
		}

		fmt.Fprintf(&buf, "}\n")
		if isHandCoded {
			buf.WriteString("*/\n\n")
		} else {
			buf.WriteString("\n")
		}
	}

	if err := os.WriteFile(opt.OutputFile, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("writing output file: %w", err)
	}
	fmt.Printf("Successfully generated %s from %s\n", opt.OutputFile, opt.SchemaFile)
	return nil
}

func shouldIgnoreField(typeName string, fieldName string, ignoreRules []string) bool {
	for _, rule := range ignoreRules {
		parts := strings.Split(rule, ":")
		if len(parts) != 2 {
			continue
		}
		tRule, fRule := parts[0], parts[1]
		if (tRule == "*" || tRule == typeName) && (fRule == "*" || fRule == fieldName) {
			return true
		}
	}
	return false
}

func isFieldRequired(typeName string, fieldName string, requiredRules []string) bool {
	for _, rule := range requiredRules {
		parts := strings.Split(rule, ":")
		if len(parts) != 2 {
			continue
		}
		tRule, fRule := parts[0], parts[1]
		if (tRule == "*" || tRule == typeName) && (fRule == "*" || fRule == fieldName) {
			return true
		}
	}
	return false
}

func findReachable(schemas map[string]*Schema, root string, visited map[string]bool) {
	if visited[root] {
		return
	}
	visited[root] = true
	s, ok := schemas[root]
	if !ok {
		return
	}
	for _, prop := range s.Properties {
		resolveRefs(schemas, prop, visited)
	}
}

func resolveRefs(schemas map[string]*Schema, prop *Schema, visited map[string]bool) {
	if prop.Ref != "" {
		findReachable(schemas, prop.Ref, visited)
	}
	if prop.Items != nil {
		resolveRefs(schemas, prop.Items, visited)
	}
	if prop.Properties != nil {
		for _, subProp := range prop.Properties {
			resolveRefs(schemas, subProp, visited)
		}
	}
}

// findHandCodedOpenAPIIDs uses the standard Go AST parser to locate Go struct declarations
// in a directory (excluding types.generated.go) and extracts any associated '+openapi:<id>' tag comments.
func findHandCodedOpenAPIIDs(dir string) (map[string]bool, error) {
	handCodedIDs := make(map[string]bool)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(fi os.FileInfo) bool {
		return !fi.IsDir() && !strings.HasSuffix(fi.Name(), "types.generated.go") && strings.HasSuffix(fi.Name(), ".go")
	}, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`\+openapi:([A-Za-z0-9_]+)`)

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				genDecl, ok := n.(*ast.GenDecl)
				if !ok || genDecl.Tok != token.TYPE {
					return true
				}
				if genDecl.Doc != nil {
					for _, comment := range genDecl.Doc.List {
						matches := re.FindStringSubmatch(comment.Text)
						if len(matches) > 1 {
							handCodedIDs[matches[1]] = true
						}
					}
				}
				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if typeSpec.Doc != nil {
						for _, comment := range typeSpec.Doc.List {
							matches := re.FindStringSubmatch(comment.Text)
							if len(matches) > 1 {
								handCodedIDs[matches[1]] = true
							}
						}
					}
					if typeSpec.Comment != nil {
						for _, comment := range typeSpec.Comment.List {
							matches := re.FindStringSubmatch(comment.Text)
							if len(matches) > 1 {
								handCodedIDs[matches[1]] = true
							}
						}
					}
				}
				return true
			})
		}
	}
	return handCodedIDs, nil
}

// splitWords is now imported from shared package
// goFieldName is now imported from shared package

func goType(prop *Schema) string {
	if prop.Ref != "" {
		return "*" + prop.Ref
	}
	switch prop.Type {
	case "string":
		if prop.Format == "int64" {
			return "*int64"
		}
		if prop.Format == "uint64" {
			return "*uint64"
		}
		return "*string"
	case "integer":
		if prop.Format == "int64" {
			return "*int64"
		}
		if prop.Format == "uint64" {
			return "*uint64"
		}
		return "*int64"
	case "boolean":
		return "*bool"
	case "number":
		return "*float64"
	case "array":
		itemType := "string"
		if prop.Items != nil {
			if prop.Items.Ref != "" {
				itemType = prop.Items.Ref
			} else {
				switch prop.Items.Type {
				case "string":
					itemType = "string"
				case "integer":
					itemType = "int64"
				case "boolean":
					itemType = "bool"
				case "number":
					itemType = "float64"
				}
			}
		}
		return "[]" + itemType
	case "object":
		if prop.AdditionalProperties != nil {
			valType := "string"
			if prop.AdditionalProperties.Type != "" {
				switch prop.AdditionalProperties.Type {
				case "string":
					valType = "string"
				case "integer":
					valType = "int64"
				}
			}
			return "map[string]" + valType
		}
		return "map[string]any"
	}
	return "any"
}

// commentBlock takes an OpenAPI/Discovery field description string, sanitizes it (replacing
// newlines and escaped block comments), and formats it as a neat comment block with wrap limits.
func commentBlock(desc string, indent string) string {
	if desc == "" {
		return ""
	}
	// Sanitize desc (such as escaped quotes or backslashes, and prevent premature comment termination)
	desc = strings.ReplaceAll(desc, "\n", " ")
	desc = strings.ReplaceAll(desc, "*/", "* /")
	words := strings.Fields(desc)
	var lines []string
	var currentLine []string
	currentLen := 0
	for _, word := range words {
		if currentLen+len(word)+1 > 80 {
			lines = append(lines, indent+"// "+strings.Join(currentLine, " "))
			currentLine = []string{word}
			currentLen = len(word)
		} else {
			currentLine = append(currentLine, word)
			currentLen += len(word) + 1
		}
	}
	if len(currentLine) > 0 {
		lines = append(lines, indent+"// "+strings.Join(currentLine, " "))
	}
	return strings.Join(lines, "\n") + "\n"
}
