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

package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	opcorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	k8scontrollertype "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// TestMigrationToDirect tests the transition/migration of a resource from being
// managed by the legacy (TF or DCL) controller to the Direct controller.
// It first provisions the resource using the legacy controller, then applies the
// `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to force the Direct
// controller to take over. It verifies that the takeover is smooth and does not
// trigger any unexpected updates (i.e., a no-op reconciliation).
func TestMigrationToDirect(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	subtestTimeout := time.Hour
	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "mock" {
		subtestTimeout = 5 * time.Minute
	}

	t.Run("fixtures", func(t *testing.T) {
		// Load all fixtures using the same filters as TestAllInSeries
		lightFilter := func(name string, testType resourcefixture.TestType) bool {
			return !strings.Contains(name, "iam-bigqueryconnectionconnectionref") &&
				!strings.Contains(name, "iam-logsinkref") &&
				!strings.Contains(name, "iam-serviceaccountref") &&
				!strings.Contains(name, "iam-serviceidentityref") &&
				!strings.Contains(name, "iam-sqlinstanceref")
		}
		pathFilter := func(path string) bool {
			return !strings.Contains(path, "testdata/iam/iampartialpolicy")
		}

		fixtures := resourcefixture.LoadWithPathFilter(t, pathFilter, lightFilter, nil)
		for _, fixture := range fixtures {
			fixture := fixture
			group := fixture.GVK.Group

			if s := os.Getenv("SKIP_TEST_APIGROUP"); s != "" {
				skippedGroups := strings.Split(s, ",")
				if slices.Contains(skippedGroups, group) {
					continue
				}
			}
			if s := os.Getenv("ONLY_TEST_APIGROUPS"); s != "" {
				groups := strings.Split(s, ",")
				if !slices.Contains(groups, group) {
					continue
				}
			}

			// We only want to test migration if the default controller is TF/DCL, but Direct is supported.
			config, err := resourceconfig.LoadConfig().GetControllersForGVK(fixture.GVK)
			if err != nil {
				t.Logf("skipping GVK %v: %v", fixture.GVK, err)
				continue
			}

			hasDirect := false
			var oldController k8scontrollertype.ReconcilerType
			for _, c := range config.SupportedControllers {
				if c == k8scontrollertype.ReconcilerTypeDirect {
					hasDirect = true
				} else {
					oldController = c
				}
			}
			if !hasDirect || oldController == "" {
				// Skip because we can't migrate from TF/DCL to Direct (or Direct is already the only/default controller)
				continue
			}

			// Skip fixtures that explicitly specify the direct reconciler annotation in their create.yaml,
			// as they are already direct-first tests, not migration tests.
			if strings.Contains(string(fixture.Create), "alpha.cnrm.cloud.google.com/reconciler: direct") {
				continue
			}

			testName := fixture.Name
			if os.Getenv("USE_FULL_TEST_NAMES") == "true" {
				testName = "pkg/test/resourcefixture/testdata/" + fixture.TestKey
			}

			t.Run(testName, func(t *testing.T) {
				if os.Getenv("SKIP_ALL") != "" {
					t.Skip("SKIP_ALL is set")
				}

				subCtx := addTestTimeout(ctx, t, subtestTimeout, fixture.TestKey)
				runMigrationScenario(subCtx, t, fixture, oldController)
			})
		}
	})
}

func runMigrationScenario(ctx context.Context, t *testing.T, fixture resourcefixture.ResourceFixture, oldController k8scontrollertype.ReconcilerType) {
	uniqueID := testvariable.NewUniqueID()

	// Load fixture data
	loadFixture := func(project testgcp.GCPProject, uniqueID string) (*unstructured.Unstructured, create.CreateDeleteTestOptions) {
		primaryResource := bytesToUnstructured(t, fixture.Create, uniqueID, project)
		opt := create.CreateDeleteTestOptions{CleanupResources: false}

		if fixture.Dependencies != nil {
			dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
			for _, dependBytes := range dependencyYamls {
				depUnstruct := bytesToUnstructured(t, dependBytes, uniqueID, project)
				opt.Create = append(opt.Create, depUnstruct)
			}
		}
		opt.Create = append(opt.Create, primaryResource)
		opt.PrimaryResource = primaryResource
		return primaryResource, opt
	}

	// Build harness options (filter CRDs)
	_, dummyOpt := loadFixture(testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789}, uniqueID)
	keepCRDs := map[schema.GroupKind]bool{}
	for _, obj := range dummyOpt.Create {
		keepCRDs[obj.GroupVersionKind().GroupKind()] = true
	}
	harnessOptions := []create.HarnessOption{buildCRDFilter(keepCRDs)}

	// Create custom structured reporting listener to capture diffs
	diffListener := &migrationDiffListener{}
	ctx = structuredreporting.ContextWithListener(ctx, diffListener)

	// Create harness
	h := create.NewHarness(ctx, t, harnessOptions...)
	project := h.Project

	primaryResource, opt := loadFixture(project, uniqueID)

	// Setup namespaces
	create.SetupNamespacesAndApplyDefaults(h, opt.Create, project)

	// Hack: set project-id because mockkubeapiserver does not support webhooks
	for _, u := range opt.Create {
		annotations := u.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}
		if annotations["cnrm.cloud.google.com/project-id"] == "" {
			annotations["cnrm.cloud.google.com/project-id"] = project.ProjectID
			u.SetAnnotations(annotations)
		}
	}

	// Create ConfigConnector
	cc := &opcorev1beta1.ConfigConnector{}
	cc.Name = "configconnector.core.cnrm.cloud.google.com"
	cc.Spec.Mode = "namespaced"
	if err := h.GetClient().Create(ctx, cc); err != nil {
		t.Fatalf("FAIL: error creating CC: %v", err)
	}

	// Create ConfigConnectorContext with controllerOverride to force old controller
	ccc := &opcorev1beta1.ConfigConnectorContext{}
	ccc.Name = "configconnectorcontext.core.cnrm.cloud.google.com"
	ccc.Namespace = primaryResource.GetNamespace()

	primaryGK := primaryResource.GroupVersionKind().GroupKind()
	controllerOverrides := map[string]k8scontrollertype.ReconcilerType{
		fmt.Sprintf("%s.%s", primaryGK.Kind, primaryGK.Group): oldController,
	}
	ccc.Spec.Experiments = &opcorev1beta1.Experiments{
		ControllerOverrides: controllerOverrides,
	}
	if err := h.GetClient().Create(ctx, ccc); err != nil {
		t.Fatalf("FAIL: error creating CCC: %v", err)
	}

	t.Logf("Phase 1: Creating resource using %v...", oldController)
	// Create resources (dependencies + primary)
	for _, u := range opt.Create {
		t.Log("creating object", "GVK", u.GroupVersionKind().String(), "name", u.GetName())
		if err := h.GetClient().Patch(ctx, u, client.Apply, client.FieldOwner("kcc-tests")); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
	}
	// Wait for them to be ready
	create.WaitForReady(h, create.DefaultWaitForReadyTimeout, opt.Create...)

	// Record HTTP log for Phase 1
	eventsPhase1 := h.Events.GetHTTPEvents()
	if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		got, normalizers := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(eventsPhase1))
		h.CompareGoldenFile(filepath.Join(fixture.AbsoluteSourceDir, "_http_migration_phase1_legacy_create.log"), got, normalizers...)
	}

	t.Logf("Phase 2: Re-reconciling resource using %v (no-op update)...", oldController)
	// Get pre-patch resource version to wait for reconciliation
	prePatchRVPhase2 := getResourceVersion(h, primaryResource)

	// Update primary resource with a no-op annotation to trigger a re-reconciliation with the legacy controller
	uPhase2 := &unstructured.Unstructured{}
	uPhase2.SetGroupVersionKind(primaryResource.GroupVersionKind())
	uPhase2.SetName(primaryResource.GetName())
	uPhase2.SetNamespace(primaryResource.GetNamespace())

	existingPhase2 := readObject(h, primaryResource.GroupVersionKind(), primaryResource.GetNamespace(), primaryResource.GetName())
	annotationsPhase2 := existingPhase2.GetAnnotations()
	if annotationsPhase2 == nil {
		annotationsPhase2 = make(map[string]string)
	}
	annotationsPhase2["test.cnrm.cloud.google.com/reconcile-cookie"] = "re-reconcile-legacy-v1"
	uPhase2.SetAnnotations(annotationsPhase2)

	t.Logf("Applying no-op annotation to %s/%s to trigger %v re-reconciliation", uPhase2.GetNamespace(), uPhase2.GetName(), oldController)
	if err := h.GetClient().Patch(ctx, uPhase2, client.Apply, client.FieldOwner("kcc-test-migration-touch"), client.ForceOwnership); err != nil {
		t.Fatalf("error applying no-op annotation: %v", err)
	}

	// Wait for legacy controller to reconcile it
	waitForReconciliationAfterPatch(h, primaryResource, prePatchRVPhase2)

	// Record HTTP log for Phase 2
	eventsPhase2 := h.Events.GetHTTPEvents()[len(eventsPhase1):]
	if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		got, normalizers := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(eventsPhase2))
		h.CompareGoldenFile(filepath.Join(fixture.AbsoluteSourceDir, "_http_migration_phase2_legacy_re-reconciliation.log"), got, normalizers...)
	}

	t.Log("Phase 3: Migrating to Direct controller...")
	// Get pre-patch resource version to wait for reconciliation
	prePatchRVPhase3 := getResourceVersion(h, primaryResource)

	// Update primary resource with direct reconciler annotation
	uPhase3 := &unstructured.Unstructured{}
	uPhase3.SetGroupVersionKind(primaryResource.GroupVersionKind())
	uPhase3.SetName(primaryResource.GetName())
	uPhase3.SetNamespace(primaryResource.GetNamespace())

	// Get existing annotations
	existingPhase3 := readObject(h, primaryResource.GroupVersionKind(), primaryResource.GetNamespace(), primaryResource.GetName())
	annotationsPhase3 := existingPhase3.GetAnnotations()
	if annotationsPhase3 == nil {
		annotationsPhase3 = make(map[string]string)
	}
	annotationsPhase3["alpha.cnrm.cloud.google.com/reconciler"] = "direct"
	annotationsPhase3["test.cnrm.cloud.google.com/reconcile-cookie"] = "migration-v1"
	uPhase3.SetAnnotations(annotationsPhase3)

	t.Logf("Applying direct reconciler annotation to %s/%s", uPhase3.GetNamespace(), uPhase3.GetName())
	if err := h.GetClient().Patch(ctx, uPhase3, client.Apply, client.FieldOwner("kcc-test-migration-touch"), client.ForceOwnership); err != nil {
		t.Fatalf("error applying direct reconciler annotation: %v", err)
	}

	// Wait for direct controller to reconcile it
	waitForReconciliationAfterPatch(h, primaryResource, prePatchRVPhase3)

	// Verify HTTP events during Phase 3 (Direct take over)
	eventsPhase3 := h.Events.GetHTTPEvents()[len(eventsPhase1)+len(eventsPhase2):]

	// The direct controller should not perform any updates (no-op reconciliation)
	for _, event := range eventsPhase3 {
		isReadOnly := false
		switch event.Request.Method {
		case "GET":
			isReadOnly = true
		case "GRPC":
			if strings.Contains(event.Request.URL, "/Get") || strings.Contains(event.Request.URL, "/List") {
				isReadOnly = true
			}
		}
		if !isReadOnly {
			t.Errorf("FAIL: unexpected write request during migration reconciliation: %v %v", event.Request.Method, event.Request.URL)
		}
	}

	// Record HTTP log for Phase 3
	if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		got, normalizers := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(eventsPhase3))
		h.CompareGoldenFile(filepath.Join(fixture.AbsoluteSourceDir, "_http_migration_phase3_direct_takeover.log"), got, normalizers...)
	}

	t.Log("Phase 4: Re-reconciling resource using Direct controller (no-op update)...")
	// Get pre-patch resource version to wait for reconciliation
	prePatchRVPhase4 := getResourceVersion(h, primaryResource)

	// Update primary resource with a no-op annotation to trigger a re-reconciliation with the Direct controller
	uPhase4 := &unstructured.Unstructured{}
	uPhase4.SetGroupVersionKind(primaryResource.GroupVersionKind())
	uPhase4.SetName(primaryResource.GetName())
	uPhase4.SetNamespace(primaryResource.GetNamespace())

	existingPhase4 := readObject(h, primaryResource.GroupVersionKind(), primaryResource.GetNamespace(), primaryResource.GetName())
	annotationsPhase4 := existingPhase4.GetAnnotations()
	if annotationsPhase4 == nil {
		annotationsPhase4 = make(map[string]string)
	}
	annotationsPhase4["test.cnrm.cloud.google.com/reconcile-cookie"] = "re-reconcile-direct-v1"
	uPhase4.SetAnnotations(annotationsPhase4)

	t.Logf("Applying no-op annotation to %s/%s to trigger Direct re-reconciliation", uPhase4.GetNamespace(), uPhase4.GetName())
	if err := h.GetClient().Patch(ctx, uPhase4, client.Apply, client.FieldOwner("kcc-test-migration-touch"), client.ForceOwnership); err != nil {
		t.Fatalf("error applying no-op annotation: %v", err)
	}

	// Wait for Direct controller to re-reconcile it
	waitForReconciliationAfterPatch(h, primaryResource, prePatchRVPhase4)

	// Verify HTTP events during Phase 4 (Direct re-reconciliation)
	eventsPhase4 := h.Events.GetHTTPEvents()[len(eventsPhase1)+len(eventsPhase2)+len(eventsPhase3):]

	// The direct controller should not perform any updates (no-op reconciliation)
	for _, event := range eventsPhase4 {
		isReadOnly := false
		switch event.Request.Method {
		case "GET":
			isReadOnly = true
		case "GRPC":
			if strings.Contains(event.Request.URL, "/Get") || strings.Contains(event.Request.URL, "/List") {
				isReadOnly = true
			}
		}
		if !isReadOnly {
			t.Errorf("FAIL: unexpected write request during Direct re-reconciliation: %v %v", event.Request.Method, event.Request.URL)
		}
	}

	// Record HTTP log for Phase 4
	if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		got, normalizers := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(eventsPhase4))
		h.CompareGoldenFile(filepath.Join(fixture.AbsoluteSourceDir, "_http_migration_phase4_direct_re-reconciliation.log"), got, normalizers...)
	}

	// Record raw structured diffs
	if os.Getenv("GOLDEN_OBJECT_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		rawDiffsStr := formatDiffsRaw(t, diffListener)
		_, normalizers := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(h.Events.GetHTTPEvents()))
		h.CompareGoldenFile(filepath.Join(fixture.AbsoluteSourceDir, "_migration_diffs.json"), rawDiffsStr, normalizers...)
	}

	// Cleanup
	t.Log("Cleaning up resources...")
	opt.CleanupResources = true
	create.DeleteResources(h, opt)
}

func getResourceVersion(h *create.Harness, obj *unstructured.Unstructured) string {
	existing := readObject(h, obj.GroupVersionKind(), obj.GetNamespace(), obj.GetName())
	return existing.GetResourceVersion()
}

type migrationDiffListener struct {
	mu    sync.Mutex
	diffs []*structuredreporting.Diff
}

func (l *migrationDiffListener) OnDiff(ctx context.Context, diff *structuredreporting.Diff) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Clone the diff because the underlying object might be modified
	clone := &structuredreporting.Diff{
		Controller:  diff.Controller,
		IsNewObject: diff.IsNewObject,
	}
	if diff.Object != nil {
		clone.Object = diff.Object.DeepCopy()
	}
	for _, f := range diff.Fields {
		clone.Fields = append(clone.Fields, structuredreporting.DiffField{
			ID:                   f.ID,
			ProtoFieldDescriptor: f.ProtoFieldDescriptor,
			Old:                  f.Old,
			New:                  f.New,
		})
	}
	l.diffs = append(l.diffs, clone)
}

func (l *migrationDiffListener) OnError(ctx context.Context, err error, args ...any) {}
func (l *migrationDiffListener) OnReconcileStart(ctx context.Context, u *unstructured.Unstructured, t k8scontrollertype.ReconcilerType) {
}
func (l *migrationDiffListener) OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error, t k8scontrollertype.ReconcilerType) {
}

type rawDiffField struct {
	ID  string `json:"id"`
	Old any    `json:"old,omitempty"`
	New any    `json:"new,omitempty"`
}

type rawDiff struct {
	Controller  string         `json:"controller,omitempty"`
	IsNewObject bool           `json:"isNewObject"`
	Resource    string         `json:"resource"`
	Fields      []rawDiffField `json:"fields,omitempty"`
}

func formatDiffsRaw(t *testing.T, listener *migrationDiffListener) string {
	var rawDiffs []rawDiff
	for _, diff := range listener.diffs {
		rd := rawDiff{
			Controller:  string(diff.Controller),
			IsNewObject: diff.IsNewObject,
		}

		if diff.Object != nil {
			rd.Resource = fmt.Sprintf("%s/%s", diff.Object.GetKind(), diff.Object.GetName())
		}

		// Sort fields by ID to ensure deterministic output
		fields := append([]structuredreporting.DiffField{}, diff.Fields...)
		sort.Slice(fields, func(i, j int) bool {
			return fields[i].ID < fields[j].ID
		})

		for _, f := range fields {
			rd.Fields = append(rd.Fields, rawDiffField{
				ID:  f.ID,
				Old: f.Old,
				New: f.New,
			})
		}
		rawDiffs = append(rawDiffs, rd)
	}

	// Marshal to pretty JSON
	bytes, err := json.MarshalIndent(rawDiffs, "", "  ")
	if err != nil {
		t.Fatalf("FAIL: error marshaling diffs to JSON: %v", err)
	}
	return string(bytes) + "\n"
}
