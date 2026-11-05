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

package lint

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/sourcelinks"
)

// sourceLinkMismatches pairs the source-link guess lines in apis/*/*/*_types.go
// with the open source-link entries in apis/*/judgement_queue.yaml, and
// returns one message for each side without a partner.
//
// generate-types --emit-source-links writes the two together. A guess line
// with no open entry means someone resolved the entry but left the guess. An
// open entry with no guess line means someone fixed the link but did not
// resolve the entry.
func sourceLinkMismatches(apisDir string) ([]string, error) {
	type key struct{ service, kind, reason string }

	queued := map[key]string{} // -> queue file
	queues, err := filepath.Glob(filepath.Join(apisDir, "*", judgement.FileName))
	if err != nil {
		return nil, err
	}
	for _, path := range queues {
		q, err := judgement.Read(path)
		if err != nil {
			return nil, err
		}
		service := filepath.Base(filepath.Dir(path))
		for _, e := range q.Entries {
			if e.IsOpen() && judgement.IsSourceLinkReason(e.Reason) {
				queued[key{service, e.Kind, e.Reason}] = path
			}
		}
	}

	guessed := map[key]string{} // -> types file
	typeFiles, err := filepath.Glob(filepath.Join(apisDir, "*", "*", "*_types.go"))
	if err != nil {
		return nil, err
	}
	for _, path := range typeFiles {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		if !bytes.Contains(data, []byte("+kcc:guess="+sourcelinks.GuessKind)) {
			continue
		}
		h := sourcelinks.Parse(data)
		service := filepath.Base(filepath.Dir(filepath.Dir(path)))
		for _, reason := range h.Guesses() {
			guessed[key{service, h.Kind, reason}] = path
		}
	}

	var out []string
	for k, path := range guessed {
		if _, ok := queued[k]; !ok {
			out = append(out, fmt.Sprintf("%s: %s has a +kcc:guess=%s reason=%s line but no open entry in apis/%s/%s; "+
				"fix the link and delete the guess line, or reopen the entry",
				path, k.kind, sourcelinks.GuessKind, k.reason, k.service, judgement.FileName))
		}
	}
	for k, path := range queued {
		if _, ok := guessed[k]; !ok {
			out = append(out, fmt.Sprintf("%s: open %s entry for %s has no +kcc:guess=%s line in the types file; "+
				"resolve the entry if the link was fixed",
				path, k.reason, k.kind, sourcelinks.GuessKind))
		}
	}
	sort.Strings(out)
	return out, nil
}

// TestSourceLinkGuessesMatchQueue checks the real files in the repo.
func TestSourceLinkGuessesMatchQueue(t *testing.T) {
	mismatches, err := sourceLinkMismatches("../../apis")
	if err != nil {
		t.Fatal(err)
	}
	for _, m := range mismatches {
		t.Error(m)
	}
}

func TestSourceLinkMismatches(t *testing.T) {
	dir := t.TempDir()
	writeQueueFile(t, dir, "bigtable", `entries:
- kind: BigtablePaired
  group: bigtable.cnrm.cloud.google.com
  reason: verify-resource-docs-link
  status: open
- kind: BigtableFixedNotResolved
  group: bigtable.cnrm.cloud.google.com
  reason: verify-resource-docs-link
  status: open
- kind: BigtableResolvedNotFixed
  group: bigtable.cnrm.cloud.google.com
  reason: verify-service-docs-link
  status: resolved
  resolution: edited
  note: replaced the link
- kind: BigtableDone
  group: bigtable.cnrm.cloud.google.com
  reason: verify-service-docs-link
  status: resolved
  resolution: edited
  note: replaced the link
`)
	writeTypes := func(kind string, guess string) {
		t.Helper()
		links := []sourcelinks.Link{
			{Key: sourcelinks.KeyProto, URL: "https://github.com/googleapis/googleapis/blob/abc/google/bigtable/admin/v2/instance.proto"},
			{Key: sourcelinks.KeyServiceDocs, URL: "https://cloud.google.com/bigtable/"},
			{Key: sourcelinks.KeyResourceDocs, URL: "https://docs.cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances"},
		}
		switch guess {
		case judgement.ReasonVerifyServiceDocsLink:
			links[1].Guess = guess
		case judgement.ReasonVerifyResourceDocsLink:
			links[2].Guess = guess
		}
		src := "// Copyright\n\n" + sourcelinks.Render(kind, links) + "\n\npackage v1alpha1\n"
		d := filepath.Join(dir, "bigtable", "v1alpha1")
		if err := os.MkdirAll(d, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, kind+"_types.go"), []byte(src), 0644); err != nil {
			t.Fatal(err)
		}
	}
	writeTypes("BigtablePaired", judgement.ReasonVerifyResourceDocsLink)
	writeTypes("BigtableFixedNotResolved", "")
	writeTypes("BigtableResolvedNotFixed", judgement.ReasonVerifyServiceDocsLink)
	writeTypes("BigtableDone", "")

	got, err := sourceLinkMismatches(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("got %d mismatches, want 2:\n%q", len(got), got)
	}
	for i, want := range []string{"BigtableFixedNotResolved", "BigtableResolvedNotFixed"} {
		if !bytes.Contains([]byte(got[i]), []byte(want)) {
			t.Errorf("mismatch %d = %q, want it to name %s", i, got[i], want)
		}
	}
}
