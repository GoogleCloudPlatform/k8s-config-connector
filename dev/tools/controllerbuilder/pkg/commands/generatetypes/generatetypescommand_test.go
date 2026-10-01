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
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"
)

// generate.sh calls generate-types once per proto version, so the queue has to
// survive several calls for the same service.
func TestWriteJudgementQueueKeepsEarlierCalls(t *testing.T) {
	// Arrange
	apiDir := t.TempDir()
	goPackage := "discoveryengine/v1alpha1"
	path := filepath.Join(apiDir, "discoveryengine", judgement.FileName)
	a := openEntry("KindA", ".spec.a", "x")
	b := openEntry("KindB", ".spec.b", "y")

	// Act
	if err := writeJudgementQueue(apiDir, goPackage, []judgement.Entry{a}); err != nil {
		t.Fatalf("first call: %v", err)
	}
	if err := writeJudgementQueue(apiDir, goPackage, []judgement.Entry{b}); err != nil {
		t.Fatalf("second call: %v", err)
	}
	afterTwo := readQueueFile(t, path)
	if err := writeJudgementQueue(apiDir, goPackage, []judgement.Entry{b}); err != nil {
		t.Fatalf("third call: %v", err)
	}
	afterThree := readQueueFile(t, path)

	// Assert
	q, err := judgement.Read(path)
	if err != nil {
		t.Fatalf("reading queue: %v", err)
	}
	if want := []judgement.Entry{a, b}; !reflect.DeepEqual(q.Entries, want) {
		t.Errorf("entries = %+v, want %+v", q.Entries, want)
	}
	if strings.Count(afterTwo, "# Judgement queue") != 1 {
		t.Errorf("want the header exactly once, got:\n%s", afterTwo)
	}
	if afterThree != afterTwo {
		t.Errorf("repeating a call changed the queue:\nbefore:\n%s\nafter:\n%s", afterTwo, afterThree)
	}
}

// A person resolves an entry by editing the file. Regenerating must keep that.
func TestWriteJudgementQueueKeepsResolutions(t *testing.T) {
	// Arrange
	apiDir := t.TempDir()
	goPackage := "discoveryengine/v1alpha1"
	path := filepath.Join(apiDir, "discoveryengine", judgement.FileName)
	a := openEntry("KindA", ".spec.a", "possible-reference")
	if err := writeJudgementQueue(apiDir, goPackage, []judgement.Entry{a}); err != nil {
		t.Fatalf("first call: %v", err)
	}
	q, err := judgement.Read(path)
	if err != nil {
		t.Fatalf("reading queue: %v", err)
	}
	q.Entries[0].Status = judgement.StatusResolved
	q.Entries[0].Resolution = judgement.ResolutionEdited
	q.Entries[0].Note = "changed to aRef"
	if err := judgement.Write(path, q); err != nil {
		t.Fatalf("writing the resolution: %v", err)
	}

	// Act
	if err := writeJudgementQueue(apiDir, goPackage, []judgement.Entry{a}); err != nil {
		t.Fatalf("regenerating: %v", err)
	}

	// Assert
	got, err := judgement.Read(path)
	if err != nil {
		t.Fatalf("reading queue: %v", err)
	}
	if len(got.Entries) != 1 || got.Entries[0].Status != judgement.StatusResolved || got.Entries[0].Note != "changed to aRef" {
		t.Errorf("regenerating reopened or lost the resolution: %+v", got.Entries)
	}
}

// A queue that does not parse may hold resolutions someone typed by hand, so
// the generator stops rather than overwriting it.
func TestWriteJudgementQueueRejectsABrokenFile(t *testing.T) {
	apiDir := t.TempDir()
	path := filepath.Join(apiDir, "discoveryengine", judgement.FileName)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	broken := "entries:\n- kind: KindA\n  reason: x\n  status: open\n"
	if err := os.WriteFile(path, []byte(broken), 0644); err != nil {
		t.Fatal(err)
	}

	err := writeJudgementQueue(apiDir, "discoveryengine/v1alpha1", []judgement.Entry{openEntry("KindB", ".spec.b", "y")})

	if err == nil {
		t.Fatal("expected an error for a queue that does not validate")
	}
	if got := readQueueFile(t, path); got != broken {
		t.Errorf("the broken file was overwritten:\n%s", got)
	}
}

func readQueueFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading queue: %v", err)
	}
	return string(b)
}

func openEntry(kind, field, reason string) judgement.Entry {
	return judgement.Entry{Kind: kind, Group: "discoveryengine.cnrm.cloud.google.com", Field: field, Reason: reason, Status: judgement.StatusOpen}
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

// generate-types only checks output-only comments while prepopulating the
// Spec, so --detect-output-only-in-comments on its own is rejected rather than
// ignored.
func TestValidateDetectOutputOnlyNeedsPrepopulateSpec(t *testing.T) {
	for _, tc := range []struct {
		name             string
		prepopulateSpec  bool
		detectOutputOnly bool
		wantErr          string
	}{
		{
			name: "neither flag",
		},
		{
			name:            "only --prepopulate-spec",
			prepopulateSpec: true,
		},
		{
			name:             "both flags",
			prepopulateSpec:  true,
			detectOutputOnly: true,
		},
		{
			name:             "only --detect-output-only-in-comments",
			detectOutputOnly: true,
			wantErr:          "`--detect-output-only-in-comments` requires `--prepopulate-spec`",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			o := &GenerateCRDOptions{
				GenerateOptions:  &options.GenerateOptions{ProtoSourcePath: "googleapis.pb"},
				ServiceName:      "google.cloud.example.v1",
				Resources:        options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
				PrepopulateSpec:  tc.prepopulateSpec,
				DetectOutputOnly: tc.detectOutputOnly,
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

// generate-types only moves server-set fields out of a prepopulated Spec, so
// --place-server-set-fields on its own is rejected rather than ignored.
func TestValidatePlaceServerSetFieldsNeedsPrepopulateSpec(t *testing.T) {
	for _, tc := range []struct {
		name                 string
		prepopulateSpec      bool
		placeServerSetFields bool
		wantErr              string
	}{
		{
			name: "neither flag",
		},
		{
			name:            "only --prepopulate-spec",
			prepopulateSpec: true,
		},
		{
			name:                 "both flags",
			prepopulateSpec:      true,
			placeServerSetFields: true,
		},
		{
			name:                 "only --place-server-set-fields",
			placeServerSetFields: true,
			wantErr:              "`--place-server-set-fields` requires `--prepopulate-spec`",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			o := &GenerateCRDOptions{
				GenerateOptions:      &options.GenerateOptions{ProtoSourcePath: "googleapis.pb"},
				ServiceName:          "google.cloud.example.v1",
				Resources:            options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
				PrepopulateSpec:      tc.prepopulateSpec,
				PlaceServerSetFields: tc.placeServerSetFields,
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

// A bare --resource name that exists in more than one listed service gets a
// resource-level queue entry naming every match and the one that was used.
func TestAmbiguousResourceItem(t *testing.T) {
	resource := options.Resource{Kind: "NotebookInstanceV2", ProtoName: "Instance"}
	for _, tc := range []struct {
		name    string
		matches []string
		want    *scaffold.JudgementItem
	}{
		{
			name: "qualified name or no match",
		},
		{
			name:    "one match",
			matches: []string{"google.cloud.notebooks.v2.Instance"},
		},
		{
			name:    "two matches",
			matches: []string{"google.cloud.notebooks.v1.Instance", "google.cloud.notebooks.v2.Instance"},
			want: &scaffold.JudgementItem{
				Reason: "ambiguous-proto-name",
				Detail: "Instance exists in 2 of the listed services (google.cloud.notebooks.v1.Instance, google.cloud.notebooks.v2.Instance); " +
					"generated from google.cloud.notebooks.v1.Instance. Put the full proto name in --resource to choose one",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			item, ok := ambiguousResourceItem(resource, tc.matches)

			// Assert
			if ok != (tc.want != nil) {
				t.Fatalf("ok = %v, want %v", ok, tc.want != nil)
			}
			if ok && item != *tc.want {
				t.Errorf("item:\n got %+v\nwant %+v", item, *tc.want)
			}
		})
	}
}
