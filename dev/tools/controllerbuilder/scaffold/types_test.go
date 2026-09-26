// Copyright 2026 Google LLC
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
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/apis"
)

// TestTypesTemplateRendersValidGo verifies that the types template renders valid Go
// code for both stub and prepopulated configurations.
func TestTypesTemplateRendersValidGo(t *testing.T) {
	base := apis.APIArgs{
		Group:                "networkservices.cnrm.cloud.google.com",
		Version:              "v1alpha1",
		Kind:                 "NetworkServicesLBTrafficExtension",
		ProtoResource:        "LbTrafficExtension",
		PackageProtoTag:      "google.cloud.networkservices.v1",
		KindProtoTag:         "google.cloud.networkservices.v1.LbTrafficExtension",
		ProtoMessageName:     "LbTrafficExtension",
		ProtoMessageFullName: "google.cloud.networkservices.v1.LbTrafficExtension",
	}
	prepopulated := base
	prepopulated.SpecFields = "\t// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.description\n" +
		"\tDescription *string `json:\"description,omitempty\"`\n"
	prepopulated.ObservedStateFields = "\t// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.create_time\n" +
		"\tCreateTime *string `json:\"createTime,omitempty\"`\n"
	prepopulated.ExtraImports = []string{`common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"`}

	for _, tc := range []struct {
		name string
		args apis.APIArgs
	}{
		{name: "stub", args: base},
		{name: "prepopulated", args: prepopulated},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			tmpl, err := template.New(tc.args.Kind).Funcs(funcMap).Parse(apis.TypesTemplate)
			if err != nil {
				t.Fatalf("parsing types template: %v", err)
			}

			// Act
			var buf bytes.Buffer
			execErr := tmpl.Execute(&buf, &tc.args)

			// Assert
			if execErr != nil {
				t.Fatalf("executing types template: %v", execErr)
			}
			if _, err := parser.ParseFile(token.NewFileSet(), "types.go", buf.Bytes(), parser.AllErrors); err != nil {
				t.Errorf("rendered types file is not valid Go: %v\n%s", err, buf.String())
			}
		})
	}
}

// TestAddTypeFileWritesPrepopulatedBodies verifies that AddTypeFile writes all
// components of PrepopulateResult (Spec fields, ObservedState fields, and ExtraImports)
// into the scaffolded file.
func TestAddTypeFileWritesPrepopulatedBodies(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	scaffolder := &APIScaffolder{
		BaseDir:         dir,
		GoPackage:       "networkservices/v1alpha1",
		Group:           "networkservices.cnrm.cloud.google.com",
		Version:         "v1alpha1",
		PackageProtoTag: "google.cloud.networkservices.v1",
	}
	resource := options.Resource{Kind: "NetworkServicesLBTrafficExtension", ProtoName: "LbTrafficExtension"}
	prepopulated := &PrepopulateResult{
		SpecFields:          "\tDescription *string `json:\"description,omitempty\"`\n",
		ObservedStateFields: "\tError *common.Status `json:\"error,omitempty\"`\n",
		ExtraImports:        []string{`common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"`},
	}

	// Act
	if err := scaffolder.AddTypeFile(resource, prepopulated); err != nil {
		t.Fatalf("AddTypeFile: %v", err)
	}

	// Assert
	b, err := os.ReadFile(filepath.Join(dir, scaffolder.GoPackage, "networkserviceslbtrafficextension_types.go"))
	if err != nil {
		t.Fatalf("reading scaffolded file: %v", err)
	}
	got := string(b)
	if !strings.Contains(got, "Description *string `json:\"description,omitempty\"`") {
		t.Errorf("the Spec body is missing from the scaffolded file:\n%s", got)
	}
	if !strings.Contains(got, "Error *common.Status `json:\"error,omitempty\"`") {
		t.Errorf("the ObservedState body is missing from the scaffolded file:\n%s", got)
	}
	if !strings.Contains(got, `common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"`) {
		t.Errorf("the import the ObservedState body needs is missing from the scaffolded file:\n%s", got)
	}
}
