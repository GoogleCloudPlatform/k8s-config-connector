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
	"regexp"
	"strings"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/apis"
	"github.com/fatih/color"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type APIScaffolder struct {
	BaseDir         string
	GoPackage       string
	Group           string
	Version         string
	PackageProtoTag string

	// Proto is optional. When set, the scaffolder reads google.api.resource to
	// learn the resource's real collection segment and parent shape instead of
	// guessing them. Without it the templates fall back to their old guesses,
	// which are wrong for most resources: the collection segment disagrees with
	// the declared pattern for 752 of 1417 annotated messages, and the assumed
	// projects/locations parent holds for about a third.
	Proto *protoapi.Proto

	// EmitParentRefs adds one Spec field naming the resource's direct parent,
	// where the pattern declares one below project and location and a reference
	// type for it already exists. Off by default: it adds a field to the CRD of a
	// resource people already use, so a service opts in one at a time.
	EmitParentRefs bool
}

// resourceMetadata looks up what the proto states about a resource, or nil if we
// have no proto loaded or the message is not an annotated resource.
func (a *APIScaffolder) resourceMetadata(fullName string) *protoapi.ResourceMetadata {
	if a.Proto == nil || fullName == "" {
		return nil
	}
	d, err := a.Proto.Files().FindDescriptorByName(protoreflect.FullName(fullName))
	if err != nil {
		klog.V(2).Infof("no descriptor for %q, scaffolding will guess: %v", fullName, err)
		return nil
	}
	msg, ok := d.(protoreflect.MessageDescriptor)
	if !ok {
		return nil
	}
	return protoapi.GetResourceMetadata(msg)
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
	fileName := strings.ToLower(resource.Kind) + "_reference.go"
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
	// Format generated code and organize imports.
	if err := FormatImports(path, out.Bytes()); err != nil {
		return err
	}
	color.HiGreen("New identity file added %s\nPlease EDIT it!\n", path)
	return nil
}

func (a *APIScaffolder) IdentityFileExist(resource options.Resource) bool {
	return fileExists(a.PathToIdentityFile(resource))
}

func (a *APIScaffolder) PathToIdentityFile(resource options.Resource) string {
	fileName := strings.ToLower(resource.Kind) + "_identity.go"
	return filepath.Join(a.BaseDir, a.GoPackage, fileName)
}

// Populates an APIArgs for templating.  Arguments are optional.
func (a *APIScaffolder) buildAPIArgs(resource *options.Resource) *apis.APIArgs {
	args := &apis.APIArgs{
		Group:           a.Group,
		Version:         a.Version,
		PackageProtoTag: a.PackageProtoTag,
	}
	// Without a pattern to read, every resource gets projectRef and a required
	// location. AddTypeFile replaces both from the pattern when prepopulating.
	args.RootRefField, _ = a.rootRef("")
	args.LocationField, _ = a.locationRef("")

	if resource != nil {
		args.Kind = resource.Kind

		if strings.Contains(resource.ProtoName, ".") {
			args.KindProtoTag = resource.ProtoName
		} else {
			pkg := a.PackageProtoTag
			if strings.Contains(pkg, ",") {
				pkg = strings.Split(pkg, ",")[0]
			}
			args.KindProtoTag = pkg + "." + resource.ProtoName
		}
		args.ProtoResource = resource.ProtoName

		args.ProtoMessageName = resource.ProtoMessageName()
		args.ProtoMessageFullName = resource.ProtoMessageFullName(a.PackageProtoTag)

		if md := a.resourceMetadata(args.ProtoMessageFullName); md != nil {
			args.Collection = md.Collection
			args.ParentStyle = string(md.ParentStyle)
			args.ResourcePattern = md.Pattern
			args.ResourcePatterns = md.Patterns
		}
		if args.Collection == "" {
			// Fall back to lowercased message name + 's' if no resource pattern is declared.
			args.Collection = strings.ToLower(args.ProtoMessageName) + "s"
		}
		if args.ParentStyle == "" {
			args.ParentStyle = string(protoapi.ParentUnknown)
		}
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
	fileName := strings.ToLower(resource.Kind) + "_types.go"
	return filepath.Join(a.BaseDir, a.GoPackage, fileName)
}

// AddTypeFile scaffolds <kind>_types.go.
//
// When prepopulated is non-nil, the Spec and ObservedState struct bodies are
// generated from the proto message definitions along with any required package imports.
func (a *APIScaffolder) AddTypeFile(resource options.Resource, prepopulated *PrepopulateResult) error {
	typeFilePath := a.PathToTypeFile(resource)
	cArgs := a.buildAPIArgs(&resource)
	cArgs.SkipGVK = packageDeclaresGVK(filepath.Join(a.BaseDir, a.GoPackage), cArgs.Kind)
	if prepopulated != nil {
		cArgs.SpecFields = prepopulated.SpecFields
		cArgs.ObservedStateFields = prepopulated.ObservedStateFields
		cArgs.ExtraImports = prepopulated.ExtraImports
		root, item := a.rootRef(cArgs.ResourcePattern)
		cArgs.RootRefField = root
		if item != nil {
			prepopulated.Judgement = append(prepopulated.Judgement, *item)
		}
		location, item := a.locationRef(cArgs.ResourcePattern)
		cArgs.LocationField = location
		if item != nil {
			prepopulated.Judgement = append(prepopulated.Judgement, *item)
		}

		// The parent field rides on prepopulated because that is the only way its
		// queue entry reaches the caller. A field emitted without its entry would
		// leave a +kcc:guess in the Spec that nothing flags.
		if a.EmitParentRefs {
			field, item := a.parentRef(cArgs.ResourcePattern)
			cArgs.ParentRefField = field
			if item != nil {
				prepopulated.Judgement = append(prepopulated.Judgement, *item)
			}
		}
	}
	return scaffoldTypeFile(typeFilePath, cArgs)
}

// packageDeclaresGVK reports whether a Go file in dir, other than a
// _types.go file, has a top-level "var <kind>GVK" line. It does not see a
// GVK declared inside a grouped var block, as dataform and iam declare theirs.
//
// The types template declares the GVK, and so do most hand-written
// <kind>_reference.go files in apis/. Without this check, scaffolding a
// Kind's types next to its reference file declares the variable twice, and
// the package does not compile.
func packageDeclaresGVK(dir, kind string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	want := regexp.MustCompile(`(?m)^var ` + regexp.QuoteMeta(kind) + `GVK\b`)
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") ||
			strings.HasSuffix(e.Name(), "_types.go") {
			continue
		}
		body, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err == nil && want.Match(body) {
			return true
		}
	}
	return false
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
