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
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/util/sets"
)

// judgementQueueGlob finds the per-service work queues written by the bulk
// generator. The files are per service, not global, so that generating two
// services in parallel never produces a conflicting diff in the same file.
const judgementQueueGlob = "../../apis/*/" + judgement.FileName

// judgementQueue holds every entry from every per-service queue file.
type judgementQueue struct {
	entries []judgement.Entry
	// openRefFields holds refFieldKey(kind, group, field) for each open entry
	// with a reference reason.
	openRefFields sets.String
}

func refFieldKey(kind, group, field string) string {
	return kind + "." + group + "|" + field
}

// SuppressesRef reports whether an open queue entry already says this field
// may need to be a reference. TestMissingRefs skips such a finding: the
// generator has flagged it, and a person still has to decide.
//
// Only reference reasons count. An open entry about something else, such as
// a field placed in status or the untriaged-bulk-generation marker, does not
// hide a [refs] finding.
func (q *judgementQueue) SuppressesRef(kind, group, fieldPath string) bool {
	if q == nil {
		return false
	}
	return q.openRefFields.Has(refFieldKey(kind, group, fieldPath))
}

// loadJudgementQueue reads and validates every queue file matching glob.
// A file that does not validate is an error, so a typo cannot quietly change
// what is suppressed. So is a field whose reference entries do not all have
// the same status; see checkRefStatuses.
func loadJudgementQueue(glob string) (*judgementQueue, error) {
	paths, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	out := &judgementQueue{openRefFields: sets.NewString()}
	for _, path := range paths {
		q, err := judgement.Read(path)
		if err != nil {
			return nil, err
		}
		for _, e := range q.Entries {
			out.entries = append(out.entries, e)
			if e.IsOpen() && e.Kind != "" && e.Field != "" && judgement.IsReferenceReason(e.Reason) {
				out.openRefFields.Insert(refFieldKey(e.Kind, e.Group, e.Field))
			}
		}
	}
	if err := checkRefStatuses(out.entries); err != nil {
		return nil, err
	}
	return out, nil
}

// checkRefStatuses returns an error naming each field whose reference entries
// do not all have the same status. The generator files one entry per signal
// that a field may be a reference, and keeps them all because the signals are
// independent. SuppressesRef hides the field while any of them is open, so
// resolving one and forgetting another would keep the field hidden.
func checkRefStatuses(entries []judgement.Entry) error {
	var fields []string
	byField := map[string][]judgement.Entry{}
	for _, e := range entries {
		if e.Kind == "" || e.Field == "" || !judgement.IsReferenceReason(e.Reason) {
			continue
		}
		key := refFieldKey(e.Kind, e.Group, e.Field)
		if _, ok := byField[key]; !ok {
			fields = append(fields, key)
		}
		byField[key] = append(byField[key], e)
	}

	var problems []string
	for _, key := range fields {
		statuses := sets.NewString()
		var parts []string
		for _, e := range byField[key] {
			statuses.Insert(string(e.Status))
			parts = append(parts, e.Reason+" is "+string(e.Status))
		}
		if statuses.Len() > 1 {
			e := byField[key][0]
			problems = append(problems, fmt.Sprintf("  kind %s, group %s, field %s: %s", e.Kind, e.Group, e.Field, strings.Join(parts, ", ")))
		}
	}
	if len(problems) == 0 {
		return nil
	}
	return fmt.Errorf("the reference entries for a field must all have the same status, because TestMissingRefs skips the field while any of them is open. Resolve the open ones too, or reopen the resolved ones:\n%s", strings.Join(problems, "\n"))
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

// carryForwardSuppressed returns the baseline entries that match a suppressed
// finding, so that suppressing a finding does not read as having fixed it.
func carryForwardSuppressed(baselinePath string, suppressed sets.String) ([]string, error) {
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
		if suppressed.Has(line) {
			out = append(out, line)
		}
	}
	return out, nil
}

func TestCarryForwardSuppressed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missingrefs.txt")
	a := `[refs] crd=queued.example.com version=v1alpha1: field ".spec.a" should be a reference`
	b := `[refs] crd=other.example.com version=v1beta1: field ".spec.b" should be a reference`
	c := `[refs] crd=queued.example.com version=v1alpha1: field ".spec.c" should be a reference`
	if err := os.WriteFile(path, []byte("# baseline\n"+a+"\n"+b+"\n"+c+"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Only .spec.a is suppressed. .spec.c belongs to the same CRD but has no
	// queue entry, so it is checked as usual and not carried.
	got, err := carryForwardSuppressed(path, sets.NewString(a))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || got[0] != a {
		t.Fatalf("carried %v, want only %q", got, a)
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
	const (
		lb      = "NetworkServicesLBTrafficExtension"
		lbGroup = "networkservices.cnrm.cloud.google.com"
	)
	dir := t.TempDir()
	writeQueueFile(t, dir, "networkservices", `entries:
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: open
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.forwardingRules
  reason: possible-reference-by-description
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

	if len(q.entries) != 5 {
		t.Errorf("entries = %d, want 5", len(q.entries))
	}
	for _, tc := range []struct {
		name         string
		kind, group  string
		field        string
		wantSuppress bool
	}{
		{"open reference entry", lb, lbGroup, ".spec.forwardingRules", true},
		{"same resource, field with no entry", lb, lbGroup, ".spec.network", false},
		{"open entry with a non-reference reason", lb, lbGroup, ".spec.labels", false},
		{"resolved reference entry", "DataprocBatch", "dataproc.cnrm.cloud.google.com", ".spec.serviceAccount", false},
		{"unrelated resource", "StorageBucket", "storage.cnrm.cloud.google.com", ".spec.forwardingRules", false},
	} {
		if got := q.SuppressesRef(tc.kind, tc.group, tc.field); got != tc.wantSuppress {
			t.Errorf("%s: SuppressesRef(%s, %s) = %v, want %v", tc.name, tc.kind, tc.field, got, tc.wantSuppress)
		}
	}
	// The untriaged marker and the message entry name no field of a Kind, so
	// only the one reference entry counts.
	if got := q.openRefFields.Len(); got != 1 {
		t.Errorf("open reference fields = %d, want 1", got)
	}
}

func TestLoadJudgementQueueNoFiles(t *testing.T) {
	// No service has been bulk-generated, so nothing is suppressed and
	// TestMissingRefs behaves exactly as before.
	q, err := loadJudgementQueue(filepath.Join(t.TempDir(), "*", judgement.FileName))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(q.entries) != 0 || q.openRefFields.Len() != 0 {
		t.Errorf("expected an empty queue, got %d entries and %d suppressed fields", len(q.entries), q.openRefFields.Len())
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

// TestLoadJudgementQueueRejectsMixedReferenceStatuses pins the error for a
// field whose reference entries have different statuses. It names each such
// field with all of its reference entries. network's entries agree, so it is
// not named.
func TestLoadJudgementQueueRejectsMixedReferenceStatuses(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	writeQueueFile(t, dir, "networkservices", `entries:
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.forwardingRules
  reason: possible-reference
  status: resolved
  resolution: accepted
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.forwardingRules
  reason: possible-reference-by-description
  status: open
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.network
  reason: possible-reference-by-name
  status: open
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.network
  reason: possible-reference-by-description
  status: open
`)
	writeQueueFile(t, dir, "dataproc", `entries:
- kind: DataprocBatch
  group: dataproc.cnrm.cloud.google.com
  field: .spec.serviceAccount
  reason: possible-reference
  status: open
- kind: DataprocBatch
  group: dataproc.cnrm.cloud.google.com
  field: .spec.serviceAccount
  reason: possible-reference-by-description
  status: resolved
  resolution: edited
  note: changed to serviceAccountRef
- kind: DataprocBatch
  group: dataproc.cnrm.cloud.google.com
  field: .spec.serviceAccount
  reason: possible-reference-by-name
  status: open
`)
	want := `the reference entries for a field must all have the same status, because TestMissingRefs skips the field while any of them is open. Resolve the open ones too, or reopen the resolved ones:
  kind DataprocBatch, group dataproc.cnrm.cloud.google.com, field .spec.serviceAccount: possible-reference is open, possible-reference-by-description is resolved, possible-reference-by-name is open
  kind NetworkServicesLBTrafficExtension, group networkservices.cnrm.cloud.google.com, field .spec.forwardingRules: possible-reference is resolved, possible-reference-by-description is open`

	// Act
	_, err := loadJudgementQueue(filepath.Join(dir, "*", judgement.FileName))

	// Assert
	if err == nil {
		t.Fatal("loadJudgementQueue() succeeded, want an error")
	}
	if diff := cmp.Diff(want, err.Error()); diff != "" {
		t.Errorf("loadJudgementQueue() error mismatch (-want +got):\n%s", diff)
	}
}

// TestLoadJudgementQueueAcceptsMatchingReferenceStatuses pins what the status
// check allows: reference entries that are all open or all resolved, and an
// entry with another reason next to them, whatever its status. A field is
// suppressed when its reference entries are open.
func TestLoadJudgementQueueAcceptsMatchingReferenceStatuses(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	writeQueueFile(t, dir, "networkservices", `entries:
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.forwardingRules
  reason: possible-reference
  status: open
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.forwardingRules
  reason: possible-reference-by-description
  status: open
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.network
  reason: possible-reference
  status: resolved
  resolution: accepted
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.network
  reason: possible-reference-by-name
  status: resolved
  resolution: accepted
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.service
  reason: possible-reference-by-description
  status: open
- kind: NetworkServicesLBTrafficExtension
  group: networkservices.cnrm.cloud.google.com
  field: .spec.service
  reason: output-only-mentioned-in-comment
  status: resolved
  resolution: accepted
`)
	want := []string{
		"NetworkServicesLBTrafficExtension.networkservices.cnrm.cloud.google.com|.spec.forwardingRules",
		"NetworkServicesLBTrafficExtension.networkservices.cnrm.cloud.google.com|.spec.service",
	}

	// Act
	q, err := loadJudgementQueue(filepath.Join(dir, "*", judgement.FileName))

	// Assert
	if err != nil {
		t.Fatalf("loadJudgementQueue() error: %v", err)
	}
	if diff := cmp.Diff(want, q.openRefFields.List()); diff != "" {
		t.Errorf("open reference fields mismatch (-want +got):\n%s", diff)
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
	t.Logf("%d entries, %d open reference fields\n%s", len(q.entries), q.openRefFields.Len(), q.summary())
}
