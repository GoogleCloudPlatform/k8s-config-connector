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
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	"k8s.io/apimachinery/pkg/util/sets"
)

// judgementQueueGlob finds the per-service work queues written by the bulk
// generator. The files are per service, not global, so that generating two
// services in parallel never produces a conflicting diff in the same file.
const judgementQueueGlob = "../../apis/*/" + judgement.FileName

// judgementQueue holds every entry from every per-service queue file.
type judgementQueue struct {
	entries []judgement.Entry
	// openResources holds "<Kind>.<group>" for each resource with at least
	// one open entry.
	openResources sets.String
}

// Has reports whether a resource still has open entries. TestMissingRefs
// skips the [refs] findings of such a resource.
func (q *judgementQueue) Has(kind, group string) bool {
	if q == nil {
		return false
	}
	return q.openResources.Has(kind + "." + group)
}

// loadJudgementQueue reads and validates every queue file matching glob.
// A file that does not validate is an error, so a typo cannot quietly change
// what is suppressed.
func loadJudgementQueue(glob string) (*judgementQueue, error) {
	paths, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	out := &judgementQueue{openResources: sets.NewString()}
	for _, path := range paths {
		q, err := judgement.Read(path)
		if err != nil {
			return nil, err
		}
		for _, e := range q.Entries {
			out.entries = append(out.entries, e)
			if e.IsOpen() && e.Kind != "" {
				out.openResources.Insert(e.Kind + "." + e.Group)
			}
		}
	}
	return out, nil
}

// summary counts entries by reason, split into open and each resolution.
// It shows which guesses people keep and which they rewrite, which is what
// tells us where the generator needs work.
func (q *judgementQueue) summary() string {
	counts := map[string]map[string]int{}
	for _, e := range q.entries {
		state := string(e.Status)
		if e.Status == judgement.StatusResolved {
			state = string(e.Resolution)
		}
		if counts[e.Reason] == nil {
			counts[e.Reason] = map[string]int{}
		}
		counts[e.Reason][state]++
	}
	reasons := make([]string, 0, len(counts))
	for r := range counts {
		reasons = append(reasons, r)
	}
	sort.Strings(reasons)
	var b strings.Builder
	for _, r := range reasons {
		b.WriteString(r)
		for _, state := range []string{"open", "accepted", "edited", "deferred", "not-applicable"} {
			if n := counts[r][state]; n > 0 {
				fmt.Fprintf(&b, " %s=%d", state, n)
			}
		}
		b.WriteString("\n")
	}
	return b.String()
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

func writeQueueFile(t *testing.T, dir, service, content string) {
	t.Helper()
	d := filepath.Join(dir, service)
	if err := os.MkdirAll(d, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(d, judgement.FileName), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadJudgementQueue(t *testing.T) {
	dir := t.TempDir()
	writeQueueFile(t, dir, "networkservices", `entries:
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.forwardingRules
  reason: possible-reference
  status: open
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.labels
  reason: deliberate-omission
  status: open
- protoMessage: google.cloud.networkservices.v1.ExtensionChain
  field: name
  reason: possible-reference-by-sibling
  status: open
`)
	writeQueueFile(t, dir, "dataproc", `entries:
- kind: DataprocBatch
  group: dataproc.cnrm.cloud.google.com
  field: .spec.serviceAccount
  reason: possible-reference
  status: resolved
  resolution: edited
  note: changed to serviceAccountRef
`)

	q, err := loadJudgementQueue(filepath.Join(dir, "*", judgement.FileName))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(q.entries) != 4 {
		t.Errorf("entries = %d, want 4", len(q.entries))
	}
	if !q.Has("NetworkServicesLBTrafficExtension", "networkservices.cnrm.cloud.google.com") {
		t.Error("expected the networkservices resource to be queued")
	}
	// Every dataproc entry is resolved, so the resource is checked in full.
	if q.Has("DataprocBatch", "dataproc.cnrm.cloud.google.com") {
		t.Error("a resource whose entries are all resolved must not be suppressed")
	}
	if q.Has("StorageBucket", "storage.cnrm.cloud.google.com") {
		t.Error("an unrelated resource must not be suppressed")
	}
	// A message entry names no Kind, so it suppresses nothing.
	if got := q.openResources.Len(); got != 1 {
		t.Errorf("open resources = %d, want 1", got)
	}
}

func TestLoadJudgementQueueNoFiles(t *testing.T) {
	// No service has been bulk-generated, so nothing is suppressed and
	// TestMissingRefs behaves exactly as before.
	q, err := loadJudgementQueue(filepath.Join(t.TempDir(), "*", judgement.FileName))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(q.entries) != 0 || q.openResources.Len() != 0 {
		t.Errorf("expected an empty queue, got %d entries across %d resources", len(q.entries), q.openResources.Len())
	}
}

func TestLoadJudgementQueueRejectsBadEntry(t *testing.T) {
	dir := t.TempDir()
	writeQueueFile(t, dir, "svc", "entries:\n- kind: Foo\n  group: example.com\n  field: .spec.bar\n  status: open\n")

	_, err := loadJudgementQueue(filepath.Join(dir, "*", judgement.FileName))
	if err == nil {
		t.Fatal("expected an error for an entry with no reason")
	}
	if !strings.Contains(err.Error(), filepath.Join("svc", judgement.FileName)) {
		t.Errorf("error should name the file, got: %v", err)
	}
}

// TestJudgementQueueIsWellFormed validates the real files in the repo, so a
// malformed entry fails here rather than silently changing what is
// suppressed. Run it with -v to see the counts by reason.
func TestJudgementQueueIsWellFormed(t *testing.T) {
	q, err := loadJudgementQueue(judgementQueueGlob)
	if err != nil {
		t.Fatalf("error loading judgement queues: %v", err)
	}
	t.Logf("%d entries, %d resources with open entries\n%s", len(q.entries), q.openResources.Len(), q.summary())
}
