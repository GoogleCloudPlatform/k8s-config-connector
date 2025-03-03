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

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/apis"
	"github.com/fatih/color"
)

type APIScaffolder struct {
	BaseDir         string
	GoPackage       string
	Group           string
	Version         string
	PackageProtoTag string
}

func (a *APIScaffolder) RefsFileExist(kind, resourceProtoName string) bool {
	refsFilePath := a.PathToRefsFile(kind, resourceProtoName)
	_, err := os.Stat(refsFilePath)
	if err == nil {
		return true
	}
	return !errors.Is(err, os.ErrNotExist)
}

func (a *APIScaffolder) PathToRefsFile(kind, resourceProtoName string) string {
	fileName := strings.ToLower(resourceProtoName) + "_reference.go"
	return filepath.Join(a.BaseDir, a.GoPackage, fileName)
}

func (a *APIScaffolder) AddRefsFile(kind, resourceProtoName string) error {
	refsFilePath := a.PathToRefsFile(kind, resourceProtoName)
	cArgs := &apis.APIArgs{
		Group:           a.Group,
		Version:         a.Version,
		Kind:            kind,
		PackageProtoTag: a.PackageProtoTag,
		KindProtoTag:    a.PackageProtoTag + "." + resourceProtoName,
		ProtoResource:   resourceProtoName,
	}
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

func (a *APIScaffolder) IdentityFileExist(kind, resourceProtoName string) bool {
	refsFilePath := a.PathToIdentityFile(kind, resourceProtoName)
	_, err := os.Stat(refsFilePath)
	if err == nil {
		return true
	}
	return !errors.Is(err, os.ErrNotExist)
}

func (a *APIScaffolder) PathToIdentityFile(kind, resourceProtoName string) string {
	fileName := strings.ToLower(resourceProtoName) + "_identity.go"
	return filepath.Join(a.BaseDir, a.GoPackage, fileName)
}

func (a *APIScaffolder) AddIdentityFile(kind, resourceProtoName string) error {
	refsFilePath := a.PathToIdentityFile(kind, resourceProtoName)
	cArgs := &apis.APIArgs{
		Group:           a.Group,
		Version:         a.Version,
		Kind:            kind,
		PackageProtoTag: a.PackageProtoTag,
		KindProtoTag:    a.PackageProtoTag + "." + resourceProtoName,
		ProtoResource:   resourceProtoName,
	}
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

func (a *APIScaffolder) TypeFileNotExist(resourceProtoName string) bool {
	typeFilePath := a.PathToTypeFile(resourceProtoName)
	_, err := os.Stat(typeFilePath)
	if err == nil {
		return false
	}
	return errors.Is(err, os.ErrNotExist)
}

func (a *APIScaffolder) PathToTypeFile(resourceProtoName string) string {
	fileName := strings.ToLower(resourceProtoName) + "_types.go"
	return filepath.Join(a.BaseDir, a.GoPackage, fileName)
}

func (a *APIScaffolder) AddTypeFile(resourceProtoName, kind string) error {
	typeFilePath := a.PathToTypeFile(resourceProtoName)
	cArgs := &apis.APIArgs{
		Group:           a.Group,
		Version:         a.Version,
		Kind:            kind,
		PackageProtoTag: a.PackageProtoTag,
		KindProtoTag:    a.PackageProtoTag + "." + resourceProtoName,
		ProtoResource:   resourceProtoName,
	}
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
	cArgs := &apis.APIArgs{
		Group:           a.Group,
		Version:         a.Version,
		PackageProtoTag: a.PackageProtoTag,
	}
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
	cArgs := &apis.APIArgs{
		Group:           a.Group,
		Version:         a.Version,
		PackageProtoTag: a.PackageProtoTag,
	}
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

// AddTypeFileV2 creates a new types.go file using the V2 template
func (s *APIScaffolder) AddTypeFileV2(protoName string, kind string) error {
	typeFilePath := s.PathToTypeFile(protoName)
	args := &apis.APIArgs{
		Group:           s.Group,
		Version:         s.Version,
		Kind:            kind,
		KindProtoTag:    s.PackageProtoTag + "." + protoName,
		PackageProtoTag: s.PackageProtoTag,
		ProtoResource:   protoName,
	}
	return scaffoldTypeFileV2(typeFilePath, args)
}

func scaffoldTypeFileV2(path string, args *apis.APIArgs) error {
	tmpl, err := template.New(args.Kind).Funcs(funcMap).Parse(apis.TypesV2Template)
	if err != nil {
		return fmt.Errorf("parse %s_types.go template: %w", strings.ToLower(args.ProtoResource), err)
	}
	// Apply the APIArgs args to the template
	out := &bytes.Buffer{}
	if err := tmpl.Execute(out, args); err != nil {
		return err
	}
	// Write the generated <kind>_types.go
	if err := WriteToFile(path, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New API file added %s\nPlease EDIT it!\n", path)
	return nil
}
