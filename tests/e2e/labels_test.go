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
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	opcorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	k8scontrollertype "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TestBrownfieldLabelsInDirect tests direct controller support of labels on brownfield resources
// against simulated mock GCP services.
func TestBrownfieldLabelsInDirect(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	if os.Getenv("E2E_GCP_TARGET") == "real" {
		t.Fatalf("TestBrownfieldLabelsInDirect must only run against mock GCP")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	subtestTimeout := 5 * time.Minute

	t.Run("labels-fixtures", func(t *testing.T) {
		pathFilter := func(path string) bool {
			return strings.Contains(path, "testdata/labels")
		}

		fixtures := resourcefixture.LoadWithPathFilter(t, pathFilter, nil, nil)
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

			if oldController == "" {
				// No legacy controller supported for labels comparison
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
				runBrownfieldLabelsScenario(subCtx, t, fixture, oldController, hasDirect)
			})
		}
	})
}

func runBrownfieldLabelsScenario(ctx context.Context, t *testing.T, fixture resourcefixture.ResourceFixture, oldController k8scontrollertype.ReconcilerType, hasDirect bool) {
	uniqueID := testvariable.NewUniqueID()

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

	primaryResourceForSkip, _ := loadFixture(testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789}, uniqueID)
	create.MaybeSkip(t, fixture.TestKey, []*unstructured.Unstructured{primaryResourceForSkip})

	runController := func(controller k8scontrollertype.ReconcilerType, isOld bool) []string {
		harnessCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		keepCRDs := map[schema.GroupKind]bool{}
		_, dummyOpt := loadFixture(testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789}, uniqueID)
		for _, obj := range dummyOpt.Create {
			keepCRDs[obj.GroupVersionKind().GroupKind()] = true
		}
		harnessOptions := []create.HarnessOption{buildCRDFilter(keepCRDs)}

		h := create.NewHarness(harnessCtx, t, harnessOptions...)
		project := h.Project

		primaryResource, opt := loadFixture(project, uniqueID)

		// Setup namespaces
		create.SetupNamespacesAndApplyDefaults(h, opt.Create, project)

		// Create ConfigConnector
		cc := &opcorev1beta1.ConfigConnector{}
		cc.Name = "configconnector.core.cnrm.cloud.google.com"
		cc.Spec.Mode = "namespaced"
		if err := h.GetClient().Create(harnessCtx, cc); err != nil {
			t.Fatalf("FAIL: error creating CC: %v", err)
		}

		// Create ConfigConnectorContext with controllerOverride to force controller
		ccc := &opcorev1beta1.ConfigConnectorContext{}
		ccc.Name = "configconnectorcontext.core.cnrm.cloud.google.com"
		ccc.Namespace = primaryResource.GetNamespace()

		primaryGK := primaryResource.GroupVersionKind().GroupKind()
		controllerOverrides := map[string]k8scontrollertype.ReconcilerType{
			fmt.Sprintf("%s.%s", primaryGK.Kind, primaryGK.Group): controller,
		}
		ccc.Spec.Experiments = &opcorev1beta1.Experiments{
			ControllerOverrides: controllerOverrides,
		}
		if err := h.GetClient().Create(harnessCtx, ccc); err != nil {
			t.Fatalf("FAIL: error creating CCC: %v", err)
		}

		// 4. Apply create.yaml
		t.Logf("Creating resource using %v...", controller)
		for _, u := range opt.Create {
			t.Log("creating object", "GVK", u.GroupVersionKind().String(), "name", u.GetName())
			if err := h.GetClient().Patch(harnessCtx, u, client.Apply, client.FieldOwner("kcc-tests")); err != nil {
				t.Fatalf("error creating resource: %v", err)
			}
		}
		create.WaitForReady(h, create.DefaultWaitForReadyTimeout, opt.Create...)

		eventsCreate := h.Events.GetHTTPEvents()
		createLog, _ := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(eventsCreate))

		suffix := ""
		if isOld {
			suffix = "_old_controller"
		}
		createLogPath := filepath.Join(fixture.AbsoluteSourceDir, fmt.Sprintf("_http_create%s.log", suffix))
		if err := os.WriteFile(createLogPath, []byte(createLog), 0644); err != nil {
			t.Fatalf("failed to write %s: %v", createLogPath, err)
		}

		// 5. Apply update.yaml
		t.Logf("Updating resource using %v...", controller)
		updateBytes := fixture.Update
		if updateBytes == nil {
			t.Fatalf("FAIL: labels testcase %s must have update.yaml", fixture.Name)
		}
		updateResource := bytesToUnstructured(t, updateBytes, uniqueID, project)

		prePatchRVUpdate := getResourceVersion(h, primaryResource)
		if err := h.GetClient().Patch(harnessCtx, updateResource, client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
			t.Fatalf("error updating resource: %v", err)
		}
		create.WaitForReady(h, create.DefaultWaitForReadyTimeout, updateResource)
		waitForReconciliationAfterPatch(h, primaryResource, prePatchRVUpdate)

		eventsUpdate := h.Events.GetHTTPEvents()[len(eventsCreate):]
		updateLog, _ := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(eventsUpdate))
		updateLogPath := filepath.Join(fixture.AbsoluteSourceDir, fmt.Sprintf("_http_update%s.log", suffix))
		if err := os.WriteFile(updateLogPath, []byte(updateLog), 0644); err != nil {
			t.Fatalf("failed to write %s: %v", updateLogPath, err)
		}

		// 6. Do a touch to trigger re-reconciliation with no change
		t.Logf("Re-reconciling resource using %v (no-op update)...", controller)
		prePatchRVTouch := getResourceVersion(h, primaryResource)

		uTouch := &unstructured.Unstructured{}
		uTouch.SetGroupVersionKind(primaryResource.GroupVersionKind())
		uTouch.SetName(primaryResource.GetName())
		uTouch.SetNamespace(primaryResource.GetNamespace())

		existingTouch := readObject(h, primaryResource.GroupVersionKind(), primaryResource.GetNamespace(), primaryResource.GetName())
		annotationsTouch := existingTouch.GetAnnotations()
		if annotationsTouch == nil {
			annotationsTouch = make(map[string]string)
		}
		annotationsTouch["test.cnrm.cloud.google.com/reconcile-cookie"] = "re-reconcile-v1"
		uTouch.SetAnnotations(annotationsTouch)

		if err := h.GetClient().Patch(harnessCtx, uTouch, client.Apply, client.FieldOwner("kcc-test-touch"), client.ForceOwnership); err != nil {
			t.Fatalf("error applying no-op annotation: %v", err)
		}

		waitForReconciliationAfterPatch(h, primaryResource, prePatchRVTouch)

		eventsTouch := h.Events.GetHTTPEvents()[len(eventsCreate)+len(eventsUpdate):]
		touchLog, _ := LegacyNormalize(t, h, project, uniqueID, test.LogEntries(eventsTouch))
		touchLogPath := filepath.Join(fixture.AbsoluteSourceDir, fmt.Sprintf("_http_touch%s.log", suffix))
		if err := os.WriteFile(touchLogPath, []byte(touchLog), 0644); err != nil {
			t.Fatalf("failed to write %s: %v", touchLogPath, err)
		}

		// 7. Delete the target resource
		t.Logf("Deleting target resource %s using %v...", primaryResource.GetName(), controller)
		opt.CleanupResources = true
		create.DeleteResources(h, opt)

		return []string{createLogPath, updateLogPath, touchLogPath}
	}

	oldFiles := runController(oldController, true)

	if hasDirect {
		newFiles := runController(k8scontrollertype.ReconcilerTypeDirect, false)

		for i := 0; i < len(oldFiles); i++ {
			compareLogFiles(t, oldFiles[i], newFiles[i])
		}
	}
}

func compareLogFiles(t *testing.T, oldFile, newFile string) {
	oldBytes, err := os.ReadFile(oldFile)
	if err != nil {
		t.Fatalf("failed to read %s: %v", oldFile, err)
	}
	newBytes, err := os.ReadFile(newFile)
	if err != nil {
		t.Fatalf("failed to read %s: %v", newFile, err)
	}
	oldStr := string(oldBytes)
	newStr := string(newBytes)
	if oldStr != newStr {
		t.Errorf("DISCREPANCY: labels behavior differs between old and new controllers. Compare %s and %s.", oldFile, newFile)
	}
}
