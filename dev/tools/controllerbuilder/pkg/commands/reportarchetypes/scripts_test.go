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

package reportarchetypes

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseScript(t *testing.T) {
	// Arrange
	script := `#!/bin/bash
set -o errexit
CONTROLLERBUILDER="go run ${REPO_ROOT}/dev/tools/controllerbuilder"
cd ${REPO_ROOT}/dev/tools/controllerbuilder
./generate-proto.sh

# --- v1alpha1 ---
${CONTROLLERBUILDER} generate-types \
  --service google.cloud.foo.v1 \
  --api-version foo.cnrm.cloud.google.com/v1alpha1 \
  --resource FooWidget:Widget \
  --resource FooGadget:Gadget \
  --skip-scaffold-files

# A generate-mapper invocation maps no kinds.
${CONTROLLERBUILDER} generate-mapper \
  --service google.cloud.foo.v1 \
  --api-version foo.cnrm.cloud.google.com/v1alpha1

go run . generate-types --service=google.cloud.foo.v1beta,google.cloud.foo.v1 \
  -v "foo.cnrm.cloud.google.com/v1beta1" --resource='FooWidget:Widget'; echo done

# ${CONTROLLERBUILDER} generate-types --service google.cloud.old.v1 \
  --api-version old.cnrm.cloud.google.com/v1alpha1 --resource Old:Old
`

	// Act
	got, err := parseScript(script, "foo/generate.sh", "/repo")

	// Assert
	if err != nil {
		t.Fatalf("parseScript: %v", err)
	}
	want := []Mapping{
		{Kind: "FooWidget", APIVersion: "foo.cnrm.cloud.google.com/v1alpha1", Services: []string{"google.cloud.foo.v1"}, ProtoName: "Widget", Source: "foo/generate.sh:8"},
		{Kind: "FooGadget", APIVersion: "foo.cnrm.cloud.google.com/v1alpha1", Services: []string{"google.cloud.foo.v1"}, ProtoName: "Gadget", Source: "foo/generate.sh:8"},
		{Kind: "FooWidget", APIVersion: "foo.cnrm.cloud.google.com/v1beta1", Services: []string{"google.cloud.foo.v1beta", "google.cloud.foo.v1"}, ProtoName: "Widget", Source: "foo/generate.sh:20"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parseScript mismatch (-want +got):\n%s", diff)
	}
}

// A commented-out line ends the command it sits in, as it does in bash, so
// the flags after it belong to a command that is not generate-types.
func TestParseScriptCommentEndsCommand(t *testing.T) {
	// Arrange
	script := `${CONTROLLERBUILDER} generate-types \
  --service google.cloud.foo.v1 \
  --api-version foo.cnrm.cloud.google.com/v1alpha1 \
  --resource FooWidget:Widget \
# --resource FooGadget:Gadget \
  --resource FooGizmo:Gizmo
`

	// Act
	got, err := parseScript(script, "foo/generate.sh", "/repo")

	// Assert
	if err != nil {
		t.Fatalf("parseScript: %v", err)
	}
	var kinds []string
	for _, m := range got {
		kinds = append(kinds, m.Kind)
	}
	if diff := cmp.Diff([]string{"FooWidget"}, kinds); diff != "" {
		t.Errorf("kinds mismatch (-want +got):\n%s", diff)
	}
}

func TestParseScriptConfig(t *testing.T) {
	// Arrange
	repoRoot := t.TempDir()
	configDir := filepath.Join(repoRoot, "apis", "foo", "v1beta1")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		t.Fatal(err)
	}
	config := `service: google.cloud.foo.v1
apiVersion: foo.cnrm.cloud.google.com/v1beta1
resources:
  - kind: FooWidget
    protoName: Widget
  - kind: FooGadget
    protoName: Gadget
`
	if err := os.WriteFile(filepath.Join(configDir, "generatetypes.yaml"), []byte(config), 0o644); err != nil {
		t.Fatal(err)
	}
	script := `${CONTROLLERBUILDER} generate-types \
    --config ${REPO_ROOT}/apis/foo/v1beta1/generatetypes.yaml
`

	// Act
	got, err := parseScript(script, "foo/generate.sh", repoRoot)

	// Assert
	if err != nil {
		t.Fatalf("parseScript: %v", err)
	}
	want := []Mapping{
		{Kind: "FooWidget", APIVersion: "foo.cnrm.cloud.google.com/v1beta1", Services: []string{"google.cloud.foo.v1"}, ProtoName: "Widget", Source: "foo/generate.sh:1"},
		{Kind: "FooGadget", APIVersion: "foo.cnrm.cloud.google.com/v1beta1", Services: []string{"google.cloud.foo.v1"}, ProtoName: "Gadget", Source: "foo/generate.sh:1"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parseScript mismatch (-want +got):\n%s", diff)
	}
}

func TestParseScriptProtoSource(t *testing.T) {
	// Arrange
	script := `PROTO_SHA="0123abcd"
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"
./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}
x generate-types -s google.cloud.foo.v1 -v foo.cnrm.cloud.google.com/v1alpha1 --resource Pinned:Pinned
x generate-types -s google.cloud.foo.v1 -v foo.cnrm.cloud.google.com/v1alpha1 --resource Variable:Variable \
  --proto-source-path ${PROTO_OUT}
x generate-types -s google.cloud.foo.v1 -v foo.cnrm.cloud.google.com/v1alpha1 --resource Literal:Literal \
  --proto-source-path "${REPO_ROOT}/.build/googleapis-foo.pb"
./generate-proto.sh HEAD
x generate-types -s google.cloud.foo.v1 -v foo.cnrm.cloud.google.com/v1alpha1 --resource Head:Head
x generate-types -s google.cloud.foo.v1 -v foo.cnrm.cloud.google.com/v1alpha1 --resource Unassigned:Unassigned \
  --proto-source-path=$UNASSIGNED
`

	// Act
	got, err := parseScript(script, "foo/generate.sh", "/repo")

	// Assert
	if err != nil {
		t.Fatalf("parseScript: %v", err)
	}
	var sources [][2]string
	for _, m := range got {
		sources = append(sources, [2]string{m.Kind, m.ProtoSource})
	}
	want := [][2]string{
		// generate-proto.sh with an output path leaves the default alone.
		{"Pinned", ""},
		{"Variable", "googleapis-${PROTO_SHA}.pb"},
		{"Literal", "googleapis-foo.pb"},
		{"Head", "HEAD"},
		{"Unassigned", "$UNASSIGNED"},
	}
	if diff := cmp.Diff(want, sources); diff != "" {
		t.Errorf("ProtoSource mismatch (-want +got):\n%s", diff)
	}
}

func TestParseScriptErrors(t *testing.T) {
	for _, tc := range []struct {
		name    string
		script  string
		wantErr string
	}{
		{
			name:    "no service",
			script:  "x generate-types --api-version foo.cnrm.cloud.google.com/v1alpha1 --resource A:B\n",
			wantErr: "foo/generate.sh:1: generate-types has no --service",
		},
		{
			name:    "no api version",
			script:  "\nx generate-types --service google.cloud.foo.v1 --resource A:B\n",
			wantErr: "foo/generate.sh:2: generate-types has no --api-version",
		},
		{
			name:    "no resource",
			script:  "x generate-types --service google.cloud.foo.v1 --api-version foo.cnrm.cloud.google.com/v1alpha1\n",
			wantErr: "foo/generate.sh:1: generate-types has no --resource",
		},
		{
			name:    "malformed resource",
			script:  "x generate-types --service google.cloud.foo.v1 --api-version foo.cnrm.cloud.google.com/v1alpha1 --resource Widget\n",
			wantErr: `expected [KRMKind]:[ProtoResourceName], got "Widget"`,
		},
		{
			name:    "flag without a value",
			script:  "x generate-types --service google.cloud.foo.v1 --resource\n",
			wantErr: "flag --resource has no value",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			_, err := parseScript(tc.script, "foo/generate.sh", "/repo")

			// Assert
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("parseScript error = %v, want one containing %q", err, tc.wantErr)
			}
		})
	}
}

func TestShellCommands(t *testing.T) {
	// Arrange
	script := "a 'b c' \"d \\\"e\\\"\" f\\ g # comment \\\nh; i && j | k\nl \\\n  m#n\n"

	// Act
	got := shellCommands(script)

	// Assert
	want := []shellCommand{
		{words: []string{"a", "b c", `d "e"`, "f g"}, line: 1},
		{words: []string{"h"}, line: 2},
		{words: []string{"i"}, line: 2},
		{words: []string{"j"}, line: 2},
		{words: []string{"k"}, line: 2},
		{words: []string{"l", "m#n"}, line: 3},
	}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(shellCommand{})); diff != "" {
		t.Errorf("shellCommands mismatch (-want +got):\n%s", diff)
	}
}

func TestScanGenerateScripts(t *testing.T) {
	// Arrange
	apisDir := filepath.Join(t.TempDir(), "apis")
	scripts := map[string]string{
		"bar/generate.sh": "x generate-types --service google.cloud.bar.v1 --api-version bar.cnrm.cloud.google.com/v1beta1 --resource BarThing:Thing\n",
		"foo/generate.sh": "x generate-types --service google.cloud.foo.v1 --api-version foo.cnrm.cloud.google.com/v1alpha1 --resource FooThing:Thing\n",
		"foo/other.sh":    "x generate-types --service google.cloud.other.v1 --api-version other.cnrm.cloud.google.com/v1alpha1 --resource Other:Other\n",
	}
	for name, content := range scripts {
		path := filepath.Join(apisDir, name)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	// Act
	got, err := ScanGenerateScripts(apisDir)

	// Assert
	if err != nil {
		t.Fatalf("ScanGenerateScripts: %v", err)
	}
	var sources []string
	for _, m := range got {
		sources = append(sources, m.Kind+" "+m.Source)
	}
	if diff := cmp.Diff([]string{"BarThing bar/generate.sh:1", "FooThing foo/generate.sh:1"}, sources); diff != "" {
		t.Errorf("ScanGenerateScripts mismatch (-want +got):\n%s", diff)
	}
}
