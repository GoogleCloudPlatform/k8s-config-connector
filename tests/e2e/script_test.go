// Copyright 2024 Google LLC
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
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	kccyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/yaml"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

// TestE2EScript runs a Scenario test that runs step-by-step.
// See testdata/scenarios/README.md for more information.
func TestE2EScript(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("RUN_E2E not set; skipping")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
	})

	type gvkNN struct {
		gvk schema.GroupVersionKind
		nn  types.NamespacedName
	}

	logCheckTimeout := 10 * time.Second
	t.Run("scenarios", func(t *testing.T) {
		scenarioDir := "testdata/scenarios"
		scenarioPaths := findScripts(t, scenarioDir)

		for _, scenarioPath := range scenarioPaths {
			scenarioPath := scenarioPath

			t.Run(scenarioPath, func(t *testing.T) {
				uniqueID := testvariable.NewUniqueID()

				// Quickly load the sample with a dummy project, just to see if we should skip it
				{
					dummy := loadScript(t, filepath.Join(scenarioDir, scenarioPath), uniqueID, testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789})
					create.MaybeSkip(t, dummy.Name, dummy.Objects)
				}

				h := create.NewHarness(ctx, t)
				project := h.Project
				script := loadScript(t, filepath.Join(scenarioDir, scenarioPath), uniqueID, project)

				create.SetupNamespacesAndApplyDefaults(h, script.Objects, project)

				var objectsToDelete []*unstructured.Unstructured
				t.Cleanup(func() {
					create.DeleteResources(h, create.CreateDeleteTestOptions{Create: objectsToDelete})
				})

				var eventsByStep [][]*test.LogEntry
				eventsBefore := h.Events.HTTPEvents
				captureHTTPLogEvents := func() {
					var stepEvents []*test.LogEntry
					for i := len(eventsBefore); i < len(h.Events.HTTPEvents); i++ {
						stepEvents = append(stepEvents, h.Events.HTTPEvents[i])
					}
					eventsByStep = append(eventsByStep, stepEvents)
					eventsBefore = h.Events.HTTPEvents
				}

				// tracks the set of all applied objects as keyed by a gvk, namespaced name tuple
				appliedObjects := map[gvkNN]*unstructured.Unstructured{}

				for i, obj := range script.Objects {
					testCommand := ""
					v, ok := obj.Object["TEST"]
					if ok {
						testCommand = v.(string)
					}
					if testCommand == "" {
						testCommand = "APPLY"
					}

					if obj.GroupVersionKind().Kind == "RunCLI" {
						argsObjects := obj.Object["args"].([]any)
						var args []string
						for _, arg := range argsObjects {
							args = append(args, arg.(string))
						}
						baseOutputPath := filepath.Join(script.SourceDir, fmt.Sprintf("_cli-%d-", i))
						runCLI(h, args, uniqueID, baseOutputPath)
						continue
					}

					if obj.GroupVersionKind().Kind == "MockGCPBackdoor" {
						if h.MockGCP != nil {
							service, _, _ := unstructured.NestedString(obj.Object, "service")
							verb, _, _ := unstructured.NestedString(obj.Object, "verb")

							if err := h.MockGCP.RunTestCommand(ctx, service, verb); err != nil {
								h.Fatalf("running test command: %v", err)
							}
						} else {
							h.T.Logf("skipping MockGCPBackdoor command, because not running against mockgcp")
						}

						captureHTTPLogEvents()
						continue
					}

					// Try to delete this object as part of cleanup
					objectsToDelete = append(objectsToDelete, obj)

					exportResource := obj.DeepCopy()
					shouldGetKubeObject := true
					v, ok = obj.Object["WRITE-KUBE-OBJECT"]
					if ok {
						shouldGetKubeObject = v.(bool)
					}

					k := gvkNN{
						gvk: obj.GroupVersionKind(),
						nn: types.NamespacedName{
							Name:      obj.GetName(),
							Namespace: obj.GetNamespace(),
						},
					}

					var targetStepForReadAndCompare int
					switch testCommand {
					case "APPLY":
						applyObject(h, obj)
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)
						appliedObjects[k] = obj

					case "APPLY-10-SEC":
						applyObject(h, obj)
						time.Sleep(10 * time.Second)

					case "READ-OBJECT":
						appliedObjects[k] = obj

					case "READ-OBJECT-AND-COMPARE-SPEC":
						v, ok := obj.Object["TARGET_STEP_FOR_READ_AND_COMPARE"]
						if !ok {
							t.Fatalf("did not find key TARGET_STEP_FOR_READ_AND_COMPARE in the READ-OBJECT-AND-COMPARE-SPEC step")
						}
						targetStepForReadAndCompare = int(v.(int64))
						if targetStepForReadAndCompare <= 0 {
							t.Fatalf("value of TARGET_STEP_FOR_READ_AND_COMPARE should be an integer > 0")
						}
						appliedObjects[k] = obj

					case "APPLY-NO-WAIT":
						applyObject(h, obj)
						appliedObjects[k] = obj
						exportResource = nil
						shouldGetKubeObject = false

					case "PATCH-EXTERNALLY-MANAGED-FIELDS":
						patchObjectWithExternallyManagedFields(h, obj)
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)

					case "TOUCH":
						// Force re-reconciliation with an annotation
						touchObject(h, obj)
						// Pause to allow re-reconciliation
						// (annotations don't change the generation, so we can't wait for observedGeneration)
						time.Sleep(2 * time.Second)

					case "DELETE":
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{obj}})
						exportResource = nil
						shouldGetKubeObject = false

					case "SLEEP":
						// Allow some time for reconcile
						// Maybe we should instead wait for observedState
						time.Sleep(2 * time.Second)

					case "DELETE-NO-WAIT":
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{obj}, SkipWaitForDelete: true})

						// Allow some time for reconcile
						// Maybe we should instead wait for observedState
						time.Sleep(2 * time.Second)

						// The object probably still exists (that's probably
						// why we're using DELETE-NO-WAIT), so export the
						// resource and the kube object.
					case "WAIT-FOR-HTTP-REQUEST":
						applyObject(h, obj)

						val, ok := obj.Object["VALUE_PRESENT"]
						if !ok {
							t.Fatalf("did not find key VALUE_PRESENT in step")
						}
						sval := val.(string)

						ticker := time.NewTicker(1 * time.Second)
						for {
							stopWaiting := false
							select {
							case <-time.After(logCheckTimeout):
								t.Fatalf("timed out looking for value %s in http log", sval)
								stopWaiting = true
							case <-ticker.C:
								// todo(acpana): find better asympotatic approach
								for _, l := range h.Events.HTTPEvents {
									if strings.Contains(l.Response.Body, sval) {
										stopWaiting = true
										break
									}
								}
							}
							if stopWaiting {
								break
							}
						}

					case "ABANDON":
						setAnnotation(h, obj, "cnrm.cloud.google.com/deletion-policy", "abandon")
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{obj}})
						// continue to export the resource
						shouldGetKubeObject = false

					case "ABANDON-AND-REACQUIRE":
						existing := readObject(h, obj.GroupVersionKind(), obj.GetNamespace(), obj.GetName())
						resourceID, _, _ := unstructured.NestedString(existing.Object, "spec", "resourceID")
						if resourceID == "" {
							h.Fatalf("object did not have spec.resource: %v", existing)
						}
						setAnnotation(h, obj, "cnrm.cloud.google.com/deletion-policy", "abandon")
						deleteObj := obj.DeepCopy()
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{deleteObj}})
						if err := unstructured.SetNestedField(obj.Object, resourceID, "spec", "resourceID"); err != nil {
							h.Fatalf("error setting spec.resourceID: %v", err)
						}
						applyObject(h, obj)
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)
						appliedObjects[k] = obj

					case "ABANDON-AND-REACQUIRE-WITH-GENERATED-ID":
						existing := readObject(h, obj.GroupVersionKind(), obj.GetNamespace(), obj.GetName())
						resourceID, _, _ := unstructured.NestedString(existing.Object, "spec", "resourceID")
						if resourceID == "" {
							h.Fatalf("object did not have spec.resource: %v", existing)
						}
						setAnnotation(h, obj, "cnrm.cloud.google.com/deletion-policy", "abandon")
						deleteObj := obj.DeepCopy()
						create.DeleteResources(h, create.CreateDeleteTestOptions{Create: []*unstructured.Unstructured{deleteObj}})
						configuredID, _, _ := unstructured.NestedString(obj.Object, "spec", "resourceID")
						if configuredID == "" {
							h.Fatalf("object does not have resourceID configured: %v", obj)
						}
						err := unstructured.SetNestedField(obj.Object, strings.ReplaceAll(configuredID, "${TEST_GENERATED_ID}", resourceID), "spec", "resourceID")
						if err != nil {
							h.Fatalf("error setting spec.resourceID: %v", err)
						}
						applyObject(h, obj)
						create.WaitForReady(h, create.DefaultWaitForReadyTimeout, obj)
						appliedObjects[k] = obj

					default:
						t.Errorf("unknown TEST command %q", testCommand)
						continue
					}

					if exportResource != nil {
						u := exportResourceAsUnstructured(h, exportResource)
						if u == nil {
							t.Logf("ignoring failure to export resource of gvk %v", exportResource.GroupVersionKind())
							// t.Errorf("failed to export resource of gvk %v", exportResource.GroupVersionKind())
						} else {
							if err := normalizeKRMObject(t, u, project, uniqueID); err != nil {
								t.Fatalf("error from normalizeObject: %v", err)
							}
							got, err := yaml.Marshal(u)
							if err != nil {
								t.Errorf("failed to convert kube object to yaml: %v", err)
							}

							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_export%d.yaml", i))
							normalizers := []func(string) string{
								IgnoreComments,
							}
							h.CompareGoldenFile(expectedPath, string(got), normalizers...)
						}
					}

					if shouldGetKubeObject {
						u := &unstructured.Unstructured{}
						u.SetGroupVersionKind(obj.GroupVersionKind())
						id := types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
						if err := h.GetClient().Get(ctx, id, u); err != nil {
							t.Errorf("failed to get kube object: %v", err)
						} else {
							if err := normalizeKRMObject(t, u, project, uniqueID); err != nil {
								t.Fatalf("error from normalizeObject: %v", err)
							}
							got, err := yaml.Marshal(u)
							if err != nil {
								t.Errorf("failed to convert kube object to yaml: %v", err)
							}
							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", i))
							normalizers := []func(string) string{
								IgnoreComments,
								IgnoreAnnotations(map[string]struct{}{
									"cnrm.cloud.google.com/mutable-but-unreadable-fields": {},
								}),
							}
							h.CompareGoldenFile(expectedPath, string(got), normalizers...)
							// Compares the kube object spec read at the current
							// step (which should be equivalent to the golden
							// file, i.e. "_object%02d.yaml") and the kube
							// object read at a different step.
							// targetStepForReadAndCompare contains the step to
							// compare with. The step number follows 1-based
							// numbering.
							if targetStepForReadAndCompare > 0 {
								wantPath := filepath.Join(script.SourceDir, fmt.Sprintf("_object%02d.yaml", targetStepForReadAndCompare-1))
								gotPath := expectedPath
								wantObj, err := getKubeObjectInStringFromFile(wantPath)
								if err != nil {
									h.Fatalf("%s", err.Error())
								}
								gotObj, err := getKubeObjectInStringFromFile(gotPath)
								if err != nil {
									h.Fatalf("%s", err.Error())
								}
								diff := getDiffInSpecs(wantObj, gotObj)
								if diff != "" {
									t.Errorf("unexpected diff when comparing with the kube object spec in step %v: %s", targetStepForReadAndCompare, diff)
								}
							}
						}
					}

					captureHTTPLogEvents()
				}

				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" || os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
					{
						x := NewNormalizer(uniqueID, project)

						for _, stepEvents := range eventsByStep {
							x.Preprocess(stepEvents)
						}

						for i, stepEvents := range eventsByStep {
							expectedPath := filepath.Join(script.SourceDir, fmt.Sprintf("_http%02d.log", i))
							NormalizeHTTPLog(t, stepEvents, h.RegisteredServices(), project, uniqueID, "", "")
							got := x.Render(stepEvents)
							h.CompareGoldenFile(expectedPath, got, IgnoreComments)
						}
					}
				}

				objSet := []*unstructured.Unstructured{}
				for k := range appliedObjects {
					objSet = append(objSet, appliedObjects[k])
				}
				create.DeleteResources(h, create.CreateDeleteTestOptions{Create: objSet})

				h.NoExtraGoldenFiles(filepath.Join(script.SourceDir, "_*.yaml"))
			})
		}
	})
}

func applyObject(h *create.Harness, obj *unstructured.Unstructured) {
	if err := h.GetClient().Patch(h.Ctx, removeTestFields(obj), client.Apply, client.FieldOwner("kcc-tests"), client.ForceOwnership); err != nil {
		h.Fatalf("error applying resource: %v", err)
	}
}

// removes fields like "TEST" from a copy of the provided unstructured.
func removeTestFields(obj *unstructured.Unstructured) *unstructured.Unstructured {
	o := obj.DeepCopy()

	delete(o.Object, "TEST")
	delete(o.Object, "VALUE_PRESENT")
	delete(o.Object, "WRITE-KUBE-OBJECT")
	delete(o.Object, "TARGET_STEP_FOR_READ_AND_COMPARE")

	return o
}

func patchObjectWithExternallyManagedFields(h *create.Harness, obj *unstructured.Unstructured) {
	if err := h.GetClient().Patch(h.Ctx, removeTestFields(obj), client.Apply, client.FieldOwner(k8s.ControllerManagedFieldManager)); err != nil {
		h.Fatalf("error updating resource with externally managed fields: %v", err)
	}
}

// touchObject sets a new annotation that forces a re-reconciliation
func touchObject(h *create.Harness, obj *unstructured.Unstructured) {
	existing := &unstructured.Unstructured{}
	{
		existing.SetGroupVersionKind(obj.GroupVersionKind())
		existing.SetName(obj.GetName())
		existing.SetNamespace(obj.GetNamespace())

		key := types.NamespacedName{
			Namespace: obj.GetNamespace(),
			Name:      obj.GetName(),
		}
		if err := h.GetClient().Get(h.Ctx, key, existing); err != nil {
			h.Fatalf("error getting object %v: %v", key, err)
		}
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(obj.GroupVersionKind())
	u.SetName(obj.GetName())
	u.SetNamespace(obj.GetNamespace())

	oldAnnotation := existing.GetAnnotations()["test.cnrm.cloud.google.com/reconcile-cookie"]
	if oldAnnotation == "" {
		u.SetAnnotations(map[string]string{
			"test.cnrm.cloud.google.com/reconcile-cookie": "v1",
		})
	} else {
		n, err := strconv.ParseInt(strings.TrimPrefix(oldAnnotation, "v"), 10, 64)
		if err != nil {
			h.Fatalf("could not parse annotation test.cnrm.cloud.google.com/reconcile-cookie=%q", oldAnnotation)
		}
		newAnnotation := fmt.Sprintf("v%d", n+1)
		u.SetAnnotations(map[string]string{
			"test.cnrm.cloud.google.com/reconcile-cookie": newAnnotation,
		})

	}

	if err := h.GetClient().Patch(h.Ctx, u, client.Apply, client.FieldOwner("kcc-test-touch")); err != nil {
		h.Fatalf("error doing object touch (setting annotation): %v", err)
	}
}

func setAnnotation(h *create.Harness, obj *unstructured.Unstructured, k, v string) {
	patch := &unstructured.Unstructured{}
	patch.SetGroupVersionKind(obj.GroupVersionKind())
	patch.SetNamespace(obj.GetNamespace())
	patch.SetName(obj.GetName())
	annotations := map[string]string{
		k: v,
	}
	patch.SetAnnotations(annotations)

	if err := h.GetClient().Patch(h.Ctx, patch, client.Apply, client.FieldOwner("kcc-tests-setannotation"), client.ForceOwnership); err != nil {
		h.Fatalf("error setting annotations on resource: %v", err)
	}
}

func readObject(h *create.Harness, gvk schema.GroupVersionKind, namespace, name string) *unstructured.Unstructured {
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(gvk)
	key := types.NamespacedName{Namespace: namespace, Name: name}

	if err := h.GetClient().Get(h.Ctx, key, obj); err != nil {
		h.Fatalf("error reading object %v %v on resource: %v", gvk, key, err)
	}
	return obj
}

func findScripts(t *testing.T, rootDir string) []string {
	var relPaths []string
	if err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == "script.yaml" {
			relPath, err := filepath.Rel(rootDir, filepath.Dir(path))
			if err != nil {
				return fmt.Errorf("getting relative path during directory walk: %w", err)
			}
			relPaths = append(relPaths, relPath)
		}
		return nil
	}); err != nil {
		t.Fatalf("error walking directory %q: %v", rootDir, err)
	}
	return relPaths
}

type Script struct {
	Name      string
	SourceDir string
	Objects   []*unstructured.Unstructured
}

func loadScript(t *testing.T, dir string, testID string, project testgcp.GCPProject) *Script {
	s := &Script{
		Name:      dir,
		SourceDir: dir,
	}
	b := test.MustReadFile(t, filepath.Join(dir, "script.yaml"))

	b = testcontroller.ReplaceTestVars(t, b, testID, project)

	// split into yaml objects
	yamls, err := kccyaml.SplitYAML(b)
	if err != nil {
		t.Fatalf("error splitting bytes into YAMLs: %v", err)
	}

	// Parse to objects
	var objects []*unstructured.Unstructured
	for _, y := range yamls {
		obj := &unstructured.Unstructured{}
		if err := yaml.Unmarshal(y, obj); err != nil {
			t.Fatalf("error parsing object: %v", err)
		}

		// Set namespace to match project
		if obj.GetNamespace() == "" &&
			!isConfigConnectorObject(obj.GroupVersionKind()) {
			obj.SetNamespace(project.ProjectID)
		}

		// Hack: set project-id because mockkubeapiserver does not support webhooks
		if obj.GetAnnotations()["cnrm.cloud.google.com/project-id"] == "" &&
			!isConfigConnectorObject(obj.GroupVersionKind()) &&
			!isConfigConnectorContextObject(obj.GroupVersionKind()) {
			annotations := obj.GetAnnotations()
			if annotations == nil {
				annotations = make(map[string]string)
			}
			annotations["cnrm.cloud.google.com/project-id"] = project.ProjectID
			obj.SetAnnotations(annotations)
		}

		objects = append(objects, obj)
	}
	s.Objects = objects

	return s
}

func isConfigConnectorObject(gvk schema.GroupVersionKind) bool {
	if gvk.Kind == "ConfigConnector" &&
		gvk.Group == "core.cnrm.cloud.google.com" {
		return true
	}
	return false
}

func isConfigConnectorContextObject(gvk schema.GroupVersionKind) bool {
	if gvk.Kind == "ConfigConnectorContext" &&
		gvk.Group == "core.cnrm.cloud.google.com" {
		return true
	}
	return false
}

func createKubeconfigFromRestConfig(restConfig *rest.Config) ([]byte, error) {
	clusters := make(map[string]*clientcmdapi.Cluster)
	clusters["default-cluster"] = &clientcmdapi.Cluster{
		Server:                   restConfig.Host,
		CertificateAuthorityData: restConfig.CAData,
	}
	contexts := make(map[string]*clientcmdapi.Context)
	contexts["default-context"] = &clientcmdapi.Context{
		Cluster:  "default-cluster",
		AuthInfo: "default-user",
	}
	authInfos := make(map[string]*clientcmdapi.AuthInfo)
	authInfos["default-user"] = &clientcmdapi.AuthInfo{
		ClientCertificateData: restConfig.CertData,
		ClientKeyData:         restConfig.KeyData,
	}
	clientConfig := clientcmdapi.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		Clusters:       clusters,
		Contexts:       contexts,
		CurrentContext: "default-context",
		AuthInfos:      authInfos,
	}
	return clientcmd.Write(clientConfig)
}

// runCLI runs the config-connector CLI tool with the specified arguments
func runCLI(h *create.Harness, args []string, uniqueID string, baseOutputPath string) {
	project := h.Project
	t := h.T

	var options cmd.TestInvocationOptions

	for i, arg := range args {
		// Replace any substitutions in the args
		arg = strings.ReplaceAll(arg, "${projectId}", project.ProjectID)
		arg = strings.ReplaceAll(arg, "${uniqueId}", uniqueID)
		args[i] = arg
	}

	// Split the args into flags and positional arguments, so we can add more flags

	// Add some flags for kubeconfig and impersonation
	{
		tempDir := t.TempDir()
		p := filepath.Join(tempDir, "kubeconfig")

		kubeconfig, err := createKubeconfigFromRestConfig(h.GetRESTConfig())
		if err != nil {
			t.Fatalf("error creating kubeconfig: %v", err)
		}
		if err := os.WriteFile(p, kubeconfig, 0644); err != nil {
			t.Fatalf("error writing kubeconfig to %q: %v", p, err)
		}

		args = append(args, "--kubeconfig="+p)
		args = append(args, "--as=admin")
		args = append(args, "--as-group=system:masters")
	}

	options.Args = []string{"config-connector"}
	options.Args = append(options.Args, args...)

	t.Logf("running cli with args %+v", options.Args)
	if err := cmd.ExecuteFromTest(&options); err != nil {
		t.Errorf("cli execution (args=%+v) failed: %v", options.Args, err)
	}

	stdout := options.Stdout.String()
	t.Logf("stdout: %v", stdout)
	stdout = strings.ReplaceAll(stdout, project.ProjectID, "${projectID}")
	stdout = strings.ReplaceAll(stdout, uniqueID, "${uniqueId}")
	test.CompareGoldenFile(t, baseOutputPath+"stdout.log", stdout)

	stderr := options.Stderr.String()
	t.Logf("stderr: %v", stderr)
	stderr = strings.ReplaceAll(stderr, project.ProjectID, "${projectID}")
	stderr = strings.ReplaceAll(stderr, uniqueID, "${uniqueId}")
	test.CompareGoldenFile(t, baseOutputPath+"stderr.log", stderr)
}

func getKubeObjectInStringFromFile(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("error converting path %q to absolute path: %w", path, err)
	}
	objInBytes, err := os.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %q: %w", absPath, err)
	}
	objInString := string(objInBytes)
	return objInString, nil
}

func getDiffInSpecs(wantObj, gotObj string) string {
	wantSpec := extractOutSpecFromKubeObjectStrings(wantObj)
	gotSpec := extractOutSpecFromKubeObjectStrings(gotObj)
	return cmp.Diff(wantSpec, gotSpec)
}

func extractOutSpecFromKubeObjectStrings(obj string) string {
	lines := strings.Split(obj, "\n")
	var specLine int
	var statusLine int
	for i, line := range lines {
		if strings.HasPrefix(line, "spec:") {
			specLine = i
		} else if strings.HasPrefix(line, "status:") {
			statusLine = i
		}
	}
	specLines := lines[specLine:statusLine]
	spec := strings.Join(specLines, "\n")
	return spec
}
