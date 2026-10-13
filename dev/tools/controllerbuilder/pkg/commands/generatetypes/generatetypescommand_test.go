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
	"context"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/scaffold"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
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

// TestRunGenerateCRDChecksUntriagedKindAgain covers a Kind whose types file
// exists and whose untriaged-bulk-generation entry is still open. The
// per-Kind checks run again and queue what they find: here, the network hint
// and the empty ObservedState. The untriaged entry gets the current detail.
// The types file does not change, and entries that only scaffolding queues,
// such as location-parent-unknown, are not added.
func TestRunGenerateCRDChecksUntriagedKindAgain(t *testing.T) {
	// Arrange
	o, typesPath, queuePath := widgetService(t, `entries:
- kind: ExampleWidget
  group: example.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: open
`)
	want := []judgement.Entry{
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Reason: "untriaged-bulk-generation",
			Detail: "spec was generated from proto definition; verify refs, omissions, and KRM conventions",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Field:  ".spec.network",
			Reason: "possible-reference-by-description",
			Detail: `the description has the resource-name template "projects/{project}/global/networks/{network}"`,
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Reason: "empty-observedstate",
			Detail: "nothing was generated into status.observedState; the proto probably marks no field OUTPUT_ONLY, so output fields landed in the Spec. Decide what belongs in status",
			Status: judgement.StatusOpen,
		},
	}

	// Act
	err := RunGenerateCRD(context.Background(), o)

	// Assert
	if err != nil {
		t.Fatalf("RunGenerateCRD() error: %v", err)
	}
	if diff := cmp.Diff(widgetTypesFile, readFile(t, typesPath)); diff != "" {
		t.Errorf("RunGenerateCRD() changed the types file (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(want, readQueueEntries(t, queuePath)); diff != "" {
		t.Errorf("queue entries mismatch (-want +got):\n%s", diff)
	}
}

// TestRunGenerateCRDLeavesReviewedKindAlone checks that a Kind whose
// untriaged-bulk-generation entry is resolved is not checked again, so the
// network hint is not queued.
func TestRunGenerateCRDLeavesReviewedKindAlone(t *testing.T) {
	// Arrange
	o, _, queuePath := widgetService(t, `entries:
- kind: ExampleWidget
  group: example.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: resolved
  resolution: accepted
`)
	want := []judgement.Entry{
		{
			Kind:       "ExampleWidget",
			Group:      "example.cnrm.cloud.google.com",
			Reason:     "untriaged-bulk-generation",
			Status:     judgement.StatusResolved,
			Resolution: judgement.ResolutionAccepted,
		},
	}

	// Act
	err := RunGenerateCRD(context.Background(), o)

	// Assert
	if err != nil {
		t.Fatalf("RunGenerateCRD() error: %v", err)
	}
	if diff := cmp.Diff(want, readQueueEntries(t, queuePath)); diff != "" {
		t.Errorf("queue entries mismatch (-want +got):\n%s", diff)
	}
}

// TestRunGenerateCRDChecksUntriagedKindOncePerEntry checks that re-checking an
// untriaged Kind adds no duplicates: a second run leaves the queue file
// exactly as the first run wrote it.
func TestRunGenerateCRDChecksUntriagedKindOncePerEntry(t *testing.T) {
	// Arrange
	o, _, queuePath := widgetService(t, `entries:
- kind: ExampleWidget
  group: example.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: open
`)

	// Act
	if err := RunGenerateCRD(context.Background(), o); err != nil {
		t.Fatalf("first run: %v", err)
	}
	afterFirst := readFile(t, queuePath)
	if err := RunGenerateCRD(context.Background(), o); err != nil {
		t.Fatalf("second run: %v", err)
	}
	afterSecond := readFile(t, queuePath)

	// Assert
	if diff := cmp.Diff(afterFirst, afterSecond); diff != "" {
		t.Errorf("the second run changed the queue file (-first +second):\n%s", diff)
	}
}

// widgetTypesFile stands in for an ExampleWidget types file that someone has
// edited since it was scaffolded.
const widgetTypesFile = `package v1alpha1

// ExampleWidgetSpec was edited by hand.
type ExampleWidgetSpec struct{}
`

// widgetService sets up a service with one Kind, ExampleWidget, whose types
// file already exists, plus the given queue file. It returns the options
// for a run with --prepopulate-spec and --emit-reference-hints, and the paths
// of the types file and the queue file.
func widgetService(t *testing.T, queue string) (o *GenerateCRDOptions, typesPath, queuePath string) {
	t.Helper()
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "apis")
	typesPath = filepath.Join(apiDir, "example", "v1alpha1", "examplewidget_types.go")
	queuePath = filepath.Join(apiDir, "example", judgement.FileName)
	writeFile(t, typesPath, widgetTypesFile)
	writeFile(t, queuePath, queue)
	o = &GenerateCRDOptions{
		GenerateOptions: &options.GenerateOptions{
			ProtoSourcePath: writeWidgetProto(t, dir),
			APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
		},
		ServiceName:        "google.cloud.example.v1",
		OutputAPIDirectory: apiDir,
		Resources:          options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
		PrepopulateSpec:    true,
		EmitReferenceHints: true,
	}
	return o, typesPath, queuePath
}

// writeWidgetProto writes a descriptor set holding google.cloud.example.v1.Widget
// to dir and returns its path. Widget has no OUTPUT_ONLY field, and the
// comment on network spells out a resource-name template.
func writeWidgetProto(t *testing.T, dir string) string {
	t.Helper()
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{{
		Name:    proto.String("google/cloud/example/v1/widget.proto"),
		Package: proto.String("google.cloud.example.v1"),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: proto.String("Widget"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{Name: proto.String("name"), Number: proto.Int32(1), Type: str},
				{Name: proto.String("network"), Number: proto.Int32(2), Type: str},
			},
		}},
		SourceCodeInfo: &descriptorpb.SourceCodeInfo{Location: []*descriptorpb.SourceCodeInfo_Location{{
			// [4=message_type 0, 2=field 1] is Widget.network.
			Path:            []int32{4, 0, 2, 1},
			Span:            []int32{0, 0, 1},
			LeadingComments: proto.String(" The network to join. Format: projects/{project}/global/networks/{network}\n"),
		}}},
	}}}
	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	path := filepath.Join(dir, "widget.pb")
	writeFile(t, path, string(b))
	return path
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func readQueueEntries(t *testing.T, path string) []judgement.Entry {
	t.Helper()
	q, err := judgement.Read(path)
	if err != nil {
		t.Fatalf("reading queue: %v", err)
	}
	return q.Entries
}
