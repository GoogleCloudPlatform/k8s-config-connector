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
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/judgement"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/yaml"
)

// TestCRDBackwardCompatibility verifies that CRD changes relative to the base
// branch are backward-compatible for all pre-existing and already-triaged Kinds.
//
// Kinds that are currently untriaged in apis/<service>/judgement_queue.yaml
// (or are being triaged for the first time in the current branch from an open
// untriaged-bulk-generation entry at the base commit) may make breaking changes
// to their own CRD schema, provided they do not break any shared struct used by
// a pre-existing or already-triaged Kind.
func TestCRDBackwardCompatibility(t *testing.T) {
	t.Parallel()

	repoRoot := findGitRepoRoot(t)
	baseRev := resolveBaseRevision(t, repoRoot)
	changedCRDs := changedCRDFiles(t, repoRoot, baseRev)
	if len(changedCRDs) == 0 {
		return
	}

	var failures []string
	for _, relPath := range changedCRDs {
		if msg := checkCRDFileBackwardCompatibility(repoRoot, baseRev, relPath); msg != "" {
			failures = append(failures, msg)
		}
	}

	if len(failures) > 0 {
		t.Fatalf("backward-incompatible CRD changes detected against %s:\n\n%s\n\n"+
			"Hint: If this change was caused by modifying a shared nested struct in types.generated.go "+
			"while triaging another resource, keep the original struct untouched in types.generated.go "+
			"and define a forked struct for your resource using the 4-suffix convention "+
			"(<Struct>WithRequired, <Struct>WithRefs, <Struct>Spec, or <Struct>WithMultipleChanges).",
			baseRev, strings.Join(failures, "\n\n"))
	}
}

func checkCRDFileBackwardCompatibility(repoRoot, baseRev, relCRDPath string) string {
	oldBytes, isNew, err := gitShowFile(repoRoot, baseRev, relCRDPath)
	if err != nil {
		return fmt.Sprintf("%s: reading base CRD at %s: %v", relCRDPath, baseRev, err)
	}
	if isNew {
		return ""
	}

	oldCRD, err := unmarshalCRD(oldBytes)
	if err != nil {
		return fmt.Sprintf("%s: parsing base CRD at %s: %v", relCRDPath, baseRev, err)
	}

	kind := oldCRD.Spec.Names.Kind
	group := oldCRD.Spec.Group
	baseQueue := loadServiceQueueAtRef(repoRoot, baseRev, group)
	headQueue := loadServiceQueueFromDisk(repoRoot, group)
	if isKindUntriaged(kind, group, baseQueue, headQueue) {
		return ""
	}

	newBytes, err := os.ReadFile(filepath.Join(repoRoot, relCRDPath))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Sprintf("%s (%s.%s): CRD file was deleted for a triaged or pre-existing resource", relCRDPath, kind, group)
		}
		return fmt.Sprintf("%s: reading current CRD: %v", relCRDPath, err)
	}

	newCRD, err := unmarshalCRD(newBytes)
	if err != nil {
		return fmt.Sprintf("%s: parsing current CRD: %v", relCRDPath, err)
	}

	diffs := compareCRDBackwardCompatibility(oldCRD, newCRD)
	if len(diffs) == 0 {
		return ""
	}
	return fmt.Sprintf("%s (%s.%s):\n  %s", relCRDPath, kind, group, strings.Join(diffs, "\n  "))
}

// isKindUntriaged reports whether a Kind that already existed at baseRev is
// exempt from backward-compatibility enforcement because it is still untriaged
// or is being triaged for the first time from an open untriaged-bulk-generation
// entry at baseRev.
func isKindUntriaged(kind, group string, baseQueue, headQueue *judgement.Queue) bool {
	baseStatus := untriagedEntryStatus(baseQueue, kind, group)
	headStatus := untriagedEntryStatus(headQueue, kind, group)
	if baseStatus == judgement.StatusOpen {
		return true
	}
	if baseStatus == judgement.StatusResolved {
		return false
	}
	return headStatus == judgement.StatusOpen
}

func untriagedEntryStatus(q *judgement.Queue, kind, group string) judgement.Status {
	if q == nil {
		return ""
	}
	for _, e := range q.Entries {
		if e.Kind == kind && e.Group == group && e.Reason == judgement.ReasonUntriagedBulkGeneration {
			return e.Status
		}
	}
	return ""
}

func loadServiceQueueAtRef(repoRoot, rev, group string) *judgement.Queue {
	relPath := serviceQueueRelPath(group)
	if relPath == "" {
		return nil
	}
	data, isNew, err := gitShowFile(repoRoot, rev, relPath)
	if err != nil || isNew {
		return nil
	}
	var q judgement.Queue
	if err := yaml.Unmarshal(data, &q); err != nil {
		return nil
	}
	return &q
}

func loadServiceQueueFromDisk(repoRoot, group string) *judgement.Queue {
	relPath := serviceQueueRelPath(group)
	if relPath == "" {
		return nil
	}
	q, err := judgement.Read(filepath.Join(repoRoot, relPath))
	if err != nil {
		return nil
	}
	return q
}

func serviceQueueRelPath(group string) string {
	service := strings.TrimSuffix(group, ".cnrm.cloud.google.com")
	if service == "" || service == group {
		return ""
	}
	return filepath.Join("apis", service, judgement.FileName)
}

// compareCRDBackwardCompatibility returns breaking schema diffs between oldCRD and newCRD.
func compareCRDBackwardCompatibility(oldCRD, newCRD *apiextensionsv1.CustomResourceDefinition) []string {
	var diffs []string

	oldVersions := flattenCRDVersions(oldCRD)
	newVersions := flattenCRDVersions(newCRD)

	for _, vName := range slices.Sorted(maps.Keys(oldVersions)) {
		if _, ok := newVersions[vName]; !ok {
			diffs = append(diffs, fmt.Sprintf("version %q removed", vName))
		}
	}

	for _, vName := range slices.Sorted(maps.Keys(oldVersions)) {
		newPaths, ok := newVersions[vName]
		if !ok {
			continue
		}
		diffs = append(diffs, diffVersionBackwardCompatibility(vName, oldVersions[vName], newPaths)...)
	}
	return diffs
}

func diffVersionBackwardCompatibility(version string, oldPaths, newPaths map[string]string) []string {
	var diffs []string
	prefix := fmt.Sprintf("[%s] ", version)

	for _, path := range slices.Sorted(maps.Keys(oldPaths)) {
		if _, _, ok := splitRequiredMarkerPath(path); ok {
			continue
		}
		oldType := oldPaths[path]
		newType, ok := newPaths[path]
		if !ok {
			if !hasRemovedAncestor(path, oldPaths, newPaths) {
				diffs = append(diffs, fmt.Sprintf("%sfield removed: %s (was %s)", prefix, path, oldType))
			}
			continue
		}
		if oldType != newType && !isAllowedStatusIntegerWidening(path, oldType, newType) {
			diffs = append(diffs, fmt.Sprintf("%sfield type changed: %s (%s -> %s)", prefix, path, oldType, newType))
		}
	}

	for _, path := range slices.Sorted(maps.Keys(newPaths)) {
		if parentPath, fieldPath, ok := splitRequiredMarkerPath(path); ok {
			if _, wasRequired := oldPaths[path]; !wasRequired {
				if parentPath == "" || oldPaths[parentPath] != "" {
					diffs = append(diffs, fmt.Sprintf("%sfield made required: %s", prefix, fieldPath))
				}
			}
			continue
		}
		if _, existed := oldPaths[path]; !existed && newPaths[path] == "empty-object" {
			diffs = append(diffs, fmt.Sprintf("%sfield added is an empty object: %s", prefix, path))
		}
	}
	return diffs
}

func flattenCRDVersions(crd *apiextensionsv1.CustomResourceDefinition) map[string]map[string]string {
	out := make(map[string]map[string]string)
	for _, v := range crd.Spec.Versions {
		if v.Schema == nil || v.Schema.OpenAPIV3Schema == nil {
			continue
		}
		paths := make(map[string]string)
		flattenOpenAPISchema("", walkOpenAPISchema(v.Schema.OpenAPIV3Schema, 0), paths)
		out[v.Name] = paths
	}
	return out
}

func walkOpenAPISchema(s *apiextensionsv1.JSONSchemaProps, depth int) any {
	if s == nil {
		return "unknown"
	}
	if depth > 50 {
		return "unknown (depth limit exceeded)"
	}

	var res any
	switch {
	case s.XPreserveUnknownFields != nil && *s.XPreserveUnknownFields:
		res = "json"
	case len(s.Properties) > 0:
		m := make(map[string]any, len(s.Properties))
		for k, v := range s.Properties {
			val := v
			m[k] = walkOpenAPISchema(&val, depth+1)
		}
		res = m
	case s.Type == "array":
		if s.Items != nil && s.Items.Schema != nil {
			res = []any{walkOpenAPISchema(s.Items.Schema, depth+1)}
		} else {
			res = []any{"any"}
		}
	case s.AdditionalProperties != nil:
		res = walkAdditionalProperties(s.AdditionalProperties, depth)
	case s.Type == "object":
		res = "empty-object"
	default:
		res = walkScalarOrComposite(s)
	}

	if len(s.Required) > 0 {
		m := make(map[string]any)
		if existing, ok := res.(map[string]any); ok {
			maps.Copy(m, existing)
		} else {
			m[":type"] = res
		}
		for _, req := range s.Required {
			m[":required."+req] = "required"
		}
		return m
	}
	return res
}

func walkAdditionalProperties(ap *apiextensionsv1.JSONSchemaPropsOrBool, depth int) any {
	if ap.Schema != nil {
		val := walkOpenAPISchema(ap.Schema, depth+1)
		if str, ok := val.(string); ok {
			return fmt.Sprintf("map[string]%s", str)
		}
		return map[string]any{"KEY": val}
	}
	if !ap.Allows {
		return "map[string]none"
	}
	return "map[string]any"
}

func walkScalarOrComposite(s *apiextensionsv1.JSONSchemaProps) any {
	if s.Type == "" {
		return "any"
	}
	if s.Format == "int32" || s.Format == "int64" {
		return s.Format
	}
	if s.Format != "" {
		return fmt.Sprintf("%s (%s)", s.Type, s.Format)
	}
	return s.Type
}

func flattenOpenAPISchema(path string, schema any, out map[string]string) {
	switch v := schema.(type) {
	case map[string]any:
		if path != "" {
			out[path] = "object"
		}
		for k, child := range v {
			if k == ":type" {
				flattenOpenAPISchema(path, child, out)
				continue
			}
			childPath := k
			if path != "" {
				childPath = path + "." + k
			}
			flattenOpenAPISchema(childPath, child, out)
		}
	case []any:
		if path != "" {
			out[path] = "array"
		}
		for _, item := range v {
			flattenOpenAPISchema(path+"[]", item, out)
		}
	case string:
		out[path] = v
	}
}

func splitRequiredMarkerPath(path string) (parentPath, fieldPath string, ok bool) {
	if parent, field, found := strings.Cut(path, ".:required."); found {
		return parent, parent + "." + field, true
	}
	if field, found := strings.CutPrefix(path, ":required."); found {
		return "", field, true
	}
	return "", "", false
}

func hasRemovedAncestor(path string, oldPaths, newPaths map[string]string) bool {
	for i := 0; i < len(path); i++ {
		if path[i] == '.' || path[i] == '[' {
			anc := path[:i]
			if _, inOld := oldPaths[anc]; inOld {
				if _, inNew := newPaths[anc]; !inNew {
					return true
				}
			}
		}
	}
	return false
}

func isAllowedStatusIntegerWidening(path, oldType, newType string) bool {
	if path != "status" && !strings.HasPrefix(path, "status.") && !strings.HasPrefix(path, "status[") {
		return false
	}
	isInt := func(t string) bool {
		return t == "integer" || t == "int32" || t == "int64"
	}
	if isInt(oldType) && isInt(newType) {
		return !(oldType == "int64" && newType == "int32")
	}
	return false
}

func unmarshalCRD(data []byte) (*apiextensionsv1.CustomResourceDefinition, error) {
	var crd apiextensionsv1.CustomResourceDefinition
	if err := yaml.Unmarshal(data, &crd); err != nil {
		return nil, err
	}
	return &crd, nil
}

func findGitRepoRoot(t *testing.T) string {
	t.Helper()
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Skipf("skipping TestCRDBackwardCompatibility: not in a git repository (%v)", err)
	}
	return strings.TrimSpace(string(out))
}

func resolveBaseRevision(t *testing.T, repoRoot string) string {
	t.Helper()
	for _, env := range []string{"COMPARISON_REF", "BASE_COMMIT"} {
		if ref := strings.TrimSpace(os.Getenv(env)); ref != "" && !strings.HasPrefix(ref, "-") {
			if gitRefExists(repoRoot, ref) {
				return mergeBaseOrRef(repoRoot, ref)
			}
		}
	}
	if baseBranch := strings.TrimSpace(os.Getenv("GITHUB_BASE_REF")); baseBranch != "" && !strings.HasPrefix(baseBranch, "-") {
		candidate := "origin/" + baseBranch
		if !gitRefExists(repoRoot, candidate) {
			_ = exec.Command("git", "-C", repoRoot, "fetch", "origin", baseBranch, "--depth=1", "--quiet").Run()
		}
		if gitRefExists(repoRoot, candidate) {
			return mergeBaseOrRef(repoRoot, candidate)
		}
	}
	for _, candidate := range []string{"upstream/master", "origin/master", "master"} {
		if gitRefExists(repoRoot, candidate) {
			return mergeBaseOrRef(repoRoot, candidate)
		}
	}
	t.Skip("skipping TestCRDBackwardCompatibility: no git base revision found")
	return ""
}

func gitRefExists(repoRoot, ref string) bool {
	return exec.Command("git", "-C", repoRoot, "rev-parse", "--verify", "--quiet", ref).Run() == nil
}

func mergeBaseOrRef(repoRoot, ref string) string {
	out, err := exec.Command("git", "-C", repoRoot, "merge-base", ref, "HEAD").Output()
	if err == nil {
		if mb := strings.TrimSpace(string(out)); mb != "" {
			return mb
		}
	}
	return ref
}

func changedCRDFiles(t *testing.T, repoRoot, baseRev string) []string {
	t.Helper()
	out, err := exec.Command("git", "-C", repoRoot, "diff", "-z", "--name-only", baseRev, "--", "config/crds/resources").Output()
	if err != nil {
		t.Fatalf("git diff --name-only %s: %v", baseRev, err)
	}
	var files []string
	for _, part := range bytes.Split(out, []byte{0}) {
		if s := string(part); s != "" && strings.HasSuffix(s, ".yaml") {
			files = append(files, s)
		}
	}
	slices.Sort(files)
	return files
}

func gitShowFile(repoRoot, rev, relPath string) ([]byte, bool, error) {
	cmd := exec.Command("git", "-C", repoRoot, "show", fmt.Sprintf("%s:%s", rev, relPath))
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			stderr := string(exitErr.Stderr)
			if strings.Contains(stderr, "exists on disk, but not in") ||
				strings.Contains(stderr, "does not exist in") {
				return nil, true, nil
			}
		}
		return nil, false, err
	}
	return out, false, nil
}

func TestIsKindUntriaged(t *testing.T) {
	const (
		kind  = "FooResource"
		group = "foo.cnrm.cloud.google.com"
	)
	queueWithStatus := func(st judgement.Status) *judgement.Queue {
		return &judgement.Queue{
			Entries: []judgement.Entry{
				{
					Kind:   kind,
					Group:  group,
					Reason: judgement.ReasonUntriagedBulkGeneration,
					Status: st,
				},
			},
		}
	}

	for _, tc := range []struct {
		name      string
		baseQueue *judgement.Queue
		headQueue *judgement.Queue
		want      bool
	}{
		{
			name:      "pre-existing resource with no queue entry is not untriaged",
			baseQueue: nil,
			headQueue: nil,
			want:      false,
		},
		{
			name:      "untriaged at both base and head is exempt",
			baseQueue: queueWithStatus(judgement.StatusOpen),
			headQueue: queueWithStatus(judgement.StatusOpen),
			want:      true,
		},
		{
			name:      "triage PR transitioning open at base to resolved at head is exempt",
			baseQueue: queueWithStatus(judgement.StatusOpen),
			headQueue: queueWithStatus(judgement.StatusResolved),
			want:      true,
		},
		{
			name:      "already-triaged resource (resolved at base and head) is not exempt",
			baseQueue: queueWithStatus(judgement.StatusResolved),
			headQueue: queueWithStatus(judgement.StatusResolved),
			want:      false,
		},
		{
			name:      "already-triaged resource at base cannot be re-opened at head to bypass check",
			baseQueue: queueWithStatus(judgement.StatusResolved),
			headQueue: queueWithStatus(judgement.StatusOpen),
			want:      false,
		},
		{
			name:      "open at head with no prior base entry is exempt",
			baseQueue: nil,
			headQueue: queueWithStatus(judgement.StatusOpen),
			want:      true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isKindUntriaged(kind, group, tc.baseQueue, tc.headQueue); got != tc.want {
				t.Errorf("isKindUntriaged() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCompareCRDBackwardCompatibility(t *testing.T) {
	baseYAML := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.foo.cnrm.cloud.google.com
spec:
  group: foo.cnrm.cloud.google.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            required:
            - existingReq
            properties:
              existingReq:
                type: string
              sharedConfig:
                type: object
                properties:
                  network:
                    type: string
                  outputOnlyField:
                    type: string
                  optionalField:
                    type: string
          status:
            type: object
            properties:
              observedGeneration:
                type: integer
`
	baseCRD, err := unmarshalCRD([]byte(baseYAML))
	if err != nil {
		t.Fatalf("unmarshal baseCRD: %v", err)
	}

	t.Run("compatible additions and relaxing required to optional", func(t *testing.T) {
		compatibleYAML := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.foo.cnrm.cloud.google.com
spec:
  group: foo.cnrm.cloud.google.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              existingReq:
                type: string
              newOptionalField:
                type: string
              newOptionalStruct:
                type: object
                required:
                - innerReq
                properties:
                  innerReq:
                    type: string
              sharedConfig:
                type: object
                properties:
                  network:
                    type: string
                  outputOnlyField:
                    type: string
                  optionalField:
                    type: string
          status:
            type: object
            properties:
              observedGeneration:
                type: integer
                format: int64
`
		newCRD, err := unmarshalCRD([]byte(compatibleYAML))
		if err != nil {
			t.Fatalf("unmarshal compatibleCRD: %v", err)
		}
		if diffs := compareCRDBackwardCompatibility(baseCRD, newCRD); len(diffs) != 0 {
			t.Errorf("expected no diffs for compatible CRD update, got: %v", diffs)
		}
	})

	t.Run("breaking changes on shared struct (field removed, replaced by ref, or made required)", func(t *testing.T) {
		breakingYAML := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: foos.foo.cnrm.cloud.google.com
spec:
  group: foo.cnrm.cloud.google.com
  names:
    kind: Foo
    plural: foos
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            required:
            - existingReq
            properties:
              existingReq:
                type: string
              sharedConfig:
                type: object
                required:
                - optionalField
                properties:
                  networkRef:
                    type: object
                    properties:
                      external:
                        type: string
                  optionalField:
                    type: string
          status:
            type: object
            properties:
              observedGeneration:
                type: integer
`
		newCRD, err := unmarshalCRD([]byte(breakingYAML))
		if err != nil {
			t.Fatalf("unmarshal breakingCRD: %v", err)
		}
		diffs := compareCRDBackwardCompatibility(baseCRD, newCRD)
		want := []string{
			"[v1alpha1] field removed: spec.sharedConfig.network (was string)",
			"[v1alpha1] field removed: spec.sharedConfig.outputOnlyField (was string)",
			"[v1alpha1] field made required: spec.sharedConfig.optionalField",
		}
		if !slices.Equal(diffs, want) {
			t.Errorf("compareCRDBackwardCompatibility() diffs = %v, want %v", diffs, want)
		}
	})
}

func TestCheckCRDFileBackwardCompatibilityGitRepo(t *testing.T) {
	repo := t.TempDir()
	runGit := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", append([]string{"-C", repo}, args...)...)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	writeRepoFile := func(relPath, content string) {
		t.Helper()
		full := filepath.Join(repo, relPath)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	crdForKind := func(kind, fieldName string) string {
		return fmt.Sprintf(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: %ss.svc.cnrm.cloud.google.com
spec:
  group: svc.cnrm.cloud.google.com
  names:
    kind: %s
    plural: %ss
  scope: Namespaced
  versions:
  - name: v1alpha1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              %s:
                type: string
`, strings.ToLower(kind), kind, strings.ToLower(kind), fieldName)
	}

	runGit("init")
	runGit("config", "user.email", "test@example.com")
	runGit("config", "user.name", "Test")

	const (
		untriagedCRDPath   = "config/crds/resources/untriaged.yaml"
		triagedCRDPath     = "config/crds/resources/triaged.yaml"
		preexistingCRDPath = "config/crds/resources/preexisting.yaml"
		newCRDPath         = "config/crds/resources/brandnew.yaml"
		queuePath          = "apis/svc/judgement_queue.yaml"
	)

	writeRepoFile(untriagedCRDPath, crdForKind("SvcUntriaged", "rawNetwork"))
	writeRepoFile(triagedCRDPath, crdForKind("SvcTriaged", "sharedField"))
	writeRepoFile(preexistingCRDPath, crdForKind("SvcPreexisting", "sharedField"))
	writeRepoFile(queuePath, `entries:
- kind: SvcTriaged
  group: svc.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: resolved
  resolution: accepted
- kind: SvcUntriaged
  group: svc.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: open
`)
	runGit("add", ".")
	runGit("commit", "-m", "base commit")

	// Simulate a triage PR that resolves SvcUntriaged while mutating all CRDs.
	writeRepoFile(untriagedCRDPath, crdForKind("SvcUntriaged", "networkRef"))
	writeRepoFile(triagedCRDPath, crdForKind("SvcTriaged", "renamedField"))
	writeRepoFile(preexistingCRDPath, crdForKind("SvcPreexisting", "renamedField"))
	writeRepoFile(newCRDPath, crdForKind("SvcBrandNew", "anyField"))
	writeRepoFile(queuePath, `entries:
- kind: SvcTriaged
  group: svc.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: resolved
  resolution: accepted
- kind: SvcUntriaged
  group: svc.cnrm.cloud.google.com
  reason: untriaged-bulk-generation
  status: resolved
  resolution: edited
  note: triaged in this PR
`)

	if got := checkCRDFileBackwardCompatibility(repo, "HEAD", untriagedCRDPath); got != "" {
		t.Errorf("untriaged Kind being triaged should allow breaking changes, got: %s", got)
	}
	if got := checkCRDFileBackwardCompatibility(repo, "HEAD", newCRDPath); got != "" {
		t.Errorf("brand-new CRD should be allowed, got: %s", got)
	}
	if got := checkCRDFileBackwardCompatibility(repo, "HEAD", triagedCRDPath); !strings.Contains(got, "field removed: spec.sharedField") {
		t.Errorf("already-triaged Kind should reject breaking changes, got: %q", got)
	}
	if got := checkCRDFileBackwardCompatibility(repo, "HEAD", preexistingCRDPath); !strings.Contains(got, "field removed: spec.sharedField") {
		t.Errorf("pre-existing Kind should reject breaking changes, got: %q", got)
	}
}

