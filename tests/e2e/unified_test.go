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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestAllInSeries(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	t.Run("samples", func(t *testing.T) {
		samples := create.ListAllSamples(t)

		for _, sampleKey := range samples {
			sampleKey := sampleKey
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			t.Run(sampleKey.FullName, func(t *testing.T) {
				// Quickly load the sample with a dummy project, just to see if we should skip it
				{
					dummySample := create.LoadSample(t, sampleKey, testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, sampleKey.FullName, dummySample.Resources)
				}

				h := create.NewHarness(t, ctx)
				project := h.Project
				s := create.LoadSample(t, sampleKey, project)

				create.SetupNamespacesAndApplyDefaults(h, s.Resources, project)

				// Hack: set project-id because mockkubeapiserver does not support webhooks
				for _, u := range s.Resources {
					annotations := u.GetAnnotations()
					if annotations == nil {
						annotations = make(map[string]string)
					}
					annotations["cnrm.cloud.google.com/project-id"] = project.ProjectID
					u.SetAnnotations(annotations)
				}

				create.RunCreateDeleteTest(h, create.CreateDeleteTestOptions{Create: s.Resources, CleanupResources: true})
			})
		}
	})

	t.Run("fixtures", func(t *testing.T) {
		fixtures := resourcefixture.Load(t)
		for _, fixture := range fixtures {
			fixture := fixture
			// TODO(b/259496928): Randomize the resource names for parallel execution when/if needed.

			t.Run(fixture.FullName, func(t *testing.T) {
				uniqueID := testvariable.NewUniqueId()

				loadFixture := func(project testgcp.GCPProject) (*unstructured.Unstructured, create.CreateDeleteTestOptions) {
					primaryResource := bytesToUnstructured(t, fixture.Create, uniqueID, project)

					opt := create.CreateDeleteTestOptions{CleanupResources: true}
					opt.Create = append(opt.Create, primaryResource)

					if fixture.Dependencies != nil {
						dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
						for _, dependBytes := range dependencyYamls {
							depUnstruct := bytesToUnstructured(t, dependBytes, uniqueID, project)
							opt.Create = append(opt.Create, depUnstruct)
						}
					}

					if fixture.Update != nil {
						u := bytesToUnstructured(t, fixture.Update, uniqueID, project)
						opt.Updates = append(opt.Updates, u)
					}
					return primaryResource, opt
				}

				// Quickly load the fixture with a dummy project, just to see if we should skip it
				{
					_, opt := loadFixture(testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, fixture.FullName, opt.Create)
				}

				h := create.NewHarness(t, ctx)
				project := h.Project

				primaryResource, opt := loadFixture(project)

				exportResources := []*unstructured.Unstructured{primaryResource}

				create.SetupNamespacesAndApplyDefaults(h, opt.Create, project)

				opt.CleanupResources = false // We delete explicitly below
				create.RunCreateDeleteTest(h, opt)

				for _, exportResource := range exportResources {
					exportURI := ""

					gvk := exportResource.GroupVersionKind()
					switch gvk.GroupKind() {
					case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
						name := exportResource.GetName()
						projectID := project.ProjectID
						exportURI = "//serviceusage.googleapis.com/projects/" + projectID + "/services/" + name
					}

					if exportURI == "" {
						continue
					}

					exportParams := h.ExportParams()
					exportParams.IAMFormat = "partialpolicy"
					exportParams.ResourceFormat = "krm"
					outputDir := h.TempDir()
					outputPath := filepath.Join(outputDir, "export.yaml")
					exportParams.Output = outputPath
					exportParams.URI = exportURI
					if err := export.Execute(h.Ctx, &exportParams); err != nil {
						t.Errorf("error from export.Execute: %v", err)
						continue
					}

					expectedPath := filepath.Join(fixture.SourceDir, "export.yaml")
					output := h.MustReadFile(outputPath)
					h.CompareGoldenFile(expectedPath, string(output), h.IgnoreComments, h.ReplaceString(project.ProjectID, "example-project-id"))
				}

				create.DeleteResources(h, opt.Create)

				// Verify events against golden file
				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" {
					events := h.Events

					operationIDs := map[string]bool{}
					for _, req := range events.HTTPEvents {
						url := req.Request.URL
						id := ""
						if index := strings.Index(url, "/v3/operations/"); index != -1 {
							id = strings.TrimPrefix(url[index:], "/v3/operations/")
						}
						id = strings.TrimSuffix(id, "?alt=json")
						if id != "" {
							operationIDs[id] = true
						}
					}

					for _, req := range events.HTTPEvents {
						url := req.Request.URL
						for id := range operationIDs {
							if strings.Contains(url, "/operations/"+id) {
								url = strings.ReplaceAll(url, "/operations/"+id, "/operations/${operationID}")
							}
						}
						req.Request.URL = url
					}

					pathIDs := map[string]string{}
					for _, req := range events.HTTPEvents {
						if !strings.Contains(req.Request.URL, "/operations/${operationID}") {
							continue
						}
						responseBody := req.Response.ParseBody()
						if responseBody == nil {
							continue
						}
						name, _, _ := unstructured.NestedString(responseBody, "response", "name")
						if strings.HasPrefix(name, "tagKeys/") {
							pathIDs[name] = "tagKeys/${tagKeyId}"
						}
					}

					// Replace any dynamic IDs that appear in URLs
					for _, req := range events.HTTPEvents {
						url := req.Request.URL
						for k, v := range pathIDs {
							url = strings.ReplaceAll(url, "/"+k, "/"+v)
						}
						req.Request.URL = url
					}

					// TODO: Fix how we poll / wait for objects being ready.
					events.RemoveRequests(func(e *test.LogEntry) bool {
						if e.Response.StatusCode == 404 && e.Request.Method == "GET" {
							return true
						}
						return false
					})

					// Remove operation polling requests (ones where the operation is not ready)
					events.RemoveRequests(func(e *test.LogEntry) bool {
						if !strings.Contains(e.Request.URL, "/operations/${operationID}") {
							return false
						}
						responseBody := e.Response.ParseBody()
						if responseBody == nil {
							return false
						}
						done, _, _ := unstructured.NestedBool(responseBody, "done")
						return !done // remove if not done
					})

					jsonMutators := []test.JSONMutator{}
					addReplacement := func(path string, newValue string) {
						tokens := strings.Split(path, ".")
						jsonMutators = append(jsonMutators, func(obj map[string]any) {
							_, found, _ := unstructured.NestedString(obj, tokens...)
							if found {
								unstructured.SetNestedField(obj, newValue, tokens...)
							}
						})
					}

					addReplacement("uniqueId", "111111111111111111111")
					addReplacement("oauth2ClientId", "888888888888888888888")

					addReplacement("etag", "abcdef0123A=")
					addReplacement("serviceAccount.etag", "abcdef0123A=")
					addReplacement("response.etag", "abcdef0123A=")

					addReplacement("createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.createTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("updateTime", "2024-04-01T12:34:56.123456Z")
					addReplacement("response.updateTime", "2024-04-01T12:34:56.123456Z")

					events.PrettifyJSON(jsonMutators...)

					// Remove headers that just aren't very relevant to testing
					events.RemoveHTTPResponseHeader("Date")
					events.RemoveHTTPResponseHeader("Alt-Svc")

					got := events.FormatHTTP()
					expectedPath := filepath.Join(fixture.SourceDir, "_http.log")
					normalizers := []func(string) string{}
					normalizers = append(normalizers, h.IgnoreComments)
					normalizers = append(normalizers, h.ReplaceString(uniqueID, "${uniqueId}"))
					normalizers = append(normalizers, h.ReplaceString(project.ProjectID, "${projectId}"))
					normalizers = append(normalizers, h.ReplaceString(fmt.Sprintf("%d", project.ProjectNumber), "${projectNumber}"))
					for k, v := range pathIDs {
						normalizers = append(normalizers, h.ReplaceString(k, v))
					}
					for k := range operationIDs {
						normalizers = append(normalizers, h.ReplaceString("operations/"+k, "operations/${operationId}"))
					}
					h.CompareGoldenFile(expectedPath, got, normalizers...)
				}
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
