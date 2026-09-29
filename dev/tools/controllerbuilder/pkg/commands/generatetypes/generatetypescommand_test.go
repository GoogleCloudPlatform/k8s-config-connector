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

package generatetypes

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
)

func TestMergeQueueLines(t *testing.T) {
	header := "# header line one\n#\n# header line two\n\n"
	for _, tc := range []struct {
		name     string
		existing string
		added    string
		want     []string
	}{
		{
			name:     "existing lines come first, then new ones",
			existing: header + "KindA .spec.a reason=x\n",
			added:    "KindB .spec.b reason=y\n",
			want:     []string{"KindA .spec.a reason=x", "KindB .spec.b reason=y"},
		},
		{
			name:     "a line already present is not repeated",
			existing: header + "KindA .spec.a reason=x\n",
			added:    "KindA .spec.a reason=x\nKindB .spec.b reason=y\n",
			want:     []string{"KindA .spec.a reason=x", "KindB .spec.b reason=y"},
		},
		{
			name:     "header lines are dropped and other comments kept",
			existing: header + "# dropped: google.cloud.x.v1.Msg.field reason=z\n",
			added:    "",
			want:     []string{"# dropped: google.cloud.x.v1.Msg.field reason=z"},
		},
		{
			name:     "derived message map comment lines are preserved",
			existing: header + "# derived-message-map: google.cloud.x.v1.Msg.field proto=google.cloud.x.v1.Sub goType=Sub\n",
			added:    "",
			want:     []string{"# derived-message-map: google.cloud.x.v1.Msg.field proto=google.cloud.x.v1.Sub goType=Sub"},
		},
		{
			name:     "no existing file",
			existing: "",
			added:    "KindA .spec.a reason=x\n\n",
			want:     []string{"KindA .spec.a reason=x"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := mergeQueueLines(header, tc.existing, tc.added)

			// Assert
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

// generate.sh calls generate-types once per proto version, so the queue has to
// survive several calls for the same service.
func TestWriteJudgementQueueKeepsEarlierCalls(t *testing.T) {
	// Arrange
	apiDir := t.TempDir()
	goPackage := "discoveryengine/v1alpha1"
	path := filepath.Join(apiDir, "discoveryengine", "needs_judgement_call.txt")

	// Act
	if err := writeJudgementQueue(apiDir, goPackage, []string{"KindA .spec.a reason=x\n"}); err != nil {
		t.Fatalf("first call: %v", err)
	}
	if err := writeJudgementQueue(apiDir, goPackage, []string{"KindB .spec.b reason=y\n"}); err != nil {
		t.Fatalf("second call: %v", err)
	}
	afterTwo := readQueue(t, path)
	if err := writeJudgementQueue(apiDir, goPackage, []string{"KindB .spec.b reason=y\n"}); err != nil {
		t.Fatalf("third call: %v", err)
	}
	afterThree := readQueue(t, path)

	// Assert
	for _, entry := range []string{"KindA .spec.a reason=x", "KindB .spec.b reason=y"} {
		if strings.Count(afterTwo, entry) != 1 {
			t.Errorf("want %q exactly once after two calls, got:\n%s", entry, afterTwo)
		}
	}
	if strings.Count(afterTwo, "# Fields emitted during generation that require human review") != 1 {
		t.Errorf("want the header exactly once, got:\n%s", afterTwo)
	}
	if afterThree != afterTwo {
		t.Errorf("repeating a call changed the queue:\nbefore:\n%s\nafter:\n%s", afterTwo, afterThree)
	}
}

func readQueue(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading queue: %v", err)
	}
	return string(b)
}

// generate-types only adds a parent reference to a prepopulated Spec, so
// --emit-parent-refs on its own is rejected rather than ignored.
func TestValidateEmitParentRefsNeedsPrepopulateSpec(t *testing.T) {
	for _, tc := range []struct {
		name            string
		prepopulateSpec bool
		emitParentRefs  bool
		wantErr         string
	}{
		{
			name: "neither flag",
		},
		{
			name:            "only --prepopulate-spec",
			prepopulateSpec: true,
		},
		{
			name:            "both flags",
			prepopulateSpec: true,
			emitParentRefs:  true,
		},
		{
			name:           "only --emit-parent-refs",
			emitParentRefs: true,
			wantErr:        "`--emit-parent-refs` requires `--prepopulate-spec`",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			o := &GenerateCRDOptions{
				GenerateOptions: &options.GenerateOptions{ProtoSourcePath: "googleapis.pb"},
				ServiceName:     "google.cloud.example.v1",
				Resources:       options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
				PrepopulateSpec: tc.prepopulateSpec,
				EmitParentRefs:  tc.emitParentRefs,
			}

			// Act
			err := o.validate()

			// Assert
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}
			if gotErr != tc.wantErr {
				t.Errorf("validate() error = %q, want %q", gotErr, tc.wantErr)
			}
		})
	}
}

// generate-types only collects reference hints while prepopulating the Spec,
// so --emit-reference-hints on its own is rejected rather than ignored.
func TestValidateEmitReferenceHintsNeedsPrepopulateSpec(t *testing.T) {
	for _, tc := range []struct {
		name               string
		prepopulateSpec    bool
		emitReferenceHints bool
		wantErr            string
	}{
		{
			name: "neither flag",
		},
		{
			name:            "only --prepopulate-spec",
			prepopulateSpec: true,
		},
		{
			name:               "both flags",
			prepopulateSpec:    true,
			emitReferenceHints: true,
		},
		{
			name:               "only --emit-reference-hints",
			emitReferenceHints: true,
			wantErr:            "`--emit-reference-hints` requires `--prepopulate-spec`",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			o := &GenerateCRDOptions{
				GenerateOptions:    &options.GenerateOptions{ProtoSourcePath: "googleapis.pb"},
				ServiceName:        "google.cloud.example.v1",
				Resources:          options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
				PrepopulateSpec:    tc.prepopulateSpec,
				EmitReferenceHints: tc.emitReferenceHints,
			}

			// Act
			err := o.validate()

			// Assert
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}
			if gotErr != tc.wantErr {
				t.Errorf("validate() error = %q, want %q", gotErr, tc.wantErr)
			}
		})
	}
}
