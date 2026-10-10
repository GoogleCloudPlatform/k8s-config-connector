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
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// goldenPath is the checked-in report for every kind mapped in
// apis/*/generate.sh.
const goldenPath = "testdata/archetypes.tsv"

// TestArchetypesGolden regenerates the report against the descriptor set for
// the googleapis version pinned in apis/git.versions, and diffs it against
// goldenPath. Set WRITE_GOLDEN_OUTPUT=1 to rewrite goldenPath instead.
//
// The test is skipped when that descriptor set has not been built, because
// .build/googleapis.pb holds whichever version a generate.sh built last.
func TestArchetypesGolden(t *testing.T) {
	// Arrange
	root, err := options.RepoRoot()
	if err != nil {
		t.Fatalf("finding the repository root: %v", err)
	}
	sha, err := pinnedGoogleapisVersion(filepath.Join(root, "apis", "git.versions"))
	if err != nil {
		t.Fatal(err)
	}
	descriptors := filepath.Join(root, ".build", "googleapis-"+sha+".pb")
	if _, err := os.Stat(descriptors); errors.Is(err, fs.ErrNotExist) {
		t.Skipf("skipping: %s does not exist. It is the descriptor set for the googleapis version pinned in apis/git.versions; build it with dev/tools/controllerbuilder/generate-proto.sh", descriptors)
	} else if err != nil {
		t.Fatal(err)
	}
	api, err := protoapi.LoadProto(descriptors, "")
	if err != nil {
		t.Fatalf("loading %s: %v", descriptors, err)
	}

	// Act
	rows, err := Report(api, filepath.Join(root, "apis"))
	if err != nil {
		t.Fatalf("Report: %v", err)
	}
	var got bytes.Buffer
	if err := WriteTSV(&got, rows); err != nil {
		t.Fatalf("WriteTSV: %v", err)
	}

	// Assert
	if os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(goldenPath, got.Bytes(), 0o644); err != nil {
			t.Fatalf("writing %s: %v", goldenPath, err)
		}
		t.Logf("wrote %s", goldenPath)
		return
	}
	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("reading %s: %v; run with WRITE_GOLDEN_OUTPUT=1 to create it", goldenPath, err)
	}
	if diff := cmp.Diff(strings.Split(string(want), "\n"), strings.Split(got.String(), "\n")); diff != "" {
		t.Errorf("report does not match %s (-want +got); if the change is expected, rerun with WRITE_GOLDEN_OUTPUT=1:\n%s", goldenPath, diff)
	}
}

// pinnedGoogleapisVersion returns the googleapis commit that apis/git.versions
// pins, read the way generate-proto.sh reads it.
func pinnedGoogleapisVersion(gitVersions string) (string, error) {
	b, err := os.ReadFile(gitVersions)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(b), "\n") {
		if fields := strings.Fields(line); len(fields) == 2 && fields[0] == "https://github.com/googleapis/googleapis" {
			return fields[1], nil
		}
	}
	return "", fmt.Errorf("%s pins no googleapis version", gitVersions)
}

func TestWriteTSV(t *testing.T) {
	// Arrange
	file, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    proto.String("google/cloud/foo/v1/foo.proto"),
		Package: proto.String("google.cloud.foo.v1"),
		Syntax:  proto.String("proto3"),
		Service: []*descriptorpb.ServiceDescriptorProto{{Name: proto.String("FooService")}},
	}, nil)
	if err != nil {
		t.Fatalf("protodesc.NewFile: %v", err)
	}
	svc := file.Services().Get(0)
	msg := (&emptypb.Empty{}).ProtoReflect().Descriptor()
	lro := &protoapi.LROInfo{}
	rows := []Row{
		{
			Kind:       "FooThing",
			APIVersion: "foo.cnrm.cloud.google.com/v1beta1",
			Message:    "google.cloud.foo.v1.Thing",
			Archetype:  protoapi.ArchetypeStandard,
			API: &protoapi.ResourceAPI{
				Message:     msg,
				Service:     svc,
				Style:       protoapi.StyleAIP,
				DefaultHost: "foo.googleapis.com",
				Metadata: &protoapi.ResourceMetadata{
					Pattern:  "projects/{project}/locations/{location}/things/{thing}",
					Patterns: []string{"projects/{project}/locations/{location}/things/{thing}", "folders/{folder}/things/{thing}"},
				},
				Get:    &protoapi.StandardMethod{HTTPVerb: "GET"},
				Create: &protoapi.StandardMethod{HTTPVerb: "POST", IDField: "thing_id", LRO: lro},
				Update: &protoapi.StandardMethod{HTTPVerb: "PATCH", UpdateMaskField: "update_mask", LRO: lro},
				Delete: &protoapi.StandardMethod{HTTPVerb: "DELETE", EtagField: "etag", LRO: lro},
				List:   &protoapi.StandardMethod{},
			},
		},
		{
			Kind:       "FooOdd",
			APIVersion: "foo.cnrm.cloud.google.com/v1alpha1",
			Message:    "google.cloud.foo.v1.Odd",
			Archetype:  protoapi.ArchetypeNonstandard,
			API: &protoapi.ResourceAPI{
				Message: msg,
				Service: svc,
				Style:   protoapi.StyleAIP,
				Get:     &protoapi.StandardMethod{HTTPVerb: "GET", NonstandardRequest: true},
				Update:  &protoapi.StandardMethod{HTTPVerb: "PUT", NonstandardRequest: true, UpdateMaskField: "update_mask"},
			},
		},
		{
			Kind:        "FooMissing",
			APIVersion:  "foo.cnrm.cloud.google.com/v1alpha1",
			Message:     "google.cloud.foo.v1.Missing",
			ProtoSource: "HEAD",
			Archetype:   protoapi.ArchetypeUnresolved,
		},
	}
	var got bytes.Buffer

	// Act
	err = WriteTSV(&got, rows)

	// Assert
	if err != nil {
		t.Fatalf("WriteTSV: %v", err)
	}
	want := strings.Join([]string{
		"kind\tapiVersion\tmessage\tprotoSource\tarchetype\tstyle\tdefaultHost\tservice\tpattern\tpatternCount\tget\tcreate\tupdate\tdelete\tlist",
		"FooThing\tfoo.cnrm.cloud.google.com/v1beta1\tgoogle.cloud.foo.v1.Thing\t-\tA-standard\taip\tfoo.googleapis.com\tFooService\tprojects/{project}/locations/{location}/things/{thing}\t2\tGET\tPOST id=thing_id lro\tPATCH mask lro\tDELETE etag lro\tRPC",
		"FooOdd\tfoo.cnrm.cloud.google.com/v1alpha1\tgoogle.cloud.foo.v1.Odd\t-\tE-nonstandard\taip\t-\tFooService\t-\t0\tGET nonstandard\t-\tPUT nonstandard mask\t-\t-",
		"FooMissing\tfoo.cnrm.cloud.google.com/v1alpha1\tgoogle.cloud.foo.v1.Missing\tHEAD\tI-unresolved\t-\t-\t-\t-\t0\t-\t-\t-\t-\t-",
	}, "\n") + "\n"
	if diff := cmp.Diff(want, got.String()); diff != "" {
		t.Errorf("WriteTSV mismatch (-want +got):\n%s", diff)
	}
}

func TestLatestMappings(t *testing.T) {
	// Arrange
	mappings := []Mapping{
		{Kind: "Zeta", APIVersion: "foo.cnrm.cloud.google.com/v1alpha1", Source: "zeta alpha"},
		{Kind: "Alpha", APIVersion: "foo.cnrm.cloud.google.com/v1beta1", Source: "alpha beta"},
		{Kind: "Alpha", APIVersion: "foo.cnrm.cloud.google.com/v1", Source: "alpha ga"},
		{Kind: "Alpha", APIVersion: "foo.cnrm.cloud.google.com/v1alpha1", Source: "alpha alpha"},
		{Kind: "Beta", APIVersion: "foo.cnrm.cloud.google.com/v1beta1", Source: "beta first"},
		{Kind: "Beta", APIVersion: "foo.cnrm.cloud.google.com/v1beta1", Source: "beta second"},
		{Kind: "Beta", APIVersion: "foo.cnrm.cloud.google.com/v1alpha1", Source: "beta alpha"},
	}

	// Act
	got := latestMappings(mappings)

	// Assert
	var sources []string
	for _, m := range got {
		sources = append(sources, m.Source)
	}
	if diff := cmp.Diff([]string{"alpha ga", "beta second", "zeta alpha"}, sources); diff != "" {
		t.Errorf("latestMappings mismatch (-want +got):\n%s", diff)
	}
}

func TestMessageCandidates(t *testing.T) {
	for _, tc := range []struct {
		name    string
		mapping Mapping
		want    []string
	}{
		{
			name:    "a message relative to the service",
			mapping: Mapping{Services: []string{"google.cloud.foo.v1"}, ProtoName: "Thing"},
			want:    []string{"google.cloud.foo.v1.Thing"},
		},
		{
			name:    "each service in turn",
			mapping: Mapping{Services: []string{"google.cloud.foo.v1beta", "google.cloud.foo.v1"}, ProtoName: "Thing"},
			want:    []string{"google.cloud.foo.v1beta.Thing", "google.cloud.foo.v1.Thing"},
		},
		{
			name:    "a nested message is relative to the service first",
			mapping: Mapping{Services: []string{"google.cloud.foo.v1"}, ProtoName: "Thing.Part"},
			want:    []string{"google.cloud.foo.v1.Thing.Part", "Thing.Part"},
		},
		{
			name:    "a fully-qualified message is tried as written first",
			mapping: Mapping{Services: []string{"google.monitoring.v3"}, ProtoName: "google.api.MetricDescriptor"},
			want:    []string{"google.api.MetricDescriptor", "google.monitoring.v3.google.api.MetricDescriptor"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := messageCandidates(tc.mapping)

			// Assert
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("messageCandidates mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
