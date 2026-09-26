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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
)

// judgementQueueGlob finds the per-service work queues written by the bulk
// generator. The files are per service, not global, so that generating two
// services in parallel never produces a conflicting diff in the same file.
const judgementQueueGlob = "../../apis/*/needs_judgement_call.txt"

// judgementQueue tracks resources whose generated types have pending triage items.
// Resources with entries in this queue have their [refs] findings suppressed in
// TestMissingRefs until triaged.
type judgementQueue struct {
	// resources holds "<Kind>.<group>" keys.
	resources sets.String
	// entries counts individual field-level items, for reporting.
	entries int
}

func (q *judgementQueue) Has(kind, group string) bool {
	if q == nil {
		return false
	}
	return q.resources.Has(kind + "." + group)
}

// loadJudgementQueue reads every per-service queue file matching glob.
//
// Each line names one field a human still has to decide about:
//
//	kind=<Kind> group=<group>: field ".spec.foo" reason=<why>
//
// reason= is mandatory, for the same rationale as refs_deferred.txt: an entry
// with no stated reason is indistinguishable from an oversight, and this file
// suppresses a check, so an accidental entry silently weakens the ratchet.
func loadJudgementQueue(glob string) (*judgementQueue, error) {
	paths, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	q := &judgementQueue{resources: sets.NewString()}
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		for i, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			kind, group, err := parseJudgementEntry(line)
			if err != nil {
				return nil, fmt.Errorf("%s:%d: %w", path, i+1, err)
			}
			q.resources.Insert(kind + "." + group)
			q.entries++
		}
	}
	return q, nil
}

// parseJudgementEntry pulls the resource identity out of one queue line. The
// resource is keyed by Kind and group rather than CRD name because the generator
// writes this file before the CRD exists, and knows only the Kind and group.
func parseJudgementEntry(line string) (kind, group string, err error) {
	if !strings.Contains(line, "reason=") {
		return "", "", fmt.Errorf("entry has no reason=: %q", line)
	}
	head, _, _ := strings.Cut(line, ":")
	for _, tok := range strings.Fields(head) {
		k, v, ok := strings.Cut(tok, "=")
		if !ok {
			continue
		}
		switch k {
		case "kind":
			kind = v
		case "group":
			group = v
		}
	}
	if kind == "" || group == "" {
		return "", "", fmt.Errorf("entry must set kind= and group=: %q", line)
	}
	return kind, group, nil
}

// carryForwardSuppressed returns the baseline entries belonging to suppressed
// CRDs, so that suppressing a resource does not read as having fixed the
// findings it already owed.
func carryForwardSuppressed(baselinePath string, suppressedCRDs sets.String) ([]string, error) {
	data, err := os.ReadFile(baselinePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var out []string
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if name := crdNameFromEntry(line); name != "" && suppressedCRDs.Has(name) {
			out = append(out, line)
		}
	}
	return out, nil
}

// crdNameFromEntry pulls the CRD name out of an entry like
// `[refs] crd=foos.example.com version=v1beta1: field "..." should be a reference`.
func crdNameFromEntry(entry string) string {
	_, rest, ok := strings.Cut(entry, "crd=")
	if !ok {
		return ""
	}
	name, _, _ := strings.Cut(rest, " ")
	return strings.TrimSuffix(name, ":")
}

func TestCRDNameFromEntry(t *testing.T) {
	grid := []struct {
		entry string
		want  string
	}{
		{`[refs] crd=foos.example.com version=v1beta1: field ".spec.bar" should be a reference`, "foos.example.com"},
		{`[refs] crd=foos.example.com: field ".spec.bar" should be a reference`, "foos.example.com"},
		{`no crd token here`, ""},
	}
	for _, g := range grid {
		if got := crdNameFromEntry(g.entry); got != g.want {
			t.Errorf("crdNameFromEntry(%q) = %q, want %q", g.entry, got, g.want)
		}
	}
}

func TestCarryForwardSuppressed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missingrefs.txt")
	content := `# baseline
[refs] crd=queued.example.com version=v1alpha1: field ".spec.a" should be a reference
[refs] crd=other.example.com version=v1beta1: field ".spec.b" should be a reference
[refs] crd=queued.example.com version=v1alpha1: field ".spec.c" should be a reference
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	got, err := carryForwardSuppressed(path, sets.NewString("queued.example.com"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("carried %d entries, want 2: %v", len(got), got)
	}
	for _, e := range got {
		if !strings.Contains(e, "queued.example.com") {
			t.Errorf("carried an entry for a resource that is not suppressed: %q", e)
		}
	}

	// Nothing suppressed means nothing carried, so the normal path is untouched.
	none, err := carryForwardSuppressed(path, sets.NewString())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(none) != 0 {
		t.Errorf("expected no carried entries, got %v", none)
	}
}

func TestParseJudgementEntry(t *testing.T) {
	grid := []struct {
		name      string
		line      string
		wantKind  string
		wantGroup string
		wantErr   string
	}{
		{
			name:      "well formed",
			line:      `kind=NetworkServicesLBTrafficExtension group=networkservices.cnrm.cloud.google.com: field ".spec.forwardingRules" reason=possible-reference`,
			wantKind:  "NetworkServicesLBTrafficExtension",
			wantGroup: "networkservices.cnrm.cloud.google.com",
		},
		{
			// The field path contains a colon in some descriptions; only the
			// first colon separates the head.
			name:      "colon in the tail is fine",
			line:      `kind=Foo group=example.com: field ".spec.bar" reason=renamed: bar -> barRef`,
			wantKind:  "Foo",
			wantGroup: "example.com",
		},
		{
			name:    "missing reason is rejected",
			line:    `kind=Foo group=example.com: field ".spec.bar"`,
			wantErr: "no reason=",
		},
		{
			name:    "missing group is rejected",
			line:    `kind=Foo: field ".spec.bar" reason=x`,
			wantErr: "must set kind= and group=",
		},
		{
			name:    "missing kind is rejected",
			line:    `group=example.com: field ".spec.bar" reason=x`,
			wantErr: "must set kind= and group=",
		},
	}

	for _, g := range grid {
		t.Run(g.name, func(t *testing.T) {
			kind, group, err := parseJudgementEntry(g.line)
			if g.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), g.wantErr) {
					t.Fatalf("err = %v, want it to contain %q", err, g.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if kind != g.wantKind || group != g.wantGroup {
				t.Errorf("got kind=%q group=%q, want kind=%q group=%q", kind, group, g.wantKind, g.wantGroup)
			}
		})
	}
}

func TestLoadJudgementQueue(t *testing.T) {
	dir := t.TempDir()
	write := func(service, content string) {
		t.Helper()
		d := filepath.Join(dir, service)
		if err := os.MkdirAll(d, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, "needs_judgement_call.txt"), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	write("networkservices", `# Fields the generator could not decide.
kind=NetworkServicesLBTrafficExtension group=networkservices.cnrm.cloud.google.com: field ".spec.forwardingRules" reason=possible-reference

kind=NetworkServicesLBTrafficExtension group=networkservices.cnrm.cloud.google.com: field ".spec.labels" reason=deliberate-omission
`)
	write("dataproc", `kind=DataprocBatch group=dataproc.cnrm.cloud.google.com: field ".spec.serviceAccount" reason=possible-reference
`)

	q, err := loadJudgementQueue(filepath.Join(dir, "*", "needs_judgement_call.txt"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if q.entries != 3 {
		t.Errorf("entries = %d, want 3", q.entries)
	}
	// Two entries for one resource collapse to a single suppressed resource.
	if got := q.resources.Len(); got != 2 {
		t.Errorf("resources = %d, want 2", got)
	}
	if !q.Has("NetworkServicesLBTrafficExtension", "networkservices.cnrm.cloud.google.com") {
		t.Error("expected the networkservices resource to be queued")
	}
	if !q.Has("DataprocBatch", "dataproc.cnrm.cloud.google.com") {
		t.Error("expected the dataproc resource to be queued")
	}
	if q.Has("StorageBucket", "storage.cnrm.cloud.google.com") {
		t.Error("an unrelated resource must not be suppressed")
	}
}

func TestLoadJudgementQueueNoFiles(t *testing.T) {
	// The normal state of the repo today: no service has been bulk-generated, so
	// nothing is suppressed and TestMissingRefs behaves exactly as before.
	q, err := loadJudgementQueue(filepath.Join(t.TempDir(), "*", "needs_judgement_call.txt"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if q.entries != 0 || q.resources.Len() != 0 {
		t.Errorf("expected an empty queue, got %d entries across %d resources", q.entries, q.resources.Len())
	}
	if q.Has("StorageBucket", "storage.cnrm.cloud.google.com") {
		t.Error("empty queue must not suppress anything")
	}
}

func TestLoadJudgementQueueRejectsBadEntry(t *testing.T) {
	dir := t.TempDir()
	d := filepath.Join(dir, "svc")
	if err := os.MkdirAll(d, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(d, "needs_judgement_call.txt"),
		[]byte("kind=Foo group=example.com: field \".spec.bar\"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	if _, err := loadJudgementQueue(filepath.Join(dir, "*", "needs_judgement_call.txt")); err == nil {
		t.Fatal("expected an error for an entry with no reason=")
	} else if !strings.Contains(err.Error(), "needs_judgement_call.txt:1") {
		t.Errorf("error should name the file and line, got: %v", err)
	}
}

// TestJudgementQueueIsWellFormed validates the real files in the repo, so a
// malformed entry fails here rather than silently widening the suppression.
func TestJudgementQueueIsWellFormed(t *testing.T) {
	q, err := loadJudgementQueue(judgementQueueGlob)
	if err != nil {
		t.Fatalf("error loading judgement queues: %v", err)
	}
	t.Logf("%d open entries across %d resources", q.entries, q.resources.Len())
}
