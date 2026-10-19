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

package scaffold

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/sourcelinks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/apis"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// TestTypesTemplateRendersValidGo verifies that the types template renders valid Go
// code for both stub and prepopulated configurations.
func TestTypesTemplateRendersValidGo(t *testing.T) {
	base := apis.APIArgs{
		Group:                "networkservices.cnrm.cloud.google.com",
		Version:              "v1alpha1",
		Kind:                 "NetworkServicesLBTrafficExtension",
		ProtoResource:        "LbTrafficExtension",
		PackageProtoTag:      "google.cloud.networkservices.v1",
		KindProtoTag:         "google.cloud.networkservices.v1.LbTrafficExtension",
		ProtoMessageName:     "LbTrafficExtension",
		ProtoMessageFullName: "google.cloud.networkservices.v1.LbTrafficExtension",
	}
	base.RootRefField = fixedRootField("ProjectRef", "projectRef", "project")
	base.LocationField = legacyLocationField
	prepopulated := base
	prepopulated.LocationField = locationField
	prepopulated.SpecFields = "\t// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.description\n" +
		"\tDescription *string `json:\"description,omitempty\"`\n"
	prepopulated.ObservedStateFields = "\t// +kcc:proto:field=google.cloud.networkservices.v1.LbTrafficExtension.create_time\n" +
		"\tCreateTime *string `json:\"createTime,omitempty\"`\n"
	prepopulated.ExtraImports = []string{`common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"`}
	organizationRooted := prepopulated
	organizationRooted.RootRefField = fixedRootField("OrganizationRef", "organizationRef", "organization")
	selfRooted := base
	selfRooted.RootRefField = ""
	selfRooted.LocationField = ""

	for _, tc := range []struct {
		name string
		args apis.APIArgs
	}{
		{name: "stub", args: base},
		{name: "prepopulated", args: prepopulated},
		{name: "organization-rooted", args: organizationRooted},
		{name: "no root or location in the name, as for a Folder", args: selfRooted},
		{name: "GVK declared elsewhere", args: func() apis.APIArgs { a := base; a.SkipGVK = true; return a }()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			tmpl, err := template.New(tc.args.Kind).Funcs(funcMap).Parse(apis.TypesTemplate)
			if err != nil {
				t.Fatalf("parsing types template: %v", err)
			}

			// Act
			var buf bytes.Buffer
			execErr := tmpl.Execute(&buf, &tc.args)

			// Assert
			if execErr != nil {
				t.Fatalf("executing types template: %v", execErr)
			}
			if _, err := parser.ParseFile(token.NewFileSet(), "types.go", buf.Bytes(), parser.AllErrors); err != nil {
				t.Errorf("rendered types file is not valid Go: %v\n%s", err, buf.String())
			}
		})
	}
}

// TestAddTypeFileWritesPrepopulatedBodies verifies that AddTypeFile writes all
// components of PrepopulateResult (Spec fields, ObservedState fields, and ExtraImports)
// into the scaffolded file, and emits Location as *string when prepopulating vs
// string when not prepopulating.
func TestAddTypeFileWritesPrepopulatedBodies(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	scaffolder := &APIScaffolder{
		BaseDir:         dir,
		GoPackage:       "networkservices/v1alpha1",
		Group:           "networkservices.cnrm.cloud.google.com",
		Version:         "v1alpha1",
		PackageProtoTag: "google.cloud.networkservices.v1",
	}
	resource := options.Resource{Kind: "NetworkServicesLBTrafficExtension", ProtoName: "LbTrafficExtension"}
	prepopulated := &PrepopulateResult{
		SpecFields:          "\tDescription *string `json:\"description,omitempty\"`\n",
		ObservedStateFields: "\tError *common.Status `json:\"error,omitempty\"`\n",
		ExtraImports:        []string{`common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"`},
	}

	// Act
	if err := scaffolder.AddTypeFile(resource, prepopulated, ""); err != nil {
		t.Fatalf("AddTypeFile: %v", err)
	}

	// Assert
	b, err := os.ReadFile(filepath.Join(dir, scaffolder.GoPackage, "networkserviceslbtrafficextension_types.go"))
	if err != nil {
		t.Fatalf("reading scaffolded file: %v", err)
	}
	got := string(b)
	if !strings.Contains(got, "Location *string `json:\"location\"`") {
		t.Errorf("expected Location *string when prepopulating:\n%s", got)
	}
	if !strings.Contains(got, "Description *string `json:\"description,omitempty\"`") {
		t.Errorf("the Spec body is missing from the scaffolded file:\n%s", got)
	}
	if !strings.Contains(got, "Error *common.Status `json:\"error,omitempty\"`") {
		t.Errorf("the ObservedState body is missing from the scaffolded file:\n%s", got)
	}
	if !strings.Contains(got, `common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"`) {
		t.Errorf("the import the ObservedState body needs is missing from the scaffolded file:\n%s", got)
	}

	stubResource := options.Resource{Kind: "NetworkServicesStub", ProtoName: "Stub"}
	if err := scaffolder.AddTypeFile(stubResource, nil, ""); err != nil {
		t.Fatalf("AddTypeFile (stub): %v", err)
	}
	stubBytes, err := os.ReadFile(filepath.Join(dir, scaffolder.GoPackage, "networkservicesstub_types.go"))
	if err != nil {
		t.Fatalf("reading stub file: %v", err)
	}
	if !strings.Contains(string(stubBytes), "Location string `json:\"location\"`") {
		t.Errorf("expected Location string when not prepopulating:\n%s", string(stubBytes))
	}
}

func TestAddTypeFileCommentsOutEmptyObservedStateWhenPrepopulating(t *testing.T) {
	dir := t.TempDir()
	scaffolder := &APIScaffolder{
		BaseDir:         dir,
		GoPackage:       "composer/v1alpha1",
		Group:           "composer.cnrm.cloud.google.com",
		Version:         "v1alpha1",
		PackageProtoTag: "google.cloud.orchestration.airflow.service.v1",
	}
	resource := options.Resource{Kind: "ComposerUserWorkloadsSecret", ProtoName: "UserWorkloadsSecret"}
	prepopulated := &PrepopulateResult{
		SpecFields:          "\tData map[string]string `json:\"data,omitempty\"`\n",
		ObservedStateFields: "",
	}
	if err := scaffolder.AddTypeFile(resource, prepopulated, ""); err != nil {
		t.Fatalf("AddTypeFile: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(dir, scaffolder.GoPackage, "composeruserworkloadssecret_types.go"))
	if err != nil {
		t.Fatalf("reading scaffolded file: %v", err)
	}
	got := string(b)
	if !strings.Contains(got, "// ObservedState *ComposerUserWorkloadsSecretObservedState") {
		t.Errorf("expected ObservedState field to be commented out when empty:\n%s", got)
	}
	if !strings.Contains(got, "// type ComposerUserWorkloadsSecretObservedState struct {") {
		t.Errorf("expected ObservedState struct to be commented out when empty:\n%s", got)
	}
}

// TestPackageDeclaresGVK pins which files count as declaring <Kind>GVK: a
// top-level var for this exact Kind, in any Go file but a _types.go.
// AddTypeFile sets SkipGVK from the answer.
func TestPackageDeclaresGVK(t *testing.T) {
	for _, tc := range []struct {
		name  string
		files map[string]string
		want  bool
	}{
		{
			name:  "a reference file declares it",
			files: map[string]string{"widget_reference.go": "package v1\n\nvar WidgetGVK = GroupVersion.WithKind(\"Widget\")\n"},
			want:  true,
		},
		{
			name:  "a types file declaring it is the one being replaced",
			files: map[string]string{"widget_types.go": "package v1\n\nvar WidgetGVK = GroupVersion.WithKind(\"Widget\")\n"},
			want:  false,
		},
		{
			name:  "another Kind's GVK whose name starts with this one",
			files: map[string]string{"widgetpart_reference.go": "package v1\n\nvar WidgetPartGVK = GroupVersion.WithKind(\"WidgetPart\")\n"},
			want:  false,
		},
		{
			name:  "a mention that is not a declaration",
			files: map[string]string{"widget_identity.go": "package v1\n\nfunc f() { _ = WidgetGVK }\n"},
			want:  false,
		},
		{
			name: "no package directory yet",
			want: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			dir := filepath.Join(t.TempDir(), "v1")
			if tc.files != nil {
				if err := os.MkdirAll(dir, 0o755); err != nil {
					t.Fatal(err)
				}
			}
			for name, body := range tc.files {
				if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
					t.Fatal(err)
				}
			}

			// Act
			got := packageDeclaresGVK(dir, "Widget")

			// Assert
			if got != tc.want {
				t.Errorf("packageDeclaresGVK() = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestAddTypeFileWritesSourceLinks checks that the source link block lands
// between the license and the package clause, with a blank line on each side
// so Go does not read it as the package doc, and that no block is written
// without it.
func TestAddTypeFileWritesSourceLinks(t *testing.T) {
	block := sourcelinks.Render("NetworkServicesLBTrafficExtension", []sourcelinks.Link{
		{Key: sourcelinks.KeyProto, URL: "https://github.com/googleapis/googleapis/blob/abc/google/cloud/networkservices/v1/dep.proto"},
		{Key: sourcelinks.KeyServiceDocs, URL: "https://cloud.google.com/networking/"},
		{Key: sourcelinks.KeyResourceDocs, URL: "https://docs.cloud.google.com/service-extensions/docs/reference/rest/v1/projects.locations.lbTrafficExtensions",
			Guess: judgement.ReasonVerifyResourceDocsLink},
	})
	for _, withLinks := range []bool{true, false} {
		t.Run(fmt.Sprintf("withLinks=%v", withLinks), func(t *testing.T) {
			dir := t.TempDir()
			scaffolder := &APIScaffolder{
				BaseDir:         dir,
				GoPackage:       "networkservices/v1alpha1",
				Group:           "networkservices.cnrm.cloud.google.com",
				Version:         "v1alpha1",
				PackageProtoTag: "google.cloud.networkservices.v1",
			}
			resource := options.Resource{Kind: "NetworkServicesLBTrafficExtension", ProtoName: "LbTrafficExtension"}
			links := ""
			if withLinks {
				links = block
			}
			if err := scaffolder.AddTypeFile(resource, nil, links); err != nil {
				t.Fatalf("AddTypeFile: %v", err)
			}
			b, err := os.ReadFile(scaffolder.PathToTypeFile(resource))
			if err != nil {
				t.Fatal(err)
			}
			got := string(b)
			f, err := parser.ParseFile(token.NewFileSet(), "types.go", b, parser.ParseComments)
			if err != nil {
				t.Fatalf("scaffolded file is not valid Go: %v", err)
			}
			if f.Doc != nil {
				t.Errorf("the file has a package doc comment, so the block is attached to the package clause:\n%s", got)
			}
			h := sourcelinks.Parse(b)
			if !withLinks {
				if h.Kind != "" || len(h.Links) != 0 {
					t.Errorf("source links were written without being asked for:\n%s", got)
				}
				return
			}
			if !strings.Contains(got, "limitations under the License.\n\n"+block+"\n\npackage v1alpha1\n") {
				t.Errorf("the block is not between the license and the package clause:\n%s", got)
			}
			if h.Kind != resource.Kind || len(h.Links) != 3 {
				t.Errorf("Parse = %+v", h)
			}
		})
	}
}

// TestAddTypeFileWritesRequiredFromProtoMarker checks that a Kind scaffolded
// with RequiredFromProto has the marker in its Spec's doc comment, and that a
// Kind scaffolded without it doesn't.
func TestAddTypeFileWritesRequiredFromProtoMarker(t *testing.T) {
	for _, tc := range []struct {
		name              string
		requiredFromProto bool
	}{
		{name: "with the flag", requiredFromProto: true},
		{name: "without the flag", requiredFromProto: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			dir := t.TempDir()
			scaffolder := &APIScaffolder{
				BaseDir:           dir,
				GoPackage:         "networkservices/v1alpha1",
				Group:             "networkservices.cnrm.cloud.google.com",
				Version:           "v1alpha1",
				PackageProtoTag:   "google.cloud.networkservices.v1",
				RequiredFromProto: tc.requiredFromProto,
			}
			resource := options.Resource{Kind: "NetworkServicesLBTrafficExtension", ProtoName: "LbTrafficExtension"}

			// Act
			if err := scaffolder.AddTypeFile(resource, nil, ""); err != nil {
				t.Fatalf("AddTypeFile: %v", err)
			}

			// Assert
			p := filepath.Join(dir, scaffolder.GoPackage, "networkserviceslbtrafficextension_types.go")
			f, err := parser.ParseFile(token.NewFileSet(), p, nil, parser.ParseComments)
			if err != nil {
				t.Fatalf("parsing scaffolded file: %v", err)
			}
			marked := false
			found := false
			ast.Inspect(f, func(n ast.Node) bool {
				gd, ok := n.(*ast.GenDecl)
				if !ok || gd.Tok != token.TYPE {
					return true
				}
				for _, spec := range gd.Specs {
					if ts, ok := spec.(*ast.TypeSpec); ok && ts.Name.Name == resource.Kind+"Spec" {
						found = true
						if gd.Doc != nil {
							for _, c := range gd.Doc.List {
								marked = marked || codegen.IsRequiredFromProtoMarker(c.Text)
							}
						}
					}
				}
				return false
			})
			if !found {
				t.Fatalf("%sSpec is not declared in %s", resource.Kind, p)
			}
			if marked != tc.requiredFromProto {
				b, _ := os.ReadFile(p)
				t.Errorf("%sSpec has the %s marker = %v, want %v:\n%s", resource.Kind, codegen.RequiredFromProtoMarker, marked, tc.requiredFromProto, b)
			}
		})
	}
}

func TestBuildAPIArgsMultiServiceProtoTag(t *testing.T) {
	msgOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(msgOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "spanner.googleapis.com/InstancePartition",
		Pattern: []string{"projects/{project}/instances/{instance}/instancePartitions/{instance_partition}"},
	})
	fds := &descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{
			{
				Name:    proto.String("spanner_database.proto"),
				Package: proto.String("google.spanner.admin.database.v1"),
				MessageType: []*descriptorpb.DescriptorProto{
					{Name: proto.String("Database")},
				},
			},
			{
				Name:    proto.String("spanner_instance.proto"),
				Package: proto.String("google.spanner.admin.instance.v1"),
				MessageType: []*descriptorpb.DescriptorProto{
					{Name: proto.String("InstancePartition"), Options: msgOpts},
				},
			},
		},
	}
	tmpProto := filepath.Join(t.TempDir(), "test.pb")
	rawBytes, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("proto.Marshal: %v", err)
	}
	if err := os.WriteFile(tmpProto, rawBytes, 0o644); err != nil {
		t.Fatalf("os.WriteFile: %v", err)
	}
	loadedProto, err := protoapi.LoadProto(tmpProto, "")
	if err != nil {
		t.Fatalf("LoadProto: %v", err)
	}

	scaffolder := &APIScaffolder{
		BaseDir:         t.TempDir(),
		GoPackage:       "spanner/v1alpha1",
		Group:           "spanner.cnrm.cloud.google.com",
		Version:         "v1alpha1",
		PackageProtoTag: "google.spanner.admin.database.v1,google.spanner.admin.instance.v1",
		Proto:           loadedProto,
	}
	res := options.Resource{Kind: "SpannerInstancePartition", ProtoName: "InstancePartition"}
	args := scaffolder.buildAPIArgs(&res)
	const wantFull = "google.spanner.admin.instance.v1.InstancePartition"
	if args.KindProtoTag != wantFull {
		t.Errorf("KindProtoTag = %q, want %q", args.KindProtoTag, wantFull)
	}
	if args.ProtoMessageFullName != wantFull {
		t.Errorf("ProtoMessageFullName = %q, want %q", args.ProtoMessageFullName, wantFull)
	}
	const wantPat = "projects/{project}/instances/{instance}/instancePartitions/{instance_partition}"
	if args.ResourcePattern != wantPat {
		t.Errorf("ResourcePattern = %q, want %q", args.ResourcePattern, wantPat)
	}
}

func TestAddTypeFilePluralizesShortName(t *testing.T) {
	for _, tc := range []struct {
		name     string
		service  string
		protoPkg string
		resource options.Resource
		want     string
	}{
		{
			name:     "consonant + y",
			service:  "jobs",
			protoPkg: "google.cloud.talent.v4",
			resource: options.Resource{Kind: "JobsCompany", ProtoName: "Company"},
			want:     "// +kubebuilder:resource:categories=gcp,shortName=gcpjobscompany;gcpjobscompanies",
		},
		{
			name:     "kind already plural",
			service:  "iap",
			protoPkg: "google.cloud.iap.v1",
			resource: options.Resource{Kind: "IAPSettings", ProtoName: "IapSettings"},
			want:     "// +kubebuilder:resource:categories=gcp,shortName=gcpiapsettings;gcpiapsettings",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			scaffolder := &APIScaffolder{
				BaseDir:         t.TempDir(),
				GoPackage:       tc.service + "/v1alpha1",
				Group:           tc.service + ".cnrm.cloud.google.com",
				Version:         "v1alpha1",
				PackageProtoTag: tc.protoPkg,
			}
			if err := scaffolder.AddTypeFile(tc.resource, nil, ""); err != nil {
				t.Fatalf("AddTypeFile: %v", err)
			}
			b, err := os.ReadFile(scaffolder.PathToTypeFile(tc.resource))
			if err != nil {
				t.Fatalf("reading scaffolded file: %v", err)
			}
			if !strings.Contains(string(b), tc.want) {
				t.Errorf("expected %q in scaffolded file:\n%s", tc.want, string(b))
			}
		})
	}
}
