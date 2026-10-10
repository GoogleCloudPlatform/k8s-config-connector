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
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestIdentityGenerator_StandardProjectLocation(t *testing.T) {
	tmp := t.TempDir()
	pkgDir := filepath.Join(tmp, "apis", "networkconnectivity", "v1alpha1")
	writeFakeTypesFile(t, pkgDir, "v1alpha1", `
import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkConnectivityServiceConnectionMapSpec struct {
	ProjectRef *parent.ProjectRef `+"`"+`json:"projectRef"`+"`"+`
	Location string `+"`"+`json:"location"`+"`"+`
	ResourceID *string `+"`"+`json:"resourceID,omitempty"`+"`"+`
}

type NetworkConnectivityServiceConnectionMapStatus struct {
	ExternalRef *string `+"`"+`json:"externalRef,omitempty"`+"`"+`
}

type NetworkConnectivityServiceConnectionMap struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec NetworkConnectivityServiceConnectionMapSpec
	Status NetworkConnectivityServiceConnectionMapStatus
}
`)

	msg := buildTestResourceDescriptor(
		t,
		"google.cloud.networkconnectivity.v1",
		"ServiceConnectionMap",
		"networkconnectivity.googleapis.com/ServiceConnectionMap",
		[]string{"projects/{project}/locations/{location}/serviceConnectionMaps/{service_connection_map}"},
	)

	gen := NewIdentityGenerator(IdentityGeneratorOptions{
		Group:        "networkconnectivity.cnrm.cloud.google.com",
		Version:      "v1alpha1",
		ProtoService: "google.cloud.networkconnectivity.v1",
		APIDirectory: filepath.Join(tmp, "apis"),
		EmitTests:    true,
	})

	res, err := gen.Generate("NetworkConnectivityServiceConnectionMap", msg)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if res.IdentitySkipped || res.ReferenceSkipped || res.TestSkipped {
		t.Fatalf("expected all 3 files emitted, got skipped=(%v, %v, %v)", res.IdentitySkipped, res.ReferenceSkipped, res.TestSkipped)
	}

	idBytes, err := os.ReadFile(res.IdentityPath)
	if err != nil {
		t.Fatalf("ReadFile identity: %v", err)
	}
	assertValidGoSource(t, res.IdentityPath, idBytes)
	idSrc := string(idBytes)

	for _, want := range []string{
		`gcpurls.Template[NetworkConnectivityServiceConnectionMapIdentity]("networkconnectivity.googleapis.com", "projects/{project}/locations/{location}/serviceConnectionMaps/{serviceConnectionMap}")`,
		"_ identity.IdentityV2 = &NetworkConnectivityServiceConnectionMapIdentity{}",
		"func (i *NetworkConnectivityServiceConnectionMapIdentity) ParentString() string",
		"func (obj *NetworkConnectivityServiceConnectionMap) GetIdentity(",
	} {
		if !strings.Contains(idSrc, want) {
			t.Errorf("identity output missing %q:\n%s", want, idSrc)
		}
	}
	if strings.Contains(idSrc, "HasIdentitySpecified") {
		t.Errorf("did not expect ServerGeneratedIdentity for client-assignable resource:\n%s", idSrc)
	}

	refBytes, err := os.ReadFile(res.ReferencePath)
	if err != nil {
		t.Fatalf("ReadFile reference: %v", err)
	}
	assertValidGoSource(t, res.ReferencePath, refBytes)
	refSrc := string(refBytes)
	for _, want := range []string{
		"var NetworkConnectivityServiceConnectionMapGVK = schema.GroupVersionKind{",
		"refs.Register(&NetworkConnectivityServiceConnectionMapRef{}, &NetworkConnectivityServiceConnectionMap{})",
		"projects/{{projectID}}/locations/{{location}}/serviceConnectionMaps/{{serviceConnectionMapID}}",
	} {
		if !strings.Contains(refSrc, want) {
			t.Errorf("reference output missing %q:\n%s", want, refSrc)
		}
	}

	testBytes, err := os.ReadFile(res.TestPath)
	if err != nil {
		t.Fatalf("ReadFile test: %v", err)
	}
	assertValidGoSource(t, res.TestPath, testBytes)
	if !strings.Contains(string(testBytes), "identity.AssertConformance") {
		t.Errorf("test output missing identity.AssertConformance:\n%s", string(testBytes))
	}
}

func TestIdentityGenerator_ServerGeneratedIDAndMultiPattern(t *testing.T) {
	tmp := t.TempDir()
	pkgDir := filepath.Join(tmp, "apis", "cloudbuild", "v1beta1")
	writeFakeTypesFile(t, pkgDir, "v1beta1", `
import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CloudBuildTriggerSpec struct {
	ProjectRef *parent.ProjectRef `+"`"+`json:"projectRef"`+"`"+`
	Location string `+"`"+`json:"location"`+"`"+`
	// The Trigger ID. Server-generated if not specified.
	ResourceID *string `+"`"+`json:"resourceID,omitempty"`+"`"+`
}

type CloudBuildTriggerStatus struct {
	ExternalRef *string `+"`"+`json:"externalRef,omitempty"`+"`"+`
}

type CloudBuildTrigger struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec CloudBuildTriggerSpec
	Status CloudBuildTriggerStatus
}
`)

	msg := buildTestResourceDescriptor(
		t,
		"google.devtools.cloudbuild.v1",
		"BuildTrigger",
		"cloudbuild.googleapis.com/BuildTrigger",
		[]string{
			"projects/{project}/triggers/{trigger}",
			"projects/{project}/locations/{location}/triggers/{trigger}",
		},
	)

	gen := NewIdentityGenerator(IdentityGeneratorOptions{
		Group:        "cloudbuild.cnrm.cloud.google.com",
		Version:      "v1beta1",
		ProtoService: "google.devtools.cloudbuild.v1",
		APIDirectory: filepath.Join(tmp, "apis"),
		EmitTests:    true,
	})

	plan, err := gen.Plan("CloudBuildTrigger", msg)
	if err != nil {
		t.Fatalf("Plan failed: %v", err)
	}
	if plan.TemplatePattern != "projects/{project}/locations/{location}/triggers/{trigger}" {
		t.Fatalf("expected location pattern selected from Spec, got %q", plan.TemplatePattern)
	}
	if !plan.ServerGeneratedID {
		t.Fatalf("expected ServerGeneratedID=true from comment")
	}
	if len(plan.Judgement) != 1 || plan.Judgement[0].Reason != judgement.ReasonIdentityMultiPattern {
		t.Fatalf("expected [%s] judgement entry, got %+v", judgement.ReasonIdentityMultiPattern, plan.Judgement)
	}

	idBytes, emitted, err := gen.RenderIdentityFile(plan, nil)
	if err != nil || !emitted {
		t.Fatalf("RenderIdentityFile failed: emitted=%v err=%v", emitted, err)
	}
	assertValidGoSource(t, "cloudbuildtrigger_identity.generated.go", idBytes)
	if !strings.Contains(string(idBytes), "_ identity.ServerGeneratedIdentity = &CloudBuildTriggerIdentity{}") {
		t.Errorf("missing ServerGeneratedIdentity assertion:\n%s", string(idBytes))
	}
	if !strings.Contains(string(idBytes), "func (i *CloudBuildTriggerIdentity) HasIdentitySpecified() bool") {
		t.Errorf("missing HasIdentitySpecified method:\n%s", string(idBytes))
	}
}

func TestIdentityGenerator_NestedParentRefAndPartialOverride(t *testing.T) {
	tmp := t.TempDir()
	pkgDir := filepath.Join(tmp, "apis", "spanner", "v1beta1")
	writeFakeTypesFile(t, pkgDir, "v1beta1", `
import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SpannerInstanceRef struct {
	External string
}

type SpannerDatabaseSpec struct {
	InstanceRef *SpannerInstanceRef `+"`"+`json:"instanceRef"`+"`"+`
	ResourceID *string `+"`"+`json:"resourceID,omitempty"`+"`"+`
}

type SpannerDatabaseStatus struct {
	ExternalRef *string `+"`"+`json:"externalRef,omitempty"`+"`"+`
}

type SpannerDatabase struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec SpannerDatabaseSpec
	Status SpannerDatabaseStatus
}
`)

	msg := buildTestResourceDescriptor(
		t,
		"google.spanner.admin.database.v1",
		"Database",
		"spanner.googleapis.com/Database",
		[]string{"projects/{project}/instances/{instance}/databases/{database}"},
	)

	gen := NewIdentityGenerator(IdentityGeneratorOptions{
		Group:        "spanner.cnrm.cloud.google.com",
		Version:      "v1beta1",
		ProtoService: "google.spanner.admin.database.v1",
		APIDirectory: filepath.Join(tmp, "apis"),
	})

	plan, err := gen.Plan("SpannerDatabase", msg)
	if err != nil {
		t.Fatalf("Plan failed: %v", err)
	}
	if plan.ParentRef == nil || plan.ParentRef.SpecFieldName != "InstanceRef" {
		t.Fatalf("expected ParentRef InstanceRef, got %+v", plan.ParentRef)
	}

	// Verify full generation uses InstanceRef.NormalizedExternal.
	fullBytes, emitted, err := gen.RenderIdentityFile(plan, nil)
	if err != nil || !emitted {
		t.Fatalf("RenderIdentityFile failed: %v", err)
	}
	assertValidGoSource(t, "spannerdatabase_identity.generated.go", fullBytes)
	if !strings.Contains(string(fullBytes), "parentRef.Normalize(ctx, reader, obj.GetNamespace())") {
		t.Errorf("expected NormalizedExternal call on InstanceRef:\n%s", string(fullBytes))
	}

	// Verify partial hand-written override: if getIdentityFromSpannerDatabaseSpec is hand-written,
	// RenderIdentityFile omits only that function while still emitting the struct, methods, and GetIdentity.
	ov := &HandWrittenOverrides{
		HasGetIdentityFromSpec: true,
	}
	partialBytes, emitted, err := gen.RenderIdentityFile(plan, ov)
	if err != nil || !emitted {
		t.Fatalf("RenderIdentityFile partial failed: %v", err)
	}
	assertValidGoSource(t, "spannerdatabase_identity.generated.go", partialBytes)
	partialSrc := string(partialBytes)
	if strings.Contains(partialSrc, "func getIdentityFromSpannerDatabaseSpec(") {
		t.Errorf("expected hand-written getIdentityFromSpannerDatabaseSpec to be omitted:\n%s", partialSrc)
	}
	if !strings.Contains(partialSrc, "func (obj *SpannerDatabase) GetIdentity(") {
		t.Errorf("expected GetIdentity to still be emitted:\n%s", partialSrc)
	}
}

func TestIdentityGenerator_PatternOverride(t *testing.T) {
	tmp := t.TempDir()
	pkgDir := filepath.Join(tmp, "apis", "storage", "v1beta1")
	writeFakeTypesFile(t, pkgDir, "v1beta1", `
import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StorageBucketSpec struct {
	ProjectRef *parent.ProjectRef `+"`"+`json:"projectRef"`+"`"+`
	ResourceID *string `+"`"+`json:"resourceID,omitempty"`+"`"+`
}

type StorageBucketStatus struct {
	ExternalRef *string `+"`"+`json:"externalRef,omitempty"`+"`"+`
}

type StorageBucket struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec StorageBucketSpec
	Status StorageBucketStatus
}
`)

	// Build a proto message WITHOUT (google.api.resource) options.
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("storage.proto"),
		Package: proto.String("google.storage.v2"),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("Bucket"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("name"),
						Number: proto.Int32(1),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("protodesc.NewFile: %v", err)
	}
	msg := fd.Messages().ByName("Bucket")

	// 1. Without PatternOverrides, Plan skips and emits identity-root-unknown.
	genNoOverride := NewIdentityGenerator(IdentityGeneratorOptions{
		Group:        "storage.cnrm.cloud.google.com",
		Version:      "v1beta1",
		ProtoService: "google.storage.v2",
		APIDirectory: filepath.Join(tmp, "apis"),
		EmitTests:    true,
	})
	planSkipped, err := genNoOverride.Plan("StorageBucket", msg)
	if err != nil {
		t.Fatalf("Plan without override failed: %v", err)
	}
	if planSkipped.TemplatePattern != "" {
		t.Fatalf("expected Plan without override to be skipped")
	}
	if len(planSkipped.Judgement) != 1 || planSkipped.Judgement[0].Reason != judgement.ReasonIdentityRootUnknown {
		t.Fatalf("expected [%s] judgement entry when skipped, got %+v", judgement.ReasonIdentityRootUnknown, planSkipped.Judgement)
	}

	// 2. With PatternOverrides (--pattern StorageBucket=//storage.googleapis.com/projects/{project}/buckets/{bucket}),
	// Generate succeeds with zero judgement entries and emits all 3 files.
	genWithOverride := NewIdentityGenerator(IdentityGeneratorOptions{
		Group:        "storage.cnrm.cloud.google.com",
		Version:      "v1beta1",
		ProtoService: "google.storage.v2",
		APIDirectory: filepath.Join(tmp, "apis"),
		EmitTests:    true,
		PatternOverrides: map[string]string{
			"StorageBucket": "//storage.googleapis.com/projects/{project}/buckets/{bucket}",
		},
	})
	res, err := genWithOverride.Generate("StorageBucket", msg)
	if err != nil {
		t.Fatalf("Generate with override failed: %v", err)
	}
	if res.IdentitySkipped || res.ReferenceSkipped || res.TestSkipped {
		t.Fatalf("expected all 3 files emitted with override, got skipped=(%v, %v, %v)", res.IdentitySkipped, res.ReferenceSkipped, res.TestSkipped)
	}
	if len(res.Plan.Judgement) != 0 {
		t.Fatalf("expected zero judgement entries when explicit --pattern override is provided, got %+v", res.Plan.Judgement)
	}
	idBytes, err := os.ReadFile(res.IdentityPath)
	if err != nil {
		t.Fatalf("ReadFile identity: %v", err)
	}
	assertValidGoSource(t, res.IdentityPath, idBytes)
	if !strings.Contains(string(idBytes), `gcpurls.Template[StorageBucketIdentity]("storage.googleapis.com", "projects/{project}/buckets/{bucket}")`) {
		t.Errorf("unexpected template in generated identity:\n%s", string(idBytes))
	}
}


func TestIdentityGenerator_SingleParentToMultiParentRerun(t *testing.T) {
	tmp := t.TempDir()
	pkgDir := filepath.Join(tmp, "apis", "privilegedaccessmanager", "v1beta1")

	// Pass 1: initial types only declare ProjectRef.
	writeFakeTypesFile(t, pkgDir, "v1beta1", `
import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PrivilegedAccessManagerEntitlementSpec struct {
	ProjectRef *parent.ProjectRef ` + "`" + `json:"projectRef,omitempty"` + "`" + `
	Location *string ` + "`" + `json:"location"` + "`" + `
	ResourceID *string ` + "`" + `json:"resourceID,omitempty"` + "`" + `
}

type PrivilegedAccessManagerEntitlementStatus struct {
	ExternalRef *string ` + "`" + `json:"externalRef,omitempty"` + "`" + `
}

type PrivilegedAccessManagerEntitlement struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec PrivilegedAccessManagerEntitlementSpec
	Status PrivilegedAccessManagerEntitlementStatus
}
`)

	msg := buildTestResourceDescriptor(
		t,
		"google.cloud.privilegedaccessmanager.v1",
		"Entitlement",
		"privilegedaccessmanager.googleapis.com/Entitlement",
		[]string{
			"projects/{project}/locations/{location}/entitlements/{entitlement}",
			"folders/{folder}/locations/{location}/entitlements/{entitlement}",
			"organizations/{organization}/locations/{location}/entitlements/{entitlement}",
		},
	)

	gen := NewIdentityGenerator(IdentityGeneratorOptions{
		Group:        "privilegedaccessmanager.cnrm.cloud.google.com",
		Version:      "v1beta1",
		ProtoService: "google.cloud.privilegedaccessmanager.v1",
		APIDirectory: filepath.Join(tmp, "apis"),
		EmitTests:    true,
	})

	res1, err := gen.Generate("PrivilegedAccessManagerEntitlement", msg)
	if err != nil {
		t.Fatalf("Pass 1 Generate failed: %v", err)
	}
	if len(res1.Plan.MultiParentVariants) != 0 {
		t.Fatalf("Pass 1 expected single-parent plan, got %d multi-parent plans", len(res1.Plan.MultiParentVariants))
	}
	idBytes1, err := os.ReadFile(res1.IdentityPath)
	if err != nil {
		t.Fatalf("ReadFile Pass 1 identity: %v", err)
	}
	assertValidGoSource(t, res1.IdentityPath, idBytes1)
	if !strings.Contains(string(idBytes1), "PrivilegedAccessManagerEntitlementIdentityFormat = gcpurls.Template") {
		t.Errorf("Pass 1 expected single-parent template var:\n%s", string(idBytes1))
	}

	// Pass 2: corrective types pass adds FolderRef and OrganizationRef, then re-runs Generate.
	writeFakeTypesFile(t, pkgDir, "v1beta1", `
import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PrivilegedAccessManagerEntitlementSpec struct {
	ProjectRef *parent.ProjectRef ` + "`" + `json:"projectRef,omitempty"` + "`" + `
	FolderRef *parent.FolderRef ` + "`" + `json:"folderRef,omitempty"` + "`" + `
	OrganizationRef *parent.OrganizationRef ` + "`" + `json:"organizationRef,omitempty"` + "`" + `
	Location *string ` + "`" + `json:"location"` + "`" + `
	ResourceID *string ` + "`" + `json:"resourceID,omitempty"` + "`" + `
}

type PrivilegedAccessManagerEntitlementStatus struct {
	ExternalRef *string ` + "`" + `json:"externalRef,omitempty"` + "`" + `
}

type PrivilegedAccessManagerEntitlement struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec PrivilegedAccessManagerEntitlementSpec
	Status PrivilegedAccessManagerEntitlementStatus
}
`)

	res2, err := gen.Generate("PrivilegedAccessManagerEntitlement", msg)
	if err != nil {
		t.Fatalf("Pass 2 Generate failed: %v", err)
	}
	if res2.IdentitySkipped || res2.ReferenceSkipped || res2.TestSkipped {
		t.Fatalf("Pass 2 expected previously generated files to be overwritten, got skipped=(%v, %v, %v)", res2.IdentitySkipped, res2.ReferenceSkipped, res2.TestSkipped)
	}
	if len(res2.Plan.MultiParentVariants) != 3 {
		t.Fatalf("Pass 2 expected 3 multi-parent plans, got %d", len(res2.Plan.MultiParentVariants))
	}
	idBytes2, err := os.ReadFile(res2.IdentityPath)
	if err != nil {
		t.Fatalf("ReadFile Pass 2 identity: %v", err)
	}
	assertValidGoSource(t, res2.IdentityPath, idBytes2)
	idSrc2 := string(idBytes2)
	for _, want := range []string{
		"ProjectPrivilegedAccessManagerEntitlementIdentityFormat",
		"FolderPrivilegedAccessManagerEntitlementIdentityFormat",
		"OrganizationPrivilegedAccessManagerEntitlementIdentityFormat",
		"if i.Project != \"\"",
		"if i.Folder != \"\"",
		"if i.Organization != \"\"",
		"at most one of spec.projectRef, spec.folderRef, spec.organizationRef can be set",
	} {
		if !strings.Contains(idSrc2, want) {
			t.Errorf("Pass 2 identity output missing %q:\n%s", want, idSrc2)
		}
	}
}

func buildTestResourceDescriptor(t *testing.T, pkg, msgName, resType string, patterns []string) protoreflect.MessageDescriptor {
	t.Helper()
	msgOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(msgOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    resType,
		Pattern: patterns,
	})
	fdp := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test.proto"),
		Package: proto.String(pkg),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name:    proto.String(msgName),
				Options: msgOpts,
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   proto.String("name"),
						Number: proto.Int32(1),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
				},
			},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("protodesc.NewFile: %v", err)
	}
	return fd.Messages().ByName(protoreflect.Name(msgName))
}

func writeFakeTypesFile(t *testing.T, dir, pkgName, content string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "types_test_fixture.go"), []byte("package "+pkgName+"\n\n"+content), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
}

func assertValidGoSource(t *testing.T, filename string, src []byte) {
	t.Helper()
	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, filename, src, parser.AllErrors); err != nil {
		t.Fatalf("generated %s is not valid Go:\n%s\nerr: %v", filename, string(src), err)
	}
}
