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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	fuzzerTestProtoPackage = "google.cloud.fuzztest.v1"
	fuzzerTestPBGoPackage  = "cloud.google.com/go/fuzztest/apiv1/fuzztestpb"
	fuzzerTestAPIGoPackage = "github.com/GoogleCloudPlatform/k8s-config-connector/apis"
	fuzzerTestKRMDir       = "apis/fuzztest/v1alpha1"
	fuzzerTestDirectDir    = "pkg/controller/direct/fuzztest"
)

// fzMessage builds a message of the synthetic proto package. Fields are numbered in the
// order they are added.
type fzMessage struct {
	d *descriptorpb.DescriptorProto
}

func newFzMessage(name string) *fzMessage {
	return &fzMessage{d: &descriptorpb.DescriptorProto{Name: proto.String(name)}}
}

func (m *fzMessage) add(name string, typ descriptorpb.FieldDescriptorProto_Type, typeName string, label descriptorpb.FieldDescriptorProto_Label) *fzMessage {
	f := &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(name),
		Number: proto.Int32(int32(len(m.d.Field) + 1)),
		Type:   typ.Enum(),
		Label:  label.Enum(),
	}
	if typeName != "" {
		if !strings.HasPrefix(typeName, ".") {
			typeName = "." + fuzzerTestProtoPackage + "." + typeName
		}
		f.TypeName = proto.String(typeName)
	}
	m.d.Field = append(m.d.Field, f)
	return m
}

func (m *fzMessage) str(name string) *fzMessage {
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_STRING, "", descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL)
}

func (m *fzMessage) i64(name string) *fzMessage {
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_INT64, "", descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL)
}

func (m *fzMessage) strs(name string) *fzMessage {
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_STRING, "", descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
}

// msg adds a message field. typeName is relative to the test package, or fully qualified
// with a leading dot.
func (m *fzMessage) msg(name, typeName string) *fzMessage {
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, typeName, descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL)
}

func (m *fzMessage) msgs(name, typeName string) *fzMessage {
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, typeName, descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
}

// stringMap adds a map<string, value> field. name must be a single word.
func (m *fzMessage) stringMap(name string, value descriptorpb.FieldDescriptorProto_Type) *fzMessage {
	entry := strings.ToUpper(name[:1]) + name[1:] + "Entry"
	optional := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	m.d.NestedType = append(m.d.NestedType, &descriptorpb.DescriptorProto{
		Name: proto.String(entry),
		Field: []*descriptorpb.FieldDescriptorProto{
			{Name: proto.String("key"), Number: proto.Int32(1), Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(), Label: optional.Enum()},
			{Name: proto.String("value"), Number: proto.Int32(2), Type: value.Enum(), Label: optional.Enum()},
		},
		Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
	})
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, m.d.GetName()+"."+entry, descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
}

// fuzzerTestProto builds the synthetic proto package that the tests pair with KRM types.
func fuzzerTestProto(t *testing.T) protoreflect.FileDescriptor {
	t.Helper()
	stringType := descriptorpb.FieldDescriptorProto_TYPE_STRING
	messages := []*fzMessage{
		newFzMessage("Basic").str("name").str("display_name").str("state").str("unmapped"),
		newFzMessage("Special").str("name").str("etag").stringMap("labels", stringType).stringMap("annotations", stringType).msg("meta", "Meta"),
		newFzMessage("Meta").str("name").str("etag").stringMap("labels", stringType).str("owner"),
		newFzMessage("OddLabels").stringMap("labels", descriptorpb.FieldDescriptorProto_TYPE_INT64).str("description"),
		newFzMessage("Nested").msg("config", "Config"),
		newFzMessage("Config").add("size", descriptorpb.FieldDescriptorProto_TYPE_INT32, "", descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL).str("state").str("extra"),
		newFzMessage("Repeated").msgs("rules", "Rule"),
		newFzMessage("Rule").str("action").i64("hit_count"),
		newFzMessage("OneSide").msg("spec_only", "Config").msg("status_only", "Config"),
		newFzMessage("Refs").str("network").strs("subnetworks").str("kms_key"),
		newFzMessage("WellKnown").msg("create_time", ".google.protobuf.Timestamp").msg("ttl", ".google.protobuf.Duration").
			msg("params", ".google.protobuf.Struct").msg("update_time", ".google.protobuf.Timestamp").msg("other", "Config"),
		newFzMessage("Both").str("description").msg("config", "Config"),
		newFzMessage("Tree").msg("root", "Node"),
		newFzMessage("Node").str("value").msgs("children", "Node").str("note"),
		newFzMessage("Lists").strs("tags").str("zone").msgs("rules", "Rule"),
		newFzMessage("Widget").str("name").str("etag").stringMap("labels", stringType).str("display_name").
			msg("create_time", ".google.protobuf.Timestamp").msg("config", "Config").msgs("rules", "Rule").str("network").str("unmapped"),
	}
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("google/cloud/fuzztest/v1/fuzztest.proto"),
		Package: proto.String(fuzzerTestProtoPackage),
		Syntax:  proto.String("proto3"),
		Dependency: []string{
			"google/protobuf/duration.proto",
			"google/protobuf/struct.proto",
			"google/protobuf/timestamp.proto",
		},
		Options: &descriptorpb.FileOptions{GoPackage: proto.String(fuzzerTestPBGoPackage + ";fuzztestpb")},
	}
	for _, m := range messages {
		fdp.MessageType = append(fdp.MessageType, m.d)
	}
	files, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{
		protodesc.ToFileDescriptorProto(durationpb.File_google_protobuf_duration_proto),
		protodesc.ToFileDescriptorProto(structpb.File_google_protobuf_struct_proto),
		protodesc.ToFileDescriptorProto(timestamppb.File_google_protobuf_timestamp_proto),
		fdp,
	}})
	if err != nil {
		t.Fatalf("building the test proto: %v", err)
	}
	fd, err := files.FindFileByPath(fdp.GetName())
	if err != nil {
		t.Fatalf("finding the test proto: %v", err)
	}
	return fd
}

func fuzzerTestMessage(t *testing.T, name string) protoreflect.MessageDescriptor {
	t.Helper()
	msg := fuzzerTestProto(t).Messages().ByName(protoreflect.Name(name))
	if msg == nil {
		t.Fatalf("the test proto has no message %q", name)
	}
	return msg
}

// writeFuzzerTestTree writes files, keyed by slash-separated paths, under a new temporary
// directory, and returns the directory.
func writeFuzzerTestTree(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for name, content := range files {
		p := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
			t.Fatalf("creating directory for %q: %v", p, err)
		}
		if err := os.WriteFile(p, []byte(content), 0644); err != nil {
			t.Fatalf("writing %q: %v", p, err)
		}
	}
	return root
}

// fuzzerTestKRMFile is a file of the fake KRM package; body holds the struct declarations.
func fuzzerTestKRMFile(body string) string {
	return `package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	othersvc "github.com/GoogleCloudPlatform/k8s-config-connector/apis/othersvc/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var _ refsv1beta1.ComputeNetworkRef
var _ othersvc.Config
var _ apiextensionsv1.JSON

` + body
}

// fuzzerTestMapperFile declares <base>_FromProto and <base>_ToProto for each base, with the
// signatures that generate-mapper writes. A base is a KRM type, with an optional version
// specifier, like BasicSpec or BasicSpec_v1alpha1. The file is only parsed.
func fuzzerTestMapperFile(message string, bases ...string) string {
	var b strings.Builder
	b.WriteString(`package fuzztest

import (
	pb "cloud.google.com/go/fuzztest/apiv1/fuzztestpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/fuzztest/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
`)
	for _, base := range bases {
		krmType, _, _ := strings.Cut(base, "_")
		fmt.Fprintf(&b, "\nfunc %s_FromProto(mapCtx *direct.MapContext, in *pb.%s) *krm.%s { return nil }\n", base, message, krmType)
		fmt.Fprintf(&b, "\nfunc %s_ToProto(mapCtx *direct.MapContext, in *krm.%s) *pb.%s { return nil }\n", base, krmType, message)
	}
	return b.String()
}

func planFuzzerForTest(t *testing.T, kind, message, version string, files map[string]string) (*FuzzerPlan, error) {
	t.Helper()
	root := writeFuzzerTestTree(t, files)
	direct, err := loadGoPackageIndex(filepath.Join(root, filepath.FromSlash(fuzzerTestDirectDir)))
	if err != nil {
		t.Fatalf("loading the direct package: %v", err)
	}
	return planFuzzer(fuzzerPlanInput{
		kind:    kind,
		version: version,
		message: fuzzerTestMessage(t, message),
		direct:  direct,
		krm:     newFuzzerKRMLoader(fuzzerTestAPIGoPackage, filepath.Join(root, "apis")),
	})
}

func formatFuzzerEntries(entries []FuzzerEntry) []string {
	var out []string
	for _, e := range entries {
		out = append(out, fmt.Sprintf("%s(%q)", e.Method, e.Path))
	}
	return out
}

func TestPlanFuzzerClassifiesFields(t *testing.T) {
	for _, tc := range []struct {
		name    string
		message string
		krm     string
		// mappers are the mapper pairs of the direct package; nil means Spec and ObservedState.
		mappers      []string
		want         []string
		wantSpecOnly bool
	}{
		{
			name:    "top-level spec, status and unmapped fields",
			message: "Basic",
			krm: `
type BasicSpec struct {
	DisplayName *string ` + "`json:\"displayName,omitempty\"`" + `
}

type BasicObservedState struct {
	State *string ` + "`json:\"state,omitempty\"`" + `
}
`,
			want: []string{
				`SpecField(".display_name")`,
				`StatusField(".state")`,
				`Unimplemented_Identity(".name")`,
				`Unimplemented_NotYetTriaged(".unmapped")`,
			},
		},
		{
			name:    "the resource name, etag, labels and annotations get their own methods only at the top level",
			message: "Special",
			krm: `
type SpecialSpec struct {
	Meta *Meta
}

type Meta struct {
	Owner *string
}

type SpecialObservedState struct {
}
`,
			want: []string{
				`SpecField(".meta")`,
				`Unimplemented_Identity(".name")`,
				`Unimplemented_Etag(".etag")`,
				`Unimplemented_LabelsAnnotations(".annotations")`,
				`Unimplemented_LabelsAnnotations(".labels")`,
				`Unimplemented_NotYetTriaged(".meta.etag")`,
				`Unimplemented_NotYetTriaged(".meta.labels")`,
				`Unimplemented_NotYetTriaged(".meta.name")`,
			},
		},
		{
			name:    "labels that are not a map of strings are not triaged",
			message: "OddLabels",
			krm: `
type OddLabelsSpec struct {
	Description *string
}

type OddLabelsObservedState struct {
}
`,
			want: []string{
				`SpecField(".description")`,
				`Unimplemented_NotYetTriaged(".labels")`,
			},
		},
		{
			name:    "a nested message split between spec and status",
			message: "Nested",
			krm: `
type NestedSpec struct {
	Config *ConfigSpec
}

type ConfigSpec struct {
	Size *int32
}

type NestedObservedState struct {
	Config *ConfigObservedState
}

type ConfigObservedState struct {
	State *string
}
`,
			want: []string{
				`SpecField(".config.size")`,
				`StatusField(".config.state")`,
				`Unimplemented_NotYetTriaged(".config.extra")`,
			},
		},
		{
			name:    "a repeated message split between spec and status",
			message: "Repeated",
			krm: `
type RepeatedSpec struct {
	Rules []Rule
}

type Rule struct {
	Action *string
}

type RepeatedObservedState struct {
	Rules []*RuleObservedState
}

type RuleObservedState struct {
	HitCount *int64
}
`,
			want: []string{
				`SpecField(".rules[].action")`,
				`StatusField(".rules[].hit_count")`,
			},
		},
		{
			name:    "a message that one side maps is cleared as a whole by the other side",
			message: "OneSide",
			krm: `
type OneSideSpec struct {
	SpecOnly *ConfigSpec
}

type ConfigSpec struct {
	Size *int32
}

type OneSideObservedState struct {
	StatusOnly *ConfigObservedState
}

type ConfigObservedState struct {
	State *string
}
`,
			want: []string{
				`SpecField(".spec_only")`,
				`StatusField(".status_only")`,
				`Unimplemented_NotYetTriaged(".spec_only.extra")`,
				`Unimplemented_NotYetTriaged(".spec_only.state")`,
				`Unimplemented_NotYetTriaged(".status_only.extra")`,
				`Unimplemented_NotYetTriaged(".status_only.size")`,
			},
		},
		{
			name:    "references map the field they are named after",
			message: "Refs",
			krm: `
type RefsSpec struct {
	NetworkRef     *refsv1beta1.ComputeNetworkRef
	SubnetworkRefs []refsv1beta1.ComputeSubnetworkRef
}

type RefsObservedState struct {
}
`,
			want: []string{
				`SpecField(".network")`,
				`SpecField(".subnetworks")`,
				`Unimplemented_NotYetTriaged(".kms_key")`,
			},
		},
		{
			name:    "well-known types and types of other packages are leaves",
			message: "WellKnown",
			krm: `
type WellKnownSpec struct {
	TTL    *string
	Params *apiextensionsv1.JSON
	Other  *othersvc.Config
}

type WellKnownObservedState struct {
	CreateTime *string
}
`,
			want: []string{
				`SpecField(".other")`,
				`SpecField(".params")`,
				`SpecField(".ttl")`,
				`StatusField(".create_time")`,
				`Unimplemented_NotYetTriaged(".update_time")`,
			},
		},
		{
			name:    "a field that both sides map is a spec field",
			message: "Both",
			krm: `
type BothSpec struct {
	Description *string
	Config      *ConfigSpec
}

type ConfigSpec struct {
	Size *int32
}

type BothObservedState struct {
	Description *string
	Config      *ConfigObservedState
}

type ConfigObservedState struct {
	Size  *int32
	State *string
}
`,
			want: []string{
				`SpecField(".config.size")`,
				`SpecField(".description")`,
				`StatusField(".config.state")`,
				`Unimplemented_NotYetTriaged(".config.extra")`,
			},
		},
		{
			name:    "a KRM type that contains itself is not expanded again",
			message: "Tree",
			krm: `
type TreeSpec struct {
	Root *Node
}

type Node struct {
	Value    *string
	Children []Node
}

type TreeObservedState struct {
}
`,
			want: []string{
				`SpecField(".root")`,
				`Unimplemented_NotYetTriaged(".root.children")`,
				`Unimplemented_NotYetTriaged(".root.note")`,
			},
		},
		{
			name:    "a list mapped to a single KRM value is unmapped",
			message: "Lists",
			krm: `
type ListsSpec struct {
	Tags  *string
	Zone  []string
	Rules *Rule
}

type Rule struct {
	Action *string
}

type ListsObservedState struct {
}
`,
			want: []string{
				`SpecField(".zone")`,
				`Unimplemented_NotYetTriaged(".rules")`,
				`Unimplemented_NotYetTriaged(".tags")`,
			},
		},
		{
			name:    "without ObservedState mappers the fuzzer is spec-only, and the ObservedState struct still names the status fields",
			message: "Basic",
			krm: `
type BasicSpec struct {
	DisplayName *string
}

type BasicObservedState struct {
	State *string
}
`,
			mappers: []string{"BasicSpec"},
			want: []string{
				`SpecField(".display_name")`,
				`StatusField(".state")`,
				`Unimplemented_Identity(".name")`,
				`Unimplemented_NotYetTriaged(".unmapped")`,
			},
			wantSpecOnly: true,
		},
		{
			name:    "without an ObservedState struct every field the spec doesn't map is not triaged",
			message: "Basic",
			krm: `
type BasicSpec struct {
	DisplayName *string
}
`,
			mappers: []string{"BasicSpec"},
			want: []string{
				`SpecField(".display_name")`,
				`Unimplemented_Identity(".name")`,
				`Unimplemented_NotYetTriaged(".state")`,
				`Unimplemented_NotYetTriaged(".unmapped")`,
			},
			wantSpecOnly: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mappers := tc.mappers
			if mappers == nil {
				mappers = []string{tc.message + "Spec", tc.message + "ObservedState"}
			}
			files := map[string]string{
				fuzzerTestKRMDir + "/types.go":               fuzzerTestKRMFile(tc.krm),
				fuzzerTestDirectDir + "/mapper.generated.go": fuzzerTestMapperFile(tc.message, mappers...),
			}

			// Act
			plan, err := planFuzzerForTest(t, tc.message, tc.message, "v1alpha1", files)

			// Assert
			if err != nil {
				t.Fatalf("planFuzzer: %v", err)
			}
			if diff := cmp.Diff(tc.want, formatFuzzerEntries(plan.Entries)); diff != "" {
				t.Errorf("entries differ (-want +got):\n%s", diff)
			}
			if gotSpecOnly := plan.Status == nil; gotSpecOnly != tc.wantSpecOnly {
				t.Errorf("spec-only = %v, want %v", gotSpecOnly, tc.wantSpecOnly)
			}
		})
	}
}

func TestPlanFuzzerFindsMappers(t *testing.T) {
	const krm = `
type BasicSpec struct {
	DisplayName *string
}

type BasicObservedState struct {
	State *string
}

type BasicStatus struct {
	ObservedState *CustomObservedStateStatus
	State         *string
}

type CustomObservedStateStatus struct {
	State *string
}
`
	for _, tc := range []struct {
		name          string
		directFiles   map[string]string
		wantSpec      string
		wantStatus    string
		wantErrSubstr string
	}{
		{
			name: "mappers with a version specifier, as generate-mapper --multiversion names them",
			directFiles: map[string]string{
				"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicSpec_v1alpha1", "BasicObservedState_v1alpha1"),
			},
			wantSpec:   "BasicSpec_v1alpha1_FromProto",
			wantStatus: "BasicObservedState_v1alpha1_FromProto",
		},
		{
			name: "hand-written and generated mappers both count",
			directFiles: map[string]string{
				"basic_mappers.go":    fuzzerTestMapperFile("Basic", "BasicSpec"),
				"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicObservedState"),
			},
			wantSpec:   "BasicSpec_FromProto",
			wantStatus: "BasicObservedState_FromProto",
		},
		{
			name: "ObservedState mappers named after the ObservedState field of the Status struct",
			directFiles: map[string]string{
				"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicSpec", "CustomObservedStateStatus"),
			},
			wantSpec:   "BasicSpec_FromProto",
			wantStatus: "CustomObservedStateStatus_FromProto",
		},
		{
			name: "Status mappers when there are no ObservedState mappers",
			directFiles: map[string]string{
				"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicSpec", "BasicStatus"),
			},
			wantSpec:   "BasicSpec_FromProto",
			wantStatus: "BasicStatus_FromProto",
		},
		{
			name: "no Spec mappers",
			directFiles: map[string]string{
				"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicObservedState"),
			},
			wantErrSubstr: "no Spec mappers for kind Basic",
		},
		{
			name: "Spec mappers of another proto message",
			directFiles: map[string]string{
				"mapper.generated.go": fuzzerTestMapperFile("Nested", "BasicSpec"),
			},
			wantErrSubstr: "BasicSpec_FromProto converts *Nested",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			files := map[string]string{fuzzerTestKRMDir + "/types.go": fuzzerTestKRMFile(krm)}
			for name, content := range tc.directFiles {
				files[fuzzerTestDirectDir+"/"+name] = content
			}

			// Act
			plan, err := planFuzzerForTest(t, "Basic", "Basic", "v1alpha1", files)

			// Assert
			if tc.wantErrSubstr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErrSubstr) {
					t.Fatalf("planFuzzer error = %v, want an error containing %q", err, tc.wantErrSubstr)
				}
				return
			}
			if err != nil {
				t.Fatalf("planFuzzer: %v", err)
			}
			if plan.Spec.FromProto != tc.wantSpec {
				t.Errorf("Spec mapper = %q, want %q", plan.Spec.FromProto, tc.wantSpec)
			}
			if plan.Status == nil || plan.Status.FromProto != tc.wantStatus {
				t.Errorf("Status mapper = %+v, want %q", plan.Status, tc.wantStatus)
			}
			if plan.Spec.ProtoGoPackage != fuzzerTestPBGoPackage {
				t.Errorf("proto Go package = %q, want %q", plan.Spec.ProtoGoPackage, fuzzerTestPBGoPackage)
			}
		})
	}
}

// fuzzerTestWidgetKRM is the KRM package of the FuzzTestWidget kind.
const fuzzerTestWidgetKRM = `
type FuzzTestWidgetSpec struct {
	DisplayName *string
	Config      *ConfigSpec
	Rules       []Rule
	NetworkRef  *refsv1beta1.ComputeNetworkRef
}

type ConfigSpec struct {
	Size *int32
}

type Rule struct {
	Action *string
}

type FuzzTestWidgetObservedState struct {
	CreateTime *string
	Config     *ConfigObservedState
	Rules      []RuleObservedState
}

type ConfigObservedState struct {
	State *string
}

type RuleObservedState struct {
	HitCount *int64
}
`

func newFuzzerTestGenerator(root string) *FuzzerGenerator {
	return NewFuzzerGenerator(FuzzerGeneratorOptions{
		Group:            "fuzztest.cnrm.cloud.google.com",
		Version:          "v1alpha1",
		ProtoService:     fuzzerTestProtoPackage,
		APIGoPackagePath: fuzzerTestAPIGoPackage + "/",
		APIDirectory:     filepath.Join(root, "apis"),
		DirectDirectory:  filepath.Join(root, filepath.FromSlash(fuzzerTestDirectDir)),
		Year:             2026,
	})
}

func widgetFuzzerTestFiles(mappers ...string) map[string]string {
	if len(mappers) == 0 {
		mappers = []string{"FuzzTestWidgetSpec", "FuzzTestWidgetObservedState"}
	}
	return map[string]string{
		fuzzerTestKRMDir + "/fuzztestwidget_types.go": fuzzerTestKRMFile(fuzzerTestWidgetKRM),
		fuzzerTestDirectDir + "/mapper.generated.go":  fuzzerTestMapperFile("Widget", mappers...),
	}
}

const wantWidgetFuzzer = `// Copyright 2026 Google LLC
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

// Code generated by dev/tasks/generate-all. DO NOT EDIT.
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// +generated:fuzzer
// krm.group: fuzztest.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.message: google.cloud.fuzztest.v1.Widget
// proto.service: google.cloud.fuzztest.v1

package fuzztest

import (
	pb "cloud.google.com/go/fuzztest/apiv1/fuzztestpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzTestWidgetFuzzer())
}

func fuzzTestWidgetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Widget{},
		FuzzTestWidgetSpec_FromProto, FuzzTestWidgetSpec_ToProto,
		FuzzTestWidgetObservedState_FromProto, FuzzTestWidgetObservedState_ToProto,
	)

	// Spec fields
	f.SpecField(".config.size")
	f.SpecField(".display_name")
	f.SpecField(".network")
	f.SpecField(".rules[].action")

	// Status fields
	f.StatusField(".config.state")
	f.StatusField(".create_time")
	f.StatusField(".rules[].hit_count")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Etag()
	f.Unimplemented_LabelsAnnotations(".labels")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".config.extra")
	f.Unimplemented_NotYetTriaged(".unmapped")

	return f
}
`

func TestFuzzerGeneratorWritesFuzzer(t *testing.T) {
	// Arrange
	root := writeFuzzerTestTree(t, widgetFuzzerTestFiles())
	g := newFuzzerTestGenerator(root)

	// Act
	result, err := g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))

	// Assert
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if result.Skipped {
		t.Fatalf("Generate skipped the kind: %s", result.SkipReason)
	}
	wantPath := filepath.Join(root, filepath.FromSlash(fuzzerTestDirectDir), "fuzztestwidget_fuzzer.generated.go")
	if result.Path != wantPath {
		t.Errorf("Path = %q, want %q", result.Path, wantPath)
	}
	got, err := os.ReadFile(wantPath)
	if err != nil {
		t.Fatalf("reading the generated fuzzer: %v", err)
	}
	if diff := cmp.Diff(wantWidgetFuzzer, string(got)); diff != "" {
		t.Errorf("generated fuzzer differs (-want +got):\n%s", diff)
	}
}

func TestFuzzerGeneratorWritesSpecOnlyFuzzer(t *testing.T) {
	// Arrange
	root := writeFuzzerTestTree(t, widgetFuzzerTestFiles("FuzzTestWidgetSpec"))
	g := newFuzzerTestGenerator(root)

	// Act
	result, err := g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))

	// Assert
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	got, err := os.ReadFile(result.Path)
	if err != nil {
		t.Fatalf("reading the generated fuzzer: %v", err)
	}
	for _, want := range []string{
		"fuzztesting.RegisterKRMSpecFuzzer(fuzzTestWidgetFuzzer())",
		"f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Widget{},\n\t\tFuzzTestWidgetSpec_FromProto, FuzzTestWidgetSpec_ToProto,\n\t)",
		`f.StatusField(".create_time")`,
	} {
		if !strings.Contains(string(got), want) {
			t.Errorf("the spec-only fuzzer does not contain %q:\n%s", want, got)
		}
	}
	if strings.Contains(string(got), "ObservedState_FromProto") {
		t.Errorf("the spec-only fuzzer refers to ObservedState mappers:\n%s", got)
	}
}

func TestFuzzerGeneratorWiresFilterHooks(t *testing.T) {
	// Arrange
	files := widgetFuzzerTestFiles()
	files[fuzzerTestDirectDir+"/fuzztestwidget_fuzzer_filters.go"] = `package fuzztest

import pb "cloud.google.com/go/fuzztest/apiv1/fuzztestpb"

func fuzzTestWidgetFuzzerFilterSpec(in *pb.Widget) {}

func fuzzTestWidgetFuzzerFilterStatus(in *pb.Widget) {}
`
	root := writeFuzzerTestTree(t, files)
	g := newFuzzerTestGenerator(root)

	// Act
	result, err := g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))

	// Assert
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	got, err := os.ReadFile(result.Path)
	if err != nil {
		t.Fatalf("reading the generated fuzzer: %v", err)
	}
	want := "\n\tf.FilterSpec = fuzzTestWidgetFuzzerFilterSpec\n\tf.FilterStatus = fuzzTestWidgetFuzzerFilterStatus\n\n\treturn f\n}\n"
	if !strings.HasSuffix(string(got), want) {
		t.Errorf("the fuzzer does not end with the filter hooks %q:\n%s", want, got)
	}
}

func TestFuzzerGeneratorSkipsHandwrittenFuzzers(t *testing.T) {
	const llmFuzzer = `// +tool:fuzz-gen
// proto.message: google.cloud.fuzztest.v1.Widget

package %s

import (
	pb "cloud.google.com/go/fuzztest/apiv1/fuzztestpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	%s
)

func init() {
	fuzztesting.RegisterKRMFuzzer(widgetFuzzer())
}

func widgetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Widget{},
		%sFuzzTestWidgetSpec_FromProto, %sFuzzTestWidgetSpec_ToProto,
		%sFuzzTestWidgetObservedState_FromProto, %sFuzzTestWidgetObservedState_ToProto,
	)
	return f
}
`
	for _, tc := range []struct {
		name       string
		file       string
		content    string
		wantReason string
	}{
		{
			name: "a file that is not generated declares the fuzzer function",
			file: fuzzerTestDirectDir + "/fuzztestwidget_fuzzer.go",
			content: `package fuzztest

import "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"

func fuzzTestWidgetFuzzer() fuzztesting.KRMFuzzer { return nil }
`,
			wantReason: "fuzztestwidget_fuzzer.go declares fuzzTestWidgetFuzzer",
		},
		{
			name:       "a fuzzer written by an LLM registers a fuzzer built from the Spec mappers",
			file:       fuzzerTestDirectDir + "/widget_fuzzer.go",
			content:    fmt.Sprintf(llmFuzzer, "fuzztest", "", "", "", "", ""),
			wantReason: "widget_fuzzer.go registers a fuzzer built from FuzzTestWidgetSpec_FromProto",
		},
		{
			name: "a fuzzer in a subpackage uses the mappers of the package",
			file: fuzzerTestDirectDir + "/widget/widget_fuzzer.go",
			content: fmt.Sprintf(llmFuzzer, "widget", `"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/fuzztest"`,
				"fuzztest.", "fuzztest.", "fuzztest.", "fuzztest."),
			wantReason: "widget_fuzzer.go registers a fuzzer built from FuzzTestWidgetSpec_FromProto",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			files := widgetFuzzerTestFiles()
			files[tc.file] = tc.content
			stale := fuzzerTestDirectDir + "/fuzztestwidget_fuzzer.generated.go"
			files[stale] = "package fuzztest\n"
			root := writeFuzzerTestTree(t, files)
			g := newFuzzerTestGenerator(root)

			// Act
			result, err := g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))

			// Assert
			if err != nil {
				t.Fatalf("Generate: %v", err)
			}
			if !result.Skipped {
				t.Fatalf("Generate did not skip the kind")
			}
			if !strings.HasSuffix(result.SkipReason, tc.wantReason) {
				t.Errorf("SkipReason = %q, want it to end with %q", result.SkipReason, tc.wantReason)
			}
			if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(stale))); !os.IsNotExist(err) {
				t.Errorf("the stale generated fuzzer was not removed: %v", err)
			}
		})
	}
}

func TestFuzzerGeneratorIsStable(t *testing.T) {
	// Arrange
	root := writeFuzzerTestTree(t, widgetFuzzerTestFiles())
	g := newFuzzerTestGenerator(root)
	first, err := g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	b, err := os.ReadFile(first.Path)
	if err != nil {
		t.Fatalf("reading the generated fuzzer: %v", err)
	}
	// A file generated in an earlier year keeps its copyright year.
	earlier := strings.Replace(string(b), "Copyright 2026", "Copyright 2025", 1)
	if err := os.WriteFile(first.Path, []byte(earlier), 0644); err != nil {
		t.Fatalf("writing the generated fuzzer: %v", err)
	}

	// Act
	if _, err := g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget")); err != nil {
		t.Fatalf("Generate: %v", err)
	}

	// Assert
	got, err := os.ReadFile(first.Path)
	if err != nil {
		t.Fatalf("reading the generated fuzzer: %v", err)
	}
	if diff := cmp.Diff(earlier, string(got)); diff != "" {
		t.Errorf("regenerating changed the fuzzer (-want +got):\n%s", diff)
	}
}

func TestFuzzerGeneratorRejectsMappersOfAnotherVersion(t *testing.T) {
	// Arrange
	files := widgetFuzzerTestFiles()
	files["apis/fuzztest/v1beta1/fuzztestwidget_types.go"] = strings.Replace(fuzzerTestKRMFile(fuzzerTestWidgetKRM), "package v1alpha1", "package v1beta1", 1)
	files[fuzzerTestDirectDir+"/mapper.generated.go"] = strings.ReplaceAll(fuzzerTestMapperFile("Widget", "FuzzTestWidgetSpec", "FuzzTestWidgetObservedState"),
		"apis/fuzztest/v1alpha1", "apis/fuzztest/v1beta1")
	root := writeFuzzerTestTree(t, files)
	g := newFuzzerTestGenerator(root)

	// Act
	_, err := g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))

	// Assert
	want := "FuzzTestWidgetSpec_FromProto returns a type of " + fuzzerTestAPIGoPackage + "/fuzztest/v1beta1"
	if err == nil || !strings.Contains(err.Error(), want) {
		t.Errorf("Generate error = %v, want an error containing %q", err, want)
	}
}

func TestFuzzerFuncName(t *testing.T) {
	for _, tc := range []struct {
		kind string
		want string
	}{
		{kind: "ComputeTargetVPNGateway", want: "computeTargetVPNGatewayFuzzer"},
		{kind: "BigQueryDataPolicy", want: "bigQueryDataPolicyFuzzer"},
		{kind: "IAMDenyPolicy", want: "iamDenyPolicyFuzzer"},
		{kind: "SQLInstance", want: "sqlInstanceFuzzer"},
		{kind: "KMS", want: "kmsFuzzer"},
	} {
		t.Run(tc.kind, func(t *testing.T) {
			if got := fuzzerFuncName(tc.kind); got != tc.want {
				t.Errorf("fuzzerFuncName(%q) = %q, want %q", tc.kind, got, tc.want)
			}
		})
	}
}
