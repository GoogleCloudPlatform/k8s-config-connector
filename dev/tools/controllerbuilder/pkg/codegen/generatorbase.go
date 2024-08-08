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
	"sort"
	"strings"
	"time"

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
	baseDir  string
	key      generatedFileKey
	contents bytes.Buffer
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

func (f *generatedFile) Write(addCopyright bool) error {
	if f.contents.Len() == 0 {
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

	if err := f.maybeImportRefs(); err != nil {
		return err
	}

	f.contents.WriteTo(&w)

	p := filepath.Join(dir, f.key.FileName)
	klog.Infof("writing file %v", p)
	if err := os.WriteFile(p, w.Bytes(), 0644); err != nil {
		return fmt.Errorf("writing %q: %w", p, err)
	}

	return nil
}

func (v *generatorBase) WriteFiles(addCopyright bool) error {
	for _, f := range v.generatedFiles {
		if err := f.Write(addCopyright); err != nil {
			return err
		}
	}
	return nil
}

func (g *TypeGenerator) findTypeDeclaration(goTypeName string, srcDir string, skipGenerated bool) (*string, error) {
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

func (g *TypeGenerator) findTypeDeclarationWithProtoTag(protoTag string, srcDir string, skipGenerated bool) (*string, error) {
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
			line = strings.TrimSpace(line)
			line = strings.TrimPrefix(line, "//")
			line = strings.TrimSpace(line)
			line += " "
			if strings.HasPrefix(line, "+kcc:proto="+protoTag+" ") {
				return &line, nil
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

func (f *generatedFile) maybeImportRefs() error {
	if bytes.Contains(f.contents.Bytes(), []byte("refs.")) {
		return addImport(&f.contents, "refs", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1")
	}
	return nil
}

func addImport(buffer *bytes.Buffer, alias, importPackage string) error {
	content := buffer.String()

	importStr := fmt.Sprintf(`"%s"`, importPackage)
	if alias != "" {
		importStr = fmt.Sprintf(`%s "%s"`, alias, importPackage)
	}

	if strings.Contains(content, importPackage) { // import already exists
		return nil
	}

	startIndex := strings.Index(content, "import (")
	if startIndex == -1 {
		// No import block found, insert a new import block
		// Find the index after the package statement
		packageIndex := strings.Index(content, "package ")
		packageEndIndex := strings.Index(content[packageIndex:], "\n") + packageIndex + 1
		newImportBlock := fmt.Sprintf("\nimport (\n\t%s\n)\n", importStr)
		content = content[:packageEndIndex] + newImportBlock + content[packageEndIndex:]
	} else {
		// Insert the new import into the existing block
		insertIndex := strings.Index(content[startIndex:], ")") + startIndex
		newImport := fmt.Sprintf("\t%s\n", importStr)
		content = content[:insertIndex] + newImport + content[insertIndex:]
	}

	// Ensure imports are ordered correctly
	content = ensureImportOrder(content)

	buffer.Reset()
	buffer.WriteString(content)

	return nil
}

func ensureImportOrder(content string) string {
	lines := strings.Split(content, "\n")
	importStart := -1
	importEnd := -1
	var imports []string
	for i, line := range lines {
		if strings.Contains(line, "import (") {
			importStart = i
		}
		if importStart != -1 && line == ")" {
			importEnd = i
			break
		}
		if importStart != -1 && importStart != i && strings.Contains(line, "\"") {
			imports = append(imports, line)
		}
	}

	if importStart == -1 || importEnd == -1 {
		return content
	}

	// Sort imports
	sort.Slice(imports, func(i, j int) bool {
		importI := extractImportPackage(imports[i])
		importJ := extractImportPackage(imports[j])
		return importI < importJ
	})

	newImportBlock := "import (\n"
	for _, imp := range imports {
		newImportBlock += imp + "\n" // imp already includes the leading tab
	}
	newImportBlock += ")"

	// Rebuild the entire content
	newContent := strings.Join(lines[:importStart], "\n") + "\n" + newImportBlock + "\n" + strings.Join(lines[importEnd+1:], "\n")

	return newContent
}

// extractImportPackage removes leading tabs and alias of an imported package
func extractImportPackage(importLine string) string {
	p := strings.TrimSpace(importLine)
	startQuoteIndex := strings.Index(p, "\"")
	p = p[startQuoteIndex:]
	return p
}
