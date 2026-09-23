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

package protoapi

import (
	"testing"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestSplitPattern(t *testing.T) {
	grid := []struct {
		name           string
		pattern        string
		wantCollection string
		wantParent     string
	}{
		{
			name:           "project and location",
			pattern:        "projects/{project}/locations/{location}/foos/{foo}",
			wantCollection: "foos",
			wantParent:     "projects/locations",
		},
		{
			// The casing here is the whole point: the template's ToLower would
			// produce "lbtrafficextensions", which is not a valid resource name.
			name:           "camelCase collection is preserved",
			pattern:        "projects/{project}/locations/{location}/lbTrafficExtensions/{lb_traffic_extension}",
			wantCollection: "lbTrafficExtensions",
			wantParent:     "projects/locations",
		},
		{
			name:           "irregular english plural",
			pattern:        "projects/{project}/policies/{policy}",
			wantCollection: "policies",
			wantParent:     "projects",
		},
		{
			name:           "organization parent",
			pattern:        "organizations/{organization}/bars/{bar}",
			wantCollection: "bars",
			wantParent:     "organizations",
		},
		{
			name:           "deeply nested parent",
			pattern:        "projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{node_pool}",
			wantCollection: "nodePools",
			wantParent:     "projects/locations/clusters",
		},
		{
			// Some patterns name a collection with no trailing id.
			name:           "pattern ending in a literal has no collection",
			pattern:        "projects/{project}/locations",
			wantCollection: "",
			wantParent:     "projects/locations",
		},
		{
			name:           "singleton at root",
			pattern:        "foos/{foo}",
			wantCollection: "foos",
			wantParent:     "",
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			collection, parent := splitPattern(g.pattern)
			if collection != g.wantCollection {
				t.Errorf("collection = %q, want %q", collection, g.wantCollection)
			}
			if parent != g.wantParent {
				t.Errorf("parent = %q, want %q", parent, g.wantParent)
			}
		})
	}
}

func TestClassifyParent(t *testing.T) {
	grid := []struct {
		parentPath string
		want       ParentStyle
	}{
		{"projects/locations", ParentProjectLocation},
		{"projects", ParentProject},
		{"organizations", ParentOrganization},
		{"folders", ParentFolder},
		{"projects/locations/clusters", ParentOther},
		{"properties", ParentOther},
		{"", ParentUnknown},
	}

	for _, g := range grid {
		t.Run(g.parentPath, func(t *testing.T) {
			if got := classifyParent(g.parentPath); got != g.want {
				t.Errorf("classifyParent(%q) = %q, want %q", g.parentPath, got, g.want)
			}
		})
	}
}

func TestGetResourceMetadata(t *testing.T) {
	rdSingle := &annotations.ResourceDescriptor{
		Type:    "networkservices.googleapis.com/LbTrafficExtension",
		Pattern: []string{"projects/{project}/locations/{location}/lbTrafficExtensions/{lb_traffic_extension}"},
		Plural:  "lbTrafficExtensions",
	}
	rdMulti := &annotations.ResourceDescriptor{
		Type: "logging.googleapis.com/LogSink",
		Pattern: []string{
			"projects/{project}/sinks/{sink}",
			"organizations/{organization}/sinks/{sink}",
			"folders/{folder}/sinks/{sink}",
		},
		Plural: "sinks",
	}

	optsSingle := &descriptorpb.MessageOptions{}
	proto.SetExtension(optsSingle, annotations.E_Resource, rdSingle)

	optsMulti := &descriptorpb.MessageOptions{}
	proto.SetExtension(optsMulti, annotations.E_Resource, rdMulti)

	fdp := &descriptorpb.FileDescriptorProto{
		Name:    proto.String("test.proto"),
		Package: proto.String("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name:    proto.String("LbTrafficExtension"),
				Options: optsSingle,
			},
			{
				Name:    proto.String("LogSink"),
				Options: optsMulti,
			},
		},
	}

	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("protodesc.NewFile: %v", err)
	}

	singleMD := GetResourceMetadata(fd.Messages().ByName("LbTrafficExtension"))
	if singleMD == nil {
		t.Fatal("expected non-nil ResourceMetadata for LbTrafficExtension")
	}
	if singleMD.ParentStyle != ParentProjectLocation {
		t.Errorf("ParentStyle = %q, want %q", singleMD.ParentStyle, ParentProjectLocation)
	}
	if singleMD.Collection != "lbTrafficExtensions" {
		t.Errorf("Collection = %q, want %q", singleMD.Collection, "lbTrafficExtensions")
	}
	if len(singleMD.Patterns) != 1 {
		t.Errorf("len(Patterns) = %d, want 1", len(singleMD.Patterns))
	}

	multiMD := GetResourceMetadata(fd.Messages().ByName("LogSink"))
	if multiMD == nil {
		t.Fatal("expected non-nil ResourceMetadata for LogSink")
	}
	if multiMD.ParentStyle != ParentMulti {
		t.Errorf("ParentStyle = %q, want %q", multiMD.ParentStyle, ParentMulti)
	}
	if multiMD.Collection != "sinks" {
		t.Errorf("Collection = %q, want %q", multiMD.Collection, "sinks")
	}
	if len(multiMD.Patterns) != 3 {
		t.Errorf("len(Patterns) = %d, want 3", len(multiMD.Patterns))
	}
}

// TestParentPair pins which segment ParentPair calls the parent. The scaffolder
// names a reference field from the collection it returns, so a pattern read one
// pair off puts the wrong resource in somebody's CRD.
func TestParentPair(t *testing.T) {
	grid := []struct {
		name            string
		pattern         string
		wantCollection  string
		wantPlaceholder string
	}{
		{
			name:            "ends on its own collection and id",
			pattern:         "projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{node_pool}",
			wantCollection:  "clusters",
			wantPlaceholder: "cluster",
		},
		{
			// 274 of the 3160 patterns in googleapis end in a literal naming a
			// singleton, so the last collection/{id} pair is already the parent.
			name:            "ends in a singleton literal",
			pattern:         "accounts/{account}/programs/{program}/checkoutSettings",
			wantCollection:  "programs",
			wantPlaceholder: "program",
		},
		{
			name:            "project and location parent",
			pattern:         "projects/{project}/locations/{location}/foos/{foo}",
			wantCollection:  "locations",
			wantPlaceholder: "location",
		},
		{
			// The only patterns in googleapis with two placeholders in a row are
			// 2 of the 3160, both healthcare FHIR. Nothing between them names a
			// collection, so the pattern has no parent to read.
			name:            "two placeholders in a row",
			pattern:         "projects/{project}/locations/{location}/datasets/{dataset}/fhirStores/{fhir_store}/fhir/{resource_type}/{fhir_resource_id}",
			wantCollection:  "",
			wantPlaceholder: "",
		},
		{
			name:            "no parent at all",
			pattern:         "foos/{foo}",
			wantCollection:  "",
			wantPlaceholder: "",
		},
		{
			name:            "pattern ending in a literal names no parent",
			pattern:         "projects/{project}/locations",
			wantCollection:  "",
			wantPlaceholder: "",
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			// Act
			collection, placeholder := ParentPair(g.pattern)

			// Assert
			if collection != g.wantCollection {
				t.Errorf("ParentPair(%q) collection = %q, want %q", g.pattern, collection, g.wantCollection)
			}
			if placeholder != g.wantPlaceholder {
				t.Errorf("ParentPair(%q) placeholder = %q, want %q", g.pattern, placeholder, g.wantPlaceholder)
			}
		})
	}
}
