// Copyright 2022 Google LLC
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
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var (
	// This manager is only used to get the rest.Config from the test framework.
	unusedManager manager.Manager
)

func TestMain(m *testing.M) {
	testmain.TestMainForIntegrationTests(m, &unusedManager)
}

func TestAllInSeries(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	projectId := testgcp.GetDefaultProjectID(t)

	ctx := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	mgr, err := kccmanager.New(unusedManager.GetConfig(), kccmanager.Config{})
	if err != nil {
		t.Fatalf("error creating new manager: %v", err)
	}
	// Register the deletion defender controller.
	if err := registration.Add(mgr, nil, nil, nil, nil, registration.RegisterDeletionDefenderController); err != nil {
		t.Fatalf("error adding registration controller for deletion defender controllers: %v", err)
	}
	// Start the manager, Start(...) is a blocking operation so it needs to be done asynchronously.
	errors := make(chan error)
	go func() {
		errors <- mgr.Start(ctx)
	}()

	t.Run("samples", func(t *testing.T) {
		samples := create.LoadSamples(t)

		for _, s := range samples {
			s := s
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			t.Run(s.Name, func(t *testing.T) {
				cleanupResources := true

				h := create.NewHarness(t, ctx, mgr)
				create.SetupNamespacesAndApplyDefaults(h, []create.Sample{s})
				create.RunCreateDeleteTest(h, s.Resources, cleanupResources)
			})
		}
	})

	t.Run("fixtures", func(t *testing.T) {
		fixtures := resourcefixture.Load(t)
		for _, fixture := range fixtures {
			fixture := fixture
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			testID := testvariable.NewUniqueId()

			s := create.Sample{
				Name: fixture.Name,
			}

			initialUnstruct := bytesToUnstructured(t, fixture.Create, testID, projectId)
			s.Resources = append(s.Resources, initialUnstruct)

			if fixture.Dependencies != nil {
				dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
				for _, dependBytes := range dependencyYamls {
					depUnstruct := bytesToUnstructured(t, dependBytes, testID, projectId)
					s.Resources = append(s.Resources, depUnstruct)
				}
			}

			t.Run(s.Name, func(t *testing.T) {
				cleanupResources := true

				h := create.NewHarness(t, ctx, mgr)
				create.SetupNamespacesAndApplyDefaults(h, []create.Sample{s})
				// TODO(b/259496928): Do resource update
				create.RunCreateDeleteTest(h, s.Resources, cleanupResources)
			})
		}
	})

	// Do a cleanup while we can still handle the error.
	t.Logf("shutting down manager")
	cancel()

	if err := <-errors; err != nil {
		t.Fatalf("error from mgr.Start: %v", err)
	}
}

func bytesToUnstructured(t *testing.T, bytes []byte, testID string, projectId string) *unstructured.Unstructured {
	t.Helper()
	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testID, projectId)
	return test.ToUnstructWithNamespace(t, updatedBytes, testID)
}
