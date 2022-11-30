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
	"strconv"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func TestAllInSeries(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	var project testgcp.GCPProject
	if os.Getenv("E2E_GCP_TARGET") == "mock" {
		projectNumber := time.Now().Unix()
		project = testgcp.GCPProject{
			ProjectID:     "mock-project-" + strconv.FormatInt(projectNumber, 10),
			ProjectNumber: projectNumber,
		}
	} else {
		project = testgcp.GetDefaultProject(t)
	}

	ctx := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	t.Run("samples", func(t *testing.T) {
		samples := create.LoadSamples(t, project)

		for _, s := range samples {
			s := s
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			t.Run(s.Name, func(t *testing.T) {
				create.MaybeSkip(t, s.Name, s.Resources)

				h := create.NewHarness(t, ctx)

				cleanupResources := true

				create.SetupNamespacesAndApplyDefaults(h, []create.Sample{s}, project)

				// Hack: set project-id because mockkubeapiserver does not support webhooks
				for _, u := range s.Resources {
					annotations := u.GetAnnotations()
					if annotations == nil {
						annotations = make(map[string]string)
					}
					annotations["cnrm.cloud.google.com/project-id"] = project.ProjectID
					u.SetAnnotations(annotations)
				}

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

			initialUnstruct := bytesToUnstructured(t, fixture.Create, testID, project)
			s.Resources = append(s.Resources, initialUnstruct)

			if fixture.Dependencies != nil {
				dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
				for _, dependBytes := range dependencyYamls {
					depUnstruct := bytesToUnstructured(t, dependBytes, testID, project)
					s.Resources = append(s.Resources, depUnstruct)
				}
			}

			t.Run(s.Name, func(t *testing.T) {
				create.MaybeSkip(t, s.Name, s.Resources)

				h := create.NewHarness(t, ctx)

				cleanupResources := true

				create.SetupNamespacesAndApplyDefaults(h, []create.Sample{s}, project)
				create.RunCreateDeleteTest(h, s.Resources, cleanupResources)
			})
		}
	})

	// Do a cleanup while we can still handle the error.
	t.Logf("shutting down manager")
	cancel()
}

func bytesToUnstructured(t *testing.T, bytes []byte, testID string, project testgcp.GCPProject) *unstructured.Unstructured {
	t.Helper()
	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testID, project)
	return test.ToUnstructWithNamespace(t, updatedBytes, testID)
}
