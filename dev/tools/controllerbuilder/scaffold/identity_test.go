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
	"go/parser"
	"go/token"
	"strings"
	"testing"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/template/apis"
)

func renderIdentity(t *testing.T, args *apis.APIArgs) string {
	t.Helper()
	tmpl, err := template.New(args.Kind).Funcs(funcMap).Parse(apis.IdentityTemplate)
	if err != nil {
		t.Fatalf("parsing identity template: %v", err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, args); err != nil {
		t.Fatalf("executing identity template: %v", err)
	}
	return buf.String()
}

func TestIdentityTemplateRendersValidGo(t *testing.T) {
	grid := []struct {
		name string
		args *apis.APIArgs
		// substrings that must appear, and must not
		want    []string
		notWant []string
	}{
		{
			name: "project and location parent",
			args: &apis.APIArgs{
				Version:              "v1alpha1",
				Kind:                 "NetworkServicesLBTrafficExtension",
				ProtoMessageName:     "LbTrafficExtension",
				ProtoMessageFullName: "google.cloud.networkservices.v1.LbTrafficExtension",
				Collection:           "lbTrafficExtensions",
				ParentStyle:          "project_location",
				ResourcePattern:      "projects/{project}/locations/{location}/lbTrafficExtensions/{lb_traffic_extension}",
			},
			// The casing is the whole point: ToLower would emit lbtrafficextensions.
			want: []string{
				`"/lbTrafficExtensions/"`,
				`tokens[4] != "lbTrafficExtensions"`,
				"Location  string",
				"len(tokens) != 6",
			},
			notWant: []string{"lbtrafficextensions"},
		},
		{
			name: "project-only parent drops location",
			args: &apis.APIArgs{
				Version:              "v1alpha1",
				Kind:                 "ExamplePolicy",
				ProtoMessageName:     "Policy",
				ProtoMessageFullName: "google.cloud.example.v1.Policy",
				Collection:           "policies",
				ParentStyle:          "project",
				ResourcePattern:      "projects/{project}/policies/{policy}",
			},
			// Irregular plural: ToLower+"s" would emit policys.
			want: []string{
				`"/policies/"`,
				"len(tokens) != 4",
				`tokens[2] != "policies"`,
			},
			notWant: []string{"policys", "Location  string", "obj.Spec.Location"},
		},
		{
			name: "unmodelled parent still renders, with a warning",
			args: &apis.APIArgs{
				Version:              "v1alpha1",
				Kind:                 "ExampleNodePool",
				ProtoMessageName:     "NodePool",
				ProtoMessageFullName: "google.cloud.example.v1.NodePool",
				Collection:           "nodePools",
				ParentStyle:          "other",
				ResourcePattern:      "projects/{p}/locations/{l}/clusters/{c}/nodePools/{n}",
			},
			want: []string{"TODO", "nodePools", "clusters/{c}"},
		},
		{
			// An organization-parented resource compiles but is wrong: the Parent
			// struct below still carries ProjectID and Location. The TODO is what
			// tells the reader that, so it has to be there.
			name: "organization parent renders, with a warning",
			args: &apis.APIArgs{
				Version:              "v1alpha1",
				Kind:                 "ExamplePolicy",
				ProtoMessageName:     "Policy",
				ProtoMessageFullName: "google.cloud.example.v1.Policy",
				Collection:           "policies",
				ParentStyle:          "organization",
				ResourcePattern:      "organizations/{o}/policies/{p}",
			},
			want: []string{"TODO", "policies", "organizations/{o}"},
		},
		{
			name: "folder parent renders, with a warning",
			args: &apis.APIArgs{
				Version:              "v1alpha1",
				Kind:                 "ExampleFeed",
				ProtoMessageName:     "Feed",
				ProtoMessageFullName: "google.cloud.example.v1.Feed",
				Collection:           "feeds",
				ParentStyle:          "folder",
				ResourcePattern:      "folders/{f}/feeds/{feed}",
			},
			want: []string{"TODO", "feeds", "folders/{f}"},
		},
		{
			name: "multi-parent resource renders all patterns and warning",
			args: &apis.APIArgs{
				Version:              "v1alpha1",
				Kind:                 "LoggingLogSink",
				ProtoMessageName:     "LogSink",
				ProtoMessageFullName: "google.logging.v2.LogSink",
				Collection:           "sinks",
				ParentStyle:          "multi",
				ResourcePattern:      "projects/{project}/sinks/{sink}",
				ResourcePatterns: []string{
					"projects/{project}/sinks/{sink}",
					"organizations/{organization}/sinks/{sink}",
					"folders/{folder}/sinks/{sink}",
				},
			},
			want: []string{
				"TODO: this resource supports multiple parent hierarchies",
				"projects/{project}/sinks/{sink}",
				"organizations/{organization}/sinks/{sink}",
				"folders/{folder}/sinks/{sink}",
				`"/sinks/"`,
			},
		},
		{
			name: "no annotation falls back to the old guess",
			args: &apis.APIArgs{
				Version:              "v1alpha1",
				Kind:                 "ExampleThing",
				ProtoMessageName:     "Thing",
				ProtoMessageFullName: "google.cloud.example.v1.Thing",
				Collection:           "things", // what buildAPIArgs falls back to
				ParentStyle:          "unknown",
			},
			want: []string{"declares no google.api.resource pattern", `"/things/"`},
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			got := renderIdentity(t, g.args)

			// The generated file has to compile; a template that emits a stray
			// brace or an unused variable is worse than one that guesses.
			if _, err := parser.ParseFile(token.NewFileSet(), "identity.go", got, parser.AllErrors); err != nil {
				t.Fatalf("rendered identity file is not valid Go: %v\n\n%s", err, got)
			}

			for _, w := range g.want {
				if !strings.Contains(got, w) {
					t.Errorf("missing %q in rendered output:\n%s", w, got)
				}
			}
			for _, w := range g.notWant {
				if strings.Contains(got, w) {
					t.Errorf("unexpected %q in rendered output:\n%s", w, got)
				}
			}
		})
	}
}
