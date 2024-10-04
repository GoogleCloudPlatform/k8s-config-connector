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
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	ccTemplate "github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/controller"
	"github.com/fatih/color"
	"golang.org/x/tools/imports"
)

const (
	currRelPath             = "dev/tools/controllerbuilder"
	directControllerRelPath = "pkg/controller/direct"
)

var funcMap = template.FuncMap{
	"ToLower": strings.ToLower,
}

func Scaffold(service, kind string, cArgs *ccTemplate.ControllerArgs) error {
	if err := generateController(service, kind, cArgs); err != nil {
		return err
	}
	return nil
}

func generateController(service, kind string, cArgs *ccTemplate.ControllerArgs) error {
	tmpl, err := template.New(cArgs.Kind).Funcs(funcMap).Parse(ccTemplate.ControllerTemplate)
	if err != nil {
		return fmt.Errorf("parse controller template: %w", err)
	}
	// Apply the `service` and `resource` args to the controller and external resource templates
	controllerOutput := &bytes.Buffer{}
	if err := tmpl.Execute(controllerOutput, cArgs); err != nil {
		return err
	}

	controllerFilePath, err := buildControllerPath(service, kind)
	if err != nil {
		return err
	}

	// Write the generated controller.go to  pkg/controller/direct/<service>/<resource>/<resource>_controller.go
	if err := WriteToFile(controllerFilePath, controllerOutput.Bytes()); err != nil {
		return err
	}
	// Format and adjust the go imports in the generated controller file.
	if err := FormatImports(controllerFilePath, controllerOutput.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New controller %s has been generated. \nEnjoy it!\n", kind)
	return nil
}

func buildResourcePath(service, filename string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get current working directory: %w", err)
	}
	abs, err := filepath.Abs(pwd)
	if err != nil {
		return "", fmt.Errorf("get absolute path %s: %w", pwd, err)
	}
	seg := strings.Split(abs, currRelPath)
	controllerDir := filepath.Join(seg[0], directControllerRelPath, service, kind)
	err = os.MkdirAll(controllerDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create controller directory %s: %w", controllerDir, err)
	}
	resourceFilePath := filepath.Join(controllerDir, filename)
	if _, err = os.Stat(resourceFilePath); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return "", fmt.Errorf("could not stat path %s: %w", resourceFilePath, err)
		}
		// otherwise create the file
		return resourceFilePath, nil
	}
	return "", fmt.Errorf("file %s already exist", resourceFilePath)
}

func buildControllerPath(service, kind string) (string, error) {
	kind = strings.ToLower(kind)
	return buildResourcePath(service, kind, kind+"_controller.go")
}

func buildExternalResourcePath(service, kind string) (string, error) {
	kind = strings.ToLower(kind)
	return buildResourcePath(service, kind, kind+"_externalresource.go")
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
