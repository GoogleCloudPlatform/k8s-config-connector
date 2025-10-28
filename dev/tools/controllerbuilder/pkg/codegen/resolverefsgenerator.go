// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/gocode"
	"k8s.io/klog/v2"
)

type ResolveRefsGenerator struct {
	generatorBase
	goPackages []*gocode.Package
}

func NewResolveRefsGenerator(outputBaseDir string) *ResolveRefsGenerator {
	g := &ResolveRefsGenerator{}
	g.generatorBase.init(outputBaseDir)
	return g
}

func (g *ResolveRefsGenerator) VisitGoCode(goPackage string, basePath string) error {
	packages, err := gocode.LoadPackageTree(goPackage, basePath)
	if err != nil {
		return fmt.Errorf("inspecting go code: %w", err)
	}

	for _, pkg := range packages {
		g.goPackages = append(g.goPackages, pkg)
	}

	return nil
}

func (g *ResolveRefsGenerator) GenerateResolveRefs(service, version string) error {
	for _, pkg := range g.goPackages {
		for _, s := range pkg.Structs {
			var kind string
			var specStruct *gocode.GoStruct

			for _, f := range s.Fields {
				if f.Name == "Spec" {
					specTypeName := strings.TrimPrefix(f.Type, "*")
					for _, s2 := range pkg.Structs {
						if s2.Name == specTypeName {
							kind = s.Name
							specStruct = s2
							break
						}
					}
				}
				if kind != "" {
					break
				}
			}

			if kind == "" || specStruct == nil {
				continue
			}

			// todo: Distinguish between parent references and other references.
			//       Parent refs should be excluded from resolution because they are already resolved when initializing
			//       the resource identity.
			refFields := g.findRefFields(pkg, specStruct)

			if refFields == nil {
				klog.Infof("no reference fields found for kind %s in package %q", kind, pkg.GoPackage)
				continue
			}

			var b bytes.Buffer
			p := func(format string, args ...interface{}) {
				fmt.Fprintf(&b, format, args...)
				fmt.Fprintln(&b)
			}

			p("func Resolve%sRefs(ctx context.Context, kube client.Reader, obj *krm.%s) error {", kind, kind)
			// todo: Leave a reminder here to ensure all references are resolved.
			//       Once all references are updated to implement the Normalize function, we can simplify this by
			//       generating the resolver code directly.
			for _, fieldName := range refFields {
				p("\t// TODO: resolve %s", fieldName)
			}
			p("\treturn nil")
			p("}")

			formatted, err := format.Source(b.Bytes())
			if err != nil {
				return fmt.Errorf("formatting generated code for %s: %w\n%s", kind, err, b.String())
			}

			outputServiceDir := filepath.Join(g.outputBaseDir, service)

			// Create file name, e.g. forwardingrule_resolverefs.go
			kindPrefix := strings.ToLower(kind)
			shortKind := strings.TrimPrefix(kindPrefix, service)
			outputFileName := strings.ToLower(shortKind) + "_resolverefs.go"
			outputFilePath := filepath.Join(outputServiceDir, outputFileName)

			if _, err := os.Stat(outputFilePath); err == nil {
				klog.Infof("file %s already exists, skipping", outputFilePath)
				continue
			}

			k := generatedFileKey{
				GoPackage: service,
				FileName:  outputFileName,
			}
			out := g.getOutputFile(k)
			out.goPackage = service
			out.addImport("", "context")
			out.addImport("krm", fmt.Sprintf("github.com/GoogleCloudPlatform/k8s-config-connector/apis/%s/%s", service, version))
			out.addImport("", "sigs.k8s.io/controller-runtime/pkg/client")
			out.body.Write(formatted)
		}
	}
	return nil
}

func (g *ResolveRefsGenerator) findRefFields(pkg *gocode.Package, specStruct *gocode.GoStruct) []string {
	visited := make(map[*gocode.GoStruct]bool)
	return g.findNestedRefFields(pkg, specStruct, "", visited)
}

func (g *ResolveRefsGenerator) findNestedRefFields(pkg *gocode.Package, goStruct *gocode.GoStruct, prefix string, visited map[*gocode.GoStruct]bool) []string {
	if visited[goStruct] {
		return nil
	}
	visited[goStruct] = true

	var refFields []string

	for _, field := range goStruct.Fields {
		fieldName := field.Name

		currentPath := fieldName
		if prefix != "" {
			currentPath = prefix + "." + fieldName
		}

		if strings.HasSuffix(fieldName, "Ref") || strings.HasSuffix(fieldName, "Refs") {
			refFields = append(refFields, currentPath)
		}

		// Recurse into nested structs
		typeName := strings.TrimPrefix(field.Type, "*")
		typeName = strings.TrimPrefix(typeName, "[]")
		typeName = strings.TrimPrefix(typeName, "*") // for slice of pointers

		var nextStruct *gocode.GoStruct
		for _, s := range pkg.Structs {
			if s.Name == typeName {
				nextStruct = s
				break
			}
		}

		if nextStruct != nil {
			nested := g.findNestedRefFields(pkg, nextStruct, currentPath, visited)
			refFields = append(refFields, nested...)
		}
	}

	return refFields
}
