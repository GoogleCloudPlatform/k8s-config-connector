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
	"fmt"
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

func TestGeneratedFileWriteAdditiveTypes(t *testing.T) {
	baseDir := t.TempDir()
	key := generatedFileKey{
		GoPackage: "dialogflowcx/v1alpha1",
		FileName:  "types.generated.go",
	}

	writeFirst := func() {
		f1 := &generatedFile{
			baseDir:   baseDir,
			key:       key,
			goPackage: "v1alpha1",
			fileAnnotation: &annotations.FileAnnotation{
				Key: "+generated:types",
				Attributes: map[string][]string{
					"krm.group":     {"dialogflowcx.cnrm.cloud.google.com"},
					"krm.version":   {"v1alpha1"},
					"proto.service": {"google.cloud.dialogflow.cx.v3"},
					"resource":      {"DialogflowCXAgent:Agent"},
				},
			},
		}
		f1.body.WriteString(`
// +kcc:proto=google.cloud.dialogflow.cx.v3.SpeechToTextSettings
type SpeechToTextSettings struct {
	EnableSpeechAdaptation *bool ` + "`json:\"enableSpeechAdaptation,omitempty\"`" + `
}

// +kcc:observedstate:proto=google.cloud.dialogflow.cx.v3.SpeechToTextSettings
type SpeechToTextSettingsObservedState struct {
	OutputField *string ` + "`json:\"outputField,omitempty\"`" + `
}
`)
		if err := f1.Write(true, true, true); err != nil {
			t.Fatalf("first Write failed: %v", err)
		}
	}

	writeSecond := func() {
		f2 := &generatedFile{
			baseDir:   baseDir,
			key:       key,
			goPackage: "v1alpha1",
			fileAnnotation: &annotations.FileAnnotation{
				Key: "+generated:types",
				Attributes: map[string][]string{
					"krm.group":     {"dialogflowcx.cnrm.cloud.google.com"},
					"krm.version":   {"v1alpha1"},
					"proto.service": {"google.cloud.dialogflow.v2"},
					"resource":      {"DialogflowCXIntent:Intent"},
				},
			},
		}
		f2.body.WriteString(`
// +kcc:proto=google.cloud.dialogflow.v2.Intent.Message
type Intent_Message struct {
	Platform *string ` + "`json:\"platform,omitempty\"`" + `
}
`)
		if err := f2.Write(true, true, true); err != nil {
			t.Fatalf("second Write failed: %v", err)
		}
	}

	writeFirst()
	writeSecond()

	targetPath := filepath.Join(baseDir, "dialogflowcx", "v1alpha1", "types.generated.go")
	afterFirstPassBytes, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("reading generated file: %v", err)
	}
	afterFirstPass := string(afterFirstPassBytes)

	for _, wantSubstr := range []string{
		"// proto.service: google.cloud.dialogflow.cx.v3",
		"// proto.service: google.cloud.dialogflow.v2",
		"// resource: DialogflowCXAgent:Agent",
		"// resource: DialogflowCXIntent:Intent",
		"type SpeechToTextSettings struct",
		"type Intent_Message struct",
		"type SpeechToTextSettingsObservedState struct",
	} {
		if !strings.Contains(afterFirstPass, wantSubstr) {
			t.Errorf("expected types.generated.go to contain %q, got:\n%s", wantSubstr, afterFirstPass)
		}
	}

	// Regular structs must precede ObservedState structs.
	if strings.Index(afterFirstPass, "type Intent_Message struct") > strings.Index(afterFirstPass, "type SpeechToTextSettingsObservedState struct") {
		t.Errorf("expected regular struct Intent_Message before ObservedState struct:\n%s", afterFirstPass)
	}

	// Re-running first then second must be completely idempotent.
	writeFirst()
	writeSecond()
	afterSecondPassBytes, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("reading generated file after re-run: %v", err)
	}
	if diff := cmp.Diff(afterFirstPass, string(afterSecondPassBytes)); diff != "" {
		t.Errorf("re-running additive Write was not idempotent (-first +second):\n%s", diff)
	}
}

func TestGeneratedFileWriteAdditiveMappers(t *testing.T) {
	baseDir := t.TempDir()
	key := generatedFileKey{
		GoPackage: "merchantapi",
		FileName:  "mapper.generated.go",
	}

	writeAccounts := func() {
		mg := NewMapperGenerator(nil, baseDir, &annotations.FileAnnotation{
			Key: "+generated:mapper",
			Attributes: map[string][]string{
				"krm.group":     {"merchantapi.cnrm.cloud.google.com"},
				"krm.version":   {"v1alpha1"},
				"proto.service": {"google.shopping.merchant.accounts.v1"},
			},
		}, false)
		mg.LoadExistingImports("merchantapi")
		pbAlias := mg.AddGoImportAlias("cloud.google.com/go/shopping/merchant/accounts/apiv1/accountspb", "pb")

		f := mg.getOutputFile(key)
		f.goPackage = "merchantapi"
		f.fileAnnotation = mg.generatedFileAnnotation
		f.validKRMTypes = map[string]bool{"AccountType": true, "DataSourceType": true}
		f.addImport("", "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct")
		f.addImport("krm", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/merchantapi/v1alpha1")
		for _, imp := range mg.importedPackages {
			f.addImport(imp.alias, imp.goPackage)
		}
		fmt.Fprintf(&f.body, `
func AccountType_FromProto(mapCtx *direct.MapContext, in *%s.AccountType) *krm.AccountType {
	if in == nil {
		return nil
	}
	out := &krm.AccountType{}
	return out
}
`, pbAlias)
		if err := f.Write(true, true, false); err != nil {
			t.Fatalf("writeAccounts failed: %v", err)
		}
	}

	writeDataSources := func() {
		mg := NewMapperGenerator(nil, baseDir, &annotations.FileAnnotation{
			Key: "+generated:mapper",
			Attributes: map[string][]string{
				"krm.group":     {"merchantapi.cnrm.cloud.google.com"},
				"krm.version":   {"v1alpha1"},
				"proto.service": {"google.shopping.merchant.datasources.v1"},
			},
		}, false)
		mg.LoadExistingImports("merchantapi")
		pbAlias := mg.AddGoImportAlias("cloud.google.com/go/shopping/merchant/datasources/apiv1/datasourcespb", "pb")
		if pbAlias == "pb" {
			t.Fatalf("expected non-colliding alias for datasourcespb when pb is already bound to accountspb, got %q", pbAlias)
		}

		f := mg.getOutputFile(key)
		f.goPackage = "merchantapi"
		f.fileAnnotation = mg.generatedFileAnnotation
		f.validKRMTypes = map[string]bool{"AccountType": true, "DataSourceType": true}
		f.addImport("", "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct")
		f.addImport("krm", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/merchantapi/v1alpha1")
		for _, imp := range mg.importedPackages {
			f.addImport(imp.alias, imp.goPackage)
		}
		fmt.Fprintf(&f.body, `
func DataSourceType_FromProto(mapCtx *direct.MapContext, in *%s.DataSourceType) *krm.DataSourceType {
	if in == nil {
		return nil
	}
	out := &krm.DataSourceType{}
	return out
}
`, pbAlias)
		if err := f.Write(true, true, false); err != nil {
			t.Fatalf("writeDataSources failed: %v", err)
		}
	}

	writeAccounts()
	writeDataSources()

	targetPath := filepath.Join(baseDir, "merchantapi", "mapper.generated.go")
	afterFirstPassBytes, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("reading mapper.generated.go: %v", err)
	}
	afterFirstPass := string(afterFirstPassBytes)

	for _, wantSubstr := range []string{
		"// proto.service: google.shopping.merchant.accounts.v1",
		"// proto.service: google.shopping.merchant.datasources.v1",
		`pb "cloud.google.com/go/shopping/merchant/accounts/apiv1/accountspb"`,
		`datasourcespb "cloud.google.com/go/shopping/merchant/datasources/apiv1/datasourcespb"`,
		"func AccountType_FromProto(mapCtx *direct.MapContext, in *pb.AccountType) *krm.AccountType",
		"func DataSourceType_FromProto(mapCtx *direct.MapContext, in *datasourcespb.DataSourceType) *krm.DataSourceType",
	} {
		if !strings.Contains(afterFirstPass, wantSubstr) {
			t.Errorf("expected mapper.generated.go to contain %q, got:\n%s", wantSubstr, afterFirstPass)
		}
	}

	// Re-running only the first service (writeAccounts) must update in-place without flipping function order.
	writeAccounts()
	afterAccountsOnlyBytes, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("reading mapper.generated.go after single-service re-run: %v", err)
	}
	if diff := cmp.Diff(afterFirstPass, string(afterAccountsOnlyBytes)); diff != "" {
		t.Errorf("re-running first service alone flipped or altered mapper.generated.go (-first +afterAccountsOnly):\n%s", diff)
	}

	// Re-running both in sequence must be idempotent.
	writeAccounts()
	writeDataSources()
	afterSecondPassBytes, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatalf("reading mapper.generated.go after re-run: %v", err)
	}
	if diff := cmp.Diff(afterFirstPass, string(afterSecondPassBytes)); diff != "" {
		t.Errorf("re-running additive mapper Write was not idempotent (-first +second):\n%s", diff)
	}
}
