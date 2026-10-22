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

type fzMessage struct{ d *descriptorpb.DescriptorProto }

func newFzMessage(name string) *fzMessage {
	return &fzMessage{d: &descriptorpb.DescriptorProto{Name: proto.String(name)}}
}

func (m *fzMessage) add(name string, typ descriptorpb.FieldDescriptorProto_Type, typeName string, label descriptorpb.FieldDescriptorProto_Label) *fzMessage {
	if typeName != "" && !strings.HasPrefix(typeName, ".") {
		typeName = "." + fuzzerTestProtoPackage + "." + typeName
	}
	f := &descriptorpb.FieldDescriptorProto{
		Name: proto.String(name), Number: proto.Int32(int32(len(m.d.Field) + 1)), Type: typ.Enum(), Label: label.Enum(),
	}
	if typeName != "" {
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
func (m *fzMessage) msg(name, typeName string) *fzMessage {
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, typeName, descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL)
}
func (m *fzMessage) msgs(name, typeName string) *fzMessage {
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, typeName, descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
}

func (m *fzMessage) stringMap(name string, value descriptorpb.FieldDescriptorProto_Type) *fzMessage {
	entry := strings.ToUpper(name[:1]) + name[1:] + "Entry"
	opt := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	m.d.NestedType = append(m.d.NestedType, &descriptorpb.DescriptorProto{
		Name: proto.String(entry),
		Field: []*descriptorpb.FieldDescriptorProto{
			{Name: proto.String("key"), Number: proto.Int32(1), Type: descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(), Label: opt.Enum()},
			{Name: proto.String("value"), Number: proto.Int32(2), Type: value.Enum(), Label: opt.Enum()},
		},
		Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
	})
	return m.add(name, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, m.d.GetName()+"."+entry, descriptorpb.FieldDescriptorProto_LABEL_REPEATED)
}

func fuzzerTestMessage(t *testing.T, name string) protoreflect.MessageDescriptor {
	t.Helper()
	strT := descriptorpb.FieldDescriptorProto_TYPE_STRING
	messages := []*fzMessage{
		newFzMessage("Basic").str("name").str("display_name").str("state").str("unmapped"),
		newFzMessage("Special").str("name").str("etag").stringMap("labels", strT).stringMap("annotations", strT).msg("meta", "Meta"),
		newFzMessage("Meta").str("name").str("etag").stringMap("labels", strT).str("owner"),
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
		newFzMessage("Widget").str("name").str("etag").stringMap("labels", strT).str("display_name").
			msg("create_time", ".google.protobuf.Timestamp").msg("config", "Config").msgs("rules", "Rule").str("network").str("unmapped"),
	}
	fdp := &descriptorpb.FileDescriptorProto{
		Name:       proto.String("google/cloud/fuzztest/v1/fuzztest.proto"),
		Package:    proto.String(fuzzerTestProtoPackage),
		Syntax:     proto.String("proto3"),
		Dependency: []string{"google/protobuf/duration.proto", "google/protobuf/struct.proto", "google/protobuf/timestamp.proto"},
		Options:    &descriptorpb.FileOptions{GoPackage: proto.String(fuzzerTestPBGoPackage + ";fuzztestpb")},
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
		t.Fatalf("building test proto: %v", err)
	}
	fd, _ := files.FindFileByPath(fdp.GetName())
	return fd.Messages().ByName(protoreflect.Name(name))
}

func writeFuzzerTestTree(t *testing.T, files map[string]string) string {
	t.Helper()
	root := t.TempDir()
	for name, content := range files {
		p := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func fuzzerTestKRMFile(body string) string {
	return `package v1alpha1
import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	othersvc "github.com/GoogleCloudPlatform/k8s-config-connector/apis/othersvc/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)
var ( _ refsv1beta1.ComputeNetworkRef; _ othersvc.Config; _ apiextensionsv1.JSON )
` + body
}

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
		fmt.Fprintf(&b, "func %s_FromProto(*direct.MapContext, *pb.%s) *krm.%s { return nil }\n", base, message, krmType)
		fmt.Fprintf(&b, "func %s_ToProto(*direct.MapContext, *krm.%s) *pb.%s { return nil }\n", base, krmType, message)
	}
	return b.String()
}

func planFuzzerForTest(t *testing.T, kind, message, version string, files map[string]string) (*FuzzerPlan, error) {
	t.Helper()
	root := writeFuzzerTestTree(t, files)
	direct, err := loadGoPackageIndex(filepath.Join(root, filepath.FromSlash(fuzzerTestDirectDir)))
	if err != nil {
		t.Fatal(err)
	}
	return planFuzzer(fuzzerPlanInput{
		kind: kind, version: version, message: fuzzerTestMessage(t, message), direct: direct,
		krm: newFuzzerKRMLoader(fuzzerTestAPIGoPackage, filepath.Join(root, "apis")),
	})
}

func TestPlanFuzzerClassifiesFields(t *testing.T) {
	for _, tc := range []struct {
		name         string
		message      string
		krm          string
		mappers      []string
		want         []string
		wantSpecOnly bool
	}{
		{"top-level", "Basic", "type BasicSpec struct { DisplayName *string }\ntype BasicObservedState struct { State *string }", nil,
			[]string{`SpecField(".display_name")`, `StatusField(".state")`, `Unimplemented_Identity(".name")`, `Unimplemented_NotYetTriaged(".unmapped")`}, false},
		{"special top-level", "Special", "type SpecialSpec struct { Meta *Meta }\ntype Meta struct { Owner *string }\ntype SpecialObservedState struct{}", nil,
			[]string{`SpecField(".meta")`, `Unimplemented_Identity(".name")`, `Unimplemented_Etag(".etag")`, `Unimplemented_LabelsAnnotations(".annotations")`, `Unimplemented_LabelsAnnotations(".labels")`, `Unimplemented_NotYetTriaged(".meta.etag")`, `Unimplemented_NotYetTriaged(".meta.labels")`, `Unimplemented_NotYetTriaged(".meta.name")`}, false},
		{"non-string map labels", "OddLabels", "type OddLabelsSpec struct { Description *string }\ntype OddLabelsObservedState struct{}", nil,
			[]string{`SpecField(".description")`, `Unimplemented_NotYetTriaged(".labels")`}, false},
		{"nested split", "Nested", "type NestedSpec struct { Config *ConfigSpec }\ntype ConfigSpec struct { Size *int32 }\ntype NestedObservedState struct { Config *ConfigObservedState }\ntype ConfigObservedState struct { State *string }", nil,
			[]string{`SpecField(".config.size")`, `StatusField(".config.state")`, `Unimplemented_NotYetTriaged(".config.extra")`}, false},
		{"repeated split", "Repeated", "type RepeatedSpec struct { Rules []Rule }\ntype Rule struct { Action *string }\ntype RepeatedObservedState struct { Rules []*RuleObservedState }\ntype RuleObservedState struct { HitCount *int64 }", nil,
			[]string{`SpecField(".rules[].action")`, `StatusField(".rules[].hit_count")`}, false},
		{"one-side message", "OneSide", "type OneSideSpec struct { SpecOnly *ConfigSpec }\ntype ConfigSpec struct { Size *int32 }\ntype OneSideObservedState struct { StatusOnly *ConfigObservedState }\ntype ConfigObservedState struct { State *string }", nil,
			[]string{`SpecField(".spec_only")`, `StatusField(".status_only")`, `Unimplemented_NotYetTriaged(".spec_only.extra")`, `Unimplemented_NotYetTriaged(".spec_only.state")`, `Unimplemented_NotYetTriaged(".status_only.extra")`, `Unimplemented_NotYetTriaged(".status_only.size")`}, false},
		{"references", "Refs", "type RefsSpec struct { NetworkRef *refsv1beta1.ComputeNetworkRef; SubnetworkRefs []refsv1beta1.ComputeSubnetworkRef }\ntype RefsObservedState struct{}", nil,
			[]string{`SpecField(".network")`, `SpecField(".subnetworks")`, `Unimplemented_NotYetTriaged(".kms_key")`}, false},
		{"well-known types", "WellKnown", "type WellKnownSpec struct { TTL *string; Params *apiextensionsv1.JSON; Other *othersvc.Config }\ntype WellKnownObservedState struct { CreateTime *string }", nil,
			[]string{`SpecField(".other")`, `SpecField(".params")`, `SpecField(".ttl")`, `StatusField(".create_time")`, `Unimplemented_NotYetTriaged(".update_time")`}, false},
		{"both sides is spec", "Both", "type BothSpec struct { Description *string; Config *ConfigSpec }\ntype ConfigSpec struct { Size *int32 }\ntype BothObservedState struct { Description *string; Config *ConfigObservedState }\ntype ConfigObservedState struct { Size *int32; State *string }", nil,
			[]string{`SpecField(".config.size")`, `SpecField(".description")`, `StatusField(".config.state")`, `Unimplemented_NotYetTriaged(".config.extra")`}, false},
		{"recursive KRM type", "Tree", "type TreeSpec struct { Root *Node }\ntype Node struct { Value *string; Children []Node }\ntype TreeObservedState struct{}", nil,
			[]string{`SpecField(".root")`, `Unimplemented_NotYetTriaged(".root.children")`, `Unimplemented_NotYetTriaged(".root.note")`}, false},
		{"list mapped to scalar", "Lists", "type ListsSpec struct { Tags *string; Zone []string; Rules *Rule }\ntype Rule struct { Action *string }\ntype ListsObservedState struct{}", nil,
			[]string{`SpecField(".zone")`, `Unimplemented_NotYetTriaged(".rules")`, `Unimplemented_NotYetTriaged(".tags")`}, false},
		{"spec-only with ObservedState", "Basic", "type BasicSpec struct { DisplayName *string }\ntype BasicObservedState struct { State *string }", []string{"BasicSpec"},
			[]string{`SpecField(".display_name")`, `StatusField(".state")`, `Unimplemented_Identity(".name")`, `Unimplemented_NotYetTriaged(".unmapped")`}, true},
		{"spec-only without ObservedState", "Basic", "type BasicSpec struct { DisplayName *string }", []string{"BasicSpec"},
			[]string{`SpecField(".display_name")`, `Unimplemented_Identity(".name")`, `Unimplemented_NotYetTriaged(".state")`, `Unimplemented_NotYetTriaged(".unmapped")`}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mappers := tc.mappers
			if mappers == nil {
				mappers = []string{tc.message + "Spec", tc.message + "ObservedState"}
			}
			plan, err := planFuzzerForTest(t, tc.message, tc.message, "v1alpha1", map[string]string{
				fuzzerTestKRMDir + "/types.go":               fuzzerTestKRMFile(tc.krm),
				fuzzerTestDirectDir + "/mapper.generated.go": fuzzerTestMapperFile(tc.message, mappers...),
			})
			if err != nil {
				t.Fatalf("planFuzzer: %v", err)
			}
			var got []string
			for _, e := range plan.Entries {
				got = append(got, fmt.Sprintf("%s(%q)", e.Method, e.Path))
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("entries differ (-want +got):\n%s", diff)
			}
			if (plan.Status == nil) != tc.wantSpecOnly {
				t.Errorf("spec-only = %v, want %v", plan.Status == nil, tc.wantSpecOnly)
			}
		})
	}
}

func TestPlanFuzzerFindsMappers(t *testing.T) {
	const krm = `
type BasicSpec struct { DisplayName *string }
type BasicObservedState struct { State *string }
type BasicStatus struct { ObservedState *CustomObservedStateStatus; State *string }
type CustomObservedStateStatus struct { State *string }
`
	for _, tc := range []struct {
		name                            string
		directFiles                     map[string]string
		wantSpec, wantStatus, wantError string
	}{
		{"versioned", map[string]string{"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicSpec_v1alpha1", "BasicObservedState_v1alpha1")}, "BasicSpec_v1alpha1_FromProto", "BasicObservedState_v1alpha1_FromProto", ""},
		{"handwritten+generated", map[string]string{"basic_mappers.go": fuzzerTestMapperFile("Basic", "BasicSpec"), "mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicObservedState")}, "BasicSpec_FromProto", "BasicObservedState_FromProto", ""},
		{"custom ObservedState", map[string]string{"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicSpec", "CustomObservedStateStatus")}, "BasicSpec_FromProto", "CustomObservedStateStatus_FromProto", ""},
		{"Status fallback", map[string]string{"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicSpec", "BasicStatus")}, "BasicSpec_FromProto", "BasicStatus_FromProto", ""},
		{"no Spec mappers", map[string]string{"mapper.generated.go": fuzzerTestMapperFile("Basic", "BasicObservedState")}, "", "", "no Spec mappers for kind Basic"},
		{"wrong proto message", map[string]string{"mapper.generated.go": fuzzerTestMapperFile("Nested", "BasicSpec")}, "", "", "BasicSpec_FromProto converts *Nested"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			files := map[string]string{fuzzerTestKRMDir + "/types.go": fuzzerTestKRMFile(krm)}
			for name, content := range tc.directFiles {
				files[fuzzerTestDirectDir+"/"+name] = content
			}
			plan, err := planFuzzerForTest(t, "Basic", "Basic", "v1alpha1", files)
			if tc.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantError) {
					t.Fatalf("error = %v, want %q", err, tc.wantError)
				}
				return
			}
			if err != nil || plan.Spec.FromProto != tc.wantSpec || plan.Status == nil || plan.Status.FromProto != tc.wantStatus {
				t.Errorf("got (%v, %v), want (%q, %q)", plan, err, tc.wantSpec, tc.wantStatus)
			}
		})
	}
}

const fuzzerTestWidgetKRM = `
type FuzzTestWidgetSpec struct { DisplayName *string; Config *ConfigSpec; Rules []Rule; NetworkRef *refsv1beta1.ComputeNetworkRef }
type ConfigSpec struct { Size *int32 }
type Rule struct { Action *string }
type FuzzTestWidgetObservedState struct { CreateTime *string; Config *ConfigObservedState; Rules []RuleObservedState }
type ConfigObservedState struct { State *string }
type RuleObservedState struct { HitCount *int64 }
`

func newFuzzerTestGenerator(root string) *FuzzerGenerator {
	return NewFuzzerGenerator(FuzzerGeneratorOptions{
		Group: "fuzztest.cnrm.cloud.google.com", Version: "v1alpha1", ProtoService: fuzzerTestProtoPackage,
		APIGoPackagePath: fuzzerTestAPIGoPackage + "/", APIDirectory: filepath.Join(root, "apis"),
		DirectDirectory: filepath.Join(root, filepath.FromSlash(fuzzerTestDirectDir)), Year: 2026,
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

func TestFuzzerGeneratorEndToEnd(t *testing.T) {
	t.Run("writes full fuzzer and preserves copyright year", func(t *testing.T) {
		root := writeFuzzerTestTree(t, widgetFuzzerTestFiles())
		g := newFuzzerTestGenerator(root)
		result, err := g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))
		if err != nil || result.Skipped {
			t.Fatalf("Generate: err=%v skipped=%v", err, result.Skipped)
		}
		got, _ := os.ReadFile(result.Path)
		for _, want := range []string{
			"// +generated:fuzzer\n// krm.group: fuzztest.cnrm.cloud.google.com\n// krm.version: v1alpha1\n// proto.message: google.cloud.fuzztest.v1.Widget\n// proto.service: google.cloud.fuzztest.v1",
			"fuzztesting.RegisterKRMFuzzer(fuzzTestWidgetFuzzer())",
			`f.SpecField(".config.size")`, `f.SpecField(".display_name")`, `f.SpecField(".network")`, `f.SpecField(".rules[].action")`,
			`f.StatusField(".config.state")`, `f.StatusField(".create_time")`, `f.StatusField(".rules[].hit_count")`,
			`f.Unimplemented_Identity(".name")`, `f.Unimplemented_Etag()`, `f.Unimplemented_LabelsAnnotations(".labels")`,
			`f.Unimplemented_NotYetTriaged(".config.extra")`, `f.Unimplemented_NotYetTriaged(".unmapped")`,
		} {
			if !strings.Contains(string(got), want) {
				t.Errorf("generated fuzzer missing %q:\n%s", want, got)
			}
		}
		earlier := strings.Replace(string(got), "Copyright 2026", "Copyright 2025", 1)
		_ = os.WriteFile(result.Path, []byte(earlier), 0644)
		_, _ = g.Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))
		got2, _ := os.ReadFile(result.Path)
		if diff := cmp.Diff(earlier, string(got2)); diff != "" {
			t.Errorf("regenerating changed fuzzer:\n%s", diff)
		}
	})

	t.Run("writes spec-only fuzzer and wires filter hooks", func(t *testing.T) {
		files := widgetFuzzerTestFiles("FuzzTestWidgetSpec")
		files[fuzzerTestDirectDir+"/fuzztestwidget_fuzzer_filters.go"] = "package fuzztest\nimport pb \"cloud.google.com/go/fuzztest/apiv1/fuzztestpb\"\nfunc fuzzTestWidgetFuzzerFilterSpec(in *pb.Widget) {}\nfunc fuzzTestWidgetFuzzerFilterStatus(in *pb.Widget) {}\n"
		root := writeFuzzerTestTree(t, files)
		result, err := newFuzzerTestGenerator(root).Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))
		if err != nil {
			t.Fatalf("Generate: %v", err)
		}
		got, _ := os.ReadFile(result.Path)
		for _, want := range []string{
			"fuzztesting.RegisterKRMSpecFuzzer(fuzzTestWidgetFuzzer())",
			`f.StatusField(".create_time")`,
			"f.FilterSpec = fuzzTestWidgetFuzzerFilterSpec\n\tf.FilterStatus = fuzzTestWidgetFuzzerFilterStatus",
		} {
			if !strings.Contains(string(got), want) {
				t.Errorf("missing %q in:\n%s", want, got)
			}
		}
	})

	t.Run("skips handwritten fuzzers and removes stale generated file", func(t *testing.T) {
		files := widgetFuzzerTestFiles()
		files[fuzzerTestDirectDir+"/fuzztestwidget_fuzzer.go"] = "package fuzztest\nimport \"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting\"\nfunc fuzzTestWidgetFuzzer() fuzztesting.KRMFuzzer { return nil }\n"
		stale := fuzzerTestDirectDir + "/fuzztestwidget_fuzzer.generated.go"
		files[stale] = "package fuzztest\n"
		root := writeFuzzerTestTree(t, files)
		result, err := newFuzzerTestGenerator(root).Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))
		if err != nil || !result.Skipped {
			t.Fatalf("got err=%v skipped=%v", err, result.Skipped)
		}
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(stale))); !os.IsNotExist(err) {
			t.Errorf("stale generated fuzzer was not removed: %v", err)
		}
	})

	t.Run("emits hint on conditionally guarded mapper field", func(t *testing.T) {
		files := widgetFuzzerTestFiles()
		files[fuzzerTestDirectDir+"/mapper.generated.go"] = `// +generated:mapper
package fuzztest
import (
	pb "cloud.google.com/go/fuzztest/apiv1/fuzztestpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/fuzztest/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func FuzzTestWidgetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Widget) *krm.FuzzTestWidgetSpec {
	if in == nil { return nil }
	out := &krm.FuzzTestWidgetSpec{}
	if in.GetNetwork() != "" { out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()} }
	if in.GetNetwork() == "SPECIAL" { out.DisplayName = direct.LazyPtr(in.GetDisplayName()) }
	return out
}
func FuzzTestWidgetSpec_ToProto(mapCtx *direct.MapContext, in *krm.FuzzTestWidgetSpec) *pb.Widget {
	if in == nil { return nil }
	out := &pb.Widget{DisplayName: direct.ValueOf(in.DisplayName)}
	if in.NetworkRef != nil { out.Network = in.NetworkRef.External }
	return out
}
`
		root := writeFuzzerTestTree(t, files)
		result, err := newFuzzerTestGenerator(root).Generate("FuzzTestWidget", fuzzerTestMessage(t, "Widget"))
		if err != nil {
			t.Fatalf("Generate: %v", err)
		}
		got, _ := os.ReadFile(result.Path)
		for _, want := range []string{
			`f.SpecField(".display_name") // HINT: mapper guards this field behind a value-dependent 'if' condition`,
			"f.SpecField(\".network\")\n",
		} {
			if !strings.Contains(string(got), want) {
				t.Errorf("missing expected line %q:\n%s", want, got)
			}
		}
	})
}

func TestFuzzerFuncName(t *testing.T) {
	for _, tc := range []struct{ kind, want string }{
		{"ComputeTargetVPNGateway", "computeTargetVPNGatewayFuzzer"},
		{"BigQueryDataPolicy", "bigQueryDataPolicyFuzzer"},
		{"IAMDenyPolicy", "iamDenyPolicyFuzzer"},
		{"SQLInstance", "sqlInstanceFuzzer"},
		{"KMS", "kmsFuzzer"},
	} {
		if got := fuzzerFuncName(tc.kind); got != tc.want {
			t.Errorf("fuzzerFuncName(%q) = %q, want %q", tc.kind, got, tc.want)
		}
	}
}
