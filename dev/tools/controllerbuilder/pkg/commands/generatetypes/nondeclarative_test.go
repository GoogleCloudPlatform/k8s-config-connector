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
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestNonDeclarativePatternReason(t *testing.T) {
	for _, tc := range []struct {
		pattern    string
		wantMatch  bool
		wantSubstr string
	}{
		// Declarative control-plane patterns
		{
			pattern:   "projects/{project}/locations/{location}/widgets/{widget}",
			wantMatch: false,
		},
		{
			// "versions" is an ancestor segment, not the leaf collection ("specs")
			pattern:   "projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec}",
			wantMatch: false,
		},
		{
			// Wildcard "*" segments are ignored just like "{param}" segments
			pattern:   "projects/*/locations/*/apis/*/versions/*/specs/*",
			wantMatch: false,
		},
		{
			pattern:    "projects/*/locations/*/agents/*/versions/*",
			wantMatch:  true,
			wantSubstr: `version snapshot leaf collection "versions"`,
		},
		{
			pattern:   "projects/{project}/locations/{location}/backupPolicies/{backup_policy}",
			wantMatch: false,
		},
		{
			pattern:   "projects/{project}/locations/{location}/backupVaults/{backup_vault}",
			wantMatch: false,
		},
		// Criterion 2: Ephemeral runtime parent / leaf segments
		{
			pattern:    "projects/{project}/locations/{location}/reasoningEngines/{reasoning_engine}/sessions/{session}",
			wantMatch:  true,
			wantSubstr: `ephemeral runtime collection "sessions"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/sessions/{session}/events/{event}",
			wantMatch:  true,
			wantSubstr: `ephemeral runtime collection "sessions"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/conversations/{conversation}/participants/{participant}",
			wantMatch:  true,
			wantSubstr: `ephemeral runtime collection "conversations"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/processes/{process}/runs/{run}",
			wantMatch:  true,
			wantSubstr: `ephemeral execution collection "runs"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/repositories/{repository}/pullRequests/{pull_request}/pullRequestComments/{pull_request_comment}",
			wantMatch:  true,
			wantSubstr: `ephemeral activity collection "pullRequests"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/repositories/{repository}/issues/{issue}/issueComments/{issue_comment}",
			wantMatch:  true,
			wantSubstr: `ephemeral activity collection "issues"`,
		},
		// Criterion 3.1: High-cardinality data-plane records / documents / SKUs / telemetry
		{
			pattern:    "projects/{project}/locations/{location}/instances/{instance}/dataTables/{data_table}/dataTableRows/{data_table_row}",
			wantMatch:  true,
			wantSubstr: `data-plane record leaf collection "dataTableRows"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document}",
			wantMatch:  true,
			wantSubstr: `data-plane document leaf collection "documents"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/catalogs/{catalog}/branches/{branch}/products/{product}",
			wantMatch:  true,
			wantSubstr: `data-plane catalog item leaf collection "products"`,
		},
		{
			pattern:    "projects/{project}/occurrences/{occurrence}",
			wantMatch:  true,
			wantSubstr: `data-plane telemetry leaf collection "occurrences"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/metadataStores/{metadata_store}/artifacts/{artifact}",
			wantMatch:  true,
			wantSubstr: `data-plane lineage leaf collection "artifacts"`,
		},
		// Criterion 3.2: Version snapshots, tags, experiments, test cases & examples
		{
			pattern:    "projects/{project}/locations/{location}/datasets/{dataset}/datasetVersions/{dataset_version}",
			wantMatch:  true,
			wantSubstr: `version snapshot leaf collection "datasetVersions"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/agents/{agent}/versions/{version}",
			wantMatch:  true,
			wantSubstr: `version snapshot leaf collection "versions"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/repositories/{repository}/packages/{package}/tags/{tag}",
			wantMatch:  true,
			wantSubstr: `artifact tag leaf collection "tags"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/agents/{agent}/environments/{environment}/experiments/{experiment}",
			wantMatch:  true,
			wantSubstr: `evaluation/experiment leaf collection "experiments"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/agents/{agent}/testCases/{test_case}",
			wantMatch:  true,
			wantSubstr: `evaluation/test-case leaf collection "testCases"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/agents/{agent}/playbooks/{playbook}/examples/{example}",
			wantMatch:  true,
			wantSubstr: `training/evaluation example leaf collection "examples"`,
		},
		// Criterion 3.3: Point-in-time backups & snapshots
		{
			pattern:    "projects/{project}/locations/{location}/backupVaults/{backup_vault}/backups/{backup}",
			wantMatch:  true,
			wantSubstr: `point-in-time backup leaf collection "backups"`,
		},
		{
			pattern:    "projects/{project}/locations/{location}/volumes/{volume}/snapshots/{snapshot}",
			wantMatch:  true,
			wantSubstr: `point-in-time snapshot leaf collection "snapshots"`,
		},
	} {
		reason, ok := nonDeclarativePatternReason(tc.pattern)
		if ok != tc.wantMatch {
			t.Errorf("nonDeclarativePatternReason(%q) ok = %v, want %v (reason: %q)", tc.pattern, ok, tc.wantMatch, reason)
			continue
		}
		if tc.wantMatch && !strings.Contains(reason, tc.wantSubstr) {
			t.Errorf("nonDeclarativePatternReason(%q) = %q, want substring %q", tc.pattern, reason, tc.wantSubstr)
		}
	}
}

func TestNonDeclarativeRPCReason(t *testing.T) {
	for _, tc := range []struct {
		name       string
		methods    []string
		wantMatch  bool
		wantSubstr string
	}{
		{
			name:      "declarative CRUD resource",
			methods:   []string{"CreateWidget", "GetWidget", "UpdateWidget", "DeleteWidget"},
			wantMatch: false,
		},
		{
			name:      "declarative immutable resource with Create + Get + Delete",
			methods:   []string{"CreateWidget", "GetWidget", "DeleteWidget"},
			wantMatch: false,
		},
		{
			name:      "declarative destroy-protected resource with Create + Get + Update",
			methods:   []string{"CreateWidget", "GetWidget", "UpdateWidget"},
			wantMatch: false,
		},
		{
			name:       "read-only discovered resource (Get only)",
			methods:    []string{"GetWidget", "ListWidgets"},
			wantMatch:  true,
			wantSubstr: "read-only/discovered resource",
		},
		{
			name:       "one-shot execution (Create + Get only, no Update or Delete)",
			methods:    []string{"CreateWidget", "GetWidget"},
			wantMatch:  true,
			wantSubstr: "one-shot execution",
		},
		{
			name:       "one-shot execution with Cancel RPC and no Update",
			methods:    []string{"CreateWidget", "GetWidget", "DeleteWidget", "CancelWidget"},
			wantMatch:  true,
			wantSubstr: "one-shot execution",
		},
		{
			name:       "high-cardinality data-plane collection with Purge RPC",
			methods:    []string{"CreateWidget", "GetWidget", "UpdateWidget", "DeleteWidget", "PurgeWidgets"},
			wantMatch:  true,
			wantSubstr: "bulk/purge data-plane RPCs [PurgeWidgets]",
		},
		{
			name:       "high-cardinality data-plane collection with BulkCreate RPC",
			methods:    []string{"CreateWidget", "GetWidget", "UpdateWidget", "DeleteWidget", "BulkCreateWidgets"},
			wantMatch:  true,
			wantSubstr: "bulk/purge data-plane RPCs [BulkCreateWidgets]",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pbPath := writeWidgetServiceProtoWithMethods(t, t.TempDir(), "projects/{project}/locations/{location}/widgets/{widget}", tc.methods)
			api, err := protoapi.LoadProto(pbPath, "")
			if err != nil {
				t.Fatalf("LoadProto: %v", err)
			}
			d, err := api.Files().FindDescriptorByName(protoreflect.FullName("google.cloud.example.v1.Widget"))
			if err != nil {
				t.Fatalf("FindDescriptorByName: %v", err)
			}
			msg := d.(protoreflect.MessageDescriptor)
			reason, ok := nonDeclarativeResourceReason(api, msg)
			if ok != tc.wantMatch {
				t.Fatalf("nonDeclarativeResourceReason() ok = %v, want %v (reason: %q)", ok, tc.wantMatch, reason)
			}
			if tc.wantMatch && !strings.Contains(reason, tc.wantSubstr) {
				t.Errorf("nonDeclarativeResourceReason() = %q, want substring %q", reason, tc.wantSubstr)
			}
		})
	}
}

func TestRunGenerateCRDRejectsNewNonDeclarativeResourceWithoutFlag(t *testing.T) {
	apiDir := t.TempDir()
	pbPath := writeSessionWidgetProto(t, t.TempDir())

	baseOpts := func() *GenerateCRDOptions {
		return &GenerateCRDOptions{
			GenerateOptions: &options.GenerateOptions{
				ProtoSourcePath: pbPath,
				APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
			},
			ServiceName:         "google.cloud.example.v1",
			OutputAPIDirectory:  apiDir,
			Resources:           options.ResourceList{{Kind: "ExampleSession", ProtoName: "Session"}},
			PrepopulateResource: true,
		}
	}

	// 1. Scaffolding a new non-declarative resource fails without --allow-non-declarative.
	err := RunGenerateCRD(context.Background(), baseOpts())
	if err == nil || !strings.Contains(err.Error(), "--allow-non-declarative") {
		t.Fatalf("expected --allow-non-declarative error, got: %v", err)
	}
	typeFile := filepath.Join(apiDir, "example/v1alpha1/examplesession_types.go")
	if _, statErr := os.Stat(typeFile); !os.IsNotExist(statErr) {
		t.Fatalf("expected %s not to be created on failure", typeFile)
	}

	// 2. Passing --allow-non-declarative succeeds and creates <kind>_types.go.
	allowed := baseOpts()
	allowed.AllowNonDeclarative = true
	if err := RunGenerateCRD(context.Background(), allowed); err != nil {
		t.Fatalf("RunGenerateCRD with AllowNonDeclarative: %v", err)
	}
	if _, statErr := os.Stat(typeFile); statErr != nil {
		t.Fatalf("expected %s to exist after AllowNonDeclarative: %v", typeFile, statErr)
	}

	// 3. Re-running when <kind>_types.go already exists succeeds even without --allow-non-declarative.
	if err := RunGenerateCRD(context.Background(), baseOpts()); err != nil {
		t.Fatalf("RunGenerateCRD on existing resource without AllowNonDeclarative: %v", err)
	}
}

func TestRunGenerateCRDReportsAllRejectedResources(t *testing.T) {
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	sessionOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(sessionOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/Session",
		Pattern: []string{"projects/{project}/locations/{location}/sessions/{session}"},
	})
	productOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(productOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/Product",
		Pattern: []string{"projects/{project}/locations/{location}/catalogs/{catalog}/products/{product}"},
	})
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{
		{
			Name:    proto.String("google/cloud/example/v1/resources.proto"),
			Package: proto.String("google.cloud.example.v1"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{
				{
					Name:    proto.String("Session"),
					Options: sessionOpts,
					Field:   []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: str}},
				},
				{
					Name:    proto.String("Product"),
					Options: productOpts,
					Field:   []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: str}},
				},
			},
		},
		{
			Name:    proto.String("google/cloud/example/v1beta/beta.proto"),
			Package: proto.String("google.cloud.example.v1beta"),
			Syntax:  proto.String("proto3"),
			MessageType: []*descriptorpb.DescriptorProto{
				{
					Name:  proto.String("AlphaWidget"),
					Field: []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: str}},
				},
				{
					Name:  proto.String("BetaGadget"),
					Field: []*descriptorpb.FieldDescriptorProto{{Name: proto.String("name"), Number: proto.Int32(1), Type: str}},
				},
			},
		},
	}}
	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	pbPath := filepath.Join(t.TempDir(), "multi.pb")
	writeFile(t, pbPath, string(b))

	// 1. Multiple non-GA resources are all listed in the error report.
	nonGAErr := RunGenerateCRD(context.Background(), &GenerateCRDOptions{
		GenerateOptions: &options.GenerateOptions{
			ProtoSourcePath: pbPath,
			APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
		},
		ServiceName:        "google.cloud.example.v1beta",
		OutputAPIDirectory: t.TempDir(),
		Resources: options.ResourceList{
			{Kind: "ExampleAlphaWidget", ProtoName: "AlphaWidget"},
			{Kind: "ExampleBetaGadget", ProtoName: "BetaGadget"},
		},
		PrepopulateResource: true,
	})
	if nonGAErr == nil {
		t.Fatalf("expected non-GA rejection error, got nil")
	}
	for _, want := range []string{"rejected 2 non-GA resource(s)", "ExampleAlphaWidget", "ExampleBetaGadget", "--allow-non-ga-proto"} {
		if !strings.Contains(nonGAErr.Error(), want) {
			t.Errorf("expected non-GA error %q to contain %q", nonGAErr.Error(), want)
		}
	}

	// 2. Multiple non-declarative resources are all listed in the error report.
	nonDeclErr := RunGenerateCRD(context.Background(), &GenerateCRDOptions{
		GenerateOptions: &options.GenerateOptions{
			ProtoSourcePath: pbPath,
			APIVersion:      "example.cnrm.cloud.google.com/v1alpha1",
		},
		ServiceName:        "google.cloud.example.v1",
		OutputAPIDirectory: t.TempDir(),
		Resources: options.ResourceList{
			{Kind: "ExampleSession", ProtoName: "Session"},
			{Kind: "ExampleProduct", ProtoName: "Product"},
		},
		PrepopulateResource: true,
	})
	if nonDeclErr == nil {
		t.Fatalf("expected non-declarative rejection error, got nil")
	}
	for _, want := range []string{"rejected 2 suspected non-declarative resource(s)", "ExampleSession", "ExampleProduct", "--allow-non-declarative"} {
		if !strings.Contains(nonDeclErr.Error(), want) {
			t.Errorf("expected non-declarative error %q to contain %q", nonDeclErr.Error(), want)
		}
	}
}

func writeSessionWidgetProto(t *testing.T, dir string) string {
	t.Helper()
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	msgOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(msgOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/Session",
		Pattern: []string{"projects/{project}/locations/{location}/sessions/{session}"},
	})
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{{
		Name:    proto.String("google/cloud/example/v1/session.proto"),
		Package: proto.String("google.cloud.example.v1"),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name:    proto.String("Session"),
			Options: msgOpts,
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
	path := filepath.Join(dir, "session.pb")
	writeFile(t, path, string(b))
	return path
}

func writeWidgetServiceProtoWithMethods(t *testing.T, dir, pattern string, methodNames []string) string {
	t.Helper()
	str := descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum()
	msgOpts := &descriptorpb.MessageOptions{}
	proto.SetExtension(msgOpts, annotations.E_Resource, &annotations.ResourceDescriptor{
		Type:    "example.googleapis.com/Widget",
		Pattern: []string{pattern},
	})
	var methods []*descriptorpb.MethodDescriptorProto
	for _, m := range methodNames {
		methods = append(methods, &descriptorpb.MethodDescriptorProto{
			Name:       proto.String(m),
			InputType:  proto.String(".google.cloud.example.v1.Widget"),
			OutputType: proto.String(".google.cloud.example.v1.Widget"),
		})
	}
	fds := &descriptorpb.FileDescriptorSet{File: []*descriptorpb.FileDescriptorProto{{
		Name:    proto.String("google/cloud/example/v1/widget.proto"),
		Package: proto.String("google.cloud.example.v1"),
		Syntax:  proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name:    proto.String("Widget"),
			Options: msgOpts,
			Field: []*descriptorpb.FieldDescriptorProto{
				{Name: proto.String("name"), Number: proto.Int32(1), Type: str},
			},
		}},
		Service: []*descriptorpb.ServiceDescriptorProto{{
			Name:   proto.String("WidgetService"),
			Method: methods,
		}},
	}}}
	b, err := proto.Marshal(fds)
	if err != nil {
		t.Fatalf("marshalling descriptor set: %v", err)
	}
	path := filepath.Join(dir, "widget_service.pb")
	writeFile(t, path, string(b))
	return path
}
