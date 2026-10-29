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
			Reason: "empty-observedstate",
			Detail: "nothing was generated into status.observedState; the proto probably marks no field OUTPUT_ONLY, so output fields landed in the Spec. Decide what belongs in status",
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
		PrepopulateResource: true,
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

func TestRunGenerateCRDSkipScaffoldFilesSuppressesJudgements(t *testing.T) {
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "apis")
	queuePath := filepath.Join(apiDir, "example", judgement.FileName)

	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	msgType := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{{
		Name:    proto.String("google/cloud/example/v1/widget.proto"),
		Package: proto.String("google.cloud.example.v1"),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("Widget"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: proto.String("name"), Number: proto.Int32(1), Type: str},
					{Name: proto.String("config"), Number: proto.Int32(2), Type: msgType, TypeName: proto.String(".google.cloud.example.v1.WidgetConfig")},
				},
			},
			{
				Name: proto.String("WidgetConfig"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: proto.String("widget"), Number: proto.Int32(1), Type: str},
				},
			},
		},
	}}}
	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	pbPath := filepath.Join(dir, "widget.pb")
	writeFile(t, pbPath, string(b))

	o := &GenerateCRDOptions{
		GenerateOptions: &options.GenerateOptions{
			ProtoSourcePath: pbPath,
			APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
		},
		ServiceName:         "google.cloud.example.v1",
		OutputAPIDirectory:  apiDir,
		Resources:           options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
		SkipScaffoldFiles:   true,
		PrepopulateResource: true,
	}

	if err := RunGenerateCRD(context.Background(), o); err != nil {
		t.Fatalf("RunGenerateCRD() error: %v", err)
	}
	if _, statErr := os.Stat(queuePath); !os.IsNotExist(statErr) {
		t.Errorf("expected no judgement_queue.yaml when SkipScaffoldFiles is true, got stat err: %v", statErr)
	}
	genTypes := readFile(t, filepath.Join(apiDir, "example", "v1alpha1", "types.generated.go"))
	if strings.Contains(genTypes, "+kcc:guess=") {
		t.Errorf("expected no +kcc:guess comments in types.generated.go when SkipScaffoldFiles is true, got:\n%s", genTypes)
	}
}

func TestIsNonGAProto(t *testing.T) {
	for _, tc := range []struct {
		fullName string
		want     bool
	}{
		{"google.cloud.apihub.v1.PluginInstance", false},
		{"google.cloud.retail.v2.Control", false},
		{"google.cloud.talent.v4.Job", false},
		{"google.cloud.ces.v1beta.Deployment", true},
		{"google.cloud.dataform.v1beta1.Config", true},
		{"google.cloud.gdchardwaremanagement.v1alpha.Hardware", true},
		{"google.cloud.bigquery.migration.v2alpha.MigrationWorkflow", true},
		{"google.cloud.speech.v1p1beta1.CustomClass", true},
	} {
		if got := isNonGAProto(tc.fullName); got != tc.want {
			t.Errorf("isNonGAProto(%q) = %v, want %v", tc.fullName, got, tc.want)
		}
	}
}

func TestRunGenerateCRDRejectsNewNonGAProtoWithoutFlag(t *testing.T) {
	apiDir := t.TempDir()
	pbPath := writeBetaWidgetProto(t, t.TempDir())

	baseOpts := func() *GenerateCRDOptions {
		o := &GenerateCRDOptions{
			GenerateOptions: &options.GenerateOptions{
				ProtoSourcePath: pbPath,
				APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
			},
			ServiceName:         "google.cloud.example.v1beta",
			OutputAPIDirectory:  apiDir,
			Resources:           options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
			PrepopulateResource: true,
		}
		return o
	}

	// 1. Scaffolding a new resource from a non-GA proto fails without --allow-non-ga-proto.
	err := RunGenerateCRD(context.Background(), baseOpts())
	if err == nil || !strings.Contains(err.Error(), "--allow-non-ga-proto") {
		t.Fatalf("expected --allow-non-ga-proto error, got: %v", err)
	}
	typeFile := filepath.Join(apiDir, "example/v1alpha1/examplewidget_types.go")
	if _, statErr := os.Stat(typeFile); !os.IsNotExist(statErr) {
		t.Fatalf("expected %s not to be created on failure", typeFile)
	}

	// 2. Passing --allow-non-ga-proto succeeds and creates <kind>_types.go.
	allowed := baseOpts()
	allowed.AllowNonGAProto = true
	if err := RunGenerateCRD(context.Background(), allowed); err != nil {
		t.Fatalf("RunGenerateCRD with AllowNonGAProto: %v", err)
	}
	if _, statErr := os.Stat(typeFile); statErr != nil {
		t.Fatalf("expected %s to exist after AllowNonGAProto: %v", typeFile, statErr)
	}

	// 3. Re-running when <kind>_types.go already exists succeeds even without --allow-non-ga-proto.
	if err := RunGenerateCRD(context.Background(), baseOpts()); err != nil {
		t.Fatalf("RunGenerateCRD on existing resource without AllowNonGAProto: %v", err)
	}
}

func writeBetaWidgetProto(t *testing.T, dir string) string {
	t.Helper()
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{{
		Name:    proto.String("google/cloud/example/v1beta/widget.proto"),
		Package: proto.String("google.cloud.example.v1beta"),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: proto.String("Widget"),
			Field: []*descriptorpb.FieldDescriptorProto{
				{Name: proto.String("name"), Number: proto.Int32(1), Type: str},
				{Name: proto.String("display_name"), Number: proto.Int32(2), Type: str},
			},
		}},
	}}}
	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	path := filepath.Join(dir, "widget_beta.pb")
	writeFile(t, path, string(b))
	return path
}

func TestRunGenerateCRDQueuesNestedJudgementsPerKind(t *testing.T) {
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "apis")
	queuePath := filepath.Join(apiDir, "example", judgement.FileName)

	// ExampleGadget is already reviewed; ExampleWidget is newly scaffolded.
	writeFile(t, filepath.Join(apiDir, "example", "v1alpha1", "examplegadget_types.go"), "package v1alpha1\n\ntype ExampleGadgetSpec struct{}\n")
	writeFile(t, queuePath, `entries:
- kind: ExampleGadget
  group: example.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: resolved
  resolution: accepted
`)

	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	i32 := descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum()
	msgType := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
	rep := descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{{
		Name:    proto.String("google/cloud/example/v1/widget.proto"),
		Package: proto.String("google.cloud.example.v1"),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("SubTask"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: proto.String("label"), Number: proto.Int32(1), Type: str},
				},
			},
			{
				Name: proto.String("SharedConfig"),
				NestedType: []*descriptorpb.DescriptorProto{
					{
						Name: proto.String("ByIndexEntry"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: proto.String("key"), Number: proto.Int32(1), Type: i32},
							{Name: proto.String("value"), Number: proto.Int32(2), Type: str},
						},
						Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
					},
					{
						Name: proto.String("TasksEntry"),
						Field: []*descriptorpb.FieldDescriptorProto{
							{Name: proto.String("key"), Number: proto.Int32(1), Type: str},
							{Name: proto.String("value"), Number: proto.Int32(2), Type: msgType, TypeName: proto.String(".google.cloud.example.v1.SubTask")},
						},
						Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
					},
				},
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: proto.String("gadget"), Number: proto.Int32(1), Type: str},
					{Name: proto.String("by_index"), Number: proto.Int32(2), Type: msgType, TypeName: proto.String(".google.cloud.example.v1.SharedConfig.ByIndexEntry"), Label: rep},
					{Name: proto.String("tasks"), Number: proto.Int32(3), Type: msgType, TypeName: proto.String(".google.cloud.example.v1.SharedConfig.TasksEntry"), Label: rep},
				},
			},
			{
				Name: proto.String("Widget"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: proto.String("name"), Number: proto.Int32(1), Type: str},
					{Name: proto.String("config"), Number: proto.Int32(2), Type: msgType, TypeName: proto.String(".google.cloud.example.v1.SharedConfig")},
				},
			},
			{
				Name: proto.String("Gadget"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: proto.String("name"), Number: proto.Int32(1), Type: str},
					{Name: proto.String("config"), Number: proto.Int32(2), Type: msgType, TypeName: proto.String(".google.cloud.example.v1.SharedConfig")},
				},
			},
		},
	}}}
	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	pbPath := filepath.Join(dir, "shared.pb")
	writeFile(t, pbPath, string(b))

	o := &GenerateCRDOptions{
		GenerateOptions: &options.GenerateOptions{
			ProtoSourcePath: pbPath,
			APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
		},
		ServiceName:        "google.cloud.example.v1",
		OutputAPIDirectory: apiDir,
		Resources: options.ResourceList{
			{Kind: "ExampleWidget", ProtoName: "Widget"},
			{Kind: "ExampleGadget", ProtoName: "Gadget"},
		},
		PrepopulateResource: true,
	}

	if err := RunGenerateCRD(context.Background(), o); err != nil {
		t.Fatalf("RunGenerateCRD() error: %v", err)
	}

	got := readQueueEntries(t, queuePath)
	for _, e := range got {
		if e.ProtoMessage != "" {
			t.Errorf("unexpected ProtoMessage entry in queue: %+v", e)
		}
	}
	want := []judgement.Entry{
		{
			Kind:       "ExampleGadget",
			Group:      "example.cnrm.cloud.google.com",
			Reason:     judgement.ReasonUntriagedBulkGeneration,
			Status:     judgement.StatusResolved,
			Resolution: judgement.ResolutionAccepted,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Reason: judgement.ReasonUntriagedBulkGeneration,
			Detail: "spec was generated from proto definition; verify refs, omissions, and KRM conventions",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Reason: judgement.ReasonEmptyObservedState,
			Detail: "nothing was generated into status.observedState; the proto probably marks no field OUTPUT_ONLY, so output fields landed in the Spec. Decide what belongs in status",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Reason: judgement.ReasonVerifyResourceDocsLink,
			Detail: "found no docs root for this API; find the REST reference page by hand",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Reason: judgement.ReasonVerifyServiceDocsLink,
			Detail: "neither the Discovery directory nor the service config names a docs page; add one",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Field:  ".spec.config.byIndex",
			Reason: judgement.ReasonUnsupportedFieldType,
			Detail: "unsupported map type with key int32 and value string",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Field:  ".spec.config.gadget",
			Reason: judgement.ReasonPossibleReferenceBySibling,
			Detail: "the name matches ExampleGadget, a resource this service declares; confirm whether it should be a reference",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Field:  ".spec.config.tasks",
			Reason: judgement.ReasonMessageMapDerived,
			Detail: "derived as map[string]SubTask from proto google.cloud.example.v1.SubTask; verify Go type and mapping",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Field:  ".spec.location",
			Reason: "location-parent-unknown",
			Detail: "the proto declares no google.api.resource, so the parent shape is unknown; drop location if the resource is not regional",
			Status: judgement.StatusOpen,
		},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("queue entries mismatch (-want +got):\n%s", diff)
	}
}

func TestRunGenerateCRDScaffoldsBeforeWriteVisitedMessages(t *testing.T) {
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "apis")

	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	msgType := descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum()
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{
		{
			Name:    proto.String("google/cloud/example/v1/widget.proto"),
			Package: proto.String("google.cloud.example.v1"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{
				{
					Name: proto.String("WidgetSettings"),
					Field: []*descriptorpb.FieldDescriptorProto{
						{Name: proto.String("mode"), Number: proto.Int32(1), Type: str},
					},
				},
				{
					Name: proto.String("Widget"),
					Field: []*descriptorpb.FieldDescriptorProto{
						{Name: proto.String("name"), Number: proto.Int32(1), Type: str},
						{Name: proto.String("settings"), Number: proto.Int32(2), Type: msgType, TypeName: proto.String(".google.cloud.example.v1.WidgetSettings")},
					},
				},
			},
		},
	}}
	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	pbPath := filepath.Join(dir, "widget.pb")
	writeFile(t, pbPath, string(b))

	runV1 := func() {
		o1 := &GenerateCRDOptions{
			GenerateOptions: &options.GenerateOptions{
				ProtoSourcePath: pbPath,
				APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
			},
			ServiceName:         "google.cloud.example.v1",
			OutputAPIDirectory:  apiDir,
			Resources:           options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
			PrepopulateResource: true,
		}
		if err := RunGenerateCRD(context.Background(), o1); err != nil {
			t.Fatalf("RunGenerateCRD(v1) error: %v", err)
		}
	}

	runV1()

	genPath := filepath.Join(apiDir, "example", "v1alpha1", "types.generated.go")
	afterFirstPass := readFile(t, genPath)
	for _, wantSubstr := range []string{
		"// proto.service: google.cloud.example.v1",
		"// resource: ExampleWidget:Widget",
		"type WidgetSettings struct",
	} {
		if !strings.Contains(afterFirstPass, wantSubstr) {
			t.Errorf("expected types.generated.go to contain %q, got:\n%s", wantSubstr, afterFirstPass)
		}
	}
	if strings.Contains(afterFirstPass, "type Widget struct") {
		t.Errorf("expected types.generated.go not to contain top-level resource struct \"type Widget struct\" during fresh scaffolding, got:\n%s", afterFirstPass)
	}

	// Re-running v1 must be idempotent from the very first pass.
	runV1()
	afterSecondPass := readFile(t, genPath)
	if diff := cmp.Diff(afterFirstPass, afterSecondPass); diff != "" {
		t.Errorf("re-running RunGenerateCRD was not idempotent (-first +second):\n%s", diff)
	}
}

// TestRunGenerateCRDRecoversMissingQueueForPrepopulatedKind verifies that when
// a prepopulated <kind>_types.go file (carrying +kcc:required-from-proto)
// already exists on disk, but its entries in judgement_queue.yaml were lost
// (for example, clobbered during an add/add rebase conflict when another
// resource was generated in the same service), re-running RunGenerateCRD
// recovers both the spec/reference entries and the parent/location entries
// without rewriting <kind>_types.go.
func TestRunGenerateCRDRecoversMissingQueueForPrepopulatedKind(t *testing.T) {
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "apis")
	typesPath := filepath.Join(apiDir, "example", "v1alpha1", "examplewidget_types.go")
	queuePath := filepath.Join(apiDir, "example", judgement.FileName)

	prepopulatedTypesFile := `// API sources for ExampleWidget, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/master/google/cloud/example/v1/widget.proto
// +kcc:guess=source-link reason=verify-service-docs-link
// +kcc:source:service-docs=

package v1alpha1

// ExampleWidgetSpec was scaffolded with --prepopulate-resource and later edited.
// +kcc:spec:proto=google.cloud.example.v1.Widget
// +kcc:required-from-proto
type ExampleWidgetSpec struct{}
`
	writeFile(t, typesPath, prepopulatedTypesFile)
	// Simulate a rebase conflict where only another Kind's entry survived in
	// judgement_queue.yaml and ExampleWidget's entries were lost.
	writeFile(t, queuePath, `entries:
- kind: ExampleOther
  group: example.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: open
`)

	o := &GenerateCRDOptions{
		GenerateOptions: &options.GenerateOptions{
			ProtoSourcePath: writeWidgetProto(t, dir),
			APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
		},
		ServiceName:         "google.cloud.example.v1",
		OutputAPIDirectory:  apiDir,
		Resources:           options.ResourceList{{Kind: "ExampleWidget", ProtoName: "Widget"}},
		PrepopulateResource: true,
	}

	if err := RunGenerateCRD(context.Background(), o); err != nil {
		t.Fatalf("RunGenerateCRD() error: %v", err)
	}
	if diff := cmp.Diff(prepopulatedTypesFile, readFile(t, typesPath)); diff != "" {
		t.Errorf("RunGenerateCRD() changed the types file (-want +got):\n%s", diff)
	}

	want := []judgement.Entry{
		{
			Kind:   "ExampleOther",
			Group:  "example.cnrm.cloud.google.com",
			Reason: "untriaged-bulk-generation",
			Status: judgement.StatusOpen,
		},
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
			Reason: "empty-observedstate",
			Detail: "nothing was generated into status.observedState; the proto probably marks no field OUTPUT_ONLY, so output fields landed in the Spec. Decide what belongs in status",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Reason: "verify-service-docs-link",
			Detail: "neither the Discovery directory nor the service config names a docs page; add one",
			Status: judgement.StatusOpen,
		},
		{
			Kind:   "ExampleWidget",
			Group:  "example.cnrm.cloud.google.com",
			Field:  ".spec.location",
			Reason: "location-parent-unknown",
			Detail: "the proto declares no google.api.resource, so the parent shape is unknown; drop location if the resource is not regional",
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
	}
	if diff := cmp.Diff(want, readQueueEntries(t, queuePath)); diff != "" {
		t.Errorf("queue entries mismatch (-want +got):\n%s", diff)
	}
}
