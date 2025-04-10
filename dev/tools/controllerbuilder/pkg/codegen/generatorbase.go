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

package codegen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"k8s.io/klog/v2"
)

type generatorBase struct {
	outputBaseDir  string
	generatedFiles map[generatedFileKey]*generatedFile
	errors         []error
}

func (g *generatorBase) init(outputBaseDir string) {
	g.outputBaseDir = outputBaseDir
	g.generatedFiles = make(map[generatedFileKey]*generatedFile)
}

func (g *generatorBase) getOutputFile(k generatedFileKey) *generatedFile {
	out := g.generatedFiles[k]
	if out == nil {
		out = &generatedFile{key: k, baseDir: g.outputBaseDir}
		g.generatedFiles[k] = out
	}
	return out
}

func (g *generatorBase) Errorf(msg string, args ...any) {
	g.errors = append(g.errors, fmt.Errorf(msg, args...))
}

type generatedFile struct {
	baseDir     string
	key         generatedFileKey
	packageName string
	body        bytes.Buffer

	fileAnnotation *annotations.FileAnnotation

	imports map[string]string
}

type generatedFileKey struct {
	GoPackage string

	FileName string
}

func (f *generatedFile) OutputDir() string {
	tokens := strings.Split(f.key.GoPackage, ".")
	dirTokens := []string{f.baseDir}
	dirTokens = append(dirTokens, tokens...)
	dir := filepath.Join(dirTokens...)
	return dir
}

func (f *generatedFile) addImport(alias string, pkgName string) {
	if f.imports == nil {
		f.imports = make(map[string]string)
	}
	f.imports[pkgName] = alias
}

func (f *generatedFile) Write(addCopyright bool) error {
	if f.body.Len() == 0 {
		return nil
	}

	dir := f.OutputDir()

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating directory %q: %w", dir, err)
	}

	var w bytes.Buffer

	if addCopyright {
		writeCopyright(&w, time.Now().Year())
	}

	if f.fileAnnotation != nil {
		s := f.fileAnnotation.FormatGo()
		fmt.Fprintf(&w, "%s\n", s)
	}

	if f.packageName != "" {
		fmt.Fprintf(&w, "package %s\n", f.packageName)
		fmt.Fprintf(&w, "\n")
	}

	if len(f.imports) != 0 {
		w.WriteString("import (\n")
		for pkgName, alias := range f.imports {
			if alias == "" {
				w.WriteString(fmt.Sprintf("\t%q\n", pkgName))
			} else {
				w.WriteString(fmt.Sprintf("\t%s %q\n", alias, pkgName))
			}
		}
		w.WriteString(")\n")
	}
	f.body.WriteTo(&w)

	p := filepath.Join(dir, f.key.FileName)
	klog.Infof("writing file %v", p)
	if err := os.WriteFile(p, w.Bytes(), 0644); err != nil {
		return fmt.Errorf("writing %q: %w", p, err)
	}

	return nil
}

func (g *generatorBase) WriteFiles(addCopyright bool) error {
	for _, f := range g.generatedFiles {
		if err := f.Write(addCopyright); err != nil {
			return err
		}
	}
	return nil
}

func (g *generatorBase) findTypeDeclaration(goTypeName string, srcDir string, skipGenerated bool) (*string, error) {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		if os.IsNotExist(err) { // type declaration does not exist
			return nil, nil
		}
		return nil, fmt.Errorf("reading directory %q: %w", srcDir, err)
	}

	for _, f := range files {
		p := filepath.Join(srcDir, f.Name())
		if !strings.HasSuffix(p, ".go") {
			continue
		}
		if skipGenerated && strings.HasSuffix(p, "generated.go") {
			continue
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("reading file %q: %w", p, err)
		}

		for _, line := range strings.Split(string(b), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "type "+goTypeName+" ") {
				return &line, nil
			}
		}
	}

	return nil, nil
}

func (g *generatorBase) findTypeDeclarationWithProtoTag(protoTag string, srcDir string, skipGenerated bool) (*string, error) {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %q: %w", srcDir, err)
	}

	for _, f := range files {
		p := filepath.Join(srcDir, f.Name())
		if !strings.HasSuffix(p, ".go") {
			continue
		}
		if skipGenerated && strings.HasSuffix(p, "generated.go") {
			continue
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("reading file %q: %w", p, err)
		}

		for _, line := range strings.Split(string(b), "\n") {
			if proto, ok := GetProtoMessageFromAnnotation(line); ok {
				if proto == protoTag {
					return &line, nil
				}
			}
		}
	}

	return nil, nil
}

func (g *generatorBase) findFuncDeclaration(goFuncName string, srcDir string, skipGenerated bool) *string {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		g.Errorf("reading directory %q: %w", srcDir, err)
		return nil
	}

	for _, f := range files {
		p := filepath.Join(srcDir, f.Name())
		if !strings.HasSuffix(p, ".go") {
			continue
		}
		if skipGenerated && strings.HasSuffix(p, "generated.go") {
			continue
		}
		b, err := os.ReadFile(p)
		if err != nil {
			g.Errorf("reading file %q: %w", p, err)
			return nil
		}

		for _, line := range strings.Split(string(b), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "func "+goFuncName+"(") {
				return &line
			}
		}
	}

	return nil
}
