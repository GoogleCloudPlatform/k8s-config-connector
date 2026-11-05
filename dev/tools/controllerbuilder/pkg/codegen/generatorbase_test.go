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

package codegen

import (
	"maps"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/google/go-cmp/cmp"
)

func TestUsedImports(t *testing.T) {
	const (
		refsPkg     = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
		directPkg   = "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
		commonPkg   = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
		apiextPkg   = "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
		unusedAlias = "secretmanagerv1beta1"
		unusedPkg   = "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	)
	for _, tc := range []struct {
		name    string
		body    string
		imports map[string]string
		want    map[string]string
	}{
		{
			name:    "an aliased import is kept when the body uses the alias",
			body:    "var x *refsv1beta1.ProjectRef\n",
			imports: map[string]string{refsPkg: "refsv1beta1"},
			want:    map[string]string{refsPkg: "refsv1beta1"},
		},
		{
			name:    "an import without an alias is matched on its last path element",
			body:    "var x = direct.ValueOf(y)\n",
			imports: map[string]string{directPkg: ""},
			want:    map[string]string{directPkg: ""},
		},
		{
			name:    "an import the body never uses is dropped",
			body:    "var x common.Status\n",
			imports: map[string]string{commonPkg: "common", unusedPkg: unusedAlias},
			want:    map[string]string{commonPkg: "common"},
		},
		{
			name: "a qualifier from QualifierImports is added when the body uses it",
			body: "Labels *apiextensionsv1.JSON `json:\"labels,omitempty\"`\n",
			want: map[string]string{apiextPkg: "apiextensionsv1"},
		},
		{
			name: "no imports and no qualifiers in the body",
			body: "var x string\n",
			want: nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			f := &generatedFile{}
			f.body.WriteString(tc.body)
			for pkgName, alias := range tc.imports {
				f.addImport(alias, pkgName)
			}

			// Act
			got := f.usedImports()

			// Assert
			if !maps.Equal(got, tc.want) {
				t.Errorf("usedImports() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGeneratedFileWriteAdditiveMerge(t *testing.T) {
	baseDir := t.TempDir()
	key := generatedFileKey{
		GoPackage: "dialogflowcx/v1alpha1",
		FileName:  "types.generated.go",
	}

	writeService := func(service, resource, body string) {
		f := &generatedFile{
			baseDir:   baseDir,
			key:       key,
			goPackage: "v1alpha1",
			fileAnnotation: &annotations.FileAnnotation{
				Key: "+generated:types",
				Attributes: map[string][]string{
					"krm.group":     {"dialogflowcx.cnrm.cloud.google.com"},
					"krm.version":   {"v1alpha1"},
					"proto.service": {service},
					"resource":      {resource},
				},
			},
		}

		f.body.WriteString(body)
		if err := f.Write(true, true, true); err != nil {
			t.Fatalf("Write(%s) failed: %v", service, err)
		}
	}

	firstBody := `
// +kcc:proto=google.cloud.dialogflow.cx.v3.SpeechToTextSettings
type SpeechToTextSettings struct {
	EnableSpeechAdaptation *bool ` + "`json:\"enableSpeechAdaptation,omitempty\"`" + `
}
`
	secondBody := `
// +kcc:proto=google.cloud.dialogflow.v2.Intent.Message
type Intent_Message struct {
	Platform *string ` + "`json:\"platform,omitempty\"`" + `
}
`

	writeService("google.cloud.dialogflow.cx.v3", "DialogflowCXAgent:Agent", firstBody)
	writeService("google.cloud.dialogflow.v2", "DialogflowCXIntent:Intent", secondBody)

	targetPath := filepath.Join(baseDir, "dialogflowcx", "v1alpha1", "types.generated.go")
	firstPassBytes, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("reading generated file: %v", err)
	}

	firstPass := string(firstPassBytes)

	for _, want := range []string{
		"// proto.service: google.cloud.dialogflow.cx.v3",
		"// proto.service: google.cloud.dialogflow.v2",
		"type SpeechToTextSettings struct",
		"type Intent_Message struct",
	} {
		if !strings.Contains(firstPass, want) {
			t.Errorf("expected types.generated.go to contain %q, got:\n%s", want, firstPass)
		}
	}

	// Re-running both writes must be idempotent.
	writeService("google.cloud.dialogflow.cx.v3", "DialogflowCXAgent:Agent", firstBody)
	writeService("google.cloud.dialogflow.v2", "DialogflowCXIntent:Intent", secondBody)
	secondPassBytes, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("reading generated file after re-run: %v", err)
	}

	if diff := cmp.Diff(firstPass, string(secondPassBytes)); diff != "" {
		t.Errorf("re-running additive Write was not idempotent (-first +second):\n%s", diff)
	}
}
