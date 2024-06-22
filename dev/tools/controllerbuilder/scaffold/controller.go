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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	ccTemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template"
	"github.com/fatih/color"
	"golang.org/x/tools/imports"
)

const (
	currRelPath             = "dev/tools/controllerbuilder"
	directControllerRelPath = "pkg/controller/direct"
)

func Scaffold(path string, cArgs *ccTemplate.ControllerArgs) error {
	tmpl, err := template.New(cArgs.Kind).Parse(ccTemplate.ControllerTemplate)
	if err != nil {
		return fmt.Errorf("parse controller template: %s", err)
	}
	// Apply the `service` and `resource` args to the controller template
	out := &bytes.Buffer{}
	if err := tmpl.Execute(out, cArgs); err != nil {
		return err
	}
	// Write the generated controller.go to  pkg/controller/direct/<service>/<resource>_controller.go
	if err := WriteToFile(path, out.Bytes()); err != nil {
		return err
	}
	// Format and adjust the go imports in the generated controller file.
	if err := FormatImports(path, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New controller %s\nEnjoy it!\n", path)
	return nil
}

func BuildControllerPath(service, kind string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get current working directory: %w", err)
	}
	abs, err := filepath.Abs(pwd)
	if err != nil {
		return "", fmt.Errorf("get absolute path %s: %w", pwd, err)
	}
	seg := strings.Split(abs, currRelPath)
	controllerDir := filepath.Join(seg[0], directControllerRelPath, service)
	err = os.MkdirAll(controllerDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create controller directory %s: %w", controllerDir, err)
	}
	controllerFilePath := filepath.Join(controllerDir, strings.ToLower(kind)+"_controller.go")
	if _, err := os.Stat(controllerFilePath); err == nil {
		return "", fmt.Errorf("controller file %s may already exist: %w", controllerFilePath, err)
	}
	return controllerFilePath, nil
}

func FormatImports(path string, out []byte) error {
	importOps := &imports.Options{
		Comments:  true,
		AllErrors: true,
		Fragment:  true}
	formatedOut, err := imports.Process(path, out, importOps)
	if err != nil {
		return fmt.Errorf("format controller file %s: %w", path, err)
	}
	return WriteToFile(path, formatedOut)
}

func WriteToFile(path string, out []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create controller file %s: %w", path, err)
	}
	err = os.WriteFile(path, out, 0644)
	if err != nil {
		return fmt.Errorf("write controller file %s: %w", path, err)
	}
	defer f.Close()
	return nil
}
