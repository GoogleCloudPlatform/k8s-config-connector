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
	"runtime/debug"
	"strings"
	"text/template"

	ccTemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"github.com/fatih/color"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
)

const (
	currRelPath             = "dev/tools/controllerbuilder"
	directControllerRelPath = "pkg/controller/direct"
)

var funcMap = template.FuncMap{
	"ToLower": strings.ToLower,
}

func RegisterController(service, kind string) error {
	// Read register file
	directControllerPkgPath, err := buildDirectControllerPath()
	if err != nil {
		return nil
	}
	registerFilePath := filepath.Join(directControllerPkgPath, "register", "register.go")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, registerFilePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// Get main model name
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return fmt.Errorf("could not read build info")
	}
	modelPath := strings.TrimSuffix(bi.Main.Path, currRelPath)

	importPath := filepath.Join(modelPath, directControllerRelPath, service)
	added := astutil.AddNamedImport(fset, f, "_", importPath)
	if !added {
		fmt.Printf("skip registering controller %s\n", service)
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
	color.HiGreen("New controller %s has been registered.\n", kind)
	return nil
}

func GenerateController(service, kind string, cArgs *ccTemplate.ControllerArgs) error {
	tmpl, err := template.New(cArgs.Kind).Funcs(funcMap).Parse(ccTemplate.ControllerTemplate)
	if err != nil {
		return fmt.Errorf("parse controller template: %w", err)
	}
	// Apply the `service` and `resource` args to the controller and external resource templates
	controllerOutput := &bytes.Buffer{}
	if err := tmpl.Execute(controllerOutput, cArgs); err != nil {
		return err
	}

	controllerFilePath, err := buildControllerPath(service, cArgs.ProtoResource)
	if err != nil {
		return err
	}
	if _, err := os.Stat(controllerFilePath); err == nil {
		fmt.Printf("file %s already exists, skipping\n", controllerFilePath)
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("unexpected controller file: %w", err)
	}

	// Write the generated controller.go to  pkg/controller/direct/<service>/<resource>_controller.go
	if err := WriteToFile(controllerFilePath, controllerOutput.Bytes()); err != nil {
		return err
	}
	// Format and adjust the go imports in the generated controller file.
	if err := FormatImports(controllerFilePath, controllerOutput.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New controller %s has been generated.", kind)
	return nil
}

func buildDirectControllerPath() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get current working directory: %w", err)
	}
	abs, err := filepath.Abs(pwd)
	if err != nil {
		return "", fmt.Errorf("get absolute path %s: %w", pwd, err)
	}
	seg := strings.Split(abs, currRelPath)
	return filepath.Join(seg[0], directControllerRelPath), nil
}

func buildControllerPath(service, protoResource string) (string, error) {
	filename := strings.ToLower(protoResource) + "_controller.go"
	directControllerPkgPath, err := buildDirectControllerPath()
	if err != nil {
		return "", nil
	}
	controllerDir := filepath.Join(directControllerPkgPath, service)
	err = os.MkdirAll(controllerDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create controller directory %s: %w", controllerDir, err)
	}
	return filepath.Join(controllerDir, filename), nil
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
