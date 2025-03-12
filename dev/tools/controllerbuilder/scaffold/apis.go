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
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/apis"
	"github.com/fatih/color"
	"k8s.io/klog/v2"
)

type APIScaffolder struct {
	BaseDir         string
	GoPackage       string
	Group           string
	Version         string
	PackageProtoTag string
}

func fileExists(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	klog.Fatalf("unexpected error checking for file %q: %v", p, err)
	return false
}

func (a *APIScaffolder) RefsFileExist(resource options.Resource) bool {
	return fileExists(a.PathToRefsFile(resource))
}

func (a *APIScaffolder) PathToRefsFile(resource options.Resource) string {
	fileName := strings.ToLower(resource.ProtoMessageName()) + "_reference.go"
	return filepath.Join(a.BaseDir, a.GoPackage, fileName)
}

func (a *APIScaffolder) AddRefsFile(resource options.Resource) error {
	refsFilePath := a.PathToRefsFile(resource)
	cArgs := a.buildAPIArgs(&resource)
	return scaffoldRefsFile(refsFilePath, cArgs)
}

func scaffoldIdentityFile(path string, cArgs *apis.APIArgs) error {
	tmpl, err := template.New(cArgs.Kind).Funcs(funcMap).Parse(apis.IdentityTemplate)
	if err != nil {
		return fmt.Errorf("parse %s_identity.go template: %w", strings.ToLower(cArgs.ProtoResource), err)
	}
	// Apply the APIArgs args to the template
	out := &bytes.Buffer{}
	if err := tmpl.Execute(out, cArgs); err != nil {
		return err
	}
	// Write the generated <kind>_types.go
	if err := WriteToFile(path, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New identity file added %s\nPlease EDIT it!\n", path)
	return nil
}

func (a *APIScaffolder) IdentityFileExist(resource options.Resource) bool {
	return fileExists(a.PathToIdentityFile(resource))
}

func (a *APIScaffolder) PathToIdentityFile(resource options.Resource) string {
	fileName := strings.ToLower(resource.ProtoMessageName()) + "_identity.go"
	return filepath.Join(a.BaseDir, a.GoPackage, fileName)
}

// Populates an APIArgs for templating.  Arguments are optional.
func (a *APIScaffolder) buildAPIArgs(resource *options.Resource) *apis.APIArgs {
	args := &apis.APIArgs{
		Group:           a.Group,
		Version:         a.Version,
		PackageProtoTag: a.PackageProtoTag,
	}

	if resource != nil {
		args.Kind = resource.Kind

		args.KindProtoTag = a.PackageProtoTag + "." + resource.ProtoName
		args.ProtoResource = resource.ProtoName

		args.ProtoMessageName = resource.ProtoMessageName()
		args.ProtoMessageFullName = resource.ProtoMessageFullName(a.PackageProtoTag)
	}

	return args
}

func (a *APIScaffolder) AddIdentityFile(resource options.Resource) error {
	refsFilePath := a.PathToIdentityFile(resource)
	cArgs := a.buildAPIArgs(&resource)
	return scaffoldIdentityFile(refsFilePath, cArgs)
}

func scaffoldRefsFile(path string, cArgs *apis.APIArgs) error {
	tmpl, err := template.New(cArgs.Kind).Funcs(funcMap).Parse(apis.RefsHeaderTemplate)
	if err != nil {
		return fmt.Errorf("parse %s_reference.go template: %w", strings.ToLower(cArgs.ProtoResource), err)
	}
	// Apply the APIArgs args to the template
	out := &bytes.Buffer{}
	if err := tmpl.Execute(out, cArgs); err != nil {
		return err
	}
	// Write the generated <kind>_types.go
	if err := WriteToFile(path, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New reference file added %s\nPlease EDIT it!\n", path)
	return nil
}

func (a *APIScaffolder) TypeFileExists(resource options.Resource) bool {
	return fileExists(a.PathToTypeFile(resource))
}

func (a *APIScaffolder) PathToTypeFile(resource options.Resource) string {
	fileName := strings.ToLower(resource.ProtoMessageName()) + "_types.go"
	return filepath.Join(a.BaseDir, a.GoPackage, fileName)
}

func (a *APIScaffolder) AddTypeFile(resource options.Resource) error {
	typeFilePath := a.PathToTypeFile(resource)
	cArgs := a.buildAPIArgs(&resource)
	return scaffoldTypeFile(typeFilePath, cArgs)
}

func scaffoldTypeFile(path string, cArgs *apis.APIArgs) error {
	tmpl, err := template.New(cArgs.Kind).Funcs(funcMap).Parse(apis.TypesTemplate)
	if err != nil {
		return fmt.Errorf("parse %s_types.go template: %w", strings.ToLower(cArgs.ProtoResource), err)
	}
	// Apply the APIArgs args to the template
	out := &bytes.Buffer{}
	if err := tmpl.Execute(out, cArgs); err != nil {
		return err
	}
	// Write the generated <kind>_types.go
	if err := WriteToFile(path, out.Bytes()); err != nil {
		return err
	}
	// Format and adjust the go imports in the generated files.
	if err := FormatImports(path, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New API file added %s\nPlease EDIT it!\n", path)
	return nil
}

func (a *APIScaffolder) GroupVersionFileNotExist() bool {
	docFilePath := filepath.Join(a.BaseDir, a.GoPackage, "groupversion_info.go")
	_, err := os.Stat(docFilePath)
	if err == nil {
		return false
	}
	return errors.Is(err, os.ErrNotExist)
}

func (a *APIScaffolder) AddGroupVersionFile() error {
	docFilePath := filepath.Join(a.BaseDir, a.GoPackage, "groupversion_info.go")
	cArgs := a.buildAPIArgs(nil)
	return scaffoldGroupVersionFile(docFilePath, cArgs)
}

func (a *APIScaffolder) DocFileNotExist() bool {
	docFilePath := filepath.Join(a.BaseDir, a.GoPackage, "doc.go")
	_, err := os.Stat(docFilePath)
	if err == nil {
		return false
	}
	return errors.Is(err, os.ErrNotExist)
}

func (a *APIScaffolder) AddDocFile() error {
	docFilePath := filepath.Join(a.BaseDir, a.GoPackage, "doc.go")
	cArgs := a.buildAPIArgs(nil)
	return scaffoldDocFile(docFilePath, cArgs)
}

func scaffoldDocFile(path string, cArgs *apis.APIArgs) error {
	tmpl, err := template.New("doc.go").Parse(apis.DocTemplate)
	if err != nil {
		return fmt.Errorf("parse doc.go template: %w", err)
	}
	out := &bytes.Buffer{}
	if err := tmpl.Execute(out, cArgs); err != nil {
		return err
	}
	if err := WriteToFile(path, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New file added %q\n", path)
	return nil
}

func scaffoldGroupVersionFile(path string, cArgs *apis.APIArgs) error {
	tmpl, err := template.New("groupversioninfo.go").Parse(apis.GroupVersionInfoTemplate)
	if err != nil {
		return fmt.Errorf("parse groupversion_info.go template: %w", err)
	}
	out := &bytes.Buffer{}
	if err := tmpl.Execute(out, cArgs); err != nil {
		return err
	}
	if err := WriteToFile(path, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New file added %q\n", path)
	return nil
}
