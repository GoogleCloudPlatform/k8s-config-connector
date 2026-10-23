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

package scaffold

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	ccTemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"github.com/fatih/color"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
	"k8s.io/klog/v2"
)

const (
	currRelPath             = "dev/tools/controllerbuilder"
	directControllerRelPath = "pkg/controller/direct"

	// GeneratedControllerMarker marks controller files produced by generate-controller --prepopulate-resource.
	GeneratedControllerMarker = "// +generated:controller"
)

var funcMap = template.FuncMap{
	"ToLower": strings.ToLower,
	"Plural":  codegen.Plural,
}

func NewControllerBuilder(rootPath, service, proto string) *ControllerBuilder {
	return &ControllerBuilder{
		rootPath: rootPath,
		service:  service,
		proto:    proto,
	}
}

type ControllerBuilder struct {
	rootPath            string
	service             string
	proto               string
	overwrite           bool
	prepopulateResource bool
}

func (c *ControllerBuilder) WithOverwrite(overwrite bool) *ControllerBuilder {
	c.overwrite = overwrite
	return c
}

func (c *ControllerBuilder) WithPrepopulateResource(prepopulate bool) *ControllerBuilder {
	c.prepopulateResource = prepopulate
	return c
}

func (c *ControllerBuilder) RegisterController() error {
	// Read register file
	registerFilePath := filepath.Join(c.getDirectPath(), "register", "register.go")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, registerFilePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	importPath := filepath.Join("github.com/GoogleCloudPlatform/k8s-config-connector", directControllerRelPath, c.service)
	added := astutil.AddNamedImport(fset, f, "_", importPath)
	if !added {
		fmt.Printf("skip registering controller %s\n", c.service)
		return nil
	}

	out := &bytes.Buffer{}
	err = format.Node(out, fset, f)
	if err != nil {
		return fmt.Errorf("error formatting code: %w", err)
	}

	if err := FormatImports(registerFilePath, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New controller %s has been registered.\n", c.proto)
	return nil
}

// RenderController executes the GreenfieldControllerTemplate with cArgs and formats
// the resulting Go source code and imports.
func RenderController(cArgs *ccTemplate.ControllerArgs) ([]byte, error) {
	cArgs.EnsureDefaults()
	tmpl, err := template.New(cArgs.Kind).Funcs(funcMap).Parse(ccTemplate.GreenfieldControllerTemplate)
	if err != nil {
		return nil, fmt.Errorf("parse controller template: %w", err)
	}
	controllerOutput := &bytes.Buffer{}
	if err := tmpl.Execute(controllerOutput, cArgs); err != nil {
		return nil, err
	}
	importOps := &imports.Options{
		Comments:  true,
		AllErrors: true,
		Fragment:  true,
	}
	formattedOut, err := imports.Process(strings.ToLower(cArgs.Kind)+"_controller.go", controllerOutput.Bytes(), importOps)
	if err != nil {
		return nil, fmt.Errorf("format rendered controller for %s: %w", cArgs.Kind, err)
	}
	return formattedOut, nil
}

func (c *ControllerBuilder) GenerateController(cArgs *ccTemplate.ControllerArgs) error {
	usePrepopulate := c.prepopulateResource || cArgs.PassesApplied
	tmplText := ccTemplate.ControllerTemplate
	if usePrepopulate {
		cArgs.EnsureDefaults()
		tmplText = ccTemplate.GreenfieldControllerTemplate
	}
	tmpl, err := template.New(cArgs.Kind).Funcs(funcMap).Parse(tmplText)
	if err != nil {
		return fmt.Errorf("parse controller template: %w", err)
	}
	// Apply the `service` and `resource` args to the controller and external resource templates
	controllerOutput := &bytes.Buffer{}
	if err := tmpl.Execute(controllerOutput, cArgs); err != nil {
		return err
	}

	controllerFilePath, err := c.getControllerPath()
	if err != nil {
		return err
	}
	if usePrepopulate {
		shouldSkip, err := c.shouldSkipControllerWrite(controllerFilePath, cArgs)
		if err != nil {
			return err
		}
		if shouldSkip {
			return nil
		}
	} else if !c.overwrite {
		if _, err := os.Stat(controllerFilePath); err == nil {
			klog.V(1).Infof("file %s already exists, skipping\n", controllerFilePath)
			return nil
		} else if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("unexpected controller file: %w", err)
		}
	}

	// Write the generated controller.go to pkg/controller/direct/<service>/<resource>_controller.go
	if err := WriteToFile(controllerFilePath, controllerOutput.Bytes()); err != nil {
		return err
	}
	// Format and adjust the go imports in the generated controller file.
	if err := FormatImports(controllerFilePath, controllerOutput.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New controller %s has been generated.", c.proto)
	return nil
}

func (c *ControllerBuilder) shouldSkipControllerWrite(controllerFilePath string, cArgs *ccTemplate.ControllerArgs) (bool, error) {
	if c.overwrite {
		return false, nil
	}
	existing, err := os.ReadFile(controllerFilePath)
	if err == nil {
		if (c.prepopulateResource || cArgs.PassesApplied) && bytes.Contains(existing, []byte(GeneratedControllerMarker)) {
			return false, nil
		}
		klog.V(1).Infof("file %s already exists (handwritten), skipping\n", controllerFilePath)
		return true, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return false, fmt.Errorf("unexpected controller file: %w", err)
	}
	if hasExistingHandwrittenController(filepath.Dir(controllerFilePath), cArgs.Kind) {
		klog.V(1).Infof("handwritten controller for %s already exists in %s, skipping\n", cArgs.Kind, filepath.Dir(controllerFilePath))
		return true, nil
	}
	return false, nil
}

func hasExistingHandwrittenController(serviceDir, kind string) bool {
	if kind == "" {
		return false
	}
	gvkToken := "." + kind + "GVK"
	found := false
	_ = filepath.WalkDir(serviceDir, func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(d.Name(), ".go") || strings.HasSuffix(d.Name(), "_test.go") {
			return nil
		}
		src, err := os.ReadFile(p)
		if err != nil {
			return nil
		}
		if bytes.Contains(src, []byte(GeneratedControllerMarker)) {
			return nil
		}
		if bytes.Contains(src, []byte("RegisterModel(")) && bytes.Contains(src, []byte(gvkToken)) {
			found = true
			return filepath.SkipAll
		}
		return nil
	})
	return found
}

func (c *ControllerBuilder) getControllerPath() (string, error) {
	filename := strings.ToLower(c.proto) + "_controller.go"
	direct := c.getDirectPath()
	controllerDir := filepath.Join(direct, c.service)
	err := os.MkdirAll(controllerDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create controller directory %s: %w", controllerDir, err)
	}
	return filepath.Join(controllerDir, filename), nil
}

func (c *ControllerBuilder) getDirectPath() string {
	seg := strings.Split(c.rootPath, currRelPath)
	return filepath.Join(seg[0], directControllerRelPath)
}

func FormatImports(path string, out []byte) error {
	importOps := &imports.Options{
		Comments:  true,
		AllErrors: true,
		Fragment:  true}
	formattedOut, err := imports.Process(path, out, importOps)
	if err != nil {
		return fmt.Errorf("format controller file %s: %w", path, err)
	}
	return WriteToFile(path, formattedOut)
}

func WriteToFile(path string, out []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", filepath.Dir(path), err)
	}
	// Use O_TRUNC to truncate the file
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(out)
	if err != nil {
		return fmt.Errorf("write file %s: %w", path, err)
	}
	return nil
}
