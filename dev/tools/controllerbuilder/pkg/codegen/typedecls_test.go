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
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// A type gets the doc comment of its declaration only when the declaration
// has no other types. Files that include rejects are not read.
func TestParseTypeDecls(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	for name, src := range map[string]string{
		"a.go": "package test\n\n" +
			"// Alone is declared alone.\n" +
			"type Alone struct{}\n\n" +
			"// The group's comment.\n" +
			"type (\n" +
			"\t// First has its own comment.\n" +
			"\tFirst struct{}\n" +
			"\tSecond int\n" +
			")\n",
		"skipped.go": "package test\n\ntype Skipped struct{}\n",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	includeA := func(e fs.DirEntry) bool { return e.Name() == "a.go" }
	type decl struct{ File, Name, Doc string }
	want := []decl{
		{File: "a.go", Name: "Alone", Doc: "Alone is declared alone.\n"},
		{File: "a.go", Name: "First", Doc: "First has its own comment.\n"},
		{File: "a.go", Name: "Second", Doc: ""},
	}

	// Act
	decls, err := ParseTypeDecls(dir, includeA)

	// Assert
	if err != nil {
		t.Fatalf("ParseTypeDecls: %v", err)
	}
	var got []decl
	for _, d := range decls {
		got = append(got, decl{File: d.File, Name: d.Name(), Doc: d.Doc.Text()})
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ParseTypeDecls() mismatch (-want +got):\n%s", diff)
	}
}

func TestParseTypeDeclsWithoutDirectory(t *testing.T) {
	// Act
	_, err := ParseTypeDecls(filepath.Join(t.TempDir(), "missing"), func(fs.DirEntry) bool { return true })

	// Assert
	if !os.IsNotExist(err) {
		t.Errorf("ParseTypeDecls() error = %v, want one that os.IsNotExist accepts", err)
	}
}

func TestParseTypeDeclsReportsTheFile(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "broken.go"), []byte("package test\n\ntype {\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Act
	_, err := ParseTypeDecls(dir, func(fs.DirEntry) bool { return true })

	// Assert
	if err == nil || !strings.HasPrefix(err.Error(), "parsing "+filepath.Join(dir, "broken.go")+": ") {
		t.Errorf("ParseTypeDecls() error = %v, want it to start with the file path", err)
	}
}
