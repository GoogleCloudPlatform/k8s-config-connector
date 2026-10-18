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
		var gb generatorBase
		gb.init(baseDir)
		f := gb.getOutputFile(key)
		f.goPackage = "merchantapi"
		f.fileAnnotation = &annotations.FileAnnotation{
			Key: "+generated:mapper",
			Attributes: map[string][]string{
				"krm.group":     {"merchantapi.cnrm.cloud.google.com"},
				"krm.version":   {"v1alpha1"},
				"proto.service": {"google.shopping.merchant.accounts.v1"},
			},
		}
		f.validKRMTypes = map[string]bool{"AccountType": true, "DataSourceType": true}
		f.addImport("", "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct")
		f.addImport("krm", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/merchantapi/v1alpha1")
		f.addImport("pb", "cloud.google.com/go/shopping/merchant/accounts/apiv1/accountspb")
		f.body.WriteString(`
func AccountType_FromProto(mapCtx *direct.MapContext, in *pb.AccountType) *krm.AccountType {
	if in == nil {
		return nil
	}
	out := &krm.AccountType{}
	return out
}
`)
		if err := f.Write(true, true, false); err != nil {
			t.Fatalf("writeAccounts failed: %v", err)
		}
	}

	writeDataSources := func() {
		var gb generatorBase
		gb.init(baseDir)
		f := gb.getOutputFile(key)
		f.goPackage = "merchantapi"
		f.fileAnnotation = &annotations.FileAnnotation{
			Key: "+generated:mapper",
			Attributes: map[string][]string{
				"krm.group":     {"merchantapi.cnrm.cloud.google.com"},
				"krm.version":   {"v1alpha1"},
				"proto.service": {"google.shopping.merchant.datasources.v1"},
			},
		}
		f.validKRMTypes = map[string]bool{"AccountType": true, "DataSourceType": true}
		f.addImport("", "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct")
		f.addImport("krm", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/merchantapi/v1alpha1")
		f.addImport("datasourcespb", "cloud.google.com/go/shopping/merchant/datasources/apiv1/datasourcespb")
		f.body.WriteString(`
func DataSourceType_FromProto(mapCtx *direct.MapContext, in *datasourcespb.DataSourceType) *krm.DataSourceType {
	if in == nil {
		return nil
	}
	out := &krm.DataSourceType{}
	return out
}
`)
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

func TestGeneratedFileWriteSkipsCommentOnlyExternalBlocks(t *testing.T) {
	baseDir := t.TempDir()
	key := generatedFileKey{
		GoPackage: "osconfig/v1beta1",
		FileName:  "types.generated.go",
	}

	// First run has only commented-out blocks (all active types already moved to handwritten files).
	var gb1 generatorBase
	gb1.init(baseDir)
	f1 := gb1.getOutputFile(key)
	f1.goPackage = "v1beta1"
	f1.fileAnnotation = &annotations.FileAnnotation{
		Key: "+generated:types",
		Attributes: map[string][]string{
			"krm.group":     {"osconfig.cnrm.cloud.google.com"},
			"krm.version":   {"v1beta1"},
			"proto.service": {"google.cloud.osconfig.v1"},
			"resource":      {"OSConfigOSPolicyAssignment:OSPolicyAssignment"},
		},
	}
	f1.body.WriteString(`
/* found existing non-generated go type "OSPolicyAssignment", skipping
// +kcc:proto=google.cloud.osconfig.v1.OSPolicyAssignment
type OSPolicyAssignment struct {
}
*/
`)
	if err := f1.Write(true, true, true); err != nil {
		t.Fatalf("f1.Write failed: %v", err)
	}

	// Second run must replace without merging dead comments or file annotations from the comment-only first run.
	var gb2 generatorBase
	gb2.init(baseDir)
	f2 := gb2.getOutputFile(key)
	f2.goPackage = "v1beta1"
	f2.fileAnnotation = &annotations.FileAnnotation{
		Key: "+generated:types",
		Attributes: map[string][]string{
			"krm.group":     {"osconfig.cnrm.cloud.google.com"},
			"krm.version":   {"v1beta1"},
			"proto.service": {"google.cloud.osconfig.v1beta"},
			"resource":      {"OSConfigGuestPolicy:GuestPolicy"},
		},
	}
	f2.body.WriteString(`
/* found existing non-generated go type "GuestPolicy", skipping
// +kcc:proto=google.cloud.osconfig.v1beta.GuestPolicy
type GuestPolicy struct {
}
*/
`)
	if err := f2.Write(true, true, true); err != nil {
		t.Fatalf("f2.Write failed: %v", err)
	}

	gotBytes, err := os.ReadFile(filepath.Join(baseDir, "osconfig", "v1beta1", "types.generated.go"))
	if err != nil {
		t.Fatalf("reading types.generated.go: %v", err)
	}
	got := string(gotBytes)
	if strings.Contains(got, "OSPolicyAssignment") || strings.Contains(got, "google.cloud.osconfig.v1\n") {
		t.Errorf("expected comment-only first run not to be merged into second run, got:\n%s", got)
	}
}

func TestGeneratedFileWriteSkipsSeparateGeneratedTypesTarget(t *testing.T) {
	baseDir := t.TempDir()
	pkgDir := filepath.Join(baseDir, "dialogflow", "v1alpha1")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	// Existing securitysettings_types.generated.go claims DialogflowSecuritySettings:SecuritySettings.
	sepFile := filepath.Join(pkgDir, "securitysettings_types.generated.go")
	if err := os.WriteFile(sepFile, []byte(`// +generated:types
// krm.group: dialogflow.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.dialogflow.cx.v3
// resource: DialogflowSecuritySettings:SecuritySettings

package v1alpha1
`), 0644); err != nil {
		t.Fatalf("writing securitysettings_types.generated.go: %v", err)
	}

	key := generatedFileKey{
		GoPackage: "dialogflow/v1alpha1",
		FileName:  "types.generated.go",
	}

	// Step 1 writes Dialogflow v2 types to types.generated.go.
	var gb1 generatorBase
	gb1.init(baseDir)
	f1 := gb1.getOutputFile(key)
	f1.goPackage = "v1alpha1"
	f1.fileAnnotation = &annotations.FileAnnotation{
		Key: "+generated:types",
		Attributes: map[string][]string{
			"krm.group":     {"dialogflow.cnrm.cloud.google.com"},
			"krm.version":   {"v1alpha1"},
			"proto.service": {"google.cloud.dialogflow.v2"},
			"resource":      {"DialogflowGenerator:Generator"},
		},
	}
	f1.body.WriteString(`
// +kcc:proto=google.cloud.dialogflow.v2.ConversationContext
type ConversationContext struct {
}
`)
	if err := f1.Write(true, true, true); err != nil {
		t.Fatalf("f1.Write failed: %v", err)
	}

	// Step 2 regenerates DialogflowSecuritySettings:SecuritySettings (which generate.sh will mv to securitysettings_types.generated.go).
	// It must NOT merge ConversationContext from types.generated.go.
	var gb2 generatorBase
	gb2.init(baseDir)
	f2 := gb2.getOutputFile(key)
	f2.goPackage = "v1alpha1"
	f2.fileAnnotation = &annotations.FileAnnotation{
		Key: "+generated:types",
		Attributes: map[string][]string{
			"krm.group":     {"dialogflow.cnrm.cloud.google.com"},
			"krm.version":   {"v1alpha1"},
			"proto.service": {"google.cloud.dialogflow.cx.v3"},
			"resource":      {"DialogflowSecuritySettings:SecuritySettings"},
		},
	}
	f2.body.WriteString(`
// +kcc:proto=google.cloud.dialogflow.cx.v3.SecuritySettings.AudioExportSettings
type SecuritySettings_AudioExportSettings struct {
}
`)
	if err := f2.Write(true, true, true); err != nil {
		t.Fatalf("f2.Write failed: %v", err)
	}

	gotBytes, err := os.ReadFile(filepath.Join(pkgDir, "types.generated.go"))
	if err != nil {
		t.Fatalf("reading types.generated.go: %v", err)
	}
	got := string(gotBytes)
	if strings.Contains(got, "ConversationContext") {
		t.Errorf("expected separate *_types.generated.go target not to merge ConversationContext, got:\n%s", got)
	}
}

func TestGeneratedFileWritePreservesActiveTypeOverCommentedDifferentProto(t *testing.T) {
	baseDir := t.TempDir()
	key := generatedFileKey{
		GoPackage: "foo/v1alpha1",
		FileName:  "types.generated.go",
	}

	var gb1 generatorBase
	gb1.init(baseDir)
	f1 := gb1.getOutputFile(key)
	f1.goPackage = "v1alpha1"
	f1.body.WriteString(`
// +kcc:proto=google.cloud.foo.v1.sub1.EncryptionConfig
type EncryptionConfig struct {
	KmsKey *string ` + "`json:\"kmsKey,omitempty\"`" + `
}
`)
	if err := f1.Write(true, true, true); err != nil {
		t.Fatalf("f1.Write failed: %v", err)
	}

	// Second service emits a commented-out EncryptionConfig from a DIFFERENT proto message.
	var gb2 generatorBase
	gb2.init(baseDir)
	f2 := gb2.getOutputFile(key)
	f2.goPackage = "v1alpha1"
	f2.body.WriteString(`
/* unreachable type:
// +kcc:proto=google.cloud.foo.v1.sub2.EncryptionConfig
type EncryptionConfig struct {
	OtherField *string ` + "`json:\"otherField,omitempty\"`" + `
}
*/

// +kcc:proto=google.cloud.foo.v1.sub2.Sub2Config
type Sub2Config struct {
}
`)
	if err := f2.Write(true, true, true); err != nil {
		t.Fatalf("f2.Write failed: %v", err)
	}

	gotBytes, err := os.ReadFile(filepath.Join(baseDir, "foo", "v1alpha1", "types.generated.go"))
	if err != nil {
		t.Fatalf("reading types.generated.go: %v", err)
	}
	got := string(gotBytes)
	if !strings.Contains(got, "// +kcc:proto=google.cloud.foo.v1.sub1.EncryptionConfig\ntype EncryptionConfig struct") {
		t.Errorf("expected active EncryptionConfig from sub1 to be preserved over commented sub2 block, got:\n%s", got)
	}
	if !strings.Contains(got, "type Sub2Config struct") {
		t.Errorf("expected Sub2Config to be present, got:\n%s", got)
	}
}

func TestGeneratedFileWriteAutoDiscoversValidKRMTypesFromAPIDir(t *testing.T) {
	baseDir := t.TempDir()
	apiDir := filepath.Join(baseDir, "apis", "foo", "v1alpha1")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	// Only KeptSpec is an active type; PrunedSpec is commented out.
	if err := os.WriteFile(filepath.Join(apiDir, "types.generated.go"), []byte(`package v1alpha1

type KeptSpec struct{}
type CurrentSpec struct{}

/* unreachable type:
type PrunedSpec struct{}
*/
`), 0644); err != nil {
		t.Fatalf("writing types.generated.go: %v", err)
	}

	key := generatedFileKey{
		GoPackage: "pkg/controller/direct/foo",
		FileName:  "mapper.generated.go",
	}

	var gb1 generatorBase
	gb1.init(baseDir)
	f1 := gb1.getOutputFile(key)
	f1.goPackage = "foo"
	f1.addImport("krm", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1")
	f1.body.WriteString(`
func KeptSpec_FromProto(mapCtx *direct.MapContext, in * string) *krm.KeptSpec {
	return nil
}

func PrunedSpec_FromProto(mapCtx *direct.MapContext, in *string) *krm.PrunedSpec {
	return nil
}
`)
	if err := f1.Write(true, true, false); err != nil {
		t.Fatalf("f1.Write failed: %v", err)
	}

	// Second run leaves f2.validKRMTypes nil so discoverValidKRMTypes scans apis/foo/v1alpha1.
	var gb2 generatorBase
	gb2.init(baseDir)
	f2 := gb2.getOutputFile(key)
	f2.goPackage = "foo"
	f2.addImport("krm", "github.com/GoogleCloudPlatform/k8s-config-connector/apis/foo/v1alpha1")
	f2.body.WriteString(`
func CurrentSpec_FromProto(mapCtx *direct.MapContext, in *string) *krm.CurrentSpec {
	return nil
}
`)
	if err := f2.Write(true, true, false); err != nil {
		t.Fatalf("f2.Write failed: %v", err)
	}

	gotBytes, err := os.ReadFile(filepath.Join(baseDir, "pkg", "controller", "direct", "foo", "mapper.generated.go"))
	if err != nil {
		t.Fatalf("reading mapper.generated.go: %v", err)
	}
	got := string(gotBytes)
	if !strings.Contains(got, "func KeptSpec_FromProto") {
		t.Errorf("expected KeptSpec_FromProto to be preserved, got:\n%s", got)
	}
	if strings.Contains(got, "func PrunedSpec_FromProto") {
		t.Errorf("expected PrunedSpec_FromProto to be pruned via auto-discovered validKRMTypes, got:\n%s", got)
	}
	if !strings.Contains(got, "func CurrentSpec_FromProto") {
		t.Errorf("expected CurrentSpec_FromProto to be present, got:\n%s", got)
	}
}

func TestGeneratedFileWriteDisambiguatesCollidingImportAliases(t *testing.T) {
	baseDir := t.TempDir()
	key := generatedFileKey{
		GoPackage: "pkg/controller/direct/foo",
		FileName:  "mapper.generated.go",
	}

	var gb1 generatorBase
	gb1.init(baseDir)
	f1 := gb1.getOutputFile(key)
	f1.goPackage = "foo"
	f1.addImport("pb", "cloud.google.com/go/foo/apiv1/foopb")
	f1.body.WriteString(`
func Alpha_FromProto(in *pb.Alpha) *pb.Alpha {
	out := &pb.Alpha{}
	return out
}
`)
	if err := f1.Write(true, true, true); err != nil {
		t.Fatalf("f1.Write failed: %v", err)
	}

	// Second run binds "pb" to a different proto package.
	var gb2 generatorBase
	gb2.init(baseDir)
	f2 := gb2.getOutputFile(key)
	f2.goPackage = "foo"
	f2.addImport("pb", "cloud.google.com/go/foo/sub/apiv1/subpb")
	f2.body.WriteString(`
func Beta_FromProto(in *pb.Beta) *pb.Beta {
	out := &pb.Beta{}
	return out
}
`)
	if err := f2.Write(true, true, true); err != nil {
		t.Fatalf("f2.Write failed: %v", err)
	}

	gotBytes, err := os.ReadFile(filepath.Join(baseDir, "pkg", "controller", "direct", "foo", "mapper.generated.go"))
	if err != nil {
		t.Fatalf("reading mapper.generated.go: %v", err)
	}
	got := string(gotBytes)
	if !strings.Contains(got, `pb_2 "cloud.google.com/go/foo/apiv1/foopb"`) {
		t.Errorf("expected disambiguated import alias pb_2 for foopb, got:\n%s", got)
	}
	if !strings.Contains(got, `pb "cloud.google.com/go/foo/sub/apiv1/subpb"`) {
		t.Errorf("expected import alias pb for subpb, got:\n%s", got)
	}
	if !strings.Contains(got, "func Alpha_FromProto(in *pb_2.Alpha) *pb_2.Alpha") {
		t.Errorf("expected Alpha_FromProto to have pb rewritten to pb_2, got:\n%s", got)
	}
	if !strings.Contains(got, "func Beta_FromProto(in *pb.Beta) *pb.Beta") {
		t.Errorf("expected Beta_FromProto to keep pb.Beta, got:\n%s", got)
	}
}
