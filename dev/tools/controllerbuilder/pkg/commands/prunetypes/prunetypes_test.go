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

package prunetypes

import (
	"go/parser"
	"go/token"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDropUnusedImports(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{
			// Verifies that imports referenced only by commented-out types are dropped.
			name: "an import only a commented-out type uses is dropped",
			in: `package p

import (
	"fmt"

	krm "example.com/apis/krm"
)

var _ = fmt.Sprint

// type Foo struct {
// 	Bar *krm.Bar
// }
`,
			want: []string{"fmt"},
		},
		{
			name: "an import a live type uses is kept",
			in: `package p

import krm "example.com/apis/krm"

type Foo struct {
	Bar *krm.Bar
}
`,
			want: []string{"example.com/apis/krm"},
		},
		{
			name: "an alias is matched rather than the last path segment",
			in: `package p

import (
	refs "example.com/apis/refs/v1beta1"
	unused "example.com/apis/common"
)

type Foo struct {
	Ref *refs.ProjectRef
}
`,
			want: []string{"example.com/apis/refs/v1beta1"},
		},
		{
			// Verifies that multiple unused imports are all safely dropped.
			name: "two unused imports are both dropped",
			in: `package p

import (
	"fmt"
	"strings"
)

type Foo struct{}
`,
			want: nil,
		},
		{
			name: "blank and dot imports are kept",
			in: `package p

import (
	_ "embed"
	. "example.com/dot"
)

type Foo struct{}
`,
			want: []string{"embed", "example.com/dot"},
		},
		{
			// Without an alias the qualifier is the package name, which the
			// file does not state. Where the last path segment cannot be that
			// name, the import has to stay: dropping a live import breaks the
			// build just as surely as keeping a dead one.
			name: "an import whose package name is not its last path segment is kept",
			in: `package p

import (
	"example.com/mod/v2"
	"gopkg.in/yaml.v3"
)

var _ = yaml.Marshal
var _ = mod.New
`,
			want: []string{"example.com/mod/v2", "gopkg.in/yaml.v3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dropUnusedImports("test.go", []byte(tt.in))
			if err != nil {
				t.Fatalf("dropUnusedImports: %v", err)
			}
			if diff := cmp.Diff(tt.want, importPaths(t, got)); diff != "" {
				t.Errorf("imports mismatch (-want +got):\n%s\noutput:\n%s", diff, got)
			}
			// The commented-out type is the pruner's record of what it
			// removed, so rewriting the file must not lose it.
			for _, line := range strings.Split(tt.in, "\n") {
				if strings.HasPrefix(line, "// type ") && !strings.Contains(string(got), line) {
					t.Errorf("output lost the comment %q:\n%s", line, got)
				}
			}
		})
	}
}

func TestDropUnusedImportsLeavesACleanFileAlone(t *testing.T) {
	in := "package p\n\nimport \"fmt\"\n\nvar _   =   fmt.Sprint\n"
	got, err := dropUnusedImports("test.go", []byte(in))
	if err != nil {
		t.Fatalf("dropUnusedImports: %v", err)
	}
	// Byte-identical, including the spacing gofmt would change: a file
	// with nothing to drop is not reformatted.
	if string(got) != in {
		t.Errorf("got %q, want the input unchanged", got)
	}
}

func TestDropUnusedImportsReportsAParseError(t *testing.T) {
	if _, err := dropUnusedImports("test.go", []byte("package p\n\nfunc {")); err == nil {
		t.Error("got no error for a file that does not parse")
	}
}

func importPaths(t *testing.T, src []byte) []string {
	t.Helper()
	f, err := parser.ParseFile(token.NewFileSet(), "out.go", src, parser.ImportsOnly)
	if err != nil {
		t.Fatalf("output does not parse: %v\n%s", err, src)
	}
	var paths []string
	for _, imp := range f.Imports {
		p, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			t.Fatalf("unquoting %s: %v", imp.Path.Value, err)
		}
		paths = append(paths, p)
	}
	return paths
}
