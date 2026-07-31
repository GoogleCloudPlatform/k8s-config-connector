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
	"reflect"
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

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

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

// TestBrownfieldLabelsInDirect tests the behaviors of labels in migrated brownfield resources.
func TestBrownfieldLabelsInDirect(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	// 1. Return an error if the env var is asking for real GCP.
	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "real" {
		t.Fatal("TestBrownfieldLabelsInDirect only runs against mock GCP")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	subtestTimeout := 5 * time.Minute

	t.Run("fixtures", func(t *testing.T) {
		fixtures := resourcefixture.LoadLabels(t)
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

			testName := fixture.Name
			t.Run(testName, func(t *testing.T) {
				subCtx := addTestTimeout(ctx, t, subtestTimeout, fixture.TestKey)
				runBrownfieldLabelsScenario(subCtx, t, fixture)
			})
		}
	})
}

func runBrownfieldLabelsScenario(ctx context.Context, t *testing.T, fixture resourcefixture.ResourceFixture) {
	uniqueID := testvariable.NewUniqueID()

	// Load objects temporarily to get GVKs for CRD filter
	dummyProject := testgcp.GCPProject{
		ProjectID:     "test-labels",
		ProjectNumber: 123456789,
	}
	dummyPrimary := bytesToUnstructured(t, fixture.Create, uniqueID, dummyProject)
	var dummyDeps []*unstructured.Unstructured
	if fixture.Dependencies != nil {
		dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
		for _, dependBytes := range dependencyYamls {
			depUnstruct := bytesToUnstructured(t, dependBytes, uniqueID, dummyProject)
			dummyDeps = append(dummyDeps, depUnstruct)
		}
	}

	// Build CRD filter and construct harness
	keepCRDs := map[schema.GroupKind]bool{}
	keepCRDs[dummyPrimary.GroupVersionKind().GroupKind()] = true
	for _, dep := range dummyDeps {
		keepCRDs[dep.GroupVersionKind().GroupKind()] = true
	}
	harnessOptions := []create.HarnessOption{buildCRDFilter(keepCRDs)}
	h := create.NewHarness(ctx, t, harnessOptions...)
	project := h.Project

	// Skip if the target resource doesn't support mock GCP.
	primaryResource := bytesToUnstructured(t, fixture.Create, uniqueID, project)
	var dependencies []*unstructured.Unstructured
	if fixture.Dependencies != nil {
		dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
		for _, dependBytes := range dependencyYamls {
			depUnstruct := bytesToUnstructured(t, dependBytes, uniqueID, project)
			dependencies = append(dependencies, depUnstruct)
		}
	}

	// Call MaybeSkip
	create.MaybeSkip(t, fixture.TestKey, append(dependencies, primaryResource))

	// Find if the resource supports direct controller and what old controller it has
	config, err := resourceconfig.LoadConfig().GetControllersForGVK(fixture.GVK)
	if err != nil {
		t.Fatalf("error getting controllers for GVK %v: %v", fixture.GVK, err)
	}

	var oldController k8scontrollertype.ReconcilerType
	for _, c := range config.SupportedControllers {
		if c == k8scontrollertype.ReconcilerTypeTerraform || c == k8scontrollertype.ReconcilerTypeDCL {
			oldController = c
			break
		}
	}
	if oldController == "" {
		t.Fatalf("no legacy controller (TF or DCL) found for GVK %s", fixture.GVK)
	}

	supportsDirect := false
	for _, c := range config.SupportedControllers {
		if c == k8scontrollertype.ReconcilerTypeDirect {
			supportsDirect = true
			break
		}
	}

	// Setup namespaces
	create.SetupNamespacesAndApplyDefaults(h, append(dependencies, primaryResource), project)

	// Hack: set project-id because mockkubeapiserver does not support webhooks
	for _, u := range append(dependencies, primaryResource) {
		ensureProjectIDAnnotation(u, project.ProjectID)
	}

	// Helper to set override on ConfigConnectorContext
	setCCCControllerOverride := func(reconciler k8scontrollertype.ReconcilerType) {
		// First delete any existing ConfigConnectorContext
		existingCCC := &opcorev1beta1.ConfigConnectorContext{}
		existingCCC.Name = "configconnectorcontext.core.cnrm.cloud.google.com"
		existingCCC.Namespace = primaryResource.GetNamespace()
		_ = h.GetClient().Delete(ctx, existingCCC)

		ccc := &opcorev1beta1.ConfigConnectorContext{}
		ccc.Name = "configconnectorcontext.core.cnrm.cloud.google.com"
		ccc.Namespace = primaryResource.GetNamespace()

		primaryGK := primaryResource.GroupVersionKind().GroupKind()
		controllerOverrides := map[string]k8scontrollertype.ReconcilerType{
			fmt.Sprintf("%s.%s", primaryGK.Kind, primaryGK.Group): reconciler,
		}
		ccc.Spec.Experiments = &opcorev1beta1.Experiments{
			ControllerOverrides: controllerOverrides,
		}
		if err := h.GetClient().Create(ctx, ccc); err != nil {
			t.Fatalf("FAIL: error creating CCC with reconciler %s: %v", reconciler, err)
		}
	}

	// Create ConfigConnector
	cc := &opcorev1beta1.ConfigConnector{}
	cc.Name = "configconnector.core.cnrm.cloud.google.com"
	cc.Spec.Mode = "namespaced"
	if err := h.GetClient().Create(ctx, cc); err != nil {
		t.Logf("ConfigConnector might already exist or failed: %v", err)
	}

	// Create the ConfigConnectorContext first so that dependencies can reconcile.
	setCCCControllerOverride(oldController)

	// Create the dependencies as needed.
	for _, u := range dependencies {
		t.Logf("creating dependency resource GVK: %s, name: %s", u.GroupVersionKind(), u.GetName())
		if err := h.GetClient().Patch(ctx, u, client.Apply, client.FieldOwner("kcc-tests")); err != nil {
			t.Fatalf("error creating dependency resource: %v", err)
		}
	}
	if len(dependencies) > 0 {
		create.WaitForReady(h, create.DefaultWaitForReadyTimeout, dependencies...)
	}

	// Run Phase 1 with legacy controller
	oldCreateLogStr, oldUpdateLogStr, oldTouchLogStr := runLabelScenarioPhase(ctx, t, h, project, uniqueID, fixture, oldController, "_old_controller", setCCCControllerOverride)

	if !supportsDirect {
		t.Logf("GVK %s does not support direct controller, skipping direct controller phase", fixture.GVK)
		return
	}

	// Run Phase 2 with direct controller
	newCreateLogStr, newUpdateLogStr, newTouchLogStr := runLabelScenarioPhase(ctx, t, h, project, uniqueID, fixture, k8scontrollertype.ReconcilerTypeDirect, "", setCCCControllerOverride)

	// Clean up dependencies
	if len(dependencies) > 0 {
		t.Logf("Deleting dependency resources...")
		optDeleteDeps := create.CreateDeleteTestOptions{
			Create:           dependencies,
			CleanupResources: true,
		}
		create.DeleteResources(h, optDeleteDeps)
	}

	// Compare _http_[operation]_old_controller.log and _http_[operation].log files
	ops := []struct {
		name string
		old  string
		new  string
	}{
		{name: "create", old: oldCreateLogStr, new: newCreateLogStr},
		{name: "update", old: oldUpdateLogStr, new: newUpdateLogStr},
		{name: "touch", old: oldTouchLogStr, new: newTouchLogStr},
	}

	for _, op := range ops {
		oldEvents := test.ParseHTTPLog(op.old)
		newEvents := test.ParseHTTPLog(op.new)

		var oldLabels []map[string]string
		for _, ev := range oldEvents {
			if ev.Method == "GET" || ev.Method == "DELETE" {
				continue
			}
			oldLabels = append(oldLabels, extractLabelsFromJSON(ev.RequestBody)...)
		}

		var newLabels []map[string]string
		for _, ev := range newEvents {
			if ev.Method == "GET" || ev.Method == "DELETE" {
				continue
			}
			newLabels = append(newLabels, extractLabelsFromJSON(ev.RequestBody)...)
		}

		if err := compareLabels(op.name, oldLabels, newLabels); err != nil {
			t.Errorf("Labels discrepancy found in %s: %v", fixture.Name, err)
		}
	}
}

func ensureProjectIDAnnotation(u *unstructured.Unstructured, projectID string) {
	annotations := u.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	if annotations["cnrm.cloud.google.com/project-id"] == "" {
		annotations["cnrm.cloud.google.com/project-id"] = projectID
		u.SetAnnotations(annotations)
	}
}

func runLabelScenarioPhase(ctx context.Context, t *testing.T, h *create.Harness, project testgcp.GCPProject, uniqueID string, fixture resourcefixture.ResourceFixture, reconciler k8scontrollertype.ReconcilerType, suffix string, setCCCControllerOverride func(k8scontrollertype.ReconcilerType)) (createLogStr, updateLogStr, touchLogStr string) {
	setCCCControllerOverride(reconciler)

	// Create primary resource
	primaryResource := bytesToUnstructured(t, fixture.Create, uniqueID, project)
	create.SetupNamespacesAndApplyDefaults(h, []*unstructured.Unstructured{primaryResource}, project)
	ensureProjectIDAnnotation(primaryResource, project.ProjectID)

	// Apply create.yaml
	h.Events.Clear()
	t.Logf("Creating primary resource GVK: %s, name: %s via %s controller", primaryResource.GroupVersionKind(), primaryResource.GetName(), reconciler)
	if err := h.GetClient().Patch(ctx, primaryResource, client.Apply, client.FieldOwner("kcc-tests")); err != nil {
		t.Fatalf("error applying primary resource: %v", err)
	}
	create.WaitForReady(h, create.DefaultWaitForReadyTimeout, primaryResource)

	eventsCreate := h.Events.GetHTTPEvents()
	createPath := filepath.Join(fixture.AbsoluteSourceDir, fmt.Sprintf("_http_create%s.log", suffix))
	createLogStr = normalizeAndWriteLog(t, h, project, uniqueID, eventsCreate, createPath)

	// Apply update.yaml
	if fixture.Update != nil {
		h.Events.Clear()
		updateResource := bytesToUnstructured(t, fixture.Update, uniqueID, project)
		create.SetupNamespacesAndApplyDefaults(h, []*unstructured.Unstructured{updateResource}, project)
		ensureProjectIDAnnotation(updateResource, project.ProjectID)

		t.Logf("Updating primary resource GVK: %s, name: %s via %s controller", updateResource.GroupVersionKind(), updateResource.GetName(), reconciler)
		if err := h.GetClient().Patch(ctx, updateResource, client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
			t.Fatalf("error applying update to primary resource: %v", err)
		}
		create.WaitForReady(h, create.DefaultWaitForReadyTimeout, updateResource)

		eventsUpdate := h.Events.GetHTTPEvents()
		updatePath := filepath.Join(fixture.AbsoluteSourceDir, fmt.Sprintf("_http_update%s.log", suffix))
		updateLogStr = normalizeAndWriteLog(t, h, project, uniqueID, eventsUpdate, updatePath)
	} else {
		updatePath := filepath.Join(fixture.AbsoluteSourceDir, fmt.Sprintf("_http_update%s.log", suffix))
		if err := os.WriteFile(updatePath, []byte(""), 0644); err != nil {
			t.Fatalf("error writing empty update log: %v", err)
		}
	}

	// Do a touch
	h.Events.Clear()
	preTouchRV := getResourceVersion(h, primaryResource)

	uTouch := &unstructured.Unstructured{}
	uTouch.SetGroupVersionKind(primaryResource.GroupVersionKind())
	uTouch.SetName(primaryResource.GetName())
	uTouch.SetNamespace(primaryResource.GetNamespace())

	existingTouch := readObject(h, primaryResource.GroupVersionKind(), primaryResource.GetNamespace(), primaryResource.GetName())
	annotationsTouch := existingTouch.GetAnnotations()
	if annotationsTouch == nil {
		annotationsTouch = make(map[string]string)
	}
	annotationsTouch["test.cnrm.cloud.google.com/reconcile-cookie"] = fmt.Sprintf("re-reconcile-%s-v1", reconciler)
	uTouch.SetAnnotations(annotationsTouch)

	t.Logf("Touching primary resource GVK: %s, name: %s via %s controller", uTouch.GroupVersionKind(), uTouch.GetName(), reconciler)
	if err := h.GetClient().Patch(ctx, uTouch, client.Apply, client.FieldOwner("kcc-test-touch"), client.ForceOwnership); err != nil {
		t.Fatalf("error applying touch patch to primary resource: %v", err)
	}

	waitForReconciliationAfterPatch(h, primaryResource, preTouchRV)

	eventsTouch := h.Events.GetHTTPEvents()
	touchPath := filepath.Join(fixture.AbsoluteSourceDir, fmt.Sprintf("_http_touch%s.log", suffix))
	touchLogStr = normalizeAndWriteLog(t, h, project, uniqueID, eventsTouch, touchPath)

	// Delete primary resource
	t.Logf("Deleting primary resource GVK: %s, name: %s", primaryResource.GroupVersionKind(), primaryResource.GetName())
	optDelete := create.CreateDeleteTestOptions{
		Create:           []*unstructured.Unstructured{primaryResource},
		CleanupResources: true,
	}
	create.DeleteResources(h, optDelete)

	return createLogStr, updateLogStr, touchLogStr
}

func normalizeAndWriteLog(t *testing.T, h *create.Harness, project testgcp.GCPProject, uniqueID string, events []*test.LogEntry, filePath string) string {
	got, normalizers := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(events))
	for _, n := range normalizers {
		got = n(got)
	}
	if err := os.WriteFile(filePath, []byte(got), 0644); err != nil {
		t.Fatalf("error writing HTTP log to %s: %v", filePath, err)
	}
	return got
}

func extractLabelsFromJSON(jsonStr string) []map[string]string {
	if jsonStr == "" {
		return nil
	}
	var val any
	if err := json.Unmarshal([]byte(jsonStr), &val); err != nil {
		return nil
	}
	var results []map[string]string
	findLabels(val, &results)
	return results
}

func findLabels(val any, results *[]map[string]string) {
	switch v := val.(type) {
	case map[string]any:
		for k, val := range v {
			if k == "labels" || k == "userLabels" {
				if m, ok := val.(map[string]any); ok {
					labelsMap := make(map[string]string)
					for lk, lv := range m {
						if sv, ok := lv.(string); ok {
							labelsMap[lk] = sv
						}
					}
					*results = append(*results, labelsMap)
				}
			} else {
				findLabels(val, results)
			}
		}
	case []any:
		for _, item := range v {
			findLabels(item, results)
		}
	}
}

func compareLabels(op string, oldLabels, newLabels []map[string]string) error {
	normalize := func(list []map[string]string) []map[string]string {
		var res []map[string]string
		for _, m := range list {
			if len(m) == 0 {
				continue
			}
			res = append(res, m)
		}
		return res
	}

	normOld := normalize(oldLabels)
	normNew := normalize(newLabels)

	if len(normOld) != len(normNew) {
		return fmt.Errorf("discrepancy in %s: old controller sent %d label-writing requests, but new controller sent %d", op, len(normOld), len(normNew))
	}

	for i := range normOld {
		o := normOld[i]
		n := normNew[i]
		if !reflect.DeepEqual(o, n) {
			return fmt.Errorf("discrepancy in %s: write request %d labels differ. Old: %v, New: %v", op, i, o, n)
		}
	}
	return nil
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
				Old: jsonFriendlyValue(f.Old),
				New: jsonFriendlyValue(f.New),
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

func jsonFriendlyValue(val any) any {
	if val == nil {
		return nil
	}

	// If it's a protoreflect.Value, extract its interface
	if pv, ok := val.(protoreflect.Value); ok {
		if !pv.IsValid() {
			return nil
		}
		val = pv.Interface()
	}

	switch v := val.(type) {
	case protoreflect.Message:
		return jsonFriendlyMessage(v)
	case protoreflect.List:
		return jsonFriendlyList(v)
	case protoreflect.Map:
		return jsonFriendlyMap(v)
	case proto.Message:
		return jsonFriendlyMessage(v.ProtoReflect())
	case []byte:
		return string(v)
	default:
		return v
	}
}

func jsonFriendlyMessage(m protoreflect.Message) any {
	if m == nil || !m.IsValid() {
		return nil
	}
	res := make(map[string]any)
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		res[string(fd.Name())] = jsonFriendlyValue(v)
		return true
	})
	return res
}

func jsonFriendlyList(l protoreflect.List) any {
	if l == nil {
		return nil
	}
	var res []any
	for i := 0; i < l.Len(); i++ {
		res = append(res, jsonFriendlyValue(l.Get(i)))
	}
	return res
}

func jsonFriendlyMap(m protoreflect.Map) any {
	if m == nil {
		return nil
	}
	res := make(map[string]any)
	m.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		res[k.String()] = jsonFriendlyValue(v)
		return true
	})
	return res
}
